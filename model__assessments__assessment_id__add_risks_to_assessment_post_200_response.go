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

// checks if the AssessmentsAssessmentIdAddRisksToAssessmentPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssessmentsAssessmentIdAddRisksToAssessmentPost200Response{}

// AssessmentsAssessmentIdAddRisksToAssessmentPost200Response struct for AssessmentsAssessmentIdAddRisksToAssessmentPost200Response
type AssessmentsAssessmentIdAddRisksToAssessmentPost200Response struct {
	Assessments []Assessments `json:"assessments,omitempty"`
	AssessmentResponses []AssessmentResponses `json:"assessment_responses,omitempty"`
	Risks []Risks `json:"risks,omitempty"`
}

// NewAssessmentsAssessmentIdAddRisksToAssessmentPost200Response instantiates a new AssessmentsAssessmentIdAddRisksToAssessmentPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssessmentsAssessmentIdAddRisksToAssessmentPost200Response() *AssessmentsAssessmentIdAddRisksToAssessmentPost200Response {
	this := AssessmentsAssessmentIdAddRisksToAssessmentPost200Response{}
	return &this
}

// NewAssessmentsAssessmentIdAddRisksToAssessmentPost200ResponseWithDefaults instantiates a new AssessmentsAssessmentIdAddRisksToAssessmentPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssessmentsAssessmentIdAddRisksToAssessmentPost200ResponseWithDefaults() *AssessmentsAssessmentIdAddRisksToAssessmentPost200Response {
	this := AssessmentsAssessmentIdAddRisksToAssessmentPost200Response{}
	return &this
}

// GetAssessments returns the Assessments field value if set, zero value otherwise.
func (o *AssessmentsAssessmentIdAddRisksToAssessmentPost200Response) GetAssessments() []Assessments {
	if o == nil || IsNil(o.Assessments) {
		var ret []Assessments
		return ret
	}
	return o.Assessments
}

// GetAssessmentsOk returns a tuple with the Assessments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentsAssessmentIdAddRisksToAssessmentPost200Response) GetAssessmentsOk() ([]Assessments, bool) {
	if o == nil || IsNil(o.Assessments) {
		return nil, false
	}
	return o.Assessments, true
}

// HasAssessments returns a boolean if a field has been set.
func (o *AssessmentsAssessmentIdAddRisksToAssessmentPost200Response) HasAssessments() bool {
	if o != nil && !IsNil(o.Assessments) {
		return true
	}

	return false
}

// SetAssessments gets a reference to the given []Assessments and assigns it to the Assessments field.
func (o *AssessmentsAssessmentIdAddRisksToAssessmentPost200Response) SetAssessments(v []Assessments) {
	o.Assessments = v
}

// GetAssessmentResponses returns the AssessmentResponses field value if set, zero value otherwise.
func (o *AssessmentsAssessmentIdAddRisksToAssessmentPost200Response) GetAssessmentResponses() []AssessmentResponses {
	if o == nil || IsNil(o.AssessmentResponses) {
		var ret []AssessmentResponses
		return ret
	}
	return o.AssessmentResponses
}

// GetAssessmentResponsesOk returns a tuple with the AssessmentResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentsAssessmentIdAddRisksToAssessmentPost200Response) GetAssessmentResponsesOk() ([]AssessmentResponses, bool) {
	if o == nil || IsNil(o.AssessmentResponses) {
		return nil, false
	}
	return o.AssessmentResponses, true
}

// HasAssessmentResponses returns a boolean if a field has been set.
func (o *AssessmentsAssessmentIdAddRisksToAssessmentPost200Response) HasAssessmentResponses() bool {
	if o != nil && !IsNil(o.AssessmentResponses) {
		return true
	}

	return false
}

// SetAssessmentResponses gets a reference to the given []AssessmentResponses and assigns it to the AssessmentResponses field.
func (o *AssessmentsAssessmentIdAddRisksToAssessmentPost200Response) SetAssessmentResponses(v []AssessmentResponses) {
	o.AssessmentResponses = v
}

// GetRisks returns the Risks field value if set, zero value otherwise.
func (o *AssessmentsAssessmentIdAddRisksToAssessmentPost200Response) GetRisks() []Risks {
	if o == nil || IsNil(o.Risks) {
		var ret []Risks
		return ret
	}
	return o.Risks
}

// GetRisksOk returns a tuple with the Risks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentsAssessmentIdAddRisksToAssessmentPost200Response) GetRisksOk() ([]Risks, bool) {
	if o == nil || IsNil(o.Risks) {
		return nil, false
	}
	return o.Risks, true
}

// HasRisks returns a boolean if a field has been set.
func (o *AssessmentsAssessmentIdAddRisksToAssessmentPost200Response) HasRisks() bool {
	if o != nil && !IsNil(o.Risks) {
		return true
	}

	return false
}

// SetRisks gets a reference to the given []Risks and assigns it to the Risks field.
func (o *AssessmentsAssessmentIdAddRisksToAssessmentPost200Response) SetRisks(v []Risks) {
	o.Risks = v
}

func (o AssessmentsAssessmentIdAddRisksToAssessmentPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssessmentsAssessmentIdAddRisksToAssessmentPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assessments) {
		toSerialize["assessments"] = o.Assessments
	}
	if !IsNil(o.AssessmentResponses) {
		toSerialize["assessment_responses"] = o.AssessmentResponses
	}
	if !IsNil(o.Risks) {
		toSerialize["risks"] = o.Risks
	}
	return toSerialize, nil
}

type NullableAssessmentsAssessmentIdAddRisksToAssessmentPost200Response struct {
	value *AssessmentsAssessmentIdAddRisksToAssessmentPost200Response
	isSet bool
}

func (v NullableAssessmentsAssessmentIdAddRisksToAssessmentPost200Response) Get() *AssessmentsAssessmentIdAddRisksToAssessmentPost200Response {
	return v.value
}

func (v *NullableAssessmentsAssessmentIdAddRisksToAssessmentPost200Response) Set(val *AssessmentsAssessmentIdAddRisksToAssessmentPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAssessmentsAssessmentIdAddRisksToAssessmentPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAssessmentsAssessmentIdAddRisksToAssessmentPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssessmentsAssessmentIdAddRisksToAssessmentPost200Response(val *AssessmentsAssessmentIdAddRisksToAssessmentPost200Response) *NullableAssessmentsAssessmentIdAddRisksToAssessmentPost200Response {
	return &NullableAssessmentsAssessmentIdAddRisksToAssessmentPost200Response{value: val, isSet: true}
}

func (v NullableAssessmentsAssessmentIdAddRisksToAssessmentPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssessmentsAssessmentIdAddRisksToAssessmentPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

