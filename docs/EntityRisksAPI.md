# \EntityRisksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EntityRisksEntityRiskIdGet**](EntityRisksAPI.md#EntityRisksEntityRiskIdGet) | **Get** /entity_risks/{entity_risk_id} | 
[**EntityRisksEntityRiskIdPut**](EntityRisksAPI.md#EntityRisksEntityRiskIdPut) | **Put** /entity_risks/{entity_risk_id} | 
[**EntityRisksGet**](EntityRisksAPI.md#EntityRisksGet) | **Get** /entity_risks | 



## EntityRisksEntityRiskIdGet

> EntityRisks EntityRisksEntityRiskIdGet(ctx, entityRiskId).Include(include).Execute()



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
	entityRiskId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntityRisksAPI.EntityRisksEntityRiskIdGet(context.Background(), entityRiskId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntityRisksAPI.EntityRisksEntityRiskIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntityRisksEntityRiskIdGet`: EntityRisks
	fmt.Fprintf(os.Stdout, "Response from `EntityRisksAPI.EntityRisksEntityRiskIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityRiskId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntityRisksEntityRiskIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EntityRisks**](EntityRisks.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntityRisksEntityRiskIdPut

> EntityRisks EntityRisksEntityRiskIdPut(ctx, entityRiskId).EntityRisksPut(entityRisksPut).Execute()



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
	entityRiskId := int64(789) // int64 | Model id
	entityRisksPut := *openapiclient.NewEntityRisksPut() // EntityRisksPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntityRisksAPI.EntityRisksEntityRiskIdPut(context.Background(), entityRiskId).EntityRisksPut(entityRisksPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntityRisksAPI.EntityRisksEntityRiskIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntityRisksEntityRiskIdPut`: EntityRisks
	fmt.Fprintf(os.Stdout, "Response from `EntityRisksAPI.EntityRisksEntityRiskIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityRiskId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntityRisksEntityRiskIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entityRisksPut** | [**EntityRisksPut**](EntityRisksPut.md) |  | 

### Return type

[**EntityRisks**](EntityRisks.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntityRisksGet

> EntityRisksGet200Response EntityRisksGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.EntityRisksAPI.EntityRisksGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntityRisksAPI.EntityRisksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntityRisksGet`: EntityRisksGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EntityRisksAPI.EntityRisksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntityRisksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**EntityRisksGet200Response**](EntityRisksGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

