/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// VolumeTypeEnum Type of volume.  * Primary - A base object.  * Clone - A read-write object that shares storage with the object from which it is sourced.  * Snapshot - A read-only object created from a volume or clone.  * Proxy - A proxy object that is used for remote snapshot access.  Values was added in 3.5.0.0: Proxy.
type VolumeTypeEnum string

// List of VolumeTypeEnum
const (
	VOLUMETYPEENUM_PRIMARY  VolumeTypeEnum = "Primary"
	VOLUMETYPEENUM_CLONE    VolumeTypeEnum = "Clone"
	VOLUMETYPEENUM_SNAPSHOT VolumeTypeEnum = "Snapshot"
	VOLUMETYPEENUM_PROXY    VolumeTypeEnum = "Proxy"
)

// All allowed values of VolumeTypeEnum enum
var AllowedVolumeTypeEnumEnumValues = []VolumeTypeEnum{
	"Primary",
	"Clone",
	"Snapshot",
	"Proxy",
}

func (v *VolumeTypeEnum) Value() string {
	return string(*v)
}
