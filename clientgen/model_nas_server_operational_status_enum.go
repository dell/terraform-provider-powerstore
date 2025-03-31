/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// NASServerOperationalStatusEnum NAS server operational status: * Stopped - NAS server is stopped. * Starting - NAS server is starting. * Started - NAS server is started. * Stopping - NAS server is stopping. * Failover - NAS server has failed over. * Degraded - NAS server is degraded (running without backup). * Unknown - NAS server state is unknown.
type NASServerOperationalStatusEnum string

// List of NASServerOperationalStatusEnum
const (
	NASSERVEROPERATIONALSTATUSENUM_STOPPED  NASServerOperationalStatusEnum = "Stopped"
	NASSERVEROPERATIONALSTATUSENUM_STARTING NASServerOperationalStatusEnum = "Starting"
	NASSERVEROPERATIONALSTATUSENUM_STARTED  NASServerOperationalStatusEnum = "Started"
	NASSERVEROPERATIONALSTATUSENUM_STOPPING NASServerOperationalStatusEnum = "Stopping"
	NASSERVEROPERATIONALSTATUSENUM_FAILOVER NASServerOperationalStatusEnum = "Failover"
	NASSERVEROPERATIONALSTATUSENUM_DEGRADED NASServerOperationalStatusEnum = "Degraded"
	NASSERVEROPERATIONALSTATUSENUM_UNKNOWN  NASServerOperationalStatusEnum = "Unknown"
)

// All allowed values of NASServerOperationalStatusEnum enum
var AllowedNASServerOperationalStatusEnumEnumValues = []NASServerOperationalStatusEnum{
	"Stopped",
	"Starting",
	"Started",
	"Stopping",
	"Failover",
	"Degraded",
	"Unknown",
}

func (v *NASServerOperationalStatusEnum) Value() string {
	return string(*v)
}
