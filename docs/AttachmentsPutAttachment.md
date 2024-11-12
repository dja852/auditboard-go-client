# AttachmentsPutAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Type** | **string** |  | 
**SheetId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;sheets.id&#x60;.&lt;fk table&#x3D;&#39;sheets&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**DocumentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;documents.id&#x60;.&lt;fk table&#x3D;&#39;documents&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AttachableId** | Pointer to **int32** |  | [optional] 
**AttachableType** | Pointer to **string** |  | [optional] 
**LockUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**LockedAt** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**FileId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;files.id&#x60;.&lt;fk table&#x3D;&#39;files&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**SortOrder** | Pointer to **int32** |  | [optional] 
**LinkifySourceFileId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;files.id&#x60;.&lt;fk table&#x3D;&#39;files&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IsLinkified** | Pointer to **bool** |  | [optional] [default to false]
**LastLockUserId** | Pointer to **int32** |  | [optional] 
**WordReportOptions** | Pointer to **interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [default to "unverified"]
**ReviewedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ReviewedDate** | Pointer to **string** |  | [optional] 
**AmendedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AmendedDate** | Pointer to **string** |  | [optional] 
**BrokenLinkCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewAttachmentsPutAttachment

`func NewAttachmentsPutAttachment(type_ string, ) *AttachmentsPutAttachment`

NewAttachmentsPutAttachment instantiates a new AttachmentsPutAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentsPutAttachmentWithDefaults

`func NewAttachmentsPutAttachmentWithDefaults() *AttachmentsPutAttachment`

NewAttachmentsPutAttachmentWithDefaults instantiates a new AttachmentsPutAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AttachmentsPutAttachment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttachmentsPutAttachment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttachmentsPutAttachment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AttachmentsPutAttachment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AttachmentsPutAttachment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AttachmentsPutAttachment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AttachmentsPutAttachment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AttachmentsPutAttachment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AttachmentsPutAttachment) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AttachmentsPutAttachment) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AttachmentsPutAttachment) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AttachmentsPutAttachment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AttachmentsPutAttachment) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AttachmentsPutAttachment) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AttachmentsPutAttachment) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AttachmentsPutAttachment) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetType

`func (o *AttachmentsPutAttachment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttachmentsPutAttachment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttachmentsPutAttachment) SetType(v string)`

SetType sets Type field to given value.


### GetSheetId

`func (o *AttachmentsPutAttachment) GetSheetId() int32`

GetSheetId returns the SheetId field if non-nil, zero value otherwise.

### GetSheetIdOk

`func (o *AttachmentsPutAttachment) GetSheetIdOk() (*int32, bool)`

GetSheetIdOk returns a tuple with the SheetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheetId

`func (o *AttachmentsPutAttachment) SetSheetId(v int32)`

SetSheetId sets SheetId field to given value.

### HasSheetId

`func (o *AttachmentsPutAttachment) HasSheetId() bool`

HasSheetId returns a boolean if a field has been set.

### GetDocumentId

`func (o *AttachmentsPutAttachment) GetDocumentId() int32`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *AttachmentsPutAttachment) GetDocumentIdOk() (*int32, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *AttachmentsPutAttachment) SetDocumentId(v int32)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *AttachmentsPutAttachment) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetAttachableId

`func (o *AttachmentsPutAttachment) GetAttachableId() int32`

GetAttachableId returns the AttachableId field if non-nil, zero value otherwise.

### GetAttachableIdOk

`func (o *AttachmentsPutAttachment) GetAttachableIdOk() (*int32, bool)`

GetAttachableIdOk returns a tuple with the AttachableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachableId

`func (o *AttachmentsPutAttachment) SetAttachableId(v int32)`

SetAttachableId sets AttachableId field to given value.

### HasAttachableId

`func (o *AttachmentsPutAttachment) HasAttachableId() bool`

HasAttachableId returns a boolean if a field has been set.

### GetAttachableType

`func (o *AttachmentsPutAttachment) GetAttachableType() string`

GetAttachableType returns the AttachableType field if non-nil, zero value otherwise.

### GetAttachableTypeOk

`func (o *AttachmentsPutAttachment) GetAttachableTypeOk() (*string, bool)`

GetAttachableTypeOk returns a tuple with the AttachableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachableType

`func (o *AttachmentsPutAttachment) SetAttachableType(v string)`

SetAttachableType sets AttachableType field to given value.

### HasAttachableType

`func (o *AttachmentsPutAttachment) HasAttachableType() bool`

HasAttachableType returns a boolean if a field has been set.

### GetLockUserId

`func (o *AttachmentsPutAttachment) GetLockUserId() int32`

GetLockUserId returns the LockUserId field if non-nil, zero value otherwise.

### GetLockUserIdOk

`func (o *AttachmentsPutAttachment) GetLockUserIdOk() (*int32, bool)`

GetLockUserIdOk returns a tuple with the LockUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockUserId

`func (o *AttachmentsPutAttachment) SetLockUserId(v int32)`

SetLockUserId sets LockUserId field to given value.

### HasLockUserId

`func (o *AttachmentsPutAttachment) HasLockUserId() bool`

HasLockUserId returns a boolean if a field has been set.

### GetLockedAt

`func (o *AttachmentsPutAttachment) GetLockedAt() string`

GetLockedAt returns the LockedAt field if non-nil, zero value otherwise.

### GetLockedAtOk

`func (o *AttachmentsPutAttachment) GetLockedAtOk() (*string, bool)`

GetLockedAtOk returns a tuple with the LockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedAt

`func (o *AttachmentsPutAttachment) SetLockedAt(v string)`

SetLockedAt sets LockedAt field to given value.

### HasLockedAt

`func (o *AttachmentsPutAttachment) HasLockedAt() bool`

HasLockedAt returns a boolean if a field has been set.

### GetName

`func (o *AttachmentsPutAttachment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttachmentsPutAttachment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttachmentsPutAttachment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AttachmentsPutAttachment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFileId

`func (o *AttachmentsPutAttachment) GetFileId() int32`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *AttachmentsPutAttachment) GetFileIdOk() (*int32, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *AttachmentsPutAttachment) SetFileId(v int32)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *AttachmentsPutAttachment) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetSortOrder

`func (o *AttachmentsPutAttachment) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AttachmentsPutAttachment) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AttachmentsPutAttachment) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *AttachmentsPutAttachment) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetLinkifySourceFileId

`func (o *AttachmentsPutAttachment) GetLinkifySourceFileId() int32`

GetLinkifySourceFileId returns the LinkifySourceFileId field if non-nil, zero value otherwise.

### GetLinkifySourceFileIdOk

`func (o *AttachmentsPutAttachment) GetLinkifySourceFileIdOk() (*int32, bool)`

GetLinkifySourceFileIdOk returns a tuple with the LinkifySourceFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkifySourceFileId

`func (o *AttachmentsPutAttachment) SetLinkifySourceFileId(v int32)`

SetLinkifySourceFileId sets LinkifySourceFileId field to given value.

### HasLinkifySourceFileId

`func (o *AttachmentsPutAttachment) HasLinkifySourceFileId() bool`

HasLinkifySourceFileId returns a boolean if a field has been set.

### GetIsLinkified

`func (o *AttachmentsPutAttachment) GetIsLinkified() bool`

GetIsLinkified returns the IsLinkified field if non-nil, zero value otherwise.

### GetIsLinkifiedOk

`func (o *AttachmentsPutAttachment) GetIsLinkifiedOk() (*bool, bool)`

GetIsLinkifiedOk returns a tuple with the IsLinkified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLinkified

`func (o *AttachmentsPutAttachment) SetIsLinkified(v bool)`

SetIsLinkified sets IsLinkified field to given value.

### HasIsLinkified

`func (o *AttachmentsPutAttachment) HasIsLinkified() bool`

HasIsLinkified returns a boolean if a field has been set.

### GetLastLockUserId

`func (o *AttachmentsPutAttachment) GetLastLockUserId() int32`

GetLastLockUserId returns the LastLockUserId field if non-nil, zero value otherwise.

### GetLastLockUserIdOk

`func (o *AttachmentsPutAttachment) GetLastLockUserIdOk() (*int32, bool)`

GetLastLockUserIdOk returns a tuple with the LastLockUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLockUserId

`func (o *AttachmentsPutAttachment) SetLastLockUserId(v int32)`

SetLastLockUserId sets LastLockUserId field to given value.

### HasLastLockUserId

`func (o *AttachmentsPutAttachment) HasLastLockUserId() bool`

HasLastLockUserId returns a boolean if a field has been set.

### GetWordReportOptions

`func (o *AttachmentsPutAttachment) GetWordReportOptions() interface{}`

GetWordReportOptions returns the WordReportOptions field if non-nil, zero value otherwise.

### GetWordReportOptionsOk

`func (o *AttachmentsPutAttachment) GetWordReportOptionsOk() (*interface{}, bool)`

GetWordReportOptionsOk returns a tuple with the WordReportOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordReportOptions

`func (o *AttachmentsPutAttachment) SetWordReportOptions(v interface{})`

SetWordReportOptions sets WordReportOptions field to given value.

### HasWordReportOptions

`func (o *AttachmentsPutAttachment) HasWordReportOptions() bool`

HasWordReportOptions returns a boolean if a field has been set.

### SetWordReportOptionsNil

`func (o *AttachmentsPutAttachment) SetWordReportOptionsNil(b bool)`

 SetWordReportOptionsNil sets the value for WordReportOptions to be an explicit nil

### UnsetWordReportOptions
`func (o *AttachmentsPutAttachment) UnsetWordReportOptions()`

UnsetWordReportOptions ensures that no value is present for WordReportOptions, not even an explicit nil
### GetStatus

`func (o *AttachmentsPutAttachment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AttachmentsPutAttachment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AttachmentsPutAttachment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AttachmentsPutAttachment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReviewedByUserId

`func (o *AttachmentsPutAttachment) GetReviewedByUserId() int32`

GetReviewedByUserId returns the ReviewedByUserId field if non-nil, zero value otherwise.

### GetReviewedByUserIdOk

`func (o *AttachmentsPutAttachment) GetReviewedByUserIdOk() (*int32, bool)`

GetReviewedByUserIdOk returns a tuple with the ReviewedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedByUserId

`func (o *AttachmentsPutAttachment) SetReviewedByUserId(v int32)`

SetReviewedByUserId sets ReviewedByUserId field to given value.

### HasReviewedByUserId

`func (o *AttachmentsPutAttachment) HasReviewedByUserId() bool`

HasReviewedByUserId returns a boolean if a field has been set.

### GetReviewedDate

`func (o *AttachmentsPutAttachment) GetReviewedDate() string`

GetReviewedDate returns the ReviewedDate field if non-nil, zero value otherwise.

### GetReviewedDateOk

`func (o *AttachmentsPutAttachment) GetReviewedDateOk() (*string, bool)`

GetReviewedDateOk returns a tuple with the ReviewedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedDate

`func (o *AttachmentsPutAttachment) SetReviewedDate(v string)`

SetReviewedDate sets ReviewedDate field to given value.

### HasReviewedDate

`func (o *AttachmentsPutAttachment) HasReviewedDate() bool`

HasReviewedDate returns a boolean if a field has been set.

### GetAmendedByUserId

`func (o *AttachmentsPutAttachment) GetAmendedByUserId() int32`

GetAmendedByUserId returns the AmendedByUserId field if non-nil, zero value otherwise.

### GetAmendedByUserIdOk

`func (o *AttachmentsPutAttachment) GetAmendedByUserIdOk() (*int32, bool)`

GetAmendedByUserIdOk returns a tuple with the AmendedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendedByUserId

`func (o *AttachmentsPutAttachment) SetAmendedByUserId(v int32)`

SetAmendedByUserId sets AmendedByUserId field to given value.

### HasAmendedByUserId

`func (o *AttachmentsPutAttachment) HasAmendedByUserId() bool`

HasAmendedByUserId returns a boolean if a field has been set.

### GetAmendedDate

`func (o *AttachmentsPutAttachment) GetAmendedDate() string`

GetAmendedDate returns the AmendedDate field if non-nil, zero value otherwise.

### GetAmendedDateOk

`func (o *AttachmentsPutAttachment) GetAmendedDateOk() (*string, bool)`

GetAmendedDateOk returns a tuple with the AmendedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendedDate

`func (o *AttachmentsPutAttachment) SetAmendedDate(v string)`

SetAmendedDate sets AmendedDate field to given value.

### HasAmendedDate

`func (o *AttachmentsPutAttachment) HasAmendedDate() bool`

HasAmendedDate returns a boolean if a field has been set.

### GetBrokenLinkCount

`func (o *AttachmentsPutAttachment) GetBrokenLinkCount() int32`

GetBrokenLinkCount returns the BrokenLinkCount field if non-nil, zero value otherwise.

### GetBrokenLinkCountOk

`func (o *AttachmentsPutAttachment) GetBrokenLinkCountOk() (*int32, bool)`

GetBrokenLinkCountOk returns a tuple with the BrokenLinkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokenLinkCount

`func (o *AttachmentsPutAttachment) SetBrokenLinkCount(v int32)`

SetBrokenLinkCount sets BrokenLinkCount field to given value.

### HasBrokenLinkCount

`func (o *AttachmentsPutAttachment) HasBrokenLinkCount() bool`

HasBrokenLinkCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


