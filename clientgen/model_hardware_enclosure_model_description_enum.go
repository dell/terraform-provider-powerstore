/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// HardwareEnclosureModelDescriptionEnum The Enclosure Model Description. Available on the Expansion_Enclosure hardware type. Current enclosure models are:  * ENS24_2_5_NVMe_Drive_Enc - ENS24 2.5\" NVMe Expansion Enclosure.  * 2_5_SAS_Drive_Enc - 2.5\" SAS Expansion Enclosure.  Was added in version 3.0.0.0.
type HardwareEnclosureModelDescriptionEnum string

// List of HardwareEnclosureModelDescriptionEnum
const (
	HARDWAREENCLOSUREMODELDESCRIPTIONENUM_ENS24_2_5_NVME_DRIVE_ENC HardwareEnclosureModelDescriptionEnum = "ENS24_2_5_NVMe_Drive_Enc"
	HARDWAREENCLOSUREMODELDESCRIPTIONENUM__2_5_SAS_DRIVE_ENC       HardwareEnclosureModelDescriptionEnum = "2_5_SAS_Drive_Enc"
)

// All allowed values of HardwareEnclosureModelDescriptionEnum enum
var AllowedHardwareEnclosureModelDescriptionEnumEnumValues = []HardwareEnclosureModelDescriptionEnum{
	"ENS24_2_5_NVMe_Drive_Enc",
	"2_5_SAS_Drive_Enc",
}

func (v *HardwareEnclosureModelDescriptionEnum) Value() string {
	return string(*v)
}
