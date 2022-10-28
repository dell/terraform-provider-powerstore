package models

import "github.com/hashicorp/terraform-plugin-framework/types"

// Volume - volume properties
type Volume struct {
	ID                       types.String `tfsdk:"id"`
	Name                     types.String `tfsdk:"name"`
	Size                     types.Int64  `tfsdk:"size"`
	HostID                   types.String `tfsdk:"host_id"`
	HostGroupID              types.String `tfsdk:"host_group_id"`
	LogicalUnitNumber        types.Int64  `tfsdk:"logical_unit_number"`
	VolumeGroupID            types.String `tfsdk:"volume_group_id"`
	MinimumSize              types.Int64  `tfsdk:"min_size"`
	SectorSize               types.Int64  `tfsdk:"sector_size"`
	Description              types.String `tfsdk:"description"`
	ApplianceID              types.String `tfsdk:"appliance_id"`
	ProtectionPolicyID       types.String `tfsdk:"protection_policy_id"`
	PerformancePolicyID      types.String `tfsdk:"performance_policy_id"`
	CreationTimeStamp        types.String `tfsdk:"creation_timestamp"`
	IsReplicationDestination types.Bool   `tfsdk:"is_replication_destination"`
	NodeAffinity             types.String `tfsdk:"node_affinity"`
	Type                     types.String `tfsdk:"type"`
	WWN                      types.String `tfsdk:"wwn"`
	State                    types.String `tfsdk:"state"`
	LogicalUsed              types.Int64  `tfsdk:"logical_used"`
	AppType                  types.String `tfsdk:"app_type"`
	AppTypeOther             types.String `tfsdk:"app_type_other"`
	Nsid                     types.Int64  `tfsdk:"nsid"`
	Nguid                    types.String `tfsdk:"nguid"`
}
