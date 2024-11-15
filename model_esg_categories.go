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

// checks if the EsgCategories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsgCategories{}

// EsgCategories struct for EsgCategories
type EsgCategories struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Name string `json:"name"`
	SortOrder int32 `json:"sort_order"`
	Scopes interface{} `json:"scopes,omitempty"`
}

type _EsgCategories EsgCategories

// NewEsgCategories instantiates a new EsgCategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsgCategories(name string, sortOrder int32) *EsgCategories {
	this := EsgCategories{}
	this.Name = name
	this.SortOrder = sortOrder
	return &this
}

// NewEsgCategoriesWithDefaults instantiates a new EsgCategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsgCategoriesWithDefaults() *EsgCategories {
	this := EsgCategories{}
	var sortOrder int32 = 0
	this.SortOrder = sortOrder
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EsgCategories) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgCategories) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EsgCategories) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EsgCategories) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EsgCategories) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgCategories) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EsgCategories) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *EsgCategories) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EsgCategories) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgCategories) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EsgCategories) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *EsgCategories) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *EsgCategories) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgCategories) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *EsgCategories) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *EsgCategories) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetName returns the Name field value
func (o *EsgCategories) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EsgCategories) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EsgCategories) SetName(v string) {
	o.Name = v
}

// GetSortOrder returns the SortOrder field value
func (o *EsgCategories) GetSortOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value
// and a boolean to check if the value has been set.
func (o *EsgCategories) GetSortOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortOrder, true
}

// SetSortOrder sets field value
func (o *EsgCategories) SetSortOrder(v int32) {
	o.SortOrder = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EsgCategories) GetScopes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EsgCategories) GetScopesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return &o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *EsgCategories) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given interface{} and assigns it to the Scopes field.
func (o *EsgCategories) SetScopes(v interface{}) {
	o.Scopes = v
}

func (o EsgCategories) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsgCategories) ToMap() (map[string]interface{}, error) {
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
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

type NullableEsgCategories struct {
	value *EsgCategories
	isSet bool
}

func (v NullableEsgCategories) Get() *EsgCategories {
	return v.value
}

func (v *NullableEsgCategories) Set(val *EsgCategories) {
	v.value = val
	v.isSet = true
}

func (v NullableEsgCategories) IsSet() bool {
	return v.isSet
}

func (v *NullableEsgCategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsgCategories(val *EsgCategories) *NullableEsgCategories {
	return &NullableEsgCategories{value: val, isSet: true}
}

func (v NullableEsgCategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsgCategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


