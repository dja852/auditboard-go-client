# RiskCategoriesPutRiskCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**SortOrder** | **int32** |  | [default to 0]
**FormTemplateId** | Pointer to **int32** |  | [optional] 
**IsDefault** | **bool** |  | [default to false]
**Key** | Pointer to **string** |  | [optional] 

## Methods

### NewRiskCategoriesPutRiskCategory

`func NewRiskCategoriesPutRiskCategory(name string, sortOrder int32, isDefault bool, ) *RiskCategoriesPutRiskCategory`

NewRiskCategoriesPutRiskCategory instantiates a new RiskCategoriesPutRiskCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskCategoriesPutRiskCategoryWithDefaults

`func NewRiskCategoriesPutRiskCategoryWithDefaults() *RiskCategoriesPutRiskCategory`

NewRiskCategoriesPutRiskCategoryWithDefaults instantiates a new RiskCategoriesPutRiskCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RiskCategoriesPutRiskCategory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskCategoriesPutRiskCategory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskCategoriesPutRiskCategory) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RiskCategoriesPutRiskCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RiskCategoriesPutRiskCategory) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RiskCategoriesPutRiskCategory) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RiskCategoriesPutRiskCategory) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RiskCategoriesPutRiskCategory) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RiskCategoriesPutRiskCategory) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RiskCategoriesPutRiskCategory) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RiskCategoriesPutRiskCategory) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RiskCategoriesPutRiskCategory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *RiskCategoriesPutRiskCategory) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *RiskCategoriesPutRiskCategory) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *RiskCategoriesPutRiskCategory) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *RiskCategoriesPutRiskCategory) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *RiskCategoriesPutRiskCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskCategoriesPutRiskCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskCategoriesPutRiskCategory) SetName(v string)`

SetName sets Name field to given value.


### GetSortOrder

`func (o *RiskCategoriesPutRiskCategory) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *RiskCategoriesPutRiskCategory) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *RiskCategoriesPutRiskCategory) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.


### GetFormTemplateId

`func (o *RiskCategoriesPutRiskCategory) GetFormTemplateId() int32`

GetFormTemplateId returns the FormTemplateId field if non-nil, zero value otherwise.

### GetFormTemplateIdOk

`func (o *RiskCategoriesPutRiskCategory) GetFormTemplateIdOk() (*int32, bool)`

GetFormTemplateIdOk returns a tuple with the FormTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormTemplateId

`func (o *RiskCategoriesPutRiskCategory) SetFormTemplateId(v int32)`

SetFormTemplateId sets FormTemplateId field to given value.

### HasFormTemplateId

`func (o *RiskCategoriesPutRiskCategory) HasFormTemplateId() bool`

HasFormTemplateId returns a boolean if a field has been set.

### GetIsDefault

`func (o *RiskCategoriesPutRiskCategory) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *RiskCategoriesPutRiskCategory) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *RiskCategoriesPutRiskCategory) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetKey

`func (o *RiskCategoriesPutRiskCategory) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RiskCategoriesPutRiskCategory) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RiskCategoriesPutRiskCategory) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RiskCategoriesPutRiskCategory) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


