# \AuditDivisionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditDivisionsAuditDivisionIdDelete**](AuditDivisionsAPI.md#AuditDivisionsAuditDivisionIdDelete) | **Delete** /audit_divisions/{audit_division_id} | 
[**AuditDivisionsAuditDivisionIdGet**](AuditDivisionsAPI.md#AuditDivisionsAuditDivisionIdGet) | **Get** /audit_divisions/{audit_division_id} | 
[**AuditDivisionsAuditDivisionIdPut**](AuditDivisionsAPI.md#AuditDivisionsAuditDivisionIdPut) | **Put** /audit_divisions/{audit_division_id} | 
[**AuditDivisionsGet**](AuditDivisionsAPI.md#AuditDivisionsGet) | **Get** /audit_divisions | 
[**AuditDivisionsPost**](AuditDivisionsAPI.md#AuditDivisionsPost) | **Post** /audit_divisions | 



## AuditDivisionsAuditDivisionIdDelete

> AuditDivisions AuditDivisionsAuditDivisionIdDelete(ctx, auditDivisionId).Execute()



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
	auditDivisionId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditDivisionsAPI.AuditDivisionsAuditDivisionIdDelete(context.Background(), auditDivisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditDivisionsAPI.AuditDivisionsAuditDivisionIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditDivisionsAuditDivisionIdDelete`: AuditDivisions
	fmt.Fprintf(os.Stdout, "Response from `AuditDivisionsAPI.AuditDivisionsAuditDivisionIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditDivisionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditDivisionsAuditDivisionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuditDivisions**](AuditDivisions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditDivisionsAuditDivisionIdGet

> AuditDivisions AuditDivisionsAuditDivisionIdGet(ctx, auditDivisionId).Include(include).Execute()



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
	auditDivisionId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditDivisionsAPI.AuditDivisionsAuditDivisionIdGet(context.Background(), auditDivisionId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditDivisionsAPI.AuditDivisionsAuditDivisionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditDivisionsAuditDivisionIdGet`: AuditDivisions
	fmt.Fprintf(os.Stdout, "Response from `AuditDivisionsAPI.AuditDivisionsAuditDivisionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditDivisionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditDivisionsAuditDivisionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditDivisions**](AuditDivisions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditDivisionsAuditDivisionIdPut

> AuditDivisions AuditDivisionsAuditDivisionIdPut(ctx, auditDivisionId).AuditDivisionsPut(auditDivisionsPut).Execute()



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
	auditDivisionId := int64(789) // int64 | Model id
	auditDivisionsPut := *openapiclient.NewAuditDivisionsPut() // AuditDivisionsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditDivisionsAPI.AuditDivisionsAuditDivisionIdPut(context.Background(), auditDivisionId).AuditDivisionsPut(auditDivisionsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditDivisionsAPI.AuditDivisionsAuditDivisionIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditDivisionsAuditDivisionIdPut`: AuditDivisions
	fmt.Fprintf(os.Stdout, "Response from `AuditDivisionsAPI.AuditDivisionsAuditDivisionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditDivisionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditDivisionsAuditDivisionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auditDivisionsPut** | [**AuditDivisionsPut**](AuditDivisionsPut.md) |  | 

### Return type

[**AuditDivisions**](AuditDivisions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditDivisionsGet

> AuditDivisionsGet200Response AuditDivisionsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.AuditDivisionsAPI.AuditDivisionsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditDivisionsAPI.AuditDivisionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditDivisionsGet`: AuditDivisionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditDivisionsAPI.AuditDivisionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditDivisionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditDivisionsGet200Response**](AuditDivisionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditDivisionsPost

> AuditDivisionsGet200Response AuditDivisionsPost(ctx).AuditDivisionsPostRequest(auditDivisionsPostRequest).Execute()



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
	auditDivisionsPostRequest := *openapiclient.NewAuditDivisionsPostRequest() // AuditDivisionsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditDivisionsAPI.AuditDivisionsPost(context.Background()).AuditDivisionsPostRequest(auditDivisionsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditDivisionsAPI.AuditDivisionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditDivisionsPost`: AuditDivisionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditDivisionsAPI.AuditDivisionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditDivisionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auditDivisionsPostRequest** | [**AuditDivisionsPostRequest**](AuditDivisionsPostRequest.md) |  | 

### Return type

[**AuditDivisionsGet200Response**](AuditDivisionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

