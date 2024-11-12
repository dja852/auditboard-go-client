# \FilesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesFileIdDelete**](FilesAPI.md#FilesFileIdDelete) | **Delete** /files/{file_id} | 
[**FilesFileIdGet**](FilesAPI.md#FilesFileIdGet) | **Get** /files/{file_id} | 
[**FilesFileIdPut**](FilesAPI.md#FilesFileIdPut) | **Put** /files/{file_id} | 
[**FilesGet**](FilesAPI.md#FilesGet) | **Get** /files | 
[**FilesPost**](FilesAPI.md#FilesPost) | **Post** /files | 



## FilesFileIdDelete

> Files FilesFileIdDelete(ctx, fileId).Execute()



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
	fileId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.FilesFileIdDelete(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.FilesFileIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilesFileIdDelete`: Files
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.FilesFileIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesFileIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Files**](Files.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesFileIdGet

> Files FilesFileIdGet(ctx, fileId).Include(include).Execute()



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
	fileId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.FilesFileIdGet(context.Background(), fileId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.FilesFileIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilesFileIdGet`: Files
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.FilesFileIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesFileIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**Files**](Files.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesFileIdPut

> Files FilesFileIdPut(ctx, fileId).FilesPut(filesPut).Execute()



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
	fileId := int64(789) // int64 | Model id
	filesPut := *openapiclient.NewFilesPut() // FilesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.FilesFileIdPut(context.Background(), fileId).FilesPut(filesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.FilesFileIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilesFileIdPut`: Files
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.FilesFileIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesFileIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filesPut** | [**FilesPut**](FilesPut.md) |  | 

### Return type

[**Files**](Files.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesGet

> FilesGet200Response FilesGet(ctx).FilterDateTimeKeyFilterOperator(filterDateTimeKeyFilterOperator).FilterModelModelId(filterModelModelId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/dja852/auditboard-go-client"
)

func main() {
	filterDateTimeKeyFilterOperator := time.Now() // time.Time | Refer to the API filtering section <a href=\"https://developer.auditboard.com/docs/api-filtering\">here</a> for more information about filter operators. This query parameter can be used to filter the results by created_at dates. Ex: `?filter[created_at][$gt]=[2022-01-19T18:11:51.107Z]` (optional)
	filterModelModelId := []string{"FilterModelModelId_example"} // []string | Filters the results by Model and Optional Model Id value(s). Ex. `?filter[assessment]` , `?filter[action_plan][action_plan_id]=40,41` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.FilesGet(context.Background()).FilterDateTimeKeyFilterOperator(filterDateTimeKeyFilterOperator).FilterModelModelId(filterModelModelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.FilesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilesGet`: FilesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.FilesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterDateTimeKeyFilterOperator** | **time.Time** | Refer to the API filtering section &lt;a href&#x3D;\&quot;https://developer.auditboard.com/docs/api-filtering\&quot;&gt;here&lt;/a&gt; for more information about filter operators. This query parameter can be used to filter the results by created_at dates. Ex: &#x60;?filter[created_at][$gt]&#x3D;[2022-01-19T18:11:51.107Z]&#x60; | 
 **filterModelModelId** | **[]string** | Filters the results by Model and Optional Model Id value(s). Ex. &#x60;?filter[assessment]&#x60; , &#x60;?filter[action_plan][action_plan_id]&#x3D;40,41&#x60; | 

### Return type

[**FilesGet200Response**](FilesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesPost

> FilesGet200Response FilesPost(ctx).FilesPostRequest(filesPostRequest).Execute()



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
	filesPostRequest := *openapiclient.NewFilesPostRequest() // FilesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.FilesPost(context.Background()).FilesPostRequest(filesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.FilesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilesPost`: FilesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.FilesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filesPostRequest** | [**FilesPostRequest**](FilesPostRequest.md) |  | 

### Return type

[**FilesGet200Response**](FilesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

