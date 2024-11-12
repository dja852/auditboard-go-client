# RcwProjects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Name** | **string** |  | 
**Status** | **string** |  | 
**CreatedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**LaunchDate** | Pointer to **string** |  | [optional] 
**LaunchByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CompletedDate** | Pointer to **string** |  | [optional] 
**CompletedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**PeriodStart** | Pointer to **string** |  | [optional] 
**PeriodEnd** | Pointer to **string** |  | [optional] 
**Totals** | Pointer to **interface{}** |  | [optional] 
**Type** | **string** |  | [default to "Generic"]
**IsDraft** | **bool** |  | [default to true]

## Methods

### NewRcwProjects

`func NewRcwProjects(name string, status string, type_ string, isDraft bool, ) *RcwProjects`

NewRcwProjects instantiates a new RcwProjects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRcwProjectsWithDefaults

`func NewRcwProjectsWithDefaults() *RcwProjects`

NewRcwProjectsWithDefaults instantiates a new RcwProjects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RcwProjects) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RcwProjects) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RcwProjects) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RcwProjects) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RcwProjects) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RcwProjects) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RcwProjects) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *RcwProjects) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RcwProjects) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RcwProjects) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedByUserId

`func (o *RcwProjects) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *RcwProjects) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *RcwProjects) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *RcwProjects) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RcwProjects) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RcwProjects) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RcwProjects) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RcwProjects) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RcwProjects) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RcwProjects) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RcwProjects) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RcwProjects) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *RcwProjects) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *RcwProjects) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *RcwProjects) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *RcwProjects) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetLaunchDate

`func (o *RcwProjects) GetLaunchDate() string`

GetLaunchDate returns the LaunchDate field if non-nil, zero value otherwise.

### GetLaunchDateOk

`func (o *RcwProjects) GetLaunchDateOk() (*string, bool)`

GetLaunchDateOk returns a tuple with the LaunchDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchDate

`func (o *RcwProjects) SetLaunchDate(v string)`

SetLaunchDate sets LaunchDate field to given value.

### HasLaunchDate

`func (o *RcwProjects) HasLaunchDate() bool`

HasLaunchDate returns a boolean if a field has been set.

### GetLaunchByUserId

`func (o *RcwProjects) GetLaunchByUserId() int32`

GetLaunchByUserId returns the LaunchByUserId field if non-nil, zero value otherwise.

### GetLaunchByUserIdOk

`func (o *RcwProjects) GetLaunchByUserIdOk() (*int32, bool)`

GetLaunchByUserIdOk returns a tuple with the LaunchByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchByUserId

`func (o *RcwProjects) SetLaunchByUserId(v int32)`

SetLaunchByUserId sets LaunchByUserId field to given value.

### HasLaunchByUserId

`func (o *RcwProjects) HasLaunchByUserId() bool`

HasLaunchByUserId returns a boolean if a field has been set.

### GetCompletedDate

`func (o *RcwProjects) GetCompletedDate() string`

GetCompletedDate returns the CompletedDate field if non-nil, zero value otherwise.

### GetCompletedDateOk

`func (o *RcwProjects) GetCompletedDateOk() (*string, bool)`

GetCompletedDateOk returns a tuple with the CompletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDate

`func (o *RcwProjects) SetCompletedDate(v string)`

SetCompletedDate sets CompletedDate field to given value.

### HasCompletedDate

`func (o *RcwProjects) HasCompletedDate() bool`

HasCompletedDate returns a boolean if a field has been set.

### GetCompletedByUserId

`func (o *RcwProjects) GetCompletedByUserId() int32`

GetCompletedByUserId returns the CompletedByUserId field if non-nil, zero value otherwise.

### GetCompletedByUserIdOk

`func (o *RcwProjects) GetCompletedByUserIdOk() (*int32, bool)`

GetCompletedByUserIdOk returns a tuple with the CompletedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedByUserId

`func (o *RcwProjects) SetCompletedByUserId(v int32)`

SetCompletedByUserId sets CompletedByUserId field to given value.

### HasCompletedByUserId

`func (o *RcwProjects) HasCompletedByUserId() bool`

HasCompletedByUserId returns a boolean if a field has been set.

### GetDueDate

`func (o *RcwProjects) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *RcwProjects) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *RcwProjects) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *RcwProjects) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetPeriodStart

`func (o *RcwProjects) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *RcwProjects) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *RcwProjects) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *RcwProjects) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *RcwProjects) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *RcwProjects) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *RcwProjects) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *RcwProjects) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetTotals

`func (o *RcwProjects) GetTotals() interface{}`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *RcwProjects) GetTotalsOk() (*interface{}, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *RcwProjects) SetTotals(v interface{})`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *RcwProjects) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### SetTotalsNil

`func (o *RcwProjects) SetTotalsNil(b bool)`

 SetTotalsNil sets the value for Totals to be an explicit nil

### UnsetTotals
`func (o *RcwProjects) UnsetTotals()`

UnsetTotals ensures that no value is present for Totals, not even an explicit nil
### GetType

`func (o *RcwProjects) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RcwProjects) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RcwProjects) SetType(v string)`

SetType sets Type field to given value.


### GetIsDraft

`func (o *RcwProjects) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *RcwProjects) GetIsDraftOk() (*bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDraft

`func (o *RcwProjects) SetIsDraft(v bool)`

SetIsDraft sets IsDraft field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


