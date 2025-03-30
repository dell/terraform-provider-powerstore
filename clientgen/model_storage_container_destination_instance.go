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

// checks if the StorageContainerDestinationInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageContainerDestinationInstance{}

// StorageContainerDestinationInstance A storage container destination defines replication destination for a local storage container on a remote system. New replication groups will use destination storage containers to create their vVol replicas.  Was added in version 3.0.0.0. This resource type has queriable associations from storage_container, remote_system
type StorageContainerDestinationInstance struct {
	// The unique id of the storage container destination.
	Id *string `json:"id,omitempty"`
	// The unique id of the local storage container.
	StorageContainerId *string `json:"storage_container_id,omitempty"`
	// The unique id of the remote system.
	RemoteSystemId *string `json:"remote_system_id,omitempty"`
	// The unique id of the destination storage container on the remote system.
	RemoteStorageContainerId *string `json:"remote_storage_container_id,omitempty"`
	StorageContainer *StorageContainerInstance `json:"storage_container,omitempty"`
	RemoteSystem *RemoteSystemInstance `json:"remote_system,omitempty"`
}

// NewStorageContainerDestinationInstance instantiates a new StorageContainerDestinationInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageContainerDestinationInstance() *StorageContainerDestinationInstance {
	this := StorageContainerDestinationInstance{}
	return &this
}

// NewStorageContainerDestinationInstanceWithDefaults instantiates a new StorageContainerDestinationInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageContainerDestinationInstanceWithDefaults() *StorageContainerDestinationInstance {
	this := StorageContainerDestinationInstance{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StorageContainerDestinationInstance) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageContainerDestinationInstance) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StorageContainerDestinationInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StorageContainerDestinationInstance) SetId(v string) {
	o.Id = &v
}

// GetStorageContainerId returns the StorageContainerId field value if set, zero value otherwise.
func (o *StorageContainerDestinationInstance) GetStorageContainerId() string {
	if o == nil || IsNil(o.StorageContainerId) {
		var ret string
		return ret
	}
	return *o.StorageContainerId
}

// GetStorageContainerIdOk returns a tuple with the StorageContainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageContainerDestinationInstance) GetStorageContainerIdOk() (*string, bool) {
	if o == nil || IsNil(o.StorageContainerId) {
		return nil, false
	}
	return o.StorageContainerId, true
}

// HasStorageContainerId returns a boolean if a field has been set.
func (o *StorageContainerDestinationInstance) HasStorageContainerId() bool {
	if o != nil && !IsNil(o.StorageContainerId) {
		return true
	}

	return false
}

// SetStorageContainerId gets a reference to the given string and assigns it to the StorageContainerId field.
func (o *StorageContainerDestinationInstance) SetStorageContainerId(v string) {
	o.StorageContainerId = &v
}

// GetRemoteSystemId returns the RemoteSystemId field value if set, zero value otherwise.
func (o *StorageContainerDestinationInstance) GetRemoteSystemId() string {
	if o == nil || IsNil(o.RemoteSystemId) {
		var ret string
		return ret
	}
	return *o.RemoteSystemId
}

// GetRemoteSystemIdOk returns a tuple with the RemoteSystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageContainerDestinationInstance) GetRemoteSystemIdOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteSystemId) {
		return nil, false
	}
	return o.RemoteSystemId, true
}

// HasRemoteSystemId returns a boolean if a field has been set.
func (o *StorageContainerDestinationInstance) HasRemoteSystemId() bool {
	if o != nil && !IsNil(o.RemoteSystemId) {
		return true
	}

	return false
}

// SetRemoteSystemId gets a reference to the given string and assigns it to the RemoteSystemId field.
func (o *StorageContainerDestinationInstance) SetRemoteSystemId(v string) {
	o.RemoteSystemId = &v
}

// GetRemoteStorageContainerId returns the RemoteStorageContainerId field value if set, zero value otherwise.
func (o *StorageContainerDestinationInstance) GetRemoteStorageContainerId() string {
	if o == nil || IsNil(o.RemoteStorageContainerId) {
		var ret string
		return ret
	}
	return *o.RemoteStorageContainerId
}

// GetRemoteStorageContainerIdOk returns a tuple with the RemoteStorageContainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageContainerDestinationInstance) GetRemoteStorageContainerIdOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteStorageContainerId) {
		return nil, false
	}
	return o.RemoteStorageContainerId, true
}

// HasRemoteStorageContainerId returns a boolean if a field has been set.
func (o *StorageContainerDestinationInstance) HasRemoteStorageContainerId() bool {
	if o != nil && !IsNil(o.RemoteStorageContainerId) {
		return true
	}

	return false
}

// SetRemoteStorageContainerId gets a reference to the given string and assigns it to the RemoteStorageContainerId field.
func (o *StorageContainerDestinationInstance) SetRemoteStorageContainerId(v string) {
	o.RemoteStorageContainerId = &v
}

// GetStorageContainer returns the StorageContainer field value if set, zero value otherwise.
func (o *StorageContainerDestinationInstance) GetStorageContainer() StorageContainerInstance {
	if o == nil || IsNil(o.StorageContainer) {
		var ret StorageContainerInstance
		return ret
	}
	return *o.StorageContainer
}

// GetStorageContainerOk returns a tuple with the StorageContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageContainerDestinationInstance) GetStorageContainerOk() (*StorageContainerInstance, bool) {
	if o == nil || IsNil(o.StorageContainer) {
		return nil, false
	}
	return o.StorageContainer, true
}

// HasStorageContainer returns a boolean if a field has been set.
func (o *StorageContainerDestinationInstance) HasStorageContainer() bool {
	if o != nil && !IsNil(o.StorageContainer) {
		return true
	}

	return false
}

// SetStorageContainer gets a reference to the given StorageContainerInstance and assigns it to the StorageContainer field.
func (o *StorageContainerDestinationInstance) SetStorageContainer(v StorageContainerInstance) {
	o.StorageContainer = &v
}

// GetRemoteSystem returns the RemoteSystem field value if set, zero value otherwise.
func (o *StorageContainerDestinationInstance) GetRemoteSystem() RemoteSystemInstance {
	if o == nil || IsNil(o.RemoteSystem) {
		var ret RemoteSystemInstance
		return ret
	}
	return *o.RemoteSystem
}

// GetRemoteSystemOk returns a tuple with the RemoteSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageContainerDestinationInstance) GetRemoteSystemOk() (*RemoteSystemInstance, bool) {
	if o == nil || IsNil(o.RemoteSystem) {
		return nil, false
	}
	return o.RemoteSystem, true
}

// HasRemoteSystem returns a boolean if a field has been set.
func (o *StorageContainerDestinationInstance) HasRemoteSystem() bool {
	if o != nil && !IsNil(o.RemoteSystem) {
		return true
	}

	return false
}

// SetRemoteSystem gets a reference to the given RemoteSystemInstance and assigns it to the RemoteSystem field.
func (o *StorageContainerDestinationInstance) SetRemoteSystem(v RemoteSystemInstance) {
	o.RemoteSystem = &v
}

func (o StorageContainerDestinationInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageContainerDestinationInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.StorageContainerId) {
		toSerialize["storage_container_id"] = o.StorageContainerId
	}
	if !IsNil(o.RemoteSystemId) {
		toSerialize["remote_system_id"] = o.RemoteSystemId
	}
	if !IsNil(o.RemoteStorageContainerId) {
		toSerialize["remote_storage_container_id"] = o.RemoteStorageContainerId
	}
	if !IsNil(o.StorageContainer) {
		toSerialize["storage_container"] = o.StorageContainer
	}
	if !IsNil(o.RemoteSystem) {
		toSerialize["remote_system"] = o.RemoteSystem
	}
	return toSerialize, nil
}

type NullableStorageContainerDestinationInstance struct {
	value *StorageContainerDestinationInstance
	isSet bool
}

func (v NullableStorageContainerDestinationInstance) Get() *StorageContainerDestinationInstance {
	return v.value
}

func (v *NullableStorageContainerDestinationInstance) Set(val *StorageContainerDestinationInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageContainerDestinationInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageContainerDestinationInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageContainerDestinationInstance(val *StorageContainerDestinationInstance) *NullableStorageContainerDestinationInstance {
	return &NullableStorageContainerDestinationInstance{value: val, isSet: true}
}

func (v NullableStorageContainerDestinationInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageContainerDestinationInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


