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

// checks if the EsgTopics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsgTopics{}

// EsgTopics struct for EsgTopics
type EsgTopics struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Uid *string `json:"uid,omitempty"`
	FieldData interface{} `json:"field_data,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	ReviewerUserId *int32 `json:"reviewer_user_id,omitempty"`
	EffectiveDate *string `json:"effective_date,omitempty"`
	OrganizationImpact interface{} `json:"organization_impact,omitempty"`
	StakeholderImpact interface{} `json:"stakeholder_impact,omitempty"`
	MaterialityRating interface{} `json:"materiality_rating,omitempty"`
	Category *string `json:"category,omitempty"`
	Scope *string `json:"scope,omitempty"`
	// Note: This is a Foreign Key to `esg_categories.id`.<fk table='esg_categories' column='id'/>
	EsgCategoryId *int32 `json:"esg_category_id,omitempty"`
	// Note: This is a Foreign Key to `esg_topic_custom_select1_options.id`.<fk table='esg_topic_custom_select1_options' column='id'/>
	EsgTopicCustomSelect1OptionId *int32 `json:"esg_topic_custom_select1_option_id,omitempty"`
	// Note: This is a Foreign Key to `esg_topic_custom_select2_options.id`.<fk table='esg_topic_custom_select2_options' column='id'/>
	EsgTopicCustomSelect2OptionId *int32 `json:"esg_topic_custom_select2_option_id,omitempty"`
	// Note: This is a Foreign Key to `esg_topic_custom_select3_options.id`.<fk table='esg_topic_custom_select3_options' column='id'/>
	EsgTopicCustomSelect3OptionId *int32 `json:"esg_topic_custom_select3_option_id,omitempty"`
	// Note: This is a Foreign Key to `esg_topic_custom_select4_options.id`.<fk table='esg_topic_custom_select4_options' column='id'/>
	EsgTopicCustomSelect4OptionId *int32 `json:"esg_topic_custom_select4_option_id,omitempty"`
	// Note: This is a Foreign Key to `esg_topic_custom_select5_options.id`.<fk table='esg_topic_custom_select5_options' column='id'/>
	EsgTopicCustomSelect5OptionId *int32 `json:"esg_topic_custom_select5_option_id,omitempty"`
	CustomDate1 *string `json:"custom_date1,omitempty"`
	CustomDate2 *string `json:"custom_date2,omitempty"`
	CustomDate3 *string `json:"custom_date3,omitempty"`
	CustomText1 *string `json:"custom_text1,omitempty"`
	CustomText2 *string `json:"custom_text2,omitempty"`
	CustomText3 *string `json:"custom_text3,omitempty"`
	CustomText4 *string `json:"custom_text4,omitempty"`
	CustomText5 *string `json:"custom_text5,omitempty"`
	CustomText6 *string `json:"custom_text6,omitempty"`
}

type _EsgTopics EsgTopics

// NewEsgTopics instantiates a new EsgTopics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsgTopics(name string) *EsgTopics {
	this := EsgTopics{}
	this.Name = name
	var scope string = "Not In Scope"
	this.Scope = &scope
	return &this
}

// NewEsgTopicsWithDefaults instantiates a new EsgTopics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsgTopicsWithDefaults() *EsgTopics {
	this := EsgTopics{}
	var scope string = "Not In Scope"
	this.Scope = &scope
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EsgTopics) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EsgTopics) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EsgTopics) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EsgTopics) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EsgTopics) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *EsgTopics) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EsgTopics) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EsgTopics) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *EsgTopics) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *EsgTopics) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *EsgTopics) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *EsgTopics) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetName returns the Name field value
func (o *EsgTopics) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EsgTopics) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EsgTopics) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EsgTopics) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EsgTopics) SetDescription(v string) {
	o.Description = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *EsgTopics) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *EsgTopics) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *EsgTopics) SetUid(v string) {
	o.Uid = &v
}

// GetFieldData returns the FieldData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EsgTopics) GetFieldData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FieldData
}

// GetFieldDataOk returns a tuple with the FieldData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EsgTopics) GetFieldDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FieldData) {
		return nil, false
	}
	return &o.FieldData, true
}

// HasFieldData returns a boolean if a field has been set.
func (o *EsgTopics) HasFieldData() bool {
	if o != nil && !IsNil(o.FieldData) {
		return true
	}

	return false
}

// SetFieldData gets a reference to the given interface{} and assigns it to the FieldData field.
func (o *EsgTopics) SetFieldData(v interface{}) {
	o.FieldData = v
}

// GetReviewerUserId returns the ReviewerUserId field value if set, zero value otherwise.
func (o *EsgTopics) GetReviewerUserId() int32 {
	if o == nil || IsNil(o.ReviewerUserId) {
		var ret int32
		return ret
	}
	return *o.ReviewerUserId
}

// GetReviewerUserIdOk returns a tuple with the ReviewerUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetReviewerUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ReviewerUserId) {
		return nil, false
	}
	return o.ReviewerUserId, true
}

// HasReviewerUserId returns a boolean if a field has been set.
func (o *EsgTopics) HasReviewerUserId() bool {
	if o != nil && !IsNil(o.ReviewerUserId) {
		return true
	}

	return false
}

// SetReviewerUserId gets a reference to the given int32 and assigns it to the ReviewerUserId field.
func (o *EsgTopics) SetReviewerUserId(v int32) {
	o.ReviewerUserId = &v
}

// GetEffectiveDate returns the EffectiveDate field value if set, zero value otherwise.
func (o *EsgTopics) GetEffectiveDate() string {
	if o == nil || IsNil(o.EffectiveDate) {
		var ret string
		return ret
	}
	return *o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetEffectiveDateOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveDate) {
		return nil, false
	}
	return o.EffectiveDate, true
}

// HasEffectiveDate returns a boolean if a field has been set.
func (o *EsgTopics) HasEffectiveDate() bool {
	if o != nil && !IsNil(o.EffectiveDate) {
		return true
	}

	return false
}

// SetEffectiveDate gets a reference to the given string and assigns it to the EffectiveDate field.
func (o *EsgTopics) SetEffectiveDate(v string) {
	o.EffectiveDate = &v
}

// GetOrganizationImpact returns the OrganizationImpact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EsgTopics) GetOrganizationImpact() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.OrganizationImpact
}

// GetOrganizationImpactOk returns a tuple with the OrganizationImpact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EsgTopics) GetOrganizationImpactOk() (*interface{}, bool) {
	if o == nil || IsNil(o.OrganizationImpact) {
		return nil, false
	}
	return &o.OrganizationImpact, true
}

// HasOrganizationImpact returns a boolean if a field has been set.
func (o *EsgTopics) HasOrganizationImpact() bool {
	if o != nil && !IsNil(o.OrganizationImpact) {
		return true
	}

	return false
}

// SetOrganizationImpact gets a reference to the given interface{} and assigns it to the OrganizationImpact field.
func (o *EsgTopics) SetOrganizationImpact(v interface{}) {
	o.OrganizationImpact = v
}

// GetStakeholderImpact returns the StakeholderImpact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EsgTopics) GetStakeholderImpact() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StakeholderImpact
}

// GetStakeholderImpactOk returns a tuple with the StakeholderImpact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EsgTopics) GetStakeholderImpactOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StakeholderImpact) {
		return nil, false
	}
	return &o.StakeholderImpact, true
}

// HasStakeholderImpact returns a boolean if a field has been set.
func (o *EsgTopics) HasStakeholderImpact() bool {
	if o != nil && !IsNil(o.StakeholderImpact) {
		return true
	}

	return false
}

// SetStakeholderImpact gets a reference to the given interface{} and assigns it to the StakeholderImpact field.
func (o *EsgTopics) SetStakeholderImpact(v interface{}) {
	o.StakeholderImpact = v
}

// GetMaterialityRating returns the MaterialityRating field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EsgTopics) GetMaterialityRating() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MaterialityRating
}

// GetMaterialityRatingOk returns a tuple with the MaterialityRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EsgTopics) GetMaterialityRatingOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MaterialityRating) {
		return nil, false
	}
	return &o.MaterialityRating, true
}

// HasMaterialityRating returns a boolean if a field has been set.
func (o *EsgTopics) HasMaterialityRating() bool {
	if o != nil && !IsNil(o.MaterialityRating) {
		return true
	}

	return false
}

// SetMaterialityRating gets a reference to the given interface{} and assigns it to the MaterialityRating field.
func (o *EsgTopics) SetMaterialityRating(v interface{}) {
	o.MaterialityRating = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *EsgTopics) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *EsgTopics) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *EsgTopics) SetCategory(v string) {
	o.Category = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *EsgTopics) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *EsgTopics) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *EsgTopics) SetScope(v string) {
	o.Scope = &v
}

// GetEsgCategoryId returns the EsgCategoryId field value if set, zero value otherwise.
func (o *EsgTopics) GetEsgCategoryId() int32 {
	if o == nil || IsNil(o.EsgCategoryId) {
		var ret int32
		return ret
	}
	return *o.EsgCategoryId
}

// GetEsgCategoryIdOk returns a tuple with the EsgCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetEsgCategoryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EsgCategoryId) {
		return nil, false
	}
	return o.EsgCategoryId, true
}

// HasEsgCategoryId returns a boolean if a field has been set.
func (o *EsgTopics) HasEsgCategoryId() bool {
	if o != nil && !IsNil(o.EsgCategoryId) {
		return true
	}

	return false
}

// SetEsgCategoryId gets a reference to the given int32 and assigns it to the EsgCategoryId field.
func (o *EsgTopics) SetEsgCategoryId(v int32) {
	o.EsgCategoryId = &v
}

// GetEsgTopicCustomSelect1OptionId returns the EsgTopicCustomSelect1OptionId field value if set, zero value otherwise.
func (o *EsgTopics) GetEsgTopicCustomSelect1OptionId() int32 {
	if o == nil || IsNil(o.EsgTopicCustomSelect1OptionId) {
		var ret int32
		return ret
	}
	return *o.EsgTopicCustomSelect1OptionId
}

// GetEsgTopicCustomSelect1OptionIdOk returns a tuple with the EsgTopicCustomSelect1OptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetEsgTopicCustomSelect1OptionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EsgTopicCustomSelect1OptionId) {
		return nil, false
	}
	return o.EsgTopicCustomSelect1OptionId, true
}

// HasEsgTopicCustomSelect1OptionId returns a boolean if a field has been set.
func (o *EsgTopics) HasEsgTopicCustomSelect1OptionId() bool {
	if o != nil && !IsNil(o.EsgTopicCustomSelect1OptionId) {
		return true
	}

	return false
}

// SetEsgTopicCustomSelect1OptionId gets a reference to the given int32 and assigns it to the EsgTopicCustomSelect1OptionId field.
func (o *EsgTopics) SetEsgTopicCustomSelect1OptionId(v int32) {
	o.EsgTopicCustomSelect1OptionId = &v
}

// GetEsgTopicCustomSelect2OptionId returns the EsgTopicCustomSelect2OptionId field value if set, zero value otherwise.
func (o *EsgTopics) GetEsgTopicCustomSelect2OptionId() int32 {
	if o == nil || IsNil(o.EsgTopicCustomSelect2OptionId) {
		var ret int32
		return ret
	}
	return *o.EsgTopicCustomSelect2OptionId
}

// GetEsgTopicCustomSelect2OptionIdOk returns a tuple with the EsgTopicCustomSelect2OptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetEsgTopicCustomSelect2OptionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EsgTopicCustomSelect2OptionId) {
		return nil, false
	}
	return o.EsgTopicCustomSelect2OptionId, true
}

// HasEsgTopicCustomSelect2OptionId returns a boolean if a field has been set.
func (o *EsgTopics) HasEsgTopicCustomSelect2OptionId() bool {
	if o != nil && !IsNil(o.EsgTopicCustomSelect2OptionId) {
		return true
	}

	return false
}

// SetEsgTopicCustomSelect2OptionId gets a reference to the given int32 and assigns it to the EsgTopicCustomSelect2OptionId field.
func (o *EsgTopics) SetEsgTopicCustomSelect2OptionId(v int32) {
	o.EsgTopicCustomSelect2OptionId = &v
}

// GetEsgTopicCustomSelect3OptionId returns the EsgTopicCustomSelect3OptionId field value if set, zero value otherwise.
func (o *EsgTopics) GetEsgTopicCustomSelect3OptionId() int32 {
	if o == nil || IsNil(o.EsgTopicCustomSelect3OptionId) {
		var ret int32
		return ret
	}
	return *o.EsgTopicCustomSelect3OptionId
}

// GetEsgTopicCustomSelect3OptionIdOk returns a tuple with the EsgTopicCustomSelect3OptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetEsgTopicCustomSelect3OptionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EsgTopicCustomSelect3OptionId) {
		return nil, false
	}
	return o.EsgTopicCustomSelect3OptionId, true
}

// HasEsgTopicCustomSelect3OptionId returns a boolean if a field has been set.
func (o *EsgTopics) HasEsgTopicCustomSelect3OptionId() bool {
	if o != nil && !IsNil(o.EsgTopicCustomSelect3OptionId) {
		return true
	}

	return false
}

// SetEsgTopicCustomSelect3OptionId gets a reference to the given int32 and assigns it to the EsgTopicCustomSelect3OptionId field.
func (o *EsgTopics) SetEsgTopicCustomSelect3OptionId(v int32) {
	o.EsgTopicCustomSelect3OptionId = &v
}

// GetEsgTopicCustomSelect4OptionId returns the EsgTopicCustomSelect4OptionId field value if set, zero value otherwise.
func (o *EsgTopics) GetEsgTopicCustomSelect4OptionId() int32 {
	if o == nil || IsNil(o.EsgTopicCustomSelect4OptionId) {
		var ret int32
		return ret
	}
	return *o.EsgTopicCustomSelect4OptionId
}

// GetEsgTopicCustomSelect4OptionIdOk returns a tuple with the EsgTopicCustomSelect4OptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetEsgTopicCustomSelect4OptionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EsgTopicCustomSelect4OptionId) {
		return nil, false
	}
	return o.EsgTopicCustomSelect4OptionId, true
}

// HasEsgTopicCustomSelect4OptionId returns a boolean if a field has been set.
func (o *EsgTopics) HasEsgTopicCustomSelect4OptionId() bool {
	if o != nil && !IsNil(o.EsgTopicCustomSelect4OptionId) {
		return true
	}

	return false
}

// SetEsgTopicCustomSelect4OptionId gets a reference to the given int32 and assigns it to the EsgTopicCustomSelect4OptionId field.
func (o *EsgTopics) SetEsgTopicCustomSelect4OptionId(v int32) {
	o.EsgTopicCustomSelect4OptionId = &v
}

// GetEsgTopicCustomSelect5OptionId returns the EsgTopicCustomSelect5OptionId field value if set, zero value otherwise.
func (o *EsgTopics) GetEsgTopicCustomSelect5OptionId() int32 {
	if o == nil || IsNil(o.EsgTopicCustomSelect5OptionId) {
		var ret int32
		return ret
	}
	return *o.EsgTopicCustomSelect5OptionId
}

// GetEsgTopicCustomSelect5OptionIdOk returns a tuple with the EsgTopicCustomSelect5OptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetEsgTopicCustomSelect5OptionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EsgTopicCustomSelect5OptionId) {
		return nil, false
	}
	return o.EsgTopicCustomSelect5OptionId, true
}

// HasEsgTopicCustomSelect5OptionId returns a boolean if a field has been set.
func (o *EsgTopics) HasEsgTopicCustomSelect5OptionId() bool {
	if o != nil && !IsNil(o.EsgTopicCustomSelect5OptionId) {
		return true
	}

	return false
}

// SetEsgTopicCustomSelect5OptionId gets a reference to the given int32 and assigns it to the EsgTopicCustomSelect5OptionId field.
func (o *EsgTopics) SetEsgTopicCustomSelect5OptionId(v int32) {
	o.EsgTopicCustomSelect5OptionId = &v
}

// GetCustomDate1 returns the CustomDate1 field value if set, zero value otherwise.
func (o *EsgTopics) GetCustomDate1() string {
	if o == nil || IsNil(o.CustomDate1) {
		var ret string
		return ret
	}
	return *o.CustomDate1
}

// GetCustomDate1Ok returns a tuple with the CustomDate1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetCustomDate1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomDate1) {
		return nil, false
	}
	return o.CustomDate1, true
}

// HasCustomDate1 returns a boolean if a field has been set.
func (o *EsgTopics) HasCustomDate1() bool {
	if o != nil && !IsNil(o.CustomDate1) {
		return true
	}

	return false
}

// SetCustomDate1 gets a reference to the given string and assigns it to the CustomDate1 field.
func (o *EsgTopics) SetCustomDate1(v string) {
	o.CustomDate1 = &v
}

// GetCustomDate2 returns the CustomDate2 field value if set, zero value otherwise.
func (o *EsgTopics) GetCustomDate2() string {
	if o == nil || IsNil(o.CustomDate2) {
		var ret string
		return ret
	}
	return *o.CustomDate2
}

// GetCustomDate2Ok returns a tuple with the CustomDate2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetCustomDate2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomDate2) {
		return nil, false
	}
	return o.CustomDate2, true
}

// HasCustomDate2 returns a boolean if a field has been set.
func (o *EsgTopics) HasCustomDate2() bool {
	if o != nil && !IsNil(o.CustomDate2) {
		return true
	}

	return false
}

// SetCustomDate2 gets a reference to the given string and assigns it to the CustomDate2 field.
func (o *EsgTopics) SetCustomDate2(v string) {
	o.CustomDate2 = &v
}

// GetCustomDate3 returns the CustomDate3 field value if set, zero value otherwise.
func (o *EsgTopics) GetCustomDate3() string {
	if o == nil || IsNil(o.CustomDate3) {
		var ret string
		return ret
	}
	return *o.CustomDate3
}

// GetCustomDate3Ok returns a tuple with the CustomDate3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetCustomDate3Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomDate3) {
		return nil, false
	}
	return o.CustomDate3, true
}

// HasCustomDate3 returns a boolean if a field has been set.
func (o *EsgTopics) HasCustomDate3() bool {
	if o != nil && !IsNil(o.CustomDate3) {
		return true
	}

	return false
}

// SetCustomDate3 gets a reference to the given string and assigns it to the CustomDate3 field.
func (o *EsgTopics) SetCustomDate3(v string) {
	o.CustomDate3 = &v
}

// GetCustomText1 returns the CustomText1 field value if set, zero value otherwise.
func (o *EsgTopics) GetCustomText1() string {
	if o == nil || IsNil(o.CustomText1) {
		var ret string
		return ret
	}
	return *o.CustomText1
}

// GetCustomText1Ok returns a tuple with the CustomText1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetCustomText1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomText1) {
		return nil, false
	}
	return o.CustomText1, true
}

// HasCustomText1 returns a boolean if a field has been set.
func (o *EsgTopics) HasCustomText1() bool {
	if o != nil && !IsNil(o.CustomText1) {
		return true
	}

	return false
}

// SetCustomText1 gets a reference to the given string and assigns it to the CustomText1 field.
func (o *EsgTopics) SetCustomText1(v string) {
	o.CustomText1 = &v
}

// GetCustomText2 returns the CustomText2 field value if set, zero value otherwise.
func (o *EsgTopics) GetCustomText2() string {
	if o == nil || IsNil(o.CustomText2) {
		var ret string
		return ret
	}
	return *o.CustomText2
}

// GetCustomText2Ok returns a tuple with the CustomText2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetCustomText2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomText2) {
		return nil, false
	}
	return o.CustomText2, true
}

// HasCustomText2 returns a boolean if a field has been set.
func (o *EsgTopics) HasCustomText2() bool {
	if o != nil && !IsNil(o.CustomText2) {
		return true
	}

	return false
}

// SetCustomText2 gets a reference to the given string and assigns it to the CustomText2 field.
func (o *EsgTopics) SetCustomText2(v string) {
	o.CustomText2 = &v
}

// GetCustomText3 returns the CustomText3 field value if set, zero value otherwise.
func (o *EsgTopics) GetCustomText3() string {
	if o == nil || IsNil(o.CustomText3) {
		var ret string
		return ret
	}
	return *o.CustomText3
}

// GetCustomText3Ok returns a tuple with the CustomText3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetCustomText3Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomText3) {
		return nil, false
	}
	return o.CustomText3, true
}

// HasCustomText3 returns a boolean if a field has been set.
func (o *EsgTopics) HasCustomText3() bool {
	if o != nil && !IsNil(o.CustomText3) {
		return true
	}

	return false
}

// SetCustomText3 gets a reference to the given string and assigns it to the CustomText3 field.
func (o *EsgTopics) SetCustomText3(v string) {
	o.CustomText3 = &v
}

// GetCustomText4 returns the CustomText4 field value if set, zero value otherwise.
func (o *EsgTopics) GetCustomText4() string {
	if o == nil || IsNil(o.CustomText4) {
		var ret string
		return ret
	}
	return *o.CustomText4
}

// GetCustomText4Ok returns a tuple with the CustomText4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetCustomText4Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomText4) {
		return nil, false
	}
	return o.CustomText4, true
}

// HasCustomText4 returns a boolean if a field has been set.
func (o *EsgTopics) HasCustomText4() bool {
	if o != nil && !IsNil(o.CustomText4) {
		return true
	}

	return false
}

// SetCustomText4 gets a reference to the given string and assigns it to the CustomText4 field.
func (o *EsgTopics) SetCustomText4(v string) {
	o.CustomText4 = &v
}

// GetCustomText5 returns the CustomText5 field value if set, zero value otherwise.
func (o *EsgTopics) GetCustomText5() string {
	if o == nil || IsNil(o.CustomText5) {
		var ret string
		return ret
	}
	return *o.CustomText5
}

// GetCustomText5Ok returns a tuple with the CustomText5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetCustomText5Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomText5) {
		return nil, false
	}
	return o.CustomText5, true
}

// HasCustomText5 returns a boolean if a field has been set.
func (o *EsgTopics) HasCustomText5() bool {
	if o != nil && !IsNil(o.CustomText5) {
		return true
	}

	return false
}

// SetCustomText5 gets a reference to the given string and assigns it to the CustomText5 field.
func (o *EsgTopics) SetCustomText5(v string) {
	o.CustomText5 = &v
}

// GetCustomText6 returns the CustomText6 field value if set, zero value otherwise.
func (o *EsgTopics) GetCustomText6() string {
	if o == nil || IsNil(o.CustomText6) {
		var ret string
		return ret
	}
	return *o.CustomText6
}

// GetCustomText6Ok returns a tuple with the CustomText6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgTopics) GetCustomText6Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomText6) {
		return nil, false
	}
	return o.CustomText6, true
}

// HasCustomText6 returns a boolean if a field has been set.
func (o *EsgTopics) HasCustomText6() bool {
	if o != nil && !IsNil(o.CustomText6) {
		return true
	}

	return false
}

// SetCustomText6 gets a reference to the given string and assigns it to the CustomText6 field.
func (o *EsgTopics) SetCustomText6(v string) {
	o.CustomText6 = &v
}

func (o EsgTopics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsgTopics) ToMap() (map[string]interface{}, error) {
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
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	if o.FieldData != nil {
		toSerialize["field_data"] = o.FieldData
	}
	if !IsNil(o.ReviewerUserId) {
		toSerialize["reviewer_user_id"] = o.ReviewerUserId
	}
	if !IsNil(o.EffectiveDate) {
		toSerialize["effective_date"] = o.EffectiveDate
	}
	if o.OrganizationImpact != nil {
		toSerialize["organization_impact"] = o.OrganizationImpact
	}
	if o.StakeholderImpact != nil {
		toSerialize["stakeholder_impact"] = o.StakeholderImpact
	}
	if o.MaterialityRating != nil {
		toSerialize["materiality_rating"] = o.MaterialityRating
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.EsgCategoryId) {
		toSerialize["esg_category_id"] = o.EsgCategoryId
	}
	if !IsNil(o.EsgTopicCustomSelect1OptionId) {
		toSerialize["esg_topic_custom_select1_option_id"] = o.EsgTopicCustomSelect1OptionId
	}
	if !IsNil(o.EsgTopicCustomSelect2OptionId) {
		toSerialize["esg_topic_custom_select2_option_id"] = o.EsgTopicCustomSelect2OptionId
	}
	if !IsNil(o.EsgTopicCustomSelect3OptionId) {
		toSerialize["esg_topic_custom_select3_option_id"] = o.EsgTopicCustomSelect3OptionId
	}
	if !IsNil(o.EsgTopicCustomSelect4OptionId) {
		toSerialize["esg_topic_custom_select4_option_id"] = o.EsgTopicCustomSelect4OptionId
	}
	if !IsNil(o.EsgTopicCustomSelect5OptionId) {
		toSerialize["esg_topic_custom_select5_option_id"] = o.EsgTopicCustomSelect5OptionId
	}
	if !IsNil(o.CustomDate1) {
		toSerialize["custom_date1"] = o.CustomDate1
	}
	if !IsNil(o.CustomDate2) {
		toSerialize["custom_date2"] = o.CustomDate2
	}
	if !IsNil(o.CustomDate3) {
		toSerialize["custom_date3"] = o.CustomDate3
	}
	if !IsNil(o.CustomText1) {
		toSerialize["custom_text1"] = o.CustomText1
	}
	if !IsNil(o.CustomText2) {
		toSerialize["custom_text2"] = o.CustomText2
	}
	if !IsNil(o.CustomText3) {
		toSerialize["custom_text3"] = o.CustomText3
	}
	if !IsNil(o.CustomText4) {
		toSerialize["custom_text4"] = o.CustomText4
	}
	if !IsNil(o.CustomText5) {
		toSerialize["custom_text5"] = o.CustomText5
	}
	if !IsNil(o.CustomText6) {
		toSerialize["custom_text6"] = o.CustomText6
	}
	return toSerialize, nil
}

type NullableEsgTopics struct {
	value *EsgTopics
	isSet bool
}

func (v NullableEsgTopics) Get() *EsgTopics {
	return v.value
}

func (v *NullableEsgTopics) Set(val *EsgTopics) {
	v.value = val
	v.isSet = true
}

func (v NullableEsgTopics) IsSet() bool {
	return v.isSet
}

func (v *NullableEsgTopics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsgTopics(val *EsgTopics) *NullableEsgTopics {
	return &NullableEsgTopics{value: val, isSet: true}
}

func (v NullableEsgTopics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsgTopics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


