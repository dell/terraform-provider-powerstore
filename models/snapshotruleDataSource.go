package models

import "github.com/hashicorp/terraform-plugin-framework/types"

// SnapshotRuleDataSource - SnapshotRule DataSource properties
type SnapshotRuleDataSource struct {
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
	IntervalL10N     types.String `tfsdk:"interval_l10n"`
	TimeZoneL10N     types.String `tfsdk:"timezone_l10n"`
	DaysOfWeek10N    types.List   `tfsdk:"days_of_week_l10n"`
	NASAccessType10N types.String `tfsdk:"nas_access_type_l10n"`
	ManagedByID10N   types.String `tfsdk:"managed_by_l10n"`
	Policies         []Policies   `tfsdk:"policies"`
}

// Policies - Policies associated with Snapshot Rule
type Policies struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}
