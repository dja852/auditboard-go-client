# RolePermissionsPutRolePermission

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

### NewRolePermissionsPutRolePermission

`func NewRolePermissionsPutRolePermission(globalPermissionId int32, coreModules interface{}, ) *RolePermissionsPutRolePermission`

NewRolePermissionsPutRolePermission instantiates a new RolePermissionsPutRolePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePermissionsPutRolePermissionWithDefaults

`func NewRolePermissionsPutRolePermissionWithDefaults() *RolePermissionsPutRolePermission`

NewRolePermissionsPutRolePermissionWithDefaults instantiates a new RolePermissionsPutRolePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RolePermissionsPutRolePermission) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RolePermissionsPutRolePermission) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RolePermissionsPutRolePermission) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RolePermissionsPutRolePermission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRoleId

`func (o *RolePermissionsPutRolePermission) GetRoleId() int32`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *RolePermissionsPutRolePermission) GetRoleIdOk() (*int32, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *RolePermissionsPutRolePermission) SetRoleId(v int32)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *RolePermissionsPutRolePermission) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetSection

`func (o *RolePermissionsPutRolePermission) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *RolePermissionsPutRolePermission) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *RolePermissionsPutRolePermission) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *RolePermissionsPutRolePermission) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetKey

`func (o *RolePermissionsPutRolePermission) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RolePermissionsPutRolePermission) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RolePermissionsPutRolePermission) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RolePermissionsPutRolePermission) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *RolePermissionsPutRolePermission) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RolePermissionsPutRolePermission) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RolePermissionsPutRolePermission) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RolePermissionsPutRolePermission) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RolePermissionsPutRolePermission) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RolePermissionsPutRolePermission) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RolePermissionsPutRolePermission) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RolePermissionsPutRolePermission) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RolePermissionsPutRolePermission) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RolePermissionsPutRolePermission) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RolePermissionsPutRolePermission) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RolePermissionsPutRolePermission) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *RolePermissionsPutRolePermission) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *RolePermissionsPutRolePermission) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *RolePermissionsPutRolePermission) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *RolePermissionsPutRolePermission) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetGlobalPermissionId

`func (o *RolePermissionsPutRolePermission) GetGlobalPermissionId() int32`

GetGlobalPermissionId returns the GlobalPermissionId field if non-nil, zero value otherwise.

### GetGlobalPermissionIdOk

`func (o *RolePermissionsPutRolePermission) GetGlobalPermissionIdOk() (*int32, bool)`

GetGlobalPermissionIdOk returns a tuple with the GlobalPermissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPermissionId

`func (o *RolePermissionsPutRolePermission) SetGlobalPermissionId(v int32)`

SetGlobalPermissionId sets GlobalPermissionId field to given value.


### GetScope

`func (o *RolePermissionsPutRolePermission) GetScope() interface{}`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RolePermissionsPutRolePermission) GetScopeOk() (*interface{}, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *RolePermissionsPutRolePermission) SetScope(v interface{})`

SetScope sets Scope field to given value.

### HasScope

`func (o *RolePermissionsPutRolePermission) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *RolePermissionsPutRolePermission) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *RolePermissionsPutRolePermission) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetCoreModules

`func (o *RolePermissionsPutRolePermission) GetCoreModules() interface{}`

GetCoreModules returns the CoreModules field if non-nil, zero value otherwise.

### GetCoreModulesOk

`func (o *RolePermissionsPutRolePermission) GetCoreModulesOk() (*interface{}, bool)`

GetCoreModulesOk returns a tuple with the CoreModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreModules

`func (o *RolePermissionsPutRolePermission) SetCoreModules(v interface{})`

SetCoreModules sets CoreModules field to given value.


### SetCoreModulesNil

`func (o *RolePermissionsPutRolePermission) SetCoreModulesNil(b bool)`

 SetCoreModulesNil sets the value for CoreModules to be an explicit nil

### UnsetCoreModules
`func (o *RolePermissionsPutRolePermission) UnsetCoreModules()`

UnsetCoreModules ensures that no value is present for CoreModules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


