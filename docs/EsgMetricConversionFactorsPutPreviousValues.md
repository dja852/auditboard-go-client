# EsgMetricConversionFactorsPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**FromUnitId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_metric_unit_options.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_unit_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ToUnitId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_metric_unit_options.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_unit_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ConversionFactor** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewEsgMetricConversionFactorsPutPreviousValues

`func NewEsgMetricConversionFactorsPutPreviousValues() *EsgMetricConversionFactorsPutPreviousValues`

NewEsgMetricConversionFactorsPutPreviousValues instantiates a new EsgMetricConversionFactorsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsgMetricConversionFactorsPutPreviousValuesWithDefaults

`func NewEsgMetricConversionFactorsPutPreviousValuesWithDefaults() *EsgMetricConversionFactorsPutPreviousValues`

NewEsgMetricConversionFactorsPutPreviousValuesWithDefaults instantiates a new EsgMetricConversionFactorsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EsgMetricConversionFactorsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsgMetricConversionFactorsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsgMetricConversionFactorsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EsgMetricConversionFactorsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EsgMetricConversionFactorsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EsgMetricConversionFactorsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EsgMetricConversionFactorsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EsgMetricConversionFactorsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EsgMetricConversionFactorsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EsgMetricConversionFactorsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EsgMetricConversionFactorsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EsgMetricConversionFactorsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EsgMetricConversionFactorsPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EsgMetricConversionFactorsPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EsgMetricConversionFactorsPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EsgMetricConversionFactorsPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetFromUnitId

`func (o *EsgMetricConversionFactorsPutPreviousValues) GetFromUnitId() int32`

GetFromUnitId returns the FromUnitId field if non-nil, zero value otherwise.

### GetFromUnitIdOk

`func (o *EsgMetricConversionFactorsPutPreviousValues) GetFromUnitIdOk() (*int32, bool)`

GetFromUnitIdOk returns a tuple with the FromUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUnitId

`func (o *EsgMetricConversionFactorsPutPreviousValues) SetFromUnitId(v int32)`

SetFromUnitId sets FromUnitId field to given value.

### HasFromUnitId

`func (o *EsgMetricConversionFactorsPutPreviousValues) HasFromUnitId() bool`

HasFromUnitId returns a boolean if a field has been set.

### GetToUnitId

`func (o *EsgMetricConversionFactorsPutPreviousValues) GetToUnitId() int32`

GetToUnitId returns the ToUnitId field if non-nil, zero value otherwise.

### GetToUnitIdOk

`func (o *EsgMetricConversionFactorsPutPreviousValues) GetToUnitIdOk() (*int32, bool)`

GetToUnitIdOk returns a tuple with the ToUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUnitId

`func (o *EsgMetricConversionFactorsPutPreviousValues) SetToUnitId(v int32)`

SetToUnitId sets ToUnitId field to given value.

### HasToUnitId

`func (o *EsgMetricConversionFactorsPutPreviousValues) HasToUnitId() bool`

HasToUnitId returns a boolean if a field has been set.

### GetConversionFactor

`func (o *EsgMetricConversionFactorsPutPreviousValues) GetConversionFactor() float32`

GetConversionFactor returns the ConversionFactor field if non-nil, zero value otherwise.

### GetConversionFactorOk

`func (o *EsgMetricConversionFactorsPutPreviousValues) GetConversionFactorOk() (*float32, bool)`

GetConversionFactorOk returns a tuple with the ConversionFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionFactor

`func (o *EsgMetricConversionFactorsPutPreviousValues) SetConversionFactor(v float32)`

SetConversionFactor sets ConversionFactor field to given value.

### HasConversionFactor

`func (o *EsgMetricConversionFactorsPutPreviousValues) HasConversionFactor() bool`

HasConversionFactor returns a boolean if a field has been set.

### GetType

`func (o *EsgMetricConversionFactorsPutPreviousValues) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EsgMetricConversionFactorsPutPreviousValues) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EsgMetricConversionFactorsPutPreviousValues) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EsgMetricConversionFactorsPutPreviousValues) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


