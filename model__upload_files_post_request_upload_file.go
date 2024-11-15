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

// checks if the UploadFilesPostRequestUploadFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadFilesPostRequestUploadFile{}

// UploadFilesPostRequestUploadFile struct for UploadFilesPostRequestUploadFile
type UploadFilesPostRequestUploadFile struct {
	// AuditBoard models have a fileable type associated with the model. The fileable type will determine what model type the file will be uploaded (i.e. Issues module) and the fileable id will determine the specific model the file will be uploaded to.
	FileableType []string `json:"fileable_type,omitempty"`
	// The fileable_id determines the actual model where the file will be uploaded. Ex. \"21\" with fileable_type \"Issue\" will upload the file to Issue with id of 21.
	FileableId *int32 `json:"fileable_id,omitempty"`
	Name *string `json:"name,omitempty"`
	// The type field should be the type of the file that is being uploaded. List of possible values: https://www.iana.org/assignments/media-types/media-types.xhtml
	Type *string `json:"type,omitempty"`
}

// NewUploadFilesPostRequestUploadFile instantiates a new UploadFilesPostRequestUploadFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadFilesPostRequestUploadFile() *UploadFilesPostRequestUploadFile {
	this := UploadFilesPostRequestUploadFile{}
	return &this
}

// NewUploadFilesPostRequestUploadFileWithDefaults instantiates a new UploadFilesPostRequestUploadFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadFilesPostRequestUploadFileWithDefaults() *UploadFilesPostRequestUploadFile {
	this := UploadFilesPostRequestUploadFile{}
	return &this
}

// GetFileableType returns the FileableType field value if set, zero value otherwise.
func (o *UploadFilesPostRequestUploadFile) GetFileableType() []string {
	if o == nil || IsNil(o.FileableType) {
		var ret []string
		return ret
	}
	return o.FileableType
}

// GetFileableTypeOk returns a tuple with the FileableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadFilesPostRequestUploadFile) GetFileableTypeOk() ([]string, bool) {
	if o == nil || IsNil(o.FileableType) {
		return nil, false
	}
	return o.FileableType, true
}

// HasFileableType returns a boolean if a field has been set.
func (o *UploadFilesPostRequestUploadFile) HasFileableType() bool {
	if o != nil && !IsNil(o.FileableType) {
		return true
	}

	return false
}

// SetFileableType gets a reference to the given []string and assigns it to the FileableType field.
func (o *UploadFilesPostRequestUploadFile) SetFileableType(v []string) {
	o.FileableType = v
}

// GetFileableId returns the FileableId field value if set, zero value otherwise.
func (o *UploadFilesPostRequestUploadFile) GetFileableId() int32 {
	if o == nil || IsNil(o.FileableId) {
		var ret int32
		return ret
	}
	return *o.FileableId
}

// GetFileableIdOk returns a tuple with the FileableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadFilesPostRequestUploadFile) GetFileableIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FileableId) {
		return nil, false
	}
	return o.FileableId, true
}

// HasFileableId returns a boolean if a field has been set.
func (o *UploadFilesPostRequestUploadFile) HasFileableId() bool {
	if o != nil && !IsNil(o.FileableId) {
		return true
	}

	return false
}

// SetFileableId gets a reference to the given int32 and assigns it to the FileableId field.
func (o *UploadFilesPostRequestUploadFile) SetFileableId(v int32) {
	o.FileableId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UploadFilesPostRequestUploadFile) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadFilesPostRequestUploadFile) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UploadFilesPostRequestUploadFile) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UploadFilesPostRequestUploadFile) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UploadFilesPostRequestUploadFile) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadFilesPostRequestUploadFile) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UploadFilesPostRequestUploadFile) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UploadFilesPostRequestUploadFile) SetType(v string) {
	o.Type = &v
}

func (o UploadFilesPostRequestUploadFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadFilesPostRequestUploadFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileableType) {
		toSerialize["fileable_type"] = o.FileableType
	}
	if !IsNil(o.FileableId) {
		toSerialize["fileable_id"] = o.FileableId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableUploadFilesPostRequestUploadFile struct {
	value *UploadFilesPostRequestUploadFile
	isSet bool
}

func (v NullableUploadFilesPostRequestUploadFile) Get() *UploadFilesPostRequestUploadFile {
	return v.value
}

func (v *NullableUploadFilesPostRequestUploadFile) Set(val *UploadFilesPostRequestUploadFile) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadFilesPostRequestUploadFile) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadFilesPostRequestUploadFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadFilesPostRequestUploadFile(val *UploadFilesPostRequestUploadFile) *NullableUploadFilesPostRequestUploadFile {
	return &NullableUploadFilesPostRequestUploadFile{value: val, isSet: true}
}

func (v NullableUploadFilesPostRequestUploadFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadFilesPostRequestUploadFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


