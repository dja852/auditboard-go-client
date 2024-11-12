# PoliciesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**LatestDocumentVersionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;document_versions.id&#x60;.&lt;fk table&#x3D;&#39;document_versions&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**LastRevisionByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**LastRevisionDate** | Pointer to **string** |  | [optional] 
**PolicyTypeId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;policy_types.id&#x60;.&lt;fk table&#x3D;&#39;policy_types&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AdminUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IsSequentialApproval** | Pointer to **bool** |  | [optional] 
**FieldData** | Pointer to **interface{}** |  | [optional] 
**RecurringScheduleId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;recurring_schedules.id&#x60;.&lt;fk table&#x3D;&#39;recurring_schedules&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FutureApprovalDate** | Pointer to **string** |  | [optional] 
**ReminderDuration** | Pointer to **string** |  | [optional] [default to "30 Days"]
**RecurrenceNotifiedAt** | Pointer to **string** |  | [optional] 
**LastApprovalDate** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPoliciesPutPreviousValues

`func NewPoliciesPutPreviousValues() *PoliciesPutPreviousValues`

NewPoliciesPutPreviousValues instantiates a new PoliciesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoliciesPutPreviousValuesWithDefaults

`func NewPoliciesPutPreviousValuesWithDefaults() *PoliciesPutPreviousValues`

NewPoliciesPutPreviousValuesWithDefaults instantiates a new PoliciesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PoliciesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoliciesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoliciesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PoliciesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *PoliciesPutPreviousValues) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *PoliciesPutPreviousValues) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *PoliciesPutPreviousValues) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *PoliciesPutPreviousValues) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PoliciesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PoliciesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PoliciesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PoliciesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PoliciesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PoliciesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PoliciesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PoliciesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *PoliciesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *PoliciesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *PoliciesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *PoliciesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *PoliciesPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoliciesPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoliciesPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PoliciesPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PoliciesPutPreviousValues) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PoliciesPutPreviousValues) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PoliciesPutPreviousValues) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PoliciesPutPreviousValues) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *PoliciesPutPreviousValues) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PoliciesPutPreviousValues) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PoliciesPutPreviousValues) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PoliciesPutPreviousValues) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetLatestDocumentVersionId

`func (o *PoliciesPutPreviousValues) GetLatestDocumentVersionId() int32`

GetLatestDocumentVersionId returns the LatestDocumentVersionId field if non-nil, zero value otherwise.

### GetLatestDocumentVersionIdOk

`func (o *PoliciesPutPreviousValues) GetLatestDocumentVersionIdOk() (*int32, bool)`

GetLatestDocumentVersionIdOk returns a tuple with the LatestDocumentVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDocumentVersionId

`func (o *PoliciesPutPreviousValues) SetLatestDocumentVersionId(v int32)`

SetLatestDocumentVersionId sets LatestDocumentVersionId field to given value.

### HasLatestDocumentVersionId

`func (o *PoliciesPutPreviousValues) HasLatestDocumentVersionId() bool`

HasLatestDocumentVersionId returns a boolean if a field has been set.

### GetLastRevisionByUserId

`func (o *PoliciesPutPreviousValues) GetLastRevisionByUserId() int32`

GetLastRevisionByUserId returns the LastRevisionByUserId field if non-nil, zero value otherwise.

### GetLastRevisionByUserIdOk

`func (o *PoliciesPutPreviousValues) GetLastRevisionByUserIdOk() (*int32, bool)`

GetLastRevisionByUserIdOk returns a tuple with the LastRevisionByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRevisionByUserId

`func (o *PoliciesPutPreviousValues) SetLastRevisionByUserId(v int32)`

SetLastRevisionByUserId sets LastRevisionByUserId field to given value.

### HasLastRevisionByUserId

`func (o *PoliciesPutPreviousValues) HasLastRevisionByUserId() bool`

HasLastRevisionByUserId returns a boolean if a field has been set.

### GetLastRevisionDate

`func (o *PoliciesPutPreviousValues) GetLastRevisionDate() string`

GetLastRevisionDate returns the LastRevisionDate field if non-nil, zero value otherwise.

### GetLastRevisionDateOk

`func (o *PoliciesPutPreviousValues) GetLastRevisionDateOk() (*string, bool)`

GetLastRevisionDateOk returns a tuple with the LastRevisionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRevisionDate

`func (o *PoliciesPutPreviousValues) SetLastRevisionDate(v string)`

SetLastRevisionDate sets LastRevisionDate field to given value.

### HasLastRevisionDate

`func (o *PoliciesPutPreviousValues) HasLastRevisionDate() bool`

HasLastRevisionDate returns a boolean if a field has been set.

### GetPolicyTypeId

`func (o *PoliciesPutPreviousValues) GetPolicyTypeId() int32`

GetPolicyTypeId returns the PolicyTypeId field if non-nil, zero value otherwise.

### GetPolicyTypeIdOk

`func (o *PoliciesPutPreviousValues) GetPolicyTypeIdOk() (*int32, bool)`

GetPolicyTypeIdOk returns a tuple with the PolicyTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTypeId

`func (o *PoliciesPutPreviousValues) SetPolicyTypeId(v int32)`

SetPolicyTypeId sets PolicyTypeId field to given value.

### HasPolicyTypeId

`func (o *PoliciesPutPreviousValues) HasPolicyTypeId() bool`

HasPolicyTypeId returns a boolean if a field has been set.

### GetAdminUserId

`func (o *PoliciesPutPreviousValues) GetAdminUserId() int32`

GetAdminUserId returns the AdminUserId field if non-nil, zero value otherwise.

### GetAdminUserIdOk

`func (o *PoliciesPutPreviousValues) GetAdminUserIdOk() (*int32, bool)`

GetAdminUserIdOk returns a tuple with the AdminUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUserId

`func (o *PoliciesPutPreviousValues) SetAdminUserId(v int32)`

SetAdminUserId sets AdminUserId field to given value.

### HasAdminUserId

`func (o *PoliciesPutPreviousValues) HasAdminUserId() bool`

HasAdminUserId returns a boolean if a field has been set.

### GetIsSequentialApproval

`func (o *PoliciesPutPreviousValues) GetIsSequentialApproval() bool`

GetIsSequentialApproval returns the IsSequentialApproval field if non-nil, zero value otherwise.

### GetIsSequentialApprovalOk

`func (o *PoliciesPutPreviousValues) GetIsSequentialApprovalOk() (*bool, bool)`

GetIsSequentialApprovalOk returns a tuple with the IsSequentialApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSequentialApproval

`func (o *PoliciesPutPreviousValues) SetIsSequentialApproval(v bool)`

SetIsSequentialApproval sets IsSequentialApproval field to given value.

### HasIsSequentialApproval

`func (o *PoliciesPutPreviousValues) HasIsSequentialApproval() bool`

HasIsSequentialApproval returns a boolean if a field has been set.

### GetFieldData

`func (o *PoliciesPutPreviousValues) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *PoliciesPutPreviousValues) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *PoliciesPutPreviousValues) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *PoliciesPutPreviousValues) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *PoliciesPutPreviousValues) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *PoliciesPutPreviousValues) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil
### GetRecurringScheduleId

`func (o *PoliciesPutPreviousValues) GetRecurringScheduleId() int32`

GetRecurringScheduleId returns the RecurringScheduleId field if non-nil, zero value otherwise.

### GetRecurringScheduleIdOk

`func (o *PoliciesPutPreviousValues) GetRecurringScheduleIdOk() (*int32, bool)`

GetRecurringScheduleIdOk returns a tuple with the RecurringScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringScheduleId

`func (o *PoliciesPutPreviousValues) SetRecurringScheduleId(v int32)`

SetRecurringScheduleId sets RecurringScheduleId field to given value.

### HasRecurringScheduleId

`func (o *PoliciesPutPreviousValues) HasRecurringScheduleId() bool`

HasRecurringScheduleId returns a boolean if a field has been set.

### GetFutureApprovalDate

`func (o *PoliciesPutPreviousValues) GetFutureApprovalDate() string`

GetFutureApprovalDate returns the FutureApprovalDate field if non-nil, zero value otherwise.

### GetFutureApprovalDateOk

`func (o *PoliciesPutPreviousValues) GetFutureApprovalDateOk() (*string, bool)`

GetFutureApprovalDateOk returns a tuple with the FutureApprovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureApprovalDate

`func (o *PoliciesPutPreviousValues) SetFutureApprovalDate(v string)`

SetFutureApprovalDate sets FutureApprovalDate field to given value.

### HasFutureApprovalDate

`func (o *PoliciesPutPreviousValues) HasFutureApprovalDate() bool`

HasFutureApprovalDate returns a boolean if a field has been set.

### GetReminderDuration

`func (o *PoliciesPutPreviousValues) GetReminderDuration() string`

GetReminderDuration returns the ReminderDuration field if non-nil, zero value otherwise.

### GetReminderDurationOk

`func (o *PoliciesPutPreviousValues) GetReminderDurationOk() (*string, bool)`

GetReminderDurationOk returns a tuple with the ReminderDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderDuration

`func (o *PoliciesPutPreviousValues) SetReminderDuration(v string)`

SetReminderDuration sets ReminderDuration field to given value.

### HasReminderDuration

`func (o *PoliciesPutPreviousValues) HasReminderDuration() bool`

HasReminderDuration returns a boolean if a field has been set.

### GetRecurrenceNotifiedAt

`func (o *PoliciesPutPreviousValues) GetRecurrenceNotifiedAt() string`

GetRecurrenceNotifiedAt returns the RecurrenceNotifiedAt field if non-nil, zero value otherwise.

### GetRecurrenceNotifiedAtOk

`func (o *PoliciesPutPreviousValues) GetRecurrenceNotifiedAtOk() (*string, bool)`

GetRecurrenceNotifiedAtOk returns a tuple with the RecurrenceNotifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceNotifiedAt

`func (o *PoliciesPutPreviousValues) SetRecurrenceNotifiedAt(v string)`

SetRecurrenceNotifiedAt sets RecurrenceNotifiedAt field to given value.

### HasRecurrenceNotifiedAt

`func (o *PoliciesPutPreviousValues) HasRecurrenceNotifiedAt() bool`

HasRecurrenceNotifiedAt returns a boolean if a field has been set.

### GetLastApprovalDate

`func (o *PoliciesPutPreviousValues) GetLastApprovalDate() string`

GetLastApprovalDate returns the LastApprovalDate field if non-nil, zero value otherwise.

### GetLastApprovalDateOk

`func (o *PoliciesPutPreviousValues) GetLastApprovalDateOk() (*string, bool)`

GetLastApprovalDateOk returns a tuple with the LastApprovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastApprovalDate

`func (o *PoliciesPutPreviousValues) SetLastApprovalDate(v string)`

SetLastApprovalDate sets LastApprovalDate field to given value.

### HasLastApprovalDate

`func (o *PoliciesPutPreviousValues) HasLastApprovalDate() bool`

HasLastApprovalDate returns a boolean if a field has been set.

### GetScopes

`func (o *PoliciesPutPreviousValues) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PoliciesPutPreviousValues) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PoliciesPutPreviousValues) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PoliciesPutPreviousValues) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *PoliciesPutPreviousValues) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *PoliciesPutPreviousValues) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


