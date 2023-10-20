/*
Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"log"
	client "terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// newStorageContainerResource returns storagecontainer new resource instance
func newStorageContainerResource() resource.Resource {
	return &resourceStorageContainer{}
}

type resourceStorageContainer struct {
	client *client.Client
}

// Metadata defines resource interface Metadata method
func (r *resourceStorageContainer) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_storagecontainer"
}

// Schema defines resource interface Schema method
func (r *resourceStorageContainer) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "This resource is used to manage the storage container entity of PowerStore Array. We can Create, Update and Delete the storage container using this resource. We can also import an existing storage container from PowerStore array.",
		Description:         "This resource is used to manage the storage container entity of PowerStore Array. We can Create, Update and Delete the storage container using this resource. We can also import an existing storage container from PowerStore array.",

		Attributes: map[string]schema.Attribute{

			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The unique identifier of the storage container.",
				MarkdownDescription: "The unique identifier of the storage container.",
			},

			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Name for the storage container.",
				MarkdownDescription: "Name for the storage container. This should be unique across all storage containers in the cluster.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"quota": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "The total number of bytes that can be provisioned/reserved against this storage container.",
				MarkdownDescription: "The total number of bytes that can be provisioned/reserved against this storage container. A value of 0 means there is no limit. ",
			},

			"storage_protocol": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The storage protocol of Storage Container.",
				MarkdownDescription: "The storage protocol of Storage Container. eg: SCSI, NVMe",
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						string(gopowerstore.StorageContainerStorageProtocolEnumNVME),
						string(gopowerstore.StorageContainerStorageProtocolEnumSCSI),
					}...),
				},
			},

			"high_water_mark": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "The percentage of the quota that can be consumed before an alert is raised.",
				MarkdownDescription: "The percentage of the quota that can be consumed before an alert is raised",
			},
		},
	}
}

// Configure - defines configuration for storage container resource
func (r *resourceStorageContainer) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// Create - method to create storage container resource
func (r *resourceStorageContainer) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan models.StorageContainer

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	storageContainerCreate := &gopowerstore.StorageContainer{
		Name:            plan.Name.ValueString(),
		Quota:           plan.Quota.ValueInt64(),
		StorageProtocol: gopowerstore.StorageContainerStorageProtocolEnum(plan.StorageProtocol.ValueString()),
		HighWaterMark:   int16(plan.HighWaterMark.ValueInt64()),
	}

	storageContainerCreateResponse, err := r.client.PStoreClient.CreateStorageContainer(context.Background(), storageContainerCreate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Storage Container",
			"Could not create Storage Container, unexpected error: "+err.Error(),
		)
		return
	}

	// Get Storage Container Details using ID retrieved above
	storageContainerResponse, err1 := r.client.PStoreClient.GetStorageContainer(context.Background(), storageContainerCreateResponse.ID)
	if err1 != nil {
		resp.Diagnostics.AddError(
			"Error getting Storage Container after creation",
			"Could not get Storage Container, unexpected error: "+err.Error(),
		)
		return
	}

	result := models.StorageContainer{}

	r.serverToState(&plan, &result, storageContainerResponse, operationCreate)

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Create")
}

// Read - reads storage container resource information
func (r *resourceStorageContainer) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state models.StorageContainer
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get Storage Container details from API and then update what is in state from what the API returns
	storageContainerID := state.ID.ValueString()
	storageContainerResponse, err := r.client.PStoreClient.GetStorageContainer(context.Background(), storageContainerID)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading storage container",
			"Could not read storageContainerID with error "+storageContainerID+": "+err.Error(),
		)
		return
	}

	r.serverToState(nil, &state, storageContainerResponse, operationRead)

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Done with Read")
}

// Update - updates storage container resource
func (r *resourceStorageContainer) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("Started Update")

	// Get plan values
	var plan models.StorageContainer
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state models.StorageContainer
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get storageContainer ID from state
	storageContainerID := state.ID.ValueString()

	// Update storageContainer by calling API
	_, err := r.client.PStoreClient.ModifyStorageContainer(
		context.Background(),
		r.planToServer(plan, state),
		storageContainerID,
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating storageContainer",
			"Could not update storageContainerID "+storageContainerID+": "+err.Error(),
		)
		return
	}

	// Get StorageContainer Details
	storageContainerResponse, err := r.client.PStoreClient.GetStorageContainer(context.Background(), storageContainerID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting storage container after update",
			"Could not get storage container after update, unexpected error: "+err.Error(),
		)
		return
	}

	r.serverToState(nil, &state, storageContainerResponse, operationUpdate)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Successfully done with Update")
}

// Delete - method to delete storage container resource
func (r *resourceStorageContainer) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("Started with Delete")

	var state models.StorageContainer
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get storage container ID from state
	storageContainerID := state.ID.ValueString()

	// Delete Storage Container by calling API
	_, err := r.client.PStoreClient.DeleteStorageContainer(context.Background(), storageContainerID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting storage container",
			"Could not delete storageContainerID "+storageContainerID+": "+err.Error(),
		)
		return
	}

	log.Printf("Done with Delete")
}

// ImportState - imports state for existing storage container
func (r *resourceStorageContainer) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r resourceStorageContainer) serverToState(plan, state *models.StorageContainer, response gopowerstore.StorageContainer, operation operation) {
	state.ID = types.StringValue(response.ID)
	state.Name = types.StringValue(response.Name)
	state.Quota = types.Int64Value(response.Quota)
	state.StorageProtocol = types.StringValue(string(response.StorageProtocol))
	state.HighWaterMark = types.Int64Value(int64(response.HighWaterMark))
}

func (r resourceStorageContainer) planToServer(plan, state models.StorageContainer) *gopowerstore.StorageContainer {

	// a workaround
	// currently PowerStore not accepting PATCH call for same values
	// so sending only updated values

	storageContainerUpdate := &gopowerstore.StorageContainer{}

	if plan.Name.ValueString() != state.Name.ValueString() {
		storageContainerUpdate.Name = plan.Name.ValueString()
	}

	if plan.Quota.ValueInt64() != state.Quota.ValueInt64() {
		storageContainerUpdate.Quota = plan.Quota.ValueInt64()
	}

	if plan.StorageProtocol.ValueString() != state.StorageProtocol.ValueString() {
		storageContainerUpdate.StorageProtocol = gopowerstore.StorageContainerStorageProtocolEnum(plan.StorageProtocol.ValueString())
	}

	if plan.HighWaterMark.ValueInt64() != state.HighWaterMark.ValueInt64() {
		storageContainerUpdate.HighWaterMark = int16(plan.HighWaterMark.ValueInt64())
	}

	return storageContainerUpdate
}
