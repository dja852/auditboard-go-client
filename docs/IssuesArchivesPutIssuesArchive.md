# IssuesArchivesPutIssuesArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Type** | **string** |  | 
**Status** | Pointer to **string** |  | [optional] 
**IssueId** | **int32** | Note: This is a Foreign Key to &#x60;issues.id&#x60;.&lt;fk table&#x3D;&#39;issues&#39; column&#x3D;&#39;id&#39;/&gt; | 
**ArchiveId** | **int32** | Note: This is a Foreign Key to &#x60;archives.id&#x60;.&lt;fk table&#x3D;&#39;archives&#39; column&#x3D;&#39;id&#39;/&gt; | 
**ControlsDataArchiveId** | **int32** | Note: This is a Foreign Key to &#x60;controls_data_archives.id&#x60;.&lt;fk table&#x3D;&#39;controls_data_archives&#39; column&#x3D;&#39;id&#39;/&gt; | 
**CreatorUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlsDatumId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;controls_data.id&#x60;.&lt;fk table&#x3D;&#39;controls_data&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;tests.id&#x60;.&lt;fk table&#x3D;&#39;tests&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestSectionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_sections.id&#x60;.&lt;fk table&#x3D;&#39;test_sections&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ReferenceIssueId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issues.id&#x60;.&lt;fk table&#x3D;&#39;issues&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Flattened** | **interface{}** |  | 
**ReferenceMeta** | Pointer to **interface{}** |  | [optional] 
**RemediationOwners** | **interface{}** |  | 
**Totals** | **interface{}** |  | 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**CreatedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 

## Methods

### NewIssuesArchivesPutIssuesArchive

`func NewIssuesArchivesPutIssuesArchive(type_ string, issueId int32, archiveId int32, controlsDataArchiveId int32, flattened interface{}, remediationOwners interface{}, totals interface{}, ) *IssuesArchivesPutIssuesArchive`

NewIssuesArchivesPutIssuesArchive instantiates a new IssuesArchivesPutIssuesArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuesArchivesPutIssuesArchiveWithDefaults

`func NewIssuesArchivesPutIssuesArchiveWithDefaults() *IssuesArchivesPutIssuesArchive`

NewIssuesArchivesPutIssuesArchiveWithDefaults instantiates a new IssuesArchivesPutIssuesArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssuesArchivesPutIssuesArchive) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssuesArchivesPutIssuesArchive) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssuesArchivesPutIssuesArchive) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IssuesArchivesPutIssuesArchive) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *IssuesArchivesPutIssuesArchive) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IssuesArchivesPutIssuesArchive) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IssuesArchivesPutIssuesArchive) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *IssuesArchivesPutIssuesArchive) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IssuesArchivesPutIssuesArchive) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IssuesArchivesPutIssuesArchive) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IssuesArchivesPutIssuesArchive) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIssueId

`func (o *IssuesArchivesPutIssuesArchive) GetIssueId() int32`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *IssuesArchivesPutIssuesArchive) GetIssueIdOk() (*int32, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *IssuesArchivesPutIssuesArchive) SetIssueId(v int32)`

SetIssueId sets IssueId field to given value.


### GetArchiveId

`func (o *IssuesArchivesPutIssuesArchive) GetArchiveId() int32`

GetArchiveId returns the ArchiveId field if non-nil, zero value otherwise.

### GetArchiveIdOk

`func (o *IssuesArchivesPutIssuesArchive) GetArchiveIdOk() (*int32, bool)`

GetArchiveIdOk returns a tuple with the ArchiveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveId

`func (o *IssuesArchivesPutIssuesArchive) SetArchiveId(v int32)`

SetArchiveId sets ArchiveId field to given value.


### GetControlsDataArchiveId

`func (o *IssuesArchivesPutIssuesArchive) GetControlsDataArchiveId() int32`

GetControlsDataArchiveId returns the ControlsDataArchiveId field if non-nil, zero value otherwise.

### GetControlsDataArchiveIdOk

`func (o *IssuesArchivesPutIssuesArchive) GetControlsDataArchiveIdOk() (*int32, bool)`

GetControlsDataArchiveIdOk returns a tuple with the ControlsDataArchiveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlsDataArchiveId

`func (o *IssuesArchivesPutIssuesArchive) SetControlsDataArchiveId(v int32)`

SetControlsDataArchiveId sets ControlsDataArchiveId field to given value.


### GetCreatorUserId

`func (o *IssuesArchivesPutIssuesArchive) GetCreatorUserId() int32`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *IssuesArchivesPutIssuesArchive) GetCreatorUserIdOk() (*int32, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *IssuesArchivesPutIssuesArchive) SetCreatorUserId(v int32)`

SetCreatorUserId sets CreatorUserId field to given value.

### HasCreatorUserId

`func (o *IssuesArchivesPutIssuesArchive) HasCreatorUserId() bool`

HasCreatorUserId returns a boolean if a field has been set.

### GetControlsDatumId

`func (o *IssuesArchivesPutIssuesArchive) GetControlsDatumId() int32`

GetControlsDatumId returns the ControlsDatumId field if non-nil, zero value otherwise.

### GetControlsDatumIdOk

`func (o *IssuesArchivesPutIssuesArchive) GetControlsDatumIdOk() (*int32, bool)`

GetControlsDatumIdOk returns a tuple with the ControlsDatumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlsDatumId

`func (o *IssuesArchivesPutIssuesArchive) SetControlsDatumId(v int32)`

SetControlsDatumId sets ControlsDatumId field to given value.

### HasControlsDatumId

`func (o *IssuesArchivesPutIssuesArchive) HasControlsDatumId() bool`

HasControlsDatumId returns a boolean if a field has been set.

### GetTestId

`func (o *IssuesArchivesPutIssuesArchive) GetTestId() int32`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *IssuesArchivesPutIssuesArchive) GetTestIdOk() (*int32, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *IssuesArchivesPutIssuesArchive) SetTestId(v int32)`

SetTestId sets TestId field to given value.

### HasTestId

`func (o *IssuesArchivesPutIssuesArchive) HasTestId() bool`

HasTestId returns a boolean if a field has been set.

### GetTestSectionId

`func (o *IssuesArchivesPutIssuesArchive) GetTestSectionId() int32`

GetTestSectionId returns the TestSectionId field if non-nil, zero value otherwise.

### GetTestSectionIdOk

`func (o *IssuesArchivesPutIssuesArchive) GetTestSectionIdOk() (*int32, bool)`

GetTestSectionIdOk returns a tuple with the TestSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSectionId

`func (o *IssuesArchivesPutIssuesArchive) SetTestSectionId(v int32)`

SetTestSectionId sets TestSectionId field to given value.

### HasTestSectionId

`func (o *IssuesArchivesPutIssuesArchive) HasTestSectionId() bool`

HasTestSectionId returns a boolean if a field has been set.

### GetReferenceIssueId

`func (o *IssuesArchivesPutIssuesArchive) GetReferenceIssueId() int32`

GetReferenceIssueId returns the ReferenceIssueId field if non-nil, zero value otherwise.

### GetReferenceIssueIdOk

`func (o *IssuesArchivesPutIssuesArchive) GetReferenceIssueIdOk() (*int32, bool)`

GetReferenceIssueIdOk returns a tuple with the ReferenceIssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceIssueId

`func (o *IssuesArchivesPutIssuesArchive) SetReferenceIssueId(v int32)`

SetReferenceIssueId sets ReferenceIssueId field to given value.

### HasReferenceIssueId

`func (o *IssuesArchivesPutIssuesArchive) HasReferenceIssueId() bool`

HasReferenceIssueId returns a boolean if a field has been set.

### GetFlattened

`func (o *IssuesArchivesPutIssuesArchive) GetFlattened() interface{}`

GetFlattened returns the Flattened field if non-nil, zero value otherwise.

### GetFlattenedOk

`func (o *IssuesArchivesPutIssuesArchive) GetFlattenedOk() (*interface{}, bool)`

GetFlattenedOk returns a tuple with the Flattened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlattened

`func (o *IssuesArchivesPutIssuesArchive) SetFlattened(v interface{})`

SetFlattened sets Flattened field to given value.


### SetFlattenedNil

`func (o *IssuesArchivesPutIssuesArchive) SetFlattenedNil(b bool)`

 SetFlattenedNil sets the value for Flattened to be an explicit nil

### UnsetFlattened
`func (o *IssuesArchivesPutIssuesArchive) UnsetFlattened()`

UnsetFlattened ensures that no value is present for Flattened, not even an explicit nil
### GetReferenceMeta

`func (o *IssuesArchivesPutIssuesArchive) GetReferenceMeta() interface{}`

GetReferenceMeta returns the ReferenceMeta field if non-nil, zero value otherwise.

### GetReferenceMetaOk

`func (o *IssuesArchivesPutIssuesArchive) GetReferenceMetaOk() (*interface{}, bool)`

GetReferenceMetaOk returns a tuple with the ReferenceMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMeta

`func (o *IssuesArchivesPutIssuesArchive) SetReferenceMeta(v interface{})`

SetReferenceMeta sets ReferenceMeta field to given value.

### HasReferenceMeta

`func (o *IssuesArchivesPutIssuesArchive) HasReferenceMeta() bool`

HasReferenceMeta returns a boolean if a field has been set.

### SetReferenceMetaNil

`func (o *IssuesArchivesPutIssuesArchive) SetReferenceMetaNil(b bool)`

 SetReferenceMetaNil sets the value for ReferenceMeta to be an explicit nil

### UnsetReferenceMeta
`func (o *IssuesArchivesPutIssuesArchive) UnsetReferenceMeta()`

UnsetReferenceMeta ensures that no value is present for ReferenceMeta, not even an explicit nil
### GetRemediationOwners

`func (o *IssuesArchivesPutIssuesArchive) GetRemediationOwners() interface{}`

GetRemediationOwners returns the RemediationOwners field if non-nil, zero value otherwise.

### GetRemediationOwnersOk

`func (o *IssuesArchivesPutIssuesArchive) GetRemediationOwnersOk() (*interface{}, bool)`

GetRemediationOwnersOk returns a tuple with the RemediationOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationOwners

`func (o *IssuesArchivesPutIssuesArchive) SetRemediationOwners(v interface{})`

SetRemediationOwners sets RemediationOwners field to given value.


### SetRemediationOwnersNil

`func (o *IssuesArchivesPutIssuesArchive) SetRemediationOwnersNil(b bool)`

 SetRemediationOwnersNil sets the value for RemediationOwners to be an explicit nil

### UnsetRemediationOwners
`func (o *IssuesArchivesPutIssuesArchive) UnsetRemediationOwners()`

UnsetRemediationOwners ensures that no value is present for RemediationOwners, not even an explicit nil
### GetTotals

`func (o *IssuesArchivesPutIssuesArchive) GetTotals() interface{}`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *IssuesArchivesPutIssuesArchive) GetTotalsOk() (*interface{}, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *IssuesArchivesPutIssuesArchive) SetTotals(v interface{})`

SetTotals sets Totals field to given value.


### SetTotalsNil

`func (o *IssuesArchivesPutIssuesArchive) SetTotalsNil(b bool)`

 SetTotalsNil sets the value for Totals to be an explicit nil

### UnsetTotals
`func (o *IssuesArchivesPutIssuesArchive) UnsetTotals()`

UnsetTotals ensures that no value is present for Totals, not even an explicit nil
### GetCreatedAt

`func (o *IssuesArchivesPutIssuesArchive) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssuesArchivesPutIssuesArchive) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssuesArchivesPutIssuesArchive) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IssuesArchivesPutIssuesArchive) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IssuesArchivesPutIssuesArchive) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IssuesArchivesPutIssuesArchive) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IssuesArchivesPutIssuesArchive) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IssuesArchivesPutIssuesArchive) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *IssuesArchivesPutIssuesArchive) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *IssuesArchivesPutIssuesArchive) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *IssuesArchivesPutIssuesArchive) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *IssuesArchivesPutIssuesArchive) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *IssuesArchivesPutIssuesArchive) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *IssuesArchivesPutIssuesArchive) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *IssuesArchivesPutIssuesArchive) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *IssuesArchivesPutIssuesArchive) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


