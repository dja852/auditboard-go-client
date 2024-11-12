# \AssessmentPeriodsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssessmentPeriodsAssessmentPeriodIdDelete**](AssessmentPeriodsAPI.md#AssessmentPeriodsAssessmentPeriodIdDelete) | **Delete** /assessment_periods/{assessment_period_id} | 
[**AssessmentPeriodsAssessmentPeriodIdGet**](AssessmentPeriodsAPI.md#AssessmentPeriodsAssessmentPeriodIdGet) | **Get** /assessment_periods/{assessment_period_id} | 
[**AssessmentPeriodsAssessmentPeriodIdPut**](AssessmentPeriodsAPI.md#AssessmentPeriodsAssessmentPeriodIdPut) | **Put** /assessment_periods/{assessment_period_id} | 
[**AssessmentPeriodsGet**](AssessmentPeriodsAPI.md#AssessmentPeriodsGet) | **Get** /assessment_periods | 
[**AssessmentPeriodsPost**](AssessmentPeriodsAPI.md#AssessmentPeriodsPost) | **Post** /assessment_periods | 



## AssessmentPeriodsAssessmentPeriodIdDelete

> AssessmentPeriods AssessmentPeriodsAssessmentPeriodIdDelete(ctx, assessmentPeriodId).Execute()



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
	assessmentPeriodId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssessmentPeriodsAPI.AssessmentPeriodsAssessmentPeriodIdDelete(context.Background(), assessmentPeriodId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssessmentPeriodsAPI.AssessmentPeriodsAssessmentPeriodIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssessmentPeriodsAssessmentPeriodIdDelete`: AssessmentPeriods
	fmt.Fprintf(os.Stdout, "Response from `AssessmentPeriodsAPI.AssessmentPeriodsAssessmentPeriodIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assessmentPeriodId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssessmentPeriodsAssessmentPeriodIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssessmentPeriods**](AssessmentPeriods.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssessmentPeriodsAssessmentPeriodIdGet

> AssessmentPeriods AssessmentPeriodsAssessmentPeriodIdGet(ctx, assessmentPeriodId).Include(include).Execute()



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
	assessmentPeriodId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssessmentPeriodsAPI.AssessmentPeriodsAssessmentPeriodIdGet(context.Background(), assessmentPeriodId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssessmentPeriodsAPI.AssessmentPeriodsAssessmentPeriodIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssessmentPeriodsAssessmentPeriodIdGet`: AssessmentPeriods
	fmt.Fprintf(os.Stdout, "Response from `AssessmentPeriodsAPI.AssessmentPeriodsAssessmentPeriodIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assessmentPeriodId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssessmentPeriodsAssessmentPeriodIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AssessmentPeriods**](AssessmentPeriods.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssessmentPeriodsAssessmentPeriodIdPut

> AssessmentPeriods AssessmentPeriodsAssessmentPeriodIdPut(ctx, assessmentPeriodId).AssessmentPeriodsPut(assessmentPeriodsPut).Execute()



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
	assessmentPeriodId := int64(789) // int64 | Model id
	assessmentPeriodsPut := *openapiclient.NewAssessmentPeriodsPut() // AssessmentPeriodsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssessmentPeriodsAPI.AssessmentPeriodsAssessmentPeriodIdPut(context.Background(), assessmentPeriodId).AssessmentPeriodsPut(assessmentPeriodsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssessmentPeriodsAPI.AssessmentPeriodsAssessmentPeriodIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssessmentPeriodsAssessmentPeriodIdPut`: AssessmentPeriods
	fmt.Fprintf(os.Stdout, "Response from `AssessmentPeriodsAPI.AssessmentPeriodsAssessmentPeriodIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assessmentPeriodId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssessmentPeriodsAssessmentPeriodIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assessmentPeriodsPut** | [**AssessmentPeriodsPut**](AssessmentPeriodsPut.md) |  | 

### Return type

[**AssessmentPeriods**](AssessmentPeriods.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssessmentPeriodsGet

> AssessmentPeriodsGet200Response AssessmentPeriodsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.AssessmentPeriodsAPI.AssessmentPeriodsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssessmentPeriodsAPI.AssessmentPeriodsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssessmentPeriodsGet`: AssessmentPeriodsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AssessmentPeriodsAPI.AssessmentPeriodsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssessmentPeriodsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AssessmentPeriodsGet200Response**](AssessmentPeriodsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssessmentPeriodsPost

> AssessmentPeriodsGet200Response AssessmentPeriodsPost(ctx).AssessmentPeriodsPostRequest(assessmentPeriodsPostRequest).Execute()



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
	assessmentPeriodsPostRequest := *openapiclient.NewAssessmentPeriodsPostRequest() // AssessmentPeriodsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssessmentPeriodsAPI.AssessmentPeriodsPost(context.Background()).AssessmentPeriodsPostRequest(assessmentPeriodsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssessmentPeriodsAPI.AssessmentPeriodsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssessmentPeriodsPost`: AssessmentPeriodsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AssessmentPeriodsAPI.AssessmentPeriodsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssessmentPeriodsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assessmentPeriodsPostRequest** | [**AssessmentPeriodsPostRequest**](AssessmentPeriodsPostRequest.md) |  | 

### Return type

[**AssessmentPeriodsGet200Response**](AssessmentPeriodsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

