# TaskItemsSubmit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskItems** | Pointer to [**[]TaskItems**](TaskItems.md) |  | [optional] 
**TaskGroups** | Pointer to [**[]TaskGroups**](TaskGroups.md) |  | [optional] 
**GlobalAudits** | Pointer to [**[]GlobalAudits**](GlobalAudits.md) |  | [optional] 

## Methods

### NewTaskItemsSubmit

`func NewTaskItemsSubmit() *TaskItemsSubmit`

NewTaskItemsSubmit instantiates a new TaskItemsSubmit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskItemsSubmitWithDefaults

`func NewTaskItemsSubmitWithDefaults() *TaskItemsSubmit`

NewTaskItemsSubmitWithDefaults instantiates a new TaskItemsSubmit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskItems

`func (o *TaskItemsSubmit) GetTaskItems() []TaskItems`

GetTaskItems returns the TaskItems field if non-nil, zero value otherwise.

### GetTaskItemsOk

`func (o *TaskItemsSubmit) GetTaskItemsOk() (*[]TaskItems, bool)`

GetTaskItemsOk returns a tuple with the TaskItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskItems

`func (o *TaskItemsSubmit) SetTaskItems(v []TaskItems)`

SetTaskItems sets TaskItems field to given value.

### HasTaskItems

`func (o *TaskItemsSubmit) HasTaskItems() bool`

HasTaskItems returns a boolean if a field has been set.

### GetTaskGroups

`func (o *TaskItemsSubmit) GetTaskGroups() []TaskGroups`

GetTaskGroups returns the TaskGroups field if non-nil, zero value otherwise.

### GetTaskGroupsOk

`func (o *TaskItemsSubmit) GetTaskGroupsOk() (*[]TaskGroups, bool)`

GetTaskGroupsOk returns a tuple with the TaskGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskGroups

`func (o *TaskItemsSubmit) SetTaskGroups(v []TaskGroups)`

SetTaskGroups sets TaskGroups field to given value.

### HasTaskGroups

`func (o *TaskItemsSubmit) HasTaskGroups() bool`

HasTaskGroups returns a boolean if a field has been set.

### GetGlobalAudits

`func (o *TaskItemsSubmit) GetGlobalAudits() []GlobalAudits`

GetGlobalAudits returns the GlobalAudits field if non-nil, zero value otherwise.

### GetGlobalAuditsOk

`func (o *TaskItemsSubmit) GetGlobalAuditsOk() (*[]GlobalAudits, bool)`

GetGlobalAuditsOk returns a tuple with the GlobalAudits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAudits

`func (o *TaskItemsSubmit) SetGlobalAudits(v []GlobalAudits)`

SetGlobalAudits sets GlobalAudits field to given value.

### HasGlobalAudits

`func (o *TaskItemsSubmit) HasGlobalAudits() bool`

HasGlobalAudits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


