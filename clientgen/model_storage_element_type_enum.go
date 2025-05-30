/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// StorageElementTypeEnum Storage element type being replicated: * volume                  - Replicating storage element type for a volume. * virtual_volume          - Replicating storage element type for a virtual volume. * file_system             - Replicating storage element type for a file system.  Values was added in 3.5.0.0: file_system.
type StorageElementTypeEnum string

// List of StorageElementTypeEnum
const (
	STORAGEELEMENTTYPEENUM_VOLUME         StorageElementTypeEnum = "volume"
	STORAGEELEMENTTYPEENUM_VIRTUAL_VOLUME StorageElementTypeEnum = "virtual_volume"
	STORAGEELEMENTTYPEENUM_FILE_SYSTEM    StorageElementTypeEnum = "file_system"
)

// All allowed values of StorageElementTypeEnum enum
var AllowedStorageElementTypeEnumEnumValues = []StorageElementTypeEnum{
	"volume",
	"virtual_volume",
	"file_system",
}

func (v *StorageElementTypeEnum) Value() string {
	return string(*v)
}
