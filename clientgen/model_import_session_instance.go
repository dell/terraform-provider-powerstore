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

// ImportSessionInstance This resource type has queriable association from remote_system
type ImportSessionInstance struct {
	// Unique identifier of the import session.
	Id   *string                `json:"id,omitempty"`
	Type *ImportSessionTypeEnum `json:"type,omitempty"`
	// User-specified name of the import session.  This property supports case-insensitive filtering.
	Name *string `json:"name,omitempty"`
	// Global storage discovery iSCSI ip address that will be used for import workflow. The address can be an IPv4 address or FQDN (Fully Qualified Domain Name).  Was added in version 3.0.0.0.
	GlobalStorageDiscoveryAddress *string `json:"global_storage_discovery_address,omitempty"`
	// User-specified description of the import session.
	Description *string `json:"description,omitempty"`
	// Unique identifier of the storage system that contains the source volume or consistency group to be imported.
	RemoteSystemId *string `json:"remote_system_id,omitempty"`
	// Unique identifier of the volume or consistency group to be imported.
	SourceResourceId *string `json:"source_resource_id,omitempty"`
	// Unique identifier of the destination volume or volume group created as part of the import process.
	DestinationResourceId   *string                            `json:"destination_resource_id,omitempty"`
	DestinationResourceType *ImportDestinationResourceTypeEnum `json:"destination_resource_type,omitempty"`
	// For a volume that is part of a consistency group import, this value is the session identifier of the import session. For an individual volume import, this value is null.
	ParentSessionId *string                 `json:"parent_session_id,omitempty"`
	State           *ImportSessionStateEnum `json:"state,omitempty"`
	// When the import is in the Copy_In_Progress state, this value indicates the estimated time at which the data copy will complete. Before the import is in the Copy_In_Progress state, the value is null.
	EstimatedCompletionTimestamp *time.Time `json:"estimated_completion_timestamp,omitempty"`
	// When the import is in the Copy_In_Progress state, this value indicates the completion percent for the import. Before the import is in the Copy_In_Progress state, this value is 0. After the cutover or if there is a failure, this value is null.
	ProgressPercentage *int32 `json:"progress_percentage,omitempty"`
	// Average transfer rate of a data import operation in bytes/sec over the whole copy period. Before and after the import is in the Copy_In_Progress state, this value is null.
	AverageTransferRate *int64 `json:"average_transfer_rate,omitempty"`
	// Current transfer rate of a data import operation in bytes/sec. Before and after the import is in the Copy_In_Progress state, this value is null.
	CurrentTransferRate *int64 `json:"current_transfer_rate,omitempty"`
	// Unique identifier of the local protection policy in the PowerStore storage system that will be applied on an imported destination volume or consistency group after cutover. Only snapshot policies are supported in an import. Once the import completes, you can add a replication policy.
	ProtectionPolicyId *string `json:"protection_policy_id,omitempty"`
	// Unique identifier of the volume group to which the destination volume will be added, if any.
	VolumeGroupId *string `json:"volume_group_id,omitempty"`
	// Indicates whether the import session cutover is manual (true) or automatic (false).
	AutomaticCutover *bool `json:"automatic_cutover,omitempty"`
	// Date and time at which the import session is scheduled to run. The date is specified in ISO 8601 format with the time expressed in UTC format.
	ScheduledTimestamp *time.Time     `json:"scheduled_timestamp,omitempty"`
	Error              *ErrorInstance `json:"error,omitempty"`
	// Date and time when was the import was last updated. This value is updated each time the import job updates.
	LastUpdateTimestamp *time.Time `json:"last_update_timestamp,omitempty"`
	// Localized message string corresponding to type Was added in version 1.0.2.
	TypeL10n *string `json:"type_l10n,omitempty"`
	// Localized message string corresponding to destination_resource_type
	DestinationResourceTypeL10n *string `json:"destination_resource_type_l10n,omitempty"`
	// Localized message string corresponding to state
	StateL10n    *string               `json:"state_l10n,omitempty"`
	RemoteSystem *RemoteSystemInstance `json:"remote_system,omitempty"`
}
