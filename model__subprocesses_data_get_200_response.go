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

// checks if the SubprocessesDataGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubprocessesDataGet200Response{}

// SubprocessesDataGet200Response struct for SubprocessesDataGet200Response
type SubprocessesDataGet200Response struct {
	SubprocessesData []SubprocessesData `json:"subprocesses_data,omitempty"`
}

// NewSubprocessesDataGet200Response instantiates a new SubprocessesDataGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubprocessesDataGet200Response() *SubprocessesDataGet200Response {
	this := SubprocessesDataGet200Response{}
	return &this
}

// NewSubprocessesDataGet200ResponseWithDefaults instantiates a new SubprocessesDataGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubprocessesDataGet200ResponseWithDefaults() *SubprocessesDataGet200Response {
	this := SubprocessesDataGet200Response{}
	return &this
}

// GetSubprocessesData returns the SubprocessesData field value if set, zero value otherwise.
func (o *SubprocessesDataGet200Response) GetSubprocessesData() []SubprocessesData {
	if o == nil || IsNil(o.SubprocessesData) {
		var ret []SubprocessesData
		return ret
	}
	return o.SubprocessesData
}

// GetSubprocessesDataOk returns a tuple with the SubprocessesData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubprocessesDataGet200Response) GetSubprocessesDataOk() ([]SubprocessesData, bool) {
	if o == nil || IsNil(o.SubprocessesData) {
		return nil, false
	}
	return o.SubprocessesData, true
}

// HasSubprocessesData returns a boolean if a field has been set.
func (o *SubprocessesDataGet200Response) HasSubprocessesData() bool {
	if o != nil && !IsNil(o.SubprocessesData) {
		return true
	}

	return false
}

// SetSubprocessesData gets a reference to the given []SubprocessesData and assigns it to the SubprocessesData field.
func (o *SubprocessesDataGet200Response) SetSubprocessesData(v []SubprocessesData) {
	o.SubprocessesData = v
}

func (o SubprocessesDataGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubprocessesDataGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubprocessesData) {
		toSerialize["subprocesses_data"] = o.SubprocessesData
	}
	return toSerialize, nil
}

type NullableSubprocessesDataGet200Response struct {
	value *SubprocessesDataGet200Response
	isSet bool
}

func (v NullableSubprocessesDataGet200Response) Get() *SubprocessesDataGet200Response {
	return v.value
}

func (v *NullableSubprocessesDataGet200Response) Set(val *SubprocessesDataGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSubprocessesDataGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSubprocessesDataGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubprocessesDataGet200Response(val *SubprocessesDataGet200Response) *NullableSubprocessesDataGet200Response {
	return &NullableSubprocessesDataGet200Response{value: val, isSet: true}
}

func (v NullableSubprocessesDataGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubprocessesDataGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

