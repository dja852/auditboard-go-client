# CustomFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Type** | **string** |  | 
**Name** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**AttrKey** | **string** |  | 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**LinkifyKey** | **string** |  | 

## Methods

### NewCustomFields

`func NewCustomFields(type_ string, name string, attrKey string, linkifyKey string, ) *CustomFields`

NewCustomFields instantiates a new CustomFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldsWithDefaults

`func NewCustomFieldsWithDefaults() *CustomFields`

NewCustomFieldsWithDefaults instantiates a new CustomFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CustomFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CustomFields) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomFields) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomFields) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CustomFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomFields) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *CustomFields) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CustomFields) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CustomFields) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CustomFields) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAttrKey

`func (o *CustomFields) GetAttrKey() string`

GetAttrKey returns the AttrKey field if non-nil, zero value otherwise.

### GetAttrKeyOk

`func (o *CustomFields) GetAttrKeyOk() (*string, bool)`

GetAttrKeyOk returns a tuple with the AttrKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrKey

`func (o *CustomFields) SetAttrKey(v string)`

SetAttrKey sets AttrKey field to given value.


### GetCreatedAt

`func (o *CustomFields) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomFields) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomFields) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CustomFields) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomFields) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomFields) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CustomFields) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *CustomFields) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CustomFields) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CustomFields) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *CustomFields) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetLinkifyKey

`func (o *CustomFields) GetLinkifyKey() string`

GetLinkifyKey returns the LinkifyKey field if non-nil, zero value otherwise.

### GetLinkifyKeyOk

`func (o *CustomFields) GetLinkifyKeyOk() (*string, bool)`

GetLinkifyKeyOk returns a tuple with the LinkifyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkifyKey

`func (o *CustomFields) SetLinkifyKey(v string)`

SetLinkifyKey sets LinkifyKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


