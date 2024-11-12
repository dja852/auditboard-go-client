# \TeamsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsGet**](TeamsAPI.md#TeamsGet) | **Get** /teams | 
[**TeamsPost**](TeamsAPI.md#TeamsPost) | **Post** /teams | 
[**TeamsTeamIdDelete**](TeamsAPI.md#TeamsTeamIdDelete) | **Delete** /teams/{team_id} | 
[**TeamsTeamIdGet**](TeamsAPI.md#TeamsTeamIdGet) | **Get** /teams/{team_id} | 
[**TeamsTeamIdPut**](TeamsAPI.md#TeamsTeamIdPut) | **Put** /teams/{team_id} | 



## TeamsGet

> TeamsGet200Response TeamsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.TeamsAPI.TeamsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.TeamsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsGet`: TeamsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.TeamsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**TeamsGet200Response**](TeamsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPost

> TeamsGet200Response TeamsPost(ctx).TeamsPostRequest(teamsPostRequest).Execute()



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
	teamsPostRequest := *openapiclient.NewTeamsPostRequest() // TeamsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.TeamsPost(context.Background()).TeamsPostRequest(teamsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.TeamsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsPost`: TeamsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.TeamsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamsPostRequest** | [**TeamsPostRequest**](TeamsPostRequest.md) |  | 

### Return type

[**TeamsGet200Response**](TeamsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsTeamIdDelete

> Teams TeamsTeamIdDelete(ctx, teamId).Execute()



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
	teamId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.TeamsTeamIdDelete(context.Background(), teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.TeamsTeamIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsTeamIdDelete`: Teams
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.TeamsTeamIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Teams**](Teams.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsTeamIdGet

> Teams TeamsTeamIdGet(ctx, teamId).Include(include).Execute()



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
	teamId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.TeamsTeamIdGet(context.Background(), teamId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.TeamsTeamIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsTeamIdGet`: Teams
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.TeamsTeamIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**Teams**](Teams.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsTeamIdPut

> Teams TeamsTeamIdPut(ctx, teamId).TeamsPut(teamsPut).Execute()



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
	teamId := int64(789) // int64 | Model id
	teamsPut := *openapiclient.NewTeamsPut() // TeamsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.TeamsTeamIdPut(context.Background(), teamId).TeamsPut(teamsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.TeamsTeamIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsTeamIdPut`: Teams
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.TeamsTeamIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamsPut** | [**TeamsPut**](TeamsPut.md) |  | 

### Return type

[**Teams**](Teams.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

