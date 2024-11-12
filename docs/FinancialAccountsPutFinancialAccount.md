# FinancialAccountsPutFinancialAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**SortOrder** | **int32** |  | [default to 0]
**Balance** | Pointer to **float32** |  | [optional] 
**FieldData** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewFinancialAccountsPutFinancialAccount

`func NewFinancialAccountsPutFinancialAccount(name string, sortOrder int32, ) *FinancialAccountsPutFinancialAccount`

NewFinancialAccountsPutFinancialAccount instantiates a new FinancialAccountsPutFinancialAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialAccountsPutFinancialAccountWithDefaults

`func NewFinancialAccountsPutFinancialAccountWithDefaults() *FinancialAccountsPutFinancialAccount`

NewFinancialAccountsPutFinancialAccountWithDefaults instantiates a new FinancialAccountsPutFinancialAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FinancialAccountsPutFinancialAccount) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FinancialAccountsPutFinancialAccount) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FinancialAccountsPutFinancialAccount) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FinancialAccountsPutFinancialAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FinancialAccountsPutFinancialAccount) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FinancialAccountsPutFinancialAccount) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FinancialAccountsPutFinancialAccount) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FinancialAccountsPutFinancialAccount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FinancialAccountsPutFinancialAccount) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FinancialAccountsPutFinancialAccount) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FinancialAccountsPutFinancialAccount) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FinancialAccountsPutFinancialAccount) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *FinancialAccountsPutFinancialAccount) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *FinancialAccountsPutFinancialAccount) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *FinancialAccountsPutFinancialAccount) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *FinancialAccountsPutFinancialAccount) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *FinancialAccountsPutFinancialAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FinancialAccountsPutFinancialAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FinancialAccountsPutFinancialAccount) SetName(v string)`

SetName sets Name field to given value.


### GetSortOrder

`func (o *FinancialAccountsPutFinancialAccount) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *FinancialAccountsPutFinancialAccount) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *FinancialAccountsPutFinancialAccount) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.


### GetBalance

`func (o *FinancialAccountsPutFinancialAccount) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *FinancialAccountsPutFinancialAccount) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *FinancialAccountsPutFinancialAccount) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *FinancialAccountsPutFinancialAccount) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetFieldData

`func (o *FinancialAccountsPutFinancialAccount) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *FinancialAccountsPutFinancialAccount) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *FinancialAccountsPutFinancialAccount) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *FinancialAccountsPutFinancialAccount) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *FinancialAccountsPutFinancialAccount) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *FinancialAccountsPutFinancialAccount) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


