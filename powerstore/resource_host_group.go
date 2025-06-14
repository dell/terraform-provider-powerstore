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
	"strings"

	client "terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"

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

// newHostGroupResource returns host group new resource instance
func newHostGroupResource() resource.Resource {
	return &resourceHostGroup{}
}

type resourceHostGroup struct {
	client *client.Client
}

// Metadata defines resource interface Metadata method
func (r *resourceHostGroup) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hostgroup"
}

// Schema defines resource interface Schema method
func (r *resourceHostGroup) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "This resource is used to manage the host group entity of PowerStore Array. We can Create, Update and Delete the host group using this resource. We can also import an existing host group from PowerStore array.",
		Description:         "This resource is used to manage the host group entity of PowerStore Array. We can Create, Update and Delete the host group using this resource. We can also import an existing host group from PowerStore array.",

		Attributes: map[string]schema.Attribute{

			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "Unique identifier of the host group.",
				MarkdownDescription: "Unique identifier of the host group.",
			},

			"name": schema.StringAttribute{
				Required:            true,
				Description:         "The host group name.",
				MarkdownDescription: "The host group name.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "An optional description for the host group.",
				MarkdownDescription: "An optional description for the host group.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"host_ids": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "The list of host IDs to include in the host group. Conflicts with `host_names`.",
				MarkdownDescription: "The list of host IDs to include in the host group. Conflicts with `host_names`.",
				Validators: []validator.Set{
					setvalidator.ConflictsWith(path.Expressions{
						path.MatchRoot("host_names"),
					}...),
				},
			},

			"host_names": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "The list of host names to include in the host group. Conflicts with `host_ids`.",
				MarkdownDescription: "The list of host names to include in the host group. Conflicts with `host_ids`.",
			},

			"host_connectivity": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Connectivity type for hosts and host groups.",
				MarkdownDescription: "Connectivity type for hosts and host groups.",
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						string(gopowerstore.HostConnectivityEnumLocalOnly),
						string(gopowerstore.HostConnectivityEnumMetroOptimizeBoth),
						string(gopowerstore.HostConnectivityEnumMetroOptimizeLocal),
						string(gopowerstore.HostConnectivityEnumMetroOptimizeRemote),
					}...),
				},
			},
		},
	}
}

// Configure - defines configuration for host group resource
func (r *resourceHostGroup) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.client = client
}

// Create - method to create host group resource
func (r *resourceHostGroup) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan models.HostGroup

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	errmsg := r.fetchByName(&plan)
	if errmsg != "" {
		resp.Diagnostics.AddError(
			"Error creating host group",
			"Could not create host group, unexpected error: "+errmsg,
		)
		return
	}

	if !plan.HostConnectivity.IsUnknown() {
		resp.Diagnostics.AddError(
			"Error creating host group",
			"Could not set host_connectivity while creating host group",
		)
		return
	}

	hostGroupCreate := r.planToHostGroupParam(plan)

	//Create New HostGroup
	hostGroupCreateResponse, err := r.client.PStoreClient.CreateHostGroup(context.Background(), hostGroupCreate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating host group",
			"Could not create host group, unexpected error: "+err.Error(),
		)
		return
	}

	//Get Host Group details using ID retrived above
	hostGroupResponse, err := r.client.PStoreClient.GetHostGroup(context.Background(), hostGroupCreateResponse.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting host group after creation",
			"Could not get host group, unexpected error: "+err.Error(),
		)
		return
	}

	result := models.HostGroup{}
	r.updateHostGroupState(&result, hostGroupResponse, &plan)

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Create")
}

// Delete - method to delete host group resource
func (r *resourceHostGroup) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("Started with the Delete")

	var state models.HostGroup
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Get Host Group ID from state
	hostGroupID := state.ID.ValueString()

	// Get ids of hosts to be added/removed from host group
	removeHostIds := GetHostIDsFromState(&state)

	hostGroupUpdate := &gopowerstore.HostGroupModify{
		Name:             state.Name.ValueString(),
		Description:      state.Description.ValueString(),
		HostConnectivity: state.HostConnectivity.ValueString(),
		RemoveHostIDs:    removeHostIds,
		AddHostIDs:       []string{},
	}

	//Update Host Group by calling API
	_, err := r.client.PStoreClient.ModifyHostGroup(context.Background(), hostGroupUpdate, hostGroupID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating host group",
			"Could not update hostGroupID "+hostGroupID+": "+err.Error(),
		)
	}

	//Delete Host Group by calling API
	_, err = r.client.PStoreClient.DeleteHostGroup(context.Background(), hostGroupID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting host group",
			"Could not delete hostGroupID "+hostGroupID+": "+err.Error(),
		)
		return
	}

	log.Printf("Done with Delete")
}

// GetHostIDsFromState - get host ids from state
func GetHostIDsFromState(state *models.HostGroup) []string {

	//Get host ids from state into a slice
	var stateHostIds []string
	for _, host := range state.HostIDs.Elements() {
		stateHostIds = append(stateHostIds, strings.Trim(host.String(), "\""))
	}

	return stateHostIds
}

// Read - reads host group resource information
func (r *resourceHostGroup) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	log.Printf("Reading Host Group")
	var state models.HostGroup
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Get host group details from API and update what is in state from what the API returns
	id := state.ID.ValueString()
	response, err := r.client.PStoreClient.GetHostGroup(context.Background(), id)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading host group",
			"Could not read host group with error "+id+": "+err.Error(),
		)
		return
	}

	r.updateHostGroupState(&state, response, &state)

	//Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Read")
}

// Update - updates host group resource
func (r *resourceHostGroup) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("Started Update")

	//Get plan values
	var plan models.HostGroup
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Get current state
	var state models.HostGroup
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	errmsg := r.fetchByName(&plan)
	if errmsg != "" {
		resp.Diagnostics.AddError(
			"Error updating host group",
			"Could not update host group, unexpected error: "+errmsg,
		)
		return
	}

	// Get Host Group ID from state
	hostGroupID := state.ID.ValueString()

	// Get ids of hosts to be added/removed from host group
	addHostIds, removeHostIds, errmsg := GetHostDetails(plan, &state)
	if errmsg != "" {
		resp.Diagnostics.AddError(
			"Error updating host group",
			"Could not update hostGroupID "+hostGroupID+": "+errmsg,
		)
		return
	}

	hostGroupUpdate := &gopowerstore.HostGroupModify{
		Name:             plan.Name.ValueString(),
		Description:      plan.Description.ValueString(),
		HostConnectivity: plan.HostConnectivity.ValueString(),
		RemoveHostIDs:    removeHostIds,
		AddHostIDs:       addHostIds,
	}

	//Update Host Group by calling API
	_, err := r.client.PStoreClient.ModifyHostGroup(context.Background(), hostGroupUpdate, hostGroupID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating host group",
			"Could not update hostGroupID "+hostGroupID+": "+err.Error(),
		)
	}

	//Get Host Group details
	getRes, err := r.client.PStoreClient.GetHostGroup(context.Background(), hostGroupID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting host group after update",
			"Could not get host group, unexpected error: "+err.Error(),
		)
		return
	}

	r.updateHostGroupState(&state, getRes, &plan)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Successfully done with Update")
}

// ImportState import state for existing host group
func (r *resourceHostGroup) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// planToHostGroupParam - Create HostGroupCreate instance
func (r resourceHostGroup) planToHostGroupParam(plan models.HostGroup) *gopowerstore.HostGroupCreate {

	var hostIds []string
	for _, hostID := range plan.HostIDs.Elements() {
		hostIds = append(hostIds, strings.Trim(hostID.String(), "\""))
	}

	hostGroupCreate := &gopowerstore.HostGroupCreate{
		Name:        plan.Name.ValueString(),
		Description: plan.Description.ValueString(),
		HostIDs:     hostIds,
	}
	return hostGroupCreate
}

// updateHostGroupState - Update host group details from response to state
func (r resourceHostGroup) updateHostGroupState(hostGroupState *models.HostGroup, hostGroupResponse gopowerstore.HostGroup, hostGroupPlan *models.HostGroup) {
	// Update value from Host Group Response to State
	hostGroupState.ID = types.StringValue(hostGroupResponse.ID)
	hostGroupState.Name = types.StringValue(hostGroupResponse.Name)
	hostGroupState.Description = types.StringValue(hostGroupResponse.Description)
	hostGroupState.HostConnectivity = types.StringValue(string(hostGroupResponse.HostConnectivity))

	//Update HostIDs value from Response to State
	var hostIDs []string
	for _, host := range hostGroupResponse.Hosts {
		hostIDs = append(hostIDs, host.ID)
	}

	hostIDList := []attr.Value{}
	for _, hostID := range hostIDs {
		hostIDList = append(hostIDList, types.StringValue(string(hostID)))
	}
	hostGroupState.HostIDs, _ = types.SetValue(types.StringType, hostIDList)

	//Update HostNames value from Plan to State
	var hostNames []string
	for _, hostName := range hostGroupPlan.HostNames.Elements() {
		hostNames = append(hostNames, strings.Trim(hostName.String(), "\""))
	}
	hostNameList := []attr.Value{}
	for _, hostName := range hostNames {
		hostNameList = append(hostNameList, types.StringValue(string(hostName)))
	}
	hostGroupState.HostNames, _ = types.SetValue(types.StringType, hostNameList)

}

// GetHostDetails - Get details of hosts to be added/removed from host group
func GetHostDetails(plan models.HostGroup, state *models.HostGroup) ([]string, []string, string) {
	//Get host ids from plan into a slice
	var planHostIds []string
	for _, host := range plan.HostIDs.Elements() {
		planHostIds = append(planHostIds, strings.Trim(host.String(), "\""))
	}

	//Get host ids from state into a slice
	var stateHostIds []string
	for _, host := range state.HostIDs.Elements() {
		stateHostIds = append(stateHostIds, strings.Trim(host.String(), "\""))
	}

	//Get host ids from plan into map for optimized element search
	planHostIdsMap := make(map[string]string)
	if len(planHostIds) != 0 {
		for i := 0; i < len(planHostIds); i++ {
			planHostIdsMap[planHostIds[i]] = planHostIds[i]
		}
	}

	//Get host ids from state into a map for optimized element search
	stateHostIdsMap := make(map[string]string)
	if len(stateHostIds) != 0 {
		for i := 0; i < len(stateHostIds); i++ {
			stateHostIdsMap[stateHostIds[i]] = stateHostIds[i]
		}
	}

	//Create map of host ids to be removed by comparing plan and state host ids
	removeHostIdsMap := make(map[string]string)
	for i := 0; i < len(stateHostIds); i++ {
		_, found := planHostIdsMap[stateHostIds[i]]
		if !found {
			removeHostIdsMap[stateHostIds[i]] = stateHostIds[i]
		}
	}

	//Get list of host ids to be removed into a slice
	removeHostIdsSlice := []string{}
	for _, hostID := range removeHostIdsMap {
		removeHostIdsSlice = append(removeHostIdsSlice, hostID)
	}

	errmsg := ""
	// check to disallow removing all the hosts from the host group when host connectivity is Local_Only
	if plan.HostConnectivity.ValueString() == "Local_Only" && planHostIds == nil && removeHostIdsSlice != nil {
		errmsg = "Cannot remove all the hosts when host_connectivity is set"
	}

	//Create map of host ids to be added by comparing plan and state host ids
	addHostIdsMap := make(map[string]string)
	for i := 0; i < len(planHostIds); i++ {
		_, found := stateHostIdsMap[planHostIds[i]]
		if !found {
			addHostIdsMap[planHostIds[i]] = planHostIds[i]
		}
	}

	//Get list of host ids to be added into a slice
	addHostIdsSlice := []string{}
	for _, hostID := range addHostIdsMap {
		addHostIdsSlice = append(addHostIdsSlice, hostID)
	}

	return addHostIdsSlice, removeHostIdsSlice, errmsg
}

// fetchByName fetches hosts using name and updates respective ids in plan
func (r resourceHostGroup) fetchByName(plan *models.HostGroup) string {
	var hostIds []string
	if len(plan.HostNames.Elements()) != 0 {
		for _, hostName := range plan.HostNames.Elements() {
			host, err := r.client.PStoreClient.GetHostByName(context.Background(), strings.Trim(hostName.String(), "\""))
			if err != nil {
				return "Error getting host with name: " + strings.Trim(hostName.String(), "\"")
			}
			hostIds = append(hostIds, strings.Trim(host.ID, "\""))
		}
		hostList := []attr.Value{}
		for i := 0; i < len(hostIds); i++ {
			hostList = append(hostList, types.StringValue(string(hostIds[i])))
		}
		plan.HostIDs, _ = types.SetValue(types.StringType, hostList)
	}

	return ""
}
