# Bookings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**ResourceId** | **int32** | Note: This is a Foreign Key to &#x60;resources.id&#x60;.&lt;fk table&#x3D;&#39;resources&#39; column&#x3D;&#39;id&#39;/&gt; | 
**ProjectPlanId** | **int32** | Note: This is a Foreign Key to &#x60;project_plans.id&#x60;.&lt;fk table&#x3D;&#39;project_plans&#39; column&#x3D;&#39;id&#39;/&gt; | 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**BookedHours** | **int32** |  | 
**ProjectPlanPhaseId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;project_plan_phases.id&#x60;.&lt;fk table&#x3D;&#39;project_plan_phases&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 

## Methods

### NewBookings

`func NewBookings(resourceId int32, projectPlanId int32, bookedHours int32, ) *Bookings`

NewBookings instantiates a new Bookings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingsWithDefaults

`func NewBookingsWithDefaults() *Bookings`

NewBookingsWithDefaults instantiates a new Bookings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Bookings) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bookings) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bookings) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Bookings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceId

`func (o *Bookings) GetResourceId() int32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Bookings) GetResourceIdOk() (*int32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Bookings) SetResourceId(v int32)`

SetResourceId sets ResourceId field to given value.


### GetProjectPlanId

`func (o *Bookings) GetProjectPlanId() int32`

GetProjectPlanId returns the ProjectPlanId field if non-nil, zero value otherwise.

### GetProjectPlanIdOk

`func (o *Bookings) GetProjectPlanIdOk() (*int32, bool)`

GetProjectPlanIdOk returns a tuple with the ProjectPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPlanId

`func (o *Bookings) SetProjectPlanId(v int32)`

SetProjectPlanId sets ProjectPlanId field to given value.


### GetStartDate

`func (o *Bookings) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Bookings) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Bookings) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Bookings) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *Bookings) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Bookings) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Bookings) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Bookings) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Bookings) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Bookings) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Bookings) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Bookings) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Bookings) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Bookings) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Bookings) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Bookings) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Bookings) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Bookings) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Bookings) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Bookings) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetBookedHours

`func (o *Bookings) GetBookedHours() int32`

GetBookedHours returns the BookedHours field if non-nil, zero value otherwise.

### GetBookedHoursOk

`func (o *Bookings) GetBookedHoursOk() (*int32, bool)`

GetBookedHoursOk returns a tuple with the BookedHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookedHours

`func (o *Bookings) SetBookedHours(v int32)`

SetBookedHours sets BookedHours field to given value.


### GetProjectPlanPhaseId

`func (o *Bookings) GetProjectPlanPhaseId() int32`

GetProjectPlanPhaseId returns the ProjectPlanPhaseId field if non-nil, zero value otherwise.

### GetProjectPlanPhaseIdOk

`func (o *Bookings) GetProjectPlanPhaseIdOk() (*int32, bool)`

GetProjectPlanPhaseIdOk returns a tuple with the ProjectPlanPhaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPlanPhaseId

`func (o *Bookings) SetProjectPlanPhaseId(v int32)`

SetProjectPlanPhaseId sets ProjectPlanPhaseId field to given value.

### HasProjectPlanPhaseId

`func (o *Bookings) HasProjectPlanPhaseId() bool`

HasProjectPlanPhaseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


