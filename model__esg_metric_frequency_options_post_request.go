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

// checks if the EsgMetricFrequencyOptionsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsgMetricFrequencyOptionsPostRequest{}

// EsgMetricFrequencyOptionsPostRequest struct for EsgMetricFrequencyOptionsPostRequest
type EsgMetricFrequencyOptionsPostRequest struct {
	EsgMetricFrequencyOption *EsgMetricFrequencyOptions `json:"esg_metric _frequency _option,omitempty"`
}

// NewEsgMetricFrequencyOptionsPostRequest instantiates a new EsgMetricFrequencyOptionsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsgMetricFrequencyOptionsPostRequest() *EsgMetricFrequencyOptionsPostRequest {
	this := EsgMetricFrequencyOptionsPostRequest{}
	return &this
}

// NewEsgMetricFrequencyOptionsPostRequestWithDefaults instantiates a new EsgMetricFrequencyOptionsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsgMetricFrequencyOptionsPostRequestWithDefaults() *EsgMetricFrequencyOptionsPostRequest {
	this := EsgMetricFrequencyOptionsPostRequest{}
	return &this
}

// GetEsgMetricFrequencyOption returns the EsgMetricFrequencyOption field value if set, zero value otherwise.
func (o *EsgMetricFrequencyOptionsPostRequest) GetEsgMetricFrequencyOption() EsgMetricFrequencyOptions {
	if o == nil || IsNil(o.EsgMetricFrequencyOption) {
		var ret EsgMetricFrequencyOptions
		return ret
	}
	return *o.EsgMetricFrequencyOption
}

// GetEsgMetricFrequencyOptionOk returns a tuple with the EsgMetricFrequencyOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricFrequencyOptionsPostRequest) GetEsgMetricFrequencyOptionOk() (*EsgMetricFrequencyOptions, bool) {
	if o == nil || IsNil(o.EsgMetricFrequencyOption) {
		return nil, false
	}
	return o.EsgMetricFrequencyOption, true
}

// HasEsgMetricFrequencyOption returns a boolean if a field has been set.
func (o *EsgMetricFrequencyOptionsPostRequest) HasEsgMetricFrequencyOption() bool {
	if o != nil && !IsNil(o.EsgMetricFrequencyOption) {
		return true
	}

	return false
}

// SetEsgMetricFrequencyOption gets a reference to the given EsgMetricFrequencyOptions and assigns it to the EsgMetricFrequencyOption field.
func (o *EsgMetricFrequencyOptionsPostRequest) SetEsgMetricFrequencyOption(v EsgMetricFrequencyOptions) {
	o.EsgMetricFrequencyOption = &v
}

func (o EsgMetricFrequencyOptionsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsgMetricFrequencyOptionsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EsgMetricFrequencyOption) {
		toSerialize["esg_metric _frequency _option"] = o.EsgMetricFrequencyOption
	}
	return toSerialize, nil
}

type NullableEsgMetricFrequencyOptionsPostRequest struct {
	value *EsgMetricFrequencyOptionsPostRequest
	isSet bool
}

func (v NullableEsgMetricFrequencyOptionsPostRequest) Get() *EsgMetricFrequencyOptionsPostRequest {
	return v.value
}

func (v *NullableEsgMetricFrequencyOptionsPostRequest) Set(val *EsgMetricFrequencyOptionsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEsgMetricFrequencyOptionsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEsgMetricFrequencyOptionsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsgMetricFrequencyOptionsPostRequest(val *EsgMetricFrequencyOptionsPostRequest) *NullableEsgMetricFrequencyOptionsPostRequest {
	return &NullableEsgMetricFrequencyOptionsPostRequest{value: val, isSet: true}
}

func (v NullableEsgMetricFrequencyOptionsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsgMetricFrequencyOptionsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

