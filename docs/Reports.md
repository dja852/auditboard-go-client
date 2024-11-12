# Reports

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

### NewReports

`func NewReports(name string, queryFields string, queryFilters string, type_ string, ) *Reports`

NewReports instantiates a new Reports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportsWithDefaults

`func NewReportsWithDefaults() *Reports`

NewReportsWithDefaults instantiates a new Reports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Reports) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Reports) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Reports) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Reports) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Reports) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Reports) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Reports) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Reports) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Reports) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Reports) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Reports) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Reports) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Reports) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Reports) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Reports) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Reports) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *Reports) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Reports) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Reports) SetName(v string)`

SetName sets Name field to given value.


### GetQueryFields

`func (o *Reports) GetQueryFields() string`

GetQueryFields returns the QueryFields field if non-nil, zero value otherwise.

### GetQueryFieldsOk

`func (o *Reports) GetQueryFieldsOk() (*string, bool)`

GetQueryFieldsOk returns a tuple with the QueryFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryFields

`func (o *Reports) SetQueryFields(v string)`

SetQueryFields sets QueryFields field to given value.


### GetQueryFilters

`func (o *Reports) GetQueryFilters() string`

GetQueryFilters returns the QueryFilters field if non-nil, zero value otherwise.

### GetQueryFiltersOk

`func (o *Reports) GetQueryFiltersOk() (*string, bool)`

GetQueryFiltersOk returns a tuple with the QueryFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryFilters

`func (o *Reports) SetQueryFilters(v string)`

SetQueryFilters sets QueryFilters field to given value.


### GetShareToken

`func (o *Reports) GetShareToken() string`

GetShareToken returns the ShareToken field if non-nil, zero value otherwise.

### GetShareTokenOk

`func (o *Reports) GetShareTokenOk() (*string, bool)`

GetShareTokenOk returns a tuple with the ShareToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareToken

`func (o *Reports) SetShareToken(v string)`

SetShareToken sets ShareToken field to given value.

### HasShareToken

`func (o *Reports) HasShareToken() bool`

HasShareToken returns a boolean if a field has been set.

### GetType

`func (o *Reports) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Reports) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Reports) SetType(v string)`

SetType sets Type field to given value.


### GetFieldsGrouping

`func (o *Reports) GetFieldsGrouping() string`

GetFieldsGrouping returns the FieldsGrouping field if non-nil, zero value otherwise.

### GetFieldsGroupingOk

`func (o *Reports) GetFieldsGroupingOk() (*string, bool)`

GetFieldsGroupingOk returns a tuple with the FieldsGrouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsGrouping

`func (o *Reports) SetFieldsGrouping(v string)`

SetFieldsGrouping sets FieldsGrouping field to given value.

### HasFieldsGrouping

`func (o *Reports) HasFieldsGrouping() bool`

HasFieldsGrouping returns a boolean if a field has been set.

### GetCanonicalFields

`func (o *Reports) GetCanonicalFields() interface{}`

GetCanonicalFields returns the CanonicalFields field if non-nil, zero value otherwise.

### GetCanonicalFieldsOk

`func (o *Reports) GetCanonicalFieldsOk() (*interface{}, bool)`

GetCanonicalFieldsOk returns a tuple with the CanonicalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalFields

`func (o *Reports) SetCanonicalFields(v interface{})`

SetCanonicalFields sets CanonicalFields field to given value.

### HasCanonicalFields

`func (o *Reports) HasCanonicalFields() bool`

HasCanonicalFields returns a boolean if a field has been set.

### SetCanonicalFieldsNil

`func (o *Reports) SetCanonicalFieldsNil(b bool)`

 SetCanonicalFieldsNil sets the value for CanonicalFields to be an explicit nil

### UnsetCanonicalFields
`func (o *Reports) UnsetCanonicalFields()`

UnsetCanonicalFields ensures that no value is present for CanonicalFields, not even an explicit nil
### GetShareLinkUserId

`func (o *Reports) GetShareLinkUserId() int32`

GetShareLinkUserId returns the ShareLinkUserId field if non-nil, zero value otherwise.

### GetShareLinkUserIdOk

`func (o *Reports) GetShareLinkUserIdOk() (*int32, bool)`

GetShareLinkUserIdOk returns a tuple with the ShareLinkUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareLinkUserId

`func (o *Reports) SetShareLinkUserId(v int32)`

SetShareLinkUserId sets ShareLinkUserId field to given value.

### HasShareLinkUserId

`func (o *Reports) HasShareLinkUserId() bool`

HasShareLinkUserId returns a boolean if a field has been set.

### GetReportOptions

`func (o *Reports) GetReportOptions() interface{}`

GetReportOptions returns the ReportOptions field if non-nil, zero value otherwise.

### GetReportOptionsOk

`func (o *Reports) GetReportOptionsOk() (*interface{}, bool)`

GetReportOptionsOk returns a tuple with the ReportOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportOptions

`func (o *Reports) SetReportOptions(v interface{})`

SetReportOptions sets ReportOptions field to given value.

### HasReportOptions

`func (o *Reports) HasReportOptions() bool`

HasReportOptions returns a boolean if a field has been set.

### SetReportOptionsNil

`func (o *Reports) SetReportOptionsNil(b bool)`

 SetReportOptionsNil sets the value for ReportOptions to be an explicit nil

### UnsetReportOptions
`func (o *Reports) UnsetReportOptions()`

UnsetReportOptions ensures that no value is present for ReportOptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


