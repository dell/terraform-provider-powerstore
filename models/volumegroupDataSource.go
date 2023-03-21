package models

import "github.com/hashicorp/terraform-plugin-framework/types"

// VolumeGroupDataSource - Volume Group DataSource properties
type VolumeGroupDataSource struct {
	ID                       types.String        `tfsdk:"id"`
	Name                     types.String        `tfsdk:"name"`
	Description              types.String        `tfsdk:"description"`
	CreationTimestamp        types.String        `tfsdk:"creation_timestamp"`
	IsProtectable            types.Bool          `tfsdk:"is_protectable"`
	ProtectionPolicyID       types.String        `tfsdk:"protection_policy_id"`
	MigrationSessionID       types.String        `tfsdk:"migration_session_id"`
	IsWriteOrderConsistent   types.Bool          `tfsdk:"is_write_order_consistent"`
	PlacementRule            types.String        `tfsdk:"placement_rule"`
	Type                     types.String        `tfsdk:"type"`
	IsReplicationDestination types.Bool          `tfsdk:"is_replication_destination"`
	ProtectionData           ProtectionData      `tfsdk:"protection_data"`
	IsImporting              types.Bool          `tfsdk:"is_importing"`
	LocationHistory          []LocationHistory   `tfsdk:"location_history"`
	TypeL10                  types.String        `tfsdk:"type_l10n"`
	ProtectionPolicy         VolProtectionPolicy `tfsdk:"protection_policy"`
	MigrationSession         MigrationSession    `tfsdk:"migration_session"`
	Volumes                  []Volumes           `tfsdk:"volumes"`
}
