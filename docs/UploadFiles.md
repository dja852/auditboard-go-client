# UploadFiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Name** | **string** |  | 
**Url** | **string** |  | 
**Key** | **string** |  | 
**StorageType** | **string** |  | 
**Type** | **string** |  | 
**FileableType** | Pointer to **string** |  | [optional] 
**FileableId** | Pointer to **int32** |  | [optional] 
**UploadUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewUploadFiles

`func NewUploadFiles(name string, url string, key string, storageType string, type_ string, ) *UploadFiles`

NewUploadFiles instantiates a new UploadFiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadFilesWithDefaults

`func NewUploadFilesWithDefaults() *UploadFiles`

NewUploadFilesWithDefaults instantiates a new UploadFiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UploadFiles) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UploadFiles) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UploadFiles) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UploadFiles) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UploadFiles) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UploadFiles) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UploadFiles) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *UploadFiles) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UploadFiles) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UploadFiles) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetKey

`func (o *UploadFiles) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UploadFiles) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UploadFiles) SetKey(v string)`

SetKey sets Key field to given value.


### GetStorageType

`func (o *UploadFiles) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *UploadFiles) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *UploadFiles) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.


### GetType

`func (o *UploadFiles) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UploadFiles) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UploadFiles) SetType(v string)`

SetType sets Type field to given value.


### GetFileableType

`func (o *UploadFiles) GetFileableType() string`

GetFileableType returns the FileableType field if non-nil, zero value otherwise.

### GetFileableTypeOk

`func (o *UploadFiles) GetFileableTypeOk() (*string, bool)`

GetFileableTypeOk returns a tuple with the FileableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileableType

`func (o *UploadFiles) SetFileableType(v string)`

SetFileableType sets FileableType field to given value.

### HasFileableType

`func (o *UploadFiles) HasFileableType() bool`

HasFileableType returns a boolean if a field has been set.

### GetFileableId

`func (o *UploadFiles) GetFileableId() int32`

GetFileableId returns the FileableId field if non-nil, zero value otherwise.

### GetFileableIdOk

`func (o *UploadFiles) GetFileableIdOk() (*int32, bool)`

GetFileableIdOk returns a tuple with the FileableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileableId

`func (o *UploadFiles) SetFileableId(v int32)`

SetFileableId sets FileableId field to given value.

### HasFileableId

`func (o *UploadFiles) HasFileableId() bool`

HasFileableId returns a boolean if a field has been set.

### GetUploadUserId

`func (o *UploadFiles) GetUploadUserId() int32`

GetUploadUserId returns the UploadUserId field if non-nil, zero value otherwise.

### GetUploadUserIdOk

`func (o *UploadFiles) GetUploadUserIdOk() (*int32, bool)`

GetUploadUserIdOk returns a tuple with the UploadUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUserId

`func (o *UploadFiles) SetUploadUserId(v int32)`

SetUploadUserId sets UploadUserId field to given value.

### HasUploadUserId

`func (o *UploadFiles) HasUploadUserId() bool`

HasUploadUserId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UploadFiles) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UploadFiles) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UploadFiles) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UploadFiles) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UploadFiles) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UploadFiles) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UploadFiles) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UploadFiles) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *UploadFiles) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *UploadFiles) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *UploadFiles) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *UploadFiles) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


