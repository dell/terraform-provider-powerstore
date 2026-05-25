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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// newQosPolicyResource returns qos_policy resource instance
func newQosPolicyResource() resource.Resource {
	return &resourceQosPolicy{}
}

type resourceQosPolicy struct {
	client *clientgen.APIClient
}

// Metadata defines resource interface Metadata method
func (r *resourceQosPolicy) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_qos_policy"
}

// Schema defines resource interface Schema method
func (r *resourceQosPolicy) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This resource is used to manage the QoS policy entity of PowerStore Array. We can Create, Update and Delete the QoS policy using this resource. We can also import an existing QoS policy from PowerStore array.",
		Description:         "This resource is used to manage the QoS policy entity of PowerStore Array. We can Create, Update and Delete the QoS policy using this resource. We can also import an existing QoS policy from PowerStore array.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "Unique identifier of the QoS policy.",
				MarkdownDescription: "Unique identifier of the QoS policy.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},

			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Name of the QoS policy.",
				MarkdownDescription: "Name of the QoS policy.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Description of the QoS policy.",
				MarkdownDescription: "Description of the QoS policy.",
			},

			"type": schema.StringAttribute{
				Required:            true,
				Description:         "Type of the QoS policy. Must be '" + string(clientgen.POLICYTYPEENUM_QO_S) + "' for block storage or '" + string(clientgen.POLICYTYPEENUM_FILE_PERFORMANCE) + "' for file storage.",
				MarkdownDescription: "Type of the QoS policy. Must be '" + string(clientgen.POLICYTYPEENUM_QO_S) + "' for block storage or '" + string(clientgen.POLICYTYPEENUM_FILE_PERFORMANCE) + "' for file storage.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.OneOf(
						string(clientgen.POLICYTYPEENUM_QO_S),
						string(clientgen.POLICYTYPEENUM_FILE_PERFORMANCE),
					),
				},
			},

			"io_limit_rule_id": schema.StringAttribute{
				Optional:            true,
				Description:         "I/O limit rule identifier included in this policy. This attribute is only used for the QoS Performance Policy type.",
				MarkdownDescription: "I/O limit rule identifier included in this policy. This attribute is only used for the QoS Performance Policy type.",
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.Expressions{
						path.MatchRoot("file_io_limit_rule_id"),
					}...),
				},
			},

			"file_io_limit_rule_id": schema.StringAttribute{
				Optional:            true,
				Description:         "File I/O limit rule identifier included in this policy. This attribute is only valid for the File_Performance Policy type.",
				MarkdownDescription: "File I/O limit rule identifier included in this policy. This attribute is only valid for the File_Performance Policy type.",
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.Expressions{
						path.MatchRoot("io_limit_rule_id"),
					}...),
				},
			},
		},
	}
}

// Configure defines configuration for qos_policy resource
func (r *resourceQosPolicy) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = c.GenClient
}

// Create method to create qos_policy resource
func (r *resourceQosPolicy) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.QosPolicy

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create New QoS Policy
	qosPolicyCreateResponse, _, err := r.client.PolicyApi.PostAllPolicys(ctx).Body(clientgen.PolicyCreate{
		Name:              plan.Name.ValueString(),
		Description:       helper.ValueToPointer[string](plan.Description),
		IoLimitRuleId:     helper.ValueToPointer[string](plan.IoLimitRuleId),
		FileIoLimitRuleId: helper.ValueToPointer[string](plan.FileIoLimitRuleId),
	}).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating QoS policy",
			"Could not create QoS policy, unexpected error: "+err.Error(),
		)
		return
	}

	// Get QoS Policy details using ID retrieved above
	qosPolicyResponse, err := r.ReadAPI(context.Background(), *qosPolicyCreateResponse.Id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting QoS policy after creation",
			"Could not get QoS policy, unexpected error: "+err.Error(),
		)
		return
	}

	result := r.updateQosPolicyState(qosPolicyResponse)

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Create")
}

// Delete method to delete qos_policy resource
func (r *resourceQosPolicy) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("Started with the Delete")

	var state models.QosPolicy
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get QoS Policy ID from state
	qosPolicyID := state.ID.ValueString()

	// Delete QoS Policy by calling API
	_, err := r.client.PolicyApi.DeletePolicyById(ctx, qosPolicyID).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting QoS policy",
			"Could not delete QoS policy "+qosPolicyID+": "+err.Error(),
		)
		return
	}

	log.Printf("Done with Delete")
}

// Read method to read qos_policy resource information
func (r *resourceQosPolicy) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	log.Printf("Reading QoS Policy")
	var state models.QosPolicy
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get QoS policy details from API and update what is in state from what the API returns
	id := state.ID.ValueString()
	response, err := r.ReadAPI(context.Background(), id)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading QoS policy",
			"Could not read QoS policy with error "+id+": "+err.Error(),
		)
		return
	}

	state = r.updateQosPolicyState(response)

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Read")
}

// Update method to update qos_policy resource
func (r *resourceQosPolicy) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("Started Update")

	// Get plan values
	var plan models.QosPolicy
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state models.QosPolicy
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get QoS Policy ID from state
	qosPolicyID := state.ID.ValueString()

	// Update QoS Policy by calling API
	_, err := r.client.PolicyApi.PatchPolicyById(ctx, qosPolicyID).Body(clientgen.PolicyModify{
		Name:        helper.ValueToPointer[string](plan.Name),
		Description: helper.ValueToPointer[string](plan.Description),
		IoLimitRuleId: func() *string {
			if plan.Type.ValueString() == string(clientgen.POLICYTYPEENUM_QO_S) {
				return helper.ValueToPointer[string](plan.IoLimitRuleId)
			}
			return nil
		}(),
		FileIoLimitRuleId: func() *string {
			if plan.Type.ValueString() == string(clientgen.POLICYTYPEENUM_FILE_PERFORMANCE) {
				return helper.ValueToPointer[string](plan.FileIoLimitRuleId)
			}
			return nil
		}(),
	}).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating QoS policy",
			"Could not update QoS policy "+qosPolicyID+": "+err.Error(),
		)
		return
	}

	// Get QoS Policy details
	getRes, err := r.ReadAPI(context.Background(), qosPolicyID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting QoS policy after update",
			"Could not get QoS policy, unexpected error: "+err.Error(),
		)
		return
	}

	state = r.updateQosPolicyState(getRes)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Successfully done with Update")
}

// ImportState import state for existing qos_policy
func (r *resourceQosPolicy) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *resourceQosPolicy) ReadAPI(ctx context.Context, id string) (*clientgen.PolicyInstance, error) {
	queries := make(url.Values)
	queries.Set("select", "id,name,description,type,managed_by,managed_by_id,is_read_only,is_replica,file_io_limit_rule_id,io_limit_rule(id)")
	response, _, err := r.client.PolicyApi.GetPolicyById(context.Background(), id).Queries(queries).Execute()
	return response, err
}

// updateQosPolicyState method to update terraform state
func (r resourceQosPolicy) updateQosPolicyState(qosPolicyResponse *clientgen.PolicyInstance) models.QosPolicy {
	return models.QosPolicy{
		ID:                helper.TfString(helper.SetDefault(qosPolicyResponse.Id, "")),
		Name:              helper.TfString(helper.SetDefault(qosPolicyResponse.Name, "")),
		Description:       helper.TfString(helper.SetDefault(qosPolicyResponse.Description, "")),
		Type:              helper.TfString(qosPolicyResponse.Type),
		IoLimitRuleId:     helper.TfObject(qosPolicyResponse.IoLimitRule, func(r clientgen.IoLimitRuleInstance) types.String { return helper.TfString(r.Id) }),
		FileIoLimitRuleId: helper.TfString(helper.SetDefault(qosPolicyResponse.FileIoLimitRuleId, "")),
	}
}
