# \RcwProjectsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RcwProjectsGet**](RcwProjectsAPI.md#RcwProjectsGet) | **Get** /rcw_projects | 
[**RcwProjectsPost**](RcwProjectsAPI.md#RcwProjectsPost) | **Post** /rcw_projects | 
[**RcwProjectsRcwProjectIdDelete**](RcwProjectsAPI.md#RcwProjectsRcwProjectIdDelete) | **Delete** /rcw_projects/{rcw_project_id} | 
[**RcwProjectsRcwProjectIdGet**](RcwProjectsAPI.md#RcwProjectsRcwProjectIdGet) | **Get** /rcw_projects/{rcw_project_id} | 
[**RcwProjectsRcwProjectIdPut**](RcwProjectsAPI.md#RcwProjectsRcwProjectIdPut) | **Put** /rcw_projects/{rcw_project_id} | 



## RcwProjectsGet

> RcwProjectsGet200Response RcwProjectsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.RcwProjectsAPI.RcwProjectsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwProjectsAPI.RcwProjectsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwProjectsGet`: RcwProjectsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RcwProjectsAPI.RcwProjectsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRcwProjectsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**RcwProjectsGet200Response**](RcwProjectsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RcwProjectsPost

> RcwProjectsGet200Response RcwProjectsPost(ctx).RcwProjectsPostRequest(rcwProjectsPostRequest).Execute()



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
	rcwProjectsPostRequest := *openapiclient.NewRcwProjectsPostRequest() // RcwProjectsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RcwProjectsAPI.RcwProjectsPost(context.Background()).RcwProjectsPostRequest(rcwProjectsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwProjectsAPI.RcwProjectsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwProjectsPost`: RcwProjectsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RcwProjectsAPI.RcwProjectsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRcwProjectsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rcwProjectsPostRequest** | [**RcwProjectsPostRequest**](RcwProjectsPostRequest.md) |  | 

### Return type

[**RcwProjectsGet200Response**](RcwProjectsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RcwProjectsRcwProjectIdDelete

> RcwProjects RcwProjectsRcwProjectIdDelete(ctx, rcwProjectId).Execute()



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
	rcwProjectId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RcwProjectsAPI.RcwProjectsRcwProjectIdDelete(context.Background(), rcwProjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwProjectsAPI.RcwProjectsRcwProjectIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwProjectsRcwProjectIdDelete`: RcwProjects
	fmt.Fprintf(os.Stdout, "Response from `RcwProjectsAPI.RcwProjectsRcwProjectIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rcwProjectId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRcwProjectsRcwProjectIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RcwProjects**](RcwProjects.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RcwProjectsRcwProjectIdGet

> RcwProjects RcwProjectsRcwProjectIdGet(ctx, rcwProjectId).Include(include).Execute()



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
	rcwProjectId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RcwProjectsAPI.RcwProjectsRcwProjectIdGet(context.Background(), rcwProjectId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwProjectsAPI.RcwProjectsRcwProjectIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwProjectsRcwProjectIdGet`: RcwProjects
	fmt.Fprintf(os.Stdout, "Response from `RcwProjectsAPI.RcwProjectsRcwProjectIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rcwProjectId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRcwProjectsRcwProjectIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**RcwProjects**](RcwProjects.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RcwProjectsRcwProjectIdPut

> RcwProjects RcwProjectsRcwProjectIdPut(ctx, rcwProjectId).RcwProjectsPut(rcwProjectsPut).Execute()



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
	rcwProjectId := int64(789) // int64 | Model id
	rcwProjectsPut := *openapiclient.NewRcwProjectsPut() // RcwProjectsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RcwProjectsAPI.RcwProjectsRcwProjectIdPut(context.Background(), rcwProjectId).RcwProjectsPut(rcwProjectsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwProjectsAPI.RcwProjectsRcwProjectIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwProjectsRcwProjectIdPut`: RcwProjects
	fmt.Fprintf(os.Stdout, "Response from `RcwProjectsAPI.RcwProjectsRcwProjectIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rcwProjectId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRcwProjectsRcwProjectIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rcwProjectsPut** | [**RcwProjectsPut**](RcwProjectsPut.md) |  | 

### Return type

[**RcwProjects**](RcwProjects.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

