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

// Host - Host properties
type Host struct {
	ID               types.String            `tfsdk:"id"`
	Name             types.String            `tfsdk:"name"`
	Description      types.String            `tfsdk:"description"`
	HostGroupID      types.String            `tfsdk:"host_group_id"`
	OsType           types.String            `tfsdk:"os_type"`
	Initiators       []InitiatorCreateModify `tfsdk:"initiators"`
	HostConnectivity types.String            `tfsdk:"host_connectivity"`
}

// InitiatorCreateModify - for adding and modifying initiator to the host
type InitiatorCreateModify struct {
	PortName           types.String `tfsdk:"port_name"`
	PortType           types.String `tfsdk:"port_type"`
	ChapMutualPassword types.String `tfsdk:"chap_mutual_password"`
	ChapMutualUsername types.String `tfsdk:"chap_mutual_username"`
	ChapSinglePassword types.String `tfsdk:"chap_single_password"`
	ChapSingleUsername types.String `tfsdk:"chap_single_username"`
}
