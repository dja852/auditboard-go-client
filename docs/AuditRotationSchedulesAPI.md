# \AuditRotationSchedulesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditRotationSchedulesAuditRotationScheduleIdDelete**](AuditRotationSchedulesAPI.md#AuditRotationSchedulesAuditRotationScheduleIdDelete) | **Delete** /audit_rotation_schedules/{audit_rotation_schedule_id} | 
[**AuditRotationSchedulesAuditRotationScheduleIdGet**](AuditRotationSchedulesAPI.md#AuditRotationSchedulesAuditRotationScheduleIdGet) | **Get** /audit_rotation_schedules/{audit_rotation_schedule_id} | 
[**AuditRotationSchedulesAuditRotationScheduleIdPut**](AuditRotationSchedulesAPI.md#AuditRotationSchedulesAuditRotationScheduleIdPut) | **Put** /audit_rotation_schedules/{audit_rotation_schedule_id} | 
[**AuditRotationSchedulesGet**](AuditRotationSchedulesAPI.md#AuditRotationSchedulesGet) | **Get** /audit_rotation_schedules | 
[**AuditRotationSchedulesPost**](AuditRotationSchedulesAPI.md#AuditRotationSchedulesPost) | **Post** /audit_rotation_schedules | 



## AuditRotationSchedulesAuditRotationScheduleIdDelete

> AuditRotationSchedules AuditRotationSchedulesAuditRotationScheduleIdDelete(ctx, auditRotationScheduleId).Execute()



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
	auditRotationScheduleId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditRotationSchedulesAPI.AuditRotationSchedulesAuditRotationScheduleIdDelete(context.Background(), auditRotationScheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditRotationSchedulesAPI.AuditRotationSchedulesAuditRotationScheduleIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditRotationSchedulesAuditRotationScheduleIdDelete`: AuditRotationSchedules
	fmt.Fprintf(os.Stdout, "Response from `AuditRotationSchedulesAPI.AuditRotationSchedulesAuditRotationScheduleIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditRotationScheduleId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditRotationSchedulesAuditRotationScheduleIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuditRotationSchedules**](AuditRotationSchedules.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditRotationSchedulesAuditRotationScheduleIdGet

> AuditRotationSchedules AuditRotationSchedulesAuditRotationScheduleIdGet(ctx, auditRotationScheduleId).Include(include).Execute()



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
	auditRotationScheduleId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditRotationSchedulesAPI.AuditRotationSchedulesAuditRotationScheduleIdGet(context.Background(), auditRotationScheduleId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditRotationSchedulesAPI.AuditRotationSchedulesAuditRotationScheduleIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditRotationSchedulesAuditRotationScheduleIdGet`: AuditRotationSchedules
	fmt.Fprintf(os.Stdout, "Response from `AuditRotationSchedulesAPI.AuditRotationSchedulesAuditRotationScheduleIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditRotationScheduleId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditRotationSchedulesAuditRotationScheduleIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditRotationSchedules**](AuditRotationSchedules.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditRotationSchedulesAuditRotationScheduleIdPut

> AuditRotationSchedules AuditRotationSchedulesAuditRotationScheduleIdPut(ctx, auditRotationScheduleId).AuditRotationSchedulesPut(auditRotationSchedulesPut).Execute()



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
	auditRotationScheduleId := int64(789) // int64 | Model id
	auditRotationSchedulesPut := *openapiclient.NewAuditRotationSchedulesPut() // AuditRotationSchedulesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditRotationSchedulesAPI.AuditRotationSchedulesAuditRotationScheduleIdPut(context.Background(), auditRotationScheduleId).AuditRotationSchedulesPut(auditRotationSchedulesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditRotationSchedulesAPI.AuditRotationSchedulesAuditRotationScheduleIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditRotationSchedulesAuditRotationScheduleIdPut`: AuditRotationSchedules
	fmt.Fprintf(os.Stdout, "Response from `AuditRotationSchedulesAPI.AuditRotationSchedulesAuditRotationScheduleIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditRotationScheduleId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditRotationSchedulesAuditRotationScheduleIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auditRotationSchedulesPut** | [**AuditRotationSchedulesPut**](AuditRotationSchedulesPut.md) |  | 

### Return type

[**AuditRotationSchedules**](AuditRotationSchedules.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditRotationSchedulesGet

> AuditRotationSchedulesGet200Response AuditRotationSchedulesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.AuditRotationSchedulesAPI.AuditRotationSchedulesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditRotationSchedulesAPI.AuditRotationSchedulesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditRotationSchedulesGet`: AuditRotationSchedulesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditRotationSchedulesAPI.AuditRotationSchedulesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditRotationSchedulesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AuditRotationSchedulesGet200Response**](AuditRotationSchedulesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditRotationSchedulesPost

> AuditRotationSchedulesGet200Response AuditRotationSchedulesPost(ctx).AuditRotationSchedulesPostRequest(auditRotationSchedulesPostRequest).Execute()



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
	auditRotationSchedulesPostRequest := *openapiclient.NewAuditRotationSchedulesPostRequest() // AuditRotationSchedulesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditRotationSchedulesAPI.AuditRotationSchedulesPost(context.Background()).AuditRotationSchedulesPostRequest(auditRotationSchedulesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditRotationSchedulesAPI.AuditRotationSchedulesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditRotationSchedulesPost`: AuditRotationSchedulesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditRotationSchedulesAPI.AuditRotationSchedulesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditRotationSchedulesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auditRotationSchedulesPostRequest** | [**AuditRotationSchedulesPostRequest**](AuditRotationSchedulesPostRequest.md) |  | 

### Return type

[**AuditRotationSchedulesGet200Response**](AuditRotationSchedulesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

