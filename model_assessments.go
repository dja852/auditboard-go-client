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

// checks if the Assessments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Assessments{}

// Assessments struct for Assessments
type Assessments struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
	Type *string `json:"type,omitempty"`
	// Note: This is a Foreign Key to `assessment_templates.id`.<fk table='assessment_templates' column='id'/>
	AssessmentTemplateId *int32 `json:"assessment_template_id,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	FinalizedByUserId *int32 `json:"finalized_by_user_id,omitempty"`
	FinalizedDate *string `json:"finalized_date,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	StartedByUserId *int32 `json:"started_by_user_id,omitempty"`
	StartedDate *string `json:"started_date,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	CanceledByUserId *int32 `json:"canceled_by_user_id,omitempty"`
	CanceledDate *string `json:"canceled_date,omitempty"`
	SurveyType *string `json:"survey_type,omitempty"`
	ProjectOptions interface{} `json:"project_options,omitempty"`
	DueDate *string `json:"due_date,omitempty"`
	EmailProjectStart interface{} `json:"email_project_start,omitempty"`
	LandingText *string `json:"landing_text,omitempty"`
	ConfirmText *string `json:"confirm_text,omitempty"`
	InstructionText *string `json:"instruction_text,omitempty"`
	// Note: This is a Foreign Key to `assessment_periods.id`.<fk table='assessment_periods' column='id'/>
	AssessmentPeriodId *int32 `json:"assessment_period_id,omitempty"`
	Assessables interface{} `json:"assessables,omitempty"`
	Info interface{} `json:"info,omitempty"`
}

// NewAssessments instantiates a new Assessments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssessments() *Assessments {
	this := Assessments{}
	return &this
}

// NewAssessmentsWithDefaults instantiates a new Assessments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssessmentsWithDefaults() *Assessments {
	this := Assessments{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Assessments) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assessments) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Assessments) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Assessments) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Assessments) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assessments) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Assessments) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Assessments) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Assessments) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assessments) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Assessments) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Assessments) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Assessments) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assessments) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Assessments) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Assessments) SetType(v string) {
	o.Type = &v
}

// GetAssessmentTemplateId returns the AssessmentTemplateId field value if set, zero value otherwise.
func (o *Assessments) GetAssessmentTemplateId() int32 {
	if o == nil || IsNil(o.AssessmentTemplateId) {
		var ret int32
		return ret
	}
	return *o.AssessmentTemplateId
}

// GetAssessmentTemplateIdOk returns a tuple with the AssessmentTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assessments) GetAssessmentTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AssessmentTemplateId) {
		return nil, false
	}
	return o.AssessmentTemplateId, true
}

// HasAssessmentTemplateId returns a boolean if a field has been set.
func (o *Assessments) HasAssessmentTemplateId() bool {
	if o != nil && !IsNil(o.AssessmentTemplateId) {
		return true
	}

	return false
}

// SetAssessmentTemplateId gets a reference to the given int32 and assigns it to the AssessmentTemplateId field.
func (o *Assessments) SetAssessmentTemplateId(v int32) {
	o.AssessmentTemplateId = &v
}

// GetFinalizedByUserId returns the FinalizedByUserId field value if set, zero value otherwise.
func (o *Assessments) GetFinalizedByUserId() int32 {
	if o == nil || IsNil(o.FinalizedByUserId) {
		var ret int32
		return ret
	}
	return *o.FinalizedByUserId
}

// GetFinalizedByUserIdOk returns a tuple with the FinalizedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assessments) GetFinalizedByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FinalizedByUserId) {
		return nil, false
	}
	return o.FinalizedByUserId, true
}

// HasFinalizedByUserId returns a boolean if a field has been set.
func (o *Assessments) HasFinalizedByUserId() bool {
	if o != nil && !IsNil(o.FinalizedByUserId) {
		return true
	}

	return false
}

// SetFinalizedByUserId gets a reference to the given int32 and assigns it to the FinalizedByUserId field.
func (o *Assessments) SetFinalizedByUserId(v int32) {
	o.FinalizedByUserId = &v
}

// GetFinalizedDate returns the FinalizedDate field value if set, zero value otherwise.
func (o *Assessments) GetFinalizedDate() string {
	if o == nil || IsNil(o.FinalizedDate) {
		var ret string
		return ret
	}
	return *o.FinalizedDate
}

// GetFinalizedDateOk returns a tuple with the FinalizedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assessments) GetFinalizedDateOk() (*string, bool) {
	if o == nil || IsNil(o.FinalizedDate) {
		return nil, false
	}
	return o.FinalizedDate, true
}

// HasFinalizedDate returns a boolean if a field has been set.
func (o *Assessments) HasFinalizedDate() bool {
	if o != nil && !IsNil(o.FinalizedDate) {
		return true
	}

	return false
}

// SetFinalizedDate gets a reference to the given string and assigns it to the FinalizedDate field.
func (o *Assessments) SetFinalizedDate(v string) {
	o.FinalizedDate = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Assessments) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assessments) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Assessments) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Assessments) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Assessments) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assessments) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Assessments) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Assessments) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Assessments) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assessments) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Assessments) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *Assessments) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetStartedByUserId returns the StartedByUserId field value if set, zero value otherwise.
func (o *Assessments) GetStartedByUserId() int32 {
	if o == nil || IsNil(o.StartedByUserId) {
		var ret int32
		return ret
	}
	return *o.StartedByUserId
}

// GetStartedByUserIdOk returns a tuple with the StartedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assessments) GetStartedByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.StartedByUserId) {
		return nil, false
	}
	return o.StartedByUserId, true
}

// HasStartedByUserId returns a boolean if a field has been set.
func (o *Assessments) HasStartedByUserId() bool {
	if o != nil && !IsNil(o.StartedByUserId) {
		return true
	}

	return false
}

// SetStartedByUserId gets a reference to the given int32 and assigns it to the StartedByUserId field.
func (o *Assessments) SetStartedByUserId(v int32) {
	o.StartedByUserId = &v
}

// GetStartedDate returns the StartedDate field value if set, zero value otherwise.
func (o *Assessments) GetStartedDate() string {
	if o == nil || IsNil(o.StartedDate) {
		var ret string
		return ret
	}
	return *o.StartedDate
}

// GetStartedDateOk returns a tuple with the StartedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assessments) GetStartedDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartedDate) {
		return nil, false
	}
	return o.StartedDate, true
}

// HasStartedDate returns a boolean if a field has been set.
func (o *Assessments) HasStartedDate() bool {
	if o != nil && !IsNil(o.StartedDate) {
		return true
	}

	return false
}

// SetStartedDate gets a reference to the given string and assigns it to the StartedDate field.
func (o *Assessments) SetStartedDate(v string) {
	o.StartedDate = &v
}

// GetCanceledByUserId returns the CanceledByUserId field value if set, zero value otherwise.
func (o *Assessments) GetCanceledByUserId() int32 {
	if o == nil || IsNil(o.CanceledByUserId) {
		var ret int32
		return ret
	}
	return *o.CanceledByUserId
}

// GetCanceledByUserIdOk returns a tuple with the CanceledByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assessments) GetCanceledByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CanceledByUserId) {
		return nil, false
	}
	return o.CanceledByUserId, true
}

// HasCanceledByUserId returns a boolean if a field has been set.
func (o *Assessments) HasCanceledByUserId() bool {
	if o != nil && !IsNil(o.CanceledByUserId) {
		return true
	}

	return false
}

// SetCanceledByUserId gets a reference to the given int32 and assigns it to the CanceledByUserId field.
func (o *Assessments) SetCanceledByUserId(v int32) {
	o.CanceledByUserId = &v
}

// GetCanceledDate returns the CanceledDate field value if set, zero value otherwise.
func (o *Assessments) GetCanceledDate() string {
	if o == nil || IsNil(o.CanceledDate) {
		var ret string
		return ret
	}
	return *o.CanceledDate
}

// GetCanceledDateOk returns a tuple with the CanceledDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assessments) GetCanceledDateOk() (*string, bool) {
	if o == nil || IsNil(o.CanceledDate) {
		return nil, false
	}
	return o.CanceledDate, true
}

// HasCanceledDate returns a boolean if a field has been set.
func (o *Assessments) HasCanceledDate() bool {
	if o != nil && !IsNil(o.CanceledDate) {
		return true
	}

	return false
}

// SetCanceledDate gets a reference to the given string and assigns it to the CanceledDate field.
func (o *Assessments) SetCanceledDate(v string) {
	o.CanceledDate = &v
}

// GetSurveyType returns the SurveyType field value if set, zero value otherwise.
func (o *Assessments) GetSurveyType() string {
	if o == nil || IsNil(o.SurveyType) {
		var ret string
		return ret
	}
	return *o.SurveyType
}

// GetSurveyTypeOk returns a tuple with the SurveyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assessments) GetSurveyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SurveyType) {
		return nil, false
	}
	return o.SurveyType, true
}

// HasSurveyType returns a boolean if a field has been set.
func (o *Assessments) HasSurveyType() bool {
	if o != nil && !IsNil(o.SurveyType) {
		return true
	}

	return false
}

// SetSurveyType gets a reference to the given string and assigns it to the SurveyType field.
func (o *Assessments) SetSurveyType(v string) {
	o.SurveyType = &v
}

// GetProjectOptions returns the ProjectOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Assessments) GetProjectOptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ProjectOptions
}

// GetProjectOptionsOk returns a tuple with the ProjectOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Assessments) GetProjectOptionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ProjectOptions) {
		return nil, false
	}
	return &o.ProjectOptions, true
}

// HasProjectOptions returns a boolean if a field has been set.
func (o *Assessments) HasProjectOptions() bool {
	if o != nil && !IsNil(o.ProjectOptions) {
		return true
	}

	return false
}

// SetProjectOptions gets a reference to the given interface{} and assigns it to the ProjectOptions field.
func (o *Assessments) SetProjectOptions(v interface{}) {
	o.ProjectOptions = v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *Assessments) GetDueDate() string {
	if o == nil || IsNil(o.DueDate) {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assessments) GetDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *Assessments) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *Assessments) SetDueDate(v string) {
	o.DueDate = &v
}

// GetEmailProjectStart returns the EmailProjectStart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Assessments) GetEmailProjectStart() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EmailProjectStart
}

// GetEmailProjectStartOk returns a tuple with the EmailProjectStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Assessments) GetEmailProjectStartOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EmailProjectStart) {
		return nil, false
	}
	return &o.EmailProjectStart, true
}

// HasEmailProjectStart returns a boolean if a field has been set.
func (o *Assessments) HasEmailProjectStart() bool {
	if o != nil && !IsNil(o.EmailProjectStart) {
		return true
	}

	return false
}

// SetEmailProjectStart gets a reference to the given interface{} and assigns it to the EmailProjectStart field.
func (o *Assessments) SetEmailProjectStart(v interface{}) {
	o.EmailProjectStart = v
}

// GetLandingText returns the LandingText field value if set, zero value otherwise.
func (o *Assessments) GetLandingText() string {
	if o == nil || IsNil(o.LandingText) {
		var ret string
		return ret
	}
	return *o.LandingText
}

// GetLandingTextOk returns a tuple with the LandingText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assessments) GetLandingTextOk() (*string, bool) {
	if o == nil || IsNil(o.LandingText) {
		return nil, false
	}
	return o.LandingText, true
}

// HasLandingText returns a boolean if a field has been set.
func (o *Assessments) HasLandingText() bool {
	if o != nil && !IsNil(o.LandingText) {
		return true
	}

	return false
}

// SetLandingText gets a reference to the given string and assigns it to the LandingText field.
func (o *Assessments) SetLandingText(v string) {
	o.LandingText = &v
}

// GetConfirmText returns the ConfirmText field value if set, zero value otherwise.
func (o *Assessments) GetConfirmText() string {
	if o == nil || IsNil(o.ConfirmText) {
		var ret string
		return ret
	}
	return *o.ConfirmText
}

// GetConfirmTextOk returns a tuple with the ConfirmText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assessments) GetConfirmTextOk() (*string, bool) {
	if o == nil || IsNil(o.ConfirmText) {
		return nil, false
	}
	return o.ConfirmText, true
}

// HasConfirmText returns a boolean if a field has been set.
func (o *Assessments) HasConfirmText() bool {
	if o != nil && !IsNil(o.ConfirmText) {
		return true
	}

	return false
}

// SetConfirmText gets a reference to the given string and assigns it to the ConfirmText field.
func (o *Assessments) SetConfirmText(v string) {
	o.ConfirmText = &v
}

// GetInstructionText returns the InstructionText field value if set, zero value otherwise.
func (o *Assessments) GetInstructionText() string {
	if o == nil || IsNil(o.InstructionText) {
		var ret string
		return ret
	}
	return *o.InstructionText
}

// GetInstructionTextOk returns a tuple with the InstructionText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assessments) GetInstructionTextOk() (*string, bool) {
	if o == nil || IsNil(o.InstructionText) {
		return nil, false
	}
	return o.InstructionText, true
}

// HasInstructionText returns a boolean if a field has been set.
func (o *Assessments) HasInstructionText() bool {
	if o != nil && !IsNil(o.InstructionText) {
		return true
	}

	return false
}

// SetInstructionText gets a reference to the given string and assigns it to the InstructionText field.
func (o *Assessments) SetInstructionText(v string) {
	o.InstructionText = &v
}

// GetAssessmentPeriodId returns the AssessmentPeriodId field value if set, zero value otherwise.
func (o *Assessments) GetAssessmentPeriodId() int32 {
	if o == nil || IsNil(o.AssessmentPeriodId) {
		var ret int32
		return ret
	}
	return *o.AssessmentPeriodId
}

// GetAssessmentPeriodIdOk returns a tuple with the AssessmentPeriodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assessments) GetAssessmentPeriodIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AssessmentPeriodId) {
		return nil, false
	}
	return o.AssessmentPeriodId, true
}

// HasAssessmentPeriodId returns a boolean if a field has been set.
func (o *Assessments) HasAssessmentPeriodId() bool {
	if o != nil && !IsNil(o.AssessmentPeriodId) {
		return true
	}

	return false
}

// SetAssessmentPeriodId gets a reference to the given int32 and assigns it to the AssessmentPeriodId field.
func (o *Assessments) SetAssessmentPeriodId(v int32) {
	o.AssessmentPeriodId = &v
}

// GetAssessables returns the Assessables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Assessments) GetAssessables() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Assessables
}

// GetAssessablesOk returns a tuple with the Assessables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Assessments) GetAssessablesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Assessables) {
		return nil, false
	}
	return &o.Assessables, true
}

// HasAssessables returns a boolean if a field has been set.
func (o *Assessments) HasAssessables() bool {
	if o != nil && !IsNil(o.Assessables) {
		return true
	}

	return false
}

// SetAssessables gets a reference to the given interface{} and assigns it to the Assessables field.
func (o *Assessments) SetAssessables(v interface{}) {
	o.Assessables = v
}

// GetInfo returns the Info field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Assessments) GetInfo() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Assessments) GetInfoOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return &o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *Assessments) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given interface{} and assigns it to the Info field.
func (o *Assessments) SetInfo(v interface{}) {
	o.Info = v
}

func (o Assessments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Assessments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.AssessmentTemplateId) {
		toSerialize["assessment_template_id"] = o.AssessmentTemplateId
	}
	if !IsNil(o.FinalizedByUserId) {
		toSerialize["finalized_by_user_id"] = o.FinalizedByUserId
	}
	if !IsNil(o.FinalizedDate) {
		toSerialize["finalized_date"] = o.FinalizedDate
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
	if !IsNil(o.StartedByUserId) {
		toSerialize["started_by_user_id"] = o.StartedByUserId
	}
	if !IsNil(o.StartedDate) {
		toSerialize["started_date"] = o.StartedDate
	}
	if !IsNil(o.CanceledByUserId) {
		toSerialize["canceled_by_user_id"] = o.CanceledByUserId
	}
	if !IsNil(o.CanceledDate) {
		toSerialize["canceled_date"] = o.CanceledDate
	}
	if !IsNil(o.SurveyType) {
		toSerialize["survey_type"] = o.SurveyType
	}
	if o.ProjectOptions != nil {
		toSerialize["project_options"] = o.ProjectOptions
	}
	if !IsNil(o.DueDate) {
		toSerialize["due_date"] = o.DueDate
	}
	if o.EmailProjectStart != nil {
		toSerialize["email_project_start"] = o.EmailProjectStart
	}
	if !IsNil(o.LandingText) {
		toSerialize["landing_text"] = o.LandingText
	}
	if !IsNil(o.ConfirmText) {
		toSerialize["confirm_text"] = o.ConfirmText
	}
	if !IsNil(o.InstructionText) {
		toSerialize["instruction_text"] = o.InstructionText
	}
	if !IsNil(o.AssessmentPeriodId) {
		toSerialize["assessment_period_id"] = o.AssessmentPeriodId
	}
	if o.Assessables != nil {
		toSerialize["assessables"] = o.Assessables
	}
	if o.Info != nil {
		toSerialize["info"] = o.Info
	}
	return toSerialize, nil
}

type NullableAssessments struct {
	value *Assessments
	isSet bool
}

func (v NullableAssessments) Get() *Assessments {
	return v.value
}

func (v *NullableAssessments) Set(val *Assessments) {
	v.value = val
	v.isSet = true
}

func (v NullableAssessments) IsSet() bool {
	return v.isSet
}

func (v *NullableAssessments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssessments(val *Assessments) *NullableAssessments {
	return &NullableAssessments{value: val, isSet: true}
}

func (v NullableAssessments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssessments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


