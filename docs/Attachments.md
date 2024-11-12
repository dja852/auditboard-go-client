# Attachments

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

### NewAttachments

`func NewAttachments(type_ string, ) *Attachments`

NewAttachments instantiates a new Attachments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentsWithDefaults

`func NewAttachmentsWithDefaults() *Attachments`

NewAttachmentsWithDefaults instantiates a new Attachments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Attachments) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Attachments) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Attachments) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Attachments) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Attachments) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Attachments) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Attachments) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Attachments) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Attachments) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Attachments) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Attachments) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Attachments) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Attachments) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Attachments) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Attachments) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Attachments) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetType

`func (o *Attachments) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Attachments) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Attachments) SetType(v string)`

SetType sets Type field to given value.


### GetSheetId

`func (o *Attachments) GetSheetId() int32`

GetSheetId returns the SheetId field if non-nil, zero value otherwise.

### GetSheetIdOk

`func (o *Attachments) GetSheetIdOk() (*int32, bool)`

GetSheetIdOk returns a tuple with the SheetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheetId

`func (o *Attachments) SetSheetId(v int32)`

SetSheetId sets SheetId field to given value.

### HasSheetId

`func (o *Attachments) HasSheetId() bool`

HasSheetId returns a boolean if a field has been set.

### GetDocumentId

`func (o *Attachments) GetDocumentId() int32`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *Attachments) GetDocumentIdOk() (*int32, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *Attachments) SetDocumentId(v int32)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *Attachments) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetAttachableId

`func (o *Attachments) GetAttachableId() int32`

GetAttachableId returns the AttachableId field if non-nil, zero value otherwise.

### GetAttachableIdOk

`func (o *Attachments) GetAttachableIdOk() (*int32, bool)`

GetAttachableIdOk returns a tuple with the AttachableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachableId

`func (o *Attachments) SetAttachableId(v int32)`

SetAttachableId sets AttachableId field to given value.

### HasAttachableId

`func (o *Attachments) HasAttachableId() bool`

HasAttachableId returns a boolean if a field has been set.

### GetAttachableType

`func (o *Attachments) GetAttachableType() string`

GetAttachableType returns the AttachableType field if non-nil, zero value otherwise.

### GetAttachableTypeOk

`func (o *Attachments) GetAttachableTypeOk() (*string, bool)`

GetAttachableTypeOk returns a tuple with the AttachableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachableType

`func (o *Attachments) SetAttachableType(v string)`

SetAttachableType sets AttachableType field to given value.

### HasAttachableType

`func (o *Attachments) HasAttachableType() bool`

HasAttachableType returns a boolean if a field has been set.

### GetLockUserId

`func (o *Attachments) GetLockUserId() int32`

GetLockUserId returns the LockUserId field if non-nil, zero value otherwise.

### GetLockUserIdOk

`func (o *Attachments) GetLockUserIdOk() (*int32, bool)`

GetLockUserIdOk returns a tuple with the LockUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockUserId

`func (o *Attachments) SetLockUserId(v int32)`

SetLockUserId sets LockUserId field to given value.

### HasLockUserId

`func (o *Attachments) HasLockUserId() bool`

HasLockUserId returns a boolean if a field has been set.

### GetLockedAt

`func (o *Attachments) GetLockedAt() string`

GetLockedAt returns the LockedAt field if non-nil, zero value otherwise.

### GetLockedAtOk

`func (o *Attachments) GetLockedAtOk() (*string, bool)`

GetLockedAtOk returns a tuple with the LockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedAt

`func (o *Attachments) SetLockedAt(v string)`

SetLockedAt sets LockedAt field to given value.

### HasLockedAt

`func (o *Attachments) HasLockedAt() bool`

HasLockedAt returns a boolean if a field has been set.

### GetName

`func (o *Attachments) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Attachments) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Attachments) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Attachments) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFileId

`func (o *Attachments) GetFileId() int32`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *Attachments) GetFileIdOk() (*int32, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *Attachments) SetFileId(v int32)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *Attachments) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetSortOrder

`func (o *Attachments) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *Attachments) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *Attachments) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *Attachments) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetLinkifySourceFileId

`func (o *Attachments) GetLinkifySourceFileId() int32`

GetLinkifySourceFileId returns the LinkifySourceFileId field if non-nil, zero value otherwise.

### GetLinkifySourceFileIdOk

`func (o *Attachments) GetLinkifySourceFileIdOk() (*int32, bool)`

GetLinkifySourceFileIdOk returns a tuple with the LinkifySourceFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkifySourceFileId

`func (o *Attachments) SetLinkifySourceFileId(v int32)`

SetLinkifySourceFileId sets LinkifySourceFileId field to given value.

### HasLinkifySourceFileId

`func (o *Attachments) HasLinkifySourceFileId() bool`

HasLinkifySourceFileId returns a boolean if a field has been set.

### GetIsLinkified

`func (o *Attachments) GetIsLinkified() bool`

GetIsLinkified returns the IsLinkified field if non-nil, zero value otherwise.

### GetIsLinkifiedOk

`func (o *Attachments) GetIsLinkifiedOk() (*bool, bool)`

GetIsLinkifiedOk returns a tuple with the IsLinkified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLinkified

`func (o *Attachments) SetIsLinkified(v bool)`

SetIsLinkified sets IsLinkified field to given value.

### HasIsLinkified

`func (o *Attachments) HasIsLinkified() bool`

HasIsLinkified returns a boolean if a field has been set.

### GetLastLockUserId

`func (o *Attachments) GetLastLockUserId() int32`

GetLastLockUserId returns the LastLockUserId field if non-nil, zero value otherwise.

### GetLastLockUserIdOk

`func (o *Attachments) GetLastLockUserIdOk() (*int32, bool)`

GetLastLockUserIdOk returns a tuple with the LastLockUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLockUserId

`func (o *Attachments) SetLastLockUserId(v int32)`

SetLastLockUserId sets LastLockUserId field to given value.

### HasLastLockUserId

`func (o *Attachments) HasLastLockUserId() bool`

HasLastLockUserId returns a boolean if a field has been set.

### GetWordReportOptions

`func (o *Attachments) GetWordReportOptions() interface{}`

GetWordReportOptions returns the WordReportOptions field if non-nil, zero value otherwise.

### GetWordReportOptionsOk

`func (o *Attachments) GetWordReportOptionsOk() (*interface{}, bool)`

GetWordReportOptionsOk returns a tuple with the WordReportOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordReportOptions

`func (o *Attachments) SetWordReportOptions(v interface{})`

SetWordReportOptions sets WordReportOptions field to given value.

### HasWordReportOptions

`func (o *Attachments) HasWordReportOptions() bool`

HasWordReportOptions returns a boolean if a field has been set.

### SetWordReportOptionsNil

`func (o *Attachments) SetWordReportOptionsNil(b bool)`

 SetWordReportOptionsNil sets the value for WordReportOptions to be an explicit nil

### UnsetWordReportOptions
`func (o *Attachments) UnsetWordReportOptions()`

UnsetWordReportOptions ensures that no value is present for WordReportOptions, not even an explicit nil
### GetStatus

`func (o *Attachments) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Attachments) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Attachments) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Attachments) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReviewedByUserId

`func (o *Attachments) GetReviewedByUserId() int32`

GetReviewedByUserId returns the ReviewedByUserId field if non-nil, zero value otherwise.

### GetReviewedByUserIdOk

`func (o *Attachments) GetReviewedByUserIdOk() (*int32, bool)`

GetReviewedByUserIdOk returns a tuple with the ReviewedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedByUserId

`func (o *Attachments) SetReviewedByUserId(v int32)`

SetReviewedByUserId sets ReviewedByUserId field to given value.

### HasReviewedByUserId

`func (o *Attachments) HasReviewedByUserId() bool`

HasReviewedByUserId returns a boolean if a field has been set.

### GetReviewedDate

`func (o *Attachments) GetReviewedDate() string`

GetReviewedDate returns the ReviewedDate field if non-nil, zero value otherwise.

### GetReviewedDateOk

`func (o *Attachments) GetReviewedDateOk() (*string, bool)`

GetReviewedDateOk returns a tuple with the ReviewedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedDate

`func (o *Attachments) SetReviewedDate(v string)`

SetReviewedDate sets ReviewedDate field to given value.

### HasReviewedDate

`func (o *Attachments) HasReviewedDate() bool`

HasReviewedDate returns a boolean if a field has been set.

### GetAmendedByUserId

`func (o *Attachments) GetAmendedByUserId() int32`

GetAmendedByUserId returns the AmendedByUserId field if non-nil, zero value otherwise.

### GetAmendedByUserIdOk

`func (o *Attachments) GetAmendedByUserIdOk() (*int32, bool)`

GetAmendedByUserIdOk returns a tuple with the AmendedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendedByUserId

`func (o *Attachments) SetAmendedByUserId(v int32)`

SetAmendedByUserId sets AmendedByUserId field to given value.

### HasAmendedByUserId

`func (o *Attachments) HasAmendedByUserId() bool`

HasAmendedByUserId returns a boolean if a field has been set.

### GetAmendedDate

`func (o *Attachments) GetAmendedDate() string`

GetAmendedDate returns the AmendedDate field if non-nil, zero value otherwise.

### GetAmendedDateOk

`func (o *Attachments) GetAmendedDateOk() (*string, bool)`

GetAmendedDateOk returns a tuple with the AmendedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendedDate

`func (o *Attachments) SetAmendedDate(v string)`

SetAmendedDate sets AmendedDate field to given value.

### HasAmendedDate

`func (o *Attachments) HasAmendedDate() bool`

HasAmendedDate returns a boolean if a field has been set.

### GetBrokenLinkCount

`func (o *Attachments) GetBrokenLinkCount() int32`

GetBrokenLinkCount returns the BrokenLinkCount field if non-nil, zero value otherwise.

### GetBrokenLinkCountOk

`func (o *Attachments) GetBrokenLinkCountOk() (*int32, bool)`

GetBrokenLinkCountOk returns a tuple with the BrokenLinkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokenLinkCount

`func (o *Attachments) SetBrokenLinkCount(v int32)`

SetBrokenLinkCount sets BrokenLinkCount field to given value.

### HasBrokenLinkCount

`func (o *Attachments) HasBrokenLinkCount() bool`

HasBrokenLinkCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


