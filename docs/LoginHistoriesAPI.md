# \LoginHistoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoginHistoriesGet**](LoginHistoriesAPI.md#LoginHistoriesGet) | **Get** /login_histories | 
[**LoginHistoriesLoginHistoryIdGet**](LoginHistoriesAPI.md#LoginHistoriesLoginHistoryIdGet) | **Get** /login_histories/{login_history_id} | 



## LoginHistoriesGet

> LoginHistoriesGet200Response LoginHistoriesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.LoginHistoriesAPI.LoginHistoriesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoginHistoriesAPI.LoginHistoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoginHistoriesGet`: LoginHistoriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `LoginHistoriesAPI.LoginHistoriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginHistoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**LoginHistoriesGet200Response**](LoginHistoriesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginHistoriesLoginHistoryIdGet

> LoginHistories LoginHistoriesLoginHistoryIdGet(ctx, loginHistoryId).Include(include).Execute()



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
	loginHistoryId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoginHistoriesAPI.LoginHistoriesLoginHistoryIdGet(context.Background(), loginHistoryId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoginHistoriesAPI.LoginHistoriesLoginHistoryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoginHistoriesLoginHistoryIdGet`: LoginHistories
	fmt.Fprintf(os.Stdout, "Response from `LoginHistoriesAPI.LoginHistoriesLoginHistoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loginHistoryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoginHistoriesLoginHistoryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**LoginHistories**](LoginHistories.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

