# \SubprocessesDataAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubprocessesDataGet**](SubprocessesDataAPI.md#SubprocessesDataGet) | **Get** /subprocesses_data | 
[**SubprocessesDataPost**](SubprocessesDataAPI.md#SubprocessesDataPost) | **Post** /subprocesses_data | 
[**SubprocessesDataSubprocessesDatumIdDelete**](SubprocessesDataAPI.md#SubprocessesDataSubprocessesDatumIdDelete) | **Delete** /subprocesses_data/{subprocesses_datum_id} | 
[**SubprocessesDataSubprocessesDatumIdGet**](SubprocessesDataAPI.md#SubprocessesDataSubprocessesDatumIdGet) | **Get** /subprocesses_data/{subprocesses_datum_id} | 
[**SubprocessesDataSubprocessesDatumIdPut**](SubprocessesDataAPI.md#SubprocessesDataSubprocessesDatumIdPut) | **Put** /subprocesses_data/{subprocesses_datum_id} | 



## SubprocessesDataGet

> SubprocessesDataGet200Response SubprocessesDataGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.SubprocessesDataAPI.SubprocessesDataGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubprocessesDataAPI.SubprocessesDataGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubprocessesDataGet`: SubprocessesDataGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SubprocessesDataAPI.SubprocessesDataGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubprocessesDataGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**SubprocessesDataGet200Response**](SubprocessesDataGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubprocessesDataPost

> SubprocessesDataGet200Response SubprocessesDataPost(ctx).SubprocessesDataPostRequest(subprocessesDataPostRequest).Execute()



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
	subprocessesDataPostRequest := *openapiclient.NewSubprocessesDataPostRequest() // SubprocessesDataPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubprocessesDataAPI.SubprocessesDataPost(context.Background()).SubprocessesDataPostRequest(subprocessesDataPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubprocessesDataAPI.SubprocessesDataPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubprocessesDataPost`: SubprocessesDataGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SubprocessesDataAPI.SubprocessesDataPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubprocessesDataPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subprocessesDataPostRequest** | [**SubprocessesDataPostRequest**](SubprocessesDataPostRequest.md) |  | 

### Return type

[**SubprocessesDataGet200Response**](SubprocessesDataGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubprocessesDataSubprocessesDatumIdDelete

> SubprocessesData SubprocessesDataSubprocessesDatumIdDelete(ctx, subprocessesDatumId).Execute()



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
	subprocessesDatumId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubprocessesDataAPI.SubprocessesDataSubprocessesDatumIdDelete(context.Background(), subprocessesDatumId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubprocessesDataAPI.SubprocessesDataSubprocessesDatumIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubprocessesDataSubprocessesDatumIdDelete`: SubprocessesData
	fmt.Fprintf(os.Stdout, "Response from `SubprocessesDataAPI.SubprocessesDataSubprocessesDatumIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subprocessesDatumId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubprocessesDataSubprocessesDatumIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubprocessesData**](SubprocessesData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubprocessesDataSubprocessesDatumIdGet

> SubprocessesData SubprocessesDataSubprocessesDatumIdGet(ctx, subprocessesDatumId).Include(include).Execute()



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
	subprocessesDatumId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubprocessesDataAPI.SubprocessesDataSubprocessesDatumIdGet(context.Background(), subprocessesDatumId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubprocessesDataAPI.SubprocessesDataSubprocessesDatumIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubprocessesDataSubprocessesDatumIdGet`: SubprocessesData
	fmt.Fprintf(os.Stdout, "Response from `SubprocessesDataAPI.SubprocessesDataSubprocessesDatumIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subprocessesDatumId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubprocessesDataSubprocessesDatumIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**SubprocessesData**](SubprocessesData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubprocessesDataSubprocessesDatumIdPut

> SubprocessesData SubprocessesDataSubprocessesDatumIdPut(ctx, subprocessesDatumId).SubprocessesDataPut(subprocessesDataPut).Execute()



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
	subprocessesDatumId := int64(789) // int64 | Model id
	subprocessesDataPut := *openapiclient.NewSubprocessesDataPut() // SubprocessesDataPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubprocessesDataAPI.SubprocessesDataSubprocessesDatumIdPut(context.Background(), subprocessesDatumId).SubprocessesDataPut(subprocessesDataPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubprocessesDataAPI.SubprocessesDataSubprocessesDatumIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubprocessesDataSubprocessesDatumIdPut`: SubprocessesData
	fmt.Fprintf(os.Stdout, "Response from `SubprocessesDataAPI.SubprocessesDataSubprocessesDatumIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subprocessesDatumId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubprocessesDataSubprocessesDatumIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subprocessesDataPut** | [**SubprocessesDataPut**](SubprocessesDataPut.md) |  | 

### Return type

[**SubprocessesData**](SubprocessesData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

