/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// PpddStorageUnitDetailsInstance PowerProtect DD details to register with Powerstore as a remote system.  Was added in version 3.5.0.0.  Filtering on the fields of this embedded resource is not supported.
type PpddStorageUnitDetailsInstance struct {
	// DDBoost address of the PowerProtect DD remote system. IPv4 and FQDN is the supported DDBoost address.
	DdBoostAddress *string `json:"dd_boost_address,omitempty"`
	// Name of the storage unit in the PowerProtect DD.
	StorageUnitName *string `json:"storage_unit_name,omitempty"`
	// Username used to access the storage unit of a PowerProtect DD remote system.
	DdBoostUsername *string `json:"dd_boost_username,omitempty"`
	// Enable or Disable encryption for all backup session from Powerstore to a storage unit in PowerProtect DD.
	IsDataEncryptionEnabled *bool `json:"is_data_encryption_enabled,omitempty"`
}
