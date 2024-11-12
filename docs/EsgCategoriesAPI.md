# \EsgCategoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EsgCategoriesEsgCategoryIdDelete**](EsgCategoriesAPI.md#EsgCategoriesEsgCategoryIdDelete) | **Delete** /esg_categories/{esg_category_id} | 
[**EsgCategoriesEsgCategoryIdGet**](EsgCategoriesAPI.md#EsgCategoriesEsgCategoryIdGet) | **Get** /esg_categories/{esg_category_id} | 
[**EsgCategoriesEsgCategoryIdPut**](EsgCategoriesAPI.md#EsgCategoriesEsgCategoryIdPut) | **Put** /esg_categories/{esg_category_id} | 
[**EsgCategoriesGet**](EsgCategoriesAPI.md#EsgCategoriesGet) | **Get** /esg_categories | 
[**EsgCategoriesPost**](EsgCategoriesAPI.md#EsgCategoriesPost) | **Post** /esg_categories | 



## EsgCategoriesEsgCategoryIdDelete

> EsgCategories EsgCategoriesEsgCategoryIdDelete(ctx, esgCategoryId).Execute()



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
	esgCategoryId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgCategoriesAPI.EsgCategoriesEsgCategoryIdDelete(context.Background(), esgCategoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgCategoriesAPI.EsgCategoriesEsgCategoryIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgCategoriesEsgCategoryIdDelete`: EsgCategories
	fmt.Fprintf(os.Stdout, "Response from `EsgCategoriesAPI.EsgCategoriesEsgCategoryIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgCategoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgCategoriesEsgCategoryIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EsgCategories**](EsgCategories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgCategoriesEsgCategoryIdGet

> EsgCategories EsgCategoriesEsgCategoryIdGet(ctx, esgCategoryId).Include(include).Execute()



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
	esgCategoryId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgCategoriesAPI.EsgCategoriesEsgCategoryIdGet(context.Background(), esgCategoryId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgCategoriesAPI.EsgCategoriesEsgCategoryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgCategoriesEsgCategoryIdGet`: EsgCategories
	fmt.Fprintf(os.Stdout, "Response from `EsgCategoriesAPI.EsgCategoriesEsgCategoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgCategoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgCategoriesEsgCategoryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EsgCategories**](EsgCategories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgCategoriesEsgCategoryIdPut

> EsgCategories EsgCategoriesEsgCategoryIdPut(ctx, esgCategoryId).EsgCategoriesPut(esgCategoriesPut).Execute()



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
	esgCategoryId := int64(789) // int64 | Model id
	esgCategoriesPut := *openapiclient.NewEsgCategoriesPut() // EsgCategoriesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgCategoriesAPI.EsgCategoriesEsgCategoryIdPut(context.Background(), esgCategoryId).EsgCategoriesPut(esgCategoriesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgCategoriesAPI.EsgCategoriesEsgCategoryIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgCategoriesEsgCategoryIdPut`: EsgCategories
	fmt.Fprintf(os.Stdout, "Response from `EsgCategoriesAPI.EsgCategoriesEsgCategoryIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgCategoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgCategoriesEsgCategoryIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **esgCategoriesPut** | [**EsgCategoriesPut**](EsgCategoriesPut.md) |  | 

### Return type

[**EsgCategories**](EsgCategories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgCategoriesGet

> EsgCategoriesGet200Response EsgCategoriesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.EsgCategoriesAPI.EsgCategoriesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgCategoriesAPI.EsgCategoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgCategoriesGet`: EsgCategoriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EsgCategoriesAPI.EsgCategoriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsgCategoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EsgCategoriesGet200Response**](EsgCategoriesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgCategoriesPost

> EsgCategoriesGet200Response EsgCategoriesPost(ctx).EsgCategoriesPostRequest(esgCategoriesPostRequest).Execute()



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
	esgCategoriesPostRequest := *openapiclient.NewEsgCategoriesPostRequest() // EsgCategoriesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgCategoriesAPI.EsgCategoriesPost(context.Background()).EsgCategoriesPostRequest(esgCategoriesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgCategoriesAPI.EsgCategoriesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgCategoriesPost`: EsgCategoriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EsgCategoriesAPI.EsgCategoriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsgCategoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **esgCategoriesPostRequest** | [**EsgCategoriesPostRequest**](EsgCategoriesPostRequest.md) |  | 

### Return type

[**EsgCategoriesGet200Response**](EsgCategoriesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

