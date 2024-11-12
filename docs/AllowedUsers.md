# AllowedUsers

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

### NewAllowedUsers

`func NewAllowedUsers(userId int32, ) *AllowedUsers`

NewAllowedUsers instantiates a new AllowedUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedUsersWithDefaults

`func NewAllowedUsersWithDefaults() *AllowedUsers`

NewAllowedUsersWithDefaults instantiates a new AllowedUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AllowedUsers) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AllowedUsers) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AllowedUsers) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AllowedUsers) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserableId

`func (o *AllowedUsers) GetUserableId() int32`

GetUserableId returns the UserableId field if non-nil, zero value otherwise.

### GetUserableIdOk

`func (o *AllowedUsers) GetUserableIdOk() (*int32, bool)`

GetUserableIdOk returns a tuple with the UserableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserableId

`func (o *AllowedUsers) SetUserableId(v int32)`

SetUserableId sets UserableId field to given value.

### HasUserableId

`func (o *AllowedUsers) HasUserableId() bool`

HasUserableId returns a boolean if a field has been set.

### GetUserableType

`func (o *AllowedUsers) GetUserableType() string`

GetUserableType returns the UserableType field if non-nil, zero value otherwise.

### GetUserableTypeOk

`func (o *AllowedUsers) GetUserableTypeOk() (*string, bool)`

GetUserableTypeOk returns a tuple with the UserableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserableType

`func (o *AllowedUsers) SetUserableType(v string)`

SetUserableType sets UserableType field to given value.

### HasUserableType

`func (o *AllowedUsers) HasUserableType() bool`

HasUserableType returns a boolean if a field has been set.

### GetPermission

`func (o *AllowedUsers) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *AllowedUsers) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *AllowedUsers) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *AllowedUsers) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetUserId

`func (o *AllowedUsers) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AllowedUsers) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AllowedUsers) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetCreatedAt

`func (o *AllowedUsers) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AllowedUsers) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AllowedUsers) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AllowedUsers) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AllowedUsers) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AllowedUsers) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AllowedUsers) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AllowedUsers) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


