# AllowedUsersPutAllowedUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**UserableId** | Pointer to **int32** |  | [optional] 
**UserableType** | Pointer to **string** |  | [optional] 
**Permission** | Pointer to **string** |  | [optional] 
**UserId** | **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewAllowedUsersPutAllowedUser

`func NewAllowedUsersPutAllowedUser(userId int32, ) *AllowedUsersPutAllowedUser`

NewAllowedUsersPutAllowedUser instantiates a new AllowedUsersPutAllowedUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedUsersPutAllowedUserWithDefaults

`func NewAllowedUsersPutAllowedUserWithDefaults() *AllowedUsersPutAllowedUser`

NewAllowedUsersPutAllowedUserWithDefaults instantiates a new AllowedUsersPutAllowedUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AllowedUsersPutAllowedUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AllowedUsersPutAllowedUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AllowedUsersPutAllowedUser) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AllowedUsersPutAllowedUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserableId

`func (o *AllowedUsersPutAllowedUser) GetUserableId() int32`

GetUserableId returns the UserableId field if non-nil, zero value otherwise.

### GetUserableIdOk

`func (o *AllowedUsersPutAllowedUser) GetUserableIdOk() (*int32, bool)`

GetUserableIdOk returns a tuple with the UserableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserableId

`func (o *AllowedUsersPutAllowedUser) SetUserableId(v int32)`

SetUserableId sets UserableId field to given value.

### HasUserableId

`func (o *AllowedUsersPutAllowedUser) HasUserableId() bool`

HasUserableId returns a boolean if a field has been set.

### GetUserableType

`func (o *AllowedUsersPutAllowedUser) GetUserableType() string`

GetUserableType returns the UserableType field if non-nil, zero value otherwise.

### GetUserableTypeOk

`func (o *AllowedUsersPutAllowedUser) GetUserableTypeOk() (*string, bool)`

GetUserableTypeOk returns a tuple with the UserableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserableType

`func (o *AllowedUsersPutAllowedUser) SetUserableType(v string)`

SetUserableType sets UserableType field to given value.

### HasUserableType

`func (o *AllowedUsersPutAllowedUser) HasUserableType() bool`

HasUserableType returns a boolean if a field has been set.

### GetPermission

`func (o *AllowedUsersPutAllowedUser) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *AllowedUsersPutAllowedUser) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *AllowedUsersPutAllowedUser) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *AllowedUsersPutAllowedUser) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetUserId

`func (o *AllowedUsersPutAllowedUser) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AllowedUsersPutAllowedUser) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AllowedUsersPutAllowedUser) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetCreatedAt

`func (o *AllowedUsersPutAllowedUser) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AllowedUsersPutAllowedUser) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AllowedUsersPutAllowedUser) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AllowedUsersPutAllowedUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AllowedUsersPutAllowedUser) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AllowedUsersPutAllowedUser) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AllowedUsersPutAllowedUser) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AllowedUsersPutAllowedUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


