/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

import (
	"encoding/json"
)

// checks if the ApplianceInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplianceInstance{}

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
	Model *string `json:"model,omitempty"`
	Mode *ApplianceModeEnum `json:"mode,omitempty"`
	// The number of nodes deployed on an appliance. Was added in version 3.0.0.0.
	NodeCount *int32 `json:"node_count,omitempty"`
	DriveFailureToleranceLevel *DriveFailureToleranceLevelEnum `json:"drive_failure_tolerance_level,omitempty"`
	StorageClass *ApplianceStorageClassEnum `json:"storage_class,omitempty"`
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

// NewApplianceInstance instantiates a new ApplianceInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceInstance() *ApplianceInstance {
	this := ApplianceInstance{}
	var nodeCount int32 = 2
	this.NodeCount = &nodeCount
	return &this
}

// NewApplianceInstanceWithDefaults instantiates a new ApplianceInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceInstanceWithDefaults() *ApplianceInstance {
	this := ApplianceInstance{}
	var nodeCount int32 = 2
	this.NodeCount = &nodeCount
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplianceInstance) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplianceInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplianceInstance) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplianceInstance) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplianceInstance) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplianceInstance) SetName(v string) {
	o.Name = &v
}

// GetServiceTag returns the ServiceTag field value if set, zero value otherwise.
func (o *ApplianceInstance) GetServiceTag() string {
	if o == nil || IsNil(o.ServiceTag) {
		var ret string
		return ret
	}
	return *o.ServiceTag
}

// GetServiceTagOk returns a tuple with the ServiceTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetServiceTagOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceTag) {
		return nil, false
	}
	return o.ServiceTag, true
}

// HasServiceTag returns a boolean if a field has been set.
func (o *ApplianceInstance) HasServiceTag() bool {
	if o != nil && !IsNil(o.ServiceTag) {
		return true
	}

	return false
}

// SetServiceTag gets a reference to the given string and assigns it to the ServiceTag field.
func (o *ApplianceInstance) SetServiceTag(v string) {
	o.ServiceTag = &v
}

// GetExpressServiceCode returns the ExpressServiceCode field value if set, zero value otherwise.
func (o *ApplianceInstance) GetExpressServiceCode() string {
	if o == nil || IsNil(o.ExpressServiceCode) {
		var ret string
		return ret
	}
	return *o.ExpressServiceCode
}

// GetExpressServiceCodeOk returns a tuple with the ExpressServiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetExpressServiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpressServiceCode) {
		return nil, false
	}
	return o.ExpressServiceCode, true
}

// HasExpressServiceCode returns a boolean if a field has been set.
func (o *ApplianceInstance) HasExpressServiceCode() bool {
	if o != nil && !IsNil(o.ExpressServiceCode) {
		return true
	}

	return false
}

// SetExpressServiceCode gets a reference to the given string and assigns it to the ExpressServiceCode field.
func (o *ApplianceInstance) SetExpressServiceCode(v string) {
	o.ExpressServiceCode = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ApplianceInstance) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ApplianceInstance) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *ApplianceInstance) SetModel(v string) {
	o.Model = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ApplianceInstance) GetMode() ApplianceModeEnum {
	if o == nil || IsNil(o.Mode) {
		var ret ApplianceModeEnum
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetModeOk() (*ApplianceModeEnum, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ApplianceInstance) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given ApplianceModeEnum and assigns it to the Mode field.
func (o *ApplianceInstance) SetMode(v ApplianceModeEnum) {
	o.Mode = &v
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise.
func (o *ApplianceInstance) GetNodeCount() int32 {
	if o == nil || IsNil(o.NodeCount) {
		var ret int32
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetNodeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.NodeCount) {
		return nil, false
	}
	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *ApplianceInstance) HasNodeCount() bool {
	if o != nil && !IsNil(o.NodeCount) {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int32 and assigns it to the NodeCount field.
func (o *ApplianceInstance) SetNodeCount(v int32) {
	o.NodeCount = &v
}

// GetDriveFailureToleranceLevel returns the DriveFailureToleranceLevel field value if set, zero value otherwise.
func (o *ApplianceInstance) GetDriveFailureToleranceLevel() DriveFailureToleranceLevelEnum {
	if o == nil || IsNil(o.DriveFailureToleranceLevel) {
		var ret DriveFailureToleranceLevelEnum
		return ret
	}
	return *o.DriveFailureToleranceLevel
}

// GetDriveFailureToleranceLevelOk returns a tuple with the DriveFailureToleranceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetDriveFailureToleranceLevelOk() (*DriveFailureToleranceLevelEnum, bool) {
	if o == nil || IsNil(o.DriveFailureToleranceLevel) {
		return nil, false
	}
	return o.DriveFailureToleranceLevel, true
}

// HasDriveFailureToleranceLevel returns a boolean if a field has been set.
func (o *ApplianceInstance) HasDriveFailureToleranceLevel() bool {
	if o != nil && !IsNil(o.DriveFailureToleranceLevel) {
		return true
	}

	return false
}

// SetDriveFailureToleranceLevel gets a reference to the given DriveFailureToleranceLevelEnum and assigns it to the DriveFailureToleranceLevel field.
func (o *ApplianceInstance) SetDriveFailureToleranceLevel(v DriveFailureToleranceLevelEnum) {
	o.DriveFailureToleranceLevel = &v
}

// GetStorageClass returns the StorageClass field value if set, zero value otherwise.
func (o *ApplianceInstance) GetStorageClass() ApplianceStorageClassEnum {
	if o == nil || IsNil(o.StorageClass) {
		var ret ApplianceStorageClassEnum
		return ret
	}
	return *o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetStorageClassOk() (*ApplianceStorageClassEnum, bool) {
	if o == nil || IsNil(o.StorageClass) {
		return nil, false
	}
	return o.StorageClass, true
}

// HasStorageClass returns a boolean if a field has been set.
func (o *ApplianceInstance) HasStorageClass() bool {
	if o != nil && !IsNil(o.StorageClass) {
		return true
	}

	return false
}

// SetStorageClass gets a reference to the given ApplianceStorageClassEnum and assigns it to the StorageClass field.
func (o *ApplianceInstance) SetStorageClass(v ApplianceStorageClassEnum) {
	o.StorageClass = &v
}

// GetIsHyperConverged returns the IsHyperConverged field value if set, zero value otherwise.
func (o *ApplianceInstance) GetIsHyperConverged() bool {
	if o == nil || IsNil(o.IsHyperConverged) {
		var ret bool
		return ret
	}
	return *o.IsHyperConverged
}

// GetIsHyperConvergedOk returns a tuple with the IsHyperConverged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetIsHyperConvergedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsHyperConverged) {
		return nil, false
	}
	return o.IsHyperConverged, true
}

// HasIsHyperConverged returns a boolean if a field has been set.
func (o *ApplianceInstance) HasIsHyperConverged() bool {
	if o != nil && !IsNil(o.IsHyperConverged) {
		return true
	}

	return false
}

// SetIsHyperConverged gets a reference to the given bool and assigns it to the IsHyperConverged field.
func (o *ApplianceInstance) SetIsHyperConverged(v bool) {
	o.IsHyperConverged = &v
}

// GetModeL10n returns the ModeL10n field value if set, zero value otherwise.
func (o *ApplianceInstance) GetModeL10n() string {
	if o == nil || IsNil(o.ModeL10n) {
		var ret string
		return ret
	}
	return *o.ModeL10n
}

// GetModeL10nOk returns a tuple with the ModeL10n field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetModeL10nOk() (*string, bool) {
	if o == nil || IsNil(o.ModeL10n) {
		return nil, false
	}
	return o.ModeL10n, true
}

// HasModeL10n returns a boolean if a field has been set.
func (o *ApplianceInstance) HasModeL10n() bool {
	if o != nil && !IsNil(o.ModeL10n) {
		return true
	}

	return false
}

// SetModeL10n gets a reference to the given string and assigns it to the ModeL10n field.
func (o *ApplianceInstance) SetModeL10n(v string) {
	o.ModeL10n = &v
}

// GetDriveFailureToleranceLevelL10n returns the DriveFailureToleranceLevelL10n field value if set, zero value otherwise.
func (o *ApplianceInstance) GetDriveFailureToleranceLevelL10n() string {
	if o == nil || IsNil(o.DriveFailureToleranceLevelL10n) {
		var ret string
		return ret
	}
	return *o.DriveFailureToleranceLevelL10n
}

// GetDriveFailureToleranceLevelL10nOk returns a tuple with the DriveFailureToleranceLevelL10n field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetDriveFailureToleranceLevelL10nOk() (*string, bool) {
	if o == nil || IsNil(o.DriveFailureToleranceLevelL10n) {
		return nil, false
	}
	return o.DriveFailureToleranceLevelL10n, true
}

// HasDriveFailureToleranceLevelL10n returns a boolean if a field has been set.
func (o *ApplianceInstance) HasDriveFailureToleranceLevelL10n() bool {
	if o != nil && !IsNil(o.DriveFailureToleranceLevelL10n) {
		return true
	}

	return false
}

// SetDriveFailureToleranceLevelL10n gets a reference to the given string and assigns it to the DriveFailureToleranceLevelL10n field.
func (o *ApplianceInstance) SetDriveFailureToleranceLevelL10n(v string) {
	o.DriveFailureToleranceLevelL10n = &v
}

// GetStorageClassL10n returns the StorageClassL10n field value if set, zero value otherwise.
func (o *ApplianceInstance) GetStorageClassL10n() string {
	if o == nil || IsNil(o.StorageClassL10n) {
		var ret string
		return ret
	}
	return *o.StorageClassL10n
}

// GetStorageClassL10nOk returns a tuple with the StorageClassL10n field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetStorageClassL10nOk() (*string, bool) {
	if o == nil || IsNil(o.StorageClassL10n) {
		return nil, false
	}
	return o.StorageClassL10n, true
}

// HasStorageClassL10n returns a boolean if a field has been set.
func (o *ApplianceInstance) HasStorageClassL10n() bool {
	if o != nil && !IsNil(o.StorageClassL10n) {
		return true
	}

	return false
}

// SetStorageClassL10n gets a reference to the given string and assigns it to the StorageClassL10n field.
func (o *ApplianceInstance) SetStorageClassL10n(v string) {
	o.StorageClassL10n = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *ApplianceInstance) GetNodes() []NodeInstance {
	if o == nil || IsNil(o.Nodes) {
		var ret []NodeInstance
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetNodesOk() ([]NodeInstance, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *ApplianceInstance) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given []NodeInstance and assigns it to the Nodes field.
func (o *ApplianceInstance) SetNodes(v []NodeInstance) {
	o.Nodes = v
}

// GetIpPoolAddresses returns the IpPoolAddresses field value if set, zero value otherwise.
func (o *ApplianceInstance) GetIpPoolAddresses() []IpPoolAddressInstance {
	if o == nil || IsNil(o.IpPoolAddresses) {
		var ret []IpPoolAddressInstance
		return ret
	}
	return o.IpPoolAddresses
}

// GetIpPoolAddressesOk returns a tuple with the IpPoolAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetIpPoolAddressesOk() ([]IpPoolAddressInstance, bool) {
	if o == nil || IsNil(o.IpPoolAddresses) {
		return nil, false
	}
	return o.IpPoolAddresses, true
}

// HasIpPoolAddresses returns a boolean if a field has been set.
func (o *ApplianceInstance) HasIpPoolAddresses() bool {
	if o != nil && !IsNil(o.IpPoolAddresses) {
		return true
	}

	return false
}

// SetIpPoolAddresses gets a reference to the given []IpPoolAddressInstance and assigns it to the IpPoolAddresses field.
func (o *ApplianceInstance) SetIpPoolAddresses(v []IpPoolAddressInstance) {
	o.IpPoolAddresses = v
}

// GetFsns returns the Fsns field value if set, zero value otherwise.
func (o *ApplianceInstance) GetFsns() []FsnInstance {
	if o == nil || IsNil(o.Fsns) {
		var ret []FsnInstance
		return ret
	}
	return o.Fsns
}

// GetFsnsOk returns a tuple with the Fsns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetFsnsOk() ([]FsnInstance, bool) {
	if o == nil || IsNil(o.Fsns) {
		return nil, false
	}
	return o.Fsns, true
}

// HasFsns returns a boolean if a field has been set.
func (o *ApplianceInstance) HasFsns() bool {
	if o != nil && !IsNil(o.Fsns) {
		return true
	}

	return false
}

// SetFsns gets a reference to the given []FsnInstance and assigns it to the Fsns field.
func (o *ApplianceInstance) SetFsns(v []FsnInstance) {
	o.Fsns = v
}

// GetVethPorts returns the VethPorts field value if set, zero value otherwise.
func (o *ApplianceInstance) GetVethPorts() []VethPortInstance {
	if o == nil || IsNil(o.VethPorts) {
		var ret []VethPortInstance
		return ret
	}
	return o.VethPorts
}

// GetVethPortsOk returns a tuple with the VethPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetVethPortsOk() ([]VethPortInstance, bool) {
	if o == nil || IsNil(o.VethPorts) {
		return nil, false
	}
	return o.VethPorts, true
}

// HasVethPorts returns a boolean if a field has been set.
func (o *ApplianceInstance) HasVethPorts() bool {
	if o != nil && !IsNil(o.VethPorts) {
		return true
	}

	return false
}

// SetVethPorts gets a reference to the given []VethPortInstance and assigns it to the VethPorts field.
func (o *ApplianceInstance) SetVethPorts(v []VethPortInstance) {
	o.VethPorts = v
}

// GetVirtualVolumes returns the VirtualVolumes field value if set, zero value otherwise.
func (o *ApplianceInstance) GetVirtualVolumes() []VirtualVolumeInstance {
	if o == nil || IsNil(o.VirtualVolumes) {
		var ret []VirtualVolumeInstance
		return ret
	}
	return o.VirtualVolumes
}

// GetVirtualVolumesOk returns a tuple with the VirtualVolumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetVirtualVolumesOk() ([]VirtualVolumeInstance, bool) {
	if o == nil || IsNil(o.VirtualVolumes) {
		return nil, false
	}
	return o.VirtualVolumes, true
}

// HasVirtualVolumes returns a boolean if a field has been set.
func (o *ApplianceInstance) HasVirtualVolumes() bool {
	if o != nil && !IsNil(o.VirtualVolumes) {
		return true
	}

	return false
}

// SetVirtualVolumes gets a reference to the given []VirtualVolumeInstance and assigns it to the VirtualVolumes field.
func (o *ApplianceInstance) SetVirtualVolumes(v []VirtualVolumeInstance) {
	o.VirtualVolumes = v
}

// GetMaintenanceWindows returns the MaintenanceWindows field value if set, zero value otherwise.
func (o *ApplianceInstance) GetMaintenanceWindows() []MaintenanceWindowInstance {
	if o == nil || IsNil(o.MaintenanceWindows) {
		var ret []MaintenanceWindowInstance
		return ret
	}
	return o.MaintenanceWindows
}

// GetMaintenanceWindowsOk returns a tuple with the MaintenanceWindows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetMaintenanceWindowsOk() ([]MaintenanceWindowInstance, bool) {
	if o == nil || IsNil(o.MaintenanceWindows) {
		return nil, false
	}
	return o.MaintenanceWindows, true
}

// HasMaintenanceWindows returns a boolean if a field has been set.
func (o *ApplianceInstance) HasMaintenanceWindows() bool {
	if o != nil && !IsNil(o.MaintenanceWindows) {
		return true
	}

	return false
}

// SetMaintenanceWindows gets a reference to the given []MaintenanceWindowInstance and assigns it to the MaintenanceWindows field.
func (o *ApplianceInstance) SetMaintenanceWindows(v []MaintenanceWindowInstance) {
	o.MaintenanceWindows = v
}

// GetFcPorts returns the FcPorts field value if set, zero value otherwise.
func (o *ApplianceInstance) GetFcPorts() []FcPortInstance {
	if o == nil || IsNil(o.FcPorts) {
		var ret []FcPortInstance
		return ret
	}
	return o.FcPorts
}

// GetFcPortsOk returns a tuple with the FcPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetFcPortsOk() ([]FcPortInstance, bool) {
	if o == nil || IsNil(o.FcPorts) {
		return nil, false
	}
	return o.FcPorts, true
}

// HasFcPorts returns a boolean if a field has been set.
func (o *ApplianceInstance) HasFcPorts() bool {
	if o != nil && !IsNil(o.FcPorts) {
		return true
	}

	return false
}

// SetFcPorts gets a reference to the given []FcPortInstance and assigns it to the FcPorts field.
func (o *ApplianceInstance) SetFcPorts(v []FcPortInstance) {
	o.FcPorts = v
}

// GetSasPorts returns the SasPorts field value if set, zero value otherwise.
func (o *ApplianceInstance) GetSasPorts() []SasPortInstance {
	if o == nil || IsNil(o.SasPorts) {
		var ret []SasPortInstance
		return ret
	}
	return o.SasPorts
}

// GetSasPortsOk returns a tuple with the SasPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetSasPortsOk() ([]SasPortInstance, bool) {
	if o == nil || IsNil(o.SasPorts) {
		return nil, false
	}
	return o.SasPorts, true
}

// HasSasPorts returns a boolean if a field has been set.
func (o *ApplianceInstance) HasSasPorts() bool {
	if o != nil && !IsNil(o.SasPorts) {
		return true
	}

	return false
}

// SetSasPorts gets a reference to the given []SasPortInstance and assigns it to the SasPorts field.
func (o *ApplianceInstance) SetSasPorts(v []SasPortInstance) {
	o.SasPorts = v
}

// GetEthPorts returns the EthPorts field value if set, zero value otherwise.
func (o *ApplianceInstance) GetEthPorts() []EthPortInstance {
	if o == nil || IsNil(o.EthPorts) {
		var ret []EthPortInstance
		return ret
	}
	return o.EthPorts
}

// GetEthPortsOk returns a tuple with the EthPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetEthPortsOk() ([]EthPortInstance, bool) {
	if o == nil || IsNil(o.EthPorts) {
		return nil, false
	}
	return o.EthPorts, true
}

// HasEthPorts returns a boolean if a field has been set.
func (o *ApplianceInstance) HasEthPorts() bool {
	if o != nil && !IsNil(o.EthPorts) {
		return true
	}

	return false
}

// SetEthPorts gets a reference to the given []EthPortInstance and assigns it to the EthPorts field.
func (o *ApplianceInstance) SetEthPorts(v []EthPortInstance) {
	o.EthPorts = v
}

// GetEthBePorts returns the EthBePorts field value if set, zero value otherwise.
func (o *ApplianceInstance) GetEthBePorts() []EthBePortInstance {
	if o == nil || IsNil(o.EthBePorts) {
		var ret []EthBePortInstance
		return ret
	}
	return o.EthBePorts
}

// GetEthBePortsOk returns a tuple with the EthBePorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetEthBePortsOk() ([]EthBePortInstance, bool) {
	if o == nil || IsNil(o.EthBePorts) {
		return nil, false
	}
	return o.EthBePorts, true
}

// HasEthBePorts returns a boolean if a field has been set.
func (o *ApplianceInstance) HasEthBePorts() bool {
	if o != nil && !IsNil(o.EthBePorts) {
		return true
	}

	return false
}

// SetEthBePorts gets a reference to the given []EthBePortInstance and assigns it to the EthBePorts field.
func (o *ApplianceInstance) SetEthBePorts(v []EthBePortInstance) {
	o.EthBePorts = v
}

// GetSoftwareInstalled returns the SoftwareInstalled field value if set, zero value otherwise.
func (o *ApplianceInstance) GetSoftwareInstalled() []SoftwareInstalledInstance {
	if o == nil || IsNil(o.SoftwareInstalled) {
		var ret []SoftwareInstalledInstance
		return ret
	}
	return o.SoftwareInstalled
}

// GetSoftwareInstalledOk returns a tuple with the SoftwareInstalled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetSoftwareInstalledOk() ([]SoftwareInstalledInstance, bool) {
	if o == nil || IsNil(o.SoftwareInstalled) {
		return nil, false
	}
	return o.SoftwareInstalled, true
}

// HasSoftwareInstalled returns a boolean if a field has been set.
func (o *ApplianceInstance) HasSoftwareInstalled() bool {
	if o != nil && !IsNil(o.SoftwareInstalled) {
		return true
	}

	return false
}

// SetSoftwareInstalled gets a reference to the given []SoftwareInstalledInstance and assigns it to the SoftwareInstalled field.
func (o *ApplianceInstance) SetSoftwareInstalled(v []SoftwareInstalledInstance) {
	o.SoftwareInstalled = v
}

// GetHardware returns the Hardware field value if set, zero value otherwise.
func (o *ApplianceInstance) GetHardware() []HardwareInstance {
	if o == nil || IsNil(o.Hardware) {
		var ret []HardwareInstance
		return ret
	}
	return o.Hardware
}

// GetHardwareOk returns a tuple with the Hardware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetHardwareOk() ([]HardwareInstance, bool) {
	if o == nil || IsNil(o.Hardware) {
		return nil, false
	}
	return o.Hardware, true
}

// HasHardware returns a boolean if a field has been set.
func (o *ApplianceInstance) HasHardware() bool {
	if o != nil && !IsNil(o.Hardware) {
		return true
	}

	return false
}

// SetHardware gets a reference to the given []HardwareInstance and assigns it to the Hardware field.
func (o *ApplianceInstance) SetHardware(v []HardwareInstance) {
	o.Hardware = v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *ApplianceInstance) GetVolumes() []VolumeInstance {
	if o == nil || IsNil(o.Volumes) {
		var ret []VolumeInstance
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceInstance) GetVolumesOk() ([]VolumeInstance, bool) {
	if o == nil || IsNil(o.Volumes) {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *ApplianceInstance) HasVolumes() bool {
	if o != nil && !IsNil(o.Volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []VolumeInstance and assigns it to the Volumes field.
func (o *ApplianceInstance) SetVolumes(v []VolumeInstance) {
	o.Volumes = v
}

func (o ApplianceInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplianceInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ServiceTag) {
		toSerialize["service_tag"] = o.ServiceTag
	}
	if !IsNil(o.ExpressServiceCode) {
		toSerialize["express_service_code"] = o.ExpressServiceCode
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.NodeCount) {
		toSerialize["node_count"] = o.NodeCount
	}
	if !IsNil(o.DriveFailureToleranceLevel) {
		toSerialize["drive_failure_tolerance_level"] = o.DriveFailureToleranceLevel
	}
	if !IsNil(o.StorageClass) {
		toSerialize["storage_class"] = o.StorageClass
	}
	if !IsNil(o.IsHyperConverged) {
		toSerialize["is_hyper_converged"] = o.IsHyperConverged
	}
	if !IsNil(o.ModeL10n) {
		toSerialize["mode_l10n"] = o.ModeL10n
	}
	if !IsNil(o.DriveFailureToleranceLevelL10n) {
		toSerialize["drive_failure_tolerance_level_l10n"] = o.DriveFailureToleranceLevelL10n
	}
	if !IsNil(o.StorageClassL10n) {
		toSerialize["storage_class_l10n"] = o.StorageClassL10n
	}
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	if !IsNil(o.IpPoolAddresses) {
		toSerialize["ip_pool_addresses"] = o.IpPoolAddresses
	}
	if !IsNil(o.Fsns) {
		toSerialize["fsns"] = o.Fsns
	}
	if !IsNil(o.VethPorts) {
		toSerialize["veth_ports"] = o.VethPorts
	}
	if !IsNil(o.VirtualVolumes) {
		toSerialize["virtual_volumes"] = o.VirtualVolumes
	}
	if !IsNil(o.MaintenanceWindows) {
		toSerialize["maintenance_windows"] = o.MaintenanceWindows
	}
	if !IsNil(o.FcPorts) {
		toSerialize["fc_ports"] = o.FcPorts
	}
	if !IsNil(o.SasPorts) {
		toSerialize["sas_ports"] = o.SasPorts
	}
	if !IsNil(o.EthPorts) {
		toSerialize["eth_ports"] = o.EthPorts
	}
	if !IsNil(o.EthBePorts) {
		toSerialize["eth_be_ports"] = o.EthBePorts
	}
	if !IsNil(o.SoftwareInstalled) {
		toSerialize["software_installed"] = o.SoftwareInstalled
	}
	if !IsNil(o.Hardware) {
		toSerialize["hardware"] = o.Hardware
	}
	if !IsNil(o.Volumes) {
		toSerialize["volumes"] = o.Volumes
	}
	return toSerialize, nil
}

type NullableApplianceInstance struct {
	value *ApplianceInstance
	isSet bool
}

func (v NullableApplianceInstance) Get() *ApplianceInstance {
	return v.value
}

func (v *NullableApplianceInstance) Set(val *ApplianceInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceInstance(val *ApplianceInstance) *NullableApplianceInstance {
	return &NullableApplianceInstance{value: val, isSet: true}
}

func (v NullableApplianceInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


