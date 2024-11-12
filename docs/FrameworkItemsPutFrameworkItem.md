# FrameworkItemsPutFrameworkItem

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

### NewFrameworkItemsPutFrameworkItem

`func NewFrameworkItemsPutFrameworkItem(status string, referenceMeta interface{}, isCcf bool, minimumRequirementLevel string, atRiskStatus string, ) *FrameworkItemsPutFrameworkItem`

NewFrameworkItemsPutFrameworkItem instantiates a new FrameworkItemsPutFrameworkItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameworkItemsPutFrameworkItemWithDefaults

`func NewFrameworkItemsPutFrameworkItemWithDefaults() *FrameworkItemsPutFrameworkItem`

NewFrameworkItemsPutFrameworkItemWithDefaults instantiates a new FrameworkItemsPutFrameworkItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FrameworkItemsPutFrameworkItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FrameworkItemsPutFrameworkItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FrameworkItemsPutFrameworkItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FrameworkItemsPutFrameworkItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFrameworkId

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkId() int32`

GetFrameworkId returns the FrameworkId field if non-nil, zero value otherwise.

### GetFrameworkIdOk

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkIdOk() (*int32, bool)`

GetFrameworkIdOk returns a tuple with the FrameworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkId

`func (o *FrameworkItemsPutFrameworkItem) SetFrameworkId(v int32)`

SetFrameworkId sets FrameworkId field to given value.

### HasFrameworkId

`func (o *FrameworkItemsPutFrameworkItem) HasFrameworkId() bool`

HasFrameworkId returns a boolean if a field has been set.

### GetSection

`func (o *FrameworkItemsPutFrameworkItem) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *FrameworkItemsPutFrameworkItem) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *FrameworkItemsPutFrameworkItem) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *FrameworkItemsPutFrameworkItem) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetSubsection

`func (o *FrameworkItemsPutFrameworkItem) GetSubsection() string`

GetSubsection returns the Subsection field if non-nil, zero value otherwise.

### GetSubsectionOk

`func (o *FrameworkItemsPutFrameworkItem) GetSubsectionOk() (*string, bool)`

GetSubsectionOk returns a tuple with the Subsection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsection

`func (o *FrameworkItemsPutFrameworkItem) SetSubsection(v string)`

SetSubsection sets Subsection field to given value.

### HasSubsection

`func (o *FrameworkItemsPutFrameworkItem) HasSubsection() bool`

HasSubsection returns a boolean if a field has been set.

### GetUid

`func (o *FrameworkItemsPutFrameworkItem) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *FrameworkItemsPutFrameworkItem) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *FrameworkItemsPutFrameworkItem) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *FrameworkItemsPutFrameworkItem) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetName

`func (o *FrameworkItemsPutFrameworkItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FrameworkItemsPutFrameworkItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FrameworkItemsPutFrameworkItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FrameworkItemsPutFrameworkItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FrameworkItemsPutFrameworkItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FrameworkItemsPutFrameworkItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FrameworkItemsPutFrameworkItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FrameworkItemsPutFrameworkItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *FrameworkItemsPutFrameworkItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FrameworkItemsPutFrameworkItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FrameworkItemsPutFrameworkItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSortOrder

`func (o *FrameworkItemsPutFrameworkItem) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *FrameworkItemsPutFrameworkItem) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *FrameworkItemsPutFrameworkItem) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *FrameworkItemsPutFrameworkItem) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FrameworkItemsPutFrameworkItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FrameworkItemsPutFrameworkItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FrameworkItemsPutFrameworkItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FrameworkItemsPutFrameworkItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FrameworkItemsPutFrameworkItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FrameworkItemsPutFrameworkItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FrameworkItemsPutFrameworkItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FrameworkItemsPutFrameworkItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *FrameworkItemsPutFrameworkItem) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *FrameworkItemsPutFrameworkItem) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *FrameworkItemsPutFrameworkItem) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *FrameworkItemsPutFrameworkItem) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetScope

`func (o *FrameworkItemsPutFrameworkItem) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *FrameworkItemsPutFrameworkItem) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *FrameworkItemsPutFrameworkItem) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *FrameworkItemsPutFrameworkItem) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetCustomDate1

`func (o *FrameworkItemsPutFrameworkItem) GetCustomDate1() string`

GetCustomDate1 returns the CustomDate1 field if non-nil, zero value otherwise.

### GetCustomDate1Ok

`func (o *FrameworkItemsPutFrameworkItem) GetCustomDate1Ok() (*string, bool)`

GetCustomDate1Ok returns a tuple with the CustomDate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate1

`func (o *FrameworkItemsPutFrameworkItem) SetCustomDate1(v string)`

SetCustomDate1 sets CustomDate1 field to given value.

### HasCustomDate1

`func (o *FrameworkItemsPutFrameworkItem) HasCustomDate1() bool`

HasCustomDate1 returns a boolean if a field has been set.

### GetCustomDate2

`func (o *FrameworkItemsPutFrameworkItem) GetCustomDate2() string`

GetCustomDate2 returns the CustomDate2 field if non-nil, zero value otherwise.

### GetCustomDate2Ok

`func (o *FrameworkItemsPutFrameworkItem) GetCustomDate2Ok() (*string, bool)`

GetCustomDate2Ok returns a tuple with the CustomDate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate2

`func (o *FrameworkItemsPutFrameworkItem) SetCustomDate2(v string)`

SetCustomDate2 sets CustomDate2 field to given value.

### HasCustomDate2

`func (o *FrameworkItemsPutFrameworkItem) HasCustomDate2() bool`

HasCustomDate2 returns a boolean if a field has been set.

### GetCustomDate3

`func (o *FrameworkItemsPutFrameworkItem) GetCustomDate3() string`

GetCustomDate3 returns the CustomDate3 field if non-nil, zero value otherwise.

### GetCustomDate3Ok

`func (o *FrameworkItemsPutFrameworkItem) GetCustomDate3Ok() (*string, bool)`

GetCustomDate3Ok returns a tuple with the CustomDate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate3

`func (o *FrameworkItemsPutFrameworkItem) SetCustomDate3(v string)`

SetCustomDate3 sets CustomDate3 field to given value.

### HasCustomDate3

`func (o *FrameworkItemsPutFrameworkItem) HasCustomDate3() bool`

HasCustomDate3 returns a boolean if a field has been set.

### GetCustomText1

`func (o *FrameworkItemsPutFrameworkItem) GetCustomText1() string`

GetCustomText1 returns the CustomText1 field if non-nil, zero value otherwise.

### GetCustomText1Ok

`func (o *FrameworkItemsPutFrameworkItem) GetCustomText1Ok() (*string, bool)`

GetCustomText1Ok returns a tuple with the CustomText1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText1

`func (o *FrameworkItemsPutFrameworkItem) SetCustomText1(v string)`

SetCustomText1 sets CustomText1 field to given value.

### HasCustomText1

`func (o *FrameworkItemsPutFrameworkItem) HasCustomText1() bool`

HasCustomText1 returns a boolean if a field has been set.

### GetCustomText2

`func (o *FrameworkItemsPutFrameworkItem) GetCustomText2() string`

GetCustomText2 returns the CustomText2 field if non-nil, zero value otherwise.

### GetCustomText2Ok

`func (o *FrameworkItemsPutFrameworkItem) GetCustomText2Ok() (*string, bool)`

GetCustomText2Ok returns a tuple with the CustomText2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText2

`func (o *FrameworkItemsPutFrameworkItem) SetCustomText2(v string)`

SetCustomText2 sets CustomText2 field to given value.

### HasCustomText2

`func (o *FrameworkItemsPutFrameworkItem) HasCustomText2() bool`

HasCustomText2 returns a boolean if a field has been set.

### GetCustomText3

`func (o *FrameworkItemsPutFrameworkItem) GetCustomText3() string`

GetCustomText3 returns the CustomText3 field if non-nil, zero value otherwise.

### GetCustomText3Ok

`func (o *FrameworkItemsPutFrameworkItem) GetCustomText3Ok() (*string, bool)`

GetCustomText3Ok returns a tuple with the CustomText3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText3

`func (o *FrameworkItemsPutFrameworkItem) SetCustomText3(v string)`

SetCustomText3 sets CustomText3 field to given value.

### HasCustomText3

`func (o *FrameworkItemsPutFrameworkItem) HasCustomText3() bool`

HasCustomText3 returns a boolean if a field has been set.

### GetFrameworkItemCustomSelect1OptionId

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkItemCustomSelect1OptionId() int32`

GetFrameworkItemCustomSelect1OptionId returns the FrameworkItemCustomSelect1OptionId field if non-nil, zero value otherwise.

### GetFrameworkItemCustomSelect1OptionIdOk

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkItemCustomSelect1OptionIdOk() (*int32, bool)`

GetFrameworkItemCustomSelect1OptionIdOk returns a tuple with the FrameworkItemCustomSelect1OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemCustomSelect1OptionId

`func (o *FrameworkItemsPutFrameworkItem) SetFrameworkItemCustomSelect1OptionId(v int32)`

SetFrameworkItemCustomSelect1OptionId sets FrameworkItemCustomSelect1OptionId field to given value.

### HasFrameworkItemCustomSelect1OptionId

`func (o *FrameworkItemsPutFrameworkItem) HasFrameworkItemCustomSelect1OptionId() bool`

HasFrameworkItemCustomSelect1OptionId returns a boolean if a field has been set.

### GetFrameworkItemCustomSelect2OptionId

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkItemCustomSelect2OptionId() int32`

GetFrameworkItemCustomSelect2OptionId returns the FrameworkItemCustomSelect2OptionId field if non-nil, zero value otherwise.

### GetFrameworkItemCustomSelect2OptionIdOk

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkItemCustomSelect2OptionIdOk() (*int32, bool)`

GetFrameworkItemCustomSelect2OptionIdOk returns a tuple with the FrameworkItemCustomSelect2OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemCustomSelect2OptionId

`func (o *FrameworkItemsPutFrameworkItem) SetFrameworkItemCustomSelect2OptionId(v int32)`

SetFrameworkItemCustomSelect2OptionId sets FrameworkItemCustomSelect2OptionId field to given value.

### HasFrameworkItemCustomSelect2OptionId

`func (o *FrameworkItemsPutFrameworkItem) HasFrameworkItemCustomSelect2OptionId() bool`

HasFrameworkItemCustomSelect2OptionId returns a boolean if a field has been set.

### GetFrameworkItemCustomSelect3OptionId

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkItemCustomSelect3OptionId() int32`

GetFrameworkItemCustomSelect3OptionId returns the FrameworkItemCustomSelect3OptionId field if non-nil, zero value otherwise.

### GetFrameworkItemCustomSelect3OptionIdOk

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkItemCustomSelect3OptionIdOk() (*int32, bool)`

GetFrameworkItemCustomSelect3OptionIdOk returns a tuple with the FrameworkItemCustomSelect3OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemCustomSelect3OptionId

`func (o *FrameworkItemsPutFrameworkItem) SetFrameworkItemCustomSelect3OptionId(v int32)`

SetFrameworkItemCustomSelect3OptionId sets FrameworkItemCustomSelect3OptionId field to given value.

### HasFrameworkItemCustomSelect3OptionId

`func (o *FrameworkItemsPutFrameworkItem) HasFrameworkItemCustomSelect3OptionId() bool`

HasFrameworkItemCustomSelect3OptionId returns a boolean if a field has been set.

### GetRequirementOwnerUserId

`func (o *FrameworkItemsPutFrameworkItem) GetRequirementOwnerUserId() int32`

GetRequirementOwnerUserId returns the RequirementOwnerUserId field if non-nil, zero value otherwise.

### GetRequirementOwnerUserIdOk

`func (o *FrameworkItemsPutFrameworkItem) GetRequirementOwnerUserIdOk() (*int32, bool)`

GetRequirementOwnerUserIdOk returns a tuple with the RequirementOwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementOwnerUserId

`func (o *FrameworkItemsPutFrameworkItem) SetRequirementOwnerUserId(v int32)`

SetRequirementOwnerUserId sets RequirementOwnerUserId field to given value.

### HasRequirementOwnerUserId

`func (o *FrameworkItemsPutFrameworkItem) HasRequirementOwnerUserId() bool`

HasRequirementOwnerUserId returns a boolean if a field has been set.

### GetRequirementReviewerUserId

`func (o *FrameworkItemsPutFrameworkItem) GetRequirementReviewerUserId() int32`

GetRequirementReviewerUserId returns the RequirementReviewerUserId field if non-nil, zero value otherwise.

### GetRequirementReviewerUserIdOk

`func (o *FrameworkItemsPutFrameworkItem) GetRequirementReviewerUserIdOk() (*int32, bool)`

GetRequirementReviewerUserIdOk returns a tuple with the RequirementReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementReviewerUserId

`func (o *FrameworkItemsPutFrameworkItem) SetRequirementReviewerUserId(v int32)`

SetRequirementReviewerUserId sets RequirementReviewerUserId field to given value.

### HasRequirementReviewerUserId

`func (o *FrameworkItemsPutFrameworkItem) HasRequirementReviewerUserId() bool`

HasRequirementReviewerUserId returns a boolean if a field has been set.

### GetRequest

`func (o *FrameworkItemsPutFrameworkItem) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *FrameworkItemsPutFrameworkItem) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *FrameworkItemsPutFrameworkItem) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *FrameworkItemsPutFrameworkItem) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetAdditionalRequest

`func (o *FrameworkItemsPutFrameworkItem) GetAdditionalRequest() string`

GetAdditionalRequest returns the AdditionalRequest field if non-nil, zero value otherwise.

### GetAdditionalRequestOk

`func (o *FrameworkItemsPutFrameworkItem) GetAdditionalRequestOk() (*string, bool)`

GetAdditionalRequestOk returns a tuple with the AdditionalRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalRequest

`func (o *FrameworkItemsPutFrameworkItem) SetAdditionalRequest(v string)`

SetAdditionalRequest sets AdditionalRequest field to given value.

### HasAdditionalRequest

`func (o *FrameworkItemsPutFrameworkItem) HasAdditionalRequest() bool`

HasAdditionalRequest returns a boolean if a field has been set.

### GetFrameworkItemCustomSelect4OptionId

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkItemCustomSelect4OptionId() int32`

GetFrameworkItemCustomSelect4OptionId returns the FrameworkItemCustomSelect4OptionId field if non-nil, zero value otherwise.

### GetFrameworkItemCustomSelect4OptionIdOk

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkItemCustomSelect4OptionIdOk() (*int32, bool)`

GetFrameworkItemCustomSelect4OptionIdOk returns a tuple with the FrameworkItemCustomSelect4OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemCustomSelect4OptionId

`func (o *FrameworkItemsPutFrameworkItem) SetFrameworkItemCustomSelect4OptionId(v int32)`

SetFrameworkItemCustomSelect4OptionId sets FrameworkItemCustomSelect4OptionId field to given value.

### HasFrameworkItemCustomSelect4OptionId

`func (o *FrameworkItemsPutFrameworkItem) HasFrameworkItemCustomSelect4OptionId() bool`

HasFrameworkItemCustomSelect4OptionId returns a boolean if a field has been set.

### GetFrameworkItemCustomSelect5OptionId

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkItemCustomSelect5OptionId() int32`

GetFrameworkItemCustomSelect5OptionId returns the FrameworkItemCustomSelect5OptionId field if non-nil, zero value otherwise.

### GetFrameworkItemCustomSelect5OptionIdOk

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkItemCustomSelect5OptionIdOk() (*int32, bool)`

GetFrameworkItemCustomSelect5OptionIdOk returns a tuple with the FrameworkItemCustomSelect5OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemCustomSelect5OptionId

`func (o *FrameworkItemsPutFrameworkItem) SetFrameworkItemCustomSelect5OptionId(v int32)`

SetFrameworkItemCustomSelect5OptionId sets FrameworkItemCustomSelect5OptionId field to given value.

### HasFrameworkItemCustomSelect5OptionId

`func (o *FrameworkItemsPutFrameworkItem) HasFrameworkItemCustomSelect5OptionId() bool`

HasFrameworkItemCustomSelect5OptionId returns a boolean if a field has been set.

### GetImplementationGuidance

`func (o *FrameworkItemsPutFrameworkItem) GetImplementationGuidance() string`

GetImplementationGuidance returns the ImplementationGuidance field if non-nil, zero value otherwise.

### GetImplementationGuidanceOk

`func (o *FrameworkItemsPutFrameworkItem) GetImplementationGuidanceOk() (*string, bool)`

GetImplementationGuidanceOk returns a tuple with the ImplementationGuidance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationGuidance

`func (o *FrameworkItemsPutFrameworkItem) SetImplementationGuidance(v string)`

SetImplementationGuidance sets ImplementationGuidance field to given value.

### HasImplementationGuidance

`func (o *FrameworkItemsPutFrameworkItem) HasImplementationGuidance() bool`

HasImplementationGuidance returns a boolean if a field has been set.

### GetFrameworkItemImplementationStatusId

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkItemImplementationStatusId() int32`

GetFrameworkItemImplementationStatusId returns the FrameworkItemImplementationStatusId field if non-nil, zero value otherwise.

### GetFrameworkItemImplementationStatusIdOk

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkItemImplementationStatusIdOk() (*int32, bool)`

GetFrameworkItemImplementationStatusIdOk returns a tuple with the FrameworkItemImplementationStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemImplementationStatusId

`func (o *FrameworkItemsPutFrameworkItem) SetFrameworkItemImplementationStatusId(v int32)`

SetFrameworkItemImplementationStatusId sets FrameworkItemImplementationStatusId field to given value.

### HasFrameworkItemImplementationStatusId

`func (o *FrameworkItemsPutFrameworkItem) HasFrameworkItemImplementationStatusId() bool`

HasFrameworkItemImplementationStatusId returns a boolean if a field has been set.

### GetFrameworkItemRatingId

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkItemRatingId() int32`

GetFrameworkItemRatingId returns the FrameworkItemRatingId field if non-nil, zero value otherwise.

### GetFrameworkItemRatingIdOk

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkItemRatingIdOk() (*int32, bool)`

GetFrameworkItemRatingIdOk returns a tuple with the FrameworkItemRatingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemRatingId

`func (o *FrameworkItemsPutFrameworkItem) SetFrameworkItemRatingId(v int32)`

SetFrameworkItemRatingId sets FrameworkItemRatingId field to given value.

### HasFrameworkItemRatingId

`func (o *FrameworkItemsPutFrameworkItem) HasFrameworkItemRatingId() bool`

HasFrameworkItemRatingId returns a boolean if a field has been set.

### GetCustomText4

`func (o *FrameworkItemsPutFrameworkItem) GetCustomText4() string`

GetCustomText4 returns the CustomText4 field if non-nil, zero value otherwise.

### GetCustomText4Ok

`func (o *FrameworkItemsPutFrameworkItem) GetCustomText4Ok() (*string, bool)`

GetCustomText4Ok returns a tuple with the CustomText4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText4

`func (o *FrameworkItemsPutFrameworkItem) SetCustomText4(v string)`

SetCustomText4 sets CustomText4 field to given value.

### HasCustomText4

`func (o *FrameworkItemsPutFrameworkItem) HasCustomText4() bool`

HasCustomText4 returns a boolean if a field has been set.

### GetCustomText5

`func (o *FrameworkItemsPutFrameworkItem) GetCustomText5() string`

GetCustomText5 returns the CustomText5 field if non-nil, zero value otherwise.

### GetCustomText5Ok

`func (o *FrameworkItemsPutFrameworkItem) GetCustomText5Ok() (*string, bool)`

GetCustomText5Ok returns a tuple with the CustomText5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText5

`func (o *FrameworkItemsPutFrameworkItem) SetCustomText5(v string)`

SetCustomText5 sets CustomText5 field to given value.

### HasCustomText5

`func (o *FrameworkItemsPutFrameworkItem) HasCustomText5() bool`

HasCustomText5 returns a boolean if a field has been set.

### GetCustomText6

`func (o *FrameworkItemsPutFrameworkItem) GetCustomText6() string`

GetCustomText6 returns the CustomText6 field if non-nil, zero value otherwise.

### GetCustomText6Ok

`func (o *FrameworkItemsPutFrameworkItem) GetCustomText6Ok() (*string, bool)`

GetCustomText6Ok returns a tuple with the CustomText6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText6

`func (o *FrameworkItemsPutFrameworkItem) SetCustomText6(v string)`

SetCustomText6 sets CustomText6 field to given value.

### HasCustomText6

`func (o *FrameworkItemsPutFrameworkItem) HasCustomText6() bool`

HasCustomText6 returns a boolean if a field has been set.

### GetReferenceMeta

`func (o *FrameworkItemsPutFrameworkItem) GetReferenceMeta() interface{}`

GetReferenceMeta returns the ReferenceMeta field if non-nil, zero value otherwise.

### GetReferenceMetaOk

`func (o *FrameworkItemsPutFrameworkItem) GetReferenceMetaOk() (*interface{}, bool)`

GetReferenceMetaOk returns a tuple with the ReferenceMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMeta

`func (o *FrameworkItemsPutFrameworkItem) SetReferenceMeta(v interface{})`

SetReferenceMeta sets ReferenceMeta field to given value.


### SetReferenceMetaNil

`func (o *FrameworkItemsPutFrameworkItem) SetReferenceMetaNil(b bool)`

 SetReferenceMetaNil sets the value for ReferenceMeta to be an explicit nil

### UnsetReferenceMeta
`func (o *FrameworkItemsPutFrameworkItem) UnsetReferenceMeta()`

UnsetReferenceMeta ensures that no value is present for ReferenceMeta, not even an explicit nil
### GetCoordinatorUserId

`func (o *FrameworkItemsPutFrameworkItem) GetCoordinatorUserId() int32`

GetCoordinatorUserId returns the CoordinatorUserId field if non-nil, zero value otherwise.

### GetCoordinatorUserIdOk

`func (o *FrameworkItemsPutFrameworkItem) GetCoordinatorUserIdOk() (*int32, bool)`

GetCoordinatorUserIdOk returns a tuple with the CoordinatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinatorUserId

`func (o *FrameworkItemsPutFrameworkItem) SetCoordinatorUserId(v int32)`

SetCoordinatorUserId sets CoordinatorUserId field to given value.

### HasCoordinatorUserId

`func (o *FrameworkItemsPutFrameworkItem) HasCoordinatorUserId() bool`

HasCoordinatorUserId returns a boolean if a field has been set.

### GetFrameworkItemActualMaturityId

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkItemActualMaturityId() int32`

GetFrameworkItemActualMaturityId returns the FrameworkItemActualMaturityId field if non-nil, zero value otherwise.

### GetFrameworkItemActualMaturityIdOk

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkItemActualMaturityIdOk() (*int32, bool)`

GetFrameworkItemActualMaturityIdOk returns a tuple with the FrameworkItemActualMaturityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemActualMaturityId

`func (o *FrameworkItemsPutFrameworkItem) SetFrameworkItemActualMaturityId(v int32)`

SetFrameworkItemActualMaturityId sets FrameworkItemActualMaturityId field to given value.

### HasFrameworkItemActualMaturityId

`func (o *FrameworkItemsPutFrameworkItem) HasFrameworkItemActualMaturityId() bool`

HasFrameworkItemActualMaturityId returns a boolean if a field has been set.

### GetFrameworkItemTargetMaturityId

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkItemTargetMaturityId() int32`

GetFrameworkItemTargetMaturityId returns the FrameworkItemTargetMaturityId field if non-nil, zero value otherwise.

### GetFrameworkItemTargetMaturityIdOk

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkItemTargetMaturityIdOk() (*int32, bool)`

GetFrameworkItemTargetMaturityIdOk returns a tuple with the FrameworkItemTargetMaturityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemTargetMaturityId

`func (o *FrameworkItemsPutFrameworkItem) SetFrameworkItemTargetMaturityId(v int32)`

SetFrameworkItemTargetMaturityId sets FrameworkItemTargetMaturityId field to given value.

### HasFrameworkItemTargetMaturityId

`func (o *FrameworkItemsPutFrameworkItem) HasFrameworkItemTargetMaturityId() bool`

HasFrameworkItemTargetMaturityId returns a boolean if a field has been set.

### GetContentServiceId

`func (o *FrameworkItemsPutFrameworkItem) GetContentServiceId() string`

GetContentServiceId returns the ContentServiceId field if non-nil, zero value otherwise.

### GetContentServiceIdOk

`func (o *FrameworkItemsPutFrameworkItem) GetContentServiceIdOk() (*string, bool)`

GetContentServiceIdOk returns a tuple with the ContentServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentServiceId

`func (o *FrameworkItemsPutFrameworkItem) SetContentServiceId(v string)`

SetContentServiceId sets ContentServiceId field to given value.

### HasContentServiceId

`func (o *FrameworkItemsPutFrameworkItem) HasContentServiceId() bool`

HasContentServiceId returns a boolean if a field has been set.

### GetDescriptionWithTags

`func (o *FrameworkItemsPutFrameworkItem) GetDescriptionWithTags() string`

GetDescriptionWithTags returns the DescriptionWithTags field if non-nil, zero value otherwise.

### GetDescriptionWithTagsOk

`func (o *FrameworkItemsPutFrameworkItem) GetDescriptionWithTagsOk() (*string, bool)`

GetDescriptionWithTagsOk returns a tuple with the DescriptionWithTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionWithTags

`func (o *FrameworkItemsPutFrameworkItem) SetDescriptionWithTags(v string)`

SetDescriptionWithTags sets DescriptionWithTags field to given value.

### HasDescriptionWithTags

`func (o *FrameworkItemsPutFrameworkItem) HasDescriptionWithTags() bool`

HasDescriptionWithTags returns a boolean if a field has been set.

### GetParentId

`func (o *FrameworkItemsPutFrameworkItem) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *FrameworkItemsPutFrameworkItem) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *FrameworkItemsPutFrameworkItem) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *FrameworkItemsPutFrameworkItem) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetIsCcf

`func (o *FrameworkItemsPutFrameworkItem) GetIsCcf() bool`

GetIsCcf returns the IsCcf field if non-nil, zero value otherwise.

### GetIsCcfOk

`func (o *FrameworkItemsPutFrameworkItem) GetIsCcfOk() (*bool, bool)`

GetIsCcfOk returns a tuple with the IsCcf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCcf

`func (o *FrameworkItemsPutFrameworkItem) SetIsCcf(v bool)`

SetIsCcf sets IsCcf field to given value.


### GetType

`func (o *FrameworkItemsPutFrameworkItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FrameworkItemsPutFrameworkItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FrameworkItemsPutFrameworkItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FrameworkItemsPutFrameworkItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetClassification

`func (o *FrameworkItemsPutFrameworkItem) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *FrameworkItemsPutFrameworkItem) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *FrameworkItemsPutFrameworkItem) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *FrameworkItemsPutFrameworkItem) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetMinimumRequirementLevel

`func (o *FrameworkItemsPutFrameworkItem) GetMinimumRequirementLevel() string`

GetMinimumRequirementLevel returns the MinimumRequirementLevel field if non-nil, zero value otherwise.

### GetMinimumRequirementLevelOk

`func (o *FrameworkItemsPutFrameworkItem) GetMinimumRequirementLevelOk() (*string, bool)`

GetMinimumRequirementLevelOk returns a tuple with the MinimumRequirementLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRequirementLevel

`func (o *FrameworkItemsPutFrameworkItem) SetMinimumRequirementLevel(v string)`

SetMinimumRequirementLevel sets MinimumRequirementLevel field to given value.


### GetFieldData

`func (o *FrameworkItemsPutFrameworkItem) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *FrameworkItemsPutFrameworkItem) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *FrameworkItemsPutFrameworkItem) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *FrameworkItemsPutFrameworkItem) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *FrameworkItemsPutFrameworkItem) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *FrameworkItemsPutFrameworkItem) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil
### GetResponse

`func (o *FrameworkItemsPutFrameworkItem) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *FrameworkItemsPutFrameworkItem) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *FrameworkItemsPutFrameworkItem) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *FrameworkItemsPutFrameworkItem) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetIdCode

`func (o *FrameworkItemsPutFrameworkItem) GetIdCode() string`

GetIdCode returns the IdCode field if non-nil, zero value otherwise.

### GetIdCodeOk

`func (o *FrameworkItemsPutFrameworkItem) GetIdCodeOk() (*string, bool)`

GetIdCodeOk returns a tuple with the IdCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdCode

`func (o *FrameworkItemsPutFrameworkItem) SetIdCode(v string)`

SetIdCode sets IdCode field to given value.

### HasIdCode

`func (o *FrameworkItemsPutFrameworkItem) HasIdCode() bool`

HasIdCode returns a boolean if a field has been set.

### GetShortIdCode

`func (o *FrameworkItemsPutFrameworkItem) GetShortIdCode() string`

GetShortIdCode returns the ShortIdCode field if non-nil, zero value otherwise.

### GetShortIdCodeOk

`func (o *FrameworkItemsPutFrameworkItem) GetShortIdCodeOk() (*string, bool)`

GetShortIdCodeOk returns a tuple with the ShortIdCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortIdCode

`func (o *FrameworkItemsPutFrameworkItem) SetShortIdCode(v string)`

SetShortIdCode sets ShortIdCode field to given value.

### HasShortIdCode

`func (o *FrameworkItemsPutFrameworkItem) HasShortIdCode() bool`

HasShortIdCode returns a boolean if a field has been set.

### GetScopes

`func (o *FrameworkItemsPutFrameworkItem) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *FrameworkItemsPutFrameworkItem) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *FrameworkItemsPutFrameworkItem) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *FrameworkItemsPutFrameworkItem) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *FrameworkItemsPutFrameworkItem) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *FrameworkItemsPutFrameworkItem) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetTargetMaturityScore

`func (o *FrameworkItemsPutFrameworkItem) GetTargetMaturityScore() int32`

GetTargetMaturityScore returns the TargetMaturityScore field if non-nil, zero value otherwise.

### GetTargetMaturityScoreOk

`func (o *FrameworkItemsPutFrameworkItem) GetTargetMaturityScoreOk() (*int32, bool)`

GetTargetMaturityScoreOk returns a tuple with the TargetMaturityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMaturityScore

`func (o *FrameworkItemsPutFrameworkItem) SetTargetMaturityScore(v int32)`

SetTargetMaturityScore sets TargetMaturityScore field to given value.

### HasTargetMaturityScore

`func (o *FrameworkItemsPutFrameworkItem) HasTargetMaturityScore() bool`

HasTargetMaturityScore returns a boolean if a field has been set.

### GetXcTotals

`func (o *FrameworkItemsPutFrameworkItem) GetXcTotals() interface{}`

GetXcTotals returns the XcTotals field if non-nil, zero value otherwise.

### GetXcTotalsOk

`func (o *FrameworkItemsPutFrameworkItem) GetXcTotalsOk() (*interface{}, bool)`

GetXcTotalsOk returns a tuple with the XcTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXcTotals

`func (o *FrameworkItemsPutFrameworkItem) SetXcTotals(v interface{})`

SetXcTotals sets XcTotals field to given value.

### HasXcTotals

`func (o *FrameworkItemsPutFrameworkItem) HasXcTotals() bool`

HasXcTotals returns a boolean if a field has been set.

### SetXcTotalsNil

`func (o *FrameworkItemsPutFrameworkItem) SetXcTotalsNil(b bool)`

 SetXcTotalsNil sets the value for XcTotals to be an explicit nil

### UnsetXcTotals
`func (o *FrameworkItemsPutFrameworkItem) UnsetXcTotals()`

UnsetXcTotals ensures that no value is present for XcTotals, not even an explicit nil
### GetAtRiskStatus

`func (o *FrameworkItemsPutFrameworkItem) GetAtRiskStatus() string`

GetAtRiskStatus returns the AtRiskStatus field if non-nil, zero value otherwise.

### GetAtRiskStatusOk

`func (o *FrameworkItemsPutFrameworkItem) GetAtRiskStatusOk() (*string, bool)`

GetAtRiskStatusOk returns a tuple with the AtRiskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtRiskStatus

`func (o *FrameworkItemsPutFrameworkItem) SetAtRiskStatus(v string)`

SetAtRiskStatus sets AtRiskStatus field to given value.


### GetFrameworkImplementationGuidance

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkImplementationGuidance() string`

GetFrameworkImplementationGuidance returns the FrameworkImplementationGuidance field if non-nil, zero value otherwise.

### GetFrameworkImplementationGuidanceOk

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkImplementationGuidanceOk() (*string, bool)`

GetFrameworkImplementationGuidanceOk returns a tuple with the FrameworkImplementationGuidance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkImplementationGuidance

`func (o *FrameworkItemsPutFrameworkItem) SetFrameworkImplementationGuidance(v string)`

SetFrameworkImplementationGuidance sets FrameworkImplementationGuidance field to given value.

### HasFrameworkImplementationGuidance

`func (o *FrameworkItemsPutFrameworkItem) HasFrameworkImplementationGuidance() bool`

HasFrameworkImplementationGuidance returns a boolean if a field has been set.

### GetFrameworkAssessmentProcedures

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkAssessmentProcedures() string`

GetFrameworkAssessmentProcedures returns the FrameworkAssessmentProcedures field if non-nil, zero value otherwise.

### GetFrameworkAssessmentProceduresOk

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkAssessmentProceduresOk() (*string, bool)`

GetFrameworkAssessmentProceduresOk returns a tuple with the FrameworkAssessmentProcedures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkAssessmentProcedures

`func (o *FrameworkItemsPutFrameworkItem) SetFrameworkAssessmentProcedures(v string)`

SetFrameworkAssessmentProcedures sets FrameworkAssessmentProcedures field to given value.

### HasFrameworkAssessmentProcedures

`func (o *FrameworkItemsPutFrameworkItem) HasFrameworkAssessmentProcedures() bool`

HasFrameworkAssessmentProcedures returns a boolean if a field has been set.

### GetFrameworkRequirementPurpose

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkRequirementPurpose() string`

GetFrameworkRequirementPurpose returns the FrameworkRequirementPurpose field if non-nil, zero value otherwise.

### GetFrameworkRequirementPurposeOk

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkRequirementPurposeOk() (*string, bool)`

GetFrameworkRequirementPurposeOk returns a tuple with the FrameworkRequirementPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkRequirementPurpose

`func (o *FrameworkItemsPutFrameworkItem) SetFrameworkRequirementPurpose(v string)`

SetFrameworkRequirementPurpose sets FrameworkRequirementPurpose field to given value.

### HasFrameworkRequirementPurpose

`func (o *FrameworkItemsPutFrameworkItem) HasFrameworkRequirementPurpose() bool`

HasFrameworkRequirementPurpose returns a boolean if a field has been set.

### GetFrameworkItemFunctionId

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkItemFunctionId() int32`

GetFrameworkItemFunctionId returns the FrameworkItemFunctionId field if non-nil, zero value otherwise.

### GetFrameworkItemFunctionIdOk

`func (o *FrameworkItemsPutFrameworkItem) GetFrameworkItemFunctionIdOk() (*int32, bool)`

GetFrameworkItemFunctionIdOk returns a tuple with the FrameworkItemFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkItemFunctionId

`func (o *FrameworkItemsPutFrameworkItem) SetFrameworkItemFunctionId(v int32)`

SetFrameworkItemFunctionId sets FrameworkItemFunctionId field to given value.

### HasFrameworkItemFunctionId

`func (o *FrameworkItemsPutFrameworkItem) HasFrameworkItemFunctionId() bool`

HasFrameworkItemFunctionId returns a boolean if a field has been set.

### GetRelativeControlWeighting

`func (o *FrameworkItemsPutFrameworkItem) GetRelativeControlWeighting() string`

GetRelativeControlWeighting returns the RelativeControlWeighting field if non-nil, zero value otherwise.

### GetRelativeControlWeightingOk

`func (o *FrameworkItemsPutFrameworkItem) GetRelativeControlWeightingOk() (*string, bool)`

GetRelativeControlWeightingOk returns a tuple with the RelativeControlWeighting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeControlWeighting

`func (o *FrameworkItemsPutFrameworkItem) SetRelativeControlWeighting(v string)`

SetRelativeControlWeighting sets RelativeControlWeighting field to given value.

### HasRelativeControlWeighting

`func (o *FrameworkItemsPutFrameworkItem) HasRelativeControlWeighting() bool`

HasRelativeControlWeighting returns a boolean if a field has been set.

### GetCmmLevel0

`func (o *FrameworkItemsPutFrameworkItem) GetCmmLevel0() string`

GetCmmLevel0 returns the CmmLevel0 field if non-nil, zero value otherwise.

### GetCmmLevel0Ok

`func (o *FrameworkItemsPutFrameworkItem) GetCmmLevel0Ok() (*string, bool)`

GetCmmLevel0Ok returns a tuple with the CmmLevel0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmmLevel0

`func (o *FrameworkItemsPutFrameworkItem) SetCmmLevel0(v string)`

SetCmmLevel0 sets CmmLevel0 field to given value.

### HasCmmLevel0

`func (o *FrameworkItemsPutFrameworkItem) HasCmmLevel0() bool`

HasCmmLevel0 returns a boolean if a field has been set.

### GetCmmLevel1

`func (o *FrameworkItemsPutFrameworkItem) GetCmmLevel1() string`

GetCmmLevel1 returns the CmmLevel1 field if non-nil, zero value otherwise.

### GetCmmLevel1Ok

`func (o *FrameworkItemsPutFrameworkItem) GetCmmLevel1Ok() (*string, bool)`

GetCmmLevel1Ok returns a tuple with the CmmLevel1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmmLevel1

`func (o *FrameworkItemsPutFrameworkItem) SetCmmLevel1(v string)`

SetCmmLevel1 sets CmmLevel1 field to given value.

### HasCmmLevel1

`func (o *FrameworkItemsPutFrameworkItem) HasCmmLevel1() bool`

HasCmmLevel1 returns a boolean if a field has been set.

### GetCmmLevel2

`func (o *FrameworkItemsPutFrameworkItem) GetCmmLevel2() string`

GetCmmLevel2 returns the CmmLevel2 field if non-nil, zero value otherwise.

### GetCmmLevel2Ok

`func (o *FrameworkItemsPutFrameworkItem) GetCmmLevel2Ok() (*string, bool)`

GetCmmLevel2Ok returns a tuple with the CmmLevel2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmmLevel2

`func (o *FrameworkItemsPutFrameworkItem) SetCmmLevel2(v string)`

SetCmmLevel2 sets CmmLevel2 field to given value.

### HasCmmLevel2

`func (o *FrameworkItemsPutFrameworkItem) HasCmmLevel2() bool`

HasCmmLevel2 returns a boolean if a field has been set.

### GetCmmLevel3

`func (o *FrameworkItemsPutFrameworkItem) GetCmmLevel3() string`

GetCmmLevel3 returns the CmmLevel3 field if non-nil, zero value otherwise.

### GetCmmLevel3Ok

`func (o *FrameworkItemsPutFrameworkItem) GetCmmLevel3Ok() (*string, bool)`

GetCmmLevel3Ok returns a tuple with the CmmLevel3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmmLevel3

`func (o *FrameworkItemsPutFrameworkItem) SetCmmLevel3(v string)`

SetCmmLevel3 sets CmmLevel3 field to given value.

### HasCmmLevel3

`func (o *FrameworkItemsPutFrameworkItem) HasCmmLevel3() bool`

HasCmmLevel3 returns a boolean if a field has been set.

### GetCmmLevel4

`func (o *FrameworkItemsPutFrameworkItem) GetCmmLevel4() string`

GetCmmLevel4 returns the CmmLevel4 field if non-nil, zero value otherwise.

### GetCmmLevel4Ok

`func (o *FrameworkItemsPutFrameworkItem) GetCmmLevel4Ok() (*string, bool)`

GetCmmLevel4Ok returns a tuple with the CmmLevel4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmmLevel4

`func (o *FrameworkItemsPutFrameworkItem) SetCmmLevel4(v string)`

SetCmmLevel4 sets CmmLevel4 field to given value.

### HasCmmLevel4

`func (o *FrameworkItemsPutFrameworkItem) HasCmmLevel4() bool`

HasCmmLevel4 returns a boolean if a field has been set.

### GetCmmLevel5

`func (o *FrameworkItemsPutFrameworkItem) GetCmmLevel5() string`

GetCmmLevel5 returns the CmmLevel5 field if non-nil, zero value otherwise.

### GetCmmLevel5Ok

`func (o *FrameworkItemsPutFrameworkItem) GetCmmLevel5Ok() (*string, bool)`

GetCmmLevel5Ok returns a tuple with the CmmLevel5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmmLevel5

`func (o *FrameworkItemsPutFrameworkItem) SetCmmLevel5(v string)`

SetCmmLevel5 sets CmmLevel5 field to given value.

### HasCmmLevel5

`func (o *FrameworkItemsPutFrameworkItem) HasCmmLevel5() bool`

HasCmmLevel5 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


