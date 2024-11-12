# \SubprocessesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubprocessesGet**](SubprocessesAPI.md#SubprocessesGet) | **Get** /subprocesses | 
[**SubprocessesPost**](SubprocessesAPI.md#SubprocessesPost) | **Post** /subprocesses | 
[**SubprocessesSubprocessIdDelete**](SubprocessesAPI.md#SubprocessesSubprocessIdDelete) | **Delete** /subprocesses/{subprocess_id} | 
[**SubprocessesSubprocessIdGet**](SubprocessesAPI.md#SubprocessesSubprocessIdGet) | **Get** /subprocesses/{subprocess_id} | 
[**SubprocessesSubprocessIdPut**](SubprocessesAPI.md#SubprocessesSubprocessIdPut) | **Put** /subprocesses/{subprocess_id} | 



## SubprocessesGet

> SubprocessesGet200Response SubprocessesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.SubprocessesAPI.SubprocessesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubprocessesAPI.SubprocessesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubprocessesGet`: SubprocessesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SubprocessesAPI.SubprocessesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubprocessesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**SubprocessesGet200Response**](SubprocessesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubprocessesPost

> SubprocessesGet200Response SubprocessesPost(ctx).SubprocessesPostRequest(subprocessesPostRequest).Execute()



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
	subprocessesPostRequest := *openapiclient.NewSubprocessesPostRequest() // SubprocessesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubprocessesAPI.SubprocessesPost(context.Background()).SubprocessesPostRequest(subprocessesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubprocessesAPI.SubprocessesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubprocessesPost`: SubprocessesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SubprocessesAPI.SubprocessesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubprocessesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subprocessesPostRequest** | [**SubprocessesPostRequest**](SubprocessesPostRequest.md) |  | 

### Return type

[**SubprocessesGet200Response**](SubprocessesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubprocessesSubprocessIdDelete

> Subprocesses SubprocessesSubprocessIdDelete(ctx, subprocessId).Execute()



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
	subprocessId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubprocessesAPI.SubprocessesSubprocessIdDelete(context.Background(), subprocessId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubprocessesAPI.SubprocessesSubprocessIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubprocessesSubprocessIdDelete`: Subprocesses
	fmt.Fprintf(os.Stdout, "Response from `SubprocessesAPI.SubprocessesSubprocessIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subprocessId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubprocessesSubprocessIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subprocesses**](Subprocesses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubprocessesSubprocessIdGet

> Subprocesses SubprocessesSubprocessIdGet(ctx, subprocessId).Include(include).Execute()



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
	subprocessId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubprocessesAPI.SubprocessesSubprocessIdGet(context.Background(), subprocessId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubprocessesAPI.SubprocessesSubprocessIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubprocessesSubprocessIdGet`: Subprocesses
	fmt.Fprintf(os.Stdout, "Response from `SubprocessesAPI.SubprocessesSubprocessIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subprocessId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubprocessesSubprocessIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**Subprocesses**](Subprocesses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubprocessesSubprocessIdPut

> Subprocesses SubprocessesSubprocessIdPut(ctx, subprocessId).SubprocessesPut(subprocessesPut).Execute()



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
	subprocessId := int64(789) // int64 | Model id
	subprocessesPut := *openapiclient.NewSubprocessesPut() // SubprocessesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubprocessesAPI.SubprocessesSubprocessIdPut(context.Background(), subprocessId).SubprocessesPut(subprocessesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubprocessesAPI.SubprocessesSubprocessIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubprocessesSubprocessIdPut`: Subprocesses
	fmt.Fprintf(os.Stdout, "Response from `SubprocessesAPI.SubprocessesSubprocessIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subprocessId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubprocessesSubprocessIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subprocessesPut** | [**SubprocessesPut**](SubprocessesPut.md) |  | 

### Return type

[**Subprocesses**](Subprocesses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

