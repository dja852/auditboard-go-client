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

// checks if the Surveys type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Surveys{}

// Surveys struct for Surveys
type Surveys struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	MetaConfig *string `json:"meta_config,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	TaskTitle *string `json:"task_title,omitempty"`
	Type *string `json:"type,omitempty"`
	Status string `json:"status"`
	// Note: This is a Foreign Key to `task_categories.id`.<fk table='task_categories' column='id'/>
	TaskCategoryId *int32 `json:"task_category_id,omitempty"`
	IsArchived bool `json:"is_archived"`
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
	PackagingIsStandard bool `json:"packaging_is_standard"`
}

type _Surveys Surveys

// NewSurveys instantiates a new Surveys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSurveys(status string, isArchived bool, packagingIsStandard bool) *Surveys {
	this := Surveys{}
	var taskTitle string = ""
	this.TaskTitle = &taskTitle
	var type_ string = "standalone"
	this.Type = &type_
	this.Status = status
	this.IsArchived = isArchived
	var additionalPreparersEnabled bool = false
	this.AdditionalPreparersEnabled = &additionalPreparersEnabled
	this.PackagingIsStandard = packagingIsStandard
	return &this
}

// NewSurveysWithDefaults instantiates a new Surveys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSurveysWithDefaults() *Surveys {
	this := Surveys{}
	var taskTitle string = ""
	this.TaskTitle = &taskTitle
	var type_ string = "standalone"
	this.Type = &type_
	var status string = "Unlocked"
	this.Status = status
	var isArchived bool = false
	this.IsArchived = isArchived
	var additionalPreparersEnabled bool = false
	this.AdditionalPreparersEnabled = &additionalPreparersEnabled
	var packagingIsStandard bool = false
	this.PackagingIsStandard = packagingIsStandard
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Surveys) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Surveys) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Surveys) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Surveys) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Surveys) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Surveys) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Surveys) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Surveys) SetName(v string) {
	o.Name = &v
}

// GetMetaConfig returns the MetaConfig field value if set, zero value otherwise.
func (o *Surveys) GetMetaConfig() string {
	if o == nil || IsNil(o.MetaConfig) {
		var ret string
		return ret
	}
	return *o.MetaConfig
}

// GetMetaConfigOk returns a tuple with the MetaConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Surveys) GetMetaConfigOk() (*string, bool) {
	if o == nil || IsNil(o.MetaConfig) {
		return nil, false
	}
	return o.MetaConfig, true
}

// HasMetaConfig returns a boolean if a field has been set.
func (o *Surveys) HasMetaConfig() bool {
	if o != nil && !IsNil(o.MetaConfig) {
		return true
	}

	return false
}

// SetMetaConfig gets a reference to the given string and assigns it to the MetaConfig field.
func (o *Surveys) SetMetaConfig(v string) {
	o.MetaConfig = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Surveys) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Surveys) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Surveys) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Surveys) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Surveys) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Surveys) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Surveys) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Surveys) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Surveys) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Surveys) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Surveys) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *Surveys) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetTaskTitle returns the TaskTitle field value if set, zero value otherwise.
func (o *Surveys) GetTaskTitle() string {
	if o == nil || IsNil(o.TaskTitle) {
		var ret string
		return ret
	}
	return *o.TaskTitle
}

// GetTaskTitleOk returns a tuple with the TaskTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Surveys) GetTaskTitleOk() (*string, bool) {
	if o == nil || IsNil(o.TaskTitle) {
		return nil, false
	}
	return o.TaskTitle, true
}

// HasTaskTitle returns a boolean if a field has been set.
func (o *Surveys) HasTaskTitle() bool {
	if o != nil && !IsNil(o.TaskTitle) {
		return true
	}

	return false
}

// SetTaskTitle gets a reference to the given string and assigns it to the TaskTitle field.
func (o *Surveys) SetTaskTitle(v string) {
	o.TaskTitle = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Surveys) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Surveys) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Surveys) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Surveys) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value
func (o *Surveys) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Surveys) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Surveys) SetStatus(v string) {
	o.Status = v
}

// GetTaskCategoryId returns the TaskCategoryId field value if set, zero value otherwise.
func (o *Surveys) GetTaskCategoryId() int32 {
	if o == nil || IsNil(o.TaskCategoryId) {
		var ret int32
		return ret
	}
	return *o.TaskCategoryId
}

// GetTaskCategoryIdOk returns a tuple with the TaskCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Surveys) GetTaskCategoryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TaskCategoryId) {
		return nil, false
	}
	return o.TaskCategoryId, true
}

// HasTaskCategoryId returns a boolean if a field has been set.
func (o *Surveys) HasTaskCategoryId() bool {
	if o != nil && !IsNil(o.TaskCategoryId) {
		return true
	}

	return false
}

// SetTaskCategoryId gets a reference to the given int32 and assigns it to the TaskCategoryId field.
func (o *Surveys) SetTaskCategoryId(v int32) {
	o.TaskCategoryId = &v
}

// GetIsArchived returns the IsArchived field value
func (o *Surveys) GetIsArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsArchived
}

// GetIsArchivedOk returns a tuple with the IsArchived field value
// and a boolean to check if the value has been set.
func (o *Surveys) GetIsArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsArchived, true
}

// SetIsArchived sets field value
func (o *Surveys) SetIsArchived(v bool) {
	o.IsArchived = v
}

// GetDefaultUsers returns the DefaultUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Surveys) GetDefaultUsers() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DefaultUsers
}

// GetDefaultUsersOk returns a tuple with the DefaultUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Surveys) GetDefaultUsersOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DefaultUsers) {
		return nil, false
	}
	return &o.DefaultUsers, true
}

// HasDefaultUsers returns a boolean if a field has been set.
func (o *Surveys) HasDefaultUsers() bool {
	if o != nil && !IsNil(o.DefaultUsers) {
		return true
	}

	return false
}

// SetDefaultUsers gets a reference to the given interface{} and assigns it to the DefaultUsers field.
func (o *Surveys) SetDefaultUsers(v interface{}) {
	o.DefaultUsers = v
}

// GetEmailProjectStart returns the EmailProjectStart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Surveys) GetEmailProjectStart() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EmailProjectStart
}

// GetEmailProjectStartOk returns a tuple with the EmailProjectStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Surveys) GetEmailProjectStartOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EmailProjectStart) {
		return nil, false
	}
	return &o.EmailProjectStart, true
}

// HasEmailProjectStart returns a boolean if a field has been set.
func (o *Surveys) HasEmailProjectStart() bool {
	if o != nil && !IsNil(o.EmailProjectStart) {
		return true
	}

	return false
}

// SetEmailProjectStart gets a reference to the given interface{} and assigns it to the EmailProjectStart field.
func (o *Surveys) SetEmailProjectStart(v interface{}) {
	o.EmailProjectStart = v
}

// GetEmailPreparerDigest returns the EmailPreparerDigest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Surveys) GetEmailPreparerDigest() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EmailPreparerDigest
}

// GetEmailPreparerDigestOk returns a tuple with the EmailPreparerDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Surveys) GetEmailPreparerDigestOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EmailPreparerDigest) {
		return nil, false
	}
	return &o.EmailPreparerDigest, true
}

// HasEmailPreparerDigest returns a boolean if a field has been set.
func (o *Surveys) HasEmailPreparerDigest() bool {
	if o != nil && !IsNil(o.EmailPreparerDigest) {
		return true
	}

	return false
}

// SetEmailPreparerDigest gets a reference to the given interface{} and assigns it to the EmailPreparerDigest field.
func (o *Surveys) SetEmailPreparerDigest(v interface{}) {
	o.EmailPreparerDigest = v
}

// GetEmailReviewerDigest returns the EmailReviewerDigest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Surveys) GetEmailReviewerDigest() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EmailReviewerDigest
}

// GetEmailReviewerDigestOk returns a tuple with the EmailReviewerDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Surveys) GetEmailReviewerDigestOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EmailReviewerDigest) {
		return nil, false
	}
	return &o.EmailReviewerDigest, true
}

// HasEmailReviewerDigest returns a boolean if a field has been set.
func (o *Surveys) HasEmailReviewerDigest() bool {
	if o != nil && !IsNil(o.EmailReviewerDigest) {
		return true
	}

	return false
}

// SetEmailReviewerDigest gets a reference to the given interface{} and assigns it to the EmailReviewerDigest field.
func (o *Surveys) SetEmailReviewerDigest(v interface{}) {
	o.EmailReviewerDigest = v
}

// GetEmailAdminDigest returns the EmailAdminDigest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Surveys) GetEmailAdminDigest() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EmailAdminDigest
}

// GetEmailAdminDigestOk returns a tuple with the EmailAdminDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Surveys) GetEmailAdminDigestOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EmailAdminDigest) {
		return nil, false
	}
	return &o.EmailAdminDigest, true
}

// HasEmailAdminDigest returns a boolean if a field has been set.
func (o *Surveys) HasEmailAdminDigest() bool {
	if o != nil && !IsNil(o.EmailAdminDigest) {
		return true
	}

	return false
}

// SetEmailAdminDigest gets a reference to the given interface{} and assigns it to the EmailAdminDigest field.
func (o *Surveys) SetEmailAdminDigest(v interface{}) {
	o.EmailAdminDigest = v
}

// GetEmailSubscriberDigest returns the EmailSubscriberDigest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Surveys) GetEmailSubscriberDigest() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EmailSubscriberDigest
}

// GetEmailSubscriberDigestOk returns a tuple with the EmailSubscriberDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Surveys) GetEmailSubscriberDigestOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EmailSubscriberDigest) {
		return nil, false
	}
	return &o.EmailSubscriberDigest, true
}

// HasEmailSubscriberDigest returns a boolean if a field has been set.
func (o *Surveys) HasEmailSubscriberDigest() bool {
	if o != nil && !IsNil(o.EmailSubscriberDigest) {
		return true
	}

	return false
}

// SetEmailSubscriberDigest gets a reference to the given interface{} and assigns it to the EmailSubscriberDigest field.
func (o *Surveys) SetEmailSubscriberDigest(v interface{}) {
	o.EmailSubscriberDigest = v
}

// GetAdditionalPreparersEnabled returns the AdditionalPreparersEnabled field value if set, zero value otherwise.
func (o *Surveys) GetAdditionalPreparersEnabled() bool {
	if o == nil || IsNil(o.AdditionalPreparersEnabled) {
		var ret bool
		return ret
	}
	return *o.AdditionalPreparersEnabled
}

// GetAdditionalPreparersEnabledOk returns a tuple with the AdditionalPreparersEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Surveys) GetAdditionalPreparersEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AdditionalPreparersEnabled) {
		return nil, false
	}
	return o.AdditionalPreparersEnabled, true
}

// HasAdditionalPreparersEnabled returns a boolean if a field has been set.
func (o *Surveys) HasAdditionalPreparersEnabled() bool {
	if o != nil && !IsNil(o.AdditionalPreparersEnabled) {
		return true
	}

	return false
}

// SetAdditionalPreparersEnabled gets a reference to the given bool and assigns it to the AdditionalPreparersEnabled field.
func (o *Surveys) SetAdditionalPreparersEnabled(v bool) {
	o.AdditionalPreparersEnabled = &v
}

// GetEmailAdditionalPreparerDigest returns the EmailAdditionalPreparerDigest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Surveys) GetEmailAdditionalPreparerDigest() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EmailAdditionalPreparerDigest
}

// GetEmailAdditionalPreparerDigestOk returns a tuple with the EmailAdditionalPreparerDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Surveys) GetEmailAdditionalPreparerDigestOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EmailAdditionalPreparerDigest) {
		return nil, false
	}
	return &o.EmailAdditionalPreparerDigest, true
}

// HasEmailAdditionalPreparerDigest returns a boolean if a field has been set.
func (o *Surveys) HasEmailAdditionalPreparerDigest() bool {
	if o != nil && !IsNil(o.EmailAdditionalPreparerDigest) {
		return true
	}

	return false
}

// SetEmailAdditionalPreparerDigest gets a reference to the given interface{} and assigns it to the EmailAdditionalPreparerDigest field.
func (o *Surveys) SetEmailAdditionalPreparerDigest(v interface{}) {
	o.EmailAdditionalPreparerDigest = v
}

// GetProjectOptions returns the ProjectOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Surveys) GetProjectOptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ProjectOptions
}

// GetProjectOptionsOk returns a tuple with the ProjectOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Surveys) GetProjectOptionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ProjectOptions) {
		return nil, false
	}
	return &o.ProjectOptions, true
}

// HasProjectOptions returns a boolean if a field has been set.
func (o *Surveys) HasProjectOptions() bool {
	if o != nil && !IsNil(o.ProjectOptions) {
		return true
	}

	return false
}

// SetProjectOptions gets a reference to the given interface{} and assigns it to the ProjectOptions field.
func (o *Surveys) SetProjectOptions(v interface{}) {
	o.ProjectOptions = v
}

// GetAllowYesToAll returns the AllowYesToAll field value if set, zero value otherwise.
func (o *Surveys) GetAllowYesToAll() bool {
	if o == nil || IsNil(o.AllowYesToAll) {
		var ret bool
		return ret
	}
	return *o.AllowYesToAll
}

// GetAllowYesToAllOk returns a tuple with the AllowYesToAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Surveys) GetAllowYesToAllOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowYesToAll) {
		return nil, false
	}
	return o.AllowYesToAll, true
}

// HasAllowYesToAll returns a boolean if a field has been set.
func (o *Surveys) HasAllowYesToAll() bool {
	if o != nil && !IsNil(o.AllowYesToAll) {
		return true
	}

	return false
}

// SetAllowYesToAll gets a reference to the given bool and assigns it to the AllowYesToAll field.
func (o *Surveys) SetAllowYesToAll(v bool) {
	o.AllowYesToAll = &v
}

// GetPackagingIsStandard returns the PackagingIsStandard field value
func (o *Surveys) GetPackagingIsStandard() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PackagingIsStandard
}

// GetPackagingIsStandardOk returns a tuple with the PackagingIsStandard field value
// and a boolean to check if the value has been set.
func (o *Surveys) GetPackagingIsStandardOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackagingIsStandard, true
}

// SetPackagingIsStandard sets field value
func (o *Surveys) SetPackagingIsStandard(v bool) {
	o.PackagingIsStandard = v
}

func (o Surveys) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Surveys) ToMap() (map[string]interface{}, error) {
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
	toSerialize["status"] = o.Status
	if !IsNil(o.TaskCategoryId) {
		toSerialize["task_category_id"] = o.TaskCategoryId
	}
	toSerialize["is_archived"] = o.IsArchived
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
	toSerialize["packaging_is_standard"] = o.PackagingIsStandard
	return toSerialize, nil
}

func (o *Surveys) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"is_archived",
		"packaging_is_standard",
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

	varSurveys := _Surveys{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSurveys)

	if err != nil {
		return err
	}

	*o = Surveys(varSurveys)

	return err
}

type NullableSurveys struct {
	value *Surveys
	isSet bool
}

func (v NullableSurveys) Get() *Surveys {
	return v.value
}

func (v *NullableSurveys) Set(val *Surveys) {
	v.value = val
	v.isSet = true
}

func (v NullableSurveys) IsSet() bool {
	return v.isSet
}

func (v *NullableSurveys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSurveys(val *Surveys) *NullableSurveys {
	return &NullableSurveys{value: val, isSet: true}
}

func (v NullableSurveys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSurveys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


