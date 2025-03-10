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
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// tfList is union(set, list)
type tfList interface {
	ElementsAs(context.Context, interface{}, bool) diag.Diagnostics
}

// TFListToSlice converts tfList (List or Set) to slice, ignores diagnostics
// use this if no errors can happen
func TFListToSlice[out any, in tfList](input in) []out {
	list, _ := TFListToSliceWithDiags[out](input)
	return list
}

// TFListToSliceWithDiags converts tfList (List or Set) to slice, returns diagnostics
func TFListToSliceWithDiags[out any, in tfList](input in) ([]out, diag.Diagnostics) {
	list := make([]out, 0)
	diags := input.ElementsAs(context.Background(), &list, false)
	return list, diags
}
