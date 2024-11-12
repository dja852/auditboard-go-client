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

// checks if the TaskCategoriesGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskCategoriesGet200Response{}

// TaskCategoriesGet200Response struct for TaskCategoriesGet200Response
type TaskCategoriesGet200Response struct {
	TaskCategories []TaskCategories `json:"task_categories,omitempty"`
}

// NewTaskCategoriesGet200Response instantiates a new TaskCategoriesGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskCategoriesGet200Response() *TaskCategoriesGet200Response {
	this := TaskCategoriesGet200Response{}
	return &this
}

// NewTaskCategoriesGet200ResponseWithDefaults instantiates a new TaskCategoriesGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskCategoriesGet200ResponseWithDefaults() *TaskCategoriesGet200Response {
	this := TaskCategoriesGet200Response{}
	return &this
}

// GetTaskCategories returns the TaskCategories field value if set, zero value otherwise.
func (o *TaskCategoriesGet200Response) GetTaskCategories() []TaskCategories {
	if o == nil || IsNil(o.TaskCategories) {
		var ret []TaskCategories
		return ret
	}
	return o.TaskCategories
}

// GetTaskCategoriesOk returns a tuple with the TaskCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCategoriesGet200Response) GetTaskCategoriesOk() ([]TaskCategories, bool) {
	if o == nil || IsNil(o.TaskCategories) {
		return nil, false
	}
	return o.TaskCategories, true
}

// HasTaskCategories returns a boolean if a field has been set.
func (o *TaskCategoriesGet200Response) HasTaskCategories() bool {
	if o != nil && !IsNil(o.TaskCategories) {
		return true
	}

	return false
}

// SetTaskCategories gets a reference to the given []TaskCategories and assigns it to the TaskCategories field.
func (o *TaskCategoriesGet200Response) SetTaskCategories(v []TaskCategories) {
	o.TaskCategories = v
}

func (o TaskCategoriesGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskCategoriesGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TaskCategories) {
		toSerialize["task_categories"] = o.TaskCategories
	}
	return toSerialize, nil
}

type NullableTaskCategoriesGet200Response struct {
	value *TaskCategoriesGet200Response
	isSet bool
}

func (v NullableTaskCategoriesGet200Response) Get() *TaskCategoriesGet200Response {
	return v.value
}

func (v *NullableTaskCategoriesGet200Response) Set(val *TaskCategoriesGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskCategoriesGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskCategoriesGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskCategoriesGet200Response(val *TaskCategoriesGet200Response) *NullableTaskCategoriesGet200Response {
	return &NullableTaskCategoriesGet200Response{value: val, isSet: true}
}

func (v NullableTaskCategoriesGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskCategoriesGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

