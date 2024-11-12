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

// checks if the CustomFieldsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFieldsGet200Response{}

// CustomFieldsGet200Response struct for CustomFieldsGet200Response
type CustomFieldsGet200Response struct {
	CustomFields []CustomFields `json:"custom_fields,omitempty"`
}

// NewCustomFieldsGet200Response instantiates a new CustomFieldsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFieldsGet200Response() *CustomFieldsGet200Response {
	this := CustomFieldsGet200Response{}
	return &this
}

// NewCustomFieldsGet200ResponseWithDefaults instantiates a new CustomFieldsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFieldsGet200ResponseWithDefaults() *CustomFieldsGet200Response {
	this := CustomFieldsGet200Response{}
	return &this
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *CustomFieldsGet200Response) GetCustomFields() []CustomFields {
	if o == nil || IsNil(o.CustomFields) {
		var ret []CustomFields
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldsGet200Response) GetCustomFieldsOk() ([]CustomFields, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *CustomFieldsGet200Response) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []CustomFields and assigns it to the CustomFields field.
func (o *CustomFieldsGet200Response) SetCustomFields(v []CustomFields) {
	o.CustomFields = v
}

func (o CustomFieldsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFieldsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableCustomFieldsGet200Response struct {
	value *CustomFieldsGet200Response
	isSet bool
}

func (v NullableCustomFieldsGet200Response) Get() *CustomFieldsGet200Response {
	return v.value
}

func (v *NullableCustomFieldsGet200Response) Set(val *CustomFieldsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldsGet200Response(val *CustomFieldsGet200Response) *NullableCustomFieldsGet200Response {
	return &NullableCustomFieldsGet200Response{value: val, isSet: true}
}

func (v NullableCustomFieldsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

