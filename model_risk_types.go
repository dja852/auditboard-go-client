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

// checks if the RiskTypes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskTypes{}

// RiskTypes struct for RiskTypes
type RiskTypes struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Name string `json:"name"`
	SortOrder int32 `json:"sort_order"`
	FormTemplateId *int32 `json:"form_template_id,omitempty"`
}

type _RiskTypes RiskTypes

// NewRiskTypes instantiates a new RiskTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskTypes(name string, sortOrder int32) *RiskTypes {
	this := RiskTypes{}
	this.Name = name
	this.SortOrder = sortOrder
	return &this
}

// NewRiskTypesWithDefaults instantiates a new RiskTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskTypesWithDefaults() *RiskTypes {
	this := RiskTypes{}
	var sortOrder int32 = 0
	this.SortOrder = sortOrder
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RiskTypes) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskTypes) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RiskTypes) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RiskTypes) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RiskTypes) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskTypes) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RiskTypes) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *RiskTypes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RiskTypes) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskTypes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RiskTypes) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *RiskTypes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *RiskTypes) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskTypes) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *RiskTypes) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *RiskTypes) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetName returns the Name field value
func (o *RiskTypes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RiskTypes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RiskTypes) SetName(v string) {
	o.Name = v
}

// GetSortOrder returns the SortOrder field value
func (o *RiskTypes) GetSortOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value
// and a boolean to check if the value has been set.
func (o *RiskTypes) GetSortOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortOrder, true
}

// SetSortOrder sets field value
func (o *RiskTypes) SetSortOrder(v int32) {
	o.SortOrder = v
}

// GetFormTemplateId returns the FormTemplateId field value if set, zero value otherwise.
func (o *RiskTypes) GetFormTemplateId() int32 {
	if o == nil || IsNil(o.FormTemplateId) {
		var ret int32
		return ret
	}
	return *o.FormTemplateId
}

// GetFormTemplateIdOk returns a tuple with the FormTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskTypes) GetFormTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FormTemplateId) {
		return nil, false
	}
	return o.FormTemplateId, true
}

// HasFormTemplateId returns a boolean if a field has been set.
func (o *RiskTypes) HasFormTemplateId() bool {
	if o != nil && !IsNil(o.FormTemplateId) {
		return true
	}

	return false
}

// SetFormTemplateId gets a reference to the given int32 and assigns it to the FormTemplateId field.
func (o *RiskTypes) SetFormTemplateId(v int32) {
	o.FormTemplateId = &v
}

func (o RiskTypes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskTypes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
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
	toSerialize["name"] = o.Name
	toSerialize["sort_order"] = o.SortOrder
	if !IsNil(o.FormTemplateId) {
		toSerialize["form_template_id"] = o.FormTemplateId
	}
	return toSerialize, nil
}

func (o *RiskTypes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"sort_order",
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

	varRiskTypes := _RiskTypes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRiskTypes)

	if err != nil {
		return err
	}

	*o = RiskTypes(varRiskTypes)

	return err
}

type NullableRiskTypes struct {
	value *RiskTypes
	isSet bool
}

func (v NullableRiskTypes) Get() *RiskTypes {
	return v.value
}

func (v *NullableRiskTypes) Set(val *RiskTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskTypes(val *RiskTypes) *NullableRiskTypes {
	return &NullableRiskTypes{value: val, isSet: true}
}

func (v NullableRiskTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


