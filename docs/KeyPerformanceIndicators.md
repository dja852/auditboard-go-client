# KeyPerformanceIndicators

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**CompletionPercentage** | Pointer to **string** |  | [optional] 
**SumValues** | Pointer to **bool** |  | [optional] 
**OwnerUserId** | Pointer to **int64** |  | [optional] 
**ExecutiveOwnerUserId** | Pointer to **int64** |  | [optional] 
**FieldData** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewKeyPerformanceIndicators

`func NewKeyPerformanceIndicators() *KeyPerformanceIndicators`

NewKeyPerformanceIndicators instantiates a new KeyPerformanceIndicators object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyPerformanceIndicatorsWithDefaults

`func NewKeyPerformanceIndicatorsWithDefaults() *KeyPerformanceIndicators`

NewKeyPerformanceIndicatorsWithDefaults instantiates a new KeyPerformanceIndicators object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyPerformanceIndicators) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyPerformanceIndicators) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyPerformanceIndicators) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KeyPerformanceIndicators) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *KeyPerformanceIndicators) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *KeyPerformanceIndicators) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *KeyPerformanceIndicators) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *KeyPerformanceIndicators) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetName

`func (o *KeyPerformanceIndicators) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyPerformanceIndicators) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyPerformanceIndicators) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyPerformanceIndicators) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *KeyPerformanceIndicators) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KeyPerformanceIndicators) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KeyPerformanceIndicators) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KeyPerformanceIndicators) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTarget

`func (o *KeyPerformanceIndicators) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *KeyPerformanceIndicators) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *KeyPerformanceIndicators) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *KeyPerformanceIndicators) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetValue

`func (o *KeyPerformanceIndicators) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KeyPerformanceIndicators) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KeyPerformanceIndicators) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *KeyPerformanceIndicators) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCompletionPercentage

`func (o *KeyPerformanceIndicators) GetCompletionPercentage() string`

GetCompletionPercentage returns the CompletionPercentage field if non-nil, zero value otherwise.

### GetCompletionPercentageOk

`func (o *KeyPerformanceIndicators) GetCompletionPercentageOk() (*string, bool)`

GetCompletionPercentageOk returns a tuple with the CompletionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionPercentage

`func (o *KeyPerformanceIndicators) SetCompletionPercentage(v string)`

SetCompletionPercentage sets CompletionPercentage field to given value.

### HasCompletionPercentage

`func (o *KeyPerformanceIndicators) HasCompletionPercentage() bool`

HasCompletionPercentage returns a boolean if a field has been set.

### GetSumValues

`func (o *KeyPerformanceIndicators) GetSumValues() bool`

GetSumValues returns the SumValues field if non-nil, zero value otherwise.

### GetSumValuesOk

`func (o *KeyPerformanceIndicators) GetSumValuesOk() (*bool, bool)`

GetSumValuesOk returns a tuple with the SumValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumValues

`func (o *KeyPerformanceIndicators) SetSumValues(v bool)`

SetSumValues sets SumValues field to given value.

### HasSumValues

`func (o *KeyPerformanceIndicators) HasSumValues() bool`

HasSumValues returns a boolean if a field has been set.

### GetOwnerUserId

`func (o *KeyPerformanceIndicators) GetOwnerUserId() int64`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *KeyPerformanceIndicators) GetOwnerUserIdOk() (*int64, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *KeyPerformanceIndicators) SetOwnerUserId(v int64)`

SetOwnerUserId sets OwnerUserId field to given value.

### HasOwnerUserId

`func (o *KeyPerformanceIndicators) HasOwnerUserId() bool`

HasOwnerUserId returns a boolean if a field has been set.

### GetExecutiveOwnerUserId

`func (o *KeyPerformanceIndicators) GetExecutiveOwnerUserId() int64`

GetExecutiveOwnerUserId returns the ExecutiveOwnerUserId field if non-nil, zero value otherwise.

### GetExecutiveOwnerUserIdOk

`func (o *KeyPerformanceIndicators) GetExecutiveOwnerUserIdOk() (*int64, bool)`

GetExecutiveOwnerUserIdOk returns a tuple with the ExecutiveOwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutiveOwnerUserId

`func (o *KeyPerformanceIndicators) SetExecutiveOwnerUserId(v int64)`

SetExecutiveOwnerUserId sets ExecutiveOwnerUserId field to given value.

### HasExecutiveOwnerUserId

`func (o *KeyPerformanceIndicators) HasExecutiveOwnerUserId() bool`

HasExecutiveOwnerUserId returns a boolean if a field has been set.

### GetFieldData

`func (o *KeyPerformanceIndicators) GetFieldData() map[string]interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *KeyPerformanceIndicators) GetFieldDataOk() (*map[string]interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *KeyPerformanceIndicators) SetFieldData(v map[string]interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *KeyPerformanceIndicators) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KeyPerformanceIndicators) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KeyPerformanceIndicators) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KeyPerformanceIndicators) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KeyPerformanceIndicators) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KeyPerformanceIndicators) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KeyPerformanceIndicators) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KeyPerformanceIndicators) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KeyPerformanceIndicators) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *KeyPerformanceIndicators) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *KeyPerformanceIndicators) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *KeyPerformanceIndicators) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *KeyPerformanceIndicators) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetPermissions

`func (o *KeyPerformanceIndicators) GetPermissions() map[string]interface{}`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *KeyPerformanceIndicators) GetPermissionsOk() (*map[string]interface{}, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *KeyPerformanceIndicators) SetPermissions(v map[string]interface{})`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *KeyPerformanceIndicators) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


