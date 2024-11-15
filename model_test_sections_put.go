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

// checks if the TestSectionsPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestSectionsPut{}

// TestSectionsPut struct for TestSectionsPut
type TestSectionsPut struct {
	TestSection *TestSectionsPutTestSection `json:"test_section,omitempty"`
	PreviousValues *TestSectionsPutPreviousValues `json:"previous_values,omitempty"`
}

// NewTestSectionsPut instantiates a new TestSectionsPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestSectionsPut() *TestSectionsPut {
	this := TestSectionsPut{}
	return &this
}

// NewTestSectionsPutWithDefaults instantiates a new TestSectionsPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestSectionsPutWithDefaults() *TestSectionsPut {
	this := TestSectionsPut{}
	return &this
}

// GetTestSection returns the TestSection field value if set, zero value otherwise.
func (o *TestSectionsPut) GetTestSection() TestSectionsPutTestSection {
	if o == nil || IsNil(o.TestSection) {
		var ret TestSectionsPutTestSection
		return ret
	}
	return *o.TestSection
}

// GetTestSectionOk returns a tuple with the TestSection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSectionsPut) GetTestSectionOk() (*TestSectionsPutTestSection, bool) {
	if o == nil || IsNil(o.TestSection) {
		return nil, false
	}
	return o.TestSection, true
}

// HasTestSection returns a boolean if a field has been set.
func (o *TestSectionsPut) HasTestSection() bool {
	if o != nil && !IsNil(o.TestSection) {
		return true
	}

	return false
}

// SetTestSection gets a reference to the given TestSectionsPutTestSection and assigns it to the TestSection field.
func (o *TestSectionsPut) SetTestSection(v TestSectionsPutTestSection) {
	o.TestSection = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *TestSectionsPut) GetPreviousValues() TestSectionsPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret TestSectionsPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSectionsPut) GetPreviousValuesOk() (*TestSectionsPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *TestSectionsPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given TestSectionsPutPreviousValues and assigns it to the PreviousValues field.
func (o *TestSectionsPut) SetPreviousValues(v TestSectionsPutPreviousValues) {
	o.PreviousValues = &v
}

func (o TestSectionsPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestSectionsPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TestSection) {
		toSerialize["test_section"] = o.TestSection
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableTestSectionsPut struct {
	value *TestSectionsPut
	isSet bool
}

func (v NullableTestSectionsPut) Get() *TestSectionsPut {
	return v.value
}

func (v *NullableTestSectionsPut) Set(val *TestSectionsPut) {
	v.value = val
	v.isSet = true
}

func (v NullableTestSectionsPut) IsSet() bool {
	return v.isSet
}

func (v *NullableTestSectionsPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestSectionsPut(val *TestSectionsPut) *NullableTestSectionsPut {
	return &NullableTestSectionsPut{value: val, isSet: true}
}

func (v NullableTestSectionsPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestSectionsPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


