/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

import (
	"time"
)

// VirtualVolumeInstance A virtual volume. This resource type has queriable associations from appliance, storage_container, replication_group, migration_session, virtual_volume, policy, host_virtual_volume_mapping, virtual_machine
type VirtualVolumeInstance struct {
	// The unique identifier of the virtual volume.
	Id *string `json:"id,omitempty"`
	// The name of the virtual volume, based on metadata provided by vSphere.   This property supports case-insensitive filtering.
	Name *string `json:"name,omitempty"`
	// The size of the virtual volume in bytes.
	Size      *int64                      `json:"size,omitempty"`
	Type      *VirtualVolumeTypeEnum      `json:"type,omitempty"`
	UsageType *VirtualVolumeUsageTypeEnum `json:"usage_type,omitempty"`
	// The appliance where the virtual volume resides.
	ApplianceId *string `json:"appliance_id,omitempty"`
	// The storage container where the virtual volume resides.
	StorageContainerId *string         `json:"storage_container_id,omitempty"`
	IoPriority         *IoPriorityEnum `json:"io_priority,omitempty"`
	// The ID of the storage profile governing this virtual volume.
	ProfileId *string `json:"profile_id,omitempty"`
	// The unique identifier of the replication group object that this virtual volume belongs to. Was added in version 3.0.0.0.
	ReplicationGroupId *string                 `json:"replication_group_id,omitempty"`
	CreatorType        *StorageCreatorTypeEnum `json:"creator_type,omitempty"`
	// Indicates whether the virtual volume is read-only.
	IsReadonly *bool `json:"is_readonly,omitempty"`
	// If the virtual volume is part of a migration activity, the session ID for that migration.
	MigrationSessionId *string `json:"migration_session_id,omitempty"`
	// UUID of the virtual machine that owns this virtual volume.
	VirtualMachineUuid *string `json:"virtual_machine_uuid,omitempty"`
	// Family id of the virtual volume. This is the id of the primary object at the root of the family tree. For a primary virtual volume this will be the same as the id of the object. For snap-sets and clone vVols it will be set to the source objects family ID.
	FamilyId *string `json:"family_id,omitempty"`
	// For snapshots and clones, the ID of the parent virtual volume. The parent_id is set when an virtual volume is created and will only change if its parent virtual volume is deleted.
	ParentId *string `json:"parent_id,omitempty"`
	// Id of the virtual volume from which the content has been sourced. Data is sourced from another virtual volume when a snapshot or clone is created, or when a refresh or restore occurs. Only applies to snap and clones.
	SourceId *string `json:"source_id,omitempty"`
	// The source data time-stamp of the virtual volume.
	SourceTimestamp *time.Time `json:"source_timestamp,omitempty"`
	// Timestamp of the moment virtual volume was created at.
	CreationTimestamp *time.Time `json:"creation_timestamp,omitempty"`
	// The NAA name used by hosts for I/O.  This is the VASA equivalent of a LUN's WWN. Was added in version 3.0.0.0.
	NaaName *string `json:"naa_name,omitempty"`
	// Indicates whether virtual volume is replication destination or not. Was added in version 3.0.0.0.
	IsReplicationDestination *bool `json:"is_replication_destination,omitempty"`
	// Filtering on the fields of this embedded resource is not supported.
	LocationHistory []LocationHistoryInstance `json:"location_history,omitempty"`
	// The unique identifier of the protection policy applied to this virtual volume. Was added in version 3.0.0.0.
	ProtectionPolicyId *string `json:"protection_policy_id,omitempty"`
	// NVMe Namespace unique identifier in the NVMe subsystem. Was added in version 3.0.0.0.
	Nsid *int32 `json:"nsid,omitempty"`
	// NVMe Namespace globally unique identifier. Was added in version 3.0.0.0.
	Nguid *string `json:"nguid,omitempty"`
	// Localized message string corresponding to type
	TypeL10n *string `json:"type_l10n,omitempty"`
	// Localized message string corresponding to usage_type
	UsageTypeL10n *string `json:"usage_type_l10n,omitempty"`
	// Localized message string corresponding to io_priority
	IoPriorityL10n *string `json:"io_priority_l10n,omitempty"`
	// Localized message string corresponding to creator_type
	CreatorTypeL10n  *string                   `json:"creator_type_l10n,omitempty"`
	Appliance        *ApplianceInstance        `json:"appliance,omitempty"`
	StorageContainer *StorageContainerInstance `json:"storage_container,omitempty"`
	ReplicationGroup *ReplicationGroupInstance `json:"replication_group,omitempty"`
	MigrationSession *MigrationSessionInstance `json:"migration_session,omitempty"`
	Parent           *VirtualVolumeInstance    `json:"parent,omitempty"`
	// This is the inverse of the resource type virtual_volume association.
	ChildVirtualVolumes []VirtualVolumeInstance `json:"child_virtual_volumes,omitempty"`
	Source              *VirtualVolumeInstance  `json:"source,omitempty"`
	// This is the inverse of the resource type virtual_volume association.
	TargetVirtualVolumes []VirtualVolumeInstance `json:"target_virtual_volumes,omitempty"`
	ProtectionPolicy     *PolicyInstance         `json:"protection_policy,omitempty"`
	// This is the inverse of the resource type host_virtual_volume_mapping association.
	HostVirtualVolumeMappings []HostVirtualVolumeMappingInstance `json:"host_virtual_volume_mappings,omitempty"`
	// List of the virtual_machines that are associated with this virtual_volume.
	VirtualMachines []VirtualMachineInstance `json:"virtual_machines,omitempty"`
}
