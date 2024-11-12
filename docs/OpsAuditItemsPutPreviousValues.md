# OpsAuditItemsPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**OpsAuditId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audits.id&#x60;.&lt;fk table&#x3D;&#39;ops_audits&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditSectionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_sections.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_sections&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**PreparerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ReviewerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [default to "Open"]
**Notes** | Pointer to **string** |  | [optional] 
**Results** | Pointer to **string** |  | [optional] 
**SecondaryReviewerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**SortOrder** | Pointer to **int32** |  | [optional] [default to 0]
**Title** | Pointer to **string** |  | [optional] 
**Procedures** | Pointer to **string** |  | [optional] [default to ""]
**CompleteDate** | Pointer to **string** |  | [optional] 
**CompleteByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**UnderReview1Date** | Pointer to **string** |  | [optional] 
**UnderReview1ByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**UnderReview2Date** | Pointer to **string** |  | [optional] 
**UnderReview2ByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ReviewedDate** | Pointer to **string** |  | [optional] 
**ReviewedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Objective** | Pointer to **string** |  | [optional] [default to ""]
**Source** | Pointer to **string** |  | [optional] [default to ""]
**Scope** | Pointer to **string** |  | [optional] [default to ""]
**OpsRiskLevelId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_risk_levels.id&#x60;.&lt;fk table&#x3D;&#39;ops_risk_levels&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsRatingId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_ratings.id&#x60;.&lt;fk table&#x3D;&#39;ops_ratings&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**PbcRequest** | Pointer to **string** |  | [optional] 
**PbcRequestAdditional** | Pointer to **string** |  | [optional] 
**PopulationDetails** | Pointer to **string** |  | [optional] 
**SampleSize** | Pointer to **string** |  | [optional] 
**SampleSelection** | Pointer to **string** |  | [optional] 
**OpsCustomSelect1OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_custom_select1_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_custom_select1_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsCustomSelect2OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_custom_select2_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_custom_select2_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsCustomDate1** | Pointer to **string** |  | [optional] 
**OpsCustomDate2** | Pointer to **string** |  | [optional] 
**OpsCustomDate3** | Pointer to **string** |  | [optional] 
**Effectiveness** | Pointer to **string** |  | [optional] [default to "Not Tested"]
**OpsAuditSectionInstanceId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_section_instances.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_section_instances&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditSubsectionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_subsections.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_subsections&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomText1** | Pointer to **string** |  | [optional] 
**CustomText2** | Pointer to **string** |  | [optional] 
**CustomText3** | Pointer to **string** |  | [optional] 
**OpsCustomSelect3OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_custom_select3_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_custom_select3_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsCustomSelect4OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_custom_select4_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_custom_select4_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsCustomDate4** | Pointer to **string** |  | [optional] 
**OpsCustomDate5** | Pointer to **string** |  | [optional] 
**OpsCustomDate6** | Pointer to **string** |  | [optional] 
**OpsCustomDate7** | Pointer to **string** |  | [optional] 
**OwnerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomText4** | Pointer to **string** |  | [optional] 
**CustomText5** | Pointer to **string** |  | [optional] 
**CustomText6** | Pointer to **string** |  | [optional] 
**OpsCustomSelect5OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_custom_select5_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_custom_select5_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsCustomSelect6OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_custom_select6_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_custom_select6_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomText7** | Pointer to **string** |  | [optional] 
**CustomText8** | Pointer to **string** |  | [optional] 
**CustomText9** | Pointer to **string** |  | [optional] 
**CustomText10** | Pointer to **string** |  | [optional] 
**CustomText11** | Pointer to **string** |  | [optional] 
**CustomText12** | Pointer to **string** |  | [optional] 
**OpsCustomSelect7OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_custom_select7_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_custom_select7_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsCustomSelect8OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_custom_select8_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_custom_select8_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsCustomSelect9OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_custom_select9_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_custom_select9_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsCustomSelect10OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_custom_select10_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_custom_select10_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsCustomSelect11OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_custom_select11_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_custom_select11_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsCustomSelect12OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_custom_select12_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_custom_select12_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsCustomSelect13OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_custom_select13_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_custom_select13_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsCustomSelect14OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_custom_select14_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_custom_select14_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsCustomSelect15OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_custom_select15_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_custom_select15_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomText13** | Pointer to **string** |  | [optional] 
**CustomText14** | Pointer to **string** |  | [optional] 
**CustomText15** | Pointer to **string** |  | [optional] 
**OpsCustomSelect16OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_custom_select16_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_custom_select16_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsCustomSelect17OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_custom_select17_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_custom_select17_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsFrequencyOptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_frequency_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_frequency_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Totals** | Pointer to **interface{}** |  | [optional] 
**FieldData** | Pointer to **interface{}** |  | [optional] 
**FirstCompleteDate** | Pointer to **string** |  | [optional] 
**FirstReviewedDate** | Pointer to **string** |  | [optional] 
**FirstUnderReview1Date** | Pointer to **string** |  | [optional] 
**FirstUnderReview2Date** | Pointer to **string** |  | [optional] 
**Flattened** | Pointer to **interface{}** |  | [optional] 
**IsFlattened** | Pointer to **bool** |  | [optional] [default to false]
**BudgetHours** | Pointer to **float32** |  | [optional] 
**ActualHours** | Pointer to **float32** |  | [optional] 
**IdString** | Pointer to **string** |  | [optional] 
**ScopeRationale** | Pointer to **string** |  | [optional] 

## Methods

### NewOpsAuditItemsPutPreviousValues

`func NewOpsAuditItemsPutPreviousValues() *OpsAuditItemsPutPreviousValues`

NewOpsAuditItemsPutPreviousValues instantiates a new OpsAuditItemsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsAuditItemsPutPreviousValuesWithDefaults

`func NewOpsAuditItemsPutPreviousValuesWithDefaults() *OpsAuditItemsPutPreviousValues`

NewOpsAuditItemsPutPreviousValuesWithDefaults instantiates a new OpsAuditItemsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpsAuditItemsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpsAuditItemsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OpsAuditItemsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OpsAuditItemsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OpsAuditItemsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OpsAuditItemsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OpsAuditItemsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OpsAuditItemsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OpsAuditItemsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OpsAuditItemsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OpsAuditItemsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *OpsAuditItemsPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *OpsAuditItemsPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *OpsAuditItemsPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *OpsAuditItemsPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetOpsAuditId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsAuditId() int32`

GetOpsAuditId returns the OpsAuditId field if non-nil, zero value otherwise.

### GetOpsAuditIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsAuditIdOk() (*int32, bool)`

GetOpsAuditIdOk returns a tuple with the OpsAuditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsAuditId(v int32)`

SetOpsAuditId sets OpsAuditId field to given value.

### HasOpsAuditId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsAuditId() bool`

HasOpsAuditId returns a boolean if a field has been set.

### GetOpsAuditSectionId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsAuditSectionId() int32`

GetOpsAuditSectionId returns the OpsAuditSectionId field if non-nil, zero value otherwise.

### GetOpsAuditSectionIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsAuditSectionIdOk() (*int32, bool)`

GetOpsAuditSectionIdOk returns a tuple with the OpsAuditSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditSectionId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsAuditSectionId(v int32)`

SetOpsAuditSectionId sets OpsAuditSectionId field to given value.

### HasOpsAuditSectionId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsAuditSectionId() bool`

HasOpsAuditSectionId returns a boolean if a field has been set.

### GetPreparerUserId

`func (o *OpsAuditItemsPutPreviousValues) GetPreparerUserId() int32`

GetPreparerUserId returns the PreparerUserId field if non-nil, zero value otherwise.

### GetPreparerUserIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetPreparerUserIdOk() (*int32, bool)`

GetPreparerUserIdOk returns a tuple with the PreparerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparerUserId

`func (o *OpsAuditItemsPutPreviousValues) SetPreparerUserId(v int32)`

SetPreparerUserId sets PreparerUserId field to given value.

### HasPreparerUserId

`func (o *OpsAuditItemsPutPreviousValues) HasPreparerUserId() bool`

HasPreparerUserId returns a boolean if a field has been set.

### GetReviewerUserId

`func (o *OpsAuditItemsPutPreviousValues) GetReviewerUserId() int32`

GetReviewerUserId returns the ReviewerUserId field if non-nil, zero value otherwise.

### GetReviewerUserIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetReviewerUserIdOk() (*int32, bool)`

GetReviewerUserIdOk returns a tuple with the ReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerUserId

`func (o *OpsAuditItemsPutPreviousValues) SetReviewerUserId(v int32)`

SetReviewerUserId sets ReviewerUserId field to given value.

### HasReviewerUserId

`func (o *OpsAuditItemsPutPreviousValues) HasReviewerUserId() bool`

HasReviewerUserId returns a boolean if a field has been set.

### GetDueDate

`func (o *OpsAuditItemsPutPreviousValues) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *OpsAuditItemsPutPreviousValues) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *OpsAuditItemsPutPreviousValues) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *OpsAuditItemsPutPreviousValues) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetDescription

`func (o *OpsAuditItemsPutPreviousValues) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpsAuditItemsPutPreviousValues) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpsAuditItemsPutPreviousValues) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OpsAuditItemsPutPreviousValues) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *OpsAuditItemsPutPreviousValues) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpsAuditItemsPutPreviousValues) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpsAuditItemsPutPreviousValues) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OpsAuditItemsPutPreviousValues) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNotes

`func (o *OpsAuditItemsPutPreviousValues) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OpsAuditItemsPutPreviousValues) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OpsAuditItemsPutPreviousValues) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OpsAuditItemsPutPreviousValues) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetResults

`func (o *OpsAuditItemsPutPreviousValues) GetResults() string`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *OpsAuditItemsPutPreviousValues) GetResultsOk() (*string, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *OpsAuditItemsPutPreviousValues) SetResults(v string)`

SetResults sets Results field to given value.

### HasResults

`func (o *OpsAuditItemsPutPreviousValues) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetSecondaryReviewerUserId

`func (o *OpsAuditItemsPutPreviousValues) GetSecondaryReviewerUserId() int32`

GetSecondaryReviewerUserId returns the SecondaryReviewerUserId field if non-nil, zero value otherwise.

### GetSecondaryReviewerUserIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetSecondaryReviewerUserIdOk() (*int32, bool)`

GetSecondaryReviewerUserIdOk returns a tuple with the SecondaryReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryReviewerUserId

`func (o *OpsAuditItemsPutPreviousValues) SetSecondaryReviewerUserId(v int32)`

SetSecondaryReviewerUserId sets SecondaryReviewerUserId field to given value.

### HasSecondaryReviewerUserId

`func (o *OpsAuditItemsPutPreviousValues) HasSecondaryReviewerUserId() bool`

HasSecondaryReviewerUserId returns a boolean if a field has been set.

### GetSortOrder

`func (o *OpsAuditItemsPutPreviousValues) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *OpsAuditItemsPutPreviousValues) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *OpsAuditItemsPutPreviousValues) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *OpsAuditItemsPutPreviousValues) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetTitle

`func (o *OpsAuditItemsPutPreviousValues) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OpsAuditItemsPutPreviousValues) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OpsAuditItemsPutPreviousValues) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *OpsAuditItemsPutPreviousValues) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetProcedures

`func (o *OpsAuditItemsPutPreviousValues) GetProcedures() string`

GetProcedures returns the Procedures field if non-nil, zero value otherwise.

### GetProceduresOk

`func (o *OpsAuditItemsPutPreviousValues) GetProceduresOk() (*string, bool)`

GetProceduresOk returns a tuple with the Procedures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcedures

`func (o *OpsAuditItemsPutPreviousValues) SetProcedures(v string)`

SetProcedures sets Procedures field to given value.

### HasProcedures

`func (o *OpsAuditItemsPutPreviousValues) HasProcedures() bool`

HasProcedures returns a boolean if a field has been set.

### GetCompleteDate

`func (o *OpsAuditItemsPutPreviousValues) GetCompleteDate() string`

GetCompleteDate returns the CompleteDate field if non-nil, zero value otherwise.

### GetCompleteDateOk

`func (o *OpsAuditItemsPutPreviousValues) GetCompleteDateOk() (*string, bool)`

GetCompleteDateOk returns a tuple with the CompleteDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteDate

`func (o *OpsAuditItemsPutPreviousValues) SetCompleteDate(v string)`

SetCompleteDate sets CompleteDate field to given value.

### HasCompleteDate

`func (o *OpsAuditItemsPutPreviousValues) HasCompleteDate() bool`

HasCompleteDate returns a boolean if a field has been set.

### GetCompleteByUserId

`func (o *OpsAuditItemsPutPreviousValues) GetCompleteByUserId() int32`

GetCompleteByUserId returns the CompleteByUserId field if non-nil, zero value otherwise.

### GetCompleteByUserIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetCompleteByUserIdOk() (*int32, bool)`

GetCompleteByUserIdOk returns a tuple with the CompleteByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteByUserId

`func (o *OpsAuditItemsPutPreviousValues) SetCompleteByUserId(v int32)`

SetCompleteByUserId sets CompleteByUserId field to given value.

### HasCompleteByUserId

`func (o *OpsAuditItemsPutPreviousValues) HasCompleteByUserId() bool`

HasCompleteByUserId returns a boolean if a field has been set.

### GetUnderReview1Date

`func (o *OpsAuditItemsPutPreviousValues) GetUnderReview1Date() string`

GetUnderReview1Date returns the UnderReview1Date field if non-nil, zero value otherwise.

### GetUnderReview1DateOk

`func (o *OpsAuditItemsPutPreviousValues) GetUnderReview1DateOk() (*string, bool)`

GetUnderReview1DateOk returns a tuple with the UnderReview1Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderReview1Date

`func (o *OpsAuditItemsPutPreviousValues) SetUnderReview1Date(v string)`

SetUnderReview1Date sets UnderReview1Date field to given value.

### HasUnderReview1Date

`func (o *OpsAuditItemsPutPreviousValues) HasUnderReview1Date() bool`

HasUnderReview1Date returns a boolean if a field has been set.

### GetUnderReview1ByUserId

`func (o *OpsAuditItemsPutPreviousValues) GetUnderReview1ByUserId() int32`

GetUnderReview1ByUserId returns the UnderReview1ByUserId field if non-nil, zero value otherwise.

### GetUnderReview1ByUserIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetUnderReview1ByUserIdOk() (*int32, bool)`

GetUnderReview1ByUserIdOk returns a tuple with the UnderReview1ByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderReview1ByUserId

`func (o *OpsAuditItemsPutPreviousValues) SetUnderReview1ByUserId(v int32)`

SetUnderReview1ByUserId sets UnderReview1ByUserId field to given value.

### HasUnderReview1ByUserId

`func (o *OpsAuditItemsPutPreviousValues) HasUnderReview1ByUserId() bool`

HasUnderReview1ByUserId returns a boolean if a field has been set.

### GetUnderReview2Date

`func (o *OpsAuditItemsPutPreviousValues) GetUnderReview2Date() string`

GetUnderReview2Date returns the UnderReview2Date field if non-nil, zero value otherwise.

### GetUnderReview2DateOk

`func (o *OpsAuditItemsPutPreviousValues) GetUnderReview2DateOk() (*string, bool)`

GetUnderReview2DateOk returns a tuple with the UnderReview2Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderReview2Date

`func (o *OpsAuditItemsPutPreviousValues) SetUnderReview2Date(v string)`

SetUnderReview2Date sets UnderReview2Date field to given value.

### HasUnderReview2Date

`func (o *OpsAuditItemsPutPreviousValues) HasUnderReview2Date() bool`

HasUnderReview2Date returns a boolean if a field has been set.

### GetUnderReview2ByUserId

`func (o *OpsAuditItemsPutPreviousValues) GetUnderReview2ByUserId() int32`

GetUnderReview2ByUserId returns the UnderReview2ByUserId field if non-nil, zero value otherwise.

### GetUnderReview2ByUserIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetUnderReview2ByUserIdOk() (*int32, bool)`

GetUnderReview2ByUserIdOk returns a tuple with the UnderReview2ByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderReview2ByUserId

`func (o *OpsAuditItemsPutPreviousValues) SetUnderReview2ByUserId(v int32)`

SetUnderReview2ByUserId sets UnderReview2ByUserId field to given value.

### HasUnderReview2ByUserId

`func (o *OpsAuditItemsPutPreviousValues) HasUnderReview2ByUserId() bool`

HasUnderReview2ByUserId returns a boolean if a field has been set.

### GetReviewedDate

`func (o *OpsAuditItemsPutPreviousValues) GetReviewedDate() string`

GetReviewedDate returns the ReviewedDate field if non-nil, zero value otherwise.

### GetReviewedDateOk

`func (o *OpsAuditItemsPutPreviousValues) GetReviewedDateOk() (*string, bool)`

GetReviewedDateOk returns a tuple with the ReviewedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedDate

`func (o *OpsAuditItemsPutPreviousValues) SetReviewedDate(v string)`

SetReviewedDate sets ReviewedDate field to given value.

### HasReviewedDate

`func (o *OpsAuditItemsPutPreviousValues) HasReviewedDate() bool`

HasReviewedDate returns a boolean if a field has been set.

### GetReviewedByUserId

`func (o *OpsAuditItemsPutPreviousValues) GetReviewedByUserId() int32`

GetReviewedByUserId returns the ReviewedByUserId field if non-nil, zero value otherwise.

### GetReviewedByUserIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetReviewedByUserIdOk() (*int32, bool)`

GetReviewedByUserIdOk returns a tuple with the ReviewedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedByUserId

`func (o *OpsAuditItemsPutPreviousValues) SetReviewedByUserId(v int32)`

SetReviewedByUserId sets ReviewedByUserId field to given value.

### HasReviewedByUserId

`func (o *OpsAuditItemsPutPreviousValues) HasReviewedByUserId() bool`

HasReviewedByUserId returns a boolean if a field has been set.

### GetObjective

`func (o *OpsAuditItemsPutPreviousValues) GetObjective() string`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *OpsAuditItemsPutPreviousValues) GetObjectiveOk() (*string, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *OpsAuditItemsPutPreviousValues) SetObjective(v string)`

SetObjective sets Objective field to given value.

### HasObjective

`func (o *OpsAuditItemsPutPreviousValues) HasObjective() bool`

HasObjective returns a boolean if a field has been set.

### GetSource

`func (o *OpsAuditItemsPutPreviousValues) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *OpsAuditItemsPutPreviousValues) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *OpsAuditItemsPutPreviousValues) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *OpsAuditItemsPutPreviousValues) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetScope

`func (o *OpsAuditItemsPutPreviousValues) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OpsAuditItemsPutPreviousValues) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OpsAuditItemsPutPreviousValues) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OpsAuditItemsPutPreviousValues) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetOpsRiskLevelId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsRiskLevelId() int32`

GetOpsRiskLevelId returns the OpsRiskLevelId field if non-nil, zero value otherwise.

### GetOpsRiskLevelIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsRiskLevelIdOk() (*int32, bool)`

GetOpsRiskLevelIdOk returns a tuple with the OpsRiskLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsRiskLevelId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsRiskLevelId(v int32)`

SetOpsRiskLevelId sets OpsRiskLevelId field to given value.

### HasOpsRiskLevelId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsRiskLevelId() bool`

HasOpsRiskLevelId returns a boolean if a field has been set.

### GetOpsRatingId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsRatingId() int32`

GetOpsRatingId returns the OpsRatingId field if non-nil, zero value otherwise.

### GetOpsRatingIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsRatingIdOk() (*int32, bool)`

GetOpsRatingIdOk returns a tuple with the OpsRatingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsRatingId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsRatingId(v int32)`

SetOpsRatingId sets OpsRatingId field to given value.

### HasOpsRatingId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsRatingId() bool`

HasOpsRatingId returns a boolean if a field has been set.

### GetPbcRequest

`func (o *OpsAuditItemsPutPreviousValues) GetPbcRequest() string`

GetPbcRequest returns the PbcRequest field if non-nil, zero value otherwise.

### GetPbcRequestOk

`func (o *OpsAuditItemsPutPreviousValues) GetPbcRequestOk() (*string, bool)`

GetPbcRequestOk returns a tuple with the PbcRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbcRequest

`func (o *OpsAuditItemsPutPreviousValues) SetPbcRequest(v string)`

SetPbcRequest sets PbcRequest field to given value.

### HasPbcRequest

`func (o *OpsAuditItemsPutPreviousValues) HasPbcRequest() bool`

HasPbcRequest returns a boolean if a field has been set.

### GetPbcRequestAdditional

`func (o *OpsAuditItemsPutPreviousValues) GetPbcRequestAdditional() string`

GetPbcRequestAdditional returns the PbcRequestAdditional field if non-nil, zero value otherwise.

### GetPbcRequestAdditionalOk

`func (o *OpsAuditItemsPutPreviousValues) GetPbcRequestAdditionalOk() (*string, bool)`

GetPbcRequestAdditionalOk returns a tuple with the PbcRequestAdditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbcRequestAdditional

`func (o *OpsAuditItemsPutPreviousValues) SetPbcRequestAdditional(v string)`

SetPbcRequestAdditional sets PbcRequestAdditional field to given value.

### HasPbcRequestAdditional

`func (o *OpsAuditItemsPutPreviousValues) HasPbcRequestAdditional() bool`

HasPbcRequestAdditional returns a boolean if a field has been set.

### GetPopulationDetails

`func (o *OpsAuditItemsPutPreviousValues) GetPopulationDetails() string`

GetPopulationDetails returns the PopulationDetails field if non-nil, zero value otherwise.

### GetPopulationDetailsOk

`func (o *OpsAuditItemsPutPreviousValues) GetPopulationDetailsOk() (*string, bool)`

GetPopulationDetailsOk returns a tuple with the PopulationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationDetails

`func (o *OpsAuditItemsPutPreviousValues) SetPopulationDetails(v string)`

SetPopulationDetails sets PopulationDetails field to given value.

### HasPopulationDetails

`func (o *OpsAuditItemsPutPreviousValues) HasPopulationDetails() bool`

HasPopulationDetails returns a boolean if a field has been set.

### GetSampleSize

`func (o *OpsAuditItemsPutPreviousValues) GetSampleSize() string`

GetSampleSize returns the SampleSize field if non-nil, zero value otherwise.

### GetSampleSizeOk

`func (o *OpsAuditItemsPutPreviousValues) GetSampleSizeOk() (*string, bool)`

GetSampleSizeOk returns a tuple with the SampleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleSize

`func (o *OpsAuditItemsPutPreviousValues) SetSampleSize(v string)`

SetSampleSize sets SampleSize field to given value.

### HasSampleSize

`func (o *OpsAuditItemsPutPreviousValues) HasSampleSize() bool`

HasSampleSize returns a boolean if a field has been set.

### GetSampleSelection

`func (o *OpsAuditItemsPutPreviousValues) GetSampleSelection() string`

GetSampleSelection returns the SampleSelection field if non-nil, zero value otherwise.

### GetSampleSelectionOk

`func (o *OpsAuditItemsPutPreviousValues) GetSampleSelectionOk() (*string, bool)`

GetSampleSelectionOk returns a tuple with the SampleSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleSelection

`func (o *OpsAuditItemsPutPreviousValues) SetSampleSelection(v string)`

SetSampleSelection sets SampleSelection field to given value.

### HasSampleSelection

`func (o *OpsAuditItemsPutPreviousValues) HasSampleSelection() bool`

HasSampleSelection returns a boolean if a field has been set.

### GetOpsCustomSelect1OptionId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect1OptionId() int32`

GetOpsCustomSelect1OptionId returns the OpsCustomSelect1OptionId field if non-nil, zero value otherwise.

### GetOpsCustomSelect1OptionIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect1OptionIdOk() (*int32, bool)`

GetOpsCustomSelect1OptionIdOk returns a tuple with the OpsCustomSelect1OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomSelect1OptionId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomSelect1OptionId(v int32)`

SetOpsCustomSelect1OptionId sets OpsCustomSelect1OptionId field to given value.

### HasOpsCustomSelect1OptionId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomSelect1OptionId() bool`

HasOpsCustomSelect1OptionId returns a boolean if a field has been set.

### GetOpsCustomSelect2OptionId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect2OptionId() int32`

GetOpsCustomSelect2OptionId returns the OpsCustomSelect2OptionId field if non-nil, zero value otherwise.

### GetOpsCustomSelect2OptionIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect2OptionIdOk() (*int32, bool)`

GetOpsCustomSelect2OptionIdOk returns a tuple with the OpsCustomSelect2OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomSelect2OptionId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomSelect2OptionId(v int32)`

SetOpsCustomSelect2OptionId sets OpsCustomSelect2OptionId field to given value.

### HasOpsCustomSelect2OptionId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomSelect2OptionId() bool`

HasOpsCustomSelect2OptionId returns a boolean if a field has been set.

### GetOpsCustomDate1

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomDate1() string`

GetOpsCustomDate1 returns the OpsCustomDate1 field if non-nil, zero value otherwise.

### GetOpsCustomDate1Ok

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomDate1Ok() (*string, bool)`

GetOpsCustomDate1Ok returns a tuple with the OpsCustomDate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomDate1

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomDate1(v string)`

SetOpsCustomDate1 sets OpsCustomDate1 field to given value.

### HasOpsCustomDate1

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomDate1() bool`

HasOpsCustomDate1 returns a boolean if a field has been set.

### GetOpsCustomDate2

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomDate2() string`

GetOpsCustomDate2 returns the OpsCustomDate2 field if non-nil, zero value otherwise.

### GetOpsCustomDate2Ok

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomDate2Ok() (*string, bool)`

GetOpsCustomDate2Ok returns a tuple with the OpsCustomDate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomDate2

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomDate2(v string)`

SetOpsCustomDate2 sets OpsCustomDate2 field to given value.

### HasOpsCustomDate2

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomDate2() bool`

HasOpsCustomDate2 returns a boolean if a field has been set.

### GetOpsCustomDate3

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomDate3() string`

GetOpsCustomDate3 returns the OpsCustomDate3 field if non-nil, zero value otherwise.

### GetOpsCustomDate3Ok

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomDate3Ok() (*string, bool)`

GetOpsCustomDate3Ok returns a tuple with the OpsCustomDate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomDate3

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomDate3(v string)`

SetOpsCustomDate3 sets OpsCustomDate3 field to given value.

### HasOpsCustomDate3

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomDate3() bool`

HasOpsCustomDate3 returns a boolean if a field has been set.

### GetEffectiveness

`func (o *OpsAuditItemsPutPreviousValues) GetEffectiveness() string`

GetEffectiveness returns the Effectiveness field if non-nil, zero value otherwise.

### GetEffectivenessOk

`func (o *OpsAuditItemsPutPreviousValues) GetEffectivenessOk() (*string, bool)`

GetEffectivenessOk returns a tuple with the Effectiveness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveness

`func (o *OpsAuditItemsPutPreviousValues) SetEffectiveness(v string)`

SetEffectiveness sets Effectiveness field to given value.

### HasEffectiveness

`func (o *OpsAuditItemsPutPreviousValues) HasEffectiveness() bool`

HasEffectiveness returns a boolean if a field has been set.

### GetOpsAuditSectionInstanceId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsAuditSectionInstanceId() int32`

GetOpsAuditSectionInstanceId returns the OpsAuditSectionInstanceId field if non-nil, zero value otherwise.

### GetOpsAuditSectionInstanceIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsAuditSectionInstanceIdOk() (*int32, bool)`

GetOpsAuditSectionInstanceIdOk returns a tuple with the OpsAuditSectionInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditSectionInstanceId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsAuditSectionInstanceId(v int32)`

SetOpsAuditSectionInstanceId sets OpsAuditSectionInstanceId field to given value.

### HasOpsAuditSectionInstanceId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsAuditSectionInstanceId() bool`

HasOpsAuditSectionInstanceId returns a boolean if a field has been set.

### GetOpsAuditSubsectionId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsAuditSubsectionId() int32`

GetOpsAuditSubsectionId returns the OpsAuditSubsectionId field if non-nil, zero value otherwise.

### GetOpsAuditSubsectionIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsAuditSubsectionIdOk() (*int32, bool)`

GetOpsAuditSubsectionIdOk returns a tuple with the OpsAuditSubsectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditSubsectionId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsAuditSubsectionId(v int32)`

SetOpsAuditSubsectionId sets OpsAuditSubsectionId field to given value.

### HasOpsAuditSubsectionId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsAuditSubsectionId() bool`

HasOpsAuditSubsectionId returns a boolean if a field has been set.

### GetCustomText1

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText1() string`

GetCustomText1 returns the CustomText1 field if non-nil, zero value otherwise.

### GetCustomText1Ok

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText1Ok() (*string, bool)`

GetCustomText1Ok returns a tuple with the CustomText1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText1

`func (o *OpsAuditItemsPutPreviousValues) SetCustomText1(v string)`

SetCustomText1 sets CustomText1 field to given value.

### HasCustomText1

`func (o *OpsAuditItemsPutPreviousValues) HasCustomText1() bool`

HasCustomText1 returns a boolean if a field has been set.

### GetCustomText2

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText2() string`

GetCustomText2 returns the CustomText2 field if non-nil, zero value otherwise.

### GetCustomText2Ok

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText2Ok() (*string, bool)`

GetCustomText2Ok returns a tuple with the CustomText2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText2

`func (o *OpsAuditItemsPutPreviousValues) SetCustomText2(v string)`

SetCustomText2 sets CustomText2 field to given value.

### HasCustomText2

`func (o *OpsAuditItemsPutPreviousValues) HasCustomText2() bool`

HasCustomText2 returns a boolean if a field has been set.

### GetCustomText3

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText3() string`

GetCustomText3 returns the CustomText3 field if non-nil, zero value otherwise.

### GetCustomText3Ok

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText3Ok() (*string, bool)`

GetCustomText3Ok returns a tuple with the CustomText3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText3

`func (o *OpsAuditItemsPutPreviousValues) SetCustomText3(v string)`

SetCustomText3 sets CustomText3 field to given value.

### HasCustomText3

`func (o *OpsAuditItemsPutPreviousValues) HasCustomText3() bool`

HasCustomText3 returns a boolean if a field has been set.

### GetOpsCustomSelect3OptionId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect3OptionId() int32`

GetOpsCustomSelect3OptionId returns the OpsCustomSelect3OptionId field if non-nil, zero value otherwise.

### GetOpsCustomSelect3OptionIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect3OptionIdOk() (*int32, bool)`

GetOpsCustomSelect3OptionIdOk returns a tuple with the OpsCustomSelect3OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomSelect3OptionId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomSelect3OptionId(v int32)`

SetOpsCustomSelect3OptionId sets OpsCustomSelect3OptionId field to given value.

### HasOpsCustomSelect3OptionId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomSelect3OptionId() bool`

HasOpsCustomSelect3OptionId returns a boolean if a field has been set.

### GetOpsCustomSelect4OptionId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect4OptionId() int32`

GetOpsCustomSelect4OptionId returns the OpsCustomSelect4OptionId field if non-nil, zero value otherwise.

### GetOpsCustomSelect4OptionIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect4OptionIdOk() (*int32, bool)`

GetOpsCustomSelect4OptionIdOk returns a tuple with the OpsCustomSelect4OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomSelect4OptionId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomSelect4OptionId(v int32)`

SetOpsCustomSelect4OptionId sets OpsCustomSelect4OptionId field to given value.

### HasOpsCustomSelect4OptionId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomSelect4OptionId() bool`

HasOpsCustomSelect4OptionId returns a boolean if a field has been set.

### GetOpsCustomDate4

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomDate4() string`

GetOpsCustomDate4 returns the OpsCustomDate4 field if non-nil, zero value otherwise.

### GetOpsCustomDate4Ok

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomDate4Ok() (*string, bool)`

GetOpsCustomDate4Ok returns a tuple with the OpsCustomDate4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomDate4

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomDate4(v string)`

SetOpsCustomDate4 sets OpsCustomDate4 field to given value.

### HasOpsCustomDate4

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomDate4() bool`

HasOpsCustomDate4 returns a boolean if a field has been set.

### GetOpsCustomDate5

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomDate5() string`

GetOpsCustomDate5 returns the OpsCustomDate5 field if non-nil, zero value otherwise.

### GetOpsCustomDate5Ok

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomDate5Ok() (*string, bool)`

GetOpsCustomDate5Ok returns a tuple with the OpsCustomDate5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomDate5

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomDate5(v string)`

SetOpsCustomDate5 sets OpsCustomDate5 field to given value.

### HasOpsCustomDate5

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomDate5() bool`

HasOpsCustomDate5 returns a boolean if a field has been set.

### GetOpsCustomDate6

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomDate6() string`

GetOpsCustomDate6 returns the OpsCustomDate6 field if non-nil, zero value otherwise.

### GetOpsCustomDate6Ok

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomDate6Ok() (*string, bool)`

GetOpsCustomDate6Ok returns a tuple with the OpsCustomDate6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomDate6

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomDate6(v string)`

SetOpsCustomDate6 sets OpsCustomDate6 field to given value.

### HasOpsCustomDate6

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomDate6() bool`

HasOpsCustomDate6 returns a boolean if a field has been set.

### GetOpsCustomDate7

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomDate7() string`

GetOpsCustomDate7 returns the OpsCustomDate7 field if non-nil, zero value otherwise.

### GetOpsCustomDate7Ok

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomDate7Ok() (*string, bool)`

GetOpsCustomDate7Ok returns a tuple with the OpsCustomDate7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomDate7

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomDate7(v string)`

SetOpsCustomDate7 sets OpsCustomDate7 field to given value.

### HasOpsCustomDate7

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomDate7() bool`

HasOpsCustomDate7 returns a boolean if a field has been set.

### GetOwnerUserId

`func (o *OpsAuditItemsPutPreviousValues) GetOwnerUserId() int32`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOwnerUserIdOk() (*int32, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *OpsAuditItemsPutPreviousValues) SetOwnerUserId(v int32)`

SetOwnerUserId sets OwnerUserId field to given value.

### HasOwnerUserId

`func (o *OpsAuditItemsPutPreviousValues) HasOwnerUserId() bool`

HasOwnerUserId returns a boolean if a field has been set.

### GetCustomText4

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText4() string`

GetCustomText4 returns the CustomText4 field if non-nil, zero value otherwise.

### GetCustomText4Ok

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText4Ok() (*string, bool)`

GetCustomText4Ok returns a tuple with the CustomText4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText4

`func (o *OpsAuditItemsPutPreviousValues) SetCustomText4(v string)`

SetCustomText4 sets CustomText4 field to given value.

### HasCustomText4

`func (o *OpsAuditItemsPutPreviousValues) HasCustomText4() bool`

HasCustomText4 returns a boolean if a field has been set.

### GetCustomText5

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText5() string`

GetCustomText5 returns the CustomText5 field if non-nil, zero value otherwise.

### GetCustomText5Ok

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText5Ok() (*string, bool)`

GetCustomText5Ok returns a tuple with the CustomText5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText5

`func (o *OpsAuditItemsPutPreviousValues) SetCustomText5(v string)`

SetCustomText5 sets CustomText5 field to given value.

### HasCustomText5

`func (o *OpsAuditItemsPutPreviousValues) HasCustomText5() bool`

HasCustomText5 returns a boolean if a field has been set.

### GetCustomText6

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText6() string`

GetCustomText6 returns the CustomText6 field if non-nil, zero value otherwise.

### GetCustomText6Ok

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText6Ok() (*string, bool)`

GetCustomText6Ok returns a tuple with the CustomText6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText6

`func (o *OpsAuditItemsPutPreviousValues) SetCustomText6(v string)`

SetCustomText6 sets CustomText6 field to given value.

### HasCustomText6

`func (o *OpsAuditItemsPutPreviousValues) HasCustomText6() bool`

HasCustomText6 returns a boolean if a field has been set.

### GetOpsCustomSelect5OptionId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect5OptionId() int32`

GetOpsCustomSelect5OptionId returns the OpsCustomSelect5OptionId field if non-nil, zero value otherwise.

### GetOpsCustomSelect5OptionIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect5OptionIdOk() (*int32, bool)`

GetOpsCustomSelect5OptionIdOk returns a tuple with the OpsCustomSelect5OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomSelect5OptionId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomSelect5OptionId(v int32)`

SetOpsCustomSelect5OptionId sets OpsCustomSelect5OptionId field to given value.

### HasOpsCustomSelect5OptionId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomSelect5OptionId() bool`

HasOpsCustomSelect5OptionId returns a boolean if a field has been set.

### GetOpsCustomSelect6OptionId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect6OptionId() int32`

GetOpsCustomSelect6OptionId returns the OpsCustomSelect6OptionId field if non-nil, zero value otherwise.

### GetOpsCustomSelect6OptionIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect6OptionIdOk() (*int32, bool)`

GetOpsCustomSelect6OptionIdOk returns a tuple with the OpsCustomSelect6OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomSelect6OptionId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomSelect6OptionId(v int32)`

SetOpsCustomSelect6OptionId sets OpsCustomSelect6OptionId field to given value.

### HasOpsCustomSelect6OptionId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomSelect6OptionId() bool`

HasOpsCustomSelect6OptionId returns a boolean if a field has been set.

### GetCustomText7

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText7() string`

GetCustomText7 returns the CustomText7 field if non-nil, zero value otherwise.

### GetCustomText7Ok

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText7Ok() (*string, bool)`

GetCustomText7Ok returns a tuple with the CustomText7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText7

`func (o *OpsAuditItemsPutPreviousValues) SetCustomText7(v string)`

SetCustomText7 sets CustomText7 field to given value.

### HasCustomText7

`func (o *OpsAuditItemsPutPreviousValues) HasCustomText7() bool`

HasCustomText7 returns a boolean if a field has been set.

### GetCustomText8

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText8() string`

GetCustomText8 returns the CustomText8 field if non-nil, zero value otherwise.

### GetCustomText8Ok

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText8Ok() (*string, bool)`

GetCustomText8Ok returns a tuple with the CustomText8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText8

`func (o *OpsAuditItemsPutPreviousValues) SetCustomText8(v string)`

SetCustomText8 sets CustomText8 field to given value.

### HasCustomText8

`func (o *OpsAuditItemsPutPreviousValues) HasCustomText8() bool`

HasCustomText8 returns a boolean if a field has been set.

### GetCustomText9

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText9() string`

GetCustomText9 returns the CustomText9 field if non-nil, zero value otherwise.

### GetCustomText9Ok

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText9Ok() (*string, bool)`

GetCustomText9Ok returns a tuple with the CustomText9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText9

`func (o *OpsAuditItemsPutPreviousValues) SetCustomText9(v string)`

SetCustomText9 sets CustomText9 field to given value.

### HasCustomText9

`func (o *OpsAuditItemsPutPreviousValues) HasCustomText9() bool`

HasCustomText9 returns a boolean if a field has been set.

### GetCustomText10

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText10() string`

GetCustomText10 returns the CustomText10 field if non-nil, zero value otherwise.

### GetCustomText10Ok

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText10Ok() (*string, bool)`

GetCustomText10Ok returns a tuple with the CustomText10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText10

`func (o *OpsAuditItemsPutPreviousValues) SetCustomText10(v string)`

SetCustomText10 sets CustomText10 field to given value.

### HasCustomText10

`func (o *OpsAuditItemsPutPreviousValues) HasCustomText10() bool`

HasCustomText10 returns a boolean if a field has been set.

### GetCustomText11

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText11() string`

GetCustomText11 returns the CustomText11 field if non-nil, zero value otherwise.

### GetCustomText11Ok

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText11Ok() (*string, bool)`

GetCustomText11Ok returns a tuple with the CustomText11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText11

`func (o *OpsAuditItemsPutPreviousValues) SetCustomText11(v string)`

SetCustomText11 sets CustomText11 field to given value.

### HasCustomText11

`func (o *OpsAuditItemsPutPreviousValues) HasCustomText11() bool`

HasCustomText11 returns a boolean if a field has been set.

### GetCustomText12

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText12() string`

GetCustomText12 returns the CustomText12 field if non-nil, zero value otherwise.

### GetCustomText12Ok

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText12Ok() (*string, bool)`

GetCustomText12Ok returns a tuple with the CustomText12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText12

`func (o *OpsAuditItemsPutPreviousValues) SetCustomText12(v string)`

SetCustomText12 sets CustomText12 field to given value.

### HasCustomText12

`func (o *OpsAuditItemsPutPreviousValues) HasCustomText12() bool`

HasCustomText12 returns a boolean if a field has been set.

### GetOpsCustomSelect7OptionId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect7OptionId() int32`

GetOpsCustomSelect7OptionId returns the OpsCustomSelect7OptionId field if non-nil, zero value otherwise.

### GetOpsCustomSelect7OptionIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect7OptionIdOk() (*int32, bool)`

GetOpsCustomSelect7OptionIdOk returns a tuple with the OpsCustomSelect7OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomSelect7OptionId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomSelect7OptionId(v int32)`

SetOpsCustomSelect7OptionId sets OpsCustomSelect7OptionId field to given value.

### HasOpsCustomSelect7OptionId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomSelect7OptionId() bool`

HasOpsCustomSelect7OptionId returns a boolean if a field has been set.

### GetOpsCustomSelect8OptionId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect8OptionId() int32`

GetOpsCustomSelect8OptionId returns the OpsCustomSelect8OptionId field if non-nil, zero value otherwise.

### GetOpsCustomSelect8OptionIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect8OptionIdOk() (*int32, bool)`

GetOpsCustomSelect8OptionIdOk returns a tuple with the OpsCustomSelect8OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomSelect8OptionId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomSelect8OptionId(v int32)`

SetOpsCustomSelect8OptionId sets OpsCustomSelect8OptionId field to given value.

### HasOpsCustomSelect8OptionId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomSelect8OptionId() bool`

HasOpsCustomSelect8OptionId returns a boolean if a field has been set.

### GetOpsCustomSelect9OptionId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect9OptionId() int32`

GetOpsCustomSelect9OptionId returns the OpsCustomSelect9OptionId field if non-nil, zero value otherwise.

### GetOpsCustomSelect9OptionIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect9OptionIdOk() (*int32, bool)`

GetOpsCustomSelect9OptionIdOk returns a tuple with the OpsCustomSelect9OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomSelect9OptionId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomSelect9OptionId(v int32)`

SetOpsCustomSelect9OptionId sets OpsCustomSelect9OptionId field to given value.

### HasOpsCustomSelect9OptionId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomSelect9OptionId() bool`

HasOpsCustomSelect9OptionId returns a boolean if a field has been set.

### GetOpsCustomSelect10OptionId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect10OptionId() int32`

GetOpsCustomSelect10OptionId returns the OpsCustomSelect10OptionId field if non-nil, zero value otherwise.

### GetOpsCustomSelect10OptionIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect10OptionIdOk() (*int32, bool)`

GetOpsCustomSelect10OptionIdOk returns a tuple with the OpsCustomSelect10OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomSelect10OptionId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomSelect10OptionId(v int32)`

SetOpsCustomSelect10OptionId sets OpsCustomSelect10OptionId field to given value.

### HasOpsCustomSelect10OptionId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomSelect10OptionId() bool`

HasOpsCustomSelect10OptionId returns a boolean if a field has been set.

### GetOpsCustomSelect11OptionId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect11OptionId() int32`

GetOpsCustomSelect11OptionId returns the OpsCustomSelect11OptionId field if non-nil, zero value otherwise.

### GetOpsCustomSelect11OptionIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect11OptionIdOk() (*int32, bool)`

GetOpsCustomSelect11OptionIdOk returns a tuple with the OpsCustomSelect11OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomSelect11OptionId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomSelect11OptionId(v int32)`

SetOpsCustomSelect11OptionId sets OpsCustomSelect11OptionId field to given value.

### HasOpsCustomSelect11OptionId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomSelect11OptionId() bool`

HasOpsCustomSelect11OptionId returns a boolean if a field has been set.

### GetOpsCustomSelect12OptionId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect12OptionId() int32`

GetOpsCustomSelect12OptionId returns the OpsCustomSelect12OptionId field if non-nil, zero value otherwise.

### GetOpsCustomSelect12OptionIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect12OptionIdOk() (*int32, bool)`

GetOpsCustomSelect12OptionIdOk returns a tuple with the OpsCustomSelect12OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomSelect12OptionId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomSelect12OptionId(v int32)`

SetOpsCustomSelect12OptionId sets OpsCustomSelect12OptionId field to given value.

### HasOpsCustomSelect12OptionId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomSelect12OptionId() bool`

HasOpsCustomSelect12OptionId returns a boolean if a field has been set.

### GetOpsCustomSelect13OptionId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect13OptionId() int32`

GetOpsCustomSelect13OptionId returns the OpsCustomSelect13OptionId field if non-nil, zero value otherwise.

### GetOpsCustomSelect13OptionIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect13OptionIdOk() (*int32, bool)`

GetOpsCustomSelect13OptionIdOk returns a tuple with the OpsCustomSelect13OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomSelect13OptionId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomSelect13OptionId(v int32)`

SetOpsCustomSelect13OptionId sets OpsCustomSelect13OptionId field to given value.

### HasOpsCustomSelect13OptionId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomSelect13OptionId() bool`

HasOpsCustomSelect13OptionId returns a boolean if a field has been set.

### GetOpsCustomSelect14OptionId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect14OptionId() int32`

GetOpsCustomSelect14OptionId returns the OpsCustomSelect14OptionId field if non-nil, zero value otherwise.

### GetOpsCustomSelect14OptionIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect14OptionIdOk() (*int32, bool)`

GetOpsCustomSelect14OptionIdOk returns a tuple with the OpsCustomSelect14OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomSelect14OptionId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomSelect14OptionId(v int32)`

SetOpsCustomSelect14OptionId sets OpsCustomSelect14OptionId field to given value.

### HasOpsCustomSelect14OptionId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomSelect14OptionId() bool`

HasOpsCustomSelect14OptionId returns a boolean if a field has been set.

### GetOpsCustomSelect15OptionId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect15OptionId() int32`

GetOpsCustomSelect15OptionId returns the OpsCustomSelect15OptionId field if non-nil, zero value otherwise.

### GetOpsCustomSelect15OptionIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect15OptionIdOk() (*int32, bool)`

GetOpsCustomSelect15OptionIdOk returns a tuple with the OpsCustomSelect15OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomSelect15OptionId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomSelect15OptionId(v int32)`

SetOpsCustomSelect15OptionId sets OpsCustomSelect15OptionId field to given value.

### HasOpsCustomSelect15OptionId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomSelect15OptionId() bool`

HasOpsCustomSelect15OptionId returns a boolean if a field has been set.

### GetCustomText13

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText13() string`

GetCustomText13 returns the CustomText13 field if non-nil, zero value otherwise.

### GetCustomText13Ok

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText13Ok() (*string, bool)`

GetCustomText13Ok returns a tuple with the CustomText13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText13

`func (o *OpsAuditItemsPutPreviousValues) SetCustomText13(v string)`

SetCustomText13 sets CustomText13 field to given value.

### HasCustomText13

`func (o *OpsAuditItemsPutPreviousValues) HasCustomText13() bool`

HasCustomText13 returns a boolean if a field has been set.

### GetCustomText14

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText14() string`

GetCustomText14 returns the CustomText14 field if non-nil, zero value otherwise.

### GetCustomText14Ok

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText14Ok() (*string, bool)`

GetCustomText14Ok returns a tuple with the CustomText14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText14

`func (o *OpsAuditItemsPutPreviousValues) SetCustomText14(v string)`

SetCustomText14 sets CustomText14 field to given value.

### HasCustomText14

`func (o *OpsAuditItemsPutPreviousValues) HasCustomText14() bool`

HasCustomText14 returns a boolean if a field has been set.

### GetCustomText15

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText15() string`

GetCustomText15 returns the CustomText15 field if non-nil, zero value otherwise.

### GetCustomText15Ok

`func (o *OpsAuditItemsPutPreviousValues) GetCustomText15Ok() (*string, bool)`

GetCustomText15Ok returns a tuple with the CustomText15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText15

`func (o *OpsAuditItemsPutPreviousValues) SetCustomText15(v string)`

SetCustomText15 sets CustomText15 field to given value.

### HasCustomText15

`func (o *OpsAuditItemsPutPreviousValues) HasCustomText15() bool`

HasCustomText15 returns a boolean if a field has been set.

### GetOpsCustomSelect16OptionId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect16OptionId() int32`

GetOpsCustomSelect16OptionId returns the OpsCustomSelect16OptionId field if non-nil, zero value otherwise.

### GetOpsCustomSelect16OptionIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect16OptionIdOk() (*int32, bool)`

GetOpsCustomSelect16OptionIdOk returns a tuple with the OpsCustomSelect16OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomSelect16OptionId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomSelect16OptionId(v int32)`

SetOpsCustomSelect16OptionId sets OpsCustomSelect16OptionId field to given value.

### HasOpsCustomSelect16OptionId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomSelect16OptionId() bool`

HasOpsCustomSelect16OptionId returns a boolean if a field has been set.

### GetOpsCustomSelect17OptionId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect17OptionId() int32`

GetOpsCustomSelect17OptionId returns the OpsCustomSelect17OptionId field if non-nil, zero value otherwise.

### GetOpsCustomSelect17OptionIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsCustomSelect17OptionIdOk() (*int32, bool)`

GetOpsCustomSelect17OptionIdOk returns a tuple with the OpsCustomSelect17OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCustomSelect17OptionId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsCustomSelect17OptionId(v int32)`

SetOpsCustomSelect17OptionId sets OpsCustomSelect17OptionId field to given value.

### HasOpsCustomSelect17OptionId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsCustomSelect17OptionId() bool`

HasOpsCustomSelect17OptionId returns a boolean if a field has been set.

### GetOpsFrequencyOptionId

`func (o *OpsAuditItemsPutPreviousValues) GetOpsFrequencyOptionId() int32`

GetOpsFrequencyOptionId returns the OpsFrequencyOptionId field if non-nil, zero value otherwise.

### GetOpsFrequencyOptionIdOk

`func (o *OpsAuditItemsPutPreviousValues) GetOpsFrequencyOptionIdOk() (*int32, bool)`

GetOpsFrequencyOptionIdOk returns a tuple with the OpsFrequencyOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsFrequencyOptionId

`func (o *OpsAuditItemsPutPreviousValues) SetOpsFrequencyOptionId(v int32)`

SetOpsFrequencyOptionId sets OpsFrequencyOptionId field to given value.

### HasOpsFrequencyOptionId

`func (o *OpsAuditItemsPutPreviousValues) HasOpsFrequencyOptionId() bool`

HasOpsFrequencyOptionId returns a boolean if a field has been set.

### GetTotals

`func (o *OpsAuditItemsPutPreviousValues) GetTotals() interface{}`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *OpsAuditItemsPutPreviousValues) GetTotalsOk() (*interface{}, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *OpsAuditItemsPutPreviousValues) SetTotals(v interface{})`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *OpsAuditItemsPutPreviousValues) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### SetTotalsNil

`func (o *OpsAuditItemsPutPreviousValues) SetTotalsNil(b bool)`

 SetTotalsNil sets the value for Totals to be an explicit nil

### UnsetTotals
`func (o *OpsAuditItemsPutPreviousValues) UnsetTotals()`

UnsetTotals ensures that no value is present for Totals, not even an explicit nil
### GetFieldData

`func (o *OpsAuditItemsPutPreviousValues) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *OpsAuditItemsPutPreviousValues) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *OpsAuditItemsPutPreviousValues) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *OpsAuditItemsPutPreviousValues) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *OpsAuditItemsPutPreviousValues) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *OpsAuditItemsPutPreviousValues) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil
### GetFirstCompleteDate

`func (o *OpsAuditItemsPutPreviousValues) GetFirstCompleteDate() string`

GetFirstCompleteDate returns the FirstCompleteDate field if non-nil, zero value otherwise.

### GetFirstCompleteDateOk

`func (o *OpsAuditItemsPutPreviousValues) GetFirstCompleteDateOk() (*string, bool)`

GetFirstCompleteDateOk returns a tuple with the FirstCompleteDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCompleteDate

`func (o *OpsAuditItemsPutPreviousValues) SetFirstCompleteDate(v string)`

SetFirstCompleteDate sets FirstCompleteDate field to given value.

### HasFirstCompleteDate

`func (o *OpsAuditItemsPutPreviousValues) HasFirstCompleteDate() bool`

HasFirstCompleteDate returns a boolean if a field has been set.

### GetFirstReviewedDate

`func (o *OpsAuditItemsPutPreviousValues) GetFirstReviewedDate() string`

GetFirstReviewedDate returns the FirstReviewedDate field if non-nil, zero value otherwise.

### GetFirstReviewedDateOk

`func (o *OpsAuditItemsPutPreviousValues) GetFirstReviewedDateOk() (*string, bool)`

GetFirstReviewedDateOk returns a tuple with the FirstReviewedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstReviewedDate

`func (o *OpsAuditItemsPutPreviousValues) SetFirstReviewedDate(v string)`

SetFirstReviewedDate sets FirstReviewedDate field to given value.

### HasFirstReviewedDate

`func (o *OpsAuditItemsPutPreviousValues) HasFirstReviewedDate() bool`

HasFirstReviewedDate returns a boolean if a field has been set.

### GetFirstUnderReview1Date

`func (o *OpsAuditItemsPutPreviousValues) GetFirstUnderReview1Date() string`

GetFirstUnderReview1Date returns the FirstUnderReview1Date field if non-nil, zero value otherwise.

### GetFirstUnderReview1DateOk

`func (o *OpsAuditItemsPutPreviousValues) GetFirstUnderReview1DateOk() (*string, bool)`

GetFirstUnderReview1DateOk returns a tuple with the FirstUnderReview1Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstUnderReview1Date

`func (o *OpsAuditItemsPutPreviousValues) SetFirstUnderReview1Date(v string)`

SetFirstUnderReview1Date sets FirstUnderReview1Date field to given value.

### HasFirstUnderReview1Date

`func (o *OpsAuditItemsPutPreviousValues) HasFirstUnderReview1Date() bool`

HasFirstUnderReview1Date returns a boolean if a field has been set.

### GetFirstUnderReview2Date

`func (o *OpsAuditItemsPutPreviousValues) GetFirstUnderReview2Date() string`

GetFirstUnderReview2Date returns the FirstUnderReview2Date field if non-nil, zero value otherwise.

### GetFirstUnderReview2DateOk

`func (o *OpsAuditItemsPutPreviousValues) GetFirstUnderReview2DateOk() (*string, bool)`

GetFirstUnderReview2DateOk returns a tuple with the FirstUnderReview2Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstUnderReview2Date

`func (o *OpsAuditItemsPutPreviousValues) SetFirstUnderReview2Date(v string)`

SetFirstUnderReview2Date sets FirstUnderReview2Date field to given value.

### HasFirstUnderReview2Date

`func (o *OpsAuditItemsPutPreviousValues) HasFirstUnderReview2Date() bool`

HasFirstUnderReview2Date returns a boolean if a field has been set.

### GetFlattened

`func (o *OpsAuditItemsPutPreviousValues) GetFlattened() interface{}`

GetFlattened returns the Flattened field if non-nil, zero value otherwise.

### GetFlattenedOk

`func (o *OpsAuditItemsPutPreviousValues) GetFlattenedOk() (*interface{}, bool)`

GetFlattenedOk returns a tuple with the Flattened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlattened

`func (o *OpsAuditItemsPutPreviousValues) SetFlattened(v interface{})`

SetFlattened sets Flattened field to given value.

### HasFlattened

`func (o *OpsAuditItemsPutPreviousValues) HasFlattened() bool`

HasFlattened returns a boolean if a field has been set.

### SetFlattenedNil

`func (o *OpsAuditItemsPutPreviousValues) SetFlattenedNil(b bool)`

 SetFlattenedNil sets the value for Flattened to be an explicit nil

### UnsetFlattened
`func (o *OpsAuditItemsPutPreviousValues) UnsetFlattened()`

UnsetFlattened ensures that no value is present for Flattened, not even an explicit nil
### GetIsFlattened

`func (o *OpsAuditItemsPutPreviousValues) GetIsFlattened() bool`

GetIsFlattened returns the IsFlattened field if non-nil, zero value otherwise.

### GetIsFlattenedOk

`func (o *OpsAuditItemsPutPreviousValues) GetIsFlattenedOk() (*bool, bool)`

GetIsFlattenedOk returns a tuple with the IsFlattened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlattened

`func (o *OpsAuditItemsPutPreviousValues) SetIsFlattened(v bool)`

SetIsFlattened sets IsFlattened field to given value.

### HasIsFlattened

`func (o *OpsAuditItemsPutPreviousValues) HasIsFlattened() bool`

HasIsFlattened returns a boolean if a field has been set.

### GetBudgetHours

`func (o *OpsAuditItemsPutPreviousValues) GetBudgetHours() float32`

GetBudgetHours returns the BudgetHours field if non-nil, zero value otherwise.

### GetBudgetHoursOk

`func (o *OpsAuditItemsPutPreviousValues) GetBudgetHoursOk() (*float32, bool)`

GetBudgetHoursOk returns a tuple with the BudgetHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetHours

`func (o *OpsAuditItemsPutPreviousValues) SetBudgetHours(v float32)`

SetBudgetHours sets BudgetHours field to given value.

### HasBudgetHours

`func (o *OpsAuditItemsPutPreviousValues) HasBudgetHours() bool`

HasBudgetHours returns a boolean if a field has been set.

### GetActualHours

`func (o *OpsAuditItemsPutPreviousValues) GetActualHours() float32`

GetActualHours returns the ActualHours field if non-nil, zero value otherwise.

### GetActualHoursOk

`func (o *OpsAuditItemsPutPreviousValues) GetActualHoursOk() (*float32, bool)`

GetActualHoursOk returns a tuple with the ActualHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualHours

`func (o *OpsAuditItemsPutPreviousValues) SetActualHours(v float32)`

SetActualHours sets ActualHours field to given value.

### HasActualHours

`func (o *OpsAuditItemsPutPreviousValues) HasActualHours() bool`

HasActualHours returns a boolean if a field has been set.

### GetIdString

`func (o *OpsAuditItemsPutPreviousValues) GetIdString() string`

GetIdString returns the IdString field if non-nil, zero value otherwise.

### GetIdStringOk

`func (o *OpsAuditItemsPutPreviousValues) GetIdStringOk() (*string, bool)`

GetIdStringOk returns a tuple with the IdString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdString

`func (o *OpsAuditItemsPutPreviousValues) SetIdString(v string)`

SetIdString sets IdString field to given value.

### HasIdString

`func (o *OpsAuditItemsPutPreviousValues) HasIdString() bool`

HasIdString returns a boolean if a field has been set.

### GetScopeRationale

`func (o *OpsAuditItemsPutPreviousValues) GetScopeRationale() string`

GetScopeRationale returns the ScopeRationale field if non-nil, zero value otherwise.

### GetScopeRationaleOk

`func (o *OpsAuditItemsPutPreviousValues) GetScopeRationaleOk() (*string, bool)`

GetScopeRationaleOk returns a tuple with the ScopeRationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeRationale

`func (o *OpsAuditItemsPutPreviousValues) SetScopeRationale(v string)`

SetScopeRationale sets ScopeRationale field to given value.

### HasScopeRationale

`func (o *OpsAuditItemsPutPreviousValues) HasScopeRationale() bool`

HasScopeRationale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


