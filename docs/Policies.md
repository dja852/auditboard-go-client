# Policies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Uid** | **string** |  | 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**LatestDocumentVersionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;document_versions.id&#x60;.&lt;fk table&#x3D;&#39;document_versions&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**LastRevisionByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**LastRevisionDate** | Pointer to **string** |  | [optional] 
**PolicyTypeId** | **int32** | Note: This is a Foreign Key to &#x60;policy_types.id&#x60;.&lt;fk table&#x3D;&#39;policy_types&#39; column&#x3D;&#39;id&#39;/&gt; | 
**AdminUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IsSequentialApproval** | Pointer to **bool** |  | [optional] 
**FieldData** | Pointer to **interface{}** |  | [optional] 
**RecurringScheduleId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;recurring_schedules.id&#x60;.&lt;fk table&#x3D;&#39;recurring_schedules&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FutureApprovalDate** | Pointer to **string** |  | [optional] 
**ReminderDuration** | **string** |  | [default to "30 Days"]
**RecurrenceNotifiedAt** | Pointer to **string** |  | [optional] 
**LastApprovalDate** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPolicies

`func NewPolicies(uid string, name string, policyTypeId int32, reminderDuration string, ) *Policies`

NewPolicies instantiates a new Policies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoliciesWithDefaults

`func NewPoliciesWithDefaults() *Policies`

NewPoliciesWithDefaults instantiates a new Policies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Policies) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Policies) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Policies) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Policies) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *Policies) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Policies) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Policies) SetUid(v string)`

SetUid sets Uid field to given value.


### GetCreatedAt

`func (o *Policies) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Policies) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Policies) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Policies) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Policies) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Policies) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Policies) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Policies) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Policies) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Policies) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Policies) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Policies) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *Policies) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Policies) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Policies) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Policies) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Policies) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Policies) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Policies) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *Policies) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Policies) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Policies) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Policies) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetLatestDocumentVersionId

`func (o *Policies) GetLatestDocumentVersionId() int32`

GetLatestDocumentVersionId returns the LatestDocumentVersionId field if non-nil, zero value otherwise.

### GetLatestDocumentVersionIdOk

`func (o *Policies) GetLatestDocumentVersionIdOk() (*int32, bool)`

GetLatestDocumentVersionIdOk returns a tuple with the LatestDocumentVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDocumentVersionId

`func (o *Policies) SetLatestDocumentVersionId(v int32)`

SetLatestDocumentVersionId sets LatestDocumentVersionId field to given value.

### HasLatestDocumentVersionId

`func (o *Policies) HasLatestDocumentVersionId() bool`

HasLatestDocumentVersionId returns a boolean if a field has been set.

### GetLastRevisionByUserId

`func (o *Policies) GetLastRevisionByUserId() int32`

GetLastRevisionByUserId returns the LastRevisionByUserId field if non-nil, zero value otherwise.

### GetLastRevisionByUserIdOk

`func (o *Policies) GetLastRevisionByUserIdOk() (*int32, bool)`

GetLastRevisionByUserIdOk returns a tuple with the LastRevisionByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRevisionByUserId

`func (o *Policies) SetLastRevisionByUserId(v int32)`

SetLastRevisionByUserId sets LastRevisionByUserId field to given value.

### HasLastRevisionByUserId

`func (o *Policies) HasLastRevisionByUserId() bool`

HasLastRevisionByUserId returns a boolean if a field has been set.

### GetLastRevisionDate

`func (o *Policies) GetLastRevisionDate() string`

GetLastRevisionDate returns the LastRevisionDate field if non-nil, zero value otherwise.

### GetLastRevisionDateOk

`func (o *Policies) GetLastRevisionDateOk() (*string, bool)`

GetLastRevisionDateOk returns a tuple with the LastRevisionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRevisionDate

`func (o *Policies) SetLastRevisionDate(v string)`

SetLastRevisionDate sets LastRevisionDate field to given value.

### HasLastRevisionDate

`func (o *Policies) HasLastRevisionDate() bool`

HasLastRevisionDate returns a boolean if a field has been set.

### GetPolicyTypeId

`func (o *Policies) GetPolicyTypeId() int32`

GetPolicyTypeId returns the PolicyTypeId field if non-nil, zero value otherwise.

### GetPolicyTypeIdOk

`func (o *Policies) GetPolicyTypeIdOk() (*int32, bool)`

GetPolicyTypeIdOk returns a tuple with the PolicyTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTypeId

`func (o *Policies) SetPolicyTypeId(v int32)`

SetPolicyTypeId sets PolicyTypeId field to given value.


### GetAdminUserId

`func (o *Policies) GetAdminUserId() int32`

GetAdminUserId returns the AdminUserId field if non-nil, zero value otherwise.

### GetAdminUserIdOk

`func (o *Policies) GetAdminUserIdOk() (*int32, bool)`

GetAdminUserIdOk returns a tuple with the AdminUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUserId

`func (o *Policies) SetAdminUserId(v int32)`

SetAdminUserId sets AdminUserId field to given value.

### HasAdminUserId

`func (o *Policies) HasAdminUserId() bool`

HasAdminUserId returns a boolean if a field has been set.

### GetIsSequentialApproval

`func (o *Policies) GetIsSequentialApproval() bool`

GetIsSequentialApproval returns the IsSequentialApproval field if non-nil, zero value otherwise.

### GetIsSequentialApprovalOk

`func (o *Policies) GetIsSequentialApprovalOk() (*bool, bool)`

GetIsSequentialApprovalOk returns a tuple with the IsSequentialApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSequentialApproval

`func (o *Policies) SetIsSequentialApproval(v bool)`

SetIsSequentialApproval sets IsSequentialApproval field to given value.

### HasIsSequentialApproval

`func (o *Policies) HasIsSequentialApproval() bool`

HasIsSequentialApproval returns a boolean if a field has been set.

### GetFieldData

`func (o *Policies) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *Policies) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *Policies) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *Policies) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *Policies) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *Policies) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil
### GetRecurringScheduleId

`func (o *Policies) GetRecurringScheduleId() int32`

GetRecurringScheduleId returns the RecurringScheduleId field if non-nil, zero value otherwise.

### GetRecurringScheduleIdOk

`func (o *Policies) GetRecurringScheduleIdOk() (*int32, bool)`

GetRecurringScheduleIdOk returns a tuple with the RecurringScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringScheduleId

`func (o *Policies) SetRecurringScheduleId(v int32)`

SetRecurringScheduleId sets RecurringScheduleId field to given value.

### HasRecurringScheduleId

`func (o *Policies) HasRecurringScheduleId() bool`

HasRecurringScheduleId returns a boolean if a field has been set.

### GetFutureApprovalDate

`func (o *Policies) GetFutureApprovalDate() string`

GetFutureApprovalDate returns the FutureApprovalDate field if non-nil, zero value otherwise.

### GetFutureApprovalDateOk

`func (o *Policies) GetFutureApprovalDateOk() (*string, bool)`

GetFutureApprovalDateOk returns a tuple with the FutureApprovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureApprovalDate

`func (o *Policies) SetFutureApprovalDate(v string)`

SetFutureApprovalDate sets FutureApprovalDate field to given value.

### HasFutureApprovalDate

`func (o *Policies) HasFutureApprovalDate() bool`

HasFutureApprovalDate returns a boolean if a field has been set.

### GetReminderDuration

`func (o *Policies) GetReminderDuration() string`

GetReminderDuration returns the ReminderDuration field if non-nil, zero value otherwise.

### GetReminderDurationOk

`func (o *Policies) GetReminderDurationOk() (*string, bool)`

GetReminderDurationOk returns a tuple with the ReminderDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderDuration

`func (o *Policies) SetReminderDuration(v string)`

SetReminderDuration sets ReminderDuration field to given value.


### GetRecurrenceNotifiedAt

`func (o *Policies) GetRecurrenceNotifiedAt() string`

GetRecurrenceNotifiedAt returns the RecurrenceNotifiedAt field if non-nil, zero value otherwise.

### GetRecurrenceNotifiedAtOk

`func (o *Policies) GetRecurrenceNotifiedAtOk() (*string, bool)`

GetRecurrenceNotifiedAtOk returns a tuple with the RecurrenceNotifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceNotifiedAt

`func (o *Policies) SetRecurrenceNotifiedAt(v string)`

SetRecurrenceNotifiedAt sets RecurrenceNotifiedAt field to given value.

### HasRecurrenceNotifiedAt

`func (o *Policies) HasRecurrenceNotifiedAt() bool`

HasRecurrenceNotifiedAt returns a boolean if a field has been set.

### GetLastApprovalDate

`func (o *Policies) GetLastApprovalDate() string`

GetLastApprovalDate returns the LastApprovalDate field if non-nil, zero value otherwise.

### GetLastApprovalDateOk

`func (o *Policies) GetLastApprovalDateOk() (*string, bool)`

GetLastApprovalDateOk returns a tuple with the LastApprovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastApprovalDate

`func (o *Policies) SetLastApprovalDate(v string)`

SetLastApprovalDate sets LastApprovalDate field to given value.

### HasLastApprovalDate

`func (o *Policies) HasLastApprovalDate() bool`

HasLastApprovalDate returns a boolean if a field has been set.

### GetScopes

`func (o *Policies) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *Policies) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *Policies) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *Policies) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *Policies) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *Policies) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


