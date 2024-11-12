# \ActionPlansArchivesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionPlansArchivesActionPlansArchiveIdGet**](ActionPlansArchivesAPI.md#ActionPlansArchivesActionPlansArchiveIdGet) | **Get** /action_plans_archives/{action_plans_archive_id} | 
[**ActionPlansArchivesGet**](ActionPlansArchivesAPI.md#ActionPlansArchivesGet) | **Get** /action_plans_archives | 



## ActionPlansArchivesActionPlansArchiveIdGet

> ActionPlansArchives ActionPlansArchivesActionPlansArchiveIdGet(ctx, actionPlansArchiveId).Include(include).Execute()



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
	actionPlansArchiveId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionPlansArchivesAPI.ActionPlansArchivesActionPlansArchiveIdGet(context.Background(), actionPlansArchiveId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionPlansArchivesAPI.ActionPlansArchivesActionPlansArchiveIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionPlansArchivesActionPlansArchiveIdGet`: ActionPlansArchives
	fmt.Fprintf(os.Stdout, "Response from `ActionPlansArchivesAPI.ActionPlansArchivesActionPlansArchiveIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionPlansArchiveId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionPlansArchivesActionPlansArchiveIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**ActionPlansArchives**](ActionPlansArchives.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionPlansArchivesGet

> ActionPlansArchivesGet200Response ActionPlansArchivesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.ActionPlansArchivesAPI.ActionPlansArchivesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionPlansArchivesAPI.ActionPlansArchivesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionPlansArchivesGet`: ActionPlansArchivesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ActionPlansArchivesAPI.ActionPlansArchivesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActionPlansArchivesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**ActionPlansArchivesGet200Response**](ActionPlansArchivesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

