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

// checks if the AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues{}

// AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues struct for AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues
type AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	// Note: This is a Foreign Key to `audit_survey_templates.id`.<fk table='audit_survey_templates' column='id'/>
	AuditSurveyTemplateId *int32 `json:"audit_survey_template_id,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	UserId *int32 `json:"user_id,omitempty"`
	SortOrder *int32 `json:"sort_order,omitempty"`
}

// NewAuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues instantiates a new AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues() *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues {
	this := AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues{}
	var sortOrder int32 = 0
	this.SortOrder = &sortOrder
	return &this
}

// NewAuditSurveyTemplatesDefaultReviewerUsersPutPreviousValuesWithDefaults instantiates a new AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditSurveyTemplatesDefaultReviewerUsersPutPreviousValuesWithDefaults() *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues {
	this := AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues{}
	var sortOrder int32 = 0
	this.SortOrder = &sortOrder
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetAuditSurveyTemplateId returns the AuditSurveyTemplateId field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetAuditSurveyTemplateId() int32 {
	if o == nil || IsNil(o.AuditSurveyTemplateId) {
		var ret int32
		return ret
	}
	return *o.AuditSurveyTemplateId
}

// GetAuditSurveyTemplateIdOk returns a tuple with the AuditSurveyTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetAuditSurveyTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AuditSurveyTemplateId) {
		return nil, false
	}
	return o.AuditSurveyTemplateId, true
}

// HasAuditSurveyTemplateId returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) HasAuditSurveyTemplateId() bool {
	if o != nil && !IsNil(o.AuditSurveyTemplateId) {
		return true
	}

	return false
}

// SetAuditSurveyTemplateId gets a reference to the given int32 and assigns it to the AuditSurveyTemplateId field.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) SetAuditSurveyTemplateId(v int32) {
	o.AuditSurveyTemplateId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) SetUserId(v int32) {
	o.UserId = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetSortOrder() int32 {
	if o == nil || IsNil(o.SortOrder) {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetSortOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) SetSortOrder(v int32) {
	o.SortOrder = &v
}

func (o AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AuditSurveyTemplateId) {
		toSerialize["audit_survey_template_id"] = o.AuditSurveyTemplateId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.SortOrder) {
		toSerialize["sort_order"] = o.SortOrder
	}
	return toSerialize, nil
}

type NullableAuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues struct {
	value *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues
	isSet bool
}

func (v NullableAuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) Get() *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues {
	return v.value
}

func (v *NullableAuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) Set(val *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues(val *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) *NullableAuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues {
	return &NullableAuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues{value: val, isSet: true}
}

func (v NullableAuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


