# \AssessmentTemplatesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssessmentTemplatesAssessmentTemplateIdDelete**](AssessmentTemplatesAPI.md#AssessmentTemplatesAssessmentTemplateIdDelete) | **Delete** /assessment_templates/{assessment_template_id} | 
[**AssessmentTemplatesAssessmentTemplateIdGet**](AssessmentTemplatesAPI.md#AssessmentTemplatesAssessmentTemplateIdGet) | **Get** /assessment_templates/{assessment_template_id} | 
[**AssessmentTemplatesAssessmentTemplateIdPut**](AssessmentTemplatesAPI.md#AssessmentTemplatesAssessmentTemplateIdPut) | **Put** /assessment_templates/{assessment_template_id} | 
[**AssessmentTemplatesGet**](AssessmentTemplatesAPI.md#AssessmentTemplatesGet) | **Get** /assessment_templates | 
[**AssessmentTemplatesPost**](AssessmentTemplatesAPI.md#AssessmentTemplatesPost) | **Post** /assessment_templates | 



## AssessmentTemplatesAssessmentTemplateIdDelete

> AssessmentTemplates AssessmentTemplatesAssessmentTemplateIdDelete(ctx, assessmentTemplateId).Execute()



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
	assessmentTemplateId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssessmentTemplatesAPI.AssessmentTemplatesAssessmentTemplateIdDelete(context.Background(), assessmentTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssessmentTemplatesAPI.AssessmentTemplatesAssessmentTemplateIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssessmentTemplatesAssessmentTemplateIdDelete`: AssessmentTemplates
	fmt.Fprintf(os.Stdout, "Response from `AssessmentTemplatesAPI.AssessmentTemplatesAssessmentTemplateIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assessmentTemplateId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssessmentTemplatesAssessmentTemplateIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssessmentTemplates**](AssessmentTemplates.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssessmentTemplatesAssessmentTemplateIdGet

> AssessmentTemplates AssessmentTemplatesAssessmentTemplateIdGet(ctx, assessmentTemplateId).Include(include).Execute()



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
	assessmentTemplateId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssessmentTemplatesAPI.AssessmentTemplatesAssessmentTemplateIdGet(context.Background(), assessmentTemplateId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssessmentTemplatesAPI.AssessmentTemplatesAssessmentTemplateIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssessmentTemplatesAssessmentTemplateIdGet`: AssessmentTemplates
	fmt.Fprintf(os.Stdout, "Response from `AssessmentTemplatesAPI.AssessmentTemplatesAssessmentTemplateIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assessmentTemplateId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssessmentTemplatesAssessmentTemplateIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AssessmentTemplates**](AssessmentTemplates.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssessmentTemplatesAssessmentTemplateIdPut

> AssessmentTemplates AssessmentTemplatesAssessmentTemplateIdPut(ctx, assessmentTemplateId).AssessmentTemplatesPut(assessmentTemplatesPut).Execute()



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
	assessmentTemplateId := int64(789) // int64 | Model id
	assessmentTemplatesPut := *openapiclient.NewAssessmentTemplatesPut() // AssessmentTemplatesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssessmentTemplatesAPI.AssessmentTemplatesAssessmentTemplateIdPut(context.Background(), assessmentTemplateId).AssessmentTemplatesPut(assessmentTemplatesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssessmentTemplatesAPI.AssessmentTemplatesAssessmentTemplateIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssessmentTemplatesAssessmentTemplateIdPut`: AssessmentTemplates
	fmt.Fprintf(os.Stdout, "Response from `AssessmentTemplatesAPI.AssessmentTemplatesAssessmentTemplateIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assessmentTemplateId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssessmentTemplatesAssessmentTemplateIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assessmentTemplatesPut** | [**AssessmentTemplatesPut**](AssessmentTemplatesPut.md) |  | 

### Return type

[**AssessmentTemplates**](AssessmentTemplates.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssessmentTemplatesGet

> AssessmentTemplatesGet200Response AssessmentTemplatesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.AssessmentTemplatesAPI.AssessmentTemplatesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssessmentTemplatesAPI.AssessmentTemplatesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssessmentTemplatesGet`: AssessmentTemplatesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AssessmentTemplatesAPI.AssessmentTemplatesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssessmentTemplatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AssessmentTemplatesGet200Response**](AssessmentTemplatesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssessmentTemplatesPost

> AssessmentTemplatesGet200Response AssessmentTemplatesPost(ctx).AssessmentTemplatesPostRequest(assessmentTemplatesPostRequest).Execute()



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
	assessmentTemplatesPostRequest := *openapiclient.NewAssessmentTemplatesPostRequest() // AssessmentTemplatesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssessmentTemplatesAPI.AssessmentTemplatesPost(context.Background()).AssessmentTemplatesPostRequest(assessmentTemplatesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssessmentTemplatesAPI.AssessmentTemplatesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssessmentTemplatesPost`: AssessmentTemplatesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AssessmentTemplatesAPI.AssessmentTemplatesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssessmentTemplatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assessmentTemplatesPostRequest** | [**AssessmentTemplatesPostRequest**](AssessmentTemplatesPostRequest.md) |  | 

### Return type

[**AssessmentTemplatesGet200Response**](AssessmentTemplatesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

