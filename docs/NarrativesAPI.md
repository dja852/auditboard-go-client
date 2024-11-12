# \NarrativesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NarrativesGet**](NarrativesAPI.md#NarrativesGet) | **Get** /narratives | 
[**NarrativesNarrativeIdDelete**](NarrativesAPI.md#NarrativesNarrativeIdDelete) | **Delete** /narratives/{narrative_id} | 
[**NarrativesNarrativeIdGet**](NarrativesAPI.md#NarrativesNarrativeIdGet) | **Get** /narratives/{narrative_id} | 
[**NarrativesNarrativeIdPut**](NarrativesAPI.md#NarrativesNarrativeIdPut) | **Put** /narratives/{narrative_id} | 
[**NarrativesPost**](NarrativesAPI.md#NarrativesPost) | **Post** /narratives | 



## NarrativesGet

> NarrativesGet200Response NarrativesGet(ctx).Include(include).Execute()



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
	include := []string{"Include_example"} // []string | Get Attachments or m2m models. To use both query parameters, they must be written as: `?include[]=attachments&include[]=m2m`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NarrativesAPI.NarrativesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NarrativesAPI.NarrativesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NarrativesGet`: NarrativesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `NarrativesAPI.NarrativesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNarrativesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Get Attachments or m2m models. To use both query parameters, they must be written as: &#x60;?include[]&#x3D;attachments&amp;include[]&#x3D;m2m&#x60;. | 

### Return type

[**NarrativesGet200Response**](NarrativesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NarrativesNarrativeIdDelete

> Narratives NarrativesNarrativeIdDelete(ctx, narrativeId).Execute()



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
	narrativeId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NarrativesAPI.NarrativesNarrativeIdDelete(context.Background(), narrativeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NarrativesAPI.NarrativesNarrativeIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NarrativesNarrativeIdDelete`: Narratives
	fmt.Fprintf(os.Stdout, "Response from `NarrativesAPI.NarrativesNarrativeIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**narrativeId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNarrativesNarrativeIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Narratives**](Narratives.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NarrativesNarrativeIdGet

> Narratives NarrativesNarrativeIdGet(ctx, narrativeId).Include(include).Execute()



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
	narrativeId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NarrativesAPI.NarrativesNarrativeIdGet(context.Background(), narrativeId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NarrativesAPI.NarrativesNarrativeIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NarrativesNarrativeIdGet`: Narratives
	fmt.Fprintf(os.Stdout, "Response from `NarrativesAPI.NarrativesNarrativeIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**narrativeId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNarrativesNarrativeIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**Narratives**](Narratives.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NarrativesNarrativeIdPut

> Narratives NarrativesNarrativeIdPut(ctx, narrativeId).NarrativesPut(narrativesPut).Execute()



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
	narrativeId := int64(789) // int64 | Model id
	narrativesPut := *openapiclient.NewNarrativesPut() // NarrativesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NarrativesAPI.NarrativesNarrativeIdPut(context.Background(), narrativeId).NarrativesPut(narrativesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NarrativesAPI.NarrativesNarrativeIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NarrativesNarrativeIdPut`: Narratives
	fmt.Fprintf(os.Stdout, "Response from `NarrativesAPI.NarrativesNarrativeIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**narrativeId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNarrativesNarrativeIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **narrativesPut** | [**NarrativesPut**](NarrativesPut.md) |  | 

### Return type

[**Narratives**](Narratives.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NarrativesPost

> NarrativesGet200Response NarrativesPost(ctx).NarrativesPostRequest(narrativesPostRequest).Execute()



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
	narrativesPostRequest := *openapiclient.NewNarrativesPostRequest() // NarrativesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NarrativesAPI.NarrativesPost(context.Background()).NarrativesPostRequest(narrativesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NarrativesAPI.NarrativesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NarrativesPost`: NarrativesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `NarrativesAPI.NarrativesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNarrativesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **narrativesPostRequest** | [**NarrativesPostRequest**](NarrativesPostRequest.md) |  | 

### Return type

[**NarrativesGet200Response**](NarrativesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

