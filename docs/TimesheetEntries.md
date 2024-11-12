# TimesheetEntries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Date** | **string** |  | 
**Hours** | **float32** |  | 
**UserId** | **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | 
**TimesheetableType** | **string** |  | 
**TimesheetableId** | **int32** |  | 
**TimesheetableMeta** | Pointer to **string** |  | [optional] 
**TimesheetProjectId** | **int32** | Note: This is a Foreign Key to &#x60;timesheet_projects.id&#x60;.&lt;fk table&#x3D;&#39;timesheet_projects&#39; column&#x3D;&#39;id&#39;/&gt; | 
**Comment** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewTimesheetEntries

`func NewTimesheetEntries(date string, hours float32, userId int32, timesheetableType string, timesheetableId int32, timesheetProjectId int32, ) *TimesheetEntries`

NewTimesheetEntries instantiates a new TimesheetEntries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimesheetEntriesWithDefaults

`func NewTimesheetEntriesWithDefaults() *TimesheetEntries`

NewTimesheetEntriesWithDefaults instantiates a new TimesheetEntries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimesheetEntries) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimesheetEntries) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimesheetEntries) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TimesheetEntries) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDate

`func (o *TimesheetEntries) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TimesheetEntries) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TimesheetEntries) SetDate(v string)`

SetDate sets Date field to given value.


### GetHours

`func (o *TimesheetEntries) GetHours() float32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *TimesheetEntries) GetHoursOk() (*float32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *TimesheetEntries) SetHours(v float32)`

SetHours sets Hours field to given value.


### GetUserId

`func (o *TimesheetEntries) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TimesheetEntries) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TimesheetEntries) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetTimesheetableType

`func (o *TimesheetEntries) GetTimesheetableType() string`

GetTimesheetableType returns the TimesheetableType field if non-nil, zero value otherwise.

### GetTimesheetableTypeOk

`func (o *TimesheetEntries) GetTimesheetableTypeOk() (*string, bool)`

GetTimesheetableTypeOk returns a tuple with the TimesheetableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesheetableType

`func (o *TimesheetEntries) SetTimesheetableType(v string)`

SetTimesheetableType sets TimesheetableType field to given value.


### GetTimesheetableId

`func (o *TimesheetEntries) GetTimesheetableId() int32`

GetTimesheetableId returns the TimesheetableId field if non-nil, zero value otherwise.

### GetTimesheetableIdOk

`func (o *TimesheetEntries) GetTimesheetableIdOk() (*int32, bool)`

GetTimesheetableIdOk returns a tuple with the TimesheetableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesheetableId

`func (o *TimesheetEntries) SetTimesheetableId(v int32)`

SetTimesheetableId sets TimesheetableId field to given value.


### GetTimesheetableMeta

`func (o *TimesheetEntries) GetTimesheetableMeta() string`

GetTimesheetableMeta returns the TimesheetableMeta field if non-nil, zero value otherwise.

### GetTimesheetableMetaOk

`func (o *TimesheetEntries) GetTimesheetableMetaOk() (*string, bool)`

GetTimesheetableMetaOk returns a tuple with the TimesheetableMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesheetableMeta

`func (o *TimesheetEntries) SetTimesheetableMeta(v string)`

SetTimesheetableMeta sets TimesheetableMeta field to given value.

### HasTimesheetableMeta

`func (o *TimesheetEntries) HasTimesheetableMeta() bool`

HasTimesheetableMeta returns a boolean if a field has been set.

### GetTimesheetProjectId

`func (o *TimesheetEntries) GetTimesheetProjectId() int32`

GetTimesheetProjectId returns the TimesheetProjectId field if non-nil, zero value otherwise.

### GetTimesheetProjectIdOk

`func (o *TimesheetEntries) GetTimesheetProjectIdOk() (*int32, bool)`

GetTimesheetProjectIdOk returns a tuple with the TimesheetProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesheetProjectId

`func (o *TimesheetEntries) SetTimesheetProjectId(v int32)`

SetTimesheetProjectId sets TimesheetProjectId field to given value.


### GetComment

`func (o *TimesheetEntries) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TimesheetEntries) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TimesheetEntries) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TimesheetEntries) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDeletedAt

`func (o *TimesheetEntries) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *TimesheetEntries) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *TimesheetEntries) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *TimesheetEntries) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TimesheetEntries) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TimesheetEntries) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TimesheetEntries) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TimesheetEntries) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TimesheetEntries) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TimesheetEntries) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TimesheetEntries) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TimesheetEntries) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


