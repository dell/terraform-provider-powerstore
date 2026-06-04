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

// newFileIoLimitRuleResource returns file_io_limit_rule resource instance
func newFileIoLimitRuleResource() resource.Resource {
	return &resourceFileIoLimitRule{}
}

type resourceFileIoLimitRule struct {
	client *clientgen.APIClient
}

// Metadata defines resource interface Metadata method
func (r *resourceFileIoLimitRule) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_file_io_limit_rule"
}

// Schema defines resource interface Schema method
func (r *resourceFileIoLimitRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This resource is used to manage the file I/O limit rule entity of PowerStore Array. We can Create, Update and Delete the file I/O limit rule using this resource. We can also import an existing file I/O limit rule from PowerStore array.",
		Description:         "This resource is used to manage the file I/O limit rule entity of PowerStore Array. We can Create, Update and Delete the file I/O limit rule using this resource. We can also import an existing file I/O limit rule from PowerStore array.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "Unique identifier of the file I/O limit rule.",
				MarkdownDescription: "Unique identifier of the file I/O limit rule.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},

			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Name of the file I/O limit rule.",
				MarkdownDescription: "Name of the file I/O limit rule.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"max_bw": schema.Int32Attribute{
				Required:            true,
				Description:         "Maximum allowable bandwidth in MB/second for the file_system or nas_server resource. When applied to a nas_server, all the file_systems of the nas_server share the maximum bandwidth.",
				MarkdownDescription: "Maximum allowable bandwidth in MB/second for the file_system or nas_server resource. When applied to a nas_server, all the file_systems of the nas_server share the maximum bandwidth.",
			},
		},
	}
}

// Configure defines configuration for file_io_limit_rule resource
func (r *resourceFileIoLimitRule) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// Create method to create file_io_limit_rule resource
func (r *resourceFileIoLimitRule) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.FileIoLimitRule

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create New File I/O Limit Rule
	fileIoLimitRuleCreateResponse, _, err := r.client.FileIoLimitRuleApi.
		PostAllFileIoLimitRules(ctx).
		Body(clientgen.FileIoLimitRuleCreate{
			Name:  plan.Name.ValueString(),
			MaxBw: plan.MaxBw.ValueInt32(),
		}).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating file I/O limit rule",
			"Could not create file I/O limit rule, unexpected error: "+err.Error(),
		)
		return
	}

	// Get File I/O Limit Rule details using ID retrieved above
	fileIoLimitRuleResponse, err := r.ReadAPI(context.Background(), *fileIoLimitRuleCreateResponse.Id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting file I/O limit rule after creation",
			"Could not get file I/O limit rule, unexpected error: "+err.Error(),
		)
		return
	}

	result := r.updateFileIoLimitRuleState(fileIoLimitRuleResponse)

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Create")
}

// Delete method to delete file_io_limit_rule resource
func (r *resourceFileIoLimitRule) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("Started with the Delete")

	var state models.FileIoLimitRule
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get File I/O Limit Rule ID from state
	fileIoLimitRuleID := state.ID.ValueString()

	// Delete File I/O Limit Rule by calling API
	_, err := r.client.FileIoLimitRuleApi.DeleteFileIoLimitRuleById(ctx, fileIoLimitRuleID).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting file I/O limit rule",
			"Could not delete file I/O limit rule "+fileIoLimitRuleID+": "+err.Error(),
		)
		return
	}

	log.Printf("Done with Delete")
}

// Read method to read file_io_limit_rule resource information
func (r *resourceFileIoLimitRule) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	log.Printf("Reading File I/O Limit Rule")
	var state models.FileIoLimitRule
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get file I/O limit rule details from API and update what is in state from what the API returns
	id := state.ID.ValueString()
	response, err := r.ReadAPI(context.Background(), id)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading file I/O limit rule",
			"Could not read file I/O limit rule with error "+id+": "+err.Error(),
		)
		return
	}

	state = r.updateFileIoLimitRuleState(response)

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Read")
}

// Update method to update file_io_limit_rule resource
func (r *resourceFileIoLimitRule) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("Started Update")

	// Get plan values
	var plan models.FileIoLimitRule
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state models.FileIoLimitRule
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get File I/O Limit Rule ID from state
	fileIoLimitRuleID := state.ID.ValueString()

	// Update File I/O Limit Rule by calling API
	_, err := r.client.FileIoLimitRuleApi.PatchFileIoLimitRuleById(ctx, fileIoLimitRuleID).Body(clientgen.FileIoLimitRuleModify{
		Name:  helper.ValueToPointer[string](plan.Name),
		MaxBw: helper.ValueToPointer[int32](plan.MaxBw),
	}).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating file I/O limit rule",
			"Could not update file I/O limit rule "+fileIoLimitRuleID+": "+err.Error(),
		)
		return
	}

	// Get File I/O Limit Rule details
	getRes, err := r.ReadAPI(context.Background(), fileIoLimitRuleID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting file I/O limit rule after update",
			"Could not get file I/O limit rule, unexpected error: "+err.Error(),
		)
		return
	}

	state = r.updateFileIoLimitRuleState(getRes)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Successfully done with Update")
}

// ImportState import state for existing file_io_limit_rule
func (r *resourceFileIoLimitRule) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *resourceFileIoLimitRule) ReadAPI(ctx context.Context, id string) (*clientgen.FileIoLimitRuleInstance, error) {
	queries := make(url.Values)
	queries.Set("select", "id,name,max_bw")
	response, _, err := r.client.FileIoLimitRuleApi.GetFileIoLimitRuleById(context.Background(), id).Queries(queries).Execute()
	return response, err
}

// updateFileIoLimitRuleState method to update terraform state
func (r resourceFileIoLimitRule) updateFileIoLimitRuleState(fileIoLimitRuleResponse *clientgen.FileIoLimitRuleInstance) models.FileIoLimitRule {
	// Update value from File I/O Limit Rule Response to State
	return models.FileIoLimitRule{
		ID:    helper.TfString(helper.SetDefault(fileIoLimitRuleResponse.Id, "")),
		Name:  helper.TfString(helper.SetDefault(fileIoLimitRuleResponse.Name, "")),
		MaxBw: helper.TfInt32(helper.SetDefault(fileIoLimitRuleResponse.MaxBw, 0)),
	}
}
