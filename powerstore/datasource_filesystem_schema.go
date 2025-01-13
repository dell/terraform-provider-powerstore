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

package powerstore

import "github.com/hashicorp/terraform-plugin-framework/datasource/schema"

// FileSystemDatasourceSchema is a function that returns the schema for filesystem
func FileSystemDatasourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"access_policy": schema.StringAttribute{
			MarkdownDescription: "Access Policy of the File System",
			Description:         "Access Policy of the File System",
			Computed:            true,
		},
		"access_type": schema.StringAttribute{
			MarkdownDescription: "Access Type of the File System",
			Description:         "Access Type of the File System",
			Computed:            true,
		},
		"config_type": schema.StringAttribute{
			MarkdownDescription: "Config Type of the File System",
			Description:         "Config Type of the File System",
			Computed:            true,
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "Description of the File System",
			Description:         "Description of the File System",
			Computed:            true,
		},
		"expiration_timestamp": schema.StringAttribute{
			MarkdownDescription: "Expiration Timestamp of the File System",
			Description:         "Expiration Timestamp of the File System",
			Computed:            true,
		},
		"filesystem_type": schema.StringAttribute{
			MarkdownDescription: "Filesystem Type of the File System",
			Description:         "Filesystem Type of the File System",
			Computed:            true,
		},
		"flr_attributes": schema.SingleNestedAttribute{
			MarkdownDescription: "Flr Attributes of the File System",
			Description:         "Flr Attributes of the File System",
			Computed:            true,
			Attributes:          FLRAttributeSchema(),
		},
		"folder_rename_policy": schema.StringAttribute{
			MarkdownDescription: "Folder Rename Policy of the File System",
			Description:         "Folder Rename Policy of the File System",
			Computed:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "ID of the File System",
			Description:         "ID of the File System",
			Computed:            true,
		},
		"is_async_m_time_enabled": schema.BoolAttribute{
			MarkdownDescription: "Is Async MTime Enabled of the File System",
			Description:         "Is Async MTime Enabled of the File System",
			Computed:            true,
		},
		"is_smb_no_notify_enabled": schema.BoolAttribute{
			MarkdownDescription: "Is Smb No Notify Enabled of the File System",
			Description:         "Is Smb No Notify Enabled of the File System",
			Computed:            true,
		},
		"is_smb_notify_on_access_enabled": schema.BoolAttribute{
			MarkdownDescription: "Is Smb Notify On Access Enabled of the File System",
			Description:         "Is Smb Notify On Access Enabled of the File System",
			Computed:            true,
		},
		"is_smb_notify_on_write_enabled": schema.BoolAttribute{
			MarkdownDescription: "Is Smb Notify On Write Enabled of the File System",
			Description:         "Is Smb Notify On Write Enabled of the File System",
			Computed:            true,
		},
		"is_smb_op_locks_enabled": schema.BoolAttribute{
			MarkdownDescription: "Is Smb Op Locks Enabled of the File System",
			Description:         "Is Smb Op Locks Enabled of the File System",
			Computed:            true,
		},
		"is_smb_sync_writes_enabled": schema.BoolAttribute{
			MarkdownDescription: "Is Smb Sync Writes Enabled of the File System",
			Description:         "Is Smb Sync Writes Enabled of the File System",
			Computed:            true,
		},
		"locking_policy": schema.StringAttribute{
			MarkdownDescription: "Locking Policy of the File System",
			Description:         "Locking Policy of the File System",
			Computed:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "Name of the File System",
			Description:         "Name of the File System",
			Computed:            true,
		},
		"nas_server_id": schema.StringAttribute{
			MarkdownDescription: "Nas Server ID of the File System",
			Description:         "Nas Server ID of the File System",
			Computed:            true,
		},
		"parent_id": schema.StringAttribute{
			MarkdownDescription: "Parent ID of the File System",
			Description:         "Parent ID of the File System",
			Computed:            true,
		},
		"protection_policy_id": schema.StringAttribute{
			MarkdownDescription: "Protection Policy ID of the File System",
			Description:         "Protection Policy ID of the File System",
			Computed:            true,
		},
		"size_total": schema.Int64Attribute{
			MarkdownDescription: "Size Total of the File System",
			Description:         "Size Total of the File System",
			Computed:            true,
		},
		"size_used": schema.Int64Attribute{
			MarkdownDescription: "Size Used of the File System",
			Description:         "Size Used of the File System",
			Computed:            true,
		},
		"smb_notify_on_change_dir_depth": schema.Int64Attribute{
			MarkdownDescription: "Smb Notify On Change Dir Depth of the File System",
			Description:         "Smb Notify On Change Dir Depth of the File System",
			Computed:            true,
		},
		"is_quota_enabled": schema.BoolAttribute{
			MarkdownDescription: "Is Quota Enabled of the File System",
			Description:         "Is Quota Enabled of the File System",
			Computed:            true,
		},
		"grace_period": schema.Int64Attribute{
			MarkdownDescription: "Grace Period of the File System",
			Description:         "Grace Period of the File System",
			Computed:            true,
		},
		"default_hard_limit": schema.Int64Attribute{
			MarkdownDescription: "Default Hard Limit of the File System",
			Description:         "Default Hard Limit of the File System",
			Computed:            true,
		},
		"default_soft_limit": schema.Int64Attribute{
			MarkdownDescription: "Default Soft Limit of the File System",
			Description:         "Default Soft Limit of the File System",
			Computed:            true,
		},
		"creation_timestamp": schema.StringAttribute{
			MarkdownDescription: "Creation Timestamp of the File System",
			Description:         "Creation Timestamp of the File System",
			Computed:            true,
		},
		"last_refresh_timestamp": schema.StringAttribute{
			MarkdownDescription: "Last Refresh Timestamp of the File System",
			Description:         "Last Refresh Timestamp of the File System",
			Computed:            true,
		},
		"last_writable_timestamp": schema.StringAttribute{
			MarkdownDescription: "Last Writable Timestamp of the File System",
			Description:         "Last Writable Timestamp of the File System",
			Computed:            true,
		},
		"is_modified": schema.BoolAttribute{
			MarkdownDescription: "Is Modified of the File System",
			Description:         "Is Modified of the File System",
			Computed:            true,
		},
		"creator_type": schema.StringAttribute{
			MarkdownDescription: "Creator Type of the File System",
			Description:         "Creator Type of the File System",
			Computed:            true,
		},
	}
}

// FLRAttributeSchema is a function that returns the schema for FLR attributes
func FLRAttributeSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"default_retention": schema.StringAttribute{
			MarkdownDescription: "Default Retention of the File System",
			Description:         "Default Retention of the File System",
			Computed:            true,
		},
		"maximum_retention": schema.StringAttribute{
			MarkdownDescription: "Maximum Retention of the File System",
			Description:         "Maximum Retention of the File System",
			Computed:            true,
		},
		"minimum_retention": schema.StringAttribute{
			MarkdownDescription: "Minimum Retention of the File System",
			Description:         "Minimum Retention of the File System",
			Computed:            true,
		},
		"mode": schema.StringAttribute{
			MarkdownDescription: "Mode of the File System",
			Description:         "Mode of the File System",
			Computed:            true,
		},
		"auto_lock": schema.BoolAttribute{
			MarkdownDescription: "Auto Lock of the File System",
			Description:         "Auto Lock of the File System",
			Computed:            true,
		},
		"auto_delete": schema.BoolAttribute{
			MarkdownDescription: "Auto Delete of the File System",
			Description:         "Auto Delete of the File System",
			Computed:            true,
		},
		"policy_interval": schema.Int64Attribute{
			MarkdownDescription: "Policy Interval of the File System",
			Description:         "Policy Interval of the File System",
			Computed:            true,
		},
		"has_protected_files": schema.BoolAttribute{
			MarkdownDescription: "Has Protected Files of the File System",
			Description:         "Has Protected Files of the File System",
			Computed:            true,
		},
		"clock_time": schema.StringAttribute{
			MarkdownDescription: "Clock Time of the File System",
			Description:         "Clock Time of the File System",
			Computed:            true,
		},
		"maximum_retention_date": schema.StringAttribute{
			MarkdownDescription: "Maximum Retention Date of the File System",
			Description:         "Maximum Retention Date of the File System",
			Computed:            true,
		},
	}
}
