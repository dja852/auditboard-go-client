# \TimesheetEntriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TimesheetEntriesGet**](TimesheetEntriesAPI.md#TimesheetEntriesGet) | **Get** /timesheet_entries | 
[**TimesheetEntriesPost**](TimesheetEntriesAPI.md#TimesheetEntriesPost) | **Post** /timesheet_entries | 
[**TimesheetEntriesTimesheetEntryIdDelete**](TimesheetEntriesAPI.md#TimesheetEntriesTimesheetEntryIdDelete) | **Delete** /timesheet_entries/{timesheet_entry_id} | 
[**TimesheetEntriesTimesheetEntryIdGet**](TimesheetEntriesAPI.md#TimesheetEntriesTimesheetEntryIdGet) | **Get** /timesheet_entries/{timesheet_entry_id} | 
[**TimesheetEntriesTimesheetEntryIdPut**](TimesheetEntriesAPI.md#TimesheetEntriesTimesheetEntryIdPut) | **Put** /timesheet_entries/{timesheet_entry_id} | 



## TimesheetEntriesGet

> TimesheetEntriesGet200Response TimesheetEntriesGet(ctx).Start(start).End(end).SkewForWeek(skewForWeek).TimesheetProjectId(timesheetProjectId).TimesheetableType(timesheetableType).TimesheetableId(timesheetableId).UserId(userId).Limit(limit).Execute()



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
	start := "start_example" // string | The start date for timesheets to find (optional)
	end := "end_example" // string | The end date for timesheets to find (optional)
	skewForWeek := true // bool | \"Skews\" start and end date to the start and end of the week respectively (optional)
	timesheetProjectId := int32(56) // int32 |  (optional)
	timesheetableType := "timesheetableType_example" // string |  (optional)
	timesheetableId := int32(56) // int32 |  (optional)
	userId := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Limits the returned result set (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimesheetEntriesAPI.TimesheetEntriesGet(context.Background()).Start(start).End(end).SkewForWeek(skewForWeek).TimesheetProjectId(timesheetProjectId).TimesheetableType(timesheetableType).TimesheetableId(timesheetableId).UserId(userId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimesheetEntriesAPI.TimesheetEntriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TimesheetEntriesGet`: TimesheetEntriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TimesheetEntriesAPI.TimesheetEntriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTimesheetEntriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | The start date for timesheets to find | 
 **end** | **string** | The end date for timesheets to find | 
 **skewForWeek** | **bool** | \&quot;Skews\&quot; start and end date to the start and end of the week respectively | 
 **timesheetProjectId** | **int32** |  | 
 **timesheetableType** | **string** |  | 
 **timesheetableId** | **int32** |  | 
 **userId** | **int32** |  | 
 **limit** | **int32** | Limits the returned result set | 

### Return type

[**TimesheetEntriesGet200Response**](TimesheetEntriesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TimesheetEntriesPost

> TimesheetEntriesGet200Response TimesheetEntriesPost(ctx).TimesheetEntriesPostRequest(timesheetEntriesPostRequest).Execute()



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
	timesheetEntriesPostRequest := *openapiclient.NewTimesheetEntriesPostRequest() // TimesheetEntriesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimesheetEntriesAPI.TimesheetEntriesPost(context.Background()).TimesheetEntriesPostRequest(timesheetEntriesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimesheetEntriesAPI.TimesheetEntriesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TimesheetEntriesPost`: TimesheetEntriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TimesheetEntriesAPI.TimesheetEntriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTimesheetEntriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timesheetEntriesPostRequest** | [**TimesheetEntriesPostRequest**](TimesheetEntriesPostRequest.md) |  | 

### Return type

[**TimesheetEntriesGet200Response**](TimesheetEntriesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TimesheetEntriesTimesheetEntryIdDelete

> TimesheetEntries TimesheetEntriesTimesheetEntryIdDelete(ctx, timesheetEntryId).Execute()



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
	timesheetEntryId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimesheetEntriesAPI.TimesheetEntriesTimesheetEntryIdDelete(context.Background(), timesheetEntryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimesheetEntriesAPI.TimesheetEntriesTimesheetEntryIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TimesheetEntriesTimesheetEntryIdDelete`: TimesheetEntries
	fmt.Fprintf(os.Stdout, "Response from `TimesheetEntriesAPI.TimesheetEntriesTimesheetEntryIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timesheetEntryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTimesheetEntriesTimesheetEntryIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TimesheetEntries**](TimesheetEntries.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TimesheetEntriesTimesheetEntryIdGet

> TimesheetEntries TimesheetEntriesTimesheetEntryIdGet(ctx, timesheetEntryId).Include(include).Execute()



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
	timesheetEntryId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimesheetEntriesAPI.TimesheetEntriesTimesheetEntryIdGet(context.Background(), timesheetEntryId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimesheetEntriesAPI.TimesheetEntriesTimesheetEntryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TimesheetEntriesTimesheetEntryIdGet`: TimesheetEntries
	fmt.Fprintf(os.Stdout, "Response from `TimesheetEntriesAPI.TimesheetEntriesTimesheetEntryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timesheetEntryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTimesheetEntriesTimesheetEntryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**TimesheetEntries**](TimesheetEntries.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TimesheetEntriesTimesheetEntryIdPut

> TimesheetEntries TimesheetEntriesTimesheetEntryIdPut(ctx, timesheetEntryId).TimesheetEntriesPut(timesheetEntriesPut).Execute()



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
	timesheetEntryId := int64(789) // int64 | Model id
	timesheetEntriesPut := *openapiclient.NewTimesheetEntriesPut() // TimesheetEntriesPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimesheetEntriesAPI.TimesheetEntriesTimesheetEntryIdPut(context.Background(), timesheetEntryId).TimesheetEntriesPut(timesheetEntriesPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimesheetEntriesAPI.TimesheetEntriesTimesheetEntryIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TimesheetEntriesTimesheetEntryIdPut`: TimesheetEntries
	fmt.Fprintf(os.Stdout, "Response from `TimesheetEntriesAPI.TimesheetEntriesTimesheetEntryIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timesheetEntryId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTimesheetEntriesTimesheetEntryIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timesheetEntriesPut** | [**TimesheetEntriesPut**](TimesheetEntriesPut.md) |  | 

### Return type

[**TimesheetEntries**](TimesheetEntries.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

