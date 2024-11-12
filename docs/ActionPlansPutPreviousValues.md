# ActionPlansPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**IsFlattened** | Pointer to **bool** |  | [optional] [default to false]
**Flattened** | Pointer to **interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DiscussionDate** | Pointer to **string** |  | [optional] 
**RemediationDate** | Pointer to **string** |  | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**ExternalPlannedConfirmDate** | Pointer to **string** |  | [optional] 
**ExternalConfirmDate** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**FinalizedDate** | Pointer to **string** |  | [optional] 
**CompletedDate** | Pointer to **string** |  | [optional] 
**PlannedCloseDate** | Pointer to **string** |  | [optional] 
**PlannedConfirmDate** | Pointer to **string** |  | [optional] 
**VicePresidentUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ExecutiveVicePresidentUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**BusinessOwnerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ExecutiveOwnerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RemediationOwnerId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OfficerUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AuditDirectorUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AuditContactUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**BusinessContactUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**RemediationAction** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ManagementResponse** | Pointer to **string** |  | [optional] 
**IssueId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issues.id&#x60;.&lt;fk table&#x3D;&#39;issues&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueRemediationCustomSelect1OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_remediation_custom_select1_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_remediation_custom_select1_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueRemediationCustomSelect2OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_remediation_custom_select2_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_remediation_custom_select2_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audits.id&#x60;.&lt;fk table&#x3D;&#39;ops_audits&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditItemId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_items.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_items&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomText1** | Pointer to **string** |  | [optional] 
**CustomText2** | Pointer to **string** |  | [optional] 
**CustomText3** | Pointer to **string** |  | [optional] 
**IssueRemediationCustomSelect3OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_remediation_custom_select3_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_remediation_custom_select3_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueRemediationCustomSelect4OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_remediation_custom_select4_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_remediation_custom_select4_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ExternalIntegrationUrl** | Pointer to **string** |  | [optional] 
**IssueRemediationCustomSelect5OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_remediation_custom_select5_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_remediation_custom_select5_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueRemediationCustomSelect6OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_remediation_custom_select6_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_remediation_custom_select6_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueRemediationCustomSelect7OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_remediation_custom_select7_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_remediation_custom_select7_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueRemediationCustomSelect8OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_remediation_custom_select8_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_remediation_custom_select8_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CustomText4** | Pointer to **string** |  | [optional] 
**CustomText5** | Pointer to **string** |  | [optional] 
**CustomText6** | Pointer to **string** |  | [optional] 
**FirstCompletedDate** | Pointer to **string** |  | [optional] 
**CustomCurrency1** | Pointer to **float32** |  | [optional] 
**CustomCurrency2** | Pointer to **float32** |  | [optional] 
**CustomCurrency3** | Pointer to **float32** |  | [optional] 
**CustomCurrency4** | Pointer to **float32** |  | [optional] 
**CustomText7** | Pointer to **string** |  | [optional] 
**CustomText8** | Pointer to **string** |  | [optional] 
**CustomText9** | Pointer to **string** |  | [optional] 
**CustomText10** | Pointer to **string** |  | [optional] 
**CustomText11** | Pointer to **string** |  | [optional] 
**CustomText12** | Pointer to **string** |  | [optional] 
**CustomText13** | Pointer to **string** |  | [optional] 
**CustomText14** | Pointer to **string** |  | [optional] 
**CustomText15** | Pointer to **string** |  | [optional] 
**CustomText16** | Pointer to **string** |  | [optional] 
**CustomText17** | Pointer to **string** |  | [optional] 
**CustomText18** | Pointer to **string** |  | [optional] 
**CustomText19** | Pointer to **string** |  | [optional] 
**CustomText20** | Pointer to **string** |  | [optional] 
**CustomText21** | Pointer to **string** |  | [optional] 
**IssueRemediationCustomSelect9OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_remediation_custom_select9_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_remediation_custom_select9_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueRemediationCustomSelect10OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_remediation_custom_select10_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_remediation_custom_select10_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueRemediationCustomSelect11OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_remediation_custom_select11_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_remediation_custom_select11_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueRemediationCustomSelect12OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_remediation_custom_select12_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_remediation_custom_select12_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueRemediationCustomSelect13OptionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issue_remediation_custom_select13_options.id&#x60;.&lt;fk table&#x3D;&#39;issue_remediation_custom_select13_options&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FieldData** | Pointer to **interface{}** |  | [optional] 
**CompleteReason** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **interface{}** |  | [optional] 
**ComplianceAssessmentId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;compliance_assessments.id&#x60;.&lt;fk table&#x3D;&#39;compliance_assessments&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ComplianceAssessmentItemId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;compliance_assessment_items.id&#x60;.&lt;fk table&#x3D;&#39;compliance_assessment_items&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FirstPendingRemediationDueDate** | Pointer to **string** |  | [optional] 
**PendingRemediationDueDateCount** | Pointer to **int32** |  | [optional] [default to 0]
**FirstOpenDueDate** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewActionPlansPutPreviousValues

`func NewActionPlansPutPreviousValues() *ActionPlansPutPreviousValues`

NewActionPlansPutPreviousValues instantiates a new ActionPlansPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionPlansPutPreviousValuesWithDefaults

`func NewActionPlansPutPreviousValuesWithDefaults() *ActionPlansPutPreviousValues`

NewActionPlansPutPreviousValuesWithDefaults instantiates a new ActionPlansPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionPlansPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionPlansPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionPlansPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ActionPlansPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ActionPlansPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ActionPlansPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ActionPlansPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ActionPlansPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ActionPlansPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ActionPlansPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ActionPlansPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ActionPlansPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ActionPlansPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ActionPlansPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ActionPlansPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ActionPlansPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetIsFlattened

`func (o *ActionPlansPutPreviousValues) GetIsFlattened() bool`

GetIsFlattened returns the IsFlattened field if non-nil, zero value otherwise.

### GetIsFlattenedOk

`func (o *ActionPlansPutPreviousValues) GetIsFlattenedOk() (*bool, bool)`

GetIsFlattenedOk returns a tuple with the IsFlattened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlattened

`func (o *ActionPlansPutPreviousValues) SetIsFlattened(v bool)`

SetIsFlattened sets IsFlattened field to given value.

### HasIsFlattened

`func (o *ActionPlansPutPreviousValues) HasIsFlattened() bool`

HasIsFlattened returns a boolean if a field has been set.

### GetFlattened

`func (o *ActionPlansPutPreviousValues) GetFlattened() interface{}`

GetFlattened returns the Flattened field if non-nil, zero value otherwise.

### GetFlattenedOk

`func (o *ActionPlansPutPreviousValues) GetFlattenedOk() (*interface{}, bool)`

GetFlattenedOk returns a tuple with the Flattened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlattened

`func (o *ActionPlansPutPreviousValues) SetFlattened(v interface{})`

SetFlattened sets Flattened field to given value.

### HasFlattened

`func (o *ActionPlansPutPreviousValues) HasFlattened() bool`

HasFlattened returns a boolean if a field has been set.

### SetFlattenedNil

`func (o *ActionPlansPutPreviousValues) SetFlattenedNil(b bool)`

 SetFlattenedNil sets the value for Flattened to be an explicit nil

### UnsetFlattened
`func (o *ActionPlansPutPreviousValues) UnsetFlattened()`

UnsetFlattened ensures that no value is present for Flattened, not even an explicit nil
### GetStatus

`func (o *ActionPlansPutPreviousValues) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionPlansPutPreviousValues) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionPlansPutPreviousValues) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActionPlansPutPreviousValues) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDiscussionDate

`func (o *ActionPlansPutPreviousValues) GetDiscussionDate() string`

GetDiscussionDate returns the DiscussionDate field if non-nil, zero value otherwise.

### GetDiscussionDateOk

`func (o *ActionPlansPutPreviousValues) GetDiscussionDateOk() (*string, bool)`

GetDiscussionDateOk returns a tuple with the DiscussionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscussionDate

`func (o *ActionPlansPutPreviousValues) SetDiscussionDate(v string)`

SetDiscussionDate sets DiscussionDate field to given value.

### HasDiscussionDate

`func (o *ActionPlansPutPreviousValues) HasDiscussionDate() bool`

HasDiscussionDate returns a boolean if a field has been set.

### GetRemediationDate

`func (o *ActionPlansPutPreviousValues) GetRemediationDate() string`

GetRemediationDate returns the RemediationDate field if non-nil, zero value otherwise.

### GetRemediationDateOk

`func (o *ActionPlansPutPreviousValues) GetRemediationDateOk() (*string, bool)`

GetRemediationDateOk returns a tuple with the RemediationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationDate

`func (o *ActionPlansPutPreviousValues) SetRemediationDate(v string)`

SetRemediationDate sets RemediationDate field to given value.

### HasRemediationDate

`func (o *ActionPlansPutPreviousValues) HasRemediationDate() bool`

HasRemediationDate returns a boolean if a field has been set.

### GetDueDate

`func (o *ActionPlansPutPreviousValues) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *ActionPlansPutPreviousValues) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *ActionPlansPutPreviousValues) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *ActionPlansPutPreviousValues) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetExternalPlannedConfirmDate

`func (o *ActionPlansPutPreviousValues) GetExternalPlannedConfirmDate() string`

GetExternalPlannedConfirmDate returns the ExternalPlannedConfirmDate field if non-nil, zero value otherwise.

### GetExternalPlannedConfirmDateOk

`func (o *ActionPlansPutPreviousValues) GetExternalPlannedConfirmDateOk() (*string, bool)`

GetExternalPlannedConfirmDateOk returns a tuple with the ExternalPlannedConfirmDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPlannedConfirmDate

`func (o *ActionPlansPutPreviousValues) SetExternalPlannedConfirmDate(v string)`

SetExternalPlannedConfirmDate sets ExternalPlannedConfirmDate field to given value.

### HasExternalPlannedConfirmDate

`func (o *ActionPlansPutPreviousValues) HasExternalPlannedConfirmDate() bool`

HasExternalPlannedConfirmDate returns a boolean if a field has been set.

### GetExternalConfirmDate

`func (o *ActionPlansPutPreviousValues) GetExternalConfirmDate() string`

GetExternalConfirmDate returns the ExternalConfirmDate field if non-nil, zero value otherwise.

### GetExternalConfirmDateOk

`func (o *ActionPlansPutPreviousValues) GetExternalConfirmDateOk() (*string, bool)`

GetExternalConfirmDateOk returns a tuple with the ExternalConfirmDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConfirmDate

`func (o *ActionPlansPutPreviousValues) SetExternalConfirmDate(v string)`

SetExternalConfirmDate sets ExternalConfirmDate field to given value.

### HasExternalConfirmDate

`func (o *ActionPlansPutPreviousValues) HasExternalConfirmDate() bool`

HasExternalConfirmDate returns a boolean if a field has been set.

### GetStartDate

`func (o *ActionPlansPutPreviousValues) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ActionPlansPutPreviousValues) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ActionPlansPutPreviousValues) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ActionPlansPutPreviousValues) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetFinalizedDate

`func (o *ActionPlansPutPreviousValues) GetFinalizedDate() string`

GetFinalizedDate returns the FinalizedDate field if non-nil, zero value otherwise.

### GetFinalizedDateOk

`func (o *ActionPlansPutPreviousValues) GetFinalizedDateOk() (*string, bool)`

GetFinalizedDateOk returns a tuple with the FinalizedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizedDate

`func (o *ActionPlansPutPreviousValues) SetFinalizedDate(v string)`

SetFinalizedDate sets FinalizedDate field to given value.

### HasFinalizedDate

`func (o *ActionPlansPutPreviousValues) HasFinalizedDate() bool`

HasFinalizedDate returns a boolean if a field has been set.

### GetCompletedDate

`func (o *ActionPlansPutPreviousValues) GetCompletedDate() string`

GetCompletedDate returns the CompletedDate field if non-nil, zero value otherwise.

### GetCompletedDateOk

`func (o *ActionPlansPutPreviousValues) GetCompletedDateOk() (*string, bool)`

GetCompletedDateOk returns a tuple with the CompletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDate

`func (o *ActionPlansPutPreviousValues) SetCompletedDate(v string)`

SetCompletedDate sets CompletedDate field to given value.

### HasCompletedDate

`func (o *ActionPlansPutPreviousValues) HasCompletedDate() bool`

HasCompletedDate returns a boolean if a field has been set.

### GetPlannedCloseDate

`func (o *ActionPlansPutPreviousValues) GetPlannedCloseDate() string`

GetPlannedCloseDate returns the PlannedCloseDate field if non-nil, zero value otherwise.

### GetPlannedCloseDateOk

`func (o *ActionPlansPutPreviousValues) GetPlannedCloseDateOk() (*string, bool)`

GetPlannedCloseDateOk returns a tuple with the PlannedCloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedCloseDate

`func (o *ActionPlansPutPreviousValues) SetPlannedCloseDate(v string)`

SetPlannedCloseDate sets PlannedCloseDate field to given value.

### HasPlannedCloseDate

`func (o *ActionPlansPutPreviousValues) HasPlannedCloseDate() bool`

HasPlannedCloseDate returns a boolean if a field has been set.

### GetPlannedConfirmDate

`func (o *ActionPlansPutPreviousValues) GetPlannedConfirmDate() string`

GetPlannedConfirmDate returns the PlannedConfirmDate field if non-nil, zero value otherwise.

### GetPlannedConfirmDateOk

`func (o *ActionPlansPutPreviousValues) GetPlannedConfirmDateOk() (*string, bool)`

GetPlannedConfirmDateOk returns a tuple with the PlannedConfirmDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedConfirmDate

`func (o *ActionPlansPutPreviousValues) SetPlannedConfirmDate(v string)`

SetPlannedConfirmDate sets PlannedConfirmDate field to given value.

### HasPlannedConfirmDate

`func (o *ActionPlansPutPreviousValues) HasPlannedConfirmDate() bool`

HasPlannedConfirmDate returns a boolean if a field has been set.

### GetVicePresidentUserId

`func (o *ActionPlansPutPreviousValues) GetVicePresidentUserId() int32`

GetVicePresidentUserId returns the VicePresidentUserId field if non-nil, zero value otherwise.

### GetVicePresidentUserIdOk

`func (o *ActionPlansPutPreviousValues) GetVicePresidentUserIdOk() (*int32, bool)`

GetVicePresidentUserIdOk returns a tuple with the VicePresidentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVicePresidentUserId

`func (o *ActionPlansPutPreviousValues) SetVicePresidentUserId(v int32)`

SetVicePresidentUserId sets VicePresidentUserId field to given value.

### HasVicePresidentUserId

`func (o *ActionPlansPutPreviousValues) HasVicePresidentUserId() bool`

HasVicePresidentUserId returns a boolean if a field has been set.

### GetExecutiveVicePresidentUserId

`func (o *ActionPlansPutPreviousValues) GetExecutiveVicePresidentUserId() int32`

GetExecutiveVicePresidentUserId returns the ExecutiveVicePresidentUserId field if non-nil, zero value otherwise.

### GetExecutiveVicePresidentUserIdOk

`func (o *ActionPlansPutPreviousValues) GetExecutiveVicePresidentUserIdOk() (*int32, bool)`

GetExecutiveVicePresidentUserIdOk returns a tuple with the ExecutiveVicePresidentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutiveVicePresidentUserId

`func (o *ActionPlansPutPreviousValues) SetExecutiveVicePresidentUserId(v int32)`

SetExecutiveVicePresidentUserId sets ExecutiveVicePresidentUserId field to given value.

### HasExecutiveVicePresidentUserId

`func (o *ActionPlansPutPreviousValues) HasExecutiveVicePresidentUserId() bool`

HasExecutiveVicePresidentUserId returns a boolean if a field has been set.

### GetBusinessOwnerUserId

`func (o *ActionPlansPutPreviousValues) GetBusinessOwnerUserId() int32`

GetBusinessOwnerUserId returns the BusinessOwnerUserId field if non-nil, zero value otherwise.

### GetBusinessOwnerUserIdOk

`func (o *ActionPlansPutPreviousValues) GetBusinessOwnerUserIdOk() (*int32, bool)`

GetBusinessOwnerUserIdOk returns a tuple with the BusinessOwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessOwnerUserId

`func (o *ActionPlansPutPreviousValues) SetBusinessOwnerUserId(v int32)`

SetBusinessOwnerUserId sets BusinessOwnerUserId field to given value.

### HasBusinessOwnerUserId

`func (o *ActionPlansPutPreviousValues) HasBusinessOwnerUserId() bool`

HasBusinessOwnerUserId returns a boolean if a field has been set.

### GetExecutiveOwnerUserId

`func (o *ActionPlansPutPreviousValues) GetExecutiveOwnerUserId() int32`

GetExecutiveOwnerUserId returns the ExecutiveOwnerUserId field if non-nil, zero value otherwise.

### GetExecutiveOwnerUserIdOk

`func (o *ActionPlansPutPreviousValues) GetExecutiveOwnerUserIdOk() (*int32, bool)`

GetExecutiveOwnerUserIdOk returns a tuple with the ExecutiveOwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutiveOwnerUserId

`func (o *ActionPlansPutPreviousValues) SetExecutiveOwnerUserId(v int32)`

SetExecutiveOwnerUserId sets ExecutiveOwnerUserId field to given value.

### HasExecutiveOwnerUserId

`func (o *ActionPlansPutPreviousValues) HasExecutiveOwnerUserId() bool`

HasExecutiveOwnerUserId returns a boolean if a field has been set.

### GetRemediationOwnerId

`func (o *ActionPlansPutPreviousValues) GetRemediationOwnerId() int32`

GetRemediationOwnerId returns the RemediationOwnerId field if non-nil, zero value otherwise.

### GetRemediationOwnerIdOk

`func (o *ActionPlansPutPreviousValues) GetRemediationOwnerIdOk() (*int32, bool)`

GetRemediationOwnerIdOk returns a tuple with the RemediationOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationOwnerId

`func (o *ActionPlansPutPreviousValues) SetRemediationOwnerId(v int32)`

SetRemediationOwnerId sets RemediationOwnerId field to given value.

### HasRemediationOwnerId

`func (o *ActionPlansPutPreviousValues) HasRemediationOwnerId() bool`

HasRemediationOwnerId returns a boolean if a field has been set.

### GetOfficerUserId

`func (o *ActionPlansPutPreviousValues) GetOfficerUserId() int32`

GetOfficerUserId returns the OfficerUserId field if non-nil, zero value otherwise.

### GetOfficerUserIdOk

`func (o *ActionPlansPutPreviousValues) GetOfficerUserIdOk() (*int32, bool)`

GetOfficerUserIdOk returns a tuple with the OfficerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficerUserId

`func (o *ActionPlansPutPreviousValues) SetOfficerUserId(v int32)`

SetOfficerUserId sets OfficerUserId field to given value.

### HasOfficerUserId

`func (o *ActionPlansPutPreviousValues) HasOfficerUserId() bool`

HasOfficerUserId returns a boolean if a field has been set.

### GetAuditDirectorUserId

`func (o *ActionPlansPutPreviousValues) GetAuditDirectorUserId() int32`

GetAuditDirectorUserId returns the AuditDirectorUserId field if non-nil, zero value otherwise.

### GetAuditDirectorUserIdOk

`func (o *ActionPlansPutPreviousValues) GetAuditDirectorUserIdOk() (*int32, bool)`

GetAuditDirectorUserIdOk returns a tuple with the AuditDirectorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditDirectorUserId

`func (o *ActionPlansPutPreviousValues) SetAuditDirectorUserId(v int32)`

SetAuditDirectorUserId sets AuditDirectorUserId field to given value.

### HasAuditDirectorUserId

`func (o *ActionPlansPutPreviousValues) HasAuditDirectorUserId() bool`

HasAuditDirectorUserId returns a boolean if a field has been set.

### GetAuditContactUserId

`func (o *ActionPlansPutPreviousValues) GetAuditContactUserId() int32`

GetAuditContactUserId returns the AuditContactUserId field if non-nil, zero value otherwise.

### GetAuditContactUserIdOk

`func (o *ActionPlansPutPreviousValues) GetAuditContactUserIdOk() (*int32, bool)`

GetAuditContactUserIdOk returns a tuple with the AuditContactUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditContactUserId

`func (o *ActionPlansPutPreviousValues) SetAuditContactUserId(v int32)`

SetAuditContactUserId sets AuditContactUserId field to given value.

### HasAuditContactUserId

`func (o *ActionPlansPutPreviousValues) HasAuditContactUserId() bool`

HasAuditContactUserId returns a boolean if a field has been set.

### GetBusinessContactUserId

`func (o *ActionPlansPutPreviousValues) GetBusinessContactUserId() int32`

GetBusinessContactUserId returns the BusinessContactUserId field if non-nil, zero value otherwise.

### GetBusinessContactUserIdOk

`func (o *ActionPlansPutPreviousValues) GetBusinessContactUserIdOk() (*int32, bool)`

GetBusinessContactUserIdOk returns a tuple with the BusinessContactUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessContactUserId

`func (o *ActionPlansPutPreviousValues) SetBusinessContactUserId(v int32)`

SetBusinessContactUserId sets BusinessContactUserId field to given value.

### HasBusinessContactUserId

`func (o *ActionPlansPutPreviousValues) HasBusinessContactUserId() bool`

HasBusinessContactUserId returns a boolean if a field has been set.

### GetRemediationAction

`func (o *ActionPlansPutPreviousValues) GetRemediationAction() string`

GetRemediationAction returns the RemediationAction field if non-nil, zero value otherwise.

### GetRemediationActionOk

`func (o *ActionPlansPutPreviousValues) GetRemediationActionOk() (*string, bool)`

GetRemediationActionOk returns a tuple with the RemediationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationAction

`func (o *ActionPlansPutPreviousValues) SetRemediationAction(v string)`

SetRemediationAction sets RemediationAction field to given value.

### HasRemediationAction

`func (o *ActionPlansPutPreviousValues) HasRemediationAction() bool`

HasRemediationAction returns a boolean if a field has been set.

### GetDescription

`func (o *ActionPlansPutPreviousValues) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ActionPlansPutPreviousValues) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ActionPlansPutPreviousValues) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ActionPlansPutPreviousValues) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetManagementResponse

`func (o *ActionPlansPutPreviousValues) GetManagementResponse() string`

GetManagementResponse returns the ManagementResponse field if non-nil, zero value otherwise.

### GetManagementResponseOk

`func (o *ActionPlansPutPreviousValues) GetManagementResponseOk() (*string, bool)`

GetManagementResponseOk returns a tuple with the ManagementResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementResponse

`func (o *ActionPlansPutPreviousValues) SetManagementResponse(v string)`

SetManagementResponse sets ManagementResponse field to given value.

### HasManagementResponse

`func (o *ActionPlansPutPreviousValues) HasManagementResponse() bool`

HasManagementResponse returns a boolean if a field has been set.

### GetIssueId

`func (o *ActionPlansPutPreviousValues) GetIssueId() int32`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *ActionPlansPutPreviousValues) GetIssueIdOk() (*int32, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *ActionPlansPutPreviousValues) SetIssueId(v int32)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *ActionPlansPutPreviousValues) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.

### GetIssueRemediationCustomSelect1OptionId

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect1OptionId() int32`

GetIssueRemediationCustomSelect1OptionId returns the IssueRemediationCustomSelect1OptionId field if non-nil, zero value otherwise.

### GetIssueRemediationCustomSelect1OptionIdOk

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect1OptionIdOk() (*int32, bool)`

GetIssueRemediationCustomSelect1OptionIdOk returns a tuple with the IssueRemediationCustomSelect1OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueRemediationCustomSelect1OptionId

`func (o *ActionPlansPutPreviousValues) SetIssueRemediationCustomSelect1OptionId(v int32)`

SetIssueRemediationCustomSelect1OptionId sets IssueRemediationCustomSelect1OptionId field to given value.

### HasIssueRemediationCustomSelect1OptionId

`func (o *ActionPlansPutPreviousValues) HasIssueRemediationCustomSelect1OptionId() bool`

HasIssueRemediationCustomSelect1OptionId returns a boolean if a field has been set.

### GetIssueRemediationCustomSelect2OptionId

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect2OptionId() int32`

GetIssueRemediationCustomSelect2OptionId returns the IssueRemediationCustomSelect2OptionId field if non-nil, zero value otherwise.

### GetIssueRemediationCustomSelect2OptionIdOk

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect2OptionIdOk() (*int32, bool)`

GetIssueRemediationCustomSelect2OptionIdOk returns a tuple with the IssueRemediationCustomSelect2OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueRemediationCustomSelect2OptionId

`func (o *ActionPlansPutPreviousValues) SetIssueRemediationCustomSelect2OptionId(v int32)`

SetIssueRemediationCustomSelect2OptionId sets IssueRemediationCustomSelect2OptionId field to given value.

### HasIssueRemediationCustomSelect2OptionId

`func (o *ActionPlansPutPreviousValues) HasIssueRemediationCustomSelect2OptionId() bool`

HasIssueRemediationCustomSelect2OptionId returns a boolean if a field has been set.

### GetOpsAuditId

`func (o *ActionPlansPutPreviousValues) GetOpsAuditId() int32`

GetOpsAuditId returns the OpsAuditId field if non-nil, zero value otherwise.

### GetOpsAuditIdOk

`func (o *ActionPlansPutPreviousValues) GetOpsAuditIdOk() (*int32, bool)`

GetOpsAuditIdOk returns a tuple with the OpsAuditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditId

`func (o *ActionPlansPutPreviousValues) SetOpsAuditId(v int32)`

SetOpsAuditId sets OpsAuditId field to given value.

### HasOpsAuditId

`func (o *ActionPlansPutPreviousValues) HasOpsAuditId() bool`

HasOpsAuditId returns a boolean if a field has been set.

### GetOpsAuditItemId

`func (o *ActionPlansPutPreviousValues) GetOpsAuditItemId() int32`

GetOpsAuditItemId returns the OpsAuditItemId field if non-nil, zero value otherwise.

### GetOpsAuditItemIdOk

`func (o *ActionPlansPutPreviousValues) GetOpsAuditItemIdOk() (*int32, bool)`

GetOpsAuditItemIdOk returns a tuple with the OpsAuditItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditItemId

`func (o *ActionPlansPutPreviousValues) SetOpsAuditItemId(v int32)`

SetOpsAuditItemId sets OpsAuditItemId field to given value.

### HasOpsAuditItemId

`func (o *ActionPlansPutPreviousValues) HasOpsAuditItemId() bool`

HasOpsAuditItemId returns a boolean if a field has been set.

### GetCustomText1

`func (o *ActionPlansPutPreviousValues) GetCustomText1() string`

GetCustomText1 returns the CustomText1 field if non-nil, zero value otherwise.

### GetCustomText1Ok

`func (o *ActionPlansPutPreviousValues) GetCustomText1Ok() (*string, bool)`

GetCustomText1Ok returns a tuple with the CustomText1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText1

`func (o *ActionPlansPutPreviousValues) SetCustomText1(v string)`

SetCustomText1 sets CustomText1 field to given value.

### HasCustomText1

`func (o *ActionPlansPutPreviousValues) HasCustomText1() bool`

HasCustomText1 returns a boolean if a field has been set.

### GetCustomText2

`func (o *ActionPlansPutPreviousValues) GetCustomText2() string`

GetCustomText2 returns the CustomText2 field if non-nil, zero value otherwise.

### GetCustomText2Ok

`func (o *ActionPlansPutPreviousValues) GetCustomText2Ok() (*string, bool)`

GetCustomText2Ok returns a tuple with the CustomText2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText2

`func (o *ActionPlansPutPreviousValues) SetCustomText2(v string)`

SetCustomText2 sets CustomText2 field to given value.

### HasCustomText2

`func (o *ActionPlansPutPreviousValues) HasCustomText2() bool`

HasCustomText2 returns a boolean if a field has been set.

### GetCustomText3

`func (o *ActionPlansPutPreviousValues) GetCustomText3() string`

GetCustomText3 returns the CustomText3 field if non-nil, zero value otherwise.

### GetCustomText3Ok

`func (o *ActionPlansPutPreviousValues) GetCustomText3Ok() (*string, bool)`

GetCustomText3Ok returns a tuple with the CustomText3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText3

`func (o *ActionPlansPutPreviousValues) SetCustomText3(v string)`

SetCustomText3 sets CustomText3 field to given value.

### HasCustomText3

`func (o *ActionPlansPutPreviousValues) HasCustomText3() bool`

HasCustomText3 returns a boolean if a field has been set.

### GetIssueRemediationCustomSelect3OptionId

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect3OptionId() int32`

GetIssueRemediationCustomSelect3OptionId returns the IssueRemediationCustomSelect3OptionId field if non-nil, zero value otherwise.

### GetIssueRemediationCustomSelect3OptionIdOk

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect3OptionIdOk() (*int32, bool)`

GetIssueRemediationCustomSelect3OptionIdOk returns a tuple with the IssueRemediationCustomSelect3OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueRemediationCustomSelect3OptionId

`func (o *ActionPlansPutPreviousValues) SetIssueRemediationCustomSelect3OptionId(v int32)`

SetIssueRemediationCustomSelect3OptionId sets IssueRemediationCustomSelect3OptionId field to given value.

### HasIssueRemediationCustomSelect3OptionId

`func (o *ActionPlansPutPreviousValues) HasIssueRemediationCustomSelect3OptionId() bool`

HasIssueRemediationCustomSelect3OptionId returns a boolean if a field has been set.

### GetIssueRemediationCustomSelect4OptionId

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect4OptionId() int32`

GetIssueRemediationCustomSelect4OptionId returns the IssueRemediationCustomSelect4OptionId field if non-nil, zero value otherwise.

### GetIssueRemediationCustomSelect4OptionIdOk

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect4OptionIdOk() (*int32, bool)`

GetIssueRemediationCustomSelect4OptionIdOk returns a tuple with the IssueRemediationCustomSelect4OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueRemediationCustomSelect4OptionId

`func (o *ActionPlansPutPreviousValues) SetIssueRemediationCustomSelect4OptionId(v int32)`

SetIssueRemediationCustomSelect4OptionId sets IssueRemediationCustomSelect4OptionId field to given value.

### HasIssueRemediationCustomSelect4OptionId

`func (o *ActionPlansPutPreviousValues) HasIssueRemediationCustomSelect4OptionId() bool`

HasIssueRemediationCustomSelect4OptionId returns a boolean if a field has been set.

### GetExternalIntegrationUrl

`func (o *ActionPlansPutPreviousValues) GetExternalIntegrationUrl() string`

GetExternalIntegrationUrl returns the ExternalIntegrationUrl field if non-nil, zero value otherwise.

### GetExternalIntegrationUrlOk

`func (o *ActionPlansPutPreviousValues) GetExternalIntegrationUrlOk() (*string, bool)`

GetExternalIntegrationUrlOk returns a tuple with the ExternalIntegrationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIntegrationUrl

`func (o *ActionPlansPutPreviousValues) SetExternalIntegrationUrl(v string)`

SetExternalIntegrationUrl sets ExternalIntegrationUrl field to given value.

### HasExternalIntegrationUrl

`func (o *ActionPlansPutPreviousValues) HasExternalIntegrationUrl() bool`

HasExternalIntegrationUrl returns a boolean if a field has been set.

### GetIssueRemediationCustomSelect5OptionId

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect5OptionId() int32`

GetIssueRemediationCustomSelect5OptionId returns the IssueRemediationCustomSelect5OptionId field if non-nil, zero value otherwise.

### GetIssueRemediationCustomSelect5OptionIdOk

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect5OptionIdOk() (*int32, bool)`

GetIssueRemediationCustomSelect5OptionIdOk returns a tuple with the IssueRemediationCustomSelect5OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueRemediationCustomSelect5OptionId

`func (o *ActionPlansPutPreviousValues) SetIssueRemediationCustomSelect5OptionId(v int32)`

SetIssueRemediationCustomSelect5OptionId sets IssueRemediationCustomSelect5OptionId field to given value.

### HasIssueRemediationCustomSelect5OptionId

`func (o *ActionPlansPutPreviousValues) HasIssueRemediationCustomSelect5OptionId() bool`

HasIssueRemediationCustomSelect5OptionId returns a boolean if a field has been set.

### GetIssueRemediationCustomSelect6OptionId

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect6OptionId() int32`

GetIssueRemediationCustomSelect6OptionId returns the IssueRemediationCustomSelect6OptionId field if non-nil, zero value otherwise.

### GetIssueRemediationCustomSelect6OptionIdOk

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect6OptionIdOk() (*int32, bool)`

GetIssueRemediationCustomSelect6OptionIdOk returns a tuple with the IssueRemediationCustomSelect6OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueRemediationCustomSelect6OptionId

`func (o *ActionPlansPutPreviousValues) SetIssueRemediationCustomSelect6OptionId(v int32)`

SetIssueRemediationCustomSelect6OptionId sets IssueRemediationCustomSelect6OptionId field to given value.

### HasIssueRemediationCustomSelect6OptionId

`func (o *ActionPlansPutPreviousValues) HasIssueRemediationCustomSelect6OptionId() bool`

HasIssueRemediationCustomSelect6OptionId returns a boolean if a field has been set.

### GetIssueRemediationCustomSelect7OptionId

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect7OptionId() int32`

GetIssueRemediationCustomSelect7OptionId returns the IssueRemediationCustomSelect7OptionId field if non-nil, zero value otherwise.

### GetIssueRemediationCustomSelect7OptionIdOk

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect7OptionIdOk() (*int32, bool)`

GetIssueRemediationCustomSelect7OptionIdOk returns a tuple with the IssueRemediationCustomSelect7OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueRemediationCustomSelect7OptionId

`func (o *ActionPlansPutPreviousValues) SetIssueRemediationCustomSelect7OptionId(v int32)`

SetIssueRemediationCustomSelect7OptionId sets IssueRemediationCustomSelect7OptionId field to given value.

### HasIssueRemediationCustomSelect7OptionId

`func (o *ActionPlansPutPreviousValues) HasIssueRemediationCustomSelect7OptionId() bool`

HasIssueRemediationCustomSelect7OptionId returns a boolean if a field has been set.

### GetIssueRemediationCustomSelect8OptionId

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect8OptionId() int32`

GetIssueRemediationCustomSelect8OptionId returns the IssueRemediationCustomSelect8OptionId field if non-nil, zero value otherwise.

### GetIssueRemediationCustomSelect8OptionIdOk

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect8OptionIdOk() (*int32, bool)`

GetIssueRemediationCustomSelect8OptionIdOk returns a tuple with the IssueRemediationCustomSelect8OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueRemediationCustomSelect8OptionId

`func (o *ActionPlansPutPreviousValues) SetIssueRemediationCustomSelect8OptionId(v int32)`

SetIssueRemediationCustomSelect8OptionId sets IssueRemediationCustomSelect8OptionId field to given value.

### HasIssueRemediationCustomSelect8OptionId

`func (o *ActionPlansPutPreviousValues) HasIssueRemediationCustomSelect8OptionId() bool`

HasIssueRemediationCustomSelect8OptionId returns a boolean if a field has been set.

### GetCustomText4

`func (o *ActionPlansPutPreviousValues) GetCustomText4() string`

GetCustomText4 returns the CustomText4 field if non-nil, zero value otherwise.

### GetCustomText4Ok

`func (o *ActionPlansPutPreviousValues) GetCustomText4Ok() (*string, bool)`

GetCustomText4Ok returns a tuple with the CustomText4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText4

`func (o *ActionPlansPutPreviousValues) SetCustomText4(v string)`

SetCustomText4 sets CustomText4 field to given value.

### HasCustomText4

`func (o *ActionPlansPutPreviousValues) HasCustomText4() bool`

HasCustomText4 returns a boolean if a field has been set.

### GetCustomText5

`func (o *ActionPlansPutPreviousValues) GetCustomText5() string`

GetCustomText5 returns the CustomText5 field if non-nil, zero value otherwise.

### GetCustomText5Ok

`func (o *ActionPlansPutPreviousValues) GetCustomText5Ok() (*string, bool)`

GetCustomText5Ok returns a tuple with the CustomText5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText5

`func (o *ActionPlansPutPreviousValues) SetCustomText5(v string)`

SetCustomText5 sets CustomText5 field to given value.

### HasCustomText5

`func (o *ActionPlansPutPreviousValues) HasCustomText5() bool`

HasCustomText5 returns a boolean if a field has been set.

### GetCustomText6

`func (o *ActionPlansPutPreviousValues) GetCustomText6() string`

GetCustomText6 returns the CustomText6 field if non-nil, zero value otherwise.

### GetCustomText6Ok

`func (o *ActionPlansPutPreviousValues) GetCustomText6Ok() (*string, bool)`

GetCustomText6Ok returns a tuple with the CustomText6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText6

`func (o *ActionPlansPutPreviousValues) SetCustomText6(v string)`

SetCustomText6 sets CustomText6 field to given value.

### HasCustomText6

`func (o *ActionPlansPutPreviousValues) HasCustomText6() bool`

HasCustomText6 returns a boolean if a field has been set.

### GetFirstCompletedDate

`func (o *ActionPlansPutPreviousValues) GetFirstCompletedDate() string`

GetFirstCompletedDate returns the FirstCompletedDate field if non-nil, zero value otherwise.

### GetFirstCompletedDateOk

`func (o *ActionPlansPutPreviousValues) GetFirstCompletedDateOk() (*string, bool)`

GetFirstCompletedDateOk returns a tuple with the FirstCompletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCompletedDate

`func (o *ActionPlansPutPreviousValues) SetFirstCompletedDate(v string)`

SetFirstCompletedDate sets FirstCompletedDate field to given value.

### HasFirstCompletedDate

`func (o *ActionPlansPutPreviousValues) HasFirstCompletedDate() bool`

HasFirstCompletedDate returns a boolean if a field has been set.

### GetCustomCurrency1

`func (o *ActionPlansPutPreviousValues) GetCustomCurrency1() float32`

GetCustomCurrency1 returns the CustomCurrency1 field if non-nil, zero value otherwise.

### GetCustomCurrency1Ok

`func (o *ActionPlansPutPreviousValues) GetCustomCurrency1Ok() (*float32, bool)`

GetCustomCurrency1Ok returns a tuple with the CustomCurrency1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCurrency1

`func (o *ActionPlansPutPreviousValues) SetCustomCurrency1(v float32)`

SetCustomCurrency1 sets CustomCurrency1 field to given value.

### HasCustomCurrency1

`func (o *ActionPlansPutPreviousValues) HasCustomCurrency1() bool`

HasCustomCurrency1 returns a boolean if a field has been set.

### GetCustomCurrency2

`func (o *ActionPlansPutPreviousValues) GetCustomCurrency2() float32`

GetCustomCurrency2 returns the CustomCurrency2 field if non-nil, zero value otherwise.

### GetCustomCurrency2Ok

`func (o *ActionPlansPutPreviousValues) GetCustomCurrency2Ok() (*float32, bool)`

GetCustomCurrency2Ok returns a tuple with the CustomCurrency2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCurrency2

`func (o *ActionPlansPutPreviousValues) SetCustomCurrency2(v float32)`

SetCustomCurrency2 sets CustomCurrency2 field to given value.

### HasCustomCurrency2

`func (o *ActionPlansPutPreviousValues) HasCustomCurrency2() bool`

HasCustomCurrency2 returns a boolean if a field has been set.

### GetCustomCurrency3

`func (o *ActionPlansPutPreviousValues) GetCustomCurrency3() float32`

GetCustomCurrency3 returns the CustomCurrency3 field if non-nil, zero value otherwise.

### GetCustomCurrency3Ok

`func (o *ActionPlansPutPreviousValues) GetCustomCurrency3Ok() (*float32, bool)`

GetCustomCurrency3Ok returns a tuple with the CustomCurrency3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCurrency3

`func (o *ActionPlansPutPreviousValues) SetCustomCurrency3(v float32)`

SetCustomCurrency3 sets CustomCurrency3 field to given value.

### HasCustomCurrency3

`func (o *ActionPlansPutPreviousValues) HasCustomCurrency3() bool`

HasCustomCurrency3 returns a boolean if a field has been set.

### GetCustomCurrency4

`func (o *ActionPlansPutPreviousValues) GetCustomCurrency4() float32`

GetCustomCurrency4 returns the CustomCurrency4 field if non-nil, zero value otherwise.

### GetCustomCurrency4Ok

`func (o *ActionPlansPutPreviousValues) GetCustomCurrency4Ok() (*float32, bool)`

GetCustomCurrency4Ok returns a tuple with the CustomCurrency4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCurrency4

`func (o *ActionPlansPutPreviousValues) SetCustomCurrency4(v float32)`

SetCustomCurrency4 sets CustomCurrency4 field to given value.

### HasCustomCurrency4

`func (o *ActionPlansPutPreviousValues) HasCustomCurrency4() bool`

HasCustomCurrency4 returns a boolean if a field has been set.

### GetCustomText7

`func (o *ActionPlansPutPreviousValues) GetCustomText7() string`

GetCustomText7 returns the CustomText7 field if non-nil, zero value otherwise.

### GetCustomText7Ok

`func (o *ActionPlansPutPreviousValues) GetCustomText7Ok() (*string, bool)`

GetCustomText7Ok returns a tuple with the CustomText7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText7

`func (o *ActionPlansPutPreviousValues) SetCustomText7(v string)`

SetCustomText7 sets CustomText7 field to given value.

### HasCustomText7

`func (o *ActionPlansPutPreviousValues) HasCustomText7() bool`

HasCustomText7 returns a boolean if a field has been set.

### GetCustomText8

`func (o *ActionPlansPutPreviousValues) GetCustomText8() string`

GetCustomText8 returns the CustomText8 field if non-nil, zero value otherwise.

### GetCustomText8Ok

`func (o *ActionPlansPutPreviousValues) GetCustomText8Ok() (*string, bool)`

GetCustomText8Ok returns a tuple with the CustomText8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText8

`func (o *ActionPlansPutPreviousValues) SetCustomText8(v string)`

SetCustomText8 sets CustomText8 field to given value.

### HasCustomText8

`func (o *ActionPlansPutPreviousValues) HasCustomText8() bool`

HasCustomText8 returns a boolean if a field has been set.

### GetCustomText9

`func (o *ActionPlansPutPreviousValues) GetCustomText9() string`

GetCustomText9 returns the CustomText9 field if non-nil, zero value otherwise.

### GetCustomText9Ok

`func (o *ActionPlansPutPreviousValues) GetCustomText9Ok() (*string, bool)`

GetCustomText9Ok returns a tuple with the CustomText9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText9

`func (o *ActionPlansPutPreviousValues) SetCustomText9(v string)`

SetCustomText9 sets CustomText9 field to given value.

### HasCustomText9

`func (o *ActionPlansPutPreviousValues) HasCustomText9() bool`

HasCustomText9 returns a boolean if a field has been set.

### GetCustomText10

`func (o *ActionPlansPutPreviousValues) GetCustomText10() string`

GetCustomText10 returns the CustomText10 field if non-nil, zero value otherwise.

### GetCustomText10Ok

`func (o *ActionPlansPutPreviousValues) GetCustomText10Ok() (*string, bool)`

GetCustomText10Ok returns a tuple with the CustomText10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText10

`func (o *ActionPlansPutPreviousValues) SetCustomText10(v string)`

SetCustomText10 sets CustomText10 field to given value.

### HasCustomText10

`func (o *ActionPlansPutPreviousValues) HasCustomText10() bool`

HasCustomText10 returns a boolean if a field has been set.

### GetCustomText11

`func (o *ActionPlansPutPreviousValues) GetCustomText11() string`

GetCustomText11 returns the CustomText11 field if non-nil, zero value otherwise.

### GetCustomText11Ok

`func (o *ActionPlansPutPreviousValues) GetCustomText11Ok() (*string, bool)`

GetCustomText11Ok returns a tuple with the CustomText11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText11

`func (o *ActionPlansPutPreviousValues) SetCustomText11(v string)`

SetCustomText11 sets CustomText11 field to given value.

### HasCustomText11

`func (o *ActionPlansPutPreviousValues) HasCustomText11() bool`

HasCustomText11 returns a boolean if a field has been set.

### GetCustomText12

`func (o *ActionPlansPutPreviousValues) GetCustomText12() string`

GetCustomText12 returns the CustomText12 field if non-nil, zero value otherwise.

### GetCustomText12Ok

`func (o *ActionPlansPutPreviousValues) GetCustomText12Ok() (*string, bool)`

GetCustomText12Ok returns a tuple with the CustomText12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText12

`func (o *ActionPlansPutPreviousValues) SetCustomText12(v string)`

SetCustomText12 sets CustomText12 field to given value.

### HasCustomText12

`func (o *ActionPlansPutPreviousValues) HasCustomText12() bool`

HasCustomText12 returns a boolean if a field has been set.

### GetCustomText13

`func (o *ActionPlansPutPreviousValues) GetCustomText13() string`

GetCustomText13 returns the CustomText13 field if non-nil, zero value otherwise.

### GetCustomText13Ok

`func (o *ActionPlansPutPreviousValues) GetCustomText13Ok() (*string, bool)`

GetCustomText13Ok returns a tuple with the CustomText13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText13

`func (o *ActionPlansPutPreviousValues) SetCustomText13(v string)`

SetCustomText13 sets CustomText13 field to given value.

### HasCustomText13

`func (o *ActionPlansPutPreviousValues) HasCustomText13() bool`

HasCustomText13 returns a boolean if a field has been set.

### GetCustomText14

`func (o *ActionPlansPutPreviousValues) GetCustomText14() string`

GetCustomText14 returns the CustomText14 field if non-nil, zero value otherwise.

### GetCustomText14Ok

`func (o *ActionPlansPutPreviousValues) GetCustomText14Ok() (*string, bool)`

GetCustomText14Ok returns a tuple with the CustomText14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText14

`func (o *ActionPlansPutPreviousValues) SetCustomText14(v string)`

SetCustomText14 sets CustomText14 field to given value.

### HasCustomText14

`func (o *ActionPlansPutPreviousValues) HasCustomText14() bool`

HasCustomText14 returns a boolean if a field has been set.

### GetCustomText15

`func (o *ActionPlansPutPreviousValues) GetCustomText15() string`

GetCustomText15 returns the CustomText15 field if non-nil, zero value otherwise.

### GetCustomText15Ok

`func (o *ActionPlansPutPreviousValues) GetCustomText15Ok() (*string, bool)`

GetCustomText15Ok returns a tuple with the CustomText15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText15

`func (o *ActionPlansPutPreviousValues) SetCustomText15(v string)`

SetCustomText15 sets CustomText15 field to given value.

### HasCustomText15

`func (o *ActionPlansPutPreviousValues) HasCustomText15() bool`

HasCustomText15 returns a boolean if a field has been set.

### GetCustomText16

`func (o *ActionPlansPutPreviousValues) GetCustomText16() string`

GetCustomText16 returns the CustomText16 field if non-nil, zero value otherwise.

### GetCustomText16Ok

`func (o *ActionPlansPutPreviousValues) GetCustomText16Ok() (*string, bool)`

GetCustomText16Ok returns a tuple with the CustomText16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText16

`func (o *ActionPlansPutPreviousValues) SetCustomText16(v string)`

SetCustomText16 sets CustomText16 field to given value.

### HasCustomText16

`func (o *ActionPlansPutPreviousValues) HasCustomText16() bool`

HasCustomText16 returns a boolean if a field has been set.

### GetCustomText17

`func (o *ActionPlansPutPreviousValues) GetCustomText17() string`

GetCustomText17 returns the CustomText17 field if non-nil, zero value otherwise.

### GetCustomText17Ok

`func (o *ActionPlansPutPreviousValues) GetCustomText17Ok() (*string, bool)`

GetCustomText17Ok returns a tuple with the CustomText17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText17

`func (o *ActionPlansPutPreviousValues) SetCustomText17(v string)`

SetCustomText17 sets CustomText17 field to given value.

### HasCustomText17

`func (o *ActionPlansPutPreviousValues) HasCustomText17() bool`

HasCustomText17 returns a boolean if a field has been set.

### GetCustomText18

`func (o *ActionPlansPutPreviousValues) GetCustomText18() string`

GetCustomText18 returns the CustomText18 field if non-nil, zero value otherwise.

### GetCustomText18Ok

`func (o *ActionPlansPutPreviousValues) GetCustomText18Ok() (*string, bool)`

GetCustomText18Ok returns a tuple with the CustomText18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText18

`func (o *ActionPlansPutPreviousValues) SetCustomText18(v string)`

SetCustomText18 sets CustomText18 field to given value.

### HasCustomText18

`func (o *ActionPlansPutPreviousValues) HasCustomText18() bool`

HasCustomText18 returns a boolean if a field has been set.

### GetCustomText19

`func (o *ActionPlansPutPreviousValues) GetCustomText19() string`

GetCustomText19 returns the CustomText19 field if non-nil, zero value otherwise.

### GetCustomText19Ok

`func (o *ActionPlansPutPreviousValues) GetCustomText19Ok() (*string, bool)`

GetCustomText19Ok returns a tuple with the CustomText19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText19

`func (o *ActionPlansPutPreviousValues) SetCustomText19(v string)`

SetCustomText19 sets CustomText19 field to given value.

### HasCustomText19

`func (o *ActionPlansPutPreviousValues) HasCustomText19() bool`

HasCustomText19 returns a boolean if a field has been set.

### GetCustomText20

`func (o *ActionPlansPutPreviousValues) GetCustomText20() string`

GetCustomText20 returns the CustomText20 field if non-nil, zero value otherwise.

### GetCustomText20Ok

`func (o *ActionPlansPutPreviousValues) GetCustomText20Ok() (*string, bool)`

GetCustomText20Ok returns a tuple with the CustomText20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText20

`func (o *ActionPlansPutPreviousValues) SetCustomText20(v string)`

SetCustomText20 sets CustomText20 field to given value.

### HasCustomText20

`func (o *ActionPlansPutPreviousValues) HasCustomText20() bool`

HasCustomText20 returns a boolean if a field has been set.

### GetCustomText21

`func (o *ActionPlansPutPreviousValues) GetCustomText21() string`

GetCustomText21 returns the CustomText21 field if non-nil, zero value otherwise.

### GetCustomText21Ok

`func (o *ActionPlansPutPreviousValues) GetCustomText21Ok() (*string, bool)`

GetCustomText21Ok returns a tuple with the CustomText21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText21

`func (o *ActionPlansPutPreviousValues) SetCustomText21(v string)`

SetCustomText21 sets CustomText21 field to given value.

### HasCustomText21

`func (o *ActionPlansPutPreviousValues) HasCustomText21() bool`

HasCustomText21 returns a boolean if a field has been set.

### GetIssueRemediationCustomSelect9OptionId

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect9OptionId() int32`

GetIssueRemediationCustomSelect9OptionId returns the IssueRemediationCustomSelect9OptionId field if non-nil, zero value otherwise.

### GetIssueRemediationCustomSelect9OptionIdOk

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect9OptionIdOk() (*int32, bool)`

GetIssueRemediationCustomSelect9OptionIdOk returns a tuple with the IssueRemediationCustomSelect9OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueRemediationCustomSelect9OptionId

`func (o *ActionPlansPutPreviousValues) SetIssueRemediationCustomSelect9OptionId(v int32)`

SetIssueRemediationCustomSelect9OptionId sets IssueRemediationCustomSelect9OptionId field to given value.

### HasIssueRemediationCustomSelect9OptionId

`func (o *ActionPlansPutPreviousValues) HasIssueRemediationCustomSelect9OptionId() bool`

HasIssueRemediationCustomSelect9OptionId returns a boolean if a field has been set.

### GetIssueRemediationCustomSelect10OptionId

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect10OptionId() int32`

GetIssueRemediationCustomSelect10OptionId returns the IssueRemediationCustomSelect10OptionId field if non-nil, zero value otherwise.

### GetIssueRemediationCustomSelect10OptionIdOk

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect10OptionIdOk() (*int32, bool)`

GetIssueRemediationCustomSelect10OptionIdOk returns a tuple with the IssueRemediationCustomSelect10OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueRemediationCustomSelect10OptionId

`func (o *ActionPlansPutPreviousValues) SetIssueRemediationCustomSelect10OptionId(v int32)`

SetIssueRemediationCustomSelect10OptionId sets IssueRemediationCustomSelect10OptionId field to given value.

### HasIssueRemediationCustomSelect10OptionId

`func (o *ActionPlansPutPreviousValues) HasIssueRemediationCustomSelect10OptionId() bool`

HasIssueRemediationCustomSelect10OptionId returns a boolean if a field has been set.

### GetIssueRemediationCustomSelect11OptionId

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect11OptionId() int32`

GetIssueRemediationCustomSelect11OptionId returns the IssueRemediationCustomSelect11OptionId field if non-nil, zero value otherwise.

### GetIssueRemediationCustomSelect11OptionIdOk

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect11OptionIdOk() (*int32, bool)`

GetIssueRemediationCustomSelect11OptionIdOk returns a tuple with the IssueRemediationCustomSelect11OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueRemediationCustomSelect11OptionId

`func (o *ActionPlansPutPreviousValues) SetIssueRemediationCustomSelect11OptionId(v int32)`

SetIssueRemediationCustomSelect11OptionId sets IssueRemediationCustomSelect11OptionId field to given value.

### HasIssueRemediationCustomSelect11OptionId

`func (o *ActionPlansPutPreviousValues) HasIssueRemediationCustomSelect11OptionId() bool`

HasIssueRemediationCustomSelect11OptionId returns a boolean if a field has been set.

### GetIssueRemediationCustomSelect12OptionId

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect12OptionId() int32`

GetIssueRemediationCustomSelect12OptionId returns the IssueRemediationCustomSelect12OptionId field if non-nil, zero value otherwise.

### GetIssueRemediationCustomSelect12OptionIdOk

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect12OptionIdOk() (*int32, bool)`

GetIssueRemediationCustomSelect12OptionIdOk returns a tuple with the IssueRemediationCustomSelect12OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueRemediationCustomSelect12OptionId

`func (o *ActionPlansPutPreviousValues) SetIssueRemediationCustomSelect12OptionId(v int32)`

SetIssueRemediationCustomSelect12OptionId sets IssueRemediationCustomSelect12OptionId field to given value.

### HasIssueRemediationCustomSelect12OptionId

`func (o *ActionPlansPutPreviousValues) HasIssueRemediationCustomSelect12OptionId() bool`

HasIssueRemediationCustomSelect12OptionId returns a boolean if a field has been set.

### GetIssueRemediationCustomSelect13OptionId

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect13OptionId() int32`

GetIssueRemediationCustomSelect13OptionId returns the IssueRemediationCustomSelect13OptionId field if non-nil, zero value otherwise.

### GetIssueRemediationCustomSelect13OptionIdOk

`func (o *ActionPlansPutPreviousValues) GetIssueRemediationCustomSelect13OptionIdOk() (*int32, bool)`

GetIssueRemediationCustomSelect13OptionIdOk returns a tuple with the IssueRemediationCustomSelect13OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueRemediationCustomSelect13OptionId

`func (o *ActionPlansPutPreviousValues) SetIssueRemediationCustomSelect13OptionId(v int32)`

SetIssueRemediationCustomSelect13OptionId sets IssueRemediationCustomSelect13OptionId field to given value.

### HasIssueRemediationCustomSelect13OptionId

`func (o *ActionPlansPutPreviousValues) HasIssueRemediationCustomSelect13OptionId() bool`

HasIssueRemediationCustomSelect13OptionId returns a boolean if a field has been set.

### GetFieldData

`func (o *ActionPlansPutPreviousValues) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *ActionPlansPutPreviousValues) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *ActionPlansPutPreviousValues) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *ActionPlansPutPreviousValues) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *ActionPlansPutPreviousValues) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *ActionPlansPutPreviousValues) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil
### GetCompleteReason

`func (o *ActionPlansPutPreviousValues) GetCompleteReason() string`

GetCompleteReason returns the CompleteReason field if non-nil, zero value otherwise.

### GetCompleteReasonOk

`func (o *ActionPlansPutPreviousValues) GetCompleteReasonOk() (*string, bool)`

GetCompleteReasonOk returns a tuple with the CompleteReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteReason

`func (o *ActionPlansPutPreviousValues) SetCompleteReason(v string)`

SetCompleteReason sets CompleteReason field to given value.

### HasCompleteReason

`func (o *ActionPlansPutPreviousValues) HasCompleteReason() bool`

HasCompleteReason returns a boolean if a field has been set.

### GetScopes

`func (o *ActionPlansPutPreviousValues) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ActionPlansPutPreviousValues) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ActionPlansPutPreviousValues) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ActionPlansPutPreviousValues) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *ActionPlansPutPreviousValues) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *ActionPlansPutPreviousValues) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetComplianceAssessmentId

`func (o *ActionPlansPutPreviousValues) GetComplianceAssessmentId() int32`

GetComplianceAssessmentId returns the ComplianceAssessmentId field if non-nil, zero value otherwise.

### GetComplianceAssessmentIdOk

`func (o *ActionPlansPutPreviousValues) GetComplianceAssessmentIdOk() (*int32, bool)`

GetComplianceAssessmentIdOk returns a tuple with the ComplianceAssessmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceAssessmentId

`func (o *ActionPlansPutPreviousValues) SetComplianceAssessmentId(v int32)`

SetComplianceAssessmentId sets ComplianceAssessmentId field to given value.

### HasComplianceAssessmentId

`func (o *ActionPlansPutPreviousValues) HasComplianceAssessmentId() bool`

HasComplianceAssessmentId returns a boolean if a field has been set.

### GetComplianceAssessmentItemId

`func (o *ActionPlansPutPreviousValues) GetComplianceAssessmentItemId() int32`

GetComplianceAssessmentItemId returns the ComplianceAssessmentItemId field if non-nil, zero value otherwise.

### GetComplianceAssessmentItemIdOk

`func (o *ActionPlansPutPreviousValues) GetComplianceAssessmentItemIdOk() (*int32, bool)`

GetComplianceAssessmentItemIdOk returns a tuple with the ComplianceAssessmentItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceAssessmentItemId

`func (o *ActionPlansPutPreviousValues) SetComplianceAssessmentItemId(v int32)`

SetComplianceAssessmentItemId sets ComplianceAssessmentItemId field to given value.

### HasComplianceAssessmentItemId

`func (o *ActionPlansPutPreviousValues) HasComplianceAssessmentItemId() bool`

HasComplianceAssessmentItemId returns a boolean if a field has been set.

### GetFirstPendingRemediationDueDate

`func (o *ActionPlansPutPreviousValues) GetFirstPendingRemediationDueDate() string`

GetFirstPendingRemediationDueDate returns the FirstPendingRemediationDueDate field if non-nil, zero value otherwise.

### GetFirstPendingRemediationDueDateOk

`func (o *ActionPlansPutPreviousValues) GetFirstPendingRemediationDueDateOk() (*string, bool)`

GetFirstPendingRemediationDueDateOk returns a tuple with the FirstPendingRemediationDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPendingRemediationDueDate

`func (o *ActionPlansPutPreviousValues) SetFirstPendingRemediationDueDate(v string)`

SetFirstPendingRemediationDueDate sets FirstPendingRemediationDueDate field to given value.

### HasFirstPendingRemediationDueDate

`func (o *ActionPlansPutPreviousValues) HasFirstPendingRemediationDueDate() bool`

HasFirstPendingRemediationDueDate returns a boolean if a field has been set.

### GetPendingRemediationDueDateCount

`func (o *ActionPlansPutPreviousValues) GetPendingRemediationDueDateCount() int32`

GetPendingRemediationDueDateCount returns the PendingRemediationDueDateCount field if non-nil, zero value otherwise.

### GetPendingRemediationDueDateCountOk

`func (o *ActionPlansPutPreviousValues) GetPendingRemediationDueDateCountOk() (*int32, bool)`

GetPendingRemediationDueDateCountOk returns a tuple with the PendingRemediationDueDateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingRemediationDueDateCount

`func (o *ActionPlansPutPreviousValues) SetPendingRemediationDueDateCount(v int32)`

SetPendingRemediationDueDateCount sets PendingRemediationDueDateCount field to given value.

### HasPendingRemediationDueDateCount

`func (o *ActionPlansPutPreviousValues) HasPendingRemediationDueDateCount() bool`

HasPendingRemediationDueDateCount returns a boolean if a field has been set.

### GetFirstOpenDueDate

`func (o *ActionPlansPutPreviousValues) GetFirstOpenDueDate() string`

GetFirstOpenDueDate returns the FirstOpenDueDate field if non-nil, zero value otherwise.

### GetFirstOpenDueDateOk

`func (o *ActionPlansPutPreviousValues) GetFirstOpenDueDateOk() (*string, bool)`

GetFirstOpenDueDateOk returns a tuple with the FirstOpenDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstOpenDueDate

`func (o *ActionPlansPutPreviousValues) SetFirstOpenDueDate(v string)`

SetFirstOpenDueDate sets FirstOpenDueDate field to given value.

### HasFirstOpenDueDate

`func (o *ActionPlansPutPreviousValues) HasFirstOpenDueDate() bool`

HasFirstOpenDueDate returns a boolean if a field has been set.

### GetName

`func (o *ActionPlansPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionPlansPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionPlansPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActionPlansPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


