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

// checks if the TeamPermissionsPutTeamPermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamPermissionsPutTeamPermission{}

// TeamPermissionsPutTeamPermission struct for TeamPermissionsPutTeamPermission
type TeamPermissionsPutTeamPermission struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	// Note: This is a Foreign Key to `teams.id`.<fk table='teams' column='id'/>
	TeamId int32 `json:"team_id"`
	// Note: This is a Foreign Key to `global_permissions.id`.<fk table='global_permissions' column='id'/>
	GlobalPermissionId int32 `json:"global_permission_id"`
	Section string `json:"section"`
	Key string `json:"key"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

type _TeamPermissionsPutTeamPermission TeamPermissionsPutTeamPermission

// NewTeamPermissionsPutTeamPermission instantiates a new TeamPermissionsPutTeamPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamPermissionsPutTeamPermission(teamId int32, globalPermissionId int32, section string, key string) *TeamPermissionsPutTeamPermission {
	this := TeamPermissionsPutTeamPermission{}
	this.TeamId = teamId
	this.GlobalPermissionId = globalPermissionId
	this.Section = section
	this.Key = key
	return &this
}

// NewTeamPermissionsPutTeamPermissionWithDefaults instantiates a new TeamPermissionsPutTeamPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamPermissionsPutTeamPermissionWithDefaults() *TeamPermissionsPutTeamPermission {
	this := TeamPermissionsPutTeamPermission{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TeamPermissionsPutTeamPermission) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPermissionsPutTeamPermission) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TeamPermissionsPutTeamPermission) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TeamPermissionsPutTeamPermission) SetId(v int32) {
	o.Id = &v
}

// GetTeamId returns the TeamId field value
func (o *TeamPermissionsPutTeamPermission) GetTeamId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *TeamPermissionsPutTeamPermission) GetTeamIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *TeamPermissionsPutTeamPermission) SetTeamId(v int32) {
	o.TeamId = v
}

// GetGlobalPermissionId returns the GlobalPermissionId field value
func (o *TeamPermissionsPutTeamPermission) GetGlobalPermissionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GlobalPermissionId
}

// GetGlobalPermissionIdOk returns a tuple with the GlobalPermissionId field value
// and a boolean to check if the value has been set.
func (o *TeamPermissionsPutTeamPermission) GetGlobalPermissionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalPermissionId, true
}

// SetGlobalPermissionId sets field value
func (o *TeamPermissionsPutTeamPermission) SetGlobalPermissionId(v int32) {
	o.GlobalPermissionId = v
}

// GetSection returns the Section field value
func (o *TeamPermissionsPutTeamPermission) GetSection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Section
}

// GetSectionOk returns a tuple with the Section field value
// and a boolean to check if the value has been set.
func (o *TeamPermissionsPutTeamPermission) GetSectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Section, true
}

// SetSection sets field value
func (o *TeamPermissionsPutTeamPermission) SetSection(v string) {
	o.Section = v
}

// GetKey returns the Key field value
func (o *TeamPermissionsPutTeamPermission) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *TeamPermissionsPutTeamPermission) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *TeamPermissionsPutTeamPermission) SetKey(v string) {
	o.Key = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TeamPermissionsPutTeamPermission) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPermissionsPutTeamPermission) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TeamPermissionsPutTeamPermission) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *TeamPermissionsPutTeamPermission) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TeamPermissionsPutTeamPermission) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPermissionsPutTeamPermission) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TeamPermissionsPutTeamPermission) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *TeamPermissionsPutTeamPermission) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o TeamPermissionsPutTeamPermission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamPermissionsPutTeamPermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["team_id"] = o.TeamId
	toSerialize["global_permission_id"] = o.GlobalPermissionId
	toSerialize["section"] = o.Section
	toSerialize["key"] = o.Key
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableTeamPermissionsPutTeamPermission struct {
	value *TeamPermissionsPutTeamPermission
	isSet bool
}

func (v NullableTeamPermissionsPutTeamPermission) Get() *TeamPermissionsPutTeamPermission {
	return v.value
}

func (v *NullableTeamPermissionsPutTeamPermission) Set(val *TeamPermissionsPutTeamPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamPermissionsPutTeamPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamPermissionsPutTeamPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamPermissionsPutTeamPermission(val *TeamPermissionsPutTeamPermission) *NullableTeamPermissionsPutTeamPermission {
	return &NullableTeamPermissionsPutTeamPermission{value: val, isSet: true}
}

func (v NullableTeamPermissionsPutTeamPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamPermissionsPutTeamPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


