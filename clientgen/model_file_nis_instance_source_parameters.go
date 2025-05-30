/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// FileNisInstanceSourceParameters Information about the corresponding source NAS Server's File NIS settings. Only populated when is_destination_override_enabled flag is set to true. Was added in version 3.0.0.0.  Filtering on the fields of this embedded resource is not supported.
type FileNisInstanceSourceParameters struct {
	// The list of NIS server IP addresses.
	IpAddresses []string `json:"ip_addresses,omitempty"`
}
