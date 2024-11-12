# ReportsPutPreviousValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Note: This is a Primary Key.&lt;pk/&gt; | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**DeletedAt** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**QueryFields** | Pointer to **string** |  | [optional] 
**QueryFilters** | Pointer to **string** |  | [optional] 
**ShareToken** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "Control"]
**FieldsGrouping** | Pointer to **string** |  | [optional] 
**CanonicalFields** | Pointer to **interface{}** |  | [optional] 
**ShareLinkUserId** | Pointer to **int32** | Note: This is a Foreign Key to &#x60;users.id&#x60;.&lt;fk table&#x3D;&#39;users&#39; column&#x3D;&#39;id&#39;/&gt; | [optional] 
**ReportOptions** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewReportsPutPreviousValues

`func NewReportsPutPreviousValues() *ReportsPutPreviousValues`

NewReportsPutPreviousValues instantiates a new ReportsPutPreviousValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportsPutPreviousValuesWithDefaults

`func NewReportsPutPreviousValuesWithDefaults() *ReportsPutPreviousValues`

NewReportsPutPreviousValuesWithDefaults instantiates a new ReportsPutPreviousValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReportsPutPreviousValues) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportsPutPreviousValues) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportsPutPreviousValues) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReportsPutPreviousValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ReportsPutPreviousValues) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReportsPutPreviousValues) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReportsPutPreviousValues) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ReportsPutPreviousValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ReportsPutPreviousValues) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ReportsPutPreviousValues) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ReportsPutPreviousValues) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ReportsPutPreviousValues) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ReportsPutPreviousValues) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ReportsPutPreviousValues) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ReportsPutPreviousValues) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ReportsPutPreviousValues) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *ReportsPutPreviousValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportsPutPreviousValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportsPutPreviousValues) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReportsPutPreviousValues) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQueryFields

`func (o *ReportsPutPreviousValues) GetQueryFields() string`

GetQueryFields returns the QueryFields field if non-nil, zero value otherwise.

### GetQueryFieldsOk

`func (o *ReportsPutPreviousValues) GetQueryFieldsOk() (*string, bool)`

GetQueryFieldsOk returns a tuple with the QueryFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryFields

`func (o *ReportsPutPreviousValues) SetQueryFields(v string)`

SetQueryFields sets QueryFields field to given value.

### HasQueryFields

`func (o *ReportsPutPreviousValues) HasQueryFields() bool`

HasQueryFields returns a boolean if a field has been set.

### GetQueryFilters

`func (o *ReportsPutPreviousValues) GetQueryFilters() string`

GetQueryFilters returns the QueryFilters field if non-nil, zero value otherwise.

### GetQueryFiltersOk

`func (o *ReportsPutPreviousValues) GetQueryFiltersOk() (*string, bool)`

GetQueryFiltersOk returns a tuple with the QueryFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryFilters

`func (o *ReportsPutPreviousValues) SetQueryFilters(v string)`

SetQueryFilters sets QueryFilters field to given value.

### HasQueryFilters

`func (o *ReportsPutPreviousValues) HasQueryFilters() bool`

HasQueryFilters returns a boolean if a field has been set.

### GetShareToken

`func (o *ReportsPutPreviousValues) GetShareToken() string`

GetShareToken returns the ShareToken field if non-nil, zero value otherwise.

### GetShareTokenOk

`func (o *ReportsPutPreviousValues) GetShareTokenOk() (*string, bool)`

GetShareTokenOk returns a tuple with the ShareToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareToken

`func (o *ReportsPutPreviousValues) SetShareToken(v string)`

SetShareToken sets ShareToken field to given value.

### HasShareToken

`func (o *ReportsPutPreviousValues) HasShareToken() bool`

HasShareToken returns a boolean if a field has been set.

### GetType

`func (o *ReportsPutPreviousValues) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReportsPutPreviousValues) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReportsPutPreviousValues) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReportsPutPreviousValues) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFieldsGrouping

`func (o *ReportsPutPreviousValues) GetFieldsGrouping() string`

GetFieldsGrouping returns the FieldsGrouping field if non-nil, zero value otherwise.

### GetFieldsGroupingOk

`func (o *ReportsPutPreviousValues) GetFieldsGroupingOk() (*string, bool)`

GetFieldsGroupingOk returns a tuple with the FieldsGrouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsGrouping

`func (o *ReportsPutPreviousValues) SetFieldsGrouping(v string)`

SetFieldsGrouping sets FieldsGrouping field to given value.

### HasFieldsGrouping

`func (o *ReportsPutPreviousValues) HasFieldsGrouping() bool`

HasFieldsGrouping returns a boolean if a field has been set.

### GetCanonicalFields

`func (o *ReportsPutPreviousValues) GetCanonicalFields() interface{}`

GetCanonicalFields returns the CanonicalFields field if non-nil, zero value otherwise.

### GetCanonicalFieldsOk

`func (o *ReportsPutPreviousValues) GetCanonicalFieldsOk() (*interface{}, bool)`

GetCanonicalFieldsOk returns a tuple with the CanonicalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalFields

`func (o *ReportsPutPreviousValues) SetCanonicalFields(v interface{})`

SetCanonicalFields sets CanonicalFields field to given value.

### HasCanonicalFields

`func (o *ReportsPutPreviousValues) HasCanonicalFields() bool`

HasCanonicalFields returns a boolean if a field has been set.

### SetCanonicalFieldsNil

`func (o *ReportsPutPreviousValues) SetCanonicalFieldsNil(b bool)`

 SetCanonicalFieldsNil sets the value for CanonicalFields to be an explicit nil

### UnsetCanonicalFields
`func (o *ReportsPutPreviousValues) UnsetCanonicalFields()`

UnsetCanonicalFields ensures that no value is present for CanonicalFields, not even an explicit nil
### GetShareLinkUserId

`func (o *ReportsPutPreviousValues) GetShareLinkUserId() int32`

GetShareLinkUserId returns the ShareLinkUserId field if non-nil, zero value otherwise.

### GetShareLinkUserIdOk

`func (o *ReportsPutPreviousValues) GetShareLinkUserIdOk() (*int32, bool)`

GetShareLinkUserIdOk returns a tuple with the ShareLinkUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareLinkUserId

`func (o *ReportsPutPreviousValues) SetShareLinkUserId(v int32)`

SetShareLinkUserId sets ShareLinkUserId field to given value.

### HasShareLinkUserId

`func (o *ReportsPutPreviousValues) HasShareLinkUserId() bool`

HasShareLinkUserId returns a boolean if a field has been set.

### GetReportOptions

`func (o *ReportsPutPreviousValues) GetReportOptions() interface{}`

GetReportOptions returns the ReportOptions field if non-nil, zero value otherwise.

### GetReportOptionsOk

`func (o *ReportsPutPreviousValues) GetReportOptionsOk() (*interface{}, bool)`

GetReportOptionsOk returns a tuple with the ReportOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportOptions

`func (o *ReportsPutPreviousValues) SetReportOptions(v interface{})`

SetReportOptions sets ReportOptions field to given value.

### HasReportOptions

`func (o *ReportsPutPreviousValues) HasReportOptions() bool`

HasReportOptions returns a boolean if a field has been set.

### SetReportOptionsNil

`func (o *ReportsPutPreviousValues) SetReportOptionsNil(b bool)`

 SetReportOptionsNil sets the value for ReportOptions to be an explicit nil

### UnsetReportOptions
`func (o *ReportsPutPreviousValues) UnsetReportOptions()`

UnsetReportOptions ensures that no value is present for ReportOptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


