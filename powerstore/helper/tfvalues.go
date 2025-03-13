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

import "github.com/hashicorp/terraform-plugin-framework/attr"

// IsKnownValue returns true if the value is known and is not a null
// usefull to check in planmodifiers/validators if a value has been configured
func IsKnownValue(value attr.Value) bool {
	return !value.IsUnknown() && !value.IsNull()
}
