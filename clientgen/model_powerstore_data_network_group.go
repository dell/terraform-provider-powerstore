/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// PowerstoreDataNetworkGroup Group of local and remote PowerStore storage networks to use for configuring data connections for replication data transfers.  Was added in version 4.0.0.0.
type PowerstoreDataNetworkGroup struct {
	// Unique identifier of the data network group.
	Id *string `json:"id,omitempty"`
	// User or system generated name for group of local and remote replication networks.
	Name *string `json:"name,omitempty"`
	// List of local storage networks that are defined to be used for replication data transfer.
	LocalPowerstoreNetworks []PowerstoreNetworkInfo `json:"local_powerstore_networks,omitempty"`
	// List of remote storage networks that are defined to be used for replication data transfer.
	RemotePowerstoreNetworks []PowerstoreNetworkInfo `json:"remote_powerstore_networks,omitempty"`
}
