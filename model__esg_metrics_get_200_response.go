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

// checks if the EsgMetricsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsgMetricsGet200Response{}

// EsgMetricsGet200Response struct for EsgMetricsGet200Response
type EsgMetricsGet200Response struct {
	EsgMetrics []EsgMetrics `json:"esg_metrics,omitempty"`
}

// NewEsgMetricsGet200Response instantiates a new EsgMetricsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsgMetricsGet200Response() *EsgMetricsGet200Response {
	this := EsgMetricsGet200Response{}
	return &this
}

// NewEsgMetricsGet200ResponseWithDefaults instantiates a new EsgMetricsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsgMetricsGet200ResponseWithDefaults() *EsgMetricsGet200Response {
	this := EsgMetricsGet200Response{}
	return &this
}

// GetEsgMetrics returns the EsgMetrics field value if set, zero value otherwise.
func (o *EsgMetricsGet200Response) GetEsgMetrics() []EsgMetrics {
	if o == nil || IsNil(o.EsgMetrics) {
		var ret []EsgMetrics
		return ret
	}
	return o.EsgMetrics
}

// GetEsgMetricsOk returns a tuple with the EsgMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricsGet200Response) GetEsgMetricsOk() ([]EsgMetrics, bool) {
	if o == nil || IsNil(o.EsgMetrics) {
		return nil, false
	}
	return o.EsgMetrics, true
}

// HasEsgMetrics returns a boolean if a field has been set.
func (o *EsgMetricsGet200Response) HasEsgMetrics() bool {
	if o != nil && !IsNil(o.EsgMetrics) {
		return true
	}

	return false
}

// SetEsgMetrics gets a reference to the given []EsgMetrics and assigns it to the EsgMetrics field.
func (o *EsgMetricsGet200Response) SetEsgMetrics(v []EsgMetrics) {
	o.EsgMetrics = v
}

func (o EsgMetricsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsgMetricsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EsgMetrics) {
		toSerialize["esg_metrics"] = o.EsgMetrics
	}
	return toSerialize, nil
}

type NullableEsgMetricsGet200Response struct {
	value *EsgMetricsGet200Response
	isSet bool
}

func (v NullableEsgMetricsGet200Response) Get() *EsgMetricsGet200Response {
	return v.value
}

func (v *NullableEsgMetricsGet200Response) Set(val *EsgMetricsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEsgMetricsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEsgMetricsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsgMetricsGet200Response(val *EsgMetricsGet200Response) *NullableEsgMetricsGet200Response {
	return &NullableEsgMetricsGet200Response{value: val, isSet: true}
}

func (v NullableEsgMetricsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsgMetricsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


