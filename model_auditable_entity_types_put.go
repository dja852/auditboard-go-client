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

// checks if the AuditableEntityTypesPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditableEntityTypesPut{}

// AuditableEntityTypesPut struct for AuditableEntityTypesPut
type AuditableEntityTypesPut struct {
	AuditableEntityType *AuditableEntityTypesPutAuditableEntityType `json:"auditable_entity_type,omitempty"`
	PreviousValues *AuditableEntityTypesPutPreviousValues `json:"previous_values,omitempty"`
}

// NewAuditableEntityTypesPut instantiates a new AuditableEntityTypesPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditableEntityTypesPut() *AuditableEntityTypesPut {
	this := AuditableEntityTypesPut{}
	return &this
}

// NewAuditableEntityTypesPutWithDefaults instantiates a new AuditableEntityTypesPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditableEntityTypesPutWithDefaults() *AuditableEntityTypesPut {
	this := AuditableEntityTypesPut{}
	return &this
}

// GetAuditableEntityType returns the AuditableEntityType field value if set, zero value otherwise.
func (o *AuditableEntityTypesPut) GetAuditableEntityType() AuditableEntityTypesPutAuditableEntityType {
	if o == nil || IsNil(o.AuditableEntityType) {
		var ret AuditableEntityTypesPutAuditableEntityType
		return ret
	}
	return *o.AuditableEntityType
}

// GetAuditableEntityTypeOk returns a tuple with the AuditableEntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditableEntityTypesPut) GetAuditableEntityTypeOk() (*AuditableEntityTypesPutAuditableEntityType, bool) {
	if o == nil || IsNil(o.AuditableEntityType) {
		return nil, false
	}
	return o.AuditableEntityType, true
}

// HasAuditableEntityType returns a boolean if a field has been set.
func (o *AuditableEntityTypesPut) HasAuditableEntityType() bool {
	if o != nil && !IsNil(o.AuditableEntityType) {
		return true
	}

	return false
}

// SetAuditableEntityType gets a reference to the given AuditableEntityTypesPutAuditableEntityType and assigns it to the AuditableEntityType field.
func (o *AuditableEntityTypesPut) SetAuditableEntityType(v AuditableEntityTypesPutAuditableEntityType) {
	o.AuditableEntityType = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *AuditableEntityTypesPut) GetPreviousValues() AuditableEntityTypesPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret AuditableEntityTypesPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditableEntityTypesPut) GetPreviousValuesOk() (*AuditableEntityTypesPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *AuditableEntityTypesPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given AuditableEntityTypesPutPreviousValues and assigns it to the PreviousValues field.
func (o *AuditableEntityTypesPut) SetPreviousValues(v AuditableEntityTypesPutPreviousValues) {
	o.PreviousValues = &v
}

func (o AuditableEntityTypesPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditableEntityTypesPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuditableEntityType) {
		toSerialize["auditable_entity_type"] = o.AuditableEntityType
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableAuditableEntityTypesPut struct {
	value *AuditableEntityTypesPut
	isSet bool
}

func (v NullableAuditableEntityTypesPut) Get() *AuditableEntityTypesPut {
	return v.value
}

func (v *NullableAuditableEntityTypesPut) Set(val *AuditableEntityTypesPut) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditableEntityTypesPut) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditableEntityTypesPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditableEntityTypesPut(val *AuditableEntityTypesPut) *NullableAuditableEntityTypesPut {
	return &NullableAuditableEntityTypesPut{value: val, isSet: true}
}

func (v NullableAuditableEntityTypesPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditableEntityTypesPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


