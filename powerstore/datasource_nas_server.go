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
	_ datasource.DataSource              = &nasServerDataSource{}
	_ datasource.DataSourceWithConfigure = &nasServerDataSource{}
)

// newNasServerDataSource returns the nasServer data source object
func newNasServerDataSource() datasource.DataSource {
	return &nasServerDataSource{}
}

// nasServerDataSource is the data source implementation
type nasServerDataSource struct {
	client *client.Client
}

// Metadata returns the data source type name
func (d *nasServerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nas_server"
}

// Schema defines the schema for the data source
func (d *nasServerDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This datasource is used to query the existing NAS Server from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		MarkdownDescription: "This datasource is used to query the existing NAS Server from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier of the NAS Server. Conflicts with `name`.",
				MarkdownDescription: "Unique identifier of the NAS Server. Conflicts with `name`.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ConflictsWith(path.MatchRoot("name")),
				},
			},
			"name": schema.StringAttribute{
				Description:         "NAS Server name. Conflicts with `id`.",
				MarkdownDescription: "NAS Server name. Conflicts with `id`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ConflictsWith(path.MatchRoot("id")),
				},
			},

			"nas_servers": schema.ListNestedAttribute{
				Description:         "List of NAS Servers.",
				MarkdownDescription: "List of NAS Servers.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: NasServerDatasourceSchema()},
			},
		},
	}
}

// Configure adds the provider configured client to the data source
func (d *nasServerDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = req.ProviderData.(*client.Client)
}

// Read updates the Terraform state with the latest NAS Server data
func (d *nasServerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var plan models.NasServerConfigDataSource
	var nasServers []gopowerstore.NAS
	var nasServer gopowerstore.NAS
	var err error

	diags := req.Config.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Fetch NAS Server based on id/name; if nothing is mentioned, it fetches all NAS Servers
	if plan.Id.ValueString() != "" {
		nasServer, err = d.client.PStoreClient.GetNAS(context.Background(), plan.Id.ValueString())
		nasServers = append(nasServers, nasServer)
	} else if plan.Name.ValueString() != "" {
		nasServer, err = d.client.PStoreClient.GetNASByName(context.Background(), plan.Name.ValueString())
		nasServers = append(nasServers, nasServer)
	} else {
		nasServers, err = d.client.PStoreClient.GetNASServers(ctx)
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read PowerStore NAS Servers.",
			err.Error(),
		)
		return
	}

	plan.NasServers = updateNasServerState(nasServers)
	plan.Id = types.StringValue("placeholder")
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
