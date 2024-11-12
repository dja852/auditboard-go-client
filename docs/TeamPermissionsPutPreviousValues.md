# TeamPermissionsPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**TeamId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;teams.id&#x60;.&lt;fk table&#x3D;&#39;teams&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**GlobalPermissionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;global_permissions.id&#x60;.&lt;fk table&#x3D;&#39;global_permissions&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Section** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewTeamPermissionsPutPreviousValues

`func NewTeamPermissionsPutPreviousValues() *TeamPermissionsPutPreviousValues`

NewTeamPermissionsPutPreviousValues instantiates a new TeamPermissionsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamPermissionsPutPreviousValuesWithDefaults

`func NewTeamPermissionsPutPreviousValuesWithDefaults() *TeamPermissionsPutPreviousValues`

NewTeamPermissionsPutPreviousValuesWithDefaults instantiates a new TeamPermissionsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamPermissionsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamPermissionsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamPermissionsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TeamPermissionsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTeamId

`func (o *TeamPermissionsPutPreviousValues) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *TeamPermissionsPutPreviousValues) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *TeamPermissionsPutPreviousValues) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *TeamPermissionsPutPreviousValues) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetGlobalPermissionId

`func (o *TeamPermissionsPutPreviousValues) GetGlobalPermissionId() int32`

GetGlobalPermissionId returns the GlobalPermissionId field if non-nil, zero value otherwise.

### GetGlobalPermissionIdOk

`func (o *TeamPermissionsPutPreviousValues) GetGlobalPermissionIdOk() (*int32, bool)`

GetGlobalPermissionIdOk returns a tuple with the GlobalPermissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPermissionId

`func (o *TeamPermissionsPutPreviousValues) SetGlobalPermissionId(v int32)`

SetGlobalPermissionId sets GlobalPermissionId field to given value.

### HasGlobalPermissionId

`func (o *TeamPermissionsPutPreviousValues) HasGlobalPermissionId() bool`

HasGlobalPermissionId returns a boolean if a field has been set.

### GetSection

`func (o *TeamPermissionsPutPreviousValues) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *TeamPermissionsPutPreviousValues) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *TeamPermissionsPutPreviousValues) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *TeamPermissionsPutPreviousValues) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetKey

`func (o *TeamPermissionsPutPreviousValues) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TeamPermissionsPutPreviousValues) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TeamPermissionsPutPreviousValues) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TeamPermissionsPutPreviousValues) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TeamPermissionsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TeamPermissionsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TeamPermissionsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TeamPermissionsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TeamPermissionsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TeamPermissionsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TeamPermissionsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TeamPermissionsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


