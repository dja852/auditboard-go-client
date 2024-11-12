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

// checks if the ActionPlansArchivesPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionPlansArchivesPut{}

// ActionPlansArchivesPut struct for ActionPlansArchivesPut
type ActionPlansArchivesPut struct {
	ActionPlansArchive *ActionPlansArchivesPutActionPlansArchive `json:"action_plans_archive,omitempty"`
	PreviousValues *ActionPlansArchivesPutPreviousValues `json:"previous_values,omitempty"`
}

// NewActionPlansArchivesPut instantiates a new ActionPlansArchivesPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionPlansArchivesPut() *ActionPlansArchivesPut {
	this := ActionPlansArchivesPut{}
	return &this
}

// NewActionPlansArchivesPutWithDefaults instantiates a new ActionPlansArchivesPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionPlansArchivesPutWithDefaults() *ActionPlansArchivesPut {
	this := ActionPlansArchivesPut{}
	return &this
}

// GetActionPlansArchive returns the ActionPlansArchive field value if set, zero value otherwise.
func (o *ActionPlansArchivesPut) GetActionPlansArchive() ActionPlansArchivesPutActionPlansArchive {
	if o == nil || IsNil(o.ActionPlansArchive) {
		var ret ActionPlansArchivesPutActionPlansArchive
		return ret
	}
	return *o.ActionPlansArchive
}

// GetActionPlansArchiveOk returns a tuple with the ActionPlansArchive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionPlansArchivesPut) GetActionPlansArchiveOk() (*ActionPlansArchivesPutActionPlansArchive, bool) {
	if o == nil || IsNil(o.ActionPlansArchive) {
		return nil, false
	}
	return o.ActionPlansArchive, true
}

// HasActionPlansArchive returns a boolean if a field has been set.
func (o *ActionPlansArchivesPut) HasActionPlansArchive() bool {
	if o != nil && !IsNil(o.ActionPlansArchive) {
		return true
	}

	return false
}

// SetActionPlansArchive gets a reference to the given ActionPlansArchivesPutActionPlansArchive and assigns it to the ActionPlansArchive field.
func (o *ActionPlansArchivesPut) SetActionPlansArchive(v ActionPlansArchivesPutActionPlansArchive) {
	o.ActionPlansArchive = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *ActionPlansArchivesPut) GetPreviousValues() ActionPlansArchivesPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret ActionPlansArchivesPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionPlansArchivesPut) GetPreviousValuesOk() (*ActionPlansArchivesPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *ActionPlansArchivesPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given ActionPlansArchivesPutPreviousValues and assigns it to the PreviousValues field.
func (o *ActionPlansArchivesPut) SetPreviousValues(v ActionPlansArchivesPutPreviousValues) {
	o.PreviousValues = &v
}

func (o ActionPlansArchivesPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionPlansArchivesPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionPlansArchive) {
		toSerialize["action_plans_archive"] = o.ActionPlansArchive
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableActionPlansArchivesPut struct {
	value *ActionPlansArchivesPut
	isSet bool
}

func (v NullableActionPlansArchivesPut) Get() *ActionPlansArchivesPut {
	return v.value
}

func (v *NullableActionPlansArchivesPut) Set(val *ActionPlansArchivesPut) {
	v.value = val
	v.isSet = true
}

func (v NullableActionPlansArchivesPut) IsSet() bool {
	return v.isSet
}

func (v *NullableActionPlansArchivesPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionPlansArchivesPut(val *ActionPlansArchivesPut) *NullableActionPlansArchivesPut {
	return &NullableActionPlansArchivesPut{value: val, isSet: true}
}

func (v NullableActionPlansArchivesPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionPlansArchivesPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

