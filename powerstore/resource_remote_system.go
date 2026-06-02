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
	"log"
	"net/url"

	client "terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/models"
	"terraform-provider-powerstore/powerstore/helper"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// newRemoteSystemResource returns new remote system resource instance
func newRemoteSystemResource() resource.Resource {
	return &resourceRemoteSystem{}
}

type resourceRemoteSystem struct {
	client *client.Client
}

// Metadata defines resource interface Metadata method
func (r *resourceRemoteSystem) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_remote_system"
}

// Schema defines resource interface Schema method
func (r *resourceRemoteSystem) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This resource is used to manage remote system relationships on PowerStore. A remote system represents a connection to another storage system for replication or data import. For PowerStore-to-PowerStore connections, only `management_address` and `data_network_latency` are required. Non-PowerStore remote systems require additional parameters such as `remote_username`, `remote_password`, and `type`. Running `terraform destroy` will delete the remote system relationship.",
		Description:         "This resource is used to manage remote system relationships on PowerStore. A remote system represents a connection to another storage system for replication or data import. For PowerStore-to-PowerStore connections, only management_address and data_network_latency are required. Non-PowerStore remote systems require additional parameters such as remote_username, remote_password, and type. Running terraform destroy will delete the remote system relationship.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "Unique identifier of the remote system.",
				MarkdownDescription: "Unique identifier of the remote system.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"management_address": schema.StringAttribute{
				Required:            true,
				Description:         "Management address of the remote system instance. IPv4, IPv6, and FQDN are supported for PowerStore remote systems.",
				MarkdownDescription: "Management address of the remote system instance. IPv4, IPv6, and FQDN are supported for PowerStore remote systems.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "User-specified name of the remote system. Used only for non-PowerStore systems.",
				MarkdownDescription: "User-specified name of the remote system. Used only for non-PowerStore systems.",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "User-specified description of the remote system.",
				MarkdownDescription: "User-specified description of the remote system.",
			},
			"type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Type of the remote system. For PowerStore-to-PowerStore, this is auto-detected.",
				MarkdownDescription: "Type of the remote system. For PowerStore-to-PowerStore, this is auto-detected.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"data_network_latency": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Network latency for the remote system. Valid values: Low, High.",
				MarkdownDescription: "Network latency for the remote system. Valid values: `Low`, `High`.",
				Validators: []validator.String{
					stringvalidator.OneOf("Low", "High"),
				},
			},
			"remote_username": schema.StringAttribute{
				Optional:            true,
				Description:         "Username used to access the remote system. Used only for PowerProtect DD and non-PowerStore systems.",
				MarkdownDescription: "Username used to access the remote system. Used only for PowerProtect DD and non-PowerStore systems.",
			},
			"remote_password": schema.StringAttribute{
				Optional:            true,
				Sensitive:           true,
				Description:         "Password used to access the remote system. Used only for PowerProtect DD and non-PowerStore systems.",
				MarkdownDescription: "Password used to access the remote system. Used only for PowerProtect DD and non-PowerStore systems.",
			},
			"data_connection_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Data connection type. Valid values: iSCSI, FC.",
				MarkdownDescription: "Data connection type. Valid values: `iSCSI`, `FC`.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"iscsi_addresses": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "iSCSI target IP addresses for the data connection. Required for non-PowerStore remote systems with iSCSI.",
				MarkdownDescription: "iSCSI target IP addresses for the data connection. Required for non-PowerStore remote systems with iSCSI.",
				PlanModifiers:       []planmodifier.List{},
			},
			"serial_number": schema.StringAttribute{
				Computed:            true,
				Description:         "Serial number of the remote system.",
				MarkdownDescription: "Serial number of the remote system.",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				Description:         "Current state of the remote system.",
				MarkdownDescription: "Current state of the remote system.",
			},
			"data_connection_state": schema.StringAttribute{
				Computed:            true,
				Description:         "Data connection state of the remote system.",
				MarkdownDescription: "Data connection state of the remote system.",
			},
			"version": schema.StringAttribute{
				Computed:            true,
				Description:         "Version of the remote system.",
				MarkdownDescription: "Version of the remote system.",
			},
			"capabilities": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				Description:         "List of supported remote protection capabilities.",
				MarkdownDescription: "List of supported remote protection capabilities.",
			},
			"exchange_username": schema.StringAttribute{
				Optional:            true,
				Description:         "Username for certificate exchange with remote PowerStore system. Required for PowerStore-to-PowerStore connections. Can be the admin username or a temporary client_id from generate_temp_credentials API.",
				MarkdownDescription: "Username for certificate exchange with remote PowerStore system. Required for PowerStore-to-PowerStore connections. Can be the admin username or a temporary `client_id` from `generate_temp_credentials` API.",
			},
			"exchange_password": schema.StringAttribute{
				Optional:            true,
				Sensitive:           true,
				Description:         "Password for certificate exchange with remote PowerStore system. Required for PowerStore-to-PowerStore connections. Can be the admin password or a temporary secret from generate_temp_credentials API.",
				MarkdownDescription: "Password for certificate exchange with remote PowerStore system. Required for PowerStore-to-PowerStore connections. Can be the admin password or a temporary `secret` from `generate_temp_credentials` API.",
			},
		},
	}
}

// Configure - defines configuration for remote system resource
func (r *resourceRemoteSystem) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	r.client = client
}

// Create - creates a remote system relationship
func (r *resourceRemoteSystem) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.RemoteSystemResource
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Started Create Remote System")

	// Step 1: Exchange certificates if credentials are provided (for PowerStore-to-PowerStore)
	if !plan.ExchangeUsername.IsNull() && !plan.ExchangeUsername.IsUnknown() &&
		!plan.ExchangePassword.IsNull() && !plan.ExchangePassword.IsUnknown() {
		log.Printf("Exchanging certificates with remote system")
		exchangeBody := clientgen.X509CertificateExchange{
			Service:  clientgen.X509CertificateServiceEnum("Replication_HTTP"),
			Address:  plan.ManagementAddress.ValueString(),
			Port:     443,
			Username: plan.ExchangeUsername.ValueString(),
			Password: plan.ExchangePassword.ValueString(),
		}
		_, err := r.client.GenClient.X509CertificateApi.PostX509CertificateById(ctx).Body(exchangeBody).Execute()
		if err != nil {
			resp.Diagnostics.AddError(
				"Error exchanging certificates",
				"Could not exchange certificates with remote system: "+err.Error(),
			)
			return
		}
		log.Printf("Certificate exchange successful")
	}

	// Step 2: Create remote system
	createBody := clientgen.RemoteSystemCreate{
		ManagementAddress: helper.ValueToPointer[string](plan.ManagementAddress),
	}

	if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		createBody.Name = helper.ValueToPointer[string](plan.Name)
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		createBody.Description = helper.ValueToPointer[string](plan.Description)
	}
	if !plan.Type.IsNull() && !plan.Type.IsUnknown() {
		t := clientgen.RemoteSystemTypeEnum(plan.Type.ValueString())
		createBody.Type = &t
	}
	if !plan.DataNetworkLatency.IsNull() && !plan.DataNetworkLatency.IsUnknown() {
		l := clientgen.RemoteSystemLatencyEnum(plan.DataNetworkLatency.ValueString())
		createBody.DataNetworkLatency = &l
	}
	if !plan.RemoteUsername.IsNull() && !plan.RemoteUsername.IsUnknown() {
		createBody.RemoteUsername = helper.ValueToPointer[string](plan.RemoteUsername)
	}
	if !plan.RemotePassword.IsNull() && !plan.RemotePassword.IsUnknown() {
		createBody.RemotePassword = helper.ValueToPointer[string](plan.RemotePassword)
	}
	if !plan.DataConnectionType.IsNull() && !plan.DataConnectionType.IsUnknown() {
		dct := clientgen.DataConnectionTypeEnum(plan.DataConnectionType.ValueString())
		createBody.DataConnectionType = &dct
	}
	if !plan.IscsiAddresses.IsNull() && !plan.IscsiAddresses.IsUnknown() {
		var addrs []string
		diags = plan.IscsiAddresses.ElementsAs(ctx, &addrs, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		createBody.IscsiAddresses = addrs
	}

	createRes, _, err := r.client.GenClient.RemoteSystemApi.PostAllRemoteSystems(ctx).Body(createBody).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating remote system",
			"Could not create remote system: "+err.Error(),
		)
		return
	}

	rsID := ""
	if createRes != nil && createRes.Id != nil {
		rsID = *createRes.Id
	}

	if rsID == "" {
		resp.Diagnostics.AddError(
			"Error creating remote system",
			"Create succeeded but no ID was returned",
		)
		return
	}

	plan.ID = types.StringValue(rsID)

	// Read back to populate computed fields
	rsInstance, err := r.getRemoteSystem(ctx, rsID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading remote system after creation",
			"Could not read remote system "+rsID+": "+err.Error(),
		)
		return
	}

	r.updateState(ctx, &plan, rsInstance)

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	log.Printf("Successfully created Remote System: %s", rsID)
}

// Read - reads the remote system state
func (r *resourceRemoteSystem) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.RemoteSystemResource
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	rsID := state.ID.ValueString()
	rsInstance, err := r.getRemoteSystem(ctx, rsID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading remote system",
			"Could not read remote system "+rsID+": "+err.Error(),
		)
		return
	}

	r.updateState(ctx, &state, rsInstance)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Update - updates the remote system
func (r *resourceRemoteSystem) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan models.RemoteSystemResource
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state models.RemoteSystemResource
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	rsID := state.ID.ValueString()
	log.Printf("Started Update Remote System: %s", rsID)

	modifyBody := clientgen.RemoteSystemModify{}
	changed := false

	if !plan.Description.IsNull() && plan.Description.ValueString() != state.Description.ValueString() {
		modifyBody.Description = helper.ValueToPointer[string](plan.Description)
		changed = true
	}
	if !plan.ManagementAddress.IsNull() && plan.ManagementAddress.ValueString() != state.ManagementAddress.ValueString() {
		modifyBody.ManagementAddress = helper.ValueToPointer[string](plan.ManagementAddress)
		changed = true
	}
	if !plan.Name.IsNull() && !plan.Name.IsUnknown() && plan.Name.ValueString() != state.Name.ValueString() {
		modifyBody.Name = helper.ValueToPointer[string](plan.Name)
		changed = true
	}
	if !plan.DataNetworkLatency.IsNull() && !plan.DataNetworkLatency.IsUnknown() && plan.DataNetworkLatency.ValueString() != state.DataNetworkLatency.ValueString() {
		l := clientgen.RemoteSystemLatencyEnum(plan.DataNetworkLatency.ValueString())
		modifyBody.DataNetworkLatency = &l
		changed = true
	}
	if !plan.RemoteUsername.IsNull() && !plan.RemoteUsername.IsUnknown() {
		modifyBody.RemoteUsername = helper.ValueToPointer[string](plan.RemoteUsername)
		changed = true
	}
	if !plan.RemotePassword.IsNull() && !plan.RemotePassword.IsUnknown() {
		modifyBody.RemotePassword = helper.ValueToPointer[string](plan.RemotePassword)
		changed = true
	}

	if changed {
		_, err := r.client.GenClient.RemoteSystemApi.PatchRemoteSystemById(ctx, rsID).Body(modifyBody).Execute()
		if err != nil {
			resp.Diagnostics.AddError(
				"Error updating remote system",
				fmt.Sprintf("Could not update remote system %s: %s", rsID, err.Error()),
			)
			return
		}
	}

	// Read back to refresh state
	rsInstance, err := r.getRemoteSystem(ctx, rsID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading remote system after update",
			"Could not read remote system "+rsID+": "+err.Error(),
		)
		return
	}

	r.updateState(ctx, &plan, rsInstance)

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	log.Printf("Successfully updated Remote System: %s", rsID)
}

// Delete - deletes the remote system
func (r *resourceRemoteSystem) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state models.RemoteSystemResource
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	rsID := state.ID.ValueString()
	log.Printf("Started Delete Remote System: %s", rsID)

	_, err := r.client.GenClient.RemoteSystemApi.DeleteRemoteSystemById(ctx, rsID).Body(map[string]interface{}{}).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting remote system",
			"Could not delete remote system "+rsID+": "+err.Error(),
		)
		return
	}

	log.Printf("Successfully deleted Remote System: %s", rsID)
}

// ImportState - imports an existing remote system by ID
func (r *resourceRemoteSystem) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// getRemoteSystem fetches remote system details by ID
func (r *resourceRemoteSystem) getRemoteSystem(ctx context.Context, id string) (*clientgen.RemoteSystemInstance, error) {
	sel := "*"
	queries := make(url.Values)
	queries.Set("select", sel)
	response, _, err := r.client.GenClient.RemoteSystemApi.GetRemoteSystemById(ctx, id).Queries(queries).Execute()
	return response, err
}

// updateState updates the terraform state from the remote system API response
func (r *resourceRemoteSystem) updateState(ctx context.Context, state *models.RemoteSystemResource, rs *clientgen.RemoteSystemInstance) {
	if rs == nil {
		return
	}

	state.ID = helper.TfString(rs.Id)
	state.ManagementAddress = helper.TfString(rs.ManagementAddress)
	state.Name = helper.TfString(rs.Name)
	state.Description = helper.TfString(rs.Description)
	state.SerialNumber = helper.TfString(rs.SerialNumber)
	state.Version = helper.TfString(rs.Version)

	if rs.Type != nil {
		state.Type = types.StringValue(string(*rs.Type))
	}
	if rs.State != nil {
		state.State = types.StringValue(string(*rs.State))
	}
	if rs.DataConnectionState != nil {
		state.DataConnectionState = types.StringValue(string(*rs.DataConnectionState))
	}
	if rs.DataNetworkLatency != nil {
		state.DataNetworkLatency = types.StringValue(string(*rs.DataNetworkLatency))
	}
	if rs.DataConnectionType != nil {
		state.DataConnectionType = types.StringValue(string(*rs.DataConnectionType))
	}

	// Map capabilities
	if rs.Capabilities != nil {
		capVals := make([]attr.Value, len(rs.Capabilities))
		for i, c := range rs.Capabilities {
			capVals[i] = types.StringValue(string(c))
		}
		state.Capabilities, _ = types.ListValue(types.StringType, capVals)
	} else {
		state.Capabilities = types.ListNull(types.StringType)
	}
}
