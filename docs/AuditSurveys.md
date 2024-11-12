# AuditSurveys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**AuditsurveyableId** | Pointer to **int32** |  | [optional] 
**AuditsurveyableType** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**CreatedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**QuestionSortOrder** | Pointer to **interface{}** |  | [optional] 
**SourceId** | Pointer to **int32** |  | [optional] 
**IsTemplate** | **bool** |  | [default to false]
**ReferenceMeta** | Pointer to **interface{}** |  | [optional] 
**Totals** | Pointer to **interface{}** |  | [optional] 
**QuestionOrder** | **interface{}** |  | 
**Score** | Pointer to **interface{}** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**TaskItemId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;task_items.id&#x60;.&lt;fk table&#x3D;&#39;task_items&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ScoringMethodology** | Pointer to **string** |  | [optional] 
**BracketsType** | Pointer to **string** |  | [optional] [default to "NULL"]
**Brackets** | **interface{}** |  | 
**AuditSurveyTemplateId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_survey_templates.id&#x60;.&lt;fk table&#x3D;&#39;audit_survey_templates&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**SectionOrder** | **interface{}** |  | 
**StashedQuestionOrder** | Pointer to **interface{}** |  | [optional] 
**RecurringSettings** | Pointer to **interface{}** |  | [optional] 
**IsWeighted** | Pointer to **bool** |  | [optional] [default to false]
**AuditSurveyCategoryId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_survey_categories.id&#x60;.&lt;fk table&#x3D;&#39;audit_survey_categories&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AuditSurveyScoreId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_survey_scores.id&#x60;.&lt;fk table&#x3D;&#39;audit_survey_scores&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 

## Methods

### NewAuditSurveys

`func NewAuditSurveys(name string, isTemplate bool, questionOrder interface{}, brackets interface{}, sectionOrder interface{}, ) *AuditSurveys`

NewAuditSurveys instantiates a new AuditSurveys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditSurveysWithDefaults

`func NewAuditSurveysWithDefaults() *AuditSurveys`

NewAuditSurveysWithDefaults instantiates a new AuditSurveys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditSurveys) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditSurveys) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditSurveys) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuditSurveys) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuditSurveys) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditSurveys) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditSurveys) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuditSurveys) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuditSurveys) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuditSurveys) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuditSurveys) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuditSurveys) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AuditSurveys) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AuditSurveys) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AuditSurveys) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AuditSurveys) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetAuditsurveyableId

`func (o *AuditSurveys) GetAuditsurveyableId() int32`

GetAuditsurveyableId returns the AuditsurveyableId field if non-nil, zero value otherwise.

### GetAuditsurveyableIdOk

`func (o *AuditSurveys) GetAuditsurveyableIdOk() (*int32, bool)`

GetAuditsurveyableIdOk returns a tuple with the AuditsurveyableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditsurveyableId

`func (o *AuditSurveys) SetAuditsurveyableId(v int32)`

SetAuditsurveyableId sets AuditsurveyableId field to given value.

### HasAuditsurveyableId

`func (o *AuditSurveys) HasAuditsurveyableId() bool`

HasAuditsurveyableId returns a boolean if a field has been set.

### GetAuditsurveyableType

`func (o *AuditSurveys) GetAuditsurveyableType() string`

GetAuditsurveyableType returns the AuditsurveyableType field if non-nil, zero value otherwise.

### GetAuditsurveyableTypeOk

`func (o *AuditSurveys) GetAuditsurveyableTypeOk() (*string, bool)`

GetAuditsurveyableTypeOk returns a tuple with the AuditsurveyableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditsurveyableType

`func (o *AuditSurveys) SetAuditsurveyableType(v string)`

SetAuditsurveyableType sets AuditsurveyableType field to given value.

### HasAuditsurveyableType

`func (o *AuditSurveys) HasAuditsurveyableType() bool`

HasAuditsurveyableType returns a boolean if a field has been set.

### GetName

`func (o *AuditSurveys) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditSurveys) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditSurveys) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedByUserId

`func (o *AuditSurveys) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *AuditSurveys) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *AuditSurveys) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *AuditSurveys) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetQuestionSortOrder

`func (o *AuditSurveys) GetQuestionSortOrder() interface{}`

GetQuestionSortOrder returns the QuestionSortOrder field if non-nil, zero value otherwise.

### GetQuestionSortOrderOk

`func (o *AuditSurveys) GetQuestionSortOrderOk() (*interface{}, bool)`

GetQuestionSortOrderOk returns a tuple with the QuestionSortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionSortOrder

`func (o *AuditSurveys) SetQuestionSortOrder(v interface{})`

SetQuestionSortOrder sets QuestionSortOrder field to given value.

### HasQuestionSortOrder

`func (o *AuditSurveys) HasQuestionSortOrder() bool`

HasQuestionSortOrder returns a boolean if a field has been set.

### SetQuestionSortOrderNil

`func (o *AuditSurveys) SetQuestionSortOrderNil(b bool)`

 SetQuestionSortOrderNil sets the value for QuestionSortOrder to be an explicit nil

### UnsetQuestionSortOrder
`func (o *AuditSurveys) UnsetQuestionSortOrder()`

UnsetQuestionSortOrder ensures that no value is present for QuestionSortOrder, not even an explicit nil
### GetSourceId

`func (o *AuditSurveys) GetSourceId() int32`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AuditSurveys) GetSourceIdOk() (*int32, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AuditSurveys) SetSourceId(v int32)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AuditSurveys) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetIsTemplate

`func (o *AuditSurveys) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *AuditSurveys) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *AuditSurveys) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.


### GetReferenceMeta

`func (o *AuditSurveys) GetReferenceMeta() interface{}`

GetReferenceMeta returns the ReferenceMeta field if non-nil, zero value otherwise.

### GetReferenceMetaOk

`func (o *AuditSurveys) GetReferenceMetaOk() (*interface{}, bool)`

GetReferenceMetaOk returns a tuple with the ReferenceMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMeta

`func (o *AuditSurveys) SetReferenceMeta(v interface{})`

SetReferenceMeta sets ReferenceMeta field to given value.

### HasReferenceMeta

`func (o *AuditSurveys) HasReferenceMeta() bool`

HasReferenceMeta returns a boolean if a field has been set.

### SetReferenceMetaNil

`func (o *AuditSurveys) SetReferenceMetaNil(b bool)`

 SetReferenceMetaNil sets the value for ReferenceMeta to be an explicit nil

### UnsetReferenceMeta
`func (o *AuditSurveys) UnsetReferenceMeta()`

UnsetReferenceMeta ensures that no value is present for ReferenceMeta, not even an explicit nil
### GetTotals

`func (o *AuditSurveys) GetTotals() interface{}`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *AuditSurveys) GetTotalsOk() (*interface{}, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *AuditSurveys) SetTotals(v interface{})`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *AuditSurveys) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### SetTotalsNil

`func (o *AuditSurveys) SetTotalsNil(b bool)`

 SetTotalsNil sets the value for Totals to be an explicit nil

### UnsetTotals
`func (o *AuditSurveys) UnsetTotals()`

UnsetTotals ensures that no value is present for Totals, not even an explicit nil
### GetQuestionOrder

`func (o *AuditSurveys) GetQuestionOrder() interface{}`

GetQuestionOrder returns the QuestionOrder field if non-nil, zero value otherwise.

### GetQuestionOrderOk

`func (o *AuditSurveys) GetQuestionOrderOk() (*interface{}, bool)`

GetQuestionOrderOk returns a tuple with the QuestionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionOrder

`func (o *AuditSurveys) SetQuestionOrder(v interface{})`

SetQuestionOrder sets QuestionOrder field to given value.


### SetQuestionOrderNil

`func (o *AuditSurveys) SetQuestionOrderNil(b bool)`

 SetQuestionOrderNil sets the value for QuestionOrder to be an explicit nil

### UnsetQuestionOrder
`func (o *AuditSurveys) UnsetQuestionOrder()`

UnsetQuestionOrder ensures that no value is present for QuestionOrder, not even an explicit nil
### GetScore

`func (o *AuditSurveys) GetScore() interface{}`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *AuditSurveys) GetScoreOk() (*interface{}, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *AuditSurveys) SetScore(v interface{})`

SetScore sets Score field to given value.

### HasScore

`func (o *AuditSurveys) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *AuditSurveys) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *AuditSurveys) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetType

`func (o *AuditSurveys) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditSurveys) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditSurveys) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuditSurveys) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTaskItemId

`func (o *AuditSurveys) GetTaskItemId() int32`

GetTaskItemId returns the TaskItemId field if non-nil, zero value otherwise.

### GetTaskItemIdOk

`func (o *AuditSurveys) GetTaskItemIdOk() (*int32, bool)`

GetTaskItemIdOk returns a tuple with the TaskItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskItemId

`func (o *AuditSurveys) SetTaskItemId(v int32)`

SetTaskItemId sets TaskItemId field to given value.

### HasTaskItemId

`func (o *AuditSurveys) HasTaskItemId() bool`

HasTaskItemId returns a boolean if a field has been set.

### GetScoringMethodology

`func (o *AuditSurveys) GetScoringMethodology() string`

GetScoringMethodology returns the ScoringMethodology field if non-nil, zero value otherwise.

### GetScoringMethodologyOk

`func (o *AuditSurveys) GetScoringMethodologyOk() (*string, bool)`

GetScoringMethodologyOk returns a tuple with the ScoringMethodology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoringMethodology

`func (o *AuditSurveys) SetScoringMethodology(v string)`

SetScoringMethodology sets ScoringMethodology field to given value.

### HasScoringMethodology

`func (o *AuditSurveys) HasScoringMethodology() bool`

HasScoringMethodology returns a boolean if a field has been set.

### GetBracketsType

`func (o *AuditSurveys) GetBracketsType() string`

GetBracketsType returns the BracketsType field if non-nil, zero value otherwise.

### GetBracketsTypeOk

`func (o *AuditSurveys) GetBracketsTypeOk() (*string, bool)`

GetBracketsTypeOk returns a tuple with the BracketsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBracketsType

`func (o *AuditSurveys) SetBracketsType(v string)`

SetBracketsType sets BracketsType field to given value.

### HasBracketsType

`func (o *AuditSurveys) HasBracketsType() bool`

HasBracketsType returns a boolean if a field has been set.

### GetBrackets

`func (o *AuditSurveys) GetBrackets() interface{}`

GetBrackets returns the Brackets field if non-nil, zero value otherwise.

### GetBracketsOk

`func (o *AuditSurveys) GetBracketsOk() (*interface{}, bool)`

GetBracketsOk returns a tuple with the Brackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrackets

`func (o *AuditSurveys) SetBrackets(v interface{})`

SetBrackets sets Brackets field to given value.


### SetBracketsNil

`func (o *AuditSurveys) SetBracketsNil(b bool)`

 SetBracketsNil sets the value for Brackets to be an explicit nil

### UnsetBrackets
`func (o *AuditSurveys) UnsetBrackets()`

UnsetBrackets ensures that no value is present for Brackets, not even an explicit nil
### GetAuditSurveyTemplateId

`func (o *AuditSurveys) GetAuditSurveyTemplateId() int32`

GetAuditSurveyTemplateId returns the AuditSurveyTemplateId field if non-nil, zero value otherwise.

### GetAuditSurveyTemplateIdOk

`func (o *AuditSurveys) GetAuditSurveyTemplateIdOk() (*int32, bool)`

GetAuditSurveyTemplateIdOk returns a tuple with the AuditSurveyTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSurveyTemplateId

`func (o *AuditSurveys) SetAuditSurveyTemplateId(v int32)`

SetAuditSurveyTemplateId sets AuditSurveyTemplateId field to given value.

### HasAuditSurveyTemplateId

`func (o *AuditSurveys) HasAuditSurveyTemplateId() bool`

HasAuditSurveyTemplateId returns a boolean if a field has been set.

### GetSectionOrder

`func (o *AuditSurveys) GetSectionOrder() interface{}`

GetSectionOrder returns the SectionOrder field if non-nil, zero value otherwise.

### GetSectionOrderOk

`func (o *AuditSurveys) GetSectionOrderOk() (*interface{}, bool)`

GetSectionOrderOk returns a tuple with the SectionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionOrder

`func (o *AuditSurveys) SetSectionOrder(v interface{})`

SetSectionOrder sets SectionOrder field to given value.


### SetSectionOrderNil

`func (o *AuditSurveys) SetSectionOrderNil(b bool)`

 SetSectionOrderNil sets the value for SectionOrder to be an explicit nil

### UnsetSectionOrder
`func (o *AuditSurveys) UnsetSectionOrder()`

UnsetSectionOrder ensures that no value is present for SectionOrder, not even an explicit nil
### GetStashedQuestionOrder

`func (o *AuditSurveys) GetStashedQuestionOrder() interface{}`

GetStashedQuestionOrder returns the StashedQuestionOrder field if non-nil, zero value otherwise.

### GetStashedQuestionOrderOk

`func (o *AuditSurveys) GetStashedQuestionOrderOk() (*interface{}, bool)`

GetStashedQuestionOrderOk returns a tuple with the StashedQuestionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStashedQuestionOrder

`func (o *AuditSurveys) SetStashedQuestionOrder(v interface{})`

SetStashedQuestionOrder sets StashedQuestionOrder field to given value.

### HasStashedQuestionOrder

`func (o *AuditSurveys) HasStashedQuestionOrder() bool`

HasStashedQuestionOrder returns a boolean if a field has been set.

### SetStashedQuestionOrderNil

`func (o *AuditSurveys) SetStashedQuestionOrderNil(b bool)`

 SetStashedQuestionOrderNil sets the value for StashedQuestionOrder to be an explicit nil

### UnsetStashedQuestionOrder
`func (o *AuditSurveys) UnsetStashedQuestionOrder()`

UnsetStashedQuestionOrder ensures that no value is present for StashedQuestionOrder, not even an explicit nil
### GetRecurringSettings

`func (o *AuditSurveys) GetRecurringSettings() interface{}`

GetRecurringSettings returns the RecurringSettings field if non-nil, zero value otherwise.

### GetRecurringSettingsOk

`func (o *AuditSurveys) GetRecurringSettingsOk() (*interface{}, bool)`

GetRecurringSettingsOk returns a tuple with the RecurringSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringSettings

`func (o *AuditSurveys) SetRecurringSettings(v interface{})`

SetRecurringSettings sets RecurringSettings field to given value.

### HasRecurringSettings

`func (o *AuditSurveys) HasRecurringSettings() bool`

HasRecurringSettings returns a boolean if a field has been set.

### SetRecurringSettingsNil

`func (o *AuditSurveys) SetRecurringSettingsNil(b bool)`

 SetRecurringSettingsNil sets the value for RecurringSettings to be an explicit nil

### UnsetRecurringSettings
`func (o *AuditSurveys) UnsetRecurringSettings()`

UnsetRecurringSettings ensures that no value is present for RecurringSettings, not even an explicit nil
### GetIsWeighted

`func (o *AuditSurveys) GetIsWeighted() bool`

GetIsWeighted returns the IsWeighted field if non-nil, zero value otherwise.

### GetIsWeightedOk

`func (o *AuditSurveys) GetIsWeightedOk() (*bool, bool)`

GetIsWeightedOk returns a tuple with the IsWeighted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWeighted

`func (o *AuditSurveys) SetIsWeighted(v bool)`

SetIsWeighted sets IsWeighted field to given value.

### HasIsWeighted

`func (o *AuditSurveys) HasIsWeighted() bool`

HasIsWeighted returns a boolean if a field has been set.

### GetAuditSurveyCategoryId

`func (o *AuditSurveys) GetAuditSurveyCategoryId() int32`

GetAuditSurveyCategoryId returns the AuditSurveyCategoryId field if non-nil, zero value otherwise.

### GetAuditSurveyCategoryIdOk

`func (o *AuditSurveys) GetAuditSurveyCategoryIdOk() (*int32, bool)`

GetAuditSurveyCategoryIdOk returns a tuple with the AuditSurveyCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSurveyCategoryId

`func (o *AuditSurveys) SetAuditSurveyCategoryId(v int32)`

SetAuditSurveyCategoryId sets AuditSurveyCategoryId field to given value.

### HasAuditSurveyCategoryId

`func (o *AuditSurveys) HasAuditSurveyCategoryId() bool`

HasAuditSurveyCategoryId returns a boolean if a field has been set.

### GetAuditSurveyScoreId

`func (o *AuditSurveys) GetAuditSurveyScoreId() int32`

GetAuditSurveyScoreId returns the AuditSurveyScoreId field if non-nil, zero value otherwise.

### GetAuditSurveyScoreIdOk

`func (o *AuditSurveys) GetAuditSurveyScoreIdOk() (*int32, bool)`

GetAuditSurveyScoreIdOk returns a tuple with the AuditSurveyScoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSurveyScoreId

`func (o *AuditSurveys) SetAuditSurveyScoreId(v int32)`

SetAuditSurveyScoreId sets AuditSurveyScoreId field to given value.

### HasAuditSurveyScoreId

`func (o *AuditSurveys) HasAuditSurveyScoreId() bool`

HasAuditSurveyScoreId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


