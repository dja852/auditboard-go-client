# \LibraryControlTypesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LibraryControlTypesGet**](LibraryControlTypesAPI.md#LibraryControlTypesGet) | **Get** /library_control_types | 
[**LibraryControlTypesLibraryControlTypeIdDelete**](LibraryControlTypesAPI.md#LibraryControlTypesLibraryControlTypeIdDelete) | **Delete** /library_control_types/{library_control_type_id} | 
[**LibraryControlTypesLibraryControlTypeIdGet**](LibraryControlTypesAPI.md#LibraryControlTypesLibraryControlTypeIdGet) | **Get** /library_control_types/{library_control_type_id} | 
[**LibraryControlTypesLibraryControlTypeIdPut**](LibraryControlTypesAPI.md#LibraryControlTypesLibraryControlTypeIdPut) | **Put** /library_control_types/{library_control_type_id} | 
[**LibraryControlTypesPost**](LibraryControlTypesAPI.md#LibraryControlTypesPost) | **Post** /library_control_types | 



## LibraryControlTypesGet

> LibraryControlTypesGet200Response LibraryControlTypesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.LibraryControlTypesAPI.LibraryControlTypesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryControlTypesAPI.LibraryControlTypesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibraryControlTypesGet`: LibraryControlTypesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryControlTypesAPI.LibraryControlTypesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLibraryControlTypesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**LibraryControlTypesGet200Response**](LibraryControlTypesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibraryControlTypesLibraryControlTypeIdDelete

> LibraryControlTypes LibraryControlTypesLibraryControlTypeIdDelete(ctx, libraryControlTypeId).Execute()



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
	libraryControlTypeId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryControlTypesAPI.LibraryControlTypesLibraryControlTypeIdDelete(context.Background(), libraryControlTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryControlTypesAPI.LibraryControlTypesLibraryControlTypeIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibraryControlTypesLibraryControlTypeIdDelete`: LibraryControlTypes
	fmt.Fprintf(os.Stdout, "Response from `LibraryControlTypesAPI.LibraryControlTypesLibraryControlTypeIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryControlTypeId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibraryControlTypesLibraryControlTypeIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LibraryControlTypes**](LibraryControlTypes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibraryControlTypesLibraryControlTypeIdGet

> LibraryControlTypes LibraryControlTypesLibraryControlTypeIdGet(ctx, libraryControlTypeId).Include(include).Execute()



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
	libraryControlTypeId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryControlTypesAPI.LibraryControlTypesLibraryControlTypeIdGet(context.Background(), libraryControlTypeId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryControlTypesAPI.LibraryControlTypesLibraryControlTypeIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibraryControlTypesLibraryControlTypeIdGet`: LibraryControlTypes
	fmt.Fprintf(os.Stdout, "Response from `LibraryControlTypesAPI.LibraryControlTypesLibraryControlTypeIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryControlTypeId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibraryControlTypesLibraryControlTypeIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**LibraryControlTypes**](LibraryControlTypes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibraryControlTypesLibraryControlTypeIdPut

> LibraryControlTypes LibraryControlTypesLibraryControlTypeIdPut(ctx, libraryControlTypeId).LibraryControlTypesPut(libraryControlTypesPut).Execute()



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
	libraryControlTypeId := int64(789) // int64 | Model id
	libraryControlTypesPut := *openapiclient.NewLibraryControlTypesPut() // LibraryControlTypesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryControlTypesAPI.LibraryControlTypesLibraryControlTypeIdPut(context.Background(), libraryControlTypeId).LibraryControlTypesPut(libraryControlTypesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryControlTypesAPI.LibraryControlTypesLibraryControlTypeIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibraryControlTypesLibraryControlTypeIdPut`: LibraryControlTypes
	fmt.Fprintf(os.Stdout, "Response from `LibraryControlTypesAPI.LibraryControlTypesLibraryControlTypeIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryControlTypeId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibraryControlTypesLibraryControlTypeIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **libraryControlTypesPut** | [**LibraryControlTypesPut**](LibraryControlTypesPut.md) |  | 

### Return type

[**LibraryControlTypes**](LibraryControlTypes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibraryControlTypesPost

> LibraryControlTypesGet200Response LibraryControlTypesPost(ctx).LibraryControlTypesPostRequest(libraryControlTypesPostRequest).Execute()



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
	libraryControlTypesPostRequest := *openapiclient.NewLibraryControlTypesPostRequest() // LibraryControlTypesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryControlTypesAPI.LibraryControlTypesPost(context.Background()).LibraryControlTypesPostRequest(libraryControlTypesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryControlTypesAPI.LibraryControlTypesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibraryControlTypesPost`: LibraryControlTypesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryControlTypesAPI.LibraryControlTypesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLibraryControlTypesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **libraryControlTypesPostRequest** | [**LibraryControlTypesPostRequest**](LibraryControlTypesPostRequest.md) |  | 

### Return type

[**LibraryControlTypesGet200Response**](LibraryControlTypesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

