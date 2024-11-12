# \TaskCategoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TaskCategoriesGet**](TaskCategoriesAPI.md#TaskCategoriesGet) | **Get** /task_categories | 
[**TaskCategoriesPost**](TaskCategoriesAPI.md#TaskCategoriesPost) | **Post** /task_categories | 
[**TaskCategoriesTaskCategoryIdDelete**](TaskCategoriesAPI.md#TaskCategoriesTaskCategoryIdDelete) | **Delete** /task_categories/{task_category_id} | 
[**TaskCategoriesTaskCategoryIdGet**](TaskCategoriesAPI.md#TaskCategoriesTaskCategoryIdGet) | **Get** /task_categories/{task_category_id} | 
[**TaskCategoriesTaskCategoryIdPut**](TaskCategoriesAPI.md#TaskCategoriesTaskCategoryIdPut) | **Put** /task_categories/{task_category_id} | 



## TaskCategoriesGet

> TaskCategoriesGet200Response TaskCategoriesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.TaskCategoriesAPI.TaskCategoriesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskCategoriesAPI.TaskCategoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskCategoriesGet`: TaskCategoriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TaskCategoriesAPI.TaskCategoriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskCategoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**TaskCategoriesGet200Response**](TaskCategoriesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskCategoriesPost

> TaskCategoriesGet200Response TaskCategoriesPost(ctx).TaskCategoriesPostRequest(taskCategoriesPostRequest).Execute()



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
	taskCategoriesPostRequest := *openapiclient.NewTaskCategoriesPostRequest() // TaskCategoriesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskCategoriesAPI.TaskCategoriesPost(context.Background()).TaskCategoriesPostRequest(taskCategoriesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskCategoriesAPI.TaskCategoriesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskCategoriesPost`: TaskCategoriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TaskCategoriesAPI.TaskCategoriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskCategoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskCategoriesPostRequest** | [**TaskCategoriesPostRequest**](TaskCategoriesPostRequest.md) |  | 

### Return type

[**TaskCategoriesGet200Response**](TaskCategoriesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskCategoriesTaskCategoryIdDelete

> TaskCategories TaskCategoriesTaskCategoryIdDelete(ctx, taskCategoryId).Execute()



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
	taskCategoryId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskCategoriesAPI.TaskCategoriesTaskCategoryIdDelete(context.Background(), taskCategoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskCategoriesAPI.TaskCategoriesTaskCategoryIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskCategoriesTaskCategoryIdDelete`: TaskCategories
	fmt.Fprintf(os.Stdout, "Response from `TaskCategoriesAPI.TaskCategoriesTaskCategoryIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskCategoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskCategoriesTaskCategoryIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskCategories**](TaskCategories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskCategoriesTaskCategoryIdGet

> TaskCategories TaskCategoriesTaskCategoryIdGet(ctx, taskCategoryId).Include(include).Execute()



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
	taskCategoryId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskCategoriesAPI.TaskCategoriesTaskCategoryIdGet(context.Background(), taskCategoryId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskCategoriesAPI.TaskCategoriesTaskCategoryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskCategoriesTaskCategoryIdGet`: TaskCategories
	fmt.Fprintf(os.Stdout, "Response from `TaskCategoriesAPI.TaskCategoriesTaskCategoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskCategoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskCategoriesTaskCategoryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**TaskCategories**](TaskCategories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskCategoriesTaskCategoryIdPut

> TaskCategories TaskCategoriesTaskCategoryIdPut(ctx, taskCategoryId).TaskCategoriesPut(taskCategoriesPut).Execute()



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
	taskCategoryId := int64(789) // int64 | Model id
	taskCategoriesPut := *openapiclient.NewTaskCategoriesPut() // TaskCategoriesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskCategoriesAPI.TaskCategoriesTaskCategoryIdPut(context.Background(), taskCategoryId).TaskCategoriesPut(taskCategoriesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskCategoriesAPI.TaskCategoriesTaskCategoryIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskCategoriesTaskCategoryIdPut`: TaskCategories
	fmt.Fprintf(os.Stdout, "Response from `TaskCategoriesAPI.TaskCategoriesTaskCategoryIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskCategoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskCategoriesTaskCategoryIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskCategoriesPut** | [**TaskCategoriesPut**](TaskCategoriesPut.md) |  | 

### Return type

[**TaskCategories**](TaskCategories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

