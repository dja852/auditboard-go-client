# \EsgMetricAuditableEntitiesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdDelete**](EsgMetricAuditableEntitiesAPI.md#EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdDelete) | **Delete** /esg_metric_auditable_entities/{esg_metric_auditable_entity_id} | 
[**EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdGet**](EsgMetricAuditableEntitiesAPI.md#EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdGet) | **Get** /esg_metric_auditable_entities/{esg_metric_auditable_entity_id} | 
[**EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdPut**](EsgMetricAuditableEntitiesAPI.md#EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdPut) | **Put** /esg_metric_auditable_entities/{esg_metric_auditable_entity_id} | 
[**EsgMetricAuditableEntitiesGet**](EsgMetricAuditableEntitiesAPI.md#EsgMetricAuditableEntitiesGet) | **Get** /esg_metric_auditable_entities | 
[**EsgMetricAuditableEntitiesPost**](EsgMetricAuditableEntitiesAPI.md#EsgMetricAuditableEntitiesPost) | **Post** /esg_metric_auditable_entities | 



## EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdDelete

> EsgMetricAuditableEntities EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdDelete(ctx, esgMetricAuditableEntityId).Execute()



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
	esgMetricAuditableEntityId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricAuditableEntitiesAPI.EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdDelete(context.Background(), esgMetricAuditableEntityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricAuditableEntitiesAPI.EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdDelete`: EsgMetricAuditableEntities
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricAuditableEntitiesAPI.EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricAuditableEntityId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricAuditableEntitiesEsgMetricAuditableEntityIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EsgMetricAuditableEntities**](EsgMetricAuditableEntities.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdGet

> EsgMetricAuditableEntities EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdGet(ctx, esgMetricAuditableEntityId).Include(include).Execute()



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
	esgMetricAuditableEntityId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricAuditableEntitiesAPI.EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdGet(context.Background(), esgMetricAuditableEntityId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricAuditableEntitiesAPI.EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdGet`: EsgMetricAuditableEntities
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricAuditableEntitiesAPI.EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricAuditableEntityId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricAuditableEntitiesEsgMetricAuditableEntityIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EsgMetricAuditableEntities**](EsgMetricAuditableEntities.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdPut

> EsgMetricAuditableEntities EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdPut(ctx, esgMetricAuditableEntityId).EsgMetricAuditableEntitiesPut(esgMetricAuditableEntitiesPut).Execute()



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
	esgMetricAuditableEntityId := int64(789) // int64 | Model id
	esgMetricAuditableEntitiesPut := *openapiclient.NewEsgMetricAuditableEntitiesPut() // EsgMetricAuditableEntitiesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricAuditableEntitiesAPI.EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdPut(context.Background(), esgMetricAuditableEntityId).EsgMetricAuditableEntitiesPut(esgMetricAuditableEntitiesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricAuditableEntitiesAPI.EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdPut`: EsgMetricAuditableEntities
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricAuditableEntitiesAPI.EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esgMetricAuditableEntityId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricAuditableEntitiesEsgMetricAuditableEntityIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **esgMetricAuditableEntitiesPut** | [**EsgMetricAuditableEntitiesPut**](EsgMetricAuditableEntitiesPut.md) |  | 

### Return type

[**EsgMetricAuditableEntities**](EsgMetricAuditableEntities.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricAuditableEntitiesGet

> EsgMetricAuditableEntitiesGet200Response EsgMetricAuditableEntitiesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.EsgMetricAuditableEntitiesAPI.EsgMetricAuditableEntitiesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricAuditableEntitiesAPI.EsgMetricAuditableEntitiesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricAuditableEntitiesGet`: EsgMetricAuditableEntitiesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricAuditableEntitiesAPI.EsgMetricAuditableEntitiesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricAuditableEntitiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EsgMetricAuditableEntitiesGet200Response**](EsgMetricAuditableEntitiesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsgMetricAuditableEntitiesPost

> EsgMetricAuditableEntitiesGet200Response EsgMetricAuditableEntitiesPost(ctx).EsgMetricAuditableEntitiesPostRequest(esgMetricAuditableEntitiesPostRequest).Execute()



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
	esgMetricAuditableEntitiesPostRequest := *openapiclient.NewEsgMetricAuditableEntitiesPostRequest() // EsgMetricAuditableEntitiesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EsgMetricAuditableEntitiesAPI.EsgMetricAuditableEntitiesPost(context.Background()).EsgMetricAuditableEntitiesPostRequest(esgMetricAuditableEntitiesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EsgMetricAuditableEntitiesAPI.EsgMetricAuditableEntitiesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsgMetricAuditableEntitiesPost`: EsgMetricAuditableEntitiesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EsgMetricAuditableEntitiesAPI.EsgMetricAuditableEntitiesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsgMetricAuditableEntitiesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **esgMetricAuditableEntitiesPostRequest** | [**EsgMetricAuditableEntitiesPostRequest**](EsgMetricAuditableEntitiesPostRequest.md) |  | 

### Return type

[**EsgMetricAuditableEntitiesGet200Response**](EsgMetricAuditableEntitiesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

