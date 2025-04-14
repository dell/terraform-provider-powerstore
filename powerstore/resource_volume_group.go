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
	"log"
	"net/url"
	"strings"

	client "terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/models"
	"terraform-provider-powerstore/powerstore/helper"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// newVolumeGroupResource returns new volume group resource instance
func newVolumeGroupResource() resource.Resource {
	return &resourceVolumeGroup{}
}

type resourceVolumeGroup struct {
	client    *clientgen.APIClient
	allclient *client.Client
}

// Metadata defines resource interface Metadata method
func (r *resourceVolumeGroup) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_volumegroup"
}

// Schema defines resource interface Schema method
func (r *resourceVolumeGroup) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "This resource is used to manage the volumegroup entity of PowerStore Array. We can Create, Update and Delete the volumegroup using this resource. We can also import an existing host from volumegroup array.",
		Description:         "This resource is used to manage the volumegroup entity of PowerStore Array. We can Create, Update and Delete the volumegroup using this resource. We can also import an existing host from volumegroup array.",

		Attributes: map[string]schema.Attribute{

			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "Unique identifier of the volume group.",
				MarkdownDescription: "Unique identifier of the volume group.",
			},

			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Name of the volume group.",
				MarkdownDescription: "Name of the volume group.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Description for the volume group.",
				MarkdownDescription: "Description for the volume group.",
			},

			"volume_ids": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "A list of identifiers of existing volumes that should be added to the volume group. Conflicts with `volume_names`.",
				MarkdownDescription: "A list of identifiers of existing volumes that should be added to the volume group. Conflicts with `volume_names`.",
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(
						stringvalidator.LengthAtLeast(1),
					),
					setvalidator.ConflictsWith(path.Expressions{
						path.MatchRoot("volume_names"),
					}...),
				},
			},

			"is_write_order_consistent": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Determines whether snapshot sets of the group will be write order consistent.",
				MarkdownDescription: "Determines whether snapshot sets of the group will be write order consistent.",
			},

			"protection_policy_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Unique identifier of the protection policy assigned to the volume group. Conflicts with `protection_policy_name`.",
				MarkdownDescription: "Unique identifier of the protection policy assigned to the volume group. Conflicts with `protection_policy_name`.",
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.Expressions{
						path.MatchRoot("protection_policy_name"),
					}...),
				},
			},

			"volume_names": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "A list of names of existing volumes that should be added to the volume group. Conflicts with `volume_ids`.",
				MarkdownDescription: "A list of names of existing volumes that should be added to the volume group. Conflicts with `volume_ids`.",
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(
						stringvalidator.LengthAtLeast(1),
					),
				},
			},

			"protection_policy_name": schema.StringAttribute{
				Optional:            true,
				Description:         "Unique name of the protection policy assigned to the volume group. Conflicts with `protection_policy_id`.",
				MarkdownDescription: "Unique name of the protection policy assigned to the volume group. Conflicts with `protection_policy_id`.",
			},
		},
	}
}

// Configure - defines configuration for volume group resource
func (r *resourceVolumeGroup) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client.GenClient
	r.allclient = client
}

// Create - method to create volume group resource
func (r *resourceVolumeGroup) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.Volumegroup

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	errmsg := r.fetchByName(&plan)
	if errmsg != "" {
		resp.Diagnostics.AddError(
			"Error creating volume group",
			"Could not create volume group, unexpected error: "+errmsg+"",
		)
		return
	}

	var volumeIds []string
	for _, volume := range plan.VolumeIDs.Elements() {
		volumeIds = append(volumeIds, strings.Trim(volume.String(), "\""))
	}

	volumeGroupCreate := clientgen.VolumeGroupCreate{
		Name:                   plan.Name.ValueString(),
		Description:            helper.ValueToPointer[string](plan.Description),
		VolumeIds:              volumeIds,
		IsWriteOrderConsistent: helper.GetKnownBoolPointer(plan.IsWriteOrderConsistent),
		ProtectionPolicyId:     helper.ValueToPointer[string](plan.ProtectionPolicyID),
	}

	//Create New Volume Group
	volGroupCreateResponse, _, err := r.client.VolumeGroupApi.PostAllVolumeGroups(ctx).Body(volumeGroupCreate).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating volume group",
			"Could not create volume group, unexpected error: "+err.Error(),
		)
		return
	}

	//Get Volume Group details using ID retrived above
	volGroupResponse, err := r.ReadAPI(context.Background(), *volGroupCreateResponse.Id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting volume group after creation",
			"Could not get volume group, unexpected error: "+err.Error(),
		)
		return
	}

	result := models.Volumegroup{}
	r.updateVolGroupState(&result, volGroupResponse, &plan)

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Create")
}

// Delete - method to delete volume group resource
func (r *resourceVolumeGroup) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("Started with the Delete")

	var state models.Volumegroup
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Get Volume Group ID from state
	volumeGroupID := state.ID.ValueString()

	//Get Volume Group details using ID retrived above
	volGroupResponse, err := r.allclient.PStoreClient.GetVolumeGroup(context.Background(), volumeGroupID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting volume group after creation",
			"Could not get volume group, unexpected error: "+err.Error(),
		)
		return
	}

	//Remove protection policy from volume group if present
	if volGroupResponse.ProtectionPolicyID != "" {
		volGroupChangePolicy := &gopowerstore.VolumeGroupChangePolicy{
			ProtectionPolicyID: "",
		}
		_, err = r.allclient.PStoreClient.UpdateVolumeGroupProtectionPolicy(context.Background(), volumeGroupID, volGroupChangePolicy)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error deleting volume group",
				"Could not remove protection policy from volume group "+volumeGroupID+": "+err.Error(),
			)
			return
		}
	}

	//Remove volume(s) from volume group if present
	if len(volGroupResponse.Volumes) != 0 {
		var volumeIDs []string
		for _, vol := range volGroupResponse.Volumes {
			volumeIDs = append(volumeIDs, vol.ID)
		}
		volGroupMembers := &gopowerstore.VolumeGroupMembers{
			VolumeIDs: volumeIDs,
		}
		_, err = r.allclient.PStoreClient.RemoveMembersFromVolumeGroup(context.Background(), volGroupMembers, volumeGroupID)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error deleting volume group",
				"Could not remove volume from volume group "+volumeGroupID+": "+err.Error(),
			)
			return
		}
	}

	//Delete Volume Group by calling API
	_, err = r.client.VolumeGroupApi.DeleteVolumeGroupById(ctx, volumeGroupID).Body(clientgen.VolumeGroupDelete{}).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting volume group",
			"Could not delete volumeGroupID "+volumeGroupID+": "+err.Error(),
		)
		return
	}

	log.Printf("Done with Delete")
}

// Read - method to read volume group resource
func (r *resourceVolumeGroup) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	log.Printf("Reading Volume Group")
	var state models.Volumegroup
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Get volume group details from API and update what is in state from what the API returns
	id := state.ID.ValueString()
	response, err := r.ReadAPI(context.Background(), id)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading volume group",
			"Could not read volume group with error "+id+": "+err.Error(),
		)
		return
	}

	r.updateVolGroupState(&state, response, &state)

	//Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Read")
}

func (r *resourceVolumeGroup) ReadAPI(ctx context.Context, id string) (*clientgen.VolumeGroupInstance, error) {
	sel := "*,volumes(*),protection_policy(*),protection_data,location_history,migration_session(*)"
	queries := make(url.Values)
	queries.Set("select", sel)
	response, _, err := r.client.VolumeGroupApi.GetVolumeGroupById(context.Background(), id).Queries(queries).Execute()
	return response, err
}

// Update - method to update volume group resource
func (r *resourceVolumeGroup) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("Started Update")

	//Get plan values
	var plan models.Volumegroup
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Get current state
	var state models.Volumegroup
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	errmsg := r.fetchByName(&plan)
	if errmsg != "" {
		resp.Diagnostics.AddError(
			"Error updating volume group",
			"Could not update volume group, unexpected error: "+errmsg+"",
		)
		return
	}

	//Get volume ids from plan into a slice
	var planVolumeIds []string
	for _, volume := range plan.VolumeIDs.Elements() {
		planVolumeIds = append(planVolumeIds, strings.Trim(volume.String(), "\""))
	}

	//Get volume ids from state into a slice
	var stateVolumeIds []string
	for _, volume := range state.VolumeIDs.Elements() {
		stateVolumeIds = append(stateVolumeIds, strings.Trim(volume.String(), "\""))
	}

	//Get volume ids from plan into map for optimized element search
	planVolumeIdsMap := make(map[string]string)
	if len(planVolumeIds) != 0 {
		for i := 0; i < len(planVolumeIds); i++ {
			planVolumeIdsMap[planVolumeIds[i]] = planVolumeIds[i]
		}
	}

	//Get volume ids from state into a map for optimized element search
	stateVolumeIdsMap := make(map[string]string)
	if len(stateVolumeIds) != 0 {
		for i := 0; i < len(stateVolumeIds); i++ {
			stateVolumeIdsMap[stateVolumeIds[i]] = stateVolumeIds[i]
		}
	}

	//Create map of volume ids to be removed by comparing plan and state volume ids
	removeVolumeIdsMap := make(map[string]string)
	for i := 0; i < len(stateVolumeIds); i++ {
		_, found := planVolumeIdsMap[stateVolumeIds[i]]
		if !found {
			log.Printf("Volume not found in state")
			removeVolumeIdsMap[stateVolumeIds[i]] = stateVolumeIds[i]
		}
	}

	//Get list of volume ids to be removed into a slice
	removeVolumeIdsSlice := []string{}
	for _, volumeID := range removeVolumeIdsMap {
		removeVolumeIdsSlice = append(removeVolumeIdsSlice, volumeID)
	}

	//Create map of volume ids to be added by comparing plan and state volume ids
	addVolumeIdsMap := make(map[string]string)
	for i := 0; i < len(planVolumeIds); i++ {
		_, found := stateVolumeIdsMap[planVolumeIds[i]]
		if !found {
			log.Printf("Volume not found in plan")
			addVolumeIdsMap[planVolumeIds[i]] = planVolumeIds[i]
		}
	}

	//Get list of volume ids to be added into a slice
	addVolumeIdsSlice := []string{}
	for _, volumeID := range addVolumeIdsMap {
		addVolumeIdsSlice = append(addVolumeIdsSlice, volumeID)
	}

	removeVolumeGroupMembers := &gopowerstore.VolumeGroupMembers{
		VolumeIDs: removeVolumeIdsSlice,
	}

	addVolumeGroupMembers := &gopowerstore.VolumeGroupMembers{
		VolumeIDs: addVolumeIdsSlice,
	}

	// Get Volume Group ID from from state
	volumeGroupID := state.ID.ValueString()

	volumeGroupUpdate := clientgen.VolumeGroupModify{
		Description:            helper.ValueToPointer[string](plan.Description),
		ProtectionPolicyId:     helper.ValueToPointer[string](plan.ProtectionPolicyID),
		Name:                   helper.ValueToPointer[string](plan.Name),
		IsWriteOrderConsistent: helper.ValueToPointer[bool](plan.IsWriteOrderConsistent),
	}

	//Update Volume Group by calling API
	_, err := r.client.VolumeGroupApi.PatchVolumeGroupById(ctx, volumeGroupID).Body(volumeGroupUpdate).Execute()
	// r.client.PStoreClient.ModifyVolumeGroup(context.Background(), volumeGroupUpdate, volumeGroupID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating volume group",
			"Could not update volumeGroupID "+volumeGroupID+": "+err.Error(),
		)
	}

	if len(addVolumeIdsSlice) != 0 {
		//Add Volumes in Volume Group by calling API
		_, err := r.allclient.PStoreClient.AddMembersToVolumeGroup(context.Background(), addVolumeGroupMembers, volumeGroupID)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error updating volume group",
				"Could not update volumeGroupID "+volumeGroupID+": "+err.Error(),
			)
		}
	}

	if len(removeVolumeIdsSlice) != 0 {
		//Remove Volumes in Volume Group by calling API
		_, err := r.allclient.PStoreClient.RemoveMembersFromVolumeGroup(context.Background(), removeVolumeGroupMembers, volumeGroupID)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error updating volume group",
				"Could not update volumeGroupID "+volumeGroupID+": "+err.Error(),
			)
		}
	}

	//Get Volume Group details
	getRes, err := r.ReadAPI(context.Background(), volumeGroupID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting volume group after update",
			"Could not get volume group, unexpected error: "+err.Error(),
		)
		return
	}

	r.updateVolGroupState(&state, getRes, &plan)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Successfully done with Update")
}

// ImportState import state for existing volume group
func (r *resourceVolumeGroup) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// updateVolGroupState - method to update terraform state
func (r resourceVolumeGroup) updateVolGroupState(volgroupState *models.Volumegroup, volGroupResponse *clientgen.VolumeGroupInstance, volGroupPlan *models.Volumegroup) {
	// Update value from Volume Group Response to State
	volgroupState.ID = helper.PointerToStringType(volGroupResponse.Id)
	volgroupState.Name = helper.PointerToStringType(volGroupResponse.Name)
	volgroupState.Description = helper.PointerToStringType(volGroupResponse.Description)
	volgroupState.IsWriteOrderConsistent = helper.PointerToBoolType(volGroupResponse.IsWriteOrderConsistent)
	volgroupState.ProtectionPolicyID = helper.PointerToStringType(volGroupResponse.ProtectionPolicyId)

	//Update VolumeIDs value from Response to State
	volgroupState.VolumeIDs, _ = types.SetValue(
		types.StringType,
		helper.SliceTransform(volGroupResponse.Volumes, func(in clientgen.VolumeInstance) attr.Value {
			return helper.PointerToStringType(in.Id)
		}),
	)

	//Update VolumeNames value from Plan to State
	volgroupState.VolumeNames, _ = types.SetValue(
		types.StringType,
		helper.SliceTransform(volGroupPlan.VolumeNames.Elements(), func(in attr.Value) attr.Value {
			return types.StringValue(strings.Trim(in.String(), "\""))
		}),
	)

	//Update ProtectionPolicyName value from Plan to State
	volgroupState.ProtectionPolicyName = volGroupPlan.ProtectionPolicyName
}

// fetchByName fetches by name and updates respective ids in plan/root/terraform-provider-powerstore/examples
func (r *resourceVolumeGroup) fetchByName(plan *models.Volumegroup) string {
	var volumeIds []string
	if len(plan.VolumeNames.Elements()) != 0 {
		for _, volumeName := range plan.VolumeNames.Elements() {
			volume, err := r.allclient.PStoreClient.GetVolumeByName(context.Background(), strings.Trim(volumeName.String(), "\""))
			if err != nil {
				return "Error getting volume with name: " + strings.Trim(volumeName.String(), "\"")
			}
			volumeIds = append(volumeIds, strings.Trim(volume.ID, "\""))
		}
		volumeList := []attr.Value{}
		for i := 0; i < len(volumeIds); i++ {
			volumeList = append(volumeList, types.StringValue(string(volumeIds[i])))
		}
		plan.VolumeIDs, _ = types.SetValue(types.StringType, volumeList)
	}

	if plan.ProtectionPolicyName.ValueString() != "" {
		policy, err := r.allclient.PStoreClient.GetProtectionPolicyByName(context.Background(), plan.ProtectionPolicyName.ValueString())
		if err != nil {
			return "Error getting protection policy with name: " + strings.Trim(plan.ProtectionPolicyName.String(), "\"")
		}
		plan.ProtectionPolicyID = types.StringValue(policy.ID)
	}

	return ""
}
