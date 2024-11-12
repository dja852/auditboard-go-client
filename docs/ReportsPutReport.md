# ReportsPutReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**QueryFields** | **string** |  | 
**QueryFilters** | **string** |  | 
**ShareToken** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | [default to "Control"]
**FieldsGrouping** | Pointer to **string** |  | [optional] 
**CanonicalFields** | Pointer to **interface{}** |  | [optional] 
**ShareLinkUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ReportOptions** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewReportsPutReport

`func NewReportsPutReport(name string, queryFields string, queryFilters string, type_ string, ) *ReportsPutReport`

NewReportsPutReport instantiates a new ReportsPutReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportsPutReportWithDefaults

`func NewReportsPutReportWithDefaults() *ReportsPutReport`

NewReportsPutReportWithDefaults instantiates a new ReportsPutReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReportsPutReport) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportsPutReport) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportsPutReport) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReportsPutReport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ReportsPutReport) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReportsPutReport) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReportsPutReport) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ReportsPutReport) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ReportsPutReport) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ReportsPutReport) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ReportsPutReport) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ReportsPutReport) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ReportsPutReport) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ReportsPutReport) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ReportsPutReport) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ReportsPutReport) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *ReportsPutReport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportsPutReport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportsPutReport) SetName(v string)`

SetName sets Name field to given value.


### GetQueryFields

`func (o *ReportsPutReport) GetQueryFields() string`

GetQueryFields returns the QueryFields field if non-nil, zero value otherwise.

### GetQueryFieldsOk

`func (o *ReportsPutReport) GetQueryFieldsOk() (*string, bool)`

GetQueryFieldsOk returns a tuple with the QueryFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryFields

`func (o *ReportsPutReport) SetQueryFields(v string)`

SetQueryFields sets QueryFields field to given value.


### GetQueryFilters

`func (o *ReportsPutReport) GetQueryFilters() string`

GetQueryFilters returns the QueryFilters field if non-nil, zero value otherwise.

### GetQueryFiltersOk

`func (o *ReportsPutReport) GetQueryFiltersOk() (*string, bool)`

GetQueryFiltersOk returns a tuple with the QueryFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryFilters

`func (o *ReportsPutReport) SetQueryFilters(v string)`

SetQueryFilters sets QueryFilters field to given value.


### GetShareToken

`func (o *ReportsPutReport) GetShareToken() string`

GetShareToken returns the ShareToken field if non-nil, zero value otherwise.

### GetShareTokenOk

`func (o *ReportsPutReport) GetShareTokenOk() (*string, bool)`

GetShareTokenOk returns a tuple with the ShareToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareToken

`func (o *ReportsPutReport) SetShareToken(v string)`

SetShareToken sets ShareToken field to given value.

### HasShareToken

`func (o *ReportsPutReport) HasShareToken() bool`

HasShareToken returns a boolean if a field has been set.

### GetType

`func (o *ReportsPutReport) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReportsPutReport) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReportsPutReport) SetType(v string)`

SetType sets Type field to given value.


### GetFieldsGrouping

`func (o *ReportsPutReport) GetFieldsGrouping() string`

GetFieldsGrouping returns the FieldsGrouping field if non-nil, zero value otherwise.

### GetFieldsGroupingOk

`func (o *ReportsPutReport) GetFieldsGroupingOk() (*string, bool)`

GetFieldsGroupingOk returns a tuple with the FieldsGrouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsGrouping

`func (o *ReportsPutReport) SetFieldsGrouping(v string)`

SetFieldsGrouping sets FieldsGrouping field to given value.

### HasFieldsGrouping

`func (o *ReportsPutReport) HasFieldsGrouping() bool`

HasFieldsGrouping returns a boolean if a field has been set.

### GetCanonicalFields

`func (o *ReportsPutReport) GetCanonicalFields() interface{}`

GetCanonicalFields returns the CanonicalFields field if non-nil, zero value otherwise.

### GetCanonicalFieldsOk

`func (o *ReportsPutReport) GetCanonicalFieldsOk() (*interface{}, bool)`

GetCanonicalFieldsOk returns a tuple with the CanonicalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalFields

`func (o *ReportsPutReport) SetCanonicalFields(v interface{})`

SetCanonicalFields sets CanonicalFields field to given value.

### HasCanonicalFields

`func (o *ReportsPutReport) HasCanonicalFields() bool`

HasCanonicalFields returns a boolean if a field has been set.

### SetCanonicalFieldsNil

`func (o *ReportsPutReport) SetCanonicalFieldsNil(b bool)`

 SetCanonicalFieldsNil sets the value for CanonicalFields to be an explicit nil

### UnsetCanonicalFields
`func (o *ReportsPutReport) UnsetCanonicalFields()`

UnsetCanonicalFields ensures that no value is present for CanonicalFields, not even an explicit nil
### GetShareLinkUserId

`func (o *ReportsPutReport) GetShareLinkUserId() int32`

GetShareLinkUserId returns the ShareLinkUserId field if non-nil, zero value otherwise.

### GetShareLinkUserIdOk

`func (o *ReportsPutReport) GetShareLinkUserIdOk() (*int32, bool)`

GetShareLinkUserIdOk returns a tuple with the ShareLinkUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareLinkUserId

`func (o *ReportsPutReport) SetShareLinkUserId(v int32)`

SetShareLinkUserId sets ShareLinkUserId field to given value.

### HasShareLinkUserId

`func (o *ReportsPutReport) HasShareLinkUserId() bool`

HasShareLinkUserId returns a boolean if a field has been set.

### GetReportOptions

`func (o *ReportsPutReport) GetReportOptions() interface{}`

GetReportOptions returns the ReportOptions field if non-nil, zero value otherwise.

### GetReportOptionsOk

`func (o *ReportsPutReport) GetReportOptionsOk() (*interface{}, bool)`

GetReportOptionsOk returns a tuple with the ReportOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportOptions

`func (o *ReportsPutReport) SetReportOptions(v interface{})`

SetReportOptions sets ReportOptions field to given value.

### HasReportOptions

`func (o *ReportsPutReport) HasReportOptions() bool`

HasReportOptions returns a boolean if a field has been set.

### SetReportOptionsNil

`func (o *ReportsPutReport) SetReportOptionsNil(b bool)`

 SetReportOptionsNil sets the value for ReportOptions to be an explicit nil

### UnsetReportOptions
`func (o *ReportsPutReport) UnsetReportOptions()`

UnsetReportOptions ensures that no value is present for ReportOptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


