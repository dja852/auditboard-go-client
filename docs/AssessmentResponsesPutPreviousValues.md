# AssessmentResponsesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] 
**ResponseJson** | Pointer to **interface{}** |  | [optional] 
**AssessableType** | Pointer to **string** |  | [optional] 
**AssessableId** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**SubmittedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**SubmittedDate** | Pointer to **string** |  | [optional] 
**ReferenceMeta** | Pointer to **interface{}** |  | [optional] 
**AssessmentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;assessments.id&#x60;.&lt;fk table&#x3D;&#39;assessments&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**AssessmentTemplateId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;assessment_templates.id&#x60;.&lt;fk table&#x3D;&#39;assessment_templates&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FinalizedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FinalizedDate** | Pointer to **string** |  | [optional] 
**AnsweredByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AnsweredDate** | Pointer to **string** |  | [optional] 
**IsAnswered** | Pointer to **bool** |  | [optional] [default to false]
**AssigneeUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IsNotApplicable** | Pointer to **bool** |  | [optional] [default to false]
**Notes** | Pointer to **string** |  | [optional] 
**UserAssessmentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;user_assessments.id&#x60;.&lt;fk table&#x3D;&#39;user_assessments&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskResponseId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_responses.id&#x60;.&lt;fk table&#x3D;&#39;risk_responses&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskresponseactionId** | Pointer to **int32** |  | [optional] 
**RiskresponseactionType** | Pointer to **string** |  | [optional] 
**ParentAssessmentResponseId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;assessment_responses.id&#x60;.&lt;fk table&#x3D;&#39;assessment_responses&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 

## Methods

### NewAssessmentResponsesPutPreviousValues

`func NewAssessmentResponsesPutPreviousValues() *AssessmentResponsesPutPreviousValues`

NewAssessmentResponsesPutPreviousValues instantiates a new AssessmentResponsesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssessmentResponsesPutPreviousValuesWithDefaults

`func NewAssessmentResponsesPutPreviousValuesWithDefaults() *AssessmentResponsesPutPreviousValues`

NewAssessmentResponsesPutPreviousValuesWithDefaults instantiates a new AssessmentResponsesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssessmentResponsesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssessmentResponsesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssessmentResponsesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AssessmentResponsesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *AssessmentResponsesPutPreviousValues) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssessmentResponsesPutPreviousValues) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssessmentResponsesPutPreviousValues) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssessmentResponsesPutPreviousValues) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResponseJson

`func (o *AssessmentResponsesPutPreviousValues) GetResponseJson() interface{}`

GetResponseJson returns the ResponseJson field if non-nil, zero value otherwise.

### GetResponseJsonOk

`func (o *AssessmentResponsesPutPreviousValues) GetResponseJsonOk() (*interface{}, bool)`

GetResponseJsonOk returns a tuple with the ResponseJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseJson

`func (o *AssessmentResponsesPutPreviousValues) SetResponseJson(v interface{})`

SetResponseJson sets ResponseJson field to given value.

### HasResponseJson

`func (o *AssessmentResponsesPutPreviousValues) HasResponseJson() bool`

HasResponseJson returns a boolean if a field has been set.

### SetResponseJsonNil

`func (o *AssessmentResponsesPutPreviousValues) SetResponseJsonNil(b bool)`

 SetResponseJsonNil sets the value for ResponseJson to be an explicit nil

### UnsetResponseJson
`func (o *AssessmentResponsesPutPreviousValues) UnsetResponseJson()`

UnsetResponseJson ensures that no value is present for ResponseJson, not even an explicit nil
### GetAssessableType

`func (o *AssessmentResponsesPutPreviousValues) GetAssessableType() string`

GetAssessableType returns the AssessableType field if non-nil, zero value otherwise.

### GetAssessableTypeOk

`func (o *AssessmentResponsesPutPreviousValues) GetAssessableTypeOk() (*string, bool)`

GetAssessableTypeOk returns a tuple with the AssessableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessableType

`func (o *AssessmentResponsesPutPreviousValues) SetAssessableType(v string)`

SetAssessableType sets AssessableType field to given value.

### HasAssessableType

`func (o *AssessmentResponsesPutPreviousValues) HasAssessableType() bool`

HasAssessableType returns a boolean if a field has been set.

### GetAssessableId

`func (o *AssessmentResponsesPutPreviousValues) GetAssessableId() int32`

GetAssessableId returns the AssessableId field if non-nil, zero value otherwise.

### GetAssessableIdOk

`func (o *AssessmentResponsesPutPreviousValues) GetAssessableIdOk() (*int32, bool)`

GetAssessableIdOk returns a tuple with the AssessableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessableId

`func (o *AssessmentResponsesPutPreviousValues) SetAssessableId(v int32)`

SetAssessableId sets AssessableId field to given value.

### HasAssessableId

`func (o *AssessmentResponsesPutPreviousValues) HasAssessableId() bool`

HasAssessableId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AssessmentResponsesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AssessmentResponsesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AssessmentResponsesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AssessmentResponsesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AssessmentResponsesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AssessmentResponsesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AssessmentResponsesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AssessmentResponsesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AssessmentResponsesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AssessmentResponsesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AssessmentResponsesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AssessmentResponsesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetSubmittedByUserId

`func (o *AssessmentResponsesPutPreviousValues) GetSubmittedByUserId() int32`

GetSubmittedByUserId returns the SubmittedByUserId field if non-nil, zero value otherwise.

### GetSubmittedByUserIdOk

`func (o *AssessmentResponsesPutPreviousValues) GetSubmittedByUserIdOk() (*int32, bool)`

GetSubmittedByUserIdOk returns a tuple with the SubmittedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedByUserId

`func (o *AssessmentResponsesPutPreviousValues) SetSubmittedByUserId(v int32)`

SetSubmittedByUserId sets SubmittedByUserId field to given value.

### HasSubmittedByUserId

`func (o *AssessmentResponsesPutPreviousValues) HasSubmittedByUserId() bool`

HasSubmittedByUserId returns a boolean if a field has been set.

### GetSubmittedDate

`func (o *AssessmentResponsesPutPreviousValues) GetSubmittedDate() string`

GetSubmittedDate returns the SubmittedDate field if non-nil, zero value otherwise.

### GetSubmittedDateOk

`func (o *AssessmentResponsesPutPreviousValues) GetSubmittedDateOk() (*string, bool)`

GetSubmittedDateOk returns a tuple with the SubmittedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedDate

`func (o *AssessmentResponsesPutPreviousValues) SetSubmittedDate(v string)`

SetSubmittedDate sets SubmittedDate field to given value.

### HasSubmittedDate

`func (o *AssessmentResponsesPutPreviousValues) HasSubmittedDate() bool`

HasSubmittedDate returns a boolean if a field has been set.

### GetReferenceMeta

`func (o *AssessmentResponsesPutPreviousValues) GetReferenceMeta() interface{}`

GetReferenceMeta returns the ReferenceMeta field if non-nil, zero value otherwise.

### GetReferenceMetaOk

`func (o *AssessmentResponsesPutPreviousValues) GetReferenceMetaOk() (*interface{}, bool)`

GetReferenceMetaOk returns a tuple with the ReferenceMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMeta

`func (o *AssessmentResponsesPutPreviousValues) SetReferenceMeta(v interface{})`

SetReferenceMeta sets ReferenceMeta field to given value.

### HasReferenceMeta

`func (o *AssessmentResponsesPutPreviousValues) HasReferenceMeta() bool`

HasReferenceMeta returns a boolean if a field has been set.

### SetReferenceMetaNil

`func (o *AssessmentResponsesPutPreviousValues) SetReferenceMetaNil(b bool)`

 SetReferenceMetaNil sets the value for ReferenceMeta to be an explicit nil

### UnsetReferenceMeta
`func (o *AssessmentResponsesPutPreviousValues) UnsetReferenceMeta()`

UnsetReferenceMeta ensures that no value is present for ReferenceMeta, not even an explicit nil
### GetAssessmentId

`func (o *AssessmentResponsesPutPreviousValues) GetAssessmentId() int32`

GetAssessmentId returns the AssessmentId field if non-nil, zero value otherwise.

### GetAssessmentIdOk

`func (o *AssessmentResponsesPutPreviousValues) GetAssessmentIdOk() (*int32, bool)`

GetAssessmentIdOk returns a tuple with the AssessmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessmentId

`func (o *AssessmentResponsesPutPreviousValues) SetAssessmentId(v int32)`

SetAssessmentId sets AssessmentId field to given value.

### HasAssessmentId

`func (o *AssessmentResponsesPutPreviousValues) HasAssessmentId() bool`

HasAssessmentId returns a boolean if a field has been set.

### GetSource

`func (o *AssessmentResponsesPutPreviousValues) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AssessmentResponsesPutPreviousValues) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AssessmentResponsesPutPreviousValues) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AssessmentResponsesPutPreviousValues) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetAssessmentTemplateId

`func (o *AssessmentResponsesPutPreviousValues) GetAssessmentTemplateId() int32`

GetAssessmentTemplateId returns the AssessmentTemplateId field if non-nil, zero value otherwise.

### GetAssessmentTemplateIdOk

`func (o *AssessmentResponsesPutPreviousValues) GetAssessmentTemplateIdOk() (*int32, bool)`

GetAssessmentTemplateIdOk returns a tuple with the AssessmentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessmentTemplateId

`func (o *AssessmentResponsesPutPreviousValues) SetAssessmentTemplateId(v int32)`

SetAssessmentTemplateId sets AssessmentTemplateId field to given value.

### HasAssessmentTemplateId

`func (o *AssessmentResponsesPutPreviousValues) HasAssessmentTemplateId() bool`

HasAssessmentTemplateId returns a boolean if a field has been set.

### GetFinalizedByUserId

`func (o *AssessmentResponsesPutPreviousValues) GetFinalizedByUserId() int32`

GetFinalizedByUserId returns the FinalizedByUserId field if non-nil, zero value otherwise.

### GetFinalizedByUserIdOk

`func (o *AssessmentResponsesPutPreviousValues) GetFinalizedByUserIdOk() (*int32, bool)`

GetFinalizedByUserIdOk returns a tuple with the FinalizedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizedByUserId

`func (o *AssessmentResponsesPutPreviousValues) SetFinalizedByUserId(v int32)`

SetFinalizedByUserId sets FinalizedByUserId field to given value.

### HasFinalizedByUserId

`func (o *AssessmentResponsesPutPreviousValues) HasFinalizedByUserId() bool`

HasFinalizedByUserId returns a boolean if a field has been set.

### GetFinalizedDate

`func (o *AssessmentResponsesPutPreviousValues) GetFinalizedDate() string`

GetFinalizedDate returns the FinalizedDate field if non-nil, zero value otherwise.

### GetFinalizedDateOk

`func (o *AssessmentResponsesPutPreviousValues) GetFinalizedDateOk() (*string, bool)`

GetFinalizedDateOk returns a tuple with the FinalizedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizedDate

`func (o *AssessmentResponsesPutPreviousValues) SetFinalizedDate(v string)`

SetFinalizedDate sets FinalizedDate field to given value.

### HasFinalizedDate

`func (o *AssessmentResponsesPutPreviousValues) HasFinalizedDate() bool`

HasFinalizedDate returns a boolean if a field has been set.

### GetAnsweredByUserId

`func (o *AssessmentResponsesPutPreviousValues) GetAnsweredByUserId() int32`

GetAnsweredByUserId returns the AnsweredByUserId field if non-nil, zero value otherwise.

### GetAnsweredByUserIdOk

`func (o *AssessmentResponsesPutPreviousValues) GetAnsweredByUserIdOk() (*int32, bool)`

GetAnsweredByUserIdOk returns a tuple with the AnsweredByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsweredByUserId

`func (o *AssessmentResponsesPutPreviousValues) SetAnsweredByUserId(v int32)`

SetAnsweredByUserId sets AnsweredByUserId field to given value.

### HasAnsweredByUserId

`func (o *AssessmentResponsesPutPreviousValues) HasAnsweredByUserId() bool`

HasAnsweredByUserId returns a boolean if a field has been set.

### GetAnsweredDate

`func (o *AssessmentResponsesPutPreviousValues) GetAnsweredDate() string`

GetAnsweredDate returns the AnsweredDate field if non-nil, zero value otherwise.

### GetAnsweredDateOk

`func (o *AssessmentResponsesPutPreviousValues) GetAnsweredDateOk() (*string, bool)`

GetAnsweredDateOk returns a tuple with the AnsweredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsweredDate

`func (o *AssessmentResponsesPutPreviousValues) SetAnsweredDate(v string)`

SetAnsweredDate sets AnsweredDate field to given value.

### HasAnsweredDate

`func (o *AssessmentResponsesPutPreviousValues) HasAnsweredDate() bool`

HasAnsweredDate returns a boolean if a field has been set.

### GetIsAnswered

`func (o *AssessmentResponsesPutPreviousValues) GetIsAnswered() bool`

GetIsAnswered returns the IsAnswered field if non-nil, zero value otherwise.

### GetIsAnsweredOk

`func (o *AssessmentResponsesPutPreviousValues) GetIsAnsweredOk() (*bool, bool)`

GetIsAnsweredOk returns a tuple with the IsAnswered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnswered

`func (o *AssessmentResponsesPutPreviousValues) SetIsAnswered(v bool)`

SetIsAnswered sets IsAnswered field to given value.

### HasIsAnswered

`func (o *AssessmentResponsesPutPreviousValues) HasIsAnswered() bool`

HasIsAnswered returns a boolean if a field has been set.

### GetAssigneeUserId

`func (o *AssessmentResponsesPutPreviousValues) GetAssigneeUserId() int32`

GetAssigneeUserId returns the AssigneeUserId field if non-nil, zero value otherwise.

### GetAssigneeUserIdOk

`func (o *AssessmentResponsesPutPreviousValues) GetAssigneeUserIdOk() (*int32, bool)`

GetAssigneeUserIdOk returns a tuple with the AssigneeUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeUserId

`func (o *AssessmentResponsesPutPreviousValues) SetAssigneeUserId(v int32)`

SetAssigneeUserId sets AssigneeUserId field to given value.

### HasAssigneeUserId

`func (o *AssessmentResponsesPutPreviousValues) HasAssigneeUserId() bool`

HasAssigneeUserId returns a boolean if a field has been set.

### GetIsNotApplicable

`func (o *AssessmentResponsesPutPreviousValues) GetIsNotApplicable() bool`

GetIsNotApplicable returns the IsNotApplicable field if non-nil, zero value otherwise.

### GetIsNotApplicableOk

`func (o *AssessmentResponsesPutPreviousValues) GetIsNotApplicableOk() (*bool, bool)`

GetIsNotApplicableOk returns a tuple with the IsNotApplicable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotApplicable

`func (o *AssessmentResponsesPutPreviousValues) SetIsNotApplicable(v bool)`

SetIsNotApplicable sets IsNotApplicable field to given value.

### HasIsNotApplicable

`func (o *AssessmentResponsesPutPreviousValues) HasIsNotApplicable() bool`

HasIsNotApplicable returns a boolean if a field has been set.

### GetNotes

`func (o *AssessmentResponsesPutPreviousValues) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AssessmentResponsesPutPreviousValues) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AssessmentResponsesPutPreviousValues) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AssessmentResponsesPutPreviousValues) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetUserAssessmentId

`func (o *AssessmentResponsesPutPreviousValues) GetUserAssessmentId() int32`

GetUserAssessmentId returns the UserAssessmentId field if non-nil, zero value otherwise.

### GetUserAssessmentIdOk

`func (o *AssessmentResponsesPutPreviousValues) GetUserAssessmentIdOk() (*int32, bool)`

GetUserAssessmentIdOk returns a tuple with the UserAssessmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssessmentId

`func (o *AssessmentResponsesPutPreviousValues) SetUserAssessmentId(v int32)`

SetUserAssessmentId sets UserAssessmentId field to given value.

### HasUserAssessmentId

`func (o *AssessmentResponsesPutPreviousValues) HasUserAssessmentId() bool`

HasUserAssessmentId returns a boolean if a field has been set.

### GetName

`func (o *AssessmentResponsesPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssessmentResponsesPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssessmentResponsesPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssessmentResponsesPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *AssessmentResponsesPutPreviousValues) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *AssessmentResponsesPutPreviousValues) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *AssessmentResponsesPutPreviousValues) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *AssessmentResponsesPutPreviousValues) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetRiskResponseId

`func (o *AssessmentResponsesPutPreviousValues) GetRiskResponseId() int32`

GetRiskResponseId returns the RiskResponseId field if non-nil, zero value otherwise.

### GetRiskResponseIdOk

`func (o *AssessmentResponsesPutPreviousValues) GetRiskResponseIdOk() (*int32, bool)`

GetRiskResponseIdOk returns a tuple with the RiskResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskResponseId

`func (o *AssessmentResponsesPutPreviousValues) SetRiskResponseId(v int32)`

SetRiskResponseId sets RiskResponseId field to given value.

### HasRiskResponseId

`func (o *AssessmentResponsesPutPreviousValues) HasRiskResponseId() bool`

HasRiskResponseId returns a boolean if a field has been set.

### GetRiskresponseactionId

`func (o *AssessmentResponsesPutPreviousValues) GetRiskresponseactionId() int32`

GetRiskresponseactionId returns the RiskresponseactionId field if non-nil, zero value otherwise.

### GetRiskresponseactionIdOk

`func (o *AssessmentResponsesPutPreviousValues) GetRiskresponseactionIdOk() (*int32, bool)`

GetRiskresponseactionIdOk returns a tuple with the RiskresponseactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskresponseactionId

`func (o *AssessmentResponsesPutPreviousValues) SetRiskresponseactionId(v int32)`

SetRiskresponseactionId sets RiskresponseactionId field to given value.

### HasRiskresponseactionId

`func (o *AssessmentResponsesPutPreviousValues) HasRiskresponseactionId() bool`

HasRiskresponseactionId returns a boolean if a field has been set.

### GetRiskresponseactionType

`func (o *AssessmentResponsesPutPreviousValues) GetRiskresponseactionType() string`

GetRiskresponseactionType returns the RiskresponseactionType field if non-nil, zero value otherwise.

### GetRiskresponseactionTypeOk

`func (o *AssessmentResponsesPutPreviousValues) GetRiskresponseactionTypeOk() (*string, bool)`

GetRiskresponseactionTypeOk returns a tuple with the RiskresponseactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskresponseactionType

`func (o *AssessmentResponsesPutPreviousValues) SetRiskresponseactionType(v string)`

SetRiskresponseactionType sets RiskresponseactionType field to given value.

### HasRiskresponseactionType

`func (o *AssessmentResponsesPutPreviousValues) HasRiskresponseactionType() bool`

HasRiskresponseactionType returns a boolean if a field has been set.

### GetParentAssessmentResponseId

`func (o *AssessmentResponsesPutPreviousValues) GetParentAssessmentResponseId() int32`

GetParentAssessmentResponseId returns the ParentAssessmentResponseId field if non-nil, zero value otherwise.

### GetParentAssessmentResponseIdOk

`func (o *AssessmentResponsesPutPreviousValues) GetParentAssessmentResponseIdOk() (*int32, bool)`

GetParentAssessmentResponseIdOk returns a tuple with the ParentAssessmentResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAssessmentResponseId

`func (o *AssessmentResponsesPutPreviousValues) SetParentAssessmentResponseId(v int32)`

SetParentAssessmentResponseId sets ParentAssessmentResponseId field to given value.

### HasParentAssessmentResponseId

`func (o *AssessmentResponsesPutPreviousValues) HasParentAssessmentResponseId() bool`

HasParentAssessmentResponseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


