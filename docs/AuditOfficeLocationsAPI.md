# \AuditOfficeLocationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditOfficeLocationsAuditOfficeLocationIdDelete**](AuditOfficeLocationsAPI.md#AuditOfficeLocationsAuditOfficeLocationIdDelete) | **Delete** /audit_office_locations/{audit_office_location_id} | 
[**AuditOfficeLocationsAuditOfficeLocationIdGet**](AuditOfficeLocationsAPI.md#AuditOfficeLocationsAuditOfficeLocationIdGet) | **Get** /audit_office_locations/{audit_office_location_id} | 
[**AuditOfficeLocationsAuditOfficeLocationIdPut**](AuditOfficeLocationsAPI.md#AuditOfficeLocationsAuditOfficeLocationIdPut) | **Put** /audit_office_locations/{audit_office_location_id} | 
[**AuditOfficeLocationsGet**](AuditOfficeLocationsAPI.md#AuditOfficeLocationsGet) | **Get** /audit_office_locations | 
[**AuditOfficeLocationsPost**](AuditOfficeLocationsAPI.md#AuditOfficeLocationsPost) | **Post** /audit_office_locations | 



## AuditOfficeLocationsAuditOfficeLocationIdDelete

> AuditOfficeLocations AuditOfficeLocationsAuditOfficeLocationIdDelete(ctx, auditOfficeLocationId).Execute()



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
	auditOfficeLocationId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditOfficeLocationsAPI.AuditOfficeLocationsAuditOfficeLocationIdDelete(context.Background(), auditOfficeLocationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditOfficeLocationsAPI.AuditOfficeLocationsAuditOfficeLocationIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditOfficeLocationsAuditOfficeLocationIdDelete`: AuditOfficeLocations
	fmt.Fprintf(os.Stdout, "Response from `AuditOfficeLocationsAPI.AuditOfficeLocationsAuditOfficeLocationIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditOfficeLocationId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditOfficeLocationsAuditOfficeLocationIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuditOfficeLocations**](AuditOfficeLocations.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditOfficeLocationsAuditOfficeLocationIdGet

> AuditOfficeLocations AuditOfficeLocationsAuditOfficeLocationIdGet(ctx, auditOfficeLocationId).Include(include).Execute()



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
	auditOfficeLocationId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditOfficeLocationsAPI.AuditOfficeLocationsAuditOfficeLocationIdGet(context.Background(), auditOfficeLocationId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditOfficeLocationsAPI.AuditOfficeLocationsAuditOfficeLocationIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditOfficeLocationsAuditOfficeLocationIdGet`: AuditOfficeLocations
	fmt.Fprintf(os.Stdout, "Response from `AuditOfficeLocationsAPI.AuditOfficeLocationsAuditOfficeLocationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditOfficeLocationId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditOfficeLocationsAuditOfficeLocationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditOfficeLocations**](AuditOfficeLocations.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditOfficeLocationsAuditOfficeLocationIdPut

> AuditOfficeLocations AuditOfficeLocationsAuditOfficeLocationIdPut(ctx, auditOfficeLocationId).AuditOfficeLocationsPut(auditOfficeLocationsPut).Execute()



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
	auditOfficeLocationId := int64(789) // int64 | Model id
	auditOfficeLocationsPut := *openapiclient.NewAuditOfficeLocationsPut() // AuditOfficeLocationsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditOfficeLocationsAPI.AuditOfficeLocationsAuditOfficeLocationIdPut(context.Background(), auditOfficeLocationId).AuditOfficeLocationsPut(auditOfficeLocationsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditOfficeLocationsAPI.AuditOfficeLocationsAuditOfficeLocationIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditOfficeLocationsAuditOfficeLocationIdPut`: AuditOfficeLocations
	fmt.Fprintf(os.Stdout, "Response from `AuditOfficeLocationsAPI.AuditOfficeLocationsAuditOfficeLocationIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditOfficeLocationId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditOfficeLocationsAuditOfficeLocationIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auditOfficeLocationsPut** | [**AuditOfficeLocationsPut**](AuditOfficeLocationsPut.md) |  | 

### Return type

[**AuditOfficeLocations**](AuditOfficeLocations.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditOfficeLocationsGet

> AuditOfficeLocationsGet200Response AuditOfficeLocationsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.AuditOfficeLocationsAPI.AuditOfficeLocationsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditOfficeLocationsAPI.AuditOfficeLocationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditOfficeLocationsGet`: AuditOfficeLocationsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditOfficeLocationsAPI.AuditOfficeLocationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditOfficeLocationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditOfficeLocationsGet200Response**](AuditOfficeLocationsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditOfficeLocationsPost

> AuditOfficeLocationsGet200Response AuditOfficeLocationsPost(ctx).AuditOfficeLocationsPostRequest(auditOfficeLocationsPostRequest).Execute()



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
	auditOfficeLocationsPostRequest := *openapiclient.NewAuditOfficeLocationsPostRequest() // AuditOfficeLocationsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditOfficeLocationsAPI.AuditOfficeLocationsPost(context.Background()).AuditOfficeLocationsPostRequest(auditOfficeLocationsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditOfficeLocationsAPI.AuditOfficeLocationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditOfficeLocationsPost`: AuditOfficeLocationsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditOfficeLocationsAPI.AuditOfficeLocationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditOfficeLocationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auditOfficeLocationsPostRequest** | [**AuditOfficeLocationsPostRequest**](AuditOfficeLocationsPostRequest.md) |  | 

### Return type

[**AuditOfficeLocationsGet200Response**](AuditOfficeLocationsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

