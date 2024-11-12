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

// checks if the FrameworkItemsPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrameworkItemsPut{}

// FrameworkItemsPut struct for FrameworkItemsPut
type FrameworkItemsPut struct {
	FrameworkItem *FrameworkItemsPutFrameworkItem `json:"framework_item,omitempty"`
	PreviousValues *FrameworkItemsPutPreviousValues `json:"previous_values,omitempty"`
}

// NewFrameworkItemsPut instantiates a new FrameworkItemsPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameworkItemsPut() *FrameworkItemsPut {
	this := FrameworkItemsPut{}
	return &this
}

// NewFrameworkItemsPutWithDefaults instantiates a new FrameworkItemsPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameworkItemsPutWithDefaults() *FrameworkItemsPut {
	this := FrameworkItemsPut{}
	return &this
}

// GetFrameworkItem returns the FrameworkItem field value if set, zero value otherwise.
func (o *FrameworkItemsPut) GetFrameworkItem() FrameworkItemsPutFrameworkItem {
	if o == nil || IsNil(o.FrameworkItem) {
		var ret FrameworkItemsPutFrameworkItem
		return ret
	}
	return *o.FrameworkItem
}

// GetFrameworkItemOk returns a tuple with the FrameworkItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameworkItemsPut) GetFrameworkItemOk() (*FrameworkItemsPutFrameworkItem, bool) {
	if o == nil || IsNil(o.FrameworkItem) {
		return nil, false
	}
	return o.FrameworkItem, true
}

// HasFrameworkItem returns a boolean if a field has been set.
func (o *FrameworkItemsPut) HasFrameworkItem() bool {
	if o != nil && !IsNil(o.FrameworkItem) {
		return true
	}

	return false
}

// SetFrameworkItem gets a reference to the given FrameworkItemsPutFrameworkItem and assigns it to the FrameworkItem field.
func (o *FrameworkItemsPut) SetFrameworkItem(v FrameworkItemsPutFrameworkItem) {
	o.FrameworkItem = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *FrameworkItemsPut) GetPreviousValues() FrameworkItemsPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret FrameworkItemsPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameworkItemsPut) GetPreviousValuesOk() (*FrameworkItemsPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *FrameworkItemsPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given FrameworkItemsPutPreviousValues and assigns it to the PreviousValues field.
func (o *FrameworkItemsPut) SetPreviousValues(v FrameworkItemsPutPreviousValues) {
	o.PreviousValues = &v
}

func (o FrameworkItemsPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrameworkItemsPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FrameworkItem) {
		toSerialize["framework_item"] = o.FrameworkItem
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableFrameworkItemsPut struct {
	value *FrameworkItemsPut
	isSet bool
}

func (v NullableFrameworkItemsPut) Get() *FrameworkItemsPut {
	return v.value
}

func (v *NullableFrameworkItemsPut) Set(val *FrameworkItemsPut) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameworkItemsPut) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameworkItemsPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameworkItemsPut(val *FrameworkItemsPut) *NullableFrameworkItemsPut {
	return &NullableFrameworkItemsPut{value: val, isSet: true}
}

func (v NullableFrameworkItemsPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameworkItemsPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

