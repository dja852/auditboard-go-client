# NarrativesPutNarrative

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Uid** | **string** |  | 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
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

### NewNarrativesPutNarrative

`func NewNarrativesPutNarrative(uid string, name string, ) *NarrativesPutNarrative`

NewNarrativesPutNarrative instantiates a new NarrativesPutNarrative object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNarrativesPutNarrativeWithDefaults

`func NewNarrativesPutNarrativeWithDefaults() *NarrativesPutNarrative`

NewNarrativesPutNarrativeWithDefaults instantiates a new NarrativesPutNarrative object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NarrativesPutNarrative) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NarrativesPutNarrative) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NarrativesPutNarrative) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NarrativesPutNarrative) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *NarrativesPutNarrative) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *NarrativesPutNarrative) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *NarrativesPutNarrative) SetUid(v string)`

SetUid sets Uid field to given value.


### GetCreatedAt

`func (o *NarrativesPutNarrative) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NarrativesPutNarrative) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NarrativesPutNarrative) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NarrativesPutNarrative) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NarrativesPutNarrative) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NarrativesPutNarrative) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NarrativesPutNarrative) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NarrativesPutNarrative) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *NarrativesPutNarrative) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *NarrativesPutNarrative) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *NarrativesPutNarrative) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *NarrativesPutNarrative) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *NarrativesPutNarrative) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NarrativesPutNarrative) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NarrativesPutNarrative) SetName(v string)`

SetName sets Name field to given value.


### GetSummary

`func (o *NarrativesPutNarrative) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *NarrativesPutNarrative) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *NarrativesPutNarrative) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *NarrativesPutNarrative) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetBody

`func (o *NarrativesPutNarrative) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *NarrativesPutNarrative) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *NarrativesPutNarrative) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *NarrativesPutNarrative) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetLockUserId

`func (o *NarrativesPutNarrative) GetLockUserId() int32`

GetLockUserId returns the LockUserId field if non-nil, zero value otherwise.

### GetLockUserIdOk

`func (o *NarrativesPutNarrative) GetLockUserIdOk() (*int32, bool)`

GetLockUserIdOk returns a tuple with the LockUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockUserId

`func (o *NarrativesPutNarrative) SetLockUserId(v int32)`

SetLockUserId sets LockUserId field to given value.

### HasLockUserId

`func (o *NarrativesPutNarrative) HasLockUserId() bool`

HasLockUserId returns a boolean if a field has been set.

### GetLockedAt

`func (o *NarrativesPutNarrative) GetLockedAt() string`

GetLockedAt returns the LockedAt field if non-nil, zero value otherwise.

### GetLockedAtOk

`func (o *NarrativesPutNarrative) GetLockedAtOk() (*string, bool)`

GetLockedAtOk returns a tuple with the LockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedAt

`func (o *NarrativesPutNarrative) SetLockedAt(v string)`

SetLockedAt sets LockedAt field to given value.

### HasLockedAt

`func (o *NarrativesPutNarrative) HasLockedAt() bool`

HasLockedAt returns a boolean if a field has been set.

### GetPageLayout

`func (o *NarrativesPutNarrative) GetPageLayout() interface{}`

GetPageLayout returns the PageLayout field if non-nil, zero value otherwise.

### GetPageLayoutOk

`func (o *NarrativesPutNarrative) GetPageLayoutOk() (*interface{}, bool)`

GetPageLayoutOk returns a tuple with the PageLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageLayout

`func (o *NarrativesPutNarrative) SetPageLayout(v interface{})`

SetPageLayout sets PageLayout field to given value.

### HasPageLayout

`func (o *NarrativesPutNarrative) HasPageLayout() bool`

HasPageLayout returns a boolean if a field has been set.

### SetPageLayoutNil

`func (o *NarrativesPutNarrative) SetPageLayoutNil(b bool)`

 SetPageLayoutNil sets the value for PageLayout to be an explicit nil

### UnsetPageLayout
`func (o *NarrativesPutNarrative) UnsetPageLayout()`

UnsetPageLayout ensures that no value is present for PageLayout, not even an explicit nil
### GetDescription

`func (o *NarrativesPutNarrative) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NarrativesPutNarrative) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NarrativesPutNarrative) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NarrativesPutNarrative) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *NarrativesPutNarrative) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NarrativesPutNarrative) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NarrativesPutNarrative) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *NarrativesPutNarrative) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetNarrativeTypeId

`func (o *NarrativesPutNarrative) GetNarrativeTypeId() int32`

GetNarrativeTypeId returns the NarrativeTypeId field if non-nil, zero value otherwise.

### GetNarrativeTypeIdOk

`func (o *NarrativesPutNarrative) GetNarrativeTypeIdOk() (*int32, bool)`

GetNarrativeTypeIdOk returns a tuple with the NarrativeTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarrativeTypeId

`func (o *NarrativesPutNarrative) SetNarrativeTypeId(v int32)`

SetNarrativeTypeId sets NarrativeTypeId field to given value.

### HasNarrativeTypeId

`func (o *NarrativesPutNarrative) HasNarrativeTypeId() bool`

HasNarrativeTypeId returns a boolean if a field has been set.

### GetLastRevisionByUserId

`func (o *NarrativesPutNarrative) GetLastRevisionByUserId() int32`

GetLastRevisionByUserId returns the LastRevisionByUserId field if non-nil, zero value otherwise.

### GetLastRevisionByUserIdOk

`func (o *NarrativesPutNarrative) GetLastRevisionByUserIdOk() (*int32, bool)`

GetLastRevisionByUserIdOk returns a tuple with the LastRevisionByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRevisionByUserId

`func (o *NarrativesPutNarrative) SetLastRevisionByUserId(v int32)`

SetLastRevisionByUserId sets LastRevisionByUserId field to given value.

### HasLastRevisionByUserId

`func (o *NarrativesPutNarrative) HasLastRevisionByUserId() bool`

HasLastRevisionByUserId returns a boolean if a field has been set.

### GetLastRevisionDate

`func (o *NarrativesPutNarrative) GetLastRevisionDate() string`

GetLastRevisionDate returns the LastRevisionDate field if non-nil, zero value otherwise.

### GetLastRevisionDateOk

`func (o *NarrativesPutNarrative) GetLastRevisionDateOk() (*string, bool)`

GetLastRevisionDateOk returns a tuple with the LastRevisionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRevisionDate

`func (o *NarrativesPutNarrative) SetLastRevisionDate(v string)`

SetLastRevisionDate sets LastRevisionDate field to given value.

### HasLastRevisionDate

`func (o *NarrativesPutNarrative) HasLastRevisionDate() bool`

HasLastRevisionDate returns a boolean if a field has been set.

### GetLatestDocumentVersionId

`func (o *NarrativesPutNarrative) GetLatestDocumentVersionId() int32`

GetLatestDocumentVersionId returns the LatestDocumentVersionId field if non-nil, zero value otherwise.

### GetLatestDocumentVersionIdOk

`func (o *NarrativesPutNarrative) GetLatestDocumentVersionIdOk() (*int32, bool)`

GetLatestDocumentVersionIdOk returns a tuple with the LatestDocumentVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDocumentVersionId

`func (o *NarrativesPutNarrative) SetLatestDocumentVersionId(v int32)`

SetLatestDocumentVersionId sets LatestDocumentVersionId field to given value.

### HasLatestDocumentVersionId

`func (o *NarrativesPutNarrative) HasLatestDocumentVersionId() bool`

HasLatestDocumentVersionId returns a boolean if a field has been set.

### GetScopes

`func (o *NarrativesPutNarrative) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *NarrativesPutNarrative) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *NarrativesPutNarrative) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *NarrativesPutNarrative) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *NarrativesPutNarrative) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *NarrativesPutNarrative) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


