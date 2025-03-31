/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// SMBShareOfflineAvailabilityEnum Defines valid states of Offline Availability,    * Manual - Only specified files will be available offline.    * Documents - All files that users open will be available offline.    * Programs - Program will preferably run from the offline cache even when connected to the network. All files that users open will be available offline.    * None - Prevents clients from storing documents and programs in offline cache (default).
type SMBShareOfflineAvailabilityEnum string

// List of SMBShareOfflineAvailabilityEnum
const (
	SMBSHAREOFFLINEAVAILABILITYENUM_MANUAL    SMBShareOfflineAvailabilityEnum = "Manual"
	SMBSHAREOFFLINEAVAILABILITYENUM_DOCUMENTS SMBShareOfflineAvailabilityEnum = "Documents"
	SMBSHAREOFFLINEAVAILABILITYENUM_PROGRAMS  SMBShareOfflineAvailabilityEnum = "Programs"
	SMBSHAREOFFLINEAVAILABILITYENUM_NONE      SMBShareOfflineAvailabilityEnum = "None"
)

// All allowed values of SMBShareOfflineAvailabilityEnum enum
var AllowedSMBShareOfflineAvailabilityEnumEnumValues = []SMBShareOfflineAvailabilityEnum{
	"Manual",
	"Documents",
	"Programs",
	"None",
}

func (v *SMBShareOfflineAvailabilityEnum) Value() string {
	return string(*v)
}
