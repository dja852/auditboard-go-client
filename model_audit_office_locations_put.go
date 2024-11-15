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

// checks if the AuditOfficeLocationsPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditOfficeLocationsPut{}

// AuditOfficeLocationsPut struct for AuditOfficeLocationsPut
type AuditOfficeLocationsPut struct {
	AuditOfficeLocation *AuditOfficeLocationsPutAuditOfficeLocation `json:"audit_office_location,omitempty"`
	PreviousValues *AuditOfficeLocationsPutPreviousValues `json:"previous_values,omitempty"`
}

// NewAuditOfficeLocationsPut instantiates a new AuditOfficeLocationsPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditOfficeLocationsPut() *AuditOfficeLocationsPut {
	this := AuditOfficeLocationsPut{}
	return &this
}

// NewAuditOfficeLocationsPutWithDefaults instantiates a new AuditOfficeLocationsPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditOfficeLocationsPutWithDefaults() *AuditOfficeLocationsPut {
	this := AuditOfficeLocationsPut{}
	return &this
}

// GetAuditOfficeLocation returns the AuditOfficeLocation field value if set, zero value otherwise.
func (o *AuditOfficeLocationsPut) GetAuditOfficeLocation() AuditOfficeLocationsPutAuditOfficeLocation {
	if o == nil || IsNil(o.AuditOfficeLocation) {
		var ret AuditOfficeLocationsPutAuditOfficeLocation
		return ret
	}
	return *o.AuditOfficeLocation
}

// GetAuditOfficeLocationOk returns a tuple with the AuditOfficeLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditOfficeLocationsPut) GetAuditOfficeLocationOk() (*AuditOfficeLocationsPutAuditOfficeLocation, bool) {
	if o == nil || IsNil(o.AuditOfficeLocation) {
		return nil, false
	}
	return o.AuditOfficeLocation, true
}

// HasAuditOfficeLocation returns a boolean if a field has been set.
func (o *AuditOfficeLocationsPut) HasAuditOfficeLocation() bool {
	if o != nil && !IsNil(o.AuditOfficeLocation) {
		return true
	}

	return false
}

// SetAuditOfficeLocation gets a reference to the given AuditOfficeLocationsPutAuditOfficeLocation and assigns it to the AuditOfficeLocation field.
func (o *AuditOfficeLocationsPut) SetAuditOfficeLocation(v AuditOfficeLocationsPutAuditOfficeLocation) {
	o.AuditOfficeLocation = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *AuditOfficeLocationsPut) GetPreviousValues() AuditOfficeLocationsPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret AuditOfficeLocationsPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditOfficeLocationsPut) GetPreviousValuesOk() (*AuditOfficeLocationsPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *AuditOfficeLocationsPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given AuditOfficeLocationsPutPreviousValues and assigns it to the PreviousValues field.
func (o *AuditOfficeLocationsPut) SetPreviousValues(v AuditOfficeLocationsPutPreviousValues) {
	o.PreviousValues = &v
}

func (o AuditOfficeLocationsPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditOfficeLocationsPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuditOfficeLocation) {
		toSerialize["audit_office_location"] = o.AuditOfficeLocation
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableAuditOfficeLocationsPut struct {
	value *AuditOfficeLocationsPut
	isSet bool
}

func (v NullableAuditOfficeLocationsPut) Get() *AuditOfficeLocationsPut {
	return v.value
}

func (v *NullableAuditOfficeLocationsPut) Set(val *AuditOfficeLocationsPut) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditOfficeLocationsPut) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditOfficeLocationsPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditOfficeLocationsPut(val *AuditOfficeLocationsPut) *NullableAuditOfficeLocationsPut {
	return &NullableAuditOfficeLocationsPut{value: val, isSet: true}
}

func (v NullableAuditOfficeLocationsPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditOfficeLocationsPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


