/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// FileSystemFolderRenamePolicyEnum File system folder rename policies for the file system with multiprotocol access enabled. These policies control whether the directory can be renamed from NFS or SMB clients when at least one file is opened in the directory, or in one of its child directories.  * All_Allowed - All protocols are allowed to rename directories without any restrictions.  * SMB_Forbidden - A directory rename from the SMB protocol will be denied if at least one file is opened in the directory or in one of its child directories.  * All_Forbidden - Any directory rename request will be denied regardless of the protocol used, if at least one file is opened in the directory or in one of its child directories.
type FileSystemFolderRenamePolicyEnum string

// List of FileSystemFolderRenamePolicyEnum
const (
	FILESYSTEMFOLDERRENAMEPOLICYENUM_ALL_ALLOWED   FileSystemFolderRenamePolicyEnum = "All_Allowed"
	FILESYSTEMFOLDERRENAMEPOLICYENUM_SMB_FORBIDDEN FileSystemFolderRenamePolicyEnum = "SMB_Forbidden"
	FILESYSTEMFOLDERRENAMEPOLICYENUM_ALL_FORBIDDEN FileSystemFolderRenamePolicyEnum = "All_Forbidden"
)

// All allowed values of FileSystemFolderRenamePolicyEnum enum
var AllowedFileSystemFolderRenamePolicyEnumEnumValues = []FileSystemFolderRenamePolicyEnum{
	"All_Allowed",
	"SMB_Forbidden",
	"All_Forbidden",
}

func (v *FileSystemFolderRenamePolicyEnum) Value() string {
	return string(*v)
}
