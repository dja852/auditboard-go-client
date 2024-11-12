# \TaskPeriodsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TaskPeriodsGet**](TaskPeriodsAPI.md#TaskPeriodsGet) | **Get** /task_periods | 
[**TaskPeriodsPost**](TaskPeriodsAPI.md#TaskPeriodsPost) | **Post** /task_periods | 
[**TaskPeriodsTaskPeriodIdDelete**](TaskPeriodsAPI.md#TaskPeriodsTaskPeriodIdDelete) | **Delete** /task_periods/{task_period_id} | 
[**TaskPeriodsTaskPeriodIdGet**](TaskPeriodsAPI.md#TaskPeriodsTaskPeriodIdGet) | **Get** /task_periods/{task_period_id} | 
[**TaskPeriodsTaskPeriodIdPut**](TaskPeriodsAPI.md#TaskPeriodsTaskPeriodIdPut) | **Put** /task_periods/{task_period_id} | 



## TaskPeriodsGet

> TaskPeriodsGet200Response TaskPeriodsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.TaskPeriodsAPI.TaskPeriodsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskPeriodsAPI.TaskPeriodsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskPeriodsGet`: TaskPeriodsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TaskPeriodsAPI.TaskPeriodsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskPeriodsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**TaskPeriodsGet200Response**](TaskPeriodsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskPeriodsPost

> TaskPeriodsGet200Response TaskPeriodsPost(ctx).TaskPeriodsPostRequest(taskPeriodsPostRequest).Execute()



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
	taskPeriodsPostRequest := *openapiclient.NewTaskPeriodsPostRequest() // TaskPeriodsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskPeriodsAPI.TaskPeriodsPost(context.Background()).TaskPeriodsPostRequest(taskPeriodsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskPeriodsAPI.TaskPeriodsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskPeriodsPost`: TaskPeriodsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TaskPeriodsAPI.TaskPeriodsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskPeriodsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskPeriodsPostRequest** | [**TaskPeriodsPostRequest**](TaskPeriodsPostRequest.md) |  | 

### Return type

[**TaskPeriodsGet200Response**](TaskPeriodsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskPeriodsTaskPeriodIdDelete

> TaskPeriods TaskPeriodsTaskPeriodIdDelete(ctx, taskPeriodId).Execute()



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
	taskPeriodId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskPeriodsAPI.TaskPeriodsTaskPeriodIdDelete(context.Background(), taskPeriodId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskPeriodsAPI.TaskPeriodsTaskPeriodIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskPeriodsTaskPeriodIdDelete`: TaskPeriods
	fmt.Fprintf(os.Stdout, "Response from `TaskPeriodsAPI.TaskPeriodsTaskPeriodIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskPeriodId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskPeriodsTaskPeriodIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskPeriods**](TaskPeriods.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskPeriodsTaskPeriodIdGet

> TaskPeriods TaskPeriodsTaskPeriodIdGet(ctx, taskPeriodId).Include(include).Execute()



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
	taskPeriodId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskPeriodsAPI.TaskPeriodsTaskPeriodIdGet(context.Background(), taskPeriodId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskPeriodsAPI.TaskPeriodsTaskPeriodIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskPeriodsTaskPeriodIdGet`: TaskPeriods
	fmt.Fprintf(os.Stdout, "Response from `TaskPeriodsAPI.TaskPeriodsTaskPeriodIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskPeriodId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskPeriodsTaskPeriodIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**TaskPeriods**](TaskPeriods.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskPeriodsTaskPeriodIdPut

> TaskPeriods TaskPeriodsTaskPeriodIdPut(ctx, taskPeriodId).TaskPeriodsPut(taskPeriodsPut).Execute()



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
	taskPeriodId := int64(789) // int64 | Model id
	taskPeriodsPut := *openapiclient.NewTaskPeriodsPut() // TaskPeriodsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskPeriodsAPI.TaskPeriodsTaskPeriodIdPut(context.Background(), taskPeriodId).TaskPeriodsPut(taskPeriodsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskPeriodsAPI.TaskPeriodsTaskPeriodIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskPeriodsTaskPeriodIdPut`: TaskPeriods
	fmt.Fprintf(os.Stdout, "Response from `TaskPeriodsAPI.TaskPeriodsTaskPeriodIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskPeriodId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskPeriodsTaskPeriodIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskPeriodsPut** | [**TaskPeriodsPut**](TaskPeriodsPut.md) |  | 

### Return type

[**TaskPeriods**](TaskPeriods.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

