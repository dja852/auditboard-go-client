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

// checks if the UsersPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsersPut{}

// UsersPut struct for UsersPut
type UsersPut struct {
	User *UsersPutUser `json:"user,omitempty"`
	PreviousValues *UsersPutPreviousValues `json:"previous_values,omitempty"`
}

// NewUsersPut instantiates a new UsersPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersPut() *UsersPut {
	this := UsersPut{}
	return &this
}

// NewUsersPutWithDefaults instantiates a new UsersPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersPutWithDefaults() *UsersPut {
	this := UsersPut{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UsersPut) GetUser() UsersPutUser {
	if o == nil || IsNil(o.User) {
		var ret UsersPutUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersPut) GetUserOk() (*UsersPutUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UsersPut) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UsersPutUser and assigns it to the User field.
func (o *UsersPut) SetUser(v UsersPutUser) {
	o.User = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *UsersPut) GetPreviousValues() UsersPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret UsersPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersPut) GetPreviousValuesOk() (*UsersPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *UsersPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given UsersPutPreviousValues and assigns it to the PreviousValues field.
func (o *UsersPut) SetPreviousValues(v UsersPutPreviousValues) {
	o.PreviousValues = &v
}

func (o UsersPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsersPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableUsersPut struct {
	value *UsersPut
	isSet bool
}

func (v NullableUsersPut) Get() *UsersPut {
	return v.value
}

func (v *NullableUsersPut) Set(val *UsersPut) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersPut) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersPut(val *UsersPut) *NullableUsersPut {
	return &NullableUsersPut{value: val, isSet: true}
}

func (v NullableUsersPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


