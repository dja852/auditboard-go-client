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

// checks if the FormTemplates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormTemplates{}

// FormTemplates struct for FormTemplates
type FormTemplates struct {
	// Note: This is a Primary Key.<pk/>
	Id *int32 `json:"id,omitempty"`
	Name string `json:"name"`
	Type string `json:"type"`
	TemplateJson interface{} `json:"template_json,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}

type _FormTemplates FormTemplates

// NewFormTemplates instantiates a new FormTemplates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormTemplates(name string, type_ string) *FormTemplates {
	this := FormTemplates{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewFormTemplatesWithDefaults instantiates a new FormTemplates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormTemplatesWithDefaults() *FormTemplates {
	this := FormTemplates{}
	var name string = ""
	this.Name = name
	var type_ string = ""
	this.Type = type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FormTemplates) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormTemplates) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FormTemplates) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FormTemplates) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *FormTemplates) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormTemplates) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FormTemplates) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *FormTemplates) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FormTemplates) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FormTemplates) SetType(v string) {
	o.Type = v
}

// GetTemplateJson returns the TemplateJson field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FormTemplates) GetTemplateJson() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TemplateJson
}

// GetTemplateJsonOk returns a tuple with the TemplateJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormTemplates) GetTemplateJsonOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TemplateJson) {
		return nil, false
	}
	return &o.TemplateJson, true
}

// HasTemplateJson returns a boolean if a field has been set.
func (o *FormTemplates) HasTemplateJson() bool {
	if o != nil && !IsNil(o.TemplateJson) {
		return true
	}

	return false
}

// SetTemplateJson gets a reference to the given interface{} and assigns it to the TemplateJson field.
func (o *FormTemplates) SetTemplateJson(v interface{}) {
	o.TemplateJson = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FormTemplates) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormTemplates) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FormTemplates) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *FormTemplates) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FormTemplates) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormTemplates) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FormTemplates) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *FormTemplates) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *FormTemplates) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormTemplates) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *FormTemplates) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *FormTemplates) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

func (o FormTemplates) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormTemplates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if o.TemplateJson != nil {
		toSerialize["template_json"] = o.TemplateJson
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
	return toSerialize, nil
}

func (o *FormTemplates) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
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

	varFormTemplates := _FormTemplates{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFormTemplates)

	if err != nil {
		return err
	}

	*o = FormTemplates(varFormTemplates)

	return err
}

type NullableFormTemplates struct {
	value *FormTemplates
	isSet bool
}

func (v NullableFormTemplates) Get() *FormTemplates {
	return v.value
}

func (v *NullableFormTemplates) Set(val *FormTemplates) {
	v.value = val
	v.isSet = true
}

func (v NullableFormTemplates) IsSet() bool {
	return v.isSet
}

func (v *NullableFormTemplates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormTemplates(val *FormTemplates) *NullableFormTemplates {
	return &NullableFormTemplates{value: val, isSet: true}
}

func (v NullableFormTemplates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormTemplates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

