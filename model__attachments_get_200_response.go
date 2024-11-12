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

// checks if the AttachmentsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachmentsGet200Response{}

// AttachmentsGet200Response struct for AttachmentsGet200Response
type AttachmentsGet200Response struct {
	Attachments []Attachments `json:"attachments,omitempty"`
}

// NewAttachmentsGet200Response instantiates a new AttachmentsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentsGet200Response() *AttachmentsGet200Response {
	this := AttachmentsGet200Response{}
	return &this
}

// NewAttachmentsGet200ResponseWithDefaults instantiates a new AttachmentsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentsGet200ResponseWithDefaults() *AttachmentsGet200Response {
	this := AttachmentsGet200Response{}
	return &this
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *AttachmentsGet200Response) GetAttachments() []Attachments {
	if o == nil || IsNil(o.Attachments) {
		var ret []Attachments
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentsGet200Response) GetAttachmentsOk() ([]Attachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *AttachmentsGet200Response) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachments and assigns it to the Attachments field.
func (o *AttachmentsGet200Response) SetAttachments(v []Attachments) {
	o.Attachments = v
}

func (o AttachmentsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachmentsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	return toSerialize, nil
}

type NullableAttachmentsGet200Response struct {
	value *AttachmentsGet200Response
	isSet bool
}

func (v NullableAttachmentsGet200Response) Get() *AttachmentsGet200Response {
	return v.value
}

func (v *NullableAttachmentsGet200Response) Set(val *AttachmentsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentsGet200Response(val *AttachmentsGet200Response) *NullableAttachmentsGet200Response {
	return &NullableAttachmentsGet200Response{value: val, isSet: true}
}

func (v NullableAttachmentsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

