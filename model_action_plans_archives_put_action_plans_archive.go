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

// checks if the ActionPlansArchivesPutActionPlansArchive type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionPlansArchivesPutActionPlansArchive{}

// ActionPlansArchivesPutActionPlansArchive struct for ActionPlansArchivesPutActionPlansArchive
type ActionPlansArchivesPutActionPlansArchive struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	// Note: This is a Foreign Key to `action_plans.id`.<fk table='action_plans' column='id'/>
	ActionPlanId int32 `json:"action_plan_id"`
	// Note: This is a Foreign Key to `issues_archives.id`.<fk table='issues_archives' column='id'/>
	IssuesArchiveId int32 `json:"issues_archive_id"`
	// Note: This is a Foreign Key to `archives.id`.<fk table='archives' column='id'/>
	ArchiveId int32 `json:"archive_id"`
	// Note: This is a Foreign Key to `issues.id`.<fk table='issues' column='id'/>
	IssueId int32 `json:"issue_id"`
	Flattened interface{} `json:"flattened"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}

type _ActionPlansArchivesPutActionPlansArchive ActionPlansArchivesPutActionPlansArchive

// NewActionPlansArchivesPutActionPlansArchive instantiates a new ActionPlansArchivesPutActionPlansArchive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionPlansArchivesPutActionPlansArchive(actionPlanId int32, issuesArchiveId int32, archiveId int32, issueId int32, flattened interface{}) *ActionPlansArchivesPutActionPlansArchive {
	this := ActionPlansArchivesPutActionPlansArchive{}
	this.ActionPlanId = actionPlanId
	this.IssuesArchiveId = issuesArchiveId
	this.ArchiveId = archiveId
	this.IssueId = issueId
	this.Flattened = flattened
	return &this
}

// NewActionPlansArchivesPutActionPlansArchiveWithDefaults instantiates a new ActionPlansArchivesPutActionPlansArchive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionPlansArchivesPutActionPlansArchiveWithDefaults() *ActionPlansArchivesPutActionPlansArchive {
	this := ActionPlansArchivesPutActionPlansArchive{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ActionPlansArchivesPutActionPlansArchive) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionPlansArchivesPutActionPlansArchive) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ActionPlansArchivesPutActionPlansArchive) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ActionPlansArchivesPutActionPlansArchive) SetId(v int32) {
	o.Id = &v
}

// GetActionPlanId returns the ActionPlanId field value
func (o *ActionPlansArchivesPutActionPlansArchive) GetActionPlanId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ActionPlanId
}

// GetActionPlanIdOk returns a tuple with the ActionPlanId field value
// and a boolean to check if the value has been set.
func (o *ActionPlansArchivesPutActionPlansArchive) GetActionPlanIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionPlanId, true
}

// SetActionPlanId sets field value
func (o *ActionPlansArchivesPutActionPlansArchive) SetActionPlanId(v int32) {
	o.ActionPlanId = v
}

// GetIssuesArchiveId returns the IssuesArchiveId field value
func (o *ActionPlansArchivesPutActionPlansArchive) GetIssuesArchiveId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IssuesArchiveId
}

// GetIssuesArchiveIdOk returns a tuple with the IssuesArchiveId field value
// and a boolean to check if the value has been set.
func (o *ActionPlansArchivesPutActionPlansArchive) GetIssuesArchiveIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuesArchiveId, true
}

// SetIssuesArchiveId sets field value
func (o *ActionPlansArchivesPutActionPlansArchive) SetIssuesArchiveId(v int32) {
	o.IssuesArchiveId = v
}

// GetArchiveId returns the ArchiveId field value
func (o *ActionPlansArchivesPutActionPlansArchive) GetArchiveId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ArchiveId
}

// GetArchiveIdOk returns a tuple with the ArchiveId field value
// and a boolean to check if the value has been set.
func (o *ActionPlansArchivesPutActionPlansArchive) GetArchiveIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArchiveId, true
}

// SetArchiveId sets field value
func (o *ActionPlansArchivesPutActionPlansArchive) SetArchiveId(v int32) {
	o.ArchiveId = v
}

// GetIssueId returns the IssueId field value
func (o *ActionPlansArchivesPutActionPlansArchive) GetIssueId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IssueId
}

// GetIssueIdOk returns a tuple with the IssueId field value
// and a boolean to check if the value has been set.
func (o *ActionPlansArchivesPutActionPlansArchive) GetIssueIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueId, true
}

// SetIssueId sets field value
func (o *ActionPlansArchivesPutActionPlansArchive) SetIssueId(v int32) {
	o.IssueId = v
}

// GetFlattened returns the Flattened field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ActionPlansArchivesPutActionPlansArchive) GetFlattened() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Flattened
}

// GetFlattenedOk returns a tuple with the Flattened field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionPlansArchivesPutActionPlansArchive) GetFlattenedOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Flattened) {
		return nil, false
	}
	return &o.Flattened, true
}

// SetFlattened sets field value
func (o *ActionPlansArchivesPutActionPlansArchive) SetFlattened(v interface{}) {
	o.Flattened = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ActionPlansArchivesPutActionPlansArchive) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionPlansArchivesPutActionPlansArchive) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ActionPlansArchivesPutActionPlansArchive) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ActionPlansArchivesPutActionPlansArchive) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ActionPlansArchivesPutActionPlansArchive) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionPlansArchivesPutActionPlansArchive) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ActionPlansArchivesPutActionPlansArchive) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ActionPlansArchivesPutActionPlansArchive) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ActionPlansArchivesPutActionPlansArchive) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionPlansArchivesPutActionPlansArchive) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ActionPlansArchivesPutActionPlansArchive) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *ActionPlansArchivesPutActionPlansArchive) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

func (o ActionPlansArchivesPutActionPlansArchive) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionPlansArchivesPutActionPlansArchive) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["action_plan_id"] = o.ActionPlanId
	toSerialize["issues_archive_id"] = o.IssuesArchiveId
	toSerialize["archive_id"] = o.ArchiveId
	toSerialize["issue_id"] = o.IssueId
	if o.Flattened != nil {
		toSerialize["flattened"] = o.Flattened
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
	return toSerialize, nil
}

type NullableActionPlansArchivesPutActionPlansArchive struct {
	value *ActionPlansArchivesPutActionPlansArchive
	isSet bool
}

func (v NullableActionPlansArchivesPutActionPlansArchive) Get() *ActionPlansArchivesPutActionPlansArchive {
	return v.value
}

func (v *NullableActionPlansArchivesPutActionPlansArchive) Set(val *ActionPlansArchivesPutActionPlansArchive) {
	v.value = val
	v.isSet = true
}

func (v NullableActionPlansArchivesPutActionPlansArchive) IsSet() bool {
	return v.isSet
}

func (v *NullableActionPlansArchivesPutActionPlansArchive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionPlansArchivesPutActionPlansArchive(val *ActionPlansArchivesPutActionPlansArchive) *NullableActionPlansArchivesPutActionPlansArchive {
	return &NullableActionPlansArchivesPutActionPlansArchive{value: val, isSet: true}
}

func (v NullableActionPlansArchivesPutActionPlansArchive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionPlansArchivesPutActionPlansArchive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


