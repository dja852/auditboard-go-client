# AuditQuestionResponsesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] 
**AuditSurveyId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_surveys.id&#x60;.&lt;fk table&#x3D;&#39;audit_surveys&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AuditQuestionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_questions.id&#x60;.&lt;fk table&#x3D;&#39;audit_questions&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**ResponseData** | Pointer to **interface{}** |  | [optional] 
**AnsweredByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AnsweredDate** | Pointer to **string** |  | [optional] 
**IsAnswered** | Pointer to **bool** |  | [optional] [default to false]
**IsNotApplicable** | Pointer to **bool** |  | [optional] [default to false]
**IsFailedResponse** | Pointer to **bool** |  | [optional] [default to false]
**QuestionUuid** | Pointer to **string** |  | [optional] 
**OverrideData** | Pointer to **interface{}** |  | [optional] 
**OverrideReason** | Pointer to **string** |  | [optional] 
**OverriddenAt** | Pointer to **string** |  | [optional] 
**OverriddenByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IsNotApplicableOverride** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAuditQuestionResponsesPutPreviousValues

`func NewAuditQuestionResponsesPutPreviousValues() *AuditQuestionResponsesPutPreviousValues`

NewAuditQuestionResponsesPutPreviousValues instantiates a new AuditQuestionResponsesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditQuestionResponsesPutPreviousValuesWithDefaults

`func NewAuditQuestionResponsesPutPreviousValuesWithDefaults() *AuditQuestionResponsesPutPreviousValues`

NewAuditQuestionResponsesPutPreviousValuesWithDefaults instantiates a new AuditQuestionResponsesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditQuestionResponsesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditQuestionResponsesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditQuestionResponsesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuditQuestionResponsesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuditQuestionResponsesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditQuestionResponsesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditQuestionResponsesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuditQuestionResponsesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuditQuestionResponsesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuditQuestionResponsesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuditQuestionResponsesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuditQuestionResponsesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AuditQuestionResponsesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AuditQuestionResponsesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AuditQuestionResponsesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AuditQuestionResponsesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetType

`func (o *AuditQuestionResponsesPutPreviousValues) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditQuestionResponsesPutPreviousValues) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditQuestionResponsesPutPreviousValues) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuditQuestionResponsesPutPreviousValues) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAuditSurveyId

`func (o *AuditQuestionResponsesPutPreviousValues) GetAuditSurveyId() int32`

GetAuditSurveyId returns the AuditSurveyId field if non-nil, zero value otherwise.

### GetAuditSurveyIdOk

`func (o *AuditQuestionResponsesPutPreviousValues) GetAuditSurveyIdOk() (*int32, bool)`

GetAuditSurveyIdOk returns a tuple with the AuditSurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSurveyId

`func (o *AuditQuestionResponsesPutPreviousValues) SetAuditSurveyId(v int32)`

SetAuditSurveyId sets AuditSurveyId field to given value.

### HasAuditSurveyId

`func (o *AuditQuestionResponsesPutPreviousValues) HasAuditSurveyId() bool`

HasAuditSurveyId returns a boolean if a field has been set.

### GetAuditQuestionId

`func (o *AuditQuestionResponsesPutPreviousValues) GetAuditQuestionId() int32`

GetAuditQuestionId returns the AuditQuestionId field if non-nil, zero value otherwise.

### GetAuditQuestionIdOk

`func (o *AuditQuestionResponsesPutPreviousValues) GetAuditQuestionIdOk() (*int32, bool)`

GetAuditQuestionIdOk returns a tuple with the AuditQuestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditQuestionId

`func (o *AuditQuestionResponsesPutPreviousValues) SetAuditQuestionId(v int32)`

SetAuditQuestionId sets AuditQuestionId field to given value.

### HasAuditQuestionId

`func (o *AuditQuestionResponsesPutPreviousValues) HasAuditQuestionId() bool`

HasAuditQuestionId returns a boolean if a field has been set.

### GetNotes

`func (o *AuditQuestionResponsesPutPreviousValues) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AuditQuestionResponsesPutPreviousValues) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AuditQuestionResponsesPutPreviousValues) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AuditQuestionResponsesPutPreviousValues) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetResponseData

`func (o *AuditQuestionResponsesPutPreviousValues) GetResponseData() interface{}`

GetResponseData returns the ResponseData field if non-nil, zero value otherwise.

### GetResponseDataOk

`func (o *AuditQuestionResponsesPutPreviousValues) GetResponseDataOk() (*interface{}, bool)`

GetResponseDataOk returns a tuple with the ResponseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseData

`func (o *AuditQuestionResponsesPutPreviousValues) SetResponseData(v interface{})`

SetResponseData sets ResponseData field to given value.

### HasResponseData

`func (o *AuditQuestionResponsesPutPreviousValues) HasResponseData() bool`

HasResponseData returns a boolean if a field has been set.

### SetResponseDataNil

`func (o *AuditQuestionResponsesPutPreviousValues) SetResponseDataNil(b bool)`

 SetResponseDataNil sets the value for ResponseData to be an explicit nil

### UnsetResponseData
`func (o *AuditQuestionResponsesPutPreviousValues) UnsetResponseData()`

UnsetResponseData ensures that no value is present for ResponseData, not even an explicit nil
### GetAnsweredByUserId

`func (o *AuditQuestionResponsesPutPreviousValues) GetAnsweredByUserId() int32`

GetAnsweredByUserId returns the AnsweredByUserId field if non-nil, zero value otherwise.

### GetAnsweredByUserIdOk

`func (o *AuditQuestionResponsesPutPreviousValues) GetAnsweredByUserIdOk() (*int32, bool)`

GetAnsweredByUserIdOk returns a tuple with the AnsweredByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsweredByUserId

`func (o *AuditQuestionResponsesPutPreviousValues) SetAnsweredByUserId(v int32)`

SetAnsweredByUserId sets AnsweredByUserId field to given value.

### HasAnsweredByUserId

`func (o *AuditQuestionResponsesPutPreviousValues) HasAnsweredByUserId() bool`

HasAnsweredByUserId returns a boolean if a field has been set.

### GetAnsweredDate

`func (o *AuditQuestionResponsesPutPreviousValues) GetAnsweredDate() string`

GetAnsweredDate returns the AnsweredDate field if non-nil, zero value otherwise.

### GetAnsweredDateOk

`func (o *AuditQuestionResponsesPutPreviousValues) GetAnsweredDateOk() (*string, bool)`

GetAnsweredDateOk returns a tuple with the AnsweredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsweredDate

`func (o *AuditQuestionResponsesPutPreviousValues) SetAnsweredDate(v string)`

SetAnsweredDate sets AnsweredDate field to given value.

### HasAnsweredDate

`func (o *AuditQuestionResponsesPutPreviousValues) HasAnsweredDate() bool`

HasAnsweredDate returns a boolean if a field has been set.

### GetIsAnswered

`func (o *AuditQuestionResponsesPutPreviousValues) GetIsAnswered() bool`

GetIsAnswered returns the IsAnswered field if non-nil, zero value otherwise.

### GetIsAnsweredOk

`func (o *AuditQuestionResponsesPutPreviousValues) GetIsAnsweredOk() (*bool, bool)`

GetIsAnsweredOk returns a tuple with the IsAnswered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnswered

`func (o *AuditQuestionResponsesPutPreviousValues) SetIsAnswered(v bool)`

SetIsAnswered sets IsAnswered field to given value.

### HasIsAnswered

`func (o *AuditQuestionResponsesPutPreviousValues) HasIsAnswered() bool`

HasIsAnswered returns a boolean if a field has been set.

### GetIsNotApplicable

`func (o *AuditQuestionResponsesPutPreviousValues) GetIsNotApplicable() bool`

GetIsNotApplicable returns the IsNotApplicable field if non-nil, zero value otherwise.

### GetIsNotApplicableOk

`func (o *AuditQuestionResponsesPutPreviousValues) GetIsNotApplicableOk() (*bool, bool)`

GetIsNotApplicableOk returns a tuple with the IsNotApplicable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotApplicable

`func (o *AuditQuestionResponsesPutPreviousValues) SetIsNotApplicable(v bool)`

SetIsNotApplicable sets IsNotApplicable field to given value.

### HasIsNotApplicable

`func (o *AuditQuestionResponsesPutPreviousValues) HasIsNotApplicable() bool`

HasIsNotApplicable returns a boolean if a field has been set.

### GetIsFailedResponse

`func (o *AuditQuestionResponsesPutPreviousValues) GetIsFailedResponse() bool`

GetIsFailedResponse returns the IsFailedResponse field if non-nil, zero value otherwise.

### GetIsFailedResponseOk

`func (o *AuditQuestionResponsesPutPreviousValues) GetIsFailedResponseOk() (*bool, bool)`

GetIsFailedResponseOk returns a tuple with the IsFailedResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFailedResponse

`func (o *AuditQuestionResponsesPutPreviousValues) SetIsFailedResponse(v bool)`

SetIsFailedResponse sets IsFailedResponse field to given value.

### HasIsFailedResponse

`func (o *AuditQuestionResponsesPutPreviousValues) HasIsFailedResponse() bool`

HasIsFailedResponse returns a boolean if a field has been set.

### GetQuestionUuid

`func (o *AuditQuestionResponsesPutPreviousValues) GetQuestionUuid() string`

GetQuestionUuid returns the QuestionUuid field if non-nil, zero value otherwise.

### GetQuestionUuidOk

`func (o *AuditQuestionResponsesPutPreviousValues) GetQuestionUuidOk() (*string, bool)`

GetQuestionUuidOk returns a tuple with the QuestionUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionUuid

`func (o *AuditQuestionResponsesPutPreviousValues) SetQuestionUuid(v string)`

SetQuestionUuid sets QuestionUuid field to given value.

### HasQuestionUuid

`func (o *AuditQuestionResponsesPutPreviousValues) HasQuestionUuid() bool`

HasQuestionUuid returns a boolean if a field has been set.

### GetOverrideData

`func (o *AuditQuestionResponsesPutPreviousValues) GetOverrideData() interface{}`

GetOverrideData returns the OverrideData field if non-nil, zero value otherwise.

### GetOverrideDataOk

`func (o *AuditQuestionResponsesPutPreviousValues) GetOverrideDataOk() (*interface{}, bool)`

GetOverrideDataOk returns a tuple with the OverrideData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideData

`func (o *AuditQuestionResponsesPutPreviousValues) SetOverrideData(v interface{})`

SetOverrideData sets OverrideData field to given value.

### HasOverrideData

`func (o *AuditQuestionResponsesPutPreviousValues) HasOverrideData() bool`

HasOverrideData returns a boolean if a field has been set.

### SetOverrideDataNil

`func (o *AuditQuestionResponsesPutPreviousValues) SetOverrideDataNil(b bool)`

 SetOverrideDataNil sets the value for OverrideData to be an explicit nil

### UnsetOverrideData
`func (o *AuditQuestionResponsesPutPreviousValues) UnsetOverrideData()`

UnsetOverrideData ensures that no value is present for OverrideData, not even an explicit nil
### GetOverrideReason

`func (o *AuditQuestionResponsesPutPreviousValues) GetOverrideReason() string`

GetOverrideReason returns the OverrideReason field if non-nil, zero value otherwise.

### GetOverrideReasonOk

`func (o *AuditQuestionResponsesPutPreviousValues) GetOverrideReasonOk() (*string, bool)`

GetOverrideReasonOk returns a tuple with the OverrideReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideReason

`func (o *AuditQuestionResponsesPutPreviousValues) SetOverrideReason(v string)`

SetOverrideReason sets OverrideReason field to given value.

### HasOverrideReason

`func (o *AuditQuestionResponsesPutPreviousValues) HasOverrideReason() bool`

HasOverrideReason returns a boolean if a field has been set.

### GetOverriddenAt

`func (o *AuditQuestionResponsesPutPreviousValues) GetOverriddenAt() string`

GetOverriddenAt returns the OverriddenAt field if non-nil, zero value otherwise.

### GetOverriddenAtOk

`func (o *AuditQuestionResponsesPutPreviousValues) GetOverriddenAtOk() (*string, bool)`

GetOverriddenAtOk returns a tuple with the OverriddenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenAt

`func (o *AuditQuestionResponsesPutPreviousValues) SetOverriddenAt(v string)`

SetOverriddenAt sets OverriddenAt field to given value.

### HasOverriddenAt

`func (o *AuditQuestionResponsesPutPreviousValues) HasOverriddenAt() bool`

HasOverriddenAt returns a boolean if a field has been set.

### GetOverriddenByUserId

`func (o *AuditQuestionResponsesPutPreviousValues) GetOverriddenByUserId() int32`

GetOverriddenByUserId returns the OverriddenByUserId field if non-nil, zero value otherwise.

### GetOverriddenByUserIdOk

`func (o *AuditQuestionResponsesPutPreviousValues) GetOverriddenByUserIdOk() (*int32, bool)`

GetOverriddenByUserIdOk returns a tuple with the OverriddenByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenByUserId

`func (o *AuditQuestionResponsesPutPreviousValues) SetOverriddenByUserId(v int32)`

SetOverriddenByUserId sets OverriddenByUserId field to given value.

### HasOverriddenByUserId

`func (o *AuditQuestionResponsesPutPreviousValues) HasOverriddenByUserId() bool`

HasOverriddenByUserId returns a boolean if a field has been set.

### GetIsNotApplicableOverride

`func (o *AuditQuestionResponsesPutPreviousValues) GetIsNotApplicableOverride() bool`

GetIsNotApplicableOverride returns the IsNotApplicableOverride field if non-nil, zero value otherwise.

### GetIsNotApplicableOverrideOk

`func (o *AuditQuestionResponsesPutPreviousValues) GetIsNotApplicableOverrideOk() (*bool, bool)`

GetIsNotApplicableOverrideOk returns a tuple with the IsNotApplicableOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotApplicableOverride

`func (o *AuditQuestionResponsesPutPreviousValues) SetIsNotApplicableOverride(v bool)`

SetIsNotApplicableOverride sets IsNotApplicableOverride field to given value.

### HasIsNotApplicableOverride

`func (o *AuditQuestionResponsesPutPreviousValues) HasIsNotApplicableOverride() bool`

HasIsNotApplicableOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


