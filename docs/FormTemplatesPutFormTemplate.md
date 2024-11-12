# FormTemplatesPutFormTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Name** | **string** |  | [default to ""]
**Type** | **string** |  | [default to ""]
**TemplateJson** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewFormTemplatesPutFormTemplate

`func NewFormTemplatesPutFormTemplate(name string, type_ string, ) *FormTemplatesPutFormTemplate`

NewFormTemplatesPutFormTemplate instantiates a new FormTemplatesPutFormTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormTemplatesPutFormTemplateWithDefaults

`func NewFormTemplatesPutFormTemplateWithDefaults() *FormTemplatesPutFormTemplate`

NewFormTemplatesPutFormTemplateWithDefaults instantiates a new FormTemplatesPutFormTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FormTemplatesPutFormTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormTemplatesPutFormTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormTemplatesPutFormTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FormTemplatesPutFormTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FormTemplatesPutFormTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormTemplatesPutFormTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormTemplatesPutFormTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *FormTemplatesPutFormTemplate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormTemplatesPutFormTemplate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormTemplatesPutFormTemplate) SetType(v string)`

SetType sets Type field to given value.


### GetTemplateJson

`func (o *FormTemplatesPutFormTemplate) GetTemplateJson() interface{}`

GetTemplateJson returns the TemplateJson field if non-nil, zero value otherwise.

### GetTemplateJsonOk

`func (o *FormTemplatesPutFormTemplate) GetTemplateJsonOk() (*interface{}, bool)`

GetTemplateJsonOk returns a tuple with the TemplateJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateJson

`func (o *FormTemplatesPutFormTemplate) SetTemplateJson(v interface{})`

SetTemplateJson sets TemplateJson field to given value.

### HasTemplateJson

`func (o *FormTemplatesPutFormTemplate) HasTemplateJson() bool`

HasTemplateJson returns a boolean if a field has been set.

### SetTemplateJsonNil

`func (o *FormTemplatesPutFormTemplate) SetTemplateJsonNil(b bool)`

 SetTemplateJsonNil sets the value for TemplateJson to be an explicit nil

### UnsetTemplateJson
`func (o *FormTemplatesPutFormTemplate) UnsetTemplateJson()`

UnsetTemplateJson ensures that no value is present for TemplateJson, not even an explicit nil
### GetCreatedAt

`func (o *FormTemplatesPutFormTemplate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FormTemplatesPutFormTemplate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FormTemplatesPutFormTemplate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FormTemplatesPutFormTemplate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FormTemplatesPutFormTemplate) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FormTemplatesPutFormTemplate) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FormTemplatesPutFormTemplate) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FormTemplatesPutFormTemplate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *FormTemplatesPutFormTemplate) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *FormTemplatesPutFormTemplate) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *FormTemplatesPutFormTemplate) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *FormTemplatesPutFormTemplate) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


