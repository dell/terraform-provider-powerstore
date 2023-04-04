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
