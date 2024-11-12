# \ControlsSegmentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ControlsSegmentsControlsSegmentIdDelete**](ControlsSegmentsAPI.md#ControlsSegmentsControlsSegmentIdDelete) | **Delete** /controls_segments/{controls_segment_id} | 
[**ControlsSegmentsControlsSegmentIdGet**](ControlsSegmentsAPI.md#ControlsSegmentsControlsSegmentIdGet) | **Get** /controls_segments/{controls_segment_id} | 
[**ControlsSegmentsControlsSegmentIdPut**](ControlsSegmentsAPI.md#ControlsSegmentsControlsSegmentIdPut) | **Put** /controls_segments/{controls_segment_id} | 
[**ControlsSegmentsGet**](ControlsSegmentsAPI.md#ControlsSegmentsGet) | **Get** /controls_segments | 
[**ControlsSegmentsPost**](ControlsSegmentsAPI.md#ControlsSegmentsPost) | **Post** /controls_segments | 



## ControlsSegmentsControlsSegmentIdDelete

> ControlsSegments ControlsSegmentsControlsSegmentIdDelete(ctx, controlsSegmentId).Execute()



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
	controlsSegmentId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControlsSegmentsAPI.ControlsSegmentsControlsSegmentIdDelete(context.Background(), controlsSegmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControlsSegmentsAPI.ControlsSegmentsControlsSegmentIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlsSegmentsControlsSegmentIdDelete`: ControlsSegments
	fmt.Fprintf(os.Stdout, "Response from `ControlsSegmentsAPI.ControlsSegmentsControlsSegmentIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**controlsSegmentId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiControlsSegmentsControlsSegmentIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ControlsSegments**](ControlsSegments.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ControlsSegmentsControlsSegmentIdGet

> ControlsSegments ControlsSegmentsControlsSegmentIdGet(ctx, controlsSegmentId).Include(include).Execute()



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
	controlsSegmentId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControlsSegmentsAPI.ControlsSegmentsControlsSegmentIdGet(context.Background(), controlsSegmentId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControlsSegmentsAPI.ControlsSegmentsControlsSegmentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlsSegmentsControlsSegmentIdGet`: ControlsSegments
	fmt.Fprintf(os.Stdout, "Response from `ControlsSegmentsAPI.ControlsSegmentsControlsSegmentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**controlsSegmentId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiControlsSegmentsControlsSegmentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**ControlsSegments**](ControlsSegments.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ControlsSegmentsControlsSegmentIdPut

> ControlsSegments ControlsSegmentsControlsSegmentIdPut(ctx, controlsSegmentId).ControlsSegmentsPut(controlsSegmentsPut).Execute()



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
	controlsSegmentId := int64(789) // int64 | Model id
	controlsSegmentsPut := *openapiclient.NewControlsSegmentsPut() // ControlsSegmentsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControlsSegmentsAPI.ControlsSegmentsControlsSegmentIdPut(context.Background(), controlsSegmentId).ControlsSegmentsPut(controlsSegmentsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControlsSegmentsAPI.ControlsSegmentsControlsSegmentIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlsSegmentsControlsSegmentIdPut`: ControlsSegments
	fmt.Fprintf(os.Stdout, "Response from `ControlsSegmentsAPI.ControlsSegmentsControlsSegmentIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**controlsSegmentId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiControlsSegmentsControlsSegmentIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **controlsSegmentsPut** | [**ControlsSegmentsPut**](ControlsSegmentsPut.md) |  | 

### Return type

[**ControlsSegments**](ControlsSegments.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ControlsSegmentsGet

> ControlsSegmentsGet200Response ControlsSegmentsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.ControlsSegmentsAPI.ControlsSegmentsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControlsSegmentsAPI.ControlsSegmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlsSegmentsGet`: ControlsSegmentsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ControlsSegmentsAPI.ControlsSegmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiControlsSegmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**ControlsSegmentsGet200Response**](ControlsSegmentsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ControlsSegmentsPost

> ControlsSegmentsGet200Response ControlsSegmentsPost(ctx).ControlsSegmentsPostRequest(controlsSegmentsPostRequest).Execute()



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
	controlsSegmentsPostRequest := *openapiclient.NewControlsSegmentsPostRequest() // ControlsSegmentsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControlsSegmentsAPI.ControlsSegmentsPost(context.Background()).ControlsSegmentsPostRequest(controlsSegmentsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControlsSegmentsAPI.ControlsSegmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlsSegmentsPost`: ControlsSegmentsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ControlsSegmentsAPI.ControlsSegmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiControlsSegmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controlsSegmentsPostRequest** | [**ControlsSegmentsPostRequest**](ControlsSegmentsPostRequest.md) |  | 

### Return type

[**ControlsSegmentsGet200Response**](ControlsSegmentsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

