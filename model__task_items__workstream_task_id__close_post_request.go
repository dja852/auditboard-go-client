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

// checks if the TaskItemsWorkstreamTaskIdClosePostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskItemsWorkstreamTaskIdClosePostRequest{}

// TaskItemsWorkstreamTaskIdClosePostRequest struct for TaskItemsWorkstreamTaskIdClosePostRequest
type TaskItemsWorkstreamTaskIdClosePostRequest struct {
	SkipAdditionalAction *TaskItemsWorkstreamTaskIdClosePostRequestSkipAdditionalAction `json:"skipAdditionalAction,omitempty"`
}

// NewTaskItemsWorkstreamTaskIdClosePostRequest instantiates a new TaskItemsWorkstreamTaskIdClosePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskItemsWorkstreamTaskIdClosePostRequest() *TaskItemsWorkstreamTaskIdClosePostRequest {
	this := TaskItemsWorkstreamTaskIdClosePostRequest{}
	return &this
}

// NewTaskItemsWorkstreamTaskIdClosePostRequestWithDefaults instantiates a new TaskItemsWorkstreamTaskIdClosePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskItemsWorkstreamTaskIdClosePostRequestWithDefaults() *TaskItemsWorkstreamTaskIdClosePostRequest {
	this := TaskItemsWorkstreamTaskIdClosePostRequest{}
	return &this
}

// GetSkipAdditionalAction returns the SkipAdditionalAction field value if set, zero value otherwise.
func (o *TaskItemsWorkstreamTaskIdClosePostRequest) GetSkipAdditionalAction() TaskItemsWorkstreamTaskIdClosePostRequestSkipAdditionalAction {
	if o == nil || IsNil(o.SkipAdditionalAction) {
		var ret TaskItemsWorkstreamTaskIdClosePostRequestSkipAdditionalAction
		return ret
	}
	return *o.SkipAdditionalAction
}

// GetSkipAdditionalActionOk returns a tuple with the SkipAdditionalAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskItemsWorkstreamTaskIdClosePostRequest) GetSkipAdditionalActionOk() (*TaskItemsWorkstreamTaskIdClosePostRequestSkipAdditionalAction, bool) {
	if o == nil || IsNil(o.SkipAdditionalAction) {
		return nil, false
	}
	return o.SkipAdditionalAction, true
}

// HasSkipAdditionalAction returns a boolean if a field has been set.
func (o *TaskItemsWorkstreamTaskIdClosePostRequest) HasSkipAdditionalAction() bool {
	if o != nil && !IsNil(o.SkipAdditionalAction) {
		return true
	}

	return false
}

// SetSkipAdditionalAction gets a reference to the given TaskItemsWorkstreamTaskIdClosePostRequestSkipAdditionalAction and assigns it to the SkipAdditionalAction field.
func (o *TaskItemsWorkstreamTaskIdClosePostRequest) SetSkipAdditionalAction(v TaskItemsWorkstreamTaskIdClosePostRequestSkipAdditionalAction) {
	o.SkipAdditionalAction = &v
}

func (o TaskItemsWorkstreamTaskIdClosePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskItemsWorkstreamTaskIdClosePostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SkipAdditionalAction) {
		toSerialize["skipAdditionalAction"] = o.SkipAdditionalAction
	}
	return toSerialize, nil
}

type NullableTaskItemsWorkstreamTaskIdClosePostRequest struct {
	value *TaskItemsWorkstreamTaskIdClosePostRequest
	isSet bool
}

func (v NullableTaskItemsWorkstreamTaskIdClosePostRequest) Get() *TaskItemsWorkstreamTaskIdClosePostRequest {
	return v.value
}

func (v *NullableTaskItemsWorkstreamTaskIdClosePostRequest) Set(val *TaskItemsWorkstreamTaskIdClosePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskItemsWorkstreamTaskIdClosePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskItemsWorkstreamTaskIdClosePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskItemsWorkstreamTaskIdClosePostRequest(val *TaskItemsWorkstreamTaskIdClosePostRequest) *NullableTaskItemsWorkstreamTaskIdClosePostRequest {
	return &NullableTaskItemsWorkstreamTaskIdClosePostRequest{value: val, isSet: true}
}

func (v NullableTaskItemsWorkstreamTaskIdClosePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskItemsWorkstreamTaskIdClosePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


