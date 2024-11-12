# \EsgMetricsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EsgMetricsEsgMetricIdDelete**](EsgMetricsAPI.md#EsgMetricsEsgMetricIdDelete) | **Delete** /esg_metrics/{esg_metric_id} | 
[**EsgMetricsEsgMetricIdGet**](EsgMetricsAPI.md#EsgMetricsEsgMetricIdGet) | **Get** /esg_metrics/{esg_metric_id} | 
[**EsgMetricsEsgMetricIdPut**](EsgMetricsAPI.md#EsgMetricsEsgMetricIdPut) | **Put** /esg_metrics/{esg_metric_id} | 
[**EsgMetricsGet**](EsgMetricsAPI.md#EsgMetricsGet) | **Get** /esg_metrics | 
[**EsgMetricsPost**](EsgMetricsAPI.md#EsgMetricsPost) | **Post** /esg_metrics | 



## EsgMetricsEsgMetricIdDelete

> EsgMetrics EsgMetricsEsgMetricIdDelete(ctx, esgMetricId).Execute()



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
	esgMetricId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricsAPI.EsgMetricsEsgMetricIdDelete(context.Background(), esgMetricId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricsAPI.EsgMetricsEsgMetricIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricsEsgMetricIdDelete`: EsgMetrics
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricsAPI.EsgMetricsEsgMetricIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricsEsgMetricIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EsgMetrics**](EsgMetrics.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricsEsgMetricIdGet

> EsgMetrics EsgMetricsEsgMetricIdGet(ctx, esgMetricId).Include(include).Execute()



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
	esgMetricId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricsAPI.EsgMetricsEsgMetricIdGet(context.Background(), esgMetricId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricsAPI.EsgMetricsEsgMetricIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricsEsgMetricIdGet`: EsgMetrics
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricsAPI.EsgMetricsEsgMetricIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricsEsgMetricIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EsgMetrics**](EsgMetrics.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricsEsgMetricIdPut

> EsgMetrics EsgMetricsEsgMetricIdPut(ctx, esgMetricId).EsgMetricsPut(esgMetricsPut).Execute()



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
	esgMetricId := int64(789) // int64 | Model id
	esgMetricsPut := *openapiclient.NewEsgMetricsPut() // EsgMetricsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricsAPI.EsgMetricsEsgMetricIdPut(context.Background(), esgMetricId).EsgMetricsPut(esgMetricsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricsAPI.EsgMetricsEsgMetricIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricsEsgMetricIdPut`: EsgMetrics
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricsAPI.EsgMetricsEsgMetricIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricsEsgMetricIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **esgMetricsPut** | [**EsgMetricsPut**](EsgMetricsPut.md) |  | 

### Return type

[**EsgMetrics**](EsgMetrics.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricsGet

> EsgMetricsGet200Response EsgMetricsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.EsgMetricsAPI.EsgMetricsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricsAPI.EsgMetricsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricsGet`: EsgMetricsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricsAPI.EsgMetricsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EsgMetricsGet200Response**](EsgMetricsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricsPost

> EsgMetricsGet200Response EsgMetricsPost(ctx).EsgMetricsPostRequest(esgMetricsPostRequest).Execute()



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
	esgMetricsPostRequest := *openapiclient.NewEsgMetricsPostRequest() // EsgMetricsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricsAPI.EsgMetricsPost(context.Background()).EsgMetricsPostRequest(esgMetricsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricsAPI.EsgMetricsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricsPost`: EsgMetricsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricsAPI.EsgMetricsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **esgMetricsPostRequest** | [**EsgMetricsPostRequest**](EsgMetricsPostRequest.md) |  | 

### Return type

[**EsgMetricsGet200Response**](EsgMetricsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

