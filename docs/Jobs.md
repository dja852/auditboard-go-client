# Jobs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**StartedAt** | Pointer to **string** |  | [optional] 
**CompletedAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Output** | Pointer to **interface{}** |  | [optional] 
**TaskArguments** | Pointer to **interface{}** |  | [optional] 
**StackTrace** | Pointer to **string** |  | [optional] 
**TaskName** | Pointer to **string** |  | [optional] 
**TriggeredByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**KueId** | Pointer to **string** |  | [optional] 
**CancelledAt** | Pointer to **string** |  | [optional] 
**CancelledByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 

## Methods

### NewJobs

`func NewJobs() *Jobs`

NewJobs instantiates a new Jobs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobsWithDefaults

`func NewJobsWithDefaults() *Jobs`

NewJobsWithDefaults instantiates a new Jobs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Jobs) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Jobs) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Jobs) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Jobs) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Jobs) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Jobs) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Jobs) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Jobs) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Jobs) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Jobs) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Jobs) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Jobs) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Jobs) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Jobs) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Jobs) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Jobs) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *Jobs) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Jobs) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Jobs) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Jobs) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *Jobs) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *Jobs) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *Jobs) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *Jobs) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetStatus

`func (o *Jobs) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Jobs) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Jobs) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Jobs) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOutput

`func (o *Jobs) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *Jobs) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *Jobs) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *Jobs) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *Jobs) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *Jobs) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetTaskArguments

`func (o *Jobs) GetTaskArguments() interface{}`

GetTaskArguments returns the TaskArguments field if non-nil, zero value otherwise.

### GetTaskArgumentsOk

`func (o *Jobs) GetTaskArgumentsOk() (*interface{}, bool)`

GetTaskArgumentsOk returns a tuple with the TaskArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskArguments

`func (o *Jobs) SetTaskArguments(v interface{})`

SetTaskArguments sets TaskArguments field to given value.

### HasTaskArguments

`func (o *Jobs) HasTaskArguments() bool`

HasTaskArguments returns a boolean if a field has been set.

### SetTaskArgumentsNil

`func (o *Jobs) SetTaskArgumentsNil(b bool)`

 SetTaskArgumentsNil sets the value for TaskArguments to be an explicit nil

### UnsetTaskArguments
`func (o *Jobs) UnsetTaskArguments()`

UnsetTaskArguments ensures that no value is present for TaskArguments, not even an explicit nil
### GetStackTrace

`func (o *Jobs) GetStackTrace() string`

GetStackTrace returns the StackTrace field if non-nil, zero value otherwise.

### GetStackTraceOk

`func (o *Jobs) GetStackTraceOk() (*string, bool)`

GetStackTraceOk returns a tuple with the StackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTrace

`func (o *Jobs) SetStackTrace(v string)`

SetStackTrace sets StackTrace field to given value.

### HasStackTrace

`func (o *Jobs) HasStackTrace() bool`

HasStackTrace returns a boolean if a field has been set.

### GetTaskName

`func (o *Jobs) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *Jobs) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *Jobs) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *Jobs) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### GetTriggeredByUserId

`func (o *Jobs) GetTriggeredByUserId() int32`

GetTriggeredByUserId returns the TriggeredByUserId field if non-nil, zero value otherwise.

### GetTriggeredByUserIdOk

`func (o *Jobs) GetTriggeredByUserIdOk() (*int32, bool)`

GetTriggeredByUserIdOk returns a tuple with the TriggeredByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredByUserId

`func (o *Jobs) SetTriggeredByUserId(v int32)`

SetTriggeredByUserId sets TriggeredByUserId field to given value.

### HasTriggeredByUserId

`func (o *Jobs) HasTriggeredByUserId() bool`

HasTriggeredByUserId returns a boolean if a field has been set.

### GetKueId

`func (o *Jobs) GetKueId() string`

GetKueId returns the KueId field if non-nil, zero value otherwise.

### GetKueIdOk

`func (o *Jobs) GetKueIdOk() (*string, bool)`

GetKueIdOk returns a tuple with the KueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKueId

`func (o *Jobs) SetKueId(v string)`

SetKueId sets KueId field to given value.

### HasKueId

`func (o *Jobs) HasKueId() bool`

HasKueId returns a boolean if a field has been set.

### GetCancelledAt

`func (o *Jobs) GetCancelledAt() string`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *Jobs) GetCancelledAtOk() (*string, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *Jobs) SetCancelledAt(v string)`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *Jobs) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### GetCancelledByUserId

`func (o *Jobs) GetCancelledByUserId() int32`

GetCancelledByUserId returns the CancelledByUserId field if non-nil, zero value otherwise.

### GetCancelledByUserIdOk

`func (o *Jobs) GetCancelledByUserIdOk() (*int32, bool)`

GetCancelledByUserIdOk returns a tuple with the CancelledByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledByUserId

`func (o *Jobs) SetCancelledByUserId(v int32)`

SetCancelledByUserId sets CancelledByUserId field to given value.

### HasCancelledByUserId

`func (o *Jobs) HasCancelledByUserId() bool`

HasCancelledByUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


