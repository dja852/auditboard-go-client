# AssessmentResponsesPutAssessmentResponse

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
**IsAnswered** | **bool** |  | [default to false]
**AssigneeUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IsNotApplicable** | **bool** |  | [default to false]
**Notes** | Pointer to **string** |  | [optional] 
**UserAssessmentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;user_assessments.id&#x60;.&lt;fk table&#x3D;&#39;user_assessments&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskResponseId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;risk_responses.id&#x60;.&lt;fk table&#x3D;&#39;risk_responses&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RiskresponseactionId** | Pointer to **int32** |  | [optional] 
**RiskresponseactionType** | Pointer to **string** |  | [optional] 
**ParentAssessmentResponseId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;assessment_responses.id&#x60;.&lt;fk table&#x3D;&#39;assessment_responses&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 

## Methods

### NewAssessmentResponsesPutAssessmentResponse

`func NewAssessmentResponsesPutAssessmentResponse(isAnswered bool, isNotApplicable bool, ) *AssessmentResponsesPutAssessmentResponse`

NewAssessmentResponsesPutAssessmentResponse instantiates a new AssessmentResponsesPutAssessmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssessmentResponsesPutAssessmentResponseWithDefaults

`func NewAssessmentResponsesPutAssessmentResponseWithDefaults() *AssessmentResponsesPutAssessmentResponse`

NewAssessmentResponsesPutAssessmentResponseWithDefaults instantiates a new AssessmentResponsesPutAssessmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssessmentResponsesPutAssessmentResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssessmentResponsesPutAssessmentResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AssessmentResponsesPutAssessmentResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *AssessmentResponsesPutAssessmentResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssessmentResponsesPutAssessmentResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssessmentResponsesPutAssessmentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResponseJson

`func (o *AssessmentResponsesPutAssessmentResponse) GetResponseJson() interface{}`

GetResponseJson returns the ResponseJson field if non-nil, zero value otherwise.

### GetResponseJsonOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetResponseJsonOk() (*interface{}, bool)`

GetResponseJsonOk returns a tuple with the ResponseJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseJson

`func (o *AssessmentResponsesPutAssessmentResponse) SetResponseJson(v interface{})`

SetResponseJson sets ResponseJson field to given value.

### HasResponseJson

`func (o *AssessmentResponsesPutAssessmentResponse) HasResponseJson() bool`

HasResponseJson returns a boolean if a field has been set.

### SetResponseJsonNil

`func (o *AssessmentResponsesPutAssessmentResponse) SetResponseJsonNil(b bool)`

 SetResponseJsonNil sets the value for ResponseJson to be an explicit nil

### UnsetResponseJson
`func (o *AssessmentResponsesPutAssessmentResponse) UnsetResponseJson()`

UnsetResponseJson ensures that no value is present for ResponseJson, not even an explicit nil
### GetAssessableType

`func (o *AssessmentResponsesPutAssessmentResponse) GetAssessableType() string`

GetAssessableType returns the AssessableType field if non-nil, zero value otherwise.

### GetAssessableTypeOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetAssessableTypeOk() (*string, bool)`

GetAssessableTypeOk returns a tuple with the AssessableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessableType

`func (o *AssessmentResponsesPutAssessmentResponse) SetAssessableType(v string)`

SetAssessableType sets AssessableType field to given value.

### HasAssessableType

`func (o *AssessmentResponsesPutAssessmentResponse) HasAssessableType() bool`

HasAssessableType returns a boolean if a field has been set.

### GetAssessableId

`func (o *AssessmentResponsesPutAssessmentResponse) GetAssessableId() int32`

GetAssessableId returns the AssessableId field if non-nil, zero value otherwise.

### GetAssessableIdOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetAssessableIdOk() (*int32, bool)`

GetAssessableIdOk returns a tuple with the AssessableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessableId

`func (o *AssessmentResponsesPutAssessmentResponse) SetAssessableId(v int32)`

SetAssessableId sets AssessableId field to given value.

### HasAssessableId

`func (o *AssessmentResponsesPutAssessmentResponse) HasAssessableId() bool`

HasAssessableId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AssessmentResponsesPutAssessmentResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AssessmentResponsesPutAssessmentResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AssessmentResponsesPutAssessmentResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AssessmentResponsesPutAssessmentResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AssessmentResponsesPutAssessmentResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AssessmentResponsesPutAssessmentResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AssessmentResponsesPutAssessmentResponse) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AssessmentResponsesPutAssessmentResponse) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AssessmentResponsesPutAssessmentResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetSubmittedByUserId

`func (o *AssessmentResponsesPutAssessmentResponse) GetSubmittedByUserId() int32`

GetSubmittedByUserId returns the SubmittedByUserId field if non-nil, zero value otherwise.

### GetSubmittedByUserIdOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetSubmittedByUserIdOk() (*int32, bool)`

GetSubmittedByUserIdOk returns a tuple with the SubmittedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedByUserId

`func (o *AssessmentResponsesPutAssessmentResponse) SetSubmittedByUserId(v int32)`

SetSubmittedByUserId sets SubmittedByUserId field to given value.

### HasSubmittedByUserId

`func (o *AssessmentResponsesPutAssessmentResponse) HasSubmittedByUserId() bool`

HasSubmittedByUserId returns a boolean if a field has been set.

### GetSubmittedDate

`func (o *AssessmentResponsesPutAssessmentResponse) GetSubmittedDate() string`

GetSubmittedDate returns the SubmittedDate field if non-nil, zero value otherwise.

### GetSubmittedDateOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetSubmittedDateOk() (*string, bool)`

GetSubmittedDateOk returns a tuple with the SubmittedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedDate

`func (o *AssessmentResponsesPutAssessmentResponse) SetSubmittedDate(v string)`

SetSubmittedDate sets SubmittedDate field to given value.

### HasSubmittedDate

`func (o *AssessmentResponsesPutAssessmentResponse) HasSubmittedDate() bool`

HasSubmittedDate returns a boolean if a field has been set.

### GetReferenceMeta

`func (o *AssessmentResponsesPutAssessmentResponse) GetReferenceMeta() interface{}`

GetReferenceMeta returns the ReferenceMeta field if non-nil, zero value otherwise.

### GetReferenceMetaOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetReferenceMetaOk() (*interface{}, bool)`

GetReferenceMetaOk returns a tuple with the ReferenceMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMeta

`func (o *AssessmentResponsesPutAssessmentResponse) SetReferenceMeta(v interface{})`

SetReferenceMeta sets ReferenceMeta field to given value.

### HasReferenceMeta

`func (o *AssessmentResponsesPutAssessmentResponse) HasReferenceMeta() bool`

HasReferenceMeta returns a boolean if a field has been set.

### SetReferenceMetaNil

`func (o *AssessmentResponsesPutAssessmentResponse) SetReferenceMetaNil(b bool)`

 SetReferenceMetaNil sets the value for ReferenceMeta to be an explicit nil

### UnsetReferenceMeta
`func (o *AssessmentResponsesPutAssessmentResponse) UnsetReferenceMeta()`

UnsetReferenceMeta ensures that no value is present for ReferenceMeta, not even an explicit nil
### GetAssessmentId

`func (o *AssessmentResponsesPutAssessmentResponse) GetAssessmentId() int32`

GetAssessmentId returns the AssessmentId field if non-nil, zero value otherwise.

### GetAssessmentIdOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetAssessmentIdOk() (*int32, bool)`

GetAssessmentIdOk returns a tuple with the AssessmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessmentId

`func (o *AssessmentResponsesPutAssessmentResponse) SetAssessmentId(v int32)`

SetAssessmentId sets AssessmentId field to given value.

### HasAssessmentId

`func (o *AssessmentResponsesPutAssessmentResponse) HasAssessmentId() bool`

HasAssessmentId returns a boolean if a field has been set.

### GetSource

`func (o *AssessmentResponsesPutAssessmentResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AssessmentResponsesPutAssessmentResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AssessmentResponsesPutAssessmentResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetAssessmentTemplateId

`func (o *AssessmentResponsesPutAssessmentResponse) GetAssessmentTemplateId() int32`

GetAssessmentTemplateId returns the AssessmentTemplateId field if non-nil, zero value otherwise.

### GetAssessmentTemplateIdOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetAssessmentTemplateIdOk() (*int32, bool)`

GetAssessmentTemplateIdOk returns a tuple with the AssessmentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessmentTemplateId

`func (o *AssessmentResponsesPutAssessmentResponse) SetAssessmentTemplateId(v int32)`

SetAssessmentTemplateId sets AssessmentTemplateId field to given value.

### HasAssessmentTemplateId

`func (o *AssessmentResponsesPutAssessmentResponse) HasAssessmentTemplateId() bool`

HasAssessmentTemplateId returns a boolean if a field has been set.

### GetFinalizedByUserId

`func (o *AssessmentResponsesPutAssessmentResponse) GetFinalizedByUserId() int32`

GetFinalizedByUserId returns the FinalizedByUserId field if non-nil, zero value otherwise.

### GetFinalizedByUserIdOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetFinalizedByUserIdOk() (*int32, bool)`

GetFinalizedByUserIdOk returns a tuple with the FinalizedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizedByUserId

`func (o *AssessmentResponsesPutAssessmentResponse) SetFinalizedByUserId(v int32)`

SetFinalizedByUserId sets FinalizedByUserId field to given value.

### HasFinalizedByUserId

`func (o *AssessmentResponsesPutAssessmentResponse) HasFinalizedByUserId() bool`

HasFinalizedByUserId returns a boolean if a field has been set.

### GetFinalizedDate

`func (o *AssessmentResponsesPutAssessmentResponse) GetFinalizedDate() string`

GetFinalizedDate returns the FinalizedDate field if non-nil, zero value otherwise.

### GetFinalizedDateOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetFinalizedDateOk() (*string, bool)`

GetFinalizedDateOk returns a tuple with the FinalizedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizedDate

`func (o *AssessmentResponsesPutAssessmentResponse) SetFinalizedDate(v string)`

SetFinalizedDate sets FinalizedDate field to given value.

### HasFinalizedDate

`func (o *AssessmentResponsesPutAssessmentResponse) HasFinalizedDate() bool`

HasFinalizedDate returns a boolean if a field has been set.

### GetAnsweredByUserId

`func (o *AssessmentResponsesPutAssessmentResponse) GetAnsweredByUserId() int32`

GetAnsweredByUserId returns the AnsweredByUserId field if non-nil, zero value otherwise.

### GetAnsweredByUserIdOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetAnsweredByUserIdOk() (*int32, bool)`

GetAnsweredByUserIdOk returns a tuple with the AnsweredByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsweredByUserId

`func (o *AssessmentResponsesPutAssessmentResponse) SetAnsweredByUserId(v int32)`

SetAnsweredByUserId sets AnsweredByUserId field to given value.

### HasAnsweredByUserId

`func (o *AssessmentResponsesPutAssessmentResponse) HasAnsweredByUserId() bool`

HasAnsweredByUserId returns a boolean if a field has been set.

### GetAnsweredDate

`func (o *AssessmentResponsesPutAssessmentResponse) GetAnsweredDate() string`

GetAnsweredDate returns the AnsweredDate field if non-nil, zero value otherwise.

### GetAnsweredDateOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetAnsweredDateOk() (*string, bool)`

GetAnsweredDateOk returns a tuple with the AnsweredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsweredDate

`func (o *AssessmentResponsesPutAssessmentResponse) SetAnsweredDate(v string)`

SetAnsweredDate sets AnsweredDate field to given value.

### HasAnsweredDate

`func (o *AssessmentResponsesPutAssessmentResponse) HasAnsweredDate() bool`

HasAnsweredDate returns a boolean if a field has been set.

### GetIsAnswered

`func (o *AssessmentResponsesPutAssessmentResponse) GetIsAnswered() bool`

GetIsAnswered returns the IsAnswered field if non-nil, zero value otherwise.

### GetIsAnsweredOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetIsAnsweredOk() (*bool, bool)`

GetIsAnsweredOk returns a tuple with the IsAnswered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnswered

`func (o *AssessmentResponsesPutAssessmentResponse) SetIsAnswered(v bool)`

SetIsAnswered sets IsAnswered field to given value.


### GetAssigneeUserId

`func (o *AssessmentResponsesPutAssessmentResponse) GetAssigneeUserId() int32`

GetAssigneeUserId returns the AssigneeUserId field if non-nil, zero value otherwise.

### GetAssigneeUserIdOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetAssigneeUserIdOk() (*int32, bool)`

GetAssigneeUserIdOk returns a tuple with the AssigneeUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeUserId

`func (o *AssessmentResponsesPutAssessmentResponse) SetAssigneeUserId(v int32)`

SetAssigneeUserId sets AssigneeUserId field to given value.

### HasAssigneeUserId

`func (o *AssessmentResponsesPutAssessmentResponse) HasAssigneeUserId() bool`

HasAssigneeUserId returns a boolean if a field has been set.

### GetIsNotApplicable

`func (o *AssessmentResponsesPutAssessmentResponse) GetIsNotApplicable() bool`

GetIsNotApplicable returns the IsNotApplicable field if non-nil, zero value otherwise.

### GetIsNotApplicableOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetIsNotApplicableOk() (*bool, bool)`

GetIsNotApplicableOk returns a tuple with the IsNotApplicable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotApplicable

`func (o *AssessmentResponsesPutAssessmentResponse) SetIsNotApplicable(v bool)`

SetIsNotApplicable sets IsNotApplicable field to given value.


### GetNotes

`func (o *AssessmentResponsesPutAssessmentResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AssessmentResponsesPutAssessmentResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AssessmentResponsesPutAssessmentResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetUserAssessmentId

`func (o *AssessmentResponsesPutAssessmentResponse) GetUserAssessmentId() int32`

GetUserAssessmentId returns the UserAssessmentId field if non-nil, zero value otherwise.

### GetUserAssessmentIdOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetUserAssessmentIdOk() (*int32, bool)`

GetUserAssessmentIdOk returns a tuple with the UserAssessmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssessmentId

`func (o *AssessmentResponsesPutAssessmentResponse) SetUserAssessmentId(v int32)`

SetUserAssessmentId sets UserAssessmentId field to given value.

### HasUserAssessmentId

`func (o *AssessmentResponsesPutAssessmentResponse) HasUserAssessmentId() bool`

HasUserAssessmentId returns a boolean if a field has been set.

### GetName

`func (o *AssessmentResponsesPutAssessmentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssessmentResponsesPutAssessmentResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssessmentResponsesPutAssessmentResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *AssessmentResponsesPutAssessmentResponse) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *AssessmentResponsesPutAssessmentResponse) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *AssessmentResponsesPutAssessmentResponse) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetRiskResponseId

`func (o *AssessmentResponsesPutAssessmentResponse) GetRiskResponseId() int32`

GetRiskResponseId returns the RiskResponseId field if non-nil, zero value otherwise.

### GetRiskResponseIdOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetRiskResponseIdOk() (*int32, bool)`

GetRiskResponseIdOk returns a tuple with the RiskResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskResponseId

`func (o *AssessmentResponsesPutAssessmentResponse) SetRiskResponseId(v int32)`

SetRiskResponseId sets RiskResponseId field to given value.

### HasRiskResponseId

`func (o *AssessmentResponsesPutAssessmentResponse) HasRiskResponseId() bool`

HasRiskResponseId returns a boolean if a field has been set.

### GetRiskresponseactionId

`func (o *AssessmentResponsesPutAssessmentResponse) GetRiskresponseactionId() int32`

GetRiskresponseactionId returns the RiskresponseactionId field if non-nil, zero value otherwise.

### GetRiskresponseactionIdOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetRiskresponseactionIdOk() (*int32, bool)`

GetRiskresponseactionIdOk returns a tuple with the RiskresponseactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskresponseactionId

`func (o *AssessmentResponsesPutAssessmentResponse) SetRiskresponseactionId(v int32)`

SetRiskresponseactionId sets RiskresponseactionId field to given value.

### HasRiskresponseactionId

`func (o *AssessmentResponsesPutAssessmentResponse) HasRiskresponseactionId() bool`

HasRiskresponseactionId returns a boolean if a field has been set.

### GetRiskresponseactionType

`func (o *AssessmentResponsesPutAssessmentResponse) GetRiskresponseactionType() string`

GetRiskresponseactionType returns the RiskresponseactionType field if non-nil, zero value otherwise.

### GetRiskresponseactionTypeOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetRiskresponseactionTypeOk() (*string, bool)`

GetRiskresponseactionTypeOk returns a tuple with the RiskresponseactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskresponseactionType

`func (o *AssessmentResponsesPutAssessmentResponse) SetRiskresponseactionType(v string)`

SetRiskresponseactionType sets RiskresponseactionType field to given value.

### HasRiskresponseactionType

`func (o *AssessmentResponsesPutAssessmentResponse) HasRiskresponseactionType() bool`

HasRiskresponseactionType returns a boolean if a field has been set.

### GetParentAssessmentResponseId

`func (o *AssessmentResponsesPutAssessmentResponse) GetParentAssessmentResponseId() int32`

GetParentAssessmentResponseId returns the ParentAssessmentResponseId field if non-nil, zero value otherwise.

### GetParentAssessmentResponseIdOk

`func (o *AssessmentResponsesPutAssessmentResponse) GetParentAssessmentResponseIdOk() (*int32, bool)`

GetParentAssessmentResponseIdOk returns a tuple with the ParentAssessmentResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAssessmentResponseId

`func (o *AssessmentResponsesPutAssessmentResponse) SetParentAssessmentResponseId(v int32)`

SetParentAssessmentResponseId sets ParentAssessmentResponseId field to given value.

### HasParentAssessmentResponseId

`func (o *AssessmentResponsesPutAssessmentResponse) HasParentAssessmentResponseId() bool`

HasParentAssessmentResponseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


