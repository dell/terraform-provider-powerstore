package models

import "github.com/hashicorp/terraform-plugin-framework/types"

// VolumeSnapshotDataSource - Volume Snapshot DataSource properties
type VolumeSnapshotDataSource struct {
	ID                        types.String      `tfsdk:"id"`
	Name                      types.String      `tfsdk:"name"`
	Size                      types.Float64     `tfsdk:"size"`
	Description               types.String      `tfsdk:"description"`
	ApplianceID               types.String      `tfsdk:"appliance_id"`
	ProtectionPolicyID        types.String      `tfsdk:"protection_policy_id"`
	PerformancePolicyID       types.String      `tfsdk:"performance_policy_id"`
	CreationTimeStamp         types.String      `tfsdk:"creation_timestamp"`
	IsReplicationDestination  types.Bool        `tfsdk:"is_replication_destination"`
	NodeAffinity              types.String      `tfsdk:"node_affinity"`
	Type                      types.String      `tfsdk:"type"`
	WWN                       types.String      `tfsdk:"wwn"`
	State                     types.String      `tfsdk:"state"`
	LogicalUsed               types.Int64       `tfsdk:"logical_used"`
	AppType                   types.String      `tfsdk:"app_type"`
	AppTypeOther              types.String      `tfsdk:"app_type_other"`
	Nsid                      types.Int64       `tfsdk:"nsid"`
	Nguid                     types.String      `tfsdk:"nguid"`
	MigrationSessionID        types.String      `tfsdk:"migration_session_id"`
	MetroReplicationSessionID types.String      `tfsdk:"metro_replication_session_id"`
	IsHostAccessAvailable     types.Bool        `tfsdk:"is_host_access_available"`
	ProtectionData            ProtectionData    `tfsdk:"protection_data"`
	LocationHistory           []LocationHistory `tfsdk:"location_history"`
	TypeL10n                  types.String      `tfsdk:"type_l10n"`
	StateL10n                 types.String      `tfsdk:"state_l10n"`
	NodeAffinityL10n          types.String      `tfsdk:"node_affinity_l10n"`
	AppTypeL10n               types.String      `tfsdk:"app_type_l10n"`
}
