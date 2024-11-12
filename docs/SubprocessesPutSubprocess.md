# SubprocessesPutSubprocess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | **string** |  | 
**Name** | **string** |  | 
**ProcessId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;processes.id&#x60;.&lt;fk table&#x3D;&#39;processes&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IdCode** | **string** |  | 

## Methods

### NewSubprocessesPutSubprocess

`func NewSubprocessesPutSubprocess(uid string, name string, idCode string, ) *SubprocessesPutSubprocess`

NewSubprocessesPutSubprocess instantiates a new SubprocessesPutSubprocess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubprocessesPutSubprocessWithDefaults

`func NewSubprocessesPutSubprocessWithDefaults() *SubprocessesPutSubprocess`

NewSubprocessesPutSubprocessWithDefaults instantiates a new SubprocessesPutSubprocess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubprocessesPutSubprocess) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubprocessesPutSubprocess) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubprocessesPutSubprocess) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubprocessesPutSubprocess) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubprocessesPutSubprocess) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubprocessesPutSubprocess) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubprocessesPutSubprocess) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubprocessesPutSubprocess) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SubprocessesPutSubprocess) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubprocessesPutSubprocess) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubprocessesPutSubprocess) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SubprocessesPutSubprocess) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *SubprocessesPutSubprocess) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SubprocessesPutSubprocess) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SubprocessesPutSubprocess) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *SubprocessesPutSubprocess) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetUid

`func (o *SubprocessesPutSubprocess) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SubprocessesPutSubprocess) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SubprocessesPutSubprocess) SetUid(v string)`

SetUid sets Uid field to given value.


### GetName

`func (o *SubprocessesPutSubprocess) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubprocessesPutSubprocess) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubprocessesPutSubprocess) SetName(v string)`

SetName sets Name field to given value.


### GetProcessId

`func (o *SubprocessesPutSubprocess) GetProcessId() int32`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *SubprocessesPutSubprocess) GetProcessIdOk() (*int32, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *SubprocessesPutSubprocess) SetProcessId(v int32)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *SubprocessesPutSubprocess) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### GetIdCode

`func (o *SubprocessesPutSubprocess) GetIdCode() string`

GetIdCode returns the IdCode field if non-nil, zero value otherwise.

### GetIdCodeOk

`func (o *SubprocessesPutSubprocess) GetIdCodeOk() (*string, bool)`

GetIdCodeOk returns a tuple with the IdCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdCode

`func (o *SubprocessesPutSubprocess) SetIdCode(v string)`

SetIdCode sets IdCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


