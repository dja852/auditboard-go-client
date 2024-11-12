# \RisksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RisksGet**](RisksAPI.md#RisksGet) | **Get** /risks | 
[**RisksPost**](RisksAPI.md#RisksPost) | **Post** /risks | 
[**RisksRiskIdDelete**](RisksAPI.md#RisksRiskIdDelete) | **Delete** /risks/{risk_id} | 
[**RisksRiskIdGet**](RisksAPI.md#RisksRiskIdGet) | **Get** /risks/{risk_id} | 
[**RisksRiskIdPut**](RisksAPI.md#RisksRiskIdPut) | **Put** /risks/{risk_id} | 



## RisksGet

> RisksGet200Response RisksGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.RisksAPI.RisksGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RisksAPI.RisksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RisksGet`: RisksGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RisksAPI.RisksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRisksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**RisksGet200Response**](RisksGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RisksPost

> RisksGet200Response RisksPost(ctx).RisksPostRequest(risksPostRequest).Execute()



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
	risksPostRequest := *openapiclient.NewRisksPostRequest() // RisksPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RisksAPI.RisksPost(context.Background()).RisksPostRequest(risksPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RisksAPI.RisksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RisksPost`: RisksGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RisksAPI.RisksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRisksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **risksPostRequest** | [**RisksPostRequest**](RisksPostRequest.md) |  | 

### Return type

[**RisksGet200Response**](RisksGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RisksRiskIdDelete

> Risks RisksRiskIdDelete(ctx, riskId).Execute()



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
	riskId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RisksAPI.RisksRiskIdDelete(context.Background(), riskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RisksAPI.RisksRiskIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RisksRiskIdDelete`: Risks
	fmt.Fprintf(os.Stdout, "Response from `RisksAPI.RisksRiskIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRisksRiskIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Risks**](Risks.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RisksRiskIdGet

> Risks RisksRiskIdGet(ctx, riskId).Include(include).Execute()



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
	riskId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RisksAPI.RisksRiskIdGet(context.Background(), riskId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RisksAPI.RisksRiskIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RisksRiskIdGet`: Risks
	fmt.Fprintf(os.Stdout, "Response from `RisksAPI.RisksRiskIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRisksRiskIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**Risks**](Risks.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RisksRiskIdPut

> Risks RisksRiskIdPut(ctx, riskId).RisksPut(risksPut).Execute()



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
	riskId := int64(789) // int64 | Model id
	risksPut := *openapiclient.NewRisksPut() // RisksPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RisksAPI.RisksRiskIdPut(context.Background(), riskId).RisksPut(risksPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RisksAPI.RisksRiskIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RisksRiskIdPut`: Risks
	fmt.Fprintf(os.Stdout, "Response from `RisksAPI.RisksRiskIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRisksRiskIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **risksPut** | [**RisksPut**](RisksPut.md) |  | 

### Return type

[**Risks**](Risks.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

