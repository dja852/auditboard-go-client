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

// checks if the RcwSchedulesGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RcwSchedulesGet200Response{}

// RcwSchedulesGet200Response struct for RcwSchedulesGet200Response
type RcwSchedulesGet200Response struct {
	RcwSchedules []RcwSchedules `json:"rcw_schedules,omitempty"`
}

// NewRcwSchedulesGet200Response instantiates a new RcwSchedulesGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRcwSchedulesGet200Response() *RcwSchedulesGet200Response {
	this := RcwSchedulesGet200Response{}
	return &this
}

// NewRcwSchedulesGet200ResponseWithDefaults instantiates a new RcwSchedulesGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRcwSchedulesGet200ResponseWithDefaults() *RcwSchedulesGet200Response {
	this := RcwSchedulesGet200Response{}
	return &this
}

// GetRcwSchedules returns the RcwSchedules field value if set, zero value otherwise.
func (o *RcwSchedulesGet200Response) GetRcwSchedules() []RcwSchedules {
	if o == nil || IsNil(o.RcwSchedules) {
		var ret []RcwSchedules
		return ret
	}
	return o.RcwSchedules
}

// GetRcwSchedulesOk returns a tuple with the RcwSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwSchedulesGet200Response) GetRcwSchedulesOk() ([]RcwSchedules, bool) {
	if o == nil || IsNil(o.RcwSchedules) {
		return nil, false
	}
	return o.RcwSchedules, true
}

// HasRcwSchedules returns a boolean if a field has been set.
func (o *RcwSchedulesGet200Response) HasRcwSchedules() bool {
	if o != nil && !IsNil(o.RcwSchedules) {
		return true
	}

	return false
}

// SetRcwSchedules gets a reference to the given []RcwSchedules and assigns it to the RcwSchedules field.
func (o *RcwSchedulesGet200Response) SetRcwSchedules(v []RcwSchedules) {
	o.RcwSchedules = v
}

func (o RcwSchedulesGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RcwSchedulesGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RcwSchedules) {
		toSerialize["rcw_schedules"] = o.RcwSchedules
	}
	return toSerialize, nil
}

type NullableRcwSchedulesGet200Response struct {
	value *RcwSchedulesGet200Response
	isSet bool
}

func (v NullableRcwSchedulesGet200Response) Get() *RcwSchedulesGet200Response {
	return v.value
}

func (v *NullableRcwSchedulesGet200Response) Set(val *RcwSchedulesGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRcwSchedulesGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRcwSchedulesGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRcwSchedulesGet200Response(val *RcwSchedulesGet200Response) *NullableRcwSchedulesGet200Response {
	return &NullableRcwSchedulesGet200Response{value: val, isSet: true}
}

func (v NullableRcwSchedulesGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRcwSchedulesGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


