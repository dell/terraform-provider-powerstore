/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// FileSystemHostIoSizeEnum Typical size of writes from the server or other computer using the VMware file system to the storage system. This setting is used to match the storage block size to the I/O of the primary application using the storage, which can optimize performance. By default Host IO size will be 8K size for VMware filesystem. This attribute only applies to VMware config type filesystems. * VMware_8K - Host I/O size is 8K for vmware datastore purpose. * VMware_16K - Host I/O size is 16K for vmware datastore purpose. * VMware_32K - Host I/O size is 32K for vmware datastore purpose. * VMware_64K - Host I/O size is 64K for vmware datastore purpose.  Was added in version 3.0.0.0.
type FileSystemHostIoSizeEnum string

// List of FileSystemHostIoSizeEnum
const (
	FILESYSTEMHOSTIOSIZEENUM__8_K  FileSystemHostIoSizeEnum = "VMware_8K"
	FILESYSTEMHOSTIOSIZEENUM__16_K FileSystemHostIoSizeEnum = "VMware_16K"
	FILESYSTEMHOSTIOSIZEENUM__32_K FileSystemHostIoSizeEnum = "VMware_32K"
	FILESYSTEMHOSTIOSIZEENUM__64_K FileSystemHostIoSizeEnum = "VMware_64K"
)

// All allowed values of FileSystemHostIoSizeEnum enum
var AllowedFileSystemHostIoSizeEnumEnumValues = []FileSystemHostIoSizeEnum{
	"VMware_8K",
	"VMware_16K",
	"VMware_32K",
	"VMware_64K",
}

func (v *FileSystemHostIoSizeEnum) Value() string {
	return string(*v)
}
