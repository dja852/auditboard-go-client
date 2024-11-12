# UploadFilesPostRequestUploadFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileableType** | Pointer to **[]string** | AuditBoard models have a fileable type associated with the model. The fileable type will determine what model type the file will be uploaded (i.e. Issues module) and the fileable id will determine the specific model the file will be uploaded to. | [optional] 
**FileableId** | Pointer to **int32** | The fileable_id determines the actual model where the file will be uploaded. Ex. \&quot;21\&quot; with fileable_type \&quot;Issue\&quot; will upload the file to Issue with id of 21. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | The type field should be the type of the file that is being uploaded. List of possible values: https://www.iana.org/assignments/media-types/media-types.xhtml | [optional] 

## Methods

### NewUploadFilesPostRequestUploadFile

`func NewUploadFilesPostRequestUploadFile() *UploadFilesPostRequestUploadFile`

NewUploadFilesPostRequestUploadFile instantiates a new UploadFilesPostRequestUploadFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadFilesPostRequestUploadFileWithDefaults

`func NewUploadFilesPostRequestUploadFileWithDefaults() *UploadFilesPostRequestUploadFile`

NewUploadFilesPostRequestUploadFileWithDefaults instantiates a new UploadFilesPostRequestUploadFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileableType

`func (o *UploadFilesPostRequestUploadFile) GetFileableType() []string`

GetFileableType returns the FileableType field if non-nil, zero value otherwise.

### GetFileableTypeOk

`func (o *UploadFilesPostRequestUploadFile) GetFileableTypeOk() (*[]string, bool)`

GetFileableTypeOk returns a tuple with the FileableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileableType

`func (o *UploadFilesPostRequestUploadFile) SetFileableType(v []string)`

SetFileableType sets FileableType field to given value.

### HasFileableType

`func (o *UploadFilesPostRequestUploadFile) HasFileableType() bool`

HasFileableType returns a boolean if a field has been set.

### GetFileableId

`func (o *UploadFilesPostRequestUploadFile) GetFileableId() int32`

GetFileableId returns the FileableId field if non-nil, zero value otherwise.

### GetFileableIdOk

`func (o *UploadFilesPostRequestUploadFile) GetFileableIdOk() (*int32, bool)`

GetFileableIdOk returns a tuple with the FileableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileableId

`func (o *UploadFilesPostRequestUploadFile) SetFileableId(v int32)`

SetFileableId sets FileableId field to given value.

### HasFileableId

`func (o *UploadFilesPostRequestUploadFile) HasFileableId() bool`

HasFileableId returns a boolean if a field has been set.

### GetName

`func (o *UploadFilesPostRequestUploadFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UploadFilesPostRequestUploadFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UploadFilesPostRequestUploadFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UploadFilesPostRequestUploadFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *UploadFilesPostRequestUploadFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UploadFilesPostRequestUploadFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UploadFilesPostRequestUploadFile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UploadFilesPostRequestUploadFile) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


