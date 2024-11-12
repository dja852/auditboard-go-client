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

// checks if the RiskCategoriesPutPreviousValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskCategoriesPutPreviousValues{}

// RiskCategoriesPutPreviousValues struct for RiskCategoriesPutPreviousValues
type RiskCategoriesPutPreviousValues struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Name *string `json:"name,omitempty"`
	SortOrder *int32 `json:"sort_order,omitempty"`
	FormTemplateId *int32 `json:"form_template_id,omitempty"`
	IsDefault *bool `json:"is_default,omitempty"`
	Key *string `json:"key,omitempty"`
}

// NewRiskCategoriesPutPreviousValues instantiates a new RiskCategoriesPutPreviousValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskCategoriesPutPreviousValues() *RiskCategoriesPutPreviousValues {
	this := RiskCategoriesPutPreviousValues{}
	var sortOrder int32 = 0
	this.SortOrder = &sortOrder
	var isDefault bool = false
	this.IsDefault = &isDefault
	return &this
}

// NewRiskCategoriesPutPreviousValuesWithDefaults instantiates a new RiskCategoriesPutPreviousValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskCategoriesPutPreviousValuesWithDefaults() *RiskCategoriesPutPreviousValues {
	this := RiskCategoriesPutPreviousValues{}
	var sortOrder int32 = 0
	this.SortOrder = &sortOrder
	var isDefault bool = false
	this.IsDefault = &isDefault
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RiskCategoriesPutPreviousValues) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskCategoriesPutPreviousValues) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RiskCategoriesPutPreviousValues) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RiskCategoriesPutPreviousValues) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RiskCategoriesPutPreviousValues) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskCategoriesPutPreviousValues) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RiskCategoriesPutPreviousValues) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *RiskCategoriesPutPreviousValues) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RiskCategoriesPutPreviousValues) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskCategoriesPutPreviousValues) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RiskCategoriesPutPreviousValues) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *RiskCategoriesPutPreviousValues) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *RiskCategoriesPutPreviousValues) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskCategoriesPutPreviousValues) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *RiskCategoriesPutPreviousValues) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *RiskCategoriesPutPreviousValues) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RiskCategoriesPutPreviousValues) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskCategoriesPutPreviousValues) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RiskCategoriesPutPreviousValues) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RiskCategoriesPutPreviousValues) SetName(v string) {
	o.Name = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *RiskCategoriesPutPreviousValues) GetSortOrder() int32 {
	if o == nil || IsNil(o.SortOrder) {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskCategoriesPutPreviousValues) GetSortOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *RiskCategoriesPutPreviousValues) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *RiskCategoriesPutPreviousValues) SetSortOrder(v int32) {
	o.SortOrder = &v
}

// GetFormTemplateId returns the FormTemplateId field value if set, zero value otherwise.
func (o *RiskCategoriesPutPreviousValues) GetFormTemplateId() int32 {
	if o == nil || IsNil(o.FormTemplateId) {
		var ret int32
		return ret
	}
	return *o.FormTemplateId
}

// GetFormTemplateIdOk returns a tuple with the FormTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskCategoriesPutPreviousValues) GetFormTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FormTemplateId) {
		return nil, false
	}
	return o.FormTemplateId, true
}

// HasFormTemplateId returns a boolean if a field has been set.
func (o *RiskCategoriesPutPreviousValues) HasFormTemplateId() bool {
	if o != nil && !IsNil(o.FormTemplateId) {
		return true
	}

	return false
}

// SetFormTemplateId gets a reference to the given int32 and assigns it to the FormTemplateId field.
func (o *RiskCategoriesPutPreviousValues) SetFormTemplateId(v int32) {
	o.FormTemplateId = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *RiskCategoriesPutPreviousValues) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskCategoriesPutPreviousValues) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *RiskCategoriesPutPreviousValues) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *RiskCategoriesPutPreviousValues) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *RiskCategoriesPutPreviousValues) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskCategoriesPutPreviousValues) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *RiskCategoriesPutPreviousValues) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *RiskCategoriesPutPreviousValues) SetKey(v string) {
	o.Key = &v
}

func (o RiskCategoriesPutPreviousValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskCategoriesPutPreviousValues) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SortOrder) {
		toSerialize["sort_order"] = o.SortOrder
	}
	if !IsNil(o.FormTemplateId) {
		toSerialize["form_template_id"] = o.FormTemplateId
	}
	if !IsNil(o.IsDefault) {
		toSerialize["is_default"] = o.IsDefault
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	return toSerialize, nil
}

type NullableRiskCategoriesPutPreviousValues struct {
	value *RiskCategoriesPutPreviousValues
	isSet bool
}

func (v NullableRiskCategoriesPutPreviousValues) Get() *RiskCategoriesPutPreviousValues {
	return v.value
}

func (v *NullableRiskCategoriesPutPreviousValues) Set(val *RiskCategoriesPutPreviousValues) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskCategoriesPutPreviousValues) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskCategoriesPutPreviousValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskCategoriesPutPreviousValues(val *RiskCategoriesPutPreviousValues) *NullableRiskCategoriesPutPreviousValues {
	return &NullableRiskCategoriesPutPreviousValues{value: val, isSet: true}
}

func (v NullableRiskCategoriesPutPreviousValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskCategoriesPutPreviousValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


