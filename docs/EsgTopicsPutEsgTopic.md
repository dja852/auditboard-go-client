# EsgTopicsPutEsgTopic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
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

### NewEsgTopicsPutEsgTopic

`func NewEsgTopicsPutEsgTopic(name string, ) *EsgTopicsPutEsgTopic`

NewEsgTopicsPutEsgTopic instantiates a new EsgTopicsPutEsgTopic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsgTopicsPutEsgTopicWithDefaults

`func NewEsgTopicsPutEsgTopicWithDefaults() *EsgTopicsPutEsgTopic`

NewEsgTopicsPutEsgTopicWithDefaults instantiates a new EsgTopicsPutEsgTopic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EsgTopicsPutEsgTopic) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsgTopicsPutEsgTopic) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsgTopicsPutEsgTopic) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EsgTopicsPutEsgTopic) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EsgTopicsPutEsgTopic) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EsgTopicsPutEsgTopic) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EsgTopicsPutEsgTopic) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EsgTopicsPutEsgTopic) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EsgTopicsPutEsgTopic) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EsgTopicsPutEsgTopic) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EsgTopicsPutEsgTopic) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EsgTopicsPutEsgTopic) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EsgTopicsPutEsgTopic) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EsgTopicsPutEsgTopic) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EsgTopicsPutEsgTopic) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EsgTopicsPutEsgTopic) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *EsgTopicsPutEsgTopic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EsgTopicsPutEsgTopic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EsgTopicsPutEsgTopic) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EsgTopicsPutEsgTopic) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EsgTopicsPutEsgTopic) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EsgTopicsPutEsgTopic) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EsgTopicsPutEsgTopic) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUid

`func (o *EsgTopicsPutEsgTopic) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *EsgTopicsPutEsgTopic) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *EsgTopicsPutEsgTopic) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *EsgTopicsPutEsgTopic) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetFieldData

`func (o *EsgTopicsPutEsgTopic) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *EsgTopicsPutEsgTopic) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *EsgTopicsPutEsgTopic) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *EsgTopicsPutEsgTopic) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *EsgTopicsPutEsgTopic) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *EsgTopicsPutEsgTopic) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil
### GetReviewerUserId

`func (o *EsgTopicsPutEsgTopic) GetReviewerUserId() int32`

GetReviewerUserId returns the ReviewerUserId field if non-nil, zero value otherwise.

### GetReviewerUserIdOk

`func (o *EsgTopicsPutEsgTopic) GetReviewerUserIdOk() (*int32, bool)`

GetReviewerUserIdOk returns a tuple with the ReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerUserId

`func (o *EsgTopicsPutEsgTopic) SetReviewerUserId(v int32)`

SetReviewerUserId sets ReviewerUserId field to given value.

### HasReviewerUserId

`func (o *EsgTopicsPutEsgTopic) HasReviewerUserId() bool`

HasReviewerUserId returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *EsgTopicsPutEsgTopic) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *EsgTopicsPutEsgTopic) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *EsgTopicsPutEsgTopic) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *EsgTopicsPutEsgTopic) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetOrganizationImpact

`func (o *EsgTopicsPutEsgTopic) GetOrganizationImpact() interface{}`

GetOrganizationImpact returns the OrganizationImpact field if non-nil, zero value otherwise.

### GetOrganizationImpactOk

`func (o *EsgTopicsPutEsgTopic) GetOrganizationImpactOk() (*interface{}, bool)`

GetOrganizationImpactOk returns a tuple with the OrganizationImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationImpact

`func (o *EsgTopicsPutEsgTopic) SetOrganizationImpact(v interface{})`

SetOrganizationImpact sets OrganizationImpact field to given value.

### HasOrganizationImpact

`func (o *EsgTopicsPutEsgTopic) HasOrganizationImpact() bool`

HasOrganizationImpact returns a boolean if a field has been set.

### SetOrganizationImpactNil

`func (o *EsgTopicsPutEsgTopic) SetOrganizationImpactNil(b bool)`

 SetOrganizationImpactNil sets the value for OrganizationImpact to be an explicit nil

### UnsetOrganizationImpact
`func (o *EsgTopicsPutEsgTopic) UnsetOrganizationImpact()`

UnsetOrganizationImpact ensures that no value is present for OrganizationImpact, not even an explicit nil
### GetStakeholderImpact

`func (o *EsgTopicsPutEsgTopic) GetStakeholderImpact() interface{}`

GetStakeholderImpact returns the StakeholderImpact field if non-nil, zero value otherwise.

### GetStakeholderImpactOk

`func (o *EsgTopicsPutEsgTopic) GetStakeholderImpactOk() (*interface{}, bool)`

GetStakeholderImpactOk returns a tuple with the StakeholderImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderImpact

`func (o *EsgTopicsPutEsgTopic) SetStakeholderImpact(v interface{})`

SetStakeholderImpact sets StakeholderImpact field to given value.

### HasStakeholderImpact

`func (o *EsgTopicsPutEsgTopic) HasStakeholderImpact() bool`

HasStakeholderImpact returns a boolean if a field has been set.

### SetStakeholderImpactNil

`func (o *EsgTopicsPutEsgTopic) SetStakeholderImpactNil(b bool)`

 SetStakeholderImpactNil sets the value for StakeholderImpact to be an explicit nil

### UnsetStakeholderImpact
`func (o *EsgTopicsPutEsgTopic) UnsetStakeholderImpact()`

UnsetStakeholderImpact ensures that no value is present for StakeholderImpact, not even an explicit nil
### GetMaterialityRating

`func (o *EsgTopicsPutEsgTopic) GetMaterialityRating() interface{}`

GetMaterialityRating returns the MaterialityRating field if non-nil, zero value otherwise.

### GetMaterialityRatingOk

`func (o *EsgTopicsPutEsgTopic) GetMaterialityRatingOk() (*interface{}, bool)`

GetMaterialityRatingOk returns a tuple with the MaterialityRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterialityRating

`func (o *EsgTopicsPutEsgTopic) SetMaterialityRating(v interface{})`

SetMaterialityRating sets MaterialityRating field to given value.

### HasMaterialityRating

`func (o *EsgTopicsPutEsgTopic) HasMaterialityRating() bool`

HasMaterialityRating returns a boolean if a field has been set.

### SetMaterialityRatingNil

`func (o *EsgTopicsPutEsgTopic) SetMaterialityRatingNil(b bool)`

 SetMaterialityRatingNil sets the value for MaterialityRating to be an explicit nil

### UnsetMaterialityRating
`func (o *EsgTopicsPutEsgTopic) UnsetMaterialityRating()`

UnsetMaterialityRating ensures that no value is present for MaterialityRating, not even an explicit nil
### GetCategory

`func (o *EsgTopicsPutEsgTopic) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *EsgTopicsPutEsgTopic) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *EsgTopicsPutEsgTopic) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *EsgTopicsPutEsgTopic) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetScope

`func (o *EsgTopicsPutEsgTopic) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EsgTopicsPutEsgTopic) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EsgTopicsPutEsgTopic) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *EsgTopicsPutEsgTopic) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetEsgCategoryId

`func (o *EsgTopicsPutEsgTopic) GetEsgCategoryId() int32`

GetEsgCategoryId returns the EsgCategoryId field if non-nil, zero value otherwise.

### GetEsgCategoryIdOk

`func (o *EsgTopicsPutEsgTopic) GetEsgCategoryIdOk() (*int32, bool)`

GetEsgCategoryIdOk returns a tuple with the EsgCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgCategoryId

`func (o *EsgTopicsPutEsgTopic) SetEsgCategoryId(v int32)`

SetEsgCategoryId sets EsgCategoryId field to given value.

### HasEsgCategoryId

`func (o *EsgTopicsPutEsgTopic) HasEsgCategoryId() bool`

HasEsgCategoryId returns a boolean if a field has been set.

### GetEsgTopicCustomSelect1OptionId

`func (o *EsgTopicsPutEsgTopic) GetEsgTopicCustomSelect1OptionId() int32`

GetEsgTopicCustomSelect1OptionId returns the EsgTopicCustomSelect1OptionId field if non-nil, zero value otherwise.

### GetEsgTopicCustomSelect1OptionIdOk

`func (o *EsgTopicsPutEsgTopic) GetEsgTopicCustomSelect1OptionIdOk() (*int32, bool)`

GetEsgTopicCustomSelect1OptionIdOk returns a tuple with the EsgTopicCustomSelect1OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgTopicCustomSelect1OptionId

`func (o *EsgTopicsPutEsgTopic) SetEsgTopicCustomSelect1OptionId(v int32)`

SetEsgTopicCustomSelect1OptionId sets EsgTopicCustomSelect1OptionId field to given value.

### HasEsgTopicCustomSelect1OptionId

`func (o *EsgTopicsPutEsgTopic) HasEsgTopicCustomSelect1OptionId() bool`

HasEsgTopicCustomSelect1OptionId returns a boolean if a field has been set.

### GetEsgTopicCustomSelect2OptionId

`func (o *EsgTopicsPutEsgTopic) GetEsgTopicCustomSelect2OptionId() int32`

GetEsgTopicCustomSelect2OptionId returns the EsgTopicCustomSelect2OptionId field if non-nil, zero value otherwise.

### GetEsgTopicCustomSelect2OptionIdOk

`func (o *EsgTopicsPutEsgTopic) GetEsgTopicCustomSelect2OptionIdOk() (*int32, bool)`

GetEsgTopicCustomSelect2OptionIdOk returns a tuple with the EsgTopicCustomSelect2OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgTopicCustomSelect2OptionId

`func (o *EsgTopicsPutEsgTopic) SetEsgTopicCustomSelect2OptionId(v int32)`

SetEsgTopicCustomSelect2OptionId sets EsgTopicCustomSelect2OptionId field to given value.

### HasEsgTopicCustomSelect2OptionId

`func (o *EsgTopicsPutEsgTopic) HasEsgTopicCustomSelect2OptionId() bool`

HasEsgTopicCustomSelect2OptionId returns a boolean if a field has been set.

### GetEsgTopicCustomSelect3OptionId

`func (o *EsgTopicsPutEsgTopic) GetEsgTopicCustomSelect3OptionId() int32`

GetEsgTopicCustomSelect3OptionId returns the EsgTopicCustomSelect3OptionId field if non-nil, zero value otherwise.

### GetEsgTopicCustomSelect3OptionIdOk

`func (o *EsgTopicsPutEsgTopic) GetEsgTopicCustomSelect3OptionIdOk() (*int32, bool)`

GetEsgTopicCustomSelect3OptionIdOk returns a tuple with the EsgTopicCustomSelect3OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgTopicCustomSelect3OptionId

`func (o *EsgTopicsPutEsgTopic) SetEsgTopicCustomSelect3OptionId(v int32)`

SetEsgTopicCustomSelect3OptionId sets EsgTopicCustomSelect3OptionId field to given value.

### HasEsgTopicCustomSelect3OptionId

`func (o *EsgTopicsPutEsgTopic) HasEsgTopicCustomSelect3OptionId() bool`

HasEsgTopicCustomSelect3OptionId returns a boolean if a field has been set.

### GetEsgTopicCustomSelect4OptionId

`func (o *EsgTopicsPutEsgTopic) GetEsgTopicCustomSelect4OptionId() int32`

GetEsgTopicCustomSelect4OptionId returns the EsgTopicCustomSelect4OptionId field if non-nil, zero value otherwise.

### GetEsgTopicCustomSelect4OptionIdOk

`func (o *EsgTopicsPutEsgTopic) GetEsgTopicCustomSelect4OptionIdOk() (*int32, bool)`

GetEsgTopicCustomSelect4OptionIdOk returns a tuple with the EsgTopicCustomSelect4OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgTopicCustomSelect4OptionId

`func (o *EsgTopicsPutEsgTopic) SetEsgTopicCustomSelect4OptionId(v int32)`

SetEsgTopicCustomSelect4OptionId sets EsgTopicCustomSelect4OptionId field to given value.

### HasEsgTopicCustomSelect4OptionId

`func (o *EsgTopicsPutEsgTopic) HasEsgTopicCustomSelect4OptionId() bool`

HasEsgTopicCustomSelect4OptionId returns a boolean if a field has been set.

### GetEsgTopicCustomSelect5OptionId

`func (o *EsgTopicsPutEsgTopic) GetEsgTopicCustomSelect5OptionId() int32`

GetEsgTopicCustomSelect5OptionId returns the EsgTopicCustomSelect5OptionId field if non-nil, zero value otherwise.

### GetEsgTopicCustomSelect5OptionIdOk

`func (o *EsgTopicsPutEsgTopic) GetEsgTopicCustomSelect5OptionIdOk() (*int32, bool)`

GetEsgTopicCustomSelect5OptionIdOk returns a tuple with the EsgTopicCustomSelect5OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgTopicCustomSelect5OptionId

`func (o *EsgTopicsPutEsgTopic) SetEsgTopicCustomSelect5OptionId(v int32)`

SetEsgTopicCustomSelect5OptionId sets EsgTopicCustomSelect5OptionId field to given value.

### HasEsgTopicCustomSelect5OptionId

`func (o *EsgTopicsPutEsgTopic) HasEsgTopicCustomSelect5OptionId() bool`

HasEsgTopicCustomSelect5OptionId returns a boolean if a field has been set.

### GetCustomDate1

`func (o *EsgTopicsPutEsgTopic) GetCustomDate1() string`

GetCustomDate1 returns the CustomDate1 field if non-nil, zero value otherwise.

### GetCustomDate1Ok

`func (o *EsgTopicsPutEsgTopic) GetCustomDate1Ok() (*string, bool)`

GetCustomDate1Ok returns a tuple with the CustomDate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate1

`func (o *EsgTopicsPutEsgTopic) SetCustomDate1(v string)`

SetCustomDate1 sets CustomDate1 field to given value.

### HasCustomDate1

`func (o *EsgTopicsPutEsgTopic) HasCustomDate1() bool`

HasCustomDate1 returns a boolean if a field has been set.

### GetCustomDate2

`func (o *EsgTopicsPutEsgTopic) GetCustomDate2() string`

GetCustomDate2 returns the CustomDate2 field if non-nil, zero value otherwise.

### GetCustomDate2Ok

`func (o *EsgTopicsPutEsgTopic) GetCustomDate2Ok() (*string, bool)`

GetCustomDate2Ok returns a tuple with the CustomDate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate2

`func (o *EsgTopicsPutEsgTopic) SetCustomDate2(v string)`

SetCustomDate2 sets CustomDate2 field to given value.

### HasCustomDate2

`func (o *EsgTopicsPutEsgTopic) HasCustomDate2() bool`

HasCustomDate2 returns a boolean if a field has been set.

### GetCustomDate3

`func (o *EsgTopicsPutEsgTopic) GetCustomDate3() string`

GetCustomDate3 returns the CustomDate3 field if non-nil, zero value otherwise.

### GetCustomDate3Ok

`func (o *EsgTopicsPutEsgTopic) GetCustomDate3Ok() (*string, bool)`

GetCustomDate3Ok returns a tuple with the CustomDate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate3

`func (o *EsgTopicsPutEsgTopic) SetCustomDate3(v string)`

SetCustomDate3 sets CustomDate3 field to given value.

### HasCustomDate3

`func (o *EsgTopicsPutEsgTopic) HasCustomDate3() bool`

HasCustomDate3 returns a boolean if a field has been set.

### GetCustomText1

`func (o *EsgTopicsPutEsgTopic) GetCustomText1() string`

GetCustomText1 returns the CustomText1 field if non-nil, zero value otherwise.

### GetCustomText1Ok

`func (o *EsgTopicsPutEsgTopic) GetCustomText1Ok() (*string, bool)`

GetCustomText1Ok returns a tuple with the CustomText1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText1

`func (o *EsgTopicsPutEsgTopic) SetCustomText1(v string)`

SetCustomText1 sets CustomText1 field to given value.

### HasCustomText1

`func (o *EsgTopicsPutEsgTopic) HasCustomText1() bool`

HasCustomText1 returns a boolean if a field has been set.

### GetCustomText2

`func (o *EsgTopicsPutEsgTopic) GetCustomText2() string`

GetCustomText2 returns the CustomText2 field if non-nil, zero value otherwise.

### GetCustomText2Ok

`func (o *EsgTopicsPutEsgTopic) GetCustomText2Ok() (*string, bool)`

GetCustomText2Ok returns a tuple with the CustomText2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText2

`func (o *EsgTopicsPutEsgTopic) SetCustomText2(v string)`

SetCustomText2 sets CustomText2 field to given value.

### HasCustomText2

`func (o *EsgTopicsPutEsgTopic) HasCustomText2() bool`

HasCustomText2 returns a boolean if a field has been set.

### GetCustomText3

`func (o *EsgTopicsPutEsgTopic) GetCustomText3() string`

GetCustomText3 returns the CustomText3 field if non-nil, zero value otherwise.

### GetCustomText3Ok

`func (o *EsgTopicsPutEsgTopic) GetCustomText3Ok() (*string, bool)`

GetCustomText3Ok returns a tuple with the CustomText3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText3

`func (o *EsgTopicsPutEsgTopic) SetCustomText3(v string)`

SetCustomText3 sets CustomText3 field to given value.

### HasCustomText3

`func (o *EsgTopicsPutEsgTopic) HasCustomText3() bool`

HasCustomText3 returns a boolean if a field has been set.

### GetCustomText4

`func (o *EsgTopicsPutEsgTopic) GetCustomText4() string`

GetCustomText4 returns the CustomText4 field if non-nil, zero value otherwise.

### GetCustomText4Ok

`func (o *EsgTopicsPutEsgTopic) GetCustomText4Ok() (*string, bool)`

GetCustomText4Ok returns a tuple with the CustomText4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText4

`func (o *EsgTopicsPutEsgTopic) SetCustomText4(v string)`

SetCustomText4 sets CustomText4 field to given value.

### HasCustomText4

`func (o *EsgTopicsPutEsgTopic) HasCustomText4() bool`

HasCustomText4 returns a boolean if a field has been set.

### GetCustomText5

`func (o *EsgTopicsPutEsgTopic) GetCustomText5() string`

GetCustomText5 returns the CustomText5 field if non-nil, zero value otherwise.

### GetCustomText5Ok

`func (o *EsgTopicsPutEsgTopic) GetCustomText5Ok() (*string, bool)`

GetCustomText5Ok returns a tuple with the CustomText5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText5

`func (o *EsgTopicsPutEsgTopic) SetCustomText5(v string)`

SetCustomText5 sets CustomText5 field to given value.

### HasCustomText5

`func (o *EsgTopicsPutEsgTopic) HasCustomText5() bool`

HasCustomText5 returns a boolean if a field has been set.

### GetCustomText6

`func (o *EsgTopicsPutEsgTopic) GetCustomText6() string`

GetCustomText6 returns the CustomText6 field if non-nil, zero value otherwise.

### GetCustomText6Ok

`func (o *EsgTopicsPutEsgTopic) GetCustomText6Ok() (*string, bool)`

GetCustomText6Ok returns a tuple with the CustomText6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText6

`func (o *EsgTopicsPutEsgTopic) SetCustomText6(v string)`

SetCustomText6 sets CustomText6 field to given value.

### HasCustomText6

`func (o *EsgTopicsPutEsgTopic) HasCustomText6() bool`

HasCustomText6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


