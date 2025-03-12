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
	"fmt"

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
