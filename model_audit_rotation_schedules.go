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

// checks if the AuditRotationSchedules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditRotationSchedules{}

// AuditRotationSchedules struct for AuditRotationSchedules
type AuditRotationSchedules struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	SortOrder *int32 `json:"sort_order,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Name string `json:"name"`
}

type _AuditRotationSchedules AuditRotationSchedules

// NewAuditRotationSchedules instantiates a new AuditRotationSchedules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditRotationSchedules(name string) *AuditRotationSchedules {
	this := AuditRotationSchedules{}
	this.Name = name
	return &this
}

// NewAuditRotationSchedulesWithDefaults instantiates a new AuditRotationSchedules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditRotationSchedulesWithDefaults() *AuditRotationSchedules {
	this := AuditRotationSchedules{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuditRotationSchedules) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditRotationSchedules) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuditRotationSchedules) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AuditRotationSchedules) SetId(v int32) {
	o.Id = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *AuditRotationSchedules) GetSortOrder() int32 {
	if o == nil || IsNil(o.SortOrder) {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditRotationSchedules) GetSortOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *AuditRotationSchedules) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *AuditRotationSchedules) SetSortOrder(v int32) {
	o.SortOrder = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AuditRotationSchedules) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditRotationSchedules) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AuditRotationSchedules) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AuditRotationSchedules) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AuditRotationSchedules) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditRotationSchedules) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AuditRotationSchedules) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AuditRotationSchedules) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *AuditRotationSchedules) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditRotationSchedules) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *AuditRotationSchedules) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *AuditRotationSchedules) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetName returns the Name field value
func (o *AuditRotationSchedules) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuditRotationSchedules) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuditRotationSchedules) SetName(v string) {
	o.Name = v
}

func (o AuditRotationSchedules) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditRotationSchedules) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

func (o *AuditRotationSchedules) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varAuditRotationSchedules := _AuditRotationSchedules{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuditRotationSchedules)

	if err != nil {
		return err
	}

	*o = AuditRotationSchedules(varAuditRotationSchedules)

	return err
}

type NullableAuditRotationSchedules struct {
	value *AuditRotationSchedules
	isSet bool
}

func (v NullableAuditRotationSchedules) Get() *AuditRotationSchedules {
	return v.value
}

func (v *NullableAuditRotationSchedules) Set(val *AuditRotationSchedules) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditRotationSchedules) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditRotationSchedules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditRotationSchedules(val *AuditRotationSchedules) *NullableAuditRotationSchedules {
	return &NullableAuditRotationSchedules{value: val, isSet: true}
}

func (v NullableAuditRotationSchedules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditRotationSchedules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

