package customtype

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.ListTypable = (*HostsType)(nil)
)

type HostsType struct {
	basetypes.ListType
}

// String returns a human readable string of the type name.
func (t HostsType) String() string {
	return "powerstore.HostsType"
}

// ValueType returns the Value type.
func (t HostsType) ValueType(ctx context.Context) attr.Value {
	return Hosts{}
}

// Equal returns true if the given type is equivalent.
func (t HostsType) Equal(o attr.Type) bool {
	other, ok := o.(HostsType)

	if !ok {
		return false
	}

	return t.ListType.Equal(other.ListType)
}

// ValueFromString returns a ListValuable type given a ListValue.
func (t HostsType) ValueFromList(ctx context.Context, in basetypes.ListValue) (basetypes.ListValuable, diag.Diagnostics) {
	return Hosts{
		ListValue: in,
	}, nil
}

// ValueFromTerraform returns a Value given a tftypes.Value.  This is meant to convert the tftypes.Value into a more convenient Go type
// for the provider to consume the data with.
func (t HostsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.ListType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	ListValue, ok := attrValue.(basetypes.ListValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	listValuable, diags := t.ValueFromList(ctx, ListValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting ListValue to ListValuable: %v", diags)
	}

	return listValuable, nil
}
