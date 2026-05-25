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

package helper

import (
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TfString - Converts *string to types.String, returns types.StringNull if input is nil
func TfString[T ~string](in *T) types.String {
	if in == nil {
		return types.StringNull()
	}
	return types.StringValue(string(*in))
}

// TfStringFromPTime - Converts *time.Time to types.String, returns types.StringNull if input is nil
func TfStringFromPTime(in *time.Time) types.String {
	if in == nil {
		return types.StringNull()
	}
	return types.StringValue((*in).String())
}

// TfBool - Converts *bool to types.Bool, returns types.BoolNull if input is nil
func TfBool(in *bool) types.Bool {
	if in == nil {
		return types.BoolNull()
	}
	return types.BoolValue(*in)
}

// TfInt32 - Converts *int32 to types.Int64, returns types.Int64Null if input is nil
func TfInt32(in *int32) types.Int64 {
	if in == nil {
		return types.Int64Null()
	}
	return types.Int64Value(int64(*in))
}

// TfInt64 - Converts *int64 to types.Int64, returns types.Int64Null if input is nil
func TfInt64(in *int64) types.Int64 {
	if in == nil {
		return types.Int64Null()
	}
	return types.Int64Value(*in)
}

// TfObject - Converts input using the transform transform function, returns empty output if input is nil
func TfObject[tfT any, jT any](in *jT, transform func(jT) tfT) tfT {
	if in == nil {
		var ret tfT
		return ret
	}
	return transform(*in)
}

// ValueToPointer - Extracts Go value pointer from attr.Value
// Returns nil if input is not known
// Supported types: types.String, types.Bool
// We can add more types in the future when required
func ValueToPointer[T bool | string, VT attr.Value](in VT) *T {
	if in.IsNull() || in.IsUnknown() {
		return nil
	}
	var ret any
	switch inv := any(in).(type) {
	case types.String:
		ret = inv.ValueString()
	case types.Bool:
		ret = inv.ValueBool()
	}

	switch retv := ret.(type) {
	case T:
		return &retv
	}
	return nil
}

// SliceTransform - Applies the transform function to each element in a slice
func SliceTransform[tfT any, jT any](in []jT, transform func(jT) tfT) []tfT {
	ret := make([]tfT, len(in))
	for i, v := range in {
		ret[i] = transform(v)
	}
	return ret
}

// SetDefault - Returns pointer to default value if input is nil, otherwise returns input
func SetDefault[T any](in *T, defaultVal T) *T {
	if in != nil {
		return in
	}
	return &defaultVal
}

// StringPtr - Converts a string to *string
func StringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// BoolPtr - Converts a bool to *bool
func BoolPtr(b bool) *bool {
	return &b
}

// Int32Ptr - Converts int32 to *int32
func Int32Ptr(i int32) *int32 {
	return &i
}

// Int32Value - Converts *int32 to int32 with default 0
func Int32Value(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}
