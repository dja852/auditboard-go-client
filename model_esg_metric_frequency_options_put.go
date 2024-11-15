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

// checks if the EsgMetricFrequencyOptionsPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsgMetricFrequencyOptionsPut{}

// EsgMetricFrequencyOptionsPut struct for EsgMetricFrequencyOptionsPut
type EsgMetricFrequencyOptionsPut struct {
	EsgMetricFrequencyOption *AssessmentPeriodsPutAssessmentPeriod `json:"esg_metric_frequency_option,omitempty"`
	PreviousValues *AssessmentPeriodsPutPreviousValues `json:"previous_values,omitempty"`
}

// NewEsgMetricFrequencyOptionsPut instantiates a new EsgMetricFrequencyOptionsPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsgMetricFrequencyOptionsPut() *EsgMetricFrequencyOptionsPut {
	this := EsgMetricFrequencyOptionsPut{}
	return &this
}

// NewEsgMetricFrequencyOptionsPutWithDefaults instantiates a new EsgMetricFrequencyOptionsPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsgMetricFrequencyOptionsPutWithDefaults() *EsgMetricFrequencyOptionsPut {
	this := EsgMetricFrequencyOptionsPut{}
	return &this
}

// GetEsgMetricFrequencyOption returns the EsgMetricFrequencyOption field value if set, zero value otherwise.
func (o *EsgMetricFrequencyOptionsPut) GetEsgMetricFrequencyOption() AssessmentPeriodsPutAssessmentPeriod {
	if o == nil || IsNil(o.EsgMetricFrequencyOption) {
		var ret AssessmentPeriodsPutAssessmentPeriod
		return ret
	}
	return *o.EsgMetricFrequencyOption
}

// GetEsgMetricFrequencyOptionOk returns a tuple with the EsgMetricFrequencyOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricFrequencyOptionsPut) GetEsgMetricFrequencyOptionOk() (*AssessmentPeriodsPutAssessmentPeriod, bool) {
	if o == nil || IsNil(o.EsgMetricFrequencyOption) {
		return nil, false
	}
	return o.EsgMetricFrequencyOption, true
}

// HasEsgMetricFrequencyOption returns a boolean if a field has been set.
func (o *EsgMetricFrequencyOptionsPut) HasEsgMetricFrequencyOption() bool {
	if o != nil && !IsNil(o.EsgMetricFrequencyOption) {
		return true
	}

	return false
}

// SetEsgMetricFrequencyOption gets a reference to the given AssessmentPeriodsPutAssessmentPeriod and assigns it to the EsgMetricFrequencyOption field.
func (o *EsgMetricFrequencyOptionsPut) SetEsgMetricFrequencyOption(v AssessmentPeriodsPutAssessmentPeriod) {
	o.EsgMetricFrequencyOption = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *EsgMetricFrequencyOptionsPut) GetPreviousValues() AssessmentPeriodsPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret AssessmentPeriodsPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricFrequencyOptionsPut) GetPreviousValuesOk() (*AssessmentPeriodsPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *EsgMetricFrequencyOptionsPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given AssessmentPeriodsPutPreviousValues and assigns it to the PreviousValues field.
func (o *EsgMetricFrequencyOptionsPut) SetPreviousValues(v AssessmentPeriodsPutPreviousValues) {
	o.PreviousValues = &v
}

func (o EsgMetricFrequencyOptionsPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsgMetricFrequencyOptionsPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EsgMetricFrequencyOption) {
		toSerialize["esg_metric_frequency_option"] = o.EsgMetricFrequencyOption
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableEsgMetricFrequencyOptionsPut struct {
	value *EsgMetricFrequencyOptionsPut
	isSet bool
}

func (v NullableEsgMetricFrequencyOptionsPut) Get() *EsgMetricFrequencyOptionsPut {
	return v.value
}

func (v *NullableEsgMetricFrequencyOptionsPut) Set(val *EsgMetricFrequencyOptionsPut) {
	v.value = val
	v.isSet = true
}

func (v NullableEsgMetricFrequencyOptionsPut) IsSet() bool {
	return v.isSet
}

func (v *NullableEsgMetricFrequencyOptionsPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsgMetricFrequencyOptionsPut(val *EsgMetricFrequencyOptionsPut) *NullableEsgMetricFrequencyOptionsPut {
	return &NullableEsgMetricFrequencyOptionsPut{value: val, isSet: true}
}

func (v NullableEsgMetricFrequencyOptionsPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsgMetricFrequencyOptionsPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


