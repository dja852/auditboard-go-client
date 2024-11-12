# \TaskGroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TaskGroupsGet**](TaskGroupsAPI.md#TaskGroupsGet) | **Get** /task_groups | 
[**TaskGroupsPost**](TaskGroupsAPI.md#TaskGroupsPost) | **Post** /task_groups | 
[**TaskGroupsTaskGroupIdBulkEnableItemsPut**](TaskGroupsAPI.md#TaskGroupsTaskGroupIdBulkEnableItemsPut) | **Put** /task_groups/{task_group_id}/bulk_enable_items | 
[**TaskGroupsTaskGroupIdDelete**](TaskGroupsAPI.md#TaskGroupsTaskGroupIdDelete) | **Delete** /task_groups/{task_group_id} | 
[**TaskGroupsTaskGroupIdGet**](TaskGroupsAPI.md#TaskGroupsTaskGroupIdGet) | **Get** /task_groups/{task_group_id} | 
[**TaskGroupsTaskGroupIdPut**](TaskGroupsAPI.md#TaskGroupsTaskGroupIdPut) | **Put** /task_groups/{task_group_id} | 
[**TaskGroupsTaskGroupIdStartPut**](TaskGroupsAPI.md#TaskGroupsTaskGroupIdStartPut) | **Put** /task_groups/{task_group_id}/start | 



## TaskGroupsGet

> TaskGroupsGet200Response TaskGroupsGet(ctx).Include(include).Execute()





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
	resp, r, err := apiClient.TaskGroupsAPI.TaskGroupsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskGroupsAPI.TaskGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskGroupsGet`: TaskGroupsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TaskGroupsAPI.TaskGroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**TaskGroupsGet200Response**](TaskGroupsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskGroupsPost

> TaskGroupsGet200Response TaskGroupsPost(ctx).TaskGroupsPostRequest(taskGroupsPostRequest).Execute()





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
	taskGroupsPostRequest := *openapiclient.NewTaskGroupsPostRequest() // TaskGroupsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskGroupsAPI.TaskGroupsPost(context.Background()).TaskGroupsPostRequest(taskGroupsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskGroupsAPI.TaskGroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskGroupsPost`: TaskGroupsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TaskGroupsAPI.TaskGroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskGroupsPostRequest** | [**TaskGroupsPostRequest**](TaskGroupsPostRequest.md) |  | 

### Return type

[**TaskGroupsGet200Response**](TaskGroupsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskGroupsTaskGroupIdBulkEnableItemsPut

> TaskGroupsGroupsBulkEnableItems TaskGroupsTaskGroupIdBulkEnableItemsPut(ctx, taskGroupId).TaskGroupsTaskGroupIdBulkEnableItemsPutRequest(taskGroupsTaskGroupIdBulkEnableItemsPutRequest).Execute()





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
	taskGroupId := int64(789) // int64 | Task group id, aka Workstream project id
	taskGroupsTaskGroupIdBulkEnableItemsPutRequest := *openapiclient.NewTaskGroupsTaskGroupIdBulkEnableItemsPutRequest() // TaskGroupsTaskGroupIdBulkEnableItemsPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskGroupsAPI.TaskGroupsTaskGroupIdBulkEnableItemsPut(context.Background(), taskGroupId).TaskGroupsTaskGroupIdBulkEnableItemsPutRequest(taskGroupsTaskGroupIdBulkEnableItemsPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskGroupsAPI.TaskGroupsTaskGroupIdBulkEnableItemsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskGroupsTaskGroupIdBulkEnableItemsPut`: TaskGroupsGroupsBulkEnableItems
	fmt.Fprintf(os.Stdout, "Response from `TaskGroupsAPI.TaskGroupsTaskGroupIdBulkEnableItemsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskGroupId** | **int64** | Task group id, aka Workstream project id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskGroupsTaskGroupIdBulkEnableItemsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskGroupsTaskGroupIdBulkEnableItemsPutRequest** | [**TaskGroupsTaskGroupIdBulkEnableItemsPutRequest**](TaskGroupsTaskGroupIdBulkEnableItemsPutRequest.md) |  | 

### Return type

[**TaskGroupsGroupsBulkEnableItems**](TaskGroupsGroupsBulkEnableItems.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskGroupsTaskGroupIdDelete

> TaskGroups TaskGroupsTaskGroupIdDelete(ctx, taskGroupId).Execute()





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
	taskGroupId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskGroupsAPI.TaskGroupsTaskGroupIdDelete(context.Background(), taskGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskGroupsAPI.TaskGroupsTaskGroupIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskGroupsTaskGroupIdDelete`: TaskGroups
	fmt.Fprintf(os.Stdout, "Response from `TaskGroupsAPI.TaskGroupsTaskGroupIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskGroupId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskGroupsTaskGroupIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskGroups**](TaskGroups.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskGroupsTaskGroupIdGet

> TaskGroups TaskGroupsTaskGroupIdGet(ctx, taskGroupId).Include(include).Execute()





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
	taskGroupId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskGroupsAPI.TaskGroupsTaskGroupIdGet(context.Background(), taskGroupId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskGroupsAPI.TaskGroupsTaskGroupIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskGroupsTaskGroupIdGet`: TaskGroups
	fmt.Fprintf(os.Stdout, "Response from `TaskGroupsAPI.TaskGroupsTaskGroupIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskGroupId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskGroupsTaskGroupIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**TaskGroups**](TaskGroups.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskGroupsTaskGroupIdPut

> TaskGroups TaskGroupsTaskGroupIdPut(ctx, taskGroupId).TaskGroupsPut(taskGroupsPut).Execute()





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
	taskGroupId := int64(789) // int64 | Model id
	taskGroupsPut := *openapiclient.NewTaskGroupsPut() // TaskGroupsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskGroupsAPI.TaskGroupsTaskGroupIdPut(context.Background(), taskGroupId).TaskGroupsPut(taskGroupsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskGroupsAPI.TaskGroupsTaskGroupIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskGroupsTaskGroupIdPut`: TaskGroups
	fmt.Fprintf(os.Stdout, "Response from `TaskGroupsAPI.TaskGroupsTaskGroupIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskGroupId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskGroupsTaskGroupIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskGroupsPut** | [**TaskGroupsPut**](TaskGroupsPut.md) |  | 

### Return type

[**TaskGroups**](TaskGroups.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskGroupsTaskGroupIdStartPut

> TaskGroupsStart TaskGroupsTaskGroupIdStartPut(ctx, taskGroupId).Body(body).Execute()



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
	taskGroupId := int64(789) // int64 | Task group id, aka Workstream project id
	body := true // bool | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskGroupsAPI.TaskGroupsTaskGroupIdStartPut(context.Background(), taskGroupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskGroupsAPI.TaskGroupsTaskGroupIdStartPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskGroupsTaskGroupIdStartPut`: TaskGroupsStart
	fmt.Fprintf(os.Stdout, "Response from `TaskGroupsAPI.TaskGroupsTaskGroupIdStartPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskGroupId** | **int64** | Task group id, aka Workstream project id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskGroupsTaskGroupIdStartPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **bool** |  | 

### Return type

[**TaskGroupsStart**](TaskGroupsStart.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

