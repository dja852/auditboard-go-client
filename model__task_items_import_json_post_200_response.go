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

// checks if the TaskItemsImportJsonPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskItemsImportJsonPost200Response{}

// TaskItemsImportJsonPost200Response struct for TaskItemsImportJsonPost200Response
type TaskItemsImportJsonPost200Response struct {
	OpsAudits []TaskItemsImportJson `json:"ops_audits,omitempty"`
}

// NewTaskItemsImportJsonPost200Response instantiates a new TaskItemsImportJsonPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskItemsImportJsonPost200Response() *TaskItemsImportJsonPost200Response {
	this := TaskItemsImportJsonPost200Response{}
	return &this
}

// NewTaskItemsImportJsonPost200ResponseWithDefaults instantiates a new TaskItemsImportJsonPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskItemsImportJsonPost200ResponseWithDefaults() *TaskItemsImportJsonPost200Response {
	this := TaskItemsImportJsonPost200Response{}
	return &this
}

// GetOpsAudits returns the OpsAudits field value if set, zero value otherwise.
func (o *TaskItemsImportJsonPost200Response) GetOpsAudits() []TaskItemsImportJson {
	if o == nil || IsNil(o.OpsAudits) {
		var ret []TaskItemsImportJson
		return ret
	}
	return o.OpsAudits
}

// GetOpsAuditsOk returns a tuple with the OpsAudits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskItemsImportJsonPost200Response) GetOpsAuditsOk() ([]TaskItemsImportJson, bool) {
	if o == nil || IsNil(o.OpsAudits) {
		return nil, false
	}
	return o.OpsAudits, true
}

// HasOpsAudits returns a boolean if a field has been set.
func (o *TaskItemsImportJsonPost200Response) HasOpsAudits() bool {
	if o != nil && !IsNil(o.OpsAudits) {
		return true
	}

	return false
}

// SetOpsAudits gets a reference to the given []TaskItemsImportJson and assigns it to the OpsAudits field.
func (o *TaskItemsImportJsonPost200Response) SetOpsAudits(v []TaskItemsImportJson) {
	o.OpsAudits = v
}

func (o TaskItemsImportJsonPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskItemsImportJsonPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OpsAudits) {
		toSerialize["ops_audits"] = o.OpsAudits
	}
	return toSerialize, nil
}

type NullableTaskItemsImportJsonPost200Response struct {
	value *TaskItemsImportJsonPost200Response
	isSet bool
}

func (v NullableTaskItemsImportJsonPost200Response) Get() *TaskItemsImportJsonPost200Response {
	return v.value
}

func (v *NullableTaskItemsImportJsonPost200Response) Set(val *TaskItemsImportJsonPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskItemsImportJsonPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskItemsImportJsonPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskItemsImportJsonPost200Response(val *TaskItemsImportJsonPost200Response) *NullableTaskItemsImportJsonPost200Response {
	return &NullableTaskItemsImportJsonPost200Response{value: val, isSet: true}
}

func (v NullableTaskItemsImportJsonPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskItemsImportJsonPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


