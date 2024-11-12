# UploadFilesPutUploadFile

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

### NewUploadFilesPutUploadFile

`func NewUploadFilesPutUploadFile(name string, url string, key string, storageType string, type_ string, ) *UploadFilesPutUploadFile`

NewUploadFilesPutUploadFile instantiates a new UploadFilesPutUploadFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadFilesPutUploadFileWithDefaults

`func NewUploadFilesPutUploadFileWithDefaults() *UploadFilesPutUploadFile`

NewUploadFilesPutUploadFileWithDefaults instantiates a new UploadFilesPutUploadFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UploadFilesPutUploadFile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UploadFilesPutUploadFile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UploadFilesPutUploadFile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UploadFilesPutUploadFile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UploadFilesPutUploadFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UploadFilesPutUploadFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UploadFilesPutUploadFile) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *UploadFilesPutUploadFile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UploadFilesPutUploadFile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UploadFilesPutUploadFile) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetKey

`func (o *UploadFilesPutUploadFile) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UploadFilesPutUploadFile) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UploadFilesPutUploadFile) SetKey(v string)`

SetKey sets Key field to given value.


### GetStorageType

`func (o *UploadFilesPutUploadFile) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *UploadFilesPutUploadFile) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *UploadFilesPutUploadFile) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.


### GetType

`func (o *UploadFilesPutUploadFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UploadFilesPutUploadFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UploadFilesPutUploadFile) SetType(v string)`

SetType sets Type field to given value.


### GetFileableType

`func (o *UploadFilesPutUploadFile) GetFileableType() string`

GetFileableType returns the FileableType field if non-nil, zero value otherwise.

### GetFileableTypeOk

`func (o *UploadFilesPutUploadFile) GetFileableTypeOk() (*string, bool)`

GetFileableTypeOk returns a tuple with the FileableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileableType

`func (o *UploadFilesPutUploadFile) SetFileableType(v string)`

SetFileableType sets FileableType field to given value.

### HasFileableType

`func (o *UploadFilesPutUploadFile) HasFileableType() bool`

HasFileableType returns a boolean if a field has been set.

### GetFileableId

`func (o *UploadFilesPutUploadFile) GetFileableId() int32`

GetFileableId returns the FileableId field if non-nil, zero value otherwise.

### GetFileableIdOk

`func (o *UploadFilesPutUploadFile) GetFileableIdOk() (*int32, bool)`

GetFileableIdOk returns a tuple with the FileableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileableId

`func (o *UploadFilesPutUploadFile) SetFileableId(v int32)`

SetFileableId sets FileableId field to given value.

### HasFileableId

`func (o *UploadFilesPutUploadFile) HasFileableId() bool`

HasFileableId returns a boolean if a field has been set.

### GetUploadUserId

`func (o *UploadFilesPutUploadFile) GetUploadUserId() int32`

GetUploadUserId returns the UploadUserId field if non-nil, zero value otherwise.

### GetUploadUserIdOk

`func (o *UploadFilesPutUploadFile) GetUploadUserIdOk() (*int32, bool)`

GetUploadUserIdOk returns a tuple with the UploadUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUserId

`func (o *UploadFilesPutUploadFile) SetUploadUserId(v int32)`

SetUploadUserId sets UploadUserId field to given value.

### HasUploadUserId

`func (o *UploadFilesPutUploadFile) HasUploadUserId() bool`

HasUploadUserId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UploadFilesPutUploadFile) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UploadFilesPutUploadFile) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UploadFilesPutUploadFile) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UploadFilesPutUploadFile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UploadFilesPutUploadFile) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UploadFilesPutUploadFile) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UploadFilesPutUploadFile) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UploadFilesPutUploadFile) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *UploadFilesPutUploadFile) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *UploadFilesPutUploadFile) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *UploadFilesPutUploadFile) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *UploadFilesPutUploadFile) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


