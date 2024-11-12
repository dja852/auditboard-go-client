# \DepartmentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DepartmentsDepartmentIdDelete**](DepartmentsAPI.md#DepartmentsDepartmentIdDelete) | **Delete** /departments/{department_id} | 
[**DepartmentsDepartmentIdGet**](DepartmentsAPI.md#DepartmentsDepartmentIdGet) | **Get** /departments/{department_id} | 
[**DepartmentsDepartmentIdPut**](DepartmentsAPI.md#DepartmentsDepartmentIdPut) | **Put** /departments/{department_id} | 
[**DepartmentsGet**](DepartmentsAPI.md#DepartmentsGet) | **Get** /departments | 
[**DepartmentsPost**](DepartmentsAPI.md#DepartmentsPost) | **Post** /departments | 



## DepartmentsDepartmentIdDelete

> Departments DepartmentsDepartmentIdDelete(ctx, departmentId).Execute()



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
	departmentId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DepartmentsAPI.DepartmentsDepartmentIdDelete(context.Background(), departmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.DepartmentsDepartmentIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DepartmentsDepartmentIdDelete`: Departments
	fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.DepartmentsDepartmentIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**departmentId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDepartmentsDepartmentIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Departments**](Departments.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DepartmentsDepartmentIdGet

> Departments DepartmentsDepartmentIdGet(ctx, departmentId).Include(include).Execute()



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
	departmentId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DepartmentsAPI.DepartmentsDepartmentIdGet(context.Background(), departmentId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.DepartmentsDepartmentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DepartmentsDepartmentIdGet`: Departments
	fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.DepartmentsDepartmentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**departmentId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDepartmentsDepartmentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**Departments**](Departments.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DepartmentsDepartmentIdPut

> Departments DepartmentsDepartmentIdPut(ctx, departmentId).DepartmentsPut(departmentsPut).Execute()



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
	departmentId := int64(789) // int64 | Model id
	departmentsPut := *openapiclient.NewDepartmentsPut() // DepartmentsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DepartmentsAPI.DepartmentsDepartmentIdPut(context.Background(), departmentId).DepartmentsPut(departmentsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.DepartmentsDepartmentIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DepartmentsDepartmentIdPut`: Departments
	fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.DepartmentsDepartmentIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**departmentId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDepartmentsDepartmentIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **departmentsPut** | [**DepartmentsPut**](DepartmentsPut.md) |  | 

### Return type

[**Departments**](Departments.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DepartmentsGet

> DepartmentsGet200Response DepartmentsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.DepartmentsAPI.DepartmentsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.DepartmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DepartmentsGet`: DepartmentsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.DepartmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDepartmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**DepartmentsGet200Response**](DepartmentsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DepartmentsPost

> DepartmentsGet200Response DepartmentsPost(ctx).DepartmentsPostRequest(departmentsPostRequest).Execute()



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
	departmentsPostRequest := *openapiclient.NewDepartmentsPostRequest() // DepartmentsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DepartmentsAPI.DepartmentsPost(context.Background()).DepartmentsPostRequest(departmentsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.DepartmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DepartmentsPost`: DepartmentsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.DepartmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDepartmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **departmentsPostRequest** | [**DepartmentsPostRequest**](DepartmentsPostRequest.md) |  | 

### Return type

[**DepartmentsGet200Response**](DepartmentsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

