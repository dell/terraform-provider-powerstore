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

// VirtualMachinePowerStateEnum The current power state of the VM in vCenter. Not applicable to VM snapshots. * Powered_Off - VM is currently powered off. * Powered_On - VM is currently powered on. * Suspended - VM is currently suspended.  Was added in version 3.0.0.0.
type VirtualMachinePowerStateEnum string

// List of VirtualMachinePowerStateEnum
const (
	VIRTUALMACHINEPOWERSTATEENUM_POWERED_OFF VirtualMachinePowerStateEnum = "Powered_Off"
	VIRTUALMACHINEPOWERSTATEENUM_POWERED_ON VirtualMachinePowerStateEnum = "Powered_On"
	VIRTUALMACHINEPOWERSTATEENUM_SUSPENDED VirtualMachinePowerStateEnum = "Suspended"
)

// All allowed values of VirtualMachinePowerStateEnum enum
var AllowedVirtualMachinePowerStateEnumEnumValues = []VirtualMachinePowerStateEnum{
	"Powered_Off",
	"Powered_On",
	"Suspended",
}

func (v *VirtualMachinePowerStateEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VirtualMachinePowerStateEnum(value)
	for _, existing := range AllowedVirtualMachinePowerStateEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VirtualMachinePowerStateEnum", value)
}

// NewVirtualMachinePowerStateEnumFromValue returns a pointer to a valid VirtualMachinePowerStateEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVirtualMachinePowerStateEnumFromValue(v string) (*VirtualMachinePowerStateEnum, error) {
	ev := VirtualMachinePowerStateEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VirtualMachinePowerStateEnum: valid values are %v", v, AllowedVirtualMachinePowerStateEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VirtualMachinePowerStateEnum) IsValid() bool {
	for _, existing := range AllowedVirtualMachinePowerStateEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VirtualMachinePowerStateEnum value
func (v VirtualMachinePowerStateEnum) Ptr() *VirtualMachinePowerStateEnum {
	return &v
}

type NullableVirtualMachinePowerStateEnum struct {
	value *VirtualMachinePowerStateEnum
	isSet bool
}

func (v NullableVirtualMachinePowerStateEnum) Get() *VirtualMachinePowerStateEnum {
	return v.value
}

func (v *NullableVirtualMachinePowerStateEnum) Set(val *VirtualMachinePowerStateEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualMachinePowerStateEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualMachinePowerStateEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualMachinePowerStateEnum(val *VirtualMachinePowerStateEnum) *NullableVirtualMachinePowerStateEnum {
	return &NullableVirtualMachinePowerStateEnum{value: val, isSet: true}
}

func (v NullableVirtualMachinePowerStateEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualMachinePowerStateEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

