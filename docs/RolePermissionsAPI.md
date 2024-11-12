# \RolePermissionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RolePermissionsGet**](RolePermissionsAPI.md#RolePermissionsGet) | **Get** /role_permissions | 
[**RolePermissionsRolePermissionIdGet**](RolePermissionsAPI.md#RolePermissionsRolePermissionIdGet) | **Get** /role_permissions/{role_permission_id} | 
[**RolePermissionsRolePermissionIdPut**](RolePermissionsAPI.md#RolePermissionsRolePermissionIdPut) | **Put** /role_permissions/{role_permission_id} | 



## RolePermissionsGet

> RolePermissionsGet200Response RolePermissionsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.RolePermissionsAPI.RolePermissionsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolePermissionsAPI.RolePermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RolePermissionsGet`: RolePermissionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RolePermissionsAPI.RolePermissionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRolePermissionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**RolePermissionsGet200Response**](RolePermissionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolePermissionsRolePermissionIdGet

> RolePermissions RolePermissionsRolePermissionIdGet(ctx, rolePermissionId).Include(include).Execute()



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
	rolePermissionId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolePermissionsAPI.RolePermissionsRolePermissionIdGet(context.Background(), rolePermissionId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolePermissionsAPI.RolePermissionsRolePermissionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RolePermissionsRolePermissionIdGet`: RolePermissions
	fmt.Fprintf(os.Stdout, "Response from `RolePermissionsAPI.RolePermissionsRolePermissionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rolePermissionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolePermissionsRolePermissionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**RolePermissions**](RolePermissions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolePermissionsRolePermissionIdPut

> RolePermissions RolePermissionsRolePermissionIdPut(ctx, rolePermissionId).RolePermissionsPut(rolePermissionsPut).Execute()



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
	rolePermissionId := int64(789) // int64 | Model id
	rolePermissionsPut := *openapiclient.NewRolePermissionsPut() // RolePermissionsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolePermissionsAPI.RolePermissionsRolePermissionIdPut(context.Background(), rolePermissionId).RolePermissionsPut(rolePermissionsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolePermissionsAPI.RolePermissionsRolePermissionIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RolePermissionsRolePermissionIdPut`: RolePermissions
	fmt.Fprintf(os.Stdout, "Response from `RolePermissionsAPI.RolePermissionsRolePermissionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rolePermissionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolePermissionsRolePermissionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rolePermissionsPut** | [**RolePermissionsPut**](RolePermissionsPut.md) |  | 

### Return type

[**RolePermissions**](RolePermissions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

