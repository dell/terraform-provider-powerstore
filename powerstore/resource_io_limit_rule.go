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
)

// newIoLimitRuleResource returns io_limit_rule resource instance
func newIoLimitRuleResource() resource.Resource {
	return &resourceIoLimitRule{}
}

type resourceIoLimitRule struct {
	client *clientgen.APIClient
}

// Metadata defines resource interface Metadata method
func (r *resourceIoLimitRule) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_io_limit_rule"
}

// Schema defines resource interface Schema method
func (r *resourceIoLimitRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This resource is used to manage the I/O limit rule entity of PowerStore Array. We can Create, Update and Delete the I/O limit rule using this resource. We can also import an existing I/O limit rule from PowerStore array.",
		Description:         "This resource is used to manage the I/O limit rule entity of PowerStore Array. We can Create, Update and Delete the I/O limit rule using this resource. We can also import an existing I/O limit rule from PowerStore array.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "Unique identifier of the I/O limit rule.",
				MarkdownDescription: "Unique identifier of the I/O limit rule.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},

			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Name of the I/O limit rule.",
				MarkdownDescription: "Name of the I/O limit rule.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"type": schema.StringAttribute{
				Required:            true,
				Description:         "Type of bandwidth limit (" + string(clientgen.BANDWIDTHLIMITTYPEENUM_ABSOLUTE) + " or " + string(clientgen.BANDWIDTHLIMITTYPEENUM_DENSITY) + "). " + string(clientgen.BANDWIDTHLIMITTYPEENUM_ABSOLUTE) + " limits are absolute values specified in I/O operations per second or bandwidth. " + string(clientgen.BANDWIDTHLIMITTYPEENUM_DENSITY) + " limits are per GB, e.g., I/O operations per second per GB.",
				MarkdownDescription: "Type of bandwidth limit (" + string(clientgen.BANDWIDTHLIMITTYPEENUM_ABSOLUTE) + " or " + string(clientgen.BANDWIDTHLIMITTYPEENUM_DENSITY) + "). " + string(clientgen.BANDWIDTHLIMITTYPEENUM_ABSOLUTE) + " limits are absolute values specified in I/O operations per second or bandwidth. " + string(clientgen.BANDWIDTHLIMITTYPEENUM_DENSITY) + " limits are per GB, e.g., I/O operations per second per GB.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.OneOf(
						string(clientgen.BANDWIDTHLIMITTYPEENUM_ABSOLUTE),
						string(clientgen.BANDWIDTHLIMITTYPEENUM_DENSITY),
					),
				},
			},

			"max_iops": schema.Int32Attribute{
				Optional:            true,
				Description:         "Maximum I/O operations in either I/O operations per second (IOPS) or I/O operations per second per GB. If type is set to " + string(clientgen.BANDWIDTHLIMITTYPEENUM_ABSOLUTE) + ", max_iops is specified in IOPS. If type is set to " + string(clientgen.BANDWIDTHLIMITTYPEENUM_DENSITY) + ", max_iops is specified in IOPS per GB.",
				MarkdownDescription: "Maximum I/O operations in either I/O operations per second (IOPS) or I/O operations per second per GB. If type is set to " + string(clientgen.BANDWIDTHLIMITTYPEENUM_ABSOLUTE) + ", max_iops is specified in IOPS. If type is set to " + string(clientgen.BANDWIDTHLIMITTYPEENUM_DENSITY) + ", max_iops is specified in IOPS per GB.",
			},

			"max_bw": schema.Int32Attribute{
				Optional:            true,
				Description:         "Maximum I/O bandwidth measured in either Kilobytes per second or Kilobytes per second / per GB. If type is set to " + string(clientgen.BANDWIDTHLIMITTYPEENUM_ABSOLUTE) + ", max_bw is specified in Kilobytes per second. If type is set to " + string(clientgen.BANDWIDTHLIMITTYPEENUM_DENSITY) + ", max_bw is specified in Kilobytes per second / per GB.",
				MarkdownDescription: "Maximum I/O bandwidth measured in either Kilobytes per second or Kilobytes per second / per GB. If type is set to " + string(clientgen.BANDWIDTHLIMITTYPEENUM_ABSOLUTE) + ", max_bw is specified in Kilobytes per second. If type is set to " + string(clientgen.BANDWIDTHLIMITTYPEENUM_DENSITY) + ", max_bw is specified in Kilobytes per second / per GB.",
			},

			"burst_percentage": schema.Int32Attribute{
				Optional:            true,
				Description:         "Percentage indicating by how much the limit may be exceeded. If I/O normally runs below the specified limit, then the volume or volume group will accumulate burst credits that can be used to exceed the limit for a short period.",
				MarkdownDescription: "Percentage indicating by how much the limit may be exceeded. If I/O normally runs below the specified limit, then the volume or volume group will accumulate burst credits that can be used to exceed the limit for a short period.",
			},
		},
	}
}

// Configure defines configuration for io_limit_rule resource
func (r *resourceIoLimitRule) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// Create method to create io_limit_rule resource
func (r *resourceIoLimitRule) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.IoLimitRule

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create New I/O Limit Rule
	ioLimitRuleCreateResponse, _, err := r.client.IoLimitRuleApi.PostAllIoLimitRules(ctx).Body(clientgen.IoLimitRuleCreate{
		Name:            plan.Name.ValueString(),
		Type:            clientgen.BandwidthLimitTypeEnum(plan.Type.ValueString()),
		MaxIops:         helper.ValueToPointer[int32](plan.MaxIops),
		MaxBw:           helper.ValueToPointer[int32](plan.MaxBw),
		BurstPercentage: helper.ValueToPointer[int32](plan.BurstPercentage),
	}).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating I/O limit rule",
			"Could not create I/O limit rule, unexpected error: "+err.Error(),
		)
		return
	}

	// Get I/O Limit Rule details using ID retrieved above
	ioLimitRuleResponse, err := r.ReadAPI(context.Background(), *ioLimitRuleCreateResponse.Id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting I/O limit rule after creation",
			"Could not get I/O limit rule, unexpected error: "+err.Error(),
		)
		return
	}

	result := r.updateIoLimitRuleState(ioLimitRuleResponse)

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Create")
}

// Delete method to delete io_limit_rule resource
func (r *resourceIoLimitRule) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("Started with the Delete")

	var state models.IoLimitRule
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get I/O Limit Rule ID from state
	ioLimitRuleID := state.ID.ValueString()

	// Delete I/O Limit Rule by calling API
	_, err := r.client.IoLimitRuleApi.DeleteIoLimitRuleById(ctx, ioLimitRuleID).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting I/O limit rule",
			"Could not delete I/O limit rule "+ioLimitRuleID+": "+err.Error(),
		)
		return
	}

	log.Printf("Done with Delete")
}

// Read method to read io_limit_rule resource information
func (r *resourceIoLimitRule) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	log.Printf("Reading I/O Limit Rule")
	var state models.IoLimitRule
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get I/O limit rule details from API and update what is in state from what the API returns
	id := state.ID.ValueString()
	response, err := r.ReadAPI(context.Background(), id)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading I/O limit rule",
			"Could not read I/O limit rule with error "+id+": "+err.Error(),
		)
		return
	}

	state = r.updateIoLimitRuleState(response)

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Read")
}

// Update method to update io_limit_rule resource
func (r *resourceIoLimitRule) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("Started Update")

	// Get plan values
	var plan models.IoLimitRule
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state models.IoLimitRule
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get I/O Limit Rule ID from state
	ioLimitRuleID := state.ID.ValueString()

	// Update I/O Limit Rule by calling API
	_, err := r.client.IoLimitRuleApi.PatchIoLimitRuleById(ctx, ioLimitRuleID).Body(clientgen.IoLimitRuleModify{
		Name:            helper.ValueToPointer[string](plan.Name),
		Type:            helper.ValueToPointer[clientgen.BandwidthLimitTypeEnum](plan.Type),
		MaxIops:         helper.ValueToPointer[int32](plan.MaxIops),
		MaxBw:           helper.ValueToPointer[int32](plan.MaxBw),
		BurstPercentage: helper.ValueToPointer[int32](plan.BurstPercentage),
	}).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating I/O limit rule",
			"Could not update I/O limit rule "+ioLimitRuleID+": "+err.Error(),
		)
		return
	}

	// Get I/O Limit Rule details
	getRes, err := r.ReadAPI(context.Background(), ioLimitRuleID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting I/O limit rule after update",
			"Could not get I/O limit rule, unexpected error: "+err.Error(),
		)
		return
	}

	state = r.updateIoLimitRuleState(getRes)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Successfully done with Update")
}

// ImportState import state for existing io_limit_rule
func (r *resourceIoLimitRule) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *resourceIoLimitRule) ReadAPI(ctx context.Context, id string) (*clientgen.IoLimitRuleInstance, error) {
	queries := make(url.Values)
	queries.Set("select", "id,name,type,max_iops,max_bw,burst_percentage,type_l10n")
	response, _, err := r.client.IoLimitRuleApi.GetIoLimitRuleById(context.Background(), id).Queries(queries).Execute()
	return response, err
}

// updateIoLimitRuleState method to update terraform state
func (r resourceIoLimitRule) updateIoLimitRuleState(ioLimitRuleResponse *clientgen.IoLimitRuleInstance) models.IoLimitRule {
	// Update value from I/O Limit Rule Response to State
	return models.IoLimitRule{
		ID:              helper.TfString(helper.SetDefault(ioLimitRuleResponse.Id, "")),
		Name:            helper.TfString(helper.SetDefault(ioLimitRuleResponse.Name, "")),
		Type:            helper.TfString(ioLimitRuleResponse.Type),
		MaxIops:         helper.TfInt32(helper.SetDefault(ioLimitRuleResponse.MaxIops, 0)),
		MaxBw:           helper.TfInt32(helper.SetDefault(ioLimitRuleResponse.MaxBw, 0)),
		BurstPercentage: helper.TfInt32(helper.SetDefault(ioLimitRuleResponse.BurstPercentage, 0)),
	}
}
