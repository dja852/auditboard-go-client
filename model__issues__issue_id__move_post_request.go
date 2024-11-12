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

// checks if the IssuesIssueIdMovePostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssuesIssueIdMovePostRequest{}

// IssuesIssueIdMovePostRequest struct for IssuesIssueIdMovePostRequest
type IssuesIssueIdMovePostRequest struct {
	Data *IssuesIssueIdMovePostRequestData `json:"data,omitempty"`
}

// NewIssuesIssueIdMovePostRequest instantiates a new IssuesIssueIdMovePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuesIssueIdMovePostRequest() *IssuesIssueIdMovePostRequest {
	this := IssuesIssueIdMovePostRequest{}
	return &this
}

// NewIssuesIssueIdMovePostRequestWithDefaults instantiates a new IssuesIssueIdMovePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuesIssueIdMovePostRequestWithDefaults() *IssuesIssueIdMovePostRequest {
	this := IssuesIssueIdMovePostRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *IssuesIssueIdMovePostRequest) GetData() IssuesIssueIdMovePostRequestData {
	if o == nil || IsNil(o.Data) {
		var ret IssuesIssueIdMovePostRequestData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesIssueIdMovePostRequest) GetDataOk() (*IssuesIssueIdMovePostRequestData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *IssuesIssueIdMovePostRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given IssuesIssueIdMovePostRequestData and assigns it to the Data field.
func (o *IssuesIssueIdMovePostRequest) SetData(v IssuesIssueIdMovePostRequestData) {
	o.Data = &v
}

func (o IssuesIssueIdMovePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssuesIssueIdMovePostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableIssuesIssueIdMovePostRequest struct {
	value *IssuesIssueIdMovePostRequest
	isSet bool
}

func (v NullableIssuesIssueIdMovePostRequest) Get() *IssuesIssueIdMovePostRequest {
	return v.value
}

func (v *NullableIssuesIssueIdMovePostRequest) Set(val *IssuesIssueIdMovePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuesIssueIdMovePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuesIssueIdMovePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuesIssueIdMovePostRequest(val *IssuesIssueIdMovePostRequest) *NullableIssuesIssueIdMovePostRequest {
	return &NullableIssuesIssueIdMovePostRequest{value: val, isSet: true}
}

func (v NullableIssuesIssueIdMovePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuesIssueIdMovePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

