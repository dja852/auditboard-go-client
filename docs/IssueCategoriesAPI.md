# \IssueCategoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IssueCategoriesGet**](IssueCategoriesAPI.md#IssueCategoriesGet) | **Get** /issue_categories | 
[**IssueCategoriesIssueCategoryIdDelete**](IssueCategoriesAPI.md#IssueCategoriesIssueCategoryIdDelete) | **Delete** /issue_categories/{issue_category_id} | 
[**IssueCategoriesIssueCategoryIdGet**](IssueCategoriesAPI.md#IssueCategoriesIssueCategoryIdGet) | **Get** /issue_categories/{issue_category_id} | 
[**IssueCategoriesIssueCategoryIdPut**](IssueCategoriesAPI.md#IssueCategoriesIssueCategoryIdPut) | **Put** /issue_categories/{issue_category_id} | 
[**IssueCategoriesPost**](IssueCategoriesAPI.md#IssueCategoriesPost) | **Post** /issue_categories | 



## IssueCategoriesGet

> IssueCategoriesGet200Response IssueCategoriesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.IssueCategoriesAPI.IssueCategoriesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCategoriesAPI.IssueCategoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueCategoriesGet`: IssueCategoriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `IssueCategoriesAPI.IssueCategoriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssueCategoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**IssueCategoriesGet200Response**](IssueCategoriesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueCategoriesIssueCategoryIdDelete

> IssueCategories IssueCategoriesIssueCategoryIdDelete(ctx, issueCategoryId).Execute()



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
	issueCategoryId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCategoriesAPI.IssueCategoriesIssueCategoryIdDelete(context.Background(), issueCategoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCategoriesAPI.IssueCategoriesIssueCategoryIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueCategoriesIssueCategoryIdDelete`: IssueCategories
	fmt.Fprintf(os.Stdout, "Response from `IssueCategoriesAPI.IssueCategoriesIssueCategoryIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueCategoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueCategoriesIssueCategoryIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IssueCategories**](IssueCategories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueCategoriesIssueCategoryIdGet

> IssueCategories IssueCategoriesIssueCategoryIdGet(ctx, issueCategoryId).Include(include).Execute()



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
	issueCategoryId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCategoriesAPI.IssueCategoriesIssueCategoryIdGet(context.Background(), issueCategoryId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCategoriesAPI.IssueCategoriesIssueCategoryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueCategoriesIssueCategoryIdGet`: IssueCategories
	fmt.Fprintf(os.Stdout, "Response from `IssueCategoriesAPI.IssueCategoriesIssueCategoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueCategoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueCategoriesIssueCategoryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**IssueCategories**](IssueCategories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueCategoriesIssueCategoryIdPut

> IssueCategories IssueCategoriesIssueCategoryIdPut(ctx, issueCategoryId).IssueCategoriesPut(issueCategoriesPut).Execute()



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
	issueCategoryId := int64(789) // int64 | Model id
	issueCategoriesPut := *openapiclient.NewIssueCategoriesPut() // IssueCategoriesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCategoriesAPI.IssueCategoriesIssueCategoryIdPut(context.Background(), issueCategoryId).IssueCategoriesPut(issueCategoriesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCategoriesAPI.IssueCategoriesIssueCategoryIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueCategoriesIssueCategoryIdPut`: IssueCategories
	fmt.Fprintf(os.Stdout, "Response from `IssueCategoriesAPI.IssueCategoriesIssueCategoryIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueCategoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueCategoriesIssueCategoryIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **issueCategoriesPut** | [**IssueCategoriesPut**](IssueCategoriesPut.md) |  | 

### Return type

[**IssueCategories**](IssueCategories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueCategoriesPost

> IssueCategoriesGet200Response IssueCategoriesPost(ctx).IssueCategoriesPostRequest(issueCategoriesPostRequest).Execute()



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
	issueCategoriesPostRequest := *openapiclient.NewIssueCategoriesPostRequest() // IssueCategoriesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCategoriesAPI.IssueCategoriesPost(context.Background()).IssueCategoriesPostRequest(issueCategoriesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCategoriesAPI.IssueCategoriesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueCategoriesPost`: IssueCategoriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `IssueCategoriesAPI.IssueCategoriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssueCategoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueCategoriesPostRequest** | [**IssueCategoriesPostRequest**](IssueCategoriesPostRequest.md) |  | 

### Return type

[**IssueCategoriesGet200Response**](IssueCategoriesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

