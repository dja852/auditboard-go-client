# OpsAuditSectionInstancesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**SortOrder** | Pointer to **int32** |  | [optional] [default to 0]
**OpsAuditId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audits.id&#x60;.&lt;fk table&#x3D;&#39;ops_audits&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**OpsAuditSectionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;ops_audit_sections.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_sections&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FormTemplateId** | Pointer to **int32** |  | [optional] [default to 1]

## Methods

### NewOpsAuditSectionInstancesPutPreviousValues

`func NewOpsAuditSectionInstancesPutPreviousValues() *OpsAuditSectionInstancesPutPreviousValues`

NewOpsAuditSectionInstancesPutPreviousValues instantiates a new OpsAuditSectionInstancesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsAuditSectionInstancesPutPreviousValuesWithDefaults

`func NewOpsAuditSectionInstancesPutPreviousValuesWithDefaults() *OpsAuditSectionInstancesPutPreviousValues`

NewOpsAuditSectionInstancesPutPreviousValuesWithDefaults instantiates a new OpsAuditSectionInstancesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpsAuditSectionInstancesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpsAuditSectionInstancesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpsAuditSectionInstancesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OpsAuditSectionInstancesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OpsAuditSectionInstancesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OpsAuditSectionInstancesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OpsAuditSectionInstancesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OpsAuditSectionInstancesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OpsAuditSectionInstancesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OpsAuditSectionInstancesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OpsAuditSectionInstancesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OpsAuditSectionInstancesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *OpsAuditSectionInstancesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *OpsAuditSectionInstancesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *OpsAuditSectionInstancesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *OpsAuditSectionInstancesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetSortOrder

`func (o *OpsAuditSectionInstancesPutPreviousValues) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *OpsAuditSectionInstancesPutPreviousValues) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *OpsAuditSectionInstancesPutPreviousValues) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *OpsAuditSectionInstancesPutPreviousValues) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetOpsAuditId

`func (o *OpsAuditSectionInstancesPutPreviousValues) GetOpsAuditId() int32`

GetOpsAuditId returns the OpsAuditId field if non-nil, zero value otherwise.

### GetOpsAuditIdOk

`func (o *OpsAuditSectionInstancesPutPreviousValues) GetOpsAuditIdOk() (*int32, bool)`

GetOpsAuditIdOk returns a tuple with the OpsAuditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditId

`func (o *OpsAuditSectionInstancesPutPreviousValues) SetOpsAuditId(v int32)`

SetOpsAuditId sets OpsAuditId field to given value.

### HasOpsAuditId

`func (o *OpsAuditSectionInstancesPutPreviousValues) HasOpsAuditId() bool`

HasOpsAuditId returns a boolean if a field has been set.

### GetOpsAuditSectionId

`func (o *OpsAuditSectionInstancesPutPreviousValues) GetOpsAuditSectionId() int32`

GetOpsAuditSectionId returns the OpsAuditSectionId field if non-nil, zero value otherwise.

### GetOpsAuditSectionIdOk

`func (o *OpsAuditSectionInstancesPutPreviousValues) GetOpsAuditSectionIdOk() (*int32, bool)`

GetOpsAuditSectionIdOk returns a tuple with the OpsAuditSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditSectionId

`func (o *OpsAuditSectionInstancesPutPreviousValues) SetOpsAuditSectionId(v int32)`

SetOpsAuditSectionId sets OpsAuditSectionId field to given value.

### HasOpsAuditSectionId

`func (o *OpsAuditSectionInstancesPutPreviousValues) HasOpsAuditSectionId() bool`

HasOpsAuditSectionId returns a boolean if a field has been set.

### GetFormTemplateId

`func (o *OpsAuditSectionInstancesPutPreviousValues) GetFormTemplateId() int32`

GetFormTemplateId returns the FormTemplateId field if non-nil, zero value otherwise.

### GetFormTemplateIdOk

`func (o *OpsAuditSectionInstancesPutPreviousValues) GetFormTemplateIdOk() (*int32, bool)`

GetFormTemplateIdOk returns a tuple with the FormTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormTemplateId

`func (o *OpsAuditSectionInstancesPutPreviousValues) SetFormTemplateId(v int32)`

SetFormTemplateId sets FormTemplateId field to given value.

### HasFormTemplateId

`func (o *OpsAuditSectionInstancesPutPreviousValues) HasFormTemplateId() bool`

HasFormTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


