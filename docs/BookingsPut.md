# BookingsPut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Booking** | Pointer to [**BookingsPutBooking**](BookingsPutBooking.md) |  | [optional] 
**PreviousValues** | Pointer to [**BookingsPutPreviousValues**](BookingsPutPreviousValues.md) |  | [optional] 

## Methods

### NewBookingsPut

`func NewBookingsPut() *BookingsPut`

NewBookingsPut instantiates a new BookingsPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingsPutWithDefaults

`func NewBookingsPutWithDefaults() *BookingsPut`

NewBookingsPutWithDefaults instantiates a new BookingsPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBooking

`func (o *BookingsPut) GetBooking() BookingsPutBooking`

GetBooking returns the Booking field if non-nil, zero value otherwise.

### GetBookingOk

`func (o *BookingsPut) GetBookingOk() (*BookingsPutBooking, bool)`

GetBookingOk returns a tuple with the Booking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooking

`func (o *BookingsPut) SetBooking(v BookingsPutBooking)`

SetBooking sets Booking field to given value.

### HasBooking

`func (o *BookingsPut) HasBooking() bool`

HasBooking returns a boolean if a field has been set.

### GetPreviousValues

`func (o *BookingsPut) GetPreviousValues() BookingsPutPreviousValues`

GetPreviousValues returns the PreviousValues field if non-nil, zero value otherwise.

### GetPreviousValuesOk

`func (o *BookingsPut) GetPreviousValuesOk() (*BookingsPutPreviousValues, bool)`

GetPreviousValuesOk returns a tuple with the PreviousValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousValues

`func (o *BookingsPut) SetPreviousValues(v BookingsPutPreviousValues)`

SetPreviousValues sets PreviousValues field to given value.

### HasPreviousValues

`func (o *BookingsPut) HasPreviousValues() bool`

HasPreviousValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


