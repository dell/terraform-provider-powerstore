package models

import "github.com/hashicorp/terraform-plugin-framework/types"

// VolumeDataSource - Volume DataSource properties
type VolumeDataSource struct {
	ID                        types.String        `tfsdk:"id"`
	Name                      types.String        `tfsdk:"name"`
	Size                      types.Float64       `tfsdk:"size"`
	CapacityUnit              types.String        `tfsdk:"capacity_unit"`
	HostID                    types.String        `tfsdk:"host_id"`
	HostGroupID               types.String        `tfsdk:"host_group_id"`
	LogicalUnitNumber         types.Int64         `tfsdk:"logical_unit_number"`
	VolumeGroupID             types.String        `tfsdk:"volume_group_id"`
	Description               types.String        `tfsdk:"description"`
	ApplianceID               types.String        `tfsdk:"appliance_id"`
	ProtectionPolicyID        types.String        `tfsdk:"protection_policy_id"`
	PerformancePolicyID       types.String        `tfsdk:"performance_policy_id"`
	CreationTimeStamp         types.String        `tfsdk:"creation_timestamp"`
	IsReplicationDestination  types.Bool          `tfsdk:"is_replication_destination"`
	NodeAffinity              types.String        `tfsdk:"node_affinity"`
	Type                      types.String        `tfsdk:"type"`
	WWN                       types.String        `tfsdk:"wwn"`
	State                     types.String        `tfsdk:"state"`
	LogicalUsed               types.Int64         `tfsdk:"logical_used"`
	AppType                   types.String        `tfsdk:"app_type"`
	AppTypeOther              types.String        `tfsdk:"app_type_other"`
	Nsid                      types.Int64         `tfsdk:"nsid"`
	Nguid                     types.String        `tfsdk:"nguid"`
	MigrationSessionID        types.String        `tfsdk:"migration_session_id"`
	MetroReplicationSessionID types.String        `tfsdk:"metro_replication_session_id"`
	IsHostAccessAvailable     types.Bool          `tfsdk:"is_host_access_available"`
	ProtectionData            ProtectionData      `tfsdk:"protection_data"`
	LocationHistory           []LocationHistory   `tfsdk:"location_history"`
	TypeL10n                  types.String        `tfsdk:"type_l10n"`
	StateL10n                 types.String        `tfsdk:"state_l10n"`
	NodeAffinityL10n          types.String        `tfsdk:"node_affinity_l10n"`
	AppTypeL10n               types.String        `tfsdk:"app_type_l10n"`
	Appliance                 Appliance           `tfsdk:"appliance"`
	ProtectionPolicy          VolProtectionPolicy `tfsdk:"protection_policy"`
	MigrationSession          MigrationSession    `tfsdk:"migration_session"`
	MappedVolumes             []MappedVolumes     `tfsdk:"mapped_volumes"`
	VolumeGroup               []VolumeGroup       `tfsdk:"volume_groups"`
	Datastores                []Datastores        `tfsdk:"datastores"`
}

// VolumeGroup - details of volume group for volume data source
type VolumeGroup struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

// Datastores - list of datastores for a volume in the volume data source
type Datastores struct {
	ID           types.String `tfsdk:"id"`
	InstanceUUID types.String `tfsdk:"instance_uuid"`
	Name         types.String `tfsdk:"name"`
}

// MappedVolumes - list of volume IDs mapped to a volume
type MappedVolumes struct {
	ID types.String `tfsdk:"id"`
}

// MigrationSession - details of migration session for a volume
type MigrationSession struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// Appliance - details of appliance attached to a volume
type Appliance struct {
	ID         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	ServiceTag types.String `tfsdk:"service_tag"`
}

// LocationHistory - details of location history for volume
type LocationHistory struct {
	FromApplianceID types.String `tfsdk:"from_appliance_id"`
	ToApplianceID   types.String `tfsdk:"to_appliance_id"`
	MigratedOn      types.String `tfsdk:"migrated_on"`
}

// ProtectionData - details of Protection data for volume
type ProtectionData struct {
	SourceID            types.String `tfsdk:"source_id"`
	CreatorType         types.String `tfsdk:"creator_type"`
	ExpirationTimestamp types.String `tfsdk:"expiration_timestamp"`
}

// VolProtectionPolicy - details of protection policy associated with volume
type VolProtectionPolicy struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}
