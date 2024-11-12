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

// checks if the FileVersionsPutFileVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileVersionsPutFileVersion{}

// FileVersionsPutFileVersion struct for FileVersionsPutFileVersion
type FileVersionsPutFileVersion struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	// Note: This is a Foreign Key to `files.id`.<fk table='files' column='id'/>
	FileId int32 `json:"file_id"`
	Name string `json:"name"`
	Size *string `json:"size,omitempty"`
	Type string `json:"type"`
	Key *string `json:"key,omitempty"`
	Url string `json:"url"`
	// Note: This is a Foreign Key to `users.id`.<fk table='users' column='id'/>
	UploadUserId *int32 `json:"upload_user_id,omitempty"`
	StorageType *string `json:"storage_type,omitempty"`
	EmbedUrl *string `json:"embed_url,omitempty"`
	Version int32 `json:"version"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	UserAgent *string `json:"user_agent,omitempty"`
	IsPublished *bool `json:"is_published,omitempty"`
	VersionName *string `json:"version_name,omitempty"`
}

type _FileVersionsPutFileVersion FileVersionsPutFileVersion

// NewFileVersionsPutFileVersion instantiates a new FileVersionsPutFileVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileVersionsPutFileVersion(fileId int32, name string, type_ string, url string, version int32) *FileVersionsPutFileVersion {
	this := FileVersionsPutFileVersion{}
	this.FileId = fileId
	this.Name = name
	this.Type = type_
	this.Url = url
	var storageType string = "NULL::character varying"
	this.StorageType = &storageType
	this.Version = version
	var userAgent string = "NULL::character varying"
	this.UserAgent = &userAgent
	return &this
}

// NewFileVersionsPutFileVersionWithDefaults instantiates a new FileVersionsPutFileVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileVersionsPutFileVersionWithDefaults() *FileVersionsPutFileVersion {
	this := FileVersionsPutFileVersion{}
	var storageType string = "NULL::character varying"
	this.StorageType = &storageType
	var userAgent string = "NULL::character varying"
	this.UserAgent = &userAgent
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FileVersionsPutFileVersion) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVersionsPutFileVersion) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FileVersionsPutFileVersion) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FileVersionsPutFileVersion) SetId(v int32) {
	o.Id = &v
}

// GetFileId returns the FileId field value
func (o *FileVersionsPutFileVersion) GetFileId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value
// and a boolean to check if the value has been set.
func (o *FileVersionsPutFileVersion) GetFileIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileId, true
}

// SetFileId sets field value
func (o *FileVersionsPutFileVersion) SetFileId(v int32) {
	o.FileId = v
}

// GetName returns the Name field value
func (o *FileVersionsPutFileVersion) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FileVersionsPutFileVersion) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FileVersionsPutFileVersion) SetName(v string) {
	o.Name = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *FileVersionsPutFileVersion) GetSize() string {
	if o == nil || IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVersionsPutFileVersion) GetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *FileVersionsPutFileVersion) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *FileVersionsPutFileVersion) SetSize(v string) {
	o.Size = &v
}

// GetType returns the Type field value
func (o *FileVersionsPutFileVersion) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FileVersionsPutFileVersion) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FileVersionsPutFileVersion) SetType(v string) {
	o.Type = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *FileVersionsPutFileVersion) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVersionsPutFileVersion) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *FileVersionsPutFileVersion) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *FileVersionsPutFileVersion) SetKey(v string) {
	o.Key = &v
}

// GetUrl returns the Url field value
func (o *FileVersionsPutFileVersion) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *FileVersionsPutFileVersion) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *FileVersionsPutFileVersion) SetUrl(v string) {
	o.Url = v
}

// GetUploadUserId returns the UploadUserId field value if set, zero value otherwise.
func (o *FileVersionsPutFileVersion) GetUploadUserId() int32 {
	if o == nil || IsNil(o.UploadUserId) {
		var ret int32
		return ret
	}
	return *o.UploadUserId
}

// GetUploadUserIdOk returns a tuple with the UploadUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVersionsPutFileVersion) GetUploadUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UploadUserId) {
		return nil, false
	}
	return o.UploadUserId, true
}

// HasUploadUserId returns a boolean if a field has been set.
func (o *FileVersionsPutFileVersion) HasUploadUserId() bool {
	if o != nil && !IsNil(o.UploadUserId) {
		return true
	}

	return false
}

// SetUploadUserId gets a reference to the given int32 and assigns it to the UploadUserId field.
func (o *FileVersionsPutFileVersion) SetUploadUserId(v int32) {
	o.UploadUserId = &v
}

// GetStorageType returns the StorageType field value if set, zero value otherwise.
func (o *FileVersionsPutFileVersion) GetStorageType() string {
	if o == nil || IsNil(o.StorageType) {
		var ret string
		return ret
	}
	return *o.StorageType
}

// GetStorageTypeOk returns a tuple with the StorageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVersionsPutFileVersion) GetStorageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.StorageType) {
		return nil, false
	}
	return o.StorageType, true
}

// HasStorageType returns a boolean if a field has been set.
func (o *FileVersionsPutFileVersion) HasStorageType() bool {
	if o != nil && !IsNil(o.StorageType) {
		return true
	}

	return false
}

// SetStorageType gets a reference to the given string and assigns it to the StorageType field.
func (o *FileVersionsPutFileVersion) SetStorageType(v string) {
	o.StorageType = &v
}

// GetEmbedUrl returns the EmbedUrl field value if set, zero value otherwise.
func (o *FileVersionsPutFileVersion) GetEmbedUrl() string {
	if o == nil || IsNil(o.EmbedUrl) {
		var ret string
		return ret
	}
	return *o.EmbedUrl
}

// GetEmbedUrlOk returns a tuple with the EmbedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVersionsPutFileVersion) GetEmbedUrlOk() (*string, bool) {
	if o == nil || IsNil(o.EmbedUrl) {
		return nil, false
	}
	return o.EmbedUrl, true
}

// HasEmbedUrl returns a boolean if a field has been set.
func (o *FileVersionsPutFileVersion) HasEmbedUrl() bool {
	if o != nil && !IsNil(o.EmbedUrl) {
		return true
	}

	return false
}

// SetEmbedUrl gets a reference to the given string and assigns it to the EmbedUrl field.
func (o *FileVersionsPutFileVersion) SetEmbedUrl(v string) {
	o.EmbedUrl = &v
}

// GetVersion returns the Version field value
func (o *FileVersionsPutFileVersion) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *FileVersionsPutFileVersion) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *FileVersionsPutFileVersion) SetVersion(v int32) {
	o.Version = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FileVersionsPutFileVersion) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVersionsPutFileVersion) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FileVersionsPutFileVersion) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *FileVersionsPutFileVersion) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FileVersionsPutFileVersion) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVersionsPutFileVersion) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FileVersionsPutFileVersion) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *FileVersionsPutFileVersion) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *FileVersionsPutFileVersion) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVersionsPutFileVersion) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *FileVersionsPutFileVersion) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *FileVersionsPutFileVersion) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *FileVersionsPutFileVersion) GetUserAgent() string {
	if o == nil || IsNil(o.UserAgent) {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVersionsPutFileVersion) GetUserAgentOk() (*string, bool) {
	if o == nil || IsNil(o.UserAgent) {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *FileVersionsPutFileVersion) HasUserAgent() bool {
	if o != nil && !IsNil(o.UserAgent) {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *FileVersionsPutFileVersion) SetUserAgent(v string) {
	o.UserAgent = &v
}

// GetIsPublished returns the IsPublished field value if set, zero value otherwise.
func (o *FileVersionsPutFileVersion) GetIsPublished() bool {
	if o == nil || IsNil(o.IsPublished) {
		var ret bool
		return ret
	}
	return *o.IsPublished
}

// GetIsPublishedOk returns a tuple with the IsPublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVersionsPutFileVersion) GetIsPublishedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublished) {
		return nil, false
	}
	return o.IsPublished, true
}

// HasIsPublished returns a boolean if a field has been set.
func (o *FileVersionsPutFileVersion) HasIsPublished() bool {
	if o != nil && !IsNil(o.IsPublished) {
		return true
	}

	return false
}

// SetIsPublished gets a reference to the given bool and assigns it to the IsPublished field.
func (o *FileVersionsPutFileVersion) SetIsPublished(v bool) {
	o.IsPublished = &v
}

// GetVersionName returns the VersionName field value if set, zero value otherwise.
func (o *FileVersionsPutFileVersion) GetVersionName() string {
	if o == nil || IsNil(o.VersionName) {
		var ret string
		return ret
	}
	return *o.VersionName
}

// GetVersionNameOk returns a tuple with the VersionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVersionsPutFileVersion) GetVersionNameOk() (*string, bool) {
	if o == nil || IsNil(o.VersionName) {
		return nil, false
	}
	return o.VersionName, true
}

// HasVersionName returns a boolean if a field has been set.
func (o *FileVersionsPutFileVersion) HasVersionName() bool {
	if o != nil && !IsNil(o.VersionName) {
		return true
	}

	return false
}

// SetVersionName gets a reference to the given string and assigns it to the VersionName field.
func (o *FileVersionsPutFileVersion) SetVersionName(v string) {
	o.VersionName = &v
}

func (o FileVersionsPutFileVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileVersionsPutFileVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["file_id"] = o.FileId
	toSerialize["name"] = o.Name
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	toSerialize["url"] = o.Url
	if !IsNil(o.UploadUserId) {
		toSerialize["upload_user_id"] = o.UploadUserId
	}
	if !IsNil(o.StorageType) {
		toSerialize["storage_type"] = o.StorageType
	}
	if !IsNil(o.EmbedUrl) {
		toSerialize["embed_url"] = o.EmbedUrl
	}
	toSerialize["version"] = o.Version
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.UserAgent) {
		toSerialize["user_agent"] = o.UserAgent
	}
	if !IsNil(o.IsPublished) {
		toSerialize["is_published"] = o.IsPublished
	}
	if !IsNil(o.VersionName) {
		toSerialize["version_name"] = o.VersionName
	}
	return toSerialize, nil
}

type NullableFileVersionsPutFileVersion struct {
	value *FileVersionsPutFileVersion
	isSet bool
}

func (v NullableFileVersionsPutFileVersion) Get() *FileVersionsPutFileVersion {
	return v.value
}

func (v *NullableFileVersionsPutFileVersion) Set(val *FileVersionsPutFileVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableFileVersionsPutFileVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableFileVersionsPutFileVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileVersionsPutFileVersion(val *FileVersionsPutFileVersion) *NullableFileVersionsPutFileVersion {
	return &NullableFileVersionsPutFileVersion{value: val, isSet: true}
}

func (v NullableFileVersionsPutFileVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileVersionsPutFileVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


