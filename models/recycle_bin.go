/*
Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

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

// RecycleBinConfigResource - Terraform state model for recycle_bin resource
type RecycleBinConfigResource struct {
	ID                 types.String `tfsdk:"id"`
	ExpirationDuration types.Int32  `tfsdk:"expiration_duration"`
}

// RecycleBinDataSourceModel - Terraform state model for recycle_bin data source
type RecycleBinDataSourceModel struct {
	ID              types.String          `tfsdk:"id"`
	ResourceType    types.String          `tfsdk:"resource_type"`
	Filters         FilterExpressionValue `tfsdk:"filter_expression"`
	RecycleBinItems []RecycleBinItemModel `tfsdk:"recycle_bin_items"`
}

// RecycleBinItemModel - Terraform state model for a single recycle bin item
type RecycleBinItemModel struct {
	ID                  types.String `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	ResourceType        types.String `tfsdk:"resource_type"`
	LogicalProvisioned  types.Int64  `tfsdk:"logical_provisioned"`
	LogicalUsed         types.Int64  `tfsdk:"logical_used"`
	ApplianceID         types.String `tfsdk:"appliance_id"`
	DeletionTimestamp   types.String `tfsdk:"deletion_timestamp"`
	ExpirationTimestamp types.String `tfsdk:"expiration_timestamp"`
	ResourceTypeL10N    types.String `tfsdk:"resource_type_l10n"`
}
