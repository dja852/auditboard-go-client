# \RegionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegionsGet**](RegionsAPI.md#RegionsGet) | **Get** /regions | 
[**RegionsPost**](RegionsAPI.md#RegionsPost) | **Post** /regions | 
[**RegionsRegionIdDelete**](RegionsAPI.md#RegionsRegionIdDelete) | **Delete** /regions/{region_id} | 
[**RegionsRegionIdGet**](RegionsAPI.md#RegionsRegionIdGet) | **Get** /regions/{region_id} | 
[**RegionsRegionIdPut**](RegionsAPI.md#RegionsRegionIdPut) | **Put** /regions/{region_id} | 



## RegionsGet

> RegionsGet200Response RegionsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.RegionsAPI.RegionsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionsAPI.RegionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegionsGet`: RegionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RegionsAPI.RegionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**RegionsGet200Response**](RegionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegionsPost

> RegionsGet200Response RegionsPost(ctx).RegionsPostRequest(regionsPostRequest).Execute()



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
	regionsPostRequest := *openapiclient.NewRegionsPostRequest() // RegionsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionsAPI.RegionsPost(context.Background()).RegionsPostRequest(regionsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionsAPI.RegionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegionsPost`: RegionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RegionsAPI.RegionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionsPostRequest** | [**RegionsPostRequest**](RegionsPostRequest.md) |  | 

### Return type

[**RegionsGet200Response**](RegionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegionsRegionIdDelete

> Regions RegionsRegionIdDelete(ctx, regionId).Execute()



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
	regionId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionsAPI.RegionsRegionIdDelete(context.Background(), regionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionsAPI.RegionsRegionIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegionsRegionIdDelete`: Regions
	fmt.Fprintf(os.Stdout, "Response from `RegionsAPI.RegionsRegionIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegionsRegionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Regions**](Regions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegionsRegionIdGet

> Regions RegionsRegionIdGet(ctx, regionId).Include(include).Execute()



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
	regionId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionsAPI.RegionsRegionIdGet(context.Background(), regionId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionsAPI.RegionsRegionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegionsRegionIdGet`: Regions
	fmt.Fprintf(os.Stdout, "Response from `RegionsAPI.RegionsRegionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegionsRegionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**Regions**](Regions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegionsRegionIdPut

> Regions RegionsRegionIdPut(ctx, regionId).RegionsPut(regionsPut).Execute()



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
	regionId := int64(789) // int64 | Model id
	regionsPut := *openapiclient.NewRegionsPut() // RegionsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionsAPI.RegionsRegionIdPut(context.Background(), regionId).RegionsPut(regionsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionsAPI.RegionsRegionIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegionsRegionIdPut`: Regions
	fmt.Fprintf(os.Stdout, "Response from `RegionsAPI.RegionsRegionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegionsRegionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **regionsPut** | [**RegionsPut**](RegionsPut.md) |  | 

### Return type

[**Regions**](Regions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

