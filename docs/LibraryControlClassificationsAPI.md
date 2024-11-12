# \LibraryControlClassificationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LibraryControlClassificationsGet**](LibraryControlClassificationsAPI.md#LibraryControlClassificationsGet) | **Get** /library_control_classifications | 
[**LibraryControlClassificationsLibraryControlClassificationIdDelete**](LibraryControlClassificationsAPI.md#LibraryControlClassificationsLibraryControlClassificationIdDelete) | **Delete** /library_control_classifications/{library_control_classification_id} | 
[**LibraryControlClassificationsLibraryControlClassificationIdGet**](LibraryControlClassificationsAPI.md#LibraryControlClassificationsLibraryControlClassificationIdGet) | **Get** /library_control_classifications/{library_control_classification_id} | 
[**LibraryControlClassificationsLibraryControlClassificationIdPut**](LibraryControlClassificationsAPI.md#LibraryControlClassificationsLibraryControlClassificationIdPut) | **Put** /library_control_classifications/{library_control_classification_id} | 
[**LibraryControlClassificationsPost**](LibraryControlClassificationsAPI.md#LibraryControlClassificationsPost) | **Post** /library_control_classifications | 



## LibraryControlClassificationsGet

> LibraryControlClassificationsGet200Response LibraryControlClassificationsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.LibraryControlClassificationsAPI.LibraryControlClassificationsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryControlClassificationsAPI.LibraryControlClassificationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibraryControlClassificationsGet`: LibraryControlClassificationsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryControlClassificationsAPI.LibraryControlClassificationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLibraryControlClassificationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**LibraryControlClassificationsGet200Response**](LibraryControlClassificationsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibraryControlClassificationsLibraryControlClassificationIdDelete

> LibraryControlClassifications LibraryControlClassificationsLibraryControlClassificationIdDelete(ctx, libraryControlClassificationId).Execute()



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
	libraryControlClassificationId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryControlClassificationsAPI.LibraryControlClassificationsLibraryControlClassificationIdDelete(context.Background(), libraryControlClassificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryControlClassificationsAPI.LibraryControlClassificationsLibraryControlClassificationIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibraryControlClassificationsLibraryControlClassificationIdDelete`: LibraryControlClassifications
	fmt.Fprintf(os.Stdout, "Response from `LibraryControlClassificationsAPI.LibraryControlClassificationsLibraryControlClassificationIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryControlClassificationId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibraryControlClassificationsLibraryControlClassificationIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LibraryControlClassifications**](LibraryControlClassifications.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibraryControlClassificationsLibraryControlClassificationIdGet

> LibraryControlClassifications LibraryControlClassificationsLibraryControlClassificationIdGet(ctx, libraryControlClassificationId).Include(include).Execute()



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
	libraryControlClassificationId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryControlClassificationsAPI.LibraryControlClassificationsLibraryControlClassificationIdGet(context.Background(), libraryControlClassificationId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryControlClassificationsAPI.LibraryControlClassificationsLibraryControlClassificationIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibraryControlClassificationsLibraryControlClassificationIdGet`: LibraryControlClassifications
	fmt.Fprintf(os.Stdout, "Response from `LibraryControlClassificationsAPI.LibraryControlClassificationsLibraryControlClassificationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryControlClassificationId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibraryControlClassificationsLibraryControlClassificationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**LibraryControlClassifications**](LibraryControlClassifications.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibraryControlClassificationsLibraryControlClassificationIdPut

> LibraryControlClassifications LibraryControlClassificationsLibraryControlClassificationIdPut(ctx, libraryControlClassificationId).LibraryControlClassificationsPut(libraryControlClassificationsPut).Execute()



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
	libraryControlClassificationId := int64(789) // int64 | Model id
	libraryControlClassificationsPut := *openapiclient.NewLibraryControlClassificationsPut() // LibraryControlClassificationsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryControlClassificationsAPI.LibraryControlClassificationsLibraryControlClassificationIdPut(context.Background(), libraryControlClassificationId).LibraryControlClassificationsPut(libraryControlClassificationsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryControlClassificationsAPI.LibraryControlClassificationsLibraryControlClassificationIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibraryControlClassificationsLibraryControlClassificationIdPut`: LibraryControlClassifications
	fmt.Fprintf(os.Stdout, "Response from `LibraryControlClassificationsAPI.LibraryControlClassificationsLibraryControlClassificationIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryControlClassificationId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLibraryControlClassificationsLibraryControlClassificationIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **libraryControlClassificationsPut** | [**LibraryControlClassificationsPut**](LibraryControlClassificationsPut.md) |  | 

### Return type

[**LibraryControlClassifications**](LibraryControlClassifications.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LibraryControlClassificationsPost

> LibraryControlClassificationsGet200Response LibraryControlClassificationsPost(ctx).LibraryControlClassificationsPostRequest(libraryControlClassificationsPostRequest).Execute()



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
	libraryControlClassificationsPostRequest := *openapiclient.NewLibraryControlClassificationsPostRequest() // LibraryControlClassificationsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryControlClassificationsAPI.LibraryControlClassificationsPost(context.Background()).LibraryControlClassificationsPostRequest(libraryControlClassificationsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryControlClassificationsAPI.LibraryControlClassificationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibraryControlClassificationsPost`: LibraryControlClassificationsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryControlClassificationsAPI.LibraryControlClassificationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLibraryControlClassificationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **libraryControlClassificationsPostRequest** | [**LibraryControlClassificationsPostRequest**](LibraryControlClassificationsPostRequest.md) |  | 

### Return type

[**LibraryControlClassificationsGet200Response**](LibraryControlClassificationsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

