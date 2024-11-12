# RolePermissionsPutPreviousValues

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
**GlobalPermissionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;global_permissions.id&#x60;.&lt;fk table&#x3D;&#39;global_permissions&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Scope** | Pointer to **interface{}** |  | [optional] 
**CoreModules** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewRolePermissionsPutPreviousValues

`func NewRolePermissionsPutPreviousValues() *RolePermissionsPutPreviousValues`

NewRolePermissionsPutPreviousValues instantiates a new RolePermissionsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePermissionsPutPreviousValuesWithDefaults

`func NewRolePermissionsPutPreviousValuesWithDefaults() *RolePermissionsPutPreviousValues`

NewRolePermissionsPutPreviousValuesWithDefaults instantiates a new RolePermissionsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RolePermissionsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RolePermissionsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RolePermissionsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RolePermissionsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRoleId

`func (o *RolePermissionsPutPreviousValues) GetRoleId() int32`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *RolePermissionsPutPreviousValues) GetRoleIdOk() (*int32, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *RolePermissionsPutPreviousValues) SetRoleId(v int32)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *RolePermissionsPutPreviousValues) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetSection

`func (o *RolePermissionsPutPreviousValues) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *RolePermissionsPutPreviousValues) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *RolePermissionsPutPreviousValues) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *RolePermissionsPutPreviousValues) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetKey

`func (o *RolePermissionsPutPreviousValues) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RolePermissionsPutPreviousValues) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RolePermissionsPutPreviousValues) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RolePermissionsPutPreviousValues) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *RolePermissionsPutPreviousValues) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RolePermissionsPutPreviousValues) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RolePermissionsPutPreviousValues) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RolePermissionsPutPreviousValues) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RolePermissionsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RolePermissionsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RolePermissionsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RolePermissionsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RolePermissionsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RolePermissionsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RolePermissionsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RolePermissionsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *RolePermissionsPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *RolePermissionsPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *RolePermissionsPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *RolePermissionsPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetGlobalPermissionId

`func (o *RolePermissionsPutPreviousValues) GetGlobalPermissionId() int32`

GetGlobalPermissionId returns the GlobalPermissionId field if non-nil, zero value otherwise.

### GetGlobalPermissionIdOk

`func (o *RolePermissionsPutPreviousValues) GetGlobalPermissionIdOk() (*int32, bool)`

GetGlobalPermissionIdOk returns a tuple with the GlobalPermissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPermissionId

`func (o *RolePermissionsPutPreviousValues) SetGlobalPermissionId(v int32)`

SetGlobalPermissionId sets GlobalPermissionId field to given value.

### HasGlobalPermissionId

`func (o *RolePermissionsPutPreviousValues) HasGlobalPermissionId() bool`

HasGlobalPermissionId returns a boolean if a field has been set.

### GetScope

`func (o *RolePermissionsPutPreviousValues) GetScope() interface{}`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RolePermissionsPutPreviousValues) GetScopeOk() (*interface{}, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *RolePermissionsPutPreviousValues) SetScope(v interface{})`

SetScope sets Scope field to given value.

### HasScope

`func (o *RolePermissionsPutPreviousValues) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *RolePermissionsPutPreviousValues) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *RolePermissionsPutPreviousValues) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetCoreModules

`func (o *RolePermissionsPutPreviousValues) GetCoreModules() interface{}`

GetCoreModules returns the CoreModules field if non-nil, zero value otherwise.

### GetCoreModulesOk

`func (o *RolePermissionsPutPreviousValues) GetCoreModulesOk() (*interface{}, bool)`

GetCoreModulesOk returns a tuple with the CoreModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreModules

`func (o *RolePermissionsPutPreviousValues) SetCoreModules(v interface{})`

SetCoreModules sets CoreModules field to given value.

### HasCoreModules

`func (o *RolePermissionsPutPreviousValues) HasCoreModules() bool`

HasCoreModules returns a boolean if a field has been set.

### SetCoreModulesNil

`func (o *RolePermissionsPutPreviousValues) SetCoreModulesNil(b bool)`

 SetCoreModulesNil sets the value for CoreModules to be an explicit nil

### UnsetCoreModules
`func (o *RolePermissionsPutPreviousValues) UnsetCoreModules()`

UnsetCoreModules ensures that no value is present for CoreModules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


