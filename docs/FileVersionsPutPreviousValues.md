# FileVersionsPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**FileId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;files.id&#x60;.&lt;fk table&#x3D;&#39;files&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UploadUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**StorageType** | Pointer to **string** |  | [optional] [default to "NULL::character varying"]
**EmbedUrl** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**UserAgent** | Pointer to **string** |  | [optional] [default to "NULL::character varying"]
**IsPublished** | Pointer to **bool** |  | [optional] 
**VersionName** | Pointer to **string** |  | [optional] 

## Methods

### NewFileVersionsPutPreviousValues

`func NewFileVersionsPutPreviousValues() *FileVersionsPutPreviousValues`

NewFileVersionsPutPreviousValues instantiates a new FileVersionsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileVersionsPutPreviousValuesWithDefaults

`func NewFileVersionsPutPreviousValuesWithDefaults() *FileVersionsPutPreviousValues`

NewFileVersionsPutPreviousValuesWithDefaults instantiates a new FileVersionsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileVersionsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileVersionsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileVersionsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FileVersionsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFileId

`func (o *FileVersionsPutPreviousValues) GetFileId() int32`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *FileVersionsPutPreviousValues) GetFileIdOk() (*int32, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *FileVersionsPutPreviousValues) SetFileId(v int32)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *FileVersionsPutPreviousValues) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetName

`func (o *FileVersionsPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileVersionsPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileVersionsPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileVersionsPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *FileVersionsPutPreviousValues) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileVersionsPutPreviousValues) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileVersionsPutPreviousValues) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *FileVersionsPutPreviousValues) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *FileVersionsPutPreviousValues) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileVersionsPutPreviousValues) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileVersionsPutPreviousValues) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FileVersionsPutPreviousValues) HasType() bool`

HasType returns a boolean if a field has been set.

### GetKey

`func (o *FileVersionsPutPreviousValues) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FileVersionsPutPreviousValues) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FileVersionsPutPreviousValues) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FileVersionsPutPreviousValues) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetUrl

`func (o *FileVersionsPutPreviousValues) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FileVersionsPutPreviousValues) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FileVersionsPutPreviousValues) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FileVersionsPutPreviousValues) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUploadUserId

`func (o *FileVersionsPutPreviousValues) GetUploadUserId() int32`

GetUploadUserId returns the UploadUserId field if non-nil, zero value otherwise.

### GetUploadUserIdOk

`func (o *FileVersionsPutPreviousValues) GetUploadUserIdOk() (*int32, bool)`

GetUploadUserIdOk returns a tuple with the UploadUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUserId

`func (o *FileVersionsPutPreviousValues) SetUploadUserId(v int32)`

SetUploadUserId sets UploadUserId field to given value.

### HasUploadUserId

`func (o *FileVersionsPutPreviousValues) HasUploadUserId() bool`

HasUploadUserId returns a boolean if a field has been set.

### GetStorageType

`func (o *FileVersionsPutPreviousValues) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *FileVersionsPutPreviousValues) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *FileVersionsPutPreviousValues) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *FileVersionsPutPreviousValues) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetEmbedUrl

`func (o *FileVersionsPutPreviousValues) GetEmbedUrl() string`

GetEmbedUrl returns the EmbedUrl field if non-nil, zero value otherwise.

### GetEmbedUrlOk

`func (o *FileVersionsPutPreviousValues) GetEmbedUrlOk() (*string, bool)`

GetEmbedUrlOk returns a tuple with the EmbedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedUrl

`func (o *FileVersionsPutPreviousValues) SetEmbedUrl(v string)`

SetEmbedUrl sets EmbedUrl field to given value.

### HasEmbedUrl

`func (o *FileVersionsPutPreviousValues) HasEmbedUrl() bool`

HasEmbedUrl returns a boolean if a field has been set.

### GetVersion

`func (o *FileVersionsPutPreviousValues) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FileVersionsPutPreviousValues) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FileVersionsPutPreviousValues) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FileVersionsPutPreviousValues) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FileVersionsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FileVersionsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FileVersionsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FileVersionsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FileVersionsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FileVersionsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FileVersionsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FileVersionsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *FileVersionsPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *FileVersionsPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *FileVersionsPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *FileVersionsPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetUserAgent

`func (o *FileVersionsPutPreviousValues) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *FileVersionsPutPreviousValues) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *FileVersionsPutPreviousValues) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *FileVersionsPutPreviousValues) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetIsPublished

`func (o *FileVersionsPutPreviousValues) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *FileVersionsPutPreviousValues) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *FileVersionsPutPreviousValues) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *FileVersionsPutPreviousValues) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetVersionName

`func (o *FileVersionsPutPreviousValues) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *FileVersionsPutPreviousValues) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *FileVersionsPutPreviousValues) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *FileVersionsPutPreviousValues) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


