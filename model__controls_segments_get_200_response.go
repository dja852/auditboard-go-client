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

// checks if the ControlsSegmentsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControlsSegmentsGet200Response{}

// ControlsSegmentsGet200Response struct for ControlsSegmentsGet200Response
type ControlsSegmentsGet200Response struct {
	ControlsSegments []ControlsSegments `json:"controls_segments,omitempty"`
}

// NewControlsSegmentsGet200Response instantiates a new ControlsSegmentsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControlsSegmentsGet200Response() *ControlsSegmentsGet200Response {
	this := ControlsSegmentsGet200Response{}
	return &this
}

// NewControlsSegmentsGet200ResponseWithDefaults instantiates a new ControlsSegmentsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControlsSegmentsGet200ResponseWithDefaults() *ControlsSegmentsGet200Response {
	this := ControlsSegmentsGet200Response{}
	return &this
}

// GetControlsSegments returns the ControlsSegments field value if set, zero value otherwise.
func (o *ControlsSegmentsGet200Response) GetControlsSegments() []ControlsSegments {
	if o == nil || IsNil(o.ControlsSegments) {
		var ret []ControlsSegments
		return ret
	}
	return o.ControlsSegments
}

// GetControlsSegmentsOk returns a tuple with the ControlsSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsSegmentsGet200Response) GetControlsSegmentsOk() ([]ControlsSegments, bool) {
	if o == nil || IsNil(o.ControlsSegments) {
		return nil, false
	}
	return o.ControlsSegments, true
}

// HasControlsSegments returns a boolean if a field has been set.
func (o *ControlsSegmentsGet200Response) HasControlsSegments() bool {
	if o != nil && !IsNil(o.ControlsSegments) {
		return true
	}

	return false
}

// SetControlsSegments gets a reference to the given []ControlsSegments and assigns it to the ControlsSegments field.
func (o *ControlsSegmentsGet200Response) SetControlsSegments(v []ControlsSegments) {
	o.ControlsSegments = v
}

func (o ControlsSegmentsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControlsSegmentsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ControlsSegments) {
		toSerialize["controls_segments"] = o.ControlsSegments
	}
	return toSerialize, nil
}

type NullableControlsSegmentsGet200Response struct {
	value *ControlsSegmentsGet200Response
	isSet bool
}

func (v NullableControlsSegmentsGet200Response) Get() *ControlsSegmentsGet200Response {
	return v.value
}

func (v *NullableControlsSegmentsGet200Response) Set(val *ControlsSegmentsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableControlsSegmentsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableControlsSegmentsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControlsSegmentsGet200Response(val *ControlsSegmentsGet200Response) *NullableControlsSegmentsGet200Response {
	return &NullableControlsSegmentsGet200Response{value: val, isSet: true}
}

func (v NullableControlsSegmentsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControlsSegmentsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

