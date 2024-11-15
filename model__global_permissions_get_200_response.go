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

// checks if the GlobalPermissionsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalPermissionsGet200Response{}

// GlobalPermissionsGet200Response struct for GlobalPermissionsGet200Response
type GlobalPermissionsGet200Response struct {
	GlobalPermissions []GlobalPermissions `json:"global_permissions,omitempty"`
}

// NewGlobalPermissionsGet200Response instantiates a new GlobalPermissionsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalPermissionsGet200Response() *GlobalPermissionsGet200Response {
	this := GlobalPermissionsGet200Response{}
	return &this
}

// NewGlobalPermissionsGet200ResponseWithDefaults instantiates a new GlobalPermissionsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalPermissionsGet200ResponseWithDefaults() *GlobalPermissionsGet200Response {
	this := GlobalPermissionsGet200Response{}
	return &this
}

// GetGlobalPermissions returns the GlobalPermissions field value if set, zero value otherwise.
func (o *GlobalPermissionsGet200Response) GetGlobalPermissions() []GlobalPermissions {
	if o == nil || IsNil(o.GlobalPermissions) {
		var ret []GlobalPermissions
		return ret
	}
	return o.GlobalPermissions
}

// GetGlobalPermissionsOk returns a tuple with the GlobalPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionsGet200Response) GetGlobalPermissionsOk() ([]GlobalPermissions, bool) {
	if o == nil || IsNil(o.GlobalPermissions) {
		return nil, false
	}
	return o.GlobalPermissions, true
}

// HasGlobalPermissions returns a boolean if a field has been set.
func (o *GlobalPermissionsGet200Response) HasGlobalPermissions() bool {
	if o != nil && !IsNil(o.GlobalPermissions) {
		return true
	}

	return false
}

// SetGlobalPermissions gets a reference to the given []GlobalPermissions and assigns it to the GlobalPermissions field.
func (o *GlobalPermissionsGet200Response) SetGlobalPermissions(v []GlobalPermissions) {
	o.GlobalPermissions = v
}

func (o GlobalPermissionsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalPermissionsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GlobalPermissions) {
		toSerialize["global_permissions"] = o.GlobalPermissions
	}
	return toSerialize, nil
}

type NullableGlobalPermissionsGet200Response struct {
	value *GlobalPermissionsGet200Response
	isSet bool
}

func (v NullableGlobalPermissionsGet200Response) Get() *GlobalPermissionsGet200Response {
	return v.value
}

func (v *NullableGlobalPermissionsGet200Response) Set(val *GlobalPermissionsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalPermissionsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalPermissionsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalPermissionsGet200Response(val *GlobalPermissionsGet200Response) *NullableGlobalPermissionsGet200Response {
	return &NullableGlobalPermissionsGet200Response{value: val, isSet: true}
}

func (v NullableGlobalPermissionsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalPermissionsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


