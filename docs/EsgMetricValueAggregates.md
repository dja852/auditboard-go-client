# EsgMetricValueAggregates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Value** | Pointer to **float32** |  | [optional] 
**EsgMetricUnitOptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_metric_unit_options.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_unit_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**EsgMetricAuditableEntityId** | **int32** | Note: This is a Foreign Key to &#x60;esg_metric_auditable_entities.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_auditable_entities&#39; column&#x3D;&#39;id&#39;/&gt; | 
**AggregateYear** | **string** |  | 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**AggregationMethod** | Pointer to **string** |  | [optional] 

## Methods

### NewEsgMetricValueAggregates

`func NewEsgMetricValueAggregates(esgMetricAuditableEntityId int32, aggregateYear string, ) *EsgMetricValueAggregates`

NewEsgMetricValueAggregates instantiates a new EsgMetricValueAggregates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsgMetricValueAggregatesWithDefaults

`func NewEsgMetricValueAggregatesWithDefaults() *EsgMetricValueAggregates`

NewEsgMetricValueAggregatesWithDefaults instantiates a new EsgMetricValueAggregates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EsgMetricValueAggregates) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsgMetricValueAggregates) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsgMetricValueAggregates) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EsgMetricValueAggregates) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *EsgMetricValueAggregates) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EsgMetricValueAggregates) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EsgMetricValueAggregates) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *EsgMetricValueAggregates) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetEsgMetricUnitOptionId

`func (o *EsgMetricValueAggregates) GetEsgMetricUnitOptionId() int32`

GetEsgMetricUnitOptionId returns the EsgMetricUnitOptionId field if non-nil, zero value otherwise.

### GetEsgMetricUnitOptionIdOk

`func (o *EsgMetricValueAggregates) GetEsgMetricUnitOptionIdOk() (*int32, bool)`

GetEsgMetricUnitOptionIdOk returns a tuple with the EsgMetricUnitOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricUnitOptionId

`func (o *EsgMetricValueAggregates) SetEsgMetricUnitOptionId(v int32)`

SetEsgMetricUnitOptionId sets EsgMetricUnitOptionId field to given value.

### HasEsgMetricUnitOptionId

`func (o *EsgMetricValueAggregates) HasEsgMetricUnitOptionId() bool`

HasEsgMetricUnitOptionId returns a boolean if a field has been set.

### GetEsgMetricAuditableEntityId

`func (o *EsgMetricValueAggregates) GetEsgMetricAuditableEntityId() int32`

GetEsgMetricAuditableEntityId returns the EsgMetricAuditableEntityId field if non-nil, zero value otherwise.

### GetEsgMetricAuditableEntityIdOk

`func (o *EsgMetricValueAggregates) GetEsgMetricAuditableEntityIdOk() (*int32, bool)`

GetEsgMetricAuditableEntityIdOk returns a tuple with the EsgMetricAuditableEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricAuditableEntityId

`func (o *EsgMetricValueAggregates) SetEsgMetricAuditableEntityId(v int32)`

SetEsgMetricAuditableEntityId sets EsgMetricAuditableEntityId field to given value.


### GetAggregateYear

`func (o *EsgMetricValueAggregates) GetAggregateYear() string`

GetAggregateYear returns the AggregateYear field if non-nil, zero value otherwise.

### GetAggregateYearOk

`func (o *EsgMetricValueAggregates) GetAggregateYearOk() (*string, bool)`

GetAggregateYearOk returns a tuple with the AggregateYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateYear

`func (o *EsgMetricValueAggregates) SetAggregateYear(v string)`

SetAggregateYear sets AggregateYear field to given value.


### GetCreatedAt

`func (o *EsgMetricValueAggregates) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EsgMetricValueAggregates) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EsgMetricValueAggregates) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EsgMetricValueAggregates) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EsgMetricValueAggregates) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EsgMetricValueAggregates) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EsgMetricValueAggregates) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EsgMetricValueAggregates) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EsgMetricValueAggregates) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EsgMetricValueAggregates) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EsgMetricValueAggregates) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EsgMetricValueAggregates) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetAggregationMethod

`func (o *EsgMetricValueAggregates) GetAggregationMethod() string`

GetAggregationMethod returns the AggregationMethod field if non-nil, zero value otherwise.

### GetAggregationMethodOk

`func (o *EsgMetricValueAggregates) GetAggregationMethodOk() (*string, bool)`

GetAggregationMethodOk returns a tuple with the AggregationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationMethod

`func (o *EsgMetricValueAggregates) SetAggregationMethod(v string)`

SetAggregationMethod sets AggregationMethod field to given value.

### HasAggregationMethod

`func (o *EsgMetricValueAggregates) HasAggregationMethod() bool`

HasAggregationMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


