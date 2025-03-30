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

// FileQuotaStateEnum State of the user quota or tree quota record period. * Ok - No quota limits are exceeded. * Soft_Exceeded - Soft limit is exceeded, and grace period is not expired. * Soft_Exceeded_And_Expired - Soft limit is exceeded, and grace period is expired. * Hard_Reached - Hard limit is reached. 
type FileQuotaStateEnum string

// List of FileQuotaStateEnum
const (
	FILEQUOTASTATEENUM_OK FileQuotaStateEnum = "Ok"
	FILEQUOTASTATEENUM_SOFT_EXCEEDED FileQuotaStateEnum = "Soft_Exceeded"
	FILEQUOTASTATEENUM_SOFT_EXCEEDED_AND_EXPIRED FileQuotaStateEnum = "Soft_Exceeded_And_Expired"
	FILEQUOTASTATEENUM_HARD_REACHED FileQuotaStateEnum = "Hard_Reached"
)

// All allowed values of FileQuotaStateEnum enum
var AllowedFileQuotaStateEnumEnumValues = []FileQuotaStateEnum{
	"Ok",
	"Soft_Exceeded",
	"Soft_Exceeded_And_Expired",
	"Hard_Reached",
}

func (v *FileQuotaStateEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FileQuotaStateEnum(value)
	for _, existing := range AllowedFileQuotaStateEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FileQuotaStateEnum", value)
}

// NewFileQuotaStateEnumFromValue returns a pointer to a valid FileQuotaStateEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileQuotaStateEnumFromValue(v string) (*FileQuotaStateEnum, error) {
	ev := FileQuotaStateEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileQuotaStateEnum: valid values are %v", v, AllowedFileQuotaStateEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileQuotaStateEnum) IsValid() bool {
	for _, existing := range AllowedFileQuotaStateEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileQuotaStateEnum value
func (v FileQuotaStateEnum) Ptr() *FileQuotaStateEnum {
	return &v
}

type NullableFileQuotaStateEnum struct {
	value *FileQuotaStateEnum
	isSet bool
}

func (v NullableFileQuotaStateEnum) Get() *FileQuotaStateEnum {
	return v.value
}

func (v *NullableFileQuotaStateEnum) Set(val *FileQuotaStateEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableFileQuotaStateEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableFileQuotaStateEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileQuotaStateEnum(val *FileQuotaStateEnum) *NullableFileQuotaStateEnum {
	return &NullableFileQuotaStateEnum{value: val, isSet: true}
}

func (v NullableFileQuotaStateEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileQuotaStateEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

