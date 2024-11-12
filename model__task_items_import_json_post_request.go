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

// checks if the TaskItemsImportJsonPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskItemsImportJsonPostRequest{}

// TaskItemsImportJsonPostRequest struct for TaskItemsImportJsonPostRequest
type TaskItemsImportJsonPostRequest struct {
	Json []TaskItemsImportJsonPostRequestJsonInner `json:"json,omitempty"`
	Options *TaskItemsImportJsonPostRequestOptions `json:"options,omitempty"`
}

// NewTaskItemsImportJsonPostRequest instantiates a new TaskItemsImportJsonPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskItemsImportJsonPostRequest() *TaskItemsImportJsonPostRequest {
	this := TaskItemsImportJsonPostRequest{}
	return &this
}

// NewTaskItemsImportJsonPostRequestWithDefaults instantiates a new TaskItemsImportJsonPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskItemsImportJsonPostRequestWithDefaults() *TaskItemsImportJsonPostRequest {
	this := TaskItemsImportJsonPostRequest{}
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *TaskItemsImportJsonPostRequest) GetJson() []TaskItemsImportJsonPostRequestJsonInner {
	if o == nil || IsNil(o.Json) {
		var ret []TaskItemsImportJsonPostRequestJsonInner
		return ret
	}
	return o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskItemsImportJsonPostRequest) GetJsonOk() ([]TaskItemsImportJsonPostRequestJsonInner, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *TaskItemsImportJsonPostRequest) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given []TaskItemsImportJsonPostRequestJsonInner and assigns it to the Json field.
func (o *TaskItemsImportJsonPostRequest) SetJson(v []TaskItemsImportJsonPostRequestJsonInner) {
	o.Json = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *TaskItemsImportJsonPostRequest) GetOptions() TaskItemsImportJsonPostRequestOptions {
	if o == nil || IsNil(o.Options) {
		var ret TaskItemsImportJsonPostRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskItemsImportJsonPostRequest) GetOptionsOk() (*TaskItemsImportJsonPostRequestOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *TaskItemsImportJsonPostRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given TaskItemsImportJsonPostRequestOptions and assigns it to the Options field.
func (o *TaskItemsImportJsonPostRequest) SetOptions(v TaskItemsImportJsonPostRequestOptions) {
	o.Options = &v
}

func (o TaskItemsImportJsonPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskItemsImportJsonPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableTaskItemsImportJsonPostRequest struct {
	value *TaskItemsImportJsonPostRequest
	isSet bool
}

func (v NullableTaskItemsImportJsonPostRequest) Get() *TaskItemsImportJsonPostRequest {
	return v.value
}

func (v *NullableTaskItemsImportJsonPostRequest) Set(val *TaskItemsImportJsonPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskItemsImportJsonPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskItemsImportJsonPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskItemsImportJsonPostRequest(val *TaskItemsImportJsonPostRequest) *NullableTaskItemsImportJsonPostRequest {
	return &NullableTaskItemsImportJsonPostRequest{value: val, isSet: true}
}

func (v NullableTaskItemsImportJsonPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskItemsImportJsonPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

