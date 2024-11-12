# \WorkspacesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkspacesGet**](WorkspacesAPI.md#WorkspacesGet) | **Get** /workspaces | 
[**WorkspacesPost**](WorkspacesAPI.md#WorkspacesPost) | **Post** /workspaces | 
[**WorkspacesWorkspaceIdDelete**](WorkspacesAPI.md#WorkspacesWorkspaceIdDelete) | **Delete** /workspaces/{workspace_id} | 
[**WorkspacesWorkspaceIdGet**](WorkspacesAPI.md#WorkspacesWorkspaceIdGet) | **Get** /workspaces/{workspace_id} | 
[**WorkspacesWorkspaceIdPut**](WorkspacesAPI.md#WorkspacesWorkspaceIdPut) | **Put** /workspaces/{workspace_id} | 



## WorkspacesGet

> WorkspacesGet200Response WorkspacesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.WorkspacesAPI.WorkspacesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.WorkspacesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesGet`: WorkspacesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.WorkspacesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**WorkspacesGet200Response**](WorkspacesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesPost

> WorkspacesGet200Response WorkspacesPost(ctx).WorkspacesPostRequest(workspacesPostRequest).Execute()



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
	workspacesPostRequest := *openapiclient.NewWorkspacesPostRequest() // WorkspacesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.WorkspacesPost(context.Background()).WorkspacesPostRequest(workspacesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.WorkspacesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesPost`: WorkspacesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.WorkspacesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspacesPostRequest** | [**WorkspacesPostRequest**](WorkspacesPostRequest.md) |  | 

### Return type

[**WorkspacesGet200Response**](WorkspacesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesWorkspaceIdDelete

> Workspaces WorkspacesWorkspaceIdDelete(ctx, workspaceId).Execute()



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
	workspaceId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.WorkspacesWorkspaceIdDelete(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.WorkspacesWorkspaceIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesWorkspaceIdDelete`: Workspaces
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.WorkspacesWorkspaceIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Workspaces**](Workspaces.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesWorkspaceIdGet

> Workspaces WorkspacesWorkspaceIdGet(ctx, workspaceId).Include(include).Execute()



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
	workspaceId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.WorkspacesWorkspaceIdGet(context.Background(), workspaceId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.WorkspacesWorkspaceIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesWorkspaceIdGet`: Workspaces
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.WorkspacesWorkspaceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**Workspaces**](Workspaces.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesWorkspaceIdPut

> Workspaces WorkspacesWorkspaceIdPut(ctx, workspaceId).WorkspacesPut(workspacesPut).Execute()



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
	workspaceId := int64(789) // int64 | Model id
	workspacesPut := *openapiclient.NewWorkspacesPut() // WorkspacesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.WorkspacesWorkspaceIdPut(context.Background(), workspaceId).WorkspacesPut(workspacesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.WorkspacesWorkspaceIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesWorkspaceIdPut`: Workspaces
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.WorkspacesWorkspaceIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspacesPut** | [**WorkspacesPut**](WorkspacesPut.md) |  | 

### Return type

[**Workspaces**](Workspaces.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

