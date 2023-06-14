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
	"log"
	"regexp"
	"strings"
	client "terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// newSnapshotRuleResource returns snapshotrule new resource instance
func newSnapshotRuleResource() resource.Resource {
	return &resourceSnapshotRule{}
}

type resourceSnapshotRule struct {
	client *client.Client
}

// Metadata defines resource interface Metadata method
func (r *resourceSnapshotRule) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snapshotrule"
}

// Schema defines resource interface Schema method
func (r *resourceSnapshotRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "SnapshotRule resource",

		Attributes: map[string]schema.Attribute{

			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The ID of the snapshot rule.",
				MarkdownDescription: "The ID of the snapshot rule.",
			},

			"name": schema.StringAttribute{
				Required:            true,
				Description:         "The name of the snapshot rule.",
				MarkdownDescription: "The name of the snapshot rule.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"interval": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The interval between snapshots taken by a snapshot rule, mutually exclusive with time_of_day parameter.",
				MarkdownDescription: "The interval between snapshots taken by a snapshot rule.",
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
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
					}...),

					stringvalidator.ConflictsWith(path.Expressions{
						path.MatchRoot("time_of_day"),
						path.MatchRoot("timezone"),
					}...),
				},
			},

			"time_of_day": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The time of the day to take a daily snapshot, with format hh:mm, mutually exclusive with interval parameter.",
				MarkdownDescription: "The time of the day to take a daily snapshot, with format hh:mm.",
				Validators: []validator.String{

					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[0-9]{2}:[0-9]{2}$`),
						"format is hh:mm",
					),

					stringvalidator.AlsoRequires(path.Expressions{
						path.MatchRoot("timezone"),
					}...),
				},
			},

			"timezone": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The time zone identifier for applying the time zone to the time_of_day for a snapshot rule.",
				MarkdownDescription: "The time zone identifier for applying the time zone to the time_of_day for a snapshot rule.",
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
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
					}...),
				},
			},

			"days_of_week": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "The days of the week when the snapshot rule should be applied.",
				MarkdownDescription: "The days of the week when the snapshot rule should be applied.",
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOf([]string{
							string(gopowerstore.DaysOfWeekEnumMonday),
							string(gopowerstore.DaysOfWeekEnumTuesday),
							string(gopowerstore.DaysOfWeekEnumWednesday),
							string(gopowerstore.DaysOfWeekEnumThursday),
							string(gopowerstore.DaysOfWeekEnumFriday),
							string(gopowerstore.DaysOfWeekEnumSaturday),
							string(gopowerstore.DaysOfWeekEnumSunday),
						}...),
					),
				},
			},

			"desired_retention": schema.Int64Attribute{
				Required:            true,
				Description:         "The Desired snapshot retention period in hours to retain snapshots for this time period.",
				MarkdownDescription: "The Desired snapshot retention period in hours to retain snapshots for this time period.",
			},

			"is_replica": schema.BoolAttribute{
				Computed:            true,
				Description:         "Indicates whether this is a replica of a snapshot rule on a remote system.",
				MarkdownDescription: "Indicates whether this is a replica of a snapshot rule on a remote system.",
			},

			"nas_access_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The NAS filesystem snapshot access method for snapshot rule.",
				MarkdownDescription: "The NAS filesystem snapshot access method for snapshot rule.",
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						string(gopowerstore.NASAccessTypeEnumSnapshot),
						string(gopowerstore.NASAccessTypeEnumProtocol),
					}...),
				},
			},

			"is_read_only": schema.BoolAttribute{
				// todo : currently unable to set on server as true
				Optional:            true,
				Computed:            true,
				Description:         "Indicates whether this snapshot rule can be modified.",
				MarkdownDescription: "Indicates whether this snapshot rule can be modified.",
			},

			"managed_by": schema.StringAttribute{
				Computed:            true,
				Description:         "The entity that owns and manages the instance.",
				MarkdownDescription: "The entity that owns and manages the instance.",
			},

			"managed_by_id": schema.StringAttribute{
				Computed:            true,
				Description:         "The unique id of the managing entity.",
				MarkdownDescription: "The unique id of the managing entity.",
			},

			"delete_snaps": schema.BoolAttribute{
				Optional:            true,
				Description:         "Specify whether all snapshots previously created by this snapshot rule should also be deleted when this rule is removed.",
				MarkdownDescription: "Specify whether all snapshots previously created by this snapshot rule should also be deleted when this rule is removed.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

// Configure - defines configuration for snapshot rule resource
func (r *resourceSnapshotRule) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// Create - method to create Snapshot rule resource
func (r *resourceSnapshotRule) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	log.Printf("Started Creating Snapshot Rule")
	var plan models.SnapshotRule

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	snapshotRuleCreate := r.planToServer(plan)

	log.Printf("Calling api to create snapshotrule")

	// Create New SnapshotRule
	// The function returns only ID of the newly created snapshot rule
	createRes, err := r.client.PStoreClient.CreateSnapshotRule(context.Background(), snapshotRuleCreate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating snapshot rule",
			"Could not create snapshot rule, unexpected error: "+err.Error(),
		)
		return
	}

	log.Printf("Calling api to get snapshotrule created info")

	// Get SnapshotRule Details using ID retrieved above
	getRes, err := r.client.PStoreClient.GetSnapshotRule(context.Background(), createRes.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting snapshot rule after creation",
			"Could not get snapshot rule, unexpected error: "+err.Error(),
		)
		return
	}

	state := models.SnapshotRule{}
	r.serverToState(&plan, &state, getRes, operationCreate)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Successfully done with Create")
}

// Read - fetch info about asked snapshot rule
func (r *resourceSnapshotRule) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state models.SnapshotRule
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get snapshot details from API and then update what is in state from what the API returns
	id := state.ID.ValueString()
	response, err := r.client.PStoreClient.GetSnapshotRule(context.Background(), id)

	// todo distnguish whether error is for resource presence, in case resource is not present
	// we should inform it like resource should be created

	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading snapshot rule",
			"Could not read snapshot rule with error "+id+": "+err.Error(),
		)
		return
	}

	// as state is like a plan here, a current state prior to this read operation
	r.serverToState(&state, &state, response, operationRead)

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Done with Read")
}

// Update - updates snapshotRule
func (r *resourceSnapshotRule) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

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

	snapshotRuleUpdate := r.planToServer(plan)

	// Get snapshotRule ID from state
	snapshotRuleID := state.ID.ValueString()

	// Update snapshotRule by calling API
	_, err := r.client.PStoreClient.ModifySnapshotRule(context.Background(), snapshotRuleUpdate, snapshotRuleID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating snapshotRule",
			"Could not update snapshotRuleID "+snapshotRuleID+": "+err.Error(),
		)
		return
	}

	// Get SnapshotRule Details
	getRes, err := r.client.PStoreClient.GetSnapshotRule(context.Background(), snapshotRuleID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting snapshot rule after update",
			"Could not get snapshot rule, unexpected error: "+err.Error(),
		)
		return
	}

	r.serverToState(&plan, &state, getRes, operationUpdate)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Successfully done with Update")
}

// Delete - deletes snapshotRule
func (r *resourceSnapshotRule) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	log.Printf("Started with Delete")

	var state models.SnapshotRule
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get snapshot rule ID from state
	snapshotRuleID := state.ID.ValueString()

	deleteParams := &gopowerstore.SnapshotRuleDelete{}

	if !state.DeleteSnaps.IsUnknown() && !state.DeleteSnaps.IsNull() {
		deleteParams.DeleteSnaps = state.DeleteSnaps.ValueBool()
	}

	// Delete snapshotRule by calling API
	_, err := r.client.PStoreClient.DeleteSnapshotRule(context.Background(), deleteParams, snapshotRuleID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting snapshotRule",
			"Could not delete snapshotRuleID "+snapshotRuleID+": "+err.Error(),
		)
		return
	}
	// todo: instead of returning error , we should check if snapshotRuleID really exists on server
	// and if not , we must return success
	// scenerio - changes from outside of terraform

	log.Printf("Done with Delete")
}

// ImportState import states for existing snapshot rule
func (r *resourceSnapshotRule) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r resourceSnapshotRule) serverToState(plan, state *models.SnapshotRule, response gopowerstore.SnapshotRule, operation operation) {

	state.ID = types.StringValue(response.ID)
	state.Name = types.StringValue(response.Name)
	state.Interval = types.StringValue(string(response.Interval))
	state.TimeOfDay = types.StringValue(response.TimeOfDay)

	// a work-around
	// converting hh:mm:ss to hh:mm in case server returns hh:mm:ss
	// client can not send hh:mm:ss , else will be a server error , so no worry
	if len(strings.Split(response.TimeOfDay, ":")) == 3 {
		state.TimeOfDay = types.StringValue(strings.TrimSuffix(response.TimeOfDay, ":00"))
	}

	// this if-else will be removed once things got fixed on powerstore side
	// as for import we don't have pre plan so let the imported value save in state
	if operation == operationImport {
		state.TimeZone = types.StringValue(string(response.TimeZone))
		state.NASAccessType = types.StringValue(string(response.NASAccessType))
		state.IsReadOnly = types.BoolValue(false)
	} else {
		// a work-around
		// to allow empty string for default values
		// if default value is returned in response, then if empty string in plan
		// update state value as empty string

		if response.TimeZone == gopowerstore.TimeZoneEnumUTC &&
			!plan.TimeZone.IsUnknown() && !plan.TimeZone.IsNull() &&
			strings.TrimSpace(strings.Trim(plan.TimeZone.ValueString(), "\"")) == "" {
			state.TimeZone = types.StringValue(plan.TimeZone.ValueString())
		} else {
			state.TimeZone = types.StringValue(string(response.TimeZone))
		}

		// as per document, snapshot is default ,  but on server protocol is default
		if response.NASAccessType == gopowerstore.NASAccessTypeEnumProtocol &&
			!plan.NASAccessType.IsUnknown() && !plan.NASAccessType.IsNull() &&
			strings.TrimSpace(strings.Trim(plan.NASAccessType.ValueString(), "\"")) == "" {
			state.NASAccessType = types.StringValue(plan.NASAccessType.ValueString())
		} else {
			state.NASAccessType = types.StringValue(string(response.NASAccessType))
		}

		// a work-around, as we cannot set is_read_only as true on server,
		// so always setting it as plan
		if !plan.IsReadOnly.IsUnknown() && !plan.IsReadOnly.IsNull() {
			state.IsReadOnly = types.BoolValue(plan.IsReadOnly.ValueBool())
		} else {
			state.IsReadOnly = types.BoolValue(false)
		}
	}

	attributeList := []attr.Value{}
	for _, day := range response.DaysOfWeek {
		attributeList = append(attributeList, types.StringValue(string(day)))
	}

	state.DaysOfWeek, _ = types.ListValue(types.StringType, attributeList)

	state.DesiredRetention = types.Int64Value(int64(response.DesiredRetention))
	state.IsReplica = types.BoolValue(response.IsReplica)
	state.ManagedBy = types.StringValue(string(response.ManagedBy))
	state.ManagedByID = types.StringValue(string(response.ManagedById))

	if operation != operationRead {
		// we are saving delete_snaps value in state from plan
		// for future deleteion, if required
		state.DeleteSnaps = plan.DeleteSnaps
	}

	// todo, check if still plan and state are not equal
	// mark resources => should be replaced
}

func (r resourceSnapshotRule) planToServer(plan models.SnapshotRule) *gopowerstore.SnapshotRuleCreate {

	snapshotRuleCreate := &gopowerstore.SnapshotRuleCreate{
		Name:             plan.Name.ValueString(),
		Interval:         gopowerstore.SnapshotRuleIntervalEnum(plan.Interval.ValueString()),
		TimeOfDay:        plan.TimeOfDay.ValueString(),
		TimeZone:         gopowerstore.TimeZoneEnum(plan.TimeZone.ValueString()),
		DesiredRetention: int32(plan.DesiredRetention.ValueInt64()),
		NASAccessType:    gopowerstore.NASAccessTypeEnum(plan.NASAccessType.ValueString()),
	}

	if len(plan.DaysOfWeek.Elements()) > 0 {
		snapshotRuleCreate.DaysOfWeek = []gopowerstore.DaysOfWeekEnum{}

		for _, d := range plan.DaysOfWeek.Elements() {
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
