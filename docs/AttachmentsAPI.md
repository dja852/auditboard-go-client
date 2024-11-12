# \AttachmentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachmentsAttachmentIdDelete**](AttachmentsAPI.md#AttachmentsAttachmentIdDelete) | **Delete** /attachments/{attachment_id} | 
[**AttachmentsAttachmentIdGet**](AttachmentsAPI.md#AttachmentsAttachmentIdGet) | **Get** /attachments/{attachment_id} | 
[**AttachmentsAttachmentIdPut**](AttachmentsAPI.md#AttachmentsAttachmentIdPut) | **Put** /attachments/{attachment_id} | 
[**AttachmentsGet**](AttachmentsAPI.md#AttachmentsGet) | **Get** /attachments | 
[**AttachmentsPost**](AttachmentsAPI.md#AttachmentsPost) | **Post** /attachments | 



## AttachmentsAttachmentIdDelete

> Attachments AttachmentsAttachmentIdDelete(ctx, attachmentId).Execute()



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
	attachmentId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.AttachmentsAttachmentIdDelete(context.Background(), attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.AttachmentsAttachmentIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachmentsAttachmentIdDelete`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.AttachmentsAttachmentIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentsAttachmentIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Attachments**](Attachments.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachmentsAttachmentIdGet

> Attachments AttachmentsAttachmentIdGet(ctx, attachmentId).Include(include).Execute()



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
	attachmentId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.AttachmentsAttachmentIdGet(context.Background(), attachmentId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.AttachmentsAttachmentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachmentsAttachmentIdGet`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.AttachmentsAttachmentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentsAttachmentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachmentsAttachmentIdPut

> Attachments AttachmentsAttachmentIdPut(ctx, attachmentId).AttachmentsPut(attachmentsPut).Execute()



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
	attachmentId := int64(789) // int64 | Model id
	attachmentsPut := *openapiclient.NewAttachmentsPut() // AttachmentsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.AttachmentsAttachmentIdPut(context.Background(), attachmentId).AttachmentsPut(attachmentsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.AttachmentsAttachmentIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachmentsAttachmentIdPut`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.AttachmentsAttachmentIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentsAttachmentIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachmentsPut** | [**AttachmentsPut**](AttachmentsPut.md) |  | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachmentsGet

> AttachmentsGet200Response AttachmentsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.AttachmentsAPI.AttachmentsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.AttachmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachmentsGet`: AttachmentsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.AttachmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AttachmentsGet200Response**](AttachmentsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachmentsPost

> AttachmentsGet200Response AttachmentsPost(ctx).AttachmentsPostRequest(attachmentsPostRequest).Execute()



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
	attachmentsPostRequest := *openapiclient.NewAttachmentsPostRequest() // AttachmentsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.AttachmentsPost(context.Background()).AttachmentsPostRequest(attachmentsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.AttachmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachmentsPost`: AttachmentsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.AttachmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachmentsPostRequest** | [**AttachmentsPostRequest**](AttachmentsPostRequest.md) |  | 

### Return type

[**AttachmentsGet200Response**](AttachmentsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

