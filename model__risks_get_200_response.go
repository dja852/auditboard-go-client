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

// checks if the RisksGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RisksGet200Response{}

// RisksGet200Response struct for RisksGet200Response
type RisksGet200Response struct {
	Risks []Risks `json:"risks,omitempty"`
}

// NewRisksGet200Response instantiates a new RisksGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRisksGet200Response() *RisksGet200Response {
	this := RisksGet200Response{}
	return &this
}

// NewRisksGet200ResponseWithDefaults instantiates a new RisksGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRisksGet200ResponseWithDefaults() *RisksGet200Response {
	this := RisksGet200Response{}
	return &this
}

// GetRisks returns the Risks field value if set, zero value otherwise.
func (o *RisksGet200Response) GetRisks() []Risks {
	if o == nil || IsNil(o.Risks) {
		var ret []Risks
		return ret
	}
	return o.Risks
}

// GetRisksOk returns a tuple with the Risks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RisksGet200Response) GetRisksOk() ([]Risks, bool) {
	if o == nil || IsNil(o.Risks) {
		return nil, false
	}
	return o.Risks, true
}

// HasRisks returns a boolean if a field has been set.
func (o *RisksGet200Response) HasRisks() bool {
	if o != nil && !IsNil(o.Risks) {
		return true
	}

	return false
}

// SetRisks gets a reference to the given []Risks and assigns it to the Risks field.
func (o *RisksGet200Response) SetRisks(v []Risks) {
	o.Risks = v
}

func (o RisksGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RisksGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Risks) {
		toSerialize["risks"] = o.Risks
	}
	return toSerialize, nil
}

type NullableRisksGet200Response struct {
	value *RisksGet200Response
	isSet bool
}

func (v NullableRisksGet200Response) Get() *RisksGet200Response {
	return v.value
}

func (v *NullableRisksGet200Response) Set(val *RisksGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRisksGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRisksGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRisksGet200Response(val *RisksGet200Response) *NullableRisksGet200Response {
	return &NullableRisksGet200Response{value: val, isSet: true}
}

func (v NullableRisksGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRisksGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


