# SubprocessesData

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

### NewSubprocessesData

`func NewSubprocessesData(subprocessId int32, processesDatumId int32, ) *SubprocessesData`

NewSubprocessesData instantiates a new SubprocessesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubprocessesDataWithDefaults

`func NewSubprocessesDataWithDefaults() *SubprocessesData`

NewSubprocessesDataWithDefaults instantiates a new SubprocessesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubprocessesData) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubprocessesData) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubprocessesData) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubprocessesData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubprocessesData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubprocessesData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubprocessesData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubprocessesData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SubprocessesData) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubprocessesData) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubprocessesData) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SubprocessesData) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *SubprocessesData) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SubprocessesData) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SubprocessesData) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *SubprocessesData) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetFile

`func (o *SubprocessesData) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *SubprocessesData) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *SubprocessesData) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *SubprocessesData) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetSubprocessId

`func (o *SubprocessesData) GetSubprocessId() int32`

GetSubprocessId returns the SubprocessId field if non-nil, zero value otherwise.

### GetSubprocessIdOk

`func (o *SubprocessesData) GetSubprocessIdOk() (*int32, bool)`

GetSubprocessIdOk returns a tuple with the SubprocessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubprocessId

`func (o *SubprocessesData) SetSubprocessId(v int32)`

SetSubprocessId sets SubprocessId field to given value.


### GetProcessesDatumId

`func (o *SubprocessesData) GetProcessesDatumId() int32`

GetProcessesDatumId returns the ProcessesDatumId field if non-nil, zero value otherwise.

### GetProcessesDatumIdOk

`func (o *SubprocessesData) GetProcessesDatumIdOk() (*int32, bool)`

GetProcessesDatumIdOk returns a tuple with the ProcessesDatumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessesDatumId

`func (o *SubprocessesData) SetProcessesDatumId(v int32)`

SetProcessesDatumId sets ProcessesDatumId field to given value.


### GetScopes

`func (o *SubprocessesData) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *SubprocessesData) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *SubprocessesData) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *SubprocessesData) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *SubprocessesData) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *SubprocessesData) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


