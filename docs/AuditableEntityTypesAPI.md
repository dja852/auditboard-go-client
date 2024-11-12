# \AuditableEntityTypesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditableEntityTypesAuditableEntityTypeIdDelete**](AuditableEntityTypesAPI.md#AuditableEntityTypesAuditableEntityTypeIdDelete) | **Delete** /auditable_entity_types/{auditable_entity_type_id} | 
[**AuditableEntityTypesAuditableEntityTypeIdGet**](AuditableEntityTypesAPI.md#AuditableEntityTypesAuditableEntityTypeIdGet) | **Get** /auditable_entity_types/{auditable_entity_type_id} | 
[**AuditableEntityTypesAuditableEntityTypeIdPut**](AuditableEntityTypesAPI.md#AuditableEntityTypesAuditableEntityTypeIdPut) | **Put** /auditable_entity_types/{auditable_entity_type_id} | 
[**AuditableEntityTypesGet**](AuditableEntityTypesAPI.md#AuditableEntityTypesGet) | **Get** /auditable_entity_types | 
[**AuditableEntityTypesPost**](AuditableEntityTypesAPI.md#AuditableEntityTypesPost) | **Post** /auditable_entity_types | 



## AuditableEntityTypesAuditableEntityTypeIdDelete

> AuditableEntityTypes AuditableEntityTypesAuditableEntityTypeIdDelete(ctx, auditableEntityTypeId).Execute()



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
	auditableEntityTypeId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditableEntityTypesAPI.AuditableEntityTypesAuditableEntityTypeIdDelete(context.Background(), auditableEntityTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditableEntityTypesAPI.AuditableEntityTypesAuditableEntityTypeIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditableEntityTypesAuditableEntityTypeIdDelete`: AuditableEntityTypes
	fmt.Fprintf(os.Stdout, "Response from `AuditableEntityTypesAPI.AuditableEntityTypesAuditableEntityTypeIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditableEntityTypeId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditableEntityTypesAuditableEntityTypeIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuditableEntityTypes**](AuditableEntityTypes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditableEntityTypesAuditableEntityTypeIdGet

> AuditableEntityTypes AuditableEntityTypesAuditableEntityTypeIdGet(ctx, auditableEntityTypeId).Include(include).Execute()



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
	auditableEntityTypeId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditableEntityTypesAPI.AuditableEntityTypesAuditableEntityTypeIdGet(context.Background(), auditableEntityTypeId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditableEntityTypesAPI.AuditableEntityTypesAuditableEntityTypeIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditableEntityTypesAuditableEntityTypeIdGet`: AuditableEntityTypes
	fmt.Fprintf(os.Stdout, "Response from `AuditableEntityTypesAPI.AuditableEntityTypesAuditableEntityTypeIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditableEntityTypeId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditableEntityTypesAuditableEntityTypeIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditableEntityTypes**](AuditableEntityTypes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditableEntityTypesAuditableEntityTypeIdPut

> AuditableEntityTypes AuditableEntityTypesAuditableEntityTypeIdPut(ctx, auditableEntityTypeId).AuditableEntityTypesPut(auditableEntityTypesPut).Execute()



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
	auditableEntityTypeId := int64(789) // int64 | Model id
	auditableEntityTypesPut := *openapiclient.NewAuditableEntityTypesPut() // AuditableEntityTypesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditableEntityTypesAPI.AuditableEntityTypesAuditableEntityTypeIdPut(context.Background(), auditableEntityTypeId).AuditableEntityTypesPut(auditableEntityTypesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditableEntityTypesAPI.AuditableEntityTypesAuditableEntityTypeIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditableEntityTypesAuditableEntityTypeIdPut`: AuditableEntityTypes
	fmt.Fprintf(os.Stdout, "Response from `AuditableEntityTypesAPI.AuditableEntityTypesAuditableEntityTypeIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditableEntityTypeId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditableEntityTypesAuditableEntityTypeIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auditableEntityTypesPut** | [**AuditableEntityTypesPut**](AuditableEntityTypesPut.md) |  | 

### Return type

[**AuditableEntityTypes**](AuditableEntityTypes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditableEntityTypesGet

> AuditableEntityTypesGet200Response AuditableEntityTypesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.AuditableEntityTypesAPI.AuditableEntityTypesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditableEntityTypesAPI.AuditableEntityTypesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditableEntityTypesGet`: AuditableEntityTypesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditableEntityTypesAPI.AuditableEntityTypesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditableEntityTypesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditableEntityTypesGet200Response**](AuditableEntityTypesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditableEntityTypesPost

> AuditableEntityTypesGet200Response AuditableEntityTypesPost(ctx).AuditableEntityTypesPostRequest(auditableEntityTypesPostRequest).Execute()



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
	auditableEntityTypesPostRequest := *openapiclient.NewAuditableEntityTypesPostRequest() // AuditableEntityTypesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditableEntityTypesAPI.AuditableEntityTypesPost(context.Background()).AuditableEntityTypesPostRequest(auditableEntityTypesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditableEntityTypesAPI.AuditableEntityTypesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditableEntityTypesPost`: AuditableEntityTypesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditableEntityTypesAPI.AuditableEntityTypesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditableEntityTypesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auditableEntityTypesPostRequest** | [**AuditableEntityTypesPostRequest**](AuditableEntityTypesPostRequest.md) |  | 

### Return type

[**AuditableEntityTypesGet200Response**](AuditableEntityTypesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

