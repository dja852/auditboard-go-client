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

// checks if the RcwProjectsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RcwProjectsGet200Response{}

// RcwProjectsGet200Response struct for RcwProjectsGet200Response
type RcwProjectsGet200Response struct {
	RcwProjects []RcwProjects `json:"rcw_projects,omitempty"`
}

// NewRcwProjectsGet200Response instantiates a new RcwProjectsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRcwProjectsGet200Response() *RcwProjectsGet200Response {
	this := RcwProjectsGet200Response{}
	return &this
}

// NewRcwProjectsGet200ResponseWithDefaults instantiates a new RcwProjectsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRcwProjectsGet200ResponseWithDefaults() *RcwProjectsGet200Response {
	this := RcwProjectsGet200Response{}
	return &this
}

// GetRcwProjects returns the RcwProjects field value if set, zero value otherwise.
func (o *RcwProjectsGet200Response) GetRcwProjects() []RcwProjects {
	if o == nil || IsNil(o.RcwProjects) {
		var ret []RcwProjects
		return ret
	}
	return o.RcwProjects
}

// GetRcwProjectsOk returns a tuple with the RcwProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwProjectsGet200Response) GetRcwProjectsOk() ([]RcwProjects, bool) {
	if o == nil || IsNil(o.RcwProjects) {
		return nil, false
	}
	return o.RcwProjects, true
}

// HasRcwProjects returns a boolean if a field has been set.
func (o *RcwProjectsGet200Response) HasRcwProjects() bool {
	if o != nil && !IsNil(o.RcwProjects) {
		return true
	}

	return false
}

// SetRcwProjects gets a reference to the given []RcwProjects and assigns it to the RcwProjects field.
func (o *RcwProjectsGet200Response) SetRcwProjects(v []RcwProjects) {
	o.RcwProjects = v
}

func (o RcwProjectsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RcwProjectsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RcwProjects) {
		toSerialize["rcw_projects"] = o.RcwProjects
	}
	return toSerialize, nil
}

type NullableRcwProjectsGet200Response struct {
	value *RcwProjectsGet200Response
	isSet bool
}

func (v NullableRcwProjectsGet200Response) Get() *RcwProjectsGet200Response {
	return v.value
}

func (v *NullableRcwProjectsGet200Response) Set(val *RcwProjectsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRcwProjectsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRcwProjectsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRcwProjectsGet200Response(val *RcwProjectsGet200Response) *NullableRcwProjectsGet200Response {
	return &NullableRcwProjectsGet200Response{value: val, isSet: true}
}

func (v NullableRcwProjectsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRcwProjectsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


