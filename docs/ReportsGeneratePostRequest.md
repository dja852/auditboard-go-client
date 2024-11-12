# ReportsGeneratePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportType** | Pointer to **string** |  | [optional] 
**QueryData** | Pointer to [**ReportsGeneratePostRequestQueryData**](ReportsGeneratePostRequestQueryData.md) |  | [optional] 

## Methods

### NewReportsGeneratePostRequest

`func NewReportsGeneratePostRequest() *ReportsGeneratePostRequest`

NewReportsGeneratePostRequest instantiates a new ReportsGeneratePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportsGeneratePostRequestWithDefaults

`func NewReportsGeneratePostRequestWithDefaults() *ReportsGeneratePostRequest`

NewReportsGeneratePostRequestWithDefaults instantiates a new ReportsGeneratePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportType

`func (o *ReportsGeneratePostRequest) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *ReportsGeneratePostRequest) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *ReportsGeneratePostRequest) SetReportType(v string)`

SetReportType sets ReportType field to given value.

### HasReportType

`func (o *ReportsGeneratePostRequest) HasReportType() bool`

HasReportType returns a boolean if a field has been set.

### GetQueryData

`func (o *ReportsGeneratePostRequest) GetQueryData() ReportsGeneratePostRequestQueryData`

GetQueryData returns the QueryData field if non-nil, zero value otherwise.

### GetQueryDataOk

`func (o *ReportsGeneratePostRequest) GetQueryDataOk() (*ReportsGeneratePostRequestQueryData, bool)`

GetQueryDataOk returns a tuple with the QueryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryData

`func (o *ReportsGeneratePostRequest) SetQueryData(v ReportsGeneratePostRequestQueryData)`

SetQueryData sets QueryData field to given value.

### HasQueryData

`func (o *ReportsGeneratePostRequest) HasQueryData() bool`

HasQueryData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


