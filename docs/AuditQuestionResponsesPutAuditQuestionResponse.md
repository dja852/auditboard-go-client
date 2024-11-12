# AuditQuestionResponsesPutAuditQuestionResponse

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

### NewAuditQuestionResponsesPutAuditQuestionResponse

`func NewAuditQuestionResponsesPutAuditQuestionResponse(type_ string, auditSurveyId int32, auditQuestionId int32, isAnswered bool, isNotApplicable bool, isFailedResponse bool, questionUuid string, isNotApplicableOverride bool, ) *AuditQuestionResponsesPutAuditQuestionResponse`

NewAuditQuestionResponsesPutAuditQuestionResponse instantiates a new AuditQuestionResponsesPutAuditQuestionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditQuestionResponsesPutAuditQuestionResponseWithDefaults

`func NewAuditQuestionResponsesPutAuditQuestionResponseWithDefaults() *AuditQuestionResponsesPutAuditQuestionResponse`

NewAuditQuestionResponsesPutAuditQuestionResponseWithDefaults instantiates a new AuditQuestionResponsesPutAuditQuestionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetType

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetType(v string)`

SetType sets Type field to given value.


### GetAuditSurveyId

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetAuditSurveyId() int32`

GetAuditSurveyId returns the AuditSurveyId field if non-nil, zero value otherwise.

### GetAuditSurveyIdOk

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetAuditSurveyIdOk() (*int32, bool)`

GetAuditSurveyIdOk returns a tuple with the AuditSurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSurveyId

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetAuditSurveyId(v int32)`

SetAuditSurveyId sets AuditSurveyId field to given value.


### GetAuditQuestionId

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetAuditQuestionId() int32`

GetAuditQuestionId returns the AuditQuestionId field if non-nil, zero value otherwise.

### GetAuditQuestionIdOk

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetAuditQuestionIdOk() (*int32, bool)`

GetAuditQuestionIdOk returns a tuple with the AuditQuestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditQuestionId

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetAuditQuestionId(v int32)`

SetAuditQuestionId sets AuditQuestionId field to given value.


### GetNotes

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetResponseData

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetResponseData() interface{}`

GetResponseData returns the ResponseData field if non-nil, zero value otherwise.

### GetResponseDataOk

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetResponseDataOk() (*interface{}, bool)`

GetResponseDataOk returns a tuple with the ResponseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseData

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetResponseData(v interface{})`

SetResponseData sets ResponseData field to given value.

### HasResponseData

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) HasResponseData() bool`

HasResponseData returns a boolean if a field has been set.

### SetResponseDataNil

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetResponseDataNil(b bool)`

 SetResponseDataNil sets the value for ResponseData to be an explicit nil

### UnsetResponseData
`func (o *AuditQuestionResponsesPutAuditQuestionResponse) UnsetResponseData()`

UnsetResponseData ensures that no value is present for ResponseData, not even an explicit nil
### GetAnsweredByUserId

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetAnsweredByUserId() int32`

GetAnsweredByUserId returns the AnsweredByUserId field if non-nil, zero value otherwise.

### GetAnsweredByUserIdOk

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetAnsweredByUserIdOk() (*int32, bool)`

GetAnsweredByUserIdOk returns a tuple with the AnsweredByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsweredByUserId

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetAnsweredByUserId(v int32)`

SetAnsweredByUserId sets AnsweredByUserId field to given value.

### HasAnsweredByUserId

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) HasAnsweredByUserId() bool`

HasAnsweredByUserId returns a boolean if a field has been set.

### GetAnsweredDate

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetAnsweredDate() string`

GetAnsweredDate returns the AnsweredDate field if non-nil, zero value otherwise.

### GetAnsweredDateOk

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetAnsweredDateOk() (*string, bool)`

GetAnsweredDateOk returns a tuple with the AnsweredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsweredDate

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetAnsweredDate(v string)`

SetAnsweredDate sets AnsweredDate field to given value.

### HasAnsweredDate

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) HasAnsweredDate() bool`

HasAnsweredDate returns a boolean if a field has been set.

### GetIsAnswered

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetIsAnswered() bool`

GetIsAnswered returns the IsAnswered field if non-nil, zero value otherwise.

### GetIsAnsweredOk

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetIsAnsweredOk() (*bool, bool)`

GetIsAnsweredOk returns a tuple with the IsAnswered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnswered

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetIsAnswered(v bool)`

SetIsAnswered sets IsAnswered field to given value.


### GetIsNotApplicable

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetIsNotApplicable() bool`

GetIsNotApplicable returns the IsNotApplicable field if non-nil, zero value otherwise.

### GetIsNotApplicableOk

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetIsNotApplicableOk() (*bool, bool)`

GetIsNotApplicableOk returns a tuple with the IsNotApplicable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotApplicable

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetIsNotApplicable(v bool)`

SetIsNotApplicable sets IsNotApplicable field to given value.


### GetIsFailedResponse

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetIsFailedResponse() bool`

GetIsFailedResponse returns the IsFailedResponse field if non-nil, zero value otherwise.

### GetIsFailedResponseOk

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetIsFailedResponseOk() (*bool, bool)`

GetIsFailedResponseOk returns a tuple with the IsFailedResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFailedResponse

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetIsFailedResponse(v bool)`

SetIsFailedResponse sets IsFailedResponse field to given value.


### GetQuestionUuid

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetQuestionUuid() string`

GetQuestionUuid returns the QuestionUuid field if non-nil, zero value otherwise.

### GetQuestionUuidOk

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetQuestionUuidOk() (*string, bool)`

GetQuestionUuidOk returns a tuple with the QuestionUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionUuid

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetQuestionUuid(v string)`

SetQuestionUuid sets QuestionUuid field to given value.


### GetOverrideData

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetOverrideData() interface{}`

GetOverrideData returns the OverrideData field if non-nil, zero value otherwise.

### GetOverrideDataOk

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetOverrideDataOk() (*interface{}, bool)`

GetOverrideDataOk returns a tuple with the OverrideData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideData

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetOverrideData(v interface{})`

SetOverrideData sets OverrideData field to given value.

### HasOverrideData

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) HasOverrideData() bool`

HasOverrideData returns a boolean if a field has been set.

### SetOverrideDataNil

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetOverrideDataNil(b bool)`

 SetOverrideDataNil sets the value for OverrideData to be an explicit nil

### UnsetOverrideData
`func (o *AuditQuestionResponsesPutAuditQuestionResponse) UnsetOverrideData()`

UnsetOverrideData ensures that no value is present for OverrideData, not even an explicit nil
### GetOverrideReason

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetOverrideReason() string`

GetOverrideReason returns the OverrideReason field if non-nil, zero value otherwise.

### GetOverrideReasonOk

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetOverrideReasonOk() (*string, bool)`

GetOverrideReasonOk returns a tuple with the OverrideReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideReason

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetOverrideReason(v string)`

SetOverrideReason sets OverrideReason field to given value.

### HasOverrideReason

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) HasOverrideReason() bool`

HasOverrideReason returns a boolean if a field has been set.

### GetOverriddenAt

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetOverriddenAt() string`

GetOverriddenAt returns the OverriddenAt field if non-nil, zero value otherwise.

### GetOverriddenAtOk

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetOverriddenAtOk() (*string, bool)`

GetOverriddenAtOk returns a tuple with the OverriddenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenAt

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetOverriddenAt(v string)`

SetOverriddenAt sets OverriddenAt field to given value.

### HasOverriddenAt

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) HasOverriddenAt() bool`

HasOverriddenAt returns a boolean if a field has been set.

### GetOverriddenByUserId

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetOverriddenByUserId() int32`

GetOverriddenByUserId returns the OverriddenByUserId field if non-nil, zero value otherwise.

### GetOverriddenByUserIdOk

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetOverriddenByUserIdOk() (*int32, bool)`

GetOverriddenByUserIdOk returns a tuple with the OverriddenByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenByUserId

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetOverriddenByUserId(v int32)`

SetOverriddenByUserId sets OverriddenByUserId field to given value.

### HasOverriddenByUserId

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) HasOverriddenByUserId() bool`

HasOverriddenByUserId returns a boolean if a field has been set.

### GetIsNotApplicableOverride

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetIsNotApplicableOverride() bool`

GetIsNotApplicableOverride returns the IsNotApplicableOverride field if non-nil, zero value otherwise.

### GetIsNotApplicableOverrideOk

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) GetIsNotApplicableOverrideOk() (*bool, bool)`

GetIsNotApplicableOverrideOk returns a tuple with the IsNotApplicableOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotApplicableOverride

`func (o *AuditQuestionResponsesPutAuditQuestionResponse) SetIsNotApplicableOverride(v bool)`

SetIsNotApplicableOverride sets IsNotApplicableOverride field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


