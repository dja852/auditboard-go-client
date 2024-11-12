# \FormTemplatesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FormTemplatesFormTemplateIdDelete**](FormTemplatesAPI.md#FormTemplatesFormTemplateIdDelete) | **Delete** /form_templates/{form_template_id} | 
[**FormTemplatesFormTemplateIdGet**](FormTemplatesAPI.md#FormTemplatesFormTemplateIdGet) | **Get** /form_templates/{form_template_id} | 
[**FormTemplatesFormTemplateIdPut**](FormTemplatesAPI.md#FormTemplatesFormTemplateIdPut) | **Put** /form_templates/{form_template_id} | 
[**FormTemplatesGet**](FormTemplatesAPI.md#FormTemplatesGet) | **Get** /form_templates | 
[**FormTemplatesPost**](FormTemplatesAPI.md#FormTemplatesPost) | **Post** /form_templates | 



## FormTemplatesFormTemplateIdDelete

> FormTemplates FormTemplatesFormTemplateIdDelete(ctx, formTemplateId).Execute()



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
	formTemplateId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormTemplatesAPI.FormTemplatesFormTemplateIdDelete(context.Background(), formTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormTemplatesAPI.FormTemplatesFormTemplateIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FormTemplatesFormTemplateIdDelete`: FormTemplates
	fmt.Fprintf(os.Stdout, "Response from `FormTemplatesAPI.FormTemplatesFormTemplateIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formTemplateId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFormTemplatesFormTemplateIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FormTemplates**](FormTemplates.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FormTemplatesFormTemplateIdGet

> FormTemplates FormTemplatesFormTemplateIdGet(ctx, formTemplateId).Execute()



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
	formTemplateId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormTemplatesAPI.FormTemplatesFormTemplateIdGet(context.Background(), formTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormTemplatesAPI.FormTemplatesFormTemplateIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FormTemplatesFormTemplateIdGet`: FormTemplates
	fmt.Fprintf(os.Stdout, "Response from `FormTemplatesAPI.FormTemplatesFormTemplateIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formTemplateId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFormTemplatesFormTemplateIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FormTemplates**](FormTemplates.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FormTemplatesFormTemplateIdPut

> FormTemplates FormTemplatesFormTemplateIdPut(ctx, formTemplateId).FormTemplatesPut(formTemplatesPut).Execute()



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
	formTemplateId := int64(789) // int64 | Model id
	formTemplatesPut := *openapiclient.NewFormTemplatesPut() // FormTemplatesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormTemplatesAPI.FormTemplatesFormTemplateIdPut(context.Background(), formTemplateId).FormTemplatesPut(formTemplatesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormTemplatesAPI.FormTemplatesFormTemplateIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FormTemplatesFormTemplateIdPut`: FormTemplates
	fmt.Fprintf(os.Stdout, "Response from `FormTemplatesAPI.FormTemplatesFormTemplateIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formTemplateId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFormTemplatesFormTemplateIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **formTemplatesPut** | [**FormTemplatesPut**](FormTemplatesPut.md) |  | 

### Return type

[**FormTemplates**](FormTemplates.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FormTemplatesGet

> FormTemplatesGet200Response FormTemplatesGet(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormTemplatesAPI.FormTemplatesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormTemplatesAPI.FormTemplatesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FormTemplatesGet`: FormTemplatesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FormTemplatesAPI.FormTemplatesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFormTemplatesGetRequest struct via the builder pattern


### Return type

[**FormTemplatesGet200Response**](FormTemplatesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FormTemplatesPost

> FormTemplatesGet200Response FormTemplatesPost(ctx).FormTemplatesPostRequest(formTemplatesPostRequest).Execute()



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
	formTemplatesPostRequest := *openapiclient.NewFormTemplatesPostRequest() // FormTemplatesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormTemplatesAPI.FormTemplatesPost(context.Background()).FormTemplatesPostRequest(formTemplatesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormTemplatesAPI.FormTemplatesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FormTemplatesPost`: FormTemplatesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FormTemplatesAPI.FormTemplatesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFormTemplatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **formTemplatesPostRequest** | [**FormTemplatesPostRequest**](FormTemplatesPostRequest.md) |  | 

### Return type

[**FormTemplatesGet200Response**](FormTemplatesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

