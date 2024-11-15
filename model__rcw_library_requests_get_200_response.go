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

// checks if the RcwLibraryRequestsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RcwLibraryRequestsGet200Response{}

// RcwLibraryRequestsGet200Response struct for RcwLibraryRequestsGet200Response
type RcwLibraryRequestsGet200Response struct {
	RcwLibraryRequests []RcwLibraryRequests `json:"rcw_library_requests,omitempty"`
}

// NewRcwLibraryRequestsGet200Response instantiates a new RcwLibraryRequestsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRcwLibraryRequestsGet200Response() *RcwLibraryRequestsGet200Response {
	this := RcwLibraryRequestsGet200Response{}
	return &this
}

// NewRcwLibraryRequestsGet200ResponseWithDefaults instantiates a new RcwLibraryRequestsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRcwLibraryRequestsGet200ResponseWithDefaults() *RcwLibraryRequestsGet200Response {
	this := RcwLibraryRequestsGet200Response{}
	return &this
}

// GetRcwLibraryRequests returns the RcwLibraryRequests field value if set, zero value otherwise.
func (o *RcwLibraryRequestsGet200Response) GetRcwLibraryRequests() []RcwLibraryRequests {
	if o == nil || IsNil(o.RcwLibraryRequests) {
		var ret []RcwLibraryRequests
		return ret
	}
	return o.RcwLibraryRequests
}

// GetRcwLibraryRequestsOk returns a tuple with the RcwLibraryRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcwLibraryRequestsGet200Response) GetRcwLibraryRequestsOk() ([]RcwLibraryRequests, bool) {
	if o == nil || IsNil(o.RcwLibraryRequests) {
		return nil, false
	}
	return o.RcwLibraryRequests, true
}

// HasRcwLibraryRequests returns a boolean if a field has been set.
func (o *RcwLibraryRequestsGet200Response) HasRcwLibraryRequests() bool {
	if o != nil && !IsNil(o.RcwLibraryRequests) {
		return true
	}

	return false
}

// SetRcwLibraryRequests gets a reference to the given []RcwLibraryRequests and assigns it to the RcwLibraryRequests field.
func (o *RcwLibraryRequestsGet200Response) SetRcwLibraryRequests(v []RcwLibraryRequests) {
	o.RcwLibraryRequests = v
}

func (o RcwLibraryRequestsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RcwLibraryRequestsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RcwLibraryRequests) {
		toSerialize["rcw_library_requests"] = o.RcwLibraryRequests
	}
	return toSerialize, nil
}

type NullableRcwLibraryRequestsGet200Response struct {
	value *RcwLibraryRequestsGet200Response
	isSet bool
}

func (v NullableRcwLibraryRequestsGet200Response) Get() *RcwLibraryRequestsGet200Response {
	return v.value
}

func (v *NullableRcwLibraryRequestsGet200Response) Set(val *RcwLibraryRequestsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRcwLibraryRequestsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRcwLibraryRequestsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRcwLibraryRequestsGet200Response(val *RcwLibraryRequestsGet200Response) *NullableRcwLibraryRequestsGet200Response {
	return &NullableRcwLibraryRequestsGet200Response{value: val, isSet: true}
}

func (v NullableRcwLibraryRequestsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRcwLibraryRequestsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


