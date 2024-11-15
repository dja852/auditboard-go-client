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

// checks if the LibraryControlTypesGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LibraryControlTypesGet200Response{}

// LibraryControlTypesGet200Response struct for LibraryControlTypesGet200Response
type LibraryControlTypesGet200Response struct {
	LibraryControlTypes []LibraryControlTypes `json:"library_control_types,omitempty"`
}

// NewLibraryControlTypesGet200Response instantiates a new LibraryControlTypesGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLibraryControlTypesGet200Response() *LibraryControlTypesGet200Response {
	this := LibraryControlTypesGet200Response{}
	return &this
}

// NewLibraryControlTypesGet200ResponseWithDefaults instantiates a new LibraryControlTypesGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLibraryControlTypesGet200ResponseWithDefaults() *LibraryControlTypesGet200Response {
	this := LibraryControlTypesGet200Response{}
	return &this
}

// GetLibraryControlTypes returns the LibraryControlTypes field value if set, zero value otherwise.
func (o *LibraryControlTypesGet200Response) GetLibraryControlTypes() []LibraryControlTypes {
	if o == nil || IsNil(o.LibraryControlTypes) {
		var ret []LibraryControlTypes
		return ret
	}
	return o.LibraryControlTypes
}

// GetLibraryControlTypesOk returns a tuple with the LibraryControlTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryControlTypesGet200Response) GetLibraryControlTypesOk() ([]LibraryControlTypes, bool) {
	if o == nil || IsNil(o.LibraryControlTypes) {
		return nil, false
	}
	return o.LibraryControlTypes, true
}

// HasLibraryControlTypes returns a boolean if a field has been set.
func (o *LibraryControlTypesGet200Response) HasLibraryControlTypes() bool {
	if o != nil && !IsNil(o.LibraryControlTypes) {
		return true
	}

	return false
}

// SetLibraryControlTypes gets a reference to the given []LibraryControlTypes and assigns it to the LibraryControlTypes field.
func (o *LibraryControlTypesGet200Response) SetLibraryControlTypes(v []LibraryControlTypes) {
	o.LibraryControlTypes = v
}

func (o LibraryControlTypesGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LibraryControlTypesGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LibraryControlTypes) {
		toSerialize["library_control_types"] = o.LibraryControlTypes
	}
	return toSerialize, nil
}

type NullableLibraryControlTypesGet200Response struct {
	value *LibraryControlTypesGet200Response
	isSet bool
}

func (v NullableLibraryControlTypesGet200Response) Get() *LibraryControlTypesGet200Response {
	return v.value
}

func (v *NullableLibraryControlTypesGet200Response) Set(val *LibraryControlTypesGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableLibraryControlTypesGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableLibraryControlTypesGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLibraryControlTypesGet200Response(val *LibraryControlTypesGet200Response) *NullableLibraryControlTypesGet200Response {
	return &NullableLibraryControlTypesGet200Response{value: val, isSet: true}
}

func (v NullableLibraryControlTypesGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLibraryControlTypesGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


