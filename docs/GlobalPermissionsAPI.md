# \GlobalPermissionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlobalPermissionsGet**](GlobalPermissionsAPI.md#GlobalPermissionsGet) | **Get** /global_permissions | 
[**GlobalPermissionsGlobalPermissionIdGet**](GlobalPermissionsAPI.md#GlobalPermissionsGlobalPermissionIdGet) | **Get** /global_permissions/{global_permission_id} | 
[**GlobalPermissionsGlobalPermissionIdPut**](GlobalPermissionsAPI.md#GlobalPermissionsGlobalPermissionIdPut) | **Put** /global_permissions/{global_permission_id} | 



## GlobalPermissionsGet

> GlobalPermissionsGet200Response GlobalPermissionsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.GlobalPermissionsAPI.GlobalPermissionsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalPermissionsAPI.GlobalPermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalPermissionsGet`: GlobalPermissionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalPermissionsAPI.GlobalPermissionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlobalPermissionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**GlobalPermissionsGet200Response**](GlobalPermissionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlobalPermissionsGlobalPermissionIdGet

> GlobalPermissions GlobalPermissionsGlobalPermissionIdGet(ctx, globalPermissionId).Include(include).Execute()



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
	globalPermissionId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalPermissionsAPI.GlobalPermissionsGlobalPermissionIdGet(context.Background(), globalPermissionId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalPermissionsAPI.GlobalPermissionsGlobalPermissionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalPermissionsGlobalPermissionIdGet`: GlobalPermissions
	fmt.Fprintf(os.Stdout, "Response from `GlobalPermissionsAPI.GlobalPermissionsGlobalPermissionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**globalPermissionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalPermissionsGlobalPermissionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**GlobalPermissions**](GlobalPermissions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlobalPermissionsGlobalPermissionIdPut

> GlobalPermissions GlobalPermissionsGlobalPermissionIdPut(ctx, globalPermissionId).GlobalPermissionsPut(globalPermissionsPut).Execute()



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
	globalPermissionId := int64(789) // int64 | Model id
	globalPermissionsPut := *openapiclient.NewGlobalPermissionsPut() // GlobalPermissionsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalPermissionsAPI.GlobalPermissionsGlobalPermissionIdPut(context.Background(), globalPermissionId).GlobalPermissionsPut(globalPermissionsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalPermissionsAPI.GlobalPermissionsGlobalPermissionIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalPermissionsGlobalPermissionIdPut`: GlobalPermissions
	fmt.Fprintf(os.Stdout, "Response from `GlobalPermissionsAPI.GlobalPermissionsGlobalPermissionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**globalPermissionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalPermissionsGlobalPermissionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **globalPermissionsPut** | [**GlobalPermissionsPut**](GlobalPermissionsPut.md) |  | 

### Return type

[**GlobalPermissions**](GlobalPermissions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

