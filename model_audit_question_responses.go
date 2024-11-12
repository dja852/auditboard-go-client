/*
AuditBoard Developer Portal API Documentation

External API reference documentation.

API version: 23.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auditboard

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AuditQuestionResponses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditQuestionResponses{}

// AuditQuestionResponses struct for AuditQuestionResponses
type AuditQuestionResponses struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Type string `json:"type"`
	// Note: This is a Foreign Key to `audit_surveys.id`.<fk table='audit_surveys' column='id'/>
	AuditSurveyId int32 `json:"audit_survey_id"`
	// Note: This is a Foreign Key to `audit_questions.id`.<fk table='audit_questions' column='id'/>
	AuditQuestionId int32 `json:"audit_question_id"`
	Notes *string `json:"notes,omitempty"`
	ResponseData interface{} `json:"response_data,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	AnsweredByUserId *int32 `json:"answered_by_user_id,omitempty"`
	AnsweredDate *string `json:"answered_date,omitempty"`
	IsAnswered bool `json:"is_answered"`
	IsNotApplicable bool `json:"is_not_applicable"`
	IsFailedResponse bool `json:"is_failed_response"`
	QuestionUuid string `json:"question_uuid"`
	OverrideData interface{} `json:"override_data,omitempty"`
	OverrideReason *string `json:"override_reason,omitempty"`
	OverriddenAt *string `json:"overridden_at,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	OverriddenByUserId *int32 `json:"overridden_by_user_id,omitempty"`
	IsNotApplicableOverride bool `json:"is_not_applicable_override"`
}

type _AuditQuestionResponses AuditQuestionResponses

// NewAuditQuestionResponses instantiates a new AuditQuestionResponses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditQuestionResponses(type_ string, auditSurveyId int32, auditQuestionId int32, isAnswered bool, isNotApplicable bool, isFailedResponse bool, questionUuid string, isNotApplicableOverride bool) *AuditQuestionResponses {
	this := AuditQuestionResponses{}
	this.Type = type_
	this.AuditSurveyId = auditSurveyId
	this.AuditQuestionId = auditQuestionId
	this.IsAnswered = isAnswered
	this.IsNotApplicable = isNotApplicable
	this.IsFailedResponse = isFailedResponse
	this.QuestionUuid = questionUuid
	this.IsNotApplicableOverride = isNotApplicableOverride
	return &this
}

// NewAuditQuestionResponsesWithDefaults instantiates a new AuditQuestionResponses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditQuestionResponsesWithDefaults() *AuditQuestionResponses {
	this := AuditQuestionResponses{}
	var isAnswered bool = false
	this.IsAnswered = isAnswered
	var isNotApplicable bool = false
	this.IsNotApplicable = isNotApplicable
	var isFailedResponse bool = false
	this.IsFailedResponse = isFailedResponse
	var isNotApplicableOverride bool = false
	this.IsNotApplicableOverride = isNotApplicableOverride
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuditQuestionResponses) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionResponses) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuditQuestionResponses) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AuditQuestionResponses) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AuditQuestionResponses) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionResponses) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AuditQuestionResponses) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AuditQuestionResponses) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AuditQuestionResponses) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionResponses) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AuditQuestionResponses) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AuditQuestionResponses) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *AuditQuestionResponses) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionResponses) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *AuditQuestionResponses) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *AuditQuestionResponses) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetType returns the Type field value
func (o *AuditQuestionResponses) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuditQuestionResponses) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuditQuestionResponses) SetType(v string) {
	o.Type = v
}

// GetAuditSurveyId returns the AuditSurveyId field value
func (o *AuditQuestionResponses) GetAuditSurveyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AuditSurveyId
}

// GetAuditSurveyIdOk returns a tuple with the AuditSurveyId field value
// and a boolean to check if the value has been set.
func (o *AuditQuestionResponses) GetAuditSurveyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuditSurveyId, true
}

// SetAuditSurveyId sets field value
func (o *AuditQuestionResponses) SetAuditSurveyId(v int32) {
	o.AuditSurveyId = v
}

// GetAuditQuestionId returns the AuditQuestionId field value
func (o *AuditQuestionResponses) GetAuditQuestionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AuditQuestionId
}

// GetAuditQuestionIdOk returns a tuple with the AuditQuestionId field value
// and a boolean to check if the value has been set.
func (o *AuditQuestionResponses) GetAuditQuestionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuditQuestionId, true
}

// SetAuditQuestionId sets field value
func (o *AuditQuestionResponses) SetAuditQuestionId(v int32) {
	o.AuditQuestionId = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *AuditQuestionResponses) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionResponses) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *AuditQuestionResponses) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *AuditQuestionResponses) SetNotes(v string) {
	o.Notes = &v
}

// GetResponseData returns the ResponseData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditQuestionResponses) GetResponseData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ResponseData
}

// GetResponseDataOk returns a tuple with the ResponseData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditQuestionResponses) GetResponseDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ResponseData) {
		return nil, false
	}
	return &o.ResponseData, true
}

// HasResponseData returns a boolean if a field has been set.
func (o *AuditQuestionResponses) HasResponseData() bool {
	if o != nil && !IsNil(o.ResponseData) {
		return true
	}

	return false
}

// SetResponseData gets a reference to the given interface{} and assigns it to the ResponseData field.
func (o *AuditQuestionResponses) SetResponseData(v interface{}) {
	o.ResponseData = v
}

// GetAnsweredByUserId returns the AnsweredByUserId field value if set, zero value otherwise.
func (o *AuditQuestionResponses) GetAnsweredByUserId() int32 {
	if o == nil || IsNil(o.AnsweredByUserId) {
		var ret int32
		return ret
	}
	return *o.AnsweredByUserId
}

// GetAnsweredByUserIdOk returns a tuple with the AnsweredByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionResponses) GetAnsweredByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AnsweredByUserId) {
		return nil, false
	}
	return o.AnsweredByUserId, true
}

// HasAnsweredByUserId returns a boolean if a field has been set.
func (o *AuditQuestionResponses) HasAnsweredByUserId() bool {
	if o != nil && !IsNil(o.AnsweredByUserId) {
		return true
	}

	return false
}

// SetAnsweredByUserId gets a reference to the given int32 and assigns it to the AnsweredByUserId field.
func (o *AuditQuestionResponses) SetAnsweredByUserId(v int32) {
	o.AnsweredByUserId = &v
}

// GetAnsweredDate returns the AnsweredDate field value if set, zero value otherwise.
func (o *AuditQuestionResponses) GetAnsweredDate() string {
	if o == nil || IsNil(o.AnsweredDate) {
		var ret string
		return ret
	}
	return *o.AnsweredDate
}

// GetAnsweredDateOk returns a tuple with the AnsweredDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionResponses) GetAnsweredDateOk() (*string, bool) {
	if o == nil || IsNil(o.AnsweredDate) {
		return nil, false
	}
	return o.AnsweredDate, true
}

// HasAnsweredDate returns a boolean if a field has been set.
func (o *AuditQuestionResponses) HasAnsweredDate() bool {
	if o != nil && !IsNil(o.AnsweredDate) {
		return true
	}

	return false
}

// SetAnsweredDate gets a reference to the given string and assigns it to the AnsweredDate field.
func (o *AuditQuestionResponses) SetAnsweredDate(v string) {
	o.AnsweredDate = &v
}

// GetIsAnswered returns the IsAnswered field value
func (o *AuditQuestionResponses) GetIsAnswered() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAnswered
}

// GetIsAnsweredOk returns a tuple with the IsAnswered field value
// and a boolean to check if the value has been set.
func (o *AuditQuestionResponses) GetIsAnsweredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAnswered, true
}

// SetIsAnswered sets field value
func (o *AuditQuestionResponses) SetIsAnswered(v bool) {
	o.IsAnswered = v
}

// GetIsNotApplicable returns the IsNotApplicable field value
func (o *AuditQuestionResponses) GetIsNotApplicable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsNotApplicable
}

// GetIsNotApplicableOk returns a tuple with the IsNotApplicable field value
// and a boolean to check if the value has been set.
func (o *AuditQuestionResponses) GetIsNotApplicableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsNotApplicable, true
}

// SetIsNotApplicable sets field value
func (o *AuditQuestionResponses) SetIsNotApplicable(v bool) {
	o.IsNotApplicable = v
}

// GetIsFailedResponse returns the IsFailedResponse field value
func (o *AuditQuestionResponses) GetIsFailedResponse() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFailedResponse
}

// GetIsFailedResponseOk returns a tuple with the IsFailedResponse field value
// and a boolean to check if the value has been set.
func (o *AuditQuestionResponses) GetIsFailedResponseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFailedResponse, true
}

// SetIsFailedResponse sets field value
func (o *AuditQuestionResponses) SetIsFailedResponse(v bool) {
	o.IsFailedResponse = v
}

// GetQuestionUuid returns the QuestionUuid field value
func (o *AuditQuestionResponses) GetQuestionUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QuestionUuid
}

// GetQuestionUuidOk returns a tuple with the QuestionUuid field value
// and a boolean to check if the value has been set.
func (o *AuditQuestionResponses) GetQuestionUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuestionUuid, true
}

// SetQuestionUuid sets field value
func (o *AuditQuestionResponses) SetQuestionUuid(v string) {
	o.QuestionUuid = v
}

// GetOverrideData returns the OverrideData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditQuestionResponses) GetOverrideData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.OverrideData
}

// GetOverrideDataOk returns a tuple with the OverrideData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditQuestionResponses) GetOverrideDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.OverrideData) {
		return nil, false
	}
	return &o.OverrideData, true
}

// HasOverrideData returns a boolean if a field has been set.
func (o *AuditQuestionResponses) HasOverrideData() bool {
	if o != nil && !IsNil(o.OverrideData) {
		return true
	}

	return false
}

// SetOverrideData gets a reference to the given interface{} and assigns it to the OverrideData field.
func (o *AuditQuestionResponses) SetOverrideData(v interface{}) {
	o.OverrideData = v
}

// GetOverrideReason returns the OverrideReason field value if set, zero value otherwise.
func (o *AuditQuestionResponses) GetOverrideReason() string {
	if o == nil || IsNil(o.OverrideReason) {
		var ret string
		return ret
	}
	return *o.OverrideReason
}

// GetOverrideReasonOk returns a tuple with the OverrideReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionResponses) GetOverrideReasonOk() (*string, bool) {
	if o == nil || IsNil(o.OverrideReason) {
		return nil, false
	}
	return o.OverrideReason, true
}

// HasOverrideReason returns a boolean if a field has been set.
func (o *AuditQuestionResponses) HasOverrideReason() bool {
	if o != nil && !IsNil(o.OverrideReason) {
		return true
	}

	return false
}

// SetOverrideReason gets a reference to the given string and assigns it to the OverrideReason field.
func (o *AuditQuestionResponses) SetOverrideReason(v string) {
	o.OverrideReason = &v
}

// GetOverriddenAt returns the OverriddenAt field value if set, zero value otherwise.
func (o *AuditQuestionResponses) GetOverriddenAt() string {
	if o == nil || IsNil(o.OverriddenAt) {
		var ret string
		return ret
	}
	return *o.OverriddenAt
}

// GetOverriddenAtOk returns a tuple with the OverriddenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionResponses) GetOverriddenAtOk() (*string, bool) {
	if o == nil || IsNil(o.OverriddenAt) {
		return nil, false
	}
	return o.OverriddenAt, true
}

// HasOverriddenAt returns a boolean if a field has been set.
func (o *AuditQuestionResponses) HasOverriddenAt() bool {
	if o != nil && !IsNil(o.OverriddenAt) {
		return true
	}

	return false
}

// SetOverriddenAt gets a reference to the given string and assigns it to the OverriddenAt field.
func (o *AuditQuestionResponses) SetOverriddenAt(v string) {
	o.OverriddenAt = &v
}

// GetOverriddenByUserId returns the OverriddenByUserId field value if set, zero value otherwise.
func (o *AuditQuestionResponses) GetOverriddenByUserId() int32 {
	if o == nil || IsNil(o.OverriddenByUserId) {
		var ret int32
		return ret
	}
	return *o.OverriddenByUserId
}

// GetOverriddenByUserIdOk returns a tuple with the OverriddenByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionResponses) GetOverriddenByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OverriddenByUserId) {
		return nil, false
	}
	return o.OverriddenByUserId, true
}

// HasOverriddenByUserId returns a boolean if a field has been set.
func (o *AuditQuestionResponses) HasOverriddenByUserId() bool {
	if o != nil && !IsNil(o.OverriddenByUserId) {
		return true
	}

	return false
}

// SetOverriddenByUserId gets a reference to the given int32 and assigns it to the OverriddenByUserId field.
func (o *AuditQuestionResponses) SetOverriddenByUserId(v int32) {
	o.OverriddenByUserId = &v
}

// GetIsNotApplicableOverride returns the IsNotApplicableOverride field value
func (o *AuditQuestionResponses) GetIsNotApplicableOverride() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsNotApplicableOverride
}

// GetIsNotApplicableOverrideOk returns a tuple with the IsNotApplicableOverride field value
// and a boolean to check if the value has been set.
func (o *AuditQuestionResponses) GetIsNotApplicableOverrideOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsNotApplicableOverride, true
}

// SetIsNotApplicableOverride sets field value
func (o *AuditQuestionResponses) SetIsNotApplicableOverride(v bool) {
	o.IsNotApplicableOverride = v
}

func (o AuditQuestionResponses) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditQuestionResponses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	toSerialize["type"] = o.Type
	toSerialize["audit_survey_id"] = o.AuditSurveyId
	toSerialize["audit_question_id"] = o.AuditQuestionId
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if o.ResponseData != nil {
		toSerialize["response_data"] = o.ResponseData
	}
	if !IsNil(o.AnsweredByUserId) {
		toSerialize["answered_by_user_id"] = o.AnsweredByUserId
	}
	if !IsNil(o.AnsweredDate) {
		toSerialize["answered_date"] = o.AnsweredDate
	}
	toSerialize["is_answered"] = o.IsAnswered
	toSerialize["is_not_applicable"] = o.IsNotApplicable
	toSerialize["is_failed_response"] = o.IsFailedResponse
	toSerialize["question_uuid"] = o.QuestionUuid
	if o.OverrideData != nil {
		toSerialize["override_data"] = o.OverrideData
	}
	if !IsNil(o.OverrideReason) {
		toSerialize["override_reason"] = o.OverrideReason
	}
	if !IsNil(o.OverriddenAt) {
		toSerialize["overridden_at"] = o.OverriddenAt
	}
	if !IsNil(o.OverriddenByUserId) {
		toSerialize["overridden_by_user_id"] = o.OverriddenByUserId
	}
	toSerialize["is_not_applicable_override"] = o.IsNotApplicableOverride
	return toSerialize, nil
}

func (o *AuditQuestionResponses) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"audit_survey_id",
		"audit_question_id",
		"is_answered",
		"is_not_applicable",
		"is_failed_response",
		"question_uuid",
		"is_not_applicable_override",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAuditQuestionResponses := _AuditQuestionResponses{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuditQuestionResponses)

	if err != nil {
		return err
	}

	*o = AuditQuestionResponses(varAuditQuestionResponses)

	return err
}

type NullableAuditQuestionResponses struct {
	value *AuditQuestionResponses
	isSet bool
}

func (v NullableAuditQuestionResponses) Get() *AuditQuestionResponses {
	return v.value
}

func (v *NullableAuditQuestionResponses) Set(val *AuditQuestionResponses) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditQuestionResponses) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditQuestionResponses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditQuestionResponses(val *AuditQuestionResponses) *NullableAuditQuestionResponses {
	return &NullableAuditQuestionResponses{value: val, isSet: true}
}

func (v NullableAuditQuestionResponses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditQuestionResponses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


