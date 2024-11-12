# \IssuesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IssuesGet**](IssuesAPI.md#IssuesGet) | **Get** /issues | 
[**IssuesIssueIdClosePost**](IssuesAPI.md#IssuesIssueIdClosePost) | **Post** /issues/{issue_id}/close | 
[**IssuesIssueIdDelete**](IssuesAPI.md#IssuesIssueIdDelete) | **Delete** /issues/{issue_id} | 
[**IssuesIssueIdGet**](IssuesAPI.md#IssuesIssueIdGet) | **Get** /issues/{issue_id} | 
[**IssuesIssueIdMovePost**](IssuesAPI.md#IssuesIssueIdMovePost) | **Post** /issues/{issue_id}/move | 
[**IssuesIssueIdPut**](IssuesAPI.md#IssuesIssueIdPut) | **Put** /issues/{issue_id} | 
[**IssuesIssueIdRemediatePost**](IssuesAPI.md#IssuesIssueIdRemediatePost) | **Post** /issues/{issue_id}/remediate | 
[**IssuesPost**](IssuesAPI.md#IssuesPost) | **Post** /issues | 



## IssuesGet

> IssuesGet200Response IssuesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.IssuesAPI.IssuesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.IssuesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssuesGet`: IssuesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.IssuesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssuesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**IssuesGet200Response**](IssuesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesIssueIdClosePost

> ReportsGenerateTimesheetsPost200Response IssuesIssueIdClosePost(ctx, issueId).Execute()





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
	issueId := int64(789) // int64 | Issue id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.IssuesIssueIdClosePost(context.Background(), issueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.IssuesIssueIdClosePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssuesIssueIdClosePost`: ReportsGenerateTimesheetsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.IssuesIssueIdClosePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueId** | **int64** | Issue id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesIssueIdClosePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReportsGenerateTimesheetsPost200Response**](ReportsGenerateTimesheetsPost200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesIssueIdDelete

> Issues IssuesIssueIdDelete(ctx, issueId).Execute()



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
	issueId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.IssuesIssueIdDelete(context.Background(), issueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.IssuesIssueIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssuesIssueIdDelete`: Issues
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.IssuesIssueIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesIssueIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Issues**](Issues.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesIssueIdGet

> Issues IssuesIssueIdGet(ctx, issueId).Include(include).Execute()



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
	issueId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.IssuesIssueIdGet(context.Background(), issueId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.IssuesIssueIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssuesIssueIdGet`: Issues
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.IssuesIssueIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesIssueIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**Issues**](Issues.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesIssueIdMovePost

> ReportsGenerateTimesheetsPost200Response IssuesIssueIdMovePost(ctx, issueId).IssuesIssueIdMovePostRequest(issuesIssueIdMovePostRequest).Execute()





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
	issueId := int64(789) // int64 | Issue id
	issuesIssueIdMovePostRequest := *openapiclient.NewIssuesIssueIdMovePostRequest() // IssuesIssueIdMovePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.IssuesIssueIdMovePost(context.Background(), issueId).IssuesIssueIdMovePostRequest(issuesIssueIdMovePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.IssuesIssueIdMovePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssuesIssueIdMovePost`: ReportsGenerateTimesheetsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.IssuesIssueIdMovePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueId** | **int64** | Issue id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesIssueIdMovePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **issuesIssueIdMovePostRequest** | [**IssuesIssueIdMovePostRequest**](IssuesIssueIdMovePostRequest.md) |  | 

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


## IssuesIssueIdPut

> Issues IssuesIssueIdPut(ctx, issueId).IssuesPut(issuesPut).Execute()



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
	issueId := int64(789) // int64 | Model id
	issuesPut := *openapiclient.NewIssuesPut() // IssuesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.IssuesIssueIdPut(context.Background(), issueId).IssuesPut(issuesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.IssuesIssueIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssuesIssueIdPut`: Issues
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.IssuesIssueIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesIssueIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **issuesPut** | [**IssuesPut**](IssuesPut.md) |  | 

### Return type

[**Issues**](Issues.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesIssueIdRemediatePost

> ReportsGenerateTimesheetsPost200Response IssuesIssueIdRemediatePost(ctx, issueId).Execute()





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
	issueId := int64(789) // int64 | Issue id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.IssuesIssueIdRemediatePost(context.Background(), issueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.IssuesIssueIdRemediatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssuesIssueIdRemediatePost`: ReportsGenerateTimesheetsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.IssuesIssueIdRemediatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueId** | **int64** | Issue id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesIssueIdRemediatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReportsGenerateTimesheetsPost200Response**](ReportsGenerateTimesheetsPost200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesPost

> IssuesGet200Response IssuesPost(ctx).IssuesPostRequest(issuesPostRequest).Execute()



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
	issuesPostRequest := *openapiclient.NewIssuesPostRequest() // IssuesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.IssuesPost(context.Background()).IssuesPostRequest(issuesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.IssuesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssuesPost`: IssuesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.IssuesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssuesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issuesPostRequest** | [**IssuesPostRequest**](IssuesPostRequest.md) |  | 

### Return type

[**IssuesGet200Response**](IssuesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

