/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// VirtualMachinePowerStateEnum The current power state of the VM in vCenter. Not applicable to VM snapshots. * Powered_Off - VM is currently powered off. * Powered_On - VM is currently powered on. * Suspended - VM is currently suspended.  Was added in version 3.0.0.0.
type VirtualMachinePowerStateEnum string

// List of VirtualMachinePowerStateEnum
const (
	VIRTUALMACHINEPOWERSTATEENUM_POWERED_OFF VirtualMachinePowerStateEnum = "Powered_Off"
	VIRTUALMACHINEPOWERSTATEENUM_POWERED_ON  VirtualMachinePowerStateEnum = "Powered_On"
	VIRTUALMACHINEPOWERSTATEENUM_SUSPENDED   VirtualMachinePowerStateEnum = "Suspended"
)

// All allowed values of VirtualMachinePowerStateEnum enum
var AllowedVirtualMachinePowerStateEnumEnumValues = []VirtualMachinePowerStateEnum{
	"Powered_Off",
	"Powered_On",
	"Suspended",
}

func (v *VirtualMachinePowerStateEnum) Value() string {
	return string(*v)
}
