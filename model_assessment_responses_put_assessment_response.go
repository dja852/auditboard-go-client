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

// checks if the AssessmentResponsesPutAssessmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssessmentResponsesPutAssessmentResponse{}

// AssessmentResponsesPutAssessmentResponse struct for AssessmentResponsesPutAssessmentResponse
type AssessmentResponsesPutAssessmentResponse struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
	ResponseJson interface{} `json:"response_json,omitempty"`
	AssessableType *string `json:"assessable_type,omitempty"`
	AssessableId *int32 `json:"assessable_id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	SubmittedByUserId *int32 `json:"submitted_by_user_id,omitempty"`
	SubmittedDate *string `json:"submitted_date,omitempty"`
	ReferenceMeta interface{} `json:"reference_meta,omitempty"`
	// Note: This is a Foreign Key to `assessments.id`.<fk table='assessments' column='id'/>
	AssessmentId *int32 `json:"assessment_id,omitempty"`
	Source *string `json:"source,omitempty"`
	// Note: This is a Foreign Key to `assessment_templates.id`.<fk table='assessment_templates' column='id'/>
	AssessmentTemplateId *int32 `json:"assessment_template_id,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	FinalizedByUserId *int32 `json:"finalized_by_user_id,omitempty"`
	FinalizedDate *string `json:"finalized_date,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	AnsweredByUserId *int32 `json:"answered_by_user_id,omitempty"`
	AnsweredDate *string `json:"answered_date,omitempty"`
	IsAnswered bool `json:"is_answered"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	AssigneeUserId *int32 `json:"assignee_user_id,omitempty"`
	IsNotApplicable bool `json:"is_not_applicable"`
	Notes *string `json:"notes,omitempty"`
	// Note: This is a Foreign Key to `user_assessments.id`.<fk table='user_assessments' column='id'/>
	UserAssessmentId *int32 `json:"user_assessment_id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	CreatedByUserId *int32 `json:"created_by_user_id,omitempty"`
	// Note: This is a Foreign Key to `risk_responses.id`.<fk table='risk_responses' column='id'/>
	RiskResponseId *int32 `json:"risk_response_id,omitempty"`
	RiskresponseactionId *int32 `json:"riskresponseaction_id,omitempty"`
	RiskresponseactionType *string `json:"riskresponseaction_type,omitempty"`
	// Note: This is a Foreign Key to `assessment_responses.id`.<fk table='assessment_responses' column='id'/>
	ParentAssessmentResponseId *int32 `json:"parent_assessment_response_id,omitempty"`
}

type _AssessmentResponsesPutAssessmentResponse AssessmentResponsesPutAssessmentResponse

// NewAssessmentResponsesPutAssessmentResponse instantiates a new AssessmentResponsesPutAssessmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssessmentResponsesPutAssessmentResponse(isAnswered bool, isNotApplicable bool) *AssessmentResponsesPutAssessmentResponse {
	this := AssessmentResponsesPutAssessmentResponse{}
	this.IsAnswered = isAnswered
	this.IsNotApplicable = isNotApplicable
	return &this
}

// NewAssessmentResponsesPutAssessmentResponseWithDefaults instantiates a new AssessmentResponsesPutAssessmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssessmentResponsesPutAssessmentResponseWithDefaults() *AssessmentResponsesPutAssessmentResponse {
	this := AssessmentResponsesPutAssessmentResponse{}
	var isAnswered bool = false
	this.IsAnswered = isAnswered
	var isNotApplicable bool = false
	this.IsNotApplicable = isNotApplicable
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AssessmentResponsesPutAssessmentResponse) SetId(v int32) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AssessmentResponsesPutAssessmentResponse) SetStatus(v string) {
	o.Status = &v
}

// GetResponseJson returns the ResponseJson field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssessmentResponsesPutAssessmentResponse) GetResponseJson() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ResponseJson
}

// GetResponseJsonOk returns a tuple with the ResponseJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssessmentResponsesPutAssessmentResponse) GetResponseJsonOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ResponseJson) {
		return nil, false
	}
	return &o.ResponseJson, true
}

// HasResponseJson returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasResponseJson() bool {
	if o != nil && !IsNil(o.ResponseJson) {
		return true
	}

	return false
}

// SetResponseJson gets a reference to the given interface{} and assigns it to the ResponseJson field.
func (o *AssessmentResponsesPutAssessmentResponse) SetResponseJson(v interface{}) {
	o.ResponseJson = v
}

// GetAssessableType returns the AssessableType field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetAssessableType() string {
	if o == nil || IsNil(o.AssessableType) {
		var ret string
		return ret
	}
	return *o.AssessableType
}

// GetAssessableTypeOk returns a tuple with the AssessableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetAssessableTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AssessableType) {
		return nil, false
	}
	return o.AssessableType, true
}

// HasAssessableType returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasAssessableType() bool {
	if o != nil && !IsNil(o.AssessableType) {
		return true
	}

	return false
}

// SetAssessableType gets a reference to the given string and assigns it to the AssessableType field.
func (o *AssessmentResponsesPutAssessmentResponse) SetAssessableType(v string) {
	o.AssessableType = &v
}

// GetAssessableId returns the AssessableId field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetAssessableId() int32 {
	if o == nil || IsNil(o.AssessableId) {
		var ret int32
		return ret
	}
	return *o.AssessableId
}

// GetAssessableIdOk returns a tuple with the AssessableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetAssessableIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AssessableId) {
		return nil, false
	}
	return o.AssessableId, true
}

// HasAssessableId returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasAssessableId() bool {
	if o != nil && !IsNil(o.AssessableId) {
		return true
	}

	return false
}

// SetAssessableId gets a reference to the given int32 and assigns it to the AssessableId field.
func (o *AssessmentResponsesPutAssessmentResponse) SetAssessableId(v int32) {
	o.AssessableId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AssessmentResponsesPutAssessmentResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AssessmentResponsesPutAssessmentResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *AssessmentResponsesPutAssessmentResponse) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetSubmittedByUserId returns the SubmittedByUserId field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetSubmittedByUserId() int32 {
	if o == nil || IsNil(o.SubmittedByUserId) {
		var ret int32
		return ret
	}
	return *o.SubmittedByUserId
}

// GetSubmittedByUserIdOk returns a tuple with the SubmittedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetSubmittedByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SubmittedByUserId) {
		return nil, false
	}
	return o.SubmittedByUserId, true
}

// HasSubmittedByUserId returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasSubmittedByUserId() bool {
	if o != nil && !IsNil(o.SubmittedByUserId) {
		return true
	}

	return false
}

// SetSubmittedByUserId gets a reference to the given int32 and assigns it to the SubmittedByUserId field.
func (o *AssessmentResponsesPutAssessmentResponse) SetSubmittedByUserId(v int32) {
	o.SubmittedByUserId = &v
}

// GetSubmittedDate returns the SubmittedDate field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetSubmittedDate() string {
	if o == nil || IsNil(o.SubmittedDate) {
		var ret string
		return ret
	}
	return *o.SubmittedDate
}

// GetSubmittedDateOk returns a tuple with the SubmittedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetSubmittedDateOk() (*string, bool) {
	if o == nil || IsNil(o.SubmittedDate) {
		return nil, false
	}
	return o.SubmittedDate, true
}

// HasSubmittedDate returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasSubmittedDate() bool {
	if o != nil && !IsNil(o.SubmittedDate) {
		return true
	}

	return false
}

// SetSubmittedDate gets a reference to the given string and assigns it to the SubmittedDate field.
func (o *AssessmentResponsesPutAssessmentResponse) SetSubmittedDate(v string) {
	o.SubmittedDate = &v
}

// GetReferenceMeta returns the ReferenceMeta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssessmentResponsesPutAssessmentResponse) GetReferenceMeta() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceMeta
}

// GetReferenceMetaOk returns a tuple with the ReferenceMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssessmentResponsesPutAssessmentResponse) GetReferenceMetaOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceMeta) {
		return nil, false
	}
	return &o.ReferenceMeta, true
}

// HasReferenceMeta returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasReferenceMeta() bool {
	if o != nil && !IsNil(o.ReferenceMeta) {
		return true
	}

	return false
}

// SetReferenceMeta gets a reference to the given interface{} and assigns it to the ReferenceMeta field.
func (o *AssessmentResponsesPutAssessmentResponse) SetReferenceMeta(v interface{}) {
	o.ReferenceMeta = v
}

// GetAssessmentId returns the AssessmentId field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetAssessmentId() int32 {
	if o == nil || IsNil(o.AssessmentId) {
		var ret int32
		return ret
	}
	return *o.AssessmentId
}

// GetAssessmentIdOk returns a tuple with the AssessmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetAssessmentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AssessmentId) {
		return nil, false
	}
	return o.AssessmentId, true
}

// HasAssessmentId returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasAssessmentId() bool {
	if o != nil && !IsNil(o.AssessmentId) {
		return true
	}

	return false
}

// SetAssessmentId gets a reference to the given int32 and assigns it to the AssessmentId field.
func (o *AssessmentResponsesPutAssessmentResponse) SetAssessmentId(v int32) {
	o.AssessmentId = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *AssessmentResponsesPutAssessmentResponse) SetSource(v string) {
	o.Source = &v
}

// GetAssessmentTemplateId returns the AssessmentTemplateId field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetAssessmentTemplateId() int32 {
	if o == nil || IsNil(o.AssessmentTemplateId) {
		var ret int32
		return ret
	}
	return *o.AssessmentTemplateId
}

// GetAssessmentTemplateIdOk returns a tuple with the AssessmentTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetAssessmentTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AssessmentTemplateId) {
		return nil, false
	}
	return o.AssessmentTemplateId, true
}

// HasAssessmentTemplateId returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasAssessmentTemplateId() bool {
	if o != nil && !IsNil(o.AssessmentTemplateId) {
		return true
	}

	return false
}

// SetAssessmentTemplateId gets a reference to the given int32 and assigns it to the AssessmentTemplateId field.
func (o *AssessmentResponsesPutAssessmentResponse) SetAssessmentTemplateId(v int32) {
	o.AssessmentTemplateId = &v
}

// GetFinalizedByUserId returns the FinalizedByUserId field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetFinalizedByUserId() int32 {
	if o == nil || IsNil(o.FinalizedByUserId) {
		var ret int32
		return ret
	}
	return *o.FinalizedByUserId
}

// GetFinalizedByUserIdOk returns a tuple with the FinalizedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetFinalizedByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FinalizedByUserId) {
		return nil, false
	}
	return o.FinalizedByUserId, true
}

// HasFinalizedByUserId returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasFinalizedByUserId() bool {
	if o != nil && !IsNil(o.FinalizedByUserId) {
		return true
	}

	return false
}

// SetFinalizedByUserId gets a reference to the given int32 and assigns it to the FinalizedByUserId field.
func (o *AssessmentResponsesPutAssessmentResponse) SetFinalizedByUserId(v int32) {
	o.FinalizedByUserId = &v
}

// GetFinalizedDate returns the FinalizedDate field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetFinalizedDate() string {
	if o == nil || IsNil(o.FinalizedDate) {
		var ret string
		return ret
	}
	return *o.FinalizedDate
}

// GetFinalizedDateOk returns a tuple with the FinalizedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetFinalizedDateOk() (*string, bool) {
	if o == nil || IsNil(o.FinalizedDate) {
		return nil, false
	}
	return o.FinalizedDate, true
}

// HasFinalizedDate returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasFinalizedDate() bool {
	if o != nil && !IsNil(o.FinalizedDate) {
		return true
	}

	return false
}

// SetFinalizedDate gets a reference to the given string and assigns it to the FinalizedDate field.
func (o *AssessmentResponsesPutAssessmentResponse) SetFinalizedDate(v string) {
	o.FinalizedDate = &v
}

// GetAnsweredByUserId returns the AnsweredByUserId field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetAnsweredByUserId() int32 {
	if o == nil || IsNil(o.AnsweredByUserId) {
		var ret int32
		return ret
	}
	return *o.AnsweredByUserId
}

// GetAnsweredByUserIdOk returns a tuple with the AnsweredByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetAnsweredByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AnsweredByUserId) {
		return nil, false
	}
	return o.AnsweredByUserId, true
}

// HasAnsweredByUserId returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasAnsweredByUserId() bool {
	if o != nil && !IsNil(o.AnsweredByUserId) {
		return true
	}

	return false
}

// SetAnsweredByUserId gets a reference to the given int32 and assigns it to the AnsweredByUserId field.
func (o *AssessmentResponsesPutAssessmentResponse) SetAnsweredByUserId(v int32) {
	o.AnsweredByUserId = &v
}

// GetAnsweredDate returns the AnsweredDate field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetAnsweredDate() string {
	if o == nil || IsNil(o.AnsweredDate) {
		var ret string
		return ret
	}
	return *o.AnsweredDate
}

// GetAnsweredDateOk returns a tuple with the AnsweredDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetAnsweredDateOk() (*string, bool) {
	if o == nil || IsNil(o.AnsweredDate) {
		return nil, false
	}
	return o.AnsweredDate, true
}

// HasAnsweredDate returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasAnsweredDate() bool {
	if o != nil && !IsNil(o.AnsweredDate) {
		return true
	}

	return false
}

// SetAnsweredDate gets a reference to the given string and assigns it to the AnsweredDate field.
func (o *AssessmentResponsesPutAssessmentResponse) SetAnsweredDate(v string) {
	o.AnsweredDate = &v
}

// GetIsAnswered returns the IsAnswered field value
func (o *AssessmentResponsesPutAssessmentResponse) GetIsAnswered() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAnswered
}

// GetIsAnsweredOk returns a tuple with the IsAnswered field value
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetIsAnsweredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAnswered, true
}

// SetIsAnswered sets field value
func (o *AssessmentResponsesPutAssessmentResponse) SetIsAnswered(v bool) {
	o.IsAnswered = v
}

// GetAssigneeUserId returns the AssigneeUserId field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetAssigneeUserId() int32 {
	if o == nil || IsNil(o.AssigneeUserId) {
		var ret int32
		return ret
	}
	return *o.AssigneeUserId
}

// GetAssigneeUserIdOk returns a tuple with the AssigneeUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetAssigneeUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AssigneeUserId) {
		return nil, false
	}
	return o.AssigneeUserId, true
}

// HasAssigneeUserId returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasAssigneeUserId() bool {
	if o != nil && !IsNil(o.AssigneeUserId) {
		return true
	}

	return false
}

// SetAssigneeUserId gets a reference to the given int32 and assigns it to the AssigneeUserId field.
func (o *AssessmentResponsesPutAssessmentResponse) SetAssigneeUserId(v int32) {
	o.AssigneeUserId = &v
}

// GetIsNotApplicable returns the IsNotApplicable field value
func (o *AssessmentResponsesPutAssessmentResponse) GetIsNotApplicable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsNotApplicable
}

// GetIsNotApplicableOk returns a tuple with the IsNotApplicable field value
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetIsNotApplicableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsNotApplicable, true
}

// SetIsNotApplicable sets field value
func (o *AssessmentResponsesPutAssessmentResponse) SetIsNotApplicable(v bool) {
	o.IsNotApplicable = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *AssessmentResponsesPutAssessmentResponse) SetNotes(v string) {
	o.Notes = &v
}

// GetUserAssessmentId returns the UserAssessmentId field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetUserAssessmentId() int32 {
	if o == nil || IsNil(o.UserAssessmentId) {
		var ret int32
		return ret
	}
	return *o.UserAssessmentId
}

// GetUserAssessmentIdOk returns a tuple with the UserAssessmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetUserAssessmentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserAssessmentId) {
		return nil, false
	}
	return o.UserAssessmentId, true
}

// HasUserAssessmentId returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasUserAssessmentId() bool {
	if o != nil && !IsNil(o.UserAssessmentId) {
		return true
	}

	return false
}

// SetUserAssessmentId gets a reference to the given int32 and assigns it to the UserAssessmentId field.
func (o *AssessmentResponsesPutAssessmentResponse) SetUserAssessmentId(v int32) {
	o.UserAssessmentId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AssessmentResponsesPutAssessmentResponse) SetName(v string) {
	o.Name = &v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetCreatedByUserId() int32 {
	if o == nil || IsNil(o.CreatedByUserId) {
		var ret int32
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetCreatedByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedByUserId) {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasCreatedByUserId() bool {
	if o != nil && !IsNil(o.CreatedByUserId) {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given int32 and assigns it to the CreatedByUserId field.
func (o *AssessmentResponsesPutAssessmentResponse) SetCreatedByUserId(v int32) {
	o.CreatedByUserId = &v
}

// GetRiskResponseId returns the RiskResponseId field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetRiskResponseId() int32 {
	if o == nil || IsNil(o.RiskResponseId) {
		var ret int32
		return ret
	}
	return *o.RiskResponseId
}

// GetRiskResponseIdOk returns a tuple with the RiskResponseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetRiskResponseIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RiskResponseId) {
		return nil, false
	}
	return o.RiskResponseId, true
}

// HasRiskResponseId returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasRiskResponseId() bool {
	if o != nil && !IsNil(o.RiskResponseId) {
		return true
	}

	return false
}

// SetRiskResponseId gets a reference to the given int32 and assigns it to the RiskResponseId field.
func (o *AssessmentResponsesPutAssessmentResponse) SetRiskResponseId(v int32) {
	o.RiskResponseId = &v
}

// GetRiskresponseactionId returns the RiskresponseactionId field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetRiskresponseactionId() int32 {
	if o == nil || IsNil(o.RiskresponseactionId) {
		var ret int32
		return ret
	}
	return *o.RiskresponseactionId
}

// GetRiskresponseactionIdOk returns a tuple with the RiskresponseactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetRiskresponseactionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RiskresponseactionId) {
		return nil, false
	}
	return o.RiskresponseactionId, true
}

// HasRiskresponseactionId returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasRiskresponseactionId() bool {
	if o != nil && !IsNil(o.RiskresponseactionId) {
		return true
	}

	return false
}

// SetRiskresponseactionId gets a reference to the given int32 and assigns it to the RiskresponseactionId field.
func (o *AssessmentResponsesPutAssessmentResponse) SetRiskresponseactionId(v int32) {
	o.RiskresponseactionId = &v
}

// GetRiskresponseactionType returns the RiskresponseactionType field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetRiskresponseactionType() string {
	if o == nil || IsNil(o.RiskresponseactionType) {
		var ret string
		return ret
	}
	return *o.RiskresponseactionType
}

// GetRiskresponseactionTypeOk returns a tuple with the RiskresponseactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetRiskresponseactionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RiskresponseactionType) {
		return nil, false
	}
	return o.RiskresponseactionType, true
}

// HasRiskresponseactionType returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasRiskresponseactionType() bool {
	if o != nil && !IsNil(o.RiskresponseactionType) {
		return true
	}

	return false
}

// SetRiskresponseactionType gets a reference to the given string and assigns it to the RiskresponseactionType field.
func (o *AssessmentResponsesPutAssessmentResponse) SetRiskresponseactionType(v string) {
	o.RiskresponseactionType = &v
}

// GetParentAssessmentResponseId returns the ParentAssessmentResponseId field value if set, zero value otherwise.
func (o *AssessmentResponsesPutAssessmentResponse) GetParentAssessmentResponseId() int32 {
	if o == nil || IsNil(o.ParentAssessmentResponseId) {
		var ret int32
		return ret
	}
	return *o.ParentAssessmentResponseId
}

// GetParentAssessmentResponseIdOk returns a tuple with the ParentAssessmentResponseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentResponsesPutAssessmentResponse) GetParentAssessmentResponseIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ParentAssessmentResponseId) {
		return nil, false
	}
	return o.ParentAssessmentResponseId, true
}

// HasParentAssessmentResponseId returns a boolean if a field has been set.
func (o *AssessmentResponsesPutAssessmentResponse) HasParentAssessmentResponseId() bool {
	if o != nil && !IsNil(o.ParentAssessmentResponseId) {
		return true
	}

	return false
}

// SetParentAssessmentResponseId gets a reference to the given int32 and assigns it to the ParentAssessmentResponseId field.
func (o *AssessmentResponsesPutAssessmentResponse) SetParentAssessmentResponseId(v int32) {
	o.ParentAssessmentResponseId = &v
}

func (o AssessmentResponsesPutAssessmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssessmentResponsesPutAssessmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.ResponseJson != nil {
		toSerialize["response_json"] = o.ResponseJson
	}
	if !IsNil(o.AssessableType) {
		toSerialize["assessable_type"] = o.AssessableType
	}
	if !IsNil(o.AssessableId) {
		toSerialize["assessable_id"] = o.AssessableId
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
	if !IsNil(o.SubmittedByUserId) {
		toSerialize["submitted_by_user_id"] = o.SubmittedByUserId
	}
	if !IsNil(o.SubmittedDate) {
		toSerialize["submitted_date"] = o.SubmittedDate
	}
	if o.ReferenceMeta != nil {
		toSerialize["reference_meta"] = o.ReferenceMeta
	}
	if !IsNil(o.AssessmentId) {
		toSerialize["assessment_id"] = o.AssessmentId
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
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
	if !IsNil(o.AnsweredByUserId) {
		toSerialize["answered_by_user_id"] = o.AnsweredByUserId
	}
	if !IsNil(o.AnsweredDate) {
		toSerialize["answered_date"] = o.AnsweredDate
	}
	toSerialize["is_answered"] = o.IsAnswered
	if !IsNil(o.AssigneeUserId) {
		toSerialize["assignee_user_id"] = o.AssigneeUserId
	}
	toSerialize["is_not_applicable"] = o.IsNotApplicable
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.UserAssessmentId) {
		toSerialize["user_assessment_id"] = o.UserAssessmentId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CreatedByUserId) {
		toSerialize["created_by_user_id"] = o.CreatedByUserId
	}
	if !IsNil(o.RiskResponseId) {
		toSerialize["risk_response_id"] = o.RiskResponseId
	}
	if !IsNil(o.RiskresponseactionId) {
		toSerialize["riskresponseaction_id"] = o.RiskresponseactionId
	}
	if !IsNil(o.RiskresponseactionType) {
		toSerialize["riskresponseaction_type"] = o.RiskresponseactionType
	}
	if !IsNil(o.ParentAssessmentResponseId) {
		toSerialize["parent_assessment_response_id"] = o.ParentAssessmentResponseId
	}
	return toSerialize, nil
}

type NullableAssessmentResponsesPutAssessmentResponse struct {
	value *AssessmentResponsesPutAssessmentResponse
	isSet bool
}

func (v NullableAssessmentResponsesPutAssessmentResponse) Get() *AssessmentResponsesPutAssessmentResponse {
	return v.value
}

func (v *NullableAssessmentResponsesPutAssessmentResponse) Set(val *AssessmentResponsesPutAssessmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssessmentResponsesPutAssessmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssessmentResponsesPutAssessmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssessmentResponsesPutAssessmentResponse(val *AssessmentResponsesPutAssessmentResponse) *NullableAssessmentResponsesPutAssessmentResponse {
	return &NullableAssessmentResponsesPutAssessmentResponse{value: val, isSet: true}
}

func (v NullableAssessmentResponsesPutAssessmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssessmentResponsesPutAssessmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


