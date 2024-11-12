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

// checks if the AssessmentPeriodsPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssessmentPeriodsPut{}

// AssessmentPeriodsPut struct for AssessmentPeriodsPut
type AssessmentPeriodsPut struct {
	AssessmentPeriod *AssessmentPeriodsPutAssessmentPeriod `json:"assessment_period,omitempty"`
	PreviousValues *AssessmentPeriodsPutPreviousValues `json:"previous_values,omitempty"`
}

// NewAssessmentPeriodsPut instantiates a new AssessmentPeriodsPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssessmentPeriodsPut() *AssessmentPeriodsPut {
	this := AssessmentPeriodsPut{}
	return &this
}

// NewAssessmentPeriodsPutWithDefaults instantiates a new AssessmentPeriodsPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssessmentPeriodsPutWithDefaults() *AssessmentPeriodsPut {
	this := AssessmentPeriodsPut{}
	return &this
}

// GetAssessmentPeriod returns the AssessmentPeriod field value if set, zero value otherwise.
func (o *AssessmentPeriodsPut) GetAssessmentPeriod() AssessmentPeriodsPutAssessmentPeriod {
	if o == nil || IsNil(o.AssessmentPeriod) {
		var ret AssessmentPeriodsPutAssessmentPeriod
		return ret
	}
	return *o.AssessmentPeriod
}

// GetAssessmentPeriodOk returns a tuple with the AssessmentPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentPeriodsPut) GetAssessmentPeriodOk() (*AssessmentPeriodsPutAssessmentPeriod, bool) {
	if o == nil || IsNil(o.AssessmentPeriod) {
		return nil, false
	}
	return o.AssessmentPeriod, true
}

// HasAssessmentPeriod returns a boolean if a field has been set.
func (o *AssessmentPeriodsPut) HasAssessmentPeriod() bool {
	if o != nil && !IsNil(o.AssessmentPeriod) {
		return true
	}

	return false
}

// SetAssessmentPeriod gets a reference to the given AssessmentPeriodsPutAssessmentPeriod and assigns it to the AssessmentPeriod field.
func (o *AssessmentPeriodsPut) SetAssessmentPeriod(v AssessmentPeriodsPutAssessmentPeriod) {
	o.AssessmentPeriod = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *AssessmentPeriodsPut) GetPreviousValues() AssessmentPeriodsPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret AssessmentPeriodsPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentPeriodsPut) GetPreviousValuesOk() (*AssessmentPeriodsPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *AssessmentPeriodsPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given AssessmentPeriodsPutPreviousValues and assigns it to the PreviousValues field.
func (o *AssessmentPeriodsPut) SetPreviousValues(v AssessmentPeriodsPutPreviousValues) {
	o.PreviousValues = &v
}

func (o AssessmentPeriodsPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssessmentPeriodsPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssessmentPeriod) {
		toSerialize["assessment_period"] = o.AssessmentPeriod
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableAssessmentPeriodsPut struct {
	value *AssessmentPeriodsPut
	isSet bool
}

func (v NullableAssessmentPeriodsPut) Get() *AssessmentPeriodsPut {
	return v.value
}

func (v *NullableAssessmentPeriodsPut) Set(val *AssessmentPeriodsPut) {
	v.value = val
	v.isSet = true
}

func (v NullableAssessmentPeriodsPut) IsSet() bool {
	return v.isSet
}

func (v *NullableAssessmentPeriodsPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssessmentPeriodsPut(val *AssessmentPeriodsPut) *NullableAssessmentPeriodsPut {
	return &NullableAssessmentPeriodsPut{value: val, isSet: true}
}

func (v NullableAssessmentPeriodsPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssessmentPeriodsPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

