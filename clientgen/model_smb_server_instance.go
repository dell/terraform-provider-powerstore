/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// SmbServerInstance This resource type has queriable association from nas_server
type SmbServerInstance struct {
	// Unique identifier of the SMB server.
	Id *string `json:"id,omitempty"`
	// Unique identifier of the NAS server.
	NasServerId *string `json:"nas_server_id,omitempty"`
	// DNS name of the associated computer account when the SMB server is joined to an Active Directory domain. This name's minimum length is 2 characters, it is limited to 63 bytes and must not contain the following characters -   - comma (.)   - tilde (~)   - colon (:)   - exclamation point (!)   - at sign (@)   - number sign (#)   - dollar sign ($)   - percent (%)   - caret (^)   - ampersand (&)   - apostrophe (')   - period (.) - note that if you enter string with period only the first word will be kept   - parentheses (())   - braces ({})   - underscore (_)   - white space (blank) as defined by the Microsoft naming convention (see https://support.microsoft.com/en-us/help/909264/)
	ComputerName *string `json:"computer_name,omitempty"`
	// Domain name where SMB server is registered in Active Directory, if applicable.
	Domain *string `json:"domain,omitempty"`
	// NetBIOS name is the network name of the standalone SMB server. SMB server joined to Active Directory also have NetBIOS Name, defaulted to the 15 first characters of the computerName attribute. Administrators can specify a custom NetBIOS Name for a SMB server using this attribute. NetBIOS Name are limited to 15 characters and cannot contain the following characters -   - backslash (\\)   - slash mark (/)   - colon (:)   - asterisk (*)   - question mark (?)   - quotation mark (\"\")   - less than sign (<)   - greater than sign (>)   - vertical bar (|) as definied by the Microsoft naming convention (see https://support.microsoft.com/en-us/help/909264/)
	NetbiosName *string `json:"netbios_name,omitempty"`
	// Applies to stand-alone SMB servers only. Windows network workgroup for the SMB server. Workgroup names are limited to 15 alphanumeric ASCII characters.
	Workgroup *string `json:"workgroup,omitempty"`
	// Description of the SMB server.
	Description *string `json:"description,omitempty"`
	// Indicates whether the SMB server is standalone. Values are: - true - SMB server is standalone. - false - SMB server is a domain SMB server to be joined to the Active Directory.
	IsStandalone *bool `json:"is_standalone,omitempty"`
	// Indicates whether the SMB server is joined to the Active Directory. Always false for standalone SMB servers.
	IsJoined  *bool              `json:"is_joined,omitempty"`
	NasServer *NasServerInstance `json:"nas_server,omitempty"`
}
