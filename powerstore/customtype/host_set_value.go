package customtype

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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

	_, verr := NewHostSetType().normalizeValues(v.Elements())
	if verr != nil {
		resp.Diagnostics.AddAttributeError(req.Path, "Error parsing host values", verr.Error())
	}
}

// SetSemanticEquals implements basetypes.SetValuableWithSemanticEquals.
func (v HostSetValue) SetSemanticEquals(ctx context.Context, other basetypes.SetValuable) (bool, diag.Diagnostics) {
	o, _ := other.ToSetValue(ctx)
	onorm, oerr := NewHostSetType().normalizeValues(o.Elements())
	tflog.Debug(ctx, fmt.Sprintf("Got error normalizing values inside SetSemantic Equals for state: %s", oerr.Error()))
	vnorm, verr := NewHostSetType().normalizeValues(v.Elements())
	tflog.Debug(ctx, fmt.Sprintf("Got error normalizing values inside SetSemantic Equals for plan: %s", verr.Error()))

	return NewHostSetType().equal(vnorm, onorm), nil
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
