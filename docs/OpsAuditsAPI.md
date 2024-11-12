# \OpsAuditsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpsAuditsGet**](OpsAuditsAPI.md#OpsAuditsGet) | **Get** /ops_audits | 
[**OpsAuditsIdClonePost**](OpsAuditsAPI.md#OpsAuditsIdClonePost) | **Post** /ops_audits/{id}/clone | 
[**OpsAuditsOpsAuditIdCancelPut**](OpsAuditsAPI.md#OpsAuditsOpsAuditIdCancelPut) | **Put** /ops_audits/{ops_audit_id}/cancel | 
[**OpsAuditsOpsAuditIdDelete**](OpsAuditsAPI.md#OpsAuditsOpsAuditIdDelete) | **Delete** /ops_audits/{ops_audit_id} | 
[**OpsAuditsOpsAuditIdExportAuditFormsPost**](OpsAuditsAPI.md#OpsAuditsOpsAuditIdExportAuditFormsPost) | **Post** /ops_audits/{ops_audit_id}/export_audit_forms | 
[**OpsAuditsOpsAuditIdGet**](OpsAuditsAPI.md#OpsAuditsOpsAuditIdGet) | **Get** /ops_audits/{ops_audit_id} | 
[**OpsAuditsOpsAuditIdMergePost**](OpsAuditsAPI.md#OpsAuditsOpsAuditIdMergePost) | **Post** /ops_audits/{ops_audit_id}/merge | 
[**OpsAuditsOpsAuditIdPut**](OpsAuditsAPI.md#OpsAuditsOpsAuditIdPut) | **Put** /ops_audits/{ops_audit_id} | 
[**OpsAuditsOpsAuditIdStartPut**](OpsAuditsAPI.md#OpsAuditsOpsAuditIdStartPut) | **Put** /ops_audits/{ops_audit_id}/start | 
[**OpsAuditsPost**](OpsAuditsAPI.md#OpsAuditsPost) | **Post** /ops_audits | 



## OpsAuditsGet

> OpsAuditsGet200Response OpsAuditsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.OpsAuditsAPI.OpsAuditsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditsAPI.OpsAuditsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditsGet`: OpsAuditsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditsAPI.OpsAuditsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**OpsAuditsGet200Response**](OpsAuditsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditsIdClonePost

> OpsAuditsGet200Response OpsAuditsIdClonePost(ctx, id).OpsAuditsIdClonePostRequest(opsAuditsIdClonePostRequest).Execute()





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
	id := int64(789) // int64 | The id will be either the id of the OpsAudit Template or the id of the OpsAudit to be     used as a basis for the clone.
	opsAuditsIdClonePostRequest := *openapiclient.NewOpsAuditsIdClonePostRequest() // OpsAuditsIdClonePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditsAPI.OpsAuditsIdClonePost(context.Background(), id).OpsAuditsIdClonePostRequest(opsAuditsIdClonePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditsAPI.OpsAuditsIdClonePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditsIdClonePost`: OpsAuditsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditsAPI.OpsAuditsIdClonePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The id will be either the id of the OpsAudit Template or the id of the OpsAudit to be     used as a basis for the clone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditsIdClonePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opsAuditsIdClonePostRequest** | [**OpsAuditsIdClonePostRequest**](OpsAuditsIdClonePostRequest.md) |  | 

### Return type

[**OpsAuditsGet200Response**](OpsAuditsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditsOpsAuditIdCancelPut

> OpsAuditCancel OpsAuditsOpsAuditIdCancelPut(ctx, opsAuditId).OpsAuditsOpsAuditIdCancelPutRequest(opsAuditsOpsAuditIdCancelPutRequest).Execute()



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
	opsAuditId := int64(789) // int64 | OpsAudit id
	opsAuditsOpsAuditIdCancelPutRequest := *openapiclient.NewOpsAuditsOpsAuditIdCancelPutRequest() // OpsAuditsOpsAuditIdCancelPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditsAPI.OpsAuditsOpsAuditIdCancelPut(context.Background(), opsAuditId).OpsAuditsOpsAuditIdCancelPutRequest(opsAuditsOpsAuditIdCancelPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditsAPI.OpsAuditsOpsAuditIdCancelPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditsOpsAuditIdCancelPut`: OpsAuditCancel
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditsAPI.OpsAuditsOpsAuditIdCancelPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditId** | **int64** | OpsAudit id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditsOpsAuditIdCancelPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opsAuditsOpsAuditIdCancelPutRequest** | [**OpsAuditsOpsAuditIdCancelPutRequest**](OpsAuditsOpsAuditIdCancelPutRequest.md) |  | 

### Return type

[**OpsAuditCancel**](OpsAuditCancel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditsOpsAuditIdDelete

> OpsAudits OpsAuditsOpsAuditIdDelete(ctx, opsAuditId).Execute()



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
	opsAuditId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditsAPI.OpsAuditsOpsAuditIdDelete(context.Background(), opsAuditId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditsAPI.OpsAuditsOpsAuditIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditsOpsAuditIdDelete`: OpsAudits
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditsAPI.OpsAuditsOpsAuditIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditsOpsAuditIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpsAudits**](OpsAudits.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditsOpsAuditIdExportAuditFormsPost

> OpsAuditsOpsAuditIdExportAuditFormsPost200Response OpsAuditsOpsAuditIdExportAuditFormsPost(ctx, opsAuditId).Execute()





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
	opsAuditId := int64(789) // int64 | The id of the ops_audit.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditsAPI.OpsAuditsOpsAuditIdExportAuditFormsPost(context.Background(), opsAuditId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditsAPI.OpsAuditsOpsAuditIdExportAuditFormsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditsOpsAuditIdExportAuditFormsPost`: OpsAuditsOpsAuditIdExportAuditFormsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditsAPI.OpsAuditsOpsAuditIdExportAuditFormsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditId** | **int64** | The id of the ops_audit. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditsOpsAuditIdExportAuditFormsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpsAuditsOpsAuditIdExportAuditFormsPost200Response**](OpsAuditsOpsAuditIdExportAuditFormsPost200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditsOpsAuditIdGet

> OpsAudits OpsAuditsOpsAuditIdGet(ctx, opsAuditId).Include(include).Execute()



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
	opsAuditId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditsAPI.OpsAuditsOpsAuditIdGet(context.Background(), opsAuditId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditsAPI.OpsAuditsOpsAuditIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditsOpsAuditIdGet`: OpsAudits
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditsAPI.OpsAuditsOpsAuditIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditsOpsAuditIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**OpsAudits**](OpsAudits.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditsOpsAuditIdMergePost

> OpsAuditsGet200Response OpsAuditsOpsAuditIdMergePost(ctx, opsAuditId).OpsAuditsOpsAuditIdMergePostRequest(opsAuditsOpsAuditIdMergePostRequest).Execute()





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
	opsAuditId := int64(789) // int64 | The id of the OpsAudit that the Workstep will be imported to.
	opsAuditsOpsAuditIdMergePostRequest := *openapiclient.NewOpsAuditsOpsAuditIdMergePostRequest() // OpsAuditsOpsAuditIdMergePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditsAPI.OpsAuditsOpsAuditIdMergePost(context.Background(), opsAuditId).OpsAuditsOpsAuditIdMergePostRequest(opsAuditsOpsAuditIdMergePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditsAPI.OpsAuditsOpsAuditIdMergePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditsOpsAuditIdMergePost`: OpsAuditsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditsAPI.OpsAuditsOpsAuditIdMergePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditId** | **int64** | The id of the OpsAudit that the Workstep will be imported to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditsOpsAuditIdMergePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opsAuditsOpsAuditIdMergePostRequest** | [**OpsAuditsOpsAuditIdMergePostRequest**](OpsAuditsOpsAuditIdMergePostRequest.md) |  | 

### Return type

[**OpsAuditsGet200Response**](OpsAuditsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditsOpsAuditIdPut

> OpsAudits OpsAuditsOpsAuditIdPut(ctx, opsAuditId).OpsAuditsPut(opsAuditsPut).Execute()



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
	opsAuditId := int64(789) // int64 | Model id
	opsAuditsPut := *openapiclient.NewOpsAuditsPut() // OpsAuditsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditsAPI.OpsAuditsOpsAuditIdPut(context.Background(), opsAuditId).OpsAuditsPut(opsAuditsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditsAPI.OpsAuditsOpsAuditIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditsOpsAuditIdPut`: OpsAudits
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditsAPI.OpsAuditsOpsAuditIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditsOpsAuditIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opsAuditsPut** | [**OpsAuditsPut**](OpsAuditsPut.md) |  | 

### Return type

[**OpsAudits**](OpsAudits.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditsOpsAuditIdStartPut

> OpsAuditStart OpsAuditsOpsAuditIdStartPut(ctx, opsAuditId).Execute()



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
	opsAuditId := int64(789) // int64 | OpsAudit id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditsAPI.OpsAuditsOpsAuditIdStartPut(context.Background(), opsAuditId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditsAPI.OpsAuditsOpsAuditIdStartPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditsOpsAuditIdStartPut`: OpsAuditStart
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditsAPI.OpsAuditsOpsAuditIdStartPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opsAuditId** | **int64** | OpsAudit id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditsOpsAuditIdStartPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpsAuditStart**](OpsAuditStart.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpsAuditsPost

> OpsAuditsGet200Response OpsAuditsPost(ctx).OpsAuditsPostRequest(opsAuditsPostRequest).Execute()



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
	opsAuditsPostRequest := *openapiclient.NewOpsAuditsPostRequest() // OpsAuditsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAuditsAPI.OpsAuditsPost(context.Background()).OpsAuditsPostRequest(opsAuditsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAuditsAPI.OpsAuditsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsAuditsPost`: OpsAuditsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OpsAuditsAPI.OpsAuditsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpsAuditsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opsAuditsPostRequest** | [**OpsAuditsPostRequest**](OpsAuditsPostRequest.md) |  | 

### Return type

[**OpsAuditsGet200Response**](OpsAuditsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

