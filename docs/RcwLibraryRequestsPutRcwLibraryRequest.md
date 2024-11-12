# RcwLibraryRequestsPutRcwLibraryRequest

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

### NewRcwLibraryRequestsPutRcwLibraryRequest

`func NewRcwLibraryRequestsPutRcwLibraryRequest(uid string, referenceMeta interface{}, ) *RcwLibraryRequestsPutRcwLibraryRequest`

NewRcwLibraryRequestsPutRcwLibraryRequest instantiates a new RcwLibraryRequestsPutRcwLibraryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRcwLibraryRequestsPutRcwLibraryRequestWithDefaults

`func NewRcwLibraryRequestsPutRcwLibraryRequestWithDefaults() *RcwLibraryRequestsPutRcwLibraryRequest`

NewRcwLibraryRequestsPutRcwLibraryRequestWithDefaults instantiates a new RcwLibraryRequestsPutRcwLibraryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetUid(v string)`

SetUid sets Uid field to given value.


### GetRcwRequestCategoryId

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetRcwRequestCategoryId() int32`

GetRcwRequestCategoryId returns the RcwRequestCategoryId field if non-nil, zero value otherwise.

### GetRcwRequestCategoryIdOk

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetRcwRequestCategoryIdOk() (*int32, bool)`

GetRcwRequestCategoryIdOk returns a tuple with the RcwRequestCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcwRequestCategoryId

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetRcwRequestCategoryId(v int32)`

SetRcwRequestCategoryId sets RcwRequestCategoryId field to given value.

### HasRcwRequestCategoryId

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) HasRcwRequestCategoryId() bool`

HasRcwRequestCategoryId returns a boolean if a field has been set.

### GetTitle

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAssigneeUserId

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetAssigneeUserId() int32`

GetAssigneeUserId returns the AssigneeUserId field if non-nil, zero value otherwise.

### GetAssigneeUserIdOk

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetAssigneeUserIdOk() (*int32, bool)`

GetAssigneeUserIdOk returns a tuple with the AssigneeUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeUserId

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetAssigneeUserId(v int32)`

SetAssigneeUserId sets AssigneeUserId field to given value.

### HasAssigneeUserId

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) HasAssigneeUserId() bool`

HasAssigneeUserId returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetReferenceMeta

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetReferenceMeta() interface{}`

GetReferenceMeta returns the ReferenceMeta field if non-nil, zero value otherwise.

### GetReferenceMetaOk

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetReferenceMetaOk() (*interface{}, bool)`

GetReferenceMetaOk returns a tuple with the ReferenceMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMeta

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetReferenceMeta(v interface{})`

SetReferenceMeta sets ReferenceMeta field to given value.


### SetReferenceMetaNil

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetReferenceMetaNil(b bool)`

 SetReferenceMetaNil sets the value for ReferenceMeta to be an explicit nil

### UnsetReferenceMeta
`func (o *RcwLibraryRequestsPutRcwLibraryRequest) UnsetReferenceMeta()`

UnsetReferenceMeta ensures that no value is present for ReferenceMeta, not even an explicit nil
### GetFieldData

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *RcwLibraryRequestsPutRcwLibraryRequest) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil
### GetRecurringScheduleId

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetRecurringScheduleId() int32`

GetRecurringScheduleId returns the RecurringScheduleId field if non-nil, zero value otherwise.

### GetRecurringScheduleIdOk

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetRecurringScheduleIdOk() (*int32, bool)`

GetRecurringScheduleIdOk returns a tuple with the RecurringScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringScheduleId

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetRecurringScheduleId(v int32)`

SetRecurringScheduleId sets RecurringScheduleId field to given value.

### HasRecurringScheduleId

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) HasRecurringScheduleId() bool`

HasRecurringScheduleId returns a boolean if a field has been set.

### GetDueDaysAfterLaunch

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetDueDaysAfterLaunch() int32`

GetDueDaysAfterLaunch returns the DueDaysAfterLaunch field if non-nil, zero value otherwise.

### GetDueDaysAfterLaunchOk

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetDueDaysAfterLaunchOk() (*int32, bool)`

GetDueDaysAfterLaunchOk returns a tuple with the DueDaysAfterLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDaysAfterLaunch

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetDueDaysAfterLaunch(v int32)`

SetDueDaysAfterLaunch sets DueDaysAfterLaunch field to given value.

### HasDueDaysAfterLaunch

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) HasDueDaysAfterLaunch() bool`

HasDueDaysAfterLaunch returns a boolean if a field has been set.

### GetPeriodDependent

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetPeriodDependent() bool`

GetPeriodDependent returns the PeriodDependent field if non-nil, zero value otherwise.

### GetPeriodDependentOk

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetPeriodDependentOk() (*bool, bool)`

GetPeriodDependentOk returns a tuple with the PeriodDependent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodDependent

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetPeriodDependent(v bool)`

SetPeriodDependent sets PeriodDependent field to given value.

### HasPeriodDependent

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) HasPeriodDependent() bool`

HasPeriodDependent returns a boolean if a field has been set.

### GetLaunchDate

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetLaunchDate() string`

GetLaunchDate returns the LaunchDate field if non-nil, zero value otherwise.

### GetLaunchDateOk

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetLaunchDateOk() (*string, bool)`

GetLaunchDateOk returns a tuple with the LaunchDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchDate

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetLaunchDate(v string)`

SetLaunchDate sets LaunchDate field to given value.

### HasLaunchDate

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) HasLaunchDate() bool`

HasLaunchDate returns a boolean if a field has been set.

### GetExternalIntegrationUrl

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetExternalIntegrationUrl() string`

GetExternalIntegrationUrl returns the ExternalIntegrationUrl field if non-nil, zero value otherwise.

### GetExternalIntegrationUrlOk

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetExternalIntegrationUrlOk() (*string, bool)`

GetExternalIntegrationUrlOk returns a tuple with the ExternalIntegrationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIntegrationUrl

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetExternalIntegrationUrl(v string)`

SetExternalIntegrationUrl sets ExternalIntegrationUrl field to given value.

### HasExternalIntegrationUrl

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) HasExternalIntegrationUrl() bool`

HasExternalIntegrationUrl returns a boolean if a field has been set.

### GetJiraSelectedProject

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetJiraSelectedProject() string`

GetJiraSelectedProject returns the JiraSelectedProject field if non-nil, zero value otherwise.

### GetJiraSelectedProjectOk

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetJiraSelectedProjectOk() (*string, bool)`

GetJiraSelectedProjectOk returns a tuple with the JiraSelectedProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraSelectedProject

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetJiraSelectedProject(v string)`

SetJiraSelectedProject sets JiraSelectedProject field to given value.

### HasJiraSelectedProject

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) HasJiraSelectedProject() bool`

HasJiraSelectedProject returns a boolean if a field has been set.

### GetJiraSelectedIssueType

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetJiraSelectedIssueType() string`

GetJiraSelectedIssueType returns the JiraSelectedIssueType field if non-nil, zero value otherwise.

### GetJiraSelectedIssueTypeOk

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetJiraSelectedIssueTypeOk() (*string, bool)`

GetJiraSelectedIssueTypeOk returns a tuple with the JiraSelectedIssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraSelectedIssueType

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetJiraSelectedIssueType(v string)`

SetJiraSelectedIssueType sets JiraSelectedIssueType field to given value.

### HasJiraSelectedIssueType

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) HasJiraSelectedIssueType() bool`

HasJiraSelectedIssueType returns a boolean if a field has been set.

### GetAutomaticallyCreateJiraForRecurrences

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetAutomaticallyCreateJiraForRecurrences() bool`

GetAutomaticallyCreateJiraForRecurrences returns the AutomaticallyCreateJiraForRecurrences field if non-nil, zero value otherwise.

### GetAutomaticallyCreateJiraForRecurrencesOk

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetAutomaticallyCreateJiraForRecurrencesOk() (*bool, bool)`

GetAutomaticallyCreateJiraForRecurrencesOk returns a tuple with the AutomaticallyCreateJiraForRecurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyCreateJiraForRecurrences

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetAutomaticallyCreateJiraForRecurrences(v bool)`

SetAutomaticallyCreateJiraForRecurrences sets AutomaticallyCreateJiraForRecurrences field to given value.

### HasAutomaticallyCreateJiraForRecurrences

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) HasAutomaticallyCreateJiraForRecurrences() bool`

HasAutomaticallyCreateJiraForRecurrences returns a boolean if a field has been set.

### GetDescription

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RcwLibraryRequestsPutRcwLibraryRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


