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

// checks if the NodeInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeInstance{}

// NodeInstance This resource type has queriable associations from appliance, ip_pool_address, veth_port
type NodeInstance struct {
	// Unique identifier of the node.
	Id *string `json:"id,omitempty"`
	// Slot number of the node.
	Slot *int32 `json:"slot,omitempty"`
	// Unique identifier of the appliance to which the node belongs.
	ApplianceId *string `json:"appliance_id,omitempty"`
	Appliance *ApplianceInstance `json:"appliance,omitempty"`
	// This is the inverse of the resource type ip_pool_address association.
	IpPoolAddresses []IpPoolAddressInstance `json:"ip_pool_addresses,omitempty"`
	// This is the inverse of the resource type veth_port association.
	VethPorts []VethPortInstance `json:"veth_ports,omitempty"`
}

// NewNodeInstance instantiates a new NodeInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeInstance() *NodeInstance {
	this := NodeInstance{}
	return &this
}

// NewNodeInstanceWithDefaults instantiates a new NodeInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeInstanceWithDefaults() *NodeInstance {
	this := NodeInstance{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NodeInstance) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInstance) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NodeInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NodeInstance) SetId(v string) {
	o.Id = &v
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *NodeInstance) GetSlot() int32 {
	if o == nil || IsNil(o.Slot) {
		var ret int32
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInstance) GetSlotOk() (*int32, bool) {
	if o == nil || IsNil(o.Slot) {
		return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *NodeInstance) HasSlot() bool {
	if o != nil && !IsNil(o.Slot) {
		return true
	}

	return false
}

// SetSlot gets a reference to the given int32 and assigns it to the Slot field.
func (o *NodeInstance) SetSlot(v int32) {
	o.Slot = &v
}

// GetApplianceId returns the ApplianceId field value if set, zero value otherwise.
func (o *NodeInstance) GetApplianceId() string {
	if o == nil || IsNil(o.ApplianceId) {
		var ret string
		return ret
	}
	return *o.ApplianceId
}

// GetApplianceIdOk returns a tuple with the ApplianceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInstance) GetApplianceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplianceId) {
		return nil, false
	}
	return o.ApplianceId, true
}

// HasApplianceId returns a boolean if a field has been set.
func (o *NodeInstance) HasApplianceId() bool {
	if o != nil && !IsNil(o.ApplianceId) {
		return true
	}

	return false
}

// SetApplianceId gets a reference to the given string and assigns it to the ApplianceId field.
func (o *NodeInstance) SetApplianceId(v string) {
	o.ApplianceId = &v
}

// GetAppliance returns the Appliance field value if set, zero value otherwise.
func (o *NodeInstance) GetAppliance() ApplianceInstance {
	if o == nil || IsNil(o.Appliance) {
		var ret ApplianceInstance
		return ret
	}
	return *o.Appliance
}

// GetApplianceOk returns a tuple with the Appliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInstance) GetApplianceOk() (*ApplianceInstance, bool) {
	if o == nil || IsNil(o.Appliance) {
		return nil, false
	}
	return o.Appliance, true
}

// HasAppliance returns a boolean if a field has been set.
func (o *NodeInstance) HasAppliance() bool {
	if o != nil && !IsNil(o.Appliance) {
		return true
	}

	return false
}

// SetAppliance gets a reference to the given ApplianceInstance and assigns it to the Appliance field.
func (o *NodeInstance) SetAppliance(v ApplianceInstance) {
	o.Appliance = &v
}

// GetIpPoolAddresses returns the IpPoolAddresses field value if set, zero value otherwise.
func (o *NodeInstance) GetIpPoolAddresses() []IpPoolAddressInstance {
	if o == nil || IsNil(o.IpPoolAddresses) {
		var ret []IpPoolAddressInstance
		return ret
	}
	return o.IpPoolAddresses
}

// GetIpPoolAddressesOk returns a tuple with the IpPoolAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInstance) GetIpPoolAddressesOk() ([]IpPoolAddressInstance, bool) {
	if o == nil || IsNil(o.IpPoolAddresses) {
		return nil, false
	}
	return o.IpPoolAddresses, true
}

// HasIpPoolAddresses returns a boolean if a field has been set.
func (o *NodeInstance) HasIpPoolAddresses() bool {
	if o != nil && !IsNil(o.IpPoolAddresses) {
		return true
	}

	return false
}

// SetIpPoolAddresses gets a reference to the given []IpPoolAddressInstance and assigns it to the IpPoolAddresses field.
func (o *NodeInstance) SetIpPoolAddresses(v []IpPoolAddressInstance) {
	o.IpPoolAddresses = v
}

// GetVethPorts returns the VethPorts field value if set, zero value otherwise.
func (o *NodeInstance) GetVethPorts() []VethPortInstance {
	if o == nil || IsNil(o.VethPorts) {
		var ret []VethPortInstance
		return ret
	}
	return o.VethPorts
}

// GetVethPortsOk returns a tuple with the VethPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInstance) GetVethPortsOk() ([]VethPortInstance, bool) {
	if o == nil || IsNil(o.VethPorts) {
		return nil, false
	}
	return o.VethPorts, true
}

// HasVethPorts returns a boolean if a field has been set.
func (o *NodeInstance) HasVethPorts() bool {
	if o != nil && !IsNil(o.VethPorts) {
		return true
	}

	return false
}

// SetVethPorts gets a reference to the given []VethPortInstance and assigns it to the VethPorts field.
func (o *NodeInstance) SetVethPorts(v []VethPortInstance) {
	o.VethPorts = v
}

func (o NodeInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Slot) {
		toSerialize["slot"] = o.Slot
	}
	if !IsNil(o.ApplianceId) {
		toSerialize["appliance_id"] = o.ApplianceId
	}
	if !IsNil(o.Appliance) {
		toSerialize["appliance"] = o.Appliance
	}
	if !IsNil(o.IpPoolAddresses) {
		toSerialize["ip_pool_addresses"] = o.IpPoolAddresses
	}
	if !IsNil(o.VethPorts) {
		toSerialize["veth_ports"] = o.VethPorts
	}
	return toSerialize, nil
}

type NullableNodeInstance struct {
	value *NodeInstance
	isSet bool
}

func (v NullableNodeInstance) Get() *NodeInstance {
	return v.value
}

func (v *NullableNodeInstance) Set(val *NodeInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeInstance(val *NodeInstance) *NullableNodeInstance {
	return &NullableNodeInstance{value: val, isSet: true}
}

func (v NullableNodeInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


