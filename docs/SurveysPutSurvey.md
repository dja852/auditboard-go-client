# SurveysPutSurvey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**MetaConfig** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**TaskTitle** | Pointer to **string** |  | [optional] [default to ""]
**Type** | Pointer to **string** |  | [optional] [default to "standalone"]
**Status** | **string** |  | [default to "Unlocked"]
**TaskCategoryId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;task_categories.id&#x60;.&lt;fk table&#x3D;&#39;task_categories&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IsArchived** | **bool** |  | [default to false]
**DefaultUsers** | Pointer to **interface{}** |  | [optional] 
**EmailProjectStart** | Pointer to **interface{}** |  | [optional] 
**EmailPreparerDigest** | Pointer to **interface{}** |  | [optional] 
**EmailReviewerDigest** | Pointer to **interface{}** |  | [optional] 
**EmailAdminDigest** | Pointer to **interface{}** |  | [optional] 
**EmailSubscriberDigest** | Pointer to **interface{}** |  | [optional] 
**AdditionalPreparersEnabled** | Pointer to **bool** |  | [optional] [default to false]
**EmailAdditionalPreparerDigest** | Pointer to **interface{}** |  | [optional] 
**ProjectOptions** | Pointer to **interface{}** |  | [optional] 
**AllowYesToAll** | Pointer to **bool** |  | [optional] 
**PackagingIsStandard** | **bool** |  | [default to false]

## Methods

### NewSurveysPutSurvey

`func NewSurveysPutSurvey(status string, isArchived bool, packagingIsStandard bool, ) *SurveysPutSurvey`

NewSurveysPutSurvey instantiates a new SurveysPutSurvey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSurveysPutSurveyWithDefaults

`func NewSurveysPutSurveyWithDefaults() *SurveysPutSurvey`

NewSurveysPutSurveyWithDefaults instantiates a new SurveysPutSurvey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SurveysPutSurvey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SurveysPutSurvey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SurveysPutSurvey) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SurveysPutSurvey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SurveysPutSurvey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SurveysPutSurvey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SurveysPutSurvey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SurveysPutSurvey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetaConfig

`func (o *SurveysPutSurvey) GetMetaConfig() string`

GetMetaConfig returns the MetaConfig field if non-nil, zero value otherwise.

### GetMetaConfigOk

`func (o *SurveysPutSurvey) GetMetaConfigOk() (*string, bool)`

GetMetaConfigOk returns a tuple with the MetaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaConfig

`func (o *SurveysPutSurvey) SetMetaConfig(v string)`

SetMetaConfig sets MetaConfig field to given value.

### HasMetaConfig

`func (o *SurveysPutSurvey) HasMetaConfig() bool`

HasMetaConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SurveysPutSurvey) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SurveysPutSurvey) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SurveysPutSurvey) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SurveysPutSurvey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SurveysPutSurvey) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SurveysPutSurvey) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SurveysPutSurvey) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SurveysPutSurvey) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *SurveysPutSurvey) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SurveysPutSurvey) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SurveysPutSurvey) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *SurveysPutSurvey) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetTaskTitle

`func (o *SurveysPutSurvey) GetTaskTitle() string`

GetTaskTitle returns the TaskTitle field if non-nil, zero value otherwise.

### GetTaskTitleOk

`func (o *SurveysPutSurvey) GetTaskTitleOk() (*string, bool)`

GetTaskTitleOk returns a tuple with the TaskTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskTitle

`func (o *SurveysPutSurvey) SetTaskTitle(v string)`

SetTaskTitle sets TaskTitle field to given value.

### HasTaskTitle

`func (o *SurveysPutSurvey) HasTaskTitle() bool`

HasTaskTitle returns a boolean if a field has been set.

### GetType

`func (o *SurveysPutSurvey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SurveysPutSurvey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SurveysPutSurvey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SurveysPutSurvey) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *SurveysPutSurvey) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SurveysPutSurvey) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SurveysPutSurvey) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTaskCategoryId

`func (o *SurveysPutSurvey) GetTaskCategoryId() int32`

GetTaskCategoryId returns the TaskCategoryId field if non-nil, zero value otherwise.

### GetTaskCategoryIdOk

`func (o *SurveysPutSurvey) GetTaskCategoryIdOk() (*int32, bool)`

GetTaskCategoryIdOk returns a tuple with the TaskCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCategoryId

`func (o *SurveysPutSurvey) SetTaskCategoryId(v int32)`

SetTaskCategoryId sets TaskCategoryId field to given value.

### HasTaskCategoryId

`func (o *SurveysPutSurvey) HasTaskCategoryId() bool`

HasTaskCategoryId returns a boolean if a field has been set.

### GetIsArchived

`func (o *SurveysPutSurvey) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *SurveysPutSurvey) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *SurveysPutSurvey) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.


### GetDefaultUsers

`func (o *SurveysPutSurvey) GetDefaultUsers() interface{}`

GetDefaultUsers returns the DefaultUsers field if non-nil, zero value otherwise.

### GetDefaultUsersOk

`func (o *SurveysPutSurvey) GetDefaultUsersOk() (*interface{}, bool)`

GetDefaultUsersOk returns a tuple with the DefaultUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUsers

`func (o *SurveysPutSurvey) SetDefaultUsers(v interface{})`

SetDefaultUsers sets DefaultUsers field to given value.

### HasDefaultUsers

`func (o *SurveysPutSurvey) HasDefaultUsers() bool`

HasDefaultUsers returns a boolean if a field has been set.

### SetDefaultUsersNil

`func (o *SurveysPutSurvey) SetDefaultUsersNil(b bool)`

 SetDefaultUsersNil sets the value for DefaultUsers to be an explicit nil

### UnsetDefaultUsers
`func (o *SurveysPutSurvey) UnsetDefaultUsers()`

UnsetDefaultUsers ensures that no value is present for DefaultUsers, not even an explicit nil
### GetEmailProjectStart

`func (o *SurveysPutSurvey) GetEmailProjectStart() interface{}`

GetEmailProjectStart returns the EmailProjectStart field if non-nil, zero value otherwise.

### GetEmailProjectStartOk

`func (o *SurveysPutSurvey) GetEmailProjectStartOk() (*interface{}, bool)`

GetEmailProjectStartOk returns a tuple with the EmailProjectStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailProjectStart

`func (o *SurveysPutSurvey) SetEmailProjectStart(v interface{})`

SetEmailProjectStart sets EmailProjectStart field to given value.

### HasEmailProjectStart

`func (o *SurveysPutSurvey) HasEmailProjectStart() bool`

HasEmailProjectStart returns a boolean if a field has been set.

### SetEmailProjectStartNil

`func (o *SurveysPutSurvey) SetEmailProjectStartNil(b bool)`

 SetEmailProjectStartNil sets the value for EmailProjectStart to be an explicit nil

### UnsetEmailProjectStart
`func (o *SurveysPutSurvey) UnsetEmailProjectStart()`

UnsetEmailProjectStart ensures that no value is present for EmailProjectStart, not even an explicit nil
### GetEmailPreparerDigest

`func (o *SurveysPutSurvey) GetEmailPreparerDigest() interface{}`

GetEmailPreparerDigest returns the EmailPreparerDigest field if non-nil, zero value otherwise.

### GetEmailPreparerDigestOk

`func (o *SurveysPutSurvey) GetEmailPreparerDigestOk() (*interface{}, bool)`

GetEmailPreparerDigestOk returns a tuple with the EmailPreparerDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPreparerDigest

`func (o *SurveysPutSurvey) SetEmailPreparerDigest(v interface{})`

SetEmailPreparerDigest sets EmailPreparerDigest field to given value.

### HasEmailPreparerDigest

`func (o *SurveysPutSurvey) HasEmailPreparerDigest() bool`

HasEmailPreparerDigest returns a boolean if a field has been set.

### SetEmailPreparerDigestNil

`func (o *SurveysPutSurvey) SetEmailPreparerDigestNil(b bool)`

 SetEmailPreparerDigestNil sets the value for EmailPreparerDigest to be an explicit nil

### UnsetEmailPreparerDigest
`func (o *SurveysPutSurvey) UnsetEmailPreparerDigest()`

UnsetEmailPreparerDigest ensures that no value is present for EmailPreparerDigest, not even an explicit nil
### GetEmailReviewerDigest

`func (o *SurveysPutSurvey) GetEmailReviewerDigest() interface{}`

GetEmailReviewerDigest returns the EmailReviewerDigest field if non-nil, zero value otherwise.

### GetEmailReviewerDigestOk

`func (o *SurveysPutSurvey) GetEmailReviewerDigestOk() (*interface{}, bool)`

GetEmailReviewerDigestOk returns a tuple with the EmailReviewerDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailReviewerDigest

`func (o *SurveysPutSurvey) SetEmailReviewerDigest(v interface{})`

SetEmailReviewerDigest sets EmailReviewerDigest field to given value.

### HasEmailReviewerDigest

`func (o *SurveysPutSurvey) HasEmailReviewerDigest() bool`

HasEmailReviewerDigest returns a boolean if a field has been set.

### SetEmailReviewerDigestNil

`func (o *SurveysPutSurvey) SetEmailReviewerDigestNil(b bool)`

 SetEmailReviewerDigestNil sets the value for EmailReviewerDigest to be an explicit nil

### UnsetEmailReviewerDigest
`func (o *SurveysPutSurvey) UnsetEmailReviewerDigest()`

UnsetEmailReviewerDigest ensures that no value is present for EmailReviewerDigest, not even an explicit nil
### GetEmailAdminDigest

`func (o *SurveysPutSurvey) GetEmailAdminDigest() interface{}`

GetEmailAdminDigest returns the EmailAdminDigest field if non-nil, zero value otherwise.

### GetEmailAdminDigestOk

`func (o *SurveysPutSurvey) GetEmailAdminDigestOk() (*interface{}, bool)`

GetEmailAdminDigestOk returns a tuple with the EmailAdminDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAdminDigest

`func (o *SurveysPutSurvey) SetEmailAdminDigest(v interface{})`

SetEmailAdminDigest sets EmailAdminDigest field to given value.

### HasEmailAdminDigest

`func (o *SurveysPutSurvey) HasEmailAdminDigest() bool`

HasEmailAdminDigest returns a boolean if a field has been set.

### SetEmailAdminDigestNil

`func (o *SurveysPutSurvey) SetEmailAdminDigestNil(b bool)`

 SetEmailAdminDigestNil sets the value for EmailAdminDigest to be an explicit nil

### UnsetEmailAdminDigest
`func (o *SurveysPutSurvey) UnsetEmailAdminDigest()`

UnsetEmailAdminDigest ensures that no value is present for EmailAdminDigest, not even an explicit nil
### GetEmailSubscriberDigest

`func (o *SurveysPutSurvey) GetEmailSubscriberDigest() interface{}`

GetEmailSubscriberDigest returns the EmailSubscriberDigest field if non-nil, zero value otherwise.

### GetEmailSubscriberDigestOk

`func (o *SurveysPutSurvey) GetEmailSubscriberDigestOk() (*interface{}, bool)`

GetEmailSubscriberDigestOk returns a tuple with the EmailSubscriberDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubscriberDigest

`func (o *SurveysPutSurvey) SetEmailSubscriberDigest(v interface{})`

SetEmailSubscriberDigest sets EmailSubscriberDigest field to given value.

### HasEmailSubscriberDigest

`func (o *SurveysPutSurvey) HasEmailSubscriberDigest() bool`

HasEmailSubscriberDigest returns a boolean if a field has been set.

### SetEmailSubscriberDigestNil

`func (o *SurveysPutSurvey) SetEmailSubscriberDigestNil(b bool)`

 SetEmailSubscriberDigestNil sets the value for EmailSubscriberDigest to be an explicit nil

### UnsetEmailSubscriberDigest
`func (o *SurveysPutSurvey) UnsetEmailSubscriberDigest()`

UnsetEmailSubscriberDigest ensures that no value is present for EmailSubscriberDigest, not even an explicit nil
### GetAdditionalPreparersEnabled

`func (o *SurveysPutSurvey) GetAdditionalPreparersEnabled() bool`

GetAdditionalPreparersEnabled returns the AdditionalPreparersEnabled field if non-nil, zero value otherwise.

### GetAdditionalPreparersEnabledOk

`func (o *SurveysPutSurvey) GetAdditionalPreparersEnabledOk() (*bool, bool)`

GetAdditionalPreparersEnabledOk returns a tuple with the AdditionalPreparersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPreparersEnabled

`func (o *SurveysPutSurvey) SetAdditionalPreparersEnabled(v bool)`

SetAdditionalPreparersEnabled sets AdditionalPreparersEnabled field to given value.

### HasAdditionalPreparersEnabled

`func (o *SurveysPutSurvey) HasAdditionalPreparersEnabled() bool`

HasAdditionalPreparersEnabled returns a boolean if a field has been set.

### GetEmailAdditionalPreparerDigest

`func (o *SurveysPutSurvey) GetEmailAdditionalPreparerDigest() interface{}`

GetEmailAdditionalPreparerDigest returns the EmailAdditionalPreparerDigest field if non-nil, zero value otherwise.

### GetEmailAdditionalPreparerDigestOk

`func (o *SurveysPutSurvey) GetEmailAdditionalPreparerDigestOk() (*interface{}, bool)`

GetEmailAdditionalPreparerDigestOk returns a tuple with the EmailAdditionalPreparerDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAdditionalPreparerDigest

`func (o *SurveysPutSurvey) SetEmailAdditionalPreparerDigest(v interface{})`

SetEmailAdditionalPreparerDigest sets EmailAdditionalPreparerDigest field to given value.

### HasEmailAdditionalPreparerDigest

`func (o *SurveysPutSurvey) HasEmailAdditionalPreparerDigest() bool`

HasEmailAdditionalPreparerDigest returns a boolean if a field has been set.

### SetEmailAdditionalPreparerDigestNil

`func (o *SurveysPutSurvey) SetEmailAdditionalPreparerDigestNil(b bool)`

 SetEmailAdditionalPreparerDigestNil sets the value for EmailAdditionalPreparerDigest to be an explicit nil

### UnsetEmailAdditionalPreparerDigest
`func (o *SurveysPutSurvey) UnsetEmailAdditionalPreparerDigest()`

UnsetEmailAdditionalPreparerDigest ensures that no value is present for EmailAdditionalPreparerDigest, not even an explicit nil
### GetProjectOptions

`func (o *SurveysPutSurvey) GetProjectOptions() interface{}`

GetProjectOptions returns the ProjectOptions field if non-nil, zero value otherwise.

### GetProjectOptionsOk

`func (o *SurveysPutSurvey) GetProjectOptionsOk() (*interface{}, bool)`

GetProjectOptionsOk returns a tuple with the ProjectOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectOptions

`func (o *SurveysPutSurvey) SetProjectOptions(v interface{})`

SetProjectOptions sets ProjectOptions field to given value.

### HasProjectOptions

`func (o *SurveysPutSurvey) HasProjectOptions() bool`

HasProjectOptions returns a boolean if a field has been set.

### SetProjectOptionsNil

`func (o *SurveysPutSurvey) SetProjectOptionsNil(b bool)`

 SetProjectOptionsNil sets the value for ProjectOptions to be an explicit nil

### UnsetProjectOptions
`func (o *SurveysPutSurvey) UnsetProjectOptions()`

UnsetProjectOptions ensures that no value is present for ProjectOptions, not even an explicit nil
### GetAllowYesToAll

`func (o *SurveysPutSurvey) GetAllowYesToAll() bool`

GetAllowYesToAll returns the AllowYesToAll field if non-nil, zero value otherwise.

### GetAllowYesToAllOk

`func (o *SurveysPutSurvey) GetAllowYesToAllOk() (*bool, bool)`

GetAllowYesToAllOk returns a tuple with the AllowYesToAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowYesToAll

`func (o *SurveysPutSurvey) SetAllowYesToAll(v bool)`

SetAllowYesToAll sets AllowYesToAll field to given value.

### HasAllowYesToAll

`func (o *SurveysPutSurvey) HasAllowYesToAll() bool`

HasAllowYesToAll returns a boolean if a field has been set.

### GetPackagingIsStandard

`func (o *SurveysPutSurvey) GetPackagingIsStandard() bool`

GetPackagingIsStandard returns the PackagingIsStandard field if non-nil, zero value otherwise.

### GetPackagingIsStandardOk

`func (o *SurveysPutSurvey) GetPackagingIsStandardOk() (*bool, bool)`

GetPackagingIsStandardOk returns a tuple with the PackagingIsStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagingIsStandard

`func (o *SurveysPutSurvey) SetPackagingIsStandard(v bool)`

SetPackagingIsStandard sets PackagingIsStandard field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


