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
	"net/url"

	client "terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/models"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
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
		MarkdownDescription: "This resource is used to manage the PowerStore recycle bin. It supports three modes of operation:\n\n" +
			"1. **Config mode**: Set `expiration_duration` (0–30 days) to configure the retention policy.\n" +
			"2. **Item action mode**: Specify `resource_id` or `resource_name` with `action` (`recover` or `delete`) to recover or permanently delete items.\n" +
			"3. **Empty mode**: Set `empty_recycle_bin = true` to permanently delete all items from the recycle bin.\n\n" +
			"**Note:** Running `terraform destroy` on this resource only removes it from Terraform state.",
		Description: "This resource is used to manage the PowerStore recycle bin. " +
			"It supports configuring the retention policy, recovering or deleting items, and emptying the recycle bin. " +
			"Running terraform destroy only removes from Terraform state.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The unique identifier of the resource.",
				MarkdownDescription: "The unique identifier of the resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"expiration_duration": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Duration in days for items to remain in the recycle bin before automatic purging. Valid range is 0 to 30.",
				MarkdownDescription: "Duration in days for items to remain in the recycle bin before automatic purging. Valid range is `0` to `30`.",
				Validators: []validator.Int32{
					int32validator.Between(0, 30),
				},
			},
			"resource_id": schema.StringAttribute{
				Optional:            true,
				Description:         "The unique identifier of the recycle bin item to recover or delete.",
				MarkdownDescription: "The unique identifier of the recycle bin item to recover or delete.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("expiration_duration")),
					stringvalidator.ConflictsWith(path.MatchRoot("resource_name")),
					stringvalidator.ConflictsWith(path.MatchRoot("empty_recycle_bin")),
				},
			},
			"resource_name": schema.StringAttribute{
				Optional:            true,
				Description:         "The name of the deleted resource in the recycle bin.",
				MarkdownDescription: "The name of the deleted resource in the recycle bin.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("expiration_duration")),
					stringvalidator.ConflictsWith(path.MatchRoot("resource_id")),
					stringvalidator.ConflictsWith(path.MatchRoot("empty_recycle_bin")),
				},
			},
			"resource_type": schema.StringAttribute{
				Optional:            true,
				Description:         "The type of resource to filter when using resource_name. Valid values: volume, volume_group.",
				MarkdownDescription: "The type of resource to filter when using `resource_name`. Valid values: `volume`, `volume_group`.",
				Validators: []validator.String{
					stringvalidator.OneOf("volume", "volume_group"),
				},
			},
			"action": schema.StringAttribute{
				Optional:            true,
				Description:         "The action to perform on the recycle bin item. Valid values: recover, delete.",
				MarkdownDescription: "The action to perform on the recycle bin item. Valid values: `recover`, `delete`.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("recover", "delete"),
					stringvalidator.ConflictsWith(path.MatchRoot("expiration_duration")),
					stringvalidator.ConflictsWith(path.MatchRoot("empty_recycle_bin")),
				},
			},
			"empty_recycle_bin": schema.BoolAttribute{
				Optional:            true,
				Description:         "When set to true, empties the entire recycle bin by permanently deleting all items.",
				MarkdownDescription: "When set to `true`, empties the entire recycle bin by permanently deleting all items.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Validators: []validator.Bool{},
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

// getRecycleBinConfig fetches the recycle bin configuration from the array.
func (r *resourceRecycleBin) getRecycleBinConfig(ctx context.Context, id string) (*clientgen.RecycleBinConfigInstance, error) {
	queries := make(url.Values)
	queries.Set("select", "id,expiration_duration")
	result, _, err := r.client.RecycleBinConfigApi.GetRecycleBinConfigById(ctx, id).Queries(queries).Execute()
	return result, err
}

// resolveRecycleBinItem finds a recycle bin item by ID or by name+type.
func (r *resourceRecycleBin) resolveRecycleBinItem(ctx context.Context, plan models.RecycleBinConfigResource) (string, error) {
	if !plan.ResourceID.IsNull() && plan.ResourceID.ValueString() != "" {
		return plan.ResourceID.ValueString(), nil
	}

	if !plan.ResourceName.IsNull() && plan.ResourceName.ValueString() != "" {
		items, _, err := r.client.RecycleBinApi.GetAllRecycleBins(ctx).Execute()
		if err != nil {
			return "", fmt.Errorf("could not list recycle bin items: %s", err.Error())
		}

		name := plan.ResourceName.ValueString()
		resourceType := ""
		if !plan.ResourceType.IsNull() {
			resourceType = plan.ResourceType.ValueString()
		}

		var matches []clientgen.RecycleBinInstance
		for _, item := range items {
			if item.Name != nil && *item.Name == name {
				if resourceType == "" || (item.ResourceType != nil && string(*item.ResourceType) == resourceType) {
					matches = append(matches, item)
				}
			}
		}

		if len(matches) == 0 {
			return "", fmt.Errorf("recycle bin item not found with name '%s'", name)
		}
		if len(matches) > 1 {
			return "", fmt.Errorf("multiple recycle bin items found with name '%s', specify resource_type to narrow", name)
		}
		return *matches[0].Id, nil
	}

	return "", fmt.Errorf("either resource_id or resource_name must be specified")
}

// ---- CRUD methods ----

// Create defines resource interface Create method
func (r *resourceRecycleBin) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.RecycleBinConfigResource
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	// Mode 1: Empty recycle bin
	if !plan.EmptyRecycleBin.IsNull() && plan.EmptyRecycleBin.ValueBool() {
		log.Printf("Started Empty Recycle Bin")
		_, err := r.client.RecycleBinApi.PostRecycleBinById(ctx).Execute()
		if err != nil {
			resp.Diagnostics.AddError(
				"Error emptying recycle bin",
				"Could not empty the recycle bin: "+err.Error(),
			)
			return
		}
		plan.ID = types.StringValue("empty")
		plan.ExpirationDuration = types.Int32Null()
		diags = resp.State.Set(ctx, &plan)
		resp.Diagnostics.Append(diags...)
		log.Printf("Successfully emptied Recycle Bin")
		return
	}

	// Mode 2: Item action (recover or delete)
	if !plan.Action.IsNull() && plan.Action.ValueString() != "" {
		log.Printf("Started Recycle Bin Item Action: %s", plan.Action.ValueString())

		itemID, err := r.resolveRecycleBinItem(ctx, plan)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error resolving recycle bin item",
				err.Error(),
			)
			return
		}

		action := plan.Action.ValueString()
		if action == "recover" {
			_, err = r.client.RecycleBinApi.RecycleBinRecover(ctx, itemID).Body(map[string]interface{}{}).Execute()
			if err != nil {
				resp.Diagnostics.AddError(
					"Error recovering recycle bin item",
					"Could not recover recycle bin item "+itemID+": "+err.Error(),
				)
				return
			}
		} else if action == "delete" {
			_, err = r.client.RecycleBinApi.DeleteRecycleBinById(ctx, itemID).Execute()
			if err != nil {
				resp.Diagnostics.AddError(
					"Error deleting recycle bin item",
					"Could not delete recycle bin item "+itemID+": "+err.Error(),
				)
				return
			}
		}

		plan.ID = types.StringValue(itemID)
		plan.ExpirationDuration = types.Int32Null()
		diags = resp.State.Set(ctx, &plan)
		resp.Diagnostics.Append(diags...)
		log.Printf("Successfully completed Recycle Bin Item Action: %s on %s", action, itemID)
		return
	}

	// Mode 3: Config mode (default)
	log.Printf("Started Creating Recycle Bin Config")
	configID := "0"
	duration := plan.ExpirationDuration.ValueInt32()

	_, err := r.client.RecycleBinConfigApi.PatchRecycleBinConfigById(ctx, configID).Body(clientgen.RecycleBinConfigModify{
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

	// For item action or empty modes, the operation already happened — preserve state as-is
	if !state.EmptyRecycleBin.IsNull() || !state.Action.IsNull() {
		log.Printf("Done with Read Recycle Bin (action/empty mode — state preserved)")
		return
	}

	// Config mode — refresh from array
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

	log.Printf("Started Update Recycle Bin Config")
	configID := "0"
	duration := plan.ExpirationDuration.ValueInt32()

	_, err := r.client.RecycleBinConfigApi.PatchRecycleBinConfigById(ctx, configID).Body(clientgen.RecycleBinConfigModify{
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
	// Removing from Terraform state is all we do for all modes.
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
