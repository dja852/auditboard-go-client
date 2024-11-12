# RolesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**DefaultRoute** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsImplementer** | Pointer to **bool** |  | [optional] [default to false]
**IsLimited** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewRolesPutPreviousValues

`func NewRolesPutPreviousValues() *RolesPutPreviousValues`

NewRolesPutPreviousValues instantiates a new RolesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolesPutPreviousValuesWithDefaults

`func NewRolesPutPreviousValuesWithDefaults() *RolesPutPreviousValues`

NewRolesPutPreviousValuesWithDefaults instantiates a new RolesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RolesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RolesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RolesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RolesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RolesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RolesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RolesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RolesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RolesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RolesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RolesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RolesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *RolesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *RolesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *RolesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *RolesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *RolesPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RolesPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RolesPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RolesPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShortName

`func (o *RolesPutPreviousValues) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *RolesPutPreviousValues) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *RolesPutPreviousValues) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *RolesPutPreviousValues) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetDefaultRoute

`func (o *RolesPutPreviousValues) GetDefaultRoute() string`

GetDefaultRoute returns the DefaultRoute field if non-nil, zero value otherwise.

### GetDefaultRouteOk

`func (o *RolesPutPreviousValues) GetDefaultRouteOk() (*string, bool)`

GetDefaultRouteOk returns a tuple with the DefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoute

`func (o *RolesPutPreviousValues) SetDefaultRoute(v string)`

SetDefaultRoute sets DefaultRoute field to given value.

### HasDefaultRoute

`func (o *RolesPutPreviousValues) HasDefaultRoute() bool`

HasDefaultRoute returns a boolean if a field has been set.

### GetDescription

`func (o *RolesPutPreviousValues) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RolesPutPreviousValues) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RolesPutPreviousValues) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RolesPutPreviousValues) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsImplementer

`func (o *RolesPutPreviousValues) GetIsImplementer() bool`

GetIsImplementer returns the IsImplementer field if non-nil, zero value otherwise.

### GetIsImplementerOk

`func (o *RolesPutPreviousValues) GetIsImplementerOk() (*bool, bool)`

GetIsImplementerOk returns a tuple with the IsImplementer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImplementer

`func (o *RolesPutPreviousValues) SetIsImplementer(v bool)`

SetIsImplementer sets IsImplementer field to given value.

### HasIsImplementer

`func (o *RolesPutPreviousValues) HasIsImplementer() bool`

HasIsImplementer returns a boolean if a field has been set.

### GetIsLimited

`func (o *RolesPutPreviousValues) GetIsLimited() bool`

GetIsLimited returns the IsLimited field if non-nil, zero value otherwise.

### GetIsLimitedOk

`func (o *RolesPutPreviousValues) GetIsLimitedOk() (*bool, bool)`

GetIsLimitedOk returns a tuple with the IsLimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLimited

`func (o *RolesPutPreviousValues) SetIsLimited(v bool)`

SetIsLimited sets IsLimited field to given value.

### HasIsLimited

`func (o *RolesPutPreviousValues) HasIsLimited() bool`

HasIsLimited returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


