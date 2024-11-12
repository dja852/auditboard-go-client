# GlobalAttributes

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

### NewGlobalAttributes

`func NewGlobalAttributes(enabled bool, scope interface{}, ) *GlobalAttributes`

NewGlobalAttributes instantiates a new GlobalAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalAttributesWithDefaults

`func NewGlobalAttributesWithDefaults() *GlobalAttributes`

NewGlobalAttributesWithDefaults instantiates a new GlobalAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GlobalAttributes) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalAttributes) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalAttributes) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GlobalAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSection

`func (o *GlobalAttributes) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *GlobalAttributes) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *GlobalAttributes) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *GlobalAttributes) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetObjectType

`func (o *GlobalAttributes) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *GlobalAttributes) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *GlobalAttributes) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *GlobalAttributes) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetKey

`func (o *GlobalAttributes) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GlobalAttributes) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GlobalAttributes) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GlobalAttributes) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetDefaultLabel

`func (o *GlobalAttributes) GetDefaultLabel() string`

GetDefaultLabel returns the DefaultLabel field if non-nil, zero value otherwise.

### GetDefaultLabelOk

`func (o *GlobalAttributes) GetDefaultLabelOk() (*string, bool)`

GetDefaultLabelOk returns a tuple with the DefaultLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLabel

`func (o *GlobalAttributes) SetDefaultLabel(v string)`

SetDefaultLabel sets DefaultLabel field to given value.

### HasDefaultLabel

`func (o *GlobalAttributes) HasDefaultLabel() bool`

HasDefaultLabel returns a boolean if a field has been set.

### GetLabel

`func (o *GlobalAttributes) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GlobalAttributes) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GlobalAttributes) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GlobalAttributes) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GlobalAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GlobalAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GlobalAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GlobalAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GlobalAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GlobalAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GlobalAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GlobalAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *GlobalAttributes) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *GlobalAttributes) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *GlobalAttributes) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *GlobalAttributes) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetEnabled

`func (o *GlobalAttributes) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GlobalAttributes) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GlobalAttributes) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetScope

`func (o *GlobalAttributes) GetScope() interface{}`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GlobalAttributes) GetScopeOk() (*interface{}, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GlobalAttributes) SetScope(v interface{})`

SetScope sets Scope field to given value.


### SetScopeNil

`func (o *GlobalAttributes) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *GlobalAttributes) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


