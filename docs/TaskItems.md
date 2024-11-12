# TaskItems

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
**TaskGroupId** | **int32** | Note: This is a Foreign Key to &#x60;task_groups.id&#x60;.&lt;fk table&#x3D;&#39;task_groups&#39; column&#x3D;&#39;id&#39;/&gt; | 
**AssigneeUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**SubmittedByUserId** | Pointer to **int32** |  | [optional] 
**CertifiedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AssigneeResponse** | Pointer to **string** |  | [optional] 
**BudgetHours** | **float32** |  | [default to 0]
**ActualHours** | **float32** |  | [default to 0]
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

### NewTaskItems

`func NewTaskItems(taskGroupId int32, budgetHours float32, actualHours float32, ) *TaskItems`

NewTaskItems instantiates a new TaskItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskItemsWithDefaults

`func NewTaskItemsWithDefaults() *TaskItems`

NewTaskItemsWithDefaults instantiates a new TaskItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskItems) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskItems) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskItems) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TaskItems) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskItems) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskItems) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskItems) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaskItems) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TaskItems) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskItems) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskItems) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TaskItems) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *TaskItems) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *TaskItems) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *TaskItems) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *TaskItems) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetIsAcknowledgePrepared

`func (o *TaskItems) GetIsAcknowledgePrepared() bool`

GetIsAcknowledgePrepared returns the IsAcknowledgePrepared field if non-nil, zero value otherwise.

### GetIsAcknowledgePreparedOk

`func (o *TaskItems) GetIsAcknowledgePreparedOk() (*bool, bool)`

GetIsAcknowledgePreparedOk returns a tuple with the IsAcknowledgePrepared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAcknowledgePrepared

`func (o *TaskItems) SetIsAcknowledgePrepared(v bool)`

SetIsAcknowledgePrepared sets IsAcknowledgePrepared field to given value.

### HasIsAcknowledgePrepared

`func (o *TaskItems) HasIsAcknowledgePrepared() bool`

HasIsAcknowledgePrepared returns a boolean if a field has been set.

### GetDueDate

`func (o *TaskItems) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *TaskItems) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *TaskItems) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *TaskItems) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetSubmittedDate

`func (o *TaskItems) GetSubmittedDate() string`

GetSubmittedDate returns the SubmittedDate field if non-nil, zero value otherwise.

### GetSubmittedDateOk

`func (o *TaskItems) GetSubmittedDateOk() (*string, bool)`

GetSubmittedDateOk returns a tuple with the SubmittedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedDate

`func (o *TaskItems) SetSubmittedDate(v string)`

SetSubmittedDate sets SubmittedDate field to given value.

### HasSubmittedDate

`func (o *TaskItems) HasSubmittedDate() bool`

HasSubmittedDate returns a boolean if a field has been set.

### GetCertifiedDate

`func (o *TaskItems) GetCertifiedDate() string`

GetCertifiedDate returns the CertifiedDate field if non-nil, zero value otherwise.

### GetCertifiedDateOk

`func (o *TaskItems) GetCertifiedDateOk() (*string, bool)`

GetCertifiedDateOk returns a tuple with the CertifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifiedDate

`func (o *TaskItems) SetCertifiedDate(v string)`

SetCertifiedDate sets CertifiedDate field to given value.

### HasCertifiedDate

`func (o *TaskItems) HasCertifiedDate() bool`

HasCertifiedDate returns a boolean if a field has been set.

### GetOrder

`func (o *TaskItems) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TaskItems) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TaskItems) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TaskItems) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetTaskGroupId

`func (o *TaskItems) GetTaskGroupId() int32`

GetTaskGroupId returns the TaskGroupId field if non-nil, zero value otherwise.

### GetTaskGroupIdOk

`func (o *TaskItems) GetTaskGroupIdOk() (*int32, bool)`

GetTaskGroupIdOk returns a tuple with the TaskGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskGroupId

`func (o *TaskItems) SetTaskGroupId(v int32)`

SetTaskGroupId sets TaskGroupId field to given value.


### GetAssigneeUserId

`func (o *TaskItems) GetAssigneeUserId() int32`

GetAssigneeUserId returns the AssigneeUserId field if non-nil, zero value otherwise.

### GetAssigneeUserIdOk

`func (o *TaskItems) GetAssigneeUserIdOk() (*int32, bool)`

GetAssigneeUserIdOk returns a tuple with the AssigneeUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeUserId

`func (o *TaskItems) SetAssigneeUserId(v int32)`

SetAssigneeUserId sets AssigneeUserId field to given value.

### HasAssigneeUserId

`func (o *TaskItems) HasAssigneeUserId() bool`

HasAssigneeUserId returns a boolean if a field has been set.

### GetSubmittedByUserId

`func (o *TaskItems) GetSubmittedByUserId() int32`

GetSubmittedByUserId returns the SubmittedByUserId field if non-nil, zero value otherwise.

### GetSubmittedByUserIdOk

`func (o *TaskItems) GetSubmittedByUserIdOk() (*int32, bool)`

GetSubmittedByUserIdOk returns a tuple with the SubmittedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedByUserId

`func (o *TaskItems) SetSubmittedByUserId(v int32)`

SetSubmittedByUserId sets SubmittedByUserId field to given value.

### HasSubmittedByUserId

`func (o *TaskItems) HasSubmittedByUserId() bool`

HasSubmittedByUserId returns a boolean if a field has been set.

### GetCertifiedByUserId

`func (o *TaskItems) GetCertifiedByUserId() int32`

GetCertifiedByUserId returns the CertifiedByUserId field if non-nil, zero value otherwise.

### GetCertifiedByUserIdOk

`func (o *TaskItems) GetCertifiedByUserIdOk() (*int32, bool)`

GetCertifiedByUserIdOk returns a tuple with the CertifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifiedByUserId

`func (o *TaskItems) SetCertifiedByUserId(v int32)`

SetCertifiedByUserId sets CertifiedByUserId field to given value.

### HasCertifiedByUserId

`func (o *TaskItems) HasCertifiedByUserId() bool`

HasCertifiedByUserId returns a boolean if a field has been set.

### GetTitle

`func (o *TaskItems) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TaskItems) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TaskItems) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TaskItems) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *TaskItems) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskItems) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskItems) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskItems) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAssigneeResponse

`func (o *TaskItems) GetAssigneeResponse() string`

GetAssigneeResponse returns the AssigneeResponse field if non-nil, zero value otherwise.

### GetAssigneeResponseOk

`func (o *TaskItems) GetAssigneeResponseOk() (*string, bool)`

GetAssigneeResponseOk returns a tuple with the AssigneeResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeResponse

`func (o *TaskItems) SetAssigneeResponse(v string)`

SetAssigneeResponse sets AssigneeResponse field to given value.

### HasAssigneeResponse

`func (o *TaskItems) HasAssigneeResponse() bool`

HasAssigneeResponse returns a boolean if a field has been set.

### GetBudgetHours

`func (o *TaskItems) GetBudgetHours() float32`

GetBudgetHours returns the BudgetHours field if non-nil, zero value otherwise.

### GetBudgetHoursOk

`func (o *TaskItems) GetBudgetHoursOk() (*float32, bool)`

GetBudgetHoursOk returns a tuple with the BudgetHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetHours

`func (o *TaskItems) SetBudgetHours(v float32)`

SetBudgetHours sets BudgetHours field to given value.


### GetActualHours

`func (o *TaskItems) GetActualHours() float32`

GetActualHours returns the ActualHours field if non-nil, zero value otherwise.

### GetActualHoursOk

`func (o *TaskItems) GetActualHoursOk() (*float32, bool)`

GetActualHoursOk returns a tuple with the ActualHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualHours

`func (o *TaskItems) SetActualHours(v float32)`

SetActualHours sets ActualHours field to given value.


### GetTaskPriorityId

`func (o *TaskItems) GetTaskPriorityId() int32`

GetTaskPriorityId returns the TaskPriorityId field if non-nil, zero value otherwise.

### GetTaskPriorityIdOk

`func (o *TaskItems) GetTaskPriorityIdOk() (*int32, bool)`

GetTaskPriorityIdOk returns a tuple with the TaskPriorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPriorityId

`func (o *TaskItems) SetTaskPriorityId(v int32)`

SetTaskPriorityId sets TaskPriorityId field to given value.

### HasTaskPriorityId

`func (o *TaskItems) HasTaskPriorityId() bool`

HasTaskPriorityId returns a boolean if a field has been set.

### GetTaskPeriodId

`func (o *TaskItems) GetTaskPeriodId() int32`

GetTaskPeriodId returns the TaskPeriodId field if non-nil, zero value otherwise.

### GetTaskPeriodIdOk

`func (o *TaskItems) GetTaskPeriodIdOk() (*int32, bool)`

GetTaskPeriodIdOk returns a tuple with the TaskPeriodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPeriodId

`func (o *TaskItems) SetTaskPeriodId(v int32)`

SetTaskPeriodId sets TaskPeriodId field to given value.

### HasTaskPeriodId

`func (o *TaskItems) HasTaskPeriodId() bool`

HasTaskPeriodId returns a boolean if a field has been set.

### GetReviewerUserId

`func (o *TaskItems) GetReviewerUserId() int32`

GetReviewerUserId returns the ReviewerUserId field if non-nil, zero value otherwise.

### GetReviewerUserIdOk

`func (o *TaskItems) GetReviewerUserIdOk() (*int32, bool)`

GetReviewerUserIdOk returns a tuple with the ReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerUserId

`func (o *TaskItems) SetReviewerUserId(v int32)`

SetReviewerUserId sets ReviewerUserId field to given value.

### HasReviewerUserId

`func (o *TaskItems) HasReviewerUserId() bool`

HasReviewerUserId returns a boolean if a field has been set.

### GetLocationId

`func (o *TaskItems) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *TaskItems) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *TaskItems) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *TaskItems) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetDepartmentId

`func (o *TaskItems) GetDepartmentId() int32`

GetDepartmentId returns the DepartmentId field if non-nil, zero value otherwise.

### GetDepartmentIdOk

`func (o *TaskItems) GetDepartmentIdOk() (*int32, bool)`

GetDepartmentIdOk returns a tuple with the DepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentId

`func (o *TaskItems) SetDepartmentId(v int32)`

SetDepartmentId sets DepartmentId field to given value.

### HasDepartmentId

`func (o *TaskItems) HasDepartmentId() bool`

HasDepartmentId returns a boolean if a field has been set.

### GetReferenceNotes

`func (o *TaskItems) GetReferenceNotes() string`

GetReferenceNotes returns the ReferenceNotes field if non-nil, zero value otherwise.

### GetReferenceNotesOk

`func (o *TaskItems) GetReferenceNotesOk() (*string, bool)`

GetReferenceNotesOk returns a tuple with the ReferenceNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNotes

`func (o *TaskItems) SetReferenceNotes(v string)`

SetReferenceNotes sets ReferenceNotes field to given value.

### HasReferenceNotes

`func (o *TaskItems) HasReferenceNotes() bool`

HasReferenceNotes returns a boolean if a field has been set.

### GetControlsDatumId

`func (o *TaskItems) GetControlsDatumId() int32`

GetControlsDatumId returns the ControlsDatumId field if non-nil, zero value otherwise.

### GetControlsDatumIdOk

`func (o *TaskItems) GetControlsDatumIdOk() (*int32, bool)`

GetControlsDatumIdOk returns a tuple with the ControlsDatumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlsDatumId

`func (o *TaskItems) SetControlsDatumId(v int32)`

SetControlsDatumId sets ControlsDatumId field to given value.

### HasControlsDatumId

`func (o *TaskItems) HasControlsDatumId() bool`

HasControlsDatumId returns a boolean if a field has been set.

### GetTaskCategoryId

`func (o *TaskItems) GetTaskCategoryId() int32`

GetTaskCategoryId returns the TaskCategoryId field if non-nil, zero value otherwise.

### GetTaskCategoryIdOk

`func (o *TaskItems) GetTaskCategoryIdOk() (*int32, bool)`

GetTaskCategoryIdOk returns a tuple with the TaskCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCategoryId

`func (o *TaskItems) SetTaskCategoryId(v int32)`

SetTaskCategoryId sets TaskCategoryId field to given value.

### HasTaskCategoryId

`func (o *TaskItems) HasTaskCategoryId() bool`

HasTaskCategoryId returns a boolean if a field has been set.

### GetCreatorUserId

`func (o *TaskItems) GetCreatorUserId() int32`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *TaskItems) GetCreatorUserIdOk() (*int32, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *TaskItems) SetCreatorUserId(v int32)`

SetCreatorUserId sets CreatorUserId field to given value.

### HasCreatorUserId

`func (o *TaskItems) HasCreatorUserId() bool`

HasCreatorUserId returns a boolean if a field has been set.

### GetMetaConfig

`func (o *TaskItems) GetMetaConfig() string`

GetMetaConfig returns the MetaConfig field if non-nil, zero value otherwise.

### GetMetaConfigOk

`func (o *TaskItems) GetMetaConfigOk() (*string, bool)`

GetMetaConfigOk returns a tuple with the MetaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaConfig

`func (o *TaskItems) SetMetaConfig(v string)`

SetMetaConfig sets MetaConfig field to given value.

### HasMetaConfig

`func (o *TaskItems) HasMetaConfig() bool`

HasMetaConfig returns a boolean if a field has been set.

### GetStatus

`func (o *TaskItems) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskItems) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskItems) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskItems) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTestSectionId

`func (o *TaskItems) GetTestSectionId() int32`

GetTestSectionId returns the TestSectionId field if non-nil, zero value otherwise.

### GetTestSectionIdOk

`func (o *TaskItems) GetTestSectionIdOk() (*int32, bool)`

GetTestSectionIdOk returns a tuple with the TestSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSectionId

`func (o *TaskItems) SetTestSectionId(v int32)`

SetTestSectionId sets TestSectionId field to given value.

### HasTestSectionId

`func (o *TaskItems) HasTestSectionId() bool`

HasTestSectionId returns a boolean if a field has been set.

### GetIssueId

`func (o *TaskItems) GetIssueId() int32`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *TaskItems) GetIssueIdOk() (*int32, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *TaskItems) SetIssueId(v int32)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *TaskItems) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.

### GetOpenedDate

`func (o *TaskItems) GetOpenedDate() string`

GetOpenedDate returns the OpenedDate field if non-nil, zero value otherwise.

### GetOpenedDateOk

`func (o *TaskItems) GetOpenedDateOk() (*string, bool)`

GetOpenedDateOk returns a tuple with the OpenedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedDate

`func (o *TaskItems) SetOpenedDate(v string)`

SetOpenedDate sets OpenedDate field to given value.

### HasOpenedDate

`func (o *TaskItems) HasOpenedDate() bool`

HasOpenedDate returns a boolean if a field has been set.

### GetOpenedByUserId

`func (o *TaskItems) GetOpenedByUserId() int32`

GetOpenedByUserId returns the OpenedByUserId field if non-nil, zero value otherwise.

### GetOpenedByUserIdOk

`func (o *TaskItems) GetOpenedByUserIdOk() (*int32, bool)`

GetOpenedByUserIdOk returns a tuple with the OpenedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedByUserId

`func (o *TaskItems) SetOpenedByUserId(v int32)`

SetOpenedByUserId sets OpenedByUserId field to given value.

### HasOpenedByUserId

`func (o *TaskItems) HasOpenedByUserId() bool`

HasOpenedByUserId returns a boolean if a field has been set.

### GetClosedDate

`func (o *TaskItems) GetClosedDate() string`

GetClosedDate returns the ClosedDate field if non-nil, zero value otherwise.

### GetClosedDateOk

`func (o *TaskItems) GetClosedDateOk() (*string, bool)`

GetClosedDateOk returns a tuple with the ClosedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedDate

`func (o *TaskItems) SetClosedDate(v string)`

SetClosedDate sets ClosedDate field to given value.

### HasClosedDate

`func (o *TaskItems) HasClosedDate() bool`

HasClosedDate returns a boolean if a field has been set.

### GetClosedByUserId

`func (o *TaskItems) GetClosedByUserId() int32`

GetClosedByUserId returns the ClosedByUserId field if non-nil, zero value otherwise.

### GetClosedByUserIdOk

`func (o *TaskItems) GetClosedByUserIdOk() (*int32, bool)`

GetClosedByUserIdOk returns a tuple with the ClosedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedByUserId

`func (o *TaskItems) SetClosedByUserId(v int32)`

SetClosedByUserId sets ClosedByUserId field to given value.

### HasClosedByUserId

`func (o *TaskItems) HasClosedByUserId() bool`

HasClosedByUserId returns a boolean if a field has been set.

### GetOpsAuditId

`func (o *TaskItems) GetOpsAuditId() int32`

GetOpsAuditId returns the OpsAuditId field if non-nil, zero value otherwise.

### GetOpsAuditIdOk

`func (o *TaskItems) GetOpsAuditIdOk() (*int32, bool)`

GetOpsAuditIdOk returns a tuple with the OpsAuditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditId

`func (o *TaskItems) SetOpsAuditId(v int32)`

SetOpsAuditId sets OpsAuditId field to given value.

### HasOpsAuditId

`func (o *TaskItems) HasOpsAuditId() bool`

HasOpsAuditId returns a boolean if a field has been set.

### GetOpsAuditItemId

`func (o *TaskItems) GetOpsAuditItemId() int32`

GetOpsAuditItemId returns the OpsAuditItemId field if non-nil, zero value otherwise.

### GetOpsAuditItemIdOk

`func (o *TaskItems) GetOpsAuditItemIdOk() (*int32, bool)`

GetOpsAuditItemIdOk returns a tuple with the OpsAuditItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditItemId

`func (o *TaskItems) SetOpsAuditItemId(v int32)`

SetOpsAuditItemId sets OpsAuditItemId field to given value.

### HasOpsAuditItemId

`func (o *TaskItems) HasOpsAuditItemId() bool`

HasOpsAuditItemId returns a boolean if a field has been set.

### GetActionPlanId

`func (o *TaskItems) GetActionPlanId() int32`

GetActionPlanId returns the ActionPlanId field if non-nil, zero value otherwise.

### GetActionPlanIdOk

`func (o *TaskItems) GetActionPlanIdOk() (*int32, bool)`

GetActionPlanIdOk returns a tuple with the ActionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPlanId

`func (o *TaskItems) SetActionPlanId(v int32)`

SetActionPlanId sets ActionPlanId field to given value.

### HasActionPlanId

`func (o *TaskItems) HasActionPlanId() bool`

HasActionPlanId returns a boolean if a field has been set.

### GetRiskId

`func (o *TaskItems) GetRiskId() int32`

GetRiskId returns the RiskId field if non-nil, zero value otherwise.

### GetRiskIdOk

`func (o *TaskItems) GetRiskIdOk() (*int32, bool)`

GetRiskIdOk returns a tuple with the RiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskId

`func (o *TaskItems) SetRiskId(v int32)`

SetRiskId sets RiskId field to given value.

### HasRiskId

`func (o *TaskItems) HasRiskId() bool`

HasRiskId returns a boolean if a field has been set.

### GetExternalIntegrationUrl

`func (o *TaskItems) GetExternalIntegrationUrl() string`

GetExternalIntegrationUrl returns the ExternalIntegrationUrl field if non-nil, zero value otherwise.

### GetExternalIntegrationUrlOk

`func (o *TaskItems) GetExternalIntegrationUrlOk() (*string, bool)`

GetExternalIntegrationUrlOk returns a tuple with the ExternalIntegrationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIntegrationUrl

`func (o *TaskItems) SetExternalIntegrationUrl(v string)`

SetExternalIntegrationUrl sets ExternalIntegrationUrl field to given value.

### HasExternalIntegrationUrl

`func (o *TaskItems) HasExternalIntegrationUrl() bool`

HasExternalIntegrationUrl returns a boolean if a field has been set.

### GetFrameworkItemId

`func (o *TaskItems) GetFrameworkItemId() int32`

GetFrameworkItemId returns the FrameworkItemId field if non-nil, zero value otherwise.

### GetFrameworkItemIdOk

`func (o *TaskItems) GetFrameworkItemIdOk() (*int32, bool)`

GetFrameworkItemIdOk returns a tuple with the FrameworkItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemId

`func (o *TaskItems) SetFrameworkItemId(v int32)`

SetFrameworkItemId sets FrameworkItemId field to given value.

### HasFrameworkItemId

`func (o *TaskItems) HasFrameworkItemId() bool`

HasFrameworkItemId returns a boolean if a field has been set.

### GetHistoricalTaskIds

`func (o *TaskItems) GetHistoricalTaskIds() interface{}`

GetHistoricalTaskIds returns the HistoricalTaskIds field if non-nil, zero value otherwise.

### GetHistoricalTaskIdsOk

`func (o *TaskItems) GetHistoricalTaskIdsOk() (*interface{}, bool)`

GetHistoricalTaskIdsOk returns a tuple with the HistoricalTaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricalTaskIds

`func (o *TaskItems) SetHistoricalTaskIds(v interface{})`

SetHistoricalTaskIds sets HistoricalTaskIds field to given value.

### HasHistoricalTaskIds

`func (o *TaskItems) HasHistoricalTaskIds() bool`

HasHistoricalTaskIds returns a boolean if a field has been set.

### SetHistoricalTaskIdsNil

`func (o *TaskItems) SetHistoricalTaskIdsNil(b bool)`

 SetHistoricalTaskIdsNil sets the value for HistoricalTaskIds to be an explicit nil

### UnsetHistoricalTaskIds
`func (o *TaskItems) UnsetHistoricalTaskIds()`

UnsetHistoricalTaskIds ensures that no value is present for HistoricalTaskIds, not even an explicit nil
### GetNarrativeId

`func (o *TaskItems) GetNarrativeId() int32`

GetNarrativeId returns the NarrativeId field if non-nil, zero value otherwise.

### GetNarrativeIdOk

`func (o *TaskItems) GetNarrativeIdOk() (*int32, bool)`

GetNarrativeIdOk returns a tuple with the NarrativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarrativeId

`func (o *TaskItems) SetNarrativeId(v int32)`

SetNarrativeId sets NarrativeId field to given value.

### HasNarrativeId

`func (o *TaskItems) HasNarrativeId() bool`

HasNarrativeId returns a boolean if a field has been set.

### GetAuditableEntityId

`func (o *TaskItems) GetAuditableEntityId() int32`

GetAuditableEntityId returns the AuditableEntityId field if non-nil, zero value otherwise.

### GetAuditableEntityIdOk

`func (o *TaskItems) GetAuditableEntityIdOk() (*int32, bool)`

GetAuditableEntityIdOk returns a tuple with the AuditableEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditableEntityId

`func (o *TaskItems) SetAuditableEntityId(v int32)`

SetAuditableEntityId sets AuditableEntityId field to given value.

### HasAuditableEntityId

`func (o *TaskItems) HasAuditableEntityId() bool`

HasAuditableEntityId returns a boolean if a field has been set.

### GetPolicyId

`func (o *TaskItems) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *TaskItems) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *TaskItems) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *TaskItems) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetFirstSubmittedDate

`func (o *TaskItems) GetFirstSubmittedDate() string`

GetFirstSubmittedDate returns the FirstSubmittedDate field if non-nil, zero value otherwise.

### GetFirstSubmittedDateOk

`func (o *TaskItems) GetFirstSubmittedDateOk() (*string, bool)`

GetFirstSubmittedDateOk returns a tuple with the FirstSubmittedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSubmittedDate

`func (o *TaskItems) SetFirstSubmittedDate(v string)`

SetFirstSubmittedDate sets FirstSubmittedDate field to given value.

### HasFirstSubmittedDate

`func (o *TaskItems) HasFirstSubmittedDate() bool`

HasFirstSubmittedDate returns a boolean if a field has been set.

### GetFirstCertifiedDate

`func (o *TaskItems) GetFirstCertifiedDate() string`

GetFirstCertifiedDate returns the FirstCertifiedDate field if non-nil, zero value otherwise.

### GetFirstCertifiedDateOk

`func (o *TaskItems) GetFirstCertifiedDateOk() (*string, bool)`

GetFirstCertifiedDateOk returns a tuple with the FirstCertifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCertifiedDate

`func (o *TaskItems) SetFirstCertifiedDate(v string)`

SetFirstCertifiedDate sets FirstCertifiedDate field to given value.

### HasFirstCertifiedDate

`func (o *TaskItems) HasFirstCertifiedDate() bool`

HasFirstCertifiedDate returns a boolean if a field has been set.

### GetFirstClosedDate

`func (o *TaskItems) GetFirstClosedDate() string`

GetFirstClosedDate returns the FirstClosedDate field if non-nil, zero value otherwise.

### GetFirstClosedDateOk

`func (o *TaskItems) GetFirstClosedDateOk() (*string, bool)`

GetFirstClosedDateOk returns a tuple with the FirstClosedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstClosedDate

`func (o *TaskItems) SetFirstClosedDate(v string)`

SetFirstClosedDate sets FirstClosedDate field to given value.

### HasFirstClosedDate

`func (o *TaskItems) HasFirstClosedDate() bool`

HasFirstClosedDate returns a boolean if a field has been set.

### GetAuditableEntityRiskId

`func (o *TaskItems) GetAuditableEntityRiskId() int32`

GetAuditableEntityRiskId returns the AuditableEntityRiskId field if non-nil, zero value otherwise.

### GetAuditableEntityRiskIdOk

`func (o *TaskItems) GetAuditableEntityRiskIdOk() (*int32, bool)`

GetAuditableEntityRiskIdOk returns a tuple with the AuditableEntityRiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditableEntityRiskId

`func (o *TaskItems) SetAuditableEntityRiskId(v int32)`

SetAuditableEntityRiskId sets AuditableEntityRiskId field to given value.

### HasAuditableEntityRiskId

`func (o *TaskItems) HasAuditableEntityRiskId() bool`

HasAuditableEntityRiskId returns a boolean if a field has been set.

### GetType

`func (o *TaskItems) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskItems) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskItems) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TaskItems) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEsgTopicId

`func (o *TaskItems) GetEsgTopicId() int32`

GetEsgTopicId returns the EsgTopicId field if non-nil, zero value otherwise.

### GetEsgTopicIdOk

`func (o *TaskItems) GetEsgTopicIdOk() (*int32, bool)`

GetEsgTopicIdOk returns a tuple with the EsgTopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgTopicId

`func (o *TaskItems) SetEsgTopicId(v int32)`

SetEsgTopicId sets EsgTopicId field to given value.

### HasEsgTopicId

`func (o *TaskItems) HasEsgTopicId() bool`

HasEsgTopicId returns a boolean if a field has been set.

### GetCreateReason

`func (o *TaskItems) GetCreateReason() string`

GetCreateReason returns the CreateReason field if non-nil, zero value otherwise.

### GetCreateReasonOk

`func (o *TaskItems) GetCreateReasonOk() (*string, bool)`

GetCreateReasonOk returns a tuple with the CreateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateReason

`func (o *TaskItems) SetCreateReason(v string)`

SetCreateReason sets CreateReason field to given value.

### HasCreateReason

`func (o *TaskItems) HasCreateReason() bool`

HasCreateReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


