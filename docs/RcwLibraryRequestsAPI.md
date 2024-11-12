# \RcwLibraryRequestsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RcwLibraryRequestsGet**](RcwLibraryRequestsAPI.md#RcwLibraryRequestsGet) | **Get** /rcw_library_requests | 
[**RcwLibraryRequestsPost**](RcwLibraryRequestsAPI.md#RcwLibraryRequestsPost) | **Post** /rcw_library_requests | 
[**RcwLibraryRequestsRcwLibraryRequestIdDelete**](RcwLibraryRequestsAPI.md#RcwLibraryRequestsRcwLibraryRequestIdDelete) | **Delete** /rcw_library_requests/{rcw_library_request_id} | 
[**RcwLibraryRequestsRcwLibraryRequestIdGet**](RcwLibraryRequestsAPI.md#RcwLibraryRequestsRcwLibraryRequestIdGet) | **Get** /rcw_library_requests/{rcw_library_request_id} | 
[**RcwLibraryRequestsRcwLibraryRequestIdPut**](RcwLibraryRequestsAPI.md#RcwLibraryRequestsRcwLibraryRequestIdPut) | **Put** /rcw_library_requests/{rcw_library_request_id} | 



## RcwLibraryRequestsGet

> RcwLibraryRequestsGet200Response RcwLibraryRequestsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.RcwLibraryRequestsAPI.RcwLibraryRequestsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwLibraryRequestsAPI.RcwLibraryRequestsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwLibraryRequestsGet`: RcwLibraryRequestsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RcwLibraryRequestsAPI.RcwLibraryRequestsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRcwLibraryRequestsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**RcwLibraryRequestsGet200Response**](RcwLibraryRequestsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RcwLibraryRequestsPost

> RcwLibraryRequestsGet200Response RcwLibraryRequestsPost(ctx).RcwLibraryRequestsPostRequest(rcwLibraryRequestsPostRequest).Execute()



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
	rcwLibraryRequestsPostRequest := *openapiclient.NewRcwLibraryRequestsPostRequest() // RcwLibraryRequestsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RcwLibraryRequestsAPI.RcwLibraryRequestsPost(context.Background()).RcwLibraryRequestsPostRequest(rcwLibraryRequestsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwLibraryRequestsAPI.RcwLibraryRequestsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwLibraryRequestsPost`: RcwLibraryRequestsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RcwLibraryRequestsAPI.RcwLibraryRequestsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRcwLibraryRequestsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rcwLibraryRequestsPostRequest** | [**RcwLibraryRequestsPostRequest**](RcwLibraryRequestsPostRequest.md) |  | 

### Return type

[**RcwLibraryRequestsGet200Response**](RcwLibraryRequestsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RcwLibraryRequestsRcwLibraryRequestIdDelete

> RcwLibraryRequests RcwLibraryRequestsRcwLibraryRequestIdDelete(ctx, rcwLibraryRequestId).Execute()



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
	rcwLibraryRequestId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RcwLibraryRequestsAPI.RcwLibraryRequestsRcwLibraryRequestIdDelete(context.Background(), rcwLibraryRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwLibraryRequestsAPI.RcwLibraryRequestsRcwLibraryRequestIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwLibraryRequestsRcwLibraryRequestIdDelete`: RcwLibraryRequests
	fmt.Fprintf(os.Stdout, "Response from `RcwLibraryRequestsAPI.RcwLibraryRequestsRcwLibraryRequestIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rcwLibraryRequestId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRcwLibraryRequestsRcwLibraryRequestIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RcwLibraryRequests**](RcwLibraryRequests.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RcwLibraryRequestsRcwLibraryRequestIdGet

> RcwLibraryRequests RcwLibraryRequestsRcwLibraryRequestIdGet(ctx, rcwLibraryRequestId).Include(include).Execute()



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
	rcwLibraryRequestId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RcwLibraryRequestsAPI.RcwLibraryRequestsRcwLibraryRequestIdGet(context.Background(), rcwLibraryRequestId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwLibraryRequestsAPI.RcwLibraryRequestsRcwLibraryRequestIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwLibraryRequestsRcwLibraryRequestIdGet`: RcwLibraryRequests
	fmt.Fprintf(os.Stdout, "Response from `RcwLibraryRequestsAPI.RcwLibraryRequestsRcwLibraryRequestIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rcwLibraryRequestId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRcwLibraryRequestsRcwLibraryRequestIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**RcwLibraryRequests**](RcwLibraryRequests.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RcwLibraryRequestsRcwLibraryRequestIdPut

> RcwLibraryRequests RcwLibraryRequestsRcwLibraryRequestIdPut(ctx, rcwLibraryRequestId).RcwLibraryRequestsPut(rcwLibraryRequestsPut).Execute()



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
	rcwLibraryRequestId := int64(789) // int64 | Model id
	rcwLibraryRequestsPut := *openapiclient.NewRcwLibraryRequestsPut() // RcwLibraryRequestsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RcwLibraryRequestsAPI.RcwLibraryRequestsRcwLibraryRequestIdPut(context.Background(), rcwLibraryRequestId).RcwLibraryRequestsPut(rcwLibraryRequestsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwLibraryRequestsAPI.RcwLibraryRequestsRcwLibraryRequestIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwLibraryRequestsRcwLibraryRequestIdPut`: RcwLibraryRequests
	fmt.Fprintf(os.Stdout, "Response from `RcwLibraryRequestsAPI.RcwLibraryRequestsRcwLibraryRequestIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rcwLibraryRequestId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRcwLibraryRequestsRcwLibraryRequestIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rcwLibraryRequestsPut** | [**RcwLibraryRequestsPut**](RcwLibraryRequestsPut.md) |  | 

### Return type

[**RcwLibraryRequests**](RcwLibraryRequests.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

