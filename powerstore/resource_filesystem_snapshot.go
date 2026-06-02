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
	"net/url"
	"strings"
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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	client *clientgen.APIClient
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
				Description:         "Expiration Timestamp of the filesystem snapshot, if not provided there will no expiration for the snapshot. To remove the expiration timestamp, specify it as an empty string. Only UTC (+Z) format is allowed eg: 2023-05-06T09:01:47Z",
				MarkdownDescription: "Expiration Timestamp of the filesystem snapshot, if not provided there will no expiration for the snapshot. To remove the expiration timestamp, specify it as an empty string. Only UTC (+Z) format is allowed eg: 2023-05-06T09:01:47Z",
				CustomType:          timetypes.RFC3339Type{},
			},
			"access_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Access type of the filesystem snapshot. Access type can be 'Snapshot' or 'Protocol'. Cannot be updated.",
				MarkdownDescription: "Access type of the filesystem snapshot. Access type can be 'Snapshot' or 'Protocol'. Cannot be updated.",
				Validators: []validator.String{
					stringvalidator.OneOf("Snapshot", "Protocol"),
				},
				Default: stringdefault.StaticString("Snapshot"),
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

// Configure - defines configuration for filesystem snapshot resource
func (r *resourceFileSystemSnapshot) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client := req.ProviderData.(*client.Client)
	r.client = client.GenClient
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

	var expirationTimestamp *time.Time
	if !plan.ExpirationTimestamp.IsNull() {
		expTime, _ := plan.ExpirationTimestamp.ValueRFC3339Time()
		if !expTime.IsZero() {
			expirationTimestamp = &expTime
		}
	}

	// Create new filesystem snapshot
	snapCreateResponse, _, err := r.client.FileSystemApi.FileSystemSnapshot(ctx, fileSystemID).
		Body(clientgen.FileSystemSnapshot{
			Name:                helper.ValueToPointer[string](plan.Name),
			Description:         helper.ValueToPointer[string](plan.Description),
			AccessType:          helper.PointerStringEnum[clientgen.FileSystemSnapshotAccessTypeEnum](plan.AccessType),
			ExpirationTimestamp: expirationTimestamp,
			IsSecure:            helper.ValueToPointer[bool](plan.IsSecure),
		}).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating filesystem snapshot",
			"Could not create filesystem snapshot, unexpected error: "+err.Error(),
		)
		return
	}
	// Get snapshot Details using ID retrieved above
	snapshotResponse, err := r.ReadAPI(context.Background(), *snapCreateResponse.Id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting filesystem snapshot after creation",
			"Could not get filesystem snapshot, unexpected error: "+err.Error(),
		)
		return
	}

	// Update details to state
	result := models.FileSystemSnapshot{}
	r.updateSnapshotState(&plan, &result, *snapshotResponse)

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
	snapshotResponse, err := r.ReadAPI(context.Background(), snapshotID)
	if err != nil {
		// Check if it's a 404 error and snapshot has expiration timestamp
		if strings.Contains(err.Error(), "404") && !state.ExpirationTimestamp.IsNull() {
			clusterTime, clusterErr := helper.DetermineClusterTime(ctx, r.client)
			if clusterErr == nil {
				expTime, _ := state.ExpirationTimestamp.ValueRFC3339Time()
				// Allow 1 minute buffer for API delays
				if expTime.Add(time.Minute).Before(clusterTime) {
					resp.Diagnostics.AddWarning(
						"Filesystem snapshot auto-deleted",
						fmt.Sprintf("Filesystem snapshot %s was automatically deleted after expiration at %s", snapshotID, expTime.Format(time.RFC3339)),
					)
					resp.State.RemoveResource(ctx)
					return
				}
			}
		}

		// Fall through to standard error handling
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

	var expirationTimestamp *time.Time
	if !plan.ExpirationTimestamp.IsNull() {
		expTime, _ := plan.ExpirationTimestamp.ValueRFC3339Time()
		if !expTime.IsZero() {
			expirationTimestamp = &expTime
		}
	}

	//Get filesystem snapshot ID from state
	filesystemSnapshotID := state.ID.ValueString()

	//Update filesystem snapshot by calling API
	_, err := r.client.FileSystemApi.PatchFileSystemById(ctx, filesystemSnapshotID).
		Body(clientgen.FileSystemModify{
			Description:         helper.ValueToPointer[string](plan.Description),
			ExpirationTimestamp: expirationTimestamp,
			IsSecure:            helper.ValueToPointer[bool](plan.IsSecure),
		}).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating filesystem snapshot resource",
			"Could not update filesystem snapshot "+filesystemSnapshotID+": "+err.Error(),
		)
		return
	}

	//Get filesystem Snapshot details
	getRes, err := r.ReadAPI(context.Background(), filesystemSnapshotID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting snapshot resource after update",
			"Could not get filesystem snapshot, unexpected error: "+err.Error(),
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

func (r *resourceFileSystemSnapshot) ReadAPI(ctx context.Context, id string) (*clientgen.FileSystemInstance, error) {
	sel := "id,name,description,expiration_timestamp,access_type,parent_id,is_secure"
	queries := make(url.Values)
	queries.Set("select", sel)
	response, _, err := r.client.FileSystemApi.GetFileSystemById(context.Background(), id).Queries(queries).Execute()
	return response, err
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
	_, err := r.client.FileSystemApi.DeleteFileSystemById(ctx, snapshotID).Execute()

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
func (r resourceFileSystemSnapshot) updateSnapshotState(plan, state *models.FileSystemSnapshot, response clientgen.FileSystemInstance) {

	state.ID = helper.TfString(response.Id)
	state.Name = helper.TfString(response.Name)
	// Handle case where API returns null for description but plan had empty string
	if response.Description == nil && plan != nil && plan.Description.ValueString() == "" {
		state.Description = types.StringValue("")
	} else {
		state.Description = helper.TfString(response.Description)
	}
	if response.ExpirationTimestamp != nil {
		state.ExpirationTimestamp = timetypes.NewRFC3339TimePointerValue(response.ExpirationTimestamp)
	}
	state.AccessType = helper.TfString(response.AccessType)
	state.FileSystemID = helper.TfString(response.ParentId)
	state.IsSecure = helper.TfBool(response.IsSecure)
}
