# IssueCategoriesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**SortOrder** | Pointer to **int32** |  | [optional] [default to 0]
**FormTemplateId** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**AllowedIssuesourceableTypes** | Pointer to **interface{}** |  | [optional] 
**IsReportable** | Pointer to **bool** |  | [optional] [default to true]
**PackagingIsStandard** | Pointer to **bool** |  | [optional] [default to false]
**IssueRequiredFields** | Pointer to **interface{}** |  | [optional] 
**ActionPlanRequiredFields** | Pointer to **interface{}** |  | [optional] 
**Scopes** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewIssueCategoriesPutPreviousValues

`func NewIssueCategoriesPutPreviousValues() *IssueCategoriesPutPreviousValues`

NewIssueCategoriesPutPreviousValues instantiates a new IssueCategoriesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueCategoriesPutPreviousValuesWithDefaults

`func NewIssueCategoriesPutPreviousValuesWithDefaults() *IssueCategoriesPutPreviousValues`

NewIssueCategoriesPutPreviousValuesWithDefaults instantiates a new IssueCategoriesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssueCategoriesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueCategoriesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueCategoriesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IssueCategoriesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IssueCategoriesPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueCategoriesPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueCategoriesPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueCategoriesPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *IssueCategoriesPutPreviousValues) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IssueCategoriesPutPreviousValues) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IssueCategoriesPutPreviousValues) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *IssueCategoriesPutPreviousValues) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSortOrder

`func (o *IssueCategoriesPutPreviousValues) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *IssueCategoriesPutPreviousValues) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *IssueCategoriesPutPreviousValues) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *IssueCategoriesPutPreviousValues) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetFormTemplateId

`func (o *IssueCategoriesPutPreviousValues) GetFormTemplateId() int32`

GetFormTemplateId returns the FormTemplateId field if non-nil, zero value otherwise.

### GetFormTemplateIdOk

`func (o *IssueCategoriesPutPreviousValues) GetFormTemplateIdOk() (*int32, bool)`

GetFormTemplateIdOk returns a tuple with the FormTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormTemplateId

`func (o *IssueCategoriesPutPreviousValues) SetFormTemplateId(v int32)`

SetFormTemplateId sets FormTemplateId field to given value.

### HasFormTemplateId

`func (o *IssueCategoriesPutPreviousValues) HasFormTemplateId() bool`

HasFormTemplateId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IssueCategoriesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssueCategoriesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssueCategoriesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IssueCategoriesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IssueCategoriesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IssueCategoriesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IssueCategoriesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IssueCategoriesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *IssueCategoriesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *IssueCategoriesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *IssueCategoriesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *IssueCategoriesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetAllowedIssuesourceableTypes

`func (o *IssueCategoriesPutPreviousValues) GetAllowedIssuesourceableTypes() interface{}`

GetAllowedIssuesourceableTypes returns the AllowedIssuesourceableTypes field if non-nil, zero value otherwise.

### GetAllowedIssuesourceableTypesOk

`func (o *IssueCategoriesPutPreviousValues) GetAllowedIssuesourceableTypesOk() (*interface{}, bool)`

GetAllowedIssuesourceableTypesOk returns a tuple with the AllowedIssuesourceableTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIssuesourceableTypes

`func (o *IssueCategoriesPutPreviousValues) SetAllowedIssuesourceableTypes(v interface{})`

SetAllowedIssuesourceableTypes sets AllowedIssuesourceableTypes field to given value.

### HasAllowedIssuesourceableTypes

`func (o *IssueCategoriesPutPreviousValues) HasAllowedIssuesourceableTypes() bool`

HasAllowedIssuesourceableTypes returns a boolean if a field has been set.

### SetAllowedIssuesourceableTypesNil

`func (o *IssueCategoriesPutPreviousValues) SetAllowedIssuesourceableTypesNil(b bool)`

 SetAllowedIssuesourceableTypesNil sets the value for AllowedIssuesourceableTypes to be an explicit nil

### UnsetAllowedIssuesourceableTypes
`func (o *IssueCategoriesPutPreviousValues) UnsetAllowedIssuesourceableTypes()`

UnsetAllowedIssuesourceableTypes ensures that no value is present for AllowedIssuesourceableTypes, not even an explicit nil
### GetIsReportable

`func (o *IssueCategoriesPutPreviousValues) GetIsReportable() bool`

GetIsReportable returns the IsReportable field if non-nil, zero value otherwise.

### GetIsReportableOk

`func (o *IssueCategoriesPutPreviousValues) GetIsReportableOk() (*bool, bool)`

GetIsReportableOk returns a tuple with the IsReportable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReportable

`func (o *IssueCategoriesPutPreviousValues) SetIsReportable(v bool)`

SetIsReportable sets IsReportable field to given value.

### HasIsReportable

`func (o *IssueCategoriesPutPreviousValues) HasIsReportable() bool`

HasIsReportable returns a boolean if a field has been set.

### GetPackagingIsStandard

`func (o *IssueCategoriesPutPreviousValues) GetPackagingIsStandard() bool`

GetPackagingIsStandard returns the PackagingIsStandard field if non-nil, zero value otherwise.

### GetPackagingIsStandardOk

`func (o *IssueCategoriesPutPreviousValues) GetPackagingIsStandardOk() (*bool, bool)`

GetPackagingIsStandardOk returns a tuple with the PackagingIsStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagingIsStandard

`func (o *IssueCategoriesPutPreviousValues) SetPackagingIsStandard(v bool)`

SetPackagingIsStandard sets PackagingIsStandard field to given value.

### HasPackagingIsStandard

`func (o *IssueCategoriesPutPreviousValues) HasPackagingIsStandard() bool`

HasPackagingIsStandard returns a boolean if a field has been set.

### GetIssueRequiredFields

`func (o *IssueCategoriesPutPreviousValues) GetIssueRequiredFields() interface{}`

GetIssueRequiredFields returns the IssueRequiredFields field if non-nil, zero value otherwise.

### GetIssueRequiredFieldsOk

`func (o *IssueCategoriesPutPreviousValues) GetIssueRequiredFieldsOk() (*interface{}, bool)`

GetIssueRequiredFieldsOk returns a tuple with the IssueRequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueRequiredFields

`func (o *IssueCategoriesPutPreviousValues) SetIssueRequiredFields(v interface{})`

SetIssueRequiredFields sets IssueRequiredFields field to given value.

### HasIssueRequiredFields

`func (o *IssueCategoriesPutPreviousValues) HasIssueRequiredFields() bool`

HasIssueRequiredFields returns a boolean if a field has been set.

### SetIssueRequiredFieldsNil

`func (o *IssueCategoriesPutPreviousValues) SetIssueRequiredFieldsNil(b bool)`

 SetIssueRequiredFieldsNil sets the value for IssueRequiredFields to be an explicit nil

### UnsetIssueRequiredFields
`func (o *IssueCategoriesPutPreviousValues) UnsetIssueRequiredFields()`

UnsetIssueRequiredFields ensures that no value is present for IssueRequiredFields, not even an explicit nil
### GetActionPlanRequiredFields

`func (o *IssueCategoriesPutPreviousValues) GetActionPlanRequiredFields() interface{}`

GetActionPlanRequiredFields returns the ActionPlanRequiredFields field if non-nil, zero value otherwise.

### GetActionPlanRequiredFieldsOk

`func (o *IssueCategoriesPutPreviousValues) GetActionPlanRequiredFieldsOk() (*interface{}, bool)`

GetActionPlanRequiredFieldsOk returns a tuple with the ActionPlanRequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPlanRequiredFields

`func (o *IssueCategoriesPutPreviousValues) SetActionPlanRequiredFields(v interface{})`

SetActionPlanRequiredFields sets ActionPlanRequiredFields field to given value.

### HasActionPlanRequiredFields

`func (o *IssueCategoriesPutPreviousValues) HasActionPlanRequiredFields() bool`

HasActionPlanRequiredFields returns a boolean if a field has been set.

### SetActionPlanRequiredFieldsNil

`func (o *IssueCategoriesPutPreviousValues) SetActionPlanRequiredFieldsNil(b bool)`

 SetActionPlanRequiredFieldsNil sets the value for ActionPlanRequiredFields to be an explicit nil

### UnsetActionPlanRequiredFields
`func (o *IssueCategoriesPutPreviousValues) UnsetActionPlanRequiredFields()`

UnsetActionPlanRequiredFields ensures that no value is present for ActionPlanRequiredFields, not even an explicit nil
### GetScopes

`func (o *IssueCategoriesPutPreviousValues) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *IssueCategoriesPutPreviousValues) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *IssueCategoriesPutPreviousValues) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *IssueCategoriesPutPreviousValues) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *IssueCategoriesPutPreviousValues) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *IssueCategoriesPutPreviousValues) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


