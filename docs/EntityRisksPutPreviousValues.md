# EntityRisksPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**EntityId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;entities.id&#x60;.&lt;fk table&#x3D;&#39;entities&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risks.id&#x60;.&lt;fk table&#x3D;&#39;risks&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Status** | Pointer to **string** |  | [optional] 
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

### NewEntityRisksPutPreviousValues

`func NewEntityRisksPutPreviousValues() *EntityRisksPutPreviousValues`

NewEntityRisksPutPreviousValues instantiates a new EntityRisksPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityRisksPutPreviousValuesWithDefaults

`func NewEntityRisksPutPreviousValuesWithDefaults() *EntityRisksPutPreviousValues`

NewEntityRisksPutPreviousValuesWithDefaults instantiates a new EntityRisksPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntityRisksPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityRisksPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityRisksPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EntityRisksPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEntityId

`func (o *EntityRisksPutPreviousValues) GetEntityId() int32`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *EntityRisksPutPreviousValues) GetEntityIdOk() (*int32, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *EntityRisksPutPreviousValues) SetEntityId(v int32)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *EntityRisksPutPreviousValues) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetRiskId

`func (o *EntityRisksPutPreviousValues) GetRiskId() int32`

GetRiskId returns the RiskId field if non-nil, zero value otherwise.

### GetRiskIdOk

`func (o *EntityRisksPutPreviousValues) GetRiskIdOk() (*int32, bool)`

GetRiskIdOk returns a tuple with the RiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskId

`func (o *EntityRisksPutPreviousValues) SetRiskId(v int32)`

SetRiskId sets RiskId field to given value.

### HasRiskId

`func (o *EntityRisksPutPreviousValues) HasRiskId() bool`

HasRiskId returns a boolean if a field has been set.

### GetStatus

`func (o *EntityRisksPutPreviousValues) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EntityRisksPutPreviousValues) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EntityRisksPutPreviousValues) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EntityRisksPutPreviousValues) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EntityRisksPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntityRisksPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntityRisksPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EntityRisksPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EntityRisksPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EntityRisksPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EntityRisksPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EntityRisksPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EntityRisksPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EntityRisksPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EntityRisksPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EntityRisksPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetUid

`func (o *EntityRisksPutPreviousValues) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *EntityRisksPutPreviousValues) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *EntityRisksPutPreviousValues) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *EntityRisksPutPreviousValues) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetDescription

`func (o *EntityRisksPutPreviousValues) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntityRisksPutPreviousValues) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntityRisksPutPreviousValues) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntityRisksPutPreviousValues) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActivity

`func (o *EntityRisksPutPreviousValues) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *EntityRisksPutPreviousValues) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *EntityRisksPutPreviousValues) SetActivity(v string)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *EntityRisksPutPreviousValues) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetMitigationFactors

`func (o *EntityRisksPutPreviousValues) GetMitigationFactors() string`

GetMitigationFactors returns the MitigationFactors field if non-nil, zero value otherwise.

### GetMitigationFactorsOk

`func (o *EntityRisksPutPreviousValues) GetMitigationFactorsOk() (*string, bool)`

GetMitigationFactorsOk returns a tuple with the MitigationFactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigationFactors

`func (o *EntityRisksPutPreviousValues) SetMitigationFactors(v string)`

SetMitigationFactors sets MitigationFactors field to given value.

### HasMitigationFactors

`func (o *EntityRisksPutPreviousValues) HasMitigationFactors() bool`

HasMitigationFactors returns a boolean if a field has been set.

### GetCustomText1

`func (o *EntityRisksPutPreviousValues) GetCustomText1() string`

GetCustomText1 returns the CustomText1 field if non-nil, zero value otherwise.

### GetCustomText1Ok

`func (o *EntityRisksPutPreviousValues) GetCustomText1Ok() (*string, bool)`

GetCustomText1Ok returns a tuple with the CustomText1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText1

`func (o *EntityRisksPutPreviousValues) SetCustomText1(v string)`

SetCustomText1 sets CustomText1 field to given value.

### HasCustomText1

`func (o *EntityRisksPutPreviousValues) HasCustomText1() bool`

HasCustomText1 returns a boolean if a field has been set.

### GetCustomText2

`func (o *EntityRisksPutPreviousValues) GetCustomText2() string`

GetCustomText2 returns the CustomText2 field if non-nil, zero value otherwise.

### GetCustomText2Ok

`func (o *EntityRisksPutPreviousValues) GetCustomText2Ok() (*string, bool)`

GetCustomText2Ok returns a tuple with the CustomText2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText2

`func (o *EntityRisksPutPreviousValues) SetCustomText2(v string)`

SetCustomText2 sets CustomText2 field to given value.

### HasCustomText2

`func (o *EntityRisksPutPreviousValues) HasCustomText2() bool`

HasCustomText2 returns a boolean if a field has been set.

### GetErCustomSelect1OptionId

`func (o *EntityRisksPutPreviousValues) GetErCustomSelect1OptionId() int32`

GetErCustomSelect1OptionId returns the ErCustomSelect1OptionId field if non-nil, zero value otherwise.

### GetErCustomSelect1OptionIdOk

`func (o *EntityRisksPutPreviousValues) GetErCustomSelect1OptionIdOk() (*int32, bool)`

GetErCustomSelect1OptionIdOk returns a tuple with the ErCustomSelect1OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErCustomSelect1OptionId

`func (o *EntityRisksPutPreviousValues) SetErCustomSelect1OptionId(v int32)`

SetErCustomSelect1OptionId sets ErCustomSelect1OptionId field to given value.

### HasErCustomSelect1OptionId

`func (o *EntityRisksPutPreviousValues) HasErCustomSelect1OptionId() bool`

HasErCustomSelect1OptionId returns a boolean if a field has been set.

### GetErCustomSelect2OptionId

`func (o *EntityRisksPutPreviousValues) GetErCustomSelect2OptionId() int32`

GetErCustomSelect2OptionId returns the ErCustomSelect2OptionId field if non-nil, zero value otherwise.

### GetErCustomSelect2OptionIdOk

`func (o *EntityRisksPutPreviousValues) GetErCustomSelect2OptionIdOk() (*int32, bool)`

GetErCustomSelect2OptionIdOk returns a tuple with the ErCustomSelect2OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErCustomSelect2OptionId

`func (o *EntityRisksPutPreviousValues) SetErCustomSelect2OptionId(v int32)`

SetErCustomSelect2OptionId sets ErCustomSelect2OptionId field to given value.

### HasErCustomSelect2OptionId

`func (o *EntityRisksPutPreviousValues) HasErCustomSelect2OptionId() bool`

HasErCustomSelect2OptionId returns a boolean if a field has been set.

### GetRiskCategoryId

`func (o *EntityRisksPutPreviousValues) GetRiskCategoryId() int32`

GetRiskCategoryId returns the RiskCategoryId field if non-nil, zero value otherwise.

### GetRiskCategoryIdOk

`func (o *EntityRisksPutPreviousValues) GetRiskCategoryIdOk() (*int32, bool)`

GetRiskCategoryIdOk returns a tuple with the RiskCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCategoryId

`func (o *EntityRisksPutPreviousValues) SetRiskCategoryId(v int32)`

SetRiskCategoryId sets RiskCategoryId field to given value.

### HasRiskCategoryId

`func (o *EntityRisksPutPreviousValues) HasRiskCategoryId() bool`

HasRiskCategoryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


