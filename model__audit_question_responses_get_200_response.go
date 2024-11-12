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

// checks if the AuditQuestionResponsesGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditQuestionResponsesGet200Response{}

// AuditQuestionResponsesGet200Response struct for AuditQuestionResponsesGet200Response
type AuditQuestionResponsesGet200Response struct {
	AuditQuestionResponses []AuditQuestionResponses `json:"audit_question_responses,omitempty"`
}

// NewAuditQuestionResponsesGet200Response instantiates a new AuditQuestionResponsesGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditQuestionResponsesGet200Response() *AuditQuestionResponsesGet200Response {
	this := AuditQuestionResponsesGet200Response{}
	return &this
}

// NewAuditQuestionResponsesGet200ResponseWithDefaults instantiates a new AuditQuestionResponsesGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditQuestionResponsesGet200ResponseWithDefaults() *AuditQuestionResponsesGet200Response {
	this := AuditQuestionResponsesGet200Response{}
	return &this
}

// GetAuditQuestionResponses returns the AuditQuestionResponses field value if set, zero value otherwise.
func (o *AuditQuestionResponsesGet200Response) GetAuditQuestionResponses() []AuditQuestionResponses {
	if o == nil || IsNil(o.AuditQuestionResponses) {
		var ret []AuditQuestionResponses
		return ret
	}
	return o.AuditQuestionResponses
}

// GetAuditQuestionResponsesOk returns a tuple with the AuditQuestionResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionResponsesGet200Response) GetAuditQuestionResponsesOk() ([]AuditQuestionResponses, bool) {
	if o == nil || IsNil(o.AuditQuestionResponses) {
		return nil, false
	}
	return o.AuditQuestionResponses, true
}

// HasAuditQuestionResponses returns a boolean if a field has been set.
func (o *AuditQuestionResponsesGet200Response) HasAuditQuestionResponses() bool {
	if o != nil && !IsNil(o.AuditQuestionResponses) {
		return true
	}

	return false
}

// SetAuditQuestionResponses gets a reference to the given []AuditQuestionResponses and assigns it to the AuditQuestionResponses field.
func (o *AuditQuestionResponsesGet200Response) SetAuditQuestionResponses(v []AuditQuestionResponses) {
	o.AuditQuestionResponses = v
}

func (o AuditQuestionResponsesGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditQuestionResponsesGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuditQuestionResponses) {
		toSerialize["audit_question_responses"] = o.AuditQuestionResponses
	}
	return toSerialize, nil
}

type NullableAuditQuestionResponsesGet200Response struct {
	value *AuditQuestionResponsesGet200Response
	isSet bool
}

func (v NullableAuditQuestionResponsesGet200Response) Get() *AuditQuestionResponsesGet200Response {
	return v.value
}

func (v *NullableAuditQuestionResponsesGet200Response) Set(val *AuditQuestionResponsesGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditQuestionResponsesGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditQuestionResponsesGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditQuestionResponsesGet200Response(val *AuditQuestionResponsesGet200Response) *NullableAuditQuestionResponsesGet200Response {
	return &NullableAuditQuestionResponsesGet200Response{value: val, isSet: true}
}

func (v NullableAuditQuestionResponsesGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditQuestionResponsesGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


