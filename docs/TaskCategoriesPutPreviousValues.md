# TaskCategoriesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**SortOrder** | Pointer to **int32** |  | [optional] [default to 0]
**MetaConfig** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewTaskCategoriesPutPreviousValues

`func NewTaskCategoriesPutPreviousValues() *TaskCategoriesPutPreviousValues`

NewTaskCategoriesPutPreviousValues instantiates a new TaskCategoriesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskCategoriesPutPreviousValuesWithDefaults

`func NewTaskCategoriesPutPreviousValuesWithDefaults() *TaskCategoriesPutPreviousValues`

NewTaskCategoriesPutPreviousValuesWithDefaults instantiates a new TaskCategoriesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskCategoriesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskCategoriesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskCategoriesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TaskCategoriesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskCategoriesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskCategoriesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskCategoriesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaskCategoriesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TaskCategoriesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskCategoriesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskCategoriesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TaskCategoriesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *TaskCategoriesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *TaskCategoriesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *TaskCategoriesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *TaskCategoriesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *TaskCategoriesPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskCategoriesPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskCategoriesPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskCategoriesPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSortOrder

`func (o *TaskCategoriesPutPreviousValues) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *TaskCategoriesPutPreviousValues) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *TaskCategoriesPutPreviousValues) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *TaskCategoriesPutPreviousValues) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetMetaConfig

`func (o *TaskCategoriesPutPreviousValues) GetMetaConfig() string`

GetMetaConfig returns the MetaConfig field if non-nil, zero value otherwise.

### GetMetaConfigOk

`func (o *TaskCategoriesPutPreviousValues) GetMetaConfigOk() (*string, bool)`

GetMetaConfigOk returns a tuple with the MetaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaConfig

`func (o *TaskCategoriesPutPreviousValues) SetMetaConfig(v string)`

SetMetaConfig sets MetaConfig field to given value.

### HasMetaConfig

`func (o *TaskCategoriesPutPreviousValues) HasMetaConfig() bool`

HasMetaConfig returns a boolean if a field has been set.

### GetColor

`func (o *TaskCategoriesPutPreviousValues) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *TaskCategoriesPutPreviousValues) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *TaskCategoriesPutPreviousValues) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *TaskCategoriesPutPreviousValues) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetShortName

`func (o *TaskCategoriesPutPreviousValues) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *TaskCategoriesPutPreviousValues) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *TaskCategoriesPutPreviousValues) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *TaskCategoriesPutPreviousValues) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetScopes

`func (o *TaskCategoriesPutPreviousValues) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TaskCategoriesPutPreviousValues) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TaskCategoriesPutPreviousValues) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TaskCategoriesPutPreviousValues) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *TaskCategoriesPutPreviousValues) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *TaskCategoriesPutPreviousValues) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


