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

// checks if the FileUserQuotaInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileUserQuotaInstance{}

// FileUserQuotaInstance This resource type has queriable associations from file_system, file_tree_quota
type FileUserQuotaInstance struct {
	// Unique identifier of the user quota.
	Id *string `json:"id,omitempty"`
	// Unique identifier of the associated filesystem.
	FileSystemId *string `json:"file_system_id,omitempty"`
	// Unique identifier of the associated tree quota. Values are: - null - if the user quota is not within a quota tree. - tree_quota instance id - if the user quota is within a quota tree. 
	TreeQuotaId *string `json:"tree_quota_id,omitempty"`
	// Unix user identifier (UID) of the user.
	Uid *int64 `json:"uid,omitempty"`
	// Unix username.
	UnixName *string `json:"unix_name,omitempty"`
	// Windows username. The format is domain\\\\user for the domain user.
	WindowsName *string `json:"windows_name,omitempty"`
	// Windows Security Identifier of the user.
	WindowsSid *string `json:"windows_sid,omitempty"`
	State *FileQuotaStateEnum `json:"state,omitempty"`
	// Hard limit of the user quota, in bytes. No hard limit when set to 0. This value can be used to compute amount of space that is consumed without limiting the space.
	HardLimit *int64 `json:"hard_limit,omitempty"`
	// Soft limit of the user quota, in bytes. No soft limit when set to 0.
	SoftLimit *int64 `json:"soft_limit,omitempty"`
	// Remaining grace period, in seconds, after the soft limit is exceeded:   - 0 - Grace period has already expired   - -1 - No grace period in-progress, or infinite grace period set The grace period of user quotas is set in the file system quota configuration. 
	RemainingGracePeriod *int32 `json:"remaining_grace_period,omitempty"`
	// Size currently consumed by the user on the filesystem, in bytes.
	SizeUsed *int64 `json:"size_used,omitempty"`
	// Localized message string corresponding to state
	StateL10n *string `json:"state_l10n,omitempty"`
	FileSystem *FileSystemInstance `json:"file_system,omitempty"`
	TreeQuota *FileTreeQuotaInstance `json:"tree_quota,omitempty"`
}

// NewFileUserQuotaInstance instantiates a new FileUserQuotaInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileUserQuotaInstance() *FileUserQuotaInstance {
	this := FileUserQuotaInstance{}
	return &this
}

// NewFileUserQuotaInstanceWithDefaults instantiates a new FileUserQuotaInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileUserQuotaInstanceWithDefaults() *FileUserQuotaInstance {
	this := FileUserQuotaInstance{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FileUserQuotaInstance) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUserQuotaInstance) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FileUserQuotaInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FileUserQuotaInstance) SetId(v string) {
	o.Id = &v
}

// GetFileSystemId returns the FileSystemId field value if set, zero value otherwise.
func (o *FileUserQuotaInstance) GetFileSystemId() string {
	if o == nil || IsNil(o.FileSystemId) {
		var ret string
		return ret
	}
	return *o.FileSystemId
}

// GetFileSystemIdOk returns a tuple with the FileSystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUserQuotaInstance) GetFileSystemIdOk() (*string, bool) {
	if o == nil || IsNil(o.FileSystemId) {
		return nil, false
	}
	return o.FileSystemId, true
}

// HasFileSystemId returns a boolean if a field has been set.
func (o *FileUserQuotaInstance) HasFileSystemId() bool {
	if o != nil && !IsNil(o.FileSystemId) {
		return true
	}

	return false
}

// SetFileSystemId gets a reference to the given string and assigns it to the FileSystemId field.
func (o *FileUserQuotaInstance) SetFileSystemId(v string) {
	o.FileSystemId = &v
}

// GetTreeQuotaId returns the TreeQuotaId field value if set, zero value otherwise.
func (o *FileUserQuotaInstance) GetTreeQuotaId() string {
	if o == nil || IsNil(o.TreeQuotaId) {
		var ret string
		return ret
	}
	return *o.TreeQuotaId
}

// GetTreeQuotaIdOk returns a tuple with the TreeQuotaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUserQuotaInstance) GetTreeQuotaIdOk() (*string, bool) {
	if o == nil || IsNil(o.TreeQuotaId) {
		return nil, false
	}
	return o.TreeQuotaId, true
}

// HasTreeQuotaId returns a boolean if a field has been set.
func (o *FileUserQuotaInstance) HasTreeQuotaId() bool {
	if o != nil && !IsNil(o.TreeQuotaId) {
		return true
	}

	return false
}

// SetTreeQuotaId gets a reference to the given string and assigns it to the TreeQuotaId field.
func (o *FileUserQuotaInstance) SetTreeQuotaId(v string) {
	o.TreeQuotaId = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *FileUserQuotaInstance) GetUid() int64 {
	if o == nil || IsNil(o.Uid) {
		var ret int64
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUserQuotaInstance) GetUidOk() (*int64, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *FileUserQuotaInstance) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given int64 and assigns it to the Uid field.
func (o *FileUserQuotaInstance) SetUid(v int64) {
	o.Uid = &v
}

// GetUnixName returns the UnixName field value if set, zero value otherwise.
func (o *FileUserQuotaInstance) GetUnixName() string {
	if o == nil || IsNil(o.UnixName) {
		var ret string
		return ret
	}
	return *o.UnixName
}

// GetUnixNameOk returns a tuple with the UnixName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUserQuotaInstance) GetUnixNameOk() (*string, bool) {
	if o == nil || IsNil(o.UnixName) {
		return nil, false
	}
	return o.UnixName, true
}

// HasUnixName returns a boolean if a field has been set.
func (o *FileUserQuotaInstance) HasUnixName() bool {
	if o != nil && !IsNil(o.UnixName) {
		return true
	}

	return false
}

// SetUnixName gets a reference to the given string and assigns it to the UnixName field.
func (o *FileUserQuotaInstance) SetUnixName(v string) {
	o.UnixName = &v
}

// GetWindowsName returns the WindowsName field value if set, zero value otherwise.
func (o *FileUserQuotaInstance) GetWindowsName() string {
	if o == nil || IsNil(o.WindowsName) {
		var ret string
		return ret
	}
	return *o.WindowsName
}

// GetWindowsNameOk returns a tuple with the WindowsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUserQuotaInstance) GetWindowsNameOk() (*string, bool) {
	if o == nil || IsNil(o.WindowsName) {
		return nil, false
	}
	return o.WindowsName, true
}

// HasWindowsName returns a boolean if a field has been set.
func (o *FileUserQuotaInstance) HasWindowsName() bool {
	if o != nil && !IsNil(o.WindowsName) {
		return true
	}

	return false
}

// SetWindowsName gets a reference to the given string and assigns it to the WindowsName field.
func (o *FileUserQuotaInstance) SetWindowsName(v string) {
	o.WindowsName = &v
}

// GetWindowsSid returns the WindowsSid field value if set, zero value otherwise.
func (o *FileUserQuotaInstance) GetWindowsSid() string {
	if o == nil || IsNil(o.WindowsSid) {
		var ret string
		return ret
	}
	return *o.WindowsSid
}

// GetWindowsSidOk returns a tuple with the WindowsSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUserQuotaInstance) GetWindowsSidOk() (*string, bool) {
	if o == nil || IsNil(o.WindowsSid) {
		return nil, false
	}
	return o.WindowsSid, true
}

// HasWindowsSid returns a boolean if a field has been set.
func (o *FileUserQuotaInstance) HasWindowsSid() bool {
	if o != nil && !IsNil(o.WindowsSid) {
		return true
	}

	return false
}

// SetWindowsSid gets a reference to the given string and assigns it to the WindowsSid field.
func (o *FileUserQuotaInstance) SetWindowsSid(v string) {
	o.WindowsSid = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *FileUserQuotaInstance) GetState() FileQuotaStateEnum {
	if o == nil || IsNil(o.State) {
		var ret FileQuotaStateEnum
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUserQuotaInstance) GetStateOk() (*FileQuotaStateEnum, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *FileUserQuotaInstance) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given FileQuotaStateEnum and assigns it to the State field.
func (o *FileUserQuotaInstance) SetState(v FileQuotaStateEnum) {
	o.State = &v
}

// GetHardLimit returns the HardLimit field value if set, zero value otherwise.
func (o *FileUserQuotaInstance) GetHardLimit() int64 {
	if o == nil || IsNil(o.HardLimit) {
		var ret int64
		return ret
	}
	return *o.HardLimit
}

// GetHardLimitOk returns a tuple with the HardLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUserQuotaInstance) GetHardLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.HardLimit) {
		return nil, false
	}
	return o.HardLimit, true
}

// HasHardLimit returns a boolean if a field has been set.
func (o *FileUserQuotaInstance) HasHardLimit() bool {
	if o != nil && !IsNil(o.HardLimit) {
		return true
	}

	return false
}

// SetHardLimit gets a reference to the given int64 and assigns it to the HardLimit field.
func (o *FileUserQuotaInstance) SetHardLimit(v int64) {
	o.HardLimit = &v
}

// GetSoftLimit returns the SoftLimit field value if set, zero value otherwise.
func (o *FileUserQuotaInstance) GetSoftLimit() int64 {
	if o == nil || IsNil(o.SoftLimit) {
		var ret int64
		return ret
	}
	return *o.SoftLimit
}

// GetSoftLimitOk returns a tuple with the SoftLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUserQuotaInstance) GetSoftLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.SoftLimit) {
		return nil, false
	}
	return o.SoftLimit, true
}

// HasSoftLimit returns a boolean if a field has been set.
func (o *FileUserQuotaInstance) HasSoftLimit() bool {
	if o != nil && !IsNil(o.SoftLimit) {
		return true
	}

	return false
}

// SetSoftLimit gets a reference to the given int64 and assigns it to the SoftLimit field.
func (o *FileUserQuotaInstance) SetSoftLimit(v int64) {
	o.SoftLimit = &v
}

// GetRemainingGracePeriod returns the RemainingGracePeriod field value if set, zero value otherwise.
func (o *FileUserQuotaInstance) GetRemainingGracePeriod() int32 {
	if o == nil || IsNil(o.RemainingGracePeriod) {
		var ret int32
		return ret
	}
	return *o.RemainingGracePeriod
}

// GetRemainingGracePeriodOk returns a tuple with the RemainingGracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUserQuotaInstance) GetRemainingGracePeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.RemainingGracePeriod) {
		return nil, false
	}
	return o.RemainingGracePeriod, true
}

// HasRemainingGracePeriod returns a boolean if a field has been set.
func (o *FileUserQuotaInstance) HasRemainingGracePeriod() bool {
	if o != nil && !IsNil(o.RemainingGracePeriod) {
		return true
	}

	return false
}

// SetRemainingGracePeriod gets a reference to the given int32 and assigns it to the RemainingGracePeriod field.
func (o *FileUserQuotaInstance) SetRemainingGracePeriod(v int32) {
	o.RemainingGracePeriod = &v
}

// GetSizeUsed returns the SizeUsed field value if set, zero value otherwise.
func (o *FileUserQuotaInstance) GetSizeUsed() int64 {
	if o == nil || IsNil(o.SizeUsed) {
		var ret int64
		return ret
	}
	return *o.SizeUsed
}

// GetSizeUsedOk returns a tuple with the SizeUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUserQuotaInstance) GetSizeUsedOk() (*int64, bool) {
	if o == nil || IsNil(o.SizeUsed) {
		return nil, false
	}
	return o.SizeUsed, true
}

// HasSizeUsed returns a boolean if a field has been set.
func (o *FileUserQuotaInstance) HasSizeUsed() bool {
	if o != nil && !IsNil(o.SizeUsed) {
		return true
	}

	return false
}

// SetSizeUsed gets a reference to the given int64 and assigns it to the SizeUsed field.
func (o *FileUserQuotaInstance) SetSizeUsed(v int64) {
	o.SizeUsed = &v
}

// GetStateL10n returns the StateL10n field value if set, zero value otherwise.
func (o *FileUserQuotaInstance) GetStateL10n() string {
	if o == nil || IsNil(o.StateL10n) {
		var ret string
		return ret
	}
	return *o.StateL10n
}

// GetStateL10nOk returns a tuple with the StateL10n field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUserQuotaInstance) GetStateL10nOk() (*string, bool) {
	if o == nil || IsNil(o.StateL10n) {
		return nil, false
	}
	return o.StateL10n, true
}

// HasStateL10n returns a boolean if a field has been set.
func (o *FileUserQuotaInstance) HasStateL10n() bool {
	if o != nil && !IsNil(o.StateL10n) {
		return true
	}

	return false
}

// SetStateL10n gets a reference to the given string and assigns it to the StateL10n field.
func (o *FileUserQuotaInstance) SetStateL10n(v string) {
	o.StateL10n = &v
}

// GetFileSystem returns the FileSystem field value if set, zero value otherwise.
func (o *FileUserQuotaInstance) GetFileSystem() FileSystemInstance {
	if o == nil || IsNil(o.FileSystem) {
		var ret FileSystemInstance
		return ret
	}
	return *o.FileSystem
}

// GetFileSystemOk returns a tuple with the FileSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUserQuotaInstance) GetFileSystemOk() (*FileSystemInstance, bool) {
	if o == nil || IsNil(o.FileSystem) {
		return nil, false
	}
	return o.FileSystem, true
}

// HasFileSystem returns a boolean if a field has been set.
func (o *FileUserQuotaInstance) HasFileSystem() bool {
	if o != nil && !IsNil(o.FileSystem) {
		return true
	}

	return false
}

// SetFileSystem gets a reference to the given FileSystemInstance and assigns it to the FileSystem field.
func (o *FileUserQuotaInstance) SetFileSystem(v FileSystemInstance) {
	o.FileSystem = &v
}

// GetTreeQuota returns the TreeQuota field value if set, zero value otherwise.
func (o *FileUserQuotaInstance) GetTreeQuota() FileTreeQuotaInstance {
	if o == nil || IsNil(o.TreeQuota) {
		var ret FileTreeQuotaInstance
		return ret
	}
	return *o.TreeQuota
}

// GetTreeQuotaOk returns a tuple with the TreeQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUserQuotaInstance) GetTreeQuotaOk() (*FileTreeQuotaInstance, bool) {
	if o == nil || IsNil(o.TreeQuota) {
		return nil, false
	}
	return o.TreeQuota, true
}

// HasTreeQuota returns a boolean if a field has been set.
func (o *FileUserQuotaInstance) HasTreeQuota() bool {
	if o != nil && !IsNil(o.TreeQuota) {
		return true
	}

	return false
}

// SetTreeQuota gets a reference to the given FileTreeQuotaInstance and assigns it to the TreeQuota field.
func (o *FileUserQuotaInstance) SetTreeQuota(v FileTreeQuotaInstance) {
	o.TreeQuota = &v
}

func (o FileUserQuotaInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileUserQuotaInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	if !IsNil(o.FileSystemId) {
		toSerialize["file_system_id"] = o.FileSystemId
	}
	if !IsNil(o.TreeQuotaId) {
		toSerialize["tree_quota_id"] = o.TreeQuotaId
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	if !IsNil(o.UnixName) {
		toSerialize["unix_name"] = o.UnixName
	}
	if !IsNil(o.WindowsName) {
		toSerialize["windows_name"] = o.WindowsName
	}
	if !IsNil(o.WindowsSid) {
		toSerialize["windows_sid"] = o.WindowsSid
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.HardLimit) {
		toSerialize["hard_limit"] = o.HardLimit
	}
	if !IsNil(o.SoftLimit) {
		toSerialize["soft_limit"] = o.SoftLimit
	}
	if !IsNil(o.RemainingGracePeriod) {
		toSerialize["remaining_grace_period"] = o.RemainingGracePeriod
	}
	if !IsNil(o.SizeUsed) {
		toSerialize["size_used"] = o.SizeUsed
	}
	if !IsNil(o.StateL10n) {
		toSerialize["state_l10n"] = o.StateL10n
	}
	if !IsNil(o.FileSystem) {
		toSerialize["file_system"] = o.FileSystem
	}
	if !IsNil(o.TreeQuota) {
		toSerialize["tree_quota"] = o.TreeQuota
	}
	return toSerialize, nil
}

type NullableFileUserQuotaInstance struct {
	value *FileUserQuotaInstance
	isSet bool
}

func (v NullableFileUserQuotaInstance) Get() *FileUserQuotaInstance {
	return v.value
}

func (v *NullableFileUserQuotaInstance) Set(val *FileUserQuotaInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableFileUserQuotaInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableFileUserQuotaInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileUserQuotaInstance(val *FileUserQuotaInstance) *NullableFileUserQuotaInstance {
	return &NullableFileUserQuotaInstance{value: val, isSet: true}
}

func (v NullableFileUserQuotaInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileUserQuotaInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


