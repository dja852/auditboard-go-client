# \IssuesArchivesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IssuesArchivesGet**](IssuesArchivesAPI.md#IssuesArchivesGet) | **Get** /issues_archives | 
[**IssuesArchivesIssuesArchiveIdGet**](IssuesArchivesAPI.md#IssuesArchivesIssuesArchiveIdGet) | **Get** /issues_archives/{issues_archive_id} | 



## IssuesArchivesGet

> IssuesArchivesGet200Response IssuesArchivesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.IssuesArchivesAPI.IssuesArchivesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesArchivesAPI.IssuesArchivesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssuesArchivesGet`: IssuesArchivesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesArchivesAPI.IssuesArchivesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssuesArchivesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**IssuesArchivesGet200Response**](IssuesArchivesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesArchivesIssuesArchiveIdGet

> IssuesArchives IssuesArchivesIssuesArchiveIdGet(ctx, issuesArchiveId).Include(include).Execute()



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
	issuesArchiveId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesArchivesAPI.IssuesArchivesIssuesArchiveIdGet(context.Background(), issuesArchiveId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesArchivesAPI.IssuesArchivesIssuesArchiveIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssuesArchivesIssuesArchiveIdGet`: IssuesArchives
	fmt.Fprintf(os.Stdout, "Response from `IssuesArchivesAPI.IssuesArchivesIssuesArchiveIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issuesArchiveId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesArchivesIssuesArchiveIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**IssuesArchives**](IssuesArchives.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

