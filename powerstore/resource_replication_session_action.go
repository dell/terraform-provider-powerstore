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
	"time"

	client "terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/models"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// newReplicationSessionActionResource returns new replication session action resource instance
func newReplicationSessionActionResource() resource.Resource {
	return &resourceReplicationSessionAction{}
}

type resourceReplicationSessionAction struct {
	client *clientgen.APIClient
}

// Metadata defines resource interface Metadata method
func (r *resourceReplicationSessionAction) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_replication_session_action"
}

// Schema defines resource interface Schema method
func (r *resourceReplicationSessionAction) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This resource is used to perform actions on an existing PowerStore replication session. Supported actions include: `sync`, `pause`, `resume`, `failover`, `reprotect`, `start_failover_test`, `stop_failover_test`. Running `terraform destroy` only removes the resource from state; it does not undo the action.",
		Description:         "This resource is used to perform actions on an existing PowerStore replication session. Supported actions include: sync, pause, resume, failover, reprotect, start_failover_test, stop_failover_test. Running terraform destroy only removes the resource from state; it does not undo the action.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "Composite identifier of the action in the format session_id/action/timestamp.",
				MarkdownDescription: "Composite identifier of the action in the format `session_id/action/timestamp`.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"session_id": schema.StringAttribute{
				Required:            true,
				Description:         "Unique identifier of the replication session to perform the action on.",
				MarkdownDescription: "Unique identifier of the replication session to perform the action on.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"action": schema.StringAttribute{
				Required:            true,
				Description:         "Action to perform on the replication session. Valid values: sync, pause, resume, failover, reprotect, start_failover_test, stop_failover_test.",
				MarkdownDescription: "Action to perform on the replication session. Valid values: `sync`, `pause`, `resume`, `failover`, `reprotect`, `start_failover_test`, `stop_failover_test`.",
				Validators: []validator.String{
					stringvalidator.OneOf("sync", "pause", "resume", "failover", "reprotect", "start_failover_test", "stop_failover_test"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"is_planned": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
				Description:         "For failover action: whether the failover is planned (true) or unplanned (false). Default is true.",
				MarkdownDescription: "For failover action: whether the failover is planned (true) or unplanned (false). Default is true.",
			},
			"reverse": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				Description:         "For failover action: whether to auto-reprotect after failover. Default is false.",
				MarkdownDescription: "For failover action: whether to auto-reprotect after failover. Default is false.",
			},
			"post_state": schema.StringAttribute{
				Computed:            true,
				Description:         "State of the replication session after the action was performed.",
				MarkdownDescription: "State of the replication session after the action was performed.",
			},
		},
	}
}

// Configure - defines configuration for replication session action resource
func (r *resourceReplicationSessionAction) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// Create - executes the specified action on the replication session
func (r *resourceReplicationSessionAction) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.ReplicationSessionActionResource
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	sessionID := plan.SessionID.ValueString()
	action := plan.Action.ValueString()

	log.Printf("Started Replication Session Action: %s on session %s", action, sessionID)

	err := r.executeAction(ctx, sessionID, action, &plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error performing replication session action",
			fmt.Sprintf("Could not perform action '%s' on session %s: %s", action, sessionID, err.Error()),
		)
		return
	}

	// Set composite ID
	compositeID := fmt.Sprintf("%s/%s/%s", sessionID, action, time.Now().UTC().Format(time.RFC3339))
	plan.ID = types.StringValue(compositeID)

	// Read post-action state
	session, err := r.getReplicationSession(ctx, sessionID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading replication session after action",
			"Could not read replication session "+sessionID+": "+err.Error(),
		)
		return
	}

	if session != nil && session.State != nil {
		plan.PostState = types.StringValue(string(*session.State))
	}

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	log.Printf("Successfully completed Replication Session Action: %s on session %s", action, sessionID)
}

// Read - reads the replication session state after action
func (r *resourceReplicationSessionAction) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.ReplicationSessionActionResource
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	sessionID := state.SessionID.ValueString()
	session, err := r.getReplicationSession(ctx, sessionID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading replication session",
			"Could not read replication session "+sessionID+": "+err.Error(),
		)
		return
	}

	if session != nil && session.State != nil {
		state.PostState = types.StringValue(string(*session.State))
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Update - not supported; all attributes require replacement
func (r *resourceReplicationSessionAction) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"Update not supported",
		"Replication session action resources do not support in-place updates. All changes require resource replacement.",
	)
}

// Delete - no-op; state removal only, does not undo the action
func (r *resourceReplicationSessionAction) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("Removing replication session action from state (no-op)")
}

// executeAction performs the specified action on the replication session
func (r *resourceReplicationSessionAction) executeAction(ctx context.Context, sessionID, action string, plan *models.ReplicationSessionActionResource) error {
	switch action {
	case "sync":
		_, err := r.client.ReplicationSessionApi.ReplicationSessionSync(ctx, sessionID).Execute()
		return err
	case "pause":
		_, err := r.client.ReplicationSessionApi.ReplicationSessionPause(ctx, sessionID).Execute()
		return err
	case "resume":
		_, err := r.client.ReplicationSessionApi.ReplicationSessionResume(ctx, sessionID).Execute()
		return err
	case "failover":
		isPlanned := true
		if !plan.IsPlanned.IsNull() {
			isPlanned = plan.IsPlanned.ValueBool()
		}
		reverse := false
		if !plan.Reverse.IsNull() {
			reverse = plan.Reverse.ValueBool()
		}
		body := clientgen.ReplicationSessionFailover{
			IsPlanned: &isPlanned,
			Reverse:   &reverse,
		}
		_, err := r.client.ReplicationSessionApi.ReplicationSessionFailover(ctx, sessionID).Body(body).Execute()
		return err
	case "reprotect":
		_, err := r.client.ReplicationSessionApi.ReplicationSessionReprotect(ctx, sessionID).Body(clientgen.ReplicationSessionReprotect{}).Execute()
		return err
	case "start_failover_test":
		_, err := r.client.ReplicationSessionApi.ReplicationSessionStartFailoverTest(ctx, sessionID).Body(clientgen.ReplicationStartFailoverTest{}).Execute()
		return err
	case "stop_failover_test":
		_, _, err := r.client.ReplicationSessionApi.ReplicationSessionStopFailoverTest(ctx, sessionID).Execute()
		return err
	default:
		return fmt.Errorf("unsupported action: %s", action)
	}
}

// getReplicationSession fetches replication session details by ID
func (r *resourceReplicationSessionAction) getReplicationSession(ctx context.Context, sessionID string) (*clientgen.ReplicationSessionInstance, error) {
	sel := "*"
	queries := make(url.Values)
	queries.Set("select", sel)
	response, _, err := r.client.ReplicationSessionApi.GetReplicationSessionById(ctx, sessionID).Queries(queries).Execute()
	return response, err
}
