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

// HostAgentStatusEnum Status of the import host system. Valid values are:  * Unknown - Agent status is unknown.  * Running - Agent is up and running.  * Conflict_Detected - Agent detected that there are multiple MPIOs installed on the host and Destination Powerstore MPIO is not able to claim destination device as some other MPIO has already claimed it.  * Version_Unsupported - Agent detected that the OS or any other dependent component does not satisfy the version as expected by the it. 
type HostAgentStatusEnum string

// List of HostAgentStatusEnum
const (
	HOSTAGENTSTATUSENUM_UNKNOWN HostAgentStatusEnum = "Unknown"
	HOSTAGENTSTATUSENUM_RUNNING HostAgentStatusEnum = "Running"
	HOSTAGENTSTATUSENUM_CONFLICT_DETECTED HostAgentStatusEnum = "Conflict_Detected"
	HOSTAGENTSTATUSENUM_VERSION_UNSUPPORTED HostAgentStatusEnum = "Version_Unsupported"
)

// All allowed values of HostAgentStatusEnum enum
var AllowedHostAgentStatusEnumEnumValues = []HostAgentStatusEnum{
	"Unknown",
	"Running",
	"Conflict_Detected",
	"Version_Unsupported",
}

func (v *HostAgentStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HostAgentStatusEnum(value)
	for _, existing := range AllowedHostAgentStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HostAgentStatusEnum", value)
}

// NewHostAgentStatusEnumFromValue returns a pointer to a valid HostAgentStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHostAgentStatusEnumFromValue(v string) (*HostAgentStatusEnum, error) {
	ev := HostAgentStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HostAgentStatusEnum: valid values are %v", v, AllowedHostAgentStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HostAgentStatusEnum) IsValid() bool {
	for _, existing := range AllowedHostAgentStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostAgentStatusEnum value
func (v HostAgentStatusEnum) Ptr() *HostAgentStatusEnum {
	return &v
}

type NullableHostAgentStatusEnum struct {
	value *HostAgentStatusEnum
	isSet bool
}

func (v NullableHostAgentStatusEnum) Get() *HostAgentStatusEnum {
	return v.value
}

func (v *NullableHostAgentStatusEnum) Set(val *HostAgentStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableHostAgentStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableHostAgentStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostAgentStatusEnum(val *HostAgentStatusEnum) *NullableHostAgentStatusEnum {
	return &NullableHostAgentStatusEnum{value: val, isSet: true}
}

func (v NullableHostAgentStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostAgentStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

