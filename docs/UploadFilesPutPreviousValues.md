# UploadFilesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**StorageType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**FileableType** | Pointer to **string** |  | [optional] 
**FileableId** | Pointer to **int32** |  | [optional] 
**UploadUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewUploadFilesPutPreviousValues

`func NewUploadFilesPutPreviousValues() *UploadFilesPutPreviousValues`

NewUploadFilesPutPreviousValues instantiates a new UploadFilesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadFilesPutPreviousValuesWithDefaults

`func NewUploadFilesPutPreviousValuesWithDefaults() *UploadFilesPutPreviousValues`

NewUploadFilesPutPreviousValuesWithDefaults instantiates a new UploadFilesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UploadFilesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UploadFilesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UploadFilesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UploadFilesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UploadFilesPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UploadFilesPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UploadFilesPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UploadFilesPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *UploadFilesPutPreviousValues) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UploadFilesPutPreviousValues) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UploadFilesPutPreviousValues) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UploadFilesPutPreviousValues) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetKey

`func (o *UploadFilesPutPreviousValues) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UploadFilesPutPreviousValues) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UploadFilesPutPreviousValues) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UploadFilesPutPreviousValues) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetStorageType

`func (o *UploadFilesPutPreviousValues) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *UploadFilesPutPreviousValues) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *UploadFilesPutPreviousValues) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *UploadFilesPutPreviousValues) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetType

`func (o *UploadFilesPutPreviousValues) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UploadFilesPutPreviousValues) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UploadFilesPutPreviousValues) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UploadFilesPutPreviousValues) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFileableType

`func (o *UploadFilesPutPreviousValues) GetFileableType() string`

GetFileableType returns the FileableType field if non-nil, zero value otherwise.

### GetFileableTypeOk

`func (o *UploadFilesPutPreviousValues) GetFileableTypeOk() (*string, bool)`

GetFileableTypeOk returns a tuple with the FileableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileableType

`func (o *UploadFilesPutPreviousValues) SetFileableType(v string)`

SetFileableType sets FileableType field to given value.

### HasFileableType

`func (o *UploadFilesPutPreviousValues) HasFileableType() bool`

HasFileableType returns a boolean if a field has been set.

### GetFileableId

`func (o *UploadFilesPutPreviousValues) GetFileableId() int32`

GetFileableId returns the FileableId field if non-nil, zero value otherwise.

### GetFileableIdOk

`func (o *UploadFilesPutPreviousValues) GetFileableIdOk() (*int32, bool)`

GetFileableIdOk returns a tuple with the FileableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileableId

`func (o *UploadFilesPutPreviousValues) SetFileableId(v int32)`

SetFileableId sets FileableId field to given value.

### HasFileableId

`func (o *UploadFilesPutPreviousValues) HasFileableId() bool`

HasFileableId returns a boolean if a field has been set.

### GetUploadUserId

`func (o *UploadFilesPutPreviousValues) GetUploadUserId() int32`

GetUploadUserId returns the UploadUserId field if non-nil, zero value otherwise.

### GetUploadUserIdOk

`func (o *UploadFilesPutPreviousValues) GetUploadUserIdOk() (*int32, bool)`

GetUploadUserIdOk returns a tuple with the UploadUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUserId

`func (o *UploadFilesPutPreviousValues) SetUploadUserId(v int32)`

SetUploadUserId sets UploadUserId field to given value.

### HasUploadUserId

`func (o *UploadFilesPutPreviousValues) HasUploadUserId() bool`

HasUploadUserId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UploadFilesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UploadFilesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UploadFilesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UploadFilesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UploadFilesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UploadFilesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UploadFilesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UploadFilesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *UploadFilesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *UploadFilesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *UploadFilesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *UploadFilesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


