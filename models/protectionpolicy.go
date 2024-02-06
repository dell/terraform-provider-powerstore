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

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtectionPolicy - protectionPolicy properties
type ProtectionPolicy struct {
	ID                   types.String `tfsdk:"id"`
	Name                 types.String `tfsdk:"name"`
	Description          types.String `tfsdk:"description"`
	Type                 types.String `tfsdk:"type"`
	IsReadOnly           types.Bool   `tfsdk:"is_read_only"`
	SnapshotRuleIDs      types.Set    `tfsdk:"snapshot_rule_ids"`
	ReplicationRuleIDs   types.Set    `tfsdk:"replication_rule_ids"`
	SnapshotRuleNames    types.Set    `tfsdk:"snapshot_rule_names"`
	ReplicationRuleNames types.Set    `tfsdk:"replication_rule_names"`
}
