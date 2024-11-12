# Roles

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

### NewRoles

`func NewRoles(name string, shortName string, ) *Roles`

NewRoles instantiates a new Roles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolesWithDefaults

`func NewRolesWithDefaults() *Roles`

NewRolesWithDefaults instantiates a new Roles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Roles) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Roles) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Roles) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Roles) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Roles) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Roles) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Roles) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Roles) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Roles) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Roles) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Roles) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Roles) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Roles) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Roles) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Roles) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Roles) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *Roles) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Roles) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Roles) SetName(v string)`

SetName sets Name field to given value.


### GetShortName

`func (o *Roles) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *Roles) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *Roles) SetShortName(v string)`

SetShortName sets ShortName field to given value.


### GetDefaultRoute

`func (o *Roles) GetDefaultRoute() string`

GetDefaultRoute returns the DefaultRoute field if non-nil, zero value otherwise.

### GetDefaultRouteOk

`func (o *Roles) GetDefaultRouteOk() (*string, bool)`

GetDefaultRouteOk returns a tuple with the DefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoute

`func (o *Roles) SetDefaultRoute(v string)`

SetDefaultRoute sets DefaultRoute field to given value.

### HasDefaultRoute

`func (o *Roles) HasDefaultRoute() bool`

HasDefaultRoute returns a boolean if a field has been set.

### GetDescription

`func (o *Roles) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Roles) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Roles) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Roles) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsImplementer

`func (o *Roles) GetIsImplementer() bool`

GetIsImplementer returns the IsImplementer field if non-nil, zero value otherwise.

### GetIsImplementerOk

`func (o *Roles) GetIsImplementerOk() (*bool, bool)`

GetIsImplementerOk returns a tuple with the IsImplementer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImplementer

`func (o *Roles) SetIsImplementer(v bool)`

SetIsImplementer sets IsImplementer field to given value.

### HasIsImplementer

`func (o *Roles) HasIsImplementer() bool`

HasIsImplementer returns a boolean if a field has been set.

### GetIsLimited

`func (o *Roles) GetIsLimited() bool`

GetIsLimited returns the IsLimited field if non-nil, zero value otherwise.

### GetIsLimitedOk

`func (o *Roles) GetIsLimitedOk() (*bool, bool)`

GetIsLimitedOk returns a tuple with the IsLimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLimited

`func (o *Roles) SetIsLimited(v bool)`

SetIsLimited sets IsLimited field to given value.

### HasIsLimited

`func (o *Roles) HasIsLimited() bool`

HasIsLimited returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


