# ControlsSegmentsPutControlsSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Name** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**FormTemplateId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;form_templates.id&#x60;.&lt;fk table&#x3D;&#39;form_templates&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 

## Methods

### NewControlsSegmentsPutControlsSegment

`func NewControlsSegmentsPutControlsSegment(name string, ) *ControlsSegmentsPutControlsSegment`

NewControlsSegmentsPutControlsSegment instantiates a new ControlsSegmentsPutControlsSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlsSegmentsPutControlsSegmentWithDefaults

`func NewControlsSegmentsPutControlsSegmentWithDefaults() *ControlsSegmentsPutControlsSegment`

NewControlsSegmentsPutControlsSegmentWithDefaults instantiates a new ControlsSegmentsPutControlsSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ControlsSegmentsPutControlsSegment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControlsSegmentsPutControlsSegment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControlsSegmentsPutControlsSegment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ControlsSegmentsPutControlsSegment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ControlsSegmentsPutControlsSegment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControlsSegmentsPutControlsSegment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControlsSegmentsPutControlsSegment) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ControlsSegmentsPutControlsSegment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ControlsSegmentsPutControlsSegment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ControlsSegmentsPutControlsSegment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ControlsSegmentsPutControlsSegment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ControlsSegmentsPutControlsSegment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ControlsSegmentsPutControlsSegment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ControlsSegmentsPutControlsSegment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ControlsSegmentsPutControlsSegment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ControlsSegmentsPutControlsSegment) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ControlsSegmentsPutControlsSegment) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ControlsSegmentsPutControlsSegment) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ControlsSegmentsPutControlsSegment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ControlsSegmentsPutControlsSegment) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ControlsSegmentsPutControlsSegment) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ControlsSegmentsPutControlsSegment) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ControlsSegmentsPutControlsSegment) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetFormTemplateId

`func (o *ControlsSegmentsPutControlsSegment) GetFormTemplateId() int32`

GetFormTemplateId returns the FormTemplateId field if non-nil, zero value otherwise.

### GetFormTemplateIdOk

`func (o *ControlsSegmentsPutControlsSegment) GetFormTemplateIdOk() (*int32, bool)`

GetFormTemplateIdOk returns a tuple with the FormTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormTemplateId

`func (o *ControlsSegmentsPutControlsSegment) SetFormTemplateId(v int32)`

SetFormTemplateId sets FormTemplateId field to given value.

### HasFormTemplateId

`func (o *ControlsSegmentsPutControlsSegment) HasFormTemplateId() bool`

HasFormTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


