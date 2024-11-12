# AuditSurveyTemplatesPutAuditSurveyTemplate

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

### NewAuditSurveyTemplatesPutAuditSurveyTemplate

`func NewAuditSurveyTemplatesPutAuditSurveyTemplate(type_ string, allowAuditableEntityOwnersViewResponses bool, multiReviewerMode string, isFixed bool, ) *AuditSurveyTemplatesPutAuditSurveyTemplate`

NewAuditSurveyTemplatesPutAuditSurveyTemplate instantiates a new AuditSurveyTemplatesPutAuditSurveyTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditSurveyTemplatesPutAuditSurveyTemplateWithDefaults

`func NewAuditSurveyTemplatesPutAuditSurveyTemplateWithDefaults() *AuditSurveyTemplatesPutAuditSurveyTemplate`

NewAuditSurveyTemplatesPutAuditSurveyTemplateWithDefaults instantiates a new AuditSurveyTemplatesPutAuditSurveyTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetName

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpsAuditCategoryId

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetOpsAuditCategoryId() int32`

GetOpsAuditCategoryId returns the OpsAuditCategoryId field if non-nil, zero value otherwise.

### GetOpsAuditCategoryIdOk

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetOpsAuditCategoryIdOk() (*int32, bool)`

GetOpsAuditCategoryIdOk returns a tuple with the OpsAuditCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditCategoryId

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) SetOpsAuditCategoryId(v int32)`

SetOpsAuditCategoryId sets OpsAuditCategoryId field to given value.

### HasOpsAuditCategoryId

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) HasOpsAuditCategoryId() bool`

HasOpsAuditCategoryId returns a boolean if a field has been set.

### GetType

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSourceSurveyId

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetSourceSurveyId() int32`

GetSourceSurveyId returns the SourceSurveyId field if non-nil, zero value otherwise.

### GetSourceSurveyIdOk

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetSourceSurveyIdOk() (*int32, bool)`

GetSourceSurveyIdOk returns a tuple with the SourceSurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSurveyId

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) SetSourceSurveyId(v int32)`

SetSourceSurveyId sets SourceSurveyId field to given value.

### HasSourceSurveyId

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) HasSourceSurveyId() bool`

HasSourceSurveyId returns a boolean if a field has been set.

### GetAllowAuditableEntityOwnersViewResponses

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetAllowAuditableEntityOwnersViewResponses() bool`

GetAllowAuditableEntityOwnersViewResponses returns the AllowAuditableEntityOwnersViewResponses field if non-nil, zero value otherwise.

### GetAllowAuditableEntityOwnersViewResponsesOk

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetAllowAuditableEntityOwnersViewResponsesOk() (*bool, bool)`

GetAllowAuditableEntityOwnersViewResponsesOk returns a tuple with the AllowAuditableEntityOwnersViewResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAuditableEntityOwnersViewResponses

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) SetAllowAuditableEntityOwnersViewResponses(v bool)`

SetAllowAuditableEntityOwnersViewResponses sets AllowAuditableEntityOwnersViewResponses field to given value.


### GetEsgMetricCategoryId

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetEsgMetricCategoryId() int32`

GetEsgMetricCategoryId returns the EsgMetricCategoryId field if non-nil, zero value otherwise.

### GetEsgMetricCategoryIdOk

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetEsgMetricCategoryIdOk() (*int32, bool)`

GetEsgMetricCategoryIdOk returns a tuple with the EsgMetricCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsgMetricCategoryId

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) SetEsgMetricCategoryId(v int32)`

SetEsgMetricCategoryId sets EsgMetricCategoryId field to given value.

### HasEsgMetricCategoryId

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) HasEsgMetricCategoryId() bool`

HasEsgMetricCategoryId returns a boolean if a field has been set.

### GetMultiReviewerMode

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetMultiReviewerMode() string`

GetMultiReviewerMode returns the MultiReviewerMode field if non-nil, zero value otherwise.

### GetMultiReviewerModeOk

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetMultiReviewerModeOk() (*string, bool)`

GetMultiReviewerModeOk returns a tuple with the MultiReviewerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiReviewerMode

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) SetMultiReviewerMode(v string)`

SetMultiReviewerMode sets MultiReviewerMode field to given value.


### GetRequiredReviewerCount

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetRequiredReviewerCount() int32`

GetRequiredReviewerCount returns the RequiredReviewerCount field if non-nil, zero value otherwise.

### GetRequiredReviewerCountOk

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetRequiredReviewerCountOk() (*int32, bool)`

GetRequiredReviewerCountOk returns a tuple with the RequiredReviewerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredReviewerCount

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) SetRequiredReviewerCount(v int32)`

SetRequiredReviewerCount sets RequiredReviewerCount field to given value.

### HasRequiredReviewerCount

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) HasRequiredReviewerCount() bool`

HasRequiredReviewerCount returns a boolean if a field has been set.

### GetDefaultPreparerType

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetDefaultPreparerType() string`

GetDefaultPreparerType returns the DefaultPreparerType field if non-nil, zero value otherwise.

### GetDefaultPreparerTypeOk

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetDefaultPreparerTypeOk() (*string, bool)`

GetDefaultPreparerTypeOk returns a tuple with the DefaultPreparerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPreparerType

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) SetDefaultPreparerType(v string)`

SetDefaultPreparerType sets DefaultPreparerType field to given value.

### HasDefaultPreparerType

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) HasDefaultPreparerType() bool`

HasDefaultPreparerType returns a boolean if a field has been set.

### GetIsFixed

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetIsFixed() bool`

GetIsFixed returns the IsFixed field if non-nil, zero value otherwise.

### GetIsFixedOk

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) GetIsFixedOk() (*bool, bool)`

GetIsFixedOk returns a tuple with the IsFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFixed

`func (o *AuditSurveyTemplatesPutAuditSurveyTemplate) SetIsFixed(v bool)`

SetIsFixed sets IsFixed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


