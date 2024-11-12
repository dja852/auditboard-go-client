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

// checks if the AuditSurveysGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditSurveysGet200Response{}

// AuditSurveysGet200Response struct for AuditSurveysGet200Response
type AuditSurveysGet200Response struct {
	AuditSurveys []AuditSurveys `json:"audit_surveys,omitempty"`
}

// NewAuditSurveysGet200Response instantiates a new AuditSurveysGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditSurveysGet200Response() *AuditSurveysGet200Response {
	this := AuditSurveysGet200Response{}
	return &this
}

// NewAuditSurveysGet200ResponseWithDefaults instantiates a new AuditSurveysGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditSurveysGet200ResponseWithDefaults() *AuditSurveysGet200Response {
	this := AuditSurveysGet200Response{}
	return &this
}

// GetAuditSurveys returns the AuditSurveys field value if set, zero value otherwise.
func (o *AuditSurveysGet200Response) GetAuditSurveys() []AuditSurveys {
	if o == nil || IsNil(o.AuditSurveys) {
		var ret []AuditSurveys
		return ret
	}
	return o.AuditSurveys
}

// GetAuditSurveysOk returns a tuple with the AuditSurveys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveysGet200Response) GetAuditSurveysOk() ([]AuditSurveys, bool) {
	if o == nil || IsNil(o.AuditSurveys) {
		return nil, false
	}
	return o.AuditSurveys, true
}

// HasAuditSurveys returns a boolean if a field has been set.
func (o *AuditSurveysGet200Response) HasAuditSurveys() bool {
	if o != nil && !IsNil(o.AuditSurveys) {
		return true
	}

	return false
}

// SetAuditSurveys gets a reference to the given []AuditSurveys and assigns it to the AuditSurveys field.
func (o *AuditSurveysGet200Response) SetAuditSurveys(v []AuditSurveys) {
	o.AuditSurveys = v
}

func (o AuditSurveysGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditSurveysGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuditSurveys) {
		toSerialize["audit_surveys"] = o.AuditSurveys
	}
	return toSerialize, nil
}

type NullableAuditSurveysGet200Response struct {
	value *AuditSurveysGet200Response
	isSet bool
}

func (v NullableAuditSurveysGet200Response) Get() *AuditSurveysGet200Response {
	return v.value
}

func (v *NullableAuditSurveysGet200Response) Set(val *AuditSurveysGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditSurveysGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditSurveysGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditSurveysGet200Response(val *AuditSurveysGet200Response) *NullableAuditSurveysGet200Response {
	return &NullableAuditSurveysGet200Response{value: val, isSet: true}
}

func (v NullableAuditSurveysGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditSurveysGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


