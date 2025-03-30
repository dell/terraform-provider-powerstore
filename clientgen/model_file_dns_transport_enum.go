/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

import (
	"encoding/json"
	"fmt"
)

// FileDNSTransportEnum Transport used when connecting to the DNS Server: * UDP - DNS uses the UDP protocol (default) * TCP - DNS uses the TCP protocol 
type FileDNSTransportEnum string

// List of FileDNSTransportEnum
const (
	FILEDNSTRANSPORTENUM_UDP FileDNSTransportEnum = "UDP"
	FILEDNSTRANSPORTENUM_TCP FileDNSTransportEnum = "TCP"
)

// All allowed values of FileDNSTransportEnum enum
var AllowedFileDNSTransportEnumEnumValues = []FileDNSTransportEnum{
	"UDP",
	"TCP",
}

func (v *FileDNSTransportEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FileDNSTransportEnum(value)
	for _, existing := range AllowedFileDNSTransportEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FileDNSTransportEnum", value)
}

// NewFileDNSTransportEnumFromValue returns a pointer to a valid FileDNSTransportEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileDNSTransportEnumFromValue(v string) (*FileDNSTransportEnum, error) {
	ev := FileDNSTransportEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileDNSTransportEnum: valid values are %v", v, AllowedFileDNSTransportEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileDNSTransportEnum) IsValid() bool {
	for _, existing := range AllowedFileDNSTransportEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileDNSTransportEnum value
func (v FileDNSTransportEnum) Ptr() *FileDNSTransportEnum {
	return &v
}

type NullableFileDNSTransportEnum struct {
	value *FileDNSTransportEnum
	isSet bool
}

func (v NullableFileDNSTransportEnum) Get() *FileDNSTransportEnum {
	return v.value
}

func (v *NullableFileDNSTransportEnum) Set(val *FileDNSTransportEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableFileDNSTransportEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableFileDNSTransportEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileDNSTransportEnum(val *FileDNSTransportEnum) *NullableFileDNSTransportEnum {
	return &NullableFileDNSTransportEnum{value: val, isSet: true}
}

func (v NullableFileDNSTransportEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileDNSTransportEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

