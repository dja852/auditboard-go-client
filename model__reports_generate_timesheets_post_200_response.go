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

// checks if the ReportsGenerateTimesheetsPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportsGenerateTimesheetsPost200Response{}

// ReportsGenerateTimesheetsPost200Response struct for ReportsGenerateTimesheetsPost200Response
type ReportsGenerateTimesheetsPost200Response struct {
	Status *string `json:"status,omitempty"`
}

// NewReportsGenerateTimesheetsPost200Response instantiates a new ReportsGenerateTimesheetsPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportsGenerateTimesheetsPost200Response() *ReportsGenerateTimesheetsPost200Response {
	this := ReportsGenerateTimesheetsPost200Response{}
	return &this
}

// NewReportsGenerateTimesheetsPost200ResponseWithDefaults instantiates a new ReportsGenerateTimesheetsPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportsGenerateTimesheetsPost200ResponseWithDefaults() *ReportsGenerateTimesheetsPost200Response {
	this := ReportsGenerateTimesheetsPost200Response{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ReportsGenerateTimesheetsPost200Response) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsGenerateTimesheetsPost200Response) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ReportsGenerateTimesheetsPost200Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ReportsGenerateTimesheetsPost200Response) SetStatus(v string) {
	o.Status = &v
}

func (o ReportsGenerateTimesheetsPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportsGenerateTimesheetsPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableReportsGenerateTimesheetsPost200Response struct {
	value *ReportsGenerateTimesheetsPost200Response
	isSet bool
}

func (v NullableReportsGenerateTimesheetsPost200Response) Get() *ReportsGenerateTimesheetsPost200Response {
	return v.value
}

func (v *NullableReportsGenerateTimesheetsPost200Response) Set(val *ReportsGenerateTimesheetsPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReportsGenerateTimesheetsPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReportsGenerateTimesheetsPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportsGenerateTimesheetsPost200Response(val *ReportsGenerateTimesheetsPost200Response) *NullableReportsGenerateTimesheetsPost200Response {
	return &NullableReportsGenerateTimesheetsPost200Response{value: val, isSet: true}
}

func (v NullableReportsGenerateTimesheetsPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportsGenerateTimesheetsPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


