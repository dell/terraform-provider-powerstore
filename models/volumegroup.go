package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Volumegroup - Volumegroup properties
type Volumegroup struct {
	ID                     types.String `tfsdk:"id"`
	Name                   types.String `tfsdk:"name"`
	Description            types.String `tfsdk:"description"`
	VolumeIDs              types.Set    `tfsdk:"volume_ids"`
	IsWriteOrderConsistent types.Bool   `tfsdk:"is_write_order_consistent"`
	ProtectionPolicyID     types.String `tfsdk:"protection_policy_id"`
}
