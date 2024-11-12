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

// checks if the RolePermissionsPutPreviousValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RolePermissionsPutPreviousValues{}

// RolePermissionsPutPreviousValues struct for RolePermissionsPutPreviousValues
type RolePermissionsPutPreviousValues struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	// Note: This is a Foreign Key to `roles.id`.<fk table='roles' column='id'/>
	RoleId *int32 `json:"role_id,omitempty"`
	Section *string `json:"section,omitempty"`
	Key *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	// Note: This is a Foreign Key to `global_permissions.id`.<fk table='global_permissions' column='id'/>
	GlobalPermissionId *int32 `json:"global_permission_id,omitempty"`
	Scope interface{} `json:"scope,omitempty"`
	CoreModules interface{} `json:"core_modules,omitempty"`
}

// NewRolePermissionsPutPreviousValues instantiates a new RolePermissionsPutPreviousValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolePermissionsPutPreviousValues() *RolePermissionsPutPreviousValues {
	this := RolePermissionsPutPreviousValues{}
	return &this
}

// NewRolePermissionsPutPreviousValuesWithDefaults instantiates a new RolePermissionsPutPreviousValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolePermissionsPutPreviousValuesWithDefaults() *RolePermissionsPutPreviousValues {
	this := RolePermissionsPutPreviousValues{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RolePermissionsPutPreviousValues) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolePermissionsPutPreviousValues) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RolePermissionsPutPreviousValues) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RolePermissionsPutPreviousValues) SetId(v int32) {
	o.Id = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *RolePermissionsPutPreviousValues) GetRoleId() int32 {
	if o == nil || IsNil(o.RoleId) {
		var ret int32
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolePermissionsPutPreviousValues) GetRoleIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RoleId) {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *RolePermissionsPutPreviousValues) HasRoleId() bool {
	if o != nil && !IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given int32 and assigns it to the RoleId field.
func (o *RolePermissionsPutPreviousValues) SetRoleId(v int32) {
	o.RoleId = &v
}

// GetSection returns the Section field value if set, zero value otherwise.
func (o *RolePermissionsPutPreviousValues) GetSection() string {
	if o == nil || IsNil(o.Section) {
		var ret string
		return ret
	}
	return *o.Section
}

// GetSectionOk returns a tuple with the Section field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolePermissionsPutPreviousValues) GetSectionOk() (*string, bool) {
	if o == nil || IsNil(o.Section) {
		return nil, false
	}
	return o.Section, true
}

// HasSection returns a boolean if a field has been set.
func (o *RolePermissionsPutPreviousValues) HasSection() bool {
	if o != nil && !IsNil(o.Section) {
		return true
	}

	return false
}

// SetSection gets a reference to the given string and assigns it to the Section field.
func (o *RolePermissionsPutPreviousValues) SetSection(v string) {
	o.Section = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *RolePermissionsPutPreviousValues) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolePermissionsPutPreviousValues) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *RolePermissionsPutPreviousValues) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *RolePermissionsPutPreviousValues) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RolePermissionsPutPreviousValues) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolePermissionsPutPreviousValues) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RolePermissionsPutPreviousValues) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RolePermissionsPutPreviousValues) SetValue(v string) {
	o.Value = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RolePermissionsPutPreviousValues) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolePermissionsPutPreviousValues) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RolePermissionsPutPreviousValues) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *RolePermissionsPutPreviousValues) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RolePermissionsPutPreviousValues) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolePermissionsPutPreviousValues) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RolePermissionsPutPreviousValues) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *RolePermissionsPutPreviousValues) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *RolePermissionsPutPreviousValues) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolePermissionsPutPreviousValues) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *RolePermissionsPutPreviousValues) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *RolePermissionsPutPreviousValues) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetGlobalPermissionId returns the GlobalPermissionId field value if set, zero value otherwise.
func (o *RolePermissionsPutPreviousValues) GetGlobalPermissionId() int32 {
	if o == nil || IsNil(o.GlobalPermissionId) {
		var ret int32
		return ret
	}
	return *o.GlobalPermissionId
}

// GetGlobalPermissionIdOk returns a tuple with the GlobalPermissionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolePermissionsPutPreviousValues) GetGlobalPermissionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.GlobalPermissionId) {
		return nil, false
	}
	return o.GlobalPermissionId, true
}

// HasGlobalPermissionId returns a boolean if a field has been set.
func (o *RolePermissionsPutPreviousValues) HasGlobalPermissionId() bool {
	if o != nil && !IsNil(o.GlobalPermissionId) {
		return true
	}

	return false
}

// SetGlobalPermissionId gets a reference to the given int32 and assigns it to the GlobalPermissionId field.
func (o *RolePermissionsPutPreviousValues) SetGlobalPermissionId(v int32) {
	o.GlobalPermissionId = &v
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RolePermissionsPutPreviousValues) GetScope() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RolePermissionsPutPreviousValues) GetScopeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return &o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *RolePermissionsPutPreviousValues) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given interface{} and assigns it to the Scope field.
func (o *RolePermissionsPutPreviousValues) SetScope(v interface{}) {
	o.Scope = v
}

// GetCoreModules returns the CoreModules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RolePermissionsPutPreviousValues) GetCoreModules() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CoreModules
}

// GetCoreModulesOk returns a tuple with the CoreModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RolePermissionsPutPreviousValues) GetCoreModulesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CoreModules) {
		return nil, false
	}
	return &o.CoreModules, true
}

// HasCoreModules returns a boolean if a field has been set.
func (o *RolePermissionsPutPreviousValues) HasCoreModules() bool {
	if o != nil && !IsNil(o.CoreModules) {
		return true
	}

	return false
}

// SetCoreModules gets a reference to the given interface{} and assigns it to the CoreModules field.
func (o *RolePermissionsPutPreviousValues) SetCoreModules(v interface{}) {
	o.CoreModules = v
}

func (o RolePermissionsPutPreviousValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RolePermissionsPutPreviousValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RoleId) {
		toSerialize["role_id"] = o.RoleId
	}
	if !IsNil(o.Section) {
		toSerialize["section"] = o.Section
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
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
	if !IsNil(o.GlobalPermissionId) {
		toSerialize["global_permission_id"] = o.GlobalPermissionId
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.CoreModules != nil {
		toSerialize["core_modules"] = o.CoreModules
	}
	return toSerialize, nil
}

type NullableRolePermissionsPutPreviousValues struct {
	value *RolePermissionsPutPreviousValues
	isSet bool
}

func (v NullableRolePermissionsPutPreviousValues) Get() *RolePermissionsPutPreviousValues {
	return v.value
}

func (v *NullableRolePermissionsPutPreviousValues) Set(val *RolePermissionsPutPreviousValues) {
	v.value = val
	v.isSet = true
}

func (v NullableRolePermissionsPutPreviousValues) IsSet() bool {
	return v.isSet
}

func (v *NullableRolePermissionsPutPreviousValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolePermissionsPutPreviousValues(val *RolePermissionsPutPreviousValues) *NullableRolePermissionsPutPreviousValues {
	return &NullableRolePermissionsPutPreviousValues{value: val, isSet: true}
}

func (v NullableRolePermissionsPutPreviousValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolePermissionsPutPreviousValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


