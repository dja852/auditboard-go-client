# \ProcessesDataAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProcessesDataGet**](ProcessesDataAPI.md#ProcessesDataGet) | **Get** /processes_data | 
[**ProcessesDataPost**](ProcessesDataAPI.md#ProcessesDataPost) | **Post** /processes_data | 
[**ProcessesDataProcessesDatumIdDelete**](ProcessesDataAPI.md#ProcessesDataProcessesDatumIdDelete) | **Delete** /processes_data/{processes_datum_id} | 
[**ProcessesDataProcessesDatumIdGet**](ProcessesDataAPI.md#ProcessesDataProcessesDatumIdGet) | **Get** /processes_data/{processes_datum_id} | 
[**ProcessesDataProcessesDatumIdPut**](ProcessesDataAPI.md#ProcessesDataProcessesDatumIdPut) | **Put** /processes_data/{processes_datum_id} | 



## ProcessesDataGet

> ProcessesDataGet200Response ProcessesDataGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.ProcessesDataAPI.ProcessesDataGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessesDataAPI.ProcessesDataGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessesDataGet`: ProcessesDataGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ProcessesDataAPI.ProcessesDataGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessesDataGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**ProcessesDataGet200Response**](ProcessesDataGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessesDataPost

> ProcessesDataGet200Response ProcessesDataPost(ctx).ProcessesDataPostRequest(processesDataPostRequest).Execute()



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
	processesDataPostRequest := *openapiclient.NewProcessesDataPostRequest() // ProcessesDataPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessesDataAPI.ProcessesDataPost(context.Background()).ProcessesDataPostRequest(processesDataPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessesDataAPI.ProcessesDataPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessesDataPost`: ProcessesDataGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ProcessesDataAPI.ProcessesDataPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessesDataPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processesDataPostRequest** | [**ProcessesDataPostRequest**](ProcessesDataPostRequest.md) |  | 

### Return type

[**ProcessesDataGet200Response**](ProcessesDataGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessesDataProcessesDatumIdDelete

> ProcessesData ProcessesDataProcessesDatumIdDelete(ctx, processesDatumId).Execute()



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
	processesDatumId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessesDataAPI.ProcessesDataProcessesDatumIdDelete(context.Background(), processesDatumId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessesDataAPI.ProcessesDataProcessesDatumIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessesDataProcessesDatumIdDelete`: ProcessesData
	fmt.Fprintf(os.Stdout, "Response from `ProcessesDataAPI.ProcessesDataProcessesDatumIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processesDatumId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessesDataProcessesDatumIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProcessesData**](ProcessesData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessesDataProcessesDatumIdGet

> ProcessesData ProcessesDataProcessesDatumIdGet(ctx, processesDatumId).Include(include).Execute()



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
	processesDatumId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessesDataAPI.ProcessesDataProcessesDatumIdGet(context.Background(), processesDatumId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessesDataAPI.ProcessesDataProcessesDatumIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessesDataProcessesDatumIdGet`: ProcessesData
	fmt.Fprintf(os.Stdout, "Response from `ProcessesDataAPI.ProcessesDataProcessesDatumIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processesDatumId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessesDataProcessesDatumIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**ProcessesData**](ProcessesData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessesDataProcessesDatumIdPut

> ProcessesData ProcessesDataProcessesDatumIdPut(ctx, processesDatumId).ProcessesDataPut(processesDataPut).Execute()



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
	processesDatumId := int64(789) // int64 | Model id
	processesDataPut := *openapiclient.NewProcessesDataPut() // ProcessesDataPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessesDataAPI.ProcessesDataProcessesDatumIdPut(context.Background(), processesDatumId).ProcessesDataPut(processesDataPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessesDataAPI.ProcessesDataProcessesDatumIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessesDataProcessesDatumIdPut`: ProcessesData
	fmt.Fprintf(os.Stdout, "Response from `ProcessesDataAPI.ProcessesDataProcessesDatumIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processesDatumId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessesDataProcessesDatumIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **processesDataPut** | [**ProcessesDataPut**](ProcessesDataPut.md) |  | 

### Return type

[**ProcessesData**](ProcessesData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

