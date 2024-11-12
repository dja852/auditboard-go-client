# \AuditBusinessSegmentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditBusinessSegmentsAuditBusinessSegmentIdDelete**](AuditBusinessSegmentsAPI.md#AuditBusinessSegmentsAuditBusinessSegmentIdDelete) | **Delete** /audit_business_segments/{audit_business_segment_id} | 
[**AuditBusinessSegmentsAuditBusinessSegmentIdGet**](AuditBusinessSegmentsAPI.md#AuditBusinessSegmentsAuditBusinessSegmentIdGet) | **Get** /audit_business_segments/{audit_business_segment_id} | 
[**AuditBusinessSegmentsAuditBusinessSegmentIdPut**](AuditBusinessSegmentsAPI.md#AuditBusinessSegmentsAuditBusinessSegmentIdPut) | **Put** /audit_business_segments/{audit_business_segment_id} | 
[**AuditBusinessSegmentsGet**](AuditBusinessSegmentsAPI.md#AuditBusinessSegmentsGet) | **Get** /audit_business_segments | 
[**AuditBusinessSegmentsPost**](AuditBusinessSegmentsAPI.md#AuditBusinessSegmentsPost) | **Post** /audit_business_segments | 



## AuditBusinessSegmentsAuditBusinessSegmentIdDelete

> AuditBusinessSegments AuditBusinessSegmentsAuditBusinessSegmentIdDelete(ctx, auditBusinessSegmentId).Execute()



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
	auditBusinessSegmentId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditBusinessSegmentsAPI.AuditBusinessSegmentsAuditBusinessSegmentIdDelete(context.Background(), auditBusinessSegmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditBusinessSegmentsAPI.AuditBusinessSegmentsAuditBusinessSegmentIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditBusinessSegmentsAuditBusinessSegmentIdDelete`: AuditBusinessSegments
	fmt.Fprintf(os.Stdout, "Response from `AuditBusinessSegmentsAPI.AuditBusinessSegmentsAuditBusinessSegmentIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditBusinessSegmentId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditBusinessSegmentsAuditBusinessSegmentIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuditBusinessSegments**](AuditBusinessSegments.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditBusinessSegmentsAuditBusinessSegmentIdGet

> AuditBusinessSegments AuditBusinessSegmentsAuditBusinessSegmentIdGet(ctx, auditBusinessSegmentId).Include(include).Execute()



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
	auditBusinessSegmentId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditBusinessSegmentsAPI.AuditBusinessSegmentsAuditBusinessSegmentIdGet(context.Background(), auditBusinessSegmentId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditBusinessSegmentsAPI.AuditBusinessSegmentsAuditBusinessSegmentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditBusinessSegmentsAuditBusinessSegmentIdGet`: AuditBusinessSegments
	fmt.Fprintf(os.Stdout, "Response from `AuditBusinessSegmentsAPI.AuditBusinessSegmentsAuditBusinessSegmentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditBusinessSegmentId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditBusinessSegmentsAuditBusinessSegmentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditBusinessSegments**](AuditBusinessSegments.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditBusinessSegmentsAuditBusinessSegmentIdPut

> AuditBusinessSegments AuditBusinessSegmentsAuditBusinessSegmentIdPut(ctx, auditBusinessSegmentId).AuditBusinessSegmentsPut(auditBusinessSegmentsPut).Execute()



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
	auditBusinessSegmentId := int64(789) // int64 | Model id
	auditBusinessSegmentsPut := *openapiclient.NewAuditBusinessSegmentsPut() // AuditBusinessSegmentsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditBusinessSegmentsAPI.AuditBusinessSegmentsAuditBusinessSegmentIdPut(context.Background(), auditBusinessSegmentId).AuditBusinessSegmentsPut(auditBusinessSegmentsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditBusinessSegmentsAPI.AuditBusinessSegmentsAuditBusinessSegmentIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditBusinessSegmentsAuditBusinessSegmentIdPut`: AuditBusinessSegments
	fmt.Fprintf(os.Stdout, "Response from `AuditBusinessSegmentsAPI.AuditBusinessSegmentsAuditBusinessSegmentIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditBusinessSegmentId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditBusinessSegmentsAuditBusinessSegmentIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auditBusinessSegmentsPut** | [**AuditBusinessSegmentsPut**](AuditBusinessSegmentsPut.md) |  | 

### Return type

[**AuditBusinessSegments**](AuditBusinessSegments.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditBusinessSegmentsGet

> AuditBusinessSegmentsGet200Response AuditBusinessSegmentsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.AuditBusinessSegmentsAPI.AuditBusinessSegmentsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditBusinessSegmentsAPI.AuditBusinessSegmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditBusinessSegmentsGet`: AuditBusinessSegmentsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditBusinessSegmentsAPI.AuditBusinessSegmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditBusinessSegmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditBusinessSegmentsGet200Response**](AuditBusinessSegmentsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditBusinessSegmentsPost

> AuditBusinessSegmentsGet200Response AuditBusinessSegmentsPost(ctx).AuditBusinessSegmentsPostRequest(auditBusinessSegmentsPostRequest).Execute()



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
	auditBusinessSegmentsPostRequest := *openapiclient.NewAuditBusinessSegmentsPostRequest() // AuditBusinessSegmentsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditBusinessSegmentsAPI.AuditBusinessSegmentsPost(context.Background()).AuditBusinessSegmentsPostRequest(auditBusinessSegmentsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditBusinessSegmentsAPI.AuditBusinessSegmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditBusinessSegmentsPost`: AuditBusinessSegmentsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditBusinessSegmentsAPI.AuditBusinessSegmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditBusinessSegmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auditBusinessSegmentsPostRequest** | [**AuditBusinessSegmentsPostRequest**](AuditBusinessSegmentsPostRequest.md) |  | 

### Return type

[**AuditBusinessSegmentsGet200Response**](AuditBusinessSegmentsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

