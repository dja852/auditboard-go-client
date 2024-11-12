# GlobalPermissionsPutGlobalPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Section** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DefaultValue** | Pointer to **string** |  | [optional] 
**AllowedValues** | Pointer to **string** |  | [optional] 
**AllValues** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**HideScopeEditor** | Pointer to **bool** |  | [optional] [default to false]
**CoreModules** | **interface{}** |  | 
**RequireRoleLevel** | Pointer to **string** |  | [optional] 

## Methods

### NewGlobalPermissionsPutGlobalPermission

`func NewGlobalPermissionsPutGlobalPermission(coreModules interface{}, ) *GlobalPermissionsPutGlobalPermission`

NewGlobalPermissionsPutGlobalPermission instantiates a new GlobalPermissionsPutGlobalPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalPermissionsPutGlobalPermissionWithDefaults

`func NewGlobalPermissionsPutGlobalPermissionWithDefaults() *GlobalPermissionsPutGlobalPermission`

NewGlobalPermissionsPutGlobalPermissionWithDefaults instantiates a new GlobalPermissionsPutGlobalPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GlobalPermissionsPutGlobalPermission) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalPermissionsPutGlobalPermission) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalPermissionsPutGlobalPermission) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GlobalPermissionsPutGlobalPermission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSection

`func (o *GlobalPermissionsPutGlobalPermission) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *GlobalPermissionsPutGlobalPermission) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *GlobalPermissionsPutGlobalPermission) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *GlobalPermissionsPutGlobalPermission) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetKey

`func (o *GlobalPermissionsPutGlobalPermission) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GlobalPermissionsPutGlobalPermission) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GlobalPermissionsPutGlobalPermission) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GlobalPermissionsPutGlobalPermission) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *GlobalPermissionsPutGlobalPermission) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalPermissionsPutGlobalPermission) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalPermissionsPutGlobalPermission) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GlobalPermissionsPutGlobalPermission) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GlobalPermissionsPutGlobalPermission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GlobalPermissionsPutGlobalPermission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GlobalPermissionsPutGlobalPermission) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GlobalPermissionsPutGlobalPermission) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultValue

`func (o *GlobalPermissionsPutGlobalPermission) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *GlobalPermissionsPutGlobalPermission) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *GlobalPermissionsPutGlobalPermission) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *GlobalPermissionsPutGlobalPermission) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetAllowedValues

`func (o *GlobalPermissionsPutGlobalPermission) GetAllowedValues() string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *GlobalPermissionsPutGlobalPermission) GetAllowedValuesOk() (*string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *GlobalPermissionsPutGlobalPermission) SetAllowedValues(v string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *GlobalPermissionsPutGlobalPermission) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### GetAllValues

`func (o *GlobalPermissionsPutGlobalPermission) GetAllValues() string`

GetAllValues returns the AllValues field if non-nil, zero value otherwise.

### GetAllValuesOk

`func (o *GlobalPermissionsPutGlobalPermission) GetAllValuesOk() (*string, bool)`

GetAllValuesOk returns a tuple with the AllValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllValues

`func (o *GlobalPermissionsPutGlobalPermission) SetAllValues(v string)`

SetAllValues sets AllValues field to given value.

### HasAllValues

`func (o *GlobalPermissionsPutGlobalPermission) HasAllValues() bool`

HasAllValues returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GlobalPermissionsPutGlobalPermission) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GlobalPermissionsPutGlobalPermission) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GlobalPermissionsPutGlobalPermission) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GlobalPermissionsPutGlobalPermission) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GlobalPermissionsPutGlobalPermission) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GlobalPermissionsPutGlobalPermission) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GlobalPermissionsPutGlobalPermission) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GlobalPermissionsPutGlobalPermission) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *GlobalPermissionsPutGlobalPermission) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *GlobalPermissionsPutGlobalPermission) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *GlobalPermissionsPutGlobalPermission) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *GlobalPermissionsPutGlobalPermission) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetHideScopeEditor

`func (o *GlobalPermissionsPutGlobalPermission) GetHideScopeEditor() bool`

GetHideScopeEditor returns the HideScopeEditor field if non-nil, zero value otherwise.

### GetHideScopeEditorOk

`func (o *GlobalPermissionsPutGlobalPermission) GetHideScopeEditorOk() (*bool, bool)`

GetHideScopeEditorOk returns a tuple with the HideScopeEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideScopeEditor

`func (o *GlobalPermissionsPutGlobalPermission) SetHideScopeEditor(v bool)`

SetHideScopeEditor sets HideScopeEditor field to given value.

### HasHideScopeEditor

`func (o *GlobalPermissionsPutGlobalPermission) HasHideScopeEditor() bool`

HasHideScopeEditor returns a boolean if a field has been set.

### GetCoreModules

`func (o *GlobalPermissionsPutGlobalPermission) GetCoreModules() interface{}`

GetCoreModules returns the CoreModules field if non-nil, zero value otherwise.

### GetCoreModulesOk

`func (o *GlobalPermissionsPutGlobalPermission) GetCoreModulesOk() (*interface{}, bool)`

GetCoreModulesOk returns a tuple with the CoreModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreModules

`func (o *GlobalPermissionsPutGlobalPermission) SetCoreModules(v interface{})`

SetCoreModules sets CoreModules field to given value.


### SetCoreModulesNil

`func (o *GlobalPermissionsPutGlobalPermission) SetCoreModulesNil(b bool)`

 SetCoreModulesNil sets the value for CoreModules to be an explicit nil

### UnsetCoreModules
`func (o *GlobalPermissionsPutGlobalPermission) UnsetCoreModules()`

UnsetCoreModules ensures that no value is present for CoreModules, not even an explicit nil
### GetRequireRoleLevel

`func (o *GlobalPermissionsPutGlobalPermission) GetRequireRoleLevel() string`

GetRequireRoleLevel returns the RequireRoleLevel field if non-nil, zero value otherwise.

### GetRequireRoleLevelOk

`func (o *GlobalPermissionsPutGlobalPermission) GetRequireRoleLevelOk() (*string, bool)`

GetRequireRoleLevelOk returns a tuple with the RequireRoleLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireRoleLevel

`func (o *GlobalPermissionsPutGlobalPermission) SetRequireRoleLevel(v string)`

SetRequireRoleLevel sets RequireRoleLevel field to given value.

### HasRequireRoleLevel

`func (o *GlobalPermissionsPutGlobalPermission) HasRequireRoleLevel() bool`

HasRequireRoleLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


