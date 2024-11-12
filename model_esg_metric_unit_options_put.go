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

// checks if the EsgMetricUnitOptionsPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsgMetricUnitOptionsPut{}

// EsgMetricUnitOptionsPut struct for EsgMetricUnitOptionsPut
type EsgMetricUnitOptionsPut struct {
	EsgMetricUnitOption *AssessmentPeriodsPutAssessmentPeriod `json:"esg_metric_unit_option,omitempty"`
	PreviousValues *AssessmentPeriodsPutPreviousValues `json:"previous_values,omitempty"`
}

// NewEsgMetricUnitOptionsPut instantiates a new EsgMetricUnitOptionsPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsgMetricUnitOptionsPut() *EsgMetricUnitOptionsPut {
	this := EsgMetricUnitOptionsPut{}
	return &this
}

// NewEsgMetricUnitOptionsPutWithDefaults instantiates a new EsgMetricUnitOptionsPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsgMetricUnitOptionsPutWithDefaults() *EsgMetricUnitOptionsPut {
	this := EsgMetricUnitOptionsPut{}
	return &this
}

// GetEsgMetricUnitOption returns the EsgMetricUnitOption field value if set, zero value otherwise.
func (o *EsgMetricUnitOptionsPut) GetEsgMetricUnitOption() AssessmentPeriodsPutAssessmentPeriod {
	if o == nil || IsNil(o.EsgMetricUnitOption) {
		var ret AssessmentPeriodsPutAssessmentPeriod
		return ret
	}
	return *o.EsgMetricUnitOption
}

// GetEsgMetricUnitOptionOk returns a tuple with the EsgMetricUnitOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricUnitOptionsPut) GetEsgMetricUnitOptionOk() (*AssessmentPeriodsPutAssessmentPeriod, bool) {
	if o == nil || IsNil(o.EsgMetricUnitOption) {
		return nil, false
	}
	return o.EsgMetricUnitOption, true
}

// HasEsgMetricUnitOption returns a boolean if a field has been set.
func (o *EsgMetricUnitOptionsPut) HasEsgMetricUnitOption() bool {
	if o != nil && !IsNil(o.EsgMetricUnitOption) {
		return true
	}

	return false
}

// SetEsgMetricUnitOption gets a reference to the given AssessmentPeriodsPutAssessmentPeriod and assigns it to the EsgMetricUnitOption field.
func (o *EsgMetricUnitOptionsPut) SetEsgMetricUnitOption(v AssessmentPeriodsPutAssessmentPeriod) {
	o.EsgMetricUnitOption = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *EsgMetricUnitOptionsPut) GetPreviousValues() AssessmentPeriodsPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret AssessmentPeriodsPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricUnitOptionsPut) GetPreviousValuesOk() (*AssessmentPeriodsPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *EsgMetricUnitOptionsPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given AssessmentPeriodsPutPreviousValues and assigns it to the PreviousValues field.
func (o *EsgMetricUnitOptionsPut) SetPreviousValues(v AssessmentPeriodsPutPreviousValues) {
	o.PreviousValues = &v
}

func (o EsgMetricUnitOptionsPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsgMetricUnitOptionsPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EsgMetricUnitOption) {
		toSerialize["esg_metric_unit_option"] = o.EsgMetricUnitOption
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableEsgMetricUnitOptionsPut struct {
	value *EsgMetricUnitOptionsPut
	isSet bool
}

func (v NullableEsgMetricUnitOptionsPut) Get() *EsgMetricUnitOptionsPut {
	return v.value
}

func (v *NullableEsgMetricUnitOptionsPut) Set(val *EsgMetricUnitOptionsPut) {
	v.value = val
	v.isSet = true
}

func (v NullableEsgMetricUnitOptionsPut) IsSet() bool {
	return v.isSet
}

func (v *NullableEsgMetricUnitOptionsPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsgMetricUnitOptionsPut(val *EsgMetricUnitOptionsPut) *NullableEsgMetricUnitOptionsPut {
	return &NullableEsgMetricUnitOptionsPut{value: val, isSet: true}
}

func (v NullableEsgMetricUnitOptionsPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsgMetricUnitOptionsPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


