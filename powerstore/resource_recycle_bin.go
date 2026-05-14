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
	"terraform-provider-powerstore/models"

	"github.com/dell/gopowerstore/api"
	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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
	client *client.Client
}

// Metadata defines resource interface Metadata method
func (r *resourceRecycleBin) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_recycle_bin"
}

// Schema defines resource interface Schema method
func (r *resourceRecycleBin) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This resource is used to manage the PowerStore recycle bin. " +
			"It supports configuring the recycle bin expiration duration, recovering deleted volumes " +
			"and volume groups from the recycle bin, permanently deleting specific items, and emptying the entire recycle bin.\n\n" +
			"**Modes of operation:**\n" +
			"- **Configuration mode**: Set `expiration_duration` to configure the retention policy.\n" +
			"- **Item action mode**: Set `resource_id` (or `resource_name`) to recover or permanently delete a specific item. Use `action` to specify `recover` (default) or `delete`.\n" +
			"- **Empty mode**: Set `empty_recycle_bin = true` to permanently delete all items from the recycle bin.",
		Description: "This resource is used to manage the PowerStore recycle bin. " +
			"It supports configuring the recycle bin expiration duration, recovering deleted volumes " +
			"and volume groups from the recycle bin, permanently deleting specific items, and emptying the entire recycle bin.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The unique identifier of the resource instance.",
				MarkdownDescription: "The unique identifier of the resource instance.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"expiration_duration": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Duration in days for items to remain in the recycle bin before automatic purging. Valid range is 0 to 30. A value of 0 indicates items expire immediately. Used in configuration mode.",
				MarkdownDescription: "Duration in days for items to remain in the recycle bin before automatic purging. Valid range is `0` to `30`. A value of `0` indicates items expire immediately. Used in configuration mode.",
				Validators: []validator.Int32{
					int32validator.Between(0, 30),
					int32validator.ConflictsWith(
						path.MatchRoot("resource_id"),
						path.MatchRoot("resource_name"),
						path.MatchRoot("action"),
						path.MatchRoot("empty_recycle_bin"),
					),
				},
			},
			"resource_id": schema.StringAttribute{
				Optional:            true,
				Description:         "The unique identifier of a recycle bin item to recover or permanently delete.",
				MarkdownDescription: "The unique identifier of a recycle bin item to recover or permanently delete.",
				Validators: []validator.String{
					stringvalidator.ConflictsWith(
						path.MatchRoot("expiration_duration"),
						path.MatchRoot("resource_name"),
						path.MatchRoot("empty_recycle_bin"),
					),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"resource_name": schema.StringAttribute{
				Optional:            true,
				Description:         "The name of the deleted resource in the recycle bin. Used as an alternative to resource_id for identifying items.",
				MarkdownDescription: "The name of the deleted resource in the recycle bin. Used as an alternative to `resource_id`.",
				Validators: []validator.String{
					stringvalidator.ConflictsWith(
						path.MatchRoot("expiration_duration"),
						path.MatchRoot("resource_id"),
						path.MatchRoot("empty_recycle_bin"),
					),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"resource_type": schema.StringAttribute{
				Optional:            true,
				Description:         "The type of resource to filter when using resource_name. Valid values are 'volume' and 'volume_group'.",
				MarkdownDescription: "The type of resource to filter when using `resource_name`. Valid values are `volume` and `volume_group`.",
				Validators: []validator.String{
					stringvalidator.OneOf("volume", "volume_group"),
				},
			},
			"action": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The action to perform on the recycle bin item. Valid values are 'recover' (restore deleted item) and 'delete' (permanently delete item). Defaults to 'recover' when resource_id or resource_name is specified.",
				MarkdownDescription: "The action to perform on the recycle bin item. Valid values are `recover` and `delete`. Defaults to `recover` when `resource_id` or `resource_name` is specified.",
				Validators: []validator.String{
					stringvalidator.OneOf("recover", "delete"),
					stringvalidator.ConflictsWith(
						path.MatchRoot("expiration_duration"),
						path.MatchRoot("empty_recycle_bin"),
					),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"empty_recycle_bin": schema.BoolAttribute{
				Optional:            true,
				Description:         "When set to true, empties the entire recycle bin by permanently deleting all items.",
				MarkdownDescription: "When set to `true`, empties the entire recycle bin by permanently deleting all items.",
				Validators: []validator.Bool{
					boolvalidator.ConflictsWith(
						path.MatchRoot("expiration_duration"),
						path.MatchRoot("resource_id"),
						path.MatchRoot("resource_name"),
						path.MatchRoot("action"),
					),
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

	r.client = c
}

// ---- Helper methods ----

// getRecycleBinConfig fetches the recycle bin configuration
func (r *resourceRecycleBin) getRecycleBinConfig(ctx context.Context, id string) (map[string]interface{}, error) {
	var result map[string]interface{}
	qp := r.client.PStoreClient.APIClient().QueryParams()
	qp.Select("id", "expiration_duration")
	_, err := r.client.PStoreClient.APIClient().Query(
		ctx,
		api.RequestConfig{
			Method:      "GET",
			Endpoint:    "recycle_bin_config",
			ID:          id,
			QueryParams: qp,
		},
		&result)
	return result, err
}

// getRecycleBinItems fetches all recycle bin items
func (r *resourceRecycleBin) getRecycleBinItems(ctx context.Context) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	qp := r.client.PStoreClient.APIClient().QueryParams()
	_, err := r.client.PStoreClient.APIClient().Query(
		ctx,
		api.RequestConfig{
			Method:      "GET",
			Endpoint:    "recycle_bin",
			QueryParams: qp,
		},
		&result)
	return result, err
}

// recoverRecycleBinItem recovers a deleted item from the recycle bin
func (r *resourceRecycleBin) recoverRecycleBinItem(ctx context.Context, id string) error {
	_, err := r.client.PStoreClient.APIClient().Query(
		ctx,
		api.RequestConfig{
			Method:   "POST",
			Endpoint: "recycle_bin",
			ID:       id,
			Action:   "recover",
			Body:     map[string]interface{}{},
		},
		nil)
	return err
}

// deleteRecycleBinItem permanently deletes an item from the recycle bin
func (r *resourceRecycleBin) deleteRecycleBinItem(ctx context.Context, id string) error {
	_, err := r.client.PStoreClient.APIClient().Query(
		ctx,
		api.RequestConfig{
			Method:   "DELETE",
			Endpoint: "recycle_bin",
			ID:       id,
		},
		nil)
	return err
}

// emptyRecycleBin permanently deletes all items from the recycle bin
func (r *resourceRecycleBin) emptyRecycleBin(ctx context.Context) error {
	_, err := r.client.PStoreClient.APIClient().Query(
		ctx,
		api.RequestConfig{
			Method:   "POST",
			Endpoint: "recycle_bin",
			Action:   "empty",
		},
		nil)
	return err
}

// findRecycleBinItemByName finds a recycle bin item by name and optional resource type
func (r *resourceRecycleBin) findRecycleBinItemByName(ctx context.Context, name string, resourceType string) (map[string]interface{}, error) {
	items, err := r.getRecycleBinItems(ctx)
	if err != nil {
		return nil, err
	}
	var matches []map[string]interface{}
	for _, item := range items {
		itemName, ok := item["name"].(string)
		if !ok || itemName != name {
			continue
		}
		if resourceType == "" {
			matches = append(matches, item)
		} else if rt, ok := item["resource_type"].(string); ok && rt == resourceType {
			matches = append(matches, item)
		}
	}
	if len(matches) == 1 {
		return matches[0], nil
	} else if len(matches) > 1 {
		return nil, fmt.Errorf("multiple recycle bin items found with name '%s' — specify resource_type to narrow the search or use resource_id", name)
	}
	return nil, nil
}

// isConfigMode returns true if the plan/state is in configuration mode
func isConfigMode(state *models.RecycleBinResource) bool {
	return !state.ExpirationDuration.IsNull() && !state.ExpirationDuration.IsUnknown()
}

// isItemMode returns true if the plan/state is in item action mode
func isItemMode(state *models.RecycleBinResource) bool {
	return (!state.ResourceID.IsNull() && state.ResourceID.ValueString() != "") ||
		(!state.ResourceName.IsNull() && state.ResourceName.ValueString() != "")
}

// isEmptyMode returns true if the plan/state is in empty mode
func isEmptyMode(state *models.RecycleBinResource) bool {
	return !state.EmptyRecycleBin.IsNull() && state.EmptyRecycleBin.ValueBool()
}

// ---- CRUD methods ----

// Create defines resource interface Create method
func (r *resourceRecycleBin) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.RecycleBinResource
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if isConfigMode(&plan) {
		r.createConfig(ctx, &plan, resp)
	} else if isItemMode(&plan) {
		r.createItemAction(ctx, &plan, resp)
	} else if isEmptyMode(&plan) {
		r.createEmptyRecycleBin(ctx, &plan, resp)
	} else {
		resp.Diagnostics.AddError(
			"Invalid Configuration",
			"At least one of expiration_duration, resource_id, resource_name, or empty_recycle_bin must be specified.",
		)
	}
}

// createConfig handles configuration mode create (PATCH recycle_bin_config)
func (r *resourceRecycleBin) createConfig(ctx context.Context, plan *models.RecycleBinResource, resp *resource.CreateResponse) {
	log.Printf("Started Creating Recycle Bin Config")
	configID := "0"

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

	getRes, err := r.getRecycleBinConfig(ctx, configID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading recycle bin config after creation",
			"Could not read recycle bin configuration: "+err.Error(),
		)
		return
	}

	state := models.RecycleBinResource{
		ID: types.StringValue(getRes["id"].(string)),
	}
	if val, ok := getRes["expiration_duration"].(float64); ok {
		state.ExpirationDuration = types.Int32Value(int32(val))
	}

	diags := resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	log.Printf("Successfully done with Create Recycle Bin Config")
}

// createItemAction handles recover/delete of a specific recycle bin item
func (r *resourceRecycleBin) createItemAction(ctx context.Context, plan *models.RecycleBinResource, resp *resource.CreateResponse) {
	log.Printf("Started Recycle Bin Item Action")

	var itemID string

	if !plan.ResourceID.IsNull() && plan.ResourceID.ValueString() != "" {
		itemID = plan.ResourceID.ValueString()
	} else {
		// Look up by name
		resourceName := plan.ResourceName.ValueString()
		resourceType := ""
		if !plan.ResourceType.IsNull() {
			resourceType = plan.ResourceType.ValueString()
		}
		item, err := r.findRecycleBinItemByName(ctx, resourceName, resourceType)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error finding recycle bin item",
				err.Error(),
			)
			return
		}
		if item == nil {
			resp.Diagnostics.AddError(
				"Recycle bin item not found",
				fmt.Sprintf("No recycle bin item found with name '%s'", resourceName),
			)
			return
		}
		itemID = item["id"].(string)
	}

	// Determine action (default to recover)
	action := "recover"
	if !plan.Action.IsNull() && plan.Action.ValueString() != "" {
		action = plan.Action.ValueString()
	}

	switch action {
	case "recover":
		err := r.recoverRecycleBinItem(ctx, itemID)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error recovering recycle bin item",
				fmt.Sprintf("Could not recover item %s: %s", itemID, err.Error()),
			)
			return
		}
		log.Printf("Successfully recovered recycle bin item: %s", itemID)
	case "delete":
		err := r.deleteRecycleBinItem(ctx, itemID)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error deleting recycle bin item",
				fmt.Sprintf("Could not permanently delete item %s: %s", itemID, err.Error()),
			)
			return
		}
		log.Printf("Successfully deleted recycle bin item: %s", itemID)
	}

	state := models.RecycleBinResource{
		ID:         types.StringValue(itemID),
		ResourceID: types.StringValue(itemID),
		Action:     types.StringValue(action),
	}
	if !plan.ResourceName.IsNull() {
		state.ResourceName = plan.ResourceName
	}
	if !plan.ResourceType.IsNull() {
		state.ResourceType = plan.ResourceType
	}

	diags := resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	log.Printf("Successfully done with Recycle Bin Item Action: %s on %s", action, itemID)
}

// createEmptyRecycleBin handles emptying the entire recycle bin
func (r *resourceRecycleBin) createEmptyRecycleBin(ctx context.Context, plan *models.RecycleBinResource, resp *resource.CreateResponse) {
	log.Printf("Started Empty Recycle Bin")

	err := r.emptyRecycleBin(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error emptying recycle bin",
			"Could not empty recycle bin: "+err.Error(),
		)
		return
	}

	state := models.RecycleBinResource{
		ID:              types.StringValue("empty"),
		EmptyRecycleBin: types.BoolValue(true),
	}

	diags := resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	log.Printf("Successfully done with Empty Recycle Bin")
}

// Read defines resource interface Read method
func (r *resourceRecycleBin) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.RecycleBinResource
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if isConfigMode(&state) {
		// Config mode — read the current configuration from the array
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
		if val, ok := getRes["expiration_duration"].(float64); ok {
			state.ExpirationDuration = types.Int32Value(int32(val))
		}
	}
	// For item action and empty modes, state is one-shot — keep as-is

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	log.Printf("Done with Read Recycle Bin")
}

// Update defines resource interface Update method
func (r *resourceRecycleBin) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan models.RecycleBinResource
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if isConfigMode(&plan) {
		// Config mode update
		log.Printf("Started Update Recycle Bin Config")
		configID := "0"

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

		state := models.RecycleBinResource{
			ID: types.StringValue(getRes["id"].(string)),
		}
		if val, ok := getRes["expiration_duration"].(float64); ok {
			state.ExpirationDuration = types.Int32Value(int32(val))
		}

		diags = resp.State.Set(ctx, state)
		resp.Diagnostics.Append(diags...)
		log.Printf("Successfully done with Update Recycle Bin Config")
	}
	// Item and empty modes use RequiresReplace, so Update should not be called for them
}

// Delete defines resource interface Delete method
func (r *resourceRecycleBin) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("Recycle Bin Delete called — removing from Terraform state only")
	// For config mode: the recycle bin config always exists on the array and cannot be deleted.
	// For item/empty modes: the action was already performed during create.
	// In all cases, removing from Terraform state is all we do.
}

// ImportState defines resource interface ImportState method
func (r *resourceRecycleBin) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
