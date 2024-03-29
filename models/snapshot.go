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

import "github.com/hashicorp/terraform-plugin-framework/types"

// Snapshot - Snapshot properties
type Snapshot struct {
	ID                  types.String `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	Description         types.String `tfsdk:"description"`
	PerformancePolicyID types.String `tfsdk:"performance_policy_id"`
	ExpirationTimestamp types.String `tfsdk:"expiration_timestamp"`
	CreatorType         types.String `tfsdk:"creator_type"`
	VolumeID            types.String `tfsdk:"volume_id"`
	VolumeName          types.String `tfsdk:"volume_name"`
}

// VolumeGroupSnapshot - VolumeGroupSnapshot properties
type VolumeGroupSnapshot struct {
	ID                  types.String `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	Description         types.String `tfsdk:"description"`
	ExpirationTimestamp types.String `tfsdk:"expiration_timestamp"`
	VolumeGroupID       types.String `tfsdk:"volume_group_id"`
	VolumeGroupName     types.String `tfsdk:"volume_group_name"`
}
