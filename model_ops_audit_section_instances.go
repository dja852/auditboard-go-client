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

// checks if the OpsAuditSectionInstances type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpsAuditSectionInstances{}

// OpsAuditSectionInstances struct for OpsAuditSectionInstances
type OpsAuditSectionInstances struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	SortOrder int32 `json:"sort_order"`
	// Note: This is a Foreign Key to `ops_audits.id`.<fk table='ops_audits' column='id'/>
	OpsAuditId int32 `json:"ops_audit_id"`
	// Note: This is a Foreign Key to `ops_audit_sections.id`.<fk table='ops_audit_sections' column='id'/>
	OpsAuditSectionId int32 `json:"ops_audit_section_id"`
	FormTemplateId *int32 `json:"form_template_id,omitempty"`
}

type _OpsAuditSectionInstances OpsAuditSectionInstances

// NewOpsAuditSectionInstances instantiates a new OpsAuditSectionInstances object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsAuditSectionInstances(sortOrder int32, opsAuditId int32, opsAuditSectionId int32) *OpsAuditSectionInstances {
	this := OpsAuditSectionInstances{}
	this.SortOrder = sortOrder
	this.OpsAuditId = opsAuditId
	this.OpsAuditSectionId = opsAuditSectionId
	var formTemplateId int32 = 1
	this.FormTemplateId = &formTemplateId
	return &this
}

// NewOpsAuditSectionInstancesWithDefaults instantiates a new OpsAuditSectionInstances object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsAuditSectionInstancesWithDefaults() *OpsAuditSectionInstances {
	this := OpsAuditSectionInstances{}
	var sortOrder int32 = 0
	this.SortOrder = sortOrder
	var formTemplateId int32 = 1
	this.FormTemplateId = &formTemplateId
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OpsAuditSectionInstances) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditSectionInstances) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OpsAuditSectionInstances) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OpsAuditSectionInstances) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OpsAuditSectionInstances) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditSectionInstances) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OpsAuditSectionInstances) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *OpsAuditSectionInstances) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OpsAuditSectionInstances) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditSectionInstances) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OpsAuditSectionInstances) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *OpsAuditSectionInstances) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *OpsAuditSectionInstances) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditSectionInstances) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *OpsAuditSectionInstances) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *OpsAuditSectionInstances) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetSortOrder returns the SortOrder field value
func (o *OpsAuditSectionInstances) GetSortOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value
// and a boolean to check if the value has been set.
func (o *OpsAuditSectionInstances) GetSortOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortOrder, true
}

// SetSortOrder sets field value
func (o *OpsAuditSectionInstances) SetSortOrder(v int32) {
	o.SortOrder = v
}

// GetOpsAuditId returns the OpsAuditId field value
func (o *OpsAuditSectionInstances) GetOpsAuditId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OpsAuditId
}

// GetOpsAuditIdOk returns a tuple with the OpsAuditId field value
// and a boolean to check if the value has been set.
func (o *OpsAuditSectionInstances) GetOpsAuditIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpsAuditId, true
}

// SetOpsAuditId sets field value
func (o *OpsAuditSectionInstances) SetOpsAuditId(v int32) {
	o.OpsAuditId = v
}

// GetOpsAuditSectionId returns the OpsAuditSectionId field value
func (o *OpsAuditSectionInstances) GetOpsAuditSectionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OpsAuditSectionId
}

// GetOpsAuditSectionIdOk returns a tuple with the OpsAuditSectionId field value
// and a boolean to check if the value has been set.
func (o *OpsAuditSectionInstances) GetOpsAuditSectionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpsAuditSectionId, true
}

// SetOpsAuditSectionId sets field value
func (o *OpsAuditSectionInstances) SetOpsAuditSectionId(v int32) {
	o.OpsAuditSectionId = v
}

// GetFormTemplateId returns the FormTemplateId field value if set, zero value otherwise.
func (o *OpsAuditSectionInstances) GetFormTemplateId() int32 {
	if o == nil || IsNil(o.FormTemplateId) {
		var ret int32
		return ret
	}
	return *o.FormTemplateId
}

// GetFormTemplateIdOk returns a tuple with the FormTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsAuditSectionInstances) GetFormTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FormTemplateId) {
		return nil, false
	}
	return o.FormTemplateId, true
}

// HasFormTemplateId returns a boolean if a field has been set.
func (o *OpsAuditSectionInstances) HasFormTemplateId() bool {
	if o != nil && !IsNil(o.FormTemplateId) {
		return true
	}

	return false
}

// SetFormTemplateId gets a reference to the given int32 and assigns it to the FormTemplateId field.
func (o *OpsAuditSectionInstances) SetFormTemplateId(v int32) {
	o.FormTemplateId = &v
}

func (o OpsAuditSectionInstances) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpsAuditSectionInstances) ToMap() (map[string]interface{}, error) {
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
	toSerialize["sort_order"] = o.SortOrder
	toSerialize["ops_audit_id"] = o.OpsAuditId
	toSerialize["ops_audit_section_id"] = o.OpsAuditSectionId
	if !IsNil(o.FormTemplateId) {
		toSerialize["form_template_id"] = o.FormTemplateId
	}
	return toSerialize, nil
}

func (o *OpsAuditSectionInstances) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sort_order",
		"ops_audit_id",
		"ops_audit_section_id",
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

	varOpsAuditSectionInstances := _OpsAuditSectionInstances{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOpsAuditSectionInstances)

	if err != nil {
		return err
	}

	*o = OpsAuditSectionInstances(varOpsAuditSectionInstances)

	return err
}

type NullableOpsAuditSectionInstances struct {
	value *OpsAuditSectionInstances
	isSet bool
}

func (v NullableOpsAuditSectionInstances) Get() *OpsAuditSectionInstances {
	return v.value
}

func (v *NullableOpsAuditSectionInstances) Set(val *OpsAuditSectionInstances) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsAuditSectionInstances) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsAuditSectionInstances) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsAuditSectionInstances(val *OpsAuditSectionInstances) *NullableOpsAuditSectionInstances {
	return &NullableOpsAuditSectionInstances{value: val, isSet: true}
}

func (v NullableOpsAuditSectionInstances) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsAuditSectionInstances) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

