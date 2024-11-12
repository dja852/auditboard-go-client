# IssuesArchivesPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**IssueId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issues.id&#x60;.&lt;fk table&#x3D;&#39;issues&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ArchiveId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;archives.id&#x60;.&lt;fk table&#x3D;&#39;archives&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlsDataArchiveId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;controls_data_archives.id&#x60;.&lt;fk table&#x3D;&#39;controls_data_archives&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**CreatorUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ControlsDatumId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;controls_data.id&#x60;.&lt;fk table&#x3D;&#39;controls_data&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;tests.id&#x60;.&lt;fk table&#x3D;&#39;tests&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**TestSectionId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;test_sections.id&#x60;.&lt;fk table&#x3D;&#39;test_sections&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ReferenceIssueId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;issues.id&#x60;.&lt;fk table&#x3D;&#39;issues&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**Flattened** | Pointer to **interface{}** |  | [optional] 
**ReferenceMeta** | Pointer to **interface{}** |  | [optional] 
**RemediationOwners** | Pointer to **interface{}** |  | [optional] 
**Totals** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**CreatedByUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 

## Methods

### NewIssuesArchivesPutPreviousValues

`func NewIssuesArchivesPutPreviousValues() *IssuesArchivesPutPreviousValues`

NewIssuesArchivesPutPreviousValues instantiates a new IssuesArchivesPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuesArchivesPutPreviousValuesWithDefaults

`func NewIssuesArchivesPutPreviousValuesWithDefaults() *IssuesArchivesPutPreviousValues`

NewIssuesArchivesPutPreviousValuesWithDefaults instantiates a new IssuesArchivesPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssuesArchivesPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssuesArchivesPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssuesArchivesPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IssuesArchivesPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *IssuesArchivesPutPreviousValues) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IssuesArchivesPutPreviousValues) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IssuesArchivesPutPreviousValues) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IssuesArchivesPutPreviousValues) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *IssuesArchivesPutPreviousValues) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IssuesArchivesPutPreviousValues) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IssuesArchivesPutPreviousValues) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IssuesArchivesPutPreviousValues) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIssueId

`func (o *IssuesArchivesPutPreviousValues) GetIssueId() int32`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *IssuesArchivesPutPreviousValues) GetIssueIdOk() (*int32, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *IssuesArchivesPutPreviousValues) SetIssueId(v int32)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *IssuesArchivesPutPreviousValues) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.

### GetArchiveId

`func (o *IssuesArchivesPutPreviousValues) GetArchiveId() int32`

GetArchiveId returns the ArchiveId field if non-nil, zero value otherwise.

### GetArchiveIdOk

`func (o *IssuesArchivesPutPreviousValues) GetArchiveIdOk() (*int32, bool)`

GetArchiveIdOk returns a tuple with the ArchiveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveId

`func (o *IssuesArchivesPutPreviousValues) SetArchiveId(v int32)`

SetArchiveId sets ArchiveId field to given value.

### HasArchiveId

`func (o *IssuesArchivesPutPreviousValues) HasArchiveId() bool`

HasArchiveId returns a boolean if a field has been set.

### GetControlsDataArchiveId

`func (o *IssuesArchivesPutPreviousValues) GetControlsDataArchiveId() int32`

GetControlsDataArchiveId returns the ControlsDataArchiveId field if non-nil, zero value otherwise.

### GetControlsDataArchiveIdOk

`func (o *IssuesArchivesPutPreviousValues) GetControlsDataArchiveIdOk() (*int32, bool)`

GetControlsDataArchiveIdOk returns a tuple with the ControlsDataArchiveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlsDataArchiveId

`func (o *IssuesArchivesPutPreviousValues) SetControlsDataArchiveId(v int32)`

SetControlsDataArchiveId sets ControlsDataArchiveId field to given value.

### HasControlsDataArchiveId

`func (o *IssuesArchivesPutPreviousValues) HasControlsDataArchiveId() bool`

HasControlsDataArchiveId returns a boolean if a field has been set.

### GetCreatorUserId

`func (o *IssuesArchivesPutPreviousValues) GetCreatorUserId() int32`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *IssuesArchivesPutPreviousValues) GetCreatorUserIdOk() (*int32, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *IssuesArchivesPutPreviousValues) SetCreatorUserId(v int32)`

SetCreatorUserId sets CreatorUserId field to given value.

### HasCreatorUserId

`func (o *IssuesArchivesPutPreviousValues) HasCreatorUserId() bool`

HasCreatorUserId returns a boolean if a field has been set.

### GetControlsDatumId

`func (o *IssuesArchivesPutPreviousValues) GetControlsDatumId() int32`

GetControlsDatumId returns the ControlsDatumId field if non-nil, zero value otherwise.

### GetControlsDatumIdOk

`func (o *IssuesArchivesPutPreviousValues) GetControlsDatumIdOk() (*int32, bool)`

GetControlsDatumIdOk returns a tuple with the ControlsDatumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlsDatumId

`func (o *IssuesArchivesPutPreviousValues) SetControlsDatumId(v int32)`

SetControlsDatumId sets ControlsDatumId field to given value.

### HasControlsDatumId

`func (o *IssuesArchivesPutPreviousValues) HasControlsDatumId() bool`

HasControlsDatumId returns a boolean if a field has been set.

### GetTestId

`func (o *IssuesArchivesPutPreviousValues) GetTestId() int32`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *IssuesArchivesPutPreviousValues) GetTestIdOk() (*int32, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *IssuesArchivesPutPreviousValues) SetTestId(v int32)`

SetTestId sets TestId field to given value.

### HasTestId

`func (o *IssuesArchivesPutPreviousValues) HasTestId() bool`

HasTestId returns a boolean if a field has been set.

### GetTestSectionId

`func (o *IssuesArchivesPutPreviousValues) GetTestSectionId() int32`

GetTestSectionId returns the TestSectionId field if non-nil, zero value otherwise.

### GetTestSectionIdOk

`func (o *IssuesArchivesPutPreviousValues) GetTestSectionIdOk() (*int32, bool)`

GetTestSectionIdOk returns a tuple with the TestSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSectionId

`func (o *IssuesArchivesPutPreviousValues) SetTestSectionId(v int32)`

SetTestSectionId sets TestSectionId field to given value.

### HasTestSectionId

`func (o *IssuesArchivesPutPreviousValues) HasTestSectionId() bool`

HasTestSectionId returns a boolean if a field has been set.

### GetReferenceIssueId

`func (o *IssuesArchivesPutPreviousValues) GetReferenceIssueId() int32`

GetReferenceIssueId returns the ReferenceIssueId field if non-nil, zero value otherwise.

### GetReferenceIssueIdOk

`func (o *IssuesArchivesPutPreviousValues) GetReferenceIssueIdOk() (*int32, bool)`

GetReferenceIssueIdOk returns a tuple with the ReferenceIssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceIssueId

`func (o *IssuesArchivesPutPreviousValues) SetReferenceIssueId(v int32)`

SetReferenceIssueId sets ReferenceIssueId field to given value.

### HasReferenceIssueId

`func (o *IssuesArchivesPutPreviousValues) HasReferenceIssueId() bool`

HasReferenceIssueId returns a boolean if a field has been set.

### GetFlattened

`func (o *IssuesArchivesPutPreviousValues) GetFlattened() interface{}`

GetFlattened returns the Flattened field if non-nil, zero value otherwise.

### GetFlattenedOk

`func (o *IssuesArchivesPutPreviousValues) GetFlattenedOk() (*interface{}, bool)`

GetFlattenedOk returns a tuple with the Flattened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlattened

`func (o *IssuesArchivesPutPreviousValues) SetFlattened(v interface{})`

SetFlattened sets Flattened field to given value.

### HasFlattened

`func (o *IssuesArchivesPutPreviousValues) HasFlattened() bool`

HasFlattened returns a boolean if a field has been set.

### SetFlattenedNil

`func (o *IssuesArchivesPutPreviousValues) SetFlattenedNil(b bool)`

 SetFlattenedNil sets the value for Flattened to be an explicit nil

### UnsetFlattened
`func (o *IssuesArchivesPutPreviousValues) UnsetFlattened()`

UnsetFlattened ensures that no value is present for Flattened, not even an explicit nil
### GetReferenceMeta

`func (o *IssuesArchivesPutPreviousValues) GetReferenceMeta() interface{}`

GetReferenceMeta returns the ReferenceMeta field if non-nil, zero value otherwise.

### GetReferenceMetaOk

`func (o *IssuesArchivesPutPreviousValues) GetReferenceMetaOk() (*interface{}, bool)`

GetReferenceMetaOk returns a tuple with the ReferenceMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMeta

`func (o *IssuesArchivesPutPreviousValues) SetReferenceMeta(v interface{})`

SetReferenceMeta sets ReferenceMeta field to given value.

### HasReferenceMeta

`func (o *IssuesArchivesPutPreviousValues) HasReferenceMeta() bool`

HasReferenceMeta returns a boolean if a field has been set.

### SetReferenceMetaNil

`func (o *IssuesArchivesPutPreviousValues) SetReferenceMetaNil(b bool)`

 SetReferenceMetaNil sets the value for ReferenceMeta to be an explicit nil

### UnsetReferenceMeta
`func (o *IssuesArchivesPutPreviousValues) UnsetReferenceMeta()`

UnsetReferenceMeta ensures that no value is present for ReferenceMeta, not even an explicit nil
### GetRemediationOwners

`func (o *IssuesArchivesPutPreviousValues) GetRemediationOwners() interface{}`

GetRemediationOwners returns the RemediationOwners field if non-nil, zero value otherwise.

### GetRemediationOwnersOk

`func (o *IssuesArchivesPutPreviousValues) GetRemediationOwnersOk() (*interface{}, bool)`

GetRemediationOwnersOk returns a tuple with the RemediationOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationOwners

`func (o *IssuesArchivesPutPreviousValues) SetRemediationOwners(v interface{})`

SetRemediationOwners sets RemediationOwners field to given value.

### HasRemediationOwners

`func (o *IssuesArchivesPutPreviousValues) HasRemediationOwners() bool`

HasRemediationOwners returns a boolean if a field has been set.

### SetRemediationOwnersNil

`func (o *IssuesArchivesPutPreviousValues) SetRemediationOwnersNil(b bool)`

 SetRemediationOwnersNil sets the value for RemediationOwners to be an explicit nil

### UnsetRemediationOwners
`func (o *IssuesArchivesPutPreviousValues) UnsetRemediationOwners()`

UnsetRemediationOwners ensures that no value is present for RemediationOwners, not even an explicit nil
### GetTotals

`func (o *IssuesArchivesPutPreviousValues) GetTotals() interface{}`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *IssuesArchivesPutPreviousValues) GetTotalsOk() (*interface{}, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *IssuesArchivesPutPreviousValues) SetTotals(v interface{})`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *IssuesArchivesPutPreviousValues) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### SetTotalsNil

`func (o *IssuesArchivesPutPreviousValues) SetTotalsNil(b bool)`

 SetTotalsNil sets the value for Totals to be an explicit nil

### UnsetTotals
`func (o *IssuesArchivesPutPreviousValues) UnsetTotals()`

UnsetTotals ensures that no value is present for Totals, not even an explicit nil
### GetCreatedAt

`func (o *IssuesArchivesPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssuesArchivesPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssuesArchivesPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IssuesArchivesPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IssuesArchivesPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IssuesArchivesPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IssuesArchivesPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IssuesArchivesPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *IssuesArchivesPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *IssuesArchivesPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *IssuesArchivesPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *IssuesArchivesPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *IssuesArchivesPutPreviousValues) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *IssuesArchivesPutPreviousValues) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *IssuesArchivesPutPreviousValues) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *IssuesArchivesPutPreviousValues) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


