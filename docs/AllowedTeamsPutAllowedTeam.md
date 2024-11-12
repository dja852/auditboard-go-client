# AllowedTeamsPutAllowedTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**TeamableId** | Pointer to **int32** |  | [optional] 
**TeamableType** | Pointer to **string** |  | [optional] 
**Permission** | Pointer to **string** |  | [optional] 
**TeamId** | **int32** | Note: This is a Foreign Key to &#x60;teams.id&#x60;.&lt;fk table&#x3D;&#39;teams&#39; column&#x3D;&#39;id&#39;/&gt; | 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewAllowedTeamsPutAllowedTeam

`func NewAllowedTeamsPutAllowedTeam(teamId int32, ) *AllowedTeamsPutAllowedTeam`

NewAllowedTeamsPutAllowedTeam instantiates a new AllowedTeamsPutAllowedTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedTeamsPutAllowedTeamWithDefaults

`func NewAllowedTeamsPutAllowedTeamWithDefaults() *AllowedTeamsPutAllowedTeam`

NewAllowedTeamsPutAllowedTeamWithDefaults instantiates a new AllowedTeamsPutAllowedTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AllowedTeamsPutAllowedTeam) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AllowedTeamsPutAllowedTeam) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AllowedTeamsPutAllowedTeam) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AllowedTeamsPutAllowedTeam) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTeamableId

`func (o *AllowedTeamsPutAllowedTeam) GetTeamableId() int32`

GetTeamableId returns the TeamableId field if non-nil, zero value otherwise.

### GetTeamableIdOk

`func (o *AllowedTeamsPutAllowedTeam) GetTeamableIdOk() (*int32, bool)`

GetTeamableIdOk returns a tuple with the TeamableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamableId

`func (o *AllowedTeamsPutAllowedTeam) SetTeamableId(v int32)`

SetTeamableId sets TeamableId field to given value.

### HasTeamableId

`func (o *AllowedTeamsPutAllowedTeam) HasTeamableId() bool`

HasTeamableId returns a boolean if a field has been set.

### GetTeamableType

`func (o *AllowedTeamsPutAllowedTeam) GetTeamableType() string`

GetTeamableType returns the TeamableType field if non-nil, zero value otherwise.

### GetTeamableTypeOk

`func (o *AllowedTeamsPutAllowedTeam) GetTeamableTypeOk() (*string, bool)`

GetTeamableTypeOk returns a tuple with the TeamableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamableType

`func (o *AllowedTeamsPutAllowedTeam) SetTeamableType(v string)`

SetTeamableType sets TeamableType field to given value.

### HasTeamableType

`func (o *AllowedTeamsPutAllowedTeam) HasTeamableType() bool`

HasTeamableType returns a boolean if a field has been set.

### GetPermission

`func (o *AllowedTeamsPutAllowedTeam) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *AllowedTeamsPutAllowedTeam) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *AllowedTeamsPutAllowedTeam) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *AllowedTeamsPutAllowedTeam) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetTeamId

`func (o *AllowedTeamsPutAllowedTeam) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *AllowedTeamsPutAllowedTeam) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *AllowedTeamsPutAllowedTeam) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.


### GetCreatedAt

`func (o *AllowedTeamsPutAllowedTeam) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AllowedTeamsPutAllowedTeam) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AllowedTeamsPutAllowedTeam) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AllowedTeamsPutAllowedTeam) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AllowedTeamsPutAllowedTeam) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AllowedTeamsPutAllowedTeam) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AllowedTeamsPutAllowedTeam) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AllowedTeamsPutAllowedTeam) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


