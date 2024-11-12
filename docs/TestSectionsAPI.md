# \TestSectionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestSectionsGet**](TestSectionsAPI.md#TestSectionsGet) | **Get** /test_sections | 
[**TestSectionsPost**](TestSectionsAPI.md#TestSectionsPost) | **Post** /test_sections | 
[**TestSectionsTestSectionIdDelete**](TestSectionsAPI.md#TestSectionsTestSectionIdDelete) | **Delete** /test_sections/{test_section_id} | 
[**TestSectionsTestSectionIdGet**](TestSectionsAPI.md#TestSectionsTestSectionIdGet) | **Get** /test_sections/{test_section_id} | 
[**TestSectionsTestSectionIdPut**](TestSectionsAPI.md#TestSectionsTestSectionIdPut) | **Put** /test_sections/{test_section_id} | 



## TestSectionsGet

> TestSectionsGet200Response TestSectionsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.TestSectionsAPI.TestSectionsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestSectionsAPI.TestSectionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestSectionsGet`: TestSectionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TestSectionsAPI.TestSectionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestSectionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**TestSectionsGet200Response**](TestSectionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestSectionsPost

> TestSectionsGet200Response TestSectionsPost(ctx).TestSectionsPostRequest(testSectionsPostRequest).Execute()



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
	testSectionsPostRequest := *openapiclient.NewTestSectionsPostRequest() // TestSectionsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestSectionsAPI.TestSectionsPost(context.Background()).TestSectionsPostRequest(testSectionsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestSectionsAPI.TestSectionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestSectionsPost`: TestSectionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TestSectionsAPI.TestSectionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestSectionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testSectionsPostRequest** | [**TestSectionsPostRequest**](TestSectionsPostRequest.md) |  | 

### Return type

[**TestSectionsGet200Response**](TestSectionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestSectionsTestSectionIdDelete

> TestSections TestSectionsTestSectionIdDelete(ctx, testSectionId).Execute()



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
	testSectionId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestSectionsAPI.TestSectionsTestSectionIdDelete(context.Background(), testSectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestSectionsAPI.TestSectionsTestSectionIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestSectionsTestSectionIdDelete`: TestSections
	fmt.Fprintf(os.Stdout, "Response from `TestSectionsAPI.TestSectionsTestSectionIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testSectionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestSectionsTestSectionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestSections**](TestSections.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestSectionsTestSectionIdGet

> TestSections TestSectionsTestSectionIdGet(ctx, testSectionId).Include(include).Execute()



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
	testSectionId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestSectionsAPI.TestSectionsTestSectionIdGet(context.Background(), testSectionId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestSectionsAPI.TestSectionsTestSectionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestSectionsTestSectionIdGet`: TestSections
	fmt.Fprintf(os.Stdout, "Response from `TestSectionsAPI.TestSectionsTestSectionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testSectionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestSectionsTestSectionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**TestSections**](TestSections.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestSectionsTestSectionIdPut

> TestSections TestSectionsTestSectionIdPut(ctx, testSectionId).TestSectionsPut(testSectionsPut).Execute()



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
	testSectionId := int64(789) // int64 | Model id
	testSectionsPut := *openapiclient.NewTestSectionsPut() // TestSectionsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestSectionsAPI.TestSectionsTestSectionIdPut(context.Background(), testSectionId).TestSectionsPut(testSectionsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestSectionsAPI.TestSectionsTestSectionIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestSectionsTestSectionIdPut`: TestSections
	fmt.Fprintf(os.Stdout, "Response from `TestSectionsAPI.TestSectionsTestSectionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testSectionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestSectionsTestSectionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testSectionsPut** | [**TestSectionsPut**](TestSectionsPut.md) |  | 

### Return type

[**TestSections**](TestSections.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

