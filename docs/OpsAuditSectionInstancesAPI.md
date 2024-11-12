# \OpsAuditSectionInstancesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpsAuditSectionInstancesGet**](OpsAuditSectionInstancesAPI.md#OpsAuditSectionInstancesGet) | **Get** /ops_audit_section_instances | 
[**OpsAuditSectionInstancesOpsAuditSectionInstanceIdDelete**](OpsAuditSectionInstancesAPI.md#OpsAuditSectionInstancesOpsAuditSectionInstanceIdDelete) | **Delete** /ops_audit_section_instances/{ops_audit_section_instance_id} | 
[**OpsAuditSectionInstancesOpsAuditSectionInstanceIdGet**](OpsAuditSectionInstancesAPI.md#OpsAuditSectionInstancesOpsAuditSectionInstanceIdGet) | **Get** /ops_audit_section_instances/{ops_audit_section_instance_id} | 
[**OpsAuditSectionInstancesOpsAuditSectionInstanceIdPut**](OpsAuditSectionInstancesAPI.md#OpsAuditSectionInstancesOpsAuditSectionInstanceIdPut) | **Put** /ops_audit_section_instances/{ops_audit_section_instance_id} | 
[**OpsAuditSectionInstancesPost**](OpsAuditSectionInstancesAPI.md#OpsAuditSectionInstancesPost) | **Post** /ops_audit_section_instances | 



## OpsAuditSectionInstancesGet

> OpsAuditSectionInstancesGet200Response OpsAuditSectionInstancesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.OpsAuditSectionInstancesAPI.OpsAuditSectionInstancesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditSectionInstancesAPI.OpsAuditSectionInstancesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditSectionInstancesGet`: OpsAuditSectionInstancesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditSectionInstancesAPI.OpsAuditSectionInstancesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditSectionInstancesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**OpsAuditSectionInstancesGet200Response**](OpsAuditSectionInstancesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditSectionInstancesOpsAuditSectionInstanceIdDelete

> OpsAuditSectionInstances OpsAuditSectionInstancesOpsAuditSectionInstanceIdDelete(ctx, opsAuditSectionInstanceId).Execute()



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
	opsAuditSectionInstanceId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditSectionInstancesAPI.OpsAuditSectionInstancesOpsAuditSectionInstanceIdDelete(context.Background(), opsAuditSectionInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditSectionInstancesAPI.OpsAuditSectionInstancesOpsAuditSectionInstanceIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditSectionInstancesOpsAuditSectionInstanceIdDelete`: OpsAuditSectionInstances
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditSectionInstancesAPI.OpsAuditSectionInstancesOpsAuditSectionInstanceIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditSectionInstanceId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditSectionInstancesOpsAuditSectionInstanceIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpsAuditSectionInstances**](OpsAuditSectionInstances.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditSectionInstancesOpsAuditSectionInstanceIdGet

> OpsAuditSectionInstances OpsAuditSectionInstancesOpsAuditSectionInstanceIdGet(ctx, opsAuditSectionInstanceId).Include(include).Execute()



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
	opsAuditSectionInstanceId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditSectionInstancesAPI.OpsAuditSectionInstancesOpsAuditSectionInstanceIdGet(context.Background(), opsAuditSectionInstanceId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditSectionInstancesAPI.OpsAuditSectionInstancesOpsAuditSectionInstanceIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditSectionInstancesOpsAuditSectionInstanceIdGet`: OpsAuditSectionInstances
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditSectionInstancesAPI.OpsAuditSectionInstancesOpsAuditSectionInstanceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditSectionInstanceId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditSectionInstancesOpsAuditSectionInstanceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**OpsAuditSectionInstances**](OpsAuditSectionInstances.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditSectionInstancesOpsAuditSectionInstanceIdPut

> OpsAuditSectionInstances OpsAuditSectionInstancesOpsAuditSectionInstanceIdPut(ctx, opsAuditSectionInstanceId).OpsAuditSectionInstancesPut(opsAuditSectionInstancesPut).Execute()



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
	opsAuditSectionInstanceId := int64(789) // int64 | Model id
	opsAuditSectionInstancesPut := *openapiclient.NewOpsAuditSectionInstancesPut() // OpsAuditSectionInstancesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditSectionInstancesAPI.OpsAuditSectionInstancesOpsAuditSectionInstanceIdPut(context.Background(), opsAuditSectionInstanceId).OpsAuditSectionInstancesPut(opsAuditSectionInstancesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditSectionInstancesAPI.OpsAuditSectionInstancesOpsAuditSectionInstanceIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditSectionInstancesOpsAuditSectionInstanceIdPut`: OpsAuditSectionInstances
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditSectionInstancesAPI.OpsAuditSectionInstancesOpsAuditSectionInstanceIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditSectionInstanceId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditSectionInstancesOpsAuditSectionInstanceIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opsAuditSectionInstancesPut** | [**OpsAuditSectionInstancesPut**](OpsAuditSectionInstancesPut.md) |  | 

### Return type

[**OpsAuditSectionInstances**](OpsAuditSectionInstances.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditSectionInstancesPost

> OpsAuditSectionInstancesGet200Response OpsAuditSectionInstancesPost(ctx).OpsAuditSectionInstancesPostRequest(opsAuditSectionInstancesPostRequest).Execute()



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
	opsAuditSectionInstancesPostRequest := *openapiclient.NewOpsAuditSectionInstancesPostRequest() // OpsAuditSectionInstancesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditSectionInstancesAPI.OpsAuditSectionInstancesPost(context.Background()).OpsAuditSectionInstancesPostRequest(opsAuditSectionInstancesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditSectionInstancesAPI.OpsAuditSectionInstancesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditSectionInstancesPost`: OpsAuditSectionInstancesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditSectionInstancesAPI.OpsAuditSectionInstancesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditSectionInstancesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opsAuditSectionInstancesPostRequest** | [**OpsAuditSectionInstancesPostRequest**](OpsAuditSectionInstancesPostRequest.md) |  | 

### Return type

[**OpsAuditSectionInstancesGet200Response**](OpsAuditSectionInstancesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

