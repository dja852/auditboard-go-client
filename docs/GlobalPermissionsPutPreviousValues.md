# GlobalPermissionsPutPreviousValues

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
**CoreModules** | Pointer to **interface{}** |  | [optional] 
**RequireRoleLevel** | Pointer to **string** |  | [optional] 

## Methods

### NewGlobalPermissionsPutPreviousValues

`func NewGlobalPermissionsPutPreviousValues() *GlobalPermissionsPutPreviousValues`

NewGlobalPermissionsPutPreviousValues instantiates a new GlobalPermissionsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalPermissionsPutPreviousValuesWithDefaults

`func NewGlobalPermissionsPutPreviousValuesWithDefaults() *GlobalPermissionsPutPreviousValues`

NewGlobalPermissionsPutPreviousValuesWithDefaults instantiates a new GlobalPermissionsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GlobalPermissionsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalPermissionsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalPermissionsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GlobalPermissionsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSection

`func (o *GlobalPermissionsPutPreviousValues) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *GlobalPermissionsPutPreviousValues) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *GlobalPermissionsPutPreviousValues) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *GlobalPermissionsPutPreviousValues) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetKey

`func (o *GlobalPermissionsPutPreviousValues) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GlobalPermissionsPutPreviousValues) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GlobalPermissionsPutPreviousValues) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GlobalPermissionsPutPreviousValues) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *GlobalPermissionsPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalPermissionsPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalPermissionsPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GlobalPermissionsPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GlobalPermissionsPutPreviousValues) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GlobalPermissionsPutPreviousValues) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GlobalPermissionsPutPreviousValues) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GlobalPermissionsPutPreviousValues) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultValue

`func (o *GlobalPermissionsPutPreviousValues) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *GlobalPermissionsPutPreviousValues) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *GlobalPermissionsPutPreviousValues) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *GlobalPermissionsPutPreviousValues) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetAllowedValues

`func (o *GlobalPermissionsPutPreviousValues) GetAllowedValues() string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *GlobalPermissionsPutPreviousValues) GetAllowedValuesOk() (*string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *GlobalPermissionsPutPreviousValues) SetAllowedValues(v string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *GlobalPermissionsPutPreviousValues) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### GetAllValues

`func (o *GlobalPermissionsPutPreviousValues) GetAllValues() string`

GetAllValues returns the AllValues field if non-nil, zero value otherwise.

### GetAllValuesOk

`func (o *GlobalPermissionsPutPreviousValues) GetAllValuesOk() (*string, bool)`

GetAllValuesOk returns a tuple with the AllValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllValues

`func (o *GlobalPermissionsPutPreviousValues) SetAllValues(v string)`

SetAllValues sets AllValues field to given value.

### HasAllValues

`func (o *GlobalPermissionsPutPreviousValues) HasAllValues() bool`

HasAllValues returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GlobalPermissionsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GlobalPermissionsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GlobalPermissionsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GlobalPermissionsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GlobalPermissionsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GlobalPermissionsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GlobalPermissionsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GlobalPermissionsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *GlobalPermissionsPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *GlobalPermissionsPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *GlobalPermissionsPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *GlobalPermissionsPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetHideScopeEditor

`func (o *GlobalPermissionsPutPreviousValues) GetHideScopeEditor() bool`

GetHideScopeEditor returns the HideScopeEditor field if non-nil, zero value otherwise.

### GetHideScopeEditorOk

`func (o *GlobalPermissionsPutPreviousValues) GetHideScopeEditorOk() (*bool, bool)`

GetHideScopeEditorOk returns a tuple with the HideScopeEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideScopeEditor

`func (o *GlobalPermissionsPutPreviousValues) SetHideScopeEditor(v bool)`

SetHideScopeEditor sets HideScopeEditor field to given value.

### HasHideScopeEditor

`func (o *GlobalPermissionsPutPreviousValues) HasHideScopeEditor() bool`

HasHideScopeEditor returns a boolean if a field has been set.

### GetCoreModules

`func (o *GlobalPermissionsPutPreviousValues) GetCoreModules() interface{}`

GetCoreModules returns the CoreModules field if non-nil, zero value otherwise.

### GetCoreModulesOk

`func (o *GlobalPermissionsPutPreviousValues) GetCoreModulesOk() (*interface{}, bool)`

GetCoreModulesOk returns a tuple with the CoreModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreModules

`func (o *GlobalPermissionsPutPreviousValues) SetCoreModules(v interface{})`

SetCoreModules sets CoreModules field to given value.

### HasCoreModules

`func (o *GlobalPermissionsPutPreviousValues) HasCoreModules() bool`

HasCoreModules returns a boolean if a field has been set.

### SetCoreModulesNil

`func (o *GlobalPermissionsPutPreviousValues) SetCoreModulesNil(b bool)`

 SetCoreModulesNil sets the value for CoreModules to be an explicit nil

### UnsetCoreModules
`func (o *GlobalPermissionsPutPreviousValues) UnsetCoreModules()`

UnsetCoreModules ensures that no value is present for CoreModules, not even an explicit nil
### GetRequireRoleLevel

`func (o *GlobalPermissionsPutPreviousValues) GetRequireRoleLevel() string`

GetRequireRoleLevel returns the RequireRoleLevel field if non-nil, zero value otherwise.

### GetRequireRoleLevelOk

`func (o *GlobalPermissionsPutPreviousValues) GetRequireRoleLevelOk() (*string, bool)`

GetRequireRoleLevelOk returns a tuple with the RequireRoleLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireRoleLevel

`func (o *GlobalPermissionsPutPreviousValues) SetRequireRoleLevel(v string)`

SetRequireRoleLevel sets RequireRoleLevel field to given value.

### HasRequireRoleLevel

`func (o *GlobalPermissionsPutPreviousValues) HasRequireRoleLevel() bool`

HasRequireRoleLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


