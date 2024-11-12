# SubprocessesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProcessId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;processes.id&#x60;.&lt;fk table&#x3D;&#39;processes&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IdCode** | Pointer to **string** |  | [optional] 

## Methods

### NewSubprocessesPutPreviousValues

`func NewSubprocessesPutPreviousValues() *SubprocessesPutPreviousValues`

NewSubprocessesPutPreviousValues instantiates a new SubprocessesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubprocessesPutPreviousValuesWithDefaults

`func NewSubprocessesPutPreviousValuesWithDefaults() *SubprocessesPutPreviousValues`

NewSubprocessesPutPreviousValuesWithDefaults instantiates a new SubprocessesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubprocessesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubprocessesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubprocessesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubprocessesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubprocessesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubprocessesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubprocessesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubprocessesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SubprocessesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubprocessesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubprocessesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SubprocessesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *SubprocessesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SubprocessesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SubprocessesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *SubprocessesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetUid

`func (o *SubprocessesPutPreviousValues) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SubprocessesPutPreviousValues) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SubprocessesPutPreviousValues) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *SubprocessesPutPreviousValues) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetName

`func (o *SubprocessesPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubprocessesPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubprocessesPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubprocessesPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProcessId

`func (o *SubprocessesPutPreviousValues) GetProcessId() int32`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *SubprocessesPutPreviousValues) GetProcessIdOk() (*int32, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *SubprocessesPutPreviousValues) SetProcessId(v int32)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *SubprocessesPutPreviousValues) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### GetIdCode

`func (o *SubprocessesPutPreviousValues) GetIdCode() string`

GetIdCode returns the IdCode field if non-nil, zero value otherwise.

### GetIdCodeOk

`func (o *SubprocessesPutPreviousValues) GetIdCodeOk() (*string, bool)`

GetIdCodeOk returns a tuple with the IdCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdCode

`func (o *SubprocessesPutPreviousValues) SetIdCode(v string)`

SetIdCode sets IdCode field to given value.

### HasIdCode

`func (o *SubprocessesPutPreviousValues) HasIdCode() bool`

HasIdCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


