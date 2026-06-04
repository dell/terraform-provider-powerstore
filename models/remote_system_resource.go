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

// RemoteSystemResource - Terraform resource model for remote system
type RemoteSystemResource struct {
	ID                  types.String `tfsdk:"id"`
	ManagementAddress   types.String `tfsdk:"management_address"`
	Name                types.String `tfsdk:"name"`
	Description         types.String `tfsdk:"description"`
	Type                types.String `tfsdk:"type"`
	DataNetworkLatency  types.String `tfsdk:"data_network_latency"`
	RemoteUsername      types.String `tfsdk:"remote_username"`
	RemotePassword      types.String `tfsdk:"remote_password"`
	DataConnectionType  types.String `tfsdk:"data_connection_type"`
	IscsiAddresses      types.List   `tfsdk:"iscsi_addresses"`
	SerialNumber        types.String `tfsdk:"serial_number"`
	State               types.String `tfsdk:"state"`
	DataConnectionState types.String `tfsdk:"data_connection_state"`
	Version             types.String `tfsdk:"version"`
	Capabilities        types.List   `tfsdk:"capabilities"`
	// Certificate exchange credentials for PowerStore-to-PowerStore
	ExchangeUsername types.String `tfsdk:"exchange_username"`
	ExchangePassword types.String `tfsdk:"exchange_password"`
}
