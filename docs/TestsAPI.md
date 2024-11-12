# \TestsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestsGet**](TestsAPI.md#TestsGet) | **Get** /tests | 
[**TestsPost**](TestsAPI.md#TestsPost) | **Post** /tests | 
[**TestsTestIdAmendPrimaryReviewPost**](TestsAPI.md#TestsTestIdAmendPrimaryReviewPost) | **Post** /tests/{test_id}/amend_primary_review | 
[**TestsTestIdAmendSecondaryReviewPost**](TestsAPI.md#TestsTestIdAmendSecondaryReviewPost) | **Post** /tests/{test_id}/amend_secondary_review | 
[**TestsTestIdApprovePrimaryReviewPost**](TestsAPI.md#TestsTestIdApprovePrimaryReviewPost) | **Post** /tests/{test_id}/approve_primary_review | 
[**TestsTestIdApproveSecondaryReviewPost**](TestsAPI.md#TestsTestIdApproveSecondaryReviewPost) | **Post** /tests/{test_id}/approve_secondary_review | 
[**TestsTestIdBackToTesterPost**](TestsAPI.md#TestsTestIdBackToTesterPost) | **Post** /tests/{test_id}/back_to_tester | 
[**TestsTestIdCompleteTestPost**](TestsAPI.md#TestsTestIdCompleteTestPost) | **Post** /tests/{test_id}/complete_test | 
[**TestsTestIdDelete**](TestsAPI.md#TestsTestIdDelete) | **Delete** /tests/{test_id} | 
[**TestsTestIdGet**](TestsAPI.md#TestsTestIdGet) | **Get** /tests/{test_id} | 
[**TestsTestIdMarkEffectivePost**](TestsAPI.md#TestsTestIdMarkEffectivePost) | **Post** /tests/{test_id}/mark_effective | 
[**TestsTestIdMarkIneffectivePost**](TestsAPI.md#TestsTestIdMarkIneffectivePost) | **Post** /tests/{test_id}/mark_ineffective | 
[**TestsTestIdMarkNoPopulationPost**](TestsAPI.md#TestsTestIdMarkNoPopulationPost) | **Post** /tests/{test_id}/mark_no_population | 
[**TestsTestIdPut**](TestsAPI.md#TestsTestIdPut) | **Put** /tests/{test_id} | 
[**TestsTestIdReopenPost**](TestsAPI.md#TestsTestIdReopenPost) | **Post** /tests/{test_id}/reopen | 
[**TestsTestIdSendToSecondaryReviewerPost**](TestsAPI.md#TestsTestIdSendToSecondaryReviewerPost) | **Post** /tests/{test_id}/send_to_secondary_reviewer | 
[**TestsTestIdStartReviewPost**](TestsAPI.md#TestsTestIdStartReviewPost) | **Post** /tests/{test_id}/start_review | 
[**TestsTestIdStartTestPost**](TestsAPI.md#TestsTestIdStartTestPost) | **Post** /tests/{test_id}/start_test | 
[**TestsTestIdStopTestPost**](TestsAPI.md#TestsTestIdStopTestPost) | **Post** /tests/{test_id}/stop_test | 



## TestsGet

> TestsGet200Response TestsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.TestsAPI.TestsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsGet`: TestsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**TestsGet200Response**](TestsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsPost

> TestsGet200Response TestsPost(ctx).TestsPostRequest(testsPostRequest).Execute()



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
	testsPostRequest := *openapiclient.NewTestsPostRequest() // TestsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsPost(context.Background()).TestsPostRequest(testsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsPost`: TestsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testsPostRequest** | [**TestsPostRequest**](TestsPostRequest.md) |  | 

### Return type

[**TestsGet200Response**](TestsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsTestIdAmendPrimaryReviewPost

> TestStop TestsTestIdAmendPrimaryReviewPost(ctx, testId).Body(body).Execute()



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
	testId := int64(789) // int64 | Test id
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsTestIdAmendPrimaryReviewPost(context.Background(), testId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsTestIdAmendPrimaryReviewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsTestIdAmendPrimaryReviewPost`: TestStop
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsTestIdAmendPrimaryReviewPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **int64** | Test id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsTestIdAmendPrimaryReviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**TestStop**](TestStop.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsTestIdAmendSecondaryReviewPost

> TestStop TestsTestIdAmendSecondaryReviewPost(ctx, testId).Body(body).Execute()



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
	testId := int64(789) // int64 | Test id
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsTestIdAmendSecondaryReviewPost(context.Background(), testId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsTestIdAmendSecondaryReviewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsTestIdAmendSecondaryReviewPost`: TestStop
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsTestIdAmendSecondaryReviewPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **int64** | Test id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsTestIdAmendSecondaryReviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**TestStop**](TestStop.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsTestIdApprovePrimaryReviewPost

> TestStop TestsTestIdApprovePrimaryReviewPost(ctx, testId).Body(body).Execute()



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
	testId := int64(789) // int64 | Test id
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsTestIdApprovePrimaryReviewPost(context.Background(), testId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsTestIdApprovePrimaryReviewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsTestIdApprovePrimaryReviewPost`: TestStop
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsTestIdApprovePrimaryReviewPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **int64** | Test id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsTestIdApprovePrimaryReviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**TestStop**](TestStop.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsTestIdApproveSecondaryReviewPost

> TestStop TestsTestIdApproveSecondaryReviewPost(ctx, testId).Body(body).Execute()



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
	testId := int64(789) // int64 | Test id
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsTestIdApproveSecondaryReviewPost(context.Background(), testId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsTestIdApproveSecondaryReviewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsTestIdApproveSecondaryReviewPost`: TestStop
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsTestIdApproveSecondaryReviewPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **int64** | Test id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsTestIdApproveSecondaryReviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**TestStop**](TestStop.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsTestIdBackToTesterPost

> TestStop TestsTestIdBackToTesterPost(ctx, testId).Body(body).Execute()



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
	testId := int64(789) // int64 | Test id
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsTestIdBackToTesterPost(context.Background(), testId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsTestIdBackToTesterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsTestIdBackToTesterPost`: TestStop
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsTestIdBackToTesterPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **int64** | Test id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsTestIdBackToTesterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**TestStop**](TestStop.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsTestIdCompleteTestPost

> TestStop TestsTestIdCompleteTestPost(ctx, testId).Body(body).Execute()



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
	testId := int64(789) // int64 | Test id
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsTestIdCompleteTestPost(context.Background(), testId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsTestIdCompleteTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsTestIdCompleteTestPost`: TestStop
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsTestIdCompleteTestPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **int64** | Test id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsTestIdCompleteTestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**TestStop**](TestStop.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsTestIdDelete

> Tests TestsTestIdDelete(ctx, testId).Execute()



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
	testId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsTestIdDelete(context.Background(), testId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsTestIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsTestIdDelete`: Tests
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsTestIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsTestIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tests**](Tests.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsTestIdGet

> Tests TestsTestIdGet(ctx, testId).Include(include).Execute()



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
	testId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsTestIdGet(context.Background(), testId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsTestIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsTestIdGet`: Tests
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsTestIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsTestIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**Tests**](Tests.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsTestIdMarkEffectivePost

> TestStop TestsTestIdMarkEffectivePost(ctx, testId).Body(body).Execute()



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
	testId := int64(789) // int64 | Test id
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsTestIdMarkEffectivePost(context.Background(), testId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsTestIdMarkEffectivePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsTestIdMarkEffectivePost`: TestStop
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsTestIdMarkEffectivePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **int64** | Test id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsTestIdMarkEffectivePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**TestStop**](TestStop.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsTestIdMarkIneffectivePost

> TestStop TestsTestIdMarkIneffectivePost(ctx, testId).Body(body).Execute()



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
	testId := int64(789) // int64 | Test id
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsTestIdMarkIneffectivePost(context.Background(), testId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsTestIdMarkIneffectivePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsTestIdMarkIneffectivePost`: TestStop
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsTestIdMarkIneffectivePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **int64** | Test id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsTestIdMarkIneffectivePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**TestStop**](TestStop.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsTestIdMarkNoPopulationPost

> TestStop TestsTestIdMarkNoPopulationPost(ctx, testId).Body(body).Execute()



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
	testId := int64(789) // int64 | Test id
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsTestIdMarkNoPopulationPost(context.Background(), testId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsTestIdMarkNoPopulationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsTestIdMarkNoPopulationPost`: TestStop
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsTestIdMarkNoPopulationPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **int64** | Test id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsTestIdMarkNoPopulationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**TestStop**](TestStop.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsTestIdPut

> Tests TestsTestIdPut(ctx, testId).TestsPut(testsPut).Execute()



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
	testId := int64(789) // int64 | Model id
	testsPut := *openapiclient.NewTestsPut() // TestsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsTestIdPut(context.Background(), testId).TestsPut(testsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsTestIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsTestIdPut`: Tests
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsTestIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsTestIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testsPut** | [**TestsPut**](TestsPut.md) |  | 

### Return type

[**Tests**](Tests.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsTestIdReopenPost

> TestStop TestsTestIdReopenPost(ctx, testId).Body(body).Execute()



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
	testId := int64(789) // int64 | Test id
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsTestIdReopenPost(context.Background(), testId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsTestIdReopenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsTestIdReopenPost`: TestStop
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsTestIdReopenPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **int64** | Test id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsTestIdReopenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**TestStop**](TestStop.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsTestIdSendToSecondaryReviewerPost

> TestStop TestsTestIdSendToSecondaryReviewerPost(ctx, testId).Body(body).Execute()



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
	testId := int64(789) // int64 | Test id
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsTestIdSendToSecondaryReviewerPost(context.Background(), testId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsTestIdSendToSecondaryReviewerPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsTestIdSendToSecondaryReviewerPost`: TestStop
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsTestIdSendToSecondaryReviewerPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **int64** | Test id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsTestIdSendToSecondaryReviewerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**TestStop**](TestStop.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsTestIdStartReviewPost

> TestStop TestsTestIdStartReviewPost(ctx, testId).Body(body).Execute()



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
	testId := int64(789) // int64 | Test id
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsTestIdStartReviewPost(context.Background(), testId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsTestIdStartReviewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsTestIdStartReviewPost`: TestStop
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsTestIdStartReviewPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **int64** | Test id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsTestIdStartReviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**TestStop**](TestStop.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsTestIdStartTestPost

> TestStop TestsTestIdStartTestPost(ctx, testId).Body(body).Execute()



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
	testId := int64(789) // int64 | Test id
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsTestIdStartTestPost(context.Background(), testId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsTestIdStartTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsTestIdStartTestPost`: TestStop
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsTestIdStartTestPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **int64** | Test id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsTestIdStartTestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**TestStop**](TestStop.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsTestIdStopTestPost

> TestStop TestsTestIdStopTestPost(ctx, testId).Body(body).Execute()



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
	testId := int64(789) // int64 | Test id
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsTestIdStopTestPost(context.Background(), testId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsTestIdStopTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsTestIdStopTestPost`: TestStop
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsTestIdStopTestPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **int64** | Test id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsTestIdStopTestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**TestStop**](TestStop.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

