# \GlobalAttributesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlobalAttributesGet**](GlobalAttributesAPI.md#GlobalAttributesGet) | **Get** /global_attributes | 
[**GlobalAttributesGlobalAttributeIdDelete**](GlobalAttributesAPI.md#GlobalAttributesGlobalAttributeIdDelete) | **Delete** /global_attributes/{global_attribute_id} | 
[**GlobalAttributesGlobalAttributeIdGet**](GlobalAttributesAPI.md#GlobalAttributesGlobalAttributeIdGet) | **Get** /global_attributes/{global_attribute_id} | 
[**GlobalAttributesGlobalAttributeIdPut**](GlobalAttributesAPI.md#GlobalAttributesGlobalAttributeIdPut) | **Put** /global_attributes/{global_attribute_id} | 
[**GlobalAttributesPost**](GlobalAttributesAPI.md#GlobalAttributesPost) | **Post** /global_attributes | 



## GlobalAttributesGet

> GlobalAttributesGet200Response GlobalAttributesGet(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalAttributesAPI.GlobalAttributesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAttributesAPI.GlobalAttributesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalAttributesGet`: GlobalAttributesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalAttributesAPI.GlobalAttributesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalAttributesGetRequest struct via the builder pattern


### Return type

[**GlobalAttributesGet200Response**](GlobalAttributesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlobalAttributesGlobalAttributeIdDelete

> GlobalAttributes GlobalAttributesGlobalAttributeIdDelete(ctx, globalAttributeId).Execute()



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
	globalAttributeId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalAttributesAPI.GlobalAttributesGlobalAttributeIdDelete(context.Background(), globalAttributeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAttributesAPI.GlobalAttributesGlobalAttributeIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalAttributesGlobalAttributeIdDelete`: GlobalAttributes
	fmt.Fprintf(os.Stdout, "Response from `GlobalAttributesAPI.GlobalAttributesGlobalAttributeIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**globalAttributeId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalAttributesGlobalAttributeIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GlobalAttributes**](GlobalAttributes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlobalAttributesGlobalAttributeIdGet

> GlobalAttributes GlobalAttributesGlobalAttributeIdGet(ctx, globalAttributeId).Execute()



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
	globalAttributeId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalAttributesAPI.GlobalAttributesGlobalAttributeIdGet(context.Background(), globalAttributeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAttributesAPI.GlobalAttributesGlobalAttributeIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalAttributesGlobalAttributeIdGet`: GlobalAttributes
	fmt.Fprintf(os.Stdout, "Response from `GlobalAttributesAPI.GlobalAttributesGlobalAttributeIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**globalAttributeId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalAttributesGlobalAttributeIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GlobalAttributes**](GlobalAttributes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlobalAttributesGlobalAttributeIdPut

> GlobalAttributes GlobalAttributesGlobalAttributeIdPut(ctx, globalAttributeId).GlobalAttributesPut(globalAttributesPut).Execute()



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
	globalAttributeId := int64(789) // int64 | Model id
	globalAttributesPut := *openapiclient.NewGlobalAttributesPut() // GlobalAttributesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalAttributesAPI.GlobalAttributesGlobalAttributeIdPut(context.Background(), globalAttributeId).GlobalAttributesPut(globalAttributesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAttributesAPI.GlobalAttributesGlobalAttributeIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalAttributesGlobalAttributeIdPut`: GlobalAttributes
	fmt.Fprintf(os.Stdout, "Response from `GlobalAttributesAPI.GlobalAttributesGlobalAttributeIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**globalAttributeId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalAttributesGlobalAttributeIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **globalAttributesPut** | [**GlobalAttributesPut**](GlobalAttributesPut.md) |  | 

### Return type

[**GlobalAttributes**](GlobalAttributes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlobalAttributesPost

> GlobalAttributesGet200Response GlobalAttributesPost(ctx).GlobalAttributesPostRequest(globalAttributesPostRequest).Execute()



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
	globalAttributesPostRequest := *openapiclient.NewGlobalAttributesPostRequest() // GlobalAttributesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalAttributesAPI.GlobalAttributesPost(context.Background()).GlobalAttributesPostRequest(globalAttributesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAttributesAPI.GlobalAttributesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalAttributesPost`: GlobalAttributesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalAttributesAPI.GlobalAttributesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlobalAttributesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **globalAttributesPostRequest** | [**GlobalAttributesPostRequest**](GlobalAttributesPostRequest.md) |  | 

### Return type

[**GlobalAttributesGet200Response**](GlobalAttributesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

