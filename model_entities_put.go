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

// checks if the EntitiesPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitiesPut{}

// EntitiesPut struct for EntitiesPut
type EntitiesPut struct {
	Entity *EntitiesPutEntity `json:"entity,omitempty"`
	PreviousValues *EntitiesPutPreviousValues `json:"previous_values,omitempty"`
}

// NewEntitiesPut instantiates a new EntitiesPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitiesPut() *EntitiesPut {
	this := EntitiesPut{}
	return &this
}

// NewEntitiesPutWithDefaults instantiates a new EntitiesPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitiesPutWithDefaults() *EntitiesPut {
	this := EntitiesPut{}
	return &this
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *EntitiesPut) GetEntity() EntitiesPutEntity {
	if o == nil || IsNil(o.Entity) {
		var ret EntitiesPutEntity
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitiesPut) GetEntityOk() (*EntitiesPutEntity, bool) {
	if o == nil || IsNil(o.Entity) {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *EntitiesPut) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given EntitiesPutEntity and assigns it to the Entity field.
func (o *EntitiesPut) SetEntity(v EntitiesPutEntity) {
	o.Entity = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *EntitiesPut) GetPreviousValues() EntitiesPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret EntitiesPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitiesPut) GetPreviousValuesOk() (*EntitiesPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *EntitiesPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given EntitiesPutPreviousValues and assigns it to the PreviousValues field.
func (o *EntitiesPut) SetPreviousValues(v EntitiesPutPreviousValues) {
	o.PreviousValues = &v
}

func (o EntitiesPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitiesPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableEntitiesPut struct {
	value *EntitiesPut
	isSet bool
}

func (v NullableEntitiesPut) Get() *EntitiesPut {
	return v.value
}

func (v *NullableEntitiesPut) Set(val *EntitiesPut) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitiesPut) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitiesPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitiesPut(val *EntitiesPut) *NullableEntitiesPut {
	return &NullableEntitiesPut{value: val, isSet: true}
}

func (v NullableEntitiesPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitiesPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


