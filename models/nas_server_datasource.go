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

// NasServerConfigDataSource is the schema that is used to fetch NAS servers based on id or name
type NasServerConfigDataSource struct {
	Id         types.String          `tfsdk:"id"`
	Name       types.String          `tfsdk:"name"`
	NasServers []NasServerDataSource `tfsdk:"nas_servers"`
}

// NasServerDataSource represents the schema of a NAS server
type NasServerDataSource struct {
	Id                              types.String `tfsdk:"id"`
	Name                            types.String `tfsdk:"name"`
	Description                     types.String `tfsdk:"description"`
	OperationalStatus               types.String `tfsdk:"operational_status"`
	CurrentNodeId                   types.String `tfsdk:"current_node_id"`
	PreferredNodeId                 types.String `tfsdk:"preferred_node_id"`
	DefaultUnixUser                 types.String `tfsdk:"default_unix_user"`
	DefaultWindowsUser              types.String `tfsdk:"default_windows_user"`
	CurrentUnixDirectoryService     types.String `tfsdk:"current_unix_directory_service"`
	IsUsernameTranslationEnabled    types.Bool   `tfsdk:"is_username_translation_enabled"`
	IsAutoUserMappingEnabled        types.Bool   `tfsdk:"is_auto_user_mapping_enabled"`
	ProductionIPv4InterfaceId       types.String `tfsdk:"production_ipv4_interface_id"`
	ProductionIPv6InterfaceId       types.String `tfsdk:"production_ipv6_interface_id"`
	BackupIPv4InterfaceId           types.String `tfsdk:"backup_ipv4_interface_id"`
	BackupIPv6InterfaceId           types.String `tfsdk:"backup_ipv6_interface_id"`
	CurrentPreferredIPv4InterfaceId types.String `tfsdk:"current_preferred_ipv4_interface_id"`
	CurrentPreferredIPv6InterfaceId types.String `tfsdk:"current_preferred_ipv6_interface_id"`
	ProtectionPolicyId              types.String `tfsdk:"protection_policy_id"`
	FileEventsPublishingMode        types.String `tfsdk:"file_events_publishing_mode"`
	IsReplicationDestination        types.Bool   `tfsdk:"is_replication_destination"`
	IsProductionModeEnabled         types.Bool   `tfsdk:"is_production_mode_enabled"`
	IsDrTest                        types.Bool   `tfsdk:"is_dr_test"`
	OperationalStatusL10n           types.String `tfsdk:"operational_status_l10n"`
	CurrentUnixDirectoryServiceL10n types.String `tfsdk:"current_unix_directory_service_l10n"`
	FileEventsPublishingModeL10n    types.String `tfsdk:"file_events_publishing_mode_l10n"`
}
