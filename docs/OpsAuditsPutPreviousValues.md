# OpsAuditsPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**OpsAuditPeriodId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_periods.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_periods&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditCategoryId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_categories.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_categories&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CreatorUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**OpsAuditRiskLevelId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_risk_levels.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_risk_levels&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsTemplate** | Pointer to **bool** |  | [optional] [default to false]
**Status** | Pointer to **string** |  | [optional] 
**ArchivedDate** | Pointer to **string** |  | [optional] 
**ArchivedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Totals** | Pointer to **interface{}** |  | [optional] 
**SectionTotals** | Pointer to **interface{}** |  | [optional] 
**LeadUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ScopeStartDate** | Pointer to **string** |  | [optional] 
**ScopeEndDate** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**StartedDate** | Pointer to **string** |  | [optional] 
**StartedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CanceledDate** | Pointer to **string** |  | [optional] 
**CanceledByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskFactors** | Pointer to **string** |  | [optional] 
**OpsAuditCustomSelect1OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_custom_select1_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_custom_select1_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditCustomSelect2OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_custom_select2_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_custom_select2_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditCustomSelect3OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_custom_select3_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_custom_select3_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditOpinionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_opinions.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_opinions&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**SystemsReviewed** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**CustomText1** | Pointer to **string** |  | [optional] 
**CustomText2** | Pointer to **string** |  | [optional] 
**ResultCustomText1** | Pointer to **string** |  | [optional] 
**ResultCustomText2** | Pointer to **string** |  | [optional] 
**ExecutiveSummary** | Pointer to **string** |  | [optional] 
**ResultFeedback** | Pointer to **string** |  | [optional] 
**OpsAuditResultCustomSelect1OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_result_custom_select1_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_result_custom_select1_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditResultCustomSelect2OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_result_custom_select2_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_result_custom_select2_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditResultCustomSelect3OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_result_custom_select3_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_result_custom_select3_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditResultCustomSelect4OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_result_custom_select4_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_result_custom_select4_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ProjectDirectorUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ProjectLeaderUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FinalReportDate** | Pointer to **string** |  | [optional] 
**EstimatedStartDate** | Pointer to **string** |  | [optional] 
**EstimatedEndDate** | Pointer to **string** |  | [optional] 
**PreparerUsers** | Pointer to **interface{}** |  | [optional] 
**ReviewerUsers** | Pointer to **interface{}** |  | [optional] 
**SecondaryReviewerUsers** | Pointer to **interface{}** |  | [optional] 
**ProjectOptions** | Pointer to **interface{}** |  | [optional] 
**OldBudgetHours** | Pointer to **float32** |  | [optional] 
**CustomText3** | Pointer to **string** |  | [optional] 
**OldActualHours** | Pointer to **float32** |  | [optional] 
**OldInternalHours** | Pointer to **float32** |  | [optional] 
**OldExternalHours** | Pointer to **float32** |  | [optional] 
**ClosingMeetingDate** | Pointer to **string** |  | [optional] 
**OpsAuditCustomDate1** | Pointer to **string** |  | [optional] 
**OpsAuditCustomDate2** | Pointer to **string** |  | [optional] 
**OpsAuditCustomDate3** | Pointer to **string** |  | [optional] 
**OpsAuditCustomDate4** | Pointer to **string** |  | [optional] 
**OpsAuditCustomDate5** | Pointer to **string** |  | [optional] 
**OpsAuditCustomDate6** | Pointer to **string** |  | [optional] 
**CustomText4** | Pointer to **string** |  | [optional] 
**CustomText5** | Pointer to **string** |  | [optional] 
**ReactivatedDate** | Pointer to **string** |  | [optional] 
**ReactivatedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditCustomSelect4OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_custom_select4_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_custom_select4_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditCustomSelect5OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_custom_select5_options.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_custom_select5_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditCustomDate7** | Pointer to **string** |  | [optional] 
**OpsAuditCustomDate8** | Pointer to **string** |  | [optional] 
**AdditionalOwnerUsers** | Pointer to **interface{}** |  | [optional] 
**BusinessOwnerUsers** | Pointer to **interface{}** |  | [optional] 
**Scopes** | Pointer to **interface{}** |  | [optional] 
**CustomText6** | Pointer to **string** |  | [optional] 
**CustomText7** | Pointer to **string** |  | [optional] 
**CustomText8** | Pointer to **string** |  | [optional] 
**CustomText9** | Pointer to **string** |  | [optional] 
**CustomText10** | Pointer to **string** |  | [optional] 
**CustomText11** | Pointer to **string** |  | [optional] 
**FieldData** | Pointer to **interface{}** |  | [optional] 
**OpsAuditAuditReportStatusId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_audit_report_statuses.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_audit_report_statuses&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditTeamResponsibleId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_team_responsibles.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_team_responsibles&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditAcpId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_acps.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_acps&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditOtherRatingId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_other_ratings.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_other_ratings&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditDivisionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_divisions.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_divisions&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ProjectManagerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ProjectCoordinatorUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Flattened** | Pointer to **interface{}** |  | [optional] 
**IsFlattened** | Pointer to **bool** |  | [optional] [default to false]
**Purpose** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**DraftReportDate** | Pointer to **string** |  | [optional] 
**OpeningMeetingDate** | Pointer to **string** |  | [optional] 
**FrameworkId** | Pointer to **int32** |  | [optional] 
**PlanningEndDate** | Pointer to **string** |  | [optional] 
**FieldworkEndDate** | Pointer to **string** |  | [optional] 
**BudgetHours** | Pointer to **float32** |  | [optional] 
**ActualHours** | Pointer to **float32** |  | [optional] 
**ExternalHours** | Pointer to **float32** |  | [optional] 
**InternalHours** | Pointer to **float32** |  | [optional] 
**TemplateId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audits.id&#x60;.&lt;fk table&#x3D;&#39;ops_audits&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IsProcessing** | Pointer to **bool** |  | [optional] [default to false]
**AssessmentTemplateId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;assessment_templates.id&#x60;.&lt;fk table&#x3D;&#39;assessment_templates&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AllowedAssessableType** | Pointer to **string** |  | [optional] 
**IsRiskAssessmentEnabled** | Pointer to **bool** |  | [optional] [default to false]
**AssessmentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;assessments.id&#x60;.&lt;fk table&#x3D;&#39;assessments&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ProjectObjectives** | Pointer to **string** |  | [optional] 
**ProjectDrivers** | Pointer to **string** |  | [optional] 
**KeyMilestones** | Pointer to **string** |  | [optional] 
**PriorProjectKnowledge** | Pointer to **string** |  | [optional] 
**HighLevelScope** | Pointer to **string** |  | [optional] 
**ScopingConsiderationsCompletionDate** | Pointer to **string** |  | [optional] 
**ScopingConsiderationsLastUpdated** | Pointer to **string** |  | [optional] 
**OpsAuditCustomDecimal1** | Pointer to **float32** |  | [optional] 
**OpsAuditCustomDecimal2** | Pointer to **float32** |  | [optional] 
**OpsAuditCustomDecimal3** | Pointer to **float32** |  | [optional] 
**OpsAuditCustomDecimal4** | Pointer to **float32** |  | [optional] 
**OpsAuditCustomDecimal5** | Pointer to **float32** |  | [optional] 

## Methods

### NewOpsAuditsPutPreviousValues

`func NewOpsAuditsPutPreviousValues() *OpsAuditsPutPreviousValues`

NewOpsAuditsPutPreviousValues instantiates a new OpsAuditsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsAuditsPutPreviousValuesWithDefaults

`func NewOpsAuditsPutPreviousValuesWithDefaults() *OpsAuditsPutPreviousValues`

NewOpsAuditsPutPreviousValuesWithDefaults instantiates a new OpsAuditsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpsAuditsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpsAuditsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpsAuditsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OpsAuditsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OpsAuditsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OpsAuditsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OpsAuditsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OpsAuditsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OpsAuditsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OpsAuditsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OpsAuditsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OpsAuditsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *OpsAuditsPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *OpsAuditsPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *OpsAuditsPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *OpsAuditsPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetStartDate

`func (o *OpsAuditsPutPreviousValues) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *OpsAuditsPutPreviousValues) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *OpsAuditsPutPreviousValues) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *OpsAuditsPutPreviousValues) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *OpsAuditsPutPreviousValues) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *OpsAuditsPutPreviousValues) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *OpsAuditsPutPreviousValues) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *OpsAuditsPutPreviousValues) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetOpsAuditPeriodId

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditPeriodId() int32`

GetOpsAuditPeriodId returns the OpsAuditPeriodId field if non-nil, zero value otherwise.

### GetOpsAuditPeriodIdOk

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditPeriodIdOk() (*int32, bool)`

GetOpsAuditPeriodIdOk returns a tuple with the OpsAuditPeriodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditPeriodId

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditPeriodId(v int32)`

SetOpsAuditPeriodId sets OpsAuditPeriodId field to given value.

### HasOpsAuditPeriodId

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditPeriodId() bool`

HasOpsAuditPeriodId returns a boolean if a field has been set.

### GetOpsAuditCategoryId

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCategoryId() int32`

GetOpsAuditCategoryId returns the OpsAuditCategoryId field if non-nil, zero value otherwise.

### GetOpsAuditCategoryIdOk

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCategoryIdOk() (*int32, bool)`

GetOpsAuditCategoryIdOk returns a tuple with the OpsAuditCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditCategoryId

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditCategoryId(v int32)`

SetOpsAuditCategoryId sets OpsAuditCategoryId field to given value.

### HasOpsAuditCategoryId

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditCategoryId() bool`

HasOpsAuditCategoryId returns a boolean if a field has been set.

### GetCreatorUserId

`func (o *OpsAuditsPutPreviousValues) GetCreatorUserId() int32`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *OpsAuditsPutPreviousValues) GetCreatorUserIdOk() (*int32, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *OpsAuditsPutPreviousValues) SetCreatorUserId(v int32)`

SetCreatorUserId sets CreatorUserId field to given value.

### HasCreatorUserId

`func (o *OpsAuditsPutPreviousValues) HasCreatorUserId() bool`

HasCreatorUserId returns a boolean if a field has been set.

### GetTitle

`func (o *OpsAuditsPutPreviousValues) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OpsAuditsPutPreviousValues) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OpsAuditsPutPreviousValues) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *OpsAuditsPutPreviousValues) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetOpsAuditRiskLevelId

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditRiskLevelId() int32`

GetOpsAuditRiskLevelId returns the OpsAuditRiskLevelId field if non-nil, zero value otherwise.

### GetOpsAuditRiskLevelIdOk

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditRiskLevelIdOk() (*int32, bool)`

GetOpsAuditRiskLevelIdOk returns a tuple with the OpsAuditRiskLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditRiskLevelId

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditRiskLevelId(v int32)`

SetOpsAuditRiskLevelId sets OpsAuditRiskLevelId field to given value.

### HasOpsAuditRiskLevelId

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditRiskLevelId() bool`

HasOpsAuditRiskLevelId returns a boolean if a field has been set.

### GetDescription

`func (o *OpsAuditsPutPreviousValues) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpsAuditsPutPreviousValues) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpsAuditsPutPreviousValues) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OpsAuditsPutPreviousValues) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsTemplate

`func (o *OpsAuditsPutPreviousValues) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *OpsAuditsPutPreviousValues) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *OpsAuditsPutPreviousValues) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *OpsAuditsPutPreviousValues) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.

### GetStatus

`func (o *OpsAuditsPutPreviousValues) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpsAuditsPutPreviousValues) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpsAuditsPutPreviousValues) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OpsAuditsPutPreviousValues) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetArchivedDate

`func (o *OpsAuditsPutPreviousValues) GetArchivedDate() string`

GetArchivedDate returns the ArchivedDate field if non-nil, zero value otherwise.

### GetArchivedDateOk

`func (o *OpsAuditsPutPreviousValues) GetArchivedDateOk() (*string, bool)`

GetArchivedDateOk returns a tuple with the ArchivedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedDate

`func (o *OpsAuditsPutPreviousValues) SetArchivedDate(v string)`

SetArchivedDate sets ArchivedDate field to given value.

### HasArchivedDate

`func (o *OpsAuditsPutPreviousValues) HasArchivedDate() bool`

HasArchivedDate returns a boolean if a field has been set.

### GetArchivedByUserId

`func (o *OpsAuditsPutPreviousValues) GetArchivedByUserId() int32`

GetArchivedByUserId returns the ArchivedByUserId field if non-nil, zero value otherwise.

### GetArchivedByUserIdOk

`func (o *OpsAuditsPutPreviousValues) GetArchivedByUserIdOk() (*int32, bool)`

GetArchivedByUserIdOk returns a tuple with the ArchivedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedByUserId

`func (o *OpsAuditsPutPreviousValues) SetArchivedByUserId(v int32)`

SetArchivedByUserId sets ArchivedByUserId field to given value.

### HasArchivedByUserId

`func (o *OpsAuditsPutPreviousValues) HasArchivedByUserId() bool`

HasArchivedByUserId returns a boolean if a field has been set.

### GetTotals

`func (o *OpsAuditsPutPreviousValues) GetTotals() interface{}`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *OpsAuditsPutPreviousValues) GetTotalsOk() (*interface{}, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *OpsAuditsPutPreviousValues) SetTotals(v interface{})`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *OpsAuditsPutPreviousValues) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### SetTotalsNil

`func (o *OpsAuditsPutPreviousValues) SetTotalsNil(b bool)`

 SetTotalsNil sets the value for Totals to be an explicit nil

### UnsetTotals
`func (o *OpsAuditsPutPreviousValues) UnsetTotals()`

UnsetTotals ensures that no value is present for Totals, not even an explicit nil
### GetSectionTotals

`func (o *OpsAuditsPutPreviousValues) GetSectionTotals() interface{}`

GetSectionTotals returns the SectionTotals field if non-nil, zero value otherwise.

### GetSectionTotalsOk

`func (o *OpsAuditsPutPreviousValues) GetSectionTotalsOk() (*interface{}, bool)`

GetSectionTotalsOk returns a tuple with the SectionTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionTotals

`func (o *OpsAuditsPutPreviousValues) SetSectionTotals(v interface{})`

SetSectionTotals sets SectionTotals field to given value.

### HasSectionTotals

`func (o *OpsAuditsPutPreviousValues) HasSectionTotals() bool`

HasSectionTotals returns a boolean if a field has been set.

### SetSectionTotalsNil

`func (o *OpsAuditsPutPreviousValues) SetSectionTotalsNil(b bool)`

 SetSectionTotalsNil sets the value for SectionTotals to be an explicit nil

### UnsetSectionTotals
`func (o *OpsAuditsPutPreviousValues) UnsetSectionTotals()`

UnsetSectionTotals ensures that no value is present for SectionTotals, not even an explicit nil
### GetLeadUserId

`func (o *OpsAuditsPutPreviousValues) GetLeadUserId() int32`

GetLeadUserId returns the LeadUserId field if non-nil, zero value otherwise.

### GetLeadUserIdOk

`func (o *OpsAuditsPutPreviousValues) GetLeadUserIdOk() (*int32, bool)`

GetLeadUserIdOk returns a tuple with the LeadUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadUserId

`func (o *OpsAuditsPutPreviousValues) SetLeadUserId(v int32)`

SetLeadUserId sets LeadUserId field to given value.

### HasLeadUserId

`func (o *OpsAuditsPutPreviousValues) HasLeadUserId() bool`

HasLeadUserId returns a boolean if a field has been set.

### GetScopeStartDate

`func (o *OpsAuditsPutPreviousValues) GetScopeStartDate() string`

GetScopeStartDate returns the ScopeStartDate field if non-nil, zero value otherwise.

### GetScopeStartDateOk

`func (o *OpsAuditsPutPreviousValues) GetScopeStartDateOk() (*string, bool)`

GetScopeStartDateOk returns a tuple with the ScopeStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeStartDate

`func (o *OpsAuditsPutPreviousValues) SetScopeStartDate(v string)`

SetScopeStartDate sets ScopeStartDate field to given value.

### HasScopeStartDate

`func (o *OpsAuditsPutPreviousValues) HasScopeStartDate() bool`

HasScopeStartDate returns a boolean if a field has been set.

### GetScopeEndDate

`func (o *OpsAuditsPutPreviousValues) GetScopeEndDate() string`

GetScopeEndDate returns the ScopeEndDate field if non-nil, zero value otherwise.

### GetScopeEndDateOk

`func (o *OpsAuditsPutPreviousValues) GetScopeEndDateOk() (*string, bool)`

GetScopeEndDateOk returns a tuple with the ScopeEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeEndDate

`func (o *OpsAuditsPutPreviousValues) SetScopeEndDate(v string)`

SetScopeEndDate sets ScopeEndDate field to given value.

### HasScopeEndDate

`func (o *OpsAuditsPutPreviousValues) HasScopeEndDate() bool`

HasScopeEndDate returns a boolean if a field has been set.

### GetUid

`func (o *OpsAuditsPutPreviousValues) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *OpsAuditsPutPreviousValues) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *OpsAuditsPutPreviousValues) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *OpsAuditsPutPreviousValues) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetStartedDate

`func (o *OpsAuditsPutPreviousValues) GetStartedDate() string`

GetStartedDate returns the StartedDate field if non-nil, zero value otherwise.

### GetStartedDateOk

`func (o *OpsAuditsPutPreviousValues) GetStartedDateOk() (*string, bool)`

GetStartedDateOk returns a tuple with the StartedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedDate

`func (o *OpsAuditsPutPreviousValues) SetStartedDate(v string)`

SetStartedDate sets StartedDate field to given value.

### HasStartedDate

`func (o *OpsAuditsPutPreviousValues) HasStartedDate() bool`

HasStartedDate returns a boolean if a field has been set.

### GetStartedByUserId

`func (o *OpsAuditsPutPreviousValues) GetStartedByUserId() int32`

GetStartedByUserId returns the StartedByUserId field if non-nil, zero value otherwise.

### GetStartedByUserIdOk

`func (o *OpsAuditsPutPreviousValues) GetStartedByUserIdOk() (*int32, bool)`

GetStartedByUserIdOk returns a tuple with the StartedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedByUserId

`func (o *OpsAuditsPutPreviousValues) SetStartedByUserId(v int32)`

SetStartedByUserId sets StartedByUserId field to given value.

### HasStartedByUserId

`func (o *OpsAuditsPutPreviousValues) HasStartedByUserId() bool`

HasStartedByUserId returns a boolean if a field has been set.

### GetCanceledDate

`func (o *OpsAuditsPutPreviousValues) GetCanceledDate() string`

GetCanceledDate returns the CanceledDate field if non-nil, zero value otherwise.

### GetCanceledDateOk

`func (o *OpsAuditsPutPreviousValues) GetCanceledDateOk() (*string, bool)`

GetCanceledDateOk returns a tuple with the CanceledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledDate

`func (o *OpsAuditsPutPreviousValues) SetCanceledDate(v string)`

SetCanceledDate sets CanceledDate field to given value.

### HasCanceledDate

`func (o *OpsAuditsPutPreviousValues) HasCanceledDate() bool`

HasCanceledDate returns a boolean if a field has been set.

### GetCanceledByUserId

`func (o *OpsAuditsPutPreviousValues) GetCanceledByUserId() int32`

GetCanceledByUserId returns the CanceledByUserId field if non-nil, zero value otherwise.

### GetCanceledByUserIdOk

`func (o *OpsAuditsPutPreviousValues) GetCanceledByUserIdOk() (*int32, bool)`

GetCanceledByUserIdOk returns a tuple with the CanceledByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledByUserId

`func (o *OpsAuditsPutPreviousValues) SetCanceledByUserId(v int32)`

SetCanceledByUserId sets CanceledByUserId field to given value.

### HasCanceledByUserId

`func (o *OpsAuditsPutPreviousValues) HasCanceledByUserId() bool`

HasCanceledByUserId returns a boolean if a field has been set.

### GetRiskFactors

`func (o *OpsAuditsPutPreviousValues) GetRiskFactors() string`

GetRiskFactors returns the RiskFactors field if non-nil, zero value otherwise.

### GetRiskFactorsOk

`func (o *OpsAuditsPutPreviousValues) GetRiskFactorsOk() (*string, bool)`

GetRiskFactorsOk returns a tuple with the RiskFactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskFactors

`func (o *OpsAuditsPutPreviousValues) SetRiskFactors(v string)`

SetRiskFactors sets RiskFactors field to given value.

### HasRiskFactors

`func (o *OpsAuditsPutPreviousValues) HasRiskFactors() bool`

HasRiskFactors returns a boolean if a field has been set.

### GetOpsAuditCustomSelect1OptionId

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomSelect1OptionId() int32`

GetOpsAuditCustomSelect1OptionId returns the OpsAuditCustomSelect1OptionId field if non-nil, zero value otherwise.

### GetOpsAuditCustomSelect1OptionIdOk

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomSelect1OptionIdOk() (*int32, bool)`

GetOpsAuditCustomSelect1OptionIdOk returns a tuple with the OpsAuditCustomSelect1OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditCustomSelect1OptionId

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditCustomSelect1OptionId(v int32)`

SetOpsAuditCustomSelect1OptionId sets OpsAuditCustomSelect1OptionId field to given value.

### HasOpsAuditCustomSelect1OptionId

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditCustomSelect1OptionId() bool`

HasOpsAuditCustomSelect1OptionId returns a boolean if a field has been set.

### GetOpsAuditCustomSelect2OptionId

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomSelect2OptionId() int32`

GetOpsAuditCustomSelect2OptionId returns the OpsAuditCustomSelect2OptionId field if non-nil, zero value otherwise.

### GetOpsAuditCustomSelect2OptionIdOk

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomSelect2OptionIdOk() (*int32, bool)`

GetOpsAuditCustomSelect2OptionIdOk returns a tuple with the OpsAuditCustomSelect2OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditCustomSelect2OptionId

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditCustomSelect2OptionId(v int32)`

SetOpsAuditCustomSelect2OptionId sets OpsAuditCustomSelect2OptionId field to given value.

### HasOpsAuditCustomSelect2OptionId

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditCustomSelect2OptionId() bool`

HasOpsAuditCustomSelect2OptionId returns a boolean if a field has been set.

### GetOpsAuditCustomSelect3OptionId

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomSelect3OptionId() int32`

GetOpsAuditCustomSelect3OptionId returns the OpsAuditCustomSelect3OptionId field if non-nil, zero value otherwise.

### GetOpsAuditCustomSelect3OptionIdOk

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomSelect3OptionIdOk() (*int32, bool)`

GetOpsAuditCustomSelect3OptionIdOk returns a tuple with the OpsAuditCustomSelect3OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditCustomSelect3OptionId

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditCustomSelect3OptionId(v int32)`

SetOpsAuditCustomSelect3OptionId sets OpsAuditCustomSelect3OptionId field to given value.

### HasOpsAuditCustomSelect3OptionId

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditCustomSelect3OptionId() bool`

HasOpsAuditCustomSelect3OptionId returns a boolean if a field has been set.

### GetOpsAuditOpinionId

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditOpinionId() int32`

GetOpsAuditOpinionId returns the OpsAuditOpinionId field if non-nil, zero value otherwise.

### GetOpsAuditOpinionIdOk

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditOpinionIdOk() (*int32, bool)`

GetOpsAuditOpinionIdOk returns a tuple with the OpsAuditOpinionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditOpinionId

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditOpinionId(v int32)`

SetOpsAuditOpinionId sets OpsAuditOpinionId field to given value.

### HasOpsAuditOpinionId

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditOpinionId() bool`

HasOpsAuditOpinionId returns a boolean if a field has been set.

### GetSystemsReviewed

`func (o *OpsAuditsPutPreviousValues) GetSystemsReviewed() string`

GetSystemsReviewed returns the SystemsReviewed field if non-nil, zero value otherwise.

### GetSystemsReviewedOk

`func (o *OpsAuditsPutPreviousValues) GetSystemsReviewedOk() (*string, bool)`

GetSystemsReviewedOk returns a tuple with the SystemsReviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemsReviewed

`func (o *OpsAuditsPutPreviousValues) SetSystemsReviewed(v string)`

SetSystemsReviewed sets SystemsReviewed field to given value.

### HasSystemsReviewed

`func (o *OpsAuditsPutPreviousValues) HasSystemsReviewed() bool`

HasSystemsReviewed returns a boolean if a field has been set.

### GetNotes

`func (o *OpsAuditsPutPreviousValues) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OpsAuditsPutPreviousValues) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OpsAuditsPutPreviousValues) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OpsAuditsPutPreviousValues) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetCustomText1

`func (o *OpsAuditsPutPreviousValues) GetCustomText1() string`

GetCustomText1 returns the CustomText1 field if non-nil, zero value otherwise.

### GetCustomText1Ok

`func (o *OpsAuditsPutPreviousValues) GetCustomText1Ok() (*string, bool)`

GetCustomText1Ok returns a tuple with the CustomText1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText1

`func (o *OpsAuditsPutPreviousValues) SetCustomText1(v string)`

SetCustomText1 sets CustomText1 field to given value.

### HasCustomText1

`func (o *OpsAuditsPutPreviousValues) HasCustomText1() bool`

HasCustomText1 returns a boolean if a field has been set.

### GetCustomText2

`func (o *OpsAuditsPutPreviousValues) GetCustomText2() string`

GetCustomText2 returns the CustomText2 field if non-nil, zero value otherwise.

### GetCustomText2Ok

`func (o *OpsAuditsPutPreviousValues) GetCustomText2Ok() (*string, bool)`

GetCustomText2Ok returns a tuple with the CustomText2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText2

`func (o *OpsAuditsPutPreviousValues) SetCustomText2(v string)`

SetCustomText2 sets CustomText2 field to given value.

### HasCustomText2

`func (o *OpsAuditsPutPreviousValues) HasCustomText2() bool`

HasCustomText2 returns a boolean if a field has been set.

### GetResultCustomText1

`func (o *OpsAuditsPutPreviousValues) GetResultCustomText1() string`

GetResultCustomText1 returns the ResultCustomText1 field if non-nil, zero value otherwise.

### GetResultCustomText1Ok

`func (o *OpsAuditsPutPreviousValues) GetResultCustomText1Ok() (*string, bool)`

GetResultCustomText1Ok returns a tuple with the ResultCustomText1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCustomText1

`func (o *OpsAuditsPutPreviousValues) SetResultCustomText1(v string)`

SetResultCustomText1 sets ResultCustomText1 field to given value.

### HasResultCustomText1

`func (o *OpsAuditsPutPreviousValues) HasResultCustomText1() bool`

HasResultCustomText1 returns a boolean if a field has been set.

### GetResultCustomText2

`func (o *OpsAuditsPutPreviousValues) GetResultCustomText2() string`

GetResultCustomText2 returns the ResultCustomText2 field if non-nil, zero value otherwise.

### GetResultCustomText2Ok

`func (o *OpsAuditsPutPreviousValues) GetResultCustomText2Ok() (*string, bool)`

GetResultCustomText2Ok returns a tuple with the ResultCustomText2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCustomText2

`func (o *OpsAuditsPutPreviousValues) SetResultCustomText2(v string)`

SetResultCustomText2 sets ResultCustomText2 field to given value.

### HasResultCustomText2

`func (o *OpsAuditsPutPreviousValues) HasResultCustomText2() bool`

HasResultCustomText2 returns a boolean if a field has been set.

### GetExecutiveSummary

`func (o *OpsAuditsPutPreviousValues) GetExecutiveSummary() string`

GetExecutiveSummary returns the ExecutiveSummary field if non-nil, zero value otherwise.

### GetExecutiveSummaryOk

`func (o *OpsAuditsPutPreviousValues) GetExecutiveSummaryOk() (*string, bool)`

GetExecutiveSummaryOk returns a tuple with the ExecutiveSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutiveSummary

`func (o *OpsAuditsPutPreviousValues) SetExecutiveSummary(v string)`

SetExecutiveSummary sets ExecutiveSummary field to given value.

### HasExecutiveSummary

`func (o *OpsAuditsPutPreviousValues) HasExecutiveSummary() bool`

HasExecutiveSummary returns a boolean if a field has been set.

### GetResultFeedback

`func (o *OpsAuditsPutPreviousValues) GetResultFeedback() string`

GetResultFeedback returns the ResultFeedback field if non-nil, zero value otherwise.

### GetResultFeedbackOk

`func (o *OpsAuditsPutPreviousValues) GetResultFeedbackOk() (*string, bool)`

GetResultFeedbackOk returns a tuple with the ResultFeedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultFeedback

`func (o *OpsAuditsPutPreviousValues) SetResultFeedback(v string)`

SetResultFeedback sets ResultFeedback field to given value.

### HasResultFeedback

`func (o *OpsAuditsPutPreviousValues) HasResultFeedback() bool`

HasResultFeedback returns a boolean if a field has been set.

### GetOpsAuditResultCustomSelect1OptionId

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditResultCustomSelect1OptionId() int32`

GetOpsAuditResultCustomSelect1OptionId returns the OpsAuditResultCustomSelect1OptionId field if non-nil, zero value otherwise.

### GetOpsAuditResultCustomSelect1OptionIdOk

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditResultCustomSelect1OptionIdOk() (*int32, bool)`

GetOpsAuditResultCustomSelect1OptionIdOk returns a tuple with the OpsAuditResultCustomSelect1OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditResultCustomSelect1OptionId

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditResultCustomSelect1OptionId(v int32)`

SetOpsAuditResultCustomSelect1OptionId sets OpsAuditResultCustomSelect1OptionId field to given value.

### HasOpsAuditResultCustomSelect1OptionId

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditResultCustomSelect1OptionId() bool`

HasOpsAuditResultCustomSelect1OptionId returns a boolean if a field has been set.

### GetOpsAuditResultCustomSelect2OptionId

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditResultCustomSelect2OptionId() int32`

GetOpsAuditResultCustomSelect2OptionId returns the OpsAuditResultCustomSelect2OptionId field if non-nil, zero value otherwise.

### GetOpsAuditResultCustomSelect2OptionIdOk

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditResultCustomSelect2OptionIdOk() (*int32, bool)`

GetOpsAuditResultCustomSelect2OptionIdOk returns a tuple with the OpsAuditResultCustomSelect2OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditResultCustomSelect2OptionId

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditResultCustomSelect2OptionId(v int32)`

SetOpsAuditResultCustomSelect2OptionId sets OpsAuditResultCustomSelect2OptionId field to given value.

### HasOpsAuditResultCustomSelect2OptionId

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditResultCustomSelect2OptionId() bool`

HasOpsAuditResultCustomSelect2OptionId returns a boolean if a field has been set.

### GetOpsAuditResultCustomSelect3OptionId

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditResultCustomSelect3OptionId() int32`

GetOpsAuditResultCustomSelect3OptionId returns the OpsAuditResultCustomSelect3OptionId field if non-nil, zero value otherwise.

### GetOpsAuditResultCustomSelect3OptionIdOk

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditResultCustomSelect3OptionIdOk() (*int32, bool)`

GetOpsAuditResultCustomSelect3OptionIdOk returns a tuple with the OpsAuditResultCustomSelect3OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditResultCustomSelect3OptionId

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditResultCustomSelect3OptionId(v int32)`

SetOpsAuditResultCustomSelect3OptionId sets OpsAuditResultCustomSelect3OptionId field to given value.

### HasOpsAuditResultCustomSelect3OptionId

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditResultCustomSelect3OptionId() bool`

HasOpsAuditResultCustomSelect3OptionId returns a boolean if a field has been set.

### GetOpsAuditResultCustomSelect4OptionId

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditResultCustomSelect4OptionId() int32`

GetOpsAuditResultCustomSelect4OptionId returns the OpsAuditResultCustomSelect4OptionId field if non-nil, zero value otherwise.

### GetOpsAuditResultCustomSelect4OptionIdOk

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditResultCustomSelect4OptionIdOk() (*int32, bool)`

GetOpsAuditResultCustomSelect4OptionIdOk returns a tuple with the OpsAuditResultCustomSelect4OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditResultCustomSelect4OptionId

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditResultCustomSelect4OptionId(v int32)`

SetOpsAuditResultCustomSelect4OptionId sets OpsAuditResultCustomSelect4OptionId field to given value.

### HasOpsAuditResultCustomSelect4OptionId

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditResultCustomSelect4OptionId() bool`

HasOpsAuditResultCustomSelect4OptionId returns a boolean if a field has been set.

### GetProjectDirectorUserId

`func (o *OpsAuditsPutPreviousValues) GetProjectDirectorUserId() int32`

GetProjectDirectorUserId returns the ProjectDirectorUserId field if non-nil, zero value otherwise.

### GetProjectDirectorUserIdOk

`func (o *OpsAuditsPutPreviousValues) GetProjectDirectorUserIdOk() (*int32, bool)`

GetProjectDirectorUserIdOk returns a tuple with the ProjectDirectorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDirectorUserId

`func (o *OpsAuditsPutPreviousValues) SetProjectDirectorUserId(v int32)`

SetProjectDirectorUserId sets ProjectDirectorUserId field to given value.

### HasProjectDirectorUserId

`func (o *OpsAuditsPutPreviousValues) HasProjectDirectorUserId() bool`

HasProjectDirectorUserId returns a boolean if a field has been set.

### GetProjectLeaderUserId

`func (o *OpsAuditsPutPreviousValues) GetProjectLeaderUserId() int32`

GetProjectLeaderUserId returns the ProjectLeaderUserId field if non-nil, zero value otherwise.

### GetProjectLeaderUserIdOk

`func (o *OpsAuditsPutPreviousValues) GetProjectLeaderUserIdOk() (*int32, bool)`

GetProjectLeaderUserIdOk returns a tuple with the ProjectLeaderUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLeaderUserId

`func (o *OpsAuditsPutPreviousValues) SetProjectLeaderUserId(v int32)`

SetProjectLeaderUserId sets ProjectLeaderUserId field to given value.

### HasProjectLeaderUserId

`func (o *OpsAuditsPutPreviousValues) HasProjectLeaderUserId() bool`

HasProjectLeaderUserId returns a boolean if a field has been set.

### GetFinalReportDate

`func (o *OpsAuditsPutPreviousValues) GetFinalReportDate() string`

GetFinalReportDate returns the FinalReportDate field if non-nil, zero value otherwise.

### GetFinalReportDateOk

`func (o *OpsAuditsPutPreviousValues) GetFinalReportDateOk() (*string, bool)`

GetFinalReportDateOk returns a tuple with the FinalReportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalReportDate

`func (o *OpsAuditsPutPreviousValues) SetFinalReportDate(v string)`

SetFinalReportDate sets FinalReportDate field to given value.

### HasFinalReportDate

`func (o *OpsAuditsPutPreviousValues) HasFinalReportDate() bool`

HasFinalReportDate returns a boolean if a field has been set.

### GetEstimatedStartDate

`func (o *OpsAuditsPutPreviousValues) GetEstimatedStartDate() string`

GetEstimatedStartDate returns the EstimatedStartDate field if non-nil, zero value otherwise.

### GetEstimatedStartDateOk

`func (o *OpsAuditsPutPreviousValues) GetEstimatedStartDateOk() (*string, bool)`

GetEstimatedStartDateOk returns a tuple with the EstimatedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedStartDate

`func (o *OpsAuditsPutPreviousValues) SetEstimatedStartDate(v string)`

SetEstimatedStartDate sets EstimatedStartDate field to given value.

### HasEstimatedStartDate

`func (o *OpsAuditsPutPreviousValues) HasEstimatedStartDate() bool`

HasEstimatedStartDate returns a boolean if a field has been set.

### GetEstimatedEndDate

`func (o *OpsAuditsPutPreviousValues) GetEstimatedEndDate() string`

GetEstimatedEndDate returns the EstimatedEndDate field if non-nil, zero value otherwise.

### GetEstimatedEndDateOk

`func (o *OpsAuditsPutPreviousValues) GetEstimatedEndDateOk() (*string, bool)`

GetEstimatedEndDateOk returns a tuple with the EstimatedEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedEndDate

`func (o *OpsAuditsPutPreviousValues) SetEstimatedEndDate(v string)`

SetEstimatedEndDate sets EstimatedEndDate field to given value.

### HasEstimatedEndDate

`func (o *OpsAuditsPutPreviousValues) HasEstimatedEndDate() bool`

HasEstimatedEndDate returns a boolean if a field has been set.

### GetPreparerUsers

`func (o *OpsAuditsPutPreviousValues) GetPreparerUsers() interface{}`

GetPreparerUsers returns the PreparerUsers field if non-nil, zero value otherwise.

### GetPreparerUsersOk

`func (o *OpsAuditsPutPreviousValues) GetPreparerUsersOk() (*interface{}, bool)`

GetPreparerUsersOk returns a tuple with the PreparerUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparerUsers

`func (o *OpsAuditsPutPreviousValues) SetPreparerUsers(v interface{})`

SetPreparerUsers sets PreparerUsers field to given value.

### HasPreparerUsers

`func (o *OpsAuditsPutPreviousValues) HasPreparerUsers() bool`

HasPreparerUsers returns a boolean if a field has been set.

### SetPreparerUsersNil

`func (o *OpsAuditsPutPreviousValues) SetPreparerUsersNil(b bool)`

 SetPreparerUsersNil sets the value for PreparerUsers to be an explicit nil

### UnsetPreparerUsers
`func (o *OpsAuditsPutPreviousValues) UnsetPreparerUsers()`

UnsetPreparerUsers ensures that no value is present for PreparerUsers, not even an explicit nil
### GetReviewerUsers

`func (o *OpsAuditsPutPreviousValues) GetReviewerUsers() interface{}`

GetReviewerUsers returns the ReviewerUsers field if non-nil, zero value otherwise.

### GetReviewerUsersOk

`func (o *OpsAuditsPutPreviousValues) GetReviewerUsersOk() (*interface{}, bool)`

GetReviewerUsersOk returns a tuple with the ReviewerUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerUsers

`func (o *OpsAuditsPutPreviousValues) SetReviewerUsers(v interface{})`

SetReviewerUsers sets ReviewerUsers field to given value.

### HasReviewerUsers

`func (o *OpsAuditsPutPreviousValues) HasReviewerUsers() bool`

HasReviewerUsers returns a boolean if a field has been set.

### SetReviewerUsersNil

`func (o *OpsAuditsPutPreviousValues) SetReviewerUsersNil(b bool)`

 SetReviewerUsersNil sets the value for ReviewerUsers to be an explicit nil

### UnsetReviewerUsers
`func (o *OpsAuditsPutPreviousValues) UnsetReviewerUsers()`

UnsetReviewerUsers ensures that no value is present for ReviewerUsers, not even an explicit nil
### GetSecondaryReviewerUsers

`func (o *OpsAuditsPutPreviousValues) GetSecondaryReviewerUsers() interface{}`

GetSecondaryReviewerUsers returns the SecondaryReviewerUsers field if non-nil, zero value otherwise.

### GetSecondaryReviewerUsersOk

`func (o *OpsAuditsPutPreviousValues) GetSecondaryReviewerUsersOk() (*interface{}, bool)`

GetSecondaryReviewerUsersOk returns a tuple with the SecondaryReviewerUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryReviewerUsers

`func (o *OpsAuditsPutPreviousValues) SetSecondaryReviewerUsers(v interface{})`

SetSecondaryReviewerUsers sets SecondaryReviewerUsers field to given value.

### HasSecondaryReviewerUsers

`func (o *OpsAuditsPutPreviousValues) HasSecondaryReviewerUsers() bool`

HasSecondaryReviewerUsers returns a boolean if a field has been set.

### SetSecondaryReviewerUsersNil

`func (o *OpsAuditsPutPreviousValues) SetSecondaryReviewerUsersNil(b bool)`

 SetSecondaryReviewerUsersNil sets the value for SecondaryReviewerUsers to be an explicit nil

### UnsetSecondaryReviewerUsers
`func (o *OpsAuditsPutPreviousValues) UnsetSecondaryReviewerUsers()`

UnsetSecondaryReviewerUsers ensures that no value is present for SecondaryReviewerUsers, not even an explicit nil
### GetProjectOptions

`func (o *OpsAuditsPutPreviousValues) GetProjectOptions() interface{}`

GetProjectOptions returns the ProjectOptions field if non-nil, zero value otherwise.

### GetProjectOptionsOk

`func (o *OpsAuditsPutPreviousValues) GetProjectOptionsOk() (*interface{}, bool)`

GetProjectOptionsOk returns a tuple with the ProjectOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectOptions

`func (o *OpsAuditsPutPreviousValues) SetProjectOptions(v interface{})`

SetProjectOptions sets ProjectOptions field to given value.

### HasProjectOptions

`func (o *OpsAuditsPutPreviousValues) HasProjectOptions() bool`

HasProjectOptions returns a boolean if a field has been set.

### SetProjectOptionsNil

`func (o *OpsAuditsPutPreviousValues) SetProjectOptionsNil(b bool)`

 SetProjectOptionsNil sets the value for ProjectOptions to be an explicit nil

### UnsetProjectOptions
`func (o *OpsAuditsPutPreviousValues) UnsetProjectOptions()`

UnsetProjectOptions ensures that no value is present for ProjectOptions, not even an explicit nil
### GetOldBudgetHours

`func (o *OpsAuditsPutPreviousValues) GetOldBudgetHours() float32`

GetOldBudgetHours returns the OldBudgetHours field if non-nil, zero value otherwise.

### GetOldBudgetHoursOk

`func (o *OpsAuditsPutPreviousValues) GetOldBudgetHoursOk() (*float32, bool)`

GetOldBudgetHoursOk returns a tuple with the OldBudgetHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldBudgetHours

`func (o *OpsAuditsPutPreviousValues) SetOldBudgetHours(v float32)`

SetOldBudgetHours sets OldBudgetHours field to given value.

### HasOldBudgetHours

`func (o *OpsAuditsPutPreviousValues) HasOldBudgetHours() bool`

HasOldBudgetHours returns a boolean if a field has been set.

### GetCustomText3

`func (o *OpsAuditsPutPreviousValues) GetCustomText3() string`

GetCustomText3 returns the CustomText3 field if non-nil, zero value otherwise.

### GetCustomText3Ok

`func (o *OpsAuditsPutPreviousValues) GetCustomText3Ok() (*string, bool)`

GetCustomText3Ok returns a tuple with the CustomText3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText3

`func (o *OpsAuditsPutPreviousValues) SetCustomText3(v string)`

SetCustomText3 sets CustomText3 field to given value.

### HasCustomText3

`func (o *OpsAuditsPutPreviousValues) HasCustomText3() bool`

HasCustomText3 returns a boolean if a field has been set.

### GetOldActualHours

`func (o *OpsAuditsPutPreviousValues) GetOldActualHours() float32`

GetOldActualHours returns the OldActualHours field if non-nil, zero value otherwise.

### GetOldActualHoursOk

`func (o *OpsAuditsPutPreviousValues) GetOldActualHoursOk() (*float32, bool)`

GetOldActualHoursOk returns a tuple with the OldActualHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldActualHours

`func (o *OpsAuditsPutPreviousValues) SetOldActualHours(v float32)`

SetOldActualHours sets OldActualHours field to given value.

### HasOldActualHours

`func (o *OpsAuditsPutPreviousValues) HasOldActualHours() bool`

HasOldActualHours returns a boolean if a field has been set.

### GetOldInternalHours

`func (o *OpsAuditsPutPreviousValues) GetOldInternalHours() float32`

GetOldInternalHours returns the OldInternalHours field if non-nil, zero value otherwise.

### GetOldInternalHoursOk

`func (o *OpsAuditsPutPreviousValues) GetOldInternalHoursOk() (*float32, bool)`

GetOldInternalHoursOk returns a tuple with the OldInternalHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldInternalHours

`func (o *OpsAuditsPutPreviousValues) SetOldInternalHours(v float32)`

SetOldInternalHours sets OldInternalHours field to given value.

### HasOldInternalHours

`func (o *OpsAuditsPutPreviousValues) HasOldInternalHours() bool`

HasOldInternalHours returns a boolean if a field has been set.

### GetOldExternalHours

`func (o *OpsAuditsPutPreviousValues) GetOldExternalHours() float32`

GetOldExternalHours returns the OldExternalHours field if non-nil, zero value otherwise.

### GetOldExternalHoursOk

`func (o *OpsAuditsPutPreviousValues) GetOldExternalHoursOk() (*float32, bool)`

GetOldExternalHoursOk returns a tuple with the OldExternalHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldExternalHours

`func (o *OpsAuditsPutPreviousValues) SetOldExternalHours(v float32)`

SetOldExternalHours sets OldExternalHours field to given value.

### HasOldExternalHours

`func (o *OpsAuditsPutPreviousValues) HasOldExternalHours() bool`

HasOldExternalHours returns a boolean if a field has been set.

### GetClosingMeetingDate

`func (o *OpsAuditsPutPreviousValues) GetClosingMeetingDate() string`

GetClosingMeetingDate returns the ClosingMeetingDate field if non-nil, zero value otherwise.

### GetClosingMeetingDateOk

`func (o *OpsAuditsPutPreviousValues) GetClosingMeetingDateOk() (*string, bool)`

GetClosingMeetingDateOk returns a tuple with the ClosingMeetingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingMeetingDate

`func (o *OpsAuditsPutPreviousValues) SetClosingMeetingDate(v string)`

SetClosingMeetingDate sets ClosingMeetingDate field to given value.

### HasClosingMeetingDate

`func (o *OpsAuditsPutPreviousValues) HasClosingMeetingDate() bool`

HasClosingMeetingDate returns a boolean if a field has been set.

### GetOpsAuditCustomDate1

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDate1() string`

GetOpsAuditCustomDate1 returns the OpsAuditCustomDate1 field if non-nil, zero value otherwise.

### GetOpsAuditCustomDate1Ok

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDate1Ok() (*string, bool)`

GetOpsAuditCustomDate1Ok returns a tuple with the OpsAuditCustomDate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditCustomDate1

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditCustomDate1(v string)`

SetOpsAuditCustomDate1 sets OpsAuditCustomDate1 field to given value.

### HasOpsAuditCustomDate1

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditCustomDate1() bool`

HasOpsAuditCustomDate1 returns a boolean if a field has been set.

### GetOpsAuditCustomDate2

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDate2() string`

GetOpsAuditCustomDate2 returns the OpsAuditCustomDate2 field if non-nil, zero value otherwise.

### GetOpsAuditCustomDate2Ok

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDate2Ok() (*string, bool)`

GetOpsAuditCustomDate2Ok returns a tuple with the OpsAuditCustomDate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditCustomDate2

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditCustomDate2(v string)`

SetOpsAuditCustomDate2 sets OpsAuditCustomDate2 field to given value.

### HasOpsAuditCustomDate2

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditCustomDate2() bool`

HasOpsAuditCustomDate2 returns a boolean if a field has been set.

### GetOpsAuditCustomDate3

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDate3() string`

GetOpsAuditCustomDate3 returns the OpsAuditCustomDate3 field if non-nil, zero value otherwise.

### GetOpsAuditCustomDate3Ok

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDate3Ok() (*string, bool)`

GetOpsAuditCustomDate3Ok returns a tuple with the OpsAuditCustomDate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditCustomDate3

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditCustomDate3(v string)`

SetOpsAuditCustomDate3 sets OpsAuditCustomDate3 field to given value.

### HasOpsAuditCustomDate3

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditCustomDate3() bool`

HasOpsAuditCustomDate3 returns a boolean if a field has been set.

### GetOpsAuditCustomDate4

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDate4() string`

GetOpsAuditCustomDate4 returns the OpsAuditCustomDate4 field if non-nil, zero value otherwise.

### GetOpsAuditCustomDate4Ok

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDate4Ok() (*string, bool)`

GetOpsAuditCustomDate4Ok returns a tuple with the OpsAuditCustomDate4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditCustomDate4

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditCustomDate4(v string)`

SetOpsAuditCustomDate4 sets OpsAuditCustomDate4 field to given value.

### HasOpsAuditCustomDate4

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditCustomDate4() bool`

HasOpsAuditCustomDate4 returns a boolean if a field has been set.

### GetOpsAuditCustomDate5

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDate5() string`

GetOpsAuditCustomDate5 returns the OpsAuditCustomDate5 field if non-nil, zero value otherwise.

### GetOpsAuditCustomDate5Ok

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDate5Ok() (*string, bool)`

GetOpsAuditCustomDate5Ok returns a tuple with the OpsAuditCustomDate5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditCustomDate5

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditCustomDate5(v string)`

SetOpsAuditCustomDate5 sets OpsAuditCustomDate5 field to given value.

### HasOpsAuditCustomDate5

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditCustomDate5() bool`

HasOpsAuditCustomDate5 returns a boolean if a field has been set.

### GetOpsAuditCustomDate6

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDate6() string`

GetOpsAuditCustomDate6 returns the OpsAuditCustomDate6 field if non-nil, zero value otherwise.

### GetOpsAuditCustomDate6Ok

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDate6Ok() (*string, bool)`

GetOpsAuditCustomDate6Ok returns a tuple with the OpsAuditCustomDate6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditCustomDate6

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditCustomDate6(v string)`

SetOpsAuditCustomDate6 sets OpsAuditCustomDate6 field to given value.

### HasOpsAuditCustomDate6

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditCustomDate6() bool`

HasOpsAuditCustomDate6 returns a boolean if a field has been set.

### GetCustomText4

`func (o *OpsAuditsPutPreviousValues) GetCustomText4() string`

GetCustomText4 returns the CustomText4 field if non-nil, zero value otherwise.

### GetCustomText4Ok

`func (o *OpsAuditsPutPreviousValues) GetCustomText4Ok() (*string, bool)`

GetCustomText4Ok returns a tuple with the CustomText4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText4

`func (o *OpsAuditsPutPreviousValues) SetCustomText4(v string)`

SetCustomText4 sets CustomText4 field to given value.

### HasCustomText4

`func (o *OpsAuditsPutPreviousValues) HasCustomText4() bool`

HasCustomText4 returns a boolean if a field has been set.

### GetCustomText5

`func (o *OpsAuditsPutPreviousValues) GetCustomText5() string`

GetCustomText5 returns the CustomText5 field if non-nil, zero value otherwise.

### GetCustomText5Ok

`func (o *OpsAuditsPutPreviousValues) GetCustomText5Ok() (*string, bool)`

GetCustomText5Ok returns a tuple with the CustomText5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText5

`func (o *OpsAuditsPutPreviousValues) SetCustomText5(v string)`

SetCustomText5 sets CustomText5 field to given value.

### HasCustomText5

`func (o *OpsAuditsPutPreviousValues) HasCustomText5() bool`

HasCustomText5 returns a boolean if a field has been set.

### GetReactivatedDate

`func (o *OpsAuditsPutPreviousValues) GetReactivatedDate() string`

GetReactivatedDate returns the ReactivatedDate field if non-nil, zero value otherwise.

### GetReactivatedDateOk

`func (o *OpsAuditsPutPreviousValues) GetReactivatedDateOk() (*string, bool)`

GetReactivatedDateOk returns a tuple with the ReactivatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactivatedDate

`func (o *OpsAuditsPutPreviousValues) SetReactivatedDate(v string)`

SetReactivatedDate sets ReactivatedDate field to given value.

### HasReactivatedDate

`func (o *OpsAuditsPutPreviousValues) HasReactivatedDate() bool`

HasReactivatedDate returns a boolean if a field has been set.

### GetReactivatedByUserId

`func (o *OpsAuditsPutPreviousValues) GetReactivatedByUserId() int32`

GetReactivatedByUserId returns the ReactivatedByUserId field if non-nil, zero value otherwise.

### GetReactivatedByUserIdOk

`func (o *OpsAuditsPutPreviousValues) GetReactivatedByUserIdOk() (*int32, bool)`

GetReactivatedByUserIdOk returns a tuple with the ReactivatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactivatedByUserId

`func (o *OpsAuditsPutPreviousValues) SetReactivatedByUserId(v int32)`

SetReactivatedByUserId sets ReactivatedByUserId field to given value.

### HasReactivatedByUserId

`func (o *OpsAuditsPutPreviousValues) HasReactivatedByUserId() bool`

HasReactivatedByUserId returns a boolean if a field has been set.

### GetOpsAuditCustomSelect4OptionId

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomSelect4OptionId() int32`

GetOpsAuditCustomSelect4OptionId returns the OpsAuditCustomSelect4OptionId field if non-nil, zero value otherwise.

### GetOpsAuditCustomSelect4OptionIdOk

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomSelect4OptionIdOk() (*int32, bool)`

GetOpsAuditCustomSelect4OptionIdOk returns a tuple with the OpsAuditCustomSelect4OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditCustomSelect4OptionId

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditCustomSelect4OptionId(v int32)`

SetOpsAuditCustomSelect4OptionId sets OpsAuditCustomSelect4OptionId field to given value.

### HasOpsAuditCustomSelect4OptionId

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditCustomSelect4OptionId() bool`

HasOpsAuditCustomSelect4OptionId returns a boolean if a field has been set.

### GetOpsAuditCustomSelect5OptionId

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomSelect5OptionId() int32`

GetOpsAuditCustomSelect5OptionId returns the OpsAuditCustomSelect5OptionId field if non-nil, zero value otherwise.

### GetOpsAuditCustomSelect5OptionIdOk

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomSelect5OptionIdOk() (*int32, bool)`

GetOpsAuditCustomSelect5OptionIdOk returns a tuple with the OpsAuditCustomSelect5OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditCustomSelect5OptionId

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditCustomSelect5OptionId(v int32)`

SetOpsAuditCustomSelect5OptionId sets OpsAuditCustomSelect5OptionId field to given value.

### HasOpsAuditCustomSelect5OptionId

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditCustomSelect5OptionId() bool`

HasOpsAuditCustomSelect5OptionId returns a boolean if a field has been set.

### GetOpsAuditCustomDate7

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDate7() string`

GetOpsAuditCustomDate7 returns the OpsAuditCustomDate7 field if non-nil, zero value otherwise.

### GetOpsAuditCustomDate7Ok

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDate7Ok() (*string, bool)`

GetOpsAuditCustomDate7Ok returns a tuple with the OpsAuditCustomDate7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditCustomDate7

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditCustomDate7(v string)`

SetOpsAuditCustomDate7 sets OpsAuditCustomDate7 field to given value.

### HasOpsAuditCustomDate7

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditCustomDate7() bool`

HasOpsAuditCustomDate7 returns a boolean if a field has been set.

### GetOpsAuditCustomDate8

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDate8() string`

GetOpsAuditCustomDate8 returns the OpsAuditCustomDate8 field if non-nil, zero value otherwise.

### GetOpsAuditCustomDate8Ok

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDate8Ok() (*string, bool)`

GetOpsAuditCustomDate8Ok returns a tuple with the OpsAuditCustomDate8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditCustomDate8

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditCustomDate8(v string)`

SetOpsAuditCustomDate8 sets OpsAuditCustomDate8 field to given value.

### HasOpsAuditCustomDate8

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditCustomDate8() bool`

HasOpsAuditCustomDate8 returns a boolean if a field has been set.

### GetAdditionalOwnerUsers

`func (o *OpsAuditsPutPreviousValues) GetAdditionalOwnerUsers() interface{}`

GetAdditionalOwnerUsers returns the AdditionalOwnerUsers field if non-nil, zero value otherwise.

### GetAdditionalOwnerUsersOk

`func (o *OpsAuditsPutPreviousValues) GetAdditionalOwnerUsersOk() (*interface{}, bool)`

GetAdditionalOwnerUsersOk returns a tuple with the AdditionalOwnerUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalOwnerUsers

`func (o *OpsAuditsPutPreviousValues) SetAdditionalOwnerUsers(v interface{})`

SetAdditionalOwnerUsers sets AdditionalOwnerUsers field to given value.

### HasAdditionalOwnerUsers

`func (o *OpsAuditsPutPreviousValues) HasAdditionalOwnerUsers() bool`

HasAdditionalOwnerUsers returns a boolean if a field has been set.

### SetAdditionalOwnerUsersNil

`func (o *OpsAuditsPutPreviousValues) SetAdditionalOwnerUsersNil(b bool)`

 SetAdditionalOwnerUsersNil sets the value for AdditionalOwnerUsers to be an explicit nil

### UnsetAdditionalOwnerUsers
`func (o *OpsAuditsPutPreviousValues) UnsetAdditionalOwnerUsers()`

UnsetAdditionalOwnerUsers ensures that no value is present for AdditionalOwnerUsers, not even an explicit nil
### GetBusinessOwnerUsers

`func (o *OpsAuditsPutPreviousValues) GetBusinessOwnerUsers() interface{}`

GetBusinessOwnerUsers returns the BusinessOwnerUsers field if non-nil, zero value otherwise.

### GetBusinessOwnerUsersOk

`func (o *OpsAuditsPutPreviousValues) GetBusinessOwnerUsersOk() (*interface{}, bool)`

GetBusinessOwnerUsersOk returns a tuple with the BusinessOwnerUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessOwnerUsers

`func (o *OpsAuditsPutPreviousValues) SetBusinessOwnerUsers(v interface{})`

SetBusinessOwnerUsers sets BusinessOwnerUsers field to given value.

### HasBusinessOwnerUsers

`func (o *OpsAuditsPutPreviousValues) HasBusinessOwnerUsers() bool`

HasBusinessOwnerUsers returns a boolean if a field has been set.

### SetBusinessOwnerUsersNil

`func (o *OpsAuditsPutPreviousValues) SetBusinessOwnerUsersNil(b bool)`

 SetBusinessOwnerUsersNil sets the value for BusinessOwnerUsers to be an explicit nil

### UnsetBusinessOwnerUsers
`func (o *OpsAuditsPutPreviousValues) UnsetBusinessOwnerUsers()`

UnsetBusinessOwnerUsers ensures that no value is present for BusinessOwnerUsers, not even an explicit nil
### GetScopes

`func (o *OpsAuditsPutPreviousValues) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OpsAuditsPutPreviousValues) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OpsAuditsPutPreviousValues) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *OpsAuditsPutPreviousValues) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *OpsAuditsPutPreviousValues) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *OpsAuditsPutPreviousValues) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetCustomText6

`func (o *OpsAuditsPutPreviousValues) GetCustomText6() string`

GetCustomText6 returns the CustomText6 field if non-nil, zero value otherwise.

### GetCustomText6Ok

`func (o *OpsAuditsPutPreviousValues) GetCustomText6Ok() (*string, bool)`

GetCustomText6Ok returns a tuple with the CustomText6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText6

`func (o *OpsAuditsPutPreviousValues) SetCustomText6(v string)`

SetCustomText6 sets CustomText6 field to given value.

### HasCustomText6

`func (o *OpsAuditsPutPreviousValues) HasCustomText6() bool`

HasCustomText6 returns a boolean if a field has been set.

### GetCustomText7

`func (o *OpsAuditsPutPreviousValues) GetCustomText7() string`

GetCustomText7 returns the CustomText7 field if non-nil, zero value otherwise.

### GetCustomText7Ok

`func (o *OpsAuditsPutPreviousValues) GetCustomText7Ok() (*string, bool)`

GetCustomText7Ok returns a tuple with the CustomText7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText7

`func (o *OpsAuditsPutPreviousValues) SetCustomText7(v string)`

SetCustomText7 sets CustomText7 field to given value.

### HasCustomText7

`func (o *OpsAuditsPutPreviousValues) HasCustomText7() bool`

HasCustomText7 returns a boolean if a field has been set.

### GetCustomText8

`func (o *OpsAuditsPutPreviousValues) GetCustomText8() string`

GetCustomText8 returns the CustomText8 field if non-nil, zero value otherwise.

### GetCustomText8Ok

`func (o *OpsAuditsPutPreviousValues) GetCustomText8Ok() (*string, bool)`

GetCustomText8Ok returns a tuple with the CustomText8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText8

`func (o *OpsAuditsPutPreviousValues) SetCustomText8(v string)`

SetCustomText8 sets CustomText8 field to given value.

### HasCustomText8

`func (o *OpsAuditsPutPreviousValues) HasCustomText8() bool`

HasCustomText8 returns a boolean if a field has been set.

### GetCustomText9

`func (o *OpsAuditsPutPreviousValues) GetCustomText9() string`

GetCustomText9 returns the CustomText9 field if non-nil, zero value otherwise.

### GetCustomText9Ok

`func (o *OpsAuditsPutPreviousValues) GetCustomText9Ok() (*string, bool)`

GetCustomText9Ok returns a tuple with the CustomText9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText9

`func (o *OpsAuditsPutPreviousValues) SetCustomText9(v string)`

SetCustomText9 sets CustomText9 field to given value.

### HasCustomText9

`func (o *OpsAuditsPutPreviousValues) HasCustomText9() bool`

HasCustomText9 returns a boolean if a field has been set.

### GetCustomText10

`func (o *OpsAuditsPutPreviousValues) GetCustomText10() string`

GetCustomText10 returns the CustomText10 field if non-nil, zero value otherwise.

### GetCustomText10Ok

`func (o *OpsAuditsPutPreviousValues) GetCustomText10Ok() (*string, bool)`

GetCustomText10Ok returns a tuple with the CustomText10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText10

`func (o *OpsAuditsPutPreviousValues) SetCustomText10(v string)`

SetCustomText10 sets CustomText10 field to given value.

### HasCustomText10

`func (o *OpsAuditsPutPreviousValues) HasCustomText10() bool`

HasCustomText10 returns a boolean if a field has been set.

### GetCustomText11

`func (o *OpsAuditsPutPreviousValues) GetCustomText11() string`

GetCustomText11 returns the CustomText11 field if non-nil, zero value otherwise.

### GetCustomText11Ok

`func (o *OpsAuditsPutPreviousValues) GetCustomText11Ok() (*string, bool)`

GetCustomText11Ok returns a tuple with the CustomText11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText11

`func (o *OpsAuditsPutPreviousValues) SetCustomText11(v string)`

SetCustomText11 sets CustomText11 field to given value.

### HasCustomText11

`func (o *OpsAuditsPutPreviousValues) HasCustomText11() bool`

HasCustomText11 returns a boolean if a field has been set.

### GetFieldData

`func (o *OpsAuditsPutPreviousValues) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *OpsAuditsPutPreviousValues) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *OpsAuditsPutPreviousValues) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *OpsAuditsPutPreviousValues) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *OpsAuditsPutPreviousValues) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *OpsAuditsPutPreviousValues) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil
### GetOpsAuditAuditReportStatusId

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditAuditReportStatusId() int32`

GetOpsAuditAuditReportStatusId returns the OpsAuditAuditReportStatusId field if non-nil, zero value otherwise.

### GetOpsAuditAuditReportStatusIdOk

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditAuditReportStatusIdOk() (*int32, bool)`

GetOpsAuditAuditReportStatusIdOk returns a tuple with the OpsAuditAuditReportStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditAuditReportStatusId

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditAuditReportStatusId(v int32)`

SetOpsAuditAuditReportStatusId sets OpsAuditAuditReportStatusId field to given value.

### HasOpsAuditAuditReportStatusId

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditAuditReportStatusId() bool`

HasOpsAuditAuditReportStatusId returns a boolean if a field has been set.

### GetOpsAuditTeamResponsibleId

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditTeamResponsibleId() int32`

GetOpsAuditTeamResponsibleId returns the OpsAuditTeamResponsibleId field if non-nil, zero value otherwise.

### GetOpsAuditTeamResponsibleIdOk

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditTeamResponsibleIdOk() (*int32, bool)`

GetOpsAuditTeamResponsibleIdOk returns a tuple with the OpsAuditTeamResponsibleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditTeamResponsibleId

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditTeamResponsibleId(v int32)`

SetOpsAuditTeamResponsibleId sets OpsAuditTeamResponsibleId field to given value.

### HasOpsAuditTeamResponsibleId

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditTeamResponsibleId() bool`

HasOpsAuditTeamResponsibleId returns a boolean if a field has been set.

### GetOpsAuditAcpId

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditAcpId() int32`

GetOpsAuditAcpId returns the OpsAuditAcpId field if non-nil, zero value otherwise.

### GetOpsAuditAcpIdOk

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditAcpIdOk() (*int32, bool)`

GetOpsAuditAcpIdOk returns a tuple with the OpsAuditAcpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditAcpId

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditAcpId(v int32)`

SetOpsAuditAcpId sets OpsAuditAcpId field to given value.

### HasOpsAuditAcpId

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditAcpId() bool`

HasOpsAuditAcpId returns a boolean if a field has been set.

### GetOpsAuditOtherRatingId

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditOtherRatingId() int32`

GetOpsAuditOtherRatingId returns the OpsAuditOtherRatingId field if non-nil, zero value otherwise.

### GetOpsAuditOtherRatingIdOk

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditOtherRatingIdOk() (*int32, bool)`

GetOpsAuditOtherRatingIdOk returns a tuple with the OpsAuditOtherRatingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditOtherRatingId

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditOtherRatingId(v int32)`

SetOpsAuditOtherRatingId sets OpsAuditOtherRatingId field to given value.

### HasOpsAuditOtherRatingId

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditOtherRatingId() bool`

HasOpsAuditOtherRatingId returns a boolean if a field has been set.

### GetOpsAuditDivisionId

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditDivisionId() int32`

GetOpsAuditDivisionId returns the OpsAuditDivisionId field if non-nil, zero value otherwise.

### GetOpsAuditDivisionIdOk

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditDivisionIdOk() (*int32, bool)`

GetOpsAuditDivisionIdOk returns a tuple with the OpsAuditDivisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditDivisionId

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditDivisionId(v int32)`

SetOpsAuditDivisionId sets OpsAuditDivisionId field to given value.

### HasOpsAuditDivisionId

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditDivisionId() bool`

HasOpsAuditDivisionId returns a boolean if a field has been set.

### GetProjectManagerUserId

`func (o *OpsAuditsPutPreviousValues) GetProjectManagerUserId() int32`

GetProjectManagerUserId returns the ProjectManagerUserId field if non-nil, zero value otherwise.

### GetProjectManagerUserIdOk

`func (o *OpsAuditsPutPreviousValues) GetProjectManagerUserIdOk() (*int32, bool)`

GetProjectManagerUserIdOk returns a tuple with the ProjectManagerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectManagerUserId

`func (o *OpsAuditsPutPreviousValues) SetProjectManagerUserId(v int32)`

SetProjectManagerUserId sets ProjectManagerUserId field to given value.

### HasProjectManagerUserId

`func (o *OpsAuditsPutPreviousValues) HasProjectManagerUserId() bool`

HasProjectManagerUserId returns a boolean if a field has been set.

### GetProjectCoordinatorUserId

`func (o *OpsAuditsPutPreviousValues) GetProjectCoordinatorUserId() int32`

GetProjectCoordinatorUserId returns the ProjectCoordinatorUserId field if non-nil, zero value otherwise.

### GetProjectCoordinatorUserIdOk

`func (o *OpsAuditsPutPreviousValues) GetProjectCoordinatorUserIdOk() (*int32, bool)`

GetProjectCoordinatorUserIdOk returns a tuple with the ProjectCoordinatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCoordinatorUserId

`func (o *OpsAuditsPutPreviousValues) SetProjectCoordinatorUserId(v int32)`

SetProjectCoordinatorUserId sets ProjectCoordinatorUserId field to given value.

### HasProjectCoordinatorUserId

`func (o *OpsAuditsPutPreviousValues) HasProjectCoordinatorUserId() bool`

HasProjectCoordinatorUserId returns a boolean if a field has been set.

### GetFlattened

`func (o *OpsAuditsPutPreviousValues) GetFlattened() interface{}`

GetFlattened returns the Flattened field if non-nil, zero value otherwise.

### GetFlattenedOk

`func (o *OpsAuditsPutPreviousValues) GetFlattenedOk() (*interface{}, bool)`

GetFlattenedOk returns a tuple with the Flattened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlattened

`func (o *OpsAuditsPutPreviousValues) SetFlattened(v interface{})`

SetFlattened sets Flattened field to given value.

### HasFlattened

`func (o *OpsAuditsPutPreviousValues) HasFlattened() bool`

HasFlattened returns a boolean if a field has been set.

### SetFlattenedNil

`func (o *OpsAuditsPutPreviousValues) SetFlattenedNil(b bool)`

 SetFlattenedNil sets the value for Flattened to be an explicit nil

### UnsetFlattened
`func (o *OpsAuditsPutPreviousValues) UnsetFlattened()`

UnsetFlattened ensures that no value is present for Flattened, not even an explicit nil
### GetIsFlattened

`func (o *OpsAuditsPutPreviousValues) GetIsFlattened() bool`

GetIsFlattened returns the IsFlattened field if non-nil, zero value otherwise.

### GetIsFlattenedOk

`func (o *OpsAuditsPutPreviousValues) GetIsFlattenedOk() (*bool, bool)`

GetIsFlattenedOk returns a tuple with the IsFlattened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlattened

`func (o *OpsAuditsPutPreviousValues) SetIsFlattened(v bool)`

SetIsFlattened sets IsFlattened field to given value.

### HasIsFlattened

`func (o *OpsAuditsPutPreviousValues) HasIsFlattened() bool`

HasIsFlattened returns a boolean if a field has been set.

### GetPurpose

`func (o *OpsAuditsPutPreviousValues) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *OpsAuditsPutPreviousValues) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *OpsAuditsPutPreviousValues) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *OpsAuditsPutPreviousValues) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetScope

`func (o *OpsAuditsPutPreviousValues) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OpsAuditsPutPreviousValues) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OpsAuditsPutPreviousValues) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OpsAuditsPutPreviousValues) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetDraftReportDate

`func (o *OpsAuditsPutPreviousValues) GetDraftReportDate() string`

GetDraftReportDate returns the DraftReportDate field if non-nil, zero value otherwise.

### GetDraftReportDateOk

`func (o *OpsAuditsPutPreviousValues) GetDraftReportDateOk() (*string, bool)`

GetDraftReportDateOk returns a tuple with the DraftReportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftReportDate

`func (o *OpsAuditsPutPreviousValues) SetDraftReportDate(v string)`

SetDraftReportDate sets DraftReportDate field to given value.

### HasDraftReportDate

`func (o *OpsAuditsPutPreviousValues) HasDraftReportDate() bool`

HasDraftReportDate returns a boolean if a field has been set.

### GetOpeningMeetingDate

`func (o *OpsAuditsPutPreviousValues) GetOpeningMeetingDate() string`

GetOpeningMeetingDate returns the OpeningMeetingDate field if non-nil, zero value otherwise.

### GetOpeningMeetingDateOk

`func (o *OpsAuditsPutPreviousValues) GetOpeningMeetingDateOk() (*string, bool)`

GetOpeningMeetingDateOk returns a tuple with the OpeningMeetingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningMeetingDate

`func (o *OpsAuditsPutPreviousValues) SetOpeningMeetingDate(v string)`

SetOpeningMeetingDate sets OpeningMeetingDate field to given value.

### HasOpeningMeetingDate

`func (o *OpsAuditsPutPreviousValues) HasOpeningMeetingDate() bool`

HasOpeningMeetingDate returns a boolean if a field has been set.

### GetFrameworkId

`func (o *OpsAuditsPutPreviousValues) GetFrameworkId() int32`

GetFrameworkId returns the FrameworkId field if non-nil, zero value otherwise.

### GetFrameworkIdOk

`func (o *OpsAuditsPutPreviousValues) GetFrameworkIdOk() (*int32, bool)`

GetFrameworkIdOk returns a tuple with the FrameworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkId

`func (o *OpsAuditsPutPreviousValues) SetFrameworkId(v int32)`

SetFrameworkId sets FrameworkId field to given value.

### HasFrameworkId

`func (o *OpsAuditsPutPreviousValues) HasFrameworkId() bool`

HasFrameworkId returns a boolean if a field has been set.

### GetPlanningEndDate

`func (o *OpsAuditsPutPreviousValues) GetPlanningEndDate() string`

GetPlanningEndDate returns the PlanningEndDate field if non-nil, zero value otherwise.

### GetPlanningEndDateOk

`func (o *OpsAuditsPutPreviousValues) GetPlanningEndDateOk() (*string, bool)`

GetPlanningEndDateOk returns a tuple with the PlanningEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanningEndDate

`func (o *OpsAuditsPutPreviousValues) SetPlanningEndDate(v string)`

SetPlanningEndDate sets PlanningEndDate field to given value.

### HasPlanningEndDate

`func (o *OpsAuditsPutPreviousValues) HasPlanningEndDate() bool`

HasPlanningEndDate returns a boolean if a field has been set.

### GetFieldworkEndDate

`func (o *OpsAuditsPutPreviousValues) GetFieldworkEndDate() string`

GetFieldworkEndDate returns the FieldworkEndDate field if non-nil, zero value otherwise.

### GetFieldworkEndDateOk

`func (o *OpsAuditsPutPreviousValues) GetFieldworkEndDateOk() (*string, bool)`

GetFieldworkEndDateOk returns a tuple with the FieldworkEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldworkEndDate

`func (o *OpsAuditsPutPreviousValues) SetFieldworkEndDate(v string)`

SetFieldworkEndDate sets FieldworkEndDate field to given value.

### HasFieldworkEndDate

`func (o *OpsAuditsPutPreviousValues) HasFieldworkEndDate() bool`

HasFieldworkEndDate returns a boolean if a field has been set.

### GetBudgetHours

`func (o *OpsAuditsPutPreviousValues) GetBudgetHours() float32`

GetBudgetHours returns the BudgetHours field if non-nil, zero value otherwise.

### GetBudgetHoursOk

`func (o *OpsAuditsPutPreviousValues) GetBudgetHoursOk() (*float32, bool)`

GetBudgetHoursOk returns a tuple with the BudgetHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetHours

`func (o *OpsAuditsPutPreviousValues) SetBudgetHours(v float32)`

SetBudgetHours sets BudgetHours field to given value.

### HasBudgetHours

`func (o *OpsAuditsPutPreviousValues) HasBudgetHours() bool`

HasBudgetHours returns a boolean if a field has been set.

### GetActualHours

`func (o *OpsAuditsPutPreviousValues) GetActualHours() float32`

GetActualHours returns the ActualHours field if non-nil, zero value otherwise.

### GetActualHoursOk

`func (o *OpsAuditsPutPreviousValues) GetActualHoursOk() (*float32, bool)`

GetActualHoursOk returns a tuple with the ActualHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualHours

`func (o *OpsAuditsPutPreviousValues) SetActualHours(v float32)`

SetActualHours sets ActualHours field to given value.

### HasActualHours

`func (o *OpsAuditsPutPreviousValues) HasActualHours() bool`

HasActualHours returns a boolean if a field has been set.

### GetExternalHours

`func (o *OpsAuditsPutPreviousValues) GetExternalHours() float32`

GetExternalHours returns the ExternalHours field if non-nil, zero value otherwise.

### GetExternalHoursOk

`func (o *OpsAuditsPutPreviousValues) GetExternalHoursOk() (*float32, bool)`

GetExternalHoursOk returns a tuple with the ExternalHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHours

`func (o *OpsAuditsPutPreviousValues) SetExternalHours(v float32)`

SetExternalHours sets ExternalHours field to given value.

### HasExternalHours

`func (o *OpsAuditsPutPreviousValues) HasExternalHours() bool`

HasExternalHours returns a boolean if a field has been set.

### GetInternalHours

`func (o *OpsAuditsPutPreviousValues) GetInternalHours() float32`

GetInternalHours returns the InternalHours field if non-nil, zero value otherwise.

### GetInternalHoursOk

`func (o *OpsAuditsPutPreviousValues) GetInternalHoursOk() (*float32, bool)`

GetInternalHoursOk returns a tuple with the InternalHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHours

`func (o *OpsAuditsPutPreviousValues) SetInternalHours(v float32)`

SetInternalHours sets InternalHours field to given value.

### HasInternalHours

`func (o *OpsAuditsPutPreviousValues) HasInternalHours() bool`

HasInternalHours returns a boolean if a field has been set.

### GetTemplateId

`func (o *OpsAuditsPutPreviousValues) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *OpsAuditsPutPreviousValues) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *OpsAuditsPutPreviousValues) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *OpsAuditsPutPreviousValues) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetIsProcessing

`func (o *OpsAuditsPutPreviousValues) GetIsProcessing() bool`

GetIsProcessing returns the IsProcessing field if non-nil, zero value otherwise.

### GetIsProcessingOk

`func (o *OpsAuditsPutPreviousValues) GetIsProcessingOk() (*bool, bool)`

GetIsProcessingOk returns a tuple with the IsProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProcessing

`func (o *OpsAuditsPutPreviousValues) SetIsProcessing(v bool)`

SetIsProcessing sets IsProcessing field to given value.

### HasIsProcessing

`func (o *OpsAuditsPutPreviousValues) HasIsProcessing() bool`

HasIsProcessing returns a boolean if a field has been set.

### GetAssessmentTemplateId

`func (o *OpsAuditsPutPreviousValues) GetAssessmentTemplateId() int32`

GetAssessmentTemplateId returns the AssessmentTemplateId field if non-nil, zero value otherwise.

### GetAssessmentTemplateIdOk

`func (o *OpsAuditsPutPreviousValues) GetAssessmentTemplateIdOk() (*int32, bool)`

GetAssessmentTemplateIdOk returns a tuple with the AssessmentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessmentTemplateId

`func (o *OpsAuditsPutPreviousValues) SetAssessmentTemplateId(v int32)`

SetAssessmentTemplateId sets AssessmentTemplateId field to given value.

### HasAssessmentTemplateId

`func (o *OpsAuditsPutPreviousValues) HasAssessmentTemplateId() bool`

HasAssessmentTemplateId returns a boolean if a field has been set.

### GetAllowedAssessableType

`func (o *OpsAuditsPutPreviousValues) GetAllowedAssessableType() string`

GetAllowedAssessableType returns the AllowedAssessableType field if non-nil, zero value otherwise.

### GetAllowedAssessableTypeOk

`func (o *OpsAuditsPutPreviousValues) GetAllowedAssessableTypeOk() (*string, bool)`

GetAllowedAssessableTypeOk returns a tuple with the AllowedAssessableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAssessableType

`func (o *OpsAuditsPutPreviousValues) SetAllowedAssessableType(v string)`

SetAllowedAssessableType sets AllowedAssessableType field to given value.

### HasAllowedAssessableType

`func (o *OpsAuditsPutPreviousValues) HasAllowedAssessableType() bool`

HasAllowedAssessableType returns a boolean if a field has been set.

### GetIsRiskAssessmentEnabled

`func (o *OpsAuditsPutPreviousValues) GetIsRiskAssessmentEnabled() bool`

GetIsRiskAssessmentEnabled returns the IsRiskAssessmentEnabled field if non-nil, zero value otherwise.

### GetIsRiskAssessmentEnabledOk

`func (o *OpsAuditsPutPreviousValues) GetIsRiskAssessmentEnabledOk() (*bool, bool)`

GetIsRiskAssessmentEnabledOk returns a tuple with the IsRiskAssessmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRiskAssessmentEnabled

`func (o *OpsAuditsPutPreviousValues) SetIsRiskAssessmentEnabled(v bool)`

SetIsRiskAssessmentEnabled sets IsRiskAssessmentEnabled field to given value.

### HasIsRiskAssessmentEnabled

`func (o *OpsAuditsPutPreviousValues) HasIsRiskAssessmentEnabled() bool`

HasIsRiskAssessmentEnabled returns a boolean if a field has been set.

### GetAssessmentId

`func (o *OpsAuditsPutPreviousValues) GetAssessmentId() int32`

GetAssessmentId returns the AssessmentId field if non-nil, zero value otherwise.

### GetAssessmentIdOk

`func (o *OpsAuditsPutPreviousValues) GetAssessmentIdOk() (*int32, bool)`

GetAssessmentIdOk returns a tuple with the AssessmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessmentId

`func (o *OpsAuditsPutPreviousValues) SetAssessmentId(v int32)`

SetAssessmentId sets AssessmentId field to given value.

### HasAssessmentId

`func (o *OpsAuditsPutPreviousValues) HasAssessmentId() bool`

HasAssessmentId returns a boolean if a field has been set.

### GetProjectObjectives

`func (o *OpsAuditsPutPreviousValues) GetProjectObjectives() string`

GetProjectObjectives returns the ProjectObjectives field if non-nil, zero value otherwise.

### GetProjectObjectivesOk

`func (o *OpsAuditsPutPreviousValues) GetProjectObjectivesOk() (*string, bool)`

GetProjectObjectivesOk returns a tuple with the ProjectObjectives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectObjectives

`func (o *OpsAuditsPutPreviousValues) SetProjectObjectives(v string)`

SetProjectObjectives sets ProjectObjectives field to given value.

### HasProjectObjectives

`func (o *OpsAuditsPutPreviousValues) HasProjectObjectives() bool`

HasProjectObjectives returns a boolean if a field has been set.

### GetProjectDrivers

`func (o *OpsAuditsPutPreviousValues) GetProjectDrivers() string`

GetProjectDrivers returns the ProjectDrivers field if non-nil, zero value otherwise.

### GetProjectDriversOk

`func (o *OpsAuditsPutPreviousValues) GetProjectDriversOk() (*string, bool)`

GetProjectDriversOk returns a tuple with the ProjectDrivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDrivers

`func (o *OpsAuditsPutPreviousValues) SetProjectDrivers(v string)`

SetProjectDrivers sets ProjectDrivers field to given value.

### HasProjectDrivers

`func (o *OpsAuditsPutPreviousValues) HasProjectDrivers() bool`

HasProjectDrivers returns a boolean if a field has been set.

### GetKeyMilestones

`func (o *OpsAuditsPutPreviousValues) GetKeyMilestones() string`

GetKeyMilestones returns the KeyMilestones field if non-nil, zero value otherwise.

### GetKeyMilestonesOk

`func (o *OpsAuditsPutPreviousValues) GetKeyMilestonesOk() (*string, bool)`

GetKeyMilestonesOk returns a tuple with the KeyMilestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyMilestones

`func (o *OpsAuditsPutPreviousValues) SetKeyMilestones(v string)`

SetKeyMilestones sets KeyMilestones field to given value.

### HasKeyMilestones

`func (o *OpsAuditsPutPreviousValues) HasKeyMilestones() bool`

HasKeyMilestones returns a boolean if a field has been set.

### GetPriorProjectKnowledge

`func (o *OpsAuditsPutPreviousValues) GetPriorProjectKnowledge() string`

GetPriorProjectKnowledge returns the PriorProjectKnowledge field if non-nil, zero value otherwise.

### GetPriorProjectKnowledgeOk

`func (o *OpsAuditsPutPreviousValues) GetPriorProjectKnowledgeOk() (*string, bool)`

GetPriorProjectKnowledgeOk returns a tuple with the PriorProjectKnowledge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorProjectKnowledge

`func (o *OpsAuditsPutPreviousValues) SetPriorProjectKnowledge(v string)`

SetPriorProjectKnowledge sets PriorProjectKnowledge field to given value.

### HasPriorProjectKnowledge

`func (o *OpsAuditsPutPreviousValues) HasPriorProjectKnowledge() bool`

HasPriorProjectKnowledge returns a boolean if a field has been set.

### GetHighLevelScope

`func (o *OpsAuditsPutPreviousValues) GetHighLevelScope() string`

GetHighLevelScope returns the HighLevelScope field if non-nil, zero value otherwise.

### GetHighLevelScopeOk

`func (o *OpsAuditsPutPreviousValues) GetHighLevelScopeOk() (*string, bool)`

GetHighLevelScopeOk returns a tuple with the HighLevelScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighLevelScope

`func (o *OpsAuditsPutPreviousValues) SetHighLevelScope(v string)`

SetHighLevelScope sets HighLevelScope field to given value.

### HasHighLevelScope

`func (o *OpsAuditsPutPreviousValues) HasHighLevelScope() bool`

HasHighLevelScope returns a boolean if a field has been set.

### GetScopingConsiderationsCompletionDate

`func (o *OpsAuditsPutPreviousValues) GetScopingConsiderationsCompletionDate() string`

GetScopingConsiderationsCompletionDate returns the ScopingConsiderationsCompletionDate field if non-nil, zero value otherwise.

### GetScopingConsiderationsCompletionDateOk

`func (o *OpsAuditsPutPreviousValues) GetScopingConsiderationsCompletionDateOk() (*string, bool)`

GetScopingConsiderationsCompletionDateOk returns a tuple with the ScopingConsiderationsCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopingConsiderationsCompletionDate

`func (o *OpsAuditsPutPreviousValues) SetScopingConsiderationsCompletionDate(v string)`

SetScopingConsiderationsCompletionDate sets ScopingConsiderationsCompletionDate field to given value.

### HasScopingConsiderationsCompletionDate

`func (o *OpsAuditsPutPreviousValues) HasScopingConsiderationsCompletionDate() bool`

HasScopingConsiderationsCompletionDate returns a boolean if a field has been set.

### GetScopingConsiderationsLastUpdated

`func (o *OpsAuditsPutPreviousValues) GetScopingConsiderationsLastUpdated() string`

GetScopingConsiderationsLastUpdated returns the ScopingConsiderationsLastUpdated field if non-nil, zero value otherwise.

### GetScopingConsiderationsLastUpdatedOk

`func (o *OpsAuditsPutPreviousValues) GetScopingConsiderationsLastUpdatedOk() (*string, bool)`

GetScopingConsiderationsLastUpdatedOk returns a tuple with the ScopingConsiderationsLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopingConsiderationsLastUpdated

`func (o *OpsAuditsPutPreviousValues) SetScopingConsiderationsLastUpdated(v string)`

SetScopingConsiderationsLastUpdated sets ScopingConsiderationsLastUpdated field to given value.

### HasScopingConsiderationsLastUpdated

`func (o *OpsAuditsPutPreviousValues) HasScopingConsiderationsLastUpdated() bool`

HasScopingConsiderationsLastUpdated returns a boolean if a field has been set.

### GetOpsAuditCustomDecimal1

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDecimal1() float32`

GetOpsAuditCustomDecimal1 returns the OpsAuditCustomDecimal1 field if non-nil, zero value otherwise.

### GetOpsAuditCustomDecimal1Ok

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDecimal1Ok() (*float32, bool)`

GetOpsAuditCustomDecimal1Ok returns a tuple with the OpsAuditCustomDecimal1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditCustomDecimal1

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditCustomDecimal1(v float32)`

SetOpsAuditCustomDecimal1 sets OpsAuditCustomDecimal1 field to given value.

### HasOpsAuditCustomDecimal1

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditCustomDecimal1() bool`

HasOpsAuditCustomDecimal1 returns a boolean if a field has been set.

### GetOpsAuditCustomDecimal2

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDecimal2() float32`

GetOpsAuditCustomDecimal2 returns the OpsAuditCustomDecimal2 field if non-nil, zero value otherwise.

### GetOpsAuditCustomDecimal2Ok

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDecimal2Ok() (*float32, bool)`

GetOpsAuditCustomDecimal2Ok returns a tuple with the OpsAuditCustomDecimal2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditCustomDecimal2

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditCustomDecimal2(v float32)`

SetOpsAuditCustomDecimal2 sets OpsAuditCustomDecimal2 field to given value.

### HasOpsAuditCustomDecimal2

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditCustomDecimal2() bool`

HasOpsAuditCustomDecimal2 returns a boolean if a field has been set.

### GetOpsAuditCustomDecimal3

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDecimal3() float32`

GetOpsAuditCustomDecimal3 returns the OpsAuditCustomDecimal3 field if non-nil, zero value otherwise.

### GetOpsAuditCustomDecimal3Ok

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDecimal3Ok() (*float32, bool)`

GetOpsAuditCustomDecimal3Ok returns a tuple with the OpsAuditCustomDecimal3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditCustomDecimal3

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditCustomDecimal3(v float32)`

SetOpsAuditCustomDecimal3 sets OpsAuditCustomDecimal3 field to given value.

### HasOpsAuditCustomDecimal3

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditCustomDecimal3() bool`

HasOpsAuditCustomDecimal3 returns a boolean if a field has been set.

### GetOpsAuditCustomDecimal4

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDecimal4() float32`

GetOpsAuditCustomDecimal4 returns the OpsAuditCustomDecimal4 field if non-nil, zero value otherwise.

### GetOpsAuditCustomDecimal4Ok

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDecimal4Ok() (*float32, bool)`

GetOpsAuditCustomDecimal4Ok returns a tuple with the OpsAuditCustomDecimal4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditCustomDecimal4

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditCustomDecimal4(v float32)`

SetOpsAuditCustomDecimal4 sets OpsAuditCustomDecimal4 field to given value.

### HasOpsAuditCustomDecimal4

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditCustomDecimal4() bool`

HasOpsAuditCustomDecimal4 returns a boolean if a field has been set.

### GetOpsAuditCustomDecimal5

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDecimal5() float32`

GetOpsAuditCustomDecimal5 returns the OpsAuditCustomDecimal5 field if non-nil, zero value otherwise.

### GetOpsAuditCustomDecimal5Ok

`func (o *OpsAuditsPutPreviousValues) GetOpsAuditCustomDecimal5Ok() (*float32, bool)`

GetOpsAuditCustomDecimal5Ok returns a tuple with the OpsAuditCustomDecimal5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditCustomDecimal5

`func (o *OpsAuditsPutPreviousValues) SetOpsAuditCustomDecimal5(v float32)`

SetOpsAuditCustomDecimal5 sets OpsAuditCustomDecimal5 field to given value.

### HasOpsAuditCustomDecimal5

`func (o *OpsAuditsPutPreviousValues) HasOpsAuditCustomDecimal5() bool`

HasOpsAuditCustomDecimal5 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


