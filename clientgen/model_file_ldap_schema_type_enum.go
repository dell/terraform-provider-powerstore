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

// FileLDAPSchemaTypeEnum LDAP server schema type. * RFC2307 - OpenLDAP/iPlanet schema. * Microsoft - Microsoft Identity Management for UNIX (IDMU/SFU) schema. * Unknown - Unknown protocol. 
type FileLDAPSchemaTypeEnum string

// List of FileLDAPSchemaTypeEnum
const (
	FILELDAPSCHEMATYPEENUM_RFC2307 FileLDAPSchemaTypeEnum = "RFC2307"
	FILELDAPSCHEMATYPEENUM_MICROSOFT FileLDAPSchemaTypeEnum = "Microsoft"
	FILELDAPSCHEMATYPEENUM_UNKNOWN FileLDAPSchemaTypeEnum = "Unknown"
)

// All allowed values of FileLDAPSchemaTypeEnum enum
var AllowedFileLDAPSchemaTypeEnumEnumValues = []FileLDAPSchemaTypeEnum{
	"RFC2307",
	"Microsoft",
	"Unknown",
}

func (v *FileLDAPSchemaTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FileLDAPSchemaTypeEnum(value)
	for _, existing := range AllowedFileLDAPSchemaTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FileLDAPSchemaTypeEnum", value)
}

// NewFileLDAPSchemaTypeEnumFromValue returns a pointer to a valid FileLDAPSchemaTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileLDAPSchemaTypeEnumFromValue(v string) (*FileLDAPSchemaTypeEnum, error) {
	ev := FileLDAPSchemaTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileLDAPSchemaTypeEnum: valid values are %v", v, AllowedFileLDAPSchemaTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileLDAPSchemaTypeEnum) IsValid() bool {
	for _, existing := range AllowedFileLDAPSchemaTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileLDAPSchemaTypeEnum value
func (v FileLDAPSchemaTypeEnum) Ptr() *FileLDAPSchemaTypeEnum {
	return &v
}

type NullableFileLDAPSchemaTypeEnum struct {
	value *FileLDAPSchemaTypeEnum
	isSet bool
}

func (v NullableFileLDAPSchemaTypeEnum) Get() *FileLDAPSchemaTypeEnum {
	return v.value
}

func (v *NullableFileLDAPSchemaTypeEnum) Set(val *FileLDAPSchemaTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableFileLDAPSchemaTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableFileLDAPSchemaTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileLDAPSchemaTypeEnum(val *FileLDAPSchemaTypeEnum) *NullableFileLDAPSchemaTypeEnum {
	return &NullableFileLDAPSchemaTypeEnum{value: val, isSet: true}
}

func (v NullableFileLDAPSchemaTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileLDAPSchemaTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

