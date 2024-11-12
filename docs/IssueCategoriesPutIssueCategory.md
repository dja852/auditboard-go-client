# IssueCategoriesPutIssueCategory

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

### NewIssueCategoriesPutIssueCategory

`func NewIssueCategoriesPutIssueCategory(name string, sortOrder int32, allowedIssuesourceableTypes interface{}, isReportable bool, packagingIsStandard bool, ) *IssueCategoriesPutIssueCategory`

NewIssueCategoriesPutIssueCategory instantiates a new IssueCategoriesPutIssueCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueCategoriesPutIssueCategoryWithDefaults

`func NewIssueCategoriesPutIssueCategoryWithDefaults() *IssueCategoriesPutIssueCategory`

NewIssueCategoriesPutIssueCategoryWithDefaults instantiates a new IssueCategoriesPutIssueCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssueCategoriesPutIssueCategory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueCategoriesPutIssueCategory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueCategoriesPutIssueCategory) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IssueCategoriesPutIssueCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IssueCategoriesPutIssueCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueCategoriesPutIssueCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueCategoriesPutIssueCategory) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *IssueCategoriesPutIssueCategory) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IssueCategoriesPutIssueCategory) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IssueCategoriesPutIssueCategory) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *IssueCategoriesPutIssueCategory) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSortOrder

`func (o *IssueCategoriesPutIssueCategory) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *IssueCategoriesPutIssueCategory) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *IssueCategoriesPutIssueCategory) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.


### GetFormTemplateId

`func (o *IssueCategoriesPutIssueCategory) GetFormTemplateId() int32`

GetFormTemplateId returns the FormTemplateId field if non-nil, zero value otherwise.

### GetFormTemplateIdOk

`func (o *IssueCategoriesPutIssueCategory) GetFormTemplateIdOk() (*int32, bool)`

GetFormTemplateIdOk returns a tuple with the FormTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormTemplateId

`func (o *IssueCategoriesPutIssueCategory) SetFormTemplateId(v int32)`

SetFormTemplateId sets FormTemplateId field to given value.

### HasFormTemplateId

`func (o *IssueCategoriesPutIssueCategory) HasFormTemplateId() bool`

HasFormTemplateId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IssueCategoriesPutIssueCategory) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssueCategoriesPutIssueCategory) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssueCategoriesPutIssueCategory) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IssueCategoriesPutIssueCategory) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IssueCategoriesPutIssueCategory) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IssueCategoriesPutIssueCategory) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IssueCategoriesPutIssueCategory) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IssueCategoriesPutIssueCategory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *IssueCategoriesPutIssueCategory) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *IssueCategoriesPutIssueCategory) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *IssueCategoriesPutIssueCategory) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *IssueCategoriesPutIssueCategory) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetAllowedIssuesourceableTypes

`func (o *IssueCategoriesPutIssueCategory) GetAllowedIssuesourceableTypes() interface{}`

GetAllowedIssuesourceableTypes returns the AllowedIssuesourceableTypes field if non-nil, zero value otherwise.

### GetAllowedIssuesourceableTypesOk

`func (o *IssueCategoriesPutIssueCategory) GetAllowedIssuesourceableTypesOk() (*interface{}, bool)`

GetAllowedIssuesourceableTypesOk returns a tuple with the AllowedIssuesourceableTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIssuesourceableTypes

`func (o *IssueCategoriesPutIssueCategory) SetAllowedIssuesourceableTypes(v interface{})`

SetAllowedIssuesourceableTypes sets AllowedIssuesourceableTypes field to given value.


### SetAllowedIssuesourceableTypesNil

`func (o *IssueCategoriesPutIssueCategory) SetAllowedIssuesourceableTypesNil(b bool)`

 SetAllowedIssuesourceableTypesNil sets the value for AllowedIssuesourceableTypes to be an explicit nil

### UnsetAllowedIssuesourceableTypes
`func (o *IssueCategoriesPutIssueCategory) UnsetAllowedIssuesourceableTypes()`

UnsetAllowedIssuesourceableTypes ensures that no value is present for AllowedIssuesourceableTypes, not even an explicit nil
### GetIsReportable

`func (o *IssueCategoriesPutIssueCategory) GetIsReportable() bool`

GetIsReportable returns the IsReportable field if non-nil, zero value otherwise.

### GetIsReportableOk

`func (o *IssueCategoriesPutIssueCategory) GetIsReportableOk() (*bool, bool)`

GetIsReportableOk returns a tuple with the IsReportable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReportable

`func (o *IssueCategoriesPutIssueCategory) SetIsReportable(v bool)`

SetIsReportable sets IsReportable field to given value.


### GetPackagingIsStandard

`func (o *IssueCategoriesPutIssueCategory) GetPackagingIsStandard() bool`

GetPackagingIsStandard returns the PackagingIsStandard field if non-nil, zero value otherwise.

### GetPackagingIsStandardOk

`func (o *IssueCategoriesPutIssueCategory) GetPackagingIsStandardOk() (*bool, bool)`

GetPackagingIsStandardOk returns a tuple with the PackagingIsStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagingIsStandard

`func (o *IssueCategoriesPutIssueCategory) SetPackagingIsStandard(v bool)`

SetPackagingIsStandard sets PackagingIsStandard field to given value.


### GetIssueRequiredFields

`func (o *IssueCategoriesPutIssueCategory) GetIssueRequiredFields() interface{}`

GetIssueRequiredFields returns the IssueRequiredFields field if non-nil, zero value otherwise.

### GetIssueRequiredFieldsOk

`func (o *IssueCategoriesPutIssueCategory) GetIssueRequiredFieldsOk() (*interface{}, bool)`

GetIssueRequiredFieldsOk returns a tuple with the IssueRequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueRequiredFields

`func (o *IssueCategoriesPutIssueCategory) SetIssueRequiredFields(v interface{})`

SetIssueRequiredFields sets IssueRequiredFields field to given value.

### HasIssueRequiredFields

`func (o *IssueCategoriesPutIssueCategory) HasIssueRequiredFields() bool`

HasIssueRequiredFields returns a boolean if a field has been set.

### SetIssueRequiredFieldsNil

`func (o *IssueCategoriesPutIssueCategory) SetIssueRequiredFieldsNil(b bool)`

 SetIssueRequiredFieldsNil sets the value for IssueRequiredFields to be an explicit nil

### UnsetIssueRequiredFields
`func (o *IssueCategoriesPutIssueCategory) UnsetIssueRequiredFields()`

UnsetIssueRequiredFields ensures that no value is present for IssueRequiredFields, not even an explicit nil
### GetActionPlanRequiredFields

`func (o *IssueCategoriesPutIssueCategory) GetActionPlanRequiredFields() interface{}`

GetActionPlanRequiredFields returns the ActionPlanRequiredFields field if non-nil, zero value otherwise.

### GetActionPlanRequiredFieldsOk

`func (o *IssueCategoriesPutIssueCategory) GetActionPlanRequiredFieldsOk() (*interface{}, bool)`

GetActionPlanRequiredFieldsOk returns a tuple with the ActionPlanRequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPlanRequiredFields

`func (o *IssueCategoriesPutIssueCategory) SetActionPlanRequiredFields(v interface{})`

SetActionPlanRequiredFields sets ActionPlanRequiredFields field to given value.

### HasActionPlanRequiredFields

`func (o *IssueCategoriesPutIssueCategory) HasActionPlanRequiredFields() bool`

HasActionPlanRequiredFields returns a boolean if a field has been set.

### SetActionPlanRequiredFieldsNil

`func (o *IssueCategoriesPutIssueCategory) SetActionPlanRequiredFieldsNil(b bool)`

 SetActionPlanRequiredFieldsNil sets the value for ActionPlanRequiredFields to be an explicit nil

### UnsetActionPlanRequiredFields
`func (o *IssueCategoriesPutIssueCategory) UnsetActionPlanRequiredFields()`

UnsetActionPlanRequiredFields ensures that no value is present for ActionPlanRequiredFields, not even an explicit nil
### GetScopes

`func (o *IssueCategoriesPutIssueCategory) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *IssueCategoriesPutIssueCategory) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *IssueCategoriesPutIssueCategory) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *IssueCategoriesPutIssueCategory) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *IssueCategoriesPutIssueCategory) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *IssueCategoriesPutIssueCategory) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


