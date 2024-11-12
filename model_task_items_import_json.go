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

// checks if the TaskItemsImportJson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskItemsImportJson{}

// TaskItemsImportJson struct for TaskItemsImportJson
type TaskItemsImportJson struct {
	StorePayload *TaskItemsImportJsonStorePayload `json:"storePayload,omitempty"`
	Results []TaskItemsImportJsonResultsInner `json:"results,omitempty"`
	Summary *TaskItemsImportJsonSummary `json:"summary,omitempty"`
}

// NewTaskItemsImportJson instantiates a new TaskItemsImportJson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskItemsImportJson() *TaskItemsImportJson {
	this := TaskItemsImportJson{}
	return &this
}

// NewTaskItemsImportJsonWithDefaults instantiates a new TaskItemsImportJson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskItemsImportJsonWithDefaults() *TaskItemsImportJson {
	this := TaskItemsImportJson{}
	return &this
}

// GetStorePayload returns the StorePayload field value if set, zero value otherwise.
func (o *TaskItemsImportJson) GetStorePayload() TaskItemsImportJsonStorePayload {
	if o == nil || IsNil(o.StorePayload) {
		var ret TaskItemsImportJsonStorePayload
		return ret
	}
	return *o.StorePayload
}

// GetStorePayloadOk returns a tuple with the StorePayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskItemsImportJson) GetStorePayloadOk() (*TaskItemsImportJsonStorePayload, bool) {
	if o == nil || IsNil(o.StorePayload) {
		return nil, false
	}
	return o.StorePayload, true
}

// HasStorePayload returns a boolean if a field has been set.
func (o *TaskItemsImportJson) HasStorePayload() bool {
	if o != nil && !IsNil(o.StorePayload) {
		return true
	}

	return false
}

// SetStorePayload gets a reference to the given TaskItemsImportJsonStorePayload and assigns it to the StorePayload field.
func (o *TaskItemsImportJson) SetStorePayload(v TaskItemsImportJsonStorePayload) {
	o.StorePayload = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *TaskItemsImportJson) GetResults() []TaskItemsImportJsonResultsInner {
	if o == nil || IsNil(o.Results) {
		var ret []TaskItemsImportJsonResultsInner
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskItemsImportJson) GetResultsOk() ([]TaskItemsImportJsonResultsInner, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *TaskItemsImportJson) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []TaskItemsImportJsonResultsInner and assigns it to the Results field.
func (o *TaskItemsImportJson) SetResults(v []TaskItemsImportJsonResultsInner) {
	o.Results = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *TaskItemsImportJson) GetSummary() TaskItemsImportJsonSummary {
	if o == nil || IsNil(o.Summary) {
		var ret TaskItemsImportJsonSummary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskItemsImportJson) GetSummaryOk() (*TaskItemsImportJsonSummary, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *TaskItemsImportJson) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given TaskItemsImportJsonSummary and assigns it to the Summary field.
func (o *TaskItemsImportJson) SetSummary(v TaskItemsImportJsonSummary) {
	o.Summary = &v
}

func (o TaskItemsImportJson) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskItemsImportJson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StorePayload) {
		toSerialize["storePayload"] = o.StorePayload
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	return toSerialize, nil
}

type NullableTaskItemsImportJson struct {
	value *TaskItemsImportJson
	isSet bool
}

func (v NullableTaskItemsImportJson) Get() *TaskItemsImportJson {
	return v.value
}

func (v *NullableTaskItemsImportJson) Set(val *TaskItemsImportJson) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskItemsImportJson) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskItemsImportJson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskItemsImportJson(val *TaskItemsImportJson) *NullableTaskItemsImportJson {
	return &NullableTaskItemsImportJson{value: val, isSet: true}
}

func (v NullableTaskItemsImportJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskItemsImportJson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


