/*
Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0


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

// newQosPolicyDataSource returns qos_policy datasource instance
func newQosPolicyDataSource() datasource.DataSource {
	return &dataSourceQosPolicy{}
}

type dataSourceQosPolicy struct {
	client *clientgen.APIClient
}

// Metadata defines datasource interface Metadata method
func (d *dataSourceQosPolicy) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_qos_policy"
}

// Schema defines datasource interface Schema method
func (d *dataSourceQosPolicy) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This datasource is used to query QoS policy(s) from PowerStore array.",
		Description:         "This datasource is used to query QoS policy(s) from PowerStore array.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Unique identifier of the QoS policy.",
				MarkdownDescription: "Unique identifier of the QoS policy.",
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
				Description:         "Name of the QoS policy.",
				MarkdownDescription: "Name of the QoS policy.",
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("filter_expression")),
					stringvalidator.LengthAtLeast(1),
				},
			},

			"type": schema.StringAttribute{
				Optional:            true,
				Description:         "Type of the QoS policy (" + string(clientgen.POLICYTYPEENUM_QO_S) + " or " + string(clientgen.POLICYTYPEENUM_FILE_PERFORMANCE) + ").",
				MarkdownDescription: "Type of the QoS policy (" + string(clientgen.POLICYTYPEENUM_QO_S) + " or " + string(clientgen.POLICYTYPEENUM_FILE_PERFORMANCE) + ").",
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.Expressions{
						path.MatchRoot("id"),
						path.MatchRoot("name"),
						path.MatchRoot("filter_expression"),
					}...),
					stringvalidator.OneOf(
						string(clientgen.POLICYTYPEENUM_QO_S),
						string(clientgen.POLICYTYPEENUM_FILE_PERFORMANCE),
					),
				},
			},

			"qos_policies": schema.ListNestedAttribute{
				Computed:            true,
				Description:         "List of QoS policies.",
				MarkdownDescription: "List of QoS policies.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:            true,
							Description:         "Unique identifier of the QoS policy.",
							MarkdownDescription: "Unique identifier of the QoS policy.",
						},

						"name": schema.StringAttribute{
							Computed:            true,
							Description:         "Name of the QoS policy.",
							MarkdownDescription: "Name of the QoS policy.",
						},

						"description": schema.StringAttribute{
							Computed:            true,
							Description:         "Description of the QoS policy.",
							MarkdownDescription: "Description of the QoS policy.",
						},

						"type": schema.StringAttribute{
							Computed:            true,
							Description:         "Type of the QoS policy (QoS or File_Performance).",
							MarkdownDescription: "Type of the QoS policy (QoS or File_Performance).",
						},

						"managed_by": schema.StringAttribute{
							Computed:            true,
							Description:         "Entity that manages the policy.",
							MarkdownDescription: "Entity that manages the policy.",
						},

						"managed_by_id": schema.StringAttribute{
							Computed:            true,
							Description:         "Unique identifier of the managing entity.",
							MarkdownDescription: "Unique identifier of the managing entity.",
						},

						"is_read_only": schema.BoolAttribute{
							Computed:            true,
							Description:         "Indicates whether this policy can be modified.",
							MarkdownDescription: "Indicates whether this policy can be modified.",
						},

						"is_replica": schema.BoolAttribute{
							Computed:            true,
							Description:         "Indicates if this is a replica of a policy on a remote system.",
							MarkdownDescription: "Indicates if this is a replica of a policy on a remote system.",
						},

						"type_l10n": schema.StringAttribute{
							Computed:            true,
							Description:         "Localized message string corresponding to type.",
							MarkdownDescription: "Localized message string corresponding to type.",
						},

						"managed_by_l10n": schema.StringAttribute{
							Computed:            true,
							Description:         "Localized message string corresponding to managed_by.",
							MarkdownDescription: "Localized message string corresponding to managed_by.",
						},

						"io_limit_rule_id": schema.StringAttribute{
							Computed:            true,
							Description:         "I/O limit rule identifier (for QoS type).",
							MarkdownDescription: "I/O limit rule identifier (for QoS type).",
						},

						"file_io_limit_rule_id": schema.StringAttribute{
							Computed:            true,
							Description:         "File I/O limit rule identifier (for File_Performance type).",
							MarkdownDescription: "File I/O limit rule identifier (for File_Performance type).",
						},
					},
				},
			},

			"filter_expression": schema.StringAttribute{
				Optional:            true,
				Description:         "Filter expression to query QoS policies.",
				MarkdownDescription: "Filter expression to query QoS policies.",
				CustomType:          models.FilterExpressionType{},
			},
		},
	}
}

// Configure defines configuration for qos_policy datasource
func (d *dataSourceQosPolicy) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

// Read method to read qos_policy datasource information
func (d *dataSourceQosPolicy) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state models.QosPolicyDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	queries := make(url.Values)
	id := state.ID.ValueString()

	sel := "id,name,description,type,type_l10n,managed_by,managed_by_id,is_read_only,is_replica,file_io_limit_rule_id,io_limit_rule(id)"
	queries.Set("select", sel)

	if !state.Name.IsNull() {
		queries.Set("name", "eq."+state.Name.ValueString())
	} else if !state.Type.IsNull() {
		queries.Set("type", "eq."+state.Type.ValueString())
	} else if !state.Filters.IsNull() {
		queries = helper.MergeValues(queries, state.Filters.ValueQueries())
	}

	dsreq := helper.DsReq[clientgen.PolicyInstance, clientgen.ApiGetPolicyByIdRequest, clientgen.ApiGetAllPolicysRequest]{
		Instance:   d.client.PolicyApi.GetPolicyById,
		Collection: d.client.PolicyApi.GetAllPolicys,
	}

	policies, err := dsreq.Execute(ctx, queries, id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read PowerStore QoS Policy",
			err.Error(),
		)
		return
	}

	// check that there is at least one QoS policy if name or id is provided
	if (state.Name.ValueString() != "" || state.ID.ValueString() != "") && len(policies) == 0 {
		resp.Diagnostics.AddError(
			"Unable to Read PowerStore QoS Policy",
			"There is no QoS policy with name "+state.Name.ValueString()+" or id "+state.ID.ValueString(),
		)
		return
	}

	state.QosPolicies = updateQosPolicyState(policies)
	state.ID = types.StringValue("placeholder")
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// updateQosPolicyState iterates over the QoS policies list and update the state
func updateQosPolicyState(policies []clientgen.PolicyInstance) []models.QosPolicyDataSource {
	return helper.SliceTransform(policies, func(in clientgen.PolicyInstance) models.QosPolicyDataSource {
		return models.QosPolicyDataSource{
			ID:            helper.TfString(helper.SetDefault(in.Id, "")),
			Name:          helper.TfString(helper.SetDefault(in.Name, "")),
			Description:   helper.TfString(helper.SetDefault(in.Description, "")),
			TypeL10n:      helper.TfString(helper.SetDefault(in.TypeL10n, "")),
			ManagedByL10n: helper.TfString(helper.SetDefault(in.ManagedByL10n, "")),
			Type:          helper.TfString(in.Type),
			ManagedBy:     helper.TfString(in.ManagedBy),
			ManagedById:   helper.TfString(in.ManagedById),
			IsReadOnly:    helper.TfBool(in.IsReadOnly),
			IsReplica:     helper.TfBool(in.IsReplica),
			IoLimitRuleId: helper.TfObject(in.IoLimitRule, func(r clientgen.IoLimitRuleInstance) types.String {
				return helper.TfString(r.Id)
			}),
			FileIoLimitRuleId: helper.TfString(in.FileIoLimitRuleId),
		}
	})
}
