# FilesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**FileableId** | Pointer to **int32** |  | [optional] 
**FileableType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
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

### NewFilesPutPreviousValues

`func NewFilesPutPreviousValues() *FilesPutPreviousValues`

NewFilesPutPreviousValues instantiates a new FilesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesPutPreviousValuesWithDefaults

`func NewFilesPutPreviousValuesWithDefaults() *FilesPutPreviousValues`

NewFilesPutPreviousValuesWithDefaults instantiates a new FilesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FilesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FilesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FilesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FilesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FilesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FilesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FilesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FilesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FilesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FilesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FilesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FilesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetFileableId

`func (o *FilesPutPreviousValues) GetFileableId() int32`

GetFileableId returns the FileableId field if non-nil, zero value otherwise.

### GetFileableIdOk

`func (o *FilesPutPreviousValues) GetFileableIdOk() (*int32, bool)`

GetFileableIdOk returns a tuple with the FileableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileableId

`func (o *FilesPutPreviousValues) SetFileableId(v int32)`

SetFileableId sets FileableId field to given value.

### HasFileableId

`func (o *FilesPutPreviousValues) HasFileableId() bool`

HasFileableId returns a boolean if a field has been set.

### GetFileableType

`func (o *FilesPutPreviousValues) GetFileableType() string`

GetFileableType returns the FileableType field if non-nil, zero value otherwise.

### GetFileableTypeOk

`func (o *FilesPutPreviousValues) GetFileableTypeOk() (*string, bool)`

GetFileableTypeOk returns a tuple with the FileableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileableType

`func (o *FilesPutPreviousValues) SetFileableType(v string)`

SetFileableType sets FileableType field to given value.

### HasFileableType

`func (o *FilesPutPreviousValues) HasFileableType() bool`

HasFileableType returns a boolean if a field has been set.

### GetName

`func (o *FilesPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FilesPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FilesPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FilesPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *FilesPutPreviousValues) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FilesPutPreviousValues) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FilesPutPreviousValues) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *FilesPutPreviousValues) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *FilesPutPreviousValues) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FilesPutPreviousValues) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FilesPutPreviousValues) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FilesPutPreviousValues) HasType() bool`

HasType returns a boolean if a field has been set.

### GetKey

`func (o *FilesPutPreviousValues) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FilesPutPreviousValues) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FilesPutPreviousValues) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FilesPutPreviousValues) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetUrl

`func (o *FilesPutPreviousValues) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FilesPutPreviousValues) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FilesPutPreviousValues) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FilesPutPreviousValues) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUploadUserId

`func (o *FilesPutPreviousValues) GetUploadUserId() int32`

GetUploadUserId returns the UploadUserId field if non-nil, zero value otherwise.

### GetUploadUserIdOk

`func (o *FilesPutPreviousValues) GetUploadUserIdOk() (*int32, bool)`

GetUploadUserIdOk returns a tuple with the UploadUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUserId

`func (o *FilesPutPreviousValues) SetUploadUserId(v int32)`

SetUploadUserId sets UploadUserId field to given value.

### HasUploadUserId

`func (o *FilesPutPreviousValues) HasUploadUserId() bool`

HasUploadUserId returns a boolean if a field has been set.

### GetComments

`func (o *FilesPutPreviousValues) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *FilesPutPreviousValues) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *FilesPutPreviousValues) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *FilesPutPreviousValues) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetDeletedAt

`func (o *FilesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *FilesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *FilesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *FilesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetStorageType

`func (o *FilesPutPreviousValues) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *FilesPutPreviousValues) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *FilesPutPreviousValues) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *FilesPutPreviousValues) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetEmbedUrl

`func (o *FilesPutPreviousValues) GetEmbedUrl() string`

GetEmbedUrl returns the EmbedUrl field if non-nil, zero value otherwise.

### GetEmbedUrlOk

`func (o *FilesPutPreviousValues) GetEmbedUrlOk() (*string, bool)`

GetEmbedUrlOk returns a tuple with the EmbedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedUrl

`func (o *FilesPutPreviousValues) SetEmbedUrl(v string)`

SetEmbedUrl sets EmbedUrl field to given value.

### HasEmbedUrl

`func (o *FilesPutPreviousValues) HasEmbedUrl() bool`

HasEmbedUrl returns a boolean if a field has been set.

### GetThumbUrl

`func (o *FilesPutPreviousValues) GetThumbUrl() string`

GetThumbUrl returns the ThumbUrl field if non-nil, zero value otherwise.

### GetThumbUrlOk

`func (o *FilesPutPreviousValues) GetThumbUrlOk() (*string, bool)`

GetThumbUrlOk returns a tuple with the ThumbUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbUrl

`func (o *FilesPutPreviousValues) SetThumbUrl(v string)`

SetThumbUrl sets ThumbUrl field to given value.

### HasThumbUrl

`func (o *FilesPutPreviousValues) HasThumbUrl() bool`

HasThumbUrl returns a boolean if a field has been set.

### GetImageUrl

`func (o *FilesPutPreviousValues) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *FilesPutPreviousValues) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *FilesPutPreviousValues) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *FilesPutPreviousValues) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetMeta

`func (o *FilesPutPreviousValues) GetMeta() interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FilesPutPreviousValues) GetMetaOk() (*interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FilesPutPreviousValues) SetMeta(v interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FilesPutPreviousValues) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *FilesPutPreviousValues) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *FilesPutPreviousValues) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetUserAgent

`func (o *FilesPutPreviousValues) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *FilesPutPreviousValues) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *FilesPutPreviousValues) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *FilesPutPreviousValues) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetCreatorUserId

`func (o *FilesPutPreviousValues) GetCreatorUserId() int32`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *FilesPutPreviousValues) GetCreatorUserIdOk() (*int32, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *FilesPutPreviousValues) SetCreatorUserId(v int32)`

SetCreatorUserId sets CreatorUserId field to given value.

### HasCreatorUserId

`func (o *FilesPutPreviousValues) HasCreatorUserId() bool`

HasCreatorUserId returns a boolean if a field has been set.

### GetLocalFolderPath

`func (o *FilesPutPreviousValues) GetLocalFolderPath() string`

GetLocalFolderPath returns the LocalFolderPath field if non-nil, zero value otherwise.

### GetLocalFolderPathOk

`func (o *FilesPutPreviousValues) GetLocalFolderPathOk() (*string, bool)`

GetLocalFolderPathOk returns a tuple with the LocalFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalFolderPath

`func (o *FilesPutPreviousValues) SetLocalFolderPath(v string)`

SetLocalFolderPath sets LocalFolderPath field to given value.

### HasLocalFolderPath

`func (o *FilesPutPreviousValues) HasLocalFolderPath() bool`

HasLocalFolderPath returns a boolean if a field has been set.

### GetStatus

`func (o *FilesPutPreviousValues) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FilesPutPreviousValues) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FilesPutPreviousValues) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FilesPutPreviousValues) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReviewedByUserId

`func (o *FilesPutPreviousValues) GetReviewedByUserId() int32`

GetReviewedByUserId returns the ReviewedByUserId field if non-nil, zero value otherwise.

### GetReviewedByUserIdOk

`func (o *FilesPutPreviousValues) GetReviewedByUserIdOk() (*int32, bool)`

GetReviewedByUserIdOk returns a tuple with the ReviewedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedByUserId

`func (o *FilesPutPreviousValues) SetReviewedByUserId(v int32)`

SetReviewedByUserId sets ReviewedByUserId field to given value.

### HasReviewedByUserId

`func (o *FilesPutPreviousValues) HasReviewedByUserId() bool`

HasReviewedByUserId returns a boolean if a field has been set.

### GetReviewedDate

`func (o *FilesPutPreviousValues) GetReviewedDate() string`

GetReviewedDate returns the ReviewedDate field if non-nil, zero value otherwise.

### GetReviewedDateOk

`func (o *FilesPutPreviousValues) GetReviewedDateOk() (*string, bool)`

GetReviewedDateOk returns a tuple with the ReviewedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedDate

`func (o *FilesPutPreviousValues) SetReviewedDate(v string)`

SetReviewedDate sets ReviewedDate field to given value.

### HasReviewedDate

`func (o *FilesPutPreviousValues) HasReviewedDate() bool`

HasReviewedDate returns a boolean if a field has been set.

### GetAmendedByUserId

`func (o *FilesPutPreviousValues) GetAmendedByUserId() int32`

GetAmendedByUserId returns the AmendedByUserId field if non-nil, zero value otherwise.

### GetAmendedByUserIdOk

`func (o *FilesPutPreviousValues) GetAmendedByUserIdOk() (*int32, bool)`

GetAmendedByUserIdOk returns a tuple with the AmendedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendedByUserId

`func (o *FilesPutPreviousValues) SetAmendedByUserId(v int32)`

SetAmendedByUserId sets AmendedByUserId field to given value.

### HasAmendedByUserId

`func (o *FilesPutPreviousValues) HasAmendedByUserId() bool`

HasAmendedByUserId returns a boolean if a field has been set.

### GetAmendedDate

`func (o *FilesPutPreviousValues) GetAmendedDate() string`

GetAmendedDate returns the AmendedDate field if non-nil, zero value otherwise.

### GetAmendedDateOk

`func (o *FilesPutPreviousValues) GetAmendedDateOk() (*string, bool)`

GetAmendedDateOk returns a tuple with the AmendedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendedDate

`func (o *FilesPutPreviousValues) SetAmendedDate(v string)`

SetAmendedDate sets AmendedDate field to given value.

### HasAmendedDate

`func (o *FilesPutPreviousValues) HasAmendedDate() bool`

HasAmendedDate returns a boolean if a field has been set.

### GetScope

`func (o *FilesPutPreviousValues) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *FilesPutPreviousValues) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *FilesPutPreviousValues) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *FilesPutPreviousValues) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetExternalIntegrationId

`func (o *FilesPutPreviousValues) GetExternalIntegrationId() string`

GetExternalIntegrationId returns the ExternalIntegrationId field if non-nil, zero value otherwise.

### GetExternalIntegrationIdOk

`func (o *FilesPutPreviousValues) GetExternalIntegrationIdOk() (*string, bool)`

GetExternalIntegrationIdOk returns a tuple with the ExternalIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIntegrationId

`func (o *FilesPutPreviousValues) SetExternalIntegrationId(v string)`

SetExternalIntegrationId sets ExternalIntegrationId field to given value.

### HasExternalIntegrationId

`func (o *FilesPutPreviousValues) HasExternalIntegrationId() bool`

HasExternalIntegrationId returns a boolean if a field has been set.

### GetExternalIntegrationAuthor

`func (o *FilesPutPreviousValues) GetExternalIntegrationAuthor() interface{}`

GetExternalIntegrationAuthor returns the ExternalIntegrationAuthor field if non-nil, zero value otherwise.

### GetExternalIntegrationAuthorOk

`func (o *FilesPutPreviousValues) GetExternalIntegrationAuthorOk() (*interface{}, bool)`

GetExternalIntegrationAuthorOk returns a tuple with the ExternalIntegrationAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIntegrationAuthor

`func (o *FilesPutPreviousValues) SetExternalIntegrationAuthor(v interface{})`

SetExternalIntegrationAuthor sets ExternalIntegrationAuthor field to given value.

### HasExternalIntegrationAuthor

`func (o *FilesPutPreviousValues) HasExternalIntegrationAuthor() bool`

HasExternalIntegrationAuthor returns a boolean if a field has been set.

### SetExternalIntegrationAuthorNil

`func (o *FilesPutPreviousValues) SetExternalIntegrationAuthorNil(b bool)`

 SetExternalIntegrationAuthorNil sets the value for ExternalIntegrationAuthor to be an explicit nil

### UnsetExternalIntegrationAuthor
`func (o *FilesPutPreviousValues) UnsetExternalIntegrationAuthor()`

UnsetExternalIntegrationAuthor ensures that no value is present for ExternalIntegrationAuthor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


