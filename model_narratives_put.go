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

// checks if the NarrativesPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NarrativesPut{}

// NarrativesPut struct for NarrativesPut
type NarrativesPut struct {
	Narrative *NarrativesPutNarrative `json:"narrative,omitempty"`
	PreviousValues *NarrativesPutPreviousValues `json:"previous_values,omitempty"`
}

// NewNarrativesPut instantiates a new NarrativesPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNarrativesPut() *NarrativesPut {
	this := NarrativesPut{}
	return &this
}

// NewNarrativesPutWithDefaults instantiates a new NarrativesPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNarrativesPutWithDefaults() *NarrativesPut {
	this := NarrativesPut{}
	return &this
}

// GetNarrative returns the Narrative field value if set, zero value otherwise.
func (o *NarrativesPut) GetNarrative() NarrativesPutNarrative {
	if o == nil || IsNil(o.Narrative) {
		var ret NarrativesPutNarrative
		return ret
	}
	return *o.Narrative
}

// GetNarrativeOk returns a tuple with the Narrative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NarrativesPut) GetNarrativeOk() (*NarrativesPutNarrative, bool) {
	if o == nil || IsNil(o.Narrative) {
		return nil, false
	}
	return o.Narrative, true
}

// HasNarrative returns a boolean if a field has been set.
func (o *NarrativesPut) HasNarrative() bool {
	if o != nil && !IsNil(o.Narrative) {
		return true
	}

	return false
}

// SetNarrative gets a reference to the given NarrativesPutNarrative and assigns it to the Narrative field.
func (o *NarrativesPut) SetNarrative(v NarrativesPutNarrative) {
	o.Narrative = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *NarrativesPut) GetPreviousValues() NarrativesPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret NarrativesPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NarrativesPut) GetPreviousValuesOk() (*NarrativesPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *NarrativesPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given NarrativesPutPreviousValues and assigns it to the PreviousValues field.
func (o *NarrativesPut) SetPreviousValues(v NarrativesPutPreviousValues) {
	o.PreviousValues = &v
}

func (o NarrativesPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NarrativesPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Narrative) {
		toSerialize["narrative"] = o.Narrative
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableNarrativesPut struct {
	value *NarrativesPut
	isSet bool
}

func (v NullableNarrativesPut) Get() *NarrativesPut {
	return v.value
}

func (v *NullableNarrativesPut) Set(val *NarrativesPut) {
	v.value = val
	v.isSet = true
}

func (v NullableNarrativesPut) IsSet() bool {
	return v.isSet
}

func (v *NullableNarrativesPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNarrativesPut(val *NarrativesPut) *NullableNarrativesPut {
	return &NullableNarrativesPut{value: val, isSet: true}
}

func (v NullableNarrativesPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNarrativesPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

