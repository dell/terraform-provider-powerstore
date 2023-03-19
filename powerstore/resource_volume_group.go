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

// newVolumeGroupResource returns new volume group resource instance
func newVolumeGroupResource() resource.Resource {
	return &resourceVolumeGroup{}
}

type resourceVolumeGroup struct {
	client *client.Client
}

// Metadata defines resource interface Metadata method
func (r *resourceVolumeGroup) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_volumegroup"
}

// Schema defines resource interface Schema method
func (r *resourceVolumeGroup) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "VolumeGroup resource",

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
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"volume_ids": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "A list of identifiers of existing volumes that should be added to the volume group.",
				MarkdownDescription: "A list of identifiers of existing volumes that should be added to the volume group.",
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
				Description:         "Unique identifier of the protection policy assigned to the volume group.",
				MarkdownDescription: "Unique identifier of the protection policy assigned to the volume group.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ConflictsWith(path.Expressions{
						path.MatchRoot("protection_policy_name"),
					}...),
				},
			},

			"volume_names": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "A list of names of existing volumes that should be added to the volume group.",
				MarkdownDescription: "A list of names of existing volumes that should be added to the volume group.",
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(
						stringvalidator.LengthAtLeast(1),
					),
				},
			},

			"protection_policy_name": schema.StringAttribute{
				Optional:            true,
				Description:         "Unique name of the protection policy assigned to the volume group.",
				MarkdownDescription: "Unique name of the protection policy assigned to the volume group.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
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

	r.client = client
}

// Create - method to create volume group resource
func (r *resourceVolumeGroup) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.Volumegroup

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var volumeIds []string
	for _, volume := range plan.VolumeIDs.Elements() {
		volumeIds = append(volumeIds, strings.Trim(volume.String(), "\""))
	}

	volumeGroupCreate := &gopowerstore.VolumeGroupCreate{
		Name:                   plan.Name.ValueString(),
		Description:            plan.Description.ValueString(),
		VolumeIds:              volumeIds,
		IsWriteOrderConsistent: plan.IsWriteOrderConsistent.ValueBool(),
		ProtectionPolicyID:     plan.ProtectionPolicyID.ValueString(),
	}

	//Create New Volume Group
	volGroupCreateResponse, err := r.client.PStoreClient.CreateVolumeGroup(context.Background(), volumeGroupCreate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating volume group",
			"Could not create volume group, unexpected error: "+err.Error(),
		)
		return
	}

	//Get Volume Group details using ID retrived above
	volGroupResponse, err := r.client.PStoreClient.GetVolumeGroup(context.Background(), volGroupCreateResponse.ID)
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

	//Delete Volume Group by calling API
	_, err := r.client.PStoreClient.DeleteVolumeGroup(context.Background(), volumeGroupID)
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
	response, err := r.client.PStoreClient.GetVolumeGroup(context.Background(), id)

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

// Update - method to update volume group resource
func (r *resourceVolumeGroup) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// updateVolGroupState - method to update terraform state
func (r resourceVolumeGroup) updateVolGroupState(volgroupState *models.Volumegroup, volGroupResponse gopowerstore.VolumeGroup, volGroupPlan *models.Volumegroup) {
	// Update value from Volume Group Response to State
	volgroupState.ID = types.StringValue(volGroupResponse.ID)
	volgroupState.Name = types.StringValue(volGroupResponse.Name)
	volgroupState.Description = types.StringValue(volGroupResponse.Description)
	volgroupState.IsWriteOrderConsistent = types.BoolValue(volGroupResponse.IsWriteOrderConsistent)
	volgroupState.ProtectionPolicyID = types.StringValue(volGroupResponse.ProtectionPolicyID)

	//Update VolumeIDs value from Response to State
	var volumeIds []string
	for _, volume := range volGroupResponse.Volumes {
		volumeIds = append(volumeIds, volume.ID)
	}
	volumeIDList := []attr.Value{}
	for _, volumeID := range volumeIds {
		volumeIDList = append(volumeIDList, types.StringValue(string(volumeID)))
	}
	volgroupState.VolumeIDs, _ = types.SetValue(types.StringType, volumeIDList)

	//Update VolumeNames value from Plan to State
	var volumeNames []string
	for _, volumeName := range volGroupPlan.VolumeNames.Elements() {
		volumeNames = append(volumeNames, strings.Trim(volumeName.String(), "\""))
	}
	volumeNameList := []attr.Value{}
	for _, volumeName := range volumeNames {
		volumeNameList = append(volumeNameList, types.StringValue(string(volumeName)))
	}
	volgroupState.VolumeNames, _ = types.SetValue(types.StringType, volumeNameList)

	//Update ProtectionPolicyName value from Plan to State
	volgroupState.ProtectionPolicyName = volGroupPlan.ProtectionPolicyName
}

// ModifyPlan modify resource plan attribute value
func (r *resourceVolumeGroup) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var plan models.Volumegroup

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var volumeIds []string
	if len(plan.VolumeNames.Elements()) != 0 {
		for _, volumeName := range plan.VolumeNames.Elements() {
			volume, err := r.client.PStoreClient.GetVolumeByName(context.Background(), strings.Trim(volumeName.String(), "\""))
			if err != nil {
				resp.Diagnostics.AddError(
					"Error getting volume ",
					"Could not get volume with name: "+strings.Trim(volumeName.String(), "\"")+", \n unexpected error: "+err.Error(),
				)
				return
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
		policy, err := r.client.PStoreClient.GetProtectionPolicyByName(context.Background(), plan.ProtectionPolicyName.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Error getting protection policy",
				"Could not get protection policy with name: "+plan.ProtectionPolicyName.ValueString()+", \n unexpected error: "+err.Error(),
			)
			return
		}
		plan.ProtectionPolicyID = types.StringValue(policy.ID)
	}

	diags = resp.Plan.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
