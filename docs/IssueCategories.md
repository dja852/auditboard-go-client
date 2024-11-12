# IssueCategories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Name** | **string** |  | 
**Key** | Pointer to **string** |  | [optional] 
**SortOrder** | **int32** |  | [default to 0]
**FormTemplateId** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**AllowedIssuesourceableTypes** | **interface{}** |  | 
**IsReportable** | **bool** |  | [default to true]
**PackagingIsStandard** | **bool** |  | [default to false]
**IssueRequiredFields** | Pointer to **interface{}** |  | [optional] 
**ActionPlanRequiredFields** | Pointer to **interface{}** |  | [optional] 
**Scopes** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewIssueCategories

`func NewIssueCategories(name string, sortOrder int32, allowedIssuesourceableTypes interface{}, isReportable bool, packagingIsStandard bool, ) *IssueCategories`

NewIssueCategories instantiates a new IssueCategories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueCategoriesWithDefaults

`func NewIssueCategoriesWithDefaults() *IssueCategories`

NewIssueCategoriesWithDefaults instantiates a new IssueCategories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssueCategories) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueCategories) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueCategories) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IssueCategories) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IssueCategories) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueCategories) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueCategories) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *IssueCategories) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IssueCategories) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IssueCategories) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *IssueCategories) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSortOrder

`func (o *IssueCategories) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *IssueCategories) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *IssueCategories) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.


### GetFormTemplateId

`func (o *IssueCategories) GetFormTemplateId() int32`

GetFormTemplateId returns the FormTemplateId field if non-nil, zero value otherwise.

### GetFormTemplateIdOk

`func (o *IssueCategories) GetFormTemplateIdOk() (*int32, bool)`

GetFormTemplateIdOk returns a tuple with the FormTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormTemplateId

`func (o *IssueCategories) SetFormTemplateId(v int32)`

SetFormTemplateId sets FormTemplateId field to given value.

### HasFormTemplateId

`func (o *IssueCategories) HasFormTemplateId() bool`

HasFormTemplateId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IssueCategories) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssueCategories) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssueCategories) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IssueCategories) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IssueCategories) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IssueCategories) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IssueCategories) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IssueCategories) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *IssueCategories) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *IssueCategories) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *IssueCategories) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *IssueCategories) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetAllowedIssuesourceableTypes

`func (o *IssueCategories) GetAllowedIssuesourceableTypes() interface{}`

GetAllowedIssuesourceableTypes returns the AllowedIssuesourceableTypes field if non-nil, zero value otherwise.

### GetAllowedIssuesourceableTypesOk

`func (o *IssueCategories) GetAllowedIssuesourceableTypesOk() (*interface{}, bool)`

GetAllowedIssuesourceableTypesOk returns a tuple with the AllowedIssuesourceableTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIssuesourceableTypes

`func (o *IssueCategories) SetAllowedIssuesourceableTypes(v interface{})`

SetAllowedIssuesourceableTypes sets AllowedIssuesourceableTypes field to given value.


### SetAllowedIssuesourceableTypesNil

`func (o *IssueCategories) SetAllowedIssuesourceableTypesNil(b bool)`

 SetAllowedIssuesourceableTypesNil sets the value for AllowedIssuesourceableTypes to be an explicit nil

### UnsetAllowedIssuesourceableTypes
`func (o *IssueCategories) UnsetAllowedIssuesourceableTypes()`

UnsetAllowedIssuesourceableTypes ensures that no value is present for AllowedIssuesourceableTypes, not even an explicit nil
### GetIsReportable

`func (o *IssueCategories) GetIsReportable() bool`

GetIsReportable returns the IsReportable field if non-nil, zero value otherwise.

### GetIsReportableOk

`func (o *IssueCategories) GetIsReportableOk() (*bool, bool)`

GetIsReportableOk returns a tuple with the IsReportable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReportable

`func (o *IssueCategories) SetIsReportable(v bool)`

SetIsReportable sets IsReportable field to given value.


### GetPackagingIsStandard

`func (o *IssueCategories) GetPackagingIsStandard() bool`

GetPackagingIsStandard returns the PackagingIsStandard field if non-nil, zero value otherwise.

### GetPackagingIsStandardOk

`func (o *IssueCategories) GetPackagingIsStandardOk() (*bool, bool)`

GetPackagingIsStandardOk returns a tuple with the PackagingIsStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagingIsStandard

`func (o *IssueCategories) SetPackagingIsStandard(v bool)`

SetPackagingIsStandard sets PackagingIsStandard field to given value.


### GetIssueRequiredFields

`func (o *IssueCategories) GetIssueRequiredFields() interface{}`

GetIssueRequiredFields returns the IssueRequiredFields field if non-nil, zero value otherwise.

### GetIssueRequiredFieldsOk

`func (o *IssueCategories) GetIssueRequiredFieldsOk() (*interface{}, bool)`

GetIssueRequiredFieldsOk returns a tuple with the IssueRequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueRequiredFields

`func (o *IssueCategories) SetIssueRequiredFields(v interface{})`

SetIssueRequiredFields sets IssueRequiredFields field to given value.

### HasIssueRequiredFields

`func (o *IssueCategories) HasIssueRequiredFields() bool`

HasIssueRequiredFields returns a boolean if a field has been set.

### SetIssueRequiredFieldsNil

`func (o *IssueCategories) SetIssueRequiredFieldsNil(b bool)`

 SetIssueRequiredFieldsNil sets the value for IssueRequiredFields to be an explicit nil

### UnsetIssueRequiredFields
`func (o *IssueCategories) UnsetIssueRequiredFields()`

UnsetIssueRequiredFields ensures that no value is present for IssueRequiredFields, not even an explicit nil
### GetActionPlanRequiredFields

`func (o *IssueCategories) GetActionPlanRequiredFields() interface{}`

GetActionPlanRequiredFields returns the ActionPlanRequiredFields field if non-nil, zero value otherwise.

### GetActionPlanRequiredFieldsOk

`func (o *IssueCategories) GetActionPlanRequiredFieldsOk() (*interface{}, bool)`

GetActionPlanRequiredFieldsOk returns a tuple with the ActionPlanRequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPlanRequiredFields

`func (o *IssueCategories) SetActionPlanRequiredFields(v interface{})`

SetActionPlanRequiredFields sets ActionPlanRequiredFields field to given value.

### HasActionPlanRequiredFields

`func (o *IssueCategories) HasActionPlanRequiredFields() bool`

HasActionPlanRequiredFields returns a boolean if a field has been set.

### SetActionPlanRequiredFieldsNil

`func (o *IssueCategories) SetActionPlanRequiredFieldsNil(b bool)`

 SetActionPlanRequiredFieldsNil sets the value for ActionPlanRequiredFields to be an explicit nil

### UnsetActionPlanRequiredFields
`func (o *IssueCategories) UnsetActionPlanRequiredFields()`

UnsetActionPlanRequiredFields ensures that no value is present for ActionPlanRequiredFields, not even an explicit nil
### GetScopes

`func (o *IssueCategories) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *IssueCategories) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *IssueCategories) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *IssueCategories) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *IssueCategories) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *IssueCategories) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


