# \AuditableEntitiesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditableEntitiesAuditableEntityIdDelete**](AuditableEntitiesAPI.md#AuditableEntitiesAuditableEntityIdDelete) | **Delete** /auditable_entities/{auditable_entity_id} | 
[**AuditableEntitiesAuditableEntityIdGet**](AuditableEntitiesAPI.md#AuditableEntitiesAuditableEntityIdGet) | **Get** /auditable_entities/{auditable_entity_id} | 
[**AuditableEntitiesAuditableEntityIdPut**](AuditableEntitiesAPI.md#AuditableEntitiesAuditableEntityIdPut) | **Put** /auditable_entities/{auditable_entity_id} | 
[**AuditableEntitiesGet**](AuditableEntitiesAPI.md#AuditableEntitiesGet) | **Get** /auditable_entities | 
[**AuditableEntitiesPost**](AuditableEntitiesAPI.md#AuditableEntitiesPost) | **Post** /auditable_entities | 



## AuditableEntitiesAuditableEntityIdDelete

> AuditableEntities AuditableEntitiesAuditableEntityIdDelete(ctx, auditableEntityId).Execute()



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
	auditableEntityId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditableEntitiesAPI.AuditableEntitiesAuditableEntityIdDelete(context.Background(), auditableEntityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditableEntitiesAPI.AuditableEntitiesAuditableEntityIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditableEntitiesAuditableEntityIdDelete`: AuditableEntities
	fmt.Fprintf(os.Stdout, "Response from `AuditableEntitiesAPI.AuditableEntitiesAuditableEntityIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditableEntityId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditableEntitiesAuditableEntityIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuditableEntities**](AuditableEntities.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditableEntitiesAuditableEntityIdGet

> AuditableEntities AuditableEntitiesAuditableEntityIdGet(ctx, auditableEntityId).Include(include).Execute()



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
	auditableEntityId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditableEntitiesAPI.AuditableEntitiesAuditableEntityIdGet(context.Background(), auditableEntityId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditableEntitiesAPI.AuditableEntitiesAuditableEntityIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditableEntitiesAuditableEntityIdGet`: AuditableEntities
	fmt.Fprintf(os.Stdout, "Response from `AuditableEntitiesAPI.AuditableEntitiesAuditableEntityIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditableEntityId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditableEntitiesAuditableEntityIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditableEntities**](AuditableEntities.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditableEntitiesAuditableEntityIdPut

> AuditableEntities AuditableEntitiesAuditableEntityIdPut(ctx, auditableEntityId).AuditableEntitiesPut(auditableEntitiesPut).Execute()



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
	auditableEntityId := int64(789) // int64 | Model id
	auditableEntitiesPut := *openapiclient.NewAuditableEntitiesPut() // AuditableEntitiesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditableEntitiesAPI.AuditableEntitiesAuditableEntityIdPut(context.Background(), auditableEntityId).AuditableEntitiesPut(auditableEntitiesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditableEntitiesAPI.AuditableEntitiesAuditableEntityIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditableEntitiesAuditableEntityIdPut`: AuditableEntities
	fmt.Fprintf(os.Stdout, "Response from `AuditableEntitiesAPI.AuditableEntitiesAuditableEntityIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditableEntityId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditableEntitiesAuditableEntityIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auditableEntitiesPut** | [**AuditableEntitiesPut**](AuditableEntitiesPut.md) |  | 

### Return type

[**AuditableEntities**](AuditableEntities.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditableEntitiesGet

> AuditableEntitiesGet200Response AuditableEntitiesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.AuditableEntitiesAPI.AuditableEntitiesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditableEntitiesAPI.AuditableEntitiesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditableEntitiesGet`: AuditableEntitiesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditableEntitiesAPI.AuditableEntitiesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditableEntitiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditableEntitiesGet200Response**](AuditableEntitiesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditableEntitiesPost

> AuditableEntitiesGet200Response AuditableEntitiesPost(ctx).AuditableEntitiesPostRequest(auditableEntitiesPostRequest).Execute()



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
	auditableEntitiesPostRequest := *openapiclient.NewAuditableEntitiesPostRequest() // AuditableEntitiesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditableEntitiesAPI.AuditableEntitiesPost(context.Background()).AuditableEntitiesPostRequest(auditableEntitiesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditableEntitiesAPI.AuditableEntitiesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditableEntitiesPost`: AuditableEntitiesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditableEntitiesAPI.AuditableEntitiesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditableEntitiesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auditableEntitiesPostRequest** | [**AuditableEntitiesPostRequest**](AuditableEntitiesPostRequest.md) |  | 

### Return type

[**AuditableEntitiesGet200Response**](AuditableEntitiesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

