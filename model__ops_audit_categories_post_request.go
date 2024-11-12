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

// checks if the OpsAuditCategoriesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpsAuditCategoriesPostRequest{}

// OpsAuditCategoriesPostRequest struct for OpsAuditCategoriesPostRequest
type OpsAuditCategoriesPostRequest struct {
	OpsAuditCategory *OpsAuditCategories `json:"ops_audit _category,omitempty"`
}

// NewOpsAuditCategoriesPostRequest instantiates a new OpsAuditCategoriesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsAuditCategoriesPostRequest() *OpsAuditCategoriesPostRequest {
	this := OpsAuditCategoriesPostRequest{}
	return &this
}

// NewOpsAuditCategoriesPostRequestWithDefaults instantiates a new OpsAuditCategoriesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsAuditCategoriesPostRequestWithDefaults() *OpsAuditCategoriesPostRequest {
	this := OpsAuditCategoriesPostRequest{}
	return &this
}

// GetOpsAuditCategory returns the OpsAuditCategory field value if set, zero value otherwise.
func (o *OpsAuditCategoriesPostRequest) GetOpsAuditCategory() OpsAuditCategories {
	if o == nil || IsNil(o.OpsAuditCategory) {
		var ret OpsAuditCategories
		return ret
	}
	return *o.OpsAuditCategory
}

// GetOpsAuditCategoryOk returns a tuple with the OpsAuditCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditCategoriesPostRequest) GetOpsAuditCategoryOk() (*OpsAuditCategories, bool) {
	if o == nil || IsNil(o.OpsAuditCategory) {
		return nil, false
	}
	return o.OpsAuditCategory, true
}

// HasOpsAuditCategory returns a boolean if a field has been set.
func (o *OpsAuditCategoriesPostRequest) HasOpsAuditCategory() bool {
	if o != nil && !IsNil(o.OpsAuditCategory) {
		return true
	}

	return false
}

// SetOpsAuditCategory gets a reference to the given OpsAuditCategories and assigns it to the OpsAuditCategory field.
func (o *OpsAuditCategoriesPostRequest) SetOpsAuditCategory(v OpsAuditCategories) {
	o.OpsAuditCategory = &v
}

func (o OpsAuditCategoriesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpsAuditCategoriesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OpsAuditCategory) {
		toSerialize["ops_audit _category"] = o.OpsAuditCategory
	}
	return toSerialize, nil
}

type NullableOpsAuditCategoriesPostRequest struct {
	value *OpsAuditCategoriesPostRequest
	isSet bool
}

func (v NullableOpsAuditCategoriesPostRequest) Get() *OpsAuditCategoriesPostRequest {
	return v.value
}

func (v *NullableOpsAuditCategoriesPostRequest) Set(val *OpsAuditCategoriesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsAuditCategoriesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsAuditCategoriesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsAuditCategoriesPostRequest(val *OpsAuditCategoriesPostRequest) *NullableOpsAuditCategoriesPostRequest {
	return &NullableOpsAuditCategoriesPostRequest{value: val, isSet: true}
}

func (v NullableOpsAuditCategoriesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsAuditCategoriesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

