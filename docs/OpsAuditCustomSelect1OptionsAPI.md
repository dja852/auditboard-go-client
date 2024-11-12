# \OpsAuditCustomSelect1OptionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpsAuditCustomSelect1OptionsGet**](OpsAuditCustomSelect1OptionsAPI.md#OpsAuditCustomSelect1OptionsGet) | **Get** /ops_audit_custom_select1_options | 
[**OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdDelete**](OpsAuditCustomSelect1OptionsAPI.md#OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdDelete) | **Delete** /ops_audit_custom_select1_options/{ops_audit_custom_select1_option_id} | 
[**OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdGet**](OpsAuditCustomSelect1OptionsAPI.md#OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdGet) | **Get** /ops_audit_custom_select1_options/{ops_audit_custom_select1_option_id} | 
[**OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdPut**](OpsAuditCustomSelect1OptionsAPI.md#OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdPut) | **Put** /ops_audit_custom_select1_options/{ops_audit_custom_select1_option_id} | 
[**OpsAuditCustomSelect1OptionsPost**](OpsAuditCustomSelect1OptionsAPI.md#OpsAuditCustomSelect1OptionsPost) | **Post** /ops_audit_custom_select1_options | 



## OpsAuditCustomSelect1OptionsGet

> OpsAuditCustomSelect1OptionsGet200Response OpsAuditCustomSelect1OptionsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.OpsAuditCustomSelect1OptionsAPI.OpsAuditCustomSelect1OptionsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditCustomSelect1OptionsAPI.OpsAuditCustomSelect1OptionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditCustomSelect1OptionsGet`: OpsAuditCustomSelect1OptionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditCustomSelect1OptionsAPI.OpsAuditCustomSelect1OptionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditCustomSelect1OptionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**OpsAuditCustomSelect1OptionsGet200Response**](OpsAuditCustomSelect1OptionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdDelete

> OpsAuditCustomSelect1Options OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdDelete(ctx, opsAuditCustomSelect1OptionId).Execute()



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
	opsAuditCustomSelect1OptionId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditCustomSelect1OptionsAPI.OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdDelete(context.Background(), opsAuditCustomSelect1OptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditCustomSelect1OptionsAPI.OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdDelete`: OpsAuditCustomSelect1Options
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditCustomSelect1OptionsAPI.OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditCustomSelect1OptionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpsAuditCustomSelect1Options**](OpsAuditCustomSelect1Options.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdGet

> OpsAuditCustomSelect1Options OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdGet(ctx, opsAuditCustomSelect1OptionId).Include(include).Execute()



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
	opsAuditCustomSelect1OptionId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditCustomSelect1OptionsAPI.OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdGet(context.Background(), opsAuditCustomSelect1OptionId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditCustomSelect1OptionsAPI.OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdGet`: OpsAuditCustomSelect1Options
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditCustomSelect1OptionsAPI.OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditCustomSelect1OptionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**OpsAuditCustomSelect1Options**](OpsAuditCustomSelect1Options.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdPut

> OpsAuditCustomSelect1Options OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdPut(ctx, opsAuditCustomSelect1OptionId).OpsAuditCustomSelect1OptionsPut(opsAuditCustomSelect1OptionsPut).Execute()



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
	opsAuditCustomSelect1OptionId := int64(789) // int64 | Model id
	opsAuditCustomSelect1OptionsPut := *openapiclient.NewOpsAuditCustomSelect1OptionsPut() // OpsAuditCustomSelect1OptionsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditCustomSelect1OptionsAPI.OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdPut(context.Background(), opsAuditCustomSelect1OptionId).OpsAuditCustomSelect1OptionsPut(opsAuditCustomSelect1OptionsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditCustomSelect1OptionsAPI.OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdPut`: OpsAuditCustomSelect1Options
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditCustomSelect1OptionsAPI.OpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditCustomSelect1OptionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditCustomSelect1OptionsOpsAuditCustomSelect1OptionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opsAuditCustomSelect1OptionsPut** | [**OpsAuditCustomSelect1OptionsPut**](OpsAuditCustomSelect1OptionsPut.md) |  | 

### Return type

[**OpsAuditCustomSelect1Options**](OpsAuditCustomSelect1Options.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditCustomSelect1OptionsPost

> OpsAuditCustomSelect1OptionsGet200Response OpsAuditCustomSelect1OptionsPost(ctx).OpsAuditCustomSelect1OptionsPostRequest(opsAuditCustomSelect1OptionsPostRequest).Execute()



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
	opsAuditCustomSelect1OptionsPostRequest := *openapiclient.NewOpsAuditCustomSelect1OptionsPostRequest() // OpsAuditCustomSelect1OptionsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditCustomSelect1OptionsAPI.OpsAuditCustomSelect1OptionsPost(context.Background()).OpsAuditCustomSelect1OptionsPostRequest(opsAuditCustomSelect1OptionsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditCustomSelect1OptionsAPI.OpsAuditCustomSelect1OptionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditCustomSelect1OptionsPost`: OpsAuditCustomSelect1OptionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditCustomSelect1OptionsAPI.OpsAuditCustomSelect1OptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditCustomSelect1OptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opsAuditCustomSelect1OptionsPostRequest** | [**OpsAuditCustomSelect1OptionsPostRequest**](OpsAuditCustomSelect1OptionsPostRequest.md) |  | 

### Return type

[**OpsAuditCustomSelect1OptionsGet200Response**](OpsAuditCustomSelect1OptionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

