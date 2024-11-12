# RcwProjectsPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
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
**Type** | Pointer to **string** |  | [optional] [default to "Generic"]
**IsDraft** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewRcwProjectsPutPreviousValues

`func NewRcwProjectsPutPreviousValues() *RcwProjectsPutPreviousValues`

NewRcwProjectsPutPreviousValues instantiates a new RcwProjectsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRcwProjectsPutPreviousValuesWithDefaults

`func NewRcwProjectsPutPreviousValuesWithDefaults() *RcwProjectsPutPreviousValues`

NewRcwProjectsPutPreviousValuesWithDefaults instantiates a new RcwProjectsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RcwProjectsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RcwProjectsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RcwProjectsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RcwProjectsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RcwProjectsPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RcwProjectsPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RcwProjectsPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RcwProjectsPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *RcwProjectsPutPreviousValues) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RcwProjectsPutPreviousValues) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RcwProjectsPutPreviousValues) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RcwProjectsPutPreviousValues) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *RcwProjectsPutPreviousValues) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *RcwProjectsPutPreviousValues) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *RcwProjectsPutPreviousValues) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *RcwProjectsPutPreviousValues) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RcwProjectsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RcwProjectsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RcwProjectsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RcwProjectsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RcwProjectsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RcwProjectsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RcwProjectsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RcwProjectsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *RcwProjectsPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *RcwProjectsPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *RcwProjectsPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *RcwProjectsPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetLaunchDate

`func (o *RcwProjectsPutPreviousValues) GetLaunchDate() string`

GetLaunchDate returns the LaunchDate field if non-nil, zero value otherwise.

### GetLaunchDateOk

`func (o *RcwProjectsPutPreviousValues) GetLaunchDateOk() (*string, bool)`

GetLaunchDateOk returns a tuple with the LaunchDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchDate

`func (o *RcwProjectsPutPreviousValues) SetLaunchDate(v string)`

SetLaunchDate sets LaunchDate field to given value.

### HasLaunchDate

`func (o *RcwProjectsPutPreviousValues) HasLaunchDate() bool`

HasLaunchDate returns a boolean if a field has been set.

### GetLaunchByUserId

`func (o *RcwProjectsPutPreviousValues) GetLaunchByUserId() int32`

GetLaunchByUserId returns the LaunchByUserId field if non-nil, zero value otherwise.

### GetLaunchByUserIdOk

`func (o *RcwProjectsPutPreviousValues) GetLaunchByUserIdOk() (*int32, bool)`

GetLaunchByUserIdOk returns a tuple with the LaunchByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchByUserId

`func (o *RcwProjectsPutPreviousValues) SetLaunchByUserId(v int32)`

SetLaunchByUserId sets LaunchByUserId field to given value.

### HasLaunchByUserId

`func (o *RcwProjectsPutPreviousValues) HasLaunchByUserId() bool`

HasLaunchByUserId returns a boolean if a field has been set.

### GetCompletedDate

`func (o *RcwProjectsPutPreviousValues) GetCompletedDate() string`

GetCompletedDate returns the CompletedDate field if non-nil, zero value otherwise.

### GetCompletedDateOk

`func (o *RcwProjectsPutPreviousValues) GetCompletedDateOk() (*string, bool)`

GetCompletedDateOk returns a tuple with the CompletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDate

`func (o *RcwProjectsPutPreviousValues) SetCompletedDate(v string)`

SetCompletedDate sets CompletedDate field to given value.

### HasCompletedDate

`func (o *RcwProjectsPutPreviousValues) HasCompletedDate() bool`

HasCompletedDate returns a boolean if a field has been set.

### GetCompletedByUserId

`func (o *RcwProjectsPutPreviousValues) GetCompletedByUserId() int32`

GetCompletedByUserId returns the CompletedByUserId field if non-nil, zero value otherwise.

### GetCompletedByUserIdOk

`func (o *RcwProjectsPutPreviousValues) GetCompletedByUserIdOk() (*int32, bool)`

GetCompletedByUserIdOk returns a tuple with the CompletedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedByUserId

`func (o *RcwProjectsPutPreviousValues) SetCompletedByUserId(v int32)`

SetCompletedByUserId sets CompletedByUserId field to given value.

### HasCompletedByUserId

`func (o *RcwProjectsPutPreviousValues) HasCompletedByUserId() bool`

HasCompletedByUserId returns a boolean if a field has been set.

### GetDueDate

`func (o *RcwProjectsPutPreviousValues) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *RcwProjectsPutPreviousValues) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *RcwProjectsPutPreviousValues) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *RcwProjectsPutPreviousValues) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetPeriodStart

`func (o *RcwProjectsPutPreviousValues) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *RcwProjectsPutPreviousValues) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *RcwProjectsPutPreviousValues) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *RcwProjectsPutPreviousValues) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *RcwProjectsPutPreviousValues) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *RcwProjectsPutPreviousValues) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *RcwProjectsPutPreviousValues) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *RcwProjectsPutPreviousValues) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetTotals

`func (o *RcwProjectsPutPreviousValues) GetTotals() interface{}`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *RcwProjectsPutPreviousValues) GetTotalsOk() (*interface{}, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *RcwProjectsPutPreviousValues) SetTotals(v interface{})`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *RcwProjectsPutPreviousValues) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### SetTotalsNil

`func (o *RcwProjectsPutPreviousValues) SetTotalsNil(b bool)`

 SetTotalsNil sets the value for Totals to be an explicit nil

### UnsetTotals
`func (o *RcwProjectsPutPreviousValues) UnsetTotals()`

UnsetTotals ensures that no value is present for Totals, not even an explicit nil
### GetType

`func (o *RcwProjectsPutPreviousValues) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RcwProjectsPutPreviousValues) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RcwProjectsPutPreviousValues) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RcwProjectsPutPreviousValues) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsDraft

`func (o *RcwProjectsPutPreviousValues) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *RcwProjectsPutPreviousValues) GetIsDraftOk() (*bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDraft

`func (o *RcwProjectsPutPreviousValues) SetIsDraft(v bool)`

SetIsDraft sets IsDraft field to given value.

### HasIsDraft

`func (o *RcwProjectsPutPreviousValues) HasIsDraft() bool`

HasIsDraft returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


