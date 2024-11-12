# PolicyPublishedVersions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**PolicyId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;policies.id&#x60;.&lt;fk table&#x3D;&#39;policies&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Flattened** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPolicyPublishedVersions

`func NewPolicyPublishedVersions() *PolicyPublishedVersions`

NewPolicyPublishedVersions instantiates a new PolicyPublishedVersions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyPublishedVersionsWithDefaults

`func NewPolicyPublishedVersionsWithDefaults() *PolicyPublishedVersions`

NewPolicyPublishedVersionsWithDefaults instantiates a new PolicyPublishedVersions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PolicyPublishedVersions) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyPublishedVersions) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyPublishedVersions) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyPublishedVersions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PolicyPublishedVersions) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PolicyPublishedVersions) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PolicyPublishedVersions) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PolicyPublishedVersions) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PolicyPublishedVersions) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PolicyPublishedVersions) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PolicyPublishedVersions) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PolicyPublishedVersions) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *PolicyPublishedVersions) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *PolicyPublishedVersions) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *PolicyPublishedVersions) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *PolicyPublishedVersions) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetPolicyId

`func (o *PolicyPublishedVersions) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *PolicyPublishedVersions) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *PolicyPublishedVersions) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *PolicyPublishedVersions) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetFlattened

`func (o *PolicyPublishedVersions) GetFlattened() interface{}`

GetFlattened returns the Flattened field if non-nil, zero value otherwise.

### GetFlattenedOk

`func (o *PolicyPublishedVersions) GetFlattenedOk() (*interface{}, bool)`

GetFlattenedOk returns a tuple with the Flattened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlattened

`func (o *PolicyPublishedVersions) SetFlattened(v interface{})`

SetFlattened sets Flattened field to given value.

### HasFlattened

`func (o *PolicyPublishedVersions) HasFlattened() bool`

HasFlattened returns a boolean if a field has been set.

### SetFlattenedNil

`func (o *PolicyPublishedVersions) SetFlattenedNil(b bool)`

 SetFlattenedNil sets the value for Flattened to be an explicit nil

### UnsetFlattened
`func (o *PolicyPublishedVersions) UnsetFlattened()`

UnsetFlattened ensures that no value is present for Flattened, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


