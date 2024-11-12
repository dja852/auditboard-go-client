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

// checks if the FilesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesPostRequest{}

// FilesPostRequest struct for FilesPostRequest
type FilesPostRequest struct {
	File *Files `json:"file,omitempty"`
}

// NewFilesPostRequest instantiates a new FilesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesPostRequest() *FilesPostRequest {
	this := FilesPostRequest{}
	return &this
}

// NewFilesPostRequestWithDefaults instantiates a new FilesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesPostRequestWithDefaults() *FilesPostRequest {
	this := FilesPostRequest{}
	return &this
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *FilesPostRequest) GetFile() Files {
	if o == nil || IsNil(o.File) {
		var ret Files
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesPostRequest) GetFileOk() (*Files, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *FilesPostRequest) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given Files and assigns it to the File field.
func (o *FilesPostRequest) SetFile(v Files) {
	o.File = &v
}

func (o FilesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}
	return toSerialize, nil
}

type NullableFilesPostRequest struct {
	value *FilesPostRequest
	isSet bool
}

func (v NullableFilesPostRequest) Get() *FilesPostRequest {
	return v.value
}

func (v *NullableFilesPostRequest) Set(val *FilesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesPostRequest(val *FilesPostRequest) *NullableFilesPostRequest {
	return &NullableFilesPostRequest{value: val, isSet: true}
}

func (v NullableFilesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

