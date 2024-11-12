# \FinancialAccountsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FinancialAccountsFinancialAccountIdDelete**](FinancialAccountsAPI.md#FinancialAccountsFinancialAccountIdDelete) | **Delete** /financial_accounts/{financial_account_id} | 
[**FinancialAccountsFinancialAccountIdGet**](FinancialAccountsAPI.md#FinancialAccountsFinancialAccountIdGet) | **Get** /financial_accounts/{financial_account_id} | 
[**FinancialAccountsFinancialAccountIdPut**](FinancialAccountsAPI.md#FinancialAccountsFinancialAccountIdPut) | **Put** /financial_accounts/{financial_account_id} | 
[**FinancialAccountsGet**](FinancialAccountsAPI.md#FinancialAccountsGet) | **Get** /financial_accounts | 
[**FinancialAccountsPost**](FinancialAccountsAPI.md#FinancialAccountsPost) | **Post** /financial_accounts | 



## FinancialAccountsFinancialAccountIdDelete

> FinancialAccounts FinancialAccountsFinancialAccountIdDelete(ctx, financialAccountId).Execute()



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
	financialAccountId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FinancialAccountsAPI.FinancialAccountsFinancialAccountIdDelete(context.Background(), financialAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinancialAccountsAPI.FinancialAccountsFinancialAccountIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FinancialAccountsFinancialAccountIdDelete`: FinancialAccounts
	fmt.Fprintf(os.Stdout, "Response from `FinancialAccountsAPI.FinancialAccountsFinancialAccountIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**financialAccountId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFinancialAccountsFinancialAccountIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FinancialAccounts**](FinancialAccounts.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinancialAccountsFinancialAccountIdGet

> FinancialAccounts FinancialAccountsFinancialAccountIdGet(ctx, financialAccountId).Include(include).Execute()



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
	financialAccountId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FinancialAccountsAPI.FinancialAccountsFinancialAccountIdGet(context.Background(), financialAccountId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinancialAccountsAPI.FinancialAccountsFinancialAccountIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FinancialAccountsFinancialAccountIdGet`: FinancialAccounts
	fmt.Fprintf(os.Stdout, "Response from `FinancialAccountsAPI.FinancialAccountsFinancialAccountIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**financialAccountId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFinancialAccountsFinancialAccountIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**FinancialAccounts**](FinancialAccounts.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinancialAccountsFinancialAccountIdPut

> FinancialAccounts FinancialAccountsFinancialAccountIdPut(ctx, financialAccountId).FinancialAccountsPut(financialAccountsPut).Execute()



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
	financialAccountId := int64(789) // int64 | Model id
	financialAccountsPut := *openapiclient.NewFinancialAccountsPut() // FinancialAccountsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FinancialAccountsAPI.FinancialAccountsFinancialAccountIdPut(context.Background(), financialAccountId).FinancialAccountsPut(financialAccountsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinancialAccountsAPI.FinancialAccountsFinancialAccountIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FinancialAccountsFinancialAccountIdPut`: FinancialAccounts
	fmt.Fprintf(os.Stdout, "Response from `FinancialAccountsAPI.FinancialAccountsFinancialAccountIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**financialAccountId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFinancialAccountsFinancialAccountIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **financialAccountsPut** | [**FinancialAccountsPut**](FinancialAccountsPut.md) |  | 

### Return type

[**FinancialAccounts**](FinancialAccounts.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinancialAccountsGet

> FinancialAccountsGet200Response FinancialAccountsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.FinancialAccountsAPI.FinancialAccountsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinancialAccountsAPI.FinancialAccountsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FinancialAccountsGet`: FinancialAccountsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FinancialAccountsAPI.FinancialAccountsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFinancialAccountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**FinancialAccountsGet200Response**](FinancialAccountsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinancialAccountsPost

> FinancialAccountsGet200Response FinancialAccountsPost(ctx).FinancialAccountsPostRequest(financialAccountsPostRequest).Execute()



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
	financialAccountsPostRequest := *openapiclient.NewFinancialAccountsPostRequest() // FinancialAccountsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FinancialAccountsAPI.FinancialAccountsPost(context.Background()).FinancialAccountsPostRequest(financialAccountsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinancialAccountsAPI.FinancialAccountsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FinancialAccountsPost`: FinancialAccountsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FinancialAccountsAPI.FinancialAccountsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFinancialAccountsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **financialAccountsPostRequest** | [**FinancialAccountsPostRequest**](FinancialAccountsPostRequest.md) |  | 

### Return type

[**FinancialAccountsGet200Response**](FinancialAccountsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

