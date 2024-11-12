# \AuditSurveyTemplatesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditSurveyTemplatesAuditSurveyTemplateIdDelete**](AuditSurveyTemplatesAPI.md#AuditSurveyTemplatesAuditSurveyTemplateIdDelete) | **Delete** /audit_survey_templates/{audit_survey_template_id} | 
[**AuditSurveyTemplatesAuditSurveyTemplateIdGet**](AuditSurveyTemplatesAPI.md#AuditSurveyTemplatesAuditSurveyTemplateIdGet) | **Get** /audit_survey_templates/{audit_survey_template_id} | 
[**AuditSurveyTemplatesAuditSurveyTemplateIdPut**](AuditSurveyTemplatesAPI.md#AuditSurveyTemplatesAuditSurveyTemplateIdPut) | **Put** /audit_survey_templates/{audit_survey_template_id} | 
[**AuditSurveyTemplatesGet**](AuditSurveyTemplatesAPI.md#AuditSurveyTemplatesGet) | **Get** /audit_survey_templates | 
[**AuditSurveyTemplatesPost**](AuditSurveyTemplatesAPI.md#AuditSurveyTemplatesPost) | **Post** /audit_survey_templates | 



## AuditSurveyTemplatesAuditSurveyTemplateIdDelete

> AuditSurveyTemplates AuditSurveyTemplatesAuditSurveyTemplateIdDelete(ctx, auditSurveyTemplateId).Execute()



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
	auditSurveyTemplateId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditSurveyTemplatesAPI.AuditSurveyTemplatesAuditSurveyTemplateIdDelete(context.Background(), auditSurveyTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditSurveyTemplatesAPI.AuditSurveyTemplatesAuditSurveyTemplateIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditSurveyTemplatesAuditSurveyTemplateIdDelete`: AuditSurveyTemplates
	fmt.Fprintf(os.Stdout, "Response from `AuditSurveyTemplatesAPI.AuditSurveyTemplatesAuditSurveyTemplateIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditSurveyTemplateId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditSurveyTemplatesAuditSurveyTemplateIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuditSurveyTemplates**](AuditSurveyTemplates.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditSurveyTemplatesAuditSurveyTemplateIdGet

> AuditSurveyTemplates AuditSurveyTemplatesAuditSurveyTemplateIdGet(ctx, auditSurveyTemplateId).Include(include).Execute()



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
	auditSurveyTemplateId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditSurveyTemplatesAPI.AuditSurveyTemplatesAuditSurveyTemplateIdGet(context.Background(), auditSurveyTemplateId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditSurveyTemplatesAPI.AuditSurveyTemplatesAuditSurveyTemplateIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditSurveyTemplatesAuditSurveyTemplateIdGet`: AuditSurveyTemplates
	fmt.Fprintf(os.Stdout, "Response from `AuditSurveyTemplatesAPI.AuditSurveyTemplatesAuditSurveyTemplateIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditSurveyTemplateId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditSurveyTemplatesAuditSurveyTemplateIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditSurveyTemplates**](AuditSurveyTemplates.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditSurveyTemplatesAuditSurveyTemplateIdPut

> AuditSurveyTemplates AuditSurveyTemplatesAuditSurveyTemplateIdPut(ctx, auditSurveyTemplateId).AuditSurveyTemplatesPut(auditSurveyTemplatesPut).Execute()



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
	auditSurveyTemplateId := int64(789) // int64 | Model id
	auditSurveyTemplatesPut := *openapiclient.NewAuditSurveyTemplatesPut() // AuditSurveyTemplatesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditSurveyTemplatesAPI.AuditSurveyTemplatesAuditSurveyTemplateIdPut(context.Background(), auditSurveyTemplateId).AuditSurveyTemplatesPut(auditSurveyTemplatesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditSurveyTemplatesAPI.AuditSurveyTemplatesAuditSurveyTemplateIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditSurveyTemplatesAuditSurveyTemplateIdPut`: AuditSurveyTemplates
	fmt.Fprintf(os.Stdout, "Response from `AuditSurveyTemplatesAPI.AuditSurveyTemplatesAuditSurveyTemplateIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditSurveyTemplateId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditSurveyTemplatesAuditSurveyTemplateIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auditSurveyTemplatesPut** | [**AuditSurveyTemplatesPut**](AuditSurveyTemplatesPut.md) |  | 

### Return type

[**AuditSurveyTemplates**](AuditSurveyTemplates.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditSurveyTemplatesGet

> AuditSurveyTemplatesGet200Response AuditSurveyTemplatesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.AuditSurveyTemplatesAPI.AuditSurveyTemplatesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditSurveyTemplatesAPI.AuditSurveyTemplatesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditSurveyTemplatesGet`: AuditSurveyTemplatesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditSurveyTemplatesAPI.AuditSurveyTemplatesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditSurveyTemplatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditSurveyTemplatesGet200Response**](AuditSurveyTemplatesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditSurveyTemplatesPost

> AuditSurveyTemplatesGet200Response AuditSurveyTemplatesPost(ctx).AuditSurveyTemplatesPostRequest(auditSurveyTemplatesPostRequest).Execute()



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
	auditSurveyTemplatesPostRequest := *openapiclient.NewAuditSurveyTemplatesPostRequest() // AuditSurveyTemplatesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditSurveyTemplatesAPI.AuditSurveyTemplatesPost(context.Background()).AuditSurveyTemplatesPostRequest(auditSurveyTemplatesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditSurveyTemplatesAPI.AuditSurveyTemplatesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditSurveyTemplatesPost`: AuditSurveyTemplatesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditSurveyTemplatesAPI.AuditSurveyTemplatesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditSurveyTemplatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auditSurveyTemplatesPostRequest** | [**AuditSurveyTemplatesPostRequest**](AuditSurveyTemplatesPostRequest.md) |  | 

### Return type

[**AuditSurveyTemplatesGet200Response**](AuditSurveyTemplatesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

