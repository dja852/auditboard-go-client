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

// checks if the EsgCategoriesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsgCategoriesPostRequest{}

// EsgCategoriesPostRequest struct for EsgCategoriesPostRequest
type EsgCategoriesPostRequest struct {
	EsgCategory *EsgCategories `json:"esg_category,omitempty"`
}

// NewEsgCategoriesPostRequest instantiates a new EsgCategoriesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsgCategoriesPostRequest() *EsgCategoriesPostRequest {
	this := EsgCategoriesPostRequest{}
	return &this
}

// NewEsgCategoriesPostRequestWithDefaults instantiates a new EsgCategoriesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsgCategoriesPostRequestWithDefaults() *EsgCategoriesPostRequest {
	this := EsgCategoriesPostRequest{}
	return &this
}

// GetEsgCategory returns the EsgCategory field value if set, zero value otherwise.
func (o *EsgCategoriesPostRequest) GetEsgCategory() EsgCategories {
	if o == nil || IsNil(o.EsgCategory) {
		var ret EsgCategories
		return ret
	}
	return *o.EsgCategory
}

// GetEsgCategoryOk returns a tuple with the EsgCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgCategoriesPostRequest) GetEsgCategoryOk() (*EsgCategories, bool) {
	if o == nil || IsNil(o.EsgCategory) {
		return nil, false
	}
	return o.EsgCategory, true
}

// HasEsgCategory returns a boolean if a field has been set.
func (o *EsgCategoriesPostRequest) HasEsgCategory() bool {
	if o != nil && !IsNil(o.EsgCategory) {
		return true
	}

	return false
}

// SetEsgCategory gets a reference to the given EsgCategories and assigns it to the EsgCategory field.
func (o *EsgCategoriesPostRequest) SetEsgCategory(v EsgCategories) {
	o.EsgCategory = &v
}

func (o EsgCategoriesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsgCategoriesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EsgCategory) {
		toSerialize["esg_category"] = o.EsgCategory
	}
	return toSerialize, nil
}

type NullableEsgCategoriesPostRequest struct {
	value *EsgCategoriesPostRequest
	isSet bool
}

func (v NullableEsgCategoriesPostRequest) Get() *EsgCategoriesPostRequest {
	return v.value
}

func (v *NullableEsgCategoriesPostRequest) Set(val *EsgCategoriesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEsgCategoriesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEsgCategoriesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsgCategoriesPostRequest(val *EsgCategoriesPostRequest) *NullableEsgCategoriesPostRequest {
	return &NullableEsgCategoriesPostRequest{value: val, isSet: true}
}

func (v NullableEsgCategoriesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsgCategoriesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

