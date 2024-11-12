# AssessmentPeriodsPutAssessmentPeriod

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

### NewAssessmentPeriodsPutAssessmentPeriod

`func NewAssessmentPeriodsPutAssessmentPeriod(name string, sortOrder int32, ) *AssessmentPeriodsPutAssessmentPeriod`

NewAssessmentPeriodsPutAssessmentPeriod instantiates a new AssessmentPeriodsPutAssessmentPeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssessmentPeriodsPutAssessmentPeriodWithDefaults

`func NewAssessmentPeriodsPutAssessmentPeriodWithDefaults() *AssessmentPeriodsPutAssessmentPeriod`

NewAssessmentPeriodsPutAssessmentPeriodWithDefaults instantiates a new AssessmentPeriodsPutAssessmentPeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssessmentPeriodsPutAssessmentPeriod) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssessmentPeriodsPutAssessmentPeriod) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssessmentPeriodsPutAssessmentPeriod) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AssessmentPeriodsPutAssessmentPeriod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AssessmentPeriodsPutAssessmentPeriod) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AssessmentPeriodsPutAssessmentPeriod) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AssessmentPeriodsPutAssessmentPeriod) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AssessmentPeriodsPutAssessmentPeriod) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AssessmentPeriodsPutAssessmentPeriod) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AssessmentPeriodsPutAssessmentPeriod) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AssessmentPeriodsPutAssessmentPeriod) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AssessmentPeriodsPutAssessmentPeriod) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AssessmentPeriodsPutAssessmentPeriod) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AssessmentPeriodsPutAssessmentPeriod) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AssessmentPeriodsPutAssessmentPeriod) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AssessmentPeriodsPutAssessmentPeriod) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *AssessmentPeriodsPutAssessmentPeriod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssessmentPeriodsPutAssessmentPeriod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssessmentPeriodsPutAssessmentPeriod) SetName(v string)`

SetName sets Name field to given value.


### GetSortOrder

`func (o *AssessmentPeriodsPutAssessmentPeriod) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AssessmentPeriodsPutAssessmentPeriod) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AssessmentPeriodsPutAssessmentPeriod) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


