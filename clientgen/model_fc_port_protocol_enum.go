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

// FcPortProtocolEnum Protocols supported over FC port. Types of protocol:        * SCSI - SCSI protocol.        * NVMe - NVMe protocol.  Was added in version 2.0.0.0.
type FcPortProtocolEnum string

// List of FcPortProtocolEnum
const (
	FCPORTPROTOCOLENUM_SCSI FcPortProtocolEnum = "SCSI"
	FCPORTPROTOCOLENUM_NVME FcPortProtocolEnum = "NVMe"
)

// All allowed values of FcPortProtocolEnum enum
var AllowedFcPortProtocolEnumEnumValues = []FcPortProtocolEnum{
	"SCSI",
	"NVMe",
}

func (v *FcPortProtocolEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FcPortProtocolEnum(value)
	for _, existing := range AllowedFcPortProtocolEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FcPortProtocolEnum", value)
}

// NewFcPortProtocolEnumFromValue returns a pointer to a valid FcPortProtocolEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFcPortProtocolEnumFromValue(v string) (*FcPortProtocolEnum, error) {
	ev := FcPortProtocolEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FcPortProtocolEnum: valid values are %v", v, AllowedFcPortProtocolEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FcPortProtocolEnum) IsValid() bool {
	for _, existing := range AllowedFcPortProtocolEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FcPortProtocolEnum value
func (v FcPortProtocolEnum) Ptr() *FcPortProtocolEnum {
	return &v
}

type NullableFcPortProtocolEnum struct {
	value *FcPortProtocolEnum
	isSet bool
}

func (v NullableFcPortProtocolEnum) Get() *FcPortProtocolEnum {
	return v.value
}

func (v *NullableFcPortProtocolEnum) Set(val *FcPortProtocolEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableFcPortProtocolEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableFcPortProtocolEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcPortProtocolEnum(val *FcPortProtocolEnum) *NullableFcPortProtocolEnum {
	return &NullableFcPortProtocolEnum{value: val, isSet: true}
}

func (v NullableFcPortProtocolEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFcPortProtocolEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

