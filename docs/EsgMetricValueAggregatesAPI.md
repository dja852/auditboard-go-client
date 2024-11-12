# \EsgMetricValueAggregatesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EsgMetricValueAggregatesEsgMetricValueAggregateIdGet**](EsgMetricValueAggregatesAPI.md#EsgMetricValueAggregatesEsgMetricValueAggregateIdGet) | **Get** /esg_metric_value_aggregates/{esg_metric_value_aggregate_id} | 
[**EsgMetricValueAggregatesGet**](EsgMetricValueAggregatesAPI.md#EsgMetricValueAggregatesGet) | **Get** /esg_metric_value_aggregates | 



## EsgMetricValueAggregatesEsgMetricValueAggregateIdGet

> EsgMetricValueAggregates EsgMetricValueAggregatesEsgMetricValueAggregateIdGet(ctx, esgMetricValueAggregateId).Include(include).Execute()



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
	esgMetricValueAggregateId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricValueAggregatesAPI.EsgMetricValueAggregatesEsgMetricValueAggregateIdGet(context.Background(), esgMetricValueAggregateId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricValueAggregatesAPI.EsgMetricValueAggregatesEsgMetricValueAggregateIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricValueAggregatesEsgMetricValueAggregateIdGet`: EsgMetricValueAggregates
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricValueAggregatesAPI.EsgMetricValueAggregatesEsgMetricValueAggregateIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricValueAggregateId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricValueAggregatesEsgMetricValueAggregateIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EsgMetricValueAggregates**](EsgMetricValueAggregates.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricValueAggregatesGet

> EsgMetricValueAggregatesGet200Response EsgMetricValueAggregatesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.EsgMetricValueAggregatesAPI.EsgMetricValueAggregatesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricValueAggregatesAPI.EsgMetricValueAggregatesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricValueAggregatesGet`: EsgMetricValueAggregatesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricValueAggregatesAPI.EsgMetricValueAggregatesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricValueAggregatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EsgMetricValueAggregatesGet200Response**](EsgMetricValueAggregatesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

