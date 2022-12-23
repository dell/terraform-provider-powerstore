package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtectionPolicy - protectionPolicy properties
type ProtectionPolicy struct {
	ID                   types.String `tfsdk:"id"`
	Name                 types.String `tfsdk:"name"`
	Description          types.String `tfsdk:"description"`
	Type                 types.String `tfsdk:"type"`
	ManagedBy            types.String `tfsdk:"managed_by"`
	ManagedByID          types.String `tfsdk:"managed_by_id"`
	IsReadOnly           types.Bool   `tfsdk:"is_read_only"`
	IsReplica            types.Bool   `tfsdk:"is_replica"`
	SnapshotRuleIDs      types.List   `tfsdk:"snapshot_rule_ids"`
	ReplicationRuleIDs   types.List   `tfsdk:"replication_rule_ids"`
	SnapshotRuleNames    types.List   `tfsdk:"snapshot_rule_names"`
	ReplicationRuleNames types.List   `tfsdk:"replication_rule_names"`
}
