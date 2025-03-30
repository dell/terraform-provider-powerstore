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

// checks if the ImportUniversalConsistencyGroupInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportUniversalConsistencyGroupInstance{}

// ImportUniversalConsistencyGroupInstance Information about a universal consistency group. Was added in version 4.0.0.0. This resource type has queriable association from remote_system
type ImportUniversalConsistencyGroupInstance struct {
	// Unique identifier of the universal consistency group.
	Id *string `json:"id,omitempty"`
	// Name of the universal consistency group. After import is completed this name will be applied to the new consistency group.   This property supports case-insensitive filtering.
	Name *string `json:"name,omitempty"`
	// Unique identifier of the Remote storage system to which the Universal volume belongs.
	RemoteSystemId *string `json:"remote_system_id,omitempty"`
	ImportableCriteria *CGImportableCriteriaEnum `json:"importable_criteria,omitempty"`
	// Localized message string corresponding to importable_criteria Was added in version 4.0.0.0.
	ImportableCriteriaL10n *string `json:"importable_criteria_l10n,omitempty"`
	RemoteSystem *RemoteSystemInstance `json:"remote_system,omitempty"`
}

// NewImportUniversalConsistencyGroupInstance instantiates a new ImportUniversalConsistencyGroupInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportUniversalConsistencyGroupInstance() *ImportUniversalConsistencyGroupInstance {
	this := ImportUniversalConsistencyGroupInstance{}
	return &this
}

// NewImportUniversalConsistencyGroupInstanceWithDefaults instantiates a new ImportUniversalConsistencyGroupInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportUniversalConsistencyGroupInstanceWithDefaults() *ImportUniversalConsistencyGroupInstance {
	this := ImportUniversalConsistencyGroupInstance{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImportUniversalConsistencyGroupInstance) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUniversalConsistencyGroupInstance) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImportUniversalConsistencyGroupInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ImportUniversalConsistencyGroupInstance) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ImportUniversalConsistencyGroupInstance) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUniversalConsistencyGroupInstance) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ImportUniversalConsistencyGroupInstance) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ImportUniversalConsistencyGroupInstance) SetName(v string) {
	o.Name = &v
}

// GetRemoteSystemId returns the RemoteSystemId field value if set, zero value otherwise.
func (o *ImportUniversalConsistencyGroupInstance) GetRemoteSystemId() string {
	if o == nil || IsNil(o.RemoteSystemId) {
		var ret string
		return ret
	}
	return *o.RemoteSystemId
}

// GetRemoteSystemIdOk returns a tuple with the RemoteSystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUniversalConsistencyGroupInstance) GetRemoteSystemIdOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteSystemId) {
		return nil, false
	}
	return o.RemoteSystemId, true
}

// HasRemoteSystemId returns a boolean if a field has been set.
func (o *ImportUniversalConsistencyGroupInstance) HasRemoteSystemId() bool {
	if o != nil && !IsNil(o.RemoteSystemId) {
		return true
	}

	return false
}

// SetRemoteSystemId gets a reference to the given string and assigns it to the RemoteSystemId field.
func (o *ImportUniversalConsistencyGroupInstance) SetRemoteSystemId(v string) {
	o.RemoteSystemId = &v
}

// GetImportableCriteria returns the ImportableCriteria field value if set, zero value otherwise.
func (o *ImportUniversalConsistencyGroupInstance) GetImportableCriteria() CGImportableCriteriaEnum {
	if o == nil || IsNil(o.ImportableCriteria) {
		var ret CGImportableCriteriaEnum
		return ret
	}
	return *o.ImportableCriteria
}

// GetImportableCriteriaOk returns a tuple with the ImportableCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUniversalConsistencyGroupInstance) GetImportableCriteriaOk() (*CGImportableCriteriaEnum, bool) {
	if o == nil || IsNil(o.ImportableCriteria) {
		return nil, false
	}
	return o.ImportableCriteria, true
}

// HasImportableCriteria returns a boolean if a field has been set.
func (o *ImportUniversalConsistencyGroupInstance) HasImportableCriteria() bool {
	if o != nil && !IsNil(o.ImportableCriteria) {
		return true
	}

	return false
}

// SetImportableCriteria gets a reference to the given CGImportableCriteriaEnum and assigns it to the ImportableCriteria field.
func (o *ImportUniversalConsistencyGroupInstance) SetImportableCriteria(v CGImportableCriteriaEnum) {
	o.ImportableCriteria = &v
}

// GetImportableCriteriaL10n returns the ImportableCriteriaL10n field value if set, zero value otherwise.
func (o *ImportUniversalConsistencyGroupInstance) GetImportableCriteriaL10n() string {
	if o == nil || IsNil(o.ImportableCriteriaL10n) {
		var ret string
		return ret
	}
	return *o.ImportableCriteriaL10n
}

// GetImportableCriteriaL10nOk returns a tuple with the ImportableCriteriaL10n field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUniversalConsistencyGroupInstance) GetImportableCriteriaL10nOk() (*string, bool) {
	if o == nil || IsNil(o.ImportableCriteriaL10n) {
		return nil, false
	}
	return o.ImportableCriteriaL10n, true
}

// HasImportableCriteriaL10n returns a boolean if a field has been set.
func (o *ImportUniversalConsistencyGroupInstance) HasImportableCriteriaL10n() bool {
	if o != nil && !IsNil(o.ImportableCriteriaL10n) {
		return true
	}

	return false
}

// SetImportableCriteriaL10n gets a reference to the given string and assigns it to the ImportableCriteriaL10n field.
func (o *ImportUniversalConsistencyGroupInstance) SetImportableCriteriaL10n(v string) {
	o.ImportableCriteriaL10n = &v
}

// GetRemoteSystem returns the RemoteSystem field value if set, zero value otherwise.
func (o *ImportUniversalConsistencyGroupInstance) GetRemoteSystem() RemoteSystemInstance {
	if o == nil || IsNil(o.RemoteSystem) {
		var ret RemoteSystemInstance
		return ret
	}
	return *o.RemoteSystem
}

// GetRemoteSystemOk returns a tuple with the RemoteSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUniversalConsistencyGroupInstance) GetRemoteSystemOk() (*RemoteSystemInstance, bool) {
	if o == nil || IsNil(o.RemoteSystem) {
		return nil, false
	}
	return o.RemoteSystem, true
}

// HasRemoteSystem returns a boolean if a field has been set.
func (o *ImportUniversalConsistencyGroupInstance) HasRemoteSystem() bool {
	if o != nil && !IsNil(o.RemoteSystem) {
		return true
	}

	return false
}

// SetRemoteSystem gets a reference to the given RemoteSystemInstance and assigns it to the RemoteSystem field.
func (o *ImportUniversalConsistencyGroupInstance) SetRemoteSystem(v RemoteSystemInstance) {
	o.RemoteSystem = &v
}

func (o ImportUniversalConsistencyGroupInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportUniversalConsistencyGroupInstance) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ImportableCriteria) {
		toSerialize["importable_criteria"] = o.ImportableCriteria
	}
	if !IsNil(o.ImportableCriteriaL10n) {
		toSerialize["importable_criteria_l10n"] = o.ImportableCriteriaL10n
	}
	if !IsNil(o.RemoteSystem) {
		toSerialize["remote_system"] = o.RemoteSystem
	}
	return toSerialize, nil
}

type NullableImportUniversalConsistencyGroupInstance struct {
	value *ImportUniversalConsistencyGroupInstance
	isSet bool
}

func (v NullableImportUniversalConsistencyGroupInstance) Get() *ImportUniversalConsistencyGroupInstance {
	return v.value
}

func (v *NullableImportUniversalConsistencyGroupInstance) Set(val *ImportUniversalConsistencyGroupInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableImportUniversalConsistencyGroupInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableImportUniversalConsistencyGroupInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportUniversalConsistencyGroupInstance(val *ImportUniversalConsistencyGroupInstance) *NullableImportUniversalConsistencyGroupInstance {
	return &NullableImportUniversalConsistencyGroupInstance{value: val, isSet: true}
}

func (v NullableImportUniversalConsistencyGroupInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportUniversalConsistencyGroupInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


