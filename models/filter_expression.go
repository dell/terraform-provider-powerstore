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

package models

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Filter Expression Value

var _ basetypes.StringValuable = FilterExpressionValue{}
var _ xattr.ValidateableAttribute = FilterExpressionValue{}

type FilterExpressionValue struct {
	basetypes.StringValue
}

func (v FilterExpressionValue) Equal(o attr.Value) bool {
	other, ok := o.(FilterExpressionValue)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v FilterExpressionValue) Type(ctx context.Context) attr.Type {
	// FilterExpressionType defined in the schema type section
	return FilterExpressionType{}
}

// Implementation of the xattr.ValidateableAttribute interface
func (v FilterExpressionValue) ValidateAttribute(ctx context.Context, req xattr.ValidateAttributeRequest, resp *xattr.ValidateAttributeResponse) {
	if v.IsNull() || v.IsUnknown() {
		return
	}

	rawString := v.ValueString()
	if len(rawString) == 0 {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid PowerStore filter expression",
			"Expecting a string value that was expected to be in PowerStore filter expression format, got empty string value.",
		)
		return
	}
	values, err := url.ParseQuery(rawString)

	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid PowerStore filter expression",
			"An unexpected error occurred while converting a string value that was expected to be in PowerStore filter expression format: "+err.Error(),
		)

		return
	}

	// if any keys have empty values, throw an error
	for key, val := range values {
		if err := v.isValidQueryValue(val); err != nil {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid PowerStore filter value for query parameter: "+key,
				err.Error(),
			)
			return
		}
	}
}

func (v FilterExpressionValue) isValidQueryValue(in []string) error {
	for _, val := range in {
		if len(val) == 0 {
			return fmt.Errorf("empty query value provided, please provide queries in the key=value format where value is not empty string")
		}
	}
	return nil
}

func (v FilterExpressionValue) ValueQueries() url.Values {
	queries, _ := url.ParseQuery(v.ValueString())
	return queries
}

// Filter Expression Type
// Ensure the implementation satisfies the expected interfaces
var _ basetypes.StringTypable = FilterExpressionType{}

type FilterExpressionType struct {
	basetypes.StringType
	// ... potentially other fields ...
}

func (t FilterExpressionType) Equal(o attr.Type) bool {
	_, ok := o.(FilterExpressionType)
	return ok
}

func (t FilterExpressionType) String() string {
	return "PowerStore.FilterExpressionType"
}

func (t FilterExpressionType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	// FilterExpressionValue defined in the value type section
	value := FilterExpressionValue{
		StringValue: in,
	}

	return value, nil
}

func (t FilterExpressionType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	stringValue, _ := attrValue.(basetypes.StringValue)
	stringValuable, _ := t.ValueFromString(ctx, stringValue)
	return stringValuable, nil
}

func (t FilterExpressionType) ValueType(ctx context.Context) attr.Value {
	// FilterExpressionValue defined in the value type section
	return FilterExpressionValue{}
}
