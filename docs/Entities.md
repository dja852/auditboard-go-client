# Entities

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

### NewEntities

`func NewEntities(sortOrder int32, entityCode string, name string, regionId int32, isGroup bool, ) *Entities`

NewEntities instantiates a new Entities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitiesWithDefaults

`func NewEntitiesWithDefaults() *Entities`

NewEntitiesWithDefaults instantiates a new Entities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Entities) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Entities) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Entities) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Entities) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Entities) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Entities) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Entities) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Entities) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Entities) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Entities) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Entities) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Entities) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Entities) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Entities) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Entities) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Entities) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetSortOrder

`func (o *Entities) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *Entities) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *Entities) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.


### GetEntityCode

`func (o *Entities) GetEntityCode() string`

GetEntityCode returns the EntityCode field if non-nil, zero value otherwise.

### GetEntityCodeOk

`func (o *Entities) GetEntityCodeOk() (*string, bool)`

GetEntityCodeOk returns a tuple with the EntityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityCode

`func (o *Entities) SetEntityCode(v string)`

SetEntityCode sets EntityCode field to given value.


### GetName

`func (o *Entities) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Entities) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Entities) SetName(v string)`

SetName sets Name field to given value.


### GetRegionId

`func (o *Entities) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *Entities) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *Entities) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.


### GetAuditableEntityId

`func (o *Entities) GetAuditableEntityId() int32`

GetAuditableEntityId returns the AuditableEntityId field if non-nil, zero value otherwise.

### GetAuditableEntityIdOk

`func (o *Entities) GetAuditableEntityIdOk() (*int32, bool)`

GetAuditableEntityIdOk returns a tuple with the AuditableEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditableEntityId

`func (o *Entities) SetAuditableEntityId(v int32)`

SetAuditableEntityId sets AuditableEntityId field to given value.

### HasAuditableEntityId

`func (o *Entities) HasAuditableEntityId() bool`

HasAuditableEntityId returns a boolean if a field has been set.

### GetIsGroup

`func (o *Entities) GetIsGroup() bool`

GetIsGroup returns the IsGroup field if non-nil, zero value otherwise.

### GetIsGroupOk

`func (o *Entities) GetIsGroupOk() (*bool, bool)`

GetIsGroupOk returns a tuple with the IsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGroup

`func (o *Entities) SetIsGroup(v bool)`

SetIsGroup sets IsGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


