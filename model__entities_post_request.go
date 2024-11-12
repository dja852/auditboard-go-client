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

// checks if the EntitiesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitiesPostRequest{}

// EntitiesPostRequest struct for EntitiesPostRequest
type EntitiesPostRequest struct {
	Entity *Entities `json:"entity,omitempty"`
}

// NewEntitiesPostRequest instantiates a new EntitiesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitiesPostRequest() *EntitiesPostRequest {
	this := EntitiesPostRequest{}
	return &this
}

// NewEntitiesPostRequestWithDefaults instantiates a new EntitiesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitiesPostRequestWithDefaults() *EntitiesPostRequest {
	this := EntitiesPostRequest{}
	return &this
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *EntitiesPostRequest) GetEntity() Entities {
	if o == nil || IsNil(o.Entity) {
		var ret Entities
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitiesPostRequest) GetEntityOk() (*Entities, bool) {
	if o == nil || IsNil(o.Entity) {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *EntitiesPostRequest) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given Entities and assigns it to the Entity field.
func (o *EntitiesPostRequest) SetEntity(v Entities) {
	o.Entity = &v
}

func (o EntitiesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitiesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	return toSerialize, nil
}

type NullableEntitiesPostRequest struct {
	value *EntitiesPostRequest
	isSet bool
}

func (v NullableEntitiesPostRequest) Get() *EntitiesPostRequest {
	return v.value
}

func (v *NullableEntitiesPostRequest) Set(val *EntitiesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitiesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitiesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitiesPostRequest(val *EntitiesPostRequest) *NullableEntitiesPostRequest {
	return &NullableEntitiesPostRequest{value: val, isSet: true}
}

func (v NullableEntitiesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitiesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

