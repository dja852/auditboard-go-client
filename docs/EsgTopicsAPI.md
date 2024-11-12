# \EsgTopicsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EsgTopicsEsgTopicIdDelete**](EsgTopicsAPI.md#EsgTopicsEsgTopicIdDelete) | **Delete** /esg_topics/{esg_topic_id} | 
[**EsgTopicsEsgTopicIdGet**](EsgTopicsAPI.md#EsgTopicsEsgTopicIdGet) | **Get** /esg_topics/{esg_topic_id} | 
[**EsgTopicsEsgTopicIdPut**](EsgTopicsAPI.md#EsgTopicsEsgTopicIdPut) | **Put** /esg_topics/{esg_topic_id} | 
[**EsgTopicsGet**](EsgTopicsAPI.md#EsgTopicsGet) | **Get** /esg_topics | 
[**EsgTopicsPost**](EsgTopicsAPI.md#EsgTopicsPost) | **Post** /esg_topics | 



## EsgTopicsEsgTopicIdDelete

> EsgTopics EsgTopicsEsgTopicIdDelete(ctx, esgTopicId).Execute()



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
	esgTopicId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgTopicsAPI.EsgTopicsEsgTopicIdDelete(context.Background(), esgTopicId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgTopicsAPI.EsgTopicsEsgTopicIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgTopicsEsgTopicIdDelete`: EsgTopics
	fmt.Fprintf(os.Stdout, "Response from `EsgTopicsAPI.EsgTopicsEsgTopicIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgTopicId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgTopicsEsgTopicIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EsgTopics**](EsgTopics.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgTopicsEsgTopicIdGet

> EsgTopics EsgTopicsEsgTopicIdGet(ctx, esgTopicId).Include(include).Execute()



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
	esgTopicId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgTopicsAPI.EsgTopicsEsgTopicIdGet(context.Background(), esgTopicId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgTopicsAPI.EsgTopicsEsgTopicIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgTopicsEsgTopicIdGet`: EsgTopics
	fmt.Fprintf(os.Stdout, "Response from `EsgTopicsAPI.EsgTopicsEsgTopicIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgTopicId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgTopicsEsgTopicIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EsgTopics**](EsgTopics.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgTopicsEsgTopicIdPut

> EsgTopics EsgTopicsEsgTopicIdPut(ctx, esgTopicId).EsgTopicsPut(esgTopicsPut).Execute()



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
	esgTopicId := int64(789) // int64 | Model id
	esgTopicsPut := *openapiclient.NewEsgTopicsPut() // EsgTopicsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgTopicsAPI.EsgTopicsEsgTopicIdPut(context.Background(), esgTopicId).EsgTopicsPut(esgTopicsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgTopicsAPI.EsgTopicsEsgTopicIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgTopicsEsgTopicIdPut`: EsgTopics
	fmt.Fprintf(os.Stdout, "Response from `EsgTopicsAPI.EsgTopicsEsgTopicIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgTopicId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgTopicsEsgTopicIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **esgTopicsPut** | [**EsgTopicsPut**](EsgTopicsPut.md) |  | 

### Return type

[**EsgTopics**](EsgTopics.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgTopicsGet

> EsgTopicsGet200Response EsgTopicsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.EsgTopicsAPI.EsgTopicsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgTopicsAPI.EsgTopicsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgTopicsGet`: EsgTopicsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EsgTopicsAPI.EsgTopicsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsgTopicsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EsgTopicsGet200Response**](EsgTopicsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgTopicsPost

> EsgTopicsGet200Response EsgTopicsPost(ctx).EsgTopicsPostRequest(esgTopicsPostRequest).Execute()



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
	esgTopicsPostRequest := *openapiclient.NewEsgTopicsPostRequest() // EsgTopicsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgTopicsAPI.EsgTopicsPost(context.Background()).EsgTopicsPostRequest(esgTopicsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgTopicsAPI.EsgTopicsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgTopicsPost`: EsgTopicsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EsgTopicsAPI.EsgTopicsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsgTopicsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **esgTopicsPostRequest** | [**EsgTopicsPostRequest**](EsgTopicsPostRequest.md) |  | 

### Return type

[**EsgTopicsGet200Response**](EsgTopicsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

