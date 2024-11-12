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

// checks if the AuditOfficeLocationsPutPreviousValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditOfficeLocationsPutPreviousValues{}

// AuditOfficeLocationsPutPreviousValues struct for AuditOfficeLocationsPutPreviousValues
type AuditOfficeLocationsPutPreviousValues struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	SortOrder *int32 `json:"sort_order,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewAuditOfficeLocationsPutPreviousValues instantiates a new AuditOfficeLocationsPutPreviousValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditOfficeLocationsPutPreviousValues() *AuditOfficeLocationsPutPreviousValues {
	this := AuditOfficeLocationsPutPreviousValues{}
	return &this
}

// NewAuditOfficeLocationsPutPreviousValuesWithDefaults instantiates a new AuditOfficeLocationsPutPreviousValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditOfficeLocationsPutPreviousValuesWithDefaults() *AuditOfficeLocationsPutPreviousValues {
	this := AuditOfficeLocationsPutPreviousValues{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuditOfficeLocationsPutPreviousValues) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditOfficeLocationsPutPreviousValues) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuditOfficeLocationsPutPreviousValues) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AuditOfficeLocationsPutPreviousValues) SetId(v int32) {
	o.Id = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *AuditOfficeLocationsPutPreviousValues) GetSortOrder() int32 {
	if o == nil || IsNil(o.SortOrder) {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditOfficeLocationsPutPreviousValues) GetSortOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *AuditOfficeLocationsPutPreviousValues) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *AuditOfficeLocationsPutPreviousValues) SetSortOrder(v int32) {
	o.SortOrder = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AuditOfficeLocationsPutPreviousValues) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditOfficeLocationsPutPreviousValues) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AuditOfficeLocationsPutPreviousValues) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AuditOfficeLocationsPutPreviousValues) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AuditOfficeLocationsPutPreviousValues) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditOfficeLocationsPutPreviousValues) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AuditOfficeLocationsPutPreviousValues) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AuditOfficeLocationsPutPreviousValues) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *AuditOfficeLocationsPutPreviousValues) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditOfficeLocationsPutPreviousValues) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *AuditOfficeLocationsPutPreviousValues) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *AuditOfficeLocationsPutPreviousValues) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuditOfficeLocationsPutPreviousValues) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditOfficeLocationsPutPreviousValues) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuditOfficeLocationsPutPreviousValues) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuditOfficeLocationsPutPreviousValues) SetName(v string) {
	o.Name = &v
}

func (o AuditOfficeLocationsPutPreviousValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditOfficeLocationsPutPreviousValues) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableAuditOfficeLocationsPutPreviousValues struct {
	value *AuditOfficeLocationsPutPreviousValues
	isSet bool
}

func (v NullableAuditOfficeLocationsPutPreviousValues) Get() *AuditOfficeLocationsPutPreviousValues {
	return v.value
}

func (v *NullableAuditOfficeLocationsPutPreviousValues) Set(val *AuditOfficeLocationsPutPreviousValues) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditOfficeLocationsPutPreviousValues) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditOfficeLocationsPutPreviousValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditOfficeLocationsPutPreviousValues(val *AuditOfficeLocationsPutPreviousValues) *NullableAuditOfficeLocationsPutPreviousValues {
	return &NullableAuditOfficeLocationsPutPreviousValues{value: val, isSet: true}
}

func (v NullableAuditOfficeLocationsPutPreviousValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditOfficeLocationsPutPreviousValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


