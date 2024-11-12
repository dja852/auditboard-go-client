# ProcessesDataPutProcessesDatum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**File** | Pointer to **string** |  | [optional] 
**ProcessId** | **int32** | Note: This is a Foreign Key to &#x60;processes.id&#x60;.&lt;fk table&#x3D;&#39;processes&#39; column&#x3D;&#39;id&#39;/&gt; | 
**EntityId** | **int32** | Note: This is a Foreign Key to &#x60;entities.id&#x60;.&lt;fk table&#x3D;&#39;entities&#39; column&#x3D;&#39;id&#39;/&gt; | 
**AuditableEntityId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;auditable_entities.id&#x60;.&lt;fk table&#x3D;&#39;auditable_entities&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Objective** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewProcessesDataPutProcessesDatum

`func NewProcessesDataPutProcessesDatum(processId int32, entityId int32, ) *ProcessesDataPutProcessesDatum`

NewProcessesDataPutProcessesDatum instantiates a new ProcessesDataPutProcessesDatum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessesDataPutProcessesDatumWithDefaults

`func NewProcessesDataPutProcessesDatumWithDefaults() *ProcessesDataPutProcessesDatum`

NewProcessesDataPutProcessesDatumWithDefaults instantiates a new ProcessesDataPutProcessesDatum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProcessesDataPutProcessesDatum) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessesDataPutProcessesDatum) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessesDataPutProcessesDatum) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProcessesDataPutProcessesDatum) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProcessesDataPutProcessesDatum) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProcessesDataPutProcessesDatum) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProcessesDataPutProcessesDatum) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProcessesDataPutProcessesDatum) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProcessesDataPutProcessesDatum) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProcessesDataPutProcessesDatum) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProcessesDataPutProcessesDatum) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProcessesDataPutProcessesDatum) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ProcessesDataPutProcessesDatum) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ProcessesDataPutProcessesDatum) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ProcessesDataPutProcessesDatum) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ProcessesDataPutProcessesDatum) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetFile

`func (o *ProcessesDataPutProcessesDatum) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ProcessesDataPutProcessesDatum) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ProcessesDataPutProcessesDatum) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *ProcessesDataPutProcessesDatum) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetProcessId

`func (o *ProcessesDataPutProcessesDatum) GetProcessId() int32`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *ProcessesDataPutProcessesDatum) GetProcessIdOk() (*int32, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *ProcessesDataPutProcessesDatum) SetProcessId(v int32)`

SetProcessId sets ProcessId field to given value.


### GetEntityId

`func (o *ProcessesDataPutProcessesDatum) GetEntityId() int32`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ProcessesDataPutProcessesDatum) GetEntityIdOk() (*int32, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ProcessesDataPutProcessesDatum) SetEntityId(v int32)`

SetEntityId sets EntityId field to given value.


### GetAuditableEntityId

`func (o *ProcessesDataPutProcessesDatum) GetAuditableEntityId() int32`

GetAuditableEntityId returns the AuditableEntityId field if non-nil, zero value otherwise.

### GetAuditableEntityIdOk

`func (o *ProcessesDataPutProcessesDatum) GetAuditableEntityIdOk() (*int32, bool)`

GetAuditableEntityIdOk returns a tuple with the AuditableEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditableEntityId

`func (o *ProcessesDataPutProcessesDatum) SetAuditableEntityId(v int32)`

SetAuditableEntityId sets AuditableEntityId field to given value.

### HasAuditableEntityId

`func (o *ProcessesDataPutProcessesDatum) HasAuditableEntityId() bool`

HasAuditableEntityId returns a boolean if a field has been set.

### GetObjective

`func (o *ProcessesDataPutProcessesDatum) GetObjective() string`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *ProcessesDataPutProcessesDatum) GetObjectiveOk() (*string, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *ProcessesDataPutProcessesDatum) SetObjective(v string)`

SetObjective sets Objective field to given value.

### HasObjective

`func (o *ProcessesDataPutProcessesDatum) HasObjective() bool`

HasObjective returns a boolean if a field has been set.

### GetScopes

`func (o *ProcessesDataPutProcessesDatum) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ProcessesDataPutProcessesDatum) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ProcessesDataPutProcessesDatum) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ProcessesDataPutProcessesDatum) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *ProcessesDataPutProcessesDatum) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *ProcessesDataPutProcessesDatum) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


