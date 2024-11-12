# EsgMetricValuesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**EsgMetricId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_metrics.id&#x60;.&lt;fk table&#x3D;&#39;esg_metrics&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 
**Target** | Pointer to **float32** |  | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**MissingFilesReason** | Pointer to **string** |  | [optional] 
**SubmittedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**EsgMetricAuditableEntityId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_metric_auditable_entities.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_auditable_entities&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**PeriodStart** | Pointer to **string** |  | [optional] 
**PeriodEnd** | Pointer to **string** |  | [optional] 
**CertifiedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FromUnitId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_metric_unit_options.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_unit_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ConvertedValue** | Pointer to **float32** |  | [optional] 
**QualitativeQuestion** | Pointer to **string** |  | [optional] 
**QualitativeQuestionResponse** | Pointer to **string** |  | [optional] 
**NextReviewerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AuditSurveyId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_surveys.id&#x60;.&lt;fk table&#x3D;&#39;audit_surveys&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 

## Methods

### NewEsgMetricValuesPutPreviousValues

`func NewEsgMetricValuesPutPreviousValues() *EsgMetricValuesPutPreviousValues`

NewEsgMetricValuesPutPreviousValues instantiates a new EsgMetricValuesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsgMetricValuesPutPreviousValuesWithDefaults

`func NewEsgMetricValuesPutPreviousValuesWithDefaults() *EsgMetricValuesPutPreviousValues`

NewEsgMetricValuesPutPreviousValuesWithDefaults instantiates a new EsgMetricValuesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EsgMetricValuesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsgMetricValuesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsgMetricValuesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EsgMetricValuesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EsgMetricValuesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EsgMetricValuesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EsgMetricValuesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EsgMetricValuesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EsgMetricValuesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EsgMetricValuesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EsgMetricValuesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EsgMetricValuesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EsgMetricValuesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EsgMetricValuesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EsgMetricValuesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EsgMetricValuesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetEsgMetricId

`func (o *EsgMetricValuesPutPreviousValues) GetEsgMetricId() int32`

GetEsgMetricId returns the EsgMetricId field if non-nil, zero value otherwise.

### GetEsgMetricIdOk

`func (o *EsgMetricValuesPutPreviousValues) GetEsgMetricIdOk() (*int32, bool)`

GetEsgMetricIdOk returns a tuple with the EsgMetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricId

`func (o *EsgMetricValuesPutPreviousValues) SetEsgMetricId(v int32)`

SetEsgMetricId sets EsgMetricId field to given value.

### HasEsgMetricId

`func (o *EsgMetricValuesPutPreviousValues) HasEsgMetricId() bool`

HasEsgMetricId returns a boolean if a field has been set.

### GetPeriod

`func (o *EsgMetricValuesPutPreviousValues) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *EsgMetricValuesPutPreviousValues) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *EsgMetricValuesPutPreviousValues) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *EsgMetricValuesPutPreviousValues) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetValue

`func (o *EsgMetricValuesPutPreviousValues) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EsgMetricValuesPutPreviousValues) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EsgMetricValuesPutPreviousValues) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *EsgMetricValuesPutPreviousValues) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetTarget

`func (o *EsgMetricValuesPutPreviousValues) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *EsgMetricValuesPutPreviousValues) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *EsgMetricValuesPutPreviousValues) SetTarget(v float32)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *EsgMetricValuesPutPreviousValues) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetDueDate

`func (o *EsgMetricValuesPutPreviousValues) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *EsgMetricValuesPutPreviousValues) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *EsgMetricValuesPutPreviousValues) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *EsgMetricValuesPutPreviousValues) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetStatus

`func (o *EsgMetricValuesPutPreviousValues) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EsgMetricValuesPutPreviousValues) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EsgMetricValuesPutPreviousValues) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EsgMetricValuesPutPreviousValues) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNotes

`func (o *EsgMetricValuesPutPreviousValues) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *EsgMetricValuesPutPreviousValues) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *EsgMetricValuesPutPreviousValues) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *EsgMetricValuesPutPreviousValues) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetSource

`func (o *EsgMetricValuesPutPreviousValues) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EsgMetricValuesPutPreviousValues) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EsgMetricValuesPutPreviousValues) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *EsgMetricValuesPutPreviousValues) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetMissingFilesReason

`func (o *EsgMetricValuesPutPreviousValues) GetMissingFilesReason() string`

GetMissingFilesReason returns the MissingFilesReason field if non-nil, zero value otherwise.

### GetMissingFilesReasonOk

`func (o *EsgMetricValuesPutPreviousValues) GetMissingFilesReasonOk() (*string, bool)`

GetMissingFilesReasonOk returns a tuple with the MissingFilesReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingFilesReason

`func (o *EsgMetricValuesPutPreviousValues) SetMissingFilesReason(v string)`

SetMissingFilesReason sets MissingFilesReason field to given value.

### HasMissingFilesReason

`func (o *EsgMetricValuesPutPreviousValues) HasMissingFilesReason() bool`

HasMissingFilesReason returns a boolean if a field has been set.

### GetSubmittedByUserId

`func (o *EsgMetricValuesPutPreviousValues) GetSubmittedByUserId() int32`

GetSubmittedByUserId returns the SubmittedByUserId field if non-nil, zero value otherwise.

### GetSubmittedByUserIdOk

`func (o *EsgMetricValuesPutPreviousValues) GetSubmittedByUserIdOk() (*int32, bool)`

GetSubmittedByUserIdOk returns a tuple with the SubmittedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedByUserId

`func (o *EsgMetricValuesPutPreviousValues) SetSubmittedByUserId(v int32)`

SetSubmittedByUserId sets SubmittedByUserId field to given value.

### HasSubmittedByUserId

`func (o *EsgMetricValuesPutPreviousValues) HasSubmittedByUserId() bool`

HasSubmittedByUserId returns a boolean if a field has been set.

### GetEsgMetricAuditableEntityId

`func (o *EsgMetricValuesPutPreviousValues) GetEsgMetricAuditableEntityId() int32`

GetEsgMetricAuditableEntityId returns the EsgMetricAuditableEntityId field if non-nil, zero value otherwise.

### GetEsgMetricAuditableEntityIdOk

`func (o *EsgMetricValuesPutPreviousValues) GetEsgMetricAuditableEntityIdOk() (*int32, bool)`

GetEsgMetricAuditableEntityIdOk returns a tuple with the EsgMetricAuditableEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricAuditableEntityId

`func (o *EsgMetricValuesPutPreviousValues) SetEsgMetricAuditableEntityId(v int32)`

SetEsgMetricAuditableEntityId sets EsgMetricAuditableEntityId field to given value.

### HasEsgMetricAuditableEntityId

`func (o *EsgMetricValuesPutPreviousValues) HasEsgMetricAuditableEntityId() bool`

HasEsgMetricAuditableEntityId returns a boolean if a field has been set.

### GetPeriodStart

`func (o *EsgMetricValuesPutPreviousValues) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *EsgMetricValuesPutPreviousValues) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *EsgMetricValuesPutPreviousValues) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *EsgMetricValuesPutPreviousValues) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *EsgMetricValuesPutPreviousValues) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *EsgMetricValuesPutPreviousValues) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *EsgMetricValuesPutPreviousValues) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *EsgMetricValuesPutPreviousValues) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetCertifiedByUserId

`func (o *EsgMetricValuesPutPreviousValues) GetCertifiedByUserId() int32`

GetCertifiedByUserId returns the CertifiedByUserId field if non-nil, zero value otherwise.

### GetCertifiedByUserIdOk

`func (o *EsgMetricValuesPutPreviousValues) GetCertifiedByUserIdOk() (*int32, bool)`

GetCertifiedByUserIdOk returns a tuple with the CertifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifiedByUserId

`func (o *EsgMetricValuesPutPreviousValues) SetCertifiedByUserId(v int32)`

SetCertifiedByUserId sets CertifiedByUserId field to given value.

### HasCertifiedByUserId

`func (o *EsgMetricValuesPutPreviousValues) HasCertifiedByUserId() bool`

HasCertifiedByUserId returns a boolean if a field has been set.

### GetFromUnitId

`func (o *EsgMetricValuesPutPreviousValues) GetFromUnitId() int32`

GetFromUnitId returns the FromUnitId field if non-nil, zero value otherwise.

### GetFromUnitIdOk

`func (o *EsgMetricValuesPutPreviousValues) GetFromUnitIdOk() (*int32, bool)`

GetFromUnitIdOk returns a tuple with the FromUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUnitId

`func (o *EsgMetricValuesPutPreviousValues) SetFromUnitId(v int32)`

SetFromUnitId sets FromUnitId field to given value.

### HasFromUnitId

`func (o *EsgMetricValuesPutPreviousValues) HasFromUnitId() bool`

HasFromUnitId returns a boolean if a field has been set.

### GetConvertedValue

`func (o *EsgMetricValuesPutPreviousValues) GetConvertedValue() float32`

GetConvertedValue returns the ConvertedValue field if non-nil, zero value otherwise.

### GetConvertedValueOk

`func (o *EsgMetricValuesPutPreviousValues) GetConvertedValueOk() (*float32, bool)`

GetConvertedValueOk returns a tuple with the ConvertedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedValue

`func (o *EsgMetricValuesPutPreviousValues) SetConvertedValue(v float32)`

SetConvertedValue sets ConvertedValue field to given value.

### HasConvertedValue

`func (o *EsgMetricValuesPutPreviousValues) HasConvertedValue() bool`

HasConvertedValue returns a boolean if a field has been set.

### GetQualitativeQuestion

`func (o *EsgMetricValuesPutPreviousValues) GetQualitativeQuestion() string`

GetQualitativeQuestion returns the QualitativeQuestion field if non-nil, zero value otherwise.

### GetQualitativeQuestionOk

`func (o *EsgMetricValuesPutPreviousValues) GetQualitativeQuestionOk() (*string, bool)`

GetQualitativeQuestionOk returns a tuple with the QualitativeQuestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualitativeQuestion

`func (o *EsgMetricValuesPutPreviousValues) SetQualitativeQuestion(v string)`

SetQualitativeQuestion sets QualitativeQuestion field to given value.

### HasQualitativeQuestion

`func (o *EsgMetricValuesPutPreviousValues) HasQualitativeQuestion() bool`

HasQualitativeQuestion returns a boolean if a field has been set.

### GetQualitativeQuestionResponse

`func (o *EsgMetricValuesPutPreviousValues) GetQualitativeQuestionResponse() string`

GetQualitativeQuestionResponse returns the QualitativeQuestionResponse field if non-nil, zero value otherwise.

### GetQualitativeQuestionResponseOk

`func (o *EsgMetricValuesPutPreviousValues) GetQualitativeQuestionResponseOk() (*string, bool)`

GetQualitativeQuestionResponseOk returns a tuple with the QualitativeQuestionResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualitativeQuestionResponse

`func (o *EsgMetricValuesPutPreviousValues) SetQualitativeQuestionResponse(v string)`

SetQualitativeQuestionResponse sets QualitativeQuestionResponse field to given value.

### HasQualitativeQuestionResponse

`func (o *EsgMetricValuesPutPreviousValues) HasQualitativeQuestionResponse() bool`

HasQualitativeQuestionResponse returns a boolean if a field has been set.

### GetNextReviewerUserId

`func (o *EsgMetricValuesPutPreviousValues) GetNextReviewerUserId() int32`

GetNextReviewerUserId returns the NextReviewerUserId field if non-nil, zero value otherwise.

### GetNextReviewerUserIdOk

`func (o *EsgMetricValuesPutPreviousValues) GetNextReviewerUserIdOk() (*int32, bool)`

GetNextReviewerUserIdOk returns a tuple with the NextReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextReviewerUserId

`func (o *EsgMetricValuesPutPreviousValues) SetNextReviewerUserId(v int32)`

SetNextReviewerUserId sets NextReviewerUserId field to given value.

### HasNextReviewerUserId

`func (o *EsgMetricValuesPutPreviousValues) HasNextReviewerUserId() bool`

HasNextReviewerUserId returns a boolean if a field has been set.

### GetAuditSurveyId

`func (o *EsgMetricValuesPutPreviousValues) GetAuditSurveyId() int32`

GetAuditSurveyId returns the AuditSurveyId field if non-nil, zero value otherwise.

### GetAuditSurveyIdOk

`func (o *EsgMetricValuesPutPreviousValues) GetAuditSurveyIdOk() (*int32, bool)`

GetAuditSurveyIdOk returns a tuple with the AuditSurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSurveyId

`func (o *EsgMetricValuesPutPreviousValues) SetAuditSurveyId(v int32)`

SetAuditSurveyId sets AuditSurveyId field to given value.

### HasAuditSurveyId

`func (o *EsgMetricValuesPutPreviousValues) HasAuditSurveyId() bool`

HasAuditSurveyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


