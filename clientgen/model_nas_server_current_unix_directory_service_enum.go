/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// NASServerCurrentUnixDirectoryServiceEnum Define the Unix directory service used for looking up identity information for Unix such as UIDs, GIDs, net groups, and so on. Values are: * None * NIS * LDAP * Local_Files * Local_Then_NIS * Local_Then_LDAP
type NASServerCurrentUnixDirectoryServiceEnum string

// List of NASServerCurrentUnixDirectoryServiceEnum
const (
	NASSERVERCURRENTUNIXDIRECTORYSERVICEENUM_NONE            NASServerCurrentUnixDirectoryServiceEnum = "None"
	NASSERVERCURRENTUNIXDIRECTORYSERVICEENUM_NIS             NASServerCurrentUnixDirectoryServiceEnum = "NIS"
	NASSERVERCURRENTUNIXDIRECTORYSERVICEENUM_LDAP            NASServerCurrentUnixDirectoryServiceEnum = "LDAP"
	NASSERVERCURRENTUNIXDIRECTORYSERVICEENUM_LOCAL_FILES     NASServerCurrentUnixDirectoryServiceEnum = "Local_Files"
	NASSERVERCURRENTUNIXDIRECTORYSERVICEENUM_LOCAL_THEN_NIS  NASServerCurrentUnixDirectoryServiceEnum = "Local_Then_NIS"
	NASSERVERCURRENTUNIXDIRECTORYSERVICEENUM_LOCAL_THEN_LDAP NASServerCurrentUnixDirectoryServiceEnum = "Local_Then_LDAP"
)

// All allowed values of NASServerCurrentUnixDirectoryServiceEnum enum
var AllowedNASServerCurrentUnixDirectoryServiceEnumEnumValues = []NASServerCurrentUnixDirectoryServiceEnum{
	"None",
	"NIS",
	"LDAP",
	"Local_Files",
	"Local_Then_NIS",
	"Local_Then_LDAP",
}

func (v *NASServerCurrentUnixDirectoryServiceEnum) Value() string {
	return string(*v)
}
