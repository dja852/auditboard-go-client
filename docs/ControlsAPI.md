# \ControlsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ControlsControlIdDelete**](ControlsAPI.md#ControlsControlIdDelete) | **Delete** /controls/{control_id} | 
[**ControlsControlIdGet**](ControlsAPI.md#ControlsControlIdGet) | **Get** /controls/{control_id} | 
[**ControlsControlIdPut**](ControlsAPI.md#ControlsControlIdPut) | **Put** /controls/{control_id} | 
[**ControlsGet**](ControlsAPI.md#ControlsGet) | **Get** /controls | 
[**ControlsPost**](ControlsAPI.md#ControlsPost) | **Post** /controls | 



## ControlsControlIdDelete

> Controls ControlsControlIdDelete(ctx, controlId).Execute()



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
	controlId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControlsAPI.ControlsControlIdDelete(context.Background(), controlId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControlsAPI.ControlsControlIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlsControlIdDelete`: Controls
	fmt.Fprintf(os.Stdout, "Response from `ControlsAPI.ControlsControlIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**controlId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiControlsControlIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Controls**](Controls.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ControlsControlIdGet

> Controls ControlsControlIdGet(ctx, controlId).Include(include).Execute()



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
	controlId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControlsAPI.ControlsControlIdGet(context.Background(), controlId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControlsAPI.ControlsControlIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlsControlIdGet`: Controls
	fmt.Fprintf(os.Stdout, "Response from `ControlsAPI.ControlsControlIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**controlId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiControlsControlIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**Controls**](Controls.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ControlsControlIdPut

> Controls ControlsControlIdPut(ctx, controlId).ControlsPut(controlsPut).Execute()



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
	controlId := int64(789) // int64 | Model id
	controlsPut := *openapiclient.NewControlsPut() // ControlsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControlsAPI.ControlsControlIdPut(context.Background(), controlId).ControlsPut(controlsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControlsAPI.ControlsControlIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlsControlIdPut`: Controls
	fmt.Fprintf(os.Stdout, "Response from `ControlsAPI.ControlsControlIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**controlId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiControlsControlIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **controlsPut** | [**ControlsPut**](ControlsPut.md) |  | 

### Return type

[**Controls**](Controls.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ControlsGet

> ControlsGet200Response ControlsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.ControlsAPI.ControlsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControlsAPI.ControlsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlsGet`: ControlsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ControlsAPI.ControlsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiControlsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**ControlsGet200Response**](ControlsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ControlsPost

> ControlsGet200Response ControlsPost(ctx).ControlsPostRequest(controlsPostRequest).Execute()



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
	controlsPostRequest := *openapiclient.NewControlsPostRequest() // ControlsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControlsAPI.ControlsPost(context.Background()).ControlsPostRequest(controlsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControlsAPI.ControlsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlsPost`: ControlsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ControlsAPI.ControlsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiControlsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controlsPostRequest** | [**ControlsPostRequest**](ControlsPostRequest.md) |  | 

### Return type

[**ControlsGet200Response**](ControlsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

