/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// NodeInstance This resource type has queriable associations from appliance, ip_pool_address, veth_port
type NodeInstance struct {
	// Unique identifier of the node.
	Id *string `json:"id,omitempty"`
	// Slot number of the node.
	Slot *int32 `json:"slot,omitempty"`
	// Unique identifier of the appliance to which the node belongs.
	ApplianceId *string            `json:"appliance_id,omitempty"`
	Appliance   *ApplianceInstance `json:"appliance,omitempty"`
	// This is the inverse of the resource type ip_pool_address association.
	IpPoolAddresses []IpPoolAddressInstance `json:"ip_pool_addresses,omitempty"`
	// This is the inverse of the resource type veth_port association.
	VethPorts []VethPortInstance `json:"veth_ports,omitempty"`
}
