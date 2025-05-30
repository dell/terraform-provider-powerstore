/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// EthPortSpeedEnum Supported Ethernet front-end port transmission speeds. For the current_speed attribute, these values show the current transmission speed on the port. For the requested_speed attribute, these values show the transmission speed set by the user. A requested speed of Auto means that the current speed value will be automatically detected. If this file is updated, also update FrontEndPortSpeedEnum.yaml * Auto - the speed value is automatically detected * 10_Mbps - 10 Megabits per second * 100_Mbps - 100 Megabits per second * 1_Gbps - 1 Gigabits per second * 10_Gbps - 10 Gigabits per second * 25_Gbps - 25 Gigabits per second * 40_Gbps - 40 Gigabits per second * 100_Gbps - 100 Gigabits per second  Values was added in 3.0.0.0: 100_Gbps.
type EthPortSpeedEnum string

// List of EthPortSpeedEnum
const (
	ETHPORTSPEEDENUM_AUTO      EthPortSpeedEnum = "Auto"
	ETHPORTSPEEDENUM__10_MBPS  EthPortSpeedEnum = "10_Mbps"
	ETHPORTSPEEDENUM__100_MBPS EthPortSpeedEnum = "100_Mbps"
	ETHPORTSPEEDENUM__1_GBPS   EthPortSpeedEnum = "1_Gbps"
	ETHPORTSPEEDENUM__10_GBPS  EthPortSpeedEnum = "10_Gbps"
	ETHPORTSPEEDENUM__25_GBPS  EthPortSpeedEnum = "25_Gbps"
	ETHPORTSPEEDENUM__40_GBPS  EthPortSpeedEnum = "40_Gbps"
	ETHPORTSPEEDENUM__100_GBPS EthPortSpeedEnum = "100_Gbps"
)

// All allowed values of EthPortSpeedEnum enum
var AllowedEthPortSpeedEnumEnumValues = []EthPortSpeedEnum{
	"Auto",
	"10_Mbps",
	"100_Mbps",
	"1_Gbps",
	"10_Gbps",
	"25_Gbps",
	"40_Gbps",
	"100_Gbps",
}

func (v *EthPortSpeedEnum) Value() string {
	return string(*v)
}
