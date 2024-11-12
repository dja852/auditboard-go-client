# AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**AuditSurveyTemplateId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;audit_survey_templates.id&#x60;.&lt;fk table&#x3D;&#39;audit_survey_templates&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**UserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**SortOrder** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewAuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues

`func NewAuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues() *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues`

NewAuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues instantiates a new AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditSurveyTemplatesDefaultReviewerUsersPutPreviousValuesWithDefaults

`func NewAuditSurveyTemplatesDefaultReviewerUsersPutPreviousValuesWithDefaults() *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues`

NewAuditSurveyTemplatesDefaultReviewerUsersPutPreviousValuesWithDefaults instantiates a new AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetAuditSurveyTemplateId

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetAuditSurveyTemplateId() int32`

GetAuditSurveyTemplateId returns the AuditSurveyTemplateId field if non-nil, zero value otherwise.

### GetAuditSurveyTemplateIdOk

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetAuditSurveyTemplateIdOk() (*int32, bool)`

GetAuditSurveyTemplateIdOk returns a tuple with the AuditSurveyTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSurveyTemplateId

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) SetAuditSurveyTemplateId(v int32)`

SetAuditSurveyTemplateId sets AuditSurveyTemplateId field to given value.

### HasAuditSurveyTemplateId

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) HasAuditSurveyTemplateId() bool`

HasAuditSurveyTemplateId returns a boolean if a field has been set.

### GetUserId

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSortOrder

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutPreviousValues) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


