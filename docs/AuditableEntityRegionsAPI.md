# \AuditableEntityRegionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditableEntityRegionsAuditableEntityRegionIdDelete**](AuditableEntityRegionsAPI.md#AuditableEntityRegionsAuditableEntityRegionIdDelete) | **Delete** /auditable_entity_regions/{auditable_entity_region_id} | 
[**AuditableEntityRegionsAuditableEntityRegionIdGet**](AuditableEntityRegionsAPI.md#AuditableEntityRegionsAuditableEntityRegionIdGet) | **Get** /auditable_entity_regions/{auditable_entity_region_id} | 
[**AuditableEntityRegionsAuditableEntityRegionIdPut**](AuditableEntityRegionsAPI.md#AuditableEntityRegionsAuditableEntityRegionIdPut) | **Put** /auditable_entity_regions/{auditable_entity_region_id} | 
[**AuditableEntityRegionsGet**](AuditableEntityRegionsAPI.md#AuditableEntityRegionsGet) | **Get** /auditable_entity_regions | 
[**AuditableEntityRegionsPost**](AuditableEntityRegionsAPI.md#AuditableEntityRegionsPost) | **Post** /auditable_entity_regions | 



## AuditableEntityRegionsAuditableEntityRegionIdDelete

> AuditableEntityRegions AuditableEntityRegionsAuditableEntityRegionIdDelete(ctx, auditableEntityRegionId).Execute()



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
	auditableEntityRegionId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditableEntityRegionsAPI.AuditableEntityRegionsAuditableEntityRegionIdDelete(context.Background(), auditableEntityRegionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditableEntityRegionsAPI.AuditableEntityRegionsAuditableEntityRegionIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditableEntityRegionsAuditableEntityRegionIdDelete`: AuditableEntityRegions
	fmt.Fprintf(os.Stdout, "Response from `AuditableEntityRegionsAPI.AuditableEntityRegionsAuditableEntityRegionIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditableEntityRegionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditableEntityRegionsAuditableEntityRegionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuditableEntityRegions**](AuditableEntityRegions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditableEntityRegionsAuditableEntityRegionIdGet

> AuditableEntityRegions AuditableEntityRegionsAuditableEntityRegionIdGet(ctx, auditableEntityRegionId).Include(include).Execute()



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
	auditableEntityRegionId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditableEntityRegionsAPI.AuditableEntityRegionsAuditableEntityRegionIdGet(context.Background(), auditableEntityRegionId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditableEntityRegionsAPI.AuditableEntityRegionsAuditableEntityRegionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditableEntityRegionsAuditableEntityRegionIdGet`: AuditableEntityRegions
	fmt.Fprintf(os.Stdout, "Response from `AuditableEntityRegionsAPI.AuditableEntityRegionsAuditableEntityRegionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditableEntityRegionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditableEntityRegionsAuditableEntityRegionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditableEntityRegions**](AuditableEntityRegions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditableEntityRegionsAuditableEntityRegionIdPut

> AuditableEntityRegions AuditableEntityRegionsAuditableEntityRegionIdPut(ctx, auditableEntityRegionId).AuditableEntityRegionsPut(auditableEntityRegionsPut).Execute()



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
	auditableEntityRegionId := int64(789) // int64 | Model id
	auditableEntityRegionsPut := *openapiclient.NewAuditableEntityRegionsPut() // AuditableEntityRegionsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditableEntityRegionsAPI.AuditableEntityRegionsAuditableEntityRegionIdPut(context.Background(), auditableEntityRegionId).AuditableEntityRegionsPut(auditableEntityRegionsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditableEntityRegionsAPI.AuditableEntityRegionsAuditableEntityRegionIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditableEntityRegionsAuditableEntityRegionIdPut`: AuditableEntityRegions
	fmt.Fprintf(os.Stdout, "Response from `AuditableEntityRegionsAPI.AuditableEntityRegionsAuditableEntityRegionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditableEntityRegionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditableEntityRegionsAuditableEntityRegionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auditableEntityRegionsPut** | [**AuditableEntityRegionsPut**](AuditableEntityRegionsPut.md) |  | 

### Return type

[**AuditableEntityRegions**](AuditableEntityRegions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditableEntityRegionsGet

> AuditableEntityRegionsGet200Response AuditableEntityRegionsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.AuditableEntityRegionsAPI.AuditableEntityRegionsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditableEntityRegionsAPI.AuditableEntityRegionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditableEntityRegionsGet`: AuditableEntityRegionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditableEntityRegionsAPI.AuditableEntityRegionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditableEntityRegionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditableEntityRegionsGet200Response**](AuditableEntityRegionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditableEntityRegionsPost

> AuditableEntityRegionsGet200Response AuditableEntityRegionsPost(ctx).AuditableEntityRegionsPostRequest(auditableEntityRegionsPostRequest).Execute()



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
	auditableEntityRegionsPostRequest := *openapiclient.NewAuditableEntityRegionsPostRequest() // AuditableEntityRegionsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditableEntityRegionsAPI.AuditableEntityRegionsPost(context.Background()).AuditableEntityRegionsPostRequest(auditableEntityRegionsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditableEntityRegionsAPI.AuditableEntityRegionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditableEntityRegionsPost`: AuditableEntityRegionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditableEntityRegionsAPI.AuditableEntityRegionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditableEntityRegionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auditableEntityRegionsPostRequest** | [**AuditableEntityRegionsPostRequest**](AuditableEntityRegionsPostRequest.md) |  | 

### Return type

[**AuditableEntityRegionsGet200Response**](AuditableEntityRegionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

