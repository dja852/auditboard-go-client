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

// checks if the FinancialAccountsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialAccountsGet200Response{}

// FinancialAccountsGet200Response struct for FinancialAccountsGet200Response
type FinancialAccountsGet200Response struct {
	FinancialAccounts []FinancialAccounts `json:"financial_accounts,omitempty"`
}

// NewFinancialAccountsGet200Response instantiates a new FinancialAccountsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialAccountsGet200Response() *FinancialAccountsGet200Response {
	this := FinancialAccountsGet200Response{}
	return &this
}

// NewFinancialAccountsGet200ResponseWithDefaults instantiates a new FinancialAccountsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialAccountsGet200ResponseWithDefaults() *FinancialAccountsGet200Response {
	this := FinancialAccountsGet200Response{}
	return &this
}

// GetFinancialAccounts returns the FinancialAccounts field value if set, zero value otherwise.
func (o *FinancialAccountsGet200Response) GetFinancialAccounts() []FinancialAccounts {
	if o == nil || IsNil(o.FinancialAccounts) {
		var ret []FinancialAccounts
		return ret
	}
	return o.FinancialAccounts
}

// GetFinancialAccountsOk returns a tuple with the FinancialAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialAccountsGet200Response) GetFinancialAccountsOk() ([]FinancialAccounts, bool) {
	if o == nil || IsNil(o.FinancialAccounts) {
		return nil, false
	}
	return o.FinancialAccounts, true
}

// HasFinancialAccounts returns a boolean if a field has been set.
func (o *FinancialAccountsGet200Response) HasFinancialAccounts() bool {
	if o != nil && !IsNil(o.FinancialAccounts) {
		return true
	}

	return false
}

// SetFinancialAccounts gets a reference to the given []FinancialAccounts and assigns it to the FinancialAccounts field.
func (o *FinancialAccountsGet200Response) SetFinancialAccounts(v []FinancialAccounts) {
	o.FinancialAccounts = v
}

func (o FinancialAccountsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialAccountsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FinancialAccounts) {
		toSerialize["financial_accounts"] = o.FinancialAccounts
	}
	return toSerialize, nil
}

type NullableFinancialAccountsGet200Response struct {
	value *FinancialAccountsGet200Response
	isSet bool
}

func (v NullableFinancialAccountsGet200Response) Get() *FinancialAccountsGet200Response {
	return v.value
}

func (v *NullableFinancialAccountsGet200Response) Set(val *FinancialAccountsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialAccountsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialAccountsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialAccountsGet200Response(val *FinancialAccountsGet200Response) *NullableFinancialAccountsGet200Response {
	return &NullableFinancialAccountsGet200Response{value: val, isSet: true}
}

func (v NullableFinancialAccountsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialAccountsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

