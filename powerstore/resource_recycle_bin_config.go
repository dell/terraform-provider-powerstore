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

	client "terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"

	"github.com/dell/gopowerstore/api"
	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// newRecycleBinConfigResource returns recycle bin config new resource instance
func newRecycleBinConfigResource() resource.Resource {
	return &resourceRecycleBinConfig{}
}

type resourceRecycleBinConfig struct {
	client *client.Client
}

// Metadata defines resource interface Metadata method
func (r *resourceRecycleBinConfig) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_recycle_bin_config"
}

// Schema defines resource interface Schema method
func (r *resourceRecycleBinConfig) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This resource is used to manage the recycle bin configuration of a PowerStore array. " +
			"The recycle bin allows recovery of intentionally or accidentally deleted volumes and volume groups " +
			"within a configurable retention period. Use this resource to configure the expiration duration. " +
			"We can Read and Update the recycle bin configuration using this resource. We can also import an existing recycle bin configuration from PowerStore array. " +
			"Note that creation and deletion of recycle bin configuration is not supported — the PowerStore system maintains a single configuration instance.",
		Description: "This resource is used to manage the recycle bin configuration of a PowerStore array. " +
			"The recycle bin allows recovery of intentionally or accidentally deleted volumes and volume groups " +
			"within a configurable retention period. Use this resource to configure the expiration duration. " +
			"We can Read and Update the recycle bin configuration using this resource. We can also import an existing recycle bin configuration from PowerStore array. " +
			"Note that creation and deletion of recycle bin configuration is not supported — the PowerStore system maintains a single configuration instance.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("0"),
				Description:         "The unique identifier of the recycle bin configuration. Always \"0\".",
				MarkdownDescription: "The unique identifier of the recycle bin configuration. Always `\"0\"`.",
			},
			"expiration_duration": schema.Int32Attribute{
				Required:            true,
				Description:         "Duration in days for items to remain in the recycle bin before automatic purging. Valid range is 0 to 30. A value of 0 indicates items expire immediately.",
				MarkdownDescription: "Duration in days for items to remain in the recycle bin before automatic purging. Valid range is `0` to `30`. A value of `0` indicates items expire immediately.",
				Validators: []validator.Int32{
					int32validator.Between(0, 30),
				},
			},
		},
	}
}

// Configure - defines configuration for recycle bin config resource
func (r *resourceRecycleBinConfig) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.client = c
}

// getRecycleBinConfig fetches the recycle bin configuration using APIClient
func (r *resourceRecycleBinConfig) getRecycleBinConfig(ctx context.Context, id string) (map[string]interface{}, error) {
	var result map[string]interface{}
	_, err := r.client.PStoreClient.APIClient().Query(
		ctx,
		api.RequestConfig{
			Method:   "GET",
			Endpoint: "recycle_bin_config",
			ID:       id,
		},
		&result)
	return result, err
}

// Create - recycle bin config always exists on the array; "create" reads the current state and applies the desired config.
func (r *resourceRecycleBinConfig) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	log.Printf("Started Creating Recycle Bin Config")

	var plan models.RecycleBinConfigResource
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	configID := plan.ID.ValueString()
	if configID == "" {
		configID = "0"
	}

	// Modify the recycle bin configuration
	duration := plan.ExpirationDuration.ValueInt32()
	payload := map[string]interface{}{
		"expiration_duration": duration,
	}
	_, err := r.client.PStoreClient.APIClient().Query(
		ctx,
		api.RequestConfig{
			Method:   "PATCH",
			Endpoint: "recycle_bin_config",
			ID:       configID,
			Body:     payload,
		},
		nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating recycle bin config",
			"Could not set recycle bin configuration: "+err.Error(),
		)
		return
	}

	// Read back the configuration
	getRes, err := r.getRecycleBinConfig(ctx, configID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading recycle bin config after creation",
			"Could not read recycle bin configuration: "+err.Error(),
		)
		return
	}

	state := models.RecycleBinConfigResource{
		ID:                 types.StringValue(getRes["id"].(string)),
		ExpirationDuration: types.Int32Value(int32(getRes["expiration_duration"].(float64))),
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	log.Printf("Successfully done with Create Recycle Bin Config")
}

// Read - reads the recycle bin configuration from the array
func (r *resourceRecycleBinConfig) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
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

	state.ID = types.StringValue(getRes["id"].(string))
	state.ExpirationDuration = types.Int32Value(int32(getRes["expiration_duration"].(float64)))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	log.Printf("Done with Read Recycle Bin Config")
}

// Update - updates the recycle bin configuration
func (r *resourceRecycleBinConfig) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("Started Update Recycle Bin Config")

	var plan models.RecycleBinConfigResource
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state models.RecycleBinConfigResource
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	configID := state.ID.ValueString()
	if configID == "" {
		configID = "0"
	}

	duration := plan.ExpirationDuration.ValueInt32()
	payload := map[string]interface{}{
		"expiration_duration": duration,
	}
	_, err := r.client.PStoreClient.APIClient().Query(
		ctx,
		api.RequestConfig{
			Method:   "PATCH",
			Endpoint: "recycle_bin_config",
			ID:       configID,
			Body:     payload,
		},
		nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating recycle bin config",
			"Could not update recycle bin configuration "+configID+": "+err.Error(),
		)
		return
	}

	// Read back the configuration
	getRes, err := r.getRecycleBinConfig(ctx, configID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading recycle bin config after update",
			"Could not read recycle bin configuration: "+err.Error(),
		)
		return
	}

	state.ID = types.StringValue(getRes["id"].(string))
	state.ExpirationDuration = types.Int32Value(int32(getRes["expiration_duration"].(float64)))

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	log.Printf("Successfully done with Update Recycle Bin Config")
}

// Delete - recycle bin config cannot be deleted; we remove it from state only.
func (r *resourceRecycleBinConfig) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("Recycle Bin Config Delete called — removing from Terraform state only (cannot delete on array)")
	// The recycle bin config always exists on the array and cannot be deleted.
	// Removing from Terraform state is all we can do.
}

// ImportState imports an existing recycle bin configuration into Terraform state
func (r *resourceRecycleBinConfig) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
