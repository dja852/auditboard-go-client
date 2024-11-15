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

// checks if the AuditSurveyTemplatesDefaultReviewerUsers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditSurveyTemplatesDefaultReviewerUsers{}

// AuditSurveyTemplatesDefaultReviewerUsers struct for AuditSurveyTemplatesDefaultReviewerUsers
type AuditSurveyTemplatesDefaultReviewerUsers struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	// Note: This is a Foreign Key to `audit_survey_templates.id`.<fk table='audit_survey_templates' column='id'/>
	AuditSurveyTemplateId int32 `json:"audit_survey_template_id"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	UserId int32 `json:"user_id"`
	SortOrder int32 `json:"sort_order"`
}

type _AuditSurveyTemplatesDefaultReviewerUsers AuditSurveyTemplatesDefaultReviewerUsers

// NewAuditSurveyTemplatesDefaultReviewerUsers instantiates a new AuditSurveyTemplatesDefaultReviewerUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditSurveyTemplatesDefaultReviewerUsers(auditSurveyTemplateId int32, userId int32, sortOrder int32) *AuditSurveyTemplatesDefaultReviewerUsers {
	this := AuditSurveyTemplatesDefaultReviewerUsers{}
	this.AuditSurveyTemplateId = auditSurveyTemplateId
	this.UserId = userId
	this.SortOrder = sortOrder
	return &this
}

// NewAuditSurveyTemplatesDefaultReviewerUsersWithDefaults instantiates a new AuditSurveyTemplatesDefaultReviewerUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditSurveyTemplatesDefaultReviewerUsersWithDefaults() *AuditSurveyTemplatesDefaultReviewerUsers {
	this := AuditSurveyTemplatesDefaultReviewerUsers{}
	var sortOrder int32 = 0
	this.SortOrder = sortOrder
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesDefaultReviewerUsers) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsers) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsers) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AuditSurveyTemplatesDefaultReviewerUsers) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesDefaultReviewerUsers) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsers) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsers) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AuditSurveyTemplatesDefaultReviewerUsers) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesDefaultReviewerUsers) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsers) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsers) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AuditSurveyTemplatesDefaultReviewerUsers) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesDefaultReviewerUsers) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsers) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsers) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *AuditSurveyTemplatesDefaultReviewerUsers) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetAuditSurveyTemplateId returns the AuditSurveyTemplateId field value
func (o *AuditSurveyTemplatesDefaultReviewerUsers) GetAuditSurveyTemplateId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AuditSurveyTemplateId
}

// GetAuditSurveyTemplateIdOk returns a tuple with the AuditSurveyTemplateId field value
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsers) GetAuditSurveyTemplateIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuditSurveyTemplateId, true
}

// SetAuditSurveyTemplateId sets field value
func (o *AuditSurveyTemplatesDefaultReviewerUsers) SetAuditSurveyTemplateId(v int32) {
	o.AuditSurveyTemplateId = v
}

// GetUserId returns the UserId field value
func (o *AuditSurveyTemplatesDefaultReviewerUsers) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsers) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *AuditSurveyTemplatesDefaultReviewerUsers) SetUserId(v int32) {
	o.UserId = v
}

// GetSortOrder returns the SortOrder field value
func (o *AuditSurveyTemplatesDefaultReviewerUsers) GetSortOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsers) GetSortOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortOrder, true
}

// SetSortOrder sets field value
func (o *AuditSurveyTemplatesDefaultReviewerUsers) SetSortOrder(v int32) {
	o.SortOrder = v
}

func (o AuditSurveyTemplatesDefaultReviewerUsers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditSurveyTemplatesDefaultReviewerUsers) ToMap() (map[string]interface{}, error) {
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
	toSerialize["audit_survey_template_id"] = o.AuditSurveyTemplateId
	toSerialize["user_id"] = o.UserId
	toSerialize["sort_order"] = o.SortOrder
	return toSerialize, nil
}

type NullableAuditSurveyTemplatesDefaultReviewerUsers struct {
	value *AuditSurveyTemplatesDefaultReviewerUsers
	isSet bool
}

func (v NullableAuditSurveyTemplatesDefaultReviewerUsers) Get() *AuditSurveyTemplatesDefaultReviewerUsers {
	return v.value
}

func (v *NullableAuditSurveyTemplatesDefaultReviewerUsers) Set(val *AuditSurveyTemplatesDefaultReviewerUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditSurveyTemplatesDefaultReviewerUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditSurveyTemplatesDefaultReviewerUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditSurveyTemplatesDefaultReviewerUsers(val *AuditSurveyTemplatesDefaultReviewerUsers) *NullableAuditSurveyTemplatesDefaultReviewerUsers {
	return &NullableAuditSurveyTemplatesDefaultReviewerUsers{value: val, isSet: true}
}

func (v NullableAuditSurveyTemplatesDefaultReviewerUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditSurveyTemplatesDefaultReviewerUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


