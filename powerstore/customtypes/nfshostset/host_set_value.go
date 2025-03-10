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

	_, verr := NewHostSetType().normalizeValues(v.Elements())
	if verr != nil {
		resp.Diagnostics.AddAttributeError(req.Path, "Error parsing host values", verr.Error())
	}
}

// SetSemanticEquals implements basetypes.SetValuableWithSemanticEquals.
func (v HostSetValue) SetSemanticEquals(ctx context.Context, other basetypes.SetValuable) (bool, diag.Diagnostics) {
	o, _ := other.ToSetValue(ctx)
	// The only case where we get any error is if state has been imported (or is being read) and is in a format that we dont recognize.
	// If resource was created from Terraform, the state value must be valid and in a format that we recognize.
	// If resource plan is invalid, we would have got an error in the ValidateAttribute method.
	// In any case, better to fail than to silently ignore.
	onorm, oerr := NewHostSetType().normalizeValues(o.Elements())
	if oerr != nil {
		return false, helper.NewDiagnosticsFromError(
			"Got error normalizing values inside SetSemantic Equals.",
			oerr,
		)
	}
	vnorm, verr := NewHostSetType().normalizeValues(v.Elements())
	if verr != nil {
		return false, helper.NewDiagnosticsFromError(
			"Got error normalizing values inside SetSemantic Equals.",
			verr,
		)
	}

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

func NewHostSetValueFullyKnown(ctx context.Context, elements []string) HostSetValue {
	ret, _ := basetypes.NewSetValueFrom(ctx, basetypes.StringType{}, elements)
	return HostSetValue{SetValue: ret}
}
