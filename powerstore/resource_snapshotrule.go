package powerstore

import (
	"context"
	"log"
	"terraform-provider-powerstore/models"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type resourceSnapshotRuleType struct{}

// GetSchema returns the schema for snapshotrule resource.
func (r resourceSnapshotRuleType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:                types.StringType,
				Computed:            true,
				Description:         "The ID of the snapshot rule.",
				MarkdownDescription: "The ID of the snapshot rule.",
			},
			"name": {
				Type:                types.StringType,
				Required:            true,
				Description:         "The name of the snapshot rule.",
				MarkdownDescription: "The name of the snapshot rule.",
			},
			"interval": {
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "The interval between snapshots taken by a snapshot rule, mutually exclusive with time_of_day parameter.",
				MarkdownDescription: "The interval between snapshots taken by a snapshot rule.",
			},
			"time_of_day": {
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "The time of the day to take a daily snapshot, with format hh:mm, mutually exclusive with interval parameter.",
				MarkdownDescription: "The time of the day to take a daily snapshot, with format hh:mm.",
			},
			"timezone": {
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "The time zone identifier for applying the time zone to the time_of_day for a snapshot rule.",
				MarkdownDescription: "The time zone identifier for applying the time zone to the time_of_day for a snapshot rule.",
			},
			"days_of_week": {
				Type:                types.ListType{ElemType: types.StringType},
				Optional:            true,
				Computed:            true,
				Description:         "The days of the week when the snapshot rule should be applied.",
				MarkdownDescription: "The days of the week when the snapshot rule should be applied.",
			},
			"desired_retention": {
				Type:                types.Int64Type,
				Required:            true,
				Description:         "The Desired snapshot retention period in hours to retain snapshots for this time period.",
				MarkdownDescription: "The Desired snapshot retention period in hours to retain snapshots for this time period.",
			},
			"is_replica": {
				Type:                types.BoolType,
				Computed:            true,
				Description:         "Indicates whether this is a replica of a snapshot rule on a remote system.",
				MarkdownDescription: "Indicates whether this is a replica of a snapshot rule on a remote system.",
			},
			"nas_access_type": {
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "The NAS filesystem snapshot access method for snapshot rule.",
				MarkdownDescription: "The NAS filesystem snapshot access method for snapshot rule.",
			},
			"is_read_only": {
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
				Description:         "Indicates whether this snapshot rule can be modified.",
				MarkdownDescription: "Indicates whether this snapshot rule can be modified.",
			},
			"managed_by": {
				Type:                types.StringType,
				Computed:            true,
				Description:         "The entity that owns and manages the instance.",
				MarkdownDescription: "The entity that owns and manages the instance.",
			},
			"managed_by_id": {
				Type:                types.StringType,
				Computed:            true,
				Description:         "The unique id of the managing entity.",
				MarkdownDescription: "The unique id of the managing entity.",
			},
		},
	}, nil
}

// NewResource is a wrapper around provider
func (r resourceSnapshotRuleType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceSnapshotRule{
		p: *(p.(*Pstoreprovider)),
	}, nil
}

type resourceSnapshotRule struct {
	p Pstoreprovider
}

func (r resourceSnapshotRule) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	if !r.p.configured {
		resp.Diagnostics.AddError(
			"Provider not configured",
			"The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}

	log.Printf("Started Creating Snapshot Rule")
	var plan models.SnapshotRule

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	snapshotRuleCreate := &gopowerstore.SnapshotRuleCreate{
		Name:             plan.Name.Value,
		Interval:         gopowerstore.SnapshotRuleIntervalEnum(plan.Interval.Value),
		TimeOfDay:        plan.TimeOfDay.Value,
		TimeZone:         gopowerstore.TimeZoneEnum(plan.TimeZone.Value),
		DesiredRetention: int32(plan.DesiredRetention.Value),
		NASAccessType:    gopowerstore.NASAccessTypeEnum(plan.NASAccessType.Value),
		IsReadOnly:       plan.IsReadOnly.Value,
	}

	if len(plan.DaysOfWeek.Elems) > 0 {
		snapshotRuleCreate.DaysOfWeek = []gopowerstore.DaysOfWeekEnum{}
		for _, d := range plan.DaysOfWeek.Elems {
			snapshotRuleCreate.DaysOfWeek = append(snapshotRuleCreate.DaysOfWeek, gopowerstore.DaysOfWeekEnum(d.String()))
		}
	}

	log.Printf("Calling api to create snapshotrule")

	// Create New SnapshotRule
	// The function returns only ID of the newly created snapshot rule
	createRes, err := r.p.client.PStoreClient.CreateSnapshotRule(context.Background(), snapshotRuleCreate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating snapshot rule",
			"Could not create snapshot rule, unexpected error: "+err.Error(),
		)
		return
	}

	log.Printf("Calling api to get snapshotrule created info")

	// Get SnapshotRule Details using ID retrieved above
	getRes, err := r.p.client.PStoreClient.GetSnapshotRule(context.Background(), createRes.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting snapshot rule after creation",
			"Could not get snapshot rule, unexpected error: "+err.Error(),
		)
		return
	}

	state := models.SnapshotRule{}
	updateSnapshotRuleState(&state, getRes)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Successfully done with Create")
}

// Read fetch info about asked snapshot rule
func (r resourceSnapshotRule) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {

	var state models.SnapshotRule
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get snapshot details from API and then update what is in state from what the API returns
	id := state.ID.Value
	response, err := r.p.client.PStoreClient.GetSnapshotRule(context.Background(), id)

	// todo distnguish whether error is for resource presence, in case resource is not present
	// we should inform it like resource should be created

	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading snapshot rule",
			"Could not read snapshot rule with error "+id+": "+err.Error(),
		)
		return
	}

	updateSnapshotRuleState(&state, response)

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Done with Read")
}

// Update resource
func (r resourceSnapshotRule) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
}

// Delete resource
func (r resourceSnapshotRule) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
}

func updateSnapshotRuleState(state *models.SnapshotRule, response gopowerstore.SnapshotRule) {

	state.ID.Value = response.ID
	state.Name.Value = response.Name
	state.Interval.Value = string(response.Interval)
	state.TimeOfDay.Value = response.TimeOfDay
	state.TimeZone.Value = string(response.TimeZone)

	attributeList := []attr.Value{}
	for _, day := range response.DaysOfWeek {
		attributeList = append(attributeList, types.String{Value: string(day)})
	}

	state.DaysOfWeek = types.List{
		ElemType: types.StringType,
		Elems:    attributeList,
	}

	state.DesiredRetention.Value = int64(response.DesiredRetention)
	state.IsReplica.Value = response.IsReplica
	state.NASAccessType.Value = string(response.NASAccessType)
	state.IsReadOnly.Value = response.IsReadOnly
	state.ManagedBy.Value = string(response.ManagedBy)
	state.ManagedByID.Value = response.ManagedById
}
