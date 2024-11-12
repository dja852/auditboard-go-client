# \UploadFilesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadFilesPost**](UploadFilesAPI.md#UploadFilesPost) | **Post** /upload_files | 
[**UploadFilesUploadFileIdCompletePost**](UploadFilesAPI.md#UploadFilesUploadFileIdCompletePost) | **Post** /upload_files/{upload_file_id}/complete | 
[**UploadFilesUploadFileIdDelete**](UploadFilesAPI.md#UploadFilesUploadFileIdDelete) | **Delete** /upload_files/{upload_file_id} | 



## UploadFilesPost

> UploadFilesPost200Response UploadFilesPost(ctx).UploadFilesPostRequest(uploadFilesPostRequest).Execute()



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
	uploadFilesPostRequest := *openapiclient.NewUploadFilesPostRequest() // UploadFilesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadFilesAPI.UploadFilesPost(context.Background()).UploadFilesPostRequest(uploadFilesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadFilesAPI.UploadFilesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadFilesPost`: UploadFilesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UploadFilesAPI.UploadFilesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadFilesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadFilesPostRequest** | [**UploadFilesPostRequest**](UploadFilesPostRequest.md) |  | 

### Return type

[**UploadFilesPost200Response**](UploadFilesPost200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFilesUploadFileIdCompletePost

> FilesGet200Response UploadFilesUploadFileIdCompletePost(ctx, uploadFileId).Execute()



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
	uploadFileId := int64(789) // int64 | This is the id field given in the response from the endpoint /upload_files

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadFilesAPI.UploadFilesUploadFileIdCompletePost(context.Background(), uploadFileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadFilesAPI.UploadFilesUploadFileIdCompletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadFilesUploadFileIdCompletePost`: FilesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UploadFilesAPI.UploadFilesUploadFileIdCompletePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadFileId** | **int64** | This is the id field given in the response from the endpoint /upload_files | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFilesUploadFileIdCompletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UploadFilesUploadFileIdDelete

> UploadFiles UploadFilesUploadFileIdDelete(ctx, uploadFileId).Execute()



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
	uploadFileId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadFilesAPI.UploadFilesUploadFileIdDelete(context.Background(), uploadFileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadFilesAPI.UploadFilesUploadFileIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadFilesUploadFileIdDelete`: UploadFiles
	fmt.Fprintf(os.Stdout, "Response from `UploadFilesAPI.UploadFilesUploadFileIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadFileId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFilesUploadFileIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UploadFiles**](UploadFiles.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

