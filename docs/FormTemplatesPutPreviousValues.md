# FormTemplatesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [default to ""]
**Type** | Pointer to **string** |  | [optional] [default to ""]
**TemplateJson** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewFormTemplatesPutPreviousValues

`func NewFormTemplatesPutPreviousValues() *FormTemplatesPutPreviousValues`

NewFormTemplatesPutPreviousValues instantiates a new FormTemplatesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormTemplatesPutPreviousValuesWithDefaults

`func NewFormTemplatesPutPreviousValuesWithDefaults() *FormTemplatesPutPreviousValues`

NewFormTemplatesPutPreviousValuesWithDefaults instantiates a new FormTemplatesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FormTemplatesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormTemplatesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormTemplatesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FormTemplatesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FormTemplatesPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormTemplatesPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormTemplatesPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FormTemplatesPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *FormTemplatesPutPreviousValues) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormTemplatesPutPreviousValues) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormTemplatesPutPreviousValues) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FormTemplatesPutPreviousValues) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTemplateJson

`func (o *FormTemplatesPutPreviousValues) GetTemplateJson() interface{}`

GetTemplateJson returns the TemplateJson field if non-nil, zero value otherwise.

### GetTemplateJsonOk

`func (o *FormTemplatesPutPreviousValues) GetTemplateJsonOk() (*interface{}, bool)`

GetTemplateJsonOk returns a tuple with the TemplateJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateJson

`func (o *FormTemplatesPutPreviousValues) SetTemplateJson(v interface{})`

SetTemplateJson sets TemplateJson field to given value.

### HasTemplateJson

`func (o *FormTemplatesPutPreviousValues) HasTemplateJson() bool`

HasTemplateJson returns a boolean if a field has been set.

### SetTemplateJsonNil

`func (o *FormTemplatesPutPreviousValues) SetTemplateJsonNil(b bool)`

 SetTemplateJsonNil sets the value for TemplateJson to be an explicit nil

### UnsetTemplateJson
`func (o *FormTemplatesPutPreviousValues) UnsetTemplateJson()`

UnsetTemplateJson ensures that no value is present for TemplateJson, not even an explicit nil
### GetCreatedAt

`func (o *FormTemplatesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FormTemplatesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FormTemplatesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FormTemplatesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FormTemplatesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FormTemplatesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FormTemplatesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FormTemplatesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *FormTemplatesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *FormTemplatesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *FormTemplatesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *FormTemplatesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


