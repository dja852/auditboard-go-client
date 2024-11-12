# RolePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**RoleId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;roles.id&#x60;.&lt;fk table&#x3D;&#39;roles&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Section** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**GlobalPermissionId** | **int32** | Note: This is a Foreign Key to &#x60;global_permissions.id&#x60;.&lt;fk table&#x3D;&#39;global_permissions&#39; column&#x3D;&#39;id&#39;/&gt; | 
**Scope** | Pointer to **interface{}** |  | [optional] 
**CoreModules** | **interface{}** |  | 

## Methods

### NewRolePermissions

`func NewRolePermissions(globalPermissionId int32, coreModules interface{}, ) *RolePermissions`

NewRolePermissions instantiates a new RolePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePermissionsWithDefaults

`func NewRolePermissionsWithDefaults() *RolePermissions`

NewRolePermissionsWithDefaults instantiates a new RolePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RolePermissions) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RolePermissions) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RolePermissions) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RolePermissions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRoleId

`func (o *RolePermissions) GetRoleId() int32`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *RolePermissions) GetRoleIdOk() (*int32, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *RolePermissions) SetRoleId(v int32)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *RolePermissions) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetSection

`func (o *RolePermissions) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *RolePermissions) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *RolePermissions) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *RolePermissions) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetKey

`func (o *RolePermissions) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RolePermissions) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RolePermissions) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RolePermissions) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *RolePermissions) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RolePermissions) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RolePermissions) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RolePermissions) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RolePermissions) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RolePermissions) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RolePermissions) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RolePermissions) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RolePermissions) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RolePermissions) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RolePermissions) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RolePermissions) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *RolePermissions) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *RolePermissions) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *RolePermissions) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *RolePermissions) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetGlobalPermissionId

`func (o *RolePermissions) GetGlobalPermissionId() int32`

GetGlobalPermissionId returns the GlobalPermissionId field if non-nil, zero value otherwise.

### GetGlobalPermissionIdOk

`func (o *RolePermissions) GetGlobalPermissionIdOk() (*int32, bool)`

GetGlobalPermissionIdOk returns a tuple with the GlobalPermissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPermissionId

`func (o *RolePermissions) SetGlobalPermissionId(v int32)`

SetGlobalPermissionId sets GlobalPermissionId field to given value.


### GetScope

`func (o *RolePermissions) GetScope() interface{}`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RolePermissions) GetScopeOk() (*interface{}, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *RolePermissions) SetScope(v interface{})`

SetScope sets Scope field to given value.

### HasScope

`func (o *RolePermissions) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *RolePermissions) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *RolePermissions) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetCoreModules

`func (o *RolePermissions) GetCoreModules() interface{}`

GetCoreModules returns the CoreModules field if non-nil, zero value otherwise.

### GetCoreModulesOk

`func (o *RolePermissions) GetCoreModulesOk() (*interface{}, bool)`

GetCoreModulesOk returns a tuple with the CoreModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreModules

`func (o *RolePermissions) SetCoreModules(v interface{})`

SetCoreModules sets CoreModules field to given value.


### SetCoreModulesNil

`func (o *RolePermissions) SetCoreModulesNil(b bool)`

 SetCoreModulesNil sets the value for CoreModules to be an explicit nil

### UnsetCoreModules
`func (o *RolePermissions) UnsetCoreModules()`

UnsetCoreModules ensures that no value is present for CoreModules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


