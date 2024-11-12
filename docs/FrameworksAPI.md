# \FrameworksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FrameworksFrameworkIdDelete**](FrameworksAPI.md#FrameworksFrameworkIdDelete) | **Delete** /frameworks/{framework_id} | 
[**FrameworksFrameworkIdGet**](FrameworksAPI.md#FrameworksFrameworkIdGet) | **Get** /frameworks/{framework_id} | 
[**FrameworksFrameworkIdPut**](FrameworksAPI.md#FrameworksFrameworkIdPut) | **Put** /frameworks/{framework_id} | 
[**FrameworksGet**](FrameworksAPI.md#FrameworksGet) | **Get** /frameworks | 
[**FrameworksPost**](FrameworksAPI.md#FrameworksPost) | **Post** /frameworks | 



## FrameworksFrameworkIdDelete

> Frameworks FrameworksFrameworkIdDelete(ctx, frameworkId).Execute()



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
	frameworkId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworksAPI.FrameworksFrameworkIdDelete(context.Background(), frameworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworksAPI.FrameworksFrameworkIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworksFrameworkIdDelete`: Frameworks
	fmt.Fprintf(os.Stdout, "Response from `FrameworksAPI.FrameworksFrameworkIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**frameworkId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworksFrameworkIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Frameworks**](Frameworks.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworksFrameworkIdGet

> Frameworks FrameworksFrameworkIdGet(ctx, frameworkId).Include(include).Execute()



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
	frameworkId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworksAPI.FrameworksFrameworkIdGet(context.Background(), frameworkId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworksAPI.FrameworksFrameworkIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworksFrameworkIdGet`: Frameworks
	fmt.Fprintf(os.Stdout, "Response from `FrameworksAPI.FrameworksFrameworkIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**frameworkId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworksFrameworkIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**Frameworks**](Frameworks.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworksFrameworkIdPut

> Frameworks FrameworksFrameworkIdPut(ctx, frameworkId).FrameworksPut(frameworksPut).Execute()



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
	frameworkId := int64(789) // int64 | Model id
	frameworksPut := *openapiclient.NewFrameworksPut() // FrameworksPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworksAPI.FrameworksFrameworkIdPut(context.Background(), frameworkId).FrameworksPut(frameworksPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworksAPI.FrameworksFrameworkIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworksFrameworkIdPut`: Frameworks
	fmt.Fprintf(os.Stdout, "Response from `FrameworksAPI.FrameworksFrameworkIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**frameworkId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworksFrameworkIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **frameworksPut** | [**FrameworksPut**](FrameworksPut.md) |  | 

### Return type

[**Frameworks**](Frameworks.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworksGet

> FrameworksGet200Response FrameworksGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.FrameworksAPI.FrameworksGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworksAPI.FrameworksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworksGet`: FrameworksGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FrameworksAPI.FrameworksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFrameworksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**FrameworksGet200Response**](FrameworksGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworksPost

> FrameworksGet200Response FrameworksPost(ctx).FrameworksPostRequest(frameworksPostRequest).Execute()



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
	frameworksPostRequest := *openapiclient.NewFrameworksPostRequest() // FrameworksPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworksAPI.FrameworksPost(context.Background()).FrameworksPostRequest(frameworksPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworksAPI.FrameworksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworksPost`: FrameworksGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FrameworksAPI.FrameworksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFrameworksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **frameworksPostRequest** | [**FrameworksPostRequest**](FrameworksPostRequest.md) |  | 

### Return type

[**FrameworksGet200Response**](FrameworksGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

