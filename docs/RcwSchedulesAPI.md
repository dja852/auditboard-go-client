# \RcwSchedulesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RcwSchedulesGet**](RcwSchedulesAPI.md#RcwSchedulesGet) | **Get** /rcw_schedules | 
[**RcwSchedulesPost**](RcwSchedulesAPI.md#RcwSchedulesPost) | **Post** /rcw_schedules | 
[**RcwSchedulesRcwScheduleIdDelete**](RcwSchedulesAPI.md#RcwSchedulesRcwScheduleIdDelete) | **Delete** /rcw_schedules/{rcw_schedule_id} | 
[**RcwSchedulesRcwScheduleIdGet**](RcwSchedulesAPI.md#RcwSchedulesRcwScheduleIdGet) | **Get** /rcw_schedules/{rcw_schedule_id} | 
[**RcwSchedulesRcwScheduleIdPut**](RcwSchedulesAPI.md#RcwSchedulesRcwScheduleIdPut) | **Put** /rcw_schedules/{rcw_schedule_id} | 



## RcwSchedulesGet

> RcwSchedulesGet200Response RcwSchedulesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.RcwSchedulesAPI.RcwSchedulesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwSchedulesAPI.RcwSchedulesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwSchedulesGet`: RcwSchedulesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RcwSchedulesAPI.RcwSchedulesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRcwSchedulesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**RcwSchedulesGet200Response**](RcwSchedulesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RcwSchedulesPost

> RcwSchedulesGet200Response RcwSchedulesPost(ctx).RcwSchedulesPostRequest(rcwSchedulesPostRequest).Execute()



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
	rcwSchedulesPostRequest := *openapiclient.NewRcwSchedulesPostRequest() // RcwSchedulesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RcwSchedulesAPI.RcwSchedulesPost(context.Background()).RcwSchedulesPostRequest(rcwSchedulesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwSchedulesAPI.RcwSchedulesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwSchedulesPost`: RcwSchedulesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RcwSchedulesAPI.RcwSchedulesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRcwSchedulesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rcwSchedulesPostRequest** | [**RcwSchedulesPostRequest**](RcwSchedulesPostRequest.md) |  | 

### Return type

[**RcwSchedulesGet200Response**](RcwSchedulesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RcwSchedulesRcwScheduleIdDelete

> RcwSchedules RcwSchedulesRcwScheduleIdDelete(ctx, rcwScheduleId).Execute()



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
	rcwScheduleId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RcwSchedulesAPI.RcwSchedulesRcwScheduleIdDelete(context.Background(), rcwScheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwSchedulesAPI.RcwSchedulesRcwScheduleIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwSchedulesRcwScheduleIdDelete`: RcwSchedules
	fmt.Fprintf(os.Stdout, "Response from `RcwSchedulesAPI.RcwSchedulesRcwScheduleIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rcwScheduleId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRcwSchedulesRcwScheduleIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RcwSchedules**](RcwSchedules.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RcwSchedulesRcwScheduleIdGet

> RcwSchedules RcwSchedulesRcwScheduleIdGet(ctx, rcwScheduleId).Include(include).Execute()



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
	rcwScheduleId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RcwSchedulesAPI.RcwSchedulesRcwScheduleIdGet(context.Background(), rcwScheduleId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwSchedulesAPI.RcwSchedulesRcwScheduleIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwSchedulesRcwScheduleIdGet`: RcwSchedules
	fmt.Fprintf(os.Stdout, "Response from `RcwSchedulesAPI.RcwSchedulesRcwScheduleIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rcwScheduleId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRcwSchedulesRcwScheduleIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**RcwSchedules**](RcwSchedules.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RcwSchedulesRcwScheduleIdPut

> RcwSchedules RcwSchedulesRcwScheduleIdPut(ctx, rcwScheduleId).RcwSchedulesPut(rcwSchedulesPut).Execute()



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
	rcwScheduleId := int64(789) // int64 | Model id
	rcwSchedulesPut := *openapiclient.NewRcwSchedulesPut() // RcwSchedulesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RcwSchedulesAPI.RcwSchedulesRcwScheduleIdPut(context.Background(), rcwScheduleId).RcwSchedulesPut(rcwSchedulesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RcwSchedulesAPI.RcwSchedulesRcwScheduleIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RcwSchedulesRcwScheduleIdPut`: RcwSchedules
	fmt.Fprintf(os.Stdout, "Response from `RcwSchedulesAPI.RcwSchedulesRcwScheduleIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rcwScheduleId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRcwSchedulesRcwScheduleIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rcwSchedulesPut** | [**RcwSchedulesPut**](RcwSchedulesPut.md) |  | 

### Return type

[**RcwSchedules**](RcwSchedules.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

