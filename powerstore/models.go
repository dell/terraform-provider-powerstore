package powerstore

import (
	"net/http"
)

// Client - client structure for http client, host URL , username and password
type Client struct {
	HostURL    string       `json:"host_url,omitempty"`
	HTTPClient *http.Client `json:"http_client,omitempty"`
	Username   string       `json:"username,omitempty"`
	Password   string       `json:"password,omitempty"`
}

// Volume - volume properties
type Volume struct {
	ID                       string               `json:"id,omitempty"`
	Name                     string               `json:"name,omitempty"`
	Description              string               `json:"description,omitempty"`
	Type                     string               `json:"type,omitempty"`
	WWN                      string               `json:"wwn,omitempty"`
	NSID                     int                  `json:"nsid,omitempty"`
	NGUID                    string               `json:"nguid,omitempty"`
	ApplianceID              string               `json:"appliance_id,omitempty"`
	State                    string               `json:"state,omitempty"`
	Size                     int                  `json:"size,omitempty"`
	NodeAffinity             string               `json:"node_affinity,omitempty"`
	CreationTimeStamp        string               `json:"creation_timestamp,omitempty"`
	ProtectionPolicyID       string               `json:"protection_policy_id,omitempty"`
	PerformancePolicyID      string               `json:"performance_policy_id,omitempty"`
	IsReplicationDestination bool                 `json:"is_replication_destination,omitempty"`
	MigrationSessionID       string               `json:"migration_session_id,omitempty"`
	ProtectionData           ProtectionDataStruct `json:"protection_data,omitempty"`
	LocationHistory          string               `json:"location_history,omitempty"`
	TypeL10N                 string               `json:"type_l10n,omitempty"`
	StateL10N                string               `json:"state_l10n,omitempty"`
	NodeAffinityL10N         string               `json:"node_affinity_l10n,omitempty"`
}

// ProtectionDataStruct - part of volume properties
type ProtectionDataStruct struct {
	FamilyID            string `json:"family_id,omitempty"`
	ParentID            string `json:"parent_id,omitempty"`
	SoruceID            string `json:"source_id,omitempty"`
	CreatorType         string `json:"creator_type,omitempty"`
	CopySignature       string `json:"copy_signature,omitempty"`
	SourceTimeStamp     string `json:"source_timestamp,omitempty"`
	CreatorTypeL10N     string `json:"creator_type_l10n,omitempty"`
	IsAppConsistent     bool   `json:"is_app_consistent,omitempty"`
	CreatedByRuleID     string `json:"created_by_rule_id,omitempty"`
	CreatedByRuleName   string `json:"created_by_rule_name,omitempty"`
	ExpirationTimestamp string `json:"expiration_timestamp,omitempty"`
}

//SnapshotRule - snapshot rule resource schema
type SnapshotRule struct {
	ID               string   `json:"id,omitempty"`
	Name             string   `json:"name,omitempty"`
	Interval         string   `json:"interval,omitempty"`
	TimeOfDay        string   `json:"time_of_day,omitempty"`
	DaysOfWeek       []string `json:"days_of_week,omitempty"`
	DesiredRetention int      `json:"desired_retention,omitempty"`
	IsReplica        bool     `json:"is_replica,omitempty"`
	IntervalL10N     string   `json:"interval_l10n,omitempty"`
	DaysOfWeekL10N   []string `json:"days_of_week_l10n,omitempty"`
}

//StorageContainer - Storage Container resource
type StorageContainer struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Quota int    `json:"quota,omitempty"`
}

//VolRequest  - Volume Request structure
type VolRequest struct {
	ID                       string `json:"id,omitempty"`
	Name                     string `json:"name,omitempty"`
	Description              string `json:"description,omitempty"`
	IsReplicationDestination bool   `json:"is_replication_destination,omitempty"`
	ApplianceID              string `json:"appliance_id,omitempty"`
	Size                     int    `json:"size,omitempty"`
}

//ProtectionPolicy - Protection Policy resource
type ProtectionPolicy struct {
	ID               string            `json:"id,omitempty"`
	Name             string            `json:"name,omitempty"`
	Description      string            `json:"description,omitempty"`
	IsReplica        bool              `json:"is_replica,omitempty"`
	FileSystems      []FileSystems     `json:"file_systems,omitempty"`
	ReplicationRules []ReplicationRule `json:"replication_rules,omitempty"`
	SnapshotRules    []SnapshotRule    `json:"snapshot_rules,omitempty"`
	PolicyType       string            `json:"type,omitempty"`
	TypeL10n         string            `json:"type_l10n,omitempty"`
	VirtualMachines  []VirtualMachine  `json:"virtual_machines,omitempty"`
	VolumeGroups     []VolumeGroup     `json:"volume_groups,omitempty"`
	Volumes          []Volume          `json:"volumes,omitempty"`
}

//ProtectionPolicy - Protection Policy resource
type ProtectionPolicyRequest struct {
	ID                 string   `json:"id,omitempty"`
	Name               string   `json:"name,omitempty"`
	Description        string   `json:"description,omitempty"`
	ReplicationRuleIDs []string `json:"replication_rule_ids,omitempty"`
	SnapshotRulesIDs   []string `json:"snapshot_rule_ids,omitempty"`
}

type FileSystems struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ReplicationRule struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type VirtualMachine struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type VolumeGroup struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
