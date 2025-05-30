/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// HardwareModelTypeEnum Desired model to which the appliance should be upgraded.  * 1200T  * 3200T  * 5200T  * 9200T  Was added in version 3.6.0.0.
type HardwareModelTypeEnum string

// List of HardwareModelTypeEnum
const (
	HARDWAREMODELTYPEENUM__1200_T HardwareModelTypeEnum = "1200T"
	HARDWAREMODELTYPEENUM__3200_T HardwareModelTypeEnum = "3200T"
	HARDWAREMODELTYPEENUM__5200_T HardwareModelTypeEnum = "5200T"
	HARDWAREMODELTYPEENUM__9200_T HardwareModelTypeEnum = "9200T"
)

// All allowed values of HardwareModelTypeEnum enum
var AllowedHardwareModelTypeEnumEnumValues = []HardwareModelTypeEnum{
	"1200T",
	"3200T",
	"5200T",
	"9200T",
}

func (v *HardwareModelTypeEnum) Value() string {
	return string(*v)
}
