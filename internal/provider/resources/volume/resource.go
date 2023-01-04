package volume

import (
	"context"
	"fmt"
	"log"
	"strings"
	"terraform-provider-powerstore/internal/powerstore"
	"terraform-provider-powerstore/internal/provider/planmodifiers"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure Resource satsisfies resource.Resource interface
var _ resource.Resource = &Resource{}
var _ resource.ResourceWithImportState = &Resource{}

func NewResource() resource.Resource {
	return &Resource{}
}

// Resource defines the resource implementation.
type Resource struct {
	client *powerstore.Client
}

func (r *Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_volume"
}

func (r *Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "Volume resource",

		Attributes: map[string]schema.Attribute{

			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The ID of the volume.",
				MarkdownDescription: "The ID of the volume.",
			},

			"name": schema.StringAttribute{
				Required:            true,
				Description:         "The name of the volume.",
				MarkdownDescription: "The name of the volume.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"size": schema.Float64Attribute{
				Required:            true,
				Description:         "The size of the volume.",
				MarkdownDescription: "The size of the volume.",
				// todo validators
			},

			"capacity_unit": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The Capacity Unit corresponding to the size.",
				MarkdownDescription: "The Capacity Unit corresponding to the size.",
				PlanModifiers: []planmodifier.String{
					planmodifiers.DefaultAttribute("GB"),
				},
				// todo validators
			},

			"host_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The host id of the volume.",
				MarkdownDescription: "The host id of the volume.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),

					stringvalidator.ConflictsWith(path.Expressions{
						path.MatchRoot("host_name"),
						path.MatchRoot("host_group_id"),
						path.MatchRoot("host_group_name"),
					}...),
				},
			},

			"host_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The host name of the volume.",
				MarkdownDescription: "The host name of the volume.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),

					stringvalidator.ConflictsWith(path.Expressions{
						path.MatchRoot("host_group_id"),
						path.MatchRoot("host_group_name"),
					}...),
				},
			},

			"host_group_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The host group id of the volume.",
				MarkdownDescription: "The host group id of the volume.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),

					stringvalidator.ConflictsWith(path.Expressions{
						path.MatchRoot("host_group_name"),
					}...),
				},
			},

			"host_group_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The host group name of the volume.",
				MarkdownDescription: "The host group name of the volume.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"logical_unit_number": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Optional logical unit number when creating a attached volume. ",
				MarkdownDescription: "Optional logical unit number when creating a attached volume. ",
			},

			"volume_group_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The volume group id of the volume.",
				MarkdownDescription: "The volume group id of the volume.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),

					stringvalidator.ConflictsWith(path.Expressions{
						path.MatchRoot("volume_group_name"),
					}...),
				},
			},

			"volume_group_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The volume group name of the volume.",
				MarkdownDescription: "The volume group name of the volume.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"min_size": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Optional minimum size for the volume, in bytes.",
				MarkdownDescription: "Optional minimum size for the volume, in bytes.",
			},

			"sector_size": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Optional sector size, in bytes",
				MarkdownDescription: "Optional sector size, in bytes",
				Validators: []validator.Int64{
					int64validator.OneOf([]int64{512, 4096}...),
				},
				PlanModifiers: []planmodifier.Int64{
					planmodifiers.DefaultAttribute(int64(defaultSectorSize)),
				},
			},

			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The description of the volume.",
				MarkdownDescription: "The description of the volume.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"appliance_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The appliance_id of the volume.",
				MarkdownDescription: "The appliance_id of the volume.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),

					stringvalidator.ConflictsWith(path.Expressions{
						path.MatchRoot("appliance_name"),
					}...),
				},
			},

			"appliance_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The appliance name of the volume.",
				MarkdownDescription: "The appliance name of the volume.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"protection_policy_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The protection_policy_id of the volume.",
				MarkdownDescription: "The protection_policy_id of the volume.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),

					stringvalidator.ConflictsWith(path.Expressions{
						path.MatchRoot("protection_policy_name"),
					}...),
				},
			},

			"protection_policy_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The protection policy name of the volume.",
				MarkdownDescription: "The protection policy name of the volume.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"performance_policy_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The performance_policy_id of the volume.",
				MarkdownDescription: "The performance_policy_id of the volume.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				PlanModifiers: []planmodifier.String{
					planmodifiers.DefaultAttribute("default_medium"),
				},
			},

			"creation_timestamp": schema.StringAttribute{
				Computed:            true,
				Description:         "The creation_timestamp of the volume.",
				MarkdownDescription: "The creation_timestamp of the volume.",
			},

			"is_replication_destination": schema.BoolAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "The is_replication_destination of the volume.",
				MarkdownDescription: "The is_replication_destination of the volume.",
			},

			"node_affinity": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "The node_affinity of the volume.",
				MarkdownDescription: "The node_affinity of the volume.",
			},

			"type": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "The type of the volume.",
				MarkdownDescription: "The type of the volume.",
			},

			"app_type": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "The app type of the volume.",
				MarkdownDescription: "The app type of the volume.",
			},

			"app_type_other": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "The app type other of the volume.",
				MarkdownDescription: "The app type other of the volume.",
			},

			"wwn": schema.StringAttribute{
				Computed:            true,
				Description:         "The wwn of the volume.",
				MarkdownDescription: "The wwn of the volume.",
			},

			"state": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "The state of the volume.",
				MarkdownDescription: "The state of the volume.",
			},

			"nguid": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "The nguid of the volume.",
				MarkdownDescription: "The nguid of the volume.",
			},

			"nsid": schema.Int64Attribute{
				Computed:            true,
				Optional:            true,
				Description:         "The nsid of the volume.",
				MarkdownDescription: "The nsid of the volume.",
			},

			"logical_used": schema.Int64Attribute{
				Computed:            true,
				Optional:            true,
				Description:         "Current amount of data used by the volume.",
				MarkdownDescription: "Current amount of data used by the volume.",
			},
		},
	}
}

func (r *Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*powerstore.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var valInBytes int64

	log.Printf("Started Creating Volume")
	var plan model

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	valInBytes, errmsg := convertToBytes(ctx, plan)
	if errmsg != "" {
		resp.Diagnostics.AddError(
			"Error creating volume",
			"Could not create volume "+errmsg,
		)
		return
	}

	valid, errmsg := fetchByName(r.client, &plan)
	if !valid {
		resp.Diagnostics.AddError(
			"Error creating volume",
			"Could not create volume, either of "+errmsg+" should be present",
		)
		return
	}

	name := plan.Name.ValueString()
	sectorSize := plan.SectorSize.ValueInt64()

	volumeCreate := &gopowerstore.VolumeCreate{
		Name:                &name,
		Description:         plan.Description.ValueString(),
		Size:                &valInBytes,
		ApplianceID:         plan.ApplianceID.ValueString(),
		VolumeGroupID:       plan.VolumeGroupID.ValueString(),
		SectorSize:          &sectorSize,
		ProtectionPolicyID:  plan.ProtectionPolicyID.ValueString(),
		PerformancePolicyID: plan.PerformancePolicyID.ValueString(),
		AppType:             plan.AppType.ValueString(),
		AppTypeOther:        plan.AppTypeOther.ValueString(),
		MinimumSize:         plan.MinimumSize.ValueInt64(),
		HostID:              plan.HostID.ValueString(),
		HostGroupID:         plan.HostGroupID.ValueString(),
		LogicalUnitNumber:   plan.LogicalUnitNumber.ValueInt64(),
	}

	// Add validation
	valid, validErr := creationValidation(context.Background(), plan)
	if !valid {
		resp.Diagnostics.AddError(
			"Error creating volume",
			"Could not create volume "+validErr,
		)
		return
	}

	// Create New Volume
	// The function returns only ID of the newly created Volume
	volCreateResponse, err := r.client.PStoreClient.CreateVolume(context.Background(), volumeCreate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating volume",
			"Could not create volume, unexpected error: "+err.Error(),
		)
		return
	}

	// Get Volume Details using ID retrieved above
	volResponse, err1 := r.client.PStoreClient.GetVolume(context.Background(), volCreateResponse.ID)
	if err1 != nil {
		resp.Diagnostics.AddError(
			"Error getting volume after creation",
			"Could not get volume, unexpected error: "+err.Error(),
		)
		return
	}
	// Get Host Mapping from volume ID
	hostMapping, err1 := r.client.PStoreClient.GetHostVolumeMappingByVolumeID(context.Background(), volCreateResponse.ID)
	if err1 != nil {
		resp.Diagnostics.AddError(
			"Error fetching volume host mapping",
			"Could not create volume, unexpected error: "+err.Error(),
		)
		return
	}
	// Get Volume Group Mapping details from API
	volGroupMapping, err := r.client.PStoreClient.GetVolumeGroupsByVolumeID(context.Background(), volCreateResponse.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error fetching volume host mapping",
			"Could not create volume, unexpected error: "+err.Error(),
		)
		return
	}
	log.Printf("After Volume create call")

	result := model{}
	updateVolState(&result, volResponse, hostMapping, volGroupMapping, &plan, operationCreate)

	log.Printf("Added to result: %v", result)

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Create")
}

func (r *Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state model
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get volume details from API and then update what is in state from what the API returns
	volID := state.ID.ValueString()
	volResponse, err := r.client.PStoreClient.GetVolume(context.Background(), volID)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading volume",
			"Could not read volume with error "+volID+": "+err.Error(),
		)
		return
	}
	// Get Host Mapping details from API
	hostMapping, err := r.client.PStoreClient.GetHostVolumeMappingByVolumeID(context.Background(), volID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error fetching volume host mapping",
			"Could not create volume, unexpected error: "+err.Error(),
		)
		return
	}
	// Get Volume Group Mapping details from API
	volGroupMapping, err := r.client.PStoreClient.GetVolumeGroupsByVolumeID(context.Background(), volID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error fetching volume host mapping",
			"Could not create volume, unexpected error: "+err.Error(),
		)
		return
	}

	updateVolState(&state, volResponse, hostMapping, volGroupMapping, nil, operationRead)

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Done with Read")
}

func (r *Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("Started Update")

	// Get plan values
	var plan model
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state model
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	valid, errmsg := fetchByName(r.client, &plan)
	if !valid {
		resp.Diagnostics.AddError(
			"Error creating volume",
			"Could not Update volume, either of "+errmsg+" should be present",
		)
		return
	}

	// Update volume parameters. In case of validation failure, return
	updatedParams, updateFailedParameters, errMessages := updateVol(ctx, r.client, plan, state)
	if len(updateFailedParameters) > 0 && updateFailedParameters[0] == "Validation Failed" {
		resp.Diagnostics.AddError(
			"Validation Check Failed",
			errMessages[0],
		)
		return
	}
	// Get vg ID from state
	volID := state.ID.ValueString()

	if len(errMessages) > 0 || len(updateFailedParameters) > 0 {
		errMessage := strings.Join(errMessages, ",\n")
		resp.Diagnostics.AddError(
			fmt.Sprintf("Failed to update all parameters of Volume, updated parameters are %v and parameters failed to update are %v", updatedParams, updateFailedParameters),
			errMessage)
	}

	// Get volume details from volume ID
	volResponse, err := r.client.PStoreClient.GetVolume(context.Background(), volID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting volume after update",
			"Could not get after update volID "+volID+": "+err.Error(),
		)
		return
	}

	// Get Host Mapping from volume ID
	hostMapping, err := r.client.PStoreClient.GetHostVolumeMappingByVolumeID(context.Background(), volID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error fetching volume host mapping",
			"Could not update volume, unexpected error: "+err.Error(),
		)
		return
	}

	// Get Volume Group Mapping details from API
	volGroupMapping, err := r.client.PStoreClient.GetVolumeGroupsByVolumeID(context.Background(), volID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error fetching volume host mapping",
			"Could not create volume, unexpected error: "+err.Error(),
		)
		return
	}

	updateVolState(&state, volResponse, hostMapping, volGroupMapping, &plan, operationUpdate)

	//Set State
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Done with Update")
}

func (r *Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	log.Printf("Started with Delete")

	var state model
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get vg ID from state
	volID := state.ID.ValueString()
	if state.ProtectionPolicyID.ValueString() != "" {
		state.ProtectionPolicyID = types.StringValue("")
		modifyVolume(state, 0, volID, r.client)
	}

	if state.HostID.ValueString() != "" || state.HostGroupID.ValueString() != "" {
		err := detachHostFromVolume(state, model{}, r.client, volID)
		if err != nil {
			resp.Diagnostics.AddError(
				"Cannot detach volume host mapping",
				"Could not delete volume, unexpected error: "+err.Error(),
			)
			return
		}
	}

	if state.VolumeGroupID.ValueString() != "" {
		err := detachVolumeGroup(ctx, state, r.client, volID)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error detaching volume Group",
				"Could not delete volume, unexpected error: "+err.Error(),
			)
			return

		}
	}

	// Delete volume by calling API
	_, err := r.client.PStoreClient.DeleteVolume(context.Background(), nil, volID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting volume",
			"Could not delete volID "+volID+": "+err.Error(),
		)
		return
	}

	// Remove resource from state
	resp.State.RemoveResource(ctx)
	log.Printf("Done with Delete")
}

func (r *Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	log.Printf("Started with Import")

	// fetching Volume information
	response, err := r.client.PStoreClient.GetVolume(context.Background(), req.ID)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error importing Volume",
			fmt.Sprintf("Could not import volume: %s with error: %s", req.ID, err.Error()),
		)
		return
	}

	state := model{}

	// as state is like a plan here, a current state prior to this import operation
	updateVolState(&state, response, nil, gopowerstore.VolumeGroups{}, &state, operationImport)

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Done with Import")
}
