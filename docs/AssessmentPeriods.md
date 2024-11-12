# AssessmentPeriods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**SortOrder** | **int32** |  | [default to 0]

## Methods

### NewAssessmentPeriods

`func NewAssessmentPeriods(name string, sortOrder int32, ) *AssessmentPeriods`

NewAssessmentPeriods instantiates a new AssessmentPeriods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssessmentPeriodsWithDefaults

`func NewAssessmentPeriodsWithDefaults() *AssessmentPeriods`

NewAssessmentPeriodsWithDefaults instantiates a new AssessmentPeriods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssessmentPeriods) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssessmentPeriods) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssessmentPeriods) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AssessmentPeriods) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AssessmentPeriods) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AssessmentPeriods) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AssessmentPeriods) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AssessmentPeriods) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AssessmentPeriods) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AssessmentPeriods) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AssessmentPeriods) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AssessmentPeriods) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AssessmentPeriods) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AssessmentPeriods) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AssessmentPeriods) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AssessmentPeriods) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *AssessmentPeriods) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssessmentPeriods) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssessmentPeriods) SetName(v string)`

SetName sets Name field to given value.


### GetSortOrder

`func (o *AssessmentPeriods) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AssessmentPeriods) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AssessmentPeriods) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

