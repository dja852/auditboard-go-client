# Assessments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**AssessmentTemplateId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;assessment_templates.id&#x60;.&lt;fk table&#x3D;&#39;assessment_templates&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FinalizedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FinalizedDate** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**StartedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**StartedDate** | Pointer to **string** |  | [optional] 
**CanceledByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CanceledDate** | Pointer to **string** |  | [optional] 
**SurveyType** | Pointer to **string** |  | [optional] 
**ProjectOptions** | Pointer to **interface{}** |  | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**EmailProjectStart** | Pointer to **interface{}** |  | [optional] 
**LandingText** | Pointer to **string** |  | [optional] 
**ConfirmText** | Pointer to **string** |  | [optional] 
**InstructionText** | Pointer to **string** |  | [optional] 
**AssessmentPeriodId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;assessment_periods.id&#x60;.&lt;fk table&#x3D;&#39;assessment_periods&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Assessables** | Pointer to **interface{}** |  | [optional] 
**Info** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewAssessments

`func NewAssessments() *Assessments`

NewAssessments instantiates a new Assessments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssessmentsWithDefaults

`func NewAssessmentsWithDefaults() *Assessments`

NewAssessmentsWithDefaults instantiates a new Assessments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Assessments) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Assessments) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Assessments) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Assessments) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Assessments) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Assessments) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Assessments) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Assessments) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *Assessments) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Assessments) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Assessments) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Assessments) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *Assessments) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Assessments) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Assessments) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Assessments) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAssessmentTemplateId

`func (o *Assessments) GetAssessmentTemplateId() int32`

GetAssessmentTemplateId returns the AssessmentTemplateId field if non-nil, zero value otherwise.

### GetAssessmentTemplateIdOk

`func (o *Assessments) GetAssessmentTemplateIdOk() (*int32, bool)`

GetAssessmentTemplateIdOk returns a tuple with the AssessmentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessmentTemplateId

`func (o *Assessments) SetAssessmentTemplateId(v int32)`

SetAssessmentTemplateId sets AssessmentTemplateId field to given value.

### HasAssessmentTemplateId

`func (o *Assessments) HasAssessmentTemplateId() bool`

HasAssessmentTemplateId returns a boolean if a field has been set.

### GetFinalizedByUserId

`func (o *Assessments) GetFinalizedByUserId() int32`

GetFinalizedByUserId returns the FinalizedByUserId field if non-nil, zero value otherwise.

### GetFinalizedByUserIdOk

`func (o *Assessments) GetFinalizedByUserIdOk() (*int32, bool)`

GetFinalizedByUserIdOk returns a tuple with the FinalizedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizedByUserId

`func (o *Assessments) SetFinalizedByUserId(v int32)`

SetFinalizedByUserId sets FinalizedByUserId field to given value.

### HasFinalizedByUserId

`func (o *Assessments) HasFinalizedByUserId() bool`

HasFinalizedByUserId returns a boolean if a field has been set.

### GetFinalizedDate

`func (o *Assessments) GetFinalizedDate() string`

GetFinalizedDate returns the FinalizedDate field if non-nil, zero value otherwise.

### GetFinalizedDateOk

`func (o *Assessments) GetFinalizedDateOk() (*string, bool)`

GetFinalizedDateOk returns a tuple with the FinalizedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizedDate

`func (o *Assessments) SetFinalizedDate(v string)`

SetFinalizedDate sets FinalizedDate field to given value.

### HasFinalizedDate

`func (o *Assessments) HasFinalizedDate() bool`

HasFinalizedDate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Assessments) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Assessments) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Assessments) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Assessments) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Assessments) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Assessments) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Assessments) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Assessments) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Assessments) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Assessments) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Assessments) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Assessments) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetStartedByUserId

`func (o *Assessments) GetStartedByUserId() int32`

GetStartedByUserId returns the StartedByUserId field if non-nil, zero value otherwise.

### GetStartedByUserIdOk

`func (o *Assessments) GetStartedByUserIdOk() (*int32, bool)`

GetStartedByUserIdOk returns a tuple with the StartedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedByUserId

`func (o *Assessments) SetStartedByUserId(v int32)`

SetStartedByUserId sets StartedByUserId field to given value.

### HasStartedByUserId

`func (o *Assessments) HasStartedByUserId() bool`

HasStartedByUserId returns a boolean if a field has been set.

### GetStartedDate

`func (o *Assessments) GetStartedDate() string`

GetStartedDate returns the StartedDate field if non-nil, zero value otherwise.

### GetStartedDateOk

`func (o *Assessments) GetStartedDateOk() (*string, bool)`

GetStartedDateOk returns a tuple with the StartedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedDate

`func (o *Assessments) SetStartedDate(v string)`

SetStartedDate sets StartedDate field to given value.

### HasStartedDate

`func (o *Assessments) HasStartedDate() bool`

HasStartedDate returns a boolean if a field has been set.

### GetCanceledByUserId

`func (o *Assessments) GetCanceledByUserId() int32`

GetCanceledByUserId returns the CanceledByUserId field if non-nil, zero value otherwise.

### GetCanceledByUserIdOk

`func (o *Assessments) GetCanceledByUserIdOk() (*int32, bool)`

GetCanceledByUserIdOk returns a tuple with the CanceledByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledByUserId

`func (o *Assessments) SetCanceledByUserId(v int32)`

SetCanceledByUserId sets CanceledByUserId field to given value.

### HasCanceledByUserId

`func (o *Assessments) HasCanceledByUserId() bool`

HasCanceledByUserId returns a boolean if a field has been set.

### GetCanceledDate

`func (o *Assessments) GetCanceledDate() string`

GetCanceledDate returns the CanceledDate field if non-nil, zero value otherwise.

### GetCanceledDateOk

`func (o *Assessments) GetCanceledDateOk() (*string, bool)`

GetCanceledDateOk returns a tuple with the CanceledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledDate

`func (o *Assessments) SetCanceledDate(v string)`

SetCanceledDate sets CanceledDate field to given value.

### HasCanceledDate

`func (o *Assessments) HasCanceledDate() bool`

HasCanceledDate returns a boolean if a field has been set.

### GetSurveyType

`func (o *Assessments) GetSurveyType() string`

GetSurveyType returns the SurveyType field if non-nil, zero value otherwise.

### GetSurveyTypeOk

`func (o *Assessments) GetSurveyTypeOk() (*string, bool)`

GetSurveyTypeOk returns a tuple with the SurveyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyType

`func (o *Assessments) SetSurveyType(v string)`

SetSurveyType sets SurveyType field to given value.

### HasSurveyType

`func (o *Assessments) HasSurveyType() bool`

HasSurveyType returns a boolean if a field has been set.

### GetProjectOptions

`func (o *Assessments) GetProjectOptions() interface{}`

GetProjectOptions returns the ProjectOptions field if non-nil, zero value otherwise.

### GetProjectOptionsOk

`func (o *Assessments) GetProjectOptionsOk() (*interface{}, bool)`

GetProjectOptionsOk returns a tuple with the ProjectOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectOptions

`func (o *Assessments) SetProjectOptions(v interface{})`

SetProjectOptions sets ProjectOptions field to given value.

### HasProjectOptions

`func (o *Assessments) HasProjectOptions() bool`

HasProjectOptions returns a boolean if a field has been set.

### SetProjectOptionsNil

`func (o *Assessments) SetProjectOptionsNil(b bool)`

 SetProjectOptionsNil sets the value for ProjectOptions to be an explicit nil

### UnsetProjectOptions
`func (o *Assessments) UnsetProjectOptions()`

UnsetProjectOptions ensures that no value is present for ProjectOptions, not even an explicit nil
### GetDueDate

`func (o *Assessments) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Assessments) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Assessments) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *Assessments) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetEmailProjectStart

`func (o *Assessments) GetEmailProjectStart() interface{}`

GetEmailProjectStart returns the EmailProjectStart field if non-nil, zero value otherwise.

### GetEmailProjectStartOk

`func (o *Assessments) GetEmailProjectStartOk() (*interface{}, bool)`

GetEmailProjectStartOk returns a tuple with the EmailProjectStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailProjectStart

`func (o *Assessments) SetEmailProjectStart(v interface{})`

SetEmailProjectStart sets EmailProjectStart field to given value.

### HasEmailProjectStart

`func (o *Assessments) HasEmailProjectStart() bool`

HasEmailProjectStart returns a boolean if a field has been set.

### SetEmailProjectStartNil

`func (o *Assessments) SetEmailProjectStartNil(b bool)`

 SetEmailProjectStartNil sets the value for EmailProjectStart to be an explicit nil

### UnsetEmailProjectStart
`func (o *Assessments) UnsetEmailProjectStart()`

UnsetEmailProjectStart ensures that no value is present for EmailProjectStart, not even an explicit nil
### GetLandingText

`func (o *Assessments) GetLandingText() string`

GetLandingText returns the LandingText field if non-nil, zero value otherwise.

### GetLandingTextOk

`func (o *Assessments) GetLandingTextOk() (*string, bool)`

GetLandingTextOk returns a tuple with the LandingText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingText

`func (o *Assessments) SetLandingText(v string)`

SetLandingText sets LandingText field to given value.

### HasLandingText

`func (o *Assessments) HasLandingText() bool`

HasLandingText returns a boolean if a field has been set.

### GetConfirmText

`func (o *Assessments) GetConfirmText() string`

GetConfirmText returns the ConfirmText field if non-nil, zero value otherwise.

### GetConfirmTextOk

`func (o *Assessments) GetConfirmTextOk() (*string, bool)`

GetConfirmTextOk returns a tuple with the ConfirmText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmText

`func (o *Assessments) SetConfirmText(v string)`

SetConfirmText sets ConfirmText field to given value.

### HasConfirmText

`func (o *Assessments) HasConfirmText() bool`

HasConfirmText returns a boolean if a field has been set.

### GetInstructionText

`func (o *Assessments) GetInstructionText() string`

GetInstructionText returns the InstructionText field if non-nil, zero value otherwise.

### GetInstructionTextOk

`func (o *Assessments) GetInstructionTextOk() (*string, bool)`

GetInstructionTextOk returns a tuple with the InstructionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructionText

`func (o *Assessments) SetInstructionText(v string)`

SetInstructionText sets InstructionText field to given value.

### HasInstructionText

`func (o *Assessments) HasInstructionText() bool`

HasInstructionText returns a boolean if a field has been set.

### GetAssessmentPeriodId

`func (o *Assessments) GetAssessmentPeriodId() int32`

GetAssessmentPeriodId returns the AssessmentPeriodId field if non-nil, zero value otherwise.

### GetAssessmentPeriodIdOk

`func (o *Assessments) GetAssessmentPeriodIdOk() (*int32, bool)`

GetAssessmentPeriodIdOk returns a tuple with the AssessmentPeriodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessmentPeriodId

`func (o *Assessments) SetAssessmentPeriodId(v int32)`

SetAssessmentPeriodId sets AssessmentPeriodId field to given value.

### HasAssessmentPeriodId

`func (o *Assessments) HasAssessmentPeriodId() bool`

HasAssessmentPeriodId returns a boolean if a field has been set.

### GetAssessables

`func (o *Assessments) GetAssessables() interface{}`

GetAssessables returns the Assessables field if non-nil, zero value otherwise.

### GetAssessablesOk

`func (o *Assessments) GetAssessablesOk() (*interface{}, bool)`

GetAssessablesOk returns a tuple with the Assessables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessables

`func (o *Assessments) SetAssessables(v interface{})`

SetAssessables sets Assessables field to given value.

### HasAssessables

`func (o *Assessments) HasAssessables() bool`

HasAssessables returns a boolean if a field has been set.

### SetAssessablesNil

`func (o *Assessments) SetAssessablesNil(b bool)`

 SetAssessablesNil sets the value for Assessables to be an explicit nil

### UnsetAssessables
`func (o *Assessments) UnsetAssessables()`

UnsetAssessables ensures that no value is present for Assessables, not even an explicit nil
### GetInfo

`func (o *Assessments) GetInfo() interface{}`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Assessments) GetInfoOk() (*interface{}, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Assessments) SetInfo(v interface{})`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Assessments) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### SetInfoNil

`func (o *Assessments) SetInfoNil(b bool)`

 SetInfoNil sets the value for Info to be an explicit nil

### UnsetInfo
`func (o *Assessments) UnsetInfo()`

UnsetInfo ensures that no value is present for Info, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


