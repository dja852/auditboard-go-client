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

// checks if the TestStop type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestStop{}

// TestStop struct for TestStop
type TestStop struct {
	Tests []Tests `json:"tests,omitempty"`
	GlobalAudits []GlobalAudits `json:"global_audits,omitempty"`
}

// NewTestStop instantiates a new TestStop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestStop() *TestStop {
	this := TestStop{}
	return &this
}

// NewTestStopWithDefaults instantiates a new TestStop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestStopWithDefaults() *TestStop {
	this := TestStop{}
	return &this
}

// GetTests returns the Tests field value if set, zero value otherwise.
func (o *TestStop) GetTests() []Tests {
	if o == nil || IsNil(o.Tests) {
		var ret []Tests
		return ret
	}
	return o.Tests
}

// GetTestsOk returns a tuple with the Tests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestStop) GetTestsOk() ([]Tests, bool) {
	if o == nil || IsNil(o.Tests) {
		return nil, false
	}
	return o.Tests, true
}

// HasTests returns a boolean if a field has been set.
func (o *TestStop) HasTests() bool {
	if o != nil && !IsNil(o.Tests) {
		return true
	}

	return false
}

// SetTests gets a reference to the given []Tests and assigns it to the Tests field.
func (o *TestStop) SetTests(v []Tests) {
	o.Tests = v
}

// GetGlobalAudits returns the GlobalAudits field value if set, zero value otherwise.
func (o *TestStop) GetGlobalAudits() []GlobalAudits {
	if o == nil || IsNil(o.GlobalAudits) {
		var ret []GlobalAudits
		return ret
	}
	return o.GlobalAudits
}

// GetGlobalAuditsOk returns a tuple with the GlobalAudits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestStop) GetGlobalAuditsOk() ([]GlobalAudits, bool) {
	if o == nil || IsNil(o.GlobalAudits) {
		return nil, false
	}
	return o.GlobalAudits, true
}

// HasGlobalAudits returns a boolean if a field has been set.
func (o *TestStop) HasGlobalAudits() bool {
	if o != nil && !IsNil(o.GlobalAudits) {
		return true
	}

	return false
}

// SetGlobalAudits gets a reference to the given []GlobalAudits and assigns it to the GlobalAudits field.
func (o *TestStop) SetGlobalAudits(v []GlobalAudits) {
	o.GlobalAudits = v
}

func (o TestStop) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestStop) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tests) {
		toSerialize["tests"] = o.Tests
	}
	if !IsNil(o.GlobalAudits) {
		toSerialize["global_audits"] = o.GlobalAudits
	}
	return toSerialize, nil
}

type NullableTestStop struct {
	value *TestStop
	isSet bool
}

func (v NullableTestStop) Get() *TestStop {
	return v.value
}

func (v *NullableTestStop) Set(val *TestStop) {
	v.value = val
	v.isSet = true
}

func (v NullableTestStop) IsSet() bool {
	return v.isSet
}

func (v *NullableTestStop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestStop(val *TestStop) *NullableTestStop {
	return &NullableTestStop{value: val, isSet: true}
}

func (v NullableTestStop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestStop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


