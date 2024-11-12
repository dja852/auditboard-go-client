# SubprocessesDataPutSubprocessesDatum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**File** | Pointer to **string** |  | [optional] 
**SubprocessId** | **int32** | Note: This is a Foreign Key to &#x60;subprocesses.id&#x60;.&lt;fk table&#x3D;&#39;subprocesses&#39; column&#x3D;&#39;id&#39;/&gt; | 
**ProcessesDatumId** | **int32** | Note: This is a Foreign Key to &#x60;processes_data.id&#x60;.&lt;fk table&#x3D;&#39;processes_data&#39; column&#x3D;&#39;id&#39;/&gt; | 
**Scopes** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewSubprocessesDataPutSubprocessesDatum

`func NewSubprocessesDataPutSubprocessesDatum(subprocessId int32, processesDatumId int32, ) *SubprocessesDataPutSubprocessesDatum`

NewSubprocessesDataPutSubprocessesDatum instantiates a new SubprocessesDataPutSubprocessesDatum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubprocessesDataPutSubprocessesDatumWithDefaults

`func NewSubprocessesDataPutSubprocessesDatumWithDefaults() *SubprocessesDataPutSubprocessesDatum`

NewSubprocessesDataPutSubprocessesDatumWithDefaults instantiates a new SubprocessesDataPutSubprocessesDatum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubprocessesDataPutSubprocessesDatum) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubprocessesDataPutSubprocessesDatum) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubprocessesDataPutSubprocessesDatum) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubprocessesDataPutSubprocessesDatum) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubprocessesDataPutSubprocessesDatum) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubprocessesDataPutSubprocessesDatum) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubprocessesDataPutSubprocessesDatum) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubprocessesDataPutSubprocessesDatum) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SubprocessesDataPutSubprocessesDatum) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubprocessesDataPutSubprocessesDatum) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubprocessesDataPutSubprocessesDatum) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SubprocessesDataPutSubprocessesDatum) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *SubprocessesDataPutSubprocessesDatum) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SubprocessesDataPutSubprocessesDatum) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SubprocessesDataPutSubprocessesDatum) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *SubprocessesDataPutSubprocessesDatum) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetFile

`func (o *SubprocessesDataPutSubprocessesDatum) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *SubprocessesDataPutSubprocessesDatum) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *SubprocessesDataPutSubprocessesDatum) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *SubprocessesDataPutSubprocessesDatum) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetSubprocessId

`func (o *SubprocessesDataPutSubprocessesDatum) GetSubprocessId() int32`

GetSubprocessId returns the SubprocessId field if non-nil, zero value otherwise.

### GetSubprocessIdOk

`func (o *SubprocessesDataPutSubprocessesDatum) GetSubprocessIdOk() (*int32, bool)`

GetSubprocessIdOk returns a tuple with the SubprocessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubprocessId

`func (o *SubprocessesDataPutSubprocessesDatum) SetSubprocessId(v int32)`

SetSubprocessId sets SubprocessId field to given value.


### GetProcessesDatumId

`func (o *SubprocessesDataPutSubprocessesDatum) GetProcessesDatumId() int32`

GetProcessesDatumId returns the ProcessesDatumId field if non-nil, zero value otherwise.

### GetProcessesDatumIdOk

`func (o *SubprocessesDataPutSubprocessesDatum) GetProcessesDatumIdOk() (*int32, bool)`

GetProcessesDatumIdOk returns a tuple with the ProcessesDatumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessesDatumId

`func (o *SubprocessesDataPutSubprocessesDatum) SetProcessesDatumId(v int32)`

SetProcessesDatumId sets ProcessesDatumId field to given value.


### GetScopes

`func (o *SubprocessesDataPutSubprocessesDatum) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *SubprocessesDataPutSubprocessesDatum) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *SubprocessesDataPutSubprocessesDatum) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *SubprocessesDataPutSubprocessesDatum) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *SubprocessesDataPutSubprocessesDatum) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *SubprocessesDataPutSubprocessesDatum) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


