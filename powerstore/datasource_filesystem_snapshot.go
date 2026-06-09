/*
Copyright (c) 2025-2026 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	"net/url"
	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/models"
	"terraform-provider-powerstore/powerstore/helper"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &fileSystemSnapshotDataSource{}
	_ datasource.DataSourceWithConfigure = &fileSystemSnapshotDataSource{}
)

// newfileSystemSnapshotDataSource returns the fileSystemSnapshot snapshot data source object
func newFileSystemSnapshotDataSource() datasource.DataSource {
	return &fileSystemSnapshotDataSource{}
}

// fileSystemSnapshotDataSource is the data source implementation
type fileSystemSnapshotDataSource struct {
	client *clientgen.APIClient
}

// Metadata returns the data source type name
func (d *fileSystemSnapshotDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_filesystem_snapshot"
}

// Schema defines the schema for the data source
func (d *fileSystemSnapshotDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This datasource is used to query the existing File System Snapshot from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		MarkdownDescription: "This datasource is used to query the existing File System Snapshot from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier of the File System Snapshot. Conflicts with `name` and `filesystem_id`.",
				MarkdownDescription: "Unique identifier of the File System Snapshot. Conflicts with `name` and `filesystem_id`.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("name"), path.MatchRoot("filesystem_id"), path.MatchRoot("nas_server_id"), path.MatchRoot("filter_expression")),
					stringvalidator.LengthAtLeast(1),
				},
			},
			"name": schema.StringAttribute{
				Description:         "File System Snapshot name. Conflicts with `id`.",
				MarkdownDescription: "File System Snapshot name. Conflicts with `id`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ConflictsWith(path.MatchRoot("filter_expression")),
				},
			},

			"filesystem_id": schema.StringAttribute{
				Description:         "File System ID of the Snapshot. Conflicts with `id` and `nas_server_id`.",
				MarkdownDescription: "File System ID of the Snapshot. Conflicts with `id` and `nas_server_id`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ConflictsWith(path.MatchRoot("nas_server_id"), path.MatchRoot("filter_expression")),
				},
			},

			"nas_server_id": schema.StringAttribute{
				Description:         "Nas Server ID of the Snapshot. Conflicts with `id` and `filesystem_id`.",
				MarkdownDescription: "Nas Server ID of the Snapshot. Conflicts with `id` and `filesystem_id`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ConflictsWith(path.MatchRoot("filter_expression")),
				},
			},

			"filter_expression": schema.StringAttribute{
				Description:         "PowerStore filter expression to filter Filesystem Snapshots by. Conflicts with `id`, `name`, `nas_server_id` and `file_system_id`.",
				MarkdownDescription: "PowerStore filter expression to filter Filesystem Snapshots by. Conflicts with `id`, `name`, `nas_server_id` and `file_system_id`.",
				Optional:            true,
				CustomType:          models.FilterExpressionType{},
			},

			"filesystem_snapshots": schema.ListNestedAttribute{
				Description:         "List of File System Snapshots.",
				MarkdownDescription: "List of File System Snapshots.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: FileSystemDatasourceSchema()},
			},
		},
	}

}

// Configure adds the provider configured client to the data source
func (d *fileSystemSnapshotDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(*client.Client)
	d.client = client.GenClient
}

// Read refreshes the Terraform state with the latest data
func (d *fileSystemSnapshotDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state models.FileSysteSnapshotDataSource

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// sel := "*,flr_attributes(*),nas_server(*),protection_policy(*)"
	sel := "*"
	queries := make(url.Values)
	queries.Set("select", sel)
	queries.Set("filesystem_type", fmt.Sprintf("eq.%s", clientgen.FILESYSTEMTYPEENUM_SNAPSHOT))

	id := state.ID.ValueString()
	if !state.Name.IsNull() {
		queries.Set("name", "eq."+state.Name.ValueString())
	}
	if !state.FileSystemID.IsNull() {
		queries.Set("parent_id", "eq."+state.FileSystemID.ValueString())
	}
	if !state.NasServerID.IsNull() {
		queries.Set("nas_server_id", "eq."+state.NasServerID.ValueString())
	}
	if !state.Filters.IsNull() {
		err := validateFileSystemFilter(state.Filters.ValueString())
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root("filter_expression"),
				"Invalid filter expression",
				err.Error(),
			)
			return
		}
		queries = helper.MergeValues(queries, state.Filters.ValueQueries())
	}

	dsreq := helper.DsReq[clientgen.FileSystemInstance, clientgen.ApiGetFileSystemByIdRequest, clientgen.ApiGetAllFileSystemsRequest]{
		Instance:   d.client.FileSystemApi.GetFileSystemById,
		Collection: d.client.FileSystemApi.GetAllFileSystems,
	}

	fileSystems, err := dsreq.Execute(ctx, queries, id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read PowerStore File System Snapshots",
			err.Error(),
		)
		return
	}

	if state.Name.ValueString() != "" && len(fileSystems) == 0 {
		resp.Diagnostics.AddError(
			"Unable to Read PowerStore File System Snapshot",
			"There is no filesystem snapshot with name "+state.Name.ValueString(),
		)
		return
	}

	state.FileSystemSnapshots = updateFileSystemSnapshotState(fileSystems)
	state.ID = types.StringValue("placeholder")
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// updateFileSystemSnapshotState iterates over the filesystem snapshot list and update the state
func updateFileSystemSnapshotState(fileSystems []clientgen.FileSystemInstance) []models.FileSystemDatasource {
	return helper.SliceTransform(fileSystems, func(in clientgen.FileSystemInstance) models.FileSystemDatasource {
		return models.FileSystemDatasource{
			AccessPolicy:               helper.TfString(in.AccessPolicy),
			AccessType:                 helper.TfString(in.AccessType),
			ConfigType:                 helper.TfString(in.ConfigType),
			Description:                helper.TfString(in.Description),
			ExpirationTimestamp:        helper.TfStringFromPTime(in.ExpirationTimestamp),
			FilesystemType:             helper.TfString(in.FilesystemType),
			FolderRenamePolicy:         helper.TfString(in.FolderRenamePolicy),
			ID:                         helper.TfString(in.Id),
			IsAsyncMTimeEnabled:        helper.TfBool(in.IsAsyncMTimeEnabled),
			IsSmbNoNotifyEnabled:       helper.TfBool(in.IsSmbNoNotifyEnabled),
			IsSmbNotifyOnAccessEnabled: helper.TfBool(in.IsSmbNotifyOnAccessEnabled),
			IsSmbNotifyOnWriteEnabled:  helper.TfBool(in.IsSmbNotifyOnWriteEnabled),
			IsSmbOpLocksEnabled:        helper.TfBool(in.IsSmbOpLocksEnabled),
			IsSmbSyncWritesEnabled:     helper.TfBool(in.IsSmbSyncWritesEnabled),
			LockingPolicy:              helper.TfString(in.LockingPolicy),
			Name:                       helper.TfString(in.Name),
			NasServerID:                helper.TfString(in.NasServerId),
			ParentID:                   helper.TfString(in.ParentId),
			ProtectionPolicyID:         helper.TfString(in.ProtectionPolicyId),
			SizeTotal:                  helper.TfInt64(in.SizeTotal),
			SizeUsed:                   helper.TfInt64(in.SizeUsed),
			SmbNotifyOnChangeDirDepth:  helper.TfInt32AsInt64(in.SmbNotifyOnChangeDirDepth),
			IsQuotaEnabled:             helper.TfBool(in.IsQuotaEnabled),
			GracePeriod:                helper.TfInt32AsInt64(in.GracePeriod),
			DefaultHardLimit:           helper.TfInt64(in.DefaultHardLimit),
			DefaultSoftLimit:           helper.TfInt64(in.DefaultSoftLimit),
			CreationTimestamp:          helper.TfStringFromPTime(in.CreationTimestamp),
			LastRefreshTimestamp:       helper.TfStringFromPTime(in.LastRefreshTimestamp),
			LastWritableTimestamp:      helper.TfStringFromPTime(in.LastWritableTimestamp),
			IsModified:                 helper.TfBool(in.IsModified),
			CreatorType:                helper.TfString(in.CreatorType),
			FileEventsPublishingMode:   helper.TfString(in.FileEventsPublishingMode),
			HostIOSize:                 helper.TfString(in.HostIoSize),
			IsSecure:                   helper.TfBool(in.IsSecure),
			FlrAttributes: helper.TfObject(in.FlrAttributes, func(in clientgen.FlrInstance) models.FLRAttributesDatasource {
				return models.FLRAttributesDatasource{
					DefaultRetention:     helper.TfString(in.DefaultRetention),
					MaximumRetention:     helper.TfString(in.MaximumRetention),
					MinimumRetention:     helper.TfString(in.MinimumRetention),
					Mode:                 helper.TfString(in.Mode),
					AutoLock:             helper.TfBool(in.AutoLock),
					AutoDelete:           helper.TfBool(in.AutoDelete),
					PolicyInterval:       helper.TfInt32AsInt64(in.PolicyInterval),
					HasProtectedFiles:    helper.TfBool(in.HasProtectedFiles),
					ClockTime:            helper.TfString(in.ClockTime),
					MaximumRetentionDate: helper.TfString(in.MaximumRetentionDate),
				}
			}),
		}
	})
}
