# \AuditQuestionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditQuestionsAuditQuestionIdGet**](AuditQuestionsAPI.md#AuditQuestionsAuditQuestionIdGet) | **Get** /audit_questions/{audit_question_id} | 
[**AuditQuestionsGet**](AuditQuestionsAPI.md#AuditQuestionsGet) | **Get** /audit_questions | 



## AuditQuestionsAuditQuestionIdGet

> AuditQuestions AuditQuestionsAuditQuestionIdGet(ctx, auditQuestionId).Include(include).Execute()



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
	auditQuestionId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditQuestionsAPI.AuditQuestionsAuditQuestionIdGet(context.Background(), auditQuestionId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditQuestionsAPI.AuditQuestionsAuditQuestionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditQuestionsAuditQuestionIdGet`: AuditQuestions
	fmt.Fprintf(os.Stdout, "Response from `AuditQuestionsAPI.AuditQuestionsAuditQuestionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditQuestionId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditQuestionsAuditQuestionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditQuestions**](AuditQuestions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditQuestionsGet

> AuditQuestionsGet200Response AuditQuestionsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.AuditQuestionsAPI.AuditQuestionsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditQuestionsAPI.AuditQuestionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditQuestionsGet`: AuditQuestionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditQuestionsAPI.AuditQuestionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditQuestionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditQuestionsGet200Response**](AuditQuestionsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

