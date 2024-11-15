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

// checks if the AuditSurveyTemplatesPutPreviousValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditSurveyTemplatesPutPreviousValues{}

// AuditSurveyTemplatesPutPreviousValues struct for AuditSurveyTemplatesPutPreviousValues
type AuditSurveyTemplatesPutPreviousValues struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	CreatedByUserId *int32 `json:"created_by_user_id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Note: This is a Foreign Key to `ops_audit_categories.id`.<fk table='ops_audit_categories' column='id'/>
	OpsAuditCategoryId *int32 `json:"ops_audit_category_id,omitempty"`
	Type *string `json:"type,omitempty"`
	Key *string `json:"key,omitempty"`
	// Note: This is a Foreign Key to `audit_surveys.id`.<fk table='audit_surveys' column='id'/>
	SourceSurveyId *int32 `json:"source_survey_id,omitempty"`
	AllowAuditableEntityOwnersViewResponses *bool `json:"allow_auditable_entity_owners_view_responses,omitempty"`
	// Note: This is a Foreign Key to `esg_metric_categories.id`.<fk table='esg_metric_categories' column='id'/>
	EsgMetricCategoryId *int32 `json:"esg_metric_category_id,omitempty"`
	MultiReviewerMode *string `json:"multi_reviewer_mode,omitempty"`
	RequiredReviewerCount *int32 `json:"required_reviewer_count,omitempty"`
	DefaultPreparerType *string `json:"default_preparer_type,omitempty"`
	IsFixed *bool `json:"is_fixed,omitempty"`
}

// NewAuditSurveyTemplatesPutPreviousValues instantiates a new AuditSurveyTemplatesPutPreviousValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditSurveyTemplatesPutPreviousValues() *AuditSurveyTemplatesPutPreviousValues {
	this := AuditSurveyTemplatesPutPreviousValues{}
	var allowAuditableEntityOwnersViewResponses bool = false
	this.AllowAuditableEntityOwnersViewResponses = &allowAuditableEntityOwnersViewResponses
	var multiReviewerMode string = "nonsequential"
	this.MultiReviewerMode = &multiReviewerMode
	var isFixed bool = false
	this.IsFixed = &isFixed
	return &this
}

// NewAuditSurveyTemplatesPutPreviousValuesWithDefaults instantiates a new AuditSurveyTemplatesPutPreviousValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditSurveyTemplatesPutPreviousValuesWithDefaults() *AuditSurveyTemplatesPutPreviousValues {
	this := AuditSurveyTemplatesPutPreviousValues{}
	var allowAuditableEntityOwnersViewResponses bool = false
	this.AllowAuditableEntityOwnersViewResponses = &allowAuditableEntityOwnersViewResponses
	var multiReviewerMode string = "nonsequential"
	this.MultiReviewerMode = &multiReviewerMode
	var isFixed bool = false
	this.IsFixed = &isFixed
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesPutPreviousValues) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AuditSurveyTemplatesPutPreviousValues) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesPutPreviousValues) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AuditSurveyTemplatesPutPreviousValues) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesPutPreviousValues) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AuditSurveyTemplatesPutPreviousValues) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesPutPreviousValues) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *AuditSurveyTemplatesPutPreviousValues) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesPutPreviousValues) GetCreatedByUserId() int32 {
	if o == nil || IsNil(o.CreatedByUserId) {
		var ret int32
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) GetCreatedByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedByUserId) {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) HasCreatedByUserId() bool {
	if o != nil && !IsNil(o.CreatedByUserId) {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given int32 and assigns it to the CreatedByUserId field.
func (o *AuditSurveyTemplatesPutPreviousValues) SetCreatedByUserId(v int32) {
	o.CreatedByUserId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesPutPreviousValues) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuditSurveyTemplatesPutPreviousValues) SetName(v string) {
	o.Name = &v
}

// GetOpsAuditCategoryId returns the OpsAuditCategoryId field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesPutPreviousValues) GetOpsAuditCategoryId() int32 {
	if o == nil || IsNil(o.OpsAuditCategoryId) {
		var ret int32
		return ret
	}
	return *o.OpsAuditCategoryId
}

// GetOpsAuditCategoryIdOk returns a tuple with the OpsAuditCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) GetOpsAuditCategoryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OpsAuditCategoryId) {
		return nil, false
	}
	return o.OpsAuditCategoryId, true
}

// HasOpsAuditCategoryId returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) HasOpsAuditCategoryId() bool {
	if o != nil && !IsNil(o.OpsAuditCategoryId) {
		return true
	}

	return false
}

// SetOpsAuditCategoryId gets a reference to the given int32 and assigns it to the OpsAuditCategoryId field.
func (o *AuditSurveyTemplatesPutPreviousValues) SetOpsAuditCategoryId(v int32) {
	o.OpsAuditCategoryId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesPutPreviousValues) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AuditSurveyTemplatesPutPreviousValues) SetType(v string) {
	o.Type = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesPutPreviousValues) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AuditSurveyTemplatesPutPreviousValues) SetKey(v string) {
	o.Key = &v
}

// GetSourceSurveyId returns the SourceSurveyId field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesPutPreviousValues) GetSourceSurveyId() int32 {
	if o == nil || IsNil(o.SourceSurveyId) {
		var ret int32
		return ret
	}
	return *o.SourceSurveyId
}

// GetSourceSurveyIdOk returns a tuple with the SourceSurveyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) GetSourceSurveyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SourceSurveyId) {
		return nil, false
	}
	return o.SourceSurveyId, true
}

// HasSourceSurveyId returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) HasSourceSurveyId() bool {
	if o != nil && !IsNil(o.SourceSurveyId) {
		return true
	}

	return false
}

// SetSourceSurveyId gets a reference to the given int32 and assigns it to the SourceSurveyId field.
func (o *AuditSurveyTemplatesPutPreviousValues) SetSourceSurveyId(v int32) {
	o.SourceSurveyId = &v
}

// GetAllowAuditableEntityOwnersViewResponses returns the AllowAuditableEntityOwnersViewResponses field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesPutPreviousValues) GetAllowAuditableEntityOwnersViewResponses() bool {
	if o == nil || IsNil(o.AllowAuditableEntityOwnersViewResponses) {
		var ret bool
		return ret
	}
	return *o.AllowAuditableEntityOwnersViewResponses
}

// GetAllowAuditableEntityOwnersViewResponsesOk returns a tuple with the AllowAuditableEntityOwnersViewResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) GetAllowAuditableEntityOwnersViewResponsesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowAuditableEntityOwnersViewResponses) {
		return nil, false
	}
	return o.AllowAuditableEntityOwnersViewResponses, true
}

// HasAllowAuditableEntityOwnersViewResponses returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) HasAllowAuditableEntityOwnersViewResponses() bool {
	if o != nil && !IsNil(o.AllowAuditableEntityOwnersViewResponses) {
		return true
	}

	return false
}

// SetAllowAuditableEntityOwnersViewResponses gets a reference to the given bool and assigns it to the AllowAuditableEntityOwnersViewResponses field.
func (o *AuditSurveyTemplatesPutPreviousValues) SetAllowAuditableEntityOwnersViewResponses(v bool) {
	o.AllowAuditableEntityOwnersViewResponses = &v
}

// GetEsgMetricCategoryId returns the EsgMetricCategoryId field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesPutPreviousValues) GetEsgMetricCategoryId() int32 {
	if o == nil || IsNil(o.EsgMetricCategoryId) {
		var ret int32
		return ret
	}
	return *o.EsgMetricCategoryId
}

// GetEsgMetricCategoryIdOk returns a tuple with the EsgMetricCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) GetEsgMetricCategoryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EsgMetricCategoryId) {
		return nil, false
	}
	return o.EsgMetricCategoryId, true
}

// HasEsgMetricCategoryId returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) HasEsgMetricCategoryId() bool {
	if o != nil && !IsNil(o.EsgMetricCategoryId) {
		return true
	}

	return false
}

// SetEsgMetricCategoryId gets a reference to the given int32 and assigns it to the EsgMetricCategoryId field.
func (o *AuditSurveyTemplatesPutPreviousValues) SetEsgMetricCategoryId(v int32) {
	o.EsgMetricCategoryId = &v
}

// GetMultiReviewerMode returns the MultiReviewerMode field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesPutPreviousValues) GetMultiReviewerMode() string {
	if o == nil || IsNil(o.MultiReviewerMode) {
		var ret string
		return ret
	}
	return *o.MultiReviewerMode
}

// GetMultiReviewerModeOk returns a tuple with the MultiReviewerMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) GetMultiReviewerModeOk() (*string, bool) {
	if o == nil || IsNil(o.MultiReviewerMode) {
		return nil, false
	}
	return o.MultiReviewerMode, true
}

// HasMultiReviewerMode returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) HasMultiReviewerMode() bool {
	if o != nil && !IsNil(o.MultiReviewerMode) {
		return true
	}

	return false
}

// SetMultiReviewerMode gets a reference to the given string and assigns it to the MultiReviewerMode field.
func (o *AuditSurveyTemplatesPutPreviousValues) SetMultiReviewerMode(v string) {
	o.MultiReviewerMode = &v
}

// GetRequiredReviewerCount returns the RequiredReviewerCount field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesPutPreviousValues) GetRequiredReviewerCount() int32 {
	if o == nil || IsNil(o.RequiredReviewerCount) {
		var ret int32
		return ret
	}
	return *o.RequiredReviewerCount
}

// GetRequiredReviewerCountOk returns a tuple with the RequiredReviewerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) GetRequiredReviewerCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RequiredReviewerCount) {
		return nil, false
	}
	return o.RequiredReviewerCount, true
}

// HasRequiredReviewerCount returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) HasRequiredReviewerCount() bool {
	if o != nil && !IsNil(o.RequiredReviewerCount) {
		return true
	}

	return false
}

// SetRequiredReviewerCount gets a reference to the given int32 and assigns it to the RequiredReviewerCount field.
func (o *AuditSurveyTemplatesPutPreviousValues) SetRequiredReviewerCount(v int32) {
	o.RequiredReviewerCount = &v
}

// GetDefaultPreparerType returns the DefaultPreparerType field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesPutPreviousValues) GetDefaultPreparerType() string {
	if o == nil || IsNil(o.DefaultPreparerType) {
		var ret string
		return ret
	}
	return *o.DefaultPreparerType
}

// GetDefaultPreparerTypeOk returns a tuple with the DefaultPreparerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) GetDefaultPreparerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultPreparerType) {
		return nil, false
	}
	return o.DefaultPreparerType, true
}

// HasDefaultPreparerType returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) HasDefaultPreparerType() bool {
	if o != nil && !IsNil(o.DefaultPreparerType) {
		return true
	}

	return false
}

// SetDefaultPreparerType gets a reference to the given string and assigns it to the DefaultPreparerType field.
func (o *AuditSurveyTemplatesPutPreviousValues) SetDefaultPreparerType(v string) {
	o.DefaultPreparerType = &v
}

// GetIsFixed returns the IsFixed field value if set, zero value otherwise.
func (o *AuditSurveyTemplatesPutPreviousValues) GetIsFixed() bool {
	if o == nil || IsNil(o.IsFixed) {
		var ret bool
		return ret
	}
	return *o.IsFixed
}

// GetIsFixedOk returns a tuple with the IsFixed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) GetIsFixedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFixed) {
		return nil, false
	}
	return o.IsFixed, true
}

// HasIsFixed returns a boolean if a field has been set.
func (o *AuditSurveyTemplatesPutPreviousValues) HasIsFixed() bool {
	if o != nil && !IsNil(o.IsFixed) {
		return true
	}

	return false
}

// SetIsFixed gets a reference to the given bool and assigns it to the IsFixed field.
func (o *AuditSurveyTemplatesPutPreviousValues) SetIsFixed(v bool) {
	o.IsFixed = &v
}

func (o AuditSurveyTemplatesPutPreviousValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditSurveyTemplatesPutPreviousValues) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CreatedByUserId) {
		toSerialize["created_by_user_id"] = o.CreatedByUserId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OpsAuditCategoryId) {
		toSerialize["ops_audit_category_id"] = o.OpsAuditCategoryId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.SourceSurveyId) {
		toSerialize["source_survey_id"] = o.SourceSurveyId
	}
	if !IsNil(o.AllowAuditableEntityOwnersViewResponses) {
		toSerialize["allow_auditable_entity_owners_view_responses"] = o.AllowAuditableEntityOwnersViewResponses
	}
	if !IsNil(o.EsgMetricCategoryId) {
		toSerialize["esg_metric_category_id"] = o.EsgMetricCategoryId
	}
	if !IsNil(o.MultiReviewerMode) {
		toSerialize["multi_reviewer_mode"] = o.MultiReviewerMode
	}
	if !IsNil(o.RequiredReviewerCount) {
		toSerialize["required_reviewer_count"] = o.RequiredReviewerCount
	}
	if !IsNil(o.DefaultPreparerType) {
		toSerialize["default_preparer_type"] = o.DefaultPreparerType
	}
	if !IsNil(o.IsFixed) {
		toSerialize["is_fixed"] = o.IsFixed
	}
	return toSerialize, nil
}

type NullableAuditSurveyTemplatesPutPreviousValues struct {
	value *AuditSurveyTemplatesPutPreviousValues
	isSet bool
}

func (v NullableAuditSurveyTemplatesPutPreviousValues) Get() *AuditSurveyTemplatesPutPreviousValues {
	return v.value
}

func (v *NullableAuditSurveyTemplatesPutPreviousValues) Set(val *AuditSurveyTemplatesPutPreviousValues) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditSurveyTemplatesPutPreviousValues) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditSurveyTemplatesPutPreviousValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditSurveyTemplatesPutPreviousValues(val *AuditSurveyTemplatesPutPreviousValues) *NullableAuditSurveyTemplatesPutPreviousValues {
	return &NullableAuditSurveyTemplatesPutPreviousValues{value: val, isSet: true}
}

func (v NullableAuditSurveyTemplatesPutPreviousValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditSurveyTemplatesPutPreviousValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


