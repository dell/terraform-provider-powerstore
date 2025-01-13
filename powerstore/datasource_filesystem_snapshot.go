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
	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
	client *client.Client
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
					stringvalidator.ConflictsWith(path.MatchRoot("name"), path.MatchRoot("filesystem_id"), path.MatchRoot("nas_server_id")),
					stringvalidator.LengthAtLeast(1),
				},
			},
			"name": schema.StringAttribute{
				Description:         "File System Snapshot name. Conflicts with `id`.",
				MarkdownDescription: "File System Snapshot name. Conflicts with `id`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"filesystem_id": schema.StringAttribute{
				Description:         "Parent ID of the Snapshot. Conflicts with `id` and `nas_server_id`.",
				MarkdownDescription: "Parent ID of the Snapshot. Conflicts with `id` and `nas_server_id`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ConflictsWith(path.MatchRoot("nas_server_id")),
				},
			},

			"nas_server_id": schema.StringAttribute{
				Description:         "Nas Server ID of the Snapshot. Conflicts with `id` and `filesystem_id`.",
				MarkdownDescription: "Nas Server ID of the Snapshot. Conflicts with `id` and `filesystem_id`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
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
	d.client = req.ProviderData.(*client.Client)
}

// Read refreshes the Terraform state with the latest data
func (d *fileSystemSnapshotDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state models.FileSysteSnapshotDataSource
	var fileSystemSnapshots []gopowerstore.FileSystem
	var fileSystemSnapshot gopowerstore.FileSystem
	var err error

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	if state.ID.ValueString() != "" {
		fileSystemSnapshot, err = d.client.PStoreClient.GetFsSnapshot(context.Background(), state.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read PowerStore File System Snapshot by ID: "+state.ID.ValueString(),
				err.Error(),
			)
			return
		}
		fileSystemSnapshots = append(fileSystemSnapshots, fileSystemSnapshot)
	} else {
		filterMap := make(map[string]string)
		filterMap["filesystem_type"] = fmt.Sprintf("eq.%s", gopowerstore.FileSystemTypeEnumSnapshot)
		if state.Name.ValueString() != "" {
			filterMap["name"] = fmt.Sprintf("eq.%s", state.Name.ValueString())
		}
		if state.FileSystemID.ValueString() != "" {
			filterMap["parent_id"] = fmt.Sprintf("eq.%s", state.FileSystemID.ValueString())
		}
		if state.NasServerID.ValueString() != "" {
			filterMap["nas_server_id"] = fmt.Sprintf("eq.%s", state.NasServerID.ValueString())
		}
		tflog.Debug(ctx, fmt.Sprintf("PK Filter Map: %v", filterMap))
		fileSystemSnapshots, err = d.client.PStoreClient.GetFsByFilter(context.Background(), filterMap)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read PowerStore File System Snapshots",
				err.Error(),
			)
			return
		}
	}

	state.FileSystemSnapshots = updateFileSystemState(fileSystemSnapshots)
	state.ID = types.StringValue("placeholder")
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
