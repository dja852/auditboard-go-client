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

// checks if the FinancialAccountsPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialAccountsPut{}

// FinancialAccountsPut struct for FinancialAccountsPut
type FinancialAccountsPut struct {
	FinancialAccount *FinancialAccountsPutFinancialAccount `json:"financial_account,omitempty"`
	PreviousValues *FinancialAccountsPutPreviousValues `json:"previous_values,omitempty"`
}

// NewFinancialAccountsPut instantiates a new FinancialAccountsPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialAccountsPut() *FinancialAccountsPut {
	this := FinancialAccountsPut{}
	return &this
}

// NewFinancialAccountsPutWithDefaults instantiates a new FinancialAccountsPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialAccountsPutWithDefaults() *FinancialAccountsPut {
	this := FinancialAccountsPut{}
	return &this
}

// GetFinancialAccount returns the FinancialAccount field value if set, zero value otherwise.
func (o *FinancialAccountsPut) GetFinancialAccount() FinancialAccountsPutFinancialAccount {
	if o == nil || IsNil(o.FinancialAccount) {
		var ret FinancialAccountsPutFinancialAccount
		return ret
	}
	return *o.FinancialAccount
}

// GetFinancialAccountOk returns a tuple with the FinancialAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialAccountsPut) GetFinancialAccountOk() (*FinancialAccountsPutFinancialAccount, bool) {
	if o == nil || IsNil(o.FinancialAccount) {
		return nil, false
	}
	return o.FinancialAccount, true
}

// HasFinancialAccount returns a boolean if a field has been set.
func (o *FinancialAccountsPut) HasFinancialAccount() bool {
	if o != nil && !IsNil(o.FinancialAccount) {
		return true
	}

	return false
}

// SetFinancialAccount gets a reference to the given FinancialAccountsPutFinancialAccount and assigns it to the FinancialAccount field.
func (o *FinancialAccountsPut) SetFinancialAccount(v FinancialAccountsPutFinancialAccount) {
	o.FinancialAccount = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *FinancialAccountsPut) GetPreviousValues() FinancialAccountsPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret FinancialAccountsPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialAccountsPut) GetPreviousValuesOk() (*FinancialAccountsPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *FinancialAccountsPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given FinancialAccountsPutPreviousValues and assigns it to the PreviousValues field.
func (o *FinancialAccountsPut) SetPreviousValues(v FinancialAccountsPutPreviousValues) {
	o.PreviousValues = &v
}

func (o FinancialAccountsPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialAccountsPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FinancialAccount) {
		toSerialize["financial_account"] = o.FinancialAccount
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableFinancialAccountsPut struct {
	value *FinancialAccountsPut
	isSet bool
}

func (v NullableFinancialAccountsPut) Get() *FinancialAccountsPut {
	return v.value
}

func (v *NullableFinancialAccountsPut) Set(val *FinancialAccountsPut) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialAccountsPut) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialAccountsPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialAccountsPut(val *FinancialAccountsPut) *NullableFinancialAccountsPut {
	return &NullableFinancialAccountsPut{value: val, isSet: true}
}

func (v NullableFinancialAccountsPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialAccountsPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


