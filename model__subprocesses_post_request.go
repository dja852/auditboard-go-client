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

// checks if the SubprocessesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubprocessesPostRequest{}

// SubprocessesPostRequest struct for SubprocessesPostRequest
type SubprocessesPostRequest struct {
	Subprocess *Subprocesses `json:"subprocess,omitempty"`
}

// NewSubprocessesPostRequest instantiates a new SubprocessesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubprocessesPostRequest() *SubprocessesPostRequest {
	this := SubprocessesPostRequest{}
	return &this
}

// NewSubprocessesPostRequestWithDefaults instantiates a new SubprocessesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubprocessesPostRequestWithDefaults() *SubprocessesPostRequest {
	this := SubprocessesPostRequest{}
	return &this
}

// GetSubprocess returns the Subprocess field value if set, zero value otherwise.
func (o *SubprocessesPostRequest) GetSubprocess() Subprocesses {
	if o == nil || IsNil(o.Subprocess) {
		var ret Subprocesses
		return ret
	}
	return *o.Subprocess
}

// GetSubprocessOk returns a tuple with the Subprocess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubprocessesPostRequest) GetSubprocessOk() (*Subprocesses, bool) {
	if o == nil || IsNil(o.Subprocess) {
		return nil, false
	}
	return o.Subprocess, true
}

// HasSubprocess returns a boolean if a field has been set.
func (o *SubprocessesPostRequest) HasSubprocess() bool {
	if o != nil && !IsNil(o.Subprocess) {
		return true
	}

	return false
}

// SetSubprocess gets a reference to the given Subprocesses and assigns it to the Subprocess field.
func (o *SubprocessesPostRequest) SetSubprocess(v Subprocesses) {
	o.Subprocess = &v
}

func (o SubprocessesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubprocessesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subprocess) {
		toSerialize["subprocess"] = o.Subprocess
	}
	return toSerialize, nil
}

type NullableSubprocessesPostRequest struct {
	value *SubprocessesPostRequest
	isSet bool
}

func (v NullableSubprocessesPostRequest) Get() *SubprocessesPostRequest {
	return v.value
}

func (v *NullableSubprocessesPostRequest) Set(val *SubprocessesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubprocessesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubprocessesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubprocessesPostRequest(val *SubprocessesPostRequest) *NullableSubprocessesPostRequest {
	return &NullableSubprocessesPostRequest{value: val, isSet: true}
}

func (v NullableSubprocessesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubprocessesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


