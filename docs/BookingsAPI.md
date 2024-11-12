# \BookingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookingsBookingIdDelete**](BookingsAPI.md#BookingsBookingIdDelete) | **Delete** /bookings/{booking_id} | 
[**BookingsBookingIdGet**](BookingsAPI.md#BookingsBookingIdGet) | **Get** /bookings/{booking_id} | 
[**BookingsBookingIdPut**](BookingsAPI.md#BookingsBookingIdPut) | **Put** /bookings/{booking_id} | 
[**BookingsGet**](BookingsAPI.md#BookingsGet) | **Get** /bookings | 
[**BookingsPost**](BookingsAPI.md#BookingsPost) | **Post** /bookings | 



## BookingsBookingIdDelete

> Bookings BookingsBookingIdDelete(ctx, bookingId).Execute()



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
	bookingId := int64(789) // int64 | Model id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BookingsAPI.BookingsBookingIdDelete(context.Background(), bookingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookingsAPI.BookingsBookingIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BookingsBookingIdDelete`: Bookings
	fmt.Fprintf(os.Stdout, "Response from `BookingsAPI.BookingsBookingIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookingId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingsBookingIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Bookings**](Bookings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingsBookingIdGet

> Bookings BookingsBookingIdGet(ctx, bookingId).Include(include).Execute()



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
	bookingId := int64(789) // int64 | Model id
	include := []string{"Include_example"} // []string | Possible sideloaded relationships (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BookingsAPI.BookingsBookingIdGet(context.Background(), bookingId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookingsAPI.BookingsBookingIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BookingsBookingIdGet`: Bookings
	fmt.Fprintf(os.Stdout, "Response from `BookingsAPI.BookingsBookingIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookingId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingsBookingIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**Bookings**](Bookings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingsBookingIdPut

> Bookings BookingsBookingIdPut(ctx, bookingId).BookingsPut(bookingsPut).Execute()



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
	bookingId := int64(789) // int64 | Model id
	bookingsPut := *openapiclient.NewBookingsPut() // BookingsPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BookingsAPI.BookingsBookingIdPut(context.Background(), bookingId).BookingsPut(bookingsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookingsAPI.BookingsBookingIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BookingsBookingIdPut`: Bookings
	fmt.Fprintf(os.Stdout, "Response from `BookingsAPI.BookingsBookingIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookingId** | **int64** | Model id | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingsBookingIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bookingsPut** | [**BookingsPut**](BookingsPut.md) |  | 

### Return type

[**Bookings**](Bookings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingsGet

> BookingsGet200Response BookingsGet(ctx).Include(include).Execute()



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
	resp, r, err := apiClient.BookingsAPI.BookingsGet(context.Background()).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookingsAPI.BookingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BookingsGet`: BookingsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BookingsAPI.BookingsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Possible sideloaded relationships | 

### Return type

[**BookingsGet200Response**](BookingsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingsPost

> BookingsGet200Response BookingsPost(ctx).BookingsPostRequest(bookingsPostRequest).Execute()



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
	bookingsPostRequest := *openapiclient.NewBookingsPostRequest() // BookingsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BookingsAPI.BookingsPost(context.Background()).BookingsPostRequest(bookingsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookingsAPI.BookingsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BookingsPost`: BookingsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BookingsAPI.BookingsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookingsPostRequest** | [**BookingsPostRequest**](BookingsPostRequest.md) |  | 

### Return type

[**BookingsGet200Response**](BookingsGet200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

