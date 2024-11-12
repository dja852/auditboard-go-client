# IssuesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**PotentialRisk** | Pointer to **string** |  | [optional] 
**IndividualsPresent** | Pointer to **string** |  | [optional] 
**DeficiencyLevelId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;deficiency_levels.id&#x60;.&lt;fk table&#x3D;&#39;deficiency_levels&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**QualitativeFactors** | Pointer to **string** |  | [optional] 
**IssueIdentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_idents.id&#x60;.&lt;fk table&#x3D;&#39;issue_idents&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**GrossExposure** | Pointer to **string** |  | [optional] 
**MagnitudeId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;magnitudes.id&#x60;.&lt;fk table&#x3D;&#39;magnitudes&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**LikelihoodId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;likelihoods.id&#x60;.&lt;fk table&#x3D;&#39;likelihoods&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AggregationReferenceId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;aggregation_references.id&#x60;.&lt;fk table&#x3D;&#39;aggregation_references&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AdjustedExposure** | Pointer to **string** |  | [optional] 
**CreatorUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IdentifiedDate** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**ControlsDatumId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;controls_data.id&#x60;.&lt;fk table&#x3D;&#39;controls_data&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;tests.id&#x60;.&lt;fk table&#x3D;&#39;tests&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestSectionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_sections.id&#x60;.&lt;fk table&#x3D;&#39;test_sections&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**ReferenceIssueId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issues.id&#x60;.&lt;fk table&#x3D;&#39;issues&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ArchiveId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;archives.id&#x60;.&lt;fk table&#x3D;&#39;archives&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TesterUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ReviewerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audits.id&#x60;.&lt;fk table&#x3D;&#39;ops_audits&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditItemId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_items.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_items&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**InternalControlComponentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;internal_control_components.id&#x60;.&lt;fk table&#x3D;&#39;internal_control_components&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ExternalPlannedConfirmDate** | Pointer to **string** |  | [optional] 
**Disclosure** | Pointer to **string** |  | [optional] 
**ExternalConfirmDate** | Pointer to **string** |  | [optional] 
**PendingRemediationDate** | Pointer to **string** |  | [optional] 
**ClosedDate** | Pointer to **string** |  | [optional] 
**InactiveDate** | Pointer to **string** |  | [optional] 
**ReopenDate** | Pointer to **string** |  | [optional] 
**OpenDate** | Pointer to **string** |  | [optional] 
**AmendDate** | Pointer to **string** |  | [optional] 
**RemediatedDate** | Pointer to **string** |  | [optional] 
**OpenByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ReopenByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**PendingRemediationByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RemediatedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ClosedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AmendByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Flattened** | Pointer to **interface{}** |  | [optional] 
**IssueRatingId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_ratings.id&#x60;.&lt;fk table&#x3D;&#39;issue_ratings&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**SoxImpactId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;sox_impacts.id&#x60;.&lt;fk table&#x3D;&#39;sox_impacts&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueCustomSelect1OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_custom_select1_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_custom_select1_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueCustomSelect2OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_custom_select2_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_custom_select2_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**VicePresidentUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IsFlattened** | Pointer to **bool** |  | [optional] [default to false]
**RemediationOwners** | Pointer to **interface{}** |  | [optional] 
**ReferenceMeta** | Pointer to **interface{}** |  | [optional] 
**Totals** | Pointer to **interface{}** |  | [optional] 
**IsArchived** | Pointer to **bool** |  | [optional] [default to false]
**IssueCustomSelect3OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_custom_select3_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_custom_select3_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AdditionalAuditorUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**InternalAuditValidationDate** | Pointer to **string** |  | [optional] 
**CustomText1** | Pointer to **string** |  | [optional] 
**CustomText2** | Pointer to **string** |  | [optional] 
**CustomText3** | Pointer to **string** |  | [optional] 
**CustomText4** | Pointer to **string** |  | [optional] 
**CustomText5** | Pointer to **string** |  | [optional] 
**CustomText6** | Pointer to **string** |  | [optional] 
**IssueCustomSelect4OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_custom_select4_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_custom_select4_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueCustomSelect5OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_custom_select5_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_custom_select5_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueCustomSelect6OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_custom_select6_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_custom_select6_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueCustomSelect7OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_custom_select7_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_custom_select7_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueCustomSelect8OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_custom_select8_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_custom_select8_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueCategoryId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_categories.id&#x60;.&lt;fk table&#x3D;&#39;issue_categories&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssuesourceableId** | Pointer to **int32** |  | [optional] 
**IssuesourceableType** | Pointer to **string** |  | [optional] 
**FinanceProcessOwners** | Pointer to **interface{}** |  | [optional] 
**BusinessProcessOwners** | Pointer to **interface{}** |  | [optional] 
**Scopes** | Pointer to **interface{}** |  | [optional] 
**ExecutiveUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**SeniorOwnerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FirstRemediatedDate** | Pointer to **string** |  | [optional] 
**FirstPendingRemediationDate** | Pointer to **string** |  | [optional] 
**FirstClosedDate** | Pointer to **string** |  | [optional] 
**CustomCurrency1** | Pointer to **float32** |  | [optional] 
**CustomCurrency2** | Pointer to **float32** |  | [optional] 
**CustomCurrency3** | Pointer to **float32** |  | [optional] 
**CustomCurrency4** | Pointer to **float32** |  | [optional] 
**CustomDecimal1** | Pointer to **float32** |  | [optional] 
**CustomDecimal2** | Pointer to **float32** |  | [optional] 
**CustomDecimal3** | Pointer to **float32** |  | [optional] 
**FindingOwnerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomDate1** | Pointer to **string** |  | [optional] 
**CustomDate2** | Pointer to **string** |  | [optional] 
**CustomDate3** | Pointer to **string** |  | [optional] 
**CustomDate4** | Pointer to **string** |  | [optional] 
**IssueCustomSelect9OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_custom_select9_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_custom_select9_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueCustomSelect10OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_custom_select10_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_custom_select10_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueCustomSelect11OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_custom_select11_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_custom_select11_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueCustomSelect12OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_custom_select12_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_custom_select12_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueCustomSelect13OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_custom_select13_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_custom_select13_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueCustomSelect14OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_custom_select14_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_custom_select14_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CancelByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CancelDate** | Pointer to **string** |  | [optional] 
**FirstCancelDate** | Pointer to **string** |  | [optional] 
**LastActionPlanDueDate** | Pointer to **string** |  | [optional] 
**CancelReason** | Pointer to **string** |  | [optional] 
**DraftByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**DraftDate** | Pointer to **string** |  | [optional] 
**FirstDraftDate** | Pointer to **string** |  | [optional] 
**OpenRevisionCount** | Pointer to **int32** |  | [optional] [default to 0]
**CustomText7** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FieldData** | Pointer to **interface{}** |  | [optional] 
**ComplianceAssessmentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;compliance_assessments.id&#x60;.&lt;fk table&#x3D;&#39;compliance_assessments&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ComplianceAssessmentItemId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;compliance_assessment_items.id&#x60;.&lt;fk table&#x3D;&#39;compliance_assessment_items&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AuditSurveyId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_surveys.id&#x60;.&lt;fk table&#x3D;&#39;audit_surveys&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 

## Methods

### NewIssuesPutPreviousValues

`func NewIssuesPutPreviousValues() *IssuesPutPreviousValues`

NewIssuesPutPreviousValues instantiates a new IssuesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuesPutPreviousValuesWithDefaults

`func NewIssuesPutPreviousValuesWithDefaults() *IssuesPutPreviousValues`

NewIssuesPutPreviousValuesWithDefaults instantiates a new IssuesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssuesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssuesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssuesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IssuesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *IssuesPutPreviousValues) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IssuesPutPreviousValues) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IssuesPutPreviousValues) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IssuesPutPreviousValues) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *IssuesPutPreviousValues) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IssuesPutPreviousValues) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IssuesPutPreviousValues) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IssuesPutPreviousValues) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *IssuesPutPreviousValues) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssuesPutPreviousValues) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssuesPutPreviousValues) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssuesPutPreviousValues) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPotentialRisk

`func (o *IssuesPutPreviousValues) GetPotentialRisk() string`

GetPotentialRisk returns the PotentialRisk field if non-nil, zero value otherwise.

### GetPotentialRiskOk

`func (o *IssuesPutPreviousValues) GetPotentialRiskOk() (*string, bool)`

GetPotentialRiskOk returns a tuple with the PotentialRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPotentialRisk

`func (o *IssuesPutPreviousValues) SetPotentialRisk(v string)`

SetPotentialRisk sets PotentialRisk field to given value.

### HasPotentialRisk

`func (o *IssuesPutPreviousValues) HasPotentialRisk() bool`

HasPotentialRisk returns a boolean if a field has been set.

### GetIndividualsPresent

`func (o *IssuesPutPreviousValues) GetIndividualsPresent() string`

GetIndividualsPresent returns the IndividualsPresent field if non-nil, zero value otherwise.

### GetIndividualsPresentOk

`func (o *IssuesPutPreviousValues) GetIndividualsPresentOk() (*string, bool)`

GetIndividualsPresentOk returns a tuple with the IndividualsPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualsPresent

`func (o *IssuesPutPreviousValues) SetIndividualsPresent(v string)`

SetIndividualsPresent sets IndividualsPresent field to given value.

### HasIndividualsPresent

`func (o *IssuesPutPreviousValues) HasIndividualsPresent() bool`

HasIndividualsPresent returns a boolean if a field has been set.

### GetDeficiencyLevelId

`func (o *IssuesPutPreviousValues) GetDeficiencyLevelId() int32`

GetDeficiencyLevelId returns the DeficiencyLevelId field if non-nil, zero value otherwise.

### GetDeficiencyLevelIdOk

`func (o *IssuesPutPreviousValues) GetDeficiencyLevelIdOk() (*int32, bool)`

GetDeficiencyLevelIdOk returns a tuple with the DeficiencyLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeficiencyLevelId

`func (o *IssuesPutPreviousValues) SetDeficiencyLevelId(v int32)`

SetDeficiencyLevelId sets DeficiencyLevelId field to given value.

### HasDeficiencyLevelId

`func (o *IssuesPutPreviousValues) HasDeficiencyLevelId() bool`

HasDeficiencyLevelId returns a boolean if a field has been set.

### GetQualitativeFactors

`func (o *IssuesPutPreviousValues) GetQualitativeFactors() string`

GetQualitativeFactors returns the QualitativeFactors field if non-nil, zero value otherwise.

### GetQualitativeFactorsOk

`func (o *IssuesPutPreviousValues) GetQualitativeFactorsOk() (*string, bool)`

GetQualitativeFactorsOk returns a tuple with the QualitativeFactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualitativeFactors

`func (o *IssuesPutPreviousValues) SetQualitativeFactors(v string)`

SetQualitativeFactors sets QualitativeFactors field to given value.

### HasQualitativeFactors

`func (o *IssuesPutPreviousValues) HasQualitativeFactors() bool`

HasQualitativeFactors returns a boolean if a field has been set.

### GetIssueIdentId

`func (o *IssuesPutPreviousValues) GetIssueIdentId() int32`

GetIssueIdentId returns the IssueIdentId field if non-nil, zero value otherwise.

### GetIssueIdentIdOk

`func (o *IssuesPutPreviousValues) GetIssueIdentIdOk() (*int32, bool)`

GetIssueIdentIdOk returns a tuple with the IssueIdentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueIdentId

`func (o *IssuesPutPreviousValues) SetIssueIdentId(v int32)`

SetIssueIdentId sets IssueIdentId field to given value.

### HasIssueIdentId

`func (o *IssuesPutPreviousValues) HasIssueIdentId() bool`

HasIssueIdentId returns a boolean if a field has been set.

### GetGrossExposure

`func (o *IssuesPutPreviousValues) GetGrossExposure() string`

GetGrossExposure returns the GrossExposure field if non-nil, zero value otherwise.

### GetGrossExposureOk

`func (o *IssuesPutPreviousValues) GetGrossExposureOk() (*string, bool)`

GetGrossExposureOk returns a tuple with the GrossExposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossExposure

`func (o *IssuesPutPreviousValues) SetGrossExposure(v string)`

SetGrossExposure sets GrossExposure field to given value.

### HasGrossExposure

`func (o *IssuesPutPreviousValues) HasGrossExposure() bool`

HasGrossExposure returns a boolean if a field has been set.

### GetMagnitudeId

`func (o *IssuesPutPreviousValues) GetMagnitudeId() int32`

GetMagnitudeId returns the MagnitudeId field if non-nil, zero value otherwise.

### GetMagnitudeIdOk

`func (o *IssuesPutPreviousValues) GetMagnitudeIdOk() (*int32, bool)`

GetMagnitudeIdOk returns a tuple with the MagnitudeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagnitudeId

`func (o *IssuesPutPreviousValues) SetMagnitudeId(v int32)`

SetMagnitudeId sets MagnitudeId field to given value.

### HasMagnitudeId

`func (o *IssuesPutPreviousValues) HasMagnitudeId() bool`

HasMagnitudeId returns a boolean if a field has been set.

### GetLikelihoodId

`func (o *IssuesPutPreviousValues) GetLikelihoodId() int32`

GetLikelihoodId returns the LikelihoodId field if non-nil, zero value otherwise.

### GetLikelihoodIdOk

`func (o *IssuesPutPreviousValues) GetLikelihoodIdOk() (*int32, bool)`

GetLikelihoodIdOk returns a tuple with the LikelihoodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikelihoodId

`func (o *IssuesPutPreviousValues) SetLikelihoodId(v int32)`

SetLikelihoodId sets LikelihoodId field to given value.

### HasLikelihoodId

`func (o *IssuesPutPreviousValues) HasLikelihoodId() bool`

HasLikelihoodId returns a boolean if a field has been set.

### GetAggregationReferenceId

`func (o *IssuesPutPreviousValues) GetAggregationReferenceId() int32`

GetAggregationReferenceId returns the AggregationReferenceId field if non-nil, zero value otherwise.

### GetAggregationReferenceIdOk

`func (o *IssuesPutPreviousValues) GetAggregationReferenceIdOk() (*int32, bool)`

GetAggregationReferenceIdOk returns a tuple with the AggregationReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationReferenceId

`func (o *IssuesPutPreviousValues) SetAggregationReferenceId(v int32)`

SetAggregationReferenceId sets AggregationReferenceId field to given value.

### HasAggregationReferenceId

`func (o *IssuesPutPreviousValues) HasAggregationReferenceId() bool`

HasAggregationReferenceId returns a boolean if a field has been set.

### GetAdjustedExposure

`func (o *IssuesPutPreviousValues) GetAdjustedExposure() string`

GetAdjustedExposure returns the AdjustedExposure field if non-nil, zero value otherwise.

### GetAdjustedExposureOk

`func (o *IssuesPutPreviousValues) GetAdjustedExposureOk() (*string, bool)`

GetAdjustedExposureOk returns a tuple with the AdjustedExposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedExposure

`func (o *IssuesPutPreviousValues) SetAdjustedExposure(v string)`

SetAdjustedExposure sets AdjustedExposure field to given value.

### HasAdjustedExposure

`func (o *IssuesPutPreviousValues) HasAdjustedExposure() bool`

HasAdjustedExposure returns a boolean if a field has been set.

### GetCreatorUserId

`func (o *IssuesPutPreviousValues) GetCreatorUserId() int32`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *IssuesPutPreviousValues) GetCreatorUserIdOk() (*int32, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *IssuesPutPreviousValues) SetCreatorUserId(v int32)`

SetCreatorUserId sets CreatorUserId field to given value.

### HasCreatorUserId

`func (o *IssuesPutPreviousValues) HasCreatorUserId() bool`

HasCreatorUserId returns a boolean if a field has been set.

### GetIdentifiedDate

`func (o *IssuesPutPreviousValues) GetIdentifiedDate() string`

GetIdentifiedDate returns the IdentifiedDate field if non-nil, zero value otherwise.

### GetIdentifiedDateOk

`func (o *IssuesPutPreviousValues) GetIdentifiedDateOk() (*string, bool)`

GetIdentifiedDateOk returns a tuple with the IdentifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiedDate

`func (o *IssuesPutPreviousValues) SetIdentifiedDate(v string)`

SetIdentifiedDate sets IdentifiedDate field to given value.

### HasIdentifiedDate

`func (o *IssuesPutPreviousValues) HasIdentifiedDate() bool`

HasIdentifiedDate returns a boolean if a field has been set.

### GetDeletedAt

`func (o *IssuesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *IssuesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *IssuesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *IssuesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IssuesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssuesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssuesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IssuesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IssuesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IssuesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IssuesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IssuesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetControlsDatumId

`func (o *IssuesPutPreviousValues) GetControlsDatumId() int32`

GetControlsDatumId returns the ControlsDatumId field if non-nil, zero value otherwise.

### GetControlsDatumIdOk

`func (o *IssuesPutPreviousValues) GetControlsDatumIdOk() (*int32, bool)`

GetControlsDatumIdOk returns a tuple with the ControlsDatumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlsDatumId

`func (o *IssuesPutPreviousValues) SetControlsDatumId(v int32)`

SetControlsDatumId sets ControlsDatumId field to given value.

### HasControlsDatumId

`func (o *IssuesPutPreviousValues) HasControlsDatumId() bool`

HasControlsDatumId returns a boolean if a field has been set.

### GetTestId

`func (o *IssuesPutPreviousValues) GetTestId() int32`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *IssuesPutPreviousValues) GetTestIdOk() (*int32, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *IssuesPutPreviousValues) SetTestId(v int32)`

SetTestId sets TestId field to given value.

### HasTestId

`func (o *IssuesPutPreviousValues) HasTestId() bool`

HasTestId returns a boolean if a field has been set.

### GetTestSectionId

`func (o *IssuesPutPreviousValues) GetTestSectionId() int32`

GetTestSectionId returns the TestSectionId field if non-nil, zero value otherwise.

### GetTestSectionIdOk

`func (o *IssuesPutPreviousValues) GetTestSectionIdOk() (*int32, bool)`

GetTestSectionIdOk returns a tuple with the TestSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSectionId

`func (o *IssuesPutPreviousValues) SetTestSectionId(v int32)`

SetTestSectionId sets TestSectionId field to given value.

### HasTestSectionId

`func (o *IssuesPutPreviousValues) HasTestSectionId() bool`

HasTestSectionId returns a boolean if a field has been set.

### GetNotes

`func (o *IssuesPutPreviousValues) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *IssuesPutPreviousValues) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *IssuesPutPreviousValues) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *IssuesPutPreviousValues) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetReferenceIssueId

`func (o *IssuesPutPreviousValues) GetReferenceIssueId() int32`

GetReferenceIssueId returns the ReferenceIssueId field if non-nil, zero value otherwise.

### GetReferenceIssueIdOk

`func (o *IssuesPutPreviousValues) GetReferenceIssueIdOk() (*int32, bool)`

GetReferenceIssueIdOk returns a tuple with the ReferenceIssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceIssueId

`func (o *IssuesPutPreviousValues) SetReferenceIssueId(v int32)`

SetReferenceIssueId sets ReferenceIssueId field to given value.

### HasReferenceIssueId

`func (o *IssuesPutPreviousValues) HasReferenceIssueId() bool`

HasReferenceIssueId returns a boolean if a field has been set.

### GetArchiveId

`func (o *IssuesPutPreviousValues) GetArchiveId() int32`

GetArchiveId returns the ArchiveId field if non-nil, zero value otherwise.

### GetArchiveIdOk

`func (o *IssuesPutPreviousValues) GetArchiveIdOk() (*int32, bool)`

GetArchiveIdOk returns a tuple with the ArchiveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveId

`func (o *IssuesPutPreviousValues) SetArchiveId(v int32)`

SetArchiveId sets ArchiveId field to given value.

### HasArchiveId

`func (o *IssuesPutPreviousValues) HasArchiveId() bool`

HasArchiveId returns a boolean if a field has been set.

### GetTesterUserId

`func (o *IssuesPutPreviousValues) GetTesterUserId() int32`

GetTesterUserId returns the TesterUserId field if non-nil, zero value otherwise.

### GetTesterUserIdOk

`func (o *IssuesPutPreviousValues) GetTesterUserIdOk() (*int32, bool)`

GetTesterUserIdOk returns a tuple with the TesterUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesterUserId

`func (o *IssuesPutPreviousValues) SetTesterUserId(v int32)`

SetTesterUserId sets TesterUserId field to given value.

### HasTesterUserId

`func (o *IssuesPutPreviousValues) HasTesterUserId() bool`

HasTesterUserId returns a boolean if a field has been set.

### GetReviewerUserId

`func (o *IssuesPutPreviousValues) GetReviewerUserId() int32`

GetReviewerUserId returns the ReviewerUserId field if non-nil, zero value otherwise.

### GetReviewerUserIdOk

`func (o *IssuesPutPreviousValues) GetReviewerUserIdOk() (*int32, bool)`

GetReviewerUserIdOk returns a tuple with the ReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerUserId

`func (o *IssuesPutPreviousValues) SetReviewerUserId(v int32)`

SetReviewerUserId sets ReviewerUserId field to given value.

### HasReviewerUserId

`func (o *IssuesPutPreviousValues) HasReviewerUserId() bool`

HasReviewerUserId returns a boolean if a field has been set.

### GetOpsAuditId

`func (o *IssuesPutPreviousValues) GetOpsAuditId() int32`

GetOpsAuditId returns the OpsAuditId field if non-nil, zero value otherwise.

### GetOpsAuditIdOk

`func (o *IssuesPutPreviousValues) GetOpsAuditIdOk() (*int32, bool)`

GetOpsAuditIdOk returns a tuple with the OpsAuditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditId

`func (o *IssuesPutPreviousValues) SetOpsAuditId(v int32)`

SetOpsAuditId sets OpsAuditId field to given value.

### HasOpsAuditId

`func (o *IssuesPutPreviousValues) HasOpsAuditId() bool`

HasOpsAuditId returns a boolean if a field has been set.

### GetOpsAuditItemId

`func (o *IssuesPutPreviousValues) GetOpsAuditItemId() int32`

GetOpsAuditItemId returns the OpsAuditItemId field if non-nil, zero value otherwise.

### GetOpsAuditItemIdOk

`func (o *IssuesPutPreviousValues) GetOpsAuditItemIdOk() (*int32, bool)`

GetOpsAuditItemIdOk returns a tuple with the OpsAuditItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditItemId

`func (o *IssuesPutPreviousValues) SetOpsAuditItemId(v int32)`

SetOpsAuditItemId sets OpsAuditItemId field to given value.

### HasOpsAuditItemId

`func (o *IssuesPutPreviousValues) HasOpsAuditItemId() bool`

HasOpsAuditItemId returns a boolean if a field has been set.

### GetType

`func (o *IssuesPutPreviousValues) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IssuesPutPreviousValues) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IssuesPutPreviousValues) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IssuesPutPreviousValues) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInternalControlComponentId

`func (o *IssuesPutPreviousValues) GetInternalControlComponentId() int32`

GetInternalControlComponentId returns the InternalControlComponentId field if non-nil, zero value otherwise.

### GetInternalControlComponentIdOk

`func (o *IssuesPutPreviousValues) GetInternalControlComponentIdOk() (*int32, bool)`

GetInternalControlComponentIdOk returns a tuple with the InternalControlComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalControlComponentId

`func (o *IssuesPutPreviousValues) SetInternalControlComponentId(v int32)`

SetInternalControlComponentId sets InternalControlComponentId field to given value.

### HasInternalControlComponentId

`func (o *IssuesPutPreviousValues) HasInternalControlComponentId() bool`

HasInternalControlComponentId returns a boolean if a field has been set.

### GetExternalPlannedConfirmDate

`func (o *IssuesPutPreviousValues) GetExternalPlannedConfirmDate() string`

GetExternalPlannedConfirmDate returns the ExternalPlannedConfirmDate field if non-nil, zero value otherwise.

### GetExternalPlannedConfirmDateOk

`func (o *IssuesPutPreviousValues) GetExternalPlannedConfirmDateOk() (*string, bool)`

GetExternalPlannedConfirmDateOk returns a tuple with the ExternalPlannedConfirmDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPlannedConfirmDate

`func (o *IssuesPutPreviousValues) SetExternalPlannedConfirmDate(v string)`

SetExternalPlannedConfirmDate sets ExternalPlannedConfirmDate field to given value.

### HasExternalPlannedConfirmDate

`func (o *IssuesPutPreviousValues) HasExternalPlannedConfirmDate() bool`

HasExternalPlannedConfirmDate returns a boolean if a field has been set.

### GetDisclosure

`func (o *IssuesPutPreviousValues) GetDisclosure() string`

GetDisclosure returns the Disclosure field if non-nil, zero value otherwise.

### GetDisclosureOk

`func (o *IssuesPutPreviousValues) GetDisclosureOk() (*string, bool)`

GetDisclosureOk returns a tuple with the Disclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosure

`func (o *IssuesPutPreviousValues) SetDisclosure(v string)`

SetDisclosure sets Disclosure field to given value.

### HasDisclosure

`func (o *IssuesPutPreviousValues) HasDisclosure() bool`

HasDisclosure returns a boolean if a field has been set.

### GetExternalConfirmDate

`func (o *IssuesPutPreviousValues) GetExternalConfirmDate() string`

GetExternalConfirmDate returns the ExternalConfirmDate field if non-nil, zero value otherwise.

### GetExternalConfirmDateOk

`func (o *IssuesPutPreviousValues) GetExternalConfirmDateOk() (*string, bool)`

GetExternalConfirmDateOk returns a tuple with the ExternalConfirmDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConfirmDate

`func (o *IssuesPutPreviousValues) SetExternalConfirmDate(v string)`

SetExternalConfirmDate sets ExternalConfirmDate field to given value.

### HasExternalConfirmDate

`func (o *IssuesPutPreviousValues) HasExternalConfirmDate() bool`

HasExternalConfirmDate returns a boolean if a field has been set.

### GetPendingRemediationDate

`func (o *IssuesPutPreviousValues) GetPendingRemediationDate() string`

GetPendingRemediationDate returns the PendingRemediationDate field if non-nil, zero value otherwise.

### GetPendingRemediationDateOk

`func (o *IssuesPutPreviousValues) GetPendingRemediationDateOk() (*string, bool)`

GetPendingRemediationDateOk returns a tuple with the PendingRemediationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingRemediationDate

`func (o *IssuesPutPreviousValues) SetPendingRemediationDate(v string)`

SetPendingRemediationDate sets PendingRemediationDate field to given value.

### HasPendingRemediationDate

`func (o *IssuesPutPreviousValues) HasPendingRemediationDate() bool`

HasPendingRemediationDate returns a boolean if a field has been set.

### GetClosedDate

`func (o *IssuesPutPreviousValues) GetClosedDate() string`

GetClosedDate returns the ClosedDate field if non-nil, zero value otherwise.

### GetClosedDateOk

`func (o *IssuesPutPreviousValues) GetClosedDateOk() (*string, bool)`

GetClosedDateOk returns a tuple with the ClosedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedDate

`func (o *IssuesPutPreviousValues) SetClosedDate(v string)`

SetClosedDate sets ClosedDate field to given value.

### HasClosedDate

`func (o *IssuesPutPreviousValues) HasClosedDate() bool`

HasClosedDate returns a boolean if a field has been set.

### GetInactiveDate

`func (o *IssuesPutPreviousValues) GetInactiveDate() string`

GetInactiveDate returns the InactiveDate field if non-nil, zero value otherwise.

### GetInactiveDateOk

`func (o *IssuesPutPreviousValues) GetInactiveDateOk() (*string, bool)`

GetInactiveDateOk returns a tuple with the InactiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveDate

`func (o *IssuesPutPreviousValues) SetInactiveDate(v string)`

SetInactiveDate sets InactiveDate field to given value.

### HasInactiveDate

`func (o *IssuesPutPreviousValues) HasInactiveDate() bool`

HasInactiveDate returns a boolean if a field has been set.

### GetReopenDate

`func (o *IssuesPutPreviousValues) GetReopenDate() string`

GetReopenDate returns the ReopenDate field if non-nil, zero value otherwise.

### GetReopenDateOk

`func (o *IssuesPutPreviousValues) GetReopenDateOk() (*string, bool)`

GetReopenDateOk returns a tuple with the ReopenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReopenDate

`func (o *IssuesPutPreviousValues) SetReopenDate(v string)`

SetReopenDate sets ReopenDate field to given value.

### HasReopenDate

`func (o *IssuesPutPreviousValues) HasReopenDate() bool`

HasReopenDate returns a boolean if a field has been set.

### GetOpenDate

`func (o *IssuesPutPreviousValues) GetOpenDate() string`

GetOpenDate returns the OpenDate field if non-nil, zero value otherwise.

### GetOpenDateOk

`func (o *IssuesPutPreviousValues) GetOpenDateOk() (*string, bool)`

GetOpenDateOk returns a tuple with the OpenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenDate

`func (o *IssuesPutPreviousValues) SetOpenDate(v string)`

SetOpenDate sets OpenDate field to given value.

### HasOpenDate

`func (o *IssuesPutPreviousValues) HasOpenDate() bool`

HasOpenDate returns a boolean if a field has been set.

### GetAmendDate

`func (o *IssuesPutPreviousValues) GetAmendDate() string`

GetAmendDate returns the AmendDate field if non-nil, zero value otherwise.

### GetAmendDateOk

`func (o *IssuesPutPreviousValues) GetAmendDateOk() (*string, bool)`

GetAmendDateOk returns a tuple with the AmendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendDate

`func (o *IssuesPutPreviousValues) SetAmendDate(v string)`

SetAmendDate sets AmendDate field to given value.

### HasAmendDate

`func (o *IssuesPutPreviousValues) HasAmendDate() bool`

HasAmendDate returns a boolean if a field has been set.

### GetRemediatedDate

`func (o *IssuesPutPreviousValues) GetRemediatedDate() string`

GetRemediatedDate returns the RemediatedDate field if non-nil, zero value otherwise.

### GetRemediatedDateOk

`func (o *IssuesPutPreviousValues) GetRemediatedDateOk() (*string, bool)`

GetRemediatedDateOk returns a tuple with the RemediatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediatedDate

`func (o *IssuesPutPreviousValues) SetRemediatedDate(v string)`

SetRemediatedDate sets RemediatedDate field to given value.

### HasRemediatedDate

`func (o *IssuesPutPreviousValues) HasRemediatedDate() bool`

HasRemediatedDate returns a boolean if a field has been set.

### GetOpenByUserId

`func (o *IssuesPutPreviousValues) GetOpenByUserId() int32`

GetOpenByUserId returns the OpenByUserId field if non-nil, zero value otherwise.

### GetOpenByUserIdOk

`func (o *IssuesPutPreviousValues) GetOpenByUserIdOk() (*int32, bool)`

GetOpenByUserIdOk returns a tuple with the OpenByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenByUserId

`func (o *IssuesPutPreviousValues) SetOpenByUserId(v int32)`

SetOpenByUserId sets OpenByUserId field to given value.

### HasOpenByUserId

`func (o *IssuesPutPreviousValues) HasOpenByUserId() bool`

HasOpenByUserId returns a boolean if a field has been set.

### GetReopenByUserId

`func (o *IssuesPutPreviousValues) GetReopenByUserId() int32`

GetReopenByUserId returns the ReopenByUserId field if non-nil, zero value otherwise.

### GetReopenByUserIdOk

`func (o *IssuesPutPreviousValues) GetReopenByUserIdOk() (*int32, bool)`

GetReopenByUserIdOk returns a tuple with the ReopenByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReopenByUserId

`func (o *IssuesPutPreviousValues) SetReopenByUserId(v int32)`

SetReopenByUserId sets ReopenByUserId field to given value.

### HasReopenByUserId

`func (o *IssuesPutPreviousValues) HasReopenByUserId() bool`

HasReopenByUserId returns a boolean if a field has been set.

### GetPendingRemediationByUserId

`func (o *IssuesPutPreviousValues) GetPendingRemediationByUserId() int32`

GetPendingRemediationByUserId returns the PendingRemediationByUserId field if non-nil, zero value otherwise.

### GetPendingRemediationByUserIdOk

`func (o *IssuesPutPreviousValues) GetPendingRemediationByUserIdOk() (*int32, bool)`

GetPendingRemediationByUserIdOk returns a tuple with the PendingRemediationByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingRemediationByUserId

`func (o *IssuesPutPreviousValues) SetPendingRemediationByUserId(v int32)`

SetPendingRemediationByUserId sets PendingRemediationByUserId field to given value.

### HasPendingRemediationByUserId

`func (o *IssuesPutPreviousValues) HasPendingRemediationByUserId() bool`

HasPendingRemediationByUserId returns a boolean if a field has been set.

### GetRemediatedByUserId

`func (o *IssuesPutPreviousValues) GetRemediatedByUserId() int32`

GetRemediatedByUserId returns the RemediatedByUserId field if non-nil, zero value otherwise.

### GetRemediatedByUserIdOk

`func (o *IssuesPutPreviousValues) GetRemediatedByUserIdOk() (*int32, bool)`

GetRemediatedByUserIdOk returns a tuple with the RemediatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediatedByUserId

`func (o *IssuesPutPreviousValues) SetRemediatedByUserId(v int32)`

SetRemediatedByUserId sets RemediatedByUserId field to given value.

### HasRemediatedByUserId

`func (o *IssuesPutPreviousValues) HasRemediatedByUserId() bool`

HasRemediatedByUserId returns a boolean if a field has been set.

### GetClosedByUserId

`func (o *IssuesPutPreviousValues) GetClosedByUserId() int32`

GetClosedByUserId returns the ClosedByUserId field if non-nil, zero value otherwise.

### GetClosedByUserIdOk

`func (o *IssuesPutPreviousValues) GetClosedByUserIdOk() (*int32, bool)`

GetClosedByUserIdOk returns a tuple with the ClosedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedByUserId

`func (o *IssuesPutPreviousValues) SetClosedByUserId(v int32)`

SetClosedByUserId sets ClosedByUserId field to given value.

### HasClosedByUserId

`func (o *IssuesPutPreviousValues) HasClosedByUserId() bool`

HasClosedByUserId returns a boolean if a field has been set.

### GetAmendByUserId

`func (o *IssuesPutPreviousValues) GetAmendByUserId() int32`

GetAmendByUserId returns the AmendByUserId field if non-nil, zero value otherwise.

### GetAmendByUserIdOk

`func (o *IssuesPutPreviousValues) GetAmendByUserIdOk() (*int32, bool)`

GetAmendByUserIdOk returns a tuple with the AmendByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendByUserId

`func (o *IssuesPutPreviousValues) SetAmendByUserId(v int32)`

SetAmendByUserId sets AmendByUserId field to given value.

### HasAmendByUserId

`func (o *IssuesPutPreviousValues) HasAmendByUserId() bool`

HasAmendByUserId returns a boolean if a field has been set.

### GetFlattened

`func (o *IssuesPutPreviousValues) GetFlattened() interface{}`

GetFlattened returns the Flattened field if non-nil, zero value otherwise.

### GetFlattenedOk

`func (o *IssuesPutPreviousValues) GetFlattenedOk() (*interface{}, bool)`

GetFlattenedOk returns a tuple with the Flattened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlattened

`func (o *IssuesPutPreviousValues) SetFlattened(v interface{})`

SetFlattened sets Flattened field to given value.

### HasFlattened

`func (o *IssuesPutPreviousValues) HasFlattened() bool`

HasFlattened returns a boolean if a field has been set.

### SetFlattenedNil

`func (o *IssuesPutPreviousValues) SetFlattenedNil(b bool)`

 SetFlattenedNil sets the value for Flattened to be an explicit nil

### UnsetFlattened
`func (o *IssuesPutPreviousValues) UnsetFlattened()`

UnsetFlattened ensures that no value is present for Flattened, not even an explicit nil
### GetIssueRatingId

`func (o *IssuesPutPreviousValues) GetIssueRatingId() int32`

GetIssueRatingId returns the IssueRatingId field if non-nil, zero value otherwise.

### GetIssueRatingIdOk

`func (o *IssuesPutPreviousValues) GetIssueRatingIdOk() (*int32, bool)`

GetIssueRatingIdOk returns a tuple with the IssueRatingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueRatingId

`func (o *IssuesPutPreviousValues) SetIssueRatingId(v int32)`

SetIssueRatingId sets IssueRatingId field to given value.

### HasIssueRatingId

`func (o *IssuesPutPreviousValues) HasIssueRatingId() bool`

HasIssueRatingId returns a boolean if a field has been set.

### GetSoxImpactId

`func (o *IssuesPutPreviousValues) GetSoxImpactId() int32`

GetSoxImpactId returns the SoxImpactId field if non-nil, zero value otherwise.

### GetSoxImpactIdOk

`func (o *IssuesPutPreviousValues) GetSoxImpactIdOk() (*int32, bool)`

GetSoxImpactIdOk returns a tuple with the SoxImpactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoxImpactId

`func (o *IssuesPutPreviousValues) SetSoxImpactId(v int32)`

SetSoxImpactId sets SoxImpactId field to given value.

### HasSoxImpactId

`func (o *IssuesPutPreviousValues) HasSoxImpactId() bool`

HasSoxImpactId returns a boolean if a field has been set.

### GetIssueCustomSelect1OptionId

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect1OptionId() int32`

GetIssueCustomSelect1OptionId returns the IssueCustomSelect1OptionId field if non-nil, zero value otherwise.

### GetIssueCustomSelect1OptionIdOk

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect1OptionIdOk() (*int32, bool)`

GetIssueCustomSelect1OptionIdOk returns a tuple with the IssueCustomSelect1OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCustomSelect1OptionId

`func (o *IssuesPutPreviousValues) SetIssueCustomSelect1OptionId(v int32)`

SetIssueCustomSelect1OptionId sets IssueCustomSelect1OptionId field to given value.

### HasIssueCustomSelect1OptionId

`func (o *IssuesPutPreviousValues) HasIssueCustomSelect1OptionId() bool`

HasIssueCustomSelect1OptionId returns a boolean if a field has been set.

### GetIssueCustomSelect2OptionId

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect2OptionId() int32`

GetIssueCustomSelect2OptionId returns the IssueCustomSelect2OptionId field if non-nil, zero value otherwise.

### GetIssueCustomSelect2OptionIdOk

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect2OptionIdOk() (*int32, bool)`

GetIssueCustomSelect2OptionIdOk returns a tuple with the IssueCustomSelect2OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCustomSelect2OptionId

`func (o *IssuesPutPreviousValues) SetIssueCustomSelect2OptionId(v int32)`

SetIssueCustomSelect2OptionId sets IssueCustomSelect2OptionId field to given value.

### HasIssueCustomSelect2OptionId

`func (o *IssuesPutPreviousValues) HasIssueCustomSelect2OptionId() bool`

HasIssueCustomSelect2OptionId returns a boolean if a field has been set.

### GetVicePresidentUserId

`func (o *IssuesPutPreviousValues) GetVicePresidentUserId() int32`

GetVicePresidentUserId returns the VicePresidentUserId field if non-nil, zero value otherwise.

### GetVicePresidentUserIdOk

`func (o *IssuesPutPreviousValues) GetVicePresidentUserIdOk() (*int32, bool)`

GetVicePresidentUserIdOk returns a tuple with the VicePresidentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVicePresidentUserId

`func (o *IssuesPutPreviousValues) SetVicePresidentUserId(v int32)`

SetVicePresidentUserId sets VicePresidentUserId field to given value.

### HasVicePresidentUserId

`func (o *IssuesPutPreviousValues) HasVicePresidentUserId() bool`

HasVicePresidentUserId returns a boolean if a field has been set.

### GetIsFlattened

`func (o *IssuesPutPreviousValues) GetIsFlattened() bool`

GetIsFlattened returns the IsFlattened field if non-nil, zero value otherwise.

### GetIsFlattenedOk

`func (o *IssuesPutPreviousValues) GetIsFlattenedOk() (*bool, bool)`

GetIsFlattenedOk returns a tuple with the IsFlattened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlattened

`func (o *IssuesPutPreviousValues) SetIsFlattened(v bool)`

SetIsFlattened sets IsFlattened field to given value.

### HasIsFlattened

`func (o *IssuesPutPreviousValues) HasIsFlattened() bool`

HasIsFlattened returns a boolean if a field has been set.

### GetRemediationOwners

`func (o *IssuesPutPreviousValues) GetRemediationOwners() interface{}`

GetRemediationOwners returns the RemediationOwners field if non-nil, zero value otherwise.

### GetRemediationOwnersOk

`func (o *IssuesPutPreviousValues) GetRemediationOwnersOk() (*interface{}, bool)`

GetRemediationOwnersOk returns a tuple with the RemediationOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationOwners

`func (o *IssuesPutPreviousValues) SetRemediationOwners(v interface{})`

SetRemediationOwners sets RemediationOwners field to given value.

### HasRemediationOwners

`func (o *IssuesPutPreviousValues) HasRemediationOwners() bool`

HasRemediationOwners returns a boolean if a field has been set.

### SetRemediationOwnersNil

`func (o *IssuesPutPreviousValues) SetRemediationOwnersNil(b bool)`

 SetRemediationOwnersNil sets the value for RemediationOwners to be an explicit nil

### UnsetRemediationOwners
`func (o *IssuesPutPreviousValues) UnsetRemediationOwners()`

UnsetRemediationOwners ensures that no value is present for RemediationOwners, not even an explicit nil
### GetReferenceMeta

`func (o *IssuesPutPreviousValues) GetReferenceMeta() interface{}`

GetReferenceMeta returns the ReferenceMeta field if non-nil, zero value otherwise.

### GetReferenceMetaOk

`func (o *IssuesPutPreviousValues) GetReferenceMetaOk() (*interface{}, bool)`

GetReferenceMetaOk returns a tuple with the ReferenceMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMeta

`func (o *IssuesPutPreviousValues) SetReferenceMeta(v interface{})`

SetReferenceMeta sets ReferenceMeta field to given value.

### HasReferenceMeta

`func (o *IssuesPutPreviousValues) HasReferenceMeta() bool`

HasReferenceMeta returns a boolean if a field has been set.

### SetReferenceMetaNil

`func (o *IssuesPutPreviousValues) SetReferenceMetaNil(b bool)`

 SetReferenceMetaNil sets the value for ReferenceMeta to be an explicit nil

### UnsetReferenceMeta
`func (o *IssuesPutPreviousValues) UnsetReferenceMeta()`

UnsetReferenceMeta ensures that no value is present for ReferenceMeta, not even an explicit nil
### GetTotals

`func (o *IssuesPutPreviousValues) GetTotals() interface{}`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *IssuesPutPreviousValues) GetTotalsOk() (*interface{}, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *IssuesPutPreviousValues) SetTotals(v interface{})`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *IssuesPutPreviousValues) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### SetTotalsNil

`func (o *IssuesPutPreviousValues) SetTotalsNil(b bool)`

 SetTotalsNil sets the value for Totals to be an explicit nil

### UnsetTotals
`func (o *IssuesPutPreviousValues) UnsetTotals()`

UnsetTotals ensures that no value is present for Totals, not even an explicit nil
### GetIsArchived

`func (o *IssuesPutPreviousValues) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *IssuesPutPreviousValues) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *IssuesPutPreviousValues) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *IssuesPutPreviousValues) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetIssueCustomSelect3OptionId

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect3OptionId() int32`

GetIssueCustomSelect3OptionId returns the IssueCustomSelect3OptionId field if non-nil, zero value otherwise.

### GetIssueCustomSelect3OptionIdOk

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect3OptionIdOk() (*int32, bool)`

GetIssueCustomSelect3OptionIdOk returns a tuple with the IssueCustomSelect3OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCustomSelect3OptionId

`func (o *IssuesPutPreviousValues) SetIssueCustomSelect3OptionId(v int32)`

SetIssueCustomSelect3OptionId sets IssueCustomSelect3OptionId field to given value.

### HasIssueCustomSelect3OptionId

`func (o *IssuesPutPreviousValues) HasIssueCustomSelect3OptionId() bool`

HasIssueCustomSelect3OptionId returns a boolean if a field has been set.

### GetAdditionalAuditorUserId

`func (o *IssuesPutPreviousValues) GetAdditionalAuditorUserId() int32`

GetAdditionalAuditorUserId returns the AdditionalAuditorUserId field if non-nil, zero value otherwise.

### GetAdditionalAuditorUserIdOk

`func (o *IssuesPutPreviousValues) GetAdditionalAuditorUserIdOk() (*int32, bool)`

GetAdditionalAuditorUserIdOk returns a tuple with the AdditionalAuditorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAuditorUserId

`func (o *IssuesPutPreviousValues) SetAdditionalAuditorUserId(v int32)`

SetAdditionalAuditorUserId sets AdditionalAuditorUserId field to given value.

### HasAdditionalAuditorUserId

`func (o *IssuesPutPreviousValues) HasAdditionalAuditorUserId() bool`

HasAdditionalAuditorUserId returns a boolean if a field has been set.

### GetInternalAuditValidationDate

`func (o *IssuesPutPreviousValues) GetInternalAuditValidationDate() string`

GetInternalAuditValidationDate returns the InternalAuditValidationDate field if non-nil, zero value otherwise.

### GetInternalAuditValidationDateOk

`func (o *IssuesPutPreviousValues) GetInternalAuditValidationDateOk() (*string, bool)`

GetInternalAuditValidationDateOk returns a tuple with the InternalAuditValidationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAuditValidationDate

`func (o *IssuesPutPreviousValues) SetInternalAuditValidationDate(v string)`

SetInternalAuditValidationDate sets InternalAuditValidationDate field to given value.

### HasInternalAuditValidationDate

`func (o *IssuesPutPreviousValues) HasInternalAuditValidationDate() bool`

HasInternalAuditValidationDate returns a boolean if a field has been set.

### GetCustomText1

`func (o *IssuesPutPreviousValues) GetCustomText1() string`

GetCustomText1 returns the CustomText1 field if non-nil, zero value otherwise.

### GetCustomText1Ok

`func (o *IssuesPutPreviousValues) GetCustomText1Ok() (*string, bool)`

GetCustomText1Ok returns a tuple with the CustomText1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText1

`func (o *IssuesPutPreviousValues) SetCustomText1(v string)`

SetCustomText1 sets CustomText1 field to given value.

### HasCustomText1

`func (o *IssuesPutPreviousValues) HasCustomText1() bool`

HasCustomText1 returns a boolean if a field has been set.

### GetCustomText2

`func (o *IssuesPutPreviousValues) GetCustomText2() string`

GetCustomText2 returns the CustomText2 field if non-nil, zero value otherwise.

### GetCustomText2Ok

`func (o *IssuesPutPreviousValues) GetCustomText2Ok() (*string, bool)`

GetCustomText2Ok returns a tuple with the CustomText2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText2

`func (o *IssuesPutPreviousValues) SetCustomText2(v string)`

SetCustomText2 sets CustomText2 field to given value.

### HasCustomText2

`func (o *IssuesPutPreviousValues) HasCustomText2() bool`

HasCustomText2 returns a boolean if a field has been set.

### GetCustomText3

`func (o *IssuesPutPreviousValues) GetCustomText3() string`

GetCustomText3 returns the CustomText3 field if non-nil, zero value otherwise.

### GetCustomText3Ok

`func (o *IssuesPutPreviousValues) GetCustomText3Ok() (*string, bool)`

GetCustomText3Ok returns a tuple with the CustomText3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText3

`func (o *IssuesPutPreviousValues) SetCustomText3(v string)`

SetCustomText3 sets CustomText3 field to given value.

### HasCustomText3

`func (o *IssuesPutPreviousValues) HasCustomText3() bool`

HasCustomText3 returns a boolean if a field has been set.

### GetCustomText4

`func (o *IssuesPutPreviousValues) GetCustomText4() string`

GetCustomText4 returns the CustomText4 field if non-nil, zero value otherwise.

### GetCustomText4Ok

`func (o *IssuesPutPreviousValues) GetCustomText4Ok() (*string, bool)`

GetCustomText4Ok returns a tuple with the CustomText4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText4

`func (o *IssuesPutPreviousValues) SetCustomText4(v string)`

SetCustomText4 sets CustomText4 field to given value.

### HasCustomText4

`func (o *IssuesPutPreviousValues) HasCustomText4() bool`

HasCustomText4 returns a boolean if a field has been set.

### GetCustomText5

`func (o *IssuesPutPreviousValues) GetCustomText5() string`

GetCustomText5 returns the CustomText5 field if non-nil, zero value otherwise.

### GetCustomText5Ok

`func (o *IssuesPutPreviousValues) GetCustomText5Ok() (*string, bool)`

GetCustomText5Ok returns a tuple with the CustomText5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText5

`func (o *IssuesPutPreviousValues) SetCustomText5(v string)`

SetCustomText5 sets CustomText5 field to given value.

### HasCustomText5

`func (o *IssuesPutPreviousValues) HasCustomText5() bool`

HasCustomText5 returns a boolean if a field has been set.

### GetCustomText6

`func (o *IssuesPutPreviousValues) GetCustomText6() string`

GetCustomText6 returns the CustomText6 field if non-nil, zero value otherwise.

### GetCustomText6Ok

`func (o *IssuesPutPreviousValues) GetCustomText6Ok() (*string, bool)`

GetCustomText6Ok returns a tuple with the CustomText6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText6

`func (o *IssuesPutPreviousValues) SetCustomText6(v string)`

SetCustomText6 sets CustomText6 field to given value.

### HasCustomText6

`func (o *IssuesPutPreviousValues) HasCustomText6() bool`

HasCustomText6 returns a boolean if a field has been set.

### GetIssueCustomSelect4OptionId

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect4OptionId() int32`

GetIssueCustomSelect4OptionId returns the IssueCustomSelect4OptionId field if non-nil, zero value otherwise.

### GetIssueCustomSelect4OptionIdOk

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect4OptionIdOk() (*int32, bool)`

GetIssueCustomSelect4OptionIdOk returns a tuple with the IssueCustomSelect4OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCustomSelect4OptionId

`func (o *IssuesPutPreviousValues) SetIssueCustomSelect4OptionId(v int32)`

SetIssueCustomSelect4OptionId sets IssueCustomSelect4OptionId field to given value.

### HasIssueCustomSelect4OptionId

`func (o *IssuesPutPreviousValues) HasIssueCustomSelect4OptionId() bool`

HasIssueCustomSelect4OptionId returns a boolean if a field has been set.

### GetIssueCustomSelect5OptionId

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect5OptionId() int32`

GetIssueCustomSelect5OptionId returns the IssueCustomSelect5OptionId field if non-nil, zero value otherwise.

### GetIssueCustomSelect5OptionIdOk

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect5OptionIdOk() (*int32, bool)`

GetIssueCustomSelect5OptionIdOk returns a tuple with the IssueCustomSelect5OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCustomSelect5OptionId

`func (o *IssuesPutPreviousValues) SetIssueCustomSelect5OptionId(v int32)`

SetIssueCustomSelect5OptionId sets IssueCustomSelect5OptionId field to given value.

### HasIssueCustomSelect5OptionId

`func (o *IssuesPutPreviousValues) HasIssueCustomSelect5OptionId() bool`

HasIssueCustomSelect5OptionId returns a boolean if a field has been set.

### GetIssueCustomSelect6OptionId

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect6OptionId() int32`

GetIssueCustomSelect6OptionId returns the IssueCustomSelect6OptionId field if non-nil, zero value otherwise.

### GetIssueCustomSelect6OptionIdOk

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect6OptionIdOk() (*int32, bool)`

GetIssueCustomSelect6OptionIdOk returns a tuple with the IssueCustomSelect6OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCustomSelect6OptionId

`func (o *IssuesPutPreviousValues) SetIssueCustomSelect6OptionId(v int32)`

SetIssueCustomSelect6OptionId sets IssueCustomSelect6OptionId field to given value.

### HasIssueCustomSelect6OptionId

`func (o *IssuesPutPreviousValues) HasIssueCustomSelect6OptionId() bool`

HasIssueCustomSelect6OptionId returns a boolean if a field has been set.

### GetIssueCustomSelect7OptionId

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect7OptionId() int32`

GetIssueCustomSelect7OptionId returns the IssueCustomSelect7OptionId field if non-nil, zero value otherwise.

### GetIssueCustomSelect7OptionIdOk

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect7OptionIdOk() (*int32, bool)`

GetIssueCustomSelect7OptionIdOk returns a tuple with the IssueCustomSelect7OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCustomSelect7OptionId

`func (o *IssuesPutPreviousValues) SetIssueCustomSelect7OptionId(v int32)`

SetIssueCustomSelect7OptionId sets IssueCustomSelect7OptionId field to given value.

### HasIssueCustomSelect7OptionId

`func (o *IssuesPutPreviousValues) HasIssueCustomSelect7OptionId() bool`

HasIssueCustomSelect7OptionId returns a boolean if a field has been set.

### GetIssueCustomSelect8OptionId

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect8OptionId() int32`

GetIssueCustomSelect8OptionId returns the IssueCustomSelect8OptionId field if non-nil, zero value otherwise.

### GetIssueCustomSelect8OptionIdOk

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect8OptionIdOk() (*int32, bool)`

GetIssueCustomSelect8OptionIdOk returns a tuple with the IssueCustomSelect8OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCustomSelect8OptionId

`func (o *IssuesPutPreviousValues) SetIssueCustomSelect8OptionId(v int32)`

SetIssueCustomSelect8OptionId sets IssueCustomSelect8OptionId field to given value.

### HasIssueCustomSelect8OptionId

`func (o *IssuesPutPreviousValues) HasIssueCustomSelect8OptionId() bool`

HasIssueCustomSelect8OptionId returns a boolean if a field has been set.

### GetIssueCategoryId

`func (o *IssuesPutPreviousValues) GetIssueCategoryId() int32`

GetIssueCategoryId returns the IssueCategoryId field if non-nil, zero value otherwise.

### GetIssueCategoryIdOk

`func (o *IssuesPutPreviousValues) GetIssueCategoryIdOk() (*int32, bool)`

GetIssueCategoryIdOk returns a tuple with the IssueCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCategoryId

`func (o *IssuesPutPreviousValues) SetIssueCategoryId(v int32)`

SetIssueCategoryId sets IssueCategoryId field to given value.

### HasIssueCategoryId

`func (o *IssuesPutPreviousValues) HasIssueCategoryId() bool`

HasIssueCategoryId returns a boolean if a field has been set.

### GetIssuesourceableId

`func (o *IssuesPutPreviousValues) GetIssuesourceableId() int32`

GetIssuesourceableId returns the IssuesourceableId field if non-nil, zero value otherwise.

### GetIssuesourceableIdOk

`func (o *IssuesPutPreviousValues) GetIssuesourceableIdOk() (*int32, bool)`

GetIssuesourceableIdOk returns a tuple with the IssuesourceableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesourceableId

`func (o *IssuesPutPreviousValues) SetIssuesourceableId(v int32)`

SetIssuesourceableId sets IssuesourceableId field to given value.

### HasIssuesourceableId

`func (o *IssuesPutPreviousValues) HasIssuesourceableId() bool`

HasIssuesourceableId returns a boolean if a field has been set.

### GetIssuesourceableType

`func (o *IssuesPutPreviousValues) GetIssuesourceableType() string`

GetIssuesourceableType returns the IssuesourceableType field if non-nil, zero value otherwise.

### GetIssuesourceableTypeOk

`func (o *IssuesPutPreviousValues) GetIssuesourceableTypeOk() (*string, bool)`

GetIssuesourceableTypeOk returns a tuple with the IssuesourceableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesourceableType

`func (o *IssuesPutPreviousValues) SetIssuesourceableType(v string)`

SetIssuesourceableType sets IssuesourceableType field to given value.

### HasIssuesourceableType

`func (o *IssuesPutPreviousValues) HasIssuesourceableType() bool`

HasIssuesourceableType returns a boolean if a field has been set.

### GetFinanceProcessOwners

`func (o *IssuesPutPreviousValues) GetFinanceProcessOwners() interface{}`

GetFinanceProcessOwners returns the FinanceProcessOwners field if non-nil, zero value otherwise.

### GetFinanceProcessOwnersOk

`func (o *IssuesPutPreviousValues) GetFinanceProcessOwnersOk() (*interface{}, bool)`

GetFinanceProcessOwnersOk returns a tuple with the FinanceProcessOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinanceProcessOwners

`func (o *IssuesPutPreviousValues) SetFinanceProcessOwners(v interface{})`

SetFinanceProcessOwners sets FinanceProcessOwners field to given value.

### HasFinanceProcessOwners

`func (o *IssuesPutPreviousValues) HasFinanceProcessOwners() bool`

HasFinanceProcessOwners returns a boolean if a field has been set.

### SetFinanceProcessOwnersNil

`func (o *IssuesPutPreviousValues) SetFinanceProcessOwnersNil(b bool)`

 SetFinanceProcessOwnersNil sets the value for FinanceProcessOwners to be an explicit nil

### UnsetFinanceProcessOwners
`func (o *IssuesPutPreviousValues) UnsetFinanceProcessOwners()`

UnsetFinanceProcessOwners ensures that no value is present for FinanceProcessOwners, not even an explicit nil
### GetBusinessProcessOwners

`func (o *IssuesPutPreviousValues) GetBusinessProcessOwners() interface{}`

GetBusinessProcessOwners returns the BusinessProcessOwners field if non-nil, zero value otherwise.

### GetBusinessProcessOwnersOk

`func (o *IssuesPutPreviousValues) GetBusinessProcessOwnersOk() (*interface{}, bool)`

GetBusinessProcessOwnersOk returns a tuple with the BusinessProcessOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProcessOwners

`func (o *IssuesPutPreviousValues) SetBusinessProcessOwners(v interface{})`

SetBusinessProcessOwners sets BusinessProcessOwners field to given value.

### HasBusinessProcessOwners

`func (o *IssuesPutPreviousValues) HasBusinessProcessOwners() bool`

HasBusinessProcessOwners returns a boolean if a field has been set.

### SetBusinessProcessOwnersNil

`func (o *IssuesPutPreviousValues) SetBusinessProcessOwnersNil(b bool)`

 SetBusinessProcessOwnersNil sets the value for BusinessProcessOwners to be an explicit nil

### UnsetBusinessProcessOwners
`func (o *IssuesPutPreviousValues) UnsetBusinessProcessOwners()`

UnsetBusinessProcessOwners ensures that no value is present for BusinessProcessOwners, not even an explicit nil
### GetScopes

`func (o *IssuesPutPreviousValues) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *IssuesPutPreviousValues) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *IssuesPutPreviousValues) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *IssuesPutPreviousValues) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *IssuesPutPreviousValues) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *IssuesPutPreviousValues) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetExecutiveUserId

`func (o *IssuesPutPreviousValues) GetExecutiveUserId() int32`

GetExecutiveUserId returns the ExecutiveUserId field if non-nil, zero value otherwise.

### GetExecutiveUserIdOk

`func (o *IssuesPutPreviousValues) GetExecutiveUserIdOk() (*int32, bool)`

GetExecutiveUserIdOk returns a tuple with the ExecutiveUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutiveUserId

`func (o *IssuesPutPreviousValues) SetExecutiveUserId(v int32)`

SetExecutiveUserId sets ExecutiveUserId field to given value.

### HasExecutiveUserId

`func (o *IssuesPutPreviousValues) HasExecutiveUserId() bool`

HasExecutiveUserId returns a boolean if a field has been set.

### GetSeniorOwnerUserId

`func (o *IssuesPutPreviousValues) GetSeniorOwnerUserId() int32`

GetSeniorOwnerUserId returns the SeniorOwnerUserId field if non-nil, zero value otherwise.

### GetSeniorOwnerUserIdOk

`func (o *IssuesPutPreviousValues) GetSeniorOwnerUserIdOk() (*int32, bool)`

GetSeniorOwnerUserIdOk returns a tuple with the SeniorOwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeniorOwnerUserId

`func (o *IssuesPutPreviousValues) SetSeniorOwnerUserId(v int32)`

SetSeniorOwnerUserId sets SeniorOwnerUserId field to given value.

### HasSeniorOwnerUserId

`func (o *IssuesPutPreviousValues) HasSeniorOwnerUserId() bool`

HasSeniorOwnerUserId returns a boolean if a field has been set.

### GetFirstRemediatedDate

`func (o *IssuesPutPreviousValues) GetFirstRemediatedDate() string`

GetFirstRemediatedDate returns the FirstRemediatedDate field if non-nil, zero value otherwise.

### GetFirstRemediatedDateOk

`func (o *IssuesPutPreviousValues) GetFirstRemediatedDateOk() (*string, bool)`

GetFirstRemediatedDateOk returns a tuple with the FirstRemediatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstRemediatedDate

`func (o *IssuesPutPreviousValues) SetFirstRemediatedDate(v string)`

SetFirstRemediatedDate sets FirstRemediatedDate field to given value.

### HasFirstRemediatedDate

`func (o *IssuesPutPreviousValues) HasFirstRemediatedDate() bool`

HasFirstRemediatedDate returns a boolean if a field has been set.

### GetFirstPendingRemediationDate

`func (o *IssuesPutPreviousValues) GetFirstPendingRemediationDate() string`

GetFirstPendingRemediationDate returns the FirstPendingRemediationDate field if non-nil, zero value otherwise.

### GetFirstPendingRemediationDateOk

`func (o *IssuesPutPreviousValues) GetFirstPendingRemediationDateOk() (*string, bool)`

GetFirstPendingRemediationDateOk returns a tuple with the FirstPendingRemediationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPendingRemediationDate

`func (o *IssuesPutPreviousValues) SetFirstPendingRemediationDate(v string)`

SetFirstPendingRemediationDate sets FirstPendingRemediationDate field to given value.

### HasFirstPendingRemediationDate

`func (o *IssuesPutPreviousValues) HasFirstPendingRemediationDate() bool`

HasFirstPendingRemediationDate returns a boolean if a field has been set.

### GetFirstClosedDate

`func (o *IssuesPutPreviousValues) GetFirstClosedDate() string`

GetFirstClosedDate returns the FirstClosedDate field if non-nil, zero value otherwise.

### GetFirstClosedDateOk

`func (o *IssuesPutPreviousValues) GetFirstClosedDateOk() (*string, bool)`

GetFirstClosedDateOk returns a tuple with the FirstClosedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstClosedDate

`func (o *IssuesPutPreviousValues) SetFirstClosedDate(v string)`

SetFirstClosedDate sets FirstClosedDate field to given value.

### HasFirstClosedDate

`func (o *IssuesPutPreviousValues) HasFirstClosedDate() bool`

HasFirstClosedDate returns a boolean if a field has been set.

### GetCustomCurrency1

`func (o *IssuesPutPreviousValues) GetCustomCurrency1() float32`

GetCustomCurrency1 returns the CustomCurrency1 field if non-nil, zero value otherwise.

### GetCustomCurrency1Ok

`func (o *IssuesPutPreviousValues) GetCustomCurrency1Ok() (*float32, bool)`

GetCustomCurrency1Ok returns a tuple with the CustomCurrency1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCurrency1

`func (o *IssuesPutPreviousValues) SetCustomCurrency1(v float32)`

SetCustomCurrency1 sets CustomCurrency1 field to given value.

### HasCustomCurrency1

`func (o *IssuesPutPreviousValues) HasCustomCurrency1() bool`

HasCustomCurrency1 returns a boolean if a field has been set.

### GetCustomCurrency2

`func (o *IssuesPutPreviousValues) GetCustomCurrency2() float32`

GetCustomCurrency2 returns the CustomCurrency2 field if non-nil, zero value otherwise.

### GetCustomCurrency2Ok

`func (o *IssuesPutPreviousValues) GetCustomCurrency2Ok() (*float32, bool)`

GetCustomCurrency2Ok returns a tuple with the CustomCurrency2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCurrency2

`func (o *IssuesPutPreviousValues) SetCustomCurrency2(v float32)`

SetCustomCurrency2 sets CustomCurrency2 field to given value.

### HasCustomCurrency2

`func (o *IssuesPutPreviousValues) HasCustomCurrency2() bool`

HasCustomCurrency2 returns a boolean if a field has been set.

### GetCustomCurrency3

`func (o *IssuesPutPreviousValues) GetCustomCurrency3() float32`

GetCustomCurrency3 returns the CustomCurrency3 field if non-nil, zero value otherwise.

### GetCustomCurrency3Ok

`func (o *IssuesPutPreviousValues) GetCustomCurrency3Ok() (*float32, bool)`

GetCustomCurrency3Ok returns a tuple with the CustomCurrency3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCurrency3

`func (o *IssuesPutPreviousValues) SetCustomCurrency3(v float32)`

SetCustomCurrency3 sets CustomCurrency3 field to given value.

### HasCustomCurrency3

`func (o *IssuesPutPreviousValues) HasCustomCurrency3() bool`

HasCustomCurrency3 returns a boolean if a field has been set.

### GetCustomCurrency4

`func (o *IssuesPutPreviousValues) GetCustomCurrency4() float32`

GetCustomCurrency4 returns the CustomCurrency4 field if non-nil, zero value otherwise.

### GetCustomCurrency4Ok

`func (o *IssuesPutPreviousValues) GetCustomCurrency4Ok() (*float32, bool)`

GetCustomCurrency4Ok returns a tuple with the CustomCurrency4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCurrency4

`func (o *IssuesPutPreviousValues) SetCustomCurrency4(v float32)`

SetCustomCurrency4 sets CustomCurrency4 field to given value.

### HasCustomCurrency4

`func (o *IssuesPutPreviousValues) HasCustomCurrency4() bool`

HasCustomCurrency4 returns a boolean if a field has been set.

### GetCustomDecimal1

`func (o *IssuesPutPreviousValues) GetCustomDecimal1() float32`

GetCustomDecimal1 returns the CustomDecimal1 field if non-nil, zero value otherwise.

### GetCustomDecimal1Ok

`func (o *IssuesPutPreviousValues) GetCustomDecimal1Ok() (*float32, bool)`

GetCustomDecimal1Ok returns a tuple with the CustomDecimal1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDecimal1

`func (o *IssuesPutPreviousValues) SetCustomDecimal1(v float32)`

SetCustomDecimal1 sets CustomDecimal1 field to given value.

### HasCustomDecimal1

`func (o *IssuesPutPreviousValues) HasCustomDecimal1() bool`

HasCustomDecimal1 returns a boolean if a field has been set.

### GetCustomDecimal2

`func (o *IssuesPutPreviousValues) GetCustomDecimal2() float32`

GetCustomDecimal2 returns the CustomDecimal2 field if non-nil, zero value otherwise.

### GetCustomDecimal2Ok

`func (o *IssuesPutPreviousValues) GetCustomDecimal2Ok() (*float32, bool)`

GetCustomDecimal2Ok returns a tuple with the CustomDecimal2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDecimal2

`func (o *IssuesPutPreviousValues) SetCustomDecimal2(v float32)`

SetCustomDecimal2 sets CustomDecimal2 field to given value.

### HasCustomDecimal2

`func (o *IssuesPutPreviousValues) HasCustomDecimal2() bool`

HasCustomDecimal2 returns a boolean if a field has been set.

### GetCustomDecimal3

`func (o *IssuesPutPreviousValues) GetCustomDecimal3() float32`

GetCustomDecimal3 returns the CustomDecimal3 field if non-nil, zero value otherwise.

### GetCustomDecimal3Ok

`func (o *IssuesPutPreviousValues) GetCustomDecimal3Ok() (*float32, bool)`

GetCustomDecimal3Ok returns a tuple with the CustomDecimal3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDecimal3

`func (o *IssuesPutPreviousValues) SetCustomDecimal3(v float32)`

SetCustomDecimal3 sets CustomDecimal3 field to given value.

### HasCustomDecimal3

`func (o *IssuesPutPreviousValues) HasCustomDecimal3() bool`

HasCustomDecimal3 returns a boolean if a field has been set.

### GetFindingOwnerUserId

`func (o *IssuesPutPreviousValues) GetFindingOwnerUserId() int32`

GetFindingOwnerUserId returns the FindingOwnerUserId field if non-nil, zero value otherwise.

### GetFindingOwnerUserIdOk

`func (o *IssuesPutPreviousValues) GetFindingOwnerUserIdOk() (*int32, bool)`

GetFindingOwnerUserIdOk returns a tuple with the FindingOwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingOwnerUserId

`func (o *IssuesPutPreviousValues) SetFindingOwnerUserId(v int32)`

SetFindingOwnerUserId sets FindingOwnerUserId field to given value.

### HasFindingOwnerUserId

`func (o *IssuesPutPreviousValues) HasFindingOwnerUserId() bool`

HasFindingOwnerUserId returns a boolean if a field has been set.

### GetCustomDate1

`func (o *IssuesPutPreviousValues) GetCustomDate1() string`

GetCustomDate1 returns the CustomDate1 field if non-nil, zero value otherwise.

### GetCustomDate1Ok

`func (o *IssuesPutPreviousValues) GetCustomDate1Ok() (*string, bool)`

GetCustomDate1Ok returns a tuple with the CustomDate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate1

`func (o *IssuesPutPreviousValues) SetCustomDate1(v string)`

SetCustomDate1 sets CustomDate1 field to given value.

### HasCustomDate1

`func (o *IssuesPutPreviousValues) HasCustomDate1() bool`

HasCustomDate1 returns a boolean if a field has been set.

### GetCustomDate2

`func (o *IssuesPutPreviousValues) GetCustomDate2() string`

GetCustomDate2 returns the CustomDate2 field if non-nil, zero value otherwise.

### GetCustomDate2Ok

`func (o *IssuesPutPreviousValues) GetCustomDate2Ok() (*string, bool)`

GetCustomDate2Ok returns a tuple with the CustomDate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate2

`func (o *IssuesPutPreviousValues) SetCustomDate2(v string)`

SetCustomDate2 sets CustomDate2 field to given value.

### HasCustomDate2

`func (o *IssuesPutPreviousValues) HasCustomDate2() bool`

HasCustomDate2 returns a boolean if a field has been set.

### GetCustomDate3

`func (o *IssuesPutPreviousValues) GetCustomDate3() string`

GetCustomDate3 returns the CustomDate3 field if non-nil, zero value otherwise.

### GetCustomDate3Ok

`func (o *IssuesPutPreviousValues) GetCustomDate3Ok() (*string, bool)`

GetCustomDate3Ok returns a tuple with the CustomDate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate3

`func (o *IssuesPutPreviousValues) SetCustomDate3(v string)`

SetCustomDate3 sets CustomDate3 field to given value.

### HasCustomDate3

`func (o *IssuesPutPreviousValues) HasCustomDate3() bool`

HasCustomDate3 returns a boolean if a field has been set.

### GetCustomDate4

`func (o *IssuesPutPreviousValues) GetCustomDate4() string`

GetCustomDate4 returns the CustomDate4 field if non-nil, zero value otherwise.

### GetCustomDate4Ok

`func (o *IssuesPutPreviousValues) GetCustomDate4Ok() (*string, bool)`

GetCustomDate4Ok returns a tuple with the CustomDate4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate4

`func (o *IssuesPutPreviousValues) SetCustomDate4(v string)`

SetCustomDate4 sets CustomDate4 field to given value.

### HasCustomDate4

`func (o *IssuesPutPreviousValues) HasCustomDate4() bool`

HasCustomDate4 returns a boolean if a field has been set.

### GetIssueCustomSelect9OptionId

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect9OptionId() int32`

GetIssueCustomSelect9OptionId returns the IssueCustomSelect9OptionId field if non-nil, zero value otherwise.

### GetIssueCustomSelect9OptionIdOk

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect9OptionIdOk() (*int32, bool)`

GetIssueCustomSelect9OptionIdOk returns a tuple with the IssueCustomSelect9OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCustomSelect9OptionId

`func (o *IssuesPutPreviousValues) SetIssueCustomSelect9OptionId(v int32)`

SetIssueCustomSelect9OptionId sets IssueCustomSelect9OptionId field to given value.

### HasIssueCustomSelect9OptionId

`func (o *IssuesPutPreviousValues) HasIssueCustomSelect9OptionId() bool`

HasIssueCustomSelect9OptionId returns a boolean if a field has been set.

### GetIssueCustomSelect10OptionId

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect10OptionId() int32`

GetIssueCustomSelect10OptionId returns the IssueCustomSelect10OptionId field if non-nil, zero value otherwise.

### GetIssueCustomSelect10OptionIdOk

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect10OptionIdOk() (*int32, bool)`

GetIssueCustomSelect10OptionIdOk returns a tuple with the IssueCustomSelect10OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCustomSelect10OptionId

`func (o *IssuesPutPreviousValues) SetIssueCustomSelect10OptionId(v int32)`

SetIssueCustomSelect10OptionId sets IssueCustomSelect10OptionId field to given value.

### HasIssueCustomSelect10OptionId

`func (o *IssuesPutPreviousValues) HasIssueCustomSelect10OptionId() bool`

HasIssueCustomSelect10OptionId returns a boolean if a field has been set.

### GetIssueCustomSelect11OptionId

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect11OptionId() int32`

GetIssueCustomSelect11OptionId returns the IssueCustomSelect11OptionId field if non-nil, zero value otherwise.

### GetIssueCustomSelect11OptionIdOk

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect11OptionIdOk() (*int32, bool)`

GetIssueCustomSelect11OptionIdOk returns a tuple with the IssueCustomSelect11OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCustomSelect11OptionId

`func (o *IssuesPutPreviousValues) SetIssueCustomSelect11OptionId(v int32)`

SetIssueCustomSelect11OptionId sets IssueCustomSelect11OptionId field to given value.

### HasIssueCustomSelect11OptionId

`func (o *IssuesPutPreviousValues) HasIssueCustomSelect11OptionId() bool`

HasIssueCustomSelect11OptionId returns a boolean if a field has been set.

### GetIssueCustomSelect12OptionId

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect12OptionId() int32`

GetIssueCustomSelect12OptionId returns the IssueCustomSelect12OptionId field if non-nil, zero value otherwise.

### GetIssueCustomSelect12OptionIdOk

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect12OptionIdOk() (*int32, bool)`

GetIssueCustomSelect12OptionIdOk returns a tuple with the IssueCustomSelect12OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCustomSelect12OptionId

`func (o *IssuesPutPreviousValues) SetIssueCustomSelect12OptionId(v int32)`

SetIssueCustomSelect12OptionId sets IssueCustomSelect12OptionId field to given value.

### HasIssueCustomSelect12OptionId

`func (o *IssuesPutPreviousValues) HasIssueCustomSelect12OptionId() bool`

HasIssueCustomSelect12OptionId returns a boolean if a field has been set.

### GetIssueCustomSelect13OptionId

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect13OptionId() int32`

GetIssueCustomSelect13OptionId returns the IssueCustomSelect13OptionId field if non-nil, zero value otherwise.

### GetIssueCustomSelect13OptionIdOk

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect13OptionIdOk() (*int32, bool)`

GetIssueCustomSelect13OptionIdOk returns a tuple with the IssueCustomSelect13OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCustomSelect13OptionId

`func (o *IssuesPutPreviousValues) SetIssueCustomSelect13OptionId(v int32)`

SetIssueCustomSelect13OptionId sets IssueCustomSelect13OptionId field to given value.

### HasIssueCustomSelect13OptionId

`func (o *IssuesPutPreviousValues) HasIssueCustomSelect13OptionId() bool`

HasIssueCustomSelect13OptionId returns a boolean if a field has been set.

### GetIssueCustomSelect14OptionId

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect14OptionId() int32`

GetIssueCustomSelect14OptionId returns the IssueCustomSelect14OptionId field if non-nil, zero value otherwise.

### GetIssueCustomSelect14OptionIdOk

`func (o *IssuesPutPreviousValues) GetIssueCustomSelect14OptionIdOk() (*int32, bool)`

GetIssueCustomSelect14OptionIdOk returns a tuple with the IssueCustomSelect14OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCustomSelect14OptionId

`func (o *IssuesPutPreviousValues) SetIssueCustomSelect14OptionId(v int32)`

SetIssueCustomSelect14OptionId sets IssueCustomSelect14OptionId field to given value.

### HasIssueCustomSelect14OptionId

`func (o *IssuesPutPreviousValues) HasIssueCustomSelect14OptionId() bool`

HasIssueCustomSelect14OptionId returns a boolean if a field has been set.

### GetCancelByUserId

`func (o *IssuesPutPreviousValues) GetCancelByUserId() int32`

GetCancelByUserId returns the CancelByUserId field if non-nil, zero value otherwise.

### GetCancelByUserIdOk

`func (o *IssuesPutPreviousValues) GetCancelByUserIdOk() (*int32, bool)`

GetCancelByUserIdOk returns a tuple with the CancelByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelByUserId

`func (o *IssuesPutPreviousValues) SetCancelByUserId(v int32)`

SetCancelByUserId sets CancelByUserId field to given value.

### HasCancelByUserId

`func (o *IssuesPutPreviousValues) HasCancelByUserId() bool`

HasCancelByUserId returns a boolean if a field has been set.

### GetCancelDate

`func (o *IssuesPutPreviousValues) GetCancelDate() string`

GetCancelDate returns the CancelDate field if non-nil, zero value otherwise.

### GetCancelDateOk

`func (o *IssuesPutPreviousValues) GetCancelDateOk() (*string, bool)`

GetCancelDateOk returns a tuple with the CancelDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelDate

`func (o *IssuesPutPreviousValues) SetCancelDate(v string)`

SetCancelDate sets CancelDate field to given value.

### HasCancelDate

`func (o *IssuesPutPreviousValues) HasCancelDate() bool`

HasCancelDate returns a boolean if a field has been set.

### GetFirstCancelDate

`func (o *IssuesPutPreviousValues) GetFirstCancelDate() string`

GetFirstCancelDate returns the FirstCancelDate field if non-nil, zero value otherwise.

### GetFirstCancelDateOk

`func (o *IssuesPutPreviousValues) GetFirstCancelDateOk() (*string, bool)`

GetFirstCancelDateOk returns a tuple with the FirstCancelDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCancelDate

`func (o *IssuesPutPreviousValues) SetFirstCancelDate(v string)`

SetFirstCancelDate sets FirstCancelDate field to given value.

### HasFirstCancelDate

`func (o *IssuesPutPreviousValues) HasFirstCancelDate() bool`

HasFirstCancelDate returns a boolean if a field has been set.

### GetLastActionPlanDueDate

`func (o *IssuesPutPreviousValues) GetLastActionPlanDueDate() string`

GetLastActionPlanDueDate returns the LastActionPlanDueDate field if non-nil, zero value otherwise.

### GetLastActionPlanDueDateOk

`func (o *IssuesPutPreviousValues) GetLastActionPlanDueDateOk() (*string, bool)`

GetLastActionPlanDueDateOk returns a tuple with the LastActionPlanDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActionPlanDueDate

`func (o *IssuesPutPreviousValues) SetLastActionPlanDueDate(v string)`

SetLastActionPlanDueDate sets LastActionPlanDueDate field to given value.

### HasLastActionPlanDueDate

`func (o *IssuesPutPreviousValues) HasLastActionPlanDueDate() bool`

HasLastActionPlanDueDate returns a boolean if a field has been set.

### GetCancelReason

`func (o *IssuesPutPreviousValues) GetCancelReason() string`

GetCancelReason returns the CancelReason field if non-nil, zero value otherwise.

### GetCancelReasonOk

`func (o *IssuesPutPreviousValues) GetCancelReasonOk() (*string, bool)`

GetCancelReasonOk returns a tuple with the CancelReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelReason

`func (o *IssuesPutPreviousValues) SetCancelReason(v string)`

SetCancelReason sets CancelReason field to given value.

### HasCancelReason

`func (o *IssuesPutPreviousValues) HasCancelReason() bool`

HasCancelReason returns a boolean if a field has been set.

### GetDraftByUserId

`func (o *IssuesPutPreviousValues) GetDraftByUserId() int32`

GetDraftByUserId returns the DraftByUserId field if non-nil, zero value otherwise.

### GetDraftByUserIdOk

`func (o *IssuesPutPreviousValues) GetDraftByUserIdOk() (*int32, bool)`

GetDraftByUserIdOk returns a tuple with the DraftByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftByUserId

`func (o *IssuesPutPreviousValues) SetDraftByUserId(v int32)`

SetDraftByUserId sets DraftByUserId field to given value.

### HasDraftByUserId

`func (o *IssuesPutPreviousValues) HasDraftByUserId() bool`

HasDraftByUserId returns a boolean if a field has been set.

### GetDraftDate

`func (o *IssuesPutPreviousValues) GetDraftDate() string`

GetDraftDate returns the DraftDate field if non-nil, zero value otherwise.

### GetDraftDateOk

`func (o *IssuesPutPreviousValues) GetDraftDateOk() (*string, bool)`

GetDraftDateOk returns a tuple with the DraftDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftDate

`func (o *IssuesPutPreviousValues) SetDraftDate(v string)`

SetDraftDate sets DraftDate field to given value.

### HasDraftDate

`func (o *IssuesPutPreviousValues) HasDraftDate() bool`

HasDraftDate returns a boolean if a field has been set.

### GetFirstDraftDate

`func (o *IssuesPutPreviousValues) GetFirstDraftDate() string`

GetFirstDraftDate returns the FirstDraftDate field if non-nil, zero value otherwise.

### GetFirstDraftDateOk

`func (o *IssuesPutPreviousValues) GetFirstDraftDateOk() (*string, bool)`

GetFirstDraftDateOk returns a tuple with the FirstDraftDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDraftDate

`func (o *IssuesPutPreviousValues) SetFirstDraftDate(v string)`

SetFirstDraftDate sets FirstDraftDate field to given value.

### HasFirstDraftDate

`func (o *IssuesPutPreviousValues) HasFirstDraftDate() bool`

HasFirstDraftDate returns a boolean if a field has been set.

### GetOpenRevisionCount

`func (o *IssuesPutPreviousValues) GetOpenRevisionCount() int32`

GetOpenRevisionCount returns the OpenRevisionCount field if non-nil, zero value otherwise.

### GetOpenRevisionCountOk

`func (o *IssuesPutPreviousValues) GetOpenRevisionCountOk() (*int32, bool)`

GetOpenRevisionCountOk returns a tuple with the OpenRevisionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenRevisionCount

`func (o *IssuesPutPreviousValues) SetOpenRevisionCount(v int32)`

SetOpenRevisionCount sets OpenRevisionCount field to given value.

### HasOpenRevisionCount

`func (o *IssuesPutPreviousValues) HasOpenRevisionCount() bool`

HasOpenRevisionCount returns a boolean if a field has been set.

### GetCustomText7

`func (o *IssuesPutPreviousValues) GetCustomText7() string`

GetCustomText7 returns the CustomText7 field if non-nil, zero value otherwise.

### GetCustomText7Ok

`func (o *IssuesPutPreviousValues) GetCustomText7Ok() (*string, bool)`

GetCustomText7Ok returns a tuple with the CustomText7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText7

`func (o *IssuesPutPreviousValues) SetCustomText7(v string)`

SetCustomText7 sets CustomText7 field to given value.

### HasCustomText7

`func (o *IssuesPutPreviousValues) HasCustomText7() bool`

HasCustomText7 returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *IssuesPutPreviousValues) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *IssuesPutPreviousValues) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *IssuesPutPreviousValues) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *IssuesPutPreviousValues) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetFieldData

`func (o *IssuesPutPreviousValues) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *IssuesPutPreviousValues) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *IssuesPutPreviousValues) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *IssuesPutPreviousValues) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *IssuesPutPreviousValues) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *IssuesPutPreviousValues) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil
### GetComplianceAssessmentId

`func (o *IssuesPutPreviousValues) GetComplianceAssessmentId() int32`

GetComplianceAssessmentId returns the ComplianceAssessmentId field if non-nil, zero value otherwise.

### GetComplianceAssessmentIdOk

`func (o *IssuesPutPreviousValues) GetComplianceAssessmentIdOk() (*int32, bool)`

GetComplianceAssessmentIdOk returns a tuple with the ComplianceAssessmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceAssessmentId

`func (o *IssuesPutPreviousValues) SetComplianceAssessmentId(v int32)`

SetComplianceAssessmentId sets ComplianceAssessmentId field to given value.

### HasComplianceAssessmentId

`func (o *IssuesPutPreviousValues) HasComplianceAssessmentId() bool`

HasComplianceAssessmentId returns a boolean if a field has been set.

### GetComplianceAssessmentItemId

`func (o *IssuesPutPreviousValues) GetComplianceAssessmentItemId() int32`

GetComplianceAssessmentItemId returns the ComplianceAssessmentItemId field if non-nil, zero value otherwise.

### GetComplianceAssessmentItemIdOk

`func (o *IssuesPutPreviousValues) GetComplianceAssessmentItemIdOk() (*int32, bool)`

GetComplianceAssessmentItemIdOk returns a tuple with the ComplianceAssessmentItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceAssessmentItemId

`func (o *IssuesPutPreviousValues) SetComplianceAssessmentItemId(v int32)`

SetComplianceAssessmentItemId sets ComplianceAssessmentItemId field to given value.

### HasComplianceAssessmentItemId

`func (o *IssuesPutPreviousValues) HasComplianceAssessmentItemId() bool`

HasComplianceAssessmentItemId returns a boolean if a field has been set.

### GetAuditSurveyId

`func (o *IssuesPutPreviousValues) GetAuditSurveyId() int32`

GetAuditSurveyId returns the AuditSurveyId field if non-nil, zero value otherwise.

### GetAuditSurveyIdOk

`func (o *IssuesPutPreviousValues) GetAuditSurveyIdOk() (*int32, bool)`

GetAuditSurveyIdOk returns a tuple with the AuditSurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSurveyId

`func (o *IssuesPutPreviousValues) SetAuditSurveyId(v int32)`

SetAuditSurveyId sets AuditSurveyId field to given value.

### HasAuditSurveyId

`func (o *IssuesPutPreviousValues) HasAuditSurveyId() bool`

HasAuditSurveyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


