/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	client "terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// newReplicationRuleResource returns new replication rule resource instance
func newReplicationRuleResource() resource.Resource {
	return &resourceReplicationRule{}
}

type resourceReplicationRule struct {
	client *client.Client
}

// Metadata defines resource interface Metadata method
func (r *resourceReplicationRule) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_replication_rule"
}

func (r *resourceReplicationRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This resource is used to manage the replication rule entity of PowerStore Array. We can Create, Update and Delete the replication rule using this resource. We can also import an existing replication rule from PowerStore array.",
		Description:         "This resource is used to manage the replication rule entity of PowerStore Array. We can Create, Update and Delete the replication rule using this resource. We can also import an existing replication rule from PowerStore array.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The ID of the replication rule.",
				MarkdownDescription: "The ID of the replication rule.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "The name of the replication rule.",
				MarkdownDescription: "The name of the replication rule.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"rpo": schema.StringAttribute{
				Required:            true,
				Description:         "Recovery Point Objective (RPO) of the replication rule.",
				MarkdownDescription: "Recovery Point Objective (RPO) of the replication rule.",
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						string(gopowerstore.RpoFiveMinutes),
						string(gopowerstore.RpoFifteenMinutes),
						string(gopowerstore.RpoThirtyMinutes),
						string(gopowerstore.RpoOneHour),
						string(gopowerstore.RpoSixHours),
						string(gopowerstore.RpoTwelveHours),
						string(gopowerstore.RpoOneDay),
						string(gopowerstore.RpoZero),
					}...),
				},
			},
			"remote_system_id": schema.StringAttribute{
				Required:            true,
				Description:         "Unique identifier of the remote system associated with the replication rule.",
				MarkdownDescription: "Unique identifier of the remote system associated with the replication rule.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(36),
				},
			},
			"alert_threshold": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Alert threshold for the replication rule.",
				MarkdownDescription: "Alert threshold for the replication rule.",
			},
			"is_read_only": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Indicates whether the replication rule is read-only.",
				MarkdownDescription: "Indicates whether the replication rule is read-only.",
			},
		},
	}
}

// Configure - defines configuration for replication rule resource
func (r *resourceReplicationRule) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
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

// Create creates the resource and sets the initial Terraform state.
func (r *resourceReplicationRule) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Creating replication rule state")
	var plan models.ReplicationRule

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// create payload
	replicationRuleCreate := &gopowerstore.ReplicationRuleCreate{
		Name:           plan.Name.ValueString(),
		Rpo:            gopowerstore.RPOEnum(plan.RPO.ValueString()),
		RemoteSystemID: plan.RemoteSystemID.ValueString(),
		AlertThreshold: int(plan.AlertThreshold.ValueInt64()),
		IsReadOnly:     plan.IsReadOnly.ValueBool(),
	}

	// Create new replication rule.
	createRes, err := r.client.PStoreClient.CreateReplicationRule(context.Background(), replicationRuleCreate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating replication rule",
			"Could not create replication rule, unexpected error: "+err.Error(),
		)
		return
	}

	// Get replicaton rule details using ID
	response, err := r.client.PStoreClient.GetReplicationRule(context.Background(), createRes.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting replication rule after creation",
			"Could not get replication rule, unexpected error: "+err.Error(),
		)
		return
	}

	state := models.ReplicationRule{}
	r.updateState(&state, response)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Done with creating replication rule resource state")
}

// Read refreshes the Terraform state with the latest value.
func (r *resourceReplicationRule) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Info(ctx, "Reading replication rule resource state")

	var state models.ReplicationRule
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get replicaton rule details using ID
	response, err := r.client.PStoreClient.GetReplicationRule(context.Background(), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting replication rule details",
			"Could not get replication rule, unexpected error: "+err.Error(),
		)
		return
	}

	// Set state
	r.updateState(&state, response)
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Done with reading replication rule resource state")
}

// Update updates the resource and sets the updated Terraform state.
func (r *resourceReplicationRule) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Info(ctx, "Updating replication rule resource state")

	// Get plan values
	var plan models.ReplicationRule
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state models.ReplicationRule
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if the attribute is_read_only is being modified
	if !plan.IsReadOnly.IsUnknown() && plan.IsReadOnly.ValueBool() != state.IsReadOnly.ValueBool() {
		resp.Diagnostics.AddError(
			"The attribute is_read_only cannot be modified.",
			"The attribute is_read_only cannot be modified.",
		)
		return
	}

	// create payload for update
	replicationRuleUpdate := &gopowerstore.ReplicationRuleModify{
		Name:           plan.Name.ValueString(),
		Rpo:            gopowerstore.RPOEnum(plan.RPO.ValueString()),
		RemoteSystemID: plan.RemoteSystemID.ValueString(),
		AlertThreshold: int(plan.AlertThreshold.ValueInt64()),
	}

	// Get replication rule ID from state
	replicationRuleID := state.ID.ValueString()

	// Update replication rule
	_, err := r.client.PStoreClient.ModifyReplicationRule(context.Background(), replicationRuleUpdate, replicationRuleID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating replication rule",
			"Could not update replication rule with id "+replicationRuleID+": "+err.Error(),
		)
		return
	}

	// Get replicaton rule details using ID
	response, err := r.client.PStoreClient.GetReplicationRule(context.Background(), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting replication rule after updation",
			"Could not get replication rule, unexpected error: "+err.Error(),
		)
		return
	}

	// Set state
	r.updateState(&state, response)
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Done with updating replication rule resource state")
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *resourceReplicationRule) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Info(ctx, "Deleting replication rule resource state")

	var state models.ReplicationRule
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get replication rule ID from state
	replicationRuleID := state.ID.ValueString()

	// Delete replication rule
	_, err := r.client.PStoreClient.DeleteReplicationRule(context.Background(), replicationRuleID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting replication rule",
			"Could not delete replication rule with ID "+replicationRuleID+": "+err.Error(),
		)
		return
	}

	resp.State.RemoveResource(ctx)
	tflog.Info(ctx, "Done with deleting replication rule resource state")
}

// ImportState import state for existing replication rule
func (r *resourceReplicationRule) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// updateState update state from the API response
func (r *resourceReplicationRule) updateState(state *models.ReplicationRule, response gopowerstore.ReplicationRule) {
	state.ID = types.StringValue(response.ID)
	state.Name = types.StringValue(response.Name)
	state.RPO = types.StringValue(string(response.Rpo))
	state.RemoteSystemID = types.StringValue(response.RemoteSystemID)
	state.IsReadOnly = types.BoolValue(response.IsReadOnly)
	state.AlertThreshold = types.Int64Value(int64(response.AlertThreshold))
}
