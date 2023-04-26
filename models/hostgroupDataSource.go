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
