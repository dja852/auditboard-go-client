# \ProcessesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProcessesGet**](ProcessesAPI.md#ProcessesGet) | **Get** /processes | 
[**ProcessesPost**](ProcessesAPI.md#ProcessesPost) | **Post** /processes | 
[**ProcessesProcessIdDelete**](ProcessesAPI.md#ProcessesProcessIdDelete) | **Delete** /processes/{process_id} | 
[**ProcessesProcessIdGet**](ProcessesAPI.md#ProcessesProcessIdGet) | **Get** /processes/{process_id} | 
[**ProcessesProcessIdPut**](ProcessesAPI.md#ProcessesProcessIdPut) | **Put** /processes/{process_id} | 



## ProcessesGet

> ProcessesGet200Response ProcessesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.ProcessesAPI.ProcessesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessesAPI.ProcessesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessesGet`: ProcessesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ProcessesAPI.ProcessesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**ProcessesGet200Response**](ProcessesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessesPost

> ProcessesGet200Response ProcessesPost(ctx).ProcessesPostRequest(processesPostRequest).Execute()



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
	processesPostRequest := *openapiclient.NewProcessesPostRequest() // ProcessesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessesAPI.ProcessesPost(context.Background()).ProcessesPostRequest(processesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessesAPI.ProcessesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessesPost`: ProcessesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ProcessesAPI.ProcessesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processesPostRequest** | [**ProcessesPostRequest**](ProcessesPostRequest.md) |  | 

### Return type

[**ProcessesGet200Response**](ProcessesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessesProcessIdDelete

> Processes ProcessesProcessIdDelete(ctx, processId).Execute()



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
	processId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessesAPI.ProcessesProcessIdDelete(context.Background(), processId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessesAPI.ProcessesProcessIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessesProcessIdDelete`: Processes
	fmt.Fprintf(os.Stdout, "Response from `ProcessesAPI.ProcessesProcessIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessesProcessIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Processes**](Processes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessesProcessIdGet

> Processes ProcessesProcessIdGet(ctx, processId).Include(include).Execute()



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
	processId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessesAPI.ProcessesProcessIdGet(context.Background(), processId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessesAPI.ProcessesProcessIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessesProcessIdGet`: Processes
	fmt.Fprintf(os.Stdout, "Response from `ProcessesAPI.ProcessesProcessIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessesProcessIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**Processes**](Processes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessesProcessIdPut

> Processes ProcessesProcessIdPut(ctx, processId).ProcessesPut(processesPut).Execute()



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
	processId := int64(789) // int64 | Model id
	processesPut := *openapiclient.NewProcessesPut() // ProcessesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessesAPI.ProcessesProcessIdPut(context.Background(), processId).ProcessesPut(processesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessesAPI.ProcessesProcessIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessesProcessIdPut`: Processes
	fmt.Fprintf(os.Stdout, "Response from `ProcessesAPI.ProcessesProcessIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessesProcessIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **processesPut** | [**ProcessesPut**](ProcessesPut.md) |  | 

### Return type

[**Processes**](Processes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

