package powerstore

import (
	"context"
	"fmt"
	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"log"
	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"
)

// newHostResource returns host new resource instance
func newHostResource() resource.Resource {
	return &resourceHost{}
}

type resourceHost struct {
	client *client.Client
}

// Metadata defines resource interface Metadata method
func (r *resourceHost) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_host"
}

// Schema defines resource interface Schema method
func (r *resourceHost) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "Host resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The unique identifier of the host.",
				MarkdownDescription: "The unique identifier of the host.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Name of the host.",
				MarkdownDescription: "Name of the host. This should be unique across all hosts in the cluster.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Description of the host.",
				MarkdownDescription: "Description of the host.",
			},
			"host_group_id": schema.StringAttribute{
				Computed:            true,
				Description:         "Associated host group, if host is part of host group.",
				MarkdownDescription: "Associated host group, if host is part of host group.",
			},
			"os_type": schema.StringAttribute{
				Required:            true,
				Description:         "Operating system of the host.",
				MarkdownDescription: "Operating system of the host.",
			},
			"initiators": schema.SetNestedAttribute{
				Description:         "Parameters for creating or adding initiators to host.",
				MarkdownDescription: "Parameters for creating or adding initiators to host.",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"port_name": schema.StringAttribute{
							Description:         "IQN name aka address or NQN name for NVMEoF port types.",
							MarkdownDescription: "IQN name aka address or NQN name for NVMEoF port types.",
							Required:            true,
						},
						"port_type": schema.StringAttribute{
							Description:         "Protocol type of the host initiator.",
							MarkdownDescription: "Protocol type of the host initiator.",
							Required:            true,
						},
					},
				},
			},
			"host_connectivity": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Connectivity type for hosts and host groups.",
				MarkdownDescription: "Connectivity type for hosts and host groups.",
			},
		},
	}
}

// Configure - defines configuration for host resource
func (r *resourceHost) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// Create - method to create host resource
func (r *resourceHost) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan models.Host

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// traverse through initiators in plan and store them
	var initiators []gopowerstore.InitiatorCreateModify
	for _, v := range plan.Initiators {
		portName := v.PortName.ValueString()
		portType := gopowerstore.InitiatorProtocolTypeEnum(v.PortType.ValueString())
		sdsIP := gopowerstore.InitiatorCreateModify{
			PortName: &portName,
			PortType: &portType,
		}
		initiators = append(initiators, sdsIP)
	}

	name := plan.Name.ValueString()
	description := plan.Description.ValueString()
	osType := gopowerstore.OSTypeEnum(plan.OsType.ValueString())

	hostCreate := &gopowerstore.HostCreate{
		Name:             &name,
		Description:      &description,
		OsType:           &osType,
		Initiators:       &initiators,
		HostConnectivity: gopowerstore.HostConnectivityEnum(plan.HostConnectivity.ValueString()),
	}

	// Create new host
	hostCreateResponse, err := r.client.PStoreClient.CreateHost(context.Background(), hostCreate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating host",
			"Could not create host, unexpected error: "+err.Error(),
		)
		return
	}

	// Get host Details using ID retrieved above
	hostResponse, err1 := r.client.PStoreClient.GetHost(context.Background(), hostCreateResponse.ID)
	if err1 != nil {
		resp.Diagnostics.AddError(
			"Error getting host after creation",
			"Could not get host, unexpected error: "+err.Error(),
		)
		return
	}

	// Update details to state
	result := models.Host{}

	r.serverToState(&plan, &result, hostResponse, operationCreate)

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Create")
}

// Read - reads host resource information
func (r *resourceHost) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state models.Host
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get host details from API and then update what is in state from what the API returns
	hostID := state.ID.ValueString()
	hostResponse, err := r.client.PStoreClient.GetHost(context.Background(), hostID)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading host",
			"Could not read hostID with error "+hostID+": "+err.Error(),
		)
		return
	}

	r.serverToState(nil, &state, hostResponse, operationRead)

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Done with Read")
}

// Update - updates host resource
func (r *resourceHost) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete - method to delete host resource
func (r *resourceHost) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("Started with Delete")

	var state models.Host
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get host ID from state
	hostID := state.ID.ValueString()

	// Delete host by calling API
	_, err := r.client.PStoreClient.DeleteHost(context.Background(), &gopowerstore.HostDelete{}, hostID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting host",
			"Could not delete hostID "+hostID+": "+err.Error(),
		)
		return
	}

	log.Printf("Done with Delete")
}

// ImportState - imports state for existing host
func (r *resourceHost) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r resourceHost) serverToState(plan, state *models.Host, response gopowerstore.Host, operation operation) {
	state.ID = types.StringValue(response.ID)
	state.Name = types.StringValue(response.Name)
	state.Description = types.StringValue(response.Description)
	state.OsType = types.StringValue(string(response.OsType))
	state.HostConnectivity = types.StringValue(string(response.HostConnectivity))
	initiators := make([]models.InitiatorCreateModify, 0, len(response.Initiators))
	for _, initiator := range response.Initiators {
		initiators = append(initiators, models.InitiatorCreateModify{
			PortName: types.StringValue(initiator.PortName),
			PortType: types.StringValue(string(initiator.PortType)),
		})
	}
	state.Initiators = initiators
	if operation == operationRead {
		state.HostGroupID = types.StringValue(response.HostGroupID)
	}
}
