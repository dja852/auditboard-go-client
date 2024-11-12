# RegionsPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**SortOrder** | Pointer to **int32** |  | [optional] 
**RegionCode** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**WorkspaceId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;workspaces.id&#x60;.&lt;fk table&#x3D;&#39;workspaces&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**FormTemplates** | Pointer to **interface{}** |  | [optional] 
**EnabledAttributes** | Pointer to **interface{}** |  | [optional] 
**ExcludedAttributes** | Pointer to **interface{}** |  | [optional] 
**Scopes** | Pointer to **interface{}** |  | [optional] 
**Options** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewRegionsPutPreviousValues

`func NewRegionsPutPreviousValues() *RegionsPutPreviousValues`

NewRegionsPutPreviousValues instantiates a new RegionsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionsPutPreviousValuesWithDefaults

`func NewRegionsPutPreviousValuesWithDefaults() *RegionsPutPreviousValues`

NewRegionsPutPreviousValuesWithDefaults instantiates a new RegionsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegionsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegionsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegionsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RegionsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RegionsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RegionsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RegionsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RegionsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RegionsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RegionsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RegionsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RegionsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *RegionsPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *RegionsPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *RegionsPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *RegionsPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetSortOrder

`func (o *RegionsPutPreviousValues) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *RegionsPutPreviousValues) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *RegionsPutPreviousValues) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *RegionsPutPreviousValues) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetRegionCode

`func (o *RegionsPutPreviousValues) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *RegionsPutPreviousValues) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *RegionsPutPreviousValues) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *RegionsPutPreviousValues) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetName

`func (o *RegionsPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegionsPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegionsPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegionsPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *RegionsPutPreviousValues) GetWorkspaceId() int32`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *RegionsPutPreviousValues) GetWorkspaceIdOk() (*int32, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *RegionsPutPreviousValues) SetWorkspaceId(v int32)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *RegionsPutPreviousValues) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### GetFormTemplates

`func (o *RegionsPutPreviousValues) GetFormTemplates() interface{}`

GetFormTemplates returns the FormTemplates field if non-nil, zero value otherwise.

### GetFormTemplatesOk

`func (o *RegionsPutPreviousValues) GetFormTemplatesOk() (*interface{}, bool)`

GetFormTemplatesOk returns a tuple with the FormTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormTemplates

`func (o *RegionsPutPreviousValues) SetFormTemplates(v interface{})`

SetFormTemplates sets FormTemplates field to given value.

### HasFormTemplates

`func (o *RegionsPutPreviousValues) HasFormTemplates() bool`

HasFormTemplates returns a boolean if a field has been set.

### SetFormTemplatesNil

`func (o *RegionsPutPreviousValues) SetFormTemplatesNil(b bool)`

 SetFormTemplatesNil sets the value for FormTemplates to be an explicit nil

### UnsetFormTemplates
`func (o *RegionsPutPreviousValues) UnsetFormTemplates()`

UnsetFormTemplates ensures that no value is present for FormTemplates, not even an explicit nil
### GetEnabledAttributes

`func (o *RegionsPutPreviousValues) GetEnabledAttributes() interface{}`

GetEnabledAttributes returns the EnabledAttributes field if non-nil, zero value otherwise.

### GetEnabledAttributesOk

`func (o *RegionsPutPreviousValues) GetEnabledAttributesOk() (*interface{}, bool)`

GetEnabledAttributesOk returns a tuple with the EnabledAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAttributes

`func (o *RegionsPutPreviousValues) SetEnabledAttributes(v interface{})`

SetEnabledAttributes sets EnabledAttributes field to given value.

### HasEnabledAttributes

`func (o *RegionsPutPreviousValues) HasEnabledAttributes() bool`

HasEnabledAttributes returns a boolean if a field has been set.

### SetEnabledAttributesNil

`func (o *RegionsPutPreviousValues) SetEnabledAttributesNil(b bool)`

 SetEnabledAttributesNil sets the value for EnabledAttributes to be an explicit nil

### UnsetEnabledAttributes
`func (o *RegionsPutPreviousValues) UnsetEnabledAttributes()`

UnsetEnabledAttributes ensures that no value is present for EnabledAttributes, not even an explicit nil
### GetExcludedAttributes

`func (o *RegionsPutPreviousValues) GetExcludedAttributes() interface{}`

GetExcludedAttributes returns the ExcludedAttributes field if non-nil, zero value otherwise.

### GetExcludedAttributesOk

`func (o *RegionsPutPreviousValues) GetExcludedAttributesOk() (*interface{}, bool)`

GetExcludedAttributesOk returns a tuple with the ExcludedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedAttributes

`func (o *RegionsPutPreviousValues) SetExcludedAttributes(v interface{})`

SetExcludedAttributes sets ExcludedAttributes field to given value.

### HasExcludedAttributes

`func (o *RegionsPutPreviousValues) HasExcludedAttributes() bool`

HasExcludedAttributes returns a boolean if a field has been set.

### SetExcludedAttributesNil

`func (o *RegionsPutPreviousValues) SetExcludedAttributesNil(b bool)`

 SetExcludedAttributesNil sets the value for ExcludedAttributes to be an explicit nil

### UnsetExcludedAttributes
`func (o *RegionsPutPreviousValues) UnsetExcludedAttributes()`

UnsetExcludedAttributes ensures that no value is present for ExcludedAttributes, not even an explicit nil
### GetScopes

`func (o *RegionsPutPreviousValues) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *RegionsPutPreviousValues) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *RegionsPutPreviousValues) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *RegionsPutPreviousValues) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *RegionsPutPreviousValues) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *RegionsPutPreviousValues) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetOptions

`func (o *RegionsPutPreviousValues) GetOptions() interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *RegionsPutPreviousValues) GetOptionsOk() (*interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *RegionsPutPreviousValues) SetOptions(v interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *RegionsPutPreviousValues) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *RegionsPutPreviousValues) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *RegionsPutPreviousValues) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


