# FileVersions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**FileId** | **int32** | Note: This is a Foreign Key to &#x60;files.id&#x60;.&lt;fk table&#x3D;&#39;files&#39; column&#x3D;&#39;id&#39;/&gt; | 
**Name** | **string** |  | 
**Size** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Key** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 
**UploadUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**StorageType** | Pointer to **string** |  | [optional] [default to "NULL::character varying"]
**EmbedUrl** | Pointer to **string** |  | [optional] 
**Version** | **int32** |  | 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**UserAgent** | Pointer to **string** |  | [optional] [default to "NULL::character varying"]
**IsPublished** | Pointer to **bool** |  | [optional] 
**VersionName** | Pointer to **string** |  | [optional] 

## Methods

### NewFileVersions

`func NewFileVersions(fileId int32, name string, type_ string, url string, version int32, ) *FileVersions`

NewFileVersions instantiates a new FileVersions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileVersionsWithDefaults

`func NewFileVersionsWithDefaults() *FileVersions`

NewFileVersionsWithDefaults instantiates a new FileVersions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileVersions) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileVersions) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileVersions) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FileVersions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFileId

`func (o *FileVersions) GetFileId() int32`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *FileVersions) GetFileIdOk() (*int32, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *FileVersions) SetFileId(v int32)`

SetFileId sets FileId field to given value.


### GetName

`func (o *FileVersions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileVersions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileVersions) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *FileVersions) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileVersions) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileVersions) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *FileVersions) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *FileVersions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileVersions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileVersions) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *FileVersions) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FileVersions) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FileVersions) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FileVersions) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetUrl

`func (o *FileVersions) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FileVersions) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FileVersions) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUploadUserId

`func (o *FileVersions) GetUploadUserId() int32`

GetUploadUserId returns the UploadUserId field if non-nil, zero value otherwise.

### GetUploadUserIdOk

`func (o *FileVersions) GetUploadUserIdOk() (*int32, bool)`

GetUploadUserIdOk returns a tuple with the UploadUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUserId

`func (o *FileVersions) SetUploadUserId(v int32)`

SetUploadUserId sets UploadUserId field to given value.

### HasUploadUserId

`func (o *FileVersions) HasUploadUserId() bool`

HasUploadUserId returns a boolean if a field has been set.

### GetStorageType

`func (o *FileVersions) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *FileVersions) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *FileVersions) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *FileVersions) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetEmbedUrl

`func (o *FileVersions) GetEmbedUrl() string`

GetEmbedUrl returns the EmbedUrl field if non-nil, zero value otherwise.

### GetEmbedUrlOk

`func (o *FileVersions) GetEmbedUrlOk() (*string, bool)`

GetEmbedUrlOk returns a tuple with the EmbedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedUrl

`func (o *FileVersions) SetEmbedUrl(v string)`

SetEmbedUrl sets EmbedUrl field to given value.

### HasEmbedUrl

`func (o *FileVersions) HasEmbedUrl() bool`

HasEmbedUrl returns a boolean if a field has been set.

### GetVersion

`func (o *FileVersions) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FileVersions) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FileVersions) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCreatedAt

`func (o *FileVersions) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FileVersions) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FileVersions) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FileVersions) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FileVersions) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FileVersions) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FileVersions) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FileVersions) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *FileVersions) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *FileVersions) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *FileVersions) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *FileVersions) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetUserAgent

`func (o *FileVersions) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *FileVersions) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *FileVersions) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *FileVersions) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetIsPublished

`func (o *FileVersions) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *FileVersions) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *FileVersions) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *FileVersions) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetVersionName

`func (o *FileVersions) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *FileVersions) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *FileVersions) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *FileVersions) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


