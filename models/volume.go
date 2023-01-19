package models

import "github.com/hashicorp/terraform-plugin-framework/types"

// Volume - volume properties
type Volume struct {
	ID                       types.String  `tfsdk:"id"`
	Name                     types.String  `tfsdk:"name"`
	Size                     types.Float64 `tfsdk:"size"`
	CapacityUnit             types.String  `tfsdk:"capacity_unit"`
	HostID                   types.String  `tfsdk:"host_id"`
	HostName                 types.String  `tfsdk:"host_name"`
	HostGroupID              types.String  `tfsdk:"host_group_id"`
	HostGroupName            types.String  `tfsdk:"host_group_name"`
	LogicalUnitNumber        types.Int64   `tfsdk:"logical_unit_number"`
	VolumeGroupID            types.String  `tfsdk:"volume_group_id"`
	VolumeGroupName          types.String  `tfsdk:"volume_group_name"`
	MinimumSize              types.Int64   `tfsdk:"min_size"`
	SectorSize               types.Int64   `tfsdk:"sector_size"`
	Description              types.String  `tfsdk:"description"`
	ApplianceID              types.String  `tfsdk:"appliance_id"`
	ApplianceName            types.String  `tfsdk:"appliance_name"`
	ProtectionPolicyID       types.String  `tfsdk:"protection_policy_id"`
	ProtectionPolicyName     types.String  `tfsdk:"protection_policy_name"`
	PerformancePolicyID      types.String  `tfsdk:"performance_policy_id"`
	CreationTimeStamp        types.String  `tfsdk:"creation_timestamp"`
	IsReplicationDestination types.Bool    `tfsdk:"is_replication_destination"`
	NodeAffinity             types.String  `tfsdk:"node_affinity"`
	Type                     types.String  `tfsdk:"type"`
	WWN                      types.String  `tfsdk:"wwn"`
	State                    types.String  `tfsdk:"state"`
	LogicalUsed              types.Int64   `tfsdk:"logical_used"`
	AppType                  types.String  `tfsdk:"app_type"`
	AppTypeOther             types.String  `tfsdk:"app_type_other"`
	Nsid                     types.Int64   `tfsdk:"nsid"`
	Nguid                    types.String  `tfsdk:"nguid"`
}
