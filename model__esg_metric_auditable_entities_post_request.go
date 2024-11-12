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

// checks if the EsgMetricAuditableEntitiesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsgMetricAuditableEntitiesPostRequest{}

// EsgMetricAuditableEntitiesPostRequest struct for EsgMetricAuditableEntitiesPostRequest
type EsgMetricAuditableEntitiesPostRequest struct {
	EsgMetricAuditableEntity *EsgMetricAuditableEntities `json:"esg_metric _auditable _entity,omitempty"`
}

// NewEsgMetricAuditableEntitiesPostRequest instantiates a new EsgMetricAuditableEntitiesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsgMetricAuditableEntitiesPostRequest() *EsgMetricAuditableEntitiesPostRequest {
	this := EsgMetricAuditableEntitiesPostRequest{}
	return &this
}

// NewEsgMetricAuditableEntitiesPostRequestWithDefaults instantiates a new EsgMetricAuditableEntitiesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsgMetricAuditableEntitiesPostRequestWithDefaults() *EsgMetricAuditableEntitiesPostRequest {
	this := EsgMetricAuditableEntitiesPostRequest{}
	return &this
}

// GetEsgMetricAuditableEntity returns the EsgMetricAuditableEntity field value if set, zero value otherwise.
func (o *EsgMetricAuditableEntitiesPostRequest) GetEsgMetricAuditableEntity() EsgMetricAuditableEntities {
	if o == nil || IsNil(o.EsgMetricAuditableEntity) {
		var ret EsgMetricAuditableEntities
		return ret
	}
	return *o.EsgMetricAuditableEntity
}

// GetEsgMetricAuditableEntityOk returns a tuple with the EsgMetricAuditableEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricAuditableEntitiesPostRequest) GetEsgMetricAuditableEntityOk() (*EsgMetricAuditableEntities, bool) {
	if o == nil || IsNil(o.EsgMetricAuditableEntity) {
		return nil, false
	}
	return o.EsgMetricAuditableEntity, true
}

// HasEsgMetricAuditableEntity returns a boolean if a field has been set.
func (o *EsgMetricAuditableEntitiesPostRequest) HasEsgMetricAuditableEntity() bool {
	if o != nil && !IsNil(o.EsgMetricAuditableEntity) {
		return true
	}

	return false
}

// SetEsgMetricAuditableEntity gets a reference to the given EsgMetricAuditableEntities and assigns it to the EsgMetricAuditableEntity field.
func (o *EsgMetricAuditableEntitiesPostRequest) SetEsgMetricAuditableEntity(v EsgMetricAuditableEntities) {
	o.EsgMetricAuditableEntity = &v
}

func (o EsgMetricAuditableEntitiesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsgMetricAuditableEntitiesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EsgMetricAuditableEntity) {
		toSerialize["esg_metric _auditable _entity"] = o.EsgMetricAuditableEntity
	}
	return toSerialize, nil
}

type NullableEsgMetricAuditableEntitiesPostRequest struct {
	value *EsgMetricAuditableEntitiesPostRequest
	isSet bool
}

func (v NullableEsgMetricAuditableEntitiesPostRequest) Get() *EsgMetricAuditableEntitiesPostRequest {
	return v.value
}

func (v *NullableEsgMetricAuditableEntitiesPostRequest) Set(val *EsgMetricAuditableEntitiesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEsgMetricAuditableEntitiesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEsgMetricAuditableEntitiesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsgMetricAuditableEntitiesPostRequest(val *EsgMetricAuditableEntitiesPostRequest) *NullableEsgMetricAuditableEntitiesPostRequest {
	return &NullableEsgMetricAuditableEntitiesPostRequest{value: val, isSet: true}
}

func (v NullableEsgMetricAuditableEntitiesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsgMetricAuditableEntitiesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


