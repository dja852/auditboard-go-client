# \RiskTypesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RiskTypesGet**](RiskTypesAPI.md#RiskTypesGet) | **Get** /risk_types | 
[**RiskTypesPost**](RiskTypesAPI.md#RiskTypesPost) | **Post** /risk_types | 
[**RiskTypesRiskTypeIdDelete**](RiskTypesAPI.md#RiskTypesRiskTypeIdDelete) | **Delete** /risk_types/{risk_type_id} | 
[**RiskTypesRiskTypeIdGet**](RiskTypesAPI.md#RiskTypesRiskTypeIdGet) | **Get** /risk_types/{risk_type_id} | 
[**RiskTypesRiskTypeIdPut**](RiskTypesAPI.md#RiskTypesRiskTypeIdPut) | **Put** /risk_types/{risk_type_id} | 



## RiskTypesGet

> RiskTypesGet200Response RiskTypesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.RiskTypesAPI.RiskTypesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskTypesAPI.RiskTypesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskTypesGet`: RiskTypesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RiskTypesAPI.RiskTypesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskTypesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**RiskTypesGet200Response**](RiskTypesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskTypesPost

> RiskTypesGet200Response RiskTypesPost(ctx).RiskTypesPostRequest(riskTypesPostRequest).Execute()



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
	riskTypesPostRequest := *openapiclient.NewRiskTypesPostRequest() // RiskTypesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskTypesAPI.RiskTypesPost(context.Background()).RiskTypesPostRequest(riskTypesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskTypesAPI.RiskTypesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskTypesPost`: RiskTypesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RiskTypesAPI.RiskTypesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskTypesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskTypesPostRequest** | [**RiskTypesPostRequest**](RiskTypesPostRequest.md) |  | 

### Return type

[**RiskTypesGet200Response**](RiskTypesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskTypesRiskTypeIdDelete

> RiskTypes RiskTypesRiskTypeIdDelete(ctx, riskTypeId).Execute()



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
	riskTypeId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskTypesAPI.RiskTypesRiskTypeIdDelete(context.Background(), riskTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskTypesAPI.RiskTypesRiskTypeIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskTypesRiskTypeIdDelete`: RiskTypes
	fmt.Fprintf(os.Stdout, "Response from `RiskTypesAPI.RiskTypesRiskTypeIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskTypeId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskTypesRiskTypeIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RiskTypes**](RiskTypes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskTypesRiskTypeIdGet

> RiskTypes RiskTypesRiskTypeIdGet(ctx, riskTypeId).Include(include).Execute()



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
	riskTypeId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskTypesAPI.RiskTypesRiskTypeIdGet(context.Background(), riskTypeId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskTypesAPI.RiskTypesRiskTypeIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskTypesRiskTypeIdGet`: RiskTypes
	fmt.Fprintf(os.Stdout, "Response from `RiskTypesAPI.RiskTypesRiskTypeIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskTypeId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskTypesRiskTypeIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**RiskTypes**](RiskTypes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskTypesRiskTypeIdPut

> RiskTypes RiskTypesRiskTypeIdPut(ctx, riskTypeId).RiskTypesPut(riskTypesPut).Execute()



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
	riskTypeId := int64(789) // int64 | Model id
	riskTypesPut := *openapiclient.NewRiskTypesPut() // RiskTypesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskTypesAPI.RiskTypesRiskTypeIdPut(context.Background(), riskTypeId).RiskTypesPut(riskTypesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskTypesAPI.RiskTypesRiskTypeIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskTypesRiskTypeIdPut`: RiskTypes
	fmt.Fprintf(os.Stdout, "Response from `RiskTypesAPI.RiskTypesRiskTypeIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskTypeId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskTypesRiskTypeIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **riskTypesPut** | [**RiskTypesPut**](RiskTypesPut.md) |  | 

### Return type

[**RiskTypes**](RiskTypes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

