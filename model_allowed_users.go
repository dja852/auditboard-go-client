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

// checks if the AllowedUsers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllowedUsers{}

// AllowedUsers struct for AllowedUsers
type AllowedUsers struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	UserableId *int32 `json:"userable_id,omitempty"`
	UserableType *string `json:"userable_type,omitempty"`
	Permission *string `json:"permission,omitempty"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	UserId int32 `json:"user_id"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

type _AllowedUsers AllowedUsers

// NewAllowedUsers instantiates a new AllowedUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowedUsers(userId int32) *AllowedUsers {
	this := AllowedUsers{}
	this.UserId = userId
	return &this
}

// NewAllowedUsersWithDefaults instantiates a new AllowedUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowedUsersWithDefaults() *AllowedUsers {
	this := AllowedUsers{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AllowedUsers) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedUsers) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AllowedUsers) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AllowedUsers) SetId(v int32) {
	o.Id = &v
}

// GetUserableId returns the UserableId field value if set, zero value otherwise.
func (o *AllowedUsers) GetUserableId() int32 {
	if o == nil || IsNil(o.UserableId) {
		var ret int32
		return ret
	}
	return *o.UserableId
}

// GetUserableIdOk returns a tuple with the UserableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedUsers) GetUserableIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserableId) {
		return nil, false
	}
	return o.UserableId, true
}

// HasUserableId returns a boolean if a field has been set.
func (o *AllowedUsers) HasUserableId() bool {
	if o != nil && !IsNil(o.UserableId) {
		return true
	}

	return false
}

// SetUserableId gets a reference to the given int32 and assigns it to the UserableId field.
func (o *AllowedUsers) SetUserableId(v int32) {
	o.UserableId = &v
}

// GetUserableType returns the UserableType field value if set, zero value otherwise.
func (o *AllowedUsers) GetUserableType() string {
	if o == nil || IsNil(o.UserableType) {
		var ret string
		return ret
	}
	return *o.UserableType
}

// GetUserableTypeOk returns a tuple with the UserableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedUsers) GetUserableTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UserableType) {
		return nil, false
	}
	return o.UserableType, true
}

// HasUserableType returns a boolean if a field has been set.
func (o *AllowedUsers) HasUserableType() bool {
	if o != nil && !IsNil(o.UserableType) {
		return true
	}

	return false
}

// SetUserableType gets a reference to the given string and assigns it to the UserableType field.
func (o *AllowedUsers) SetUserableType(v string) {
	o.UserableType = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *AllowedUsers) GetPermission() string {
	if o == nil || IsNil(o.Permission) {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedUsers) GetPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *AllowedUsers) HasPermission() bool {
	if o != nil && !IsNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *AllowedUsers) SetPermission(v string) {
	o.Permission = &v
}

// GetUserId returns the UserId field value
func (o *AllowedUsers) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *AllowedUsers) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *AllowedUsers) SetUserId(v int32) {
	o.UserId = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AllowedUsers) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedUsers) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AllowedUsers) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AllowedUsers) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AllowedUsers) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedUsers) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AllowedUsers) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AllowedUsers) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o AllowedUsers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllowedUsers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserableId) {
		toSerialize["userable_id"] = o.UserableId
	}
	if !IsNil(o.UserableType) {
		toSerialize["userable_type"] = o.UserableType
	}
	if !IsNil(o.Permission) {
		toSerialize["permission"] = o.Permission
	}
	toSerialize["user_id"] = o.UserId
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

func (o *AllowedUsers) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_id",
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

	varAllowedUsers := _AllowedUsers{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAllowedUsers)

	if err != nil {
		return err
	}

	*o = AllowedUsers(varAllowedUsers)

	return err
}

type NullableAllowedUsers struct {
	value *AllowedUsers
	isSet bool
}

func (v NullableAllowedUsers) Get() *AllowedUsers {
	return v.value
}

func (v *NullableAllowedUsers) Set(val *AllowedUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowedUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowedUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowedUsers(val *AllowedUsers) *NullableAllowedUsers {
	return &NullableAllowedUsers{value: val, isSet: true}
}

func (v NullableAllowedUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowedUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


