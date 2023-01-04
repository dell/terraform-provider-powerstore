package protectionpolicy

import (
	"context"
	"fmt"
	"terraform-provider-powerstore/internal/powerstore"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
	resp.TypeName = req.ProviderTypeName + "_protectionpolicy"
}

func (r *Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "ProtectionPolicy resource",

		Attributes: map[string]schema.Attribute{

			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "Unique identifier of the policy.",
				MarkdownDescription: "Unique identifier of the policy.",
			},

			"name": schema.StringAttribute{
				Required:            true,
				Description:         "The name of the protection policy.",
				MarkdownDescription: "The name of the protection policy.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The description of the protection policy.",
				MarkdownDescription: "The description of the protection policy.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"snapshot_rule_ids": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "List of the snapshot rule IDs that are associated with this policy.",
				MarkdownDescription: "List of the snapshot rule IDs that are associated with this policy.",
				Validators: []validator.Set{

					setvalidator.SizeAtLeast(1),

					setvalidator.ValueStringsAre(
						stringvalidator.LengthAtLeast(1),
					),

					setvalidator.AtLeastOneOf(path.Expressions{
						path.MatchRoot("snapshot_rule_names"),
						path.MatchRoot("replication_rule_ids"),
						path.MatchRoot("replication_rule_names"),
					}...),
				},
			},

			"snapshot_rule_names": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "List of the snapshot rule names that are associated with this policy.",
				MarkdownDescription: "List of the snapshot rule names that are associated with this policy.",
				Validators: []validator.Set{

					setvalidator.SizeAtLeast(1),

					setvalidator.ValueStringsAre(
						stringvalidator.LengthAtLeast(1),
					),
				},
			},

			"replication_rule_ids": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "List of the replication rule IDs that are associated with this policy.",
				MarkdownDescription: "List of the replication rule IDs that are associated with this policy.",
				Validators: []validator.Set{

					setvalidator.SizeAtLeast(1),

					setvalidator.ValueStringsAre(
						stringvalidator.LengthAtLeast(1),
					),
				},
			},

			"replication_rule_names": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "List of the replication rule names that are associated with this policy.",
				MarkdownDescription: "List of the replication rule names that are associated with this policy.",
				Validators: []validator.Set{

					setvalidator.SizeAtLeast(1),

					setvalidator.ValueStringsAre(
						stringvalidator.LengthAtLeast(1),
					),
				},
			},

			// todo, once thse fields are added in gopowerstore
			// then these will be picked up

			// "type": schema.StringAttribute{
			// 	Computed:            true,
			// 	Description:         "The type of the protection policy.",
			// 	MarkdownDescription: "The type of the protection policy.",
			// },

			// "managed_by": schema.StringAttribute{
			// 	Computed:            true,
			// 	Description:         "The entity that owns and manages the instance.",
			// 	MarkdownDescription: "The entity that owns and manages the instance.",
			// },

			// "managed_by_id": schema.StringAttribute{
			// 	Computed:            true,
			// 	Description:         "The unique id of the managing entity.",
			// 	MarkdownDescription: "The unique id of the managing entity.",
			// },

			// "is_replica": schema.BoolAttribute{
			// 	Computed:            true,
			// 	Description:         "Indicates if this is a replica of a policy.",
			// 	MarkdownDescription: "Indicates if this is a replica of a policy.",
			// },

			// "is_read_only": schema.BoolAttribute{
			// 	Computed:            true,
			// 	Description:         "Indicates whether this policy can be modified.",
			// 	MarkdownDescription: "Indicates whether this policy can be modified.",
			// },
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

	createRes, err := r.client.PStoreClient.CreateProtectionPolicy(context.Background(), plan.serverRequest())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating protection policy",
			fmt.Sprintf("Could not create protection policy, unexpected error: %s", err.Error()),
		)
		return
	}

	serverResponse, err := r.client.PStoreClient.GetProtectionPolicy(context.Background(), createRes.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting protection policy after creation",
			fmt.Sprintf("Could not get protection policy id : %s, unexpected error: %s", createRes.ID, err.Error()),
		)
		return
	}

	resp.Diagnostics.Append(plan.saveSeverResponse(serverResponse, r.client)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state model
	// Read Terraform state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	serverResponse, err := r.client.PStoreClient.GetProtectionPolicy(context.Background(), state.ID.ValueString())

	// todo distnguish whether error is for resource presence, in case resource is not present
	// we should inform it like resource should be created

	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading protection policy",
			fmt.Sprintf("Could not read protection policy id: %s, error: %s", state.ID.ValueString(), err.Error()),
		)
		return
	}

	resp.Diagnostics.Append(state.saveSeverResponse(serverResponse, r.client)...)
	if resp.Diagnostics.HasError() {
		return
	}

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

	// Update protection policy by calling API
	_, err := r.client.PStoreClient.ModifyProtectionPolicy(context.Background(), plan.serverRequest(), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating protection policy",
			fmt.Sprintf("Could not update protection policy, id: %s, error: %s", state.ID.ValueString(), err.Error()),
		)
		return
	}

	// Get protection policy details by calling API
	serverResponse, err := r.client.PStoreClient.GetProtectionPolicy(context.Background(), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting protection policy after update",
			fmt.Sprintf("Could not get protection policy, id: %s unexpected error: %s", state.ID.ValueString(), err.Error()),
		)
		return
	}

	resp.Diagnostics.Append(plan.saveSeverResponse(serverResponse, r.client)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state model

	// Read Terraform state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete protection policy by calling API
	_, err := r.client.PStoreClient.DeleteProtectionPolicy(context.Background(), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting protection policy",
			fmt.Sprintf("Could not delete protection policy id: %s, error %s", state.ID.ValueString(), err.Error()),
		)
	}
	// todo: instead of returning error , we should check if protectionPolicy really exists on server
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

	// Get protection poolicy details from API and then update what is in state from what the API returns
	serverResponse, err := r.client.PStoreClient.GetProtectionPolicy(context.Background(), state.ID.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading protection policy",
			fmt.Sprintf("Could not import protection policy for id: %s, error %s \n if its name please mention it like name:actual_name instead of actual_name", state.ID.ValueString(), err.Error()),
		)
		return
	}

	resp.Diagnostics.Append(state.saveSeverResponse(serverResponse, r.client)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
