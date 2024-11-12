# \AuditSurveysAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditSurveysAuditSurveyIdDelete**](AuditSurveysAPI.md#AuditSurveysAuditSurveyIdDelete) | **Delete** /audit_surveys/{audit_survey_id} | 
[**AuditSurveysAuditSurveyIdGet**](AuditSurveysAPI.md#AuditSurveysAuditSurveyIdGet) | **Get** /audit_surveys/{audit_survey_id} | 
[**AuditSurveysAuditSurveyIdPut**](AuditSurveysAPI.md#AuditSurveysAuditSurveyIdPut) | **Put** /audit_surveys/{audit_survey_id} | 
[**AuditSurveysGet**](AuditSurveysAPI.md#AuditSurveysGet) | **Get** /audit_surveys | 



## AuditSurveysAuditSurveyIdDelete

> AuditSurveys AuditSurveysAuditSurveyIdDelete(ctx, auditSurveyId).Execute()



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
	auditSurveyId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditSurveysAPI.AuditSurveysAuditSurveyIdDelete(context.Background(), auditSurveyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditSurveysAPI.AuditSurveysAuditSurveyIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditSurveysAuditSurveyIdDelete`: AuditSurveys
	fmt.Fprintf(os.Stdout, "Response from `AuditSurveysAPI.AuditSurveysAuditSurveyIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditSurveyId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditSurveysAuditSurveyIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuditSurveys**](AuditSurveys.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditSurveysAuditSurveyIdGet

> AuditSurveys AuditSurveysAuditSurveyIdGet(ctx, auditSurveyId).Include(include).Execute()



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
	auditSurveyId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditSurveysAPI.AuditSurveysAuditSurveyIdGet(context.Background(), auditSurveyId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditSurveysAPI.AuditSurveysAuditSurveyIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditSurveysAuditSurveyIdGet`: AuditSurveys
	fmt.Fprintf(os.Stdout, "Response from `AuditSurveysAPI.AuditSurveysAuditSurveyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditSurveyId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditSurveysAuditSurveyIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditSurveys**](AuditSurveys.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditSurveysAuditSurveyIdPut

> AuditSurveys AuditSurveysAuditSurveyIdPut(ctx, auditSurveyId).AuditSurveysPut(auditSurveysPut).Execute()



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
	auditSurveyId := int64(789) // int64 | Model id
	auditSurveysPut := *openapiclient.NewAuditSurveysPut() // AuditSurveysPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditSurveysAPI.AuditSurveysAuditSurveyIdPut(context.Background(), auditSurveyId).AuditSurveysPut(auditSurveysPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditSurveysAPI.AuditSurveysAuditSurveyIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditSurveysAuditSurveyIdPut`: AuditSurveys
	fmt.Fprintf(os.Stdout, "Response from `AuditSurveysAPI.AuditSurveysAuditSurveyIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditSurveyId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditSurveysAuditSurveyIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auditSurveysPut** | [**AuditSurveysPut**](AuditSurveysPut.md) |  | 

### Return type

[**AuditSurveys**](AuditSurveys.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditSurveysGet

> AuditSurveysGet200Response AuditSurveysGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.AuditSurveysAPI.AuditSurveysGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditSurveysAPI.AuditSurveysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditSurveysGet`: AuditSurveysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditSurveysAPI.AuditSurveysGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditSurveysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditSurveysGet200Response**](AuditSurveysGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

