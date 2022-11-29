package powerstore

import (
	"context"
	"log"
	"strings"
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
				Validators: []tfsdk.AttributeValidator{
					emptyStringtValidator{},
				},
			},
			"interval": {
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "The interval between snapshots taken by a snapshot rule, mutually exclusive with time_of_day parameter.",
				MarkdownDescription: "The interval between snapshots taken by a snapshot rule.",
				Validators: []tfsdk.AttributeValidator{
					oneOfStringtValidator{
						acceptableStringValues: []string{
							string(gopowerstore.SnapshotRuleIntervalEnumFive_Minutes),
							string(gopowerstore.SnapshotRuleIntervalEnumFifteen_Minutes),
							string(gopowerstore.SnapshotRuleIntervalEnumThirty_Minutes),
							string(gopowerstore.SnapshotRuleIntervalEnumOne_Hour),
							string(gopowerstore.SnapshotRuleIntervalEnumTwo_Hours),
							string(gopowerstore.SnapshotRuleIntervalEnumThree_Hours),
							string(gopowerstore.SnapshotRuleIntervalEnumFour_Hours),
							string(gopowerstore.SnapshotRuleIntervalEnumSix_Hours),
							string(gopowerstore.SnapshotRuleIntervalEnumEight_Hours),
							string(gopowerstore.SnapshotRuleIntervalEnumTwelve_Hours),
							string(gopowerstore.SnapshotRuleIntervalEnumOne_Day),
						},
					},
				},
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
				PlanModifiers:       tfsdk.AttributePlanModifiers{
					// DefaultAttribute(types.String{Value: string(gopowerstore.TimeZoneEnumUTC)}),
					// as default cannot be set for the case when interval is specified,  timezone cannot be specified else server's error
				},
				Validators: []tfsdk.AttributeValidator{
					oneOfStringtValidator{
						acceptableStringValues: []string{
							"", // to allow empty string as well for default value
							string(gopowerstore.TimeZoneEnumEtc__GMT_plus_12),
							string(gopowerstore.TimeZoneEnumUS__Samoa),
							string(gopowerstore.TimeZoneEnumEtc__GMT_plus_11),
							string(gopowerstore.TimeZoneEnumAmerica__Atka),
							string(gopowerstore.TimeZoneEnumUS__Hawaii),
							string(gopowerstore.TimeZoneEnumEtc__GMT_plus_10),
							string(gopowerstore.TimeZoneEnumPacific__Marquesas),
							string(gopowerstore.TimeZoneEnumUS__Alaska),
							string(gopowerstore.TimeZoneEnumPacific__Gambier),
							string(gopowerstore.TimeZoneEnumEtc__GMT_plus_9),
							string(gopowerstore.TimeZoneEnumPST8PDT),
							string(gopowerstore.TimeZoneEnumPacific__Pitcairn),
							string(gopowerstore.TimeZoneEnumUS__Pacific),
							string(gopowerstore.TimeZoneEnumEtc__GMT_plus_8),
							string(gopowerstore.TimeZoneEnumMexico__BajaSur),
							string(gopowerstore.TimeZoneEnumAmerica__Boise),
							string(gopowerstore.TimeZoneEnumAmerica__Phoenix),
							string(gopowerstore.TimeZoneEnumMST7MDT),
							string(gopowerstore.TimeZoneEnumEtc__GMT_plus_7),
							string(gopowerstore.TimeZoneEnumCST6CDT),
							string(gopowerstore.TimeZoneEnumAmerica__Chicago),
							string(gopowerstore.TimeZoneEnumCanada__Saskatchewan),
							string(gopowerstore.TimeZoneEnumAmerica__Bahia_Banderas),
							string(gopowerstore.TimeZoneEnumEtc__GMT_plus_6),
							string(gopowerstore.TimeZoneEnumChile__EasterIsland),
							string(gopowerstore.TimeZoneEnumAmerica__Bogota),
							string(gopowerstore.TimeZoneEnumAmerica__New_York),
							string(gopowerstore.TimeZoneEnumEST5EDT),
							string(gopowerstore.TimeZoneEnumAmerica__Havana),
							string(gopowerstore.TimeZoneEnumEtc__GMT_plus_5),
							string(gopowerstore.TimeZoneEnumAmerica__Caracas),
							string(gopowerstore.TimeZoneEnumAmerica__Cuiaba),
							string(gopowerstore.TimeZoneEnumAmerica__Santo_Domingo),
							string(gopowerstore.TimeZoneEnumCanada__Atlantic),
							string(gopowerstore.TimeZoneEnumAmerica__Asuncion),
							string(gopowerstore.TimeZoneEnumEtc__GMT_plus_4),
							string(gopowerstore.TimeZoneEnumCanada__Newfoundland),
							string(gopowerstore.TimeZoneEnumChile__Continental),
							string(gopowerstore.TimeZoneEnumBrazil__East),
							string(gopowerstore.TimeZoneEnumAmerica__Godthab),
							string(gopowerstore.TimeZoneEnumAmerica__Miquelon),
							string(gopowerstore.TimeZoneEnumAmerica__Buenos_Aires),
							string(gopowerstore.TimeZoneEnumEtc__GMT_plus_3),
							string(gopowerstore.TimeZoneEnumAmerica__Noronha),
							string(gopowerstore.TimeZoneEnumEtc__GMT_plus_2),
							string(gopowerstore.TimeZoneEnumAmerica__Scoresbysund),
							string(gopowerstore.TimeZoneEnumAtlantic__Cape_Verde),
							string(gopowerstore.TimeZoneEnumEtc__GMT_plus_1),
							string(gopowerstore.TimeZoneEnumUTC),
							string(gopowerstore.TimeZoneEnumEurope__London),
							string(gopowerstore.TimeZoneEnumAfrica__Casablanca),
							string(gopowerstore.TimeZoneEnumAtlantic__Reykjavik),
							string(gopowerstore.TimeZoneEnumAntarctica__Troll),
							string(gopowerstore.TimeZoneEnumEurope__Paris),
							string(gopowerstore.TimeZoneEnumEurope__Sarajevo),
							string(gopowerstore.TimeZoneEnumEurope__Belgrade),
							string(gopowerstore.TimeZoneEnumEurope__Rome),
							string(gopowerstore.TimeZoneEnumAfrica__Tunis),
							string(gopowerstore.TimeZoneEnumEtc__GMT_minus_1),
							string(gopowerstore.TimeZoneEnumAsia__Gaza),
							string(gopowerstore.TimeZoneEnumEurope__Bucharest),
							string(gopowerstore.TimeZoneEnumEurope__Helsinki),
							string(gopowerstore.TimeZoneEnumAsia__Beirut),
							string(gopowerstore.TimeZoneEnumAfrica__Harare),
							string(gopowerstore.TimeZoneEnumAsia__Damascus),
							string(gopowerstore.TimeZoneEnumAsia__Amman),
							string(gopowerstore.TimeZoneEnumEurope__Tiraspol),
							string(gopowerstore.TimeZoneEnumAsia__Jerusalem),
							string(gopowerstore.TimeZoneEnumEtc__GMT_minus_2),
							string(gopowerstore.TimeZoneEnumAsia__Baghdad),
							string(gopowerstore.TimeZoneEnumAfrica__Asmera),
							string(gopowerstore.TimeZoneEnumEtc__GMT_minus_3),
							string(gopowerstore.TimeZoneEnumAsia__Tehran),
							string(gopowerstore.TimeZoneEnumAsia__Baku),
							string(gopowerstore.TimeZoneEnumEtc__GMT_minus_4),
							string(gopowerstore.TimeZoneEnumAsia__Kabul),
							string(gopowerstore.TimeZoneEnumAsia__Karachi),
							string(gopowerstore.TimeZoneEnumEtc__GMT_minus_5),
							string(gopowerstore.TimeZoneEnumAsia__Kolkata),
							string(gopowerstore.TimeZoneEnumAsia__Katmandu),
							string(gopowerstore.TimeZoneEnumAsia__Almaty),
							string(gopowerstore.TimeZoneEnumEtc__GMT_minus_6),
							string(gopowerstore.TimeZoneEnumAsia__Rangoon),
							string(gopowerstore.TimeZoneEnumAsia__Hovd),
							string(gopowerstore.TimeZoneEnumAsia__Bangkok),
							string(gopowerstore.TimeZoneEnumEtc__GMT_minus_7),
							string(gopowerstore.TimeZoneEnumAsia__Hong_Kong),
							string(gopowerstore.TimeZoneEnumAsia__Brunei),
							string(gopowerstore.TimeZoneEnumAsia__Singapore),
							string(gopowerstore.TimeZoneEnumEtc__GMT_minus_8),
							string(gopowerstore.TimeZoneEnumAsia__Pyongyang),
							string(gopowerstore.TimeZoneEnumAustralia__Eucla),
							string(gopowerstore.TimeZoneEnumAsia__Seoul),
							string(gopowerstore.TimeZoneEnumEtc__GMT_minus_9),
							string(gopowerstore.TimeZoneEnumAustralia__Darwin),
							string(gopowerstore.TimeZoneEnumAustralia__Adelaide),
							string(gopowerstore.TimeZoneEnumAustralia__Sydney),
							string(gopowerstore.TimeZoneEnumAustralia__Brisbane),
							string(gopowerstore.TimeZoneEnumAsia__Magadan),
							string(gopowerstore.TimeZoneEnumEtc__GMT_minus_10),
							string(gopowerstore.TimeZoneEnumAustralia__Lord_Howe),
							string(gopowerstore.TimeZoneEnumEtc__GMT_minus_11),
							string(gopowerstore.TimeZoneEnumAsia__Kamchatka),
							string(gopowerstore.TimeZoneEnumPacific__Fiji),
							string(gopowerstore.TimeZoneEnumAntarctica__South_Pole),
							string(gopowerstore.TimeZoneEnumEtc__GMT_minus_12),
							string(gopowerstore.TimeZoneEnumPacific__Chatham),
							string(gopowerstore.TimeZoneEnumPacific__Tongatapu),
							string(gopowerstore.TimeZoneEnumPacific__Apia),
							string(gopowerstore.TimeZoneEnumEtc__GMT_minus_13),
							string(gopowerstore.TimeZoneEnumPacific__Kiritimati),
							string(gopowerstore.TimeZoneEnumEtc__GMT_minus_14),
						},
					},
				},
			},
			"days_of_week": {
				Type:                types.ListType{ElemType: types.StringType},
				Optional:            true,
				Computed:            true,
				Description:         "The days of the week when the snapshot rule should be applied.",
				MarkdownDescription: "The days of the week when the snapshot rule should be applied.",
				Validators: []tfsdk.AttributeValidator{
					oneOfStringtValidator{
						isList: true,
						acceptableStringValues: []string{
							string(gopowerstore.DaysOfWeekEnumMonday),
							string(gopowerstore.DaysOfWeekEnumTuesday),
							string(gopowerstore.DaysOfWeekEnumWednesday),
							string(gopowerstore.DaysOfWeekEnumThursday),
							string(gopowerstore.DaysOfWeekEnumFriday),
							string(gopowerstore.DaysOfWeekEnumSaturday),
							string(gopowerstore.DaysOfWeekEnumSunday),
						},
					},
				},
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
				PlanModifiers: tfsdk.AttributePlanModifiers{
					DefaultAttribute(types.String{Value: string(gopowerstore.NASAccessTypeEnumSnapshot)}),
				},
				Validators: []tfsdk.AttributeValidator{
					oneOfStringtValidator{
						acceptableStringValues: []string{
							"", // to allow empty value to have default value
							string(gopowerstore.NASAccessTypeEnumSnapshot),
							string(gopowerstore.NASAccessTypeEnumProtocol),
						},
					},
				},
			},
			// todo : currently unable to set on server as true
			"is_read_only": {
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
				Description:         "Indicates whether this snapshot rule can be modified.",
				MarkdownDescription: "Indicates whether this snapshot rule can be modified.",
				PlanModifiers: tfsdk.AttributePlanModifiers{
					DefaultAttribute(types.Bool{Value: false}),
				},
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

	snapshotRuleCreate := planToSnapshotRuleParam(plan)

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
	updateSnapshotRuleState(&plan, &state, getRes)

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

	// as stats is like a plan here, a current state prior to this read operation
	updateSnapshotRuleState(&state, &state, response)

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Done with Read")
}

// Update updates snapshotRule
func (r resourceSnapshotRule) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {

	log.Printf("Started Update")

	// Get plan values
	var plan models.SnapshotRule
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state models.SnapshotRule
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	snapshotRuleUpdate := planToSnapshotRuleParam(plan)

	// Get snapshotRule ID from state
	snapshotRuleID := state.ID.Value

	// Update snapshotRule by calling API
	_, err := r.p.client.PStoreClient.ModifySnapshotRule(context.Background(), snapshotRuleUpdate, snapshotRuleID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating snapshotRule",
			"Could not update snapshotRuleID "+snapshotRuleID+": "+err.Error(),
		)
		return
	}

	// Get SnapshotRule Details
	getRes, err := r.p.client.PStoreClient.GetSnapshotRule(context.Background(), snapshotRuleID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting snapshot rule after update",
			"Could not get snapshot rule, unexpected error: "+err.Error(),
		)
		return
	}

	updateSnapshotRuleState(&plan, &state, getRes)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Successfully done with Update")
}

// Delete deletes snapshotRule
func (r resourceSnapshotRule) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {

	log.Printf("Started with Delete")

	var state models.SnapshotRule
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get snapshot rule ID from state
	snapshotRuleID := state.ID.Value

	// Delete snapshotRule by calling API
	_, err := r.p.client.PStoreClient.DeleteSnapshotRule(context.Background(), nil, snapshotRuleID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting snapshotRule",
			"Could not delete snapshotRuleID "+snapshotRuleID+": "+err.Error(),
		)
		return
	}

	log.Printf("Done with Delete")
}

func updateSnapshotRuleState(plan, state *models.SnapshotRule, response gopowerstore.SnapshotRule) {

	state.ID.Value = response.ID
	state.Name.Value = response.Name
	state.Interval.Value = string(response.Interval)
	state.TimeOfDay.Value = response.TimeOfDay

	// a work-around
	// converting hh:mm:ss to hh:mm in case server returns hh:mm:ss
	// client can not send hh:mm:ss , else will be a server error , so no worry
	if len(strings.Split(response.TimeOfDay, ":")) == 3 {
		state.TimeOfDay.Value = strings.TrimSuffix(response.TimeOfDay, ":00")
	}

	// a work-around
	// to allow empty string for default values
	// if default value is returned in response, then if empty string in plan
	// update state value as empty string
	{

		// again a hard coded value , bruh :(
		if response.TimeZone == gopowerstore.TimeZoneEnumUTC &&
			!plan.TimeZone.IsUnknown() && !plan.TimeZone.IsNull() &&
			strings.TrimSpace(strings.Trim(plan.TimeZone.String(), "\"")) == "" {
			state.TimeZone.Value = plan.TimeZone.Value
		} else {
			state.TimeZone.Value = string(response.TimeZone)
		}

		// again a hard coded value , bruh :(
		// as per document, snapshot is default ,  but on server protocol is default
		if response.NASAccessType == gopowerstore.NASAccessTypeEnumProtocol &&
			!plan.NASAccessType.IsUnknown() && !plan.NASAccessType.IsNull() &&
			strings.TrimSpace(strings.Trim(plan.NASAccessType.String(), "\"")) == "" {
			state.NASAccessType.Value = plan.NASAccessType.Value
		} else {
			state.NASAccessType.Value = string(response.NASAccessType)
		}

		// a work-around, as we cannot set is_read_only as true on server,
		// so always setting it as plan
		if !plan.IsReadOnly.IsUnknown() && !plan.IsReadOnly.IsNull() {
			state.IsReadOnly.Value = plan.IsReadOnly.Value
		} else {
			state.IsReadOnly.Value = false
		}

	}

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
	state.ManagedBy.Value = string(response.ManagedBy)
	state.ManagedByID.Value = response.ManagedById

	// todo, check if still plan and state are not equal
	// mark resources => should be replaced
}

func planToSnapshotRuleParam(plan models.SnapshotRule) *gopowerstore.SnapshotRuleCreate {

	snapshotRuleCreate := &gopowerstore.SnapshotRuleCreate{
		Name:             plan.Name.Value,
		Interval:         gopowerstore.SnapshotRuleIntervalEnum(plan.Interval.Value),
		TimeOfDay:        plan.TimeOfDay.Value,
		TimeZone:         gopowerstore.TimeZoneEnum(plan.TimeZone.Value),
		DesiredRetention: int32(plan.DesiredRetention.Value),
		NASAccessType:    gopowerstore.NASAccessTypeEnum(plan.NASAccessType.Value),
	}

	if len(plan.DaysOfWeek.Elems) > 0 {
		snapshotRuleCreate.DaysOfWeek = []gopowerstore.DaysOfWeekEnum{}

		for _, d := range plan.DaysOfWeek.Elems {
			snapshotRuleCreate.DaysOfWeek = append(
				snapshotRuleCreate.DaysOfWeek,
				gopowerstore.DaysOfWeekEnum(
					strings.Trim(d.String(), "\""),
				),
			)
		}
	}

	return snapshotRuleCreate
}
