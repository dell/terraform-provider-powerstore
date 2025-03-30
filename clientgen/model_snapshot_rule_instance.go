/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

import (
	"encoding/json"
)

// checks if the SnapshotRuleInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapshotRuleInstance{}

// SnapshotRuleInstance Snapshot rule instance. Values was added in 3.0.0.0: timezone, is_read_only, is_replica. Values was added in 3.5.0.0: remote_system_id. This resource type has queriable associations from remote_system, remote_snapshot_session, policy
type SnapshotRuleInstance struct {
	// Unique identifier of the snapshot rule.
	Id *string `json:"id,omitempty"`
	// Snapshot rule name.  This property supports case-insensitive filtering.
	Name *string `json:"name,omitempty"`
	// If set, the unique identifier of the Data Domain remote system to which snaps will be transported. Otherwise, snaps will be taken locally. Was added in version 3.5.0.0.
	RemoteSystemId *string `json:"remote_system_id,omitempty"`
	Interval *SnapRuleIntervalEnum `json:"interval,omitempty"`
	// Time of the day to take a daily snapshot, with format \"hh:mm\" using a 24 hour clock. Either the interval parameter or the time_of_day parameter will be set, but not both. 
	TimeOfDay *string `json:"time_of_day,omitempty"`
	Timezone *TimeZoneEnum `json:"timezone,omitempty"`
	// Days of the week when the snapshot rule should be applied. Days are determined based on the UTC time zone, unless the time_of_day and timezone properties are set. 
	DaysOfWeek []DaysOfWeekEnum `json:"days_of_week,omitempty"`
	// Desired snapshot retention period in hours. The system will retain snapshots for this time period. 
	DesiredRetention *int32 `json:"desired_retention,omitempty"`
	// Indicates whether this is a replica of a snapshot rule on a remote system that is the source of a replication session replicating a storage resource to the local system. 
	IsReplica *bool `json:"is_replica,omitempty"`
	NasAccessType *NASAccessTypeEnum `json:"nas_access_type,omitempty"`
	// Indicates whether this snapshot rule can be modified.  Was added in version 3.0.0.0.
	IsReadOnly *bool `json:"is_read_only,omitempty"`
	ManagedBy *PolicyManagedByEnum `json:"managed_by,omitempty"`
	// Unique identifier of the managing entity based on the value of the managed_by property, as shown below:   * User - Empty   * Metro - Unique identifier of the remote system where the policy was assigned.   * Replication - Unique identifier of the source remote system.   * VMware_vSphere - Unique identifier of the owning VMware vSphere/vCenter.  Was added in version 3.0.0.0.
	ManagedById *string `json:"managed_by_id,omitempty"`
	// Indicates whether snapshots created by this rule should be secure. Secure snapshots cannot be deleted before the expiration time, and the expiration time cannot be reduced. Secure snapshots will only be created for block volumes, volume groups and file systems.  Was added in version 3.5.0.0.
	IsSecure *bool `json:"is_secure,omitempty"`
	// Localized message string corresponding to interval
	IntervalL10n *string `json:"interval_l10n,omitempty"`
	// Localized message string corresponding to timezone Was added in version 2.0.0.0.
	TimezoneL10n *string `json:"timezone_l10n,omitempty"`
	// Localized message array corresponding to days_of_week
	DaysOfWeekL10n []string `json:"days_of_week_l10n,omitempty"`
	// Localized message string corresponding to nas_access_type Was added in version 3.0.0.0.
	NasAccessTypeL10n *string `json:"nas_access_type_l10n,omitempty"`
	// Localized message string corresponding to managed_by Was added in version 3.0.0.0.
	ManagedByL10n *string `json:"managed_by_l10n,omitempty"`
	RemoteSystem *RemoteSystemInstance `json:"remote_system,omitempty"`
	// This is the inverse of the resource type remote_snapshot_session association.
	RemoteSnapshotSessions []RemoteSnapshotSessionInstance `json:"remote_snapshot_sessions,omitempty"`
	// List of the policies that are associated with this snapshot_rule.
	Policies []PolicyInstance `json:"policies,omitempty"`
}

// NewSnapshotRuleInstance instantiates a new SnapshotRuleInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotRuleInstance() *SnapshotRuleInstance {
	this := SnapshotRuleInstance{}
	var timezone TimeZoneEnum = TIMEZONEENUM_UTC
	this.Timezone = &timezone
	var isReplica bool = false
	this.IsReplica = &isReplica
	var nasAccessType NASAccessTypeEnum = NASACCESSTYPEENUM_PROTOCOL
	this.NasAccessType = &nasAccessType
	var isReadOnly bool = false
	this.IsReadOnly = &isReadOnly
	var managedBy PolicyManagedByEnum = POLICYMANAGEDBYENUM_USER
	this.ManagedBy = &managedBy
	var isSecure bool = false
	this.IsSecure = &isSecure
	return &this
}

// NewSnapshotRuleInstanceWithDefaults instantiates a new SnapshotRuleInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotRuleInstanceWithDefaults() *SnapshotRuleInstance {
	this := SnapshotRuleInstance{}
	var timezone TimeZoneEnum = TIMEZONEENUM_UTC
	this.Timezone = &timezone
	var isReplica bool = false
	this.IsReplica = &isReplica
	var nasAccessType NASAccessTypeEnum = NASACCESSTYPEENUM_PROTOCOL
	this.NasAccessType = &nasAccessType
	var isReadOnly bool = false
	this.IsReadOnly = &isReadOnly
	var managedBy PolicyManagedByEnum = POLICYMANAGEDBYENUM_USER
	this.ManagedBy = &managedBy
	var isSecure bool = false
	this.IsSecure = &isSecure
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SnapshotRuleInstance) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SnapshotRuleInstance) SetName(v string) {
	o.Name = &v
}

// GetRemoteSystemId returns the RemoteSystemId field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetRemoteSystemId() string {
	if o == nil || IsNil(o.RemoteSystemId) {
		var ret string
		return ret
	}
	return *o.RemoteSystemId
}

// GetRemoteSystemIdOk returns a tuple with the RemoteSystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetRemoteSystemIdOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteSystemId) {
		return nil, false
	}
	return o.RemoteSystemId, true
}

// HasRemoteSystemId returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasRemoteSystemId() bool {
	if o != nil && !IsNil(o.RemoteSystemId) {
		return true
	}

	return false
}

// SetRemoteSystemId gets a reference to the given string and assigns it to the RemoteSystemId field.
func (o *SnapshotRuleInstance) SetRemoteSystemId(v string) {
	o.RemoteSystemId = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetInterval() SnapRuleIntervalEnum {
	if o == nil || IsNil(o.Interval) {
		var ret SnapRuleIntervalEnum
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetIntervalOk() (*SnapRuleIntervalEnum, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given SnapRuleIntervalEnum and assigns it to the Interval field.
func (o *SnapshotRuleInstance) SetInterval(v SnapRuleIntervalEnum) {
	o.Interval = &v
}

// GetTimeOfDay returns the TimeOfDay field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetTimeOfDay() string {
	if o == nil || IsNil(o.TimeOfDay) {
		var ret string
		return ret
	}
	return *o.TimeOfDay
}

// GetTimeOfDayOk returns a tuple with the TimeOfDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetTimeOfDayOk() (*string, bool) {
	if o == nil || IsNil(o.TimeOfDay) {
		return nil, false
	}
	return o.TimeOfDay, true
}

// HasTimeOfDay returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasTimeOfDay() bool {
	if o != nil && !IsNil(o.TimeOfDay) {
		return true
	}

	return false
}

// SetTimeOfDay gets a reference to the given string and assigns it to the TimeOfDay field.
func (o *SnapshotRuleInstance) SetTimeOfDay(v string) {
	o.TimeOfDay = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetTimezone() TimeZoneEnum {
	if o == nil || IsNil(o.Timezone) {
		var ret TimeZoneEnum
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetTimezoneOk() (*TimeZoneEnum, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given TimeZoneEnum and assigns it to the Timezone field.
func (o *SnapshotRuleInstance) SetTimezone(v TimeZoneEnum) {
	o.Timezone = &v
}

// GetDaysOfWeek returns the DaysOfWeek field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetDaysOfWeek() []DaysOfWeekEnum {
	if o == nil || IsNil(o.DaysOfWeek) {
		var ret []DaysOfWeekEnum
		return ret
	}
	return o.DaysOfWeek
}

// GetDaysOfWeekOk returns a tuple with the DaysOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetDaysOfWeekOk() ([]DaysOfWeekEnum, bool) {
	if o == nil || IsNil(o.DaysOfWeek) {
		return nil, false
	}
	return o.DaysOfWeek, true
}

// HasDaysOfWeek returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasDaysOfWeek() bool {
	if o != nil && !IsNil(o.DaysOfWeek) {
		return true
	}

	return false
}

// SetDaysOfWeek gets a reference to the given []DaysOfWeekEnum and assigns it to the DaysOfWeek field.
func (o *SnapshotRuleInstance) SetDaysOfWeek(v []DaysOfWeekEnum) {
	o.DaysOfWeek = v
}

// GetDesiredRetention returns the DesiredRetention field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetDesiredRetention() int32 {
	if o == nil || IsNil(o.DesiredRetention) {
		var ret int32
		return ret
	}
	return *o.DesiredRetention
}

// GetDesiredRetentionOk returns a tuple with the DesiredRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetDesiredRetentionOk() (*int32, bool) {
	if o == nil || IsNil(o.DesiredRetention) {
		return nil, false
	}
	return o.DesiredRetention, true
}

// HasDesiredRetention returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasDesiredRetention() bool {
	if o != nil && !IsNil(o.DesiredRetention) {
		return true
	}

	return false
}

// SetDesiredRetention gets a reference to the given int32 and assigns it to the DesiredRetention field.
func (o *SnapshotRuleInstance) SetDesiredRetention(v int32) {
	o.DesiredRetention = &v
}

// GetIsReplica returns the IsReplica field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetIsReplica() bool {
	if o == nil || IsNil(o.IsReplica) {
		var ret bool
		return ret
	}
	return *o.IsReplica
}

// GetIsReplicaOk returns a tuple with the IsReplica field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetIsReplicaOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReplica) {
		return nil, false
	}
	return o.IsReplica, true
}

// HasIsReplica returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasIsReplica() bool {
	if o != nil && !IsNil(o.IsReplica) {
		return true
	}

	return false
}

// SetIsReplica gets a reference to the given bool and assigns it to the IsReplica field.
func (o *SnapshotRuleInstance) SetIsReplica(v bool) {
	o.IsReplica = &v
}

// GetNasAccessType returns the NasAccessType field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetNasAccessType() NASAccessTypeEnum {
	if o == nil || IsNil(o.NasAccessType) {
		var ret NASAccessTypeEnum
		return ret
	}
	return *o.NasAccessType
}

// GetNasAccessTypeOk returns a tuple with the NasAccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetNasAccessTypeOk() (*NASAccessTypeEnum, bool) {
	if o == nil || IsNil(o.NasAccessType) {
		return nil, false
	}
	return o.NasAccessType, true
}

// HasNasAccessType returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasNasAccessType() bool {
	if o != nil && !IsNil(o.NasAccessType) {
		return true
	}

	return false
}

// SetNasAccessType gets a reference to the given NASAccessTypeEnum and assigns it to the NasAccessType field.
func (o *SnapshotRuleInstance) SetNasAccessType(v NASAccessTypeEnum) {
	o.NasAccessType = &v
}

// GetIsReadOnly returns the IsReadOnly field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetIsReadOnly() bool {
	if o == nil || IsNil(o.IsReadOnly) {
		var ret bool
		return ret
	}
	return *o.IsReadOnly
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetIsReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReadOnly) {
		return nil, false
	}
	return o.IsReadOnly, true
}

// HasIsReadOnly returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasIsReadOnly() bool {
	if o != nil && !IsNil(o.IsReadOnly) {
		return true
	}

	return false
}

// SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.
func (o *SnapshotRuleInstance) SetIsReadOnly(v bool) {
	o.IsReadOnly = &v
}

// GetManagedBy returns the ManagedBy field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetManagedBy() PolicyManagedByEnum {
	if o == nil || IsNil(o.ManagedBy) {
		var ret PolicyManagedByEnum
		return ret
	}
	return *o.ManagedBy
}

// GetManagedByOk returns a tuple with the ManagedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetManagedByOk() (*PolicyManagedByEnum, bool) {
	if o == nil || IsNil(o.ManagedBy) {
		return nil, false
	}
	return o.ManagedBy, true
}

// HasManagedBy returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasManagedBy() bool {
	if o != nil && !IsNil(o.ManagedBy) {
		return true
	}

	return false
}

// SetManagedBy gets a reference to the given PolicyManagedByEnum and assigns it to the ManagedBy field.
func (o *SnapshotRuleInstance) SetManagedBy(v PolicyManagedByEnum) {
	o.ManagedBy = &v
}

// GetManagedById returns the ManagedById field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetManagedById() string {
	if o == nil || IsNil(o.ManagedById) {
		var ret string
		return ret
	}
	return *o.ManagedById
}

// GetManagedByIdOk returns a tuple with the ManagedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetManagedByIdOk() (*string, bool) {
	if o == nil || IsNil(o.ManagedById) {
		return nil, false
	}
	return o.ManagedById, true
}

// HasManagedById returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasManagedById() bool {
	if o != nil && !IsNil(o.ManagedById) {
		return true
	}

	return false
}

// SetManagedById gets a reference to the given string and assigns it to the ManagedById field.
func (o *SnapshotRuleInstance) SetManagedById(v string) {
	o.ManagedById = &v
}

// GetIsSecure returns the IsSecure field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetIsSecure() bool {
	if o == nil || IsNil(o.IsSecure) {
		var ret bool
		return ret
	}
	return *o.IsSecure
}

// GetIsSecureOk returns a tuple with the IsSecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetIsSecureOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSecure) {
		return nil, false
	}
	return o.IsSecure, true
}

// HasIsSecure returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasIsSecure() bool {
	if o != nil && !IsNil(o.IsSecure) {
		return true
	}

	return false
}

// SetIsSecure gets a reference to the given bool and assigns it to the IsSecure field.
func (o *SnapshotRuleInstance) SetIsSecure(v bool) {
	o.IsSecure = &v
}

// GetIntervalL10n returns the IntervalL10n field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetIntervalL10n() string {
	if o == nil || IsNil(o.IntervalL10n) {
		var ret string
		return ret
	}
	return *o.IntervalL10n
}

// GetIntervalL10nOk returns a tuple with the IntervalL10n field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetIntervalL10nOk() (*string, bool) {
	if o == nil || IsNil(o.IntervalL10n) {
		return nil, false
	}
	return o.IntervalL10n, true
}

// HasIntervalL10n returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasIntervalL10n() bool {
	if o != nil && !IsNil(o.IntervalL10n) {
		return true
	}

	return false
}

// SetIntervalL10n gets a reference to the given string and assigns it to the IntervalL10n field.
func (o *SnapshotRuleInstance) SetIntervalL10n(v string) {
	o.IntervalL10n = &v
}

// GetTimezoneL10n returns the TimezoneL10n field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetTimezoneL10n() string {
	if o == nil || IsNil(o.TimezoneL10n) {
		var ret string
		return ret
	}
	return *o.TimezoneL10n
}

// GetTimezoneL10nOk returns a tuple with the TimezoneL10n field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetTimezoneL10nOk() (*string, bool) {
	if o == nil || IsNil(o.TimezoneL10n) {
		return nil, false
	}
	return o.TimezoneL10n, true
}

// HasTimezoneL10n returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasTimezoneL10n() bool {
	if o != nil && !IsNil(o.TimezoneL10n) {
		return true
	}

	return false
}

// SetTimezoneL10n gets a reference to the given string and assigns it to the TimezoneL10n field.
func (o *SnapshotRuleInstance) SetTimezoneL10n(v string) {
	o.TimezoneL10n = &v
}

// GetDaysOfWeekL10n returns the DaysOfWeekL10n field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetDaysOfWeekL10n() []string {
	if o == nil || IsNil(o.DaysOfWeekL10n) {
		var ret []string
		return ret
	}
	return o.DaysOfWeekL10n
}

// GetDaysOfWeekL10nOk returns a tuple with the DaysOfWeekL10n field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetDaysOfWeekL10nOk() ([]string, bool) {
	if o == nil || IsNil(o.DaysOfWeekL10n) {
		return nil, false
	}
	return o.DaysOfWeekL10n, true
}

// HasDaysOfWeekL10n returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasDaysOfWeekL10n() bool {
	if o != nil && !IsNil(o.DaysOfWeekL10n) {
		return true
	}

	return false
}

// SetDaysOfWeekL10n gets a reference to the given []string and assigns it to the DaysOfWeekL10n field.
func (o *SnapshotRuleInstance) SetDaysOfWeekL10n(v []string) {
	o.DaysOfWeekL10n = v
}

// GetNasAccessTypeL10n returns the NasAccessTypeL10n field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetNasAccessTypeL10n() string {
	if o == nil || IsNil(o.NasAccessTypeL10n) {
		var ret string
		return ret
	}
	return *o.NasAccessTypeL10n
}

// GetNasAccessTypeL10nOk returns a tuple with the NasAccessTypeL10n field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetNasAccessTypeL10nOk() (*string, bool) {
	if o == nil || IsNil(o.NasAccessTypeL10n) {
		return nil, false
	}
	return o.NasAccessTypeL10n, true
}

// HasNasAccessTypeL10n returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasNasAccessTypeL10n() bool {
	if o != nil && !IsNil(o.NasAccessTypeL10n) {
		return true
	}

	return false
}

// SetNasAccessTypeL10n gets a reference to the given string and assigns it to the NasAccessTypeL10n field.
func (o *SnapshotRuleInstance) SetNasAccessTypeL10n(v string) {
	o.NasAccessTypeL10n = &v
}

// GetManagedByL10n returns the ManagedByL10n field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetManagedByL10n() string {
	if o == nil || IsNil(o.ManagedByL10n) {
		var ret string
		return ret
	}
	return *o.ManagedByL10n
}

// GetManagedByL10nOk returns a tuple with the ManagedByL10n field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetManagedByL10nOk() (*string, bool) {
	if o == nil || IsNil(o.ManagedByL10n) {
		return nil, false
	}
	return o.ManagedByL10n, true
}

// HasManagedByL10n returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasManagedByL10n() bool {
	if o != nil && !IsNil(o.ManagedByL10n) {
		return true
	}

	return false
}

// SetManagedByL10n gets a reference to the given string and assigns it to the ManagedByL10n field.
func (o *SnapshotRuleInstance) SetManagedByL10n(v string) {
	o.ManagedByL10n = &v
}

// GetRemoteSystem returns the RemoteSystem field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetRemoteSystem() RemoteSystemInstance {
	if o == nil || IsNil(o.RemoteSystem) {
		var ret RemoteSystemInstance
		return ret
	}
	return *o.RemoteSystem
}

// GetRemoteSystemOk returns a tuple with the RemoteSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetRemoteSystemOk() (*RemoteSystemInstance, bool) {
	if o == nil || IsNil(o.RemoteSystem) {
		return nil, false
	}
	return o.RemoteSystem, true
}

// HasRemoteSystem returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasRemoteSystem() bool {
	if o != nil && !IsNil(o.RemoteSystem) {
		return true
	}

	return false
}

// SetRemoteSystem gets a reference to the given RemoteSystemInstance and assigns it to the RemoteSystem field.
func (o *SnapshotRuleInstance) SetRemoteSystem(v RemoteSystemInstance) {
	o.RemoteSystem = &v
}

// GetRemoteSnapshotSessions returns the RemoteSnapshotSessions field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetRemoteSnapshotSessions() []RemoteSnapshotSessionInstance {
	if o == nil || IsNil(o.RemoteSnapshotSessions) {
		var ret []RemoteSnapshotSessionInstance
		return ret
	}
	return o.RemoteSnapshotSessions
}

// GetRemoteSnapshotSessionsOk returns a tuple with the RemoteSnapshotSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetRemoteSnapshotSessionsOk() ([]RemoteSnapshotSessionInstance, bool) {
	if o == nil || IsNil(o.RemoteSnapshotSessions) {
		return nil, false
	}
	return o.RemoteSnapshotSessions, true
}

// HasRemoteSnapshotSessions returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasRemoteSnapshotSessions() bool {
	if o != nil && !IsNil(o.RemoteSnapshotSessions) {
		return true
	}

	return false
}

// SetRemoteSnapshotSessions gets a reference to the given []RemoteSnapshotSessionInstance and assigns it to the RemoteSnapshotSessions field.
func (o *SnapshotRuleInstance) SetRemoteSnapshotSessions(v []RemoteSnapshotSessionInstance) {
	o.RemoteSnapshotSessions = v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *SnapshotRuleInstance) GetPolicies() []PolicyInstance {
	if o == nil || IsNil(o.Policies) {
		var ret []PolicyInstance
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRuleInstance) GetPoliciesOk() ([]PolicyInstance, bool) {
	if o == nil || IsNil(o.Policies) {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *SnapshotRuleInstance) HasPolicies() bool {
	if o != nil && !IsNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []PolicyInstance and assigns it to the Policies field.
func (o *SnapshotRuleInstance) SetPolicies(v []PolicyInstance) {
	o.Policies = v
}

func (o SnapshotRuleInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnapshotRuleInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RemoteSystemId) {
		toSerialize["remote_system_id"] = o.RemoteSystemId
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.TimeOfDay) {
		toSerialize["time_of_day"] = o.TimeOfDay
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.DaysOfWeek) {
		toSerialize["days_of_week"] = o.DaysOfWeek
	}
	if !IsNil(o.DesiredRetention) {
		toSerialize["desired_retention"] = o.DesiredRetention
	}
	if !IsNil(o.IsReplica) {
		toSerialize["is_replica"] = o.IsReplica
	}
	if !IsNil(o.NasAccessType) {
		toSerialize["nas_access_type"] = o.NasAccessType
	}
	if !IsNil(o.IsReadOnly) {
		toSerialize["is_read_only"] = o.IsReadOnly
	}
	if !IsNil(o.ManagedBy) {
		toSerialize["managed_by"] = o.ManagedBy
	}
	if !IsNil(o.ManagedById) {
		toSerialize["managed_by_id"] = o.ManagedById
	}
	if !IsNil(o.IsSecure) {
		toSerialize["is_secure"] = o.IsSecure
	}
	if !IsNil(o.IntervalL10n) {
		toSerialize["interval_l10n"] = o.IntervalL10n
	}
	if !IsNil(o.TimezoneL10n) {
		toSerialize["timezone_l10n"] = o.TimezoneL10n
	}
	if !IsNil(o.DaysOfWeekL10n) {
		toSerialize["days_of_week_l10n"] = o.DaysOfWeekL10n
	}
	if !IsNil(o.NasAccessTypeL10n) {
		toSerialize["nas_access_type_l10n"] = o.NasAccessTypeL10n
	}
	if !IsNil(o.ManagedByL10n) {
		toSerialize["managed_by_l10n"] = o.ManagedByL10n
	}
	if !IsNil(o.RemoteSystem) {
		toSerialize["remote_system"] = o.RemoteSystem
	}
	if !IsNil(o.RemoteSnapshotSessions) {
		toSerialize["remote_snapshot_sessions"] = o.RemoteSnapshotSessions
	}
	if !IsNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}
	return toSerialize, nil
}

type NullableSnapshotRuleInstance struct {
	value *SnapshotRuleInstance
	isSet bool
}

func (v NullableSnapshotRuleInstance) Get() *SnapshotRuleInstance {
	return v.value
}

func (v *NullableSnapshotRuleInstance) Set(val *SnapshotRuleInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotRuleInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotRuleInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotRuleInstance(val *SnapshotRuleInstance) *NullableSnapshotRuleInstance {
	return &NullableSnapshotRuleInstance{value: val, isSet: true}
}

func (v NullableSnapshotRuleInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotRuleInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


