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

// InitiatorProtocolTypeEnum Protocol type of the host initiator.  * iSCSI - An iSCSI host initiator.  * FC - A Fibre Channel host initiator.  * NVMe -  A NVMe host initiator. Single type for all NVMe Fabrics types (NVMe/FC, NVMe/TCP, NVMe/RoCEv2).  * NVMe_vVol - A vVol specific NVMe host initiator.  Values was added in 2.0.0.0: NVMe. Values was added in 3.0.0.0: NVMe_vVol.
type InitiatorProtocolTypeEnum string

// List of InitiatorProtocolTypeEnum
const (
	INITIATORPROTOCOLTYPEENUM_I_SCSI InitiatorProtocolTypeEnum = "iSCSI"
	INITIATORPROTOCOLTYPEENUM_FC InitiatorProtocolTypeEnum = "FC"
	INITIATORPROTOCOLTYPEENUM_NVME InitiatorProtocolTypeEnum = "NVMe"
	INITIATORPROTOCOLTYPEENUM_NVME_V_VOL InitiatorProtocolTypeEnum = "NVMe_vVol"
)

// All allowed values of InitiatorProtocolTypeEnum enum
var AllowedInitiatorProtocolTypeEnumEnumValues = []InitiatorProtocolTypeEnum{
	"iSCSI",
	"FC",
	"NVMe",
	"NVMe_vVol",
}

func (v *InitiatorProtocolTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InitiatorProtocolTypeEnum(value)
	for _, existing := range AllowedInitiatorProtocolTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InitiatorProtocolTypeEnum", value)
}

// NewInitiatorProtocolTypeEnumFromValue returns a pointer to a valid InitiatorProtocolTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInitiatorProtocolTypeEnumFromValue(v string) (*InitiatorProtocolTypeEnum, error) {
	ev := InitiatorProtocolTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InitiatorProtocolTypeEnum: valid values are %v", v, AllowedInitiatorProtocolTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InitiatorProtocolTypeEnum) IsValid() bool {
	for _, existing := range AllowedInitiatorProtocolTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InitiatorProtocolTypeEnum value
func (v InitiatorProtocolTypeEnum) Ptr() *InitiatorProtocolTypeEnum {
	return &v
}

type NullableInitiatorProtocolTypeEnum struct {
	value *InitiatorProtocolTypeEnum
	isSet bool
}

func (v NullableInitiatorProtocolTypeEnum) Get() *InitiatorProtocolTypeEnum {
	return v.value
}

func (v *NullableInitiatorProtocolTypeEnum) Set(val *InitiatorProtocolTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableInitiatorProtocolTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableInitiatorProtocolTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitiatorProtocolTypeEnum(val *InitiatorProtocolTypeEnum) *NullableInitiatorProtocolTypeEnum {
	return &NullableInitiatorProtocolTypeEnum{value: val, isSet: true}
}

func (v NullableInitiatorProtocolTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitiatorProtocolTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

