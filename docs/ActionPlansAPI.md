# \ActionPlansAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionPlansActionPlanIdDelete**](ActionPlansAPI.md#ActionPlansActionPlanIdDelete) | **Delete** /action_plans/{action_plan_id} | 
[**ActionPlansActionPlanIdGet**](ActionPlansAPI.md#ActionPlansActionPlanIdGet) | **Get** /action_plans/{action_plan_id} | 
[**ActionPlansActionPlanIdPut**](ActionPlansAPI.md#ActionPlansActionPlanIdPut) | **Put** /action_plans/{action_plan_id} | 
[**ActionPlansGet**](ActionPlansAPI.md#ActionPlansGet) | **Get** /action_plans | 
[**ActionPlansPost**](ActionPlansAPI.md#ActionPlansPost) | **Post** /action_plans | 



## ActionPlansActionPlanIdDelete

> ActionPlans ActionPlansActionPlanIdDelete(ctx, actionPlanId).Execute()



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
	actionPlanId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionPlansAPI.ActionPlansActionPlanIdDelete(context.Background(), actionPlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionPlansAPI.ActionPlansActionPlanIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionPlansActionPlanIdDelete`: ActionPlans
	fmt.Fprintf(os.Stdout, "Response from `ActionPlansAPI.ActionPlansActionPlanIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionPlanId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionPlansActionPlanIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionPlans**](ActionPlans.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionPlansActionPlanIdGet

> ActionPlans ActionPlansActionPlanIdGet(ctx, actionPlanId).Include(include).Execute()



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
	actionPlanId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionPlansAPI.ActionPlansActionPlanIdGet(context.Background(), actionPlanId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionPlansAPI.ActionPlansActionPlanIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionPlansActionPlanIdGet`: ActionPlans
	fmt.Fprintf(os.Stdout, "Response from `ActionPlansAPI.ActionPlansActionPlanIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionPlanId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionPlansActionPlanIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**ActionPlans**](ActionPlans.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionPlansActionPlanIdPut

> ActionPlans ActionPlansActionPlanIdPut(ctx, actionPlanId).ActionPlansPut(actionPlansPut).Execute()



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
	actionPlanId := int64(789) // int64 | Model id
	actionPlansPut := *openapiclient.NewActionPlansPut() // ActionPlansPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionPlansAPI.ActionPlansActionPlanIdPut(context.Background(), actionPlanId).ActionPlansPut(actionPlansPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionPlansAPI.ActionPlansActionPlanIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionPlansActionPlanIdPut`: ActionPlans
	fmt.Fprintf(os.Stdout, "Response from `ActionPlansAPI.ActionPlansActionPlanIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionPlanId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionPlansActionPlanIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actionPlansPut** | [**ActionPlansPut**](ActionPlansPut.md) |  | 

### Return type

[**ActionPlans**](ActionPlans.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionPlansGet

> ActionPlansGet200Response ActionPlansGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.ActionPlansAPI.ActionPlansGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionPlansAPI.ActionPlansGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionPlansGet`: ActionPlansGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ActionPlansAPI.ActionPlansGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActionPlansGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**ActionPlansGet200Response**](ActionPlansGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionPlansPost

> ActionPlansGet200Response ActionPlansPost(ctx).ActionPlansPostRequest(actionPlansPostRequest).Execute()



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
	actionPlansPostRequest := *openapiclient.NewActionPlansPostRequest() // ActionPlansPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionPlansAPI.ActionPlansPost(context.Background()).ActionPlansPostRequest(actionPlansPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionPlansAPI.ActionPlansPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionPlansPost`: ActionPlansGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ActionPlansAPI.ActionPlansPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActionPlansPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actionPlansPostRequest** | [**ActionPlansPostRequest**](ActionPlansPostRequest.md) |  | 

### Return type

[**ActionPlansGet200Response**](ActionPlansGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

