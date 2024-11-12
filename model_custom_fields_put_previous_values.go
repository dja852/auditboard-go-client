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

// checks if the CustomFieldsPutPreviousValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFieldsPutPreviousValues{}

// CustomFieldsPutPreviousValues struct for CustomFieldsPutPreviousValues
type CustomFieldsPutPreviousValues struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	Name *string `json:"name,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	AttrKey *string `json:"attr_key,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	LinkifyKey *string `json:"linkify_key,omitempty"`
}

// NewCustomFieldsPutPreviousValues instantiates a new CustomFieldsPutPreviousValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFieldsPutPreviousValues() *CustomFieldsPutPreviousValues {
	this := CustomFieldsPutPreviousValues{}
	return &this
}

// NewCustomFieldsPutPreviousValuesWithDefaults instantiates a new CustomFieldsPutPreviousValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFieldsPutPreviousValuesWithDefaults() *CustomFieldsPutPreviousValues {
	this := CustomFieldsPutPreviousValues{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomFieldsPutPreviousValues) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldsPutPreviousValues) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomFieldsPutPreviousValues) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CustomFieldsPutPreviousValues) SetId(v int32) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CustomFieldsPutPreviousValues) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldsPutPreviousValues) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomFieldsPutPreviousValues) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CustomFieldsPutPreviousValues) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomFieldsPutPreviousValues) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldsPutPreviousValues) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomFieldsPutPreviousValues) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomFieldsPutPreviousValues) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CustomFieldsPutPreviousValues) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldsPutPreviousValues) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CustomFieldsPutPreviousValues) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CustomFieldsPutPreviousValues) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetAttrKey returns the AttrKey field value if set, zero value otherwise.
func (o *CustomFieldsPutPreviousValues) GetAttrKey() string {
	if o == nil || IsNil(o.AttrKey) {
		var ret string
		return ret
	}
	return *o.AttrKey
}

// GetAttrKeyOk returns a tuple with the AttrKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldsPutPreviousValues) GetAttrKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AttrKey) {
		return nil, false
	}
	return o.AttrKey, true
}

// HasAttrKey returns a boolean if a field has been set.
func (o *CustomFieldsPutPreviousValues) HasAttrKey() bool {
	if o != nil && !IsNil(o.AttrKey) {
		return true
	}

	return false
}

// SetAttrKey gets a reference to the given string and assigns it to the AttrKey field.
func (o *CustomFieldsPutPreviousValues) SetAttrKey(v string) {
	o.AttrKey = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CustomFieldsPutPreviousValues) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldsPutPreviousValues) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomFieldsPutPreviousValues) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *CustomFieldsPutPreviousValues) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CustomFieldsPutPreviousValues) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldsPutPreviousValues) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CustomFieldsPutPreviousValues) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *CustomFieldsPutPreviousValues) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *CustomFieldsPutPreviousValues) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldsPutPreviousValues) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *CustomFieldsPutPreviousValues) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *CustomFieldsPutPreviousValues) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetLinkifyKey returns the LinkifyKey field value if set, zero value otherwise.
func (o *CustomFieldsPutPreviousValues) GetLinkifyKey() string {
	if o == nil || IsNil(o.LinkifyKey) {
		var ret string
		return ret
	}
	return *o.LinkifyKey
}

// GetLinkifyKeyOk returns a tuple with the LinkifyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldsPutPreviousValues) GetLinkifyKeyOk() (*string, bool) {
	if o == nil || IsNil(o.LinkifyKey) {
		return nil, false
	}
	return o.LinkifyKey, true
}

// HasLinkifyKey returns a boolean if a field has been set.
func (o *CustomFieldsPutPreviousValues) HasLinkifyKey() bool {
	if o != nil && !IsNil(o.LinkifyKey) {
		return true
	}

	return false
}

// SetLinkifyKey gets a reference to the given string and assigns it to the LinkifyKey field.
func (o *CustomFieldsPutPreviousValues) SetLinkifyKey(v string) {
	o.LinkifyKey = &v
}

func (o CustomFieldsPutPreviousValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFieldsPutPreviousValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.AttrKey) {
		toSerialize["attr_key"] = o.AttrKey
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
	if !IsNil(o.LinkifyKey) {
		toSerialize["linkify_key"] = o.LinkifyKey
	}
	return toSerialize, nil
}

type NullableCustomFieldsPutPreviousValues struct {
	value *CustomFieldsPutPreviousValues
	isSet bool
}

func (v NullableCustomFieldsPutPreviousValues) Get() *CustomFieldsPutPreviousValues {
	return v.value
}

func (v *NullableCustomFieldsPutPreviousValues) Set(val *CustomFieldsPutPreviousValues) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldsPutPreviousValues) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldsPutPreviousValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldsPutPreviousValues(val *CustomFieldsPutPreviousValues) *NullableCustomFieldsPutPreviousValues {
	return &NullableCustomFieldsPutPreviousValues{value: val, isSet: true}
}

func (v NullableCustomFieldsPutPreviousValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldsPutPreviousValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


