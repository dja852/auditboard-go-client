# \CommentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommentsCommentIdDelete**](CommentsAPI.md#CommentsCommentIdDelete) | **Delete** /comments/{comment_id} | 
[**CommentsCommentIdGet**](CommentsAPI.md#CommentsCommentIdGet) | **Get** /comments/{comment_id} | 
[**CommentsCommentIdPut**](CommentsAPI.md#CommentsCommentIdPut) | **Put** /comments/{comment_id} | 
[**CommentsGet**](CommentsAPI.md#CommentsGet) | **Get** /comments | 
[**CommentsPost**](CommentsAPI.md#CommentsPost) | **Post** /comments | 



## CommentsCommentIdDelete

> Comments CommentsCommentIdDelete(ctx, commentId).Execute()



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
	commentId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.CommentsCommentIdDelete(context.Background(), commentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.CommentsCommentIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommentsCommentIdDelete`: Comments
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.CommentsCommentIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommentsCommentIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Comments**](Comments.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommentsCommentIdGet

> Comments CommentsCommentIdGet(ctx, commentId).Include(include).Execute()



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
	commentId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.CommentsCommentIdGet(context.Background(), commentId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.CommentsCommentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommentsCommentIdGet`: Comments
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.CommentsCommentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommentsCommentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**Comments**](Comments.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommentsCommentIdPut

> Comments CommentsCommentIdPut(ctx, commentId).CommentsPut(commentsPut).Execute()



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
	commentId := int64(789) // int64 | Model id
	commentsPut := *openapiclient.NewCommentsPut() // CommentsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.CommentsCommentIdPut(context.Background(), commentId).CommentsPut(commentsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.CommentsCommentIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommentsCommentIdPut`: Comments
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.CommentsCommentIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommentsCommentIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commentsPut** | [**CommentsPut**](CommentsPut.md) |  | 

### Return type

[**Comments**](Comments.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommentsGet

> CommentsGet200Response CommentsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.CommentsAPI.CommentsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.CommentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommentsGet`: CommentsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.CommentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**CommentsGet200Response**](CommentsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommentsPost

> CommentsGet200Response CommentsPost(ctx).CommentsPostRequest(commentsPostRequest).Execute()



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
	commentsPostRequest := *openapiclient.NewCommentsPostRequest() // CommentsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.CommentsPost(context.Background()).CommentsPostRequest(commentsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.CommentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommentsPost`: CommentsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.CommentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commentsPostRequest** | [**CommentsPostRequest**](CommentsPostRequest.md) |  | 

### Return type

[**CommentsGet200Response**](CommentsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

