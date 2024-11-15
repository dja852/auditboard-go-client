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

// checks if the BusinessUnitsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BusinessUnitsGet200Response{}

// BusinessUnitsGet200Response struct for BusinessUnitsGet200Response
type BusinessUnitsGet200Response struct {
	BusinessUnits []BusinessUnits `json:"business_units,omitempty"`
}

// NewBusinessUnitsGet200Response instantiates a new BusinessUnitsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessUnitsGet200Response() *BusinessUnitsGet200Response {
	this := BusinessUnitsGet200Response{}
	return &this
}

// NewBusinessUnitsGet200ResponseWithDefaults instantiates a new BusinessUnitsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessUnitsGet200ResponseWithDefaults() *BusinessUnitsGet200Response {
	this := BusinessUnitsGet200Response{}
	return &this
}

// GetBusinessUnits returns the BusinessUnits field value if set, zero value otherwise.
func (o *BusinessUnitsGet200Response) GetBusinessUnits() []BusinessUnits {
	if o == nil || IsNil(o.BusinessUnits) {
		var ret []BusinessUnits
		return ret
	}
	return o.BusinessUnits
}

// GetBusinessUnitsOk returns a tuple with the BusinessUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessUnitsGet200Response) GetBusinessUnitsOk() ([]BusinessUnits, bool) {
	if o == nil || IsNil(o.BusinessUnits) {
		return nil, false
	}
	return o.BusinessUnits, true
}

// HasBusinessUnits returns a boolean if a field has been set.
func (o *BusinessUnitsGet200Response) HasBusinessUnits() bool {
	if o != nil && !IsNil(o.BusinessUnits) {
		return true
	}

	return false
}

// SetBusinessUnits gets a reference to the given []BusinessUnits and assigns it to the BusinessUnits field.
func (o *BusinessUnitsGet200Response) SetBusinessUnits(v []BusinessUnits) {
	o.BusinessUnits = v
}

func (o BusinessUnitsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BusinessUnitsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessUnits) {
		toSerialize["business_units"] = o.BusinessUnits
	}
	return toSerialize, nil
}

type NullableBusinessUnitsGet200Response struct {
	value *BusinessUnitsGet200Response
	isSet bool
}

func (v NullableBusinessUnitsGet200Response) Get() *BusinessUnitsGet200Response {
	return v.value
}

func (v *NullableBusinessUnitsGet200Response) Set(val *BusinessUnitsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessUnitsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessUnitsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessUnitsGet200Response(val *BusinessUnitsGet200Response) *NullableBusinessUnitsGet200Response {
	return &NullableBusinessUnitsGet200Response{value: val, isSet: true}
}

func (v NullableBusinessUnitsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessUnitsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


