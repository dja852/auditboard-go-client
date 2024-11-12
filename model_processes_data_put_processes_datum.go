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

// checks if the ProcessesDataPutProcessesDatum type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessesDataPutProcessesDatum{}

// ProcessesDataPutProcessesDatum struct for ProcessesDataPutProcessesDatum
type ProcessesDataPutProcessesDatum struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	File *string `json:"file,omitempty"`
	// Note: This is a Foreign Key to `processes.id`.<fk table='processes' column='id'/>
	ProcessId int32 `json:"process_id"`
	// Note: This is a Foreign Key to `entities.id`.<fk table='entities' column='id'/>
	EntityId int32 `json:"entity_id"`
	// Note: This is a Foreign Key to `auditable_entities.id`.<fk table='auditable_entities' column='id'/>
	AuditableEntityId *int32 `json:"auditable_entity_id,omitempty"`
	Objective *string `json:"objective,omitempty"`
	Scopes interface{} `json:"scopes,omitempty"`
}

type _ProcessesDataPutProcessesDatum ProcessesDataPutProcessesDatum

// NewProcessesDataPutProcessesDatum instantiates a new ProcessesDataPutProcessesDatum object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessesDataPutProcessesDatum(processId int32, entityId int32) *ProcessesDataPutProcessesDatum {
	this := ProcessesDataPutProcessesDatum{}
	this.ProcessId = processId
	this.EntityId = entityId
	return &this
}

// NewProcessesDataPutProcessesDatumWithDefaults instantiates a new ProcessesDataPutProcessesDatum object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessesDataPutProcessesDatumWithDefaults() *ProcessesDataPutProcessesDatum {
	this := ProcessesDataPutProcessesDatum{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProcessesDataPutProcessesDatum) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessesDataPutProcessesDatum) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProcessesDataPutProcessesDatum) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProcessesDataPutProcessesDatum) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ProcessesDataPutProcessesDatum) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessesDataPutProcessesDatum) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProcessesDataPutProcessesDatum) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ProcessesDataPutProcessesDatum) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ProcessesDataPutProcessesDatum) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessesDataPutProcessesDatum) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProcessesDataPutProcessesDatum) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ProcessesDataPutProcessesDatum) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ProcessesDataPutProcessesDatum) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessesDataPutProcessesDatum) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ProcessesDataPutProcessesDatum) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *ProcessesDataPutProcessesDatum) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *ProcessesDataPutProcessesDatum) GetFile() string {
	if o == nil || IsNil(o.File) {
		var ret string
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessesDataPutProcessesDatum) GetFileOk() (*string, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *ProcessesDataPutProcessesDatum) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given string and assigns it to the File field.
func (o *ProcessesDataPutProcessesDatum) SetFile(v string) {
	o.File = &v
}

// GetProcessId returns the ProcessId field value
func (o *ProcessesDataPutProcessesDatum) GetProcessId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value
// and a boolean to check if the value has been set.
func (o *ProcessesDataPutProcessesDatum) GetProcessIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessId, true
}

// SetProcessId sets field value
func (o *ProcessesDataPutProcessesDatum) SetProcessId(v int32) {
	o.ProcessId = v
}

// GetEntityId returns the EntityId field value
func (o *ProcessesDataPutProcessesDatum) GetEntityId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *ProcessesDataPutProcessesDatum) GetEntityIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *ProcessesDataPutProcessesDatum) SetEntityId(v int32) {
	o.EntityId = v
}

// GetAuditableEntityId returns the AuditableEntityId field value if set, zero value otherwise.
func (o *ProcessesDataPutProcessesDatum) GetAuditableEntityId() int32 {
	if o == nil || IsNil(o.AuditableEntityId) {
		var ret int32
		return ret
	}
	return *o.AuditableEntityId
}

// GetAuditableEntityIdOk returns a tuple with the AuditableEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessesDataPutProcessesDatum) GetAuditableEntityIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AuditableEntityId) {
		return nil, false
	}
	return o.AuditableEntityId, true
}

// HasAuditableEntityId returns a boolean if a field has been set.
func (o *ProcessesDataPutProcessesDatum) HasAuditableEntityId() bool {
	if o != nil && !IsNil(o.AuditableEntityId) {
		return true
	}

	return false
}

// SetAuditableEntityId gets a reference to the given int32 and assigns it to the AuditableEntityId field.
func (o *ProcessesDataPutProcessesDatum) SetAuditableEntityId(v int32) {
	o.AuditableEntityId = &v
}

// GetObjective returns the Objective field value if set, zero value otherwise.
func (o *ProcessesDataPutProcessesDatum) GetObjective() string {
	if o == nil || IsNil(o.Objective) {
		var ret string
		return ret
	}
	return *o.Objective
}

// GetObjectiveOk returns a tuple with the Objective field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessesDataPutProcessesDatum) GetObjectiveOk() (*string, bool) {
	if o == nil || IsNil(o.Objective) {
		return nil, false
	}
	return o.Objective, true
}

// HasObjective returns a boolean if a field has been set.
func (o *ProcessesDataPutProcessesDatum) HasObjective() bool {
	if o != nil && !IsNil(o.Objective) {
		return true
	}

	return false
}

// SetObjective gets a reference to the given string and assigns it to the Objective field.
func (o *ProcessesDataPutProcessesDatum) SetObjective(v string) {
	o.Objective = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessesDataPutProcessesDatum) GetScopes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessesDataPutProcessesDatum) GetScopesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return &o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ProcessesDataPutProcessesDatum) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given interface{} and assigns it to the Scopes field.
func (o *ProcessesDataPutProcessesDatum) SetScopes(v interface{}) {
	o.Scopes = v
}

func (o ProcessesDataPutProcessesDatum) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessesDataPutProcessesDatum) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}
	toSerialize["process_id"] = o.ProcessId
	toSerialize["entity_id"] = o.EntityId
	if !IsNil(o.AuditableEntityId) {
		toSerialize["auditable_entity_id"] = o.AuditableEntityId
	}
	if !IsNil(o.Objective) {
		toSerialize["objective"] = o.Objective
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

type NullableProcessesDataPutProcessesDatum struct {
	value *ProcessesDataPutProcessesDatum
	isSet bool
}

func (v NullableProcessesDataPutProcessesDatum) Get() *ProcessesDataPutProcessesDatum {
	return v.value
}

func (v *NullableProcessesDataPutProcessesDatum) Set(val *ProcessesDataPutProcessesDatum) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessesDataPutProcessesDatum) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessesDataPutProcessesDatum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessesDataPutProcessesDatum(val *ProcessesDataPutProcessesDatum) *NullableProcessesDataPutProcessesDatum {
	return &NullableProcessesDataPutProcessesDatum{value: val, isSet: true}
}

func (v NullableProcessesDataPutProcessesDatum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessesDataPutProcessesDatum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


