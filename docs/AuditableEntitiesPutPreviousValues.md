# AuditableEntitiesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CustomEntityReference** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**AuditableEntityRegionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;auditable_entity_regions.id&#x60;.&lt;fk table&#x3D;&#39;auditable_entity_regions&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AuditableEntityTypeId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;auditable_entity_types.id&#x60;.&lt;fk table&#x3D;&#39;auditable_entity_types&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AuditRotationScheduleId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_rotation_schedules.id&#x60;.&lt;fk table&#x3D;&#39;audit_rotation_schedules&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AuditOfficeLocationId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_office_locations.id&#x60;.&lt;fk table&#x3D;&#39;audit_office_locations&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CustomText1** | Pointer to **string** |  | [optional] 
**CustomText2** | Pointer to **string** |  | [optional] 
**CustomText3** | Pointer to **string** |  | [optional] 
**CustomText4** | Pointer to **string** |  | [optional] 
**CustomText5** | Pointer to **string** |  | [optional] 
**CustomText6** | Pointer to **string** |  | [optional] 
**CustomText7** | Pointer to **string** |  | [optional] 
**CustomDate1** | Pointer to **string** |  | [optional] 
**AeCustomSelect1OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ae_custom_select1_options.id&#x60;.&lt;fk table&#x3D;&#39;ae_custom_select1_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AuditDivisionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_divisions.id&#x60;.&lt;fk table&#x3D;&#39;audit_divisions&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AuditBusinessSegmentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_business_segments.id&#x60;.&lt;fk table&#x3D;&#39;audit_business_segments&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AeCustomSelect2OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ae_custom_select2_options.id&#x60;.&lt;fk table&#x3D;&#39;ae_custom_select2_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomDate2** | Pointer to **string** |  | [optional] 
**AssessmentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;assessments.id&#x60;.&lt;fk table&#x3D;&#39;assessments&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomText8** | Pointer to **string** |  | [optional] 
**CustomText9** | Pointer to **string** |  | [optional] 
**CustomText10** | Pointer to **string** |  | [optional] 
**CustomText11** | Pointer to **string** |  | [optional] 
**CustomText12** | Pointer to **string** |  | [optional] 
**CustomText13** | Pointer to **string** |  | [optional] 
**CustomText14** | Pointer to **string** |  | [optional] 
**CustomText15** | Pointer to **string** |  | [optional] 
**ReviewerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AeCustomSelect3OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ae_custom_select3_options.id&#x60;.&lt;fk table&#x3D;&#39;ae_custom_select3_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AeCustomSelect4OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ae_custom_select4_options.id&#x60;.&lt;fk table&#x3D;&#39;ae_custom_select4_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AeCustomSelect5OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ae_custom_select5_options.id&#x60;.&lt;fk table&#x3D;&#39;ae_custom_select5_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AeCustomSelect6OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ae_custom_select6_options.id&#x60;.&lt;fk table&#x3D;&#39;ae_custom_select6_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AeCustomSelect7OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ae_custom_select7_options.id&#x60;.&lt;fk table&#x3D;&#39;ae_custom_select7_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AeCustomSelect8OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ae_custom_select8_options.id&#x60;.&lt;fk table&#x3D;&#39;ae_custom_select8_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AeStatusOptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ae_status_options.id&#x60;.&lt;fk table&#x3D;&#39;ae_status_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomDate3** | Pointer to **string** |  | [optional] 
**EffectiveDate** | Pointer to **string** |  | [optional] 
**AeCustomSelect9OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ae_custom_select9_options.id&#x60;.&lt;fk table&#x3D;&#39;ae_custom_select9_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AeCustomSelect10OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ae_custom_select10_options.id&#x60;.&lt;fk table&#x3D;&#39;ae_custom_select10_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AeCustomSelect11OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ae_custom_select11_options.id&#x60;.&lt;fk table&#x3D;&#39;ae_custom_select11_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AeCustomSelect12OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ae_custom_select12_options.id&#x60;.&lt;fk table&#x3D;&#39;ae_custom_select12_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomDate4** | Pointer to **string** |  | [optional] 
**ExecutiveUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**VicePresidentUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**InherentRisk** | Pointer to **interface{}** |  | [optional] 
**ResidualRisk** | Pointer to **interface{}** |  | [optional] 
**ManagerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CoordinatorUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ProductStatusId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;product_statuses.id&#x60;.&lt;fk table&#x3D;&#39;product_statuses&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**VendorCriticalityId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;vendor_criticalities.id&#x60;.&lt;fk table&#x3D;&#39;vendor_criticalities&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**LegalEntityTypeId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;legal_entity_types.id&#x60;.&lt;fk table&#x3D;&#39;legal_entity_types&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ItAssetTypeId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;it_asset_types.id&#x60;.&lt;fk table&#x3D;&#39;it_asset_types&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ItAssetStatusId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;it_asset_statuses.id&#x60;.&lt;fk table&#x3D;&#39;it_asset_statuses&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CurrentVersion** | Pointer to **string** |  | [optional] 
**ContractDetails** | Pointer to **string** |  | [optional] 
**ContactName** | Pointer to **string** |  | [optional] 
**ContactEmail** | Pointer to **string** |  | [optional] 
**ContactPhoneNumber** | Pointer to **string** |  | [optional] 
**ItServiceLevelAgreement** | Pointer to **string** |  | [optional] 
**ItRecoveryTimeObjective** | Pointer to **string** |  | [optional] 
**ItRecoveryPointObjective** | Pointer to **string** |  | [optional] 
**OldItSecurityIncidentHours** | Pointer to **float32** |  | [optional] 
**EmployeesSize** | Pointer to **float32** |  | [optional] 
**RegulatoryAuditFrequencyId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_frequencies.id&#x60;.&lt;fk table&#x3D;&#39;audit_frequencies&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RegulatoryAuditFrequencyNotes** | Pointer to **string** |  | [optional] 
**CalculatedAuditFrequencyId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_frequencies.id&#x60;.&lt;fk table&#x3D;&#39;audit_frequencies&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CreatedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomDecimal1** | Pointer to **float32** |  | [optional] 
**CustomDecimal2** | Pointer to **float32** |  | [optional] 
**CustomDecimal3** | Pointer to **float32** |  | [optional] 
**CustomDecimal4** | Pointer to **float32** |  | [optional] 
**CustomDecimal5** | Pointer to **float32** |  | [optional] 
**CustomDecimal6** | Pointer to **float32** |  | [optional] 
**CustomDecimal7** | Pointer to **float32** |  | [optional] 
**CustomDecimal8** | Pointer to **float32** |  | [optional] 
**CustomCurrency1** | Pointer to **float32** |  | [optional] 
**CustomCurrency2** | Pointer to **float32** |  | [optional] 
**CustomCurrency3** | Pointer to **float32** |  | [optional] 
**CustomCurrency4** | Pointer to **float32** |  | [optional] 
**FieldData** | Pointer to **interface{}** |  | [optional] 
**DefaultOpsAuditTemplateId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audits.id&#x60;.&lt;fk table&#x3D;&#39;ops_audits&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskScoreId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_scores.id&#x60;.&lt;fk table&#x3D;&#39;risk_scores&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**VendorStatus** | Pointer to **string** |  | [optional] 
**ContactUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**PriorAuditId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audits.id&#x60;.&lt;fk table&#x3D;&#39;ops_audits&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**UpcomingAuditId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audits.id&#x60;.&lt;fk table&#x3D;&#39;ops_audits&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**PriorAuditFinalReportDate** | Pointer to **string** |  | [optional] 
**FormulaData** | Pointer to **interface{}** |  | [optional] 
**IsFlagged** | Pointer to **bool** |  | [optional] [default to false]
**RiskNextAuditDue** | Pointer to **string** |  | [optional] 
**RegulatoryNextAuditDue** | Pointer to **string** |  | [optional] 
**RegulatoryNextAuditDueDate** | Pointer to **string** |  | [optional] 
**RiskNextAuditDueDate** | Pointer to **string** |  | [optional] 
**ItSecurityIncidentHours** | Pointer to **float32** |  | [optional] 
**ResidualRiskCalc** | Pointer to **interface{}** |  | [optional] 
**SystemComponentTypeId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;system_component_types.id&#x60;.&lt;fk table&#x3D;&#39;system_component_types&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**SoxScopeId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;sox_scopes.id&#x60;.&lt;fk table&#x3D;&#39;sox_scopes&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**EstimatedSpend** | Pointer to **float32** |  | [optional] 
**ScopeRationale** | Pointer to **string** |  | [optional] [default to "false"]
**FinancialApplicationId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;financial_applications.id&#x60;.&lt;fk table&#x3D;&#39;financial_applications&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**PriorAuditTitle** | Pointer to **string** |  | [optional] 
**UpcomingAuditTitle** | Pointer to **string** |  | [optional] 
**PriorAuditOpinion** | Pointer to **string** |  | [optional] 
**UpcomingAuditStartDate** | Pointer to **string** |  | [optional] 
**UpcomingAuditEndDate** | Pointer to **string** |  | [optional] 
**ConfidentialityImpact** | Pointer to **float32** |  | [optional] 
**IntegrityImpact** | Pointer to **float32** |  | [optional] 
**AvailabilityImpact** | Pointer to **float32** |  | [optional] 
**CiaInherentRiskCalc** | Pointer to **interface{}** |  | [optional] 
**CiaResidualRiskCalc** | Pointer to **interface{}** |  | [optional] 
**IntakeStatus** | Pointer to **string** |  | [optional] 
**CustomDate5** | Pointer to **string** |  | [optional] 
**CustomDate6** | Pointer to **string** |  | [optional] 
**CustomDate7** | Pointer to **string** |  | [optional] 
**CustomDate8** | Pointer to **string** |  | [optional] 
**CustomDate9** | Pointer to **string** |  | [optional] 
**CustomDate10** | Pointer to **string** |  | [optional] 
**AeCustomSelect13OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ae_custom_select13_options.id&#x60;.&lt;fk table&#x3D;&#39;ae_custom_select13_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AeCustomSelect14OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ae_custom_select14_options.id&#x60;.&lt;fk table&#x3D;&#39;ae_custom_select14_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AeCustomSelect15OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ae_custom_select15_options.id&#x60;.&lt;fk table&#x3D;&#39;ae_custom_select15_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**DataClassificationId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;data_classifications.id&#x60;.&lt;fk table&#x3D;&#39;data_classifications&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
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
**BusinessOwnerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TechnicalProductOwnerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ThirdPartyResidualRiskCalc** | Pointer to **interface{}** |  | [optional] 
**AuditableEntityParentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;auditable_entities.id&#x60;.&lt;fk table&#x3D;&#39;auditable_entities&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 

## Methods

### NewAuditableEntitiesPutPreviousValues

`func NewAuditableEntitiesPutPreviousValues() *AuditableEntitiesPutPreviousValues`

NewAuditableEntitiesPutPreviousValues instantiates a new AuditableEntitiesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditableEntitiesPutPreviousValuesWithDefaults

`func NewAuditableEntitiesPutPreviousValuesWithDefaults() *AuditableEntitiesPutPreviousValues`

NewAuditableEntitiesPutPreviousValuesWithDefaults instantiates a new AuditableEntitiesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditableEntitiesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditableEntitiesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuditableEntitiesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuditableEntitiesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditableEntitiesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditableEntitiesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuditableEntitiesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuditableEntitiesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuditableEntitiesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuditableEntitiesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuditableEntitiesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AuditableEntitiesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AuditableEntitiesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AuditableEntitiesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AuditableEntitiesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetUid

`func (o *AuditableEntitiesPutPreviousValues) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AuditableEntitiesPutPreviousValues) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AuditableEntitiesPutPreviousValues) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AuditableEntitiesPutPreviousValues) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetName

`func (o *AuditableEntitiesPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditableEntitiesPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditableEntitiesPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuditableEntitiesPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCustomEntityReference

`func (o *AuditableEntitiesPutPreviousValues) GetCustomEntityReference() string`

GetCustomEntityReference returns the CustomEntityReference field if non-nil, zero value otherwise.

### GetCustomEntityReferenceOk

`func (o *AuditableEntitiesPutPreviousValues) GetCustomEntityReferenceOk() (*string, bool)`

GetCustomEntityReferenceOk returns a tuple with the CustomEntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEntityReference

`func (o *AuditableEntitiesPutPreviousValues) SetCustomEntityReference(v string)`

SetCustomEntityReference sets CustomEntityReference field to given value.

### HasCustomEntityReference

`func (o *AuditableEntitiesPutPreviousValues) HasCustomEntityReference() bool`

HasCustomEntityReference returns a boolean if a field has been set.

### GetDescription

`func (o *AuditableEntitiesPutPreviousValues) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuditableEntitiesPutPreviousValues) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuditableEntitiesPutPreviousValues) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuditableEntitiesPutPreviousValues) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotes

`func (o *AuditableEntitiesPutPreviousValues) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AuditableEntitiesPutPreviousValues) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AuditableEntitiesPutPreviousValues) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AuditableEntitiesPutPreviousValues) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAuditableEntityRegionId

`func (o *AuditableEntitiesPutPreviousValues) GetAuditableEntityRegionId() int32`

GetAuditableEntityRegionId returns the AuditableEntityRegionId field if non-nil, zero value otherwise.

### GetAuditableEntityRegionIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAuditableEntityRegionIdOk() (*int32, bool)`

GetAuditableEntityRegionIdOk returns a tuple with the AuditableEntityRegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditableEntityRegionId

`func (o *AuditableEntitiesPutPreviousValues) SetAuditableEntityRegionId(v int32)`

SetAuditableEntityRegionId sets AuditableEntityRegionId field to given value.

### HasAuditableEntityRegionId

`func (o *AuditableEntitiesPutPreviousValues) HasAuditableEntityRegionId() bool`

HasAuditableEntityRegionId returns a boolean if a field has been set.

### GetAuditableEntityTypeId

`func (o *AuditableEntitiesPutPreviousValues) GetAuditableEntityTypeId() int32`

GetAuditableEntityTypeId returns the AuditableEntityTypeId field if non-nil, zero value otherwise.

### GetAuditableEntityTypeIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAuditableEntityTypeIdOk() (*int32, bool)`

GetAuditableEntityTypeIdOk returns a tuple with the AuditableEntityTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditableEntityTypeId

`func (o *AuditableEntitiesPutPreviousValues) SetAuditableEntityTypeId(v int32)`

SetAuditableEntityTypeId sets AuditableEntityTypeId field to given value.

### HasAuditableEntityTypeId

`func (o *AuditableEntitiesPutPreviousValues) HasAuditableEntityTypeId() bool`

HasAuditableEntityTypeId returns a boolean if a field has been set.

### GetAuditRotationScheduleId

`func (o *AuditableEntitiesPutPreviousValues) GetAuditRotationScheduleId() int32`

GetAuditRotationScheduleId returns the AuditRotationScheduleId field if non-nil, zero value otherwise.

### GetAuditRotationScheduleIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAuditRotationScheduleIdOk() (*int32, bool)`

GetAuditRotationScheduleIdOk returns a tuple with the AuditRotationScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditRotationScheduleId

`func (o *AuditableEntitiesPutPreviousValues) SetAuditRotationScheduleId(v int32)`

SetAuditRotationScheduleId sets AuditRotationScheduleId field to given value.

### HasAuditRotationScheduleId

`func (o *AuditableEntitiesPutPreviousValues) HasAuditRotationScheduleId() bool`

HasAuditRotationScheduleId returns a boolean if a field has been set.

### GetAuditOfficeLocationId

`func (o *AuditableEntitiesPutPreviousValues) GetAuditOfficeLocationId() int32`

GetAuditOfficeLocationId returns the AuditOfficeLocationId field if non-nil, zero value otherwise.

### GetAuditOfficeLocationIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAuditOfficeLocationIdOk() (*int32, bool)`

GetAuditOfficeLocationIdOk returns a tuple with the AuditOfficeLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditOfficeLocationId

`func (o *AuditableEntitiesPutPreviousValues) SetAuditOfficeLocationId(v int32)`

SetAuditOfficeLocationId sets AuditOfficeLocationId field to given value.

### HasAuditOfficeLocationId

`func (o *AuditableEntitiesPutPreviousValues) HasAuditOfficeLocationId() bool`

HasAuditOfficeLocationId returns a boolean if a field has been set.

### GetAddress

`func (o *AuditableEntitiesPutPreviousValues) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AuditableEntitiesPutPreviousValues) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AuditableEntitiesPutPreviousValues) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AuditableEntitiesPutPreviousValues) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCity

`func (o *AuditableEntitiesPutPreviousValues) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AuditableEntitiesPutPreviousValues) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AuditableEntitiesPutPreviousValues) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AuditableEntitiesPutPreviousValues) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *AuditableEntitiesPutPreviousValues) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AuditableEntitiesPutPreviousValues) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AuditableEntitiesPutPreviousValues) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AuditableEntitiesPutPreviousValues) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *AuditableEntitiesPutPreviousValues) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AuditableEntitiesPutPreviousValues) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AuditableEntitiesPutPreviousValues) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AuditableEntitiesPutPreviousValues) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCustomText1

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText1() string`

GetCustomText1 returns the CustomText1 field if non-nil, zero value otherwise.

### GetCustomText1Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText1Ok() (*string, bool)`

GetCustomText1Ok returns a tuple with the CustomText1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText1

`func (o *AuditableEntitiesPutPreviousValues) SetCustomText1(v string)`

SetCustomText1 sets CustomText1 field to given value.

### HasCustomText1

`func (o *AuditableEntitiesPutPreviousValues) HasCustomText1() bool`

HasCustomText1 returns a boolean if a field has been set.

### GetCustomText2

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText2() string`

GetCustomText2 returns the CustomText2 field if non-nil, zero value otherwise.

### GetCustomText2Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText2Ok() (*string, bool)`

GetCustomText2Ok returns a tuple with the CustomText2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText2

`func (o *AuditableEntitiesPutPreviousValues) SetCustomText2(v string)`

SetCustomText2 sets CustomText2 field to given value.

### HasCustomText2

`func (o *AuditableEntitiesPutPreviousValues) HasCustomText2() bool`

HasCustomText2 returns a boolean if a field has been set.

### GetCustomText3

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText3() string`

GetCustomText3 returns the CustomText3 field if non-nil, zero value otherwise.

### GetCustomText3Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText3Ok() (*string, bool)`

GetCustomText3Ok returns a tuple with the CustomText3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText3

`func (o *AuditableEntitiesPutPreviousValues) SetCustomText3(v string)`

SetCustomText3 sets CustomText3 field to given value.

### HasCustomText3

`func (o *AuditableEntitiesPutPreviousValues) HasCustomText3() bool`

HasCustomText3 returns a boolean if a field has been set.

### GetCustomText4

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText4() string`

GetCustomText4 returns the CustomText4 field if non-nil, zero value otherwise.

### GetCustomText4Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText4Ok() (*string, bool)`

GetCustomText4Ok returns a tuple with the CustomText4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText4

`func (o *AuditableEntitiesPutPreviousValues) SetCustomText4(v string)`

SetCustomText4 sets CustomText4 field to given value.

### HasCustomText4

`func (o *AuditableEntitiesPutPreviousValues) HasCustomText4() bool`

HasCustomText4 returns a boolean if a field has been set.

### GetCustomText5

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText5() string`

GetCustomText5 returns the CustomText5 field if non-nil, zero value otherwise.

### GetCustomText5Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText5Ok() (*string, bool)`

GetCustomText5Ok returns a tuple with the CustomText5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText5

`func (o *AuditableEntitiesPutPreviousValues) SetCustomText5(v string)`

SetCustomText5 sets CustomText5 field to given value.

### HasCustomText5

`func (o *AuditableEntitiesPutPreviousValues) HasCustomText5() bool`

HasCustomText5 returns a boolean if a field has been set.

### GetCustomText6

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText6() string`

GetCustomText6 returns the CustomText6 field if non-nil, zero value otherwise.

### GetCustomText6Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText6Ok() (*string, bool)`

GetCustomText6Ok returns a tuple with the CustomText6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText6

`func (o *AuditableEntitiesPutPreviousValues) SetCustomText6(v string)`

SetCustomText6 sets CustomText6 field to given value.

### HasCustomText6

`func (o *AuditableEntitiesPutPreviousValues) HasCustomText6() bool`

HasCustomText6 returns a boolean if a field has been set.

### GetCustomText7

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText7() string`

GetCustomText7 returns the CustomText7 field if non-nil, zero value otherwise.

### GetCustomText7Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText7Ok() (*string, bool)`

GetCustomText7Ok returns a tuple with the CustomText7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText7

`func (o *AuditableEntitiesPutPreviousValues) SetCustomText7(v string)`

SetCustomText7 sets CustomText7 field to given value.

### HasCustomText7

`func (o *AuditableEntitiesPutPreviousValues) HasCustomText7() bool`

HasCustomText7 returns a boolean if a field has been set.

### GetCustomDate1

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDate1() string`

GetCustomDate1 returns the CustomDate1 field if non-nil, zero value otherwise.

### GetCustomDate1Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDate1Ok() (*string, bool)`

GetCustomDate1Ok returns a tuple with the CustomDate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate1

`func (o *AuditableEntitiesPutPreviousValues) SetCustomDate1(v string)`

SetCustomDate1 sets CustomDate1 field to given value.

### HasCustomDate1

`func (o *AuditableEntitiesPutPreviousValues) HasCustomDate1() bool`

HasCustomDate1 returns a boolean if a field has been set.

### GetAeCustomSelect1OptionId

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect1OptionId() int32`

GetAeCustomSelect1OptionId returns the AeCustomSelect1OptionId field if non-nil, zero value otherwise.

### GetAeCustomSelect1OptionIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect1OptionIdOk() (*int32, bool)`

GetAeCustomSelect1OptionIdOk returns a tuple with the AeCustomSelect1OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeCustomSelect1OptionId

`func (o *AuditableEntitiesPutPreviousValues) SetAeCustomSelect1OptionId(v int32)`

SetAeCustomSelect1OptionId sets AeCustomSelect1OptionId field to given value.

### HasAeCustomSelect1OptionId

`func (o *AuditableEntitiesPutPreviousValues) HasAeCustomSelect1OptionId() bool`

HasAeCustomSelect1OptionId returns a boolean if a field has been set.

### GetAuditDivisionId

`func (o *AuditableEntitiesPutPreviousValues) GetAuditDivisionId() int32`

GetAuditDivisionId returns the AuditDivisionId field if non-nil, zero value otherwise.

### GetAuditDivisionIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAuditDivisionIdOk() (*int32, bool)`

GetAuditDivisionIdOk returns a tuple with the AuditDivisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditDivisionId

`func (o *AuditableEntitiesPutPreviousValues) SetAuditDivisionId(v int32)`

SetAuditDivisionId sets AuditDivisionId field to given value.

### HasAuditDivisionId

`func (o *AuditableEntitiesPutPreviousValues) HasAuditDivisionId() bool`

HasAuditDivisionId returns a boolean if a field has been set.

### GetAuditBusinessSegmentId

`func (o *AuditableEntitiesPutPreviousValues) GetAuditBusinessSegmentId() int32`

GetAuditBusinessSegmentId returns the AuditBusinessSegmentId field if non-nil, zero value otherwise.

### GetAuditBusinessSegmentIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAuditBusinessSegmentIdOk() (*int32, bool)`

GetAuditBusinessSegmentIdOk returns a tuple with the AuditBusinessSegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditBusinessSegmentId

`func (o *AuditableEntitiesPutPreviousValues) SetAuditBusinessSegmentId(v int32)`

SetAuditBusinessSegmentId sets AuditBusinessSegmentId field to given value.

### HasAuditBusinessSegmentId

`func (o *AuditableEntitiesPutPreviousValues) HasAuditBusinessSegmentId() bool`

HasAuditBusinessSegmentId returns a boolean if a field has been set.

### GetAeCustomSelect2OptionId

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect2OptionId() int32`

GetAeCustomSelect2OptionId returns the AeCustomSelect2OptionId field if non-nil, zero value otherwise.

### GetAeCustomSelect2OptionIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect2OptionIdOk() (*int32, bool)`

GetAeCustomSelect2OptionIdOk returns a tuple with the AeCustomSelect2OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeCustomSelect2OptionId

`func (o *AuditableEntitiesPutPreviousValues) SetAeCustomSelect2OptionId(v int32)`

SetAeCustomSelect2OptionId sets AeCustomSelect2OptionId field to given value.

### HasAeCustomSelect2OptionId

`func (o *AuditableEntitiesPutPreviousValues) HasAeCustomSelect2OptionId() bool`

HasAeCustomSelect2OptionId returns a boolean if a field has been set.

### GetCustomDate2

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDate2() string`

GetCustomDate2 returns the CustomDate2 field if non-nil, zero value otherwise.

### GetCustomDate2Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDate2Ok() (*string, bool)`

GetCustomDate2Ok returns a tuple with the CustomDate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate2

`func (o *AuditableEntitiesPutPreviousValues) SetCustomDate2(v string)`

SetCustomDate2 sets CustomDate2 field to given value.

### HasCustomDate2

`func (o *AuditableEntitiesPutPreviousValues) HasCustomDate2() bool`

HasCustomDate2 returns a boolean if a field has been set.

### GetAssessmentId

`func (o *AuditableEntitiesPutPreviousValues) GetAssessmentId() int32`

GetAssessmentId returns the AssessmentId field if non-nil, zero value otherwise.

### GetAssessmentIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAssessmentIdOk() (*int32, bool)`

GetAssessmentIdOk returns a tuple with the AssessmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessmentId

`func (o *AuditableEntitiesPutPreviousValues) SetAssessmentId(v int32)`

SetAssessmentId sets AssessmentId field to given value.

### HasAssessmentId

`func (o *AuditableEntitiesPutPreviousValues) HasAssessmentId() bool`

HasAssessmentId returns a boolean if a field has been set.

### GetCustomText8

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText8() string`

GetCustomText8 returns the CustomText8 field if non-nil, zero value otherwise.

### GetCustomText8Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText8Ok() (*string, bool)`

GetCustomText8Ok returns a tuple with the CustomText8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText8

`func (o *AuditableEntitiesPutPreviousValues) SetCustomText8(v string)`

SetCustomText8 sets CustomText8 field to given value.

### HasCustomText8

`func (o *AuditableEntitiesPutPreviousValues) HasCustomText8() bool`

HasCustomText8 returns a boolean if a field has been set.

### GetCustomText9

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText9() string`

GetCustomText9 returns the CustomText9 field if non-nil, zero value otherwise.

### GetCustomText9Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText9Ok() (*string, bool)`

GetCustomText9Ok returns a tuple with the CustomText9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText9

`func (o *AuditableEntitiesPutPreviousValues) SetCustomText9(v string)`

SetCustomText9 sets CustomText9 field to given value.

### HasCustomText9

`func (o *AuditableEntitiesPutPreviousValues) HasCustomText9() bool`

HasCustomText9 returns a boolean if a field has been set.

### GetCustomText10

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText10() string`

GetCustomText10 returns the CustomText10 field if non-nil, zero value otherwise.

### GetCustomText10Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText10Ok() (*string, bool)`

GetCustomText10Ok returns a tuple with the CustomText10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText10

`func (o *AuditableEntitiesPutPreviousValues) SetCustomText10(v string)`

SetCustomText10 sets CustomText10 field to given value.

### HasCustomText10

`func (o *AuditableEntitiesPutPreviousValues) HasCustomText10() bool`

HasCustomText10 returns a boolean if a field has been set.

### GetCustomText11

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText11() string`

GetCustomText11 returns the CustomText11 field if non-nil, zero value otherwise.

### GetCustomText11Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText11Ok() (*string, bool)`

GetCustomText11Ok returns a tuple with the CustomText11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText11

`func (o *AuditableEntitiesPutPreviousValues) SetCustomText11(v string)`

SetCustomText11 sets CustomText11 field to given value.

### HasCustomText11

`func (o *AuditableEntitiesPutPreviousValues) HasCustomText11() bool`

HasCustomText11 returns a boolean if a field has been set.

### GetCustomText12

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText12() string`

GetCustomText12 returns the CustomText12 field if non-nil, zero value otherwise.

### GetCustomText12Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText12Ok() (*string, bool)`

GetCustomText12Ok returns a tuple with the CustomText12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText12

`func (o *AuditableEntitiesPutPreviousValues) SetCustomText12(v string)`

SetCustomText12 sets CustomText12 field to given value.

### HasCustomText12

`func (o *AuditableEntitiesPutPreviousValues) HasCustomText12() bool`

HasCustomText12 returns a boolean if a field has been set.

### GetCustomText13

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText13() string`

GetCustomText13 returns the CustomText13 field if non-nil, zero value otherwise.

### GetCustomText13Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText13Ok() (*string, bool)`

GetCustomText13Ok returns a tuple with the CustomText13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText13

`func (o *AuditableEntitiesPutPreviousValues) SetCustomText13(v string)`

SetCustomText13 sets CustomText13 field to given value.

### HasCustomText13

`func (o *AuditableEntitiesPutPreviousValues) HasCustomText13() bool`

HasCustomText13 returns a boolean if a field has been set.

### GetCustomText14

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText14() string`

GetCustomText14 returns the CustomText14 field if non-nil, zero value otherwise.

### GetCustomText14Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText14Ok() (*string, bool)`

GetCustomText14Ok returns a tuple with the CustomText14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText14

`func (o *AuditableEntitiesPutPreviousValues) SetCustomText14(v string)`

SetCustomText14 sets CustomText14 field to given value.

### HasCustomText14

`func (o *AuditableEntitiesPutPreviousValues) HasCustomText14() bool`

HasCustomText14 returns a boolean if a field has been set.

### GetCustomText15

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText15() string`

GetCustomText15 returns the CustomText15 field if non-nil, zero value otherwise.

### GetCustomText15Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomText15Ok() (*string, bool)`

GetCustomText15Ok returns a tuple with the CustomText15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText15

`func (o *AuditableEntitiesPutPreviousValues) SetCustomText15(v string)`

SetCustomText15 sets CustomText15 field to given value.

### HasCustomText15

`func (o *AuditableEntitiesPutPreviousValues) HasCustomText15() bool`

HasCustomText15 returns a boolean if a field has been set.

### GetReviewerUserId

`func (o *AuditableEntitiesPutPreviousValues) GetReviewerUserId() int32`

GetReviewerUserId returns the ReviewerUserId field if non-nil, zero value otherwise.

### GetReviewerUserIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetReviewerUserIdOk() (*int32, bool)`

GetReviewerUserIdOk returns a tuple with the ReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerUserId

`func (o *AuditableEntitiesPutPreviousValues) SetReviewerUserId(v int32)`

SetReviewerUserId sets ReviewerUserId field to given value.

### HasReviewerUserId

`func (o *AuditableEntitiesPutPreviousValues) HasReviewerUserId() bool`

HasReviewerUserId returns a boolean if a field has been set.

### GetAeCustomSelect3OptionId

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect3OptionId() int32`

GetAeCustomSelect3OptionId returns the AeCustomSelect3OptionId field if non-nil, zero value otherwise.

### GetAeCustomSelect3OptionIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect3OptionIdOk() (*int32, bool)`

GetAeCustomSelect3OptionIdOk returns a tuple with the AeCustomSelect3OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeCustomSelect3OptionId

`func (o *AuditableEntitiesPutPreviousValues) SetAeCustomSelect3OptionId(v int32)`

SetAeCustomSelect3OptionId sets AeCustomSelect3OptionId field to given value.

### HasAeCustomSelect3OptionId

`func (o *AuditableEntitiesPutPreviousValues) HasAeCustomSelect3OptionId() bool`

HasAeCustomSelect3OptionId returns a boolean if a field has been set.

### GetAeCustomSelect4OptionId

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect4OptionId() int32`

GetAeCustomSelect4OptionId returns the AeCustomSelect4OptionId field if non-nil, zero value otherwise.

### GetAeCustomSelect4OptionIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect4OptionIdOk() (*int32, bool)`

GetAeCustomSelect4OptionIdOk returns a tuple with the AeCustomSelect4OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeCustomSelect4OptionId

`func (o *AuditableEntitiesPutPreviousValues) SetAeCustomSelect4OptionId(v int32)`

SetAeCustomSelect4OptionId sets AeCustomSelect4OptionId field to given value.

### HasAeCustomSelect4OptionId

`func (o *AuditableEntitiesPutPreviousValues) HasAeCustomSelect4OptionId() bool`

HasAeCustomSelect4OptionId returns a boolean if a field has been set.

### GetAeCustomSelect5OptionId

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect5OptionId() int32`

GetAeCustomSelect5OptionId returns the AeCustomSelect5OptionId field if non-nil, zero value otherwise.

### GetAeCustomSelect5OptionIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect5OptionIdOk() (*int32, bool)`

GetAeCustomSelect5OptionIdOk returns a tuple with the AeCustomSelect5OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeCustomSelect5OptionId

`func (o *AuditableEntitiesPutPreviousValues) SetAeCustomSelect5OptionId(v int32)`

SetAeCustomSelect5OptionId sets AeCustomSelect5OptionId field to given value.

### HasAeCustomSelect5OptionId

`func (o *AuditableEntitiesPutPreviousValues) HasAeCustomSelect5OptionId() bool`

HasAeCustomSelect5OptionId returns a boolean if a field has been set.

### GetAeCustomSelect6OptionId

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect6OptionId() int32`

GetAeCustomSelect6OptionId returns the AeCustomSelect6OptionId field if non-nil, zero value otherwise.

### GetAeCustomSelect6OptionIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect6OptionIdOk() (*int32, bool)`

GetAeCustomSelect6OptionIdOk returns a tuple with the AeCustomSelect6OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeCustomSelect6OptionId

`func (o *AuditableEntitiesPutPreviousValues) SetAeCustomSelect6OptionId(v int32)`

SetAeCustomSelect6OptionId sets AeCustomSelect6OptionId field to given value.

### HasAeCustomSelect6OptionId

`func (o *AuditableEntitiesPutPreviousValues) HasAeCustomSelect6OptionId() bool`

HasAeCustomSelect6OptionId returns a boolean if a field has been set.

### GetAeCustomSelect7OptionId

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect7OptionId() int32`

GetAeCustomSelect7OptionId returns the AeCustomSelect7OptionId field if non-nil, zero value otherwise.

### GetAeCustomSelect7OptionIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect7OptionIdOk() (*int32, bool)`

GetAeCustomSelect7OptionIdOk returns a tuple with the AeCustomSelect7OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeCustomSelect7OptionId

`func (o *AuditableEntitiesPutPreviousValues) SetAeCustomSelect7OptionId(v int32)`

SetAeCustomSelect7OptionId sets AeCustomSelect7OptionId field to given value.

### HasAeCustomSelect7OptionId

`func (o *AuditableEntitiesPutPreviousValues) HasAeCustomSelect7OptionId() bool`

HasAeCustomSelect7OptionId returns a boolean if a field has been set.

### GetAeCustomSelect8OptionId

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect8OptionId() int32`

GetAeCustomSelect8OptionId returns the AeCustomSelect8OptionId field if non-nil, zero value otherwise.

### GetAeCustomSelect8OptionIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect8OptionIdOk() (*int32, bool)`

GetAeCustomSelect8OptionIdOk returns a tuple with the AeCustomSelect8OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeCustomSelect8OptionId

`func (o *AuditableEntitiesPutPreviousValues) SetAeCustomSelect8OptionId(v int32)`

SetAeCustomSelect8OptionId sets AeCustomSelect8OptionId field to given value.

### HasAeCustomSelect8OptionId

`func (o *AuditableEntitiesPutPreviousValues) HasAeCustomSelect8OptionId() bool`

HasAeCustomSelect8OptionId returns a boolean if a field has been set.

### GetAeStatusOptionId

`func (o *AuditableEntitiesPutPreviousValues) GetAeStatusOptionId() int32`

GetAeStatusOptionId returns the AeStatusOptionId field if non-nil, zero value otherwise.

### GetAeStatusOptionIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAeStatusOptionIdOk() (*int32, bool)`

GetAeStatusOptionIdOk returns a tuple with the AeStatusOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeStatusOptionId

`func (o *AuditableEntitiesPutPreviousValues) SetAeStatusOptionId(v int32)`

SetAeStatusOptionId sets AeStatusOptionId field to given value.

### HasAeStatusOptionId

`func (o *AuditableEntitiesPutPreviousValues) HasAeStatusOptionId() bool`

HasAeStatusOptionId returns a boolean if a field has been set.

### GetCustomDate3

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDate3() string`

GetCustomDate3 returns the CustomDate3 field if non-nil, zero value otherwise.

### GetCustomDate3Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDate3Ok() (*string, bool)`

GetCustomDate3Ok returns a tuple with the CustomDate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate3

`func (o *AuditableEntitiesPutPreviousValues) SetCustomDate3(v string)`

SetCustomDate3 sets CustomDate3 field to given value.

### HasCustomDate3

`func (o *AuditableEntitiesPutPreviousValues) HasCustomDate3() bool`

HasCustomDate3 returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *AuditableEntitiesPutPreviousValues) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *AuditableEntitiesPutPreviousValues) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *AuditableEntitiesPutPreviousValues) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *AuditableEntitiesPutPreviousValues) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetAeCustomSelect9OptionId

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect9OptionId() int32`

GetAeCustomSelect9OptionId returns the AeCustomSelect9OptionId field if non-nil, zero value otherwise.

### GetAeCustomSelect9OptionIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect9OptionIdOk() (*int32, bool)`

GetAeCustomSelect9OptionIdOk returns a tuple with the AeCustomSelect9OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeCustomSelect9OptionId

`func (o *AuditableEntitiesPutPreviousValues) SetAeCustomSelect9OptionId(v int32)`

SetAeCustomSelect9OptionId sets AeCustomSelect9OptionId field to given value.

### HasAeCustomSelect9OptionId

`func (o *AuditableEntitiesPutPreviousValues) HasAeCustomSelect9OptionId() bool`

HasAeCustomSelect9OptionId returns a boolean if a field has been set.

### GetAeCustomSelect10OptionId

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect10OptionId() int32`

GetAeCustomSelect10OptionId returns the AeCustomSelect10OptionId field if non-nil, zero value otherwise.

### GetAeCustomSelect10OptionIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect10OptionIdOk() (*int32, bool)`

GetAeCustomSelect10OptionIdOk returns a tuple with the AeCustomSelect10OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeCustomSelect10OptionId

`func (o *AuditableEntitiesPutPreviousValues) SetAeCustomSelect10OptionId(v int32)`

SetAeCustomSelect10OptionId sets AeCustomSelect10OptionId field to given value.

### HasAeCustomSelect10OptionId

`func (o *AuditableEntitiesPutPreviousValues) HasAeCustomSelect10OptionId() bool`

HasAeCustomSelect10OptionId returns a boolean if a field has been set.

### GetAeCustomSelect11OptionId

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect11OptionId() int32`

GetAeCustomSelect11OptionId returns the AeCustomSelect11OptionId field if non-nil, zero value otherwise.

### GetAeCustomSelect11OptionIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect11OptionIdOk() (*int32, bool)`

GetAeCustomSelect11OptionIdOk returns a tuple with the AeCustomSelect11OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeCustomSelect11OptionId

`func (o *AuditableEntitiesPutPreviousValues) SetAeCustomSelect11OptionId(v int32)`

SetAeCustomSelect11OptionId sets AeCustomSelect11OptionId field to given value.

### HasAeCustomSelect11OptionId

`func (o *AuditableEntitiesPutPreviousValues) HasAeCustomSelect11OptionId() bool`

HasAeCustomSelect11OptionId returns a boolean if a field has been set.

### GetAeCustomSelect12OptionId

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect12OptionId() int32`

GetAeCustomSelect12OptionId returns the AeCustomSelect12OptionId field if non-nil, zero value otherwise.

### GetAeCustomSelect12OptionIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect12OptionIdOk() (*int32, bool)`

GetAeCustomSelect12OptionIdOk returns a tuple with the AeCustomSelect12OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeCustomSelect12OptionId

`func (o *AuditableEntitiesPutPreviousValues) SetAeCustomSelect12OptionId(v int32)`

SetAeCustomSelect12OptionId sets AeCustomSelect12OptionId field to given value.

### HasAeCustomSelect12OptionId

`func (o *AuditableEntitiesPutPreviousValues) HasAeCustomSelect12OptionId() bool`

HasAeCustomSelect12OptionId returns a boolean if a field has been set.

### GetCustomDate4

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDate4() string`

GetCustomDate4 returns the CustomDate4 field if non-nil, zero value otherwise.

### GetCustomDate4Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDate4Ok() (*string, bool)`

GetCustomDate4Ok returns a tuple with the CustomDate4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate4

`func (o *AuditableEntitiesPutPreviousValues) SetCustomDate4(v string)`

SetCustomDate4 sets CustomDate4 field to given value.

### HasCustomDate4

`func (o *AuditableEntitiesPutPreviousValues) HasCustomDate4() bool`

HasCustomDate4 returns a boolean if a field has been set.

### GetExecutiveUserId

`func (o *AuditableEntitiesPutPreviousValues) GetExecutiveUserId() int32`

GetExecutiveUserId returns the ExecutiveUserId field if non-nil, zero value otherwise.

### GetExecutiveUserIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetExecutiveUserIdOk() (*int32, bool)`

GetExecutiveUserIdOk returns a tuple with the ExecutiveUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutiveUserId

`func (o *AuditableEntitiesPutPreviousValues) SetExecutiveUserId(v int32)`

SetExecutiveUserId sets ExecutiveUserId field to given value.

### HasExecutiveUserId

`func (o *AuditableEntitiesPutPreviousValues) HasExecutiveUserId() bool`

HasExecutiveUserId returns a boolean if a field has been set.

### GetVicePresidentUserId

`func (o *AuditableEntitiesPutPreviousValues) GetVicePresidentUserId() int32`

GetVicePresidentUserId returns the VicePresidentUserId field if non-nil, zero value otherwise.

### GetVicePresidentUserIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetVicePresidentUserIdOk() (*int32, bool)`

GetVicePresidentUserIdOk returns a tuple with the VicePresidentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVicePresidentUserId

`func (o *AuditableEntitiesPutPreviousValues) SetVicePresidentUserId(v int32)`

SetVicePresidentUserId sets VicePresidentUserId field to given value.

### HasVicePresidentUserId

`func (o *AuditableEntitiesPutPreviousValues) HasVicePresidentUserId() bool`

HasVicePresidentUserId returns a boolean if a field has been set.

### GetInherentRisk

`func (o *AuditableEntitiesPutPreviousValues) GetInherentRisk() interface{}`

GetInherentRisk returns the InherentRisk field if non-nil, zero value otherwise.

### GetInherentRiskOk

`func (o *AuditableEntitiesPutPreviousValues) GetInherentRiskOk() (*interface{}, bool)`

GetInherentRiskOk returns a tuple with the InherentRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInherentRisk

`func (o *AuditableEntitiesPutPreviousValues) SetInherentRisk(v interface{})`

SetInherentRisk sets InherentRisk field to given value.

### HasInherentRisk

`func (o *AuditableEntitiesPutPreviousValues) HasInherentRisk() bool`

HasInherentRisk returns a boolean if a field has been set.

### SetInherentRiskNil

`func (o *AuditableEntitiesPutPreviousValues) SetInherentRiskNil(b bool)`

 SetInherentRiskNil sets the value for InherentRisk to be an explicit nil

### UnsetInherentRisk
`func (o *AuditableEntitiesPutPreviousValues) UnsetInherentRisk()`

UnsetInherentRisk ensures that no value is present for InherentRisk, not even an explicit nil
### GetResidualRisk

`func (o *AuditableEntitiesPutPreviousValues) GetResidualRisk() interface{}`

GetResidualRisk returns the ResidualRisk field if non-nil, zero value otherwise.

### GetResidualRiskOk

`func (o *AuditableEntitiesPutPreviousValues) GetResidualRiskOk() (*interface{}, bool)`

GetResidualRiskOk returns a tuple with the ResidualRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidualRisk

`func (o *AuditableEntitiesPutPreviousValues) SetResidualRisk(v interface{})`

SetResidualRisk sets ResidualRisk field to given value.

### HasResidualRisk

`func (o *AuditableEntitiesPutPreviousValues) HasResidualRisk() bool`

HasResidualRisk returns a boolean if a field has been set.

### SetResidualRiskNil

`func (o *AuditableEntitiesPutPreviousValues) SetResidualRiskNil(b bool)`

 SetResidualRiskNil sets the value for ResidualRisk to be an explicit nil

### UnsetResidualRisk
`func (o *AuditableEntitiesPutPreviousValues) UnsetResidualRisk()`

UnsetResidualRisk ensures that no value is present for ResidualRisk, not even an explicit nil
### GetManagerUserId

`func (o *AuditableEntitiesPutPreviousValues) GetManagerUserId() int32`

GetManagerUserId returns the ManagerUserId field if non-nil, zero value otherwise.

### GetManagerUserIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetManagerUserIdOk() (*int32, bool)`

GetManagerUserIdOk returns a tuple with the ManagerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerUserId

`func (o *AuditableEntitiesPutPreviousValues) SetManagerUserId(v int32)`

SetManagerUserId sets ManagerUserId field to given value.

### HasManagerUserId

`func (o *AuditableEntitiesPutPreviousValues) HasManagerUserId() bool`

HasManagerUserId returns a boolean if a field has been set.

### GetCoordinatorUserId

`func (o *AuditableEntitiesPutPreviousValues) GetCoordinatorUserId() int32`

GetCoordinatorUserId returns the CoordinatorUserId field if non-nil, zero value otherwise.

### GetCoordinatorUserIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetCoordinatorUserIdOk() (*int32, bool)`

GetCoordinatorUserIdOk returns a tuple with the CoordinatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinatorUserId

`func (o *AuditableEntitiesPutPreviousValues) SetCoordinatorUserId(v int32)`

SetCoordinatorUserId sets CoordinatorUserId field to given value.

### HasCoordinatorUserId

`func (o *AuditableEntitiesPutPreviousValues) HasCoordinatorUserId() bool`

HasCoordinatorUserId returns a boolean if a field has been set.

### GetProductStatusId

`func (o *AuditableEntitiesPutPreviousValues) GetProductStatusId() int32`

GetProductStatusId returns the ProductStatusId field if non-nil, zero value otherwise.

### GetProductStatusIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetProductStatusIdOk() (*int32, bool)`

GetProductStatusIdOk returns a tuple with the ProductStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductStatusId

`func (o *AuditableEntitiesPutPreviousValues) SetProductStatusId(v int32)`

SetProductStatusId sets ProductStatusId field to given value.

### HasProductStatusId

`func (o *AuditableEntitiesPutPreviousValues) HasProductStatusId() bool`

HasProductStatusId returns a boolean if a field has been set.

### GetVendorCriticalityId

`func (o *AuditableEntitiesPutPreviousValues) GetVendorCriticalityId() int32`

GetVendorCriticalityId returns the VendorCriticalityId field if non-nil, zero value otherwise.

### GetVendorCriticalityIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetVendorCriticalityIdOk() (*int32, bool)`

GetVendorCriticalityIdOk returns a tuple with the VendorCriticalityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCriticalityId

`func (o *AuditableEntitiesPutPreviousValues) SetVendorCriticalityId(v int32)`

SetVendorCriticalityId sets VendorCriticalityId field to given value.

### HasVendorCriticalityId

`func (o *AuditableEntitiesPutPreviousValues) HasVendorCriticalityId() bool`

HasVendorCriticalityId returns a boolean if a field has been set.

### GetLegalEntityTypeId

`func (o *AuditableEntitiesPutPreviousValues) GetLegalEntityTypeId() int32`

GetLegalEntityTypeId returns the LegalEntityTypeId field if non-nil, zero value otherwise.

### GetLegalEntityTypeIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetLegalEntityTypeIdOk() (*int32, bool)`

GetLegalEntityTypeIdOk returns a tuple with the LegalEntityTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalEntityTypeId

`func (o *AuditableEntitiesPutPreviousValues) SetLegalEntityTypeId(v int32)`

SetLegalEntityTypeId sets LegalEntityTypeId field to given value.

### HasLegalEntityTypeId

`func (o *AuditableEntitiesPutPreviousValues) HasLegalEntityTypeId() bool`

HasLegalEntityTypeId returns a boolean if a field has been set.

### GetItAssetTypeId

`func (o *AuditableEntitiesPutPreviousValues) GetItAssetTypeId() int32`

GetItAssetTypeId returns the ItAssetTypeId field if non-nil, zero value otherwise.

### GetItAssetTypeIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetItAssetTypeIdOk() (*int32, bool)`

GetItAssetTypeIdOk returns a tuple with the ItAssetTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItAssetTypeId

`func (o *AuditableEntitiesPutPreviousValues) SetItAssetTypeId(v int32)`

SetItAssetTypeId sets ItAssetTypeId field to given value.

### HasItAssetTypeId

`func (o *AuditableEntitiesPutPreviousValues) HasItAssetTypeId() bool`

HasItAssetTypeId returns a boolean if a field has been set.

### GetItAssetStatusId

`func (o *AuditableEntitiesPutPreviousValues) GetItAssetStatusId() int32`

GetItAssetStatusId returns the ItAssetStatusId field if non-nil, zero value otherwise.

### GetItAssetStatusIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetItAssetStatusIdOk() (*int32, bool)`

GetItAssetStatusIdOk returns a tuple with the ItAssetStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItAssetStatusId

`func (o *AuditableEntitiesPutPreviousValues) SetItAssetStatusId(v int32)`

SetItAssetStatusId sets ItAssetStatusId field to given value.

### HasItAssetStatusId

`func (o *AuditableEntitiesPutPreviousValues) HasItAssetStatusId() bool`

HasItAssetStatusId returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *AuditableEntitiesPutPreviousValues) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *AuditableEntitiesPutPreviousValues) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *AuditableEntitiesPutPreviousValues) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *AuditableEntitiesPutPreviousValues) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetContractDetails

`func (o *AuditableEntitiesPutPreviousValues) GetContractDetails() string`

GetContractDetails returns the ContractDetails field if non-nil, zero value otherwise.

### GetContractDetailsOk

`func (o *AuditableEntitiesPutPreviousValues) GetContractDetailsOk() (*string, bool)`

GetContractDetailsOk returns a tuple with the ContractDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractDetails

`func (o *AuditableEntitiesPutPreviousValues) SetContractDetails(v string)`

SetContractDetails sets ContractDetails field to given value.

### HasContractDetails

`func (o *AuditableEntitiesPutPreviousValues) HasContractDetails() bool`

HasContractDetails returns a boolean if a field has been set.

### GetContactName

`func (o *AuditableEntitiesPutPreviousValues) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *AuditableEntitiesPutPreviousValues) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *AuditableEntitiesPutPreviousValues) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *AuditableEntitiesPutPreviousValues) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetContactEmail

`func (o *AuditableEntitiesPutPreviousValues) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *AuditableEntitiesPutPreviousValues) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *AuditableEntitiesPutPreviousValues) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *AuditableEntitiesPutPreviousValues) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetContactPhoneNumber

`func (o *AuditableEntitiesPutPreviousValues) GetContactPhoneNumber() string`

GetContactPhoneNumber returns the ContactPhoneNumber field if non-nil, zero value otherwise.

### GetContactPhoneNumberOk

`func (o *AuditableEntitiesPutPreviousValues) GetContactPhoneNumberOk() (*string, bool)`

GetContactPhoneNumberOk returns a tuple with the ContactPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPhoneNumber

`func (o *AuditableEntitiesPutPreviousValues) SetContactPhoneNumber(v string)`

SetContactPhoneNumber sets ContactPhoneNumber field to given value.

### HasContactPhoneNumber

`func (o *AuditableEntitiesPutPreviousValues) HasContactPhoneNumber() bool`

HasContactPhoneNumber returns a boolean if a field has been set.

### GetItServiceLevelAgreement

`func (o *AuditableEntitiesPutPreviousValues) GetItServiceLevelAgreement() string`

GetItServiceLevelAgreement returns the ItServiceLevelAgreement field if non-nil, zero value otherwise.

### GetItServiceLevelAgreementOk

`func (o *AuditableEntitiesPutPreviousValues) GetItServiceLevelAgreementOk() (*string, bool)`

GetItServiceLevelAgreementOk returns a tuple with the ItServiceLevelAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItServiceLevelAgreement

`func (o *AuditableEntitiesPutPreviousValues) SetItServiceLevelAgreement(v string)`

SetItServiceLevelAgreement sets ItServiceLevelAgreement field to given value.

### HasItServiceLevelAgreement

`func (o *AuditableEntitiesPutPreviousValues) HasItServiceLevelAgreement() bool`

HasItServiceLevelAgreement returns a boolean if a field has been set.

### GetItRecoveryTimeObjective

`func (o *AuditableEntitiesPutPreviousValues) GetItRecoveryTimeObjective() string`

GetItRecoveryTimeObjective returns the ItRecoveryTimeObjective field if non-nil, zero value otherwise.

### GetItRecoveryTimeObjectiveOk

`func (o *AuditableEntitiesPutPreviousValues) GetItRecoveryTimeObjectiveOk() (*string, bool)`

GetItRecoveryTimeObjectiveOk returns a tuple with the ItRecoveryTimeObjective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItRecoveryTimeObjective

`func (o *AuditableEntitiesPutPreviousValues) SetItRecoveryTimeObjective(v string)`

SetItRecoveryTimeObjective sets ItRecoveryTimeObjective field to given value.

### HasItRecoveryTimeObjective

`func (o *AuditableEntitiesPutPreviousValues) HasItRecoveryTimeObjective() bool`

HasItRecoveryTimeObjective returns a boolean if a field has been set.

### GetItRecoveryPointObjective

`func (o *AuditableEntitiesPutPreviousValues) GetItRecoveryPointObjective() string`

GetItRecoveryPointObjective returns the ItRecoveryPointObjective field if non-nil, zero value otherwise.

### GetItRecoveryPointObjectiveOk

`func (o *AuditableEntitiesPutPreviousValues) GetItRecoveryPointObjectiveOk() (*string, bool)`

GetItRecoveryPointObjectiveOk returns a tuple with the ItRecoveryPointObjective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItRecoveryPointObjective

`func (o *AuditableEntitiesPutPreviousValues) SetItRecoveryPointObjective(v string)`

SetItRecoveryPointObjective sets ItRecoveryPointObjective field to given value.

### HasItRecoveryPointObjective

`func (o *AuditableEntitiesPutPreviousValues) HasItRecoveryPointObjective() bool`

HasItRecoveryPointObjective returns a boolean if a field has been set.

### GetOldItSecurityIncidentHours

`func (o *AuditableEntitiesPutPreviousValues) GetOldItSecurityIncidentHours() float32`

GetOldItSecurityIncidentHours returns the OldItSecurityIncidentHours field if non-nil, zero value otherwise.

### GetOldItSecurityIncidentHoursOk

`func (o *AuditableEntitiesPutPreviousValues) GetOldItSecurityIncidentHoursOk() (*float32, bool)`

GetOldItSecurityIncidentHoursOk returns a tuple with the OldItSecurityIncidentHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldItSecurityIncidentHours

`func (o *AuditableEntitiesPutPreviousValues) SetOldItSecurityIncidentHours(v float32)`

SetOldItSecurityIncidentHours sets OldItSecurityIncidentHours field to given value.

### HasOldItSecurityIncidentHours

`func (o *AuditableEntitiesPutPreviousValues) HasOldItSecurityIncidentHours() bool`

HasOldItSecurityIncidentHours returns a boolean if a field has been set.

### GetEmployeesSize

`func (o *AuditableEntitiesPutPreviousValues) GetEmployeesSize() float32`

GetEmployeesSize returns the EmployeesSize field if non-nil, zero value otherwise.

### GetEmployeesSizeOk

`func (o *AuditableEntitiesPutPreviousValues) GetEmployeesSizeOk() (*float32, bool)`

GetEmployeesSizeOk returns a tuple with the EmployeesSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeesSize

`func (o *AuditableEntitiesPutPreviousValues) SetEmployeesSize(v float32)`

SetEmployeesSize sets EmployeesSize field to given value.

### HasEmployeesSize

`func (o *AuditableEntitiesPutPreviousValues) HasEmployeesSize() bool`

HasEmployeesSize returns a boolean if a field has been set.

### GetRegulatoryAuditFrequencyId

`func (o *AuditableEntitiesPutPreviousValues) GetRegulatoryAuditFrequencyId() int32`

GetRegulatoryAuditFrequencyId returns the RegulatoryAuditFrequencyId field if non-nil, zero value otherwise.

### GetRegulatoryAuditFrequencyIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetRegulatoryAuditFrequencyIdOk() (*int32, bool)`

GetRegulatoryAuditFrequencyIdOk returns a tuple with the RegulatoryAuditFrequencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryAuditFrequencyId

`func (o *AuditableEntitiesPutPreviousValues) SetRegulatoryAuditFrequencyId(v int32)`

SetRegulatoryAuditFrequencyId sets RegulatoryAuditFrequencyId field to given value.

### HasRegulatoryAuditFrequencyId

`func (o *AuditableEntitiesPutPreviousValues) HasRegulatoryAuditFrequencyId() bool`

HasRegulatoryAuditFrequencyId returns a boolean if a field has been set.

### GetRegulatoryAuditFrequencyNotes

`func (o *AuditableEntitiesPutPreviousValues) GetRegulatoryAuditFrequencyNotes() string`

GetRegulatoryAuditFrequencyNotes returns the RegulatoryAuditFrequencyNotes field if non-nil, zero value otherwise.

### GetRegulatoryAuditFrequencyNotesOk

`func (o *AuditableEntitiesPutPreviousValues) GetRegulatoryAuditFrequencyNotesOk() (*string, bool)`

GetRegulatoryAuditFrequencyNotesOk returns a tuple with the RegulatoryAuditFrequencyNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryAuditFrequencyNotes

`func (o *AuditableEntitiesPutPreviousValues) SetRegulatoryAuditFrequencyNotes(v string)`

SetRegulatoryAuditFrequencyNotes sets RegulatoryAuditFrequencyNotes field to given value.

### HasRegulatoryAuditFrequencyNotes

`func (o *AuditableEntitiesPutPreviousValues) HasRegulatoryAuditFrequencyNotes() bool`

HasRegulatoryAuditFrequencyNotes returns a boolean if a field has been set.

### GetCalculatedAuditFrequencyId

`func (o *AuditableEntitiesPutPreviousValues) GetCalculatedAuditFrequencyId() int32`

GetCalculatedAuditFrequencyId returns the CalculatedAuditFrequencyId field if non-nil, zero value otherwise.

### GetCalculatedAuditFrequencyIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetCalculatedAuditFrequencyIdOk() (*int32, bool)`

GetCalculatedAuditFrequencyIdOk returns a tuple with the CalculatedAuditFrequencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedAuditFrequencyId

`func (o *AuditableEntitiesPutPreviousValues) SetCalculatedAuditFrequencyId(v int32)`

SetCalculatedAuditFrequencyId sets CalculatedAuditFrequencyId field to given value.

### HasCalculatedAuditFrequencyId

`func (o *AuditableEntitiesPutPreviousValues) HasCalculatedAuditFrequencyId() bool`

HasCalculatedAuditFrequencyId returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *AuditableEntitiesPutPreviousValues) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *AuditableEntitiesPutPreviousValues) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *AuditableEntitiesPutPreviousValues) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCustomDecimal1

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDecimal1() float32`

GetCustomDecimal1 returns the CustomDecimal1 field if non-nil, zero value otherwise.

### GetCustomDecimal1Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDecimal1Ok() (*float32, bool)`

GetCustomDecimal1Ok returns a tuple with the CustomDecimal1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDecimal1

`func (o *AuditableEntitiesPutPreviousValues) SetCustomDecimal1(v float32)`

SetCustomDecimal1 sets CustomDecimal1 field to given value.

### HasCustomDecimal1

`func (o *AuditableEntitiesPutPreviousValues) HasCustomDecimal1() bool`

HasCustomDecimal1 returns a boolean if a field has been set.

### GetCustomDecimal2

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDecimal2() float32`

GetCustomDecimal2 returns the CustomDecimal2 field if non-nil, zero value otherwise.

### GetCustomDecimal2Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDecimal2Ok() (*float32, bool)`

GetCustomDecimal2Ok returns a tuple with the CustomDecimal2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDecimal2

`func (o *AuditableEntitiesPutPreviousValues) SetCustomDecimal2(v float32)`

SetCustomDecimal2 sets CustomDecimal2 field to given value.

### HasCustomDecimal2

`func (o *AuditableEntitiesPutPreviousValues) HasCustomDecimal2() bool`

HasCustomDecimal2 returns a boolean if a field has been set.

### GetCustomDecimal3

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDecimal3() float32`

GetCustomDecimal3 returns the CustomDecimal3 field if non-nil, zero value otherwise.

### GetCustomDecimal3Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDecimal3Ok() (*float32, bool)`

GetCustomDecimal3Ok returns a tuple with the CustomDecimal3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDecimal3

`func (o *AuditableEntitiesPutPreviousValues) SetCustomDecimal3(v float32)`

SetCustomDecimal3 sets CustomDecimal3 field to given value.

### HasCustomDecimal3

`func (o *AuditableEntitiesPutPreviousValues) HasCustomDecimal3() bool`

HasCustomDecimal3 returns a boolean if a field has been set.

### GetCustomDecimal4

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDecimal4() float32`

GetCustomDecimal4 returns the CustomDecimal4 field if non-nil, zero value otherwise.

### GetCustomDecimal4Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDecimal4Ok() (*float32, bool)`

GetCustomDecimal4Ok returns a tuple with the CustomDecimal4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDecimal4

`func (o *AuditableEntitiesPutPreviousValues) SetCustomDecimal4(v float32)`

SetCustomDecimal4 sets CustomDecimal4 field to given value.

### HasCustomDecimal4

`func (o *AuditableEntitiesPutPreviousValues) HasCustomDecimal4() bool`

HasCustomDecimal4 returns a boolean if a field has been set.

### GetCustomDecimal5

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDecimal5() float32`

GetCustomDecimal5 returns the CustomDecimal5 field if non-nil, zero value otherwise.

### GetCustomDecimal5Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDecimal5Ok() (*float32, bool)`

GetCustomDecimal5Ok returns a tuple with the CustomDecimal5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDecimal5

`func (o *AuditableEntitiesPutPreviousValues) SetCustomDecimal5(v float32)`

SetCustomDecimal5 sets CustomDecimal5 field to given value.

### HasCustomDecimal5

`func (o *AuditableEntitiesPutPreviousValues) HasCustomDecimal5() bool`

HasCustomDecimal5 returns a boolean if a field has been set.

### GetCustomDecimal6

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDecimal6() float32`

GetCustomDecimal6 returns the CustomDecimal6 field if non-nil, zero value otherwise.

### GetCustomDecimal6Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDecimal6Ok() (*float32, bool)`

GetCustomDecimal6Ok returns a tuple with the CustomDecimal6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDecimal6

`func (o *AuditableEntitiesPutPreviousValues) SetCustomDecimal6(v float32)`

SetCustomDecimal6 sets CustomDecimal6 field to given value.

### HasCustomDecimal6

`func (o *AuditableEntitiesPutPreviousValues) HasCustomDecimal6() bool`

HasCustomDecimal6 returns a boolean if a field has been set.

### GetCustomDecimal7

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDecimal7() float32`

GetCustomDecimal7 returns the CustomDecimal7 field if non-nil, zero value otherwise.

### GetCustomDecimal7Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDecimal7Ok() (*float32, bool)`

GetCustomDecimal7Ok returns a tuple with the CustomDecimal7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDecimal7

`func (o *AuditableEntitiesPutPreviousValues) SetCustomDecimal7(v float32)`

SetCustomDecimal7 sets CustomDecimal7 field to given value.

### HasCustomDecimal7

`func (o *AuditableEntitiesPutPreviousValues) HasCustomDecimal7() bool`

HasCustomDecimal7 returns a boolean if a field has been set.

### GetCustomDecimal8

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDecimal8() float32`

GetCustomDecimal8 returns the CustomDecimal8 field if non-nil, zero value otherwise.

### GetCustomDecimal8Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDecimal8Ok() (*float32, bool)`

GetCustomDecimal8Ok returns a tuple with the CustomDecimal8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDecimal8

`func (o *AuditableEntitiesPutPreviousValues) SetCustomDecimal8(v float32)`

SetCustomDecimal8 sets CustomDecimal8 field to given value.

### HasCustomDecimal8

`func (o *AuditableEntitiesPutPreviousValues) HasCustomDecimal8() bool`

HasCustomDecimal8 returns a boolean if a field has been set.

### GetCustomCurrency1

`func (o *AuditableEntitiesPutPreviousValues) GetCustomCurrency1() float32`

GetCustomCurrency1 returns the CustomCurrency1 field if non-nil, zero value otherwise.

### GetCustomCurrency1Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomCurrency1Ok() (*float32, bool)`

GetCustomCurrency1Ok returns a tuple with the CustomCurrency1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCurrency1

`func (o *AuditableEntitiesPutPreviousValues) SetCustomCurrency1(v float32)`

SetCustomCurrency1 sets CustomCurrency1 field to given value.

### HasCustomCurrency1

`func (o *AuditableEntitiesPutPreviousValues) HasCustomCurrency1() bool`

HasCustomCurrency1 returns a boolean if a field has been set.

### GetCustomCurrency2

`func (o *AuditableEntitiesPutPreviousValues) GetCustomCurrency2() float32`

GetCustomCurrency2 returns the CustomCurrency2 field if non-nil, zero value otherwise.

### GetCustomCurrency2Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomCurrency2Ok() (*float32, bool)`

GetCustomCurrency2Ok returns a tuple with the CustomCurrency2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCurrency2

`func (o *AuditableEntitiesPutPreviousValues) SetCustomCurrency2(v float32)`

SetCustomCurrency2 sets CustomCurrency2 field to given value.

### HasCustomCurrency2

`func (o *AuditableEntitiesPutPreviousValues) HasCustomCurrency2() bool`

HasCustomCurrency2 returns a boolean if a field has been set.

### GetCustomCurrency3

`func (o *AuditableEntitiesPutPreviousValues) GetCustomCurrency3() float32`

GetCustomCurrency3 returns the CustomCurrency3 field if non-nil, zero value otherwise.

### GetCustomCurrency3Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomCurrency3Ok() (*float32, bool)`

GetCustomCurrency3Ok returns a tuple with the CustomCurrency3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCurrency3

`func (o *AuditableEntitiesPutPreviousValues) SetCustomCurrency3(v float32)`

SetCustomCurrency3 sets CustomCurrency3 field to given value.

### HasCustomCurrency3

`func (o *AuditableEntitiesPutPreviousValues) HasCustomCurrency3() bool`

HasCustomCurrency3 returns a boolean if a field has been set.

### GetCustomCurrency4

`func (o *AuditableEntitiesPutPreviousValues) GetCustomCurrency4() float32`

GetCustomCurrency4 returns the CustomCurrency4 field if non-nil, zero value otherwise.

### GetCustomCurrency4Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomCurrency4Ok() (*float32, bool)`

GetCustomCurrency4Ok returns a tuple with the CustomCurrency4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCurrency4

`func (o *AuditableEntitiesPutPreviousValues) SetCustomCurrency4(v float32)`

SetCustomCurrency4 sets CustomCurrency4 field to given value.

### HasCustomCurrency4

`func (o *AuditableEntitiesPutPreviousValues) HasCustomCurrency4() bool`

HasCustomCurrency4 returns a boolean if a field has been set.

### GetFieldData

`func (o *AuditableEntitiesPutPreviousValues) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *AuditableEntitiesPutPreviousValues) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *AuditableEntitiesPutPreviousValues) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *AuditableEntitiesPutPreviousValues) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *AuditableEntitiesPutPreviousValues) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *AuditableEntitiesPutPreviousValues) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil
### GetDefaultOpsAuditTemplateId

`func (o *AuditableEntitiesPutPreviousValues) GetDefaultOpsAuditTemplateId() int32`

GetDefaultOpsAuditTemplateId returns the DefaultOpsAuditTemplateId field if non-nil, zero value otherwise.

### GetDefaultOpsAuditTemplateIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetDefaultOpsAuditTemplateIdOk() (*int32, bool)`

GetDefaultOpsAuditTemplateIdOk returns a tuple with the DefaultOpsAuditTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOpsAuditTemplateId

`func (o *AuditableEntitiesPutPreviousValues) SetDefaultOpsAuditTemplateId(v int32)`

SetDefaultOpsAuditTemplateId sets DefaultOpsAuditTemplateId field to given value.

### HasDefaultOpsAuditTemplateId

`func (o *AuditableEntitiesPutPreviousValues) HasDefaultOpsAuditTemplateId() bool`

HasDefaultOpsAuditTemplateId returns a boolean if a field has been set.

### GetRiskScoreId

`func (o *AuditableEntitiesPutPreviousValues) GetRiskScoreId() int32`

GetRiskScoreId returns the RiskScoreId field if non-nil, zero value otherwise.

### GetRiskScoreIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetRiskScoreIdOk() (*int32, bool)`

GetRiskScoreIdOk returns a tuple with the RiskScoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScoreId

`func (o *AuditableEntitiesPutPreviousValues) SetRiskScoreId(v int32)`

SetRiskScoreId sets RiskScoreId field to given value.

### HasRiskScoreId

`func (o *AuditableEntitiesPutPreviousValues) HasRiskScoreId() bool`

HasRiskScoreId returns a boolean if a field has been set.

### GetVendorStatus

`func (o *AuditableEntitiesPutPreviousValues) GetVendorStatus() string`

GetVendorStatus returns the VendorStatus field if non-nil, zero value otherwise.

### GetVendorStatusOk

`func (o *AuditableEntitiesPutPreviousValues) GetVendorStatusOk() (*string, bool)`

GetVendorStatusOk returns a tuple with the VendorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorStatus

`func (o *AuditableEntitiesPutPreviousValues) SetVendorStatus(v string)`

SetVendorStatus sets VendorStatus field to given value.

### HasVendorStatus

`func (o *AuditableEntitiesPutPreviousValues) HasVendorStatus() bool`

HasVendorStatus returns a boolean if a field has been set.

### GetContactUserId

`func (o *AuditableEntitiesPutPreviousValues) GetContactUserId() int32`

GetContactUserId returns the ContactUserId field if non-nil, zero value otherwise.

### GetContactUserIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetContactUserIdOk() (*int32, bool)`

GetContactUserIdOk returns a tuple with the ContactUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactUserId

`func (o *AuditableEntitiesPutPreviousValues) SetContactUserId(v int32)`

SetContactUserId sets ContactUserId field to given value.

### HasContactUserId

`func (o *AuditableEntitiesPutPreviousValues) HasContactUserId() bool`

HasContactUserId returns a boolean if a field has been set.

### GetPriorAuditId

`func (o *AuditableEntitiesPutPreviousValues) GetPriorAuditId() int32`

GetPriorAuditId returns the PriorAuditId field if non-nil, zero value otherwise.

### GetPriorAuditIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetPriorAuditIdOk() (*int32, bool)`

GetPriorAuditIdOk returns a tuple with the PriorAuditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorAuditId

`func (o *AuditableEntitiesPutPreviousValues) SetPriorAuditId(v int32)`

SetPriorAuditId sets PriorAuditId field to given value.

### HasPriorAuditId

`func (o *AuditableEntitiesPutPreviousValues) HasPriorAuditId() bool`

HasPriorAuditId returns a boolean if a field has been set.

### GetUpcomingAuditId

`func (o *AuditableEntitiesPutPreviousValues) GetUpcomingAuditId() int32`

GetUpcomingAuditId returns the UpcomingAuditId field if non-nil, zero value otherwise.

### GetUpcomingAuditIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetUpcomingAuditIdOk() (*int32, bool)`

GetUpcomingAuditIdOk returns a tuple with the UpcomingAuditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcomingAuditId

`func (o *AuditableEntitiesPutPreviousValues) SetUpcomingAuditId(v int32)`

SetUpcomingAuditId sets UpcomingAuditId field to given value.

### HasUpcomingAuditId

`func (o *AuditableEntitiesPutPreviousValues) HasUpcomingAuditId() bool`

HasUpcomingAuditId returns a boolean if a field has been set.

### GetPriorAuditFinalReportDate

`func (o *AuditableEntitiesPutPreviousValues) GetPriorAuditFinalReportDate() string`

GetPriorAuditFinalReportDate returns the PriorAuditFinalReportDate field if non-nil, zero value otherwise.

### GetPriorAuditFinalReportDateOk

`func (o *AuditableEntitiesPutPreviousValues) GetPriorAuditFinalReportDateOk() (*string, bool)`

GetPriorAuditFinalReportDateOk returns a tuple with the PriorAuditFinalReportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorAuditFinalReportDate

`func (o *AuditableEntitiesPutPreviousValues) SetPriorAuditFinalReportDate(v string)`

SetPriorAuditFinalReportDate sets PriorAuditFinalReportDate field to given value.

### HasPriorAuditFinalReportDate

`func (o *AuditableEntitiesPutPreviousValues) HasPriorAuditFinalReportDate() bool`

HasPriorAuditFinalReportDate returns a boolean if a field has been set.

### GetFormulaData

`func (o *AuditableEntitiesPutPreviousValues) GetFormulaData() interface{}`

GetFormulaData returns the FormulaData field if non-nil, zero value otherwise.

### GetFormulaDataOk

`func (o *AuditableEntitiesPutPreviousValues) GetFormulaDataOk() (*interface{}, bool)`

GetFormulaDataOk returns a tuple with the FormulaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulaData

`func (o *AuditableEntitiesPutPreviousValues) SetFormulaData(v interface{})`

SetFormulaData sets FormulaData field to given value.

### HasFormulaData

`func (o *AuditableEntitiesPutPreviousValues) HasFormulaData() bool`

HasFormulaData returns a boolean if a field has been set.

### SetFormulaDataNil

`func (o *AuditableEntitiesPutPreviousValues) SetFormulaDataNil(b bool)`

 SetFormulaDataNil sets the value for FormulaData to be an explicit nil

### UnsetFormulaData
`func (o *AuditableEntitiesPutPreviousValues) UnsetFormulaData()`

UnsetFormulaData ensures that no value is present for FormulaData, not even an explicit nil
### GetIsFlagged

`func (o *AuditableEntitiesPutPreviousValues) GetIsFlagged() bool`

GetIsFlagged returns the IsFlagged field if non-nil, zero value otherwise.

### GetIsFlaggedOk

`func (o *AuditableEntitiesPutPreviousValues) GetIsFlaggedOk() (*bool, bool)`

GetIsFlaggedOk returns a tuple with the IsFlagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlagged

`func (o *AuditableEntitiesPutPreviousValues) SetIsFlagged(v bool)`

SetIsFlagged sets IsFlagged field to given value.

### HasIsFlagged

`func (o *AuditableEntitiesPutPreviousValues) HasIsFlagged() bool`

HasIsFlagged returns a boolean if a field has been set.

### GetRiskNextAuditDue

`func (o *AuditableEntitiesPutPreviousValues) GetRiskNextAuditDue() string`

GetRiskNextAuditDue returns the RiskNextAuditDue field if non-nil, zero value otherwise.

### GetRiskNextAuditDueOk

`func (o *AuditableEntitiesPutPreviousValues) GetRiskNextAuditDueOk() (*string, bool)`

GetRiskNextAuditDueOk returns a tuple with the RiskNextAuditDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskNextAuditDue

`func (o *AuditableEntitiesPutPreviousValues) SetRiskNextAuditDue(v string)`

SetRiskNextAuditDue sets RiskNextAuditDue field to given value.

### HasRiskNextAuditDue

`func (o *AuditableEntitiesPutPreviousValues) HasRiskNextAuditDue() bool`

HasRiskNextAuditDue returns a boolean if a field has been set.

### GetRegulatoryNextAuditDue

`func (o *AuditableEntitiesPutPreviousValues) GetRegulatoryNextAuditDue() string`

GetRegulatoryNextAuditDue returns the RegulatoryNextAuditDue field if non-nil, zero value otherwise.

### GetRegulatoryNextAuditDueOk

`func (o *AuditableEntitiesPutPreviousValues) GetRegulatoryNextAuditDueOk() (*string, bool)`

GetRegulatoryNextAuditDueOk returns a tuple with the RegulatoryNextAuditDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryNextAuditDue

`func (o *AuditableEntitiesPutPreviousValues) SetRegulatoryNextAuditDue(v string)`

SetRegulatoryNextAuditDue sets RegulatoryNextAuditDue field to given value.

### HasRegulatoryNextAuditDue

`func (o *AuditableEntitiesPutPreviousValues) HasRegulatoryNextAuditDue() bool`

HasRegulatoryNextAuditDue returns a boolean if a field has been set.

### GetRegulatoryNextAuditDueDate

`func (o *AuditableEntitiesPutPreviousValues) GetRegulatoryNextAuditDueDate() string`

GetRegulatoryNextAuditDueDate returns the RegulatoryNextAuditDueDate field if non-nil, zero value otherwise.

### GetRegulatoryNextAuditDueDateOk

`func (o *AuditableEntitiesPutPreviousValues) GetRegulatoryNextAuditDueDateOk() (*string, bool)`

GetRegulatoryNextAuditDueDateOk returns a tuple with the RegulatoryNextAuditDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryNextAuditDueDate

`func (o *AuditableEntitiesPutPreviousValues) SetRegulatoryNextAuditDueDate(v string)`

SetRegulatoryNextAuditDueDate sets RegulatoryNextAuditDueDate field to given value.

### HasRegulatoryNextAuditDueDate

`func (o *AuditableEntitiesPutPreviousValues) HasRegulatoryNextAuditDueDate() bool`

HasRegulatoryNextAuditDueDate returns a boolean if a field has been set.

### GetRiskNextAuditDueDate

`func (o *AuditableEntitiesPutPreviousValues) GetRiskNextAuditDueDate() string`

GetRiskNextAuditDueDate returns the RiskNextAuditDueDate field if non-nil, zero value otherwise.

### GetRiskNextAuditDueDateOk

`func (o *AuditableEntitiesPutPreviousValues) GetRiskNextAuditDueDateOk() (*string, bool)`

GetRiskNextAuditDueDateOk returns a tuple with the RiskNextAuditDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskNextAuditDueDate

`func (o *AuditableEntitiesPutPreviousValues) SetRiskNextAuditDueDate(v string)`

SetRiskNextAuditDueDate sets RiskNextAuditDueDate field to given value.

### HasRiskNextAuditDueDate

`func (o *AuditableEntitiesPutPreviousValues) HasRiskNextAuditDueDate() bool`

HasRiskNextAuditDueDate returns a boolean if a field has been set.

### GetItSecurityIncidentHours

`func (o *AuditableEntitiesPutPreviousValues) GetItSecurityIncidentHours() float32`

GetItSecurityIncidentHours returns the ItSecurityIncidentHours field if non-nil, zero value otherwise.

### GetItSecurityIncidentHoursOk

`func (o *AuditableEntitiesPutPreviousValues) GetItSecurityIncidentHoursOk() (*float32, bool)`

GetItSecurityIncidentHoursOk returns a tuple with the ItSecurityIncidentHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItSecurityIncidentHours

`func (o *AuditableEntitiesPutPreviousValues) SetItSecurityIncidentHours(v float32)`

SetItSecurityIncidentHours sets ItSecurityIncidentHours field to given value.

### HasItSecurityIncidentHours

`func (o *AuditableEntitiesPutPreviousValues) HasItSecurityIncidentHours() bool`

HasItSecurityIncidentHours returns a boolean if a field has been set.

### GetResidualRiskCalc

`func (o *AuditableEntitiesPutPreviousValues) GetResidualRiskCalc() interface{}`

GetResidualRiskCalc returns the ResidualRiskCalc field if non-nil, zero value otherwise.

### GetResidualRiskCalcOk

`func (o *AuditableEntitiesPutPreviousValues) GetResidualRiskCalcOk() (*interface{}, bool)`

GetResidualRiskCalcOk returns a tuple with the ResidualRiskCalc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidualRiskCalc

`func (o *AuditableEntitiesPutPreviousValues) SetResidualRiskCalc(v interface{})`

SetResidualRiskCalc sets ResidualRiskCalc field to given value.

### HasResidualRiskCalc

`func (o *AuditableEntitiesPutPreviousValues) HasResidualRiskCalc() bool`

HasResidualRiskCalc returns a boolean if a field has been set.

### SetResidualRiskCalcNil

`func (o *AuditableEntitiesPutPreviousValues) SetResidualRiskCalcNil(b bool)`

 SetResidualRiskCalcNil sets the value for ResidualRiskCalc to be an explicit nil

### UnsetResidualRiskCalc
`func (o *AuditableEntitiesPutPreviousValues) UnsetResidualRiskCalc()`

UnsetResidualRiskCalc ensures that no value is present for ResidualRiskCalc, not even an explicit nil
### GetSystemComponentTypeId

`func (o *AuditableEntitiesPutPreviousValues) GetSystemComponentTypeId() int32`

GetSystemComponentTypeId returns the SystemComponentTypeId field if non-nil, zero value otherwise.

### GetSystemComponentTypeIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetSystemComponentTypeIdOk() (*int32, bool)`

GetSystemComponentTypeIdOk returns a tuple with the SystemComponentTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemComponentTypeId

`func (o *AuditableEntitiesPutPreviousValues) SetSystemComponentTypeId(v int32)`

SetSystemComponentTypeId sets SystemComponentTypeId field to given value.

### HasSystemComponentTypeId

`func (o *AuditableEntitiesPutPreviousValues) HasSystemComponentTypeId() bool`

HasSystemComponentTypeId returns a boolean if a field has been set.

### GetSoxScopeId

`func (o *AuditableEntitiesPutPreviousValues) GetSoxScopeId() int32`

GetSoxScopeId returns the SoxScopeId field if non-nil, zero value otherwise.

### GetSoxScopeIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetSoxScopeIdOk() (*int32, bool)`

GetSoxScopeIdOk returns a tuple with the SoxScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoxScopeId

`func (o *AuditableEntitiesPutPreviousValues) SetSoxScopeId(v int32)`

SetSoxScopeId sets SoxScopeId field to given value.

### HasSoxScopeId

`func (o *AuditableEntitiesPutPreviousValues) HasSoxScopeId() bool`

HasSoxScopeId returns a boolean if a field has been set.

### GetEstimatedSpend

`func (o *AuditableEntitiesPutPreviousValues) GetEstimatedSpend() float32`

GetEstimatedSpend returns the EstimatedSpend field if non-nil, zero value otherwise.

### GetEstimatedSpendOk

`func (o *AuditableEntitiesPutPreviousValues) GetEstimatedSpendOk() (*float32, bool)`

GetEstimatedSpendOk returns a tuple with the EstimatedSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedSpend

`func (o *AuditableEntitiesPutPreviousValues) SetEstimatedSpend(v float32)`

SetEstimatedSpend sets EstimatedSpend field to given value.

### HasEstimatedSpend

`func (o *AuditableEntitiesPutPreviousValues) HasEstimatedSpend() bool`

HasEstimatedSpend returns a boolean if a field has been set.

### GetScopeRationale

`func (o *AuditableEntitiesPutPreviousValues) GetScopeRationale() string`

GetScopeRationale returns the ScopeRationale field if non-nil, zero value otherwise.

### GetScopeRationaleOk

`func (o *AuditableEntitiesPutPreviousValues) GetScopeRationaleOk() (*string, bool)`

GetScopeRationaleOk returns a tuple with the ScopeRationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeRationale

`func (o *AuditableEntitiesPutPreviousValues) SetScopeRationale(v string)`

SetScopeRationale sets ScopeRationale field to given value.

### HasScopeRationale

`func (o *AuditableEntitiesPutPreviousValues) HasScopeRationale() bool`

HasScopeRationale returns a boolean if a field has been set.

### GetFinancialApplicationId

`func (o *AuditableEntitiesPutPreviousValues) GetFinancialApplicationId() int32`

GetFinancialApplicationId returns the FinancialApplicationId field if non-nil, zero value otherwise.

### GetFinancialApplicationIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetFinancialApplicationIdOk() (*int32, bool)`

GetFinancialApplicationIdOk returns a tuple with the FinancialApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialApplicationId

`func (o *AuditableEntitiesPutPreviousValues) SetFinancialApplicationId(v int32)`

SetFinancialApplicationId sets FinancialApplicationId field to given value.

### HasFinancialApplicationId

`func (o *AuditableEntitiesPutPreviousValues) HasFinancialApplicationId() bool`

HasFinancialApplicationId returns a boolean if a field has been set.

### GetPriorAuditTitle

`func (o *AuditableEntitiesPutPreviousValues) GetPriorAuditTitle() string`

GetPriorAuditTitle returns the PriorAuditTitle field if non-nil, zero value otherwise.

### GetPriorAuditTitleOk

`func (o *AuditableEntitiesPutPreviousValues) GetPriorAuditTitleOk() (*string, bool)`

GetPriorAuditTitleOk returns a tuple with the PriorAuditTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorAuditTitle

`func (o *AuditableEntitiesPutPreviousValues) SetPriorAuditTitle(v string)`

SetPriorAuditTitle sets PriorAuditTitle field to given value.

### HasPriorAuditTitle

`func (o *AuditableEntitiesPutPreviousValues) HasPriorAuditTitle() bool`

HasPriorAuditTitle returns a boolean if a field has been set.

### GetUpcomingAuditTitle

`func (o *AuditableEntitiesPutPreviousValues) GetUpcomingAuditTitle() string`

GetUpcomingAuditTitle returns the UpcomingAuditTitle field if non-nil, zero value otherwise.

### GetUpcomingAuditTitleOk

`func (o *AuditableEntitiesPutPreviousValues) GetUpcomingAuditTitleOk() (*string, bool)`

GetUpcomingAuditTitleOk returns a tuple with the UpcomingAuditTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcomingAuditTitle

`func (o *AuditableEntitiesPutPreviousValues) SetUpcomingAuditTitle(v string)`

SetUpcomingAuditTitle sets UpcomingAuditTitle field to given value.

### HasUpcomingAuditTitle

`func (o *AuditableEntitiesPutPreviousValues) HasUpcomingAuditTitle() bool`

HasUpcomingAuditTitle returns a boolean if a field has been set.

### GetPriorAuditOpinion

`func (o *AuditableEntitiesPutPreviousValues) GetPriorAuditOpinion() string`

GetPriorAuditOpinion returns the PriorAuditOpinion field if non-nil, zero value otherwise.

### GetPriorAuditOpinionOk

`func (o *AuditableEntitiesPutPreviousValues) GetPriorAuditOpinionOk() (*string, bool)`

GetPriorAuditOpinionOk returns a tuple with the PriorAuditOpinion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorAuditOpinion

`func (o *AuditableEntitiesPutPreviousValues) SetPriorAuditOpinion(v string)`

SetPriorAuditOpinion sets PriorAuditOpinion field to given value.

### HasPriorAuditOpinion

`func (o *AuditableEntitiesPutPreviousValues) HasPriorAuditOpinion() bool`

HasPriorAuditOpinion returns a boolean if a field has been set.

### GetUpcomingAuditStartDate

`func (o *AuditableEntitiesPutPreviousValues) GetUpcomingAuditStartDate() string`

GetUpcomingAuditStartDate returns the UpcomingAuditStartDate field if non-nil, zero value otherwise.

### GetUpcomingAuditStartDateOk

`func (o *AuditableEntitiesPutPreviousValues) GetUpcomingAuditStartDateOk() (*string, bool)`

GetUpcomingAuditStartDateOk returns a tuple with the UpcomingAuditStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcomingAuditStartDate

`func (o *AuditableEntitiesPutPreviousValues) SetUpcomingAuditStartDate(v string)`

SetUpcomingAuditStartDate sets UpcomingAuditStartDate field to given value.

### HasUpcomingAuditStartDate

`func (o *AuditableEntitiesPutPreviousValues) HasUpcomingAuditStartDate() bool`

HasUpcomingAuditStartDate returns a boolean if a field has been set.

### GetUpcomingAuditEndDate

`func (o *AuditableEntitiesPutPreviousValues) GetUpcomingAuditEndDate() string`

GetUpcomingAuditEndDate returns the UpcomingAuditEndDate field if non-nil, zero value otherwise.

### GetUpcomingAuditEndDateOk

`func (o *AuditableEntitiesPutPreviousValues) GetUpcomingAuditEndDateOk() (*string, bool)`

GetUpcomingAuditEndDateOk returns a tuple with the UpcomingAuditEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcomingAuditEndDate

`func (o *AuditableEntitiesPutPreviousValues) SetUpcomingAuditEndDate(v string)`

SetUpcomingAuditEndDate sets UpcomingAuditEndDate field to given value.

### HasUpcomingAuditEndDate

`func (o *AuditableEntitiesPutPreviousValues) HasUpcomingAuditEndDate() bool`

HasUpcomingAuditEndDate returns a boolean if a field has been set.

### GetConfidentialityImpact

`func (o *AuditableEntitiesPutPreviousValues) GetConfidentialityImpact() float32`

GetConfidentialityImpact returns the ConfidentialityImpact field if non-nil, zero value otherwise.

### GetConfidentialityImpactOk

`func (o *AuditableEntitiesPutPreviousValues) GetConfidentialityImpactOk() (*float32, bool)`

GetConfidentialityImpactOk returns a tuple with the ConfidentialityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialityImpact

`func (o *AuditableEntitiesPutPreviousValues) SetConfidentialityImpact(v float32)`

SetConfidentialityImpact sets ConfidentialityImpact field to given value.

### HasConfidentialityImpact

`func (o *AuditableEntitiesPutPreviousValues) HasConfidentialityImpact() bool`

HasConfidentialityImpact returns a boolean if a field has been set.

### GetIntegrityImpact

`func (o *AuditableEntitiesPutPreviousValues) GetIntegrityImpact() float32`

GetIntegrityImpact returns the IntegrityImpact field if non-nil, zero value otherwise.

### GetIntegrityImpactOk

`func (o *AuditableEntitiesPutPreviousValues) GetIntegrityImpactOk() (*float32, bool)`

GetIntegrityImpactOk returns a tuple with the IntegrityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityImpact

`func (o *AuditableEntitiesPutPreviousValues) SetIntegrityImpact(v float32)`

SetIntegrityImpact sets IntegrityImpact field to given value.

### HasIntegrityImpact

`func (o *AuditableEntitiesPutPreviousValues) HasIntegrityImpact() bool`

HasIntegrityImpact returns a boolean if a field has been set.

### GetAvailabilityImpact

`func (o *AuditableEntitiesPutPreviousValues) GetAvailabilityImpact() float32`

GetAvailabilityImpact returns the AvailabilityImpact field if non-nil, zero value otherwise.

### GetAvailabilityImpactOk

`func (o *AuditableEntitiesPutPreviousValues) GetAvailabilityImpactOk() (*float32, bool)`

GetAvailabilityImpactOk returns a tuple with the AvailabilityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityImpact

`func (o *AuditableEntitiesPutPreviousValues) SetAvailabilityImpact(v float32)`

SetAvailabilityImpact sets AvailabilityImpact field to given value.

### HasAvailabilityImpact

`func (o *AuditableEntitiesPutPreviousValues) HasAvailabilityImpact() bool`

HasAvailabilityImpact returns a boolean if a field has been set.

### GetCiaInherentRiskCalc

`func (o *AuditableEntitiesPutPreviousValues) GetCiaInherentRiskCalc() interface{}`

GetCiaInherentRiskCalc returns the CiaInherentRiskCalc field if non-nil, zero value otherwise.

### GetCiaInherentRiskCalcOk

`func (o *AuditableEntitiesPutPreviousValues) GetCiaInherentRiskCalcOk() (*interface{}, bool)`

GetCiaInherentRiskCalcOk returns a tuple with the CiaInherentRiskCalc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiaInherentRiskCalc

`func (o *AuditableEntitiesPutPreviousValues) SetCiaInherentRiskCalc(v interface{})`

SetCiaInherentRiskCalc sets CiaInherentRiskCalc field to given value.

### HasCiaInherentRiskCalc

`func (o *AuditableEntitiesPutPreviousValues) HasCiaInherentRiskCalc() bool`

HasCiaInherentRiskCalc returns a boolean if a field has been set.

### SetCiaInherentRiskCalcNil

`func (o *AuditableEntitiesPutPreviousValues) SetCiaInherentRiskCalcNil(b bool)`

 SetCiaInherentRiskCalcNil sets the value for CiaInherentRiskCalc to be an explicit nil

### UnsetCiaInherentRiskCalc
`func (o *AuditableEntitiesPutPreviousValues) UnsetCiaInherentRiskCalc()`

UnsetCiaInherentRiskCalc ensures that no value is present for CiaInherentRiskCalc, not even an explicit nil
### GetCiaResidualRiskCalc

`func (o *AuditableEntitiesPutPreviousValues) GetCiaResidualRiskCalc() interface{}`

GetCiaResidualRiskCalc returns the CiaResidualRiskCalc field if non-nil, zero value otherwise.

### GetCiaResidualRiskCalcOk

`func (o *AuditableEntitiesPutPreviousValues) GetCiaResidualRiskCalcOk() (*interface{}, bool)`

GetCiaResidualRiskCalcOk returns a tuple with the CiaResidualRiskCalc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiaResidualRiskCalc

`func (o *AuditableEntitiesPutPreviousValues) SetCiaResidualRiskCalc(v interface{})`

SetCiaResidualRiskCalc sets CiaResidualRiskCalc field to given value.

### HasCiaResidualRiskCalc

`func (o *AuditableEntitiesPutPreviousValues) HasCiaResidualRiskCalc() bool`

HasCiaResidualRiskCalc returns a boolean if a field has been set.

### SetCiaResidualRiskCalcNil

`func (o *AuditableEntitiesPutPreviousValues) SetCiaResidualRiskCalcNil(b bool)`

 SetCiaResidualRiskCalcNil sets the value for CiaResidualRiskCalc to be an explicit nil

### UnsetCiaResidualRiskCalc
`func (o *AuditableEntitiesPutPreviousValues) UnsetCiaResidualRiskCalc()`

UnsetCiaResidualRiskCalc ensures that no value is present for CiaResidualRiskCalc, not even an explicit nil
### GetIntakeStatus

`func (o *AuditableEntitiesPutPreviousValues) GetIntakeStatus() string`

GetIntakeStatus returns the IntakeStatus field if non-nil, zero value otherwise.

### GetIntakeStatusOk

`func (o *AuditableEntitiesPutPreviousValues) GetIntakeStatusOk() (*string, bool)`

GetIntakeStatusOk returns a tuple with the IntakeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntakeStatus

`func (o *AuditableEntitiesPutPreviousValues) SetIntakeStatus(v string)`

SetIntakeStatus sets IntakeStatus field to given value.

### HasIntakeStatus

`func (o *AuditableEntitiesPutPreviousValues) HasIntakeStatus() bool`

HasIntakeStatus returns a boolean if a field has been set.

### GetCustomDate5

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDate5() string`

GetCustomDate5 returns the CustomDate5 field if non-nil, zero value otherwise.

### GetCustomDate5Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDate5Ok() (*string, bool)`

GetCustomDate5Ok returns a tuple with the CustomDate5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate5

`func (o *AuditableEntitiesPutPreviousValues) SetCustomDate5(v string)`

SetCustomDate5 sets CustomDate5 field to given value.

### HasCustomDate5

`func (o *AuditableEntitiesPutPreviousValues) HasCustomDate5() bool`

HasCustomDate5 returns a boolean if a field has been set.

### GetCustomDate6

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDate6() string`

GetCustomDate6 returns the CustomDate6 field if non-nil, zero value otherwise.

### GetCustomDate6Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDate6Ok() (*string, bool)`

GetCustomDate6Ok returns a tuple with the CustomDate6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate6

`func (o *AuditableEntitiesPutPreviousValues) SetCustomDate6(v string)`

SetCustomDate6 sets CustomDate6 field to given value.

### HasCustomDate6

`func (o *AuditableEntitiesPutPreviousValues) HasCustomDate6() bool`

HasCustomDate6 returns a boolean if a field has been set.

### GetCustomDate7

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDate7() string`

GetCustomDate7 returns the CustomDate7 field if non-nil, zero value otherwise.

### GetCustomDate7Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDate7Ok() (*string, bool)`

GetCustomDate7Ok returns a tuple with the CustomDate7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate7

`func (o *AuditableEntitiesPutPreviousValues) SetCustomDate7(v string)`

SetCustomDate7 sets CustomDate7 field to given value.

### HasCustomDate7

`func (o *AuditableEntitiesPutPreviousValues) HasCustomDate7() bool`

HasCustomDate7 returns a boolean if a field has been set.

### GetCustomDate8

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDate8() string`

GetCustomDate8 returns the CustomDate8 field if non-nil, zero value otherwise.

### GetCustomDate8Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDate8Ok() (*string, bool)`

GetCustomDate8Ok returns a tuple with the CustomDate8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate8

`func (o *AuditableEntitiesPutPreviousValues) SetCustomDate8(v string)`

SetCustomDate8 sets CustomDate8 field to given value.

### HasCustomDate8

`func (o *AuditableEntitiesPutPreviousValues) HasCustomDate8() bool`

HasCustomDate8 returns a boolean if a field has been set.

### GetCustomDate9

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDate9() string`

GetCustomDate9 returns the CustomDate9 field if non-nil, zero value otherwise.

### GetCustomDate9Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDate9Ok() (*string, bool)`

GetCustomDate9Ok returns a tuple with the CustomDate9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate9

`func (o *AuditableEntitiesPutPreviousValues) SetCustomDate9(v string)`

SetCustomDate9 sets CustomDate9 field to given value.

### HasCustomDate9

`func (o *AuditableEntitiesPutPreviousValues) HasCustomDate9() bool`

HasCustomDate9 returns a boolean if a field has been set.

### GetCustomDate10

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDate10() string`

GetCustomDate10 returns the CustomDate10 field if non-nil, zero value otherwise.

### GetCustomDate10Ok

`func (o *AuditableEntitiesPutPreviousValues) GetCustomDate10Ok() (*string, bool)`

GetCustomDate10Ok returns a tuple with the CustomDate10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate10

`func (o *AuditableEntitiesPutPreviousValues) SetCustomDate10(v string)`

SetCustomDate10 sets CustomDate10 field to given value.

### HasCustomDate10

`func (o *AuditableEntitiesPutPreviousValues) HasCustomDate10() bool`

HasCustomDate10 returns a boolean if a field has been set.

### GetAeCustomSelect13OptionId

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect13OptionId() int32`

GetAeCustomSelect13OptionId returns the AeCustomSelect13OptionId field if non-nil, zero value otherwise.

### GetAeCustomSelect13OptionIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect13OptionIdOk() (*int32, bool)`

GetAeCustomSelect13OptionIdOk returns a tuple with the AeCustomSelect13OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeCustomSelect13OptionId

`func (o *AuditableEntitiesPutPreviousValues) SetAeCustomSelect13OptionId(v int32)`

SetAeCustomSelect13OptionId sets AeCustomSelect13OptionId field to given value.

### HasAeCustomSelect13OptionId

`func (o *AuditableEntitiesPutPreviousValues) HasAeCustomSelect13OptionId() bool`

HasAeCustomSelect13OptionId returns a boolean if a field has been set.

### GetAeCustomSelect14OptionId

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect14OptionId() int32`

GetAeCustomSelect14OptionId returns the AeCustomSelect14OptionId field if non-nil, zero value otherwise.

### GetAeCustomSelect14OptionIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect14OptionIdOk() (*int32, bool)`

GetAeCustomSelect14OptionIdOk returns a tuple with the AeCustomSelect14OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeCustomSelect14OptionId

`func (o *AuditableEntitiesPutPreviousValues) SetAeCustomSelect14OptionId(v int32)`

SetAeCustomSelect14OptionId sets AeCustomSelect14OptionId field to given value.

### HasAeCustomSelect14OptionId

`func (o *AuditableEntitiesPutPreviousValues) HasAeCustomSelect14OptionId() bool`

HasAeCustomSelect14OptionId returns a boolean if a field has been set.

### GetAeCustomSelect15OptionId

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect15OptionId() int32`

GetAeCustomSelect15OptionId returns the AeCustomSelect15OptionId field if non-nil, zero value otherwise.

### GetAeCustomSelect15OptionIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAeCustomSelect15OptionIdOk() (*int32, bool)`

GetAeCustomSelect15OptionIdOk returns a tuple with the AeCustomSelect15OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeCustomSelect15OptionId

`func (o *AuditableEntitiesPutPreviousValues) SetAeCustomSelect15OptionId(v int32)`

SetAeCustomSelect15OptionId sets AeCustomSelect15OptionId field to given value.

### HasAeCustomSelect15OptionId

`func (o *AuditableEntitiesPutPreviousValues) HasAeCustomSelect15OptionId() bool`

HasAeCustomSelect15OptionId returns a boolean if a field has been set.

### GetDataClassificationId

`func (o *AuditableEntitiesPutPreviousValues) GetDataClassificationId() int32`

GetDataClassificationId returns the DataClassificationId field if non-nil, zero value otherwise.

### GetDataClassificationIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetDataClassificationIdOk() (*int32, bool)`

GetDataClassificationIdOk returns a tuple with the DataClassificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataClassificationId

`func (o *AuditableEntitiesPutPreviousValues) SetDataClassificationId(v int32)`

SetDataClassificationId sets DataClassificationId field to given value.

### HasDataClassificationId

`func (o *AuditableEntitiesPutPreviousValues) HasDataClassificationId() bool`

HasDataClassificationId returns a boolean if a field has been set.

### GetCustomUserSelect1UserId

`func (o *AuditableEntitiesPutPreviousValues) GetCustomUserSelect1UserId() int32`

GetCustomUserSelect1UserId returns the CustomUserSelect1UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect1UserIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetCustomUserSelect1UserIdOk() (*int32, bool)`

GetCustomUserSelect1UserIdOk returns a tuple with the CustomUserSelect1UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect1UserId

`func (o *AuditableEntitiesPutPreviousValues) SetCustomUserSelect1UserId(v int32)`

SetCustomUserSelect1UserId sets CustomUserSelect1UserId field to given value.

### HasCustomUserSelect1UserId

`func (o *AuditableEntitiesPutPreviousValues) HasCustomUserSelect1UserId() bool`

HasCustomUserSelect1UserId returns a boolean if a field has been set.

### GetCustomUserSelect2UserId

`func (o *AuditableEntitiesPutPreviousValues) GetCustomUserSelect2UserId() int32`

GetCustomUserSelect2UserId returns the CustomUserSelect2UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect2UserIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetCustomUserSelect2UserIdOk() (*int32, bool)`

GetCustomUserSelect2UserIdOk returns a tuple with the CustomUserSelect2UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect2UserId

`func (o *AuditableEntitiesPutPreviousValues) SetCustomUserSelect2UserId(v int32)`

SetCustomUserSelect2UserId sets CustomUserSelect2UserId field to given value.

### HasCustomUserSelect2UserId

`func (o *AuditableEntitiesPutPreviousValues) HasCustomUserSelect2UserId() bool`

HasCustomUserSelect2UserId returns a boolean if a field has been set.

### GetCustomUserSelect3UserId

`func (o *AuditableEntitiesPutPreviousValues) GetCustomUserSelect3UserId() int32`

GetCustomUserSelect3UserId returns the CustomUserSelect3UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect3UserIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetCustomUserSelect3UserIdOk() (*int32, bool)`

GetCustomUserSelect3UserIdOk returns a tuple with the CustomUserSelect3UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect3UserId

`func (o *AuditableEntitiesPutPreviousValues) SetCustomUserSelect3UserId(v int32)`

SetCustomUserSelect3UserId sets CustomUserSelect3UserId field to given value.

### HasCustomUserSelect3UserId

`func (o *AuditableEntitiesPutPreviousValues) HasCustomUserSelect3UserId() bool`

HasCustomUserSelect3UserId returns a boolean if a field has been set.

### GetCustomUserSelect4UserId

`func (o *AuditableEntitiesPutPreviousValues) GetCustomUserSelect4UserId() int32`

GetCustomUserSelect4UserId returns the CustomUserSelect4UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect4UserIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetCustomUserSelect4UserIdOk() (*int32, bool)`

GetCustomUserSelect4UserIdOk returns a tuple with the CustomUserSelect4UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect4UserId

`func (o *AuditableEntitiesPutPreviousValues) SetCustomUserSelect4UserId(v int32)`

SetCustomUserSelect4UserId sets CustomUserSelect4UserId field to given value.

### HasCustomUserSelect4UserId

`func (o *AuditableEntitiesPutPreviousValues) HasCustomUserSelect4UserId() bool`

HasCustomUserSelect4UserId returns a boolean if a field has been set.

### GetCustomUserSelect5UserId

`func (o *AuditableEntitiesPutPreviousValues) GetCustomUserSelect5UserId() int32`

GetCustomUserSelect5UserId returns the CustomUserSelect5UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect5UserIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetCustomUserSelect5UserIdOk() (*int32, bool)`

GetCustomUserSelect5UserIdOk returns a tuple with the CustomUserSelect5UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect5UserId

`func (o *AuditableEntitiesPutPreviousValues) SetCustomUserSelect5UserId(v int32)`

SetCustomUserSelect5UserId sets CustomUserSelect5UserId field to given value.

### HasCustomUserSelect5UserId

`func (o *AuditableEntitiesPutPreviousValues) HasCustomUserSelect5UserId() bool`

HasCustomUserSelect5UserId returns a boolean if a field has been set.

### GetCustomUserSelect6UserId

`func (o *AuditableEntitiesPutPreviousValues) GetCustomUserSelect6UserId() int32`

GetCustomUserSelect6UserId returns the CustomUserSelect6UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect6UserIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetCustomUserSelect6UserIdOk() (*int32, bool)`

GetCustomUserSelect6UserIdOk returns a tuple with the CustomUserSelect6UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect6UserId

`func (o *AuditableEntitiesPutPreviousValues) SetCustomUserSelect6UserId(v int32)`

SetCustomUserSelect6UserId sets CustomUserSelect6UserId field to given value.

### HasCustomUserSelect6UserId

`func (o *AuditableEntitiesPutPreviousValues) HasCustomUserSelect6UserId() bool`

HasCustomUserSelect6UserId returns a boolean if a field has been set.

### GetCustomUserSelect7UserId

`func (o *AuditableEntitiesPutPreviousValues) GetCustomUserSelect7UserId() int32`

GetCustomUserSelect7UserId returns the CustomUserSelect7UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect7UserIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetCustomUserSelect7UserIdOk() (*int32, bool)`

GetCustomUserSelect7UserIdOk returns a tuple with the CustomUserSelect7UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect7UserId

`func (o *AuditableEntitiesPutPreviousValues) SetCustomUserSelect7UserId(v int32)`

SetCustomUserSelect7UserId sets CustomUserSelect7UserId field to given value.

### HasCustomUserSelect7UserId

`func (o *AuditableEntitiesPutPreviousValues) HasCustomUserSelect7UserId() bool`

HasCustomUserSelect7UserId returns a boolean if a field has been set.

### GetCustomUserSelect8UserId

`func (o *AuditableEntitiesPutPreviousValues) GetCustomUserSelect8UserId() int32`

GetCustomUserSelect8UserId returns the CustomUserSelect8UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect8UserIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetCustomUserSelect8UserIdOk() (*int32, bool)`

GetCustomUserSelect8UserIdOk returns a tuple with the CustomUserSelect8UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect8UserId

`func (o *AuditableEntitiesPutPreviousValues) SetCustomUserSelect8UserId(v int32)`

SetCustomUserSelect8UserId sets CustomUserSelect8UserId field to given value.

### HasCustomUserSelect8UserId

`func (o *AuditableEntitiesPutPreviousValues) HasCustomUserSelect8UserId() bool`

HasCustomUserSelect8UserId returns a boolean if a field has been set.

### GetCustomUserSelect9UserId

`func (o *AuditableEntitiesPutPreviousValues) GetCustomUserSelect9UserId() int32`

GetCustomUserSelect9UserId returns the CustomUserSelect9UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect9UserIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetCustomUserSelect9UserIdOk() (*int32, bool)`

GetCustomUserSelect9UserIdOk returns a tuple with the CustomUserSelect9UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect9UserId

`func (o *AuditableEntitiesPutPreviousValues) SetCustomUserSelect9UserId(v int32)`

SetCustomUserSelect9UserId sets CustomUserSelect9UserId field to given value.

### HasCustomUserSelect9UserId

`func (o *AuditableEntitiesPutPreviousValues) HasCustomUserSelect9UserId() bool`

HasCustomUserSelect9UserId returns a boolean if a field has been set.

### GetCustomUserSelect10UserId

`func (o *AuditableEntitiesPutPreviousValues) GetCustomUserSelect10UserId() int32`

GetCustomUserSelect10UserId returns the CustomUserSelect10UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect10UserIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetCustomUserSelect10UserIdOk() (*int32, bool)`

GetCustomUserSelect10UserIdOk returns a tuple with the CustomUserSelect10UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect10UserId

`func (o *AuditableEntitiesPutPreviousValues) SetCustomUserSelect10UserId(v int32)`

SetCustomUserSelect10UserId sets CustomUserSelect10UserId field to given value.

### HasCustomUserSelect10UserId

`func (o *AuditableEntitiesPutPreviousValues) HasCustomUserSelect10UserId() bool`

HasCustomUserSelect10UserId returns a boolean if a field has been set.

### GetBusinessOwnerUserId

`func (o *AuditableEntitiesPutPreviousValues) GetBusinessOwnerUserId() int32`

GetBusinessOwnerUserId returns the BusinessOwnerUserId field if non-nil, zero value otherwise.

### GetBusinessOwnerUserIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetBusinessOwnerUserIdOk() (*int32, bool)`

GetBusinessOwnerUserIdOk returns a tuple with the BusinessOwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessOwnerUserId

`func (o *AuditableEntitiesPutPreviousValues) SetBusinessOwnerUserId(v int32)`

SetBusinessOwnerUserId sets BusinessOwnerUserId field to given value.

### HasBusinessOwnerUserId

`func (o *AuditableEntitiesPutPreviousValues) HasBusinessOwnerUserId() bool`

HasBusinessOwnerUserId returns a boolean if a field has been set.

### GetTechnicalProductOwnerUserId

`func (o *AuditableEntitiesPutPreviousValues) GetTechnicalProductOwnerUserId() int32`

GetTechnicalProductOwnerUserId returns the TechnicalProductOwnerUserId field if non-nil, zero value otherwise.

### GetTechnicalProductOwnerUserIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetTechnicalProductOwnerUserIdOk() (*int32, bool)`

GetTechnicalProductOwnerUserIdOk returns a tuple with the TechnicalProductOwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalProductOwnerUserId

`func (o *AuditableEntitiesPutPreviousValues) SetTechnicalProductOwnerUserId(v int32)`

SetTechnicalProductOwnerUserId sets TechnicalProductOwnerUserId field to given value.

### HasTechnicalProductOwnerUserId

`func (o *AuditableEntitiesPutPreviousValues) HasTechnicalProductOwnerUserId() bool`

HasTechnicalProductOwnerUserId returns a boolean if a field has been set.

### GetThirdPartyResidualRiskCalc

`func (o *AuditableEntitiesPutPreviousValues) GetThirdPartyResidualRiskCalc() interface{}`

GetThirdPartyResidualRiskCalc returns the ThirdPartyResidualRiskCalc field if non-nil, zero value otherwise.

### GetThirdPartyResidualRiskCalcOk

`func (o *AuditableEntitiesPutPreviousValues) GetThirdPartyResidualRiskCalcOk() (*interface{}, bool)`

GetThirdPartyResidualRiskCalcOk returns a tuple with the ThirdPartyResidualRiskCalc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyResidualRiskCalc

`func (o *AuditableEntitiesPutPreviousValues) SetThirdPartyResidualRiskCalc(v interface{})`

SetThirdPartyResidualRiskCalc sets ThirdPartyResidualRiskCalc field to given value.

### HasThirdPartyResidualRiskCalc

`func (o *AuditableEntitiesPutPreviousValues) HasThirdPartyResidualRiskCalc() bool`

HasThirdPartyResidualRiskCalc returns a boolean if a field has been set.

### SetThirdPartyResidualRiskCalcNil

`func (o *AuditableEntitiesPutPreviousValues) SetThirdPartyResidualRiskCalcNil(b bool)`

 SetThirdPartyResidualRiskCalcNil sets the value for ThirdPartyResidualRiskCalc to be an explicit nil

### UnsetThirdPartyResidualRiskCalc
`func (o *AuditableEntitiesPutPreviousValues) UnsetThirdPartyResidualRiskCalc()`

UnsetThirdPartyResidualRiskCalc ensures that no value is present for ThirdPartyResidualRiskCalc, not even an explicit nil
### GetAuditableEntityParentId

`func (o *AuditableEntitiesPutPreviousValues) GetAuditableEntityParentId() int32`

GetAuditableEntityParentId returns the AuditableEntityParentId field if non-nil, zero value otherwise.

### GetAuditableEntityParentIdOk

`func (o *AuditableEntitiesPutPreviousValues) GetAuditableEntityParentIdOk() (*int32, bool)`

GetAuditableEntityParentIdOk returns a tuple with the AuditableEntityParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditableEntityParentId

`func (o *AuditableEntitiesPutPreviousValues) SetAuditableEntityParentId(v int32)`

SetAuditableEntityParentId sets AuditableEntityParentId field to given value.

### HasAuditableEntityParentId

`func (o *AuditableEntitiesPutPreviousValues) HasAuditableEntityParentId() bool`

HasAuditableEntityParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


