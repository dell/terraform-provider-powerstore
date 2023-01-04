package snapshotrule

import (
	"strings"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Model specifies snapshotRule properties
type model struct {
	ID               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Interval         types.String `tfsdk:"interval"`
	TimeOfDay        types.String `tfsdk:"time_of_day"`
	TimeZone         types.String `tfsdk:"timezone"`
	DaysOfWeek       types.List   `tfsdk:"days_of_week"`
	DesiredRetention types.Int64  `tfsdk:"desired_retention"`
	NASAccessType    types.String `tfsdk:"nas_access_type"`

	// only send by server and are read only
	IsReplica   types.Bool   `tfsdk:"is_replica"`
	IsReadOnly  types.Bool   `tfsdk:"is_read_only"`
	ManagedBy   types.String `tfsdk:"managed_by"`
	ManagedByID types.String `tfsdk:"managed_by_id"`

	// only to be used while deleting
	DeleteSnaps types.Bool `tfsdk:"delete_snaps"`
}

func (m *model) saveSeverResponse(response gopowerstore.SnapshotRule) (diag diag.Diagnostics) {

	m.ID = types.StringValue(response.ID)
	m.Name = types.StringValue(response.Name)
	m.Interval = types.StringValue(string(response.Interval))
	m.TimeOfDay = types.StringValue(response.TimeOfDay)
	m.IsReadOnly = types.BoolValue(response.IsReadOnly)
	m.TimeZone = types.StringValue(string(response.TimeZone))
	m.NASAccessType = types.StringValue(string(response.NASAccessType))
	m.DesiredRetention = types.Int64Value(int64(response.DesiredRetention))
	m.IsReplica = types.BoolValue(response.IsReplica)
	m.ManagedBy = types.StringValue(string(response.ManagedBy))
	m.ManagedByID = types.StringValue(string(response.ManagedById))

	// a work-around
	// converting hh:mm:ss to hh:mm in case server returns hh:mm:ss
	// client can not send hh:mm:ss , else will be a server error , so no worry
	if len(strings.Split(response.TimeOfDay, ":")) == 3 {
		m.TimeOfDay = types.StringValue(strings.TrimSuffix(response.TimeOfDay, ":00"))
	}

	dayOfWeekList := []attr.Value{}
	for _, day := range response.DaysOfWeek {
		dayOfWeekList = append(dayOfWeekList, types.StringValue(string(day)))
	}

	m.DaysOfWeek, diag = types.ListValue(types.StringType, dayOfWeekList)

	// todo, check if still plan and state are not equal
	// mark resources => should be replaced

	return
}

func (m *model) serverRequest() *gopowerstore.SnapshotRuleCreate {

	snapshotRuleCreate := &gopowerstore.SnapshotRuleCreate{
		Name:             m.Name.ValueString(),
		Interval:         gopowerstore.SnapshotRuleIntervalEnum(m.Interval.ValueString()),
		TimeOfDay:        m.TimeOfDay.ValueString(),
		TimeZone:         gopowerstore.TimeZoneEnum(m.TimeZone.ValueString()),
		DesiredRetention: int32(m.DesiredRetention.ValueInt64()),
		NASAccessType:    gopowerstore.NASAccessTypeEnum(m.NASAccessType.ValueString()),
	}

	if len(m.DaysOfWeek.Elements()) > 0 {
		snapshotRuleCreate.DaysOfWeek = []gopowerstore.DaysOfWeekEnum{}

		for _, d := range m.DaysOfWeek.Elements() {
			snapshotRuleCreate.DaysOfWeek = append(
				snapshotRuleCreate.DaysOfWeek,
				gopowerstore.DaysOfWeekEnum(
					strings.Trim(d.String(), "\""),
				),
			)
		}
	}

	return snapshotRuleCreate
}
