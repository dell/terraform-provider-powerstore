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
	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ datasource.DataSource              = &replicationRuleDataSource{}
	_ datasource.DataSourceWithConfigure = &replicationRuleDataSource{}
)

// newReplicationRuleDataSource returns the replication rule data source object
func newReplicationRuleDataSource() datasource.DataSource {
	return &replicationRuleDataSource{}
}

type replicationRuleDataSource struct {
	client *client.Client
}

func (d *replicationRuleDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_replication_rule"
}

func (d *replicationRuleDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This datasource is used to query the existing replication rule from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		MarkdownDescription: "This datasource is used to query the existing replication rule from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier of the replication rule. Conflicts with `name`.",
				MarkdownDescription: "Unique identifier of the replication rule. Conflicts with `name`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("name")),
					stringvalidator.ConflictsWith(path.MatchRoot("filter_expression")),
					stringvalidator.LengthAtLeast(1),
				},
			},
			"name": schema.StringAttribute{
				Description:         "Name of the replication rule. Conflicts with `id`.",
				MarkdownDescription: "Name of the replication rule. Conflicts with `id`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ConflictsWith(path.MatchRoot("filter_expression")),
				},
			},
			"filter_expression": schema.StringAttribute{
				Description:         "PowerStore filter expression to filter SMB shares by. Conflicts with `id`, `name` and `file_system_id`.",
				MarkdownDescription: "PowerStore filter expression to filter SMB shares by. Conflicts with `id`, `name` and `file_system_id`.",
				Optional:            true,
				CustomType:          models.FilterExpressionType{},
			},
			"replication_rules": schema.ListNestedAttribute{
				Description:         "List of replication rules.",
				MarkdownDescription: "List of replication rules.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description:         "Unique identifier of the replication rule.",
							MarkdownDescription: "Unique identifier of the replication rule.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							Description:         "Name of the replication rule.",
							MarkdownDescription: "Name of the replication rule.",
							Computed:            true,
						},
						"rpo": schema.StringAttribute{
							Description:         "The RPO (Recovery Point Objective) of the replication rule.",
							MarkdownDescription: "The RPO (Recovery Point Objective) of the replication rule.",
							Computed:            true,
						},
						"remote_system_id": schema.StringAttribute{
							Description:         "The ID of the remote system associated with the replication rule.",
							MarkdownDescription: "The ID of the remote system associated with the replication rule.",
							Computed:            true,
						},
						"alert_threshold": schema.Int64Attribute{
							Description:         "The alert threshold for the replication rule.",
							MarkdownDescription: "The alert threshold for the replication rule.",
							Computed:            true,
						},
						"is_read_only": schema.BoolAttribute{
							Description:         "Indicates whether the replication rule is read-only.",
							MarkdownDescription: "Indicates whether the replication rule is read-only.",
							Computed:            true,
						},
						"is_replica": schema.BoolAttribute{
							Description:         "Indicates whether the replication rule is a replica.",
							MarkdownDescription: "Indicates whether the replication rule is a replica.",
							Computed:            true,
						},
						"managed_by": schema.StringAttribute{
							Description:         "The entity that manages the replication rule.",
							MarkdownDescription: "The entity that manages the replication rule.",
							Computed:            true,
						},
						"managed_by_id": schema.StringAttribute{
							Description:         "The ID of the managing entity.",
							MarkdownDescription: "The ID of the managing entity.",
							Computed:            true,
						},
						"policies": schema.ListNestedAttribute{
							Description:         "The protection policies associated with the replication rule.",
							MarkdownDescription: "The protection policies associated with the replication rule.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "The ID of the protection policy.",
										MarkdownDescription: "The ID of the protection policy.",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										Description:         "The name of the protection policy.",
										MarkdownDescription: "The name of the protection policy.",
										Computed:            true,
									},
								},
							},
						},
						"remote_system": schema.SingleNestedAttribute{
							Description:         "The remote system associated with the replication rule.",
							MarkdownDescription: "The remote system associated with the replication rule.",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description:         "The ID of the remote system.",
									MarkdownDescription: "The ID of the remote system.",
									Computed:            true,
								},
								"name": schema.StringAttribute{
									Description:         "The name of the remote system.",
									MarkdownDescription: "The name of the remote system.",
									Computed:            true,
								},
							},
						},
						"replication_sessions": schema.ListNestedAttribute{
							Description:         "The replication session associated with the replication rule.",
							MarkdownDescription: "The replication session associated with the replication rule.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "The ID of the replication session.",
										MarkdownDescription: "The ID of the replication session.",
										Computed:            true,
									},
									"state": schema.StringAttribute{
										Description:         "The state of the replication session.",
										MarkdownDescription: "The state of the replication session.",
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

func (d *replicationRuleDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = req.ProviderData.(*client.Client)
}

func (d *replicationRuleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var (
		state            models.ReplicationRuleDataSourceModel
		replicationRules []gopowerstore.ReplicationRule
		replicationRule  gopowerstore.ReplicationRule
		err              error
	)

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Read the replication rules based on replication rule id/name and if nothing is mentioned, then it returns all the replication rules
	if state.Name.ValueString() != "" {
		tflog.Info(ctx, fmt.Sprintf("Read replication rule with name: %s", state.Name.ValueString()))
		replicationRule, err = d.client.PStoreClient.GetReplicationRuleByName(context.Background(), state.Name.ValueString())
		replicationRules = append(replicationRules, replicationRule)
		tflog.Info(ctx, fmt.Sprintf("Content: %v", replicationRule))
	} else if state.ID.ValueString() != "" {
		replicationRule, err = d.client.PStoreClient.GetReplicationRule(context.Background(), state.ID.ValueString())
		replicationRules = append(replicationRules, replicationRule)
	} else {
		filters := make(map[string]string)
		if !state.Filters.IsNull() {
			filters = convertQueriesToMap(state.Filters.ValueQueries())
		}
		replicationRules, err = d.client.GetReplicationRules(context.Background(), filters)
	}

	//check if there is any error while getting the replication rules details
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read PowerStore Replication Rules",
			err.Error(),
		)
		return
	}

	state.ReplicationRules = updateReplicationRuleState(replicationRules)
	if state.ID.IsNull() {
		state.ID = types.StringValue("replication_rule_data_source")
	}
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// updateReplicationRuleState iterates over the replication rules list and update the state
func updateReplicationRuleState(ReplicationRules []gopowerstore.ReplicationRule) (response []models.ReplicationRuleDataSource) {
	for _, ReplicationRuleValue := range ReplicationRules {
		policiesList := []models.Policy{}
		for _, policy := range ReplicationRuleValue.ProtectionPolicies {
			policiesList = append(policiesList, models.Policy{
				ID:   types.StringValue(policy.ID),
				Name: types.StringValue(policy.Name),
			})
		}

		remoteSystem := models.RemoteSystem{
			ID:   types.StringValue(ReplicationRuleValue.RemoteSystem.ID),
			Name: types.StringValue(ReplicationRuleValue.RemoteSystem.Name),
		}

		sessionList := []models.ReplicationSession{}
		for _, session := range ReplicationRuleValue.ReplicationSession {
			sessionList = append(sessionList, models.ReplicationSession{
				ID:    types.StringValue(session.ID),
				State: types.StringValue(string(session.State)),
			})
		}

		var replicationRuleState = models.ReplicationRuleDataSource{
			ID:             types.StringValue(ReplicationRuleValue.ID),
			Name:           types.StringValue(ReplicationRuleValue.Name),
			RPO:            types.StringValue(string(ReplicationRuleValue.Rpo)),
			RemoteSystemID: types.StringValue(ReplicationRuleValue.RemoteSystemID),
			IsReadOnly:     types.BoolValue(ReplicationRuleValue.IsReadOnly),
			AlertThreshold: types.Int64Value(int64(ReplicationRuleValue.AlertThreshold)),
			IsReplica:      types.BoolValue(ReplicationRuleValue.IsReplica),
			ManagedBy:      types.StringValue(string(ReplicationRuleValue.ManagedBy)),
			ManagedByID:    types.StringValue(string(ReplicationRuleValue.ManagedByID)),
		}
		replicationRuleState.Policies = policiesList
		replicationRuleState.RemoteSystem = remoteSystem
		replicationRuleState.ReplicationSession = sessionList

		response = append(response, replicationRuleState)
	}
	return response
}
