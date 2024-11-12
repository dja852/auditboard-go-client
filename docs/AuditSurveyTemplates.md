# AuditSurveyTemplates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**CreatedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OpsAuditCategoryId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_categories.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_categories&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Type** | **string** |  | 
**Key** | Pointer to **string** |  | [optional] 
**SourceSurveyId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_surveys.id&#x60;.&lt;fk table&#x3D;&#39;audit_surveys&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AllowAuditableEntityOwnersViewResponses** | **bool** |  | [default to false]
**EsgMetricCategoryId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;esg_metric_categories.id&#x60;.&lt;fk table&#x3D;&#39;esg_metric_categories&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**MultiReviewerMode** | **string** |  | [default to "nonsequential"]
**RequiredReviewerCount** | Pointer to **int32** |  | [optional] 
**DefaultPreparerType** | Pointer to **string** |  | [optional] 
**IsFixed** | **bool** |  | [default to false]

## Methods

### NewAuditSurveyTemplates

`func NewAuditSurveyTemplates(type_ string, allowAuditableEntityOwnersViewResponses bool, multiReviewerMode string, isFixed bool, ) *AuditSurveyTemplates`

NewAuditSurveyTemplates instantiates a new AuditSurveyTemplates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditSurveyTemplatesWithDefaults

`func NewAuditSurveyTemplatesWithDefaults() *AuditSurveyTemplates`

NewAuditSurveyTemplatesWithDefaults instantiates a new AuditSurveyTemplates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditSurveyTemplates) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditSurveyTemplates) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditSurveyTemplates) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuditSurveyTemplates) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuditSurveyTemplates) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditSurveyTemplates) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditSurveyTemplates) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuditSurveyTemplates) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuditSurveyTemplates) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuditSurveyTemplates) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuditSurveyTemplates) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuditSurveyTemplates) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AuditSurveyTemplates) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AuditSurveyTemplates) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AuditSurveyTemplates) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AuditSurveyTemplates) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *AuditSurveyTemplates) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *AuditSurveyTemplates) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *AuditSurveyTemplates) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *AuditSurveyTemplates) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetName

`func (o *AuditSurveyTemplates) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditSurveyTemplates) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditSurveyTemplates) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuditSurveyTemplates) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpsAuditCategoryId

`func (o *AuditSurveyTemplates) GetOpsAuditCategoryId() int32`

GetOpsAuditCategoryId returns the OpsAuditCategoryId field if non-nil, zero value otherwise.

### GetOpsAuditCategoryIdOk

`func (o *AuditSurveyTemplates) GetOpsAuditCategoryIdOk() (*int32, bool)`

GetOpsAuditCategoryIdOk returns a tuple with the OpsAuditCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditCategoryId

`func (o *AuditSurveyTemplates) SetOpsAuditCategoryId(v int32)`

SetOpsAuditCategoryId sets OpsAuditCategoryId field to given value.

### HasOpsAuditCategoryId

`func (o *AuditSurveyTemplates) HasOpsAuditCategoryId() bool`

HasOpsAuditCategoryId returns a boolean if a field has been set.

### GetType

`func (o *AuditSurveyTemplates) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditSurveyTemplates) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditSurveyTemplates) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *AuditSurveyTemplates) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AuditSurveyTemplates) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AuditSurveyTemplates) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AuditSurveyTemplates) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSourceSurveyId

`func (o *AuditSurveyTemplates) GetSourceSurveyId() int32`

GetSourceSurveyId returns the SourceSurveyId field if non-nil, zero value otherwise.

### GetSourceSurveyIdOk

`func (o *AuditSurveyTemplates) GetSourceSurveyIdOk() (*int32, bool)`

GetSourceSurveyIdOk returns a tuple with the SourceSurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSurveyId

`func (o *AuditSurveyTemplates) SetSourceSurveyId(v int32)`

SetSourceSurveyId sets SourceSurveyId field to given value.

### HasSourceSurveyId

`func (o *AuditSurveyTemplates) HasSourceSurveyId() bool`

HasSourceSurveyId returns a boolean if a field has been set.

### GetAllowAuditableEntityOwnersViewResponses

`func (o *AuditSurveyTemplates) GetAllowAuditableEntityOwnersViewResponses() bool`

GetAllowAuditableEntityOwnersViewResponses returns the AllowAuditableEntityOwnersViewResponses field if non-nil, zero value otherwise.

### GetAllowAuditableEntityOwnersViewResponsesOk

`func (o *AuditSurveyTemplates) GetAllowAuditableEntityOwnersViewResponsesOk() (*bool, bool)`

GetAllowAuditableEntityOwnersViewResponsesOk returns a tuple with the AllowAuditableEntityOwnersViewResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAuditableEntityOwnersViewResponses

`func (o *AuditSurveyTemplates) SetAllowAuditableEntityOwnersViewResponses(v bool)`

SetAllowAuditableEntityOwnersViewResponses sets AllowAuditableEntityOwnersViewResponses field to given value.


### GetEsgMetricCategoryId

`func (o *AuditSurveyTemplates) GetEsgMetricCategoryId() int32`

GetEsgMetricCategoryId returns the EsgMetricCategoryId field if non-nil, zero value otherwise.

### GetEsgMetricCategoryIdOk

`func (o *AuditSurveyTemplates) GetEsgMetricCategoryIdOk() (*int32, bool)`

GetEsgMetricCategoryIdOk returns a tuple with the EsgMetricCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricCategoryId

`func (o *AuditSurveyTemplates) SetEsgMetricCategoryId(v int32)`

SetEsgMetricCategoryId sets EsgMetricCategoryId field to given value.

### HasEsgMetricCategoryId

`func (o *AuditSurveyTemplates) HasEsgMetricCategoryId() bool`

HasEsgMetricCategoryId returns a boolean if a field has been set.

### GetMultiReviewerMode

`func (o *AuditSurveyTemplates) GetMultiReviewerMode() string`

GetMultiReviewerMode returns the MultiReviewerMode field if non-nil, zero value otherwise.

### GetMultiReviewerModeOk

`func (o *AuditSurveyTemplates) GetMultiReviewerModeOk() (*string, bool)`

GetMultiReviewerModeOk returns a tuple with the MultiReviewerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiReviewerMode

`func (o *AuditSurveyTemplates) SetMultiReviewerMode(v string)`

SetMultiReviewerMode sets MultiReviewerMode field to given value.


### GetRequiredReviewerCount

`func (o *AuditSurveyTemplates) GetRequiredReviewerCount() int32`

GetRequiredReviewerCount returns the RequiredReviewerCount field if non-nil, zero value otherwise.

### GetRequiredReviewerCountOk

`func (o *AuditSurveyTemplates) GetRequiredReviewerCountOk() (*int32, bool)`

GetRequiredReviewerCountOk returns a tuple with the RequiredReviewerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredReviewerCount

`func (o *AuditSurveyTemplates) SetRequiredReviewerCount(v int32)`

SetRequiredReviewerCount sets RequiredReviewerCount field to given value.

### HasRequiredReviewerCount

`func (o *AuditSurveyTemplates) HasRequiredReviewerCount() bool`

HasRequiredReviewerCount returns a boolean if a field has been set.

### GetDefaultPreparerType

`func (o *AuditSurveyTemplates) GetDefaultPreparerType() string`

GetDefaultPreparerType returns the DefaultPreparerType field if non-nil, zero value otherwise.

### GetDefaultPreparerTypeOk

`func (o *AuditSurveyTemplates) GetDefaultPreparerTypeOk() (*string, bool)`

GetDefaultPreparerTypeOk returns a tuple with the DefaultPreparerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPreparerType

`func (o *AuditSurveyTemplates) SetDefaultPreparerType(v string)`

SetDefaultPreparerType sets DefaultPreparerType field to given value.

### HasDefaultPreparerType

`func (o *AuditSurveyTemplates) HasDefaultPreparerType() bool`

HasDefaultPreparerType returns a boolean if a field has been set.

### GetIsFixed

`func (o *AuditSurveyTemplates) GetIsFixed() bool`

GetIsFixed returns the IsFixed field if non-nil, zero value otherwise.

### GetIsFixedOk

`func (o *AuditSurveyTemplates) GetIsFixedOk() (*bool, bool)`

GetIsFixedOk returns a tuple with the IsFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFixed

`func (o *AuditSurveyTemplates) SetIsFixed(v bool)`

SetIsFixed sets IsFixed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


