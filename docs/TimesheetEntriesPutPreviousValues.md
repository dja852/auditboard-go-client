# TimesheetEntriesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Date** | Pointer to **string** |  | [optional] 
**Hours** | Pointer to **float32** |  | [optional] 
**UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TimesheetableType** | Pointer to **string** |  | [optional] 
**TimesheetableId** | Pointer to **int32** |  | [optional] 
**TimesheetableMeta** | Pointer to **string** |  | [optional] 
**TimesheetProjectId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;timesheet_projects.id&#x60;.&lt;fk table&#x3D;&#39;timesheet_projects&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewTimesheetEntriesPutPreviousValues

`func NewTimesheetEntriesPutPreviousValues() *TimesheetEntriesPutPreviousValues`

NewTimesheetEntriesPutPreviousValues instantiates a new TimesheetEntriesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimesheetEntriesPutPreviousValuesWithDefaults

`func NewTimesheetEntriesPutPreviousValuesWithDefaults() *TimesheetEntriesPutPreviousValues`

NewTimesheetEntriesPutPreviousValuesWithDefaults instantiates a new TimesheetEntriesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimesheetEntriesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimesheetEntriesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimesheetEntriesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TimesheetEntriesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDate

`func (o *TimesheetEntriesPutPreviousValues) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TimesheetEntriesPutPreviousValues) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TimesheetEntriesPutPreviousValues) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *TimesheetEntriesPutPreviousValues) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetHours

`func (o *TimesheetEntriesPutPreviousValues) GetHours() float32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *TimesheetEntriesPutPreviousValues) GetHoursOk() (*float32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *TimesheetEntriesPutPreviousValues) SetHours(v float32)`

SetHours sets Hours field to given value.

### HasHours

`func (o *TimesheetEntriesPutPreviousValues) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetUserId

`func (o *TimesheetEntriesPutPreviousValues) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TimesheetEntriesPutPreviousValues) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TimesheetEntriesPutPreviousValues) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *TimesheetEntriesPutPreviousValues) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetTimesheetableType

`func (o *TimesheetEntriesPutPreviousValues) GetTimesheetableType() string`

GetTimesheetableType returns the TimesheetableType field if non-nil, zero value otherwise.

### GetTimesheetableTypeOk

`func (o *TimesheetEntriesPutPreviousValues) GetTimesheetableTypeOk() (*string, bool)`

GetTimesheetableTypeOk returns a tuple with the TimesheetableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesheetableType

`func (o *TimesheetEntriesPutPreviousValues) SetTimesheetableType(v string)`

SetTimesheetableType sets TimesheetableType field to given value.

### HasTimesheetableType

`func (o *TimesheetEntriesPutPreviousValues) HasTimesheetableType() bool`

HasTimesheetableType returns a boolean if a field has been set.

### GetTimesheetableId

`func (o *TimesheetEntriesPutPreviousValues) GetTimesheetableId() int32`

GetTimesheetableId returns the TimesheetableId field if non-nil, zero value otherwise.

### GetTimesheetableIdOk

`func (o *TimesheetEntriesPutPreviousValues) GetTimesheetableIdOk() (*int32, bool)`

GetTimesheetableIdOk returns a tuple with the TimesheetableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesheetableId

`func (o *TimesheetEntriesPutPreviousValues) SetTimesheetableId(v int32)`

SetTimesheetableId sets TimesheetableId field to given value.

### HasTimesheetableId

`func (o *TimesheetEntriesPutPreviousValues) HasTimesheetableId() bool`

HasTimesheetableId returns a boolean if a field has been set.

### GetTimesheetableMeta

`func (o *TimesheetEntriesPutPreviousValues) GetTimesheetableMeta() string`

GetTimesheetableMeta returns the TimesheetableMeta field if non-nil, zero value otherwise.

### GetTimesheetableMetaOk

`func (o *TimesheetEntriesPutPreviousValues) GetTimesheetableMetaOk() (*string, bool)`

GetTimesheetableMetaOk returns a tuple with the TimesheetableMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesheetableMeta

`func (o *TimesheetEntriesPutPreviousValues) SetTimesheetableMeta(v string)`

SetTimesheetableMeta sets TimesheetableMeta field to given value.

### HasTimesheetableMeta

`func (o *TimesheetEntriesPutPreviousValues) HasTimesheetableMeta() bool`

HasTimesheetableMeta returns a boolean if a field has been set.

### GetTimesheetProjectId

`func (o *TimesheetEntriesPutPreviousValues) GetTimesheetProjectId() int32`

GetTimesheetProjectId returns the TimesheetProjectId field if non-nil, zero value otherwise.

### GetTimesheetProjectIdOk

`func (o *TimesheetEntriesPutPreviousValues) GetTimesheetProjectIdOk() (*int32, bool)`

GetTimesheetProjectIdOk returns a tuple with the TimesheetProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesheetProjectId

`func (o *TimesheetEntriesPutPreviousValues) SetTimesheetProjectId(v int32)`

SetTimesheetProjectId sets TimesheetProjectId field to given value.

### HasTimesheetProjectId

`func (o *TimesheetEntriesPutPreviousValues) HasTimesheetProjectId() bool`

HasTimesheetProjectId returns a boolean if a field has been set.

### GetComment

`func (o *TimesheetEntriesPutPreviousValues) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TimesheetEntriesPutPreviousValues) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TimesheetEntriesPutPreviousValues) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TimesheetEntriesPutPreviousValues) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDeletedAt

`func (o *TimesheetEntriesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *TimesheetEntriesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *TimesheetEntriesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *TimesheetEntriesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TimesheetEntriesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TimesheetEntriesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TimesheetEntriesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TimesheetEntriesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TimesheetEntriesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TimesheetEntriesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TimesheetEntriesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TimesheetEntriesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


