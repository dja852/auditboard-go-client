# AuditSurveysPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**AuditsurveyableId** | Pointer to **int32** |  | [optional] 
**AuditsurveyableType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**QuestionSortOrder** | Pointer to **interface{}** |  | [optional] 
**SourceId** | Pointer to **int32** |  | [optional] 
**IsTemplate** | Pointer to **bool** |  | [optional] [default to false]
**ReferenceMeta** | Pointer to **interface{}** |  | [optional] 
**Totals** | Pointer to **interface{}** |  | [optional] 
**QuestionOrder** | Pointer to **interface{}** |  | [optional] 
**Score** | Pointer to **interface{}** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**TaskItemId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;task_items.id&#x60;.&lt;fk table&#x3D;&#39;task_items&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ScoringMethodology** | Pointer to **string** |  | [optional] 
**BracketsType** | Pointer to **string** |  | [optional] [default to "NULL"]
**Brackets** | Pointer to **interface{}** |  | [optional] 
**AuditSurveyTemplateId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_survey_templates.id&#x60;.&lt;fk table&#x3D;&#39;audit_survey_templates&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**SectionOrder** | Pointer to **interface{}** |  | [optional] 
**StashedQuestionOrder** | Pointer to **interface{}** |  | [optional] 
**RecurringSettings** | Pointer to **interface{}** |  | [optional] 
**IsWeighted** | Pointer to **bool** |  | [optional] [default to false]
**AuditSurveyCategoryId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_survey_categories.id&#x60;.&lt;fk table&#x3D;&#39;audit_survey_categories&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AuditSurveyScoreId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_survey_scores.id&#x60;.&lt;fk table&#x3D;&#39;audit_survey_scores&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 

## Methods

### NewAuditSurveysPutPreviousValues

`func NewAuditSurveysPutPreviousValues() *AuditSurveysPutPreviousValues`

NewAuditSurveysPutPreviousValues instantiates a new AuditSurveysPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditSurveysPutPreviousValuesWithDefaults

`func NewAuditSurveysPutPreviousValuesWithDefaults() *AuditSurveysPutPreviousValues`

NewAuditSurveysPutPreviousValuesWithDefaults instantiates a new AuditSurveysPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditSurveysPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditSurveysPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditSurveysPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuditSurveysPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuditSurveysPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditSurveysPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditSurveysPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuditSurveysPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuditSurveysPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuditSurveysPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuditSurveysPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuditSurveysPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AuditSurveysPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AuditSurveysPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AuditSurveysPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AuditSurveysPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetAuditsurveyableId

`func (o *AuditSurveysPutPreviousValues) GetAuditsurveyableId() int32`

GetAuditsurveyableId returns the AuditsurveyableId field if non-nil, zero value otherwise.

### GetAuditsurveyableIdOk

`func (o *AuditSurveysPutPreviousValues) GetAuditsurveyableIdOk() (*int32, bool)`

GetAuditsurveyableIdOk returns a tuple with the AuditsurveyableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditsurveyableId

`func (o *AuditSurveysPutPreviousValues) SetAuditsurveyableId(v int32)`

SetAuditsurveyableId sets AuditsurveyableId field to given value.

### HasAuditsurveyableId

`func (o *AuditSurveysPutPreviousValues) HasAuditsurveyableId() bool`

HasAuditsurveyableId returns a boolean if a field has been set.

### GetAuditsurveyableType

`func (o *AuditSurveysPutPreviousValues) GetAuditsurveyableType() string`

GetAuditsurveyableType returns the AuditsurveyableType field if non-nil, zero value otherwise.

### GetAuditsurveyableTypeOk

`func (o *AuditSurveysPutPreviousValues) GetAuditsurveyableTypeOk() (*string, bool)`

GetAuditsurveyableTypeOk returns a tuple with the AuditsurveyableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditsurveyableType

`func (o *AuditSurveysPutPreviousValues) SetAuditsurveyableType(v string)`

SetAuditsurveyableType sets AuditsurveyableType field to given value.

### HasAuditsurveyableType

`func (o *AuditSurveysPutPreviousValues) HasAuditsurveyableType() bool`

HasAuditsurveyableType returns a boolean if a field has been set.

### GetName

`func (o *AuditSurveysPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditSurveysPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditSurveysPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuditSurveysPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *AuditSurveysPutPreviousValues) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *AuditSurveysPutPreviousValues) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *AuditSurveysPutPreviousValues) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *AuditSurveysPutPreviousValues) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetQuestionSortOrder

`func (o *AuditSurveysPutPreviousValues) GetQuestionSortOrder() interface{}`

GetQuestionSortOrder returns the QuestionSortOrder field if non-nil, zero value otherwise.

### GetQuestionSortOrderOk

`func (o *AuditSurveysPutPreviousValues) GetQuestionSortOrderOk() (*interface{}, bool)`

GetQuestionSortOrderOk returns a tuple with the QuestionSortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionSortOrder

`func (o *AuditSurveysPutPreviousValues) SetQuestionSortOrder(v interface{})`

SetQuestionSortOrder sets QuestionSortOrder field to given value.

### HasQuestionSortOrder

`func (o *AuditSurveysPutPreviousValues) HasQuestionSortOrder() bool`

HasQuestionSortOrder returns a boolean if a field has been set.

### SetQuestionSortOrderNil

`func (o *AuditSurveysPutPreviousValues) SetQuestionSortOrderNil(b bool)`

 SetQuestionSortOrderNil sets the value for QuestionSortOrder to be an explicit nil

### UnsetQuestionSortOrder
`func (o *AuditSurveysPutPreviousValues) UnsetQuestionSortOrder()`

UnsetQuestionSortOrder ensures that no value is present for QuestionSortOrder, not even an explicit nil
### GetSourceId

`func (o *AuditSurveysPutPreviousValues) GetSourceId() int32`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AuditSurveysPutPreviousValues) GetSourceIdOk() (*int32, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AuditSurveysPutPreviousValues) SetSourceId(v int32)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AuditSurveysPutPreviousValues) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetIsTemplate

`func (o *AuditSurveysPutPreviousValues) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *AuditSurveysPutPreviousValues) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *AuditSurveysPutPreviousValues) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *AuditSurveysPutPreviousValues) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.

### GetReferenceMeta

`func (o *AuditSurveysPutPreviousValues) GetReferenceMeta() interface{}`

GetReferenceMeta returns the ReferenceMeta field if non-nil, zero value otherwise.

### GetReferenceMetaOk

`func (o *AuditSurveysPutPreviousValues) GetReferenceMetaOk() (*interface{}, bool)`

GetReferenceMetaOk returns a tuple with the ReferenceMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMeta

`func (o *AuditSurveysPutPreviousValues) SetReferenceMeta(v interface{})`

SetReferenceMeta sets ReferenceMeta field to given value.

### HasReferenceMeta

`func (o *AuditSurveysPutPreviousValues) HasReferenceMeta() bool`

HasReferenceMeta returns a boolean if a field has been set.

### SetReferenceMetaNil

`func (o *AuditSurveysPutPreviousValues) SetReferenceMetaNil(b bool)`

 SetReferenceMetaNil sets the value for ReferenceMeta to be an explicit nil

### UnsetReferenceMeta
`func (o *AuditSurveysPutPreviousValues) UnsetReferenceMeta()`

UnsetReferenceMeta ensures that no value is present for ReferenceMeta, not even an explicit nil
### GetTotals

`func (o *AuditSurveysPutPreviousValues) GetTotals() interface{}`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *AuditSurveysPutPreviousValues) GetTotalsOk() (*interface{}, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *AuditSurveysPutPreviousValues) SetTotals(v interface{})`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *AuditSurveysPutPreviousValues) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### SetTotalsNil

`func (o *AuditSurveysPutPreviousValues) SetTotalsNil(b bool)`

 SetTotalsNil sets the value for Totals to be an explicit nil

### UnsetTotals
`func (o *AuditSurveysPutPreviousValues) UnsetTotals()`

UnsetTotals ensures that no value is present for Totals, not even an explicit nil
### GetQuestionOrder

`func (o *AuditSurveysPutPreviousValues) GetQuestionOrder() interface{}`

GetQuestionOrder returns the QuestionOrder field if non-nil, zero value otherwise.

### GetQuestionOrderOk

`func (o *AuditSurveysPutPreviousValues) GetQuestionOrderOk() (*interface{}, bool)`

GetQuestionOrderOk returns a tuple with the QuestionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionOrder

`func (o *AuditSurveysPutPreviousValues) SetQuestionOrder(v interface{})`

SetQuestionOrder sets QuestionOrder field to given value.

### HasQuestionOrder

`func (o *AuditSurveysPutPreviousValues) HasQuestionOrder() bool`

HasQuestionOrder returns a boolean if a field has been set.

### SetQuestionOrderNil

`func (o *AuditSurveysPutPreviousValues) SetQuestionOrderNil(b bool)`

 SetQuestionOrderNil sets the value for QuestionOrder to be an explicit nil

### UnsetQuestionOrder
`func (o *AuditSurveysPutPreviousValues) UnsetQuestionOrder()`

UnsetQuestionOrder ensures that no value is present for QuestionOrder, not even an explicit nil
### GetScore

`func (o *AuditSurveysPutPreviousValues) GetScore() interface{}`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *AuditSurveysPutPreviousValues) GetScoreOk() (*interface{}, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *AuditSurveysPutPreviousValues) SetScore(v interface{})`

SetScore sets Score field to given value.

### HasScore

`func (o *AuditSurveysPutPreviousValues) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *AuditSurveysPutPreviousValues) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *AuditSurveysPutPreviousValues) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetType

`func (o *AuditSurveysPutPreviousValues) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditSurveysPutPreviousValues) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditSurveysPutPreviousValues) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuditSurveysPutPreviousValues) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTaskItemId

`func (o *AuditSurveysPutPreviousValues) GetTaskItemId() int32`

GetTaskItemId returns the TaskItemId field if non-nil, zero value otherwise.

### GetTaskItemIdOk

`func (o *AuditSurveysPutPreviousValues) GetTaskItemIdOk() (*int32, bool)`

GetTaskItemIdOk returns a tuple with the TaskItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskItemId

`func (o *AuditSurveysPutPreviousValues) SetTaskItemId(v int32)`

SetTaskItemId sets TaskItemId field to given value.

### HasTaskItemId

`func (o *AuditSurveysPutPreviousValues) HasTaskItemId() bool`

HasTaskItemId returns a boolean if a field has been set.

### GetScoringMethodology

`func (o *AuditSurveysPutPreviousValues) GetScoringMethodology() string`

GetScoringMethodology returns the ScoringMethodology field if non-nil, zero value otherwise.

### GetScoringMethodologyOk

`func (o *AuditSurveysPutPreviousValues) GetScoringMethodologyOk() (*string, bool)`

GetScoringMethodologyOk returns a tuple with the ScoringMethodology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoringMethodology

`func (o *AuditSurveysPutPreviousValues) SetScoringMethodology(v string)`

SetScoringMethodology sets ScoringMethodology field to given value.

### HasScoringMethodology

`func (o *AuditSurveysPutPreviousValues) HasScoringMethodology() bool`

HasScoringMethodology returns a boolean if a field has been set.

### GetBracketsType

`func (o *AuditSurveysPutPreviousValues) GetBracketsType() string`

GetBracketsType returns the BracketsType field if non-nil, zero value otherwise.

### GetBracketsTypeOk

`func (o *AuditSurveysPutPreviousValues) GetBracketsTypeOk() (*string, bool)`

GetBracketsTypeOk returns a tuple with the BracketsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBracketsType

`func (o *AuditSurveysPutPreviousValues) SetBracketsType(v string)`

SetBracketsType sets BracketsType field to given value.

### HasBracketsType

`func (o *AuditSurveysPutPreviousValues) HasBracketsType() bool`

HasBracketsType returns a boolean if a field has been set.

### GetBrackets

`func (o *AuditSurveysPutPreviousValues) GetBrackets() interface{}`

GetBrackets returns the Brackets field if non-nil, zero value otherwise.

### GetBracketsOk

`func (o *AuditSurveysPutPreviousValues) GetBracketsOk() (*interface{}, bool)`

GetBracketsOk returns a tuple with the Brackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrackets

`func (o *AuditSurveysPutPreviousValues) SetBrackets(v interface{})`

SetBrackets sets Brackets field to given value.

### HasBrackets

`func (o *AuditSurveysPutPreviousValues) HasBrackets() bool`

HasBrackets returns a boolean if a field has been set.

### SetBracketsNil

`func (o *AuditSurveysPutPreviousValues) SetBracketsNil(b bool)`

 SetBracketsNil sets the value for Brackets to be an explicit nil

### UnsetBrackets
`func (o *AuditSurveysPutPreviousValues) UnsetBrackets()`

UnsetBrackets ensures that no value is present for Brackets, not even an explicit nil
### GetAuditSurveyTemplateId

`func (o *AuditSurveysPutPreviousValues) GetAuditSurveyTemplateId() int32`

GetAuditSurveyTemplateId returns the AuditSurveyTemplateId field if non-nil, zero value otherwise.

### GetAuditSurveyTemplateIdOk

`func (o *AuditSurveysPutPreviousValues) GetAuditSurveyTemplateIdOk() (*int32, bool)`

GetAuditSurveyTemplateIdOk returns a tuple with the AuditSurveyTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSurveyTemplateId

`func (o *AuditSurveysPutPreviousValues) SetAuditSurveyTemplateId(v int32)`

SetAuditSurveyTemplateId sets AuditSurveyTemplateId field to given value.

### HasAuditSurveyTemplateId

`func (o *AuditSurveysPutPreviousValues) HasAuditSurveyTemplateId() bool`

HasAuditSurveyTemplateId returns a boolean if a field has been set.

### GetSectionOrder

`func (o *AuditSurveysPutPreviousValues) GetSectionOrder() interface{}`

GetSectionOrder returns the SectionOrder field if non-nil, zero value otherwise.

### GetSectionOrderOk

`func (o *AuditSurveysPutPreviousValues) GetSectionOrderOk() (*interface{}, bool)`

GetSectionOrderOk returns a tuple with the SectionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionOrder

`func (o *AuditSurveysPutPreviousValues) SetSectionOrder(v interface{})`

SetSectionOrder sets SectionOrder field to given value.

### HasSectionOrder

`func (o *AuditSurveysPutPreviousValues) HasSectionOrder() bool`

HasSectionOrder returns a boolean if a field has been set.

### SetSectionOrderNil

`func (o *AuditSurveysPutPreviousValues) SetSectionOrderNil(b bool)`

 SetSectionOrderNil sets the value for SectionOrder to be an explicit nil

### UnsetSectionOrder
`func (o *AuditSurveysPutPreviousValues) UnsetSectionOrder()`

UnsetSectionOrder ensures that no value is present for SectionOrder, not even an explicit nil
### GetStashedQuestionOrder

`func (o *AuditSurveysPutPreviousValues) GetStashedQuestionOrder() interface{}`

GetStashedQuestionOrder returns the StashedQuestionOrder field if non-nil, zero value otherwise.

### GetStashedQuestionOrderOk

`func (o *AuditSurveysPutPreviousValues) GetStashedQuestionOrderOk() (*interface{}, bool)`

GetStashedQuestionOrderOk returns a tuple with the StashedQuestionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStashedQuestionOrder

`func (o *AuditSurveysPutPreviousValues) SetStashedQuestionOrder(v interface{})`

SetStashedQuestionOrder sets StashedQuestionOrder field to given value.

### HasStashedQuestionOrder

`func (o *AuditSurveysPutPreviousValues) HasStashedQuestionOrder() bool`

HasStashedQuestionOrder returns a boolean if a field has been set.

### SetStashedQuestionOrderNil

`func (o *AuditSurveysPutPreviousValues) SetStashedQuestionOrderNil(b bool)`

 SetStashedQuestionOrderNil sets the value for StashedQuestionOrder to be an explicit nil

### UnsetStashedQuestionOrder
`func (o *AuditSurveysPutPreviousValues) UnsetStashedQuestionOrder()`

UnsetStashedQuestionOrder ensures that no value is present for StashedQuestionOrder, not even an explicit nil
### GetRecurringSettings

`func (o *AuditSurveysPutPreviousValues) GetRecurringSettings() interface{}`

GetRecurringSettings returns the RecurringSettings field if non-nil, zero value otherwise.

### GetRecurringSettingsOk

`func (o *AuditSurveysPutPreviousValues) GetRecurringSettingsOk() (*interface{}, bool)`

GetRecurringSettingsOk returns a tuple with the RecurringSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringSettings

`func (o *AuditSurveysPutPreviousValues) SetRecurringSettings(v interface{})`

SetRecurringSettings sets RecurringSettings field to given value.

### HasRecurringSettings

`func (o *AuditSurveysPutPreviousValues) HasRecurringSettings() bool`

HasRecurringSettings returns a boolean if a field has been set.

### SetRecurringSettingsNil

`func (o *AuditSurveysPutPreviousValues) SetRecurringSettingsNil(b bool)`

 SetRecurringSettingsNil sets the value for RecurringSettings to be an explicit nil

### UnsetRecurringSettings
`func (o *AuditSurveysPutPreviousValues) UnsetRecurringSettings()`

UnsetRecurringSettings ensures that no value is present for RecurringSettings, not even an explicit nil
### GetIsWeighted

`func (o *AuditSurveysPutPreviousValues) GetIsWeighted() bool`

GetIsWeighted returns the IsWeighted field if non-nil, zero value otherwise.

### GetIsWeightedOk

`func (o *AuditSurveysPutPreviousValues) GetIsWeightedOk() (*bool, bool)`

GetIsWeightedOk returns a tuple with the IsWeighted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWeighted

`func (o *AuditSurveysPutPreviousValues) SetIsWeighted(v bool)`

SetIsWeighted sets IsWeighted field to given value.

### HasIsWeighted

`func (o *AuditSurveysPutPreviousValues) HasIsWeighted() bool`

HasIsWeighted returns a boolean if a field has been set.

### GetAuditSurveyCategoryId

`func (o *AuditSurveysPutPreviousValues) GetAuditSurveyCategoryId() int32`

GetAuditSurveyCategoryId returns the AuditSurveyCategoryId field if non-nil, zero value otherwise.

### GetAuditSurveyCategoryIdOk

`func (o *AuditSurveysPutPreviousValues) GetAuditSurveyCategoryIdOk() (*int32, bool)`

GetAuditSurveyCategoryIdOk returns a tuple with the AuditSurveyCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSurveyCategoryId

`func (o *AuditSurveysPutPreviousValues) SetAuditSurveyCategoryId(v int32)`

SetAuditSurveyCategoryId sets AuditSurveyCategoryId field to given value.

### HasAuditSurveyCategoryId

`func (o *AuditSurveysPutPreviousValues) HasAuditSurveyCategoryId() bool`

HasAuditSurveyCategoryId returns a boolean if a field has been set.

### GetAuditSurveyScoreId

`func (o *AuditSurveysPutPreviousValues) GetAuditSurveyScoreId() int32`

GetAuditSurveyScoreId returns the AuditSurveyScoreId field if non-nil, zero value otherwise.

### GetAuditSurveyScoreIdOk

`func (o *AuditSurveysPutPreviousValues) GetAuditSurveyScoreIdOk() (*int32, bool)`

GetAuditSurveyScoreIdOk returns a tuple with the AuditSurveyScoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSurveyScoreId

`func (o *AuditSurveysPutPreviousValues) SetAuditSurveyScoreId(v int32)`

SetAuditSurveyScoreId sets AuditSurveyScoreId field to given value.

### HasAuditSurveyScoreId

`func (o *AuditSurveysPutPreviousValues) HasAuditSurveyScoreId() bool`

HasAuditSurveyScoreId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


