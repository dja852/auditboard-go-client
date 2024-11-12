# ControlsDataArchives

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**ControlsDatumId** | **int32** | Note: This is a Foreign Key to &#x60;controls_data.id&#x60;.&lt;fk table&#x3D;&#39;controls_data&#39; column&#x3D;&#39;id&#39;/&gt; | 
**ArchiveId** | **int32** | Note: This is a Foreign Key to &#x60;archives.id&#x60;.&lt;fk table&#x3D;&#39;archives&#39; column&#x3D;&#39;id&#39;/&gt; | 
**ArchiveData** | **string** |  | 
**ControlsDatumDescription** | Pointer to **string** |  | [optional] 
**EntityName** | Pointer to **string** |  | [optional] 
**ControlName** | Pointer to **string** |  | [optional] 
**ControlUid** | Pointer to **string** |  | [optional] 
**EntityCode** | Pointer to **string** |  | [optional] 
**RegionCode** | Pointer to **string** |  | [optional] 
**RegionName** | Pointer to **string** |  | [optional] 
**ProcessUid** | Pointer to **string** |  | [optional] 
**ProcessName** | Pointer to **string** |  | [optional] 
**SubprocessUid** | Pointer to **string** |  | [optional] 
**SubprocessName** | Pointer to **string** |  | [optional] 
**RegionId** | Pointer to **int32** |  | [optional] 
**EntityId** | Pointer to **int32** |  | [optional] 
**SubprocessId** | Pointer to **int32** |  | [optional] 
**ProcessId** | Pointer to **int32** |  | [optional] 
**ControlId** | Pointer to **int32** |  | [optional] 
**FieldData** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewControlsDataArchives

`func NewControlsDataArchives(controlsDatumId int32, archiveId int32, archiveData string, ) *ControlsDataArchives`

NewControlsDataArchives instantiates a new ControlsDataArchives object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsDataArchivesWithDefaults

`func NewControlsDataArchivesWithDefaults() *ControlsDataArchives`

NewControlsDataArchivesWithDefaults instantiates a new ControlsDataArchives object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ControlsDataArchives) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControlsDataArchives) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControlsDataArchives) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ControlsDataArchives) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ControlsDataArchives) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ControlsDataArchives) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ControlsDataArchives) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ControlsDataArchives) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ControlsDataArchives) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ControlsDataArchives) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ControlsDataArchives) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ControlsDataArchives) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ControlsDataArchives) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ControlsDataArchives) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ControlsDataArchives) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ControlsDataArchives) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetControlsDatumId

`func (o *ControlsDataArchives) GetControlsDatumId() int32`

GetControlsDatumId returns the ControlsDatumId field if non-nil, zero value otherwise.

### GetControlsDatumIdOk

`func (o *ControlsDataArchives) GetControlsDatumIdOk() (*int32, bool)`

GetControlsDatumIdOk returns a tuple with the ControlsDatumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlsDatumId

`func (o *ControlsDataArchives) SetControlsDatumId(v int32)`

SetControlsDatumId sets ControlsDatumId field to given value.


### GetArchiveId

`func (o *ControlsDataArchives) GetArchiveId() int32`

GetArchiveId returns the ArchiveId field if non-nil, zero value otherwise.

### GetArchiveIdOk

`func (o *ControlsDataArchives) GetArchiveIdOk() (*int32, bool)`

GetArchiveIdOk returns a tuple with the ArchiveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveId

`func (o *ControlsDataArchives) SetArchiveId(v int32)`

SetArchiveId sets ArchiveId field to given value.


### GetArchiveData

`func (o *ControlsDataArchives) GetArchiveData() string`

GetArchiveData returns the ArchiveData field if non-nil, zero value otherwise.

### GetArchiveDataOk

`func (o *ControlsDataArchives) GetArchiveDataOk() (*string, bool)`

GetArchiveDataOk returns a tuple with the ArchiveData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveData

`func (o *ControlsDataArchives) SetArchiveData(v string)`

SetArchiveData sets ArchiveData field to given value.


### GetControlsDatumDescription

`func (o *ControlsDataArchives) GetControlsDatumDescription() string`

GetControlsDatumDescription returns the ControlsDatumDescription field if non-nil, zero value otherwise.

### GetControlsDatumDescriptionOk

`func (o *ControlsDataArchives) GetControlsDatumDescriptionOk() (*string, bool)`

GetControlsDatumDescriptionOk returns a tuple with the ControlsDatumDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlsDatumDescription

`func (o *ControlsDataArchives) SetControlsDatumDescription(v string)`

SetControlsDatumDescription sets ControlsDatumDescription field to given value.

### HasControlsDatumDescription

`func (o *ControlsDataArchives) HasControlsDatumDescription() bool`

HasControlsDatumDescription returns a boolean if a field has been set.

### GetEntityName

`func (o *ControlsDataArchives) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *ControlsDataArchives) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *ControlsDataArchives) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *ControlsDataArchives) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetControlName

`func (o *ControlsDataArchives) GetControlName() string`

GetControlName returns the ControlName field if non-nil, zero value otherwise.

### GetControlNameOk

`func (o *ControlsDataArchives) GetControlNameOk() (*string, bool)`

GetControlNameOk returns a tuple with the ControlName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlName

`func (o *ControlsDataArchives) SetControlName(v string)`

SetControlName sets ControlName field to given value.

### HasControlName

`func (o *ControlsDataArchives) HasControlName() bool`

HasControlName returns a boolean if a field has been set.

### GetControlUid

`func (o *ControlsDataArchives) GetControlUid() string`

GetControlUid returns the ControlUid field if non-nil, zero value otherwise.

### GetControlUidOk

`func (o *ControlsDataArchives) GetControlUidOk() (*string, bool)`

GetControlUidOk returns a tuple with the ControlUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlUid

`func (o *ControlsDataArchives) SetControlUid(v string)`

SetControlUid sets ControlUid field to given value.

### HasControlUid

`func (o *ControlsDataArchives) HasControlUid() bool`

HasControlUid returns a boolean if a field has been set.

### GetEntityCode

`func (o *ControlsDataArchives) GetEntityCode() string`

GetEntityCode returns the EntityCode field if non-nil, zero value otherwise.

### GetEntityCodeOk

`func (o *ControlsDataArchives) GetEntityCodeOk() (*string, bool)`

GetEntityCodeOk returns a tuple with the EntityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityCode

`func (o *ControlsDataArchives) SetEntityCode(v string)`

SetEntityCode sets EntityCode field to given value.

### HasEntityCode

`func (o *ControlsDataArchives) HasEntityCode() bool`

HasEntityCode returns a boolean if a field has been set.

### GetRegionCode

`func (o *ControlsDataArchives) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *ControlsDataArchives) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *ControlsDataArchives) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *ControlsDataArchives) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetRegionName

`func (o *ControlsDataArchives) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ControlsDataArchives) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ControlsDataArchives) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ControlsDataArchives) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetProcessUid

`func (o *ControlsDataArchives) GetProcessUid() string`

GetProcessUid returns the ProcessUid field if non-nil, zero value otherwise.

### GetProcessUidOk

`func (o *ControlsDataArchives) GetProcessUidOk() (*string, bool)`

GetProcessUidOk returns a tuple with the ProcessUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessUid

`func (o *ControlsDataArchives) SetProcessUid(v string)`

SetProcessUid sets ProcessUid field to given value.

### HasProcessUid

`func (o *ControlsDataArchives) HasProcessUid() bool`

HasProcessUid returns a boolean if a field has been set.

### GetProcessName

`func (o *ControlsDataArchives) GetProcessName() string`

GetProcessName returns the ProcessName field if non-nil, zero value otherwise.

### GetProcessNameOk

`func (o *ControlsDataArchives) GetProcessNameOk() (*string, bool)`

GetProcessNameOk returns a tuple with the ProcessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessName

`func (o *ControlsDataArchives) SetProcessName(v string)`

SetProcessName sets ProcessName field to given value.

### HasProcessName

`func (o *ControlsDataArchives) HasProcessName() bool`

HasProcessName returns a boolean if a field has been set.

### GetSubprocessUid

`func (o *ControlsDataArchives) GetSubprocessUid() string`

GetSubprocessUid returns the SubprocessUid field if non-nil, zero value otherwise.

### GetSubprocessUidOk

`func (o *ControlsDataArchives) GetSubprocessUidOk() (*string, bool)`

GetSubprocessUidOk returns a tuple with the SubprocessUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubprocessUid

`func (o *ControlsDataArchives) SetSubprocessUid(v string)`

SetSubprocessUid sets SubprocessUid field to given value.

### HasSubprocessUid

`func (o *ControlsDataArchives) HasSubprocessUid() bool`

HasSubprocessUid returns a boolean if a field has been set.

### GetSubprocessName

`func (o *ControlsDataArchives) GetSubprocessName() string`

GetSubprocessName returns the SubprocessName field if non-nil, zero value otherwise.

### GetSubprocessNameOk

`func (o *ControlsDataArchives) GetSubprocessNameOk() (*string, bool)`

GetSubprocessNameOk returns a tuple with the SubprocessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubprocessName

`func (o *ControlsDataArchives) SetSubprocessName(v string)`

SetSubprocessName sets SubprocessName field to given value.

### HasSubprocessName

`func (o *ControlsDataArchives) HasSubprocessName() bool`

HasSubprocessName returns a boolean if a field has been set.

### GetRegionId

`func (o *ControlsDataArchives) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *ControlsDataArchives) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *ControlsDataArchives) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *ControlsDataArchives) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetEntityId

`func (o *ControlsDataArchives) GetEntityId() int32`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ControlsDataArchives) GetEntityIdOk() (*int32, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ControlsDataArchives) SetEntityId(v int32)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *ControlsDataArchives) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetSubprocessId

`func (o *ControlsDataArchives) GetSubprocessId() int32`

GetSubprocessId returns the SubprocessId field if non-nil, zero value otherwise.

### GetSubprocessIdOk

`func (o *ControlsDataArchives) GetSubprocessIdOk() (*int32, bool)`

GetSubprocessIdOk returns a tuple with the SubprocessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubprocessId

`func (o *ControlsDataArchives) SetSubprocessId(v int32)`

SetSubprocessId sets SubprocessId field to given value.

### HasSubprocessId

`func (o *ControlsDataArchives) HasSubprocessId() bool`

HasSubprocessId returns a boolean if a field has been set.

### GetProcessId

`func (o *ControlsDataArchives) GetProcessId() int32`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *ControlsDataArchives) GetProcessIdOk() (*int32, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *ControlsDataArchives) SetProcessId(v int32)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *ControlsDataArchives) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### GetControlId

`func (o *ControlsDataArchives) GetControlId() int32`

GetControlId returns the ControlId field if non-nil, zero value otherwise.

### GetControlIdOk

`func (o *ControlsDataArchives) GetControlIdOk() (*int32, bool)`

GetControlIdOk returns a tuple with the ControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlId

`func (o *ControlsDataArchives) SetControlId(v int32)`

SetControlId sets ControlId field to given value.

### HasControlId

`func (o *ControlsDataArchives) HasControlId() bool`

HasControlId returns a boolean if a field has been set.

### GetFieldData

`func (o *ControlsDataArchives) GetFieldData() interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *ControlsDataArchives) GetFieldDataOk() (*interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *ControlsDataArchives) SetFieldData(v interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *ControlsDataArchives) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### SetFieldDataNil

`func (o *ControlsDataArchives) SetFieldDataNil(b bool)`

 SetFieldDataNil sets the value for FieldData to be an explicit nil

### UnsetFieldData
`func (o *ControlsDataArchives) UnsetFieldData()`

UnsetFieldData ensures that no value is present for FieldData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


