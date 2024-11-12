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

// checks if the OpsAuditsIdClonePostRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpsAuditsIdClonePostRequestData{}

// OpsAuditsIdClonePostRequestData struct for OpsAuditsIdClonePostRequestData
type OpsAuditsIdClonePostRequestData struct {
	IsTemplate *bool `json:"is_template,omitempty"`
}

// NewOpsAuditsIdClonePostRequestData instantiates a new OpsAuditsIdClonePostRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsAuditsIdClonePostRequestData() *OpsAuditsIdClonePostRequestData {
	this := OpsAuditsIdClonePostRequestData{}
	return &this
}

// NewOpsAuditsIdClonePostRequestDataWithDefaults instantiates a new OpsAuditsIdClonePostRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsAuditsIdClonePostRequestDataWithDefaults() *OpsAuditsIdClonePostRequestData {
	this := OpsAuditsIdClonePostRequestData{}
	return &this
}

// GetIsTemplate returns the IsTemplate field value if set, zero value otherwise.
func (o *OpsAuditsIdClonePostRequestData) GetIsTemplate() bool {
	if o == nil || IsNil(o.IsTemplate) {
		var ret bool
		return ret
	}
	return *o.IsTemplate
}

// GetIsTemplateOk returns a tuple with the IsTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditsIdClonePostRequestData) GetIsTemplateOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTemplate) {
		return nil, false
	}
	return o.IsTemplate, true
}

// HasIsTemplate returns a boolean if a field has been set.
func (o *OpsAuditsIdClonePostRequestData) HasIsTemplate() bool {
	if o != nil && !IsNil(o.IsTemplate) {
		return true
	}

	return false
}

// SetIsTemplate gets a reference to the given bool and assigns it to the IsTemplate field.
func (o *OpsAuditsIdClonePostRequestData) SetIsTemplate(v bool) {
	o.IsTemplate = &v
}

func (o OpsAuditsIdClonePostRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpsAuditsIdClonePostRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsTemplate) {
		toSerialize["is_template"] = o.IsTemplate
	}
	return toSerialize, nil
}

type NullableOpsAuditsIdClonePostRequestData struct {
	value *OpsAuditsIdClonePostRequestData
	isSet bool
}

func (v NullableOpsAuditsIdClonePostRequestData) Get() *OpsAuditsIdClonePostRequestData {
	return v.value
}

func (v *NullableOpsAuditsIdClonePostRequestData) Set(val *OpsAuditsIdClonePostRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsAuditsIdClonePostRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsAuditsIdClonePostRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsAuditsIdClonePostRequestData(val *OpsAuditsIdClonePostRequestData) *NullableOpsAuditsIdClonePostRequestData {
	return &NullableOpsAuditsIdClonePostRequestData{value: val, isSet: true}
}

func (v NullableOpsAuditsIdClonePostRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsAuditsIdClonePostRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

