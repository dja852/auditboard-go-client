# ActionPlansArchivesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**ActionPlanId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;action_plans.id&#x60;.&lt;fk table&#x3D;&#39;action_plans&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssuesArchiveId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issues_archives.id&#x60;.&lt;fk table&#x3D;&#39;issues_archives&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ArchiveId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;archives.id&#x60;.&lt;fk table&#x3D;&#39;archives&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**IssueId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issues.id&#x60;.&lt;fk table&#x3D;&#39;issues&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Flattened** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewActionPlansArchivesPutPreviousValues

`func NewActionPlansArchivesPutPreviousValues() *ActionPlansArchivesPutPreviousValues`

NewActionPlansArchivesPutPreviousValues instantiates a new ActionPlansArchivesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionPlansArchivesPutPreviousValuesWithDefaults

`func NewActionPlansArchivesPutPreviousValuesWithDefaults() *ActionPlansArchivesPutPreviousValues`

NewActionPlansArchivesPutPreviousValuesWithDefaults instantiates a new ActionPlansArchivesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionPlansArchivesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionPlansArchivesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionPlansArchivesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ActionPlansArchivesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActionPlanId

`func (o *ActionPlansArchivesPutPreviousValues) GetActionPlanId() int32`

GetActionPlanId returns the ActionPlanId field if non-nil, zero value otherwise.

### GetActionPlanIdOk

`func (o *ActionPlansArchivesPutPreviousValues) GetActionPlanIdOk() (*int32, bool)`

GetActionPlanIdOk returns a tuple with the ActionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPlanId

`func (o *ActionPlansArchivesPutPreviousValues) SetActionPlanId(v int32)`

SetActionPlanId sets ActionPlanId field to given value.

### HasActionPlanId

`func (o *ActionPlansArchivesPutPreviousValues) HasActionPlanId() bool`

HasActionPlanId returns a boolean if a field has been set.

### GetIssuesArchiveId

`func (o *ActionPlansArchivesPutPreviousValues) GetIssuesArchiveId() int32`

GetIssuesArchiveId returns the IssuesArchiveId field if non-nil, zero value otherwise.

### GetIssuesArchiveIdOk

`func (o *ActionPlansArchivesPutPreviousValues) GetIssuesArchiveIdOk() (*int32, bool)`

GetIssuesArchiveIdOk returns a tuple with the IssuesArchiveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesArchiveId

`func (o *ActionPlansArchivesPutPreviousValues) SetIssuesArchiveId(v int32)`

SetIssuesArchiveId sets IssuesArchiveId field to given value.

### HasIssuesArchiveId

`func (o *ActionPlansArchivesPutPreviousValues) HasIssuesArchiveId() bool`

HasIssuesArchiveId returns a boolean if a field has been set.

### GetArchiveId

`func (o *ActionPlansArchivesPutPreviousValues) GetArchiveId() int32`

GetArchiveId returns the ArchiveId field if non-nil, zero value otherwise.

### GetArchiveIdOk

`func (o *ActionPlansArchivesPutPreviousValues) GetArchiveIdOk() (*int32, bool)`

GetArchiveIdOk returns a tuple with the ArchiveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveId

`func (o *ActionPlansArchivesPutPreviousValues) SetArchiveId(v int32)`

SetArchiveId sets ArchiveId field to given value.

### HasArchiveId

`func (o *ActionPlansArchivesPutPreviousValues) HasArchiveId() bool`

HasArchiveId returns a boolean if a field has been set.

### GetIssueId

`func (o *ActionPlansArchivesPutPreviousValues) GetIssueId() int32`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *ActionPlansArchivesPutPreviousValues) GetIssueIdOk() (*int32, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *ActionPlansArchivesPutPreviousValues) SetIssueId(v int32)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *ActionPlansArchivesPutPreviousValues) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.

### GetFlattened

`func (o *ActionPlansArchivesPutPreviousValues) GetFlattened() interface{}`

GetFlattened returns the Flattened field if non-nil, zero value otherwise.

### GetFlattenedOk

`func (o *ActionPlansArchivesPutPreviousValues) GetFlattenedOk() (*interface{}, bool)`

GetFlattenedOk returns a tuple with the Flattened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlattened

`func (o *ActionPlansArchivesPutPreviousValues) SetFlattened(v interface{})`

SetFlattened sets Flattened field to given value.

### HasFlattened

`func (o *ActionPlansArchivesPutPreviousValues) HasFlattened() bool`

HasFlattened returns a boolean if a field has been set.

### SetFlattenedNil

`func (o *ActionPlansArchivesPutPreviousValues) SetFlattenedNil(b bool)`

 SetFlattenedNil sets the value for Flattened to be an explicit nil

### UnsetFlattened
`func (o *ActionPlansArchivesPutPreviousValues) UnsetFlattened()`

UnsetFlattened ensures that no value is present for Flattened, not even an explicit nil
### GetCreatedAt

`func (o *ActionPlansArchivesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ActionPlansArchivesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ActionPlansArchivesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ActionPlansArchivesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ActionPlansArchivesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ActionPlansArchivesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ActionPlansArchivesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ActionPlansArchivesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ActionPlansArchivesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ActionPlansArchivesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ActionPlansArchivesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ActionPlansArchivesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


