# OpsAuditSectionInstances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**SortOrder** | **int32** |  | [default to 0]
**OpsAuditId** | **int32** | Note: This is a Foreign Key to &#x60;ops_audits.id&#x60;.&lt;fk table&#x3D;&#39;ops_audits&#39; column&#x3D;&#39;id&#39;/&gt; | 
**OpsAuditSectionId** | **int32** | Note: This is a Foreign Key to &#x60;ops_audit_sections.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_sections&#39; column&#x3D;&#39;id&#39;/&gt; | 
**FormTemplateId** | Pointer to **int32** |  | [optional] [default to 1]

## Methods

### NewOpsAuditSectionInstances

`func NewOpsAuditSectionInstances(sortOrder int32, opsAuditId int32, opsAuditSectionId int32, ) *OpsAuditSectionInstances`

NewOpsAuditSectionInstances instantiates a new OpsAuditSectionInstances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsAuditSectionInstancesWithDefaults

`func NewOpsAuditSectionInstancesWithDefaults() *OpsAuditSectionInstances`

NewOpsAuditSectionInstancesWithDefaults instantiates a new OpsAuditSectionInstances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpsAuditSectionInstances) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpsAuditSectionInstances) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpsAuditSectionInstances) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OpsAuditSectionInstances) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OpsAuditSectionInstances) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OpsAuditSectionInstances) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OpsAuditSectionInstances) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OpsAuditSectionInstances) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OpsAuditSectionInstances) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OpsAuditSectionInstances) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OpsAuditSectionInstances) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OpsAuditSectionInstances) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *OpsAuditSectionInstances) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *OpsAuditSectionInstances) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *OpsAuditSectionInstances) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *OpsAuditSectionInstances) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetSortOrder

`func (o *OpsAuditSectionInstances) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *OpsAuditSectionInstances) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *OpsAuditSectionInstances) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.


### GetOpsAuditId

`func (o *OpsAuditSectionInstances) GetOpsAuditId() int32`

GetOpsAuditId returns the OpsAuditId field if non-nil, zero value otherwise.

### GetOpsAuditIdOk

`func (o *OpsAuditSectionInstances) GetOpsAuditIdOk() (*int32, bool)`

GetOpsAuditIdOk returns a tuple with the OpsAuditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditId

`func (o *OpsAuditSectionInstances) SetOpsAuditId(v int32)`

SetOpsAuditId sets OpsAuditId field to given value.


### GetOpsAuditSectionId

`func (o *OpsAuditSectionInstances) GetOpsAuditSectionId() int32`

GetOpsAuditSectionId returns the OpsAuditSectionId field if non-nil, zero value otherwise.

### GetOpsAuditSectionIdOk

`func (o *OpsAuditSectionInstances) GetOpsAuditSectionIdOk() (*int32, bool)`

GetOpsAuditSectionIdOk returns a tuple with the OpsAuditSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditSectionId

`func (o *OpsAuditSectionInstances) SetOpsAuditSectionId(v int32)`

SetOpsAuditSectionId sets OpsAuditSectionId field to given value.


### GetFormTemplateId

`func (o *OpsAuditSectionInstances) GetFormTemplateId() int32`

GetFormTemplateId returns the FormTemplateId field if non-nil, zero value otherwise.

### GetFormTemplateIdOk

`func (o *OpsAuditSectionInstances) GetFormTemplateIdOk() (*int32, bool)`

GetFormTemplateIdOk returns a tuple with the FormTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormTemplateId

`func (o *OpsAuditSectionInstances) SetFormTemplateId(v int32)`

SetFormTemplateId sets FormTemplateId field to given value.

### HasFormTemplateId

`func (o *OpsAuditSectionInstances) HasFormTemplateId() bool`

HasFormTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


