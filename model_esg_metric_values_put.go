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

// checks if the EsgMetricValuesPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsgMetricValuesPut{}

// EsgMetricValuesPut struct for EsgMetricValuesPut
type EsgMetricValuesPut struct {
	EsgMetricValue *EsgMetricValuesPutEsgMetricValue `json:"esg_metric_value,omitempty"`
	PreviousValues *EsgMetricValuesPutPreviousValues `json:"previous_values,omitempty"`
}

// NewEsgMetricValuesPut instantiates a new EsgMetricValuesPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsgMetricValuesPut() *EsgMetricValuesPut {
	this := EsgMetricValuesPut{}
	return &this
}

// NewEsgMetricValuesPutWithDefaults instantiates a new EsgMetricValuesPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsgMetricValuesPutWithDefaults() *EsgMetricValuesPut {
	this := EsgMetricValuesPut{}
	return &this
}

// GetEsgMetricValue returns the EsgMetricValue field value if set, zero value otherwise.
func (o *EsgMetricValuesPut) GetEsgMetricValue() EsgMetricValuesPutEsgMetricValue {
	if o == nil || IsNil(o.EsgMetricValue) {
		var ret EsgMetricValuesPutEsgMetricValue
		return ret
	}
	return *o.EsgMetricValue
}

// GetEsgMetricValueOk returns a tuple with the EsgMetricValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricValuesPut) GetEsgMetricValueOk() (*EsgMetricValuesPutEsgMetricValue, bool) {
	if o == nil || IsNil(o.EsgMetricValue) {
		return nil, false
	}
	return o.EsgMetricValue, true
}

// HasEsgMetricValue returns a boolean if a field has been set.
func (o *EsgMetricValuesPut) HasEsgMetricValue() bool {
	if o != nil && !IsNil(o.EsgMetricValue) {
		return true
	}

	return false
}

// SetEsgMetricValue gets a reference to the given EsgMetricValuesPutEsgMetricValue and assigns it to the EsgMetricValue field.
func (o *EsgMetricValuesPut) SetEsgMetricValue(v EsgMetricValuesPutEsgMetricValue) {
	o.EsgMetricValue = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *EsgMetricValuesPut) GetPreviousValues() EsgMetricValuesPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret EsgMetricValuesPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricValuesPut) GetPreviousValuesOk() (*EsgMetricValuesPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *EsgMetricValuesPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given EsgMetricValuesPutPreviousValues and assigns it to the PreviousValues field.
func (o *EsgMetricValuesPut) SetPreviousValues(v EsgMetricValuesPutPreviousValues) {
	o.PreviousValues = &v
}

func (o EsgMetricValuesPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsgMetricValuesPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EsgMetricValue) {
		toSerialize["esg_metric_value"] = o.EsgMetricValue
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableEsgMetricValuesPut struct {
	value *EsgMetricValuesPut
	isSet bool
}

func (v NullableEsgMetricValuesPut) Get() *EsgMetricValuesPut {
	return v.value
}

func (v *NullableEsgMetricValuesPut) Set(val *EsgMetricValuesPut) {
	v.value = val
	v.isSet = true
}

func (v NullableEsgMetricValuesPut) IsSet() bool {
	return v.isSet
}

func (v *NullableEsgMetricValuesPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsgMetricValuesPut(val *EsgMetricValuesPut) *NullableEsgMetricValuesPut {
	return &NullableEsgMetricValuesPut{value: val, isSet: true}
}

func (v NullableEsgMetricValuesPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsgMetricValuesPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


