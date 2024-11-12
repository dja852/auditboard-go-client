# \KeyPerformanceIndicatorsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePost**](KeyPerformanceIndicatorsAPI.md#KeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePost) | **Post** /key_performance_indicators/{key_performance_indicator_id}/add_value | 



## KeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePost

> KeyPerformanceIndicatorsAddValue KeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePost(ctx, keyPerformanceIndicatorId).KeyPerformanceIndicatorAddValuePost(keyPerformanceIndicatorAddValuePost).Execute()



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
	keyPerformanceIndicatorId := int64(789) // int64 | Key Performance Indicator Id
	keyPerformanceIndicatorAddValuePost := *openapiclient.NewKeyPerformanceIndicatorAddValuePost(float32(123), float32(123), "HistoricalDate_example") // KeyPerformanceIndicatorAddValuePost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyPerformanceIndicatorsAPI.KeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePost(context.Background(), keyPerformanceIndicatorId).KeyPerformanceIndicatorAddValuePost(keyPerformanceIndicatorAddValuePost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyPerformanceIndicatorsAPI.KeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePost`: KeyPerformanceIndicatorsAddValue
	fmt.Fprintf(os.Stdout, "Response from `KeyPerformanceIndicatorsAPI.KeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyPerformanceIndicatorId** | **int64** | Key Performance Indicator Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keyPerformanceIndicatorAddValuePost** | [**KeyPerformanceIndicatorAddValuePost**](KeyPerformanceIndicatorAddValuePost.md) |  | 

### Return type

[**KeyPerformanceIndicatorsAddValue**](KeyPerformanceIndicatorsAddValue.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

