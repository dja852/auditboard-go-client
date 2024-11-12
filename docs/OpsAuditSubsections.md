# OpsAuditSubsections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**SortOrder** | **int32** |  | [default to 0]
**Name** | **string** |  | 
**OpsAuditId** | **int32** | Note: This is a Foreign Key to &#x60;ops_audits.id&#x60;.&lt;fk table&#x3D;&#39;ops_audits&#39; column&#x3D;&#39;id&#39;/&gt; | 
**OpsAuditSectionInstanceId** | **int32** | Note: This is a Foreign Key to &#x60;ops_audit_section_instances.id&#x60;.&lt;fk table&#x3D;&#39;ops_audit_section_instances&#39; column&#x3D;&#39;id&#39;/&gt; | 

## Methods

### NewOpsAuditSubsections

`func NewOpsAuditSubsections(sortOrder int32, name string, opsAuditId int32, opsAuditSectionInstanceId int32, ) *OpsAuditSubsections`

NewOpsAuditSubsections instantiates a new OpsAuditSubsections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsAuditSubsectionsWithDefaults

`func NewOpsAuditSubsectionsWithDefaults() *OpsAuditSubsections`

NewOpsAuditSubsectionsWithDefaults instantiates a new OpsAuditSubsections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpsAuditSubsections) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpsAuditSubsections) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpsAuditSubsections) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OpsAuditSubsections) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OpsAuditSubsections) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OpsAuditSubsections) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OpsAuditSubsections) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OpsAuditSubsections) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OpsAuditSubsections) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OpsAuditSubsections) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OpsAuditSubsections) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OpsAuditSubsections) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *OpsAuditSubsections) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *OpsAuditSubsections) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *OpsAuditSubsections) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *OpsAuditSubsections) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetSortOrder

`func (o *OpsAuditSubsections) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *OpsAuditSubsections) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *OpsAuditSubsections) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.


### GetName

`func (o *OpsAuditSubsections) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpsAuditSubsections) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpsAuditSubsections) SetName(v string)`

SetName sets Name field to given value.


### GetOpsAuditId

`func (o *OpsAuditSubsections) GetOpsAuditId() int32`

GetOpsAuditId returns the OpsAuditId field if non-nil, zero value otherwise.

### GetOpsAuditIdOk

`func (o *OpsAuditSubsections) GetOpsAuditIdOk() (*int32, bool)`

GetOpsAuditIdOk returns a tuple with the OpsAuditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditId

`func (o *OpsAuditSubsections) SetOpsAuditId(v int32)`

SetOpsAuditId sets OpsAuditId field to given value.


### GetOpsAuditSectionInstanceId

`func (o *OpsAuditSubsections) GetOpsAuditSectionInstanceId() int32`

GetOpsAuditSectionInstanceId returns the OpsAuditSectionInstanceId field if non-nil, zero value otherwise.

### GetOpsAuditSectionInstanceIdOk

`func (o *OpsAuditSubsections) GetOpsAuditSectionInstanceIdOk() (*int32, bool)`

GetOpsAuditSectionInstanceIdOk returns a tuple with the OpsAuditSectionInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditSectionInstanceId

`func (o *OpsAuditSubsections) SetOpsAuditSectionInstanceId(v int32)`

SetOpsAuditSectionInstanceId sets OpsAuditSectionInstanceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


