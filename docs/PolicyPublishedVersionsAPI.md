# \PolicyPublishedVersionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PolicyPublishedVersionsGet**](PolicyPublishedVersionsAPI.md#PolicyPublishedVersionsGet) | **Get** /policy_published_versions | 
[**PolicyPublishedVersionsPolicyPublishedVersionIdDelete**](PolicyPublishedVersionsAPI.md#PolicyPublishedVersionsPolicyPublishedVersionIdDelete) | **Delete** /policy_published_versions/{policy_published_version_id} | 
[**PolicyPublishedVersionsPolicyPublishedVersionIdGet**](PolicyPublishedVersionsAPI.md#PolicyPublishedVersionsPolicyPublishedVersionIdGet) | **Get** /policy_published_versions/{policy_published_version_id} | 
[**PolicyPublishedVersionsPolicyPublishedVersionIdPut**](PolicyPublishedVersionsAPI.md#PolicyPublishedVersionsPolicyPublishedVersionIdPut) | **Put** /policy_published_versions/{policy_published_version_id} | 
[**PolicyPublishedVersionsPost**](PolicyPublishedVersionsAPI.md#PolicyPublishedVersionsPost) | **Post** /policy_published_versions | 



## PolicyPublishedVersionsGet

> PolicyPublishedVersionsGet200Response PolicyPublishedVersionsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.PolicyPublishedVersionsAPI.PolicyPublishedVersionsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyPublishedVersionsAPI.PolicyPublishedVersionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PolicyPublishedVersionsGet`: PolicyPublishedVersionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PolicyPublishedVersionsAPI.PolicyPublishedVersionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPolicyPublishedVersionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**PolicyPublishedVersionsGet200Response**](PolicyPublishedVersionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PolicyPublishedVersionsPolicyPublishedVersionIdDelete

> PolicyPublishedVersions PolicyPublishedVersionsPolicyPublishedVersionIdDelete(ctx, policyPublishedVersionId).Execute()



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
	policyPublishedVersionId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyPublishedVersionsAPI.PolicyPublishedVersionsPolicyPublishedVersionIdDelete(context.Background(), policyPublishedVersionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyPublishedVersionsAPI.PolicyPublishedVersionsPolicyPublishedVersionIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PolicyPublishedVersionsPolicyPublishedVersionIdDelete`: PolicyPublishedVersions
	fmt.Fprintf(os.Stdout, "Response from `PolicyPublishedVersionsAPI.PolicyPublishedVersionsPolicyPublishedVersionIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyPublishedVersionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyPublishedVersionsPolicyPublishedVersionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyPublishedVersions**](PolicyPublishedVersions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PolicyPublishedVersionsPolicyPublishedVersionIdGet

> PolicyPublishedVersions PolicyPublishedVersionsPolicyPublishedVersionIdGet(ctx, policyPublishedVersionId).Include(include).Execute()



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
	policyPublishedVersionId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyPublishedVersionsAPI.PolicyPublishedVersionsPolicyPublishedVersionIdGet(context.Background(), policyPublishedVersionId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyPublishedVersionsAPI.PolicyPublishedVersionsPolicyPublishedVersionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PolicyPublishedVersionsPolicyPublishedVersionIdGet`: PolicyPublishedVersions
	fmt.Fprintf(os.Stdout, "Response from `PolicyPublishedVersionsAPI.PolicyPublishedVersionsPolicyPublishedVersionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyPublishedVersionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyPublishedVersionsPolicyPublishedVersionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**PolicyPublishedVersions**](PolicyPublishedVersions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PolicyPublishedVersionsPolicyPublishedVersionIdPut

> PolicyPublishedVersions PolicyPublishedVersionsPolicyPublishedVersionIdPut(ctx, policyPublishedVersionId).Body(body).Execute()



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
	policyPublishedVersionId := int64(789) // int64 | Model id
	body := interface{}(987) // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyPublishedVersionsAPI.PolicyPublishedVersionsPolicyPublishedVersionIdPut(context.Background(), policyPublishedVersionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyPublishedVersionsAPI.PolicyPublishedVersionsPolicyPublishedVersionIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PolicyPublishedVersionsPolicyPublishedVersionIdPut`: PolicyPublishedVersions
	fmt.Fprintf(os.Stdout, "Response from `PolicyPublishedVersionsAPI.PolicyPublishedVersionsPolicyPublishedVersionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyPublishedVersionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyPublishedVersionsPolicyPublishedVersionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** |  | 

### Return type

[**PolicyPublishedVersions**](PolicyPublishedVersions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PolicyPublishedVersionsPost

> PolicyPublishedVersionsGet200Response PolicyPublishedVersionsPost(ctx).PolicyPublishedVersionsPostRequest(policyPublishedVersionsPostRequest).Execute()



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
	policyPublishedVersionsPostRequest := *openapiclient.NewPolicyPublishedVersionsPostRequest() // PolicyPublishedVersionsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyPublishedVersionsAPI.PolicyPublishedVersionsPost(context.Background()).PolicyPublishedVersionsPostRequest(policyPublishedVersionsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyPublishedVersionsAPI.PolicyPublishedVersionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PolicyPublishedVersionsPost`: PolicyPublishedVersionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PolicyPublishedVersionsAPI.PolicyPublishedVersionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPolicyPublishedVersionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyPublishedVersionsPostRequest** | [**PolicyPublishedVersionsPostRequest**](PolicyPublishedVersionsPostRequest.md) |  | 

### Return type

[**PolicyPublishedVersionsGet200Response**](PolicyPublishedVersionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

