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
	"regexp"
	"strings"
	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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
	client *client.Client
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
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`(^([0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}Z)$|^$)`),
						"Only UTC (+Z) format is allowed eg: 2023-05-06T09:01:47Z",
					),
				},
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
				Optional:            true,
				Computed:            true,
				Description:         "Indicates whether this snapshot is secure. Secure snapshots cannot be deleted before their expiration time. This is a one-way lock: once set to true, it cannot be changed back to false.",
				MarkdownDescription: "Indicates whether this snapshot is secure. Secure snapshots cannot be deleted before their expiration time. This is a one-way lock: once set to true, it cannot be changed back to false.",
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
		volResponse, err := r.client.PStoreClient.GetVolumeByName(context.Background(), plan.VolumeName.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Error creating volume snapshot",
				"Could not fetch volume ID from volume name, unexpected error: "+err.Error(),
			)
			return
		}
		volID = volResponse.ID
		plan.VolumeID = types.StringValue(volID)
	}

	// Validate secure snapshot parameters
	if !validateSecureSnapshotParams(plan.IsSecure, plan.ExpirationTimestamp, &resp.Diagnostics) {
		return
	}

	name := plan.Name.ValueString()
	description := plan.Description.ValueString()
	performancePolicyID := plan.PerformancePolicyID.ValueString()
	expirationTimestamp := plan.ExpirationTimestamp.ValueString()
	creatorType := plan.CreatorType.ValueString()

	// If name of the snapshot is not present, the default name of the volume snapshot is the date and time when the snapshot is taken.
	if name == "" {
		cluster, err := r.client.PStoreClient.GetCluster(ctx)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error creating volume snapshot",
				"Could not fetch name of the cluster, unexpected error: "+err.Error(),
			)
			return
		}
		name = cluster.SystemTime
	}

	// Create new volume snapshot
	isSecure := plan.IsSecure.ValueBool()
	snapCreate := &gopowerstore.SnapshotCreate{
		Name:                &name,
		Description:         &description,
		PerformancePolicyID: performancePolicyID,
		ExpirationTimestamp: expirationTimestamp,
		CreatorType:         gopowerstore.StorageCreatorTypeEnum(creatorType),
		IsSecure:            &isSecure,
	}

	snapCreateResponse, err := r.client.PStoreClient.CreateSnapshot(context.Background(), snapCreate, volID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating volume snapshot",
			"Could not create volume snapshot, unexpected error: "+err.Error(),
		)
		return
	}
	// Get snapshot Details using ID retrieved above
	snapshotResponse, err1 := r.client.PStoreClient.GetSnapshot(context.Background(), snapCreateResponse.ID)
	if err1 != nil {
		resp.Diagnostics.AddError(
			"Error getting volume snapshot after creation",
			"Could not get volume snapshot, unexpected error: "+err.Error(),
		)
		return
	}

	// Update details to state
	result := models.Snapshot{}
	r.updateSnapshotState(&plan, &result, snapshotResponse)

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
	snapshotResponse, err := r.client.PStoreClient.GetSnapshot(context.Background(), snapshotID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading snapshot",
			"Could not read snapshotID with error "+snapshotID+": "+err.Error(),
		)
		return
	}
	r.updateSnapshotState(nil, &state, snapshotResponse)

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

	// Validate secure snapshot parameters
	if !validateSecureSnapshotParams(plan.IsSecure, plan.ExpirationTimestamp, &resp.Diagnostics) {
		return
	}

	// Validate one-way lock: cannot change is_secure from true to false
	if !validateOneWayLock(state.IsSecure, plan.IsSecure, &resp.Diagnostics) {
		return
	}

	var errFlag bool
	// if volume name is present instead of ID
	if plan.VolumeID.IsUnknown() {
		volResponse, err := r.client.PStoreClient.GetVolumeByName(context.Background(), plan.VolumeName.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Error updating volume snapshot",
				"Could not fetch volume ID from volume name, unexpected error: "+err.Error(),
			)
			return
		}
		if volResponse.ID != state.VolumeID.ValueString() {
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

	volModify := r.planToServer(plan)

	//Get volume snapshot ID from state
	volumeSnapshotID := state.ID.ValueString()

	//Update volume snapshot by calling API
	_, err := r.client.PStoreClient.ModifyVolume(context.Background(), volModify, volumeSnapshotID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating volume snapshot resource",
			"Could not update volume snapshot "+volumeSnapshotID+": "+err.Error(),
		)
		return
	}

	//Get Volume Snapshot details
	getRes, err := r.client.PStoreClient.GetSnapshot(context.Background(), volumeSnapshotID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting snapshot resource after update",
			"Could not get volume snapshot, unexpected error: "+err.Error(),
		)
		return
	}

	r.updateSnapshotState(&plan, &state, getRes)

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
	_, err := r.client.PStoreClient.DeleteSnapshot(context.Background(), nil, snapshotID)

	if err != nil {
		errMsg := strings.ToLower(err.Error())
		// Match various secure snapshot retention error patterns from the API
		isSecureRetentionErr := strings.Contains(errMsg, "secure") &&
			(strings.Contains(errMsg, "cannot") || strings.Contains(errMsg, "can not")) ||
			strings.Contains(errMsg, "retention")
		if isSecureRetentionErr {
			log.Printf("[WARN] Cannot delete secure snapshot %s (retention period active). Removing from state; it will auto-expire. Error: %v", snapshotID, err)
			resp.State.RemoveResource(ctx)
			return
		}
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

// updateSnapshotState - method to update terraform state
func (r resourceVolumeSnapshot) updateSnapshotState(plan, state *models.Snapshot, response gopowerstore.Volume) {

	expTime := response.ProtectionData.ExpirationTimeStamp
	state.ID = types.StringValue(response.ID)
	state.Name = types.StringValue(response.Name)
	state.Description = types.StringValue(response.Description)
	// if expiration timestamp is not present then set to null.
	if expTime == "" {
		state.ExpirationTimestamp = types.StringValue("")
	} else {
		state.ExpirationTimestamp = types.StringValue(expTime[:len(expTime)-6] + "Z")
	}
	state.VolumeID = types.StringValue(response.ProtectionData.ParentID)
	state.PerformancePolicyID = types.StringValue(response.PerformancePolicyID)
	// Map is_secure from ProtectionData response
	if response.ProtectionData.IsSecure != nil {
		state.IsSecure = types.BoolValue(*response.ProtectionData.IsSecure)
	} else if plan != nil && !plan.IsSecure.IsNull() && !plan.IsSecure.IsUnknown() {
		state.IsSecure = plan.IsSecure
	} else {
		state.IsSecure = types.BoolValue(false)
	}
	if plan != nil {
		state.VolumeName = plan.VolumeName
		state.CreatorType = plan.CreatorType
	}
}

func (r resourceVolumeSnapshot) planToServer(plan models.Snapshot) *gopowerstore.VolumeModify {
	name := plan.Name.ValueString()
	description := plan.Description.ValueString()
	performancePolicyID := plan.PerformancePolicyID.ValueString()
	expirationTimeStamp := plan.ExpirationTimestamp.ValueString()

	isSecure := plan.IsSecure.ValueBool()
	volSnapshotUpdate := &gopowerstore.VolumeModify{
		Description:         description,
		Name:                name,
		PerformancePolicyID: performancePolicyID,
		ExpirationTimestamp: &expirationTimeStamp,
		IsSecure:            &isSecure,
	}
	return volSnapshotUpdate
}

// validateSecureSnapshotParams validates that expiration_timestamp is set when is_secure is true
func validateSecureSnapshotParams(isSecure types.Bool, expirationTimestamp types.String, diags *diag.Diagnostics) bool {
	if !isSecure.IsNull() && !isSecure.IsUnknown() && isSecure.ValueBool() {
		if expirationTimestamp.IsNull() || expirationTimestamp.IsUnknown() || expirationTimestamp.ValueString() == "" {
			diags.AddError(
				"Missing required parameter",
				"Secure snapshots require an expiration timestamp. "+
					"Please provide expiration_timestamp when is_secure is true.",
			)
			return false
		}
	}
	return true
}

// validateOneWayLock validates that is_secure cannot be changed from true to false
func validateOneWayLock(stateIsSecure, planIsSecure types.Bool, diags *diag.Diagnostics) bool {
	if !stateIsSecure.IsNull() && !stateIsSecure.IsUnknown() && stateIsSecure.ValueBool() {
		if !planIsSecure.IsNull() && !planIsSecure.IsUnknown() && !planIsSecure.ValueBool() {
			diags.AddError(
				"Invalid is_secure change",
				"Secure snapshots cannot be unlocked. The is_secure attribute "+
					"is a one-way lock and cannot be changed from true to false.",
			)
			return false
		}
	}
	return true
}
