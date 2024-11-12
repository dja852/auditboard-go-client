# KeyRiskIndicators

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomDate1** | Pointer to **string** |  | [optional] 
**LinkifyUid** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**PriorStatus** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**OwnerUserId** | Pointer to **int64** |  | [optional] 
**ExecutiveOwnerUserId** | Pointer to **int64** |  | [optional] 
**CustomUserSelect1UserId** | Pointer to **int64** |  | [optional] 
**Formula** | Pointer to **string** |  | [optional] 
**FormulaData** | Pointer to **map[string]interface{}** |  | [optional] 
**Brackets** | Pointer to **[]int32** |  | [optional] 
**CurrentValue** | Pointer to **string** |  | [optional] 
**NormalizedCurrentValue** | Pointer to **string** |  | [optional] 
**PriorValue** | Pointer to **string** |  | [optional] 
**NormalizedPriorValue** | Pointer to **string** |  | [optional] 
**ValueLastUpdated** | Pointer to **string** |  | [optional] 
**KriCustomSelect1OptionId** | Pointer to **int32** |  | [optional] 
**KriCustomSelect2OptionId** | Pointer to **int32** |  | [optional] 
**KriCustomSelect3OptionId** | Pointer to **int32** |  | [optional] 
**KriCustomSelect4OptionId** | Pointer to **int32** |  | [optional] 
**KriCustomSelect5OptionId** | Pointer to **int32** |  | [optional] 
**CustomText1** | Pointer to **string** |  | [optional] 
**CustomText2** | Pointer to **string** |  | [optional] 
**CustomText3** | Pointer to **string** |  | [optional] 
**CustomText4** | Pointer to **string** |  | [optional] 
**CustomText5** | Pointer to **string** |  | [optional] 
**NeedsUpdate** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewKeyRiskIndicators

`func NewKeyRiskIndicators() *KeyRiskIndicators`

NewKeyRiskIndicators instantiates a new KeyRiskIndicators object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyRiskIndicatorsWithDefaults

`func NewKeyRiskIndicatorsWithDefaults() *KeyRiskIndicators`

NewKeyRiskIndicatorsWithDefaults instantiates a new KeyRiskIndicators object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomDate1

`func (o *KeyRiskIndicators) GetCustomDate1() string`

GetCustomDate1 returns the CustomDate1 field if non-nil, zero value otherwise.

### GetCustomDate1Ok

`func (o *KeyRiskIndicators) GetCustomDate1Ok() (*string, bool)`

GetCustomDate1Ok returns a tuple with the CustomDate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate1

`func (o *KeyRiskIndicators) SetCustomDate1(v string)`

SetCustomDate1 sets CustomDate1 field to given value.

### HasCustomDate1

`func (o *KeyRiskIndicators) HasCustomDate1() bool`

HasCustomDate1 returns a boolean if a field has been set.

### GetLinkifyUid

`func (o *KeyRiskIndicators) GetLinkifyUid() string`

GetLinkifyUid returns the LinkifyUid field if non-nil, zero value otherwise.

### GetLinkifyUidOk

`func (o *KeyRiskIndicators) GetLinkifyUidOk() (*string, bool)`

GetLinkifyUidOk returns a tuple with the LinkifyUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkifyUid

`func (o *KeyRiskIndicators) SetLinkifyUid(v string)`

SetLinkifyUid sets LinkifyUid field to given value.

### HasLinkifyUid

`func (o *KeyRiskIndicators) HasLinkifyUid() bool`

HasLinkifyUid returns a boolean if a field has been set.

### GetId

`func (o *KeyRiskIndicators) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyRiskIndicators) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyRiskIndicators) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KeyRiskIndicators) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *KeyRiskIndicators) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *KeyRiskIndicators) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *KeyRiskIndicators) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *KeyRiskIndicators) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetName

`func (o *KeyRiskIndicators) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyRiskIndicators) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyRiskIndicators) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyRiskIndicators) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *KeyRiskIndicators) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KeyRiskIndicators) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KeyRiskIndicators) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KeyRiskIndicators) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriorStatus

`func (o *KeyRiskIndicators) GetPriorStatus() string`

GetPriorStatus returns the PriorStatus field if non-nil, zero value otherwise.

### GetPriorStatusOk

`func (o *KeyRiskIndicators) GetPriorStatusOk() (*string, bool)`

GetPriorStatusOk returns a tuple with the PriorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorStatus

`func (o *KeyRiskIndicators) SetPriorStatus(v string)`

SetPriorStatus sets PriorStatus field to given value.

### HasPriorStatus

`func (o *KeyRiskIndicators) HasPriorStatus() bool`

HasPriorStatus returns a boolean if a field has been set.

### GetDescription

`func (o *KeyRiskIndicators) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KeyRiskIndicators) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KeyRiskIndicators) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KeyRiskIndicators) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwnerUserId

`func (o *KeyRiskIndicators) GetOwnerUserId() int64`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *KeyRiskIndicators) GetOwnerUserIdOk() (*int64, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *KeyRiskIndicators) SetOwnerUserId(v int64)`

SetOwnerUserId sets OwnerUserId field to given value.

### HasOwnerUserId

`func (o *KeyRiskIndicators) HasOwnerUserId() bool`

HasOwnerUserId returns a boolean if a field has been set.

### GetExecutiveOwnerUserId

`func (o *KeyRiskIndicators) GetExecutiveOwnerUserId() int64`

GetExecutiveOwnerUserId returns the ExecutiveOwnerUserId field if non-nil, zero value otherwise.

### GetExecutiveOwnerUserIdOk

`func (o *KeyRiskIndicators) GetExecutiveOwnerUserIdOk() (*int64, bool)`

GetExecutiveOwnerUserIdOk returns a tuple with the ExecutiveOwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutiveOwnerUserId

`func (o *KeyRiskIndicators) SetExecutiveOwnerUserId(v int64)`

SetExecutiveOwnerUserId sets ExecutiveOwnerUserId field to given value.

### HasExecutiveOwnerUserId

`func (o *KeyRiskIndicators) HasExecutiveOwnerUserId() bool`

HasExecutiveOwnerUserId returns a boolean if a field has been set.

### GetCustomUserSelect1UserId

`func (o *KeyRiskIndicators) GetCustomUserSelect1UserId() int64`

GetCustomUserSelect1UserId returns the CustomUserSelect1UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect1UserIdOk

`func (o *KeyRiskIndicators) GetCustomUserSelect1UserIdOk() (*int64, bool)`

GetCustomUserSelect1UserIdOk returns a tuple with the CustomUserSelect1UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect1UserId

`func (o *KeyRiskIndicators) SetCustomUserSelect1UserId(v int64)`

SetCustomUserSelect1UserId sets CustomUserSelect1UserId field to given value.

### HasCustomUserSelect1UserId

`func (o *KeyRiskIndicators) HasCustomUserSelect1UserId() bool`

HasCustomUserSelect1UserId returns a boolean if a field has been set.

### GetFormula

`func (o *KeyRiskIndicators) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *KeyRiskIndicators) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *KeyRiskIndicators) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *KeyRiskIndicators) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetFormulaData

`func (o *KeyRiskIndicators) GetFormulaData() map[string]interface{}`

GetFormulaData returns the FormulaData field if non-nil, zero value otherwise.

### GetFormulaDataOk

`func (o *KeyRiskIndicators) GetFormulaDataOk() (*map[string]interface{}, bool)`

GetFormulaDataOk returns a tuple with the FormulaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulaData

`func (o *KeyRiskIndicators) SetFormulaData(v map[string]interface{})`

SetFormulaData sets FormulaData field to given value.

### HasFormulaData

`func (o *KeyRiskIndicators) HasFormulaData() bool`

HasFormulaData returns a boolean if a field has been set.

### GetBrackets

`func (o *KeyRiskIndicators) GetBrackets() []int32`

GetBrackets returns the Brackets field if non-nil, zero value otherwise.

### GetBracketsOk

`func (o *KeyRiskIndicators) GetBracketsOk() (*[]int32, bool)`

GetBracketsOk returns a tuple with the Brackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrackets

`func (o *KeyRiskIndicators) SetBrackets(v []int32)`

SetBrackets sets Brackets field to given value.

### HasBrackets

`func (o *KeyRiskIndicators) HasBrackets() bool`

HasBrackets returns a boolean if a field has been set.

### GetCurrentValue

`func (o *KeyRiskIndicators) GetCurrentValue() string`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *KeyRiskIndicators) GetCurrentValueOk() (*string, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *KeyRiskIndicators) SetCurrentValue(v string)`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *KeyRiskIndicators) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### GetNormalizedCurrentValue

`func (o *KeyRiskIndicators) GetNormalizedCurrentValue() string`

GetNormalizedCurrentValue returns the NormalizedCurrentValue field if non-nil, zero value otherwise.

### GetNormalizedCurrentValueOk

`func (o *KeyRiskIndicators) GetNormalizedCurrentValueOk() (*string, bool)`

GetNormalizedCurrentValueOk returns a tuple with the NormalizedCurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedCurrentValue

`func (o *KeyRiskIndicators) SetNormalizedCurrentValue(v string)`

SetNormalizedCurrentValue sets NormalizedCurrentValue field to given value.

### HasNormalizedCurrentValue

`func (o *KeyRiskIndicators) HasNormalizedCurrentValue() bool`

HasNormalizedCurrentValue returns a boolean if a field has been set.

### GetPriorValue

`func (o *KeyRiskIndicators) GetPriorValue() string`

GetPriorValue returns the PriorValue field if non-nil, zero value otherwise.

### GetPriorValueOk

`func (o *KeyRiskIndicators) GetPriorValueOk() (*string, bool)`

GetPriorValueOk returns a tuple with the PriorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorValue

`func (o *KeyRiskIndicators) SetPriorValue(v string)`

SetPriorValue sets PriorValue field to given value.

### HasPriorValue

`func (o *KeyRiskIndicators) HasPriorValue() bool`

HasPriorValue returns a boolean if a field has been set.

### GetNormalizedPriorValue

`func (o *KeyRiskIndicators) GetNormalizedPriorValue() string`

GetNormalizedPriorValue returns the NormalizedPriorValue field if non-nil, zero value otherwise.

### GetNormalizedPriorValueOk

`func (o *KeyRiskIndicators) GetNormalizedPriorValueOk() (*string, bool)`

GetNormalizedPriorValueOk returns a tuple with the NormalizedPriorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedPriorValue

`func (o *KeyRiskIndicators) SetNormalizedPriorValue(v string)`

SetNormalizedPriorValue sets NormalizedPriorValue field to given value.

### HasNormalizedPriorValue

`func (o *KeyRiskIndicators) HasNormalizedPriorValue() bool`

HasNormalizedPriorValue returns a boolean if a field has been set.

### GetValueLastUpdated

`func (o *KeyRiskIndicators) GetValueLastUpdated() string`

GetValueLastUpdated returns the ValueLastUpdated field if non-nil, zero value otherwise.

### GetValueLastUpdatedOk

`func (o *KeyRiskIndicators) GetValueLastUpdatedOk() (*string, bool)`

GetValueLastUpdatedOk returns a tuple with the ValueLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueLastUpdated

`func (o *KeyRiskIndicators) SetValueLastUpdated(v string)`

SetValueLastUpdated sets ValueLastUpdated field to given value.

### HasValueLastUpdated

`func (o *KeyRiskIndicators) HasValueLastUpdated() bool`

HasValueLastUpdated returns a boolean if a field has been set.

### GetKriCustomSelect1OptionId

`func (o *KeyRiskIndicators) GetKriCustomSelect1OptionId() int32`

GetKriCustomSelect1OptionId returns the KriCustomSelect1OptionId field if non-nil, zero value otherwise.

### GetKriCustomSelect1OptionIdOk

`func (o *KeyRiskIndicators) GetKriCustomSelect1OptionIdOk() (*int32, bool)`

GetKriCustomSelect1OptionIdOk returns a tuple with the KriCustomSelect1OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKriCustomSelect1OptionId

`func (o *KeyRiskIndicators) SetKriCustomSelect1OptionId(v int32)`

SetKriCustomSelect1OptionId sets KriCustomSelect1OptionId field to given value.

### HasKriCustomSelect1OptionId

`func (o *KeyRiskIndicators) HasKriCustomSelect1OptionId() bool`

HasKriCustomSelect1OptionId returns a boolean if a field has been set.

### GetKriCustomSelect2OptionId

`func (o *KeyRiskIndicators) GetKriCustomSelect2OptionId() int32`

GetKriCustomSelect2OptionId returns the KriCustomSelect2OptionId field if non-nil, zero value otherwise.

### GetKriCustomSelect2OptionIdOk

`func (o *KeyRiskIndicators) GetKriCustomSelect2OptionIdOk() (*int32, bool)`

GetKriCustomSelect2OptionIdOk returns a tuple with the KriCustomSelect2OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKriCustomSelect2OptionId

`func (o *KeyRiskIndicators) SetKriCustomSelect2OptionId(v int32)`

SetKriCustomSelect2OptionId sets KriCustomSelect2OptionId field to given value.

### HasKriCustomSelect2OptionId

`func (o *KeyRiskIndicators) HasKriCustomSelect2OptionId() bool`

HasKriCustomSelect2OptionId returns a boolean if a field has been set.

### GetKriCustomSelect3OptionId

`func (o *KeyRiskIndicators) GetKriCustomSelect3OptionId() int32`

GetKriCustomSelect3OptionId returns the KriCustomSelect3OptionId field if non-nil, zero value otherwise.

### GetKriCustomSelect3OptionIdOk

`func (o *KeyRiskIndicators) GetKriCustomSelect3OptionIdOk() (*int32, bool)`

GetKriCustomSelect3OptionIdOk returns a tuple with the KriCustomSelect3OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKriCustomSelect3OptionId

`func (o *KeyRiskIndicators) SetKriCustomSelect3OptionId(v int32)`

SetKriCustomSelect3OptionId sets KriCustomSelect3OptionId field to given value.

### HasKriCustomSelect3OptionId

`func (o *KeyRiskIndicators) HasKriCustomSelect3OptionId() bool`

HasKriCustomSelect3OptionId returns a boolean if a field has been set.

### GetKriCustomSelect4OptionId

`func (o *KeyRiskIndicators) GetKriCustomSelect4OptionId() int32`

GetKriCustomSelect4OptionId returns the KriCustomSelect4OptionId field if non-nil, zero value otherwise.

### GetKriCustomSelect4OptionIdOk

`func (o *KeyRiskIndicators) GetKriCustomSelect4OptionIdOk() (*int32, bool)`

GetKriCustomSelect4OptionIdOk returns a tuple with the KriCustomSelect4OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKriCustomSelect4OptionId

`func (o *KeyRiskIndicators) SetKriCustomSelect4OptionId(v int32)`

SetKriCustomSelect4OptionId sets KriCustomSelect4OptionId field to given value.

### HasKriCustomSelect4OptionId

`func (o *KeyRiskIndicators) HasKriCustomSelect4OptionId() bool`

HasKriCustomSelect4OptionId returns a boolean if a field has been set.

### GetKriCustomSelect5OptionId

`func (o *KeyRiskIndicators) GetKriCustomSelect5OptionId() int32`

GetKriCustomSelect5OptionId returns the KriCustomSelect5OptionId field if non-nil, zero value otherwise.

### GetKriCustomSelect5OptionIdOk

`func (o *KeyRiskIndicators) GetKriCustomSelect5OptionIdOk() (*int32, bool)`

GetKriCustomSelect5OptionIdOk returns a tuple with the KriCustomSelect5OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKriCustomSelect5OptionId

`func (o *KeyRiskIndicators) SetKriCustomSelect5OptionId(v int32)`

SetKriCustomSelect5OptionId sets KriCustomSelect5OptionId field to given value.

### HasKriCustomSelect5OptionId

`func (o *KeyRiskIndicators) HasKriCustomSelect5OptionId() bool`

HasKriCustomSelect5OptionId returns a boolean if a field has been set.

### GetCustomText1

`func (o *KeyRiskIndicators) GetCustomText1() string`

GetCustomText1 returns the CustomText1 field if non-nil, zero value otherwise.

### GetCustomText1Ok

`func (o *KeyRiskIndicators) GetCustomText1Ok() (*string, bool)`

GetCustomText1Ok returns a tuple with the CustomText1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText1

`func (o *KeyRiskIndicators) SetCustomText1(v string)`

SetCustomText1 sets CustomText1 field to given value.

### HasCustomText1

`func (o *KeyRiskIndicators) HasCustomText1() bool`

HasCustomText1 returns a boolean if a field has been set.

### GetCustomText2

`func (o *KeyRiskIndicators) GetCustomText2() string`

GetCustomText2 returns the CustomText2 field if non-nil, zero value otherwise.

### GetCustomText2Ok

`func (o *KeyRiskIndicators) GetCustomText2Ok() (*string, bool)`

GetCustomText2Ok returns a tuple with the CustomText2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText2

`func (o *KeyRiskIndicators) SetCustomText2(v string)`

SetCustomText2 sets CustomText2 field to given value.

### HasCustomText2

`func (o *KeyRiskIndicators) HasCustomText2() bool`

HasCustomText2 returns a boolean if a field has been set.

### GetCustomText3

`func (o *KeyRiskIndicators) GetCustomText3() string`

GetCustomText3 returns the CustomText3 field if non-nil, zero value otherwise.

### GetCustomText3Ok

`func (o *KeyRiskIndicators) GetCustomText3Ok() (*string, bool)`

GetCustomText3Ok returns a tuple with the CustomText3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText3

`func (o *KeyRiskIndicators) SetCustomText3(v string)`

SetCustomText3 sets CustomText3 field to given value.

### HasCustomText3

`func (o *KeyRiskIndicators) HasCustomText3() bool`

HasCustomText3 returns a boolean if a field has been set.

### GetCustomText4

`func (o *KeyRiskIndicators) GetCustomText4() string`

GetCustomText4 returns the CustomText4 field if non-nil, zero value otherwise.

### GetCustomText4Ok

`func (o *KeyRiskIndicators) GetCustomText4Ok() (*string, bool)`

GetCustomText4Ok returns a tuple with the CustomText4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText4

`func (o *KeyRiskIndicators) SetCustomText4(v string)`

SetCustomText4 sets CustomText4 field to given value.

### HasCustomText4

`func (o *KeyRiskIndicators) HasCustomText4() bool`

HasCustomText4 returns a boolean if a field has been set.

### GetCustomText5

`func (o *KeyRiskIndicators) GetCustomText5() string`

GetCustomText5 returns the CustomText5 field if non-nil, zero value otherwise.

### GetCustomText5Ok

`func (o *KeyRiskIndicators) GetCustomText5Ok() (*string, bool)`

GetCustomText5Ok returns a tuple with the CustomText5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText5

`func (o *KeyRiskIndicators) SetCustomText5(v string)`

SetCustomText5 sets CustomText5 field to given value.

### HasCustomText5

`func (o *KeyRiskIndicators) HasCustomText5() bool`

HasCustomText5 returns a boolean if a field has been set.

### GetNeedsUpdate

`func (o *KeyRiskIndicators) GetNeedsUpdate() bool`

GetNeedsUpdate returns the NeedsUpdate field if non-nil, zero value otherwise.

### GetNeedsUpdateOk

`func (o *KeyRiskIndicators) GetNeedsUpdateOk() (*bool, bool)`

GetNeedsUpdateOk returns a tuple with the NeedsUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsUpdate

`func (o *KeyRiskIndicators) SetNeedsUpdate(v bool)`

SetNeedsUpdate sets NeedsUpdate field to given value.

### HasNeedsUpdate

`func (o *KeyRiskIndicators) HasNeedsUpdate() bool`

HasNeedsUpdate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KeyRiskIndicators) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KeyRiskIndicators) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KeyRiskIndicators) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KeyRiskIndicators) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KeyRiskIndicators) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KeyRiskIndicators) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KeyRiskIndicators) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KeyRiskIndicators) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *KeyRiskIndicators) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *KeyRiskIndicators) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *KeyRiskIndicators) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *KeyRiskIndicators) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetPermissions

`func (o *KeyRiskIndicators) GetPermissions() map[string]interface{}`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *KeyRiskIndicators) GetPermissionsOk() (*map[string]interface{}, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *KeyRiskIndicators) SetPermissions(v map[string]interface{})`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *KeyRiskIndicators) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


