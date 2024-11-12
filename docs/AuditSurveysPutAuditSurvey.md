# AuditSurveysPutAuditSurvey

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

### NewAuditSurveysPutAuditSurvey

`func NewAuditSurveysPutAuditSurvey(name string, isTemplate bool, questionOrder interface{}, brackets interface{}, sectionOrder interface{}, ) *AuditSurveysPutAuditSurvey`

NewAuditSurveysPutAuditSurvey instantiates a new AuditSurveysPutAuditSurvey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditSurveysPutAuditSurveyWithDefaults

`func NewAuditSurveysPutAuditSurveyWithDefaults() *AuditSurveysPutAuditSurvey`

NewAuditSurveysPutAuditSurveyWithDefaults instantiates a new AuditSurveysPutAuditSurvey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditSurveysPutAuditSurvey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditSurveysPutAuditSurvey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditSurveysPutAuditSurvey) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuditSurveysPutAuditSurvey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuditSurveysPutAuditSurvey) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditSurveysPutAuditSurvey) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditSurveysPutAuditSurvey) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuditSurveysPutAuditSurvey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuditSurveysPutAuditSurvey) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuditSurveysPutAuditSurvey) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuditSurveysPutAuditSurvey) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuditSurveysPutAuditSurvey) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AuditSurveysPutAuditSurvey) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AuditSurveysPutAuditSurvey) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AuditSurveysPutAuditSurvey) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AuditSurveysPutAuditSurvey) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetAuditsurveyableId

`func (o *AuditSurveysPutAuditSurvey) GetAuditsurveyableId() int32`

GetAuditsurveyableId returns the AuditsurveyableId field if non-nil, zero value otherwise.

### GetAuditsurveyableIdOk

`func (o *AuditSurveysPutAuditSurvey) GetAuditsurveyableIdOk() (*int32, bool)`

GetAuditsurveyableIdOk returns a tuple with the AuditsurveyableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditsurveyableId

`func (o *AuditSurveysPutAuditSurvey) SetAuditsurveyableId(v int32)`

SetAuditsurveyableId sets AuditsurveyableId field to given value.

### HasAuditsurveyableId

`func (o *AuditSurveysPutAuditSurvey) HasAuditsurveyableId() bool`

HasAuditsurveyableId returns a boolean if a field has been set.

### GetAuditsurveyableType

`func (o *AuditSurveysPutAuditSurvey) GetAuditsurveyableType() string`

GetAuditsurveyableType returns the AuditsurveyableType field if non-nil, zero value otherwise.

### GetAuditsurveyableTypeOk

`func (o *AuditSurveysPutAuditSurvey) GetAuditsurveyableTypeOk() (*string, bool)`

GetAuditsurveyableTypeOk returns a tuple with the AuditsurveyableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditsurveyableType

`func (o *AuditSurveysPutAuditSurvey) SetAuditsurveyableType(v string)`

SetAuditsurveyableType sets AuditsurveyableType field to given value.

### HasAuditsurveyableType

`func (o *AuditSurveysPutAuditSurvey) HasAuditsurveyableType() bool`

HasAuditsurveyableType returns a boolean if a field has been set.

### GetName

`func (o *AuditSurveysPutAuditSurvey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditSurveysPutAuditSurvey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditSurveysPutAuditSurvey) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedByUserId

`func (o *AuditSurveysPutAuditSurvey) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *AuditSurveysPutAuditSurvey) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *AuditSurveysPutAuditSurvey) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *AuditSurveysPutAuditSurvey) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetQuestionSortOrder

`func (o *AuditSurveysPutAuditSurvey) GetQuestionSortOrder() interface{}`

GetQuestionSortOrder returns the QuestionSortOrder field if non-nil, zero value otherwise.

### GetQuestionSortOrderOk

`func (o *AuditSurveysPutAuditSurvey) GetQuestionSortOrderOk() (*interface{}, bool)`

GetQuestionSortOrderOk returns a tuple with the QuestionSortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionSortOrder

`func (o *AuditSurveysPutAuditSurvey) SetQuestionSortOrder(v interface{})`

SetQuestionSortOrder sets QuestionSortOrder field to given value.

### HasQuestionSortOrder

`func (o *AuditSurveysPutAuditSurvey) HasQuestionSortOrder() bool`

HasQuestionSortOrder returns a boolean if a field has been set.

### SetQuestionSortOrderNil

`func (o *AuditSurveysPutAuditSurvey) SetQuestionSortOrderNil(b bool)`

 SetQuestionSortOrderNil sets the value for QuestionSortOrder to be an explicit nil

### UnsetQuestionSortOrder
`func (o *AuditSurveysPutAuditSurvey) UnsetQuestionSortOrder()`

UnsetQuestionSortOrder ensures that no value is present for QuestionSortOrder, not even an explicit nil
### GetSourceId

`func (o *AuditSurveysPutAuditSurvey) GetSourceId() int32`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AuditSurveysPutAuditSurvey) GetSourceIdOk() (*int32, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AuditSurveysPutAuditSurvey) SetSourceId(v int32)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AuditSurveysPutAuditSurvey) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetIsTemplate

`func (o *AuditSurveysPutAuditSurvey) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *AuditSurveysPutAuditSurvey) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *AuditSurveysPutAuditSurvey) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.


### GetReferenceMeta

`func (o *AuditSurveysPutAuditSurvey) GetReferenceMeta() interface{}`

GetReferenceMeta returns the ReferenceMeta field if non-nil, zero value otherwise.

### GetReferenceMetaOk

`func (o *AuditSurveysPutAuditSurvey) GetReferenceMetaOk() (*interface{}, bool)`

GetReferenceMetaOk returns a tuple with the ReferenceMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMeta

`func (o *AuditSurveysPutAuditSurvey) SetReferenceMeta(v interface{})`

SetReferenceMeta sets ReferenceMeta field to given value.

### HasReferenceMeta

`func (o *AuditSurveysPutAuditSurvey) HasReferenceMeta() bool`

HasReferenceMeta returns a boolean if a field has been set.

### SetReferenceMetaNil

`func (o *AuditSurveysPutAuditSurvey) SetReferenceMetaNil(b bool)`

 SetReferenceMetaNil sets the value for ReferenceMeta to be an explicit nil

### UnsetReferenceMeta
`func (o *AuditSurveysPutAuditSurvey) UnsetReferenceMeta()`

UnsetReferenceMeta ensures that no value is present for ReferenceMeta, not even an explicit nil
### GetTotals

`func (o *AuditSurveysPutAuditSurvey) GetTotals() interface{}`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *AuditSurveysPutAuditSurvey) GetTotalsOk() (*interface{}, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *AuditSurveysPutAuditSurvey) SetTotals(v interface{})`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *AuditSurveysPutAuditSurvey) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### SetTotalsNil

`func (o *AuditSurveysPutAuditSurvey) SetTotalsNil(b bool)`

 SetTotalsNil sets the value for Totals to be an explicit nil

### UnsetTotals
`func (o *AuditSurveysPutAuditSurvey) UnsetTotals()`

UnsetTotals ensures that no value is present for Totals, not even an explicit nil
### GetQuestionOrder

`func (o *AuditSurveysPutAuditSurvey) GetQuestionOrder() interface{}`

GetQuestionOrder returns the QuestionOrder field if non-nil, zero value otherwise.

### GetQuestionOrderOk

`func (o *AuditSurveysPutAuditSurvey) GetQuestionOrderOk() (*interface{}, bool)`

GetQuestionOrderOk returns a tuple with the QuestionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionOrder

`func (o *AuditSurveysPutAuditSurvey) SetQuestionOrder(v interface{})`

SetQuestionOrder sets QuestionOrder field to given value.


### SetQuestionOrderNil

`func (o *AuditSurveysPutAuditSurvey) SetQuestionOrderNil(b bool)`

 SetQuestionOrderNil sets the value for QuestionOrder to be an explicit nil

### UnsetQuestionOrder
`func (o *AuditSurveysPutAuditSurvey) UnsetQuestionOrder()`

UnsetQuestionOrder ensures that no value is present for QuestionOrder, not even an explicit nil
### GetScore

`func (o *AuditSurveysPutAuditSurvey) GetScore() interface{}`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *AuditSurveysPutAuditSurvey) GetScoreOk() (*interface{}, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *AuditSurveysPutAuditSurvey) SetScore(v interface{})`

SetScore sets Score field to given value.

### HasScore

`func (o *AuditSurveysPutAuditSurvey) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *AuditSurveysPutAuditSurvey) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *AuditSurveysPutAuditSurvey) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetType

`func (o *AuditSurveysPutAuditSurvey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditSurveysPutAuditSurvey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditSurveysPutAuditSurvey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuditSurveysPutAuditSurvey) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTaskItemId

`func (o *AuditSurveysPutAuditSurvey) GetTaskItemId() int32`

GetTaskItemId returns the TaskItemId field if non-nil, zero value otherwise.

### GetTaskItemIdOk

`func (o *AuditSurveysPutAuditSurvey) GetTaskItemIdOk() (*int32, bool)`

GetTaskItemIdOk returns a tuple with the TaskItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskItemId

`func (o *AuditSurveysPutAuditSurvey) SetTaskItemId(v int32)`

SetTaskItemId sets TaskItemId field to given value.

### HasTaskItemId

`func (o *AuditSurveysPutAuditSurvey) HasTaskItemId() bool`

HasTaskItemId returns a boolean if a field has been set.

### GetScoringMethodology

`func (o *AuditSurveysPutAuditSurvey) GetScoringMethodology() string`

GetScoringMethodology returns the ScoringMethodology field if non-nil, zero value otherwise.

### GetScoringMethodologyOk

`func (o *AuditSurveysPutAuditSurvey) GetScoringMethodologyOk() (*string, bool)`

GetScoringMethodologyOk returns a tuple with the ScoringMethodology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoringMethodology

`func (o *AuditSurveysPutAuditSurvey) SetScoringMethodology(v string)`

SetScoringMethodology sets ScoringMethodology field to given value.

### HasScoringMethodology

`func (o *AuditSurveysPutAuditSurvey) HasScoringMethodology() bool`

HasScoringMethodology returns a boolean if a field has been set.

### GetBracketsType

`func (o *AuditSurveysPutAuditSurvey) GetBracketsType() string`

GetBracketsType returns the BracketsType field if non-nil, zero value otherwise.

### GetBracketsTypeOk

`func (o *AuditSurveysPutAuditSurvey) GetBracketsTypeOk() (*string, bool)`

GetBracketsTypeOk returns a tuple with the BracketsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBracketsType

`func (o *AuditSurveysPutAuditSurvey) SetBracketsType(v string)`

SetBracketsType sets BracketsType field to given value.

### HasBracketsType

`func (o *AuditSurveysPutAuditSurvey) HasBracketsType() bool`

HasBracketsType returns a boolean if a field has been set.

### GetBrackets

`func (o *AuditSurveysPutAuditSurvey) GetBrackets() interface{}`

GetBrackets returns the Brackets field if non-nil, zero value otherwise.

### GetBracketsOk

`func (o *AuditSurveysPutAuditSurvey) GetBracketsOk() (*interface{}, bool)`

GetBracketsOk returns a tuple with the Brackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrackets

`func (o *AuditSurveysPutAuditSurvey) SetBrackets(v interface{})`

SetBrackets sets Brackets field to given value.


### SetBracketsNil

`func (o *AuditSurveysPutAuditSurvey) SetBracketsNil(b bool)`

 SetBracketsNil sets the value for Brackets to be an explicit nil

### UnsetBrackets
`func (o *AuditSurveysPutAuditSurvey) UnsetBrackets()`

UnsetBrackets ensures that no value is present for Brackets, not even an explicit nil
### GetAuditSurveyTemplateId

`func (o *AuditSurveysPutAuditSurvey) GetAuditSurveyTemplateId() int32`

GetAuditSurveyTemplateId returns the AuditSurveyTemplateId field if non-nil, zero value otherwise.

### GetAuditSurveyTemplateIdOk

`func (o *AuditSurveysPutAuditSurvey) GetAuditSurveyTemplateIdOk() (*int32, bool)`

GetAuditSurveyTemplateIdOk returns a tuple with the AuditSurveyTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSurveyTemplateId

`func (o *AuditSurveysPutAuditSurvey) SetAuditSurveyTemplateId(v int32)`

SetAuditSurveyTemplateId sets AuditSurveyTemplateId field to given value.

### HasAuditSurveyTemplateId

`func (o *AuditSurveysPutAuditSurvey) HasAuditSurveyTemplateId() bool`

HasAuditSurveyTemplateId returns a boolean if a field has been set.

### GetSectionOrder

`func (o *AuditSurveysPutAuditSurvey) GetSectionOrder() interface{}`

GetSectionOrder returns the SectionOrder field if non-nil, zero value otherwise.

### GetSectionOrderOk

`func (o *AuditSurveysPutAuditSurvey) GetSectionOrderOk() (*interface{}, bool)`

GetSectionOrderOk returns a tuple with the SectionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionOrder

`func (o *AuditSurveysPutAuditSurvey) SetSectionOrder(v interface{})`

SetSectionOrder sets SectionOrder field to given value.


### SetSectionOrderNil

`func (o *AuditSurveysPutAuditSurvey) SetSectionOrderNil(b bool)`

 SetSectionOrderNil sets the value for SectionOrder to be an explicit nil

### UnsetSectionOrder
`func (o *AuditSurveysPutAuditSurvey) UnsetSectionOrder()`

UnsetSectionOrder ensures that no value is present for SectionOrder, not even an explicit nil
### GetStashedQuestionOrder

`func (o *AuditSurveysPutAuditSurvey) GetStashedQuestionOrder() interface{}`

GetStashedQuestionOrder returns the StashedQuestionOrder field if non-nil, zero value otherwise.

### GetStashedQuestionOrderOk

`func (o *AuditSurveysPutAuditSurvey) GetStashedQuestionOrderOk() (*interface{}, bool)`

GetStashedQuestionOrderOk returns a tuple with the StashedQuestionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStashedQuestionOrder

`func (o *AuditSurveysPutAuditSurvey) SetStashedQuestionOrder(v interface{})`

SetStashedQuestionOrder sets StashedQuestionOrder field to given value.

### HasStashedQuestionOrder

`func (o *AuditSurveysPutAuditSurvey) HasStashedQuestionOrder() bool`

HasStashedQuestionOrder returns a boolean if a field has been set.

### SetStashedQuestionOrderNil

`func (o *AuditSurveysPutAuditSurvey) SetStashedQuestionOrderNil(b bool)`

 SetStashedQuestionOrderNil sets the value for StashedQuestionOrder to be an explicit nil

### UnsetStashedQuestionOrder
`func (o *AuditSurveysPutAuditSurvey) UnsetStashedQuestionOrder()`

UnsetStashedQuestionOrder ensures that no value is present for StashedQuestionOrder, not even an explicit nil
### GetRecurringSettings

`func (o *AuditSurveysPutAuditSurvey) GetRecurringSettings() interface{}`

GetRecurringSettings returns the RecurringSettings field if non-nil, zero value otherwise.

### GetRecurringSettingsOk

`func (o *AuditSurveysPutAuditSurvey) GetRecurringSettingsOk() (*interface{}, bool)`

GetRecurringSettingsOk returns a tuple with the RecurringSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringSettings

`func (o *AuditSurveysPutAuditSurvey) SetRecurringSettings(v interface{})`

SetRecurringSettings sets RecurringSettings field to given value.

### HasRecurringSettings

`func (o *AuditSurveysPutAuditSurvey) HasRecurringSettings() bool`

HasRecurringSettings returns a boolean if a field has been set.

### SetRecurringSettingsNil

`func (o *AuditSurveysPutAuditSurvey) SetRecurringSettingsNil(b bool)`

 SetRecurringSettingsNil sets the value for RecurringSettings to be an explicit nil

### UnsetRecurringSettings
`func (o *AuditSurveysPutAuditSurvey) UnsetRecurringSettings()`

UnsetRecurringSettings ensures that no value is present for RecurringSettings, not even an explicit nil
### GetIsWeighted

`func (o *AuditSurveysPutAuditSurvey) GetIsWeighted() bool`

GetIsWeighted returns the IsWeighted field if non-nil, zero value otherwise.

### GetIsWeightedOk

`func (o *AuditSurveysPutAuditSurvey) GetIsWeightedOk() (*bool, bool)`

GetIsWeightedOk returns a tuple with the IsWeighted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWeighted

`func (o *AuditSurveysPutAuditSurvey) SetIsWeighted(v bool)`

SetIsWeighted sets IsWeighted field to given value.

### HasIsWeighted

`func (o *AuditSurveysPutAuditSurvey) HasIsWeighted() bool`

HasIsWeighted returns a boolean if a field has been set.

### GetAuditSurveyCategoryId

`func (o *AuditSurveysPutAuditSurvey) GetAuditSurveyCategoryId() int32`

GetAuditSurveyCategoryId returns the AuditSurveyCategoryId field if non-nil, zero value otherwise.

### GetAuditSurveyCategoryIdOk

`func (o *AuditSurveysPutAuditSurvey) GetAuditSurveyCategoryIdOk() (*int32, bool)`

GetAuditSurveyCategoryIdOk returns a tuple with the AuditSurveyCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSurveyCategoryId

`func (o *AuditSurveysPutAuditSurvey) SetAuditSurveyCategoryId(v int32)`

SetAuditSurveyCategoryId sets AuditSurveyCategoryId field to given value.

### HasAuditSurveyCategoryId

`func (o *AuditSurveysPutAuditSurvey) HasAuditSurveyCategoryId() bool`

HasAuditSurveyCategoryId returns a boolean if a field has been set.

### GetAuditSurveyScoreId

`func (o *AuditSurveysPutAuditSurvey) GetAuditSurveyScoreId() int32`

GetAuditSurveyScoreId returns the AuditSurveyScoreId field if non-nil, zero value otherwise.

### GetAuditSurveyScoreIdOk

`func (o *AuditSurveysPutAuditSurvey) GetAuditSurveyScoreIdOk() (*int32, bool)`

GetAuditSurveyScoreIdOk returns a tuple with the AuditSurveyScoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSurveyScoreId

`func (o *AuditSurveysPutAuditSurvey) SetAuditSurveyScoreId(v int32)`

SetAuditSurveyScoreId sets AuditSurveyScoreId field to given value.

### HasAuditSurveyScoreId

`func (o *AuditSurveysPutAuditSurvey) HasAuditSurveyScoreId() bool`

HasAuditSurveyScoreId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


