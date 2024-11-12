# \NotificationMessagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationMessagesGet**](NotificationMessagesAPI.md#NotificationMessagesGet) | **Get** /notification_messages | 
[**NotificationMessagesNotificationMessageIdDelete**](NotificationMessagesAPI.md#NotificationMessagesNotificationMessageIdDelete) | **Delete** /notification_messages/{notification_message_id} | 
[**NotificationMessagesNotificationMessageIdGet**](NotificationMessagesAPI.md#NotificationMessagesNotificationMessageIdGet) | **Get** /notification_messages/{notification_message_id} | 



## NotificationMessagesGet

> NotificationMessagesGet200Response NotificationMessagesGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.NotificationMessagesAPI.NotificationMessagesGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationMessagesAPI.NotificationMessagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationMessagesGet`: NotificationMessagesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationMessagesAPI.NotificationMessagesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationMessagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**NotificationMessagesGet200Response**](NotificationMessagesGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationMessagesNotificationMessageIdDelete

> NotificationMessages NotificationMessagesNotificationMessageIdDelete(ctx, notificationMessageId).Execute()



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
	notificationMessageId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationMessagesAPI.NotificationMessagesNotificationMessageIdDelete(context.Background(), notificationMessageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationMessagesAPI.NotificationMessagesNotificationMessageIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationMessagesNotificationMessageIdDelete`: NotificationMessages
	fmt.Fprintf(os.Stdout, "Response from `NotificationMessagesAPI.NotificationMessagesNotificationMessageIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationMessageId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationMessagesNotificationMessageIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationMessages**](NotificationMessages.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationMessagesNotificationMessageIdGet

> NotificationMessages NotificationMessagesNotificationMessageIdGet(ctx, notificationMessageId).Include(include).Execute()



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
	notificationMessageId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationMessagesAPI.NotificationMessagesNotificationMessageIdGet(context.Background(), notificationMessageId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationMessagesAPI.NotificationMessagesNotificationMessageIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationMessagesNotificationMessageIdGet`: NotificationMessages
	fmt.Fprintf(os.Stdout, "Response from `NotificationMessagesAPI.NotificationMessagesNotificationMessageIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationMessageId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationMessagesNotificationMessageIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**NotificationMessages**](NotificationMessages.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

