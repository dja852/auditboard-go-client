# AuditQuestionsPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Requirements** | Pointer to **interface{}** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**QuestionData** | Pointer to **interface{}** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**SourceId** | Pointer to **int32** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ConditionChoiceUuid** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**ChangeableType** | Pointer to **string** |  | [optional] 
**ManagedChangeField** | Pointer to **string** |  | [optional] 
**ManagedChangeFieldType** | Pointer to **string** |  | [optional] 
**QuestionOptions** | Pointer to **interface{}** |  | [optional] 
**AllowNaResponse** | Pointer to **bool** |  | [optional] [default to false]
**AllowExplanation** | Pointer to **bool** |  | [optional] [default to false]
**AllowFiles** | Pointer to **bool** |  | [optional] [default to false]
**RequireFiles** | Pointer to **bool** |  | [optional] [default to false]
**QuestionUuid** | Pointer to **string** |  | [optional] 
**ConditionQuestionUuid** | Pointer to **string** |  | [optional] 
**Condition** | Pointer to **interface{}** |  | [optional] 
**RequireNaExplanation** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAuditQuestionsPutPreviousValues

`func NewAuditQuestionsPutPreviousValues() *AuditQuestionsPutPreviousValues`

NewAuditQuestionsPutPreviousValues instantiates a new AuditQuestionsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditQuestionsPutPreviousValuesWithDefaults

`func NewAuditQuestionsPutPreviousValuesWithDefaults() *AuditQuestionsPutPreviousValues`

NewAuditQuestionsPutPreviousValuesWithDefaults instantiates a new AuditQuestionsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditQuestionsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditQuestionsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditQuestionsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuditQuestionsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuditQuestionsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditQuestionsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditQuestionsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuditQuestionsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuditQuestionsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuditQuestionsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuditQuestionsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuditQuestionsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AuditQuestionsPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AuditQuestionsPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AuditQuestionsPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AuditQuestionsPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetRequirements

`func (o *AuditQuestionsPutPreviousValues) GetRequirements() interface{}`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *AuditQuestionsPutPreviousValues) GetRequirementsOk() (*interface{}, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *AuditQuestionsPutPreviousValues) SetRequirements(v interface{})`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *AuditQuestionsPutPreviousValues) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### SetRequirementsNil

`func (o *AuditQuestionsPutPreviousValues) SetRequirementsNil(b bool)`

 SetRequirementsNil sets the value for Requirements to be an explicit nil

### UnsetRequirements
`func (o *AuditQuestionsPutPreviousValues) UnsetRequirements()`

UnsetRequirements ensures that no value is present for Requirements, not even an explicit nil
### GetType

`func (o *AuditQuestionsPutPreviousValues) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditQuestionsPutPreviousValues) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditQuestionsPutPreviousValues) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuditQuestionsPutPreviousValues) HasType() bool`

HasType returns a boolean if a field has been set.

### GetQuestionData

`func (o *AuditQuestionsPutPreviousValues) GetQuestionData() interface{}`

GetQuestionData returns the QuestionData field if non-nil, zero value otherwise.

### GetQuestionDataOk

`func (o *AuditQuestionsPutPreviousValues) GetQuestionDataOk() (*interface{}, bool)`

GetQuestionDataOk returns a tuple with the QuestionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionData

`func (o *AuditQuestionsPutPreviousValues) SetQuestionData(v interface{})`

SetQuestionData sets QuestionData field to given value.

### HasQuestionData

`func (o *AuditQuestionsPutPreviousValues) HasQuestionData() bool`

HasQuestionData returns a boolean if a field has been set.

### SetQuestionDataNil

`func (o *AuditQuestionsPutPreviousValues) SetQuestionDataNil(b bool)`

 SetQuestionDataNil sets the value for QuestionData to be an explicit nil

### UnsetQuestionData
`func (o *AuditQuestionsPutPreviousValues) UnsetQuestionData()`

UnsetQuestionData ensures that no value is present for QuestionData, not even an explicit nil
### GetText

`func (o *AuditQuestionsPutPreviousValues) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AuditQuestionsPutPreviousValues) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AuditQuestionsPutPreviousValues) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *AuditQuestionsPutPreviousValues) HasText() bool`

HasText returns a boolean if a field has been set.

### GetDescription

`func (o *AuditQuestionsPutPreviousValues) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuditQuestionsPutPreviousValues) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuditQuestionsPutPreviousValues) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuditQuestionsPutPreviousValues) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSourceId

`func (o *AuditQuestionsPutPreviousValues) GetSourceId() int32`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AuditQuestionsPutPreviousValues) GetSourceIdOk() (*int32, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AuditQuestionsPutPreviousValues) SetSourceId(v int32)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AuditQuestionsPutPreviousValues) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *AuditQuestionsPutPreviousValues) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *AuditQuestionsPutPreviousValues) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *AuditQuestionsPutPreviousValues) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *AuditQuestionsPutPreviousValues) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetConditionChoiceUuid

`func (o *AuditQuestionsPutPreviousValues) GetConditionChoiceUuid() string`

GetConditionChoiceUuid returns the ConditionChoiceUuid field if non-nil, zero value otherwise.

### GetConditionChoiceUuidOk

`func (o *AuditQuestionsPutPreviousValues) GetConditionChoiceUuidOk() (*string, bool)`

GetConditionChoiceUuidOk returns a tuple with the ConditionChoiceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionChoiceUuid

`func (o *AuditQuestionsPutPreviousValues) SetConditionChoiceUuid(v string)`

SetConditionChoiceUuid sets ConditionChoiceUuid field to given value.

### HasConditionChoiceUuid

`func (o *AuditQuestionsPutPreviousValues) HasConditionChoiceUuid() bool`

HasConditionChoiceUuid returns a boolean if a field has been set.

### GetUid

`func (o *AuditQuestionsPutPreviousValues) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AuditQuestionsPutPreviousValues) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AuditQuestionsPutPreviousValues) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AuditQuestionsPutPreviousValues) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetChangeableType

`func (o *AuditQuestionsPutPreviousValues) GetChangeableType() string`

GetChangeableType returns the ChangeableType field if non-nil, zero value otherwise.

### GetChangeableTypeOk

`func (o *AuditQuestionsPutPreviousValues) GetChangeableTypeOk() (*string, bool)`

GetChangeableTypeOk returns a tuple with the ChangeableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeableType

`func (o *AuditQuestionsPutPreviousValues) SetChangeableType(v string)`

SetChangeableType sets ChangeableType field to given value.

### HasChangeableType

`func (o *AuditQuestionsPutPreviousValues) HasChangeableType() bool`

HasChangeableType returns a boolean if a field has been set.

### GetManagedChangeField

`func (o *AuditQuestionsPutPreviousValues) GetManagedChangeField() string`

GetManagedChangeField returns the ManagedChangeField field if non-nil, zero value otherwise.

### GetManagedChangeFieldOk

`func (o *AuditQuestionsPutPreviousValues) GetManagedChangeFieldOk() (*string, bool)`

GetManagedChangeFieldOk returns a tuple with the ManagedChangeField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedChangeField

`func (o *AuditQuestionsPutPreviousValues) SetManagedChangeField(v string)`

SetManagedChangeField sets ManagedChangeField field to given value.

### HasManagedChangeField

`func (o *AuditQuestionsPutPreviousValues) HasManagedChangeField() bool`

HasManagedChangeField returns a boolean if a field has been set.

### GetManagedChangeFieldType

`func (o *AuditQuestionsPutPreviousValues) GetManagedChangeFieldType() string`

GetManagedChangeFieldType returns the ManagedChangeFieldType field if non-nil, zero value otherwise.

### GetManagedChangeFieldTypeOk

`func (o *AuditQuestionsPutPreviousValues) GetManagedChangeFieldTypeOk() (*string, bool)`

GetManagedChangeFieldTypeOk returns a tuple with the ManagedChangeFieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedChangeFieldType

`func (o *AuditQuestionsPutPreviousValues) SetManagedChangeFieldType(v string)`

SetManagedChangeFieldType sets ManagedChangeFieldType field to given value.

### HasManagedChangeFieldType

`func (o *AuditQuestionsPutPreviousValues) HasManagedChangeFieldType() bool`

HasManagedChangeFieldType returns a boolean if a field has been set.

### GetQuestionOptions

`func (o *AuditQuestionsPutPreviousValues) GetQuestionOptions() interface{}`

GetQuestionOptions returns the QuestionOptions field if non-nil, zero value otherwise.

### GetQuestionOptionsOk

`func (o *AuditQuestionsPutPreviousValues) GetQuestionOptionsOk() (*interface{}, bool)`

GetQuestionOptionsOk returns a tuple with the QuestionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionOptions

`func (o *AuditQuestionsPutPreviousValues) SetQuestionOptions(v interface{})`

SetQuestionOptions sets QuestionOptions field to given value.

### HasQuestionOptions

`func (o *AuditQuestionsPutPreviousValues) HasQuestionOptions() bool`

HasQuestionOptions returns a boolean if a field has been set.

### SetQuestionOptionsNil

`func (o *AuditQuestionsPutPreviousValues) SetQuestionOptionsNil(b bool)`

 SetQuestionOptionsNil sets the value for QuestionOptions to be an explicit nil

### UnsetQuestionOptions
`func (o *AuditQuestionsPutPreviousValues) UnsetQuestionOptions()`

UnsetQuestionOptions ensures that no value is present for QuestionOptions, not even an explicit nil
### GetAllowNaResponse

`func (o *AuditQuestionsPutPreviousValues) GetAllowNaResponse() bool`

GetAllowNaResponse returns the AllowNaResponse field if non-nil, zero value otherwise.

### GetAllowNaResponseOk

`func (o *AuditQuestionsPutPreviousValues) GetAllowNaResponseOk() (*bool, bool)`

GetAllowNaResponseOk returns a tuple with the AllowNaResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNaResponse

`func (o *AuditQuestionsPutPreviousValues) SetAllowNaResponse(v bool)`

SetAllowNaResponse sets AllowNaResponse field to given value.

### HasAllowNaResponse

`func (o *AuditQuestionsPutPreviousValues) HasAllowNaResponse() bool`

HasAllowNaResponse returns a boolean if a field has been set.

### GetAllowExplanation

`func (o *AuditQuestionsPutPreviousValues) GetAllowExplanation() bool`

GetAllowExplanation returns the AllowExplanation field if non-nil, zero value otherwise.

### GetAllowExplanationOk

`func (o *AuditQuestionsPutPreviousValues) GetAllowExplanationOk() (*bool, bool)`

GetAllowExplanationOk returns a tuple with the AllowExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExplanation

`func (o *AuditQuestionsPutPreviousValues) SetAllowExplanation(v bool)`

SetAllowExplanation sets AllowExplanation field to given value.

### HasAllowExplanation

`func (o *AuditQuestionsPutPreviousValues) HasAllowExplanation() bool`

HasAllowExplanation returns a boolean if a field has been set.

### GetAllowFiles

`func (o *AuditQuestionsPutPreviousValues) GetAllowFiles() bool`

GetAllowFiles returns the AllowFiles field if non-nil, zero value otherwise.

### GetAllowFilesOk

`func (o *AuditQuestionsPutPreviousValues) GetAllowFilesOk() (*bool, bool)`

GetAllowFilesOk returns a tuple with the AllowFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFiles

`func (o *AuditQuestionsPutPreviousValues) SetAllowFiles(v bool)`

SetAllowFiles sets AllowFiles field to given value.

### HasAllowFiles

`func (o *AuditQuestionsPutPreviousValues) HasAllowFiles() bool`

HasAllowFiles returns a boolean if a field has been set.

### GetRequireFiles

`func (o *AuditQuestionsPutPreviousValues) GetRequireFiles() bool`

GetRequireFiles returns the RequireFiles field if non-nil, zero value otherwise.

### GetRequireFilesOk

`func (o *AuditQuestionsPutPreviousValues) GetRequireFilesOk() (*bool, bool)`

GetRequireFilesOk returns a tuple with the RequireFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireFiles

`func (o *AuditQuestionsPutPreviousValues) SetRequireFiles(v bool)`

SetRequireFiles sets RequireFiles field to given value.

### HasRequireFiles

`func (o *AuditQuestionsPutPreviousValues) HasRequireFiles() bool`

HasRequireFiles returns a boolean if a field has been set.

### GetQuestionUuid

`func (o *AuditQuestionsPutPreviousValues) GetQuestionUuid() string`

GetQuestionUuid returns the QuestionUuid field if non-nil, zero value otherwise.

### GetQuestionUuidOk

`func (o *AuditQuestionsPutPreviousValues) GetQuestionUuidOk() (*string, bool)`

GetQuestionUuidOk returns a tuple with the QuestionUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionUuid

`func (o *AuditQuestionsPutPreviousValues) SetQuestionUuid(v string)`

SetQuestionUuid sets QuestionUuid field to given value.

### HasQuestionUuid

`func (o *AuditQuestionsPutPreviousValues) HasQuestionUuid() bool`

HasQuestionUuid returns a boolean if a field has been set.

### GetConditionQuestionUuid

`func (o *AuditQuestionsPutPreviousValues) GetConditionQuestionUuid() string`

GetConditionQuestionUuid returns the ConditionQuestionUuid field if non-nil, zero value otherwise.

### GetConditionQuestionUuidOk

`func (o *AuditQuestionsPutPreviousValues) GetConditionQuestionUuidOk() (*string, bool)`

GetConditionQuestionUuidOk returns a tuple with the ConditionQuestionUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionQuestionUuid

`func (o *AuditQuestionsPutPreviousValues) SetConditionQuestionUuid(v string)`

SetConditionQuestionUuid sets ConditionQuestionUuid field to given value.

### HasConditionQuestionUuid

`func (o *AuditQuestionsPutPreviousValues) HasConditionQuestionUuid() bool`

HasConditionQuestionUuid returns a boolean if a field has been set.

### GetCondition

`func (o *AuditQuestionsPutPreviousValues) GetCondition() interface{}`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AuditQuestionsPutPreviousValues) GetConditionOk() (*interface{}, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AuditQuestionsPutPreviousValues) SetCondition(v interface{})`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *AuditQuestionsPutPreviousValues) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *AuditQuestionsPutPreviousValues) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *AuditQuestionsPutPreviousValues) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetRequireNaExplanation

`func (o *AuditQuestionsPutPreviousValues) GetRequireNaExplanation() bool`

GetRequireNaExplanation returns the RequireNaExplanation field if non-nil, zero value otherwise.

### GetRequireNaExplanationOk

`func (o *AuditQuestionsPutPreviousValues) GetRequireNaExplanationOk() (*bool, bool)`

GetRequireNaExplanationOk returns a tuple with the RequireNaExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireNaExplanation

`func (o *AuditQuestionsPutPreviousValues) SetRequireNaExplanation(v bool)`

SetRequireNaExplanation sets RequireNaExplanation field to given value.

### HasRequireNaExplanation

`func (o *AuditQuestionsPutPreviousValues) HasRequireNaExplanation() bool`

HasRequireNaExplanation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


