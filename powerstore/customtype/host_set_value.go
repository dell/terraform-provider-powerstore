package customtype

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ basetypes.SetValuable                   = HostSetValue{}
	_ basetypes.SetValuableWithSemanticEquals = HostSetValue{}
)

type HostSetValue struct {
	basetypes.SetValue
}

// SetSemanticEquals implements basetypes.SetValuableWithSemanticEquals.
func (v HostSetValue) SetSemanticEquals(ctx context.Context, other basetypes.SetValuable) (bool, diag.Diagnostics) {
	o, diags := other.ToSetValue(ctx)
	if diags.HasError() {
		return false, diags
	}
	onorm, _ := NewHostSetType().normalizeValues(o.Elements())
	vnorm, _ := NewHostSetType().normalizeValues(v.Elements())

	log.Print("Inside SetSemantic Equals")
	log.Println("Equality value is ", NewHostSetType().equal(vnorm, onorm))

	return NewHostSetType().equal(vnorm, onorm), nil
}

func (v HostSetValue) Type(_ context.Context) attr.Type {
	return NewHostSetType()
}

// func (v HostSetValue) Equal(o attr.Value) bool {
// 	other, ok := o.(HostSetValue)

// 	if !ok {
// 		return false
// 	}

// 	// return v.SetValue.Equal(other.SetValue)
// 	if !((v.IsNull() && other.IsNull()) || (v.IsUnknown() && other.IsUnknown())) {
// 		return false
// 	}

// 	if v.IsNull() || v.IsUnknown() {
// 		return true
// 	}

// 	onorm, _ := NewHostSetType().normalizeValues(other.Elements())
// 	vnorm, _ := NewHostSetType().normalizeValues(v.Elements())

// 	return NewHostSetType().equal(vnorm, onorm)
// }

// func (v HostSetValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
// 	return v.SetValue.ToTerraformValue(ctx)
// }

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
	// elements, err := NewHostSetType().normalizeValues(elements)
	// if err != nil {
	// 	return HostSetValue{}, diag.Diagnostics{
	// 		diag.NewErrorDiagnostic(
	// 			"Invalid Host Set Value",
	// 			err.Error(),
	// 		),
	// 	}
	// }
	SetValue, diags := basetypes.NewSetValue(basetypes.StringType{}, elements)

	return HostSetValue{
		SetValue: SetValue,
	}, diags
}

// func NewHostSetValueFullyKnown(elements []string) HostSetValue {
// 	elemAttrs := make([]attr.Value, 0, len(elements))
// 	for _, elem := range elements {
// 		elemAttrs = append(elemAttrs, elem)
// 	}
// 	SetValue, diags := basetypes.NewSetValue(basetypes.StringType{}, elemAttrs)

// 	return HostSetValue{
// 		SetValue: SetValue,
// 	}, diags
// }
