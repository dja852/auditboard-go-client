# QuestionResponses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**QuestionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;questions.id&#x60;.&lt;fk table&#x3D;&#39;questions&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**SurveyResponseId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;survey_responses.id&#x60;.&lt;fk table&#x3D;&#39;survey_responses&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Response** | Pointer to **string** |  | [optional] 
**MetaConfig** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**ResponseExplanation** | Pointer to **string** |  | [optional] [default to ""]
**ResponseData** | Pointer to **string** |  | [optional] 
**ReferenceText** | Pointer to **string** |  | [optional] [default to ""]
**MetaFlat** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**AssessmentResponseId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;assessment_responses.id&#x60;.&lt;fk table&#x3D;&#39;assessment_responses&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FlatProperties** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewQuestionResponses

`func NewQuestionResponses() *QuestionResponses`

NewQuestionResponses instantiates a new QuestionResponses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionResponsesWithDefaults

`func NewQuestionResponsesWithDefaults() *QuestionResponses`

NewQuestionResponsesWithDefaults instantiates a new QuestionResponses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuestionResponses) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuestionResponses) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuestionResponses) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *QuestionResponses) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQuestionId

`func (o *QuestionResponses) GetQuestionId() int32`

GetQuestionId returns the QuestionId field if non-nil, zero value otherwise.

### GetQuestionIdOk

`func (o *QuestionResponses) GetQuestionIdOk() (*int32, bool)`

GetQuestionIdOk returns a tuple with the QuestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionId

`func (o *QuestionResponses) SetQuestionId(v int32)`

SetQuestionId sets QuestionId field to given value.

### HasQuestionId

`func (o *QuestionResponses) HasQuestionId() bool`

HasQuestionId returns a boolean if a field has been set.

### GetSurveyResponseId

`func (o *QuestionResponses) GetSurveyResponseId() int32`

GetSurveyResponseId returns the SurveyResponseId field if non-nil, zero value otherwise.

### GetSurveyResponseIdOk

`func (o *QuestionResponses) GetSurveyResponseIdOk() (*int32, bool)`

GetSurveyResponseIdOk returns a tuple with the SurveyResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyResponseId

`func (o *QuestionResponses) SetSurveyResponseId(v int32)`

SetSurveyResponseId sets SurveyResponseId field to given value.

### HasSurveyResponseId

`func (o *QuestionResponses) HasSurveyResponseId() bool`

HasSurveyResponseId returns a boolean if a field has been set.

### GetStatus

`func (o *QuestionResponses) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QuestionResponses) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QuestionResponses) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QuestionResponses) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResponse

`func (o *QuestionResponses) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *QuestionResponses) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *QuestionResponses) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *QuestionResponses) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetMetaConfig

`func (o *QuestionResponses) GetMetaConfig() string`

GetMetaConfig returns the MetaConfig field if non-nil, zero value otherwise.

### GetMetaConfigOk

`func (o *QuestionResponses) GetMetaConfigOk() (*string, bool)`

GetMetaConfigOk returns a tuple with the MetaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaConfig

`func (o *QuestionResponses) SetMetaConfig(v string)`

SetMetaConfig sets MetaConfig field to given value.

### HasMetaConfig

`func (o *QuestionResponses) HasMetaConfig() bool`

HasMetaConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *QuestionResponses) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QuestionResponses) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QuestionResponses) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *QuestionResponses) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *QuestionResponses) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QuestionResponses) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QuestionResponses) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *QuestionResponses) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *QuestionResponses) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *QuestionResponses) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *QuestionResponses) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *QuestionResponses) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetResponseExplanation

`func (o *QuestionResponses) GetResponseExplanation() string`

GetResponseExplanation returns the ResponseExplanation field if non-nil, zero value otherwise.

### GetResponseExplanationOk

`func (o *QuestionResponses) GetResponseExplanationOk() (*string, bool)`

GetResponseExplanationOk returns a tuple with the ResponseExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseExplanation

`func (o *QuestionResponses) SetResponseExplanation(v string)`

SetResponseExplanation sets ResponseExplanation field to given value.

### HasResponseExplanation

`func (o *QuestionResponses) HasResponseExplanation() bool`

HasResponseExplanation returns a boolean if a field has been set.

### GetResponseData

`func (o *QuestionResponses) GetResponseData() string`

GetResponseData returns the ResponseData field if non-nil, zero value otherwise.

### GetResponseDataOk

`func (o *QuestionResponses) GetResponseDataOk() (*string, bool)`

GetResponseDataOk returns a tuple with the ResponseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseData

`func (o *QuestionResponses) SetResponseData(v string)`

SetResponseData sets ResponseData field to given value.

### HasResponseData

`func (o *QuestionResponses) HasResponseData() bool`

HasResponseData returns a boolean if a field has been set.

### GetReferenceText

`func (o *QuestionResponses) GetReferenceText() string`

GetReferenceText returns the ReferenceText field if non-nil, zero value otherwise.

### GetReferenceTextOk

`func (o *QuestionResponses) GetReferenceTextOk() (*string, bool)`

GetReferenceTextOk returns a tuple with the ReferenceText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceText

`func (o *QuestionResponses) SetReferenceText(v string)`

SetReferenceText sets ReferenceText field to given value.

### HasReferenceText

`func (o *QuestionResponses) HasReferenceText() bool`

HasReferenceText returns a boolean if a field has been set.

### GetMetaFlat

`func (o *QuestionResponses) GetMetaFlat() string`

GetMetaFlat returns the MetaFlat field if non-nil, zero value otherwise.

### GetMetaFlatOk

`func (o *QuestionResponses) GetMetaFlatOk() (*string, bool)`

GetMetaFlatOk returns a tuple with the MetaFlat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaFlat

`func (o *QuestionResponses) SetMetaFlat(v string)`

SetMetaFlat sets MetaFlat field to given value.

### HasMetaFlat

`func (o *QuestionResponses) HasMetaFlat() bool`

HasMetaFlat returns a boolean if a field has been set.

### GetType

`func (o *QuestionResponses) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QuestionResponses) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QuestionResponses) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *QuestionResponses) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAssessmentResponseId

`func (o *QuestionResponses) GetAssessmentResponseId() int32`

GetAssessmentResponseId returns the AssessmentResponseId field if non-nil, zero value otherwise.

### GetAssessmentResponseIdOk

`func (o *QuestionResponses) GetAssessmentResponseIdOk() (*int32, bool)`

GetAssessmentResponseIdOk returns a tuple with the AssessmentResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessmentResponseId

`func (o *QuestionResponses) SetAssessmentResponseId(v int32)`

SetAssessmentResponseId sets AssessmentResponseId field to given value.

### HasAssessmentResponseId

`func (o *QuestionResponses) HasAssessmentResponseId() bool`

HasAssessmentResponseId returns a boolean if a field has been set.

### GetFlatProperties

`func (o *QuestionResponses) GetFlatProperties() interface{}`

GetFlatProperties returns the FlatProperties field if non-nil, zero value otherwise.

### GetFlatPropertiesOk

`func (o *QuestionResponses) GetFlatPropertiesOk() (*interface{}, bool)`

GetFlatPropertiesOk returns a tuple with the FlatProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatProperties

`func (o *QuestionResponses) SetFlatProperties(v interface{})`

SetFlatProperties sets FlatProperties field to given value.

### HasFlatProperties

`func (o *QuestionResponses) HasFlatProperties() bool`

HasFlatProperties returns a boolean if a field has been set.

### SetFlatPropertiesNil

`func (o *QuestionResponses) SetFlatPropertiesNil(b bool)`

 SetFlatPropertiesNil sets the value for FlatProperties to be an explicit nil

### UnsetFlatProperties
`func (o *QuestionResponses) UnsetFlatProperties()`

UnsetFlatProperties ensures that no value is present for FlatProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


