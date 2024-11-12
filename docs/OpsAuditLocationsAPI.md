# \OpsAuditLocationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpsAuditLocationsGet**](OpsAuditLocationsAPI.md#OpsAuditLocationsGet) | **Get** /ops_audit_locations | 
[**OpsAuditLocationsOpsAuditLocationIdDelete**](OpsAuditLocationsAPI.md#OpsAuditLocationsOpsAuditLocationIdDelete) | **Delete** /ops_audit_locations/{ops_audit_location_id} | 
[**OpsAuditLocationsOpsAuditLocationIdGet**](OpsAuditLocationsAPI.md#OpsAuditLocationsOpsAuditLocationIdGet) | **Get** /ops_audit_locations/{ops_audit_location_id} | 
[**OpsAuditLocationsOpsAuditLocationIdPut**](OpsAuditLocationsAPI.md#OpsAuditLocationsOpsAuditLocationIdPut) | **Put** /ops_audit_locations/{ops_audit_location_id} | 
[**OpsAuditLocationsPost**](OpsAuditLocationsAPI.md#OpsAuditLocationsPost) | **Post** /ops_audit_locations | 



## OpsAuditLocationsGet

> OpsAuditLocationsGet200Response OpsAuditLocationsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.OpsAuditLocationsAPI.OpsAuditLocationsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditLocationsAPI.OpsAuditLocationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditLocationsGet`: OpsAuditLocationsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditLocationsAPI.OpsAuditLocationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditLocationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**OpsAuditLocationsGet200Response**](OpsAuditLocationsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditLocationsOpsAuditLocationIdDelete

> OpsAuditLocations OpsAuditLocationsOpsAuditLocationIdDelete(ctx, opsAuditLocationId).Execute()



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
	opsAuditLocationId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditLocationsAPI.OpsAuditLocationsOpsAuditLocationIdDelete(context.Background(), opsAuditLocationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditLocationsAPI.OpsAuditLocationsOpsAuditLocationIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditLocationsOpsAuditLocationIdDelete`: OpsAuditLocations
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditLocationsAPI.OpsAuditLocationsOpsAuditLocationIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditLocationId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditLocationsOpsAuditLocationIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpsAuditLocations**](OpsAuditLocations.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditLocationsOpsAuditLocationIdGet

> OpsAuditLocations OpsAuditLocationsOpsAuditLocationIdGet(ctx, opsAuditLocationId).Include(include).Execute()



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
	opsAuditLocationId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditLocationsAPI.OpsAuditLocationsOpsAuditLocationIdGet(context.Background(), opsAuditLocationId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditLocationsAPI.OpsAuditLocationsOpsAuditLocationIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditLocationsOpsAuditLocationIdGet`: OpsAuditLocations
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditLocationsAPI.OpsAuditLocationsOpsAuditLocationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditLocationId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditLocationsOpsAuditLocationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**OpsAuditLocations**](OpsAuditLocations.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditLocationsOpsAuditLocationIdPut

> OpsAuditLocations OpsAuditLocationsOpsAuditLocationIdPut(ctx, opsAuditLocationId).OpsAuditLocationsPut(opsAuditLocationsPut).Execute()



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
	opsAuditLocationId := int64(789) // int64 | Model id
	opsAuditLocationsPut := *openapiclient.NewOpsAuditLocationsPut() // OpsAuditLocationsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditLocationsAPI.OpsAuditLocationsOpsAuditLocationIdPut(context.Background(), opsAuditLocationId).OpsAuditLocationsPut(opsAuditLocationsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditLocationsAPI.OpsAuditLocationsOpsAuditLocationIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditLocationsOpsAuditLocationIdPut`: OpsAuditLocations
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditLocationsAPI.OpsAuditLocationsOpsAuditLocationIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditLocationId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditLocationsOpsAuditLocationIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opsAuditLocationsPut** | [**OpsAuditLocationsPut**](OpsAuditLocationsPut.md) |  | 

### Return type

[**OpsAuditLocations**](OpsAuditLocations.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditLocationsPost

> OpsAuditLocationsGet200Response OpsAuditLocationsPost(ctx).OpsAuditLocationsPostRequest(opsAuditLocationsPostRequest).Execute()



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
	opsAuditLocationsPostRequest := *openapiclient.NewOpsAuditLocationsPostRequest() // OpsAuditLocationsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditLocationsAPI.OpsAuditLocationsPost(context.Background()).OpsAuditLocationsPostRequest(opsAuditLocationsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditLocationsAPI.OpsAuditLocationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditLocationsPost`: OpsAuditLocationsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditLocationsAPI.OpsAuditLocationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditLocationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opsAuditLocationsPostRequest** | [**OpsAuditLocationsPostRequest**](OpsAuditLocationsPostRequest.md) |  | 

### Return type

[**OpsAuditLocationsGet200Response**](OpsAuditLocationsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

