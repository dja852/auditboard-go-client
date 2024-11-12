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

// checks if the AuditableEntityTypes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditableEntityTypes{}

// AuditableEntityTypes struct for AuditableEntityTypes
type AuditableEntityTypes struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	SortOrder *int32 `json:"sort_order,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Name string `json:"name"`
	FormTemplateId *int32 `json:"form_template_id,omitempty"`
	AllowedSections interface{} `json:"allowed_sections"`
	EnabledAttributes interface{} `json:"enabled_attributes,omitempty"`
	ExcludedAttributes interface{} `json:"excluded_attributes,omitempty"`
	Key *string `json:"key,omitempty"`
	// Note: This is a Foreign Key to `audit_survey_templates.id`.<fk table='audit_survey_templates' column='id'/>
	IntakeFormTemplateId *int32 `json:"intake_form_template_id,omitempty"`
	Level *int32 `json:"level,omitempty"`
	InventoryClass *string `json:"inventory_class,omitempty"`
}

type _AuditableEntityTypes AuditableEntityTypes

// NewAuditableEntityTypes instantiates a new AuditableEntityTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditableEntityTypes(name string, allowedSections interface{}) *AuditableEntityTypes {
	this := AuditableEntityTypes{}
	this.Name = name
	this.AllowedSections = allowedSections
	return &this
}

// NewAuditableEntityTypesWithDefaults instantiates a new AuditableEntityTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditableEntityTypesWithDefaults() *AuditableEntityTypes {
	this := AuditableEntityTypes{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuditableEntityTypes) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditableEntityTypes) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuditableEntityTypes) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AuditableEntityTypes) SetId(v int32) {
	o.Id = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *AuditableEntityTypes) GetSortOrder() int32 {
	if o == nil || IsNil(o.SortOrder) {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditableEntityTypes) GetSortOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *AuditableEntityTypes) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *AuditableEntityTypes) SetSortOrder(v int32) {
	o.SortOrder = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AuditableEntityTypes) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditableEntityTypes) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AuditableEntityTypes) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AuditableEntityTypes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AuditableEntityTypes) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditableEntityTypes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AuditableEntityTypes) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AuditableEntityTypes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *AuditableEntityTypes) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditableEntityTypes) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *AuditableEntityTypes) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *AuditableEntityTypes) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetName returns the Name field value
func (o *AuditableEntityTypes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuditableEntityTypes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuditableEntityTypes) SetName(v string) {
	o.Name = v
}

// GetFormTemplateId returns the FormTemplateId field value if set, zero value otherwise.
func (o *AuditableEntityTypes) GetFormTemplateId() int32 {
	if o == nil || IsNil(o.FormTemplateId) {
		var ret int32
		return ret
	}
	return *o.FormTemplateId
}

// GetFormTemplateIdOk returns a tuple with the FormTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditableEntityTypes) GetFormTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FormTemplateId) {
		return nil, false
	}
	return o.FormTemplateId, true
}

// HasFormTemplateId returns a boolean if a field has been set.
func (o *AuditableEntityTypes) HasFormTemplateId() bool {
	if o != nil && !IsNil(o.FormTemplateId) {
		return true
	}

	return false
}

// SetFormTemplateId gets a reference to the given int32 and assigns it to the FormTemplateId field.
func (o *AuditableEntityTypes) SetFormTemplateId(v int32) {
	o.FormTemplateId = &v
}

// GetAllowedSections returns the AllowedSections field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AuditableEntityTypes) GetAllowedSections() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.AllowedSections
}

// GetAllowedSectionsOk returns a tuple with the AllowedSections field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditableEntityTypes) GetAllowedSectionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AllowedSections) {
		return nil, false
	}
	return &o.AllowedSections, true
}

// SetAllowedSections sets field value
func (o *AuditableEntityTypes) SetAllowedSections(v interface{}) {
	o.AllowedSections = v
}

// GetEnabledAttributes returns the EnabledAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditableEntityTypes) GetEnabledAttributes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EnabledAttributes
}

// GetEnabledAttributesOk returns a tuple with the EnabledAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditableEntityTypes) GetEnabledAttributesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EnabledAttributes) {
		return nil, false
	}
	return &o.EnabledAttributes, true
}

// HasEnabledAttributes returns a boolean if a field has been set.
func (o *AuditableEntityTypes) HasEnabledAttributes() bool {
	if o != nil && !IsNil(o.EnabledAttributes) {
		return true
	}

	return false
}

// SetEnabledAttributes gets a reference to the given interface{} and assigns it to the EnabledAttributes field.
func (o *AuditableEntityTypes) SetEnabledAttributes(v interface{}) {
	o.EnabledAttributes = v
}

// GetExcludedAttributes returns the ExcludedAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditableEntityTypes) GetExcludedAttributes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExcludedAttributes
}

// GetExcludedAttributesOk returns a tuple with the ExcludedAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditableEntityTypes) GetExcludedAttributesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExcludedAttributes) {
		return nil, false
	}
	return &o.ExcludedAttributes, true
}

// HasExcludedAttributes returns a boolean if a field has been set.
func (o *AuditableEntityTypes) HasExcludedAttributes() bool {
	if o != nil && !IsNil(o.ExcludedAttributes) {
		return true
	}

	return false
}

// SetExcludedAttributes gets a reference to the given interface{} and assigns it to the ExcludedAttributes field.
func (o *AuditableEntityTypes) SetExcludedAttributes(v interface{}) {
	o.ExcludedAttributes = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AuditableEntityTypes) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditableEntityTypes) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AuditableEntityTypes) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AuditableEntityTypes) SetKey(v string) {
	o.Key = &v
}

// GetIntakeFormTemplateId returns the IntakeFormTemplateId field value if set, zero value otherwise.
func (o *AuditableEntityTypes) GetIntakeFormTemplateId() int32 {
	if o == nil || IsNil(o.IntakeFormTemplateId) {
		var ret int32
		return ret
	}
	return *o.IntakeFormTemplateId
}

// GetIntakeFormTemplateIdOk returns a tuple with the IntakeFormTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditableEntityTypes) GetIntakeFormTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.IntakeFormTemplateId) {
		return nil, false
	}
	return o.IntakeFormTemplateId, true
}

// HasIntakeFormTemplateId returns a boolean if a field has been set.
func (o *AuditableEntityTypes) HasIntakeFormTemplateId() bool {
	if o != nil && !IsNil(o.IntakeFormTemplateId) {
		return true
	}

	return false
}

// SetIntakeFormTemplateId gets a reference to the given int32 and assigns it to the IntakeFormTemplateId field.
func (o *AuditableEntityTypes) SetIntakeFormTemplateId(v int32) {
	o.IntakeFormTemplateId = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *AuditableEntityTypes) GetLevel() int32 {
	if o == nil || IsNil(o.Level) {
		var ret int32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditableEntityTypes) GetLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *AuditableEntityTypes) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given int32 and assigns it to the Level field.
func (o *AuditableEntityTypes) SetLevel(v int32) {
	o.Level = &v
}

// GetInventoryClass returns the InventoryClass field value if set, zero value otherwise.
func (o *AuditableEntityTypes) GetInventoryClass() string {
	if o == nil || IsNil(o.InventoryClass) {
		var ret string
		return ret
	}
	return *o.InventoryClass
}

// GetInventoryClassOk returns a tuple with the InventoryClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditableEntityTypes) GetInventoryClassOk() (*string, bool) {
	if o == nil || IsNil(o.InventoryClass) {
		return nil, false
	}
	return o.InventoryClass, true
}

// HasInventoryClass returns a boolean if a field has been set.
func (o *AuditableEntityTypes) HasInventoryClass() bool {
	if o != nil && !IsNil(o.InventoryClass) {
		return true
	}

	return false
}

// SetInventoryClass gets a reference to the given string and assigns it to the InventoryClass field.
func (o *AuditableEntityTypes) SetInventoryClass(v string) {
	o.InventoryClass = &v
}

func (o AuditableEntityTypes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditableEntityTypes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SortOrder) {
		toSerialize["sort_order"] = o.SortOrder
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
	if !IsNil(o.FormTemplateId) {
		toSerialize["form_template_id"] = o.FormTemplateId
	}
	if o.AllowedSections != nil {
		toSerialize["allowed_sections"] = o.AllowedSections
	}
	if o.EnabledAttributes != nil {
		toSerialize["enabled_attributes"] = o.EnabledAttributes
	}
	if o.ExcludedAttributes != nil {
		toSerialize["excluded_attributes"] = o.ExcludedAttributes
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.IntakeFormTemplateId) {
		toSerialize["intake_form_template_id"] = o.IntakeFormTemplateId
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.InventoryClass) {
		toSerialize["inventory_class"] = o.InventoryClass
	}
	return toSerialize, nil
}

func (o *AuditableEntityTypes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"allowed_sections",
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

	varAuditableEntityTypes := _AuditableEntityTypes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuditableEntityTypes)

	if err != nil {
		return err
	}

	*o = AuditableEntityTypes(varAuditableEntityTypes)

	return err
}

type NullableAuditableEntityTypes struct {
	value *AuditableEntityTypes
	isSet bool
}

func (v NullableAuditableEntityTypes) Get() *AuditableEntityTypes {
	return v.value
}

func (v *NullableAuditableEntityTypes) Set(val *AuditableEntityTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditableEntityTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditableEntityTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditableEntityTypes(val *AuditableEntityTypes) *NullableAuditableEntityTypes {
	return &NullableAuditableEntityTypes{value: val, isSet: true}
}

func (v NullableAuditableEntityTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditableEntityTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


