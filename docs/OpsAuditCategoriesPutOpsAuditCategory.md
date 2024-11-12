# OpsAuditCategoriesPutOpsAuditCategory

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

### NewOpsAuditCategoriesPutOpsAuditCategory

`func NewOpsAuditCategoriesPutOpsAuditCategory(name string, sortOrder int32, status string, workspaceType string, reserved bool, ) *OpsAuditCategoriesPutOpsAuditCategory`

NewOpsAuditCategoriesPutOpsAuditCategory instantiates a new OpsAuditCategoriesPutOpsAuditCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsAuditCategoriesPutOpsAuditCategoryWithDefaults

`func NewOpsAuditCategoriesPutOpsAuditCategoryWithDefaults() *OpsAuditCategoriesPutOpsAuditCategory`

NewOpsAuditCategoriesPutOpsAuditCategoryWithDefaults instantiates a new OpsAuditCategoriesPutOpsAuditCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpsAuditCategoriesPutOpsAuditCategory) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OpsAuditCategoriesPutOpsAuditCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OpsAuditCategoriesPutOpsAuditCategory) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OpsAuditCategoriesPutOpsAuditCategory) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OpsAuditCategoriesPutOpsAuditCategory) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OpsAuditCategoriesPutOpsAuditCategory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *OpsAuditCategoriesPutOpsAuditCategory) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *OpsAuditCategoriesPutOpsAuditCategory) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpsAuditCategoriesPutOpsAuditCategory) SetName(v string)`

SetName sets Name field to given value.


### GetSortOrder

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *OpsAuditCategoriesPutOpsAuditCategory) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.


### GetStatus

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpsAuditCategoriesPutOpsAuditCategory) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFormTemplates

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetFormTemplates() interface{}`

GetFormTemplates returns the FormTemplates field if non-nil, zero value otherwise.

### GetFormTemplatesOk

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetFormTemplatesOk() (*interface{}, bool)`

GetFormTemplatesOk returns a tuple with the FormTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormTemplates

`func (o *OpsAuditCategoriesPutOpsAuditCategory) SetFormTemplates(v interface{})`

SetFormTemplates sets FormTemplates field to given value.

### HasFormTemplates

`func (o *OpsAuditCategoriesPutOpsAuditCategory) HasFormTemplates() bool`

HasFormTemplates returns a boolean if a field has been set.

### SetFormTemplatesNil

`func (o *OpsAuditCategoriesPutOpsAuditCategory) SetFormTemplatesNil(b bool)`

 SetFormTemplatesNil sets the value for FormTemplates to be an explicit nil

### UnsetFormTemplates
`func (o *OpsAuditCategoriesPutOpsAuditCategory) UnsetFormTemplates()`

UnsetFormTemplates ensures that no value is present for FormTemplates, not even an explicit nil
### GetOpsAuditRequiredFields

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetOpsAuditRequiredFields() interface{}`

GetOpsAuditRequiredFields returns the OpsAuditRequiredFields field if non-nil, zero value otherwise.

### GetOpsAuditRequiredFieldsOk

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetOpsAuditRequiredFieldsOk() (*interface{}, bool)`

GetOpsAuditRequiredFieldsOk returns a tuple with the OpsAuditRequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditRequiredFields

`func (o *OpsAuditCategoriesPutOpsAuditCategory) SetOpsAuditRequiredFields(v interface{})`

SetOpsAuditRequiredFields sets OpsAuditRequiredFields field to given value.

### HasOpsAuditRequiredFields

`func (o *OpsAuditCategoriesPutOpsAuditCategory) HasOpsAuditRequiredFields() bool`

HasOpsAuditRequiredFields returns a boolean if a field has been set.

### SetOpsAuditRequiredFieldsNil

`func (o *OpsAuditCategoriesPutOpsAuditCategory) SetOpsAuditRequiredFieldsNil(b bool)`

 SetOpsAuditRequiredFieldsNil sets the value for OpsAuditRequiredFields to be an explicit nil

### UnsetOpsAuditRequiredFields
`func (o *OpsAuditCategoriesPutOpsAuditCategory) UnsetOpsAuditRequiredFields()`

UnsetOpsAuditRequiredFields ensures that no value is present for OpsAuditRequiredFields, not even an explicit nil
### GetOpsAuditItemRequiredFields

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetOpsAuditItemRequiredFields() interface{}`

GetOpsAuditItemRequiredFields returns the OpsAuditItemRequiredFields field if non-nil, zero value otherwise.

### GetOpsAuditItemRequiredFieldsOk

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetOpsAuditItemRequiredFieldsOk() (*interface{}, bool)`

GetOpsAuditItemRequiredFieldsOk returns a tuple with the OpsAuditItemRequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsAuditItemRequiredFields

`func (o *OpsAuditCategoriesPutOpsAuditCategory) SetOpsAuditItemRequiredFields(v interface{})`

SetOpsAuditItemRequiredFields sets OpsAuditItemRequiredFields field to given value.

### HasOpsAuditItemRequiredFields

`func (o *OpsAuditCategoriesPutOpsAuditCategory) HasOpsAuditItemRequiredFields() bool`

HasOpsAuditItemRequiredFields returns a boolean if a field has been set.

### SetOpsAuditItemRequiredFieldsNil

`func (o *OpsAuditCategoriesPutOpsAuditCategory) SetOpsAuditItemRequiredFieldsNil(b bool)`

 SetOpsAuditItemRequiredFieldsNil sets the value for OpsAuditItemRequiredFields to be an explicit nil

### UnsetOpsAuditItemRequiredFields
`func (o *OpsAuditCategoriesPutOpsAuditCategory) UnsetOpsAuditItemRequiredFields()`

UnsetOpsAuditItemRequiredFields ensures that no value is present for OpsAuditItemRequiredFields, not even an explicit nil
### GetWorkspaceType

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetWorkspaceType() string`

GetWorkspaceType returns the WorkspaceType field if non-nil, zero value otherwise.

### GetWorkspaceTypeOk

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetWorkspaceTypeOk() (*string, bool)`

GetWorkspaceTypeOk returns a tuple with the WorkspaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceType

`func (o *OpsAuditCategoriesPutOpsAuditCategory) SetWorkspaceType(v string)`

SetWorkspaceType sets WorkspaceType field to given value.


### GetReserved

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *OpsAuditCategoriesPutOpsAuditCategory) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *OpsAuditCategoriesPutOpsAuditCategory) SetReserved(v bool)`

SetReserved sets Reserved field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


