# AuditQuestionsPutAuditQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Requirements** | Pointer to **interface{}** |  | [optional] 
**Type** | **string** |  | 
**QuestionData** | **interface{}** |  | 
**Text** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**SourceId** | Pointer to **int32** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ConditionChoiceUuid** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**ChangeableType** | Pointer to **string** |  | [optional] 
**ManagedChangeField** | Pointer to **string** |  | [optional] 
**ManagedChangeFieldType** | Pointer to **string** |  | [optional] 
**QuestionOptions** | **interface{}** |  | 
**AllowNaResponse** | **bool** |  | [default to false]
**AllowExplanation** | **bool** |  | [default to false]
**AllowFiles** | **bool** |  | [default to false]
**RequireFiles** | **bool** |  | [default to false]
**QuestionUuid** | Pointer to **string** |  | [optional] 
**ConditionQuestionUuid** | Pointer to **string** |  | [optional] 
**Condition** | Pointer to **interface{}** |  | [optional] 
**RequireNaExplanation** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAuditQuestionsPutAuditQuestion

`func NewAuditQuestionsPutAuditQuestion(type_ string, questionData interface{}, text string, questionOptions interface{}, allowNaResponse bool, allowExplanation bool, allowFiles bool, requireFiles bool, ) *AuditQuestionsPutAuditQuestion`

NewAuditQuestionsPutAuditQuestion instantiates a new AuditQuestionsPutAuditQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditQuestionsPutAuditQuestionWithDefaults

`func NewAuditQuestionsPutAuditQuestionWithDefaults() *AuditQuestionsPutAuditQuestion`

NewAuditQuestionsPutAuditQuestionWithDefaults instantiates a new AuditQuestionsPutAuditQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditQuestionsPutAuditQuestion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditQuestionsPutAuditQuestion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditQuestionsPutAuditQuestion) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuditQuestionsPutAuditQuestion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuditQuestionsPutAuditQuestion) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditQuestionsPutAuditQuestion) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditQuestionsPutAuditQuestion) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuditQuestionsPutAuditQuestion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuditQuestionsPutAuditQuestion) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuditQuestionsPutAuditQuestion) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuditQuestionsPutAuditQuestion) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuditQuestionsPutAuditQuestion) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AuditQuestionsPutAuditQuestion) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AuditQuestionsPutAuditQuestion) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AuditQuestionsPutAuditQuestion) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AuditQuestionsPutAuditQuestion) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetRequirements

`func (o *AuditQuestionsPutAuditQuestion) GetRequirements() interface{}`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *AuditQuestionsPutAuditQuestion) GetRequirementsOk() (*interface{}, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *AuditQuestionsPutAuditQuestion) SetRequirements(v interface{})`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *AuditQuestionsPutAuditQuestion) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### SetRequirementsNil

`func (o *AuditQuestionsPutAuditQuestion) SetRequirementsNil(b bool)`

 SetRequirementsNil sets the value for Requirements to be an explicit nil

### UnsetRequirements
`func (o *AuditQuestionsPutAuditQuestion) UnsetRequirements()`

UnsetRequirements ensures that no value is present for Requirements, not even an explicit nil
### GetType

`func (o *AuditQuestionsPutAuditQuestion) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditQuestionsPutAuditQuestion) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditQuestionsPutAuditQuestion) SetType(v string)`

SetType sets Type field to given value.


### GetQuestionData

`func (o *AuditQuestionsPutAuditQuestion) GetQuestionData() interface{}`

GetQuestionData returns the QuestionData field if non-nil, zero value otherwise.

### GetQuestionDataOk

`func (o *AuditQuestionsPutAuditQuestion) GetQuestionDataOk() (*interface{}, bool)`

GetQuestionDataOk returns a tuple with the QuestionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionData

`func (o *AuditQuestionsPutAuditQuestion) SetQuestionData(v interface{})`

SetQuestionData sets QuestionData field to given value.


### SetQuestionDataNil

`func (o *AuditQuestionsPutAuditQuestion) SetQuestionDataNil(b bool)`

 SetQuestionDataNil sets the value for QuestionData to be an explicit nil

### UnsetQuestionData
`func (o *AuditQuestionsPutAuditQuestion) UnsetQuestionData()`

UnsetQuestionData ensures that no value is present for QuestionData, not even an explicit nil
### GetText

`func (o *AuditQuestionsPutAuditQuestion) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AuditQuestionsPutAuditQuestion) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AuditQuestionsPutAuditQuestion) SetText(v string)`

SetText sets Text field to given value.


### GetDescription

`func (o *AuditQuestionsPutAuditQuestion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuditQuestionsPutAuditQuestion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuditQuestionsPutAuditQuestion) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuditQuestionsPutAuditQuestion) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSourceId

`func (o *AuditQuestionsPutAuditQuestion) GetSourceId() int32`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AuditQuestionsPutAuditQuestion) GetSourceIdOk() (*int32, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AuditQuestionsPutAuditQuestion) SetSourceId(v int32)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AuditQuestionsPutAuditQuestion) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *AuditQuestionsPutAuditQuestion) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *AuditQuestionsPutAuditQuestion) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *AuditQuestionsPutAuditQuestion) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *AuditQuestionsPutAuditQuestion) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetConditionChoiceUuid

`func (o *AuditQuestionsPutAuditQuestion) GetConditionChoiceUuid() string`

GetConditionChoiceUuid returns the ConditionChoiceUuid field if non-nil, zero value otherwise.

### GetConditionChoiceUuidOk

`func (o *AuditQuestionsPutAuditQuestion) GetConditionChoiceUuidOk() (*string, bool)`

GetConditionChoiceUuidOk returns a tuple with the ConditionChoiceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionChoiceUuid

`func (o *AuditQuestionsPutAuditQuestion) SetConditionChoiceUuid(v string)`

SetConditionChoiceUuid sets ConditionChoiceUuid field to given value.

### HasConditionChoiceUuid

`func (o *AuditQuestionsPutAuditQuestion) HasConditionChoiceUuid() bool`

HasConditionChoiceUuid returns a boolean if a field has been set.

### GetUid

`func (o *AuditQuestionsPutAuditQuestion) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AuditQuestionsPutAuditQuestion) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AuditQuestionsPutAuditQuestion) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AuditQuestionsPutAuditQuestion) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetChangeableType

`func (o *AuditQuestionsPutAuditQuestion) GetChangeableType() string`

GetChangeableType returns the ChangeableType field if non-nil, zero value otherwise.

### GetChangeableTypeOk

`func (o *AuditQuestionsPutAuditQuestion) GetChangeableTypeOk() (*string, bool)`

GetChangeableTypeOk returns a tuple with the ChangeableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeableType

`func (o *AuditQuestionsPutAuditQuestion) SetChangeableType(v string)`

SetChangeableType sets ChangeableType field to given value.

### HasChangeableType

`func (o *AuditQuestionsPutAuditQuestion) HasChangeableType() bool`

HasChangeableType returns a boolean if a field has been set.

### GetManagedChangeField

`func (o *AuditQuestionsPutAuditQuestion) GetManagedChangeField() string`

GetManagedChangeField returns the ManagedChangeField field if non-nil, zero value otherwise.

### GetManagedChangeFieldOk

`func (o *AuditQuestionsPutAuditQuestion) GetManagedChangeFieldOk() (*string, bool)`

GetManagedChangeFieldOk returns a tuple with the ManagedChangeField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedChangeField

`func (o *AuditQuestionsPutAuditQuestion) SetManagedChangeField(v string)`

SetManagedChangeField sets ManagedChangeField field to given value.

### HasManagedChangeField

`func (o *AuditQuestionsPutAuditQuestion) HasManagedChangeField() bool`

HasManagedChangeField returns a boolean if a field has been set.

### GetManagedChangeFieldType

`func (o *AuditQuestionsPutAuditQuestion) GetManagedChangeFieldType() string`

GetManagedChangeFieldType returns the ManagedChangeFieldType field if non-nil, zero value otherwise.

### GetManagedChangeFieldTypeOk

`func (o *AuditQuestionsPutAuditQuestion) GetManagedChangeFieldTypeOk() (*string, bool)`

GetManagedChangeFieldTypeOk returns a tuple with the ManagedChangeFieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedChangeFieldType

`func (o *AuditQuestionsPutAuditQuestion) SetManagedChangeFieldType(v string)`

SetManagedChangeFieldType sets ManagedChangeFieldType field to given value.

### HasManagedChangeFieldType

`func (o *AuditQuestionsPutAuditQuestion) HasManagedChangeFieldType() bool`

HasManagedChangeFieldType returns a boolean if a field has been set.

### GetQuestionOptions

`func (o *AuditQuestionsPutAuditQuestion) GetQuestionOptions() interface{}`

GetQuestionOptions returns the QuestionOptions field if non-nil, zero value otherwise.

### GetQuestionOptionsOk

`func (o *AuditQuestionsPutAuditQuestion) GetQuestionOptionsOk() (*interface{}, bool)`

GetQuestionOptionsOk returns a tuple with the QuestionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionOptions

`func (o *AuditQuestionsPutAuditQuestion) SetQuestionOptions(v interface{})`

SetQuestionOptions sets QuestionOptions field to given value.


### SetQuestionOptionsNil

`func (o *AuditQuestionsPutAuditQuestion) SetQuestionOptionsNil(b bool)`

 SetQuestionOptionsNil sets the value for QuestionOptions to be an explicit nil

### UnsetQuestionOptions
`func (o *AuditQuestionsPutAuditQuestion) UnsetQuestionOptions()`

UnsetQuestionOptions ensures that no value is present for QuestionOptions, not even an explicit nil
### GetAllowNaResponse

`func (o *AuditQuestionsPutAuditQuestion) GetAllowNaResponse() bool`

GetAllowNaResponse returns the AllowNaResponse field if non-nil, zero value otherwise.

### GetAllowNaResponseOk

`func (o *AuditQuestionsPutAuditQuestion) GetAllowNaResponseOk() (*bool, bool)`

GetAllowNaResponseOk returns a tuple with the AllowNaResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNaResponse

`func (o *AuditQuestionsPutAuditQuestion) SetAllowNaResponse(v bool)`

SetAllowNaResponse sets AllowNaResponse field to given value.


### GetAllowExplanation

`func (o *AuditQuestionsPutAuditQuestion) GetAllowExplanation() bool`

GetAllowExplanation returns the AllowExplanation field if non-nil, zero value otherwise.

### GetAllowExplanationOk

`func (o *AuditQuestionsPutAuditQuestion) GetAllowExplanationOk() (*bool, bool)`

GetAllowExplanationOk returns a tuple with the AllowExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExplanation

`func (o *AuditQuestionsPutAuditQuestion) SetAllowExplanation(v bool)`

SetAllowExplanation sets AllowExplanation field to given value.


### GetAllowFiles

`func (o *AuditQuestionsPutAuditQuestion) GetAllowFiles() bool`

GetAllowFiles returns the AllowFiles field if non-nil, zero value otherwise.

### GetAllowFilesOk

`func (o *AuditQuestionsPutAuditQuestion) GetAllowFilesOk() (*bool, bool)`

GetAllowFilesOk returns a tuple with the AllowFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFiles

`func (o *AuditQuestionsPutAuditQuestion) SetAllowFiles(v bool)`

SetAllowFiles sets AllowFiles field to given value.


### GetRequireFiles

`func (o *AuditQuestionsPutAuditQuestion) GetRequireFiles() bool`

GetRequireFiles returns the RequireFiles field if non-nil, zero value otherwise.

### GetRequireFilesOk

`func (o *AuditQuestionsPutAuditQuestion) GetRequireFilesOk() (*bool, bool)`

GetRequireFilesOk returns a tuple with the RequireFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireFiles

`func (o *AuditQuestionsPutAuditQuestion) SetRequireFiles(v bool)`

SetRequireFiles sets RequireFiles field to given value.


### GetQuestionUuid

`func (o *AuditQuestionsPutAuditQuestion) GetQuestionUuid() string`

GetQuestionUuid returns the QuestionUuid field if non-nil, zero value otherwise.

### GetQuestionUuidOk

`func (o *AuditQuestionsPutAuditQuestion) GetQuestionUuidOk() (*string, bool)`

GetQuestionUuidOk returns a tuple with the QuestionUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionUuid

`func (o *AuditQuestionsPutAuditQuestion) SetQuestionUuid(v string)`

SetQuestionUuid sets QuestionUuid field to given value.

### HasQuestionUuid

`func (o *AuditQuestionsPutAuditQuestion) HasQuestionUuid() bool`

HasQuestionUuid returns a boolean if a field has been set.

### GetConditionQuestionUuid

`func (o *AuditQuestionsPutAuditQuestion) GetConditionQuestionUuid() string`

GetConditionQuestionUuid returns the ConditionQuestionUuid field if non-nil, zero value otherwise.

### GetConditionQuestionUuidOk

`func (o *AuditQuestionsPutAuditQuestion) GetConditionQuestionUuidOk() (*string, bool)`

GetConditionQuestionUuidOk returns a tuple with the ConditionQuestionUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionQuestionUuid

`func (o *AuditQuestionsPutAuditQuestion) SetConditionQuestionUuid(v string)`

SetConditionQuestionUuid sets ConditionQuestionUuid field to given value.

### HasConditionQuestionUuid

`func (o *AuditQuestionsPutAuditQuestion) HasConditionQuestionUuid() bool`

HasConditionQuestionUuid returns a boolean if a field has been set.

### GetCondition

`func (o *AuditQuestionsPutAuditQuestion) GetCondition() interface{}`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AuditQuestionsPutAuditQuestion) GetConditionOk() (*interface{}, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AuditQuestionsPutAuditQuestion) SetCondition(v interface{})`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *AuditQuestionsPutAuditQuestion) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *AuditQuestionsPutAuditQuestion) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *AuditQuestionsPutAuditQuestion) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetRequireNaExplanation

`func (o *AuditQuestionsPutAuditQuestion) GetRequireNaExplanation() bool`

GetRequireNaExplanation returns the RequireNaExplanation field if non-nil, zero value otherwise.

### GetRequireNaExplanationOk

`func (o *AuditQuestionsPutAuditQuestion) GetRequireNaExplanationOk() (*bool, bool)`

GetRequireNaExplanationOk returns a tuple with the RequireNaExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireNaExplanation

`func (o *AuditQuestionsPutAuditQuestion) SetRequireNaExplanation(v bool)`

SetRequireNaExplanation sets RequireNaExplanation field to given value.

### HasRequireNaExplanation

`func (o *AuditQuestionsPutAuditQuestion) HasRequireNaExplanation() bool`

HasRequireNaExplanation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


