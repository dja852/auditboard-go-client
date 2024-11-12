# \LocationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocationsGet**](LocationsAPI.md#LocationsGet) | **Get** /locations | 
[**LocationsLocationIdDelete**](LocationsAPI.md#LocationsLocationIdDelete) | **Delete** /locations/{location_id} | 
[**LocationsLocationIdGet**](LocationsAPI.md#LocationsLocationIdGet) | **Get** /locations/{location_id} | 
[**LocationsLocationIdPut**](LocationsAPI.md#LocationsLocationIdPut) | **Put** /locations/{location_id} | 
[**LocationsPost**](LocationsAPI.md#LocationsPost) | **Post** /locations | 



## LocationsGet

> LocationsGet200Response LocationsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.LocationsAPI.LocationsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationsAPI.LocationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LocationsGet`: LocationsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `LocationsAPI.LocationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLocationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**LocationsGet200Response**](LocationsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocationsLocationIdDelete

> Locations LocationsLocationIdDelete(ctx, locationId).Execute()



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
	locationId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationsAPI.LocationsLocationIdDelete(context.Background(), locationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationsAPI.LocationsLocationIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LocationsLocationIdDelete`: Locations
	fmt.Fprintf(os.Stdout, "Response from `LocationsAPI.LocationsLocationIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLocationsLocationIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Locations**](Locations.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocationsLocationIdGet

> Locations LocationsLocationIdGet(ctx, locationId).Include(include).Execute()



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
	locationId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationsAPI.LocationsLocationIdGet(context.Background(), locationId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationsAPI.LocationsLocationIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LocationsLocationIdGet`: Locations
	fmt.Fprintf(os.Stdout, "Response from `LocationsAPI.LocationsLocationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLocationsLocationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**Locations**](Locations.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocationsLocationIdPut

> Locations LocationsLocationIdPut(ctx, locationId).LocationsPut(locationsPut).Execute()



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
	locationId := int64(789) // int64 | Model id
	locationsPut := *openapiclient.NewLocationsPut() // LocationsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationsAPI.LocationsLocationIdPut(context.Background(), locationId).LocationsPut(locationsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationsAPI.LocationsLocationIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LocationsLocationIdPut`: Locations
	fmt.Fprintf(os.Stdout, "Response from `LocationsAPI.LocationsLocationIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLocationsLocationIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locationsPut** | [**LocationsPut**](LocationsPut.md) |  | 

### Return type

[**Locations**](Locations.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocationsPost

> LocationsGet200Response LocationsPost(ctx).LocationsPostRequest(locationsPostRequest).Execute()



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
	locationsPostRequest := *openapiclient.NewLocationsPostRequest() // LocationsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationsAPI.LocationsPost(context.Background()).LocationsPostRequest(locationsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationsAPI.LocationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LocationsPost`: LocationsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `LocationsAPI.LocationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLocationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locationsPostRequest** | [**LocationsPostRequest**](LocationsPostRequest.md) |  | 

### Return type

[**LocationsGet200Response**](LocationsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

