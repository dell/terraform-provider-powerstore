/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

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

package models

import "github.com/hashicorp/terraform-plugin-framework/types"

// ReplicationRule defines the model for replication rule resource
type ReplicationRule struct {
	ID             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	RPO            types.String `tfsdk:"rpo"`
	RemoteSystemID types.String `tfsdk:"remote_system_id"`
	AlertThreshold types.Int64  `tfsdk:"alert_threshold"`
	IsReadOnly     types.Bool   `tfsdk:"is_read_only"`
}

// ReplicationRuleDataSourceModel defines the model for replication rule data source
type ReplicationRuleDataSourceModel struct {
	ReplicationRules []ReplicationRuleDataSource `tfsdk:"replication_rules"`
	ID               types.String                `tfsdk:"id"`
	Name             types.String                `tfsdk:"name"`
}

// ReplicationRuleDataSource defines the model for replication rule details
type ReplicationRuleDataSource struct {
	ID                 types.String         `tfsdk:"id"`
	Name               types.String         `tfsdk:"name"`
	RPO                types.String         `tfsdk:"rpo"`
	RemoteSystemID     types.String         `tfsdk:"remote_system_id"`
	AlertThreshold     types.Int64          `tfsdk:"alert_threshold"`
	IsReadOnly         types.Bool           `tfsdk:"is_read_only"`
	IsReplica          types.Bool           `tfsdk:"is_replica"`
	ManagedBy          types.String         `tfsdk:"managed_by"`
	ManagedByID        types.String         `tfsdk:"managed_by_id"`
	Policies           []Policy             `tfsdk:"policies"`
	RemoteSystem       RemoteSystem         `tfsdk:"remote_system"`
	ReplicationSession []ReplicationSession `tfsdk:"replication_sessions"`
}

// Policy defines the model for policy
type Policy struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// RemoteSystem defines the model for remote system
type RemoteSystem struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// ReplicationSession defines the model for replication session
type ReplicationSession struct {
	ID    types.String `tfsdk:"id"`
	State types.String `tfsdk:"state"`
}
