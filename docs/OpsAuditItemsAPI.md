# \OpsAuditItemsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpsAuditItemsGet**](OpsAuditItemsAPI.md#OpsAuditItemsGet) | **Get** /ops_audit_items | 
[**OpsAuditItemsOpsAuditItemIdDelete**](OpsAuditItemsAPI.md#OpsAuditItemsOpsAuditItemIdDelete) | **Delete** /ops_audit_items/{ops_audit_item_id} | 
[**OpsAuditItemsOpsAuditItemIdGet**](OpsAuditItemsAPI.md#OpsAuditItemsOpsAuditItemIdGet) | **Get** /ops_audit_items/{ops_audit_item_id} | 
[**OpsAuditItemsOpsAuditItemIdPut**](OpsAuditItemsAPI.md#OpsAuditItemsOpsAuditItemIdPut) | **Put** /ops_audit_items/{ops_audit_item_id} | 
[**OpsAuditItemsPost**](OpsAuditItemsAPI.md#OpsAuditItemsPost) | **Post** /ops_audit_items | 



## OpsAuditItemsGet

> OpsAuditItemsGet200Response OpsAuditItemsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.OpsAuditItemsAPI.OpsAuditItemsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditItemsAPI.OpsAuditItemsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditItemsGet`: OpsAuditItemsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditItemsAPI.OpsAuditItemsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditItemsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**OpsAuditItemsGet200Response**](OpsAuditItemsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditItemsOpsAuditItemIdDelete

> OpsAuditItems OpsAuditItemsOpsAuditItemIdDelete(ctx, opsAuditItemId).Execute()



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
	opsAuditItemId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditItemsAPI.OpsAuditItemsOpsAuditItemIdDelete(context.Background(), opsAuditItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditItemsAPI.OpsAuditItemsOpsAuditItemIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditItemsOpsAuditItemIdDelete`: OpsAuditItems
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditItemsAPI.OpsAuditItemsOpsAuditItemIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditItemId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditItemsOpsAuditItemIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpsAuditItems**](OpsAuditItems.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditItemsOpsAuditItemIdGet

> OpsAuditItems OpsAuditItemsOpsAuditItemIdGet(ctx, opsAuditItemId).Include(include).Execute()



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
	opsAuditItemId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditItemsAPI.OpsAuditItemsOpsAuditItemIdGet(context.Background(), opsAuditItemId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditItemsAPI.OpsAuditItemsOpsAuditItemIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditItemsOpsAuditItemIdGet`: OpsAuditItems
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditItemsAPI.OpsAuditItemsOpsAuditItemIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditItemId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditItemsOpsAuditItemIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**OpsAuditItems**](OpsAuditItems.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditItemsOpsAuditItemIdPut

> OpsAuditItems OpsAuditItemsOpsAuditItemIdPut(ctx, opsAuditItemId).OpsAuditItemsPut(opsAuditItemsPut).Execute()



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
	opsAuditItemId := int64(789) // int64 | Model id
	opsAuditItemsPut := *openapiclient.NewOpsAuditItemsPut() // OpsAuditItemsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditItemsAPI.OpsAuditItemsOpsAuditItemIdPut(context.Background(), opsAuditItemId).OpsAuditItemsPut(opsAuditItemsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditItemsAPI.OpsAuditItemsOpsAuditItemIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditItemsOpsAuditItemIdPut`: OpsAuditItems
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditItemsAPI.OpsAuditItemsOpsAuditItemIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditItemId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditItemsOpsAuditItemIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opsAuditItemsPut** | [**OpsAuditItemsPut**](OpsAuditItemsPut.md) |  | 

### Return type

[**OpsAuditItems**](OpsAuditItems.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditItemsPost

> OpsAuditItemsGet200Response OpsAuditItemsPost(ctx).OpsAuditItemsPostRequest(opsAuditItemsPostRequest).Execute()



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
	opsAuditItemsPostRequest := *openapiclient.NewOpsAuditItemsPostRequest() // OpsAuditItemsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditItemsAPI.OpsAuditItemsPost(context.Background()).OpsAuditItemsPostRequest(opsAuditItemsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditItemsAPI.OpsAuditItemsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditItemsPost`: OpsAuditItemsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditItemsAPI.OpsAuditItemsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditItemsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opsAuditItemsPostRequest** | [**OpsAuditItemsPostRequest**](OpsAuditItemsPostRequest.md) |  | 

### Return type

[**OpsAuditItemsGet200Response**](OpsAuditItemsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

