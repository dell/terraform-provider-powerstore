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
	"log"
	"net/url"

	client "terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/models"
	"terraform-provider-powerstore/powerstore/helper"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// newReplicationSessionDataSource returns new replication session data source instance
func newReplicationSessionDataSource() datasource.DataSource {
	return &datasourceReplicationSession{}
}

type datasourceReplicationSession struct {
	client *clientgen.APIClient
}

// Metadata defines data source interface Metadata method
func (d *datasourceReplicationSession) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_replication_session"
}

// Schema defines data source interface Schema method
func (d *datasourceReplicationSession) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This datasource is used to query replication sessions from PowerStore array. The results can be used to monitor replication health, status, and session details for metro and async replication.",
		Description:         "This datasource is used to query replication sessions from PowerStore array. The results can be used to monitor replication health, status, and session details for metro and async replication.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Unique identifier of a specific replication session to query. If provided, only this session is returned.",
				MarkdownDescription: "Unique identifier of a specific replication session to query. If provided, only this session is returned.",
			},
			"replication_sessions": schema.ListNestedAttribute{
				Computed:            true,
				Description:         "List of replication sessions matching the query.",
				MarkdownDescription: "List of replication sessions matching the query.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "Unique identifier of the replication session.",
						},
						"state": schema.StringAttribute{
							Computed:    true,
							Description: "Current state of the replication session.",
						},
						"role": schema.StringAttribute{
							Computed:    true,
							Description: "Role of the replication session.",
						},
						"resource_type": schema.StringAttribute{
							Computed:    true,
							Description: "Type of the storage resource.",
						},
						"data_transfer_state": schema.StringAttribute{
							Computed:    true,
							Description: "Current data transfer state.",
						},
						"type": schema.StringAttribute{
							Computed:    true,
							Description: "Replication session type (e.g., Metro_Active_Active).",
						},
						"last_sync_timestamp": schema.StringAttribute{
							Computed:    true,
							Description: "Time of last successful synchronization.",
						},
						"local_resource_id": schema.StringAttribute{
							Computed:    true,
							Description: "Unique identifier of the local storage resource.",
						},
						"remote_resource_id": schema.StringAttribute{
							Computed:    true,
							Description: "Unique identifier of the remote storage resource.",
						},
						"remote_system_id": schema.StringAttribute{
							Computed:    true,
							Description: "Unique identifier of the remote system.",
						},
						"progress_percentage": schema.Int64Attribute{
							Computed:    true,
							Description: "Progress of the current replication operation.",
						},
						"replication_rule_id": schema.StringAttribute{
							Computed:    true,
							Description: "Associated replication rule instance ID.",
						},
						"last_sync_duration": schema.Int64Attribute{
							Computed:    true,
							Description: "Elapsed time of the last synchronization in milliseconds.",
						},
						"failover_test_in_progress": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether a test failover is in progress.",
						},
						"state_l10n": schema.StringAttribute{
							Computed:    true,
							Description: "Localized state string.",
						},
						"role_l10n": schema.StringAttribute{
							Computed:    true,
							Description: "Localized role string.",
						},
						"resource_type_l10n": schema.StringAttribute{
							Computed:    true,
							Description: "Localized resource type string.",
						},
						"type_l10n": schema.StringAttribute{
							Computed:    true,
							Description: "Localized type string.",
						},
					},
				},
			},
		},
	}
}

// Configure - defines configuration for replication session data source
func (d *datasourceReplicationSession) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	d.client = client.GenClient
}

// Read - reads replication session data
func (d *datasourceReplicationSession) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state models.ReplicationSessionDataSource
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Reading Replication Sessions")

	sel := "*"
	queries := make(url.Values)
	queries.Set("select", sel)

	var sessions []clientgen.ReplicationSessionInstance

	// If specific ID is provided, query by ID
	if !state.ID.IsNull() && state.ID.ValueString() != "" {
		session, _, err := d.client.ReplicationSessionApi.GetReplicationSessionById(ctx, state.ID.ValueString()).Queries(queries).Execute()
		if err != nil {
			resp.Diagnostics.AddError(
				"Error reading replication session",
				"Could not read replication session "+state.ID.ValueString()+": "+err.Error(),
			)
			return
		}
		if session != nil {
			sessions = append(sessions, *session)
		}
	} else {
		// Query all sessions
		allSessions, _, err := d.client.ReplicationSessionApi.GetAllReplicationSessions(ctx).Queries(queries).Execute()
		if err != nil {
			resp.Diagnostics.AddError(
				"Error reading replication sessions",
				"Could not read replication sessions: "+err.Error(),
			)
			return
		}
		sessions = allSessions
	}

	state.ReplicationSessions = mapReplicationSessionsToState(sessions)

	if state.ID.IsNull() || state.ID.ValueString() == "" {
		state.ID = types.StringValue("placeholder")
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	log.Printf("Done reading Replication Sessions, found %d", len(sessions))
}

// mapReplicationSessionsToState converts API response to Terraform state model
func mapReplicationSessionsToState(sessions []clientgen.ReplicationSessionInstance) []models.ReplicationSessionItemModel {
	result := make([]models.ReplicationSessionItemModel, len(sessions))
	for i, s := range sessions {
		item := models.ReplicationSessionItemModel{
			ID:               helper.TfString(s.Id),
			LocalResourceID:  helper.TfString(s.LocalResourceId),
			RemoteResourceID: helper.TfString(s.RemoteResourceId),
			RemoteSystemID:   helper.TfString(s.RemoteSystemId),
			ReplicationRuleID: helper.TfString(s.ReplicationRuleId),
			StateL10n:        helper.TfString(s.StateL10n),
			RoleL10n:         helper.TfString(s.RoleL10n),
			ResourceTypeL10n: helper.TfString(s.ResourceTypeL10n),
			TypeL10n:         helper.TfString(s.TypeL10n),
			LastSyncTimestamp: helper.TfStringFromPTime(s.LastSyncTimestamp),
			FailoverTestInProgress: helper.TfBool(s.FailoverTestInProgress),
		}

		if s.State != nil {
			item.State = types.StringValue(string(*s.State))
		} else {
			item.State = types.StringNull()
		}
		if s.Role != nil {
			item.Role = types.StringValue(string(*s.Role))
		} else {
			item.Role = types.StringNull()
		}
		if s.ResourceType != nil {
			item.ResourceType = types.StringValue(string(*s.ResourceType))
		} else {
			item.ResourceType = types.StringNull()
		}
		if s.DataTransferState != nil {
			item.DataTransferState = types.StringValue(string(*s.DataTransferState))
		} else {
			item.DataTransferState = types.StringNull()
		}
		if s.Type != nil {
			item.Type = types.StringValue(string(*s.Type))
		} else {
			item.Type = types.StringNull()
		}
		if s.ProgressPercentage != nil {
			item.ProgressPercentage = types.Int64Value(int64(*s.ProgressPercentage))
		} else {
			item.ProgressPercentage = types.Int64Null()
		}
		if s.LastSyncDuration != nil {
			item.LastSyncDuration = types.Int64Value(int64(*s.LastSyncDuration))
		} else {
			item.LastSyncDuration = types.Int64Null()
		}

		result[i] = item
	}
	return result
}
