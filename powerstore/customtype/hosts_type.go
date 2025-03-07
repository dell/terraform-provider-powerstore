package customtype

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.ListTypable  = (*HostsType)(nil)
	_ xattr.TypeWithValidate = (*HostsType)(nil)
)

type HostsType struct {
	basetypes.ListType
}

func (t HostsType) String() string {
	return "powerstore.Hosts"
}

func (l HostsType) ElementType() attr.Type {
	return basetypes.StringType{}
}

func (l HostsType) WithElementType(typ attr.Type) attr.TypeWithElementType {
	return HostsType{
		ListType: basetypes.ListType{
			ElemType: basetypes.StringType{},
		},
	}
}


func (t HostsType) ValueType(ctx context.Context) attr.Value {
	return Hosts{}
}

func (t HostsType) Equal(o attr.Type) bool {
	_, ok := o.(HostsType)

	return ok
}

func (t HostsType) Validate(ctx context.Context, in tftypes.Value, path path.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	if in.Type() == nil {
		return diags
	}

	return diags
}

func (t HostsType) ValueFromList(ctx context.Context, in basetypes.ListValue) (basetypes.ListValuable, diag.Diagnostics) {
	return Hosts{
		ListValue: in,
	}, nil
}

func (t HostsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return basetypes.NewListNull(t.ElementType()), nil
	}

	if !in.Type().Is(tftypes.List{}) {
		return nil, fmt.Errorf("can't use %s as value of List with ElementType %T", in.String(), t.ElementType())
	}

	if !in.IsKnown() {
		return basetypes.NewListUnknown(t.ElementType()), nil
	}

	if in.IsNull() {
		return basetypes.NewListNull(t.ElementType()), nil
	}

	val := []tftypes.Value{}
	err := in.As(&val)
	if err != nil {
		return nil, err
	}
	elems := make([]attr.Value, 0, len(val))
	for _, elem := range val {
		av, err := t.ElementType().ValueFromTerraform(ctx, elem)
		if err != nil {
			return nil, err
		}
		elems = append(elems, av)
	}

	listValue := basetypes.NewListValueMust(t.ElementType(), elems)

	listValuable, diags := t.ValueFromList(ctx, listValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting ListValue to ListValuable: %v", diags)
	}

	return listValuable, nil
}
