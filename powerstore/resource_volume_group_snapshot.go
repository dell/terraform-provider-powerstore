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
	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// newVGSnapshotResource returns snapshot new resource instance
func newVGSnapshotResource() resource.Resource {
	return &resourceVGSnapshot{}
}

type resourceVGSnapshot struct {
	client *client.Client
}

// Metadata defines resource interface Metadata method
func (r *resourceVGSnapshot) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_volumegroup_snapshot"
}

// Schema defines resource interface Schema method
func (r *resourceVGSnapshot) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "This resource is used to manage the volumegroup snapshot entity of PowerStore Array. We can Create, Update and Delete the volumegroup snapshot using this resource. We can also import an existing host from volumegroup snapshot array.",
		Description:         "This resource is used to manage the volumegroup snapshot entity of PowerStore Array. We can Create, Update and Delete the volumegroup snapshot using this resource. We can also import an existing host from volumegroup snapshot array.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The unique identifier of the volume group snapshot.",
				MarkdownDescription: "The unique identifier of the volume group snapshot.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Name of the volume group snapshot.",
				MarkdownDescription: "Name of the volume group snapshot.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"volume_group_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "ID of the volume group to take snapshot. Conflicts with `volume_group_name`. Cannot be updated.",
				MarkdownDescription: "ID of the volume group to take snapshot. Conflicts with `volume_group_name`. Cannot be updated.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ExactlyOneOf(path.MatchRoot("volume_group_name")),
				},
			},
			"volume_group_name": schema.StringAttribute{
				Optional:            true,
				Description:         "Name of the volume group to take snapshot. Conflicts with `volume_group_id`. Cannot be updated.",
				MarkdownDescription: "Name of the volume group to take snapshot. Conflicts with `volume_group_id`. Cannot be updated.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ExactlyOneOf(path.MatchRoot("volume_group_id")),
				},
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Description of the volume group snapshot.",
				MarkdownDescription: "Description of the volume group snapshot.",
			},
			"expiration_timestamp": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Expiration Timestamp of the volume group snapshot.Only UTC (+Z) format is allowed",
				MarkdownDescription: "Expiration Timestamp of the volume group snapshot.Only UTC (+Z) format is allowed",
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`(^([0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}Z)$|^$)`),
						"Only UTC (+Z) format is allowed eg: 2023-05-06T09:01:47Z",
					),
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

// Configure - defines configuration for volume group snapshot resource
func (r *resourceVGSnapshot) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// Create - method to create volume group snapshot resource
func (r *resourceVGSnapshot) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan models.VolumeGroupSnapshot

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate secure snapshot parameters
	if !validateSecureSnapshotParams(plan.IsSecure, plan.ExpirationTimestamp, &resp.Diagnostics) {
		return
	}

	name := plan.Name.ValueString()
	description := plan.Description.ValueString()
	expirationTimestamp := plan.ExpirationTimestamp.ValueString()

	volGroupID := plan.VolumeGroupID.ValueString()

	// if volume group name is present instead of ID
	if volGroupID == "" {
		volGroupResponse, err := r.client.PStoreClient.GetVolumeGroupByName(context.Background(), plan.VolumeGroupName.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Error creating volume group snapshot",
				"Could not fetch volume group ID from volume group name, unexpected error: "+err.Error(),
			)
			return
		}
		volGroupID = volGroupResponse.ID
		plan.VolumeGroupID = types.StringValue(volGroupID)
	}

	// Create new volume group snapshot
	isSecure := plan.IsSecure.ValueBool()
	vgSnapCreate := &gopowerstore.VolumeGroupSnapshotCreate{
		Name:                name,
		Description:         description,
		ExpirationTimestamp: expirationTimestamp,
		IsSecure:            &isSecure,
	}

	snapCreateResponse, err := r.client.PStoreClient.CreateVolumeGroupSnapshot(context.Background(), volGroupID, vgSnapCreate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating volume group snapshot",
			"Could not create volume group snapshot, unexpected error: "+err.Error(),
		)
		return
	}
	// Get volume group snapshot Details using ID retrieved above
	snapshotResponse, err1 := r.client.PStoreClient.GetVolumeGroupSnapshot(context.Background(), snapCreateResponse.ID)
	if err1 != nil {
		resp.Diagnostics.AddError(
			"Error getting volume group snapshot after creation",
			"Could not get volume group snapshot, unexpected error: "+err1.Error(),
		)
		return
	}

	// Update details to state
	result := models.VolumeGroupSnapshot{}

	r.updateVGSnapshotState(&plan, &result, snapshotResponse)

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Create")
}

// Read - reads volume group snapshot resource information
func (r *resourceVGSnapshot) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state models.VolumeGroupSnapshot
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	snapshotID := state.ID.ValueString()
	// Get snapshot details from API and then update what is in state from what the API returns

	snapshotResponse, err := r.client.PStoreClient.GetVolumeGroupSnapshot(context.Background(), snapshotID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading snapshot",
			"Could not read snapshotID with error "+snapshotID+": "+err.Error(),
		)
		return
	}
	r.updateVGSnapshotState(nil, &state, snapshotResponse)

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Done with Read")
}

// Update - updates volume group snapshot resource
func (r *resourceVGSnapshot) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("Started Update")

	//Get plan values
	var plan models.VolumeGroupSnapshot
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Get current state
	var state models.VolumeGroupSnapshot
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

	volGroupID := plan.VolumeGroupID.ValueString()
	var errFlag bool
	// if volume group name is present instead of ID
	if volGroupID == "" {
		volGroupResponse, err := r.client.PStoreClient.GetVolumeGroupByName(context.Background(), plan.VolumeGroupName.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Error updating volume group snapshot",
				"Could not fetch volume group ID from volume group name, unexpected error: "+err.Error(),
			)
			return
		}
		if volGroupResponse.ID != state.VolumeGroupID.ValueString() {
			errFlag = true
		}
	} else if volGroupID != "" && volGroupID != state.VolumeGroupID.ValueString() {
		errFlag = true
	}
	if errFlag {
		resp.Diagnostics.AddError(
			"Error updating volume group snapshot resource",
			"Volume group Name or Volume group ID cannot be updated")
		return
	}

	volModify := r.planToServer(plan)

	//Get volume group snapshot ID from state
	volumeGroupSnapshotID := state.ID.ValueString()

	//Update volume group snapshot by calling API
	_, err := r.client.PStoreClient.ModifyVolumeGroupSnapshot(context.Background(), volModify, volumeGroupSnapshotID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating volume group snapshot resource",
			"Could not update volume group snapshot "+volumeGroupSnapshotID+": "+err.Error(),
		)
		return
	}

	//Get Volume Snapshot details
	getRes, err := r.client.PStoreClient.GetVolumeGroupSnapshot(context.Background(), volumeGroupSnapshotID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting volume group snapshot resource after update",
			"Could not get volume group snapshot, unexpected error: "+err.Error(),
		)
		return
	}

	r.updateVGSnapshotState(&plan, &state, getRes)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Successfully done with Update")
}

// Delete - method to delete volume group snapshot resource
func (r *resourceVGSnapshot) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("Started with Delete")

	var state models.VolumeGroupSnapshot
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get volume group snapshot ID from state
	snapshotID := state.ID.ValueString()

	var err error
	// Delete volume group snapshot by calling API
	_, err = r.client.PStoreClient.DeleteVolumeGroup(context.Background(), snapshotID)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting snapshot",
			"Could not delete snapshotID "+snapshotID+": "+err.Error(),
		)
		return
	}

	log.Printf("Done with Delete")
}

// ImportState - imports state for existing volume group snapshot
func (r *resourceVGSnapshot) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// updateVGSnapshotState - method to update terraform state
func (r resourceVGSnapshot) updateVGSnapshotState(plan, state *models.VolumeGroupSnapshot, response gopowerstore.VolumeGroup) {

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
	state.VolumeGroupID = types.StringValue(response.ProtectionData.ParentID)
	// Map is_secure from ProtectionData response
	if response.ProtectionData.IsSecure != nil {
		state.IsSecure = types.BoolValue(*response.ProtectionData.IsSecure)
	} else if plan != nil && !plan.IsSecure.IsNull() && !plan.IsSecure.IsUnknown() {
		state.IsSecure = plan.IsSecure
	} else {
		state.IsSecure = types.BoolValue(false)
	}
	if plan != nil {
		state.VolumeGroupName = plan.VolumeGroupName
	}
}

func (r resourceVGSnapshot) planToServer(plan models.VolumeGroupSnapshot) *gopowerstore.VolumeGroupSnapshotModify {
	name := plan.Name.ValueString()
	description := plan.Description.ValueString()
	expirationTimeStamp := plan.ExpirationTimestamp.ValueString()

	isSecure := plan.IsSecure.ValueBool()
	volGroupSnapshotUpdate := &gopowerstore.VolumeGroupSnapshotModify{
		Description:         description,
		Name:                name,
		ExpirationTimestamp: &expirationTimeStamp,
		IsSecure:            &isSecure,
	}
	return volGroupSnapshotUpdate
}
