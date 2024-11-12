# TestsPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**ControlsDatumId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;controls_data.id&#x60;.&lt;fk table&#x3D;&#39;controls_data&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SampleSize** | Pointer to **string** |  | [optional] 
**Selections** | Pointer to **string** |  | [optional] 
**TesterUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ReviewerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**Results** | Pointer to **string** |  | [optional] 
**TestSheet** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**EffectivenessOptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;effectiveness_options.id&#x60;.&lt;fk table&#x3D;&#39;effectiveness_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**StatusOptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;status_options.id&#x60;.&lt;fk table&#x3D;&#39;status_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Archived** | Pointer to **bool** |  | [optional] [default to false]
**ArchivedAt** | Pointer to **string** |  | [optional] 
**OldHours** | Pointer to **string** |  | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**OldBudgetHours** | Pointer to **string** |  | [optional] 
**OldPercentComplete** | Pointer to **string** |  | [optional] 
**TestPeriod** | Pointer to **string** |  | [optional] 
**EntityId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;entities.id&#x60;.&lt;fk table&#x3D;&#39;entities&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;controls.id&#x60;.&lt;fk table&#x3D;&#39;controls&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**PrecisionLevel** | Pointer to **string** |  | [optional] 
**MrcEvidence** | Pointer to **string** |  | [optional] 
**SheetId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;sheets.id&#x60;.&lt;fk table&#x3D;&#39;sheets&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IsInactive** | Pointer to **bool** |  | [optional] [default to false]
**CompleteDate** | Pointer to **string** |  | [optional] 
**ReviewedDate** | Pointer to **string** |  | [optional] 
**ReviewNotesDate** | Pointer to **string** |  | [optional] 
**CompleteByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ReviewedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ReviewNotesByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestSectionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_sections.id&#x60;.&lt;fk table&#x3D;&#39;test_sections&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Procedures** | Pointer to **string** |  | [optional] 
**SheetFileId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;files.id&#x60;.&lt;fk table&#x3D;&#39;files&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OldReviewerHours** | Pointer to **string** |  | [optional] 
**OldReviewerBudgetHours** | Pointer to **string** |  | [optional] 
**SecondaryReviewerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**PopulationSize** | Pointer to **string** |  | [optional] 
**TestPbcRequest** | Pointer to **string** |  | [optional] 
**TestPbcRequestAdditional** | Pointer to **string** |  | [optional] 
**PopulationCompleteness** | Pointer to **string** |  | [optional] 
**PopulationSource** | Pointer to **string** |  | [optional] 
**ExternalAuditorStatusOptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;external_auditor_status_options.id&#x60;.&lt;fk table&#x3D;&#39;external_auditor_status_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ReviewerDueDate** | Pointer to **string** |  | [optional] 
**OpenDate** | Pointer to **string** |  | [optional] 
**UnderReviewTDate** | Pointer to **string** |  | [optional] 
**UnderReviewR2Date** | Pointer to **string** |  | [optional] 
**OpenByUserId** | Pointer to **int32** |  | [optional] 
**UnderReviewTByUserId** | Pointer to **int32** |  | [optional] 
**UnderReviewR2ByUserId** | Pointer to **int32** |  | [optional] 
**TestCustomSelect1OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_custom_select1_options.id&#x60;.&lt;fk table&#x3D;&#39;test_custom_select1_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestCustomSelect2OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_custom_select2_options.id&#x60;.&lt;fk table&#x3D;&#39;test_custom_select2_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestCustomSelect3OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_custom_select3_options.id&#x60;.&lt;fk table&#x3D;&#39;test_custom_select3_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**InformationReliability** | Pointer to **string** |  | [optional] 
**ControlCharacteristics** | Pointer to **string** |  | [optional] 
**ExceptionNature** | Pointer to **string** |  | [optional] 
**SampleSizeAdditional** | Pointer to **string** |  | [optional] 
**NumberExceptions** | Pointer to **string** |  | [optional] 
**SecondaryReviewerDueDate** | Pointer to **string** |  | [optional] 
**ArchiveId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;archives.id&#x60;.&lt;fk table&#x3D;&#39;archives&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestCustomSelect4OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_custom_select4_options.id&#x60;.&lt;fk table&#x3D;&#39;test_custom_select4_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FirstCompleteDate** | Pointer to **string** |  | [optional] 
**FirstReviewedDate** | Pointer to **string** |  | [optional] 
**FirstUnderReviewTDate** | Pointer to **string** |  | [optional] 
**FirstUnderReviewR2Date** | Pointer to **string** |  | [optional] 
**CustomText1** | Pointer to **string** |  | [optional] 
**CustomText2** | Pointer to **string** |  | [optional] 
**CustomText3** | Pointer to **string** |  | [optional] 
**CustomText4** | Pointer to **string** |  | [optional] 
**CustomText5** | Pointer to **string** |  | [optional] 
**CustomText6** | Pointer to **string** |  | [optional] 
**CustomText7** | Pointer to **string** |  | [optional] 
**CustomText8** | Pointer to **string** |  | [optional] 
**HoldByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**HoldDate** | Pointer to **string** |  | [optional] 
**FirstHoldDate** | Pointer to **string** |  | [optional] 
**HoldStatusOptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;status_options.id&#x60;.&lt;fk table&#x3D;&#39;status_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestCustomSelect5OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_custom_select5_options.id&#x60;.&lt;fk table&#x3D;&#39;test_custom_select5_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestCustomSelect6OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_custom_select6_options.id&#x60;.&lt;fk table&#x3D;&#39;test_custom_select6_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestCustomSelect7OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_custom_select7_options.id&#x60;.&lt;fk table&#x3D;&#39;test_custom_select7_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**BudgetHours** | Pointer to **float32** |  | [optional] 
**Hours** | Pointer to **float32** |  | [optional] 
**ReviewerBudgetHours** | Pointer to **float32** |  | [optional] 
**ReviewerHours** | Pointer to **float32** |  | [optional] 
**PercentComplete** | Pointer to **float32** |  | [optional] 
**CustomText9** | Pointer to **string** |  | [optional] 
**CustomText10** | Pointer to **string** |  | [optional] 
**CustomText11** | Pointer to **string** |  | [optional] 
**CustomText12** | Pointer to **string** |  | [optional] 
**CustomText13** | Pointer to **string** |  | [optional] 
**CustomText14** | Pointer to **string** |  | [optional] 
**CustomText15** | Pointer to **string** |  | [optional] 
**CustomText16** | Pointer to **string** |  | [optional] 
**CustomText17** | Pointer to **string** |  | [optional] 
**CustomText18** | Pointer to **string** |  | [optional] 
**CustomText19** | Pointer to **string** |  | [optional] 
**CustomText20** | Pointer to **string** |  | [optional] 
**CustomText21** | Pointer to **string** |  | [optional] 
**CustomText22** | Pointer to **string** |  | [optional] 
**CustomText23** | Pointer to **string** |  | [optional] 
**CustomText24** | Pointer to **string** |  | [optional] 
**CustomText25** | Pointer to **string** |  | [optional] 
**CustomText26** | Pointer to **string** |  | [optional] 
**CustomText27** | Pointer to **string** |  | [optional] 
**CustomText28** | Pointer to **string** |  | [optional] 
**CustomText29** | Pointer to **string** |  | [optional] 
**CustomText30** | Pointer to **string** |  | [optional] 
**CustomText31** | Pointer to **string** |  | [optional] 
**CustomText32** | Pointer to **string** |  | [optional] 
**CustomText33** | Pointer to **string** |  | [optional] 
**TestCustomSelect8OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_custom_select8_options.id&#x60;.&lt;fk table&#x3D;&#39;test_custom_select8_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestCustomSelect9OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_custom_select9_options.id&#x60;.&lt;fk table&#x3D;&#39;test_custom_select9_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestCustomSelect10OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_custom_select10_options.id&#x60;.&lt;fk table&#x3D;&#39;test_custom_select10_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestCustomSelect11OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_custom_select11_options.id&#x60;.&lt;fk table&#x3D;&#39;test_custom_select11_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestCustomSelect12OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_custom_select12_options.id&#x60;.&lt;fk table&#x3D;&#39;test_custom_select12_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestCustomSelect13OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_custom_select13_options.id&#x60;.&lt;fk table&#x3D;&#39;test_custom_select13_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestCustomSelect14OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_custom_select14_options.id&#x60;.&lt;fk table&#x3D;&#39;test_custom_select14_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestCustomSelect15OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_custom_select15_options.id&#x60;.&lt;fk table&#x3D;&#39;test_custom_select15_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestCustomSelect16OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_custom_select16_options.id&#x60;.&lt;fk table&#x3D;&#39;test_custom_select16_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestCustomSelect17OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_custom_select17_options.id&#x60;.&lt;fk table&#x3D;&#39;test_custom_select17_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FieldData** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewTestsPutPreviousValues

`func NewTestsPutPreviousValues() *TestsPutPreviousValues`

NewTestsPutPreviousValues instantiates a new TestsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestsPutPreviousValuesWithDefaults

`func NewTestsPutPreviousValuesWithDefaults() *TestsPutPreviousValues`

NewTestsPutPreviousValuesWithDefaults instantiates a new TestsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TestsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TestsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TestsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TestsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TestsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TestsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TestsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TestsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TestsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *TestsPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *TestsPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *TestsPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *TestsPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetControlsDatumId

`func (o *TestsPutPreviousValues) GetControlsDatumId() int32`

GetControlsDatumId returns the ControlsDatumId field if non-nil, zero value otherwise.

### GetControlsDatumIdOk

`func (o *TestsPutPreviousValues) GetControlsDatumIdOk() (*int32, bool)`

GetControlsDatumIdOk returns a tuple with the ControlsDatumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlsDatumId

`func (o *TestsPutPreviousValues) SetControlsDatumId(v int32)`

SetControlsDatumId sets ControlsDatumId field to given value.

### HasControlsDatumId

`func (o *TestsPutPreviousValues) HasControlsDatumId() bool`

HasControlsDatumId returns a boolean if a field has been set.

### GetOrder

`func (o *TestsPutPreviousValues) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TestsPutPreviousValues) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TestsPutPreviousValues) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TestsPutPreviousValues) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetName

`func (o *TestsPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestsPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestsPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestsPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSampleSize

`func (o *TestsPutPreviousValues) GetSampleSize() string`

GetSampleSize returns the SampleSize field if non-nil, zero value otherwise.

### GetSampleSizeOk

`func (o *TestsPutPreviousValues) GetSampleSizeOk() (*string, bool)`

GetSampleSizeOk returns a tuple with the SampleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleSize

`func (o *TestsPutPreviousValues) SetSampleSize(v string)`

SetSampleSize sets SampleSize field to given value.

### HasSampleSize

`func (o *TestsPutPreviousValues) HasSampleSize() bool`

HasSampleSize returns a boolean if a field has been set.

### GetSelections

`func (o *TestsPutPreviousValues) GetSelections() string`

GetSelections returns the Selections field if non-nil, zero value otherwise.

### GetSelectionsOk

`func (o *TestsPutPreviousValues) GetSelectionsOk() (*string, bool)`

GetSelectionsOk returns a tuple with the Selections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelections

`func (o *TestsPutPreviousValues) SetSelections(v string)`

SetSelections sets Selections field to given value.

### HasSelections

`func (o *TestsPutPreviousValues) HasSelections() bool`

HasSelections returns a boolean if a field has been set.

### GetTesterUserId

`func (o *TestsPutPreviousValues) GetTesterUserId() int32`

GetTesterUserId returns the TesterUserId field if non-nil, zero value otherwise.

### GetTesterUserIdOk

`func (o *TestsPutPreviousValues) GetTesterUserIdOk() (*int32, bool)`

GetTesterUserIdOk returns a tuple with the TesterUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesterUserId

`func (o *TestsPutPreviousValues) SetTesterUserId(v int32)`

SetTesterUserId sets TesterUserId field to given value.

### HasTesterUserId

`func (o *TestsPutPreviousValues) HasTesterUserId() bool`

HasTesterUserId returns a boolean if a field has been set.

### GetReviewerUserId

`func (o *TestsPutPreviousValues) GetReviewerUserId() int32`

GetReviewerUserId returns the ReviewerUserId field if non-nil, zero value otherwise.

### GetReviewerUserIdOk

`func (o *TestsPutPreviousValues) GetReviewerUserIdOk() (*int32, bool)`

GetReviewerUserIdOk returns a tuple with the ReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerUserId

`func (o *TestsPutPreviousValues) SetReviewerUserId(v int32)`

SetReviewerUserId sets ReviewerUserId field to given value.

### HasReviewerUserId

`func (o *TestsPutPreviousValues) HasReviewerUserId() bool`

HasReviewerUserId returns a boolean if a field has been set.

### GetStartDate

`func (o *TestsPutPreviousValues) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TestsPutPreviousValues) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TestsPutPreviousValues) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TestsPutPreviousValues) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *TestsPutPreviousValues) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TestsPutPreviousValues) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TestsPutPreviousValues) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TestsPutPreviousValues) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetResults

`func (o *TestsPutPreviousValues) GetResults() string`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TestsPutPreviousValues) GetResultsOk() (*string, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TestsPutPreviousValues) SetResults(v string)`

SetResults sets Results field to given value.

### HasResults

`func (o *TestsPutPreviousValues) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTestSheet

`func (o *TestsPutPreviousValues) GetTestSheet() string`

GetTestSheet returns the TestSheet field if non-nil, zero value otherwise.

### GetTestSheetOk

`func (o *TestsPutPreviousValues) GetTestSheetOk() (*string, bool)`

GetTestSheetOk returns a tuple with the TestSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSheet

`func (o *TestsPutPreviousValues) SetTestSheet(v string)`

SetTestSheet sets TestSheet field to given value.

### HasTestSheet

`func (o *TestsPutPreviousValues) HasTestSheet() bool`

HasTestSheet returns a boolean if a field has been set.

### GetNotes

`func (o *TestsPutPreviousValues) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TestsPutPreviousValues) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TestsPutPreviousValues) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TestsPutPreviousValues) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetEffectivenessOptionId

`func (o *TestsPutPreviousValues) GetEffectivenessOptionId() int32`

GetEffectivenessOptionId returns the EffectivenessOptionId field if non-nil, zero value otherwise.

### GetEffectivenessOptionIdOk

`func (o *TestsPutPreviousValues) GetEffectivenessOptionIdOk() (*int32, bool)`

GetEffectivenessOptionIdOk returns a tuple with the EffectivenessOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectivenessOptionId

`func (o *TestsPutPreviousValues) SetEffectivenessOptionId(v int32)`

SetEffectivenessOptionId sets EffectivenessOptionId field to given value.

### HasEffectivenessOptionId

`func (o *TestsPutPreviousValues) HasEffectivenessOptionId() bool`

HasEffectivenessOptionId returns a boolean if a field has been set.

### GetStatusOptionId

`func (o *TestsPutPreviousValues) GetStatusOptionId() int32`

GetStatusOptionId returns the StatusOptionId field if non-nil, zero value otherwise.

### GetStatusOptionIdOk

`func (o *TestsPutPreviousValues) GetStatusOptionIdOk() (*int32, bool)`

GetStatusOptionIdOk returns a tuple with the StatusOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusOptionId

`func (o *TestsPutPreviousValues) SetStatusOptionId(v int32)`

SetStatusOptionId sets StatusOptionId field to given value.

### HasStatusOptionId

`func (o *TestsPutPreviousValues) HasStatusOptionId() bool`

HasStatusOptionId returns a boolean if a field has been set.

### GetArchived

`func (o *TestsPutPreviousValues) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *TestsPutPreviousValues) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *TestsPutPreviousValues) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *TestsPutPreviousValues) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetArchivedAt

`func (o *TestsPutPreviousValues) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *TestsPutPreviousValues) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *TestsPutPreviousValues) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *TestsPutPreviousValues) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetOldHours

`func (o *TestsPutPreviousValues) GetOldHours() string`

GetOldHours returns the OldHours field if non-nil, zero value otherwise.

### GetOldHoursOk

`func (o *TestsPutPreviousValues) GetOldHoursOk() (*string, bool)`

GetOldHoursOk returns a tuple with the OldHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldHours

`func (o *TestsPutPreviousValues) SetOldHours(v string)`

SetOldHours sets OldHours field to given value.

### HasOldHours

`func (o *TestsPutPreviousValues) HasOldHours() bool`

HasOldHours returns a boolean if a field has been set.

### GetDueDate

`func (o *TestsPutPreviousValues) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *TestsPutPreviousValues) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *TestsPutPreviousValues) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *TestsPutPreviousValues) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetOldBudgetHours

`func (o *TestsPutPreviousValues) GetOldBudgetHours() string`

GetOldBudgetHours returns the OldBudgetHours field if non-nil, zero value otherwise.

### GetOldBudgetHoursOk

`func (o *TestsPutPreviousValues) GetOldBudgetHoursOk() (*string, bool)`

GetOldBudgetHoursOk returns a tuple with the OldBudgetHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldBudgetHours

`func (o *TestsPutPreviousValues) SetOldBudgetHours(v string)`

SetOldBudgetHours sets OldBudgetHours field to given value.

### HasOldBudgetHours

`func (o *TestsPutPreviousValues) HasOldBudgetHours() bool`

HasOldBudgetHours returns a boolean if a field has been set.

### GetOldPercentComplete

`func (o *TestsPutPreviousValues) GetOldPercentComplete() string`

GetOldPercentComplete returns the OldPercentComplete field if non-nil, zero value otherwise.

### GetOldPercentCompleteOk

`func (o *TestsPutPreviousValues) GetOldPercentCompleteOk() (*string, bool)`

GetOldPercentCompleteOk returns a tuple with the OldPercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPercentComplete

`func (o *TestsPutPreviousValues) SetOldPercentComplete(v string)`

SetOldPercentComplete sets OldPercentComplete field to given value.

### HasOldPercentComplete

`func (o *TestsPutPreviousValues) HasOldPercentComplete() bool`

HasOldPercentComplete returns a boolean if a field has been set.

### GetTestPeriod

`func (o *TestsPutPreviousValues) GetTestPeriod() string`

GetTestPeriod returns the TestPeriod field if non-nil, zero value otherwise.

### GetTestPeriodOk

`func (o *TestsPutPreviousValues) GetTestPeriodOk() (*string, bool)`

GetTestPeriodOk returns a tuple with the TestPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPeriod

`func (o *TestsPutPreviousValues) SetTestPeriod(v string)`

SetTestPeriod sets TestPeriod field to given value.

### HasTestPeriod

`func (o *TestsPutPreviousValues) HasTestPeriod() bool`

HasTestPeriod returns a boolean if a field has been set.

### GetEntityId

`func (o *TestsPutPreviousValues) GetEntityId() int32`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *TestsPutPreviousValues) GetEntityIdOk() (*int32, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *TestsPutPreviousValues) SetEntityId(v int32)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *TestsPutPreviousValues) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetControlId

`func (o *TestsPutPreviousValues) GetControlId() int32`

GetControlId returns the ControlId field if non-nil, zero value otherwise.

### GetControlIdOk

`func (o *TestsPutPreviousValues) GetControlIdOk() (*int32, bool)`

GetControlIdOk returns a tuple with the ControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlId

`func (o *TestsPutPreviousValues) SetControlId(v int32)`

SetControlId sets ControlId field to given value.

### HasControlId

`func (o *TestsPutPreviousValues) HasControlId() bool`

HasControlId returns a boolean if a field has been set.

### GetPrecisionLevel

`func (o *TestsPutPreviousValues) GetPrecisionLevel() string`

GetPrecisionLevel returns the PrecisionLevel field if non-nil, zero value otherwise.

### GetPrecisionLevelOk

`func (o *TestsPutPreviousValues) GetPrecisionLevelOk() (*string, bool)`

GetPrecisionLevelOk returns a tuple with the PrecisionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecisionLevel

`func (o *TestsPutPreviousValues) SetPrecisionLevel(v string)`

SetPrecisionLevel sets PrecisionLevel field to given value.

### HasPrecisionLevel

`func (o *TestsPutPreviousValues) HasPrecisionLevel() bool`

HasPrecisionLevel returns a boolean if a field has been set.

### GetMrcEvidence

`func (o *TestsPutPreviousValues) GetMrcEvidence() string`

GetMrcEvidence returns the MrcEvidence field if non-nil, zero value otherwise.

### GetMrcEvidenceOk

`func (o *TestsPutPreviousValues) GetMrcEvidenceOk() (*string, bool)`

GetMrcEvidenceOk returns a tuple with the MrcEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrcEvidence

`func (o *TestsPutPreviousValues) SetMrcEvidence(v string)`

SetMrcEvidence sets MrcEvidence field to given value.

### HasMrcEvidence

`func (o *TestsPutPreviousValues) HasMrcEvidence() bool`

HasMrcEvidence returns a boolean if a field has been set.

### GetSheetId

`func (o *TestsPutPreviousValues) GetSheetId() int32`

GetSheetId returns the SheetId field if non-nil, zero value otherwise.

### GetSheetIdOk

`func (o *TestsPutPreviousValues) GetSheetIdOk() (*int32, bool)`

GetSheetIdOk returns a tuple with the SheetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheetId

`func (o *TestsPutPreviousValues) SetSheetId(v int32)`

SetSheetId sets SheetId field to given value.

### HasSheetId

`func (o *TestsPutPreviousValues) HasSheetId() bool`

HasSheetId returns a boolean if a field has been set.

### GetIsInactive

`func (o *TestsPutPreviousValues) GetIsInactive() bool`

GetIsInactive returns the IsInactive field if non-nil, zero value otherwise.

### GetIsInactiveOk

`func (o *TestsPutPreviousValues) GetIsInactiveOk() (*bool, bool)`

GetIsInactiveOk returns a tuple with the IsInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInactive

`func (o *TestsPutPreviousValues) SetIsInactive(v bool)`

SetIsInactive sets IsInactive field to given value.

### HasIsInactive

`func (o *TestsPutPreviousValues) HasIsInactive() bool`

HasIsInactive returns a boolean if a field has been set.

### GetCompleteDate

`func (o *TestsPutPreviousValues) GetCompleteDate() string`

GetCompleteDate returns the CompleteDate field if non-nil, zero value otherwise.

### GetCompleteDateOk

`func (o *TestsPutPreviousValues) GetCompleteDateOk() (*string, bool)`

GetCompleteDateOk returns a tuple with the CompleteDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteDate

`func (o *TestsPutPreviousValues) SetCompleteDate(v string)`

SetCompleteDate sets CompleteDate field to given value.

### HasCompleteDate

`func (o *TestsPutPreviousValues) HasCompleteDate() bool`

HasCompleteDate returns a boolean if a field has been set.

### GetReviewedDate

`func (o *TestsPutPreviousValues) GetReviewedDate() string`

GetReviewedDate returns the ReviewedDate field if non-nil, zero value otherwise.

### GetReviewedDateOk

`func (o *TestsPutPreviousValues) GetReviewedDateOk() (*string, bool)`

GetReviewedDateOk returns a tuple with the ReviewedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedDate

`func (o *TestsPutPreviousValues) SetReviewedDate(v string)`

SetReviewedDate sets ReviewedDate field to given value.

### HasReviewedDate

`func (o *TestsPutPreviousValues) HasReviewedDate() bool`

HasReviewedDate returns a boolean if a field has been set.

### GetReviewNotesDate

`func (o *TestsPutPreviousValues) GetReviewNotesDate() string`

GetReviewNotesDate returns the ReviewNotesDate field if non-nil, zero value otherwise.

### GetReviewNotesDateOk

`func (o *TestsPutPreviousValues) GetReviewNotesDateOk() (*string, bool)`

GetReviewNotesDateOk returns a tuple with the ReviewNotesDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewNotesDate

`func (o *TestsPutPreviousValues) SetReviewNotesDate(v string)`

SetReviewNotesDate sets ReviewNotesDate field to given value.

### HasReviewNotesDate

`func (o *TestsPutPreviousValues) HasReviewNotesDate() bool`

HasReviewNotesDate returns a boolean if a field has been set.

### GetCompleteByUserId

`func (o *TestsPutPreviousValues) GetCompleteByUserId() int32`

GetCompleteByUserId returns the CompleteByUserId field if non-nil, zero value otherwise.

### GetCompleteByUserIdOk

`func (o *TestsPutPreviousValues) GetCompleteByUserIdOk() (*int32, bool)`

GetCompleteByUserIdOk returns a tuple with the CompleteByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteByUserId

`func (o *TestsPutPreviousValues) SetCompleteByUserId(v int32)`

SetCompleteByUserId sets CompleteByUserId field to given value.

### HasCompleteByUserId

`func (o *TestsPutPreviousValues) HasCompleteByUserId() bool`

HasCompleteByUserId returns a boolean if a field has been set.

### GetReviewedByUserId

`func (o *TestsPutPreviousValues) GetReviewedByUserId() int32`

GetReviewedByUserId returns the ReviewedByUserId field if non-nil, zero value otherwise.

### GetReviewedByUserIdOk

`func (o *TestsPutPreviousValues) GetReviewedByUserIdOk() (*int32, bool)`

GetReviewedByUserIdOk returns a tuple with the ReviewedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedByUserId

`func (o *TestsPutPreviousValues) SetReviewedByUserId(v int32)`

SetReviewedByUserId sets ReviewedByUserId field to given value.

### HasReviewedByUserId

`func (o *TestsPutPreviousValues) HasReviewedByUserId() bool`

HasReviewedByUserId returns a boolean if a field has been set.

### GetReviewNotesByUserId

`func (o *TestsPutPreviousValues) GetReviewNotesByUserId() int32`

GetReviewNotesByUserId returns the ReviewNotesByUserId field if non-nil, zero value otherwise.

### GetReviewNotesByUserIdOk

`func (o *TestsPutPreviousValues) GetReviewNotesByUserIdOk() (*int32, bool)`

GetReviewNotesByUserIdOk returns a tuple with the ReviewNotesByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewNotesByUserId

`func (o *TestsPutPreviousValues) SetReviewNotesByUserId(v int32)`

SetReviewNotesByUserId sets ReviewNotesByUserId field to given value.

### HasReviewNotesByUserId

`func (o *TestsPutPreviousValues) HasReviewNotesByUserId() bool`

HasReviewNotesByUserId returns a boolean if a field has been set.

### GetTestSectionId

`func (o *TestsPutPreviousValues) GetTestSectionId() int32`

GetTestSectionId returns the TestSectionId field if non-nil, zero value otherwise.

### GetTestSectionIdOk

`func (o *TestsPutPreviousValues) GetTestSectionIdOk() (*int32, bool)`

GetTestSectionIdOk returns a tuple with the TestSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSectionId

`func (o *TestsPutPreviousValues) SetTestSectionId(v int32)`

SetTestSectionId sets TestSectionId field to given value.

### HasTestSectionId

`func (o *TestsPutPreviousValues) HasTestSectionId() bool`

HasTestSectionId returns a boolean if a field has been set.

### GetProcedures

`func (o *TestsPutPreviousValues) GetProcedures() string`

GetProcedures returns the Procedures field if non-nil, zero value otherwise.

### GetProceduresOk

`func (o *TestsPutPreviousValues) GetProceduresOk() (*string, bool)`

GetProceduresOk returns a tuple with the Procedures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcedures

`func (o *TestsPutPreviousValues) SetProcedures(v string)`

SetProcedures sets Procedures field to given value.

### HasProcedures

`func (o *TestsPutPreviousValues) HasProcedures() bool`

HasProcedures returns a boolean if a field has been set.

### GetSheetFileId

`func (o *TestsPutPreviousValues) GetSheetFileId() int32`

GetSheetFileId returns the SheetFileId field if non-nil, zero value otherwise.

### GetSheetFileIdOk

`func (o *TestsPutPreviousValues) GetSheetFileIdOk() (*int32, bool)`

GetSheetFileIdOk returns a tuple with the SheetFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheetFileId

`func (o *TestsPutPreviousValues) SetSheetFileId(v int32)`

SetSheetFileId sets SheetFileId field to given value.

### HasSheetFileId

`func (o *TestsPutPreviousValues) HasSheetFileId() bool`

HasSheetFileId returns a boolean if a field has been set.

### GetOldReviewerHours

`func (o *TestsPutPreviousValues) GetOldReviewerHours() string`

GetOldReviewerHours returns the OldReviewerHours field if non-nil, zero value otherwise.

### GetOldReviewerHoursOk

`func (o *TestsPutPreviousValues) GetOldReviewerHoursOk() (*string, bool)`

GetOldReviewerHoursOk returns a tuple with the OldReviewerHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldReviewerHours

`func (o *TestsPutPreviousValues) SetOldReviewerHours(v string)`

SetOldReviewerHours sets OldReviewerHours field to given value.

### HasOldReviewerHours

`func (o *TestsPutPreviousValues) HasOldReviewerHours() bool`

HasOldReviewerHours returns a boolean if a field has been set.

### GetOldReviewerBudgetHours

`func (o *TestsPutPreviousValues) GetOldReviewerBudgetHours() string`

GetOldReviewerBudgetHours returns the OldReviewerBudgetHours field if non-nil, zero value otherwise.

### GetOldReviewerBudgetHoursOk

`func (o *TestsPutPreviousValues) GetOldReviewerBudgetHoursOk() (*string, bool)`

GetOldReviewerBudgetHoursOk returns a tuple with the OldReviewerBudgetHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldReviewerBudgetHours

`func (o *TestsPutPreviousValues) SetOldReviewerBudgetHours(v string)`

SetOldReviewerBudgetHours sets OldReviewerBudgetHours field to given value.

### HasOldReviewerBudgetHours

`func (o *TestsPutPreviousValues) HasOldReviewerBudgetHours() bool`

HasOldReviewerBudgetHours returns a boolean if a field has been set.

### GetSecondaryReviewerUserId

`func (o *TestsPutPreviousValues) GetSecondaryReviewerUserId() int32`

GetSecondaryReviewerUserId returns the SecondaryReviewerUserId field if non-nil, zero value otherwise.

### GetSecondaryReviewerUserIdOk

`func (o *TestsPutPreviousValues) GetSecondaryReviewerUserIdOk() (*int32, bool)`

GetSecondaryReviewerUserIdOk returns a tuple with the SecondaryReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryReviewerUserId

`func (o *TestsPutPreviousValues) SetSecondaryReviewerUserId(v int32)`

SetSecondaryReviewerUserId sets SecondaryReviewerUserId field to given value.

### HasSecondaryReviewerUserId

`func (o *TestsPutPreviousValues) HasSecondaryReviewerUserId() bool`

HasSecondaryReviewerUserId returns a boolean if a field has been set.

### GetPopulationSize

`func (o *TestsPutPreviousValues) GetPopulationSize() string`

GetPopulationSize returns the PopulationSize field if non-nil, zero value otherwise.

### GetPopulationSizeOk

`func (o *TestsPutPreviousValues) GetPopulationSizeOk() (*string, bool)`

GetPopulationSizeOk returns a tuple with the PopulationSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationSize

`func (o *TestsPutPreviousValues) SetPopulationSize(v string)`

SetPopulationSize sets PopulationSize field to given value.

### HasPopulationSize

`func (o *TestsPutPreviousValues) HasPopulationSize() bool`

HasPopulationSize returns a boolean if a field has been set.

### GetTestPbcRequest

`func (o *TestsPutPreviousValues) GetTestPbcRequest() string`

GetTestPbcRequest returns the TestPbcRequest field if non-nil, zero value otherwise.

### GetTestPbcRequestOk

`func (o *TestsPutPreviousValues) GetTestPbcRequestOk() (*string, bool)`

GetTestPbcRequestOk returns a tuple with the TestPbcRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPbcRequest

`func (o *TestsPutPreviousValues) SetTestPbcRequest(v string)`

SetTestPbcRequest sets TestPbcRequest field to given value.

### HasTestPbcRequest

`func (o *TestsPutPreviousValues) HasTestPbcRequest() bool`

HasTestPbcRequest returns a boolean if a field has been set.

### GetTestPbcRequestAdditional

`func (o *TestsPutPreviousValues) GetTestPbcRequestAdditional() string`

GetTestPbcRequestAdditional returns the TestPbcRequestAdditional field if non-nil, zero value otherwise.

### GetTestPbcRequestAdditionalOk

`func (o *TestsPutPreviousValues) GetTestPbcRequestAdditionalOk() (*string, bool)`

GetTestPbcRequestAdditionalOk returns a tuple with the TestPbcRequestAdditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPbcRequestAdditional

`func (o *TestsPutPreviousValues) SetTestPbcRequestAdditional(v string)`

SetTestPbcRequestAdditional sets TestPbcRequestAdditional field to given value.

### HasTestPbcRequestAdditional

`func (o *TestsPutPreviousValues) HasTestPbcRequestAdditional() bool`

HasTestPbcRequestAdditional returns a boolean if a field has been set.

### GetPopulationCompleteness

`func (o *TestsPutPreviousValues) GetPopulationCompleteness() string`

GetPopulationCompleteness returns the PopulationCompleteness field if non-nil, zero value otherwise.

### GetPopulationCompletenessOk

`func (o *TestsPutPreviousValues) GetPopulationCompletenessOk() (*string, bool)`

GetPopulationCompletenessOk returns a tuple with the PopulationCompleteness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationCompleteness

`func (o *TestsPutPreviousValues) SetPopulationCompleteness(v string)`

SetPopulationCompleteness sets PopulationCompleteness field to given value.

### HasPopulationCompleteness

`func (o *TestsPutPreviousValues) HasPopulationCompleteness() bool`

HasPopulationCompleteness returns a boolean if a field has been set.

### GetPopulationSource

`func (o *TestsPutPreviousValues) GetPopulationSource() string`

GetPopulationSource returns the PopulationSource field if non-nil, zero value otherwise.

### GetPopulationSourceOk

`func (o *TestsPutPreviousValues) GetPopulationSourceOk() (*string, bool)`

GetPopulationSourceOk returns a tuple with the PopulationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationSource

`func (o *TestsPutPreviousValues) SetPopulationSource(v string)`

SetPopulationSource sets PopulationSource field to given value.

### HasPopulationSource

`func (o *TestsPutPreviousValues) HasPopulationSource() bool`

HasPopulationSource returns a boolean if a field has been set.

### GetExternalAuditorStatusOptionId

`func (o *TestsPutPreviousValues) GetExternalAuditorStatusOptionId() int32`

GetExternalAuditorStatusOptionId returns the ExternalAuditorStatusOptionId field if non-nil, zero value otherwise.

### GetExternalAuditorStatusOptionIdOk

`func (o *TestsPutPreviousValues) GetExternalAuditorStatusOptionIdOk() (*int32, bool)`

GetExternalAuditorStatusOptionIdOk returns a tuple with the ExternalAuditorStatusOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAuditorStatusOptionId

`func (o *TestsPutPreviousValues) SetExternalAuditorStatusOptionId(v int32)`

SetExternalAuditorStatusOptionId sets ExternalAuditorStatusOptionId field to given value.

### HasExternalAuditorStatusOptionId

`func (o *TestsPutPreviousValues) HasExternalAuditorStatusOptionId() bool`

HasExternalAuditorStatusOptionId returns a boolean if a field has been set.

### GetReviewerDueDate

`func (o *TestsPutPreviousValues) GetReviewerDueDate() string`

GetReviewerDueDate returns the ReviewerDueDate field if non-nil, zero value otherwise.

### GetReviewerDueDateOk

`func (o *TestsPutPreviousValues) GetReviewerDueDateOk() (*string, bool)`

GetReviewerDueDateOk returns a tuple with the ReviewerDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerDueDate

`func (o *TestsPutPreviousValues) SetReviewerDueDate(v string)`

SetReviewerDueDate sets ReviewerDueDate field to given value.

### HasReviewerDueDate

`func (o *TestsPutPreviousValues) HasReviewerDueDate() bool`

HasReviewerDueDate returns a boolean if a field has been set.

### GetOpenDate

`func (o *TestsPutPreviousValues) GetOpenDate() string`

GetOpenDate returns the OpenDate field if non-nil, zero value otherwise.

### GetOpenDateOk

`func (o *TestsPutPreviousValues) GetOpenDateOk() (*string, bool)`

GetOpenDateOk returns a tuple with the OpenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenDate

`func (o *TestsPutPreviousValues) SetOpenDate(v string)`

SetOpenDate sets OpenDate field to given value.

### HasOpenDate

`func (o *TestsPutPreviousValues) HasOpenDate() bool`

HasOpenDate returns a boolean if a field has been set.

### GetUnderReviewTDate

`func (o *TestsPutPreviousValues) GetUnderReviewTDate() string`

GetUnderReviewTDate returns the UnderReviewTDate field if non-nil, zero value otherwise.

### GetUnderReviewTDateOk

`func (o *TestsPutPreviousValues) GetUnderReviewTDateOk() (*string, bool)`

GetUnderReviewTDateOk returns a tuple with the UnderReviewTDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderReviewTDate

`func (o *TestsPutPreviousValues) SetUnderReviewTDate(v string)`

SetUnderReviewTDate sets UnderReviewTDate field to given value.

### HasUnderReviewTDate

`func (o *TestsPutPreviousValues) HasUnderReviewTDate() bool`

HasUnderReviewTDate returns a boolean if a field has been set.

### GetUnderReviewR2Date

`func (o *TestsPutPreviousValues) GetUnderReviewR2Date() string`

GetUnderReviewR2Date returns the UnderReviewR2Date field if non-nil, zero value otherwise.

### GetUnderReviewR2DateOk

`func (o *TestsPutPreviousValues) GetUnderReviewR2DateOk() (*string, bool)`

GetUnderReviewR2DateOk returns a tuple with the UnderReviewR2Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderReviewR2Date

`func (o *TestsPutPreviousValues) SetUnderReviewR2Date(v string)`

SetUnderReviewR2Date sets UnderReviewR2Date field to given value.

### HasUnderReviewR2Date

`func (o *TestsPutPreviousValues) HasUnderReviewR2Date() bool`

HasUnderReviewR2Date returns a boolean if a field has been set.

### GetOpenByUserId

`func (o *TestsPutPreviousValues) GetOpenByUserId() int32`

GetOpenByUserId returns the OpenByUserId field if non-nil, zero value otherwise.

### GetOpenByUserIdOk

`func (o *TestsPutPreviousValues) GetOpenByUserIdOk() (*int32, bool)`

GetOpenByUserIdOk returns a tuple with the OpenByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenByUserId

`func (o *TestsPutPreviousValues) SetOpenByUserId(v int32)`

SetOpenByUserId sets OpenByUserId field to given value.

### HasOpenByUserId

`func (o *TestsPutPreviousValues) HasOpenByUserId() bool`

HasOpenByUserId returns a boolean if a field has been set.

### GetUnderReviewTByUserId

`func (o *TestsPutPreviousValues) GetUnderReviewTByUserId() int32`

GetUnderReviewTByUserId returns the UnderReviewTByUserId field if non-nil, zero value otherwise.

### GetUnderReviewTByUserIdOk

`func (o *TestsPutPreviousValues) GetUnderReviewTByUserIdOk() (*int32, bool)`

GetUnderReviewTByUserIdOk returns a tuple with the UnderReviewTByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderReviewTByUserId

`func (o *TestsPutPreviousValues) SetUnderReviewTByUserId(v int32)`

SetUnderReviewTByUserId sets UnderReviewTByUserId field to given value.

### HasUnderReviewTByUserId

`func (o *TestsPutPreviousValues) HasUnderReviewTByUserId() bool`

HasUnderReviewTByUserId returns a boolean if a field has been set.

### GetUnderReviewR2ByUserId

`func (o *TestsPutPreviousValues) GetUnderReviewR2ByUserId() int32`

GetUnderReviewR2ByUserId returns the UnderReviewR2ByUserId field if non-nil, zero value otherwise.

### GetUnderReviewR2ByUserIdOk

`func (o *TestsPutPreviousValues) GetUnderReviewR2ByUserIdOk() (*int32, bool)`

GetUnderReviewR2ByUserIdOk returns a tuple with the UnderReviewR2ByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderReviewR2ByUserId

`func (o *TestsPutPreviousValues) SetUnderReviewR2ByUserId(v int32)`

SetUnderReviewR2ByUserId sets UnderReviewR2ByUserId field to given value.

### HasUnderReviewR2ByUserId

`func (o *TestsPutPreviousValues) HasUnderReviewR2ByUserId() bool`

HasUnderReviewR2ByUserId returns a boolean if a field has been set.

### GetTestCustomSelect1OptionId

`func (o *TestsPutPreviousValues) GetTestCustomSelect1OptionId() int32`

GetTestCustomSelect1OptionId returns the TestCustomSelect1OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect1OptionIdOk

`func (o *TestsPutPreviousValues) GetTestCustomSelect1OptionIdOk() (*int32, bool)`

GetTestCustomSelect1OptionIdOk returns a tuple with the TestCustomSelect1OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect1OptionId

`func (o *TestsPutPreviousValues) SetTestCustomSelect1OptionId(v int32)`

SetTestCustomSelect1OptionId sets TestCustomSelect1OptionId field to given value.

### HasTestCustomSelect1OptionId

`func (o *TestsPutPreviousValues) HasTestCustomSelect1OptionId() bool`

HasTestCustomSelect1OptionId returns a boolean if a field has been set.

### GetTestCustomSelect2OptionId

`func (o *TestsPutPreviousValues) GetTestCustomSelect2OptionId() int32`

GetTestCustomSelect2OptionId returns the TestCustomSelect2OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect2OptionIdOk

`func (o *TestsPutPreviousValues) GetTestCustomSelect2OptionIdOk() (*int32, bool)`

GetTestCustomSelect2OptionIdOk returns a tuple with the TestCustomSelect2OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect2OptionId

`func (o *TestsPutPreviousValues) SetTestCustomSelect2OptionId(v int32)`

SetTestCustomSelect2OptionId sets TestCustomSelect2OptionId field to given value.

### HasTestCustomSelect2OptionId

`func (o *TestsPutPreviousValues) HasTestCustomSelect2OptionId() bool`

HasTestCustomSelect2OptionId returns a boolean if a field has been set.

### GetTestCustomSelect3OptionId

`func (o *TestsPutPreviousValues) GetTestCustomSelect3OptionId() int32`

GetTestCustomSelect3OptionId returns the TestCustomSelect3OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect3OptionIdOk

`func (o *TestsPutPreviousValues) GetTestCustomSelect3OptionIdOk() (*int32, bool)`

GetTestCustomSelect3OptionIdOk returns a tuple with the TestCustomSelect3OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect3OptionId

`func (o *TestsPutPreviousValues) SetTestCustomSelect3OptionId(v int32)`

SetTestCustomSelect3OptionId sets TestCustomSelect3OptionId field to given value.

### HasTestCustomSelect3OptionId

`func (o *TestsPutPreviousValues) HasTestCustomSelect3OptionId() bool`

HasTestCustomSelect3OptionId returns a boolean if a field has been set.

### GetInformationReliability

`func (o *TestsPutPreviousValues) GetInformationReliability() string`

GetInformationReliability returns the InformationReliability field if non-nil, zero value otherwise.

### GetInformationReliabilityOk

`func (o *TestsPutPreviousValues) GetInformationReliabilityOk() (*string, bool)`

GetInformationReliabilityOk returns a tuple with the InformationReliability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformationReliability

`func (o *TestsPutPreviousValues) SetInformationReliability(v string)`

SetInformationReliability sets InformationReliability field to given value.

### HasInformationReliability

`func (o *TestsPutPreviousValues) HasInformationReliability() bool`

HasInformationReliability returns a boolean if a field has been set.

### GetControlCharacteristics

`func (o *TestsPutPreviousValues) GetControlCharacteristics() string`

GetControlCharacteristics returns the ControlCharacteristics field if non-nil, zero value otherwise.

### GetControlCharacteristicsOk

`func (o *TestsPutPreviousValues) GetControlCharacteristicsOk() (*string, bool)`

GetControlCharacteristicsOk returns a tuple with the ControlCharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCharacteristics

`func (o *TestsPutPreviousValues) SetControlCharacteristics(v string)`

SetControlCharacteristics sets ControlCharacteristics field to given value.

### HasControlCharacteristics

`func (o *TestsPutPreviousValues) HasControlCharacteristics() bool`

HasControlCharacteristics returns a boolean if a field has been set.

### GetExceptionNature

`func (o *TestsPutPreviousValues) GetExceptionNature() string`

GetExceptionNature returns the ExceptionNature field if non-nil, zero value otherwise.

### GetExceptionNatureOk

`func (o *TestsPutPreviousValues) GetExceptionNatureOk() (*string, bool)`

GetExceptionNatureOk returns a tuple with the ExceptionNature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionNature

`func (o *TestsPutPreviousValues) SetExceptionNature(v string)`

SetExceptionNature sets ExceptionNature field to given value.

### HasExceptionNature

`func (o *TestsPutPreviousValues) HasExceptionNature() bool`

HasExceptionNature returns a boolean if a field has been set.

### GetSampleSizeAdditional

`func (o *TestsPutPreviousValues) GetSampleSizeAdditional() string`

GetSampleSizeAdditional returns the SampleSizeAdditional field if non-nil, zero value otherwise.

### GetSampleSizeAdditionalOk

`func (o *TestsPutPreviousValues) GetSampleSizeAdditionalOk() (*string, bool)`

GetSampleSizeAdditionalOk returns a tuple with the SampleSizeAdditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleSizeAdditional

`func (o *TestsPutPreviousValues) SetSampleSizeAdditional(v string)`

SetSampleSizeAdditional sets SampleSizeAdditional field to given value.

### HasSampleSizeAdditional

`func (o *TestsPutPreviousValues) HasSampleSizeAdditional() bool`

HasSampleSizeAdditional returns a boolean if a field has been set.

### GetNumberExceptions

`func (o *TestsPutPreviousValues) GetNumberExceptions() string`

GetNumberExceptions returns the NumberExceptions field if non-nil, zero value otherwise.

### GetNumberExceptionsOk

`func (o *TestsPutPreviousValues) GetNumberExceptionsOk() (*string, bool)`

GetNumberExceptionsOk returns a tuple with the NumberExceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberExceptions

`func (o *TestsPutPreviousValues) SetNumberExceptions(v string)`

SetNumberExceptions sets NumberExceptions field to given value.

### HasNumberExceptions

`func (o *TestsPutPreviousValues) HasNumberExceptions() bool`

HasNumberExceptions returns a boolean if a field has been set.

### GetSecondaryReviewerDueDate

`func (o *TestsPutPreviousValues) GetSecondaryReviewerDueDate() string`

GetSecondaryReviewerDueDate returns the SecondaryReviewerDueDate field if non-nil, zero value otherwise.

### GetSecondaryReviewerDueDateOk

`func (o *TestsPutPreviousValues) GetSecondaryReviewerDueDateOk() (*string, bool)`

GetSecondaryReviewerDueDateOk returns a tuple with the SecondaryReviewerDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryReviewerDueDate

`func (o *TestsPutPreviousValues) SetSecondaryReviewerDueDate(v string)`

SetSecondaryReviewerDueDate sets SecondaryReviewerDueDate field to given value.

### HasSecondaryReviewerDueDate

`func (o *TestsPutPreviousValues) HasSecondaryReviewerDueDate() bool`

HasSecondaryReviewerDueDate returns a boolean if a field has been set.

### GetArchiveId

`func (o *TestsPutPreviousValues) GetArchiveId() int32`

GetArchiveId returns the ArchiveId field if non-nil, zero value otherwise.

### GetArchiveIdOk

`func (o *TestsPutPreviousValues) GetArchiveIdOk() (*int32, bool)`

GetArchiveIdOk returns a tuple with the ArchiveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveId

`func (o *TestsPutPreviousValues) SetArchiveId(v int32)`

SetArchiveId sets ArchiveId field to given value.

### HasArchiveId

`func (o *TestsPutPreviousValues) HasArchiveId() bool`

HasArchiveId returns a boolean if a field has been set.

### GetTestCustomSelect4OptionId

`func (o *TestsPutPreviousValues) GetTestCustomSelect4OptionId() int32`

GetTestCustomSelect4OptionId returns the TestCustomSelect4OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect4OptionIdOk

`func (o *TestsPutPreviousValues) GetTestCustomSelect4OptionIdOk() (*int32, bool)`

GetTestCustomSelect4OptionIdOk returns a tuple with the TestCustomSelect4OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect4OptionId

`func (o *TestsPutPreviousValues) SetTestCustomSelect4OptionId(v int32)`

SetTestCustomSelect4OptionId sets TestCustomSelect4OptionId field to given value.

### HasTestCustomSelect4OptionId

`func (o *TestsPutPreviousValues) HasTestCustomSelect4OptionId() bool`

HasTestCustomSelect4OptionId returns a boolean if a field has been set.

### GetFirstCompleteDate

`func (o *TestsPutPreviousValues) GetFirstCompleteDate() string`

GetFirstCompleteDate returns the FirstCompleteDate field if non-nil, zero value otherwise.

### GetFirstCompleteDateOk

`func (o *TestsPutPreviousValues) GetFirstCompleteDateOk() (*string, bool)`

GetFirstCompleteDateOk returns a tuple with the FirstCompleteDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCompleteDate

`func (o *TestsPutPreviousValues) SetFirstCompleteDate(v string)`

SetFirstCompleteDate sets FirstCompleteDate field to given value.

### HasFirstCompleteDate

`func (o *TestsPutPreviousValues) HasFirstCompleteDate() bool`

HasFirstCompleteDate returns a boolean if a field has been set.

### GetFirstReviewedDate

`func (o *TestsPutPreviousValues) GetFirstReviewedDate() string`

GetFirstReviewedDate returns the FirstReviewedDate field if non-nil, zero value otherwise.

### GetFirstReviewedDateOk

`func (o *TestsPutPreviousValues) GetFirstReviewedDateOk() (*string, bool)`

GetFirstReviewedDateOk returns a tuple with the FirstReviewedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstReviewedDate

`func (o *TestsPutPreviousValues) SetFirstReviewedDate(v string)`

SetFirstReviewedDate sets FirstReviewedDate field to given value.

### HasFirstReviewedDate

`func (o *TestsPutPreviousValues) HasFirstReviewedDate() bool`

HasFirstReviewedDate returns a boolean if a field has been set.

### GetFirstUnderReviewTDate

`func (o *TestsPutPreviousValues) GetFirstUnderReviewTDate() string`

GetFirstUnderReviewTDate returns the FirstUnderReviewTDate field if non-nil, zero value otherwise.

### GetFirstUnderReviewTDateOk

`func (o *TestsPutPreviousValues) GetFirstUnderReviewTDateOk() (*string, bool)`

GetFirstUnderReviewTDateOk returns a tuple with the FirstUnderReviewTDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstUnderReviewTDate

`func (o *TestsPutPreviousValues) SetFirstUnderReviewTDate(v string)`

SetFirstUnderReviewTDate sets FirstUnderReviewTDate field to given value.

### HasFirstUnderReviewTDate

`func (o *TestsPutPreviousValues) HasFirstUnderReviewTDate() bool`

HasFirstUnderReviewTDate returns a boolean if a field has been set.

### GetFirstUnderReviewR2Date

`func (o *TestsPutPreviousValues) GetFirstUnderReviewR2Date() string`

GetFirstUnderReviewR2Date returns the FirstUnderReviewR2Date field if non-nil, zero value otherwise.

### GetFirstUnderReviewR2DateOk

`func (o *TestsPutPreviousValues) GetFirstUnderReviewR2DateOk() (*string, bool)`

GetFirstUnderReviewR2DateOk returns a tuple with the FirstUnderReviewR2Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstUnderReviewR2Date

`func (o *TestsPutPreviousValues) SetFirstUnderReviewR2Date(v string)`

SetFirstUnderReviewR2Date sets FirstUnderReviewR2Date field to given value.

### HasFirstUnderReviewR2Date

`func (o *TestsPutPreviousValues) HasFirstUnderReviewR2Date() bool`

HasFirstUnderReviewR2Date returns a boolean if a field has been set.

### GetCustomText1

`func (o *TestsPutPreviousValues) GetCustomText1() string`

GetCustomText1 returns the CustomText1 field if non-nil, zero value otherwise.

### GetCustomText1Ok

`func (o *TestsPutPreviousValues) GetCustomText1Ok() (*string, bool)`

GetCustomText1Ok returns a tuple with the CustomText1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText1

`func (o *TestsPutPreviousValues) SetCustomText1(v string)`

SetCustomText1 sets CustomText1 field to given value.

### HasCustomText1

`func (o *TestsPutPreviousValues) HasCustomText1() bool`

HasCustomText1 returns a boolean if a field has been set.

### GetCustomText2

`func (o *TestsPutPreviousValues) GetCustomText2() string`

GetCustomText2 returns the CustomText2 field if non-nil, zero value otherwise.

### GetCustomText2Ok

`func (o *TestsPutPreviousValues) GetCustomText2Ok() (*string, bool)`

GetCustomText2Ok returns a tuple with the CustomText2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText2

`func (o *TestsPutPreviousValues) SetCustomText2(v string)`

SetCustomText2 sets CustomText2 field to given value.

### HasCustomText2

`func (o *TestsPutPreviousValues) HasCustomText2() bool`

HasCustomText2 returns a boolean if a field has been set.

### GetCustomText3

`func (o *TestsPutPreviousValues) GetCustomText3() string`

GetCustomText3 returns the CustomText3 field if non-nil, zero value otherwise.

### GetCustomText3Ok

`func (o *TestsPutPreviousValues) GetCustomText3Ok() (*string, bool)`

GetCustomText3Ok returns a tuple with the CustomText3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText3

`func (o *TestsPutPreviousValues) SetCustomText3(v string)`

SetCustomText3 sets CustomText3 field to given value.

### HasCustomText3

`func (o *TestsPutPreviousValues) HasCustomText3() bool`

HasCustomText3 returns a boolean if a field has been set.

### GetCustomText4

`func (o *TestsPutPreviousValues) GetCustomText4() string`

GetCustomText4 returns the CustomText4 field if non-nil, zero value otherwise.

### GetCustomText4Ok

`func (o *TestsPutPreviousValues) GetCustomText4Ok() (*string, bool)`

GetCustomText4Ok returns a tuple with the CustomText4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText4

`func (o *TestsPutPreviousValues) SetCustomText4(v string)`

SetCustomText4 sets CustomText4 field to given value.

### HasCustomText4

`func (o *TestsPutPreviousValues) HasCustomText4() bool`

HasCustomText4 returns a boolean if a field has been set.

### GetCustomText5

`func (o *TestsPutPreviousValues) GetCustomText5() string`

GetCustomText5 returns the CustomText5 field if non-nil, zero value otherwise.

### GetCustomText5Ok

`func (o *TestsPutPreviousValues) GetCustomText5Ok() (*string, bool)`

GetCustomText5Ok returns a tuple with the CustomText5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText5

`func (o *TestsPutPreviousValues) SetCustomText5(v string)`

SetCustomText5 sets CustomText5 field to given value.

### HasCustomText5

`func (o *TestsPutPreviousValues) HasCustomText5() bool`

HasCustomText5 returns a boolean if a field has been set.

### GetCustomText6

`func (o *TestsPutPreviousValues) GetCustomText6() string`

GetCustomText6 returns the CustomText6 field if non-nil, zero value otherwise.

### GetCustomText6Ok

`func (o *TestsPutPreviousValues) GetCustomText6Ok() (*string, bool)`

GetCustomText6Ok returns a tuple with the CustomText6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText6

`func (o *TestsPutPreviousValues) SetCustomText6(v string)`

SetCustomText6 sets CustomText6 field to given value.

### HasCustomText6

`func (o *TestsPutPreviousValues) HasCustomText6() bool`

HasCustomText6 returns a boolean if a field has been set.

### GetCustomText7

`func (o *TestsPutPreviousValues) GetCustomText7() string`

GetCustomText7 returns the CustomText7 field if non-nil, zero value otherwise.

### GetCustomText7Ok

`func (o *TestsPutPreviousValues) GetCustomText7Ok() (*string, bool)`

GetCustomText7Ok returns a tuple with the CustomText7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText7

`func (o *TestsPutPreviousValues) SetCustomText7(v string)`

SetCustomText7 sets CustomText7 field to given value.

### HasCustomText7

`func (o *TestsPutPreviousValues) HasCustomText7() bool`

HasCustomText7 returns a boolean if a field has been set.

### GetCustomText8

`func (o *TestsPutPreviousValues) GetCustomText8() string`

GetCustomText8 returns the CustomText8 field if non-nil, zero value otherwise.

### GetCustomText8Ok

`func (o *TestsPutPreviousValues) GetCustomText8Ok() (*string, bool)`

GetCustomText8Ok returns a tuple with the CustomText8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText8

`func (o *TestsPutPreviousValues) SetCustomText8(v string)`

SetCustomText8 sets CustomText8 field to given value.

### HasCustomText8

`func (o *TestsPutPreviousValues) HasCustomText8() bool`

HasCustomText8 returns a boolean if a field has been set.

### GetHoldByUserId

`func (o *TestsPutPreviousValues) GetHoldByUserId() int32`

GetHoldByUserId returns the HoldByUserId field if non-nil, zero value otherwise.

### GetHoldByUserIdOk

`func (o *TestsPutPreviousValues) GetHoldByUserIdOk() (*int32, bool)`

GetHoldByUserIdOk returns a tuple with the HoldByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldByUserId

`func (o *TestsPutPreviousValues) SetHoldByUserId(v int32)`

SetHoldByUserId sets HoldByUserId field to given value.

### HasHoldByUserId

`func (o *TestsPutPreviousValues) HasHoldByUserId() bool`

HasHoldByUserId returns a boolean if a field has been set.

### GetHoldDate

`func (o *TestsPutPreviousValues) GetHoldDate() string`

GetHoldDate returns the HoldDate field if non-nil, zero value otherwise.

### GetHoldDateOk

`func (o *TestsPutPreviousValues) GetHoldDateOk() (*string, bool)`

GetHoldDateOk returns a tuple with the HoldDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldDate

`func (o *TestsPutPreviousValues) SetHoldDate(v string)`

SetHoldDate sets HoldDate field to given value.

### HasHoldDate

`func (o *TestsPutPreviousValues) HasHoldDate() bool`

HasHoldDate returns a boolean if a field has been set.

### GetFirstHoldDate

`func (o *TestsPutPreviousValues) GetFirstHoldDate() string`

GetFirstHoldDate returns the FirstHoldDate field if non-nil, zero value otherwise.

### GetFirstHoldDateOk

`func (o *TestsPutPreviousValues) GetFirstHoldDateOk() (*string, bool)`

GetFirstHoldDateOk returns a tuple with the FirstHoldDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstHoldDate

`func (o *TestsPutPreviousValues) SetFirstHoldDate(v string)`

SetFirstHoldDate sets FirstHoldDate field to given value.

### HasFirstHoldDate

`func (o *TestsPutPreviousValues) HasFirstHoldDate() bool`

HasFirstHoldDate returns a boolean if a field has been set.

### GetHoldStatusOptionId

`func (o *TestsPutPreviousValues) GetHoldStatusOptionId() int32`

GetHoldStatusOptionId returns the HoldStatusOptionId field if non-nil, zero value otherwise.

### GetHoldStatusOptionIdOk

`func (o *TestsPutPreviousValues) GetHoldStatusOptionIdOk() (*int32, bool)`

GetHoldStatusOptionIdOk returns a tuple with the HoldStatusOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldStatusOptionId

`func (o *TestsPutPreviousValues) SetHoldStatusOptionId(v int32)`

SetHoldStatusOptionId sets HoldStatusOptionId field to given value.

### HasHoldStatusOptionId

`func (o *TestsPutPreviousValues) HasHoldStatusOptionId() bool`

HasHoldStatusOptionId returns a boolean if a field has been set.

### GetTestCustomSelect5OptionId

`func (o *TestsPutPreviousValues) GetTestCustomSelect5OptionId() int32`

GetTestCustomSelect5OptionId returns the TestCustomSelect5OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect5OptionIdOk

`func (o *TestsPutPreviousValues) GetTestCustomSelect5OptionIdOk() (*int32, bool)`

GetTestCustomSelect5OptionIdOk returns a tuple with the TestCustomSelect5OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect5OptionId

`func (o *TestsPutPreviousValues) SetTestCustomSelect5OptionId(v int32)`

SetTestCustomSelect5OptionId sets TestCustomSelect5OptionId field to given value.

### HasTestCustomSelect5OptionId

`func (o *TestsPutPreviousValues) HasTestCustomSelect5OptionId() bool`

HasTestCustomSelect5OptionId returns a boolean if a field has been set.

### GetTestCustomSelect6OptionId

`func (o *TestsPutPreviousValues) GetTestCustomSelect6OptionId() int32`

GetTestCustomSelect6OptionId returns the TestCustomSelect6OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect6OptionIdOk

`func (o *TestsPutPreviousValues) GetTestCustomSelect6OptionIdOk() (*int32, bool)`

GetTestCustomSelect6OptionIdOk returns a tuple with the TestCustomSelect6OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect6OptionId

`func (o *TestsPutPreviousValues) SetTestCustomSelect6OptionId(v int32)`

SetTestCustomSelect6OptionId sets TestCustomSelect6OptionId field to given value.

### HasTestCustomSelect6OptionId

`func (o *TestsPutPreviousValues) HasTestCustomSelect6OptionId() bool`

HasTestCustomSelect6OptionId returns a boolean if a field has been set.

### GetTestCustomSelect7OptionId

`func (o *TestsPutPreviousValues) GetTestCustomSelect7OptionId() int32`

GetTestCustomSelect7OptionId returns the TestCustomSelect7OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect7OptionIdOk

`func (o *TestsPutPreviousValues) GetTestCustomSelect7OptionIdOk() (*int32, bool)`

GetTestCustomSelect7OptionIdOk returns a tuple with the TestCustomSelect7OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect7OptionId

`func (o *TestsPutPreviousValues) SetTestCustomSelect7OptionId(v int32)`

SetTestCustomSelect7OptionId sets TestCustomSelect7OptionId field to given value.

### HasTestCustomSelect7OptionId

`func (o *TestsPutPreviousValues) HasTestCustomSelect7OptionId() bool`

HasTestCustomSelect7OptionId returns a boolean if a field has been set.

### GetBudgetHours

`func (o *TestsPutPreviousValues) GetBudgetHours() float32`

GetBudgetHours returns the BudgetHours field if non-nil, zero value otherwise.

### GetBudgetHoursOk

`func (o *TestsPutPreviousValues) GetBudgetHoursOk() (*float32, bool)`

GetBudgetHoursOk returns a tuple with the BudgetHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetHours

`func (o *TestsPutPreviousValues) SetBudgetHours(v float32)`

SetBudgetHours sets BudgetHours field to given value.

### HasBudgetHours

`func (o *TestsPutPreviousValues) HasBudgetHours() bool`

HasBudgetHours returns a boolean if a field has been set.

### GetHours

`func (o *TestsPutPreviousValues) GetHours() float32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *TestsPutPreviousValues) GetHoursOk() (*float32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *TestsPutPreviousValues) SetHours(v float32)`

SetHours sets Hours field to given value.

### HasHours

`func (o *TestsPutPreviousValues) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetReviewerBudgetHours

`func (o *TestsPutPreviousValues) GetReviewerBudgetHours() float32`

GetReviewerBudgetHours returns the ReviewerBudgetHours field if non-nil, zero value otherwise.

### GetReviewerBudgetHoursOk

`func (o *TestsPutPreviousValues) GetReviewerBudgetHoursOk() (*float32, bool)`

GetReviewerBudgetHoursOk returns a tuple with the ReviewerBudgetHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerBudgetHours

`func (o *TestsPutPreviousValues) SetReviewerBudgetHours(v float32)`

SetReviewerBudgetHours sets ReviewerBudgetHours field to given value.

### HasReviewerBudgetHours

`func (o *TestsPutPreviousValues) HasReviewerBudgetHours() bool`

HasReviewerBudgetHours returns a boolean if a field has been set.

### GetReviewerHours

`func (o *TestsPutPreviousValues) GetReviewerHours() float32`

GetReviewerHours returns the ReviewerHours field if non-nil, zero value otherwise.

### GetReviewerHoursOk

`func (o *TestsPutPreviousValues) GetReviewerHoursOk() (*float32, bool)`

GetReviewerHoursOk returns a tuple with the ReviewerHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerHours

`func (o *TestsPutPreviousValues) SetReviewerHours(v float32)`

SetReviewerHours sets ReviewerHours field to given value.

### HasReviewerHours

`func (o *TestsPutPreviousValues) HasReviewerHours() bool`

HasReviewerHours returns a boolean if a field has been set.

### GetPercentComplete

`func (o *TestsPutPreviousValues) GetPercentComplete() float32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *TestsPutPreviousValues) GetPercentCompleteOk() (*float32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *TestsPutPreviousValues) SetPercentComplete(v float32)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *TestsPutPreviousValues) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### GetCustomText9

`func (o *TestsPutPreviousValues) GetCustomText9() string`

GetCustomText9 returns the CustomText9 field if non-nil, zero value otherwise.

### GetCustomText9Ok

`func (o *TestsPutPreviousValues) GetCustomText9Ok() (*string, bool)`

GetCustomText9Ok returns a tuple with the CustomText9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText9

`func (o *TestsPutPreviousValues) SetCustomText9(v string)`

SetCustomText9 sets CustomText9 field to given value.

### HasCustomText9

`func (o *TestsPutPreviousValues) HasCustomText9() bool`

HasCustomText9 returns a boolean if a field has been set.

### GetCustomText10

`func (o *TestsPutPreviousValues) GetCustomText10() string`

GetCustomText10 returns the CustomText10 field if non-nil, zero value otherwise.

### GetCustomText10Ok

`func (o *TestsPutPreviousValues) GetCustomText10Ok() (*string, bool)`

GetCustomText10Ok returns a tuple with the CustomText10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText10

`func (o *TestsPutPreviousValues) SetCustomText10(v string)`

SetCustomText10 sets CustomText10 field to given value.

### HasCustomText10

`func (o *TestsPutPreviousValues) HasCustomText10() bool`

HasCustomText10 returns a boolean if a field has been set.

### GetCustomText11

`func (o *TestsPutPreviousValues) GetCustomText11() string`

GetCustomText11 returns the CustomText11 field if non-nil, zero value otherwise.

### GetCustomText11Ok

`func (o *TestsPutPreviousValues) GetCustomText11Ok() (*string, bool)`

GetCustomText11Ok returns a tuple with the CustomText11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText11

`func (o *TestsPutPreviousValues) SetCustomText11(v string)`

SetCustomText11 sets CustomText11 field to given value.

### HasCustomText11

`func (o *TestsPutPreviousValues) HasCustomText11() bool`

HasCustomText11 returns a boolean if a field has been set.

### GetCustomText12

`func (o *TestsPutPreviousValues) GetCustomText12() string`

GetCustomText12 returns the CustomText12 field if non-nil, zero value otherwise.

### GetCustomText12Ok

`func (o *TestsPutPreviousValues) GetCustomText12Ok() (*string, bool)`

GetCustomText12Ok returns a tuple with the CustomText12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText12

`func (o *TestsPutPreviousValues) SetCustomText12(v string)`

SetCustomText12 sets CustomText12 field to given value.

### HasCustomText12

`func (o *TestsPutPreviousValues) HasCustomText12() bool`

HasCustomText12 returns a boolean if a field has been set.

### GetCustomText13

`func (o *TestsPutPreviousValues) GetCustomText13() string`

GetCustomText13 returns the CustomText13 field if non-nil, zero value otherwise.

### GetCustomText13Ok

`func (o *TestsPutPreviousValues) GetCustomText13Ok() (*string, bool)`

GetCustomText13Ok returns a tuple with the CustomText13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText13

`func (o *TestsPutPreviousValues) SetCustomText13(v string)`

SetCustomText13 sets CustomText13 field to given value.

### HasCustomText13

`func (o *TestsPutPreviousValues) HasCustomText13() bool`

HasCustomText13 returns a boolean if a field has been set.

### GetCustomText14

`func (o *TestsPutPreviousValues) GetCustomText14() string`

GetCustomText14 returns the CustomText14 field if non-nil, zero value otherwise.

### GetCustomText14Ok

`func (o *TestsPutPreviousValues) GetCustomText14Ok() (*string, bool)`

GetCustomText14Ok returns a tuple with the CustomText14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText14

`func (o *TestsPutPreviousValues) SetCustomText14(v string)`

SetCustomText14 sets CustomText14 field to given value.

### HasCustomText14

`func (o *TestsPutPreviousValues) HasCustomText14() bool`

HasCustomText14 returns a boolean if a field has been set.

### GetCustomText15

`func (o *TestsPutPreviousValues) GetCustomText15() string`

GetCustomText15 returns the CustomText15 field if non-nil, zero value otherwise.

### GetCustomText15Ok

`func (o *TestsPutPreviousValues) GetCustomText15Ok() (*string, bool)`

GetCustomText15Ok returns a tuple with the CustomText15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText15

`func (o *TestsPutPreviousValues) SetCustomText15(v string)`

SetCustomText15 sets CustomText15 field to given value.

### HasCustomText15

`func (o *TestsPutPreviousValues) HasCustomText15() bool`

HasCustomText15 returns a boolean if a field has been set.

### GetCustomText16

`func (o *TestsPutPreviousValues) GetCustomText16() string`

GetCustomText16 returns the CustomText16 field if non-nil, zero value otherwise.

### GetCustomText16Ok

`func (o *TestsPutPreviousValues) GetCustomText16Ok() (*string, bool)`

GetCustomText16Ok returns a tuple with the CustomText16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText16

`func (o *TestsPutPreviousValues) SetCustomText16(v string)`

SetCustomText16 sets CustomText16 field to given value.

### HasCustomText16

`func (o *TestsPutPreviousValues) HasCustomText16() bool`

HasCustomText16 returns a boolean if a field has been set.

### GetCustomText17

`func (o *TestsPutPreviousValues) GetCustomText17() string`

GetCustomText17 returns the CustomText17 field if non-nil, zero value otherwise.

### GetCustomText17Ok

`func (o *TestsPutPreviousValues) GetCustomText17Ok() (*string, bool)`

GetCustomText17Ok returns a tuple with the CustomText17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText17

`func (o *TestsPutPreviousValues) SetCustomText17(v string)`

SetCustomText17 sets CustomText17 field to given value.

### HasCustomText17

`func (o *TestsPutPreviousValues) HasCustomText17() bool`

HasCustomText17 returns a boolean if a field has been set.

### GetCustomText18

`func (o *TestsPutPreviousValues) GetCustomText18() string`

GetCustomText18 returns the CustomText18 field if non-nil, zero value otherwise.

### GetCustomText18Ok

`func (o *TestsPutPreviousValues) GetCustomText18Ok() (*string, bool)`

GetCustomText18Ok returns a tuple with the CustomText18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText18

`func (o *TestsPutPreviousValues) SetCustomText18(v string)`

SetCustomText18 sets CustomText18 field to given value.

### HasCustomText18

`func (o *TestsPutPreviousValues) HasCustomText18() bool`

HasCustomText18 returns a boolean if a field has been set.

### GetCustomText19

`func (o *TestsPutPreviousValues) GetCustomText19() string`

GetCustomText19 returns the CustomText19 field if non-nil, zero value otherwise.

### GetCustomText19Ok

`func (o *TestsPutPreviousValues) GetCustomText19Ok() (*string, bool)`

GetCustomText19Ok returns a tuple with the CustomText19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText19

`func (o *TestsPutPreviousValues) SetCustomText19(v string)`

SetCustomText19 sets CustomText19 field to given value.

### HasCustomText19

`func (o *TestsPutPreviousValues) HasCustomText19() bool`

HasCustomText19 returns a boolean if a field has been set.

### GetCustomText20

`func (o *TestsPutPreviousValues) GetCustomText20() string`

GetCustomText20 returns the CustomText20 field if non-nil, zero value otherwise.

### GetCustomText20Ok

`func (o *TestsPutPreviousValues) GetCustomText20Ok() (*string, bool)`

GetCustomText20Ok returns a tuple with the CustomText20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText20

`func (o *TestsPutPreviousValues) SetCustomText20(v string)`

SetCustomText20 sets CustomText20 field to given value.

### HasCustomText20

`func (o *TestsPutPreviousValues) HasCustomText20() bool`

HasCustomText20 returns a boolean if a field has been set.

### GetCustomText21

`func (o *TestsPutPreviousValues) GetCustomText21() string`

GetCustomText21 returns the CustomText21 field if non-nil, zero value otherwise.

### GetCustomText21Ok

`func (o *TestsPutPreviousValues) GetCustomText21Ok() (*string, bool)`

GetCustomText21Ok returns a tuple with the CustomText21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText21

`func (o *TestsPutPreviousValues) SetCustomText21(v string)`

SetCustomText21 sets CustomText21 field to given value.

### HasCustomText21

`func (o *TestsPutPreviousValues) HasCustomText21() bool`

HasCustomText21 returns a boolean if a field has been set.

### GetCustomText22

`func (o *TestsPutPreviousValues) GetCustomText22() string`

GetCustomText22 returns the CustomText22 field if non-nil, zero value otherwise.

### GetCustomText22Ok

`func (o *TestsPutPreviousValues) GetCustomText22Ok() (*string, bool)`

GetCustomText22Ok returns a tuple with the CustomText22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText22

`func (o *TestsPutPreviousValues) SetCustomText22(v string)`

SetCustomText22 sets CustomText22 field to given value.

### HasCustomText22

`func (o *TestsPutPreviousValues) HasCustomText22() bool`

HasCustomText22 returns a boolean if a field has been set.

### GetCustomText23

`func (o *TestsPutPreviousValues) GetCustomText23() string`

GetCustomText23 returns the CustomText23 field if non-nil, zero value otherwise.

### GetCustomText23Ok

`func (o *TestsPutPreviousValues) GetCustomText23Ok() (*string, bool)`

GetCustomText23Ok returns a tuple with the CustomText23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText23

`func (o *TestsPutPreviousValues) SetCustomText23(v string)`

SetCustomText23 sets CustomText23 field to given value.

### HasCustomText23

`func (o *TestsPutPreviousValues) HasCustomText23() bool`

HasCustomText23 returns a boolean if a field has been set.

### GetCustomText24

`func (o *TestsPutPreviousValues) GetCustomText24() string`

GetCustomText24 returns the CustomText24 field if non-nil, zero value otherwise.

### GetCustomText24Ok

`func (o *TestsPutPreviousValues) GetCustomText24Ok() (*string, bool)`

GetCustomText24Ok returns a tuple with the CustomText24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText24

`func (o *TestsPutPreviousValues) SetCustomText24(v string)`

SetCustomText24 sets CustomText24 field to given value.

### HasCustomText24

`func (o *TestsPutPreviousValues) HasCustomText24() bool`

HasCustomText24 returns a boolean if a field has been set.

### GetCustomText25

`func (o *TestsPutPreviousValues) GetCustomText25() string`

GetCustomText25 returns the CustomText25 field if non-nil, zero value otherwise.

### GetCustomText25Ok

`func (o *TestsPutPreviousValues) GetCustomText25Ok() (*string, bool)`

GetCustomText25Ok returns a tuple with the CustomText25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText25

`func (o *TestsPutPreviousValues) SetCustomText25(v string)`

SetCustomText25 sets CustomText25 field to given value.

### HasCustomText25

`func (o *TestsPutPreviousValues) HasCustomText25() bool`

HasCustomText25 returns a boolean if a field has been set.

### GetCustomText26

`func (o *TestsPutPreviousValues) GetCustomText26() string`

GetCustomText26 returns the CustomText26 field if non-nil, zero value otherwise.

### GetCustomText26Ok

`func (o *TestsPutPreviousValues) GetCustomText26Ok() (*string, bool)`

GetCustomText26Ok returns a tuple with the CustomText26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText26

`func (o *TestsPutPreviousValues) SetCustomText26(v string)`

SetCustomText26 sets CustomText26 field to given value.

### HasCustomText26

`func (o *TestsPutPreviousValues) HasCustomText26() bool`

HasCustomText26 returns a boolean if a field has been set.

### GetCustomText27

`func (o *TestsPutPreviousValues) GetCustomText27() string`

GetCustomText27 returns the CustomText27 field if non-nil, zero value otherwise.

### GetCustomText27Ok

`func (o *TestsPutPreviousValues) GetCustomText27Ok() (*string, bool)`

GetCustomText27Ok returns a tuple with the CustomText27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText27

`func (o *TestsPutPreviousValues) SetCustomText27(v string)`

SetCustomText27 sets CustomText27 field to given value.

### HasCustomText27

`func (o *TestsPutPreviousValues) HasCustomText27() bool`

HasCustomText27 returns a boolean if a field has been set.

### GetCustomText28

`func (o *TestsPutPreviousValues) GetCustomText28() string`

GetCustomText28 returns the CustomText28 field if non-nil, zero value otherwise.

### GetCustomText28Ok

`func (o *TestsPutPreviousValues) GetCustomText28Ok() (*string, bool)`

GetCustomText28Ok returns a tuple with the CustomText28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText28

`func (o *TestsPutPreviousValues) SetCustomText28(v string)`

SetCustomText28 sets CustomText28 field to given value.

### HasCustomText28

`func (o *TestsPutPreviousValues) HasCustomText28() bool`

HasCustomText28 returns a boolean if a field has been set.

### GetCustomText29

`func (o *TestsPutPreviousValues) GetCustomText29() string`

GetCustomText29 returns the CustomText29 field if non-nil, zero value otherwise.

### GetCustomText29Ok

`func (o *TestsPutPreviousValues) GetCustomText29Ok() (*string, bool)`

GetCustomText29Ok returns a tuple with the CustomText29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText29

`func (o *TestsPutPreviousValues) SetCustomText29(v string)`

SetCustomText29 sets CustomText29 field to given value.

### HasCustomText29

`func (o *TestsPutPreviousValues) HasCustomText29() bool`

HasCustomText29 returns a boolean if a field has been set.

### GetCustomText30

`func (o *TestsPutPreviousValues) GetCustomText30() string`

GetCustomText30 returns the CustomText30 field if non-nil, zero value otherwise.

### GetCustomText30Ok

`func (o *TestsPutPreviousValues) GetCustomText30Ok() (*string, bool)`

GetCustomText30Ok returns a tuple with the CustomText30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText30

`func (o *TestsPutPreviousValues) SetCustomText30(v string)`

SetCustomText30 sets CustomText30 field to given value.

### HasCustomText30

`func (o *TestsPutPreviousValues) HasCustomText30() bool`

HasCustomText30 returns a boolean if a field has been set.

### GetCustomText31

`func (o *TestsPutPreviousValues) GetCustomText31() string`

GetCustomText31 returns the CustomText31 field if non-nil, zero value otherwise.

### GetCustomText31Ok

`func (o *TestsPutPreviousValues) GetCustomText31Ok() (*string, bool)`

GetCustomText31Ok returns a tuple with the CustomText31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText31

`func (o *TestsPutPreviousValues) SetCustomText31(v string)`

SetCustomText31 sets CustomText31 field to given value.

### HasCustomText31

`func (o *TestsPutPreviousValues) HasCustomText31() bool`

HasCustomText31 returns a boolean if a field has been set.

### GetCustomText32

`func (o *TestsPutPreviousValues) GetCustomText32() string`

GetCustomText32 returns the CustomText32 field if non-nil, zero value otherwise.

### GetCustomText32Ok

`func (o *TestsPutPreviousValues) GetCustomText32Ok() (*string, bool)`

GetCustomText32Ok returns a tuple with the CustomText32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText32

`func (o *TestsPutPreviousValues) SetCustomText32(v string)`

SetCustomText32 sets CustomText32 field to given value.

### HasCustomText32

`func (o *TestsPutPreviousValues) HasCustomText32() bool`

HasCustomText32 returns a boolean if a field has been set.

### GetCustomText33

`func (o *TestsPutPreviousValues) GetCustomText33() string`

GetCustomText33 returns the CustomText33 field if non-nil, zero value otherwise.

### GetCustomText33Ok

`func (o *TestsPutPreviousValues) GetCustomText33Ok() (*string, bool)`

GetCustomText33Ok returns a tuple with the CustomText33 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText33

`func (o *TestsPutPreviousValues) SetCustomText33(v string)`

SetCustomText33 sets CustomText33 field to given value.

### HasCustomText33

`func (o *TestsPutPreviousValues) HasCustomText33() bool`

HasCustomText33 returns a boolean if a field has been set.

### GetTestCustomSelect8OptionId

`func (o *TestsPutPreviousValues) GetTestCustomSelect8OptionId() int32`

GetTestCustomSelect8OptionId returns the TestCustomSelect8OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect8OptionIdOk

`func (o *TestsPutPreviousValues) GetTestCustomSelect8OptionIdOk() (*int32, bool)`

GetTestCustomSelect8OptionIdOk returns a tuple with the TestCustomSelect8OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect8OptionId

`func (o *TestsPutPreviousValues) SetTestCustomSelect8OptionId(v int32)`

SetTestCustomSelect8OptionId sets TestCustomSelect8OptionId field to given value.

### HasTestCustomSelect8OptionId

`func (o *TestsPutPreviousValues) HasTestCustomSelect8OptionId() bool`

HasTestCustomSelect8OptionId returns a boolean if a field has been set.

### GetTestCustomSelect9OptionId

`func (o *TestsPutPreviousValues) GetTestCustomSelect9OptionId() int32`

GetTestCustomSelect9OptionId returns the TestCustomSelect9OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect9OptionIdOk

`func (o *TestsPutPreviousValues) GetTestCustomSelect9OptionIdOk() (*int32, bool)`

GetTestCustomSelect9OptionIdOk returns a tuple with the TestCustomSelect9OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect9OptionId

`func (o *TestsPutPreviousValues) SetTestCustomSelect9OptionId(v int32)`

SetTestCustomSelect9OptionId sets TestCustomSelect9OptionId field to given value.

### HasTestCustomSelect9OptionId

`func (o *TestsPutPreviousValues) HasTestCustomSelect9OptionId() bool`

HasTestCustomSelect9OptionId returns a boolean if a field has been set.

### GetTestCustomSelect10OptionId

`func (o *TestsPutPreviousValues) GetTestCustomSelect10OptionId() int32`

GetTestCustomSelect10OptionId returns the TestCustomSelect10OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect10OptionIdOk

`func (o *TestsPutPreviousValues) GetTestCustomSelect10OptionIdOk() (*int32, bool)`

GetTestCustomSelect10OptionIdOk returns a tuple with the TestCustomSelect10OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect10OptionId

`func (o *TestsPutPreviousValues) SetTestCustomSelect10OptionId(v int32)`

SetTestCustomSelect10OptionId sets TestCustomSelect10OptionId field to given value.

### HasTestCustomSelect10OptionId

`func (o *TestsPutPreviousValues) HasTestCustomSelect10OptionId() bool`

HasTestCustomSelect10OptionId returns a boolean if a field has been set.

### GetTestCustomSelect11OptionId

`func (o *TestsPutPreviousValues) GetTestCustomSelect11OptionId() int32`

GetTestCustomSelect11OptionId returns the TestCustomSelect11OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect11OptionIdOk

`func (o *TestsPutPreviousValues) GetTestCustomSelect11OptionIdOk() (*int32, bool)`

GetTestCustomSelect11OptionIdOk returns a tuple with the TestCustomSelect11OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect11OptionId

`func (o *TestsPutPreviousValues) SetTestCustomSelect11OptionId(v int32)`

SetTestCustomSelect11OptionId sets TestCustomSelect11OptionId field to given value.

### HasTestCustomSelect11OptionId

`func (o *TestsPutPreviousValues) HasTestCustomSelect11OptionId() bool`

HasTestCustomSelect11OptionId returns a boolean if a field has been set.

### GetTestCustomSelect12OptionId

`func (o *TestsPutPreviousValues) GetTestCustomSelect12OptionId() int32`

GetTestCustomSelect12OptionId returns the TestCustomSelect12OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect12OptionIdOk

`func (o *TestsPutPreviousValues) GetTestCustomSelect12OptionIdOk() (*int32, bool)`

GetTestCustomSelect12OptionIdOk returns a tuple with the TestCustomSelect12OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect12OptionId

`func (o *TestsPutPreviousValues) SetTestCustomSelect12OptionId(v int32)`

SetTestCustomSelect12OptionId sets TestCustomSelect12OptionId field to given value.

### HasTestCustomSelect12OptionId

`func (o *TestsPutPreviousValues) HasTestCustomSelect12OptionId() bool`

HasTestCustomSelect12OptionId returns a boolean if a field has been set.

### GetTestCustomSelect13OptionId

`func (o *TestsPutPreviousValues) GetTestCustomSelect13OptionId() int32`

GetTestCustomSelect13OptionId returns the TestCustomSelect13OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect13OptionIdOk

`func (o *TestsPutPreviousValues) GetTestCustomSelect13OptionIdOk() (*int32, bool)`

GetTestCustomSelect13OptionIdOk returns a tuple with the TestCustomSelect13OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect13OptionId

`func (o *TestsPutPreviousValues) SetTestCustomSelect13OptionId(v int32)`

SetTestCustomSelect13OptionId sets TestCustomSelect13OptionId field to given value.

### HasTestCustomSelect13OptionId

`func (o *TestsPutPreviousValues) HasTestCustomSelect13OptionId() bool`

HasTestCustomSelect13OptionId returns a boolean if a field has been set.

### GetTestCustomSelect14OptionId

`func (o *TestsPutPreviousValues) GetTestCustomSelect14OptionId() int32`

GetTestCustomSelect14OptionId returns the TestCustomSelect14OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect14OptionIdOk

`func (o *TestsPutPreviousValues) GetTestCustomSelect14OptionIdOk() (*int32, bool)`

GetTestCustomSelect14OptionIdOk returns a tuple with the TestCustomSelect14OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect14OptionId

`func (o *TestsPutPreviousValues) SetTestCustomSelect14OptionId(v int32)`

SetTestCustomSelect14OptionId sets TestCustomSelect14OptionId field to given value.

### HasTestCustomSelect14OptionId

`func (o *TestsPutPreviousValues) HasTestCustomSelect14OptionId() bool`

HasTestCustomSelect14OptionId returns a boolean if a field has been set.

### GetTestCustomSelect15OptionId

`func (o *TestsPutPreviousValues) GetTestCustomSelect15OptionId() int32`

GetTestCustomSelect15OptionId returns the TestCustomSelect15OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect15OptionIdOk

`func (o *TestsPutPreviousValues) GetTestCustomSelect15OptionIdOk() (*int32, bool)`

GetTestCustomSelect15OptionIdOk returns a tuple with the TestCustomSelect15OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect15OptionId

`func (o *TestsPutPreviousValues) SetTestCustomSelect15OptionId(v int32)`

SetTestCustomSelect15OptionId sets TestCustomSelect15OptionId field to given value.

### HasTestCustomSelect15OptionId

`func (o *TestsPutPreviousValues) HasTestCustomSelect15OptionId() bool`

HasTestCustomSelect15OptionId returns a boolean if a field has been set.

### GetTestCustomSelect16OptionId

`func (o *TestsPutPreviousValues) GetTestCustomSelect16OptionId() int32`

GetTestCustomSelect16OptionId returns the TestCustomSelect16OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect16OptionIdOk

`func (o *TestsPutPreviousValues) GetTestCustomSelect16OptionIdOk() (*int32, bool)`

GetTestCustomSelect16OptionIdOk returns a tuple with the TestCustomSelect16OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect16OptionId

`func (o *TestsPutPreviousValues) SetTestCustomSelect16OptionId(v int32)`

SetTestCustomSelect16OptionId sets TestCustomSelect16OptionId field to given value.

### HasTestCustomSelect16OptionId

`func (o *TestsPutPreviousValues) HasTestCustomSelect16OptionId() bool`

HasTestCustomSelect16OptionId returns a boolean if a field has been set.

### GetTestCustomSelect17OptionId

`func (o *TestsPutPreviousValues) GetTestCustomSelect17OptionId() int32`

GetTestCustomSelect17OptionId returns the TestCustomSelect17OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect17OptionIdOk

`func (o *TestsPutPreviousValues) GetTestCustomSelect17OptionIdOk() (*int32, bool)`

GetTestCustomSelect17OptionIdOk returns a tuple with the TestCustomSelect17OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect17OptionId

`func (o *TestsPutPreviousValues) SetTestCustomSelect17OptionId(v int32)`

SetTestCustomSelect17OptionId sets TestCustomSelect17OptionId field to given value.

### HasTestCustomSelect17OptionId

`func (o *TestsPutPreviousValues) HasTestCustomSelect17OptionId() bool`

HasTestCustomSelect17OptionId returns a boolean if a field has been set.

### GetFieldData

`func (o *TestsPutPreviousValues) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *TestsPutPreviousValues) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *TestsPutPreviousValues) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *TestsPutPreviousValues) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *TestsPutPreviousValues) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *TestsPutPreviousValues) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


