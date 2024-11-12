# EsgMetricValuesPutEsgMetricValue

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

### NewEsgMetricValuesPutEsgMetricValue

`func NewEsgMetricValuesPutEsgMetricValue(esgMetricId int32, period string, value float32, status string, periodStart string, periodEnd string, ) *EsgMetricValuesPutEsgMetricValue`

NewEsgMetricValuesPutEsgMetricValue instantiates a new EsgMetricValuesPutEsgMetricValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsgMetricValuesPutEsgMetricValueWithDefaults

`func NewEsgMetricValuesPutEsgMetricValueWithDefaults() *EsgMetricValuesPutEsgMetricValue`

NewEsgMetricValuesPutEsgMetricValueWithDefaults instantiates a new EsgMetricValuesPutEsgMetricValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EsgMetricValuesPutEsgMetricValue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsgMetricValuesPutEsgMetricValue) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EsgMetricValuesPutEsgMetricValue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EsgMetricValuesPutEsgMetricValue) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EsgMetricValuesPutEsgMetricValue) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EsgMetricValuesPutEsgMetricValue) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EsgMetricValuesPutEsgMetricValue) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EsgMetricValuesPutEsgMetricValue) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EsgMetricValuesPutEsgMetricValue) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EsgMetricValuesPutEsgMetricValue) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EsgMetricValuesPutEsgMetricValue) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EsgMetricValuesPutEsgMetricValue) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetEsgMetricId

`func (o *EsgMetricValuesPutEsgMetricValue) GetEsgMetricId() int32`

GetEsgMetricId returns the EsgMetricId field if non-nil, zero value otherwise.

### GetEsgMetricIdOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetEsgMetricIdOk() (*int32, bool)`

GetEsgMetricIdOk returns a tuple with the EsgMetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricId

`func (o *EsgMetricValuesPutEsgMetricValue) SetEsgMetricId(v int32)`

SetEsgMetricId sets EsgMetricId field to given value.


### GetPeriod

`func (o *EsgMetricValuesPutEsgMetricValue) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *EsgMetricValuesPutEsgMetricValue) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### GetValue

`func (o *EsgMetricValuesPutEsgMetricValue) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EsgMetricValuesPutEsgMetricValue) SetValue(v float32)`

SetValue sets Value field to given value.


### GetTarget

`func (o *EsgMetricValuesPutEsgMetricValue) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *EsgMetricValuesPutEsgMetricValue) SetTarget(v float32)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *EsgMetricValuesPutEsgMetricValue) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetDueDate

`func (o *EsgMetricValuesPutEsgMetricValue) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *EsgMetricValuesPutEsgMetricValue) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *EsgMetricValuesPutEsgMetricValue) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetStatus

`func (o *EsgMetricValuesPutEsgMetricValue) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EsgMetricValuesPutEsgMetricValue) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetNotes

`func (o *EsgMetricValuesPutEsgMetricValue) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *EsgMetricValuesPutEsgMetricValue) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *EsgMetricValuesPutEsgMetricValue) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetSource

`func (o *EsgMetricValuesPutEsgMetricValue) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EsgMetricValuesPutEsgMetricValue) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *EsgMetricValuesPutEsgMetricValue) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetMissingFilesReason

`func (o *EsgMetricValuesPutEsgMetricValue) GetMissingFilesReason() string`

GetMissingFilesReason returns the MissingFilesReason field if non-nil, zero value otherwise.

### GetMissingFilesReasonOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetMissingFilesReasonOk() (*string, bool)`

GetMissingFilesReasonOk returns a tuple with the MissingFilesReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingFilesReason

`func (o *EsgMetricValuesPutEsgMetricValue) SetMissingFilesReason(v string)`

SetMissingFilesReason sets MissingFilesReason field to given value.

### HasMissingFilesReason

`func (o *EsgMetricValuesPutEsgMetricValue) HasMissingFilesReason() bool`

HasMissingFilesReason returns a boolean if a field has been set.

### GetSubmittedByUserId

`func (o *EsgMetricValuesPutEsgMetricValue) GetSubmittedByUserId() int32`

GetSubmittedByUserId returns the SubmittedByUserId field if non-nil, zero value otherwise.

### GetSubmittedByUserIdOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetSubmittedByUserIdOk() (*int32, bool)`

GetSubmittedByUserIdOk returns a tuple with the SubmittedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedByUserId

`func (o *EsgMetricValuesPutEsgMetricValue) SetSubmittedByUserId(v int32)`

SetSubmittedByUserId sets SubmittedByUserId field to given value.

### HasSubmittedByUserId

`func (o *EsgMetricValuesPutEsgMetricValue) HasSubmittedByUserId() bool`

HasSubmittedByUserId returns a boolean if a field has been set.

### GetEsgMetricAuditableEntityId

`func (o *EsgMetricValuesPutEsgMetricValue) GetEsgMetricAuditableEntityId() int32`

GetEsgMetricAuditableEntityId returns the EsgMetricAuditableEntityId field if non-nil, zero value otherwise.

### GetEsgMetricAuditableEntityIdOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetEsgMetricAuditableEntityIdOk() (*int32, bool)`

GetEsgMetricAuditableEntityIdOk returns a tuple with the EsgMetricAuditableEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricAuditableEntityId

`func (o *EsgMetricValuesPutEsgMetricValue) SetEsgMetricAuditableEntityId(v int32)`

SetEsgMetricAuditableEntityId sets EsgMetricAuditableEntityId field to given value.

### HasEsgMetricAuditableEntityId

`func (o *EsgMetricValuesPutEsgMetricValue) HasEsgMetricAuditableEntityId() bool`

HasEsgMetricAuditableEntityId returns a boolean if a field has been set.

### GetPeriodStart

`func (o *EsgMetricValuesPutEsgMetricValue) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *EsgMetricValuesPutEsgMetricValue) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *EsgMetricValuesPutEsgMetricValue) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *EsgMetricValuesPutEsgMetricValue) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetCertifiedByUserId

`func (o *EsgMetricValuesPutEsgMetricValue) GetCertifiedByUserId() int32`

GetCertifiedByUserId returns the CertifiedByUserId field if non-nil, zero value otherwise.

### GetCertifiedByUserIdOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetCertifiedByUserIdOk() (*int32, bool)`

GetCertifiedByUserIdOk returns a tuple with the CertifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifiedByUserId

`func (o *EsgMetricValuesPutEsgMetricValue) SetCertifiedByUserId(v int32)`

SetCertifiedByUserId sets CertifiedByUserId field to given value.

### HasCertifiedByUserId

`func (o *EsgMetricValuesPutEsgMetricValue) HasCertifiedByUserId() bool`

HasCertifiedByUserId returns a boolean if a field has been set.

### GetFromUnitId

`func (o *EsgMetricValuesPutEsgMetricValue) GetFromUnitId() int32`

GetFromUnitId returns the FromUnitId field if non-nil, zero value otherwise.

### GetFromUnitIdOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetFromUnitIdOk() (*int32, bool)`

GetFromUnitIdOk returns a tuple with the FromUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUnitId

`func (o *EsgMetricValuesPutEsgMetricValue) SetFromUnitId(v int32)`

SetFromUnitId sets FromUnitId field to given value.

### HasFromUnitId

`func (o *EsgMetricValuesPutEsgMetricValue) HasFromUnitId() bool`

HasFromUnitId returns a boolean if a field has been set.

### GetConvertedValue

`func (o *EsgMetricValuesPutEsgMetricValue) GetConvertedValue() float32`

GetConvertedValue returns the ConvertedValue field if non-nil, zero value otherwise.

### GetConvertedValueOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetConvertedValueOk() (*float32, bool)`

GetConvertedValueOk returns a tuple with the ConvertedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedValue

`func (o *EsgMetricValuesPutEsgMetricValue) SetConvertedValue(v float32)`

SetConvertedValue sets ConvertedValue field to given value.

### HasConvertedValue

`func (o *EsgMetricValuesPutEsgMetricValue) HasConvertedValue() bool`

HasConvertedValue returns a boolean if a field has been set.

### GetQualitativeQuestion

`func (o *EsgMetricValuesPutEsgMetricValue) GetQualitativeQuestion() string`

GetQualitativeQuestion returns the QualitativeQuestion field if non-nil, zero value otherwise.

### GetQualitativeQuestionOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetQualitativeQuestionOk() (*string, bool)`

GetQualitativeQuestionOk returns a tuple with the QualitativeQuestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualitativeQuestion

`func (o *EsgMetricValuesPutEsgMetricValue) SetQualitativeQuestion(v string)`

SetQualitativeQuestion sets QualitativeQuestion field to given value.

### HasQualitativeQuestion

`func (o *EsgMetricValuesPutEsgMetricValue) HasQualitativeQuestion() bool`

HasQualitativeQuestion returns a boolean if a field has been set.

### GetQualitativeQuestionResponse

`func (o *EsgMetricValuesPutEsgMetricValue) GetQualitativeQuestionResponse() string`

GetQualitativeQuestionResponse returns the QualitativeQuestionResponse field if non-nil, zero value otherwise.

### GetQualitativeQuestionResponseOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetQualitativeQuestionResponseOk() (*string, bool)`

GetQualitativeQuestionResponseOk returns a tuple with the QualitativeQuestionResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualitativeQuestionResponse

`func (o *EsgMetricValuesPutEsgMetricValue) SetQualitativeQuestionResponse(v string)`

SetQualitativeQuestionResponse sets QualitativeQuestionResponse field to given value.

### HasQualitativeQuestionResponse

`func (o *EsgMetricValuesPutEsgMetricValue) HasQualitativeQuestionResponse() bool`

HasQualitativeQuestionResponse returns a boolean if a field has been set.

### GetNextReviewerUserId

`func (o *EsgMetricValuesPutEsgMetricValue) GetNextReviewerUserId() int32`

GetNextReviewerUserId returns the NextReviewerUserId field if non-nil, zero value otherwise.

### GetNextReviewerUserIdOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetNextReviewerUserIdOk() (*int32, bool)`

GetNextReviewerUserIdOk returns a tuple with the NextReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextReviewerUserId

`func (o *EsgMetricValuesPutEsgMetricValue) SetNextReviewerUserId(v int32)`

SetNextReviewerUserId sets NextReviewerUserId field to given value.

### HasNextReviewerUserId

`func (o *EsgMetricValuesPutEsgMetricValue) HasNextReviewerUserId() bool`

HasNextReviewerUserId returns a boolean if a field has been set.

### GetAuditSurveyId

`func (o *EsgMetricValuesPutEsgMetricValue) GetAuditSurveyId() int32`

GetAuditSurveyId returns the AuditSurveyId field if non-nil, zero value otherwise.

### GetAuditSurveyIdOk

`func (o *EsgMetricValuesPutEsgMetricValue) GetAuditSurveyIdOk() (*int32, bool)`

GetAuditSurveyIdOk returns a tuple with the AuditSurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSurveyId

`func (o *EsgMetricValuesPutEsgMetricValue) SetAuditSurveyId(v int32)`

SetAuditSurveyId sets AuditSurveyId field to given value.

### HasAuditSurveyId

`func (o *EsgMetricValuesPutEsgMetricValue) HasAuditSurveyId() bool`

HasAuditSurveyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


