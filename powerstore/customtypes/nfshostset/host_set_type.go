/*
Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nfshostset

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.SetTypable = HostSetType{}
)

type HostSetType struct {
	basetypes.SetType
}

func NewHostSetType() HostSetType {
	return HostSetType{
		SetType: basetypes.SetType{
			ElemType: basetypes.StringType{},
		},
	}
}

func (t HostSetType) String() string {
	return "powerstore.HostSetType"
}

func (l HostSetType) ElementType() attr.Type {
	return basetypes.StringType{}
}

func (t HostSetType) ValueType(ctx context.Context) attr.Value {
	return NewHostSetValueNull()
}

// TerraformType returns the tftypes.Type that should be used to
// represent this type. This constrains what user input will be
// accepted and what kind of data can be set in state. The framework
// will use this to translate the AttributeType to something Terraform
// can understand.
func (t HostSetType) TerraformType(ctx context.Context) tftypes.Type {
	return tftypes.Set{
		ElementType: tftypes.String,
	}
}

func (t HostSetType) Equal(o attr.Type) bool {
	_, ok := o.(HostSetType)

	return ok
}

func (t HostSetType) ValueFromSet(ctx context.Context, in basetypes.SetValue) (basetypes.SetValuable, diag.Diagnostics) {
	if in.ElementType(ctx) != t.ElementType() {
		return nil, diag.Diagnostics{
			diag.NewErrorDiagnostic(
				"Invalid Set Type",
				"Only a set of strings is allowed. Received "+in.Type(ctx).String(),
			),
		}
	}
	if in.IsUnknown() {
		return NewHostSetValueUnknown(), nil
	}
	if in.IsNull() {
		return NewHostSetValueNull(), nil
	}

	return NewHostSetValue(in.Elements())
}

func (t HostSetType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	// setValue := basetypes.NewSetValueMust(t.ElementType(), elems)
	setVal, err := t.SetType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}
	setValue := setVal.(basetypes.SetValue)

	setValuable, diags := t.ValueFromSet(ctx, setValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting SetValue to SetValuable: %+v", diags)
	}

	return setValuable, nil
}

func (t HostSetType) normalizeValues(in []attr.Value) ([]attr.Value, error) {
	ret := make([]attr.Value, 0, len(in))
	toNormalize := make([]string, 0, len(in))
	for _, val := range in {
		strVal := val.(basetypes.StringValue)
		if strVal.IsNull() || strVal.IsUnknown() {
			ret = append(ret, val)
			continue
		}
		str := strVal.ValueString()
		toNormalize = append(toNormalize, str)
	}
	normalized, err := t.normalizeStrings(toNormalize)
	if err != nil {
		return nil, err
	}
	for _, val := range normalized {
		ret = append(ret, basetypes.NewStringValue(val))
	}
	return ret, nil
}

func (t HostSetType) normalizeStrings(in []string) ([]string, error) {
	retCommon := make([]string, 0, len(in))
	cidrs, ips := make([]*net.IPNet, 0, len(in)), make([]net.IP, 0, len(in))
	var perr error
	for _, val := range in {
		// if val ends in /128, it is an IPv6 address with 128 mask
		if strings.HasSuffix(val, "/128") {
			ip := net.ParseIP(strings.TrimSuffix(val, "/128"))
			if ip == nil || ip.To16() == nil {
				perr = errors.Join(perr, fmt.Errorf("invalid IPv6 address entry %s", val))
				continue
			}
			ips = append(ips, ip)
		} else if strings.Contains(val, "/") {
			// if val contains /, it is a cidr
			// first check if both parts are valid IPv4
			// if so, convert to prefix length format
			splitCidr := strings.SplitN(val, "/", 2)
			if ip, mask := net.ParseIP(splitCidr[0]), net.ParseIP(splitCidr[1]); ip != nil && ip.To4() != nil && mask != nil && mask.To4() != nil {
				// check if mask is valid
				maskBytes := net.IPMask(mask.To4())
				ones, mlen := maskBytes.Size()
				if mlen == 0 {
					// return nil, fmt.Errorf("invalid IPv4 mask %s in CIDR entry %s", mask.String(), val)
					perr = errors.Join(perr, fmt.Errorf("invalid IPv4 mask %s in CIDR entry %s", mask.String(), val))
					continue
				}
				val = fmt.Sprintf("%s/%d", splitCidr[0], ones)
			}
			// convert to CIDR
			_, ipNet, err := net.ParseCIDR(val)
			if err != nil {
				perr = errors.Join(perr, fmt.Errorf("unable to parse CIDR: %w", err))
				continue
			}
			cidrs = append(cidrs, ipNet)
		} else if ipVal := net.ParseIP(val); ipVal != nil {
			// check if it is a valid IP address
			ips = append(ips, ipVal)
		} else {
			// these are custom hostnames, dns domains or netgroups
			retCommon = append(retCommon, val)
		}
	}
	// if any values were invalid, return error
	if perr != nil {
		return nil, perr
	}
	// deduplicate CIDRs by removing subnets
	uniqueCidrs := make([]*net.IPNet, 0, len(cidrs))
candidateLoop:
	for i := range cidrs {
		sizei, _ := cidrs[i].Mask.Size()
		for j := 0; j < len(cidrs); j++ {
			if i == j {
				continue
			}
			sizej, _ := cidrs[j].Mask.Size()
			if cidrs[j].Contains(cidrs[i].IP) && (sizej < sizei || (sizej == sizei && i > j)) { // superset exists
				continue candidateLoop
			}
		}
		uniqueCidrs = append(uniqueCidrs, cidrs[i])
	}
	// deduplicate IPs by removing those that are already in CIDRs
	uniqueIps := make([]net.IP, 0, len(ips))
ipcandidateLoop:
	for _, ip := range ips {
		for _, cidr := range uniqueCidrs {
			if cidr.Contains(ip) {
				continue ipcandidateLoop
			}
		}
		uniqueIps = append(uniqueIps, ip)
	}
	// merge CIDRs and IPs to ret
	ret := make([]string, 0, len(uniqueCidrs)+len(uniqueIps)+len(retCommon))
	ret = append(ret, retCommon...)
	for _, cidr := range uniqueCidrs {
		ret = append(ret, cidr.String())
	}
	for _, ip := range uniqueIps {
		ret = append(ret, ip.String())
	}
	// sort the strings
	return ret, nil
}

func (v HostSetType) equal(ins, outs []attr.Value) bool {
outerLoop:
	for _, elem := range ins {
		if elem.IsNull() || elem.IsUnknown() {
			return false
		}
		for _, out := range outs {
			if elem.Equal(out) {
				continue outerLoop
			}
		}
		return false
	}

	return true
}
