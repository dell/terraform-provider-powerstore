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

// newRemoteSystemDatasource returns new remote system datasource instance
func newRemoteSystemDatasource() datasource.DataSource {
	return &datasourceRemoteSystem{}
}

type datasourceRemoteSystem struct {
	client *client.Client
}

// Metadata defines datasource interface Metadata method
func (r *datasourceRemoteSystem) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_remote_system"
}

// Schema defines datasource interface Schema method
func (r *datasourceRemoteSystem) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "This datasource is used to query the existing Remote Systems from a PowerStore Array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		Description:         "This datasource is used to query the existing Remote Systems from a PowerStore Array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier of the Remote System to be fetched. Conflicts with `name`, `management_address` and `filter_expression`.",
				MarkdownDescription: "Unique identifier of the Remote System to be fetched. Conflicts with `name`, `management_address` and `filter_expression`.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ConflictsWith(
						path.MatchRoot("name"),
						path.MatchRoot("filter_expression"),
					),
				},
			},
			"name": schema.StringAttribute{
				Description:         "Name of the Remote System to be fetched. Conflicts with `id`, `management_address` and `filter_expression`.",
				MarkdownDescription: "Name of the Remote System to be fetched. Conflicts with `id`, `management_address` and `filter_expression`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ConflictsWith(
						path.MatchRoot("filter_expression"),
					),
				},
			},
			"filter_expression": schema.StringAttribute{
				Description:         "PowerStore filter expression to filter Remote Systems by. Conflicts with `id`, `name` and `management_address`.",
				MarkdownDescription: "PowerStore filter expression to filter Remote Systems by. Conflicts with `id`, `name` and `management_address`.",
				Optional:            true,
				CustomType:          models.FilterExpressionType{},
			},
			"remote_systems": schema.ListNestedAttribute{
				Description:         "List of Remote Systems fetched from PowerStore array.",
				MarkdownDescription: "List of Remote Systems fetched from PowerStore array.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: r.RemoteSystemDsSchema()},
			},
		},
	}
}

// RemoteSystemDsSchema defines datasource interface Schema method
func (r *datasourceRemoteSystem) RemoteSystemDsSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Unique identifier of the remote system instance.",
			Description:         "Unique identifier of the remote system instance.",
		},
		"name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Name of the remote system instance.",
			Description:         "Name of the remote system instance.",
		},
		"description": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Description of the remote system instance.",
			Description:         "Description of the remote system instance.",
		},
		"serial_number": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Serial number of the remote system instance.",
			Description:         "Serial number of the remote system instance.",
		},
		"type": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Type of the remote system instance.",
			Description:         "Type of the remote system instance.",
		},
		"management_address": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Management IP address of the remote system instance.",
			Description:         "Management IP address of the remote system instance.",
		},
		"data_connection_state": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Data connection state of the remote system.",
			Description:         "Data connection state of the remote system.",
		},
		"data_network_latency": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Data network latency of the remote system.",
			Description:         "Data network latency of the remote system.",
		},
		"capabilities": schema.ListAttribute{
			ElementType:         types.StringType,
			Computed:            true,
			MarkdownDescription: "List of supported remote protection capabilities.",
			Description:         "List of supported remote protection capabilities.",
		},
	}
}

// Configure - defines configuration for Remote System datasource
func (r *datasourceRemoteSystem) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

// Read - reads Remote System datasource information
func (r *datasourceRemoteSystem) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var conf models.RemoteSystemDs
	diags := req.Config.Get(ctx, &conf)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var remoteSystems []gopowerstore.RemoteSystem

	if !conf.ID.IsNull() && !conf.ID.IsUnknown() {
		remoteSystemID := conf.ID.ValueString()

		// Get remoteSystem details from API and then update what is in state from what the API returns
		remoteSystemResponse, err := r.client.PStoreClient.GetRemoteSystem(context.Background(), remoteSystemID)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error reading Remote System",
				"Could not read Remote System with id "+remoteSystemID+": "+err.Error(),
			)
			return
		}

		remoteSystems = append(remoteSystems, remoteSystemResponse)
		conf.Name = types.StringValue(remoteSystemResponse.Name)
	} else if !conf.Name.IsNull() && !conf.Name.IsUnknown() {
		remoteSystemResponse, err := r.client.PStoreClient.GetRemoteSystem(context.Background(), "name:"+conf.Name.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Error reading Remote System",
				"Could not read Remote System with name "+conf.Name.ValueString()+": "+err.Error(),
			)
			return
		}
		remoteSystems = append(remoteSystems, remoteSystemResponse)
		conf.ID = types.StringValue(remoteSystemResponse.ID)
	} else {
		filters := make(map[string]string)
		if !conf.Filters.IsNull() {
			filters = convertQueriesToMap(conf.Filters.ValueQueries())
		}

		remoteSystemResponse, err := r.client.PStoreClient.GetRemoteSystems(ctx, filters)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error reading Remote Systems",
				"Could not read Remote Systems with error "+err.Error(),
			)
			return
		}
		remoteSystems = append(remoteSystems, remoteSystemResponse...)
		conf.ID = types.StringValue("dummy id")
	}

	// Set state
	diags = resp.State.Set(ctx, r.getState(conf, remoteSystems))
	resp.Diagnostics.Append(diags...)
}

// getState - method to update terraform state
func (r *datasourceRemoteSystem) getState(cfg models.RemoteSystemDs, input []gopowerstore.RemoteSystem) *models.RemoteSystemDs {
	ret := models.RemoteSystemDs{
		ID:      cfg.ID,
		Name:    cfg.Name,
		Filters: cfg.Filters,
		Items:   make([]models.RemoteSystemDsItem, 0, len(input)),
	}
	for _, remoteSystem := range input {
		ret.Items = append(ret.Items, r.getItemState(remoteSystem))
	}
	return &ret
}

// RemoteSystemDsState - method to update terraform state
func (r *datasourceRemoteSystem) getItemState(input gopowerstore.RemoteSystem) models.RemoteSystemDsItem {
	return models.RemoteSystemDsItem{
		ID:                  input.ID,
		Name:                input.Name,
		Description:         input.Description,
		SerialNumber:        input.SerialNumber,
		Type:                input.Type,
		ManagementAddress:   input.ManagementAddress,
		DataConnectionState: input.DataConnectionState,
		DataNetworkLatency:  input.DataNetworkLatency,
		Capabilities:        input.Capabilities,
	}
}
