/*
Copyright (c) 2024-2026 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	"log"
	"net/url"
	"regexp"
	"strings"
	client "terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/models"
	"terraform-provider-powerstore/powerstore/helper"

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
	client *clientgen.APIClient
}

// Metadata defines resource interface Metadata method
func (r *resourceSnapshotRule) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snapshotrule"
}

// Schema defines resource interface Schema method
func (r *resourceSnapshotRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "This resource is used to manage the snapshot rule entity of PowerStore Array. We can Create, Update and Delete the snapshot rule using this resource. We can also import an existing snapshot rule from PowerStore array.",
		Description:         "This resource is used to manage the snapshot rule entity of PowerStore Array. We can Create, Update and Delete the snapshot rule using this resource. We can also import an existing snapshot rule from PowerStore array.",

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
						"Five_Minutes",
						"Fifteen_Minutes",
						"Thirty_Minutes",
						"One_Hour",
						"Two_Hours",
						"Three_Hours",
						"Four_Hours",
						"Six_Hours",
						"Eight_Hours",
						"Twelve_Hours",
						"One_Day",
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
						"Etc__GMT_plus_12",
						"US__Samoa",
						"Etc__GMT_plus_11",
						"America__Atka",
						"US__Hawaii",
						"Etc__GMT_plus_10",
						"Pacific__Marquesas",
						"US__Alaska",
						"Pacific__Gambier",
						"Etc__GMT_plus_9",
						"PST8PDT",
						"Pacific__Pitcairn",
						"US__Pacific",
						"Etc__GMT_plus_8",
						"Mexico__BajaSur",
						"America__Boise",
						"America__Phoenix",
						"MST7MDT",
						"Etc__GMT_plus_7",
						"CST6CDT",
						"America__Chicago",
						"Canada__Saskatchewan",
						"America__Bahia_Banderas",
						"Etc__GMT_plus_6",
						"Chile__EasterIsland",
						"America__Bogota",
						"America__New_York",
						"EST5EDT",
						"America__Havana",
						"Etc__GMT_plus_5",
						"America__Caracas",
						"America__Cuiaba",
						"America__Santo_Domingo",
						"Canada__Atlantic",
						"America__Asuncion",
						"Etc__GMT_plus_4",
						"Canada__Newfoundland",
						"Chile__Continental",
						"Brazil__East",
						"America__Godthab",
						"America__Miquelon",
						"America__Buenos_Aires",
						"Etc__GMT_plus_3",
						"America__Noronha",
						"Etc__GMT_plus_2",
						"America__Scoresbysund",
						"Atlantic__Cape_Verde",
						"Etc__GMT_plus_1",
						"UTC",
						"Europe__London",
						"Africa__Casablanca",
						"Atlantic__Reykjavik",
						"Antarctica__Troll",
						"Europe__Paris",
						"Europe__Sarajevo",
						"Europe__Belgrade",
						"Europe__Rome",
						"Africa__Tunis",
						"Etc__GMT_minus_1",
						"Asia__Gaza",
						"Europe__Bucharest",
						"Europe__Helsinki",
						"Asia__Beirut",
						"Africa__Harare",
						"Asia__Damascus",
						"Asia__Amman",
						"Europe__Tiraspol",
						"Asia__Jerusalem",
						"Etc__GMT_minus_2",
						"Asia__Baghdad",
						"Africa__Asmera",
						"Etc__GMT_minus_3",
						"Asia__Tehran",
						"Asia__Baku",
						"Etc__GMT_minus_4",
						"Asia__Kabul",
						"Asia__Karachi",
						"Etc__GMT_minus_5",
						"Asia__Kolkata",
						"Asia__Katmandu",
						"Asia__Almaty",
						"Etc__GMT_minus_6",
						"Asia__Rangoon",
						"Asia__Hovd",
						"Asia__Bangkok",
						"Etc__GMT_minus_7",
						"Asia__Hong_Kong",
						"Asia__Brunei",
						"Asia__Singapore",
						"Etc__GMT_minus_8",
						"Asia__Pyongyang",
						"Australia__Eucla",
						"Asia__Seoul",
						"Etc__GMT_minus_9",
						"Australia__Darwin",
						"Australia__Adelaide",
						"Australia__Sydney",
						"Australia__Brisbane",
						"Asia__Magadan",
						"Etc__GMT_minus_10",
						"Australia__Lord_Howe",
						"Etc__GMT_minus_11",
						"Asia__Kamchatka",
						"Pacific__Fiji",
						"Antarctica__South_Pole",
						"Etc__GMT_minus_12",
						"Pacific__Chatham",
						"Pacific__Tongatapu",
						"Pacific__Apia",
						"Etc__GMT_minus_13",
						"Pacific__Kiritimati",
						"Etc__GMT_minus_14",
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
							"Monday",
							"Tuesday",
							"Wednesday",
							"Thursday",
							"Friday",
							"Saturday",
							"Sunday",
						}...),
					),
				},
			},

			"desired_retention": schema.Int32Attribute{
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
						"Snapshot",
						"Protocol",
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
			"is_secure": schema.BoolAttribute{
				Description:         "Indicates whether snapshots created by this rule should be secure. Secure snapshots cannot be deleted before the expiration time, and the expiration time cannot be reduced.",
				MarkdownDescription: "Indicates whether snapshots created by this rule should be secure. Secure snapshots cannot be deleted before the expiration time, and the expiration time cannot be reduced.",
				Optional:            true,
				Computed:            true,
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

	client := req.ProviderData.(*client.Client)
	r.client = client.GenClient
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

	var daysOfWeek []clientgen.DaysOfWeekEnum
	if len(plan.DaysOfWeek.Elements()) > 0 {
		daysOfWeek = []clientgen.DaysOfWeekEnum{}
		for _, d := range plan.DaysOfWeek.Elements() {
			daysOfWeek = append(
				daysOfWeek,
				clientgen.DaysOfWeekEnum(strings.Trim(d.String(), "\"")),
			)
		}
	}

	log.Printf("Calling api to create snapshotrule")

	// Create New SnapshotRule
	// The function returns only ID of the newly created snapshot rule
	createRes, _, err := r.client.SnapshotRuleApi.PostAllSnapshotRules(ctx).
		Body(clientgen.SnapshotRuleCreate{
			Name:             plan.Name.ValueString(),
			DesiredRetention: plan.DesiredRetention.ValueInt32(),
			Interval:         helper.PointerStringEnum[clientgen.SnapRuleIntervalEnum](plan.Interval),
			TimeOfDay:        helper.ValueToPointer[string](plan.TimeOfDay),
			Timezone:         helper.PointerStringEnum[clientgen.TimeZoneEnum](plan.TimeZone),
			NasAccessType:    helper.PointerStringEnum[clientgen.NASAccessTypeEnum](plan.NASAccessType),
			IsReadOnly:       helper.ValueToPointer[bool](plan.IsReadOnly),
			IsSecure:         helper.ValueToPointer[bool](plan.IsSecure),
			DaysOfWeek:       daysOfWeek,
		}).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating snapshot rule",
			"Could not create snapshot rule, unexpected error: "+err.Error(),
		)
		return
	}

	log.Printf("Calling api to get snapshotrule created info")

	// Get SnapshotRule Details using ID retrieved above
	getRes, err := r.ReadAPI(context.Background(), *createRes.Id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting snapshot rule after creation",
			"Could not get snapshot rule, unexpected error: "+err.Error(),
		)
		return
	}

	state := models.SnapshotRule{}
	r.serverToState(&plan, &state, *getRes, operationCreate)

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
	response, err := r.ReadAPI(context.Background(), id)

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
	r.serverToState(&state, &state, *response, operationRead)

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

	var daysOfWeek []clientgen.DaysOfWeekEnum
	if len(plan.DaysOfWeek.Elements()) > 0 {
		daysOfWeek = []clientgen.DaysOfWeekEnum{}
		for _, d := range plan.DaysOfWeek.Elements() {
			daysOfWeek = append(
				daysOfWeek,
				clientgen.DaysOfWeekEnum(strings.Trim(d.String(), "\"")),
			)
		}
	}

	// Get snapshotRule ID from state
	snapshotRuleID := state.ID.ValueString()

	// Update snapshotRule by calling API
	_, err := r.client.SnapshotRuleApi.PatchSnapshotRuleById(ctx, snapshotRuleID).
		Body(clientgen.SnapshotRuleModify{
			Name:             helper.ValueToPointer[string](plan.Name),
			Interval:         helper.PointerStringEnum[clientgen.SnapRuleIntervalEnum](plan.Interval),
			TimeOfDay:        helper.ValueToPointer[string](plan.TimeOfDay),
			Timezone:         helper.PointerStringEnum[clientgen.TimeZoneEnum](plan.TimeZone),
			NasAccessType:    helper.PointerStringEnum[clientgen.NASAccessTypeEnum](plan.NASAccessType),
			IsSecure:         helper.ValueToPointer[bool](plan.IsSecure),
			DesiredRetention: helper.ValueToPointer[int32](plan.DesiredRetention),
			DaysOfWeek:       daysOfWeek,
		}).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating snapshotRule",
			"Could not update snapshotRuleID "+snapshotRuleID+": "+err.Error(),
		)
		return
	}

	// Get SnapshotRule Details
	getRes, err := r.ReadAPI(context.Background(), snapshotRuleID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting snapshot rule after update",
			"Could not get snapshot rule, unexpected error: "+err.Error(),
		)
		return
	}

	r.serverToState(&plan, &state, *getRes, operationUpdate)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Successfully done with Update")
}

func (r *resourceSnapshotRule) ReadAPI(ctx context.Context, id string) (*clientgen.SnapshotRuleInstance, error) {
	sel := "id,name,remote_system_id,interval,time_of_day,timezone,days_of_week,desired_retention,is_replica,nas_access_type,is_read_only,managed_by,managed_by_id,is_secure"
	queries := make(url.Values)
	queries.Set("select", sel)
	response, _, err := r.client.SnapshotRuleApi.GetSnapshotRuleById(context.Background(), id).Queries(queries).Execute()
	return response, err
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

	deleteParams := clientgen.SnapshotRuleDelete{
		DeleteSnaps: helper.ValueToPointer[bool](state.DeleteSnaps),
	}

	// Delete snapshotRule by calling API
	_, err := r.client.SnapshotRuleApi.DeleteSnapshotRuleById(ctx, snapshotRuleID).Body(deleteParams).Execute()
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

func (r resourceSnapshotRule) serverToState(plan, state *models.SnapshotRule, response clientgen.SnapshotRuleInstance, operation operation) {

	state.ID = helper.TfString(response.Id)
	state.Name = helper.TfString(response.Name)
	state.Interval = helper.TfString(response.Interval)
	state.TimeOfDay = helper.TfString(response.TimeOfDay)

	// a work-around
	// converting hh:mm:ss to hh:mm in case server returns hh:mm:ss
	// client can not send hh:mm:ss , else will be a server error , so no worry
	if response.TimeOfDay != nil && len(strings.Split(*response.TimeOfDay, ":")) == 3 {
		trimmedTime := strings.TrimSuffix(*response.TimeOfDay, ":00")
		state.TimeOfDay = types.StringValue(trimmedTime)
	}

	// this if-else will be removed once things got fixed on powerstore side
	// as for import we don't have pre plan so let the imported value save in state
	if operation == operationImport {
		state.TimeZone = helper.TfString(response.Timezone)
		state.NASAccessType = helper.TfString(response.NasAccessType)
		state.IsReadOnly = types.BoolValue(false)
	} else {
		// a work-around
		// to allow empty string for default values
		// if default value is returned in response, then if empty string in plan
		// update state value as empty string

		if response.Timezone != nil && *response.Timezone == clientgen.TIMEZONEENUM_UTC &&
			!plan.TimeZone.IsUnknown() && !plan.TimeZone.IsNull() &&
			strings.TrimSpace(strings.Trim(plan.TimeZone.ValueString(), "\"")) == "" {
			state.TimeZone = types.StringValue(plan.TimeZone.ValueString())
		} else {
			state.TimeZone = helper.TfString(response.Timezone)
		}

		// as per document, snapshot is default ,  but on server protocol is default
		if response.NasAccessType != nil && *response.NasAccessType == clientgen.NASACCESSTYPEENUM_PROTOCOL &&
			!plan.NASAccessType.IsUnknown() && !plan.NASAccessType.IsNull() &&
			strings.TrimSpace(strings.Trim(plan.NASAccessType.ValueString(), "\"")) == "" {
			state.NASAccessType = types.StringValue(plan.NASAccessType.ValueString())
		} else {
			state.NASAccessType = helper.TfString(response.NasAccessType)
		}

		// a work-around, as we cannot set is_read_only as true on server,
		// so always setting it as plan
		if !plan.IsReadOnly.IsUnknown() && !plan.IsReadOnly.IsNull() {
			state.IsReadOnly = types.BoolValue(plan.IsReadOnly.ValueBool())
		} else {
			state.IsReadOnly = types.BoolValue(false)
		}
	}
	state.IsSecure = helper.TfBool(response.IsSecure)

	// DaysOfWeek mapping
	slice := helper.SliceTransform(response.DaysOfWeek, func(in clientgen.DaysOfWeekEnum) attr.Value {
		return types.StringValue(string(in))
	})
	list, _ := types.ListValue(types.StringType, slice)
	state.DaysOfWeek = list

	state.DesiredRetention = types.Int32Value(helper.Int32Value(response.DesiredRetention))
	state.IsReplica = helper.TfBool(response.IsReplica)
	state.ManagedBy = helper.TfString(response.ManagedBy)
	state.ManagedByID = helper.TfString(response.ManagedById)

	if operation != operationRead {
		// we are saving delete_snaps value in state from plan
		// for future deleteion, if required
		state.DeleteSnaps = plan.DeleteSnaps
	}
}
