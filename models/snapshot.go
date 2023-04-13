package models

import "github.com/hashicorp/terraform-plugin-framework/types"

// Snapshot - Snapshot properties
type Snapshot struct {
	ID                  types.String `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	Description         types.String `tfsdk:"description"`
	PerformancePolicyID types.String `tfsdk:"performance_policy_id"`
	ExpirationTimestamp types.String `tfsdk:"expiration_timestamp"`
	CreatorType         types.String `tfsdk:"creator_type"`
	VolumeID            types.String `tfsdk:"volume_id"`
}

// VolumeGroupSnapshot - VolumeGroupSnapshot properties
type VolumeGroupSnapshot struct {
	ID                  types.String `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	Description         types.String `tfsdk:"description"`
	ExpirationTimestamp types.String `tfsdk:"expiration_timestamp"`
	VolumeGroupID       types.String `tfsdk:"volume_group_id"`
}

