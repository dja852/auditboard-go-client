# \CustomFieldsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomFieldsCustomFieldIdDelete**](CustomFieldsAPI.md#CustomFieldsCustomFieldIdDelete) | **Delete** /custom_fields/{custom_field_id} | 
[**CustomFieldsCustomFieldIdGet**](CustomFieldsAPI.md#CustomFieldsCustomFieldIdGet) | **Get** /custom_fields/{custom_field_id} | 
[**CustomFieldsCustomFieldIdPut**](CustomFieldsAPI.md#CustomFieldsCustomFieldIdPut) | **Put** /custom_fields/{custom_field_id} | 
[**CustomFieldsGet**](CustomFieldsAPI.md#CustomFieldsGet) | **Get** /custom_fields | 
[**CustomFieldsPost**](CustomFieldsAPI.md#CustomFieldsPost) | **Post** /custom_fields | 



## CustomFieldsCustomFieldIdDelete

> CustomFields CustomFieldsCustomFieldIdDelete(ctx, customFieldId).Execute()



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
	customFieldId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldsAPI.CustomFieldsCustomFieldIdDelete(context.Background(), customFieldId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsAPI.CustomFieldsCustomFieldIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomFieldsCustomFieldIdDelete`: CustomFields
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldsAPI.CustomFieldsCustomFieldIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customFieldId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomFieldsCustomFieldIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomFields**](CustomFields.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomFieldsCustomFieldIdGet

> CustomFields CustomFieldsCustomFieldIdGet(ctx, customFieldId).Include(include).Execute()



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
	customFieldId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldsAPI.CustomFieldsCustomFieldIdGet(context.Background(), customFieldId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsAPI.CustomFieldsCustomFieldIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomFieldsCustomFieldIdGet`: CustomFields
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldsAPI.CustomFieldsCustomFieldIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customFieldId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomFieldsCustomFieldIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**CustomFields**](CustomFields.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomFieldsCustomFieldIdPut

> CustomFields CustomFieldsCustomFieldIdPut(ctx, customFieldId).CustomFieldsPut(customFieldsPut).Execute()



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
	customFieldId := int64(789) // int64 | Model id
	customFieldsPut := *openapiclient.NewCustomFieldsPut() // CustomFieldsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldsAPI.CustomFieldsCustomFieldIdPut(context.Background(), customFieldId).CustomFieldsPut(customFieldsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsAPI.CustomFieldsCustomFieldIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomFieldsCustomFieldIdPut`: CustomFields
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldsAPI.CustomFieldsCustomFieldIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customFieldId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomFieldsCustomFieldIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customFieldsPut** | [**CustomFieldsPut**](CustomFieldsPut.md) |  | 

### Return type

[**CustomFields**](CustomFields.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomFieldsGet

> CustomFieldsGet200Response CustomFieldsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.CustomFieldsAPI.CustomFieldsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsAPI.CustomFieldsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomFieldsGet`: CustomFieldsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldsAPI.CustomFieldsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomFieldsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**CustomFieldsGet200Response**](CustomFieldsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomFieldsPost

> CustomFieldsGet200Response CustomFieldsPost(ctx).CustomFieldsPostRequest(customFieldsPostRequest).Execute()



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
	customFieldsPostRequest := *openapiclient.NewCustomFieldsPostRequest() // CustomFieldsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldsAPI.CustomFieldsPost(context.Background()).CustomFieldsPostRequest(customFieldsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsAPI.CustomFieldsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomFieldsPost`: CustomFieldsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldsAPI.CustomFieldsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomFieldsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customFieldsPostRequest** | [**CustomFieldsPostRequest**](CustomFieldsPostRequest.md) |  | 

### Return type

[**CustomFieldsGet200Response**](CustomFieldsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

