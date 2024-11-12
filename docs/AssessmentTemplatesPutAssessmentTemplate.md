# AssessmentTemplatesPutAssessmentTemplate

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
**PackagingIsStandard** | **bool** |  | [default to false]
**FrequenciesNeedRefresh** | Pointer to **bool** |  | [optional] 
**DecimalPrecision** | **int32** |  | [default to 4]

## Methods

### NewAssessmentTemplatesPutAssessmentTemplate

`func NewAssessmentTemplatesPutAssessmentTemplate(packagingIsStandard bool, decimalPrecision int32, ) *AssessmentTemplatesPutAssessmentTemplate`

NewAssessmentTemplatesPutAssessmentTemplate instantiates a new AssessmentTemplatesPutAssessmentTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssessmentTemplatesPutAssessmentTemplateWithDefaults

`func NewAssessmentTemplatesPutAssessmentTemplateWithDefaults() *AssessmentTemplatesPutAssessmentTemplate`

NewAssessmentTemplatesPutAssessmentTemplateWithDefaults instantiates a new AssessmentTemplatesPutAssessmentTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssessmentTemplatesPutAssessmentTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AssessmentTemplatesPutAssessmentTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssessmentTemplatesPutAssessmentTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssessmentTemplatesPutAssessmentTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssessmentTemplatesPutAssessmentTemplate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssessmentTemplatesPutAssessmentTemplate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemplateJson

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetTemplateJson() interface{}`

GetTemplateJson returns the TemplateJson field if non-nil, zero value otherwise.

### GetTemplateJsonOk

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetTemplateJsonOk() (*interface{}, bool)`

GetTemplateJsonOk returns a tuple with the TemplateJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateJson

`func (o *AssessmentTemplatesPutAssessmentTemplate) SetTemplateJson(v interface{})`

SetTemplateJson sets TemplateJson field to given value.

### HasTemplateJson

`func (o *AssessmentTemplatesPutAssessmentTemplate) HasTemplateJson() bool`

HasTemplateJson returns a boolean if a field has been set.

### SetTemplateJsonNil

`func (o *AssessmentTemplatesPutAssessmentTemplate) SetTemplateJsonNil(b bool)`

 SetTemplateJsonNil sets the value for TemplateJson to be an explicit nil

### UnsetTemplateJson
`func (o *AssessmentTemplatesPutAssessmentTemplate) UnsetTemplateJson()`

UnsetTemplateJson ensures that no value is present for TemplateJson, not even an explicit nil
### GetCreatedAt

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AssessmentTemplatesPutAssessmentTemplate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AssessmentTemplatesPutAssessmentTemplate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AssessmentTemplatesPutAssessmentTemplate) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AssessmentTemplatesPutAssessmentTemplate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AssessmentTemplatesPutAssessmentTemplate) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AssessmentTemplatesPutAssessmentTemplate) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetProjectOptions

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetProjectOptions() interface{}`

GetProjectOptions returns the ProjectOptions field if non-nil, zero value otherwise.

### GetProjectOptionsOk

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetProjectOptionsOk() (*interface{}, bool)`

GetProjectOptionsOk returns a tuple with the ProjectOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectOptions

`func (o *AssessmentTemplatesPutAssessmentTemplate) SetProjectOptions(v interface{})`

SetProjectOptions sets ProjectOptions field to given value.

### HasProjectOptions

`func (o *AssessmentTemplatesPutAssessmentTemplate) HasProjectOptions() bool`

HasProjectOptions returns a boolean if a field has been set.

### SetProjectOptionsNil

`func (o *AssessmentTemplatesPutAssessmentTemplate) SetProjectOptionsNil(b bool)`

 SetProjectOptionsNil sets the value for ProjectOptions to be an explicit nil

### UnsetProjectOptions
`func (o *AssessmentTemplatesPutAssessmentTemplate) UnsetProjectOptions()`

UnsetProjectOptions ensures that no value is present for ProjectOptions, not even an explicit nil
### GetPowerbiMapping

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetPowerbiMapping() interface{}`

GetPowerbiMapping returns the PowerbiMapping field if non-nil, zero value otherwise.

### GetPowerbiMappingOk

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetPowerbiMappingOk() (*interface{}, bool)`

GetPowerbiMappingOk returns a tuple with the PowerbiMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerbiMapping

`func (o *AssessmentTemplatesPutAssessmentTemplate) SetPowerbiMapping(v interface{})`

SetPowerbiMapping sets PowerbiMapping field to given value.

### HasPowerbiMapping

`func (o *AssessmentTemplatesPutAssessmentTemplate) HasPowerbiMapping() bool`

HasPowerbiMapping returns a boolean if a field has been set.

### SetPowerbiMappingNil

`func (o *AssessmentTemplatesPutAssessmentTemplate) SetPowerbiMappingNil(b bool)`

 SetPowerbiMappingNil sets the value for PowerbiMapping to be an explicit nil

### UnsetPowerbiMapping
`func (o *AssessmentTemplatesPutAssessmentTemplate) UnsetPowerbiMapping()`

UnsetPowerbiMapping ensures that no value is present for PowerbiMapping, not even an explicit nil
### GetRiskMapping

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetRiskMapping() interface{}`

GetRiskMapping returns the RiskMapping field if non-nil, zero value otherwise.

### GetRiskMappingOk

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetRiskMappingOk() (*interface{}, bool)`

GetRiskMappingOk returns a tuple with the RiskMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskMapping

`func (o *AssessmentTemplatesPutAssessmentTemplate) SetRiskMapping(v interface{})`

SetRiskMapping sets RiskMapping field to given value.

### HasRiskMapping

`func (o *AssessmentTemplatesPutAssessmentTemplate) HasRiskMapping() bool`

HasRiskMapping returns a boolean if a field has been set.

### SetRiskMappingNil

`func (o *AssessmentTemplatesPutAssessmentTemplate) SetRiskMappingNil(b bool)`

 SetRiskMappingNil sets the value for RiskMapping to be an explicit nil

### UnsetRiskMapping
`func (o *AssessmentTemplatesPutAssessmentTemplate) UnsetRiskMapping()`

UnsetRiskMapping ensures that no value is present for RiskMapping, not even an explicit nil
### GetActive

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AssessmentTemplatesPutAssessmentTemplate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AssessmentTemplatesPutAssessmentTemplate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetType

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssessmentTemplatesPutAssessmentTemplate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AssessmentTemplatesPutAssessmentTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAuditFrequencyCalculationKey

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetAuditFrequencyCalculationKey() string`

GetAuditFrequencyCalculationKey returns the AuditFrequencyCalculationKey field if non-nil, zero value otherwise.

### GetAuditFrequencyCalculationKeyOk

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetAuditFrequencyCalculationKeyOk() (*string, bool)`

GetAuditFrequencyCalculationKeyOk returns a tuple with the AuditFrequencyCalculationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditFrequencyCalculationKey

`func (o *AssessmentTemplatesPutAssessmentTemplate) SetAuditFrequencyCalculationKey(v string)`

SetAuditFrequencyCalculationKey sets AuditFrequencyCalculationKey field to given value.

### HasAuditFrequencyCalculationKey

`func (o *AssessmentTemplatesPutAssessmentTemplate) HasAuditFrequencyCalculationKey() bool`

HasAuditFrequencyCalculationKey returns a boolean if a field has been set.

### GetPackagingIsStandard

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetPackagingIsStandard() bool`

GetPackagingIsStandard returns the PackagingIsStandard field if non-nil, zero value otherwise.

### GetPackagingIsStandardOk

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetPackagingIsStandardOk() (*bool, bool)`

GetPackagingIsStandardOk returns a tuple with the PackagingIsStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagingIsStandard

`func (o *AssessmentTemplatesPutAssessmentTemplate) SetPackagingIsStandard(v bool)`

SetPackagingIsStandard sets PackagingIsStandard field to given value.


### GetFrequenciesNeedRefresh

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetFrequenciesNeedRefresh() bool`

GetFrequenciesNeedRefresh returns the FrequenciesNeedRefresh field if non-nil, zero value otherwise.

### GetFrequenciesNeedRefreshOk

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetFrequenciesNeedRefreshOk() (*bool, bool)`

GetFrequenciesNeedRefreshOk returns a tuple with the FrequenciesNeedRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequenciesNeedRefresh

`func (o *AssessmentTemplatesPutAssessmentTemplate) SetFrequenciesNeedRefresh(v bool)`

SetFrequenciesNeedRefresh sets FrequenciesNeedRefresh field to given value.

### HasFrequenciesNeedRefresh

`func (o *AssessmentTemplatesPutAssessmentTemplate) HasFrequenciesNeedRefresh() bool`

HasFrequenciesNeedRefresh returns a boolean if a field has been set.

### GetDecimalPrecision

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetDecimalPrecision() int32`

GetDecimalPrecision returns the DecimalPrecision field if non-nil, zero value otherwise.

### GetDecimalPrecisionOk

`func (o *AssessmentTemplatesPutAssessmentTemplate) GetDecimalPrecisionOk() (*int32, bool)`

GetDecimalPrecisionOk returns a tuple with the DecimalPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPrecision

`func (o *AssessmentTemplatesPutAssessmentTemplate) SetDecimalPrecision(v int32)`

SetDecimalPrecision sets DecimalPrecision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


