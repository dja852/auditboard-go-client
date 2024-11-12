# \TeamPermissionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamPermissionsGet**](TeamPermissionsAPI.md#TeamPermissionsGet) | **Get** /team_permissions | 
[**TeamPermissionsPost**](TeamPermissionsAPI.md#TeamPermissionsPost) | **Post** /team_permissions | 
[**TeamPermissionsTeamPermissionIdDelete**](TeamPermissionsAPI.md#TeamPermissionsTeamPermissionIdDelete) | **Delete** /team_permissions/{team_permission_id} | 
[**TeamPermissionsTeamPermissionIdGet**](TeamPermissionsAPI.md#TeamPermissionsTeamPermissionIdGet) | **Get** /team_permissions/{team_permission_id} | 



## TeamPermissionsGet

> TeamPermissionsGet200Response TeamPermissionsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.TeamPermissionsAPI.TeamPermissionsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamPermissionsAPI.TeamPermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamPermissionsGet`: TeamPermissionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamPermissionsAPI.TeamPermissionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamPermissionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**TeamPermissionsGet200Response**](TeamPermissionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamPermissionsPost

> TeamPermissionsGet200Response TeamPermissionsPost(ctx).TeamPermissionsPostRequest(teamPermissionsPostRequest).Execute()



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
	teamPermissionsPostRequest := *openapiclient.NewTeamPermissionsPostRequest() // TeamPermissionsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamPermissionsAPI.TeamPermissionsPost(context.Background()).TeamPermissionsPostRequest(teamPermissionsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamPermissionsAPI.TeamPermissionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamPermissionsPost`: TeamPermissionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamPermissionsAPI.TeamPermissionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamPermissionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamPermissionsPostRequest** | [**TeamPermissionsPostRequest**](TeamPermissionsPostRequest.md) |  | 

### Return type

[**TeamPermissionsGet200Response**](TeamPermissionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamPermissionsTeamPermissionIdDelete

> TeamPermissions TeamPermissionsTeamPermissionIdDelete(ctx, teamPermissionId).Execute()



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
	teamPermissionId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamPermissionsAPI.TeamPermissionsTeamPermissionIdDelete(context.Background(), teamPermissionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamPermissionsAPI.TeamPermissionsTeamPermissionIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamPermissionsTeamPermissionIdDelete`: TeamPermissions
	fmt.Fprintf(os.Stdout, "Response from `TeamPermissionsAPI.TeamPermissionsTeamPermissionIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamPermissionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamPermissionsTeamPermissionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TeamPermissions**](TeamPermissions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamPermissionsTeamPermissionIdGet

> TeamPermissions TeamPermissionsTeamPermissionIdGet(ctx, teamPermissionId).Include(include).Execute()



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
	teamPermissionId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamPermissionsAPI.TeamPermissionsTeamPermissionIdGet(context.Background(), teamPermissionId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamPermissionsAPI.TeamPermissionsTeamPermissionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamPermissionsTeamPermissionIdGet`: TeamPermissions
	fmt.Fprintf(os.Stdout, "Response from `TeamPermissionsAPI.TeamPermissionsTeamPermissionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamPermissionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamPermissionsTeamPermissionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**TeamPermissions**](TeamPermissions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

