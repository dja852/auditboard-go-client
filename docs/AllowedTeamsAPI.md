# \AllowedTeamsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllowedTeamsAllowedTeamIdDelete**](AllowedTeamsAPI.md#AllowedTeamsAllowedTeamIdDelete) | **Delete** /allowed_teams/{allowed_team_id} | 
[**AllowedTeamsAllowedTeamIdGet**](AllowedTeamsAPI.md#AllowedTeamsAllowedTeamIdGet) | **Get** /allowed_teams/{allowed_team_id} | 
[**AllowedTeamsAllowedTeamIdPut**](AllowedTeamsAPI.md#AllowedTeamsAllowedTeamIdPut) | **Put** /allowed_teams/{allowed_team_id} | 
[**AllowedTeamsGet**](AllowedTeamsAPI.md#AllowedTeamsGet) | **Get** /allowed_teams | 
[**AllowedTeamsPost**](AllowedTeamsAPI.md#AllowedTeamsPost) | **Post** /allowed_teams | 



## AllowedTeamsAllowedTeamIdDelete

> AllowedTeams AllowedTeamsAllowedTeamIdDelete(ctx, allowedTeamId).Execute()



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
	allowedTeamId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllowedTeamsAPI.AllowedTeamsAllowedTeamIdDelete(context.Background(), allowedTeamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedTeamsAPI.AllowedTeamsAllowedTeamIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AllowedTeamsAllowedTeamIdDelete`: AllowedTeams
	fmt.Fprintf(os.Stdout, "Response from `AllowedTeamsAPI.AllowedTeamsAllowedTeamIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allowedTeamId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAllowedTeamsAllowedTeamIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AllowedTeams**](AllowedTeams.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllowedTeamsAllowedTeamIdGet

> AllowedTeams AllowedTeamsAllowedTeamIdGet(ctx, allowedTeamId).Include(include).Execute()



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
	allowedTeamId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllowedTeamsAPI.AllowedTeamsAllowedTeamIdGet(context.Background(), allowedTeamId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedTeamsAPI.AllowedTeamsAllowedTeamIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AllowedTeamsAllowedTeamIdGet`: AllowedTeams
	fmt.Fprintf(os.Stdout, "Response from `AllowedTeamsAPI.AllowedTeamsAllowedTeamIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allowedTeamId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAllowedTeamsAllowedTeamIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AllowedTeams**](AllowedTeams.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllowedTeamsAllowedTeamIdPut

> AllowedTeams AllowedTeamsAllowedTeamIdPut(ctx, allowedTeamId).AllowedTeamsPut(allowedTeamsPut).Execute()



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
	allowedTeamId := int64(789) // int64 | Model id
	allowedTeamsPut := *openapiclient.NewAllowedTeamsPut() // AllowedTeamsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllowedTeamsAPI.AllowedTeamsAllowedTeamIdPut(context.Background(), allowedTeamId).AllowedTeamsPut(allowedTeamsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedTeamsAPI.AllowedTeamsAllowedTeamIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AllowedTeamsAllowedTeamIdPut`: AllowedTeams
	fmt.Fprintf(os.Stdout, "Response from `AllowedTeamsAPI.AllowedTeamsAllowedTeamIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allowedTeamId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAllowedTeamsAllowedTeamIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allowedTeamsPut** | [**AllowedTeamsPut**](AllowedTeamsPut.md) |  | 

### Return type

[**AllowedTeams**](AllowedTeams.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllowedTeamsGet

> AllowedTeamsGet200Response AllowedTeamsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.AllowedTeamsAPI.AllowedTeamsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedTeamsAPI.AllowedTeamsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AllowedTeamsGet`: AllowedTeamsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AllowedTeamsAPI.AllowedTeamsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllowedTeamsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**AllowedTeamsGet200Response**](AllowedTeamsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllowedTeamsPost

> AllowedTeamsGet200Response AllowedTeamsPost(ctx).AllowedTeamsPostRequest(allowedTeamsPostRequest).Execute()



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
	allowedTeamsPostRequest := *openapiclient.NewAllowedTeamsPostRequest() // AllowedTeamsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllowedTeamsAPI.AllowedTeamsPost(context.Background()).AllowedTeamsPostRequest(allowedTeamsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedTeamsAPI.AllowedTeamsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AllowedTeamsPost`: AllowedTeamsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AllowedTeamsAPI.AllowedTeamsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllowedTeamsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allowedTeamsPostRequest** | [**AllowedTeamsPostRequest**](AllowedTeamsPostRequest.md) |  | 

### Return type

[**AllowedTeamsGet200Response**](AllowedTeamsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

