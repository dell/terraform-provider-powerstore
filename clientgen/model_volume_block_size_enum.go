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

// VolumeBlockSizeEnum Block size of the volume. Valid values are:  * 512_Bytes - 512 byte block size.  * 4K_Bytes - 4096 byte block size.  * Unknown - Block size cannot be determined. 
type VolumeBlockSizeEnum string

// List of VolumeBlockSizeEnum
const (
	VOLUMEBLOCKSIZEENUM__512_BYTES VolumeBlockSizeEnum = "512_Bytes"
	VOLUMEBLOCKSIZEENUM__4_K_BYTES VolumeBlockSizeEnum = "4K_Bytes"
	VOLUMEBLOCKSIZEENUM_UNKNOWN VolumeBlockSizeEnum = "Unknown"
)

// All allowed values of VolumeBlockSizeEnum enum
var AllowedVolumeBlockSizeEnumEnumValues = []VolumeBlockSizeEnum{
	"512_Bytes",
	"4K_Bytes",
	"Unknown",
}

func (v *VolumeBlockSizeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VolumeBlockSizeEnum(value)
	for _, existing := range AllowedVolumeBlockSizeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VolumeBlockSizeEnum", value)
}

// NewVolumeBlockSizeEnumFromValue returns a pointer to a valid VolumeBlockSizeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVolumeBlockSizeEnumFromValue(v string) (*VolumeBlockSizeEnum, error) {
	ev := VolumeBlockSizeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VolumeBlockSizeEnum: valid values are %v", v, AllowedVolumeBlockSizeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VolumeBlockSizeEnum) IsValid() bool {
	for _, existing := range AllowedVolumeBlockSizeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VolumeBlockSizeEnum value
func (v VolumeBlockSizeEnum) Ptr() *VolumeBlockSizeEnum {
	return &v
}

type NullableVolumeBlockSizeEnum struct {
	value *VolumeBlockSizeEnum
	isSet bool
}

func (v NullableVolumeBlockSizeEnum) Get() *VolumeBlockSizeEnum {
	return v.value
}

func (v *NullableVolumeBlockSizeEnum) Set(val *VolumeBlockSizeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeBlockSizeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeBlockSizeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeBlockSizeEnum(val *VolumeBlockSizeEnum) *NullableVolumeBlockSizeEnum {
	return &NullableVolumeBlockSizeEnum{value: val, isSet: true}
}

func (v NullableVolumeBlockSizeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeBlockSizeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

