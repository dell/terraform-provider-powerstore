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

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"
)

const (
	// DefaultExpirationTimestamp is the default expiration timestamp when not specified
	DefaultExpirationTimestamp = "1970-01-01T00:00:00.000Z"
	// SpaceDescription is the default description when not specified
	SpaceDescription = " "
)

// newFileSystemSnapshotResource returns snapshot new resource instance
func newFileSystemSnapshotResource() resource.Resource {
	return &resourceFileSystemSnapshot{}
}

type resourceFileSystemSnapshot struct {
	client *client.Client
}

// Metadata defines resource interface Metadata method
func (r *resourceFileSystemSnapshot) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_filesystem_snapshot"
}

// Schema defines resource interface Schema method
func (r *resourceFileSystemSnapshot) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "This resource is used to manage the filesystem snapshot entity of PowerStore Array. We can Create, Update and Delete the filesystem snapshot using this resource. We can also import an existing filesystem snapshot from PowerStore array.",
		Description:         "This resource is used to manage the filesystem snapshot entity of PowerStore Array. We can Create, Update and Delete the filesystem snapshot using this resource. We can also import an existing filesystem snapshot from PowerStore array.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The unique identifier of the filesystem snapshot.",
				MarkdownDescription: "The unique identifier of the filesystem snapshot.",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Name of the filesystem snapshot.The default name of the filesystem snapshot is the date and time when the snapshot is taken.",
				MarkdownDescription: "Name of the filesystem snapshot.The default name of the filesystem snapshot is the date and time when the snapshot is taken.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"filesystem_id": schema.StringAttribute{
				Required:            true,
				Description:         "ID of the filesystem to take snapshot. Cannot be updated.",
				MarkdownDescription: "ID of the filesystem to take snapshot. Cannot be updated.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Description of the filesystem snapshot.",
				MarkdownDescription: "Description of the filesystem snapshot.",
			},
			"expiration_timestamp": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Expiration Timestamp of the filesystem snapshot, if not provided there will no expiration for the snapshot.Only UTC (+Z) format is allowed eg: 2023-05-06T09:01:47Z",
				MarkdownDescription: "Expiration Timestamp of the filesystem snapshot, if not provided there will no expiration for the snapshot.Only UTC (+Z) format is allowed eg: 2023-05-06T09:01:47Z",
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`(^([0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}Z)$|^$)`),
						"Only UTC (+Z) format is allowed eg: 2023-05-06T09:01:47Z",
					),
				},
			},
			"access_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Access type of the filesystem snapshot. Access type can be 'Snapshot' or 'Protocol'.Cannot be updated.",
				MarkdownDescription: "Access type of the filesystem snapshot. Access type can be 'Snapshot' or 'Protocol'. Cannot be updated.",
				Validators: []validator.String{
					stringvalidator.OneOf("Snapshot", "Protocol"),
				},
				Default: stringdefault.StaticString("Snapshot"),
			},
		},
	}
}

// Configure - defines configuration for filesystem snapshot resource
func (r *resourceFileSystemSnapshot) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// Create - method to create filesystem snapshot resource
func (r *resourceFileSystemSnapshot) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan models.FileSystemSnapshot

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	fileSystemID := plan.FileSystemID.ValueString()

	// Create new filesystem snapshot
	snapCreate := &gopowerstore.SnapshotFSCreate{
		Name:                plan.Name.ValueString(),
		Description:         plan.Description.ValueString(),
		ExpirationTimestamp: plan.ExpirationTimestamp.ValueString(),
		AccessType:          plan.AccessType.ValueString(),
	}

	snapCreateResponse, err := r.client.PStoreClient.CreateFsSnapshot(context.Background(), snapCreate, fileSystemID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating filesystem snapshot",
			"Could not create filesystem snapshot, unexpected error: "+err.Error(),
		)
		return
	}
	// Get snapshot Details using ID retrieved above
	snapshotResponse, err1 := r.client.PStoreClient.GetFS(context.Background(), snapCreateResponse.ID)
	if err1 != nil {
		resp.Diagnostics.AddError(
			"Error getting filesystem snapshot after creation",
			"Could not get filesystem snapshot, unexpected error: "+err1.Error(),
		)
		return
	}

	// Update details to state
	result := models.FileSystemSnapshot{}
	r.updateSnapshotState(&plan, &result, snapshotResponse)

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Create")
}

// Read - reads filesystem snapshot resource information
func (r *resourceFileSystemSnapshot) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state models.FileSystemSnapshot
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	snapshotID := state.ID.ValueString()

	// Get snapshot details from API and then update what is in state from what the API returns
	snapshotResponse, err := r.client.PStoreClient.GetFS(context.Background(), snapshotID)
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

// Update - updates filesystem snapshot resource
func (r *resourceFileSystemSnapshot) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("Started Update")

	//Get plan values
	var plan models.FileSystemSnapshot
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Get current state
	var state models.FileSystemSnapshot
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check not modifiable attributes
	if !plan.Name.Equal(state.Name) || !plan.FileSystemID.Equal(state.FileSystemID) || !plan.AccessType.Equal(state.AccessType) {
		resp.Diagnostics.AddError(
			"Error updating filesystem snapshot resource",
			"filesystem snapshot attributes [name, filesystem_id, access_type] are not modifiable",
		)
		return

	}

	snapshotModify := r.planToServer(plan)

	//Get filesystem snapshot ID from state
	filesystemSnapshotID := state.ID.ValueString()

	//Update filesystem snapshot by calling API
	_, err := r.client.PStoreClient.ModifyFS(context.Background(), snapshotModify, filesystemSnapshotID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating filesystem snapshot resource",
			"Could not update filesystem snapshot "+filesystemSnapshotID+": "+err.Error(),
		)
		return
	}

	//Get filesystem Snapshot details
	getRes, err := r.client.PStoreClient.GetFS(context.Background(), filesystemSnapshotID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting snapshot resource after update",
			"Could not get filesystem snapshot, unexpected error: "+err.Error(),
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

// Delete - method to delete filesystem snapshot resource
func (r *resourceFileSystemSnapshot) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("Started with Delete")

	var state models.FileSystemSnapshot
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get snapshot ID from state
	snapshotID := state.ID.ValueString()

	// Delete snapshot by calling API
	_, err := r.client.PStoreClient.DeleteFsSnapshot(context.Background(), snapshotID)

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
func (r *resourceFileSystemSnapshot) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// updateSnapshotState - method to update terraform state
func (r resourceFileSystemSnapshot) updateSnapshotState(_, state *models.FileSystemSnapshot, response gopowerstore.FileSystem) {

	expTime := response.ExpirationTimestamp
	state.ID = types.StringValue(response.ID)
	state.Name = types.StringValue(response.Name)
	state.Description = types.StringValue(response.Description)
	// if expiration timestamp is not present then set to null.
	if expTime == "" {
		state.ExpirationTimestamp = types.StringValue("")
	} else {
		state.ExpirationTimestamp = types.StringValue(expTime[:len(expTime)-6] + "Z")
	}
	state.AccessType = types.StringValue(response.AccessType)
	state.FileSystemID = types.StringValue(response.ParentID)
}

func (r resourceFileSystemSnapshot) planToServer(plan models.FileSystemSnapshot) *gopowerstore.FSModify {
	description := r.getNonEmptyString(plan.Description.ValueString(), SpaceDescription)
	expirationTimeStamp := r.getNonEmptyString(plan.ExpirationTimestamp.ValueString(), DefaultExpirationTimestamp)
	volSnapshotUpdate := &gopowerstore.FSModify{
		Description:         description,
		ExpirationTimestamp: expirationTimeStamp,
	}
	return volSnapshotUpdate
}

func (r resourceFileSystemSnapshot) getNonEmptyString(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
