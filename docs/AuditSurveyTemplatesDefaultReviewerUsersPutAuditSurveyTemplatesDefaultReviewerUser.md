# AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**AuditSurveyTemplateId** | **int32** | Note: This is a Foreign Key to &#x60;audit_survey_templates.id&#x60;.&lt;fk table&#x3D;&#39;audit_survey_templates&#39; column&#x3D;&#39;id&#39;/&gt; | 
**UserId** | **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | 
**SortOrder** | **int32** |  | [default to 0]

## Methods

### NewAuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser

`func NewAuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser(auditSurveyTemplateId int32, userId int32, sortOrder int32, ) *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser`

NewAuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser instantiates a new AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUserWithDefaults

`func NewAuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUserWithDefaults() *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser`

NewAuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUserWithDefaults instantiates a new AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetAuditSurveyTemplateId

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) GetAuditSurveyTemplateId() int32`

GetAuditSurveyTemplateId returns the AuditSurveyTemplateId field if non-nil, zero value otherwise.

### GetAuditSurveyTemplateIdOk

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) GetAuditSurveyTemplateIdOk() (*int32, bool)`

GetAuditSurveyTemplateIdOk returns a tuple with the AuditSurveyTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditSurveyTemplateId

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) SetAuditSurveyTemplateId(v int32)`

SetAuditSurveyTemplateId sets AuditSurveyTemplateId field to given value.


### GetUserId

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetSortOrder

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AuditSurveyTemplatesDefaultReviewerUsersPutAuditSurveyTemplatesDefaultReviewerUser) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


