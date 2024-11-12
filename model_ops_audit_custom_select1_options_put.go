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

// checks if the OpsAuditCustomSelect1OptionsPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpsAuditCustomSelect1OptionsPut{}

// OpsAuditCustomSelect1OptionsPut struct for OpsAuditCustomSelect1OptionsPut
type OpsAuditCustomSelect1OptionsPut struct {
	OpsAuditCustomSelect1Option *AssessmentPeriodsPutAssessmentPeriod `json:"ops_audit_custom_select1_option,omitempty"`
	PreviousValues *AssessmentPeriodsPutPreviousValues `json:"previous_values,omitempty"`
}

// NewOpsAuditCustomSelect1OptionsPut instantiates a new OpsAuditCustomSelect1OptionsPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsAuditCustomSelect1OptionsPut() *OpsAuditCustomSelect1OptionsPut {
	this := OpsAuditCustomSelect1OptionsPut{}
	return &this
}

// NewOpsAuditCustomSelect1OptionsPutWithDefaults instantiates a new OpsAuditCustomSelect1OptionsPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsAuditCustomSelect1OptionsPutWithDefaults() *OpsAuditCustomSelect1OptionsPut {
	this := OpsAuditCustomSelect1OptionsPut{}
	return &this
}

// GetOpsAuditCustomSelect1Option returns the OpsAuditCustomSelect1Option field value if set, zero value otherwise.
func (o *OpsAuditCustomSelect1OptionsPut) GetOpsAuditCustomSelect1Option() AssessmentPeriodsPutAssessmentPeriod {
	if o == nil || IsNil(o.OpsAuditCustomSelect1Option) {
		var ret AssessmentPeriodsPutAssessmentPeriod
		return ret
	}
	return *o.OpsAuditCustomSelect1Option
}

// GetOpsAuditCustomSelect1OptionOk returns a tuple with the OpsAuditCustomSelect1Option field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditCustomSelect1OptionsPut) GetOpsAuditCustomSelect1OptionOk() (*AssessmentPeriodsPutAssessmentPeriod, bool) {
	if o == nil || IsNil(o.OpsAuditCustomSelect1Option) {
		return nil, false
	}
	return o.OpsAuditCustomSelect1Option, true
}

// HasOpsAuditCustomSelect1Option returns a boolean if a field has been set.
func (o *OpsAuditCustomSelect1OptionsPut) HasOpsAuditCustomSelect1Option() bool {
	if o != nil && !IsNil(o.OpsAuditCustomSelect1Option) {
		return true
	}

	return false
}

// SetOpsAuditCustomSelect1Option gets a reference to the given AssessmentPeriodsPutAssessmentPeriod and assigns it to the OpsAuditCustomSelect1Option field.
func (o *OpsAuditCustomSelect1OptionsPut) SetOpsAuditCustomSelect1Option(v AssessmentPeriodsPutAssessmentPeriod) {
	o.OpsAuditCustomSelect1Option = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *OpsAuditCustomSelect1OptionsPut) GetPreviousValues() AssessmentPeriodsPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret AssessmentPeriodsPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditCustomSelect1OptionsPut) GetPreviousValuesOk() (*AssessmentPeriodsPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *OpsAuditCustomSelect1OptionsPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given AssessmentPeriodsPutPreviousValues and assigns it to the PreviousValues field.
func (o *OpsAuditCustomSelect1OptionsPut) SetPreviousValues(v AssessmentPeriodsPutPreviousValues) {
	o.PreviousValues = &v
}

func (o OpsAuditCustomSelect1OptionsPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpsAuditCustomSelect1OptionsPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OpsAuditCustomSelect1Option) {
		toSerialize["ops_audit_custom_select1_option"] = o.OpsAuditCustomSelect1Option
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableOpsAuditCustomSelect1OptionsPut struct {
	value *OpsAuditCustomSelect1OptionsPut
	isSet bool
}

func (v NullableOpsAuditCustomSelect1OptionsPut) Get() *OpsAuditCustomSelect1OptionsPut {
	return v.value
}

func (v *NullableOpsAuditCustomSelect1OptionsPut) Set(val *OpsAuditCustomSelect1OptionsPut) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsAuditCustomSelect1OptionsPut) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsAuditCustomSelect1OptionsPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsAuditCustomSelect1OptionsPut(val *OpsAuditCustomSelect1OptionsPut) *NullableOpsAuditCustomSelect1OptionsPut {
	return &NullableOpsAuditCustomSelect1OptionsPut{value: val, isSet: true}
}

func (v NullableOpsAuditCustomSelect1OptionsPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsAuditCustomSelect1OptionsPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

