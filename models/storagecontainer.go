package models

import "github.com/hashicorp/terraform-plugin-framework/types"

// StorageContainer - StorageContainer properties
type StorageContainer struct {
	ID              types.String `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
	Quota           types.Int64  `tfsdk:"quota"`
	StorageProtocol types.String `tfsdk:"storage_protocol"`
	HighWaterMark   types.Int64  `tfsdk:"high_water_mark"`
}
