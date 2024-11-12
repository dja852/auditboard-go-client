# \SurveysAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SurveysGet**](SurveysAPI.md#SurveysGet) | **Get** /surveys | 
[**SurveysPost**](SurveysAPI.md#SurveysPost) | **Post** /surveys | 
[**SurveysSurveyIdDelete**](SurveysAPI.md#SurveysSurveyIdDelete) | **Delete** /surveys/{survey_id} | 
[**SurveysSurveyIdGet**](SurveysAPI.md#SurveysSurveyIdGet) | **Get** /surveys/{survey_id} | 
[**SurveysSurveyIdPut**](SurveysAPI.md#SurveysSurveyIdPut) | **Put** /surveys/{survey_id} | 



## SurveysGet

> SurveysGet200Response SurveysGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.SurveysAPI.SurveysGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveysAPI.SurveysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveysGet`: SurveysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SurveysAPI.SurveysGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSurveysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**SurveysGet200Response**](SurveysGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveysPost

> SurveysGet200Response SurveysPost(ctx).SurveysPostRequest(surveysPostRequest).Execute()



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
	surveysPostRequest := *openapiclient.NewSurveysPostRequest() // SurveysPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveysAPI.SurveysPost(context.Background()).SurveysPostRequest(surveysPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveysAPI.SurveysPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveysPost`: SurveysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SurveysAPI.SurveysPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSurveysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **surveysPostRequest** | [**SurveysPostRequest**](SurveysPostRequest.md) |  | 

### Return type

[**SurveysGet200Response**](SurveysGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveysSurveyIdDelete

> Surveys SurveysSurveyIdDelete(ctx, surveyId).Execute()



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
	surveyId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveysAPI.SurveysSurveyIdDelete(context.Background(), surveyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveysAPI.SurveysSurveyIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveysSurveyIdDelete`: Surveys
	fmt.Fprintf(os.Stdout, "Response from `SurveysAPI.SurveysSurveyIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**surveyId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSurveysSurveyIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Surveys**](Surveys.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveysSurveyIdGet

> Surveys SurveysSurveyIdGet(ctx, surveyId).Include(include).Execute()



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
	surveyId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveysAPI.SurveysSurveyIdGet(context.Background(), surveyId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveysAPI.SurveysSurveyIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveysSurveyIdGet`: Surveys
	fmt.Fprintf(os.Stdout, "Response from `SurveysAPI.SurveysSurveyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**surveyId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSurveysSurveyIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**Surveys**](Surveys.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveysSurveyIdPut

> Surveys SurveysSurveyIdPut(ctx, surveyId).SurveysPut(surveysPut).Execute()



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
	surveyId := int64(789) // int64 | Model id
	surveysPut := *openapiclient.NewSurveysPut() // SurveysPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveysAPI.SurveysSurveyIdPut(context.Background(), surveyId).SurveysPut(surveysPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveysAPI.SurveysSurveyIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveysSurveyIdPut`: Surveys
	fmt.Fprintf(os.Stdout, "Response from `SurveysAPI.SurveysSurveyIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**surveyId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSurveysSurveyIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **surveysPut** | [**SurveysPut**](SurveysPut.md) |  | 

### Return type

[**Surveys**](Surveys.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

