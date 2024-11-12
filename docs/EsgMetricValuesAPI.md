# \EsgMetricValuesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EsgMetricValuesEsgMetricValueIdDelete**](EsgMetricValuesAPI.md#EsgMetricValuesEsgMetricValueIdDelete) | **Delete** /esg_metric_values/{esg_metric_value_id} | 
[**EsgMetricValuesEsgMetricValueIdGet**](EsgMetricValuesAPI.md#EsgMetricValuesEsgMetricValueIdGet) | **Get** /esg_metric_values/{esg_metric_value_id} | 
[**EsgMetricValuesEsgMetricValueIdPut**](EsgMetricValuesAPI.md#EsgMetricValuesEsgMetricValueIdPut) | **Put** /esg_metric_values/{esg_metric_value_id} | 
[**EsgMetricValuesGet**](EsgMetricValuesAPI.md#EsgMetricValuesGet) | **Get** /esg_metric_values | 
[**EsgMetricValuesPost**](EsgMetricValuesAPI.md#EsgMetricValuesPost) | **Post** /esg_metric_values | 



## EsgMetricValuesEsgMetricValueIdDelete

> EsgMetricValues EsgMetricValuesEsgMetricValueIdDelete(ctx, esgMetricValueId).Execute()



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
	esgMetricValueId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricValuesAPI.EsgMetricValuesEsgMetricValueIdDelete(context.Background(), esgMetricValueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricValuesAPI.EsgMetricValuesEsgMetricValueIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricValuesEsgMetricValueIdDelete`: EsgMetricValues
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricValuesAPI.EsgMetricValuesEsgMetricValueIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricValueId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricValuesEsgMetricValueIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EsgMetricValues**](EsgMetricValues.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricValuesEsgMetricValueIdGet

> EsgMetricValues EsgMetricValuesEsgMetricValueIdGet(ctx, esgMetricValueId).Include(include).Execute()



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
	esgMetricValueId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricValuesAPI.EsgMetricValuesEsgMetricValueIdGet(context.Background(), esgMetricValueId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricValuesAPI.EsgMetricValuesEsgMetricValueIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricValuesEsgMetricValueIdGet`: EsgMetricValues
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricValuesAPI.EsgMetricValuesEsgMetricValueIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricValueId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricValuesEsgMetricValueIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EsgMetricValues**](EsgMetricValues.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricValuesEsgMetricValueIdPut

> EsgMetricValues EsgMetricValuesEsgMetricValueIdPut(ctx, esgMetricValueId).EsgMetricValuesPut(esgMetricValuesPut).Execute()



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
	esgMetricValueId := int64(789) // int64 | Model id
	esgMetricValuesPut := *openapiclient.NewEsgMetricValuesPut() // EsgMetricValuesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricValuesAPI.EsgMetricValuesEsgMetricValueIdPut(context.Background(), esgMetricValueId).EsgMetricValuesPut(esgMetricValuesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricValuesAPI.EsgMetricValuesEsgMetricValueIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricValuesEsgMetricValueIdPut`: EsgMetricValues
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricValuesAPI.EsgMetricValuesEsgMetricValueIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricValueId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricValuesEsgMetricValueIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **esgMetricValuesPut** | [**EsgMetricValuesPut**](EsgMetricValuesPut.md) |  | 

### Return type

[**EsgMetricValues**](EsgMetricValues.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricValuesGet

> EsgMetricValuesGet200Response EsgMetricValuesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.EsgMetricValuesAPI.EsgMetricValuesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricValuesAPI.EsgMetricValuesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricValuesGet`: EsgMetricValuesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricValuesAPI.EsgMetricValuesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricValuesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EsgMetricValuesGet200Response**](EsgMetricValuesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricValuesPost

> EsgMetricValuesGet200Response EsgMetricValuesPost(ctx).EsgMetricValuesPostRequest(esgMetricValuesPostRequest).Execute()



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
	esgMetricValuesPostRequest := *openapiclient.NewEsgMetricValuesPostRequest() // EsgMetricValuesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricValuesAPI.EsgMetricValuesPost(context.Background()).EsgMetricValuesPostRequest(esgMetricValuesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricValuesAPI.EsgMetricValuesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricValuesPost`: EsgMetricValuesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricValuesAPI.EsgMetricValuesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricValuesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **esgMetricValuesPostRequest** | [**EsgMetricValuesPostRequest**](EsgMetricValuesPostRequest.md) |  | 

### Return type

[**EsgMetricValuesGet200Response**](EsgMetricValuesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

