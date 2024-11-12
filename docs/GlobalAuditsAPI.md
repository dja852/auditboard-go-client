# \GlobalAuditsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlobalAuditsGet**](GlobalAuditsAPI.md#GlobalAuditsGet) | **Get** /global_audits | 
[**GlobalAuditsGlobalAuditIdGet**](GlobalAuditsAPI.md#GlobalAuditsGlobalAuditIdGet) | **Get** /global_audits/{global_audit_id} | 



## GlobalAuditsGet

> GlobalAuditsGet200Response GlobalAuditsGet(ctx).CreatedAtFilterOperator(createdAtFilterOperator).ObjectType(objectType).Limit(limit).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/dja852/auditboard-go-client"
)

func main() {
	createdAtFilterOperator := time.Now() // time.Time | Refer to the API filtering section <a href=\"https://developer.auditboard.com/docs/api-filtering\">here</a> for more information about filter operators. This query parameter can be used to filter the results by created_at dates. Ex: `?created_at[$gt]=[2022-01-19T18:11:51.107Z]&created_at[$lt]=[2022-03-11]` (optional)
	objectType := []string{"ObjectType_example"} // []string | The object_type parameter will filter the response to contain only results that are of the object_type. (optional)
	limit := int32(56) // int32 | Limits the returned result set (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalAuditsAPI.GlobalAuditsGet(context.Background()).CreatedAtFilterOperator(createdAtFilterOperator).ObjectType(objectType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAuditsAPI.GlobalAuditsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalAuditsGet`: GlobalAuditsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalAuditsAPI.GlobalAuditsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlobalAuditsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAtFilterOperator** | **time.Time** | Refer to the API filtering section &lt;a href&#x3D;\&quot;https://developer.auditboard.com/docs/api-filtering\&quot;&gt;here&lt;/a&gt; for more information about filter operators. This query parameter can be used to filter the results by created_at dates. Ex: &#x60;?created_at[$gt]&#x3D;[2022-01-19T18:11:51.107Z]&amp;created_at[$lt]&#x3D;[2022-03-11]&#x60; | 
 **objectType** | **[]string** | The object_type parameter will filter the response to contain only results that are of the object_type. | 
 **limit** | **int32** | Limits the returned result set | 

### Return type

[**GlobalAuditsGet200Response**](GlobalAuditsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlobalAuditsGlobalAuditIdGet

> GlobalAudits GlobalAuditsGlobalAuditIdGet(ctx, globalAuditId).Include(include).Execute()



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
	globalAuditId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalAuditsAPI.GlobalAuditsGlobalAuditIdGet(context.Background(), globalAuditId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAuditsAPI.GlobalAuditsGlobalAuditIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalAuditsGlobalAuditIdGet`: GlobalAudits
	fmt.Fprintf(os.Stdout, "Response from `GlobalAuditsAPI.GlobalAuditsGlobalAuditIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**globalAuditId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalAuditsGlobalAuditIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**GlobalAudits**](GlobalAudits.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

