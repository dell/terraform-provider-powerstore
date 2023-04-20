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
	"regexp"
	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"
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

		MarkdownDescription: "volume snapshot resource",

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
				Description:         "ID of the volume to take snapshot.",
				MarkdownDescription: "ID of the volume to take snapshot.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ExactlyOneOf(path.MatchRoot("volume_name")),
				},
			},
			"volume_name": schema.StringAttribute{
				Optional:            true,
				Description:         "Name of the volume to take snapshot.",
				MarkdownDescription: "Name of the volume to take snapshot.",
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
				Description:         "Performance Policy id of the volume snapshot.",
				MarkdownDescription: "Performance Policy id of the volume snapshot.",
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
				Description:         "Expiration Timestamp of the volume snapshot.",
				MarkdownDescription: "Expiration Timestamp of the volume snapshot.",
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`\b[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}Z\b`),
						"Only UTC (+Z) format is allowed eg: 2023-05-06T09:01:47Z",
					),
				},
			},
			"creator_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Creator Type of the volume snapshot.",
				MarkdownDescription: "Creator Type of the volume snapshot.",
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						"User",
						"System",
						"Scheduler",
					}...),
				},
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

	// if volume name is present instead of ID
	if volID == "" {
		volResponse, err := r.client.PStoreClient.GetVolumeByName(context.Background(), plan.VolumeName.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Error creating volume snapshot",
				"Could not fetch volume ID from volume name, unexpected error: "+err.Error(),
			)
			return
		}
		volID = volResponse.ID
	}

	// Create new volume snapshot
	snapCreate := &gopowerstore.SnapshotCreate{
		Name:                &name,
		Description:         &description,
		PerformancePolicyID: performancePolicyID,
		ExpirationTimestamp: expirationTimestamp,
		CreatorType:         gopowerstore.StorageCreatorTypeEnum(creatorType),
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
		state.ExpirationTimestamp = types.StringNull()
	} else {
		state.ExpirationTimestamp = types.StringValue(expTime[:len(expTime)-6] + "Z")
	}
	state.CreatorType = types.StringValue(response.ProtectionData.CreatorType)
	state.PerformancePolicyID = types.StringValue(response.PerformancePolicyID)
	state.VolumeID = types.StringValue(response.ProtectionData.ParentID)
	if plan != nil {
		state.VolumeName = plan.VolumeName
	}
}
