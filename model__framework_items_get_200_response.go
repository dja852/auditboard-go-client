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

// checks if the FrameworkItemsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrameworkItemsGet200Response{}

// FrameworkItemsGet200Response struct for FrameworkItemsGet200Response
type FrameworkItemsGet200Response struct {
	FrameworkItems []FrameworkItems `json:"framework_items,omitempty"`
}

// NewFrameworkItemsGet200Response instantiates a new FrameworkItemsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameworkItemsGet200Response() *FrameworkItemsGet200Response {
	this := FrameworkItemsGet200Response{}
	return &this
}

// NewFrameworkItemsGet200ResponseWithDefaults instantiates a new FrameworkItemsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameworkItemsGet200ResponseWithDefaults() *FrameworkItemsGet200Response {
	this := FrameworkItemsGet200Response{}
	return &this
}

// GetFrameworkItems returns the FrameworkItems field value if set, zero value otherwise.
func (o *FrameworkItemsGet200Response) GetFrameworkItems() []FrameworkItems {
	if o == nil || IsNil(o.FrameworkItems) {
		var ret []FrameworkItems
		return ret
	}
	return o.FrameworkItems
}

// GetFrameworkItemsOk returns a tuple with the FrameworkItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameworkItemsGet200Response) GetFrameworkItemsOk() ([]FrameworkItems, bool) {
	if o == nil || IsNil(o.FrameworkItems) {
		return nil, false
	}
	return o.FrameworkItems, true
}

// HasFrameworkItems returns a boolean if a field has been set.
func (o *FrameworkItemsGet200Response) HasFrameworkItems() bool {
	if o != nil && !IsNil(o.FrameworkItems) {
		return true
	}

	return false
}

// SetFrameworkItems gets a reference to the given []FrameworkItems and assigns it to the FrameworkItems field.
func (o *FrameworkItemsGet200Response) SetFrameworkItems(v []FrameworkItems) {
	o.FrameworkItems = v
}

func (o FrameworkItemsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrameworkItemsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FrameworkItems) {
		toSerialize["framework_items"] = o.FrameworkItems
	}
	return toSerialize, nil
}

type NullableFrameworkItemsGet200Response struct {
	value *FrameworkItemsGet200Response
	isSet bool
}

func (v NullableFrameworkItemsGet200Response) Get() *FrameworkItemsGet200Response {
	return v.value
}

func (v *NullableFrameworkItemsGet200Response) Set(val *FrameworkItemsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameworkItemsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameworkItemsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameworkItemsGet200Response(val *FrameworkItemsGet200Response) *NullableFrameworkItemsGet200Response {
	return &NullableFrameworkItemsGet200Response{value: val, isSet: true}
}

func (v NullableFrameworkItemsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameworkItemsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


