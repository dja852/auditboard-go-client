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

// checks if the AuditSurveysPutAuditSurvey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditSurveysPutAuditSurvey{}

// AuditSurveysPutAuditSurvey struct for AuditSurveysPutAuditSurvey
type AuditSurveysPutAuditSurvey struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	AuditsurveyableId *int32 `json:"auditsurveyable_id,omitempty"`
	AuditsurveyableType *string `json:"auditsurveyable_type,omitempty"`
	Name string `json:"name"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	CreatedByUserId *int32 `json:"created_by_user_id,omitempty"`
	QuestionSortOrder interface{} `json:"question_sort_order,omitempty"`
	SourceId *int32 `json:"source_id,omitempty"`
	IsTemplate bool `json:"is_template"`
	ReferenceMeta interface{} `json:"reference_meta,omitempty"`
	Totals interface{} `json:"totals,omitempty"`
	QuestionOrder interface{} `json:"question_order"`
	Score interface{} `json:"score,omitempty"`
	Type *string `json:"type,omitempty"`
	// Note: This is a Foreign Key to `task_items.id`.<fk table='task_items' column='id'/>
	TaskItemId *int32 `json:"task_item_id,omitempty"`
	ScoringMethodology *string `json:"scoring_methodology,omitempty"`
	BracketsType *string `json:"brackets_type,omitempty"`
	Brackets interface{} `json:"brackets"`
	// Note: This is a Foreign Key to `audit_survey_templates.id`.<fk table='audit_survey_templates' column='id'/>
	AuditSurveyTemplateId *int32 `json:"audit_survey_template_id,omitempty"`
	SectionOrder interface{} `json:"section_order"`
	StashedQuestionOrder interface{} `json:"stashed_question_order,omitempty"`
	RecurringSettings interface{} `json:"recurring_settings,omitempty"`
	IsWeighted *bool `json:"is_weighted,omitempty"`
	// Note: This is a Foreign Key to `audit_survey_categories.id`.<fk table='audit_survey_categories' column='id'/>
	AuditSurveyCategoryId *int32 `json:"audit_survey_category_id,omitempty"`
	// Note: This is a Foreign Key to `audit_survey_scores.id`.<fk table='audit_survey_scores' column='id'/>
	AuditSurveyScoreId *int32 `json:"audit_survey_score_id,omitempty"`
}

type _AuditSurveysPutAuditSurvey AuditSurveysPutAuditSurvey

// NewAuditSurveysPutAuditSurvey instantiates a new AuditSurveysPutAuditSurvey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditSurveysPutAuditSurvey(name string, isTemplate bool, questionOrder interface{}, brackets interface{}, sectionOrder interface{}) *AuditSurveysPutAuditSurvey {
	this := AuditSurveysPutAuditSurvey{}
	this.Name = name
	this.IsTemplate = isTemplate
	this.QuestionOrder = questionOrder
	var bracketsType string = "NULL"
	this.BracketsType = &bracketsType
	this.Brackets = brackets
	this.SectionOrder = sectionOrder
	var isWeighted bool = false
	this.IsWeighted = &isWeighted
	return &this
}

// NewAuditSurveysPutAuditSurveyWithDefaults instantiates a new AuditSurveysPutAuditSurvey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditSurveysPutAuditSurveyWithDefaults() *AuditSurveysPutAuditSurvey {
	this := AuditSurveysPutAuditSurvey{}
	var isTemplate bool = false
	this.IsTemplate = isTemplate
	var bracketsType string = "NULL"
	this.BracketsType = &bracketsType
	var isWeighted bool = false
	this.IsWeighted = &isWeighted
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuditSurveysPutAuditSurvey) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveysPutAuditSurvey) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AuditSurveysPutAuditSurvey) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AuditSurveysPutAuditSurvey) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveysPutAuditSurvey) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AuditSurveysPutAuditSurvey) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AuditSurveysPutAuditSurvey) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveysPutAuditSurvey) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AuditSurveysPutAuditSurvey) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *AuditSurveysPutAuditSurvey) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveysPutAuditSurvey) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *AuditSurveysPutAuditSurvey) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetAuditsurveyableId returns the AuditsurveyableId field value if set, zero value otherwise.
func (o *AuditSurveysPutAuditSurvey) GetAuditsurveyableId() int32 {
	if o == nil || IsNil(o.AuditsurveyableId) {
		var ret int32
		return ret
	}
	return *o.AuditsurveyableId
}

// GetAuditsurveyableIdOk returns a tuple with the AuditsurveyableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveysPutAuditSurvey) GetAuditsurveyableIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AuditsurveyableId) {
		return nil, false
	}
	return o.AuditsurveyableId, true
}

// HasAuditsurveyableId returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasAuditsurveyableId() bool {
	if o != nil && !IsNil(o.AuditsurveyableId) {
		return true
	}

	return false
}

// SetAuditsurveyableId gets a reference to the given int32 and assigns it to the AuditsurveyableId field.
func (o *AuditSurveysPutAuditSurvey) SetAuditsurveyableId(v int32) {
	o.AuditsurveyableId = &v
}

// GetAuditsurveyableType returns the AuditsurveyableType field value if set, zero value otherwise.
func (o *AuditSurveysPutAuditSurvey) GetAuditsurveyableType() string {
	if o == nil || IsNil(o.AuditsurveyableType) {
		var ret string
		return ret
	}
	return *o.AuditsurveyableType
}

// GetAuditsurveyableTypeOk returns a tuple with the AuditsurveyableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveysPutAuditSurvey) GetAuditsurveyableTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AuditsurveyableType) {
		return nil, false
	}
	return o.AuditsurveyableType, true
}

// HasAuditsurveyableType returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasAuditsurveyableType() bool {
	if o != nil && !IsNil(o.AuditsurveyableType) {
		return true
	}

	return false
}

// SetAuditsurveyableType gets a reference to the given string and assigns it to the AuditsurveyableType field.
func (o *AuditSurveysPutAuditSurvey) SetAuditsurveyableType(v string) {
	o.AuditsurveyableType = &v
}

// GetName returns the Name field value
func (o *AuditSurveysPutAuditSurvey) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuditSurveysPutAuditSurvey) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuditSurveysPutAuditSurvey) SetName(v string) {
	o.Name = v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *AuditSurveysPutAuditSurvey) GetCreatedByUserId() int32 {
	if o == nil || IsNil(o.CreatedByUserId) {
		var ret int32
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveysPutAuditSurvey) GetCreatedByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedByUserId) {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasCreatedByUserId() bool {
	if o != nil && !IsNil(o.CreatedByUserId) {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given int32 and assigns it to the CreatedByUserId field.
func (o *AuditSurveysPutAuditSurvey) SetCreatedByUserId(v int32) {
	o.CreatedByUserId = &v
}

// GetQuestionSortOrder returns the QuestionSortOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditSurveysPutAuditSurvey) GetQuestionSortOrder() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.QuestionSortOrder
}

// GetQuestionSortOrderOk returns a tuple with the QuestionSortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditSurveysPutAuditSurvey) GetQuestionSortOrderOk() (*interface{}, bool) {
	if o == nil || IsNil(o.QuestionSortOrder) {
		return nil, false
	}
	return &o.QuestionSortOrder, true
}

// HasQuestionSortOrder returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasQuestionSortOrder() bool {
	if o != nil && !IsNil(o.QuestionSortOrder) {
		return true
	}

	return false
}

// SetQuestionSortOrder gets a reference to the given interface{} and assigns it to the QuestionSortOrder field.
func (o *AuditSurveysPutAuditSurvey) SetQuestionSortOrder(v interface{}) {
	o.QuestionSortOrder = v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *AuditSurveysPutAuditSurvey) GetSourceId() int32 {
	if o == nil || IsNil(o.SourceId) {
		var ret int32
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveysPutAuditSurvey) GetSourceIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasSourceId() bool {
	if o != nil && !IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given int32 and assigns it to the SourceId field.
func (o *AuditSurveysPutAuditSurvey) SetSourceId(v int32) {
	o.SourceId = &v
}

// GetIsTemplate returns the IsTemplate field value
func (o *AuditSurveysPutAuditSurvey) GetIsTemplate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsTemplate
}

// GetIsTemplateOk returns a tuple with the IsTemplate field value
// and a boolean to check if the value has been set.
func (o *AuditSurveysPutAuditSurvey) GetIsTemplateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTemplate, true
}

// SetIsTemplate sets field value
func (o *AuditSurveysPutAuditSurvey) SetIsTemplate(v bool) {
	o.IsTemplate = v
}

// GetReferenceMeta returns the ReferenceMeta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditSurveysPutAuditSurvey) GetReferenceMeta() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceMeta
}

// GetReferenceMetaOk returns a tuple with the ReferenceMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditSurveysPutAuditSurvey) GetReferenceMetaOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceMeta) {
		return nil, false
	}
	return &o.ReferenceMeta, true
}

// HasReferenceMeta returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasReferenceMeta() bool {
	if o != nil && !IsNil(o.ReferenceMeta) {
		return true
	}

	return false
}

// SetReferenceMeta gets a reference to the given interface{} and assigns it to the ReferenceMeta field.
func (o *AuditSurveysPutAuditSurvey) SetReferenceMeta(v interface{}) {
	o.ReferenceMeta = v
}

// GetTotals returns the Totals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditSurveysPutAuditSurvey) GetTotals() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Totals
}

// GetTotalsOk returns a tuple with the Totals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditSurveysPutAuditSurvey) GetTotalsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Totals) {
		return nil, false
	}
	return &o.Totals, true
}

// HasTotals returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasTotals() bool {
	if o != nil && !IsNil(o.Totals) {
		return true
	}

	return false
}

// SetTotals gets a reference to the given interface{} and assigns it to the Totals field.
func (o *AuditSurveysPutAuditSurvey) SetTotals(v interface{}) {
	o.Totals = v
}

// GetQuestionOrder returns the QuestionOrder field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AuditSurveysPutAuditSurvey) GetQuestionOrder() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.QuestionOrder
}

// GetQuestionOrderOk returns a tuple with the QuestionOrder field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditSurveysPutAuditSurvey) GetQuestionOrderOk() (*interface{}, bool) {
	if o == nil || IsNil(o.QuestionOrder) {
		return nil, false
	}
	return &o.QuestionOrder, true
}

// SetQuestionOrder sets field value
func (o *AuditSurveysPutAuditSurvey) SetQuestionOrder(v interface{}) {
	o.QuestionOrder = v
}

// GetScore returns the Score field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditSurveysPutAuditSurvey) GetScore() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditSurveysPutAuditSurvey) GetScoreOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return &o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given interface{} and assigns it to the Score field.
func (o *AuditSurveysPutAuditSurvey) SetScore(v interface{}) {
	o.Score = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuditSurveysPutAuditSurvey) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveysPutAuditSurvey) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AuditSurveysPutAuditSurvey) SetType(v string) {
	o.Type = &v
}

// GetTaskItemId returns the TaskItemId field value if set, zero value otherwise.
func (o *AuditSurveysPutAuditSurvey) GetTaskItemId() int32 {
	if o == nil || IsNil(o.TaskItemId) {
		var ret int32
		return ret
	}
	return *o.TaskItemId
}

// GetTaskItemIdOk returns a tuple with the TaskItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveysPutAuditSurvey) GetTaskItemIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TaskItemId) {
		return nil, false
	}
	return o.TaskItemId, true
}

// HasTaskItemId returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasTaskItemId() bool {
	if o != nil && !IsNil(o.TaskItemId) {
		return true
	}

	return false
}

// SetTaskItemId gets a reference to the given int32 and assigns it to the TaskItemId field.
func (o *AuditSurveysPutAuditSurvey) SetTaskItemId(v int32) {
	o.TaskItemId = &v
}

// GetScoringMethodology returns the ScoringMethodology field value if set, zero value otherwise.
func (o *AuditSurveysPutAuditSurvey) GetScoringMethodology() string {
	if o == nil || IsNil(o.ScoringMethodology) {
		var ret string
		return ret
	}
	return *o.ScoringMethodology
}

// GetScoringMethodologyOk returns a tuple with the ScoringMethodology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveysPutAuditSurvey) GetScoringMethodologyOk() (*string, bool) {
	if o == nil || IsNil(o.ScoringMethodology) {
		return nil, false
	}
	return o.ScoringMethodology, true
}

// HasScoringMethodology returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasScoringMethodology() bool {
	if o != nil && !IsNil(o.ScoringMethodology) {
		return true
	}

	return false
}

// SetScoringMethodology gets a reference to the given string and assigns it to the ScoringMethodology field.
func (o *AuditSurveysPutAuditSurvey) SetScoringMethodology(v string) {
	o.ScoringMethodology = &v
}

// GetBracketsType returns the BracketsType field value if set, zero value otherwise.
func (o *AuditSurveysPutAuditSurvey) GetBracketsType() string {
	if o == nil || IsNil(o.BracketsType) {
		var ret string
		return ret
	}
	return *o.BracketsType
}

// GetBracketsTypeOk returns a tuple with the BracketsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveysPutAuditSurvey) GetBracketsTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BracketsType) {
		return nil, false
	}
	return o.BracketsType, true
}

// HasBracketsType returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasBracketsType() bool {
	if o != nil && !IsNil(o.BracketsType) {
		return true
	}

	return false
}

// SetBracketsType gets a reference to the given string and assigns it to the BracketsType field.
func (o *AuditSurveysPutAuditSurvey) SetBracketsType(v string) {
	o.BracketsType = &v
}

// GetBrackets returns the Brackets field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AuditSurveysPutAuditSurvey) GetBrackets() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Brackets
}

// GetBracketsOk returns a tuple with the Brackets field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditSurveysPutAuditSurvey) GetBracketsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Brackets) {
		return nil, false
	}
	return &o.Brackets, true
}

// SetBrackets sets field value
func (o *AuditSurveysPutAuditSurvey) SetBrackets(v interface{}) {
	o.Brackets = v
}

// GetAuditSurveyTemplateId returns the AuditSurveyTemplateId field value if set, zero value otherwise.
func (o *AuditSurveysPutAuditSurvey) GetAuditSurveyTemplateId() int32 {
	if o == nil || IsNil(o.AuditSurveyTemplateId) {
		var ret int32
		return ret
	}
	return *o.AuditSurveyTemplateId
}

// GetAuditSurveyTemplateIdOk returns a tuple with the AuditSurveyTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveysPutAuditSurvey) GetAuditSurveyTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AuditSurveyTemplateId) {
		return nil, false
	}
	return o.AuditSurveyTemplateId, true
}

// HasAuditSurveyTemplateId returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasAuditSurveyTemplateId() bool {
	if o != nil && !IsNil(o.AuditSurveyTemplateId) {
		return true
	}

	return false
}

// SetAuditSurveyTemplateId gets a reference to the given int32 and assigns it to the AuditSurveyTemplateId field.
func (o *AuditSurveysPutAuditSurvey) SetAuditSurveyTemplateId(v int32) {
	o.AuditSurveyTemplateId = &v
}

// GetSectionOrder returns the SectionOrder field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AuditSurveysPutAuditSurvey) GetSectionOrder() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.SectionOrder
}

// GetSectionOrderOk returns a tuple with the SectionOrder field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditSurveysPutAuditSurvey) GetSectionOrderOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SectionOrder) {
		return nil, false
	}
	return &o.SectionOrder, true
}

// SetSectionOrder sets field value
func (o *AuditSurveysPutAuditSurvey) SetSectionOrder(v interface{}) {
	o.SectionOrder = v
}

// GetStashedQuestionOrder returns the StashedQuestionOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditSurveysPutAuditSurvey) GetStashedQuestionOrder() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StashedQuestionOrder
}

// GetStashedQuestionOrderOk returns a tuple with the StashedQuestionOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditSurveysPutAuditSurvey) GetStashedQuestionOrderOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StashedQuestionOrder) {
		return nil, false
	}
	return &o.StashedQuestionOrder, true
}

// HasStashedQuestionOrder returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasStashedQuestionOrder() bool {
	if o != nil && !IsNil(o.StashedQuestionOrder) {
		return true
	}

	return false
}

// SetStashedQuestionOrder gets a reference to the given interface{} and assigns it to the StashedQuestionOrder field.
func (o *AuditSurveysPutAuditSurvey) SetStashedQuestionOrder(v interface{}) {
	o.StashedQuestionOrder = v
}

// GetRecurringSettings returns the RecurringSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditSurveysPutAuditSurvey) GetRecurringSettings() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RecurringSettings
}

// GetRecurringSettingsOk returns a tuple with the RecurringSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditSurveysPutAuditSurvey) GetRecurringSettingsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RecurringSettings) {
		return nil, false
	}
	return &o.RecurringSettings, true
}

// HasRecurringSettings returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasRecurringSettings() bool {
	if o != nil && !IsNil(o.RecurringSettings) {
		return true
	}

	return false
}

// SetRecurringSettings gets a reference to the given interface{} and assigns it to the RecurringSettings field.
func (o *AuditSurveysPutAuditSurvey) SetRecurringSettings(v interface{}) {
	o.RecurringSettings = v
}

// GetIsWeighted returns the IsWeighted field value if set, zero value otherwise.
func (o *AuditSurveysPutAuditSurvey) GetIsWeighted() bool {
	if o == nil || IsNil(o.IsWeighted) {
		var ret bool
		return ret
	}
	return *o.IsWeighted
}

// GetIsWeightedOk returns a tuple with the IsWeighted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveysPutAuditSurvey) GetIsWeightedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsWeighted) {
		return nil, false
	}
	return o.IsWeighted, true
}

// HasIsWeighted returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasIsWeighted() bool {
	if o != nil && !IsNil(o.IsWeighted) {
		return true
	}

	return false
}

// SetIsWeighted gets a reference to the given bool and assigns it to the IsWeighted field.
func (o *AuditSurveysPutAuditSurvey) SetIsWeighted(v bool) {
	o.IsWeighted = &v
}

// GetAuditSurveyCategoryId returns the AuditSurveyCategoryId field value if set, zero value otherwise.
func (o *AuditSurveysPutAuditSurvey) GetAuditSurveyCategoryId() int32 {
	if o == nil || IsNil(o.AuditSurveyCategoryId) {
		var ret int32
		return ret
	}
	return *o.AuditSurveyCategoryId
}

// GetAuditSurveyCategoryIdOk returns a tuple with the AuditSurveyCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveysPutAuditSurvey) GetAuditSurveyCategoryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AuditSurveyCategoryId) {
		return nil, false
	}
	return o.AuditSurveyCategoryId, true
}

// HasAuditSurveyCategoryId returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasAuditSurveyCategoryId() bool {
	if o != nil && !IsNil(o.AuditSurveyCategoryId) {
		return true
	}

	return false
}

// SetAuditSurveyCategoryId gets a reference to the given int32 and assigns it to the AuditSurveyCategoryId field.
func (o *AuditSurveysPutAuditSurvey) SetAuditSurveyCategoryId(v int32) {
	o.AuditSurveyCategoryId = &v
}

// GetAuditSurveyScoreId returns the AuditSurveyScoreId field value if set, zero value otherwise.
func (o *AuditSurveysPutAuditSurvey) GetAuditSurveyScoreId() int32 {
	if o == nil || IsNil(o.AuditSurveyScoreId) {
		var ret int32
		return ret
	}
	return *o.AuditSurveyScoreId
}

// GetAuditSurveyScoreIdOk returns a tuple with the AuditSurveyScoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSurveysPutAuditSurvey) GetAuditSurveyScoreIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AuditSurveyScoreId) {
		return nil, false
	}
	return o.AuditSurveyScoreId, true
}

// HasAuditSurveyScoreId returns a boolean if a field has been set.
func (o *AuditSurveysPutAuditSurvey) HasAuditSurveyScoreId() bool {
	if o != nil && !IsNil(o.AuditSurveyScoreId) {
		return true
	}

	return false
}

// SetAuditSurveyScoreId gets a reference to the given int32 and assigns it to the AuditSurveyScoreId field.
func (o *AuditSurveysPutAuditSurvey) SetAuditSurveyScoreId(v int32) {
	o.AuditSurveyScoreId = &v
}

func (o AuditSurveysPutAuditSurvey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditSurveysPutAuditSurvey) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AuditsurveyableId) {
		toSerialize["auditsurveyable_id"] = o.AuditsurveyableId
	}
	if !IsNil(o.AuditsurveyableType) {
		toSerialize["auditsurveyable_type"] = o.AuditsurveyableType
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.CreatedByUserId) {
		toSerialize["created_by_user_id"] = o.CreatedByUserId
	}
	if o.QuestionSortOrder != nil {
		toSerialize["question_sort_order"] = o.QuestionSortOrder
	}
	if !IsNil(o.SourceId) {
		toSerialize["source_id"] = o.SourceId
	}
	toSerialize["is_template"] = o.IsTemplate
	if o.ReferenceMeta != nil {
		toSerialize["reference_meta"] = o.ReferenceMeta
	}
	if o.Totals != nil {
		toSerialize["totals"] = o.Totals
	}
	if o.QuestionOrder != nil {
		toSerialize["question_order"] = o.QuestionOrder
	}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.TaskItemId) {
		toSerialize["task_item_id"] = o.TaskItemId
	}
	if !IsNil(o.ScoringMethodology) {
		toSerialize["scoring_methodology"] = o.ScoringMethodology
	}
	if !IsNil(o.BracketsType) {
		toSerialize["brackets_type"] = o.BracketsType
	}
	if o.Brackets != nil {
		toSerialize["brackets"] = o.Brackets
	}
	if !IsNil(o.AuditSurveyTemplateId) {
		toSerialize["audit_survey_template_id"] = o.AuditSurveyTemplateId
	}
	if o.SectionOrder != nil {
		toSerialize["section_order"] = o.SectionOrder
	}
	if o.StashedQuestionOrder != nil {
		toSerialize["stashed_question_order"] = o.StashedQuestionOrder
	}
	if o.RecurringSettings != nil {
		toSerialize["recurring_settings"] = o.RecurringSettings
	}
	if !IsNil(o.IsWeighted) {
		toSerialize["is_weighted"] = o.IsWeighted
	}
	if !IsNil(o.AuditSurveyCategoryId) {
		toSerialize["audit_survey_category_id"] = o.AuditSurveyCategoryId
	}
	if !IsNil(o.AuditSurveyScoreId) {
		toSerialize["audit_survey_score_id"] = o.AuditSurveyScoreId
	}
	return toSerialize, nil
}

func (o *AuditSurveysPutAuditSurvey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"is_template",
		"question_order",
		"brackets",
		"section_order",
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

	varAuditSurveysPutAuditSurvey := _AuditSurveysPutAuditSurvey{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuditSurveysPutAuditSurvey)

	if err != nil {
		return err
	}

	*o = AuditSurveysPutAuditSurvey(varAuditSurveysPutAuditSurvey)

	return err
}

type NullableAuditSurveysPutAuditSurvey struct {
	value *AuditSurveysPutAuditSurvey
	isSet bool
}

func (v NullableAuditSurveysPutAuditSurvey) Get() *AuditSurveysPutAuditSurvey {
	return v.value
}

func (v *NullableAuditSurveysPutAuditSurvey) Set(val *AuditSurveysPutAuditSurvey) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditSurveysPutAuditSurvey) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditSurveysPutAuditSurvey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditSurveysPutAuditSurvey(val *AuditSurveysPutAuditSurvey) *NullableAuditSurveysPutAuditSurvey {
	return &NullableAuditSurveysPutAuditSurvey{value: val, isSet: true}
}

func (v NullableAuditSurveysPutAuditSurvey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditSurveysPutAuditSurvey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


