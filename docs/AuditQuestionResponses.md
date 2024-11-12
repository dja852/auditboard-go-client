# AuditQuestionResponses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Type** | **string** |  | 
**AuditSurveyId** | **int32** | Note: This is a Foreign Key to &#x60;audit_surveys.id&#x60;.&lt;fk table&#x3D;&#39;audit_surveys&#39; column&#x3D;&#39;id&#39;/&gt; | 
**AuditQuestionId** | **int32** | Note: This is a Foreign Key to &#x60;audit_questions.id&#x60;.&lt;fk table&#x3D;&#39;audit_questions&#39; column&#x3D;&#39;id&#39;/&gt; | 
**Notes** | Pointer to **string** |  | [optional] 
**ResponseData** | Pointer to **interface{}** |  | [optional] 
**AnsweredByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AnsweredDate** | Pointer to **string** |  | [optional] 
**IsAnswered** | **bool** |  | [default to false]
**IsNotApplicable** | **bool** |  | [default to false]
**IsFailedResponse** | **bool** |  | [default to false]
**QuestionUuid** | **string** |  | 
**OverrideData** | Pointer to **interface{}** |  | [optional] 
**OverrideReason** | Pointer to **string** |  | [optional] 
**OverriddenAt** | Pointer to **string** |  | [optional] 
**OverriddenByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IsNotApplicableOverride** | **bool** |  | [default to false]

## Methods

### NewAuditQuestionResponses

`func NewAuditQuestionResponses(type_ string, auditSurveyId int32, auditQuestionId int32, isAnswered bool, isNotApplicable bool, isFailedResponse bool, questionUuid string, isNotApplicableOverride bool, ) *AuditQuestionResponses`

NewAuditQuestionResponses instantiates a new AuditQuestionResponses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditQuestionResponsesWithDefaults

`func NewAuditQuestionResponsesWithDefaults() *AuditQuestionResponses`

NewAuditQuestionResponsesWithDefaults instantiates a new AuditQuestionResponses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditQuestionResponses) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditQuestionResponses) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditQuestionResponses) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuditQuestionResponses) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuditQuestionResponses) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditQuestionResponses) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditQuestionResponses) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuditQuestionResponses) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuditQuestionResponses) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuditQuestionResponses) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuditQuestionResponses) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuditQuestionResponses) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AuditQuestionResponses) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AuditQuestionResponses) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AuditQuestionResponses) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AuditQuestionResponses) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetType

`func (o *AuditQuestionResponses) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditQuestionResponses) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditQuestionResponses) SetType(v string)`

SetType sets Type field to given value.


### GetAuditSurveyId

`func (o *AuditQuestionResponses) GetAuditSurveyId() int32`

GetAuditSurveyId returns the AuditSurveyId field if non-nil, zero value otherwise.

### GetAuditSurveyIdOk

`func (o *AuditQuestionResponses) GetAuditSurveyIdOk() (*int32, bool)`

GetAuditSurveyIdOk returns a tuple with the AuditSurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSurveyId

`func (o *AuditQuestionResponses) SetAuditSurveyId(v int32)`

SetAuditSurveyId sets AuditSurveyId field to given value.


### GetAuditQuestionId

`func (o *AuditQuestionResponses) GetAuditQuestionId() int32`

GetAuditQuestionId returns the AuditQuestionId field if non-nil, zero value otherwise.

### GetAuditQuestionIdOk

`func (o *AuditQuestionResponses) GetAuditQuestionIdOk() (*int32, bool)`

GetAuditQuestionIdOk returns a tuple with the AuditQuestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditQuestionId

`func (o *AuditQuestionResponses) SetAuditQuestionId(v int32)`

SetAuditQuestionId sets AuditQuestionId field to given value.


### GetNotes

`func (o *AuditQuestionResponses) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AuditQuestionResponses) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AuditQuestionResponses) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AuditQuestionResponses) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetResponseData

`func (o *AuditQuestionResponses) GetResponseData() interface{}`

GetResponseData returns the ResponseData field if non-nil, zero value otherwise.

### GetResponseDataOk

`func (o *AuditQuestionResponses) GetResponseDataOk() (*interface{}, bool)`

GetResponseDataOk returns a tuple with the ResponseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseData

`func (o *AuditQuestionResponses) SetResponseData(v interface{})`

SetResponseData sets ResponseData field to given value.

### HasResponseData

`func (o *AuditQuestionResponses) HasResponseData() bool`

HasResponseData returns a boolean if a field has been set.

### SetResponseDataNil

`func (o *AuditQuestionResponses) SetResponseDataNil(b bool)`

 SetResponseDataNil sets the value for ResponseData to be an explicit nil

### UnsetResponseData
`func (o *AuditQuestionResponses) UnsetResponseData()`

UnsetResponseData ensures that no value is present for ResponseData, not even an explicit nil
### GetAnsweredByUserId

`func (o *AuditQuestionResponses) GetAnsweredByUserId() int32`

GetAnsweredByUserId returns the AnsweredByUserId field if non-nil, zero value otherwise.

### GetAnsweredByUserIdOk

`func (o *AuditQuestionResponses) GetAnsweredByUserIdOk() (*int32, bool)`

GetAnsweredByUserIdOk returns a tuple with the AnsweredByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsweredByUserId

`func (o *AuditQuestionResponses) SetAnsweredByUserId(v int32)`

SetAnsweredByUserId sets AnsweredByUserId field to given value.

### HasAnsweredByUserId

`func (o *AuditQuestionResponses) HasAnsweredByUserId() bool`

HasAnsweredByUserId returns a boolean if a field has been set.

### GetAnsweredDate

`func (o *AuditQuestionResponses) GetAnsweredDate() string`

GetAnsweredDate returns the AnsweredDate field if non-nil, zero value otherwise.

### GetAnsweredDateOk

`func (o *AuditQuestionResponses) GetAnsweredDateOk() (*string, bool)`

GetAnsweredDateOk returns a tuple with the AnsweredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsweredDate

`func (o *AuditQuestionResponses) SetAnsweredDate(v string)`

SetAnsweredDate sets AnsweredDate field to given value.

### HasAnsweredDate

`func (o *AuditQuestionResponses) HasAnsweredDate() bool`

HasAnsweredDate returns a boolean if a field has been set.

### GetIsAnswered

`func (o *AuditQuestionResponses) GetIsAnswered() bool`

GetIsAnswered returns the IsAnswered field if non-nil, zero value otherwise.

### GetIsAnsweredOk

`func (o *AuditQuestionResponses) GetIsAnsweredOk() (*bool, bool)`

GetIsAnsweredOk returns a tuple with the IsAnswered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnswered

`func (o *AuditQuestionResponses) SetIsAnswered(v bool)`

SetIsAnswered sets IsAnswered field to given value.


### GetIsNotApplicable

`func (o *AuditQuestionResponses) GetIsNotApplicable() bool`

GetIsNotApplicable returns the IsNotApplicable field if non-nil, zero value otherwise.

### GetIsNotApplicableOk

`func (o *AuditQuestionResponses) GetIsNotApplicableOk() (*bool, bool)`

GetIsNotApplicableOk returns a tuple with the IsNotApplicable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotApplicable

`func (o *AuditQuestionResponses) SetIsNotApplicable(v bool)`

SetIsNotApplicable sets IsNotApplicable field to given value.


### GetIsFailedResponse

`func (o *AuditQuestionResponses) GetIsFailedResponse() bool`

GetIsFailedResponse returns the IsFailedResponse field if non-nil, zero value otherwise.

### GetIsFailedResponseOk

`func (o *AuditQuestionResponses) GetIsFailedResponseOk() (*bool, bool)`

GetIsFailedResponseOk returns a tuple with the IsFailedResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFailedResponse

`func (o *AuditQuestionResponses) SetIsFailedResponse(v bool)`

SetIsFailedResponse sets IsFailedResponse field to given value.


### GetQuestionUuid

`func (o *AuditQuestionResponses) GetQuestionUuid() string`

GetQuestionUuid returns the QuestionUuid field if non-nil, zero value otherwise.

### GetQuestionUuidOk

`func (o *AuditQuestionResponses) GetQuestionUuidOk() (*string, bool)`

GetQuestionUuidOk returns a tuple with the QuestionUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionUuid

`func (o *AuditQuestionResponses) SetQuestionUuid(v string)`

SetQuestionUuid sets QuestionUuid field to given value.


### GetOverrideData

`func (o *AuditQuestionResponses) GetOverrideData() interface{}`

GetOverrideData returns the OverrideData field if non-nil, zero value otherwise.

### GetOverrideDataOk

`func (o *AuditQuestionResponses) GetOverrideDataOk() (*interface{}, bool)`

GetOverrideDataOk returns a tuple with the OverrideData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideData

`func (o *AuditQuestionResponses) SetOverrideData(v interface{})`

SetOverrideData sets OverrideData field to given value.

### HasOverrideData

`func (o *AuditQuestionResponses) HasOverrideData() bool`

HasOverrideData returns a boolean if a field has been set.

### SetOverrideDataNil

`func (o *AuditQuestionResponses) SetOverrideDataNil(b bool)`

 SetOverrideDataNil sets the value for OverrideData to be an explicit nil

### UnsetOverrideData
`func (o *AuditQuestionResponses) UnsetOverrideData()`

UnsetOverrideData ensures that no value is present for OverrideData, not even an explicit nil
### GetOverrideReason

`func (o *AuditQuestionResponses) GetOverrideReason() string`

GetOverrideReason returns the OverrideReason field if non-nil, zero value otherwise.

### GetOverrideReasonOk

`func (o *AuditQuestionResponses) GetOverrideReasonOk() (*string, bool)`

GetOverrideReasonOk returns a tuple with the OverrideReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideReason

`func (o *AuditQuestionResponses) SetOverrideReason(v string)`

SetOverrideReason sets OverrideReason field to given value.

### HasOverrideReason

`func (o *AuditQuestionResponses) HasOverrideReason() bool`

HasOverrideReason returns a boolean if a field has been set.

### GetOverriddenAt

`func (o *AuditQuestionResponses) GetOverriddenAt() string`

GetOverriddenAt returns the OverriddenAt field if non-nil, zero value otherwise.

### GetOverriddenAtOk

`func (o *AuditQuestionResponses) GetOverriddenAtOk() (*string, bool)`

GetOverriddenAtOk returns a tuple with the OverriddenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenAt

`func (o *AuditQuestionResponses) SetOverriddenAt(v string)`

SetOverriddenAt sets OverriddenAt field to given value.

### HasOverriddenAt

`func (o *AuditQuestionResponses) HasOverriddenAt() bool`

HasOverriddenAt returns a boolean if a field has been set.

### GetOverriddenByUserId

`func (o *AuditQuestionResponses) GetOverriddenByUserId() int32`

GetOverriddenByUserId returns the OverriddenByUserId field if non-nil, zero value otherwise.

### GetOverriddenByUserIdOk

`func (o *AuditQuestionResponses) GetOverriddenByUserIdOk() (*int32, bool)`

GetOverriddenByUserIdOk returns a tuple with the OverriddenByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenByUserId

`func (o *AuditQuestionResponses) SetOverriddenByUserId(v int32)`

SetOverriddenByUserId sets OverriddenByUserId field to given value.

### HasOverriddenByUserId

`func (o *AuditQuestionResponses) HasOverriddenByUserId() bool`

HasOverriddenByUserId returns a boolean if a field has been set.

### GetIsNotApplicableOverride

`func (o *AuditQuestionResponses) GetIsNotApplicableOverride() bool`

GetIsNotApplicableOverride returns the IsNotApplicableOverride field if non-nil, zero value otherwise.

### GetIsNotApplicableOverrideOk

`func (o *AuditQuestionResponses) GetIsNotApplicableOverrideOk() (*bool, bool)`

GetIsNotApplicableOverrideOk returns a tuple with the IsNotApplicableOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotApplicableOverride

`func (o *AuditQuestionResponses) SetIsNotApplicableOverride(v bool)`

SetIsNotApplicableOverride sets IsNotApplicableOverride field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


