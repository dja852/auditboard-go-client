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

// checks if the GlobalAuditsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalAuditsGet200Response{}

// GlobalAuditsGet200Response struct for GlobalAuditsGet200Response
type GlobalAuditsGet200Response struct {
	GlobalAudits []GlobalAudits `json:"global_audits,omitempty"`
}

// NewGlobalAuditsGet200Response instantiates a new GlobalAuditsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalAuditsGet200Response() *GlobalAuditsGet200Response {
	this := GlobalAuditsGet200Response{}
	return &this
}

// NewGlobalAuditsGet200ResponseWithDefaults instantiates a new GlobalAuditsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalAuditsGet200ResponseWithDefaults() *GlobalAuditsGet200Response {
	this := GlobalAuditsGet200Response{}
	return &this
}

// GetGlobalAudits returns the GlobalAudits field value if set, zero value otherwise.
func (o *GlobalAuditsGet200Response) GetGlobalAudits() []GlobalAudits {
	if o == nil || IsNil(o.GlobalAudits) {
		var ret []GlobalAudits
		return ret
	}
	return o.GlobalAudits
}

// GetGlobalAuditsOk returns a tuple with the GlobalAudits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalAuditsGet200Response) GetGlobalAuditsOk() ([]GlobalAudits, bool) {
	if o == nil || IsNil(o.GlobalAudits) {
		return nil, false
	}
	return o.GlobalAudits, true
}

// HasGlobalAudits returns a boolean if a field has been set.
func (o *GlobalAuditsGet200Response) HasGlobalAudits() bool {
	if o != nil && !IsNil(o.GlobalAudits) {
		return true
	}

	return false
}

// SetGlobalAudits gets a reference to the given []GlobalAudits and assigns it to the GlobalAudits field.
func (o *GlobalAuditsGet200Response) SetGlobalAudits(v []GlobalAudits) {
	o.GlobalAudits = v
}

func (o GlobalAuditsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalAuditsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GlobalAudits) {
		toSerialize["global_audits"] = o.GlobalAudits
	}
	return toSerialize, nil
}

type NullableGlobalAuditsGet200Response struct {
	value *GlobalAuditsGet200Response
	isSet bool
}

func (v NullableGlobalAuditsGet200Response) Get() *GlobalAuditsGet200Response {
	return v.value
}

func (v *NullableGlobalAuditsGet200Response) Set(val *GlobalAuditsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalAuditsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalAuditsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalAuditsGet200Response(val *GlobalAuditsGet200Response) *NullableGlobalAuditsGet200Response {
	return &NullableGlobalAuditsGet200Response{value: val, isSet: true}
}

func (v NullableGlobalAuditsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalAuditsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

