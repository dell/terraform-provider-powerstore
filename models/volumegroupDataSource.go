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

// VolumeGroupDataSource - Volume Group DataSource properties
type VolumeGroupDataSource struct {
	ID                       types.String        `tfsdk:"id"`
	Name                     types.String        `tfsdk:"name"`
	Description              types.String        `tfsdk:"description"`
	CreationTimestamp        types.String        `tfsdk:"creation_timestamp"`
	IsProtectable            types.Bool          `tfsdk:"is_protectable"`
	ProtectionPolicyID       types.String        `tfsdk:"protection_policy_id"`
	MigrationSessionID       types.String        `tfsdk:"migration_session_id"`
	IsWriteOrderConsistent   types.Bool          `tfsdk:"is_write_order_consistent"`
	PlacementRule            types.String        `tfsdk:"placement_rule"`
	Type                     types.String        `tfsdk:"type"`
	IsReplicationDestination types.Bool          `tfsdk:"is_replication_destination"`
	ProtectionData           ProtectionData      `tfsdk:"protection_data"`
	IsImporting              types.Bool          `tfsdk:"is_importing"`
	LocationHistory          []LocationHistory   `tfsdk:"location_history"`
	TypeL10                  types.String        `tfsdk:"type_l10n"`
	ProtectionPolicy         VolProtectionPolicy `tfsdk:"protection_policy"`
	MigrationSession         MigrationSession    `tfsdk:"migration_session"`
	Volumes                  []Volumes           `tfsdk:"volumes"`
}
