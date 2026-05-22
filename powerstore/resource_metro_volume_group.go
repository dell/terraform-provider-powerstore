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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// newMetroVolumeGroupResource returns new metro volume group resource instance
func newMetroVolumeGroupResource() resource.Resource {
	return &resourceMetroVolumeGroup{}
}

type resourceMetroVolumeGroup struct {
	client *clientgen.APIClient
}

// Metadata defines resource interface Metadata method
func (r *resourceMetroVolumeGroup) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metro_volume_group"
}

// Schema defines resource interface Schema method
func (r *resourceMetroVolumeGroup) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This resource is used to manage metro volume group replication on PowerStore. Metro replication provides synchronous replication (RPO=0) between two PowerStore clusters for disaster recovery. The resource takes an existing volume group ID and configures it as a metro volume group with a remote system. Running `terraform destroy` will end the metro configuration.",
		Description:         "This resource is used to manage metro volume group replication on PowerStore. Metro replication provides synchronous replication (RPO=0) between two PowerStore clusters for disaster recovery. The resource takes an existing volume group ID and configures it as a metro volume group with a remote system. Running terraform destroy will end the metro configuration.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "Unique identifier of the metro replication session.",
				MarkdownDescription: "Unique identifier of the metro replication session.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"volume_group_id": schema.StringAttribute{
				Required:            true,
				Description:         "Unique identifier of the existing volume group to configure as a metro volume group.",
				MarkdownDescription: "Unique identifier of the existing volume group to configure as a metro volume group.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"remote_system_id": schema.StringAttribute{
				Required:            true,
				Description:         "Unique identifier of the remote system for metro replication. The remote system must support metro volumes.",
				MarkdownDescription: "Unique identifier of the remote system for metro replication. The remote system must support metro volumes.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"remote_appliance_id": schema.StringAttribute{
				Optional:            true,
				Description:         "Specific remote appliance to assign the volume group. If not specified, the system chooses based on space, load, and connectivity.",
				MarkdownDescription: "Specific remote appliance to assign the volume group. If not specified, the system chooses based on space, load, and connectivity.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"delete_remote_volume_group": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				Description:         "Whether to delete the remote volume group when ending metro configuration. Default is false.",
				MarkdownDescription: "Whether to delete the remote volume group when ending metro configuration. Default is false.",
			},
			"force": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				Description:         "Force end metro even if the remote side has errors. Not recommended unless the remote is known to be down.",
				MarkdownDescription: "Force end metro even if the remote side has errors. Not recommended unless the remote is known to be down.",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				Description:         "Current state of the replication session.",
				MarkdownDescription: "Current state of the replication session.",
			},
			"remote_resource_id": schema.StringAttribute{
				Computed:            true,
				Description:         "Unique identifier of the remote storage resource.",
				MarkdownDescription: "Unique identifier of the remote storage resource.",
			},
			"data_transfer_state": schema.StringAttribute{
				Computed:            true,
				Description:         "Current data transfer state of the replication session.",
				MarkdownDescription: "Current data transfer state of the replication session.",
			},
			"metro_replication_session_id": schema.StringAttribute{
				Computed:            true,
				Description:         "Unique identifier of the metro replication session assigned to the volume group.",
				MarkdownDescription: "Unique identifier of the metro replication session assigned to the volume group.",
			},
		},
	}
}

// Configure - defines configuration for metro volume group resource
func (r *resourceMetroVolumeGroup) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = client.GenClient
}

// Create - configures metro replication on an existing volume group
func (r *resourceMetroVolumeGroup) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.MetroVolumeGroupResource
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Started Configure Metro Volume Group")

	volumeGroupID := plan.VolumeGroupID.ValueString()

	configureMetro := clientgen.VolumeGroupConfigureMetro{
		RemoteSystemId: plan.RemoteSystemID.ValueString(),
	}
	if !plan.RemoteApplianceID.IsNull() && !plan.RemoteApplianceID.IsUnknown() {
		remoteAppID := plan.RemoteApplianceID.ValueString()
		configureMetro.RemoteApplianceId = &remoteAppID
	}

	metroResponse, _, err := r.client.VolumeGroupApi.VolumeGroupConfigureMetro(ctx, volumeGroupID).Body(configureMetro).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error configuring metro volume group",
			"Could not configure metro on volume group "+volumeGroupID+": "+err.Error(),
		)
		return
	}

	sessionID := ""
	if metroResponse != nil && metroResponse.MetroReplicationSessionId != nil {
		sessionID = *metroResponse.MetroReplicationSessionId
	}

	if sessionID == "" {
		resp.Diagnostics.AddError(
			"Error configuring metro volume group",
			"Configure metro succeeded but no session ID was returned for volume group "+volumeGroupID,
		)
		return
	}

	plan.ID = types.StringValue(sessionID)
	plan.MetroReplicationSessionID = types.StringValue(sessionID)

	session, err := r.getReplicationSession(ctx, sessionID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading replication session after metro configuration",
			"Could not read replication session "+sessionID+": "+err.Error(),
		)
		return
	}

	r.updateState(&plan, session)

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	log.Printf("Successfully configured Metro Volume Group, session: %s", sessionID)
}

// Read - reads the metro volume group replication session state
func (r *resourceMetroVolumeGroup) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.MetroVolumeGroupResource
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	sessionID := state.ID.ValueString()
	session, err := r.getReplicationSession(ctx, sessionID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading metro volume group replication session",
			"Could not read replication session "+sessionID+": "+err.Error(),
		)
		return
	}

	r.updateState(&state, session)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Update - metro volume group does not support in-place update; mutable attrs trigger replacement
func (r *resourceMetroVolumeGroup) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan models.MetroVolumeGroupResource
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state models.MetroVolumeGroupResource
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.ID = state.ID
	plan.MetroReplicationSessionID = state.MetroReplicationSessionID

	sessionID := state.ID.ValueString()
	session, err := r.getReplicationSession(ctx, sessionID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading metro volume group replication session during update",
			"Could not read replication session "+sessionID+": "+err.Error(),
		)
		return
	}

	r.updateState(&plan, session)

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Delete - ends metro configuration on the volume group
func (r *resourceMetroVolumeGroup) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state models.MetroVolumeGroupResource
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Started End Metro Volume Group")

	volumeGroupID := state.VolumeGroupID.ValueString()

	deleteRemote := false
	if !state.DeleteRemoteVolumeGroup.IsNull() {
		deleteRemote = state.DeleteRemoteVolumeGroup.ValueBool()
	}
	forceEnd := false
	if !state.Force.IsNull() {
		forceEnd = state.Force.ValueBool()
	}

	endMetro := clientgen.VolumeGroupEndMetro{
		DeleteRemoteVolumeGroup: &deleteRemote,
		Force:                   &forceEnd,
	}

	_, err := r.client.VolumeGroupApi.VolumeGroupEndMetro(ctx, volumeGroupID).Body(endMetro).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error ending metro volume group configuration",
			"Could not end metro on volume group "+volumeGroupID+": "+err.Error(),
		)
		return
	}

	log.Printf("Successfully ended Metro Volume Group configuration for volume group: %s", volumeGroupID)
}

// ImportState - imports an existing metro volume group by replication session ID
func (r *resourceMetroVolumeGroup) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// getReplicationSession fetches replication session details by ID
func (r *resourceMetroVolumeGroup) getReplicationSession(ctx context.Context, sessionID string) (*clientgen.ReplicationSessionInstance, error) {
	sel := "*"
	queries := make(url.Values)
	queries.Set("select", sel)
	response, _, err := r.client.ReplicationSessionApi.GetReplicationSessionById(ctx, sessionID).Queries(queries).Execute()
	return response, err
}

// updateState updates the terraform state from the replication session response
func (r *resourceMetroVolumeGroup) updateState(state *models.MetroVolumeGroupResource, session *clientgen.ReplicationSessionInstance) {
	if session == nil {
		return
	}
	state.ID = helper.TfString(session.Id)
	state.MetroReplicationSessionID = helper.TfString(session.Id)
	if session.State != nil {
		state.State = types.StringValue(string(*session.State))
	}
	state.RemoteResourceID = helper.TfString(session.RemoteResourceId)
	if session.DataTransferState != nil {
		state.DataTransferState = types.StringValue(string(*session.DataTransferState))
	}

	if session.LocalResourceId != nil {
		state.VolumeGroupID = types.StringValue(*session.LocalResourceId)
	}
	if session.RemoteSystemId != nil {
		state.RemoteSystemID = types.StringValue(*session.RemoteSystemId)
	}
}
