# \BusinessUnitsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BusinessUnitsBusinessUnitIdDelete**](BusinessUnitsAPI.md#BusinessUnitsBusinessUnitIdDelete) | **Delete** /business_units/{business_unit_id} | 
[**BusinessUnitsBusinessUnitIdGet**](BusinessUnitsAPI.md#BusinessUnitsBusinessUnitIdGet) | **Get** /business_units/{business_unit_id} | 
[**BusinessUnitsBusinessUnitIdPut**](BusinessUnitsAPI.md#BusinessUnitsBusinessUnitIdPut) | **Put** /business_units/{business_unit_id} | 
[**BusinessUnitsGet**](BusinessUnitsAPI.md#BusinessUnitsGet) | **Get** /business_units | 
[**BusinessUnitsPost**](BusinessUnitsAPI.md#BusinessUnitsPost) | **Post** /business_units | 



## BusinessUnitsBusinessUnitIdDelete

> BusinessUnits BusinessUnitsBusinessUnitIdDelete(ctx, businessUnitId).Execute()



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
	businessUnitId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessUnitsAPI.BusinessUnitsBusinessUnitIdDelete(context.Background(), businessUnitId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessUnitsAPI.BusinessUnitsBusinessUnitIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BusinessUnitsBusinessUnitIdDelete`: BusinessUnits
	fmt.Fprintf(os.Stdout, "Response from `BusinessUnitsAPI.BusinessUnitsBusinessUnitIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessUnitId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiBusinessUnitsBusinessUnitIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BusinessUnits**](BusinessUnits.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessUnitsBusinessUnitIdGet

> BusinessUnits BusinessUnitsBusinessUnitIdGet(ctx, businessUnitId).Include(include).Execute()



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
	businessUnitId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessUnitsAPI.BusinessUnitsBusinessUnitIdGet(context.Background(), businessUnitId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessUnitsAPI.BusinessUnitsBusinessUnitIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BusinessUnitsBusinessUnitIdGet`: BusinessUnits
	fmt.Fprintf(os.Stdout, "Response from `BusinessUnitsAPI.BusinessUnitsBusinessUnitIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessUnitId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiBusinessUnitsBusinessUnitIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**BusinessUnits**](BusinessUnits.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessUnitsBusinessUnitIdPut

> BusinessUnits BusinessUnitsBusinessUnitIdPut(ctx, businessUnitId).BusinessUnitsPut(businessUnitsPut).Execute()



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
	businessUnitId := int64(789) // int64 | Model id
	businessUnitsPut := *openapiclient.NewBusinessUnitsPut() // BusinessUnitsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessUnitsAPI.BusinessUnitsBusinessUnitIdPut(context.Background(), businessUnitId).BusinessUnitsPut(businessUnitsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessUnitsAPI.BusinessUnitsBusinessUnitIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BusinessUnitsBusinessUnitIdPut`: BusinessUnits
	fmt.Fprintf(os.Stdout, "Response from `BusinessUnitsAPI.BusinessUnitsBusinessUnitIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessUnitId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiBusinessUnitsBusinessUnitIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **businessUnitsPut** | [**BusinessUnitsPut**](BusinessUnitsPut.md) |  | 

### Return type

[**BusinessUnits**](BusinessUnits.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessUnitsGet

> BusinessUnitsGet200Response BusinessUnitsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.BusinessUnitsAPI.BusinessUnitsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessUnitsAPI.BusinessUnitsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BusinessUnitsGet`: BusinessUnitsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BusinessUnitsAPI.BusinessUnitsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBusinessUnitsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**BusinessUnitsGet200Response**](BusinessUnitsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessUnitsPost

> BusinessUnitsGet200Response BusinessUnitsPost(ctx).BusinessUnitsPostRequest(businessUnitsPostRequest).Execute()



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
	businessUnitsPostRequest := *openapiclient.NewBusinessUnitsPostRequest() // BusinessUnitsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessUnitsAPI.BusinessUnitsPost(context.Background()).BusinessUnitsPostRequest(businessUnitsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessUnitsAPI.BusinessUnitsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BusinessUnitsPost`: BusinessUnitsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BusinessUnitsAPI.BusinessUnitsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBusinessUnitsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **businessUnitsPostRequest** | [**BusinessUnitsPostRequest**](BusinessUnitsPostRequest.md) |  | 

### Return type

[**BusinessUnitsGet200Response**](BusinessUnitsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

