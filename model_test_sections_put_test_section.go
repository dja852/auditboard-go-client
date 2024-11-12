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

// checks if the TestSectionsPutTestSection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestSectionsPutTestSection{}

// TestSectionsPutTestSection struct for TestSectionsPutTestSection
type TestSectionsPutTestSection struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	SortOrder *int32 `json:"sort_order,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Name string `json:"name"`
	ShortName *string `json:"short_name,omitempty"`
	Uid *string `json:"uid,omitempty"`
}

type _TestSectionsPutTestSection TestSectionsPutTestSection

// NewTestSectionsPutTestSection instantiates a new TestSectionsPutTestSection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestSectionsPutTestSection(name string) *TestSectionsPutTestSection {
	this := TestSectionsPutTestSection{}
	this.Name = name
	return &this
}

// NewTestSectionsPutTestSectionWithDefaults instantiates a new TestSectionsPutTestSection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestSectionsPutTestSectionWithDefaults() *TestSectionsPutTestSection {
	this := TestSectionsPutTestSection{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TestSectionsPutTestSection) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSectionsPutTestSection) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TestSectionsPutTestSection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TestSectionsPutTestSection) SetId(v int32) {
	o.Id = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *TestSectionsPutTestSection) GetSortOrder() int32 {
	if o == nil || IsNil(o.SortOrder) {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSectionsPutTestSection) GetSortOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *TestSectionsPutTestSection) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *TestSectionsPutTestSection) SetSortOrder(v int32) {
	o.SortOrder = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TestSectionsPutTestSection) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSectionsPutTestSection) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TestSectionsPutTestSection) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *TestSectionsPutTestSection) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TestSectionsPutTestSection) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSectionsPutTestSection) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TestSectionsPutTestSection) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *TestSectionsPutTestSection) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *TestSectionsPutTestSection) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSectionsPutTestSection) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *TestSectionsPutTestSection) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *TestSectionsPutTestSection) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetName returns the Name field value
func (o *TestSectionsPutTestSection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TestSectionsPutTestSection) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TestSectionsPutTestSection) SetName(v string) {
	o.Name = v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *TestSectionsPutTestSection) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSectionsPutTestSection) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *TestSectionsPutTestSection) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *TestSectionsPutTestSection) SetShortName(v string) {
	o.ShortName = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *TestSectionsPutTestSection) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSectionsPutTestSection) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *TestSectionsPutTestSection) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *TestSectionsPutTestSection) SetUid(v string) {
	o.Uid = &v
}

func (o TestSectionsPutTestSection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestSectionsPutTestSection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SortOrder) {
		toSerialize["sort_order"] = o.SortOrder
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
	if !IsNil(o.ShortName) {
		toSerialize["short_name"] = o.ShortName
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	return toSerialize, nil
}

type NullableTestSectionsPutTestSection struct {
	value *TestSectionsPutTestSection
	isSet bool
}

func (v NullableTestSectionsPutTestSection) Get() *TestSectionsPutTestSection {
	return v.value
}

func (v *NullableTestSectionsPutTestSection) Set(val *TestSectionsPutTestSection) {
	v.value = val
	v.isSet = true
}

func (v NullableTestSectionsPutTestSection) IsSet() bool {
	return v.isSet
}

func (v *NullableTestSectionsPutTestSection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestSectionsPutTestSection(val *TestSectionsPutTestSection) *NullableTestSectionsPutTestSection {
	return &NullableTestSectionsPutTestSection{value: val, isSet: true}
}

func (v NullableTestSectionsPutTestSection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestSectionsPutTestSection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


