# RcwLibraryRequests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Uid** | **string** |  | 
**RcwRequestCategoryId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;rcw_request_categories.id&#x60;.&lt;fk table&#x3D;&#39;rcw_request_categories&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**AssigneeUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CreatedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**ReferenceMeta** | **interface{}** |  | 
**FieldData** | Pointer to **interface{}** |  | [optional] 
**RecurringScheduleId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;recurring_schedules.id&#x60;.&lt;fk table&#x3D;&#39;recurring_schedules&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**DueDaysAfterLaunch** | Pointer to **int32** |  | [optional] 
**PeriodDependent** | Pointer to **bool** |  | [optional] 
**LaunchDate** | Pointer to **string** |  | [optional] 
**ExternalIntegrationUrl** | Pointer to **string** |  | [optional] 
**JiraSelectedProject** | Pointer to **string** |  | [optional] 
**JiraSelectedIssueType** | Pointer to **string** |  | [optional] 
**AutomaticallyCreateJiraForRecurrences** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewRcwLibraryRequests

`func NewRcwLibraryRequests(uid string, referenceMeta interface{}, ) *RcwLibraryRequests`

NewRcwLibraryRequests instantiates a new RcwLibraryRequests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRcwLibraryRequestsWithDefaults

`func NewRcwLibraryRequestsWithDefaults() *RcwLibraryRequests`

NewRcwLibraryRequestsWithDefaults instantiates a new RcwLibraryRequests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RcwLibraryRequests) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RcwLibraryRequests) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RcwLibraryRequests) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RcwLibraryRequests) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *RcwLibraryRequests) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *RcwLibraryRequests) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *RcwLibraryRequests) SetUid(v string)`

SetUid sets Uid field to given value.


### GetRcwRequestCategoryId

`func (o *RcwLibraryRequests) GetRcwRequestCategoryId() int32`

GetRcwRequestCategoryId returns the RcwRequestCategoryId field if non-nil, zero value otherwise.

### GetRcwRequestCategoryIdOk

`func (o *RcwLibraryRequests) GetRcwRequestCategoryIdOk() (*int32, bool)`

GetRcwRequestCategoryIdOk returns a tuple with the RcwRequestCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcwRequestCategoryId

`func (o *RcwLibraryRequests) SetRcwRequestCategoryId(v int32)`

SetRcwRequestCategoryId sets RcwRequestCategoryId field to given value.

### HasRcwRequestCategoryId

`func (o *RcwLibraryRequests) HasRcwRequestCategoryId() bool`

HasRcwRequestCategoryId returns a boolean if a field has been set.

### GetTitle

`func (o *RcwLibraryRequests) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RcwLibraryRequests) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RcwLibraryRequests) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RcwLibraryRequests) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAssigneeUserId

`func (o *RcwLibraryRequests) GetAssigneeUserId() int32`

GetAssigneeUserId returns the AssigneeUserId field if non-nil, zero value otherwise.

### GetAssigneeUserIdOk

`func (o *RcwLibraryRequests) GetAssigneeUserIdOk() (*int32, bool)`

GetAssigneeUserIdOk returns a tuple with the AssigneeUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeUserId

`func (o *RcwLibraryRequests) SetAssigneeUserId(v int32)`

SetAssigneeUserId sets AssigneeUserId field to given value.

### HasAssigneeUserId

`func (o *RcwLibraryRequests) HasAssigneeUserId() bool`

HasAssigneeUserId returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *RcwLibraryRequests) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *RcwLibraryRequests) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *RcwLibraryRequests) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *RcwLibraryRequests) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RcwLibraryRequests) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RcwLibraryRequests) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RcwLibraryRequests) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RcwLibraryRequests) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RcwLibraryRequests) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RcwLibraryRequests) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RcwLibraryRequests) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RcwLibraryRequests) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *RcwLibraryRequests) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *RcwLibraryRequests) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *RcwLibraryRequests) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *RcwLibraryRequests) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetReferenceMeta

`func (o *RcwLibraryRequests) GetReferenceMeta() interface{}`

GetReferenceMeta returns the ReferenceMeta field if non-nil, zero value otherwise.

### GetReferenceMetaOk

`func (o *RcwLibraryRequests) GetReferenceMetaOk() (*interface{}, bool)`

GetReferenceMetaOk returns a tuple with the ReferenceMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMeta

`func (o *RcwLibraryRequests) SetReferenceMeta(v interface{})`

SetReferenceMeta sets ReferenceMeta field to given value.


### SetReferenceMetaNil

`func (o *RcwLibraryRequests) SetReferenceMetaNil(b bool)`

 SetReferenceMetaNil sets the value for ReferenceMeta to be an explicit nil

### UnsetReferenceMeta
`func (o *RcwLibraryRequests) UnsetReferenceMeta()`

UnsetReferenceMeta ensures that no value is present for ReferenceMeta, not even an explicit nil
### GetFieldData

`func (o *RcwLibraryRequests) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *RcwLibraryRequests) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *RcwLibraryRequests) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *RcwLibraryRequests) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *RcwLibraryRequests) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *RcwLibraryRequests) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil
### GetRecurringScheduleId

`func (o *RcwLibraryRequests) GetRecurringScheduleId() int32`

GetRecurringScheduleId returns the RecurringScheduleId field if non-nil, zero value otherwise.

### GetRecurringScheduleIdOk

`func (o *RcwLibraryRequests) GetRecurringScheduleIdOk() (*int32, bool)`

GetRecurringScheduleIdOk returns a tuple with the RecurringScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringScheduleId

`func (o *RcwLibraryRequests) SetRecurringScheduleId(v int32)`

SetRecurringScheduleId sets RecurringScheduleId field to given value.

### HasRecurringScheduleId

`func (o *RcwLibraryRequests) HasRecurringScheduleId() bool`

HasRecurringScheduleId returns a boolean if a field has been set.

### GetDueDaysAfterLaunch

`func (o *RcwLibraryRequests) GetDueDaysAfterLaunch() int32`

GetDueDaysAfterLaunch returns the DueDaysAfterLaunch field if non-nil, zero value otherwise.

### GetDueDaysAfterLaunchOk

`func (o *RcwLibraryRequests) GetDueDaysAfterLaunchOk() (*int32, bool)`

GetDueDaysAfterLaunchOk returns a tuple with the DueDaysAfterLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDaysAfterLaunch

`func (o *RcwLibraryRequests) SetDueDaysAfterLaunch(v int32)`

SetDueDaysAfterLaunch sets DueDaysAfterLaunch field to given value.

### HasDueDaysAfterLaunch

`func (o *RcwLibraryRequests) HasDueDaysAfterLaunch() bool`

HasDueDaysAfterLaunch returns a boolean if a field has been set.

### GetPeriodDependent

`func (o *RcwLibraryRequests) GetPeriodDependent() bool`

GetPeriodDependent returns the PeriodDependent field if non-nil, zero value otherwise.

### GetPeriodDependentOk

`func (o *RcwLibraryRequests) GetPeriodDependentOk() (*bool, bool)`

GetPeriodDependentOk returns a tuple with the PeriodDependent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodDependent

`func (o *RcwLibraryRequests) SetPeriodDependent(v bool)`

SetPeriodDependent sets PeriodDependent field to given value.

### HasPeriodDependent

`func (o *RcwLibraryRequests) HasPeriodDependent() bool`

HasPeriodDependent returns a boolean if a field has been set.

### GetLaunchDate

`func (o *RcwLibraryRequests) GetLaunchDate() string`

GetLaunchDate returns the LaunchDate field if non-nil, zero value otherwise.

### GetLaunchDateOk

`func (o *RcwLibraryRequests) GetLaunchDateOk() (*string, bool)`

GetLaunchDateOk returns a tuple with the LaunchDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchDate

`func (o *RcwLibraryRequests) SetLaunchDate(v string)`

SetLaunchDate sets LaunchDate field to given value.

### HasLaunchDate

`func (o *RcwLibraryRequests) HasLaunchDate() bool`

HasLaunchDate returns a boolean if a field has been set.

### GetExternalIntegrationUrl

`func (o *RcwLibraryRequests) GetExternalIntegrationUrl() string`

GetExternalIntegrationUrl returns the ExternalIntegrationUrl field if non-nil, zero value otherwise.

### GetExternalIntegrationUrlOk

`func (o *RcwLibraryRequests) GetExternalIntegrationUrlOk() (*string, bool)`

GetExternalIntegrationUrlOk returns a tuple with the ExternalIntegrationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIntegrationUrl

`func (o *RcwLibraryRequests) SetExternalIntegrationUrl(v string)`

SetExternalIntegrationUrl sets ExternalIntegrationUrl field to given value.

### HasExternalIntegrationUrl

`func (o *RcwLibraryRequests) HasExternalIntegrationUrl() bool`

HasExternalIntegrationUrl returns a boolean if a field has been set.

### GetJiraSelectedProject

`func (o *RcwLibraryRequests) GetJiraSelectedProject() string`

GetJiraSelectedProject returns the JiraSelectedProject field if non-nil, zero value otherwise.

### GetJiraSelectedProjectOk

`func (o *RcwLibraryRequests) GetJiraSelectedProjectOk() (*string, bool)`

GetJiraSelectedProjectOk returns a tuple with the JiraSelectedProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraSelectedProject

`func (o *RcwLibraryRequests) SetJiraSelectedProject(v string)`

SetJiraSelectedProject sets JiraSelectedProject field to given value.

### HasJiraSelectedProject

`func (o *RcwLibraryRequests) HasJiraSelectedProject() bool`

HasJiraSelectedProject returns a boolean if a field has been set.

### GetJiraSelectedIssueType

`func (o *RcwLibraryRequests) GetJiraSelectedIssueType() string`

GetJiraSelectedIssueType returns the JiraSelectedIssueType field if non-nil, zero value otherwise.

### GetJiraSelectedIssueTypeOk

`func (o *RcwLibraryRequests) GetJiraSelectedIssueTypeOk() (*string, bool)`

GetJiraSelectedIssueTypeOk returns a tuple with the JiraSelectedIssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraSelectedIssueType

`func (o *RcwLibraryRequests) SetJiraSelectedIssueType(v string)`

SetJiraSelectedIssueType sets JiraSelectedIssueType field to given value.

### HasJiraSelectedIssueType

`func (o *RcwLibraryRequests) HasJiraSelectedIssueType() bool`

HasJiraSelectedIssueType returns a boolean if a field has been set.

### GetAutomaticallyCreateJiraForRecurrences

`func (o *RcwLibraryRequests) GetAutomaticallyCreateJiraForRecurrences() bool`

GetAutomaticallyCreateJiraForRecurrences returns the AutomaticallyCreateJiraForRecurrences field if non-nil, zero value otherwise.

### GetAutomaticallyCreateJiraForRecurrencesOk

`func (o *RcwLibraryRequests) GetAutomaticallyCreateJiraForRecurrencesOk() (*bool, bool)`

GetAutomaticallyCreateJiraForRecurrencesOk returns a tuple with the AutomaticallyCreateJiraForRecurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyCreateJiraForRecurrences

`func (o *RcwLibraryRequests) SetAutomaticallyCreateJiraForRecurrences(v bool)`

SetAutomaticallyCreateJiraForRecurrences sets AutomaticallyCreateJiraForRecurrences field to given value.

### HasAutomaticallyCreateJiraForRecurrences

`func (o *RcwLibraryRequests) HasAutomaticallyCreateJiraForRecurrences() bool`

HasAutomaticallyCreateJiraForRecurrences returns a boolean if a field has been set.

### GetDescription

`func (o *RcwLibraryRequests) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RcwLibraryRequests) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RcwLibraryRequests) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RcwLibraryRequests) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


