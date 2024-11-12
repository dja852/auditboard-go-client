/*
AuditBoard Developer Portal API Documentation

External API reference documentation.

API version: 23.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auditboard

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EsgMetricValueAggregatesPutEsgMetricValueAggregate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsgMetricValueAggregatesPutEsgMetricValueAggregate{}

// EsgMetricValueAggregatesPutEsgMetricValueAggregate struct for EsgMetricValueAggregatesPutEsgMetricValueAggregate
type EsgMetricValueAggregatesPutEsgMetricValueAggregate struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	Value *float32 `json:"value,omitempty"`
	// Note: This is a Foreign Key to `esg_metric_unit_options.id`.<fk table='esg_metric_unit_options' column='id'/>
	EsgMetricUnitOptionId *int32 `json:"esg_metric_unit_option_id,omitempty"`
	// Note: This is a Foreign Key to `esg_metric_auditable_entities.id`.<fk table='esg_metric_auditable_entities' column='id'/>
	EsgMetricAuditableEntityId int32 `json:"esg_metric_auditable_entity_id"`
	AggregateYear string `json:"aggregate_year"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	AggregationMethod *string `json:"aggregation_method,omitempty"`
}

type _EsgMetricValueAggregatesPutEsgMetricValueAggregate EsgMetricValueAggregatesPutEsgMetricValueAggregate

// NewEsgMetricValueAggregatesPutEsgMetricValueAggregate instantiates a new EsgMetricValueAggregatesPutEsgMetricValueAggregate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsgMetricValueAggregatesPutEsgMetricValueAggregate(esgMetricAuditableEntityId int32, aggregateYear string) *EsgMetricValueAggregatesPutEsgMetricValueAggregate {
	this := EsgMetricValueAggregatesPutEsgMetricValueAggregate{}
	this.EsgMetricAuditableEntityId = esgMetricAuditableEntityId
	this.AggregateYear = aggregateYear
	return &this
}

// NewEsgMetricValueAggregatesPutEsgMetricValueAggregateWithDefaults instantiates a new EsgMetricValueAggregatesPutEsgMetricValueAggregate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsgMetricValueAggregatesPutEsgMetricValueAggregateWithDefaults() *EsgMetricValueAggregatesPutEsgMetricValueAggregate {
	this := EsgMetricValueAggregatesPutEsgMetricValueAggregate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) SetId(v int32) {
	o.Id = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) SetValue(v float32) {
	o.Value = &v
}

// GetEsgMetricUnitOptionId returns the EsgMetricUnitOptionId field value if set, zero value otherwise.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) GetEsgMetricUnitOptionId() int32 {
	if o == nil || IsNil(o.EsgMetricUnitOptionId) {
		var ret int32
		return ret
	}
	return *o.EsgMetricUnitOptionId
}

// GetEsgMetricUnitOptionIdOk returns a tuple with the EsgMetricUnitOptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) GetEsgMetricUnitOptionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EsgMetricUnitOptionId) {
		return nil, false
	}
	return o.EsgMetricUnitOptionId, true
}

// HasEsgMetricUnitOptionId returns a boolean if a field has been set.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) HasEsgMetricUnitOptionId() bool {
	if o != nil && !IsNil(o.EsgMetricUnitOptionId) {
		return true
	}

	return false
}

// SetEsgMetricUnitOptionId gets a reference to the given int32 and assigns it to the EsgMetricUnitOptionId field.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) SetEsgMetricUnitOptionId(v int32) {
	o.EsgMetricUnitOptionId = &v
}

// GetEsgMetricAuditableEntityId returns the EsgMetricAuditableEntityId field value
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) GetEsgMetricAuditableEntityId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EsgMetricAuditableEntityId
}

// GetEsgMetricAuditableEntityIdOk returns a tuple with the EsgMetricAuditableEntityId field value
// and a boolean to check if the value has been set.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) GetEsgMetricAuditableEntityIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EsgMetricAuditableEntityId, true
}

// SetEsgMetricAuditableEntityId sets field value
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) SetEsgMetricAuditableEntityId(v int32) {
	o.EsgMetricAuditableEntityId = v
}

// GetAggregateYear returns the AggregateYear field value
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) GetAggregateYear() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AggregateYear
}

// GetAggregateYearOk returns a tuple with the AggregateYear field value
// and a boolean to check if the value has been set.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) GetAggregateYearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AggregateYear, true
}

// SetAggregateYear sets field value
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) SetAggregateYear(v string) {
	o.AggregateYear = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetAggregationMethod returns the AggregationMethod field value if set, zero value otherwise.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) GetAggregationMethod() string {
	if o == nil || IsNil(o.AggregationMethod) {
		var ret string
		return ret
	}
	return *o.AggregationMethod
}

// GetAggregationMethodOk returns a tuple with the AggregationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) GetAggregationMethodOk() (*string, bool) {
	if o == nil || IsNil(o.AggregationMethod) {
		return nil, false
	}
	return o.AggregationMethod, true
}

// HasAggregationMethod returns a boolean if a field has been set.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) HasAggregationMethod() bool {
	if o != nil && !IsNil(o.AggregationMethod) {
		return true
	}

	return false
}

// SetAggregationMethod gets a reference to the given string and assigns it to the AggregationMethod field.
func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) SetAggregationMethod(v string) {
	o.AggregationMethod = &v
}

func (o EsgMetricValueAggregatesPutEsgMetricValueAggregate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsgMetricValueAggregatesPutEsgMetricValueAggregate) ToMap() (map[string]interface{}, error) {
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
	toSerialize["esg_metric_auditable_entity_id"] = o.EsgMetricAuditableEntityId
	toSerialize["aggregate_year"] = o.AggregateYear
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

func (o *EsgMetricValueAggregatesPutEsgMetricValueAggregate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"esg_metric_auditable_entity_id",
		"aggregate_year",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEsgMetricValueAggregatesPutEsgMetricValueAggregate := _EsgMetricValueAggregatesPutEsgMetricValueAggregate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEsgMetricValueAggregatesPutEsgMetricValueAggregate)

	if err != nil {
		return err
	}

	*o = EsgMetricValueAggregatesPutEsgMetricValueAggregate(varEsgMetricValueAggregatesPutEsgMetricValueAggregate)

	return err
}

type NullableEsgMetricValueAggregatesPutEsgMetricValueAggregate struct {
	value *EsgMetricValueAggregatesPutEsgMetricValueAggregate
	isSet bool
}

func (v NullableEsgMetricValueAggregatesPutEsgMetricValueAggregate) Get() *EsgMetricValueAggregatesPutEsgMetricValueAggregate {
	return v.value
}

func (v *NullableEsgMetricValueAggregatesPutEsgMetricValueAggregate) Set(val *EsgMetricValueAggregatesPutEsgMetricValueAggregate) {
	v.value = val
	v.isSet = true
}

func (v NullableEsgMetricValueAggregatesPutEsgMetricValueAggregate) IsSet() bool {
	return v.isSet
}

func (v *NullableEsgMetricValueAggregatesPutEsgMetricValueAggregate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsgMetricValueAggregatesPutEsgMetricValueAggregate(val *EsgMetricValueAggregatesPutEsgMetricValueAggregate) *NullableEsgMetricValueAggregatesPutEsgMetricValueAggregate {
	return &NullableEsgMetricValueAggregatesPutEsgMetricValueAggregate{value: val, isSet: true}
}

func (v NullableEsgMetricValueAggregatesPutEsgMetricValueAggregate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsgMetricValueAggregatesPutEsgMetricValueAggregate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


