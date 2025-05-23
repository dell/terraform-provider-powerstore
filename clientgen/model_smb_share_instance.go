/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// SmbShareInstance This resource type has queriable association from file_system
type SmbShareInstance struct {
	// Id of the SMB Share.
	Id *string `json:"id,omitempty"`
	// The file system from which the share was created.
	FileSystemId *string `json:"file_system_id,omitempty"`
	// SMB share name.  This property supports case-insensitive filtering.
	Name *string `json:"name,omitempty"`
	// Local path to the file system or any existing sub-folder of the file system that is shared over the network. This path is relative to the NAS Server and must start with the filesystem's mountpoint path, which is the filesystem name. For example to share the top-level of a filesystem named svr1fs1, which is on the /svr1fs1 mountpoint of the NAS Server, use /svr1fs1 in the path parameter. SMB shares allow you to create multiple network shares for the same local path.
	Path *string `json:"path,omitempty"`
	// User defined SMB share description.
	Description *string `json:"description,omitempty"`
	// Indicates whether continuous availability for Server Message Block (SMB) 3.0 is enabled for the SMB Share. Values are: - true - Continuous availability for SMB 3.0 is enabled for the SMB Share. - false - Continuous availability for SMB 3.0 is disabled for the SMB Share.
	IsContinuousAvailabilityEnabled *bool `json:"is_continuous_availability_enabled,omitempty"`
	// Indicates whether encryption for Server Message Block (SMB) 3.0 is enabled at the shared folder level. Values are: - true - Encryption for SMB 3.0 is enabled. - false - Encryption for SMB 3.0 is disabled.
	IsEncryptionEnabled *bool `json:"is_encryption_enabled,omitempty"`
	// Indicates whether Access-based Enumeration (ABE) is enabled. ABE filters the list of available files and folders on a server to include only those to which the requesting user has access. Values are: - true - ABE is enabled. - false - ABE is disabled.
	IsABEEnabled *bool `json:"is_ABE_enabled,omitempty"`
	// Indicates whether BranchCache optimization is enabled. BranchCache optimization technology copies content from your main office or hosted cloud content servers and caches the content at branch office locations, allowing client computers at branch offices to access the content locally rather than over the WAN. Values are: - true - BranchCache is enabled. - false - BranchCache is disabled.
	IsBranchCacheEnabled *bool                            `json:"is_branch_cache_enabled,omitempty"`
	OfflineAvailability  *SMBShareOfflineAvailabilityEnum `json:"offline_availability,omitempty"`
	// The default UNIX umask for new files created on the Share. If not specified the umask defaults to 022.
	Umask *string `json:"umask,omitempty"`
	// Localized message string corresponding to offline_availability
	OfflineAvailabilityL10n *string             `json:"offline_availability_l10n,omitempty"`
	FileSystem              *FileSystemInstance `json:"file_system,omitempty"`
}
