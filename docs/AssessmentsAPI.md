# \AssessmentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssessmentsAssessmentIdAddRisksToAssessmentPost**](AssessmentsAPI.md#AssessmentsAssessmentIdAddRisksToAssessmentPost) | **Post** /assessments/{assessment_id}/add_risks_to_assessment | 
[**AssessmentsAssessmentIdGet**](AssessmentsAPI.md#AssessmentsAssessmentIdGet) | **Get** /assessments/{assessment_id} | 
[**AssessmentsAssessmentIdPut**](AssessmentsAPI.md#AssessmentsAssessmentIdPut) | **Put** /assessments/{assessment_id} | 
[**AssessmentsGet**](AssessmentsAPI.md#AssessmentsGet) | **Get** /assessments | 
[**AssessmentsPost**](AssessmentsAPI.md#AssessmentsPost) | **Post** /assessments | 



## AssessmentsAssessmentIdAddRisksToAssessmentPost

> AssessmentsAssessmentIdAddRisksToAssessmentPost200Response AssessmentsAssessmentIdAddRisksToAssessmentPost(ctx, assessmentId).AssessmentsAssessmentIdAddRisksToAssessmentPostRequest(assessmentsAssessmentIdAddRisksToAssessmentPostRequest).Execute()



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
	assessmentId := int64(789) // int64 | The id of the assessment where the risk(s) will be added.
	assessmentsAssessmentIdAddRisksToAssessmentPostRequest := *openapiclient.NewAssessmentsAssessmentIdAddRisksToAssessmentPostRequest() // AssessmentsAssessmentIdAddRisksToAssessmentPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssessmentsAPI.AssessmentsAssessmentIdAddRisksToAssessmentPost(context.Background(), assessmentId).AssessmentsAssessmentIdAddRisksToAssessmentPostRequest(assessmentsAssessmentIdAddRisksToAssessmentPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssessmentsAPI.AssessmentsAssessmentIdAddRisksToAssessmentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssessmentsAssessmentIdAddRisksToAssessmentPost`: AssessmentsAssessmentIdAddRisksToAssessmentPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AssessmentsAPI.AssessmentsAssessmentIdAddRisksToAssessmentPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assessmentId** | **int64** | The id of the assessment where the risk(s) will be added. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssessmentsAssessmentIdAddRisksToAssessmentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assessmentsAssessmentIdAddRisksToAssessmentPostRequest** | [**AssessmentsAssessmentIdAddRisksToAssessmentPostRequest**](AssessmentsAssessmentIdAddRisksToAssessmentPostRequest.md) |  | 

### Return type

[**AssessmentsAssessmentIdAddRisksToAssessmentPost200Response**](AssessmentsAssessmentIdAddRisksToAssessmentPost200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssessmentsAssessmentIdGet

> Assessments AssessmentsAssessmentIdGet(ctx, assessmentId).Include(include).Execute()



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
	assessmentId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssessmentsAPI.AssessmentsAssessmentIdGet(context.Background(), assessmentId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssessmentsAPI.AssessmentsAssessmentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssessmentsAssessmentIdGet`: Assessments
	fmt.Fprintf(os.Stdout, "Response from `AssessmentsAPI.AssessmentsAssessmentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assessmentId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssessmentsAssessmentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**Assessments**](Assessments.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssessmentsAssessmentIdPut

> Assessments AssessmentsAssessmentIdPut(ctx, assessmentId).Body(body).Execute()



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
	assessmentId := int64(789) // int64 | Model id
	body := interface{}(987) // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssessmentsAPI.AssessmentsAssessmentIdPut(context.Background(), assessmentId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssessmentsAPI.AssessmentsAssessmentIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssessmentsAssessmentIdPut`: Assessments
	fmt.Fprintf(os.Stdout, "Response from `AssessmentsAPI.AssessmentsAssessmentIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assessmentId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssessmentsAssessmentIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** |  | 

### Return type

[**Assessments**](Assessments.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssessmentsGet

> AssessmentsGet200Response AssessmentsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.AssessmentsAPI.AssessmentsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssessmentsAPI.AssessmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssessmentsGet`: AssessmentsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AssessmentsAPI.AssessmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssessmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AssessmentsGet200Response**](AssessmentsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssessmentsPost

> AssessmentsGet200Response AssessmentsPost(ctx).AssessmentsPostRequest(assessmentsPostRequest).Execute()



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
	assessmentsPostRequest := *openapiclient.NewAssessmentsPostRequest() // AssessmentsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssessmentsAPI.AssessmentsPost(context.Background()).AssessmentsPostRequest(assessmentsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssessmentsAPI.AssessmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssessmentsPost`: AssessmentsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AssessmentsAPI.AssessmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssessmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assessmentsPostRequest** | [**AssessmentsPostRequest**](AssessmentsPostRequest.md) |  | 

### Return type

[**AssessmentsGet200Response**](AssessmentsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

