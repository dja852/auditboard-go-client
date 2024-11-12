# RcwLibraryRequestsPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] 
**RcwRequestCategoryId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;rcw_request_categories.id&#x60;.&lt;fk table&#x3D;&#39;rcw_request_categories&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**AssigneeUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CreatedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**ReferenceMeta** | Pointer to **interface{}** |  | [optional] 
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

### NewRcwLibraryRequestsPutPreviousValues

`func NewRcwLibraryRequestsPutPreviousValues() *RcwLibraryRequestsPutPreviousValues`

NewRcwLibraryRequestsPutPreviousValues instantiates a new RcwLibraryRequestsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRcwLibraryRequestsPutPreviousValuesWithDefaults

`func NewRcwLibraryRequestsPutPreviousValuesWithDefaults() *RcwLibraryRequestsPutPreviousValues`

NewRcwLibraryRequestsPutPreviousValuesWithDefaults instantiates a new RcwLibraryRequestsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RcwLibraryRequestsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RcwLibraryRequestsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RcwLibraryRequestsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RcwLibraryRequestsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *RcwLibraryRequestsPutPreviousValues) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *RcwLibraryRequestsPutPreviousValues) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *RcwLibraryRequestsPutPreviousValues) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *RcwLibraryRequestsPutPreviousValues) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetRcwRequestCategoryId

`func (o *RcwLibraryRequestsPutPreviousValues) GetRcwRequestCategoryId() int32`

GetRcwRequestCategoryId returns the RcwRequestCategoryId field if non-nil, zero value otherwise.

### GetRcwRequestCategoryIdOk

`func (o *RcwLibraryRequestsPutPreviousValues) GetRcwRequestCategoryIdOk() (*int32, bool)`

GetRcwRequestCategoryIdOk returns a tuple with the RcwRequestCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcwRequestCategoryId

`func (o *RcwLibraryRequestsPutPreviousValues) SetRcwRequestCategoryId(v int32)`

SetRcwRequestCategoryId sets RcwRequestCategoryId field to given value.

### HasRcwRequestCategoryId

`func (o *RcwLibraryRequestsPutPreviousValues) HasRcwRequestCategoryId() bool`

HasRcwRequestCategoryId returns a boolean if a field has been set.

### GetTitle

`func (o *RcwLibraryRequestsPutPreviousValues) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RcwLibraryRequestsPutPreviousValues) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RcwLibraryRequestsPutPreviousValues) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RcwLibraryRequestsPutPreviousValues) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAssigneeUserId

`func (o *RcwLibraryRequestsPutPreviousValues) GetAssigneeUserId() int32`

GetAssigneeUserId returns the AssigneeUserId field if non-nil, zero value otherwise.

### GetAssigneeUserIdOk

`func (o *RcwLibraryRequestsPutPreviousValues) GetAssigneeUserIdOk() (*int32, bool)`

GetAssigneeUserIdOk returns a tuple with the AssigneeUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeUserId

`func (o *RcwLibraryRequestsPutPreviousValues) SetAssigneeUserId(v int32)`

SetAssigneeUserId sets AssigneeUserId field to given value.

### HasAssigneeUserId

`func (o *RcwLibraryRequestsPutPreviousValues) HasAssigneeUserId() bool`

HasAssigneeUserId returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *RcwLibraryRequestsPutPreviousValues) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *RcwLibraryRequestsPutPreviousValues) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *RcwLibraryRequestsPutPreviousValues) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *RcwLibraryRequestsPutPreviousValues) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RcwLibraryRequestsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RcwLibraryRequestsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RcwLibraryRequestsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RcwLibraryRequestsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RcwLibraryRequestsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RcwLibraryRequestsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RcwLibraryRequestsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RcwLibraryRequestsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *RcwLibraryRequestsPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *RcwLibraryRequestsPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *RcwLibraryRequestsPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *RcwLibraryRequestsPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetReferenceMeta

`func (o *RcwLibraryRequestsPutPreviousValues) GetReferenceMeta() interface{}`

GetReferenceMeta returns the ReferenceMeta field if non-nil, zero value otherwise.

### GetReferenceMetaOk

`func (o *RcwLibraryRequestsPutPreviousValues) GetReferenceMetaOk() (*interface{}, bool)`

GetReferenceMetaOk returns a tuple with the ReferenceMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMeta

`func (o *RcwLibraryRequestsPutPreviousValues) SetReferenceMeta(v interface{})`

SetReferenceMeta sets ReferenceMeta field to given value.

### HasReferenceMeta

`func (o *RcwLibraryRequestsPutPreviousValues) HasReferenceMeta() bool`

HasReferenceMeta returns a boolean if a field has been set.

### SetReferenceMetaNil

`func (o *RcwLibraryRequestsPutPreviousValues) SetReferenceMetaNil(b bool)`

 SetReferenceMetaNil sets the value for ReferenceMeta to be an explicit nil

### UnsetReferenceMeta
`func (o *RcwLibraryRequestsPutPreviousValues) UnsetReferenceMeta()`

UnsetReferenceMeta ensures that no value is present for ReferenceMeta, not even an explicit nil
### GetFieldData

`func (o *RcwLibraryRequestsPutPreviousValues) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *RcwLibraryRequestsPutPreviousValues) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *RcwLibraryRequestsPutPreviousValues) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *RcwLibraryRequestsPutPreviousValues) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *RcwLibraryRequestsPutPreviousValues) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *RcwLibraryRequestsPutPreviousValues) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil
### GetRecurringScheduleId

`func (o *RcwLibraryRequestsPutPreviousValues) GetRecurringScheduleId() int32`

GetRecurringScheduleId returns the RecurringScheduleId field if non-nil, zero value otherwise.

### GetRecurringScheduleIdOk

`func (o *RcwLibraryRequestsPutPreviousValues) GetRecurringScheduleIdOk() (*int32, bool)`

GetRecurringScheduleIdOk returns a tuple with the RecurringScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringScheduleId

`func (o *RcwLibraryRequestsPutPreviousValues) SetRecurringScheduleId(v int32)`

SetRecurringScheduleId sets RecurringScheduleId field to given value.

### HasRecurringScheduleId

`func (o *RcwLibraryRequestsPutPreviousValues) HasRecurringScheduleId() bool`

HasRecurringScheduleId returns a boolean if a field has been set.

### GetDueDaysAfterLaunch

`func (o *RcwLibraryRequestsPutPreviousValues) GetDueDaysAfterLaunch() int32`

GetDueDaysAfterLaunch returns the DueDaysAfterLaunch field if non-nil, zero value otherwise.

### GetDueDaysAfterLaunchOk

`func (o *RcwLibraryRequestsPutPreviousValues) GetDueDaysAfterLaunchOk() (*int32, bool)`

GetDueDaysAfterLaunchOk returns a tuple with the DueDaysAfterLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDaysAfterLaunch

`func (o *RcwLibraryRequestsPutPreviousValues) SetDueDaysAfterLaunch(v int32)`

SetDueDaysAfterLaunch sets DueDaysAfterLaunch field to given value.

### HasDueDaysAfterLaunch

`func (o *RcwLibraryRequestsPutPreviousValues) HasDueDaysAfterLaunch() bool`

HasDueDaysAfterLaunch returns a boolean if a field has been set.

### GetPeriodDependent

`func (o *RcwLibraryRequestsPutPreviousValues) GetPeriodDependent() bool`

GetPeriodDependent returns the PeriodDependent field if non-nil, zero value otherwise.

### GetPeriodDependentOk

`func (o *RcwLibraryRequestsPutPreviousValues) GetPeriodDependentOk() (*bool, bool)`

GetPeriodDependentOk returns a tuple with the PeriodDependent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodDependent

`func (o *RcwLibraryRequestsPutPreviousValues) SetPeriodDependent(v bool)`

SetPeriodDependent sets PeriodDependent field to given value.

### HasPeriodDependent

`func (o *RcwLibraryRequestsPutPreviousValues) HasPeriodDependent() bool`

HasPeriodDependent returns a boolean if a field has been set.

### GetLaunchDate

`func (o *RcwLibraryRequestsPutPreviousValues) GetLaunchDate() string`

GetLaunchDate returns the LaunchDate field if non-nil, zero value otherwise.

### GetLaunchDateOk

`func (o *RcwLibraryRequestsPutPreviousValues) GetLaunchDateOk() (*string, bool)`

GetLaunchDateOk returns a tuple with the LaunchDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchDate

`func (o *RcwLibraryRequestsPutPreviousValues) SetLaunchDate(v string)`

SetLaunchDate sets LaunchDate field to given value.

### HasLaunchDate

`func (o *RcwLibraryRequestsPutPreviousValues) HasLaunchDate() bool`

HasLaunchDate returns a boolean if a field has been set.

### GetExternalIntegrationUrl

`func (o *RcwLibraryRequestsPutPreviousValues) GetExternalIntegrationUrl() string`

GetExternalIntegrationUrl returns the ExternalIntegrationUrl field if non-nil, zero value otherwise.

### GetExternalIntegrationUrlOk

`func (o *RcwLibraryRequestsPutPreviousValues) GetExternalIntegrationUrlOk() (*string, bool)`

GetExternalIntegrationUrlOk returns a tuple with the ExternalIntegrationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIntegrationUrl

`func (o *RcwLibraryRequestsPutPreviousValues) SetExternalIntegrationUrl(v string)`

SetExternalIntegrationUrl sets ExternalIntegrationUrl field to given value.

### HasExternalIntegrationUrl

`func (o *RcwLibraryRequestsPutPreviousValues) HasExternalIntegrationUrl() bool`

HasExternalIntegrationUrl returns a boolean if a field has been set.

### GetJiraSelectedProject

`func (o *RcwLibraryRequestsPutPreviousValues) GetJiraSelectedProject() string`

GetJiraSelectedProject returns the JiraSelectedProject field if non-nil, zero value otherwise.

### GetJiraSelectedProjectOk

`func (o *RcwLibraryRequestsPutPreviousValues) GetJiraSelectedProjectOk() (*string, bool)`

GetJiraSelectedProjectOk returns a tuple with the JiraSelectedProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraSelectedProject

`func (o *RcwLibraryRequestsPutPreviousValues) SetJiraSelectedProject(v string)`

SetJiraSelectedProject sets JiraSelectedProject field to given value.

### HasJiraSelectedProject

`func (o *RcwLibraryRequestsPutPreviousValues) HasJiraSelectedProject() bool`

HasJiraSelectedProject returns a boolean if a field has been set.

### GetJiraSelectedIssueType

`func (o *RcwLibraryRequestsPutPreviousValues) GetJiraSelectedIssueType() string`

GetJiraSelectedIssueType returns the JiraSelectedIssueType field if non-nil, zero value otherwise.

### GetJiraSelectedIssueTypeOk

`func (o *RcwLibraryRequestsPutPreviousValues) GetJiraSelectedIssueTypeOk() (*string, bool)`

GetJiraSelectedIssueTypeOk returns a tuple with the JiraSelectedIssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraSelectedIssueType

`func (o *RcwLibraryRequestsPutPreviousValues) SetJiraSelectedIssueType(v string)`

SetJiraSelectedIssueType sets JiraSelectedIssueType field to given value.

### HasJiraSelectedIssueType

`func (o *RcwLibraryRequestsPutPreviousValues) HasJiraSelectedIssueType() bool`

HasJiraSelectedIssueType returns a boolean if a field has been set.

### GetAutomaticallyCreateJiraForRecurrences

`func (o *RcwLibraryRequestsPutPreviousValues) GetAutomaticallyCreateJiraForRecurrences() bool`

GetAutomaticallyCreateJiraForRecurrences returns the AutomaticallyCreateJiraForRecurrences field if non-nil, zero value otherwise.

### GetAutomaticallyCreateJiraForRecurrencesOk

`func (o *RcwLibraryRequestsPutPreviousValues) GetAutomaticallyCreateJiraForRecurrencesOk() (*bool, bool)`

GetAutomaticallyCreateJiraForRecurrencesOk returns a tuple with the AutomaticallyCreateJiraForRecurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyCreateJiraForRecurrences

`func (o *RcwLibraryRequestsPutPreviousValues) SetAutomaticallyCreateJiraForRecurrences(v bool)`

SetAutomaticallyCreateJiraForRecurrences sets AutomaticallyCreateJiraForRecurrences field to given value.

### HasAutomaticallyCreateJiraForRecurrences

`func (o *RcwLibraryRequestsPutPreviousValues) HasAutomaticallyCreateJiraForRecurrences() bool`

HasAutomaticallyCreateJiraForRecurrences returns a boolean if a field has been set.

### GetDescription

`func (o *RcwLibraryRequestsPutPreviousValues) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RcwLibraryRequestsPutPreviousValues) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RcwLibraryRequestsPutPreviousValues) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RcwLibraryRequestsPutPreviousValues) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


