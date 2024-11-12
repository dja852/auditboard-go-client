# \OpsAuditCategoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpsAuditCategoriesGet**](OpsAuditCategoriesAPI.md#OpsAuditCategoriesGet) | **Get** /ops_audit_categories | 
[**OpsAuditCategoriesOpsAuditCategoryIdDelete**](OpsAuditCategoriesAPI.md#OpsAuditCategoriesOpsAuditCategoryIdDelete) | **Delete** /ops_audit_categories/{ops_audit_category_id} | 
[**OpsAuditCategoriesOpsAuditCategoryIdGet**](OpsAuditCategoriesAPI.md#OpsAuditCategoriesOpsAuditCategoryIdGet) | **Get** /ops_audit_categories/{ops_audit_category_id} | 
[**OpsAuditCategoriesOpsAuditCategoryIdPut**](OpsAuditCategoriesAPI.md#OpsAuditCategoriesOpsAuditCategoryIdPut) | **Put** /ops_audit_categories/{ops_audit_category_id} | 
[**OpsAuditCategoriesPost**](OpsAuditCategoriesAPI.md#OpsAuditCategoriesPost) | **Post** /ops_audit_categories | 



## OpsAuditCategoriesGet

> OpsAuditCategoriesGet200Response OpsAuditCategoriesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.OpsAuditCategoriesAPI.OpsAuditCategoriesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditCategoriesAPI.OpsAuditCategoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditCategoriesGet`: OpsAuditCategoriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditCategoriesAPI.OpsAuditCategoriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditCategoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**OpsAuditCategoriesGet200Response**](OpsAuditCategoriesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditCategoriesOpsAuditCategoryIdDelete

> OpsAuditCategories OpsAuditCategoriesOpsAuditCategoryIdDelete(ctx, opsAuditCategoryId).Execute()



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
	opsAuditCategoryId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditCategoriesAPI.OpsAuditCategoriesOpsAuditCategoryIdDelete(context.Background(), opsAuditCategoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditCategoriesAPI.OpsAuditCategoriesOpsAuditCategoryIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditCategoriesOpsAuditCategoryIdDelete`: OpsAuditCategories
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditCategoriesAPI.OpsAuditCategoriesOpsAuditCategoryIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditCategoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditCategoriesOpsAuditCategoryIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpsAuditCategories**](OpsAuditCategories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditCategoriesOpsAuditCategoryIdGet

> OpsAuditCategories OpsAuditCategoriesOpsAuditCategoryIdGet(ctx, opsAuditCategoryId).Include(include).Execute()



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
	opsAuditCategoryId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditCategoriesAPI.OpsAuditCategoriesOpsAuditCategoryIdGet(context.Background(), opsAuditCategoryId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditCategoriesAPI.OpsAuditCategoriesOpsAuditCategoryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditCategoriesOpsAuditCategoryIdGet`: OpsAuditCategories
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditCategoriesAPI.OpsAuditCategoriesOpsAuditCategoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditCategoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditCategoriesOpsAuditCategoryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**OpsAuditCategories**](OpsAuditCategories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditCategoriesOpsAuditCategoryIdPut

> OpsAuditCategories OpsAuditCategoriesOpsAuditCategoryIdPut(ctx, opsAuditCategoryId).OpsAuditCategoriesPut(opsAuditCategoriesPut).Execute()



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
	opsAuditCategoryId := int64(789) // int64 | Model id
	opsAuditCategoriesPut := *openapiclient.NewOpsAuditCategoriesPut() // OpsAuditCategoriesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditCategoriesAPI.OpsAuditCategoriesOpsAuditCategoryIdPut(context.Background(), opsAuditCategoryId).OpsAuditCategoriesPut(opsAuditCategoriesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditCategoriesAPI.OpsAuditCategoriesOpsAuditCategoryIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditCategoriesOpsAuditCategoryIdPut`: OpsAuditCategories
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditCategoriesAPI.OpsAuditCategoriesOpsAuditCategoryIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditCategoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditCategoriesOpsAuditCategoryIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opsAuditCategoriesPut** | [**OpsAuditCategoriesPut**](OpsAuditCategoriesPut.md) |  | 

### Return type

[**OpsAuditCategories**](OpsAuditCategories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditCategoriesPost

> OpsAuditCategoriesGet200Response OpsAuditCategoriesPost(ctx).OpsAuditCategoriesPostRequest(opsAuditCategoriesPostRequest).Execute()



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
	opsAuditCategoriesPostRequest := *openapiclient.NewOpsAuditCategoriesPostRequest() // OpsAuditCategoriesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditCategoriesAPI.OpsAuditCategoriesPost(context.Background()).OpsAuditCategoriesPostRequest(opsAuditCategoriesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditCategoriesAPI.OpsAuditCategoriesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditCategoriesPost`: OpsAuditCategoriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditCategoriesAPI.OpsAuditCategoriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditCategoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opsAuditCategoriesPostRequest** | [**OpsAuditCategoriesPostRequest**](OpsAuditCategoriesPostRequest.md) |  | 

### Return type

[**OpsAuditCategoriesGet200Response**](OpsAuditCategoriesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

