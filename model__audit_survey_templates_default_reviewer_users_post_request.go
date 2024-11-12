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

// checks if the AuditSurveyTemplatesDefaultReviewerUsersPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditSurveyTemplatesDefaultReviewerUsersPostRequest{}

// AuditSurveyTemplatesDefaultReviewerUsersPostRequest struct for AuditSurveyTemplatesDefaultReviewerUsersPostRequest
type AuditSurveyTemplatesDefaultReviewerUsersPostRequest struct {
	AuditSurveyTemplatesDefaultReviewerUser *AuditSurveyTemplatesDefaultReviewerUsers `json:"audit_survey _templates _default _reviewer _user,omitempty"`
}

// NewAuditSurveyTemplatesDefaultReviewerUsersPostRequest instantiates a new AuditSurveyTemplatesDefaultReviewerUsersPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditSurveyTemplatesDefaultReviewerUsersPostRequest() *AuditSurveyTemplatesDefaultReviewerUsersPostRequest {
	this := AuditSurveyTemplatesDefaultReviewerUsersPostRequest{}
	return &this
}

// NewAuditSurveyTemplatesDefaultReviewerUsersPostRequestWithDefaults instantiates a new AuditSurveyTemplatesDefaultReviewerUsersPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditSurveyTemplatesDefaultReviewerUsersPostRequestWithDefaults() *AuditSurveyTemplatesDefaultReviewerUsersPostRequest {
	this := AuditSurveyTemplatesDefaultReviewerUsersPostRequest{}
	return &this
}

// GetAuditSurveyTemplatesDefaultReviewerUser returns the AuditSurveyTemplatesDefaultReviewerUser field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPostRequest) GetAuditSurveyTemplatesDefaultReviewerUser() AuditSurveyTemplatesDefaultReviewerUsers {
	if o == nil || IsNil(o.AuditSurveyTemplatesDefaultReviewerUser) {
		var ret AuditSurveyTemplatesDefaultReviewerUsers
		return ret
	}
	return *o.AuditSurveyTemplatesDefaultReviewerUser
}

// GetAuditSurveyTemplatesDefaultReviewerUserOk returns a tuple with the AuditSurveyTemplatesDefaultReviewerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPostRequest) GetAuditSurveyTemplatesDefaultReviewerUserOk() (*AuditSurveyTemplatesDefaultReviewerUsers, bool) {
	if o == nil || IsNil(o.AuditSurveyTemplatesDefaultReviewerUser) {
		return nil, false
	}
	return o.AuditSurveyTemplatesDefaultReviewerUser, true
}

// HasAuditSurveyTemplatesDefaultReviewerUser returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPostRequest) HasAuditSurveyTemplatesDefaultReviewerUser() bool {
	if o != nil && !IsNil(o.AuditSurveyTemplatesDefaultReviewerUser) {
		return true
	}

	return false
}

// SetAuditSurveyTemplatesDefaultReviewerUser gets a reference to the given AuditSurveyTemplatesDefaultReviewerUsers and assigns it to the AuditSurveyTemplatesDefaultReviewerUser field.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPostRequest) SetAuditSurveyTemplatesDefaultReviewerUser(v AuditSurveyTemplatesDefaultReviewerUsers) {
	o.AuditSurveyTemplatesDefaultReviewerUser = &v
}

func (o AuditSurveyTemplatesDefaultReviewerUsersPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditSurveyTemplatesDefaultReviewerUsersPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuditSurveyTemplatesDefaultReviewerUser) {
		toSerialize["audit_survey _templates _default _reviewer _user"] = o.AuditSurveyTemplatesDefaultReviewerUser
	}
	return toSerialize, nil
}

type NullableAuditSurveyTemplatesDefaultReviewerUsersPostRequest struct {
	value *AuditSurveyTemplatesDefaultReviewerUsersPostRequest
	isSet bool
}

func (v NullableAuditSurveyTemplatesDefaultReviewerUsersPostRequest) Get() *AuditSurveyTemplatesDefaultReviewerUsersPostRequest {
	return v.value
}

func (v *NullableAuditSurveyTemplatesDefaultReviewerUsersPostRequest) Set(val *AuditSurveyTemplatesDefaultReviewerUsersPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditSurveyTemplatesDefaultReviewerUsersPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditSurveyTemplatesDefaultReviewerUsersPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditSurveyTemplatesDefaultReviewerUsersPostRequest(val *AuditSurveyTemplatesDefaultReviewerUsersPostRequest) *NullableAuditSurveyTemplatesDefaultReviewerUsersPostRequest {
	return &NullableAuditSurveyTemplatesDefaultReviewerUsersPostRequest{value: val, isSet: true}
}

func (v NullableAuditSurveyTemplatesDefaultReviewerUsersPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditSurveyTemplatesDefaultReviewerUsersPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

