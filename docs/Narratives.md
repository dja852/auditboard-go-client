# Narratives

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

### NewNarratives

`func NewNarratives(uid string, name string, ) *Narratives`

NewNarratives instantiates a new Narratives object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNarrativesWithDefaults

`func NewNarrativesWithDefaults() *Narratives`

NewNarrativesWithDefaults instantiates a new Narratives object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Narratives) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Narratives) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Narratives) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Narratives) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *Narratives) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Narratives) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Narratives) SetUid(v string)`

SetUid sets Uid field to given value.


### GetCreatedAt

`func (o *Narratives) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Narratives) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Narratives) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Narratives) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Narratives) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Narratives) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Narratives) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Narratives) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Narratives) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Narratives) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Narratives) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Narratives) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *Narratives) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Narratives) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Narratives) SetName(v string)`

SetName sets Name field to given value.


### GetSummary

`func (o *Narratives) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Narratives) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Narratives) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Narratives) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetBody

`func (o *Narratives) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Narratives) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Narratives) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *Narratives) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetLockUserId

`func (o *Narratives) GetLockUserId() int32`

GetLockUserId returns the LockUserId field if non-nil, zero value otherwise.

### GetLockUserIdOk

`func (o *Narratives) GetLockUserIdOk() (*int32, bool)`

GetLockUserIdOk returns a tuple with the LockUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockUserId

`func (o *Narratives) SetLockUserId(v int32)`

SetLockUserId sets LockUserId field to given value.

### HasLockUserId

`func (o *Narratives) HasLockUserId() bool`

HasLockUserId returns a boolean if a field has been set.

### GetLockedAt

`func (o *Narratives) GetLockedAt() string`

GetLockedAt returns the LockedAt field if non-nil, zero value otherwise.

### GetLockedAtOk

`func (o *Narratives) GetLockedAtOk() (*string, bool)`

GetLockedAtOk returns a tuple with the LockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedAt

`func (o *Narratives) SetLockedAt(v string)`

SetLockedAt sets LockedAt field to given value.

### HasLockedAt

`func (o *Narratives) HasLockedAt() bool`

HasLockedAt returns a boolean if a field has been set.

### GetPageLayout

`func (o *Narratives) GetPageLayout() interface{}`

GetPageLayout returns the PageLayout field if non-nil, zero value otherwise.

### GetPageLayoutOk

`func (o *Narratives) GetPageLayoutOk() (*interface{}, bool)`

GetPageLayoutOk returns a tuple with the PageLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageLayout

`func (o *Narratives) SetPageLayout(v interface{})`

SetPageLayout sets PageLayout field to given value.

### HasPageLayout

`func (o *Narratives) HasPageLayout() bool`

HasPageLayout returns a boolean if a field has been set.

### SetPageLayoutNil

`func (o *Narratives) SetPageLayoutNil(b bool)`

 SetPageLayoutNil sets the value for PageLayout to be an explicit nil

### UnsetPageLayout
`func (o *Narratives) UnsetPageLayout()`

UnsetPageLayout ensures that no value is present for PageLayout, not even an explicit nil
### GetDescription

`func (o *Narratives) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Narratives) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Narratives) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Narratives) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *Narratives) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Narratives) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Narratives) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Narratives) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetNarrativeTypeId

`func (o *Narratives) GetNarrativeTypeId() int32`

GetNarrativeTypeId returns the NarrativeTypeId field if non-nil, zero value otherwise.

### GetNarrativeTypeIdOk

`func (o *Narratives) GetNarrativeTypeIdOk() (*int32, bool)`

GetNarrativeTypeIdOk returns a tuple with the NarrativeTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarrativeTypeId

`func (o *Narratives) SetNarrativeTypeId(v int32)`

SetNarrativeTypeId sets NarrativeTypeId field to given value.

### HasNarrativeTypeId

`func (o *Narratives) HasNarrativeTypeId() bool`

HasNarrativeTypeId returns a boolean if a field has been set.

### GetLastRevisionByUserId

`func (o *Narratives) GetLastRevisionByUserId() int32`

GetLastRevisionByUserId returns the LastRevisionByUserId field if non-nil, zero value otherwise.

### GetLastRevisionByUserIdOk

`func (o *Narratives) GetLastRevisionByUserIdOk() (*int32, bool)`

GetLastRevisionByUserIdOk returns a tuple with the LastRevisionByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRevisionByUserId

`func (o *Narratives) SetLastRevisionByUserId(v int32)`

SetLastRevisionByUserId sets LastRevisionByUserId field to given value.

### HasLastRevisionByUserId

`func (o *Narratives) HasLastRevisionByUserId() bool`

HasLastRevisionByUserId returns a boolean if a field has been set.

### GetLastRevisionDate

`func (o *Narratives) GetLastRevisionDate() string`

GetLastRevisionDate returns the LastRevisionDate field if non-nil, zero value otherwise.

### GetLastRevisionDateOk

`func (o *Narratives) GetLastRevisionDateOk() (*string, bool)`

GetLastRevisionDateOk returns a tuple with the LastRevisionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRevisionDate

`func (o *Narratives) SetLastRevisionDate(v string)`

SetLastRevisionDate sets LastRevisionDate field to given value.

### HasLastRevisionDate

`func (o *Narratives) HasLastRevisionDate() bool`

HasLastRevisionDate returns a boolean if a field has been set.

### GetLatestDocumentVersionId

`func (o *Narratives) GetLatestDocumentVersionId() int32`

GetLatestDocumentVersionId returns the LatestDocumentVersionId field if non-nil, zero value otherwise.

### GetLatestDocumentVersionIdOk

`func (o *Narratives) GetLatestDocumentVersionIdOk() (*int32, bool)`

GetLatestDocumentVersionIdOk returns a tuple with the LatestDocumentVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDocumentVersionId

`func (o *Narratives) SetLatestDocumentVersionId(v int32)`

SetLatestDocumentVersionId sets LatestDocumentVersionId field to given value.

### HasLatestDocumentVersionId

`func (o *Narratives) HasLatestDocumentVersionId() bool`

HasLatestDocumentVersionId returns a boolean if a field has been set.

### GetScopes

`func (o *Narratives) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *Narratives) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *Narratives) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *Narratives) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *Narratives) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *Narratives) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


