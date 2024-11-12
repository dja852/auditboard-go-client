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

// checks if the LibraryControlClassificationsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LibraryControlClassificationsGet200Response{}

// LibraryControlClassificationsGet200Response struct for LibraryControlClassificationsGet200Response
type LibraryControlClassificationsGet200Response struct {
	LibraryControlClassifications []LibraryControlClassifications `json:"library_control_classifications,omitempty"`
}

// NewLibraryControlClassificationsGet200Response instantiates a new LibraryControlClassificationsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLibraryControlClassificationsGet200Response() *LibraryControlClassificationsGet200Response {
	this := LibraryControlClassificationsGet200Response{}
	return &this
}

// NewLibraryControlClassificationsGet200ResponseWithDefaults instantiates a new LibraryControlClassificationsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLibraryControlClassificationsGet200ResponseWithDefaults() *LibraryControlClassificationsGet200Response {
	this := LibraryControlClassificationsGet200Response{}
	return &this
}

// GetLibraryControlClassifications returns the LibraryControlClassifications field value if set, zero value otherwise.
func (o *LibraryControlClassificationsGet200Response) GetLibraryControlClassifications() []LibraryControlClassifications {
	if o == nil || IsNil(o.LibraryControlClassifications) {
		var ret []LibraryControlClassifications
		return ret
	}
	return o.LibraryControlClassifications
}

// GetLibraryControlClassificationsOk returns a tuple with the LibraryControlClassifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryControlClassificationsGet200Response) GetLibraryControlClassificationsOk() ([]LibraryControlClassifications, bool) {
	if o == nil || IsNil(o.LibraryControlClassifications) {
		return nil, false
	}
	return o.LibraryControlClassifications, true
}

// HasLibraryControlClassifications returns a boolean if a field has been set.
func (o *LibraryControlClassificationsGet200Response) HasLibraryControlClassifications() bool {
	if o != nil && !IsNil(o.LibraryControlClassifications) {
		return true
	}

	return false
}

// SetLibraryControlClassifications gets a reference to the given []LibraryControlClassifications and assigns it to the LibraryControlClassifications field.
func (o *LibraryControlClassificationsGet200Response) SetLibraryControlClassifications(v []LibraryControlClassifications) {
	o.LibraryControlClassifications = v
}

func (o LibraryControlClassificationsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LibraryControlClassificationsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LibraryControlClassifications) {
		toSerialize["library_control_classifications"] = o.LibraryControlClassifications
	}
	return toSerialize, nil
}

type NullableLibraryControlClassificationsGet200Response struct {
	value *LibraryControlClassificationsGet200Response
	isSet bool
}

func (v NullableLibraryControlClassificationsGet200Response) Get() *LibraryControlClassificationsGet200Response {
	return v.value
}

func (v *NullableLibraryControlClassificationsGet200Response) Set(val *LibraryControlClassificationsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableLibraryControlClassificationsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableLibraryControlClassificationsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLibraryControlClassificationsGet200Response(val *LibraryControlClassificationsGet200Response) *NullableLibraryControlClassificationsGet200Response {
	return &NullableLibraryControlClassificationsGet200Response{value: val, isSet: true}
}

func (v NullableLibraryControlClassificationsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLibraryControlClassificationsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


