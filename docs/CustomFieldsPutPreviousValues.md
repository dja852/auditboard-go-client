# CustomFieldsPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**AttrKey** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**LinkifyKey** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomFieldsPutPreviousValues

`func NewCustomFieldsPutPreviousValues() *CustomFieldsPutPreviousValues`

NewCustomFieldsPutPreviousValues instantiates a new CustomFieldsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldsPutPreviousValuesWithDefaults

`func NewCustomFieldsPutPreviousValuesWithDefaults() *CustomFieldsPutPreviousValues`

NewCustomFieldsPutPreviousValuesWithDefaults instantiates a new CustomFieldsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomFieldsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomFieldsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomFieldsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CustomFieldsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CustomFieldsPutPreviousValues) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomFieldsPutPreviousValues) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomFieldsPutPreviousValues) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomFieldsPutPreviousValues) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *CustomFieldsPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomFieldsPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomFieldsPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomFieldsPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *CustomFieldsPutPreviousValues) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CustomFieldsPutPreviousValues) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CustomFieldsPutPreviousValues) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CustomFieldsPutPreviousValues) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAttrKey

`func (o *CustomFieldsPutPreviousValues) GetAttrKey() string`

GetAttrKey returns the AttrKey field if non-nil, zero value otherwise.

### GetAttrKeyOk

`func (o *CustomFieldsPutPreviousValues) GetAttrKeyOk() (*string, bool)`

GetAttrKeyOk returns a tuple with the AttrKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrKey

`func (o *CustomFieldsPutPreviousValues) SetAttrKey(v string)`

SetAttrKey sets AttrKey field to given value.

### HasAttrKey

`func (o *CustomFieldsPutPreviousValues) HasAttrKey() bool`

HasAttrKey returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CustomFieldsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomFieldsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomFieldsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomFieldsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CustomFieldsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomFieldsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomFieldsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CustomFieldsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *CustomFieldsPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CustomFieldsPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CustomFieldsPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *CustomFieldsPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetLinkifyKey

`func (o *CustomFieldsPutPreviousValues) GetLinkifyKey() string`

GetLinkifyKey returns the LinkifyKey field if non-nil, zero value otherwise.

### GetLinkifyKeyOk

`func (o *CustomFieldsPutPreviousValues) GetLinkifyKeyOk() (*string, bool)`

GetLinkifyKeyOk returns a tuple with the LinkifyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkifyKey

`func (o *CustomFieldsPutPreviousValues) SetLinkifyKey(v string)`

SetLinkifyKey sets LinkifyKey field to given value.

### HasLinkifyKey

`func (o *CustomFieldsPutPreviousValues) HasLinkifyKey() bool`

HasLinkifyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


