# TaskItemsPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**IsAcknowledgePrepared** | Pointer to **bool** |  | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**SubmittedDate** | Pointer to **string** |  | [optional] 
**CertifiedDate** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**TaskGroupId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;task_groups.id&#x60;.&lt;fk table&#x3D;&#39;task_groups&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AssigneeUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**SubmittedByUserId** | Pointer to **int32** |  | [optional] 
**CertifiedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AssigneeResponse** | Pointer to **string** |  | [optional] 
**BudgetHours** | Pointer to **float32** |  | [optional] [default to 0]
**ActualHours** | Pointer to **float32** |  | [optional] [default to 0]
**TaskPriorityId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;task_priorities.id&#x60;.&lt;fk table&#x3D;&#39;task_priorities&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TaskPeriodId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;task_periods.id&#x60;.&lt;fk table&#x3D;&#39;task_periods&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ReviewerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**LocationId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;locations.id&#x60;.&lt;fk table&#x3D;&#39;locations&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**DepartmentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;departments.id&#x60;.&lt;fk table&#x3D;&#39;departments&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ReferenceNotes** | Pointer to **string** |  | [optional] 
**ControlsDatumId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;controls_data.id&#x60;.&lt;fk table&#x3D;&#39;controls_data&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TaskCategoryId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;task_categories.id&#x60;.&lt;fk table&#x3D;&#39;task_categories&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CreatorUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**MetaConfig** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [default to ""]
**TestSectionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_sections.id&#x60;.&lt;fk table&#x3D;&#39;test_sections&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issues.id&#x60;.&lt;fk table&#x3D;&#39;issues&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpenedDate** | Pointer to **string** |  | [optional] 
**OpenedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ClosedDate** | Pointer to **string** |  | [optional] 
**ClosedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audits.id&#x60;.&lt;fk table&#x3D;&#39;ops_audits&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditItemId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_items.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_items&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ActionPlanId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;action_plans.id&#x60;.&lt;fk table&#x3D;&#39;action_plans&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risks.id&#x60;.&lt;fk table&#x3D;&#39;risks&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ExternalIntegrationUrl** | Pointer to **string** |  | [optional] 
**FrameworkItemId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;framework_items.id&#x60;.&lt;fk table&#x3D;&#39;framework_items&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**HistoricalTaskIds** | Pointer to **interface{}** |  | [optional] 
**NarrativeId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;narratives.id&#x60;.&lt;fk table&#x3D;&#39;narratives&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AuditableEntityId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;auditable_entities.id&#x60;.&lt;fk table&#x3D;&#39;auditable_entities&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**PolicyId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;policies.id&#x60;.&lt;fk table&#x3D;&#39;policies&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FirstSubmittedDate** | Pointer to **string** |  | [optional] 
**FirstCertifiedDate** | Pointer to **string** |  | [optional] 
**FirstClosedDate** | Pointer to **string** |  | [optional] 
**AuditableEntityRiskId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;auditable_entity_risks.id&#x60;.&lt;fk table&#x3D;&#39;auditable_entity_risks&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "workstream"]
**EsgTopicId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_topics.id&#x60;.&lt;fk table&#x3D;&#39;esg_topics&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CreateReason** | Pointer to **string** |  | [optional] 

## Methods

### NewTaskItemsPutPreviousValues

`func NewTaskItemsPutPreviousValues() *TaskItemsPutPreviousValues`

NewTaskItemsPutPreviousValues instantiates a new TaskItemsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskItemsPutPreviousValuesWithDefaults

`func NewTaskItemsPutPreviousValuesWithDefaults() *TaskItemsPutPreviousValues`

NewTaskItemsPutPreviousValuesWithDefaults instantiates a new TaskItemsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskItemsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskItemsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskItemsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TaskItemsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskItemsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskItemsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskItemsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaskItemsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TaskItemsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskItemsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskItemsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TaskItemsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *TaskItemsPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *TaskItemsPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *TaskItemsPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *TaskItemsPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetIsAcknowledgePrepared

`func (o *TaskItemsPutPreviousValues) GetIsAcknowledgePrepared() bool`

GetIsAcknowledgePrepared returns the IsAcknowledgePrepared field if non-nil, zero value otherwise.

### GetIsAcknowledgePreparedOk

`func (o *TaskItemsPutPreviousValues) GetIsAcknowledgePreparedOk() (*bool, bool)`

GetIsAcknowledgePreparedOk returns a tuple with the IsAcknowledgePrepared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAcknowledgePrepared

`func (o *TaskItemsPutPreviousValues) SetIsAcknowledgePrepared(v bool)`

SetIsAcknowledgePrepared sets IsAcknowledgePrepared field to given value.

### HasIsAcknowledgePrepared

`func (o *TaskItemsPutPreviousValues) HasIsAcknowledgePrepared() bool`

HasIsAcknowledgePrepared returns a boolean if a field has been set.

### GetDueDate

`func (o *TaskItemsPutPreviousValues) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *TaskItemsPutPreviousValues) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *TaskItemsPutPreviousValues) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *TaskItemsPutPreviousValues) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetSubmittedDate

`func (o *TaskItemsPutPreviousValues) GetSubmittedDate() string`

GetSubmittedDate returns the SubmittedDate field if non-nil, zero value otherwise.

### GetSubmittedDateOk

`func (o *TaskItemsPutPreviousValues) GetSubmittedDateOk() (*string, bool)`

GetSubmittedDateOk returns a tuple with the SubmittedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedDate

`func (o *TaskItemsPutPreviousValues) SetSubmittedDate(v string)`

SetSubmittedDate sets SubmittedDate field to given value.

### HasSubmittedDate

`func (o *TaskItemsPutPreviousValues) HasSubmittedDate() bool`

HasSubmittedDate returns a boolean if a field has been set.

### GetCertifiedDate

`func (o *TaskItemsPutPreviousValues) GetCertifiedDate() string`

GetCertifiedDate returns the CertifiedDate field if non-nil, zero value otherwise.

### GetCertifiedDateOk

`func (o *TaskItemsPutPreviousValues) GetCertifiedDateOk() (*string, bool)`

GetCertifiedDateOk returns a tuple with the CertifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifiedDate

`func (o *TaskItemsPutPreviousValues) SetCertifiedDate(v string)`

SetCertifiedDate sets CertifiedDate field to given value.

### HasCertifiedDate

`func (o *TaskItemsPutPreviousValues) HasCertifiedDate() bool`

HasCertifiedDate returns a boolean if a field has been set.

### GetOrder

`func (o *TaskItemsPutPreviousValues) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TaskItemsPutPreviousValues) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TaskItemsPutPreviousValues) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TaskItemsPutPreviousValues) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetTaskGroupId

`func (o *TaskItemsPutPreviousValues) GetTaskGroupId() int32`

GetTaskGroupId returns the TaskGroupId field if non-nil, zero value otherwise.

### GetTaskGroupIdOk

`func (o *TaskItemsPutPreviousValues) GetTaskGroupIdOk() (*int32, bool)`

GetTaskGroupIdOk returns a tuple with the TaskGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskGroupId

`func (o *TaskItemsPutPreviousValues) SetTaskGroupId(v int32)`

SetTaskGroupId sets TaskGroupId field to given value.

### HasTaskGroupId

`func (o *TaskItemsPutPreviousValues) HasTaskGroupId() bool`

HasTaskGroupId returns a boolean if a field has been set.

### GetAssigneeUserId

`func (o *TaskItemsPutPreviousValues) GetAssigneeUserId() int32`

GetAssigneeUserId returns the AssigneeUserId field if non-nil, zero value otherwise.

### GetAssigneeUserIdOk

`func (o *TaskItemsPutPreviousValues) GetAssigneeUserIdOk() (*int32, bool)`

GetAssigneeUserIdOk returns a tuple with the AssigneeUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeUserId

`func (o *TaskItemsPutPreviousValues) SetAssigneeUserId(v int32)`

SetAssigneeUserId sets AssigneeUserId field to given value.

### HasAssigneeUserId

`func (o *TaskItemsPutPreviousValues) HasAssigneeUserId() bool`

HasAssigneeUserId returns a boolean if a field has been set.

### GetSubmittedByUserId

`func (o *TaskItemsPutPreviousValues) GetSubmittedByUserId() int32`

GetSubmittedByUserId returns the SubmittedByUserId field if non-nil, zero value otherwise.

### GetSubmittedByUserIdOk

`func (o *TaskItemsPutPreviousValues) GetSubmittedByUserIdOk() (*int32, bool)`

GetSubmittedByUserIdOk returns a tuple with the SubmittedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedByUserId

`func (o *TaskItemsPutPreviousValues) SetSubmittedByUserId(v int32)`

SetSubmittedByUserId sets SubmittedByUserId field to given value.

### HasSubmittedByUserId

`func (o *TaskItemsPutPreviousValues) HasSubmittedByUserId() bool`

HasSubmittedByUserId returns a boolean if a field has been set.

### GetCertifiedByUserId

`func (o *TaskItemsPutPreviousValues) GetCertifiedByUserId() int32`

GetCertifiedByUserId returns the CertifiedByUserId field if non-nil, zero value otherwise.

### GetCertifiedByUserIdOk

`func (o *TaskItemsPutPreviousValues) GetCertifiedByUserIdOk() (*int32, bool)`

GetCertifiedByUserIdOk returns a tuple with the CertifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifiedByUserId

`func (o *TaskItemsPutPreviousValues) SetCertifiedByUserId(v int32)`

SetCertifiedByUserId sets CertifiedByUserId field to given value.

### HasCertifiedByUserId

`func (o *TaskItemsPutPreviousValues) HasCertifiedByUserId() bool`

HasCertifiedByUserId returns a boolean if a field has been set.

### GetTitle

`func (o *TaskItemsPutPreviousValues) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TaskItemsPutPreviousValues) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TaskItemsPutPreviousValues) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TaskItemsPutPreviousValues) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *TaskItemsPutPreviousValues) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskItemsPutPreviousValues) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskItemsPutPreviousValues) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskItemsPutPreviousValues) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAssigneeResponse

`func (o *TaskItemsPutPreviousValues) GetAssigneeResponse() string`

GetAssigneeResponse returns the AssigneeResponse field if non-nil, zero value otherwise.

### GetAssigneeResponseOk

`func (o *TaskItemsPutPreviousValues) GetAssigneeResponseOk() (*string, bool)`

GetAssigneeResponseOk returns a tuple with the AssigneeResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeResponse

`func (o *TaskItemsPutPreviousValues) SetAssigneeResponse(v string)`

SetAssigneeResponse sets AssigneeResponse field to given value.

### HasAssigneeResponse

`func (o *TaskItemsPutPreviousValues) HasAssigneeResponse() bool`

HasAssigneeResponse returns a boolean if a field has been set.

### GetBudgetHours

`func (o *TaskItemsPutPreviousValues) GetBudgetHours() float32`

GetBudgetHours returns the BudgetHours field if non-nil, zero value otherwise.

### GetBudgetHoursOk

`func (o *TaskItemsPutPreviousValues) GetBudgetHoursOk() (*float32, bool)`

GetBudgetHoursOk returns a tuple with the BudgetHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetHours

`func (o *TaskItemsPutPreviousValues) SetBudgetHours(v float32)`

SetBudgetHours sets BudgetHours field to given value.

### HasBudgetHours

`func (o *TaskItemsPutPreviousValues) HasBudgetHours() bool`

HasBudgetHours returns a boolean if a field has been set.

### GetActualHours

`func (o *TaskItemsPutPreviousValues) GetActualHours() float32`

GetActualHours returns the ActualHours field if non-nil, zero value otherwise.

### GetActualHoursOk

`func (o *TaskItemsPutPreviousValues) GetActualHoursOk() (*float32, bool)`

GetActualHoursOk returns a tuple with the ActualHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualHours

`func (o *TaskItemsPutPreviousValues) SetActualHours(v float32)`

SetActualHours sets ActualHours field to given value.

### HasActualHours

`func (o *TaskItemsPutPreviousValues) HasActualHours() bool`

HasActualHours returns a boolean if a field has been set.

### GetTaskPriorityId

`func (o *TaskItemsPutPreviousValues) GetTaskPriorityId() int32`

GetTaskPriorityId returns the TaskPriorityId field if non-nil, zero value otherwise.

### GetTaskPriorityIdOk

`func (o *TaskItemsPutPreviousValues) GetTaskPriorityIdOk() (*int32, bool)`

GetTaskPriorityIdOk returns a tuple with the TaskPriorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPriorityId

`func (o *TaskItemsPutPreviousValues) SetTaskPriorityId(v int32)`

SetTaskPriorityId sets TaskPriorityId field to given value.

### HasTaskPriorityId

`func (o *TaskItemsPutPreviousValues) HasTaskPriorityId() bool`

HasTaskPriorityId returns a boolean if a field has been set.

### GetTaskPeriodId

`func (o *TaskItemsPutPreviousValues) GetTaskPeriodId() int32`

GetTaskPeriodId returns the TaskPeriodId field if non-nil, zero value otherwise.

### GetTaskPeriodIdOk

`func (o *TaskItemsPutPreviousValues) GetTaskPeriodIdOk() (*int32, bool)`

GetTaskPeriodIdOk returns a tuple with the TaskPeriodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPeriodId

`func (o *TaskItemsPutPreviousValues) SetTaskPeriodId(v int32)`

SetTaskPeriodId sets TaskPeriodId field to given value.

### HasTaskPeriodId

`func (o *TaskItemsPutPreviousValues) HasTaskPeriodId() bool`

HasTaskPeriodId returns a boolean if a field has been set.

### GetReviewerUserId

`func (o *TaskItemsPutPreviousValues) GetReviewerUserId() int32`

GetReviewerUserId returns the ReviewerUserId field if non-nil, zero value otherwise.

### GetReviewerUserIdOk

`func (o *TaskItemsPutPreviousValues) GetReviewerUserIdOk() (*int32, bool)`

GetReviewerUserIdOk returns a tuple with the ReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerUserId

`func (o *TaskItemsPutPreviousValues) SetReviewerUserId(v int32)`

SetReviewerUserId sets ReviewerUserId field to given value.

### HasReviewerUserId

`func (o *TaskItemsPutPreviousValues) HasReviewerUserId() bool`

HasReviewerUserId returns a boolean if a field has been set.

### GetLocationId

`func (o *TaskItemsPutPreviousValues) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *TaskItemsPutPreviousValues) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *TaskItemsPutPreviousValues) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *TaskItemsPutPreviousValues) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetDepartmentId

`func (o *TaskItemsPutPreviousValues) GetDepartmentId() int32`

GetDepartmentId returns the DepartmentId field if non-nil, zero value otherwise.

### GetDepartmentIdOk

`func (o *TaskItemsPutPreviousValues) GetDepartmentIdOk() (*int32, bool)`

GetDepartmentIdOk returns a tuple with the DepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentId

`func (o *TaskItemsPutPreviousValues) SetDepartmentId(v int32)`

SetDepartmentId sets DepartmentId field to given value.

### HasDepartmentId

`func (o *TaskItemsPutPreviousValues) HasDepartmentId() bool`

HasDepartmentId returns a boolean if a field has been set.

### GetReferenceNotes

`func (o *TaskItemsPutPreviousValues) GetReferenceNotes() string`

GetReferenceNotes returns the ReferenceNotes field if non-nil, zero value otherwise.

### GetReferenceNotesOk

`func (o *TaskItemsPutPreviousValues) GetReferenceNotesOk() (*string, bool)`

GetReferenceNotesOk returns a tuple with the ReferenceNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNotes

`func (o *TaskItemsPutPreviousValues) SetReferenceNotes(v string)`

SetReferenceNotes sets ReferenceNotes field to given value.

### HasReferenceNotes

`func (o *TaskItemsPutPreviousValues) HasReferenceNotes() bool`

HasReferenceNotes returns a boolean if a field has been set.

### GetControlsDatumId

`func (o *TaskItemsPutPreviousValues) GetControlsDatumId() int32`

GetControlsDatumId returns the ControlsDatumId field if non-nil, zero value otherwise.

### GetControlsDatumIdOk

`func (o *TaskItemsPutPreviousValues) GetControlsDatumIdOk() (*int32, bool)`

GetControlsDatumIdOk returns a tuple with the ControlsDatumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlsDatumId

`func (o *TaskItemsPutPreviousValues) SetControlsDatumId(v int32)`

SetControlsDatumId sets ControlsDatumId field to given value.

### HasControlsDatumId

`func (o *TaskItemsPutPreviousValues) HasControlsDatumId() bool`

HasControlsDatumId returns a boolean if a field has been set.

### GetTaskCategoryId

`func (o *TaskItemsPutPreviousValues) GetTaskCategoryId() int32`

GetTaskCategoryId returns the TaskCategoryId field if non-nil, zero value otherwise.

### GetTaskCategoryIdOk

`func (o *TaskItemsPutPreviousValues) GetTaskCategoryIdOk() (*int32, bool)`

GetTaskCategoryIdOk returns a tuple with the TaskCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCategoryId

`func (o *TaskItemsPutPreviousValues) SetTaskCategoryId(v int32)`

SetTaskCategoryId sets TaskCategoryId field to given value.

### HasTaskCategoryId

`func (o *TaskItemsPutPreviousValues) HasTaskCategoryId() bool`

HasTaskCategoryId returns a boolean if a field has been set.

### GetCreatorUserId

`func (o *TaskItemsPutPreviousValues) GetCreatorUserId() int32`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *TaskItemsPutPreviousValues) GetCreatorUserIdOk() (*int32, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *TaskItemsPutPreviousValues) SetCreatorUserId(v int32)`

SetCreatorUserId sets CreatorUserId field to given value.

### HasCreatorUserId

`func (o *TaskItemsPutPreviousValues) HasCreatorUserId() bool`

HasCreatorUserId returns a boolean if a field has been set.

### GetMetaConfig

`func (o *TaskItemsPutPreviousValues) GetMetaConfig() string`

GetMetaConfig returns the MetaConfig field if non-nil, zero value otherwise.

### GetMetaConfigOk

`func (o *TaskItemsPutPreviousValues) GetMetaConfigOk() (*string, bool)`

GetMetaConfigOk returns a tuple with the MetaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaConfig

`func (o *TaskItemsPutPreviousValues) SetMetaConfig(v string)`

SetMetaConfig sets MetaConfig field to given value.

### HasMetaConfig

`func (o *TaskItemsPutPreviousValues) HasMetaConfig() bool`

HasMetaConfig returns a boolean if a field has been set.

### GetStatus

`func (o *TaskItemsPutPreviousValues) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskItemsPutPreviousValues) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskItemsPutPreviousValues) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskItemsPutPreviousValues) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTestSectionId

`func (o *TaskItemsPutPreviousValues) GetTestSectionId() int32`

GetTestSectionId returns the TestSectionId field if non-nil, zero value otherwise.

### GetTestSectionIdOk

`func (o *TaskItemsPutPreviousValues) GetTestSectionIdOk() (*int32, bool)`

GetTestSectionIdOk returns a tuple with the TestSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSectionId

`func (o *TaskItemsPutPreviousValues) SetTestSectionId(v int32)`

SetTestSectionId sets TestSectionId field to given value.

### HasTestSectionId

`func (o *TaskItemsPutPreviousValues) HasTestSectionId() bool`

HasTestSectionId returns a boolean if a field has been set.

### GetIssueId

`func (o *TaskItemsPutPreviousValues) GetIssueId() int32`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *TaskItemsPutPreviousValues) GetIssueIdOk() (*int32, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *TaskItemsPutPreviousValues) SetIssueId(v int32)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *TaskItemsPutPreviousValues) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.

### GetOpenedDate

`func (o *TaskItemsPutPreviousValues) GetOpenedDate() string`

GetOpenedDate returns the OpenedDate field if non-nil, zero value otherwise.

### GetOpenedDateOk

`func (o *TaskItemsPutPreviousValues) GetOpenedDateOk() (*string, bool)`

GetOpenedDateOk returns a tuple with the OpenedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedDate

`func (o *TaskItemsPutPreviousValues) SetOpenedDate(v string)`

SetOpenedDate sets OpenedDate field to given value.

### HasOpenedDate

`func (o *TaskItemsPutPreviousValues) HasOpenedDate() bool`

HasOpenedDate returns a boolean if a field has been set.

### GetOpenedByUserId

`func (o *TaskItemsPutPreviousValues) GetOpenedByUserId() int32`

GetOpenedByUserId returns the OpenedByUserId field if non-nil, zero value otherwise.

### GetOpenedByUserIdOk

`func (o *TaskItemsPutPreviousValues) GetOpenedByUserIdOk() (*int32, bool)`

GetOpenedByUserIdOk returns a tuple with the OpenedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedByUserId

`func (o *TaskItemsPutPreviousValues) SetOpenedByUserId(v int32)`

SetOpenedByUserId sets OpenedByUserId field to given value.

### HasOpenedByUserId

`func (o *TaskItemsPutPreviousValues) HasOpenedByUserId() bool`

HasOpenedByUserId returns a boolean if a field has been set.

### GetClosedDate

`func (o *TaskItemsPutPreviousValues) GetClosedDate() string`

GetClosedDate returns the ClosedDate field if non-nil, zero value otherwise.

### GetClosedDateOk

`func (o *TaskItemsPutPreviousValues) GetClosedDateOk() (*string, bool)`

GetClosedDateOk returns a tuple with the ClosedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedDate

`func (o *TaskItemsPutPreviousValues) SetClosedDate(v string)`

SetClosedDate sets ClosedDate field to given value.

### HasClosedDate

`func (o *TaskItemsPutPreviousValues) HasClosedDate() bool`

HasClosedDate returns a boolean if a field has been set.

### GetClosedByUserId

`func (o *TaskItemsPutPreviousValues) GetClosedByUserId() int32`

GetClosedByUserId returns the ClosedByUserId field if non-nil, zero value otherwise.

### GetClosedByUserIdOk

`func (o *TaskItemsPutPreviousValues) GetClosedByUserIdOk() (*int32, bool)`

GetClosedByUserIdOk returns a tuple with the ClosedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedByUserId

`func (o *TaskItemsPutPreviousValues) SetClosedByUserId(v int32)`

SetClosedByUserId sets ClosedByUserId field to given value.

### HasClosedByUserId

`func (o *TaskItemsPutPreviousValues) HasClosedByUserId() bool`

HasClosedByUserId returns a boolean if a field has been set.

### GetOpsAuditId

`func (o *TaskItemsPutPreviousValues) GetOpsAuditId() int32`

GetOpsAuditId returns the OpsAuditId field if non-nil, zero value otherwise.

### GetOpsAuditIdOk

`func (o *TaskItemsPutPreviousValues) GetOpsAuditIdOk() (*int32, bool)`

GetOpsAuditIdOk returns a tuple with the OpsAuditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditId

`func (o *TaskItemsPutPreviousValues) SetOpsAuditId(v int32)`

SetOpsAuditId sets OpsAuditId field to given value.

### HasOpsAuditId

`func (o *TaskItemsPutPreviousValues) HasOpsAuditId() bool`

HasOpsAuditId returns a boolean if a field has been set.

### GetOpsAuditItemId

`func (o *TaskItemsPutPreviousValues) GetOpsAuditItemId() int32`

GetOpsAuditItemId returns the OpsAuditItemId field if non-nil, zero value otherwise.

### GetOpsAuditItemIdOk

`func (o *TaskItemsPutPreviousValues) GetOpsAuditItemIdOk() (*int32, bool)`

GetOpsAuditItemIdOk returns a tuple with the OpsAuditItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditItemId

`func (o *TaskItemsPutPreviousValues) SetOpsAuditItemId(v int32)`

SetOpsAuditItemId sets OpsAuditItemId field to given value.

### HasOpsAuditItemId

`func (o *TaskItemsPutPreviousValues) HasOpsAuditItemId() bool`

HasOpsAuditItemId returns a boolean if a field has been set.

### GetActionPlanId

`func (o *TaskItemsPutPreviousValues) GetActionPlanId() int32`

GetActionPlanId returns the ActionPlanId field if non-nil, zero value otherwise.

### GetActionPlanIdOk

`func (o *TaskItemsPutPreviousValues) GetActionPlanIdOk() (*int32, bool)`

GetActionPlanIdOk returns a tuple with the ActionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPlanId

`func (o *TaskItemsPutPreviousValues) SetActionPlanId(v int32)`

SetActionPlanId sets ActionPlanId field to given value.

### HasActionPlanId

`func (o *TaskItemsPutPreviousValues) HasActionPlanId() bool`

HasActionPlanId returns a boolean if a field has been set.

### GetRiskId

`func (o *TaskItemsPutPreviousValues) GetRiskId() int32`

GetRiskId returns the RiskId field if non-nil, zero value otherwise.

### GetRiskIdOk

`func (o *TaskItemsPutPreviousValues) GetRiskIdOk() (*int32, bool)`

GetRiskIdOk returns a tuple with the RiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskId

`func (o *TaskItemsPutPreviousValues) SetRiskId(v int32)`

SetRiskId sets RiskId field to given value.

### HasRiskId

`func (o *TaskItemsPutPreviousValues) HasRiskId() bool`

HasRiskId returns a boolean if a field has been set.

### GetExternalIntegrationUrl

`func (o *TaskItemsPutPreviousValues) GetExternalIntegrationUrl() string`

GetExternalIntegrationUrl returns the ExternalIntegrationUrl field if non-nil, zero value otherwise.

### GetExternalIntegrationUrlOk

`func (o *TaskItemsPutPreviousValues) GetExternalIntegrationUrlOk() (*string, bool)`

GetExternalIntegrationUrlOk returns a tuple with the ExternalIntegrationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIntegrationUrl

`func (o *TaskItemsPutPreviousValues) SetExternalIntegrationUrl(v string)`

SetExternalIntegrationUrl sets ExternalIntegrationUrl field to given value.

### HasExternalIntegrationUrl

`func (o *TaskItemsPutPreviousValues) HasExternalIntegrationUrl() bool`

HasExternalIntegrationUrl returns a boolean if a field has been set.

### GetFrameworkItemId

`func (o *TaskItemsPutPreviousValues) GetFrameworkItemId() int32`

GetFrameworkItemId returns the FrameworkItemId field if non-nil, zero value otherwise.

### GetFrameworkItemIdOk

`func (o *TaskItemsPutPreviousValues) GetFrameworkItemIdOk() (*int32, bool)`

GetFrameworkItemIdOk returns a tuple with the FrameworkItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemId

`func (o *TaskItemsPutPreviousValues) SetFrameworkItemId(v int32)`

SetFrameworkItemId sets FrameworkItemId field to given value.

### HasFrameworkItemId

`func (o *TaskItemsPutPreviousValues) HasFrameworkItemId() bool`

HasFrameworkItemId returns a boolean if a field has been set.

### GetHistoricalTaskIds

`func (o *TaskItemsPutPreviousValues) GetHistoricalTaskIds() interface{}`

GetHistoricalTaskIds returns the HistoricalTaskIds field if non-nil, zero value otherwise.

### GetHistoricalTaskIdsOk

`func (o *TaskItemsPutPreviousValues) GetHistoricalTaskIdsOk() (*interface{}, bool)`

GetHistoricalTaskIdsOk returns a tuple with the HistoricalTaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricalTaskIds

`func (o *TaskItemsPutPreviousValues) SetHistoricalTaskIds(v interface{})`

SetHistoricalTaskIds sets HistoricalTaskIds field to given value.

### HasHistoricalTaskIds

`func (o *TaskItemsPutPreviousValues) HasHistoricalTaskIds() bool`

HasHistoricalTaskIds returns a boolean if a field has been set.

### SetHistoricalTaskIdsNil

`func (o *TaskItemsPutPreviousValues) SetHistoricalTaskIdsNil(b bool)`

 SetHistoricalTaskIdsNil sets the value for HistoricalTaskIds to be an explicit nil

### UnsetHistoricalTaskIds
`func (o *TaskItemsPutPreviousValues) UnsetHistoricalTaskIds()`

UnsetHistoricalTaskIds ensures that no value is present for HistoricalTaskIds, not even an explicit nil
### GetNarrativeId

`func (o *TaskItemsPutPreviousValues) GetNarrativeId() int32`

GetNarrativeId returns the NarrativeId field if non-nil, zero value otherwise.

### GetNarrativeIdOk

`func (o *TaskItemsPutPreviousValues) GetNarrativeIdOk() (*int32, bool)`

GetNarrativeIdOk returns a tuple with the NarrativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarrativeId

`func (o *TaskItemsPutPreviousValues) SetNarrativeId(v int32)`

SetNarrativeId sets NarrativeId field to given value.

### HasNarrativeId

`func (o *TaskItemsPutPreviousValues) HasNarrativeId() bool`

HasNarrativeId returns a boolean if a field has been set.

### GetAuditableEntityId

`func (o *TaskItemsPutPreviousValues) GetAuditableEntityId() int32`

GetAuditableEntityId returns the AuditableEntityId field if non-nil, zero value otherwise.

### GetAuditableEntityIdOk

`func (o *TaskItemsPutPreviousValues) GetAuditableEntityIdOk() (*int32, bool)`

GetAuditableEntityIdOk returns a tuple with the AuditableEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditableEntityId

`func (o *TaskItemsPutPreviousValues) SetAuditableEntityId(v int32)`

SetAuditableEntityId sets AuditableEntityId field to given value.

### HasAuditableEntityId

`func (o *TaskItemsPutPreviousValues) HasAuditableEntityId() bool`

HasAuditableEntityId returns a boolean if a field has been set.

### GetPolicyId

`func (o *TaskItemsPutPreviousValues) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *TaskItemsPutPreviousValues) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *TaskItemsPutPreviousValues) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *TaskItemsPutPreviousValues) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetFirstSubmittedDate

`func (o *TaskItemsPutPreviousValues) GetFirstSubmittedDate() string`

GetFirstSubmittedDate returns the FirstSubmittedDate field if non-nil, zero value otherwise.

### GetFirstSubmittedDateOk

`func (o *TaskItemsPutPreviousValues) GetFirstSubmittedDateOk() (*string, bool)`

GetFirstSubmittedDateOk returns a tuple with the FirstSubmittedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSubmittedDate

`func (o *TaskItemsPutPreviousValues) SetFirstSubmittedDate(v string)`

SetFirstSubmittedDate sets FirstSubmittedDate field to given value.

### HasFirstSubmittedDate

`func (o *TaskItemsPutPreviousValues) HasFirstSubmittedDate() bool`

HasFirstSubmittedDate returns a boolean if a field has been set.

### GetFirstCertifiedDate

`func (o *TaskItemsPutPreviousValues) GetFirstCertifiedDate() string`

GetFirstCertifiedDate returns the FirstCertifiedDate field if non-nil, zero value otherwise.

### GetFirstCertifiedDateOk

`func (o *TaskItemsPutPreviousValues) GetFirstCertifiedDateOk() (*string, bool)`

GetFirstCertifiedDateOk returns a tuple with the FirstCertifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCertifiedDate

`func (o *TaskItemsPutPreviousValues) SetFirstCertifiedDate(v string)`

SetFirstCertifiedDate sets FirstCertifiedDate field to given value.

### HasFirstCertifiedDate

`func (o *TaskItemsPutPreviousValues) HasFirstCertifiedDate() bool`

HasFirstCertifiedDate returns a boolean if a field has been set.

### GetFirstClosedDate

`func (o *TaskItemsPutPreviousValues) GetFirstClosedDate() string`

GetFirstClosedDate returns the FirstClosedDate field if non-nil, zero value otherwise.

### GetFirstClosedDateOk

`func (o *TaskItemsPutPreviousValues) GetFirstClosedDateOk() (*string, bool)`

GetFirstClosedDateOk returns a tuple with the FirstClosedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstClosedDate

`func (o *TaskItemsPutPreviousValues) SetFirstClosedDate(v string)`

SetFirstClosedDate sets FirstClosedDate field to given value.

### HasFirstClosedDate

`func (o *TaskItemsPutPreviousValues) HasFirstClosedDate() bool`

HasFirstClosedDate returns a boolean if a field has been set.

### GetAuditableEntityRiskId

`func (o *TaskItemsPutPreviousValues) GetAuditableEntityRiskId() int32`

GetAuditableEntityRiskId returns the AuditableEntityRiskId field if non-nil, zero value otherwise.

### GetAuditableEntityRiskIdOk

`func (o *TaskItemsPutPreviousValues) GetAuditableEntityRiskIdOk() (*int32, bool)`

GetAuditableEntityRiskIdOk returns a tuple with the AuditableEntityRiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditableEntityRiskId

`func (o *TaskItemsPutPreviousValues) SetAuditableEntityRiskId(v int32)`

SetAuditableEntityRiskId sets AuditableEntityRiskId field to given value.

### HasAuditableEntityRiskId

`func (o *TaskItemsPutPreviousValues) HasAuditableEntityRiskId() bool`

HasAuditableEntityRiskId returns a boolean if a field has been set.

### GetType

`func (o *TaskItemsPutPreviousValues) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskItemsPutPreviousValues) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskItemsPutPreviousValues) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TaskItemsPutPreviousValues) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEsgTopicId

`func (o *TaskItemsPutPreviousValues) GetEsgTopicId() int32`

GetEsgTopicId returns the EsgTopicId field if non-nil, zero value otherwise.

### GetEsgTopicIdOk

`func (o *TaskItemsPutPreviousValues) GetEsgTopicIdOk() (*int32, bool)`

GetEsgTopicIdOk returns a tuple with the EsgTopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgTopicId

`func (o *TaskItemsPutPreviousValues) SetEsgTopicId(v int32)`

SetEsgTopicId sets EsgTopicId field to given value.

### HasEsgTopicId

`func (o *TaskItemsPutPreviousValues) HasEsgTopicId() bool`

HasEsgTopicId returns a boolean if a field has been set.

### GetCreateReason

`func (o *TaskItemsPutPreviousValues) GetCreateReason() string`

GetCreateReason returns the CreateReason field if non-nil, zero value otherwise.

### GetCreateReasonOk

`func (o *TaskItemsPutPreviousValues) GetCreateReasonOk() (*string, bool)`

GetCreateReasonOk returns a tuple with the CreateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateReason

`func (o *TaskItemsPutPreviousValues) SetCreateReason(v string)`

SetCreateReason sets CreateReason field to given value.

### HasCreateReason

`func (o *TaskItemsPutPreviousValues) HasCreateReason() bool`

HasCreateReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


