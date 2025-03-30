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

// checks if the FileDnsInstanceSourceParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileDnsInstanceSourceParameters{}

// FileDnsInstanceSourceParameters Information about the corresponding source NAS Server. Only populated when is_destination_override_enabled flag is set to true. Was added in version 3.0.0.0.  Filtering on the fields of this embedded resource is not supported.
type FileDnsInstanceSourceParameters struct {
	// The list of DNS server IP addresses. The addresses may be IPv4 or IPv6.
	IpAddresses []string `json:"ip_addresses,omitempty"`
}

// NewFileDnsInstanceSourceParameters instantiates a new FileDnsInstanceSourceParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileDnsInstanceSourceParameters() *FileDnsInstanceSourceParameters {
	this := FileDnsInstanceSourceParameters{}
	return &this
}

// NewFileDnsInstanceSourceParametersWithDefaults instantiates a new FileDnsInstanceSourceParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileDnsInstanceSourceParametersWithDefaults() *FileDnsInstanceSourceParameters {
	this := FileDnsInstanceSourceParameters{}
	return &this
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *FileDnsInstanceSourceParameters) GetIpAddresses() []string {
	if o == nil || IsNil(o.IpAddresses) {
		var ret []string
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDnsInstanceSourceParameters) GetIpAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.IpAddresses) {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *FileDnsInstanceSourceParameters) HasIpAddresses() bool {
	if o != nil && !IsNil(o.IpAddresses) {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *FileDnsInstanceSourceParameters) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

func (o FileDnsInstanceSourceParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileDnsInstanceSourceParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IpAddresses) {
		toSerialize["ip_addresses"] = o.IpAddresses
	}
	return toSerialize, nil
}

type NullableFileDnsInstanceSourceParameters struct {
	value *FileDnsInstanceSourceParameters
	isSet bool
}

func (v NullableFileDnsInstanceSourceParameters) Get() *FileDnsInstanceSourceParameters {
	return v.value
}

func (v *NullableFileDnsInstanceSourceParameters) Set(val *FileDnsInstanceSourceParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableFileDnsInstanceSourceParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableFileDnsInstanceSourceParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileDnsInstanceSourceParameters(val *FileDnsInstanceSourceParameters) *NullableFileDnsInstanceSourceParameters {
	return &NullableFileDnsInstanceSourceParameters{value: val, isSet: true}
}

func (v NullableFileDnsInstanceSourceParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileDnsInstanceSourceParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


