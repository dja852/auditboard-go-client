# \EsgMetricFrequencyOptionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdDelete**](EsgMetricFrequencyOptionsAPI.md#EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdDelete) | **Delete** /esg_metric_frequency_options/{esg_metric_frequency_option_id} | 
[**EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdGet**](EsgMetricFrequencyOptionsAPI.md#EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdGet) | **Get** /esg_metric_frequency_options/{esg_metric_frequency_option_id} | 
[**EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdPut**](EsgMetricFrequencyOptionsAPI.md#EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdPut) | **Put** /esg_metric_frequency_options/{esg_metric_frequency_option_id} | 
[**EsgMetricFrequencyOptionsGet**](EsgMetricFrequencyOptionsAPI.md#EsgMetricFrequencyOptionsGet) | **Get** /esg_metric_frequency_options | 
[**EsgMetricFrequencyOptionsPost**](EsgMetricFrequencyOptionsAPI.md#EsgMetricFrequencyOptionsPost) | **Post** /esg_metric_frequency_options | 



## EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdDelete

> EsgMetricFrequencyOptions EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdDelete(ctx, esgMetricFrequencyOptionId).Execute()



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
	esgMetricFrequencyOptionId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricFrequencyOptionsAPI.EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdDelete(context.Background(), esgMetricFrequencyOptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricFrequencyOptionsAPI.EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdDelete`: EsgMetricFrequencyOptions
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricFrequencyOptionsAPI.EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricFrequencyOptionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EsgMetricFrequencyOptions**](EsgMetricFrequencyOptions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdGet

> EsgMetricFrequencyOptions EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdGet(ctx, esgMetricFrequencyOptionId).Include(include).Execute()



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
	esgMetricFrequencyOptionId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricFrequencyOptionsAPI.EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdGet(context.Background(), esgMetricFrequencyOptionId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricFrequencyOptionsAPI.EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdGet`: EsgMetricFrequencyOptions
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricFrequencyOptionsAPI.EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricFrequencyOptionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EsgMetricFrequencyOptions**](EsgMetricFrequencyOptions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdPut

> EsgMetricFrequencyOptions EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdPut(ctx, esgMetricFrequencyOptionId).EsgMetricFrequencyOptionsPut(esgMetricFrequencyOptionsPut).Execute()



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
	esgMetricFrequencyOptionId := int64(789) // int64 | Model id
	esgMetricFrequencyOptionsPut := *openapiclient.NewEsgMetricFrequencyOptionsPut() // EsgMetricFrequencyOptionsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricFrequencyOptionsAPI.EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdPut(context.Background(), esgMetricFrequencyOptionId).EsgMetricFrequencyOptionsPut(esgMetricFrequencyOptionsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricFrequencyOptionsAPI.EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdPut`: EsgMetricFrequencyOptions
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricFrequencyOptionsAPI.EsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricFrequencyOptionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricFrequencyOptionsEsgMetricFrequencyOptionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **esgMetricFrequencyOptionsPut** | [**EsgMetricFrequencyOptionsPut**](EsgMetricFrequencyOptionsPut.md) |  | 

### Return type

[**EsgMetricFrequencyOptions**](EsgMetricFrequencyOptions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricFrequencyOptionsGet

> EsgMetricFrequencyOptionsGet200Response EsgMetricFrequencyOptionsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.EsgMetricFrequencyOptionsAPI.EsgMetricFrequencyOptionsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricFrequencyOptionsAPI.EsgMetricFrequencyOptionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricFrequencyOptionsGet`: EsgMetricFrequencyOptionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricFrequencyOptionsAPI.EsgMetricFrequencyOptionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricFrequencyOptionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EsgMetricFrequencyOptionsGet200Response**](EsgMetricFrequencyOptionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricFrequencyOptionsPost

> EsgMetricFrequencyOptionsGet200Response EsgMetricFrequencyOptionsPost(ctx).EsgMetricFrequencyOptionsPostRequest(esgMetricFrequencyOptionsPostRequest).Execute()



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
	esgMetricFrequencyOptionsPostRequest := *openapiclient.NewEsgMetricFrequencyOptionsPostRequest() // EsgMetricFrequencyOptionsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricFrequencyOptionsAPI.EsgMetricFrequencyOptionsPost(context.Background()).EsgMetricFrequencyOptionsPostRequest(esgMetricFrequencyOptionsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricFrequencyOptionsAPI.EsgMetricFrequencyOptionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricFrequencyOptionsPost`: EsgMetricFrequencyOptionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricFrequencyOptionsAPI.EsgMetricFrequencyOptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricFrequencyOptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **esgMetricFrequencyOptionsPostRequest** | [**EsgMetricFrequencyOptionsPostRequest**](EsgMetricFrequencyOptionsPostRequest.md) |  | 

### Return type

[**EsgMetricFrequencyOptionsGet200Response**](EsgMetricFrequencyOptionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

