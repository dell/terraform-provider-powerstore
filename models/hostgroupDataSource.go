/*
Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.

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

// HostGroupDataSource - HostGroupDataSource properties
type HostGroupDataSource struct {
	ID                        types.String                `tfsdk:"id"`
	Name                      types.String                `tfsdk:"name"`
	Description               types.String                `tfsdk:"description"`
	HostConnectivity          types.String                `tfsdk:"host_connectivity"`
	HostConnectivityL10n      types.String                `tfsdk:"host_connectivity_l10n"`
	Hosts                     []Hosts                     `tfsdk:"hosts"`
	MappedHostGroups          []MappedHostGroup           `tfsdk:"mapped_host_groups"`
	HostVirtualVolumeMappings []HostVirtualVolumeMappings `tfsdk:"host_virtual_volume_mappings"`
}

// MappedHostGroup - Details about a configured host or host group attached to a volume.
type MappedHostGroup struct {
	ID         types.String `tfsdk:"id"`
	HostID     types.String `tfsdk:"host_id"`
	VolumeID   types.String `tfsdk:"volume_id"`
	VolumeName types.String `tfsdk:"volume_name"`
}

// Hosts - Properties of a host.
type Hosts struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}
