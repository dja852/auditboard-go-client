# \EsgMetricUnitOptionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EsgMetricUnitOptionsEsgMetricUnitOptionIdDelete**](EsgMetricUnitOptionsAPI.md#EsgMetricUnitOptionsEsgMetricUnitOptionIdDelete) | **Delete** /esg_metric_unit_options/{esg_metric_unit_option_id} | 
[**EsgMetricUnitOptionsEsgMetricUnitOptionIdGet**](EsgMetricUnitOptionsAPI.md#EsgMetricUnitOptionsEsgMetricUnitOptionIdGet) | **Get** /esg_metric_unit_options/{esg_metric_unit_option_id} | 
[**EsgMetricUnitOptionsEsgMetricUnitOptionIdPut**](EsgMetricUnitOptionsAPI.md#EsgMetricUnitOptionsEsgMetricUnitOptionIdPut) | **Put** /esg_metric_unit_options/{esg_metric_unit_option_id} | 
[**EsgMetricUnitOptionsGet**](EsgMetricUnitOptionsAPI.md#EsgMetricUnitOptionsGet) | **Get** /esg_metric_unit_options | 
[**EsgMetricUnitOptionsPost**](EsgMetricUnitOptionsAPI.md#EsgMetricUnitOptionsPost) | **Post** /esg_metric_unit_options | 



## EsgMetricUnitOptionsEsgMetricUnitOptionIdDelete

> EsgMetricUnitOptions EsgMetricUnitOptionsEsgMetricUnitOptionIdDelete(ctx, esgMetricUnitOptionId).Execute()



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
	esgMetricUnitOptionId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricUnitOptionsAPI.EsgMetricUnitOptionsEsgMetricUnitOptionIdDelete(context.Background(), esgMetricUnitOptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricUnitOptionsAPI.EsgMetricUnitOptionsEsgMetricUnitOptionIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricUnitOptionsEsgMetricUnitOptionIdDelete`: EsgMetricUnitOptions
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricUnitOptionsAPI.EsgMetricUnitOptionsEsgMetricUnitOptionIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricUnitOptionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricUnitOptionsEsgMetricUnitOptionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EsgMetricUnitOptions**](EsgMetricUnitOptions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricUnitOptionsEsgMetricUnitOptionIdGet

> EsgMetricUnitOptions EsgMetricUnitOptionsEsgMetricUnitOptionIdGet(ctx, esgMetricUnitOptionId).Include(include).Execute()



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
	esgMetricUnitOptionId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricUnitOptionsAPI.EsgMetricUnitOptionsEsgMetricUnitOptionIdGet(context.Background(), esgMetricUnitOptionId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricUnitOptionsAPI.EsgMetricUnitOptionsEsgMetricUnitOptionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricUnitOptionsEsgMetricUnitOptionIdGet`: EsgMetricUnitOptions
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricUnitOptionsAPI.EsgMetricUnitOptionsEsgMetricUnitOptionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricUnitOptionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricUnitOptionsEsgMetricUnitOptionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EsgMetricUnitOptions**](EsgMetricUnitOptions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricUnitOptionsEsgMetricUnitOptionIdPut

> EsgMetricUnitOptions EsgMetricUnitOptionsEsgMetricUnitOptionIdPut(ctx, esgMetricUnitOptionId).EsgMetricUnitOptionsPut(esgMetricUnitOptionsPut).Execute()



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
	esgMetricUnitOptionId := int64(789) // int64 | Model id
	esgMetricUnitOptionsPut := *openapiclient.NewEsgMetricUnitOptionsPut() // EsgMetricUnitOptionsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricUnitOptionsAPI.EsgMetricUnitOptionsEsgMetricUnitOptionIdPut(context.Background(), esgMetricUnitOptionId).EsgMetricUnitOptionsPut(esgMetricUnitOptionsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricUnitOptionsAPI.EsgMetricUnitOptionsEsgMetricUnitOptionIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricUnitOptionsEsgMetricUnitOptionIdPut`: EsgMetricUnitOptions
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricUnitOptionsAPI.EsgMetricUnitOptionsEsgMetricUnitOptionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricUnitOptionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricUnitOptionsEsgMetricUnitOptionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **esgMetricUnitOptionsPut** | [**EsgMetricUnitOptionsPut**](EsgMetricUnitOptionsPut.md) |  | 

### Return type

[**EsgMetricUnitOptions**](EsgMetricUnitOptions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricUnitOptionsGet

> EsgMetricUnitOptionsGet200Response EsgMetricUnitOptionsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.EsgMetricUnitOptionsAPI.EsgMetricUnitOptionsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricUnitOptionsAPI.EsgMetricUnitOptionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricUnitOptionsGet`: EsgMetricUnitOptionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricUnitOptionsAPI.EsgMetricUnitOptionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricUnitOptionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EsgMetricUnitOptionsGet200Response**](EsgMetricUnitOptionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricUnitOptionsPost

> EsgMetricUnitOptionsGet200Response EsgMetricUnitOptionsPost(ctx).EsgMetricUnitOptionsPostRequest(esgMetricUnitOptionsPostRequest).Execute()



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
	esgMetricUnitOptionsPostRequest := *openapiclient.NewEsgMetricUnitOptionsPostRequest() // EsgMetricUnitOptionsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricUnitOptionsAPI.EsgMetricUnitOptionsPost(context.Background()).EsgMetricUnitOptionsPostRequest(esgMetricUnitOptionsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricUnitOptionsAPI.EsgMetricUnitOptionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricUnitOptionsPost`: EsgMetricUnitOptionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricUnitOptionsAPI.EsgMetricUnitOptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricUnitOptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **esgMetricUnitOptionsPostRequest** | [**EsgMetricUnitOptionsPostRequest**](EsgMetricUnitOptionsPostRequest.md) |  | 

### Return type

[**EsgMetricUnitOptionsGet200Response**](EsgMetricUnitOptionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

