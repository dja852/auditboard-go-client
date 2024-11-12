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

// checks if the AuditQuestionsPutPreviousValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditQuestionsPutPreviousValues{}

// AuditQuestionsPutPreviousValues struct for AuditQuestionsPutPreviousValues
type AuditQuestionsPutPreviousValues struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Requirements interface{} `json:"requirements,omitempty"`
	Type *string `json:"type,omitempty"`
	QuestionData interface{} `json:"question_data,omitempty"`
	Text *string `json:"text,omitempty"`
	Description *string `json:"description,omitempty"`
	SourceId *int32 `json:"source_id,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	CreatedByUserId *int32 `json:"created_by_user_id,omitempty"`
	ConditionChoiceUuid *string `json:"condition_choice_uuid,omitempty"`
	Uid *string `json:"uid,omitempty"`
	ChangeableType *string `json:"changeable_type,omitempty"`
	ManagedChangeField *string `json:"managed_change_field,omitempty"`
	ManagedChangeFieldType *string `json:"managed_change_field_type,omitempty"`
	QuestionOptions interface{} `json:"question_options,omitempty"`
	AllowNaResponse *bool `json:"allow_na_response,omitempty"`
	AllowExplanation *bool `json:"allow_explanation,omitempty"`
	AllowFiles *bool `json:"allow_files,omitempty"`
	RequireFiles *bool `json:"require_files,omitempty"`
	QuestionUuid *string `json:"question_uuid,omitempty"`
	ConditionQuestionUuid *string `json:"condition_question_uuid,omitempty"`
	Condition interface{} `json:"condition,omitempty"`
	RequireNaExplanation *bool `json:"require_na_explanation,omitempty"`
}

// NewAuditQuestionsPutPreviousValues instantiates a new AuditQuestionsPutPreviousValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditQuestionsPutPreviousValues() *AuditQuestionsPutPreviousValues {
	this := AuditQuestionsPutPreviousValues{}
	var allowNaResponse bool = false
	this.AllowNaResponse = &allowNaResponse
	var allowExplanation bool = false
	this.AllowExplanation = &allowExplanation
	var allowFiles bool = false
	this.AllowFiles = &allowFiles
	var requireFiles bool = false
	this.RequireFiles = &requireFiles
	var requireNaExplanation bool = false
	this.RequireNaExplanation = &requireNaExplanation
	return &this
}

// NewAuditQuestionsPutPreviousValuesWithDefaults instantiates a new AuditQuestionsPutPreviousValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditQuestionsPutPreviousValuesWithDefaults() *AuditQuestionsPutPreviousValues {
	this := AuditQuestionsPutPreviousValues{}
	var allowNaResponse bool = false
	this.AllowNaResponse = &allowNaResponse
	var allowExplanation bool = false
	this.AllowExplanation = &allowExplanation
	var allowFiles bool = false
	this.AllowFiles = &allowFiles
	var requireFiles bool = false
	this.RequireFiles = &requireFiles
	var requireNaExplanation bool = false
	this.RequireNaExplanation = &requireNaExplanation
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuditQuestionsPutPreviousValues) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionsPutPreviousValues) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AuditQuestionsPutPreviousValues) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AuditQuestionsPutPreviousValues) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionsPutPreviousValues) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AuditQuestionsPutPreviousValues) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AuditQuestionsPutPreviousValues) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionsPutPreviousValues) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AuditQuestionsPutPreviousValues) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *AuditQuestionsPutPreviousValues) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionsPutPreviousValues) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *AuditQuestionsPutPreviousValues) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetRequirements returns the Requirements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditQuestionsPutPreviousValues) GetRequirements() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditQuestionsPutPreviousValues) GetRequirementsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Requirements) {
		return nil, false
	}
	return &o.Requirements, true
}

// HasRequirements returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasRequirements() bool {
	if o != nil && !IsNil(o.Requirements) {
		return true
	}

	return false
}

// SetRequirements gets a reference to the given interface{} and assigns it to the Requirements field.
func (o *AuditQuestionsPutPreviousValues) SetRequirements(v interface{}) {
	o.Requirements = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuditQuestionsPutPreviousValues) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionsPutPreviousValues) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AuditQuestionsPutPreviousValues) SetType(v string) {
	o.Type = &v
}

// GetQuestionData returns the QuestionData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditQuestionsPutPreviousValues) GetQuestionData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.QuestionData
}

// GetQuestionDataOk returns a tuple with the QuestionData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditQuestionsPutPreviousValues) GetQuestionDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.QuestionData) {
		return nil, false
	}
	return &o.QuestionData, true
}

// HasQuestionData returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasQuestionData() bool {
	if o != nil && !IsNil(o.QuestionData) {
		return true
	}

	return false
}

// SetQuestionData gets a reference to the given interface{} and assigns it to the QuestionData field.
func (o *AuditQuestionsPutPreviousValues) SetQuestionData(v interface{}) {
	o.QuestionData = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *AuditQuestionsPutPreviousValues) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionsPutPreviousValues) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *AuditQuestionsPutPreviousValues) SetText(v string) {
	o.Text = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuditQuestionsPutPreviousValues) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionsPutPreviousValues) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuditQuestionsPutPreviousValues) SetDescription(v string) {
	o.Description = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *AuditQuestionsPutPreviousValues) GetSourceId() int32 {
	if o == nil || IsNil(o.SourceId) {
		var ret int32
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionsPutPreviousValues) GetSourceIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasSourceId() bool {
	if o != nil && !IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given int32 and assigns it to the SourceId field.
func (o *AuditQuestionsPutPreviousValues) SetSourceId(v int32) {
	o.SourceId = &v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *AuditQuestionsPutPreviousValues) GetCreatedByUserId() int32 {
	if o == nil || IsNil(o.CreatedByUserId) {
		var ret int32
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionsPutPreviousValues) GetCreatedByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedByUserId) {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasCreatedByUserId() bool {
	if o != nil && !IsNil(o.CreatedByUserId) {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given int32 and assigns it to the CreatedByUserId field.
func (o *AuditQuestionsPutPreviousValues) SetCreatedByUserId(v int32) {
	o.CreatedByUserId = &v
}

// GetConditionChoiceUuid returns the ConditionChoiceUuid field value if set, zero value otherwise.
func (o *AuditQuestionsPutPreviousValues) GetConditionChoiceUuid() string {
	if o == nil || IsNil(o.ConditionChoiceUuid) {
		var ret string
		return ret
	}
	return *o.ConditionChoiceUuid
}

// GetConditionChoiceUuidOk returns a tuple with the ConditionChoiceUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionsPutPreviousValues) GetConditionChoiceUuidOk() (*string, bool) {
	if o == nil || IsNil(o.ConditionChoiceUuid) {
		return nil, false
	}
	return o.ConditionChoiceUuid, true
}

// HasConditionChoiceUuid returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasConditionChoiceUuid() bool {
	if o != nil && !IsNil(o.ConditionChoiceUuid) {
		return true
	}

	return false
}

// SetConditionChoiceUuid gets a reference to the given string and assigns it to the ConditionChoiceUuid field.
func (o *AuditQuestionsPutPreviousValues) SetConditionChoiceUuid(v string) {
	o.ConditionChoiceUuid = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *AuditQuestionsPutPreviousValues) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionsPutPreviousValues) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *AuditQuestionsPutPreviousValues) SetUid(v string) {
	o.Uid = &v
}

// GetChangeableType returns the ChangeableType field value if set, zero value otherwise.
func (o *AuditQuestionsPutPreviousValues) GetChangeableType() string {
	if o == nil || IsNil(o.ChangeableType) {
		var ret string
		return ret
	}
	return *o.ChangeableType
}

// GetChangeableTypeOk returns a tuple with the ChangeableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionsPutPreviousValues) GetChangeableTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ChangeableType) {
		return nil, false
	}
	return o.ChangeableType, true
}

// HasChangeableType returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasChangeableType() bool {
	if o != nil && !IsNil(o.ChangeableType) {
		return true
	}

	return false
}

// SetChangeableType gets a reference to the given string and assigns it to the ChangeableType field.
func (o *AuditQuestionsPutPreviousValues) SetChangeableType(v string) {
	o.ChangeableType = &v
}

// GetManagedChangeField returns the ManagedChangeField field value if set, zero value otherwise.
func (o *AuditQuestionsPutPreviousValues) GetManagedChangeField() string {
	if o == nil || IsNil(o.ManagedChangeField) {
		var ret string
		return ret
	}
	return *o.ManagedChangeField
}

// GetManagedChangeFieldOk returns a tuple with the ManagedChangeField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionsPutPreviousValues) GetManagedChangeFieldOk() (*string, bool) {
	if o == nil || IsNil(o.ManagedChangeField) {
		return nil, false
	}
	return o.ManagedChangeField, true
}

// HasManagedChangeField returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasManagedChangeField() bool {
	if o != nil && !IsNil(o.ManagedChangeField) {
		return true
	}

	return false
}

// SetManagedChangeField gets a reference to the given string and assigns it to the ManagedChangeField field.
func (o *AuditQuestionsPutPreviousValues) SetManagedChangeField(v string) {
	o.ManagedChangeField = &v
}

// GetManagedChangeFieldType returns the ManagedChangeFieldType field value if set, zero value otherwise.
func (o *AuditQuestionsPutPreviousValues) GetManagedChangeFieldType() string {
	if o == nil || IsNil(o.ManagedChangeFieldType) {
		var ret string
		return ret
	}
	return *o.ManagedChangeFieldType
}

// GetManagedChangeFieldTypeOk returns a tuple with the ManagedChangeFieldType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionsPutPreviousValues) GetManagedChangeFieldTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ManagedChangeFieldType) {
		return nil, false
	}
	return o.ManagedChangeFieldType, true
}

// HasManagedChangeFieldType returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasManagedChangeFieldType() bool {
	if o != nil && !IsNil(o.ManagedChangeFieldType) {
		return true
	}

	return false
}

// SetManagedChangeFieldType gets a reference to the given string and assigns it to the ManagedChangeFieldType field.
func (o *AuditQuestionsPutPreviousValues) SetManagedChangeFieldType(v string) {
	o.ManagedChangeFieldType = &v
}

// GetQuestionOptions returns the QuestionOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditQuestionsPutPreviousValues) GetQuestionOptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.QuestionOptions
}

// GetQuestionOptionsOk returns a tuple with the QuestionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditQuestionsPutPreviousValues) GetQuestionOptionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.QuestionOptions) {
		return nil, false
	}
	return &o.QuestionOptions, true
}

// HasQuestionOptions returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasQuestionOptions() bool {
	if o != nil && !IsNil(o.QuestionOptions) {
		return true
	}

	return false
}

// SetQuestionOptions gets a reference to the given interface{} and assigns it to the QuestionOptions field.
func (o *AuditQuestionsPutPreviousValues) SetQuestionOptions(v interface{}) {
	o.QuestionOptions = v
}

// GetAllowNaResponse returns the AllowNaResponse field value if set, zero value otherwise.
func (o *AuditQuestionsPutPreviousValues) GetAllowNaResponse() bool {
	if o == nil || IsNil(o.AllowNaResponse) {
		var ret bool
		return ret
	}
	return *o.AllowNaResponse
}

// GetAllowNaResponseOk returns a tuple with the AllowNaResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionsPutPreviousValues) GetAllowNaResponseOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowNaResponse) {
		return nil, false
	}
	return o.AllowNaResponse, true
}

// HasAllowNaResponse returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasAllowNaResponse() bool {
	if o != nil && !IsNil(o.AllowNaResponse) {
		return true
	}

	return false
}

// SetAllowNaResponse gets a reference to the given bool and assigns it to the AllowNaResponse field.
func (o *AuditQuestionsPutPreviousValues) SetAllowNaResponse(v bool) {
	o.AllowNaResponse = &v
}

// GetAllowExplanation returns the AllowExplanation field value if set, zero value otherwise.
func (o *AuditQuestionsPutPreviousValues) GetAllowExplanation() bool {
	if o == nil || IsNil(o.AllowExplanation) {
		var ret bool
		return ret
	}
	return *o.AllowExplanation
}

// GetAllowExplanationOk returns a tuple with the AllowExplanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionsPutPreviousValues) GetAllowExplanationOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowExplanation) {
		return nil, false
	}
	return o.AllowExplanation, true
}

// HasAllowExplanation returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasAllowExplanation() bool {
	if o != nil && !IsNil(o.AllowExplanation) {
		return true
	}

	return false
}

// SetAllowExplanation gets a reference to the given bool and assigns it to the AllowExplanation field.
func (o *AuditQuestionsPutPreviousValues) SetAllowExplanation(v bool) {
	o.AllowExplanation = &v
}

// GetAllowFiles returns the AllowFiles field value if set, zero value otherwise.
func (o *AuditQuestionsPutPreviousValues) GetAllowFiles() bool {
	if o == nil || IsNil(o.AllowFiles) {
		var ret bool
		return ret
	}
	return *o.AllowFiles
}

// GetAllowFilesOk returns a tuple with the AllowFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionsPutPreviousValues) GetAllowFilesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowFiles) {
		return nil, false
	}
	return o.AllowFiles, true
}

// HasAllowFiles returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasAllowFiles() bool {
	if o != nil && !IsNil(o.AllowFiles) {
		return true
	}

	return false
}

// SetAllowFiles gets a reference to the given bool and assigns it to the AllowFiles field.
func (o *AuditQuestionsPutPreviousValues) SetAllowFiles(v bool) {
	o.AllowFiles = &v
}

// GetRequireFiles returns the RequireFiles field value if set, zero value otherwise.
func (o *AuditQuestionsPutPreviousValues) GetRequireFiles() bool {
	if o == nil || IsNil(o.RequireFiles) {
		var ret bool
		return ret
	}
	return *o.RequireFiles
}

// GetRequireFilesOk returns a tuple with the RequireFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionsPutPreviousValues) GetRequireFilesOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireFiles) {
		return nil, false
	}
	return o.RequireFiles, true
}

// HasRequireFiles returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasRequireFiles() bool {
	if o != nil && !IsNil(o.RequireFiles) {
		return true
	}

	return false
}

// SetRequireFiles gets a reference to the given bool and assigns it to the RequireFiles field.
func (o *AuditQuestionsPutPreviousValues) SetRequireFiles(v bool) {
	o.RequireFiles = &v
}

// GetQuestionUuid returns the QuestionUuid field value if set, zero value otherwise.
func (o *AuditQuestionsPutPreviousValues) GetQuestionUuid() string {
	if o == nil || IsNil(o.QuestionUuid) {
		var ret string
		return ret
	}
	return *o.QuestionUuid
}

// GetQuestionUuidOk returns a tuple with the QuestionUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionsPutPreviousValues) GetQuestionUuidOk() (*string, bool) {
	if o == nil || IsNil(o.QuestionUuid) {
		return nil, false
	}
	return o.QuestionUuid, true
}

// HasQuestionUuid returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasQuestionUuid() bool {
	if o != nil && !IsNil(o.QuestionUuid) {
		return true
	}

	return false
}

// SetQuestionUuid gets a reference to the given string and assigns it to the QuestionUuid field.
func (o *AuditQuestionsPutPreviousValues) SetQuestionUuid(v string) {
	o.QuestionUuid = &v
}

// GetConditionQuestionUuid returns the ConditionQuestionUuid field value if set, zero value otherwise.
func (o *AuditQuestionsPutPreviousValues) GetConditionQuestionUuid() string {
	if o == nil || IsNil(o.ConditionQuestionUuid) {
		var ret string
		return ret
	}
	return *o.ConditionQuestionUuid
}

// GetConditionQuestionUuidOk returns a tuple with the ConditionQuestionUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionsPutPreviousValues) GetConditionQuestionUuidOk() (*string, bool) {
	if o == nil || IsNil(o.ConditionQuestionUuid) {
		return nil, false
	}
	return o.ConditionQuestionUuid, true
}

// HasConditionQuestionUuid returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasConditionQuestionUuid() bool {
	if o != nil && !IsNil(o.ConditionQuestionUuid) {
		return true
	}

	return false
}

// SetConditionQuestionUuid gets a reference to the given string and assigns it to the ConditionQuestionUuid field.
func (o *AuditQuestionsPutPreviousValues) SetConditionQuestionUuid(v string) {
	o.ConditionQuestionUuid = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditQuestionsPutPreviousValues) GetCondition() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditQuestionsPutPreviousValues) GetConditionOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return &o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given interface{} and assigns it to the Condition field.
func (o *AuditQuestionsPutPreviousValues) SetCondition(v interface{}) {
	o.Condition = v
}

// GetRequireNaExplanation returns the RequireNaExplanation field value if set, zero value otherwise.
func (o *AuditQuestionsPutPreviousValues) GetRequireNaExplanation() bool {
	if o == nil || IsNil(o.RequireNaExplanation) {
		var ret bool
		return ret
	}
	return *o.RequireNaExplanation
}

// GetRequireNaExplanationOk returns a tuple with the RequireNaExplanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditQuestionsPutPreviousValues) GetRequireNaExplanationOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireNaExplanation) {
		return nil, false
	}
	return o.RequireNaExplanation, true
}

// HasRequireNaExplanation returns a boolean if a field has been set.
func (o *AuditQuestionsPutPreviousValues) HasRequireNaExplanation() bool {
	if o != nil && !IsNil(o.RequireNaExplanation) {
		return true
	}

	return false
}

// SetRequireNaExplanation gets a reference to the given bool and assigns it to the RequireNaExplanation field.
func (o *AuditQuestionsPutPreviousValues) SetRequireNaExplanation(v bool) {
	o.RequireNaExplanation = &v
}

func (o AuditQuestionsPutPreviousValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditQuestionsPutPreviousValues) ToMap() (map[string]interface{}, error) {
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
	if o.Requirements != nil {
		toSerialize["requirements"] = o.Requirements
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.QuestionData != nil {
		toSerialize["question_data"] = o.QuestionData
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.SourceId) {
		toSerialize["source_id"] = o.SourceId
	}
	if !IsNil(o.CreatedByUserId) {
		toSerialize["created_by_user_id"] = o.CreatedByUserId
	}
	if !IsNil(o.ConditionChoiceUuid) {
		toSerialize["condition_choice_uuid"] = o.ConditionChoiceUuid
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	if !IsNil(o.ChangeableType) {
		toSerialize["changeable_type"] = o.ChangeableType
	}
	if !IsNil(o.ManagedChangeField) {
		toSerialize["managed_change_field"] = o.ManagedChangeField
	}
	if !IsNil(o.ManagedChangeFieldType) {
		toSerialize["managed_change_field_type"] = o.ManagedChangeFieldType
	}
	if o.QuestionOptions != nil {
		toSerialize["question_options"] = o.QuestionOptions
	}
	if !IsNil(o.AllowNaResponse) {
		toSerialize["allow_na_response"] = o.AllowNaResponse
	}
	if !IsNil(o.AllowExplanation) {
		toSerialize["allow_explanation"] = o.AllowExplanation
	}
	if !IsNil(o.AllowFiles) {
		toSerialize["allow_files"] = o.AllowFiles
	}
	if !IsNil(o.RequireFiles) {
		toSerialize["require_files"] = o.RequireFiles
	}
	if !IsNil(o.QuestionUuid) {
		toSerialize["question_uuid"] = o.QuestionUuid
	}
	if !IsNil(o.ConditionQuestionUuid) {
		toSerialize["condition_question_uuid"] = o.ConditionQuestionUuid
	}
	if o.Condition != nil {
		toSerialize["condition"] = o.Condition
	}
	if !IsNil(o.RequireNaExplanation) {
		toSerialize["require_na_explanation"] = o.RequireNaExplanation
	}
	return toSerialize, nil
}

type NullableAuditQuestionsPutPreviousValues struct {
	value *AuditQuestionsPutPreviousValues
	isSet bool
}

func (v NullableAuditQuestionsPutPreviousValues) Get() *AuditQuestionsPutPreviousValues {
	return v.value
}

func (v *NullableAuditQuestionsPutPreviousValues) Set(val *AuditQuestionsPutPreviousValues) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditQuestionsPutPreviousValues) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditQuestionsPutPreviousValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditQuestionsPutPreviousValues(val *AuditQuestionsPutPreviousValues) *NullableAuditQuestionsPutPreviousValues {
	return &NullableAuditQuestionsPutPreviousValues{value: val, isSet: true}
}

func (v NullableAuditQuestionsPutPreviousValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditQuestionsPutPreviousValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


