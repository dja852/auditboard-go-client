# OpsAuditCategoriesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**SortOrder** | Pointer to **int32** |  | [optional] [default to 0]
**Status** | Pointer to **string** |  | [optional] [default to ""]
**FormTemplates** | Pointer to **interface{}** |  | [optional] 
**OpsAuditRequiredFields** | Pointer to **interface{}** |  | [optional] 
**OpsAuditItemRequiredFields** | Pointer to **interface{}** |  | [optional] 
**WorkspaceType** | Pointer to **string** |  | [optional] [default to "OpsAudit"]
**Reserved** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewOpsAuditCategoriesPutPreviousValues

`func NewOpsAuditCategoriesPutPreviousValues() *OpsAuditCategoriesPutPreviousValues`

NewOpsAuditCategoriesPutPreviousValues instantiates a new OpsAuditCategoriesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsAuditCategoriesPutPreviousValuesWithDefaults

`func NewOpsAuditCategoriesPutPreviousValuesWithDefaults() *OpsAuditCategoriesPutPreviousValues`

NewOpsAuditCategoriesPutPreviousValuesWithDefaults instantiates a new OpsAuditCategoriesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpsAuditCategoriesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpsAuditCategoriesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpsAuditCategoriesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OpsAuditCategoriesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OpsAuditCategoriesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OpsAuditCategoriesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OpsAuditCategoriesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OpsAuditCategoriesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OpsAuditCategoriesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OpsAuditCategoriesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OpsAuditCategoriesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OpsAuditCategoriesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *OpsAuditCategoriesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *OpsAuditCategoriesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *OpsAuditCategoriesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *OpsAuditCategoriesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *OpsAuditCategoriesPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpsAuditCategoriesPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpsAuditCategoriesPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpsAuditCategoriesPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSortOrder

`func (o *OpsAuditCategoriesPutPreviousValues) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *OpsAuditCategoriesPutPreviousValues) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *OpsAuditCategoriesPutPreviousValues) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *OpsAuditCategoriesPutPreviousValues) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetStatus

`func (o *OpsAuditCategoriesPutPreviousValues) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpsAuditCategoriesPutPreviousValues) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpsAuditCategoriesPutPreviousValues) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OpsAuditCategoriesPutPreviousValues) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFormTemplates

`func (o *OpsAuditCategoriesPutPreviousValues) GetFormTemplates() interface{}`

GetFormTemplates returns the FormTemplates field if non-nil, zero value otherwise.

### GetFormTemplatesOk

`func (o *OpsAuditCategoriesPutPreviousValues) GetFormTemplatesOk() (*interface{}, bool)`

GetFormTemplatesOk returns a tuple with the FormTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormTemplates

`func (o *OpsAuditCategoriesPutPreviousValues) SetFormTemplates(v interface{})`

SetFormTemplates sets FormTemplates field to given value.

### HasFormTemplates

`func (o *OpsAuditCategoriesPutPreviousValues) HasFormTemplates() bool`

HasFormTemplates returns a boolean if a field has been set.

### SetFormTemplatesNil

`func (o *OpsAuditCategoriesPutPreviousValues) SetFormTemplatesNil(b bool)`

 SetFormTemplatesNil sets the value for FormTemplates to be an explicit nil

### UnsetFormTemplates
`func (o *OpsAuditCategoriesPutPreviousValues) UnsetFormTemplates()`

UnsetFormTemplates ensures that no value is present for FormTemplates, not even an explicit nil
### GetOpsAuditRequiredFields

`func (o *OpsAuditCategoriesPutPreviousValues) GetOpsAuditRequiredFields() interface{}`

GetOpsAuditRequiredFields returns the OpsAuditRequiredFields field if non-nil, zero value otherwise.

### GetOpsAuditRequiredFieldsOk

`func (o *OpsAuditCategoriesPutPreviousValues) GetOpsAuditRequiredFieldsOk() (*interface{}, bool)`

GetOpsAuditRequiredFieldsOk returns a tuple with the OpsAuditRequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditRequiredFields

`func (o *OpsAuditCategoriesPutPreviousValues) SetOpsAuditRequiredFields(v interface{})`

SetOpsAuditRequiredFields sets OpsAuditRequiredFields field to given value.

### HasOpsAuditRequiredFields

`func (o *OpsAuditCategoriesPutPreviousValues) HasOpsAuditRequiredFields() bool`

HasOpsAuditRequiredFields returns a boolean if a field has been set.

### SetOpsAuditRequiredFieldsNil

`func (o *OpsAuditCategoriesPutPreviousValues) SetOpsAuditRequiredFieldsNil(b bool)`

 SetOpsAuditRequiredFieldsNil sets the value for OpsAuditRequiredFields to be an explicit nil

### UnsetOpsAuditRequiredFields
`func (o *OpsAuditCategoriesPutPreviousValues) UnsetOpsAuditRequiredFields()`

UnsetOpsAuditRequiredFields ensures that no value is present for OpsAuditRequiredFields, not even an explicit nil
### GetOpsAuditItemRequiredFields

`func (o *OpsAuditCategoriesPutPreviousValues) GetOpsAuditItemRequiredFields() interface{}`

GetOpsAuditItemRequiredFields returns the OpsAuditItemRequiredFields field if non-nil, zero value otherwise.

### GetOpsAuditItemRequiredFieldsOk

`func (o *OpsAuditCategoriesPutPreviousValues) GetOpsAuditItemRequiredFieldsOk() (*interface{}, bool)`

GetOpsAuditItemRequiredFieldsOk returns a tuple with the OpsAuditItemRequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditItemRequiredFields

`func (o *OpsAuditCategoriesPutPreviousValues) SetOpsAuditItemRequiredFields(v interface{})`

SetOpsAuditItemRequiredFields sets OpsAuditItemRequiredFields field to given value.

### HasOpsAuditItemRequiredFields

`func (o *OpsAuditCategoriesPutPreviousValues) HasOpsAuditItemRequiredFields() bool`

HasOpsAuditItemRequiredFields returns a boolean if a field has been set.

### SetOpsAuditItemRequiredFieldsNil

`func (o *OpsAuditCategoriesPutPreviousValues) SetOpsAuditItemRequiredFieldsNil(b bool)`

 SetOpsAuditItemRequiredFieldsNil sets the value for OpsAuditItemRequiredFields to be an explicit nil

### UnsetOpsAuditItemRequiredFields
`func (o *OpsAuditCategoriesPutPreviousValues) UnsetOpsAuditItemRequiredFields()`

UnsetOpsAuditItemRequiredFields ensures that no value is present for OpsAuditItemRequiredFields, not even an explicit nil
### GetWorkspaceType

`func (o *OpsAuditCategoriesPutPreviousValues) GetWorkspaceType() string`

GetWorkspaceType returns the WorkspaceType field if non-nil, zero value otherwise.

### GetWorkspaceTypeOk

`func (o *OpsAuditCategoriesPutPreviousValues) GetWorkspaceTypeOk() (*string, bool)`

GetWorkspaceTypeOk returns a tuple with the WorkspaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceType

`func (o *OpsAuditCategoriesPutPreviousValues) SetWorkspaceType(v string)`

SetWorkspaceType sets WorkspaceType field to given value.

### HasWorkspaceType

`func (o *OpsAuditCategoriesPutPreviousValues) HasWorkspaceType() bool`

HasWorkspaceType returns a boolean if a field has been set.

### GetReserved

`func (o *OpsAuditCategoriesPutPreviousValues) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *OpsAuditCategoriesPutPreviousValues) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *OpsAuditCategoriesPutPreviousValues) SetReserved(v bool)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *OpsAuditCategoriesPutPreviousValues) HasReserved() bool`

HasReserved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


