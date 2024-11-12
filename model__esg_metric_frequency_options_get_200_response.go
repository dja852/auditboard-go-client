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

// checks if the EsgMetricFrequencyOptionsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsgMetricFrequencyOptionsGet200Response{}

// EsgMetricFrequencyOptionsGet200Response struct for EsgMetricFrequencyOptionsGet200Response
type EsgMetricFrequencyOptionsGet200Response struct {
	EsgMetricFrequencyOptions []EsgMetricFrequencyOptions `json:"esg_metric_frequency_options,omitempty"`
}

// NewEsgMetricFrequencyOptionsGet200Response instantiates a new EsgMetricFrequencyOptionsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsgMetricFrequencyOptionsGet200Response() *EsgMetricFrequencyOptionsGet200Response {
	this := EsgMetricFrequencyOptionsGet200Response{}
	return &this
}

// NewEsgMetricFrequencyOptionsGet200ResponseWithDefaults instantiates a new EsgMetricFrequencyOptionsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsgMetricFrequencyOptionsGet200ResponseWithDefaults() *EsgMetricFrequencyOptionsGet200Response {
	this := EsgMetricFrequencyOptionsGet200Response{}
	return &this
}

// GetEsgMetricFrequencyOptions returns the EsgMetricFrequencyOptions field value if set, zero value otherwise.
func (o *EsgMetricFrequencyOptionsGet200Response) GetEsgMetricFrequencyOptions() []EsgMetricFrequencyOptions {
	if o == nil || IsNil(o.EsgMetricFrequencyOptions) {
		var ret []EsgMetricFrequencyOptions
		return ret
	}
	return o.EsgMetricFrequencyOptions
}

// GetEsgMetricFrequencyOptionsOk returns a tuple with the EsgMetricFrequencyOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricFrequencyOptionsGet200Response) GetEsgMetricFrequencyOptionsOk() ([]EsgMetricFrequencyOptions, bool) {
	if o == nil || IsNil(o.EsgMetricFrequencyOptions) {
		return nil, false
	}
	return o.EsgMetricFrequencyOptions, true
}

// HasEsgMetricFrequencyOptions returns a boolean if a field has been set.
func (o *EsgMetricFrequencyOptionsGet200Response) HasEsgMetricFrequencyOptions() bool {
	if o != nil && !IsNil(o.EsgMetricFrequencyOptions) {
		return true
	}

	return false
}

// SetEsgMetricFrequencyOptions gets a reference to the given []EsgMetricFrequencyOptions and assigns it to the EsgMetricFrequencyOptions field.
func (o *EsgMetricFrequencyOptionsGet200Response) SetEsgMetricFrequencyOptions(v []EsgMetricFrequencyOptions) {
	o.EsgMetricFrequencyOptions = v
}

func (o EsgMetricFrequencyOptionsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsgMetricFrequencyOptionsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EsgMetricFrequencyOptions) {
		toSerialize["esg_metric_frequency_options"] = o.EsgMetricFrequencyOptions
	}
	return toSerialize, nil
}

type NullableEsgMetricFrequencyOptionsGet200Response struct {
	value *EsgMetricFrequencyOptionsGet200Response
	isSet bool
}

func (v NullableEsgMetricFrequencyOptionsGet200Response) Get() *EsgMetricFrequencyOptionsGet200Response {
	return v.value
}

func (v *NullableEsgMetricFrequencyOptionsGet200Response) Set(val *EsgMetricFrequencyOptionsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEsgMetricFrequencyOptionsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEsgMetricFrequencyOptionsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsgMetricFrequencyOptionsGet200Response(val *EsgMetricFrequencyOptionsGet200Response) *NullableEsgMetricFrequencyOptionsGet200Response {
	return &NullableEsgMetricFrequencyOptionsGet200Response{value: val, isSet: true}
}

func (v NullableEsgMetricFrequencyOptionsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsgMetricFrequencyOptionsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

