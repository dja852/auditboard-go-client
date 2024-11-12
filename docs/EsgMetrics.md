# EsgMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Uid** | **string** |  | 
**Target** | Pointer to **float32** |  | [optional] 
**RunningSum** | **bool** |  | [default to false]
**ReviewerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**EsgMetricFrequencyOptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_metric_frequency_options.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_frequency_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Direction** | Pointer to **string** |  | [optional] [default to "NULL"]
**Health** | Pointer to **string** |  | [optional] [default to "NULL"]
**EsgMetricUnitOptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_metric_unit_options.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_unit_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IsKey** | **bool** |  | [default to false]
**Source** | Pointer to **string** |  | [optional] 
**FieldData** | Pointer to **interface{}** |  | [optional] 
**LastReportedValue** | Pointer to **float32** |  | [optional] 
**LastReportedPeriod** | Pointer to **string** |  | [optional] 
**LastReportedNotes** | Pointer to **string** |  | [optional] 
**Completion** | Pointer to **float32** |  | [optional] 
**Instructions** | Pointer to **string** |  | [optional] 
**NotesRequired** | **bool** |  | [default to false]
**FilesRequired** | **bool** |  | [default to false]
**EsgMetricCategoryId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_metric_categories.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_categories&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CurrentYearToDateValue** | Pointer to **float32** |  | [optional] [default to 0]
**PreviousYearToDateValue** | Pointer to **float32** |  | [optional] [default to 0]
**QualitativeQuestion** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
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

### NewEsgMetrics

`func NewEsgMetrics(name string, uid string, runningSum bool, isKey bool, notesRequired bool, filesRequired bool, type_ string, ) *EsgMetrics`

NewEsgMetrics instantiates a new EsgMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsgMetricsWithDefaults

`func NewEsgMetricsWithDefaults() *EsgMetrics`

NewEsgMetricsWithDefaults instantiates a new EsgMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EsgMetrics) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsgMetrics) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsgMetrics) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EsgMetrics) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EsgMetrics) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EsgMetrics) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EsgMetrics) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EsgMetrics) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EsgMetrics) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EsgMetrics) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EsgMetrics) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EsgMetrics) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EsgMetrics) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EsgMetrics) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EsgMetrics) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EsgMetrics) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *EsgMetrics) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EsgMetrics) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EsgMetrics) SetName(v string)`

SetName sets Name field to given value.


### GetUid

`func (o *EsgMetrics) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *EsgMetrics) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *EsgMetrics) SetUid(v string)`

SetUid sets Uid field to given value.


### GetTarget

`func (o *EsgMetrics) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *EsgMetrics) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *EsgMetrics) SetTarget(v float32)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *EsgMetrics) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetRunningSum

`func (o *EsgMetrics) GetRunningSum() bool`

GetRunningSum returns the RunningSum field if non-nil, zero value otherwise.

### GetRunningSumOk

`func (o *EsgMetrics) GetRunningSumOk() (*bool, bool)`

GetRunningSumOk returns a tuple with the RunningSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningSum

`func (o *EsgMetrics) SetRunningSum(v bool)`

SetRunningSum sets RunningSum field to given value.


### GetReviewerUserId

`func (o *EsgMetrics) GetReviewerUserId() int32`

GetReviewerUserId returns the ReviewerUserId field if non-nil, zero value otherwise.

### GetReviewerUserIdOk

`func (o *EsgMetrics) GetReviewerUserIdOk() (*int32, bool)`

GetReviewerUserIdOk returns a tuple with the ReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerUserId

`func (o *EsgMetrics) SetReviewerUserId(v int32)`

SetReviewerUserId sets ReviewerUserId field to given value.

### HasReviewerUserId

`func (o *EsgMetrics) HasReviewerUserId() bool`

HasReviewerUserId returns a boolean if a field has been set.

### GetEsgMetricFrequencyOptionId

`func (o *EsgMetrics) GetEsgMetricFrequencyOptionId() int32`

GetEsgMetricFrequencyOptionId returns the EsgMetricFrequencyOptionId field if non-nil, zero value otherwise.

### GetEsgMetricFrequencyOptionIdOk

`func (o *EsgMetrics) GetEsgMetricFrequencyOptionIdOk() (*int32, bool)`

GetEsgMetricFrequencyOptionIdOk returns a tuple with the EsgMetricFrequencyOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricFrequencyOptionId

`func (o *EsgMetrics) SetEsgMetricFrequencyOptionId(v int32)`

SetEsgMetricFrequencyOptionId sets EsgMetricFrequencyOptionId field to given value.

### HasEsgMetricFrequencyOptionId

`func (o *EsgMetrics) HasEsgMetricFrequencyOptionId() bool`

HasEsgMetricFrequencyOptionId returns a boolean if a field has been set.

### GetDirection

`func (o *EsgMetrics) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *EsgMetrics) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *EsgMetrics) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *EsgMetrics) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetHealth

`func (o *EsgMetrics) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *EsgMetrics) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *EsgMetrics) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *EsgMetrics) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetEsgMetricUnitOptionId

`func (o *EsgMetrics) GetEsgMetricUnitOptionId() int32`

GetEsgMetricUnitOptionId returns the EsgMetricUnitOptionId field if non-nil, zero value otherwise.

### GetEsgMetricUnitOptionIdOk

`func (o *EsgMetrics) GetEsgMetricUnitOptionIdOk() (*int32, bool)`

GetEsgMetricUnitOptionIdOk returns a tuple with the EsgMetricUnitOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricUnitOptionId

`func (o *EsgMetrics) SetEsgMetricUnitOptionId(v int32)`

SetEsgMetricUnitOptionId sets EsgMetricUnitOptionId field to given value.

### HasEsgMetricUnitOptionId

`func (o *EsgMetrics) HasEsgMetricUnitOptionId() bool`

HasEsgMetricUnitOptionId returns a boolean if a field has been set.

### GetIsKey

`func (o *EsgMetrics) GetIsKey() bool`

GetIsKey returns the IsKey field if non-nil, zero value otherwise.

### GetIsKeyOk

`func (o *EsgMetrics) GetIsKeyOk() (*bool, bool)`

GetIsKeyOk returns a tuple with the IsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKey

`func (o *EsgMetrics) SetIsKey(v bool)`

SetIsKey sets IsKey field to given value.


### GetSource

`func (o *EsgMetrics) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EsgMetrics) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EsgMetrics) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *EsgMetrics) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetFieldData

`func (o *EsgMetrics) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *EsgMetrics) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *EsgMetrics) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *EsgMetrics) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *EsgMetrics) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *EsgMetrics) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil
### GetLastReportedValue

`func (o *EsgMetrics) GetLastReportedValue() float32`

GetLastReportedValue returns the LastReportedValue field if non-nil, zero value otherwise.

### GetLastReportedValueOk

`func (o *EsgMetrics) GetLastReportedValueOk() (*float32, bool)`

GetLastReportedValueOk returns a tuple with the LastReportedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedValue

`func (o *EsgMetrics) SetLastReportedValue(v float32)`

SetLastReportedValue sets LastReportedValue field to given value.

### HasLastReportedValue

`func (o *EsgMetrics) HasLastReportedValue() bool`

HasLastReportedValue returns a boolean if a field has been set.

### GetLastReportedPeriod

`func (o *EsgMetrics) GetLastReportedPeriod() string`

GetLastReportedPeriod returns the LastReportedPeriod field if non-nil, zero value otherwise.

### GetLastReportedPeriodOk

`func (o *EsgMetrics) GetLastReportedPeriodOk() (*string, bool)`

GetLastReportedPeriodOk returns a tuple with the LastReportedPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedPeriod

`func (o *EsgMetrics) SetLastReportedPeriod(v string)`

SetLastReportedPeriod sets LastReportedPeriod field to given value.

### HasLastReportedPeriod

`func (o *EsgMetrics) HasLastReportedPeriod() bool`

HasLastReportedPeriod returns a boolean if a field has been set.

### GetLastReportedNotes

`func (o *EsgMetrics) GetLastReportedNotes() string`

GetLastReportedNotes returns the LastReportedNotes field if non-nil, zero value otherwise.

### GetLastReportedNotesOk

`func (o *EsgMetrics) GetLastReportedNotesOk() (*string, bool)`

GetLastReportedNotesOk returns a tuple with the LastReportedNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedNotes

`func (o *EsgMetrics) SetLastReportedNotes(v string)`

SetLastReportedNotes sets LastReportedNotes field to given value.

### HasLastReportedNotes

`func (o *EsgMetrics) HasLastReportedNotes() bool`

HasLastReportedNotes returns a boolean if a field has been set.

### GetCompletion

`func (o *EsgMetrics) GetCompletion() float32`

GetCompletion returns the Completion field if non-nil, zero value otherwise.

### GetCompletionOk

`func (o *EsgMetrics) GetCompletionOk() (*float32, bool)`

GetCompletionOk returns a tuple with the Completion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletion

`func (o *EsgMetrics) SetCompletion(v float32)`

SetCompletion sets Completion field to given value.

### HasCompletion

`func (o *EsgMetrics) HasCompletion() bool`

HasCompletion returns a boolean if a field has been set.

### GetInstructions

`func (o *EsgMetrics) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *EsgMetrics) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *EsgMetrics) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *EsgMetrics) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetNotesRequired

`func (o *EsgMetrics) GetNotesRequired() bool`

GetNotesRequired returns the NotesRequired field if non-nil, zero value otherwise.

### GetNotesRequiredOk

`func (o *EsgMetrics) GetNotesRequiredOk() (*bool, bool)`

GetNotesRequiredOk returns a tuple with the NotesRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesRequired

`func (o *EsgMetrics) SetNotesRequired(v bool)`

SetNotesRequired sets NotesRequired field to given value.


### GetFilesRequired

`func (o *EsgMetrics) GetFilesRequired() bool`

GetFilesRequired returns the FilesRequired field if non-nil, zero value otherwise.

### GetFilesRequiredOk

`func (o *EsgMetrics) GetFilesRequiredOk() (*bool, bool)`

GetFilesRequiredOk returns a tuple with the FilesRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesRequired

`func (o *EsgMetrics) SetFilesRequired(v bool)`

SetFilesRequired sets FilesRequired field to given value.


### GetEsgMetricCategoryId

`func (o *EsgMetrics) GetEsgMetricCategoryId() int32`

GetEsgMetricCategoryId returns the EsgMetricCategoryId field if non-nil, zero value otherwise.

### GetEsgMetricCategoryIdOk

`func (o *EsgMetrics) GetEsgMetricCategoryIdOk() (*int32, bool)`

GetEsgMetricCategoryIdOk returns a tuple with the EsgMetricCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricCategoryId

`func (o *EsgMetrics) SetEsgMetricCategoryId(v int32)`

SetEsgMetricCategoryId sets EsgMetricCategoryId field to given value.

### HasEsgMetricCategoryId

`func (o *EsgMetrics) HasEsgMetricCategoryId() bool`

HasEsgMetricCategoryId returns a boolean if a field has been set.

### GetDescription

`func (o *EsgMetrics) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EsgMetrics) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EsgMetrics) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EsgMetrics) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCurrentYearToDateValue

`func (o *EsgMetrics) GetCurrentYearToDateValue() float32`

GetCurrentYearToDateValue returns the CurrentYearToDateValue field if non-nil, zero value otherwise.

### GetCurrentYearToDateValueOk

`func (o *EsgMetrics) GetCurrentYearToDateValueOk() (*float32, bool)`

GetCurrentYearToDateValueOk returns a tuple with the CurrentYearToDateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentYearToDateValue

`func (o *EsgMetrics) SetCurrentYearToDateValue(v float32)`

SetCurrentYearToDateValue sets CurrentYearToDateValue field to given value.

### HasCurrentYearToDateValue

`func (o *EsgMetrics) HasCurrentYearToDateValue() bool`

HasCurrentYearToDateValue returns a boolean if a field has been set.

### GetPreviousYearToDateValue

`func (o *EsgMetrics) GetPreviousYearToDateValue() float32`

GetPreviousYearToDateValue returns the PreviousYearToDateValue field if non-nil, zero value otherwise.

### GetPreviousYearToDateValueOk

`func (o *EsgMetrics) GetPreviousYearToDateValueOk() (*float32, bool)`

GetPreviousYearToDateValueOk returns a tuple with the PreviousYearToDateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousYearToDateValue

`func (o *EsgMetrics) SetPreviousYearToDateValue(v float32)`

SetPreviousYearToDateValue sets PreviousYearToDateValue field to given value.

### HasPreviousYearToDateValue

`func (o *EsgMetrics) HasPreviousYearToDateValue() bool`

HasPreviousYearToDateValue returns a boolean if a field has been set.

### GetQualitativeQuestion

`func (o *EsgMetrics) GetQualitativeQuestion() string`

GetQualitativeQuestion returns the QualitativeQuestion field if non-nil, zero value otherwise.

### GetQualitativeQuestionOk

`func (o *EsgMetrics) GetQualitativeQuestionOk() (*string, bool)`

GetQualitativeQuestionOk returns a tuple with the QualitativeQuestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualitativeQuestion

`func (o *EsgMetrics) SetQualitativeQuestion(v string)`

SetQualitativeQuestion sets QualitativeQuestion field to given value.

### HasQualitativeQuestion

`func (o *EsgMetrics) HasQualitativeQuestion() bool`

HasQualitativeQuestion returns a boolean if a field has been set.

### GetType

`func (o *EsgMetrics) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EsgMetrics) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EsgMetrics) SetType(v string)`

SetType sets Type field to given value.


### GetAuditSurveyId

`func (o *EsgMetrics) GetAuditSurveyId() int32`

GetAuditSurveyId returns the AuditSurveyId field if non-nil, zero value otherwise.

### GetAuditSurveyIdOk

`func (o *EsgMetrics) GetAuditSurveyIdOk() (*int32, bool)`

GetAuditSurveyIdOk returns a tuple with the AuditSurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSurveyId

`func (o *EsgMetrics) SetAuditSurveyId(v int32)`

SetAuditSurveyId sets AuditSurveyId field to given value.

### HasAuditSurveyId

`func (o *EsgMetrics) HasAuditSurveyId() bool`

HasAuditSurveyId returns a boolean if a field has been set.

### GetReportResponse

`func (o *EsgMetrics) GetReportResponse() string`

GetReportResponse returns the ReportResponse field if non-nil, zero value otherwise.

### GetReportResponseOk

`func (o *EsgMetrics) GetReportResponseOk() (*string, bool)`

GetReportResponseOk returns a tuple with the ReportResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportResponse

`func (o *EsgMetrics) SetReportResponse(v string)`

SetReportResponse sets ReportResponse field to given value.

### HasReportResponse

`func (o *EsgMetrics) HasReportResponse() bool`

HasReportResponse returns a boolean if a field has been set.

### GetTwoPriorYearToDateValue

`func (o *EsgMetrics) GetTwoPriorYearToDateValue() float32`

GetTwoPriorYearToDateValue returns the TwoPriorYearToDateValue field if non-nil, zero value otherwise.

### GetTwoPriorYearToDateValueOk

`func (o *EsgMetrics) GetTwoPriorYearToDateValueOk() (*float32, bool)`

GetTwoPriorYearToDateValueOk returns a tuple with the TwoPriorYearToDateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoPriorYearToDateValue

`func (o *EsgMetrics) SetTwoPriorYearToDateValue(v float32)`

SetTwoPriorYearToDateValue sets TwoPriorYearToDateValue field to given value.

### HasTwoPriorYearToDateValue

`func (o *EsgMetrics) HasTwoPriorYearToDateValue() bool`

HasTwoPriorYearToDateValue returns a boolean if a field has been set.

### GetAggregationMethod

`func (o *EsgMetrics) GetAggregationMethod() string`

GetAggregationMethod returns the AggregationMethod field if non-nil, zero value otherwise.

### GetAggregationMethodOk

`func (o *EsgMetrics) GetAggregationMethodOk() (*string, bool)`

GetAggregationMethodOk returns a tuple with the AggregationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationMethod

`func (o *EsgMetrics) SetAggregationMethod(v string)`

SetAggregationMethod sets AggregationMethod field to given value.

### HasAggregationMethod

`func (o *EsgMetrics) HasAggregationMethod() bool`

HasAggregationMethod returns a boolean if a field has been set.

### GetEsgMetricCustomSelect1OptionId

`func (o *EsgMetrics) GetEsgMetricCustomSelect1OptionId() int32`

GetEsgMetricCustomSelect1OptionId returns the EsgMetricCustomSelect1OptionId field if non-nil, zero value otherwise.

### GetEsgMetricCustomSelect1OptionIdOk

`func (o *EsgMetrics) GetEsgMetricCustomSelect1OptionIdOk() (*int32, bool)`

GetEsgMetricCustomSelect1OptionIdOk returns a tuple with the EsgMetricCustomSelect1OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricCustomSelect1OptionId

`func (o *EsgMetrics) SetEsgMetricCustomSelect1OptionId(v int32)`

SetEsgMetricCustomSelect1OptionId sets EsgMetricCustomSelect1OptionId field to given value.

### HasEsgMetricCustomSelect1OptionId

`func (o *EsgMetrics) HasEsgMetricCustomSelect1OptionId() bool`

HasEsgMetricCustomSelect1OptionId returns a boolean if a field has been set.

### GetEsgMetricCustomSelect2OptionId

`func (o *EsgMetrics) GetEsgMetricCustomSelect2OptionId() int32`

GetEsgMetricCustomSelect2OptionId returns the EsgMetricCustomSelect2OptionId field if non-nil, zero value otherwise.

### GetEsgMetricCustomSelect2OptionIdOk

`func (o *EsgMetrics) GetEsgMetricCustomSelect2OptionIdOk() (*int32, bool)`

GetEsgMetricCustomSelect2OptionIdOk returns a tuple with the EsgMetricCustomSelect2OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricCustomSelect2OptionId

`func (o *EsgMetrics) SetEsgMetricCustomSelect2OptionId(v int32)`

SetEsgMetricCustomSelect2OptionId sets EsgMetricCustomSelect2OptionId field to given value.

### HasEsgMetricCustomSelect2OptionId

`func (o *EsgMetrics) HasEsgMetricCustomSelect2OptionId() bool`

HasEsgMetricCustomSelect2OptionId returns a boolean if a field has been set.

### GetEsgMetricCustomSelect3OptionId

`func (o *EsgMetrics) GetEsgMetricCustomSelect3OptionId() int32`

GetEsgMetricCustomSelect3OptionId returns the EsgMetricCustomSelect3OptionId field if non-nil, zero value otherwise.

### GetEsgMetricCustomSelect3OptionIdOk

`func (o *EsgMetrics) GetEsgMetricCustomSelect3OptionIdOk() (*int32, bool)`

GetEsgMetricCustomSelect3OptionIdOk returns a tuple with the EsgMetricCustomSelect3OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricCustomSelect3OptionId

`func (o *EsgMetrics) SetEsgMetricCustomSelect3OptionId(v int32)`

SetEsgMetricCustomSelect3OptionId sets EsgMetricCustomSelect3OptionId field to given value.

### HasEsgMetricCustomSelect3OptionId

`func (o *EsgMetrics) HasEsgMetricCustomSelect3OptionId() bool`

HasEsgMetricCustomSelect3OptionId returns a boolean if a field has been set.

### GetEsgMetricCustomSelect4OptionId

`func (o *EsgMetrics) GetEsgMetricCustomSelect4OptionId() int32`

GetEsgMetricCustomSelect4OptionId returns the EsgMetricCustomSelect4OptionId field if non-nil, zero value otherwise.

### GetEsgMetricCustomSelect4OptionIdOk

`func (o *EsgMetrics) GetEsgMetricCustomSelect4OptionIdOk() (*int32, bool)`

GetEsgMetricCustomSelect4OptionIdOk returns a tuple with the EsgMetricCustomSelect4OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricCustomSelect4OptionId

`func (o *EsgMetrics) SetEsgMetricCustomSelect4OptionId(v int32)`

SetEsgMetricCustomSelect4OptionId sets EsgMetricCustomSelect4OptionId field to given value.

### HasEsgMetricCustomSelect4OptionId

`func (o *EsgMetrics) HasEsgMetricCustomSelect4OptionId() bool`

HasEsgMetricCustomSelect4OptionId returns a boolean if a field has been set.

### GetEsgMetricCustomSelect5OptionId

`func (o *EsgMetrics) GetEsgMetricCustomSelect5OptionId() int32`

GetEsgMetricCustomSelect5OptionId returns the EsgMetricCustomSelect5OptionId field if non-nil, zero value otherwise.

### GetEsgMetricCustomSelect5OptionIdOk

`func (o *EsgMetrics) GetEsgMetricCustomSelect5OptionIdOk() (*int32, bool)`

GetEsgMetricCustomSelect5OptionIdOk returns a tuple with the EsgMetricCustomSelect5OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricCustomSelect5OptionId

`func (o *EsgMetrics) SetEsgMetricCustomSelect5OptionId(v int32)`

SetEsgMetricCustomSelect5OptionId sets EsgMetricCustomSelect5OptionId field to given value.

### HasEsgMetricCustomSelect5OptionId

`func (o *EsgMetrics) HasEsgMetricCustomSelect5OptionId() bool`

HasEsgMetricCustomSelect5OptionId returns a boolean if a field has been set.

### GetCustomDate1

`func (o *EsgMetrics) GetCustomDate1() string`

GetCustomDate1 returns the CustomDate1 field if non-nil, zero value otherwise.

### GetCustomDate1Ok

`func (o *EsgMetrics) GetCustomDate1Ok() (*string, bool)`

GetCustomDate1Ok returns a tuple with the CustomDate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate1

`func (o *EsgMetrics) SetCustomDate1(v string)`

SetCustomDate1 sets CustomDate1 field to given value.

### HasCustomDate1

`func (o *EsgMetrics) HasCustomDate1() bool`

HasCustomDate1 returns a boolean if a field has been set.

### GetCustomDate2

`func (o *EsgMetrics) GetCustomDate2() string`

GetCustomDate2 returns the CustomDate2 field if non-nil, zero value otherwise.

### GetCustomDate2Ok

`func (o *EsgMetrics) GetCustomDate2Ok() (*string, bool)`

GetCustomDate2Ok returns a tuple with the CustomDate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate2

`func (o *EsgMetrics) SetCustomDate2(v string)`

SetCustomDate2 sets CustomDate2 field to given value.

### HasCustomDate2

`func (o *EsgMetrics) HasCustomDate2() bool`

HasCustomDate2 returns a boolean if a field has been set.

### GetCustomDate3

`func (o *EsgMetrics) GetCustomDate3() string`

GetCustomDate3 returns the CustomDate3 field if non-nil, zero value otherwise.

### GetCustomDate3Ok

`func (o *EsgMetrics) GetCustomDate3Ok() (*string, bool)`

GetCustomDate3Ok returns a tuple with the CustomDate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDate3

`func (o *EsgMetrics) SetCustomDate3(v string)`

SetCustomDate3 sets CustomDate3 field to given value.

### HasCustomDate3

`func (o *EsgMetrics) HasCustomDate3() bool`

HasCustomDate3 returns a boolean if a field has been set.

### GetCustomText1

`func (o *EsgMetrics) GetCustomText1() string`

GetCustomText1 returns the CustomText1 field if non-nil, zero value otherwise.

### GetCustomText1Ok

`func (o *EsgMetrics) GetCustomText1Ok() (*string, bool)`

GetCustomText1Ok returns a tuple with the CustomText1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText1

`func (o *EsgMetrics) SetCustomText1(v string)`

SetCustomText1 sets CustomText1 field to given value.

### HasCustomText1

`func (o *EsgMetrics) HasCustomText1() bool`

HasCustomText1 returns a boolean if a field has been set.

### GetCustomText2

`func (o *EsgMetrics) GetCustomText2() string`

GetCustomText2 returns the CustomText2 field if non-nil, zero value otherwise.

### GetCustomText2Ok

`func (o *EsgMetrics) GetCustomText2Ok() (*string, bool)`

GetCustomText2Ok returns a tuple with the CustomText2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText2

`func (o *EsgMetrics) SetCustomText2(v string)`

SetCustomText2 sets CustomText2 field to given value.

### HasCustomText2

`func (o *EsgMetrics) HasCustomText2() bool`

HasCustomText2 returns a boolean if a field has been set.

### GetCustomText3

`func (o *EsgMetrics) GetCustomText3() string`

GetCustomText3 returns the CustomText3 field if non-nil, zero value otherwise.

### GetCustomText3Ok

`func (o *EsgMetrics) GetCustomText3Ok() (*string, bool)`

GetCustomText3Ok returns a tuple with the CustomText3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText3

`func (o *EsgMetrics) SetCustomText3(v string)`

SetCustomText3 sets CustomText3 field to given value.

### HasCustomText3

`func (o *EsgMetrics) HasCustomText3() bool`

HasCustomText3 returns a boolean if a field has been set.

### GetCustomText4

`func (o *EsgMetrics) GetCustomText4() string`

GetCustomText4 returns the CustomText4 field if non-nil, zero value otherwise.

### GetCustomText4Ok

`func (o *EsgMetrics) GetCustomText4Ok() (*string, bool)`

GetCustomText4Ok returns a tuple with the CustomText4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText4

`func (o *EsgMetrics) SetCustomText4(v string)`

SetCustomText4 sets CustomText4 field to given value.

### HasCustomText4

`func (o *EsgMetrics) HasCustomText4() bool`

HasCustomText4 returns a boolean if a field has been set.

### GetCustomText5

`func (o *EsgMetrics) GetCustomText5() string`

GetCustomText5 returns the CustomText5 field if non-nil, zero value otherwise.

### GetCustomText5Ok

`func (o *EsgMetrics) GetCustomText5Ok() (*string, bool)`

GetCustomText5Ok returns a tuple with the CustomText5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText5

`func (o *EsgMetrics) SetCustomText5(v string)`

SetCustomText5 sets CustomText5 field to given value.

### HasCustomText5

`func (o *EsgMetrics) HasCustomText5() bool`

HasCustomText5 returns a boolean if a field has been set.

### GetCustomText6

`func (o *EsgMetrics) GetCustomText6() string`

GetCustomText6 returns the CustomText6 field if non-nil, zero value otherwise.

### GetCustomText6Ok

`func (o *EsgMetrics) GetCustomText6Ok() (*string, bool)`

GetCustomText6Ok returns a tuple with the CustomText6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText6

`func (o *EsgMetrics) SetCustomText6(v string)`

SetCustomText6 sets CustomText6 field to given value.

### HasCustomText6

`func (o *EsgMetrics) HasCustomText6() bool`

HasCustomText6 returns a boolean if a field has been set.

### GetIsIntegrationMetric

`func (o *EsgMetrics) GetIsIntegrationMetric() bool`

GetIsIntegrationMetric returns the IsIntegrationMetric field if non-nil, zero value otherwise.

### GetIsIntegrationMetricOk

`func (o *EsgMetrics) GetIsIntegrationMetricOk() (*bool, bool)`

GetIsIntegrationMetricOk returns a tuple with the IsIntegrationMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIntegrationMetric

`func (o *EsgMetrics) SetIsIntegrationMetric(v bool)`

SetIsIntegrationMetric sets IsIntegrationMetric field to given value.

### HasIsIntegrationMetric

`func (o *EsgMetrics) HasIsIntegrationMetric() bool`

HasIsIntegrationMetric returns a boolean if a field has been set.

### GetIntegrationUid

`func (o *EsgMetrics) GetIntegrationUid() string`

GetIntegrationUid returns the IntegrationUid field if non-nil, zero value otherwise.

### GetIntegrationUidOk

`func (o *EsgMetrics) GetIntegrationUidOk() (*string, bool)`

GetIntegrationUidOk returns a tuple with the IntegrationUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationUid

`func (o *EsgMetrics) SetIntegrationUid(v string)`

SetIntegrationUid sets IntegrationUid field to given value.

### HasIntegrationUid

`func (o *EsgMetrics) HasIntegrationUid() bool`

HasIntegrationUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


