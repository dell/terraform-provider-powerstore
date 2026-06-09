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
	"net/url"
	"sort"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	client "terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/models"
	"terraform-provider-powerstore/powerstore/helper"
)

// newRemoteSystemDatasource returns new remote system datasource instance
func newRemoteSystemDatasource() datasource.DataSource {
	return &datasourceRemoteSystem{}
}

type datasourceRemoteSystem struct {
	client *clientgen.APIClient
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
				Description:         "Unique identifier of the Remote System to be fetched. Conflicts with `name` and `filter_expression`.",
				MarkdownDescription: "Unique identifier of the Remote System to be fetched. Conflicts with `name` and `filter_expression`.",
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
				Description:         "Name of the Remote System to be fetched. Conflicts with `id` and `filter_expression`.",
				MarkdownDescription: "Name of the Remote System to be fetched. Conflicts with `id` and `filter_expression`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ConflictsWith(
						path.MatchRoot("filter_expression"),
					),
				},
			},
			"filter_expression": schema.StringAttribute{
				Description:         "PowerStore filter expression to filter Remote Systems by. Conflicts with `id` and `name`.",
				MarkdownDescription: "PowerStore filter expression to filter Remote Systems by. Conflicts with `id` and `name`.",
				Optional:            true,
				CustomType:          models.FilterExpressionType{},
			},
			"remote_systems": schema.ListNestedAttribute{
				Description:         "List of Remote Systems fetched from PowerStore array.",
				MarkdownDescription: "List of Remote Systems fetched from PowerStore array.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: remoteSystemDsSchema(),
				},
			},
		},
	}
}

// remoteSystemDsSchema defines the nested schema for individual remote system items
func remoteSystemDsSchema() map[string]schema.Attribute {
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
		"data_connection_type": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Data connection type of the remote system. Values: iSCSI, FC, TCP, DD_Boost.",
			Description:         "Data connection type of the remote system. Values: iSCSI, FC, TCP, DD_Boost.",
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
		"state": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Current state of the remote system.",
			Description:         "Current state of the remote system.",
		},
		"version": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Version of the remote system.",
			Description:         "Version of the remote system.",
		},
		"fc_target_wwns": schema.ListAttribute{
			ElementType:         types.StringType,
			Computed:            true,
			MarkdownDescription: "FC target World Wide Names for the data connection.",
			Description:         "FC target World Wide Names for the data connection.",
		},
		"iscsi_addresses": schema.ListAttribute{
			ElementType:         types.StringType,
			Computed:            true,
			MarkdownDescription: "iSCSI target IP addresses for the data connection.",
			Description:         "iSCSI target IP addresses for the data connection.",
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
	if req.ProviderData == nil {
		return
	}

	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Datasource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = c.GenClient
}

// Read - reads Remote System datasource information
func (r *datasourceRemoteSystem) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state models.RemoteSystemDs
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	queries := make(url.Values)
	id := state.ID.ValueString()

	sel := "*"
	queries.Set("select", sel)

	if !state.Name.IsNull() {
		queries.Set("name", "eq."+state.Name.ValueString())
	} else if !state.Filters.IsNull() {
		queries = helper.MergeValues(queries, state.Filters.ValueQueries())
	}

	dsreq := helper.DsReq[clientgen.RemoteSystemInstance, clientgen.ApiGetRemoteSystemByIdRequest, clientgen.ApiGetAllRemoteSystemsRequest]{
		Instance:   r.client.RemoteSystemApi.GetRemoteSystemById,
		Collection: r.client.RemoteSystemApi.GetAllRemoteSystems,
	}

	remoteSystems, err := dsreq.Execute(ctx, queries, id)
	if err != nil {
		errMsg := helper.ExtractErrorMessage(err)
		if id != "" {
			resp.Diagnostics.AddError(
				"Error reading Remote System",
				"Could not read Remote System with id "+id+": "+errMsg,
			)
		} else if !state.Name.IsNull() {
			resp.Diagnostics.AddError(
				"Error reading Remote System",
				"Could not read Remote System with name "+state.Name.ValueString()+": "+errMsg,
			)
		} else {
			resp.Diagnostics.AddError(
				"Error reading Remote Systems",
				"Could not read Remote Systems with error "+errMsg,
			)
		}
		return
	}

	// check that there is at least one remote system if name or id is provided
	if id != "" && len(remoteSystems) == 0 {
		resp.Diagnostics.AddError(
			"Error reading Remote System",
			"Could not read Remote System with id "+id+": no results found",
		)
		return
	}
	if !state.Name.IsNull() && len(remoteSystems) == 0 {
		resp.Diagnostics.AddError(
			"Error reading Remote System",
			"Could not read Remote System with name "+state.Name.ValueString()+": no results found",
		)
		return
	}

	state.Items = updateRemoteSystemDsState(remoteSystems)
	state.ID = types.StringValue("placeholder")
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

// updateRemoteSystemDsState converts API responses to terraform state
func updateRemoteSystemDsState(remoteSystems []clientgen.RemoteSystemInstance) []models.RemoteSystemDsItem {
	return helper.SliceTransform(remoteSystems, func(in clientgen.RemoteSystemInstance) models.RemoteSystemDsItem {
		return models.RemoteSystemDsItem{
			ID:                  helper.TfString(in.Id),
			Name:                helper.TfString(in.Name),
			Description:         helper.TfString(in.Description),
			SerialNumber:        helper.TfString(in.SerialNumber),
			Type:                helper.TfString(in.Type),
			ManagementAddress:   helper.TfString(in.ManagementAddress),
			DataConnectionType:  helper.TfString(in.DataConnectionType),
			DataConnectionState: helper.TfString(in.DataConnectionState),
			DataNetworkLatency:  helper.TfString(in.DataNetworkLatency),
			State:               helper.TfString(in.State),
			Version:             helper.TfString(in.Version),
			FcTargetWwns:        stringSliceToTfList(in.FcTargetWwns),
			IscsiAddresses:      stringSliceToTfList(in.IscsiAddresses),
			Capabilities:        capabilitiesToSortedTfList(in.Capabilities),
		}
	})
}

// stringSliceToTfList converts a Go string slice to a types.List of StringType
func stringSliceToTfList(in []string) types.List {
	if in == nil {
		return types.ListNull(types.StringType)
	}
	vals := make([]attr.Value, len(in))
	for i, v := range in {
		vals[i] = types.StringValue(v)
	}
	list, _ := types.ListValue(types.StringType, vals)
	return list
}

// capabilitiesToSortedTfList converts capabilities enum slice to a sorted types.List
func capabilitiesToSortedTfList(in []clientgen.RemoteProtectionCapabilityEnum) types.List {
	if in == nil {
		return types.ListNull(types.StringType)
	}
	strs := make([]string, len(in))
	for i, c := range in {
		strs[i] = string(c)
	}
	sort.Strings(strs)
	vals := make([]attr.Value, len(strs))
	for i, s := range strs {
		vals[i] = types.StringValue(s)
	}
	list, _ := types.ListValue(types.StringType, vals)
	return list
}
