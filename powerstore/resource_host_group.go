package powerstore

import (
	"context"
	"fmt"
	"log"
	"strings"

	client "terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
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

		MarkdownDescription: "HostGroup resource",

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
				Required:            true,
				Description:         "The list of hosts to include in the host group.",
				MarkdownDescription: "The list of hosts to include in the host group.",
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

	//Delete Host Group by calling API
	_, err := r.client.PStoreClient.DeleteHostGroup(context.Background(), hostGroupID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting host group",
			"Could not delete hostGroupID "+hostGroupID+": "+err.Error(),
		)
		return
	}

	log.Printf("Done with Delete")
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

	// Get ids of hosts to be added/removed from host group
	addHostIds, removeHostIds := GetHostDetails(plan, &state)

	// Get Host Group ID from state
	hostGroupID := state.ID.ValueString()

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
}

// GetHostDetails - Get details of hosts to be added/removed from host group
func GetHostDetails(plan models.HostGroup, state *models.HostGroup) ([]string, []string) {
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
			log.Printf("Volume not found in state")
			removeHostIdsMap[stateHostIds[i]] = stateHostIds[i]
		}
	}

	//Get list of host ids to be removed into a slice
	removeHostIdsSlice := []string{}
	for _, hostID := range removeHostIdsMap {
		removeHostIdsSlice = append(removeHostIdsSlice, hostID)
	}

	//Create map of host ids to be added by comparing plan and state host ids
	addHostIdsMap := make(map[string]string)
	for i := 0; i < len(planHostIds); i++ {
		_, found := stateHostIdsMap[planHostIds[i]]
		if !found {
			log.Printf("Host not found in plan")
			addHostIdsMap[planHostIds[i]] = planHostIds[i]
		}
	}

	//Get list of host ids to be added into a slice
	addHostIdsSlice := []string{}
	for _, hostID := range addHostIdsMap {
		addHostIdsSlice = append(addHostIdsSlice, hostID)
	}

	return addHostIdsSlice, removeHostIdsSlice
}
