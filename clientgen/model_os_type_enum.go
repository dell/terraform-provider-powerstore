/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// OSTypeEnum Operating system of the host. Values are: * Windows * Linux * ESXi * AIX * HP-UX * Solaris
type OSTypeEnum string

// List of OSTypeEnum
const (
	OSTYPEENUM_WINDOWS OSTypeEnum = "Windows"
	OSTYPEENUM_LINUX   OSTypeEnum = "Linux"
	OSTYPEENUM_ESXI    OSTypeEnum = "ESXi"
	OSTYPEENUM_AIX     OSTypeEnum = "AIX"
	OSTYPEENUM_HP_UX   OSTypeEnum = "HP-UX"
	OSTYPEENUM_SOLARIS OSTypeEnum = "Solaris"
)

// All allowed values of OSTypeEnum enum
var AllowedOSTypeEnumEnumValues = []OSTypeEnum{
	"Windows",
	"Linux",
	"ESXi",
	"AIX",
	"HP-UX",
	"Solaris",
}

func (v *OSTypeEnum) Value() string {
	return string(*v)
}
