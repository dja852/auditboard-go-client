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

// checks if the EsgMetricCategoriesGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsgMetricCategoriesGet200Response{}

// EsgMetricCategoriesGet200Response struct for EsgMetricCategoriesGet200Response
type EsgMetricCategoriesGet200Response struct {
	EsgMetricCategories []EsgMetricCategories `json:"esg_metric_categories,omitempty"`
}

// NewEsgMetricCategoriesGet200Response instantiates a new EsgMetricCategoriesGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsgMetricCategoriesGet200Response() *EsgMetricCategoriesGet200Response {
	this := EsgMetricCategoriesGet200Response{}
	return &this
}

// NewEsgMetricCategoriesGet200ResponseWithDefaults instantiates a new EsgMetricCategoriesGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsgMetricCategoriesGet200ResponseWithDefaults() *EsgMetricCategoriesGet200Response {
	this := EsgMetricCategoriesGet200Response{}
	return &this
}

// GetEsgMetricCategories returns the EsgMetricCategories field value if set, zero value otherwise.
func (o *EsgMetricCategoriesGet200Response) GetEsgMetricCategories() []EsgMetricCategories {
	if o == nil || IsNil(o.EsgMetricCategories) {
		var ret []EsgMetricCategories
		return ret
	}
	return o.EsgMetricCategories
}

// GetEsgMetricCategoriesOk returns a tuple with the EsgMetricCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricCategoriesGet200Response) GetEsgMetricCategoriesOk() ([]EsgMetricCategories, bool) {
	if o == nil || IsNil(o.EsgMetricCategories) {
		return nil, false
	}
	return o.EsgMetricCategories, true
}

// HasEsgMetricCategories returns a boolean if a field has been set.
func (o *EsgMetricCategoriesGet200Response) HasEsgMetricCategories() bool {
	if o != nil && !IsNil(o.EsgMetricCategories) {
		return true
	}

	return false
}

// SetEsgMetricCategories gets a reference to the given []EsgMetricCategories and assigns it to the EsgMetricCategories field.
func (o *EsgMetricCategoriesGet200Response) SetEsgMetricCategories(v []EsgMetricCategories) {
	o.EsgMetricCategories = v
}

func (o EsgMetricCategoriesGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsgMetricCategoriesGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EsgMetricCategories) {
		toSerialize["esg_metric_categories"] = o.EsgMetricCategories
	}
	return toSerialize, nil
}

type NullableEsgMetricCategoriesGet200Response struct {
	value *EsgMetricCategoriesGet200Response
	isSet bool
}

func (v NullableEsgMetricCategoriesGet200Response) Get() *EsgMetricCategoriesGet200Response {
	return v.value
}

func (v *NullableEsgMetricCategoriesGet200Response) Set(val *EsgMetricCategoriesGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEsgMetricCategoriesGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEsgMetricCategoriesGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsgMetricCategoriesGet200Response(val *EsgMetricCategoriesGet200Response) *NullableEsgMetricCategoriesGet200Response {
	return &NullableEsgMetricCategoriesGet200Response{value: val, isSet: true}
}

func (v NullableEsgMetricCategoriesGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsgMetricCategoriesGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

