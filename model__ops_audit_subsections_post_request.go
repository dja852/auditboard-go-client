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

// checks if the OpsAuditSubsectionsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpsAuditSubsectionsPostRequest{}

// OpsAuditSubsectionsPostRequest struct for OpsAuditSubsectionsPostRequest
type OpsAuditSubsectionsPostRequest struct {
	OpsAuditSubsection *OpsAuditSubsections `json:"ops_audit _subsection,omitempty"`
}

// NewOpsAuditSubsectionsPostRequest instantiates a new OpsAuditSubsectionsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsAuditSubsectionsPostRequest() *OpsAuditSubsectionsPostRequest {
	this := OpsAuditSubsectionsPostRequest{}
	return &this
}

// NewOpsAuditSubsectionsPostRequestWithDefaults instantiates a new OpsAuditSubsectionsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsAuditSubsectionsPostRequestWithDefaults() *OpsAuditSubsectionsPostRequest {
	this := OpsAuditSubsectionsPostRequest{}
	return &this
}

// GetOpsAuditSubsection returns the OpsAuditSubsection field value if set, zero value otherwise.
func (o *OpsAuditSubsectionsPostRequest) GetOpsAuditSubsection() OpsAuditSubsections {
	if o == nil || IsNil(o.OpsAuditSubsection) {
		var ret OpsAuditSubsections
		return ret
	}
	return *o.OpsAuditSubsection
}

// GetOpsAuditSubsectionOk returns a tuple with the OpsAuditSubsection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditSubsectionsPostRequest) GetOpsAuditSubsectionOk() (*OpsAuditSubsections, bool) {
	if o == nil || IsNil(o.OpsAuditSubsection) {
		return nil, false
	}
	return o.OpsAuditSubsection, true
}

// HasOpsAuditSubsection returns a boolean if a field has been set.
func (o *OpsAuditSubsectionsPostRequest) HasOpsAuditSubsection() bool {
	if o != nil && !IsNil(o.OpsAuditSubsection) {
		return true
	}

	return false
}

// SetOpsAuditSubsection gets a reference to the given OpsAuditSubsections and assigns it to the OpsAuditSubsection field.
func (o *OpsAuditSubsectionsPostRequest) SetOpsAuditSubsection(v OpsAuditSubsections) {
	o.OpsAuditSubsection = &v
}

func (o OpsAuditSubsectionsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpsAuditSubsectionsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OpsAuditSubsection) {
		toSerialize["ops_audit _subsection"] = o.OpsAuditSubsection
	}
	return toSerialize, nil
}

type NullableOpsAuditSubsectionsPostRequest struct {
	value *OpsAuditSubsectionsPostRequest
	isSet bool
}

func (v NullableOpsAuditSubsectionsPostRequest) Get() *OpsAuditSubsectionsPostRequest {
	return v.value
}

func (v *NullableOpsAuditSubsectionsPostRequest) Set(val *OpsAuditSubsectionsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsAuditSubsectionsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsAuditSubsectionsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsAuditSubsectionsPostRequest(val *OpsAuditSubsectionsPostRequest) *NullableOpsAuditSubsectionsPostRequest {
	return &NullableOpsAuditSubsectionsPostRequest{value: val, isSet: true}
}

func (v NullableOpsAuditSubsectionsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsAuditSubsectionsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


