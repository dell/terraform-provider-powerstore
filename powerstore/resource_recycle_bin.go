/*
Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

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

	client "terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/models"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// newRecycleBinResource returns recycle bin new resource instance
func newRecycleBinResource() resource.Resource {
	return &resourceRecycleBin{}
}

type resourceRecycleBin struct {
	client *clientgen.APIClient
}

// Metadata defines resource interface Metadata method
func (r *resourceRecycleBin) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_recycle_bin_config"
}

// Schema defines resource interface Schema method
func (r *resourceRecycleBin) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This resource is used to manage the PowerStore recycle bin configuration. " +
			"It supports declaring the recycle bin retention period (`expiration_duration`), reading the current configuration, " +
			"detecting and correcting configuration drift, and importing existing configuration into Terraform state.\n\n" +
			"**Note:** Running `terraform destroy` on this resource only removes it from Terraform state. " +
			"The recycle bin configuration is a singleton on the PowerStore array and cannot be deleted.",
		Description: "This resource is used to manage the PowerStore recycle bin configuration. " +
			"It supports declaring the recycle bin retention period, reading, updating, and importing the configuration. " +
			"Running terraform destroy only removes from Terraform state, not from the PowerStore array.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The unique identifier of the recycle bin configuration (always '0').",
				MarkdownDescription: "The unique identifier of the recycle bin configuration (always `0`).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"expiration_duration": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Duration in days for items to remain in the recycle bin before automatic purging. Valid range is 0 to 30. A value of 0 indicates items expire immediately.",
				MarkdownDescription: "Duration in days for items to remain in the recycle bin before automatic purging. Valid range is `0` to `30`. A value of `0` indicates items expire immediately.",
				Validators: []validator.Int32{
					int32validator.Between(0, 30),
				},
			},
		},
	}
}

// Configure defines resource interface Configure method
func (r *resourceRecycleBin) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// ---- Helper methods ----

// getRecycleBinConfig fetches the recycle bin configuration from the array
func (r *resourceRecycleBin) getRecycleBinConfig(ctx context.Context, id string) (*clientgen.RecycleBinConfigInstance, error) {
	result, _, err := r.client.RecycleBinApi.GetRecycleBinConfigById(ctx, id).Execute()
	return result, err
}

// ---- CRUD methods ----

// Create defines resource interface Create method
func (r *resourceRecycleBin) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.RecycleBinConfigResource
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Started Creating Recycle Bin Config")
	configID := "0"
	duration := plan.ExpirationDuration.ValueInt32()

	_, err := r.client.RecycleBinApi.PatchRecycleBinConfigById(ctx, configID).Body(clientgen.RecycleBinConfigModify{
		ExpirationDuration: &duration,
	}).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating recycle bin config",
			"Could not set recycle bin configuration: "+err.Error(),
		)
		return
	}

	getRes, err := r.getRecycleBinConfig(ctx, configID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading recycle bin config after creation",
			"Could not read recycle bin configuration: "+err.Error(),
		)
		return
	}

	state := recycleBinConfigToState(getRes)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	log.Printf("Successfully done with Create Recycle Bin Config")
}

// Read defines resource interface Read method
func (r *resourceRecycleBin) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.RecycleBinConfigResource
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	configID := state.ID.ValueString()
	if configID == "" {
		configID = "0"
	}

	getRes, err := r.getRecycleBinConfig(ctx, configID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading recycle bin config",
			"Could not read recycle bin configuration "+configID+": "+err.Error(),
		)
		return
	}

	newState := recycleBinConfigToState(getRes)
	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
	log.Printf("Done with Read Recycle Bin Config")
}

// Update defines resource interface Update method
func (r *resourceRecycleBin) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan models.RecycleBinConfigResource
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Started Update Recycle Bin Config")
	configID := "0"
	duration := plan.ExpirationDuration.ValueInt32()

	_, err := r.client.RecycleBinApi.PatchRecycleBinConfigById(ctx, configID).Body(clientgen.RecycleBinConfigModify{
		ExpirationDuration: &duration,
	}).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating recycle bin config",
			"Could not update recycle bin configuration: "+err.Error(),
		)
		return
	}

	getRes, err := r.getRecycleBinConfig(ctx, configID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading recycle bin config after update",
			"Could not read recycle bin configuration: "+err.Error(),
		)
		return
	}

	state := recycleBinConfigToState(getRes)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	log.Printf("Successfully done with Update Recycle Bin Config")
}

// Delete defines resource interface Delete method
func (r *resourceRecycleBin) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("Recycle Bin Config Delete called — removing from Terraform state only")
	// The recycle bin configuration is a singleton on the array and cannot be deleted.
	// Removing from Terraform state is all we do.
}

// recycleBinConfigToState converts a RecycleBinConfigInstance to the Terraform state model
func recycleBinConfigToState(cfg *clientgen.RecycleBinConfigInstance) models.RecycleBinConfigResource {
	state := models.RecycleBinConfigResource{}
	if cfg.Id != nil {
		state.ID = types.StringValue(*cfg.Id)
	}
	if cfg.ExpirationDuration != nil {
		state.ExpirationDuration = types.Int32Value(*cfg.ExpirationDuration)
	}
	return state
}

// ImportState defines resource interface ImportState method
func (r *resourceRecycleBin) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
