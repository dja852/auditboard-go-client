# Files

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**FileableId** | Pointer to **int32** |  | [optional] 
**FileableType** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Size** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Key** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 
**UploadUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**StorageType** | Pointer to **string** |  | [optional] [default to "NULL::character varying"]
**EmbedUrl** | Pointer to **string** |  | [optional] 
**ThumbUrl** | Pointer to **string** |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **interface{}** |  | [optional] 
**UserAgent** | Pointer to **string** |  | [optional] [default to "NULL::character varying"]
**CreatorUserId** | Pointer to **int32** |  | [optional] 
**LocalFolderPath** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [default to "unverified"]
**ReviewedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ReviewedDate** | Pointer to **string** |  | [optional] 
**AmendedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AmendedDate** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**ExternalIntegrationId** | Pointer to **string** |  | [optional] 
**ExternalIntegrationAuthor** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewFiles

`func NewFiles(name string, type_ string, url string, ) *Files`

NewFiles instantiates a new Files object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesWithDefaults

`func NewFilesWithDefaults() *Files`

NewFilesWithDefaults instantiates a new Files object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Files) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Files) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Files) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Files) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Files) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Files) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Files) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Files) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Files) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Files) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Files) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Files) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetFileableId

`func (o *Files) GetFileableId() int32`

GetFileableId returns the FileableId field if non-nil, zero value otherwise.

### GetFileableIdOk

`func (o *Files) GetFileableIdOk() (*int32, bool)`

GetFileableIdOk returns a tuple with the FileableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileableId

`func (o *Files) SetFileableId(v int32)`

SetFileableId sets FileableId field to given value.

### HasFileableId

`func (o *Files) HasFileableId() bool`

HasFileableId returns a boolean if a field has been set.

### GetFileableType

`func (o *Files) GetFileableType() string`

GetFileableType returns the FileableType field if non-nil, zero value otherwise.

### GetFileableTypeOk

`func (o *Files) GetFileableTypeOk() (*string, bool)`

GetFileableTypeOk returns a tuple with the FileableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileableType

`func (o *Files) SetFileableType(v string)`

SetFileableType sets FileableType field to given value.

### HasFileableType

`func (o *Files) HasFileableType() bool`

HasFileableType returns a boolean if a field has been set.

### GetName

`func (o *Files) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Files) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Files) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *Files) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Files) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Files) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *Files) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *Files) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Files) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Files) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *Files) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Files) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Files) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Files) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetUrl

`func (o *Files) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Files) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Files) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUploadUserId

`func (o *Files) GetUploadUserId() int32`

GetUploadUserId returns the UploadUserId field if non-nil, zero value otherwise.

### GetUploadUserIdOk

`func (o *Files) GetUploadUserIdOk() (*int32, bool)`

GetUploadUserIdOk returns a tuple with the UploadUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUserId

`func (o *Files) SetUploadUserId(v int32)`

SetUploadUserId sets UploadUserId field to given value.

### HasUploadUserId

`func (o *Files) HasUploadUserId() bool`

HasUploadUserId returns a boolean if a field has been set.

### GetComments

`func (o *Files) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Files) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Files) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Files) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Files) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Files) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Files) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Files) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetStorageType

`func (o *Files) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *Files) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *Files) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *Files) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetEmbedUrl

`func (o *Files) GetEmbedUrl() string`

GetEmbedUrl returns the EmbedUrl field if non-nil, zero value otherwise.

### GetEmbedUrlOk

`func (o *Files) GetEmbedUrlOk() (*string, bool)`

GetEmbedUrlOk returns a tuple with the EmbedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedUrl

`func (o *Files) SetEmbedUrl(v string)`

SetEmbedUrl sets EmbedUrl field to given value.

### HasEmbedUrl

`func (o *Files) HasEmbedUrl() bool`

HasEmbedUrl returns a boolean if a field has been set.

### GetThumbUrl

`func (o *Files) GetThumbUrl() string`

GetThumbUrl returns the ThumbUrl field if non-nil, zero value otherwise.

### GetThumbUrlOk

`func (o *Files) GetThumbUrlOk() (*string, bool)`

GetThumbUrlOk returns a tuple with the ThumbUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbUrl

`func (o *Files) SetThumbUrl(v string)`

SetThumbUrl sets ThumbUrl field to given value.

### HasThumbUrl

`func (o *Files) HasThumbUrl() bool`

HasThumbUrl returns a boolean if a field has been set.

### GetImageUrl

`func (o *Files) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *Files) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *Files) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *Files) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetMeta

`func (o *Files) GetMeta() interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Files) GetMetaOk() (*interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Files) SetMeta(v interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Files) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *Files) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *Files) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetUserAgent

`func (o *Files) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *Files) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *Files) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *Files) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetCreatorUserId

`func (o *Files) GetCreatorUserId() int32`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *Files) GetCreatorUserIdOk() (*int32, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *Files) SetCreatorUserId(v int32)`

SetCreatorUserId sets CreatorUserId field to given value.

### HasCreatorUserId

`func (o *Files) HasCreatorUserId() bool`

HasCreatorUserId returns a boolean if a field has been set.

### GetLocalFolderPath

`func (o *Files) GetLocalFolderPath() string`

GetLocalFolderPath returns the LocalFolderPath field if non-nil, zero value otherwise.

### GetLocalFolderPathOk

`func (o *Files) GetLocalFolderPathOk() (*string, bool)`

GetLocalFolderPathOk returns a tuple with the LocalFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalFolderPath

`func (o *Files) SetLocalFolderPath(v string)`

SetLocalFolderPath sets LocalFolderPath field to given value.

### HasLocalFolderPath

`func (o *Files) HasLocalFolderPath() bool`

HasLocalFolderPath returns a boolean if a field has been set.

### GetStatus

`func (o *Files) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Files) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Files) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Files) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReviewedByUserId

`func (o *Files) GetReviewedByUserId() int32`

GetReviewedByUserId returns the ReviewedByUserId field if non-nil, zero value otherwise.

### GetReviewedByUserIdOk

`func (o *Files) GetReviewedByUserIdOk() (*int32, bool)`

GetReviewedByUserIdOk returns a tuple with the ReviewedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedByUserId

`func (o *Files) SetReviewedByUserId(v int32)`

SetReviewedByUserId sets ReviewedByUserId field to given value.

### HasReviewedByUserId

`func (o *Files) HasReviewedByUserId() bool`

HasReviewedByUserId returns a boolean if a field has been set.

### GetReviewedDate

`func (o *Files) GetReviewedDate() string`

GetReviewedDate returns the ReviewedDate field if non-nil, zero value otherwise.

### GetReviewedDateOk

`func (o *Files) GetReviewedDateOk() (*string, bool)`

GetReviewedDateOk returns a tuple with the ReviewedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedDate

`func (o *Files) SetReviewedDate(v string)`

SetReviewedDate sets ReviewedDate field to given value.

### HasReviewedDate

`func (o *Files) HasReviewedDate() bool`

HasReviewedDate returns a boolean if a field has been set.

### GetAmendedByUserId

`func (o *Files) GetAmendedByUserId() int32`

GetAmendedByUserId returns the AmendedByUserId field if non-nil, zero value otherwise.

### GetAmendedByUserIdOk

`func (o *Files) GetAmendedByUserIdOk() (*int32, bool)`

GetAmendedByUserIdOk returns a tuple with the AmendedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendedByUserId

`func (o *Files) SetAmendedByUserId(v int32)`

SetAmendedByUserId sets AmendedByUserId field to given value.

### HasAmendedByUserId

`func (o *Files) HasAmendedByUserId() bool`

HasAmendedByUserId returns a boolean if a field has been set.

### GetAmendedDate

`func (o *Files) GetAmendedDate() string`

GetAmendedDate returns the AmendedDate field if non-nil, zero value otherwise.

### GetAmendedDateOk

`func (o *Files) GetAmendedDateOk() (*string, bool)`

GetAmendedDateOk returns a tuple with the AmendedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendedDate

`func (o *Files) SetAmendedDate(v string)`

SetAmendedDate sets AmendedDate field to given value.

### HasAmendedDate

`func (o *Files) HasAmendedDate() bool`

HasAmendedDate returns a boolean if a field has been set.

### GetScope

`func (o *Files) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Files) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Files) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Files) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetExternalIntegrationId

`func (o *Files) GetExternalIntegrationId() string`

GetExternalIntegrationId returns the ExternalIntegrationId field if non-nil, zero value otherwise.

### GetExternalIntegrationIdOk

`func (o *Files) GetExternalIntegrationIdOk() (*string, bool)`

GetExternalIntegrationIdOk returns a tuple with the ExternalIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIntegrationId

`func (o *Files) SetExternalIntegrationId(v string)`

SetExternalIntegrationId sets ExternalIntegrationId field to given value.

### HasExternalIntegrationId

`func (o *Files) HasExternalIntegrationId() bool`

HasExternalIntegrationId returns a boolean if a field has been set.

### GetExternalIntegrationAuthor

`func (o *Files) GetExternalIntegrationAuthor() interface{}`

GetExternalIntegrationAuthor returns the ExternalIntegrationAuthor field if non-nil, zero value otherwise.

### GetExternalIntegrationAuthorOk

`func (o *Files) GetExternalIntegrationAuthorOk() (*interface{}, bool)`

GetExternalIntegrationAuthorOk returns a tuple with the ExternalIntegrationAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIntegrationAuthor

`func (o *Files) SetExternalIntegrationAuthor(v interface{})`

SetExternalIntegrationAuthor sets ExternalIntegrationAuthor field to given value.

### HasExternalIntegrationAuthor

`func (o *Files) HasExternalIntegrationAuthor() bool`

HasExternalIntegrationAuthor returns a boolean if a field has been set.

### SetExternalIntegrationAuthorNil

`func (o *Files) SetExternalIntegrationAuthorNil(b bool)`

 SetExternalIntegrationAuthorNil sets the value for ExternalIntegrationAuthor to be an explicit nil

### UnsetExternalIntegrationAuthor
`func (o *Files) UnsetExternalIntegrationAuthor()`

UnsetExternalIntegrationAuthor ensures that no value is present for ExternalIntegrationAuthor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


