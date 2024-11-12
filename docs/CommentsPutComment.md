# CommentsPutComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**ParentCommentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;comments.id&#x60;.&lt;fk table&#x3D;&#39;comments&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OwnerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**MessageRaw** | **string** |  | 
**MessageHtml** | **string** |  | 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**CommentableId** | Pointer to **int32** |  | [optional] 
**CommentableType** | Pointer to **string** |  | [optional] 
**ReferenceMeta** | Pointer to **interface{}** |  | [optional] 
**OriginAttachmentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;attachments.id&#x60;.&lt;fk table&#x3D;&#39;attachments&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OriginFileId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;files.id&#x60;.&lt;fk table&#x3D;&#39;files&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Scopes** | **interface{}** |  | 
**ExternalIntegrationId** | **string** |  | [default to ""]
**ExternalIntegrationAuthor** | Pointer to **interface{}** |  | [optional] 
**OriginatingPath** | Pointer to **string** |  | [optional] 

## Methods

### NewCommentsPutComment

`func NewCommentsPutComment(messageRaw string, messageHtml string, scopes interface{}, externalIntegrationId string, ) *CommentsPutComment`

NewCommentsPutComment instantiates a new CommentsPutComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentsPutCommentWithDefaults

`func NewCommentsPutCommentWithDefaults() *CommentsPutComment`

NewCommentsPutCommentWithDefaults instantiates a new CommentsPutComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommentsPutComment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommentsPutComment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommentsPutComment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CommentsPutComment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentCommentId

`func (o *CommentsPutComment) GetParentCommentId() int32`

GetParentCommentId returns the ParentCommentId field if non-nil, zero value otherwise.

### GetParentCommentIdOk

`func (o *CommentsPutComment) GetParentCommentIdOk() (*int32, bool)`

GetParentCommentIdOk returns a tuple with the ParentCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommentId

`func (o *CommentsPutComment) SetParentCommentId(v int32)`

SetParentCommentId sets ParentCommentId field to given value.

### HasParentCommentId

`func (o *CommentsPutComment) HasParentCommentId() bool`

HasParentCommentId returns a boolean if a field has been set.

### GetOwnerUserId

`func (o *CommentsPutComment) GetOwnerUserId() int32`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *CommentsPutComment) GetOwnerUserIdOk() (*int32, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *CommentsPutComment) SetOwnerUserId(v int32)`

SetOwnerUserId sets OwnerUserId field to given value.

### HasOwnerUserId

`func (o *CommentsPutComment) HasOwnerUserId() bool`

HasOwnerUserId returns a boolean if a field has been set.

### GetType

`func (o *CommentsPutComment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommentsPutComment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommentsPutComment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CommentsPutComment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *CommentsPutComment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommentsPutComment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommentsPutComment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommentsPutComment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessageRaw

`func (o *CommentsPutComment) GetMessageRaw() string`

GetMessageRaw returns the MessageRaw field if non-nil, zero value otherwise.

### GetMessageRawOk

`func (o *CommentsPutComment) GetMessageRawOk() (*string, bool)`

GetMessageRawOk returns a tuple with the MessageRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageRaw

`func (o *CommentsPutComment) SetMessageRaw(v string)`

SetMessageRaw sets MessageRaw field to given value.


### GetMessageHtml

`func (o *CommentsPutComment) GetMessageHtml() string`

GetMessageHtml returns the MessageHtml field if non-nil, zero value otherwise.

### GetMessageHtmlOk

`func (o *CommentsPutComment) GetMessageHtmlOk() (*string, bool)`

GetMessageHtmlOk returns a tuple with the MessageHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageHtml

`func (o *CommentsPutComment) SetMessageHtml(v string)`

SetMessageHtml sets MessageHtml field to given value.


### GetCreatedAt

`func (o *CommentsPutComment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommentsPutComment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommentsPutComment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommentsPutComment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CommentsPutComment) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CommentsPutComment) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CommentsPutComment) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CommentsPutComment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *CommentsPutComment) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CommentsPutComment) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CommentsPutComment) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *CommentsPutComment) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetCommentableId

`func (o *CommentsPutComment) GetCommentableId() int32`

GetCommentableId returns the CommentableId field if non-nil, zero value otherwise.

### GetCommentableIdOk

`func (o *CommentsPutComment) GetCommentableIdOk() (*int32, bool)`

GetCommentableIdOk returns a tuple with the CommentableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentableId

`func (o *CommentsPutComment) SetCommentableId(v int32)`

SetCommentableId sets CommentableId field to given value.

### HasCommentableId

`func (o *CommentsPutComment) HasCommentableId() bool`

HasCommentableId returns a boolean if a field has been set.

### GetCommentableType

`func (o *CommentsPutComment) GetCommentableType() string`

GetCommentableType returns the CommentableType field if non-nil, zero value otherwise.

### GetCommentableTypeOk

`func (o *CommentsPutComment) GetCommentableTypeOk() (*string, bool)`

GetCommentableTypeOk returns a tuple with the CommentableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentableType

`func (o *CommentsPutComment) SetCommentableType(v string)`

SetCommentableType sets CommentableType field to given value.

### HasCommentableType

`func (o *CommentsPutComment) HasCommentableType() bool`

HasCommentableType returns a boolean if a field has been set.

### GetReferenceMeta

`func (o *CommentsPutComment) GetReferenceMeta() interface{}`

GetReferenceMeta returns the ReferenceMeta field if non-nil, zero value otherwise.

### GetReferenceMetaOk

`func (o *CommentsPutComment) GetReferenceMetaOk() (*interface{}, bool)`

GetReferenceMetaOk returns a tuple with the ReferenceMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMeta

`func (o *CommentsPutComment) SetReferenceMeta(v interface{})`

SetReferenceMeta sets ReferenceMeta field to given value.

### HasReferenceMeta

`func (o *CommentsPutComment) HasReferenceMeta() bool`

HasReferenceMeta returns a boolean if a field has been set.

### SetReferenceMetaNil

`func (o *CommentsPutComment) SetReferenceMetaNil(b bool)`

 SetReferenceMetaNil sets the value for ReferenceMeta to be an explicit nil

### UnsetReferenceMeta
`func (o *CommentsPutComment) UnsetReferenceMeta()`

UnsetReferenceMeta ensures that no value is present for ReferenceMeta, not even an explicit nil
### GetOriginAttachmentId

`func (o *CommentsPutComment) GetOriginAttachmentId() int32`

GetOriginAttachmentId returns the OriginAttachmentId field if non-nil, zero value otherwise.

### GetOriginAttachmentIdOk

`func (o *CommentsPutComment) GetOriginAttachmentIdOk() (*int32, bool)`

GetOriginAttachmentIdOk returns a tuple with the OriginAttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAttachmentId

`func (o *CommentsPutComment) SetOriginAttachmentId(v int32)`

SetOriginAttachmentId sets OriginAttachmentId field to given value.

### HasOriginAttachmentId

`func (o *CommentsPutComment) HasOriginAttachmentId() bool`

HasOriginAttachmentId returns a boolean if a field has been set.

### GetOriginFileId

`func (o *CommentsPutComment) GetOriginFileId() int32`

GetOriginFileId returns the OriginFileId field if non-nil, zero value otherwise.

### GetOriginFileIdOk

`func (o *CommentsPutComment) GetOriginFileIdOk() (*int32, bool)`

GetOriginFileIdOk returns a tuple with the OriginFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFileId

`func (o *CommentsPutComment) SetOriginFileId(v int32)`

SetOriginFileId sets OriginFileId field to given value.

### HasOriginFileId

`func (o *CommentsPutComment) HasOriginFileId() bool`

HasOriginFileId returns a boolean if a field has been set.

### GetScopes

`func (o *CommentsPutComment) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CommentsPutComment) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CommentsPutComment) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.


### SetScopesNil

`func (o *CommentsPutComment) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *CommentsPutComment) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetExternalIntegrationId

`func (o *CommentsPutComment) GetExternalIntegrationId() string`

GetExternalIntegrationId returns the ExternalIntegrationId field if non-nil, zero value otherwise.

### GetExternalIntegrationIdOk

`func (o *CommentsPutComment) GetExternalIntegrationIdOk() (*string, bool)`

GetExternalIntegrationIdOk returns a tuple with the ExternalIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIntegrationId

`func (o *CommentsPutComment) SetExternalIntegrationId(v string)`

SetExternalIntegrationId sets ExternalIntegrationId field to given value.


### GetExternalIntegrationAuthor

`func (o *CommentsPutComment) GetExternalIntegrationAuthor() interface{}`

GetExternalIntegrationAuthor returns the ExternalIntegrationAuthor field if non-nil, zero value otherwise.

### GetExternalIntegrationAuthorOk

`func (o *CommentsPutComment) GetExternalIntegrationAuthorOk() (*interface{}, bool)`

GetExternalIntegrationAuthorOk returns a tuple with the ExternalIntegrationAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIntegrationAuthor

`func (o *CommentsPutComment) SetExternalIntegrationAuthor(v interface{})`

SetExternalIntegrationAuthor sets ExternalIntegrationAuthor field to given value.

### HasExternalIntegrationAuthor

`func (o *CommentsPutComment) HasExternalIntegrationAuthor() bool`

HasExternalIntegrationAuthor returns a boolean if a field has been set.

### SetExternalIntegrationAuthorNil

`func (o *CommentsPutComment) SetExternalIntegrationAuthorNil(b bool)`

 SetExternalIntegrationAuthorNil sets the value for ExternalIntegrationAuthor to be an explicit nil

### UnsetExternalIntegrationAuthor
`func (o *CommentsPutComment) UnsetExternalIntegrationAuthor()`

UnsetExternalIntegrationAuthor ensures that no value is present for ExternalIntegrationAuthor, not even an explicit nil
### GetOriginatingPath

`func (o *CommentsPutComment) GetOriginatingPath() string`

GetOriginatingPath returns the OriginatingPath field if non-nil, zero value otherwise.

### GetOriginatingPathOk

`func (o *CommentsPutComment) GetOriginatingPathOk() (*string, bool)`

GetOriginatingPathOk returns a tuple with the OriginatingPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingPath

`func (o *CommentsPutComment) SetOriginatingPath(v string)`

SetOriginatingPath sets OriginatingPath field to given value.

### HasOriginatingPath

`func (o *CommentsPutComment) HasOriginatingPath() bool`

HasOriginatingPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


