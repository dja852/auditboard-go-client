# TaskGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**CreatorUserId** | **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | 
**TaskCategoryId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;task_categories.id&#x60;.&lt;fk table&#x3D;&#39;task_categories&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ArchivedAt** | Pointer to **string** |  | [optional] 
**TaskPeriodId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;task_periods.id&#x60;.&lt;fk table&#x3D;&#39;task_periods&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**SurveyId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;surveys.id&#x60;.&lt;fk table&#x3D;&#39;surveys&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**AllowYesToAll** | **bool** |  | [default to false]
**TestSectionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_sections.id&#x60;.&lt;fk table&#x3D;&#39;test_sections&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**EmailProjectStart** | Pointer to **interface{}** |  | [optional] 
**EmailPreparerDigest** | Pointer to **interface{}** |  | [optional] 
**EmailReviewerDigest** | Pointer to **interface{}** |  | [optional] 
**EmailAdminDigest** | Pointer to **interface{}** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**ReviewerUsers** | Pointer to **interface{}** |  | [optional] 
**AssigneeUsers** | Pointer to **interface{}** |  | [optional] 
**IsMultiReviewer** | Pointer to **bool** |  | [optional] [default to false]
**ProjectOptions** | Pointer to **interface{}** |  | [optional] 
**IsManagedChange** | Pointer to **bool** |  | [optional] [default to false]
**SubscriberUsers** | Pointer to **interface{}** |  | [optional] 
**EmailSubscriberDigest** | Pointer to **interface{}** |  | [optional] 
**AssessmentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;assessments.id&#x60;.&lt;fk table&#x3D;&#39;assessments&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AdditionalAssigneeUsers** | Pointer to **interface{}** |  | [optional] 
**AdditionalPreparersEnabled** | Pointer to **bool** |  | [optional] [default to false]
**EmailAdditionalPreparerDigest** | Pointer to **interface{}** |  | [optional] 
**HistoricalProjectId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;task_groups.id&#x60;.&lt;fk table&#x3D;&#39;task_groups&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Scopes** | **interface{}** |  | 
**ScheduledStartDate** | Pointer to **string** |  | [optional] 
**Totals** | Pointer to **interface{}** |  | [optional] 
**ShortTitle** | Pointer to **string** |  | [optional] 
**RecurringTaskGroupId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;recurring_task_groups.id&#x60;.&lt;fk table&#x3D;&#39;recurring_task_groups&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AuditableEntityId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;auditable_entities.id&#x60;.&lt;fk table&#x3D;&#39;auditable_entities&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**MultiReviewerMode** | **string** |  | [default to "sequential"]
**RequiredReviewerCount** | Pointer to **int32** |  | [optional] 
**UserTotals** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewTaskGroups

`func NewTaskGroups(creatorUserId int32, allowYesToAll bool, scopes interface{}, multiReviewerMode string, ) *TaskGroups`

NewTaskGroups instantiates a new TaskGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskGroupsWithDefaults

`func NewTaskGroupsWithDefaults() *TaskGroups`

NewTaskGroupsWithDefaults instantiates a new TaskGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskGroups) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskGroups) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskGroups) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TaskGroups) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskGroups) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskGroups) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskGroups) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaskGroups) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TaskGroups) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskGroups) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskGroups) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TaskGroups) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *TaskGroups) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *TaskGroups) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *TaskGroups) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *TaskGroups) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetCreatorUserId

`func (o *TaskGroups) GetCreatorUserId() int32`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *TaskGroups) GetCreatorUserIdOk() (*int32, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *TaskGroups) SetCreatorUserId(v int32)`

SetCreatorUserId sets CreatorUserId field to given value.


### GetTaskCategoryId

`func (o *TaskGroups) GetTaskCategoryId() int32`

GetTaskCategoryId returns the TaskCategoryId field if non-nil, zero value otherwise.

### GetTaskCategoryIdOk

`func (o *TaskGroups) GetTaskCategoryIdOk() (*int32, bool)`

GetTaskCategoryIdOk returns a tuple with the TaskCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCategoryId

`func (o *TaskGroups) SetTaskCategoryId(v int32)`

SetTaskCategoryId sets TaskCategoryId field to given value.

### HasTaskCategoryId

`func (o *TaskGroups) HasTaskCategoryId() bool`

HasTaskCategoryId returns a boolean if a field has been set.

### GetArchivedAt

`func (o *TaskGroups) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *TaskGroups) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *TaskGroups) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *TaskGroups) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetTaskPeriodId

`func (o *TaskGroups) GetTaskPeriodId() int32`

GetTaskPeriodId returns the TaskPeriodId field if non-nil, zero value otherwise.

### GetTaskPeriodIdOk

`func (o *TaskGroups) GetTaskPeriodIdOk() (*int32, bool)`

GetTaskPeriodIdOk returns a tuple with the TaskPeriodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPeriodId

`func (o *TaskGroups) SetTaskPeriodId(v int32)`

SetTaskPeriodId sets TaskPeriodId field to given value.

### HasTaskPeriodId

`func (o *TaskGroups) HasTaskPeriodId() bool`

HasTaskPeriodId returns a boolean if a field has been set.

### GetSurveyId

`func (o *TaskGroups) GetSurveyId() int32`

GetSurveyId returns the SurveyId field if non-nil, zero value otherwise.

### GetSurveyIdOk

`func (o *TaskGroups) GetSurveyIdOk() (*int32, bool)`

GetSurveyIdOk returns a tuple with the SurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyId

`func (o *TaskGroups) SetSurveyId(v int32)`

SetSurveyId sets SurveyId field to given value.

### HasSurveyId

`func (o *TaskGroups) HasSurveyId() bool`

HasSurveyId returns a boolean if a field has been set.

### GetStartedAt

`func (o *TaskGroups) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *TaskGroups) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *TaskGroups) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *TaskGroups) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *TaskGroups) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskGroups) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskGroups) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskGroups) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *TaskGroups) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TaskGroups) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TaskGroups) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TaskGroups) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAllowYesToAll

`func (o *TaskGroups) GetAllowYesToAll() bool`

GetAllowYesToAll returns the AllowYesToAll field if non-nil, zero value otherwise.

### GetAllowYesToAllOk

`func (o *TaskGroups) GetAllowYesToAllOk() (*bool, bool)`

GetAllowYesToAllOk returns a tuple with the AllowYesToAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowYesToAll

`func (o *TaskGroups) SetAllowYesToAll(v bool)`

SetAllowYesToAll sets AllowYesToAll field to given value.


### GetTestSectionId

`func (o *TaskGroups) GetTestSectionId() int32`

GetTestSectionId returns the TestSectionId field if non-nil, zero value otherwise.

### GetTestSectionIdOk

`func (o *TaskGroups) GetTestSectionIdOk() (*int32, bool)`

GetTestSectionIdOk returns a tuple with the TestSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSectionId

`func (o *TaskGroups) SetTestSectionId(v int32)`

SetTestSectionId sets TestSectionId field to given value.

### HasTestSectionId

`func (o *TaskGroups) HasTestSectionId() bool`

HasTestSectionId returns a boolean if a field has been set.

### GetEmailProjectStart

`func (o *TaskGroups) GetEmailProjectStart() interface{}`

GetEmailProjectStart returns the EmailProjectStart field if non-nil, zero value otherwise.

### GetEmailProjectStartOk

`func (o *TaskGroups) GetEmailProjectStartOk() (*interface{}, bool)`

GetEmailProjectStartOk returns a tuple with the EmailProjectStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailProjectStart

`func (o *TaskGroups) SetEmailProjectStart(v interface{})`

SetEmailProjectStart sets EmailProjectStart field to given value.

### HasEmailProjectStart

`func (o *TaskGroups) HasEmailProjectStart() bool`

HasEmailProjectStart returns a boolean if a field has been set.

### SetEmailProjectStartNil

`func (o *TaskGroups) SetEmailProjectStartNil(b bool)`

 SetEmailProjectStartNil sets the value for EmailProjectStart to be an explicit nil

### UnsetEmailProjectStart
`func (o *TaskGroups) UnsetEmailProjectStart()`

UnsetEmailProjectStart ensures that no value is present for EmailProjectStart, not even an explicit nil
### GetEmailPreparerDigest

`func (o *TaskGroups) GetEmailPreparerDigest() interface{}`

GetEmailPreparerDigest returns the EmailPreparerDigest field if non-nil, zero value otherwise.

### GetEmailPreparerDigestOk

`func (o *TaskGroups) GetEmailPreparerDigestOk() (*interface{}, bool)`

GetEmailPreparerDigestOk returns a tuple with the EmailPreparerDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPreparerDigest

`func (o *TaskGroups) SetEmailPreparerDigest(v interface{})`

SetEmailPreparerDigest sets EmailPreparerDigest field to given value.

### HasEmailPreparerDigest

`func (o *TaskGroups) HasEmailPreparerDigest() bool`

HasEmailPreparerDigest returns a boolean if a field has been set.

### SetEmailPreparerDigestNil

`func (o *TaskGroups) SetEmailPreparerDigestNil(b bool)`

 SetEmailPreparerDigestNil sets the value for EmailPreparerDigest to be an explicit nil

### UnsetEmailPreparerDigest
`func (o *TaskGroups) UnsetEmailPreparerDigest()`

UnsetEmailPreparerDigest ensures that no value is present for EmailPreparerDigest, not even an explicit nil
### GetEmailReviewerDigest

`func (o *TaskGroups) GetEmailReviewerDigest() interface{}`

GetEmailReviewerDigest returns the EmailReviewerDigest field if non-nil, zero value otherwise.

### GetEmailReviewerDigestOk

`func (o *TaskGroups) GetEmailReviewerDigestOk() (*interface{}, bool)`

GetEmailReviewerDigestOk returns a tuple with the EmailReviewerDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailReviewerDigest

`func (o *TaskGroups) SetEmailReviewerDigest(v interface{})`

SetEmailReviewerDigest sets EmailReviewerDigest field to given value.

### HasEmailReviewerDigest

`func (o *TaskGroups) HasEmailReviewerDigest() bool`

HasEmailReviewerDigest returns a boolean if a field has been set.

### SetEmailReviewerDigestNil

`func (o *TaskGroups) SetEmailReviewerDigestNil(b bool)`

 SetEmailReviewerDigestNil sets the value for EmailReviewerDigest to be an explicit nil

### UnsetEmailReviewerDigest
`func (o *TaskGroups) UnsetEmailReviewerDigest()`

UnsetEmailReviewerDigest ensures that no value is present for EmailReviewerDigest, not even an explicit nil
### GetEmailAdminDigest

`func (o *TaskGroups) GetEmailAdminDigest() interface{}`

GetEmailAdminDigest returns the EmailAdminDigest field if non-nil, zero value otherwise.

### GetEmailAdminDigestOk

`func (o *TaskGroups) GetEmailAdminDigestOk() (*interface{}, bool)`

GetEmailAdminDigestOk returns a tuple with the EmailAdminDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAdminDigest

`func (o *TaskGroups) SetEmailAdminDigest(v interface{})`

SetEmailAdminDigest sets EmailAdminDigest field to given value.

### HasEmailAdminDigest

`func (o *TaskGroups) HasEmailAdminDigest() bool`

HasEmailAdminDigest returns a boolean if a field has been set.

### SetEmailAdminDigestNil

`func (o *TaskGroups) SetEmailAdminDigestNil(b bool)`

 SetEmailAdminDigestNil sets the value for EmailAdminDigest to be an explicit nil

### UnsetEmailAdminDigest
`func (o *TaskGroups) UnsetEmailAdminDigest()`

UnsetEmailAdminDigest ensures that no value is present for EmailAdminDigest, not even an explicit nil
### GetUid

`func (o *TaskGroups) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *TaskGroups) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *TaskGroups) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *TaskGroups) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetReviewerUsers

`func (o *TaskGroups) GetReviewerUsers() interface{}`

GetReviewerUsers returns the ReviewerUsers field if non-nil, zero value otherwise.

### GetReviewerUsersOk

`func (o *TaskGroups) GetReviewerUsersOk() (*interface{}, bool)`

GetReviewerUsersOk returns a tuple with the ReviewerUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerUsers

`func (o *TaskGroups) SetReviewerUsers(v interface{})`

SetReviewerUsers sets ReviewerUsers field to given value.

### HasReviewerUsers

`func (o *TaskGroups) HasReviewerUsers() bool`

HasReviewerUsers returns a boolean if a field has been set.

### SetReviewerUsersNil

`func (o *TaskGroups) SetReviewerUsersNil(b bool)`

 SetReviewerUsersNil sets the value for ReviewerUsers to be an explicit nil

### UnsetReviewerUsers
`func (o *TaskGroups) UnsetReviewerUsers()`

UnsetReviewerUsers ensures that no value is present for ReviewerUsers, not even an explicit nil
### GetAssigneeUsers

`func (o *TaskGroups) GetAssigneeUsers() interface{}`

GetAssigneeUsers returns the AssigneeUsers field if non-nil, zero value otherwise.

### GetAssigneeUsersOk

`func (o *TaskGroups) GetAssigneeUsersOk() (*interface{}, bool)`

GetAssigneeUsersOk returns a tuple with the AssigneeUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeUsers

`func (o *TaskGroups) SetAssigneeUsers(v interface{})`

SetAssigneeUsers sets AssigneeUsers field to given value.

### HasAssigneeUsers

`func (o *TaskGroups) HasAssigneeUsers() bool`

HasAssigneeUsers returns a boolean if a field has been set.

### SetAssigneeUsersNil

`func (o *TaskGroups) SetAssigneeUsersNil(b bool)`

 SetAssigneeUsersNil sets the value for AssigneeUsers to be an explicit nil

### UnsetAssigneeUsers
`func (o *TaskGroups) UnsetAssigneeUsers()`

UnsetAssigneeUsers ensures that no value is present for AssigneeUsers, not even an explicit nil
### GetIsMultiReviewer

`func (o *TaskGroups) GetIsMultiReviewer() bool`

GetIsMultiReviewer returns the IsMultiReviewer field if non-nil, zero value otherwise.

### GetIsMultiReviewerOk

`func (o *TaskGroups) GetIsMultiReviewerOk() (*bool, bool)`

GetIsMultiReviewerOk returns a tuple with the IsMultiReviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiReviewer

`func (o *TaskGroups) SetIsMultiReviewer(v bool)`

SetIsMultiReviewer sets IsMultiReviewer field to given value.

### HasIsMultiReviewer

`func (o *TaskGroups) HasIsMultiReviewer() bool`

HasIsMultiReviewer returns a boolean if a field has been set.

### GetProjectOptions

`func (o *TaskGroups) GetProjectOptions() interface{}`

GetProjectOptions returns the ProjectOptions field if non-nil, zero value otherwise.

### GetProjectOptionsOk

`func (o *TaskGroups) GetProjectOptionsOk() (*interface{}, bool)`

GetProjectOptionsOk returns a tuple with the ProjectOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectOptions

`func (o *TaskGroups) SetProjectOptions(v interface{})`

SetProjectOptions sets ProjectOptions field to given value.

### HasProjectOptions

`func (o *TaskGroups) HasProjectOptions() bool`

HasProjectOptions returns a boolean if a field has been set.

### SetProjectOptionsNil

`func (o *TaskGroups) SetProjectOptionsNil(b bool)`

 SetProjectOptionsNil sets the value for ProjectOptions to be an explicit nil

### UnsetProjectOptions
`func (o *TaskGroups) UnsetProjectOptions()`

UnsetProjectOptions ensures that no value is present for ProjectOptions, not even an explicit nil
### GetIsManagedChange

`func (o *TaskGroups) GetIsManagedChange() bool`

GetIsManagedChange returns the IsManagedChange field if non-nil, zero value otherwise.

### GetIsManagedChangeOk

`func (o *TaskGroups) GetIsManagedChangeOk() (*bool, bool)`

GetIsManagedChangeOk returns a tuple with the IsManagedChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManagedChange

`func (o *TaskGroups) SetIsManagedChange(v bool)`

SetIsManagedChange sets IsManagedChange field to given value.

### HasIsManagedChange

`func (o *TaskGroups) HasIsManagedChange() bool`

HasIsManagedChange returns a boolean if a field has been set.

### GetSubscriberUsers

`func (o *TaskGroups) GetSubscriberUsers() interface{}`

GetSubscriberUsers returns the SubscriberUsers field if non-nil, zero value otherwise.

### GetSubscriberUsersOk

`func (o *TaskGroups) GetSubscriberUsersOk() (*interface{}, bool)`

GetSubscriberUsersOk returns a tuple with the SubscriberUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberUsers

`func (o *TaskGroups) SetSubscriberUsers(v interface{})`

SetSubscriberUsers sets SubscriberUsers field to given value.

### HasSubscriberUsers

`func (o *TaskGroups) HasSubscriberUsers() bool`

HasSubscriberUsers returns a boolean if a field has been set.

### SetSubscriberUsersNil

`func (o *TaskGroups) SetSubscriberUsersNil(b bool)`

 SetSubscriberUsersNil sets the value for SubscriberUsers to be an explicit nil

### UnsetSubscriberUsers
`func (o *TaskGroups) UnsetSubscriberUsers()`

UnsetSubscriberUsers ensures that no value is present for SubscriberUsers, not even an explicit nil
### GetEmailSubscriberDigest

`func (o *TaskGroups) GetEmailSubscriberDigest() interface{}`

GetEmailSubscriberDigest returns the EmailSubscriberDigest field if non-nil, zero value otherwise.

### GetEmailSubscriberDigestOk

`func (o *TaskGroups) GetEmailSubscriberDigestOk() (*interface{}, bool)`

GetEmailSubscriberDigestOk returns a tuple with the EmailSubscriberDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubscriberDigest

`func (o *TaskGroups) SetEmailSubscriberDigest(v interface{})`

SetEmailSubscriberDigest sets EmailSubscriberDigest field to given value.

### HasEmailSubscriberDigest

`func (o *TaskGroups) HasEmailSubscriberDigest() bool`

HasEmailSubscriberDigest returns a boolean if a field has been set.

### SetEmailSubscriberDigestNil

`func (o *TaskGroups) SetEmailSubscriberDigestNil(b bool)`

 SetEmailSubscriberDigestNil sets the value for EmailSubscriberDigest to be an explicit nil

### UnsetEmailSubscriberDigest
`func (o *TaskGroups) UnsetEmailSubscriberDigest()`

UnsetEmailSubscriberDigest ensures that no value is present for EmailSubscriberDigest, not even an explicit nil
### GetAssessmentId

`func (o *TaskGroups) GetAssessmentId() int32`

GetAssessmentId returns the AssessmentId field if non-nil, zero value otherwise.

### GetAssessmentIdOk

`func (o *TaskGroups) GetAssessmentIdOk() (*int32, bool)`

GetAssessmentIdOk returns a tuple with the AssessmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessmentId

`func (o *TaskGroups) SetAssessmentId(v int32)`

SetAssessmentId sets AssessmentId field to given value.

### HasAssessmentId

`func (o *TaskGroups) HasAssessmentId() bool`

HasAssessmentId returns a boolean if a field has been set.

### GetAdditionalAssigneeUsers

`func (o *TaskGroups) GetAdditionalAssigneeUsers() interface{}`

GetAdditionalAssigneeUsers returns the AdditionalAssigneeUsers field if non-nil, zero value otherwise.

### GetAdditionalAssigneeUsersOk

`func (o *TaskGroups) GetAdditionalAssigneeUsersOk() (*interface{}, bool)`

GetAdditionalAssigneeUsersOk returns a tuple with the AdditionalAssigneeUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAssigneeUsers

`func (o *TaskGroups) SetAdditionalAssigneeUsers(v interface{})`

SetAdditionalAssigneeUsers sets AdditionalAssigneeUsers field to given value.

### HasAdditionalAssigneeUsers

`func (o *TaskGroups) HasAdditionalAssigneeUsers() bool`

HasAdditionalAssigneeUsers returns a boolean if a field has been set.

### SetAdditionalAssigneeUsersNil

`func (o *TaskGroups) SetAdditionalAssigneeUsersNil(b bool)`

 SetAdditionalAssigneeUsersNil sets the value for AdditionalAssigneeUsers to be an explicit nil

### UnsetAdditionalAssigneeUsers
`func (o *TaskGroups) UnsetAdditionalAssigneeUsers()`

UnsetAdditionalAssigneeUsers ensures that no value is present for AdditionalAssigneeUsers, not even an explicit nil
### GetAdditionalPreparersEnabled

`func (o *TaskGroups) GetAdditionalPreparersEnabled() bool`

GetAdditionalPreparersEnabled returns the AdditionalPreparersEnabled field if non-nil, zero value otherwise.

### GetAdditionalPreparersEnabledOk

`func (o *TaskGroups) GetAdditionalPreparersEnabledOk() (*bool, bool)`

GetAdditionalPreparersEnabledOk returns a tuple with the AdditionalPreparersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPreparersEnabled

`func (o *TaskGroups) SetAdditionalPreparersEnabled(v bool)`

SetAdditionalPreparersEnabled sets AdditionalPreparersEnabled field to given value.

### HasAdditionalPreparersEnabled

`func (o *TaskGroups) HasAdditionalPreparersEnabled() bool`

HasAdditionalPreparersEnabled returns a boolean if a field has been set.

### GetEmailAdditionalPreparerDigest

`func (o *TaskGroups) GetEmailAdditionalPreparerDigest() interface{}`

GetEmailAdditionalPreparerDigest returns the EmailAdditionalPreparerDigest field if non-nil, zero value otherwise.

### GetEmailAdditionalPreparerDigestOk

`func (o *TaskGroups) GetEmailAdditionalPreparerDigestOk() (*interface{}, bool)`

GetEmailAdditionalPreparerDigestOk returns a tuple with the EmailAdditionalPreparerDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAdditionalPreparerDigest

`func (o *TaskGroups) SetEmailAdditionalPreparerDigest(v interface{})`

SetEmailAdditionalPreparerDigest sets EmailAdditionalPreparerDigest field to given value.

### HasEmailAdditionalPreparerDigest

`func (o *TaskGroups) HasEmailAdditionalPreparerDigest() bool`

HasEmailAdditionalPreparerDigest returns a boolean if a field has been set.

### SetEmailAdditionalPreparerDigestNil

`func (o *TaskGroups) SetEmailAdditionalPreparerDigestNil(b bool)`

 SetEmailAdditionalPreparerDigestNil sets the value for EmailAdditionalPreparerDigest to be an explicit nil

### UnsetEmailAdditionalPreparerDigest
`func (o *TaskGroups) UnsetEmailAdditionalPreparerDigest()`

UnsetEmailAdditionalPreparerDigest ensures that no value is present for EmailAdditionalPreparerDigest, not even an explicit nil
### GetHistoricalProjectId

`func (o *TaskGroups) GetHistoricalProjectId() int32`

GetHistoricalProjectId returns the HistoricalProjectId field if non-nil, zero value otherwise.

### GetHistoricalProjectIdOk

`func (o *TaskGroups) GetHistoricalProjectIdOk() (*int32, bool)`

GetHistoricalProjectIdOk returns a tuple with the HistoricalProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricalProjectId

`func (o *TaskGroups) SetHistoricalProjectId(v int32)`

SetHistoricalProjectId sets HistoricalProjectId field to given value.

### HasHistoricalProjectId

`func (o *TaskGroups) HasHistoricalProjectId() bool`

HasHistoricalProjectId returns a boolean if a field has been set.

### GetScopes

`func (o *TaskGroups) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TaskGroups) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TaskGroups) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.


### SetScopesNil

`func (o *TaskGroups) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *TaskGroups) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetScheduledStartDate

`func (o *TaskGroups) GetScheduledStartDate() string`

GetScheduledStartDate returns the ScheduledStartDate field if non-nil, zero value otherwise.

### GetScheduledStartDateOk

`func (o *TaskGroups) GetScheduledStartDateOk() (*string, bool)`

GetScheduledStartDateOk returns a tuple with the ScheduledStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartDate

`func (o *TaskGroups) SetScheduledStartDate(v string)`

SetScheduledStartDate sets ScheduledStartDate field to given value.

### HasScheduledStartDate

`func (o *TaskGroups) HasScheduledStartDate() bool`

HasScheduledStartDate returns a boolean if a field has been set.

### GetTotals

`func (o *TaskGroups) GetTotals() interface{}`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *TaskGroups) GetTotalsOk() (*interface{}, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *TaskGroups) SetTotals(v interface{})`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *TaskGroups) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### SetTotalsNil

`func (o *TaskGroups) SetTotalsNil(b bool)`

 SetTotalsNil sets the value for Totals to be an explicit nil

### UnsetTotals
`func (o *TaskGroups) UnsetTotals()`

UnsetTotals ensures that no value is present for Totals, not even an explicit nil
### GetShortTitle

`func (o *TaskGroups) GetShortTitle() string`

GetShortTitle returns the ShortTitle field if non-nil, zero value otherwise.

### GetShortTitleOk

`func (o *TaskGroups) GetShortTitleOk() (*string, bool)`

GetShortTitleOk returns a tuple with the ShortTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortTitle

`func (o *TaskGroups) SetShortTitle(v string)`

SetShortTitle sets ShortTitle field to given value.

### HasShortTitle

`func (o *TaskGroups) HasShortTitle() bool`

HasShortTitle returns a boolean if a field has been set.

### GetRecurringTaskGroupId

`func (o *TaskGroups) GetRecurringTaskGroupId() int32`

GetRecurringTaskGroupId returns the RecurringTaskGroupId field if non-nil, zero value otherwise.

### GetRecurringTaskGroupIdOk

`func (o *TaskGroups) GetRecurringTaskGroupIdOk() (*int32, bool)`

GetRecurringTaskGroupIdOk returns a tuple with the RecurringTaskGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringTaskGroupId

`func (o *TaskGroups) SetRecurringTaskGroupId(v int32)`

SetRecurringTaskGroupId sets RecurringTaskGroupId field to given value.

### HasRecurringTaskGroupId

`func (o *TaskGroups) HasRecurringTaskGroupId() bool`

HasRecurringTaskGroupId returns a boolean if a field has been set.

### GetAuditableEntityId

`func (o *TaskGroups) GetAuditableEntityId() int32`

GetAuditableEntityId returns the AuditableEntityId field if non-nil, zero value otherwise.

### GetAuditableEntityIdOk

`func (o *TaskGroups) GetAuditableEntityIdOk() (*int32, bool)`

GetAuditableEntityIdOk returns a tuple with the AuditableEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditableEntityId

`func (o *TaskGroups) SetAuditableEntityId(v int32)`

SetAuditableEntityId sets AuditableEntityId field to given value.

### HasAuditableEntityId

`func (o *TaskGroups) HasAuditableEntityId() bool`

HasAuditableEntityId returns a boolean if a field has been set.

### GetMultiReviewerMode

`func (o *TaskGroups) GetMultiReviewerMode() string`

GetMultiReviewerMode returns the MultiReviewerMode field if non-nil, zero value otherwise.

### GetMultiReviewerModeOk

`func (o *TaskGroups) GetMultiReviewerModeOk() (*string, bool)`

GetMultiReviewerModeOk returns a tuple with the MultiReviewerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiReviewerMode

`func (o *TaskGroups) SetMultiReviewerMode(v string)`

SetMultiReviewerMode sets MultiReviewerMode field to given value.


### GetRequiredReviewerCount

`func (o *TaskGroups) GetRequiredReviewerCount() int32`

GetRequiredReviewerCount returns the RequiredReviewerCount field if non-nil, zero value otherwise.

### GetRequiredReviewerCountOk

`func (o *TaskGroups) GetRequiredReviewerCountOk() (*int32, bool)`

GetRequiredReviewerCountOk returns a tuple with the RequiredReviewerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredReviewerCount

`func (o *TaskGroups) SetRequiredReviewerCount(v int32)`

SetRequiredReviewerCount sets RequiredReviewerCount field to given value.

### HasRequiredReviewerCount

`func (o *TaskGroups) HasRequiredReviewerCount() bool`

HasRequiredReviewerCount returns a boolean if a field has been set.

### GetUserTotals

`func (o *TaskGroups) GetUserTotals() interface{}`

GetUserTotals returns the UserTotals field if non-nil, zero value otherwise.

### GetUserTotalsOk

`func (o *TaskGroups) GetUserTotalsOk() (*interface{}, bool)`

GetUserTotalsOk returns a tuple with the UserTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTotals

`func (o *TaskGroups) SetUserTotals(v interface{})`

SetUserTotals sets UserTotals field to given value.

### HasUserTotals

`func (o *TaskGroups) HasUserTotals() bool`

HasUserTotals returns a boolean if a field has been set.

### SetUserTotalsNil

`func (o *TaskGroups) SetUserTotalsNil(b bool)`

 SetUserTotalsNil sets the value for UserTotals to be an explicit nil

### UnsetUserTotals
`func (o *TaskGroups) UnsetUserTotals()`

UnsetUserTotals ensures that no value is present for UserTotals, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


