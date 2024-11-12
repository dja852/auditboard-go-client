# AllowedTeams

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

### NewAllowedTeams

`func NewAllowedTeams(teamId int32, ) *AllowedTeams`

NewAllowedTeams instantiates a new AllowedTeams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedTeamsWithDefaults

`func NewAllowedTeamsWithDefaults() *AllowedTeams`

NewAllowedTeamsWithDefaults instantiates a new AllowedTeams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AllowedTeams) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AllowedTeams) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AllowedTeams) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AllowedTeams) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTeamableId

`func (o *AllowedTeams) GetTeamableId() int32`

GetTeamableId returns the TeamableId field if non-nil, zero value otherwise.

### GetTeamableIdOk

`func (o *AllowedTeams) GetTeamableIdOk() (*int32, bool)`

GetTeamableIdOk returns a tuple with the TeamableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamableId

`func (o *AllowedTeams) SetTeamableId(v int32)`

SetTeamableId sets TeamableId field to given value.

### HasTeamableId

`func (o *AllowedTeams) HasTeamableId() bool`

HasTeamableId returns a boolean if a field has been set.

### GetTeamableType

`func (o *AllowedTeams) GetTeamableType() string`

GetTeamableType returns the TeamableType field if non-nil, zero value otherwise.

### GetTeamableTypeOk

`func (o *AllowedTeams) GetTeamableTypeOk() (*string, bool)`

GetTeamableTypeOk returns a tuple with the TeamableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamableType

`func (o *AllowedTeams) SetTeamableType(v string)`

SetTeamableType sets TeamableType field to given value.

### HasTeamableType

`func (o *AllowedTeams) HasTeamableType() bool`

HasTeamableType returns a boolean if a field has been set.

### GetPermission

`func (o *AllowedTeams) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *AllowedTeams) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *AllowedTeams) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *AllowedTeams) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetTeamId

`func (o *AllowedTeams) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *AllowedTeams) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *AllowedTeams) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.


### GetCreatedAt

`func (o *AllowedTeams) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AllowedTeams) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AllowedTeams) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AllowedTeams) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AllowedTeams) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AllowedTeams) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AllowedTeams) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AllowedTeams) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


