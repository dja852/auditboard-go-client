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

// checks if the UploadFilesPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadFilesPost200Response{}

// UploadFilesPost200Response struct for UploadFilesPost200Response
type UploadFilesPost200Response struct {
	UploadFiles []UploadFiles `json:"upload_files,omitempty"`
}

// NewUploadFilesPost200Response instantiates a new UploadFilesPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadFilesPost200Response() *UploadFilesPost200Response {
	this := UploadFilesPost200Response{}
	return &this
}

// NewUploadFilesPost200ResponseWithDefaults instantiates a new UploadFilesPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadFilesPost200ResponseWithDefaults() *UploadFilesPost200Response {
	this := UploadFilesPost200Response{}
	return &this
}

// GetUploadFiles returns the UploadFiles field value if set, zero value otherwise.
func (o *UploadFilesPost200Response) GetUploadFiles() []UploadFiles {
	if o == nil || IsNil(o.UploadFiles) {
		var ret []UploadFiles
		return ret
	}
	return o.UploadFiles
}

// GetUploadFilesOk returns a tuple with the UploadFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadFilesPost200Response) GetUploadFilesOk() ([]UploadFiles, bool) {
	if o == nil || IsNil(o.UploadFiles) {
		return nil, false
	}
	return o.UploadFiles, true
}

// HasUploadFiles returns a boolean if a field has been set.
func (o *UploadFilesPost200Response) HasUploadFiles() bool {
	if o != nil && !IsNil(o.UploadFiles) {
		return true
	}

	return false
}

// SetUploadFiles gets a reference to the given []UploadFiles and assigns it to the UploadFiles field.
func (o *UploadFilesPost200Response) SetUploadFiles(v []UploadFiles) {
	o.UploadFiles = v
}

func (o UploadFilesPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadFilesPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UploadFiles) {
		toSerialize["upload_files"] = o.UploadFiles
	}
	return toSerialize, nil
}

type NullableUploadFilesPost200Response struct {
	value *UploadFilesPost200Response
	isSet bool
}

func (v NullableUploadFilesPost200Response) Get() *UploadFilesPost200Response {
	return v.value
}

func (v *NullableUploadFilesPost200Response) Set(val *UploadFilesPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadFilesPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadFilesPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadFilesPost200Response(val *UploadFilesPost200Response) *NullableUploadFilesPost200Response {
	return &NullableUploadFilesPost200Response{value: val, isSet: true}
}

func (v NullableUploadFilesPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadFilesPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

