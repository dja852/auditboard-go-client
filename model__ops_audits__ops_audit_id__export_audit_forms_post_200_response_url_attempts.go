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

// checks if the OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts{}

// OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts struct for OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts
type OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts struct {
	Made *int32 `json:"made,omitempty"`
	Remaining *int32 `json:"remaining,omitempty"`
	Max *int32 `json:"max,omitempty"`
}

// NewOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts instantiates a new OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts() *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts {
	this := OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts{}
	return &this
}

// NewOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttemptsWithDefaults instantiates a new OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttemptsWithDefaults() *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts {
	this := OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts{}
	return &this
}

// GetMade returns the Made field value if set, zero value otherwise.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) GetMade() int32 {
	if o == nil || IsNil(o.Made) {
		var ret int32
		return ret
	}
	return *o.Made
}

// GetMadeOk returns a tuple with the Made field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) GetMadeOk() (*int32, bool) {
	if o == nil || IsNil(o.Made) {
		return nil, false
	}
	return o.Made, true
}

// HasMade returns a boolean if a field has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) HasMade() bool {
	if o != nil && !IsNil(o.Made) {
		return true
	}

	return false
}

// SetMade gets a reference to the given int32 and assigns it to the Made field.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) SetMade(v int32) {
	o.Made = &v
}

// GetRemaining returns the Remaining field value if set, zero value otherwise.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) GetRemaining() int32 {
	if o == nil || IsNil(o.Remaining) {
		var ret int32
		return ret
	}
	return *o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) GetRemainingOk() (*int32, bool) {
	if o == nil || IsNil(o.Remaining) {
		return nil, false
	}
	return o.Remaining, true
}

// HasRemaining returns a boolean if a field has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) HasRemaining() bool {
	if o != nil && !IsNil(o.Remaining) {
		return true
	}

	return false
}

// SetRemaining gets a reference to the given int32 and assigns it to the Remaining field.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) SetRemaining(v int32) {
	o.Remaining = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) SetMax(v int32) {
	o.Max = &v
}

func (o OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Made) {
		toSerialize["made"] = o.Made
	}
	if !IsNil(o.Remaining) {
		toSerialize["remaining"] = o.Remaining
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullableOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts struct {
	value *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts
	isSet bool
}

func (v NullableOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) Get() *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts {
	return v.value
}

func (v *NullableOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) Set(val *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts(val *OpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) *NullableOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts {
	return &NullableOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts{value: val, isSet: true}
}

func (v NullableOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsAuditsOpsAuditIdExportAuditFormsPost200ResponseUrlAttempts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

