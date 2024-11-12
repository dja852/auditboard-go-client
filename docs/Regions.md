# Regions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**SortOrder** | **int32** |  | 
**RegionCode** | **string** |  | 
**Name** | **string** |  | 
**WorkspaceId** | **int32** | Note: This is a Foreign Key to &#x60;workspaces.id&#x60;.&lt;fk table&#x3D;&#39;workspaces&#39; column&#x3D;&#39;id&#39;/&gt; | 
**FormTemplates** | Pointer to **interface{}** |  | [optional] 
**EnabledAttributes** | Pointer to **interface{}** |  | [optional] 
**ExcludedAttributes** | Pointer to **interface{}** |  | [optional] 
**Scopes** | **interface{}** |  | 
**Options** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewRegions

`func NewRegions(sortOrder int32, regionCode string, name string, workspaceId int32, scopes interface{}, ) *Regions`

NewRegions instantiates a new Regions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionsWithDefaults

`func NewRegionsWithDefaults() *Regions`

NewRegionsWithDefaults instantiates a new Regions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Regions) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Regions) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Regions) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Regions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Regions) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Regions) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Regions) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Regions) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Regions) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Regions) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Regions) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Regions) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Regions) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Regions) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Regions) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Regions) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetSortOrder

`func (o *Regions) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *Regions) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *Regions) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.


### GetRegionCode

`func (o *Regions) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *Regions) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *Regions) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.


### GetName

`func (o *Regions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Regions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Regions) SetName(v string)`

SetName sets Name field to given value.


### GetWorkspaceId

`func (o *Regions) GetWorkspaceId() int32`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *Regions) GetWorkspaceIdOk() (*int32, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *Regions) SetWorkspaceId(v int32)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetFormTemplates

`func (o *Regions) GetFormTemplates() interface{}`

GetFormTemplates returns the FormTemplates field if non-nil, zero value otherwise.

### GetFormTemplatesOk

`func (o *Regions) GetFormTemplatesOk() (*interface{}, bool)`

GetFormTemplatesOk returns a tuple with the FormTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormTemplates

`func (o *Regions) SetFormTemplates(v interface{})`

SetFormTemplates sets FormTemplates field to given value.

### HasFormTemplates

`func (o *Regions) HasFormTemplates() bool`

HasFormTemplates returns a boolean if a field has been set.

### SetFormTemplatesNil

`func (o *Regions) SetFormTemplatesNil(b bool)`

 SetFormTemplatesNil sets the value for FormTemplates to be an explicit nil

### UnsetFormTemplates
`func (o *Regions) UnsetFormTemplates()`

UnsetFormTemplates ensures that no value is present for FormTemplates, not even an explicit nil
### GetEnabledAttributes

`func (o *Regions) GetEnabledAttributes() interface{}`

GetEnabledAttributes returns the EnabledAttributes field if non-nil, zero value otherwise.

### GetEnabledAttributesOk

`func (o *Regions) GetEnabledAttributesOk() (*interface{}, bool)`

GetEnabledAttributesOk returns a tuple with the EnabledAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAttributes

`func (o *Regions) SetEnabledAttributes(v interface{})`

SetEnabledAttributes sets EnabledAttributes field to given value.

### HasEnabledAttributes

`func (o *Regions) HasEnabledAttributes() bool`

HasEnabledAttributes returns a boolean if a field has been set.

### SetEnabledAttributesNil

`func (o *Regions) SetEnabledAttributesNil(b bool)`

 SetEnabledAttributesNil sets the value for EnabledAttributes to be an explicit nil

### UnsetEnabledAttributes
`func (o *Regions) UnsetEnabledAttributes()`

UnsetEnabledAttributes ensures that no value is present for EnabledAttributes, not even an explicit nil
### GetExcludedAttributes

`func (o *Regions) GetExcludedAttributes() interface{}`

GetExcludedAttributes returns the ExcludedAttributes field if non-nil, zero value otherwise.

### GetExcludedAttributesOk

`func (o *Regions) GetExcludedAttributesOk() (*interface{}, bool)`

GetExcludedAttributesOk returns a tuple with the ExcludedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedAttributes

`func (o *Regions) SetExcludedAttributes(v interface{})`

SetExcludedAttributes sets ExcludedAttributes field to given value.

### HasExcludedAttributes

`func (o *Regions) HasExcludedAttributes() bool`

HasExcludedAttributes returns a boolean if a field has been set.

### SetExcludedAttributesNil

`func (o *Regions) SetExcludedAttributesNil(b bool)`

 SetExcludedAttributesNil sets the value for ExcludedAttributes to be an explicit nil

### UnsetExcludedAttributes
`func (o *Regions) UnsetExcludedAttributes()`

UnsetExcludedAttributes ensures that no value is present for ExcludedAttributes, not even an explicit nil
### GetScopes

`func (o *Regions) GetScopes() interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *Regions) GetScopesOk() (*interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *Regions) SetScopes(v interface{})`

SetScopes sets Scopes field to given value.


### SetScopesNil

`func (o *Regions) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *Regions) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetOptions

`func (o *Regions) GetOptions() interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Regions) GetOptionsOk() (*interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Regions) SetOptions(v interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Regions) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *Regions) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *Regions) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


