# TaskPeriodsPutTaskPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**SortOrder** | **int32** |  | [default to 0]
**ShortName** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 

## Methods

### NewTaskPeriodsPutTaskPeriod

`func NewTaskPeriodsPutTaskPeriod(name string, sortOrder int32, ) *TaskPeriodsPutTaskPeriod`

NewTaskPeriodsPutTaskPeriod instantiates a new TaskPeriodsPutTaskPeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskPeriodsPutTaskPeriodWithDefaults

`func NewTaskPeriodsPutTaskPeriodWithDefaults() *TaskPeriodsPutTaskPeriod`

NewTaskPeriodsPutTaskPeriodWithDefaults instantiates a new TaskPeriodsPutTaskPeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskPeriodsPutTaskPeriod) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskPeriodsPutTaskPeriod) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskPeriodsPutTaskPeriod) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TaskPeriodsPutTaskPeriod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskPeriodsPutTaskPeriod) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskPeriodsPutTaskPeriod) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskPeriodsPutTaskPeriod) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaskPeriodsPutTaskPeriod) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TaskPeriodsPutTaskPeriod) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskPeriodsPutTaskPeriod) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskPeriodsPutTaskPeriod) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TaskPeriodsPutTaskPeriod) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *TaskPeriodsPutTaskPeriod) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *TaskPeriodsPutTaskPeriod) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *TaskPeriodsPutTaskPeriod) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *TaskPeriodsPutTaskPeriod) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *TaskPeriodsPutTaskPeriod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskPeriodsPutTaskPeriod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskPeriodsPutTaskPeriod) SetName(v string)`

SetName sets Name field to given value.


### GetSortOrder

`func (o *TaskPeriodsPutTaskPeriod) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *TaskPeriodsPutTaskPeriod) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *TaskPeriodsPutTaskPeriod) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.


### GetShortName

`func (o *TaskPeriodsPutTaskPeriod) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *TaskPeriodsPutTaskPeriod) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *TaskPeriodsPutTaskPeriod) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *TaskPeriodsPutTaskPeriod) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetStartDate

`func (o *TaskPeriodsPutTaskPeriod) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TaskPeriodsPutTaskPeriod) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TaskPeriodsPutTaskPeriod) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TaskPeriodsPutTaskPeriod) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *TaskPeriodsPutTaskPeriod) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TaskPeriodsPutTaskPeriod) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TaskPeriodsPutTaskPeriod) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TaskPeriodsPutTaskPeriod) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


