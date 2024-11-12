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

// checks if the BusinessUnitsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BusinessUnitsPostRequest{}

// BusinessUnitsPostRequest struct for BusinessUnitsPostRequest
type BusinessUnitsPostRequest struct {
	BusinessUnit *BusinessUnits `json:"business_unit,omitempty"`
}

// NewBusinessUnitsPostRequest instantiates a new BusinessUnitsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessUnitsPostRequest() *BusinessUnitsPostRequest {
	this := BusinessUnitsPostRequest{}
	return &this
}

// NewBusinessUnitsPostRequestWithDefaults instantiates a new BusinessUnitsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessUnitsPostRequestWithDefaults() *BusinessUnitsPostRequest {
	this := BusinessUnitsPostRequest{}
	return &this
}

// GetBusinessUnit returns the BusinessUnit field value if set, zero value otherwise.
func (o *BusinessUnitsPostRequest) GetBusinessUnit() BusinessUnits {
	if o == nil || IsNil(o.BusinessUnit) {
		var ret BusinessUnits
		return ret
	}
	return *o.BusinessUnit
}

// GetBusinessUnitOk returns a tuple with the BusinessUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessUnitsPostRequest) GetBusinessUnitOk() (*BusinessUnits, bool) {
	if o == nil || IsNil(o.BusinessUnit) {
		return nil, false
	}
	return o.BusinessUnit, true
}

// HasBusinessUnit returns a boolean if a field has been set.
func (o *BusinessUnitsPostRequest) HasBusinessUnit() bool {
	if o != nil && !IsNil(o.BusinessUnit) {
		return true
	}

	return false
}

// SetBusinessUnit gets a reference to the given BusinessUnits and assigns it to the BusinessUnit field.
func (o *BusinessUnitsPostRequest) SetBusinessUnit(v BusinessUnits) {
	o.BusinessUnit = &v
}

func (o BusinessUnitsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BusinessUnitsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessUnit) {
		toSerialize["business_unit"] = o.BusinessUnit
	}
	return toSerialize, nil
}

type NullableBusinessUnitsPostRequest struct {
	value *BusinessUnitsPostRequest
	isSet bool
}

func (v NullableBusinessUnitsPostRequest) Get() *BusinessUnitsPostRequest {
	return v.value
}

func (v *NullableBusinessUnitsPostRequest) Set(val *BusinessUnitsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessUnitsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessUnitsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessUnitsPostRequest(val *BusinessUnitsPostRequest) *NullableBusinessUnitsPostRequest {
	return &NullableBusinessUnitsPostRequest{value: val, isSet: true}
}

func (v NullableBusinessUnitsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessUnitsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

