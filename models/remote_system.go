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

// RemoteSystemDs
type RemoteSystemDs struct {
	ID      types.String          `tfsdk:"id"`
	Name    types.String          `tfsdk:"name"`
	Filters FilterExpressionValue `tfsdk:"filter_expression"`
	Items   []RemoteSystemDsItem  `tfsdk:"remote_systems"`
}

// RemoteSystem - RemoteSystem properties
type RemoteSystemDsItem struct {
	ID                  string   `tfsdk:"id"`
	Name                string   `tfsdk:"name"`
	Description         string   `tfsdk:"description"`
	SerialNumber        string   `tfsdk:"serial_number"`
	Type                string   `tfsdk:"type"`
	ManagementAddress   string   `tfsdk:"management_address"`
	DataConnectionState string   `tfsdk:"data_connection_state"`
	DataNetworkLatency  string   `tfsdk:"data_network_latency"`
	Capabilities        []string `tfsdk:"capabilities"`
}
