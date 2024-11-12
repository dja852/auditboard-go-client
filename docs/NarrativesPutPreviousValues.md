# NarrativesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**LockUserId** | Pointer to **int32** |  | [optional] 
**LockedAt** | Pointer to **string** |  | [optional] 
**PageLayout** | Pointer to **interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**NarrativeTypeId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;narrative_types.id&#x60;.&lt;fk table&#x3D;&#39;narrative_types&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**LastRevisionByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**LastRevisionDate** | Pointer to **string** |  | [optional] 
**LatestDocumentVersionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;document_versions.id&#x60;.&lt;fk table&#x3D;&#39;document_versions&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Scopes** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewNarrativesPutPreviousValues

`func NewNarrativesPutPreviousValues() *NarrativesPutPreviousValues`

NewNarrativesPutPreviousValues instantiates a new NarrativesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNarrativesPutPreviousValuesWithDefaults

`func NewNarrativesPutPreviousValuesWithDefaults() *NarrativesPutPreviousValues`

NewNarrativesPutPreviousValuesWithDefaults instantiates a new NarrativesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NarrativesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NarrativesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NarrativesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NarrativesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *NarrativesPutPreviousValues) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *NarrativesPutPreviousValues) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *NarrativesPutPreviousValues) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *NarrativesPutPreviousValues) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NarrativesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NarrativesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NarrativesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NarrativesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NarrativesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NarrativesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NarrativesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NarrativesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *NarrativesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *NarrativesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *NarrativesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *NarrativesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *NarrativesPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NarrativesPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NarrativesPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NarrativesPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSummary

`func (o *NarrativesPutPreviousValues) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *NarrativesPutPreviousValues) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *NarrativesPutPreviousValues) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *NarrativesPutPreviousValues) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetBody

`func (o *NarrativesPutPreviousValues) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *NarrativesPutPreviousValues) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *NarrativesPutPreviousValues) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *NarrativesPutPreviousValues) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetLockUserId

`func (o *NarrativesPutPreviousValues) GetLockUserId() int32`

GetLockUserId returns the LockUserId field if non-nil, zero value otherwise.

### GetLockUserIdOk

`func (o *NarrativesPutPreviousValues) GetLockUserIdOk() (*int32, bool)`

GetLockUserIdOk returns a tuple with the LockUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockUserId

`func (o *NarrativesPutPreviousValues) SetLockUserId(v int32)`

SetLockUserId sets LockUserId field to given value.

### HasLockUserId

`func (o *NarrativesPutPreviousValues) HasLockUserId() bool`

HasLockUserId returns a boolean if a field has been set.

### GetLockedAt

`func (o *NarrativesPutPreviousValues) GetLockedAt() string`

GetLockedAt returns the LockedAt field if non-nil, zero value otherwise.

### GetLockedAtOk

`func (o *NarrativesPutPreviousValues) GetLockedAtOk() (*string, bool)`

GetLockedAtOk returns a tuple with the LockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedAt

`func (o *NarrativesPutPreviousValues) SetLockedAt(v string)`

SetLockedAt sets LockedAt field to given value.

### HasLockedAt

`func (o *NarrativesPutPreviousValues) HasLockedAt() bool`

HasLockedAt returns a boolean if a field has been set.

### GetPageLayout

`func (o *NarrativesPutPreviousValues) GetPageLayout() interface{}`

GetPageLayout returns the PageLayout field if non-nil, zero value otherwise.

### GetPageLayoutOk

`func (o *NarrativesPutPreviousValues) GetPageLayoutOk() (*interface{}, bool)`

GetPageLayoutOk returns a tuple with the PageLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageLayout

`func (o *NarrativesPutPreviousValues) SetPageLayout(v interface{})`

SetPageLayout sets PageLayout field to given value.

### HasPageLayout

`func (o *NarrativesPutPreviousValues) HasPageLayout() bool`

HasPageLayout returns a boolean if a field has been set.

### SetPageLayoutNil

`func (o *NarrativesPutPreviousValues) SetPageLayoutNil(b bool)`

 SetPageLayoutNil sets the value for PageLayout to be an explicit nil

### UnsetPageLayout
`func (o *NarrativesPutPreviousValues) UnsetPageLayout()`

UnsetPageLayout ensures that no value is present for PageLayout, not even an explicit nil
### GetDescription

`func (o *NarrativesPutPreviousValues) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NarrativesPutPreviousValues) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NarrativesPutPreviousValues) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NarrativesPutPreviousValues) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *NarrativesPutPreviousValues) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NarrativesPutPreviousValues) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NarrativesPutPreviousValues) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *NarrativesPutPreviousValues) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetNarrativeTypeId

`func (o *NarrativesPutPreviousValues) GetNarrativeTypeId() int32`

GetNarrativeTypeId returns the NarrativeTypeId field if non-nil, zero value otherwise.

### GetNarrativeTypeIdOk

`func (o *NarrativesPutPreviousValues) GetNarrativeTypeIdOk() (*int32, bool)`

GetNarrativeTypeIdOk returns a tuple with the NarrativeTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarrativeTypeId

`func (o *NarrativesPutPreviousValues) SetNarrativeTypeId(v int32)`

SetNarrativeTypeId sets NarrativeTypeId field to given value.

### HasNarrativeTypeId

`func (o *NarrativesPutPreviousValues) HasNarrativeTypeId() bool`

HasNarrativeTypeId returns a boolean if a field has been set.

### GetLastRevisionByUserId

`func (o *NarrativesPutPreviousValues) GetLastRevisionByUserId() int32`

GetLastRevisionByUserId returns the LastRevisionByUserId field if non-nil, zero value otherwise.

### GetLastRevisionByUserIdOk

`func (o *NarrativesPutPreviousValues) GetLastRevisionByUserIdOk() (*int32, bool)`

GetLastRevisionByUserIdOk returns a tuple with the LastRevisionByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRevisionByUserId

`func (o *NarrativesPutPreviousValues) SetLastRevisionByUserId(v int32)`

SetLastRevisionByUserId sets LastRevisionByUserId field to given value.

### HasLastRevisionByUserId

`func (o *NarrativesPutPreviousValues) HasLastRevisionByUserId() bool`

HasLastRevisionByUserId returns a boolean if a field has been set.

### GetLastRevisionDate

`func (o *NarrativesPutPreviousValues) GetLastRevisionDate() string`

GetLastRevisionDate returns the LastRevisionDate field if non-nil, zero value otherwise.

### GetLastRevisionDateOk

`func (o *NarrativesPutPreviousValues) GetLastRevisionDateOk() (*string, bool)`

GetLastRevisionDateOk returns a tuple with the LastRevisionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRevisionDate

`func (o *NarrativesPutPreviousValues) SetLastRevisionDate(v string)`

SetLastRevisionDate sets LastRevisionDate field to given value.

### HasLastRevisionDate

`func (o *NarrativesPutPreviousValues) HasLastRevisionDate() bool`

HasLastRevisionDate returns a boolean if a field has been set.

### GetLatestDocumentVersionId

`func (o *NarrativesPutPreviousValues) GetLatestDocumentVersionId() int32`

GetLatestDocumentVersionId returns the LatestDocumentVersionId field if non-nil, zero value otherwise.

### GetLatestDocumentVersionIdOk

`func (o *NarrativesPutPreviousValues) GetLatestDocumentVersionIdOk() (*int32, bool)`

GetLatestDocumentVersionIdOk returns a tuple with the LatestDocumentVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDocumentVersionId

`func (o *NarrativesPutPreviousValues) SetLatestDocumentVersionId(v int32)`

SetLatestDocumentVersionId sets LatestDocumentVersionId field to given value.

### HasLatestDocumentVersionId

`func (o *NarrativesPutPreviousValues) HasLatestDocumentVersionId() bool`

HasLatestDocumentVersionId returns a boolean if a field has been set.

### GetScopes

`func (o *NarrativesPutPreviousValues) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *NarrativesPutPreviousValues) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *NarrativesPutPreviousValues) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *NarrativesPutPreviousValues) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *NarrativesPutPreviousValues) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *NarrativesPutPreviousValues) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


