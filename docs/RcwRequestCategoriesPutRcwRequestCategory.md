# RcwRequestCategoriesPutRcwRequestCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Name** | **string** |  | 
**CreatedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRcwRequestCategoriesPutRcwRequestCategory

`func NewRcwRequestCategoriesPutRcwRequestCategory(name string, ) *RcwRequestCategoriesPutRcwRequestCategory`

NewRcwRequestCategoriesPutRcwRequestCategory instantiates a new RcwRequestCategoriesPutRcwRequestCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRcwRequestCategoriesPutRcwRequestCategoryWithDefaults

`func NewRcwRequestCategoriesPutRcwRequestCategoryWithDefaults() *RcwRequestCategoriesPutRcwRequestCategory`

NewRcwRequestCategoriesPutRcwRequestCategoryWithDefaults instantiates a new RcwRequestCategoriesPutRcwRequestCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RcwRequestCategoriesPutRcwRequestCategory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RcwRequestCategoriesPutRcwRequestCategory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RcwRequestCategoriesPutRcwRequestCategory) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RcwRequestCategoriesPutRcwRequestCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RcwRequestCategoriesPutRcwRequestCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RcwRequestCategoriesPutRcwRequestCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RcwRequestCategoriesPutRcwRequestCategory) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedByUserId

`func (o *RcwRequestCategoriesPutRcwRequestCategory) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *RcwRequestCategoriesPutRcwRequestCategory) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *RcwRequestCategoriesPutRcwRequestCategory) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *RcwRequestCategoriesPutRcwRequestCategory) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RcwRequestCategoriesPutRcwRequestCategory) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RcwRequestCategoriesPutRcwRequestCategory) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RcwRequestCategoriesPutRcwRequestCategory) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RcwRequestCategoriesPutRcwRequestCategory) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RcwRequestCategoriesPutRcwRequestCategory) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RcwRequestCategoriesPutRcwRequestCategory) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RcwRequestCategoriesPutRcwRequestCategory) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RcwRequestCategoriesPutRcwRequestCategory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *RcwRequestCategoriesPutRcwRequestCategory) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *RcwRequestCategoriesPutRcwRequestCategory) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *RcwRequestCategoriesPutRcwRequestCategory) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *RcwRequestCategoriesPutRcwRequestCategory) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


