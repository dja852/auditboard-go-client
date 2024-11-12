# EmailLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**ToUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FromUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IntendedRecipient** | Pointer to **string** |  | [optional] 
**ActualRecipient** | Pointer to **string** |  | [optional] 
**EmailType** | Pointer to **string** |  | [optional] 
**EmailCategory** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**ReasonType** | Pointer to **string** |  | [optional] 
**EmailBody** | Pointer to **string** |  | [optional] 
**EmailHeaders** | Pointer to **interface{}** |  | [optional] 
**EmailSubject** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**OpsAuditItemId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_items.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_items&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TaskItemId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;task_items.id&#x60;.&lt;fk table&#x3D;&#39;task_items&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issues.id&#x60;.&lt;fk table&#x3D;&#39;issues&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CommentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;comments.id&#x60;.&lt;fk table&#x3D;&#39;comments&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;tests.id&#x60;.&lt;fk table&#x3D;&#39;tests&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlsDatumId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;controls_data.id&#x60;.&lt;fk table&#x3D;&#39;controls_data&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ActionPlanId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;action_plans.id&#x60;.&lt;fk table&#x3D;&#39;action_plans&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 

## Methods

### NewEmailLogs

`func NewEmailLogs() *EmailLogs`

NewEmailLogs instantiates a new EmailLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailLogsWithDefaults

`func NewEmailLogsWithDefaults() *EmailLogs`

NewEmailLogsWithDefaults instantiates a new EmailLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailLogs) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailLogs) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailLogs) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EmailLogs) HasId() bool`

HasId returns a boolean if a field has been set.

### GetToUserId

`func (o *EmailLogs) GetToUserId() int32`

GetToUserId returns the ToUserId field if non-nil, zero value otherwise.

### GetToUserIdOk

`func (o *EmailLogs) GetToUserIdOk() (*int32, bool)`

GetToUserIdOk returns a tuple with the ToUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUserId

`func (o *EmailLogs) SetToUserId(v int32)`

SetToUserId sets ToUserId field to given value.

### HasToUserId

`func (o *EmailLogs) HasToUserId() bool`

HasToUserId returns a boolean if a field has been set.

### GetFromUserId

`func (o *EmailLogs) GetFromUserId() int32`

GetFromUserId returns the FromUserId field if non-nil, zero value otherwise.

### GetFromUserIdOk

`func (o *EmailLogs) GetFromUserIdOk() (*int32, bool)`

GetFromUserIdOk returns a tuple with the FromUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserId

`func (o *EmailLogs) SetFromUserId(v int32)`

SetFromUserId sets FromUserId field to given value.

### HasFromUserId

`func (o *EmailLogs) HasFromUserId() bool`

HasFromUserId returns a boolean if a field has been set.

### GetIntendedRecipient

`func (o *EmailLogs) GetIntendedRecipient() string`

GetIntendedRecipient returns the IntendedRecipient field if non-nil, zero value otherwise.

### GetIntendedRecipientOk

`func (o *EmailLogs) GetIntendedRecipientOk() (*string, bool)`

GetIntendedRecipientOk returns a tuple with the IntendedRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntendedRecipient

`func (o *EmailLogs) SetIntendedRecipient(v string)`

SetIntendedRecipient sets IntendedRecipient field to given value.

### HasIntendedRecipient

`func (o *EmailLogs) HasIntendedRecipient() bool`

HasIntendedRecipient returns a boolean if a field has been set.

### GetActualRecipient

`func (o *EmailLogs) GetActualRecipient() string`

GetActualRecipient returns the ActualRecipient field if non-nil, zero value otherwise.

### GetActualRecipientOk

`func (o *EmailLogs) GetActualRecipientOk() (*string, bool)`

GetActualRecipientOk returns a tuple with the ActualRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualRecipient

`func (o *EmailLogs) SetActualRecipient(v string)`

SetActualRecipient sets ActualRecipient field to given value.

### HasActualRecipient

`func (o *EmailLogs) HasActualRecipient() bool`

HasActualRecipient returns a boolean if a field has been set.

### GetEmailType

`func (o *EmailLogs) GetEmailType() string`

GetEmailType returns the EmailType field if non-nil, zero value otherwise.

### GetEmailTypeOk

`func (o *EmailLogs) GetEmailTypeOk() (*string, bool)`

GetEmailTypeOk returns a tuple with the EmailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailType

`func (o *EmailLogs) SetEmailType(v string)`

SetEmailType sets EmailType field to given value.

### HasEmailType

`func (o *EmailLogs) HasEmailType() bool`

HasEmailType returns a boolean if a field has been set.

### GetEmailCategory

`func (o *EmailLogs) GetEmailCategory() string`

GetEmailCategory returns the EmailCategory field if non-nil, zero value otherwise.

### GetEmailCategoryOk

`func (o *EmailLogs) GetEmailCategoryOk() (*string, bool)`

GetEmailCategoryOk returns a tuple with the EmailCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCategory

`func (o *EmailLogs) SetEmailCategory(v string)`

SetEmailCategory sets EmailCategory field to given value.

### HasEmailCategory

`func (o *EmailLogs) HasEmailCategory() bool`

HasEmailCategory returns a boolean if a field has been set.

### GetStatus

`func (o *EmailLogs) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailLogs) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailLogs) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailLogs) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *EmailLogs) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EmailLogs) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EmailLogs) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EmailLogs) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonType

`func (o *EmailLogs) GetReasonType() string`

GetReasonType returns the ReasonType field if non-nil, zero value otherwise.

### GetReasonTypeOk

`func (o *EmailLogs) GetReasonTypeOk() (*string, bool)`

GetReasonTypeOk returns a tuple with the ReasonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonType

`func (o *EmailLogs) SetReasonType(v string)`

SetReasonType sets ReasonType field to given value.

### HasReasonType

`func (o *EmailLogs) HasReasonType() bool`

HasReasonType returns a boolean if a field has been set.

### GetEmailBody

`func (o *EmailLogs) GetEmailBody() string`

GetEmailBody returns the EmailBody field if non-nil, zero value otherwise.

### GetEmailBodyOk

`func (o *EmailLogs) GetEmailBodyOk() (*string, bool)`

GetEmailBodyOk returns a tuple with the EmailBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailBody

`func (o *EmailLogs) SetEmailBody(v string)`

SetEmailBody sets EmailBody field to given value.

### HasEmailBody

`func (o *EmailLogs) HasEmailBody() bool`

HasEmailBody returns a boolean if a field has been set.

### GetEmailHeaders

`func (o *EmailLogs) GetEmailHeaders() interface{}`

GetEmailHeaders returns the EmailHeaders field if non-nil, zero value otherwise.

### GetEmailHeadersOk

`func (o *EmailLogs) GetEmailHeadersOk() (*interface{}, bool)`

GetEmailHeadersOk returns a tuple with the EmailHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailHeaders

`func (o *EmailLogs) SetEmailHeaders(v interface{})`

SetEmailHeaders sets EmailHeaders field to given value.

### HasEmailHeaders

`func (o *EmailLogs) HasEmailHeaders() bool`

HasEmailHeaders returns a boolean if a field has been set.

### SetEmailHeadersNil

`func (o *EmailLogs) SetEmailHeadersNil(b bool)`

 SetEmailHeadersNil sets the value for EmailHeaders to be an explicit nil

### UnsetEmailHeaders
`func (o *EmailLogs) UnsetEmailHeaders()`

UnsetEmailHeaders ensures that no value is present for EmailHeaders, not even an explicit nil
### GetEmailSubject

`func (o *EmailLogs) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *EmailLogs) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *EmailLogs) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *EmailLogs) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EmailLogs) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmailLogs) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmailLogs) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EmailLogs) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EmailLogs) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EmailLogs) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EmailLogs) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EmailLogs) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EmailLogs) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EmailLogs) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EmailLogs) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EmailLogs) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetOpsAuditItemId

`func (o *EmailLogs) GetOpsAuditItemId() int32`

GetOpsAuditItemId returns the OpsAuditItemId field if non-nil, zero value otherwise.

### GetOpsAuditItemIdOk

`func (o *EmailLogs) GetOpsAuditItemIdOk() (*int32, bool)`

GetOpsAuditItemIdOk returns a tuple with the OpsAuditItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditItemId

`func (o *EmailLogs) SetOpsAuditItemId(v int32)`

SetOpsAuditItemId sets OpsAuditItemId field to given value.

### HasOpsAuditItemId

`func (o *EmailLogs) HasOpsAuditItemId() bool`

HasOpsAuditItemId returns a boolean if a field has been set.

### GetTaskItemId

`func (o *EmailLogs) GetTaskItemId() int32`

GetTaskItemId returns the TaskItemId field if non-nil, zero value otherwise.

### GetTaskItemIdOk

`func (o *EmailLogs) GetTaskItemIdOk() (*int32, bool)`

GetTaskItemIdOk returns a tuple with the TaskItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskItemId

`func (o *EmailLogs) SetTaskItemId(v int32)`

SetTaskItemId sets TaskItemId field to given value.

### HasTaskItemId

`func (o *EmailLogs) HasTaskItemId() bool`

HasTaskItemId returns a boolean if a field has been set.

### GetIssueId

`func (o *EmailLogs) GetIssueId() int32`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *EmailLogs) GetIssueIdOk() (*int32, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *EmailLogs) SetIssueId(v int32)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *EmailLogs) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.

### GetCommentId

`func (o *EmailLogs) GetCommentId() int32`

GetCommentId returns the CommentId field if non-nil, zero value otherwise.

### GetCommentIdOk

`func (o *EmailLogs) GetCommentIdOk() (*int32, bool)`

GetCommentIdOk returns a tuple with the CommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentId

`func (o *EmailLogs) SetCommentId(v int32)`

SetCommentId sets CommentId field to given value.

### HasCommentId

`func (o *EmailLogs) HasCommentId() bool`

HasCommentId returns a boolean if a field has been set.

### GetTestId

`func (o *EmailLogs) GetTestId() int32`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *EmailLogs) GetTestIdOk() (*int32, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *EmailLogs) SetTestId(v int32)`

SetTestId sets TestId field to given value.

### HasTestId

`func (o *EmailLogs) HasTestId() bool`

HasTestId returns a boolean if a field has been set.

### GetControlsDatumId

`func (o *EmailLogs) GetControlsDatumId() int32`

GetControlsDatumId returns the ControlsDatumId field if non-nil, zero value otherwise.

### GetControlsDatumIdOk

`func (o *EmailLogs) GetControlsDatumIdOk() (*int32, bool)`

GetControlsDatumIdOk returns a tuple with the ControlsDatumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlsDatumId

`func (o *EmailLogs) SetControlsDatumId(v int32)`

SetControlsDatumId sets ControlsDatumId field to given value.

### HasControlsDatumId

`func (o *EmailLogs) HasControlsDatumId() bool`

HasControlsDatumId returns a boolean if a field has been set.

### GetActionPlanId

`func (o *EmailLogs) GetActionPlanId() int32`

GetActionPlanId returns the ActionPlanId field if non-nil, zero value otherwise.

### GetActionPlanIdOk

`func (o *EmailLogs) GetActionPlanIdOk() (*int32, bool)`

GetActionPlanIdOk returns a tuple with the ActionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPlanId

`func (o *EmailLogs) SetActionPlanId(v int32)`

SetActionPlanId sets ActionPlanId field to given value.

### HasActionPlanId

`func (o *EmailLogs) HasActionPlanId() bool`

HasActionPlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


