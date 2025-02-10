package customtype

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ basetypes.ListValuable                   = (*Hosts)(nil)
	_ basetypes.ListValuableWithSemanticEquals = (*Hosts)(nil)
	_ xattr.ValidateableAttribute              = (*Hosts)(nil)
	_ function.ValidateableParameter           = (*Hosts)(nil)
)

type Hosts struct {
	basetypes.ListValue
}

// Type returns an HostsType.
func (v Hosts) Type(_ context.Context) attr.Type {
	return HostsType{}
}

// Equal returns true if the given value is equivalent.
func (v Hosts) Equal(o attr.Value) bool {
	other, ok := o.(Hosts)

	if !ok {
		return false
	}

	return v.ListValue.Equal(other.ListValue)
}

func (v Hosts) ListSemanticEquals(_ context.Context, newValuable basetypes.ListValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(Hosts)
	if !ok {
		diags.AddError(
			"Semantic Equality Check Error",
			"An unexpected value type was received while performing semantic equality checks. "+
				"Please report this to the provider developers.\n\n"+
				"Expected Value Type: "+fmt.Sprintf("%T", v)+"\n"+
				"Got Value Type: "+fmt.Sprintf("%T", newValuable),
		)

		return false, diags
	}
	v.ElementsAs(context.Background(), nil, true)
	newValue.ElementsAs(context.Background(), nil, true)
	// newIpAddr, _ := netip.ParseAddr(newValue.ValueString())
	// currentIpAddr, _ := netip.ParseAddr(v.ValueString())

	return true, diags
}

func (v Hosts) ValidateAttribute(ctx context.Context, req xattr.ValidateAttributeRequest, resp *xattr.ValidateAttributeResponse) {
	if v.IsUnknown() || v.IsNull() {
		return
	}

	// ipAddr, err := netip.ParseAddr(v.ValueString())
	// if err != nil {
	// 	resp.Diagnostics.AddAttributeError(
	// 		req.Path,
	// 		"Invalid IPv6 Address String Value",
	// 		"A string value was provided that is not valid IPv6 string format (RFC 4291).\n\n"+
	// 			"Given Value: "+v.ValueString()+"\n"+
	// 			"Error: "+err.Error(),
	// 	)

	// 	return
	// }

	// if ipAddr.Is4() {
	// 	resp.Diagnostics.AddAttributeError(
	// 		req.Path,
	// 		"Invalid IPv6 Address String Value",
	// 		"An IPv4 string format was provided, string value must be IPv6 string format or IPv4-Mapped IPv6 string format (RFC 4291).\n\n"+
	// 			"Given Value: "+v.ValueString()+"\n",
	// 	)

	// 	return
	// }

	// if !ipAddr.IsValid() || !ipAddr.Is6() {
	// 	resp.Diagnostics.AddAttributeError(
	// 		req.Path,
	// 		"Invalid IPv6 Address String Value",
	// 		"A string value was provided that is not valid IPv6 string format (RFC 4291).\n\n"+
	// 			"Given Value: "+v.ValueString()+"\n",
	// 	)

	// 	return
	// }
}

func (v Hosts) ValidateParameter(ctx context.Context, req function.ValidateParameterRequest, resp *function.ValidateParameterResponse) {
	if v.IsUnknown() || v.IsNull() {
		return
	}

	// ipAddr, err := netip.ParseAddr(v.ValueString())
	// if err != nil {
	// 	resp.Error = function.NewArgumentFuncError(
	// 		req.Position,
	// 		"Invalid IPv6 Address String Value: "+
	// 			"A string value was provided that is not valid IPv6 string format (RFC 4291).\n\n"+
	// 			"Given Value: "+v.ValueString()+"\n"+
	// 			"Error: "+err.Error(),
	// 	)

	// 	return
	// }

	// if ipAddr.Is4() {
	// 	resp.Error = function.NewArgumentFuncError(
	// 		req.Position,
	// 		"Invalid IPv6 Address String Value: "+
	// 			"An IPv4 string format was provided, string value must be IPv6 string format or IPv4-Mapped IPv6 string format (RFC 4291).\n\n"+
	// 			"Given Value: "+v.ValueString()+"\n",
	// 	)

	// 	return
	// }

	// if !ipAddr.IsValid() || !ipAddr.Is6() {
	// 	resp.Error = function.NewArgumentFuncError(
	// 		req.Position,
	// 		"Invalid IPv6 Address String Value: "+
	// 			"A string value was provided that is not valid IPv6 string format (RFC 4291).\n\n"+
	// 			"Given Value: "+v.ValueString()+"\n",
	// 	)

	// 	return
	// }
}
