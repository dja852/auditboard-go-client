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

// checks if the Subprocesses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Subprocesses{}

// Subprocesses struct for Subprocesses
type Subprocesses struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Uid string `json:"uid"`
	Name string `json:"name"`
	// Note: This is a Foreign Key to `processes.id`.<fk table='processes' column='id'/>
	ProcessId *int32 `json:"process_id,omitempty"`
	IdCode string `json:"id_code"`
}

type _Subprocesses Subprocesses

// NewSubprocesses instantiates a new Subprocesses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubprocesses(uid string, name string, idCode string) *Subprocesses {
	this := Subprocesses{}
	this.Uid = uid
	this.Name = name
	this.IdCode = idCode
	return &this
}

// NewSubprocessesWithDefaults instantiates a new Subprocesses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubprocessesWithDefaults() *Subprocesses {
	this := Subprocesses{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Subprocesses) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subprocesses) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Subprocesses) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Subprocesses) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Subprocesses) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subprocesses) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Subprocesses) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Subprocesses) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Subprocesses) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subprocesses) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Subprocesses) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Subprocesses) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Subprocesses) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subprocesses) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Subprocesses) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *Subprocesses) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetUid returns the Uid field value
func (o *Subprocesses) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *Subprocesses) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *Subprocesses) SetUid(v string) {
	o.Uid = v
}

// GetName returns the Name field value
func (o *Subprocesses) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Subprocesses) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Subprocesses) SetName(v string) {
	o.Name = v
}

// GetProcessId returns the ProcessId field value if set, zero value otherwise.
func (o *Subprocesses) GetProcessId() int32 {
	if o == nil || IsNil(o.ProcessId) {
		var ret int32
		return ret
	}
	return *o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subprocesses) GetProcessIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProcessId) {
		return nil, false
	}
	return o.ProcessId, true
}

// HasProcessId returns a boolean if a field has been set.
func (o *Subprocesses) HasProcessId() bool {
	if o != nil && !IsNil(o.ProcessId) {
		return true
	}

	return false
}

// SetProcessId gets a reference to the given int32 and assigns it to the ProcessId field.
func (o *Subprocesses) SetProcessId(v int32) {
	o.ProcessId = &v
}

// GetIdCode returns the IdCode field value
func (o *Subprocesses) GetIdCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdCode
}

// GetIdCodeOk returns a tuple with the IdCode field value
// and a boolean to check if the value has been set.
func (o *Subprocesses) GetIdCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdCode, true
}

// SetIdCode sets field value
func (o *Subprocesses) SetIdCode(v string) {
	o.IdCode = v
}

func (o Subprocesses) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Subprocesses) ToMap() (map[string]interface{}, error) {
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
	toSerialize["uid"] = o.Uid
	toSerialize["name"] = o.Name
	if !IsNil(o.ProcessId) {
		toSerialize["process_id"] = o.ProcessId
	}
	toSerialize["id_code"] = o.IdCode
	return toSerialize, nil
}

type NullableSubprocesses struct {
	value *Subprocesses
	isSet bool
}

func (v NullableSubprocesses) Get() *Subprocesses {
	return v.value
}

func (v *NullableSubprocesses) Set(val *Subprocesses) {
	v.value = val
	v.isSet = true
}

func (v NullableSubprocesses) IsSet() bool {
	return v.isSet
}

func (v *NullableSubprocesses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubprocesses(val *Subprocesses) *NullableSubprocesses {
	return &NullableSubprocesses{value: val, isSet: true}
}

func (v NullableSubprocesses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubprocesses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


