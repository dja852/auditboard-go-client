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

// checks if the EsgMetricAuditableEntitiesPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsgMetricAuditableEntitiesPut{}

// EsgMetricAuditableEntitiesPut struct for EsgMetricAuditableEntitiesPut
type EsgMetricAuditableEntitiesPut struct {
	EsgMetricAuditableEntity *EsgMetricAuditableEntitiesPutEsgMetricAuditableEntity `json:"esg_metric_auditable_entity,omitempty"`
	PreviousValues *EsgMetricAuditableEntitiesPutPreviousValues `json:"previous_values,omitempty"`
}

// NewEsgMetricAuditableEntitiesPut instantiates a new EsgMetricAuditableEntitiesPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsgMetricAuditableEntitiesPut() *EsgMetricAuditableEntitiesPut {
	this := EsgMetricAuditableEntitiesPut{}
	return &this
}

// NewEsgMetricAuditableEntitiesPutWithDefaults instantiates a new EsgMetricAuditableEntitiesPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsgMetricAuditableEntitiesPutWithDefaults() *EsgMetricAuditableEntitiesPut {
	this := EsgMetricAuditableEntitiesPut{}
	return &this
}

// GetEsgMetricAuditableEntity returns the EsgMetricAuditableEntity field value if set, zero value otherwise.
func (o *EsgMetricAuditableEntitiesPut) GetEsgMetricAuditableEntity() EsgMetricAuditableEntitiesPutEsgMetricAuditableEntity {
	if o == nil || IsNil(o.EsgMetricAuditableEntity) {
		var ret EsgMetricAuditableEntitiesPutEsgMetricAuditableEntity
		return ret
	}
	return *o.EsgMetricAuditableEntity
}

// GetEsgMetricAuditableEntityOk returns a tuple with the EsgMetricAuditableEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricAuditableEntitiesPut) GetEsgMetricAuditableEntityOk() (*EsgMetricAuditableEntitiesPutEsgMetricAuditableEntity, bool) {
	if o == nil || IsNil(o.EsgMetricAuditableEntity) {
		return nil, false
	}
	return o.EsgMetricAuditableEntity, true
}

// HasEsgMetricAuditableEntity returns a boolean if a field has been set.
func (o *EsgMetricAuditableEntitiesPut) HasEsgMetricAuditableEntity() bool {
	if o != nil && !IsNil(o.EsgMetricAuditableEntity) {
		return true
	}

	return false
}

// SetEsgMetricAuditableEntity gets a reference to the given EsgMetricAuditableEntitiesPutEsgMetricAuditableEntity and assigns it to the EsgMetricAuditableEntity field.
func (o *EsgMetricAuditableEntitiesPut) SetEsgMetricAuditableEntity(v EsgMetricAuditableEntitiesPutEsgMetricAuditableEntity) {
	o.EsgMetricAuditableEntity = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *EsgMetricAuditableEntitiesPut) GetPreviousValues() EsgMetricAuditableEntitiesPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret EsgMetricAuditableEntitiesPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricAuditableEntitiesPut) GetPreviousValuesOk() (*EsgMetricAuditableEntitiesPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *EsgMetricAuditableEntitiesPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given EsgMetricAuditableEntitiesPutPreviousValues and assigns it to the PreviousValues field.
func (o *EsgMetricAuditableEntitiesPut) SetPreviousValues(v EsgMetricAuditableEntitiesPutPreviousValues) {
	o.PreviousValues = &v
}

func (o EsgMetricAuditableEntitiesPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsgMetricAuditableEntitiesPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EsgMetricAuditableEntity) {
		toSerialize["esg_metric_auditable_entity"] = o.EsgMetricAuditableEntity
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableEsgMetricAuditableEntitiesPut struct {
	value *EsgMetricAuditableEntitiesPut
	isSet bool
}

func (v NullableEsgMetricAuditableEntitiesPut) Get() *EsgMetricAuditableEntitiesPut {
	return v.value
}

func (v *NullableEsgMetricAuditableEntitiesPut) Set(val *EsgMetricAuditableEntitiesPut) {
	v.value = val
	v.isSet = true
}

func (v NullableEsgMetricAuditableEntitiesPut) IsSet() bool {
	return v.isSet
}

func (v *NullableEsgMetricAuditableEntitiesPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsgMetricAuditableEntitiesPut(val *EsgMetricAuditableEntitiesPut) *NullableEsgMetricAuditableEntitiesPut {
	return &NullableEsgMetricAuditableEntitiesPut{value: val, isSet: true}
}

func (v NullableEsgMetricAuditableEntitiesPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsgMetricAuditableEntitiesPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


