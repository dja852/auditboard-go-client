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

// checks if the ControlsSegments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControlsSegments{}

// ControlsSegments struct for ControlsSegments
type ControlsSegments struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	Name string `json:"name"`
	Type *string `json:"type,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	// Note: This is a Foreign Key to `form_templates.id`.<fk table='form_templates' column='id'/>
	FormTemplateId *int32 `json:"form_template_id,omitempty"`
}

type _ControlsSegments ControlsSegments

// NewControlsSegments instantiates a new ControlsSegments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControlsSegments(name string) *ControlsSegments {
	this := ControlsSegments{}
	this.Name = name
	return &this
}

// NewControlsSegmentsWithDefaults instantiates a new ControlsSegments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControlsSegmentsWithDefaults() *ControlsSegments {
	this := ControlsSegments{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ControlsSegments) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsSegments) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ControlsSegments) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ControlsSegments) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *ControlsSegments) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ControlsSegments) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ControlsSegments) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ControlsSegments) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsSegments) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ControlsSegments) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ControlsSegments) SetType(v string) {
	o.Type = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ControlsSegments) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsSegments) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ControlsSegments) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ControlsSegments) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ControlsSegments) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsSegments) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ControlsSegments) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ControlsSegments) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ControlsSegments) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsSegments) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ControlsSegments) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *ControlsSegments) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetFormTemplateId returns the FormTemplateId field value if set, zero value otherwise.
func (o *ControlsSegments) GetFormTemplateId() int32 {
	if o == nil || IsNil(o.FormTemplateId) {
		var ret int32
		return ret
	}
	return *o.FormTemplateId
}

// GetFormTemplateIdOk returns a tuple with the FormTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsSegments) GetFormTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FormTemplateId) {
		return nil, false
	}
	return o.FormTemplateId, true
}

// HasFormTemplateId returns a boolean if a field has been set.
func (o *ControlsSegments) HasFormTemplateId() bool {
	if o != nil && !IsNil(o.FormTemplateId) {
		return true
	}

	return false
}

// SetFormTemplateId gets a reference to the given int32 and assigns it to the FormTemplateId field.
func (o *ControlsSegments) SetFormTemplateId(v int32) {
	o.FormTemplateId = &v
}

func (o ControlsSegments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControlsSegments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
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
	if !IsNil(o.FormTemplateId) {
		toSerialize["form_template_id"] = o.FormTemplateId
	}
	return toSerialize, nil
}

func (o *ControlsSegments) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varControlsSegments := _ControlsSegments{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varControlsSegments)

	if err != nil {
		return err
	}

	*o = ControlsSegments(varControlsSegments)

	return err
}

type NullableControlsSegments struct {
	value *ControlsSegments
	isSet bool
}

func (v NullableControlsSegments) Get() *ControlsSegments {
	return v.value
}

func (v *NullableControlsSegments) Set(val *ControlsSegments) {
	v.value = val
	v.isSet = true
}

func (v NullableControlsSegments) IsSet() bool {
	return v.isSet
}

func (v *NullableControlsSegments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControlsSegments(val *ControlsSegments) *NullableControlsSegments {
	return &NullableControlsSegments{value: val, isSet: true}
}

func (v NullableControlsSegments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControlsSegments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

