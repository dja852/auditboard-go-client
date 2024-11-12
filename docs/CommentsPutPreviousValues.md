# CommentsPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**ParentCommentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;comments.id&#x60;.&lt;fk table&#x3D;&#39;comments&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OwnerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**MessageRaw** | Pointer to **string** |  | [optional] 
**MessageHtml** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**CommentableId** | Pointer to **int32** |  | [optional] 
**CommentableType** | Pointer to **string** |  | [optional] 
**ReferenceMeta** | Pointer to **interface{}** |  | [optional] 
**OriginAttachmentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;attachments.id&#x60;.&lt;fk table&#x3D;&#39;attachments&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OriginFileId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;files.id&#x60;.&lt;fk table&#x3D;&#39;files&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Scopes** | Pointer to **interface{}** |  | [optional] 
**ExternalIntegrationId** | Pointer to **string** |  | [optional] [default to ""]
**ExternalIntegrationAuthor** | Pointer to **interface{}** |  | [optional] 
**OriginatingPath** | Pointer to **string** |  | [optional] 

## Methods

### NewCommentsPutPreviousValues

`func NewCommentsPutPreviousValues() *CommentsPutPreviousValues`

NewCommentsPutPreviousValues instantiates a new CommentsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentsPutPreviousValuesWithDefaults

`func NewCommentsPutPreviousValuesWithDefaults() *CommentsPutPreviousValues`

NewCommentsPutPreviousValuesWithDefaults instantiates a new CommentsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommentsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommentsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommentsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CommentsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentCommentId

`func (o *CommentsPutPreviousValues) GetParentCommentId() int32`

GetParentCommentId returns the ParentCommentId field if non-nil, zero value otherwise.

### GetParentCommentIdOk

`func (o *CommentsPutPreviousValues) GetParentCommentIdOk() (*int32, bool)`

GetParentCommentIdOk returns a tuple with the ParentCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommentId

`func (o *CommentsPutPreviousValues) SetParentCommentId(v int32)`

SetParentCommentId sets ParentCommentId field to given value.

### HasParentCommentId

`func (o *CommentsPutPreviousValues) HasParentCommentId() bool`

HasParentCommentId returns a boolean if a field has been set.

### GetOwnerUserId

`func (o *CommentsPutPreviousValues) GetOwnerUserId() int32`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *CommentsPutPreviousValues) GetOwnerUserIdOk() (*int32, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *CommentsPutPreviousValues) SetOwnerUserId(v int32)`

SetOwnerUserId sets OwnerUserId field to given value.

### HasOwnerUserId

`func (o *CommentsPutPreviousValues) HasOwnerUserId() bool`

HasOwnerUserId returns a boolean if a field has been set.

### GetType

`func (o *CommentsPutPreviousValues) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommentsPutPreviousValues) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommentsPutPreviousValues) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CommentsPutPreviousValues) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *CommentsPutPreviousValues) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommentsPutPreviousValues) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommentsPutPreviousValues) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommentsPutPreviousValues) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessageRaw

`func (o *CommentsPutPreviousValues) GetMessageRaw() string`

GetMessageRaw returns the MessageRaw field if non-nil, zero value otherwise.

### GetMessageRawOk

`func (o *CommentsPutPreviousValues) GetMessageRawOk() (*string, bool)`

GetMessageRawOk returns a tuple with the MessageRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageRaw

`func (o *CommentsPutPreviousValues) SetMessageRaw(v string)`

SetMessageRaw sets MessageRaw field to given value.

### HasMessageRaw

`func (o *CommentsPutPreviousValues) HasMessageRaw() bool`

HasMessageRaw returns a boolean if a field has been set.

### GetMessageHtml

`func (o *CommentsPutPreviousValues) GetMessageHtml() string`

GetMessageHtml returns the MessageHtml field if non-nil, zero value otherwise.

### GetMessageHtmlOk

`func (o *CommentsPutPreviousValues) GetMessageHtmlOk() (*string, bool)`

GetMessageHtmlOk returns a tuple with the MessageHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageHtml

`func (o *CommentsPutPreviousValues) SetMessageHtml(v string)`

SetMessageHtml sets MessageHtml field to given value.

### HasMessageHtml

`func (o *CommentsPutPreviousValues) HasMessageHtml() bool`

HasMessageHtml returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CommentsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommentsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommentsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommentsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CommentsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CommentsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CommentsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CommentsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *CommentsPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CommentsPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CommentsPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *CommentsPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetCommentableId

`func (o *CommentsPutPreviousValues) GetCommentableId() int32`

GetCommentableId returns the CommentableId field if non-nil, zero value otherwise.

### GetCommentableIdOk

`func (o *CommentsPutPreviousValues) GetCommentableIdOk() (*int32, bool)`

GetCommentableIdOk returns a tuple with the CommentableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentableId

`func (o *CommentsPutPreviousValues) SetCommentableId(v int32)`

SetCommentableId sets CommentableId field to given value.

### HasCommentableId

`func (o *CommentsPutPreviousValues) HasCommentableId() bool`

HasCommentableId returns a boolean if a field has been set.

### GetCommentableType

`func (o *CommentsPutPreviousValues) GetCommentableType() string`

GetCommentableType returns the CommentableType field if non-nil, zero value otherwise.

### GetCommentableTypeOk

`func (o *CommentsPutPreviousValues) GetCommentableTypeOk() (*string, bool)`

GetCommentableTypeOk returns a tuple with the CommentableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentableType

`func (o *CommentsPutPreviousValues) SetCommentableType(v string)`

SetCommentableType sets CommentableType field to given value.

### HasCommentableType

`func (o *CommentsPutPreviousValues) HasCommentableType() bool`

HasCommentableType returns a boolean if a field has been set.

### GetReferenceMeta

`func (o *CommentsPutPreviousValues) GetReferenceMeta() interface{}`

GetReferenceMeta returns the ReferenceMeta field if non-nil, zero value otherwise.

### GetReferenceMetaOk

`func (o *CommentsPutPreviousValues) GetReferenceMetaOk() (*interface{}, bool)`

GetReferenceMetaOk returns a tuple with the ReferenceMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMeta

`func (o *CommentsPutPreviousValues) SetReferenceMeta(v interface{})`

SetReferenceMeta sets ReferenceMeta field to given value.

### HasReferenceMeta

`func (o *CommentsPutPreviousValues) HasReferenceMeta() bool`

HasReferenceMeta returns a boolean if a field has been set.

### SetReferenceMetaNil

`func (o *CommentsPutPreviousValues) SetReferenceMetaNil(b bool)`

 SetReferenceMetaNil sets the value for ReferenceMeta to be an explicit nil

### UnsetReferenceMeta
`func (o *CommentsPutPreviousValues) UnsetReferenceMeta()`

UnsetReferenceMeta ensures that no value is present for ReferenceMeta, not even an explicit nil
### GetOriginAttachmentId

`func (o *CommentsPutPreviousValues) GetOriginAttachmentId() int32`

GetOriginAttachmentId returns the OriginAttachmentId field if non-nil, zero value otherwise.

### GetOriginAttachmentIdOk

`func (o *CommentsPutPreviousValues) GetOriginAttachmentIdOk() (*int32, bool)`

GetOriginAttachmentIdOk returns a tuple with the OriginAttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAttachmentId

`func (o *CommentsPutPreviousValues) SetOriginAttachmentId(v int32)`

SetOriginAttachmentId sets OriginAttachmentId field to given value.

### HasOriginAttachmentId

`func (o *CommentsPutPreviousValues) HasOriginAttachmentId() bool`

HasOriginAttachmentId returns a boolean if a field has been set.

### GetOriginFileId

`func (o *CommentsPutPreviousValues) GetOriginFileId() int32`

GetOriginFileId returns the OriginFileId field if non-nil, zero value otherwise.

### GetOriginFileIdOk

`func (o *CommentsPutPreviousValues) GetOriginFileIdOk() (*int32, bool)`

GetOriginFileIdOk returns a tuple with the OriginFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFileId

`func (o *CommentsPutPreviousValues) SetOriginFileId(v int32)`

SetOriginFileId sets OriginFileId field to given value.

### HasOriginFileId

`func (o *CommentsPutPreviousValues) HasOriginFileId() bool`

HasOriginFileId returns a boolean if a field has been set.

### GetScopes

`func (o *CommentsPutPreviousValues) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CommentsPutPreviousValues) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CommentsPutPreviousValues) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *CommentsPutPreviousValues) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *CommentsPutPreviousValues) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *CommentsPutPreviousValues) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetExternalIntegrationId

`func (o *CommentsPutPreviousValues) GetExternalIntegrationId() string`

GetExternalIntegrationId returns the ExternalIntegrationId field if non-nil, zero value otherwise.

### GetExternalIntegrationIdOk

`func (o *CommentsPutPreviousValues) GetExternalIntegrationIdOk() (*string, bool)`

GetExternalIntegrationIdOk returns a tuple with the ExternalIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIntegrationId

`func (o *CommentsPutPreviousValues) SetExternalIntegrationId(v string)`

SetExternalIntegrationId sets ExternalIntegrationId field to given value.

### HasExternalIntegrationId

`func (o *CommentsPutPreviousValues) HasExternalIntegrationId() bool`

HasExternalIntegrationId returns a boolean if a field has been set.

### GetExternalIntegrationAuthor

`func (o *CommentsPutPreviousValues) GetExternalIntegrationAuthor() interface{}`

GetExternalIntegrationAuthor returns the ExternalIntegrationAuthor field if non-nil, zero value otherwise.

### GetExternalIntegrationAuthorOk

`func (o *CommentsPutPreviousValues) GetExternalIntegrationAuthorOk() (*interface{}, bool)`

GetExternalIntegrationAuthorOk returns a tuple with the ExternalIntegrationAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIntegrationAuthor

`func (o *CommentsPutPreviousValues) SetExternalIntegrationAuthor(v interface{})`

SetExternalIntegrationAuthor sets ExternalIntegrationAuthor field to given value.

### HasExternalIntegrationAuthor

`func (o *CommentsPutPreviousValues) HasExternalIntegrationAuthor() bool`

HasExternalIntegrationAuthor returns a boolean if a field has been set.

### SetExternalIntegrationAuthorNil

`func (o *CommentsPutPreviousValues) SetExternalIntegrationAuthorNil(b bool)`

 SetExternalIntegrationAuthorNil sets the value for ExternalIntegrationAuthor to be an explicit nil

### UnsetExternalIntegrationAuthor
`func (o *CommentsPutPreviousValues) UnsetExternalIntegrationAuthor()`

UnsetExternalIntegrationAuthor ensures that no value is present for ExternalIntegrationAuthor, not even an explicit nil
### GetOriginatingPath

`func (o *CommentsPutPreviousValues) GetOriginatingPath() string`

GetOriginatingPath returns the OriginatingPath field if non-nil, zero value otherwise.

### GetOriginatingPathOk

`func (o *CommentsPutPreviousValues) GetOriginatingPathOk() (*string, bool)`

GetOriginatingPathOk returns a tuple with the OriginatingPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingPath

`func (o *CommentsPutPreviousValues) SetOriginatingPath(v string)`

SetOriginatingPath sets OriginatingPath field to given value.

### HasOriginatingPath

`func (o *CommentsPutPreviousValues) HasOriginatingPath() bool`

HasOriginatingPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


