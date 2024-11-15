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

// checks if the EsgMetricValueAggregatesPutPreviousValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsgMetricValueAggregatesPutPreviousValues{}

// EsgMetricValueAggregatesPutPreviousValues struct for EsgMetricValueAggregatesPutPreviousValues
type EsgMetricValueAggregatesPutPreviousValues struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	Value *float32 `json:"value,omitempty"`
	// Note: This is a Foreign Key to `esg_metric_unit_options.id`.<fk table='esg_metric_unit_options' column='id'/>
	EsgMetricUnitOptionId *int32 `json:"esg_metric_unit_option_id,omitempty"`
	// Note: This is a Foreign Key to `esg_metric_auditable_entities.id`.<fk table='esg_metric_auditable_entities' column='id'/>
	EsgMetricAuditableEntityId *int32 `json:"esg_metric_auditable_entity_id,omitempty"`
	AggregateYear *string `json:"aggregate_year,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	AggregationMethod *string `json:"aggregation_method,omitempty"`
}

// NewEsgMetricValueAggregatesPutPreviousValues instantiates a new EsgMetricValueAggregatesPutPreviousValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsgMetricValueAggregatesPutPreviousValues() *EsgMetricValueAggregatesPutPreviousValues {
	this := EsgMetricValueAggregatesPutPreviousValues{}
	return &this
}

// NewEsgMetricValueAggregatesPutPreviousValuesWithDefaults instantiates a new EsgMetricValueAggregatesPutPreviousValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsgMetricValueAggregatesPutPreviousValuesWithDefaults() *EsgMetricValueAggregatesPutPreviousValues {
	this := EsgMetricValueAggregatesPutPreviousValues{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EsgMetricValueAggregatesPutPreviousValues) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricValueAggregatesPutPreviousValues) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EsgMetricValueAggregatesPutPreviousValues) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EsgMetricValueAggregatesPutPreviousValues) SetId(v int32) {
	o.Id = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EsgMetricValueAggregatesPutPreviousValues) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricValueAggregatesPutPreviousValues) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EsgMetricValueAggregatesPutPreviousValues) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *EsgMetricValueAggregatesPutPreviousValues) SetValue(v float32) {
	o.Value = &v
}

// GetEsgMetricUnitOptionId returns the EsgMetricUnitOptionId field value if set, zero value otherwise.
func (o *EsgMetricValueAggregatesPutPreviousValues) GetEsgMetricUnitOptionId() int32 {
	if o == nil || IsNil(o.EsgMetricUnitOptionId) {
		var ret int32
		return ret
	}
	return *o.EsgMetricUnitOptionId
}

// GetEsgMetricUnitOptionIdOk returns a tuple with the EsgMetricUnitOptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricValueAggregatesPutPreviousValues) GetEsgMetricUnitOptionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EsgMetricUnitOptionId) {
		return nil, false
	}
	return o.EsgMetricUnitOptionId, true
}

// HasEsgMetricUnitOptionId returns a boolean if a field has been set.
func (o *EsgMetricValueAggregatesPutPreviousValues) HasEsgMetricUnitOptionId() bool {
	if o != nil && !IsNil(o.EsgMetricUnitOptionId) {
		return true
	}

	return false
}

// SetEsgMetricUnitOptionId gets a reference to the given int32 and assigns it to the EsgMetricUnitOptionId field.
func (o *EsgMetricValueAggregatesPutPreviousValues) SetEsgMetricUnitOptionId(v int32) {
	o.EsgMetricUnitOptionId = &v
}

// GetEsgMetricAuditableEntityId returns the EsgMetricAuditableEntityId field value if set, zero value otherwise.
func (o *EsgMetricValueAggregatesPutPreviousValues) GetEsgMetricAuditableEntityId() int32 {
	if o == nil || IsNil(o.EsgMetricAuditableEntityId) {
		var ret int32
		return ret
	}
	return *o.EsgMetricAuditableEntityId
}

// GetEsgMetricAuditableEntityIdOk returns a tuple with the EsgMetricAuditableEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricValueAggregatesPutPreviousValues) GetEsgMetricAuditableEntityIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EsgMetricAuditableEntityId) {
		return nil, false
	}
	return o.EsgMetricAuditableEntityId, true
}

// HasEsgMetricAuditableEntityId returns a boolean if a field has been set.
func (o *EsgMetricValueAggregatesPutPreviousValues) HasEsgMetricAuditableEntityId() bool {
	if o != nil && !IsNil(o.EsgMetricAuditableEntityId) {
		return true
	}

	return false
}

// SetEsgMetricAuditableEntityId gets a reference to the given int32 and assigns it to the EsgMetricAuditableEntityId field.
func (o *EsgMetricValueAggregatesPutPreviousValues) SetEsgMetricAuditableEntityId(v int32) {
	o.EsgMetricAuditableEntityId = &v
}

// GetAggregateYear returns the AggregateYear field value if set, zero value otherwise.
func (o *EsgMetricValueAggregatesPutPreviousValues) GetAggregateYear() string {
	if o == nil || IsNil(o.AggregateYear) {
		var ret string
		return ret
	}
	return *o.AggregateYear
}

// GetAggregateYearOk returns a tuple with the AggregateYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricValueAggregatesPutPreviousValues) GetAggregateYearOk() (*string, bool) {
	if o == nil || IsNil(o.AggregateYear) {
		return nil, false
	}
	return o.AggregateYear, true
}

// HasAggregateYear returns a boolean if a field has been set.
func (o *EsgMetricValueAggregatesPutPreviousValues) HasAggregateYear() bool {
	if o != nil && !IsNil(o.AggregateYear) {
		return true
	}

	return false
}

// SetAggregateYear gets a reference to the given string and assigns it to the AggregateYear field.
func (o *EsgMetricValueAggregatesPutPreviousValues) SetAggregateYear(v string) {
	o.AggregateYear = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EsgMetricValueAggregatesPutPreviousValues) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricValueAggregatesPutPreviousValues) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EsgMetricValueAggregatesPutPreviousValues) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *EsgMetricValueAggregatesPutPreviousValues) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EsgMetricValueAggregatesPutPreviousValues) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricValueAggregatesPutPreviousValues) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EsgMetricValueAggregatesPutPreviousValues) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *EsgMetricValueAggregatesPutPreviousValues) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *EsgMetricValueAggregatesPutPreviousValues) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricValueAggregatesPutPreviousValues) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *EsgMetricValueAggregatesPutPreviousValues) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *EsgMetricValueAggregatesPutPreviousValues) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetAggregationMethod returns the AggregationMethod field value if set, zero value otherwise.
func (o *EsgMetricValueAggregatesPutPreviousValues) GetAggregationMethod() string {
	if o == nil || IsNil(o.AggregationMethod) {
		var ret string
		return ret
	}
	return *o.AggregationMethod
}

// GetAggregationMethodOk returns a tuple with the AggregationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricValueAggregatesPutPreviousValues) GetAggregationMethodOk() (*string, bool) {
	if o == nil || IsNil(o.AggregationMethod) {
		return nil, false
	}
	return o.AggregationMethod, true
}

// HasAggregationMethod returns a boolean if a field has been set.
func (o *EsgMetricValueAggregatesPutPreviousValues) HasAggregationMethod() bool {
	if o != nil && !IsNil(o.AggregationMethod) {
		return true
	}

	return false
}

// SetAggregationMethod gets a reference to the given string and assigns it to the AggregationMethod field.
func (o *EsgMetricValueAggregatesPutPreviousValues) SetAggregationMethod(v string) {
	o.AggregationMethod = &v
}

func (o EsgMetricValueAggregatesPutPreviousValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsgMetricValueAggregatesPutPreviousValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.EsgMetricUnitOptionId) {
		toSerialize["esg_metric_unit_option_id"] = o.EsgMetricUnitOptionId
	}
	if !IsNil(o.EsgMetricAuditableEntityId) {
		toSerialize["esg_metric_auditable_entity_id"] = o.EsgMetricAuditableEntityId
	}
	if !IsNil(o.AggregateYear) {
		toSerialize["aggregate_year"] = o.AggregateYear
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.AggregationMethod) {
		toSerialize["aggregation_method"] = o.AggregationMethod
	}
	return toSerialize, nil
}

type NullableEsgMetricValueAggregatesPutPreviousValues struct {
	value *EsgMetricValueAggregatesPutPreviousValues
	isSet bool
}

func (v NullableEsgMetricValueAggregatesPutPreviousValues) Get() *EsgMetricValueAggregatesPutPreviousValues {
	return v.value
}

func (v *NullableEsgMetricValueAggregatesPutPreviousValues) Set(val *EsgMetricValueAggregatesPutPreviousValues) {
	v.value = val
	v.isSet = true
}

func (v NullableEsgMetricValueAggregatesPutPreviousValues) IsSet() bool {
	return v.isSet
}

func (v *NullableEsgMetricValueAggregatesPutPreviousValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsgMetricValueAggregatesPutPreviousValues(val *EsgMetricValueAggregatesPutPreviousValues) *NullableEsgMetricValueAggregatesPutPreviousValues {
	return &NullableEsgMetricValueAggregatesPutPreviousValues{value: val, isSet: true}
}

func (v NullableEsgMetricValueAggregatesPutPreviousValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsgMetricValueAggregatesPutPreviousValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


