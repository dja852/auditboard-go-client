# TestsPutTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**ControlsDatumId** | **int32** | Note: This is a Foreign Key to &#x60;controls_data.id&#x60;.&lt;fk table&#x3D;&#39;controls_data&#39; column&#x3D;&#39;id&#39;/&gt; | 
**Order** | **int32** |  | 
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
**Archived** | **bool** |  | [default to false]
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
**IsInactive** | **bool** |  | [default to false]
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

### NewTestsPutTest

`func NewTestsPutTest(controlsDatumId int32, order int32, archived bool, isInactive bool, ) *TestsPutTest`

NewTestsPutTest instantiates a new TestsPutTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestsPutTestWithDefaults

`func NewTestsPutTestWithDefaults() *TestsPutTest`

NewTestsPutTestWithDefaults instantiates a new TestsPutTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestsPutTest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestsPutTest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestsPutTest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TestsPutTest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TestsPutTest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TestsPutTest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TestsPutTest) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TestsPutTest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TestsPutTest) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TestsPutTest) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TestsPutTest) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TestsPutTest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *TestsPutTest) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *TestsPutTest) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *TestsPutTest) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *TestsPutTest) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetControlsDatumId

`func (o *TestsPutTest) GetControlsDatumId() int32`

GetControlsDatumId returns the ControlsDatumId field if non-nil, zero value otherwise.

### GetControlsDatumIdOk

`func (o *TestsPutTest) GetControlsDatumIdOk() (*int32, bool)`

GetControlsDatumIdOk returns a tuple with the ControlsDatumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlsDatumId

`func (o *TestsPutTest) SetControlsDatumId(v int32)`

SetControlsDatumId sets ControlsDatumId field to given value.


### GetOrder

`func (o *TestsPutTest) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TestsPutTest) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TestsPutTest) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetName

`func (o *TestsPutTest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestsPutTest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestsPutTest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestsPutTest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSampleSize

`func (o *TestsPutTest) GetSampleSize() string`

GetSampleSize returns the SampleSize field if non-nil, zero value otherwise.

### GetSampleSizeOk

`func (o *TestsPutTest) GetSampleSizeOk() (*string, bool)`

GetSampleSizeOk returns a tuple with the SampleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleSize

`func (o *TestsPutTest) SetSampleSize(v string)`

SetSampleSize sets SampleSize field to given value.

### HasSampleSize

`func (o *TestsPutTest) HasSampleSize() bool`

HasSampleSize returns a boolean if a field has been set.

### GetSelections

`func (o *TestsPutTest) GetSelections() string`

GetSelections returns the Selections field if non-nil, zero value otherwise.

### GetSelectionsOk

`func (o *TestsPutTest) GetSelectionsOk() (*string, bool)`

GetSelectionsOk returns a tuple with the Selections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelections

`func (o *TestsPutTest) SetSelections(v string)`

SetSelections sets Selections field to given value.

### HasSelections

`func (o *TestsPutTest) HasSelections() bool`

HasSelections returns a boolean if a field has been set.

### GetTesterUserId

`func (o *TestsPutTest) GetTesterUserId() int32`

GetTesterUserId returns the TesterUserId field if non-nil, zero value otherwise.

### GetTesterUserIdOk

`func (o *TestsPutTest) GetTesterUserIdOk() (*int32, bool)`

GetTesterUserIdOk returns a tuple with the TesterUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesterUserId

`func (o *TestsPutTest) SetTesterUserId(v int32)`

SetTesterUserId sets TesterUserId field to given value.

### HasTesterUserId

`func (o *TestsPutTest) HasTesterUserId() bool`

HasTesterUserId returns a boolean if a field has been set.

### GetReviewerUserId

`func (o *TestsPutTest) GetReviewerUserId() int32`

GetReviewerUserId returns the ReviewerUserId field if non-nil, zero value otherwise.

### GetReviewerUserIdOk

`func (o *TestsPutTest) GetReviewerUserIdOk() (*int32, bool)`

GetReviewerUserIdOk returns a tuple with the ReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerUserId

`func (o *TestsPutTest) SetReviewerUserId(v int32)`

SetReviewerUserId sets ReviewerUserId field to given value.

### HasReviewerUserId

`func (o *TestsPutTest) HasReviewerUserId() bool`

HasReviewerUserId returns a boolean if a field has been set.

### GetStartDate

`func (o *TestsPutTest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TestsPutTest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TestsPutTest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TestsPutTest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *TestsPutTest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TestsPutTest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TestsPutTest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TestsPutTest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetResults

`func (o *TestsPutTest) GetResults() string`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TestsPutTest) GetResultsOk() (*string, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TestsPutTest) SetResults(v string)`

SetResults sets Results field to given value.

### HasResults

`func (o *TestsPutTest) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTestSheet

`func (o *TestsPutTest) GetTestSheet() string`

GetTestSheet returns the TestSheet field if non-nil, zero value otherwise.

### GetTestSheetOk

`func (o *TestsPutTest) GetTestSheetOk() (*string, bool)`

GetTestSheetOk returns a tuple with the TestSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSheet

`func (o *TestsPutTest) SetTestSheet(v string)`

SetTestSheet sets TestSheet field to given value.

### HasTestSheet

`func (o *TestsPutTest) HasTestSheet() bool`

HasTestSheet returns a boolean if a field has been set.

### GetNotes

`func (o *TestsPutTest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TestsPutTest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TestsPutTest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TestsPutTest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetEffectivenessOptionId

`func (o *TestsPutTest) GetEffectivenessOptionId() int32`

GetEffectivenessOptionId returns the EffectivenessOptionId field if non-nil, zero value otherwise.

### GetEffectivenessOptionIdOk

`func (o *TestsPutTest) GetEffectivenessOptionIdOk() (*int32, bool)`

GetEffectivenessOptionIdOk returns a tuple with the EffectivenessOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectivenessOptionId

`func (o *TestsPutTest) SetEffectivenessOptionId(v int32)`

SetEffectivenessOptionId sets EffectivenessOptionId field to given value.

### HasEffectivenessOptionId

`func (o *TestsPutTest) HasEffectivenessOptionId() bool`

HasEffectivenessOptionId returns a boolean if a field has been set.

### GetStatusOptionId

`func (o *TestsPutTest) GetStatusOptionId() int32`

GetStatusOptionId returns the StatusOptionId field if non-nil, zero value otherwise.

### GetStatusOptionIdOk

`func (o *TestsPutTest) GetStatusOptionIdOk() (*int32, bool)`

GetStatusOptionIdOk returns a tuple with the StatusOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusOptionId

`func (o *TestsPutTest) SetStatusOptionId(v int32)`

SetStatusOptionId sets StatusOptionId field to given value.

### HasStatusOptionId

`func (o *TestsPutTest) HasStatusOptionId() bool`

HasStatusOptionId returns a boolean if a field has been set.

### GetArchived

`func (o *TestsPutTest) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *TestsPutTest) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *TestsPutTest) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetArchivedAt

`func (o *TestsPutTest) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *TestsPutTest) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *TestsPutTest) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *TestsPutTest) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetOldHours

`func (o *TestsPutTest) GetOldHours() string`

GetOldHours returns the OldHours field if non-nil, zero value otherwise.

### GetOldHoursOk

`func (o *TestsPutTest) GetOldHoursOk() (*string, bool)`

GetOldHoursOk returns a tuple with the OldHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldHours

`func (o *TestsPutTest) SetOldHours(v string)`

SetOldHours sets OldHours field to given value.

### HasOldHours

`func (o *TestsPutTest) HasOldHours() bool`

HasOldHours returns a boolean if a field has been set.

### GetDueDate

`func (o *TestsPutTest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *TestsPutTest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *TestsPutTest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *TestsPutTest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetOldBudgetHours

`func (o *TestsPutTest) GetOldBudgetHours() string`

GetOldBudgetHours returns the OldBudgetHours field if non-nil, zero value otherwise.

### GetOldBudgetHoursOk

`func (o *TestsPutTest) GetOldBudgetHoursOk() (*string, bool)`

GetOldBudgetHoursOk returns a tuple with the OldBudgetHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldBudgetHours

`func (o *TestsPutTest) SetOldBudgetHours(v string)`

SetOldBudgetHours sets OldBudgetHours field to given value.

### HasOldBudgetHours

`func (o *TestsPutTest) HasOldBudgetHours() bool`

HasOldBudgetHours returns a boolean if a field has been set.

### GetOldPercentComplete

`func (o *TestsPutTest) GetOldPercentComplete() string`

GetOldPercentComplete returns the OldPercentComplete field if non-nil, zero value otherwise.

### GetOldPercentCompleteOk

`func (o *TestsPutTest) GetOldPercentCompleteOk() (*string, bool)`

GetOldPercentCompleteOk returns a tuple with the OldPercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPercentComplete

`func (o *TestsPutTest) SetOldPercentComplete(v string)`

SetOldPercentComplete sets OldPercentComplete field to given value.

### HasOldPercentComplete

`func (o *TestsPutTest) HasOldPercentComplete() bool`

HasOldPercentComplete returns a boolean if a field has been set.

### GetTestPeriod

`func (o *TestsPutTest) GetTestPeriod() string`

GetTestPeriod returns the TestPeriod field if non-nil, zero value otherwise.

### GetTestPeriodOk

`func (o *TestsPutTest) GetTestPeriodOk() (*string, bool)`

GetTestPeriodOk returns a tuple with the TestPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPeriod

`func (o *TestsPutTest) SetTestPeriod(v string)`

SetTestPeriod sets TestPeriod field to given value.

### HasTestPeriod

`func (o *TestsPutTest) HasTestPeriod() bool`

HasTestPeriod returns a boolean if a field has been set.

### GetEntityId

`func (o *TestsPutTest) GetEntityId() int32`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *TestsPutTest) GetEntityIdOk() (*int32, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *TestsPutTest) SetEntityId(v int32)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *TestsPutTest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetControlId

`func (o *TestsPutTest) GetControlId() int32`

GetControlId returns the ControlId field if non-nil, zero value otherwise.

### GetControlIdOk

`func (o *TestsPutTest) GetControlIdOk() (*int32, bool)`

GetControlIdOk returns a tuple with the ControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlId

`func (o *TestsPutTest) SetControlId(v int32)`

SetControlId sets ControlId field to given value.

### HasControlId

`func (o *TestsPutTest) HasControlId() bool`

HasControlId returns a boolean if a field has been set.

### GetPrecisionLevel

`func (o *TestsPutTest) GetPrecisionLevel() string`

GetPrecisionLevel returns the PrecisionLevel field if non-nil, zero value otherwise.

### GetPrecisionLevelOk

`func (o *TestsPutTest) GetPrecisionLevelOk() (*string, bool)`

GetPrecisionLevelOk returns a tuple with the PrecisionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecisionLevel

`func (o *TestsPutTest) SetPrecisionLevel(v string)`

SetPrecisionLevel sets PrecisionLevel field to given value.

### HasPrecisionLevel

`func (o *TestsPutTest) HasPrecisionLevel() bool`

HasPrecisionLevel returns a boolean if a field has been set.

### GetMrcEvidence

`func (o *TestsPutTest) GetMrcEvidence() string`

GetMrcEvidence returns the MrcEvidence field if non-nil, zero value otherwise.

### GetMrcEvidenceOk

`func (o *TestsPutTest) GetMrcEvidenceOk() (*string, bool)`

GetMrcEvidenceOk returns a tuple with the MrcEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrcEvidence

`func (o *TestsPutTest) SetMrcEvidence(v string)`

SetMrcEvidence sets MrcEvidence field to given value.

### HasMrcEvidence

`func (o *TestsPutTest) HasMrcEvidence() bool`

HasMrcEvidence returns a boolean if a field has been set.

### GetSheetId

`func (o *TestsPutTest) GetSheetId() int32`

GetSheetId returns the SheetId field if non-nil, zero value otherwise.

### GetSheetIdOk

`func (o *TestsPutTest) GetSheetIdOk() (*int32, bool)`

GetSheetIdOk returns a tuple with the SheetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheetId

`func (o *TestsPutTest) SetSheetId(v int32)`

SetSheetId sets SheetId field to given value.

### HasSheetId

`func (o *TestsPutTest) HasSheetId() bool`

HasSheetId returns a boolean if a field has been set.

### GetIsInactive

`func (o *TestsPutTest) GetIsInactive() bool`

GetIsInactive returns the IsInactive field if non-nil, zero value otherwise.

### GetIsInactiveOk

`func (o *TestsPutTest) GetIsInactiveOk() (*bool, bool)`

GetIsInactiveOk returns a tuple with the IsInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInactive

`func (o *TestsPutTest) SetIsInactive(v bool)`

SetIsInactive sets IsInactive field to given value.


### GetCompleteDate

`func (o *TestsPutTest) GetCompleteDate() string`

GetCompleteDate returns the CompleteDate field if non-nil, zero value otherwise.

### GetCompleteDateOk

`func (o *TestsPutTest) GetCompleteDateOk() (*string, bool)`

GetCompleteDateOk returns a tuple with the CompleteDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteDate

`func (o *TestsPutTest) SetCompleteDate(v string)`

SetCompleteDate sets CompleteDate field to given value.

### HasCompleteDate

`func (o *TestsPutTest) HasCompleteDate() bool`

HasCompleteDate returns a boolean if a field has been set.

### GetReviewedDate

`func (o *TestsPutTest) GetReviewedDate() string`

GetReviewedDate returns the ReviewedDate field if non-nil, zero value otherwise.

### GetReviewedDateOk

`func (o *TestsPutTest) GetReviewedDateOk() (*string, bool)`

GetReviewedDateOk returns a tuple with the ReviewedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedDate

`func (o *TestsPutTest) SetReviewedDate(v string)`

SetReviewedDate sets ReviewedDate field to given value.

### HasReviewedDate

`func (o *TestsPutTest) HasReviewedDate() bool`

HasReviewedDate returns a boolean if a field has been set.

### GetReviewNotesDate

`func (o *TestsPutTest) GetReviewNotesDate() string`

GetReviewNotesDate returns the ReviewNotesDate field if non-nil, zero value otherwise.

### GetReviewNotesDateOk

`func (o *TestsPutTest) GetReviewNotesDateOk() (*string, bool)`

GetReviewNotesDateOk returns a tuple with the ReviewNotesDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewNotesDate

`func (o *TestsPutTest) SetReviewNotesDate(v string)`

SetReviewNotesDate sets ReviewNotesDate field to given value.

### HasReviewNotesDate

`func (o *TestsPutTest) HasReviewNotesDate() bool`

HasReviewNotesDate returns a boolean if a field has been set.

### GetCompleteByUserId

`func (o *TestsPutTest) GetCompleteByUserId() int32`

GetCompleteByUserId returns the CompleteByUserId field if non-nil, zero value otherwise.

### GetCompleteByUserIdOk

`func (o *TestsPutTest) GetCompleteByUserIdOk() (*int32, bool)`

GetCompleteByUserIdOk returns a tuple with the CompleteByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteByUserId

`func (o *TestsPutTest) SetCompleteByUserId(v int32)`

SetCompleteByUserId sets CompleteByUserId field to given value.

### HasCompleteByUserId

`func (o *TestsPutTest) HasCompleteByUserId() bool`

HasCompleteByUserId returns a boolean if a field has been set.

### GetReviewedByUserId

`func (o *TestsPutTest) GetReviewedByUserId() int32`

GetReviewedByUserId returns the ReviewedByUserId field if non-nil, zero value otherwise.

### GetReviewedByUserIdOk

`func (o *TestsPutTest) GetReviewedByUserIdOk() (*int32, bool)`

GetReviewedByUserIdOk returns a tuple with the ReviewedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedByUserId

`func (o *TestsPutTest) SetReviewedByUserId(v int32)`

SetReviewedByUserId sets ReviewedByUserId field to given value.

### HasReviewedByUserId

`func (o *TestsPutTest) HasReviewedByUserId() bool`

HasReviewedByUserId returns a boolean if a field has been set.

### GetReviewNotesByUserId

`func (o *TestsPutTest) GetReviewNotesByUserId() int32`

GetReviewNotesByUserId returns the ReviewNotesByUserId field if non-nil, zero value otherwise.

### GetReviewNotesByUserIdOk

`func (o *TestsPutTest) GetReviewNotesByUserIdOk() (*int32, bool)`

GetReviewNotesByUserIdOk returns a tuple with the ReviewNotesByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewNotesByUserId

`func (o *TestsPutTest) SetReviewNotesByUserId(v int32)`

SetReviewNotesByUserId sets ReviewNotesByUserId field to given value.

### HasReviewNotesByUserId

`func (o *TestsPutTest) HasReviewNotesByUserId() bool`

HasReviewNotesByUserId returns a boolean if a field has been set.

### GetTestSectionId

`func (o *TestsPutTest) GetTestSectionId() int32`

GetTestSectionId returns the TestSectionId field if non-nil, zero value otherwise.

### GetTestSectionIdOk

`func (o *TestsPutTest) GetTestSectionIdOk() (*int32, bool)`

GetTestSectionIdOk returns a tuple with the TestSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSectionId

`func (o *TestsPutTest) SetTestSectionId(v int32)`

SetTestSectionId sets TestSectionId field to given value.

### HasTestSectionId

`func (o *TestsPutTest) HasTestSectionId() bool`

HasTestSectionId returns a boolean if a field has been set.

### GetProcedures

`func (o *TestsPutTest) GetProcedures() string`

GetProcedures returns the Procedures field if non-nil, zero value otherwise.

### GetProceduresOk

`func (o *TestsPutTest) GetProceduresOk() (*string, bool)`

GetProceduresOk returns a tuple with the Procedures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcedures

`func (o *TestsPutTest) SetProcedures(v string)`

SetProcedures sets Procedures field to given value.

### HasProcedures

`func (o *TestsPutTest) HasProcedures() bool`

HasProcedures returns a boolean if a field has been set.

### GetSheetFileId

`func (o *TestsPutTest) GetSheetFileId() int32`

GetSheetFileId returns the SheetFileId field if non-nil, zero value otherwise.

### GetSheetFileIdOk

`func (o *TestsPutTest) GetSheetFileIdOk() (*int32, bool)`

GetSheetFileIdOk returns a tuple with the SheetFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheetFileId

`func (o *TestsPutTest) SetSheetFileId(v int32)`

SetSheetFileId sets SheetFileId field to given value.

### HasSheetFileId

`func (o *TestsPutTest) HasSheetFileId() bool`

HasSheetFileId returns a boolean if a field has been set.

### GetOldReviewerHours

`func (o *TestsPutTest) GetOldReviewerHours() string`

GetOldReviewerHours returns the OldReviewerHours field if non-nil, zero value otherwise.

### GetOldReviewerHoursOk

`func (o *TestsPutTest) GetOldReviewerHoursOk() (*string, bool)`

GetOldReviewerHoursOk returns a tuple with the OldReviewerHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldReviewerHours

`func (o *TestsPutTest) SetOldReviewerHours(v string)`

SetOldReviewerHours sets OldReviewerHours field to given value.

### HasOldReviewerHours

`func (o *TestsPutTest) HasOldReviewerHours() bool`

HasOldReviewerHours returns a boolean if a field has been set.

### GetOldReviewerBudgetHours

`func (o *TestsPutTest) GetOldReviewerBudgetHours() string`

GetOldReviewerBudgetHours returns the OldReviewerBudgetHours field if non-nil, zero value otherwise.

### GetOldReviewerBudgetHoursOk

`func (o *TestsPutTest) GetOldReviewerBudgetHoursOk() (*string, bool)`

GetOldReviewerBudgetHoursOk returns a tuple with the OldReviewerBudgetHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldReviewerBudgetHours

`func (o *TestsPutTest) SetOldReviewerBudgetHours(v string)`

SetOldReviewerBudgetHours sets OldReviewerBudgetHours field to given value.

### HasOldReviewerBudgetHours

`func (o *TestsPutTest) HasOldReviewerBudgetHours() bool`

HasOldReviewerBudgetHours returns a boolean if a field has been set.

### GetSecondaryReviewerUserId

`func (o *TestsPutTest) GetSecondaryReviewerUserId() int32`

GetSecondaryReviewerUserId returns the SecondaryReviewerUserId field if non-nil, zero value otherwise.

### GetSecondaryReviewerUserIdOk

`func (o *TestsPutTest) GetSecondaryReviewerUserIdOk() (*int32, bool)`

GetSecondaryReviewerUserIdOk returns a tuple with the SecondaryReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryReviewerUserId

`func (o *TestsPutTest) SetSecondaryReviewerUserId(v int32)`

SetSecondaryReviewerUserId sets SecondaryReviewerUserId field to given value.

### HasSecondaryReviewerUserId

`func (o *TestsPutTest) HasSecondaryReviewerUserId() bool`

HasSecondaryReviewerUserId returns a boolean if a field has been set.

### GetPopulationSize

`func (o *TestsPutTest) GetPopulationSize() string`

GetPopulationSize returns the PopulationSize field if non-nil, zero value otherwise.

### GetPopulationSizeOk

`func (o *TestsPutTest) GetPopulationSizeOk() (*string, bool)`

GetPopulationSizeOk returns a tuple with the PopulationSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationSize

`func (o *TestsPutTest) SetPopulationSize(v string)`

SetPopulationSize sets PopulationSize field to given value.

### HasPopulationSize

`func (o *TestsPutTest) HasPopulationSize() bool`

HasPopulationSize returns a boolean if a field has been set.

### GetTestPbcRequest

`func (o *TestsPutTest) GetTestPbcRequest() string`

GetTestPbcRequest returns the TestPbcRequest field if non-nil, zero value otherwise.

### GetTestPbcRequestOk

`func (o *TestsPutTest) GetTestPbcRequestOk() (*string, bool)`

GetTestPbcRequestOk returns a tuple with the TestPbcRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPbcRequest

`func (o *TestsPutTest) SetTestPbcRequest(v string)`

SetTestPbcRequest sets TestPbcRequest field to given value.

### HasTestPbcRequest

`func (o *TestsPutTest) HasTestPbcRequest() bool`

HasTestPbcRequest returns a boolean if a field has been set.

### GetTestPbcRequestAdditional

`func (o *TestsPutTest) GetTestPbcRequestAdditional() string`

GetTestPbcRequestAdditional returns the TestPbcRequestAdditional field if non-nil, zero value otherwise.

### GetTestPbcRequestAdditionalOk

`func (o *TestsPutTest) GetTestPbcRequestAdditionalOk() (*string, bool)`

GetTestPbcRequestAdditionalOk returns a tuple with the TestPbcRequestAdditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPbcRequestAdditional

`func (o *TestsPutTest) SetTestPbcRequestAdditional(v string)`

SetTestPbcRequestAdditional sets TestPbcRequestAdditional field to given value.

### HasTestPbcRequestAdditional

`func (o *TestsPutTest) HasTestPbcRequestAdditional() bool`

HasTestPbcRequestAdditional returns a boolean if a field has been set.

### GetPopulationCompleteness

`func (o *TestsPutTest) GetPopulationCompleteness() string`

GetPopulationCompleteness returns the PopulationCompleteness field if non-nil, zero value otherwise.

### GetPopulationCompletenessOk

`func (o *TestsPutTest) GetPopulationCompletenessOk() (*string, bool)`

GetPopulationCompletenessOk returns a tuple with the PopulationCompleteness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationCompleteness

`func (o *TestsPutTest) SetPopulationCompleteness(v string)`

SetPopulationCompleteness sets PopulationCompleteness field to given value.

### HasPopulationCompleteness

`func (o *TestsPutTest) HasPopulationCompleteness() bool`

HasPopulationCompleteness returns a boolean if a field has been set.

### GetPopulationSource

`func (o *TestsPutTest) GetPopulationSource() string`

GetPopulationSource returns the PopulationSource field if non-nil, zero value otherwise.

### GetPopulationSourceOk

`func (o *TestsPutTest) GetPopulationSourceOk() (*string, bool)`

GetPopulationSourceOk returns a tuple with the PopulationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationSource

`func (o *TestsPutTest) SetPopulationSource(v string)`

SetPopulationSource sets PopulationSource field to given value.

### HasPopulationSource

`func (o *TestsPutTest) HasPopulationSource() bool`

HasPopulationSource returns a boolean if a field has been set.

### GetExternalAuditorStatusOptionId

`func (o *TestsPutTest) GetExternalAuditorStatusOptionId() int32`

GetExternalAuditorStatusOptionId returns the ExternalAuditorStatusOptionId field if non-nil, zero value otherwise.

### GetExternalAuditorStatusOptionIdOk

`func (o *TestsPutTest) GetExternalAuditorStatusOptionIdOk() (*int32, bool)`

GetExternalAuditorStatusOptionIdOk returns a tuple with the ExternalAuditorStatusOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAuditorStatusOptionId

`func (o *TestsPutTest) SetExternalAuditorStatusOptionId(v int32)`

SetExternalAuditorStatusOptionId sets ExternalAuditorStatusOptionId field to given value.

### HasExternalAuditorStatusOptionId

`func (o *TestsPutTest) HasExternalAuditorStatusOptionId() bool`

HasExternalAuditorStatusOptionId returns a boolean if a field has been set.

### GetReviewerDueDate

`func (o *TestsPutTest) GetReviewerDueDate() string`

GetReviewerDueDate returns the ReviewerDueDate field if non-nil, zero value otherwise.

### GetReviewerDueDateOk

`func (o *TestsPutTest) GetReviewerDueDateOk() (*string, bool)`

GetReviewerDueDateOk returns a tuple with the ReviewerDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerDueDate

`func (o *TestsPutTest) SetReviewerDueDate(v string)`

SetReviewerDueDate sets ReviewerDueDate field to given value.

### HasReviewerDueDate

`func (o *TestsPutTest) HasReviewerDueDate() bool`

HasReviewerDueDate returns a boolean if a field has been set.

### GetOpenDate

`func (o *TestsPutTest) GetOpenDate() string`

GetOpenDate returns the OpenDate field if non-nil, zero value otherwise.

### GetOpenDateOk

`func (o *TestsPutTest) GetOpenDateOk() (*string, bool)`

GetOpenDateOk returns a tuple with the OpenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenDate

`func (o *TestsPutTest) SetOpenDate(v string)`

SetOpenDate sets OpenDate field to given value.

### HasOpenDate

`func (o *TestsPutTest) HasOpenDate() bool`

HasOpenDate returns a boolean if a field has been set.

### GetUnderReviewTDate

`func (o *TestsPutTest) GetUnderReviewTDate() string`

GetUnderReviewTDate returns the UnderReviewTDate field if non-nil, zero value otherwise.

### GetUnderReviewTDateOk

`func (o *TestsPutTest) GetUnderReviewTDateOk() (*string, bool)`

GetUnderReviewTDateOk returns a tuple with the UnderReviewTDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderReviewTDate

`func (o *TestsPutTest) SetUnderReviewTDate(v string)`

SetUnderReviewTDate sets UnderReviewTDate field to given value.

### HasUnderReviewTDate

`func (o *TestsPutTest) HasUnderReviewTDate() bool`

HasUnderReviewTDate returns a boolean if a field has been set.

### GetUnderReviewR2Date

`func (o *TestsPutTest) GetUnderReviewR2Date() string`

GetUnderReviewR2Date returns the UnderReviewR2Date field if non-nil, zero value otherwise.

### GetUnderReviewR2DateOk

`func (o *TestsPutTest) GetUnderReviewR2DateOk() (*string, bool)`

GetUnderReviewR2DateOk returns a tuple with the UnderReviewR2Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderReviewR2Date

`func (o *TestsPutTest) SetUnderReviewR2Date(v string)`

SetUnderReviewR2Date sets UnderReviewR2Date field to given value.

### HasUnderReviewR2Date

`func (o *TestsPutTest) HasUnderReviewR2Date() bool`

HasUnderReviewR2Date returns a boolean if a field has been set.

### GetOpenByUserId

`func (o *TestsPutTest) GetOpenByUserId() int32`

GetOpenByUserId returns the OpenByUserId field if non-nil, zero value otherwise.

### GetOpenByUserIdOk

`func (o *TestsPutTest) GetOpenByUserIdOk() (*int32, bool)`

GetOpenByUserIdOk returns a tuple with the OpenByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenByUserId

`func (o *TestsPutTest) SetOpenByUserId(v int32)`

SetOpenByUserId sets OpenByUserId field to given value.

### HasOpenByUserId

`func (o *TestsPutTest) HasOpenByUserId() bool`

HasOpenByUserId returns a boolean if a field has been set.

### GetUnderReviewTByUserId

`func (o *TestsPutTest) GetUnderReviewTByUserId() int32`

GetUnderReviewTByUserId returns the UnderReviewTByUserId field if non-nil, zero value otherwise.

### GetUnderReviewTByUserIdOk

`func (o *TestsPutTest) GetUnderReviewTByUserIdOk() (*int32, bool)`

GetUnderReviewTByUserIdOk returns a tuple with the UnderReviewTByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderReviewTByUserId

`func (o *TestsPutTest) SetUnderReviewTByUserId(v int32)`

SetUnderReviewTByUserId sets UnderReviewTByUserId field to given value.

### HasUnderReviewTByUserId

`func (o *TestsPutTest) HasUnderReviewTByUserId() bool`

HasUnderReviewTByUserId returns a boolean if a field has been set.

### GetUnderReviewR2ByUserId

`func (o *TestsPutTest) GetUnderReviewR2ByUserId() int32`

GetUnderReviewR2ByUserId returns the UnderReviewR2ByUserId field if non-nil, zero value otherwise.

### GetUnderReviewR2ByUserIdOk

`func (o *TestsPutTest) GetUnderReviewR2ByUserIdOk() (*int32, bool)`

GetUnderReviewR2ByUserIdOk returns a tuple with the UnderReviewR2ByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderReviewR2ByUserId

`func (o *TestsPutTest) SetUnderReviewR2ByUserId(v int32)`

SetUnderReviewR2ByUserId sets UnderReviewR2ByUserId field to given value.

### HasUnderReviewR2ByUserId

`func (o *TestsPutTest) HasUnderReviewR2ByUserId() bool`

HasUnderReviewR2ByUserId returns a boolean if a field has been set.

### GetTestCustomSelect1OptionId

`func (o *TestsPutTest) GetTestCustomSelect1OptionId() int32`

GetTestCustomSelect1OptionId returns the TestCustomSelect1OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect1OptionIdOk

`func (o *TestsPutTest) GetTestCustomSelect1OptionIdOk() (*int32, bool)`

GetTestCustomSelect1OptionIdOk returns a tuple with the TestCustomSelect1OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect1OptionId

`func (o *TestsPutTest) SetTestCustomSelect1OptionId(v int32)`

SetTestCustomSelect1OptionId sets TestCustomSelect1OptionId field to given value.

### HasTestCustomSelect1OptionId

`func (o *TestsPutTest) HasTestCustomSelect1OptionId() bool`

HasTestCustomSelect1OptionId returns a boolean if a field has been set.

### GetTestCustomSelect2OptionId

`func (o *TestsPutTest) GetTestCustomSelect2OptionId() int32`

GetTestCustomSelect2OptionId returns the TestCustomSelect2OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect2OptionIdOk

`func (o *TestsPutTest) GetTestCustomSelect2OptionIdOk() (*int32, bool)`

GetTestCustomSelect2OptionIdOk returns a tuple with the TestCustomSelect2OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect2OptionId

`func (o *TestsPutTest) SetTestCustomSelect2OptionId(v int32)`

SetTestCustomSelect2OptionId sets TestCustomSelect2OptionId field to given value.

### HasTestCustomSelect2OptionId

`func (o *TestsPutTest) HasTestCustomSelect2OptionId() bool`

HasTestCustomSelect2OptionId returns a boolean if a field has been set.

### GetTestCustomSelect3OptionId

`func (o *TestsPutTest) GetTestCustomSelect3OptionId() int32`

GetTestCustomSelect3OptionId returns the TestCustomSelect3OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect3OptionIdOk

`func (o *TestsPutTest) GetTestCustomSelect3OptionIdOk() (*int32, bool)`

GetTestCustomSelect3OptionIdOk returns a tuple with the TestCustomSelect3OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect3OptionId

`func (o *TestsPutTest) SetTestCustomSelect3OptionId(v int32)`

SetTestCustomSelect3OptionId sets TestCustomSelect3OptionId field to given value.

### HasTestCustomSelect3OptionId

`func (o *TestsPutTest) HasTestCustomSelect3OptionId() bool`

HasTestCustomSelect3OptionId returns a boolean if a field has been set.

### GetInformationReliability

`func (o *TestsPutTest) GetInformationReliability() string`

GetInformationReliability returns the InformationReliability field if non-nil, zero value otherwise.

### GetInformationReliabilityOk

`func (o *TestsPutTest) GetInformationReliabilityOk() (*string, bool)`

GetInformationReliabilityOk returns a tuple with the InformationReliability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformationReliability

`func (o *TestsPutTest) SetInformationReliability(v string)`

SetInformationReliability sets InformationReliability field to given value.

### HasInformationReliability

`func (o *TestsPutTest) HasInformationReliability() bool`

HasInformationReliability returns a boolean if a field has been set.

### GetControlCharacteristics

`func (o *TestsPutTest) GetControlCharacteristics() string`

GetControlCharacteristics returns the ControlCharacteristics field if non-nil, zero value otherwise.

### GetControlCharacteristicsOk

`func (o *TestsPutTest) GetControlCharacteristicsOk() (*string, bool)`

GetControlCharacteristicsOk returns a tuple with the ControlCharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCharacteristics

`func (o *TestsPutTest) SetControlCharacteristics(v string)`

SetControlCharacteristics sets ControlCharacteristics field to given value.

### HasControlCharacteristics

`func (o *TestsPutTest) HasControlCharacteristics() bool`

HasControlCharacteristics returns a boolean if a field has been set.

### GetExceptionNature

`func (o *TestsPutTest) GetExceptionNature() string`

GetExceptionNature returns the ExceptionNature field if non-nil, zero value otherwise.

### GetExceptionNatureOk

`func (o *TestsPutTest) GetExceptionNatureOk() (*string, bool)`

GetExceptionNatureOk returns a tuple with the ExceptionNature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionNature

`func (o *TestsPutTest) SetExceptionNature(v string)`

SetExceptionNature sets ExceptionNature field to given value.

### HasExceptionNature

`func (o *TestsPutTest) HasExceptionNature() bool`

HasExceptionNature returns a boolean if a field has been set.

### GetSampleSizeAdditional

`func (o *TestsPutTest) GetSampleSizeAdditional() string`

GetSampleSizeAdditional returns the SampleSizeAdditional field if non-nil, zero value otherwise.

### GetSampleSizeAdditionalOk

`func (o *TestsPutTest) GetSampleSizeAdditionalOk() (*string, bool)`

GetSampleSizeAdditionalOk returns a tuple with the SampleSizeAdditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleSizeAdditional

`func (o *TestsPutTest) SetSampleSizeAdditional(v string)`

SetSampleSizeAdditional sets SampleSizeAdditional field to given value.

### HasSampleSizeAdditional

`func (o *TestsPutTest) HasSampleSizeAdditional() bool`

HasSampleSizeAdditional returns a boolean if a field has been set.

### GetNumberExceptions

`func (o *TestsPutTest) GetNumberExceptions() string`

GetNumberExceptions returns the NumberExceptions field if non-nil, zero value otherwise.

### GetNumberExceptionsOk

`func (o *TestsPutTest) GetNumberExceptionsOk() (*string, bool)`

GetNumberExceptionsOk returns a tuple with the NumberExceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberExceptions

`func (o *TestsPutTest) SetNumberExceptions(v string)`

SetNumberExceptions sets NumberExceptions field to given value.

### HasNumberExceptions

`func (o *TestsPutTest) HasNumberExceptions() bool`

HasNumberExceptions returns a boolean if a field has been set.

### GetSecondaryReviewerDueDate

`func (o *TestsPutTest) GetSecondaryReviewerDueDate() string`

GetSecondaryReviewerDueDate returns the SecondaryReviewerDueDate field if non-nil, zero value otherwise.

### GetSecondaryReviewerDueDateOk

`func (o *TestsPutTest) GetSecondaryReviewerDueDateOk() (*string, bool)`

GetSecondaryReviewerDueDateOk returns a tuple with the SecondaryReviewerDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryReviewerDueDate

`func (o *TestsPutTest) SetSecondaryReviewerDueDate(v string)`

SetSecondaryReviewerDueDate sets SecondaryReviewerDueDate field to given value.

### HasSecondaryReviewerDueDate

`func (o *TestsPutTest) HasSecondaryReviewerDueDate() bool`

HasSecondaryReviewerDueDate returns a boolean if a field has been set.

### GetArchiveId

`func (o *TestsPutTest) GetArchiveId() int32`

GetArchiveId returns the ArchiveId field if non-nil, zero value otherwise.

### GetArchiveIdOk

`func (o *TestsPutTest) GetArchiveIdOk() (*int32, bool)`

GetArchiveIdOk returns a tuple with the ArchiveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveId

`func (o *TestsPutTest) SetArchiveId(v int32)`

SetArchiveId sets ArchiveId field to given value.

### HasArchiveId

`func (o *TestsPutTest) HasArchiveId() bool`

HasArchiveId returns a boolean if a field has been set.

### GetTestCustomSelect4OptionId

`func (o *TestsPutTest) GetTestCustomSelect4OptionId() int32`

GetTestCustomSelect4OptionId returns the TestCustomSelect4OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect4OptionIdOk

`func (o *TestsPutTest) GetTestCustomSelect4OptionIdOk() (*int32, bool)`

GetTestCustomSelect4OptionIdOk returns a tuple with the TestCustomSelect4OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect4OptionId

`func (o *TestsPutTest) SetTestCustomSelect4OptionId(v int32)`

SetTestCustomSelect4OptionId sets TestCustomSelect4OptionId field to given value.

### HasTestCustomSelect4OptionId

`func (o *TestsPutTest) HasTestCustomSelect4OptionId() bool`

HasTestCustomSelect4OptionId returns a boolean if a field has been set.

### GetFirstCompleteDate

`func (o *TestsPutTest) GetFirstCompleteDate() string`

GetFirstCompleteDate returns the FirstCompleteDate field if non-nil, zero value otherwise.

### GetFirstCompleteDateOk

`func (o *TestsPutTest) GetFirstCompleteDateOk() (*string, bool)`

GetFirstCompleteDateOk returns a tuple with the FirstCompleteDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCompleteDate

`func (o *TestsPutTest) SetFirstCompleteDate(v string)`

SetFirstCompleteDate sets FirstCompleteDate field to given value.

### HasFirstCompleteDate

`func (o *TestsPutTest) HasFirstCompleteDate() bool`

HasFirstCompleteDate returns a boolean if a field has been set.

### GetFirstReviewedDate

`func (o *TestsPutTest) GetFirstReviewedDate() string`

GetFirstReviewedDate returns the FirstReviewedDate field if non-nil, zero value otherwise.

### GetFirstReviewedDateOk

`func (o *TestsPutTest) GetFirstReviewedDateOk() (*string, bool)`

GetFirstReviewedDateOk returns a tuple with the FirstReviewedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstReviewedDate

`func (o *TestsPutTest) SetFirstReviewedDate(v string)`

SetFirstReviewedDate sets FirstReviewedDate field to given value.

### HasFirstReviewedDate

`func (o *TestsPutTest) HasFirstReviewedDate() bool`

HasFirstReviewedDate returns a boolean if a field has been set.

### GetFirstUnderReviewTDate

`func (o *TestsPutTest) GetFirstUnderReviewTDate() string`

GetFirstUnderReviewTDate returns the FirstUnderReviewTDate field if non-nil, zero value otherwise.

### GetFirstUnderReviewTDateOk

`func (o *TestsPutTest) GetFirstUnderReviewTDateOk() (*string, bool)`

GetFirstUnderReviewTDateOk returns a tuple with the FirstUnderReviewTDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstUnderReviewTDate

`func (o *TestsPutTest) SetFirstUnderReviewTDate(v string)`

SetFirstUnderReviewTDate sets FirstUnderReviewTDate field to given value.

### HasFirstUnderReviewTDate

`func (o *TestsPutTest) HasFirstUnderReviewTDate() bool`

HasFirstUnderReviewTDate returns a boolean if a field has been set.

### GetFirstUnderReviewR2Date

`func (o *TestsPutTest) GetFirstUnderReviewR2Date() string`

GetFirstUnderReviewR2Date returns the FirstUnderReviewR2Date field if non-nil, zero value otherwise.

### GetFirstUnderReviewR2DateOk

`func (o *TestsPutTest) GetFirstUnderReviewR2DateOk() (*string, bool)`

GetFirstUnderReviewR2DateOk returns a tuple with the FirstUnderReviewR2Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstUnderReviewR2Date

`func (o *TestsPutTest) SetFirstUnderReviewR2Date(v string)`

SetFirstUnderReviewR2Date sets FirstUnderReviewR2Date field to given value.

### HasFirstUnderReviewR2Date

`func (o *TestsPutTest) HasFirstUnderReviewR2Date() bool`

HasFirstUnderReviewR2Date returns a boolean if a field has been set.

### GetCustomText1

`func (o *TestsPutTest) GetCustomText1() string`

GetCustomText1 returns the CustomText1 field if non-nil, zero value otherwise.

### GetCustomText1Ok

`func (o *TestsPutTest) GetCustomText1Ok() (*string, bool)`

GetCustomText1Ok returns a tuple with the CustomText1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText1

`func (o *TestsPutTest) SetCustomText1(v string)`

SetCustomText1 sets CustomText1 field to given value.

### HasCustomText1

`func (o *TestsPutTest) HasCustomText1() bool`

HasCustomText1 returns a boolean if a field has been set.

### GetCustomText2

`func (o *TestsPutTest) GetCustomText2() string`

GetCustomText2 returns the CustomText2 field if non-nil, zero value otherwise.

### GetCustomText2Ok

`func (o *TestsPutTest) GetCustomText2Ok() (*string, bool)`

GetCustomText2Ok returns a tuple with the CustomText2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText2

`func (o *TestsPutTest) SetCustomText2(v string)`

SetCustomText2 sets CustomText2 field to given value.

### HasCustomText2

`func (o *TestsPutTest) HasCustomText2() bool`

HasCustomText2 returns a boolean if a field has been set.

### GetCustomText3

`func (o *TestsPutTest) GetCustomText3() string`

GetCustomText3 returns the CustomText3 field if non-nil, zero value otherwise.

### GetCustomText3Ok

`func (o *TestsPutTest) GetCustomText3Ok() (*string, bool)`

GetCustomText3Ok returns a tuple with the CustomText3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText3

`func (o *TestsPutTest) SetCustomText3(v string)`

SetCustomText3 sets CustomText3 field to given value.

### HasCustomText3

`func (o *TestsPutTest) HasCustomText3() bool`

HasCustomText3 returns a boolean if a field has been set.

### GetCustomText4

`func (o *TestsPutTest) GetCustomText4() string`

GetCustomText4 returns the CustomText4 field if non-nil, zero value otherwise.

### GetCustomText4Ok

`func (o *TestsPutTest) GetCustomText4Ok() (*string, bool)`

GetCustomText4Ok returns a tuple with the CustomText4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText4

`func (o *TestsPutTest) SetCustomText4(v string)`

SetCustomText4 sets CustomText4 field to given value.

### HasCustomText4

`func (o *TestsPutTest) HasCustomText4() bool`

HasCustomText4 returns a boolean if a field has been set.

### GetCustomText5

`func (o *TestsPutTest) GetCustomText5() string`

GetCustomText5 returns the CustomText5 field if non-nil, zero value otherwise.

### GetCustomText5Ok

`func (o *TestsPutTest) GetCustomText5Ok() (*string, bool)`

GetCustomText5Ok returns a tuple with the CustomText5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText5

`func (o *TestsPutTest) SetCustomText5(v string)`

SetCustomText5 sets CustomText5 field to given value.

### HasCustomText5

`func (o *TestsPutTest) HasCustomText5() bool`

HasCustomText5 returns a boolean if a field has been set.

### GetCustomText6

`func (o *TestsPutTest) GetCustomText6() string`

GetCustomText6 returns the CustomText6 field if non-nil, zero value otherwise.

### GetCustomText6Ok

`func (o *TestsPutTest) GetCustomText6Ok() (*string, bool)`

GetCustomText6Ok returns a tuple with the CustomText6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText6

`func (o *TestsPutTest) SetCustomText6(v string)`

SetCustomText6 sets CustomText6 field to given value.

### HasCustomText6

`func (o *TestsPutTest) HasCustomText6() bool`

HasCustomText6 returns a boolean if a field has been set.

### GetCustomText7

`func (o *TestsPutTest) GetCustomText7() string`

GetCustomText7 returns the CustomText7 field if non-nil, zero value otherwise.

### GetCustomText7Ok

`func (o *TestsPutTest) GetCustomText7Ok() (*string, bool)`

GetCustomText7Ok returns a tuple with the CustomText7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText7

`func (o *TestsPutTest) SetCustomText7(v string)`

SetCustomText7 sets CustomText7 field to given value.

### HasCustomText7

`func (o *TestsPutTest) HasCustomText7() bool`

HasCustomText7 returns a boolean if a field has been set.

### GetCustomText8

`func (o *TestsPutTest) GetCustomText8() string`

GetCustomText8 returns the CustomText8 field if non-nil, zero value otherwise.

### GetCustomText8Ok

`func (o *TestsPutTest) GetCustomText8Ok() (*string, bool)`

GetCustomText8Ok returns a tuple with the CustomText8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText8

`func (o *TestsPutTest) SetCustomText8(v string)`

SetCustomText8 sets CustomText8 field to given value.

### HasCustomText8

`func (o *TestsPutTest) HasCustomText8() bool`

HasCustomText8 returns a boolean if a field has been set.

### GetHoldByUserId

`func (o *TestsPutTest) GetHoldByUserId() int32`

GetHoldByUserId returns the HoldByUserId field if non-nil, zero value otherwise.

### GetHoldByUserIdOk

`func (o *TestsPutTest) GetHoldByUserIdOk() (*int32, bool)`

GetHoldByUserIdOk returns a tuple with the HoldByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldByUserId

`func (o *TestsPutTest) SetHoldByUserId(v int32)`

SetHoldByUserId sets HoldByUserId field to given value.

### HasHoldByUserId

`func (o *TestsPutTest) HasHoldByUserId() bool`

HasHoldByUserId returns a boolean if a field has been set.

### GetHoldDate

`func (o *TestsPutTest) GetHoldDate() string`

GetHoldDate returns the HoldDate field if non-nil, zero value otherwise.

### GetHoldDateOk

`func (o *TestsPutTest) GetHoldDateOk() (*string, bool)`

GetHoldDateOk returns a tuple with the HoldDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldDate

`func (o *TestsPutTest) SetHoldDate(v string)`

SetHoldDate sets HoldDate field to given value.

### HasHoldDate

`func (o *TestsPutTest) HasHoldDate() bool`

HasHoldDate returns a boolean if a field has been set.

### GetFirstHoldDate

`func (o *TestsPutTest) GetFirstHoldDate() string`

GetFirstHoldDate returns the FirstHoldDate field if non-nil, zero value otherwise.

### GetFirstHoldDateOk

`func (o *TestsPutTest) GetFirstHoldDateOk() (*string, bool)`

GetFirstHoldDateOk returns a tuple with the FirstHoldDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstHoldDate

`func (o *TestsPutTest) SetFirstHoldDate(v string)`

SetFirstHoldDate sets FirstHoldDate field to given value.

### HasFirstHoldDate

`func (o *TestsPutTest) HasFirstHoldDate() bool`

HasFirstHoldDate returns a boolean if a field has been set.

### GetHoldStatusOptionId

`func (o *TestsPutTest) GetHoldStatusOptionId() int32`

GetHoldStatusOptionId returns the HoldStatusOptionId field if non-nil, zero value otherwise.

### GetHoldStatusOptionIdOk

`func (o *TestsPutTest) GetHoldStatusOptionIdOk() (*int32, bool)`

GetHoldStatusOptionIdOk returns a tuple with the HoldStatusOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldStatusOptionId

`func (o *TestsPutTest) SetHoldStatusOptionId(v int32)`

SetHoldStatusOptionId sets HoldStatusOptionId field to given value.

### HasHoldStatusOptionId

`func (o *TestsPutTest) HasHoldStatusOptionId() bool`

HasHoldStatusOptionId returns a boolean if a field has been set.

### GetTestCustomSelect5OptionId

`func (o *TestsPutTest) GetTestCustomSelect5OptionId() int32`

GetTestCustomSelect5OptionId returns the TestCustomSelect5OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect5OptionIdOk

`func (o *TestsPutTest) GetTestCustomSelect5OptionIdOk() (*int32, bool)`

GetTestCustomSelect5OptionIdOk returns a tuple with the TestCustomSelect5OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect5OptionId

`func (o *TestsPutTest) SetTestCustomSelect5OptionId(v int32)`

SetTestCustomSelect5OptionId sets TestCustomSelect5OptionId field to given value.

### HasTestCustomSelect5OptionId

`func (o *TestsPutTest) HasTestCustomSelect5OptionId() bool`

HasTestCustomSelect5OptionId returns a boolean if a field has been set.

### GetTestCustomSelect6OptionId

`func (o *TestsPutTest) GetTestCustomSelect6OptionId() int32`

GetTestCustomSelect6OptionId returns the TestCustomSelect6OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect6OptionIdOk

`func (o *TestsPutTest) GetTestCustomSelect6OptionIdOk() (*int32, bool)`

GetTestCustomSelect6OptionIdOk returns a tuple with the TestCustomSelect6OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect6OptionId

`func (o *TestsPutTest) SetTestCustomSelect6OptionId(v int32)`

SetTestCustomSelect6OptionId sets TestCustomSelect6OptionId field to given value.

### HasTestCustomSelect6OptionId

`func (o *TestsPutTest) HasTestCustomSelect6OptionId() bool`

HasTestCustomSelect6OptionId returns a boolean if a field has been set.

### GetTestCustomSelect7OptionId

`func (o *TestsPutTest) GetTestCustomSelect7OptionId() int32`

GetTestCustomSelect7OptionId returns the TestCustomSelect7OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect7OptionIdOk

`func (o *TestsPutTest) GetTestCustomSelect7OptionIdOk() (*int32, bool)`

GetTestCustomSelect7OptionIdOk returns a tuple with the TestCustomSelect7OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect7OptionId

`func (o *TestsPutTest) SetTestCustomSelect7OptionId(v int32)`

SetTestCustomSelect7OptionId sets TestCustomSelect7OptionId field to given value.

### HasTestCustomSelect7OptionId

`func (o *TestsPutTest) HasTestCustomSelect7OptionId() bool`

HasTestCustomSelect7OptionId returns a boolean if a field has been set.

### GetBudgetHours

`func (o *TestsPutTest) GetBudgetHours() float32`

GetBudgetHours returns the BudgetHours field if non-nil, zero value otherwise.

### GetBudgetHoursOk

`func (o *TestsPutTest) GetBudgetHoursOk() (*float32, bool)`

GetBudgetHoursOk returns a tuple with the BudgetHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetHours

`func (o *TestsPutTest) SetBudgetHours(v float32)`

SetBudgetHours sets BudgetHours field to given value.

### HasBudgetHours

`func (o *TestsPutTest) HasBudgetHours() bool`

HasBudgetHours returns a boolean if a field has been set.

### GetHours

`func (o *TestsPutTest) GetHours() float32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *TestsPutTest) GetHoursOk() (*float32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *TestsPutTest) SetHours(v float32)`

SetHours sets Hours field to given value.

### HasHours

`func (o *TestsPutTest) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetReviewerBudgetHours

`func (o *TestsPutTest) GetReviewerBudgetHours() float32`

GetReviewerBudgetHours returns the ReviewerBudgetHours field if non-nil, zero value otherwise.

### GetReviewerBudgetHoursOk

`func (o *TestsPutTest) GetReviewerBudgetHoursOk() (*float32, bool)`

GetReviewerBudgetHoursOk returns a tuple with the ReviewerBudgetHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerBudgetHours

`func (o *TestsPutTest) SetReviewerBudgetHours(v float32)`

SetReviewerBudgetHours sets ReviewerBudgetHours field to given value.

### HasReviewerBudgetHours

`func (o *TestsPutTest) HasReviewerBudgetHours() bool`

HasReviewerBudgetHours returns a boolean if a field has been set.

### GetReviewerHours

`func (o *TestsPutTest) GetReviewerHours() float32`

GetReviewerHours returns the ReviewerHours field if non-nil, zero value otherwise.

### GetReviewerHoursOk

`func (o *TestsPutTest) GetReviewerHoursOk() (*float32, bool)`

GetReviewerHoursOk returns a tuple with the ReviewerHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerHours

`func (o *TestsPutTest) SetReviewerHours(v float32)`

SetReviewerHours sets ReviewerHours field to given value.

### HasReviewerHours

`func (o *TestsPutTest) HasReviewerHours() bool`

HasReviewerHours returns a boolean if a field has been set.

### GetPercentComplete

`func (o *TestsPutTest) GetPercentComplete() float32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *TestsPutTest) GetPercentCompleteOk() (*float32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *TestsPutTest) SetPercentComplete(v float32)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *TestsPutTest) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### GetCustomText9

`func (o *TestsPutTest) GetCustomText9() string`

GetCustomText9 returns the CustomText9 field if non-nil, zero value otherwise.

### GetCustomText9Ok

`func (o *TestsPutTest) GetCustomText9Ok() (*string, bool)`

GetCustomText9Ok returns a tuple with the CustomText9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText9

`func (o *TestsPutTest) SetCustomText9(v string)`

SetCustomText9 sets CustomText9 field to given value.

### HasCustomText9

`func (o *TestsPutTest) HasCustomText9() bool`

HasCustomText9 returns a boolean if a field has been set.

### GetCustomText10

`func (o *TestsPutTest) GetCustomText10() string`

GetCustomText10 returns the CustomText10 field if non-nil, zero value otherwise.

### GetCustomText10Ok

`func (o *TestsPutTest) GetCustomText10Ok() (*string, bool)`

GetCustomText10Ok returns a tuple with the CustomText10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText10

`func (o *TestsPutTest) SetCustomText10(v string)`

SetCustomText10 sets CustomText10 field to given value.

### HasCustomText10

`func (o *TestsPutTest) HasCustomText10() bool`

HasCustomText10 returns a boolean if a field has been set.

### GetCustomText11

`func (o *TestsPutTest) GetCustomText11() string`

GetCustomText11 returns the CustomText11 field if non-nil, zero value otherwise.

### GetCustomText11Ok

`func (o *TestsPutTest) GetCustomText11Ok() (*string, bool)`

GetCustomText11Ok returns a tuple with the CustomText11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText11

`func (o *TestsPutTest) SetCustomText11(v string)`

SetCustomText11 sets CustomText11 field to given value.

### HasCustomText11

`func (o *TestsPutTest) HasCustomText11() bool`

HasCustomText11 returns a boolean if a field has been set.

### GetCustomText12

`func (o *TestsPutTest) GetCustomText12() string`

GetCustomText12 returns the CustomText12 field if non-nil, zero value otherwise.

### GetCustomText12Ok

`func (o *TestsPutTest) GetCustomText12Ok() (*string, bool)`

GetCustomText12Ok returns a tuple with the CustomText12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText12

`func (o *TestsPutTest) SetCustomText12(v string)`

SetCustomText12 sets CustomText12 field to given value.

### HasCustomText12

`func (o *TestsPutTest) HasCustomText12() bool`

HasCustomText12 returns a boolean if a field has been set.

### GetCustomText13

`func (o *TestsPutTest) GetCustomText13() string`

GetCustomText13 returns the CustomText13 field if non-nil, zero value otherwise.

### GetCustomText13Ok

`func (o *TestsPutTest) GetCustomText13Ok() (*string, bool)`

GetCustomText13Ok returns a tuple with the CustomText13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText13

`func (o *TestsPutTest) SetCustomText13(v string)`

SetCustomText13 sets CustomText13 field to given value.

### HasCustomText13

`func (o *TestsPutTest) HasCustomText13() bool`

HasCustomText13 returns a boolean if a field has been set.

### GetCustomText14

`func (o *TestsPutTest) GetCustomText14() string`

GetCustomText14 returns the CustomText14 field if non-nil, zero value otherwise.

### GetCustomText14Ok

`func (o *TestsPutTest) GetCustomText14Ok() (*string, bool)`

GetCustomText14Ok returns a tuple with the CustomText14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText14

`func (o *TestsPutTest) SetCustomText14(v string)`

SetCustomText14 sets CustomText14 field to given value.

### HasCustomText14

`func (o *TestsPutTest) HasCustomText14() bool`

HasCustomText14 returns a boolean if a field has been set.

### GetCustomText15

`func (o *TestsPutTest) GetCustomText15() string`

GetCustomText15 returns the CustomText15 field if non-nil, zero value otherwise.

### GetCustomText15Ok

`func (o *TestsPutTest) GetCustomText15Ok() (*string, bool)`

GetCustomText15Ok returns a tuple with the CustomText15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText15

`func (o *TestsPutTest) SetCustomText15(v string)`

SetCustomText15 sets CustomText15 field to given value.

### HasCustomText15

`func (o *TestsPutTest) HasCustomText15() bool`

HasCustomText15 returns a boolean if a field has been set.

### GetCustomText16

`func (o *TestsPutTest) GetCustomText16() string`

GetCustomText16 returns the CustomText16 field if non-nil, zero value otherwise.

### GetCustomText16Ok

`func (o *TestsPutTest) GetCustomText16Ok() (*string, bool)`

GetCustomText16Ok returns a tuple with the CustomText16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText16

`func (o *TestsPutTest) SetCustomText16(v string)`

SetCustomText16 sets CustomText16 field to given value.

### HasCustomText16

`func (o *TestsPutTest) HasCustomText16() bool`

HasCustomText16 returns a boolean if a field has been set.

### GetCustomText17

`func (o *TestsPutTest) GetCustomText17() string`

GetCustomText17 returns the CustomText17 field if non-nil, zero value otherwise.

### GetCustomText17Ok

`func (o *TestsPutTest) GetCustomText17Ok() (*string, bool)`

GetCustomText17Ok returns a tuple with the CustomText17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText17

`func (o *TestsPutTest) SetCustomText17(v string)`

SetCustomText17 sets CustomText17 field to given value.

### HasCustomText17

`func (o *TestsPutTest) HasCustomText17() bool`

HasCustomText17 returns a boolean if a field has been set.

### GetCustomText18

`func (o *TestsPutTest) GetCustomText18() string`

GetCustomText18 returns the CustomText18 field if non-nil, zero value otherwise.

### GetCustomText18Ok

`func (o *TestsPutTest) GetCustomText18Ok() (*string, bool)`

GetCustomText18Ok returns a tuple with the CustomText18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText18

`func (o *TestsPutTest) SetCustomText18(v string)`

SetCustomText18 sets CustomText18 field to given value.

### HasCustomText18

`func (o *TestsPutTest) HasCustomText18() bool`

HasCustomText18 returns a boolean if a field has been set.

### GetCustomText19

`func (o *TestsPutTest) GetCustomText19() string`

GetCustomText19 returns the CustomText19 field if non-nil, zero value otherwise.

### GetCustomText19Ok

`func (o *TestsPutTest) GetCustomText19Ok() (*string, bool)`

GetCustomText19Ok returns a tuple with the CustomText19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText19

`func (o *TestsPutTest) SetCustomText19(v string)`

SetCustomText19 sets CustomText19 field to given value.

### HasCustomText19

`func (o *TestsPutTest) HasCustomText19() bool`

HasCustomText19 returns a boolean if a field has been set.

### GetCustomText20

`func (o *TestsPutTest) GetCustomText20() string`

GetCustomText20 returns the CustomText20 field if non-nil, zero value otherwise.

### GetCustomText20Ok

`func (o *TestsPutTest) GetCustomText20Ok() (*string, bool)`

GetCustomText20Ok returns a tuple with the CustomText20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText20

`func (o *TestsPutTest) SetCustomText20(v string)`

SetCustomText20 sets CustomText20 field to given value.

### HasCustomText20

`func (o *TestsPutTest) HasCustomText20() bool`

HasCustomText20 returns a boolean if a field has been set.

### GetCustomText21

`func (o *TestsPutTest) GetCustomText21() string`

GetCustomText21 returns the CustomText21 field if non-nil, zero value otherwise.

### GetCustomText21Ok

`func (o *TestsPutTest) GetCustomText21Ok() (*string, bool)`

GetCustomText21Ok returns a tuple with the CustomText21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText21

`func (o *TestsPutTest) SetCustomText21(v string)`

SetCustomText21 sets CustomText21 field to given value.

### HasCustomText21

`func (o *TestsPutTest) HasCustomText21() bool`

HasCustomText21 returns a boolean if a field has been set.

### GetCustomText22

`func (o *TestsPutTest) GetCustomText22() string`

GetCustomText22 returns the CustomText22 field if non-nil, zero value otherwise.

### GetCustomText22Ok

`func (o *TestsPutTest) GetCustomText22Ok() (*string, bool)`

GetCustomText22Ok returns a tuple with the CustomText22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText22

`func (o *TestsPutTest) SetCustomText22(v string)`

SetCustomText22 sets CustomText22 field to given value.

### HasCustomText22

`func (o *TestsPutTest) HasCustomText22() bool`

HasCustomText22 returns a boolean if a field has been set.

### GetCustomText23

`func (o *TestsPutTest) GetCustomText23() string`

GetCustomText23 returns the CustomText23 field if non-nil, zero value otherwise.

### GetCustomText23Ok

`func (o *TestsPutTest) GetCustomText23Ok() (*string, bool)`

GetCustomText23Ok returns a tuple with the CustomText23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText23

`func (o *TestsPutTest) SetCustomText23(v string)`

SetCustomText23 sets CustomText23 field to given value.

### HasCustomText23

`func (o *TestsPutTest) HasCustomText23() bool`

HasCustomText23 returns a boolean if a field has been set.

### GetCustomText24

`func (o *TestsPutTest) GetCustomText24() string`

GetCustomText24 returns the CustomText24 field if non-nil, zero value otherwise.

### GetCustomText24Ok

`func (o *TestsPutTest) GetCustomText24Ok() (*string, bool)`

GetCustomText24Ok returns a tuple with the CustomText24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText24

`func (o *TestsPutTest) SetCustomText24(v string)`

SetCustomText24 sets CustomText24 field to given value.

### HasCustomText24

`func (o *TestsPutTest) HasCustomText24() bool`

HasCustomText24 returns a boolean if a field has been set.

### GetCustomText25

`func (o *TestsPutTest) GetCustomText25() string`

GetCustomText25 returns the CustomText25 field if non-nil, zero value otherwise.

### GetCustomText25Ok

`func (o *TestsPutTest) GetCustomText25Ok() (*string, bool)`

GetCustomText25Ok returns a tuple with the CustomText25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText25

`func (o *TestsPutTest) SetCustomText25(v string)`

SetCustomText25 sets CustomText25 field to given value.

### HasCustomText25

`func (o *TestsPutTest) HasCustomText25() bool`

HasCustomText25 returns a boolean if a field has been set.

### GetCustomText26

`func (o *TestsPutTest) GetCustomText26() string`

GetCustomText26 returns the CustomText26 field if non-nil, zero value otherwise.

### GetCustomText26Ok

`func (o *TestsPutTest) GetCustomText26Ok() (*string, bool)`

GetCustomText26Ok returns a tuple with the CustomText26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText26

`func (o *TestsPutTest) SetCustomText26(v string)`

SetCustomText26 sets CustomText26 field to given value.

### HasCustomText26

`func (o *TestsPutTest) HasCustomText26() bool`

HasCustomText26 returns a boolean if a field has been set.

### GetCustomText27

`func (o *TestsPutTest) GetCustomText27() string`

GetCustomText27 returns the CustomText27 field if non-nil, zero value otherwise.

### GetCustomText27Ok

`func (o *TestsPutTest) GetCustomText27Ok() (*string, bool)`

GetCustomText27Ok returns a tuple with the CustomText27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText27

`func (o *TestsPutTest) SetCustomText27(v string)`

SetCustomText27 sets CustomText27 field to given value.

### HasCustomText27

`func (o *TestsPutTest) HasCustomText27() bool`

HasCustomText27 returns a boolean if a field has been set.

### GetCustomText28

`func (o *TestsPutTest) GetCustomText28() string`

GetCustomText28 returns the CustomText28 field if non-nil, zero value otherwise.

### GetCustomText28Ok

`func (o *TestsPutTest) GetCustomText28Ok() (*string, bool)`

GetCustomText28Ok returns a tuple with the CustomText28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText28

`func (o *TestsPutTest) SetCustomText28(v string)`

SetCustomText28 sets CustomText28 field to given value.

### HasCustomText28

`func (o *TestsPutTest) HasCustomText28() bool`

HasCustomText28 returns a boolean if a field has been set.

### GetCustomText29

`func (o *TestsPutTest) GetCustomText29() string`

GetCustomText29 returns the CustomText29 field if non-nil, zero value otherwise.

### GetCustomText29Ok

`func (o *TestsPutTest) GetCustomText29Ok() (*string, bool)`

GetCustomText29Ok returns a tuple with the CustomText29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText29

`func (o *TestsPutTest) SetCustomText29(v string)`

SetCustomText29 sets CustomText29 field to given value.

### HasCustomText29

`func (o *TestsPutTest) HasCustomText29() bool`

HasCustomText29 returns a boolean if a field has been set.

### GetCustomText30

`func (o *TestsPutTest) GetCustomText30() string`

GetCustomText30 returns the CustomText30 field if non-nil, zero value otherwise.

### GetCustomText30Ok

`func (o *TestsPutTest) GetCustomText30Ok() (*string, bool)`

GetCustomText30Ok returns a tuple with the CustomText30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText30

`func (o *TestsPutTest) SetCustomText30(v string)`

SetCustomText30 sets CustomText30 field to given value.

### HasCustomText30

`func (o *TestsPutTest) HasCustomText30() bool`

HasCustomText30 returns a boolean if a field has been set.

### GetCustomText31

`func (o *TestsPutTest) GetCustomText31() string`

GetCustomText31 returns the CustomText31 field if non-nil, zero value otherwise.

### GetCustomText31Ok

`func (o *TestsPutTest) GetCustomText31Ok() (*string, bool)`

GetCustomText31Ok returns a tuple with the CustomText31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText31

`func (o *TestsPutTest) SetCustomText31(v string)`

SetCustomText31 sets CustomText31 field to given value.

### HasCustomText31

`func (o *TestsPutTest) HasCustomText31() bool`

HasCustomText31 returns a boolean if a field has been set.

### GetCustomText32

`func (o *TestsPutTest) GetCustomText32() string`

GetCustomText32 returns the CustomText32 field if non-nil, zero value otherwise.

### GetCustomText32Ok

`func (o *TestsPutTest) GetCustomText32Ok() (*string, bool)`

GetCustomText32Ok returns a tuple with the CustomText32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText32

`func (o *TestsPutTest) SetCustomText32(v string)`

SetCustomText32 sets CustomText32 field to given value.

### HasCustomText32

`func (o *TestsPutTest) HasCustomText32() bool`

HasCustomText32 returns a boolean if a field has been set.

### GetCustomText33

`func (o *TestsPutTest) GetCustomText33() string`

GetCustomText33 returns the CustomText33 field if non-nil, zero value otherwise.

### GetCustomText33Ok

`func (o *TestsPutTest) GetCustomText33Ok() (*string, bool)`

GetCustomText33Ok returns a tuple with the CustomText33 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText33

`func (o *TestsPutTest) SetCustomText33(v string)`

SetCustomText33 sets CustomText33 field to given value.

### HasCustomText33

`func (o *TestsPutTest) HasCustomText33() bool`

HasCustomText33 returns a boolean if a field has been set.

### GetTestCustomSelect8OptionId

`func (o *TestsPutTest) GetTestCustomSelect8OptionId() int32`

GetTestCustomSelect8OptionId returns the TestCustomSelect8OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect8OptionIdOk

`func (o *TestsPutTest) GetTestCustomSelect8OptionIdOk() (*int32, bool)`

GetTestCustomSelect8OptionIdOk returns a tuple with the TestCustomSelect8OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect8OptionId

`func (o *TestsPutTest) SetTestCustomSelect8OptionId(v int32)`

SetTestCustomSelect8OptionId sets TestCustomSelect8OptionId field to given value.

### HasTestCustomSelect8OptionId

`func (o *TestsPutTest) HasTestCustomSelect8OptionId() bool`

HasTestCustomSelect8OptionId returns a boolean if a field has been set.

### GetTestCustomSelect9OptionId

`func (o *TestsPutTest) GetTestCustomSelect9OptionId() int32`

GetTestCustomSelect9OptionId returns the TestCustomSelect9OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect9OptionIdOk

`func (o *TestsPutTest) GetTestCustomSelect9OptionIdOk() (*int32, bool)`

GetTestCustomSelect9OptionIdOk returns a tuple with the TestCustomSelect9OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect9OptionId

`func (o *TestsPutTest) SetTestCustomSelect9OptionId(v int32)`

SetTestCustomSelect9OptionId sets TestCustomSelect9OptionId field to given value.

### HasTestCustomSelect9OptionId

`func (o *TestsPutTest) HasTestCustomSelect9OptionId() bool`

HasTestCustomSelect9OptionId returns a boolean if a field has been set.

### GetTestCustomSelect10OptionId

`func (o *TestsPutTest) GetTestCustomSelect10OptionId() int32`

GetTestCustomSelect10OptionId returns the TestCustomSelect10OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect10OptionIdOk

`func (o *TestsPutTest) GetTestCustomSelect10OptionIdOk() (*int32, bool)`

GetTestCustomSelect10OptionIdOk returns a tuple with the TestCustomSelect10OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect10OptionId

`func (o *TestsPutTest) SetTestCustomSelect10OptionId(v int32)`

SetTestCustomSelect10OptionId sets TestCustomSelect10OptionId field to given value.

### HasTestCustomSelect10OptionId

`func (o *TestsPutTest) HasTestCustomSelect10OptionId() bool`

HasTestCustomSelect10OptionId returns a boolean if a field has been set.

### GetTestCustomSelect11OptionId

`func (o *TestsPutTest) GetTestCustomSelect11OptionId() int32`

GetTestCustomSelect11OptionId returns the TestCustomSelect11OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect11OptionIdOk

`func (o *TestsPutTest) GetTestCustomSelect11OptionIdOk() (*int32, bool)`

GetTestCustomSelect11OptionIdOk returns a tuple with the TestCustomSelect11OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect11OptionId

`func (o *TestsPutTest) SetTestCustomSelect11OptionId(v int32)`

SetTestCustomSelect11OptionId sets TestCustomSelect11OptionId field to given value.

### HasTestCustomSelect11OptionId

`func (o *TestsPutTest) HasTestCustomSelect11OptionId() bool`

HasTestCustomSelect11OptionId returns a boolean if a field has been set.

### GetTestCustomSelect12OptionId

`func (o *TestsPutTest) GetTestCustomSelect12OptionId() int32`

GetTestCustomSelect12OptionId returns the TestCustomSelect12OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect12OptionIdOk

`func (o *TestsPutTest) GetTestCustomSelect12OptionIdOk() (*int32, bool)`

GetTestCustomSelect12OptionIdOk returns a tuple with the TestCustomSelect12OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect12OptionId

`func (o *TestsPutTest) SetTestCustomSelect12OptionId(v int32)`

SetTestCustomSelect12OptionId sets TestCustomSelect12OptionId field to given value.

### HasTestCustomSelect12OptionId

`func (o *TestsPutTest) HasTestCustomSelect12OptionId() bool`

HasTestCustomSelect12OptionId returns a boolean if a field has been set.

### GetTestCustomSelect13OptionId

`func (o *TestsPutTest) GetTestCustomSelect13OptionId() int32`

GetTestCustomSelect13OptionId returns the TestCustomSelect13OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect13OptionIdOk

`func (o *TestsPutTest) GetTestCustomSelect13OptionIdOk() (*int32, bool)`

GetTestCustomSelect13OptionIdOk returns a tuple with the TestCustomSelect13OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect13OptionId

`func (o *TestsPutTest) SetTestCustomSelect13OptionId(v int32)`

SetTestCustomSelect13OptionId sets TestCustomSelect13OptionId field to given value.

### HasTestCustomSelect13OptionId

`func (o *TestsPutTest) HasTestCustomSelect13OptionId() bool`

HasTestCustomSelect13OptionId returns a boolean if a field has been set.

### GetTestCustomSelect14OptionId

`func (o *TestsPutTest) GetTestCustomSelect14OptionId() int32`

GetTestCustomSelect14OptionId returns the TestCustomSelect14OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect14OptionIdOk

`func (o *TestsPutTest) GetTestCustomSelect14OptionIdOk() (*int32, bool)`

GetTestCustomSelect14OptionIdOk returns a tuple with the TestCustomSelect14OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect14OptionId

`func (o *TestsPutTest) SetTestCustomSelect14OptionId(v int32)`

SetTestCustomSelect14OptionId sets TestCustomSelect14OptionId field to given value.

### HasTestCustomSelect14OptionId

`func (o *TestsPutTest) HasTestCustomSelect14OptionId() bool`

HasTestCustomSelect14OptionId returns a boolean if a field has been set.

### GetTestCustomSelect15OptionId

`func (o *TestsPutTest) GetTestCustomSelect15OptionId() int32`

GetTestCustomSelect15OptionId returns the TestCustomSelect15OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect15OptionIdOk

`func (o *TestsPutTest) GetTestCustomSelect15OptionIdOk() (*int32, bool)`

GetTestCustomSelect15OptionIdOk returns a tuple with the TestCustomSelect15OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect15OptionId

`func (o *TestsPutTest) SetTestCustomSelect15OptionId(v int32)`

SetTestCustomSelect15OptionId sets TestCustomSelect15OptionId field to given value.

### HasTestCustomSelect15OptionId

`func (o *TestsPutTest) HasTestCustomSelect15OptionId() bool`

HasTestCustomSelect15OptionId returns a boolean if a field has been set.

### GetTestCustomSelect16OptionId

`func (o *TestsPutTest) GetTestCustomSelect16OptionId() int32`

GetTestCustomSelect16OptionId returns the TestCustomSelect16OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect16OptionIdOk

`func (o *TestsPutTest) GetTestCustomSelect16OptionIdOk() (*int32, bool)`

GetTestCustomSelect16OptionIdOk returns a tuple with the TestCustomSelect16OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect16OptionId

`func (o *TestsPutTest) SetTestCustomSelect16OptionId(v int32)`

SetTestCustomSelect16OptionId sets TestCustomSelect16OptionId field to given value.

### HasTestCustomSelect16OptionId

`func (o *TestsPutTest) HasTestCustomSelect16OptionId() bool`

HasTestCustomSelect16OptionId returns a boolean if a field has been set.

### GetTestCustomSelect17OptionId

`func (o *TestsPutTest) GetTestCustomSelect17OptionId() int32`

GetTestCustomSelect17OptionId returns the TestCustomSelect17OptionId field if non-nil, zero value otherwise.

### GetTestCustomSelect17OptionIdOk

`func (o *TestsPutTest) GetTestCustomSelect17OptionIdOk() (*int32, bool)`

GetTestCustomSelect17OptionIdOk returns a tuple with the TestCustomSelect17OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCustomSelect17OptionId

`func (o *TestsPutTest) SetTestCustomSelect17OptionId(v int32)`

SetTestCustomSelect17OptionId sets TestCustomSelect17OptionId field to given value.

### HasTestCustomSelect17OptionId

`func (o *TestsPutTest) HasTestCustomSelect17OptionId() bool`

HasTestCustomSelect17OptionId returns a boolean if a field has been set.

### GetFieldData

`func (o *TestsPutTest) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *TestsPutTest) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *TestsPutTest) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *TestsPutTest) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *TestsPutTest) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *TestsPutTest) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


