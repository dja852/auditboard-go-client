# \AssessmentResponsesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssessmentResponsesAssessmentResponseIdDelete**](AssessmentResponsesAPI.md#AssessmentResponsesAssessmentResponseIdDelete) | **Delete** /assessment_responses/{assessment_response_id} | 
[**AssessmentResponsesAssessmentResponseIdGet**](AssessmentResponsesAPI.md#AssessmentResponsesAssessmentResponseIdGet) | **Get** /assessment_responses/{assessment_response_id} | 
[**AssessmentResponsesAssessmentResponseIdPut**](AssessmentResponsesAPI.md#AssessmentResponsesAssessmentResponseIdPut) | **Put** /assessment_responses/{assessment_response_id} | 
[**AssessmentResponsesGet**](AssessmentResponsesAPI.md#AssessmentResponsesGet) | **Get** /assessment_responses | 



## AssessmentResponsesAssessmentResponseIdDelete

> AssessmentResponses AssessmentResponsesAssessmentResponseIdDelete(ctx, assessmentResponseId).Execute()



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
	assessmentResponseId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssessmentResponsesAPI.AssessmentResponsesAssessmentResponseIdDelete(context.Background(), assessmentResponseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssessmentResponsesAPI.AssessmentResponsesAssessmentResponseIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssessmentResponsesAssessmentResponseIdDelete`: AssessmentResponses
	fmt.Fprintf(os.Stdout, "Response from `AssessmentResponsesAPI.AssessmentResponsesAssessmentResponseIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assessmentResponseId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssessmentResponsesAssessmentResponseIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssessmentResponses**](AssessmentResponses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssessmentResponsesAssessmentResponseIdGet

> AssessmentResponses AssessmentResponsesAssessmentResponseIdGet(ctx, assessmentResponseId).Include(include).Execute()



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
	assessmentResponseId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssessmentResponsesAPI.AssessmentResponsesAssessmentResponseIdGet(context.Background(), assessmentResponseId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssessmentResponsesAPI.AssessmentResponsesAssessmentResponseIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssessmentResponsesAssessmentResponseIdGet`: AssessmentResponses
	fmt.Fprintf(os.Stdout, "Response from `AssessmentResponsesAPI.AssessmentResponsesAssessmentResponseIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assessmentResponseId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssessmentResponsesAssessmentResponseIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AssessmentResponses**](AssessmentResponses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssessmentResponsesAssessmentResponseIdPut

> AssessmentResponses AssessmentResponsesAssessmentResponseIdPut(ctx, assessmentResponseId).AssessmentResponsesPut(assessmentResponsesPut).Execute()



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
	assessmentResponseId := int64(789) // int64 | Model id
	assessmentResponsesPut := *openapiclient.NewAssessmentResponsesPut() // AssessmentResponsesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssessmentResponsesAPI.AssessmentResponsesAssessmentResponseIdPut(context.Background(), assessmentResponseId).AssessmentResponsesPut(assessmentResponsesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssessmentResponsesAPI.AssessmentResponsesAssessmentResponseIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssessmentResponsesAssessmentResponseIdPut`: AssessmentResponses
	fmt.Fprintf(os.Stdout, "Response from `AssessmentResponsesAPI.AssessmentResponsesAssessmentResponseIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assessmentResponseId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssessmentResponsesAssessmentResponseIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assessmentResponsesPut** | [**AssessmentResponsesPut**](AssessmentResponsesPut.md) |  | 

### Return type

[**AssessmentResponses**](AssessmentResponses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssessmentResponsesGet

> AssessmentResponsesGet200Response AssessmentResponsesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.AssessmentResponsesAPI.AssessmentResponsesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssessmentResponsesAPI.AssessmentResponsesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssessmentResponsesGet`: AssessmentResponsesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AssessmentResponsesAPI.AssessmentResponsesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssessmentResponsesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AssessmentResponsesGet200Response**](AssessmentResponsesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

