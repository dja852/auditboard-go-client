# \RcwRequestsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RcwRequestsGet**](RcwRequestsAPI.md#RcwRequestsGet) | **Get** /rcw_requests | 
[**RcwRequestsPost**](RcwRequestsAPI.md#RcwRequestsPost) | **Post** /rcw_requests | 
[**RcwRequestsRcwRequestIdDelete**](RcwRequestsAPI.md#RcwRequestsRcwRequestIdDelete) | **Delete** /rcw_requests/{rcw_request_id} | 
[**RcwRequestsRcwRequestIdGet**](RcwRequestsAPI.md#RcwRequestsRcwRequestIdGet) | **Get** /rcw_requests/{rcw_request_id} | 
[**RcwRequestsRcwRequestIdPut**](RcwRequestsAPI.md#RcwRequestsRcwRequestIdPut) | **Put** /rcw_requests/{rcw_request_id} | 



## RcwRequestsGet

> RcwRequestsGet200Response RcwRequestsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.RcwRequestsAPI.RcwRequestsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwRequestsAPI.RcwRequestsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwRequestsGet`: RcwRequestsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RcwRequestsAPI.RcwRequestsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRcwRequestsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**RcwRequestsGet200Response**](RcwRequestsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RcwRequestsPost

> RcwRequestsGet200Response RcwRequestsPost(ctx).RcwRequestsPostRequest(rcwRequestsPostRequest).Execute()



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
	rcwRequestsPostRequest := *openapiclient.NewRcwRequestsPostRequest() // RcwRequestsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RcwRequestsAPI.RcwRequestsPost(context.Background()).RcwRequestsPostRequest(rcwRequestsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwRequestsAPI.RcwRequestsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwRequestsPost`: RcwRequestsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RcwRequestsAPI.RcwRequestsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRcwRequestsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rcwRequestsPostRequest** | [**RcwRequestsPostRequest**](RcwRequestsPostRequest.md) |  | 

### Return type

[**RcwRequestsGet200Response**](RcwRequestsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RcwRequestsRcwRequestIdDelete

> RcwRequests RcwRequestsRcwRequestIdDelete(ctx, rcwRequestId).Execute()



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
	rcwRequestId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RcwRequestsAPI.RcwRequestsRcwRequestIdDelete(context.Background(), rcwRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwRequestsAPI.RcwRequestsRcwRequestIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwRequestsRcwRequestIdDelete`: RcwRequests
	fmt.Fprintf(os.Stdout, "Response from `RcwRequestsAPI.RcwRequestsRcwRequestIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rcwRequestId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRcwRequestsRcwRequestIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RcwRequests**](RcwRequests.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RcwRequestsRcwRequestIdGet

> RcwRequests RcwRequestsRcwRequestIdGet(ctx, rcwRequestId).Include(include).Execute()



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
	rcwRequestId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RcwRequestsAPI.RcwRequestsRcwRequestIdGet(context.Background(), rcwRequestId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwRequestsAPI.RcwRequestsRcwRequestIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwRequestsRcwRequestIdGet`: RcwRequests
	fmt.Fprintf(os.Stdout, "Response from `RcwRequestsAPI.RcwRequestsRcwRequestIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rcwRequestId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRcwRequestsRcwRequestIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**RcwRequests**](RcwRequests.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RcwRequestsRcwRequestIdPut

> RcwRequests RcwRequestsRcwRequestIdPut(ctx, rcwRequestId).RcwRequestsPut(rcwRequestsPut).Execute()



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
	rcwRequestId := int64(789) // int64 | Model id
	rcwRequestsPut := *openapiclient.NewRcwRequestsPut() // RcwRequestsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RcwRequestsAPI.RcwRequestsRcwRequestIdPut(context.Background(), rcwRequestId).RcwRequestsPut(rcwRequestsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwRequestsAPI.RcwRequestsRcwRequestIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwRequestsRcwRequestIdPut`: RcwRequests
	fmt.Fprintf(os.Stdout, "Response from `RcwRequestsAPI.RcwRequestsRcwRequestIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rcwRequestId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRcwRequestsRcwRequestIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rcwRequestsPut** | [**RcwRequestsPut**](RcwRequestsPut.md) |  | 

### Return type

[**RcwRequests**](RcwRequests.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

