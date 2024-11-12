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

// checks if the TaskCategories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskCategories{}

// TaskCategories struct for TaskCategories
type TaskCategories struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Name string `json:"name"`
	SortOrder int32 `json:"sort_order"`
	MetaConfig *string `json:"meta_config,omitempty"`
	Color *string `json:"color,omitempty"`
	ShortName *string `json:"short_name,omitempty"`
	Scopes interface{} `json:"scopes,omitempty"`
}

type _TaskCategories TaskCategories

// NewTaskCategories instantiates a new TaskCategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskCategories(name string, sortOrder int32) *TaskCategories {
	this := TaskCategories{}
	this.Name = name
	this.SortOrder = sortOrder
	return &this
}

// NewTaskCategoriesWithDefaults instantiates a new TaskCategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskCategoriesWithDefaults() *TaskCategories {
	this := TaskCategories{}
	var sortOrder int32 = 0
	this.SortOrder = sortOrder
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaskCategories) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCategories) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TaskCategories) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TaskCategories) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TaskCategories) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCategories) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TaskCategories) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *TaskCategories) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TaskCategories) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCategories) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TaskCategories) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *TaskCategories) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *TaskCategories) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCategories) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *TaskCategories) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *TaskCategories) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetName returns the Name field value
func (o *TaskCategories) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TaskCategories) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TaskCategories) SetName(v string) {
	o.Name = v
}

// GetSortOrder returns the SortOrder field value
func (o *TaskCategories) GetSortOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value
// and a boolean to check if the value has been set.
func (o *TaskCategories) GetSortOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortOrder, true
}

// SetSortOrder sets field value
func (o *TaskCategories) SetSortOrder(v int32) {
	o.SortOrder = v
}

// GetMetaConfig returns the MetaConfig field value if set, zero value otherwise.
func (o *TaskCategories) GetMetaConfig() string {
	if o == nil || IsNil(o.MetaConfig) {
		var ret string
		return ret
	}
	return *o.MetaConfig
}

// GetMetaConfigOk returns a tuple with the MetaConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCategories) GetMetaConfigOk() (*string, bool) {
	if o == nil || IsNil(o.MetaConfig) {
		return nil, false
	}
	return o.MetaConfig, true
}

// HasMetaConfig returns a boolean if a field has been set.
func (o *TaskCategories) HasMetaConfig() bool {
	if o != nil && !IsNil(o.MetaConfig) {
		return true
	}

	return false
}

// SetMetaConfig gets a reference to the given string and assigns it to the MetaConfig field.
func (o *TaskCategories) SetMetaConfig(v string) {
	o.MetaConfig = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *TaskCategories) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCategories) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *TaskCategories) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *TaskCategories) SetColor(v string) {
	o.Color = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *TaskCategories) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCategories) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *TaskCategories) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *TaskCategories) SetShortName(v string) {
	o.ShortName = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskCategories) GetScopes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskCategories) GetScopesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return &o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *TaskCategories) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given interface{} and assigns it to the Scopes field.
func (o *TaskCategories) SetScopes(v interface{}) {
	o.Scopes = v
}

func (o TaskCategories) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskCategories) ToMap() (map[string]interface{}, error) {
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
	toSerialize["name"] = o.Name
	toSerialize["sort_order"] = o.SortOrder
	if !IsNil(o.MetaConfig) {
		toSerialize["meta_config"] = o.MetaConfig
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.ShortName) {
		toSerialize["short_name"] = o.ShortName
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

func (o *TaskCategories) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"sort_order",
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

	varTaskCategories := _TaskCategories{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTaskCategories)

	if err != nil {
		return err
	}

	*o = TaskCategories(varTaskCategories)

	return err
}

type NullableTaskCategories struct {
	value *TaskCategories
	isSet bool
}

func (v NullableTaskCategories) Get() *TaskCategories {
	return v.value
}

func (v *NullableTaskCategories) Set(val *TaskCategories) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskCategories) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskCategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskCategories(val *TaskCategories) *NullableTaskCategories {
	return &NullableTaskCategories{value: val, isSet: true}
}

func (v NullableTaskCategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskCategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

