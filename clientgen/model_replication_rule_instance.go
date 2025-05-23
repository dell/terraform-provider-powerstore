/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// ReplicationRuleInstance Replication rule instance. Values was added in 3.0.0.0: alert_threshold, is_read_only. This resource type has queriable associations from remote_system, replication_session, policy
type ReplicationRuleInstance struct {
	// Unique identifier of the replication rule.
	Id *string `json:"id,omitempty"`
	// Name of the replication rule.  This property supports case-insensitive filtering.
	Name *string  `json:"name,omitempty"`
	Rpo  *RPOEnum `json:"rpo,omitempty"`
	// Unique identifier of the remote system to which this replication rule will replicate the associated storage resources.
	RemoteSystemId *string `json:"remote_system_id,omitempty"`
	// Indicates whether this is a replica of a replication rule on a remote system that is the source of a replication session replicating a storage resource to the local system.
	IsReplica *bool `json:"is_replica,omitempty"`
	// Indicates whether this replication rule can be modified.  Was added in version 3.0.0.0.
	IsReadOnly *bool `json:"is_read_only,omitempty"`
	// Number of minutes the system will wait before generating a compliance alert when a replication session does not meet the RPO. By default, this will be set to the number of minutes in the configured RPO.
	AlertThreshold *int32               `json:"alert_threshold,omitempty"`
	ManagedBy      *PolicyManagedByEnum `json:"managed_by,omitempty"`
	// Unique identifier of the managing entity based on the value of the managed_by property, as shown below:   * User - Empty   * Metro - Unique identifier of the remote system where the policy was assigned.   * Replication - Unique identifier of the source remote system.   * VMware_vSphere - Unique identifier of the owning VMware vSphere/vCenter.  Was added in version 3.0.0.0.
	ManagedById *string `json:"managed_by_id,omitempty"`
	// Localized message string corresponding to rpo
	RpoL10n *string `json:"rpo_l10n,omitempty"`
	// Localized message string corresponding to managed_by Was added in version 3.0.0.0.
	ManagedByL10n *string               `json:"managed_by_l10n,omitempty"`
	RemoteSystem  *RemoteSystemInstance `json:"remote_system,omitempty"`
	// This is the inverse of the resource type replication_session association.
	ReplicationSessions []ReplicationSessionInstance `json:"replication_sessions,omitempty"`
	// List of the policies that are associated with this replication_rule.
	Policies []PolicyInstance `json:"policies,omitempty"`
}
