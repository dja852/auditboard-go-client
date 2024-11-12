# GlobalAttributesPutGlobalAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Section** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**DefaultLabel** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**Enabled** | **bool** |  | [default to true]
**Scope** | **interface{}** |  | 

## Methods

### NewGlobalAttributesPutGlobalAttribute

`func NewGlobalAttributesPutGlobalAttribute(enabled bool, scope interface{}, ) *GlobalAttributesPutGlobalAttribute`

NewGlobalAttributesPutGlobalAttribute instantiates a new GlobalAttributesPutGlobalAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalAttributesPutGlobalAttributeWithDefaults

`func NewGlobalAttributesPutGlobalAttributeWithDefaults() *GlobalAttributesPutGlobalAttribute`

NewGlobalAttributesPutGlobalAttributeWithDefaults instantiates a new GlobalAttributesPutGlobalAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GlobalAttributesPutGlobalAttribute) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalAttributesPutGlobalAttribute) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalAttributesPutGlobalAttribute) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GlobalAttributesPutGlobalAttribute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSection

`func (o *GlobalAttributesPutGlobalAttribute) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *GlobalAttributesPutGlobalAttribute) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *GlobalAttributesPutGlobalAttribute) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *GlobalAttributesPutGlobalAttribute) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetObjectType

`func (o *GlobalAttributesPutGlobalAttribute) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *GlobalAttributesPutGlobalAttribute) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *GlobalAttributesPutGlobalAttribute) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *GlobalAttributesPutGlobalAttribute) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetKey

`func (o *GlobalAttributesPutGlobalAttribute) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GlobalAttributesPutGlobalAttribute) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GlobalAttributesPutGlobalAttribute) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GlobalAttributesPutGlobalAttribute) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetDefaultLabel

`func (o *GlobalAttributesPutGlobalAttribute) GetDefaultLabel() string`

GetDefaultLabel returns the DefaultLabel field if non-nil, zero value otherwise.

### GetDefaultLabelOk

`func (o *GlobalAttributesPutGlobalAttribute) GetDefaultLabelOk() (*string, bool)`

GetDefaultLabelOk returns a tuple with the DefaultLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLabel

`func (o *GlobalAttributesPutGlobalAttribute) SetDefaultLabel(v string)`

SetDefaultLabel sets DefaultLabel field to given value.

### HasDefaultLabel

`func (o *GlobalAttributesPutGlobalAttribute) HasDefaultLabel() bool`

HasDefaultLabel returns a boolean if a field has been set.

### GetLabel

`func (o *GlobalAttributesPutGlobalAttribute) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GlobalAttributesPutGlobalAttribute) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GlobalAttributesPutGlobalAttribute) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GlobalAttributesPutGlobalAttribute) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GlobalAttributesPutGlobalAttribute) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GlobalAttributesPutGlobalAttribute) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GlobalAttributesPutGlobalAttribute) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GlobalAttributesPutGlobalAttribute) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GlobalAttributesPutGlobalAttribute) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GlobalAttributesPutGlobalAttribute) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GlobalAttributesPutGlobalAttribute) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GlobalAttributesPutGlobalAttribute) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *GlobalAttributesPutGlobalAttribute) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *GlobalAttributesPutGlobalAttribute) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *GlobalAttributesPutGlobalAttribute) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *GlobalAttributesPutGlobalAttribute) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetEnabled

`func (o *GlobalAttributesPutGlobalAttribute) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GlobalAttributesPutGlobalAttribute) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GlobalAttributesPutGlobalAttribute) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetScope

`func (o *GlobalAttributesPutGlobalAttribute) GetScope() interface{}`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GlobalAttributesPutGlobalAttribute) GetScopeOk() (*interface{}, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GlobalAttributesPutGlobalAttribute) SetScope(v interface{})`

SetScope sets Scope field to given value.


### SetScopeNil

`func (o *GlobalAttributesPutGlobalAttribute) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *GlobalAttributesPutGlobalAttribute) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


