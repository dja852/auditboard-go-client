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

// checks if the SurveysPutPreviousValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SurveysPutPreviousValues{}

// SurveysPutPreviousValues struct for SurveysPutPreviousValues
type SurveysPutPreviousValues struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	MetaConfig *string `json:"meta_config,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	TaskTitle *string `json:"task_title,omitempty"`
	Type *string `json:"type,omitempty"`
	Status *string `json:"status,omitempty"`
	// Note: This is a Foreign Key to `task_categories.id`.<fk table='task_categories' column='id'/>
	TaskCategoryId *int32 `json:"task_category_id,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	DefaultUsers interface{} `json:"default_users,omitempty"`
	EmailProjectStart interface{} `json:"email_project_start,omitempty"`
	EmailPreparerDigest interface{} `json:"email_preparer_digest,omitempty"`
	EmailReviewerDigest interface{} `json:"email_reviewer_digest,omitempty"`
	EmailAdminDigest interface{} `json:"email_admin_digest,omitempty"`
	EmailSubscriberDigest interface{} `json:"email_subscriber_digest,omitempty"`
	AdditionalPreparersEnabled *bool `json:"additional_preparers_enabled,omitempty"`
	EmailAdditionalPreparerDigest interface{} `json:"email_additional_preparer_digest,omitempty"`
	ProjectOptions interface{} `json:"project_options,omitempty"`
	AllowYesToAll *bool `json:"allow_yes_to_all,omitempty"`
	PackagingIsStandard *bool `json:"packaging_is_standard,omitempty"`
}

// NewSurveysPutPreviousValues instantiates a new SurveysPutPreviousValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSurveysPutPreviousValues() *SurveysPutPreviousValues {
	this := SurveysPutPreviousValues{}
	var taskTitle string = ""
	this.TaskTitle = &taskTitle
	var type_ string = "standalone"
	this.Type = &type_
	var status string = "Unlocked"
	this.Status = &status
	var isArchived bool = false
	this.IsArchived = &isArchived
	var additionalPreparersEnabled bool = false
	this.AdditionalPreparersEnabled = &additionalPreparersEnabled
	var packagingIsStandard bool = false
	this.PackagingIsStandard = &packagingIsStandard
	return &this
}

// NewSurveysPutPreviousValuesWithDefaults instantiates a new SurveysPutPreviousValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSurveysPutPreviousValuesWithDefaults() *SurveysPutPreviousValues {
	this := SurveysPutPreviousValues{}
	var taskTitle string = ""
	this.TaskTitle = &taskTitle
	var type_ string = "standalone"
	this.Type = &type_
	var status string = "Unlocked"
	this.Status = &status
	var isArchived bool = false
	this.IsArchived = &isArchived
	var additionalPreparersEnabled bool = false
	this.AdditionalPreparersEnabled = &additionalPreparersEnabled
	var packagingIsStandard bool = false
	this.PackagingIsStandard = &packagingIsStandard
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SurveysPutPreviousValues) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveysPutPreviousValues) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SurveysPutPreviousValues) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SurveysPutPreviousValues) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveysPutPreviousValues) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SurveysPutPreviousValues) SetName(v string) {
	o.Name = &v
}

// GetMetaConfig returns the MetaConfig field value if set, zero value otherwise.
func (o *SurveysPutPreviousValues) GetMetaConfig() string {
	if o == nil || IsNil(o.MetaConfig) {
		var ret string
		return ret
	}
	return *o.MetaConfig
}

// GetMetaConfigOk returns a tuple with the MetaConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveysPutPreviousValues) GetMetaConfigOk() (*string, bool) {
	if o == nil || IsNil(o.MetaConfig) {
		return nil, false
	}
	return o.MetaConfig, true
}

// HasMetaConfig returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasMetaConfig() bool {
	if o != nil && !IsNil(o.MetaConfig) {
		return true
	}

	return false
}

// SetMetaConfig gets a reference to the given string and assigns it to the MetaConfig field.
func (o *SurveysPutPreviousValues) SetMetaConfig(v string) {
	o.MetaConfig = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SurveysPutPreviousValues) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveysPutPreviousValues) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *SurveysPutPreviousValues) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SurveysPutPreviousValues) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveysPutPreviousValues) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *SurveysPutPreviousValues) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *SurveysPutPreviousValues) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveysPutPreviousValues) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *SurveysPutPreviousValues) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetTaskTitle returns the TaskTitle field value if set, zero value otherwise.
func (o *SurveysPutPreviousValues) GetTaskTitle() string {
	if o == nil || IsNil(o.TaskTitle) {
		var ret string
		return ret
	}
	return *o.TaskTitle
}

// GetTaskTitleOk returns a tuple with the TaskTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveysPutPreviousValues) GetTaskTitleOk() (*string, bool) {
	if o == nil || IsNil(o.TaskTitle) {
		return nil, false
	}
	return o.TaskTitle, true
}

// HasTaskTitle returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasTaskTitle() bool {
	if o != nil && !IsNil(o.TaskTitle) {
		return true
	}

	return false
}

// SetTaskTitle gets a reference to the given string and assigns it to the TaskTitle field.
func (o *SurveysPutPreviousValues) SetTaskTitle(v string) {
	o.TaskTitle = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SurveysPutPreviousValues) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveysPutPreviousValues) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SurveysPutPreviousValues) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SurveysPutPreviousValues) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveysPutPreviousValues) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SurveysPutPreviousValues) SetStatus(v string) {
	o.Status = &v
}

// GetTaskCategoryId returns the TaskCategoryId field value if set, zero value otherwise.
func (o *SurveysPutPreviousValues) GetTaskCategoryId() int32 {
	if o == nil || IsNil(o.TaskCategoryId) {
		var ret int32
		return ret
	}
	return *o.TaskCategoryId
}

// GetTaskCategoryIdOk returns a tuple with the TaskCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveysPutPreviousValues) GetTaskCategoryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TaskCategoryId) {
		return nil, false
	}
	return o.TaskCategoryId, true
}

// HasTaskCategoryId returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasTaskCategoryId() bool {
	if o != nil && !IsNil(o.TaskCategoryId) {
		return true
	}

	return false
}

// SetTaskCategoryId gets a reference to the given int32 and assigns it to the TaskCategoryId field.
func (o *SurveysPutPreviousValues) SetTaskCategoryId(v int32) {
	o.TaskCategoryId = &v
}

// GetIsArchived returns the IsArchived field value if set, zero value otherwise.
func (o *SurveysPutPreviousValues) GetIsArchived() bool {
	if o == nil || IsNil(o.IsArchived) {
		var ret bool
		return ret
	}
	return *o.IsArchived
}

// GetIsArchivedOk returns a tuple with the IsArchived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveysPutPreviousValues) GetIsArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsArchived) {
		return nil, false
	}
	return o.IsArchived, true
}

// HasIsArchived returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasIsArchived() bool {
	if o != nil && !IsNil(o.IsArchived) {
		return true
	}

	return false
}

// SetIsArchived gets a reference to the given bool and assigns it to the IsArchived field.
func (o *SurveysPutPreviousValues) SetIsArchived(v bool) {
	o.IsArchived = &v
}

// GetDefaultUsers returns the DefaultUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SurveysPutPreviousValues) GetDefaultUsers() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DefaultUsers
}

// GetDefaultUsersOk returns a tuple with the DefaultUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SurveysPutPreviousValues) GetDefaultUsersOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DefaultUsers) {
		return nil, false
	}
	return &o.DefaultUsers, true
}

// HasDefaultUsers returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasDefaultUsers() bool {
	if o != nil && !IsNil(o.DefaultUsers) {
		return true
	}

	return false
}

// SetDefaultUsers gets a reference to the given interface{} and assigns it to the DefaultUsers field.
func (o *SurveysPutPreviousValues) SetDefaultUsers(v interface{}) {
	o.DefaultUsers = v
}

// GetEmailProjectStart returns the EmailProjectStart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SurveysPutPreviousValues) GetEmailProjectStart() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EmailProjectStart
}

// GetEmailProjectStartOk returns a tuple with the EmailProjectStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SurveysPutPreviousValues) GetEmailProjectStartOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EmailProjectStart) {
		return nil, false
	}
	return &o.EmailProjectStart, true
}

// HasEmailProjectStart returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasEmailProjectStart() bool {
	if o != nil && !IsNil(o.EmailProjectStart) {
		return true
	}

	return false
}

// SetEmailProjectStart gets a reference to the given interface{} and assigns it to the EmailProjectStart field.
func (o *SurveysPutPreviousValues) SetEmailProjectStart(v interface{}) {
	o.EmailProjectStart = v
}

// GetEmailPreparerDigest returns the EmailPreparerDigest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SurveysPutPreviousValues) GetEmailPreparerDigest() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EmailPreparerDigest
}

// GetEmailPreparerDigestOk returns a tuple with the EmailPreparerDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SurveysPutPreviousValues) GetEmailPreparerDigestOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EmailPreparerDigest) {
		return nil, false
	}
	return &o.EmailPreparerDigest, true
}

// HasEmailPreparerDigest returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasEmailPreparerDigest() bool {
	if o != nil && !IsNil(o.EmailPreparerDigest) {
		return true
	}

	return false
}

// SetEmailPreparerDigest gets a reference to the given interface{} and assigns it to the EmailPreparerDigest field.
func (o *SurveysPutPreviousValues) SetEmailPreparerDigest(v interface{}) {
	o.EmailPreparerDigest = v
}

// GetEmailReviewerDigest returns the EmailReviewerDigest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SurveysPutPreviousValues) GetEmailReviewerDigest() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EmailReviewerDigest
}

// GetEmailReviewerDigestOk returns a tuple with the EmailReviewerDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SurveysPutPreviousValues) GetEmailReviewerDigestOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EmailReviewerDigest) {
		return nil, false
	}
	return &o.EmailReviewerDigest, true
}

// HasEmailReviewerDigest returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasEmailReviewerDigest() bool {
	if o != nil && !IsNil(o.EmailReviewerDigest) {
		return true
	}

	return false
}

// SetEmailReviewerDigest gets a reference to the given interface{} and assigns it to the EmailReviewerDigest field.
func (o *SurveysPutPreviousValues) SetEmailReviewerDigest(v interface{}) {
	o.EmailReviewerDigest = v
}

// GetEmailAdminDigest returns the EmailAdminDigest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SurveysPutPreviousValues) GetEmailAdminDigest() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EmailAdminDigest
}

// GetEmailAdminDigestOk returns a tuple with the EmailAdminDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SurveysPutPreviousValues) GetEmailAdminDigestOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EmailAdminDigest) {
		return nil, false
	}
	return &o.EmailAdminDigest, true
}

// HasEmailAdminDigest returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasEmailAdminDigest() bool {
	if o != nil && !IsNil(o.EmailAdminDigest) {
		return true
	}

	return false
}

// SetEmailAdminDigest gets a reference to the given interface{} and assigns it to the EmailAdminDigest field.
func (o *SurveysPutPreviousValues) SetEmailAdminDigest(v interface{}) {
	o.EmailAdminDigest = v
}

// GetEmailSubscriberDigest returns the EmailSubscriberDigest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SurveysPutPreviousValues) GetEmailSubscriberDigest() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EmailSubscriberDigest
}

// GetEmailSubscriberDigestOk returns a tuple with the EmailSubscriberDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SurveysPutPreviousValues) GetEmailSubscriberDigestOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EmailSubscriberDigest) {
		return nil, false
	}
	return &o.EmailSubscriberDigest, true
}

// HasEmailSubscriberDigest returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasEmailSubscriberDigest() bool {
	if o != nil && !IsNil(o.EmailSubscriberDigest) {
		return true
	}

	return false
}

// SetEmailSubscriberDigest gets a reference to the given interface{} and assigns it to the EmailSubscriberDigest field.
func (o *SurveysPutPreviousValues) SetEmailSubscriberDigest(v interface{}) {
	o.EmailSubscriberDigest = v
}

// GetAdditionalPreparersEnabled returns the AdditionalPreparersEnabled field value if set, zero value otherwise.
func (o *SurveysPutPreviousValues) GetAdditionalPreparersEnabled() bool {
	if o == nil || IsNil(o.AdditionalPreparersEnabled) {
		var ret bool
		return ret
	}
	return *o.AdditionalPreparersEnabled
}

// GetAdditionalPreparersEnabledOk returns a tuple with the AdditionalPreparersEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveysPutPreviousValues) GetAdditionalPreparersEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AdditionalPreparersEnabled) {
		return nil, false
	}
	return o.AdditionalPreparersEnabled, true
}

// HasAdditionalPreparersEnabled returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasAdditionalPreparersEnabled() bool {
	if o != nil && !IsNil(o.AdditionalPreparersEnabled) {
		return true
	}

	return false
}

// SetAdditionalPreparersEnabled gets a reference to the given bool and assigns it to the AdditionalPreparersEnabled field.
func (o *SurveysPutPreviousValues) SetAdditionalPreparersEnabled(v bool) {
	o.AdditionalPreparersEnabled = &v
}

// GetEmailAdditionalPreparerDigest returns the EmailAdditionalPreparerDigest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SurveysPutPreviousValues) GetEmailAdditionalPreparerDigest() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EmailAdditionalPreparerDigest
}

// GetEmailAdditionalPreparerDigestOk returns a tuple with the EmailAdditionalPreparerDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SurveysPutPreviousValues) GetEmailAdditionalPreparerDigestOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EmailAdditionalPreparerDigest) {
		return nil, false
	}
	return &o.EmailAdditionalPreparerDigest, true
}

// HasEmailAdditionalPreparerDigest returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasEmailAdditionalPreparerDigest() bool {
	if o != nil && !IsNil(o.EmailAdditionalPreparerDigest) {
		return true
	}

	return false
}

// SetEmailAdditionalPreparerDigest gets a reference to the given interface{} and assigns it to the EmailAdditionalPreparerDigest field.
func (o *SurveysPutPreviousValues) SetEmailAdditionalPreparerDigest(v interface{}) {
	o.EmailAdditionalPreparerDigest = v
}

// GetProjectOptions returns the ProjectOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SurveysPutPreviousValues) GetProjectOptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ProjectOptions
}

// GetProjectOptionsOk returns a tuple with the ProjectOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SurveysPutPreviousValues) GetProjectOptionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ProjectOptions) {
		return nil, false
	}
	return &o.ProjectOptions, true
}

// HasProjectOptions returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasProjectOptions() bool {
	if o != nil && !IsNil(o.ProjectOptions) {
		return true
	}

	return false
}

// SetProjectOptions gets a reference to the given interface{} and assigns it to the ProjectOptions field.
func (o *SurveysPutPreviousValues) SetProjectOptions(v interface{}) {
	o.ProjectOptions = v
}

// GetAllowYesToAll returns the AllowYesToAll field value if set, zero value otherwise.
func (o *SurveysPutPreviousValues) GetAllowYesToAll() bool {
	if o == nil || IsNil(o.AllowYesToAll) {
		var ret bool
		return ret
	}
	return *o.AllowYesToAll
}

// GetAllowYesToAllOk returns a tuple with the AllowYesToAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveysPutPreviousValues) GetAllowYesToAllOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowYesToAll) {
		return nil, false
	}
	return o.AllowYesToAll, true
}

// HasAllowYesToAll returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasAllowYesToAll() bool {
	if o != nil && !IsNil(o.AllowYesToAll) {
		return true
	}

	return false
}

// SetAllowYesToAll gets a reference to the given bool and assigns it to the AllowYesToAll field.
func (o *SurveysPutPreviousValues) SetAllowYesToAll(v bool) {
	o.AllowYesToAll = &v
}

// GetPackagingIsStandard returns the PackagingIsStandard field value if set, zero value otherwise.
func (o *SurveysPutPreviousValues) GetPackagingIsStandard() bool {
	if o == nil || IsNil(o.PackagingIsStandard) {
		var ret bool
		return ret
	}
	return *o.PackagingIsStandard
}

// GetPackagingIsStandardOk returns a tuple with the PackagingIsStandard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveysPutPreviousValues) GetPackagingIsStandardOk() (*bool, bool) {
	if o == nil || IsNil(o.PackagingIsStandard) {
		return nil, false
	}
	return o.PackagingIsStandard, true
}

// HasPackagingIsStandard returns a boolean if a field has been set.
func (o *SurveysPutPreviousValues) HasPackagingIsStandard() bool {
	if o != nil && !IsNil(o.PackagingIsStandard) {
		return true
	}

	return false
}

// SetPackagingIsStandard gets a reference to the given bool and assigns it to the PackagingIsStandard field.
func (o *SurveysPutPreviousValues) SetPackagingIsStandard(v bool) {
	o.PackagingIsStandard = &v
}

func (o SurveysPutPreviousValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SurveysPutPreviousValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.MetaConfig) {
		toSerialize["meta_config"] = o.MetaConfig
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
	if !IsNil(o.TaskTitle) {
		toSerialize["task_title"] = o.TaskTitle
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TaskCategoryId) {
		toSerialize["task_category_id"] = o.TaskCategoryId
	}
	if !IsNil(o.IsArchived) {
		toSerialize["is_archived"] = o.IsArchived
	}
	if o.DefaultUsers != nil {
		toSerialize["default_users"] = o.DefaultUsers
	}
	if o.EmailProjectStart != nil {
		toSerialize["email_project_start"] = o.EmailProjectStart
	}
	if o.EmailPreparerDigest != nil {
		toSerialize["email_preparer_digest"] = o.EmailPreparerDigest
	}
	if o.EmailReviewerDigest != nil {
		toSerialize["email_reviewer_digest"] = o.EmailReviewerDigest
	}
	if o.EmailAdminDigest != nil {
		toSerialize["email_admin_digest"] = o.EmailAdminDigest
	}
	if o.EmailSubscriberDigest != nil {
		toSerialize["email_subscriber_digest"] = o.EmailSubscriberDigest
	}
	if !IsNil(o.AdditionalPreparersEnabled) {
		toSerialize["additional_preparers_enabled"] = o.AdditionalPreparersEnabled
	}
	if o.EmailAdditionalPreparerDigest != nil {
		toSerialize["email_additional_preparer_digest"] = o.EmailAdditionalPreparerDigest
	}
	if o.ProjectOptions != nil {
		toSerialize["project_options"] = o.ProjectOptions
	}
	if !IsNil(o.AllowYesToAll) {
		toSerialize["allow_yes_to_all"] = o.AllowYesToAll
	}
	if !IsNil(o.PackagingIsStandard) {
		toSerialize["packaging_is_standard"] = o.PackagingIsStandard
	}
	return toSerialize, nil
}

type NullableSurveysPutPreviousValues struct {
	value *SurveysPutPreviousValues
	isSet bool
}

func (v NullableSurveysPutPreviousValues) Get() *SurveysPutPreviousValues {
	return v.value
}

func (v *NullableSurveysPutPreviousValues) Set(val *SurveysPutPreviousValues) {
	v.value = val
	v.isSet = true
}

func (v NullableSurveysPutPreviousValues) IsSet() bool {
	return v.isSet
}

func (v *NullableSurveysPutPreviousValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSurveysPutPreviousValues(val *SurveysPutPreviousValues) *NullableSurveysPutPreviousValues {
	return &NullableSurveysPutPreviousValues{value: val, isSet: true}
}

func (v NullableSurveysPutPreviousValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSurveysPutPreviousValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

