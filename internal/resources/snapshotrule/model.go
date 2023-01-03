package snapshotrule

import "github.com/hashicorp/terraform-plugin-framework/types"

// Model specifies snapshotRule properties
type model struct {
	ID               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Interval         types.String `tfsdk:"interval"`
	TimeOfDay        types.String `tfsdk:"time_of_day"`
	TimeZone         types.String `tfsdk:"timezone"`
	DaysOfWeek       types.List   `tfsdk:"days_of_week"`
	DesiredRetention types.Int64  `tfsdk:"desired_retention"`
	IsReplica        types.Bool   `tfsdk:"is_replica"`
	NASAccessType    types.String `tfsdk:"nas_access_type"`
	IsReadOnly       types.Bool   `tfsdk:"is_read_only"`
	ManagedBy        types.String `tfsdk:"managed_by"`
	ManagedByID      types.String `tfsdk:"managed_by_id"`
	DeleteSnaps      types.Bool   `tfsdk:"delete_snaps"`
}
