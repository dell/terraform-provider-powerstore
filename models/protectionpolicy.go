package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtectionPolicy - protectionPolicy properties
type ProtectionPolicy struct {
	ID                   types.String `tfsdk:"id"`
	Name                 types.String `tfsdk:"name"`
	Description          types.String `tfsdk:"description"`
	SnapshotRuleIDs      types.Set    `tfsdk:"snapshot_rule_ids"`
	ReplicationRuleIDs   types.Set    `tfsdk:"replication_rule_ids"`
	SnapshotRuleNames    types.Set    `tfsdk:"snapshot_rule_names"`
	ReplicationRuleNames types.Set    `tfsdk:"replication_rule_names"`
}
