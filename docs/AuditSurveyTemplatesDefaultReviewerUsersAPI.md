# \AuditSurveyTemplatesDefaultReviewerUsersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdDelete**](AuditSurveyTemplatesDefaultReviewerUsersAPI.md#AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdDelete) | **Delete** /audit_survey_templates_default_reviewer_users/{audit_survey_templates_default_reviewer_user_id} | 
[**AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdGet**](AuditSurveyTemplatesDefaultReviewerUsersAPI.md#AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdGet) | **Get** /audit_survey_templates_default_reviewer_users/{audit_survey_templates_default_reviewer_user_id} | 
[**AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdPut**](AuditSurveyTemplatesDefaultReviewerUsersAPI.md#AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdPut) | **Put** /audit_survey_templates_default_reviewer_users/{audit_survey_templates_default_reviewer_user_id} | 
[**AuditSurveyTemplatesDefaultReviewerUsersGet**](AuditSurveyTemplatesDefaultReviewerUsersAPI.md#AuditSurveyTemplatesDefaultReviewerUsersGet) | **Get** /audit_survey_templates_default_reviewer_users | 
[**AuditSurveyTemplatesDefaultReviewerUsersPost**](AuditSurveyTemplatesDefaultReviewerUsersAPI.md#AuditSurveyTemplatesDefaultReviewerUsersPost) | **Post** /audit_survey_templates_default_reviewer_users | 



## AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdDelete

> AuditSurveyTemplatesDefaultReviewerUsers AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdDelete(ctx, auditSurveyTemplatesDefaultReviewerUserId).Execute()



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
	auditSurveyTemplatesDefaultReviewerUserId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditSurveyTemplatesDefaultReviewerUsersAPI.AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdDelete(context.Background(), auditSurveyTemplatesDefaultReviewerUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditSurveyTemplatesDefaultReviewerUsersAPI.AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdDelete`: AuditSurveyTemplatesDefaultReviewerUsers
	fmt.Fprintf(os.Stdout, "Response from `AuditSurveyTemplatesDefaultReviewerUsersAPI.AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditSurveyTemplatesDefaultReviewerUserId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuditSurveyTemplatesDefaultReviewerUsers**](AuditSurveyTemplatesDefaultReviewerUsers.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdGet

> AuditSurveyTemplatesDefaultReviewerUsers AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdGet(ctx, auditSurveyTemplatesDefaultReviewerUserId).Include(include).Execute()



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
	auditSurveyTemplatesDefaultReviewerUserId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditSurveyTemplatesDefaultReviewerUsersAPI.AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdGet(context.Background(), auditSurveyTemplatesDefaultReviewerUserId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditSurveyTemplatesDefaultReviewerUsersAPI.AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdGet`: AuditSurveyTemplatesDefaultReviewerUsers
	fmt.Fprintf(os.Stdout, "Response from `AuditSurveyTemplatesDefaultReviewerUsersAPI.AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditSurveyTemplatesDefaultReviewerUserId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditSurveyTemplatesDefaultReviewerUsers**](AuditSurveyTemplatesDefaultReviewerUsers.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdPut

> AuditSurveyTemplatesDefaultReviewerUsers AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdPut(ctx, auditSurveyTemplatesDefaultReviewerUserId).AuditSurveyTemplatesDefaultReviewerUsersPut(auditSurveyTemplatesDefaultReviewerUsersPut).Execute()



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
	auditSurveyTemplatesDefaultReviewerUserId := int64(789) // int64 | Model id
	auditSurveyTemplatesDefaultReviewerUsersPut := *openapiclient.NewAuditSurveyTemplatesDefaultReviewerUsersPut() // AuditSurveyTemplatesDefaultReviewerUsersPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditSurveyTemplatesDefaultReviewerUsersAPI.AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdPut(context.Background(), auditSurveyTemplatesDefaultReviewerUserId).AuditSurveyTemplatesDefaultReviewerUsersPut(auditSurveyTemplatesDefaultReviewerUsersPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditSurveyTemplatesDefaultReviewerUsersAPI.AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdPut`: AuditSurveyTemplatesDefaultReviewerUsers
	fmt.Fprintf(os.Stdout, "Response from `AuditSurveyTemplatesDefaultReviewerUsersAPI.AuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditSurveyTemplatesDefaultReviewerUserId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditSurveyTemplatesDefaultReviewerUsersAuditSurveyTemplatesDefaultReviewerUserIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auditSurveyTemplatesDefaultReviewerUsersPut** | [**AuditSurveyTemplatesDefaultReviewerUsersPut**](AuditSurveyTemplatesDefaultReviewerUsersPut.md) |  | 

### Return type

[**AuditSurveyTemplatesDefaultReviewerUsers**](AuditSurveyTemplatesDefaultReviewerUsers.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditSurveyTemplatesDefaultReviewerUsersGet

> AuditSurveyTemplatesDefaultReviewerUsersGet200Response AuditSurveyTemplatesDefaultReviewerUsersGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.AuditSurveyTemplatesDefaultReviewerUsersAPI.AuditSurveyTemplatesDefaultReviewerUsersGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditSurveyTemplatesDefaultReviewerUsersAPI.AuditSurveyTemplatesDefaultReviewerUsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditSurveyTemplatesDefaultReviewerUsersGet`: AuditSurveyTemplatesDefaultReviewerUsersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditSurveyTemplatesDefaultReviewerUsersAPI.AuditSurveyTemplatesDefaultReviewerUsersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditSurveyTemplatesDefaultReviewerUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditSurveyTemplatesDefaultReviewerUsersGet200Response**](AuditSurveyTemplatesDefaultReviewerUsersGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditSurveyTemplatesDefaultReviewerUsersPost

> AuditSurveyTemplatesDefaultReviewerUsersGet200Response AuditSurveyTemplatesDefaultReviewerUsersPost(ctx).AuditSurveyTemplatesDefaultReviewerUsersPostRequest(auditSurveyTemplatesDefaultReviewerUsersPostRequest).Execute()



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
	auditSurveyTemplatesDefaultReviewerUsersPostRequest := *openapiclient.NewAuditSurveyTemplatesDefaultReviewerUsersPostRequest() // AuditSurveyTemplatesDefaultReviewerUsersPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditSurveyTemplatesDefaultReviewerUsersAPI.AuditSurveyTemplatesDefaultReviewerUsersPost(context.Background()).AuditSurveyTemplatesDefaultReviewerUsersPostRequest(auditSurveyTemplatesDefaultReviewerUsersPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditSurveyTemplatesDefaultReviewerUsersAPI.AuditSurveyTemplatesDefaultReviewerUsersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditSurveyTemplatesDefaultReviewerUsersPost`: AuditSurveyTemplatesDefaultReviewerUsersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditSurveyTemplatesDefaultReviewerUsersAPI.AuditSurveyTemplatesDefaultReviewerUsersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditSurveyTemplatesDefaultReviewerUsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auditSurveyTemplatesDefaultReviewerUsersPostRequest** | [**AuditSurveyTemplatesDefaultReviewerUsersPostRequest**](AuditSurveyTemplatesDefaultReviewerUsersPostRequest.md) |  | 

### Return type

[**AuditSurveyTemplatesDefaultReviewerUsersGet200Response**](AuditSurveyTemplatesDefaultReviewerUsersGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

