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
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const defaultSectorSize = 512

type volumeResource struct {
	client *client.Client
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &volumeResource{}
	_ resource.ResourceWithConfigure   = &volumeResource{}
	_ resource.ResourceWithImportState = &volumeResource{}
)

// NewVolumeResource is a helper function to simplify the provider implementation.
func newVolumeResource() resource.Resource {
	return &volumeResource{}
}

func (r volumeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_volume"
}

func (r volumeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This resource is used to manage the volume entity of PowerStore Array. We can Create, Update and Delete the volume using this resource. We can also import an existing volume from PowerStore array.",
		MarkdownDescription: "This resource is used to manage the volume entity of PowerStore Array. We can Create, Update and Delete the volume using this resource. We can also import an existing volume from PowerStore array.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "The ID of the volume.",
				MarkdownDescription: "The ID of the volume.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				Description:         "The name of the volume.",
				MarkdownDescription: "The name of the volume.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"size": schema.Float64Attribute{
				Description:         "The size of the volume.",
				MarkdownDescription: "The size of the volume.",
				Required:            true,
			},
			"capacity_unit": schema.StringAttribute{
				Description:         "The Capacity Unit corresponding to the size.",
				MarkdownDescription: "The Capacity Unit corresponding to the size.",
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					DefaultAttribute("GB"),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						"MB",
						"GB",
						"TB",
					}...),
				},
			},
			"host_id": schema.StringAttribute{
				Description:         "The host id of the volume.",
				MarkdownDescription: "The host id of the volume.",
				Computed:            true,
				Optional:            true,
			},

			"host_name": schema.StringAttribute{
				Optional:            true,
				Description:         "The host name of the volume.",
				MarkdownDescription: "The host name of the volume.",
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("host_id")),
				},
			},

			"host_group_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The host group id of the volume.",
				MarkdownDescription: "The host group id of the volume.",
			},
			"host_group_name": schema.StringAttribute{
				Optional:            true,
				Description:         "The host group name of the volume.",
				MarkdownDescription: "The host group name of the volume.",
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("host_group_id")),
				},
			},
			"logical_unit_number": schema.Int64Attribute{
				Computed:            true,
				Optional:            true,
				Description:         "The current amount of data written to the volume.",
				MarkdownDescription: "The current amount of data written to the volume.",
			},
			"volume_group_id": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "The volume group id of the volume.",
				MarkdownDescription: "The volume group id of the volume.",
			},
			"volume_group_name": schema.StringAttribute{
				Optional:            true,
				Description:         "The volume group name of the volume.",
				MarkdownDescription: "The volume group name of the volume.",
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("volume_group_id")),
				},
			},
			"min_size": schema.Int64Attribute{
				Optional:            true,
				Description:         "The minimum size  of the volume.",
				MarkdownDescription: "The minimum size of the volume.",
			},
			"sector_size": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "The sector size of the volume.",
				MarkdownDescription: "The sector size of the volume.",
				PlanModifiers: []planmodifier.Int64{
					DefaultAttribute(int64(defaultSectorSize)),
				},
			},
			"description": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "The description of the volume.",
				MarkdownDescription: "The description of the volume.",
			},
			"appliance_id": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "The appliance_id of the volume.",
				MarkdownDescription: "The appliance_id of the volume.",
			},
			"appliance_name": schema.StringAttribute{
				Optional:            true,
				Description:         "The appliance name of the volume.",
				MarkdownDescription: "The appliance name of the volume.",
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("appliance_id")),
				},
			},
			"protection_policy_id": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "The protection_policy_id of the volume.",
				MarkdownDescription: "The protection_policy_id of the volume.",
			},
			"protection_policy_name": schema.StringAttribute{
				Optional:            true,
				Description:         "The protection policy name of the volume.",
				MarkdownDescription: "The protection policy name of the volume.",
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("protection_policy_id")),
				},
			},
			"performance_policy_id": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "The performance_policy_id of the volume.",
				MarkdownDescription: "The performance_policy_id of the volume.",
				PlanModifiers: []planmodifier.String{
					DefaultAttribute("default_medium"),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						"default_medium",
						"default_low",
						"default_high",
					}...),
				},
			},
			"creation_timestamp": schema.StringAttribute{
				Computed:            true,
				Description:         "The creation_timestamp of the volume.",
				MarkdownDescription: "The creation_timestamp of the volume.",
			},
			"is_replication_destination": schema.BoolAttribute{
				Computed:            true,
				Description:         "The is_replication_destination of the volume.",
				MarkdownDescription: "The is_replication_destination of the volume.",
			},
			"node_affinity": schema.StringAttribute{
				Computed:            true,
				Description:         "The node_affinity of the volume.",
				MarkdownDescription: "The node_affinity of the volume.",
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						string(gopowerstore.NodeAffinityEnumSelectAtAttach),
						string(gopowerstore.NodeAffinityEnumSelectNodeA),
						string(gopowerstore.NodeAffinityEnumSelectNodeB),
						string(gopowerstore.NodeAffinityEnumPreferredNodeA),
						string(gopowerstore.NodeAffinityEnumPreferredNodeB),
					}...),
				},
			},
			"type": schema.StringAttribute{
				Computed:            true,
				Description:         "The type of the volume.",
				MarkdownDescription: "The type of the volume.",
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						string(gopowerstore.VolumeTypeEnumPrimary),
						string(gopowerstore.VolumeTypeEnumSnapshot),
						string(gopowerstore.VolumeTypeEnumClone),
					}...),
				},
			},
			"app_type": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "The app type of the volume.",
				MarkdownDescription: "The app type of the volume.",
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						string(gopowerstore.AppTypeEnumRelationDB),
						string(gopowerstore.AppTypeEnumOracle),
						string(gopowerstore.AppTypeEnumSQLServer),
						string(gopowerstore.AppTypeEnumPostgreSQL),
						string(gopowerstore.AppTypeEnumMySQL),
						string(gopowerstore.AppTypeEnumIBMDB2),
						string(gopowerstore.AppTypeEnumBigData),
						string(gopowerstore.AppTypeEnumMongoDB),
						string(gopowerstore.AppTypeEnumCassandra),
						string(gopowerstore.AppTypeEnumSAPHANA),
						string(gopowerstore.AppTypeEnumSpark),
						string(gopowerstore.AppTypeEnumSplunk),
						string(gopowerstore.AppTypeEnumElasticSearch),
						string(gopowerstore.AppTypeEnumExchange),
						string(gopowerstore.AppTypeEnumSharepoint),
						string(gopowerstore.AppTypeEnumRBusinessApplicationsOther),
						string(gopowerstore.AppTypeEnumRelationERPSAP),
						string(gopowerstore.AppTypeEnumCRM),
						string(gopowerstore.AppTypeEnumHealthcareOther),
						string(gopowerstore.AppTypeEnumEpic),
						string(gopowerstore.AppTypeEnumMEDITECH),
						string(gopowerstore.AppTypeEnumAllscripts),
						string(gopowerstore.AppTypeEnumCerner),
						string(gopowerstore.AppTypeEnumVirtualization),
						string(gopowerstore.AppTypeEnumVirtualServers),
						string(gopowerstore.AppTypeEnumContainers),
						string(gopowerstore.AppTypeEnumVirtualDesktops),
						string(gopowerstore.AppTypeEnumRelationOther),
					}...),
				},
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
				Description:         "The state of the volume.",
				MarkdownDescription: "The state of the volume.",
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						string(gopowerstore.VolumeStateEnumReady),
						string(gopowerstore.VolumeStateEnumInitializing),
						string(gopowerstore.VolumeStateEnumOffline),
						string(gopowerstore.VolumeStateEnumDestroying),
					}...),
				},
			},
			"nguid": schema.StringAttribute{
				Computed:            true,
				Description:         "The nguid of the volume.",
				MarkdownDescription: "The nguid of the volume.",
			},
			"nsid": schema.Int64Attribute{
				Computed:            true,
				Description:         "The nsid of the volume.",
				MarkdownDescription: "The nsid of the volume.",
			},
			"logical_used": schema.Int64Attribute{
				Computed:            true,
				Description:         "Current amount of data used by the volume.",
				MarkdownDescription: "Current amount of data used by the volume.",
			},
		},
	}
}

// Configure - defines configuration for volume resource.
func (r *volumeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// Create - method to create volume resource
func (r volumeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var valInBytes int64

	log.Printf("Started Creating Volume")
	var plan models.Volume

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

	errmsg = fetchByName(*r.client, &plan)
	if errmsg != "" {
		resp.Diagnostics.AddError(
			"Error creating volume",
			"Could not create volume, "+errmsg,
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
		AppType:             gopowerstore.AppTypeEnum(plan.AppType.ValueString()),
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
			"Could not get volume, unexpected error: "+err1.Error(),
		)
		return
	}
	// Get Host Mapping from volume ID
	hostMapping, err1 := r.client.PStoreClient.GetHostVolumeMappingByVolumeID(context.Background(), volCreateResponse.ID)
	if err1 != nil {
		resp.Diagnostics.AddError(
			"Error fetching volume host mapping",
			"Could not create volume, unexpected error: "+err1.Error(),
		)
		return
	}
	// Get Volume Group Mapping details from API
	volGroupMapping, err := r.client.PStoreClient.GetVolumeGroupsByVolumeID(context.Background(), volCreateResponse.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error fetching volume group mapping",
			"Could not create volume, unexpected error: "+err.Error(),
		)
		return
	}
	log.Printf("After Volume create call")

	result := models.Volume{}
	updateVolState(&result, volResponse, hostMapping, volGroupMapping, &plan, operationCreate)

	log.Printf("Added to result: %v", result)

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Create")
}

// Read - reads volume resource
func (r volumeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.Volume
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	volID := state.ID.ValueString()
	err := r.performRead(ctx, volID, &state)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading volume",
			"Could not read volume with error "+volID+": "+err.Error(),
		)
	}

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Done with Read")
}

// Update - updates volume resource
func (r volumeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("Started Update")

	// Get plan values
	var plan models.Volume
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state models.Volume
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	errmsg := fetchByName(*r.client, &plan)
	if errmsg != "" {
		resp.Diagnostics.AddError(
			"Error Updating volume",
			"Could not Update volume, "+errmsg,
		)
		return
	}

	// Update volume parameters. In case of validation failure, return
	updatedParams, updateFailedParameters, errMessages := updateVol(ctx, *r.client, plan, state)
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
	hostMapping, err := r.client.PStoreClient.GetHostVolumeMappingByVolumeID(context.Background(), volResponse.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error fetching volume host mapping",
			"Could not update volume, unexpected error: "+err.Error(),
		)
		return
	}

	// Get Volume Group Mapping details from API
	volGroupMapping, err := r.client.PStoreClient.GetVolumeGroupsByVolumeID(context.Background(), volResponse.ID)
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

// Delete - method to delete volume resource
func (r volumeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("Started with Delete")

	var state models.Volume
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get vg ID from state
	volID := state.ID.ValueString()

	// perform read refresh to update state before deletion
	err := r.performRead(ctx, volID, &state)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading volume",
			"Could not read volume with error "+volID+": "+err.Error(),
		)
	}

	// Detach protection policy from volume
	if state.ProtectionPolicyID.ValueString() != "" {
		state.ProtectionPolicyID = types.StringNull()
		err := modifyVolume(state, 0, volID, *r.client)
		if err != nil {
			resp.Diagnostics.AddError(
				"Cannot detach protection policy",
				"Could not delete volume, unexpected error: "+err.Error(),
			)
			return
		}
	}

	if state.HostID.ValueString() != "" || state.HostGroupID.ValueString() != "" {
		err := detachHostFromVolume(state, models.Volume{}, *r.client, volID)
		if err != nil {
			resp.Diagnostics.AddError(
				"Cannot detach volume host mapping",
				"Could not delete volume, unexpected error: "+err.Error(),
			)
			return
		}
	}

	if state.VolumeGroupID.ValueString() != "" {
		err := detachVolumeGroup(ctx, state, *r.client, volID)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error detaching volume Group",
				"Could not delete volume, unexpected error: "+err.Error(),
			)
			return

		}
	}

	// Delete volume by calling API
	_, err = r.client.PStoreClient.DeleteVolume(context.Background(), nil, volID)
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

// ImportState import state for existing volume
func (r volumeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
