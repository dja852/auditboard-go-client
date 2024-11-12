# EntityRisksPutEntityRisk

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

### NewEntityRisksPutEntityRisk

`func NewEntityRisksPutEntityRisk(entityId int32, riskId int32, status string, ) *EntityRisksPutEntityRisk`

NewEntityRisksPutEntityRisk instantiates a new EntityRisksPutEntityRisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityRisksPutEntityRiskWithDefaults

`func NewEntityRisksPutEntityRiskWithDefaults() *EntityRisksPutEntityRisk`

NewEntityRisksPutEntityRiskWithDefaults instantiates a new EntityRisksPutEntityRisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntityRisksPutEntityRisk) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityRisksPutEntityRisk) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityRisksPutEntityRisk) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EntityRisksPutEntityRisk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEntityId

`func (o *EntityRisksPutEntityRisk) GetEntityId() int32`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *EntityRisksPutEntityRisk) GetEntityIdOk() (*int32, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *EntityRisksPutEntityRisk) SetEntityId(v int32)`

SetEntityId sets EntityId field to given value.


### GetRiskId

`func (o *EntityRisksPutEntityRisk) GetRiskId() int32`

GetRiskId returns the RiskId field if non-nil, zero value otherwise.

### GetRiskIdOk

`func (o *EntityRisksPutEntityRisk) GetRiskIdOk() (*int32, bool)`

GetRiskIdOk returns a tuple with the RiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskId

`func (o *EntityRisksPutEntityRisk) SetRiskId(v int32)`

SetRiskId sets RiskId field to given value.


### GetStatus

`func (o *EntityRisksPutEntityRisk) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EntityRisksPutEntityRisk) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EntityRisksPutEntityRisk) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *EntityRisksPutEntityRisk) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntityRisksPutEntityRisk) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntityRisksPutEntityRisk) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EntityRisksPutEntityRisk) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EntityRisksPutEntityRisk) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EntityRisksPutEntityRisk) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EntityRisksPutEntityRisk) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EntityRisksPutEntityRisk) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EntityRisksPutEntityRisk) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EntityRisksPutEntityRisk) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EntityRisksPutEntityRisk) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EntityRisksPutEntityRisk) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetUid

`func (o *EntityRisksPutEntityRisk) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *EntityRisksPutEntityRisk) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *EntityRisksPutEntityRisk) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *EntityRisksPutEntityRisk) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetDescription

`func (o *EntityRisksPutEntityRisk) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntityRisksPutEntityRisk) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntityRisksPutEntityRisk) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntityRisksPutEntityRisk) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActivity

`func (o *EntityRisksPutEntityRisk) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *EntityRisksPutEntityRisk) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *EntityRisksPutEntityRisk) SetActivity(v string)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *EntityRisksPutEntityRisk) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetMitigationFactors

`func (o *EntityRisksPutEntityRisk) GetMitigationFactors() string`

GetMitigationFactors returns the MitigationFactors field if non-nil, zero value otherwise.

### GetMitigationFactorsOk

`func (o *EntityRisksPutEntityRisk) GetMitigationFactorsOk() (*string, bool)`

GetMitigationFactorsOk returns a tuple with the MitigationFactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigationFactors

`func (o *EntityRisksPutEntityRisk) SetMitigationFactors(v string)`

SetMitigationFactors sets MitigationFactors field to given value.

### HasMitigationFactors

`func (o *EntityRisksPutEntityRisk) HasMitigationFactors() bool`

HasMitigationFactors returns a boolean if a field has been set.

### GetCustomText1

`func (o *EntityRisksPutEntityRisk) GetCustomText1() string`

GetCustomText1 returns the CustomText1 field if non-nil, zero value otherwise.

### GetCustomText1Ok

`func (o *EntityRisksPutEntityRisk) GetCustomText1Ok() (*string, bool)`

GetCustomText1Ok returns a tuple with the CustomText1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText1

`func (o *EntityRisksPutEntityRisk) SetCustomText1(v string)`

SetCustomText1 sets CustomText1 field to given value.

### HasCustomText1

`func (o *EntityRisksPutEntityRisk) HasCustomText1() bool`

HasCustomText1 returns a boolean if a field has been set.

### GetCustomText2

`func (o *EntityRisksPutEntityRisk) GetCustomText2() string`

GetCustomText2 returns the CustomText2 field if non-nil, zero value otherwise.

### GetCustomText2Ok

`func (o *EntityRisksPutEntityRisk) GetCustomText2Ok() (*string, bool)`

GetCustomText2Ok returns a tuple with the CustomText2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText2

`func (o *EntityRisksPutEntityRisk) SetCustomText2(v string)`

SetCustomText2 sets CustomText2 field to given value.

### HasCustomText2

`func (o *EntityRisksPutEntityRisk) HasCustomText2() bool`

HasCustomText2 returns a boolean if a field has been set.

### GetErCustomSelect1OptionId

`func (o *EntityRisksPutEntityRisk) GetErCustomSelect1OptionId() int32`

GetErCustomSelect1OptionId returns the ErCustomSelect1OptionId field if non-nil, zero value otherwise.

### GetErCustomSelect1OptionIdOk

`func (o *EntityRisksPutEntityRisk) GetErCustomSelect1OptionIdOk() (*int32, bool)`

GetErCustomSelect1OptionIdOk returns a tuple with the ErCustomSelect1OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErCustomSelect1OptionId

`func (o *EntityRisksPutEntityRisk) SetErCustomSelect1OptionId(v int32)`

SetErCustomSelect1OptionId sets ErCustomSelect1OptionId field to given value.

### HasErCustomSelect1OptionId

`func (o *EntityRisksPutEntityRisk) HasErCustomSelect1OptionId() bool`

HasErCustomSelect1OptionId returns a boolean if a field has been set.

### GetErCustomSelect2OptionId

`func (o *EntityRisksPutEntityRisk) GetErCustomSelect2OptionId() int32`

GetErCustomSelect2OptionId returns the ErCustomSelect2OptionId field if non-nil, zero value otherwise.

### GetErCustomSelect2OptionIdOk

`func (o *EntityRisksPutEntityRisk) GetErCustomSelect2OptionIdOk() (*int32, bool)`

GetErCustomSelect2OptionIdOk returns a tuple with the ErCustomSelect2OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErCustomSelect2OptionId

`func (o *EntityRisksPutEntityRisk) SetErCustomSelect2OptionId(v int32)`

SetErCustomSelect2OptionId sets ErCustomSelect2OptionId field to given value.

### HasErCustomSelect2OptionId

`func (o *EntityRisksPutEntityRisk) HasErCustomSelect2OptionId() bool`

HasErCustomSelect2OptionId returns a boolean if a field has been set.

### GetRiskCategoryId

`func (o *EntityRisksPutEntityRisk) GetRiskCategoryId() int32`

GetRiskCategoryId returns the RiskCategoryId field if non-nil, zero value otherwise.

### GetRiskCategoryIdOk

`func (o *EntityRisksPutEntityRisk) GetRiskCategoryIdOk() (*int32, bool)`

GetRiskCategoryIdOk returns a tuple with the RiskCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCategoryId

`func (o *EntityRisksPutEntityRisk) SetRiskCategoryId(v int32)`

SetRiskCategoryId sets RiskCategoryId field to given value.

### HasRiskCategoryId

`func (o *EntityRisksPutEntityRisk) HasRiskCategoryId() bool`

HasRiskCategoryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


