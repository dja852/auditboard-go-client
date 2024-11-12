# FrameworkItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**FrameworkId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;frameworks.id&#x60;.&lt;fk table&#x3D;&#39;frameworks&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Section** | Pointer to **string** |  | [optional] 
**Subsection** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | [default to "Active"]
**SortOrder** | Pointer to **int32** |  | [optional] [default to 0]
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Scope** | Pointer to **string** |  | [optional] [default to "In Scope"]
**CustomDate1** | Pointer to **string** |  | [optional] 
**CustomDate2** | Pointer to **string** |  | [optional] 
**CustomDate3** | Pointer to **string** |  | [optional] 
**CustomText1** | Pointer to **string** |  | [optional] 
**CustomText2** | Pointer to **string** |  | [optional] 
**CustomText3** | Pointer to **string** |  | [optional] 
**FrameworkItemCustomSelect1OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;framework_item_custom_select1_options.id&#x60;.&lt;fk table&#x3D;&#39;framework_item_custom_select1_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FrameworkItemCustomSelect2OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;framework_item_custom_select2_options.id&#x60;.&lt;fk table&#x3D;&#39;framework_item_custom_select2_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FrameworkItemCustomSelect3OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;framework_item_custom_select3_options.id&#x60;.&lt;fk table&#x3D;&#39;framework_item_custom_select3_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RequirementOwnerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RequirementReviewerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Request** | Pointer to **string** |  | [optional] 
**AdditionalRequest** | Pointer to **string** |  | [optional] 
**FrameworkItemCustomSelect4OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;framework_item_custom_select4_options.id&#x60;.&lt;fk table&#x3D;&#39;framework_item_custom_select4_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FrameworkItemCustomSelect5OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;framework_item_custom_select5_options.id&#x60;.&lt;fk table&#x3D;&#39;framework_item_custom_select5_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ImplementationGuidance** | Pointer to **string** |  | [optional] 
**FrameworkItemImplementationStatusId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;framework_item_implementation_statuses.id&#x60;.&lt;fk table&#x3D;&#39;framework_item_implementation_statuses&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FrameworkItemRatingId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;framework_item_ratings.id&#x60;.&lt;fk table&#x3D;&#39;framework_item_ratings&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomText4** | Pointer to **string** |  | [optional] 
**CustomText5** | Pointer to **string** |  | [optional] 
**CustomText6** | Pointer to **string** |  | [optional] 
**ReferenceMeta** | **interface{}** |  | 
**CoordinatorUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FrameworkItemActualMaturityId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;framework_item_actual_maturities.id&#x60;.&lt;fk table&#x3D;&#39;framework_item_actual_maturities&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FrameworkItemTargetMaturityId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;framework_item_target_maturities.id&#x60;.&lt;fk table&#x3D;&#39;framework_item_target_maturities&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ContentServiceId** | Pointer to **string** |  | [optional] 
**DescriptionWithTags** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **int32** |  | [optional] 
**IsCcf** | **bool** |  | [default to false]
**Type** | Pointer to **string** |  | [optional] [default to ""]
**Classification** | Pointer to **string** |  | [optional] [default to ""]
**MinimumRequirementLevel** | **string** |  | [default to "0"]
**FieldData** | Pointer to **interface{}** |  | [optional] 
**Response** | Pointer to **string** |  | [optional] 
**IdCode** | Pointer to **string** |  | [optional] 
**ShortIdCode** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **interface{}** |  | [optional] 
**TargetMaturityScore** | Pointer to **int32** |  | [optional] 
**XcTotals** | Pointer to **interface{}** |  | [optional] 
**AtRiskStatus** | **string** |  | [default to "Not At Risk"]
**FrameworkImplementationGuidance** | Pointer to **string** |  | [optional] 
**FrameworkAssessmentProcedures** | Pointer to **string** |  | [optional] 
**FrameworkRequirementPurpose** | Pointer to **string** |  | [optional] 
**FrameworkItemFunctionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;framework_item_functions.id&#x60;.&lt;fk table&#x3D;&#39;framework_item_functions&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RelativeControlWeighting** | Pointer to **string** |  | [optional] 
**CmmLevel0** | Pointer to **string** |  | [optional] 
**CmmLevel1** | Pointer to **string** |  | [optional] 
**CmmLevel2** | Pointer to **string** |  | [optional] 
**CmmLevel3** | Pointer to **string** |  | [optional] 
**CmmLevel4** | Pointer to **string** |  | [optional] 
**CmmLevel5** | Pointer to **string** |  | [optional] 

## Methods

### NewFrameworkItems

`func NewFrameworkItems(status string, referenceMeta interface{}, isCcf bool, minimumRequirementLevel string, atRiskStatus string, ) *FrameworkItems`

NewFrameworkItems instantiates a new FrameworkItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameworkItemsWithDefaults

`func NewFrameworkItemsWithDefaults() *FrameworkItems`

NewFrameworkItemsWithDefaults instantiates a new FrameworkItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FrameworkItems) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FrameworkItems) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FrameworkItems) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FrameworkItems) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFrameworkId

`func (o *FrameworkItems) GetFrameworkId() int32`

GetFrameworkId returns the FrameworkId field if non-nil, zero value otherwise.

### GetFrameworkIdOk

`func (o *FrameworkItems) GetFrameworkIdOk() (*int32, bool)`

GetFrameworkIdOk returns a tuple with the FrameworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkId

`func (o *FrameworkItems) SetFrameworkId(v int32)`

SetFrameworkId sets FrameworkId field to given value.

### HasFrameworkId

`func (o *FrameworkItems) HasFrameworkId() bool`

HasFrameworkId returns a boolean if a field has been set.

### GetSection

`func (o *FrameworkItems) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *FrameworkItems) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *FrameworkItems) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *FrameworkItems) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetSubsection

`func (o *FrameworkItems) GetSubsection() string`

GetSubsection returns the Subsection field if non-nil, zero value otherwise.

### GetSubsectionOk

`func (o *FrameworkItems) GetSubsectionOk() (*string, bool)`

GetSubsectionOk returns a tuple with the Subsection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsection

`func (o *FrameworkItems) SetSubsection(v string)`

SetSubsection sets Subsection field to given value.

### HasSubsection

`func (o *FrameworkItems) HasSubsection() bool`

HasSubsection returns a boolean if a field has been set.

### GetUid

`func (o *FrameworkItems) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *FrameworkItems) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *FrameworkItems) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *FrameworkItems) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetName

`func (o *FrameworkItems) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FrameworkItems) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FrameworkItems) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FrameworkItems) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FrameworkItems) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FrameworkItems) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FrameworkItems) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FrameworkItems) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *FrameworkItems) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FrameworkItems) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FrameworkItems) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSortOrder

`func (o *FrameworkItems) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *FrameworkItems) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *FrameworkItems) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *FrameworkItems) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FrameworkItems) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FrameworkItems) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FrameworkItems) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FrameworkItems) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FrameworkItems) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FrameworkItems) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FrameworkItems) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FrameworkItems) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *FrameworkItems) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *FrameworkItems) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *FrameworkItems) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *FrameworkItems) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetScope

`func (o *FrameworkItems) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *FrameworkItems) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *FrameworkItems) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *FrameworkItems) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetCustomDate1

`func (o *FrameworkItems) GetCustomDate1() string`

GetCustomDate1 returns the CustomDate1 field if non-nil, zero value otherwise.

### GetCustomDate1Ok

`func (o *FrameworkItems) GetCustomDate1Ok() (*string, bool)`

GetCustomDate1Ok returns a tuple with the CustomDate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate1

`func (o *FrameworkItems) SetCustomDate1(v string)`

SetCustomDate1 sets CustomDate1 field to given value.

### HasCustomDate1

`func (o *FrameworkItems) HasCustomDate1() bool`

HasCustomDate1 returns a boolean if a field has been set.

### GetCustomDate2

`func (o *FrameworkItems) GetCustomDate2() string`

GetCustomDate2 returns the CustomDate2 field if non-nil, zero value otherwise.

### GetCustomDate2Ok

`func (o *FrameworkItems) GetCustomDate2Ok() (*string, bool)`

GetCustomDate2Ok returns a tuple with the CustomDate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate2

`func (o *FrameworkItems) SetCustomDate2(v string)`

SetCustomDate2 sets CustomDate2 field to given value.

### HasCustomDate2

`func (o *FrameworkItems) HasCustomDate2() bool`

HasCustomDate2 returns a boolean if a field has been set.

### GetCustomDate3

`func (o *FrameworkItems) GetCustomDate3() string`

GetCustomDate3 returns the CustomDate3 field if non-nil, zero value otherwise.

### GetCustomDate3Ok

`func (o *FrameworkItems) GetCustomDate3Ok() (*string, bool)`

GetCustomDate3Ok returns a tuple with the CustomDate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate3

`func (o *FrameworkItems) SetCustomDate3(v string)`

SetCustomDate3 sets CustomDate3 field to given value.

### HasCustomDate3

`func (o *FrameworkItems) HasCustomDate3() bool`

HasCustomDate3 returns a boolean if a field has been set.

### GetCustomText1

`func (o *FrameworkItems) GetCustomText1() string`

GetCustomText1 returns the CustomText1 field if non-nil, zero value otherwise.

### GetCustomText1Ok

`func (o *FrameworkItems) GetCustomText1Ok() (*string, bool)`

GetCustomText1Ok returns a tuple with the CustomText1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText1

`func (o *FrameworkItems) SetCustomText1(v string)`

SetCustomText1 sets CustomText1 field to given value.

### HasCustomText1

`func (o *FrameworkItems) HasCustomText1() bool`

HasCustomText1 returns a boolean if a field has been set.

### GetCustomText2

`func (o *FrameworkItems) GetCustomText2() string`

GetCustomText2 returns the CustomText2 field if non-nil, zero value otherwise.

### GetCustomText2Ok

`func (o *FrameworkItems) GetCustomText2Ok() (*string, bool)`

GetCustomText2Ok returns a tuple with the CustomText2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText2

`func (o *FrameworkItems) SetCustomText2(v string)`

SetCustomText2 sets CustomText2 field to given value.

### HasCustomText2

`func (o *FrameworkItems) HasCustomText2() bool`

HasCustomText2 returns a boolean if a field has been set.

### GetCustomText3

`func (o *FrameworkItems) GetCustomText3() string`

GetCustomText3 returns the CustomText3 field if non-nil, zero value otherwise.

### GetCustomText3Ok

`func (o *FrameworkItems) GetCustomText3Ok() (*string, bool)`

GetCustomText3Ok returns a tuple with the CustomText3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText3

`func (o *FrameworkItems) SetCustomText3(v string)`

SetCustomText3 sets CustomText3 field to given value.

### HasCustomText3

`func (o *FrameworkItems) HasCustomText3() bool`

HasCustomText3 returns a boolean if a field has been set.

### GetFrameworkItemCustomSelect1OptionId

`func (o *FrameworkItems) GetFrameworkItemCustomSelect1OptionId() int32`

GetFrameworkItemCustomSelect1OptionId returns the FrameworkItemCustomSelect1OptionId field if non-nil, zero value otherwise.

### GetFrameworkItemCustomSelect1OptionIdOk

`func (o *FrameworkItems) GetFrameworkItemCustomSelect1OptionIdOk() (*int32, bool)`

GetFrameworkItemCustomSelect1OptionIdOk returns a tuple with the FrameworkItemCustomSelect1OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemCustomSelect1OptionId

`func (o *FrameworkItems) SetFrameworkItemCustomSelect1OptionId(v int32)`

SetFrameworkItemCustomSelect1OptionId sets FrameworkItemCustomSelect1OptionId field to given value.

### HasFrameworkItemCustomSelect1OptionId

`func (o *FrameworkItems) HasFrameworkItemCustomSelect1OptionId() bool`

HasFrameworkItemCustomSelect1OptionId returns a boolean if a field has been set.

### GetFrameworkItemCustomSelect2OptionId

`func (o *FrameworkItems) GetFrameworkItemCustomSelect2OptionId() int32`

GetFrameworkItemCustomSelect2OptionId returns the FrameworkItemCustomSelect2OptionId field if non-nil, zero value otherwise.

### GetFrameworkItemCustomSelect2OptionIdOk

`func (o *FrameworkItems) GetFrameworkItemCustomSelect2OptionIdOk() (*int32, bool)`

GetFrameworkItemCustomSelect2OptionIdOk returns a tuple with the FrameworkItemCustomSelect2OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemCustomSelect2OptionId

`func (o *FrameworkItems) SetFrameworkItemCustomSelect2OptionId(v int32)`

SetFrameworkItemCustomSelect2OptionId sets FrameworkItemCustomSelect2OptionId field to given value.

### HasFrameworkItemCustomSelect2OptionId

`func (o *FrameworkItems) HasFrameworkItemCustomSelect2OptionId() bool`

HasFrameworkItemCustomSelect2OptionId returns a boolean if a field has been set.

### GetFrameworkItemCustomSelect3OptionId

`func (o *FrameworkItems) GetFrameworkItemCustomSelect3OptionId() int32`

GetFrameworkItemCustomSelect3OptionId returns the FrameworkItemCustomSelect3OptionId field if non-nil, zero value otherwise.

### GetFrameworkItemCustomSelect3OptionIdOk

`func (o *FrameworkItems) GetFrameworkItemCustomSelect3OptionIdOk() (*int32, bool)`

GetFrameworkItemCustomSelect3OptionIdOk returns a tuple with the FrameworkItemCustomSelect3OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemCustomSelect3OptionId

`func (o *FrameworkItems) SetFrameworkItemCustomSelect3OptionId(v int32)`

SetFrameworkItemCustomSelect3OptionId sets FrameworkItemCustomSelect3OptionId field to given value.

### HasFrameworkItemCustomSelect3OptionId

`func (o *FrameworkItems) HasFrameworkItemCustomSelect3OptionId() bool`

HasFrameworkItemCustomSelect3OptionId returns a boolean if a field has been set.

### GetRequirementOwnerUserId

`func (o *FrameworkItems) GetRequirementOwnerUserId() int32`

GetRequirementOwnerUserId returns the RequirementOwnerUserId field if non-nil, zero value otherwise.

### GetRequirementOwnerUserIdOk

`func (o *FrameworkItems) GetRequirementOwnerUserIdOk() (*int32, bool)`

GetRequirementOwnerUserIdOk returns a tuple with the RequirementOwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementOwnerUserId

`func (o *FrameworkItems) SetRequirementOwnerUserId(v int32)`

SetRequirementOwnerUserId sets RequirementOwnerUserId field to given value.

### HasRequirementOwnerUserId

`func (o *FrameworkItems) HasRequirementOwnerUserId() bool`

HasRequirementOwnerUserId returns a boolean if a field has been set.

### GetRequirementReviewerUserId

`func (o *FrameworkItems) GetRequirementReviewerUserId() int32`

GetRequirementReviewerUserId returns the RequirementReviewerUserId field if non-nil, zero value otherwise.

### GetRequirementReviewerUserIdOk

`func (o *FrameworkItems) GetRequirementReviewerUserIdOk() (*int32, bool)`

GetRequirementReviewerUserIdOk returns a tuple with the RequirementReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementReviewerUserId

`func (o *FrameworkItems) SetRequirementReviewerUserId(v int32)`

SetRequirementReviewerUserId sets RequirementReviewerUserId field to given value.

### HasRequirementReviewerUserId

`func (o *FrameworkItems) HasRequirementReviewerUserId() bool`

HasRequirementReviewerUserId returns a boolean if a field has been set.

### GetRequest

`func (o *FrameworkItems) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *FrameworkItems) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *FrameworkItems) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *FrameworkItems) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetAdditionalRequest

`func (o *FrameworkItems) GetAdditionalRequest() string`

GetAdditionalRequest returns the AdditionalRequest field if non-nil, zero value otherwise.

### GetAdditionalRequestOk

`func (o *FrameworkItems) GetAdditionalRequestOk() (*string, bool)`

GetAdditionalRequestOk returns a tuple with the AdditionalRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalRequest

`func (o *FrameworkItems) SetAdditionalRequest(v string)`

SetAdditionalRequest sets AdditionalRequest field to given value.

### HasAdditionalRequest

`func (o *FrameworkItems) HasAdditionalRequest() bool`

HasAdditionalRequest returns a boolean if a field has been set.

### GetFrameworkItemCustomSelect4OptionId

`func (o *FrameworkItems) GetFrameworkItemCustomSelect4OptionId() int32`

GetFrameworkItemCustomSelect4OptionId returns the FrameworkItemCustomSelect4OptionId field if non-nil, zero value otherwise.

### GetFrameworkItemCustomSelect4OptionIdOk

`func (o *FrameworkItems) GetFrameworkItemCustomSelect4OptionIdOk() (*int32, bool)`

GetFrameworkItemCustomSelect4OptionIdOk returns a tuple with the FrameworkItemCustomSelect4OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemCustomSelect4OptionId

`func (o *FrameworkItems) SetFrameworkItemCustomSelect4OptionId(v int32)`

SetFrameworkItemCustomSelect4OptionId sets FrameworkItemCustomSelect4OptionId field to given value.

### HasFrameworkItemCustomSelect4OptionId

`func (o *FrameworkItems) HasFrameworkItemCustomSelect4OptionId() bool`

HasFrameworkItemCustomSelect4OptionId returns a boolean if a field has been set.

### GetFrameworkItemCustomSelect5OptionId

`func (o *FrameworkItems) GetFrameworkItemCustomSelect5OptionId() int32`

GetFrameworkItemCustomSelect5OptionId returns the FrameworkItemCustomSelect5OptionId field if non-nil, zero value otherwise.

### GetFrameworkItemCustomSelect5OptionIdOk

`func (o *FrameworkItems) GetFrameworkItemCustomSelect5OptionIdOk() (*int32, bool)`

GetFrameworkItemCustomSelect5OptionIdOk returns a tuple with the FrameworkItemCustomSelect5OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemCustomSelect5OptionId

`func (o *FrameworkItems) SetFrameworkItemCustomSelect5OptionId(v int32)`

SetFrameworkItemCustomSelect5OptionId sets FrameworkItemCustomSelect5OptionId field to given value.

### HasFrameworkItemCustomSelect5OptionId

`func (o *FrameworkItems) HasFrameworkItemCustomSelect5OptionId() bool`

HasFrameworkItemCustomSelect5OptionId returns a boolean if a field has been set.

### GetImplementationGuidance

`func (o *FrameworkItems) GetImplementationGuidance() string`

GetImplementationGuidance returns the ImplementationGuidance field if non-nil, zero value otherwise.

### GetImplementationGuidanceOk

`func (o *FrameworkItems) GetImplementationGuidanceOk() (*string, bool)`

GetImplementationGuidanceOk returns a tuple with the ImplementationGuidance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationGuidance

`func (o *FrameworkItems) SetImplementationGuidance(v string)`

SetImplementationGuidance sets ImplementationGuidance field to given value.

### HasImplementationGuidance

`func (o *FrameworkItems) HasImplementationGuidance() bool`

HasImplementationGuidance returns a boolean if a field has been set.

### GetFrameworkItemImplementationStatusId

`func (o *FrameworkItems) GetFrameworkItemImplementationStatusId() int32`

GetFrameworkItemImplementationStatusId returns the FrameworkItemImplementationStatusId field if non-nil, zero value otherwise.

### GetFrameworkItemImplementationStatusIdOk

`func (o *FrameworkItems) GetFrameworkItemImplementationStatusIdOk() (*int32, bool)`

GetFrameworkItemImplementationStatusIdOk returns a tuple with the FrameworkItemImplementationStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemImplementationStatusId

`func (o *FrameworkItems) SetFrameworkItemImplementationStatusId(v int32)`

SetFrameworkItemImplementationStatusId sets FrameworkItemImplementationStatusId field to given value.

### HasFrameworkItemImplementationStatusId

`func (o *FrameworkItems) HasFrameworkItemImplementationStatusId() bool`

HasFrameworkItemImplementationStatusId returns a boolean if a field has been set.

### GetFrameworkItemRatingId

`func (o *FrameworkItems) GetFrameworkItemRatingId() int32`

GetFrameworkItemRatingId returns the FrameworkItemRatingId field if non-nil, zero value otherwise.

### GetFrameworkItemRatingIdOk

`func (o *FrameworkItems) GetFrameworkItemRatingIdOk() (*int32, bool)`

GetFrameworkItemRatingIdOk returns a tuple with the FrameworkItemRatingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemRatingId

`func (o *FrameworkItems) SetFrameworkItemRatingId(v int32)`

SetFrameworkItemRatingId sets FrameworkItemRatingId field to given value.

### HasFrameworkItemRatingId

`func (o *FrameworkItems) HasFrameworkItemRatingId() bool`

HasFrameworkItemRatingId returns a boolean if a field has been set.

### GetCustomText4

`func (o *FrameworkItems) GetCustomText4() string`

GetCustomText4 returns the CustomText4 field if non-nil, zero value otherwise.

### GetCustomText4Ok

`func (o *FrameworkItems) GetCustomText4Ok() (*string, bool)`

GetCustomText4Ok returns a tuple with the CustomText4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText4

`func (o *FrameworkItems) SetCustomText4(v string)`

SetCustomText4 sets CustomText4 field to given value.

### HasCustomText4

`func (o *FrameworkItems) HasCustomText4() bool`

HasCustomText4 returns a boolean if a field has been set.

### GetCustomText5

`func (o *FrameworkItems) GetCustomText5() string`

GetCustomText5 returns the CustomText5 field if non-nil, zero value otherwise.

### GetCustomText5Ok

`func (o *FrameworkItems) GetCustomText5Ok() (*string, bool)`

GetCustomText5Ok returns a tuple with the CustomText5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText5

`func (o *FrameworkItems) SetCustomText5(v string)`

SetCustomText5 sets CustomText5 field to given value.

### HasCustomText5

`func (o *FrameworkItems) HasCustomText5() bool`

HasCustomText5 returns a boolean if a field has been set.

### GetCustomText6

`func (o *FrameworkItems) GetCustomText6() string`

GetCustomText6 returns the CustomText6 field if non-nil, zero value otherwise.

### GetCustomText6Ok

`func (o *FrameworkItems) GetCustomText6Ok() (*string, bool)`

GetCustomText6Ok returns a tuple with the CustomText6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText6

`func (o *FrameworkItems) SetCustomText6(v string)`

SetCustomText6 sets CustomText6 field to given value.

### HasCustomText6

`func (o *FrameworkItems) HasCustomText6() bool`

HasCustomText6 returns a boolean if a field has been set.

### GetReferenceMeta

`func (o *FrameworkItems) GetReferenceMeta() interface{}`

GetReferenceMeta returns the ReferenceMeta field if non-nil, zero value otherwise.

### GetReferenceMetaOk

`func (o *FrameworkItems) GetReferenceMetaOk() (*interface{}, bool)`

GetReferenceMetaOk returns a tuple with the ReferenceMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMeta

`func (o *FrameworkItems) SetReferenceMeta(v interface{})`

SetReferenceMeta sets ReferenceMeta field to given value.


### SetReferenceMetaNil

`func (o *FrameworkItems) SetReferenceMetaNil(b bool)`

 SetReferenceMetaNil sets the value for ReferenceMeta to be an explicit nil

### UnsetReferenceMeta
`func (o *FrameworkItems) UnsetReferenceMeta()`

UnsetReferenceMeta ensures that no value is present for ReferenceMeta, not even an explicit nil
### GetCoordinatorUserId

`func (o *FrameworkItems) GetCoordinatorUserId() int32`

GetCoordinatorUserId returns the CoordinatorUserId field if non-nil, zero value otherwise.

### GetCoordinatorUserIdOk

`func (o *FrameworkItems) GetCoordinatorUserIdOk() (*int32, bool)`

GetCoordinatorUserIdOk returns a tuple with the CoordinatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinatorUserId

`func (o *FrameworkItems) SetCoordinatorUserId(v int32)`

SetCoordinatorUserId sets CoordinatorUserId field to given value.

### HasCoordinatorUserId

`func (o *FrameworkItems) HasCoordinatorUserId() bool`

HasCoordinatorUserId returns a boolean if a field has been set.

### GetFrameworkItemActualMaturityId

`func (o *FrameworkItems) GetFrameworkItemActualMaturityId() int32`

GetFrameworkItemActualMaturityId returns the FrameworkItemActualMaturityId field if non-nil, zero value otherwise.

### GetFrameworkItemActualMaturityIdOk

`func (o *FrameworkItems) GetFrameworkItemActualMaturityIdOk() (*int32, bool)`

GetFrameworkItemActualMaturityIdOk returns a tuple with the FrameworkItemActualMaturityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemActualMaturityId

`func (o *FrameworkItems) SetFrameworkItemActualMaturityId(v int32)`

SetFrameworkItemActualMaturityId sets FrameworkItemActualMaturityId field to given value.

### HasFrameworkItemActualMaturityId

`func (o *FrameworkItems) HasFrameworkItemActualMaturityId() bool`

HasFrameworkItemActualMaturityId returns a boolean if a field has been set.

### GetFrameworkItemTargetMaturityId

`func (o *FrameworkItems) GetFrameworkItemTargetMaturityId() int32`

GetFrameworkItemTargetMaturityId returns the FrameworkItemTargetMaturityId field if non-nil, zero value otherwise.

### GetFrameworkItemTargetMaturityIdOk

`func (o *FrameworkItems) GetFrameworkItemTargetMaturityIdOk() (*int32, bool)`

GetFrameworkItemTargetMaturityIdOk returns a tuple with the FrameworkItemTargetMaturityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemTargetMaturityId

`func (o *FrameworkItems) SetFrameworkItemTargetMaturityId(v int32)`

SetFrameworkItemTargetMaturityId sets FrameworkItemTargetMaturityId field to given value.

### HasFrameworkItemTargetMaturityId

`func (o *FrameworkItems) HasFrameworkItemTargetMaturityId() bool`

HasFrameworkItemTargetMaturityId returns a boolean if a field has been set.

### GetContentServiceId

`func (o *FrameworkItems) GetContentServiceId() string`

GetContentServiceId returns the ContentServiceId field if non-nil, zero value otherwise.

### GetContentServiceIdOk

`func (o *FrameworkItems) GetContentServiceIdOk() (*string, bool)`

GetContentServiceIdOk returns a tuple with the ContentServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentServiceId

`func (o *FrameworkItems) SetContentServiceId(v string)`

SetContentServiceId sets ContentServiceId field to given value.

### HasContentServiceId

`func (o *FrameworkItems) HasContentServiceId() bool`

HasContentServiceId returns a boolean if a field has been set.

### GetDescriptionWithTags

`func (o *FrameworkItems) GetDescriptionWithTags() string`

GetDescriptionWithTags returns the DescriptionWithTags field if non-nil, zero value otherwise.

### GetDescriptionWithTagsOk

`func (o *FrameworkItems) GetDescriptionWithTagsOk() (*string, bool)`

GetDescriptionWithTagsOk returns a tuple with the DescriptionWithTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionWithTags

`func (o *FrameworkItems) SetDescriptionWithTags(v string)`

SetDescriptionWithTags sets DescriptionWithTags field to given value.

### HasDescriptionWithTags

`func (o *FrameworkItems) HasDescriptionWithTags() bool`

HasDescriptionWithTags returns a boolean if a field has been set.

### GetParentId

`func (o *FrameworkItems) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *FrameworkItems) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *FrameworkItems) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *FrameworkItems) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetIsCcf

`func (o *FrameworkItems) GetIsCcf() bool`

GetIsCcf returns the IsCcf field if non-nil, zero value otherwise.

### GetIsCcfOk

`func (o *FrameworkItems) GetIsCcfOk() (*bool, bool)`

GetIsCcfOk returns a tuple with the IsCcf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCcf

`func (o *FrameworkItems) SetIsCcf(v bool)`

SetIsCcf sets IsCcf field to given value.


### GetType

`func (o *FrameworkItems) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FrameworkItems) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FrameworkItems) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FrameworkItems) HasType() bool`

HasType returns a boolean if a field has been set.

### GetClassification

`func (o *FrameworkItems) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *FrameworkItems) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *FrameworkItems) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *FrameworkItems) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetMinimumRequirementLevel

`func (o *FrameworkItems) GetMinimumRequirementLevel() string`

GetMinimumRequirementLevel returns the MinimumRequirementLevel field if non-nil, zero value otherwise.

### GetMinimumRequirementLevelOk

`func (o *FrameworkItems) GetMinimumRequirementLevelOk() (*string, bool)`

GetMinimumRequirementLevelOk returns a tuple with the MinimumRequirementLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRequirementLevel

`func (o *FrameworkItems) SetMinimumRequirementLevel(v string)`

SetMinimumRequirementLevel sets MinimumRequirementLevel field to given value.


### GetFieldData

`func (o *FrameworkItems) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *FrameworkItems) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *FrameworkItems) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *FrameworkItems) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *FrameworkItems) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *FrameworkItems) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil
### GetResponse

`func (o *FrameworkItems) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *FrameworkItems) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *FrameworkItems) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *FrameworkItems) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetIdCode

`func (o *FrameworkItems) GetIdCode() string`

GetIdCode returns the IdCode field if non-nil, zero value otherwise.

### GetIdCodeOk

`func (o *FrameworkItems) GetIdCodeOk() (*string, bool)`

GetIdCodeOk returns a tuple with the IdCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdCode

`func (o *FrameworkItems) SetIdCode(v string)`

SetIdCode sets IdCode field to given value.

### HasIdCode

`func (o *FrameworkItems) HasIdCode() bool`

HasIdCode returns a boolean if a field has been set.

### GetShortIdCode

`func (o *FrameworkItems) GetShortIdCode() string`

GetShortIdCode returns the ShortIdCode field if non-nil, zero value otherwise.

### GetShortIdCodeOk

`func (o *FrameworkItems) GetShortIdCodeOk() (*string, bool)`

GetShortIdCodeOk returns a tuple with the ShortIdCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortIdCode

`func (o *FrameworkItems) SetShortIdCode(v string)`

SetShortIdCode sets ShortIdCode field to given value.

### HasShortIdCode

`func (o *FrameworkItems) HasShortIdCode() bool`

HasShortIdCode returns a boolean if a field has been set.

### GetScopes

`func (o *FrameworkItems) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *FrameworkItems) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *FrameworkItems) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *FrameworkItems) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *FrameworkItems) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *FrameworkItems) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetTargetMaturityScore

`func (o *FrameworkItems) GetTargetMaturityScore() int32`

GetTargetMaturityScore returns the TargetMaturityScore field if non-nil, zero value otherwise.

### GetTargetMaturityScoreOk

`func (o *FrameworkItems) GetTargetMaturityScoreOk() (*int32, bool)`

GetTargetMaturityScoreOk returns a tuple with the TargetMaturityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMaturityScore

`func (o *FrameworkItems) SetTargetMaturityScore(v int32)`

SetTargetMaturityScore sets TargetMaturityScore field to given value.

### HasTargetMaturityScore

`func (o *FrameworkItems) HasTargetMaturityScore() bool`

HasTargetMaturityScore returns a boolean if a field has been set.

### GetXcTotals

`func (o *FrameworkItems) GetXcTotals() interface{}`

GetXcTotals returns the XcTotals field if non-nil, zero value otherwise.

### GetXcTotalsOk

`func (o *FrameworkItems) GetXcTotalsOk() (*interface{}, bool)`

GetXcTotalsOk returns a tuple with the XcTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXcTotals

`func (o *FrameworkItems) SetXcTotals(v interface{})`

SetXcTotals sets XcTotals field to given value.

### HasXcTotals

`func (o *FrameworkItems) HasXcTotals() bool`

HasXcTotals returns a boolean if a field has been set.

### SetXcTotalsNil

`func (o *FrameworkItems) SetXcTotalsNil(b bool)`

 SetXcTotalsNil sets the value for XcTotals to be an explicit nil

### UnsetXcTotals
`func (o *FrameworkItems) UnsetXcTotals()`

UnsetXcTotals ensures that no value is present for XcTotals, not even an explicit nil
### GetAtRiskStatus

`func (o *FrameworkItems) GetAtRiskStatus() string`

GetAtRiskStatus returns the AtRiskStatus field if non-nil, zero value otherwise.

### GetAtRiskStatusOk

`func (o *FrameworkItems) GetAtRiskStatusOk() (*string, bool)`

GetAtRiskStatusOk returns a tuple with the AtRiskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtRiskStatus

`func (o *FrameworkItems) SetAtRiskStatus(v string)`

SetAtRiskStatus sets AtRiskStatus field to given value.


### GetFrameworkImplementationGuidance

`func (o *FrameworkItems) GetFrameworkImplementationGuidance() string`

GetFrameworkImplementationGuidance returns the FrameworkImplementationGuidance field if non-nil, zero value otherwise.

### GetFrameworkImplementationGuidanceOk

`func (o *FrameworkItems) GetFrameworkImplementationGuidanceOk() (*string, bool)`

GetFrameworkImplementationGuidanceOk returns a tuple with the FrameworkImplementationGuidance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkImplementationGuidance

`func (o *FrameworkItems) SetFrameworkImplementationGuidance(v string)`

SetFrameworkImplementationGuidance sets FrameworkImplementationGuidance field to given value.

### HasFrameworkImplementationGuidance

`func (o *FrameworkItems) HasFrameworkImplementationGuidance() bool`

HasFrameworkImplementationGuidance returns a boolean if a field has been set.

### GetFrameworkAssessmentProcedures

`func (o *FrameworkItems) GetFrameworkAssessmentProcedures() string`

GetFrameworkAssessmentProcedures returns the FrameworkAssessmentProcedures field if non-nil, zero value otherwise.

### GetFrameworkAssessmentProceduresOk

`func (o *FrameworkItems) GetFrameworkAssessmentProceduresOk() (*string, bool)`

GetFrameworkAssessmentProceduresOk returns a tuple with the FrameworkAssessmentProcedures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkAssessmentProcedures

`func (o *FrameworkItems) SetFrameworkAssessmentProcedures(v string)`

SetFrameworkAssessmentProcedures sets FrameworkAssessmentProcedures field to given value.

### HasFrameworkAssessmentProcedures

`func (o *FrameworkItems) HasFrameworkAssessmentProcedures() bool`

HasFrameworkAssessmentProcedures returns a boolean if a field has been set.

### GetFrameworkRequirementPurpose

`func (o *FrameworkItems) GetFrameworkRequirementPurpose() string`

GetFrameworkRequirementPurpose returns the FrameworkRequirementPurpose field if non-nil, zero value otherwise.

### GetFrameworkRequirementPurposeOk

`func (o *FrameworkItems) GetFrameworkRequirementPurposeOk() (*string, bool)`

GetFrameworkRequirementPurposeOk returns a tuple with the FrameworkRequirementPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkRequirementPurpose

`func (o *FrameworkItems) SetFrameworkRequirementPurpose(v string)`

SetFrameworkRequirementPurpose sets FrameworkRequirementPurpose field to given value.

### HasFrameworkRequirementPurpose

`func (o *FrameworkItems) HasFrameworkRequirementPurpose() bool`

HasFrameworkRequirementPurpose returns a boolean if a field has been set.

### GetFrameworkItemFunctionId

`func (o *FrameworkItems) GetFrameworkItemFunctionId() int32`

GetFrameworkItemFunctionId returns the FrameworkItemFunctionId field if non-nil, zero value otherwise.

### GetFrameworkItemFunctionIdOk

`func (o *FrameworkItems) GetFrameworkItemFunctionIdOk() (*int32, bool)`

GetFrameworkItemFunctionIdOk returns a tuple with the FrameworkItemFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemFunctionId

`func (o *FrameworkItems) SetFrameworkItemFunctionId(v int32)`

SetFrameworkItemFunctionId sets FrameworkItemFunctionId field to given value.

### HasFrameworkItemFunctionId

`func (o *FrameworkItems) HasFrameworkItemFunctionId() bool`

HasFrameworkItemFunctionId returns a boolean if a field has been set.

### GetRelativeControlWeighting

`func (o *FrameworkItems) GetRelativeControlWeighting() string`

GetRelativeControlWeighting returns the RelativeControlWeighting field if non-nil, zero value otherwise.

### GetRelativeControlWeightingOk

`func (o *FrameworkItems) GetRelativeControlWeightingOk() (*string, bool)`

GetRelativeControlWeightingOk returns a tuple with the RelativeControlWeighting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeControlWeighting

`func (o *FrameworkItems) SetRelativeControlWeighting(v string)`

SetRelativeControlWeighting sets RelativeControlWeighting field to given value.

### HasRelativeControlWeighting

`func (o *FrameworkItems) HasRelativeControlWeighting() bool`

HasRelativeControlWeighting returns a boolean if a field has been set.

### GetCmmLevel0

`func (o *FrameworkItems) GetCmmLevel0() string`

GetCmmLevel0 returns the CmmLevel0 field if non-nil, zero value otherwise.

### GetCmmLevel0Ok

`func (o *FrameworkItems) GetCmmLevel0Ok() (*string, bool)`

GetCmmLevel0Ok returns a tuple with the CmmLevel0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmmLevel0

`func (o *FrameworkItems) SetCmmLevel0(v string)`

SetCmmLevel0 sets CmmLevel0 field to given value.

### HasCmmLevel0

`func (o *FrameworkItems) HasCmmLevel0() bool`

HasCmmLevel0 returns a boolean if a field has been set.

### GetCmmLevel1

`func (o *FrameworkItems) GetCmmLevel1() string`

GetCmmLevel1 returns the CmmLevel1 field if non-nil, zero value otherwise.

### GetCmmLevel1Ok

`func (o *FrameworkItems) GetCmmLevel1Ok() (*string, bool)`

GetCmmLevel1Ok returns a tuple with the CmmLevel1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmmLevel1

`func (o *FrameworkItems) SetCmmLevel1(v string)`

SetCmmLevel1 sets CmmLevel1 field to given value.

### HasCmmLevel1

`func (o *FrameworkItems) HasCmmLevel1() bool`

HasCmmLevel1 returns a boolean if a field has been set.

### GetCmmLevel2

`func (o *FrameworkItems) GetCmmLevel2() string`

GetCmmLevel2 returns the CmmLevel2 field if non-nil, zero value otherwise.

### GetCmmLevel2Ok

`func (o *FrameworkItems) GetCmmLevel2Ok() (*string, bool)`

GetCmmLevel2Ok returns a tuple with the CmmLevel2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmmLevel2

`func (o *FrameworkItems) SetCmmLevel2(v string)`

SetCmmLevel2 sets CmmLevel2 field to given value.

### HasCmmLevel2

`func (o *FrameworkItems) HasCmmLevel2() bool`

HasCmmLevel2 returns a boolean if a field has been set.

### GetCmmLevel3

`func (o *FrameworkItems) GetCmmLevel3() string`

GetCmmLevel3 returns the CmmLevel3 field if non-nil, zero value otherwise.

### GetCmmLevel3Ok

`func (o *FrameworkItems) GetCmmLevel3Ok() (*string, bool)`

GetCmmLevel3Ok returns a tuple with the CmmLevel3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmmLevel3

`func (o *FrameworkItems) SetCmmLevel3(v string)`

SetCmmLevel3 sets CmmLevel3 field to given value.

### HasCmmLevel3

`func (o *FrameworkItems) HasCmmLevel3() bool`

HasCmmLevel3 returns a boolean if a field has been set.

### GetCmmLevel4

`func (o *FrameworkItems) GetCmmLevel4() string`

GetCmmLevel4 returns the CmmLevel4 field if non-nil, zero value otherwise.

### GetCmmLevel4Ok

`func (o *FrameworkItems) GetCmmLevel4Ok() (*string, bool)`

GetCmmLevel4Ok returns a tuple with the CmmLevel4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmmLevel4

`func (o *FrameworkItems) SetCmmLevel4(v string)`

SetCmmLevel4 sets CmmLevel4 field to given value.

### HasCmmLevel4

`func (o *FrameworkItems) HasCmmLevel4() bool`

HasCmmLevel4 returns a boolean if a field has been set.

### GetCmmLevel5

`func (o *FrameworkItems) GetCmmLevel5() string`

GetCmmLevel5 returns the CmmLevel5 field if non-nil, zero value otherwise.

### GetCmmLevel5Ok

`func (o *FrameworkItems) GetCmmLevel5Ok() (*string, bool)`

GetCmmLevel5Ok returns a tuple with the CmmLevel5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmmLevel5

`func (o *FrameworkItems) SetCmmLevel5(v string)`

SetCmmLevel5 sets CmmLevel5 field to given value.

### HasCmmLevel5

`func (o *FrameworkItems) HasCmmLevel5() bool`

HasCmmLevel5 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


