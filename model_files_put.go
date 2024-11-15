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

// checks if the FilesPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesPut{}

// FilesPut struct for FilesPut
type FilesPut struct {
	File *FilesPutFile `json:"file,omitempty"`
	PreviousValues *FilesPutPreviousValues `json:"previous_values,omitempty"`
}

// NewFilesPut instantiates a new FilesPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesPut() *FilesPut {
	this := FilesPut{}
	return &this
}

// NewFilesPutWithDefaults instantiates a new FilesPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesPutWithDefaults() *FilesPut {
	this := FilesPut{}
	return &this
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *FilesPut) GetFile() FilesPutFile {
	if o == nil || IsNil(o.File) {
		var ret FilesPutFile
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesPut) GetFileOk() (*FilesPutFile, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *FilesPut) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given FilesPutFile and assigns it to the File field.
func (o *FilesPut) SetFile(v FilesPutFile) {
	o.File = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *FilesPut) GetPreviousValues() FilesPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret FilesPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesPut) GetPreviousValuesOk() (*FilesPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *FilesPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given FilesPutPreviousValues and assigns it to the PreviousValues field.
func (o *FilesPut) SetPreviousValues(v FilesPutPreviousValues) {
	o.PreviousValues = &v
}

func (o FilesPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableFilesPut struct {
	value *FilesPut
	isSet bool
}

func (v NullableFilesPut) Get() *FilesPut {
	return v.value
}

func (v *NullableFilesPut) Set(val *FilesPut) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesPut) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesPut(val *FilesPut) *NullableFilesPut {
	return &NullableFilesPut{value: val, isSet: true}
}

func (v NullableFilesPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


