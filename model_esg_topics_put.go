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

// checks if the EsgTopicsPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsgTopicsPut{}

// EsgTopicsPut struct for EsgTopicsPut
type EsgTopicsPut struct {
	EsgTopic *EsgTopicsPutEsgTopic `json:"esg_topic,omitempty"`
	PreviousValues *EsgTopicsPutPreviousValues `json:"previous_values,omitempty"`
}

// NewEsgTopicsPut instantiates a new EsgTopicsPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsgTopicsPut() *EsgTopicsPut {
	this := EsgTopicsPut{}
	return &this
}

// NewEsgTopicsPutWithDefaults instantiates a new EsgTopicsPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsgTopicsPutWithDefaults() *EsgTopicsPut {
	this := EsgTopicsPut{}
	return &this
}

// GetEsgTopic returns the EsgTopic field value if set, zero value otherwise.
func (o *EsgTopicsPut) GetEsgTopic() EsgTopicsPutEsgTopic {
	if o == nil || IsNil(o.EsgTopic) {
		var ret EsgTopicsPutEsgTopic
		return ret
	}
	return *o.EsgTopic
}

// GetEsgTopicOk returns a tuple with the EsgTopic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopicsPut) GetEsgTopicOk() (*EsgTopicsPutEsgTopic, bool) {
	if o == nil || IsNil(o.EsgTopic) {
		return nil, false
	}
	return o.EsgTopic, true
}

// HasEsgTopic returns a boolean if a field has been set.
func (o *EsgTopicsPut) HasEsgTopic() bool {
	if o != nil && !IsNil(o.EsgTopic) {
		return true
	}

	return false
}

// SetEsgTopic gets a reference to the given EsgTopicsPutEsgTopic and assigns it to the EsgTopic field.
func (o *EsgTopicsPut) SetEsgTopic(v EsgTopicsPutEsgTopic) {
	o.EsgTopic = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *EsgTopicsPut) GetPreviousValues() EsgTopicsPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret EsgTopicsPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopicsPut) GetPreviousValuesOk() (*EsgTopicsPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *EsgTopicsPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given EsgTopicsPutPreviousValues and assigns it to the PreviousValues field.
func (o *EsgTopicsPut) SetPreviousValues(v EsgTopicsPutPreviousValues) {
	o.PreviousValues = &v
}

func (o EsgTopicsPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsgTopicsPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EsgTopic) {
		toSerialize["esg_topic"] = o.EsgTopic
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableEsgTopicsPut struct {
	value *EsgTopicsPut
	isSet bool
}

func (v NullableEsgTopicsPut) Get() *EsgTopicsPut {
	return v.value
}

func (v *NullableEsgTopicsPut) Set(val *EsgTopicsPut) {
	v.value = val
	v.isSet = true
}

func (v NullableEsgTopicsPut) IsSet() bool {
	return v.isSet
}

func (v *NullableEsgTopicsPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsgTopicsPut(val *EsgTopicsPut) *NullableEsgTopicsPut {
	return &NullableEsgTopicsPut{value: val, isSet: true}
}

func (v NullableEsgTopicsPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsgTopicsPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

