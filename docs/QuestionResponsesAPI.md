# \QuestionResponsesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QuestionResponsesGet**](QuestionResponsesAPI.md#QuestionResponsesGet) | **Get** /question_responses | 
[**QuestionResponsesPost**](QuestionResponsesAPI.md#QuestionResponsesPost) | **Post** /question_responses | 
[**QuestionResponsesQuestionResponseIdDelete**](QuestionResponsesAPI.md#QuestionResponsesQuestionResponseIdDelete) | **Delete** /question_responses/{question_response_id} | 
[**QuestionResponsesQuestionResponseIdGet**](QuestionResponsesAPI.md#QuestionResponsesQuestionResponseIdGet) | **Get** /question_responses/{question_response_id} | 
[**QuestionResponsesQuestionResponseIdPut**](QuestionResponsesAPI.md#QuestionResponsesQuestionResponseIdPut) | **Put** /question_responses/{question_response_id} | 



## QuestionResponsesGet

> QuestionResponsesGet200Response QuestionResponsesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.QuestionResponsesAPI.QuestionResponsesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuestionResponsesAPI.QuestionResponsesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuestionResponsesGet`: QuestionResponsesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `QuestionResponsesAPI.QuestionResponsesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuestionResponsesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**QuestionResponsesGet200Response**](QuestionResponsesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuestionResponsesPost

> QuestionResponsesGet200Response QuestionResponsesPost(ctx).QuestionResponsesPostRequest(questionResponsesPostRequest).Execute()



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
	questionResponsesPostRequest := *openapiclient.NewQuestionResponsesPostRequest() // QuestionResponsesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuestionResponsesAPI.QuestionResponsesPost(context.Background()).QuestionResponsesPostRequest(questionResponsesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuestionResponsesAPI.QuestionResponsesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuestionResponsesPost`: QuestionResponsesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `QuestionResponsesAPI.QuestionResponsesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuestionResponsesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questionResponsesPostRequest** | [**QuestionResponsesPostRequest**](QuestionResponsesPostRequest.md) |  | 

### Return type

[**QuestionResponsesGet200Response**](QuestionResponsesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuestionResponsesQuestionResponseIdDelete

> QuestionResponses QuestionResponsesQuestionResponseIdDelete(ctx, questionResponseId).Execute()



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
	questionResponseId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuestionResponsesAPI.QuestionResponsesQuestionResponseIdDelete(context.Background(), questionResponseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuestionResponsesAPI.QuestionResponsesQuestionResponseIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuestionResponsesQuestionResponseIdDelete`: QuestionResponses
	fmt.Fprintf(os.Stdout, "Response from `QuestionResponsesAPI.QuestionResponsesQuestionResponseIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionResponseId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuestionResponsesQuestionResponseIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuestionResponses**](QuestionResponses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuestionResponsesQuestionResponseIdGet

> QuestionResponses QuestionResponsesQuestionResponseIdGet(ctx, questionResponseId).Include(include).Execute()



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
	questionResponseId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuestionResponsesAPI.QuestionResponsesQuestionResponseIdGet(context.Background(), questionResponseId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuestionResponsesAPI.QuestionResponsesQuestionResponseIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuestionResponsesQuestionResponseIdGet`: QuestionResponses
	fmt.Fprintf(os.Stdout, "Response from `QuestionResponsesAPI.QuestionResponsesQuestionResponseIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionResponseId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuestionResponsesQuestionResponseIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**QuestionResponses**](QuestionResponses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuestionResponsesQuestionResponseIdPut

> QuestionResponses QuestionResponsesQuestionResponseIdPut(ctx, questionResponseId).Body(body).Execute()



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
	questionResponseId := int64(789) // int64 | Model id
	body := interface{}(987) // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuestionResponsesAPI.QuestionResponsesQuestionResponseIdPut(context.Background(), questionResponseId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuestionResponsesAPI.QuestionResponsesQuestionResponseIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuestionResponsesQuestionResponseIdPut`: QuestionResponses
	fmt.Fprintf(os.Stdout, "Response from `QuestionResponsesAPI.QuestionResponsesQuestionResponseIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionResponseId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuestionResponsesQuestionResponseIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** |  | 

### Return type

[**QuestionResponses**](QuestionResponses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

