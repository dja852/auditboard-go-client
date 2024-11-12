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

// checks if the AuditableEntityTypesGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditableEntityTypesGet200Response{}

// AuditableEntityTypesGet200Response struct for AuditableEntityTypesGet200Response
type AuditableEntityTypesGet200Response struct {
	AuditableEntityTypes []AuditableEntityTypes `json:"auditable_entity_types,omitempty"`
}

// NewAuditableEntityTypesGet200Response instantiates a new AuditableEntityTypesGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditableEntityTypesGet200Response() *AuditableEntityTypesGet200Response {
	this := AuditableEntityTypesGet200Response{}
	return &this
}

// NewAuditableEntityTypesGet200ResponseWithDefaults instantiates a new AuditableEntityTypesGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditableEntityTypesGet200ResponseWithDefaults() *AuditableEntityTypesGet200Response {
	this := AuditableEntityTypesGet200Response{}
	return &this
}

// GetAuditableEntityTypes returns the AuditableEntityTypes field value if set, zero value otherwise.
func (o *AuditableEntityTypesGet200Response) GetAuditableEntityTypes() []AuditableEntityTypes {
	if o == nil || IsNil(o.AuditableEntityTypes) {
		var ret []AuditableEntityTypes
		return ret
	}
	return o.AuditableEntityTypes
}

// GetAuditableEntityTypesOk returns a tuple with the AuditableEntityTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditableEntityTypesGet200Response) GetAuditableEntityTypesOk() ([]AuditableEntityTypes, bool) {
	if o == nil || IsNil(o.AuditableEntityTypes) {
		return nil, false
	}
	return o.AuditableEntityTypes, true
}

// HasAuditableEntityTypes returns a boolean if a field has been set.
func (o *AuditableEntityTypesGet200Response) HasAuditableEntityTypes() bool {
	if o != nil && !IsNil(o.AuditableEntityTypes) {
		return true
	}

	return false
}

// SetAuditableEntityTypes gets a reference to the given []AuditableEntityTypes and assigns it to the AuditableEntityTypes field.
func (o *AuditableEntityTypesGet200Response) SetAuditableEntityTypes(v []AuditableEntityTypes) {
	o.AuditableEntityTypes = v
}

func (o AuditableEntityTypesGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditableEntityTypesGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuditableEntityTypes) {
		toSerialize["auditable_entity_types"] = o.AuditableEntityTypes
	}
	return toSerialize, nil
}

type NullableAuditableEntityTypesGet200Response struct {
	value *AuditableEntityTypesGet200Response
	isSet bool
}

func (v NullableAuditableEntityTypesGet200Response) Get() *AuditableEntityTypesGet200Response {
	return v.value
}

func (v *NullableAuditableEntityTypesGet200Response) Set(val *AuditableEntityTypesGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditableEntityTypesGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditableEntityTypesGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditableEntityTypesGet200Response(val *AuditableEntityTypesGet200Response) *NullableAuditableEntityTypesGet200Response {
	return &NullableAuditableEntityTypesGet200Response{value: val, isSet: true}
}

func (v NullableAuditableEntityTypesGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditableEntityTypesGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


