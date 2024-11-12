# EsgMetricAuditableEntities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] 
**EsgMetricId** | **int32** | Note: This is a Foreign Key to &#x60;esg_metrics.id&#x60;.&lt;fk table&#x3D;&#39;esg_metrics&#39; column&#x3D;&#39;id&#39;/&gt; | 
**AuditableEntityId** | **int32** | Note: This is a Foreign Key to &#x60;auditable_entities.id&#x60;.&lt;fk table&#x3D;&#39;auditable_entities&#39; column&#x3D;&#39;id&#39;/&gt; | 
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

### NewEsgMetricAuditableEntities

`func NewEsgMetricAuditableEntities(esgMetricId int32, auditableEntityId int32, ) *EsgMetricAuditableEntities`

NewEsgMetricAuditableEntities instantiates a new EsgMetricAuditableEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsgMetricAuditableEntitiesWithDefaults

`func NewEsgMetricAuditableEntitiesWithDefaults() *EsgMetricAuditableEntities`

NewEsgMetricAuditableEntitiesWithDefaults instantiates a new EsgMetricAuditableEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EsgMetricAuditableEntities) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsgMetricAuditableEntities) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsgMetricAuditableEntities) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EsgMetricAuditableEntities) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EsgMetricAuditableEntities) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EsgMetricAuditableEntities) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EsgMetricAuditableEntities) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EsgMetricAuditableEntities) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EsgMetricAuditableEntities) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EsgMetricAuditableEntities) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EsgMetricAuditableEntities) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EsgMetricAuditableEntities) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EsgMetricAuditableEntities) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EsgMetricAuditableEntities) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EsgMetricAuditableEntities) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EsgMetricAuditableEntities) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetUid

`func (o *EsgMetricAuditableEntities) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *EsgMetricAuditableEntities) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *EsgMetricAuditableEntities) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *EsgMetricAuditableEntities) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetEsgMetricId

`func (o *EsgMetricAuditableEntities) GetEsgMetricId() int32`

GetEsgMetricId returns the EsgMetricId field if non-nil, zero value otherwise.

### GetEsgMetricIdOk

`func (o *EsgMetricAuditableEntities) GetEsgMetricIdOk() (*int32, bool)`

GetEsgMetricIdOk returns a tuple with the EsgMetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricId

`func (o *EsgMetricAuditableEntities) SetEsgMetricId(v int32)`

SetEsgMetricId sets EsgMetricId field to given value.


### GetAuditableEntityId

`func (o *EsgMetricAuditableEntities) GetAuditableEntityId() int32`

GetAuditableEntityId returns the AuditableEntityId field if non-nil, zero value otherwise.

### GetAuditableEntityIdOk

`func (o *EsgMetricAuditableEntities) GetAuditableEntityIdOk() (*int32, bool)`

GetAuditableEntityIdOk returns a tuple with the AuditableEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditableEntityId

`func (o *EsgMetricAuditableEntities) SetAuditableEntityId(v int32)`

SetAuditableEntityId sets AuditableEntityId field to given value.


### GetReviewerUserId

`func (o *EsgMetricAuditableEntities) GetReviewerUserId() int32`

GetReviewerUserId returns the ReviewerUserId field if non-nil, zero value otherwise.

### GetReviewerUserIdOk

`func (o *EsgMetricAuditableEntities) GetReviewerUserIdOk() (*int32, bool)`

GetReviewerUserIdOk returns a tuple with the ReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerUserId

`func (o *EsgMetricAuditableEntities) SetReviewerUserId(v int32)`

SetReviewerUserId sets ReviewerUserId field to given value.

### HasReviewerUserId

`func (o *EsgMetricAuditableEntities) HasReviewerUserId() bool`

HasReviewerUserId returns a boolean if a field has been set.

### GetLastCertifiedValue

`func (o *EsgMetricAuditableEntities) GetLastCertifiedValue() float32`

GetLastCertifiedValue returns the LastCertifiedValue field if non-nil, zero value otherwise.

### GetLastCertifiedValueOk

`func (o *EsgMetricAuditableEntities) GetLastCertifiedValueOk() (*float32, bool)`

GetLastCertifiedValueOk returns a tuple with the LastCertifiedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCertifiedValue

`func (o *EsgMetricAuditableEntities) SetLastCertifiedValue(v float32)`

SetLastCertifiedValue sets LastCertifiedValue field to given value.

### HasLastCertifiedValue

`func (o *EsgMetricAuditableEntities) HasLastCertifiedValue() bool`

HasLastCertifiedValue returns a boolean if a field has been set.

### GetLastCertifiedNotes

`func (o *EsgMetricAuditableEntities) GetLastCertifiedNotes() string`

GetLastCertifiedNotes returns the LastCertifiedNotes field if non-nil, zero value otherwise.

### GetLastCertifiedNotesOk

`func (o *EsgMetricAuditableEntities) GetLastCertifiedNotesOk() (*string, bool)`

GetLastCertifiedNotesOk returns a tuple with the LastCertifiedNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCertifiedNotes

`func (o *EsgMetricAuditableEntities) SetLastCertifiedNotes(v string)`

SetLastCertifiedNotes sets LastCertifiedNotes field to given value.

### HasLastCertifiedNotes

`func (o *EsgMetricAuditableEntities) HasLastCertifiedNotes() bool`

HasLastCertifiedNotes returns a boolean if a field has been set.

### GetLastCertifiedPeriodStart

`func (o *EsgMetricAuditableEntities) GetLastCertifiedPeriodStart() string`

GetLastCertifiedPeriodStart returns the LastCertifiedPeriodStart field if non-nil, zero value otherwise.

### GetLastCertifiedPeriodStartOk

`func (o *EsgMetricAuditableEntities) GetLastCertifiedPeriodStartOk() (*string, bool)`

GetLastCertifiedPeriodStartOk returns a tuple with the LastCertifiedPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCertifiedPeriodStart

`func (o *EsgMetricAuditableEntities) SetLastCertifiedPeriodStart(v string)`

SetLastCertifiedPeriodStart sets LastCertifiedPeriodStart field to given value.

### HasLastCertifiedPeriodStart

`func (o *EsgMetricAuditableEntities) HasLastCertifiedPeriodStart() bool`

HasLastCertifiedPeriodStart returns a boolean if a field has been set.

### GetLastCertifiedPeriodEnd

`func (o *EsgMetricAuditableEntities) GetLastCertifiedPeriodEnd() string`

GetLastCertifiedPeriodEnd returns the LastCertifiedPeriodEnd field if non-nil, zero value otherwise.

### GetLastCertifiedPeriodEndOk

`func (o *EsgMetricAuditableEntities) GetLastCertifiedPeriodEndOk() (*string, bool)`

GetLastCertifiedPeriodEndOk returns a tuple with the LastCertifiedPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCertifiedPeriodEnd

`func (o *EsgMetricAuditableEntities) SetLastCertifiedPeriodEnd(v string)`

SetLastCertifiedPeriodEnd sets LastCertifiedPeriodEnd field to given value.

### HasLastCertifiedPeriodEnd

`func (o *EsgMetricAuditableEntities) HasLastCertifiedPeriodEnd() bool`

HasLastCertifiedPeriodEnd returns a boolean if a field has been set.

### GetLastPeriodRequested

`func (o *EsgMetricAuditableEntities) GetLastPeriodRequested() string`

GetLastPeriodRequested returns the LastPeriodRequested field if non-nil, zero value otherwise.

### GetLastPeriodRequestedOk

`func (o *EsgMetricAuditableEntities) GetLastPeriodRequestedOk() (*string, bool)`

GetLastPeriodRequestedOk returns a tuple with the LastPeriodRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPeriodRequested

`func (o *EsgMetricAuditableEntities) SetLastPeriodRequested(v string)`

SetLastPeriodRequested sets LastPeriodRequested field to given value.

### HasLastPeriodRequested

`func (o *EsgMetricAuditableEntities) HasLastPeriodRequested() bool`

HasLastPeriodRequested returns a boolean if a field has been set.

### GetCurrentYearToDateValue

`func (o *EsgMetricAuditableEntities) GetCurrentYearToDateValue() float32`

GetCurrentYearToDateValue returns the CurrentYearToDateValue field if non-nil, zero value otherwise.

### GetCurrentYearToDateValueOk

`func (o *EsgMetricAuditableEntities) GetCurrentYearToDateValueOk() (*float32, bool)`

GetCurrentYearToDateValueOk returns a tuple with the CurrentYearToDateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentYearToDateValue

`func (o *EsgMetricAuditableEntities) SetCurrentYearToDateValue(v float32)`

SetCurrentYearToDateValue sets CurrentYearToDateValue field to given value.

### HasCurrentYearToDateValue

`func (o *EsgMetricAuditableEntities) HasCurrentYearToDateValue() bool`

HasCurrentYearToDateValue returns a boolean if a field has been set.

### GetPreviousYearToDateValue

`func (o *EsgMetricAuditableEntities) GetPreviousYearToDateValue() float32`

GetPreviousYearToDateValue returns the PreviousYearToDateValue field if non-nil, zero value otherwise.

### GetPreviousYearToDateValueOk

`func (o *EsgMetricAuditableEntities) GetPreviousYearToDateValueOk() (*float32, bool)`

GetPreviousYearToDateValueOk returns a tuple with the PreviousYearToDateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousYearToDateValue

`func (o *EsgMetricAuditableEntities) SetPreviousYearToDateValue(v float32)`

SetPreviousYearToDateValue sets PreviousYearToDateValue field to given value.

### HasPreviousYearToDateValue

`func (o *EsgMetricAuditableEntities) HasPreviousYearToDateValue() bool`

HasPreviousYearToDateValue returns a boolean if a field has been set.

### GetQuestion

`func (o *EsgMetricAuditableEntities) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *EsgMetricAuditableEntities) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *EsgMetricAuditableEntities) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *EsgMetricAuditableEntities) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetLastCertifiedResponse

`func (o *EsgMetricAuditableEntities) GetLastCertifiedResponse() string`

GetLastCertifiedResponse returns the LastCertifiedResponse field if non-nil, zero value otherwise.

### GetLastCertifiedResponseOk

`func (o *EsgMetricAuditableEntities) GetLastCertifiedResponseOk() (*string, bool)`

GetLastCertifiedResponseOk returns a tuple with the LastCertifiedResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCertifiedResponse

`func (o *EsgMetricAuditableEntities) SetLastCertifiedResponse(v string)`

SetLastCertifiedResponse sets LastCertifiedResponse field to given value.

### HasLastCertifiedResponse

`func (o *EsgMetricAuditableEntities) HasLastCertifiedResponse() bool`

HasLastCertifiedResponse returns a boolean if a field has been set.

### GetTwoPriorYearToDateValue

`func (o *EsgMetricAuditableEntities) GetTwoPriorYearToDateValue() float32`

GetTwoPriorYearToDateValue returns the TwoPriorYearToDateValue field if non-nil, zero value otherwise.

### GetTwoPriorYearToDateValueOk

`func (o *EsgMetricAuditableEntities) GetTwoPriorYearToDateValueOk() (*float32, bool)`

GetTwoPriorYearToDateValueOk returns a tuple with the TwoPriorYearToDateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoPriorYearToDateValue

`func (o *EsgMetricAuditableEntities) SetTwoPriorYearToDateValue(v float32)`

SetTwoPriorYearToDateValue sets TwoPriorYearToDateValue field to given value.

### HasTwoPriorYearToDateValue

`func (o *EsgMetricAuditableEntities) HasTwoPriorYearToDateValue() bool`

HasTwoPriorYearToDateValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


