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

// checks if the RcwSchedules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RcwSchedules{}

// RcwSchedules struct for RcwSchedules
type RcwSchedules struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	Name string `json:"name"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	CreatedByUserId *int32 `json:"created_by_user_id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}

type _RcwSchedules RcwSchedules

// NewRcwSchedules instantiates a new RcwSchedules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRcwSchedules(name string) *RcwSchedules {
	this := RcwSchedules{}
	this.Name = name
	return &this
}

// NewRcwSchedulesWithDefaults instantiates a new RcwSchedules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRcwSchedulesWithDefaults() *RcwSchedules {
	this := RcwSchedules{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RcwSchedules) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwSchedules) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RcwSchedules) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RcwSchedules) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *RcwSchedules) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RcwSchedules) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RcwSchedules) SetName(v string) {
	o.Name = v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *RcwSchedules) GetCreatedByUserId() int32 {
	if o == nil || IsNil(o.CreatedByUserId) {
		var ret int32
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwSchedules) GetCreatedByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedByUserId) {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *RcwSchedules) HasCreatedByUserId() bool {
	if o != nil && !IsNil(o.CreatedByUserId) {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given int32 and assigns it to the CreatedByUserId field.
func (o *RcwSchedules) SetCreatedByUserId(v int32) {
	o.CreatedByUserId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RcwSchedules) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwSchedules) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RcwSchedules) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *RcwSchedules) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RcwSchedules) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwSchedules) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RcwSchedules) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *RcwSchedules) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *RcwSchedules) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwSchedules) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *RcwSchedules) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *RcwSchedules) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

func (o RcwSchedules) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RcwSchedules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.CreatedByUserId) {
		toSerialize["created_by_user_id"] = o.CreatedByUserId
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
	return toSerialize, nil
}

type NullableRcwSchedules struct {
	value *RcwSchedules
	isSet bool
}

func (v NullableRcwSchedules) Get() *RcwSchedules {
	return v.value
}

func (v *NullableRcwSchedules) Set(val *RcwSchedules) {
	v.value = val
	v.isSet = true
}

func (v NullableRcwSchedules) IsSet() bool {
	return v.isSet
}

func (v *NullableRcwSchedules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRcwSchedules(val *RcwSchedules) *NullableRcwSchedules {
	return &NullableRcwSchedules{value: val, isSet: true}
}

func (v NullableRcwSchedules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRcwSchedules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


