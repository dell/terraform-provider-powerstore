package snapshotrule

import (
	"context"
	"fmt"
	"regexp"
	"terraform-provider-powerstore/internal/powerstore"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure Resource satsisfies resource.Resource interface
var _ resource.Resource = &Resource{}
var _ resource.ResourceWithImportState = &Resource{}

// NewResource returns snapshotrule new resource instance
func NewResource() resource.Resource {
	return &Resource{}
}

// Resource defines the resource implementation.
type Resource struct {
	client *powerstore.Client
}

// Metadata defines resource interface Metadata method
func (r *Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snapshotrule"
}

// Schema defines resource interface Schema method
func (r *Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					stringvalidator.ExactlyOneOf(path.Expressions{
						path.MatchRoot("interval"),
					}...),

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
				// Optional:            true,
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

// Configure defines resource interface Configure method
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

// Create defines resource interface Create method
func (r *Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan model

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create New SnapshotRule
	// The function returns only ID of the newly created snapshot rule
	createRes, err := r.client.PStoreClient.CreateSnapshotRule(context.Background(), plan.serverRequest())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating snapshot rule",
			fmt.Sprintf("Could not create snapshot rule, unexpected error: %s", err.Error()),
		)
		return
	}

	// Get SnapshotRule Details using ID retrieved above
	serverResponse, err := r.client.PStoreClient.GetSnapshotRule(context.Background(), createRes.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting snapshot rule after creation",
			fmt.Sprintf("Could not get snapshot rule id : %s, unexpected error: %s", createRes.ID, err.Error()),
		)
		return
	}

	diag := plan.saveSeverResponse(serverResponse)
	resp.Diagnostics.Append(diag...)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Read defines resource interface Read method
func (r *Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state model
	// Read Terraform state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	serverResponse, err := r.client.PStoreClient.GetSnapshotRule(context.Background(), state.ID.ValueString())

	// todo distnguish whether error is for resource presence, in case resource is not present
	// we should inform it like resource should be created

	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading snapshot rule",
			fmt.Sprintf("Could not read snapshot rule id: %s, error: %s", state.ID.ValueString(), err.Error()),
		)
		return
	}

	diag := state.saveSeverResponse(serverResponse)
	resp.Diagnostics.Append(diag...)

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

// Update defines resource interface Update method
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

	// Update snapshotRule by calling API
	_, err := r.client.PStoreClient.ModifySnapshotRule(context.Background(), plan.serverRequest(), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating snapshotRule",
			fmt.Sprintf("Could not update snapshotRuleID %s, error: %s", state.ID.ValueString(), err.Error()),
		)
		return
	}

	// Get SnapshotRule Details
	serverResponse, err := r.client.PStoreClient.GetSnapshotRule(context.Background(), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting snapshot rule after update",
			fmt.Sprintf("Could not get snapshot rule, id: %s unexpected error: %s", state.ID.ValueString(), err.Error()),
		)
		return
	}

	diag := plan.saveSeverResponse(serverResponse)
	resp.Diagnostics.Append(diag...)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Delete defines resource interface Delete method
func (r *Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state model

	// Read Terraform state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	deleteParams := &gopowerstore.SnapshotRuleDelete{}

	if !state.DeleteSnaps.IsUnknown() && !state.DeleteSnaps.IsNull() {
		deleteParams.DeleteSnaps = state.DeleteSnaps.ValueBool()
	}

	// Delete snapshotRule by calling API
	_, err := r.client.PStoreClient.DeleteSnapshotRule(context.Background(), deleteParams, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting snapshotRule",
			fmt.Sprintf("Could not delete snapshotRuleID id: %s, error %s", state.ID.ValueString(), err.Error()),
		)
	}
	// todo: instead of returning error , we should check if snapshotRuleID really exists on server
	// and if not , we must return success
	// scenerio - changes from outside of terraform
}

// ImportState defines resource interface ImportState method
func (r *Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	var state model

	// Read Terraform state data into the model
	resp.Diagnostics.Append(resp.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get snapshot details from API and then update what is in state from what the API returns
	serverResponse, err := r.client.PStoreClient.GetSnapshotRule(context.Background(), state.ID.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading snapshot rule",
			fmt.Sprintf("Could not import snapshot rule for id: %s, error %s", state.ID.ValueString(), err.Error()),
		)
		return
	}

	diag := state.saveSeverResponse(serverResponse)
	resp.Diagnostics.Append(diag...)

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
