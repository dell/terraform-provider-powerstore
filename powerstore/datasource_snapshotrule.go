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
	"net/url"

	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/models"
	"terraform-provider-powerstore/powerstore/helper"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &snapshotRuleDataSource{}
	_ datasource.DataSourceWithConfigure = &snapshotRuleDataSource{}
)

// newSnapshotRuleDataSource returns the snapshot rule data source object
func newSnapshotRuleDataSource() datasource.DataSource {
	return &snapshotRuleDataSource{}
}

type snapshotRuleDataSource struct {
	client *clientgen.APIClient
}

type snapshotRuleDataSourceModel struct {
	SnapshotRules []models.SnapshotRuleDataSource `tfsdk:"snapshot_rules"`
	ID            types.String                    `tfsdk:"id"`
	Name          types.String                    `tfsdk:"name"`
	Filters       models.FilterExpressionValue    `tfsdk:"filter_expression"`
}

func (d *snapshotRuleDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snapshotrule"
}

func (d *snapshotRuleDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This datasource is used to query the existing snapshot rule from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		MarkdownDescription: "This datasource is used to query the existing snapshot rule from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier of the snapshot rule instance. Conflicts with `name`.",
				MarkdownDescription: "Unique identifier of the snapshot rule instance. Conflicts with `name`.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("filter_expression")),
					stringvalidator.ConflictsWith(path.MatchRoot("name")),
					stringvalidator.LengthAtLeast(1),
				},
			},
			"name": schema.StringAttribute{
				Description:         "Name of the snapshot rule. Conflicts with `id`.",
				MarkdownDescription: "Name of the snapshot rule. Conflicts with `id`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("filter_expression")),
					stringvalidator.LengthAtLeast(1),
				},
			},
			"filter_expression": schema.StringAttribute{
				Description:         "PowerStore filter expression to filter Host by. Conflicts with `id` and `name`.",
				MarkdownDescription: "PowerStore filter expression to filter Host by. Conflicts with `id` and `name`.",
				Optional:            true,
				CustomType:          models.FilterExpressionType{},
			},
			"snapshot_rules": schema.ListNestedAttribute{
				Description:         "List of snapshot rules.",
				MarkdownDescription: "List of snapshot rules.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description:         "The ID of the snapshot rule.",
							MarkdownDescription: "The ID of the snapshot rule.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							Description:         "Name of the snapshot rule.",
							MarkdownDescription: "Name of the snapshot rule.",
							Computed:            true,
						},
						"interval": schema.StringAttribute{
							Description:         "The interval of the snapshot rule.",
							MarkdownDescription: "The interval of the snapshot rule.",
							Computed:            true,
						},
						"time_of_day": schema.StringAttribute{
							Description:         "The time of the day to take a daily snapshot, with format hh:mm.",
							MarkdownDescription: "The time of the day to take a daily snapshot, with format hh:mm.",
							Computed:            true,
						},
						"timezone": schema.StringAttribute{
							Description:         "The time zone identifier for applying the time zone to the time_of_day for a snapshot rule.",
							MarkdownDescription: "The time zone identifier for applying the time zone to the time_of_day for a snapshot rule.",
							Computed:            true,
						},
						"days_of_week": schema.ListAttribute{
							ElementType:         types.StringType,
							Optional:            true,
							Computed:            true,
							Description:         "The days of the week when the snapshot rule should be applied.",
							MarkdownDescription: "The days of the week when the snapshot rule should be applied.",
						},
						"desired_retention": schema.Int64Attribute{
							Description:         "The Desired snapshot retention period in hours to retain snapshots for this time period.",
							MarkdownDescription: "The Desired snapshot retention period in hours to retain snapshots for this time period.",
							Computed:            true,
						},
						"is_replica": schema.BoolAttribute{
							Description:         "Indicates whether this is a replica of a snapshot rule on a remote system.",
							MarkdownDescription: "Indicates whether this is a replica of a snapshot rule on a remote system.",
							Computed:            true,
						},
						"nas_access_type": schema.StringAttribute{
							Description:         "The NAS filesystem snapshot access method for snapshot rule.",
							MarkdownDescription: "The NAS filesystem snapshot access method for snapshot rule.",
							Computed:            true,
						},
						"is_read_only": schema.BoolAttribute{
							Description:         "Indicates whether this snapshot rule can be modified.",
							MarkdownDescription: "Indicates whether this snapshot rule can be modified.",
							Computed:            true,
						},
						"is_secure": schema.BoolAttribute{
							Description:         "Indicates whether snapshots created by this rule should be secure. Secure snapshots cannot be deleted before the expiration time, and the expiration time cannot be reduced.",
							MarkdownDescription: "Indicates whether snapshots created by this rule should be secure. Secure snapshots cannot be deleted before the expiration time, and the expiration time cannot be reduced.",
							Computed:            true,
						},
						"managed_by": schema.StringAttribute{
							Description:         "The entity that owns and manages the instance.",
							MarkdownDescription: "The entity that owns and manages the instance.",
							Computed:            true,
						},
						"managed_by_id": schema.StringAttribute{
							Description:         "The unique id of the managing entity.",
							MarkdownDescription: "The unique id of the managing entity.",
							Computed:            true,
						},
						"interval_l10n": schema.StringAttribute{
							Description:         "Localized message string corresponding to interval",
							MarkdownDescription: "Localized message string corresponding to interval",
							Computed:            true,
						},
						"timezone_l10n": schema.StringAttribute{
							Description:         "Localized message string corresponding to timezone.",
							MarkdownDescription: "Localized message string corresponding to timezone.",
							Computed:            true,
						},
						"days_of_week_l10n": schema.ListAttribute{
							ElementType:         types.StringType,
							Description:         "Localized message array corresponding to days_of_week",
							MarkdownDescription: "Localized message array corresponding to days_of_week",
							Computed:            true,
						},
						"nas_access_type_l10n": schema.StringAttribute{
							Description:         "Localized message string corresponding to nas_access_type.",
							MarkdownDescription: "Localized message string corresponding to nas_access_type.",
							Computed:            true,
						},
						"managed_by_l10n": schema.StringAttribute{
							Description:         "Localized message string corresponding to managed_by.",
							MarkdownDescription: "Localized message string corresponding to managed_by.",
							Computed:            true,
						},
						"policies": schema.ListNestedAttribute{
							Description:         "List of the protection policies that are associated with the snapshot_rule.",
							MarkdownDescription: "List of the protection policies that are associated with the snapshot_rule..",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "Unique identifier of the protection policy.",
										MarkdownDescription: "Unique identifier of the protection policy.",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										Description:         "Name of the protection policy.",
										MarkdownDescription: "Name of the protection policy.",
										Computed:            true,
									},
									"description": schema.StringAttribute{
										Description:         "Description of the protection policy.",
										MarkdownDescription: "Description of the protection policy.",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *snapshotRuleDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(*client.Client)
	d.client = client.GenClient
}

func (d *snapshotRuleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state snapshotRuleDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	sel := "*,policies(*)"
	queries := make(url.Values)
	queries.Set("select", sel)

	id := state.ID.ValueString()
	if !state.Name.IsNull() {
		queries.Set("name", "eq."+state.Name.ValueString())
	}
	if !state.Filters.IsNull() {
		queries = helper.MergeValues(queries, state.Filters.ValueQueries())
	}

	dsreq := helper.DsReq[clientgen.SnapshotRuleInstance, clientgen.ApiGetSnapshotRuleByIdRequest, clientgen.ApiGetAllSnapshotRulesRequest]{
		Instance:   d.client.SnapshotRuleApi.GetSnapshotRuleById,
		Collection: d.client.SnapshotRuleApi.GetAllSnapshotRules,
	}

	snapshotRules, err := dsreq.Execute(ctx, queries, id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read PowerStore Snapshot Rules",
			err.Error(),
		)
		return
	}

	if state.Name.ValueString() != "" && len(snapshotRules) == 0 {
		resp.Diagnostics.AddError(
			"Unable to Read PowerStore Snapshot Rule",
			"There is no snapshot rule with name "+state.Name.ValueString(),
		)
		return
	}

	state.SnapshotRules = updateSnapshotRuleState(snapshotRules)
	state.ID = types.StringValue("placeholder")
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// updateSnapshotRuleState iterates over the snapshot rules list and update the state
func updateSnapshotRuleState(snapshotRules []clientgen.SnapshotRuleInstance) []models.SnapshotRuleDataSource {
	return helper.SliceTransform(snapshotRules, func(in clientgen.SnapshotRuleInstance) models.SnapshotRuleDataSource {
		return models.SnapshotRuleDataSource{
			ID:               helper.TfString(in.Id),
			Name:             helper.TfString(in.Name),
			Interval:         helper.TfString(in.Interval),
			TimeOfDay:        helper.TfString(in.TimeOfDay),
			TimeZone:         helper.TfString(in.Timezone),
			DesiredRetention: helper.TfInt32AsInt64(in.DesiredRetention),
			IsReplica:        helper.TfBool(in.IsReplica),
			NASAccessType:    helper.TfString(in.NasAccessType),
			IsReadOnly:       helper.TfBool(in.IsReadOnly),
			IsSecure:         helper.TfBool(in.IsSecure),
			ManagedBy:        helper.TfString(in.ManagedBy),
			ManagedByID:      helper.TfString(in.ManagedById),
			IntervalL10N:     helper.TfString(in.IntervalL10n),
			TimeZoneL10N:     helper.TfString(in.TimezoneL10n),
			NASAccessType10N: helper.TfString(in.NasAccessTypeL10n),
			ManagedByID10N:   helper.TfString(in.ManagedByL10n),
			DaysOfWeek: func() types.List {
				slice := helper.SliceTransform(in.DaysOfWeek, func(in clientgen.DaysOfWeekEnum) attr.Value {
					return types.StringValue(string(in))
				})
				list, _ := types.ListValue(types.StringType, slice)
				return list
			}(),
			DaysOfWeek10N: func() types.List {
				slice := helper.SliceTransform(in.DaysOfWeekL10n, func(in string) attr.Value {
					return types.StringValue(in)
				})
				list, _ := types.ListValue(types.StringType, slice)
				return list
			}(),
			Policies: helper.SliceTransform(in.Policies, func(in clientgen.PolicyInstance) models.Policies {
				return models.Policies{
					ID:          helper.TfString(in.Id),
					Name:        helper.TfString(in.Name),
					Description: helper.TfString(in.Description),
				}
			}),
		}
	})
}
