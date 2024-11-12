# TaskCategories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**SortOrder** | **int32** |  | [default to 0]
**MetaConfig** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewTaskCategories

`func NewTaskCategories(name string, sortOrder int32, ) *TaskCategories`

NewTaskCategories instantiates a new TaskCategories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskCategoriesWithDefaults

`func NewTaskCategoriesWithDefaults() *TaskCategories`

NewTaskCategoriesWithDefaults instantiates a new TaskCategories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskCategories) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskCategories) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskCategories) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TaskCategories) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskCategories) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskCategories) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskCategories) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaskCategories) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TaskCategories) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskCategories) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskCategories) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TaskCategories) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *TaskCategories) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *TaskCategories) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *TaskCategories) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *TaskCategories) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *TaskCategories) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskCategories) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskCategories) SetName(v string)`

SetName sets Name field to given value.


### GetSortOrder

`func (o *TaskCategories) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *TaskCategories) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *TaskCategories) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.


### GetMetaConfig

`func (o *TaskCategories) GetMetaConfig() string`

GetMetaConfig returns the MetaConfig field if non-nil, zero value otherwise.

### GetMetaConfigOk

`func (o *TaskCategories) GetMetaConfigOk() (*string, bool)`

GetMetaConfigOk returns a tuple with the MetaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaConfig

`func (o *TaskCategories) SetMetaConfig(v string)`

SetMetaConfig sets MetaConfig field to given value.

### HasMetaConfig

`func (o *TaskCategories) HasMetaConfig() bool`

HasMetaConfig returns a boolean if a field has been set.

### GetColor

`func (o *TaskCategories) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *TaskCategories) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *TaskCategories) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *TaskCategories) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetShortName

`func (o *TaskCategories) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *TaskCategories) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *TaskCategories) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *TaskCategories) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetScopes

`func (o *TaskCategories) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TaskCategories) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TaskCategories) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TaskCategories) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *TaskCategories) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *TaskCategories) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


