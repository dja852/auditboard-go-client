# TeamPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**TeamId** | **int32** | Note: This is a Foreign Key to &#x60;teams.id&#x60;.&lt;fk table&#x3D;&#39;teams&#39; column&#x3D;&#39;id&#39;/&gt; | 
**GlobalPermissionId** | **int32** | Note: This is a Foreign Key to &#x60;global_permissions.id&#x60;.&lt;fk table&#x3D;&#39;global_permissions&#39; column&#x3D;&#39;id&#39;/&gt; | 
**Section** | **string** |  | 
**Key** | **string** |  | 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewTeamPermissions

`func NewTeamPermissions(teamId int32, globalPermissionId int32, section string, key string, ) *TeamPermissions`

NewTeamPermissions instantiates a new TeamPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamPermissionsWithDefaults

`func NewTeamPermissionsWithDefaults() *TeamPermissions`

NewTeamPermissionsWithDefaults instantiates a new TeamPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamPermissions) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamPermissions) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamPermissions) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TeamPermissions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTeamId

`func (o *TeamPermissions) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *TeamPermissions) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *TeamPermissions) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.


### GetGlobalPermissionId

`func (o *TeamPermissions) GetGlobalPermissionId() int32`

GetGlobalPermissionId returns the GlobalPermissionId field if non-nil, zero value otherwise.

### GetGlobalPermissionIdOk

`func (o *TeamPermissions) GetGlobalPermissionIdOk() (*int32, bool)`

GetGlobalPermissionIdOk returns a tuple with the GlobalPermissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPermissionId

`func (o *TeamPermissions) SetGlobalPermissionId(v int32)`

SetGlobalPermissionId sets GlobalPermissionId field to given value.


### GetSection

`func (o *TeamPermissions) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *TeamPermissions) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *TeamPermissions) SetSection(v string)`

SetSection sets Section field to given value.


### GetKey

`func (o *TeamPermissions) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TeamPermissions) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TeamPermissions) SetKey(v string)`

SetKey sets Key field to given value.


### GetCreatedAt

`func (o *TeamPermissions) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TeamPermissions) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TeamPermissions) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TeamPermissions) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TeamPermissions) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TeamPermissions) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TeamPermissions) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TeamPermissions) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


