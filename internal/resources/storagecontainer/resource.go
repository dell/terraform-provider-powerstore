package storagecontainer

import (
	"context"
	"fmt"
	"terraform-provider-powerstore/internal/powerstore"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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
	resp.TypeName = req.ProviderTypeName + "_storagecontainer"
}

func (r *Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "StorageContainer resource",

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
	var plan model

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.PStoreClient.CreateStorageContainer(context.Background(), plan.serverRequest(model{}))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Storage Container",
			fmt.Sprintf("Could not create Storage Container, unexpected error: %s", err.Error()),
		)
		return
	}

	// Get Storage Container Details using ID retrieved above
	serverResponse, err1 := r.client.PStoreClient.GetStorageContainer(context.Background(), res.ID)
	if err1 != nil {
		resp.Diagnostics.AddError(
			"Error getting Storage Container after creation",
			fmt.Sprintf("Could not get Storage Container id: %s, unexpected error: %s", res.ID, err.Error()),
		)
		return
	}

	plan.saveSeverResponse(serverResponse)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state model
	// Read Terraform state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	serverResponse, err := r.client.PStoreClient.GetStorageContainer(context.Background(), state.ID.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading storage container",
			fmt.Sprintf("Could not read storageContainerID: %s, error: %s ", state.ID.ValueString(), err.Error()),
		)
		return
	}

	state.saveSeverResponse(serverResponse)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var plan, state model
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read Terraform state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update storageContainer by calling API
	_, err := r.client.PStoreClient.ModifyStorageContainer(context.Background(), plan.serverRequest(state), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating storageContainer",
			fmt.Sprintf("Could not update storageContainerID %s, error: %s", state.ID.ValueString(), err.Error()),
		)
		return
	}

	// Get StorageContainer Details
	serverResponse, err := r.client.PStoreClient.GetStorageContainer(context.Background(), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting storage container after update",
			fmt.Sprintf("Could not get storage container id: %s after update, unexpected error: %s", state.ID.ValueString(), err.Error()),
		)
		return
	}

	plan.saveSeverResponse(serverResponse)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state model

	// Read Terraform state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete Storage Container by calling API
	_, err := r.client.PStoreClient.DeleteStorageContainer(context.Background(), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting storage container",
			fmt.Sprintf("Could not delete storageContainerID: %s, error: %s", state.ID.ValueString(), err.Error()),
		)
		return
	}
	// todo: instead of returning error , we should check if storage container id really exists on server
	// and if not , we must return success
	// scenerio - changes from outside of terraform
}

func (r *Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	var state model

	// Read Terraform state data into the model
	resp.Diagnostics.Append(resp.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// fetching asked storage container ID's information
	serverResponse, err := r.client.PStoreClient.GetStorageContainer(context.Background(), state.ID.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(
			"Error importing storage container",
			fmt.Sprintf("Could not import storage container ID: %s with error: %s", state.ID.ValueString(), err.Error()),
		)
		return
	}

	state.saveSeverResponse(serverResponse)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
