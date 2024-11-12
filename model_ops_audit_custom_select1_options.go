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

// checks if the OpsAuditCustomSelect1Options type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpsAuditCustomSelect1Options{}

// OpsAuditCustomSelect1Options struct for OpsAuditCustomSelect1Options
type OpsAuditCustomSelect1Options struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Name string `json:"name"`
	SortOrder int32 `json:"sort_order"`
}

type _OpsAuditCustomSelect1Options OpsAuditCustomSelect1Options

// NewOpsAuditCustomSelect1Options instantiates a new OpsAuditCustomSelect1Options object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsAuditCustomSelect1Options(name string, sortOrder int32) *OpsAuditCustomSelect1Options {
	this := OpsAuditCustomSelect1Options{}
	this.Name = name
	this.SortOrder = sortOrder
	return &this
}

// NewOpsAuditCustomSelect1OptionsWithDefaults instantiates a new OpsAuditCustomSelect1Options object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsAuditCustomSelect1OptionsWithDefaults() *OpsAuditCustomSelect1Options {
	this := OpsAuditCustomSelect1Options{}
	var sortOrder int32 = 0
	this.SortOrder = sortOrder
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OpsAuditCustomSelect1Options) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditCustomSelect1Options) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OpsAuditCustomSelect1Options) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OpsAuditCustomSelect1Options) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OpsAuditCustomSelect1Options) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditCustomSelect1Options) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OpsAuditCustomSelect1Options) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *OpsAuditCustomSelect1Options) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OpsAuditCustomSelect1Options) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditCustomSelect1Options) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OpsAuditCustomSelect1Options) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *OpsAuditCustomSelect1Options) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *OpsAuditCustomSelect1Options) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditCustomSelect1Options) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *OpsAuditCustomSelect1Options) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *OpsAuditCustomSelect1Options) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetName returns the Name field value
func (o *OpsAuditCustomSelect1Options) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OpsAuditCustomSelect1Options) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OpsAuditCustomSelect1Options) SetName(v string) {
	o.Name = v
}

// GetSortOrder returns the SortOrder field value
func (o *OpsAuditCustomSelect1Options) GetSortOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value
// and a boolean to check if the value has been set.
func (o *OpsAuditCustomSelect1Options) GetSortOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortOrder, true
}

// SetSortOrder sets field value
func (o *OpsAuditCustomSelect1Options) SetSortOrder(v int32) {
	o.SortOrder = v
}

func (o OpsAuditCustomSelect1Options) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpsAuditCustomSelect1Options) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableOpsAuditCustomSelect1Options struct {
	value *OpsAuditCustomSelect1Options
	isSet bool
}

func (v NullableOpsAuditCustomSelect1Options) Get() *OpsAuditCustomSelect1Options {
	return v.value
}

func (v *NullableOpsAuditCustomSelect1Options) Set(val *OpsAuditCustomSelect1Options) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsAuditCustomSelect1Options) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsAuditCustomSelect1Options) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsAuditCustomSelect1Options(val *OpsAuditCustomSelect1Options) *NullableOpsAuditCustomSelect1Options {
	return &NullableOpsAuditCustomSelect1Options{value: val, isSet: true}
}

func (v NullableOpsAuditCustomSelect1Options) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsAuditCustomSelect1Options) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


