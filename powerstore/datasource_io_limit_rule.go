/*
Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	"net/url"

	client "terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/models"
	"terraform-provider-powerstore/powerstore/helper"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// newIoLimitRuleDataSource returns io_limit_rule datasource instance
func newIoLimitRuleDataSource() datasource.DataSource {
	return &dataSourceIoLimitRule{}
}

type dataSourceIoLimitRule struct {
	client *clientgen.APIClient
}

// Metadata defines datasource interface Metadata method
func (d *dataSourceIoLimitRule) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_io_limit_rule"
}

// Schema defines datasource interface Schema method
func (d *dataSourceIoLimitRule) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This datasource is used to query I/O limit rule(s) from PowerStore array.",
		Description:         "This datasource is used to query I/O limit rule(s) from PowerStore array.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Unique identifier of the I/O limit rule.",
				MarkdownDescription: "Unique identifier of the I/O limit rule.",
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.Expressions{
						path.MatchRoot("name"),
						path.MatchRoot("filter_expression"),
					}...),
					stringvalidator.LengthAtLeast(1),
				},
			},

			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Name of the I/O limit rule.",
				MarkdownDescription: "Name of the I/O limit rule.",
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("filter_expression")),
					stringvalidator.LengthAtLeast(1),
				},
			},

			"type": schema.StringAttribute{
				Optional:            true,
				Description:         "Type of the I/O limit rule (" + string(clientgen.BANDWIDTHLIMITTYPEENUM_ABSOLUTE) + " or " + string(clientgen.BANDWIDTHLIMITTYPEENUM_DENSITY) + ").",
				MarkdownDescription: "Type of the I/O limit rule (" + string(clientgen.BANDWIDTHLIMITTYPEENUM_ABSOLUTE) + " or " + string(clientgen.BANDWIDTHLIMITTYPEENUM_DENSITY) + ").",
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.Expressions{
						path.MatchRoot("id"),
						path.MatchRoot("name"),
						path.MatchRoot("filter_expression"),
					}...),
					stringvalidator.OneOf(
						string(clientgen.BANDWIDTHLIMITTYPEENUM_ABSOLUTE),
						string(clientgen.BANDWIDTHLIMITTYPEENUM_DENSITY),
					),
				},
			},

			"io_limit_rules": schema.ListNestedAttribute{
				Computed:            true,
				Description:         "List of I/O limit rules.",
				MarkdownDescription: "List of I/O limit rules.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:            true,
							Description:         "Unique identifier of the I/O limit rule.",
							MarkdownDescription: "Unique identifier of the I/O limit rule.",
						},

						"name": schema.StringAttribute{
							Computed:            true,
							Description:         "Name of the I/O limit rule.",
							MarkdownDescription: "Name of the I/O limit rule.",
						},

						"type": schema.StringAttribute{
							Computed:            true,
							Description:         "Type of bandwidth limit (Absolute or Density).",
							MarkdownDescription: "Type of bandwidth limit (Absolute or Density).",
						},

						"max_iops": schema.Int32Attribute{
							Computed:            true,
							Description:         "Maximum I/O operations in IOPS or IOPS per GB.",
							MarkdownDescription: "Maximum I/O operations in IOPS or IOPS per GB.",
						},

						"max_bw": schema.Int32Attribute{
							Computed:            true,
							Description:         "Maximum I/O bandwidth in Kilobytes per second or Kilobytes per second / per GB.",
							MarkdownDescription: "Maximum I/O bandwidth in Kilobytes per second or Kilobytes per second / per GB.",
						},

						"burst_percentage": schema.Int32Attribute{
							Computed:            true,
							Description:         "Percentage indicating by how much the limit may be exceeded.",
							MarkdownDescription: "Percentage indicating by how much the limit may be exceeded.",
						},

						"type_l10n": schema.StringAttribute{
							Computed:            true,
							Description:         "Localized message string corresponding to type.",
							MarkdownDescription: "Localized message string corresponding to type.",
						},
					},
				},
			},

			"filter_expression": schema.StringAttribute{
				Optional:            true,
				Description:         "Filter expression to query I/O limit rules.",
				MarkdownDescription: "Filter expression to query I/O limit rules.",
				CustomType:          models.FilterExpressionType{},
			},
		},
	}
}

// Configure defines configuration for io_limit_rule datasource
func (d *dataSourceIoLimitRule) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Datasource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = c.GenClient
}

// Read method to read io_limit_rule datasource information
func (d *dataSourceIoLimitRule) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state models.IoLimitRuleDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	queries := make(url.Values)
	id := state.ID.ValueString()

	sel := "id,name,type,max_iops,max_bw,burst_percentage,type_l10n"
	queries.Set("select", sel)

	if !state.Name.IsNull() {
		queries.Set("name", "eq."+state.Name.ValueString())
	} else if !state.Type.IsNull() {
		queries.Set("type", "eq."+state.Type.ValueString())
	} else if !state.Filters.IsNull() {
		queries = helper.MergeValues(queries, state.Filters.ValueQueries())
	}

	dsreq := helper.DsReq[clientgen.IoLimitRuleInstance, clientgen.ApiGetIoLimitRuleByIdRequest, clientgen.ApiGetAllIoLimitRulesRequest]{
		Instance:   d.client.IoLimitRuleApi.GetIoLimitRuleById,
		Collection: d.client.IoLimitRuleApi.GetAllIoLimitRules,
	}

	ioLimitRules, err := dsreq.Execute(ctx, queries, id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read PowerStore I/O Limit Rule",
			err.Error(),
		)
		return
	}

	// check that there is at least one I/O limit rule if name or id is provided
	if (state.Name.ValueString() != "" || state.ID.ValueString() != "") && len(ioLimitRules) == 0 {
		resp.Diagnostics.AddError(
			"Unable to Read PowerStore I/O Limit Rule",
			"There is no I/O limit rule with name "+state.Name.ValueString()+" or id "+state.ID.ValueString(),
		)
		return
	}

	state.IoLimitRules = updateIoLimitRuleState(ioLimitRules)
	state.ID = types.StringValue("placeholder")
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// updateIoLimitRuleState iterates over the I/O limit rules list and update the state
func updateIoLimitRuleState(ioLimitRules []clientgen.IoLimitRuleInstance) []models.IoLimitRuleDataSource {
	return helper.SliceTransform(ioLimitRules, func(in clientgen.IoLimitRuleInstance) models.IoLimitRuleDataSource {
		return models.IoLimitRuleDataSource{
			ID:              helper.TfString(helper.SetDefault(in.Id, "")),
			Name:            helper.TfString(helper.SetDefault(in.Name, "")),
			TypeL10n:        helper.TfString(helper.SetDefault(in.TypeL10n, "")),
			Type:            helper.TfString(in.Type),
			MaxIops:         helper.TfInt32(helper.SetDefault(in.MaxIops, 0)),
			MaxBw:           helper.TfInt32(helper.SetDefault(in.MaxBw, 0)),
			BurstPercentage: helper.TfInt32(helper.SetDefault(in.BurstPercentage, 0)),
		}
	})
}
