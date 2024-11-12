# \EntitiesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EntitiesEntityIdAddControlsPost**](EntitiesAPI.md#EntitiesEntityIdAddControlsPost) | **Post** /entities/{entity_id}/add_controls | 
[**EntitiesEntityIdDelete**](EntitiesAPI.md#EntitiesEntityIdDelete) | **Delete** /entities/{entity_id} | 
[**EntitiesEntityIdGet**](EntitiesAPI.md#EntitiesEntityIdGet) | **Get** /entities/{entity_id} | 
[**EntitiesEntityIdPut**](EntitiesAPI.md#EntitiesEntityIdPut) | **Put** /entities/{entity_id} | 
[**EntitiesGet**](EntitiesAPI.md#EntitiesGet) | **Get** /entities | 
[**EntitiesPost**](EntitiesAPI.md#EntitiesPost) | **Post** /entities | 



## EntitiesEntityIdAddControlsPost

> EntitiesEntityIdAddControlsPost200Response EntitiesEntityIdAddControlsPost(ctx, entityId).EntitiesEntityIdAddControlsPostRequest(entitiesEntityIdAddControlsPostRequest).Execute()



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
	entityId := int64(789) // int64 | The id of the entity to be added to the control(s).
	entitiesEntityIdAddControlsPostRequest := *openapiclient.NewEntitiesEntityIdAddControlsPostRequest() // EntitiesEntityIdAddControlsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitiesAPI.EntitiesEntityIdAddControlsPost(context.Background(), entityId).EntitiesEntityIdAddControlsPostRequest(entitiesEntityIdAddControlsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitiesAPI.EntitiesEntityIdAddControlsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntitiesEntityIdAddControlsPost`: EntitiesEntityIdAddControlsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `EntitiesAPI.EntitiesEntityIdAddControlsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **int64** | The id of the entity to be added to the control(s). | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntitiesEntityIdAddControlsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entitiesEntityIdAddControlsPostRequest** | [**EntitiesEntityIdAddControlsPostRequest**](EntitiesEntityIdAddControlsPostRequest.md) |  | 

### Return type

[**EntitiesEntityIdAddControlsPost200Response**](EntitiesEntityIdAddControlsPost200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitiesEntityIdDelete

> Entities EntitiesEntityIdDelete(ctx, entityId).Execute()



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
	entityId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitiesAPI.EntitiesEntityIdDelete(context.Background(), entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitiesAPI.EntitiesEntityIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntitiesEntityIdDelete`: Entities
	fmt.Fprintf(os.Stdout, "Response from `EntitiesAPI.EntitiesEntityIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntitiesEntityIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Entities**](Entities.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitiesEntityIdGet

> Entities EntitiesEntityIdGet(ctx, entityId).Include(include).Execute()



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
	entityId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitiesAPI.EntitiesEntityIdGet(context.Background(), entityId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitiesAPI.EntitiesEntityIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntitiesEntityIdGet`: Entities
	fmt.Fprintf(os.Stdout, "Response from `EntitiesAPI.EntitiesEntityIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntitiesEntityIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**Entities**](Entities.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitiesEntityIdPut

> Entities EntitiesEntityIdPut(ctx, entityId).EntitiesPut(entitiesPut).Execute()



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
	entityId := int64(789) // int64 | Model id
	entitiesPut := *openapiclient.NewEntitiesPut() // EntitiesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitiesAPI.EntitiesEntityIdPut(context.Background(), entityId).EntitiesPut(entitiesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitiesAPI.EntitiesEntityIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntitiesEntityIdPut`: Entities
	fmt.Fprintf(os.Stdout, "Response from `EntitiesAPI.EntitiesEntityIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntitiesEntityIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entitiesPut** | [**EntitiesPut**](EntitiesPut.md) |  | 

### Return type

[**Entities**](Entities.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitiesGet

> EntitiesGet200Response EntitiesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.EntitiesAPI.EntitiesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitiesAPI.EntitiesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntitiesGet`: EntitiesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EntitiesAPI.EntitiesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntitiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EntitiesGet200Response**](EntitiesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitiesPost

> EntitiesGet200Response EntitiesPost(ctx).EntitiesPostRequest(entitiesPostRequest).Execute()



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
	entitiesPostRequest := *openapiclient.NewEntitiesPostRequest() // EntitiesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitiesAPI.EntitiesPost(context.Background()).EntitiesPostRequest(entitiesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitiesAPI.EntitiesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntitiesPost`: EntitiesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EntitiesAPI.EntitiesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntitiesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entitiesPostRequest** | [**EntitiesPostRequest**](EntitiesPostRequest.md) |  | 

### Return type

[**EntitiesGet200Response**](EntitiesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

