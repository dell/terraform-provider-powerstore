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

package models

import "github.com/hashicorp/terraform-plugin-framework/types"

// RemoteSystemDs - datasource model wrapper
type RemoteSystemDs struct {
	ID      types.String          `tfsdk:"id"`
	Name    types.String          `tfsdk:"name"`
	Filters FilterExpressionValue `tfsdk:"filter_expression"`
	Items   []RemoteSystemDsItem  `tfsdk:"remote_systems"`
}

// RemoteSystemDsItem - Remote System datasource item properties
type RemoteSystemDsItem struct {
	ID                  types.String `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	Description         types.String `tfsdk:"description"`
	SerialNumber        types.String `tfsdk:"serial_number"`
	Type                types.String `tfsdk:"type"`
	ManagementAddress   types.String `tfsdk:"management_address"`
	DataConnectionType  types.String `tfsdk:"data_connection_type"`
	DataConnectionState types.String `tfsdk:"data_connection_state"`
	DataNetworkLatency  types.String `tfsdk:"data_network_latency"`
	State               types.String `tfsdk:"state"`
	Version             types.String `tfsdk:"version"`
	FcTargetWwns        types.List   `tfsdk:"fc_target_wwns"`
	IscsiAddresses      types.List   `tfsdk:"iscsi_addresses"`
	Capabilities        types.List   `tfsdk:"capabilities"`
}
