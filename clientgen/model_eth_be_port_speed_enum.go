/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// EthBEPortSpeedEnum Ethernet Backend port transmission speed. * 25_Gbps - 25 Gigabits per second. * 100_Gbps - 100 Gigabits per second.  Was added in version 3.0.0.0.
type EthBEPortSpeedEnum string

// List of EthBEPortSpeedEnum
const (
	ETHBEPORTSPEEDENUM__25_GBPS  EthBEPortSpeedEnum = "25_Gbps"
	ETHBEPORTSPEEDENUM__100_GBPS EthBEPortSpeedEnum = "100_Gbps"
)

// All allowed values of EthBEPortSpeedEnum enum
var AllowedEthBEPortSpeedEnumEnumValues = []EthBEPortSpeedEnum{
	"25_Gbps",
	"100_Gbps",
}

func (v *EthBEPortSpeedEnum) Value() string {
	return string(*v)
}
