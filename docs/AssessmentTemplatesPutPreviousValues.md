# AssessmentTemplatesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TemplateJson** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**ProjectOptions** | Pointer to **interface{}** |  | [optional] 
**PowerbiMapping** | Pointer to **interface{}** |  | [optional] 
**RiskMapping** | Pointer to **interface{}** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**AuditFrequencyCalculationKey** | Pointer to **string** |  | [optional] 
**PackagingIsStandard** | Pointer to **bool** |  | [optional] [default to false]
**FrequenciesNeedRefresh** | Pointer to **bool** |  | [optional] 
**DecimalPrecision** | Pointer to **int32** |  | [optional] [default to 4]

## Methods

### NewAssessmentTemplatesPutPreviousValues

`func NewAssessmentTemplatesPutPreviousValues() *AssessmentTemplatesPutPreviousValues`

NewAssessmentTemplatesPutPreviousValues instantiates a new AssessmentTemplatesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssessmentTemplatesPutPreviousValuesWithDefaults

`func NewAssessmentTemplatesPutPreviousValuesWithDefaults() *AssessmentTemplatesPutPreviousValues`

NewAssessmentTemplatesPutPreviousValuesWithDefaults instantiates a new AssessmentTemplatesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssessmentTemplatesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssessmentTemplatesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssessmentTemplatesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AssessmentTemplatesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AssessmentTemplatesPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssessmentTemplatesPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssessmentTemplatesPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssessmentTemplatesPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *AssessmentTemplatesPutPreviousValues) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssessmentTemplatesPutPreviousValues) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssessmentTemplatesPutPreviousValues) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssessmentTemplatesPutPreviousValues) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemplateJson

`func (o *AssessmentTemplatesPutPreviousValues) GetTemplateJson() interface{}`

GetTemplateJson returns the TemplateJson field if non-nil, zero value otherwise.

### GetTemplateJsonOk

`func (o *AssessmentTemplatesPutPreviousValues) GetTemplateJsonOk() (*interface{}, bool)`

GetTemplateJsonOk returns a tuple with the TemplateJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateJson

`func (o *AssessmentTemplatesPutPreviousValues) SetTemplateJson(v interface{})`

SetTemplateJson sets TemplateJson field to given value.

### HasTemplateJson

`func (o *AssessmentTemplatesPutPreviousValues) HasTemplateJson() bool`

HasTemplateJson returns a boolean if a field has been set.

### SetTemplateJsonNil

`func (o *AssessmentTemplatesPutPreviousValues) SetTemplateJsonNil(b bool)`

 SetTemplateJsonNil sets the value for TemplateJson to be an explicit nil

### UnsetTemplateJson
`func (o *AssessmentTemplatesPutPreviousValues) UnsetTemplateJson()`

UnsetTemplateJson ensures that no value is present for TemplateJson, not even an explicit nil
### GetCreatedAt

`func (o *AssessmentTemplatesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AssessmentTemplatesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AssessmentTemplatesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AssessmentTemplatesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AssessmentTemplatesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AssessmentTemplatesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AssessmentTemplatesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AssessmentTemplatesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AssessmentTemplatesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AssessmentTemplatesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AssessmentTemplatesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AssessmentTemplatesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetProjectOptions

`func (o *AssessmentTemplatesPutPreviousValues) GetProjectOptions() interface{}`

GetProjectOptions returns the ProjectOptions field if non-nil, zero value otherwise.

### GetProjectOptionsOk

`func (o *AssessmentTemplatesPutPreviousValues) GetProjectOptionsOk() (*interface{}, bool)`

GetProjectOptionsOk returns a tuple with the ProjectOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectOptions

`func (o *AssessmentTemplatesPutPreviousValues) SetProjectOptions(v interface{})`

SetProjectOptions sets ProjectOptions field to given value.

### HasProjectOptions

`func (o *AssessmentTemplatesPutPreviousValues) HasProjectOptions() bool`

HasProjectOptions returns a boolean if a field has been set.

### SetProjectOptionsNil

`func (o *AssessmentTemplatesPutPreviousValues) SetProjectOptionsNil(b bool)`

 SetProjectOptionsNil sets the value for ProjectOptions to be an explicit nil

### UnsetProjectOptions
`func (o *AssessmentTemplatesPutPreviousValues) UnsetProjectOptions()`

UnsetProjectOptions ensures that no value is present for ProjectOptions, not even an explicit nil
### GetPowerbiMapping

`func (o *AssessmentTemplatesPutPreviousValues) GetPowerbiMapping() interface{}`

GetPowerbiMapping returns the PowerbiMapping field if non-nil, zero value otherwise.

### GetPowerbiMappingOk

`func (o *AssessmentTemplatesPutPreviousValues) GetPowerbiMappingOk() (*interface{}, bool)`

GetPowerbiMappingOk returns a tuple with the PowerbiMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerbiMapping

`func (o *AssessmentTemplatesPutPreviousValues) SetPowerbiMapping(v interface{})`

SetPowerbiMapping sets PowerbiMapping field to given value.

### HasPowerbiMapping

`func (o *AssessmentTemplatesPutPreviousValues) HasPowerbiMapping() bool`

HasPowerbiMapping returns a boolean if a field has been set.

### SetPowerbiMappingNil

`func (o *AssessmentTemplatesPutPreviousValues) SetPowerbiMappingNil(b bool)`

 SetPowerbiMappingNil sets the value for PowerbiMapping to be an explicit nil

### UnsetPowerbiMapping
`func (o *AssessmentTemplatesPutPreviousValues) UnsetPowerbiMapping()`

UnsetPowerbiMapping ensures that no value is present for PowerbiMapping, not even an explicit nil
### GetRiskMapping

`func (o *AssessmentTemplatesPutPreviousValues) GetRiskMapping() interface{}`

GetRiskMapping returns the RiskMapping field if non-nil, zero value otherwise.

### GetRiskMappingOk

`func (o *AssessmentTemplatesPutPreviousValues) GetRiskMappingOk() (*interface{}, bool)`

GetRiskMappingOk returns a tuple with the RiskMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskMapping

`func (o *AssessmentTemplatesPutPreviousValues) SetRiskMapping(v interface{})`

SetRiskMapping sets RiskMapping field to given value.

### HasRiskMapping

`func (o *AssessmentTemplatesPutPreviousValues) HasRiskMapping() bool`

HasRiskMapping returns a boolean if a field has been set.

### SetRiskMappingNil

`func (o *AssessmentTemplatesPutPreviousValues) SetRiskMappingNil(b bool)`

 SetRiskMappingNil sets the value for RiskMapping to be an explicit nil

### UnsetRiskMapping
`func (o *AssessmentTemplatesPutPreviousValues) UnsetRiskMapping()`

UnsetRiskMapping ensures that no value is present for RiskMapping, not even an explicit nil
### GetActive

`func (o *AssessmentTemplatesPutPreviousValues) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AssessmentTemplatesPutPreviousValues) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AssessmentTemplatesPutPreviousValues) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AssessmentTemplatesPutPreviousValues) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetType

`func (o *AssessmentTemplatesPutPreviousValues) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssessmentTemplatesPutPreviousValues) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssessmentTemplatesPutPreviousValues) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AssessmentTemplatesPutPreviousValues) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAuditFrequencyCalculationKey

`func (o *AssessmentTemplatesPutPreviousValues) GetAuditFrequencyCalculationKey() string`

GetAuditFrequencyCalculationKey returns the AuditFrequencyCalculationKey field if non-nil, zero value otherwise.

### GetAuditFrequencyCalculationKeyOk

`func (o *AssessmentTemplatesPutPreviousValues) GetAuditFrequencyCalculationKeyOk() (*string, bool)`

GetAuditFrequencyCalculationKeyOk returns a tuple with the AuditFrequencyCalculationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditFrequencyCalculationKey

`func (o *AssessmentTemplatesPutPreviousValues) SetAuditFrequencyCalculationKey(v string)`

SetAuditFrequencyCalculationKey sets AuditFrequencyCalculationKey field to given value.

### HasAuditFrequencyCalculationKey

`func (o *AssessmentTemplatesPutPreviousValues) HasAuditFrequencyCalculationKey() bool`

HasAuditFrequencyCalculationKey returns a boolean if a field has been set.

### GetPackagingIsStandard

`func (o *AssessmentTemplatesPutPreviousValues) GetPackagingIsStandard() bool`

GetPackagingIsStandard returns the PackagingIsStandard field if non-nil, zero value otherwise.

### GetPackagingIsStandardOk

`func (o *AssessmentTemplatesPutPreviousValues) GetPackagingIsStandardOk() (*bool, bool)`

GetPackagingIsStandardOk returns a tuple with the PackagingIsStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagingIsStandard

`func (o *AssessmentTemplatesPutPreviousValues) SetPackagingIsStandard(v bool)`

SetPackagingIsStandard sets PackagingIsStandard field to given value.

### HasPackagingIsStandard

`func (o *AssessmentTemplatesPutPreviousValues) HasPackagingIsStandard() bool`

HasPackagingIsStandard returns a boolean if a field has been set.

### GetFrequenciesNeedRefresh

`func (o *AssessmentTemplatesPutPreviousValues) GetFrequenciesNeedRefresh() bool`

GetFrequenciesNeedRefresh returns the FrequenciesNeedRefresh field if non-nil, zero value otherwise.

### GetFrequenciesNeedRefreshOk

`func (o *AssessmentTemplatesPutPreviousValues) GetFrequenciesNeedRefreshOk() (*bool, bool)`

GetFrequenciesNeedRefreshOk returns a tuple with the FrequenciesNeedRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequenciesNeedRefresh

`func (o *AssessmentTemplatesPutPreviousValues) SetFrequenciesNeedRefresh(v bool)`

SetFrequenciesNeedRefresh sets FrequenciesNeedRefresh field to given value.

### HasFrequenciesNeedRefresh

`func (o *AssessmentTemplatesPutPreviousValues) HasFrequenciesNeedRefresh() bool`

HasFrequenciesNeedRefresh returns a boolean if a field has been set.

### GetDecimalPrecision

`func (o *AssessmentTemplatesPutPreviousValues) GetDecimalPrecision() int32`

GetDecimalPrecision returns the DecimalPrecision field if non-nil, zero value otherwise.

### GetDecimalPrecisionOk

`func (o *AssessmentTemplatesPutPreviousValues) GetDecimalPrecisionOk() (*int32, bool)`

GetDecimalPrecisionOk returns a tuple with the DecimalPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPrecision

`func (o *AssessmentTemplatesPutPreviousValues) SetDecimalPrecision(v int32)`

SetDecimalPrecision sets DecimalPrecision field to given value.

### HasDecimalPrecision

`func (o *AssessmentTemplatesPutPreviousValues) HasDecimalPrecision() bool`

HasDecimalPrecision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


