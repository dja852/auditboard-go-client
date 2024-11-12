# SubprocessesDataPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**File** | Pointer to **string** |  | [optional] 
**SubprocessId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;subprocesses.id&#x60;.&lt;fk table&#x3D;&#39;subprocesses&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ProcessesDatumId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;processes_data.id&#x60;.&lt;fk table&#x3D;&#39;processes_data&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Scopes** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewSubprocessesDataPutPreviousValues

`func NewSubprocessesDataPutPreviousValues() *SubprocessesDataPutPreviousValues`

NewSubprocessesDataPutPreviousValues instantiates a new SubprocessesDataPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubprocessesDataPutPreviousValuesWithDefaults

`func NewSubprocessesDataPutPreviousValuesWithDefaults() *SubprocessesDataPutPreviousValues`

NewSubprocessesDataPutPreviousValuesWithDefaults instantiates a new SubprocessesDataPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubprocessesDataPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubprocessesDataPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubprocessesDataPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubprocessesDataPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubprocessesDataPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubprocessesDataPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubprocessesDataPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubprocessesDataPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SubprocessesDataPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubprocessesDataPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubprocessesDataPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SubprocessesDataPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *SubprocessesDataPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SubprocessesDataPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SubprocessesDataPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *SubprocessesDataPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetFile

`func (o *SubprocessesDataPutPreviousValues) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *SubprocessesDataPutPreviousValues) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *SubprocessesDataPutPreviousValues) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *SubprocessesDataPutPreviousValues) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetSubprocessId

`func (o *SubprocessesDataPutPreviousValues) GetSubprocessId() int32`

GetSubprocessId returns the SubprocessId field if non-nil, zero value otherwise.

### GetSubprocessIdOk

`func (o *SubprocessesDataPutPreviousValues) GetSubprocessIdOk() (*int32, bool)`

GetSubprocessIdOk returns a tuple with the SubprocessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubprocessId

`func (o *SubprocessesDataPutPreviousValues) SetSubprocessId(v int32)`

SetSubprocessId sets SubprocessId field to given value.

### HasSubprocessId

`func (o *SubprocessesDataPutPreviousValues) HasSubprocessId() bool`

HasSubprocessId returns a boolean if a field has been set.

### GetProcessesDatumId

`func (o *SubprocessesDataPutPreviousValues) GetProcessesDatumId() int32`

GetProcessesDatumId returns the ProcessesDatumId field if non-nil, zero value otherwise.

### GetProcessesDatumIdOk

`func (o *SubprocessesDataPutPreviousValues) GetProcessesDatumIdOk() (*int32, bool)`

GetProcessesDatumIdOk returns a tuple with the ProcessesDatumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessesDatumId

`func (o *SubprocessesDataPutPreviousValues) SetProcessesDatumId(v int32)`

SetProcessesDatumId sets ProcessesDatumId field to given value.

### HasProcessesDatumId

`func (o *SubprocessesDataPutPreviousValues) HasProcessesDatumId() bool`

HasProcessesDatumId returns a boolean if a field has been set.

### GetScopes

`func (o *SubprocessesDataPutPreviousValues) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *SubprocessesDataPutPreviousValues) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *SubprocessesDataPutPreviousValues) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *SubprocessesDataPutPreviousValues) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *SubprocessesDataPutPreviousValues) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *SubprocessesDataPutPreviousValues) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


