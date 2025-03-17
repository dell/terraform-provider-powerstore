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
	"reflect"
	"strings"

	"terraform-provider-powerstore/powerstore/helper"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ basetypes.SetValuable                   = HostSetValue{}
	_ basetypes.SetValuableWithSemanticEquals = HostSetValue{}
	_ xattr.ValidateableAttribute             = HostSetValue{}
)

type HostSetValue struct {
	basetypes.SetValue
}

// ValidateAttribute implements xattr.ValidateableAttribute.
func (v HostSetValue) ValidateAttribute(_ context.Context, req xattr.ValidateAttributeRequest, resp *xattr.ValidateAttributeResponse) {
	if v.IsNull() || v.IsUnknown() {
		return
	}

	_, verr, _ := v.normalizeValues()
	if verr != nil {
		resp.Diagnostics.AddAttributeError(req.Path, "Error parsing host values", verr.Error())
	}
}

// SetSemanticEquals implements basetypes.SetValuableWithSemanticEquals.
func (v HostSetValue) SetSemanticEquals(ctx context.Context, other basetypes.SetValuable) (bool, diag.Diagnostics) {
	oSet, _ := other.ToSetValue(ctx)
	o, _ := NewHostSetValue(oSet.Elements())
	// The only case where we get any error is if state has been imported (or is being read) and is in a format that we dont recognize.
	// If resource was created from Terraform, the state value must be valid and in a format that we recognize.
	// If resource plan is invalid, we would have got an error in the ValidateAttribute method.
	// In any case, better to fail than to silently ignore.
	onorm, oerr, ook := o.normalizeValues()
	if oerr != nil {
		return false, helper.NewDiagnosticsFromError(
			"Got error normalizing values inside SetSemantic Equals.",
			oerr,
		)
	}
	vnorm, verr, vok := v.normalizeValues()
	if verr != nil {
		return false, helper.NewDiagnosticsFromError(
			"Got error normalizing values inside SetSemantic Equals.",
			verr,
		)
	}

	// if either values lists have unknowns, they are not semantically equal
	if !ook || !vok {
		return false, nil
	}

	return v.equal(vnorm, onorm), nil
}

// normalizes value list by removing all unknown and null values
// also returns false in last return value if there are any unknown or null values
func (t HostSetValue) normalizeValues() (map[string]bool, error, bool) {
	in := t.Elements()
	ok := true
	toNormalize := make([]string, 0, len(in))
	for _, val := range in {
		strVal := val.(basetypes.StringValue)
		if strVal.IsNull() || strVal.IsUnknown() {
			ok = false
			continue
		}
		str := strVal.ValueString()
		toNormalize = append(toNormalize, str)
	}
	normalized, err := t.normalizeStrings(toNormalize)
	return normalized, err, ok
}

// normalizeStrings - normalizes list of IP addresses, CIDRs and custom name (as strings) by:
// - removing duplicates
// - removing IPs that are part of a CIDR that is already in the list
// - removing CIDRs that are proper subset of a CIDR that is already in the list
// - normalizing CIDRs to use <first-ip>/<prefix-length> format
// - converting CIDR in <ipv4>/<subnet mask as ipv4> format to <first-ip>/<prefix-length>
// - normalizing IPv6 addresses to remove the 128 mask
// - normalizing IPv6 addresses to remove excess zeros and replace by ::
//
// # Returns set of normalized strings
//
// Returns error if a CIDR (string has a "/" char) is not in valid format
// everything else is treated either as an ip address or custom name
//
// PowerStore accepts custom names for dns domains, hostnames and netgroups.
// Anything that is not a CIDR or an IP address is treated as a custom name.
func (t HostSetValue) normalizeStrings(in []string) (map[string]bool, error) {
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
		} else if strings.HasSuffix(val, "/255.255.255.255") {
			// if val ends in /255.255.255.255, it is an IPv4 address
			ip := net.ParseIP(strings.TrimSuffix(val, "/255.255.255.255"))
			if ip == nil || ip.To4() == nil {
				perr = errors.Join(perr, fmt.Errorf("invalid IPv4 address entry %s", val))
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
	// deduplicate strings and return map
	return t.deduplicateStrings(ret), nil
}

func (t HostSetValue) deduplicateStrings(in []string) map[string]bool {
	ret := make(map[string]bool)
	for _, val := range in {
		ret[val] = true
	}
	return ret
}

func (v HostSetValue) equal(ins, outs map[string]bool) bool {
	return reflect.DeepEqual(ins, outs)
}

func (v HostSetValue) Type(_ context.Context) attr.Type {
	return NewHostSetType()
}

func NewHostSetValueNull() HostSetValue {
	return HostSetValue{
		SetValue: basetypes.NewSetNull(basetypes.StringType{}),
	}
}

func NewHostSetValueUnknown() HostSetValue {
	return HostSetValue{
		SetValue: basetypes.NewSetUnknown(basetypes.NewStringUnknown().Type(context.Background())),
	}
}

func NewHostSetValue(elements []attr.Value) (HostSetValue, diag.Diagnostics) {
	SetValue, diags := basetypes.NewSetValue(basetypes.StringType{}, elements)

	return HostSetValue{
		SetValue: SetValue,
	}, diags
}

func NewHostSetValueFullyKnown(ctx context.Context, elements []string) HostSetValue {
	ret, _ := basetypes.NewSetValueFrom(ctx, basetypes.StringType{}, elements)
	return HostSetValue{SetValue: ret}
}
