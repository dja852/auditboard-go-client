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

// checks if the AssessmentPeriodsPutPreviousValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssessmentPeriodsPutPreviousValues{}

// AssessmentPeriodsPutPreviousValues struct for AssessmentPeriodsPutPreviousValues
type AssessmentPeriodsPutPreviousValues struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Name *string `json:"name,omitempty"`
	SortOrder *int32 `json:"sort_order,omitempty"`
}

// NewAssessmentPeriodsPutPreviousValues instantiates a new AssessmentPeriodsPutPreviousValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssessmentPeriodsPutPreviousValues() *AssessmentPeriodsPutPreviousValues {
	this := AssessmentPeriodsPutPreviousValues{}
	var sortOrder int32 = 0
	this.SortOrder = &sortOrder
	return &this
}

// NewAssessmentPeriodsPutPreviousValuesWithDefaults instantiates a new AssessmentPeriodsPutPreviousValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssessmentPeriodsPutPreviousValuesWithDefaults() *AssessmentPeriodsPutPreviousValues {
	this := AssessmentPeriodsPutPreviousValues{}
	var sortOrder int32 = 0
	this.SortOrder = &sortOrder
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssessmentPeriodsPutPreviousValues) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentPeriodsPutPreviousValues) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssessmentPeriodsPutPreviousValues) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AssessmentPeriodsPutPreviousValues) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AssessmentPeriodsPutPreviousValues) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentPeriodsPutPreviousValues) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AssessmentPeriodsPutPreviousValues) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AssessmentPeriodsPutPreviousValues) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AssessmentPeriodsPutPreviousValues) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentPeriodsPutPreviousValues) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AssessmentPeriodsPutPreviousValues) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AssessmentPeriodsPutPreviousValues) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *AssessmentPeriodsPutPreviousValues) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentPeriodsPutPreviousValues) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *AssessmentPeriodsPutPreviousValues) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *AssessmentPeriodsPutPreviousValues) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AssessmentPeriodsPutPreviousValues) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentPeriodsPutPreviousValues) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AssessmentPeriodsPutPreviousValues) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AssessmentPeriodsPutPreviousValues) SetName(v string) {
	o.Name = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *AssessmentPeriodsPutPreviousValues) GetSortOrder() int32 {
	if o == nil || IsNil(o.SortOrder) {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentPeriodsPutPreviousValues) GetSortOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *AssessmentPeriodsPutPreviousValues) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *AssessmentPeriodsPutPreviousValues) SetSortOrder(v int32) {
	o.SortOrder = &v
}

func (o AssessmentPeriodsPutPreviousValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssessmentPeriodsPutPreviousValues) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SortOrder) {
		toSerialize["sort_order"] = o.SortOrder
	}
	return toSerialize, nil
}

type NullableAssessmentPeriodsPutPreviousValues struct {
	value *AssessmentPeriodsPutPreviousValues
	isSet bool
}

func (v NullableAssessmentPeriodsPutPreviousValues) Get() *AssessmentPeriodsPutPreviousValues {
	return v.value
}

func (v *NullableAssessmentPeriodsPutPreviousValues) Set(val *AssessmentPeriodsPutPreviousValues) {
	v.value = val
	v.isSet = true
}

func (v NullableAssessmentPeriodsPutPreviousValues) IsSet() bool {
	return v.isSet
}

func (v *NullableAssessmentPeriodsPutPreviousValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssessmentPeriodsPutPreviousValues(val *AssessmentPeriodsPutPreviousValues) *NullableAssessmentPeriodsPutPreviousValues {
	return &NullableAssessmentPeriodsPutPreviousValues{value: val, isSet: true}
}

func (v NullableAssessmentPeriodsPutPreviousValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssessmentPeriodsPutPreviousValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


