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

// checks if the EsgMetricAuditableEntitiesGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsgMetricAuditableEntitiesGet200Response{}

// EsgMetricAuditableEntitiesGet200Response struct for EsgMetricAuditableEntitiesGet200Response
type EsgMetricAuditableEntitiesGet200Response struct {
	EsgMetricAuditableEntities []EsgMetricAuditableEntities `json:"esg_metric_auditable_entities,omitempty"`
}

// NewEsgMetricAuditableEntitiesGet200Response instantiates a new EsgMetricAuditableEntitiesGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsgMetricAuditableEntitiesGet200Response() *EsgMetricAuditableEntitiesGet200Response {
	this := EsgMetricAuditableEntitiesGet200Response{}
	return &this
}

// NewEsgMetricAuditableEntitiesGet200ResponseWithDefaults instantiates a new EsgMetricAuditableEntitiesGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsgMetricAuditableEntitiesGet200ResponseWithDefaults() *EsgMetricAuditableEntitiesGet200Response {
	this := EsgMetricAuditableEntitiesGet200Response{}
	return &this
}

// GetEsgMetricAuditableEntities returns the EsgMetricAuditableEntities field value if set, zero value otherwise.
func (o *EsgMetricAuditableEntitiesGet200Response) GetEsgMetricAuditableEntities() []EsgMetricAuditableEntities {
	if o == nil || IsNil(o.EsgMetricAuditableEntities) {
		var ret []EsgMetricAuditableEntities
		return ret
	}
	return o.EsgMetricAuditableEntities
}

// GetEsgMetricAuditableEntitiesOk returns a tuple with the EsgMetricAuditableEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricAuditableEntitiesGet200Response) GetEsgMetricAuditableEntitiesOk() ([]EsgMetricAuditableEntities, bool) {
	if o == nil || IsNil(o.EsgMetricAuditableEntities) {
		return nil, false
	}
	return o.EsgMetricAuditableEntities, true
}

// HasEsgMetricAuditableEntities returns a boolean if a field has been set.
func (o *EsgMetricAuditableEntitiesGet200Response) HasEsgMetricAuditableEntities() bool {
	if o != nil && !IsNil(o.EsgMetricAuditableEntities) {
		return true
	}

	return false
}

// SetEsgMetricAuditableEntities gets a reference to the given []EsgMetricAuditableEntities and assigns it to the EsgMetricAuditableEntities field.
func (o *EsgMetricAuditableEntitiesGet200Response) SetEsgMetricAuditableEntities(v []EsgMetricAuditableEntities) {
	o.EsgMetricAuditableEntities = v
}

func (o EsgMetricAuditableEntitiesGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsgMetricAuditableEntitiesGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EsgMetricAuditableEntities) {
		toSerialize["esg_metric_auditable_entities"] = o.EsgMetricAuditableEntities
	}
	return toSerialize, nil
}

type NullableEsgMetricAuditableEntitiesGet200Response struct {
	value *EsgMetricAuditableEntitiesGet200Response
	isSet bool
}

func (v NullableEsgMetricAuditableEntitiesGet200Response) Get() *EsgMetricAuditableEntitiesGet200Response {
	return v.value
}

func (v *NullableEsgMetricAuditableEntitiesGet200Response) Set(val *EsgMetricAuditableEntitiesGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEsgMetricAuditableEntitiesGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEsgMetricAuditableEntitiesGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsgMetricAuditableEntitiesGet200Response(val *EsgMetricAuditableEntitiesGet200Response) *NullableEsgMetricAuditableEntitiesGet200Response {
	return &NullableEsgMetricAuditableEntitiesGet200Response{value: val, isSet: true}
}

func (v NullableEsgMetricAuditableEntitiesGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsgMetricAuditableEntitiesGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


