/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// FileSystemLockingPolicyEnum File system locking policies. These policy choices control whether the NFSv4 range locks are honored. Because NFSv3 is advisory by design, this policy specifies that the NFSv4 locking feature behaves like NFSv3 (advisory mode), for backward compatiblity with applications expecting an advisory locking scheme.    * Advisory - No lock checking for NFS and honor SMB lock range only for SMB.  * Mandatory - Honor SMB and NFS lock range.
type FileSystemLockingPolicyEnum string

// List of FileSystemLockingPolicyEnum
const (
	FILESYSTEMLOCKINGPOLICYENUM_ADVISORY  FileSystemLockingPolicyEnum = "Advisory"
	FILESYSTEMLOCKINGPOLICYENUM_MANDATORY FileSystemLockingPolicyEnum = "Mandatory"
)

// All allowed values of FileSystemLockingPolicyEnum enum
var AllowedFileSystemLockingPolicyEnumEnumValues = []FileSystemLockingPolicyEnum{
	"Advisory",
	"Mandatory",
}

func (v *FileSystemLockingPolicyEnum) Value() string {
	return string(*v)
}
