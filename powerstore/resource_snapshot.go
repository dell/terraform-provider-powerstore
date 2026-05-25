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
	"log"
	"net/url"
	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/models"
	"terraform-provider-powerstore/powerstore/helper"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// newVolumeSnapshotResource returns snapshot new resource instance
func newVolumeSnapshotResource() resource.Resource {
	return &resourceVolumeSnapshot{}
}

type resourceVolumeSnapshot struct {
	client *clientgen.APIClient
}

// Metadata defines resource interface Metadata method
func (r *resourceVolumeSnapshot) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_volume_snapshot"
}

// Schema defines resource interface Schema method
func (r *resourceVolumeSnapshot) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "This resource is used to manage the volume snapshot entity of PowerStore Array. We can Create, Update and Delete the volume snapshot using this resource. We can also import an existing volume snapshot from PowerStore array.",
		Description:         "This resource is used to manage the volume snapshot entity of PowerStore Array. We can Create, Update and Delete the volume snapshot using this resource. We can also import an existing volume snapshot from PowerStore array.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The unique identifier of the volume snapshot.",
				MarkdownDescription: "The unique identifier of the volume snapshot.",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Name of the volume snapshot.The default name of the volume snapshot is the date and time when the snapshot is taken.",
				MarkdownDescription: "Name of the volume snapshot.The default name of the volume snapshot is the date and time when the snapshot is taken.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"volume_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "ID of the volume to take snapshot. Conflicts with `volume_name`. Cannot be updated.",
				MarkdownDescription: "ID of the volume to take snapshot. Conflicts with `volume_name`. Cannot be updated.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ExactlyOneOf(path.MatchRoot("volume_name")),
				},
			},
			"volume_name": schema.StringAttribute{
				Optional:            true,
				Description:         "Name of the volume to take snapshot. Conflicts with `volume_id`. Cannot be updated.",
				MarkdownDescription: "Name of the volume to take snapshot. Conflicts with `volume_id`. Cannot be updated.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ExactlyOneOf(path.MatchRoot("volume_id")),
				},
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Description of the volume snapshot.",
				MarkdownDescription: "Description of the volume snapshot.",
			},
			"performance_policy_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Performance Policy id of the volume snapshot. Valid values are default_medium, default_low, default_high.",
				MarkdownDescription: "Performance Policy id of the volume snapshot. Valid values are default_medium, default_low, default_high.",
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						"default_medium",
						"default_low",
						"default_high",
					}...),
				},
			},
			"expiration_timestamp": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Expiration Timestamp of the volume snapshot.Only UTC (+Z) format is allowed",
				MarkdownDescription: "Expiration Timestamp of the volume snapshot.Only UTC (+Z) format is allowed.",
				CustomType:          timetypes.RFC3339Type{},
			},
			"creator_type": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "Creator Type of the volume snapshot.",
				MarkdownDescription: "Creator Type of the volume snapshot.",
				PlanModifiers: []planmodifier.String{
					DefaultAttribute("User"),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						"User",
					}...),
				},
			},
			"is_secure": schema.BoolAttribute{
				Description:         "Indicates whether the snapshot is secure. Secure snapshots cannot be deleted before the expiration time, and the expiration time cannot be reduced.",
				MarkdownDescription: "Indicates whether the snapshot is secure. Secure snapshots cannot be deleted before the expiration time, and the expiration time cannot be reduced.",
				Optional:            true,
				Computed:            true,
			},
		},
	}
}

// Configure - defines configuration for volume snapshot resource
func (r *resourceVolumeSnapshot) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client := req.ProviderData.(*client.Client)
	r.client = client.GenClient
}

// Create - method to create volume snapshot resource
func (r *resourceVolumeSnapshot) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan models.Snapshot

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	volID := plan.VolumeID.ValueString()

	// if volume name is present instead of ID
	if plan.VolumeID.ValueString() == "" {
		queries := make(url.Values)
		queries.Set("name", "eq."+plan.VolumeName.ValueString())
		volResponse, _, err := r.client.VolumeApi.GetAllVolumes(ctx).Queries(queries).Execute()
		if err != nil {
			resp.Diagnostics.AddError(
				"Error creating volume snapshot",
				"Could not fetch volume ID from volume name, unexpected error: "+err.Error(),
			)
			return
		}
		if len(volResponse) == 0 {
			resp.Diagnostics.AddError(
				"Error creating volume snapshot",
				"Volume not found",
			)
			return
		}
		volID = *volResponse[0].Id
		plan.VolumeID = types.StringValue(volID)
	}

	name := plan.Name.ValueString()

	// If name of the snapshot is not present, the default name of the volume snapshot is the date and time when the snapshot is taken.
	if name == "" {
		sel := "system_time"
		queries := make(url.Values)
		queries.Set("select", sel)
		clusterResponse, _, err := r.client.ClusterApi.GetAllClusters(ctx).Queries(queries).Execute()
		if err != nil {
			resp.Diagnostics.AddError(
				"Error creating volume snapshot",
				"Could not fetch cluster time, unexpected error: "+err.Error(),
			)
			return
		}
		if len(clusterResponse) == 0 {
			resp.Diagnostics.AddError(
				"Error creating volume snapshot",
				"Cluster not found",
			)
			return
		}
		clusterTime := clusterResponse[0].SystemTime
		if clusterTime != nil {
			name = clusterTime.Format("2006-01-02T15:04:05Z")
		}
	}

	var expirationTimestamp *time.Time
	if !plan.ExpirationTimestamp.IsNull() {
		expTime, _ := plan.ExpirationTimestamp.ValueRFC3339Time()
		if !expTime.IsZero() {
			expirationTimestamp = &expTime
		}
	}

	creatorTypeEnum := clientgen.StorageCreatorTypeEnum(plan.CreatorType.ValueString())

	// Create new volume snapshot
	snapCreateResponse, _, err := r.client.VolumeApi.VolumeSnapshot(ctx, volID).
		Body(clientgen.VolumeSnapshot{
			// Name:                helper.ValueToPointer[string](types.StringValue(name)),
			Name:                &name,
			Description:         helper.ValueToPointer[string](plan.Description),
			PerformancePolicyId: helper.ValueToPointer[string](plan.PerformancePolicyID),
			ExpirationTimestamp: expirationTimestamp,
			CreatorType:         &creatorTypeEnum,
			IsSecure:            helper.ValueToPointer[bool](plan.IsSecure),
		}).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating volume snapshot",
			"Could not create volume snapshot, unexpected error: "+err.Error(),
		)
		return
	}

	// Get snapshot Details using ID retrieved above
	snapshotResponse, err := r.ReadAPI(context.Background(), *snapCreateResponse.Id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting volume snapshot after creation",
			"Could not get volume snapshot, unexpected error: "+err.Error(),
		)
		return
	}

	// Update details to state
	result := models.Snapshot{}
	r.updateSnapshotState(&plan, &result, *snapshotResponse)

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Create")
}

// Read - reads volume snapshot resource information
func (r *resourceVolumeSnapshot) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state models.Snapshot
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	snapshotID := state.ID.ValueString()

	// Get snapshot details from API and then update what is in state from what the API returns
	snapshotResponse, err := r.ReadAPI(context.Background(), snapshotID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading snapshot",
			"Could not read snapshotID with error "+snapshotID+": "+err.Error(),
		)
		return
	}
	r.updateSnapshotState(nil, &state, *snapshotResponse)

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Done with Read")
}

// Update - updates volume snapshot resource
func (r *resourceVolumeSnapshot) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("Started Update")

	//Get plan values
	var plan models.Snapshot
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Get current state
	var state models.Snapshot
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var errFlag bool
	// if volume name is present instead of ID
	if plan.VolumeID.IsUnknown() {
		queries := make(url.Values)
		queries.Set("name", "eq."+plan.VolumeName.ValueString())
		volResponse, _, err := r.client.VolumeApi.GetAllVolumes(ctx).Queries(queries).Execute()
		if err != nil {
			resp.Diagnostics.AddError(
				"Error updating volume snapshot",
				"Could not fetch volume ID from volume name, unexpected error: "+err.Error(),
			)
			return
		}
		if len(volResponse) == 0 {
			resp.Diagnostics.AddError(
				"Error updating volume snapshot",
				"Volume not found",
			)
			return
		}
		if *volResponse[0].Id != state.VolumeID.ValueString() {
			errFlag = true
		}
	}

	if plan.VolumeID.ValueString() != "" && (plan.VolumeID.ValueString() != state.VolumeID.ValueString()) {
		errFlag = true
	}
	if errFlag {
		resp.Diagnostics.AddError(
			"Error updating volume snapshot resource",
			"Volume Name or Volume ID cannot be updated")
		return
	}

	var expirationTimestamp *time.Time
	if !plan.ExpirationTimestamp.IsNull() {
		expTime, _ := plan.ExpirationTimestamp.ValueRFC3339Time()
		if !expTime.IsZero() {
			expirationTimestamp = &expTime
		}
	}

	//Get volume snapshot ID from state
	volumeSnapshotID := state.ID.ValueString()

	//Update volume snapshot by calling API
	_, err := r.client.VolumeApi.PatchVolumeById(ctx, volumeSnapshotID).
		Body(clientgen.VolumeModify{
			Name:                helper.ValueToPointer[string](plan.Name),
			Description:         helper.ValueToPointer[string](plan.Description),
			PerformancePolicyId: helper.ValueToPointer[string](plan.PerformancePolicyID),
			ExpirationTimestamp: expirationTimestamp,
			IsSecure:            helper.ValueToPointer[bool](plan.IsSecure),
		}).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating volume snapshot resource",
			"Could not update volume snapshot "+volumeSnapshotID+": "+err.Error(),
		)
		return
	}

	//Get Volume Snapshot details
	getRes, err := r.ReadAPI(context.Background(), volumeSnapshotID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting snapshot resource after update",
			"Could not get volume snapshot, unexpected error: "+err.Error(),
		)
		return
	}

	r.updateSnapshotState(&plan, &state, *getRes)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Successfully done with Update")
}

// Delete - method to delete volume snapshot resource
func (r *resourceVolumeSnapshot) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("Started with Delete")

	var state models.Snapshot
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get snapshot ID from state
	snapshotID := state.ID.ValueString()

	// Delete snapshot by calling API
	_, err := r.client.VolumeApi.DeleteVolumeById(ctx, snapshotID).Body(clientgen.VolumeDelete{}).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting snapshot",
			"Could not delete snapshotID "+snapshotID+": "+err.Error(),
		)
		return
	}

	log.Printf("Done with Delete")
}

// ImportState - imports state for existing snapshot
func (r *resourceVolumeSnapshot) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r resourceVolumeSnapshot) ReadAPI(ctx context.Context, id string) (*clientgen.VolumeInstance, error) {
	sel := "id,name,description,protection_data,performance_policy_id"
	queries := make(url.Values)
	queries.Set("select", sel)
	response, _, err := r.client.VolumeApi.GetVolumeById(context.Background(), id).Queries(queries).Execute()
	return response, err
}

// updateSnapshotState - method to update terraform state
func (r resourceVolumeSnapshot) updateSnapshotState(plan, state *models.Snapshot, response clientgen.VolumeInstance) {

	state.ID = helper.TfString(response.Id)
	state.Name = helper.TfString(response.Name)
	state.Description = helper.TfString(response.Description)
	if response.ProtectionData != nil {
		state.ExpirationTimestamp = timetypes.NewRFC3339TimePointerValue(response.ProtectionData.ExpirationTimestamp)
		state.VolumeID = helper.TfString(response.ProtectionData.ParentId)
		state.IsSecure = helper.TfBool(response.ProtectionData.IsSecure)
	}
	state.PerformancePolicyID = helper.TfString(response.PerformancePolicyId)
	if plan != nil {
		state.VolumeName = plan.VolumeName
		state.CreatorType = plan.CreatorType
	}
}
