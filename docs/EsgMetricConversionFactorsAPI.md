# \EsgMetricConversionFactorsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EsgMetricConversionFactorsEsgMetricConversionFactorIdDelete**](EsgMetricConversionFactorsAPI.md#EsgMetricConversionFactorsEsgMetricConversionFactorIdDelete) | **Delete** /esg_metric_conversion_factors/{esg_metric_conversion_factor_id} | 
[**EsgMetricConversionFactorsEsgMetricConversionFactorIdGet**](EsgMetricConversionFactorsAPI.md#EsgMetricConversionFactorsEsgMetricConversionFactorIdGet) | **Get** /esg_metric_conversion_factors/{esg_metric_conversion_factor_id} | 
[**EsgMetricConversionFactorsEsgMetricConversionFactorIdPut**](EsgMetricConversionFactorsAPI.md#EsgMetricConversionFactorsEsgMetricConversionFactorIdPut) | **Put** /esg_metric_conversion_factors/{esg_metric_conversion_factor_id} | 
[**EsgMetricConversionFactorsGet**](EsgMetricConversionFactorsAPI.md#EsgMetricConversionFactorsGet) | **Get** /esg_metric_conversion_factors | 
[**EsgMetricConversionFactorsPost**](EsgMetricConversionFactorsAPI.md#EsgMetricConversionFactorsPost) | **Post** /esg_metric_conversion_factors | 



## EsgMetricConversionFactorsEsgMetricConversionFactorIdDelete

> EsgMetricConversionFactors EsgMetricConversionFactorsEsgMetricConversionFactorIdDelete(ctx, esgMetricConversionFactorId).Execute()



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
	esgMetricConversionFactorId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricConversionFactorsAPI.EsgMetricConversionFactorsEsgMetricConversionFactorIdDelete(context.Background(), esgMetricConversionFactorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricConversionFactorsAPI.EsgMetricConversionFactorsEsgMetricConversionFactorIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricConversionFactorsEsgMetricConversionFactorIdDelete`: EsgMetricConversionFactors
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricConversionFactorsAPI.EsgMetricConversionFactorsEsgMetricConversionFactorIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricConversionFactorId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricConversionFactorsEsgMetricConversionFactorIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EsgMetricConversionFactors**](EsgMetricConversionFactors.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricConversionFactorsEsgMetricConversionFactorIdGet

> EsgMetricConversionFactors EsgMetricConversionFactorsEsgMetricConversionFactorIdGet(ctx, esgMetricConversionFactorId).Include(include).Execute()



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
	esgMetricConversionFactorId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricConversionFactorsAPI.EsgMetricConversionFactorsEsgMetricConversionFactorIdGet(context.Background(), esgMetricConversionFactorId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricConversionFactorsAPI.EsgMetricConversionFactorsEsgMetricConversionFactorIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricConversionFactorsEsgMetricConversionFactorIdGet`: EsgMetricConversionFactors
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricConversionFactorsAPI.EsgMetricConversionFactorsEsgMetricConversionFactorIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricConversionFactorId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricConversionFactorsEsgMetricConversionFactorIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EsgMetricConversionFactors**](EsgMetricConversionFactors.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricConversionFactorsEsgMetricConversionFactorIdPut

> EsgMetricConversionFactors EsgMetricConversionFactorsEsgMetricConversionFactorIdPut(ctx, esgMetricConversionFactorId).EsgMetricConversionFactorsPut(esgMetricConversionFactorsPut).Execute()



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
	esgMetricConversionFactorId := int64(789) // int64 | Model id
	esgMetricConversionFactorsPut := *openapiclient.NewEsgMetricConversionFactorsPut() // EsgMetricConversionFactorsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricConversionFactorsAPI.EsgMetricConversionFactorsEsgMetricConversionFactorIdPut(context.Background(), esgMetricConversionFactorId).EsgMetricConversionFactorsPut(esgMetricConversionFactorsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricConversionFactorsAPI.EsgMetricConversionFactorsEsgMetricConversionFactorIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricConversionFactorsEsgMetricConversionFactorIdPut`: EsgMetricConversionFactors
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricConversionFactorsAPI.EsgMetricConversionFactorsEsgMetricConversionFactorIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricConversionFactorId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricConversionFactorsEsgMetricConversionFactorIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **esgMetricConversionFactorsPut** | [**EsgMetricConversionFactorsPut**](EsgMetricConversionFactorsPut.md) |  | 

### Return type

[**EsgMetricConversionFactors**](EsgMetricConversionFactors.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricConversionFactorsGet

> EsgMetricConversionFactorsGet200Response EsgMetricConversionFactorsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.EsgMetricConversionFactorsAPI.EsgMetricConversionFactorsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricConversionFactorsAPI.EsgMetricConversionFactorsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricConversionFactorsGet`: EsgMetricConversionFactorsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricConversionFactorsAPI.EsgMetricConversionFactorsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricConversionFactorsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EsgMetricConversionFactorsGet200Response**](EsgMetricConversionFactorsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricConversionFactorsPost

> EsgMetricConversionFactorsGet200Response EsgMetricConversionFactorsPost(ctx).EsgMetricConversionFactorsPostRequest(esgMetricConversionFactorsPostRequest).Execute()



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
	esgMetricConversionFactorsPostRequest := *openapiclient.NewEsgMetricConversionFactorsPostRequest() // EsgMetricConversionFactorsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricConversionFactorsAPI.EsgMetricConversionFactorsPost(context.Background()).EsgMetricConversionFactorsPostRequest(esgMetricConversionFactorsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricConversionFactorsAPI.EsgMetricConversionFactorsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricConversionFactorsPost`: EsgMetricConversionFactorsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricConversionFactorsAPI.EsgMetricConversionFactorsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricConversionFactorsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **esgMetricConversionFactorsPostRequest** | [**EsgMetricConversionFactorsPostRequest**](EsgMetricConversionFactorsPostRequest.md) |  | 

### Return type

[**EsgMetricConversionFactorsGet200Response**](EsgMetricConversionFactorsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

