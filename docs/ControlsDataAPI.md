# \ControlsDataAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ControlsDataControlsDatumIdDelete**](ControlsDataAPI.md#ControlsDataControlsDatumIdDelete) | **Delete** /controls_data/{controls_datum_id} | 
[**ControlsDataControlsDatumIdGet**](ControlsDataAPI.md#ControlsDataControlsDatumIdGet) | **Get** /controls_data/{controls_datum_id} | 
[**ControlsDataControlsDatumIdPut**](ControlsDataAPI.md#ControlsDataControlsDatumIdPut) | **Put** /controls_data/{controls_datum_id} | 
[**ControlsDataGet**](ControlsDataAPI.md#ControlsDataGet) | **Get** /controls_data | 
[**ControlsDataPost**](ControlsDataAPI.md#ControlsDataPost) | **Post** /controls_data | 
[**ControlsDataRequestZipPost**](ControlsDataAPI.md#ControlsDataRequestZipPost) | **Post** /controls_data/request_zip | 



## ControlsDataControlsDatumIdDelete

> ControlsData ControlsDataControlsDatumIdDelete(ctx, controlsDatumId).Execute()



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
	controlsDatumId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControlsDataAPI.ControlsDataControlsDatumIdDelete(context.Background(), controlsDatumId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControlsDataAPI.ControlsDataControlsDatumIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlsDataControlsDatumIdDelete`: ControlsData
	fmt.Fprintf(os.Stdout, "Response from `ControlsDataAPI.ControlsDataControlsDatumIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**controlsDatumId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiControlsDataControlsDatumIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ControlsData**](ControlsData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ControlsDataControlsDatumIdGet

> ControlsData ControlsDataControlsDatumIdGet(ctx, controlsDatumId).Include(include).Execute()



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
	controlsDatumId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControlsDataAPI.ControlsDataControlsDatumIdGet(context.Background(), controlsDatumId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControlsDataAPI.ControlsDataControlsDatumIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlsDataControlsDatumIdGet`: ControlsData
	fmt.Fprintf(os.Stdout, "Response from `ControlsDataAPI.ControlsDataControlsDatumIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**controlsDatumId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiControlsDataControlsDatumIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**ControlsData**](ControlsData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ControlsDataControlsDatumIdPut

> ControlsData ControlsDataControlsDatumIdPut(ctx, controlsDatumId).ControlsDataPut(controlsDataPut).Execute()



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
	controlsDatumId := int64(789) // int64 | Model id
	controlsDataPut := *openapiclient.NewControlsDataPut() // ControlsDataPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControlsDataAPI.ControlsDataControlsDatumIdPut(context.Background(), controlsDatumId).ControlsDataPut(controlsDataPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControlsDataAPI.ControlsDataControlsDatumIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlsDataControlsDatumIdPut`: ControlsData
	fmt.Fprintf(os.Stdout, "Response from `ControlsDataAPI.ControlsDataControlsDatumIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**controlsDatumId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiControlsDataControlsDatumIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **controlsDataPut** | [**ControlsDataPut**](ControlsDataPut.md) |  | 

### Return type

[**ControlsData**](ControlsData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ControlsDataGet

> ControlsDataGet200Response ControlsDataGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.ControlsDataAPI.ControlsDataGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControlsDataAPI.ControlsDataGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlsDataGet`: ControlsDataGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ControlsDataAPI.ControlsDataGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiControlsDataGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**ControlsDataGet200Response**](ControlsDataGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ControlsDataPost

> ControlsDataGet200Response ControlsDataPost(ctx).ControlsDataPostRequest(controlsDataPostRequest).Execute()



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
	controlsDataPostRequest := *openapiclient.NewControlsDataPostRequest() // ControlsDataPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControlsDataAPI.ControlsDataPost(context.Background()).ControlsDataPostRequest(controlsDataPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControlsDataAPI.ControlsDataPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlsDataPost`: ControlsDataGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ControlsDataAPI.ControlsDataPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiControlsDataPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controlsDataPostRequest** | [**ControlsDataPostRequest**](ControlsDataPostRequest.md) |  | 

### Return type

[**ControlsDataGet200Response**](ControlsDataGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ControlsDataRequestZipPost

> ReportsGenerateTimesheetsPost200Response ControlsDataRequestZipPost(ctx).ControlsDataRequestZipPostRequest(controlsDataRequestZipPostRequest).Execute()





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
	controlsDataRequestZipPostRequest := *openapiclient.NewControlsDataRequestZipPostRequest() // ControlsDataRequestZipPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControlsDataAPI.ControlsDataRequestZipPost(context.Background()).ControlsDataRequestZipPostRequest(controlsDataRequestZipPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControlsDataAPI.ControlsDataRequestZipPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlsDataRequestZipPost`: ReportsGenerateTimesheetsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ControlsDataAPI.ControlsDataRequestZipPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiControlsDataRequestZipPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controlsDataRequestZipPostRequest** | [**ControlsDataRequestZipPostRequest**](ControlsDataRequestZipPostRequest.md) |  | 

### Return type

[**ReportsGenerateTimesheetsPost200Response**](ReportsGenerateTimesheetsPost200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

