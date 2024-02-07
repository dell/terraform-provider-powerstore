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

// HostDataSource - HostDataSource properties
type HostDataSource struct {
	ID                        types.String                `tfsdk:"id"`
	Name                      types.String                `tfsdk:"name"`
	Description               types.String                `tfsdk:"description"`
	HostGroupID               types.String                `tfsdk:"host_group_id"`
	OsType                    types.String                `tfsdk:"os_type"`
	Initiators                []InitiatorInstance         `tfsdk:"initiators"`
	HostConnectivity          types.String                `tfsdk:"host_connectivity"`
	Type                      types.String                `tfsdk:"type"`
	TypeL10n                  types.String                `tfsdk:"type_l10n"`
	OsTypeL10n                types.String                `tfsdk:"os_type_l10n"`
	HostConnectivityL10n      types.String                `tfsdk:"host_connectivity_l10n"`
	ImportHostSystem          ImportHostSystem            `tfsdk:"import_host_system"`
	MappedHosts               []MappedHosts               `tfsdk:"mapped_hosts"`
	HostVirtualVolumeMappings []HostVirtualVolumeMappings `tfsdk:"host_virtual_volume_mappings"`
	VsphereHosts              []VsphereHosts              `tfsdk:"vsphere_hosts"`
}

// ImportHostSystem - Details about an import host system.
type ImportHostSystem struct {
	ID           types.String `tfsdk:"id"`
	AgentAddress types.String `tfsdk:"agent_address"`
}

// MappedHosts - Details about a configured host or host group attached to a volume.
type MappedHosts struct {
	ID       types.String `tfsdk:"id"`
	HostID   types.String `tfsdk:"host_id"`
	VolumeID types.String `tfsdk:"volume_id"`
}

// HostVirtualVolumeMappings - Virtual volume mapping details.
type HostVirtualVolumeMappings struct {
	ID                types.String `tfsdk:"id"`
	HostID            types.String `tfsdk:"host_id"`
	VirtualVolumeID   types.String `tfsdk:"virtual_volume_id"`
	VirtualVolumeName types.String `tfsdk:"virtual_volume_name"`
}

// VsphereHosts - List of the vsphere hosts that are associated with this host.
type VsphereHosts struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// InitiatorInstance - for reading initiator to the host
type InitiatorInstance struct {
	PortName           types.String `tfsdk:"port_name"`
	PortType           types.String `tfsdk:"port_type"`
	ChapMutualUsername types.String `tfsdk:"chap_mutual_username"`
	ChapSingleUsername types.String `tfsdk:"chap_single_username"`
}
