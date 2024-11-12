# \EsgMetricCategoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EsgMetricCategoriesEsgMetricCategoryIdDelete**](EsgMetricCategoriesAPI.md#EsgMetricCategoriesEsgMetricCategoryIdDelete) | **Delete** /esg_metric_categories/{esg_metric_category_id} | 
[**EsgMetricCategoriesEsgMetricCategoryIdGet**](EsgMetricCategoriesAPI.md#EsgMetricCategoriesEsgMetricCategoryIdGet) | **Get** /esg_metric_categories/{esg_metric_category_id} | 
[**EsgMetricCategoriesEsgMetricCategoryIdPut**](EsgMetricCategoriesAPI.md#EsgMetricCategoriesEsgMetricCategoryIdPut) | **Put** /esg_metric_categories/{esg_metric_category_id} | 
[**EsgMetricCategoriesGet**](EsgMetricCategoriesAPI.md#EsgMetricCategoriesGet) | **Get** /esg_metric_categories | 
[**EsgMetricCategoriesPost**](EsgMetricCategoriesAPI.md#EsgMetricCategoriesPost) | **Post** /esg_metric_categories | 



## EsgMetricCategoriesEsgMetricCategoryIdDelete

> EsgMetricCategories EsgMetricCategoriesEsgMetricCategoryIdDelete(ctx, esgMetricCategoryId).Execute()



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
	esgMetricCategoryId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricCategoriesAPI.EsgMetricCategoriesEsgMetricCategoryIdDelete(context.Background(), esgMetricCategoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricCategoriesAPI.EsgMetricCategoriesEsgMetricCategoryIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricCategoriesEsgMetricCategoryIdDelete`: EsgMetricCategories
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricCategoriesAPI.EsgMetricCategoriesEsgMetricCategoryIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricCategoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricCategoriesEsgMetricCategoryIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EsgMetricCategories**](EsgMetricCategories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricCategoriesEsgMetricCategoryIdGet

> EsgMetricCategories EsgMetricCategoriesEsgMetricCategoryIdGet(ctx, esgMetricCategoryId).Include(include).Execute()



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
	esgMetricCategoryId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricCategoriesAPI.EsgMetricCategoriesEsgMetricCategoryIdGet(context.Background(), esgMetricCategoryId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricCategoriesAPI.EsgMetricCategoriesEsgMetricCategoryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricCategoriesEsgMetricCategoryIdGet`: EsgMetricCategories
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricCategoriesAPI.EsgMetricCategoriesEsgMetricCategoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricCategoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricCategoriesEsgMetricCategoryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EsgMetricCategories**](EsgMetricCategories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricCategoriesEsgMetricCategoryIdPut

> EsgMetricCategories EsgMetricCategoriesEsgMetricCategoryIdPut(ctx, esgMetricCategoryId).EsgMetricCategoriesPut(esgMetricCategoriesPut).Execute()



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
	esgMetricCategoryId := int64(789) // int64 | Model id
	esgMetricCategoriesPut := *openapiclient.NewEsgMetricCategoriesPut() // EsgMetricCategoriesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricCategoriesAPI.EsgMetricCategoriesEsgMetricCategoryIdPut(context.Background(), esgMetricCategoryId).EsgMetricCategoriesPut(esgMetricCategoriesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricCategoriesAPI.EsgMetricCategoriesEsgMetricCategoryIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricCategoriesEsgMetricCategoryIdPut`: EsgMetricCategories
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricCategoriesAPI.EsgMetricCategoriesEsgMetricCategoryIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricCategoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricCategoriesEsgMetricCategoryIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **esgMetricCategoriesPut** | [**EsgMetricCategoriesPut**](EsgMetricCategoriesPut.md) |  | 

### Return type

[**EsgMetricCategories**](EsgMetricCategories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricCategoriesGet

> EsgMetricCategoriesGet200Response EsgMetricCategoriesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.EsgMetricCategoriesAPI.EsgMetricCategoriesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricCategoriesAPI.EsgMetricCategoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricCategoriesGet`: EsgMetricCategoriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricCategoriesAPI.EsgMetricCategoriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricCategoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EsgMetricCategoriesGet200Response**](EsgMetricCategoriesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricCategoriesPost

> EsgMetricCategoriesGet200Response EsgMetricCategoriesPost(ctx).EsgMetricCategoriesPostRequest(esgMetricCategoriesPostRequest).Execute()



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
	esgMetricCategoriesPostRequest := *openapiclient.NewEsgMetricCategoriesPostRequest() // EsgMetricCategoriesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricCategoriesAPI.EsgMetricCategoriesPost(context.Background()).EsgMetricCategoriesPostRequest(esgMetricCategoriesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricCategoriesAPI.EsgMetricCategoriesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricCategoriesPost`: EsgMetricCategoriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricCategoriesAPI.EsgMetricCategoriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricCategoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **esgMetricCategoriesPostRequest** | [**EsgMetricCategoriesPostRequest**](EsgMetricCategoriesPostRequest.md) |  | 

### Return type

[**EsgMetricCategoriesGet200Response**](EsgMetricCategoriesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

