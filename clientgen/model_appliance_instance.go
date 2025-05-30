/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// ApplianceInstance Properties of an appliance. This resource type has queriable associations from node, ip_pool_address, fsn, veth_port, virtual_volume, maintenance_window, fc_port, sas_port, eth_port, eth_be_port, software_installed, hardware, volume
type ApplianceInstance struct {
	// Unique identifier of the appliance.
	Id *string `json:"id,omitempty"`
	// Name of the appliance.  This property supports case-insensitive filtering.
	Name *string `json:"name,omitempty"`
	// Dell Service Tag.
	ServiceTag *string `json:"service_tag,omitempty"`
	// Express Service Code.
	ExpressServiceCode *string `json:"express_service_code,omitempty"`
	// Model of the appliance.
	Model *string            `json:"model,omitempty"`
	Mode  *ApplianceModeEnum `json:"mode,omitempty"`
	// The number of nodes deployed on an appliance. Was added in version 3.0.0.0.
	NodeCount                  *int32                          `json:"node_count,omitempty"`
	DriveFailureToleranceLevel *DriveFailureToleranceLevelEnum `json:"drive_failure_tolerance_level,omitempty"`
	StorageClass               *ApplianceStorageClassEnum      `json:"storage_class,omitempty"`
	// Is this a HyperConverged Appliance Was added in version 3.2.0.0.
	IsHyperConverged *bool `json:"is_hyper_converged,omitempty"`
	// Localized message string corresponding to mode Was added in version 4.0.0.0.
	ModeL10n *string `json:"mode_l10n,omitempty"`
	// Localized message string corresponding to drive_failure_tolerance_level Was added in version 2.0.0.0.
	DriveFailureToleranceLevelL10n *string `json:"drive_failure_tolerance_level_l10n,omitempty"`
	// Localized message string corresponding to storage_class Was added in version 4.0.0.0.
	StorageClassL10n *string `json:"storage_class_l10n,omitempty"`
	// This is the inverse of the resource type node association.
	Nodes []NodeInstance `json:"nodes,omitempty"`
	// This is the inverse of the resource type ip_pool_address association.
	IpPoolAddresses []IpPoolAddressInstance `json:"ip_pool_addresses,omitempty"`
	// This is the inverse of the resource type fsn association.
	Fsns []FsnInstance `json:"fsns,omitempty"`
	// This is the inverse of the resource type veth_port association.
	VethPorts []VethPortInstance `json:"veth_ports,omitempty"`
	// This is the inverse of the resource type virtual_volume association.
	VirtualVolumes []VirtualVolumeInstance `json:"virtual_volumes,omitempty"`
	// This is the inverse of the resource type maintenance_window association.
	MaintenanceWindows []MaintenanceWindowInstance `json:"maintenance_windows,omitempty"`
	// This is the inverse of the resource type fc_port association.
	FcPorts []FcPortInstance `json:"fc_ports,omitempty"`
	// This is the inverse of the resource type sas_port association.
	SasPorts []SasPortInstance `json:"sas_ports,omitempty"`
	// This is the inverse of the resource type eth_port association.
	EthPorts []EthPortInstance `json:"eth_ports,omitempty"`
	// This is the inverse of the resource type eth_be_port association.
	EthBePorts []EthBePortInstance `json:"eth_be_ports,omitempty"`
	// This is the inverse of the resource type software_installed association.
	SoftwareInstalled []SoftwareInstalledInstance `json:"software_installed,omitempty"`
	// This is the inverse of the resource type hardware association.
	Hardware []HardwareInstance `json:"hardware,omitempty"`
	// This is the inverse of the resource type volume association.
	Volumes []VolumeInstance `json:"volumes,omitempty"`
}
