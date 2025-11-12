/*
Copyright (c) 2024-2025 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"
	"terraform-provider-powerstore/powerstore/helper"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// newHostResource returns host new resource instance
func newHostResource() resource.Resource {
	return &resourceHost{}
}

type resourceHost struct {
	client *client.Client
}

// Metadata defines resource interface Metadata method
func (r *resourceHost) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_host"
}

// Schema defines resource interface Schema method
func (r *resourceHost) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "This resource is used to manage the host entity of PowerStore Array. We can Create, Update and Delete the host using this resource. We can also import an existing host from PowerStore array.",
		Description:         "This resource is used to manage the host entity of PowerStore Array. We can Create, Update and Delete the host using this resource. We can also import an existing host from PowerStore array.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The unique identifier of the host.",
				MarkdownDescription: "The unique identifier of the host.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Name of the host. This should be unique across all hosts in the cluster.",
				MarkdownDescription: "Name of the host. This should be unique across all hosts in the cluster.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Description of the host.",
				MarkdownDescription: "Description of the host.",
			},
			"host_group_id": schema.StringAttribute{
				Computed:            true,
				Description:         "Associated host group, if host is part of host group.",
				MarkdownDescription: "Associated host group, if host is part of host group.",
			},
			"os_type": schema.StringAttribute{
				Required:            true,
				Description:         "Operating system of the host. This cannot be updated",
				MarkdownDescription: "Operating system of the host. This cannot be updated",
				Validators: []validator.String{stringvalidator.OneOf(
					string(gopowerstore.OSTypeEnumWindows),
					string(gopowerstore.OSTypeEnumLinux),
					string(gopowerstore.OSTypeEnumESXi),
				)},
			},
			"initiators": schema.SetNestedAttribute{
				Description:         "Parameters for creating or adding initiators to host.",
				MarkdownDescription: "Parameters for creating or adding initiators to host.",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"port_name": schema.StringAttribute{
							Description:         "IQN name aka address or NQN name for NVMEoF port types.",
							MarkdownDescription: "IQN name aka address or NQN name for NVMEoF port types.",
							Required:            true,
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(1),
							},
						},
						"port_type": schema.StringAttribute{
							Description:         "Protocol type of the host initiator.",
							MarkdownDescription: "Protocol type of the host initiator.",
							Computed:            true,
						},
						"chap_mutual_password": schema.StringAttribute{
							Description:         "Password for CHAP authentication. This value must be 12 to 64 UTF-8 characters. This password cannot be queried. CHAP password is required when the cluster CHAP mode is mutual authentication.",
							MarkdownDescription: "Password for CHAP authentication. This value must be 12 to 64 UTF-8 characters. This password cannot be queried. CHAP password is required when the cluster CHAP mode is mutual authentication.",
							Optional:            true,
							Sensitive:           true,
						},
						"chap_mutual_username": schema.StringAttribute{
							Description:         "Username for CHAP authentication. This value must be 1 to 64 UTF-8 characters. CHAP username is required when the cluster CHAP mode is mutual authentication.",
							MarkdownDescription: "Username for CHAP authentication. This value must be 1 to 64 UTF-8 characters. CHAP username is required when the cluster CHAP mode is mutual authentication.",
							Optional:            true,
						},
						"chap_single_username": schema.StringAttribute{
							Description:         "Username for CHAP authentication. This value must be 1 to 64 UTF-8 characters. CHAP username is required when the cluster CHAP mode is single authentication.",
							MarkdownDescription: "Username for CHAP authentication. This value must be 1 to 64 UTF-8 characters. CHAP username is required when the cluster CHAP mode is single authentication.",
							Optional:            true,
						},
						"chap_single_password": schema.StringAttribute{
							Description:         "Password for CHAP authentication. This value must be 12 to 64 UTF-8 characters. This password cannot be queried. CHAP password is required when the cluster CHAP mode is single authentication.",
							MarkdownDescription: "Password for CHAP authentication. This value must be 12 to 64 UTF-8 characters. This password cannot be queried. CHAP password is required when the cluster CHAP mode is single authentication.",
							Optional:            true,
							Sensitive:           true,
						},
					},
				},
			},
			"host_connectivity": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Connectivity type for hosts and host groups.",
				MarkdownDescription: "Connectivity type for hosts and host groups.",
				Validators: []validator.String{stringvalidator.OneOf(
					string(gopowerstore.HostConnectivityEnumLocalOnly),
					string(gopowerstore.HostConnectivityEnumMetroOptimizeBoth),
					string(gopowerstore.HostConnectivityEnumMetroOptimizeLocal),
					string(gopowerstore.HostConnectivityEnumMetroOptimizeRemote),
				)},
			},
		},
	}
}

// Configure - defines configuration for host resource
func (r *resourceHost) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// Create - method to create host resource
func (r *resourceHost) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan models.Host

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if initiators are known before processing
	if plan.Initiators.IsUnknown() || plan.Initiators.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid initiators",
			"The 'initiators' attribute is unknown or null at plan time. This may happen when using dynamic blocks or for_each. Please ensure the value is known before applying.",
		)
		return
	}

	// traverse through initiators in plan and store them
	var initiators []gopowerstore.InitiatorCreateModify
	for _, v := range plan.Initiators.Elements() {
		objVal := v.(types.Object)
		initiatorModel := models.InitiatorCreateModify{}
		diags := objVal.As(ctx, &initiatorModel, basetypes.ObjectAsOptions{})
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		initiator := r.addInitiatorFromPlan(initiatorModel)
		initiators = append(initiators, initiator)
	}

	name := plan.Name.ValueString()
	description := plan.Description.ValueString()
	osType := gopowerstore.OSTypeEnum(plan.OsType.ValueString())

	hostCreate := &gopowerstore.HostCreate{
		Name:             &name,
		Description:      &description,
		OsType:           &osType,
		Initiators:       &initiators,
		HostConnectivity: gopowerstore.HostConnectivityEnum(plan.HostConnectivity.ValueString()),
	}

	// Create new host
	hostCreateResponse, err := r.client.PStoreClient.CreateHost(context.Background(), hostCreate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating host",
			"Could not create host, unexpected error: "+err.Error(),
		)
		return
	}

	// Get host Details using ID retrieved above
	hostResponse, err1 := r.client.PStoreClient.GetHost(context.Background(), hostCreateResponse.ID)
	if err1 != nil {
		resp.Diagnostics.AddError(
			"Error getting host after creation",
			"Could not get host, unexpected error: "+err.Error(),
		)
		return
	}

	// Update details to state
	result := models.Host{}

	r.serverToState(ctx, &plan, &result, hostResponse, operationCreate)

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Create")
}

// Read - reads host resource information
func (r *resourceHost) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state models.Host
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get host details from API and then update what is in state from what the API returns
	hostID := state.ID.ValueString()
	hostResponse, err := r.client.PStoreClient.GetHost(context.Background(), hostID)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading host",
			"Could not read hostID with error "+hostID+": "+err.Error(),
		)
		return
	}

	r.serverToState(ctx, nil, &state, hostResponse, operationRead)

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Done with Read")
}

// Update - updates host resource
func (r *resourceHost) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("Started Update")

	// Get plan values
	var plan models.Host
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state models.Host
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get host ID from state
	hostID := state.ID.ValueString()

	// Get host Details
	_, err := r.client.PStoreClient.GetHost(context.Background(), hostID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting host after update",
			"Could not get host after update, unexpected error: "+err.Error(),
		)
		return
	}

	// Update host by calling API
	// Since either update/add/remove can be performed in a single call so moved modify and remove separately.
	// first check if there is any addition in initiators
	_, err = r.client.PStoreClient.ModifyHost(
		context.Background(),
		r.addInitiators(ctx, plan, state),
		hostID,
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating host",
			"Could not update hostID "+hostID+": "+err.Error(),
		)
	}

	// Check if there is any removal in initiators
	_, err = r.client.PStoreClient.ModifyHost(
		context.Background(),
		r.removeInitiators(ctx, plan, state),
		hostID,
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating host",
			"Could not update hostID "+hostID+": "+err.Error(),
		)
	}

	// Modify CHAP credentials based on PortName.
	modifyInitiators := r.modifyOperation(ctx, plan, state)
	if len(modifyInitiators) > 0 {
		hostUpdate := &gopowerstore.HostModify{
			ModifyInitiators: &modifyInitiators,
		}
		_, err = r.client.PStoreClient.ModifyHost(
			context.Background(),
			hostUpdate,
			hostID,
		)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error updating host",
				"Could not update hostID "+hostID+": "+err.Error(),
			)
		}
	}

	// Get host Details
	hostResponse, err := r.client.PStoreClient.GetHost(context.Background(), hostID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting host after update",
			"Could not get host after update, unexpected error: "+err.Error(),
		)
		return
	}

	// Update the data to state
	r.serverToState(ctx, &plan, &state, hostResponse, operationUpdate)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Successfully done with Update")
}

// Delete - method to delete host resource
func (r *resourceHost) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("Started with Delete")

	var state models.Host
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get host ID from state
	hostID := state.ID.ValueString()

	// Delete host by calling API
	_, err := r.client.PStoreClient.DeleteHost(context.Background(), &gopowerstore.HostDelete{}, hostID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting host",
			"Could not delete hostID "+hostID+": "+err.Error(),
		)
		return
	}

	log.Printf("Done with Delete")
}

// ImportState - imports state for existing host
func (r *resourceHost) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Update details to state from the API repsonse
func (r resourceHost) serverToState(ctx context.Context, plan, state *models.Host, response gopowerstore.Host, operation operation) {

	if response.ID != "" {
		state.ID = types.StringValue(response.ID)
	}
	if response.Name != "" {
		state.Name = types.StringValue(response.Name)
	}
	state.Description = types.StringValue(response.Description)
	if response.OsType != "" {
		state.OsType = types.StringValue(string(response.OsType))
	}
	if response.HostConnectivity != "" {
		state.HostConnectivity = types.StringValue(string(response.HostConnectivity))
	}
	initiators := make([]models.InitiatorCreateModify, 0, len(response.Initiators))

	// fetch the plan data to get password value.
	// Passwords are not queryable so in order to maintain state it is taken from the plan.
	// Convert plan.Initiators to map
	planInitiatorMap := make(map[string]models.InitiatorCreateModify)
	if plan != nil && !plan.Initiators.IsNull() && !plan.Initiators.IsUnknown() {
		for _, elem := range plan.Initiators.Elements() {
			objVal := elem.(types.Object)
			var model models.InitiatorCreateModify
			diags := objVal.As(ctx, &model, basetypes.ObjectAsOptions{})
			if diags.HasError() {
				// optionally log or append diagnostics
				continue
			}
			portName := model.PortName.ValueString()
			planInitiatorMap[portName] = model
		}
	}

	// fetch the plan data to get password value.
	// Passwords are not queryable so in order to maintain state it is taken from the plan.
	// Convert state.Initiators to map
	stateInitiatorMap := make(map[string]models.InitiatorCreateModify)
	if state != nil && !state.Initiators.IsNull() && !state.Initiators.IsUnknown() {
		for _, elem := range state.Initiators.Elements() {
			objVal := elem.(types.Object)
			var model models.InitiatorCreateModify
			diags := objVal.As(ctx, &model, basetypes.ObjectAsOptions{})
			if diags.HasError() {
				// optionally log or append diagnostics
				continue
			}
			portName := model.PortName.ValueString()
			stateInitiatorMap[portName] = model
		}
	}

	for _, initiator := range response.Initiators {
		initiatorModel := models.InitiatorCreateModify{}
		if initiator.PortType != "iSCSI" || (initiator.ChapMutualUsername == "" && initiator.ChapSingleUsername == "") {
			initiatorModel.PortName = types.StringValue(initiator.PortName)
			initiatorModel.PortType = types.StringValue(string(initiator.PortType))
		} else if initiator.ChapMutualUsername == "" {
			initiatorModel.PortName = types.StringValue(initiator.PortName)
			initiatorModel.PortType = types.StringValue(string(initiator.PortType))
			if operation != operationRead {
				initiatorModel.ChapSinglePassword = planInitiatorMap[initiator.PortName].ChapSinglePassword
			} else {
				initiatorModel.ChapSinglePassword = stateInitiatorMap[initiator.PortName].ChapSinglePassword
			}
			initiatorModel.ChapSingleUsername = types.StringValue(initiator.ChapSingleUsername)
		} else {
			initiatorModel.PortName = types.StringValue(initiator.PortName)
			initiatorModel.PortType = types.StringValue(string(initiator.PortType))
			if operation != operationRead {
				initiatorModel.ChapSinglePassword = planInitiatorMap[initiator.PortName].ChapSinglePassword
				initiatorModel.ChapMutualPassword = planInitiatorMap[initiator.PortName].ChapMutualPassword
			} else {
				initiatorModel.ChapSinglePassword = stateInitiatorMap[initiator.PortName].ChapSinglePassword
				initiatorModel.ChapMutualPassword = stateInitiatorMap[initiator.PortName].ChapMutualPassword
			}
			initiatorModel.ChapMutualUsername = types.StringValue(initiator.ChapMutualUsername)
			initiatorModel.ChapSingleUsername = types.StringValue(initiator.ChapSingleUsername)
		}

		initiators = append(initiators, initiatorModel)
	}
	initiatorAttrTypes := map[string]attr.Type{
		"port_name":            types.StringType,
		"port_type":            types.StringType,
		"chap_mutual_password": types.StringType,
		"chap_mutual_username": types.StringType,
		"chap_single_password": types.StringType,
		"chap_single_username": types.StringType,
	}

	initiatorObjects := make([]attr.Value, 0, len(initiators))
	for _, i := range initiators {
		obj, _ := types.ObjectValueFrom(ctx, initiatorAttrTypes, i)
		initiatorObjects = append(initiatorObjects, obj)
	}

	setVal, _ := types.SetValue(types.ObjectType{AttrTypes: initiatorAttrTypes}, initiatorObjects)
	state.Initiators = setVal

	if operation == operationRead {
		state.HostGroupID = types.StringValue(response.HostGroupID)
	}
}

// Attributes to be updated in update operation
func (r resourceHost) addInitiators(ctx context.Context, plan, state models.Host) *gopowerstore.HostModify {

	hostUpdate := &gopowerstore.HostModify{}
	name := plan.Name.ValueString()
	description := plan.Description.ValueString()

	if plan.HostConnectivity.ValueString() != state.HostConnectivity.ValueString() {
		hostUpdate.HostConnectivity = gopowerstore.HostConnectivityEnum(plan.HostConnectivity.ValueString())
	}

	// Create a map of initiators from state with PortName as key, as it is unique
	stateInitiatorsMap := r.getInitiatorMap(ctx, state.Initiators)

	// Create map of initiators to be added
	addInitiatorsMap := make(map[types.String]models.InitiatorCreateModify)
	for _, elem := range plan.Initiators.Elements() {
		objVal := elem.(types.Object)
		var model models.InitiatorCreateModify
		_ = objVal.As(ctx, &model, basetypes.ObjectAsOptions{})
		_, found := stateInitiatorsMap[model.PortName.ValueString()]
		if !found {
			addInitiatorsMap[model.PortName] = model
		}
	}

	addInitiators := make([]gopowerstore.InitiatorCreateModify, 0, len(addInitiatorsMap))
	for _, initiator := range addInitiatorsMap {
		currentInitiator := r.addInitiatorFromPlan(initiator)
		addInitiators = append(addInitiators, currentInitiator)
	}

	hostUpdate = &gopowerstore.HostModify{
		Description:      &description,
		HostConnectivity: gopowerstore.HostConnectivityEnum(plan.HostConnectivity.ValueString()),
		Name:             &name,
		AddInitiators:    &addInitiators,
	}
	return hostUpdate
}

// Attributes to be updated in update operation
func (r resourceHost) removeInitiators(ctx context.Context, plan, state models.Host) *gopowerstore.HostModify {

	// Create a map of initiators from plan with PortName as key, as it is unique
	planInitiatorsMap := r.getInitiatorMap(ctx, plan.Initiators)

	// create a map to find initiators to be removed
	removeInitiatorsMap := make(map[types.String]models.InitiatorCreateModify)
	for _, elem := range state.Initiators.Elements() {
		objVal := elem.(types.Object)
		var model models.InitiatorCreateModify
		_ = objVal.As(ctx, &model, basetypes.ObjectAsOptions{})
		_, found := planInitiatorsMap[model.PortName.ValueString()]
		if !found {
			removeInitiatorsMap[model.PortName] = model
		}
	}

	// Fetch keys (port names) to be removed
	var removeInitiators []string
	for removeID := range removeInitiatorsMap {
		removeInitiators = append(removeInitiators, removeID.ValueString())
	}

	hostUpdate := &gopowerstore.HostModify{
		RemoveInitiators: &removeInitiators,
	}

	return hostUpdate
}

// to perform modify operation in update
func (r resourceHost) modifyOperation(ctx context.Context, plan, state models.Host) []gopowerstore.UpdateInitiatorInHost {
	// update CHAP credentials based on port name
	modifyInitiators := make([]gopowerstore.UpdateInitiatorInHost, 0, len(plan.Initiators.Elements()))

	stateInitiatorMap := r.getInitiatorMap(ctx, state.Initiators)
	for _, elem := range plan.Initiators.Elements() {
		objVal := elem.(types.Object)
		var initiator models.InitiatorCreateModify
		_ = objVal.As(ctx, &initiator, basetypes.ObjectAsOptions{})
		if _, ok := stateInitiatorMap[initiator.PortName.ValueString()]; ok {

			var updateInitiator gopowerstore.UpdateInitiatorInHost

			portName := initiator.PortName.ValueString()
			portType := r.getPortType(portName)
			chapMutualPassword := initiator.ChapMutualPassword.ValueString()
			chapMutualUsername := initiator.ChapMutualUsername.ValueString()
			chapSinglePassword := initiator.ChapSinglePassword.ValueString()
			chapSingleUsername := initiator.ChapSingleUsername.ValueString()

			var chapSingleUsername1 *string = &chapSingleUsername
			var chapSinglePassword1 *string = &chapSinglePassword
			var chapMutualUsername1 *string = &chapMutualUsername
			var chapMutualPassword1 *string = &chapMutualPassword

			// update is supported only for CHAP credentials and they are absent in case of `NVME` and `FC` so modification is not applicable for them
			if portType == string(gopowerstore.InitiatorProtocolTypeEnumISCSI) {
				if initiator.ChapMutualUsername.IsNull() && initiator.ChapSingleUsername.IsNull() {
					updateInitiator = gopowerstore.UpdateInitiatorInHost{
						PortName: &portName,
					}
				} else if initiator.ChapMutualUsername.IsNull() {
					updateInitiator = gopowerstore.UpdateInitiatorInHost{
						PortName:           &portName,
						ChapSinglePassword: chapSinglePassword1,
						ChapSingleUsername: chapSingleUsername1,
					}
				} else if !initiator.ChapMutualUsername.IsNull() {
					updateInitiator = gopowerstore.UpdateInitiatorInHost{
						PortName:           &portName,
						ChapMutualPassword: chapMutualPassword1,
						ChapMutualUsername: chapMutualUsername1,
						ChapSinglePassword: chapSinglePassword1,
						ChapSingleUsername: chapSingleUsername1,
					}
				}
				modifyInitiators = append(modifyInitiators, updateInitiator)
			}
		}
	}
	return modifyInitiators
}

func (r *resourceHost) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.Host
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if data.Initiators.IsNull() || data.Initiators.IsUnknown() {
		return
	}

	for _, elem := range data.Initiators.Elements() {
		objVal := elem.(types.Object)
		var initiator models.InitiatorCreateModify
		diags := objVal.As(ctx, &initiator, basetypes.ObjectAsOptions{})
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		// PortName is required. It can be unknown, but not null. Ignore if unknown.
		if !initiator.PortName.IsUnknown() &&
			// if port type is not iSCSI then check further
			r.getPortType(initiator.PortName.ValueString()) != string(gopowerstore.InitiatorProtocolTypeEnumISCSI) &&
			// check if any of the chap creds have not been configured
			(helper.IsKnownValue(initiator.ChapSingleUsername) ||
				helper.IsKnownValue(initiator.ChapMutualUsername) ||
				helper.IsKnownValue(initiator.ChapMutualPassword) ||
				helper.IsKnownValue(initiator.ChapSinglePassword)) {
			resp.Diagnostics.AddError(
				"Error validating config host",
				"chap credentials are supported only with iSCSI protocol",
			)
		}
		if helper.IsKnownValue(initiator.ChapMutualUsername) && initiator.ChapSingleUsername.IsNull() {
			resp.Diagnostics.AddError(
				"Error validating config host",
				"`chap_mutual_username` cannot be present without `chap_single_username`",
			)
		}
		if helper.IsKnownValue(initiator.ChapMutualPassword) && initiator.ChapSinglePassword.IsNull() {
			resp.Diagnostics.AddError(
				"Error validating config host",
				"`chap_mutual_password` cannot be present without `chap_mutual_username`",
			)
		}
		if helper.IsKnownValue(initiator.ChapSinglePassword) && initiator.ChapSingleUsername.IsNull() {
			resp.Diagnostics.AddError(
				"Error validating config host",
				"`chap_single_password` cannot be present without `chap_single_username`",
			)
		}
	}
}

func (r resourceHost) getPortType(portName string) string {
	var portType string
	if strings.HasPrefix(portName, "iqn") {
		portType = string(gopowerstore.InitiatorProtocolTypeEnumISCSI)
	} else if strings.HasPrefix(portName, "nqn") {
		portType = string(gopowerstore.InitiatorProtocolTypeEnumNVME)
	} else {
		portType = string(gopowerstore.InitiatorProtocolTypeEnumFC)
	}
	return portType
}

func (r resourceHost) getInitiatorMap(ctx context.Context, set types.Set) map[string]models.InitiatorCreateModify {
	initiatorsMap := make(map[string]models.InitiatorCreateModify)

	if !set.IsNull() && !set.IsUnknown() {
		for _, elem := range set.Elements() {
			objVal := elem.(types.Object)
			var model models.InitiatorCreateModify
			_ = objVal.As(ctx, &model, basetypes.ObjectAsOptions{})
			initiatorsMap[model.PortName.ValueString()] = model
		}
	}
	return initiatorsMap
}

func (r resourceHost) addInitiatorFromPlan(v models.InitiatorCreateModify) gopowerstore.InitiatorCreateModify {

	initiator := gopowerstore.InitiatorCreateModify{}

	portName := v.PortName.ValueString()
	portType := r.getPortType(portName)
	chapMutualPassword := v.ChapMutualPassword.ValueString()
	chapMutualUsername := v.ChapMutualUsername.ValueString()
	chapSinglePassword := v.ChapSinglePassword.ValueString()
	chapSingleUsername := v.ChapSingleUsername.ValueString()

	// When port type is iSCSI only then look for CHAP Username and Password
	if portType != "iSCSI" || (chapMutualUsername == "" && chapSingleUsername == "") {
		initiator.PortName = &portName
		initiator.PortType = (*gopowerstore.InitiatorProtocolTypeEnum)(&portType)
	} else if chapMutualUsername == "" {
		initiator.PortName = &portName
		initiator.PortType = (*gopowerstore.InitiatorProtocolTypeEnum)(&portType)
		initiator.ChapSinglePassword = &chapSinglePassword
		initiator.ChapSingleUsername = &chapSingleUsername
	} else {
		initiator.PortName = &portName
		initiator.PortType = (*gopowerstore.InitiatorProtocolTypeEnum)(&portType)
		initiator.ChapMutualPassword = &chapMutualPassword
		initiator.ChapMutualUsername = &chapMutualUsername
		initiator.ChapSinglePassword = &chapSinglePassword
		initiator.ChapSingleUsername = &chapSingleUsername
	}
	return initiator
}
