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

// checks if the PoliciesPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoliciesPut{}

// PoliciesPut struct for PoliciesPut
type PoliciesPut struct {
	Policy *PoliciesPutPolicy `json:"policy,omitempty"`
	PreviousValues *PoliciesPutPreviousValues `json:"previous_values,omitempty"`
}

// NewPoliciesPut instantiates a new PoliciesPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoliciesPut() *PoliciesPut {
	this := PoliciesPut{}
	return &this
}

// NewPoliciesPutWithDefaults instantiates a new PoliciesPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoliciesPutWithDefaults() *PoliciesPut {
	this := PoliciesPut{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *PoliciesPut) GetPolicy() PoliciesPutPolicy {
	if o == nil || IsNil(o.Policy) {
		var ret PoliciesPutPolicy
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesPut) GetPolicyOk() (*PoliciesPutPolicy, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *PoliciesPut) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given PoliciesPutPolicy and assigns it to the Policy field.
func (o *PoliciesPut) SetPolicy(v PoliciesPutPolicy) {
	o.Policy = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *PoliciesPut) GetPreviousValues() PoliciesPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret PoliciesPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesPut) GetPreviousValuesOk() (*PoliciesPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *PoliciesPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given PoliciesPutPreviousValues and assigns it to the PreviousValues field.
func (o *PoliciesPut) SetPreviousValues(v PoliciesPutPreviousValues) {
	o.PreviousValues = &v
}

func (o PoliciesPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoliciesPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullablePoliciesPut struct {
	value *PoliciesPut
	isSet bool
}

func (v NullablePoliciesPut) Get() *PoliciesPut {
	return v.value
}

func (v *NullablePoliciesPut) Set(val *PoliciesPut) {
	v.value = val
	v.isSet = true
}

func (v NullablePoliciesPut) IsSet() bool {
	return v.isSet
}

func (v *NullablePoliciesPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoliciesPut(val *PoliciesPut) *NullablePoliciesPut {
	return &NullablePoliciesPut{value: val, isSet: true}
}

func (v NullablePoliciesPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoliciesPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


