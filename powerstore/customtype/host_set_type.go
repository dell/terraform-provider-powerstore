package customtype

import (
	"context"
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

// func (t HostSetType) Validate(ctx context.Context, in tftypes.Value, path path.Path) diag.Diagnostics {
// 	var diags diag.Diagnostics

// 	if in.Type() == nil {
// 		return diags
// 	}

// 	return diags
// }

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
	// if in.Type() == nil {
	// 	return basetypes.NewSetNull(t.ElementType()), nil
	// }

	// if !in.Type().Equal(t.TerraformType(ctx)) {
	// 	return nil, fmt.Errorf("can't use %s as value of Set with ElementType %T, can only use %s values", in.String(), t.ElementType(), t.ElementType().TerraformType(ctx).String())
	// }

	// if !in.IsKnown() {
	// 	return basetypes.NewSetUnknown(t.ElementType()), nil
	// }

	// if in.IsNull() {
	// 	return basetypes.NewSetNull(t.ElementType()), nil
	// }

	// val := []tftypes.Value{}
	// err := in.As(&val)
	// if err != nil {
	// 	return nil, err
	// }
	// elems := make([]attr.Value, 0, len(val))
	// for _, elem := range val {
	// 	av, err := t.ElementType().ValueFromTerraform(ctx, elem)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	elems = append(elems, av)
	// }

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
	for _, val := range in {
		// if val contains /, it is a cidr
		if strings.Contains(val, "/") {
			// check if both parts are valid IPv4
			// if so, convert to prefix length format
			splitCidr := strings.SplitN(val, "/", 2)
			if ip, mask := net.ParseIP(splitCidr[0]), net.ParseIP(splitCidr[1]); ip != nil && ip.To4() != nil && mask != nil && mask.To4() != nil {
				// check if mask is valid
				maskBytes := net.IPMask(mask.To4())
				ones, mlen := maskBytes.Size()
				if mlen == 0 {
					return nil, fmt.Errorf("invalid IPv4 mask %s in CIDR entry %s", mask.String(), val)
				}
				val = fmt.Sprintf("%s/%d", splitCidr[0], ones)
			}
			// convert to CIDR
			_, ipNet, err := net.ParseCIDR(val)
			if err != nil {
				return nil, err
			}
			cidrs = append(cidrs, ipNet)
		} else if ipVal := net.ParseIP(val); ipVal != nil {
			// check if it is a valid IP address
			ips = append(ips, ipVal)
		} else {
			// these are custom hostnames
			retCommon = append(retCommon, val)
		}
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
			continue
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
