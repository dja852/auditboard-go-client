# \StatusOptionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatusOptionsGet**](StatusOptionsAPI.md#StatusOptionsGet) | **Get** /status_options | 
[**StatusOptionsPost**](StatusOptionsAPI.md#StatusOptionsPost) | **Post** /status_options | 
[**StatusOptionsStatusOptionIdDelete**](StatusOptionsAPI.md#StatusOptionsStatusOptionIdDelete) | **Delete** /status_options/{status_option_id} | 
[**StatusOptionsStatusOptionIdGet**](StatusOptionsAPI.md#StatusOptionsStatusOptionIdGet) | **Get** /status_options/{status_option_id} | 
[**StatusOptionsStatusOptionIdPut**](StatusOptionsAPI.md#StatusOptionsStatusOptionIdPut) | **Put** /status_options/{status_option_id} | 



## StatusOptionsGet

> StatusOptionsGet200Response StatusOptionsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.StatusOptionsAPI.StatusOptionsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusOptionsAPI.StatusOptionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatusOptionsGet`: StatusOptionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `StatusOptionsAPI.StatusOptionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStatusOptionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**StatusOptionsGet200Response**](StatusOptionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatusOptionsPost

> StatusOptionsGet200Response StatusOptionsPost(ctx).StatusOptionsPostRequest(statusOptionsPostRequest).Execute()



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
	statusOptionsPostRequest := *openapiclient.NewStatusOptionsPostRequest() // StatusOptionsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatusOptionsAPI.StatusOptionsPost(context.Background()).StatusOptionsPostRequest(statusOptionsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusOptionsAPI.StatusOptionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatusOptionsPost`: StatusOptionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `StatusOptionsAPI.StatusOptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStatusOptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **statusOptionsPostRequest** | [**StatusOptionsPostRequest**](StatusOptionsPostRequest.md) |  | 

### Return type

[**StatusOptionsGet200Response**](StatusOptionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatusOptionsStatusOptionIdDelete

> StatusOptions StatusOptionsStatusOptionIdDelete(ctx, statusOptionId).Execute()



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
	statusOptionId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatusOptionsAPI.StatusOptionsStatusOptionIdDelete(context.Background(), statusOptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusOptionsAPI.StatusOptionsStatusOptionIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatusOptionsStatusOptionIdDelete`: StatusOptions
	fmt.Fprintf(os.Stdout, "Response from `StatusOptionsAPI.StatusOptionsStatusOptionIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**statusOptionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatusOptionsStatusOptionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StatusOptions**](StatusOptions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatusOptionsStatusOptionIdGet

> StatusOptions StatusOptionsStatusOptionIdGet(ctx, statusOptionId).Include(include).Execute()



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
	statusOptionId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatusOptionsAPI.StatusOptionsStatusOptionIdGet(context.Background(), statusOptionId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusOptionsAPI.StatusOptionsStatusOptionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatusOptionsStatusOptionIdGet`: StatusOptions
	fmt.Fprintf(os.Stdout, "Response from `StatusOptionsAPI.StatusOptionsStatusOptionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**statusOptionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatusOptionsStatusOptionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**StatusOptions**](StatusOptions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatusOptionsStatusOptionIdPut

> StatusOptions StatusOptionsStatusOptionIdPut(ctx, statusOptionId).StatusOptionsPut(statusOptionsPut).Execute()



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
	statusOptionId := int64(789) // int64 | Model id
	statusOptionsPut := *openapiclient.NewStatusOptionsPut() // StatusOptionsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatusOptionsAPI.StatusOptionsStatusOptionIdPut(context.Background(), statusOptionId).StatusOptionsPut(statusOptionsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusOptionsAPI.StatusOptionsStatusOptionIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatusOptionsStatusOptionIdPut`: StatusOptions
	fmt.Fprintf(os.Stdout, "Response from `StatusOptionsAPI.StatusOptionsStatusOptionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**statusOptionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatusOptionsStatusOptionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **statusOptionsPut** | [**StatusOptionsPut**](StatusOptionsPut.md) |  | 

### Return type

[**StatusOptions**](StatusOptions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

