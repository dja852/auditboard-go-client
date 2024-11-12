# \OpsAuditSubsectionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpsAuditSubsectionsGet**](OpsAuditSubsectionsAPI.md#OpsAuditSubsectionsGet) | **Get** /ops_audit_subsections | 
[**OpsAuditSubsectionsOpsAuditSubsectionIdDelete**](OpsAuditSubsectionsAPI.md#OpsAuditSubsectionsOpsAuditSubsectionIdDelete) | **Delete** /ops_audit_subsections/{ops_audit_subsection_id} | 
[**OpsAuditSubsectionsOpsAuditSubsectionIdGet**](OpsAuditSubsectionsAPI.md#OpsAuditSubsectionsOpsAuditSubsectionIdGet) | **Get** /ops_audit_subsections/{ops_audit_subsection_id} | 
[**OpsAuditSubsectionsOpsAuditSubsectionIdPut**](OpsAuditSubsectionsAPI.md#OpsAuditSubsectionsOpsAuditSubsectionIdPut) | **Put** /ops_audit_subsections/{ops_audit_subsection_id} | 
[**OpsAuditSubsectionsPost**](OpsAuditSubsectionsAPI.md#OpsAuditSubsectionsPost) | **Post** /ops_audit_subsections | 



## OpsAuditSubsectionsGet

> OpsAuditSubsectionsGet200Response OpsAuditSubsectionsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.OpsAuditSubsectionsAPI.OpsAuditSubsectionsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditSubsectionsAPI.OpsAuditSubsectionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditSubsectionsGet`: OpsAuditSubsectionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditSubsectionsAPI.OpsAuditSubsectionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditSubsectionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**OpsAuditSubsectionsGet200Response**](OpsAuditSubsectionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditSubsectionsOpsAuditSubsectionIdDelete

> OpsAuditSubsections OpsAuditSubsectionsOpsAuditSubsectionIdDelete(ctx, opsAuditSubsectionId).Execute()



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
	opsAuditSubsectionId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditSubsectionsAPI.OpsAuditSubsectionsOpsAuditSubsectionIdDelete(context.Background(), opsAuditSubsectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditSubsectionsAPI.OpsAuditSubsectionsOpsAuditSubsectionIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditSubsectionsOpsAuditSubsectionIdDelete`: OpsAuditSubsections
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditSubsectionsAPI.OpsAuditSubsectionsOpsAuditSubsectionIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditSubsectionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditSubsectionsOpsAuditSubsectionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpsAuditSubsections**](OpsAuditSubsections.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditSubsectionsOpsAuditSubsectionIdGet

> OpsAuditSubsections OpsAuditSubsectionsOpsAuditSubsectionIdGet(ctx, opsAuditSubsectionId).Include(include).Execute()



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
	opsAuditSubsectionId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditSubsectionsAPI.OpsAuditSubsectionsOpsAuditSubsectionIdGet(context.Background(), opsAuditSubsectionId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditSubsectionsAPI.OpsAuditSubsectionsOpsAuditSubsectionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditSubsectionsOpsAuditSubsectionIdGet`: OpsAuditSubsections
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditSubsectionsAPI.OpsAuditSubsectionsOpsAuditSubsectionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditSubsectionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditSubsectionsOpsAuditSubsectionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**OpsAuditSubsections**](OpsAuditSubsections.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditSubsectionsOpsAuditSubsectionIdPut

> OpsAuditSubsections OpsAuditSubsectionsOpsAuditSubsectionIdPut(ctx, opsAuditSubsectionId).OpsAuditSubsectionsPut(opsAuditSubsectionsPut).Execute()



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
	opsAuditSubsectionId := int64(789) // int64 | Model id
	opsAuditSubsectionsPut := *openapiclient.NewOpsAuditSubsectionsPut() // OpsAuditSubsectionsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditSubsectionsAPI.OpsAuditSubsectionsOpsAuditSubsectionIdPut(context.Background(), opsAuditSubsectionId).OpsAuditSubsectionsPut(opsAuditSubsectionsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditSubsectionsAPI.OpsAuditSubsectionsOpsAuditSubsectionIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditSubsectionsOpsAuditSubsectionIdPut`: OpsAuditSubsections
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditSubsectionsAPI.OpsAuditSubsectionsOpsAuditSubsectionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditSubsectionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditSubsectionsOpsAuditSubsectionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opsAuditSubsectionsPut** | [**OpsAuditSubsectionsPut**](OpsAuditSubsectionsPut.md) |  | 

### Return type

[**OpsAuditSubsections**](OpsAuditSubsections.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditSubsectionsPost

> OpsAuditSubsectionsGet200Response OpsAuditSubsectionsPost(ctx).OpsAuditSubsectionsPostRequest(opsAuditSubsectionsPostRequest).Execute()



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
	opsAuditSubsectionsPostRequest := *openapiclient.NewOpsAuditSubsectionsPostRequest() // OpsAuditSubsectionsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditSubsectionsAPI.OpsAuditSubsectionsPost(context.Background()).OpsAuditSubsectionsPostRequest(opsAuditSubsectionsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditSubsectionsAPI.OpsAuditSubsectionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditSubsectionsPost`: OpsAuditSubsectionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditSubsectionsAPI.OpsAuditSubsectionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditSubsectionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opsAuditSubsectionsPostRequest** | [**OpsAuditSubsectionsPostRequest**](OpsAuditSubsectionsPostRequest.md) |  | 

### Return type

[**OpsAuditSubsectionsGet200Response**](OpsAuditSubsectionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

