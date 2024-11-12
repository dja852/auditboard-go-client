# TimesheetEntriesPutTimesheetEntry

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

### NewTimesheetEntriesPutTimesheetEntry

`func NewTimesheetEntriesPutTimesheetEntry(date string, hours float32, userId int32, timesheetableType string, timesheetableId int32, timesheetProjectId int32, ) *TimesheetEntriesPutTimesheetEntry`

NewTimesheetEntriesPutTimesheetEntry instantiates a new TimesheetEntriesPutTimesheetEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimesheetEntriesPutTimesheetEntryWithDefaults

`func NewTimesheetEntriesPutTimesheetEntryWithDefaults() *TimesheetEntriesPutTimesheetEntry`

NewTimesheetEntriesPutTimesheetEntryWithDefaults instantiates a new TimesheetEntriesPutTimesheetEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimesheetEntriesPutTimesheetEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimesheetEntriesPutTimesheetEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimesheetEntriesPutTimesheetEntry) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TimesheetEntriesPutTimesheetEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDate

`func (o *TimesheetEntriesPutTimesheetEntry) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TimesheetEntriesPutTimesheetEntry) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TimesheetEntriesPutTimesheetEntry) SetDate(v string)`

SetDate sets Date field to given value.


### GetHours

`func (o *TimesheetEntriesPutTimesheetEntry) GetHours() float32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *TimesheetEntriesPutTimesheetEntry) GetHoursOk() (*float32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *TimesheetEntriesPutTimesheetEntry) SetHours(v float32)`

SetHours sets Hours field to given value.


### GetUserId

`func (o *TimesheetEntriesPutTimesheetEntry) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TimesheetEntriesPutTimesheetEntry) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TimesheetEntriesPutTimesheetEntry) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetTimesheetableType

`func (o *TimesheetEntriesPutTimesheetEntry) GetTimesheetableType() string`

GetTimesheetableType returns the TimesheetableType field if non-nil, zero value otherwise.

### GetTimesheetableTypeOk

`func (o *TimesheetEntriesPutTimesheetEntry) GetTimesheetableTypeOk() (*string, bool)`

GetTimesheetableTypeOk returns a tuple with the TimesheetableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesheetableType

`func (o *TimesheetEntriesPutTimesheetEntry) SetTimesheetableType(v string)`

SetTimesheetableType sets TimesheetableType field to given value.


### GetTimesheetableId

`func (o *TimesheetEntriesPutTimesheetEntry) GetTimesheetableId() int32`

GetTimesheetableId returns the TimesheetableId field if non-nil, zero value otherwise.

### GetTimesheetableIdOk

`func (o *TimesheetEntriesPutTimesheetEntry) GetTimesheetableIdOk() (*int32, bool)`

GetTimesheetableIdOk returns a tuple with the TimesheetableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesheetableId

`func (o *TimesheetEntriesPutTimesheetEntry) SetTimesheetableId(v int32)`

SetTimesheetableId sets TimesheetableId field to given value.


### GetTimesheetableMeta

`func (o *TimesheetEntriesPutTimesheetEntry) GetTimesheetableMeta() string`

GetTimesheetableMeta returns the TimesheetableMeta field if non-nil, zero value otherwise.

### GetTimesheetableMetaOk

`func (o *TimesheetEntriesPutTimesheetEntry) GetTimesheetableMetaOk() (*string, bool)`

GetTimesheetableMetaOk returns a tuple with the TimesheetableMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesheetableMeta

`func (o *TimesheetEntriesPutTimesheetEntry) SetTimesheetableMeta(v string)`

SetTimesheetableMeta sets TimesheetableMeta field to given value.

### HasTimesheetableMeta

`func (o *TimesheetEntriesPutTimesheetEntry) HasTimesheetableMeta() bool`

HasTimesheetableMeta returns a boolean if a field has been set.

### GetTimesheetProjectId

`func (o *TimesheetEntriesPutTimesheetEntry) GetTimesheetProjectId() int32`

GetTimesheetProjectId returns the TimesheetProjectId field if non-nil, zero value otherwise.

### GetTimesheetProjectIdOk

`func (o *TimesheetEntriesPutTimesheetEntry) GetTimesheetProjectIdOk() (*int32, bool)`

GetTimesheetProjectIdOk returns a tuple with the TimesheetProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesheetProjectId

`func (o *TimesheetEntriesPutTimesheetEntry) SetTimesheetProjectId(v int32)`

SetTimesheetProjectId sets TimesheetProjectId field to given value.


### GetComment

`func (o *TimesheetEntriesPutTimesheetEntry) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TimesheetEntriesPutTimesheetEntry) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TimesheetEntriesPutTimesheetEntry) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TimesheetEntriesPutTimesheetEntry) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDeletedAt

`func (o *TimesheetEntriesPutTimesheetEntry) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *TimesheetEntriesPutTimesheetEntry) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *TimesheetEntriesPutTimesheetEntry) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *TimesheetEntriesPutTimesheetEntry) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TimesheetEntriesPutTimesheetEntry) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TimesheetEntriesPutTimesheetEntry) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TimesheetEntriesPutTimesheetEntry) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TimesheetEntriesPutTimesheetEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TimesheetEntriesPutTimesheetEntry) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TimesheetEntriesPutTimesheetEntry) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TimesheetEntriesPutTimesheetEntry) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TimesheetEntriesPutTimesheetEntry) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


