# \FrameworkItemsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FrameworkItemsFrameworkItemIdDelete**](FrameworkItemsAPI.md#FrameworkItemsFrameworkItemIdDelete) | **Delete** /framework_items/{framework_item_id} | 
[**FrameworkItemsFrameworkItemIdGet**](FrameworkItemsAPI.md#FrameworkItemsFrameworkItemIdGet) | **Get** /framework_items/{framework_item_id} | 
[**FrameworkItemsFrameworkItemIdPut**](FrameworkItemsAPI.md#FrameworkItemsFrameworkItemIdPut) | **Put** /framework_items/{framework_item_id} | 
[**FrameworkItemsGet**](FrameworkItemsAPI.md#FrameworkItemsGet) | **Get** /framework_items | 
[**FrameworkItemsPost**](FrameworkItemsAPI.md#FrameworkItemsPost) | **Post** /framework_items | 



## FrameworkItemsFrameworkItemIdDelete

> FrameworkItems FrameworkItemsFrameworkItemIdDelete(ctx, frameworkItemId).Execute()



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
	frameworkItemId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkItemsAPI.FrameworkItemsFrameworkItemIdDelete(context.Background(), frameworkItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkItemsAPI.FrameworkItemsFrameworkItemIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkItemsFrameworkItemIdDelete`: FrameworkItems
	fmt.Fprintf(os.Stdout, "Response from `FrameworkItemsAPI.FrameworkItemsFrameworkItemIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**frameworkItemId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkItemsFrameworkItemIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FrameworkItems**](FrameworkItems.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkItemsFrameworkItemIdGet

> FrameworkItems FrameworkItemsFrameworkItemIdGet(ctx, frameworkItemId).Include(include).Execute()



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
	frameworkItemId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkItemsAPI.FrameworkItemsFrameworkItemIdGet(context.Background(), frameworkItemId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkItemsAPI.FrameworkItemsFrameworkItemIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkItemsFrameworkItemIdGet`: FrameworkItems
	fmt.Fprintf(os.Stdout, "Response from `FrameworkItemsAPI.FrameworkItemsFrameworkItemIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**frameworkItemId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkItemsFrameworkItemIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**FrameworkItems**](FrameworkItems.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkItemsFrameworkItemIdPut

> FrameworkItems FrameworkItemsFrameworkItemIdPut(ctx, frameworkItemId).FrameworkItemsPut(frameworkItemsPut).Execute()



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
	frameworkItemId := int64(789) // int64 | Model id
	frameworkItemsPut := *openapiclient.NewFrameworkItemsPut() // FrameworkItemsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkItemsAPI.FrameworkItemsFrameworkItemIdPut(context.Background(), frameworkItemId).FrameworkItemsPut(frameworkItemsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkItemsAPI.FrameworkItemsFrameworkItemIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkItemsFrameworkItemIdPut`: FrameworkItems
	fmt.Fprintf(os.Stdout, "Response from `FrameworkItemsAPI.FrameworkItemsFrameworkItemIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**frameworkItemId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkItemsFrameworkItemIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **frameworkItemsPut** | [**FrameworkItemsPut**](FrameworkItemsPut.md) |  | 

### Return type

[**FrameworkItems**](FrameworkItems.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkItemsGet

> FrameworkItemsGet200Response FrameworkItemsGet(ctx).FrameworkId(frameworkId).Execute()



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
	frameworkId := int32(56) // int32 | id of the framework to find (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkItemsAPI.FrameworkItemsGet(context.Background()).FrameworkId(frameworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkItemsAPI.FrameworkItemsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkItemsGet`: FrameworkItemsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FrameworkItemsAPI.FrameworkItemsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkItemsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **frameworkId** | **int32** | id of the framework to find | 

### Return type

[**FrameworkItemsGet200Response**](FrameworkItemsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkItemsPost

> FrameworkItemsGet200Response FrameworkItemsPost(ctx).FrameworkItemsPostRequest(frameworkItemsPostRequest).Execute()



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
	frameworkItemsPostRequest := *openapiclient.NewFrameworkItemsPostRequest() // FrameworkItemsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkItemsAPI.FrameworkItemsPost(context.Background()).FrameworkItemsPostRequest(frameworkItemsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkItemsAPI.FrameworkItemsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkItemsPost`: FrameworkItemsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FrameworkItemsAPI.FrameworkItemsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkItemsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **frameworkItemsPostRequest** | [**FrameworkItemsPostRequest**](FrameworkItemsPostRequest.md) |  | 

### Return type

[**FrameworkItemsGet200Response**](FrameworkItemsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

