/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// TransitConnectionStatusEnum Possible transit connection statuses: * Login_Success               - Login to target successful. * Authentication_Failure      - Failed to authenticate the connection. * Connection_Refused          - Connection was refused. * Login_Timeout               - Login to target timed out. * Network_Error - Network error * General_Failure             - Other failure not listed. * Login_Success_No_Ports      - Login successful after discovery failure. Used only for PS EqualLogic systems. * Discovery_Success           - Discovery of target IP successful. * Discovery_Authentication_Failure - Authentication failure during discovery of target. * Discovery_Connection_Refused     - Connection was refused during discovery of target. * Discovery_Timeout                - Discovery of target timed out. * Operation_Timeout               - Operations on the connection timed out. Used only for PowerProtect DD systems.  Values was added in 3.5.0.0: Operation_Timeout.
type TransitConnectionStatusEnum string

// List of TransitConnectionStatusEnum
const (
	TRANSITCONNECTIONSTATUSENUM_LOGIN_SUCCESS                    TransitConnectionStatusEnum = "Login_Success"
	TRANSITCONNECTIONSTATUSENUM_AUTHENTICATION_FAILURE           TransitConnectionStatusEnum = "Authentication_Failure"
	TRANSITCONNECTIONSTATUSENUM_CONNECTION_REFUSED               TransitConnectionStatusEnum = "Connection_Refused"
	TRANSITCONNECTIONSTATUSENUM_LOGIN_TIMEOUT                    TransitConnectionStatusEnum = "Login_Timeout"
	TRANSITCONNECTIONSTATUSENUM_NETWORK_ERROR                    TransitConnectionStatusEnum = "Network_Error"
	TRANSITCONNECTIONSTATUSENUM_GENERAL_FAILURE                  TransitConnectionStatusEnum = "General_Failure"
	TRANSITCONNECTIONSTATUSENUM_LOGIN_SUCCESS_NO_PORTS           TransitConnectionStatusEnum = "Login_Success_No_Ports"
	TRANSITCONNECTIONSTATUSENUM_DISCOVERY_SUCCESS                TransitConnectionStatusEnum = "Discovery_Success"
	TRANSITCONNECTIONSTATUSENUM_DISCOVERY_AUTHENTICATION_FAILURE TransitConnectionStatusEnum = "Discovery_Authentication_Failure"
	TRANSITCONNECTIONSTATUSENUM_DISCOVERY_CONNECTION_REFUSED     TransitConnectionStatusEnum = "Discovery_Connection_Refused"
	TRANSITCONNECTIONSTATUSENUM_DISCOVERY_TIMEOUT                TransitConnectionStatusEnum = "Discovery_Timeout"
	TRANSITCONNECTIONSTATUSENUM_OPERATION_TIMEOUT                TransitConnectionStatusEnum = "Operation_Timeout"
)

// All allowed values of TransitConnectionStatusEnum enum
var AllowedTransitConnectionStatusEnumEnumValues = []TransitConnectionStatusEnum{
	"Login_Success",
	"Authentication_Failure",
	"Connection_Refused",
	"Login_Timeout",
	"Network_Error",
	"General_Failure",
	"Login_Success_No_Ports",
	"Discovery_Success",
	"Discovery_Authentication_Failure",
	"Discovery_Connection_Refused",
	"Discovery_Timeout",
	"Operation_Timeout",
}

func (v *TransitConnectionStatusEnum) Value() string {
	return string(*v)
}
