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

// checks if the FormTemplatesPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormTemplatesPut{}

// FormTemplatesPut struct for FormTemplatesPut
type FormTemplatesPut struct {
	FormTemplate *FormTemplatesPutFormTemplate `json:"form_template,omitempty"`
	PreviousValues *FormTemplatesPutPreviousValues `json:"previous_values,omitempty"`
}

// NewFormTemplatesPut instantiates a new FormTemplatesPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormTemplatesPut() *FormTemplatesPut {
	this := FormTemplatesPut{}
	return &this
}

// NewFormTemplatesPutWithDefaults instantiates a new FormTemplatesPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormTemplatesPutWithDefaults() *FormTemplatesPut {
	this := FormTemplatesPut{}
	return &this
}

// GetFormTemplate returns the FormTemplate field value if set, zero value otherwise.
func (o *FormTemplatesPut) GetFormTemplate() FormTemplatesPutFormTemplate {
	if o == nil || IsNil(o.FormTemplate) {
		var ret FormTemplatesPutFormTemplate
		return ret
	}
	return *o.FormTemplate
}

// GetFormTemplateOk returns a tuple with the FormTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormTemplatesPut) GetFormTemplateOk() (*FormTemplatesPutFormTemplate, bool) {
	if o == nil || IsNil(o.FormTemplate) {
		return nil, false
	}
	return o.FormTemplate, true
}

// HasFormTemplate returns a boolean if a field has been set.
func (o *FormTemplatesPut) HasFormTemplate() bool {
	if o != nil && !IsNil(o.FormTemplate) {
		return true
	}

	return false
}

// SetFormTemplate gets a reference to the given FormTemplatesPutFormTemplate and assigns it to the FormTemplate field.
func (o *FormTemplatesPut) SetFormTemplate(v FormTemplatesPutFormTemplate) {
	o.FormTemplate = &v
}

// GetPreviousValues returns the PreviousValues field value if set, zero value otherwise.
func (o *FormTemplatesPut) GetPreviousValues() FormTemplatesPutPreviousValues {
	if o == nil || IsNil(o.PreviousValues) {
		var ret FormTemplatesPutPreviousValues
		return ret
	}
	return *o.PreviousValues
}

// GetPreviousValuesOk returns a tuple with the PreviousValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormTemplatesPut) GetPreviousValuesOk() (*FormTemplatesPutPreviousValues, bool) {
	if o == nil || IsNil(o.PreviousValues) {
		return nil, false
	}
	return o.PreviousValues, true
}

// HasPreviousValues returns a boolean if a field has been set.
func (o *FormTemplatesPut) HasPreviousValues() bool {
	if o != nil && !IsNil(o.PreviousValues) {
		return true
	}

	return false
}

// SetPreviousValues gets a reference to the given FormTemplatesPutPreviousValues and assigns it to the PreviousValues field.
func (o *FormTemplatesPut) SetPreviousValues(v FormTemplatesPutPreviousValues) {
	o.PreviousValues = &v
}

func (o FormTemplatesPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormTemplatesPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FormTemplate) {
		toSerialize["form_template"] = o.FormTemplate
	}
	if !IsNil(o.PreviousValues) {
		toSerialize["previous_values"] = o.PreviousValues
	}
	return toSerialize, nil
}

type NullableFormTemplatesPut struct {
	value *FormTemplatesPut
	isSet bool
}

func (v NullableFormTemplatesPut) Get() *FormTemplatesPut {
	return v.value
}

func (v *NullableFormTemplatesPut) Set(val *FormTemplatesPut) {
	v.value = val
	v.isSet = true
}

func (v NullableFormTemplatesPut) IsSet() bool {
	return v.isSet
}

func (v *NullableFormTemplatesPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormTemplatesPut(val *FormTemplatesPut) *NullableFormTemplatesPut {
	return &NullableFormTemplatesPut{value: val, isSet: true}
}

func (v NullableFormTemplatesPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormTemplatesPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


