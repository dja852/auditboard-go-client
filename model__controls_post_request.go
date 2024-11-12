/*
AuditBoard Developer Portal API Documentation

External API reference documentation.

API version: 23.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auditboard

import (
	"encoding/json"
)

// checks if the ControlsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControlsPostRequest{}

// ControlsPostRequest struct for ControlsPostRequest
type ControlsPostRequest struct {
	Control *Controls `json:"control,omitempty"`
}

// NewControlsPostRequest instantiates a new ControlsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControlsPostRequest() *ControlsPostRequest {
	this := ControlsPostRequest{}
	return &this
}

// NewControlsPostRequestWithDefaults instantiates a new ControlsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControlsPostRequestWithDefaults() *ControlsPostRequest {
	this := ControlsPostRequest{}
	return &this
}

// GetControl returns the Control field value if set, zero value otherwise.
func (o *ControlsPostRequest) GetControl() Controls {
	if o == nil || IsNil(o.Control) {
		var ret Controls
		return ret
	}
	return *o.Control
}

// GetControlOk returns a tuple with the Control field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsPostRequest) GetControlOk() (*Controls, bool) {
	if o == nil || IsNil(o.Control) {
		return nil, false
	}
	return o.Control, true
}

// HasControl returns a boolean if a field has been set.
func (o *ControlsPostRequest) HasControl() bool {
	if o != nil && !IsNil(o.Control) {
		return true
	}

	return false
}

// SetControl gets a reference to the given Controls and assigns it to the Control field.
func (o *ControlsPostRequest) SetControl(v Controls) {
	o.Control = &v
}

func (o ControlsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControlsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Control) {
		toSerialize["control"] = o.Control
	}
	return toSerialize, nil
}

type NullableControlsPostRequest struct {
	value *ControlsPostRequest
	isSet bool
}

func (v NullableControlsPostRequest) Get() *ControlsPostRequest {
	return v.value
}

func (v *NullableControlsPostRequest) Set(val *ControlsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableControlsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableControlsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControlsPostRequest(val *ControlsPostRequest) *NullableControlsPostRequest {
	return &NullableControlsPostRequest{value: val, isSet: true}
}

func (v NullableControlsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControlsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


