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

// checks if the PoliciesPutPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoliciesPutPolicy{}

// PoliciesPutPolicy struct for PoliciesPutPolicy
type PoliciesPutPolicy struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	Uid string `json:"uid"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Category *string `json:"category,omitempty"`
	// Note: This is a Foreign Key to `document_versions.id`.<fk table='document_versions' column='id'/>
	LatestDocumentVersionId *int32 `json:"latest_document_version_id,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	LastRevisionByUserId *int32 `json:"last_revision_by_user_id,omitempty"`
	LastRevisionDate *string `json:"last_revision_date,omitempty"`
	// Note: This is a Foreign Key to `policy_types.id`.<fk table='policy_types' column='id'/>
	PolicyTypeId int32 `json:"policy_type_id"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	AdminUserId *int32 `json:"admin_user_id,omitempty"`
	IsSequentialApproval *bool `json:"is_sequential_approval,omitempty"`
	FieldData interface{} `json:"field_data,omitempty"`
	// Note: This is a Foreign Key to `recurring_schedules.id`.<fk table='recurring_schedules' column='id'/>
	RecurringScheduleId *int32 `json:"recurring_schedule_id,omitempty"`
	FutureApprovalDate *string `json:"future_approval_date,omitempty"`
	ReminderDuration string `json:"reminder_duration"`
	RecurrenceNotifiedAt *string `json:"recurrence_notified_at,omitempty"`
	LastApprovalDate *string `json:"last_approval_date,omitempty"`
	Scopes interface{} `json:"scopes,omitempty"`
}

type _PoliciesPutPolicy PoliciesPutPolicy

// NewPoliciesPutPolicy instantiates a new PoliciesPutPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoliciesPutPolicy(uid string, name string, policyTypeId int32, reminderDuration string) *PoliciesPutPolicy {
	this := PoliciesPutPolicy{}
	this.Uid = uid
	this.Name = name
	this.PolicyTypeId = policyTypeId
	this.ReminderDuration = reminderDuration
	return &this
}

// NewPoliciesPutPolicyWithDefaults instantiates a new PoliciesPutPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoliciesPutPolicyWithDefaults() *PoliciesPutPolicy {
	this := PoliciesPutPolicy{}
	var reminderDuration string = "30 Days"
	this.ReminderDuration = reminderDuration
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PoliciesPutPolicy) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesPutPolicy) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PoliciesPutPolicy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PoliciesPutPolicy) SetId(v int32) {
	o.Id = &v
}

// GetUid returns the Uid field value
func (o *PoliciesPutPolicy) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *PoliciesPutPolicy) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *PoliciesPutPolicy) SetUid(v string) {
	o.Uid = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PoliciesPutPolicy) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesPutPolicy) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PoliciesPutPolicy) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *PoliciesPutPolicy) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PoliciesPutPolicy) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesPutPolicy) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PoliciesPutPolicy) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *PoliciesPutPolicy) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *PoliciesPutPolicy) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesPutPolicy) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *PoliciesPutPolicy) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *PoliciesPutPolicy) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetName returns the Name field value
func (o *PoliciesPutPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PoliciesPutPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PoliciesPutPolicy) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PoliciesPutPolicy) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesPutPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PoliciesPutPolicy) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PoliciesPutPolicy) SetDescription(v string) {
	o.Description = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *PoliciesPutPolicy) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesPutPolicy) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *PoliciesPutPolicy) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *PoliciesPutPolicy) SetCategory(v string) {
	o.Category = &v
}

// GetLatestDocumentVersionId returns the LatestDocumentVersionId field value if set, zero value otherwise.
func (o *PoliciesPutPolicy) GetLatestDocumentVersionId() int32 {
	if o == nil || IsNil(o.LatestDocumentVersionId) {
		var ret int32
		return ret
	}
	return *o.LatestDocumentVersionId
}

// GetLatestDocumentVersionIdOk returns a tuple with the LatestDocumentVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesPutPolicy) GetLatestDocumentVersionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LatestDocumentVersionId) {
		return nil, false
	}
	return o.LatestDocumentVersionId, true
}

// HasLatestDocumentVersionId returns a boolean if a field has been set.
func (o *PoliciesPutPolicy) HasLatestDocumentVersionId() bool {
	if o != nil && !IsNil(o.LatestDocumentVersionId) {
		return true
	}

	return false
}

// SetLatestDocumentVersionId gets a reference to the given int32 and assigns it to the LatestDocumentVersionId field.
func (o *PoliciesPutPolicy) SetLatestDocumentVersionId(v int32) {
	o.LatestDocumentVersionId = &v
}

// GetLastRevisionByUserId returns the LastRevisionByUserId field value if set, zero value otherwise.
func (o *PoliciesPutPolicy) GetLastRevisionByUserId() int32 {
	if o == nil || IsNil(o.LastRevisionByUserId) {
		var ret int32
		return ret
	}
	return *o.LastRevisionByUserId
}

// GetLastRevisionByUserIdOk returns a tuple with the LastRevisionByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesPutPolicy) GetLastRevisionByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LastRevisionByUserId) {
		return nil, false
	}
	return o.LastRevisionByUserId, true
}

// HasLastRevisionByUserId returns a boolean if a field has been set.
func (o *PoliciesPutPolicy) HasLastRevisionByUserId() bool {
	if o != nil && !IsNil(o.LastRevisionByUserId) {
		return true
	}

	return false
}

// SetLastRevisionByUserId gets a reference to the given int32 and assigns it to the LastRevisionByUserId field.
func (o *PoliciesPutPolicy) SetLastRevisionByUserId(v int32) {
	o.LastRevisionByUserId = &v
}

// GetLastRevisionDate returns the LastRevisionDate field value if set, zero value otherwise.
func (o *PoliciesPutPolicy) GetLastRevisionDate() string {
	if o == nil || IsNil(o.LastRevisionDate) {
		var ret string
		return ret
	}
	return *o.LastRevisionDate
}

// GetLastRevisionDateOk returns a tuple with the LastRevisionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesPutPolicy) GetLastRevisionDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastRevisionDate) {
		return nil, false
	}
	return o.LastRevisionDate, true
}

// HasLastRevisionDate returns a boolean if a field has been set.
func (o *PoliciesPutPolicy) HasLastRevisionDate() bool {
	if o != nil && !IsNil(o.LastRevisionDate) {
		return true
	}

	return false
}

// SetLastRevisionDate gets a reference to the given string and assigns it to the LastRevisionDate field.
func (o *PoliciesPutPolicy) SetLastRevisionDate(v string) {
	o.LastRevisionDate = &v
}

// GetPolicyTypeId returns the PolicyTypeId field value
func (o *PoliciesPutPolicy) GetPolicyTypeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PolicyTypeId
}

// GetPolicyTypeIdOk returns a tuple with the PolicyTypeId field value
// and a boolean to check if the value has been set.
func (o *PoliciesPutPolicy) GetPolicyTypeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyTypeId, true
}

// SetPolicyTypeId sets field value
func (o *PoliciesPutPolicy) SetPolicyTypeId(v int32) {
	o.PolicyTypeId = v
}

// GetAdminUserId returns the AdminUserId field value if set, zero value otherwise.
func (o *PoliciesPutPolicy) GetAdminUserId() int32 {
	if o == nil || IsNil(o.AdminUserId) {
		var ret int32
		return ret
	}
	return *o.AdminUserId
}

// GetAdminUserIdOk returns a tuple with the AdminUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesPutPolicy) GetAdminUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AdminUserId) {
		return nil, false
	}
	return o.AdminUserId, true
}

// HasAdminUserId returns a boolean if a field has been set.
func (o *PoliciesPutPolicy) HasAdminUserId() bool {
	if o != nil && !IsNil(o.AdminUserId) {
		return true
	}

	return false
}

// SetAdminUserId gets a reference to the given int32 and assigns it to the AdminUserId field.
func (o *PoliciesPutPolicy) SetAdminUserId(v int32) {
	o.AdminUserId = &v
}

// GetIsSequentialApproval returns the IsSequentialApproval field value if set, zero value otherwise.
func (o *PoliciesPutPolicy) GetIsSequentialApproval() bool {
	if o == nil || IsNil(o.IsSequentialApproval) {
		var ret bool
		return ret
	}
	return *o.IsSequentialApproval
}

// GetIsSequentialApprovalOk returns a tuple with the IsSequentialApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesPutPolicy) GetIsSequentialApprovalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSequentialApproval) {
		return nil, false
	}
	return o.IsSequentialApproval, true
}

// HasIsSequentialApproval returns a boolean if a field has been set.
func (o *PoliciesPutPolicy) HasIsSequentialApproval() bool {
	if o != nil && !IsNil(o.IsSequentialApproval) {
		return true
	}

	return false
}

// SetIsSequentialApproval gets a reference to the given bool and assigns it to the IsSequentialApproval field.
func (o *PoliciesPutPolicy) SetIsSequentialApproval(v bool) {
	o.IsSequentialApproval = &v
}

// GetFieldData returns the FieldData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoliciesPutPolicy) GetFieldData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FieldData
}

// GetFieldDataOk returns a tuple with the FieldData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoliciesPutPolicy) GetFieldDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FieldData) {
		return nil, false
	}
	return &o.FieldData, true
}

// HasFieldData returns a boolean if a field has been set.
func (o *PoliciesPutPolicy) HasFieldData() bool {
	if o != nil && !IsNil(o.FieldData) {
		return true
	}

	return false
}

// SetFieldData gets a reference to the given interface{} and assigns it to the FieldData field.
func (o *PoliciesPutPolicy) SetFieldData(v interface{}) {
	o.FieldData = v
}

// GetRecurringScheduleId returns the RecurringScheduleId field value if set, zero value otherwise.
func (o *PoliciesPutPolicy) GetRecurringScheduleId() int32 {
	if o == nil || IsNil(o.RecurringScheduleId) {
		var ret int32
		return ret
	}
	return *o.RecurringScheduleId
}

// GetRecurringScheduleIdOk returns a tuple with the RecurringScheduleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesPutPolicy) GetRecurringScheduleIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RecurringScheduleId) {
		return nil, false
	}
	return o.RecurringScheduleId, true
}

// HasRecurringScheduleId returns a boolean if a field has been set.
func (o *PoliciesPutPolicy) HasRecurringScheduleId() bool {
	if o != nil && !IsNil(o.RecurringScheduleId) {
		return true
	}

	return false
}

// SetRecurringScheduleId gets a reference to the given int32 and assigns it to the RecurringScheduleId field.
func (o *PoliciesPutPolicy) SetRecurringScheduleId(v int32) {
	o.RecurringScheduleId = &v
}

// GetFutureApprovalDate returns the FutureApprovalDate field value if set, zero value otherwise.
func (o *PoliciesPutPolicy) GetFutureApprovalDate() string {
	if o == nil || IsNil(o.FutureApprovalDate) {
		var ret string
		return ret
	}
	return *o.FutureApprovalDate
}

// GetFutureApprovalDateOk returns a tuple with the FutureApprovalDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesPutPolicy) GetFutureApprovalDateOk() (*string, bool) {
	if o == nil || IsNil(o.FutureApprovalDate) {
		return nil, false
	}
	return o.FutureApprovalDate, true
}

// HasFutureApprovalDate returns a boolean if a field has been set.
func (o *PoliciesPutPolicy) HasFutureApprovalDate() bool {
	if o != nil && !IsNil(o.FutureApprovalDate) {
		return true
	}

	return false
}

// SetFutureApprovalDate gets a reference to the given string and assigns it to the FutureApprovalDate field.
func (o *PoliciesPutPolicy) SetFutureApprovalDate(v string) {
	o.FutureApprovalDate = &v
}

// GetReminderDuration returns the ReminderDuration field value
func (o *PoliciesPutPolicy) GetReminderDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReminderDuration
}

// GetReminderDurationOk returns a tuple with the ReminderDuration field value
// and a boolean to check if the value has been set.
func (o *PoliciesPutPolicy) GetReminderDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReminderDuration, true
}

// SetReminderDuration sets field value
func (o *PoliciesPutPolicy) SetReminderDuration(v string) {
	o.ReminderDuration = v
}

// GetRecurrenceNotifiedAt returns the RecurrenceNotifiedAt field value if set, zero value otherwise.
func (o *PoliciesPutPolicy) GetRecurrenceNotifiedAt() string {
	if o == nil || IsNil(o.RecurrenceNotifiedAt) {
		var ret string
		return ret
	}
	return *o.RecurrenceNotifiedAt
}

// GetRecurrenceNotifiedAtOk returns a tuple with the RecurrenceNotifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesPutPolicy) GetRecurrenceNotifiedAtOk() (*string, bool) {
	if o == nil || IsNil(o.RecurrenceNotifiedAt) {
		return nil, false
	}
	return o.RecurrenceNotifiedAt, true
}

// HasRecurrenceNotifiedAt returns a boolean if a field has been set.
func (o *PoliciesPutPolicy) HasRecurrenceNotifiedAt() bool {
	if o != nil && !IsNil(o.RecurrenceNotifiedAt) {
		return true
	}

	return false
}

// SetRecurrenceNotifiedAt gets a reference to the given string and assigns it to the RecurrenceNotifiedAt field.
func (o *PoliciesPutPolicy) SetRecurrenceNotifiedAt(v string) {
	o.RecurrenceNotifiedAt = &v
}

// GetLastApprovalDate returns the LastApprovalDate field value if set, zero value otherwise.
func (o *PoliciesPutPolicy) GetLastApprovalDate() string {
	if o == nil || IsNil(o.LastApprovalDate) {
		var ret string
		return ret
	}
	return *o.LastApprovalDate
}

// GetLastApprovalDateOk returns a tuple with the LastApprovalDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesPutPolicy) GetLastApprovalDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastApprovalDate) {
		return nil, false
	}
	return o.LastApprovalDate, true
}

// HasLastApprovalDate returns a boolean if a field has been set.
func (o *PoliciesPutPolicy) HasLastApprovalDate() bool {
	if o != nil && !IsNil(o.LastApprovalDate) {
		return true
	}

	return false
}

// SetLastApprovalDate gets a reference to the given string and assigns it to the LastApprovalDate field.
func (o *PoliciesPutPolicy) SetLastApprovalDate(v string) {
	o.LastApprovalDate = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoliciesPutPolicy) GetScopes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoliciesPutPolicy) GetScopesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return &o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *PoliciesPutPolicy) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given interface{} and assigns it to the Scopes field.
func (o *PoliciesPutPolicy) SetScopes(v interface{}) {
	o.Scopes = v
}

func (o PoliciesPutPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoliciesPutPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["uid"] = o.Uid
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
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.LatestDocumentVersionId) {
		toSerialize["latest_document_version_id"] = o.LatestDocumentVersionId
	}
	if !IsNil(o.LastRevisionByUserId) {
		toSerialize["last_revision_by_user_id"] = o.LastRevisionByUserId
	}
	if !IsNil(o.LastRevisionDate) {
		toSerialize["last_revision_date"] = o.LastRevisionDate
	}
	toSerialize["policy_type_id"] = o.PolicyTypeId
	if !IsNil(o.AdminUserId) {
		toSerialize["admin_user_id"] = o.AdminUserId
	}
	if !IsNil(o.IsSequentialApproval) {
		toSerialize["is_sequential_approval"] = o.IsSequentialApproval
	}
	if o.FieldData != nil {
		toSerialize["field_data"] = o.FieldData
	}
	if !IsNil(o.RecurringScheduleId) {
		toSerialize["recurring_schedule_id"] = o.RecurringScheduleId
	}
	if !IsNil(o.FutureApprovalDate) {
		toSerialize["future_approval_date"] = o.FutureApprovalDate
	}
	toSerialize["reminder_duration"] = o.ReminderDuration
	if !IsNil(o.RecurrenceNotifiedAt) {
		toSerialize["recurrence_notified_at"] = o.RecurrenceNotifiedAt
	}
	if !IsNil(o.LastApprovalDate) {
		toSerialize["last_approval_date"] = o.LastApprovalDate
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

type NullablePoliciesPutPolicy struct {
	value *PoliciesPutPolicy
	isSet bool
}

func (v NullablePoliciesPutPolicy) Get() *PoliciesPutPolicy {
	return v.value
}

func (v *NullablePoliciesPutPolicy) Set(val *PoliciesPutPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullablePoliciesPutPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullablePoliciesPutPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoliciesPutPolicy(val *PoliciesPutPolicy) *NullablePoliciesPutPolicy {
	return &NullablePoliciesPutPolicy{value: val, isSet: true}
}

func (v NullablePoliciesPutPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoliciesPutPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


