/*
Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package models

import "github.com/hashicorp/terraform-plugin-framework/types"

// MetroVolumeResource - Terraform state model for metro_volume resource
type MetroVolumeResource struct {
	ID                        types.String `tfsdk:"id"`
	VolumeID                  types.String `tfsdk:"volume_id"`
	RemoteSystemID            types.String `tfsdk:"remote_system_id"`
	RemoteApplianceID         types.String `tfsdk:"remote_appliance_id"`
	IsReplicationPaused       types.Bool   `tfsdk:"is_replication_paused"`
	DeleteRemoteVolume        types.Bool   `tfsdk:"delete_remote_volume"`
	Force                     types.Bool   `tfsdk:"force"`
	State                     types.String `tfsdk:"state"`
	RemoteResourceID          types.String `tfsdk:"remote_resource_id"`
	DataTransferState         types.String `tfsdk:"data_transfer_state"`
	MetroReplicationSessionID types.String `tfsdk:"metro_replication_session_id"`
}

// MetroVolumeGroupResource - Terraform state model for metro_volume_group resource
type MetroVolumeGroupResource struct {
	ID                        types.String `tfsdk:"id"`
	VolumeGroupID             types.String `tfsdk:"volume_group_id"`
	RemoteSystemID            types.String `tfsdk:"remote_system_id"`
	RemoteApplianceID         types.String `tfsdk:"remote_appliance_id"`
	IsReplicationPaused       types.Bool   `tfsdk:"is_replication_paused"`
	DeleteRemoteVolumeGroup   types.Bool   `tfsdk:"delete_remote_volume_group"`
	Force                     types.Bool   `tfsdk:"force"`
	State                     types.String `tfsdk:"state"`
	RemoteResourceID          types.String `tfsdk:"remote_resource_id"`
	DataTransferState         types.String `tfsdk:"data_transfer_state"`
	MetroReplicationSessionID types.String `tfsdk:"metro_replication_session_id"`
}

// ReplicationSessionActionResource - Terraform state model for replication_session_action resource
type ReplicationSessionActionResource struct {
	ID        types.String `tfsdk:"id"`
	SessionID types.String `tfsdk:"session_id"`
	Action    types.String `tfsdk:"action"`
	IsPlanned types.Bool   `tfsdk:"is_planned"`
	Reverse   types.Bool   `tfsdk:"reverse"`
	PostState types.String `tfsdk:"post_state"`
}

// ReplicationSessionDataSource - Terraform state model for replication_session data source
type ReplicationSessionDataSource struct {
	ID                  types.String                  `tfsdk:"id"`
	ReplicationSessions []ReplicationSessionItemModel `tfsdk:"replication_sessions"`
}

// ReplicationSessionItemModel - Terraform state model for a single replication session item
type ReplicationSessionItemModel struct {
	ID                     types.String `tfsdk:"id"`
	State                  types.String `tfsdk:"state"`
	Role                   types.String `tfsdk:"role"`
	ResourceType           types.String `tfsdk:"resource_type"`
	DataTransferState      types.String `tfsdk:"data_transfer_state"`
	Type                   types.String `tfsdk:"type"`
	LastSyncTimestamp      types.String `tfsdk:"last_sync_timestamp"`
	LocalResourceID        types.String `tfsdk:"local_resource_id"`
	RemoteResourceID       types.String `tfsdk:"remote_resource_id"`
	RemoteSystemID         types.String `tfsdk:"remote_system_id"`
	ProgressPercentage     types.Int64  `tfsdk:"progress_percentage"`
	ReplicationRuleID      types.String `tfsdk:"replication_rule_id"`
	LastSyncDuration       types.Int64  `tfsdk:"last_sync_duration"`
	FailoverTestInProgress types.Bool   `tfsdk:"failover_test_in_progress"`
	StateL10n              types.String `tfsdk:"state_l10n"`
	RoleL10n               types.String `tfsdk:"role_l10n"`
	ResourceTypeL10n       types.String `tfsdk:"resource_type_l10n"`
	TypeL10n               types.String `tfsdk:"type_l10n"`
}
