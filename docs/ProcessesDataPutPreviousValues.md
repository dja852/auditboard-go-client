# ProcessesDataPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**File** | Pointer to **string** |  | [optional] 
**ProcessId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;processes.id&#x60;.&lt;fk table&#x3D;&#39;processes&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**EntityId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;entities.id&#x60;.&lt;fk table&#x3D;&#39;entities&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**AuditableEntityId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;auditable_entities.id&#x60;.&lt;fk table&#x3D;&#39;auditable_entities&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Objective** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewProcessesDataPutPreviousValues

`func NewProcessesDataPutPreviousValues() *ProcessesDataPutPreviousValues`

NewProcessesDataPutPreviousValues instantiates a new ProcessesDataPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessesDataPutPreviousValuesWithDefaults

`func NewProcessesDataPutPreviousValuesWithDefaults() *ProcessesDataPutPreviousValues`

NewProcessesDataPutPreviousValuesWithDefaults instantiates a new ProcessesDataPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProcessesDataPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessesDataPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessesDataPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProcessesDataPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProcessesDataPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProcessesDataPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProcessesDataPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProcessesDataPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProcessesDataPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProcessesDataPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProcessesDataPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProcessesDataPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ProcessesDataPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ProcessesDataPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ProcessesDataPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ProcessesDataPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetFile

`func (o *ProcessesDataPutPreviousValues) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ProcessesDataPutPreviousValues) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ProcessesDataPutPreviousValues) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *ProcessesDataPutPreviousValues) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetProcessId

`func (o *ProcessesDataPutPreviousValues) GetProcessId() int32`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *ProcessesDataPutPreviousValues) GetProcessIdOk() (*int32, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *ProcessesDataPutPreviousValues) SetProcessId(v int32)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *ProcessesDataPutPreviousValues) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### GetEntityId

`func (o *ProcessesDataPutPreviousValues) GetEntityId() int32`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ProcessesDataPutPreviousValues) GetEntityIdOk() (*int32, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ProcessesDataPutPreviousValues) SetEntityId(v int32)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *ProcessesDataPutPreviousValues) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetAuditableEntityId

`func (o *ProcessesDataPutPreviousValues) GetAuditableEntityId() int32`

GetAuditableEntityId returns the AuditableEntityId field if non-nil, zero value otherwise.

### GetAuditableEntityIdOk

`func (o *ProcessesDataPutPreviousValues) GetAuditableEntityIdOk() (*int32, bool)`

GetAuditableEntityIdOk returns a tuple with the AuditableEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditableEntityId

`func (o *ProcessesDataPutPreviousValues) SetAuditableEntityId(v int32)`

SetAuditableEntityId sets AuditableEntityId field to given value.

### HasAuditableEntityId

`func (o *ProcessesDataPutPreviousValues) HasAuditableEntityId() bool`

HasAuditableEntityId returns a boolean if a field has been set.

### GetObjective

`func (o *ProcessesDataPutPreviousValues) GetObjective() string`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *ProcessesDataPutPreviousValues) GetObjectiveOk() (*string, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *ProcessesDataPutPreviousValues) SetObjective(v string)`

SetObjective sets Objective field to given value.

### HasObjective

`func (o *ProcessesDataPutPreviousValues) HasObjective() bool`

HasObjective returns a boolean if a field has been set.

### GetScopes

`func (o *ProcessesDataPutPreviousValues) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ProcessesDataPutPreviousValues) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ProcessesDataPutPreviousValues) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ProcessesDataPutPreviousValues) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *ProcessesDataPutPreviousValues) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *ProcessesDataPutPreviousValues) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


