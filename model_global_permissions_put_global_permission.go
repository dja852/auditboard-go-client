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

// checks if the GlobalPermissionsPutGlobalPermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalPermissionsPutGlobalPermission{}

// GlobalPermissionsPutGlobalPermission struct for GlobalPermissionsPutGlobalPermission
type GlobalPermissionsPutGlobalPermission struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	Section *string `json:"section,omitempty"`
	Key *string `json:"key,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	DefaultValue *string `json:"default_value,omitempty"`
	AllowedValues *string `json:"allowed_values,omitempty"`
	AllValues *string `json:"all_values,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	HideScopeEditor *bool `json:"hide_scope_editor,omitempty"`
	CoreModules interface{} `json:"core_modules"`
	RequireRoleLevel *string `json:"require_role_level,omitempty"`
}

type _GlobalPermissionsPutGlobalPermission GlobalPermissionsPutGlobalPermission

// NewGlobalPermissionsPutGlobalPermission instantiates a new GlobalPermissionsPutGlobalPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalPermissionsPutGlobalPermission(coreModules interface{}) *GlobalPermissionsPutGlobalPermission {
	this := GlobalPermissionsPutGlobalPermission{}
	var hideScopeEditor bool = false
	this.HideScopeEditor = &hideScopeEditor
	this.CoreModules = coreModules
	return &this
}

// NewGlobalPermissionsPutGlobalPermissionWithDefaults instantiates a new GlobalPermissionsPutGlobalPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalPermissionsPutGlobalPermissionWithDefaults() *GlobalPermissionsPutGlobalPermission {
	this := GlobalPermissionsPutGlobalPermission{}
	var hideScopeEditor bool = false
	this.HideScopeEditor = &hideScopeEditor
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GlobalPermissionsPutGlobalPermission) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionsPutGlobalPermission) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GlobalPermissionsPutGlobalPermission) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GlobalPermissionsPutGlobalPermission) SetId(v int32) {
	o.Id = &v
}

// GetSection returns the Section field value if set, zero value otherwise.
func (o *GlobalPermissionsPutGlobalPermission) GetSection() string {
	if o == nil || IsNil(o.Section) {
		var ret string
		return ret
	}
	return *o.Section
}

// GetSectionOk returns a tuple with the Section field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionsPutGlobalPermission) GetSectionOk() (*string, bool) {
	if o == nil || IsNil(o.Section) {
		return nil, false
	}
	return o.Section, true
}

// HasSection returns a boolean if a field has been set.
func (o *GlobalPermissionsPutGlobalPermission) HasSection() bool {
	if o != nil && !IsNil(o.Section) {
		return true
	}

	return false
}

// SetSection gets a reference to the given string and assigns it to the Section field.
func (o *GlobalPermissionsPutGlobalPermission) SetSection(v string) {
	o.Section = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *GlobalPermissionsPutGlobalPermission) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionsPutGlobalPermission) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *GlobalPermissionsPutGlobalPermission) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *GlobalPermissionsPutGlobalPermission) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GlobalPermissionsPutGlobalPermission) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionsPutGlobalPermission) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GlobalPermissionsPutGlobalPermission) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GlobalPermissionsPutGlobalPermission) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GlobalPermissionsPutGlobalPermission) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionsPutGlobalPermission) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GlobalPermissionsPutGlobalPermission) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GlobalPermissionsPutGlobalPermission) SetDescription(v string) {
	o.Description = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *GlobalPermissionsPutGlobalPermission) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionsPutGlobalPermission) GetDefaultValueOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *GlobalPermissionsPutGlobalPermission) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *GlobalPermissionsPutGlobalPermission) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetAllowedValues returns the AllowedValues field value if set, zero value otherwise.
func (o *GlobalPermissionsPutGlobalPermission) GetAllowedValues() string {
	if o == nil || IsNil(o.AllowedValues) {
		var ret string
		return ret
	}
	return *o.AllowedValues
}

// GetAllowedValuesOk returns a tuple with the AllowedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionsPutGlobalPermission) GetAllowedValuesOk() (*string, bool) {
	if o == nil || IsNil(o.AllowedValues) {
		return nil, false
	}
	return o.AllowedValues, true
}

// HasAllowedValues returns a boolean if a field has been set.
func (o *GlobalPermissionsPutGlobalPermission) HasAllowedValues() bool {
	if o != nil && !IsNil(o.AllowedValues) {
		return true
	}

	return false
}

// SetAllowedValues gets a reference to the given string and assigns it to the AllowedValues field.
func (o *GlobalPermissionsPutGlobalPermission) SetAllowedValues(v string) {
	o.AllowedValues = &v
}

// GetAllValues returns the AllValues field value if set, zero value otherwise.
func (o *GlobalPermissionsPutGlobalPermission) GetAllValues() string {
	if o == nil || IsNil(o.AllValues) {
		var ret string
		return ret
	}
	return *o.AllValues
}

// GetAllValuesOk returns a tuple with the AllValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionsPutGlobalPermission) GetAllValuesOk() (*string, bool) {
	if o == nil || IsNil(o.AllValues) {
		return nil, false
	}
	return o.AllValues, true
}

// HasAllValues returns a boolean if a field has been set.
func (o *GlobalPermissionsPutGlobalPermission) HasAllValues() bool {
	if o != nil && !IsNil(o.AllValues) {
		return true
	}

	return false
}

// SetAllValues gets a reference to the given string and assigns it to the AllValues field.
func (o *GlobalPermissionsPutGlobalPermission) SetAllValues(v string) {
	o.AllValues = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GlobalPermissionsPutGlobalPermission) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionsPutGlobalPermission) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GlobalPermissionsPutGlobalPermission) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GlobalPermissionsPutGlobalPermission) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GlobalPermissionsPutGlobalPermission) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionsPutGlobalPermission) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GlobalPermissionsPutGlobalPermission) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GlobalPermissionsPutGlobalPermission) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *GlobalPermissionsPutGlobalPermission) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionsPutGlobalPermission) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *GlobalPermissionsPutGlobalPermission) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *GlobalPermissionsPutGlobalPermission) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetHideScopeEditor returns the HideScopeEditor field value if set, zero value otherwise.
func (o *GlobalPermissionsPutGlobalPermission) GetHideScopeEditor() bool {
	if o == nil || IsNil(o.HideScopeEditor) {
		var ret bool
		return ret
	}
	return *o.HideScopeEditor
}

// GetHideScopeEditorOk returns a tuple with the HideScopeEditor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionsPutGlobalPermission) GetHideScopeEditorOk() (*bool, bool) {
	if o == nil || IsNil(o.HideScopeEditor) {
		return nil, false
	}
	return o.HideScopeEditor, true
}

// HasHideScopeEditor returns a boolean if a field has been set.
func (o *GlobalPermissionsPutGlobalPermission) HasHideScopeEditor() bool {
	if o != nil && !IsNil(o.HideScopeEditor) {
		return true
	}

	return false
}

// SetHideScopeEditor gets a reference to the given bool and assigns it to the HideScopeEditor field.
func (o *GlobalPermissionsPutGlobalPermission) SetHideScopeEditor(v bool) {
	o.HideScopeEditor = &v
}

// GetCoreModules returns the CoreModules field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GlobalPermissionsPutGlobalPermission) GetCoreModules() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.CoreModules
}

// GetCoreModulesOk returns a tuple with the CoreModules field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GlobalPermissionsPutGlobalPermission) GetCoreModulesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CoreModules) {
		return nil, false
	}
	return &o.CoreModules, true
}

// SetCoreModules sets field value
func (o *GlobalPermissionsPutGlobalPermission) SetCoreModules(v interface{}) {
	o.CoreModules = v
}

// GetRequireRoleLevel returns the RequireRoleLevel field value if set, zero value otherwise.
func (o *GlobalPermissionsPutGlobalPermission) GetRequireRoleLevel() string {
	if o == nil || IsNil(o.RequireRoleLevel) {
		var ret string
		return ret
	}
	return *o.RequireRoleLevel
}

// GetRequireRoleLevelOk returns a tuple with the RequireRoleLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalPermissionsPutGlobalPermission) GetRequireRoleLevelOk() (*string, bool) {
	if o == nil || IsNil(o.RequireRoleLevel) {
		return nil, false
	}
	return o.RequireRoleLevel, true
}

// HasRequireRoleLevel returns a boolean if a field has been set.
func (o *GlobalPermissionsPutGlobalPermission) HasRequireRoleLevel() bool {
	if o != nil && !IsNil(o.RequireRoleLevel) {
		return true
	}

	return false
}

// SetRequireRoleLevel gets a reference to the given string and assigns it to the RequireRoleLevel field.
func (o *GlobalPermissionsPutGlobalPermission) SetRequireRoleLevel(v string) {
	o.RequireRoleLevel = &v
}

func (o GlobalPermissionsPutGlobalPermission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalPermissionsPutGlobalPermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Section) {
		toSerialize["section"] = o.Section
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DefaultValue) {
		toSerialize["default_value"] = o.DefaultValue
	}
	if !IsNil(o.AllowedValues) {
		toSerialize["allowed_values"] = o.AllowedValues
	}
	if !IsNil(o.AllValues) {
		toSerialize["all_values"] = o.AllValues
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
	if !IsNil(o.HideScopeEditor) {
		toSerialize["hide_scope_editor"] = o.HideScopeEditor
	}
	if o.CoreModules != nil {
		toSerialize["core_modules"] = o.CoreModules
	}
	if !IsNil(o.RequireRoleLevel) {
		toSerialize["require_role_level"] = o.RequireRoleLevel
	}
	return toSerialize, nil
}

type NullableGlobalPermissionsPutGlobalPermission struct {
	value *GlobalPermissionsPutGlobalPermission
	isSet bool
}

func (v NullableGlobalPermissionsPutGlobalPermission) Get() *GlobalPermissionsPutGlobalPermission {
	return v.value
}

func (v *NullableGlobalPermissionsPutGlobalPermission) Set(val *GlobalPermissionsPutGlobalPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalPermissionsPutGlobalPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalPermissionsPutGlobalPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalPermissionsPutGlobalPermission(val *GlobalPermissionsPutGlobalPermission) *NullableGlobalPermissionsPutGlobalPermission {
	return &NullableGlobalPermissionsPutGlobalPermission{value: val, isSet: true}
}

func (v NullableGlobalPermissionsPutGlobalPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalPermissionsPutGlobalPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


