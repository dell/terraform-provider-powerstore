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

// checks if the FileInterfaceSourceParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileInterfaceSourceParameters{}

// FileInterfaceSourceParameters Information about the corresponding source NAS server's File Interface settings. Only populated when is_destination_override_enabled flag is set to true. Was added in version 3.0.0.0.  Filtering on the fields of this embedded resource is not supported.
type FileInterfaceSourceParameters struct {
	// IP address of the network interface. IPv4 and IPv6 are supported.
	IpAddress *string `json:"ip_address,omitempty"`
	// Prefix length for the interface. IPv4 and IPv6 are supported.
	PrefixLength *int32 `json:"prefix_length,omitempty"`
	// Gateway address for the network interface. IPv4 and IPv6 are supported.
	Gateway *string `json:"gateway,omitempty"`
	// Virtual Local Area Network (VLAN) identifier for the interface. The interface uses the identifier to accept packets that have matching VLAN tags.
	VlanId *int32 `json:"vlan_id,omitempty"`
	// Unique identifier of the IP Port that is associated with the file interfacesinterface. Was added in version 3.0.0.0.
	IpPortId *string `json:"ip_port_id,omitempty"`
}

// NewFileInterfaceSourceParameters instantiates a new FileInterfaceSourceParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileInterfaceSourceParameters() *FileInterfaceSourceParameters {
	this := FileInterfaceSourceParameters{}
	var vlanId int32 = 0
	this.VlanId = &vlanId
	return &this
}

// NewFileInterfaceSourceParametersWithDefaults instantiates a new FileInterfaceSourceParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileInterfaceSourceParametersWithDefaults() *FileInterfaceSourceParameters {
	this := FileInterfaceSourceParameters{}
	var vlanId int32 = 0
	this.VlanId = &vlanId
	return &this
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *FileInterfaceSourceParameters) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileInterfaceSourceParameters) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *FileInterfaceSourceParameters) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *FileInterfaceSourceParameters) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetPrefixLength returns the PrefixLength field value if set, zero value otherwise.
func (o *FileInterfaceSourceParameters) GetPrefixLength() int32 {
	if o == nil || IsNil(o.PrefixLength) {
		var ret int32
		return ret
	}
	return *o.PrefixLength
}

// GetPrefixLengthOk returns a tuple with the PrefixLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileInterfaceSourceParameters) GetPrefixLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.PrefixLength) {
		return nil, false
	}
	return o.PrefixLength, true
}

// HasPrefixLength returns a boolean if a field has been set.
func (o *FileInterfaceSourceParameters) HasPrefixLength() bool {
	if o != nil && !IsNil(o.PrefixLength) {
		return true
	}

	return false
}

// SetPrefixLength gets a reference to the given int32 and assigns it to the PrefixLength field.
func (o *FileInterfaceSourceParameters) SetPrefixLength(v int32) {
	o.PrefixLength = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *FileInterfaceSourceParameters) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileInterfaceSourceParameters) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *FileInterfaceSourceParameters) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *FileInterfaceSourceParameters) SetGateway(v string) {
	o.Gateway = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *FileInterfaceSourceParameters) GetVlanId() int32 {
	if o == nil || IsNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileInterfaceSourceParameters) GetVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *FileInterfaceSourceParameters) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *FileInterfaceSourceParameters) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetIpPortId returns the IpPortId field value if set, zero value otherwise.
func (o *FileInterfaceSourceParameters) GetIpPortId() string {
	if o == nil || IsNil(o.IpPortId) {
		var ret string
		return ret
	}
	return *o.IpPortId
}

// GetIpPortIdOk returns a tuple with the IpPortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileInterfaceSourceParameters) GetIpPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.IpPortId) {
		return nil, false
	}
	return o.IpPortId, true
}

// HasIpPortId returns a boolean if a field has been set.
func (o *FileInterfaceSourceParameters) HasIpPortId() bool {
	if o != nil && !IsNil(o.IpPortId) {
		return true
	}

	return false
}

// SetIpPortId gets a reference to the given string and assigns it to the IpPortId field.
func (o *FileInterfaceSourceParameters) SetIpPortId(v string) {
	o.IpPortId = &v
}

func (o FileInterfaceSourceParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileInterfaceSourceParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IpAddress) {
		toSerialize["ip_address"] = o.IpAddress
	}
	if !IsNil(o.PrefixLength) {
		toSerialize["prefix_length"] = o.PrefixLength
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.VlanId) {
		toSerialize["vlan_id"] = o.VlanId
	}
	if !IsNil(o.IpPortId) {
		toSerialize["ip_port_id"] = o.IpPortId
	}
	return toSerialize, nil
}

type NullableFileInterfaceSourceParameters struct {
	value *FileInterfaceSourceParameters
	isSet bool
}

func (v NullableFileInterfaceSourceParameters) Get() *FileInterfaceSourceParameters {
	return v.value
}

func (v *NullableFileInterfaceSourceParameters) Set(val *FileInterfaceSourceParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableFileInterfaceSourceParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableFileInterfaceSourceParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileInterfaceSourceParameters(val *FileInterfaceSourceParameters) *NullableFileInterfaceSourceParameters {
	return &NullableFileInterfaceSourceParameters{value: val, isSet: true}
}

func (v NullableFileInterfaceSourceParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileInterfaceSourceParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


