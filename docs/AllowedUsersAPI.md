# \AllowedUsersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllowedUsersAllowedUserIdDelete**](AllowedUsersAPI.md#AllowedUsersAllowedUserIdDelete) | **Delete** /allowed_users/{allowed_user_id} | 
[**AllowedUsersAllowedUserIdGet**](AllowedUsersAPI.md#AllowedUsersAllowedUserIdGet) | **Get** /allowed_users/{allowed_user_id} | 
[**AllowedUsersAllowedUserIdPut**](AllowedUsersAPI.md#AllowedUsersAllowedUserIdPut) | **Put** /allowed_users/{allowed_user_id} | 
[**AllowedUsersGet**](AllowedUsersAPI.md#AllowedUsersGet) | **Get** /allowed_users | 
[**AllowedUsersPost**](AllowedUsersAPI.md#AllowedUsersPost) | **Post** /allowed_users | 



## AllowedUsersAllowedUserIdDelete

> AllowedUsers AllowedUsersAllowedUserIdDelete(ctx, allowedUserId).Execute()



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
	allowedUserId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllowedUsersAPI.AllowedUsersAllowedUserIdDelete(context.Background(), allowedUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedUsersAPI.AllowedUsersAllowedUserIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AllowedUsersAllowedUserIdDelete`: AllowedUsers
	fmt.Fprintf(os.Stdout, "Response from `AllowedUsersAPI.AllowedUsersAllowedUserIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allowedUserId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAllowedUsersAllowedUserIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AllowedUsers**](AllowedUsers.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllowedUsersAllowedUserIdGet

> AllowedUsers AllowedUsersAllowedUserIdGet(ctx, allowedUserId).Include(include).Execute()



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
	allowedUserId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllowedUsersAPI.AllowedUsersAllowedUserIdGet(context.Background(), allowedUserId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedUsersAPI.AllowedUsersAllowedUserIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AllowedUsersAllowedUserIdGet`: AllowedUsers
	fmt.Fprintf(os.Stdout, "Response from `AllowedUsersAPI.AllowedUsersAllowedUserIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allowedUserId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAllowedUsersAllowedUserIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AllowedUsers**](AllowedUsers.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllowedUsersAllowedUserIdPut

> AllowedUsers AllowedUsersAllowedUserIdPut(ctx, allowedUserId).AllowedUsersPut(allowedUsersPut).Execute()



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
	allowedUserId := int64(789) // int64 | Model id
	allowedUsersPut := *openapiclient.NewAllowedUsersPut() // AllowedUsersPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllowedUsersAPI.AllowedUsersAllowedUserIdPut(context.Background(), allowedUserId).AllowedUsersPut(allowedUsersPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedUsersAPI.AllowedUsersAllowedUserIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AllowedUsersAllowedUserIdPut`: AllowedUsers
	fmt.Fprintf(os.Stdout, "Response from `AllowedUsersAPI.AllowedUsersAllowedUserIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allowedUserId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAllowedUsersAllowedUserIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allowedUsersPut** | [**AllowedUsersPut**](AllowedUsersPut.md) |  | 

### Return type

[**AllowedUsers**](AllowedUsers.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllowedUsersGet

> AllowedUsersGet200Response AllowedUsersGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.AllowedUsersAPI.AllowedUsersGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedUsersAPI.AllowedUsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AllowedUsersGet`: AllowedUsersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AllowedUsersAPI.AllowedUsersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllowedUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AllowedUsersGet200Response**](AllowedUsersGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllowedUsersPost

> AllowedUsersGet200Response AllowedUsersPost(ctx).AllowedUsersPostRequest(allowedUsersPostRequest).Execute()



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
	allowedUsersPostRequest := *openapiclient.NewAllowedUsersPostRequest() // AllowedUsersPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllowedUsersAPI.AllowedUsersPost(context.Background()).AllowedUsersPostRequest(allowedUsersPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedUsersAPI.AllowedUsersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AllowedUsersPost`: AllowedUsersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AllowedUsersAPI.AllowedUsersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllowedUsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allowedUsersPostRequest** | [**AllowedUsersPostRequest**](AllowedUsersPostRequest.md) |  | 

### Return type

[**AllowedUsersGet200Response**](AllowedUsersGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

