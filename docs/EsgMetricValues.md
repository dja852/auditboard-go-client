# EsgMetricValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**EsgMetricId** | **int32** | Note: This is a Foreign Key to &#x60;esg_metrics.id&#x60;.&lt;fk table&#x3D;&#39;esg_metrics&#39; column&#x3D;&#39;id&#39;/&gt; | 
**Period** | **string** |  | 
**Value** | **float32** |  | 
**Target** | Pointer to **float32** |  | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**Notes** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**MissingFilesReason** | Pointer to **string** |  | [optional] 
**SubmittedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**EsgMetricAuditableEntityId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_metric_auditable_entities.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_auditable_entities&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**PeriodStart** | **string** |  | 
**PeriodEnd** | **string** |  | 
**CertifiedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FromUnitId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_metric_unit_options.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_unit_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ConvertedValue** | Pointer to **float32** |  | [optional] 
**QualitativeQuestion** | Pointer to **string** |  | [optional] 
**QualitativeQuestionResponse** | Pointer to **string** |  | [optional] 
**NextReviewerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AuditSurveyId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_surveys.id&#x60;.&lt;fk table&#x3D;&#39;audit_surveys&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 

## Methods

### NewEsgMetricValues

`func NewEsgMetricValues(esgMetricId int32, period string, value float32, status string, periodStart string, periodEnd string, ) *EsgMetricValues`

NewEsgMetricValues instantiates a new EsgMetricValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsgMetricValuesWithDefaults

`func NewEsgMetricValuesWithDefaults() *EsgMetricValues`

NewEsgMetricValuesWithDefaults instantiates a new EsgMetricValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EsgMetricValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsgMetricValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsgMetricValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EsgMetricValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EsgMetricValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EsgMetricValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EsgMetricValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EsgMetricValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EsgMetricValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EsgMetricValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EsgMetricValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EsgMetricValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EsgMetricValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EsgMetricValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EsgMetricValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EsgMetricValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetEsgMetricId

`func (o *EsgMetricValues) GetEsgMetricId() int32`

GetEsgMetricId returns the EsgMetricId field if non-nil, zero value otherwise.

### GetEsgMetricIdOk

`func (o *EsgMetricValues) GetEsgMetricIdOk() (*int32, bool)`

GetEsgMetricIdOk returns a tuple with the EsgMetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricId

`func (o *EsgMetricValues) SetEsgMetricId(v int32)`

SetEsgMetricId sets EsgMetricId field to given value.


### GetPeriod

`func (o *EsgMetricValues) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *EsgMetricValues) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *EsgMetricValues) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### GetValue

`func (o *EsgMetricValues) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EsgMetricValues) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EsgMetricValues) SetValue(v float32)`

SetValue sets Value field to given value.


### GetTarget

`func (o *EsgMetricValues) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *EsgMetricValues) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *EsgMetricValues) SetTarget(v float32)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *EsgMetricValues) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetDueDate

`func (o *EsgMetricValues) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *EsgMetricValues) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *EsgMetricValues) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *EsgMetricValues) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetStatus

`func (o *EsgMetricValues) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EsgMetricValues) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EsgMetricValues) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetNotes

`func (o *EsgMetricValues) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *EsgMetricValues) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *EsgMetricValues) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *EsgMetricValues) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetSource

`func (o *EsgMetricValues) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EsgMetricValues) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EsgMetricValues) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *EsgMetricValues) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetMissingFilesReason

`func (o *EsgMetricValues) GetMissingFilesReason() string`

GetMissingFilesReason returns the MissingFilesReason field if non-nil, zero value otherwise.

### GetMissingFilesReasonOk

`func (o *EsgMetricValues) GetMissingFilesReasonOk() (*string, bool)`

GetMissingFilesReasonOk returns a tuple with the MissingFilesReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingFilesReason

`func (o *EsgMetricValues) SetMissingFilesReason(v string)`

SetMissingFilesReason sets MissingFilesReason field to given value.

### HasMissingFilesReason

`func (o *EsgMetricValues) HasMissingFilesReason() bool`

HasMissingFilesReason returns a boolean if a field has been set.

### GetSubmittedByUserId

`func (o *EsgMetricValues) GetSubmittedByUserId() int32`

GetSubmittedByUserId returns the SubmittedByUserId field if non-nil, zero value otherwise.

### GetSubmittedByUserIdOk

`func (o *EsgMetricValues) GetSubmittedByUserIdOk() (*int32, bool)`

GetSubmittedByUserIdOk returns a tuple with the SubmittedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedByUserId

`func (o *EsgMetricValues) SetSubmittedByUserId(v int32)`

SetSubmittedByUserId sets SubmittedByUserId field to given value.

### HasSubmittedByUserId

`func (o *EsgMetricValues) HasSubmittedByUserId() bool`

HasSubmittedByUserId returns a boolean if a field has been set.

### GetEsgMetricAuditableEntityId

`func (o *EsgMetricValues) GetEsgMetricAuditableEntityId() int32`

GetEsgMetricAuditableEntityId returns the EsgMetricAuditableEntityId field if non-nil, zero value otherwise.

### GetEsgMetricAuditableEntityIdOk

`func (o *EsgMetricValues) GetEsgMetricAuditableEntityIdOk() (*int32, bool)`

GetEsgMetricAuditableEntityIdOk returns a tuple with the EsgMetricAuditableEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricAuditableEntityId

`func (o *EsgMetricValues) SetEsgMetricAuditableEntityId(v int32)`

SetEsgMetricAuditableEntityId sets EsgMetricAuditableEntityId field to given value.

### HasEsgMetricAuditableEntityId

`func (o *EsgMetricValues) HasEsgMetricAuditableEntityId() bool`

HasEsgMetricAuditableEntityId returns a boolean if a field has been set.

### GetPeriodStart

`func (o *EsgMetricValues) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *EsgMetricValues) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *EsgMetricValues) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *EsgMetricValues) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *EsgMetricValues) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *EsgMetricValues) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetCertifiedByUserId

`func (o *EsgMetricValues) GetCertifiedByUserId() int32`

GetCertifiedByUserId returns the CertifiedByUserId field if non-nil, zero value otherwise.

### GetCertifiedByUserIdOk

`func (o *EsgMetricValues) GetCertifiedByUserIdOk() (*int32, bool)`

GetCertifiedByUserIdOk returns a tuple with the CertifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifiedByUserId

`func (o *EsgMetricValues) SetCertifiedByUserId(v int32)`

SetCertifiedByUserId sets CertifiedByUserId field to given value.

### HasCertifiedByUserId

`func (o *EsgMetricValues) HasCertifiedByUserId() bool`

HasCertifiedByUserId returns a boolean if a field has been set.

### GetFromUnitId

`func (o *EsgMetricValues) GetFromUnitId() int32`

GetFromUnitId returns the FromUnitId field if non-nil, zero value otherwise.

### GetFromUnitIdOk

`func (o *EsgMetricValues) GetFromUnitIdOk() (*int32, bool)`

GetFromUnitIdOk returns a tuple with the FromUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUnitId

`func (o *EsgMetricValues) SetFromUnitId(v int32)`

SetFromUnitId sets FromUnitId field to given value.

### HasFromUnitId

`func (o *EsgMetricValues) HasFromUnitId() bool`

HasFromUnitId returns a boolean if a field has been set.

### GetConvertedValue

`func (o *EsgMetricValues) GetConvertedValue() float32`

GetConvertedValue returns the ConvertedValue field if non-nil, zero value otherwise.

### GetConvertedValueOk

`func (o *EsgMetricValues) GetConvertedValueOk() (*float32, bool)`

GetConvertedValueOk returns a tuple with the ConvertedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedValue

`func (o *EsgMetricValues) SetConvertedValue(v float32)`

SetConvertedValue sets ConvertedValue field to given value.

### HasConvertedValue

`func (o *EsgMetricValues) HasConvertedValue() bool`

HasConvertedValue returns a boolean if a field has been set.

### GetQualitativeQuestion

`func (o *EsgMetricValues) GetQualitativeQuestion() string`

GetQualitativeQuestion returns the QualitativeQuestion field if non-nil, zero value otherwise.

### GetQualitativeQuestionOk

`func (o *EsgMetricValues) GetQualitativeQuestionOk() (*string, bool)`

GetQualitativeQuestionOk returns a tuple with the QualitativeQuestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualitativeQuestion

`func (o *EsgMetricValues) SetQualitativeQuestion(v string)`

SetQualitativeQuestion sets QualitativeQuestion field to given value.

### HasQualitativeQuestion

`func (o *EsgMetricValues) HasQualitativeQuestion() bool`

HasQualitativeQuestion returns a boolean if a field has been set.

### GetQualitativeQuestionResponse

`func (o *EsgMetricValues) GetQualitativeQuestionResponse() string`

GetQualitativeQuestionResponse returns the QualitativeQuestionResponse field if non-nil, zero value otherwise.

### GetQualitativeQuestionResponseOk

`func (o *EsgMetricValues) GetQualitativeQuestionResponseOk() (*string, bool)`

GetQualitativeQuestionResponseOk returns a tuple with the QualitativeQuestionResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualitativeQuestionResponse

`func (o *EsgMetricValues) SetQualitativeQuestionResponse(v string)`

SetQualitativeQuestionResponse sets QualitativeQuestionResponse field to given value.

### HasQualitativeQuestionResponse

`func (o *EsgMetricValues) HasQualitativeQuestionResponse() bool`

HasQualitativeQuestionResponse returns a boolean if a field has been set.

### GetNextReviewerUserId

`func (o *EsgMetricValues) GetNextReviewerUserId() int32`

GetNextReviewerUserId returns the NextReviewerUserId field if non-nil, zero value otherwise.

### GetNextReviewerUserIdOk

`func (o *EsgMetricValues) GetNextReviewerUserIdOk() (*int32, bool)`

GetNextReviewerUserIdOk returns a tuple with the NextReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextReviewerUserId

`func (o *EsgMetricValues) SetNextReviewerUserId(v int32)`

SetNextReviewerUserId sets NextReviewerUserId field to given value.

### HasNextReviewerUserId

`func (o *EsgMetricValues) HasNextReviewerUserId() bool`

HasNextReviewerUserId returns a boolean if a field has been set.

### GetAuditSurveyId

`func (o *EsgMetricValues) GetAuditSurveyId() int32`

GetAuditSurveyId returns the AuditSurveyId field if non-nil, zero value otherwise.

### GetAuditSurveyIdOk

`func (o *EsgMetricValues) GetAuditSurveyIdOk() (*int32, bool)`

GetAuditSurveyIdOk returns a tuple with the AuditSurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSurveyId

`func (o *EsgMetricValues) SetAuditSurveyId(v int32)`

SetAuditSurveyId sets AuditSurveyId field to given value.

### HasAuditSurveyId

`func (o *EsgMetricValues) HasAuditSurveyId() bool`

HasAuditSurveyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


