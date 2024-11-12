# RolesPutRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**ShortName** | **string** |  | 
**DefaultRoute** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsImplementer** | Pointer to **bool** |  | [optional] [default to false]
**IsLimited** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewRolesPutRole

`func NewRolesPutRole(name string, shortName string, ) *RolesPutRole`

NewRolesPutRole instantiates a new RolesPutRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolesPutRoleWithDefaults

`func NewRolesPutRoleWithDefaults() *RolesPutRole`

NewRolesPutRoleWithDefaults instantiates a new RolesPutRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RolesPutRole) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RolesPutRole) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RolesPutRole) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RolesPutRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RolesPutRole) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RolesPutRole) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RolesPutRole) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RolesPutRole) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RolesPutRole) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RolesPutRole) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RolesPutRole) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RolesPutRole) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *RolesPutRole) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *RolesPutRole) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *RolesPutRole) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *RolesPutRole) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *RolesPutRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RolesPutRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RolesPutRole) SetName(v string)`

SetName sets Name field to given value.


### GetShortName

`func (o *RolesPutRole) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *RolesPutRole) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *RolesPutRole) SetShortName(v string)`

SetShortName sets ShortName field to given value.


### GetDefaultRoute

`func (o *RolesPutRole) GetDefaultRoute() string`

GetDefaultRoute returns the DefaultRoute field if non-nil, zero value otherwise.

### GetDefaultRouteOk

`func (o *RolesPutRole) GetDefaultRouteOk() (*string, bool)`

GetDefaultRouteOk returns a tuple with the DefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoute

`func (o *RolesPutRole) SetDefaultRoute(v string)`

SetDefaultRoute sets DefaultRoute field to given value.

### HasDefaultRoute

`func (o *RolesPutRole) HasDefaultRoute() bool`

HasDefaultRoute returns a boolean if a field has been set.

### GetDescription

`func (o *RolesPutRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RolesPutRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RolesPutRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RolesPutRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsImplementer

`func (o *RolesPutRole) GetIsImplementer() bool`

GetIsImplementer returns the IsImplementer field if non-nil, zero value otherwise.

### GetIsImplementerOk

`func (o *RolesPutRole) GetIsImplementerOk() (*bool, bool)`

GetIsImplementerOk returns a tuple with the IsImplementer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImplementer

`func (o *RolesPutRole) SetIsImplementer(v bool)`

SetIsImplementer sets IsImplementer field to given value.

### HasIsImplementer

`func (o *RolesPutRole) HasIsImplementer() bool`

HasIsImplementer returns a boolean if a field has been set.

### GetIsLimited

`func (o *RolesPutRole) GetIsLimited() bool`

GetIsLimited returns the IsLimited field if non-nil, zero value otherwise.

### GetIsLimitedOk

`func (o *RolesPutRole) GetIsLimitedOk() (*bool, bool)`

GetIsLimitedOk returns a tuple with the IsLimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLimited

`func (o *RolesPutRole) SetIsLimited(v bool)`

SetIsLimited sets IsLimited field to given value.

### HasIsLimited

`func (o *RolesPutRole) HasIsLimited() bool`

HasIsLimited returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


