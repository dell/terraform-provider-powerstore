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

// QosPolicyDataSource - qos_policy datasource properties
type QosPolicyDataSource struct {
	ID                types.String `tfsdk:"id"`
	Name              types.String `tfsdk:"name"`
	Description       types.String `tfsdk:"description"`
	Type              types.String `tfsdk:"type"`
	ManagedBy         types.String `tfsdk:"managed_by"`
	ManagedById       types.String `tfsdk:"managed_by_id"`
	IsReadOnly        types.Bool   `tfsdk:"is_read_only"`
	IsReplica         types.Bool   `tfsdk:"is_replica"`
	TypeL10n          types.String `tfsdk:"type_l10n"`
	ManagedByL10n     types.String `tfsdk:"managed_by_l10n"`
	IoLimitRuleId     types.String `tfsdk:"io_limit_rule_id"`
	FileIoLimitRuleId types.String `tfsdk:"file_io_limit_rule_id"`
}

// QosPolicyDataSourceModel - datasource model wrapper
type QosPolicyDataSourceModel struct {
	QosPolicies []QosPolicyDataSource `tfsdk:"qos_policies"`
	ID          types.String          `tfsdk:"id"`
	Name        types.String          `tfsdk:"name"`
	Type        types.String          `tfsdk:"type"`
	Filters     FilterExpressionValue `tfsdk:"filter_expression"`
}
