# AttachmentsPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] 
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

### NewAttachmentsPutPreviousValues

`func NewAttachmentsPutPreviousValues() *AttachmentsPutPreviousValues`

NewAttachmentsPutPreviousValues instantiates a new AttachmentsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentsPutPreviousValuesWithDefaults

`func NewAttachmentsPutPreviousValuesWithDefaults() *AttachmentsPutPreviousValues`

NewAttachmentsPutPreviousValuesWithDefaults instantiates a new AttachmentsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AttachmentsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttachmentsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttachmentsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AttachmentsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AttachmentsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AttachmentsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AttachmentsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AttachmentsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AttachmentsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AttachmentsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AttachmentsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AttachmentsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AttachmentsPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AttachmentsPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AttachmentsPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AttachmentsPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetType

`func (o *AttachmentsPutPreviousValues) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttachmentsPutPreviousValues) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttachmentsPutPreviousValues) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AttachmentsPutPreviousValues) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSheetId

`func (o *AttachmentsPutPreviousValues) GetSheetId() int32`

GetSheetId returns the SheetId field if non-nil, zero value otherwise.

### GetSheetIdOk

`func (o *AttachmentsPutPreviousValues) GetSheetIdOk() (*int32, bool)`

GetSheetIdOk returns a tuple with the SheetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheetId

`func (o *AttachmentsPutPreviousValues) SetSheetId(v int32)`

SetSheetId sets SheetId field to given value.

### HasSheetId

`func (o *AttachmentsPutPreviousValues) HasSheetId() bool`

HasSheetId returns a boolean if a field has been set.

### GetDocumentId

`func (o *AttachmentsPutPreviousValues) GetDocumentId() int32`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *AttachmentsPutPreviousValues) GetDocumentIdOk() (*int32, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *AttachmentsPutPreviousValues) SetDocumentId(v int32)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *AttachmentsPutPreviousValues) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetAttachableId

`func (o *AttachmentsPutPreviousValues) GetAttachableId() int32`

GetAttachableId returns the AttachableId field if non-nil, zero value otherwise.

### GetAttachableIdOk

`func (o *AttachmentsPutPreviousValues) GetAttachableIdOk() (*int32, bool)`

GetAttachableIdOk returns a tuple with the AttachableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachableId

`func (o *AttachmentsPutPreviousValues) SetAttachableId(v int32)`

SetAttachableId sets AttachableId field to given value.

### HasAttachableId

`func (o *AttachmentsPutPreviousValues) HasAttachableId() bool`

HasAttachableId returns a boolean if a field has been set.

### GetAttachableType

`func (o *AttachmentsPutPreviousValues) GetAttachableType() string`

GetAttachableType returns the AttachableType field if non-nil, zero value otherwise.

### GetAttachableTypeOk

`func (o *AttachmentsPutPreviousValues) GetAttachableTypeOk() (*string, bool)`

GetAttachableTypeOk returns a tuple with the AttachableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachableType

`func (o *AttachmentsPutPreviousValues) SetAttachableType(v string)`

SetAttachableType sets AttachableType field to given value.

### HasAttachableType

`func (o *AttachmentsPutPreviousValues) HasAttachableType() bool`

HasAttachableType returns a boolean if a field has been set.

### GetLockUserId

`func (o *AttachmentsPutPreviousValues) GetLockUserId() int32`

GetLockUserId returns the LockUserId field if non-nil, zero value otherwise.

### GetLockUserIdOk

`func (o *AttachmentsPutPreviousValues) GetLockUserIdOk() (*int32, bool)`

GetLockUserIdOk returns a tuple with the LockUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockUserId

`func (o *AttachmentsPutPreviousValues) SetLockUserId(v int32)`

SetLockUserId sets LockUserId field to given value.

### HasLockUserId

`func (o *AttachmentsPutPreviousValues) HasLockUserId() bool`

HasLockUserId returns a boolean if a field has been set.

### GetLockedAt

`func (o *AttachmentsPutPreviousValues) GetLockedAt() string`

GetLockedAt returns the LockedAt field if non-nil, zero value otherwise.

### GetLockedAtOk

`func (o *AttachmentsPutPreviousValues) GetLockedAtOk() (*string, bool)`

GetLockedAtOk returns a tuple with the LockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedAt

`func (o *AttachmentsPutPreviousValues) SetLockedAt(v string)`

SetLockedAt sets LockedAt field to given value.

### HasLockedAt

`func (o *AttachmentsPutPreviousValues) HasLockedAt() bool`

HasLockedAt returns a boolean if a field has been set.

### GetName

`func (o *AttachmentsPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttachmentsPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttachmentsPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AttachmentsPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFileId

`func (o *AttachmentsPutPreviousValues) GetFileId() int32`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *AttachmentsPutPreviousValues) GetFileIdOk() (*int32, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *AttachmentsPutPreviousValues) SetFileId(v int32)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *AttachmentsPutPreviousValues) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetSortOrder

`func (o *AttachmentsPutPreviousValues) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AttachmentsPutPreviousValues) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AttachmentsPutPreviousValues) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *AttachmentsPutPreviousValues) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetLinkifySourceFileId

`func (o *AttachmentsPutPreviousValues) GetLinkifySourceFileId() int32`

GetLinkifySourceFileId returns the LinkifySourceFileId field if non-nil, zero value otherwise.

### GetLinkifySourceFileIdOk

`func (o *AttachmentsPutPreviousValues) GetLinkifySourceFileIdOk() (*int32, bool)`

GetLinkifySourceFileIdOk returns a tuple with the LinkifySourceFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkifySourceFileId

`func (o *AttachmentsPutPreviousValues) SetLinkifySourceFileId(v int32)`

SetLinkifySourceFileId sets LinkifySourceFileId field to given value.

### HasLinkifySourceFileId

`func (o *AttachmentsPutPreviousValues) HasLinkifySourceFileId() bool`

HasLinkifySourceFileId returns a boolean if a field has been set.

### GetIsLinkified

`func (o *AttachmentsPutPreviousValues) GetIsLinkified() bool`

GetIsLinkified returns the IsLinkified field if non-nil, zero value otherwise.

### GetIsLinkifiedOk

`func (o *AttachmentsPutPreviousValues) GetIsLinkifiedOk() (*bool, bool)`

GetIsLinkifiedOk returns a tuple with the IsLinkified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLinkified

`func (o *AttachmentsPutPreviousValues) SetIsLinkified(v bool)`

SetIsLinkified sets IsLinkified field to given value.

### HasIsLinkified

`func (o *AttachmentsPutPreviousValues) HasIsLinkified() bool`

HasIsLinkified returns a boolean if a field has been set.

### GetLastLockUserId

`func (o *AttachmentsPutPreviousValues) GetLastLockUserId() int32`

GetLastLockUserId returns the LastLockUserId field if non-nil, zero value otherwise.

### GetLastLockUserIdOk

`func (o *AttachmentsPutPreviousValues) GetLastLockUserIdOk() (*int32, bool)`

GetLastLockUserIdOk returns a tuple with the LastLockUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLockUserId

`func (o *AttachmentsPutPreviousValues) SetLastLockUserId(v int32)`

SetLastLockUserId sets LastLockUserId field to given value.

### HasLastLockUserId

`func (o *AttachmentsPutPreviousValues) HasLastLockUserId() bool`

HasLastLockUserId returns a boolean if a field has been set.

### GetWordReportOptions

`func (o *AttachmentsPutPreviousValues) GetWordReportOptions() interface{}`

GetWordReportOptions returns the WordReportOptions field if non-nil, zero value otherwise.

### GetWordReportOptionsOk

`func (o *AttachmentsPutPreviousValues) GetWordReportOptionsOk() (*interface{}, bool)`

GetWordReportOptionsOk returns a tuple with the WordReportOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordReportOptions

`func (o *AttachmentsPutPreviousValues) SetWordReportOptions(v interface{})`

SetWordReportOptions sets WordReportOptions field to given value.

### HasWordReportOptions

`func (o *AttachmentsPutPreviousValues) HasWordReportOptions() bool`

HasWordReportOptions returns a boolean if a field has been set.

### SetWordReportOptionsNil

`func (o *AttachmentsPutPreviousValues) SetWordReportOptionsNil(b bool)`

 SetWordReportOptionsNil sets the value for WordReportOptions to be an explicit nil

### UnsetWordReportOptions
`func (o *AttachmentsPutPreviousValues) UnsetWordReportOptions()`

UnsetWordReportOptions ensures that no value is present for WordReportOptions, not even an explicit nil
### GetStatus

`func (o *AttachmentsPutPreviousValues) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AttachmentsPutPreviousValues) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AttachmentsPutPreviousValues) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AttachmentsPutPreviousValues) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReviewedByUserId

`func (o *AttachmentsPutPreviousValues) GetReviewedByUserId() int32`

GetReviewedByUserId returns the ReviewedByUserId field if non-nil, zero value otherwise.

### GetReviewedByUserIdOk

`func (o *AttachmentsPutPreviousValues) GetReviewedByUserIdOk() (*int32, bool)`

GetReviewedByUserIdOk returns a tuple with the ReviewedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedByUserId

`func (o *AttachmentsPutPreviousValues) SetReviewedByUserId(v int32)`

SetReviewedByUserId sets ReviewedByUserId field to given value.

### HasReviewedByUserId

`func (o *AttachmentsPutPreviousValues) HasReviewedByUserId() bool`

HasReviewedByUserId returns a boolean if a field has been set.

### GetReviewedDate

`func (o *AttachmentsPutPreviousValues) GetReviewedDate() string`

GetReviewedDate returns the ReviewedDate field if non-nil, zero value otherwise.

### GetReviewedDateOk

`func (o *AttachmentsPutPreviousValues) GetReviewedDateOk() (*string, bool)`

GetReviewedDateOk returns a tuple with the ReviewedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedDate

`func (o *AttachmentsPutPreviousValues) SetReviewedDate(v string)`

SetReviewedDate sets ReviewedDate field to given value.

### HasReviewedDate

`func (o *AttachmentsPutPreviousValues) HasReviewedDate() bool`

HasReviewedDate returns a boolean if a field has been set.

### GetAmendedByUserId

`func (o *AttachmentsPutPreviousValues) GetAmendedByUserId() int32`

GetAmendedByUserId returns the AmendedByUserId field if non-nil, zero value otherwise.

### GetAmendedByUserIdOk

`func (o *AttachmentsPutPreviousValues) GetAmendedByUserIdOk() (*int32, bool)`

GetAmendedByUserIdOk returns a tuple with the AmendedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendedByUserId

`func (o *AttachmentsPutPreviousValues) SetAmendedByUserId(v int32)`

SetAmendedByUserId sets AmendedByUserId field to given value.

### HasAmendedByUserId

`func (o *AttachmentsPutPreviousValues) HasAmendedByUserId() bool`

HasAmendedByUserId returns a boolean if a field has been set.

### GetAmendedDate

`func (o *AttachmentsPutPreviousValues) GetAmendedDate() string`

GetAmendedDate returns the AmendedDate field if non-nil, zero value otherwise.

### GetAmendedDateOk

`func (o *AttachmentsPutPreviousValues) GetAmendedDateOk() (*string, bool)`

GetAmendedDateOk returns a tuple with the AmendedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendedDate

`func (o *AttachmentsPutPreviousValues) SetAmendedDate(v string)`

SetAmendedDate sets AmendedDate field to given value.

### HasAmendedDate

`func (o *AttachmentsPutPreviousValues) HasAmendedDate() bool`

HasAmendedDate returns a boolean if a field has been set.

### GetBrokenLinkCount

`func (o *AttachmentsPutPreviousValues) GetBrokenLinkCount() int32`

GetBrokenLinkCount returns the BrokenLinkCount field if non-nil, zero value otherwise.

### GetBrokenLinkCountOk

`func (o *AttachmentsPutPreviousValues) GetBrokenLinkCountOk() (*int32, bool)`

GetBrokenLinkCountOk returns a tuple with the BrokenLinkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokenLinkCount

`func (o *AttachmentsPutPreviousValues) SetBrokenLinkCount(v int32)`

SetBrokenLinkCount sets BrokenLinkCount field to given value.

### HasBrokenLinkCount

`func (o *AttachmentsPutPreviousValues) HasBrokenLinkCount() bool`

HasBrokenLinkCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


