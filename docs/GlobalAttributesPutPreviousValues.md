# GlobalAttributesPutPreviousValues

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
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**Scope** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewGlobalAttributesPutPreviousValues

`func NewGlobalAttributesPutPreviousValues() *GlobalAttributesPutPreviousValues`

NewGlobalAttributesPutPreviousValues instantiates a new GlobalAttributesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalAttributesPutPreviousValuesWithDefaults

`func NewGlobalAttributesPutPreviousValuesWithDefaults() *GlobalAttributesPutPreviousValues`

NewGlobalAttributesPutPreviousValuesWithDefaults instantiates a new GlobalAttributesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GlobalAttributesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalAttributesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalAttributesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GlobalAttributesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSection

`func (o *GlobalAttributesPutPreviousValues) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *GlobalAttributesPutPreviousValues) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *GlobalAttributesPutPreviousValues) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *GlobalAttributesPutPreviousValues) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetObjectType

`func (o *GlobalAttributesPutPreviousValues) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *GlobalAttributesPutPreviousValues) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *GlobalAttributesPutPreviousValues) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *GlobalAttributesPutPreviousValues) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetKey

`func (o *GlobalAttributesPutPreviousValues) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GlobalAttributesPutPreviousValues) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GlobalAttributesPutPreviousValues) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GlobalAttributesPutPreviousValues) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetDefaultLabel

`func (o *GlobalAttributesPutPreviousValues) GetDefaultLabel() string`

GetDefaultLabel returns the DefaultLabel field if non-nil, zero value otherwise.

### GetDefaultLabelOk

`func (o *GlobalAttributesPutPreviousValues) GetDefaultLabelOk() (*string, bool)`

GetDefaultLabelOk returns a tuple with the DefaultLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLabel

`func (o *GlobalAttributesPutPreviousValues) SetDefaultLabel(v string)`

SetDefaultLabel sets DefaultLabel field to given value.

### HasDefaultLabel

`func (o *GlobalAttributesPutPreviousValues) HasDefaultLabel() bool`

HasDefaultLabel returns a boolean if a field has been set.

### GetLabel

`func (o *GlobalAttributesPutPreviousValues) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GlobalAttributesPutPreviousValues) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GlobalAttributesPutPreviousValues) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GlobalAttributesPutPreviousValues) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GlobalAttributesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GlobalAttributesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GlobalAttributesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GlobalAttributesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GlobalAttributesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GlobalAttributesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GlobalAttributesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GlobalAttributesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *GlobalAttributesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *GlobalAttributesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *GlobalAttributesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *GlobalAttributesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetEnabled

`func (o *GlobalAttributesPutPreviousValues) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GlobalAttributesPutPreviousValues) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GlobalAttributesPutPreviousValues) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GlobalAttributesPutPreviousValues) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetScope

`func (o *GlobalAttributesPutPreviousValues) GetScope() interface{}`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GlobalAttributesPutPreviousValues) GetScopeOk() (*interface{}, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GlobalAttributesPutPreviousValues) SetScope(v interface{})`

SetScope sets Scope field to given value.

### HasScope

`func (o *GlobalAttributesPutPreviousValues) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *GlobalAttributesPutPreviousValues) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *GlobalAttributesPutPreviousValues) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


