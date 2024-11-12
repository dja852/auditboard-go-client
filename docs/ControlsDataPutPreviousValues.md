# ControlsDataPutPreviousValues

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
**ReferenceMeta** | Pointer to **interface{}** |  | [optional] 
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
**Scopes** | Pointer to **interface{}** |  | [optional] 
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
**AtRiskStatus** | Pointer to **string** |  | [optional] [default to "Not At Risk"]
**LatestResultId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;controls_datum_results.id&#x60;.&lt;fk table&#x3D;&#39;controls_datum_results&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 

## Methods

### NewControlsDataPutPreviousValues

`func NewControlsDataPutPreviousValues() *ControlsDataPutPreviousValues`

NewControlsDataPutPreviousValues instantiates a new ControlsDataPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsDataPutPreviousValuesWithDefaults

`func NewControlsDataPutPreviousValuesWithDefaults() *ControlsDataPutPreviousValues`

NewControlsDataPutPreviousValuesWithDefaults instantiates a new ControlsDataPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ControlsDataPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControlsDataPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControlsDataPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ControlsDataPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ControlsDataPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ControlsDataPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ControlsDataPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ControlsDataPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ControlsDataPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ControlsDataPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ControlsDataPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ControlsDataPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ControlsDataPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ControlsDataPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ControlsDataPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ControlsDataPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetControlId

`func (o *ControlsDataPutPreviousValues) GetControlId() int32`

GetControlId returns the ControlId field if non-nil, zero value otherwise.

### GetControlIdOk

`func (o *ControlsDataPutPreviousValues) GetControlIdOk() (*int32, bool)`

GetControlIdOk returns a tuple with the ControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlId

`func (o *ControlsDataPutPreviousValues) SetControlId(v int32)`

SetControlId sets ControlId field to given value.

### HasControlId

`func (o *ControlsDataPutPreviousValues) HasControlId() bool`

HasControlId returns a boolean if a field has been set.

### GetDescription

`func (o *ControlsDataPutPreviousValues) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ControlsDataPutPreviousValues) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ControlsDataPutPreviousValues) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ControlsDataPutPreviousValues) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubprocessesDatumId

`func (o *ControlsDataPutPreviousValues) GetSubprocessesDatumId() int32`

GetSubprocessesDatumId returns the SubprocessesDatumId field if non-nil, zero value otherwise.

### GetSubprocessesDatumIdOk

`func (o *ControlsDataPutPreviousValues) GetSubprocessesDatumIdOk() (*int32, bool)`

GetSubprocessesDatumIdOk returns a tuple with the SubprocessesDatumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubprocessesDatumId

`func (o *ControlsDataPutPreviousValues) SetSubprocessesDatumId(v int32)`

SetSubprocessesDatumId sets SubprocessesDatumId field to given value.

### HasSubprocessesDatumId

`func (o *ControlsDataPutPreviousValues) HasSubprocessesDatumId() bool`

HasSubprocessesDatumId returns a boolean if a field has been set.

### GetRiskLevelId

`func (o *ControlsDataPutPreviousValues) GetRiskLevelId() int32`

GetRiskLevelId returns the RiskLevelId field if non-nil, zero value otherwise.

### GetRiskLevelIdOk

`func (o *ControlsDataPutPreviousValues) GetRiskLevelIdOk() (*int32, bool)`

GetRiskLevelIdOk returns a tuple with the RiskLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevelId

`func (o *ControlsDataPutPreviousValues) SetRiskLevelId(v int32)`

SetRiskLevelId sets RiskLevelId field to given value.

### HasRiskLevelId

`func (o *ControlsDataPutPreviousValues) HasRiskLevelId() bool`

HasRiskLevelId returns a boolean if a field has been set.

### GetSupervisorUserId

`func (o *ControlsDataPutPreviousValues) GetSupervisorUserId() int32`

GetSupervisorUserId returns the SupervisorUserId field if non-nil, zero value otherwise.

### GetSupervisorUserIdOk

`func (o *ControlsDataPutPreviousValues) GetSupervisorUserIdOk() (*int32, bool)`

GetSupervisorUserIdOk returns a tuple with the SupervisorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisorUserId

`func (o *ControlsDataPutPreviousValues) SetSupervisorUserId(v int32)`

SetSupervisorUserId sets SupervisorUserId field to given value.

### HasSupervisorUserId

`func (o *ControlsDataPutPreviousValues) HasSupervisorUserId() bool`

HasSupervisorUserId returns a boolean if a field has been set.

### GetOwnerUserId

`func (o *ControlsDataPutPreviousValues) GetOwnerUserId() int32`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *ControlsDataPutPreviousValues) GetOwnerUserIdOk() (*int32, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *ControlsDataPutPreviousValues) SetOwnerUserId(v int32)`

SetOwnerUserId sets OwnerUserId field to given value.

### HasOwnerUserId

`func (o *ControlsDataPutPreviousValues) HasOwnerUserId() bool`

HasOwnerUserId returns a boolean if a field has been set.

### GetCosoElementId

`func (o *ControlsDataPutPreviousValues) GetCosoElementId() int32`

GetCosoElementId returns the CosoElementId field if non-nil, zero value otherwise.

### GetCosoElementIdOk

`func (o *ControlsDataPutPreviousValues) GetCosoElementIdOk() (*int32, bool)`

GetCosoElementIdOk returns a tuple with the CosoElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosoElementId

`func (o *ControlsDataPutPreviousValues) SetCosoElementId(v int32)`

SetCosoElementId sets CosoElementId field to given value.

### HasCosoElementId

`func (o *ControlsDataPutPreviousValues) HasCosoElementId() bool`

HasCosoElementId returns a boolean if a field has been set.

### GetControlTypeId

`func (o *ControlsDataPutPreviousValues) GetControlTypeId() int32`

GetControlTypeId returns the ControlTypeId field if non-nil, zero value otherwise.

### GetControlTypeIdOk

`func (o *ControlsDataPutPreviousValues) GetControlTypeIdOk() (*int32, bool)`

GetControlTypeIdOk returns a tuple with the ControlTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlTypeId

`func (o *ControlsDataPutPreviousValues) SetControlTypeId(v int32)`

SetControlTypeId sets ControlTypeId field to given value.

### HasControlTypeId

`func (o *ControlsDataPutPreviousValues) HasControlTypeId() bool`

HasControlTypeId returns a boolean if a field has been set.

### GetControlNatureOptionId

`func (o *ControlsDataPutPreviousValues) GetControlNatureOptionId() int32`

GetControlNatureOptionId returns the ControlNatureOptionId field if non-nil, zero value otherwise.

### GetControlNatureOptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlNatureOptionIdOk() (*int32, bool)`

GetControlNatureOptionIdOk returns a tuple with the ControlNatureOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlNatureOptionId

`func (o *ControlsDataPutPreviousValues) SetControlNatureOptionId(v int32)`

SetControlNatureOptionId sets ControlNatureOptionId field to given value.

### HasControlNatureOptionId

`func (o *ControlsDataPutPreviousValues) HasControlNatureOptionId() bool`

HasControlNatureOptionId returns a boolean if a field has been set.

### GetSignificanceOptionId

`func (o *ControlsDataPutPreviousValues) GetSignificanceOptionId() int32`

GetSignificanceOptionId returns the SignificanceOptionId field if non-nil, zero value otherwise.

### GetSignificanceOptionIdOk

`func (o *ControlsDataPutPreviousValues) GetSignificanceOptionIdOk() (*int32, bool)`

GetSignificanceOptionIdOk returns a tuple with the SignificanceOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificanceOptionId

`func (o *ControlsDataPutPreviousValues) SetSignificanceOptionId(v int32)`

SetSignificanceOptionId sets SignificanceOptionId field to given value.

### HasSignificanceOptionId

`func (o *ControlsDataPutPreviousValues) HasSignificanceOptionId() bool`

HasSignificanceOptionId returns a boolean if a field has been set.

### GetTestProcedures

`func (o *ControlsDataPutPreviousValues) GetTestProcedures() string`

GetTestProcedures returns the TestProcedures field if non-nil, zero value otherwise.

### GetTestProceduresOk

`func (o *ControlsDataPutPreviousValues) GetTestProceduresOk() (*string, bool)`

GetTestProceduresOk returns a tuple with the TestProcedures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestProcedures

`func (o *ControlsDataPutPreviousValues) SetTestProcedures(v string)`

SetTestProcedures sets TestProcedures field to given value.

### HasTestProcedures

`func (o *ControlsDataPutPreviousValues) HasTestProcedures() bool`

HasTestProcedures returns a boolean if a field has been set.

### GetPbcRequest

`func (o *ControlsDataPutPreviousValues) GetPbcRequest() string`

GetPbcRequest returns the PbcRequest field if non-nil, zero value otherwise.

### GetPbcRequestOk

`func (o *ControlsDataPutPreviousValues) GetPbcRequestOk() (*string, bool)`

GetPbcRequestOk returns a tuple with the PbcRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbcRequest

`func (o *ControlsDataPutPreviousValues) SetPbcRequest(v string)`

SetPbcRequest sets PbcRequest field to given value.

### HasPbcRequest

`func (o *ControlsDataPutPreviousValues) HasPbcRequest() bool`

HasPbcRequest returns a boolean if a field has been set.

### GetSourceData

`func (o *ControlsDataPutPreviousValues) GetSourceData() string`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *ControlsDataPutPreviousValues) GetSourceDataOk() (*string, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *ControlsDataPutPreviousValues) SetSourceData(v string)`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *ControlsDataPutPreviousValues) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetNotes

`func (o *ControlsDataPutPreviousValues) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ControlsDataPutPreviousValues) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ControlsDataPutPreviousValues) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ControlsDataPutPreviousValues) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetControlFrequencyId

`func (o *ControlsDataPutPreviousValues) GetControlFrequencyId() int32`

GetControlFrequencyId returns the ControlFrequencyId field if non-nil, zero value otherwise.

### GetControlFrequencyIdOk

`func (o *ControlsDataPutPreviousValues) GetControlFrequencyIdOk() (*int32, bool)`

GetControlFrequencyIdOk returns a tuple with the ControlFrequencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlFrequencyId

`func (o *ControlsDataPutPreviousValues) SetControlFrequencyId(v int32)`

SetControlFrequencyId sets ControlFrequencyId field to given value.

### HasControlFrequencyId

`func (o *ControlsDataPutPreviousValues) HasControlFrequencyId() bool`

HasControlFrequencyId returns a boolean if a field has been set.

### GetTestTimingId

`func (o *ControlsDataPutPreviousValues) GetTestTimingId() int32`

GetTestTimingId returns the TestTimingId field if non-nil, zero value otherwise.

### GetTestTimingIdOk

`func (o *ControlsDataPutPreviousValues) GetTestTimingIdOk() (*int32, bool)`

GetTestTimingIdOk returns a tuple with the TestTimingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestTimingId

`func (o *ControlsDataPutPreviousValues) SetTestTimingId(v int32)`

SetTestTimingId sets TestTimingId field to given value.

### HasTestTimingId

`func (o *ControlsDataPutPreviousValues) HasTestTimingId() bool`

HasTestTimingId returns a boolean if a field has been set.

### GetTestTypeId

`func (o *ControlsDataPutPreviousValues) GetTestTypeId() int32`

GetTestTypeId returns the TestTypeId field if non-nil, zero value otherwise.

### GetTestTypeIdOk

`func (o *ControlsDataPutPreviousValues) GetTestTypeIdOk() (*int32, bool)`

GetTestTypeIdOk returns a tuple with the TestTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestTypeId

`func (o *ControlsDataPutPreviousValues) SetTestTypeId(v int32)`

SetTestTypeId sets TestTypeId field to given value.

### HasTestTypeId

`func (o *ControlsDataPutPreviousValues) HasTestTypeId() bool`

HasTestTypeId returns a boolean if a field has been set.

### GetIssueTitle

`func (o *ControlsDataPutPreviousValues) GetIssueTitle() string`

GetIssueTitle returns the IssueTitle field if non-nil, zero value otherwise.

### GetIssueTitleOk

`func (o *ControlsDataPutPreviousValues) GetIssueTitleOk() (*string, bool)`

GetIssueTitleOk returns a tuple with the IssueTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTitle

`func (o *ControlsDataPutPreviousValues) SetIssueTitle(v string)`

SetIssueTitle sets IssueTitle field to given value.

### HasIssueTitle

`func (o *ControlsDataPutPreviousValues) HasIssueTitle() bool`

HasIssueTitle returns a boolean if a field has been set.

### GetIssueDescription

`func (o *ControlsDataPutPreviousValues) GetIssueDescription() string`

GetIssueDescription returns the IssueDescription field if non-nil, zero value otherwise.

### GetIssueDescriptionOk

`func (o *ControlsDataPutPreviousValues) GetIssueDescriptionOk() (*string, bool)`

GetIssueDescriptionOk returns a tuple with the IssueDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDescription

`func (o *ControlsDataPutPreviousValues) SetIssueDescription(v string)`

SetIssueDescription sets IssueDescription field to given value.

### HasIssueDescription

`func (o *ControlsDataPutPreviousValues) HasIssueDescription() bool`

HasIssueDescription returns a boolean if a field has been set.

### GetRemediationAction

`func (o *ControlsDataPutPreviousValues) GetRemediationAction() string`

GetRemediationAction returns the RemediationAction field if non-nil, zero value otherwise.

### GetRemediationActionOk

`func (o *ControlsDataPutPreviousValues) GetRemediationActionOk() (*string, bool)`

GetRemediationActionOk returns a tuple with the RemediationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationAction

`func (o *ControlsDataPutPreviousValues) SetRemediationAction(v string)`

SetRemediationAction sets RemediationAction field to given value.

### HasRemediationAction

`func (o *ControlsDataPutPreviousValues) HasRemediationAction() bool`

HasRemediationAction returns a boolean if a field has been set.

### GetRemediationOwnerId

`func (o *ControlsDataPutPreviousValues) GetRemediationOwnerId() int32`

GetRemediationOwnerId returns the RemediationOwnerId field if non-nil, zero value otherwise.

### GetRemediationOwnerIdOk

`func (o *ControlsDataPutPreviousValues) GetRemediationOwnerIdOk() (*int32, bool)`

GetRemediationOwnerIdOk returns a tuple with the RemediationOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationOwnerId

`func (o *ControlsDataPutPreviousValues) SetRemediationOwnerId(v int32)`

SetRemediationOwnerId sets RemediationOwnerId field to given value.

### HasRemediationOwnerId

`func (o *ControlsDataPutPreviousValues) HasRemediationOwnerId() bool`

HasRemediationOwnerId returns a boolean if a field has been set.

### GetRemediationDate

`func (o *ControlsDataPutPreviousValues) GetRemediationDate() string`

GetRemediationDate returns the RemediationDate field if non-nil, zero value otherwise.

### GetRemediationDateOk

`func (o *ControlsDataPutPreviousValues) GetRemediationDateOk() (*string, bool)`

GetRemediationDateOk returns a tuple with the RemediationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationDate

`func (o *ControlsDataPutPreviousValues) SetRemediationDate(v string)`

SetRemediationDate sets RemediationDate field to given value.

### HasRemediationDate

`func (o *ControlsDataPutPreviousValues) HasRemediationDate() bool`

HasRemediationDate returns a boolean if a field has been set.

### GetIssueIdentId

`func (o *ControlsDataPutPreviousValues) GetIssueIdentId() int32`

GetIssueIdentId returns the IssueIdentId field if non-nil, zero value otherwise.

### GetIssueIdentIdOk

`func (o *ControlsDataPutPreviousValues) GetIssueIdentIdOk() (*int32, bool)`

GetIssueIdentIdOk returns a tuple with the IssueIdentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueIdentId

`func (o *ControlsDataPutPreviousValues) SetIssueIdentId(v int32)`

SetIssueIdentId sets IssueIdentId field to given value.

### HasIssueIdentId

`func (o *ControlsDataPutPreviousValues) HasIssueIdentId() bool`

HasIssueIdentId returns a boolean if a field has been set.

### GetIssueAccount

`func (o *ControlsDataPutPreviousValues) GetIssueAccount() string`

GetIssueAccount returns the IssueAccount field if non-nil, zero value otherwise.

### GetIssueAccountOk

`func (o *ControlsDataPutPreviousValues) GetIssueAccountOk() (*string, bool)`

GetIssueAccountOk returns a tuple with the IssueAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueAccount

`func (o *ControlsDataPutPreviousValues) SetIssueAccount(v string)`

SetIssueAccount sets IssueAccount field to given value.

### HasIssueAccount

`func (o *ControlsDataPutPreviousValues) HasIssueAccount() bool`

HasIssueAccount returns a boolean if a field has been set.

### GetPotentialRisk

`func (o *ControlsDataPutPreviousValues) GetPotentialRisk() string`

GetPotentialRisk returns the PotentialRisk field if non-nil, zero value otherwise.

### GetPotentialRiskOk

`func (o *ControlsDataPutPreviousValues) GetPotentialRiskOk() (*string, bool)`

GetPotentialRiskOk returns a tuple with the PotentialRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPotentialRisk

`func (o *ControlsDataPutPreviousValues) SetPotentialRisk(v string)`

SetPotentialRisk sets PotentialRisk field to given value.

### HasPotentialRisk

`func (o *ControlsDataPutPreviousValues) HasPotentialRisk() bool`

HasPotentialRisk returns a boolean if a field has been set.

### GetComplementaryControls

`func (o *ControlsDataPutPreviousValues) GetComplementaryControls() string`

GetComplementaryControls returns the ComplementaryControls field if non-nil, zero value otherwise.

### GetComplementaryControlsOk

`func (o *ControlsDataPutPreviousValues) GetComplementaryControlsOk() (*string, bool)`

GetComplementaryControlsOk returns a tuple with the ComplementaryControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplementaryControls

`func (o *ControlsDataPutPreviousValues) SetComplementaryControls(v string)`

SetComplementaryControls sets ComplementaryControls field to given value.

### HasComplementaryControls

`func (o *ControlsDataPutPreviousValues) HasComplementaryControls() bool`

HasComplementaryControls returns a boolean if a field has been set.

### GetQualitativeFactors

`func (o *ControlsDataPutPreviousValues) GetQualitativeFactors() string`

GetQualitativeFactors returns the QualitativeFactors field if non-nil, zero value otherwise.

### GetQualitativeFactorsOk

`func (o *ControlsDataPutPreviousValues) GetQualitativeFactorsOk() (*string, bool)`

GetQualitativeFactorsOk returns a tuple with the QualitativeFactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualitativeFactors

`func (o *ControlsDataPutPreviousValues) SetQualitativeFactors(v string)`

SetQualitativeFactors sets QualitativeFactors field to given value.

### HasQualitativeFactors

`func (o *ControlsDataPutPreviousValues) HasQualitativeFactors() bool`

HasQualitativeFactors returns a boolean if a field has been set.

### GetIssueId

`func (o *ControlsDataPutPreviousValues) GetIssueId() string`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *ControlsDataPutPreviousValues) GetIssueIdOk() (*string, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *ControlsDataPutPreviousValues) SetIssueId(v string)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *ControlsDataPutPreviousValues) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.

### GetGrossExposure

`func (o *ControlsDataPutPreviousValues) GetGrossExposure() string`

GetGrossExposure returns the GrossExposure field if non-nil, zero value otherwise.

### GetGrossExposureOk

`func (o *ControlsDataPutPreviousValues) GetGrossExposureOk() (*string, bool)`

GetGrossExposureOk returns a tuple with the GrossExposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossExposure

`func (o *ControlsDataPutPreviousValues) SetGrossExposure(v string)`

SetGrossExposure sets GrossExposure field to given value.

### HasGrossExposure

`func (o *ControlsDataPutPreviousValues) HasGrossExposure() bool`

HasGrossExposure returns a boolean if a field has been set.

### GetAdjustedExposure

`func (o *ControlsDataPutPreviousValues) GetAdjustedExposure() string`

GetAdjustedExposure returns the AdjustedExposure field if non-nil, zero value otherwise.

### GetAdjustedExposureOk

`func (o *ControlsDataPutPreviousValues) GetAdjustedExposureOk() (*string, bool)`

GetAdjustedExposureOk returns a tuple with the AdjustedExposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedExposure

`func (o *ControlsDataPutPreviousValues) SetAdjustedExposure(v string)`

SetAdjustedExposure sets AdjustedExposure field to given value.

### HasAdjustedExposure

`func (o *ControlsDataPutPreviousValues) HasAdjustedExposure() bool`

HasAdjustedExposure returns a boolean if a field has been set.

### GetLikelihoodId

`func (o *ControlsDataPutPreviousValues) GetLikelihoodId() int32`

GetLikelihoodId returns the LikelihoodId field if non-nil, zero value otherwise.

### GetLikelihoodIdOk

`func (o *ControlsDataPutPreviousValues) GetLikelihoodIdOk() (*int32, bool)`

GetLikelihoodIdOk returns a tuple with the LikelihoodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikelihoodId

`func (o *ControlsDataPutPreviousValues) SetLikelihoodId(v int32)`

SetLikelihoodId sets LikelihoodId field to given value.

### HasLikelihoodId

`func (o *ControlsDataPutPreviousValues) HasLikelihoodId() bool`

HasLikelihoodId returns a boolean if a field has been set.

### GetMagnitudeId

`func (o *ControlsDataPutPreviousValues) GetMagnitudeId() int32`

GetMagnitudeId returns the MagnitudeId field if non-nil, zero value otherwise.

### GetMagnitudeIdOk

`func (o *ControlsDataPutPreviousValues) GetMagnitudeIdOk() (*int32, bool)`

GetMagnitudeIdOk returns a tuple with the MagnitudeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagnitudeId

`func (o *ControlsDataPutPreviousValues) SetMagnitudeId(v int32)`

SetMagnitudeId sets MagnitudeId field to given value.

### HasMagnitudeId

`func (o *ControlsDataPutPreviousValues) HasMagnitudeId() bool`

HasMagnitudeId returns a boolean if a field has been set.

### GetDeficiencyLevelId

`func (o *ControlsDataPutPreviousValues) GetDeficiencyLevelId() int32`

GetDeficiencyLevelId returns the DeficiencyLevelId field if non-nil, zero value otherwise.

### GetDeficiencyLevelIdOk

`func (o *ControlsDataPutPreviousValues) GetDeficiencyLevelIdOk() (*int32, bool)`

GetDeficiencyLevelIdOk returns a tuple with the DeficiencyLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeficiencyLevelId

`func (o *ControlsDataPutPreviousValues) SetDeficiencyLevelId(v int32)`

SetDeficiencyLevelId sets DeficiencyLevelId field to given value.

### HasDeficiencyLevelId

`func (o *ControlsDataPutPreviousValues) HasDeficiencyLevelId() bool`

HasDeficiencyLevelId returns a boolean if a field has been set.

### GetAggregationReferenceId

`func (o *ControlsDataPutPreviousValues) GetAggregationReferenceId() int32`

GetAggregationReferenceId returns the AggregationReferenceId field if non-nil, zero value otherwise.

### GetAggregationReferenceIdOk

`func (o *ControlsDataPutPreviousValues) GetAggregationReferenceIdOk() (*int32, bool)`

GetAggregationReferenceIdOk returns a tuple with the AggregationReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationReferenceId

`func (o *ControlsDataPutPreviousValues) SetAggregationReferenceId(v int32)`

SetAggregationReferenceId sets AggregationReferenceId field to given value.

### HasAggregationReferenceId

`func (o *ControlsDataPutPreviousValues) HasAggregationReferenceId() bool`

HasAggregationReferenceId returns a boolean if a field has been set.

### GetIssueComments

`func (o *ControlsDataPutPreviousValues) GetIssueComments() string`

GetIssueComments returns the IssueComments field if non-nil, zero value otherwise.

### GetIssueCommentsOk

`func (o *ControlsDataPutPreviousValues) GetIssueCommentsOk() (*string, bool)`

GetIssueCommentsOk returns a tuple with the IssueComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueComments

`func (o *ControlsDataPutPreviousValues) SetIssueComments(v string)`

SetIssueComments sets IssueComments field to given value.

### HasIssueComments

`func (o *ControlsDataPutPreviousValues) HasIssueComments() bool`

HasIssueComments returns a boolean if a field has been set.

### GetPbcRequestAdditional

`func (o *ControlsDataPutPreviousValues) GetPbcRequestAdditional() string`

GetPbcRequestAdditional returns the PbcRequestAdditional field if non-nil, zero value otherwise.

### GetPbcRequestAdditionalOk

`func (o *ControlsDataPutPreviousValues) GetPbcRequestAdditionalOk() (*string, bool)`

GetPbcRequestAdditionalOk returns a tuple with the PbcRequestAdditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbcRequestAdditional

`func (o *ControlsDataPutPreviousValues) SetPbcRequestAdditional(v string)`

SetPbcRequestAdditional sets PbcRequestAdditional field to given value.

### HasPbcRequestAdditional

`func (o *ControlsDataPutPreviousValues) HasPbcRequestAdditional() bool`

HasPbcRequestAdditional returns a boolean if a field has been set.

### GetSodControlId

`func (o *ControlsDataPutPreviousValues) GetSodControlId() int32`

GetSodControlId returns the SodControlId field if non-nil, zero value otherwise.

### GetSodControlIdOk

`func (o *ControlsDataPutPreviousValues) GetSodControlIdOk() (*int32, bool)`

GetSodControlIdOk returns a tuple with the SodControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSodControlId

`func (o *ControlsDataPutPreviousValues) SetSodControlId(v int32)`

SetSodControlId sets SodControlId field to given value.

### HasSodControlId

`func (o *ControlsDataPutPreviousValues) HasSodControlId() bool`

HasSodControlId returns a boolean if a field has been set.

### GetSamplingApproach

`func (o *ControlsDataPutPreviousValues) GetSamplingApproach() string`

GetSamplingApproach returns the SamplingApproach field if non-nil, zero value otherwise.

### GetSamplingApproachOk

`func (o *ControlsDataPutPreviousValues) GetSamplingApproachOk() (*string, bool)`

GetSamplingApproachOk returns a tuple with the SamplingApproach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingApproach

`func (o *ControlsDataPutPreviousValues) SetSamplingApproach(v string)`

SetSamplingApproach sets SamplingApproach field to given value.

### HasSamplingApproach

`func (o *ControlsDataPutPreviousValues) HasSamplingApproach() bool`

HasSamplingApproach returns a boolean if a field has been set.

### GetMrcControlId

`func (o *ControlsDataPutPreviousValues) GetMrcControlId() int32`

GetMrcControlId returns the MrcControlId field if non-nil, zero value otherwise.

### GetMrcControlIdOk

`func (o *ControlsDataPutPreviousValues) GetMrcControlIdOk() (*int32, bool)`

GetMrcControlIdOk returns a tuple with the MrcControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrcControlId

`func (o *ControlsDataPutPreviousValues) SetMrcControlId(v int32)`

SetMrcControlId sets MrcControlId field to given value.

### HasMrcControlId

`func (o *ControlsDataPutPreviousValues) HasMrcControlId() bool`

HasMrcControlId returns a boolean if a field has been set.

### GetFraudControlId

`func (o *ControlsDataPutPreviousValues) GetFraudControlId() int32`

GetFraudControlId returns the FraudControlId field if non-nil, zero value otherwise.

### GetFraudControlIdOk

`func (o *ControlsDataPutPreviousValues) GetFraudControlIdOk() (*int32, bool)`

GetFraudControlIdOk returns a tuple with the FraudControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudControlId

`func (o *ControlsDataPutPreviousValues) SetFraudControlId(v int32)`

SetFraudControlId sets FraudControlId field to given value.

### HasFraudControlId

`func (o *ControlsDataPutPreviousValues) HasFraudControlId() bool`

HasFraudControlId returns a boolean if a field has been set.

### GetEntityId

`func (o *ControlsDataPutPreviousValues) GetEntityId() int32`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ControlsDataPutPreviousValues) GetEntityIdOk() (*int32, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ControlsDataPutPreviousValues) SetEntityId(v int32)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *ControlsDataPutPreviousValues) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetPbcNotes

`func (o *ControlsDataPutPreviousValues) GetPbcNotes() string`

GetPbcNotes returns the PbcNotes field if non-nil, zero value otherwise.

### GetPbcNotesOk

`func (o *ControlsDataPutPreviousValues) GetPbcNotesOk() (*string, bool)`

GetPbcNotesOk returns a tuple with the PbcNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbcNotes

`func (o *ControlsDataPutPreviousValues) SetPbcNotes(v string)`

SetPbcNotes sets PbcNotes field to given value.

### HasPbcNotes

`func (o *ControlsDataPutPreviousValues) HasPbcNotes() bool`

HasPbcNotes returns a boolean if a field has been set.

### GetControlDesignOptionId

`func (o *ControlsDataPutPreviousValues) GetControlDesignOptionId() int32`

GetControlDesignOptionId returns the ControlDesignOptionId field if non-nil, zero value otherwise.

### GetControlDesignOptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlDesignOptionIdOk() (*int32, bool)`

GetControlDesignOptionIdOk returns a tuple with the ControlDesignOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlDesignOptionId

`func (o *ControlsDataPutPreviousValues) SetControlDesignOptionId(v int32)`

SetControlDesignOptionId sets ControlDesignOptionId field to given value.

### HasControlDesignOptionId

`func (o *ControlsDataPutPreviousValues) HasControlDesignOptionId() bool`

HasControlDesignOptionId returns a boolean if a field has been set.

### GetControlValidation

`func (o *ControlsDataPutPreviousValues) GetControlValidation() string`

GetControlValidation returns the ControlValidation field if non-nil, zero value otherwise.

### GetControlValidationOk

`func (o *ControlsDataPutPreviousValues) GetControlValidationOk() (*string, bool)`

GetControlValidationOk returns a tuple with the ControlValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlValidation

`func (o *ControlsDataPutPreviousValues) SetControlValidation(v string)`

SetControlValidation sets ControlValidation field to given value.

### HasControlValidation

`func (o *ControlsDataPutPreviousValues) HasControlValidation() bool`

HasControlValidation returns a boolean if a field has been set.

### GetAnnualSampleSize

`func (o *ControlsDataPutPreviousValues) GetAnnualSampleSize() string`

GetAnnualSampleSize returns the AnnualSampleSize field if non-nil, zero value otherwise.

### GetAnnualSampleSizeOk

`func (o *ControlsDataPutPreviousValues) GetAnnualSampleSizeOk() (*string, bool)`

GetAnnualSampleSizeOk returns a tuple with the AnnualSampleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualSampleSize

`func (o *ControlsDataPutPreviousValues) SetAnnualSampleSize(v string)`

SetAnnualSampleSize sets AnnualSampleSize field to given value.

### HasAnnualSampleSize

`func (o *ControlsDataPutPreviousValues) HasAnnualSampleSize() bool`

HasAnnualSampleSize returns a boolean if a field has been set.

### GetControlScopeId

`func (o *ControlsDataPutPreviousValues) GetControlScopeId() int32`

GetControlScopeId returns the ControlScopeId field if non-nil, zero value otherwise.

### GetControlScopeIdOk

`func (o *ControlsDataPutPreviousValues) GetControlScopeIdOk() (*int32, bool)`

GetControlScopeIdOk returns a tuple with the ControlScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlScopeId

`func (o *ControlsDataPutPreviousValues) SetControlScopeId(v int32)`

SetControlScopeId sets ControlScopeId field to given value.

### HasControlScopeId

`func (o *ControlsDataPutPreviousValues) HasControlScopeId() bool`

HasControlScopeId returns a boolean if a field has been set.

### GetControlClassificationId

`func (o *ControlsDataPutPreviousValues) GetControlClassificationId() int32`

GetControlClassificationId returns the ControlClassificationId field if non-nil, zero value otherwise.

### GetControlClassificationIdOk

`func (o *ControlsDataPutPreviousValues) GetControlClassificationIdOk() (*int32, bool)`

GetControlClassificationIdOk returns a tuple with the ControlClassificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlClassificationId

`func (o *ControlsDataPutPreviousValues) SetControlClassificationId(v int32)`

SetControlClassificationId sets ControlClassificationId field to given value.

### HasControlClassificationId

`func (o *ControlsDataPutPreviousValues) HasControlClassificationId() bool`

HasControlClassificationId returns a boolean if a field has been set.

### GetControlProcedures

`func (o *ControlsDataPutPreviousValues) GetControlProcedures() string`

GetControlProcedures returns the ControlProcedures field if non-nil, zero value otherwise.

### GetControlProceduresOk

`func (o *ControlsDataPutPreviousValues) GetControlProceduresOk() (*string, bool)`

GetControlProceduresOk returns a tuple with the ControlProcedures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlProcedures

`func (o *ControlsDataPutPreviousValues) SetControlProcedures(v string)`

SetControlProcedures sets ControlProcedures field to given value.

### HasControlProcedures

`func (o *ControlsDataPutPreviousValues) HasControlProcedures() bool`

HasControlProcedures returns a boolean if a field has been set.

### GetManagementResponse

`func (o *ControlsDataPutPreviousValues) GetManagementResponse() string`

GetManagementResponse returns the ManagementResponse field if non-nil, zero value otherwise.

### GetManagementResponseOk

`func (o *ControlsDataPutPreviousValues) GetManagementResponseOk() (*string, bool)`

GetManagementResponseOk returns a tuple with the ManagementResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementResponse

`func (o *ControlsDataPutPreviousValues) SetManagementResponse(v string)`

SetManagementResponse sets ManagementResponse field to given value.

### HasManagementResponse

`func (o *ControlsDataPutPreviousValues) HasManagementResponse() bool`

HasManagementResponse returns a boolean if a field has been set.

### GetControlObjective

`func (o *ControlsDataPutPreviousValues) GetControlObjective() string`

GetControlObjective returns the ControlObjective field if non-nil, zero value otherwise.

### GetControlObjectiveOk

`func (o *ControlsDataPutPreviousValues) GetControlObjectiveOk() (*string, bool)`

GetControlObjectiveOk returns a tuple with the ControlObjective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlObjective

`func (o *ControlsDataPutPreviousValues) SetControlObjective(v string)`

SetControlObjective sets ControlObjective field to given value.

### HasControlObjective

`func (o *ControlsDataPutPreviousValues) HasControlObjective() bool`

HasControlObjective returns a boolean if a field has been set.

### GetDependentControls

`func (o *ControlsDataPutPreviousValues) GetDependentControls() string`

GetDependentControls returns the DependentControls field if non-nil, zero value otherwise.

### GetDependentControlsOk

`func (o *ControlsDataPutPreviousValues) GetDependentControlsOk() (*string, bool)`

GetDependentControlsOk returns a tuple with the DependentControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentControls

`func (o *ControlsDataPutPreviousValues) SetDependentControls(v string)`

SetDependentControls sets DependentControls field to given value.

### HasDependentControls

`func (o *ControlsDataPutPreviousValues) HasDependentControls() bool`

HasDependentControls returns a boolean if a field has been set.

### GetCompensatingControls

`func (o *ControlsDataPutPreviousValues) GetCompensatingControls() string`

GetCompensatingControls returns the CompensatingControls field if non-nil, zero value otherwise.

### GetCompensatingControlsOk

`func (o *ControlsDataPutPreviousValues) GetCompensatingControlsOk() (*string, bool)`

GetCompensatingControlsOk returns a tuple with the CompensatingControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompensatingControls

`func (o *ControlsDataPutPreviousValues) SetCompensatingControls(v string)`

SetCompensatingControls sets CompensatingControls field to given value.

### HasCompensatingControls

`func (o *ControlsDataPutPreviousValues) HasCompensatingControls() bool`

HasCompensatingControls returns a boolean if a field has been set.

### GetAssertionNotes

`func (o *ControlsDataPutPreviousValues) GetAssertionNotes() string`

GetAssertionNotes returns the AssertionNotes field if non-nil, zero value otherwise.

### GetAssertionNotesOk

`func (o *ControlsDataPutPreviousValues) GetAssertionNotesOk() (*string, bool)`

GetAssertionNotesOk returns a tuple with the AssertionNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionNotes

`func (o *ControlsDataPutPreviousValues) SetAssertionNotes(v string)`

SetAssertionNotes sets AssertionNotes field to given value.

### HasAssertionNotes

`func (o *ControlsDataPutPreviousValues) HasAssertionNotes() bool`

HasAssertionNotes returns a boolean if a field has been set.

### GetPrecisionLevel

`func (o *ControlsDataPutPreviousValues) GetPrecisionLevel() string`

GetPrecisionLevel returns the PrecisionLevel field if non-nil, zero value otherwise.

### GetPrecisionLevelOk

`func (o *ControlsDataPutPreviousValues) GetPrecisionLevelOk() (*string, bool)`

GetPrecisionLevelOk returns a tuple with the PrecisionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecisionLevel

`func (o *ControlsDataPutPreviousValues) SetPrecisionLevel(v string)`

SetPrecisionLevel sets PrecisionLevel field to given value.

### HasPrecisionLevel

`func (o *ControlsDataPutPreviousValues) HasPrecisionLevel() bool`

HasPrecisionLevel returns a boolean if a field has been set.

### GetMrcEvidence

`func (o *ControlsDataPutPreviousValues) GetMrcEvidence() string`

GetMrcEvidence returns the MrcEvidence field if non-nil, zero value otherwise.

### GetMrcEvidenceOk

`func (o *ControlsDataPutPreviousValues) GetMrcEvidenceOk() (*string, bool)`

GetMrcEvidenceOk returns a tuple with the MrcEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrcEvidence

`func (o *ControlsDataPutPreviousValues) SetMrcEvidence(v string)`

SetMrcEvidence sets MrcEvidence field to given value.

### HasMrcEvidence

`func (o *ControlsDataPutPreviousValues) HasMrcEvidence() bool`

HasMrcEvidence returns a boolean if a field has been set.

### GetCobitReference

`func (o *ControlsDataPutPreviousValues) GetCobitReference() string`

GetCobitReference returns the CobitReference field if non-nil, zero value otherwise.

### GetCobitReferenceOk

`func (o *ControlsDataPutPreviousValues) GetCobitReferenceOk() (*string, bool)`

GetCobitReferenceOk returns a tuple with the CobitReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCobitReference

`func (o *ControlsDataPutPreviousValues) SetCobitReference(v string)`

SetCobitReference sets CobitReference field to given value.

### HasCobitReference

`func (o *ControlsDataPutPreviousValues) HasCobitReference() bool`

HasCobitReference returns a boolean if a field has been set.

### GetControlReference

`func (o *ControlsDataPutPreviousValues) GetControlReference() string`

GetControlReference returns the ControlReference field if non-nil, zero value otherwise.

### GetControlReferenceOk

`func (o *ControlsDataPutPreviousValues) GetControlReferenceOk() (*string, bool)`

GetControlReferenceOk returns a tuple with the ControlReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlReference

`func (o *ControlsDataPutPreviousValues) SetControlReference(v string)`

SetControlReference sets ControlReference field to given value.

### HasControlReference

`func (o *ControlsDataPutPreviousValues) HasControlReference() bool`

HasControlReference returns a boolean if a field has been set.

### GetMapDueDate

`func (o *ControlsDataPutPreviousValues) GetMapDueDate() string`

GetMapDueDate returns the MapDueDate field if non-nil, zero value otherwise.

### GetMapDueDateOk

`func (o *ControlsDataPutPreviousValues) GetMapDueDateOk() (*string, bool)`

GetMapDueDateOk returns a tuple with the MapDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapDueDate

`func (o *ControlsDataPutPreviousValues) SetMapDueDate(v string)`

SetMapDueDate sets MapDueDate field to given value.

### HasMapDueDate

`func (o *ControlsDataPutPreviousValues) HasMapDueDate() bool`

HasMapDueDate returns a boolean if a field has been set.

### GetIssueDiscussionDate

`func (o *ControlsDataPutPreviousValues) GetIssueDiscussionDate() string`

GetIssueDiscussionDate returns the IssueDiscussionDate field if non-nil, zero value otherwise.

### GetIssueDiscussionDateOk

`func (o *ControlsDataPutPreviousValues) GetIssueDiscussionDateOk() (*string, bool)`

GetIssueDiscussionDateOk returns a tuple with the IssueDiscussionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDiscussionDate

`func (o *ControlsDataPutPreviousValues) SetIssueDiscussionDate(v string)`

SetIssueDiscussionDate sets IssueDiscussionDate field to given value.

### HasIssueDiscussionDate

`func (o *ControlsDataPutPreviousValues) HasIssueDiscussionDate() bool`

HasIssueDiscussionDate returns a boolean if a field has been set.

### GetIssueIndividualsPresent

`func (o *ControlsDataPutPreviousValues) GetIssueIndividualsPresent() string`

GetIssueIndividualsPresent returns the IssueIndividualsPresent field if non-nil, zero value otherwise.

### GetIssueIndividualsPresentOk

`func (o *ControlsDataPutPreviousValues) GetIssueIndividualsPresentOk() (*string, bool)`

GetIssueIndividualsPresentOk returns a tuple with the IssueIndividualsPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueIndividualsPresent

`func (o *ControlsDataPutPreviousValues) SetIssueIndividualsPresent(v string)`

SetIssueIndividualsPresent sets IssueIndividualsPresent field to given value.

### HasIssueIndividualsPresent

`func (o *ControlsDataPutPreviousValues) HasIssueIndividualsPresent() bool`

HasIssueIndividualsPresent returns a boolean if a field has been set.

### GetReviewerUserId

`func (o *ControlsDataPutPreviousValues) GetReviewerUserId() int32`

GetReviewerUserId returns the ReviewerUserId field if non-nil, zero value otherwise.

### GetReviewerUserIdOk

`func (o *ControlsDataPutPreviousValues) GetReviewerUserIdOk() (*int32, bool)`

GetReviewerUserIdOk returns a tuple with the ReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerUserId

`func (o *ControlsDataPutPreviousValues) SetReviewerUserId(v int32)`

SetReviewerUserId sets ReviewerUserId field to given value.

### HasReviewerUserId

`func (o *ControlsDataPutPreviousValues) HasReviewerUserId() bool`

HasReviewerUserId returns a boolean if a field has been set.

### GetAnnualPopulationSize

`func (o *ControlsDataPutPreviousValues) GetAnnualPopulationSize() string`

GetAnnualPopulationSize returns the AnnualPopulationSize field if non-nil, zero value otherwise.

### GetAnnualPopulationSizeOk

`func (o *ControlsDataPutPreviousValues) GetAnnualPopulationSizeOk() (*string, bool)`

GetAnnualPopulationSizeOk returns a tuple with the AnnualPopulationSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualPopulationSize

`func (o *ControlsDataPutPreviousValues) SetAnnualPopulationSize(v string)`

SetAnnualPopulationSize sets AnnualPopulationSize field to given value.

### HasAnnualPopulationSize

`func (o *ControlsDataPutPreviousValues) HasAnnualPopulationSize() bool`

HasAnnualPopulationSize returns a boolean if a field has been set.

### GetControlInherentRiskOptionId

`func (o *ControlsDataPutPreviousValues) GetControlInherentRiskOptionId() int32`

GetControlInherentRiskOptionId returns the ControlInherentRiskOptionId field if non-nil, zero value otherwise.

### GetControlInherentRiskOptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlInherentRiskOptionIdOk() (*int32, bool)`

GetControlInherentRiskOptionIdOk returns a tuple with the ControlInherentRiskOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlInherentRiskOptionId

`func (o *ControlsDataPutPreviousValues) SetControlInherentRiskOptionId(v int32)`

SetControlInherentRiskOptionId sets ControlInherentRiskOptionId field to given value.

### HasControlInherentRiskOptionId

`func (o *ControlsDataPutPreviousValues) HasControlInherentRiskOptionId() bool`

HasControlInherentRiskOptionId returns a boolean if a field has been set.

### GetControlRoutineOptionId

`func (o *ControlsDataPutPreviousValues) GetControlRoutineOptionId() int32`

GetControlRoutineOptionId returns the ControlRoutineOptionId field if non-nil, zero value otherwise.

### GetControlRoutineOptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlRoutineOptionIdOk() (*int32, bool)`

GetControlRoutineOptionIdOk returns a tuple with the ControlRoutineOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlRoutineOptionId

`func (o *ControlsDataPutPreviousValues) SetControlRoutineOptionId(v int32)`

SetControlRoutineOptionId sets ControlRoutineOptionId field to given value.

### HasControlRoutineOptionId

`func (o *ControlsDataPutPreviousValues) HasControlRoutineOptionId() bool`

HasControlRoutineOptionId returns a boolean if a field has been set.

### GetControlJudgementDegreeOptionId

`func (o *ControlsDataPutPreviousValues) GetControlJudgementDegreeOptionId() int32`

GetControlJudgementDegreeOptionId returns the ControlJudgementDegreeOptionId field if non-nil, zero value otherwise.

### GetControlJudgementDegreeOptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlJudgementDegreeOptionIdOk() (*int32, bool)`

GetControlJudgementDegreeOptionIdOk returns a tuple with the ControlJudgementDegreeOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlJudgementDegreeOptionId

`func (o *ControlsDataPutPreviousValues) SetControlJudgementDegreeOptionId(v int32)`

SetControlJudgementDegreeOptionId sets ControlJudgementDegreeOptionId field to given value.

### HasControlJudgementDegreeOptionId

`func (o *ControlsDataPutPreviousValues) HasControlJudgementDegreeOptionId() bool`

HasControlJudgementDegreeOptionId returns a boolean if a field has been set.

### GetControlManagementOverrideOptionId

`func (o *ControlsDataPutPreviousValues) GetControlManagementOverrideOptionId() int32`

GetControlManagementOverrideOptionId returns the ControlManagementOverrideOptionId field if non-nil, zero value otherwise.

### GetControlManagementOverrideOptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlManagementOverrideOptionIdOk() (*int32, bool)`

GetControlManagementOverrideOptionIdOk returns a tuple with the ControlManagementOverrideOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlManagementOverrideOptionId

`func (o *ControlsDataPutPreviousValues) SetControlManagementOverrideOptionId(v int32)`

SetControlManagementOverrideOptionId sets ControlManagementOverrideOptionId field to given value.

### HasControlManagementOverrideOptionId

`func (o *ControlsDataPutPreviousValues) HasControlManagementOverrideOptionId() bool`

HasControlManagementOverrideOptionId returns a boolean if a field has been set.

### GetControlComplexityOptionId

`func (o *ControlsDataPutPreviousValues) GetControlComplexityOptionId() int32`

GetControlComplexityOptionId returns the ControlComplexityOptionId field if non-nil, zero value otherwise.

### GetControlComplexityOptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlComplexityOptionIdOk() (*int32, bool)`

GetControlComplexityOptionIdOk returns a tuple with the ControlComplexityOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlComplexityOptionId

`func (o *ControlsDataPutPreviousValues) SetControlComplexityOptionId(v int32)`

SetControlComplexityOptionId sets ControlComplexityOptionId field to given value.

### HasControlComplexityOptionId

`func (o *ControlsDataPutPreviousValues) HasControlComplexityOptionId() bool`

HasControlComplexityOptionId returns a boolean if a field has been set.

### GetExternalAuditorViewableStatus

`func (o *ControlsDataPutPreviousValues) GetExternalAuditorViewableStatus() string`

GetExternalAuditorViewableStatus returns the ExternalAuditorViewableStatus field if non-nil, zero value otherwise.

### GetExternalAuditorViewableStatusOk

`func (o *ControlsDataPutPreviousValues) GetExternalAuditorViewableStatusOk() (*string, bool)`

GetExternalAuditorViewableStatusOk returns a tuple with the ExternalAuditorViewableStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAuditorViewableStatus

`func (o *ControlsDataPutPreviousValues) SetExternalAuditorViewableStatus(v string)`

SetExternalAuditorViewableStatus sets ExternalAuditorViewableStatus field to given value.

### HasExternalAuditorViewableStatus

`func (o *ControlsDataPutPreviousValues) HasExternalAuditorViewableStatus() bool`

HasExternalAuditorViewableStatus returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *ControlsDataPutPreviousValues) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *ControlsDataPutPreviousValues) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *ControlsDataPutPreviousValues) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *ControlsDataPutPreviousValues) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetBaselineResultOptionId

`func (o *ControlsDataPutPreviousValues) GetBaselineResultOptionId() int32`

GetBaselineResultOptionId returns the BaselineResultOptionId field if non-nil, zero value otherwise.

### GetBaselineResultOptionIdOk

`func (o *ControlsDataPutPreviousValues) GetBaselineResultOptionIdOk() (*int32, bool)`

GetBaselineResultOptionIdOk returns a tuple with the BaselineResultOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineResultOptionId

`func (o *ControlsDataPutPreviousValues) SetBaselineResultOptionId(v int32)`

SetBaselineResultOptionId sets BaselineResultOptionId field to given value.

### HasBaselineResultOptionId

`func (o *ControlsDataPutPreviousValues) HasBaselineResultOptionId() bool`

HasBaselineResultOptionId returns a boolean if a field has been set.

### GetModificationDescription

`func (o *ControlsDataPutPreviousValues) GetModificationDescription() string`

GetModificationDescription returns the ModificationDescription field if non-nil, zero value otherwise.

### GetModificationDescriptionOk

`func (o *ControlsDataPutPreviousValues) GetModificationDescriptionOk() (*string, bool)`

GetModificationDescriptionOk returns a tuple with the ModificationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDescription

`func (o *ControlsDataPutPreviousValues) SetModificationDescription(v string)`

SetModificationDescription sets ModificationDescription field to given value.

### HasModificationDescription

`func (o *ControlsDataPutPreviousValues) HasModificationDescription() bool`

HasModificationDescription returns a boolean if a field has been set.

### GetReviewerNotes

`func (o *ControlsDataPutPreviousValues) GetReviewerNotes() string`

GetReviewerNotes returns the ReviewerNotes field if non-nil, zero value otherwise.

### GetReviewerNotesOk

`func (o *ControlsDataPutPreviousValues) GetReviewerNotesOk() (*string, bool)`

GetReviewerNotesOk returns a tuple with the ReviewerNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerNotes

`func (o *ControlsDataPutPreviousValues) SetReviewerNotes(v string)`

SetReviewerNotes sets ReviewerNotes field to given value.

### HasReviewerNotes

`func (o *ControlsDataPutPreviousValues) HasReviewerNotes() bool`

HasReviewerNotes returns a boolean if a field has been set.

### GetRationale

`func (o *ControlsDataPutPreviousValues) GetRationale() string`

GetRationale returns the Rationale field if non-nil, zero value otherwise.

### GetRationaleOk

`func (o *ControlsDataPutPreviousValues) GetRationaleOk() (*string, bool)`

GetRationaleOk returns a tuple with the Rationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRationale

`func (o *ControlsDataPutPreviousValues) SetRationale(v string)`

SetRationale sets Rationale field to given value.

### HasRationale

`func (o *ControlsDataPutPreviousValues) HasRationale() bool`

HasRationale returns a boolean if a field has been set.

### GetBaselineDate

`func (o *ControlsDataPutPreviousValues) GetBaselineDate() string`

GetBaselineDate returns the BaselineDate field if non-nil, zero value otherwise.

### GetBaselineDateOk

`func (o *ControlsDataPutPreviousValues) GetBaselineDateOk() (*string, bool)`

GetBaselineDateOk returns a tuple with the BaselineDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineDate

`func (o *ControlsDataPutPreviousValues) SetBaselineDate(v string)`

SetBaselineDate sets BaselineDate field to given value.

### HasBaselineDate

`func (o *ControlsDataPutPreviousValues) HasBaselineDate() bool`

HasBaselineDate returns a boolean if a field has been set.

### GetUid

`func (o *ControlsDataPutPreviousValues) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ControlsDataPutPreviousValues) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ControlsDataPutPreviousValues) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ControlsDataPutPreviousValues) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetProcessId

`func (o *ControlsDataPutPreviousValues) GetProcessId() int32`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *ControlsDataPutPreviousValues) GetProcessIdOk() (*int32, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *ControlsDataPutPreviousValues) SetProcessId(v int32)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *ControlsDataPutPreviousValues) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### GetSubprocessId

`func (o *ControlsDataPutPreviousValues) GetSubprocessId() int32`

GetSubprocessId returns the SubprocessId field if non-nil, zero value otherwise.

### GetSubprocessIdOk

`func (o *ControlsDataPutPreviousValues) GetSubprocessIdOk() (*int32, bool)`

GetSubprocessIdOk returns a tuple with the SubprocessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubprocessId

`func (o *ControlsDataPutPreviousValues) SetSubprocessId(v int32)`

SetSubprocessId sets SubprocessId field to given value.

### HasSubprocessId

`func (o *ControlsDataPutPreviousValues) HasSubprocessId() bool`

HasSubprocessId returns a boolean if a field has been set.

### GetLastModificationDate

`func (o *ControlsDataPutPreviousValues) GetLastModificationDate() string`

GetLastModificationDate returns the LastModificationDate field if non-nil, zero value otherwise.

### GetLastModificationDateOk

`func (o *ControlsDataPutPreviousValues) GetLastModificationDateOk() (*string, bool)`

GetLastModificationDateOk returns a tuple with the LastModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationDate

`func (o *ControlsDataPutPreviousValues) SetLastModificationDate(v string)`

SetLastModificationDate sets LastModificationDate field to given value.

### HasLastModificationDate

`func (o *ControlsDataPutPreviousValues) HasLastModificationDate() bool`

HasLastModificationDate returns a boolean if a field has been set.

### GetLastReviewDate

`func (o *ControlsDataPutPreviousValues) GetLastReviewDate() string`

GetLastReviewDate returns the LastReviewDate field if non-nil, zero value otherwise.

### GetLastReviewDateOk

`func (o *ControlsDataPutPreviousValues) GetLastReviewDateOk() (*string, bool)`

GetLastReviewDateOk returns a tuple with the LastReviewDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReviewDate

`func (o *ControlsDataPutPreviousValues) SetLastReviewDate(v string)`

SetLastReviewDate sets LastReviewDate field to given value.

### HasLastReviewDate

`func (o *ControlsDataPutPreviousValues) HasLastReviewDate() bool`

HasLastReviewDate returns a boolean if a field has been set.

### GetRelianceOptionId

`func (o *ControlsDataPutPreviousValues) GetRelianceOptionId() int32`

GetRelianceOptionId returns the RelianceOptionId field if non-nil, zero value otherwise.

### GetRelianceOptionIdOk

`func (o *ControlsDataPutPreviousValues) GetRelianceOptionIdOk() (*int32, bool)`

GetRelianceOptionIdOk returns a tuple with the RelianceOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelianceOptionId

`func (o *ControlsDataPutPreviousValues) SetRelianceOptionId(v int32)`

SetRelianceOptionId sets RelianceOptionId field to given value.

### HasRelianceOptionId

`func (o *ControlsDataPutPreviousValues) HasRelianceOptionId() bool`

HasRelianceOptionId returns a boolean if a field has been set.

### GetControlCustomSelect1OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect1OptionId() int32`

GetControlCustomSelect1OptionId returns the ControlCustomSelect1OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect1OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect1OptionIdOk() (*int32, bool)`

GetControlCustomSelect1OptionIdOk returns a tuple with the ControlCustomSelect1OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect1OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect1OptionId(v int32)`

SetControlCustomSelect1OptionId sets ControlCustomSelect1OptionId field to given value.

### HasControlCustomSelect1OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect1OptionId() bool`

HasControlCustomSelect1OptionId returns a boolean if a field has been set.

### GetCoordinatorUserId

`func (o *ControlsDataPutPreviousValues) GetCoordinatorUserId() int32`

GetCoordinatorUserId returns the CoordinatorUserId field if non-nil, zero value otherwise.

### GetCoordinatorUserIdOk

`func (o *ControlsDataPutPreviousValues) GetCoordinatorUserIdOk() (*int32, bool)`

GetCoordinatorUserIdOk returns a tuple with the CoordinatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinatorUserId

`func (o *ControlsDataPutPreviousValues) SetCoordinatorUserId(v int32)`

SetCoordinatorUserId sets CoordinatorUserId field to given value.

### HasCoordinatorUserId

`func (o *ControlsDataPutPreviousValues) HasCoordinatorUserId() bool`

HasCoordinatorUserId returns a boolean if a field has been set.

### GetControlCustomText1

`func (o *ControlsDataPutPreviousValues) GetControlCustomText1() string`

GetControlCustomText1 returns the ControlCustomText1 field if non-nil, zero value otherwise.

### GetControlCustomText1Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText1Ok() (*string, bool)`

GetControlCustomText1Ok returns a tuple with the ControlCustomText1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText1

`func (o *ControlsDataPutPreviousValues) SetControlCustomText1(v string)`

SetControlCustomText1 sets ControlCustomText1 field to given value.

### HasControlCustomText1

`func (o *ControlsDataPutPreviousValues) HasControlCustomText1() bool`

HasControlCustomText1 returns a boolean if a field has been set.

### GetControlCustomText2

`func (o *ControlsDataPutPreviousValues) GetControlCustomText2() string`

GetControlCustomText2 returns the ControlCustomText2 field if non-nil, zero value otherwise.

### GetControlCustomText2Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText2Ok() (*string, bool)`

GetControlCustomText2Ok returns a tuple with the ControlCustomText2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText2

`func (o *ControlsDataPutPreviousValues) SetControlCustomText2(v string)`

SetControlCustomText2 sets ControlCustomText2 field to given value.

### HasControlCustomText2

`func (o *ControlsDataPutPreviousValues) HasControlCustomText2() bool`

HasControlCustomText2 returns a boolean if a field has been set.

### GetControlCustomText3

`func (o *ControlsDataPutPreviousValues) GetControlCustomText3() string`

GetControlCustomText3 returns the ControlCustomText3 field if non-nil, zero value otherwise.

### GetControlCustomText3Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText3Ok() (*string, bool)`

GetControlCustomText3Ok returns a tuple with the ControlCustomText3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText3

`func (o *ControlsDataPutPreviousValues) SetControlCustomText3(v string)`

SetControlCustomText3 sets ControlCustomText3 field to given value.

### HasControlCustomText3

`func (o *ControlsDataPutPreviousValues) HasControlCustomText3() bool`

HasControlCustomText3 returns a boolean if a field has been set.

### GetControlCustomText4

`func (o *ControlsDataPutPreviousValues) GetControlCustomText4() string`

GetControlCustomText4 returns the ControlCustomText4 field if non-nil, zero value otherwise.

### GetControlCustomText4Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText4Ok() (*string, bool)`

GetControlCustomText4Ok returns a tuple with the ControlCustomText4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText4

`func (o *ControlsDataPutPreviousValues) SetControlCustomText4(v string)`

SetControlCustomText4 sets ControlCustomText4 field to given value.

### HasControlCustomText4

`func (o *ControlsDataPutPreviousValues) HasControlCustomText4() bool`

HasControlCustomText4 returns a boolean if a field has been set.

### GetControlCustomSelect2OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect2OptionId() int32`

GetControlCustomSelect2OptionId returns the ControlCustomSelect2OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect2OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect2OptionIdOk() (*int32, bool)`

GetControlCustomSelect2OptionIdOk returns a tuple with the ControlCustomSelect2OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect2OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect2OptionId(v int32)`

SetControlCustomSelect2OptionId sets ControlCustomSelect2OptionId field to given value.

### HasControlCustomSelect2OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect2OptionId() bool`

HasControlCustomSelect2OptionId returns a boolean if a field has been set.

### GetControlCustomSelect3OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect3OptionId() int32`

GetControlCustomSelect3OptionId returns the ControlCustomSelect3OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect3OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect3OptionIdOk() (*int32, bool)`

GetControlCustomSelect3OptionIdOk returns a tuple with the ControlCustomSelect3OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect3OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect3OptionId(v int32)`

SetControlCustomSelect3OptionId sets ControlCustomSelect3OptionId field to given value.

### HasControlCustomSelect3OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect3OptionId() bool`

HasControlCustomSelect3OptionId returns a boolean if a field has been set.

### GetControlCustomSelect4OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect4OptionId() int32`

GetControlCustomSelect4OptionId returns the ControlCustomSelect4OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect4OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect4OptionIdOk() (*int32, bool)`

GetControlCustomSelect4OptionIdOk returns a tuple with the ControlCustomSelect4OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect4OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect4OptionId(v int32)`

SetControlCustomSelect4OptionId sets ControlCustomSelect4OptionId field to given value.

### HasControlCustomSelect4OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect4OptionId() bool`

HasControlCustomSelect4OptionId returns a boolean if a field has been set.

### GetControlCustomSelect5OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect5OptionId() int32`

GetControlCustomSelect5OptionId returns the ControlCustomSelect5OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect5OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect5OptionIdOk() (*int32, bool)`

GetControlCustomSelect5OptionIdOk returns a tuple with the ControlCustomSelect5OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect5OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect5OptionId(v int32)`

SetControlCustomSelect5OptionId sets ControlCustomSelect5OptionId field to given value.

### HasControlCustomSelect5OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect5OptionId() bool`

HasControlCustomSelect5OptionId returns a boolean if a field has been set.

### GetReferenceMeta

`func (o *ControlsDataPutPreviousValues) GetReferenceMeta() interface{}`

GetReferenceMeta returns the ReferenceMeta field if non-nil, zero value otherwise.

### GetReferenceMetaOk

`func (o *ControlsDataPutPreviousValues) GetReferenceMetaOk() (*interface{}, bool)`

GetReferenceMetaOk returns a tuple with the ReferenceMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMeta

`func (o *ControlsDataPutPreviousValues) SetReferenceMeta(v interface{})`

SetReferenceMeta sets ReferenceMeta field to given value.

### HasReferenceMeta

`func (o *ControlsDataPutPreviousValues) HasReferenceMeta() bool`

HasReferenceMeta returns a boolean if a field has been set.

### SetReferenceMetaNil

`func (o *ControlsDataPutPreviousValues) SetReferenceMetaNil(b bool)`

 SetReferenceMetaNil sets the value for ReferenceMeta to be an explicit nil

### UnsetReferenceMeta
`func (o *ControlsDataPutPreviousValues) UnsetReferenceMeta()`

UnsetReferenceMeta ensures that no value is present for ReferenceMeta, not even an explicit nil
### GetComplianceReviewerUserId

`func (o *ControlsDataPutPreviousValues) GetComplianceReviewerUserId() int32`

GetComplianceReviewerUserId returns the ComplianceReviewerUserId field if non-nil, zero value otherwise.

### GetComplianceReviewerUserIdOk

`func (o *ControlsDataPutPreviousValues) GetComplianceReviewerUserIdOk() (*int32, bool)`

GetComplianceReviewerUserIdOk returns a tuple with the ComplianceReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceReviewerUserId

`func (o *ControlsDataPutPreviousValues) SetComplianceReviewerUserId(v int32)`

SetComplianceReviewerUserId sets ComplianceReviewerUserId field to given value.

### HasComplianceReviewerUserId

`func (o *ControlsDataPutPreviousValues) HasComplianceReviewerUserId() bool`

HasComplianceReviewerUserId returns a boolean if a field has been set.

### GetControlCustomText5

`func (o *ControlsDataPutPreviousValues) GetControlCustomText5() string`

GetControlCustomText5 returns the ControlCustomText5 field if non-nil, zero value otherwise.

### GetControlCustomText5Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText5Ok() (*string, bool)`

GetControlCustomText5Ok returns a tuple with the ControlCustomText5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText5

`func (o *ControlsDataPutPreviousValues) SetControlCustomText5(v string)`

SetControlCustomText5 sets ControlCustomText5 field to given value.

### HasControlCustomText5

`func (o *ControlsDataPutPreviousValues) HasControlCustomText5() bool`

HasControlCustomText5 returns a boolean if a field has been set.

### GetControlCustomText6

`func (o *ControlsDataPutPreviousValues) GetControlCustomText6() string`

GetControlCustomText6 returns the ControlCustomText6 field if non-nil, zero value otherwise.

### GetControlCustomText6Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText6Ok() (*string, bool)`

GetControlCustomText6Ok returns a tuple with the ControlCustomText6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText6

`func (o *ControlsDataPutPreviousValues) SetControlCustomText6(v string)`

SetControlCustomText6 sets ControlCustomText6 field to given value.

### HasControlCustomText6

`func (o *ControlsDataPutPreviousValues) HasControlCustomText6() bool`

HasControlCustomText6 returns a boolean if a field has been set.

### GetControlCustomText7

`func (o *ControlsDataPutPreviousValues) GetControlCustomText7() string`

GetControlCustomText7 returns the ControlCustomText7 field if non-nil, zero value otherwise.

### GetControlCustomText7Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText7Ok() (*string, bool)`

GetControlCustomText7Ok returns a tuple with the ControlCustomText7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText7

`func (o *ControlsDataPutPreviousValues) SetControlCustomText7(v string)`

SetControlCustomText7 sets ControlCustomText7 field to given value.

### HasControlCustomText7

`func (o *ControlsDataPutPreviousValues) HasControlCustomText7() bool`

HasControlCustomText7 returns a boolean if a field has been set.

### GetControlCustomText8

`func (o *ControlsDataPutPreviousValues) GetControlCustomText8() string`

GetControlCustomText8 returns the ControlCustomText8 field if non-nil, zero value otherwise.

### GetControlCustomText8Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText8Ok() (*string, bool)`

GetControlCustomText8Ok returns a tuple with the ControlCustomText8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText8

`func (o *ControlsDataPutPreviousValues) SetControlCustomText8(v string)`

SetControlCustomText8 sets ControlCustomText8 field to given value.

### HasControlCustomText8

`func (o *ControlsDataPutPreviousValues) HasControlCustomText8() bool`

HasControlCustomText8 returns a boolean if a field has been set.

### GetControlCustomText9

`func (o *ControlsDataPutPreviousValues) GetControlCustomText9() string`

GetControlCustomText9 returns the ControlCustomText9 field if non-nil, zero value otherwise.

### GetControlCustomText9Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText9Ok() (*string, bool)`

GetControlCustomText9Ok returns a tuple with the ControlCustomText9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText9

`func (o *ControlsDataPutPreviousValues) SetControlCustomText9(v string)`

SetControlCustomText9 sets ControlCustomText9 field to given value.

### HasControlCustomText9

`func (o *ControlsDataPutPreviousValues) HasControlCustomText9() bool`

HasControlCustomText9 returns a boolean if a field has been set.

### GetControlCustomText10

`func (o *ControlsDataPutPreviousValues) GetControlCustomText10() string`

GetControlCustomText10 returns the ControlCustomText10 field if non-nil, zero value otherwise.

### GetControlCustomText10Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText10Ok() (*string, bool)`

GetControlCustomText10Ok returns a tuple with the ControlCustomText10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText10

`func (o *ControlsDataPutPreviousValues) SetControlCustomText10(v string)`

SetControlCustomText10 sets ControlCustomText10 field to given value.

### HasControlCustomText10

`func (o *ControlsDataPutPreviousValues) HasControlCustomText10() bool`

HasControlCustomText10 returns a boolean if a field has been set.

### GetControlCustomText11

`func (o *ControlsDataPutPreviousValues) GetControlCustomText11() string`

GetControlCustomText11 returns the ControlCustomText11 field if non-nil, zero value otherwise.

### GetControlCustomText11Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText11Ok() (*string, bool)`

GetControlCustomText11Ok returns a tuple with the ControlCustomText11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText11

`func (o *ControlsDataPutPreviousValues) SetControlCustomText11(v string)`

SetControlCustomText11 sets ControlCustomText11 field to given value.

### HasControlCustomText11

`func (o *ControlsDataPutPreviousValues) HasControlCustomText11() bool`

HasControlCustomText11 returns a boolean if a field has been set.

### GetControlCustomText12

`func (o *ControlsDataPutPreviousValues) GetControlCustomText12() string`

GetControlCustomText12 returns the ControlCustomText12 field if non-nil, zero value otherwise.

### GetControlCustomText12Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText12Ok() (*string, bool)`

GetControlCustomText12Ok returns a tuple with the ControlCustomText12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText12

`func (o *ControlsDataPutPreviousValues) SetControlCustomText12(v string)`

SetControlCustomText12 sets ControlCustomText12 field to given value.

### HasControlCustomText12

`func (o *ControlsDataPutPreviousValues) HasControlCustomText12() bool`

HasControlCustomText12 returns a boolean if a field has been set.

### GetName

`func (o *ControlsDataPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControlsDataPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControlsDataPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ControlsDataPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetXuid

`func (o *ControlsDataPutPreviousValues) GetXuid() string`

GetXuid returns the Xuid field if non-nil, zero value otherwise.

### GetXuidOk

`func (o *ControlsDataPutPreviousValues) GetXuidOk() (*string, bool)`

GetXuidOk returns a tuple with the Xuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXuid

`func (o *ControlsDataPutPreviousValues) SetXuid(v string)`

SetXuid sets Xuid field to given value.

### HasXuid

`func (o *ControlsDataPutPreviousValues) HasXuid() bool`

HasXuid returns a boolean if a field has been set.

### GetScopes

`func (o *ControlsDataPutPreviousValues) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ControlsDataPutPreviousValues) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ControlsDataPutPreviousValues) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ControlsDataPutPreviousValues) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *ControlsDataPutPreviousValues) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *ControlsDataPutPreviousValues) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetFieldData

`func (o *ControlsDataPutPreviousValues) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *ControlsDataPutPreviousValues) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *ControlsDataPutPreviousValues) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *ControlsDataPutPreviousValues) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *ControlsDataPutPreviousValues) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *ControlsDataPutPreviousValues) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil
### GetControlCustomSelect6OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect6OptionId() int32`

GetControlCustomSelect6OptionId returns the ControlCustomSelect6OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect6OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect6OptionIdOk() (*int32, bool)`

GetControlCustomSelect6OptionIdOk returns a tuple with the ControlCustomSelect6OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect6OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect6OptionId(v int32)`

SetControlCustomSelect6OptionId sets ControlCustomSelect6OptionId field to given value.

### HasControlCustomSelect6OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect6OptionId() bool`

HasControlCustomSelect6OptionId returns a boolean if a field has been set.

### GetControlCustomSelect7OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect7OptionId() int32`

GetControlCustomSelect7OptionId returns the ControlCustomSelect7OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect7OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect7OptionIdOk() (*int32, bool)`

GetControlCustomSelect7OptionIdOk returns a tuple with the ControlCustomSelect7OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect7OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect7OptionId(v int32)`

SetControlCustomSelect7OptionId sets ControlCustomSelect7OptionId field to given value.

### HasControlCustomSelect7OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect7OptionId() bool`

HasControlCustomSelect7OptionId returns a boolean if a field has been set.

### GetControlCustomSelect8OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect8OptionId() int32`

GetControlCustomSelect8OptionId returns the ControlCustomSelect8OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect8OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect8OptionIdOk() (*int32, bool)`

GetControlCustomSelect8OptionIdOk returns a tuple with the ControlCustomSelect8OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect8OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect8OptionId(v int32)`

SetControlCustomSelect8OptionId sets ControlCustomSelect8OptionId field to given value.

### HasControlCustomSelect8OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect8OptionId() bool`

HasControlCustomSelect8OptionId returns a boolean if a field has been set.

### GetControlCustomSelect9OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect9OptionId() int32`

GetControlCustomSelect9OptionId returns the ControlCustomSelect9OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect9OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect9OptionIdOk() (*int32, bool)`

GetControlCustomSelect9OptionIdOk returns a tuple with the ControlCustomSelect9OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect9OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect9OptionId(v int32)`

SetControlCustomSelect9OptionId sets ControlCustomSelect9OptionId field to given value.

### HasControlCustomSelect9OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect9OptionId() bool`

HasControlCustomSelect9OptionId returns a boolean if a field has been set.

### GetControlCustomSelect10OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect10OptionId() int32`

GetControlCustomSelect10OptionId returns the ControlCustomSelect10OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect10OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect10OptionIdOk() (*int32, bool)`

GetControlCustomSelect10OptionIdOk returns a tuple with the ControlCustomSelect10OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect10OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect10OptionId(v int32)`

SetControlCustomSelect10OptionId sets ControlCustomSelect10OptionId field to given value.

### HasControlCustomSelect10OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect10OptionId() bool`

HasControlCustomSelect10OptionId returns a boolean if a field has been set.

### GetControlCustomSelect11OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect11OptionId() int32`

GetControlCustomSelect11OptionId returns the ControlCustomSelect11OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect11OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect11OptionIdOk() (*int32, bool)`

GetControlCustomSelect11OptionIdOk returns a tuple with the ControlCustomSelect11OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect11OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect11OptionId(v int32)`

SetControlCustomSelect11OptionId sets ControlCustomSelect11OptionId field to given value.

### HasControlCustomSelect11OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect11OptionId() bool`

HasControlCustomSelect11OptionId returns a boolean if a field has been set.

### GetControlCustomSelect12OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect12OptionId() int32`

GetControlCustomSelect12OptionId returns the ControlCustomSelect12OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect12OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect12OptionIdOk() (*int32, bool)`

GetControlCustomSelect12OptionIdOk returns a tuple with the ControlCustomSelect12OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect12OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect12OptionId(v int32)`

SetControlCustomSelect12OptionId sets ControlCustomSelect12OptionId field to given value.

### HasControlCustomSelect12OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect12OptionId() bool`

HasControlCustomSelect12OptionId returns a boolean if a field has been set.

### GetCustomUserSelect1UserId

`func (o *ControlsDataPutPreviousValues) GetCustomUserSelect1UserId() int32`

GetCustomUserSelect1UserId returns the CustomUserSelect1UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect1UserIdOk

`func (o *ControlsDataPutPreviousValues) GetCustomUserSelect1UserIdOk() (*int32, bool)`

GetCustomUserSelect1UserIdOk returns a tuple with the CustomUserSelect1UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect1UserId

`func (o *ControlsDataPutPreviousValues) SetCustomUserSelect1UserId(v int32)`

SetCustomUserSelect1UserId sets CustomUserSelect1UserId field to given value.

### HasCustomUserSelect1UserId

`func (o *ControlsDataPutPreviousValues) HasCustomUserSelect1UserId() bool`

HasCustomUserSelect1UserId returns a boolean if a field has been set.

### GetCustomUserSelect2UserId

`func (o *ControlsDataPutPreviousValues) GetCustomUserSelect2UserId() int32`

GetCustomUserSelect2UserId returns the CustomUserSelect2UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect2UserIdOk

`func (o *ControlsDataPutPreviousValues) GetCustomUserSelect2UserIdOk() (*int32, bool)`

GetCustomUserSelect2UserIdOk returns a tuple with the CustomUserSelect2UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect2UserId

`func (o *ControlsDataPutPreviousValues) SetCustomUserSelect2UserId(v int32)`

SetCustomUserSelect2UserId sets CustomUserSelect2UserId field to given value.

### HasCustomUserSelect2UserId

`func (o *ControlsDataPutPreviousValues) HasCustomUserSelect2UserId() bool`

HasCustomUserSelect2UserId returns a boolean if a field has been set.

### GetCustomUserSelect3UserId

`func (o *ControlsDataPutPreviousValues) GetCustomUserSelect3UserId() int32`

GetCustomUserSelect3UserId returns the CustomUserSelect3UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect3UserIdOk

`func (o *ControlsDataPutPreviousValues) GetCustomUserSelect3UserIdOk() (*int32, bool)`

GetCustomUserSelect3UserIdOk returns a tuple with the CustomUserSelect3UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect3UserId

`func (o *ControlsDataPutPreviousValues) SetCustomUserSelect3UserId(v int32)`

SetCustomUserSelect3UserId sets CustomUserSelect3UserId field to given value.

### HasCustomUserSelect3UserId

`func (o *ControlsDataPutPreviousValues) HasCustomUserSelect3UserId() bool`

HasCustomUserSelect3UserId returns a boolean if a field has been set.

### GetCustomUserSelect4UserId

`func (o *ControlsDataPutPreviousValues) GetCustomUserSelect4UserId() int32`

GetCustomUserSelect4UserId returns the CustomUserSelect4UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect4UserIdOk

`func (o *ControlsDataPutPreviousValues) GetCustomUserSelect4UserIdOk() (*int32, bool)`

GetCustomUserSelect4UserIdOk returns a tuple with the CustomUserSelect4UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect4UserId

`func (o *ControlsDataPutPreviousValues) SetCustomUserSelect4UserId(v int32)`

SetCustomUserSelect4UserId sets CustomUserSelect4UserId field to given value.

### HasCustomUserSelect4UserId

`func (o *ControlsDataPutPreviousValues) HasCustomUserSelect4UserId() bool`

HasCustomUserSelect4UserId returns a boolean if a field has been set.

### GetControlCustomText13

`func (o *ControlsDataPutPreviousValues) GetControlCustomText13() string`

GetControlCustomText13 returns the ControlCustomText13 field if non-nil, zero value otherwise.

### GetControlCustomText13Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText13Ok() (*string, bool)`

GetControlCustomText13Ok returns a tuple with the ControlCustomText13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText13

`func (o *ControlsDataPutPreviousValues) SetControlCustomText13(v string)`

SetControlCustomText13 sets ControlCustomText13 field to given value.

### HasControlCustomText13

`func (o *ControlsDataPutPreviousValues) HasControlCustomText13() bool`

HasControlCustomText13 returns a boolean if a field has been set.

### GetControlCustomText14

`func (o *ControlsDataPutPreviousValues) GetControlCustomText14() string`

GetControlCustomText14 returns the ControlCustomText14 field if non-nil, zero value otherwise.

### GetControlCustomText14Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText14Ok() (*string, bool)`

GetControlCustomText14Ok returns a tuple with the ControlCustomText14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText14

`func (o *ControlsDataPutPreviousValues) SetControlCustomText14(v string)`

SetControlCustomText14 sets ControlCustomText14 field to given value.

### HasControlCustomText14

`func (o *ControlsDataPutPreviousValues) HasControlCustomText14() bool`

HasControlCustomText14 returns a boolean if a field has been set.

### GetControlCustomText15

`func (o *ControlsDataPutPreviousValues) GetControlCustomText15() string`

GetControlCustomText15 returns the ControlCustomText15 field if non-nil, zero value otherwise.

### GetControlCustomText15Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText15Ok() (*string, bool)`

GetControlCustomText15Ok returns a tuple with the ControlCustomText15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText15

`func (o *ControlsDataPutPreviousValues) SetControlCustomText15(v string)`

SetControlCustomText15 sets ControlCustomText15 field to given value.

### HasControlCustomText15

`func (o *ControlsDataPutPreviousValues) HasControlCustomText15() bool`

HasControlCustomText15 returns a boolean if a field has been set.

### GetControlCustomText16

`func (o *ControlsDataPutPreviousValues) GetControlCustomText16() string`

GetControlCustomText16 returns the ControlCustomText16 field if non-nil, zero value otherwise.

### GetControlCustomText16Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText16Ok() (*string, bool)`

GetControlCustomText16Ok returns a tuple with the ControlCustomText16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText16

`func (o *ControlsDataPutPreviousValues) SetControlCustomText16(v string)`

SetControlCustomText16 sets ControlCustomText16 field to given value.

### HasControlCustomText16

`func (o *ControlsDataPutPreviousValues) HasControlCustomText16() bool`

HasControlCustomText16 returns a boolean if a field has been set.

### GetControlCustomText17

`func (o *ControlsDataPutPreviousValues) GetControlCustomText17() string`

GetControlCustomText17 returns the ControlCustomText17 field if non-nil, zero value otherwise.

### GetControlCustomText17Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText17Ok() (*string, bool)`

GetControlCustomText17Ok returns a tuple with the ControlCustomText17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText17

`func (o *ControlsDataPutPreviousValues) SetControlCustomText17(v string)`

SetControlCustomText17 sets ControlCustomText17 field to given value.

### HasControlCustomText17

`func (o *ControlsDataPutPreviousValues) HasControlCustomText17() bool`

HasControlCustomText17 returns a boolean if a field has been set.

### GetControlCustomText18

`func (o *ControlsDataPutPreviousValues) GetControlCustomText18() string`

GetControlCustomText18 returns the ControlCustomText18 field if non-nil, zero value otherwise.

### GetControlCustomText18Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText18Ok() (*string, bool)`

GetControlCustomText18Ok returns a tuple with the ControlCustomText18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText18

`func (o *ControlsDataPutPreviousValues) SetControlCustomText18(v string)`

SetControlCustomText18 sets ControlCustomText18 field to given value.

### HasControlCustomText18

`func (o *ControlsDataPutPreviousValues) HasControlCustomText18() bool`

HasControlCustomText18 returns a boolean if a field has been set.

### GetControlCustomText19

`func (o *ControlsDataPutPreviousValues) GetControlCustomText19() string`

GetControlCustomText19 returns the ControlCustomText19 field if non-nil, zero value otherwise.

### GetControlCustomText19Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText19Ok() (*string, bool)`

GetControlCustomText19Ok returns a tuple with the ControlCustomText19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText19

`func (o *ControlsDataPutPreviousValues) SetControlCustomText19(v string)`

SetControlCustomText19 sets ControlCustomText19 field to given value.

### HasControlCustomText19

`func (o *ControlsDataPutPreviousValues) HasControlCustomText19() bool`

HasControlCustomText19 returns a boolean if a field has been set.

### GetControlCustomText20

`func (o *ControlsDataPutPreviousValues) GetControlCustomText20() string`

GetControlCustomText20 returns the ControlCustomText20 field if non-nil, zero value otherwise.

### GetControlCustomText20Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText20Ok() (*string, bool)`

GetControlCustomText20Ok returns a tuple with the ControlCustomText20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText20

`func (o *ControlsDataPutPreviousValues) SetControlCustomText20(v string)`

SetControlCustomText20 sets ControlCustomText20 field to given value.

### HasControlCustomText20

`func (o *ControlsDataPutPreviousValues) HasControlCustomText20() bool`

HasControlCustomText20 returns a boolean if a field has been set.

### GetControlCustomText21

`func (o *ControlsDataPutPreviousValues) GetControlCustomText21() string`

GetControlCustomText21 returns the ControlCustomText21 field if non-nil, zero value otherwise.

### GetControlCustomText21Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText21Ok() (*string, bool)`

GetControlCustomText21Ok returns a tuple with the ControlCustomText21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText21

`func (o *ControlsDataPutPreviousValues) SetControlCustomText21(v string)`

SetControlCustomText21 sets ControlCustomText21 field to given value.

### HasControlCustomText21

`func (o *ControlsDataPutPreviousValues) HasControlCustomText21() bool`

HasControlCustomText21 returns a boolean if a field has been set.

### GetControlCustomText22

`func (o *ControlsDataPutPreviousValues) GetControlCustomText22() string`

GetControlCustomText22 returns the ControlCustomText22 field if non-nil, zero value otherwise.

### GetControlCustomText22Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText22Ok() (*string, bool)`

GetControlCustomText22Ok returns a tuple with the ControlCustomText22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText22

`func (o *ControlsDataPutPreviousValues) SetControlCustomText22(v string)`

SetControlCustomText22 sets ControlCustomText22 field to given value.

### HasControlCustomText22

`func (o *ControlsDataPutPreviousValues) HasControlCustomText22() bool`

HasControlCustomText22 returns a boolean if a field has been set.

### GetControlCustomText23

`func (o *ControlsDataPutPreviousValues) GetControlCustomText23() string`

GetControlCustomText23 returns the ControlCustomText23 field if non-nil, zero value otherwise.

### GetControlCustomText23Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText23Ok() (*string, bool)`

GetControlCustomText23Ok returns a tuple with the ControlCustomText23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText23

`func (o *ControlsDataPutPreviousValues) SetControlCustomText23(v string)`

SetControlCustomText23 sets ControlCustomText23 field to given value.

### HasControlCustomText23

`func (o *ControlsDataPutPreviousValues) HasControlCustomText23() bool`

HasControlCustomText23 returns a boolean if a field has been set.

### GetControlCustomText24

`func (o *ControlsDataPutPreviousValues) GetControlCustomText24() string`

GetControlCustomText24 returns the ControlCustomText24 field if non-nil, zero value otherwise.

### GetControlCustomText24Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText24Ok() (*string, bool)`

GetControlCustomText24Ok returns a tuple with the ControlCustomText24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText24

`func (o *ControlsDataPutPreviousValues) SetControlCustomText24(v string)`

SetControlCustomText24 sets ControlCustomText24 field to given value.

### HasControlCustomText24

`func (o *ControlsDataPutPreviousValues) HasControlCustomText24() bool`

HasControlCustomText24 returns a boolean if a field has been set.

### GetCustomDate1

`func (o *ControlsDataPutPreviousValues) GetCustomDate1() string`

GetCustomDate1 returns the CustomDate1 field if non-nil, zero value otherwise.

### GetCustomDate1Ok

`func (o *ControlsDataPutPreviousValues) GetCustomDate1Ok() (*string, bool)`

GetCustomDate1Ok returns a tuple with the CustomDate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate1

`func (o *ControlsDataPutPreviousValues) SetCustomDate1(v string)`

SetCustomDate1 sets CustomDate1 field to given value.

### HasCustomDate1

`func (o *ControlsDataPutPreviousValues) HasCustomDate1() bool`

HasCustomDate1 returns a boolean if a field has been set.

### GetCustomDate2

`func (o *ControlsDataPutPreviousValues) GetCustomDate2() string`

GetCustomDate2 returns the CustomDate2 field if non-nil, zero value otherwise.

### GetCustomDate2Ok

`func (o *ControlsDataPutPreviousValues) GetCustomDate2Ok() (*string, bool)`

GetCustomDate2Ok returns a tuple with the CustomDate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate2

`func (o *ControlsDataPutPreviousValues) SetCustomDate2(v string)`

SetCustomDate2 sets CustomDate2 field to given value.

### HasCustomDate2

`func (o *ControlsDataPutPreviousValues) HasCustomDate2() bool`

HasCustomDate2 returns a boolean if a field has been set.

### GetCustomDate3

`func (o *ControlsDataPutPreviousValues) GetCustomDate3() string`

GetCustomDate3 returns the CustomDate3 field if non-nil, zero value otherwise.

### GetCustomDate3Ok

`func (o *ControlsDataPutPreviousValues) GetCustomDate3Ok() (*string, bool)`

GetCustomDate3Ok returns a tuple with the CustomDate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate3

`func (o *ControlsDataPutPreviousValues) SetCustomDate3(v string)`

SetCustomDate3 sets CustomDate3 field to given value.

### HasCustomDate3

`func (o *ControlsDataPutPreviousValues) HasCustomDate3() bool`

HasCustomDate3 returns a boolean if a field has been set.

### GetCustomDate4

`func (o *ControlsDataPutPreviousValues) GetCustomDate4() string`

GetCustomDate4 returns the CustomDate4 field if non-nil, zero value otherwise.

### GetCustomDate4Ok

`func (o *ControlsDataPutPreviousValues) GetCustomDate4Ok() (*string, bool)`

GetCustomDate4Ok returns a tuple with the CustomDate4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate4

`func (o *ControlsDataPutPreviousValues) SetCustomDate4(v string)`

SetCustomDate4 sets CustomDate4 field to given value.

### HasCustomDate4

`func (o *ControlsDataPutPreviousValues) HasCustomDate4() bool`

HasCustomDate4 returns a boolean if a field has been set.

### GetCustomDate5

`func (o *ControlsDataPutPreviousValues) GetCustomDate5() string`

GetCustomDate5 returns the CustomDate5 field if non-nil, zero value otherwise.

### GetCustomDate5Ok

`func (o *ControlsDataPutPreviousValues) GetCustomDate5Ok() (*string, bool)`

GetCustomDate5Ok returns a tuple with the CustomDate5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate5

`func (o *ControlsDataPutPreviousValues) SetCustomDate5(v string)`

SetCustomDate5 sets CustomDate5 field to given value.

### HasCustomDate5

`func (o *ControlsDataPutPreviousValues) HasCustomDate5() bool`

HasCustomDate5 returns a boolean if a field has been set.

### GetCustomUserSelect5UserId

`func (o *ControlsDataPutPreviousValues) GetCustomUserSelect5UserId() int32`

GetCustomUserSelect5UserId returns the CustomUserSelect5UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect5UserIdOk

`func (o *ControlsDataPutPreviousValues) GetCustomUserSelect5UserIdOk() (*int32, bool)`

GetCustomUserSelect5UserIdOk returns a tuple with the CustomUserSelect5UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect5UserId

`func (o *ControlsDataPutPreviousValues) SetCustomUserSelect5UserId(v int32)`

SetCustomUserSelect5UserId sets CustomUserSelect5UserId field to given value.

### HasCustomUserSelect5UserId

`func (o *ControlsDataPutPreviousValues) HasCustomUserSelect5UserId() bool`

HasCustomUserSelect5UserId returns a boolean if a field has been set.

### GetCustomUserSelect6UserId

`func (o *ControlsDataPutPreviousValues) GetCustomUserSelect6UserId() int32`

GetCustomUserSelect6UserId returns the CustomUserSelect6UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect6UserIdOk

`func (o *ControlsDataPutPreviousValues) GetCustomUserSelect6UserIdOk() (*int32, bool)`

GetCustomUserSelect6UserIdOk returns a tuple with the CustomUserSelect6UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect6UserId

`func (o *ControlsDataPutPreviousValues) SetCustomUserSelect6UserId(v int32)`

SetCustomUserSelect6UserId sets CustomUserSelect6UserId field to given value.

### HasCustomUserSelect6UserId

`func (o *ControlsDataPutPreviousValues) HasCustomUserSelect6UserId() bool`

HasCustomUserSelect6UserId returns a boolean if a field has been set.

### GetCustomUserSelect7UserId

`func (o *ControlsDataPutPreviousValues) GetCustomUserSelect7UserId() int32`

GetCustomUserSelect7UserId returns the CustomUserSelect7UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect7UserIdOk

`func (o *ControlsDataPutPreviousValues) GetCustomUserSelect7UserIdOk() (*int32, bool)`

GetCustomUserSelect7UserIdOk returns a tuple with the CustomUserSelect7UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect7UserId

`func (o *ControlsDataPutPreviousValues) SetCustomUserSelect7UserId(v int32)`

SetCustomUserSelect7UserId sets CustomUserSelect7UserId field to given value.

### HasCustomUserSelect7UserId

`func (o *ControlsDataPutPreviousValues) HasCustomUserSelect7UserId() bool`

HasCustomUserSelect7UserId returns a boolean if a field has been set.

### GetCustomUserSelect8UserId

`func (o *ControlsDataPutPreviousValues) GetCustomUserSelect8UserId() int32`

GetCustomUserSelect8UserId returns the CustomUserSelect8UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect8UserIdOk

`func (o *ControlsDataPutPreviousValues) GetCustomUserSelect8UserIdOk() (*int32, bool)`

GetCustomUserSelect8UserIdOk returns a tuple with the CustomUserSelect8UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect8UserId

`func (o *ControlsDataPutPreviousValues) SetCustomUserSelect8UserId(v int32)`

SetCustomUserSelect8UserId sets CustomUserSelect8UserId field to given value.

### HasCustomUserSelect8UserId

`func (o *ControlsDataPutPreviousValues) HasCustomUserSelect8UserId() bool`

HasCustomUserSelect8UserId returns a boolean if a field has been set.

### GetCustomUserSelect9UserId

`func (o *ControlsDataPutPreviousValues) GetCustomUserSelect9UserId() int32`

GetCustomUserSelect9UserId returns the CustomUserSelect9UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect9UserIdOk

`func (o *ControlsDataPutPreviousValues) GetCustomUserSelect9UserIdOk() (*int32, bool)`

GetCustomUserSelect9UserIdOk returns a tuple with the CustomUserSelect9UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect9UserId

`func (o *ControlsDataPutPreviousValues) SetCustomUserSelect9UserId(v int32)`

SetCustomUserSelect9UserId sets CustomUserSelect9UserId field to given value.

### HasCustomUserSelect9UserId

`func (o *ControlsDataPutPreviousValues) HasCustomUserSelect9UserId() bool`

HasCustomUserSelect9UserId returns a boolean if a field has been set.

### GetCustomUserSelect10UserId

`func (o *ControlsDataPutPreviousValues) GetCustomUserSelect10UserId() int32`

GetCustomUserSelect10UserId returns the CustomUserSelect10UserId field if non-nil, zero value otherwise.

### GetCustomUserSelect10UserIdOk

`func (o *ControlsDataPutPreviousValues) GetCustomUserSelect10UserIdOk() (*int32, bool)`

GetCustomUserSelect10UserIdOk returns a tuple with the CustomUserSelect10UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserSelect10UserId

`func (o *ControlsDataPutPreviousValues) SetCustomUserSelect10UserId(v int32)`

SetCustomUserSelect10UserId sets CustomUserSelect10UserId field to given value.

### HasCustomUserSelect10UserId

`func (o *ControlsDataPutPreviousValues) HasCustomUserSelect10UserId() bool`

HasCustomUserSelect10UserId returns a boolean if a field has been set.

### GetControlCustomText25

`func (o *ControlsDataPutPreviousValues) GetControlCustomText25() string`

GetControlCustomText25 returns the ControlCustomText25 field if non-nil, zero value otherwise.

### GetControlCustomText25Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText25Ok() (*string, bool)`

GetControlCustomText25Ok returns a tuple with the ControlCustomText25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText25

`func (o *ControlsDataPutPreviousValues) SetControlCustomText25(v string)`

SetControlCustomText25 sets ControlCustomText25 field to given value.

### HasControlCustomText25

`func (o *ControlsDataPutPreviousValues) HasControlCustomText25() bool`

HasControlCustomText25 returns a boolean if a field has been set.

### GetControlCustomText26

`func (o *ControlsDataPutPreviousValues) GetControlCustomText26() string`

GetControlCustomText26 returns the ControlCustomText26 field if non-nil, zero value otherwise.

### GetControlCustomText26Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText26Ok() (*string, bool)`

GetControlCustomText26Ok returns a tuple with the ControlCustomText26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText26

`func (o *ControlsDataPutPreviousValues) SetControlCustomText26(v string)`

SetControlCustomText26 sets ControlCustomText26 field to given value.

### HasControlCustomText26

`func (o *ControlsDataPutPreviousValues) HasControlCustomText26() bool`

HasControlCustomText26 returns a boolean if a field has been set.

### GetControlCustomText27

`func (o *ControlsDataPutPreviousValues) GetControlCustomText27() string`

GetControlCustomText27 returns the ControlCustomText27 field if non-nil, zero value otherwise.

### GetControlCustomText27Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText27Ok() (*string, bool)`

GetControlCustomText27Ok returns a tuple with the ControlCustomText27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText27

`func (o *ControlsDataPutPreviousValues) SetControlCustomText27(v string)`

SetControlCustomText27 sets ControlCustomText27 field to given value.

### HasControlCustomText27

`func (o *ControlsDataPutPreviousValues) HasControlCustomText27() bool`

HasControlCustomText27 returns a boolean if a field has been set.

### GetControlCustomText28

`func (o *ControlsDataPutPreviousValues) GetControlCustomText28() string`

GetControlCustomText28 returns the ControlCustomText28 field if non-nil, zero value otherwise.

### GetControlCustomText28Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText28Ok() (*string, bool)`

GetControlCustomText28Ok returns a tuple with the ControlCustomText28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText28

`func (o *ControlsDataPutPreviousValues) SetControlCustomText28(v string)`

SetControlCustomText28 sets ControlCustomText28 field to given value.

### HasControlCustomText28

`func (o *ControlsDataPutPreviousValues) HasControlCustomText28() bool`

HasControlCustomText28 returns a boolean if a field has been set.

### GetControlCustomText29

`func (o *ControlsDataPutPreviousValues) GetControlCustomText29() string`

GetControlCustomText29 returns the ControlCustomText29 field if non-nil, zero value otherwise.

### GetControlCustomText29Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText29Ok() (*string, bool)`

GetControlCustomText29Ok returns a tuple with the ControlCustomText29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText29

`func (o *ControlsDataPutPreviousValues) SetControlCustomText29(v string)`

SetControlCustomText29 sets ControlCustomText29 field to given value.

### HasControlCustomText29

`func (o *ControlsDataPutPreviousValues) HasControlCustomText29() bool`

HasControlCustomText29 returns a boolean if a field has been set.

### GetControlCustomText30

`func (o *ControlsDataPutPreviousValues) GetControlCustomText30() string`

GetControlCustomText30 returns the ControlCustomText30 field if non-nil, zero value otherwise.

### GetControlCustomText30Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText30Ok() (*string, bool)`

GetControlCustomText30Ok returns a tuple with the ControlCustomText30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText30

`func (o *ControlsDataPutPreviousValues) SetControlCustomText30(v string)`

SetControlCustomText30 sets ControlCustomText30 field to given value.

### HasControlCustomText30

`func (o *ControlsDataPutPreviousValues) HasControlCustomText30() bool`

HasControlCustomText30 returns a boolean if a field has been set.

### GetControlCustomText31

`func (o *ControlsDataPutPreviousValues) GetControlCustomText31() string`

GetControlCustomText31 returns the ControlCustomText31 field if non-nil, zero value otherwise.

### GetControlCustomText31Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText31Ok() (*string, bool)`

GetControlCustomText31Ok returns a tuple with the ControlCustomText31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText31

`func (o *ControlsDataPutPreviousValues) SetControlCustomText31(v string)`

SetControlCustomText31 sets ControlCustomText31 field to given value.

### HasControlCustomText31

`func (o *ControlsDataPutPreviousValues) HasControlCustomText31() bool`

HasControlCustomText31 returns a boolean if a field has been set.

### GetControlCustomText32

`func (o *ControlsDataPutPreviousValues) GetControlCustomText32() string`

GetControlCustomText32 returns the ControlCustomText32 field if non-nil, zero value otherwise.

### GetControlCustomText32Ok

`func (o *ControlsDataPutPreviousValues) GetControlCustomText32Ok() (*string, bool)`

GetControlCustomText32Ok returns a tuple with the ControlCustomText32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomText32

`func (o *ControlsDataPutPreviousValues) SetControlCustomText32(v string)`

SetControlCustomText32 sets ControlCustomText32 field to given value.

### HasControlCustomText32

`func (o *ControlsDataPutPreviousValues) HasControlCustomText32() bool`

HasControlCustomText32 returns a boolean if a field has been set.

### GetControlCustomSelect13OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect13OptionId() int32`

GetControlCustomSelect13OptionId returns the ControlCustomSelect13OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect13OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect13OptionIdOk() (*int32, bool)`

GetControlCustomSelect13OptionIdOk returns a tuple with the ControlCustomSelect13OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect13OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect13OptionId(v int32)`

SetControlCustomSelect13OptionId sets ControlCustomSelect13OptionId field to given value.

### HasControlCustomSelect13OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect13OptionId() bool`

HasControlCustomSelect13OptionId returns a boolean if a field has been set.

### GetControlCustomSelect14OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect14OptionId() int32`

GetControlCustomSelect14OptionId returns the ControlCustomSelect14OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect14OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect14OptionIdOk() (*int32, bool)`

GetControlCustomSelect14OptionIdOk returns a tuple with the ControlCustomSelect14OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect14OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect14OptionId(v int32)`

SetControlCustomSelect14OptionId sets ControlCustomSelect14OptionId field to given value.

### HasControlCustomSelect14OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect14OptionId() bool`

HasControlCustomSelect14OptionId returns a boolean if a field has been set.

### GetControlCustomSelect15OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect15OptionId() int32`

GetControlCustomSelect15OptionId returns the ControlCustomSelect15OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect15OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect15OptionIdOk() (*int32, bool)`

GetControlCustomSelect15OptionIdOk returns a tuple with the ControlCustomSelect15OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect15OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect15OptionId(v int32)`

SetControlCustomSelect15OptionId sets ControlCustomSelect15OptionId field to given value.

### HasControlCustomSelect15OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect15OptionId() bool`

HasControlCustomSelect15OptionId returns a boolean if a field has been set.

### GetControlCustomSelect16OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect16OptionId() int32`

GetControlCustomSelect16OptionId returns the ControlCustomSelect16OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect16OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect16OptionIdOk() (*int32, bool)`

GetControlCustomSelect16OptionIdOk returns a tuple with the ControlCustomSelect16OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect16OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect16OptionId(v int32)`

SetControlCustomSelect16OptionId sets ControlCustomSelect16OptionId field to given value.

### HasControlCustomSelect16OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect16OptionId() bool`

HasControlCustomSelect16OptionId returns a boolean if a field has been set.

### GetControlsSegmentId

`func (o *ControlsDataPutPreviousValues) GetControlsSegmentId() int32`

GetControlsSegmentId returns the ControlsSegmentId field if non-nil, zero value otherwise.

### GetControlsSegmentIdOk

`func (o *ControlsDataPutPreviousValues) GetControlsSegmentIdOk() (*int32, bool)`

GetControlsSegmentIdOk returns a tuple with the ControlsSegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlsSegmentId

`func (o *ControlsDataPutPreviousValues) SetControlsSegmentId(v int32)`

SetControlsSegmentId sets ControlsSegmentId field to given value.

### HasControlsSegmentId

`func (o *ControlsDataPutPreviousValues) HasControlsSegmentId() bool`

HasControlsSegmentId returns a boolean if a field has been set.

### GetControlCustomSelect17OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect17OptionId() int32`

GetControlCustomSelect17OptionId returns the ControlCustomSelect17OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect17OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect17OptionIdOk() (*int32, bool)`

GetControlCustomSelect17OptionIdOk returns a tuple with the ControlCustomSelect17OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect17OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect17OptionId(v int32)`

SetControlCustomSelect17OptionId sets ControlCustomSelect17OptionId field to given value.

### HasControlCustomSelect17OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect17OptionId() bool`

HasControlCustomSelect17OptionId returns a boolean if a field has been set.

### GetControlCustomSelect18OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect18OptionId() int32`

GetControlCustomSelect18OptionId returns the ControlCustomSelect18OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect18OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect18OptionIdOk() (*int32, bool)`

GetControlCustomSelect18OptionIdOk returns a tuple with the ControlCustomSelect18OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect18OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect18OptionId(v int32)`

SetControlCustomSelect18OptionId sets ControlCustomSelect18OptionId field to given value.

### HasControlCustomSelect18OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect18OptionId() bool`

HasControlCustomSelect18OptionId returns a boolean if a field has been set.

### GetControlCustomSelect19OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect19OptionId() int32`

GetControlCustomSelect19OptionId returns the ControlCustomSelect19OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect19OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect19OptionIdOk() (*int32, bool)`

GetControlCustomSelect19OptionIdOk returns a tuple with the ControlCustomSelect19OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect19OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect19OptionId(v int32)`

SetControlCustomSelect19OptionId sets ControlCustomSelect19OptionId field to given value.

### HasControlCustomSelect19OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect19OptionId() bool`

HasControlCustomSelect19OptionId returns a boolean if a field has been set.

### GetControlCustomSelect20OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect20OptionId() int32`

GetControlCustomSelect20OptionId returns the ControlCustomSelect20OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect20OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect20OptionIdOk() (*int32, bool)`

GetControlCustomSelect20OptionIdOk returns a tuple with the ControlCustomSelect20OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect20OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect20OptionId(v int32)`

SetControlCustomSelect20OptionId sets ControlCustomSelect20OptionId field to given value.

### HasControlCustomSelect20OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect20OptionId() bool`

HasControlCustomSelect20OptionId returns a boolean if a field has been set.

### GetControlCustomSelect21OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect21OptionId() int32`

GetControlCustomSelect21OptionId returns the ControlCustomSelect21OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect21OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect21OptionIdOk() (*int32, bool)`

GetControlCustomSelect21OptionIdOk returns a tuple with the ControlCustomSelect21OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect21OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect21OptionId(v int32)`

SetControlCustomSelect21OptionId sets ControlCustomSelect21OptionId field to given value.

### HasControlCustomSelect21OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect21OptionId() bool`

HasControlCustomSelect21OptionId returns a boolean if a field has been set.

### GetControlCustomSelect22OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect22OptionId() int32`

GetControlCustomSelect22OptionId returns the ControlCustomSelect22OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect22OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect22OptionIdOk() (*int32, bool)`

GetControlCustomSelect22OptionIdOk returns a tuple with the ControlCustomSelect22OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect22OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect22OptionId(v int32)`

SetControlCustomSelect22OptionId sets ControlCustomSelect22OptionId field to given value.

### HasControlCustomSelect22OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect22OptionId() bool`

HasControlCustomSelect22OptionId returns a boolean if a field has been set.

### GetControlCustomSelect23OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect23OptionId() int32`

GetControlCustomSelect23OptionId returns the ControlCustomSelect23OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect23OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect23OptionIdOk() (*int32, bool)`

GetControlCustomSelect23OptionIdOk returns a tuple with the ControlCustomSelect23OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect23OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect23OptionId(v int32)`

SetControlCustomSelect23OptionId sets ControlCustomSelect23OptionId field to given value.

### HasControlCustomSelect23OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect23OptionId() bool`

HasControlCustomSelect23OptionId returns a boolean if a field has been set.

### GetControlCustomSelect24OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect24OptionId() int32`

GetControlCustomSelect24OptionId returns the ControlCustomSelect24OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect24OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect24OptionIdOk() (*int32, bool)`

GetControlCustomSelect24OptionIdOk returns a tuple with the ControlCustomSelect24OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect24OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect24OptionId(v int32)`

SetControlCustomSelect24OptionId sets ControlCustomSelect24OptionId field to given value.

### HasControlCustomSelect24OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect24OptionId() bool`

HasControlCustomSelect24OptionId returns a boolean if a field has been set.

### GetControlCustomSelect25OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect25OptionId() int32`

GetControlCustomSelect25OptionId returns the ControlCustomSelect25OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect25OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect25OptionIdOk() (*int32, bool)`

GetControlCustomSelect25OptionIdOk returns a tuple with the ControlCustomSelect25OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect25OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect25OptionId(v int32)`

SetControlCustomSelect25OptionId sets ControlCustomSelect25OptionId field to given value.

### HasControlCustomSelect25OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect25OptionId() bool`

HasControlCustomSelect25OptionId returns a boolean if a field has been set.

### GetControlCustomSelect26OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect26OptionId() int32`

GetControlCustomSelect26OptionId returns the ControlCustomSelect26OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect26OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect26OptionIdOk() (*int32, bool)`

GetControlCustomSelect26OptionIdOk returns a tuple with the ControlCustomSelect26OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect26OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect26OptionId(v int32)`

SetControlCustomSelect26OptionId sets ControlCustomSelect26OptionId field to given value.

### HasControlCustomSelect26OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect26OptionId() bool`

HasControlCustomSelect26OptionId returns a boolean if a field has been set.

### GetControlCustomSelect27OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect27OptionId() int32`

GetControlCustomSelect27OptionId returns the ControlCustomSelect27OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect27OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect27OptionIdOk() (*int32, bool)`

GetControlCustomSelect27OptionIdOk returns a tuple with the ControlCustomSelect27OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect27OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect27OptionId(v int32)`

SetControlCustomSelect27OptionId sets ControlCustomSelect27OptionId field to given value.

### HasControlCustomSelect27OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect27OptionId() bool`

HasControlCustomSelect27OptionId returns a boolean if a field has been set.

### GetControlCustomSelect28OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect28OptionId() int32`

GetControlCustomSelect28OptionId returns the ControlCustomSelect28OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect28OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect28OptionIdOk() (*int32, bool)`

GetControlCustomSelect28OptionIdOk returns a tuple with the ControlCustomSelect28OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect28OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect28OptionId(v int32)`

SetControlCustomSelect28OptionId sets ControlCustomSelect28OptionId field to given value.

### HasControlCustomSelect28OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect28OptionId() bool`

HasControlCustomSelect28OptionId returns a boolean if a field has been set.

### GetControlCustomSelect29OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect29OptionId() int32`

GetControlCustomSelect29OptionId returns the ControlCustomSelect29OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect29OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect29OptionIdOk() (*int32, bool)`

GetControlCustomSelect29OptionIdOk returns a tuple with the ControlCustomSelect29OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect29OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect29OptionId(v int32)`

SetControlCustomSelect29OptionId sets ControlCustomSelect29OptionId field to given value.

### HasControlCustomSelect29OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect29OptionId() bool`

HasControlCustomSelect29OptionId returns a boolean if a field has been set.

### GetControlCustomSelect30OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect30OptionId() int32`

GetControlCustomSelect30OptionId returns the ControlCustomSelect30OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect30OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect30OptionIdOk() (*int32, bool)`

GetControlCustomSelect30OptionIdOk returns a tuple with the ControlCustomSelect30OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect30OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect30OptionId(v int32)`

SetControlCustomSelect30OptionId sets ControlCustomSelect30OptionId field to given value.

### HasControlCustomSelect30OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect30OptionId() bool`

HasControlCustomSelect30OptionId returns a boolean if a field has been set.

### GetControlCustomSelect31OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect31OptionId() int32`

GetControlCustomSelect31OptionId returns the ControlCustomSelect31OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect31OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect31OptionIdOk() (*int32, bool)`

GetControlCustomSelect31OptionIdOk returns a tuple with the ControlCustomSelect31OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect31OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect31OptionId(v int32)`

SetControlCustomSelect31OptionId sets ControlCustomSelect31OptionId field to given value.

### HasControlCustomSelect31OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect31OptionId() bool`

HasControlCustomSelect31OptionId returns a boolean if a field has been set.

### GetControlCustomSelect32OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect32OptionId() int32`

GetControlCustomSelect32OptionId returns the ControlCustomSelect32OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect32OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect32OptionIdOk() (*int32, bool)`

GetControlCustomSelect32OptionIdOk returns a tuple with the ControlCustomSelect32OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect32OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect32OptionId(v int32)`

SetControlCustomSelect32OptionId sets ControlCustomSelect32OptionId field to given value.

### HasControlCustomSelect32OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect32OptionId() bool`

HasControlCustomSelect32OptionId returns a boolean if a field has been set.

### GetControlCustomSelect33OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect33OptionId() int32`

GetControlCustomSelect33OptionId returns the ControlCustomSelect33OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect33OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect33OptionIdOk() (*int32, bool)`

GetControlCustomSelect33OptionIdOk returns a tuple with the ControlCustomSelect33OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect33OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect33OptionId(v int32)`

SetControlCustomSelect33OptionId sets ControlCustomSelect33OptionId field to given value.

### HasControlCustomSelect33OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect33OptionId() bool`

HasControlCustomSelect33OptionId returns a boolean if a field has been set.

### GetControlCustomSelect34OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect34OptionId() int32`

GetControlCustomSelect34OptionId returns the ControlCustomSelect34OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect34OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect34OptionIdOk() (*int32, bool)`

GetControlCustomSelect34OptionIdOk returns a tuple with the ControlCustomSelect34OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect34OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect34OptionId(v int32)`

SetControlCustomSelect34OptionId sets ControlCustomSelect34OptionId field to given value.

### HasControlCustomSelect34OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect34OptionId() bool`

HasControlCustomSelect34OptionId returns a boolean if a field has been set.

### GetControlCustomSelect35OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect35OptionId() int32`

GetControlCustomSelect35OptionId returns the ControlCustomSelect35OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect35OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect35OptionIdOk() (*int32, bool)`

GetControlCustomSelect35OptionIdOk returns a tuple with the ControlCustomSelect35OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect35OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect35OptionId(v int32)`

SetControlCustomSelect35OptionId sets ControlCustomSelect35OptionId field to given value.

### HasControlCustomSelect35OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect35OptionId() bool`

HasControlCustomSelect35OptionId returns a boolean if a field has been set.

### GetControlCustomSelect36OptionId

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect36OptionId() int32`

GetControlCustomSelect36OptionId returns the ControlCustomSelect36OptionId field if non-nil, zero value otherwise.

### GetControlCustomSelect36OptionIdOk

`func (o *ControlsDataPutPreviousValues) GetControlCustomSelect36OptionIdOk() (*int32, bool)`

GetControlCustomSelect36OptionIdOk returns a tuple with the ControlCustomSelect36OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCustomSelect36OptionId

`func (o *ControlsDataPutPreviousValues) SetControlCustomSelect36OptionId(v int32)`

SetControlCustomSelect36OptionId sets ControlCustomSelect36OptionId field to given value.

### HasControlCustomSelect36OptionId

`func (o *ControlsDataPutPreviousValues) HasControlCustomSelect36OptionId() bool`

HasControlCustomSelect36OptionId returns a boolean if a field has been set.

### GetAtRiskStatus

`func (o *ControlsDataPutPreviousValues) GetAtRiskStatus() string`

GetAtRiskStatus returns the AtRiskStatus field if non-nil, zero value otherwise.

### GetAtRiskStatusOk

`func (o *ControlsDataPutPreviousValues) GetAtRiskStatusOk() (*string, bool)`

GetAtRiskStatusOk returns a tuple with the AtRiskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtRiskStatus

`func (o *ControlsDataPutPreviousValues) SetAtRiskStatus(v string)`

SetAtRiskStatus sets AtRiskStatus field to given value.

### HasAtRiskStatus

`func (o *ControlsDataPutPreviousValues) HasAtRiskStatus() bool`

HasAtRiskStatus returns a boolean if a field has been set.

### GetLatestResultId

`func (o *ControlsDataPutPreviousValues) GetLatestResultId() int32`

GetLatestResultId returns the LatestResultId field if non-nil, zero value otherwise.

### GetLatestResultIdOk

`func (o *ControlsDataPutPreviousValues) GetLatestResultIdOk() (*int32, bool)`

GetLatestResultIdOk returns a tuple with the LatestResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestResultId

`func (o *ControlsDataPutPreviousValues) SetLatestResultId(v int32)`

SetLatestResultId sets LatestResultId field to given value.

### HasLatestResultId

`func (o *ControlsDataPutPreviousValues) HasLatestResultId() bool`

HasLatestResultId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


