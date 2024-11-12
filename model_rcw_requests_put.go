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

// checks if the RcwRequestsPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RcwRequestsPut{}

// RcwRequestsPut struct for RcwRequestsPut
type RcwRequestsPut struct {
	RcwRequest *RcwRequestsPutRcwRequest `json:"rcw_request,omitempty"`
	PreviousValues *RcwRequestsPutPreviousValues `json:"previous_values,omitempty"`
}

// NewRcwRequestsPut instantiates a new RcwRequestsPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRcwRequestsPut() *RcwRequestsPut {
	this := RcwRequestsPut{}
	return &this
}

// NewRcwRequestsPutWithDefaults instantiates a new RcwRequestsPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRcwRequestsPutWithDefaults() *RcwRequestsPut {
	this := RcwRequestsPut{}
	return &this
}

// GetRcwRequest returns the RcwRequest field value if set, zero value otherwise.
func (o *RcwRequestsPut) GetRcwRequest() RcwRequestsPutRcwRequest {
	if o == nil || IsNil(o.RcwRequest) {
		var ret RcwRequestsPutRcwRequest
		return ret
	}
	return *o.RcwRequest
}

// GetRcwRequestOk returns a tuple with the RcwRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwRequestsPut) GetRcwRequestOk() (*RcwRequestsPutRcwRequest, bool) {
	if o == nil || IsNil(o.RcwRequest) {
		return nil, false
	}
	return o.RcwRequest, true
}

// HasRcwRequest returns a boolean if a field has been set.
func (o *RcwRequestsPut) HasRcwRequest() bool {
	if o != nil && !IsNil(o.RcwRequest) {
		return true
	}

	return false
}

// SetRcwRequest gets a reference to the given RcwRequestsPutRcwRequest and assigns it to the RcwRequest field.
func (o *RcwRequestsPut) SetRcwRequest(v RcwRequestsPutRcwRequest) {
	o.RcwRequest = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *RcwRequestsPut) GetPreviousValues() RcwRequestsPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret RcwRequestsPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwRequestsPut) GetPreviousValuesOk() (*RcwRequestsPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *RcwRequestsPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given RcwRequestsPutPreviousValues and assigns it to the PreviousValues field.
func (o *RcwRequestsPut) SetPreviousValues(v RcwRequestsPutPreviousValues) {
	o.PreviousValues = &v
}

func (o RcwRequestsPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RcwRequestsPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RcwRequest) {
		toSerialize["rcw_request"] = o.RcwRequest
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableRcwRequestsPut struct {
	value *RcwRequestsPut
	isSet bool
}

func (v NullableRcwRequestsPut) Get() *RcwRequestsPut {
	return v.value
}

func (v *NullableRcwRequestsPut) Set(val *RcwRequestsPut) {
	v.value = val
	v.isSet = true
}

func (v NullableRcwRequestsPut) IsSet() bool {
	return v.isSet
}

func (v *NullableRcwRequestsPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRcwRequestsPut(val *RcwRequestsPut) *NullableRcwRequestsPut {
	return &NullableRcwRequestsPut{value: val, isSet: true}
}

func (v NullableRcwRequestsPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRcwRequestsPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

