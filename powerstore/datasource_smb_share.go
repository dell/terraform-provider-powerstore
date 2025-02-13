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
	"context"
	"fmt"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"
)

// newSmbShareDatasource returns smbShare new datasource instance
func newSmbShareDatasource() datasource.DataSource {
	return &datasourcesmbShare{}
}

type datasourcesmbShare struct {
	client *client.Client
}

// Metadata defines datasource interface Metadata method
func (r *datasourcesmbShare) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_smb_share"
}

// Schema defines datasource interface Schema method
func (r *datasourcesmbShare) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "This datasource is used to query the existing SMB Shares from a PowerStore Array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		Description:         "This datasource is used to query the existing SMB Shares from a PowerStore Array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier of the SMB share to be fetched. Conflicts with `name`, `file_system_id` and `filter_expression`.",
				MarkdownDescription: "Unique identifier of the SMB share to be fetched. Conflicts with `name`, `file_system_id` and `filter_expression`.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ConflictsWith(
						path.MatchRoot("name"),
						path.MatchRoot("file_system_id"),
						path.MatchRoot("filter_expression"),
					),
				},
			},
			"name": schema.StringAttribute{
				Description:         "Name of the SMB share to be fetch. Conflicts with `id`, `file_system_id` and `filter_expression`.",
				MarkdownDescription: "Name of the SMB share to be fetch. Conflicts with `id`, `file_system_id` and `filter_expression`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ConflictsWith(
						path.MatchRoot("file_system_id"),
						path.MatchRoot("filter_expression"),
					),
				},
			},
			"file_system_id": schema.StringAttribute{
				Description:         "The ID of the file system whose SMB Shares are to be fetched. Conflicts with `id`, `name` and `filter_expression`.",
				MarkdownDescription: "The ID of the file system whose SMB Shares are to be fetched. Conflicts with `id`, `name` and `filter_expression`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ConflictsWith(
						path.MatchRoot("filter_expression"),
					),
				},
			},
			"filter_expression": schema.StringAttribute{
				Description:         "PowerStore filter expression to filter SMB shares by. Conflicts with `id`, `name` and `file_system_id`.",
				MarkdownDescription: "PowerStore filter expression to filter SMB shares by. Conflicts with `id`, `name` and `file_system_id`.",
				Optional:            true,
				CustomType:          models.FilterExpressionType{},
			},
			"smb_shares": schema.ListNestedAttribute{
				Description:         "List of SMB shares fetched from PowerStore array.",
				MarkdownDescription: "List of SMB shares fetched from PowerStore array.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: r.smbShareDsSchema()},
			},
		},
	}
}

// smbShareDsSchema defines datasource interface Schema method
func (r *datasourcesmbShare) smbShareDsSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			Description:         "The unique identifier of the SMB Share.",
			MarkdownDescription: "The unique identifier of the SMB Share.",
		},
		"file_system_id": schema.StringAttribute{
			MarkdownDescription: "The unique identifier of the file	system on which the SMB Share is created.",
			Description:         "The unique identifier of the file system on which the SMB Share is created.",
			Computed:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the SMB Share.",
			Description:         "The name of the SMB Share.",
			Computed:            true,
		},
		"path": schema.StringAttribute{
			MarkdownDescription: "The local path of the filesystem or any existing sub-folder of the file system exported via the SMB Share. This path is relative to the NAS Server.",
			Description:         "The local path of the filesystem or any existing sub-folder of the file system exported via the SMB Share. This path is relative to the NAS Server.",
			Computed:            true,
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "User defined SMB share description.",
			Description:         "User defined SMB share description.",
			Computed:            true,
		},
		"is_continuous_availability_enabled": schema.BoolAttribute{
			MarkdownDescription: "Whether continuous availability for Server Message Block (SMB) 3.0 is enabled for the SMB Share.",
			Description:         "Whether continuous availability for Server Message Block (SMB) 3.0 is enabled for the SMB Share.",
			Computed:            true,
		},
		"is_encryption_enabled": schema.BoolAttribute{
			MarkdownDescription: "Whether encryption for Server Message Block (SMB) 3.0 is enabled at the shared folder level.",
			Description:         "Whether encryption for Server Message Block (SMB) 3.0 is enabled at the shared folder level.",
			Computed:            true,
		},
		"is_abe_enabled": schema.BoolAttribute{
			MarkdownDescription: "Whether Access-based Enumeration (ABE) is enabled.",
			Description:         "Whether Access-based Enumeration (ABE) is enabled.",
			Computed:            true,
		},
		"is_branch_cache_enabled": schema.BoolAttribute{
			MarkdownDescription: "Whether BranchCache optimization is enabled." +
				" BranchCache optimization technology copies content from your main office or hosted cloud content servers" +
				" and caches the content at branch office locations, allowing client computers at branch offices to access" +
				" the content locally rather than over the WAN.",
			Description: "Whether BranchCache optimization is enabled." +
				" BranchCache optimization technology copies content from your main office or hosted cloud content servers" +
				" and caches the content at branch office locations, allowing client computers at branch offices to access" +
				" the content locally rather than over the WAN.",
			Computed: true,
		},
		"offline_availability": schema.StringAttribute{
			MarkdownDescription: "Defines valid states of Offline Availability, where the states are:" +
				" `Manual` - Only specified files will be available offline." +
				" `Documents` - All files that users open will be available offline." +
				" `Programs` - Program will preferably run from the offline cache even when connected to the network." +
				" All files that users open will be available offline." +
				" `None` - Prevents clients from storing documents and programs in offline cache.",
			Description: "Defines valid states of Offline Availability, where the states are:" +
				" `Manual` - Only specified files will be available offline." +
				" `Documents` - All files that users open will be available offline." +
				" `Programs` - Program will preferably run from the offline cache even when connected to the network." +
				" All files that users open will be available offline." +
				" `None` - Prevents clients from storing documents and programs in offline cache.",
			Computed: true,
		},
		"umask": schema.StringAttribute{
			MarkdownDescription: "The default UNIX umask for new files created on the Share.",
			Description:         "The default UNIX umask for new files created on the Share.",
			Computed:            true,
		},
		"offline_availability_l10n": schema.StringAttribute{
			MarkdownDescription: "Localized message string corresponding to offline_availability",
			Description:         "Localized message string corresponding to offline_availability",
			Computed:            true,
		},
	}
}

// Configure - defines configuration for smb share datasource
func (r *datasourcesmbShare) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// Read - reads smb share datasource information
func (r *datasourcesmbShare) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var conf models.SMBShareDs
	diags := req.Config.Get(ctx, &conf)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var smbShares []gopowerstore.SMBShare

	if !conf.ID.IsNull() && !conf.ID.IsUnknown() {
		smbShareID := conf.ID.ValueString()

		// Get smbShare details from API and then update what is in state from what the API returns
		smbShareResponse, err := r.client.PStoreClient.GetSMBShare(context.Background(), smbShareID)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error reading SMB Share",
				"Could not read SMB Share with id "+smbShareID+": "+err.Error(),
			)
			return
		}

		smbShares = append(smbShares, smbShareResponse)
		conf.Name = types.StringValue(smbShareResponse.Name)
	} else if !conf.Name.IsNull() && !conf.Name.IsUnknown() {
		smbShareResponse, err := r.client.PStoreClient.GetSMBShare(context.Background(), "name:"+conf.Name.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Error reading SMB Share",
				"Could not read SMB Share with name "+conf.Name.ValueString()+": "+err.Error(),
			)
			return
		}
		smbShares = append(smbShares, smbShareResponse)
		conf.ID = types.StringValue(smbShareResponse.ID)
	} else {
		filters := make(map[string]string)
		if !conf.Filters.IsNull() {
			filters = convertQueriesToMap(conf.Filters.ValueQueries())
		}
		if !conf.FileSystemID.IsNull() {
			filters["file_system_id"] = "eq." + conf.FileSystemID.ValueString()
		}

		smbShareResponse, err := r.client.PStoreClient.GetSMBShares(ctx, filters)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error reading SMB Shares",
				"Could not read SMB Shares with error "+err.Error(),
			)
			return
		}
		smbShares = append(smbShares, smbShareResponse...)
		conf.ID = types.StringValue("dummy id")
	}

	// Set state
	diags = resp.State.Set(ctx, r.getState(conf, smbShares))
	resp.Diagnostics.Append(diags...)
}

// getState - method to update terraform state
func (r *datasourcesmbShare) getState(cfg models.SMBShareDs, input []gopowerstore.SMBShare) *models.SMBShareDs {
	ret := models.SMBShareDs{
		ID:           cfg.ID,
		Name:         cfg.Name,
		FileSystemID: cfg.FileSystemID,
		Filters:      cfg.Filters,
		Items:        make([]models.SMBShareDsData, 0, len(input)),
	}
	for _, smbShare := range input {
		ret.Items = append(ret.Items, r.getItemState(smbShare))
	}
	return &ret
}

// smbShareDsState - method to update terraform state
func (r *datasourcesmbShare) getItemState(input gopowerstore.SMBShare) models.SMBShareDsData {
	return models.SMBShareDsData{
		ID:                              types.StringValue(input.ID),
		FileSystemID:                    types.StringValue(input.FileSystemID),
		Name:                            types.StringValue(input.Name),
		Description:                     types.StringValue(input.Description),
		Path:                            types.StringValue(input.Path),
		IsContinuousAvailabilityEnabled: types.BoolValue(input.IsContinuousAvailabilityEnabled),
		IsEncryptionEnabled:             types.BoolValue(input.IsEncryptionEnabled),
		IsABEEnabled:                    types.BoolValue(input.IsABEEnabled),
		IsBranchCacheEnabled:            types.BoolValue(input.IsBranchCacheEnabled),
		OfflineAvailability:             types.StringValue(input.OfflineAvailability),
		OfflineAvailabilityLocalized:    types.StringValue(input.OfflineAvailabilityL10N),
		Umask:                           types.StringValue(input.Umask),
	}
}
