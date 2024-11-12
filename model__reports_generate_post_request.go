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

// checks if the ReportsGeneratePostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportsGeneratePostRequest{}

// ReportsGeneratePostRequest struct for ReportsGeneratePostRequest
type ReportsGeneratePostRequest struct {
	ReportType *string `json:"reportType,omitempty"`
	QueryData *ReportsGeneratePostRequestQueryData `json:"queryData,omitempty"`
}

// NewReportsGeneratePostRequest instantiates a new ReportsGeneratePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportsGeneratePostRequest() *ReportsGeneratePostRequest {
	this := ReportsGeneratePostRequest{}
	return &this
}

// NewReportsGeneratePostRequestWithDefaults instantiates a new ReportsGeneratePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportsGeneratePostRequestWithDefaults() *ReportsGeneratePostRequest {
	this := ReportsGeneratePostRequest{}
	return &this
}

// GetReportType returns the ReportType field value if set, zero value otherwise.
func (o *ReportsGeneratePostRequest) GetReportType() string {
	if o == nil || IsNil(o.ReportType) {
		var ret string
		return ret
	}
	return *o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsGeneratePostRequest) GetReportTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReportType) {
		return nil, false
	}
	return o.ReportType, true
}

// HasReportType returns a boolean if a field has been set.
func (o *ReportsGeneratePostRequest) HasReportType() bool {
	if o != nil && !IsNil(o.ReportType) {
		return true
	}

	return false
}

// SetReportType gets a reference to the given string and assigns it to the ReportType field.
func (o *ReportsGeneratePostRequest) SetReportType(v string) {
	o.ReportType = &v
}

// GetQueryData returns the QueryData field value if set, zero value otherwise.
func (o *ReportsGeneratePostRequest) GetQueryData() ReportsGeneratePostRequestQueryData {
	if o == nil || IsNil(o.QueryData) {
		var ret ReportsGeneratePostRequestQueryData
		return ret
	}
	return *o.QueryData
}

// GetQueryDataOk returns a tuple with the QueryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsGeneratePostRequest) GetQueryDataOk() (*ReportsGeneratePostRequestQueryData, bool) {
	if o == nil || IsNil(o.QueryData) {
		return nil, false
	}
	return o.QueryData, true
}

// HasQueryData returns a boolean if a field has been set.
func (o *ReportsGeneratePostRequest) HasQueryData() bool {
	if o != nil && !IsNil(o.QueryData) {
		return true
	}

	return false
}

// SetQueryData gets a reference to the given ReportsGeneratePostRequestQueryData and assigns it to the QueryData field.
func (o *ReportsGeneratePostRequest) SetQueryData(v ReportsGeneratePostRequestQueryData) {
	o.QueryData = &v
}

func (o ReportsGeneratePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportsGeneratePostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReportType) {
		toSerialize["reportType"] = o.ReportType
	}
	if !IsNil(o.QueryData) {
		toSerialize["queryData"] = o.QueryData
	}
	return toSerialize, nil
}

type NullableReportsGeneratePostRequest struct {
	value *ReportsGeneratePostRequest
	isSet bool
}

func (v NullableReportsGeneratePostRequest) Get() *ReportsGeneratePostRequest {
	return v.value
}

func (v *NullableReportsGeneratePostRequest) Set(val *ReportsGeneratePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReportsGeneratePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReportsGeneratePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportsGeneratePostRequest(val *ReportsGeneratePostRequest) *NullableReportsGeneratePostRequest {
	return &NullableReportsGeneratePostRequest{value: val, isSet: true}
}

func (v NullableReportsGeneratePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportsGeneratePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

