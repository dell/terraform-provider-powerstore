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

// VolumeInstance Details about a volume, including snapshots and clones of volumes. This resource type has queriable associations from appliance, policy, migration_session, host_volume_mapping, remote_snapshot_session, remote_snapshot, volume_group, datastore, replication_session
type VolumeInstance struct {
	// Unique identifier of the volume instance.
	Id *string `json:"id,omitempty"`
	// Name of the volume. This value must contain 128 or fewer printable Unicode characters.   This property supports case-insensitive filtering.
	Name *string `json:"name,omitempty"`
	// Description of the volume. This value must contain 128 or fewer printable Unicode characters.
	Description *string         `json:"description,omitempty"`
	Type        *VolumeTypeEnum `json:"type,omitempty"`
	// World wide name of the volume.
	Wwn *string `json:"wwn,omitempty"`
	// NVMe Namespace unique identifier in the NVME subsystem. Used for volumes attached to NVMEoF hosts. Was added in version 2.0.0.0.
	Nsid *int32 `json:"nsid,omitempty"`
	// NVMe Namespace globally unique identifier. Used for volumes attached to NVMEoF hosts. Was added in version 2.0.0.0.
	Nguid *string `json:"nguid,omitempty"`
	// Unique identifier of the appliance on which the volume is provisioned.
	ApplianceId *string          `json:"appliance_id,omitempty"`
	State       *VolumeStateEnum `json:"state,omitempty"`
	//  Size of the volume in bytes. Minimum volume size is 1MB. Maximum volume size is 256TB. Size must be a multiple of 8192.
	Size *int64 `json:"size,omitempty"`
	// Current amount of data (in bytes) host has written to a volume without dedupe, compression or sharing. This metric applies to primaries, snaps and clones. The value is null initially when a volume is created and is collected at 5 minute intervals.  Was added in version 3.0.0.0.
	LogicalUsed  *int64            `json:"logical_used,omitempty"`
	NodeAffinity *NodeAffinityEnum `json:"node_affinity,omitempty"`
	// Time when the volume was created.
	CreationTimestamp *time.Time `json:"creation_timestamp,omitempty"`
	// Unique identifier of the protection policy assigned to the volume. Only applicable to primary and clone volumes.
	ProtectionPolicyId *string `json:"protection_policy_id,omitempty"`
	// Unique identifier of the performance policy assigned to the volume.
	PerformancePolicyId *string `json:"performance_policy_id,omitempty"`
	// Unique identifier of the QoS performance policy assigned to the volume. Was added in version 4.0.0.0.
	QosPerformancePolicyId *string `json:"qos_performance_policy_id,omitempty"`
	// Indicates whether this volume is a replication destination. This field is false on both ends when a volume is a metro volume. Areplication destination will be created by the system when a replication session is created. When there is an active replication session, all the user operations are restricted including modification, deletion, host operation, snapshot, clone, etc. After the replication session is deleted, the replication destination volume will remain as it is until the end user changes it to be a non-replication destination. After the change, it becomes a primary volume. If the end user keeps it as a replication destination, when the replication session is recreated, the replication destination volume could potentially be reused in the new session to avoid a time-consuming full sync. This property is only valid for primary and clone volumes.
	IsReplicationDestination *bool `json:"is_replication_destination,omitempty"`
	// Unique identifier of the migration session assigned to the volume if it is part of a migration activity.
	MigrationSessionId *string `json:"migration_session_id,omitempty"`
	// Unique identifier of the replication session assigned to the volume if it has been configured as a metro volume between two PowerStore clusters. The volume can only be modified, refreshed, or restored when the metro_replication_session is in the paused state.  Was added in version 3.0.0.0.
	MetroReplicationSessionId *string `json:"metro_replication_session_id,omitempty"`
	// Indicates whether the volume is available to host. This attribute is only applicable to primary volumes and clones.  Was added in version 3.0.0.0.
	IsHostAccessAvailable *bool                   `json:"is_host_access_available,omitempty"`
	ProtectionData        *ProtectionDataInstance `json:"protection_data,omitempty"`
	// Filtering on the fields of this embedded resource is not supported.
	LocationHistory []LocationHistoryInstance `json:"location_history,omitempty"`
	AppType         *AppTypeEnum              `json:"app_type,omitempty"`
	// An optional field used to describe application type usage for a volume. This field can only be set if app_type is set to Relational_Databases_Other, Big_Data_Analytics_Other, Business_Applications_Other, Healthcare_Other, Virtualization_Other or Other. If the app_type attribute is set to anything other than one of these values, the attribute will be cleared.  Was added in version 2.1.0.0.
	AppTypeOther *string `json:"app_type_other,omitempty"`
	// Localized message string corresponding to type
	TypeL10n *string `json:"type_l10n,omitempty"`
	// Localized message string corresponding to state
	StateL10n *string `json:"state_l10n,omitempty"`
	// Localized message string corresponding to node_affinity
	NodeAffinityL10n *string `json:"node_affinity_l10n,omitempty"`
	// Localized message string corresponding to app_type Was added in version 2.1.0.0.
	AppTypeL10n          *string                   `json:"app_type_l10n,omitempty"`
	Appliance            *ApplianceInstance        `json:"appliance,omitempty"`
	ProtectionPolicy     *PolicyInstance           `json:"protection_policy,omitempty"`
	QosPerformancePolicy *PolicyInstance           `json:"qos_performance_policy,omitempty"`
	MigrationSession     *MigrationSessionInstance `json:"migration_session,omitempty"`
	// This is the inverse of the resource type host_volume_mapping association.
	MappedVolumes []HostVolumeMappingInstance `json:"mapped_volumes,omitempty"`
	// This is the inverse of the resource type remote_snapshot_session association.
	RemoteSnapshotSessions []RemoteSnapshotSessionInstance `json:"remote_snapshot_sessions,omitempty"`
	// This is the inverse of the resource type remote_snapshot_session association.
	CurrentRemoteSnapshotSessions []RemoteSnapshotSessionInstance `json:"current_remote_snapshot_sessions,omitempty"`
	// This is the inverse of the resource type remote_snapshot association.
	RemoteSnapshots []RemoteSnapshotInstance `json:"remote_snapshots,omitempty"`
	// List of the volume_groups that are associated with this volume.
	VolumeGroups []VolumeGroupInstance `json:"volume_groups,omitempty"`
	// List of the datastores that are associated with this volume.
	Datastores []DatastoreInstance `json:"datastores,omitempty"`
	// List of the replication_sessions that are associated with this volume.
	ReplicationSessions []ReplicationSessionInstance `json:"replication_sessions,omitempty"`
}
