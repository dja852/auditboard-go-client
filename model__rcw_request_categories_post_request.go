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

// checks if the RcwRequestCategoriesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RcwRequestCategoriesPostRequest{}

// RcwRequestCategoriesPostRequest struct for RcwRequestCategoriesPostRequest
type RcwRequestCategoriesPostRequest struct {
	RcwRequestCategory *RcwRequestCategories `json:"rcw_request _category,omitempty"`
}

// NewRcwRequestCategoriesPostRequest instantiates a new RcwRequestCategoriesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRcwRequestCategoriesPostRequest() *RcwRequestCategoriesPostRequest {
	this := RcwRequestCategoriesPostRequest{}
	return &this
}

// NewRcwRequestCategoriesPostRequestWithDefaults instantiates a new RcwRequestCategoriesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRcwRequestCategoriesPostRequestWithDefaults() *RcwRequestCategoriesPostRequest {
	this := RcwRequestCategoriesPostRequest{}
	return &this
}

// GetRcwRequestCategory returns the RcwRequestCategory field value if set, zero value otherwise.
func (o *RcwRequestCategoriesPostRequest) GetRcwRequestCategory() RcwRequestCategories {
	if o == nil || IsNil(o.RcwRequestCategory) {
		var ret RcwRequestCategories
		return ret
	}
	return *o.RcwRequestCategory
}

// GetRcwRequestCategoryOk returns a tuple with the RcwRequestCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwRequestCategoriesPostRequest) GetRcwRequestCategoryOk() (*RcwRequestCategories, bool) {
	if o == nil || IsNil(o.RcwRequestCategory) {
		return nil, false
	}
	return o.RcwRequestCategory, true
}

// HasRcwRequestCategory returns a boolean if a field has been set.
func (o *RcwRequestCategoriesPostRequest) HasRcwRequestCategory() bool {
	if o != nil && !IsNil(o.RcwRequestCategory) {
		return true
	}

	return false
}

// SetRcwRequestCategory gets a reference to the given RcwRequestCategories and assigns it to the RcwRequestCategory field.
func (o *RcwRequestCategoriesPostRequest) SetRcwRequestCategory(v RcwRequestCategories) {
	o.RcwRequestCategory = &v
}

func (o RcwRequestCategoriesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RcwRequestCategoriesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RcwRequestCategory) {
		toSerialize["rcw_request _category"] = o.RcwRequestCategory
	}
	return toSerialize, nil
}

type NullableRcwRequestCategoriesPostRequest struct {
	value *RcwRequestCategoriesPostRequest
	isSet bool
}

func (v NullableRcwRequestCategoriesPostRequest) Get() *RcwRequestCategoriesPostRequest {
	return v.value
}

func (v *NullableRcwRequestCategoriesPostRequest) Set(val *RcwRequestCategoriesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRcwRequestCategoriesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRcwRequestCategoriesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRcwRequestCategoriesPostRequest(val *RcwRequestCategoriesPostRequest) *NullableRcwRequestCategoriesPostRequest {
	return &NullableRcwRequestCategoriesPostRequest{value: val, isSet: true}
}

func (v NullableRcwRequestCategoriesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRcwRequestCategoriesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


