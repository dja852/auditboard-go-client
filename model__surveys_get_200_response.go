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

// checks if the SurveysGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SurveysGet200Response{}

// SurveysGet200Response struct for SurveysGet200Response
type SurveysGet200Response struct {
	Surveys []Surveys `json:"surveys,omitempty"`
}

// NewSurveysGet200Response instantiates a new SurveysGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSurveysGet200Response() *SurveysGet200Response {
	this := SurveysGet200Response{}
	return &this
}

// NewSurveysGet200ResponseWithDefaults instantiates a new SurveysGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSurveysGet200ResponseWithDefaults() *SurveysGet200Response {
	this := SurveysGet200Response{}
	return &this
}

// GetSurveys returns the Surveys field value if set, zero value otherwise.
func (o *SurveysGet200Response) GetSurveys() []Surveys {
	if o == nil || IsNil(o.Surveys) {
		var ret []Surveys
		return ret
	}
	return o.Surveys
}

// GetSurveysOk returns a tuple with the Surveys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveysGet200Response) GetSurveysOk() ([]Surveys, bool) {
	if o == nil || IsNil(o.Surveys) {
		return nil, false
	}
	return o.Surveys, true
}

// HasSurveys returns a boolean if a field has been set.
func (o *SurveysGet200Response) HasSurveys() bool {
	if o != nil && !IsNil(o.Surveys) {
		return true
	}

	return false
}

// SetSurveys gets a reference to the given []Surveys and assigns it to the Surveys field.
func (o *SurveysGet200Response) SetSurveys(v []Surveys) {
	o.Surveys = v
}

func (o SurveysGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SurveysGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Surveys) {
		toSerialize["surveys"] = o.Surveys
	}
	return toSerialize, nil
}

type NullableSurveysGet200Response struct {
	value *SurveysGet200Response
	isSet bool
}

func (v NullableSurveysGet200Response) Get() *SurveysGet200Response {
	return v.value
}

func (v *NullableSurveysGet200Response) Set(val *SurveysGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSurveysGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSurveysGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSurveysGet200Response(val *SurveysGet200Response) *NullableSurveysGet200Response {
	return &NullableSurveysGet200Response{value: val, isSet: true}
}

func (v NullableSurveysGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSurveysGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


