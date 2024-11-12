# \RiskCategoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RiskCategoriesGet**](RiskCategoriesAPI.md#RiskCategoriesGet) | **Get** /risk_categories | 
[**RiskCategoriesPost**](RiskCategoriesAPI.md#RiskCategoriesPost) | **Post** /risk_categories | 
[**RiskCategoriesRiskCategoryIdDelete**](RiskCategoriesAPI.md#RiskCategoriesRiskCategoryIdDelete) | **Delete** /risk_categories/{risk_category_id} | 
[**RiskCategoriesRiskCategoryIdGet**](RiskCategoriesAPI.md#RiskCategoriesRiskCategoryIdGet) | **Get** /risk_categories/{risk_category_id} | 
[**RiskCategoriesRiskCategoryIdPut**](RiskCategoriesAPI.md#RiskCategoriesRiskCategoryIdPut) | **Put** /risk_categories/{risk_category_id} | 



## RiskCategoriesGet

> RiskCategoriesGet200Response RiskCategoriesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.RiskCategoriesAPI.RiskCategoriesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskCategoriesAPI.RiskCategoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskCategoriesGet`: RiskCategoriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RiskCategoriesAPI.RiskCategoriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskCategoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**RiskCategoriesGet200Response**](RiskCategoriesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskCategoriesPost

> RiskCategoriesGet200Response RiskCategoriesPost(ctx).RiskCategoriesPostRequest(riskCategoriesPostRequest).Execute()



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
	riskCategoriesPostRequest := *openapiclient.NewRiskCategoriesPostRequest() // RiskCategoriesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskCategoriesAPI.RiskCategoriesPost(context.Background()).RiskCategoriesPostRequest(riskCategoriesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskCategoriesAPI.RiskCategoriesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskCategoriesPost`: RiskCategoriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RiskCategoriesAPI.RiskCategoriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskCategoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskCategoriesPostRequest** | [**RiskCategoriesPostRequest**](RiskCategoriesPostRequest.md) |  | 

### Return type

[**RiskCategoriesGet200Response**](RiskCategoriesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskCategoriesRiskCategoryIdDelete

> RiskCategories RiskCategoriesRiskCategoryIdDelete(ctx, riskCategoryId).Execute()



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
	riskCategoryId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskCategoriesAPI.RiskCategoriesRiskCategoryIdDelete(context.Background(), riskCategoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskCategoriesAPI.RiskCategoriesRiskCategoryIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskCategoriesRiskCategoryIdDelete`: RiskCategories
	fmt.Fprintf(os.Stdout, "Response from `RiskCategoriesAPI.RiskCategoriesRiskCategoryIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskCategoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskCategoriesRiskCategoryIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RiskCategories**](RiskCategories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskCategoriesRiskCategoryIdGet

> RiskCategories RiskCategoriesRiskCategoryIdGet(ctx, riskCategoryId).Include(include).Execute()



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
	riskCategoryId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskCategoriesAPI.RiskCategoriesRiskCategoryIdGet(context.Background(), riskCategoryId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskCategoriesAPI.RiskCategoriesRiskCategoryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskCategoriesRiskCategoryIdGet`: RiskCategories
	fmt.Fprintf(os.Stdout, "Response from `RiskCategoriesAPI.RiskCategoriesRiskCategoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskCategoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskCategoriesRiskCategoryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**RiskCategories**](RiskCategories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskCategoriesRiskCategoryIdPut

> RiskCategories RiskCategoriesRiskCategoryIdPut(ctx, riskCategoryId).RiskCategoriesPut(riskCategoriesPut).Execute()



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
	riskCategoryId := int64(789) // int64 | Model id
	riskCategoriesPut := *openapiclient.NewRiskCategoriesPut() // RiskCategoriesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskCategoriesAPI.RiskCategoriesRiskCategoryIdPut(context.Background(), riskCategoryId).RiskCategoriesPut(riskCategoriesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskCategoriesAPI.RiskCategoriesRiskCategoryIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskCategoriesRiskCategoryIdPut`: RiskCategories
	fmt.Fprintf(os.Stdout, "Response from `RiskCategoriesAPI.RiskCategoriesRiskCategoryIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskCategoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskCategoriesRiskCategoryIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **riskCategoriesPut** | [**RiskCategoriesPut**](RiskCategoriesPut.md) |  | 

### Return type

[**RiskCategories**](RiskCategories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

