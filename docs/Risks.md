# Risks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**ProcessId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;processes.id&#x60;.&lt;fk table&#x3D;&#39;processes&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**SubprocessId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;subprocesses.id&#x60;.&lt;fk table&#x3D;&#39;subprocesses&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect1OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select1_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select1_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect2OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select2_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select2_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Activity** | Pointer to **string** |  | [optional] 
**MitigationFactors** | Pointer to **string** |  | [optional] 
**CustomText1** | Pointer to **string** |  | [optional] 
**CustomText2** | Pointer to **string** |  | [optional] 
**RiskCategoryId** | **int32** | Note: This is a Foreign Key to &#x60;risk_categories.id&#x60;.&lt;fk table&#x3D;&#39;risk_categories&#39; column&#x3D;&#39;id&#39;/&gt; | 
**ReviewerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AssessmentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;assessments.id&#x60;.&lt;fk table&#x3D;&#39;assessments&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect3OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select3_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select3_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect4OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select4_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select4_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect5OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select5_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select5_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect6OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select6_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select6_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomText3** | Pointer to **string** |  | [optional] 
**CustomText4** | Pointer to **string** |  | [optional] 
**CustomText5** | Pointer to **string** |  | [optional] 
**CustomText6** | Pointer to **string** |  | [optional] 
**CustomText7** | Pointer to **string** |  | [optional] 
**CustomText8** | Pointer to **string** |  | [optional] 
**CustomText9** | Pointer to **string** |  | [optional] 
**CustomText10** | Pointer to **string** |  | [optional] 
**CustomDate1** | Pointer to **string** |  | [optional] 
**CustomDate2** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**RiskTypeId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_types.id&#x60;.&lt;fk table&#x3D;&#39;risk_types&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect7OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select7_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select7_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect8OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select8_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select8_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect9OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select9_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select9_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect10OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select10_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select10_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect11OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select11_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select11_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect12OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select12_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select12_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FieldData** | Pointer to **interface{}** |  | [optional] 
**InherentRisk** | Pointer to **interface{}** |  | [optional] 
**ResidualRisk** | Pointer to **interface{}** |  | [optional] 
**ReferenceMeta** | **interface{}** |  | 
**CustomDate3** | Pointer to **string** |  | [optional] 
**CustomDate4** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskResponseId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_responses.id&#x60;.&lt;fk table&#x3D;&#39;risk_responses&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**PrimaryOwnerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FormulaData** | **interface{}** |  | 
**ResidualRiskCalc** | Pointer to **interface{}** |  | [optional] 
**IsFlagged** | **bool** |  | [default to false]
**PriorAuditId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audits.id&#x60;.&lt;fk table&#x3D;&#39;ops_audits&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**UpcomingAuditId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audits.id&#x60;.&lt;fk table&#x3D;&#39;ops_audits&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**PriorAuditFinalReportDate** | Pointer to **string** |  | [optional] 
**ExternalParentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risks.id&#x60;.&lt;fk table&#x3D;&#39;risks&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Xuid** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | [default to "Active"]
**CustomJson1** | Pointer to **interface{}** |  | [optional] 
**CustomUserSelect1UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomUserSelect2UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomUserSelect3UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomUserSelect4UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomUserSelect5UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomUserSelect6UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomUserSelect7UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomUserSelect8UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomUserSelect9UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomUserSelect10UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomDate5** | Pointer to **string** |  | [optional] 
**CustomDate6** | Pointer to **string** |  | [optional] 
**CustomDate7** | Pointer to **string** |  | [optional] 
**CustomDate8** | Pointer to **string** |  | [optional] 
**CustomDate9** | Pointer to **string** |  | [optional] 
**CustomDate10** | Pointer to **string** |  | [optional] 
**CustomDate11** | Pointer to **string** |  | [optional] 
**CustomDate12** | Pointer to **string** |  | [optional] 
**CustomDate13** | Pointer to **string** |  | [optional] 
**CustomDate14** | Pointer to **string** |  | [optional] 
**CustomText11** | Pointer to **string** |  | [optional] 
**CustomText12** | Pointer to **string** |  | [optional] 
**CustomText13** | Pointer to **string** |  | [optional] 
**CustomText14** | Pointer to **string** |  | [optional] 
**CustomText15** | Pointer to **string** |  | [optional] 
**CustomText16** | Pointer to **string** |  | [optional] 
**CustomText17** | Pointer to **string** |  | [optional] 
**CustomText18** | Pointer to **string** |  | [optional] 
**CustomText19** | Pointer to **string** |  | [optional] 
**CustomText20** | Pointer to **string** |  | [optional] 
**CustomText21** | Pointer to **string** |  | [optional] 
**CustomText22** | Pointer to **string** |  | [optional] 
**CustomText23** | Pointer to **string** |  | [optional] 
**CustomText24** | Pointer to **string** |  | [optional] 
**CustomText25** | Pointer to **string** |  | [optional] 
**RiskCustomSelect13OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select13_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select13_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect14OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select14_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select14_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect15OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select15_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select15_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect16OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select16_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select16_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect17OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select17_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select17_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect18OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select18_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select18_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect19OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select19_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select19_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect20OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select20_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select20_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect21OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select21_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select21_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect22OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select22_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select22_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect23OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select23_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select23_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect24OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select24_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select24_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect25OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select25_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select25_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect26OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select26_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select26_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskCustomSelect27OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_custom_select27_options.id&#x60;.&lt;fk table&#x3D;&#39;risk_custom_select27_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskAppetiteId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_appetites.id&#x60;.&lt;fk table&#x3D;&#39;risk_appetites&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskAppetiteScore** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewRisks

`func NewRisks(uid string, riskCategoryId int32, referenceMeta interface{}, formulaData interface{}, isFlagged bool, status string, ) *Risks`

NewRisks instantiates a new Risks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRisksWithDefaults

`func NewRisksWithDefaults() *Risks`

NewRisksWithDefaults instantiates a new Risks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Risks) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Risks) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Risks) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Risks) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Risks) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Risks) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Risks) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Risks) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Risks) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Risks) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Risks) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Risks) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Risks) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Risks) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Risks) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Risks) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetUid

`func (o *Risks) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Risks) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Risks) SetUid(v string)`

SetUid sets Uid field to given value.


### GetName

`func (o *Risks) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Risks) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Risks) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Risks) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProcessId

`func (o *Risks) GetProcessId() int32`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *Risks) GetProcessIdOk() (*int32, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *Risks) SetProcessId(v int32)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *Risks) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### GetSubprocessId

`func (o *Risks) GetSubprocessId() int32`

GetSubprocessId returns the SubprocessId field if non-nil, zero value otherwise.

### GetSubprocessIdOk

`func (o *Risks) GetSubprocessIdOk() (*int32, bool)`

GetSubprocessIdOk returns a tuple with the SubprocessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubprocessId

`func (o *Risks) SetSubprocessId(v int32)`

SetSubprocessId sets SubprocessId field to given value.

### HasSubprocessId

`func (o *Risks) HasSubprocessId() bool`

HasSubprocessId returns a boolean if a field has been set.

### GetRiskCustomSelect1OptionId

`func (o *Risks) GetRiskCustomSelect1OptionId() int32`

GetRiskCustomSelect1OptionId returns the RiskCustomSelect1OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect1OptionIdOk

`func (o *Risks) GetRiskCustomSelect1OptionIdOk() (*int32, bool)`

GetRiskCustomSelect1OptionIdOk returns a tuple with the RiskCustomSelect1OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect1OptionId

`func (o *Risks) SetRiskCustomSelect1OptionId(v int32)`

SetRiskCustomSelect1OptionId sets RiskCustomSelect1OptionId field to given value.

### HasRiskCustomSelect1OptionId

`func (o *Risks) HasRiskCustomSelect1OptionId() bool`

HasRiskCustomSelect1OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect2OptionId

`func (o *Risks) GetRiskCustomSelect2OptionId() int32`

GetRiskCustomSelect2OptionId returns the RiskCustomSelect2OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect2OptionIdOk

`func (o *Risks) GetRiskCustomSelect2OptionIdOk() (*int32, bool)`

GetRiskCustomSelect2OptionIdOk returns a tuple with the RiskCustomSelect2OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect2OptionId

`func (o *Risks) SetRiskCustomSelect2OptionId(v int32)`

SetRiskCustomSelect2OptionId sets RiskCustomSelect2OptionId field to given value.

### HasRiskCustomSelect2OptionId

`func (o *Risks) HasRiskCustomSelect2OptionId() bool`

HasRiskCustomSelect2OptionId returns a boolean if a field has been set.

### GetDescription

`func (o *Risks) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Risks) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Risks) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Risks) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActivity

`func (o *Risks) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *Risks) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *Risks) SetActivity(v string)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *Risks) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetMitigationFactors

`func (o *Risks) GetMitigationFactors() string`

GetMitigationFactors returns the MitigationFactors field if non-nil, zero value otherwise.

### GetMitigationFactorsOk

`func (o *Risks) GetMitigationFactorsOk() (*string, bool)`

GetMitigationFactorsOk returns a tuple with the MitigationFactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigationFactors

`func (o *Risks) SetMitigationFactors(v string)`

SetMitigationFactors sets MitigationFactors field to given value.

### HasMitigationFactors

`func (o *Risks) HasMitigationFactors() bool`

HasMitigationFactors returns a boolean if a field has been set.

### GetCustomText1

`func (o *Risks) GetCustomText1() string`

GetCustomText1 returns the CustomText1 field if non-nil, zero value otherwise.

### GetCustomText1Ok

`func (o *Risks) GetCustomText1Ok() (*string, bool)`

GetCustomText1Ok returns a tuple with the CustomText1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText1

`func (o *Risks) SetCustomText1(v string)`

SetCustomText1 sets CustomText1 field to given value.

### HasCustomText1

`func (o *Risks) HasCustomText1() bool`

HasCustomText1 returns a boolean if a field has been set.

### GetCustomText2

`func (o *Risks) GetCustomText2() string`

GetCustomText2 returns the CustomText2 field if non-nil, zero value otherwise.

### GetCustomText2Ok

`func (o *Risks) GetCustomText2Ok() (*string, bool)`

GetCustomText2Ok returns a tuple with the CustomText2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText2

`func (o *Risks) SetCustomText2(v string)`

SetCustomText2 sets CustomText2 field to given value.

### HasCustomText2

`func (o *Risks) HasCustomText2() bool`

HasCustomText2 returns a boolean if a field has been set.

### GetRiskCategoryId

`func (o *Risks) GetRiskCategoryId() int32`

GetRiskCategoryId returns the RiskCategoryId field if non-nil, zero value otherwise.

### GetRiskCategoryIdOk

`func (o *Risks) GetRiskCategoryIdOk() (*int32, bool)`

GetRiskCategoryIdOk returns a tuple with the RiskCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCategoryId

`func (o *Risks) SetRiskCategoryId(v int32)`

SetRiskCategoryId sets RiskCategoryId field to given value.


### GetReviewerUserId

`func (o *Risks) GetReviewerUserId() int32`

GetReviewerUserId returns the ReviewerUserId field if non-nil, zero value otherwise.

### GetReviewerUserIdOk

`func (o *Risks) GetReviewerUserIdOk() (*int32, bool)`

GetReviewerUserIdOk returns a tuple with the ReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerUserId

`func (o *Risks) SetReviewerUserId(v int32)`

SetReviewerUserId sets ReviewerUserId field to given value.

### HasReviewerUserId

`func (o *Risks) HasReviewerUserId() bool`

HasReviewerUserId returns a boolean if a field has been set.

### GetAssessmentId

`func (o *Risks) GetAssessmentId() int32`

GetAssessmentId returns the AssessmentId field if non-nil, zero value otherwise.

### GetAssessmentIdOk

`func (o *Risks) GetAssessmentIdOk() (*int32, bool)`

GetAssessmentIdOk returns a tuple with the AssessmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessmentId

`func (o *Risks) SetAssessmentId(v int32)`

SetAssessmentId sets AssessmentId field to given value.

### HasAssessmentId

`func (o *Risks) HasAssessmentId() bool`

HasAssessmentId returns a boolean if a field has been set.

### GetRiskCustomSelect3OptionId

`func (o *Risks) GetRiskCustomSelect3OptionId() int32`

GetRiskCustomSelect3OptionId returns the RiskCustomSelect3OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect3OptionIdOk

`func (o *Risks) GetRiskCustomSelect3OptionIdOk() (*int32, bool)`

GetRiskCustomSelect3OptionIdOk returns a tuple with the RiskCustomSelect3OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect3OptionId

`func (o *Risks) SetRiskCustomSelect3OptionId(v int32)`

SetRiskCustomSelect3OptionId sets RiskCustomSelect3OptionId field to given value.

### HasRiskCustomSelect3OptionId

`func (o *Risks) HasRiskCustomSelect3OptionId() bool`

HasRiskCustomSelect3OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect4OptionId

`func (o *Risks) GetRiskCustomSelect4OptionId() int32`

GetRiskCustomSelect4OptionId returns the RiskCustomSelect4OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect4OptionIdOk

`func (o *Risks) GetRiskCustomSelect4OptionIdOk() (*int32, bool)`

GetRiskCustomSelect4OptionIdOk returns a tuple with the RiskCustomSelect4OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect4OptionId

`func (o *Risks) SetRiskCustomSelect4OptionId(v int32)`

SetRiskCustomSelect4OptionId sets RiskCustomSelect4OptionId field to given value.

### HasRiskCustomSelect4OptionId

`func (o *Risks) HasRiskCustomSelect4OptionId() bool`

HasRiskCustomSelect4OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect5OptionId

`func (o *Risks) GetRiskCustomSelect5OptionId() int32`

GetRiskCustomSelect5OptionId returns the RiskCustomSelect5OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect5OptionIdOk

`func (o *Risks) GetRiskCustomSelect5OptionIdOk() (*int32, bool)`

GetRiskCustomSelect5OptionIdOk returns a tuple with the RiskCustomSelect5OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect5OptionId

`func (o *Risks) SetRiskCustomSelect5OptionId(v int32)`

SetRiskCustomSelect5OptionId sets RiskCustomSelect5OptionId field to given value.

### HasRiskCustomSelect5OptionId

`func (o *Risks) HasRiskCustomSelect5OptionId() bool`

HasRiskCustomSelect5OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect6OptionId

`func (o *Risks) GetRiskCustomSelect6OptionId() int32`

GetRiskCustomSelect6OptionId returns the RiskCustomSelect6OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect6OptionIdOk

`func (o *Risks) GetRiskCustomSelect6OptionIdOk() (*int32, bool)`

GetRiskCustomSelect6OptionIdOk returns a tuple with the RiskCustomSelect6OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect6OptionId

`func (o *Risks) SetRiskCustomSelect6OptionId(v int32)`

SetRiskCustomSelect6OptionId sets RiskCustomSelect6OptionId field to given value.

### HasRiskCustomSelect6OptionId

`func (o *Risks) HasRiskCustomSelect6OptionId() bool`

HasRiskCustomSelect6OptionId returns a boolean if a field has been set.

### GetCustomText3

`func (o *Risks) GetCustomText3() string`

GetCustomText3 returns the CustomText3 field if non-nil, zero value otherwise.

### GetCustomText3Ok

`func (o *Risks) GetCustomText3Ok() (*string, bool)`

GetCustomText3Ok returns a tuple with the CustomText3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText3

`func (o *Risks) SetCustomText3(v string)`

SetCustomText3 sets CustomText3 field to given value.

### HasCustomText3

`func (o *Risks) HasCustomText3() bool`

HasCustomText3 returns a boolean if a field has been set.

### GetCustomText4

`func (o *Risks) GetCustomText4() string`

GetCustomText4 returns the CustomText4 field if non-nil, zero value otherwise.

### GetCustomText4Ok

`func (o *Risks) GetCustomText4Ok() (*string, bool)`

GetCustomText4Ok returns a tuple with the CustomText4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText4

`func (o *Risks) SetCustomText4(v string)`

SetCustomText4 sets CustomText4 field to given value.

### HasCustomText4

`func (o *Risks) HasCustomText4() bool`

HasCustomText4 returns a boolean if a field has been set.

### GetCustomText5

`func (o *Risks) GetCustomText5() string`

GetCustomText5 returns the CustomText5 field if non-nil, zero value otherwise.

### GetCustomText5Ok

`func (o *Risks) GetCustomText5Ok() (*string, bool)`

GetCustomText5Ok returns a tuple with the CustomText5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText5

`func (o *Risks) SetCustomText5(v string)`

SetCustomText5 sets CustomText5 field to given value.

### HasCustomText5

`func (o *Risks) HasCustomText5() bool`

HasCustomText5 returns a boolean if a field has been set.

### GetCustomText6

`func (o *Risks) GetCustomText6() string`

GetCustomText6 returns the CustomText6 field if non-nil, zero value otherwise.

### GetCustomText6Ok

`func (o *Risks) GetCustomText6Ok() (*string, bool)`

GetCustomText6Ok returns a tuple with the CustomText6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText6

`func (o *Risks) SetCustomText6(v string)`

SetCustomText6 sets CustomText6 field to given value.

### HasCustomText6

`func (o *Risks) HasCustomText6() bool`

HasCustomText6 returns a boolean if a field has been set.

### GetCustomText7

`func (o *Risks) GetCustomText7() string`

GetCustomText7 returns the CustomText7 field if non-nil, zero value otherwise.

### GetCustomText7Ok

`func (o *Risks) GetCustomText7Ok() (*string, bool)`

GetCustomText7Ok returns a tuple with the CustomText7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText7

`func (o *Risks) SetCustomText7(v string)`

SetCustomText7 sets CustomText7 field to given value.

### HasCustomText7

`func (o *Risks) HasCustomText7() bool`

HasCustomText7 returns a boolean if a field has been set.

### GetCustomText8

`func (o *Risks) GetCustomText8() string`

GetCustomText8 returns the CustomText8 field if non-nil, zero value otherwise.

### GetCustomText8Ok

`func (o *Risks) GetCustomText8Ok() (*string, bool)`

GetCustomText8Ok returns a tuple with the CustomText8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText8

`func (o *Risks) SetCustomText8(v string)`

SetCustomText8 sets CustomText8 field to given value.

### HasCustomText8

`func (o *Risks) HasCustomText8() bool`

HasCustomText8 returns a boolean if a field has been set.

### GetCustomText9

`func (o *Risks) GetCustomText9() string`

GetCustomText9 returns the CustomText9 field if non-nil, zero value otherwise.

### GetCustomText9Ok

`func (o *Risks) GetCustomText9Ok() (*string, bool)`

GetCustomText9Ok returns a tuple with the CustomText9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText9

`func (o *Risks) SetCustomText9(v string)`

SetCustomText9 sets CustomText9 field to given value.

### HasCustomText9

`func (o *Risks) HasCustomText9() bool`

HasCustomText9 returns a boolean if a field has been set.

### GetCustomText10

`func (o *Risks) GetCustomText10() string`

GetCustomText10 returns the CustomText10 field if non-nil, zero value otherwise.

### GetCustomText10Ok

`func (o *Risks) GetCustomText10Ok() (*string, bool)`

GetCustomText10Ok returns a tuple with the CustomText10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText10

`func (o *Risks) SetCustomText10(v string)`

SetCustomText10 sets CustomText10 field to given value.

### HasCustomText10

`func (o *Risks) HasCustomText10() bool`

HasCustomText10 returns a boolean if a field has been set.

### GetCustomDate1

`func (o *Risks) GetCustomDate1() string`

GetCustomDate1 returns the CustomDate1 field if non-nil, zero value otherwise.

### GetCustomDate1Ok

`func (o *Risks) GetCustomDate1Ok() (*string, bool)`

GetCustomDate1Ok returns a tuple with the CustomDate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate1

`func (o *Risks) SetCustomDate1(v string)`

SetCustomDate1 sets CustomDate1 field to given value.

### HasCustomDate1

`func (o *Risks) HasCustomDate1() bool`

HasCustomDate1 returns a boolean if a field has been set.

### GetCustomDate2

`func (o *Risks) GetCustomDate2() string`

GetCustomDate2 returns the CustomDate2 field if non-nil, zero value otherwise.

### GetCustomDate2Ok

`func (o *Risks) GetCustomDate2Ok() (*string, bool)`

GetCustomDate2Ok returns a tuple with the CustomDate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate2

`func (o *Risks) SetCustomDate2(v string)`

SetCustomDate2 sets CustomDate2 field to given value.

### HasCustomDate2

`func (o *Risks) HasCustomDate2() bool`

HasCustomDate2 returns a boolean if a field has been set.

### GetNotes

`func (o *Risks) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Risks) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Risks) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Risks) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetRiskTypeId

`func (o *Risks) GetRiskTypeId() int32`

GetRiskTypeId returns the RiskTypeId field if non-nil, zero value otherwise.

### GetRiskTypeIdOk

`func (o *Risks) GetRiskTypeIdOk() (*int32, bool)`

GetRiskTypeIdOk returns a tuple with the RiskTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskTypeId

`func (o *Risks) SetRiskTypeId(v int32)`

SetRiskTypeId sets RiskTypeId field to given value.

### HasRiskTypeId

`func (o *Risks) HasRiskTypeId() bool`

HasRiskTypeId returns a boolean if a field has been set.

### GetRiskCustomSelect7OptionId

`func (o *Risks) GetRiskCustomSelect7OptionId() int32`

GetRiskCustomSelect7OptionId returns the RiskCustomSelect7OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect7OptionIdOk

`func (o *Risks) GetRiskCustomSelect7OptionIdOk() (*int32, bool)`

GetRiskCustomSelect7OptionIdOk returns a tuple with the RiskCustomSelect7OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect7OptionId

`func (o *Risks) SetRiskCustomSelect7OptionId(v int32)`

SetRiskCustomSelect7OptionId sets RiskCustomSelect7OptionId field to given value.

### HasRiskCustomSelect7OptionId

`func (o *Risks) HasRiskCustomSelect7OptionId() bool`

HasRiskCustomSelect7OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect8OptionId

`func (o *Risks) GetRiskCustomSelect8OptionId() int32`

GetRiskCustomSelect8OptionId returns the RiskCustomSelect8OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect8OptionIdOk

`func (o *Risks) GetRiskCustomSelect8OptionIdOk() (*int32, bool)`

GetRiskCustomSelect8OptionIdOk returns a tuple with the RiskCustomSelect8OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect8OptionId

`func (o *Risks) SetRiskCustomSelect8OptionId(v int32)`

SetRiskCustomSelect8OptionId sets RiskCustomSelect8OptionId field to given value.

### HasRiskCustomSelect8OptionId

`func (o *Risks) HasRiskCustomSelect8OptionId() bool`

HasRiskCustomSelect8OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect9OptionId

`func (o *Risks) GetRiskCustomSelect9OptionId() int32`

GetRiskCustomSelect9OptionId returns the RiskCustomSelect9OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect9OptionIdOk

`func (o *Risks) GetRiskCustomSelect9OptionIdOk() (*int32, bool)`

GetRiskCustomSelect9OptionIdOk returns a tuple with the RiskCustomSelect9OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect9OptionId

`func (o *Risks) SetRiskCustomSelect9OptionId(v int32)`

SetRiskCustomSelect9OptionId sets RiskCustomSelect9OptionId field to given value.

### HasRiskCustomSelect9OptionId

`func (o *Risks) HasRiskCustomSelect9OptionId() bool`

HasRiskCustomSelect9OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect10OptionId

`func (o *Risks) GetRiskCustomSelect10OptionId() int32`

GetRiskCustomSelect10OptionId returns the RiskCustomSelect10OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect10OptionIdOk

`func (o *Risks) GetRiskCustomSelect10OptionIdOk() (*int32, bool)`

GetRiskCustomSelect10OptionIdOk returns a tuple with the RiskCustomSelect10OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect10OptionId

`func (o *Risks) SetRiskCustomSelect10OptionId(v int32)`

SetRiskCustomSelect10OptionId sets RiskCustomSelect10OptionId field to given value.

### HasRiskCustomSelect10OptionId

`func (o *Risks) HasRiskCustomSelect10OptionId() bool`

HasRiskCustomSelect10OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect11OptionId

`func (o *Risks) GetRiskCustomSelect11OptionId() int32`

GetRiskCustomSelect11OptionId returns the RiskCustomSelect11OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect11OptionIdOk

`func (o *Risks) GetRiskCustomSelect11OptionIdOk() (*int32, bool)`

GetRiskCustomSelect11OptionIdOk returns a tuple with the RiskCustomSelect11OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect11OptionId

`func (o *Risks) SetRiskCustomSelect11OptionId(v int32)`

SetRiskCustomSelect11OptionId sets RiskCustomSelect11OptionId field to given value.

### HasRiskCustomSelect11OptionId

`func (o *Risks) HasRiskCustomSelect11OptionId() bool`

HasRiskCustomSelect11OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect12OptionId

`func (o *Risks) GetRiskCustomSelect12OptionId() int32`

GetRiskCustomSelect12OptionId returns the RiskCustomSelect12OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect12OptionIdOk

`func (o *Risks) GetRiskCustomSelect12OptionIdOk() (*int32, bool)`

GetRiskCustomSelect12OptionIdOk returns a tuple with the RiskCustomSelect12OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect12OptionId

`func (o *Risks) SetRiskCustomSelect12OptionId(v int32)`

SetRiskCustomSelect12OptionId sets RiskCustomSelect12OptionId field to given value.

### HasRiskCustomSelect12OptionId

`func (o *Risks) HasRiskCustomSelect12OptionId() bool`

HasRiskCustomSelect12OptionId returns a boolean if a field has been set.

### GetFieldData

`func (o *Risks) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *Risks) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *Risks) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *Risks) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *Risks) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *Risks) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil
### GetInherentRisk

`func (o *Risks) GetInherentRisk() interface{}`

GetInherentRisk returns the InherentRisk field if non-nil, zero value otherwise.

### GetInherentRiskOk

`func (o *Risks) GetInherentRiskOk() (*interface{}, bool)`

GetInherentRiskOk returns a tuple with the InherentRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInherentRisk

`func (o *Risks) SetInherentRisk(v interface{})`

SetInherentRisk sets InherentRisk field to given value.

### HasInherentRisk

`func (o *Risks) HasInherentRisk() bool`

HasInherentRisk returns a boolean if a field has been set.

### SetInherentRiskNil

`func (o *Risks) SetInherentRiskNil(b bool)`

 SetInherentRiskNil sets the value for InherentRisk to be an explicit nil

### UnsetInherentRisk
`func (o *Risks) UnsetInherentRisk()`

UnsetInherentRisk ensures that no value is present for InherentRisk, not even an explicit nil
### GetResidualRisk

`func (o *Risks) GetResidualRisk() interface{}`

GetResidualRisk returns the ResidualRisk field if non-nil, zero value otherwise.

### GetResidualRiskOk

`func (o *Risks) GetResidualRiskOk() (*interface{}, bool)`

GetResidualRiskOk returns a tuple with the ResidualRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidualRisk

`func (o *Risks) SetResidualRisk(v interface{})`

SetResidualRisk sets ResidualRisk field to given value.

### HasResidualRisk

`func (o *Risks) HasResidualRisk() bool`

HasResidualRisk returns a boolean if a field has been set.

### SetResidualRiskNil

`func (o *Risks) SetResidualRiskNil(b bool)`

 SetResidualRiskNil sets the value for ResidualRisk to be an explicit nil

### UnsetResidualRisk
`func (o *Risks) UnsetResidualRisk()`

UnsetResidualRisk ensures that no value is present for ResidualRisk, not even an explicit nil
### GetReferenceMeta

`func (o *Risks) GetReferenceMeta() interface{}`

GetReferenceMeta returns the ReferenceMeta field if non-nil, zero value otherwise.

### GetReferenceMetaOk

`func (o *Risks) GetReferenceMetaOk() (*interface{}, bool)`

GetReferenceMetaOk returns a tuple with the ReferenceMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMeta

`func (o *Risks) SetReferenceMeta(v interface{})`

SetReferenceMeta sets ReferenceMeta field to given value.


### SetReferenceMetaNil

`func (o *Risks) SetReferenceMetaNil(b bool)`

 SetReferenceMetaNil sets the value for ReferenceMeta to be an explicit nil

### UnsetReferenceMeta
`func (o *Risks) UnsetReferenceMeta()`

UnsetReferenceMeta ensures that no value is present for ReferenceMeta, not even an explicit nil
### GetCustomDate3

`func (o *Risks) GetCustomDate3() string`

GetCustomDate3 returns the CustomDate3 field if non-nil, zero value otherwise.

### GetCustomDate3Ok

`func (o *Risks) GetCustomDate3Ok() (*string, bool)`

GetCustomDate3Ok returns a tuple with the CustomDate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate3

`func (o *Risks) SetCustomDate3(v string)`

SetCustomDate3 sets CustomDate3 field to given value.

### HasCustomDate3

`func (o *Risks) HasCustomDate3() bool`

HasCustomDate3 returns a boolean if a field has been set.

### GetCustomDate4

`func (o *Risks) GetCustomDate4() string`

GetCustomDate4 returns the CustomDate4 field if non-nil, zero value otherwise.

### GetCustomDate4Ok

`func (o *Risks) GetCustomDate4Ok() (*string, bool)`

GetCustomDate4Ok returns a tuple with the CustomDate4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate4

`func (o *Risks) SetCustomDate4(v string)`

SetCustomDate4 sets CustomDate4 field to given value.

### HasCustomDate4

`func (o *Risks) HasCustomDate4() bool`

HasCustomDate4 returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *Risks) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *Risks) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *Risks) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *Risks) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetRiskResponseId

`func (o *Risks) GetRiskResponseId() int32`

GetRiskResponseId returns the RiskResponseId field if non-nil, zero value otherwise.

### GetRiskResponseIdOk

`func (o *Risks) GetRiskResponseIdOk() (*int32, bool)`

GetRiskResponseIdOk returns a tuple with the RiskResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskResponseId

`func (o *Risks) SetRiskResponseId(v int32)`

SetRiskResponseId sets RiskResponseId field to given value.

### HasRiskResponseId

`func (o *Risks) HasRiskResponseId() bool`

HasRiskResponseId returns a boolean if a field has been set.

### GetPrimaryOwnerUserId

`func (o *Risks) GetPrimaryOwnerUserId() int32`

GetPrimaryOwnerUserId returns the PrimaryOwnerUserId field if non-nil, zero value otherwise.

### GetPrimaryOwnerUserIdOk

`func (o *Risks) GetPrimaryOwnerUserIdOk() (*int32, bool)`

GetPrimaryOwnerUserIdOk returns a tuple with the PrimaryOwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryOwnerUserId

`func (o *Risks) SetPrimaryOwnerUserId(v int32)`

SetPrimaryOwnerUserId sets PrimaryOwnerUserId field to given value.

### HasPrimaryOwnerUserId

`func (o *Risks) HasPrimaryOwnerUserId() bool`

HasPrimaryOwnerUserId returns a boolean if a field has been set.

### GetFormulaData

`func (o *Risks) GetFormulaData() interface{}`

GetFormulaData returns the FormulaData field if non-nil, zero value otherwise.

### GetFormulaDataOk

`func (o *Risks) GetFormulaDataOk() (*interface{}, bool)`

GetFormulaDataOk returns a tuple with the FormulaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulaData

`func (o *Risks) SetFormulaData(v interface{})`

SetFormulaData sets FormulaData field to given value.


### SetFormulaDataNil

`func (o *Risks) SetFormulaDataNil(b bool)`

 SetFormulaDataNil sets the value for FormulaData to be an explicit nil

### UnsetFormulaData
`func (o *Risks) UnsetFormulaData()`

UnsetFormulaData ensures that no value is present for FormulaData, not even an explicit nil
### GetResidualRiskCalc

`func (o *Risks) GetResidualRiskCalc() interface{}`

GetResidualRiskCalc returns the ResidualRiskCalc field if non-nil, zero value otherwise.

### GetResidualRiskCalcOk

`func (o *Risks) GetResidualRiskCalcOk() (*interface{}, bool)`

GetResidualRiskCalcOk returns a tuple with the ResidualRiskCalc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidualRiskCalc

`func (o *Risks) SetResidualRiskCalc(v interface{})`

SetResidualRiskCalc sets ResidualRiskCalc field to given value.

### HasResidualRiskCalc

`func (o *Risks) HasResidualRiskCalc() bool`

HasResidualRiskCalc returns a boolean if a field has been set.

### SetResidualRiskCalcNil

`func (o *Risks) SetResidualRiskCalcNil(b bool)`

 SetResidualRiskCalcNil sets the value for ResidualRiskCalc to be an explicit nil

### UnsetResidualRiskCalc
`func (o *Risks) UnsetResidualRiskCalc()`

UnsetResidualRiskCalc ensures that no value is present for ResidualRiskCalc, not even an explicit nil
### GetIsFlagged

`func (o *Risks) GetIsFlagged() bool`

GetIsFlagged returns the IsFlagged field if non-nil, zero value otherwise.

### GetIsFlaggedOk

`func (o *Risks) GetIsFlaggedOk() (*bool, bool)`

GetIsFlaggedOk returns a tuple with the IsFlagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlagged

`func (o *Risks) SetIsFlagged(v bool)`

SetIsFlagged sets IsFlagged field to given value.


### GetPriorAuditId

`func (o *Risks) GetPriorAuditId() int32`

GetPriorAuditId returns the PriorAuditId field if non-nil, zero value otherwise.

### GetPriorAuditIdOk

`func (o *Risks) GetPriorAuditIdOk() (*int32, bool)`

GetPriorAuditIdOk returns a tuple with the PriorAuditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorAuditId

`func (o *Risks) SetPriorAuditId(v int32)`

SetPriorAuditId sets PriorAuditId field to given value.

### HasPriorAuditId

`func (o *Risks) HasPriorAuditId() bool`

HasPriorAuditId returns a boolean if a field has been set.

### GetUpcomingAuditId

`func (o *Risks) GetUpcomingAuditId() int32`

GetUpcomingAuditId returns the UpcomingAuditId field if non-nil, zero value otherwise.

### GetUpcomingAuditIdOk

`func (o *Risks) GetUpcomingAuditIdOk() (*int32, bool)`

GetUpcomingAuditIdOk returns a tuple with the UpcomingAuditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcomingAuditId

`func (o *Risks) SetUpcomingAuditId(v int32)`

SetUpcomingAuditId sets UpcomingAuditId field to given value.

### HasUpcomingAuditId

`func (o *Risks) HasUpcomingAuditId() bool`

HasUpcomingAuditId returns a boolean if a field has been set.

### GetPriorAuditFinalReportDate

`func (o *Risks) GetPriorAuditFinalReportDate() string`

GetPriorAuditFinalReportDate returns the PriorAuditFinalReportDate field if non-nil, zero value otherwise.

### GetPriorAuditFinalReportDateOk

`func (o *Risks) GetPriorAuditFinalReportDateOk() (*string, bool)`

GetPriorAuditFinalReportDateOk returns a tuple with the PriorAuditFinalReportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorAuditFinalReportDate

`func (o *Risks) SetPriorAuditFinalReportDate(v string)`

SetPriorAuditFinalReportDate sets PriorAuditFinalReportDate field to given value.

### HasPriorAuditFinalReportDate

`func (o *Risks) HasPriorAuditFinalReportDate() bool`

HasPriorAuditFinalReportDate returns a boolean if a field has been set.

### GetExternalParentId

`func (o *Risks) GetExternalParentId() int32`

GetExternalParentId returns the ExternalParentId field if non-nil, zero value otherwise.

### GetExternalParentIdOk

`func (o *Risks) GetExternalParentIdOk() (*int32, bool)`

GetExternalParentIdOk returns a tuple with the ExternalParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalParentId

`func (o *Risks) SetExternalParentId(v int32)`

SetExternalParentId sets ExternalParentId field to given value.

### HasExternalParentId

`func (o *Risks) HasExternalParentId() bool`

HasExternalParentId returns a boolean if a field has been set.

### GetXuid

`func (o *Risks) GetXuid() string`

GetXuid returns the Xuid field if non-nil, zero value otherwise.

### GetXuidOk

`func (o *Risks) GetXuidOk() (*string, bool)`

GetXuidOk returns a tuple with the Xuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXuid

`func (o *Risks) SetXuid(v string)`

SetXuid sets Xuid field to given value.

### HasXuid

`func (o *Risks) HasXuid() bool`

HasXuid returns a boolean if a field has been set.

### GetStatus

`func (o *Risks) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Risks) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Risks) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCustomJson1

`func (o *Risks) GetCustomJson1() interface{}`

GetCustomJson1 returns the CustomJson1 field if non-nil, zero value otherwise.

### GetCustomJson1Ok

`func (o *Risks) GetCustomJson1Ok() (*interface{}, bool)`

GetCustomJson1Ok returns a tuple with the CustomJson1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomJson1

`func (o *Risks) SetCustomJson1(v interface{})`

SetCustomJson1 sets CustomJson1 field to given value.

### HasCustomJson1

`func (o *Risks) HasCustomJson1() bool`

HasCustomJson1 returns a boolean if a field has been set.

### SetCustomJson1Nil

`func (o *Risks) SetCustomJson1Nil(b bool)`

 SetCustomJson1Nil sets the value for CustomJson1 to be an explicit nil

### UnsetCustomJson1
`func (o *Risks) UnsetCustomJson1()`

UnsetCustomJson1 ensures that no value is present for CustomJson1, not even an explicit nil
### GetCustomUserSelect1UserId

`func (o *Risks) GetCustomUserSelect1UserId() int32`

GetCustomUserSelect1UserId returns the CustomUserSelect1UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect1UserIdOk

`func (o *Risks) GetCustomUserSelect1UserIdOk() (*int32, bool)`

GetCustomUserSelect1UserIdOk returns a tuple with the CustomUserSelect1UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect1UserId

`func (o *Risks) SetCustomUserSelect1UserId(v int32)`

SetCustomUserSelect1UserId sets CustomUserSelect1UserId field to given value.

### HasCustomUserSelect1UserId

`func (o *Risks) HasCustomUserSelect1UserId() bool`

HasCustomUserSelect1UserId returns a boolean if a field has been set.

### GetCustomUserSelect2UserId

`func (o *Risks) GetCustomUserSelect2UserId() int32`

GetCustomUserSelect2UserId returns the CustomUserSelect2UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect2UserIdOk

`func (o *Risks) GetCustomUserSelect2UserIdOk() (*int32, bool)`

GetCustomUserSelect2UserIdOk returns a tuple with the CustomUserSelect2UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect2UserId

`func (o *Risks) SetCustomUserSelect2UserId(v int32)`

SetCustomUserSelect2UserId sets CustomUserSelect2UserId field to given value.

### HasCustomUserSelect2UserId

`func (o *Risks) HasCustomUserSelect2UserId() bool`

HasCustomUserSelect2UserId returns a boolean if a field has been set.

### GetCustomUserSelect3UserId

`func (o *Risks) GetCustomUserSelect3UserId() int32`

GetCustomUserSelect3UserId returns the CustomUserSelect3UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect3UserIdOk

`func (o *Risks) GetCustomUserSelect3UserIdOk() (*int32, bool)`

GetCustomUserSelect3UserIdOk returns a tuple with the CustomUserSelect3UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect3UserId

`func (o *Risks) SetCustomUserSelect3UserId(v int32)`

SetCustomUserSelect3UserId sets CustomUserSelect3UserId field to given value.

### HasCustomUserSelect3UserId

`func (o *Risks) HasCustomUserSelect3UserId() bool`

HasCustomUserSelect3UserId returns a boolean if a field has been set.

### GetCustomUserSelect4UserId

`func (o *Risks) GetCustomUserSelect4UserId() int32`

GetCustomUserSelect4UserId returns the CustomUserSelect4UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect4UserIdOk

`func (o *Risks) GetCustomUserSelect4UserIdOk() (*int32, bool)`

GetCustomUserSelect4UserIdOk returns a tuple with the CustomUserSelect4UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect4UserId

`func (o *Risks) SetCustomUserSelect4UserId(v int32)`

SetCustomUserSelect4UserId sets CustomUserSelect4UserId field to given value.

### HasCustomUserSelect4UserId

`func (o *Risks) HasCustomUserSelect4UserId() bool`

HasCustomUserSelect4UserId returns a boolean if a field has been set.

### GetCustomUserSelect5UserId

`func (o *Risks) GetCustomUserSelect5UserId() int32`

GetCustomUserSelect5UserId returns the CustomUserSelect5UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect5UserIdOk

`func (o *Risks) GetCustomUserSelect5UserIdOk() (*int32, bool)`

GetCustomUserSelect5UserIdOk returns a tuple with the CustomUserSelect5UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect5UserId

`func (o *Risks) SetCustomUserSelect5UserId(v int32)`

SetCustomUserSelect5UserId sets CustomUserSelect5UserId field to given value.

### HasCustomUserSelect5UserId

`func (o *Risks) HasCustomUserSelect5UserId() bool`

HasCustomUserSelect5UserId returns a boolean if a field has been set.

### GetCustomUserSelect6UserId

`func (o *Risks) GetCustomUserSelect6UserId() int32`

GetCustomUserSelect6UserId returns the CustomUserSelect6UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect6UserIdOk

`func (o *Risks) GetCustomUserSelect6UserIdOk() (*int32, bool)`

GetCustomUserSelect6UserIdOk returns a tuple with the CustomUserSelect6UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect6UserId

`func (o *Risks) SetCustomUserSelect6UserId(v int32)`

SetCustomUserSelect6UserId sets CustomUserSelect6UserId field to given value.

### HasCustomUserSelect6UserId

`func (o *Risks) HasCustomUserSelect6UserId() bool`

HasCustomUserSelect6UserId returns a boolean if a field has been set.

### GetCustomUserSelect7UserId

`func (o *Risks) GetCustomUserSelect7UserId() int32`

GetCustomUserSelect7UserId returns the CustomUserSelect7UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect7UserIdOk

`func (o *Risks) GetCustomUserSelect7UserIdOk() (*int32, bool)`

GetCustomUserSelect7UserIdOk returns a tuple with the CustomUserSelect7UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect7UserId

`func (o *Risks) SetCustomUserSelect7UserId(v int32)`

SetCustomUserSelect7UserId sets CustomUserSelect7UserId field to given value.

### HasCustomUserSelect7UserId

`func (o *Risks) HasCustomUserSelect7UserId() bool`

HasCustomUserSelect7UserId returns a boolean if a field has been set.

### GetCustomUserSelect8UserId

`func (o *Risks) GetCustomUserSelect8UserId() int32`

GetCustomUserSelect8UserId returns the CustomUserSelect8UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect8UserIdOk

`func (o *Risks) GetCustomUserSelect8UserIdOk() (*int32, bool)`

GetCustomUserSelect8UserIdOk returns a tuple with the CustomUserSelect8UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect8UserId

`func (o *Risks) SetCustomUserSelect8UserId(v int32)`

SetCustomUserSelect8UserId sets CustomUserSelect8UserId field to given value.

### HasCustomUserSelect8UserId

`func (o *Risks) HasCustomUserSelect8UserId() bool`

HasCustomUserSelect8UserId returns a boolean if a field has been set.

### GetCustomUserSelect9UserId

`func (o *Risks) GetCustomUserSelect9UserId() int32`

GetCustomUserSelect9UserId returns the CustomUserSelect9UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect9UserIdOk

`func (o *Risks) GetCustomUserSelect9UserIdOk() (*int32, bool)`

GetCustomUserSelect9UserIdOk returns a tuple with the CustomUserSelect9UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect9UserId

`func (o *Risks) SetCustomUserSelect9UserId(v int32)`

SetCustomUserSelect9UserId sets CustomUserSelect9UserId field to given value.

### HasCustomUserSelect9UserId

`func (o *Risks) HasCustomUserSelect9UserId() bool`

HasCustomUserSelect9UserId returns a boolean if a field has been set.

### GetCustomUserSelect10UserId

`func (o *Risks) GetCustomUserSelect10UserId() int32`

GetCustomUserSelect10UserId returns the CustomUserSelect10UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect10UserIdOk

`func (o *Risks) GetCustomUserSelect10UserIdOk() (*int32, bool)`

GetCustomUserSelect10UserIdOk returns a tuple with the CustomUserSelect10UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect10UserId

`func (o *Risks) SetCustomUserSelect10UserId(v int32)`

SetCustomUserSelect10UserId sets CustomUserSelect10UserId field to given value.

### HasCustomUserSelect10UserId

`func (o *Risks) HasCustomUserSelect10UserId() bool`

HasCustomUserSelect10UserId returns a boolean if a field has been set.

### GetCustomDate5

`func (o *Risks) GetCustomDate5() string`

GetCustomDate5 returns the CustomDate5 field if non-nil, zero value otherwise.

### GetCustomDate5Ok

`func (o *Risks) GetCustomDate5Ok() (*string, bool)`

GetCustomDate5Ok returns a tuple with the CustomDate5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate5

`func (o *Risks) SetCustomDate5(v string)`

SetCustomDate5 sets CustomDate5 field to given value.

### HasCustomDate5

`func (o *Risks) HasCustomDate5() bool`

HasCustomDate5 returns a boolean if a field has been set.

### GetCustomDate6

`func (o *Risks) GetCustomDate6() string`

GetCustomDate6 returns the CustomDate6 field if non-nil, zero value otherwise.

### GetCustomDate6Ok

`func (o *Risks) GetCustomDate6Ok() (*string, bool)`

GetCustomDate6Ok returns a tuple with the CustomDate6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate6

`func (o *Risks) SetCustomDate6(v string)`

SetCustomDate6 sets CustomDate6 field to given value.

### HasCustomDate6

`func (o *Risks) HasCustomDate6() bool`

HasCustomDate6 returns a boolean if a field has been set.

### GetCustomDate7

`func (o *Risks) GetCustomDate7() string`

GetCustomDate7 returns the CustomDate7 field if non-nil, zero value otherwise.

### GetCustomDate7Ok

`func (o *Risks) GetCustomDate7Ok() (*string, bool)`

GetCustomDate7Ok returns a tuple with the CustomDate7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate7

`func (o *Risks) SetCustomDate7(v string)`

SetCustomDate7 sets CustomDate7 field to given value.

### HasCustomDate7

`func (o *Risks) HasCustomDate7() bool`

HasCustomDate7 returns a boolean if a field has been set.

### GetCustomDate8

`func (o *Risks) GetCustomDate8() string`

GetCustomDate8 returns the CustomDate8 field if non-nil, zero value otherwise.

### GetCustomDate8Ok

`func (o *Risks) GetCustomDate8Ok() (*string, bool)`

GetCustomDate8Ok returns a tuple with the CustomDate8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate8

`func (o *Risks) SetCustomDate8(v string)`

SetCustomDate8 sets CustomDate8 field to given value.

### HasCustomDate8

`func (o *Risks) HasCustomDate8() bool`

HasCustomDate8 returns a boolean if a field has been set.

### GetCustomDate9

`func (o *Risks) GetCustomDate9() string`

GetCustomDate9 returns the CustomDate9 field if non-nil, zero value otherwise.

### GetCustomDate9Ok

`func (o *Risks) GetCustomDate9Ok() (*string, bool)`

GetCustomDate9Ok returns a tuple with the CustomDate9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate9

`func (o *Risks) SetCustomDate9(v string)`

SetCustomDate9 sets CustomDate9 field to given value.

### HasCustomDate9

`func (o *Risks) HasCustomDate9() bool`

HasCustomDate9 returns a boolean if a field has been set.

### GetCustomDate10

`func (o *Risks) GetCustomDate10() string`

GetCustomDate10 returns the CustomDate10 field if non-nil, zero value otherwise.

### GetCustomDate10Ok

`func (o *Risks) GetCustomDate10Ok() (*string, bool)`

GetCustomDate10Ok returns a tuple with the CustomDate10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate10

`func (o *Risks) SetCustomDate10(v string)`

SetCustomDate10 sets CustomDate10 field to given value.

### HasCustomDate10

`func (o *Risks) HasCustomDate10() bool`

HasCustomDate10 returns a boolean if a field has been set.

### GetCustomDate11

`func (o *Risks) GetCustomDate11() string`

GetCustomDate11 returns the CustomDate11 field if non-nil, zero value otherwise.

### GetCustomDate11Ok

`func (o *Risks) GetCustomDate11Ok() (*string, bool)`

GetCustomDate11Ok returns a tuple with the CustomDate11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate11

`func (o *Risks) SetCustomDate11(v string)`

SetCustomDate11 sets CustomDate11 field to given value.

### HasCustomDate11

`func (o *Risks) HasCustomDate11() bool`

HasCustomDate11 returns a boolean if a field has been set.

### GetCustomDate12

`func (o *Risks) GetCustomDate12() string`

GetCustomDate12 returns the CustomDate12 field if non-nil, zero value otherwise.

### GetCustomDate12Ok

`func (o *Risks) GetCustomDate12Ok() (*string, bool)`

GetCustomDate12Ok returns a tuple with the CustomDate12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate12

`func (o *Risks) SetCustomDate12(v string)`

SetCustomDate12 sets CustomDate12 field to given value.

### HasCustomDate12

`func (o *Risks) HasCustomDate12() bool`

HasCustomDate12 returns a boolean if a field has been set.

### GetCustomDate13

`func (o *Risks) GetCustomDate13() string`

GetCustomDate13 returns the CustomDate13 field if non-nil, zero value otherwise.

### GetCustomDate13Ok

`func (o *Risks) GetCustomDate13Ok() (*string, bool)`

GetCustomDate13Ok returns a tuple with the CustomDate13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate13

`func (o *Risks) SetCustomDate13(v string)`

SetCustomDate13 sets CustomDate13 field to given value.

### HasCustomDate13

`func (o *Risks) HasCustomDate13() bool`

HasCustomDate13 returns a boolean if a field has been set.

### GetCustomDate14

`func (o *Risks) GetCustomDate14() string`

GetCustomDate14 returns the CustomDate14 field if non-nil, zero value otherwise.

### GetCustomDate14Ok

`func (o *Risks) GetCustomDate14Ok() (*string, bool)`

GetCustomDate14Ok returns a tuple with the CustomDate14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate14

`func (o *Risks) SetCustomDate14(v string)`

SetCustomDate14 sets CustomDate14 field to given value.

### HasCustomDate14

`func (o *Risks) HasCustomDate14() bool`

HasCustomDate14 returns a boolean if a field has been set.

### GetCustomText11

`func (o *Risks) GetCustomText11() string`

GetCustomText11 returns the CustomText11 field if non-nil, zero value otherwise.

### GetCustomText11Ok

`func (o *Risks) GetCustomText11Ok() (*string, bool)`

GetCustomText11Ok returns a tuple with the CustomText11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText11

`func (o *Risks) SetCustomText11(v string)`

SetCustomText11 sets CustomText11 field to given value.

### HasCustomText11

`func (o *Risks) HasCustomText11() bool`

HasCustomText11 returns a boolean if a field has been set.

### GetCustomText12

`func (o *Risks) GetCustomText12() string`

GetCustomText12 returns the CustomText12 field if non-nil, zero value otherwise.

### GetCustomText12Ok

`func (o *Risks) GetCustomText12Ok() (*string, bool)`

GetCustomText12Ok returns a tuple with the CustomText12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText12

`func (o *Risks) SetCustomText12(v string)`

SetCustomText12 sets CustomText12 field to given value.

### HasCustomText12

`func (o *Risks) HasCustomText12() bool`

HasCustomText12 returns a boolean if a field has been set.

### GetCustomText13

`func (o *Risks) GetCustomText13() string`

GetCustomText13 returns the CustomText13 field if non-nil, zero value otherwise.

### GetCustomText13Ok

`func (o *Risks) GetCustomText13Ok() (*string, bool)`

GetCustomText13Ok returns a tuple with the CustomText13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText13

`func (o *Risks) SetCustomText13(v string)`

SetCustomText13 sets CustomText13 field to given value.

### HasCustomText13

`func (o *Risks) HasCustomText13() bool`

HasCustomText13 returns a boolean if a field has been set.

### GetCustomText14

`func (o *Risks) GetCustomText14() string`

GetCustomText14 returns the CustomText14 field if non-nil, zero value otherwise.

### GetCustomText14Ok

`func (o *Risks) GetCustomText14Ok() (*string, bool)`

GetCustomText14Ok returns a tuple with the CustomText14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText14

`func (o *Risks) SetCustomText14(v string)`

SetCustomText14 sets CustomText14 field to given value.

### HasCustomText14

`func (o *Risks) HasCustomText14() bool`

HasCustomText14 returns a boolean if a field has been set.

### GetCustomText15

`func (o *Risks) GetCustomText15() string`

GetCustomText15 returns the CustomText15 field if non-nil, zero value otherwise.

### GetCustomText15Ok

`func (o *Risks) GetCustomText15Ok() (*string, bool)`

GetCustomText15Ok returns a tuple with the CustomText15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText15

`func (o *Risks) SetCustomText15(v string)`

SetCustomText15 sets CustomText15 field to given value.

### HasCustomText15

`func (o *Risks) HasCustomText15() bool`

HasCustomText15 returns a boolean if a field has been set.

### GetCustomText16

`func (o *Risks) GetCustomText16() string`

GetCustomText16 returns the CustomText16 field if non-nil, zero value otherwise.

### GetCustomText16Ok

`func (o *Risks) GetCustomText16Ok() (*string, bool)`

GetCustomText16Ok returns a tuple with the CustomText16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText16

`func (o *Risks) SetCustomText16(v string)`

SetCustomText16 sets CustomText16 field to given value.

### HasCustomText16

`func (o *Risks) HasCustomText16() bool`

HasCustomText16 returns a boolean if a field has been set.

### GetCustomText17

`func (o *Risks) GetCustomText17() string`

GetCustomText17 returns the CustomText17 field if non-nil, zero value otherwise.

### GetCustomText17Ok

`func (o *Risks) GetCustomText17Ok() (*string, bool)`

GetCustomText17Ok returns a tuple with the CustomText17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText17

`func (o *Risks) SetCustomText17(v string)`

SetCustomText17 sets CustomText17 field to given value.

### HasCustomText17

`func (o *Risks) HasCustomText17() bool`

HasCustomText17 returns a boolean if a field has been set.

### GetCustomText18

`func (o *Risks) GetCustomText18() string`

GetCustomText18 returns the CustomText18 field if non-nil, zero value otherwise.

### GetCustomText18Ok

`func (o *Risks) GetCustomText18Ok() (*string, bool)`

GetCustomText18Ok returns a tuple with the CustomText18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText18

`func (o *Risks) SetCustomText18(v string)`

SetCustomText18 sets CustomText18 field to given value.

### HasCustomText18

`func (o *Risks) HasCustomText18() bool`

HasCustomText18 returns a boolean if a field has been set.

### GetCustomText19

`func (o *Risks) GetCustomText19() string`

GetCustomText19 returns the CustomText19 field if non-nil, zero value otherwise.

### GetCustomText19Ok

`func (o *Risks) GetCustomText19Ok() (*string, bool)`

GetCustomText19Ok returns a tuple with the CustomText19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText19

`func (o *Risks) SetCustomText19(v string)`

SetCustomText19 sets CustomText19 field to given value.

### HasCustomText19

`func (o *Risks) HasCustomText19() bool`

HasCustomText19 returns a boolean if a field has been set.

### GetCustomText20

`func (o *Risks) GetCustomText20() string`

GetCustomText20 returns the CustomText20 field if non-nil, zero value otherwise.

### GetCustomText20Ok

`func (o *Risks) GetCustomText20Ok() (*string, bool)`

GetCustomText20Ok returns a tuple with the CustomText20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText20

`func (o *Risks) SetCustomText20(v string)`

SetCustomText20 sets CustomText20 field to given value.

### HasCustomText20

`func (o *Risks) HasCustomText20() bool`

HasCustomText20 returns a boolean if a field has been set.

### GetCustomText21

`func (o *Risks) GetCustomText21() string`

GetCustomText21 returns the CustomText21 field if non-nil, zero value otherwise.

### GetCustomText21Ok

`func (o *Risks) GetCustomText21Ok() (*string, bool)`

GetCustomText21Ok returns a tuple with the CustomText21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText21

`func (o *Risks) SetCustomText21(v string)`

SetCustomText21 sets CustomText21 field to given value.

### HasCustomText21

`func (o *Risks) HasCustomText21() bool`

HasCustomText21 returns a boolean if a field has been set.

### GetCustomText22

`func (o *Risks) GetCustomText22() string`

GetCustomText22 returns the CustomText22 field if non-nil, zero value otherwise.

### GetCustomText22Ok

`func (o *Risks) GetCustomText22Ok() (*string, bool)`

GetCustomText22Ok returns a tuple with the CustomText22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText22

`func (o *Risks) SetCustomText22(v string)`

SetCustomText22 sets CustomText22 field to given value.

### HasCustomText22

`func (o *Risks) HasCustomText22() bool`

HasCustomText22 returns a boolean if a field has been set.

### GetCustomText23

`func (o *Risks) GetCustomText23() string`

GetCustomText23 returns the CustomText23 field if non-nil, zero value otherwise.

### GetCustomText23Ok

`func (o *Risks) GetCustomText23Ok() (*string, bool)`

GetCustomText23Ok returns a tuple with the CustomText23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText23

`func (o *Risks) SetCustomText23(v string)`

SetCustomText23 sets CustomText23 field to given value.

### HasCustomText23

`func (o *Risks) HasCustomText23() bool`

HasCustomText23 returns a boolean if a field has been set.

### GetCustomText24

`func (o *Risks) GetCustomText24() string`

GetCustomText24 returns the CustomText24 field if non-nil, zero value otherwise.

### GetCustomText24Ok

`func (o *Risks) GetCustomText24Ok() (*string, bool)`

GetCustomText24Ok returns a tuple with the CustomText24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText24

`func (o *Risks) SetCustomText24(v string)`

SetCustomText24 sets CustomText24 field to given value.

### HasCustomText24

`func (o *Risks) HasCustomText24() bool`

HasCustomText24 returns a boolean if a field has been set.

### GetCustomText25

`func (o *Risks) GetCustomText25() string`

GetCustomText25 returns the CustomText25 field if non-nil, zero value otherwise.

### GetCustomText25Ok

`func (o *Risks) GetCustomText25Ok() (*string, bool)`

GetCustomText25Ok returns a tuple with the CustomText25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText25

`func (o *Risks) SetCustomText25(v string)`

SetCustomText25 sets CustomText25 field to given value.

### HasCustomText25

`func (o *Risks) HasCustomText25() bool`

HasCustomText25 returns a boolean if a field has been set.

### GetRiskCustomSelect13OptionId

`func (o *Risks) GetRiskCustomSelect13OptionId() int32`

GetRiskCustomSelect13OptionId returns the RiskCustomSelect13OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect13OptionIdOk

`func (o *Risks) GetRiskCustomSelect13OptionIdOk() (*int32, bool)`

GetRiskCustomSelect13OptionIdOk returns a tuple with the RiskCustomSelect13OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect13OptionId

`func (o *Risks) SetRiskCustomSelect13OptionId(v int32)`

SetRiskCustomSelect13OptionId sets RiskCustomSelect13OptionId field to given value.

### HasRiskCustomSelect13OptionId

`func (o *Risks) HasRiskCustomSelect13OptionId() bool`

HasRiskCustomSelect13OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect14OptionId

`func (o *Risks) GetRiskCustomSelect14OptionId() int32`

GetRiskCustomSelect14OptionId returns the RiskCustomSelect14OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect14OptionIdOk

`func (o *Risks) GetRiskCustomSelect14OptionIdOk() (*int32, bool)`

GetRiskCustomSelect14OptionIdOk returns a tuple with the RiskCustomSelect14OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect14OptionId

`func (o *Risks) SetRiskCustomSelect14OptionId(v int32)`

SetRiskCustomSelect14OptionId sets RiskCustomSelect14OptionId field to given value.

### HasRiskCustomSelect14OptionId

`func (o *Risks) HasRiskCustomSelect14OptionId() bool`

HasRiskCustomSelect14OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect15OptionId

`func (o *Risks) GetRiskCustomSelect15OptionId() int32`

GetRiskCustomSelect15OptionId returns the RiskCustomSelect15OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect15OptionIdOk

`func (o *Risks) GetRiskCustomSelect15OptionIdOk() (*int32, bool)`

GetRiskCustomSelect15OptionIdOk returns a tuple with the RiskCustomSelect15OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect15OptionId

`func (o *Risks) SetRiskCustomSelect15OptionId(v int32)`

SetRiskCustomSelect15OptionId sets RiskCustomSelect15OptionId field to given value.

### HasRiskCustomSelect15OptionId

`func (o *Risks) HasRiskCustomSelect15OptionId() bool`

HasRiskCustomSelect15OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect16OptionId

`func (o *Risks) GetRiskCustomSelect16OptionId() int32`

GetRiskCustomSelect16OptionId returns the RiskCustomSelect16OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect16OptionIdOk

`func (o *Risks) GetRiskCustomSelect16OptionIdOk() (*int32, bool)`

GetRiskCustomSelect16OptionIdOk returns a tuple with the RiskCustomSelect16OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect16OptionId

`func (o *Risks) SetRiskCustomSelect16OptionId(v int32)`

SetRiskCustomSelect16OptionId sets RiskCustomSelect16OptionId field to given value.

### HasRiskCustomSelect16OptionId

`func (o *Risks) HasRiskCustomSelect16OptionId() bool`

HasRiskCustomSelect16OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect17OptionId

`func (o *Risks) GetRiskCustomSelect17OptionId() int32`

GetRiskCustomSelect17OptionId returns the RiskCustomSelect17OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect17OptionIdOk

`func (o *Risks) GetRiskCustomSelect17OptionIdOk() (*int32, bool)`

GetRiskCustomSelect17OptionIdOk returns a tuple with the RiskCustomSelect17OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect17OptionId

`func (o *Risks) SetRiskCustomSelect17OptionId(v int32)`

SetRiskCustomSelect17OptionId sets RiskCustomSelect17OptionId field to given value.

### HasRiskCustomSelect17OptionId

`func (o *Risks) HasRiskCustomSelect17OptionId() bool`

HasRiskCustomSelect17OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect18OptionId

`func (o *Risks) GetRiskCustomSelect18OptionId() int32`

GetRiskCustomSelect18OptionId returns the RiskCustomSelect18OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect18OptionIdOk

`func (o *Risks) GetRiskCustomSelect18OptionIdOk() (*int32, bool)`

GetRiskCustomSelect18OptionIdOk returns a tuple with the RiskCustomSelect18OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect18OptionId

`func (o *Risks) SetRiskCustomSelect18OptionId(v int32)`

SetRiskCustomSelect18OptionId sets RiskCustomSelect18OptionId field to given value.

### HasRiskCustomSelect18OptionId

`func (o *Risks) HasRiskCustomSelect18OptionId() bool`

HasRiskCustomSelect18OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect19OptionId

`func (o *Risks) GetRiskCustomSelect19OptionId() int32`

GetRiskCustomSelect19OptionId returns the RiskCustomSelect19OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect19OptionIdOk

`func (o *Risks) GetRiskCustomSelect19OptionIdOk() (*int32, bool)`

GetRiskCustomSelect19OptionIdOk returns a tuple with the RiskCustomSelect19OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect19OptionId

`func (o *Risks) SetRiskCustomSelect19OptionId(v int32)`

SetRiskCustomSelect19OptionId sets RiskCustomSelect19OptionId field to given value.

### HasRiskCustomSelect19OptionId

`func (o *Risks) HasRiskCustomSelect19OptionId() bool`

HasRiskCustomSelect19OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect20OptionId

`func (o *Risks) GetRiskCustomSelect20OptionId() int32`

GetRiskCustomSelect20OptionId returns the RiskCustomSelect20OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect20OptionIdOk

`func (o *Risks) GetRiskCustomSelect20OptionIdOk() (*int32, bool)`

GetRiskCustomSelect20OptionIdOk returns a tuple with the RiskCustomSelect20OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect20OptionId

`func (o *Risks) SetRiskCustomSelect20OptionId(v int32)`

SetRiskCustomSelect20OptionId sets RiskCustomSelect20OptionId field to given value.

### HasRiskCustomSelect20OptionId

`func (o *Risks) HasRiskCustomSelect20OptionId() bool`

HasRiskCustomSelect20OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect21OptionId

`func (o *Risks) GetRiskCustomSelect21OptionId() int32`

GetRiskCustomSelect21OptionId returns the RiskCustomSelect21OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect21OptionIdOk

`func (o *Risks) GetRiskCustomSelect21OptionIdOk() (*int32, bool)`

GetRiskCustomSelect21OptionIdOk returns a tuple with the RiskCustomSelect21OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect21OptionId

`func (o *Risks) SetRiskCustomSelect21OptionId(v int32)`

SetRiskCustomSelect21OptionId sets RiskCustomSelect21OptionId field to given value.

### HasRiskCustomSelect21OptionId

`func (o *Risks) HasRiskCustomSelect21OptionId() bool`

HasRiskCustomSelect21OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect22OptionId

`func (o *Risks) GetRiskCustomSelect22OptionId() int32`

GetRiskCustomSelect22OptionId returns the RiskCustomSelect22OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect22OptionIdOk

`func (o *Risks) GetRiskCustomSelect22OptionIdOk() (*int32, bool)`

GetRiskCustomSelect22OptionIdOk returns a tuple with the RiskCustomSelect22OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect22OptionId

`func (o *Risks) SetRiskCustomSelect22OptionId(v int32)`

SetRiskCustomSelect22OptionId sets RiskCustomSelect22OptionId field to given value.

### HasRiskCustomSelect22OptionId

`func (o *Risks) HasRiskCustomSelect22OptionId() bool`

HasRiskCustomSelect22OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect23OptionId

`func (o *Risks) GetRiskCustomSelect23OptionId() int32`

GetRiskCustomSelect23OptionId returns the RiskCustomSelect23OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect23OptionIdOk

`func (o *Risks) GetRiskCustomSelect23OptionIdOk() (*int32, bool)`

GetRiskCustomSelect23OptionIdOk returns a tuple with the RiskCustomSelect23OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect23OptionId

`func (o *Risks) SetRiskCustomSelect23OptionId(v int32)`

SetRiskCustomSelect23OptionId sets RiskCustomSelect23OptionId field to given value.

### HasRiskCustomSelect23OptionId

`func (o *Risks) HasRiskCustomSelect23OptionId() bool`

HasRiskCustomSelect23OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect24OptionId

`func (o *Risks) GetRiskCustomSelect24OptionId() int32`

GetRiskCustomSelect24OptionId returns the RiskCustomSelect24OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect24OptionIdOk

`func (o *Risks) GetRiskCustomSelect24OptionIdOk() (*int32, bool)`

GetRiskCustomSelect24OptionIdOk returns a tuple with the RiskCustomSelect24OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect24OptionId

`func (o *Risks) SetRiskCustomSelect24OptionId(v int32)`

SetRiskCustomSelect24OptionId sets RiskCustomSelect24OptionId field to given value.

### HasRiskCustomSelect24OptionId

`func (o *Risks) HasRiskCustomSelect24OptionId() bool`

HasRiskCustomSelect24OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect25OptionId

`func (o *Risks) GetRiskCustomSelect25OptionId() int32`

GetRiskCustomSelect25OptionId returns the RiskCustomSelect25OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect25OptionIdOk

`func (o *Risks) GetRiskCustomSelect25OptionIdOk() (*int32, bool)`

GetRiskCustomSelect25OptionIdOk returns a tuple with the RiskCustomSelect25OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect25OptionId

`func (o *Risks) SetRiskCustomSelect25OptionId(v int32)`

SetRiskCustomSelect25OptionId sets RiskCustomSelect25OptionId field to given value.

### HasRiskCustomSelect25OptionId

`func (o *Risks) HasRiskCustomSelect25OptionId() bool`

HasRiskCustomSelect25OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect26OptionId

`func (o *Risks) GetRiskCustomSelect26OptionId() int32`

GetRiskCustomSelect26OptionId returns the RiskCustomSelect26OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect26OptionIdOk

`func (o *Risks) GetRiskCustomSelect26OptionIdOk() (*int32, bool)`

GetRiskCustomSelect26OptionIdOk returns a tuple with the RiskCustomSelect26OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect26OptionId

`func (o *Risks) SetRiskCustomSelect26OptionId(v int32)`

SetRiskCustomSelect26OptionId sets RiskCustomSelect26OptionId field to given value.

### HasRiskCustomSelect26OptionId

`func (o *Risks) HasRiskCustomSelect26OptionId() bool`

HasRiskCustomSelect26OptionId returns a boolean if a field has been set.

### GetRiskCustomSelect27OptionId

`func (o *Risks) GetRiskCustomSelect27OptionId() int32`

GetRiskCustomSelect27OptionId returns the RiskCustomSelect27OptionId field if non-nil, zero value otherwise.

### GetRiskCustomSelect27OptionIdOk

`func (o *Risks) GetRiskCustomSelect27OptionIdOk() (*int32, bool)`

GetRiskCustomSelect27OptionIdOk returns a tuple with the RiskCustomSelect27OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCustomSelect27OptionId

`func (o *Risks) SetRiskCustomSelect27OptionId(v int32)`

SetRiskCustomSelect27OptionId sets RiskCustomSelect27OptionId field to given value.

### HasRiskCustomSelect27OptionId

`func (o *Risks) HasRiskCustomSelect27OptionId() bool`

HasRiskCustomSelect27OptionId returns a boolean if a field has been set.

### GetRiskAppetiteId

`func (o *Risks) GetRiskAppetiteId() int32`

GetRiskAppetiteId returns the RiskAppetiteId field if non-nil, zero value otherwise.

### GetRiskAppetiteIdOk

`func (o *Risks) GetRiskAppetiteIdOk() (*int32, bool)`

GetRiskAppetiteIdOk returns a tuple with the RiskAppetiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAppetiteId

`func (o *Risks) SetRiskAppetiteId(v int32)`

SetRiskAppetiteId sets RiskAppetiteId field to given value.

### HasRiskAppetiteId

`func (o *Risks) HasRiskAppetiteId() bool`

HasRiskAppetiteId returns a boolean if a field has been set.

### GetRiskAppetiteScore

`func (o *Risks) GetRiskAppetiteScore() interface{}`

GetRiskAppetiteScore returns the RiskAppetiteScore field if non-nil, zero value otherwise.

### GetRiskAppetiteScoreOk

`func (o *Risks) GetRiskAppetiteScoreOk() (*interface{}, bool)`

GetRiskAppetiteScoreOk returns a tuple with the RiskAppetiteScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAppetiteScore

`func (o *Risks) SetRiskAppetiteScore(v interface{})`

SetRiskAppetiteScore sets RiskAppetiteScore field to given value.

### HasRiskAppetiteScore

`func (o *Risks) HasRiskAppetiteScore() bool`

HasRiskAppetiteScore returns a boolean if a field has been set.

### SetRiskAppetiteScoreNil

`func (o *Risks) SetRiskAppetiteScoreNil(b bool)`

 SetRiskAppetiteScoreNil sets the value for RiskAppetiteScore to be an explicit nil

### UnsetRiskAppetiteScore
`func (o *Risks) UnsetRiskAppetiteScore()`

UnsetRiskAppetiteScore ensures that no value is present for RiskAppetiteScore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


