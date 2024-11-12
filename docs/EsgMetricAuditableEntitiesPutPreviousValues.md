# EsgMetricAuditableEntitiesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] 
**EsgMetricId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_metrics.id&#x60;.&lt;fk table&#x3D;&#39;esg_metrics&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AuditableEntityId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;auditable_entities.id&#x60;.&lt;fk table&#x3D;&#39;auditable_entities&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ReviewerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**LastCertifiedValue** | Pointer to **float32** |  | [optional] 
**LastCertifiedNotes** | Pointer to **string** |  | [optional] 
**LastCertifiedPeriodStart** | Pointer to **string** |  | [optional] 
**LastCertifiedPeriodEnd** | Pointer to **string** |  | [optional] 
**LastPeriodRequested** | Pointer to **string** |  | [optional] 
**CurrentYearToDateValue** | Pointer to **float32** |  | [optional] [default to 0]
**PreviousYearToDateValue** | Pointer to **float32** |  | [optional] [default to 0]
**Question** | Pointer to **string** |  | [optional] 
**LastCertifiedResponse** | Pointer to **string** |  | [optional] 
**TwoPriorYearToDateValue** | Pointer to **float32** |  | [optional] 

## Methods

### NewEsgMetricAuditableEntitiesPutPreviousValues

`func NewEsgMetricAuditableEntitiesPutPreviousValues() *EsgMetricAuditableEntitiesPutPreviousValues`

NewEsgMetricAuditableEntitiesPutPreviousValues instantiates a new EsgMetricAuditableEntitiesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsgMetricAuditableEntitiesPutPreviousValuesWithDefaults

`func NewEsgMetricAuditableEntitiesPutPreviousValuesWithDefaults() *EsgMetricAuditableEntitiesPutPreviousValues`

NewEsgMetricAuditableEntitiesPutPreviousValuesWithDefaults instantiates a new EsgMetricAuditableEntitiesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetUid

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetEsgMetricId

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetEsgMetricId() int32`

GetEsgMetricId returns the EsgMetricId field if non-nil, zero value otherwise.

### GetEsgMetricIdOk

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetEsgMetricIdOk() (*int32, bool)`

GetEsgMetricIdOk returns a tuple with the EsgMetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricId

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) SetEsgMetricId(v int32)`

SetEsgMetricId sets EsgMetricId field to given value.

### HasEsgMetricId

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) HasEsgMetricId() bool`

HasEsgMetricId returns a boolean if a field has been set.

### GetAuditableEntityId

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetAuditableEntityId() int32`

GetAuditableEntityId returns the AuditableEntityId field if non-nil, zero value otherwise.

### GetAuditableEntityIdOk

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetAuditableEntityIdOk() (*int32, bool)`

GetAuditableEntityIdOk returns a tuple with the AuditableEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditableEntityId

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) SetAuditableEntityId(v int32)`

SetAuditableEntityId sets AuditableEntityId field to given value.

### HasAuditableEntityId

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) HasAuditableEntityId() bool`

HasAuditableEntityId returns a boolean if a field has been set.

### GetReviewerUserId

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetReviewerUserId() int32`

GetReviewerUserId returns the ReviewerUserId field if non-nil, zero value otherwise.

### GetReviewerUserIdOk

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetReviewerUserIdOk() (*int32, bool)`

GetReviewerUserIdOk returns a tuple with the ReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerUserId

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) SetReviewerUserId(v int32)`

SetReviewerUserId sets ReviewerUserId field to given value.

### HasReviewerUserId

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) HasReviewerUserId() bool`

HasReviewerUserId returns a boolean if a field has been set.

### GetLastCertifiedValue

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetLastCertifiedValue() float32`

GetLastCertifiedValue returns the LastCertifiedValue field if non-nil, zero value otherwise.

### GetLastCertifiedValueOk

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetLastCertifiedValueOk() (*float32, bool)`

GetLastCertifiedValueOk returns a tuple with the LastCertifiedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCertifiedValue

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) SetLastCertifiedValue(v float32)`

SetLastCertifiedValue sets LastCertifiedValue field to given value.

### HasLastCertifiedValue

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) HasLastCertifiedValue() bool`

HasLastCertifiedValue returns a boolean if a field has been set.

### GetLastCertifiedNotes

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetLastCertifiedNotes() string`

GetLastCertifiedNotes returns the LastCertifiedNotes field if non-nil, zero value otherwise.

### GetLastCertifiedNotesOk

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetLastCertifiedNotesOk() (*string, bool)`

GetLastCertifiedNotesOk returns a tuple with the LastCertifiedNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCertifiedNotes

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) SetLastCertifiedNotes(v string)`

SetLastCertifiedNotes sets LastCertifiedNotes field to given value.

### HasLastCertifiedNotes

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) HasLastCertifiedNotes() bool`

HasLastCertifiedNotes returns a boolean if a field has been set.

### GetLastCertifiedPeriodStart

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetLastCertifiedPeriodStart() string`

GetLastCertifiedPeriodStart returns the LastCertifiedPeriodStart field if non-nil, zero value otherwise.

### GetLastCertifiedPeriodStartOk

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetLastCertifiedPeriodStartOk() (*string, bool)`

GetLastCertifiedPeriodStartOk returns a tuple with the LastCertifiedPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCertifiedPeriodStart

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) SetLastCertifiedPeriodStart(v string)`

SetLastCertifiedPeriodStart sets LastCertifiedPeriodStart field to given value.

### HasLastCertifiedPeriodStart

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) HasLastCertifiedPeriodStart() bool`

HasLastCertifiedPeriodStart returns a boolean if a field has been set.

### GetLastCertifiedPeriodEnd

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetLastCertifiedPeriodEnd() string`

GetLastCertifiedPeriodEnd returns the LastCertifiedPeriodEnd field if non-nil, zero value otherwise.

### GetLastCertifiedPeriodEndOk

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetLastCertifiedPeriodEndOk() (*string, bool)`

GetLastCertifiedPeriodEndOk returns a tuple with the LastCertifiedPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCertifiedPeriodEnd

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) SetLastCertifiedPeriodEnd(v string)`

SetLastCertifiedPeriodEnd sets LastCertifiedPeriodEnd field to given value.

### HasLastCertifiedPeriodEnd

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) HasLastCertifiedPeriodEnd() bool`

HasLastCertifiedPeriodEnd returns a boolean if a field has been set.

### GetLastPeriodRequested

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetLastPeriodRequested() string`

GetLastPeriodRequested returns the LastPeriodRequested field if non-nil, zero value otherwise.

### GetLastPeriodRequestedOk

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetLastPeriodRequestedOk() (*string, bool)`

GetLastPeriodRequestedOk returns a tuple with the LastPeriodRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPeriodRequested

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) SetLastPeriodRequested(v string)`

SetLastPeriodRequested sets LastPeriodRequested field to given value.

### HasLastPeriodRequested

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) HasLastPeriodRequested() bool`

HasLastPeriodRequested returns a boolean if a field has been set.

### GetCurrentYearToDateValue

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetCurrentYearToDateValue() float32`

GetCurrentYearToDateValue returns the CurrentYearToDateValue field if non-nil, zero value otherwise.

### GetCurrentYearToDateValueOk

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetCurrentYearToDateValueOk() (*float32, bool)`

GetCurrentYearToDateValueOk returns a tuple with the CurrentYearToDateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentYearToDateValue

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) SetCurrentYearToDateValue(v float32)`

SetCurrentYearToDateValue sets CurrentYearToDateValue field to given value.

### HasCurrentYearToDateValue

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) HasCurrentYearToDateValue() bool`

HasCurrentYearToDateValue returns a boolean if a field has been set.

### GetPreviousYearToDateValue

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetPreviousYearToDateValue() float32`

GetPreviousYearToDateValue returns the PreviousYearToDateValue field if non-nil, zero value otherwise.

### GetPreviousYearToDateValueOk

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetPreviousYearToDateValueOk() (*float32, bool)`

GetPreviousYearToDateValueOk returns a tuple with the PreviousYearToDateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousYearToDateValue

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) SetPreviousYearToDateValue(v float32)`

SetPreviousYearToDateValue sets PreviousYearToDateValue field to given value.

### HasPreviousYearToDateValue

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) HasPreviousYearToDateValue() bool`

HasPreviousYearToDateValue returns a boolean if a field has been set.

### GetQuestion

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetLastCertifiedResponse

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetLastCertifiedResponse() string`

GetLastCertifiedResponse returns the LastCertifiedResponse field if non-nil, zero value otherwise.

### GetLastCertifiedResponseOk

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetLastCertifiedResponseOk() (*string, bool)`

GetLastCertifiedResponseOk returns a tuple with the LastCertifiedResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCertifiedResponse

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) SetLastCertifiedResponse(v string)`

SetLastCertifiedResponse sets LastCertifiedResponse field to given value.

### HasLastCertifiedResponse

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) HasLastCertifiedResponse() bool`

HasLastCertifiedResponse returns a boolean if a field has been set.

### GetTwoPriorYearToDateValue

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetTwoPriorYearToDateValue() float32`

GetTwoPriorYearToDateValue returns the TwoPriorYearToDateValue field if non-nil, zero value otherwise.

### GetTwoPriorYearToDateValueOk

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) GetTwoPriorYearToDateValueOk() (*float32, bool)`

GetTwoPriorYearToDateValueOk returns a tuple with the TwoPriorYearToDateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoPriorYearToDateValue

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) SetTwoPriorYearToDateValue(v float32)`

SetTwoPriorYearToDateValue sets TwoPriorYearToDateValue field to given value.

### HasTwoPriorYearToDateValue

`func (o *EsgMetricAuditableEntitiesPutPreviousValues) HasTwoPriorYearToDateValue() bool`

HasTwoPriorYearToDateValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


