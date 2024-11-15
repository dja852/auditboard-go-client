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

// checks if the EsgMetricValuesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsgMetricValuesPostRequest{}

// EsgMetricValuesPostRequest struct for EsgMetricValuesPostRequest
type EsgMetricValuesPostRequest struct {
	EsgMetricValue *EsgMetricValues `json:"esg_metric _value,omitempty"`
}

// NewEsgMetricValuesPostRequest instantiates a new EsgMetricValuesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsgMetricValuesPostRequest() *EsgMetricValuesPostRequest {
	this := EsgMetricValuesPostRequest{}
	return &this
}

// NewEsgMetricValuesPostRequestWithDefaults instantiates a new EsgMetricValuesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsgMetricValuesPostRequestWithDefaults() *EsgMetricValuesPostRequest {
	this := EsgMetricValuesPostRequest{}
	return &this
}

// GetEsgMetricValue returns the EsgMetricValue field value if set, zero value otherwise.
func (o *EsgMetricValuesPostRequest) GetEsgMetricValue() EsgMetricValues {
	if o == nil || IsNil(o.EsgMetricValue) {
		var ret EsgMetricValues
		return ret
	}
	return *o.EsgMetricValue
}

// GetEsgMetricValueOk returns a tuple with the EsgMetricValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricValuesPostRequest) GetEsgMetricValueOk() (*EsgMetricValues, bool) {
	if o == nil || IsNil(o.EsgMetricValue) {
		return nil, false
	}
	return o.EsgMetricValue, true
}

// HasEsgMetricValue returns a boolean if a field has been set.
func (o *EsgMetricValuesPostRequest) HasEsgMetricValue() bool {
	if o != nil && !IsNil(o.EsgMetricValue) {
		return true
	}

	return false
}

// SetEsgMetricValue gets a reference to the given EsgMetricValues and assigns it to the EsgMetricValue field.
func (o *EsgMetricValuesPostRequest) SetEsgMetricValue(v EsgMetricValues) {
	o.EsgMetricValue = &v
}

func (o EsgMetricValuesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsgMetricValuesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EsgMetricValue) {
		toSerialize["esg_metric _value"] = o.EsgMetricValue
	}
	return toSerialize, nil
}

type NullableEsgMetricValuesPostRequest struct {
	value *EsgMetricValuesPostRequest
	isSet bool
}

func (v NullableEsgMetricValuesPostRequest) Get() *EsgMetricValuesPostRequest {
	return v.value
}

func (v *NullableEsgMetricValuesPostRequest) Set(val *EsgMetricValuesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEsgMetricValuesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEsgMetricValuesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsgMetricValuesPostRequest(val *EsgMetricValuesPostRequest) *NullableEsgMetricValuesPostRequest {
	return &NullableEsgMetricValuesPostRequest{value: val, isSet: true}
}

func (v NullableEsgMetricValuesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsgMetricValuesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


