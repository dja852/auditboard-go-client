# \ReportsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportsGenerateAuditFormsPost**](ReportsAPI.md#ReportsGenerateAuditFormsPost) | **Post** /reports/generate/audit_forms | 
[**ReportsGenerateAuditableEntityPost**](ReportsAPI.md#ReportsGenerateAuditableEntityPost) | **Post** /reports/generate/auditable_entity | 
[**ReportsGenerateOpsauditPost**](ReportsAPI.md#ReportsGenerateOpsauditPost) | **Post** /reports/generate/opsaudit | 
[**ReportsGeneratePost**](ReportsAPI.md#ReportsGeneratePost) | **Post** /reports/generate | 
[**ReportsGenerateTimesheetsPost**](ReportsAPI.md#ReportsGenerateTimesheetsPost) | **Post** /reports/generate/timesheets | 
[**ReportsGenerateWorkstepPost**](ReportsAPI.md#ReportsGenerateWorkstepPost) | **Post** /reports/generate/workstep | 
[**ReportsGenerateWorkstreamPost**](ReportsAPI.md#ReportsGenerateWorkstreamPost) | **Post** /reports/generate/workstream | 
[**ReportsGet**](ReportsAPI.md#ReportsGet) | **Get** /reports | 
[**ReportsPost**](ReportsAPI.md#ReportsPost) | **Post** /reports | 
[**ReportsReportIdDelete**](ReportsAPI.md#ReportsReportIdDelete) | **Delete** /reports/{report_id} | 
[**ReportsReportIdGet**](ReportsAPI.md#ReportsReportIdGet) | **Get** /reports/{report_id} | 
[**ReportsReportIdPut**](ReportsAPI.md#ReportsReportIdPut) | **Put** /reports/{report_id} | 



## ReportsGenerateAuditFormsPost

> ReportsGenerateWorkstreamPost200Response ReportsGenerateAuditFormsPost(ctx).ReportsGenerateAuditFormsPostRequest(reportsGenerateAuditFormsPostRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dja852/auditboard-go-client"
)

func main() {
	reportsGenerateAuditFormsPostRequest := *openapiclient.NewReportsGenerateAuditFormsPostRequest() // ReportsGenerateAuditFormsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsGenerateAuditFormsPost(context.Background()).ReportsGenerateAuditFormsPostRequest(reportsGenerateAuditFormsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsGenerateAuditFormsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsGenerateAuditFormsPost`: ReportsGenerateWorkstreamPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsGenerateAuditFormsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsGenerateAuditFormsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportsGenerateAuditFormsPostRequest** | [**ReportsGenerateAuditFormsPostRequest**](ReportsGenerateAuditFormsPostRequest.md) |  | 

### Return type

[**ReportsGenerateWorkstreamPost200Response**](ReportsGenerateWorkstreamPost200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsGenerateAuditableEntityPost

> ReportsGenerateWorkstreamPost200Response ReportsGenerateAuditableEntityPost(ctx).ReportsGenerateAuditableEntityPostRequest(reportsGenerateAuditableEntityPostRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dja852/auditboard-go-client"
)

func main() {
	reportsGenerateAuditableEntityPostRequest := *openapiclient.NewReportsGenerateAuditableEntityPostRequest() // ReportsGenerateAuditableEntityPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsGenerateAuditableEntityPost(context.Background()).ReportsGenerateAuditableEntityPostRequest(reportsGenerateAuditableEntityPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsGenerateAuditableEntityPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsGenerateAuditableEntityPost`: ReportsGenerateWorkstreamPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsGenerateAuditableEntityPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsGenerateAuditableEntityPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportsGenerateAuditableEntityPostRequest** | [**ReportsGenerateAuditableEntityPostRequest**](ReportsGenerateAuditableEntityPostRequest.md) |  | 

### Return type

[**ReportsGenerateWorkstreamPost200Response**](ReportsGenerateWorkstreamPost200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsGenerateOpsauditPost

> ReportsGenerateWorkstreamPost200Response ReportsGenerateOpsauditPost(ctx).ReportsGenerateOpsauditPostRequest(reportsGenerateOpsauditPostRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dja852/auditboard-go-client"
)

func main() {
	reportsGenerateOpsauditPostRequest := *openapiclient.NewReportsGenerateOpsauditPostRequest() // ReportsGenerateOpsauditPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsGenerateOpsauditPost(context.Background()).ReportsGenerateOpsauditPostRequest(reportsGenerateOpsauditPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsGenerateOpsauditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsGenerateOpsauditPost`: ReportsGenerateWorkstreamPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsGenerateOpsauditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsGenerateOpsauditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportsGenerateOpsauditPostRequest** | [**ReportsGenerateOpsauditPostRequest**](ReportsGenerateOpsauditPostRequest.md) |  | 

### Return type

[**ReportsGenerateWorkstreamPost200Response**](ReportsGenerateWorkstreamPost200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsGeneratePost

> ReportsGenerateWorkstreamPost200Response ReportsGeneratePost(ctx).ReportsGeneratePostRequest(reportsGeneratePostRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dja852/auditboard-go-client"
)

func main() {
	reportsGeneratePostRequest := *openapiclient.NewReportsGeneratePostRequest() // ReportsGeneratePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsGeneratePost(context.Background()).ReportsGeneratePostRequest(reportsGeneratePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsGeneratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsGeneratePost`: ReportsGenerateWorkstreamPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsGeneratePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsGeneratePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportsGeneratePostRequest** | [**ReportsGeneratePostRequest**](ReportsGeneratePostRequest.md) |  | 

### Return type

[**ReportsGenerateWorkstreamPost200Response**](ReportsGenerateWorkstreamPost200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsGenerateTimesheetsPost

> ReportsGenerateTimesheetsPost200Response ReportsGenerateTimesheetsPost(ctx).ReportsGenerateTimesheetsPostRequest(reportsGenerateTimesheetsPostRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dja852/auditboard-go-client"
)

func main() {
	reportsGenerateTimesheetsPostRequest := *openapiclient.NewReportsGenerateTimesheetsPostRequest() // ReportsGenerateTimesheetsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsGenerateTimesheetsPost(context.Background()).ReportsGenerateTimesheetsPostRequest(reportsGenerateTimesheetsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsGenerateTimesheetsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsGenerateTimesheetsPost`: ReportsGenerateTimesheetsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsGenerateTimesheetsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsGenerateTimesheetsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportsGenerateTimesheetsPostRequest** | [**ReportsGenerateTimesheetsPostRequest**](ReportsGenerateTimesheetsPostRequest.md) |  | 

### Return type

[**ReportsGenerateTimesheetsPost200Response**](ReportsGenerateTimesheetsPost200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsGenerateWorkstepPost

> ReportsGenerateWorkstepPost200Response ReportsGenerateWorkstepPost(ctx).ReportsGenerateWorkstepPostRequest(reportsGenerateWorkstepPostRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dja852/auditboard-go-client"
)

func main() {
	reportsGenerateWorkstepPostRequest := *openapiclient.NewReportsGenerateWorkstepPostRequest() // ReportsGenerateWorkstepPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsGenerateWorkstepPost(context.Background()).ReportsGenerateWorkstepPostRequest(reportsGenerateWorkstepPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsGenerateWorkstepPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsGenerateWorkstepPost`: ReportsGenerateWorkstepPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsGenerateWorkstepPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsGenerateWorkstepPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportsGenerateWorkstepPostRequest** | [**ReportsGenerateWorkstepPostRequest**](ReportsGenerateWorkstepPostRequest.md) |  | 

### Return type

[**ReportsGenerateWorkstepPost200Response**](ReportsGenerateWorkstepPost200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsGenerateWorkstreamPost

> ReportsGenerateWorkstreamPost200Response ReportsGenerateWorkstreamPost(ctx).ReportsGenerateWorkstreamPostRequest(reportsGenerateWorkstreamPostRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dja852/auditboard-go-client"
)

func main() {
	reportsGenerateWorkstreamPostRequest := *openapiclient.NewReportsGenerateWorkstreamPostRequest() // ReportsGenerateWorkstreamPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsGenerateWorkstreamPost(context.Background()).ReportsGenerateWorkstreamPostRequest(reportsGenerateWorkstreamPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsGenerateWorkstreamPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsGenerateWorkstreamPost`: ReportsGenerateWorkstreamPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsGenerateWorkstreamPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsGenerateWorkstreamPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportsGenerateWorkstreamPostRequest** | [**ReportsGenerateWorkstreamPostRequest**](ReportsGenerateWorkstreamPostRequest.md) |  | 

### Return type

[**ReportsGenerateWorkstreamPost200Response**](ReportsGenerateWorkstreamPost200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsGet

> ReportsGet200Response ReportsGet(ctx).Include(include).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dja852/auditboard-go-client"
)

func main() {
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsGet`: ReportsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**ReportsGet200Response**](ReportsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsPost

> ReportsGet200Response ReportsPost(ctx).ReportsPostRequest(reportsPostRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dja852/auditboard-go-client"
)

func main() {
	reportsPostRequest := *openapiclient.NewReportsPostRequest() // ReportsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsPost(context.Background()).ReportsPostRequest(reportsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsPost`: ReportsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportsPostRequest** | [**ReportsPostRequest**](ReportsPostRequest.md) |  | 

### Return type

[**ReportsGet200Response**](ReportsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsReportIdDelete

> Reports ReportsReportIdDelete(ctx, reportId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dja852/auditboard-go-client"
)

func main() {
	reportId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsReportIdDelete(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsReportIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsReportIdDelete`: Reports
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsReportIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsReportIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Reports**](Reports.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsReportIdGet

> Reports ReportsReportIdGet(ctx, reportId).Include(include).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dja852/auditboard-go-client"
)

func main() {
	reportId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsReportIdGet(context.Background(), reportId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsReportIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsReportIdGet`: Reports
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsReportIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsReportIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**Reports**](Reports.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsReportIdPut

> Reports ReportsReportIdPut(ctx, reportId).ReportsPut(reportsPut).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dja852/auditboard-go-client"
)

func main() {
	reportId := int64(789) // int64 | Model id
	reportsPut := *openapiclient.NewReportsPut() // ReportsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsReportIdPut(context.Background(), reportId).ReportsPut(reportsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsReportIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsReportIdPut`: Reports
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsReportIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsReportIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportsPut** | [**ReportsPut**](ReportsPut.md) |  | 

### Return type

[**Reports**](Reports.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

