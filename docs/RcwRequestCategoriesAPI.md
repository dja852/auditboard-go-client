# \RcwRequestCategoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RcwRequestCategoriesGet**](RcwRequestCategoriesAPI.md#RcwRequestCategoriesGet) | **Get** /rcw_request_categories | 
[**RcwRequestCategoriesPost**](RcwRequestCategoriesAPI.md#RcwRequestCategoriesPost) | **Post** /rcw_request_categories | 
[**RcwRequestCategoriesRcwRequestCategoryIdDelete**](RcwRequestCategoriesAPI.md#RcwRequestCategoriesRcwRequestCategoryIdDelete) | **Delete** /rcw_request_categories/{rcw_request_category_id} | 
[**RcwRequestCategoriesRcwRequestCategoryIdGet**](RcwRequestCategoriesAPI.md#RcwRequestCategoriesRcwRequestCategoryIdGet) | **Get** /rcw_request_categories/{rcw_request_category_id} | 
[**RcwRequestCategoriesRcwRequestCategoryIdPut**](RcwRequestCategoriesAPI.md#RcwRequestCategoriesRcwRequestCategoryIdPut) | **Put** /rcw_request_categories/{rcw_request_category_id} | 



## RcwRequestCategoriesGet

> RcwRequestCategoriesGet200Response RcwRequestCategoriesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.RcwRequestCategoriesAPI.RcwRequestCategoriesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwRequestCategoriesAPI.RcwRequestCategoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwRequestCategoriesGet`: RcwRequestCategoriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RcwRequestCategoriesAPI.RcwRequestCategoriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRcwRequestCategoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**RcwRequestCategoriesGet200Response**](RcwRequestCategoriesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RcwRequestCategoriesPost

> RcwRequestCategoriesGet200Response RcwRequestCategoriesPost(ctx).RcwRequestCategoriesPostRequest(rcwRequestCategoriesPostRequest).Execute()



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
	rcwRequestCategoriesPostRequest := *openapiclient.NewRcwRequestCategoriesPostRequest() // RcwRequestCategoriesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RcwRequestCategoriesAPI.RcwRequestCategoriesPost(context.Background()).RcwRequestCategoriesPostRequest(rcwRequestCategoriesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwRequestCategoriesAPI.RcwRequestCategoriesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwRequestCategoriesPost`: RcwRequestCategoriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RcwRequestCategoriesAPI.RcwRequestCategoriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRcwRequestCategoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rcwRequestCategoriesPostRequest** | [**RcwRequestCategoriesPostRequest**](RcwRequestCategoriesPostRequest.md) |  | 

### Return type

[**RcwRequestCategoriesGet200Response**](RcwRequestCategoriesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RcwRequestCategoriesRcwRequestCategoryIdDelete

> RcwRequestCategories RcwRequestCategoriesRcwRequestCategoryIdDelete(ctx, rcwRequestCategoryId).Execute()



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
	rcwRequestCategoryId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RcwRequestCategoriesAPI.RcwRequestCategoriesRcwRequestCategoryIdDelete(context.Background(), rcwRequestCategoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwRequestCategoriesAPI.RcwRequestCategoriesRcwRequestCategoryIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwRequestCategoriesRcwRequestCategoryIdDelete`: RcwRequestCategories
	fmt.Fprintf(os.Stdout, "Response from `RcwRequestCategoriesAPI.RcwRequestCategoriesRcwRequestCategoryIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rcwRequestCategoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRcwRequestCategoriesRcwRequestCategoryIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RcwRequestCategories**](RcwRequestCategories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RcwRequestCategoriesRcwRequestCategoryIdGet

> RcwRequestCategories RcwRequestCategoriesRcwRequestCategoryIdGet(ctx, rcwRequestCategoryId).Include(include).Execute()



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
	rcwRequestCategoryId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RcwRequestCategoriesAPI.RcwRequestCategoriesRcwRequestCategoryIdGet(context.Background(), rcwRequestCategoryId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwRequestCategoriesAPI.RcwRequestCategoriesRcwRequestCategoryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwRequestCategoriesRcwRequestCategoryIdGet`: RcwRequestCategories
	fmt.Fprintf(os.Stdout, "Response from `RcwRequestCategoriesAPI.RcwRequestCategoriesRcwRequestCategoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rcwRequestCategoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRcwRequestCategoriesRcwRequestCategoryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**RcwRequestCategories**](RcwRequestCategories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RcwRequestCategoriesRcwRequestCategoryIdPut

> RcwRequestCategories RcwRequestCategoriesRcwRequestCategoryIdPut(ctx, rcwRequestCategoryId).RcwRequestCategoriesPut(rcwRequestCategoriesPut).Execute()



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
	rcwRequestCategoryId := int64(789) // int64 | Model id
	rcwRequestCategoriesPut := *openapiclient.NewRcwRequestCategoriesPut() // RcwRequestCategoriesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RcwRequestCategoriesAPI.RcwRequestCategoriesRcwRequestCategoryIdPut(context.Background(), rcwRequestCategoryId).RcwRequestCategoriesPut(rcwRequestCategoriesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwRequestCategoriesAPI.RcwRequestCategoriesRcwRequestCategoryIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwRequestCategoriesRcwRequestCategoryIdPut`: RcwRequestCategories
	fmt.Fprintf(os.Stdout, "Response from `RcwRequestCategoriesAPI.RcwRequestCategoriesRcwRequestCategoryIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rcwRequestCategoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRcwRequestCategoriesRcwRequestCategoryIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rcwRequestCategoriesPut** | [**RcwRequestCategoriesPut**](RcwRequestCategoriesPut.md) |  | 

### Return type

[**RcwRequestCategories**](RcwRequestCategories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

