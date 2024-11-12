# ActionPlansArchives

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**ActionPlanId** | **int32** | Note: This is a Foreign Key to &#x60;action_plans.id&#x60;.&lt;fk table&#x3D;&#39;action_plans&#39; column&#x3D;&#39;id&#39;/&gt; | 
**IssuesArchiveId** | **int32** | Note: This is a Foreign Key to &#x60;issues_archives.id&#x60;.&lt;fk table&#x3D;&#39;issues_archives&#39; column&#x3D;&#39;id&#39;/&gt; | 
**ArchiveId** | **int32** | Note: This is a Foreign Key to &#x60;archives.id&#x60;.&lt;fk table&#x3D;&#39;archives&#39; column&#x3D;&#39;id&#39;/&gt; | 
**IssueId** | **int32** | Note: This is a Foreign Key to &#x60;issues.id&#x60;.&lt;fk table&#x3D;&#39;issues&#39; column&#x3D;&#39;id&#39;/&gt; | 
**Flattened** | **interface{}** |  | 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewActionPlansArchives

`func NewActionPlansArchives(actionPlanId int32, issuesArchiveId int32, archiveId int32, issueId int32, flattened interface{}, ) *ActionPlansArchives`

NewActionPlansArchives instantiates a new ActionPlansArchives object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionPlansArchivesWithDefaults

`func NewActionPlansArchivesWithDefaults() *ActionPlansArchives`

NewActionPlansArchivesWithDefaults instantiates a new ActionPlansArchives object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionPlansArchives) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionPlansArchives) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionPlansArchives) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ActionPlansArchives) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActionPlanId

`func (o *ActionPlansArchives) GetActionPlanId() int32`

GetActionPlanId returns the ActionPlanId field if non-nil, zero value otherwise.

### GetActionPlanIdOk

`func (o *ActionPlansArchives) GetActionPlanIdOk() (*int32, bool)`

GetActionPlanIdOk returns a tuple with the ActionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPlanId

`func (o *ActionPlansArchives) SetActionPlanId(v int32)`

SetActionPlanId sets ActionPlanId field to given value.


### GetIssuesArchiveId

`func (o *ActionPlansArchives) GetIssuesArchiveId() int32`

GetIssuesArchiveId returns the IssuesArchiveId field if non-nil, zero value otherwise.

### GetIssuesArchiveIdOk

`func (o *ActionPlansArchives) GetIssuesArchiveIdOk() (*int32, bool)`

GetIssuesArchiveIdOk returns a tuple with the IssuesArchiveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesArchiveId

`func (o *ActionPlansArchives) SetIssuesArchiveId(v int32)`

SetIssuesArchiveId sets IssuesArchiveId field to given value.


### GetArchiveId

`func (o *ActionPlansArchives) GetArchiveId() int32`

GetArchiveId returns the ArchiveId field if non-nil, zero value otherwise.

### GetArchiveIdOk

`func (o *ActionPlansArchives) GetArchiveIdOk() (*int32, bool)`

GetArchiveIdOk returns a tuple with the ArchiveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveId

`func (o *ActionPlansArchives) SetArchiveId(v int32)`

SetArchiveId sets ArchiveId field to given value.


### GetIssueId

`func (o *ActionPlansArchives) GetIssueId() int32`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *ActionPlansArchives) GetIssueIdOk() (*int32, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *ActionPlansArchives) SetIssueId(v int32)`

SetIssueId sets IssueId field to given value.


### GetFlattened

`func (o *ActionPlansArchives) GetFlattened() interface{}`

GetFlattened returns the Flattened field if non-nil, zero value otherwise.

### GetFlattenedOk

`func (o *ActionPlansArchives) GetFlattenedOk() (*interface{}, bool)`

GetFlattenedOk returns a tuple with the Flattened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlattened

`func (o *ActionPlansArchives) SetFlattened(v interface{})`

SetFlattened sets Flattened field to given value.


### SetFlattenedNil

`func (o *ActionPlansArchives) SetFlattenedNil(b bool)`

 SetFlattenedNil sets the value for Flattened to be an explicit nil

### UnsetFlattened
`func (o *ActionPlansArchives) UnsetFlattened()`

UnsetFlattened ensures that no value is present for Flattened, not even an explicit nil
### GetCreatedAt

`func (o *ActionPlansArchives) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ActionPlansArchives) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ActionPlansArchives) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ActionPlansArchives) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ActionPlansArchives) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ActionPlansArchives) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ActionPlansArchives) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ActionPlansArchives) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ActionPlansArchives) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ActionPlansArchives) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ActionPlansArchives) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ActionPlansArchives) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


