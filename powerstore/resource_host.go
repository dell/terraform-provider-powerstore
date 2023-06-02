package powerstore

import (
	"context"
	"fmt"
	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"log"
	"strings"
	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"
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

		MarkdownDescription: "Host resource",

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

	// traverse through initiators in plan and store them
	var initiators []gopowerstore.InitiatorCreateModify
	for _, v := range plan.Initiators {
		initiator := r.addInitiatorFromPlan(v)
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

	r.serverToState(&plan, &result, hostResponse, operationCreate)

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

	r.serverToState(nil, &state, hostResponse, operationRead)

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
	hostResponse, err := r.client.PStoreClient.GetHost(context.Background(), hostID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting host after update",
			"Could not get host after update, unexpected error: "+err.Error(),
		)
		return
	}

	// Update host by calling API
	//  Since either update/add/remove can be performed in a single call so moved modify and remove separately.
	// first check if there is any addition in initiators
	_, err = r.client.PStoreClient.ModifyHost(
		context.Background(),
		r.addInitiators(plan, state),
		hostID,
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating host",
			"Could not update hostID "+hostID+": "+err.Error(),
		)
	}

	// Check for removal of initiators
	_, err = r.client.PStoreClient.ModifyHost(
		context.Background(),
		r.removeInitiators(plan, state),
		hostID,
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating host",
			"Could not update hostID "+hostID+": "+err.Error(),
		)
	}

	// Modify CHAP credentials based on PortName.
	// since due to idempotency issue modify is getting called in every call.
	_, err = r.client.PStoreClient.ModifyHost(
		context.Background(),
		r.modifyOperation(plan, state),
		hostID,
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating host",
			"Could not update hostID "+hostID+": "+err.Error(),
		)
	}

	// Get host Details
	hostResponse, err = r.client.PStoreClient.GetHost(context.Background(), hostID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting host after update",
			"Could not get host after update, unexpected error: "+err.Error(),
		)
		return
	}

	// Update the data to state
	r.serverToState(&plan, &state, hostResponse, operationUpdate)

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
func (r resourceHost) serverToState(plan, state *models.Host, response gopowerstore.Host, operation operation) {
	state.ID = types.StringValue(response.ID)
	state.Name = types.StringValue(response.Name)
	state.Description = types.StringValue(response.Description)
	state.OsType = types.StringValue(string(response.OsType))
	state.HostConnectivity = types.StringValue(string(response.HostConnectivity))
	initiators := make([]models.InitiatorCreateModify, 0, len(response.Initiators))

	// fetch the plan data to get password value.
	// Passwords are not queryable so in order to maintain state it is taken from the plan.
	planInitiatorMap := make(map[types.String]models.InitiatorCreateModify)
	if plan != nil && len(plan.Initiators) != 0 {
		for i := 0; i < len(plan.Initiators); i++ {
			planInitiatorMap[plan.Initiators[i].PortName] = plan.Initiators[i]
		}
	}

	// fetch the plan data to get password value.
	// Passwords are not queryable so in order to maintain state it is taken from the plan.
	stateInitiatorMap := make(map[types.String]models.InitiatorCreateModify)
	if state != nil && len(state.Initiators) != 0 {
		for i := 0; i < len(state.Initiators); i++ {
			stateInitiatorMap[state.Initiators[i].PortName] = state.Initiators[i]
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
				initiatorModel.ChapSinglePassword = planInitiatorMap[types.StringValue(initiator.PortName)].ChapSinglePassword
			} else {
				initiatorModel.ChapSinglePassword = stateInitiatorMap[types.StringValue(initiator.PortName)].ChapSinglePassword
			}
			initiatorModel.ChapSingleUsername = types.StringValue(initiator.ChapSingleUsername)
		} else {
			initiatorModel.PortName = types.StringValue(initiator.PortName)
			initiatorModel.PortType = types.StringValue(string(initiator.PortType))
			if operation != operationRead {
				initiatorModel.ChapSinglePassword = planInitiatorMap[types.StringValue(initiator.PortName)].ChapSinglePassword
				initiatorModel.ChapMutualPassword = planInitiatorMap[types.StringValue(initiator.PortName)].ChapMutualPassword
			} else {
				initiatorModel.ChapSinglePassword = stateInitiatorMap[types.StringValue(initiator.PortName)].ChapSinglePassword
				initiatorModel.ChapMutualPassword = stateInitiatorMap[types.StringValue(initiator.PortName)].ChapMutualPassword
			}
			initiatorModel.ChapMutualUsername = types.StringValue(initiator.ChapMutualUsername)
			initiatorModel.ChapSingleUsername = types.StringValue(initiator.ChapSingleUsername)
		}

		initiators = append(initiators, initiatorModel)
	}
	state.Initiators = initiators
	if operation == operationRead {
		state.HostGroupID = types.StringValue(response.HostGroupID)
	}
}

// Attributes to be updated in update operation
func (r resourceHost) addInitiators(plan, state models.Host) *gopowerstore.HostModify {

	hostUpdate := &gopowerstore.HostModify{}
	name := plan.Name.ValueString()
	description := plan.Description.ValueString()

	if plan.HostConnectivity.ValueString() != state.HostConnectivity.ValueString() {
		hostUpdate.HostConnectivity = gopowerstore.HostConnectivityEnum(plan.HostConnectivity.ValueString())
	}

	// Create a map of initiators from state with PortName as key, as it is unique
	stateInitiatorsMap := r.getInitiatorMap(state.Initiators)

	// Create map of initiators to be added
	addInitiatorsMap := make(map[types.String]models.InitiatorCreateModify)
	for i := 0; i < len(plan.Initiators); i++ {
		_, found := stateInitiatorsMap[plan.Initiators[i].PortName]
		if !found {
			addInitiatorsMap[plan.Initiators[i].PortName] = plan.Initiators[i]
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
func (r resourceHost) removeInitiators(plan, state models.Host) *gopowerstore.HostModify {

	hostUpdate := &gopowerstore.HostModify{}

	// Create a map of initiators from plan with PortName as key, as it is unique
	planInitiatorsMap := r.getInitiatorMap(plan.Initiators)

	// create a map to find initiators to be removed
	removeInitiatorsMap := make(map[types.String]models.InitiatorCreateModify)
	for i := 0; i < len(state.Initiators); i++ {
		_, found := planInitiatorsMap[state.Initiators[i].PortName]
		if !found {
			removeInitiatorsMap[state.Initiators[i].PortName] = state.Initiators[i]
		}
	}

	// Fetch keys (port names) to be removed
	var removeInitiators []string
	for removeID := range removeInitiatorsMap {
		removeInitiators = append(removeInitiators, removeID.ValueString())
	}

	hostUpdate = &gopowerstore.HostModify{
		RemoveInitiators: &removeInitiators,
	}

	return hostUpdate
}

// to perform modify operation in update
func (r resourceHost) modifyOperation(plan, state models.Host) *gopowerstore.HostModify {
	hostUpdate := &gopowerstore.HostModify{}
	// update CHAP credentials based on port name
	modifyInitiators := make([]gopowerstore.UpdateInitiatorInHost, 0, len(plan.Initiators))

	stateInitiatorMap := r.getInitiatorMap(state.Initiators)
	for _, initiator := range plan.Initiators {
		if _, ok := stateInitiatorMap[initiator.PortName]; ok {

			var updateInitiator gopowerstore.UpdateInitiatorInHost

			portName := initiator.PortName.ValueString()
			chapMutualPassword := initiator.ChapMutualPassword.ValueString()
			chapMutualUsername := initiator.ChapMutualUsername.ValueString()
			chapSinglePassword := initiator.ChapSinglePassword.ValueString()
			chapSingleUsername := initiator.ChapSingleUsername.ValueString()

			var chapSingleUsername1 *string
			chapSingleUsername1 = &chapSingleUsername
			var chapSinglePassword1 *string
			chapSinglePassword1 = &chapSinglePassword
			var chapMutualUsername1 *string
			chapMutualUsername1 = &chapMutualUsername
			var chapMutualPassword1 *string
			chapMutualPassword1 = &chapMutualPassword

			if chapMutualUsername == "" && chapSingleUsername == "" {
				updateInitiator = gopowerstore.UpdateInitiatorInHost{
					PortName: &portName,
				}
			} else if chapMutualUsername == "" {
				updateInitiator = gopowerstore.UpdateInitiatorInHost{
					PortName:           &portName,
					ChapSinglePassword: chapSinglePassword1,
					ChapSingleUsername: chapSingleUsername1,
				}
			} else if chapMutualUsername != "" {
				updateInitiator = gopowerstore.UpdateInitiatorInHost{
					PortName:           &portName,
					ChapMutualPassword: chapMutualPassword1,
					ChapMutualUsername: chapMutualUsername1,
					ChapSinglePassword: chapSinglePassword1,
					ChapSingleUsername: chapSingleUsername1,
				}
			}

			modifyInitiators = append(modifyInitiators, updateInitiator)
			hostUpdate = &gopowerstore.HostModify{
				ModifyInitiators: &modifyInitiators,
			}
		}
	}
	return hostUpdate
}

func (r *resourceHost) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.Host
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	for _, initiator := range data.Initiators {
		if initiator.ChapMutualUsername != types.StringNull() && initiator.ChapSingleUsername == types.StringNull() {
			resp.Diagnostics.AddError(
				"Error validating config host",
				"`chap_mutual_username` cannot pe present without `chap_single_username`",
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

func (r resourceHost) getInitiatorMap(initiators []models.InitiatorCreateModify) map[types.String]models.InitiatorCreateModify {
	initiatorsMap := make(map[types.String]models.InitiatorCreateModify)
	if len(initiators) != 0 {
		for i := 0; i < len(initiators); i++ {
			initiatorsMap[initiators[i].PortName] = initiators[i]
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

	var chapSingleUsername1 *string
	chapSingleUsername1 = &chapSingleUsername
	var chapSinglePassword1 *string
	chapSinglePassword1 = &chapSinglePassword
	var chapMutualUsername1 *string
	chapMutualUsername1 = &chapMutualUsername
	var chapMutualPassword1 *string
	chapMutualPassword1 = &chapMutualPassword

	// When port type is iSCSI only then look for CHAP Username and Password
	if portType != "iSCSI" || (chapMutualUsername == "" && chapSingleUsername == "") {
		initiator = gopowerstore.InitiatorCreateModify{
			PortName: &portName,
			PortType: (*gopowerstore.InitiatorProtocolTypeEnum)(&portType),
		}
	} else if chapMutualUsername == "" {
		initiator = gopowerstore.InitiatorCreateModify{
			PortName:           &portName,
			PortType:           (*gopowerstore.InitiatorProtocolTypeEnum)(&portType),
			ChapSinglePassword: chapSinglePassword1,
			ChapSingleUsername: chapSingleUsername1,
		}
	} else {
		initiator = gopowerstore.InitiatorCreateModify{
			PortName:           &portName,
			PortType:           (*gopowerstore.InitiatorProtocolTypeEnum)(&portType),
			ChapMutualPassword: chapMutualPassword1,
			ChapMutualUsername: chapMutualUsername1,
			ChapSinglePassword: chapSinglePassword1,
			ChapSingleUsername: chapSingleUsername1,
		}
	}
	return initiator
}
