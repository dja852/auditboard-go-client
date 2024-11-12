# \ControlsDataArchivesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ControlsDataArchivesControlsDataArchiveIdGet**](ControlsDataArchivesAPI.md#ControlsDataArchivesControlsDataArchiveIdGet) | **Get** /controls_data_archives/{controls_data_archive_id} | 
[**ControlsDataArchivesGet**](ControlsDataArchivesAPI.md#ControlsDataArchivesGet) | **Get** /controls_data_archives | 



## ControlsDataArchivesControlsDataArchiveIdGet

> ControlsDataArchives ControlsDataArchivesControlsDataArchiveIdGet(ctx, controlsDataArchiveId).Include(include).Execute()



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
	controlsDataArchiveId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControlsDataArchivesAPI.ControlsDataArchivesControlsDataArchiveIdGet(context.Background(), controlsDataArchiveId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControlsDataArchivesAPI.ControlsDataArchivesControlsDataArchiveIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlsDataArchivesControlsDataArchiveIdGet`: ControlsDataArchives
	fmt.Fprintf(os.Stdout, "Response from `ControlsDataArchivesAPI.ControlsDataArchivesControlsDataArchiveIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**controlsDataArchiveId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiControlsDataArchivesControlsDataArchiveIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**ControlsDataArchives**](ControlsDataArchives.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ControlsDataArchivesGet

> ControlsDataArchivesGet200Response ControlsDataArchivesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.ControlsDataArchivesAPI.ControlsDataArchivesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControlsDataArchivesAPI.ControlsDataArchivesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlsDataArchivesGet`: ControlsDataArchivesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ControlsDataArchivesAPI.ControlsDataArchivesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiControlsDataArchivesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**ControlsDataArchivesGet200Response**](ControlsDataArchivesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

