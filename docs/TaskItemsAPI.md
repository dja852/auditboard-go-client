# \TaskItemsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TaskItemsGet**](TaskItemsAPI.md#TaskItemsGet) | **Get** /task_items | 
[**TaskItemsImportJsonPost**](TaskItemsAPI.md#TaskItemsImportJsonPost) | **Post** /task_items/import/json | 
[**TaskItemsPost**](TaskItemsAPI.md#TaskItemsPost) | **Post** /task_items | 
[**TaskItemsTaskItemIdDelete**](TaskItemsAPI.md#TaskItemsTaskItemIdDelete) | **Delete** /task_items/{task_item_id} | 
[**TaskItemsTaskItemIdGet**](TaskItemsAPI.md#TaskItemsTaskItemIdGet) | **Get** /task_items/{task_item_id} | 
[**TaskItemsTaskItemIdPut**](TaskItemsAPI.md#TaskItemsTaskItemIdPut) | **Put** /task_items/{task_item_id} | 
[**TaskItemsWorkstreamTaskIdCertifyPost**](TaskItemsAPI.md#TaskItemsWorkstreamTaskIdCertifyPost) | **Post** /task_items/{workstream_task_id}/certify | 
[**TaskItemsWorkstreamTaskIdClosePost**](TaskItemsAPI.md#TaskItemsWorkstreamTaskIdClosePost) | **Post** /task_items/{workstream_task_id}/close | 
[**TaskItemsWorkstreamTaskIdSubmitPost**](TaskItemsAPI.md#TaskItemsWorkstreamTaskIdSubmitPost) | **Post** /task_items/{workstream_task_id}/submit | 



## TaskItemsGet

> TaskItemsGet200Response TaskItemsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.TaskItemsAPI.TaskItemsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskItemsAPI.TaskItemsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskItemsGet`: TaskItemsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TaskItemsAPI.TaskItemsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskItemsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**TaskItemsGet200Response**](TaskItemsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskItemsImportJsonPost

> TaskItemsImportJsonPost200Response TaskItemsImportJsonPost(ctx).TaskItemsImportJsonPostRequest(taskItemsImportJsonPostRequest).Execute()



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
	taskItemsImportJsonPostRequest := *openapiclient.NewTaskItemsImportJsonPostRequest() // TaskItemsImportJsonPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskItemsAPI.TaskItemsImportJsonPost(context.Background()).TaskItemsImportJsonPostRequest(taskItemsImportJsonPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskItemsAPI.TaskItemsImportJsonPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskItemsImportJsonPost`: TaskItemsImportJsonPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TaskItemsAPI.TaskItemsImportJsonPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskItemsImportJsonPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskItemsImportJsonPostRequest** | [**TaskItemsImportJsonPostRequest**](TaskItemsImportJsonPostRequest.md) |  | 

### Return type

[**TaskItemsImportJsonPost200Response**](TaskItemsImportJsonPost200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskItemsPost

> TaskItemsGet200Response TaskItemsPost(ctx).TaskItemsPostRequest(taskItemsPostRequest).Execute()



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
	taskItemsPostRequest := *openapiclient.NewTaskItemsPostRequest() // TaskItemsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskItemsAPI.TaskItemsPost(context.Background()).TaskItemsPostRequest(taskItemsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskItemsAPI.TaskItemsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskItemsPost`: TaskItemsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TaskItemsAPI.TaskItemsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskItemsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskItemsPostRequest** | [**TaskItemsPostRequest**](TaskItemsPostRequest.md) |  | 

### Return type

[**TaskItemsGet200Response**](TaskItemsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskItemsTaskItemIdDelete

> TaskItems TaskItemsTaskItemIdDelete(ctx, taskItemId).Execute()



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
	taskItemId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskItemsAPI.TaskItemsTaskItemIdDelete(context.Background(), taskItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskItemsAPI.TaskItemsTaskItemIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskItemsTaskItemIdDelete`: TaskItems
	fmt.Fprintf(os.Stdout, "Response from `TaskItemsAPI.TaskItemsTaskItemIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskItemId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskItemsTaskItemIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskItems**](TaskItems.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskItemsTaskItemIdGet

> TaskItems TaskItemsTaskItemIdGet(ctx, taskItemId).Include(include).Execute()



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
	taskItemId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskItemsAPI.TaskItemsTaskItemIdGet(context.Background(), taskItemId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskItemsAPI.TaskItemsTaskItemIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskItemsTaskItemIdGet`: TaskItems
	fmt.Fprintf(os.Stdout, "Response from `TaskItemsAPI.TaskItemsTaskItemIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskItemId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskItemsTaskItemIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**TaskItems**](TaskItems.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskItemsTaskItemIdPut

> TaskItems TaskItemsTaskItemIdPut(ctx, taskItemId).TaskItemsPut(taskItemsPut).Execute()



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
	taskItemId := int64(789) // int64 | Model id
	taskItemsPut := *openapiclient.NewTaskItemsPut() // TaskItemsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskItemsAPI.TaskItemsTaskItemIdPut(context.Background(), taskItemId).TaskItemsPut(taskItemsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskItemsAPI.TaskItemsTaskItemIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskItemsTaskItemIdPut`: TaskItems
	fmt.Fprintf(os.Stdout, "Response from `TaskItemsAPI.TaskItemsTaskItemIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskItemId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskItemsTaskItemIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskItemsPut** | [**TaskItemsPut**](TaskItemsPut.md) |  | 

### Return type

[**TaskItems**](TaskItems.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskItemsWorkstreamTaskIdCertifyPost

> TaskItemsWorkstreamTaskIdClosePost200Response TaskItemsWorkstreamTaskIdCertifyPost(ctx, workstreamTaskId).Execute()





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
	workstreamTaskId := int64(789) // int64 | Workstream Task id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskItemsAPI.TaskItemsWorkstreamTaskIdCertifyPost(context.Background(), workstreamTaskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskItemsAPI.TaskItemsWorkstreamTaskIdCertifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskItemsWorkstreamTaskIdCertifyPost`: TaskItemsWorkstreamTaskIdClosePost200Response
	fmt.Fprintf(os.Stdout, "Response from `TaskItemsAPI.TaskItemsWorkstreamTaskIdCertifyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workstreamTaskId** | **int64** | Workstream Task id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskItemsWorkstreamTaskIdCertifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskItemsWorkstreamTaskIdClosePost200Response**](TaskItemsWorkstreamTaskIdClosePost200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskItemsWorkstreamTaskIdClosePost

> TaskItemsWorkstreamTaskIdClosePost200Response TaskItemsWorkstreamTaskIdClosePost(ctx, workstreamTaskId).TaskItemsWorkstreamTaskIdClosePostRequest(taskItemsWorkstreamTaskIdClosePostRequest).Execute()





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
	workstreamTaskId := int64(789) // int64 | Workstream Task id
	taskItemsWorkstreamTaskIdClosePostRequest := *openapiclient.NewTaskItemsWorkstreamTaskIdClosePostRequest() // TaskItemsWorkstreamTaskIdClosePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskItemsAPI.TaskItemsWorkstreamTaskIdClosePost(context.Background(), workstreamTaskId).TaskItemsWorkstreamTaskIdClosePostRequest(taskItemsWorkstreamTaskIdClosePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskItemsAPI.TaskItemsWorkstreamTaskIdClosePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskItemsWorkstreamTaskIdClosePost`: TaskItemsWorkstreamTaskIdClosePost200Response
	fmt.Fprintf(os.Stdout, "Response from `TaskItemsAPI.TaskItemsWorkstreamTaskIdClosePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workstreamTaskId** | **int64** | Workstream Task id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskItemsWorkstreamTaskIdClosePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskItemsWorkstreamTaskIdClosePostRequest** | [**TaskItemsWorkstreamTaskIdClosePostRequest**](TaskItemsWorkstreamTaskIdClosePostRequest.md) |  | 

### Return type

[**TaskItemsWorkstreamTaskIdClosePost200Response**](TaskItemsWorkstreamTaskIdClosePost200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskItemsWorkstreamTaskIdSubmitPost

> TaskItemsSubmit TaskItemsWorkstreamTaskIdSubmitPost(ctx, workstreamTaskId).Execute()





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
	workstreamTaskId := int64(789) // int64 | Workstream Task id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskItemsAPI.TaskItemsWorkstreamTaskIdSubmitPost(context.Background(), workstreamTaskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskItemsAPI.TaskItemsWorkstreamTaskIdSubmitPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskItemsWorkstreamTaskIdSubmitPost`: TaskItemsSubmit
	fmt.Fprintf(os.Stdout, "Response from `TaskItemsAPI.TaskItemsWorkstreamTaskIdSubmitPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workstreamTaskId** | **int64** | Workstream Task id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskItemsWorkstreamTaskIdSubmitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskItemsSubmit**](TaskItemsSubmit.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

