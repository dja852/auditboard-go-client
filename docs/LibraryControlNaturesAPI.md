# \LibraryControlNaturesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LibraryControlNaturesGet**](LibraryControlNaturesAPI.md#LibraryControlNaturesGet) | **Get** /library_control_natures | 
[**LibraryControlNaturesLibraryControlNatureIdDelete**](LibraryControlNaturesAPI.md#LibraryControlNaturesLibraryControlNatureIdDelete) | **Delete** /library_control_natures/{library_control_nature_id} | 
[**LibraryControlNaturesLibraryControlNatureIdGet**](LibraryControlNaturesAPI.md#LibraryControlNaturesLibraryControlNatureIdGet) | **Get** /library_control_natures/{library_control_nature_id} | 
[**LibraryControlNaturesLibraryControlNatureIdPut**](LibraryControlNaturesAPI.md#LibraryControlNaturesLibraryControlNatureIdPut) | **Put** /library_control_natures/{library_control_nature_id} | 
[**LibraryControlNaturesPost**](LibraryControlNaturesAPI.md#LibraryControlNaturesPost) | **Post** /library_control_natures | 



## LibraryControlNaturesGet

> LibraryControlNaturesGet200Response LibraryControlNaturesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.LibraryControlNaturesAPI.LibraryControlNaturesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryControlNaturesAPI.LibraryControlNaturesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibraryControlNaturesGet`: LibraryControlNaturesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryControlNaturesAPI.LibraryControlNaturesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLibraryControlNaturesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**LibraryControlNaturesGet200Response**](LibraryControlNaturesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibraryControlNaturesLibraryControlNatureIdDelete

> LibraryControlNatures LibraryControlNaturesLibraryControlNatureIdDelete(ctx, libraryControlNatureId).Execute()



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
	libraryControlNatureId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryControlNaturesAPI.LibraryControlNaturesLibraryControlNatureIdDelete(context.Background(), libraryControlNatureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryControlNaturesAPI.LibraryControlNaturesLibraryControlNatureIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibraryControlNaturesLibraryControlNatureIdDelete`: LibraryControlNatures
	fmt.Fprintf(os.Stdout, "Response from `LibraryControlNaturesAPI.LibraryControlNaturesLibraryControlNatureIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryControlNatureId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibraryControlNaturesLibraryControlNatureIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LibraryControlNatures**](LibraryControlNatures.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibraryControlNaturesLibraryControlNatureIdGet

> LibraryControlNatures LibraryControlNaturesLibraryControlNatureIdGet(ctx, libraryControlNatureId).Include(include).Execute()



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
	libraryControlNatureId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryControlNaturesAPI.LibraryControlNaturesLibraryControlNatureIdGet(context.Background(), libraryControlNatureId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryControlNaturesAPI.LibraryControlNaturesLibraryControlNatureIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibraryControlNaturesLibraryControlNatureIdGet`: LibraryControlNatures
	fmt.Fprintf(os.Stdout, "Response from `LibraryControlNaturesAPI.LibraryControlNaturesLibraryControlNatureIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryControlNatureId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibraryControlNaturesLibraryControlNatureIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**LibraryControlNatures**](LibraryControlNatures.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibraryControlNaturesLibraryControlNatureIdPut

> LibraryControlNatures LibraryControlNaturesLibraryControlNatureIdPut(ctx, libraryControlNatureId).LibraryControlNaturesPut(libraryControlNaturesPut).Execute()



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
	libraryControlNatureId := int64(789) // int64 | Model id
	libraryControlNaturesPut := *openapiclient.NewLibraryControlNaturesPut() // LibraryControlNaturesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryControlNaturesAPI.LibraryControlNaturesLibraryControlNatureIdPut(context.Background(), libraryControlNatureId).LibraryControlNaturesPut(libraryControlNaturesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryControlNaturesAPI.LibraryControlNaturesLibraryControlNatureIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibraryControlNaturesLibraryControlNatureIdPut`: LibraryControlNatures
	fmt.Fprintf(os.Stdout, "Response from `LibraryControlNaturesAPI.LibraryControlNaturesLibraryControlNatureIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryControlNatureId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibraryControlNaturesLibraryControlNatureIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **libraryControlNaturesPut** | [**LibraryControlNaturesPut**](LibraryControlNaturesPut.md) |  | 

### Return type

[**LibraryControlNatures**](LibraryControlNatures.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibraryControlNaturesPost

> LibraryControlNaturesGet200Response LibraryControlNaturesPost(ctx).LibraryControlNaturesPostRequest(libraryControlNaturesPostRequest).Execute()



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
	libraryControlNaturesPostRequest := *openapiclient.NewLibraryControlNaturesPostRequest() // LibraryControlNaturesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryControlNaturesAPI.LibraryControlNaturesPost(context.Background()).LibraryControlNaturesPostRequest(libraryControlNaturesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryControlNaturesAPI.LibraryControlNaturesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibraryControlNaturesPost`: LibraryControlNaturesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryControlNaturesAPI.LibraryControlNaturesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLibraryControlNaturesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **libraryControlNaturesPostRequest** | [**LibraryControlNaturesPostRequest**](LibraryControlNaturesPostRequest.md) |  | 

### Return type

[**LibraryControlNaturesGet200Response**](LibraryControlNaturesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

