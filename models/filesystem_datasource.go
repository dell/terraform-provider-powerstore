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

// FileSystemConfigDataSource is the schema for the providing the details in the config for fetching the filesystems based on id, name or nas server id.
type FileSystemConfigDataSource struct {
	ID          types.String           `tfsdk:"id"`
	Name        types.String           `tfsdk:"name"`
	NasServerID types.String           `tfsdk:"nas_server_id"`
	FileSystems []FileSystemDatasource `tfsdk:"filesystems"`
}

// FileSystemDatasource represents filesystem
type FileSystemDatasource struct {
	AccessPolicy               types.String            `tfsdk:"access_policy"`
	AccessType                 types.String            `tfsdk:"access_type"`
	ConfigType                 types.String            `tfsdk:"config_type"`
	Description                types.String            `tfsdk:"description"`
	ExpirationTimestamp        types.String            `tfsdk:"expiration_timestamp"`
	FilesystemType             types.String            `tfsdk:"filesystem_type"`
	FlrAttributes              FLRAttributesDatasource `tfsdk:"flr_attributes"`
	FolderRenamePolicy         types.String            `tfsdk:"folder_rename_policy"`
	ID                         types.String            `tfsdk:"id"`
	IsAsyncMTimeEnabled        types.Bool              `tfsdk:"is_async_m_time_enabled"`
	IsSmbNoNotifyEnabled       types.Bool              `tfsdk:"is_smb_no_notify_enabled"`
	IsSmbNotifyOnAccessEnabled types.Bool              `tfsdk:"is_smb_notify_on_access_enabled"`
	IsSmbNotifyOnWriteEnabled  types.Bool              `tfsdk:"is_smb_notify_on_write_enabled"`
	IsSmbOpLocksEnabled        types.Bool              `tfsdk:"is_smb_op_locks_enabled"`
	IsSmbSyncWritesEnabled     types.Bool              `tfsdk:"is_smb_sync_writes_enabled"`
	LockingPolicy              types.String            `tfsdk:"locking_policy"`
	Name                       types.String            `tfsdk:"name"`
	NasServerID                types.String            `tfsdk:"nas_server_id"`
	ParentID                   types.String            `tfsdk:"parent_id"`
	ProtectionPolicyID         types.String            `tfsdk:"protection_policy_id"`
	SizeTotal                  types.Int64             `tfsdk:"size_total"`
	SizeUsed                   types.Int64             `tfsdk:"size_used"`
	SmbNotifyOnChangeDirDepth  types.Int64             `tfsdk:"smb_notify_on_change_dir_depth"`
	IsQuotaEnabled             types.Bool              `tfsdk:"is_quota_enabled"`
	GracePeriod                types.Int64             `tfsdk:"grace_period"`
	DefaultHardLimit           types.Int64             `tfsdk:"default_hard_limit"`
	DefaultSoftLimit           types.Int64             `tfsdk:"default_soft_limit"`
	CreationTimestamp          types.String            `tfsdk:"creation_timestamp"`
	LastRefreshTimestamp       types.String            `tfsdk:"last_refresh_timestamp"`
	LastWritableTimestamp      types.String            `tfsdk:"last_writable_timestamp"`
	IsModified                 types.Bool              `tfsdk:"is_modified"`
	CreatorType                types.String            `tfsdk:"creator_type"`
	FileEventsPublishingMode   types.String            `tfsdk:"file_events_publishing_mode"`
	HostIOSize                 types.String            `tfsdk:"host_io_size"`
}

// FLRAttributesDatasource represents flr attributes
type FLRAttributesDatasource struct {
	DefaultRetention     types.String `tfsdk:"default_retention"`
	MaximumRetention     types.String `tfsdk:"maximum_retention"`
	MinimumRetention     types.String `tfsdk:"minimum_retention"`
	Mode                 types.String `tfsdk:"mode"`
	AutoLock             types.Bool   `tfsdk:"auto_lock"`
	AutoDelete           types.Bool   `tfsdk:"auto_delete"`
	PolicyInterval       types.Int64  `tfsdk:"policy_interval"`
	HasProtectedFiles    types.Bool   `tfsdk:"has_protected_files"`
	ClockTime            types.String `tfsdk:"clock_time"`
	MaximumRetentionDate types.String `tfsdk:"maximum_retention_date"`
}
