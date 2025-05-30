/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// VirtualVolumeTypeEnum The logical type of a virtual volume. Values are: * Primary * Snapshot * Prepared_Snapshot * Clone * Fast_Clone
type VirtualVolumeTypeEnum string

// List of VirtualVolumeTypeEnum
const (
	VIRTUALVOLUMETYPEENUM_PRIMARY           VirtualVolumeTypeEnum = "Primary"
	VIRTUALVOLUMETYPEENUM_SNAPSHOT          VirtualVolumeTypeEnum = "Snapshot"
	VIRTUALVOLUMETYPEENUM_PREPARED_SNAPSHOT VirtualVolumeTypeEnum = "Prepared_Snapshot"
	VIRTUALVOLUMETYPEENUM_CLONE             VirtualVolumeTypeEnum = "Clone"
	VIRTUALVOLUMETYPEENUM_FAST_CLONE        VirtualVolumeTypeEnum = "Fast_Clone"
)

// All allowed values of VirtualVolumeTypeEnum enum
var AllowedVirtualVolumeTypeEnumEnumValues = []VirtualVolumeTypeEnum{
	"Primary",
	"Snapshot",
	"Prepared_Snapshot",
	"Clone",
	"Fast_Clone",
}

func (v *VirtualVolumeTypeEnum) Value() string {
	return string(*v)
}
