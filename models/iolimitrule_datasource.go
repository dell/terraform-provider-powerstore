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

// IoLimitRuleDataSource - io_limit_rule datasource properties
type IoLimitRuleDataSource struct {
	ID              types.String `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
	Type            types.String `tfsdk:"type"`
	MaxIops         types.Int32  `tfsdk:"max_iops"`
	MaxBw           types.Int32  `tfsdk:"max_bw"`
	BurstPercentage types.Int32  `tfsdk:"burst_percentage"`
	TypeL10n        types.String `tfsdk:"type_l10n"`
}

// IoLimitRuleDataSourceModel - datasource model wrapper
type IoLimitRuleDataSourceModel struct {
	IoLimitRules []IoLimitRuleDataSource `tfsdk:"io_limit_rules"`
	ID           types.String            `tfsdk:"id"`
	Name         types.String            `tfsdk:"name"`
	Type         types.String            `tfsdk:"type"`
	Filters      FilterExpressionValue   `tfsdk:"filter_expression"`
}
