# EntitiesPutEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**SortOrder** | **int32** |  | 
**EntityCode** | **string** |  | 
**Name** | **string** |  | 
**RegionId** | **int32** | Note: This is a Foreign Key to &#x60;regions.id&#x60;.&lt;fk table&#x3D;&#39;regions&#39; column&#x3D;&#39;id&#39;/&gt; | 
**AuditableEntityId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;auditable_entities.id&#x60;.&lt;fk table&#x3D;&#39;auditable_entities&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IsGroup** | **bool** |  | [default to false]

## Methods

### NewEntitiesPutEntity

`func NewEntitiesPutEntity(sortOrder int32, entityCode string, name string, regionId int32, isGroup bool, ) *EntitiesPutEntity`

NewEntitiesPutEntity instantiates a new EntitiesPutEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitiesPutEntityWithDefaults

`func NewEntitiesPutEntityWithDefaults() *EntitiesPutEntity`

NewEntitiesPutEntityWithDefaults instantiates a new EntitiesPutEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntitiesPutEntity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitiesPutEntity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitiesPutEntity) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EntitiesPutEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EntitiesPutEntity) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntitiesPutEntity) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntitiesPutEntity) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EntitiesPutEntity) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EntitiesPutEntity) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EntitiesPutEntity) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EntitiesPutEntity) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EntitiesPutEntity) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EntitiesPutEntity) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EntitiesPutEntity) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EntitiesPutEntity) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EntitiesPutEntity) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetSortOrder

`func (o *EntitiesPutEntity) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *EntitiesPutEntity) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *EntitiesPutEntity) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.


### GetEntityCode

`func (o *EntitiesPutEntity) GetEntityCode() string`

GetEntityCode returns the EntityCode field if non-nil, zero value otherwise.

### GetEntityCodeOk

`func (o *EntitiesPutEntity) GetEntityCodeOk() (*string, bool)`

GetEntityCodeOk returns a tuple with the EntityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityCode

`func (o *EntitiesPutEntity) SetEntityCode(v string)`

SetEntityCode sets EntityCode field to given value.


### GetName

`func (o *EntitiesPutEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitiesPutEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitiesPutEntity) SetName(v string)`

SetName sets Name field to given value.


### GetRegionId

`func (o *EntitiesPutEntity) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *EntitiesPutEntity) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *EntitiesPutEntity) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.


### GetAuditableEntityId

`func (o *EntitiesPutEntity) GetAuditableEntityId() int32`

GetAuditableEntityId returns the AuditableEntityId field if non-nil, zero value otherwise.

### GetAuditableEntityIdOk

`func (o *EntitiesPutEntity) GetAuditableEntityIdOk() (*int32, bool)`

GetAuditableEntityIdOk returns a tuple with the AuditableEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditableEntityId

`func (o *EntitiesPutEntity) SetAuditableEntityId(v int32)`

SetAuditableEntityId sets AuditableEntityId field to given value.

### HasAuditableEntityId

`func (o *EntitiesPutEntity) HasAuditableEntityId() bool`

HasAuditableEntityId returns a boolean if a field has been set.

### GetIsGroup

`func (o *EntitiesPutEntity) GetIsGroup() bool`

GetIsGroup returns the IsGroup field if non-nil, zero value otherwise.

### GetIsGroupOk

`func (o *EntitiesPutEntity) GetIsGroupOk() (*bool, bool)`

GetIsGroupOk returns a tuple with the IsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGroup

`func (o *EntitiesPutEntity) SetIsGroup(v bool)`

SetIsGroup sets IsGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


