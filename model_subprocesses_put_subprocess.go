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

// checks if the SubprocessesPutSubprocess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubprocessesPutSubprocess{}

// SubprocessesPutSubprocess struct for SubprocessesPutSubprocess
type SubprocessesPutSubprocess struct {
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

type _SubprocessesPutSubprocess SubprocessesPutSubprocess

// NewSubprocessesPutSubprocess instantiates a new SubprocessesPutSubprocess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubprocessesPutSubprocess(uid string, name string, idCode string) *SubprocessesPutSubprocess {
	this := SubprocessesPutSubprocess{}
	this.Uid = uid
	this.Name = name
	this.IdCode = idCode
	return &this
}

// NewSubprocessesPutSubprocessWithDefaults instantiates a new SubprocessesPutSubprocess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubprocessesPutSubprocessWithDefaults() *SubprocessesPutSubprocess {
	this := SubprocessesPutSubprocess{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubprocessesPutSubprocess) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubprocessesPutSubprocess) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubprocessesPutSubprocess) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SubprocessesPutSubprocess) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SubprocessesPutSubprocess) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubprocessesPutSubprocess) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SubprocessesPutSubprocess) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *SubprocessesPutSubprocess) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SubprocessesPutSubprocess) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubprocessesPutSubprocess) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SubprocessesPutSubprocess) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *SubprocessesPutSubprocess) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *SubprocessesPutSubprocess) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubprocessesPutSubprocess) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *SubprocessesPutSubprocess) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *SubprocessesPutSubprocess) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetUid returns the Uid field value
func (o *SubprocessesPutSubprocess) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *SubprocessesPutSubprocess) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *SubprocessesPutSubprocess) SetUid(v string) {
	o.Uid = v
}

// GetName returns the Name field value
func (o *SubprocessesPutSubprocess) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SubprocessesPutSubprocess) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SubprocessesPutSubprocess) SetName(v string) {
	o.Name = v
}

// GetProcessId returns the ProcessId field value if set, zero value otherwise.
func (o *SubprocessesPutSubprocess) GetProcessId() int32 {
	if o == nil || IsNil(o.ProcessId) {
		var ret int32
		return ret
	}
	return *o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubprocessesPutSubprocess) GetProcessIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProcessId) {
		return nil, false
	}
	return o.ProcessId, true
}

// HasProcessId returns a boolean if a field has been set.
func (o *SubprocessesPutSubprocess) HasProcessId() bool {
	if o != nil && !IsNil(o.ProcessId) {
		return true
	}

	return false
}

// SetProcessId gets a reference to the given int32 and assigns it to the ProcessId field.
func (o *SubprocessesPutSubprocess) SetProcessId(v int32) {
	o.ProcessId = &v
}

// GetIdCode returns the IdCode field value
func (o *SubprocessesPutSubprocess) GetIdCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdCode
}

// GetIdCodeOk returns a tuple with the IdCode field value
// and a boolean to check if the value has been set.
func (o *SubprocessesPutSubprocess) GetIdCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdCode, true
}

// SetIdCode sets field value
func (o *SubprocessesPutSubprocess) SetIdCode(v string) {
	o.IdCode = v
}

func (o SubprocessesPutSubprocess) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubprocessesPutSubprocess) ToMap() (map[string]interface{}, error) {
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

type NullableSubprocessesPutSubprocess struct {
	value *SubprocessesPutSubprocess
	isSet bool
}

func (v NullableSubprocessesPutSubprocess) Get() *SubprocessesPutSubprocess {
	return v.value
}

func (v *NullableSubprocessesPutSubprocess) Set(val *SubprocessesPutSubprocess) {
	v.value = val
	v.isSet = true
}

func (v NullableSubprocessesPutSubprocess) IsSet() bool {
	return v.isSet
}

func (v *NullableSubprocessesPutSubprocess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubprocessesPutSubprocess(val *SubprocessesPutSubprocess) *NullableSubprocessesPutSubprocess {
	return &NullableSubprocessesPutSubprocess{value: val, isSet: true}
}

func (v NullableSubprocessesPutSubprocess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubprocessesPutSubprocess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


