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

// checks if the BookingsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookingsGet200Response{}

// BookingsGet200Response struct for BookingsGet200Response
type BookingsGet200Response struct {
	Bookings []Bookings `json:"bookings,omitempty"`
}

// NewBookingsGet200Response instantiates a new BookingsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookingsGet200Response() *BookingsGet200Response {
	this := BookingsGet200Response{}
	return &this
}

// NewBookingsGet200ResponseWithDefaults instantiates a new BookingsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookingsGet200ResponseWithDefaults() *BookingsGet200Response {
	this := BookingsGet200Response{}
	return &this
}

// GetBookings returns the Bookings field value if set, zero value otherwise.
func (o *BookingsGet200Response) GetBookings() []Bookings {
	if o == nil || IsNil(o.Bookings) {
		var ret []Bookings
		return ret
	}
	return o.Bookings
}

// GetBookingsOk returns a tuple with the Bookings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookingsGet200Response) GetBookingsOk() ([]Bookings, bool) {
	if o == nil || IsNil(o.Bookings) {
		return nil, false
	}
	return o.Bookings, true
}

// HasBookings returns a boolean if a field has been set.
func (o *BookingsGet200Response) HasBookings() bool {
	if o != nil && !IsNil(o.Bookings) {
		return true
	}

	return false
}

// SetBookings gets a reference to the given []Bookings and assigns it to the Bookings field.
func (o *BookingsGet200Response) SetBookings(v []Bookings) {
	o.Bookings = v
}

func (o BookingsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookingsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bookings) {
		toSerialize["bookings"] = o.Bookings
	}
	return toSerialize, nil
}

type NullableBookingsGet200Response struct {
	value *BookingsGet200Response
	isSet bool
}

func (v NullableBookingsGet200Response) Get() *BookingsGet200Response {
	return v.value
}

func (v *NullableBookingsGet200Response) Set(val *BookingsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableBookingsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableBookingsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookingsGet200Response(val *BookingsGet200Response) *NullableBookingsGet200Response {
	return &NullableBookingsGet200Response{value: val, isSet: true}
}

func (v NullableBookingsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookingsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


