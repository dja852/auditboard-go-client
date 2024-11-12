# SurveysPutPreviousValues

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
**Status** | Pointer to **string** |  | [optional] [default to "Unlocked"]
**TaskCategoryId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;task_categories.id&#x60;.&lt;fk table&#x3D;&#39;task_categories&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IsArchived** | Pointer to **bool** |  | [optional] [default to false]
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
**PackagingIsStandard** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSurveysPutPreviousValues

`func NewSurveysPutPreviousValues() *SurveysPutPreviousValues`

NewSurveysPutPreviousValues instantiates a new SurveysPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSurveysPutPreviousValuesWithDefaults

`func NewSurveysPutPreviousValuesWithDefaults() *SurveysPutPreviousValues`

NewSurveysPutPreviousValuesWithDefaults instantiates a new SurveysPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SurveysPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SurveysPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SurveysPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SurveysPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SurveysPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SurveysPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SurveysPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SurveysPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetaConfig

`func (o *SurveysPutPreviousValues) GetMetaConfig() string`

GetMetaConfig returns the MetaConfig field if non-nil, zero value otherwise.

### GetMetaConfigOk

`func (o *SurveysPutPreviousValues) GetMetaConfigOk() (*string, bool)`

GetMetaConfigOk returns a tuple with the MetaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaConfig

`func (o *SurveysPutPreviousValues) SetMetaConfig(v string)`

SetMetaConfig sets MetaConfig field to given value.

### HasMetaConfig

`func (o *SurveysPutPreviousValues) HasMetaConfig() bool`

HasMetaConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SurveysPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SurveysPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SurveysPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SurveysPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SurveysPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SurveysPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SurveysPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SurveysPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *SurveysPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SurveysPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SurveysPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *SurveysPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetTaskTitle

`func (o *SurveysPutPreviousValues) GetTaskTitle() string`

GetTaskTitle returns the TaskTitle field if non-nil, zero value otherwise.

### GetTaskTitleOk

`func (o *SurveysPutPreviousValues) GetTaskTitleOk() (*string, bool)`

GetTaskTitleOk returns a tuple with the TaskTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskTitle

`func (o *SurveysPutPreviousValues) SetTaskTitle(v string)`

SetTaskTitle sets TaskTitle field to given value.

### HasTaskTitle

`func (o *SurveysPutPreviousValues) HasTaskTitle() bool`

HasTaskTitle returns a boolean if a field has been set.

### GetType

`func (o *SurveysPutPreviousValues) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SurveysPutPreviousValues) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SurveysPutPreviousValues) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SurveysPutPreviousValues) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *SurveysPutPreviousValues) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SurveysPutPreviousValues) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SurveysPutPreviousValues) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SurveysPutPreviousValues) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskCategoryId

`func (o *SurveysPutPreviousValues) GetTaskCategoryId() int32`

GetTaskCategoryId returns the TaskCategoryId field if non-nil, zero value otherwise.

### GetTaskCategoryIdOk

`func (o *SurveysPutPreviousValues) GetTaskCategoryIdOk() (*int32, bool)`

GetTaskCategoryIdOk returns a tuple with the TaskCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCategoryId

`func (o *SurveysPutPreviousValues) SetTaskCategoryId(v int32)`

SetTaskCategoryId sets TaskCategoryId field to given value.

### HasTaskCategoryId

`func (o *SurveysPutPreviousValues) HasTaskCategoryId() bool`

HasTaskCategoryId returns a boolean if a field has been set.

### GetIsArchived

`func (o *SurveysPutPreviousValues) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *SurveysPutPreviousValues) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *SurveysPutPreviousValues) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *SurveysPutPreviousValues) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetDefaultUsers

`func (o *SurveysPutPreviousValues) GetDefaultUsers() interface{}`

GetDefaultUsers returns the DefaultUsers field if non-nil, zero value otherwise.

### GetDefaultUsersOk

`func (o *SurveysPutPreviousValues) GetDefaultUsersOk() (*interface{}, bool)`

GetDefaultUsersOk returns a tuple with the DefaultUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUsers

`func (o *SurveysPutPreviousValues) SetDefaultUsers(v interface{})`

SetDefaultUsers sets DefaultUsers field to given value.

### HasDefaultUsers

`func (o *SurveysPutPreviousValues) HasDefaultUsers() bool`

HasDefaultUsers returns a boolean if a field has been set.

### SetDefaultUsersNil

`func (o *SurveysPutPreviousValues) SetDefaultUsersNil(b bool)`

 SetDefaultUsersNil sets the value for DefaultUsers to be an explicit nil

### UnsetDefaultUsers
`func (o *SurveysPutPreviousValues) UnsetDefaultUsers()`

UnsetDefaultUsers ensures that no value is present for DefaultUsers, not even an explicit nil
### GetEmailProjectStart

`func (o *SurveysPutPreviousValues) GetEmailProjectStart() interface{}`

GetEmailProjectStart returns the EmailProjectStart field if non-nil, zero value otherwise.

### GetEmailProjectStartOk

`func (o *SurveysPutPreviousValues) GetEmailProjectStartOk() (*interface{}, bool)`

GetEmailProjectStartOk returns a tuple with the EmailProjectStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailProjectStart

`func (o *SurveysPutPreviousValues) SetEmailProjectStart(v interface{})`

SetEmailProjectStart sets EmailProjectStart field to given value.

### HasEmailProjectStart

`func (o *SurveysPutPreviousValues) HasEmailProjectStart() bool`

HasEmailProjectStart returns a boolean if a field has been set.

### SetEmailProjectStartNil

`func (o *SurveysPutPreviousValues) SetEmailProjectStartNil(b bool)`

 SetEmailProjectStartNil sets the value for EmailProjectStart to be an explicit nil

### UnsetEmailProjectStart
`func (o *SurveysPutPreviousValues) UnsetEmailProjectStart()`

UnsetEmailProjectStart ensures that no value is present for EmailProjectStart, not even an explicit nil
### GetEmailPreparerDigest

`func (o *SurveysPutPreviousValues) GetEmailPreparerDigest() interface{}`

GetEmailPreparerDigest returns the EmailPreparerDigest field if non-nil, zero value otherwise.

### GetEmailPreparerDigestOk

`func (o *SurveysPutPreviousValues) GetEmailPreparerDigestOk() (*interface{}, bool)`

GetEmailPreparerDigestOk returns a tuple with the EmailPreparerDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPreparerDigest

`func (o *SurveysPutPreviousValues) SetEmailPreparerDigest(v interface{})`

SetEmailPreparerDigest sets EmailPreparerDigest field to given value.

### HasEmailPreparerDigest

`func (o *SurveysPutPreviousValues) HasEmailPreparerDigest() bool`

HasEmailPreparerDigest returns a boolean if a field has been set.

### SetEmailPreparerDigestNil

`func (o *SurveysPutPreviousValues) SetEmailPreparerDigestNil(b bool)`

 SetEmailPreparerDigestNil sets the value for EmailPreparerDigest to be an explicit nil

### UnsetEmailPreparerDigest
`func (o *SurveysPutPreviousValues) UnsetEmailPreparerDigest()`

UnsetEmailPreparerDigest ensures that no value is present for EmailPreparerDigest, not even an explicit nil
### GetEmailReviewerDigest

`func (o *SurveysPutPreviousValues) GetEmailReviewerDigest() interface{}`

GetEmailReviewerDigest returns the EmailReviewerDigest field if non-nil, zero value otherwise.

### GetEmailReviewerDigestOk

`func (o *SurveysPutPreviousValues) GetEmailReviewerDigestOk() (*interface{}, bool)`

GetEmailReviewerDigestOk returns a tuple with the EmailReviewerDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailReviewerDigest

`func (o *SurveysPutPreviousValues) SetEmailReviewerDigest(v interface{})`

SetEmailReviewerDigest sets EmailReviewerDigest field to given value.

### HasEmailReviewerDigest

`func (o *SurveysPutPreviousValues) HasEmailReviewerDigest() bool`

HasEmailReviewerDigest returns a boolean if a field has been set.

### SetEmailReviewerDigestNil

`func (o *SurveysPutPreviousValues) SetEmailReviewerDigestNil(b bool)`

 SetEmailReviewerDigestNil sets the value for EmailReviewerDigest to be an explicit nil

### UnsetEmailReviewerDigest
`func (o *SurveysPutPreviousValues) UnsetEmailReviewerDigest()`

UnsetEmailReviewerDigest ensures that no value is present for EmailReviewerDigest, not even an explicit nil
### GetEmailAdminDigest

`func (o *SurveysPutPreviousValues) GetEmailAdminDigest() interface{}`

GetEmailAdminDigest returns the EmailAdminDigest field if non-nil, zero value otherwise.

### GetEmailAdminDigestOk

`func (o *SurveysPutPreviousValues) GetEmailAdminDigestOk() (*interface{}, bool)`

GetEmailAdminDigestOk returns a tuple with the EmailAdminDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAdminDigest

`func (o *SurveysPutPreviousValues) SetEmailAdminDigest(v interface{})`

SetEmailAdminDigest sets EmailAdminDigest field to given value.

### HasEmailAdminDigest

`func (o *SurveysPutPreviousValues) HasEmailAdminDigest() bool`

HasEmailAdminDigest returns a boolean if a field has been set.

### SetEmailAdminDigestNil

`func (o *SurveysPutPreviousValues) SetEmailAdminDigestNil(b bool)`

 SetEmailAdminDigestNil sets the value for EmailAdminDigest to be an explicit nil

### UnsetEmailAdminDigest
`func (o *SurveysPutPreviousValues) UnsetEmailAdminDigest()`

UnsetEmailAdminDigest ensures that no value is present for EmailAdminDigest, not even an explicit nil
### GetEmailSubscriberDigest

`func (o *SurveysPutPreviousValues) GetEmailSubscriberDigest() interface{}`

GetEmailSubscriberDigest returns the EmailSubscriberDigest field if non-nil, zero value otherwise.

### GetEmailSubscriberDigestOk

`func (o *SurveysPutPreviousValues) GetEmailSubscriberDigestOk() (*interface{}, bool)`

GetEmailSubscriberDigestOk returns a tuple with the EmailSubscriberDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubscriberDigest

`func (o *SurveysPutPreviousValues) SetEmailSubscriberDigest(v interface{})`

SetEmailSubscriberDigest sets EmailSubscriberDigest field to given value.

### HasEmailSubscriberDigest

`func (o *SurveysPutPreviousValues) HasEmailSubscriberDigest() bool`

HasEmailSubscriberDigest returns a boolean if a field has been set.

### SetEmailSubscriberDigestNil

`func (o *SurveysPutPreviousValues) SetEmailSubscriberDigestNil(b bool)`

 SetEmailSubscriberDigestNil sets the value for EmailSubscriberDigest to be an explicit nil

### UnsetEmailSubscriberDigest
`func (o *SurveysPutPreviousValues) UnsetEmailSubscriberDigest()`

UnsetEmailSubscriberDigest ensures that no value is present for EmailSubscriberDigest, not even an explicit nil
### GetAdditionalPreparersEnabled

`func (o *SurveysPutPreviousValues) GetAdditionalPreparersEnabled() bool`

GetAdditionalPreparersEnabled returns the AdditionalPreparersEnabled field if non-nil, zero value otherwise.

### GetAdditionalPreparersEnabledOk

`func (o *SurveysPutPreviousValues) GetAdditionalPreparersEnabledOk() (*bool, bool)`

GetAdditionalPreparersEnabledOk returns a tuple with the AdditionalPreparersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPreparersEnabled

`func (o *SurveysPutPreviousValues) SetAdditionalPreparersEnabled(v bool)`

SetAdditionalPreparersEnabled sets AdditionalPreparersEnabled field to given value.

### HasAdditionalPreparersEnabled

`func (o *SurveysPutPreviousValues) HasAdditionalPreparersEnabled() bool`

HasAdditionalPreparersEnabled returns a boolean if a field has been set.

### GetEmailAdditionalPreparerDigest

`func (o *SurveysPutPreviousValues) GetEmailAdditionalPreparerDigest() interface{}`

GetEmailAdditionalPreparerDigest returns the EmailAdditionalPreparerDigest field if non-nil, zero value otherwise.

### GetEmailAdditionalPreparerDigestOk

`func (o *SurveysPutPreviousValues) GetEmailAdditionalPreparerDigestOk() (*interface{}, bool)`

GetEmailAdditionalPreparerDigestOk returns a tuple with the EmailAdditionalPreparerDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAdditionalPreparerDigest

`func (o *SurveysPutPreviousValues) SetEmailAdditionalPreparerDigest(v interface{})`

SetEmailAdditionalPreparerDigest sets EmailAdditionalPreparerDigest field to given value.

### HasEmailAdditionalPreparerDigest

`func (o *SurveysPutPreviousValues) HasEmailAdditionalPreparerDigest() bool`

HasEmailAdditionalPreparerDigest returns a boolean if a field has been set.

### SetEmailAdditionalPreparerDigestNil

`func (o *SurveysPutPreviousValues) SetEmailAdditionalPreparerDigestNil(b bool)`

 SetEmailAdditionalPreparerDigestNil sets the value for EmailAdditionalPreparerDigest to be an explicit nil

### UnsetEmailAdditionalPreparerDigest
`func (o *SurveysPutPreviousValues) UnsetEmailAdditionalPreparerDigest()`

UnsetEmailAdditionalPreparerDigest ensures that no value is present for EmailAdditionalPreparerDigest, not even an explicit nil
### GetProjectOptions

`func (o *SurveysPutPreviousValues) GetProjectOptions() interface{}`

GetProjectOptions returns the ProjectOptions field if non-nil, zero value otherwise.

### GetProjectOptionsOk

`func (o *SurveysPutPreviousValues) GetProjectOptionsOk() (*interface{}, bool)`

GetProjectOptionsOk returns a tuple with the ProjectOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectOptions

`func (o *SurveysPutPreviousValues) SetProjectOptions(v interface{})`

SetProjectOptions sets ProjectOptions field to given value.

### HasProjectOptions

`func (o *SurveysPutPreviousValues) HasProjectOptions() bool`

HasProjectOptions returns a boolean if a field has been set.

### SetProjectOptionsNil

`func (o *SurveysPutPreviousValues) SetProjectOptionsNil(b bool)`

 SetProjectOptionsNil sets the value for ProjectOptions to be an explicit nil

### UnsetProjectOptions
`func (o *SurveysPutPreviousValues) UnsetProjectOptions()`

UnsetProjectOptions ensures that no value is present for ProjectOptions, not even an explicit nil
### GetAllowYesToAll

`func (o *SurveysPutPreviousValues) GetAllowYesToAll() bool`

GetAllowYesToAll returns the AllowYesToAll field if non-nil, zero value otherwise.

### GetAllowYesToAllOk

`func (o *SurveysPutPreviousValues) GetAllowYesToAllOk() (*bool, bool)`

GetAllowYesToAllOk returns a tuple with the AllowYesToAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowYesToAll

`func (o *SurveysPutPreviousValues) SetAllowYesToAll(v bool)`

SetAllowYesToAll sets AllowYesToAll field to given value.

### HasAllowYesToAll

`func (o *SurveysPutPreviousValues) HasAllowYesToAll() bool`

HasAllowYesToAll returns a boolean if a field has been set.

### GetPackagingIsStandard

`func (o *SurveysPutPreviousValues) GetPackagingIsStandard() bool`

GetPackagingIsStandard returns the PackagingIsStandard field if non-nil, zero value otherwise.

### GetPackagingIsStandardOk

`func (o *SurveysPutPreviousValues) GetPackagingIsStandardOk() (*bool, bool)`

GetPackagingIsStandardOk returns a tuple with the PackagingIsStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagingIsStandard

`func (o *SurveysPutPreviousValues) SetPackagingIsStandard(v bool)`

SetPackagingIsStandard sets PackagingIsStandard field to given value.

### HasPackagingIsStandard

`func (o *SurveysPutPreviousValues) HasPackagingIsStandard() bool`

HasPackagingIsStandard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


