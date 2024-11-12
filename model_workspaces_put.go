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

// checks if the WorkspacesPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkspacesPut{}

// WorkspacesPut struct for WorkspacesPut
type WorkspacesPut struct {
	Workspace *WorkspacesPutWorkspace `json:"workspace,omitempty"`
	PreviousValues *WorkspacesPutPreviousValues `json:"previous_values,omitempty"`
}

// NewWorkspacesPut instantiates a new WorkspacesPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspacesPut() *WorkspacesPut {
	this := WorkspacesPut{}
	return &this
}

// NewWorkspacesPutWithDefaults instantiates a new WorkspacesPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspacesPutWithDefaults() *WorkspacesPut {
	this := WorkspacesPut{}
	return &this
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise.
func (o *WorkspacesPut) GetWorkspace() WorkspacesPutWorkspace {
	if o == nil || IsNil(o.Workspace) {
		var ret WorkspacesPutWorkspace
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspacesPut) GetWorkspaceOk() (*WorkspacesPutWorkspace, bool) {
	if o == nil || IsNil(o.Workspace) {
		return nil, false
	}
	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *WorkspacesPut) HasWorkspace() bool {
	if o != nil && !IsNil(o.Workspace) {
		return true
	}

	return false
}

// SetWorkspace gets a reference to the given WorkspacesPutWorkspace and assigns it to the Workspace field.
func (o *WorkspacesPut) SetWorkspace(v WorkspacesPutWorkspace) {
	o.Workspace = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *WorkspacesPut) GetPreviousValues() WorkspacesPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret WorkspacesPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspacesPut) GetPreviousValuesOk() (*WorkspacesPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *WorkspacesPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given WorkspacesPutPreviousValues and assigns it to the PreviousValues field.
func (o *WorkspacesPut) SetPreviousValues(v WorkspacesPutPreviousValues) {
	o.PreviousValues = &v
}

func (o WorkspacesPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkspacesPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Workspace) {
		toSerialize["workspace"] = o.Workspace
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableWorkspacesPut struct {
	value *WorkspacesPut
	isSet bool
}

func (v NullableWorkspacesPut) Get() *WorkspacesPut {
	return v.value
}

func (v *NullableWorkspacesPut) Set(val *WorkspacesPut) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspacesPut) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspacesPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspacesPut(val *WorkspacesPut) *NullableWorkspacesPut {
	return &NullableWorkspacesPut{value: val, isSet: true}
}

func (v NullableWorkspacesPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspacesPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

