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

// checks if the TaskPeriodsPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskPeriodsPut{}

// TaskPeriodsPut struct for TaskPeriodsPut
type TaskPeriodsPut struct {
	TaskPeriod *TaskPeriodsPutTaskPeriod `json:"task_period,omitempty"`
	PreviousValues *TaskPeriodsPutPreviousValues `json:"previous_values,omitempty"`
}

// NewTaskPeriodsPut instantiates a new TaskPeriodsPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskPeriodsPut() *TaskPeriodsPut {
	this := TaskPeriodsPut{}
	return &this
}

// NewTaskPeriodsPutWithDefaults instantiates a new TaskPeriodsPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskPeriodsPutWithDefaults() *TaskPeriodsPut {
	this := TaskPeriodsPut{}
	return &this
}

// GetTaskPeriod returns the TaskPeriod field value if set, zero value otherwise.
func (o *TaskPeriodsPut) GetTaskPeriod() TaskPeriodsPutTaskPeriod {
	if o == nil || IsNil(o.TaskPeriod) {
		var ret TaskPeriodsPutTaskPeriod
		return ret
	}
	return *o.TaskPeriod
}

// GetTaskPeriodOk returns a tuple with the TaskPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPeriodsPut) GetTaskPeriodOk() (*TaskPeriodsPutTaskPeriod, bool) {
	if o == nil || IsNil(o.TaskPeriod) {
		return nil, false
	}
	return o.TaskPeriod, true
}

// HasTaskPeriod returns a boolean if a field has been set.
func (o *TaskPeriodsPut) HasTaskPeriod() bool {
	if o != nil && !IsNil(o.TaskPeriod) {
		return true
	}

	return false
}

// SetTaskPeriod gets a reference to the given TaskPeriodsPutTaskPeriod and assigns it to the TaskPeriod field.
func (o *TaskPeriodsPut) SetTaskPeriod(v TaskPeriodsPutTaskPeriod) {
	o.TaskPeriod = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *TaskPeriodsPut) GetPreviousValues() TaskPeriodsPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret TaskPeriodsPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPeriodsPut) GetPreviousValuesOk() (*TaskPeriodsPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *TaskPeriodsPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given TaskPeriodsPutPreviousValues and assigns it to the PreviousValues field.
func (o *TaskPeriodsPut) SetPreviousValues(v TaskPeriodsPutPreviousValues) {
	o.PreviousValues = &v
}

func (o TaskPeriodsPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskPeriodsPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TaskPeriod) {
		toSerialize["task_period"] = o.TaskPeriod
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableTaskPeriodsPut struct {
	value *TaskPeriodsPut
	isSet bool
}

func (v NullableTaskPeriodsPut) Get() *TaskPeriodsPut {
	return v.value
}

func (v *NullableTaskPeriodsPut) Set(val *TaskPeriodsPut) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskPeriodsPut) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskPeriodsPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskPeriodsPut(val *TaskPeriodsPut) *NullableTaskPeriodsPut {
	return &NullableTaskPeriodsPut{value: val, isSet: true}
}

func (v NullableTaskPeriodsPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskPeriodsPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

