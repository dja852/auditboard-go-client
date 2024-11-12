# CustomFieldsPutCustomField

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

### NewCustomFieldsPutCustomField

`func NewCustomFieldsPutCustomField(type_ string, name string, attrKey string, linkifyKey string, ) *CustomFieldsPutCustomField`

NewCustomFieldsPutCustomField instantiates a new CustomFieldsPutCustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldsPutCustomFieldWithDefaults

`func NewCustomFieldsPutCustomFieldWithDefaults() *CustomFieldsPutCustomField`

NewCustomFieldsPutCustomFieldWithDefaults instantiates a new CustomFieldsPutCustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomFieldsPutCustomField) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomFieldsPutCustomField) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomFieldsPutCustomField) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CustomFieldsPutCustomField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CustomFieldsPutCustomField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomFieldsPutCustomField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomFieldsPutCustomField) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CustomFieldsPutCustomField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomFieldsPutCustomField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomFieldsPutCustomField) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *CustomFieldsPutCustomField) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CustomFieldsPutCustomField) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CustomFieldsPutCustomField) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CustomFieldsPutCustomField) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAttrKey

`func (o *CustomFieldsPutCustomField) GetAttrKey() string`

GetAttrKey returns the AttrKey field if non-nil, zero value otherwise.

### GetAttrKeyOk

`func (o *CustomFieldsPutCustomField) GetAttrKeyOk() (*string, bool)`

GetAttrKeyOk returns a tuple with the AttrKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrKey

`func (o *CustomFieldsPutCustomField) SetAttrKey(v string)`

SetAttrKey sets AttrKey field to given value.


### GetCreatedAt

`func (o *CustomFieldsPutCustomField) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomFieldsPutCustomField) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomFieldsPutCustomField) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomFieldsPutCustomField) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CustomFieldsPutCustomField) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomFieldsPutCustomField) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomFieldsPutCustomField) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CustomFieldsPutCustomField) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *CustomFieldsPutCustomField) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CustomFieldsPutCustomField) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CustomFieldsPutCustomField) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *CustomFieldsPutCustomField) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetLinkifyKey

`func (o *CustomFieldsPutCustomField) GetLinkifyKey() string`

GetLinkifyKey returns the LinkifyKey field if non-nil, zero value otherwise.

### GetLinkifyKeyOk

`func (o *CustomFieldsPutCustomField) GetLinkifyKeyOk() (*string, bool)`

GetLinkifyKeyOk returns a tuple with the LinkifyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkifyKey

`func (o *CustomFieldsPutCustomField) SetLinkifyKey(v string)`

SetLinkifyKey sets LinkifyKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


