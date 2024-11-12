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

// checks if the TaskPeriodsPutTaskPeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskPeriodsPutTaskPeriod{}

// TaskPeriodsPutTaskPeriod struct for TaskPeriodsPutTaskPeriod
type TaskPeriodsPutTaskPeriod struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Name string `json:"name"`
	SortOrder int32 `json:"sort_order"`
	ShortName *string `json:"short_name,omitempty"`
	StartDate *string `json:"start_date,omitempty"`
	EndDate *string `json:"end_date,omitempty"`
}

type _TaskPeriodsPutTaskPeriod TaskPeriodsPutTaskPeriod

// NewTaskPeriodsPutTaskPeriod instantiates a new TaskPeriodsPutTaskPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskPeriodsPutTaskPeriod(name string, sortOrder int32) *TaskPeriodsPutTaskPeriod {
	this := TaskPeriodsPutTaskPeriod{}
	this.Name = name
	this.SortOrder = sortOrder
	return &this
}

// NewTaskPeriodsPutTaskPeriodWithDefaults instantiates a new TaskPeriodsPutTaskPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskPeriodsPutTaskPeriodWithDefaults() *TaskPeriodsPutTaskPeriod {
	this := TaskPeriodsPutTaskPeriod{}
	var sortOrder int32 = 0
	this.SortOrder = sortOrder
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaskPeriodsPutTaskPeriod) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPeriodsPutTaskPeriod) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TaskPeriodsPutTaskPeriod) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TaskPeriodsPutTaskPeriod) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TaskPeriodsPutTaskPeriod) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPeriodsPutTaskPeriod) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TaskPeriodsPutTaskPeriod) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *TaskPeriodsPutTaskPeriod) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TaskPeriodsPutTaskPeriod) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPeriodsPutTaskPeriod) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TaskPeriodsPutTaskPeriod) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *TaskPeriodsPutTaskPeriod) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *TaskPeriodsPutTaskPeriod) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPeriodsPutTaskPeriod) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *TaskPeriodsPutTaskPeriod) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *TaskPeriodsPutTaskPeriod) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetName returns the Name field value
func (o *TaskPeriodsPutTaskPeriod) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TaskPeriodsPutTaskPeriod) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TaskPeriodsPutTaskPeriod) SetName(v string) {
	o.Name = v
}

// GetSortOrder returns the SortOrder field value
func (o *TaskPeriodsPutTaskPeriod) GetSortOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value
// and a boolean to check if the value has been set.
func (o *TaskPeriodsPutTaskPeriod) GetSortOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortOrder, true
}

// SetSortOrder sets field value
func (o *TaskPeriodsPutTaskPeriod) SetSortOrder(v int32) {
	o.SortOrder = v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *TaskPeriodsPutTaskPeriod) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPeriodsPutTaskPeriod) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *TaskPeriodsPutTaskPeriod) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *TaskPeriodsPutTaskPeriod) SetShortName(v string) {
	o.ShortName = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *TaskPeriodsPutTaskPeriod) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPeriodsPutTaskPeriod) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *TaskPeriodsPutTaskPeriod) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *TaskPeriodsPutTaskPeriod) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *TaskPeriodsPutTaskPeriod) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPeriodsPutTaskPeriod) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *TaskPeriodsPutTaskPeriod) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *TaskPeriodsPutTaskPeriod) SetEndDate(v string) {
	o.EndDate = &v
}

func (o TaskPeriodsPutTaskPeriod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskPeriodsPutTaskPeriod) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ShortName) {
		toSerialize["short_name"] = o.ShortName
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	return toSerialize, nil
}

func (o *TaskPeriodsPutTaskPeriod) UnmarshalJSON(data []byte) (err error) {
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

	varTaskPeriodsPutTaskPeriod := _TaskPeriodsPutTaskPeriod{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTaskPeriodsPutTaskPeriod)

	if err != nil {
		return err
	}

	*o = TaskPeriodsPutTaskPeriod(varTaskPeriodsPutTaskPeriod)

	return err
}

type NullableTaskPeriodsPutTaskPeriod struct {
	value *TaskPeriodsPutTaskPeriod
	isSet bool
}

func (v NullableTaskPeriodsPutTaskPeriod) Get() *TaskPeriodsPutTaskPeriod {
	return v.value
}

func (v *NullableTaskPeriodsPutTaskPeriod) Set(val *TaskPeriodsPutTaskPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskPeriodsPutTaskPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskPeriodsPutTaskPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskPeriodsPutTaskPeriod(val *TaskPeriodsPutTaskPeriod) *NullableTaskPeriodsPutTaskPeriod {
	return &NullableTaskPeriodsPutTaskPeriod{value: val, isSet: true}
}

func (v NullableTaskPeriodsPutTaskPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskPeriodsPutTaskPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


