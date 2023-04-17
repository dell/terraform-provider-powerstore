package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// HostGroup - hostGroup properties
type HostGroup struct {
	ID               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Description      types.String `tfsdk:"description"`
	HostIDs          types.Set    `tfsdk:"host_ids"`
	HostNames        types.Set    `tfsdk:"host_names"`
	HostConnectivity types.String `tfsdk:"host_connectivity"`
}
