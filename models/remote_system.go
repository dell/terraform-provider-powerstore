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
	// // Unique identifier of the remote system instance.
	// ID string `json:"id,omitempty"`
	ID string `tfsdk:"id"`
	// // User-specified name of the remote system instance.
	// // This property supports case-insensitive filtering
	// Name string `json:"name,omitempty"`
	Name string `tfsdk:"name"`
	// // User-specified description of the remote system instance.
	// Description string `json:"description,omitempty"`
	Description string `tfsdk:"description"`
	// // Serial number of the remote system instance
	// SerialNumber string `json:"serial_number,omitempty"`
	SerialNumber string `tfsdk:"serial_number"`
	// // Management IP address of the remote system instance
	// ManagementAddress string `json:"management_address,omitempty"`
	ManagementAddress string `tfsdk:"management_address"`
	// // Possible data connection states of a remote system
	// DataConnectionState string `json:"data_connection_state,omitempty"`
	DataConnectionState string `tfsdk:"data_connection_state"`
	// // List of supported remote protection capabilities
	Capabilities []string `tfsdk:"capabilities"`
}
