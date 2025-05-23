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

// ReplicationSessionInstance A replication session.  This resource type has queriable associations from remote_system, replication_rule, volume, volume_group
type ReplicationSessionInstance struct {
	// Unique identifier of the replication session instance.
	Id                *string                     `json:"id,omitempty"`
	State             *ReplicationStateEnum       `json:"state,omitempty"`
	Role              *ReplicationRoleEnum        `json:"role,omitempty"`
	ResourceType      *ReplicatedResourceTypeEnum `json:"resource_type,omitempty"`
	DataTransferState *DataTransferStateEnum      `json:"data_transfer_state,omitempty"`
	Type              *ReplicationSessionTypeEnum `json:"type,omitempty"`
	// Time of last successful synchronization. For metro type replication sessions, this property is updated only during asynchronous copy phases. This is not supported for Nas Server replication sessions.
	LastSyncTimestamp *time.Time `json:"last_sync_timestamp,omitempty"`
	// Unique identifier of the local storage resource for the replication session.
	LocalResourceId *string `json:"local_resource_id,omitempty"`
	// Unique identifier of the remote storage resource for the replication session.
	RemoteResourceId *string `json:"remote_resource_id,omitempty"`
	// Unique identifier of the remote system instance.
	RemoteSystemId *string `json:"remote_system_id,omitempty"`
	// Progress of the current replication operation. This value is only available from the source system for the replication session, and is not supported for Nas Server replication sessions.
	ProgressPercentage *int32 `json:"progress_percentage,omitempty"`
	// Estimated completion time of the current replication operation. This is not supported for Nas Server replication sessions.
	EstimatedCompletionTimestamp *time.Time `json:"estimated_completion_timestamp,omitempty"`
	// Associated replication rule instance if created by policy engine.
	ReplicationRuleId *string `json:"replication_rule_id,omitempty"`
	// Elapsed time of the last synchronization operation in milliseconds.  This is not supported for Nas Server replication sessions. For metro type replication sessions, this property is updated only during asynchronous copy phases.  Was added in version 2.0.0.0.
	LastSyncDuration *int32 `json:"last_sync_duration,omitempty"`
	// Estimated start time of the next automatic synchronization operation. This is applicable to asynchronous type replication sessions. This is not supported for Nas Server replication sessions.  Was added in version 2.0.0.0.
	NextSyncTimestamp *time.Time `json:"next_sync_timestamp,omitempty"`
	// List of storage element pairs for a replication session. For a volume or volume group replication session, the replicating storage elements are of type 'volume’. For a virtual volume replication session, the replicating storage elements are of type 'virtual volume’. For a volume group replication session, there will be as many pairs of storage elements as the number of volumes in the volume group. For volume/virtual volume replication session, there will be only one storage element pair.   Filtering on the fields of this embedded resource is not supported.
	StorageElementPairs []ReplicationElementPair `json:"storage_element_pairs,omitempty"`
	// Indicates whether a test failover is in progress on the destination system of this replication session. This is not supported for Nas Server replication sessions.  Was added in version 2.0.0.0.
	FailoverTestInProgress *bool `json:"failover_test_in_progress,omitempty"`
	// Error code for the asynchronous copy phase failure.  Was added in version 3.0.0.0.
	ErrorCode           *string                  `json:"error_code,omitempty"`
	DataConnectionState *DataConnectionStateEnum `json:"data_connection_state,omitempty"`
	// Parent Replication session identifier. This is only applicable for replication sessions of type file system.  Was added in version 3.0.0.0.
	ParentReplicationSessionId *string                           `json:"parent_replication_session_id,omitempty"`
	LocalResourceState         *ReplicationResourceStateEnum     `json:"local_resource_state,omitempty"`
	WitnessDetails             *ReplicationSessionWitnessDetails `json:"witness_details,omitempty"`
	// Localized message string corresponding to state
	StateL10n *string `json:"state_l10n,omitempty"`
	// Localized message string corresponding to role
	RoleL10n *string `json:"role_l10n,omitempty"`
	// Localized message string corresponding to resource_type
	ResourceTypeL10n *string `json:"resource_type_l10n,omitempty"`
	// Localized message string corresponding to data_transfer_state Was added in version 3.0.0.0.
	DataTransferStateL10n *string `json:"data_transfer_state_l10n,omitempty"`
	// Localized message string corresponding to type Was added in version 3.0.0.0.
	TypeL10n *string `json:"type_l10n,omitempty"`
	// Localized message string corresponding to data_connection_state Was added in version 3.0.0.0.
	DataConnectionStateL10n *string `json:"data_connection_state_l10n,omitempty"`
	// Localized message string corresponding to local_resource_state Was added in version 3.0.0.0.
	LocalResourceStateL10n *string                   `json:"local_resource_state_l10n,omitempty"`
	RemoteSystem           *RemoteSystemInstance     `json:"remote_system,omitempty"`
	MigrationSession       *MigrationSessionInstance `json:"migration_session,omitempty"`
	ReplicationRule        *ReplicationRuleInstance  `json:"replication_rule,omitempty"`
	// List of the volumes that are associated with this replication_session.
	Volumes []VolumeInstance `json:"volumes,omitempty"`
	// List of the volume_groups that are associated with this replication_session.
	VolumeGroups []VolumeGroupInstance `json:"volume_groups,omitempty"`
}
