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
)

var (
	_ datasource.DataSource              = &fileSystemDataSource{}
	_ datasource.DataSourceWithConfigure = &fileSystemDataSource{}
)

// newFileSystemDataSource returns the fileSystem data source object
func newFileSystemDataSource() datasource.DataSource {
	return &fileSystemDataSource{}
}

// fileSystemDataSource is the data source implementation
type fileSystemDataSource struct {
	client *client.Client
}

// Metadata returns the data source type name
func (d *fileSystemDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_filesystem"
}

// Schema defines the schema for the data source
func (d *fileSystemDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This datasource is used to query the existing File System from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		MarkdownDescription: "This datasource is used to query the existing File System from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier of the File System. Conflicts with `name` and `nas_server_id`.",
				MarkdownDescription: "Unique identifier of the File System. Conflicts with `name` and `nas_server_id`.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("name"), path.MatchRoot("nas_server_id")),
					stringvalidator.LengthAtLeast(1),
				},
			},
			"name": schema.StringAttribute{
				Description:         "File System name. Conflicts with `id` and `nas_server_id`.",
				MarkdownDescription: "File System name. Conflicts with `id` and `nas_server_id`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ConflictsWith(path.MatchRoot("id"), path.MatchRoot("nas_server_id")),
				},
			},
			"nas_server_id": schema.StringAttribute{
				Description:         "Nas server ID. Conflicts with `id` and `name`.",
				MarkdownDescription: "Nas server ID. Conflicts with `id` and `name`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ConflictsWith(path.MatchRoot("id"), path.MatchRoot("name")),
				},
			},
			"filesystems": schema.ListNestedAttribute{
				Description:         "List of File System.",
				MarkdownDescription: "List of File System.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: FileSystemDatasourceSchema()},
			},
		},
	}

}

// Configure adds the provider configured client to the data source
func (d *fileSystemDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = req.ProviderData.(*client.Client)
}

// Read refreshes the Terraform state with the latest data
func (d *fileSystemDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var plan models.FileSystemConfigDataSource
	var fileSystems []gopowerstore.FileSystem
	var fileSystem gopowerstore.FileSystem
	var err error

	diags := req.Config.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	filterMap := make(map[string]string)
	filterMap["filesystem_type"] = fmt.Sprintf("eq.%s", gopowerstore.FileSystemTypeEnumPrimary)

	//Read the filesystem based on id/name/nas server id and if nothing is mentioned, then it returns all the file system
	if plan.Name.ValueString() != "" {
		filterMap["name"] = fmt.Sprintf("eq.%s", plan.Name.ValueString())
	} else if plan.ID.ValueString() != "" {
		fileSystem, err = d.client.PStoreClient.GetFS(context.Background(), plan.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read PowerStore fileSystem",
				err.Error(),
			)
			return
		}
		fileSystems = append(fileSystems, fileSystem)
	} else if plan.NasServerID.ValueString() != "" {
		filterMap["nas_server_id"] = fmt.Sprintf("eq.%s", plan.NasServerID.ValueString())
	}

	if plan.ID.ValueString() == "" {
		fileSystems, err = d.client.PStoreClient.GetFsByFilter(context.Background(), filterMap)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read PowerStore File Systems",
				err.Error(),
			)
			return
		}
	}

	plan.FileSystems = updateFileSystemState(fileSystems)
	plan.ID = types.StringValue("placeholder")
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}