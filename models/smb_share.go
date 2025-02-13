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

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SMBShareDs struct {
	ID           types.String          `tfsdk:"id"`
	FileSystemID types.String          `tfsdk:"file_system_id"`
	Name         types.String          `tfsdk:"name"`
	Filters      FilterExpressionValue `tfsdk:"filter_expression"`
	Items        []SMBShareDsData      `tfsdk:"smb_shares"`
}

type SMBShareDsData struct {
	ID                              types.String `tfsdk:"id"`
	FileSystemID                    types.String `tfsdk:"file_system_id"`
	Name                            types.String `tfsdk:"name"`
	Path                            types.String `tfsdk:"path"`
	Description                     types.String `tfsdk:"description"`
	IsContinuousAvailabilityEnabled types.Bool   `tfsdk:"is_continuous_availability_enabled"`
	IsEncryptionEnabled             types.Bool   `tfsdk:"is_encryption_enabled"`
	IsABEEnabled                    types.Bool   `tfsdk:"is_abe_enabled"`
	IsBranchCacheEnabled            types.Bool   `tfsdk:"is_branch_cache_enabled"`
	OfflineAvailability             types.String `tfsdk:"offline_availability"`
	OfflineAvailabilityLocalized    types.String `tfsdk:"offline_availability_l10n"`
	Umask                           types.String `tfsdk:"umask"`
}
