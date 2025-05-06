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

func PointerToStringType[T ~string](in *T) types.String {
	if in == nil {
		return types.StringNull()
	}
	return types.StringValue(string(*in))
}

func PTimeToStringType(in *time.Time) types.String {
	if in == nil {
		return types.StringNull()
	}
	return types.StringValue((*in).String())
}

func PointerToBoolType(in *bool) types.Bool {
	if in == nil {
		return types.BoolNull()
	}
	return types.BoolValue(*in)
}

func PointerToStruct[tfT any, jT any](in *jT, transform func(jT) tfT) tfT {
	if in == nil {
		var ret tfT
		return ret
	}
	return transform(*in)
}

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

func SliceTransform[tfT any, jT any](in []jT, transform func(jT) tfT) []tfT {
	ret := make([]tfT, len(in))
	for i, v := range in {
		ret[i] = transform(v)
	}
	return ret
}
