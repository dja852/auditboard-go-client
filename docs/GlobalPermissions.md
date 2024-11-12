# GlobalPermissions

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

### NewGlobalPermissions

`func NewGlobalPermissions(coreModules interface{}, ) *GlobalPermissions`

NewGlobalPermissions instantiates a new GlobalPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalPermissionsWithDefaults

`func NewGlobalPermissionsWithDefaults() *GlobalPermissions`

NewGlobalPermissionsWithDefaults instantiates a new GlobalPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GlobalPermissions) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalPermissions) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalPermissions) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GlobalPermissions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSection

`func (o *GlobalPermissions) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *GlobalPermissions) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *GlobalPermissions) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *GlobalPermissions) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetKey

`func (o *GlobalPermissions) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GlobalPermissions) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GlobalPermissions) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GlobalPermissions) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *GlobalPermissions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalPermissions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalPermissions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GlobalPermissions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GlobalPermissions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GlobalPermissions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GlobalPermissions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GlobalPermissions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultValue

`func (o *GlobalPermissions) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *GlobalPermissions) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *GlobalPermissions) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *GlobalPermissions) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetAllowedValues

`func (o *GlobalPermissions) GetAllowedValues() string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *GlobalPermissions) GetAllowedValuesOk() (*string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *GlobalPermissions) SetAllowedValues(v string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *GlobalPermissions) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### GetAllValues

`func (o *GlobalPermissions) GetAllValues() string`

GetAllValues returns the AllValues field if non-nil, zero value otherwise.

### GetAllValuesOk

`func (o *GlobalPermissions) GetAllValuesOk() (*string, bool)`

GetAllValuesOk returns a tuple with the AllValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllValues

`func (o *GlobalPermissions) SetAllValues(v string)`

SetAllValues sets AllValues field to given value.

### HasAllValues

`func (o *GlobalPermissions) HasAllValues() bool`

HasAllValues returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GlobalPermissions) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GlobalPermissions) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GlobalPermissions) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GlobalPermissions) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GlobalPermissions) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GlobalPermissions) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GlobalPermissions) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GlobalPermissions) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *GlobalPermissions) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *GlobalPermissions) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *GlobalPermissions) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *GlobalPermissions) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetHideScopeEditor

`func (o *GlobalPermissions) GetHideScopeEditor() bool`

GetHideScopeEditor returns the HideScopeEditor field if non-nil, zero value otherwise.

### GetHideScopeEditorOk

`func (o *GlobalPermissions) GetHideScopeEditorOk() (*bool, bool)`

GetHideScopeEditorOk returns a tuple with the HideScopeEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideScopeEditor

`func (o *GlobalPermissions) SetHideScopeEditor(v bool)`

SetHideScopeEditor sets HideScopeEditor field to given value.

### HasHideScopeEditor

`func (o *GlobalPermissions) HasHideScopeEditor() bool`

HasHideScopeEditor returns a boolean if a field has been set.

### GetCoreModules

`func (o *GlobalPermissions) GetCoreModules() interface{}`

GetCoreModules returns the CoreModules field if non-nil, zero value otherwise.

### GetCoreModulesOk

`func (o *GlobalPermissions) GetCoreModulesOk() (*interface{}, bool)`

GetCoreModulesOk returns a tuple with the CoreModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreModules

`func (o *GlobalPermissions) SetCoreModules(v interface{})`

SetCoreModules sets CoreModules field to given value.


### SetCoreModulesNil

`func (o *GlobalPermissions) SetCoreModulesNil(b bool)`

 SetCoreModulesNil sets the value for CoreModules to be an explicit nil

### UnsetCoreModules
`func (o *GlobalPermissions) UnsetCoreModules()`

UnsetCoreModules ensures that no value is present for CoreModules, not even an explicit nil
### GetRequireRoleLevel

`func (o *GlobalPermissions) GetRequireRoleLevel() string`

GetRequireRoleLevel returns the RequireRoleLevel field if non-nil, zero value otherwise.

### GetRequireRoleLevelOk

`func (o *GlobalPermissions) GetRequireRoleLevelOk() (*string, bool)`

GetRequireRoleLevelOk returns a tuple with the RequireRoleLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireRoleLevel

`func (o *GlobalPermissions) SetRequireRoleLevel(v string)`

SetRequireRoleLevel sets RequireRoleLevel field to given value.

### HasRequireRoleLevel

`func (o *GlobalPermissions) HasRequireRoleLevel() bool`

HasRequireRoleLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


