# \LibraryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LibraryPreloadGet**](LibraryAPI.md#LibraryPreloadGet) | **Get** /library/preload | 



## LibraryPreloadGet

> LibraryPreload LibraryPreloadGet(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.LibraryPreloadGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.LibraryPreloadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LibraryPreloadGet`: LibraryPreload
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.LibraryPreloadGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLibraryPreloadGetRequest struct via the builder pattern


### Return type

[**LibraryPreload**](LibraryPreload.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

