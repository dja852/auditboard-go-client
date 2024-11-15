/*
AuditBoard Developer Portal API Documentation

External API reference documentation.

API version: 23.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auditboard

import (
	"encoding/json"
)

// checks if the BookingsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookingsPostRequest{}

// BookingsPostRequest struct for BookingsPostRequest
type BookingsPostRequest struct {
	Booking *Bookings `json:"booking,omitempty"`
}

// NewBookingsPostRequest instantiates a new BookingsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookingsPostRequest() *BookingsPostRequest {
	this := BookingsPostRequest{}
	return &this
}

// NewBookingsPostRequestWithDefaults instantiates a new BookingsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookingsPostRequestWithDefaults() *BookingsPostRequest {
	this := BookingsPostRequest{}
	return &this
}

// GetBooking returns the Booking field value if set, zero value otherwise.
func (o *BookingsPostRequest) GetBooking() Bookings {
	if o == nil || IsNil(o.Booking) {
		var ret Bookings
		return ret
	}
	return *o.Booking
}

// GetBookingOk returns a tuple with the Booking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookingsPostRequest) GetBookingOk() (*Bookings, bool) {
	if o == nil || IsNil(o.Booking) {
		return nil, false
	}
	return o.Booking, true
}

// HasBooking returns a boolean if a field has been set.
func (o *BookingsPostRequest) HasBooking() bool {
	if o != nil && !IsNil(o.Booking) {
		return true
	}

	return false
}

// SetBooking gets a reference to the given Bookings and assigns it to the Booking field.
func (o *BookingsPostRequest) SetBooking(v Bookings) {
	o.Booking = &v
}

func (o BookingsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookingsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Booking) {
		toSerialize["booking"] = o.Booking
	}
	return toSerialize, nil
}

type NullableBookingsPostRequest struct {
	value *BookingsPostRequest
	isSet bool
}

func (v NullableBookingsPostRequest) Get() *BookingsPostRequest {
	return v.value
}

func (v *NullableBookingsPostRequest) Set(val *BookingsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBookingsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBookingsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookingsPostRequest(val *BookingsPostRequest) *NullableBookingsPostRequest {
	return &NullableBookingsPostRequest{value: val, isSet: true}
}

func (v NullableBookingsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookingsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


