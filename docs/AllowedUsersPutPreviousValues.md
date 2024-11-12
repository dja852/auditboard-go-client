# AllowedUsersPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**UserableId** | Pointer to **int32** |  | [optional] 
**UserableType** | Pointer to **string** |  | [optional] 
**Permission** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewAllowedUsersPutPreviousValues

`func NewAllowedUsersPutPreviousValues() *AllowedUsersPutPreviousValues`

NewAllowedUsersPutPreviousValues instantiates a new AllowedUsersPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedUsersPutPreviousValuesWithDefaults

`func NewAllowedUsersPutPreviousValuesWithDefaults() *AllowedUsersPutPreviousValues`

NewAllowedUsersPutPreviousValuesWithDefaults instantiates a new AllowedUsersPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AllowedUsersPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AllowedUsersPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AllowedUsersPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AllowedUsersPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserableId

`func (o *AllowedUsersPutPreviousValues) GetUserableId() int32`

GetUserableId returns the UserableId field if non-nil, zero value otherwise.

### GetUserableIdOk

`func (o *AllowedUsersPutPreviousValues) GetUserableIdOk() (*int32, bool)`

GetUserableIdOk returns a tuple with the UserableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserableId

`func (o *AllowedUsersPutPreviousValues) SetUserableId(v int32)`

SetUserableId sets UserableId field to given value.

### HasUserableId

`func (o *AllowedUsersPutPreviousValues) HasUserableId() bool`

HasUserableId returns a boolean if a field has been set.

### GetUserableType

`func (o *AllowedUsersPutPreviousValues) GetUserableType() string`

GetUserableType returns the UserableType field if non-nil, zero value otherwise.

### GetUserableTypeOk

`func (o *AllowedUsersPutPreviousValues) GetUserableTypeOk() (*string, bool)`

GetUserableTypeOk returns a tuple with the UserableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserableType

`func (o *AllowedUsersPutPreviousValues) SetUserableType(v string)`

SetUserableType sets UserableType field to given value.

### HasUserableType

`func (o *AllowedUsersPutPreviousValues) HasUserableType() bool`

HasUserableType returns a boolean if a field has been set.

### GetPermission

`func (o *AllowedUsersPutPreviousValues) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *AllowedUsersPutPreviousValues) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *AllowedUsersPutPreviousValues) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *AllowedUsersPutPreviousValues) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetUserId

`func (o *AllowedUsersPutPreviousValues) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AllowedUsersPutPreviousValues) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AllowedUsersPutPreviousValues) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AllowedUsersPutPreviousValues) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AllowedUsersPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AllowedUsersPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AllowedUsersPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AllowedUsersPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AllowedUsersPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AllowedUsersPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AllowedUsersPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AllowedUsersPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


