package powerstore

import (
	"context"
	"log"
	"terraform-provider-powerstore/models"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type resourceStorageContainerType struct{}

// GetSchema returns the schema for this resource.
func (r resourceStorageContainerType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:                types.StringType,
				Computed:            true,
				Description:         "The unique identifier of the storage container.",
				MarkdownDescription: "The unique identifier of the storage container.",
			},
			"name": {
				Type:                types.StringType,
				Required:            true,
				Description:         "Name for the storage container.",
				MarkdownDescription: "Name for the storage container. This should be unique across all storage containers in the cluster.",
			},
			"quota": {
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Description:         "The total number of bytes that can be provisioned/reserved against this storage container.",
				MarkdownDescription: "The total number of bytes that can be provisioned/reserved against this storage container. A value of 0 means there is no limit. ",
			},
			"storage_protocol": {
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "The storage protocol of Storage Container.",
				MarkdownDescription: "The storage protocol of Storage Container. eg: SCSI, NVMe",
				Validators: []tfsdk.AttributeValidator{
					oneOfStringtValidator{
						acceptableStringValues: []string{
							string(gopowerstore.StorageContainerStorageProtocolEnumNVME),
							string(gopowerstore.StorageContainerStorageProtocolEnumSCSI),
						},
					},
				},
			},
			"high_water_mark": {
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Description:         "The percentage of the quota that can be consumed before an alert is raised.",
				MarkdownDescription: "The percentage of the quota that can be consumed before an alert is raised",
			},
		},
	}, nil
}

// NewResource is a wrapper around provider
func (r resourceStorageContainerType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceStorageContainer{
		p: *(p.(*Pstoreprovider)),
	}, nil
}

type resourceStorageContainer struct {
	p Pstoreprovider
}

func (r resourceStorageContainer) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	if !r.p.configured {
		resp.Diagnostics.AddError(
			"Provider not configured",
			"The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}
	var plan models.StorageContainer

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	storageContainerCreate := &gopowerstore.StorageContainer{
		Name:            plan.Name.Value,
		Quota:           plan.Quota.Value,
		StorageProtocol: gopowerstore.StorageContainerStorageProtocolEnum(plan.StorageProtocol.Value),
		HighWaterMark:   int16(plan.HighWaterMark.Value),
	}

	storageContainerCreateResponse, err := r.p.client.PStoreClient.CreateStorageContainer(context.Background(), storageContainerCreate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Storage Container",
			"Could not create Storage Container, unexpected error: "+err.Error(),
		)
		return
	}

	// Get Storage Container Details using ID retrieved above
	storageContainerResponse, err1 := r.p.client.PStoreClient.GetStorageContainer(context.Background(), storageContainerCreateResponse.ID)
	if err1 != nil {
		resp.Diagnostics.AddError(
			"Error getting Storage Container after creation",
			"Could not get Storage Container, unexpected error: "+err.Error(),
		)
		return
	}

	result := models.StorageContainer{}

	updateStorageContainerState(&result, storageContainerResponse, &plan, "Create")

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Create")
}

// Read resource information
func (r resourceStorageContainer) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {

	var state models.StorageContainer
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get Storage Container details from API and then update what is in state from what the API returns
	storageContainerID := state.ID.Value
	storageContainerResponse, err := r.p.client.PStoreClient.GetStorageContainer(context.Background(), storageContainerID)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading storage container",
			"Could not read storageContainerID with error "+storageContainerID+": "+err.Error(),
		)
		return
	}

	updateStorageContainerState(&state, storageContainerResponse, nil, "Read")

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Done with Read")
}

// Update resource
func (r resourceStorageContainer) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
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
	storageContainerID := state.ID.Value

	// Update storageContainer by calling API
	_, err := r.p.client.PStoreClient.ModifyStorageContainer(
		context.Background(),
		r.updateRequestPayload(plan, state),
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
	storageContainerResponse, err := r.p.client.PStoreClient.GetStorageContainer(context.Background(), storageContainerID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting storage container after update",
			"Could not get storage container after update, unexpected error: "+err.Error(),
		)
		return
	}

	updateStorageContainerState(&state, storageContainerResponse, nil, "Update")

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Successfully done with Update")
}

// Delete resource
func (r resourceStorageContainer) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	log.Printf("Started with Delete")

	var state models.StorageContainer
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get storage container ID from state
	storageContainerID := state.ID.Value

	// Delete Storage Container by calling API
	_, err := r.p.client.PStoreClient.DeleteStorageContainer(context.Background(), storageContainerID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting storage container",
			"Could not delete storageContainerID "+storageContainerID+": "+err.Error(),
		)
		return
	}

	log.Printf("Done with Delete")
}
