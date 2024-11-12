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

// checks if the EsgMetricValuesGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsgMetricValuesGet200Response{}

// EsgMetricValuesGet200Response struct for EsgMetricValuesGet200Response
type EsgMetricValuesGet200Response struct {
	EsgMetricValues []EsgMetricValues `json:"esg_metric_values,omitempty"`
}

// NewEsgMetricValuesGet200Response instantiates a new EsgMetricValuesGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsgMetricValuesGet200Response() *EsgMetricValuesGet200Response {
	this := EsgMetricValuesGet200Response{}
	return &this
}

// NewEsgMetricValuesGet200ResponseWithDefaults instantiates a new EsgMetricValuesGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsgMetricValuesGet200ResponseWithDefaults() *EsgMetricValuesGet200Response {
	this := EsgMetricValuesGet200Response{}
	return &this
}

// GetEsgMetricValues returns the EsgMetricValues field value if set, zero value otherwise.
func (o *EsgMetricValuesGet200Response) GetEsgMetricValues() []EsgMetricValues {
	if o == nil || IsNil(o.EsgMetricValues) {
		var ret []EsgMetricValues
		return ret
	}
	return o.EsgMetricValues
}

// GetEsgMetricValuesOk returns a tuple with the EsgMetricValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricValuesGet200Response) GetEsgMetricValuesOk() ([]EsgMetricValues, bool) {
	if o == nil || IsNil(o.EsgMetricValues) {
		return nil, false
	}
	return o.EsgMetricValues, true
}

// HasEsgMetricValues returns a boolean if a field has been set.
func (o *EsgMetricValuesGet200Response) HasEsgMetricValues() bool {
	if o != nil && !IsNil(o.EsgMetricValues) {
		return true
	}

	return false
}

// SetEsgMetricValues gets a reference to the given []EsgMetricValues and assigns it to the EsgMetricValues field.
func (o *EsgMetricValuesGet200Response) SetEsgMetricValues(v []EsgMetricValues) {
	o.EsgMetricValues = v
}

func (o EsgMetricValuesGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsgMetricValuesGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EsgMetricValues) {
		toSerialize["esg_metric_values"] = o.EsgMetricValues
	}
	return toSerialize, nil
}

type NullableEsgMetricValuesGet200Response struct {
	value *EsgMetricValuesGet200Response
	isSet bool
}

func (v NullableEsgMetricValuesGet200Response) Get() *EsgMetricValuesGet200Response {
	return v.value
}

func (v *NullableEsgMetricValuesGet200Response) Set(val *EsgMetricValuesGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEsgMetricValuesGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEsgMetricValuesGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsgMetricValuesGet200Response(val *EsgMetricValuesGet200Response) *NullableEsgMetricValuesGet200Response {
	return &NullableEsgMetricValuesGet200Response{value: val, isSet: true}
}

func (v NullableEsgMetricValuesGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsgMetricValuesGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


