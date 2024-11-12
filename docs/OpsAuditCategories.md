# OpsAuditCategories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**SortOrder** | **int32** |  | [default to 0]
**Status** | **string** |  | [default to ""]
**FormTemplates** | Pointer to **interface{}** |  | [optional] 
**OpsAuditRequiredFields** | Pointer to **interface{}** |  | [optional] 
**OpsAuditItemRequiredFields** | Pointer to **interface{}** |  | [optional] 
**WorkspaceType** | **string** |  | [default to "OpsAudit"]
**Reserved** | **bool** |  | [default to false]

## Methods

### NewOpsAuditCategories

`func NewOpsAuditCategories(name string, sortOrder int32, status string, workspaceType string, reserved bool, ) *OpsAuditCategories`

NewOpsAuditCategories instantiates a new OpsAuditCategories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsAuditCategoriesWithDefaults

`func NewOpsAuditCategoriesWithDefaults() *OpsAuditCategories`

NewOpsAuditCategoriesWithDefaults instantiates a new OpsAuditCategories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpsAuditCategories) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpsAuditCategories) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpsAuditCategories) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OpsAuditCategories) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OpsAuditCategories) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OpsAuditCategories) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OpsAuditCategories) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OpsAuditCategories) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OpsAuditCategories) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OpsAuditCategories) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OpsAuditCategories) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OpsAuditCategories) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *OpsAuditCategories) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *OpsAuditCategories) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *OpsAuditCategories) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *OpsAuditCategories) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *OpsAuditCategories) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpsAuditCategories) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpsAuditCategories) SetName(v string)`

SetName sets Name field to given value.


### GetSortOrder

`func (o *OpsAuditCategories) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *OpsAuditCategories) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *OpsAuditCategories) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.


### GetStatus

`func (o *OpsAuditCategories) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpsAuditCategories) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpsAuditCategories) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFormTemplates

`func (o *OpsAuditCategories) GetFormTemplates() interface{}`

GetFormTemplates returns the FormTemplates field if non-nil, zero value otherwise.

### GetFormTemplatesOk

`func (o *OpsAuditCategories) GetFormTemplatesOk() (*interface{}, bool)`

GetFormTemplatesOk returns a tuple with the FormTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormTemplates

`func (o *OpsAuditCategories) SetFormTemplates(v interface{})`

SetFormTemplates sets FormTemplates field to given value.

### HasFormTemplates

`func (o *OpsAuditCategories) HasFormTemplates() bool`

HasFormTemplates returns a boolean if a field has been set.

### SetFormTemplatesNil

`func (o *OpsAuditCategories) SetFormTemplatesNil(b bool)`

 SetFormTemplatesNil sets the value for FormTemplates to be an explicit nil

### UnsetFormTemplates
`func (o *OpsAuditCategories) UnsetFormTemplates()`

UnsetFormTemplates ensures that no value is present for FormTemplates, not even an explicit nil
### GetOpsAuditRequiredFields

`func (o *OpsAuditCategories) GetOpsAuditRequiredFields() interface{}`

GetOpsAuditRequiredFields returns the OpsAuditRequiredFields field if non-nil, zero value otherwise.

### GetOpsAuditRequiredFieldsOk

`func (o *OpsAuditCategories) GetOpsAuditRequiredFieldsOk() (*interface{}, bool)`

GetOpsAuditRequiredFieldsOk returns a tuple with the OpsAuditRequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditRequiredFields

`func (o *OpsAuditCategories) SetOpsAuditRequiredFields(v interface{})`

SetOpsAuditRequiredFields sets OpsAuditRequiredFields field to given value.

### HasOpsAuditRequiredFields

`func (o *OpsAuditCategories) HasOpsAuditRequiredFields() bool`

HasOpsAuditRequiredFields returns a boolean if a field has been set.

### SetOpsAuditRequiredFieldsNil

`func (o *OpsAuditCategories) SetOpsAuditRequiredFieldsNil(b bool)`

 SetOpsAuditRequiredFieldsNil sets the value for OpsAuditRequiredFields to be an explicit nil

### UnsetOpsAuditRequiredFields
`func (o *OpsAuditCategories) UnsetOpsAuditRequiredFields()`

UnsetOpsAuditRequiredFields ensures that no value is present for OpsAuditRequiredFields, not even an explicit nil
### GetOpsAuditItemRequiredFields

`func (o *OpsAuditCategories) GetOpsAuditItemRequiredFields() interface{}`

GetOpsAuditItemRequiredFields returns the OpsAuditItemRequiredFields field if non-nil, zero value otherwise.

### GetOpsAuditItemRequiredFieldsOk

`func (o *OpsAuditCategories) GetOpsAuditItemRequiredFieldsOk() (*interface{}, bool)`

GetOpsAuditItemRequiredFieldsOk returns a tuple with the OpsAuditItemRequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditItemRequiredFields

`func (o *OpsAuditCategories) SetOpsAuditItemRequiredFields(v interface{})`

SetOpsAuditItemRequiredFields sets OpsAuditItemRequiredFields field to given value.

### HasOpsAuditItemRequiredFields

`func (o *OpsAuditCategories) HasOpsAuditItemRequiredFields() bool`

HasOpsAuditItemRequiredFields returns a boolean if a field has been set.

### SetOpsAuditItemRequiredFieldsNil

`func (o *OpsAuditCategories) SetOpsAuditItemRequiredFieldsNil(b bool)`

 SetOpsAuditItemRequiredFieldsNil sets the value for OpsAuditItemRequiredFields to be an explicit nil

### UnsetOpsAuditItemRequiredFields
`func (o *OpsAuditCategories) UnsetOpsAuditItemRequiredFields()`

UnsetOpsAuditItemRequiredFields ensures that no value is present for OpsAuditItemRequiredFields, not even an explicit nil
### GetWorkspaceType

`func (o *OpsAuditCategories) GetWorkspaceType() string`

GetWorkspaceType returns the WorkspaceType field if non-nil, zero value otherwise.

### GetWorkspaceTypeOk

`func (o *OpsAuditCategories) GetWorkspaceTypeOk() (*string, bool)`

GetWorkspaceTypeOk returns a tuple with the WorkspaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceType

`func (o *OpsAuditCategories) SetWorkspaceType(v string)`

SetWorkspaceType sets WorkspaceType field to given value.


### GetReserved

`func (o *OpsAuditCategories) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *OpsAuditCategories) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *OpsAuditCategories) SetReserved(v bool)`

SetReserved sets Reserved field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


