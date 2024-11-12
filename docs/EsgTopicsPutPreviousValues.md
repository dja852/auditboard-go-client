# EsgTopicsPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**FieldData** | Pointer to **interface{}** |  | [optional] 
**ReviewerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**EffectiveDate** | Pointer to **string** |  | [optional] 
**OrganizationImpact** | Pointer to **interface{}** |  | [optional] 
**StakeholderImpact** | Pointer to **interface{}** |  | [optional] 
**MaterialityRating** | Pointer to **interface{}** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] [default to "Not In Scope"]
**EsgCategoryId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_categories.id&#x60;.&lt;fk table&#x3D;&#39;esg_categories&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**EsgTopicCustomSelect1OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_topic_custom_select1_options.id&#x60;.&lt;fk table&#x3D;&#39;esg_topic_custom_select1_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**EsgTopicCustomSelect2OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_topic_custom_select2_options.id&#x60;.&lt;fk table&#x3D;&#39;esg_topic_custom_select2_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**EsgTopicCustomSelect3OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_topic_custom_select3_options.id&#x60;.&lt;fk table&#x3D;&#39;esg_topic_custom_select3_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**EsgTopicCustomSelect4OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_topic_custom_select4_options.id&#x60;.&lt;fk table&#x3D;&#39;esg_topic_custom_select4_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**EsgTopicCustomSelect5OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_topic_custom_select5_options.id&#x60;.&lt;fk table&#x3D;&#39;esg_topic_custom_select5_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomDate1** | Pointer to **string** |  | [optional] 
**CustomDate2** | Pointer to **string** |  | [optional] 
**CustomDate3** | Pointer to **string** |  | [optional] 
**CustomText1** | Pointer to **string** |  | [optional] 
**CustomText2** | Pointer to **string** |  | [optional] 
**CustomText3** | Pointer to **string** |  | [optional] 
**CustomText4** | Pointer to **string** |  | [optional] 
**CustomText5** | Pointer to **string** |  | [optional] 
**CustomText6** | Pointer to **string** |  | [optional] 

## Methods

### NewEsgTopicsPutPreviousValues

`func NewEsgTopicsPutPreviousValues() *EsgTopicsPutPreviousValues`

NewEsgTopicsPutPreviousValues instantiates a new EsgTopicsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsgTopicsPutPreviousValuesWithDefaults

`func NewEsgTopicsPutPreviousValuesWithDefaults() *EsgTopicsPutPreviousValues`

NewEsgTopicsPutPreviousValuesWithDefaults instantiates a new EsgTopicsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EsgTopicsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsgTopicsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsgTopicsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EsgTopicsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EsgTopicsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EsgTopicsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EsgTopicsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EsgTopicsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EsgTopicsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EsgTopicsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EsgTopicsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EsgTopicsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EsgTopicsPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EsgTopicsPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EsgTopicsPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EsgTopicsPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *EsgTopicsPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EsgTopicsPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EsgTopicsPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EsgTopicsPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *EsgTopicsPutPreviousValues) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EsgTopicsPutPreviousValues) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EsgTopicsPutPreviousValues) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EsgTopicsPutPreviousValues) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUid

`func (o *EsgTopicsPutPreviousValues) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *EsgTopicsPutPreviousValues) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *EsgTopicsPutPreviousValues) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *EsgTopicsPutPreviousValues) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetFieldData

`func (o *EsgTopicsPutPreviousValues) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *EsgTopicsPutPreviousValues) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *EsgTopicsPutPreviousValues) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *EsgTopicsPutPreviousValues) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *EsgTopicsPutPreviousValues) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *EsgTopicsPutPreviousValues) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil
### GetReviewerUserId

`func (o *EsgTopicsPutPreviousValues) GetReviewerUserId() int32`

GetReviewerUserId returns the ReviewerUserId field if non-nil, zero value otherwise.

### GetReviewerUserIdOk

`func (o *EsgTopicsPutPreviousValues) GetReviewerUserIdOk() (*int32, bool)`

GetReviewerUserIdOk returns a tuple with the ReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerUserId

`func (o *EsgTopicsPutPreviousValues) SetReviewerUserId(v int32)`

SetReviewerUserId sets ReviewerUserId field to given value.

### HasReviewerUserId

`func (o *EsgTopicsPutPreviousValues) HasReviewerUserId() bool`

HasReviewerUserId returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *EsgTopicsPutPreviousValues) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *EsgTopicsPutPreviousValues) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *EsgTopicsPutPreviousValues) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *EsgTopicsPutPreviousValues) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetOrganizationImpact

`func (o *EsgTopicsPutPreviousValues) GetOrganizationImpact() interface{}`

GetOrganizationImpact returns the OrganizationImpact field if non-nil, zero value otherwise.

### GetOrganizationImpactOk

`func (o *EsgTopicsPutPreviousValues) GetOrganizationImpactOk() (*interface{}, bool)`

GetOrganizationImpactOk returns a tuple with the OrganizationImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationImpact

`func (o *EsgTopicsPutPreviousValues) SetOrganizationImpact(v interface{})`

SetOrganizationImpact sets OrganizationImpact field to given value.

### HasOrganizationImpact

`func (o *EsgTopicsPutPreviousValues) HasOrganizationImpact() bool`

HasOrganizationImpact returns a boolean if a field has been set.

### SetOrganizationImpactNil

`func (o *EsgTopicsPutPreviousValues) SetOrganizationImpactNil(b bool)`

 SetOrganizationImpactNil sets the value for OrganizationImpact to be an explicit nil

### UnsetOrganizationImpact
`func (o *EsgTopicsPutPreviousValues) UnsetOrganizationImpact()`

UnsetOrganizationImpact ensures that no value is present for OrganizationImpact, not even an explicit nil
### GetStakeholderImpact

`func (o *EsgTopicsPutPreviousValues) GetStakeholderImpact() interface{}`

GetStakeholderImpact returns the StakeholderImpact field if non-nil, zero value otherwise.

### GetStakeholderImpactOk

`func (o *EsgTopicsPutPreviousValues) GetStakeholderImpactOk() (*interface{}, bool)`

GetStakeholderImpactOk returns a tuple with the StakeholderImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderImpact

`func (o *EsgTopicsPutPreviousValues) SetStakeholderImpact(v interface{})`

SetStakeholderImpact sets StakeholderImpact field to given value.

### HasStakeholderImpact

`func (o *EsgTopicsPutPreviousValues) HasStakeholderImpact() bool`

HasStakeholderImpact returns a boolean if a field has been set.

### SetStakeholderImpactNil

`func (o *EsgTopicsPutPreviousValues) SetStakeholderImpactNil(b bool)`

 SetStakeholderImpactNil sets the value for StakeholderImpact to be an explicit nil

### UnsetStakeholderImpact
`func (o *EsgTopicsPutPreviousValues) UnsetStakeholderImpact()`

UnsetStakeholderImpact ensures that no value is present for StakeholderImpact, not even an explicit nil
### GetMaterialityRating

`func (o *EsgTopicsPutPreviousValues) GetMaterialityRating() interface{}`

GetMaterialityRating returns the MaterialityRating field if non-nil, zero value otherwise.

### GetMaterialityRatingOk

`func (o *EsgTopicsPutPreviousValues) GetMaterialityRatingOk() (*interface{}, bool)`

GetMaterialityRatingOk returns a tuple with the MaterialityRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterialityRating

`func (o *EsgTopicsPutPreviousValues) SetMaterialityRating(v interface{})`

SetMaterialityRating sets MaterialityRating field to given value.

### HasMaterialityRating

`func (o *EsgTopicsPutPreviousValues) HasMaterialityRating() bool`

HasMaterialityRating returns a boolean if a field has been set.

### SetMaterialityRatingNil

`func (o *EsgTopicsPutPreviousValues) SetMaterialityRatingNil(b bool)`

 SetMaterialityRatingNil sets the value for MaterialityRating to be an explicit nil

### UnsetMaterialityRating
`func (o *EsgTopicsPutPreviousValues) UnsetMaterialityRating()`

UnsetMaterialityRating ensures that no value is present for MaterialityRating, not even an explicit nil
### GetCategory

`func (o *EsgTopicsPutPreviousValues) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *EsgTopicsPutPreviousValues) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *EsgTopicsPutPreviousValues) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *EsgTopicsPutPreviousValues) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetScope

`func (o *EsgTopicsPutPreviousValues) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EsgTopicsPutPreviousValues) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EsgTopicsPutPreviousValues) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *EsgTopicsPutPreviousValues) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetEsgCategoryId

`func (o *EsgTopicsPutPreviousValues) GetEsgCategoryId() int32`

GetEsgCategoryId returns the EsgCategoryId field if non-nil, zero value otherwise.

### GetEsgCategoryIdOk

`func (o *EsgTopicsPutPreviousValues) GetEsgCategoryIdOk() (*int32, bool)`

GetEsgCategoryIdOk returns a tuple with the EsgCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgCategoryId

`func (o *EsgTopicsPutPreviousValues) SetEsgCategoryId(v int32)`

SetEsgCategoryId sets EsgCategoryId field to given value.

### HasEsgCategoryId

`func (o *EsgTopicsPutPreviousValues) HasEsgCategoryId() bool`

HasEsgCategoryId returns a boolean if a field has been set.

### GetEsgTopicCustomSelect1OptionId

`func (o *EsgTopicsPutPreviousValues) GetEsgTopicCustomSelect1OptionId() int32`

GetEsgTopicCustomSelect1OptionId returns the EsgTopicCustomSelect1OptionId field if non-nil, zero value otherwise.

### GetEsgTopicCustomSelect1OptionIdOk

`func (o *EsgTopicsPutPreviousValues) GetEsgTopicCustomSelect1OptionIdOk() (*int32, bool)`

GetEsgTopicCustomSelect1OptionIdOk returns a tuple with the EsgTopicCustomSelect1OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgTopicCustomSelect1OptionId

`func (o *EsgTopicsPutPreviousValues) SetEsgTopicCustomSelect1OptionId(v int32)`

SetEsgTopicCustomSelect1OptionId sets EsgTopicCustomSelect1OptionId field to given value.

### HasEsgTopicCustomSelect1OptionId

`func (o *EsgTopicsPutPreviousValues) HasEsgTopicCustomSelect1OptionId() bool`

HasEsgTopicCustomSelect1OptionId returns a boolean if a field has been set.

### GetEsgTopicCustomSelect2OptionId

`func (o *EsgTopicsPutPreviousValues) GetEsgTopicCustomSelect2OptionId() int32`

GetEsgTopicCustomSelect2OptionId returns the EsgTopicCustomSelect2OptionId field if non-nil, zero value otherwise.

### GetEsgTopicCustomSelect2OptionIdOk

`func (o *EsgTopicsPutPreviousValues) GetEsgTopicCustomSelect2OptionIdOk() (*int32, bool)`

GetEsgTopicCustomSelect2OptionIdOk returns a tuple with the EsgTopicCustomSelect2OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgTopicCustomSelect2OptionId

`func (o *EsgTopicsPutPreviousValues) SetEsgTopicCustomSelect2OptionId(v int32)`

SetEsgTopicCustomSelect2OptionId sets EsgTopicCustomSelect2OptionId field to given value.

### HasEsgTopicCustomSelect2OptionId

`func (o *EsgTopicsPutPreviousValues) HasEsgTopicCustomSelect2OptionId() bool`

HasEsgTopicCustomSelect2OptionId returns a boolean if a field has been set.

### GetEsgTopicCustomSelect3OptionId

`func (o *EsgTopicsPutPreviousValues) GetEsgTopicCustomSelect3OptionId() int32`

GetEsgTopicCustomSelect3OptionId returns the EsgTopicCustomSelect3OptionId field if non-nil, zero value otherwise.

### GetEsgTopicCustomSelect3OptionIdOk

`func (o *EsgTopicsPutPreviousValues) GetEsgTopicCustomSelect3OptionIdOk() (*int32, bool)`

GetEsgTopicCustomSelect3OptionIdOk returns a tuple with the EsgTopicCustomSelect3OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgTopicCustomSelect3OptionId

`func (o *EsgTopicsPutPreviousValues) SetEsgTopicCustomSelect3OptionId(v int32)`

SetEsgTopicCustomSelect3OptionId sets EsgTopicCustomSelect3OptionId field to given value.

### HasEsgTopicCustomSelect3OptionId

`func (o *EsgTopicsPutPreviousValues) HasEsgTopicCustomSelect3OptionId() bool`

HasEsgTopicCustomSelect3OptionId returns a boolean if a field has been set.

### GetEsgTopicCustomSelect4OptionId

`func (o *EsgTopicsPutPreviousValues) GetEsgTopicCustomSelect4OptionId() int32`

GetEsgTopicCustomSelect4OptionId returns the EsgTopicCustomSelect4OptionId field if non-nil, zero value otherwise.

### GetEsgTopicCustomSelect4OptionIdOk

`func (o *EsgTopicsPutPreviousValues) GetEsgTopicCustomSelect4OptionIdOk() (*int32, bool)`

GetEsgTopicCustomSelect4OptionIdOk returns a tuple with the EsgTopicCustomSelect4OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgTopicCustomSelect4OptionId

`func (o *EsgTopicsPutPreviousValues) SetEsgTopicCustomSelect4OptionId(v int32)`

SetEsgTopicCustomSelect4OptionId sets EsgTopicCustomSelect4OptionId field to given value.

### HasEsgTopicCustomSelect4OptionId

`func (o *EsgTopicsPutPreviousValues) HasEsgTopicCustomSelect4OptionId() bool`

HasEsgTopicCustomSelect4OptionId returns a boolean if a field has been set.

### GetEsgTopicCustomSelect5OptionId

`func (o *EsgTopicsPutPreviousValues) GetEsgTopicCustomSelect5OptionId() int32`

GetEsgTopicCustomSelect5OptionId returns the EsgTopicCustomSelect5OptionId field if non-nil, zero value otherwise.

### GetEsgTopicCustomSelect5OptionIdOk

`func (o *EsgTopicsPutPreviousValues) GetEsgTopicCustomSelect5OptionIdOk() (*int32, bool)`

GetEsgTopicCustomSelect5OptionIdOk returns a tuple with the EsgTopicCustomSelect5OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgTopicCustomSelect5OptionId

`func (o *EsgTopicsPutPreviousValues) SetEsgTopicCustomSelect5OptionId(v int32)`

SetEsgTopicCustomSelect5OptionId sets EsgTopicCustomSelect5OptionId field to given value.

### HasEsgTopicCustomSelect5OptionId

`func (o *EsgTopicsPutPreviousValues) HasEsgTopicCustomSelect5OptionId() bool`

HasEsgTopicCustomSelect5OptionId returns a boolean if a field has been set.

### GetCustomDate1

`func (o *EsgTopicsPutPreviousValues) GetCustomDate1() string`

GetCustomDate1 returns the CustomDate1 field if non-nil, zero value otherwise.

### GetCustomDate1Ok

`func (o *EsgTopicsPutPreviousValues) GetCustomDate1Ok() (*string, bool)`

GetCustomDate1Ok returns a tuple with the CustomDate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate1

`func (o *EsgTopicsPutPreviousValues) SetCustomDate1(v string)`

SetCustomDate1 sets CustomDate1 field to given value.

### HasCustomDate1

`func (o *EsgTopicsPutPreviousValues) HasCustomDate1() bool`

HasCustomDate1 returns a boolean if a field has been set.

### GetCustomDate2

`func (o *EsgTopicsPutPreviousValues) GetCustomDate2() string`

GetCustomDate2 returns the CustomDate2 field if non-nil, zero value otherwise.

### GetCustomDate2Ok

`func (o *EsgTopicsPutPreviousValues) GetCustomDate2Ok() (*string, bool)`

GetCustomDate2Ok returns a tuple with the CustomDate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate2

`func (o *EsgTopicsPutPreviousValues) SetCustomDate2(v string)`

SetCustomDate2 sets CustomDate2 field to given value.

### HasCustomDate2

`func (o *EsgTopicsPutPreviousValues) HasCustomDate2() bool`

HasCustomDate2 returns a boolean if a field has been set.

### GetCustomDate3

`func (o *EsgTopicsPutPreviousValues) GetCustomDate3() string`

GetCustomDate3 returns the CustomDate3 field if non-nil, zero value otherwise.

### GetCustomDate3Ok

`func (o *EsgTopicsPutPreviousValues) GetCustomDate3Ok() (*string, bool)`

GetCustomDate3Ok returns a tuple with the CustomDate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate3

`func (o *EsgTopicsPutPreviousValues) SetCustomDate3(v string)`

SetCustomDate3 sets CustomDate3 field to given value.

### HasCustomDate3

`func (o *EsgTopicsPutPreviousValues) HasCustomDate3() bool`

HasCustomDate3 returns a boolean if a field has been set.

### GetCustomText1

`func (o *EsgTopicsPutPreviousValues) GetCustomText1() string`

GetCustomText1 returns the CustomText1 field if non-nil, zero value otherwise.

### GetCustomText1Ok

`func (o *EsgTopicsPutPreviousValues) GetCustomText1Ok() (*string, bool)`

GetCustomText1Ok returns a tuple with the CustomText1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText1

`func (o *EsgTopicsPutPreviousValues) SetCustomText1(v string)`

SetCustomText1 sets CustomText1 field to given value.

### HasCustomText1

`func (o *EsgTopicsPutPreviousValues) HasCustomText1() bool`

HasCustomText1 returns a boolean if a field has been set.

### GetCustomText2

`func (o *EsgTopicsPutPreviousValues) GetCustomText2() string`

GetCustomText2 returns the CustomText2 field if non-nil, zero value otherwise.

### GetCustomText2Ok

`func (o *EsgTopicsPutPreviousValues) GetCustomText2Ok() (*string, bool)`

GetCustomText2Ok returns a tuple with the CustomText2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText2

`func (o *EsgTopicsPutPreviousValues) SetCustomText2(v string)`

SetCustomText2 sets CustomText2 field to given value.

### HasCustomText2

`func (o *EsgTopicsPutPreviousValues) HasCustomText2() bool`

HasCustomText2 returns a boolean if a field has been set.

### GetCustomText3

`func (o *EsgTopicsPutPreviousValues) GetCustomText3() string`

GetCustomText3 returns the CustomText3 field if non-nil, zero value otherwise.

### GetCustomText3Ok

`func (o *EsgTopicsPutPreviousValues) GetCustomText3Ok() (*string, bool)`

GetCustomText3Ok returns a tuple with the CustomText3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText3

`func (o *EsgTopicsPutPreviousValues) SetCustomText3(v string)`

SetCustomText3 sets CustomText3 field to given value.

### HasCustomText3

`func (o *EsgTopicsPutPreviousValues) HasCustomText3() bool`

HasCustomText3 returns a boolean if a field has been set.

### GetCustomText4

`func (o *EsgTopicsPutPreviousValues) GetCustomText4() string`

GetCustomText4 returns the CustomText4 field if non-nil, zero value otherwise.

### GetCustomText4Ok

`func (o *EsgTopicsPutPreviousValues) GetCustomText4Ok() (*string, bool)`

GetCustomText4Ok returns a tuple with the CustomText4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText4

`func (o *EsgTopicsPutPreviousValues) SetCustomText4(v string)`

SetCustomText4 sets CustomText4 field to given value.

### HasCustomText4

`func (o *EsgTopicsPutPreviousValues) HasCustomText4() bool`

HasCustomText4 returns a boolean if a field has been set.

### GetCustomText5

`func (o *EsgTopicsPutPreviousValues) GetCustomText5() string`

GetCustomText5 returns the CustomText5 field if non-nil, zero value otherwise.

### GetCustomText5Ok

`func (o *EsgTopicsPutPreviousValues) GetCustomText5Ok() (*string, bool)`

GetCustomText5Ok returns a tuple with the CustomText5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText5

`func (o *EsgTopicsPutPreviousValues) SetCustomText5(v string)`

SetCustomText5 sets CustomText5 field to given value.

### HasCustomText5

`func (o *EsgTopicsPutPreviousValues) HasCustomText5() bool`

HasCustomText5 returns a boolean if a field has been set.

### GetCustomText6

`func (o *EsgTopicsPutPreviousValues) GetCustomText6() string`

GetCustomText6 returns the CustomText6 field if non-nil, zero value otherwise.

### GetCustomText6Ok

`func (o *EsgTopicsPutPreviousValues) GetCustomText6Ok() (*string, bool)`

GetCustomText6Ok returns a tuple with the CustomText6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText6

`func (o *EsgTopicsPutPreviousValues) SetCustomText6(v string)`

SetCustomText6 sets CustomText6 field to given value.

### HasCustomText6

`func (o *EsgTopicsPutPreviousValues) HasCustomText6() bool`

HasCustomText6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


