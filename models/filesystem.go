/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

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

// Volume - volume properties
type FileSystem struct {
	ID                         types.String  `tfsdk:"id"`
	Name                       types.String  `tfsdk:"name"`
	Size                       types.Float64 `tfsdk:"size"`
	CapacityUnit               types.String  `tfsdk:"capacity_unit"`
	Description                types.String  `tfsdk:"description"`
	NASServerID                types.String  `tfsdk:"nas_server_id"`
	ConfigType                 types.String  `tfsdk:"config_type"`
	AccessPolicy               types.String  `tfsdk:"access_policy"`
	LockingPolicy              types.String  `tfsdk:"locking_policy"`
	FolderRenamePolicy         types.String  `tfsdk:"folder_rename_policy"`
	IsAsyncMTimeEnabled        types.Bool    `tfsdk:"is_async_mtime_enabled"`
	ProtectionPolicyID         types.String  `tfsdk:"protection_policy_id"`
	FileEventsPublishingMode   types.String  `tfsdk:"file_events_publishing_mode"`
	HostIOSize                 types.String  `tfsdk:"host_io_size"`
	IsSmbSyncWritesEnabled     types.Bool    `tfsdk:"is_smb_sync_writes_enabled"`
	IsSmbNoNotifyEnabled       types.Bool    `tfsdk:"is_smb_no_notify_enabled"`
	IsSmbOpLocksEnabled        types.Bool    `tfsdk:"is_smb_op_locks_enabled"`
	IsSmbNotifyOnAccessEnabled types.Bool    `tfsdk:"is_smb_notify_on_access_enabled"`
	IsSmbNotifyOnWriteEnabled  types.Bool    `tfsdk:"is_smb_notify_on_write_enabled"`
	SmbNotifyOnChangeDirDepth  types.Int32   `tfsdk:"smb_notify_on_change_dir_depth"`
	FilesystemType             types.String  `tfsdk:"file_system_type"`
	ParentID                   types.String  `tfsdk:"parent_id"`
	FlrAttributes              types.Object  `tfsdk:"flr_attributes"`
}

type FlrAttributes struct {
	Mode types.String `tfsdk:"mode"`
}
