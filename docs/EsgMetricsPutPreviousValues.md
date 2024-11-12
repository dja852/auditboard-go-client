# EsgMetricsPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **float32** |  | [optional] 
**RunningSum** | Pointer to **bool** |  | [optional] [default to false]
**ReviewerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**EsgMetricFrequencyOptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_metric_frequency_options.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_frequency_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Direction** | Pointer to **string** |  | [optional] [default to "NULL"]
**Health** | Pointer to **string** |  | [optional] [default to "NULL"]
**EsgMetricUnitOptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_metric_unit_options.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_unit_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IsKey** | Pointer to **bool** |  | [optional] [default to false]
**Source** | Pointer to **string** |  | [optional] 
**FieldData** | Pointer to **interface{}** |  | [optional] 
**LastReportedValue** | Pointer to **float32** |  | [optional] 
**LastReportedPeriod** | Pointer to **string** |  | [optional] 
**LastReportedNotes** | Pointer to **string** |  | [optional] 
**Completion** | Pointer to **float32** |  | [optional] 
**Instructions** | Pointer to **string** |  | [optional] 
**NotesRequired** | Pointer to **bool** |  | [optional] [default to false]
**FilesRequired** | Pointer to **bool** |  | [optional] [default to false]
**EsgMetricCategoryId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_metric_categories.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_categories&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CurrentYearToDateValue** | Pointer to **float32** |  | [optional] [default to 0]
**PreviousYearToDateValue** | Pointer to **float32** |  | [optional] [default to 0]
**QualitativeQuestion** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**AuditSurveyId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_surveys.id&#x60;.&lt;fk table&#x3D;&#39;audit_surveys&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ReportResponse** | Pointer to **string** |  | [optional] 
**TwoPriorYearToDateValue** | Pointer to **float32** |  | [optional] 
**AggregationMethod** | Pointer to **string** |  | [optional] [default to "NULL::character varying"]
**EsgMetricCustomSelect1OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_metric_custom_select1_options.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_custom_select1_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**EsgMetricCustomSelect2OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_metric_custom_select2_options.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_custom_select2_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**EsgMetricCustomSelect3OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_metric_custom_select3_options.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_custom_select3_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**EsgMetricCustomSelect4OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_metric_custom_select4_options.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_custom_select4_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**EsgMetricCustomSelect5OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_metric_custom_select5_options.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_custom_select5_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomDate1** | Pointer to **string** |  | [optional] 
**CustomDate2** | Pointer to **string** |  | [optional] 
**CustomDate3** | Pointer to **string** |  | [optional] 
**CustomText1** | Pointer to **string** |  | [optional] 
**CustomText2** | Pointer to **string** |  | [optional] 
**CustomText3** | Pointer to **string** |  | [optional] 
**CustomText4** | Pointer to **string** |  | [optional] 
**CustomText5** | Pointer to **string** |  | [optional] 
**CustomText6** | Pointer to **string** |  | [optional] 
**IsIntegrationMetric** | Pointer to **bool** |  | [optional] 
**IntegrationUid** | Pointer to **string** |  | [optional] 

## Methods

### NewEsgMetricsPutPreviousValues

`func NewEsgMetricsPutPreviousValues() *EsgMetricsPutPreviousValues`

NewEsgMetricsPutPreviousValues instantiates a new EsgMetricsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsgMetricsPutPreviousValuesWithDefaults

`func NewEsgMetricsPutPreviousValuesWithDefaults() *EsgMetricsPutPreviousValues`

NewEsgMetricsPutPreviousValuesWithDefaults instantiates a new EsgMetricsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EsgMetricsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsgMetricsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsgMetricsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EsgMetricsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EsgMetricsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EsgMetricsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EsgMetricsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EsgMetricsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EsgMetricsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EsgMetricsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EsgMetricsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EsgMetricsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EsgMetricsPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EsgMetricsPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EsgMetricsPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EsgMetricsPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *EsgMetricsPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EsgMetricsPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EsgMetricsPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EsgMetricsPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUid

`func (o *EsgMetricsPutPreviousValues) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *EsgMetricsPutPreviousValues) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *EsgMetricsPutPreviousValues) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *EsgMetricsPutPreviousValues) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetTarget

`func (o *EsgMetricsPutPreviousValues) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *EsgMetricsPutPreviousValues) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *EsgMetricsPutPreviousValues) SetTarget(v float32)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *EsgMetricsPutPreviousValues) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetRunningSum

`func (o *EsgMetricsPutPreviousValues) GetRunningSum() bool`

GetRunningSum returns the RunningSum field if non-nil, zero value otherwise.

### GetRunningSumOk

`func (o *EsgMetricsPutPreviousValues) GetRunningSumOk() (*bool, bool)`

GetRunningSumOk returns a tuple with the RunningSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningSum

`func (o *EsgMetricsPutPreviousValues) SetRunningSum(v bool)`

SetRunningSum sets RunningSum field to given value.

### HasRunningSum

`func (o *EsgMetricsPutPreviousValues) HasRunningSum() bool`

HasRunningSum returns a boolean if a field has been set.

### GetReviewerUserId

`func (o *EsgMetricsPutPreviousValues) GetReviewerUserId() int32`

GetReviewerUserId returns the ReviewerUserId field if non-nil, zero value otherwise.

### GetReviewerUserIdOk

`func (o *EsgMetricsPutPreviousValues) GetReviewerUserIdOk() (*int32, bool)`

GetReviewerUserIdOk returns a tuple with the ReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerUserId

`func (o *EsgMetricsPutPreviousValues) SetReviewerUserId(v int32)`

SetReviewerUserId sets ReviewerUserId field to given value.

### HasReviewerUserId

`func (o *EsgMetricsPutPreviousValues) HasReviewerUserId() bool`

HasReviewerUserId returns a boolean if a field has been set.

### GetEsgMetricFrequencyOptionId

`func (o *EsgMetricsPutPreviousValues) GetEsgMetricFrequencyOptionId() int32`

GetEsgMetricFrequencyOptionId returns the EsgMetricFrequencyOptionId field if non-nil, zero value otherwise.

### GetEsgMetricFrequencyOptionIdOk

`func (o *EsgMetricsPutPreviousValues) GetEsgMetricFrequencyOptionIdOk() (*int32, bool)`

GetEsgMetricFrequencyOptionIdOk returns a tuple with the EsgMetricFrequencyOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricFrequencyOptionId

`func (o *EsgMetricsPutPreviousValues) SetEsgMetricFrequencyOptionId(v int32)`

SetEsgMetricFrequencyOptionId sets EsgMetricFrequencyOptionId field to given value.

### HasEsgMetricFrequencyOptionId

`func (o *EsgMetricsPutPreviousValues) HasEsgMetricFrequencyOptionId() bool`

HasEsgMetricFrequencyOptionId returns a boolean if a field has been set.

### GetDirection

`func (o *EsgMetricsPutPreviousValues) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *EsgMetricsPutPreviousValues) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *EsgMetricsPutPreviousValues) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *EsgMetricsPutPreviousValues) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetHealth

`func (o *EsgMetricsPutPreviousValues) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *EsgMetricsPutPreviousValues) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *EsgMetricsPutPreviousValues) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *EsgMetricsPutPreviousValues) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetEsgMetricUnitOptionId

`func (o *EsgMetricsPutPreviousValues) GetEsgMetricUnitOptionId() int32`

GetEsgMetricUnitOptionId returns the EsgMetricUnitOptionId field if non-nil, zero value otherwise.

### GetEsgMetricUnitOptionIdOk

`func (o *EsgMetricsPutPreviousValues) GetEsgMetricUnitOptionIdOk() (*int32, bool)`

GetEsgMetricUnitOptionIdOk returns a tuple with the EsgMetricUnitOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricUnitOptionId

`func (o *EsgMetricsPutPreviousValues) SetEsgMetricUnitOptionId(v int32)`

SetEsgMetricUnitOptionId sets EsgMetricUnitOptionId field to given value.

### HasEsgMetricUnitOptionId

`func (o *EsgMetricsPutPreviousValues) HasEsgMetricUnitOptionId() bool`

HasEsgMetricUnitOptionId returns a boolean if a field has been set.

### GetIsKey

`func (o *EsgMetricsPutPreviousValues) GetIsKey() bool`

GetIsKey returns the IsKey field if non-nil, zero value otherwise.

### GetIsKeyOk

`func (o *EsgMetricsPutPreviousValues) GetIsKeyOk() (*bool, bool)`

GetIsKeyOk returns a tuple with the IsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKey

`func (o *EsgMetricsPutPreviousValues) SetIsKey(v bool)`

SetIsKey sets IsKey field to given value.

### HasIsKey

`func (o *EsgMetricsPutPreviousValues) HasIsKey() bool`

HasIsKey returns a boolean if a field has been set.

### GetSource

`func (o *EsgMetricsPutPreviousValues) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EsgMetricsPutPreviousValues) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EsgMetricsPutPreviousValues) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *EsgMetricsPutPreviousValues) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetFieldData

`func (o *EsgMetricsPutPreviousValues) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *EsgMetricsPutPreviousValues) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *EsgMetricsPutPreviousValues) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *EsgMetricsPutPreviousValues) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *EsgMetricsPutPreviousValues) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *EsgMetricsPutPreviousValues) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil
### GetLastReportedValue

`func (o *EsgMetricsPutPreviousValues) GetLastReportedValue() float32`

GetLastReportedValue returns the LastReportedValue field if non-nil, zero value otherwise.

### GetLastReportedValueOk

`func (o *EsgMetricsPutPreviousValues) GetLastReportedValueOk() (*float32, bool)`

GetLastReportedValueOk returns a tuple with the LastReportedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedValue

`func (o *EsgMetricsPutPreviousValues) SetLastReportedValue(v float32)`

SetLastReportedValue sets LastReportedValue field to given value.

### HasLastReportedValue

`func (o *EsgMetricsPutPreviousValues) HasLastReportedValue() bool`

HasLastReportedValue returns a boolean if a field has been set.

### GetLastReportedPeriod

`func (o *EsgMetricsPutPreviousValues) GetLastReportedPeriod() string`

GetLastReportedPeriod returns the LastReportedPeriod field if non-nil, zero value otherwise.

### GetLastReportedPeriodOk

`func (o *EsgMetricsPutPreviousValues) GetLastReportedPeriodOk() (*string, bool)`

GetLastReportedPeriodOk returns a tuple with the LastReportedPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedPeriod

`func (o *EsgMetricsPutPreviousValues) SetLastReportedPeriod(v string)`

SetLastReportedPeriod sets LastReportedPeriod field to given value.

### HasLastReportedPeriod

`func (o *EsgMetricsPutPreviousValues) HasLastReportedPeriod() bool`

HasLastReportedPeriod returns a boolean if a field has been set.

### GetLastReportedNotes

`func (o *EsgMetricsPutPreviousValues) GetLastReportedNotes() string`

GetLastReportedNotes returns the LastReportedNotes field if non-nil, zero value otherwise.

### GetLastReportedNotesOk

`func (o *EsgMetricsPutPreviousValues) GetLastReportedNotesOk() (*string, bool)`

GetLastReportedNotesOk returns a tuple with the LastReportedNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedNotes

`func (o *EsgMetricsPutPreviousValues) SetLastReportedNotes(v string)`

SetLastReportedNotes sets LastReportedNotes field to given value.

### HasLastReportedNotes

`func (o *EsgMetricsPutPreviousValues) HasLastReportedNotes() bool`

HasLastReportedNotes returns a boolean if a field has been set.

### GetCompletion

`func (o *EsgMetricsPutPreviousValues) GetCompletion() float32`

GetCompletion returns the Completion field if non-nil, zero value otherwise.

### GetCompletionOk

`func (o *EsgMetricsPutPreviousValues) GetCompletionOk() (*float32, bool)`

GetCompletionOk returns a tuple with the Completion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletion

`func (o *EsgMetricsPutPreviousValues) SetCompletion(v float32)`

SetCompletion sets Completion field to given value.

### HasCompletion

`func (o *EsgMetricsPutPreviousValues) HasCompletion() bool`

HasCompletion returns a boolean if a field has been set.

### GetInstructions

`func (o *EsgMetricsPutPreviousValues) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *EsgMetricsPutPreviousValues) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *EsgMetricsPutPreviousValues) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *EsgMetricsPutPreviousValues) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetNotesRequired

`func (o *EsgMetricsPutPreviousValues) GetNotesRequired() bool`

GetNotesRequired returns the NotesRequired field if non-nil, zero value otherwise.

### GetNotesRequiredOk

`func (o *EsgMetricsPutPreviousValues) GetNotesRequiredOk() (*bool, bool)`

GetNotesRequiredOk returns a tuple with the NotesRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesRequired

`func (o *EsgMetricsPutPreviousValues) SetNotesRequired(v bool)`

SetNotesRequired sets NotesRequired field to given value.

### HasNotesRequired

`func (o *EsgMetricsPutPreviousValues) HasNotesRequired() bool`

HasNotesRequired returns a boolean if a field has been set.

### GetFilesRequired

`func (o *EsgMetricsPutPreviousValues) GetFilesRequired() bool`

GetFilesRequired returns the FilesRequired field if non-nil, zero value otherwise.

### GetFilesRequiredOk

`func (o *EsgMetricsPutPreviousValues) GetFilesRequiredOk() (*bool, bool)`

GetFilesRequiredOk returns a tuple with the FilesRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesRequired

`func (o *EsgMetricsPutPreviousValues) SetFilesRequired(v bool)`

SetFilesRequired sets FilesRequired field to given value.

### HasFilesRequired

`func (o *EsgMetricsPutPreviousValues) HasFilesRequired() bool`

HasFilesRequired returns a boolean if a field has been set.

### GetEsgMetricCategoryId

`func (o *EsgMetricsPutPreviousValues) GetEsgMetricCategoryId() int32`

GetEsgMetricCategoryId returns the EsgMetricCategoryId field if non-nil, zero value otherwise.

### GetEsgMetricCategoryIdOk

`func (o *EsgMetricsPutPreviousValues) GetEsgMetricCategoryIdOk() (*int32, bool)`

GetEsgMetricCategoryIdOk returns a tuple with the EsgMetricCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricCategoryId

`func (o *EsgMetricsPutPreviousValues) SetEsgMetricCategoryId(v int32)`

SetEsgMetricCategoryId sets EsgMetricCategoryId field to given value.

### HasEsgMetricCategoryId

`func (o *EsgMetricsPutPreviousValues) HasEsgMetricCategoryId() bool`

HasEsgMetricCategoryId returns a boolean if a field has been set.

### GetDescription

`func (o *EsgMetricsPutPreviousValues) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EsgMetricsPutPreviousValues) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EsgMetricsPutPreviousValues) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EsgMetricsPutPreviousValues) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCurrentYearToDateValue

`func (o *EsgMetricsPutPreviousValues) GetCurrentYearToDateValue() float32`

GetCurrentYearToDateValue returns the CurrentYearToDateValue field if non-nil, zero value otherwise.

### GetCurrentYearToDateValueOk

`func (o *EsgMetricsPutPreviousValues) GetCurrentYearToDateValueOk() (*float32, bool)`

GetCurrentYearToDateValueOk returns a tuple with the CurrentYearToDateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentYearToDateValue

`func (o *EsgMetricsPutPreviousValues) SetCurrentYearToDateValue(v float32)`

SetCurrentYearToDateValue sets CurrentYearToDateValue field to given value.

### HasCurrentYearToDateValue

`func (o *EsgMetricsPutPreviousValues) HasCurrentYearToDateValue() bool`

HasCurrentYearToDateValue returns a boolean if a field has been set.

### GetPreviousYearToDateValue

`func (o *EsgMetricsPutPreviousValues) GetPreviousYearToDateValue() float32`

GetPreviousYearToDateValue returns the PreviousYearToDateValue field if non-nil, zero value otherwise.

### GetPreviousYearToDateValueOk

`func (o *EsgMetricsPutPreviousValues) GetPreviousYearToDateValueOk() (*float32, bool)`

GetPreviousYearToDateValueOk returns a tuple with the PreviousYearToDateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousYearToDateValue

`func (o *EsgMetricsPutPreviousValues) SetPreviousYearToDateValue(v float32)`

SetPreviousYearToDateValue sets PreviousYearToDateValue field to given value.

### HasPreviousYearToDateValue

`func (o *EsgMetricsPutPreviousValues) HasPreviousYearToDateValue() bool`

HasPreviousYearToDateValue returns a boolean if a field has been set.

### GetQualitativeQuestion

`func (o *EsgMetricsPutPreviousValues) GetQualitativeQuestion() string`

GetQualitativeQuestion returns the QualitativeQuestion field if non-nil, zero value otherwise.

### GetQualitativeQuestionOk

`func (o *EsgMetricsPutPreviousValues) GetQualitativeQuestionOk() (*string, bool)`

GetQualitativeQuestionOk returns a tuple with the QualitativeQuestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualitativeQuestion

`func (o *EsgMetricsPutPreviousValues) SetQualitativeQuestion(v string)`

SetQualitativeQuestion sets QualitativeQuestion field to given value.

### HasQualitativeQuestion

`func (o *EsgMetricsPutPreviousValues) HasQualitativeQuestion() bool`

HasQualitativeQuestion returns a boolean if a field has been set.

### GetType

`func (o *EsgMetricsPutPreviousValues) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EsgMetricsPutPreviousValues) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EsgMetricsPutPreviousValues) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EsgMetricsPutPreviousValues) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAuditSurveyId

`func (o *EsgMetricsPutPreviousValues) GetAuditSurveyId() int32`

GetAuditSurveyId returns the AuditSurveyId field if non-nil, zero value otherwise.

### GetAuditSurveyIdOk

`func (o *EsgMetricsPutPreviousValues) GetAuditSurveyIdOk() (*int32, bool)`

GetAuditSurveyIdOk returns a tuple with the AuditSurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSurveyId

`func (o *EsgMetricsPutPreviousValues) SetAuditSurveyId(v int32)`

SetAuditSurveyId sets AuditSurveyId field to given value.

### HasAuditSurveyId

`func (o *EsgMetricsPutPreviousValues) HasAuditSurveyId() bool`

HasAuditSurveyId returns a boolean if a field has been set.

### GetReportResponse

`func (o *EsgMetricsPutPreviousValues) GetReportResponse() string`

GetReportResponse returns the ReportResponse field if non-nil, zero value otherwise.

### GetReportResponseOk

`func (o *EsgMetricsPutPreviousValues) GetReportResponseOk() (*string, bool)`

GetReportResponseOk returns a tuple with the ReportResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportResponse

`func (o *EsgMetricsPutPreviousValues) SetReportResponse(v string)`

SetReportResponse sets ReportResponse field to given value.

### HasReportResponse

`func (o *EsgMetricsPutPreviousValues) HasReportResponse() bool`

HasReportResponse returns a boolean if a field has been set.

### GetTwoPriorYearToDateValue

`func (o *EsgMetricsPutPreviousValues) GetTwoPriorYearToDateValue() float32`

GetTwoPriorYearToDateValue returns the TwoPriorYearToDateValue field if non-nil, zero value otherwise.

### GetTwoPriorYearToDateValueOk

`func (o *EsgMetricsPutPreviousValues) GetTwoPriorYearToDateValueOk() (*float32, bool)`

GetTwoPriorYearToDateValueOk returns a tuple with the TwoPriorYearToDateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoPriorYearToDateValue

`func (o *EsgMetricsPutPreviousValues) SetTwoPriorYearToDateValue(v float32)`

SetTwoPriorYearToDateValue sets TwoPriorYearToDateValue field to given value.

### HasTwoPriorYearToDateValue

`func (o *EsgMetricsPutPreviousValues) HasTwoPriorYearToDateValue() bool`

HasTwoPriorYearToDateValue returns a boolean if a field has been set.

### GetAggregationMethod

`func (o *EsgMetricsPutPreviousValues) GetAggregationMethod() string`

GetAggregationMethod returns the AggregationMethod field if non-nil, zero value otherwise.

### GetAggregationMethodOk

`func (o *EsgMetricsPutPreviousValues) GetAggregationMethodOk() (*string, bool)`

GetAggregationMethodOk returns a tuple with the AggregationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationMethod

`func (o *EsgMetricsPutPreviousValues) SetAggregationMethod(v string)`

SetAggregationMethod sets AggregationMethod field to given value.

### HasAggregationMethod

`func (o *EsgMetricsPutPreviousValues) HasAggregationMethod() bool`

HasAggregationMethod returns a boolean if a field has been set.

### GetEsgMetricCustomSelect1OptionId

`func (o *EsgMetricsPutPreviousValues) GetEsgMetricCustomSelect1OptionId() int32`

GetEsgMetricCustomSelect1OptionId returns the EsgMetricCustomSelect1OptionId field if non-nil, zero value otherwise.

### GetEsgMetricCustomSelect1OptionIdOk

`func (o *EsgMetricsPutPreviousValues) GetEsgMetricCustomSelect1OptionIdOk() (*int32, bool)`

GetEsgMetricCustomSelect1OptionIdOk returns a tuple with the EsgMetricCustomSelect1OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricCustomSelect1OptionId

`func (o *EsgMetricsPutPreviousValues) SetEsgMetricCustomSelect1OptionId(v int32)`

SetEsgMetricCustomSelect1OptionId sets EsgMetricCustomSelect1OptionId field to given value.

### HasEsgMetricCustomSelect1OptionId

`func (o *EsgMetricsPutPreviousValues) HasEsgMetricCustomSelect1OptionId() bool`

HasEsgMetricCustomSelect1OptionId returns a boolean if a field has been set.

### GetEsgMetricCustomSelect2OptionId

`func (o *EsgMetricsPutPreviousValues) GetEsgMetricCustomSelect2OptionId() int32`

GetEsgMetricCustomSelect2OptionId returns the EsgMetricCustomSelect2OptionId field if non-nil, zero value otherwise.

### GetEsgMetricCustomSelect2OptionIdOk

`func (o *EsgMetricsPutPreviousValues) GetEsgMetricCustomSelect2OptionIdOk() (*int32, bool)`

GetEsgMetricCustomSelect2OptionIdOk returns a tuple with the EsgMetricCustomSelect2OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricCustomSelect2OptionId

`func (o *EsgMetricsPutPreviousValues) SetEsgMetricCustomSelect2OptionId(v int32)`

SetEsgMetricCustomSelect2OptionId sets EsgMetricCustomSelect2OptionId field to given value.

### HasEsgMetricCustomSelect2OptionId

`func (o *EsgMetricsPutPreviousValues) HasEsgMetricCustomSelect2OptionId() bool`

HasEsgMetricCustomSelect2OptionId returns a boolean if a field has been set.

### GetEsgMetricCustomSelect3OptionId

`func (o *EsgMetricsPutPreviousValues) GetEsgMetricCustomSelect3OptionId() int32`

GetEsgMetricCustomSelect3OptionId returns the EsgMetricCustomSelect3OptionId field if non-nil, zero value otherwise.

### GetEsgMetricCustomSelect3OptionIdOk

`func (o *EsgMetricsPutPreviousValues) GetEsgMetricCustomSelect3OptionIdOk() (*int32, bool)`

GetEsgMetricCustomSelect3OptionIdOk returns a tuple with the EsgMetricCustomSelect3OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricCustomSelect3OptionId

`func (o *EsgMetricsPutPreviousValues) SetEsgMetricCustomSelect3OptionId(v int32)`

SetEsgMetricCustomSelect3OptionId sets EsgMetricCustomSelect3OptionId field to given value.

### HasEsgMetricCustomSelect3OptionId

`func (o *EsgMetricsPutPreviousValues) HasEsgMetricCustomSelect3OptionId() bool`

HasEsgMetricCustomSelect3OptionId returns a boolean if a field has been set.

### GetEsgMetricCustomSelect4OptionId

`func (o *EsgMetricsPutPreviousValues) GetEsgMetricCustomSelect4OptionId() int32`

GetEsgMetricCustomSelect4OptionId returns the EsgMetricCustomSelect4OptionId field if non-nil, zero value otherwise.

### GetEsgMetricCustomSelect4OptionIdOk

`func (o *EsgMetricsPutPreviousValues) GetEsgMetricCustomSelect4OptionIdOk() (*int32, bool)`

GetEsgMetricCustomSelect4OptionIdOk returns a tuple with the EsgMetricCustomSelect4OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricCustomSelect4OptionId

`func (o *EsgMetricsPutPreviousValues) SetEsgMetricCustomSelect4OptionId(v int32)`

SetEsgMetricCustomSelect4OptionId sets EsgMetricCustomSelect4OptionId field to given value.

### HasEsgMetricCustomSelect4OptionId

`func (o *EsgMetricsPutPreviousValues) HasEsgMetricCustomSelect4OptionId() bool`

HasEsgMetricCustomSelect4OptionId returns a boolean if a field has been set.

### GetEsgMetricCustomSelect5OptionId

`func (o *EsgMetricsPutPreviousValues) GetEsgMetricCustomSelect5OptionId() int32`

GetEsgMetricCustomSelect5OptionId returns the EsgMetricCustomSelect5OptionId field if non-nil, zero value otherwise.

### GetEsgMetricCustomSelect5OptionIdOk

`func (o *EsgMetricsPutPreviousValues) GetEsgMetricCustomSelect5OptionIdOk() (*int32, bool)`

GetEsgMetricCustomSelect5OptionIdOk returns a tuple with the EsgMetricCustomSelect5OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricCustomSelect5OptionId

`func (o *EsgMetricsPutPreviousValues) SetEsgMetricCustomSelect5OptionId(v int32)`

SetEsgMetricCustomSelect5OptionId sets EsgMetricCustomSelect5OptionId field to given value.

### HasEsgMetricCustomSelect5OptionId

`func (o *EsgMetricsPutPreviousValues) HasEsgMetricCustomSelect5OptionId() bool`

HasEsgMetricCustomSelect5OptionId returns a boolean if a field has been set.

### GetCustomDate1

`func (o *EsgMetricsPutPreviousValues) GetCustomDate1() string`

GetCustomDate1 returns the CustomDate1 field if non-nil, zero value otherwise.

### GetCustomDate1Ok

`func (o *EsgMetricsPutPreviousValues) GetCustomDate1Ok() (*string, bool)`

GetCustomDate1Ok returns a tuple with the CustomDate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate1

`func (o *EsgMetricsPutPreviousValues) SetCustomDate1(v string)`

SetCustomDate1 sets CustomDate1 field to given value.

### HasCustomDate1

`func (o *EsgMetricsPutPreviousValues) HasCustomDate1() bool`

HasCustomDate1 returns a boolean if a field has been set.

### GetCustomDate2

`func (o *EsgMetricsPutPreviousValues) GetCustomDate2() string`

GetCustomDate2 returns the CustomDate2 field if non-nil, zero value otherwise.

### GetCustomDate2Ok

`func (o *EsgMetricsPutPreviousValues) GetCustomDate2Ok() (*string, bool)`

GetCustomDate2Ok returns a tuple with the CustomDate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate2

`func (o *EsgMetricsPutPreviousValues) SetCustomDate2(v string)`

SetCustomDate2 sets CustomDate2 field to given value.

### HasCustomDate2

`func (o *EsgMetricsPutPreviousValues) HasCustomDate2() bool`

HasCustomDate2 returns a boolean if a field has been set.

### GetCustomDate3

`func (o *EsgMetricsPutPreviousValues) GetCustomDate3() string`

GetCustomDate3 returns the CustomDate3 field if non-nil, zero value otherwise.

### GetCustomDate3Ok

`func (o *EsgMetricsPutPreviousValues) GetCustomDate3Ok() (*string, bool)`

GetCustomDate3Ok returns a tuple with the CustomDate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate3

`func (o *EsgMetricsPutPreviousValues) SetCustomDate3(v string)`

SetCustomDate3 sets CustomDate3 field to given value.

### HasCustomDate3

`func (o *EsgMetricsPutPreviousValues) HasCustomDate3() bool`

HasCustomDate3 returns a boolean if a field has been set.

### GetCustomText1

`func (o *EsgMetricsPutPreviousValues) GetCustomText1() string`

GetCustomText1 returns the CustomText1 field if non-nil, zero value otherwise.

### GetCustomText1Ok

`func (o *EsgMetricsPutPreviousValues) GetCustomText1Ok() (*string, bool)`

GetCustomText1Ok returns a tuple with the CustomText1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText1

`func (o *EsgMetricsPutPreviousValues) SetCustomText1(v string)`

SetCustomText1 sets CustomText1 field to given value.

### HasCustomText1

`func (o *EsgMetricsPutPreviousValues) HasCustomText1() bool`

HasCustomText1 returns a boolean if a field has been set.

### GetCustomText2

`func (o *EsgMetricsPutPreviousValues) GetCustomText2() string`

GetCustomText2 returns the CustomText2 field if non-nil, zero value otherwise.

### GetCustomText2Ok

`func (o *EsgMetricsPutPreviousValues) GetCustomText2Ok() (*string, bool)`

GetCustomText2Ok returns a tuple with the CustomText2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText2

`func (o *EsgMetricsPutPreviousValues) SetCustomText2(v string)`

SetCustomText2 sets CustomText2 field to given value.

### HasCustomText2

`func (o *EsgMetricsPutPreviousValues) HasCustomText2() bool`

HasCustomText2 returns a boolean if a field has been set.

### GetCustomText3

`func (o *EsgMetricsPutPreviousValues) GetCustomText3() string`

GetCustomText3 returns the CustomText3 field if non-nil, zero value otherwise.

### GetCustomText3Ok

`func (o *EsgMetricsPutPreviousValues) GetCustomText3Ok() (*string, bool)`

GetCustomText3Ok returns a tuple with the CustomText3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText3

`func (o *EsgMetricsPutPreviousValues) SetCustomText3(v string)`

SetCustomText3 sets CustomText3 field to given value.

### HasCustomText3

`func (o *EsgMetricsPutPreviousValues) HasCustomText3() bool`

HasCustomText3 returns a boolean if a field has been set.

### GetCustomText4

`func (o *EsgMetricsPutPreviousValues) GetCustomText4() string`

GetCustomText4 returns the CustomText4 field if non-nil, zero value otherwise.

### GetCustomText4Ok

`func (o *EsgMetricsPutPreviousValues) GetCustomText4Ok() (*string, bool)`

GetCustomText4Ok returns a tuple with the CustomText4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText4

`func (o *EsgMetricsPutPreviousValues) SetCustomText4(v string)`

SetCustomText4 sets CustomText4 field to given value.

### HasCustomText4

`func (o *EsgMetricsPutPreviousValues) HasCustomText4() bool`

HasCustomText4 returns a boolean if a field has been set.

### GetCustomText5

`func (o *EsgMetricsPutPreviousValues) GetCustomText5() string`

GetCustomText5 returns the CustomText5 field if non-nil, zero value otherwise.

### GetCustomText5Ok

`func (o *EsgMetricsPutPreviousValues) GetCustomText5Ok() (*string, bool)`

GetCustomText5Ok returns a tuple with the CustomText5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText5

`func (o *EsgMetricsPutPreviousValues) SetCustomText5(v string)`

SetCustomText5 sets CustomText5 field to given value.

### HasCustomText5

`func (o *EsgMetricsPutPreviousValues) HasCustomText5() bool`

HasCustomText5 returns a boolean if a field has been set.

### GetCustomText6

`func (o *EsgMetricsPutPreviousValues) GetCustomText6() string`

GetCustomText6 returns the CustomText6 field if non-nil, zero value otherwise.

### GetCustomText6Ok

`func (o *EsgMetricsPutPreviousValues) GetCustomText6Ok() (*string, bool)`

GetCustomText6Ok returns a tuple with the CustomText6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText6

`func (o *EsgMetricsPutPreviousValues) SetCustomText6(v string)`

SetCustomText6 sets CustomText6 field to given value.

### HasCustomText6

`func (o *EsgMetricsPutPreviousValues) HasCustomText6() bool`

HasCustomText6 returns a boolean if a field has been set.

### GetIsIntegrationMetric

`func (o *EsgMetricsPutPreviousValues) GetIsIntegrationMetric() bool`

GetIsIntegrationMetric returns the IsIntegrationMetric field if non-nil, zero value otherwise.

### GetIsIntegrationMetricOk

`func (o *EsgMetricsPutPreviousValues) GetIsIntegrationMetricOk() (*bool, bool)`

GetIsIntegrationMetricOk returns a tuple with the IsIntegrationMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIntegrationMetric

`func (o *EsgMetricsPutPreviousValues) SetIsIntegrationMetric(v bool)`

SetIsIntegrationMetric sets IsIntegrationMetric field to given value.

### HasIsIntegrationMetric

`func (o *EsgMetricsPutPreviousValues) HasIsIntegrationMetric() bool`

HasIsIntegrationMetric returns a boolean if a field has been set.

### GetIntegrationUid

`func (o *EsgMetricsPutPreviousValues) GetIntegrationUid() string`

GetIntegrationUid returns the IntegrationUid field if non-nil, zero value otherwise.

### GetIntegrationUidOk

`func (o *EsgMetricsPutPreviousValues) GetIntegrationUidOk() (*string, bool)`

GetIntegrationUidOk returns a tuple with the IntegrationUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationUid

`func (o *EsgMetricsPutPreviousValues) SetIntegrationUid(v string)`

SetIntegrationUid sets IntegrationUid field to given value.

### HasIntegrationUid

`func (o *EsgMetricsPutPreviousValues) HasIntegrationUid() bool`

HasIntegrationUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


