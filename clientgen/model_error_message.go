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

// checks if the ErrorMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorMessage{}

// ErrorMessage Error message for an operation.
type ErrorMessage struct {
	// Hexadecimal error code of the message.
	Code *string `json:"code,omitempty"`
	Severity *MessageSeverityEnum `json:"severity,omitempty"`
	// The message description in the specified locale with arguments substituted. 
	MessageL10n *string `json:"message_l10n,omitempty"`
	// Arguments (if applicable) for the error message.
	Arguments []string `json:"arguments,omitempty"`
	// Localized message string corresponding to severity
	SeverityL10n *string `json:"severity_l10n,omitempty"`
}

// NewErrorMessage instantiates a new ErrorMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorMessage() *ErrorMessage {
	this := ErrorMessage{}
	return &this
}

// NewErrorMessageWithDefaults instantiates a new ErrorMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorMessageWithDefaults() *ErrorMessage {
	this := ErrorMessage{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ErrorMessage) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorMessage) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ErrorMessage) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ErrorMessage) SetCode(v string) {
	o.Code = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ErrorMessage) GetSeverity() MessageSeverityEnum {
	if o == nil || IsNil(o.Severity) {
		var ret MessageSeverityEnum
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorMessage) GetSeverityOk() (*MessageSeverityEnum, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ErrorMessage) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given MessageSeverityEnum and assigns it to the Severity field.
func (o *ErrorMessage) SetSeverity(v MessageSeverityEnum) {
	o.Severity = &v
}

// GetMessageL10n returns the MessageL10n field value if set, zero value otherwise.
func (o *ErrorMessage) GetMessageL10n() string {
	if o == nil || IsNil(o.MessageL10n) {
		var ret string
		return ret
	}
	return *o.MessageL10n
}

// GetMessageL10nOk returns a tuple with the MessageL10n field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorMessage) GetMessageL10nOk() (*string, bool) {
	if o == nil || IsNil(o.MessageL10n) {
		return nil, false
	}
	return o.MessageL10n, true
}

// HasMessageL10n returns a boolean if a field has been set.
func (o *ErrorMessage) HasMessageL10n() bool {
	if o != nil && !IsNil(o.MessageL10n) {
		return true
	}

	return false
}

// SetMessageL10n gets a reference to the given string and assigns it to the MessageL10n field.
func (o *ErrorMessage) SetMessageL10n(v string) {
	o.MessageL10n = &v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *ErrorMessage) GetArguments() []string {
	if o == nil || IsNil(o.Arguments) {
		var ret []string
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorMessage) GetArgumentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Arguments) {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *ErrorMessage) HasArguments() bool {
	if o != nil && !IsNil(o.Arguments) {
		return true
	}

	return false
}

// SetArguments gets a reference to the given []string and assigns it to the Arguments field.
func (o *ErrorMessage) SetArguments(v []string) {
	o.Arguments = v
}

// GetSeverityL10n returns the SeverityL10n field value if set, zero value otherwise.
func (o *ErrorMessage) GetSeverityL10n() string {
	if o == nil || IsNil(o.SeverityL10n) {
		var ret string
		return ret
	}
	return *o.SeverityL10n
}

// GetSeverityL10nOk returns a tuple with the SeverityL10n field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorMessage) GetSeverityL10nOk() (*string, bool) {
	if o == nil || IsNil(o.SeverityL10n) {
		return nil, false
	}
	return o.SeverityL10n, true
}

// HasSeverityL10n returns a boolean if a field has been set.
func (o *ErrorMessage) HasSeverityL10n() bool {
	if o != nil && !IsNil(o.SeverityL10n) {
		return true
	}

	return false
}

// SetSeverityL10n gets a reference to the given string and assigns it to the SeverityL10n field.
func (o *ErrorMessage) SetSeverityL10n(v string) {
	o.SeverityL10n = &v
}

func (o ErrorMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.MessageL10n) {
		toSerialize["message_l10n"] = o.MessageL10n
	}
	if !IsNil(o.Arguments) {
		toSerialize["arguments"] = o.Arguments
	}
	if !IsNil(o.SeverityL10n) {
		toSerialize["severity_l10n"] = o.SeverityL10n
	}
	return toSerialize, nil
}

type NullableErrorMessage struct {
	value *ErrorMessage
	isSet bool
}

func (v NullableErrorMessage) Get() *ErrorMessage {
	return v.value
}

func (v *NullableErrorMessage) Set(val *ErrorMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorMessage(val *ErrorMessage) *NullableErrorMessage {
	return &NullableErrorMessage{value: val, isSet: true}
}

func (v NullableErrorMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


