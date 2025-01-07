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
			MarkdownDescription: "Access Policy",
			Description:         "Access Policy",
			Computed:            true,
		},
		"access_type": schema.StringAttribute{
			MarkdownDescription: "Access Type",
			Description:         "Access Type",
			Computed:            true,
		},
		"config_type": schema.StringAttribute{
			MarkdownDescription: "Config Type",
			Description:         "Config Type",
			Computed:            true,
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "Description",
			Description:         "Description",
			Computed:            true,
		},
		"expiration_timestamp": schema.StringAttribute{
			MarkdownDescription: "Expiration Timestamp",
			Description:         "Expiration Timestamp",
			Computed:            true,
		},
		"filesystem_type": schema.StringAttribute{
			MarkdownDescription: "Filesystem Type",
			Description:         "Filesystem Type",
			Computed:            true,
		},
		"flr_attributes": schema.SingleNestedAttribute{
			MarkdownDescription: "Flr Attributes",
			Description:         "Flr Attributes",
			Computed:            true,
			Attributes:          FLRAttributeSchema(),
		},
		"folder_rename_policy": schema.StringAttribute{
			MarkdownDescription: "Folder Rename Policy",
			Description:         "Folder Rename Policy",
			Computed:            true,
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "ID",
			Description:         "ID",
			Computed:            true,
		},
		"is_async_m_time_enabled": schema.BoolAttribute{
			MarkdownDescription: "Is Async MTime Enabled",
			Description:         "Is Async MTime Enabled",
			Computed:            true,
		},
		"is_smb_no_notify_enabled": schema.BoolAttribute{
			MarkdownDescription: "Is Smb No Notify Enabled",
			Description:         "Is Smb No Notify Enabled",
			Computed:            true,
		},
		"is_smb_notify_on_access_enabled": schema.BoolAttribute{
			MarkdownDescription: "Is Smb Notify On Access Enabled",
			Description:         "Is Smb Notify On Access Enabled",
			Computed:            true,
		},
		"is_smb_notify_on_write_enabled": schema.BoolAttribute{
			MarkdownDescription: "Is Smb Notify On Write Enabled",
			Description:         "Is Smb Notify On Write Enabled",
			Computed:            true,
		},
		"is_smb_op_locks_enabled": schema.BoolAttribute{
			MarkdownDescription: "Is Smb Op Locks Enabled",
			Description:         "Is Smb Op Locks Enabled",
			Computed:            true,
		},
		"is_smb_sync_writes_enabled": schema.BoolAttribute{
			MarkdownDescription: "Is Smb Sync Writes Enabled",
			Description:         "Is Smb Sync Writes Enabled",
			Computed:            true,
		},
		"locking_policy": schema.StringAttribute{
			MarkdownDescription: "Locking Policy",
			Description:         "Locking Policy",
			Computed:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "Name",
			Description:         "Name",
			Computed:            true,
		},
		"nas_server_id": schema.StringAttribute{
			MarkdownDescription: "Nas Server ID",
			Description:         "Nas Server ID",
			Computed:            true,
		},
		"parent_id": schema.StringAttribute{
			MarkdownDescription: "Parent ID",
			Description:         "Parent ID",
			Computed:            true,
		},
		"protection_policy_id": schema.StringAttribute{
			MarkdownDescription: "Protection Policy ID",
			Description:         "Protection Policy ID",
			Computed:            true,
		},
		"size_total": schema.Int64Attribute{
			MarkdownDescription: "Size Total",
			Description:         "Size Total",
			Computed:            true,
		},
		"size_used": schema.Int64Attribute{
			MarkdownDescription: "Size Used",
			Description:         "Size Used",
			Computed:            true,
		},
		"smb_notify_on_change_dir_depth": schema.Int64Attribute{
			MarkdownDescription: "Smb Notify On Change Dir Depth",
			Description:         "Smb Notify On Change Dir Depth",
			Computed:            true,
		},
	}
}

// FLRAttributeSchema is a function that returns the schema for FLR attributes
func FLRAttributeSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"default_retention": schema.StringAttribute{
			MarkdownDescription: "Default Retention",
			Description:         "Default Retention",
			Computed:            true,
		},
		"maximum_retention": schema.StringAttribute{
			MarkdownDescription: "Maximum Retention",
			Description:         "Maximum Retention",
			Computed:            true,
		},
		"minimum_retention": schema.StringAttribute{
			MarkdownDescription: "Minimum Retention",
			Description:         "Minimum Retention",
			Computed:            true,
		},
		"mode": schema.StringAttribute{
			MarkdownDescription: "Mode",
			Description:         "Mode",
			Computed:            true,
		},
	}
}
