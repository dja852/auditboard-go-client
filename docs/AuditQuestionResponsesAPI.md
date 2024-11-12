# \AuditQuestionResponsesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditQuestionResponsesAuditQuestionResponseIdGet**](AuditQuestionResponsesAPI.md#AuditQuestionResponsesAuditQuestionResponseIdGet) | **Get** /audit_question_responses/{audit_question_response_id} | 
[**AuditQuestionResponsesAuditQuestionResponseIdPut**](AuditQuestionResponsesAPI.md#AuditQuestionResponsesAuditQuestionResponseIdPut) | **Put** /audit_question_responses/{audit_question_response_id} | 
[**AuditQuestionResponsesGet**](AuditQuestionResponsesAPI.md#AuditQuestionResponsesGet) | **Get** /audit_question_responses | 



## AuditQuestionResponsesAuditQuestionResponseIdGet

> AuditQuestionResponses AuditQuestionResponsesAuditQuestionResponseIdGet(ctx, auditQuestionResponseId).Include(include).Execute()



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
	auditQuestionResponseId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditQuestionResponsesAPI.AuditQuestionResponsesAuditQuestionResponseIdGet(context.Background(), auditQuestionResponseId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditQuestionResponsesAPI.AuditQuestionResponsesAuditQuestionResponseIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditQuestionResponsesAuditQuestionResponseIdGet`: AuditQuestionResponses
	fmt.Fprintf(os.Stdout, "Response from `AuditQuestionResponsesAPI.AuditQuestionResponsesAuditQuestionResponseIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditQuestionResponseId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditQuestionResponsesAuditQuestionResponseIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditQuestionResponses**](AuditQuestionResponses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditQuestionResponsesAuditQuestionResponseIdPut

> AuditQuestionResponses AuditQuestionResponsesAuditQuestionResponseIdPut(ctx, auditQuestionResponseId).AuditQuestionResponsesPut(auditQuestionResponsesPut).Execute()



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
	auditQuestionResponseId := int64(789) // int64 | Model id
	auditQuestionResponsesPut := *openapiclient.NewAuditQuestionResponsesPut() // AuditQuestionResponsesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditQuestionResponsesAPI.AuditQuestionResponsesAuditQuestionResponseIdPut(context.Background(), auditQuestionResponseId).AuditQuestionResponsesPut(auditQuestionResponsesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditQuestionResponsesAPI.AuditQuestionResponsesAuditQuestionResponseIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditQuestionResponsesAuditQuestionResponseIdPut`: AuditQuestionResponses
	fmt.Fprintf(os.Stdout, "Response from `AuditQuestionResponsesAPI.AuditQuestionResponsesAuditQuestionResponseIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditQuestionResponseId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditQuestionResponsesAuditQuestionResponseIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auditQuestionResponsesPut** | [**AuditQuestionResponsesPut**](AuditQuestionResponsesPut.md) |  | 

### Return type

[**AuditQuestionResponses**](AuditQuestionResponses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditQuestionResponsesGet

> AuditQuestionResponsesGet200Response AuditQuestionResponsesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.AuditQuestionResponsesAPI.AuditQuestionResponsesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditQuestionResponsesAPI.AuditQuestionResponsesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditQuestionResponsesGet`: AuditQuestionResponsesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditQuestionResponsesAPI.AuditQuestionResponsesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditQuestionResponsesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditQuestionResponsesGet200Response**](AuditQuestionResponsesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

