/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

import (
	"time"
)

// MaintenanceWindowInstance This resource type has queriable association from appliance
type MaintenanceWindowInstance struct {
	// Unique identifier of the maintenance window.
	Id *string `json:"id,omitempty"`
	// Appliance id on which this maintenance window is configured.
	ApplianceId *string `json:"appliance_id,omitempty"`
	// Whether the maintenance window is active.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Time when the maintenance window will close (or did close).
	EndTime   *time.Time         `json:"end_time,omitempty"`
	Appliance *ApplianceInstance `json:"appliance,omitempty"`
}
