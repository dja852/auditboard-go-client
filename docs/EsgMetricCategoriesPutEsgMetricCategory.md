# EsgMetricCategoriesPutEsgMetricCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**SortOrder** | **int32** |  | [default to 0]
**Scopes** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewEsgMetricCategoriesPutEsgMetricCategory

`func NewEsgMetricCategoriesPutEsgMetricCategory(name string, sortOrder int32, ) *EsgMetricCategoriesPutEsgMetricCategory`

NewEsgMetricCategoriesPutEsgMetricCategory instantiates a new EsgMetricCategoriesPutEsgMetricCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsgMetricCategoriesPutEsgMetricCategoryWithDefaults

`func NewEsgMetricCategoriesPutEsgMetricCategoryWithDefaults() *EsgMetricCategoriesPutEsgMetricCategory`

NewEsgMetricCategoriesPutEsgMetricCategoryWithDefaults instantiates a new EsgMetricCategoriesPutEsgMetricCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EsgMetricCategoriesPutEsgMetricCategory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsgMetricCategoriesPutEsgMetricCategory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsgMetricCategoriesPutEsgMetricCategory) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EsgMetricCategoriesPutEsgMetricCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EsgMetricCategoriesPutEsgMetricCategory) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EsgMetricCategoriesPutEsgMetricCategory) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EsgMetricCategoriesPutEsgMetricCategory) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EsgMetricCategoriesPutEsgMetricCategory) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EsgMetricCategoriesPutEsgMetricCategory) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EsgMetricCategoriesPutEsgMetricCategory) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EsgMetricCategoriesPutEsgMetricCategory) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EsgMetricCategoriesPutEsgMetricCategory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EsgMetricCategoriesPutEsgMetricCategory) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EsgMetricCategoriesPutEsgMetricCategory) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EsgMetricCategoriesPutEsgMetricCategory) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EsgMetricCategoriesPutEsgMetricCategory) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *EsgMetricCategoriesPutEsgMetricCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EsgMetricCategoriesPutEsgMetricCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EsgMetricCategoriesPutEsgMetricCategory) SetName(v string)`

SetName sets Name field to given value.


### GetSortOrder

`func (o *EsgMetricCategoriesPutEsgMetricCategory) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *EsgMetricCategoriesPutEsgMetricCategory) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *EsgMetricCategoriesPutEsgMetricCategory) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.


### GetScopes

`func (o *EsgMetricCategoriesPutEsgMetricCategory) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *EsgMetricCategoriesPutEsgMetricCategory) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *EsgMetricCategoriesPutEsgMetricCategory) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *EsgMetricCategoriesPutEsgMetricCategory) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *EsgMetricCategoriesPutEsgMetricCategory) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *EsgMetricCategoriesPutEsgMetricCategory) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


