# \KeyRiskIndicatorsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KeyRiskIndicatorsKeyRiskIndicatorIdAddValuePost**](KeyRiskIndicatorsAPI.md#KeyRiskIndicatorsKeyRiskIndicatorIdAddValuePost) | **Post** /key_risk_indicators/{key_risk_indicator_id}/add_value | 



## KeyRiskIndicatorsKeyRiskIndicatorIdAddValuePost

> KeyRiskIndicatorsAddValue KeyRiskIndicatorsKeyRiskIndicatorIdAddValuePost(ctx, keyRiskIndicatorId).KeyRiskIndicatorAddValuePost(keyRiskIndicatorAddValuePost).Execute()



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
	keyRiskIndicatorId := int64(789) // int64 | Key Risk Indicator Id
	keyRiskIndicatorAddValuePost := *openapiclient.NewKeyRiskIndicatorAddValuePost(float32(123), "HistoricalDate_example") // KeyRiskIndicatorAddValuePost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyRiskIndicatorsAPI.KeyRiskIndicatorsKeyRiskIndicatorIdAddValuePost(context.Background(), keyRiskIndicatorId).KeyRiskIndicatorAddValuePost(keyRiskIndicatorAddValuePost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyRiskIndicatorsAPI.KeyRiskIndicatorsKeyRiskIndicatorIdAddValuePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeyRiskIndicatorsKeyRiskIndicatorIdAddValuePost`: KeyRiskIndicatorsAddValue
	fmt.Fprintf(os.Stdout, "Response from `KeyRiskIndicatorsAPI.KeyRiskIndicatorsKeyRiskIndicatorIdAddValuePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyRiskIndicatorId** | **int64** | Key Risk Indicator Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeyRiskIndicatorsKeyRiskIndicatorIdAddValuePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keyRiskIndicatorAddValuePost** | [**KeyRiskIndicatorAddValuePost**](KeyRiskIndicatorAddValuePost.md) |  | 

### Return type

[**KeyRiskIndicatorsAddValue**](KeyRiskIndicatorsAddValue.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

