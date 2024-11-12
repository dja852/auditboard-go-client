# ControlsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**ControlId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;controls.id&#x60;.&lt;fk table&#x3D;&#39;controls&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**SubprocessesDatumId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;subprocesses_data.id&#x60;.&lt;fk table&#x3D;&#39;subprocesses_data&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskLevelId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_levels.id&#x60;.&lt;fk table&#x3D;&#39;risk_levels&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**SupervisorUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OwnerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CosoElementId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;coso_elements.id&#x60;.&lt;fk table&#x3D;&#39;coso_elements&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlTypeId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_types.id&#x60;.&lt;fk table&#x3D;&#39;control_types&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlNatureOptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_nature_options.id&#x60;.&lt;fk table&#x3D;&#39;control_nature_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**SignificanceOptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;significance_options.id&#x60;.&lt;fk table&#x3D;&#39;significance_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestProcedures** | Pointer to **string** |  | [optional] 
**PbcRequest** | Pointer to **string** |  | [optional] 
**SourceData** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**ControlFrequencyId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_frequencies.id&#x60;.&lt;fk table&#x3D;&#39;control_frequencies&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestTimingId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_timings.id&#x60;.&lt;fk table&#x3D;&#39;test_timings&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestTypeId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_types.id&#x60;.&lt;fk table&#x3D;&#39;test_types&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueTitle** | Pointer to **string** |  | [optional] 
**IssueDescription** | Pointer to **string** |  | [optional] 
**RemediationAction** | Pointer to **string** |  | [optional] 
**RemediationOwnerId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RemediationDate** | Pointer to **string** |  | [optional] 
**IssueIdentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_idents.id&#x60;.&lt;fk table&#x3D;&#39;issue_idents&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueAccount** | Pointer to **string** |  | [optional] 
**PotentialRisk** | Pointer to **string** |  | [optional] 
**ComplementaryControls** | Pointer to **string** |  | [optional] 
**QualitativeFactors** | Pointer to **string** |  | [optional] 
**IssueId** | Pointer to **string** |  | [optional] 
**GrossExposure** | Pointer to **string** |  | [optional] 
**AdjustedExposure** | Pointer to **string** |  | [optional] 
**LikelihoodId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;likelihoods.id&#x60;.&lt;fk table&#x3D;&#39;likelihoods&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**MagnitudeId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;magnitudes.id&#x60;.&lt;fk table&#x3D;&#39;magnitudes&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**DeficiencyLevelId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;deficiency_levels.id&#x60;.&lt;fk table&#x3D;&#39;deficiency_levels&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AggregationReferenceId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;aggregation_references.id&#x60;.&lt;fk table&#x3D;&#39;aggregation_references&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueComments** | Pointer to **string** |  | [optional] 
**PbcRequestAdditional** | Pointer to **string** |  | [optional] 
**SodControlId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;sod_controls.id&#x60;.&lt;fk table&#x3D;&#39;sod_controls&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**SamplingApproach** | Pointer to **string** |  | [optional] 
**MrcControlId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;mrc_controls.id&#x60;.&lt;fk table&#x3D;&#39;mrc_controls&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FraudControlId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;fraud_controls.id&#x60;.&lt;fk table&#x3D;&#39;fraud_controls&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**EntityId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;entities.id&#x60;.&lt;fk table&#x3D;&#39;entities&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**PbcNotes** | Pointer to **string** |  | [optional] 
**ControlDesignOptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_design_options.id&#x60;.&lt;fk table&#x3D;&#39;control_design_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlValidation** | Pointer to **string** |  | [optional] 
**AnnualSampleSize** | Pointer to **string** |  | [optional] 
**ControlScopeId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_scopes.id&#x60;.&lt;fk table&#x3D;&#39;control_scopes&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlClassificationId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_classifications.id&#x60;.&lt;fk table&#x3D;&#39;control_classifications&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlProcedures** | Pointer to **string** |  | [optional] 
**ManagementResponse** | Pointer to **string** |  | [optional] 
**ControlObjective** | Pointer to **string** |  | [optional] 
**DependentControls** | Pointer to **string** |  | [optional] 
**CompensatingControls** | Pointer to **string** |  | [optional] 
**AssertionNotes** | Pointer to **string** |  | [optional] 
**PrecisionLevel** | Pointer to **string** |  | [optional] 
**MrcEvidence** | Pointer to **string** |  | [optional] 
**CobitReference** | Pointer to **string** |  | [optional] 
**ControlReference** | Pointer to **string** |  | [optional] 
**MapDueDate** | Pointer to **string** |  | [optional] 
**IssueDiscussionDate** | Pointer to **string** |  | [optional] 
**IssueIndividualsPresent** | Pointer to **string** |  | [optional] 
**ReviewerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AnnualPopulationSize** | Pointer to **string** |  | [optional] 
**ControlInherentRiskOptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_inherent_risk_options.id&#x60;.&lt;fk table&#x3D;&#39;control_inherent_risk_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlRoutineOptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_routine_options.id&#x60;.&lt;fk table&#x3D;&#39;control_routine_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlJudgementDegreeOptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_judgement_degree_options.id&#x60;.&lt;fk table&#x3D;&#39;control_judgement_degree_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlManagementOverrideOptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_management_override_options.id&#x60;.&lt;fk table&#x3D;&#39;control_management_override_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlComplexityOptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_complexity_options.id&#x60;.&lt;fk table&#x3D;&#39;control_complexity_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ExternalAuditorViewableStatus** | Pointer to **string** |  | [optional] [default to "Visible"]
**EffectiveDate** | Pointer to **string** |  | [optional] 
**BaselineResultOptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;baseline_result_options.id&#x60;.&lt;fk table&#x3D;&#39;baseline_result_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ModificationDescription** | Pointer to **string** |  | [optional] [default to ""]
**ReviewerNotes** | Pointer to **string** |  | [optional] [default to ""]
**Rationale** | Pointer to **string** |  | [optional] [default to ""]
**BaselineDate** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**ProcessId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;processes.id&#x60;.&lt;fk table&#x3D;&#39;processes&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**SubprocessId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;subprocesses.id&#x60;.&lt;fk table&#x3D;&#39;subprocesses&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**LastModificationDate** | Pointer to **string** |  | [optional] 
**LastReviewDate** | Pointer to **string** |  | [optional] 
**RelianceOptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;reliance_options.id&#x60;.&lt;fk table&#x3D;&#39;reliance_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect1OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select1_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select1_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CoordinatorUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomText1** | Pointer to **string** |  | [optional] 
**ControlCustomText2** | Pointer to **string** |  | [optional] 
**ControlCustomText3** | Pointer to **string** |  | [optional] 
**ControlCustomText4** | Pointer to **string** |  | [optional] 
**ControlCustomSelect2OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select2_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select2_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect3OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select3_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select3_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect4OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select4_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select4_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect5OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select5_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select5_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ReferenceMeta** | **interface{}** |  | 
**ComplianceReviewerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomText5** | Pointer to **string** |  | [optional] 
**ControlCustomText6** | Pointer to **string** |  | [optional] 
**ControlCustomText7** | Pointer to **string** |  | [optional] 
**ControlCustomText8** | Pointer to **string** |  | [optional] 
**ControlCustomText9** | Pointer to **string** |  | [optional] 
**ControlCustomText10** | Pointer to **string** |  | [optional] 
**ControlCustomText11** | Pointer to **string** |  | [optional] 
**ControlCustomText12** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Xuid** | Pointer to **string** |  | [optional] 
**Scopes** | **interface{}** |  | 
**FieldData** | Pointer to **interface{}** |  | [optional] 
**ControlCustomSelect6OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select6_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select6_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect7OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select7_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select7_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect8OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select8_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select8_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect9OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select9_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select9_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect10OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select10_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select10_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect11OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select11_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select11_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect12OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select12_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select12_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomUserSelect1UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomUserSelect2UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomUserSelect3UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomUserSelect4UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomText13** | Pointer to **string** |  | [optional] 
**ControlCustomText14** | Pointer to **string** |  | [optional] 
**ControlCustomText15** | Pointer to **string** |  | [optional] 
**ControlCustomText16** | Pointer to **string** |  | [optional] 
**ControlCustomText17** | Pointer to **string** |  | [optional] 
**ControlCustomText18** | Pointer to **string** |  | [optional] 
**ControlCustomText19** | Pointer to **string** |  | [optional] 
**ControlCustomText20** | Pointer to **string** |  | [optional] 
**ControlCustomText21** | Pointer to **string** |  | [optional] 
**ControlCustomText22** | Pointer to **string** |  | [optional] 
**ControlCustomText23** | Pointer to **string** |  | [optional] 
**ControlCustomText24** | Pointer to **string** |  | [optional] 
**CustomDate1** | Pointer to **string** |  | [optional] 
**CustomDate2** | Pointer to **string** |  | [optional] 
**CustomDate3** | Pointer to **string** |  | [optional] 
**CustomDate4** | Pointer to **string** |  | [optional] 
**CustomDate5** | Pointer to **string** |  | [optional] 
**CustomUserSelect5UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomUserSelect6UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomUserSelect7UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomUserSelect8UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomUserSelect9UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomUserSelect10UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomText25** | Pointer to **string** |  | [optional] 
**ControlCustomText26** | Pointer to **string** |  | [optional] 
**ControlCustomText27** | Pointer to **string** |  | [optional] 
**ControlCustomText28** | Pointer to **string** |  | [optional] 
**ControlCustomText29** | Pointer to **string** |  | [optional] 
**ControlCustomText30** | Pointer to **string** |  | [optional] 
**ControlCustomText31** | Pointer to **string** |  | [optional] 
**ControlCustomText32** | Pointer to **string** |  | [optional] 
**ControlCustomSelect13OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select13_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select13_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect14OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select14_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select14_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect15OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select15_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select15_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect16OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select16_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select16_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlsSegmentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;controls_segments.id&#x60;.&lt;fk table&#x3D;&#39;controls_segments&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect17OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select17_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select17_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect18OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select18_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select18_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect19OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select19_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select19_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect20OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select20_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select20_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect21OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select21_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select21_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect22OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select22_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select22_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect23OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select23_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select23_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect24OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select24_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select24_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect25OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select25_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select25_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect26OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select26_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select26_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect27OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select27_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select27_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect28OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select28_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select28_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect29OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select29_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select29_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect30OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select30_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select30_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect31OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select31_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select31_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect32OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select32_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select32_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect33OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select33_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select33_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect34OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select34_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select34_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect35OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select35_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select35_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlCustomSelect36OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;control_custom_select36_options.id&#x60;.&lt;fk table&#x3D;&#39;control_custom_select36_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AtRiskStatus** | **string** |  | [default to "Not At Risk"]
**LatestResultId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;controls_datum_results.id&#x60;.&lt;fk table&#x3D;&#39;controls_datum_results&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 

## Methods

### NewControlsData

`func NewControlsData(referenceMeta interface{}, scopes interface{}, atRiskStatus string, ) *ControlsData`

NewControlsData instantiates a new ControlsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsDataWithDefaults

`func NewControlsDataWithDefaults() *ControlsData`

NewControlsDataWithDefaults instantiates a new ControlsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ControlsData) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControlsData) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControlsData) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ControlsData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ControlsData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ControlsData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ControlsData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ControlsData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ControlsData) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ControlsData) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ControlsData) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ControlsData) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ControlsData) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ControlsData) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ControlsData) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ControlsData) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetControlId

`func (o *ControlsData) GetControlId() int32`

GetControlId returns the ControlId field if non-nil, zero value otherwise.

### GetControlIdOk

`func (o *ControlsData) GetControlIdOk() (*int32, bool)`

GetControlIdOk returns a tuple with the ControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlId

`func (o *ControlsData) SetControlId(v int32)`

SetControlId sets ControlId field to given value.

### HasControlId

`func (o *ControlsData) HasControlId() bool`

HasControlId returns a boolean if a field has been set.

### GetDescription

`func (o *ControlsData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ControlsData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ControlsData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ControlsData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubprocessesDatumId

`func (o *ControlsData) GetSubprocessesDatumId() int32`

GetSubprocessesDatumId returns the SubprocessesDatumId field if non-nil, zero value otherwise.

### GetSubprocessesDatumIdOk

`func (o *ControlsData) GetSubprocessesDatumIdOk() (*int32, bool)`

GetSubprocessesDatumIdOk returns a tuple with the SubprocessesDatumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubprocessesDatumId

`func (o *ControlsData) SetSubprocessesDatumId(v int32)`

SetSubprocessesDatumId sets SubprocessesDatumId field to given value.

### HasSubprocessesDatumId

`func (o *ControlsData) HasSubprocessesDatumId() bool`

HasSubprocessesDatumId returns a boolean if a field has been set.

### GetRiskLevelId

`func (o *ControlsData) GetRiskLevelId() int32`

GetRiskLevelId returns the RiskLevelId field if non-nil, zero value otherwise.

### GetRiskLevelIdOk

`func (o *ControlsData) GetRiskLevelIdOk() (*int32, bool)`

GetRiskLevelIdOk returns a tuple with the RiskLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevelId

`func (o *ControlsData) SetRiskLevelId(v int32)`

SetRiskLevelId sets RiskLevelId field to given value.

### HasRiskLevelId

`func (o *ControlsData) HasRiskLevelId() bool`

HasRiskLevelId returns a boolean if a field has been set.

### GetSupervisorUserId

`func (o *ControlsData) GetSupervisorUserId() int32`

GetSupervisorUserId returns the SupervisorUserId field if non-nil, zero value otherwise.

### GetSupervisorUserIdOk

`func (o *ControlsData) GetSupervisorUserIdOk() (*int32, bool)`

GetSupervisorUserIdOk returns a tuple with the SupervisorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisorUserId

`func (o *ControlsData) SetSupervisorUserId(v int32)`

SetSupervisorUserId sets SupervisorUserId field to given value.

### HasSupervisorUserId

`func (o *ControlsData) HasSupervisorUserId() bool`

HasSupervisorUserId returns a boolean if a field has been set.

### GetOwnerUserId

`func (o *ControlsData) GetOwnerUserId() int32`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *ControlsData) GetOwnerUserIdOk() (*int32, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *ControlsData) SetOwnerUserId(v int32)`

SetOwnerUserId sets OwnerUserId field to given value.

### HasOwnerUserId

`func (o *ControlsData) HasOwnerUserId() bool`

HasOwnerUserId returns a boolean if a field has been set.

### GetCosoElementId

`func (o *ControlsData) GetCosoElementId() int32`

GetCosoElementId returns the CosoElementId field if non-nil, zero value otherwise.

### GetCosoElementIdOk

`func (o *ControlsData) GetCosoElementIdOk() (*int32, bool)`

GetCosoElementIdOk returns a tuple with the CosoElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosoElementId

`func (o *ControlsData) SetCosoElementId(v int32)`

SetCosoElementId sets CosoElementId field to given value.

### HasCosoElementId

`func (o *ControlsData) HasCosoElementId() bool`

HasCosoElementId returns a boolean if a field has been set.

### GetControlTypeId

`func (o *ControlsData) GetControlTypeId() int32`

GetControlTypeId returns the ControlTypeId field if non-nil, zero value otherwise.

### GetControlTypeIdOk

`func (o *ControlsData) GetControlTypeIdOk() (*int32, bool)`

GetControlTypeIdOk returns a tuple with the ControlTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlTypeId

`func (o *ControlsData) SetControlTypeId(v int32)`

SetControlTypeId sets ControlTypeId field to given value.

### HasControlTypeId

`func (o *ControlsData) HasControlTypeId() bool`

HasControlTypeId returns a boolean if a field has been set.

### GetControlNatureOptionId

`func (o *ControlsData) GetControlNatureOptionId() int32`

GetControlNatureOptionId returns the ControlNatureOptionId field if non-nil, zero value otherwise.

### GetControlNatureOptionIdOk

`func (o *ControlsData) GetControlNatureOptionIdOk() (*int32, bool)`

GetControlNatureOptionIdOk returns a tuple with the ControlNatureOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlNatureOptionId

`func (o *ControlsData) SetControlNatureOptionId(v int32)`

SetControlNatureOptionId sets ControlNatureOptionId field to given value.

### HasControlNatureOptionId

`func (o *ControlsData) HasControlNatureOptionId() bool`

HasControlNatureOptionId returns a boolean if a field has been set.

### GetSignificanceOptionId

`func (o *ControlsData) GetSignificanceOptionId() int32`

GetSignificanceOptionId returns the SignificanceOptionId field if non-nil, zero value otherwise.

### GetSignificanceOptionIdOk

`func (o *ControlsData) GetSignificanceOptionIdOk() (*int32, bool)`

GetSignificanceOptionIdOk returns a tuple with the SignificanceOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificanceOptionId

`func (o *ControlsData) SetSignificanceOptionId(v int32)`

SetSignificanceOptionId sets SignificanceOptionId field to given value.

### HasSignificanceOptionId

`func (o *ControlsData) HasSignificanceOptionId() bool`

HasSignificanceOptionId returns a boolean if a field has been set.

### GetTestProcedures

`func (o *ControlsData) GetTestProcedures() string`

GetTestProcedures returns the TestProcedures field if non-nil, zero value otherwise.

### GetTestProceduresOk

`func (o *ControlsData) GetTestProceduresOk() (*string, bool)`

GetTestProceduresOk returns a tuple with the TestProcedures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestProcedures

`func (o *ControlsData) SetTestProcedures(v string)`

SetTestProcedures sets TestProcedures field to given value.

### HasTestProcedures

`func (o *ControlsData) HasTestProcedures() bool`

HasTestProcedures returns a boolean if a field has been set.

### GetPbcRequest

`func (o *ControlsData) GetPbcRequest() string`

GetPbcRequest returns the PbcRequest field if non-nil, zero value otherwise.

### GetPbcRequestOk

`func (o *ControlsData) GetPbcRequestOk() (*string, bool)`

GetPbcRequestOk returns a tuple with the PbcRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbcRequest

`func (o *ControlsData) SetPbcRequest(v string)`

SetPbcRequest sets PbcRequest field to given value.

### HasPbcRequest

`func (o *ControlsData) HasPbcRequest() bool`

HasPbcRequest returns a boolean if a field has been set.

### GetSourceData

`func (o *ControlsData) GetSourceData() string`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *ControlsData) GetSourceDataOk() (*string, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *ControlsData) SetSourceData(v string)`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *ControlsData) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetNotes

`func (o *ControlsData) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ControlsData) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ControlsData) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ControlsData) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetControlFrequencyId

`func (o *ControlsData) GetControlFrequencyId() int32`

GetControlFrequencyId returns the ControlFrequencyId field if non-nil, zero value otherwise.

### GetControlFrequencyIdOk

`func (o *ControlsData) GetControlFrequencyIdOk() (*int32, bool)`

GetControlFrequencyIdOk returns a tuple with the ControlFrequencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlFrequencyId

`func (o *ControlsData) SetControlFrequencyId(v int32)`

SetControlFrequencyId sets ControlFrequencyId field to given value.

### HasControlFrequencyId

`func (o *ControlsData) HasControlFrequencyId() bool`

HasControlFrequencyId returns a boolean if a field has been set.

### GetTestTimingId

`func (o *ControlsData) GetTestTimingId() int32`

GetTestTimingId returns the TestTimingId field if non-nil, zero value otherwise.

### GetTestTimingIdOk

`func (o *ControlsData) GetTestTimingIdOk() (*int32, bool)`

GetTestTimingIdOk returns a tuple with the TestTimingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestTimingId

`func (o *ControlsData) SetTestTimingId(v int32)`

SetTestTimingId sets TestTimingId field to given value.

### HasTestTimingId

`func (o *ControlsData) HasTestTimingId() bool`

HasTestTimingId returns a boolean if a field has been set.

### GetTestTypeId

`func (o *ControlsData) GetTestTypeId() int32`

GetTestTypeId returns the TestTypeId field if non-nil, zero value otherwise.

### GetTestTypeIdOk

`func (o *ControlsData) GetTestTypeIdOk() (*int32, bool)`

GetTestTypeIdOk returns a tuple with the TestTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestTypeId

`func (o *ControlsData) SetTestTypeId(v int32)`

SetTestTypeId sets TestTypeId field to given value.

### HasTestTypeId

`func (o *ControlsData) HasTestTypeId() bool`

HasTestTypeId returns a boolean if a field has been set.

### GetIssueTitle

`func (o *ControlsData) GetIssueTitle() string`

GetIssueTitle returns the IssueTitle field if non-nil, zero value otherwise.

### GetIssueTitleOk

`func (o *ControlsData) GetIssueTitleOk() (*string, bool)`

GetIssueTitleOk returns a tuple with the IssueTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTitle

`func (o *ControlsData) SetIssueTitle(v string)`

SetIssueTitle sets IssueTitle field to given value.

### HasIssueTitle

`func (o *ControlsData) HasIssueTitle() bool`

HasIssueTitle returns a boolean if a field has been set.

### GetIssueDescription

`func (o *ControlsData) GetIssueDescription() string`

GetIssueDescription returns the IssueDescription field if non-nil, zero value otherwise.

### GetIssueDescriptionOk

`func (o *ControlsData) GetIssueDescriptionOk() (*string, bool)`

GetIssueDescriptionOk returns a tuple with the IssueDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDescription

`func (o *ControlsData) SetIssueDescription(v string)`

SetIssueDescription sets IssueDescription field to given value.

### HasIssueDescription

`func (o *ControlsData) HasIssueDescription() bool`

HasIssueDescription returns a boolean if a field has been set.

### GetRemediationAction

`func (o *ControlsData) GetRemediationAction() string`

GetRemediationAction returns the RemediationAction field if non-nil, zero value otherwise.

### GetRemediationActionOk

`func (o *ControlsData) GetRemediationActionOk() (*string, bool)`

GetRemediationActionOk returns a tuple with the RemediationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationAction

`func (o *ControlsData) SetRemediationAction(v string)`

SetRemediationAction sets RemediationAction field to given value.

### HasRemediationAction

`func (o *ControlsData) HasRemediationAction() bool`

HasRemediationAction returns a boolean if a field has been set.

### GetRemediationOwnerId

`func (o *ControlsData) GetRemediationOwnerId() int32`

GetRemediationOwnerId returns the RemediationOwnerId field if non-nil, zero value otherwise.

### GetRemediationOwnerIdOk

`func (o *ControlsData) GetRemediationOwnerIdOk() (*int32, bool)`

GetRemediationOwnerIdOk returns a tuple with the RemediationOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationOwnerId

`func (o *ControlsData) SetRemediationOwnerId(v int32)`

SetRemediationOwnerId sets RemediationOwnerId field to given value.

### HasRemediationOwnerId

`func (o *ControlsData) HasRemediationOwnerId() bool`

HasRemediationOwnerId returns a boolean if a field has been set.

### GetRemediationDate

`func (o *ControlsData) GetRemediationDate() string`

GetRemediationDate returns the RemediationDate field if non-nil, zero value otherwise.

### GetRemediationDateOk

`func (o *ControlsData) GetRemediationDateOk() (*string, bool)`

GetRemediationDateOk returns a tuple with the RemediationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationDate

`func (o *ControlsData) SetRemediationDate(v string)`

SetRemediationDate sets RemediationDate field to given value.

### HasRemediationDate

`func (o *ControlsData) HasRemediationDate() bool`

HasRemediationDate returns a boolean if a field has been set.

### GetIssueIdentId

`func (o *ControlsData) GetIssueIdentId() int32`

GetIssueIdentId returns the IssueIdentId field if non-nil, zero value otherwise.

### GetIssueIdentIdOk

`func (o *ControlsData) GetIssueIdentIdOk() (*int32, bool)`

GetIssueIdentIdOk returns a tuple with the IssueIdentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueIdentId

`func (o *ControlsData) SetIssueIdentId(v int32)`

SetIssueIdentId sets IssueIdentId field to given value.

### HasIssueIdentId

`func (o *ControlsData) HasIssueIdentId() bool`

HasIssueIdentId returns a boolean if a field has been set.

### GetIssueAccount

`func (o *ControlsData) GetIssueAccount() string`

GetIssueAccount returns the IssueAccount field if non-nil, zero value otherwise.

### GetIssueAccountOk

`func (o *ControlsData) GetIssueAccountOk() (*string, bool)`

GetIssueAccountOk returns a tuple with the IssueAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueAccount

`func (o *ControlsData) SetIssueAccount(v string)`

SetIssueAccount sets IssueAccount field to given value.

### HasIssueAccount

`func (o *ControlsData) HasIssueAccount() bool`

HasIssueAccount returns a boolean if a field has been set.

### GetPotentialRisk

`func (o *ControlsData) GetPotentialRisk() string`

GetPotentialRisk returns the PotentialRisk field if non-nil, zero value otherwise.

### GetPotentialRiskOk

`func (o *ControlsData) GetPotentialRiskOk() (*string, bool)`

GetPotentialRiskOk returns a tuple with the PotentialRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPotentialRisk

`func (o *ControlsData) SetPotentialRisk(v string)`

SetPotentialRisk sets PotentialRisk field to given value.

### HasPotentialRisk

`func (o *ControlsData) HasPotentialRisk() bool`

HasPotentialRisk returns a boolean if a field has been set.

### GetComplementaryControls

`func (o *ControlsData) GetComplementaryControls() string`

GetComplementaryControls returns the ComplementaryControls field if non-nil, zero value otherwise.

### GetComplementaryControlsOk

`func (o *ControlsData) GetComplementaryControlsOk() (*string, bool)`

GetComplementaryControlsOk returns a tuple with the ComplementaryControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplementaryControls

`func (o *ControlsData) SetComplementaryControls(v string)`

SetComplementaryControls sets ComplementaryControls field to given value.

### HasComplementaryControls

`func (o *ControlsData) HasComplementaryControls() bool`

HasComplementaryControls returns a boolean if a field has been set.

### GetQualitativeFactors

`func (o *ControlsData) GetQualitativeFactors() string`

GetQualitativeFactors returns the QualitativeFactors field if non-nil, zero value otherwise.

### GetQualitativeFactorsOk

`func (o *ControlsData) GetQualitativeFactorsOk() (*string, bool)`

GetQualitativeFactorsOk returns a tuple with the QualitativeFactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualitativeFactors

`func (o *ControlsData) SetQualitativeFactors(v string)`

SetQualitativeFactors sets QualitativeFactors field to given value.

### HasQualitativeFactors

`func (o *ControlsData) HasQualitativeFactors() bool`

HasQualitativeFactors returns a boolean if a field has been set.

### GetIssueId

`func (o *ControlsData) GetIssueId() string`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *ControlsData) GetIssueIdOk() (*string, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *ControlsData) SetIssueId(v string)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *ControlsData) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.

### GetGrossExposure

`func (o *ControlsData) GetGrossExposure() string`

GetGrossExposure returns the GrossExposure field if non-nil, zero value otherwise.

### GetGrossExposureOk

`func (o *ControlsData) GetGrossExposureOk() (*string, bool)`

GetGrossExposureOk returns a tuple with the GrossExposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossExposure

`func (o *ControlsData) SetGrossExposure(v string)`

SetGrossExposure sets GrossExposure field to given value.

### HasGrossExposure

`func (o *ControlsData) HasGrossExposure() bool`

HasGrossExposure returns a boolean if a field has been set.

### GetAdjustedExposure

`func (o *ControlsData) GetAdjustedExposure() string`

GetAdjustedExposure returns the AdjustedExposure field if non-nil, zero value otherwise.

### GetAdjustedExposureOk

`func (o *ControlsData) GetAdjustedExposureOk() (*string, bool)`

GetAdjustedExposureOk returns a tuple with the AdjustedExposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedExposure

`func (o *ControlsData) SetAdjustedExposure(v string)`

SetAdjustedExposure sets AdjustedExposure field to given value.

### HasAdjustedExposure

`func (o *ControlsData) HasAdjustedExposure() bool`

HasAdjustedExposure returns a boolean if a field has been set.

### GetLikelihoodId

`func (o *ControlsData) GetLikelihoodId() int32`

GetLikelihoodId returns the LikelihoodId field if non-nil, zero value otherwise.

### GetLikelihoodIdOk

`func (o *ControlsData) GetLikelihoodIdOk() (*int32, bool)`

GetLikelihoodIdOk returns a tuple with the LikelihoodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikelihoodId

`func (o *ControlsData) SetLikelihoodId(v int32)`

SetLikelihoodId sets LikelihoodId field to given value.

### HasLikelihoodId

`func (o *ControlsData) HasLikelihoodId() bool`

HasLikelihoodId returns a boolean if a field has been set.

### GetMagnitudeId

`func (o *ControlsData) GetMagnitudeId() int32`

GetMagnitudeId returns the MagnitudeId field if non-nil, zero value otherwise.

### GetMagnitudeIdOk

`func (o *ControlsData) GetMagnitudeIdOk() (*int32, bool)`

GetMagnitudeIdOk returns a tuple with the MagnitudeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagnitudeId

`func (o *ControlsData) SetMagnitudeId(v int32)`

SetMagnitudeId sets MagnitudeId field to given value.

### HasMagnitudeId

`func (o *ControlsData) HasMagnitudeId() bool`

HasMagnitudeId returns a boolean if a field has been set.

### GetDeficiencyLevelId

`func (o *ControlsData) GetDeficiencyLevelId() int32`

GetDeficiencyLevelId returns the DeficiencyLevelId field if non-nil, zero value otherwise.

### GetDeficiencyLevelIdOk

`func (o *ControlsData) GetDeficiencyLevelIdOk() (*int32, bool)`

GetDeficiencyLevelIdOk returns a tuple with the DeficiencyLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeficiencyLevelId

`func (o *ControlsData) SetDeficiencyLevelId(v int32)`

SetDeficiencyLevelId sets DeficiencyLevelId field to given value.

### HasDeficiencyLevelId

`func (o *ControlsData) HasDeficiencyLevelId() bool`

HasDeficiencyLevelId returns a boolean if a field has been set.

### GetAggregationReferenceId

`func (o *ControlsData) GetAggregationReferenceId() int32`

GetAggregationReferenceId returns the AggregationReferenceId field if non-nil, zero value otherwise.

### GetAggregationReferenceIdOk

`func (o *ControlsData) GetAggregationReferenceIdOk() (*int32, bool)`

GetAggregationReferenceIdOk returns a tuple with the AggregationReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationReferenceId

`func (o *ControlsData) SetAggregationReferenceId(v int32)`

SetAggregationReferenceId sets AggregationReferenceId field to given value.

### HasAggregationReferenceId

`func (o *ControlsData) HasAggregationReferenceId() bool`

HasAggregationReferenceId returns a boolean if a field has been set.

### GetIssueComments

`func (o *ControlsData) GetIssueComments() string`

GetIssueComments returns the IssueComments field if non-nil, zero value otherwise.

### GetIssueCommentsOk

`func (o *ControlsData) GetIssueCommentsOk() (*string, bool)`

GetIssueCommentsOk returns a tuple with the IssueComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueComments

`func (o *ControlsData) SetIssueComments(v string)`

SetIssueComments sets IssueComments field to given value.

### HasIssueComments

`func (o *ControlsData) HasIssueComments() bool`

HasIssueComments returns a boolean if a field has been set.

### GetPbcRequestAdditional

`func (o *ControlsData) GetPbcRequestAdditional() string`

GetPbcRequestAdditional returns the PbcRequestAdditional field if non-nil, zero value otherwise.

### GetPbcRequestAdditionalOk

`func (o *ControlsData) GetPbcRequestAdditionalOk() (*string, bool)`

GetPbcRequestAdditionalOk returns a tuple with the PbcRequestAdditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbcRequestAdditional

`func (o *ControlsData) SetPbcRequestAdditional(v string)`

SetPbcRequestAdditional sets PbcRequestAdditional field to given value.

### HasPbcRequestAdditional

`func (o *ControlsData) HasPbcRequestAdditional() bool`

HasPbcRequestAdditional returns a boolean if a field has been set.

### GetSodControlId

`func (o *ControlsData) GetSodControlId() int32`

GetSodControlId returns the SodControlId field if non-nil, zero value otherwise.

### GetSodControlIdOk

`func (o *ControlsData) GetSodControlIdOk() (*int32, bool)`

GetSodControlIdOk returns a tuple with the SodControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSodControlId

`func (o *ControlsData) SetSodControlId(v int32)`

SetSodControlId sets SodControlId field to given value.

### HasSodControlId

`func (o *ControlsData) HasSodControlId() bool`

HasSodControlId returns a boolean if a field has been set.

### GetSamplingApproach

`func (o *ControlsData) GetSamplingApproach() string`

GetSamplingApproach returns the SamplingApproach field if non-nil, zero value otherwise.

### GetSamplingApproachOk

`func (o *ControlsData) GetSamplingApproachOk() (*string, bool)`

GetSamplingApproachOk returns a tuple with the SamplingApproach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingApproach

`func (o *ControlsData) SetSamplingApproach(v string)`

SetSamplingApproach sets SamplingApproach field to given value.

### HasSamplingApproach

`func (o *ControlsData) HasSamplingApproach() bool`

HasSamplingApproach returns a boolean if a field has been set.

### GetMrcControlId

`func (o *ControlsData) GetMrcControlId() int32`

GetMrcControlId returns the MrcControlId field if non-nil, zero value otherwise.

### GetMrcControlIdOk

`func (o *ControlsData) GetMrcControlIdOk() (*int32, bool)`

GetMrcControlIdOk returns a tuple with the MrcControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrcControlId

`func (o *ControlsData) SetMrcControlId(v int32)`

SetMrcControlId sets MrcControlId field to given value.

### HasMrcControlId

`func (o *ControlsData) HasMrcControlId() bool`

HasMrcControlId returns a boolean if a field has been set.

### GetFraudControlId

`func (o *ControlsData) GetFraudControlId() int32`

GetFraudControlId returns the FraudControlId field if non-nil, zero value otherwise.

### GetFraudControlIdOk

`func (o *ControlsData) GetFraudControlIdOk() (*int32, bool)`

GetFraudControlIdOk returns a tuple with the FraudControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudControlId

`func (o *ControlsData) SetFraudControlId(v int32)`

SetFraudControlId sets FraudControlId field to given value.

### HasFraudControlId

`func (o *ControlsData) HasFraudControlId() bool`

HasFraudControlId returns a boolean if a field has been set.

### GetEntityId

`func (o *ControlsData) GetEntityId() int32`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ControlsData) GetEntityIdOk() (*int32, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ControlsData) SetEntityId(v int32)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *ControlsData) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetPbcNotes

`func (o *ControlsData) GetPbcNotes() string`

GetPbcNotes returns the PbcNotes field if non-nil, zero value otherwise.

### GetPbcNotesOk

`func (o *ControlsData) GetPbcNotesOk() (*string, bool)`

GetPbcNotesOk returns a tuple with the PbcNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbcNotes

`func (o *ControlsData) SetPbcNotes(v string)`

SetPbcNotes sets PbcNotes field to given value.

### HasPbcNotes

`func (o *ControlsData) HasPbcNotes() bool`

HasPbcNotes returns a boolean if a field has been set.

### GetControlDesignOptionId

`func (o *ControlsData) GetControlDesignOptionId() int32`

GetControlDesignOptionId returns the ControlDesignOptionId field if non-nil, zero value otherwise.

### GetControlDesignOptionIdOk

`func (o *ControlsData) GetControlDesignOptionIdOk() (*int32, bool)`

GetControlDesignOptionIdOk returns a tuple with the ControlDesignOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlDesignOptionId

`func (o *ControlsData) SetControlDesignOptionId(v int32)`

SetControlDesignOptionId sets ControlDesignOptionId field to given value.

### HasControlDesignOptionId

`func (o *ControlsData) HasControlDesignOptionId() bool`

HasControlDesignOptionId returns a boolean if a field has been set.

### GetControlValidation

`func (o *ControlsData) GetControlValidation() string`

GetControlValidation returns the ControlValidation field if non-nil, zero value otherwise.

### GetControlValidationOk

`func (o *ControlsData) GetControlValidationOk() (*string, bool)`

GetControlValidationOk returns a tuple with the ControlValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlValidation

`func (o *ControlsData) SetControlValidation(v string)`

SetControlValidation sets ControlValidation field to given value.

### HasControlValidation

`func (o *ControlsData) HasControlValidation() bool`

HasControlValidation returns a boolean if a field has been set.

### GetAnnualSampleSize

`func (o *ControlsData) GetAnnualSampleSize() string`

GetAnnualSampleSize returns the AnnualSampleSize field if non-nil, zero value otherwise.

### GetAnnualSampleSizeOk

`func (o *ControlsData) GetAnnualSampleSizeOk() (*string, bool)`

GetAnnualSampleSizeOk returns a tuple with the AnnualSampleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualSampleSize

`func (o *ControlsData) SetAnnualSampleSize(v string)`

SetAnnualSampleSize sets AnnualSampleSize field to given value.

### HasAnnualSampleSize

`func (o *ControlsData) HasAnnualSampleSize() bool`

HasAnnualSampleSize returns a boolean if a field has been set.

### GetControlScopeId

`func (o *ControlsData) GetControlScopeId() int32`

GetControlScopeId returns the ControlScopeId field if non-nil, zero value otherwise.

### GetControlScopeIdOk

`func (o *ControlsData) GetControlScopeIdOk() (*int32, bool)`

GetControlScopeIdOk returns a tuple with the ControlScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlScopeId

`func (o *ControlsData) SetControlScopeId(v int32)`

SetControlScopeId sets ControlScopeId field to given value.

### HasControlScopeId

`func (o *ControlsData) HasControlScopeId() bool`

HasControlScopeId returns a boolean if a field has been set.

### GetControlClassificationId

`func (o *ControlsData) GetControlClassificationId() int32`

GetControlClassificationId returns the ControlClassificationId field if non-nil, zero value otherwise.

### GetControlClassificationIdOk

`func (o *ControlsData) GetControlClassificationIdOk() (*int32, bool)`

GetControlClassificationIdOk returns a tuple with the ControlClassificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlClassificationId

`func (o *ControlsData) SetControlClassificationId(v int32)`

SetControlClassificationId sets ControlClassificationId field to given value.

### HasControlClassificationId

`func (o *ControlsData) HasControlClassificationId() bool`

HasControlClassificationId returns a boolean if a field has been set.

### GetControlProcedures

`func (o *ControlsData) GetControlProcedures() string`

GetControlProcedures returns the ControlProcedures field if non-nil, zero value otherwise.

### GetControlProceduresOk

`func (o *ControlsData) GetControlProceduresOk() (*string, bool)`

GetControlProceduresOk returns a tuple with the ControlProcedures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlProcedures

`func (o *ControlsData) SetControlProcedures(v string)`

SetControlProcedures sets ControlProcedures field to given value.

### HasControlProcedures

`func (o *ControlsData) HasControlProcedures() bool`

HasControlProcedures returns a boolean if a field has been set.

### GetManagementResponse

`func (o *ControlsData) GetManagementResponse() string`

GetManagementResponse returns the ManagementResponse field if non-nil, zero value otherwise.

### GetManagementResponseOk

`func (o *ControlsData) GetManagementResponseOk() (*string, bool)`

GetManagementResponseOk returns a tuple with the ManagementResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementResponse

`func (o *ControlsData) SetManagementResponse(v string)`

SetManagementResponse sets ManagementResponse field to given value.

### HasManagementResponse

`func (o *ControlsData) HasManagementResponse() bool`

HasManagementResponse returns a boolean if a field has been set.

### GetControlObjective

`func (o *ControlsData) GetControlObjective() string`

GetControlObjective returns the ControlObjective field if non-nil, zero value otherwise.

### GetControlObjectiveOk

`func (o *ControlsData) GetControlObjectiveOk() (*string, bool)`

GetControlObjectiveOk returns a tuple with the ControlObjective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlObjective

`func (o *ControlsData) SetControlObjective(v string)`

SetControlObjective sets ControlObjective field to given value.

### HasControlObjective

`func (o *ControlsData) HasControlObjective() bool`

HasControlObjective returns a boolean if a field has been set.

### GetDependentControls

`func (o *ControlsData) GetDependentControls() string`

GetDependentControls returns the DependentControls field if non-nil, zero value otherwise.

### GetDependentControlsOk

`func (o *ControlsData) GetDependentControlsOk() (*string, bool)`

GetDependentControlsOk returns a tuple with the DependentControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentControls

`func (o *ControlsData) SetDependentControls(v string)`

SetDependentControls sets DependentControls field to given value.

### HasDependentControls

`func (o *ControlsData) HasDependentControls() bool`

HasDependentControls returns a boolean if a field has been set.

### GetCompensatingControls

`func (o *ControlsData) GetCompensatingControls() string`

GetCompensatingControls returns the CompensatingControls field if non-nil, zero value otherwise.

### GetCompensatingControlsOk

`func (o *ControlsData) GetCompensatingControlsOk() (*string, bool)`

GetCompensatingControlsOk returns a tuple with the CompensatingControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompensatingControls

`func (o *ControlsData) SetCompensatingControls(v string)`

SetCompensatingControls sets CompensatingControls field to given value.

### HasCompensatingControls

`func (o *ControlsData) HasCompensatingControls() bool`

HasCompensatingControls returns a boolean if a field has been set.

### GetAssertionNotes

`func (o *ControlsData) GetAssertionNotes() string`

GetAssertionNotes returns the AssertionNotes field if non-nil, zero value otherwise.

### GetAssertionNotesOk

`func (o *ControlsData) GetAssertionNotesOk() (*string, bool)`

GetAssertionNotesOk returns a tuple with the AssertionNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionNotes

`func (o *ControlsData) SetAssertionNotes(v string)`

SetAssertionNotes sets AssertionNotes field to given value.

### HasAssertionNotes

`func (o *ControlsData) HasAssertionNotes() bool`

HasAssertionNotes returns a boolean if a field has been set.

### GetPrecisionLevel

`func (o *ControlsData) GetPrecisionLevel() string`

GetPrecisionLevel returns the PrecisionLevel field if non-nil, zero value otherwise.

### GetPrecisionLevelOk

`func (o *ControlsData) GetPrecisionLevelOk() (*string, bool)`

GetPrecisionLevelOk returns a tuple with the PrecisionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecisionLevel

`func (o *ControlsData) SetPrecisionLevel(v string)`

SetPrecisionLevel sets PrecisionLevel field to given value.

### HasPrecisionLevel

`func (o *ControlsData) HasPrecisionLevel() bool`

HasPrecisionLevel returns a boolean if a field has been set.

### GetMrcEvidence

`func (o *ControlsData) GetMrcEvidence() string`

GetMrcEvidence returns the MrcEvidence field if non-nil, zero value otherwise.

### GetMrcEvidenceOk

`func (o *ControlsData) GetMrcEvidenceOk() (*string, bool)`

GetMrcEvidenceOk returns a tuple with the MrcEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrcEvidence

`func (o *ControlsData) SetMrcEvidence(v string)`

SetMrcEvidence sets MrcEvidence field to given value.

### HasMrcEvidence

`func (o *ControlsData) HasMrcEvidence() bool`

HasMrcEvidence returns a boolean if a field has been set.

### GetCobitReference

`func (o *ControlsData) GetCobitReference() string`

GetCobitReference returns the CobitReference field if non-nil, zero value otherwise.

### GetCobitReferenceOk

`func (o *ControlsData) GetCobitReferenceOk() (*string, bool)`

GetCobitReferenceOk returns a tuple with the CobitReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCobitReference

`func (o *ControlsData) SetCobitReference(v string)`

SetCobitReference sets CobitReference field to given value.

### HasCobitReference

`func (o *ControlsData) HasCobitReference() bool`

HasCobitReference returns a boolean if a field has been set.

### GetControlReference

`func (o *ControlsData) GetControlReference() string`

GetControlReference returns the ControlReference field if non-nil, zero value otherwise.

### GetControlReferenceOk

`func (o *ControlsData) GetControlReferenceOk() (*string, bool)`

GetControlReferenceOk returns a tuple with the ControlReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlReference

`func (o *ControlsData) SetControlReference(v string)`

SetControlReference sets ControlReference field to given value.

### HasControlReference

`func (o *ControlsData) HasControlReference() bool`

HasControlReference returns a boolean if a field has been set.

### GetMapDueDate

`func (o *ControlsData) GetMapDueDate() string`

GetMapDueDate returns the MapDueDate field if non-nil, zero value otherwise.

### GetMapDueDateOk

`func (o *ControlsData) GetMapDueDateOk() (*string, bool)`

GetMapDueDateOk returns a tuple with the MapDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapDueDate

`func (o *ControlsData) SetMapDueDate(v string)`

SetMapDueDate sets MapDueDate field to given value.

### HasMapDueDate

`func (o *ControlsData) HasMapDueDate() bool`

HasMapDueDate returns a boolean if a field has been set.

### GetIssueDiscussionDate

`func (o *ControlsData) GetIssueDiscussionDate() string`

GetIssueDiscussionDate returns the IssueDiscussionDate field if non-nil, zero value otherwise.

### GetIssueDiscussionDateOk

`func (o *ControlsData) GetIssueDiscussionDateOk() (*string, bool)`

GetIssueDiscussionDateOk returns a tuple with the IssueDiscussionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDiscussionDate

`func (o *ControlsData) SetIssueDiscussionDate(v string)`

SetIssueDiscussionDate sets IssueDiscussionDate field to given value.

### HasIssueDiscussionDate

`func (o *ControlsData) HasIssueDiscussionDate() bool`

HasIssueDiscussionDate returns a boolean if a field has been set.

### GetIssueIndividualsPresent

`func (o *ControlsData) GetIssueIndividualsPresent() string`

GetIssueIndividualsPresent returns the IssueIndividualsPresent field if non-nil, zero value otherwise.

### GetIssueIndividualsPresentOk

`func (o *ControlsData) GetIssueIndividualsPresentOk() (*string, bool)`

GetIssueIndividualsPresentOk returns a tuple with the IssueIndividualsPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueIndividualsPresent

`func (o *ControlsData) SetIssueIndividualsPresent(v string)`

SetIssueIndividualsPresent sets IssueIndividualsPresent field to given value.

### HasIssueIndividualsPresent

`func (o *ControlsData) HasIssueIndividualsPresent() bool`

HasIssueIndividualsPresent returns a boolean if a field has been set.

### GetReviewerUserId

`func (o *ControlsData) GetReviewerUserId() int32`

GetReviewerUserId returns the ReviewerUserId field if non-nil, zero value otherwise.

### GetReviewerUserIdOk

`func (o *ControlsData) GetReviewerUserIdOk() (*int32, bool)`

GetReviewerUserIdOk returns a tuple with the ReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerUserId

`func (o *ControlsData) SetReviewerUserId(v int32)`

SetReviewerUserId sets ReviewerUserId field to given value.

### HasReviewerUserId

`func (o *ControlsData) HasReviewerUserId() bool`

HasReviewerUserId returns a boolean if a field has been set.

### GetAnnualPopulationSize

`func (o *ControlsData) GetAnnualPopulationSize() string`

GetAnnualPopulationSize returns the AnnualPopulationSize field if non-nil, zero value otherwise.

### GetAnnualPopulationSizeOk

`func (o *ControlsData) GetAnnualPopulationSizeOk() (*string, bool)`

GetAnnualPopulationSizeOk returns a tuple with the AnnualPopulationSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualPopulationSize

`func (o *ControlsData) SetAnnualPopulationSize(v string)`

SetAnnualPopulationSize sets AnnualPopulationSize field to given value.

### HasAnnualPopulationSize

`func (o *ControlsData) HasAnnualPopulationSize() bool`

HasAnnualPopulationSize returns a boolean if a field has been set.

### GetControlInherentRiskOptionId

`func (o *ControlsData) GetControlInherentRiskOptionId() int32`

GetControlInherentRiskOptionId returns the ControlInherentRiskOptionId field if non-nil, zero value otherwise.

### GetControlInherentRiskOptionIdOk

`func (o *ControlsData) GetControlInherentRiskOptionIdOk() (*int32, bool)`

GetControlInherentRiskOptionIdOk returns a tuple with the ControlInherentRiskOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlInherentRiskOptionId

`func (o *ControlsData) SetControlInherentRiskOptionId(v int32)`

SetControlInherentRiskOptionId sets ControlInherentRiskOptionId field to given value.

### HasControlInherentRiskOptionId

`func (o *ControlsData) HasControlInherentRiskOptionId() bool`

HasControlInherentRiskOptionId returns a boolean if a field has been set.

### GetControlRoutineOptionId

`func (o *ControlsData) GetControlRoutineOptionId() int32`

GetControlRoutineOptionId returns the ControlRoutineOptionId field if non-nil, zero value otherwise.

### GetControlRoutineOptionIdOk

`func (o *ControlsData) GetControlRoutineOptionIdOk() (*int32, bool)`

GetControlRoutineOptionIdOk returns a tuple with the ControlRoutineOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlRoutineOptionId

`func (o *ControlsData) SetControlRoutineOptionId(v int32)`

SetControlRoutineOptionId sets ControlRoutineOptionId field to given value.

### HasControlRoutineOptionId

`func (o *ControlsData) HasControlRoutineOptionId() bool`

HasControlRoutineOptionId returns a boolean if a field has been set.

### GetControlJudgementDegreeOptionId

`func (o *ControlsData) GetControlJudgementDegreeOptionId() int32`

GetControlJudgementDegreeOptionId returns the ControlJudgementDegreeOptionId field if non-nil, zero value otherwise.

### GetControlJudgementDegreeOptionIdOk

`func (o *ControlsData) GetControlJudgementDegreeOptionIdOk() (*int32, bool)`

GetControlJudgementDegreeOptionIdOk returns a tuple with the ControlJudgementDegreeOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlJudgementDegreeOptionId

`func (o *ControlsData) SetControlJudgementDegreeOptionId(v int32)`

SetControlJudgementDegreeOptionId sets ControlJudgementDegreeOptionId field to given value.

### HasControlJudgementDegreeOptionId

`func (o *ControlsData) HasControlJudgementDegreeOptionId() bool`

HasControlJudgementDegreeOptionId returns a boolean if a field has been set.

### GetControlManagementOverrideOptionId

`func (o *ControlsData) GetControlManagementOverrideOptionId() int32`

GetControlManagementOverrideOptionId returns the ControlManagementOverrideOptionId field if non-nil, zero value otherwise.

### GetControlManagementOverrideOptionIdOk

`func (o *ControlsData) GetControlManagementOverrideOptionIdOk() (*int32, bool)`

GetControlManagementOverrideOptionIdOk returns a tuple with the ControlManagementOverrideOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlManagementOverrideOptionId

`func (o *ControlsData) SetControlManagementOverrideOptionId(v int32)`

SetControlManagementOverrideOptionId sets ControlManagementOverrideOptionId field to given value.

### HasControlManagementOverrideOptionId

`func (o *ControlsData) HasControlManagementOverrideOptionId() bool`

HasControlManagementOverrideOptionId returns a boolean if a field has been set.

### GetControlComplexityOptionId

`func (o *ControlsData) GetControlComplexityOptionId() int32`

GetControlComplexityOptionId returns the ControlComplexityOptionId field if non-nil, zero value otherwise.

### GetControlComplexityOptionIdOk

`func (o *ControlsData) GetControlComplexityOptionIdOk() (*int32, bool)`

GetControlComplexityOptionIdOk returns a tuple with the ControlComplexityOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlComplexityOptionId

`func (o *ControlsData) SetControlComplexityOptionId(v int32)`

SetControlComplexityOptionId sets ControlComplexityOptionId field to given value.

### HasControlComplexityOptionId

`func (o *ControlsData) HasControlComplexityOptionId() bool`

HasControlComplexityOptionId returns a boolean if a field has been set.

### GetExternalAuditorViewableStatus

`func (o *ControlsData) GetExternalAuditorViewableStatus() string`

GetExternalAuditorViewableStatus returns the ExternalAuditorViewableStatus field if non-nil, zero value otherwise.

### GetExternalAuditorViewableStatusOk

`func (o *ControlsData) GetExternalAuditorViewableStatusOk() (*string, bool)`

GetExternalAuditorViewableStatusOk returns a tuple with the ExternalAuditorViewableStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAuditorViewableStatus

`func (o *ControlsData) SetExternalAuditorViewableStatus(v string)`

SetExternalAuditorViewableStatus sets ExternalAuditorViewableStatus field to given value.

### HasExternalAuditorViewableStatus

`func (o *ControlsData) HasExternalAuditorViewableStatus() bool`

HasExternalAuditorViewableStatus returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *ControlsData) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *ControlsData) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *ControlsData) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *ControlsData) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetBaselineResultOptionId

`func (o *ControlsData) GetBaselineResultOptionId() int32`

GetBaselineResultOptionId returns the BaselineResultOptionId field if non-nil, zero value otherwise.

### GetBaselineResultOptionIdOk

`func (o *ControlsData) GetBaselineResultOptionIdOk() (*int32, bool)`

GetBaselineResultOptionIdOk returns a tuple with the BaselineResultOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineResultOptionId

`func (o *ControlsData) SetBaselineResultOptionId(v int32)`

SetBaselineResultOptionId sets BaselineResultOptionId field to given value.

### HasBaselineResultOptionId

`func (o *ControlsData) HasBaselineResultOptionId() bool`

HasBaselineResultOptionId returns a boolean if a field has been set.

### GetModificationDescription

`func (o *ControlsData) GetModificationDescription() string`

GetModificationDescription returns the ModificationDescription field if non-nil, zero value otherwise.

### GetModificationDescriptionOk

`func (o *ControlsData) GetModificationDescriptionOk() (*string, bool)`

GetModificationDescriptionOk returns a tuple with the ModificationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDescription

`func (o *ControlsData) SetModificationDescription(v string)`

SetModificationDescription sets ModificationDescription field to given value.

### HasModificationDescription

`func (o *ControlsData) HasModificationDescription() bool`

HasModificationDescription returns a boolean if a field has been set.

### GetReviewerNotes

`func (o *ControlsData) GetReviewerNotes() string`

GetReviewerNotes returns the ReviewerNotes field if non-nil, zero value otherwise.

### GetReviewerNotesOk

`func (o *ControlsData) GetReviewerNotesOk() (*string, bool)`

GetReviewerNotesOk returns a tuple with the ReviewerNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerNotes

`func (o *ControlsData) SetReviewerNotes(v string)`

SetReviewerNotes sets ReviewerNotes field to given value.

### HasReviewerNotes

`func (o *ControlsData) HasReviewerNotes() bool`

HasReviewerNotes returns a boolean if a field has been set.

### GetRationale

`func (o *ControlsData) GetRationale() string`

GetRationale returns the Rationale field if non-nil, zero value otherwise.

### GetRationaleOk

`func (o *ControlsData) GetRationaleOk() (*string, bool)`

GetRationaleOk returns a tuple with the Rationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRationale

`func (o *ControlsData) SetRationale(v string)`

SetRationale sets Rationale field to given value.

### HasRationale

`func (o *ControlsData) HasRationale() bool`

HasRationale returns a boolean if a field has been set.

### GetBaselineDate

`func (o *ControlsData) GetBaselineDate() string`

GetBaselineDate returns the BaselineDate field if non-nil, zero value otherwise.

### GetBaselineDateOk

`func (o *ControlsData) GetBaselineDateOk() (*string, bool)`

GetBaselineDateOk returns a tuple with the BaselineDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineDate

`func (o *ControlsData) SetBaselineDate(v string)`

SetBaselineDate sets BaselineDate field to given value.

### HasBaselineDate

`func (o *ControlsData) HasBaselineDate() bool`

HasBaselineDate returns a boolean if a field has been set.

### GetUid

`func (o *ControlsData) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ControlsData) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ControlsData) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ControlsData) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetProcessId

`func (o *ControlsData) GetProcessId() int32`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *ControlsData) GetProcessIdOk() (*int32, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *ControlsData) SetProcessId(v int32)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *ControlsData) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### GetSubprocessId

`func (o *ControlsData) GetSubprocessId() int32`

GetSubprocessId returns the SubprocessId field if non-nil, zero value otherwise.

### GetSubprocessIdOk

`func (o *ControlsData) GetSubprocessIdOk() (*int32, bool)`

GetSubprocessIdOk returns a tuple with the SubprocessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubprocessId

`func (o *ControlsData) SetSubprocessId(v int32)`

SetSubprocessId sets SubprocessId field to given value.

### HasSubprocessId

`func (o *ControlsData) HasSubprocessId() bool`

HasSubprocessId returns a boolean if a field has been set.

### GetLastModificationDate

`func (o *ControlsData) GetLastModificationDate() string`

GetLastModificationDate returns the LastModificationDate field if non-nil, zero value otherwise.

### GetLastModificationDateOk

`func (o *ControlsData) GetLastModificationDateOk() (*string, bool)`

GetLastModificationDateOk returns a tuple with the LastModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationDate

`func (o *ControlsData) SetLastModificationDate(v string)`

SetLastModificationDate sets LastModificationDate field to given value.

### HasLastModificationDate

`func (o *ControlsData) HasLastModificationDate() bool`

HasLastModificationDate returns a boolean if a field has been set.

### GetLastReviewDate

`func (o *ControlsData) GetLastReviewDate() string`

GetLastReviewDate returns the LastReviewDate field if non-nil, zero value otherwise.

### GetLastReviewDateOk

`func (o *ControlsData) GetLastReviewDateOk() (*string, bool)`

GetLastReviewDateOk returns a tuple with the LastReviewDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReviewDate

`func (o *ControlsData) SetLastReviewDate(v string)`

SetLastReviewDate sets LastReviewDate field to given value.

### HasLastReviewDate

`func (o *ControlsData) HasLastReviewDate() bool`

HasLastReviewDate returns a boolean if a field has been set.

### GetRelianceOptionId

`func (o *ControlsData) GetRelianceOptionId() int32`

GetRelianceOptionId returns the RelianceOptionId field if non-nil, zero value otherwise.

### GetRelianceOptionIdOk

`func (o *ControlsData) GetRelianceOptionIdOk() (*int32, bool)`

GetRelianceOptionIdOk returns a tuple with the RelianceOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelianceOptionId

`func (o *ControlsData) SetRelianceOptionId(v int32)`

SetRelianceOptionId sets RelianceOptionId field to given value.

### HasRelianceOptionId

`func (o *ControlsData) HasRelianceOptionId() bool`

HasRelianceOptionId returns a boolean if a field has been set.

### GetControlCustomSelect1OptionId

`func (o *ControlsData) GetControlCustomSelect1OptionId() int32`

GetControlCustomSelect1OptionId returns the ControlCustomSelect1OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect1OptionIdOk

`func (o *ControlsData) GetControlCustomSelect1OptionIdOk() (*int32, bool)`

GetControlCustomSelect1OptionIdOk returns a tuple with the ControlCustomSelect1OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect1OptionId

`func (o *ControlsData) SetControlCustomSelect1OptionId(v int32)`

SetControlCustomSelect1OptionId sets ControlCustomSelect1OptionId field to given value.

### HasControlCustomSelect1OptionId

`func (o *ControlsData) HasControlCustomSelect1OptionId() bool`

HasControlCustomSelect1OptionId returns a boolean if a field has been set.

### GetCoordinatorUserId

`func (o *ControlsData) GetCoordinatorUserId() int32`

GetCoordinatorUserId returns the CoordinatorUserId field if non-nil, zero value otherwise.

### GetCoordinatorUserIdOk

`func (o *ControlsData) GetCoordinatorUserIdOk() (*int32, bool)`

GetCoordinatorUserIdOk returns a tuple with the CoordinatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinatorUserId

`func (o *ControlsData) SetCoordinatorUserId(v int32)`

SetCoordinatorUserId sets CoordinatorUserId field to given value.

### HasCoordinatorUserId

`func (o *ControlsData) HasCoordinatorUserId() bool`

HasCoordinatorUserId returns a boolean if a field has been set.

### GetControlCustomText1

`func (o *ControlsData) GetControlCustomText1() string`

GetControlCustomText1 returns the ControlCustomText1 field if non-nil, zero value otherwise.

### GetControlCustomText1Ok

`func (o *ControlsData) GetControlCustomText1Ok() (*string, bool)`

GetControlCustomText1Ok returns a tuple with the ControlCustomText1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText1

`func (o *ControlsData) SetControlCustomText1(v string)`

SetControlCustomText1 sets ControlCustomText1 field to given value.

### HasControlCustomText1

`func (o *ControlsData) HasControlCustomText1() bool`

HasControlCustomText1 returns a boolean if a field has been set.

### GetControlCustomText2

`func (o *ControlsData) GetControlCustomText2() string`

GetControlCustomText2 returns the ControlCustomText2 field if non-nil, zero value otherwise.

### GetControlCustomText2Ok

`func (o *ControlsData) GetControlCustomText2Ok() (*string, bool)`

GetControlCustomText2Ok returns a tuple with the ControlCustomText2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText2

`func (o *ControlsData) SetControlCustomText2(v string)`

SetControlCustomText2 sets ControlCustomText2 field to given value.

### HasControlCustomText2

`func (o *ControlsData) HasControlCustomText2() bool`

HasControlCustomText2 returns a boolean if a field has been set.

### GetControlCustomText3

`func (o *ControlsData) GetControlCustomText3() string`

GetControlCustomText3 returns the ControlCustomText3 field if non-nil, zero value otherwise.

### GetControlCustomText3Ok

`func (o *ControlsData) GetControlCustomText3Ok() (*string, bool)`

GetControlCustomText3Ok returns a tuple with the ControlCustomText3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText3

`func (o *ControlsData) SetControlCustomText3(v string)`

SetControlCustomText3 sets ControlCustomText3 field to given value.

### HasControlCustomText3

`func (o *ControlsData) HasControlCustomText3() bool`

HasControlCustomText3 returns a boolean if a field has been set.

### GetControlCustomText4

`func (o *ControlsData) GetControlCustomText4() string`

GetControlCustomText4 returns the ControlCustomText4 field if non-nil, zero value otherwise.

### GetControlCustomText4Ok

`func (o *ControlsData) GetControlCustomText4Ok() (*string, bool)`

GetControlCustomText4Ok returns a tuple with the ControlCustomText4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText4

`func (o *ControlsData) SetControlCustomText4(v string)`

SetControlCustomText4 sets ControlCustomText4 field to given value.

### HasControlCustomText4

`func (o *ControlsData) HasControlCustomText4() bool`

HasControlCustomText4 returns a boolean if a field has been set.

### GetControlCustomSelect2OptionId

`func (o *ControlsData) GetControlCustomSelect2OptionId() int32`

GetControlCustomSelect2OptionId returns the ControlCustomSelect2OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect2OptionIdOk

`func (o *ControlsData) GetControlCustomSelect2OptionIdOk() (*int32, bool)`

GetControlCustomSelect2OptionIdOk returns a tuple with the ControlCustomSelect2OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect2OptionId

`func (o *ControlsData) SetControlCustomSelect2OptionId(v int32)`

SetControlCustomSelect2OptionId sets ControlCustomSelect2OptionId field to given value.

### HasControlCustomSelect2OptionId

`func (o *ControlsData) HasControlCustomSelect2OptionId() bool`

HasControlCustomSelect2OptionId returns a boolean if a field has been set.

### GetControlCustomSelect3OptionId

`func (o *ControlsData) GetControlCustomSelect3OptionId() int32`

GetControlCustomSelect3OptionId returns the ControlCustomSelect3OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect3OptionIdOk

`func (o *ControlsData) GetControlCustomSelect3OptionIdOk() (*int32, bool)`

GetControlCustomSelect3OptionIdOk returns a tuple with the ControlCustomSelect3OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect3OptionId

`func (o *ControlsData) SetControlCustomSelect3OptionId(v int32)`

SetControlCustomSelect3OptionId sets ControlCustomSelect3OptionId field to given value.

### HasControlCustomSelect3OptionId

`func (o *ControlsData) HasControlCustomSelect3OptionId() bool`

HasControlCustomSelect3OptionId returns a boolean if a field has been set.

### GetControlCustomSelect4OptionId

`func (o *ControlsData) GetControlCustomSelect4OptionId() int32`

GetControlCustomSelect4OptionId returns the ControlCustomSelect4OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect4OptionIdOk

`func (o *ControlsData) GetControlCustomSelect4OptionIdOk() (*int32, bool)`

GetControlCustomSelect4OptionIdOk returns a tuple with the ControlCustomSelect4OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect4OptionId

`func (o *ControlsData) SetControlCustomSelect4OptionId(v int32)`

SetControlCustomSelect4OptionId sets ControlCustomSelect4OptionId field to given value.

### HasControlCustomSelect4OptionId

`func (o *ControlsData) HasControlCustomSelect4OptionId() bool`

HasControlCustomSelect4OptionId returns a boolean if a field has been set.

### GetControlCustomSelect5OptionId

`func (o *ControlsData) GetControlCustomSelect5OptionId() int32`

GetControlCustomSelect5OptionId returns the ControlCustomSelect5OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect5OptionIdOk

`func (o *ControlsData) GetControlCustomSelect5OptionIdOk() (*int32, bool)`

GetControlCustomSelect5OptionIdOk returns a tuple with the ControlCustomSelect5OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect5OptionId

`func (o *ControlsData) SetControlCustomSelect5OptionId(v int32)`

SetControlCustomSelect5OptionId sets ControlCustomSelect5OptionId field to given value.

### HasControlCustomSelect5OptionId

`func (o *ControlsData) HasControlCustomSelect5OptionId() bool`

HasControlCustomSelect5OptionId returns a boolean if a field has been set.

### GetReferenceMeta

`func (o *ControlsData) GetReferenceMeta() interface{}`

GetReferenceMeta returns the ReferenceMeta field if non-nil, zero value otherwise.

### GetReferenceMetaOk

`func (o *ControlsData) GetReferenceMetaOk() (*interface{}, bool)`

GetReferenceMetaOk returns a tuple with the ReferenceMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMeta

`func (o *ControlsData) SetReferenceMeta(v interface{})`

SetReferenceMeta sets ReferenceMeta field to given value.


### SetReferenceMetaNil

`func (o *ControlsData) SetReferenceMetaNil(b bool)`

 SetReferenceMetaNil sets the value for ReferenceMeta to be an explicit nil

### UnsetReferenceMeta
`func (o *ControlsData) UnsetReferenceMeta()`

UnsetReferenceMeta ensures that no value is present for ReferenceMeta, not even an explicit nil
### GetComplianceReviewerUserId

`func (o *ControlsData) GetComplianceReviewerUserId() int32`

GetComplianceReviewerUserId returns the ComplianceReviewerUserId field if non-nil, zero value otherwise.

### GetComplianceReviewerUserIdOk

`func (o *ControlsData) GetComplianceReviewerUserIdOk() (*int32, bool)`

GetComplianceReviewerUserIdOk returns a tuple with the ComplianceReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceReviewerUserId

`func (o *ControlsData) SetComplianceReviewerUserId(v int32)`

SetComplianceReviewerUserId sets ComplianceReviewerUserId field to given value.

### HasComplianceReviewerUserId

`func (o *ControlsData) HasComplianceReviewerUserId() bool`

HasComplianceReviewerUserId returns a boolean if a field has been set.

### GetControlCustomText5

`func (o *ControlsData) GetControlCustomText5() string`

GetControlCustomText5 returns the ControlCustomText5 field if non-nil, zero value otherwise.

### GetControlCustomText5Ok

`func (o *ControlsData) GetControlCustomText5Ok() (*string, bool)`

GetControlCustomText5Ok returns a tuple with the ControlCustomText5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText5

`func (o *ControlsData) SetControlCustomText5(v string)`

SetControlCustomText5 sets ControlCustomText5 field to given value.

### HasControlCustomText5

`func (o *ControlsData) HasControlCustomText5() bool`

HasControlCustomText5 returns a boolean if a field has been set.

### GetControlCustomText6

`func (o *ControlsData) GetControlCustomText6() string`

GetControlCustomText6 returns the ControlCustomText6 field if non-nil, zero value otherwise.

### GetControlCustomText6Ok

`func (o *ControlsData) GetControlCustomText6Ok() (*string, bool)`

GetControlCustomText6Ok returns a tuple with the ControlCustomText6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText6

`func (o *ControlsData) SetControlCustomText6(v string)`

SetControlCustomText6 sets ControlCustomText6 field to given value.

### HasControlCustomText6

`func (o *ControlsData) HasControlCustomText6() bool`

HasControlCustomText6 returns a boolean if a field has been set.

### GetControlCustomText7

`func (o *ControlsData) GetControlCustomText7() string`

GetControlCustomText7 returns the ControlCustomText7 field if non-nil, zero value otherwise.

### GetControlCustomText7Ok

`func (o *ControlsData) GetControlCustomText7Ok() (*string, bool)`

GetControlCustomText7Ok returns a tuple with the ControlCustomText7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText7

`func (o *ControlsData) SetControlCustomText7(v string)`

SetControlCustomText7 sets ControlCustomText7 field to given value.

### HasControlCustomText7

`func (o *ControlsData) HasControlCustomText7() bool`

HasControlCustomText7 returns a boolean if a field has been set.

### GetControlCustomText8

`func (o *ControlsData) GetControlCustomText8() string`

GetControlCustomText8 returns the ControlCustomText8 field if non-nil, zero value otherwise.

### GetControlCustomText8Ok

`func (o *ControlsData) GetControlCustomText8Ok() (*string, bool)`

GetControlCustomText8Ok returns a tuple with the ControlCustomText8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText8

`func (o *ControlsData) SetControlCustomText8(v string)`

SetControlCustomText8 sets ControlCustomText8 field to given value.

### HasControlCustomText8

`func (o *ControlsData) HasControlCustomText8() bool`

HasControlCustomText8 returns a boolean if a field has been set.

### GetControlCustomText9

`func (o *ControlsData) GetControlCustomText9() string`

GetControlCustomText9 returns the ControlCustomText9 field if non-nil, zero value otherwise.

### GetControlCustomText9Ok

`func (o *ControlsData) GetControlCustomText9Ok() (*string, bool)`

GetControlCustomText9Ok returns a tuple with the ControlCustomText9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText9

`func (o *ControlsData) SetControlCustomText9(v string)`

SetControlCustomText9 sets ControlCustomText9 field to given value.

### HasControlCustomText9

`func (o *ControlsData) HasControlCustomText9() bool`

HasControlCustomText9 returns a boolean if a field has been set.

### GetControlCustomText10

`func (o *ControlsData) GetControlCustomText10() string`

GetControlCustomText10 returns the ControlCustomText10 field if non-nil, zero value otherwise.

### GetControlCustomText10Ok

`func (o *ControlsData) GetControlCustomText10Ok() (*string, bool)`

GetControlCustomText10Ok returns a tuple with the ControlCustomText10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText10

`func (o *ControlsData) SetControlCustomText10(v string)`

SetControlCustomText10 sets ControlCustomText10 field to given value.

### HasControlCustomText10

`func (o *ControlsData) HasControlCustomText10() bool`

HasControlCustomText10 returns a boolean if a field has been set.

### GetControlCustomText11

`func (o *ControlsData) GetControlCustomText11() string`

GetControlCustomText11 returns the ControlCustomText11 field if non-nil, zero value otherwise.

### GetControlCustomText11Ok

`func (o *ControlsData) GetControlCustomText11Ok() (*string, bool)`

GetControlCustomText11Ok returns a tuple with the ControlCustomText11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText11

`func (o *ControlsData) SetControlCustomText11(v string)`

SetControlCustomText11 sets ControlCustomText11 field to given value.

### HasControlCustomText11

`func (o *ControlsData) HasControlCustomText11() bool`

HasControlCustomText11 returns a boolean if a field has been set.

### GetControlCustomText12

`func (o *ControlsData) GetControlCustomText12() string`

GetControlCustomText12 returns the ControlCustomText12 field if non-nil, zero value otherwise.

### GetControlCustomText12Ok

`func (o *ControlsData) GetControlCustomText12Ok() (*string, bool)`

GetControlCustomText12Ok returns a tuple with the ControlCustomText12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText12

`func (o *ControlsData) SetControlCustomText12(v string)`

SetControlCustomText12 sets ControlCustomText12 field to given value.

### HasControlCustomText12

`func (o *ControlsData) HasControlCustomText12() bool`

HasControlCustomText12 returns a boolean if a field has been set.

### GetName

`func (o *ControlsData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControlsData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControlsData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ControlsData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetXuid

`func (o *ControlsData) GetXuid() string`

GetXuid returns the Xuid field if non-nil, zero value otherwise.

### GetXuidOk

`func (o *ControlsData) GetXuidOk() (*string, bool)`

GetXuidOk returns a tuple with the Xuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXuid

`func (o *ControlsData) SetXuid(v string)`

SetXuid sets Xuid field to given value.

### HasXuid

`func (o *ControlsData) HasXuid() bool`

HasXuid returns a boolean if a field has been set.

### GetScopes

`func (o *ControlsData) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ControlsData) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ControlsData) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.


### SetScopesNil

`func (o *ControlsData) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *ControlsData) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetFieldData

`func (o *ControlsData) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *ControlsData) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *ControlsData) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *ControlsData) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *ControlsData) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *ControlsData) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil
### GetControlCustomSelect6OptionId

`func (o *ControlsData) GetControlCustomSelect6OptionId() int32`

GetControlCustomSelect6OptionId returns the ControlCustomSelect6OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect6OptionIdOk

`func (o *ControlsData) GetControlCustomSelect6OptionIdOk() (*int32, bool)`

GetControlCustomSelect6OptionIdOk returns a tuple with the ControlCustomSelect6OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect6OptionId

`func (o *ControlsData) SetControlCustomSelect6OptionId(v int32)`

SetControlCustomSelect6OptionId sets ControlCustomSelect6OptionId field to given value.

### HasControlCustomSelect6OptionId

`func (o *ControlsData) HasControlCustomSelect6OptionId() bool`

HasControlCustomSelect6OptionId returns a boolean if a field has been set.

### GetControlCustomSelect7OptionId

`func (o *ControlsData) GetControlCustomSelect7OptionId() int32`

GetControlCustomSelect7OptionId returns the ControlCustomSelect7OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect7OptionIdOk

`func (o *ControlsData) GetControlCustomSelect7OptionIdOk() (*int32, bool)`

GetControlCustomSelect7OptionIdOk returns a tuple with the ControlCustomSelect7OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect7OptionId

`func (o *ControlsData) SetControlCustomSelect7OptionId(v int32)`

SetControlCustomSelect7OptionId sets ControlCustomSelect7OptionId field to given value.

### HasControlCustomSelect7OptionId

`func (o *ControlsData) HasControlCustomSelect7OptionId() bool`

HasControlCustomSelect7OptionId returns a boolean if a field has been set.

### GetControlCustomSelect8OptionId

`func (o *ControlsData) GetControlCustomSelect8OptionId() int32`

GetControlCustomSelect8OptionId returns the ControlCustomSelect8OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect8OptionIdOk

`func (o *ControlsData) GetControlCustomSelect8OptionIdOk() (*int32, bool)`

GetControlCustomSelect8OptionIdOk returns a tuple with the ControlCustomSelect8OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect8OptionId

`func (o *ControlsData) SetControlCustomSelect8OptionId(v int32)`

SetControlCustomSelect8OptionId sets ControlCustomSelect8OptionId field to given value.

### HasControlCustomSelect8OptionId

`func (o *ControlsData) HasControlCustomSelect8OptionId() bool`

HasControlCustomSelect8OptionId returns a boolean if a field has been set.

### GetControlCustomSelect9OptionId

`func (o *ControlsData) GetControlCustomSelect9OptionId() int32`

GetControlCustomSelect9OptionId returns the ControlCustomSelect9OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect9OptionIdOk

`func (o *ControlsData) GetControlCustomSelect9OptionIdOk() (*int32, bool)`

GetControlCustomSelect9OptionIdOk returns a tuple with the ControlCustomSelect9OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect9OptionId

`func (o *ControlsData) SetControlCustomSelect9OptionId(v int32)`

SetControlCustomSelect9OptionId sets ControlCustomSelect9OptionId field to given value.

### HasControlCustomSelect9OptionId

`func (o *ControlsData) HasControlCustomSelect9OptionId() bool`

HasControlCustomSelect9OptionId returns a boolean if a field has been set.

### GetControlCustomSelect10OptionId

`func (o *ControlsData) GetControlCustomSelect10OptionId() int32`

GetControlCustomSelect10OptionId returns the ControlCustomSelect10OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect10OptionIdOk

`func (o *ControlsData) GetControlCustomSelect10OptionIdOk() (*int32, bool)`

GetControlCustomSelect10OptionIdOk returns a tuple with the ControlCustomSelect10OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect10OptionId

`func (o *ControlsData) SetControlCustomSelect10OptionId(v int32)`

SetControlCustomSelect10OptionId sets ControlCustomSelect10OptionId field to given value.

### HasControlCustomSelect10OptionId

`func (o *ControlsData) HasControlCustomSelect10OptionId() bool`

HasControlCustomSelect10OptionId returns a boolean if a field has been set.

### GetControlCustomSelect11OptionId

`func (o *ControlsData) GetControlCustomSelect11OptionId() int32`

GetControlCustomSelect11OptionId returns the ControlCustomSelect11OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect11OptionIdOk

`func (o *ControlsData) GetControlCustomSelect11OptionIdOk() (*int32, bool)`

GetControlCustomSelect11OptionIdOk returns a tuple with the ControlCustomSelect11OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect11OptionId

`func (o *ControlsData) SetControlCustomSelect11OptionId(v int32)`

SetControlCustomSelect11OptionId sets ControlCustomSelect11OptionId field to given value.

### HasControlCustomSelect11OptionId

`func (o *ControlsData) HasControlCustomSelect11OptionId() bool`

HasControlCustomSelect11OptionId returns a boolean if a field has been set.

### GetControlCustomSelect12OptionId

`func (o *ControlsData) GetControlCustomSelect12OptionId() int32`

GetControlCustomSelect12OptionId returns the ControlCustomSelect12OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect12OptionIdOk

`func (o *ControlsData) GetControlCustomSelect12OptionIdOk() (*int32, bool)`

GetControlCustomSelect12OptionIdOk returns a tuple with the ControlCustomSelect12OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect12OptionId

`func (o *ControlsData) SetControlCustomSelect12OptionId(v int32)`

SetControlCustomSelect12OptionId sets ControlCustomSelect12OptionId field to given value.

### HasControlCustomSelect12OptionId

`func (o *ControlsData) HasControlCustomSelect12OptionId() bool`

HasControlCustomSelect12OptionId returns a boolean if a field has been set.

### GetCustomUserSelect1UserId

`func (o *ControlsData) GetCustomUserSelect1UserId() int32`

GetCustomUserSelect1UserId returns the CustomUserSelect1UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect1UserIdOk

`func (o *ControlsData) GetCustomUserSelect1UserIdOk() (*int32, bool)`

GetCustomUserSelect1UserIdOk returns a tuple with the CustomUserSelect1UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect1UserId

`func (o *ControlsData) SetCustomUserSelect1UserId(v int32)`

SetCustomUserSelect1UserId sets CustomUserSelect1UserId field to given value.

### HasCustomUserSelect1UserId

`func (o *ControlsData) HasCustomUserSelect1UserId() bool`

HasCustomUserSelect1UserId returns a boolean if a field has been set.

### GetCustomUserSelect2UserId

`func (o *ControlsData) GetCustomUserSelect2UserId() int32`

GetCustomUserSelect2UserId returns the CustomUserSelect2UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect2UserIdOk

`func (o *ControlsData) GetCustomUserSelect2UserIdOk() (*int32, bool)`

GetCustomUserSelect2UserIdOk returns a tuple with the CustomUserSelect2UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect2UserId

`func (o *ControlsData) SetCustomUserSelect2UserId(v int32)`

SetCustomUserSelect2UserId sets CustomUserSelect2UserId field to given value.

### HasCustomUserSelect2UserId

`func (o *ControlsData) HasCustomUserSelect2UserId() bool`

HasCustomUserSelect2UserId returns a boolean if a field has been set.

### GetCustomUserSelect3UserId

`func (o *ControlsData) GetCustomUserSelect3UserId() int32`

GetCustomUserSelect3UserId returns the CustomUserSelect3UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect3UserIdOk

`func (o *ControlsData) GetCustomUserSelect3UserIdOk() (*int32, bool)`

GetCustomUserSelect3UserIdOk returns a tuple with the CustomUserSelect3UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect3UserId

`func (o *ControlsData) SetCustomUserSelect3UserId(v int32)`

SetCustomUserSelect3UserId sets CustomUserSelect3UserId field to given value.

### HasCustomUserSelect3UserId

`func (o *ControlsData) HasCustomUserSelect3UserId() bool`

HasCustomUserSelect3UserId returns a boolean if a field has been set.

### GetCustomUserSelect4UserId

`func (o *ControlsData) GetCustomUserSelect4UserId() int32`

GetCustomUserSelect4UserId returns the CustomUserSelect4UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect4UserIdOk

`func (o *ControlsData) GetCustomUserSelect4UserIdOk() (*int32, bool)`

GetCustomUserSelect4UserIdOk returns a tuple with the CustomUserSelect4UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect4UserId

`func (o *ControlsData) SetCustomUserSelect4UserId(v int32)`

SetCustomUserSelect4UserId sets CustomUserSelect4UserId field to given value.

### HasCustomUserSelect4UserId

`func (o *ControlsData) HasCustomUserSelect4UserId() bool`

HasCustomUserSelect4UserId returns a boolean if a field has been set.

### GetControlCustomText13

`func (o *ControlsData) GetControlCustomText13() string`

GetControlCustomText13 returns the ControlCustomText13 field if non-nil, zero value otherwise.

### GetControlCustomText13Ok

`func (o *ControlsData) GetControlCustomText13Ok() (*string, bool)`

GetControlCustomText13Ok returns a tuple with the ControlCustomText13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText13

`func (o *ControlsData) SetControlCustomText13(v string)`

SetControlCustomText13 sets ControlCustomText13 field to given value.

### HasControlCustomText13

`func (o *ControlsData) HasControlCustomText13() bool`

HasControlCustomText13 returns a boolean if a field has been set.

### GetControlCustomText14

`func (o *ControlsData) GetControlCustomText14() string`

GetControlCustomText14 returns the ControlCustomText14 field if non-nil, zero value otherwise.

### GetControlCustomText14Ok

`func (o *ControlsData) GetControlCustomText14Ok() (*string, bool)`

GetControlCustomText14Ok returns a tuple with the ControlCustomText14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText14

`func (o *ControlsData) SetControlCustomText14(v string)`

SetControlCustomText14 sets ControlCustomText14 field to given value.

### HasControlCustomText14

`func (o *ControlsData) HasControlCustomText14() bool`

HasControlCustomText14 returns a boolean if a field has been set.

### GetControlCustomText15

`func (o *ControlsData) GetControlCustomText15() string`

GetControlCustomText15 returns the ControlCustomText15 field if non-nil, zero value otherwise.

### GetControlCustomText15Ok

`func (o *ControlsData) GetControlCustomText15Ok() (*string, bool)`

GetControlCustomText15Ok returns a tuple with the ControlCustomText15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText15

`func (o *ControlsData) SetControlCustomText15(v string)`

SetControlCustomText15 sets ControlCustomText15 field to given value.

### HasControlCustomText15

`func (o *ControlsData) HasControlCustomText15() bool`

HasControlCustomText15 returns a boolean if a field has been set.

### GetControlCustomText16

`func (o *ControlsData) GetControlCustomText16() string`

GetControlCustomText16 returns the ControlCustomText16 field if non-nil, zero value otherwise.

### GetControlCustomText16Ok

`func (o *ControlsData) GetControlCustomText16Ok() (*string, bool)`

GetControlCustomText16Ok returns a tuple with the ControlCustomText16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText16

`func (o *ControlsData) SetControlCustomText16(v string)`

SetControlCustomText16 sets ControlCustomText16 field to given value.

### HasControlCustomText16

`func (o *ControlsData) HasControlCustomText16() bool`

HasControlCustomText16 returns a boolean if a field has been set.

### GetControlCustomText17

`func (o *ControlsData) GetControlCustomText17() string`

GetControlCustomText17 returns the ControlCustomText17 field if non-nil, zero value otherwise.

### GetControlCustomText17Ok

`func (o *ControlsData) GetControlCustomText17Ok() (*string, bool)`

GetControlCustomText17Ok returns a tuple with the ControlCustomText17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText17

`func (o *ControlsData) SetControlCustomText17(v string)`

SetControlCustomText17 sets ControlCustomText17 field to given value.

### HasControlCustomText17

`func (o *ControlsData) HasControlCustomText17() bool`

HasControlCustomText17 returns a boolean if a field has been set.

### GetControlCustomText18

`func (o *ControlsData) GetControlCustomText18() string`

GetControlCustomText18 returns the ControlCustomText18 field if non-nil, zero value otherwise.

### GetControlCustomText18Ok

`func (o *ControlsData) GetControlCustomText18Ok() (*string, bool)`

GetControlCustomText18Ok returns a tuple with the ControlCustomText18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText18

`func (o *ControlsData) SetControlCustomText18(v string)`

SetControlCustomText18 sets ControlCustomText18 field to given value.

### HasControlCustomText18

`func (o *ControlsData) HasControlCustomText18() bool`

HasControlCustomText18 returns a boolean if a field has been set.

### GetControlCustomText19

`func (o *ControlsData) GetControlCustomText19() string`

GetControlCustomText19 returns the ControlCustomText19 field if non-nil, zero value otherwise.

### GetControlCustomText19Ok

`func (o *ControlsData) GetControlCustomText19Ok() (*string, bool)`

GetControlCustomText19Ok returns a tuple with the ControlCustomText19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText19

`func (o *ControlsData) SetControlCustomText19(v string)`

SetControlCustomText19 sets ControlCustomText19 field to given value.

### HasControlCustomText19

`func (o *ControlsData) HasControlCustomText19() bool`

HasControlCustomText19 returns a boolean if a field has been set.

### GetControlCustomText20

`func (o *ControlsData) GetControlCustomText20() string`

GetControlCustomText20 returns the ControlCustomText20 field if non-nil, zero value otherwise.

### GetControlCustomText20Ok

`func (o *ControlsData) GetControlCustomText20Ok() (*string, bool)`

GetControlCustomText20Ok returns a tuple with the ControlCustomText20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText20

`func (o *ControlsData) SetControlCustomText20(v string)`

SetControlCustomText20 sets ControlCustomText20 field to given value.

### HasControlCustomText20

`func (o *ControlsData) HasControlCustomText20() bool`

HasControlCustomText20 returns a boolean if a field has been set.

### GetControlCustomText21

`func (o *ControlsData) GetControlCustomText21() string`

GetControlCustomText21 returns the ControlCustomText21 field if non-nil, zero value otherwise.

### GetControlCustomText21Ok

`func (o *ControlsData) GetControlCustomText21Ok() (*string, bool)`

GetControlCustomText21Ok returns a tuple with the ControlCustomText21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText21

`func (o *ControlsData) SetControlCustomText21(v string)`

SetControlCustomText21 sets ControlCustomText21 field to given value.

### HasControlCustomText21

`func (o *ControlsData) HasControlCustomText21() bool`

HasControlCustomText21 returns a boolean if a field has been set.

### GetControlCustomText22

`func (o *ControlsData) GetControlCustomText22() string`

GetControlCustomText22 returns the ControlCustomText22 field if non-nil, zero value otherwise.

### GetControlCustomText22Ok

`func (o *ControlsData) GetControlCustomText22Ok() (*string, bool)`

GetControlCustomText22Ok returns a tuple with the ControlCustomText22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText22

`func (o *ControlsData) SetControlCustomText22(v string)`

SetControlCustomText22 sets ControlCustomText22 field to given value.

### HasControlCustomText22

`func (o *ControlsData) HasControlCustomText22() bool`

HasControlCustomText22 returns a boolean if a field has been set.

### GetControlCustomText23

`func (o *ControlsData) GetControlCustomText23() string`

GetControlCustomText23 returns the ControlCustomText23 field if non-nil, zero value otherwise.

### GetControlCustomText23Ok

`func (o *ControlsData) GetControlCustomText23Ok() (*string, bool)`

GetControlCustomText23Ok returns a tuple with the ControlCustomText23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText23

`func (o *ControlsData) SetControlCustomText23(v string)`

SetControlCustomText23 sets ControlCustomText23 field to given value.

### HasControlCustomText23

`func (o *ControlsData) HasControlCustomText23() bool`

HasControlCustomText23 returns a boolean if a field has been set.

### GetControlCustomText24

`func (o *ControlsData) GetControlCustomText24() string`

GetControlCustomText24 returns the ControlCustomText24 field if non-nil, zero value otherwise.

### GetControlCustomText24Ok

`func (o *ControlsData) GetControlCustomText24Ok() (*string, bool)`

GetControlCustomText24Ok returns a tuple with the ControlCustomText24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText24

`func (o *ControlsData) SetControlCustomText24(v string)`

SetControlCustomText24 sets ControlCustomText24 field to given value.

### HasControlCustomText24

`func (o *ControlsData) HasControlCustomText24() bool`

HasControlCustomText24 returns a boolean if a field has been set.

### GetCustomDate1

`func (o *ControlsData) GetCustomDate1() string`

GetCustomDate1 returns the CustomDate1 field if non-nil, zero value otherwise.

### GetCustomDate1Ok

`func (o *ControlsData) GetCustomDate1Ok() (*string, bool)`

GetCustomDate1Ok returns a tuple with the CustomDate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate1

`func (o *ControlsData) SetCustomDate1(v string)`

SetCustomDate1 sets CustomDate1 field to given value.

### HasCustomDate1

`func (o *ControlsData) HasCustomDate1() bool`

HasCustomDate1 returns a boolean if a field has been set.

### GetCustomDate2

`func (o *ControlsData) GetCustomDate2() string`

GetCustomDate2 returns the CustomDate2 field if non-nil, zero value otherwise.

### GetCustomDate2Ok

`func (o *ControlsData) GetCustomDate2Ok() (*string, bool)`

GetCustomDate2Ok returns a tuple with the CustomDate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate2

`func (o *ControlsData) SetCustomDate2(v string)`

SetCustomDate2 sets CustomDate2 field to given value.

### HasCustomDate2

`func (o *ControlsData) HasCustomDate2() bool`

HasCustomDate2 returns a boolean if a field has been set.

### GetCustomDate3

`func (o *ControlsData) GetCustomDate3() string`

GetCustomDate3 returns the CustomDate3 field if non-nil, zero value otherwise.

### GetCustomDate3Ok

`func (o *ControlsData) GetCustomDate3Ok() (*string, bool)`

GetCustomDate3Ok returns a tuple with the CustomDate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate3

`func (o *ControlsData) SetCustomDate3(v string)`

SetCustomDate3 sets CustomDate3 field to given value.

### HasCustomDate3

`func (o *ControlsData) HasCustomDate3() bool`

HasCustomDate3 returns a boolean if a field has been set.

### GetCustomDate4

`func (o *ControlsData) GetCustomDate4() string`

GetCustomDate4 returns the CustomDate4 field if non-nil, zero value otherwise.

### GetCustomDate4Ok

`func (o *ControlsData) GetCustomDate4Ok() (*string, bool)`

GetCustomDate4Ok returns a tuple with the CustomDate4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate4

`func (o *ControlsData) SetCustomDate4(v string)`

SetCustomDate4 sets CustomDate4 field to given value.

### HasCustomDate4

`func (o *ControlsData) HasCustomDate4() bool`

HasCustomDate4 returns a boolean if a field has been set.

### GetCustomDate5

`func (o *ControlsData) GetCustomDate5() string`

GetCustomDate5 returns the CustomDate5 field if non-nil, zero value otherwise.

### GetCustomDate5Ok

`func (o *ControlsData) GetCustomDate5Ok() (*string, bool)`

GetCustomDate5Ok returns a tuple with the CustomDate5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate5

`func (o *ControlsData) SetCustomDate5(v string)`

SetCustomDate5 sets CustomDate5 field to given value.

### HasCustomDate5

`func (o *ControlsData) HasCustomDate5() bool`

HasCustomDate5 returns a boolean if a field has been set.

### GetCustomUserSelect5UserId

`func (o *ControlsData) GetCustomUserSelect5UserId() int32`

GetCustomUserSelect5UserId returns the CustomUserSelect5UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect5UserIdOk

`func (o *ControlsData) GetCustomUserSelect5UserIdOk() (*int32, bool)`

GetCustomUserSelect5UserIdOk returns a tuple with the CustomUserSelect5UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect5UserId

`func (o *ControlsData) SetCustomUserSelect5UserId(v int32)`

SetCustomUserSelect5UserId sets CustomUserSelect5UserId field to given value.

### HasCustomUserSelect5UserId

`func (o *ControlsData) HasCustomUserSelect5UserId() bool`

HasCustomUserSelect5UserId returns a boolean if a field has been set.

### GetCustomUserSelect6UserId

`func (o *ControlsData) GetCustomUserSelect6UserId() int32`

GetCustomUserSelect6UserId returns the CustomUserSelect6UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect6UserIdOk

`func (o *ControlsData) GetCustomUserSelect6UserIdOk() (*int32, bool)`

GetCustomUserSelect6UserIdOk returns a tuple with the CustomUserSelect6UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect6UserId

`func (o *ControlsData) SetCustomUserSelect6UserId(v int32)`

SetCustomUserSelect6UserId sets CustomUserSelect6UserId field to given value.

### HasCustomUserSelect6UserId

`func (o *ControlsData) HasCustomUserSelect6UserId() bool`

HasCustomUserSelect6UserId returns a boolean if a field has been set.

### GetCustomUserSelect7UserId

`func (o *ControlsData) GetCustomUserSelect7UserId() int32`

GetCustomUserSelect7UserId returns the CustomUserSelect7UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect7UserIdOk

`func (o *ControlsData) GetCustomUserSelect7UserIdOk() (*int32, bool)`

GetCustomUserSelect7UserIdOk returns a tuple with the CustomUserSelect7UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect7UserId

`func (o *ControlsData) SetCustomUserSelect7UserId(v int32)`

SetCustomUserSelect7UserId sets CustomUserSelect7UserId field to given value.

### HasCustomUserSelect7UserId

`func (o *ControlsData) HasCustomUserSelect7UserId() bool`

HasCustomUserSelect7UserId returns a boolean if a field has been set.

### GetCustomUserSelect8UserId

`func (o *ControlsData) GetCustomUserSelect8UserId() int32`

GetCustomUserSelect8UserId returns the CustomUserSelect8UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect8UserIdOk

`func (o *ControlsData) GetCustomUserSelect8UserIdOk() (*int32, bool)`

GetCustomUserSelect8UserIdOk returns a tuple with the CustomUserSelect8UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect8UserId

`func (o *ControlsData) SetCustomUserSelect8UserId(v int32)`

SetCustomUserSelect8UserId sets CustomUserSelect8UserId field to given value.

### HasCustomUserSelect8UserId

`func (o *ControlsData) HasCustomUserSelect8UserId() bool`

HasCustomUserSelect8UserId returns a boolean if a field has been set.

### GetCustomUserSelect9UserId

`func (o *ControlsData) GetCustomUserSelect9UserId() int32`

GetCustomUserSelect9UserId returns the CustomUserSelect9UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect9UserIdOk

`func (o *ControlsData) GetCustomUserSelect9UserIdOk() (*int32, bool)`

GetCustomUserSelect9UserIdOk returns a tuple with the CustomUserSelect9UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect9UserId

`func (o *ControlsData) SetCustomUserSelect9UserId(v int32)`

SetCustomUserSelect9UserId sets CustomUserSelect9UserId field to given value.

### HasCustomUserSelect9UserId

`func (o *ControlsData) HasCustomUserSelect9UserId() bool`

HasCustomUserSelect9UserId returns a boolean if a field has been set.

### GetCustomUserSelect10UserId

`func (o *ControlsData) GetCustomUserSelect10UserId() int32`

GetCustomUserSelect10UserId returns the CustomUserSelect10UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect10UserIdOk

`func (o *ControlsData) GetCustomUserSelect10UserIdOk() (*int32, bool)`

GetCustomUserSelect10UserIdOk returns a tuple with the CustomUserSelect10UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect10UserId

`func (o *ControlsData) SetCustomUserSelect10UserId(v int32)`

SetCustomUserSelect10UserId sets CustomUserSelect10UserId field to given value.

### HasCustomUserSelect10UserId

`func (o *ControlsData) HasCustomUserSelect10UserId() bool`

HasCustomUserSelect10UserId returns a boolean if a field has been set.

### GetControlCustomText25

`func (o *ControlsData) GetControlCustomText25() string`

GetControlCustomText25 returns the ControlCustomText25 field if non-nil, zero value otherwise.

### GetControlCustomText25Ok

`func (o *ControlsData) GetControlCustomText25Ok() (*string, bool)`

GetControlCustomText25Ok returns a tuple with the ControlCustomText25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText25

`func (o *ControlsData) SetControlCustomText25(v string)`

SetControlCustomText25 sets ControlCustomText25 field to given value.

### HasControlCustomText25

`func (o *ControlsData) HasControlCustomText25() bool`

HasControlCustomText25 returns a boolean if a field has been set.

### GetControlCustomText26

`func (o *ControlsData) GetControlCustomText26() string`

GetControlCustomText26 returns the ControlCustomText26 field if non-nil, zero value otherwise.

### GetControlCustomText26Ok

`func (o *ControlsData) GetControlCustomText26Ok() (*string, bool)`

GetControlCustomText26Ok returns a tuple with the ControlCustomText26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText26

`func (o *ControlsData) SetControlCustomText26(v string)`

SetControlCustomText26 sets ControlCustomText26 field to given value.

### HasControlCustomText26

`func (o *ControlsData) HasControlCustomText26() bool`

HasControlCustomText26 returns a boolean if a field has been set.

### GetControlCustomText27

`func (o *ControlsData) GetControlCustomText27() string`

GetControlCustomText27 returns the ControlCustomText27 field if non-nil, zero value otherwise.

### GetControlCustomText27Ok

`func (o *ControlsData) GetControlCustomText27Ok() (*string, bool)`

GetControlCustomText27Ok returns a tuple with the ControlCustomText27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText27

`func (o *ControlsData) SetControlCustomText27(v string)`

SetControlCustomText27 sets ControlCustomText27 field to given value.

### HasControlCustomText27

`func (o *ControlsData) HasControlCustomText27() bool`

HasControlCustomText27 returns a boolean if a field has been set.

### GetControlCustomText28

`func (o *ControlsData) GetControlCustomText28() string`

GetControlCustomText28 returns the ControlCustomText28 field if non-nil, zero value otherwise.

### GetControlCustomText28Ok

`func (o *ControlsData) GetControlCustomText28Ok() (*string, bool)`

GetControlCustomText28Ok returns a tuple with the ControlCustomText28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText28

`func (o *ControlsData) SetControlCustomText28(v string)`

SetControlCustomText28 sets ControlCustomText28 field to given value.

### HasControlCustomText28

`func (o *ControlsData) HasControlCustomText28() bool`

HasControlCustomText28 returns a boolean if a field has been set.

### GetControlCustomText29

`func (o *ControlsData) GetControlCustomText29() string`

GetControlCustomText29 returns the ControlCustomText29 field if non-nil, zero value otherwise.

### GetControlCustomText29Ok

`func (o *ControlsData) GetControlCustomText29Ok() (*string, bool)`

GetControlCustomText29Ok returns a tuple with the ControlCustomText29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText29

`func (o *ControlsData) SetControlCustomText29(v string)`

SetControlCustomText29 sets ControlCustomText29 field to given value.

### HasControlCustomText29

`func (o *ControlsData) HasControlCustomText29() bool`

HasControlCustomText29 returns a boolean if a field has been set.

### GetControlCustomText30

`func (o *ControlsData) GetControlCustomText30() string`

GetControlCustomText30 returns the ControlCustomText30 field if non-nil, zero value otherwise.

### GetControlCustomText30Ok

`func (o *ControlsData) GetControlCustomText30Ok() (*string, bool)`

GetControlCustomText30Ok returns a tuple with the ControlCustomText30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText30

`func (o *ControlsData) SetControlCustomText30(v string)`

SetControlCustomText30 sets ControlCustomText30 field to given value.

### HasControlCustomText30

`func (o *ControlsData) HasControlCustomText30() bool`

HasControlCustomText30 returns a boolean if a field has been set.

### GetControlCustomText31

`func (o *ControlsData) GetControlCustomText31() string`

GetControlCustomText31 returns the ControlCustomText31 field if non-nil, zero value otherwise.

### GetControlCustomText31Ok

`func (o *ControlsData) GetControlCustomText31Ok() (*string, bool)`

GetControlCustomText31Ok returns a tuple with the ControlCustomText31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText31

`func (o *ControlsData) SetControlCustomText31(v string)`

SetControlCustomText31 sets ControlCustomText31 field to given value.

### HasControlCustomText31

`func (o *ControlsData) HasControlCustomText31() bool`

HasControlCustomText31 returns a boolean if a field has been set.

### GetControlCustomText32

`func (o *ControlsData) GetControlCustomText32() string`

GetControlCustomText32 returns the ControlCustomText32 field if non-nil, zero value otherwise.

### GetControlCustomText32Ok

`func (o *ControlsData) GetControlCustomText32Ok() (*string, bool)`

GetControlCustomText32Ok returns a tuple with the ControlCustomText32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText32

`func (o *ControlsData) SetControlCustomText32(v string)`

SetControlCustomText32 sets ControlCustomText32 field to given value.

### HasControlCustomText32

`func (o *ControlsData) HasControlCustomText32() bool`

HasControlCustomText32 returns a boolean if a field has been set.

### GetControlCustomSelect13OptionId

`func (o *ControlsData) GetControlCustomSelect13OptionId() int32`

GetControlCustomSelect13OptionId returns the ControlCustomSelect13OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect13OptionIdOk

`func (o *ControlsData) GetControlCustomSelect13OptionIdOk() (*int32, bool)`

GetControlCustomSelect13OptionIdOk returns a tuple with the ControlCustomSelect13OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect13OptionId

`func (o *ControlsData) SetControlCustomSelect13OptionId(v int32)`

SetControlCustomSelect13OptionId sets ControlCustomSelect13OptionId field to given value.

### HasControlCustomSelect13OptionId

`func (o *ControlsData) HasControlCustomSelect13OptionId() bool`

HasControlCustomSelect13OptionId returns a boolean if a field has been set.

### GetControlCustomSelect14OptionId

`func (o *ControlsData) GetControlCustomSelect14OptionId() int32`

GetControlCustomSelect14OptionId returns the ControlCustomSelect14OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect14OptionIdOk

`func (o *ControlsData) GetControlCustomSelect14OptionIdOk() (*int32, bool)`

GetControlCustomSelect14OptionIdOk returns a tuple with the ControlCustomSelect14OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect14OptionId

`func (o *ControlsData) SetControlCustomSelect14OptionId(v int32)`

SetControlCustomSelect14OptionId sets ControlCustomSelect14OptionId field to given value.

### HasControlCustomSelect14OptionId

`func (o *ControlsData) HasControlCustomSelect14OptionId() bool`

HasControlCustomSelect14OptionId returns a boolean if a field has been set.

### GetControlCustomSelect15OptionId

`func (o *ControlsData) GetControlCustomSelect15OptionId() int32`

GetControlCustomSelect15OptionId returns the ControlCustomSelect15OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect15OptionIdOk

`func (o *ControlsData) GetControlCustomSelect15OptionIdOk() (*int32, bool)`

GetControlCustomSelect15OptionIdOk returns a tuple with the ControlCustomSelect15OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect15OptionId

`func (o *ControlsData) SetControlCustomSelect15OptionId(v int32)`

SetControlCustomSelect15OptionId sets ControlCustomSelect15OptionId field to given value.

### HasControlCustomSelect15OptionId

`func (o *ControlsData) HasControlCustomSelect15OptionId() bool`

HasControlCustomSelect15OptionId returns a boolean if a field has been set.

### GetControlCustomSelect16OptionId

`func (o *ControlsData) GetControlCustomSelect16OptionId() int32`

GetControlCustomSelect16OptionId returns the ControlCustomSelect16OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect16OptionIdOk

`func (o *ControlsData) GetControlCustomSelect16OptionIdOk() (*int32, bool)`

GetControlCustomSelect16OptionIdOk returns a tuple with the ControlCustomSelect16OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect16OptionId

`func (o *ControlsData) SetControlCustomSelect16OptionId(v int32)`

SetControlCustomSelect16OptionId sets ControlCustomSelect16OptionId field to given value.

### HasControlCustomSelect16OptionId

`func (o *ControlsData) HasControlCustomSelect16OptionId() bool`

HasControlCustomSelect16OptionId returns a boolean if a field has been set.

### GetControlsSegmentId

`func (o *ControlsData) GetControlsSegmentId() int32`

GetControlsSegmentId returns the ControlsSegmentId field if non-nil, zero value otherwise.

### GetControlsSegmentIdOk

`func (o *ControlsData) GetControlsSegmentIdOk() (*int32, bool)`

GetControlsSegmentIdOk returns a tuple with the ControlsSegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlsSegmentId

`func (o *ControlsData) SetControlsSegmentId(v int32)`

SetControlsSegmentId sets ControlsSegmentId field to given value.

### HasControlsSegmentId

`func (o *ControlsData) HasControlsSegmentId() bool`

HasControlsSegmentId returns a boolean if a field has been set.

### GetControlCustomSelect17OptionId

`func (o *ControlsData) GetControlCustomSelect17OptionId() int32`

GetControlCustomSelect17OptionId returns the ControlCustomSelect17OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect17OptionIdOk

`func (o *ControlsData) GetControlCustomSelect17OptionIdOk() (*int32, bool)`

GetControlCustomSelect17OptionIdOk returns a tuple with the ControlCustomSelect17OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect17OptionId

`func (o *ControlsData) SetControlCustomSelect17OptionId(v int32)`

SetControlCustomSelect17OptionId sets ControlCustomSelect17OptionId field to given value.

### HasControlCustomSelect17OptionId

`func (o *ControlsData) HasControlCustomSelect17OptionId() bool`

HasControlCustomSelect17OptionId returns a boolean if a field has been set.

### GetControlCustomSelect18OptionId

`func (o *ControlsData) GetControlCustomSelect18OptionId() int32`

GetControlCustomSelect18OptionId returns the ControlCustomSelect18OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect18OptionIdOk

`func (o *ControlsData) GetControlCustomSelect18OptionIdOk() (*int32, bool)`

GetControlCustomSelect18OptionIdOk returns a tuple with the ControlCustomSelect18OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect18OptionId

`func (o *ControlsData) SetControlCustomSelect18OptionId(v int32)`

SetControlCustomSelect18OptionId sets ControlCustomSelect18OptionId field to given value.

### HasControlCustomSelect18OptionId

`func (o *ControlsData) HasControlCustomSelect18OptionId() bool`

HasControlCustomSelect18OptionId returns a boolean if a field has been set.

### GetControlCustomSelect19OptionId

`func (o *ControlsData) GetControlCustomSelect19OptionId() int32`

GetControlCustomSelect19OptionId returns the ControlCustomSelect19OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect19OptionIdOk

`func (o *ControlsData) GetControlCustomSelect19OptionIdOk() (*int32, bool)`

GetControlCustomSelect19OptionIdOk returns a tuple with the ControlCustomSelect19OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect19OptionId

`func (o *ControlsData) SetControlCustomSelect19OptionId(v int32)`

SetControlCustomSelect19OptionId sets ControlCustomSelect19OptionId field to given value.

### HasControlCustomSelect19OptionId

`func (o *ControlsData) HasControlCustomSelect19OptionId() bool`

HasControlCustomSelect19OptionId returns a boolean if a field has been set.

### GetControlCustomSelect20OptionId

`func (o *ControlsData) GetControlCustomSelect20OptionId() int32`

GetControlCustomSelect20OptionId returns the ControlCustomSelect20OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect20OptionIdOk

`func (o *ControlsData) GetControlCustomSelect20OptionIdOk() (*int32, bool)`

GetControlCustomSelect20OptionIdOk returns a tuple with the ControlCustomSelect20OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect20OptionId

`func (o *ControlsData) SetControlCustomSelect20OptionId(v int32)`

SetControlCustomSelect20OptionId sets ControlCustomSelect20OptionId field to given value.

### HasControlCustomSelect20OptionId

`func (o *ControlsData) HasControlCustomSelect20OptionId() bool`

HasControlCustomSelect20OptionId returns a boolean if a field has been set.

### GetControlCustomSelect21OptionId

`func (o *ControlsData) GetControlCustomSelect21OptionId() int32`

GetControlCustomSelect21OptionId returns the ControlCustomSelect21OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect21OptionIdOk

`func (o *ControlsData) GetControlCustomSelect21OptionIdOk() (*int32, bool)`

GetControlCustomSelect21OptionIdOk returns a tuple with the ControlCustomSelect21OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect21OptionId

`func (o *ControlsData) SetControlCustomSelect21OptionId(v int32)`

SetControlCustomSelect21OptionId sets ControlCustomSelect21OptionId field to given value.

### HasControlCustomSelect21OptionId

`func (o *ControlsData) HasControlCustomSelect21OptionId() bool`

HasControlCustomSelect21OptionId returns a boolean if a field has been set.

### GetControlCustomSelect22OptionId

`func (o *ControlsData) GetControlCustomSelect22OptionId() int32`

GetControlCustomSelect22OptionId returns the ControlCustomSelect22OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect22OptionIdOk

`func (o *ControlsData) GetControlCustomSelect22OptionIdOk() (*int32, bool)`

GetControlCustomSelect22OptionIdOk returns a tuple with the ControlCustomSelect22OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect22OptionId

`func (o *ControlsData) SetControlCustomSelect22OptionId(v int32)`

SetControlCustomSelect22OptionId sets ControlCustomSelect22OptionId field to given value.

### HasControlCustomSelect22OptionId

`func (o *ControlsData) HasControlCustomSelect22OptionId() bool`

HasControlCustomSelect22OptionId returns a boolean if a field has been set.

### GetControlCustomSelect23OptionId

`func (o *ControlsData) GetControlCustomSelect23OptionId() int32`

GetControlCustomSelect23OptionId returns the ControlCustomSelect23OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect23OptionIdOk

`func (o *ControlsData) GetControlCustomSelect23OptionIdOk() (*int32, bool)`

GetControlCustomSelect23OptionIdOk returns a tuple with the ControlCustomSelect23OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect23OptionId

`func (o *ControlsData) SetControlCustomSelect23OptionId(v int32)`

SetControlCustomSelect23OptionId sets ControlCustomSelect23OptionId field to given value.

### HasControlCustomSelect23OptionId

`func (o *ControlsData) HasControlCustomSelect23OptionId() bool`

HasControlCustomSelect23OptionId returns a boolean if a field has been set.

### GetControlCustomSelect24OptionId

`func (o *ControlsData) GetControlCustomSelect24OptionId() int32`

GetControlCustomSelect24OptionId returns the ControlCustomSelect24OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect24OptionIdOk

`func (o *ControlsData) GetControlCustomSelect24OptionIdOk() (*int32, bool)`

GetControlCustomSelect24OptionIdOk returns a tuple with the ControlCustomSelect24OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect24OptionId

`func (o *ControlsData) SetControlCustomSelect24OptionId(v int32)`

SetControlCustomSelect24OptionId sets ControlCustomSelect24OptionId field to given value.

### HasControlCustomSelect24OptionId

`func (o *ControlsData) HasControlCustomSelect24OptionId() bool`

HasControlCustomSelect24OptionId returns a boolean if a field has been set.

### GetControlCustomSelect25OptionId

`func (o *ControlsData) GetControlCustomSelect25OptionId() int32`

GetControlCustomSelect25OptionId returns the ControlCustomSelect25OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect25OptionIdOk

`func (o *ControlsData) GetControlCustomSelect25OptionIdOk() (*int32, bool)`

GetControlCustomSelect25OptionIdOk returns a tuple with the ControlCustomSelect25OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect25OptionId

`func (o *ControlsData) SetControlCustomSelect25OptionId(v int32)`

SetControlCustomSelect25OptionId sets ControlCustomSelect25OptionId field to given value.

### HasControlCustomSelect25OptionId

`func (o *ControlsData) HasControlCustomSelect25OptionId() bool`

HasControlCustomSelect25OptionId returns a boolean if a field has been set.

### GetControlCustomSelect26OptionId

`func (o *ControlsData) GetControlCustomSelect26OptionId() int32`

GetControlCustomSelect26OptionId returns the ControlCustomSelect26OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect26OptionIdOk

`func (o *ControlsData) GetControlCustomSelect26OptionIdOk() (*int32, bool)`

GetControlCustomSelect26OptionIdOk returns a tuple with the ControlCustomSelect26OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect26OptionId

`func (o *ControlsData) SetControlCustomSelect26OptionId(v int32)`

SetControlCustomSelect26OptionId sets ControlCustomSelect26OptionId field to given value.

### HasControlCustomSelect26OptionId

`func (o *ControlsData) HasControlCustomSelect26OptionId() bool`

HasControlCustomSelect26OptionId returns a boolean if a field has been set.

### GetControlCustomSelect27OptionId

`func (o *ControlsData) GetControlCustomSelect27OptionId() int32`

GetControlCustomSelect27OptionId returns the ControlCustomSelect27OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect27OptionIdOk

`func (o *ControlsData) GetControlCustomSelect27OptionIdOk() (*int32, bool)`

GetControlCustomSelect27OptionIdOk returns a tuple with the ControlCustomSelect27OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect27OptionId

`func (o *ControlsData) SetControlCustomSelect27OptionId(v int32)`

SetControlCustomSelect27OptionId sets ControlCustomSelect27OptionId field to given value.

### HasControlCustomSelect27OptionId

`func (o *ControlsData) HasControlCustomSelect27OptionId() bool`

HasControlCustomSelect27OptionId returns a boolean if a field has been set.

### GetControlCustomSelect28OptionId

`func (o *ControlsData) GetControlCustomSelect28OptionId() int32`

GetControlCustomSelect28OptionId returns the ControlCustomSelect28OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect28OptionIdOk

`func (o *ControlsData) GetControlCustomSelect28OptionIdOk() (*int32, bool)`

GetControlCustomSelect28OptionIdOk returns a tuple with the ControlCustomSelect28OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect28OptionId

`func (o *ControlsData) SetControlCustomSelect28OptionId(v int32)`

SetControlCustomSelect28OptionId sets ControlCustomSelect28OptionId field to given value.

### HasControlCustomSelect28OptionId

`func (o *ControlsData) HasControlCustomSelect28OptionId() bool`

HasControlCustomSelect28OptionId returns a boolean if a field has been set.

### GetControlCustomSelect29OptionId

`func (o *ControlsData) GetControlCustomSelect29OptionId() int32`

GetControlCustomSelect29OptionId returns the ControlCustomSelect29OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect29OptionIdOk

`func (o *ControlsData) GetControlCustomSelect29OptionIdOk() (*int32, bool)`

GetControlCustomSelect29OptionIdOk returns a tuple with the ControlCustomSelect29OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect29OptionId

`func (o *ControlsData) SetControlCustomSelect29OptionId(v int32)`

SetControlCustomSelect29OptionId sets ControlCustomSelect29OptionId field to given value.

### HasControlCustomSelect29OptionId

`func (o *ControlsData) HasControlCustomSelect29OptionId() bool`

HasControlCustomSelect29OptionId returns a boolean if a field has been set.

### GetControlCustomSelect30OptionId

`func (o *ControlsData) GetControlCustomSelect30OptionId() int32`

GetControlCustomSelect30OptionId returns the ControlCustomSelect30OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect30OptionIdOk

`func (o *ControlsData) GetControlCustomSelect30OptionIdOk() (*int32, bool)`

GetControlCustomSelect30OptionIdOk returns a tuple with the ControlCustomSelect30OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect30OptionId

`func (o *ControlsData) SetControlCustomSelect30OptionId(v int32)`

SetControlCustomSelect30OptionId sets ControlCustomSelect30OptionId field to given value.

### HasControlCustomSelect30OptionId

`func (o *ControlsData) HasControlCustomSelect30OptionId() bool`

HasControlCustomSelect30OptionId returns a boolean if a field has been set.

### GetControlCustomSelect31OptionId

`func (o *ControlsData) GetControlCustomSelect31OptionId() int32`

GetControlCustomSelect31OptionId returns the ControlCustomSelect31OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect31OptionIdOk

`func (o *ControlsData) GetControlCustomSelect31OptionIdOk() (*int32, bool)`

GetControlCustomSelect31OptionIdOk returns a tuple with the ControlCustomSelect31OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect31OptionId

`func (o *ControlsData) SetControlCustomSelect31OptionId(v int32)`

SetControlCustomSelect31OptionId sets ControlCustomSelect31OptionId field to given value.

### HasControlCustomSelect31OptionId

`func (o *ControlsData) HasControlCustomSelect31OptionId() bool`

HasControlCustomSelect31OptionId returns a boolean if a field has been set.

### GetControlCustomSelect32OptionId

`func (o *ControlsData) GetControlCustomSelect32OptionId() int32`

GetControlCustomSelect32OptionId returns the ControlCustomSelect32OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect32OptionIdOk

`func (o *ControlsData) GetControlCustomSelect32OptionIdOk() (*int32, bool)`

GetControlCustomSelect32OptionIdOk returns a tuple with the ControlCustomSelect32OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect32OptionId

`func (o *ControlsData) SetControlCustomSelect32OptionId(v int32)`

SetControlCustomSelect32OptionId sets ControlCustomSelect32OptionId field to given value.

### HasControlCustomSelect32OptionId

`func (o *ControlsData) HasControlCustomSelect32OptionId() bool`

HasControlCustomSelect32OptionId returns a boolean if a field has been set.

### GetControlCustomSelect33OptionId

`func (o *ControlsData) GetControlCustomSelect33OptionId() int32`

GetControlCustomSelect33OptionId returns the ControlCustomSelect33OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect33OptionIdOk

`func (o *ControlsData) GetControlCustomSelect33OptionIdOk() (*int32, bool)`

GetControlCustomSelect33OptionIdOk returns a tuple with the ControlCustomSelect33OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect33OptionId

`func (o *ControlsData) SetControlCustomSelect33OptionId(v int32)`

SetControlCustomSelect33OptionId sets ControlCustomSelect33OptionId field to given value.

### HasControlCustomSelect33OptionId

`func (o *ControlsData) HasControlCustomSelect33OptionId() bool`

HasControlCustomSelect33OptionId returns a boolean if a field has been set.

### GetControlCustomSelect34OptionId

`func (o *ControlsData) GetControlCustomSelect34OptionId() int32`

GetControlCustomSelect34OptionId returns the ControlCustomSelect34OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect34OptionIdOk

`func (o *ControlsData) GetControlCustomSelect34OptionIdOk() (*int32, bool)`

GetControlCustomSelect34OptionIdOk returns a tuple with the ControlCustomSelect34OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect34OptionId

`func (o *ControlsData) SetControlCustomSelect34OptionId(v int32)`

SetControlCustomSelect34OptionId sets ControlCustomSelect34OptionId field to given value.

### HasControlCustomSelect34OptionId

`func (o *ControlsData) HasControlCustomSelect34OptionId() bool`

HasControlCustomSelect34OptionId returns a boolean if a field has been set.

### GetControlCustomSelect35OptionId

`func (o *ControlsData) GetControlCustomSelect35OptionId() int32`

GetControlCustomSelect35OptionId returns the ControlCustomSelect35OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect35OptionIdOk

`func (o *ControlsData) GetControlCustomSelect35OptionIdOk() (*int32, bool)`

GetControlCustomSelect35OptionIdOk returns a tuple with the ControlCustomSelect35OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect35OptionId

`func (o *ControlsData) SetControlCustomSelect35OptionId(v int32)`

SetControlCustomSelect35OptionId sets ControlCustomSelect35OptionId field to given value.

### HasControlCustomSelect35OptionId

`func (o *ControlsData) HasControlCustomSelect35OptionId() bool`

HasControlCustomSelect35OptionId returns a boolean if a field has been set.

### GetControlCustomSelect36OptionId

`func (o *ControlsData) GetControlCustomSelect36OptionId() int32`

GetControlCustomSelect36OptionId returns the ControlCustomSelect36OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect36OptionIdOk

`func (o *ControlsData) GetControlCustomSelect36OptionIdOk() (*int32, bool)`

GetControlCustomSelect36OptionIdOk returns a tuple with the ControlCustomSelect36OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect36OptionId

`func (o *ControlsData) SetControlCustomSelect36OptionId(v int32)`

SetControlCustomSelect36OptionId sets ControlCustomSelect36OptionId field to given value.

### HasControlCustomSelect36OptionId

`func (o *ControlsData) HasControlCustomSelect36OptionId() bool`

HasControlCustomSelect36OptionId returns a boolean if a field has been set.

### GetAtRiskStatus

`func (o *ControlsData) GetAtRiskStatus() string`

GetAtRiskStatus returns the AtRiskStatus field if non-nil, zero value otherwise.

### GetAtRiskStatusOk

`func (o *ControlsData) GetAtRiskStatusOk() (*string, bool)`

GetAtRiskStatusOk returns a tuple with the AtRiskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtRiskStatus

`func (o *ControlsData) SetAtRiskStatus(v string)`

SetAtRiskStatus sets AtRiskStatus field to given value.


### GetLatestResultId

`func (o *ControlsData) GetLatestResultId() int32`

GetLatestResultId returns the LatestResultId field if non-nil, zero value otherwise.

### GetLatestResultIdOk

`func (o *ControlsData) GetLatestResultIdOk() (*int32, bool)`

GetLatestResultIdOk returns a tuple with the LatestResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestResultId

`func (o *ControlsData) SetLatestResultId(v int32)`

SetLatestResultId sets LatestResultId field to given value.

### HasLatestResultId

`func (o *ControlsData) HasLatestResultId() bool`

HasLatestResultId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


