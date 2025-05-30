/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// FcPortInstance Properties of the FC front-end port. Values was added in 2.0.0.0: wwn_nvme. Values was added in 3.0.0.0: is_in_use. This resource type has queriable associations from appliance, hardware, fc_port
type FcPortInstance struct {
	// Unique identifier of the port.
	Id *string `json:"id,omitempty"`
	// Name of the port.  This property supports case-insensitive filtering.
	Name *string `json:"name,omitempty"`
	// Unique identifier of the appliance containing the port.
	ApplianceId *string `json:"appliance_id,omitempty"`
	// Unique identifier of the hardware instance of type 'Node' containing the port.
	NodeId *string `json:"node_id,omitempty"`
	// World Wide Name (WWN) of the port.  More specifically, it's the World Wide Port Name (WWPN) used for SCSI Protocol.
	Wwn *string `json:"wwn,omitempty"`
	// World Wide Name (WWN) of NVME port. Was added in version 2.0.0.0.
	WwnNvme *string `json:"wwn_nvme,omitempty"`
	// World Wide Name (WWN) of the Node of the port.  More specifically, it's the World Wide Node Name (WWNN). Was added in version 3.0.0.0.
	WwnNode *string `json:"wwn_node,omitempty"`
	// Indicates whether the port's link is up. Values are: * true - Link is up. * false - Link is down.
	IsLinkUp *bool `json:"is_link_up,omitempty"`
	// Indicates whether the port is in use. Values are: * true - Is in use. * false - Is not in use.  Was added in version 3.0.0.0.
	IsInUse *bool `json:"is_in_use,omitempty"`
	// List of supported transmission speeds for the port.
	SupportedSpeeds []FcPortSpeedEnum `json:"supported_speeds,omitempty"`
	CurrentSpeed    *FcPortSpeedEnum  `json:"current_speed,omitempty"`
	RequestedSpeed  *FcPortSpeedEnum  `json:"requested_speed,omitempty"`
	// Unique identifier of the hardware instance of type 'SFP' (Small Form-factor Pluggable) inserted into the port.
	SfpId *string `json:"sfp_id,omitempty"`
	// Unique identifier of the hardware instance of type 'IO_Module' handling the port. Was deprecated in version 2.0.0.0.
	IoModuleId *string `json:"io_module_id,omitempty"`
	// Unique identifier of the parent hardware instance handling the port. @added(Foothills) Was added in version 2.0.0.0.
	HardwareParentId *string `json:"hardware_parent_id,omitempty"`
	// Index of the port in the IO module.
	PortIndex         *int32                          `json:"port_index,omitempty"`
	PortConnectorType *FrontEndPortConnectionTypeEnum `json:"port_connector_type,omitempty"`
	// Unique identifier of the partner port.
	PartnerId *string `json:"partner_id,omitempty"`
	// List of supported protocols for the port. Was added in version 2.0.0.0.
	Protocols  []FcPortProtocolEnum `json:"protocols,omitempty"`
	StaleState *PortStaleStateEnum  `json:"stale_state,omitempty"`
	ScsiMode   *FcPortScsiModeEnum  `json:"scsi_mode,omitempty"`
	// Localized message array corresponding to supported_speeds
	SupportedSpeedsL10n []string `json:"supported_speeds_l10n,omitempty"`
	// Localized message string corresponding to current_speed
	CurrentSpeedL10n *string `json:"current_speed_l10n,omitempty"`
	// Localized message string corresponding to requested_speed
	RequestedSpeedL10n *string `json:"requested_speed_l10n,omitempty"`
	// Localized message string corresponding to port_connector_type
	PortConnectorTypeL10n *string `json:"port_connector_type_l10n,omitempty"`
	// Localized message array corresponding to protocols Was added in version 2.0.0.0.
	ProtocolsL10n []string `json:"protocols_l10n,omitempty"`
	// Localized message string corresponding to stale_state Was added in version 2.0.0.0.
	StaleStateL10n *string `json:"stale_state_l10n,omitempty"`
	// Localized message string corresponding to scsi_mode Was added in version 3.0.0.0.
	ScsiModeL10n   *string            `json:"scsi_mode_l10n,omitempty"`
	Appliance      *ApplianceInstance `json:"appliance,omitempty"`
	Node           *HardwareInstance  `json:"node,omitempty"`
	Sfp            *HardwareInstance  `json:"sfp,omitempty"`
	IoModule       *HardwareInstance  `json:"io_module,omitempty"`
	HardwareParent *HardwareInstance  `json:"hardware_parent,omitempty"`
	Partner        *FcPortInstance    `json:"partner,omitempty"`
}
