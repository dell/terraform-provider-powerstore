/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Volumegroup - Volumegroup properties
type Volumegroup struct {
	ID                     types.String `tfsdk:"id"`
	Name                   types.String `tfsdk:"name"`
	Description            types.String `tfsdk:"description"`
	VolumeIDs              types.Set    `tfsdk:"volume_ids"`
	IsWriteOrderConsistent types.Bool   `tfsdk:"is_write_order_consistent"`
	ProtectionPolicyID     types.String `tfsdk:"protection_policy_id"`
	VolumeNames            types.Set    `tfsdk:"volume_names"`
	ProtectionPolicyName   types.String `tfsdk:"protection_policy_name"`
}

// VolumeGroupModifyLocal - VolumeGroupModifyLocal properties needed since gopowermax causes issue with setting Protection policy on patch for VolumeGroupSnapshots
type VolumeGroupModifyLocal struct {
	Description            string  `json:"description"`
	Name                   string  `json:"name,omitempty"`
	IsWriteOrderConsistent bool    `json:"is_write_order_consistent,omitempty"`
	ExpirationTimestamp    *string `json:"expiration_timestamp,omitempty"`
}
