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

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// NasServerDatasourcsSchema is a function that returns the schema for NAS server datasource
func NasServerDatasourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			MarkdownDescription: "Unique identifier of the NAS Server",
			Description:         "Unique identifier of the NAS Server",
			Computed:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "NAS Server name",
			Description:         "NAS Server name",
			Computed:            true,
		},
		"current_node_id": schema.StringAttribute{
			MarkdownDescription: "Current Node ID",
			Description:         "Current Node ID",
			Computed:            true,
		},
		"current_preferred_ipv4_interface_id": schema.StringAttribute{
			MarkdownDescription: "Current Preferred IPv4Interface ID",
			Description:         "Current Preferred IPv4Interface ID",
			Computed:            true,
		},
		"current_preferred_ipv6_interface_id": schema.StringAttribute{
			MarkdownDescription: "Current Preferred IPv6Interface ID",
			Description:         "Current Preferred IPv6Interface ID",
			Computed:            true,
		},
		"default_unix_user": schema.StringAttribute{
			MarkdownDescription: "Default Unix User",
			Description:         "Default Unix User",
			Computed:            true,
		},
		"default_windows_user": schema.StringAttribute{
			MarkdownDescription: "Default Windows User",
			Description:         "Default Windows User",
			Computed:            true,
		},
		"production_ipv4_interface_id": schema.StringAttribute{
			MarkdownDescription: "Production IPv4 Interface ID",
			Description:         "Production IPv4 Interface ID",
			Computed:            true,
		},
		"production_ipv6_interface_id": schema.StringAttribute{
			MarkdownDescription: "Production IPv6 Interface ID",
			Description:         "Production IPv6 Interface ID",
			Computed:            true,
		},
		"protection_policy_id": schema.StringAttribute{
			MarkdownDescription: "Protection Policy ID",
			Description:         "Protection Policy ID",
			Computed:            true,
		},
		"current_unix_directory_service": schema.StringAttribute{
			MarkdownDescription: "Current Unix Directory Service",
			Description:         "Current Unix Directory Service",
			Computed:            true,
		},
		"current_unix_directory_service_l10n": schema.StringAttribute{
			MarkdownDescription: "Current Unix Directory Service L10n",
			Description:         "Current Unix Directory Service L10n",
			Computed:            true,
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "Description",
			Description:         "Description",
			Computed:            true,
		},
		"backup_ipv4_interface_id": schema.StringAttribute{
			MarkdownDescription: "Backup IPv4 Interface ID",
			Description:         "Backup IPv4 Interface ID",
			Computed:            true,
		},
		"backup_ipv6_interface_id": schema.StringAttribute{
			MarkdownDescription: "Backup IPv6 Interface ID",
			Description:         "Backup IPv6 Interface ID",
			Computed:            true,
		},
		"file_events_publishing_mode": schema.StringAttribute{
			MarkdownDescription: "File Events Publishing Mode",
			Description:         "File Events Publishing Mode",
			Computed:            true,
		},
		"file_events_publishing_mode_l10n": schema.StringAttribute{
			MarkdownDescription: "File Events Publishing Mode L10n",
			Description:         "File Events Publishing Mode L10n",
			Computed:            true,
		},
		"is_auto_user_mapping_enabled": schema.BoolAttribute{
			MarkdownDescription: "Is Auto User Mapping Enabled",
			Description:         "Is Auto User Mapping Enabled",
			Computed:            true,
		},
		"is_dr_test": schema.BoolAttribute{
			MarkdownDescription: "Is DR Test",
			Description:         "Is DR Test",
			Computed:            true,
		},
		"is_production_mode_enabled": schema.BoolAttribute{
			MarkdownDescription: "Is Production Mode Enabled",
			Description:         "Is Production Mode Enabled",
			Computed:            true,
		},
		"is_replication_destination": schema.BoolAttribute{
			MarkdownDescription: "Is Replication Destination",
			Description:         "Is Replication Destination",
			Computed:            true,
		},
		"is_username_translation_enabled": schema.BoolAttribute{
			MarkdownDescription: "Is Username Translation Enabled",
			Description:         "Is Username Translation Enabled",
			Computed:            true,
		},
		"operational_status": schema.StringAttribute{
			MarkdownDescription: "Operational Status",
			Description:         "Operational Status",
			Computed:            true,
		},
		"operational_status_l10n": schema.StringAttribute{
			MarkdownDescription: "Operational Status L10n",
			Description:         "Operational Status L10n",
			Computed:            true,
		},
		"preferred_node_id": schema.StringAttribute{
			MarkdownDescription: "Preferred Node ID",
			Description:         "Preferred Node ID",
			Computed:            true,
		},
	}
}
