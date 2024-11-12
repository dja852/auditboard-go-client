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

// checks if the AuditRotationSchedulesPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditRotationSchedulesPut{}

// AuditRotationSchedulesPut struct for AuditRotationSchedulesPut
type AuditRotationSchedulesPut struct {
	AuditRotationSchedule *AuditOfficeLocationsPutAuditOfficeLocation `json:"audit_rotation_schedule,omitempty"`
	PreviousValues *AuditOfficeLocationsPutPreviousValues `json:"previous_values,omitempty"`
}

// NewAuditRotationSchedulesPut instantiates a new AuditRotationSchedulesPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditRotationSchedulesPut() *AuditRotationSchedulesPut {
	this := AuditRotationSchedulesPut{}
	return &this
}

// NewAuditRotationSchedulesPutWithDefaults instantiates a new AuditRotationSchedulesPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditRotationSchedulesPutWithDefaults() *AuditRotationSchedulesPut {
	this := AuditRotationSchedulesPut{}
	return &this
}

// GetAuditRotationSchedule returns the AuditRotationSchedule field value if set, zero value otherwise.
func (o *AuditRotationSchedulesPut) GetAuditRotationSchedule() AuditOfficeLocationsPutAuditOfficeLocation {
	if o == nil || IsNil(o.AuditRotationSchedule) {
		var ret AuditOfficeLocationsPutAuditOfficeLocation
		return ret
	}
	return *o.AuditRotationSchedule
}

// GetAuditRotationScheduleOk returns a tuple with the AuditRotationSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditRotationSchedulesPut) GetAuditRotationScheduleOk() (*AuditOfficeLocationsPutAuditOfficeLocation, bool) {
	if o == nil || IsNil(o.AuditRotationSchedule) {
		return nil, false
	}
	return o.AuditRotationSchedule, true
}

// HasAuditRotationSchedule returns a boolean if a field has been set.
func (o *AuditRotationSchedulesPut) HasAuditRotationSchedule() bool {
	if o != nil && !IsNil(o.AuditRotationSchedule) {
		return true
	}

	return false
}

// SetAuditRotationSchedule gets a reference to the given AuditOfficeLocationsPutAuditOfficeLocation and assigns it to the AuditRotationSchedule field.
func (o *AuditRotationSchedulesPut) SetAuditRotationSchedule(v AuditOfficeLocationsPutAuditOfficeLocation) {
	o.AuditRotationSchedule = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *AuditRotationSchedulesPut) GetPreviousValues() AuditOfficeLocationsPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret AuditOfficeLocationsPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditRotationSchedulesPut) GetPreviousValuesOk() (*AuditOfficeLocationsPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *AuditRotationSchedulesPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given AuditOfficeLocationsPutPreviousValues and assigns it to the PreviousValues field.
func (o *AuditRotationSchedulesPut) SetPreviousValues(v AuditOfficeLocationsPutPreviousValues) {
	o.PreviousValues = &v
}

func (o AuditRotationSchedulesPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditRotationSchedulesPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuditRotationSchedule) {
		toSerialize["audit_rotation_schedule"] = o.AuditRotationSchedule
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableAuditRotationSchedulesPut struct {
	value *AuditRotationSchedulesPut
	isSet bool
}

func (v NullableAuditRotationSchedulesPut) Get() *AuditRotationSchedulesPut {
	return v.value
}

func (v *NullableAuditRotationSchedulesPut) Set(val *AuditRotationSchedulesPut) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditRotationSchedulesPut) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditRotationSchedulesPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditRotationSchedulesPut(val *AuditRotationSchedulesPut) *NullableAuditRotationSchedulesPut {
	return &NullableAuditRotationSchedulesPut{value: val, isSet: true}
}

func (v NullableAuditRotationSchedulesPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditRotationSchedulesPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

