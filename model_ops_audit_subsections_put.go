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

// checks if the OpsAuditSubsectionsPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpsAuditSubsectionsPut{}

// OpsAuditSubsectionsPut struct for OpsAuditSubsectionsPut
type OpsAuditSubsectionsPut struct {
	OpsAuditSubsection *OpsAuditSubsectionsPutOpsAuditSubsection `json:"ops_audit_subsection,omitempty"`
	PreviousValues *OpsAuditSubsectionsPutPreviousValues `json:"previous_values,omitempty"`
}

// NewOpsAuditSubsectionsPut instantiates a new OpsAuditSubsectionsPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsAuditSubsectionsPut() *OpsAuditSubsectionsPut {
	this := OpsAuditSubsectionsPut{}
	return &this
}

// NewOpsAuditSubsectionsPutWithDefaults instantiates a new OpsAuditSubsectionsPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsAuditSubsectionsPutWithDefaults() *OpsAuditSubsectionsPut {
	this := OpsAuditSubsectionsPut{}
	return &this
}

// GetOpsAuditSubsection returns the OpsAuditSubsection field value if set, zero value otherwise.
func (o *OpsAuditSubsectionsPut) GetOpsAuditSubsection() OpsAuditSubsectionsPutOpsAuditSubsection {
	if o == nil || IsNil(o.OpsAuditSubsection) {
		var ret OpsAuditSubsectionsPutOpsAuditSubsection
		return ret
	}
	return *o.OpsAuditSubsection
}

// GetOpsAuditSubsectionOk returns a tuple with the OpsAuditSubsection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditSubsectionsPut) GetOpsAuditSubsectionOk() (*OpsAuditSubsectionsPutOpsAuditSubsection, bool) {
	if o == nil || IsNil(o.OpsAuditSubsection) {
		return nil, false
	}
	return o.OpsAuditSubsection, true
}

// HasOpsAuditSubsection returns a boolean if a field has been set.
func (o *OpsAuditSubsectionsPut) HasOpsAuditSubsection() bool {
	if o != nil && !IsNil(o.OpsAuditSubsection) {
		return true
	}

	return false
}

// SetOpsAuditSubsection gets a reference to the given OpsAuditSubsectionsPutOpsAuditSubsection and assigns it to the OpsAuditSubsection field.
func (o *OpsAuditSubsectionsPut) SetOpsAuditSubsection(v OpsAuditSubsectionsPutOpsAuditSubsection) {
	o.OpsAuditSubsection = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *OpsAuditSubsectionsPut) GetPreviousValues() OpsAuditSubsectionsPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret OpsAuditSubsectionsPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditSubsectionsPut) GetPreviousValuesOk() (*OpsAuditSubsectionsPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *OpsAuditSubsectionsPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given OpsAuditSubsectionsPutPreviousValues and assigns it to the PreviousValues field.
func (o *OpsAuditSubsectionsPut) SetPreviousValues(v OpsAuditSubsectionsPutPreviousValues) {
	o.PreviousValues = &v
}

func (o OpsAuditSubsectionsPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpsAuditSubsectionsPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OpsAuditSubsection) {
		toSerialize["ops_audit_subsection"] = o.OpsAuditSubsection
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableOpsAuditSubsectionsPut struct {
	value *OpsAuditSubsectionsPut
	isSet bool
}

func (v NullableOpsAuditSubsectionsPut) Get() *OpsAuditSubsectionsPut {
	return v.value
}

func (v *NullableOpsAuditSubsectionsPut) Set(val *OpsAuditSubsectionsPut) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsAuditSubsectionsPut) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsAuditSubsectionsPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsAuditSubsectionsPut(val *OpsAuditSubsectionsPut) *NullableOpsAuditSubsectionsPut {
	return &NullableOpsAuditSubsectionsPut{value: val, isSet: true}
}

func (v NullableOpsAuditSubsectionsPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsAuditSubsectionsPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


