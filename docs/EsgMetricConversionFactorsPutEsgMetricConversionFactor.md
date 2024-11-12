# EsgMetricConversionFactorsPutEsgMetricConversionFactor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**FromUnitId** | **int32** | Note: This is a Foreign Key to &#x60;esg_metric_unit_options.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_unit_options&#39; column&#x3D;&#39;id&#39;/&gt; | 
**ToUnitId** | **int32** | Note: This is a Foreign Key to &#x60;esg_metric_unit_options.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_unit_options&#39; column&#x3D;&#39;id&#39;/&gt; | 
**ConversionFactor** | **float32** |  | 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewEsgMetricConversionFactorsPutEsgMetricConversionFactor

`func NewEsgMetricConversionFactorsPutEsgMetricConversionFactor(fromUnitId int32, toUnitId int32, conversionFactor float32, ) *EsgMetricConversionFactorsPutEsgMetricConversionFactor`

NewEsgMetricConversionFactorsPutEsgMetricConversionFactor instantiates a new EsgMetricConversionFactorsPutEsgMetricConversionFactor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsgMetricConversionFactorsPutEsgMetricConversionFactorWithDefaults

`func NewEsgMetricConversionFactorsPutEsgMetricConversionFactorWithDefaults() *EsgMetricConversionFactorsPutEsgMetricConversionFactor`

NewEsgMetricConversionFactorsPutEsgMetricConversionFactorWithDefaults instantiates a new EsgMetricConversionFactorsPutEsgMetricConversionFactor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetFromUnitId

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) GetFromUnitId() int32`

GetFromUnitId returns the FromUnitId field if non-nil, zero value otherwise.

### GetFromUnitIdOk

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) GetFromUnitIdOk() (*int32, bool)`

GetFromUnitIdOk returns a tuple with the FromUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUnitId

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) SetFromUnitId(v int32)`

SetFromUnitId sets FromUnitId field to given value.


### GetToUnitId

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) GetToUnitId() int32`

GetToUnitId returns the ToUnitId field if non-nil, zero value otherwise.

### GetToUnitIdOk

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) GetToUnitIdOk() (*int32, bool)`

GetToUnitIdOk returns a tuple with the ToUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUnitId

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) SetToUnitId(v int32)`

SetToUnitId sets ToUnitId field to given value.


### GetConversionFactor

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) GetConversionFactor() float32`

GetConversionFactor returns the ConversionFactor field if non-nil, zero value otherwise.

### GetConversionFactorOk

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) GetConversionFactorOk() (*float32, bool)`

GetConversionFactorOk returns a tuple with the ConversionFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionFactor

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) SetConversionFactor(v float32)`

SetConversionFactor sets ConversionFactor field to given value.


### GetType

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EsgMetricConversionFactorsPutEsgMetricConversionFactor) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


