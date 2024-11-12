# EntityRisks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**EntityId** | **int32** | Note: This is a Foreign Key to &#x60;entities.id&#x60;.&lt;fk table&#x3D;&#39;entities&#39; column&#x3D;&#39;id&#39;/&gt; | 
**RiskId** | **int32** | Note: This is a Foreign Key to &#x60;risks.id&#x60;.&lt;fk table&#x3D;&#39;risks&#39; column&#x3D;&#39;id&#39;/&gt; | 
**Status** | **string** |  | 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Activity** | Pointer to **string** |  | [optional] 
**MitigationFactors** | Pointer to **string** |  | [optional] 
**CustomText1** | Pointer to **string** |  | [optional] 
**CustomText2** | Pointer to **string** |  | [optional] 
**ErCustomSelect1OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;er_custom_select1_options.id&#x60;.&lt;fk table&#x3D;&#39;er_custom_select1_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ErCustomSelect2OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;er_custom_select2_options.id&#x60;.&lt;fk table&#x3D;&#39;er_custom_select2_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCategoryId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_categories.id&#x60;.&lt;fk table&#x3D;&#39;risk_categories&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 

## Methods

### NewEntityRisks

`func NewEntityRisks(entityId int32, riskId int32, status string, ) *EntityRisks`

NewEntityRisks instantiates a new EntityRisks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityRisksWithDefaults

`func NewEntityRisksWithDefaults() *EntityRisks`

NewEntityRisksWithDefaults instantiates a new EntityRisks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntityRisks) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityRisks) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityRisks) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EntityRisks) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEntityId

`func (o *EntityRisks) GetEntityId() int32`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *EntityRisks) GetEntityIdOk() (*int32, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *EntityRisks) SetEntityId(v int32)`

SetEntityId sets EntityId field to given value.


### GetRiskId

`func (o *EntityRisks) GetRiskId() int32`

GetRiskId returns the RiskId field if non-nil, zero value otherwise.

### GetRiskIdOk

`func (o *EntityRisks) GetRiskIdOk() (*int32, bool)`

GetRiskIdOk returns a tuple with the RiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskId

`func (o *EntityRisks) SetRiskId(v int32)`

SetRiskId sets RiskId field to given value.


### GetStatus

`func (o *EntityRisks) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EntityRisks) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EntityRisks) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *EntityRisks) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntityRisks) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntityRisks) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EntityRisks) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EntityRisks) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EntityRisks) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EntityRisks) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EntityRisks) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EntityRisks) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EntityRisks) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EntityRisks) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EntityRisks) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetUid

`func (o *EntityRisks) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *EntityRisks) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *EntityRisks) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *EntityRisks) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetDescription

`func (o *EntityRisks) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntityRisks) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntityRisks) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntityRisks) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActivity

`func (o *EntityRisks) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *EntityRisks) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *EntityRisks) SetActivity(v string)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *EntityRisks) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetMitigationFactors

`func (o *EntityRisks) GetMitigationFactors() string`

GetMitigationFactors returns the MitigationFactors field if non-nil, zero value otherwise.

### GetMitigationFactorsOk

`func (o *EntityRisks) GetMitigationFactorsOk() (*string, bool)`

GetMitigationFactorsOk returns a tuple with the MitigationFactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigationFactors

`func (o *EntityRisks) SetMitigationFactors(v string)`

SetMitigationFactors sets MitigationFactors field to given value.

### HasMitigationFactors

`func (o *EntityRisks) HasMitigationFactors() bool`

HasMitigationFactors returns a boolean if a field has been set.

### GetCustomText1

`func (o *EntityRisks) GetCustomText1() string`

GetCustomText1 returns the CustomText1 field if non-nil, zero value otherwise.

### GetCustomText1Ok

`func (o *EntityRisks) GetCustomText1Ok() (*string, bool)`

GetCustomText1Ok returns a tuple with the CustomText1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText1

`func (o *EntityRisks) SetCustomText1(v string)`

SetCustomText1 sets CustomText1 field to given value.

### HasCustomText1

`func (o *EntityRisks) HasCustomText1() bool`

HasCustomText1 returns a boolean if a field has been set.

### GetCustomText2

`func (o *EntityRisks) GetCustomText2() string`

GetCustomText2 returns the CustomText2 field if non-nil, zero value otherwise.

### GetCustomText2Ok

`func (o *EntityRisks) GetCustomText2Ok() (*string, bool)`

GetCustomText2Ok returns a tuple with the CustomText2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText2

`func (o *EntityRisks) SetCustomText2(v string)`

SetCustomText2 sets CustomText2 field to given value.

### HasCustomText2

`func (o *EntityRisks) HasCustomText2() bool`

HasCustomText2 returns a boolean if a field has been set.

### GetErCustomSelect1OptionId

`func (o *EntityRisks) GetErCustomSelect1OptionId() int32`

GetErCustomSelect1OptionId returns the ErCustomSelect1OptionId field if non-nil, zero value otherwise.

### GetErCustomSelect1OptionIdOk

`func (o *EntityRisks) GetErCustomSelect1OptionIdOk() (*int32, bool)`

GetErCustomSelect1OptionIdOk returns a tuple with the ErCustomSelect1OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErCustomSelect1OptionId

`func (o *EntityRisks) SetErCustomSelect1OptionId(v int32)`

SetErCustomSelect1OptionId sets ErCustomSelect1OptionId field to given value.

### HasErCustomSelect1OptionId

`func (o *EntityRisks) HasErCustomSelect1OptionId() bool`

HasErCustomSelect1OptionId returns a boolean if a field has been set.

### GetErCustomSelect2OptionId

`func (o *EntityRisks) GetErCustomSelect2OptionId() int32`

GetErCustomSelect2OptionId returns the ErCustomSelect2OptionId field if non-nil, zero value otherwise.

### GetErCustomSelect2OptionIdOk

`func (o *EntityRisks) GetErCustomSelect2OptionIdOk() (*int32, bool)`

GetErCustomSelect2OptionIdOk returns a tuple with the ErCustomSelect2OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErCustomSelect2OptionId

`func (o *EntityRisks) SetErCustomSelect2OptionId(v int32)`

SetErCustomSelect2OptionId sets ErCustomSelect2OptionId field to given value.

### HasErCustomSelect2OptionId

`func (o *EntityRisks) HasErCustomSelect2OptionId() bool`

HasErCustomSelect2OptionId returns a boolean if a field has been set.

### GetRiskCategoryId

`func (o *EntityRisks) GetRiskCategoryId() int32`

GetRiskCategoryId returns the RiskCategoryId field if non-nil, zero value otherwise.

### GetRiskCategoryIdOk

`func (o *EntityRisks) GetRiskCategoryIdOk() (*int32, bool)`

GetRiskCategoryIdOk returns a tuple with the RiskCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCategoryId

`func (o *EntityRisks) SetRiskCategoryId(v int32)`

SetRiskCategoryId sets RiskCategoryId field to given value.

### HasRiskCategoryId

`func (o *EntityRisks) HasRiskCategoryId() bool`

HasRiskCategoryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


